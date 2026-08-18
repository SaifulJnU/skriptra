# Skriptra on AWS — Production Deployment

**Scope.** Running Skriptra on AWS with Claude as the LLM: React frontend, Go API
and worker, PostgreSQL 17 + pgvector, NATS JetStream, Redis, and the OCR sidecar.
Target region in the examples is `eu-central-1` (Frankfurt), because the corpus is
German university material and the data should not leave the EU.

This document is the deployment counterpart to [`00-DESIGN.md`](00-DESIGN.md).
It does not repeat the architecture rationale; it says what to build, in what
order, and what has to change in the code first.

---

## 0. Read this first: two things that are not what they look like

### 0.1 A Claude Pro subscription does not give you API access

Claude Pro is a consumer subscription for claude.ai and the Claude desktop and
mobile apps. It carries **no API key and no programmatic quota**. An application
that calls Claude needs one of the three access paths below, each billed
separately from Pro, and none of which your Pro subscription discounts.

| Access path | What it is | Auth | Model ID form | Use when |
|---|---|---|---|---|
| **Anthropic API** (`api.anthropic.com`) | First-party. Every feature, same day. | `x-api-key` from [console.anthropic.com](https://console.anthropic.com) | `claude-opus-5` | Simplest. Fine if an egress call to Anthropic is acceptable. |
| **Amazon Bedrock** | AWS-operated. Partner service, feature subset, separate pricing. | AWS SigV4 / IAM role | `anthropic.claude-opus-5` | You want IAM instead of a key, PrivateLink, AWS Marketplace billing, and one bill. |
| **Claude Platform on AWS** | Anthropic-operated, reached through AWS infrastructure. Same-day API parity. | AWS SigV4 / IAM role | `claude-opus-5` (bare) | You want AWS-native IAM and billing **and** full first-party feature parity. |

**Recommendation for this project: Amazon Bedrock.** The workload uses only
plain text generation — no server-side tools, no Files API, no batches — so the
Bedrock feature subset costs you nothing, and in exchange you get an IAM task
role instead of a long-lived API key in Secrets Manager, a VPC interface
endpoint so inference traffic never touches the public internet, and one AWS
invoice. Verify Claude model availability in `eu-central-1` before you commit; if
the model you want is not there, either use an EU region that has it or fall back
to the first-party API with `LLM_API_KEY` in Secrets Manager.

### 0.2 Anthropic does not sell an embeddings API

Claude generates text. It does not produce embeddings, and there is no
`/v1/embeddings` endpoint. Skriptra needs embeddings for `chunks.embedding` and
`question_embeddings.embedding`, so **the embedder is a separate decision from
the LLM** and always a different vendor.

That decision is constrained by the schema. `migrations/000001_init.up.sql` pins
both vector columns to `vector(768)`, with a `CHECK (embedding_dim = 768)`
constraint, and `config.validate()` refuses to start when
`EMBEDDING_DIMENSIONS != 768`. Most hosted embedding models emit 1024 or 1536 —
but the Matryoshka models let you *request* an output width, and 768 is a
supported value on two of them.

| Option | Dimensions | Schema change? | Adapter work | Notes |
|---|---|---|---|---|
| **OpenAI `text-embedding-3-small`** | **768 on request** | **None** | one line (below) | Cheapest hosted path, exact dimension match. |
| **Ollama `nomic-embed-text`, self-hosted on ECS** | 768 | **None** | none | Same model as local dev. CPU-only, ~275M params, no per-token cost, no egress. |
| Google `gemini-embedding-001` | 768 on request (`outputDimensionality`) | None | new adapter | Strong multilingual. Not OpenAI-compatible. |
| Bedrock Cohere Embed Multilingual v3 (`cohere.embed-multilingual-v3`) | 1024 | Yes | SigV4 adapter | Best quality on a mixed German/English corpus. |
| Bedrock Titan Text Embeddings V2 (`amazon.titan-embed-text-v2:0`) | 1024 / 512 / 256 | Yes | SigV4 adapter | Same IAM path as the LLM. English-leaning. No 768. |
| Voyage AI (`voyage-3.5`) | 2048 / 1024 / 512 / 256 | Yes | new adapter | Anthropic's recommended embedding partner. No 768. |

**Decision: OpenAI `text-embedding-3-small` at 768 dimensions.** It matches the
schema exactly, needs no migration and no re-embed, and removes a container from
the deployment. The enabling change is already made: `embedRequest` in
`internal/provider/openaicompat/openaicompat.go` now carries a `dimensions`
field and forwards `settings.Dimensions`, so `EMBEDDING_DIMENSIONS=768` is a
request parameter rather than a wish. Everything else about the adapter is
unchanged — this runs through the existing `openai-compatible` path.

```
EMBEDDING_PROVIDER=openai
EMBEDDING_MODEL=text-embedding-3-small
EMBEDDING_BASE_URL=https://api.openai.com/v1
EMBEDDING_API_KEY=<from Secrets Manager>
EMBEDDING_DIMENSIONS=768
```

Two consequences worth being explicit about:

- **`dimensions` is now sent to every OpenAI-compatible embedding endpoint.** A
  self-hosted server that does not implement the parameter (some vLLM and TGI
  builds) will reject the request. That is not a regression: without the
  parameter such a server returns its native width and the existing check
  rejects the batch anyway. The failure just moved earlier and got a better
  error message. Local development is unaffected — Ollama has its own adapter.
- **Embedding text now leaves your infrastructure.** Question and chunk text is
  sent to OpenAI at ingest, and every user question is sent at query time. If
  the institution's position is that course material must not reach a US
  processor, this is the decision to revisit, and `nomic-embed-text` on a
  Fargate task is the drop-in that undoes it — same 768 dimensions, same schema,
  no re-embed.

Cohere Embed Multilingual v3 remains the quality play for a German corpus, but
it costs a dimension migration (§3.3). Measure with the eval harness before
paying for it.

---

## 1. Target architecture

```
                        Route 53  ─────────────────────────────────┐
                            │                                      │
                    CloudFront (OAC)                          ACM cert
                       /          \
          S3 (static SPA)      ALB  (HTTPS, WAF)
                                 │            public subnets, 3 AZ
        ─────────────────────────┼────────────────────────────────────────
                                 │            private subnets, 3 AZ
                          ┌──────┴───────┐
                          │  ECS Fargate │
                          ├──────────────┤
                          │ api      x N │──┐
                          │ worker   x M │──┤
                          │ ocr      x 1 │  │
                          │ nats     x 1 │◄─┘   (JetStream, EFS-backed)
                          └──────┬───────┘
                                 │
     ┌───────────────┬───────────┼─────────────┬──────────────────┐
     │               │           │             │                  │
  RDS Postgres   ElastiCache    S3          Secrets          VPC endpoint
  17 + pgvector   (Valkey)   (uploads)      Manager        → Bedrock / S3 / ECR
  Multi-AZ                                                    (no NAT egress)
```

Every service boundary here is the one that already exists in
`docker-compose.yml`. Only the infrastructure differs: managed Postgres instead
of a container, a load balancer in front of the API, replicas of the workers.

### 1.1 Component mapping

| Compose service | AWS | Sizing (start here) |
|---|---|---|
| `web` | S3 + CloudFront (static Vite build) | — |
| `api` | ECS Fargate service behind ALB | 0.5 vCPU / 1 GB, 2 tasks minimum |
| `worker` | ECS Fargate service, no ingress | 1 vCPU / 2 GB, 1 task, scales on queue depth |
| `ocr` | ECS Fargate service via Service Connect | 1 vCPU / 2 GB, 1 task |
| `postgres` | RDS PostgreSQL 17, Multi-AZ | `db.m7g.large`, 100 GB gp3 |
| `redis` | ElastiCache Serverless (Valkey) | — |
| `nats` | ECS Fargate + EFS for the JetStream store | 0.25 vCPU / 0.5 GB, 1 task |
| `ollama` | **not deployed** — embeddings come from OpenAI (§0.2) | — |
| uploads volume | **S3 bucket** (see §3.1) | — |

The `api` prod target is `gcr.io/distroless/static-debian12:nonroot` and the
binary is static (`CGO_ENABLED=0`), so 256 MB is genuinely enough — the 2 GB
limit in compose exists only because the dev target runs `go run` and the Go
compiler needs the headroom, not the server.

---

## 2. Model and provider configuration

### 2.1 Which Claude model

| Model | Model ID (first-party / Bedrock) | Input / Output per MTok | Where it fits here |
|---|---|---|---|
| Claude Opus 5 | `claude-opus-5` / `anthropic.claude-opus-5` | $5 / $25 | Default. Answer synthesis with citations, query routing. |
| Claude Sonnet 5 | `claude-sonnet-5` / `anthropic.claude-sonnet-5` | $3 / $15 | High-volume alternative if Q&A traffic grows. |
| Claude Haiku 4.5 | `claude-haiku-4-5` / `anthropic.claude-haiku-4-5` | $1 / $5 | The chapter-classification fallback, which is a labelling task on short text. |

Bedrock is partner-priced separately from the first-party rates above — check
[Bedrock pricing](https://aws.amazon.com/bedrock/pricing/) for the region you
deploy in.

Skriptra has two distinct LLM call sites with different economics:

- **Answer generation** (`internal/router` → `internal/retrieval`): low volume,
  quality-critical, user-facing. Opus 5.
- **Chapter classification fallback** (`internal/ingest`): runs once per
  ambiguous question at ingest, potentially thousands of calls per uploaded
  corpus, and the output is a single chapter label. Haiku 4.5.

The provider registry already supports this — `provider.Settings` is per
instance, so the ingest path can construct a second LLM with its own model
without touching the query path. If you want that split, add `CLASSIFIER_MODEL`
alongside `LLM_MODEL` in `config.Config`.

### 2.2 The Messages API is not OpenAI-compatible

`internal/provider/openaicompat` will not work against Claude. The wire formats
differ in ways that matter:

| | OpenAI chat completions | Anthropic Messages |
|---|---|---|
| Endpoint | `POST /v1/chat/completions` | `POST /v1/messages` |
| Auth header | `Authorization: Bearer` | `x-api-key` + `anthropic-version: 2023-06-01` |
| System prompt | a `messages[]` entry with `role: "system"` | a **top-level `system` field** |
| `max_tokens` | optional | **required** |
| `temperature` | accepted | **rejected with 400 on Opus 5, Sonnet 5, Opus 4.7/4.8** |
| Stream delta | `choices[].delta.content` | `content_block_delta` → `delta.text` |
| Response text | `choices[0].message.content` | `content[]`, a typed block union |

The `temperature` row is the trap. `provider.GenerateRequest` carries a
`Temperature float32`, and the current adapters forward it. An Anthropic adapter
**must drop it** — sending `temperature`, `top_p`, or `top_k` to Opus 5 returns
`400 invalid_request_error`, not a warning.

### 2.3 The adapter to write

Add `backend/internal/provider/anthropic/anthropic.go`, registered from `init()`
exactly like the existing adapters, so no switch statement anywhere changes:

```go
func init() {
    provider.RegisterLLM("anthropic", func(s provider.Settings) (provider.LLM, error) {
        return &client{settings: s, http: newHTTPClient()}, nil
    })
    // Deliberately no RegisterEmbedder: Anthropic has no embeddings API.
    // See docs/01-AWS-DEPLOYMENT.md §0.2.
}
```

The request body, following the house style of a hand-rolled HTTP client rather
than pulling in an SDK:

```jsonc
{
  "model": "claude-opus-5",
  "max_tokens": 4096,                    // required; caps thinking + text together
  "system": "<the grounding prompt>",    // top-level, not a message
  "messages": [
    { "role": "user", "content": "<retrieved context + question>" }
  ],
  "thinking":      { "type": "adaptive" },
  "output_config": { "effort": "medium" }
  // no temperature / top_p / top_k — 400 on Opus 5
}
```

Three things to get right in that adapter:

1. **`max_tokens` bounds thinking *and* visible text.** Thinking is on by default
   on Opus 5. A 1024-token budget that was fine on a non-thinking model can now
   truncate the answer mid-citation. Start at 4096 for answer generation; the
   current `GenerateRequest.MaxTokens` default should be raised to match.
2. **Prefer `effort: "medium"` over disabling thinking.** `{"type":"disabled"}`
   is accepted only at effort `high` or below, and on Opus 5 it can leak
   `<thinking>` tags into the visible response. Lower effort is the cheaper,
   safer latency lever.
3. **Map errors onto the existing sentinels.** `429` → `provider.ErrRateLimited`,
   a `400` mentioning context length → `provider.ErrContextTooLong`, a connection
   failure → `provider.ErrUnavailable`. The API layer already turns these into
   the right status codes; a generic error loses that.

Also handle `stop_reason: "refusal"` before reading `content[0]` — a refused
request returns HTTP **200** with an empty or partial `content` array. Indexing
`content[0].text` unconditionally panics on that path.

### 2.4 Bedrock instead of a raw HTTP client

Bedrock requires SigV4 request signing, which is not worth hand-rolling. Use the
official Go SDK's Bedrock client inside the adapter:

```go
import "github.com/anthropics/anthropic-sdk-go/bedrock"

client, err := bedrock.NewMantleClient(ctx, bedrock.MantleClientConfig{
    AWSRegion: os.Getenv("AWS_REGION"),
})
// then the same client.Messages.New / NewStreaming surface as first-party,
// with Model: "anthropic.claude-opus-5"
```

Credentials resolve from the ECS task role automatically — nothing to store, and
nothing to rotate. Register this under the provider name `bedrock` so
`LLM_PROVIDER=bedrock` selects it. Note that `config.requiresAPIKey()` returns
`true` for any name outside its local-runtime list, so add `"bedrock"` to that
switch or startup will demand an `LLM_API_KEY` that does not exist:

```go
case "ollama", "llamacpp", "tgi", "local", "bedrock":
    return false
```

`Config.IsLocal()` is derived from the same function, so this also keeps the UI
from claiming inference is local when it is running in Bedrock. If that reads
wrong to you, split the two concepts — "needs a key" and "stays on this machine"
are not the same predicate, and Bedrock is the first configuration where they
diverge.

### 2.5 Prompt caching

The grounding system prompt plus the course's chapter taxonomy is stable across
every question asked about a course, and the cache minimum on Opus 5 is 512
tokens. Put a `cache_control: {"type": "ephemeral"}` breakpoint on the last
system block and keep everything volatile — the user's question, the retrieved
chunks — after it. Cache reads bill at roughly 0.1×. Verify it is working by
asserting `usage.cache_read_input_tokens > 0` across two consecutive questions on
the same course; if it is zero, something per-request has leaked into the prefix.

---

## 3. Code changes required before the first deploy

These are gaps between the compose topology and a real ECS deployment. None are
large, but the deploy does not work without them.

### 3.1 Uploads must move to S3 — this is the blocking one

`api` writes to `filepath.Join(cfg.StorageDir, storageKey)`
(`internal/api/upload.go:123`) and `worker` reads from the same path
(`cmd/worker/main.go:83`). In compose both mount the `uploads` volume. **On
Fargate they are separate tasks on separate hosts and share no filesystem** —
every ingest job fails with `no such file or directory` the moment the worker
lands somewhere other than the API.

Two paths:

- **S3 (recommended).** Introduce a `Blobstore` interface with `Put`/`Get`
  alongside the existing provider interfaces, back it with the AWS SDK, and
  switch `STORAGE_DIR` to `STORAGE_BUCKET`. Roughly 150 lines across
  `upload.go`, `files.go`, and `cmd/worker`. It also fixes durability: a Fargate
  task's ephemeral storage is gone when the task is.
- **EFS (zero code change).** Mount the same EFS access point at `/data/uploads`
  in both task definitions. Works immediately, but costs more per GB, adds NFS
  latency to every read, and leaves you with a POSIX filesystem to back up
  separately.

Use EFS to get the first deploy green if you must; plan the S3 migration in the
same sprint. Note that `internal/api/files.go` already does a path-traversal
check against `filepath.Abs(StorageDir)` — the S3 version needs the equivalent
key-prefix validation, not a dropped check.

### 3.2 TLS to the database

`.env.example` ships `sslmode=disable`. RDS accepts TLS and you should require it
and verify the CA:

```
DATABASE_URL=postgres://skriptra:<pw>@<rds-endpoint>:5432/skriptra?sslmode=verify-full&sslrootcert=/etc/ssl/certs/rds-eu-central-1-bundle.pem
```

Bake the [RDS CA bundle](https://truststore.pki.rds.amazonaws.com/eu-central-1/eu-central-1-bundle.pem)
into the image, or fetch it in an init container. `sslmode=require` without
`verify-full` encrypts but does not authenticate the server — it stops passive
sniffing, not an active attacker.

### 3.3 If you change the embedding model

Not needed for the chosen setup — `text-embedding-3-small` at 768 fits the schema
as it stands. This applies only if you move to a model that cannot emit 768, such
as Cohere Multilingual v3. The dimension is load-bearing in three places, and all
three have to move together:

1. `migrations/000005_embedding_dim.up.sql` — `ALTER TABLE chunks ALTER COLUMN
   embedding TYPE vector(1024)`, the same for `question_embeddings`, then drop
   and recreate both `CHECK (embedding_dim = ...)` constraints and both HNSW
   indexes.
2. `config.validate()` — the `const schemaDimensions = 768` literal.
3. A full re-embed. Every stored vector is meaningless against the new model, so
   this is a re-ingest of the entire corpus, not a backfill. Do it against a
   restored snapshot first and measure recall@k with the eval harness before
   pointing production at it.

Rebuild the HNSW indexes **after** the re-embed, not before, and raise
`maintenance_work_mem` for the build (`SET maintenance_work_mem = '2GB'`) —
otherwise it spills to disk and takes hours.

### 3.4 Health checks

The ALB needs a target-group health check path. Confirm that whatever
`internal/api` registers returns 200 **without touching the database** — a health
check that queries Postgres turns a brief database blip into a full
deregistration of every task, which is a far worse outage than the blip.

---

## 4. Building the infrastructure

Order matters: network → data → compute → edge.

### 4.1 Network

- VPC `10.0.0.0/16`, three AZs.
- Public subnets: ALB, NAT gateway.
- Private subnets: all ECS tasks, RDS, ElastiCache.
- **VPC interface endpoints** for `bedrock-runtime`, `secretsmanager`, `ecr.api`,
  `ecr.dkr`, `logs`, and a gateway endpoint for `s3`. This keeps inference, image
  pulls, and secret fetches off the NAT gateway — meaningful on both the bill and
  the threat model.
- Security groups, each referencing the previous by ID rather than by CIDR:
  ALB ← internet:443 · api ← ALB:8080 · rds ← api,worker:5432 ·
  redis ← api:6379 · nats ← api,worker:4222 · ocr ← worker:50052.
  Egress: `api` and `worker` need outbound 443 to reach OpenAI for embeddings.

### 4.2 RDS PostgreSQL 17 with pgvector

```
Engine               PostgreSQL 17.x
Class                db.m7g.large  (2 vCPU, 8 GB) — Graviton, cheaper per unit
Storage              100 GB gp3, autoscaling to 500 GB
Multi-AZ             yes
Encryption           at rest with a customer-managed KMS key
Backups              7-day automated, PITR on
Deletion protection  on
Parameter group      custom (below)
```

The four extensions the schema needs — `vector`, `pg_trgm`, `citext`,
`uuid-ossp` — are all available on RDS PostgreSQL 17. `CREATE EXTENSION` runs as
part of `000001_init.up.sql`, executed by the master user; there is nothing extra
to enable in the console, and pgvector does not need `shared_preload_libraries`.

Custom parameter group, for the HNSW workload:

| Parameter | Value | Why |
|---|---|---|
| `maintenance_work_mem` | `2GB` | HNSW index builds are memory-bound. The default spills to disk. |
| `max_parallel_maintenance_workers` | `4` | Parallel index build. |
| `work_mem` | `64MB` | Hybrid retrieval sorts a reranked candidate set per query. |
| `shared_buffers` | ~25% of RAM (RDS default) | Leave it alone. |
| `hnsw.ef_search` | set per session, not here | Retrieval tunes recall vs latency per query. |

The indexes are built with `m = 16, ef_construction = 64` — sensible defaults.
Tune `hnsw.ef_search` (default 40) at query time: raise it for recall, lower it
for latency, and measure with the eval harness rather than by feel.

Store the credentials in Secrets Manager with rotation enabled, and grant the
task role `secretsmanager:GetSecretValue` on that one ARN.

### 4.3 ElastiCache

Serverless Valkey, cache-only. Redis holds embeddings and analytics and is
explicitly optional (`internal/cache`), so no persistence, no Multi-AZ ceremony,
and an eviction policy of `allkeys-lru`. If it goes away the app gets slower, not
wrong — keep it that way, and never put session state in it.

### 4.4 NATS JetStream

There is no managed NATS on AWS, and JetStream is load-bearing here:
`internal/queue` uses it precisely so ingestion jobs survive a restart. Run one
Fargate task with an EFS volume at `/data` for the JetStream store, fronted by
ECS Service Connect so `nats://nats:4222` resolves the same way it does in
compose.

One task is a single point of failure for *new* ingest jobs. It is not a single
point of failure for queries, which never touch it. That is an acceptable trade
for v1; if it stops being acceptable, either run a 3-node NATS cluster or replace
the queue with SQS behind the same `internal/queue` interface — the interface
exists for exactly this.

### 4.5 ECS services

One cluster, four services, all Fargate, all in private subnets, all logging to
CloudWatch via the `awslogs` driver.

- **api** — target group on 8080, ALB health check, 2 tasks minimum, autoscale on
  `ALBRequestCountPerTarget`.
- **worker** — no load balancer. Autoscale on the NATS pending-message count
  published as a custom CloudWatch metric, or fall back to CPU. Scaling to zero
  is tempting but costs you cold-start latency on the first upload after a quiet
  night; keep one warm.
- **ocr** — Service Connect, reachable at `http://ocr:50052`. Only needed if you
  accept photo and scan uploads; without it those uploads are refused with a
  message naming the missing capability, which is by design.
- **nats** — as above.

There is no embedder service: embeddings are an outbound HTTPS call to OpenAI
from `api` (query time) and `worker` (ingest time). That call is the one piece of
traffic in this design that leaves AWS, so it goes through the NAT gateway rather
than a VPC endpoint. If you later move embeddings to Bedrock Cohere, it moves
behind the `bedrock-runtime` endpoint with everything else — see §0.2.

Migrations run as a **standalone ECS task** using the `migrate/migrate:v4.17.1`
image, invoked by the deploy pipeline before the new API task set goes live — not
as a container inside the API task definition, where N tasks would race each
other on the same schema.

### 4.6 Frontend

`npm run build` produces a static bundle. There is no reason to run a Node
container in production.

```
S3 bucket (private, no website hosting)
  └─ CloudFront distribution
       ├─ Origin Access Control → S3
       ├─ Behavior  /api/*  → ALB origin (no caching, forward all headers)
       ├─ Behavior  /*      → S3, SPA fallback: 403/404 → /index.html, 200
       └─ ACM cert in us-east-1 (CloudFront requirement — the only us-east-1 thing here)
```

`VITE_API_BASE_URL` is baked in at build time, so it must be set in CI:
`VITE_API_BASE_URL=https://skriptra.example.edu/api/v1` and
`VITE_USE_MOCKS=false`. Getting `VITE_USE_MOCKS` wrong ships a convincing,
completely fake demo sitting next to a healthy live API — the compose file
already warns about this and the warning applies double in production.

Invalidate `/index.html` on every deploy; the hashed assets are immutable and
should be served with `Cache-Control: public, max-age=31536000, immutable`.

---

## 5. Production configuration

Everything below is an environment variable or a secret. The application cannot
tell the difference between this and `.env.local`, which is the point.

| Variable | Production value | Source |
|---|---|---|
| `APP_ENV` | `production` | task definition |
| `HTTP_ADDR` | `:8080` | task definition |
| `LOG_LEVEL` | `info` | task definition |
| `DATABASE_URL` | `postgres://…?sslmode=verify-full&sslrootcert=…` | Secrets Manager |
| `NATS_URL` | `nats://nats:4222` | task definition (Service Connect) |
| `REDIS_URL` | `rediss://<elasticache-endpoint>:6379` | task definition |
| `OCR_URL` | `http://ocr:50052` | task definition |
| `STORAGE_BUCKET` | `skriptra-uploads-prod` | task definition (after §3.1) |
| `MAX_UPLOAD_MB` | `50` | task definition |
| `JWT_SECRET` | 48 random bytes | **Secrets Manager** |
| `COOKIE_SECURE` | `true` | task definition — *config refuses to start otherwise* |
| `LLM_PROVIDER` | `bedrock` | task definition |
| `LLM_MODEL` | `anthropic.claude-opus-5` | task definition |
| `LLM_API_KEY` | *unset* | — (IAM task role) |
| `EMBEDDING_PROVIDER` | `openai` | task definition |
| `EMBEDDING_MODEL` | `text-embedding-3-small` | task definition |
| `EMBEDDING_BASE_URL` | `https://api.openai.com/v1` | task definition |
| `EMBEDDING_API_KEY` | OpenAI key | **Secrets Manager** |
| `EMBEDDING_DIMENSIONS` | `768` | task definition — *must stay 768; config refuses to start otherwise* |
| `AWS_REGION` | `eu-central-1` | task definition |
| `OTEL_EXPORTER_OTLP_ENDPOINT` | `http://localhost:4317` | task definition (ADOT sidecar) |

`JWT_SECRET` must be at least 32 bytes or the process refuses to start —
`openssl rand -base64 48`. Rotate it deliberately: rotating invalidates every
live session, so do it during a maintenance window, or implement a two-key
verification window first.

**IAM task role**, minimum viable:

```jsonc
{
  "Statement": [
    { "Effect": "Allow",
      "Action": ["bedrock:InvokeModel", "bedrock:InvokeModelWithResponseStream"],
      "Resource": "arn:aws:bedrock:eu-central-1::foundation-model/anthropic.claude-opus-5" },
    { "Effect": "Allow",
      "Action": ["s3:GetObject", "s3:PutObject"],
      "Resource": "arn:aws:s3:::skriptra-uploads-prod/*" },
    { "Effect": "Allow",
      "Action": ["secretsmanager:GetSecretValue"],
      "Resource": "arn:aws:secretsmanager:eu-central-1:…:secret:skriptra/*" }
  ]
}
```

Scope `bedrock:InvokeModel` to the specific model ARN, not `*`. The worker's role
needs `s3:GetObject` but not `PutObject`; the API needs both.

---

## 6. Deploy pipeline

Extend the existing GitHub Actions workflow. Authenticate with **OIDC**, not
stored access keys.

```
on: push to main
 ├─ existing CI: go build, go vet, race tests, migrate up/down/up,
 │               frontend typecheck + build
 ├─ eval harness  ← record the baseline and make this gate blocking (see PROGRESS.md)
 ├─ docker buildx bake:  backend --target api, backend --target worker,
 │                       services/ocr  → push to ECR, tagged with the git SHA
 ├─ run migrations: aws ecs run-task (migrate/migrate, waits for exit 0)
 ├─ deploy: update task definitions to the new SHA, aws ecs update-service
 │          --force-new-deployment, wait for services-stable
 └─ frontend: npm run build → aws s3 sync → cloudfront create-invalidation /index.html
```

Never tag `:latest`. Tag with the commit SHA so a rollback is `update-service`
against the previous task-definition revision, and so you can tell from
`docker inspect` exactly which commit is serving traffic.

Guard the migration step: `migrate` is not automatically reversible in the
presence of data, and the CI job that runs up/down/up proves the migrations are
structurally reversible on an *empty* database — a weaker claim than it looks.
Take an RDS snapshot before any deploy that includes a new migration.

---

## 7. Observability

`OTEL_EXPORTER_OTLP_ENDPOINT` is read by config but nothing emits spans yet (see
[`PROGRESS.md`](../PROGRESS.md)). Until that is wired, CloudWatch Logs plus these
alarms is the floor, not the ceiling:

| Metric | Alarm on | Why |
|---|---|---|
| ALB 5xx rate | > 1% over 5 min | The obvious one. |
| ALB target response time p99 | > 5 s | Retrieval or generation degrading. |
| RDS CPU / freeable memory | > 80% / < 500 MB | HNSW search is memory-hungry. |
| RDS free storage | < 20% | Vectors grow faster than people expect. |
| NATS pending messages | > 100 sustained | Workers are not keeping up; scale them. |
| Documents stuck in `parsing`/`embedding` | any older than 30 min | A wedged ingest job, invisible from HTTP metrics. |
| Bedrock `InvocationClientErrors` | any | 400s from a bad request shape — usually a `temperature` that crept back in. |
| Bedrock `InvocationThrottles` | any | Request a quota increase before it bites. |
| Worker log pattern `expected 768` | any | The embedding model stopped honouring the dimension request. Ingest is silently blocked, not silently corrupted — the check catches it, but only the log says so. |
| OpenAI 401/429 in worker logs | any | Rotated key, or an org rate limit throttling ingest. |

The stuck-documents alarm is the highest-value one on this list and the one
nobody adds. `documents.status` is a state machine; anything sitting in a
non-terminal state for half an hour is a failure the user never sees.

When you do wire OpenTelemetry, run the ADOT collector as a sidecar in each task
and export to X-Ray plus CloudWatch. Trace the ingest path end to end —
upload → parse → segment → classify → embed → index — because that is the
pipeline where a regression is otherwise invisible.

---

## 8. Cost

Rough monthly estimate, `eu-central-1`, moderate load (a few hundred students,
one corpus of a few thousand questions). Verify against the AWS pricing
calculator; these are order-of-magnitude figures, not a quote.

| Item | ~USD/month |
|---|---|
| RDS `db.m7g.large` Multi-AZ + 100 GB gp3 | 320 |
| ECS Fargate (api 2×, worker 1×, ocr 1×, nats 1×) | 130 |
| ALB | 25 |
| NAT gateway (reduced by the VPC endpoints in §4.1) | 35 |
| ElastiCache Serverless (low usage floor) | 40 |
| S3 + CloudFront + Route 53 | 15 |
| EFS (NATS store, small) | 5 |
| Secrets Manager, ECR, CloudWatch | 20 |
| **Infrastructure subtotal** | **~590** |
| Bedrock — Claude Opus 5, ~2k Q&A/month @ ~8k in / 1k out | ~105 |
| Bedrock — Haiku 4.5 classification, ~20k calls @ ~500 in / 50 out | ~5 |
| OpenAI `text-embedding-3-small` @ $0.02/MTok | **< 1** |
| **Total** | **~700** |

The obvious levers, in order of impact:

Embeddings barely register. A corpus of a few thousand questions is single-digit
millions of tokens embedded exactly once, and a query embedding is ~20 tokens.
This line item is rounding error and should not drive the vendor choice —
retrieval quality and data residency should.

1. **Single-AZ RDS in non-production.** Halves the largest line item. Do not do
   this in production.
2. **Prompt caching** on the system prompt and chapter taxonomy — cache reads
   bill at roughly 0.1×, and this workload has a large stable prefix.
3. **Sonnet 5 for answer generation** if quality holds on your eval set. A ~40%
   cut on the LLM line. Measure before switching; that is what the eval harness
   is for.
4. **Fargate Spot for the worker.** Ingest is asynchronous and restartable —
   JetStream redelivers on failure, which is exactly why it is JetStream.

---

## 9. Security checklist

- [ ] No secrets in task-definition environment variables — Secrets Manager ARNs only.
- [ ] `COOKIE_SECURE=true` (the config enforces this outside development).
- [ ] `JWT_SECRET` ≥ 32 bytes, from a CSPRNG, never in git.
- [ ] RDS not publicly accessible; TLS `verify-full` with a pinned CA bundle.
- [ ] Encryption at rest with a customer-managed KMS key: RDS, S3, EFS.
- [ ] S3 bucket: Block Public Access on, versioning on, reachable only via the task role.
- [ ] AWS WAF on the ALB or CloudFront: managed common rule set, plus rate limiting on `/api/v1/upload`.
- [ ] `bedrock:InvokeModel` scoped to one model ARN.
- [ ] `EMBEDDING_API_KEY` in Secrets Manager, never in a task-definition
      environment variable. It is the one long-lived credential left in this
      design — everything else authenticates by IAM role. Rotate it on a
      schedule, and scope the OpenAI key to a project with embeddings-only
      access if your OpenAI org supports it.
- [ ] Confirm the institution accepts that question and chunk text is processed
      by OpenAI (§0.2). If not, switch `EMBEDDING_PROVIDER` back to a
      self-hosted `nomic-embed-text` — same 768 dimensions, no re-embed.
- [ ] ECR image scanning on push; rebuild on base-image CVEs.
- [ ] Path-traversal validation preserved when `files.go` moves from filesystem paths to S3 keys.
- [ ] **Auth is a single-user stub.** Authorization checks are wired at every
      endpoint but real multi-user auth is not built (`PROGRESS.md`). Do not
      expose this to a student body until it is. Put it behind an SSO proxy or an
      IP allowlist in the meantime.
- [ ] Exam papers are usually the university's copyright. Self-hosting on the
      institution's own AWS account is the intended posture; a hosted
      multi-tenant deployment needs a rights conversation first.

---

## 10. Runbook

**Ingest is stuck.** Check `documents.status`, then worker logs, then the NATS
pending count.

```sql
SELECT id, filename, status, updated_at FROM documents
WHERE status NOT IN ('indexed','failed') AND updated_at < now() - interval '30 minutes';
```

A document stuck in `embedding` means the OpenAI call is failing. Check the
worker logs for the distinguishing error: a `401` is a rotated or revoked key, a
`429` is a rate limit (raise the org limit or add backoff), and
`returned N dimensions, expected 768` means `EMBEDDING_MODEL` was changed to
something that ignores the `dimensions` parameter. JetStream redelivers the job
once the cause is fixed — no manual requeue.

**Answers are slow.** Check ALB p99 first to separate retrieval from generation.
If retrieval: `EXPLAIN ANALYZE` the hybrid query and check whether the HNSW index
is being used — a filtered query with a very selective `course_id`/`chapter_id`
predicate may be better served by a sequential scan, and that is not a bug. If
generation: lower `output_config.effort` before anything else.

**Retrieval quality regressed after a deploy.** Run the eval harness against the
new revision and compare to the committed baseline. This is the whole reason it
exists — record the baseline and make the CI gate blocking, and this becomes a
build failure instead of a user complaint.

**Bedrock 400s after a code change.** Almost certainly a sampling parameter.
`temperature`, `top_p`, and `top_k` are rejected outright on Opus 5 and Sonnet 5
(§2.2), and `provider.GenerateRequest` still carries a `Temperature` field an
adapter can accidentally forward.

**Rollback.** `aws ecs update-service --task-definition <previous-revision>`. If
the bad deploy included a migration, restore from the pre-deploy snapshot —
rolling the schema back under live data is not a thing to attempt at speed.

---

## Appendix A — Deployment order, first time through

1. VPC, subnets, security groups, VPC endpoints.
2. RDS with the custom parameter group. Wait for `available`.
3. Secrets Manager entries: `skriptra/database-url`, `skriptra/jwt-secret`,
   `skriptra/embedding-api-key`.
4. ECR repositories: `skriptra-api`, `skriptra-worker`, `skriptra-ocr`.
5. Build and push images tagged with the git SHA.
6. Run migrations as a one-off ECS task. Verify: `\dx` shows `vector`, `pg_trgm`,
   `citext`, `uuid-ossp`; `\d chunks` shows `vector(768)`.
7. ECS cluster, then services in dependency order: nats, ocr, worker, api.
8. ALB, target group, ACM certificate, listener.
9. S3 + CloudFront + Route 53 for the frontend.
10. Smoke test: upload one PDF, watch `documents.status` walk to `indexed`, ask a
    question, confirm the answer carries page-level citations.
11. Alarms from §7 — especially the stuck-document one.
12. Load the sample corpus (`dev/seed.sql`) into a *staging* database only. Never
    into production; it is fixture data and it will pollute analytics.

## Appendix B — Verifying pgvector on RDS

```sql
-- Extensions present
SELECT extname, extversion FROM pg_extension ORDER BY extname;

-- Vector columns are the width the app expects
SELECT c.relname AS table_name, a.attname, format_type(a.atttypid, a.atttypmod) AS type
FROM pg_attribute a
JOIN pg_class c ON c.oid = a.attrelid
WHERE a.attname = 'embedding' AND a.attnum > 0;

-- HNSW indexes exist and are being used
SET hnsw.ef_search = 100;
EXPLAIN ANALYZE
SELECT id, 1 - (embedding <=> $1::vector) AS score
FROM chunks WHERE course_id = $2
ORDER BY embedding <=> $1::vector LIMIT 10;
-- Expect: Index Scan using chunks_embedding_hnsw_idx
```

If you see a sequential scan on an *unfiltered* query, the index is missing or was
never built. On a heavily filtered query a sequential scan may be the correct
plan — as the schema comment says, the filter is the optimisation.
