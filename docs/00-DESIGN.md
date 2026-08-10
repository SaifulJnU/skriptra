# Skriptra, Design Document

Version 1 · 9 August 2026

This document is the source of truth for *what* Skriptra is and *why* it is built this way. `PROGRESS.md` tracks execution state; this file records decisions and their reasoning.

---

## 1. Product

### 1.1 Problem

University students preparing for exams need past exam papers (*Altklausuren*). Today at TU Dortmund and many German universities this means paying a fee at the library, borrowing physical copies, and returning them. The papers are then unsearchable, unlinked, and un-analysable.

The questions students actually want answered are not "find me a document":

- *"Give me all Chapter 3 questions from the last five years."*
- *"Has a question like this appeared before?"*
- *"Which chapters are tested most often?"*
- *"Why is maximum likelihood used in this question, explain it from my lecture notes."*

Three of those four are not retrieval questions at all. They are **structured queries over extracted entities**. That observation drives the whole architecture.

### 1.2 Product identity

Skriptra is not an "AI exam chatbot". It is an **exam intelligence platform** with three pillars:

```
                     SKRIPTRA
                        |
        +---------------+---------------+
        |               |               |
     Search          AI Tutor       Analytics
        |               |               |
   Past papers       Explain       Topic trends
   Similar Qs        Compare       Frequency
   Filters           Practice      Difficulty
```

The AI layer is one pillar of three. A system that only chats is a demo; a system that also indexes, links and measures is a product.

### 1.3 First milestone

> A user can open the app → select a course → browse exams → open a question → see similar questions → ask an AI question → get an answer with citations.

Everything in the MVP exists to make that sentence true end to end.

---

## 2. The three hard problems

These are the parts worth building carefully, because they are what separate this from the many "chat with your PDF" projects.

### 2.1 Chapter resolution

**Problem.** A user asks for "Chapter 2 questions". Nothing in an exam question's text says "Chapter 2". Embedding similarity cannot recover this, because chapter membership is not a semantic property of the question, it is a property of the *course structure* the question belongs to.

**Approach.**

1. At course setup, ingest a syllabus or textbook table of contents and build a **chapter taxonomy**: ordered chapters with numbers, titles, and topic keywords.
2. At ingest, split each paper into individual questions, then **classify every question against that taxonomy**. Classification uses the topic keywords plus an LLM pass, and stores a confidence score.
3. At query time, a resolver maps chapter references in natural language ("chapter two", "Kapitel 2", "the MLE chapter") onto `chapter_id` **filters**, not embeddings.

**Consequence.** Chapter is a first-class column with an index, not a hope. Low-confidence classifications are surfaced in the UI for correction rather than hidden, human-in-the-loop is cheaper than perfect classification.

### 2.2 The query router

**Problem.** "Give me *all* Chapter 3 questions" cannot be answered by top-k retrieval. Retrieval returns the *k* most similar chunks; it has no notion of completeness. Asking a vector database for "all" is a category error, and it is the single most common RAG failure in production.

**Approach.** Classify the incoming query into an intent, then dispatch:

| Intent | Example | Path |
|---|---|---|
| `enumerate` | "all Chapter 3 questions from 2019-2025" | SQL over `questions` with filters, **exhaustive, ordered, paginated** |
| `explain` | "why is MLE used here" | Hybrid retrieval → LLM with citations |
| `similar` | "questions like this one" | Vector k-NN on the question embedding, excluding self |
| `analyse` | "which chapters are tested most" | SQL aggregate → chart, no LLM in the path |
| `hybrid` | "explain the most common Chapter 3 question" | `enumerate` to select, then `explain` over the result |

**Consequence.** Most queries never touch the LLM, which makes the system fast and cheap, and makes its answers exactly correct where exactness is possible. The LLM is used for explanation, not for lookup.

### 2.3 Evaluation

**Problem.** RAG systems degrade silently. A chunking tweak or a prompt edit can quietly halve retrieval quality with no error and no failing test.

**Approach.** `eval/` holds a golden dataset of real question/answer pairs drawn from actual course material, plus:

- **Retrieval metrics:** recall@k, MRR, nDCG, did the right chunk come back at all?
- **Answer metrics:** groundedness (is every claim supported by a retrieved chunk?), citation accuracy (do the cited pages actually contain the claim?), refusal correctness (does it decline when the corpus does not contain the answer?).
- **A CI gate.** A pull request that drops recall@5 below the committed baseline **fails the build**.

**Consequence.** This is the most valuable folder in the repository from a hiring perspective and the least glamorous to build. It gets built in the MVP, not "later".

---

## 3. Architecture

### 3.1 Principle: one architecture, two environments

There is no "development architecture" that gets rewritten for production. The same boundaries run in both; only infrastructure and configuration differ.

```
LOCAL (docker compose up)              PRODUCTION
  web                                    web (CDN)
  go-api                                 load balancer -> go-api (n replicas)
  worker                                 worker (n replicas)
  postgres + pgvector                    managed Postgres + pgvector
  nats                                   nats cluster
  redis                                  managed Redis
  ollama (optional profile)              hosted provider
```

Parsing runs inside the worker rather than as its own service (§4.2). The
extraction interface is defined so it can be split out later without touching
callers.

### 3.2 Provider independence

The application depends on interfaces only:

```go
type LLM interface {
    Generate(ctx context.Context, req GenerateRequest) (*GenerateResponse, error)
    Stream(ctx context.Context, req GenerateRequest) (<-chan Chunk, error)
}

type Embedder interface {
    Embed(ctx context.Context, texts []string) ([][]float32, error)
    Dimensions() int
}

type VectorStore interface {
    Upsert(ctx context.Context, chunks []Chunk) error
    Search(ctx context.Context, q Query) ([]Match, error)
}
```

Implementations are selected by configuration at dependency-initialisation time:

```
configuration -> dependency initialization -> interfaces -> application
```

There is no `if production { openai } else { ollama }` anywhere. Switching from a local model to a hosted one is `LLM_PROVIDER=...` and a restart, with no change to business logic.

**One caveat that must be designed for now:** embedding dimensions differ per model (e.g. 768 vs 1536 vs 3072). Changing `EMBEDDING_MODEL` invalidates every stored vector. The schema therefore records the embedding model and dimension on each chunk, and re-embedding is an explicit, resumable background job, not a silent corruption.

### 3.3 Ingestion pipeline

```
upload -> object store -> NATS: document.uploaded
                              |
                         ingest worker
                              |
              text + page map extraction (go-fitz)
                              |
                    question segmentation
                              |
                    chapter classification
                              |
                     chunking + embedding
                              |
                  Postgres: questions, chunks
                              |
                  NATS: document.indexed
```

Ingestion is asynchronous and idempotent. Documents are content-hashed on upload so the same paper uploaded eleven times is stored once. Status is polled by the client via `GET /api/v1/documents/{id}/status`.

### 3.4 Retrieval

Hybrid, in a single SQL statement:

- **Dense:** cosine distance over pgvector.
- **Sparse:** PostgreSQL full-text search over the same chunks.
- **Fusion:** reciprocal rank fusion, computed in SQL.
- **Filters:** `course_id`, `chapter_id`, `year`, `document_type`, applied as `WHERE` clauses, not post-filtering.
- **Source priority:** a per-source weight, so a user's own notes can outrank a shared textbook when both match.

---

## 4. Key technical decisions

### 4.1 PostgreSQL + pgvector, not a dedicated vector database

**Decision.** pgvector for the MVP, behind a `VectorStore` interface. A Qdrant adapter is planned for phase 2 specifically to produce a benchmark comparison.

**What Qdrant genuinely does better:**

- **Filterable HNSW.** Qdrant inserts extra graph links based on payload indexes, so approximate search stays accurate under selective filters. pgvector's HNSW filters after searching, which degrades recall; version 0.8's iterative index scans mitigate this but do not eliminate it.
- **Product quantization with automatic rescoring**: compressed vectors in RAM, originals on disk. pgvector has `halfvec` and binary quantization but no PQ.
- **Native multi-vector / late-interaction (ColBERT MaxSim) scoring.**
- **Native sparse vectors with server-side RRF/DBSF fusion.**
- **True horizontal sharding and replication.**

**Why pgvector wins here anyway:**

- The corpus is on the order of 200k chunks. A chapter filter reduces the candidate set to a few thousand rows, where an **exact** scan is both faster and more accurate than any approximate index. At this scale *the filter is the optimisation*, and filtered-HNSW recall, Qdrant's real advantage, does not bite until corpora two orders of magnitude larger.
- Retrieval results need joins, aggregates and window functions. The `enumerate` and `analyse` intents in §2.2 are SQL queries, not vector queries. Splitting the store would mean either duplicating relational data into the vector database or doing two round trips and joining in application code.
- One ACID transaction covers `questions` and `chunks`. No dual writes, no sync lag, no "the index thinks this document still exists".
- One backup, one migration path, one connection pool to operate.

**The trigger to revisit:** filtered result sets consistently exceeding ~100k rows, or adopting ColBERT-style late-interaction reranking. At that point it is an adapter swap, which is why the interface exists.

### 4.2 Go only, for now

**Decision: the MVP is written entirely in Go. No Python sidecar is built.**

The reflex is that RAG means Python. Component by component, that turns out to be mostly false:

| Component | Language needed |
|---|---|
| LLM calls | Go, it is an HTTP request |
| **Embeddings** | Go, also an HTTP request to Ollama or a hosted endpoint |
| Vector search, ranking, fusion | SQL |
| Chunking, dedup, orchestration, streaming | Go |
| Text + coordinates from digital PDFs | Go, `go-fitz` (MuPDF) returns text with page positions, which is exactly what citations need |
| **OCR of scanned papers** | Python, realistically, modern OCR (Surya, PaddleOCR, docTR) has no Go equivalent |
| **Layout analysis, formulas → LaTeX** | Python, effectively no Go ecosystem |

Only the last two genuinely need Python, and both are already deferred to phase 2 (§5.2). German university exam PDFs from roughly 2020 onward are usually digitally generated rather than scanned, so the first corpus is expected not to need OCR at all.

Being single-language buys real things at this budget: one container, one dependency tree, one build pipeline, one debugging story. It also produces a more distinctive result, nearly every RAG portfolio project is Python, so a production-shaped RAG system in Go is a differentiator, and it fits the backend-engineer identity this project is meant to prove rather than an AI-engineer pivot.

**The parser boundary is built; only the Python implementation is not.** `backend/internal/ingest` defines `DocumentParser` plus a `Chain` that selects an implementation **per document**, not per deployment:

- Each parser declares `Capabilities`, text layer, OCR, layout, formulas, and whether it runs in-process.
- A cheap `Probe` of the first pages decides what the document actually requires.
- The Chain picks the cheapest capable parser: a digital PDF takes the in-process Go path and never pays for a sidecar; a scan routes to whatever can read it.
- With no OCR parser registered, a scan fails with `ErrNoCapableParser` and a message naming the reason, an actionable deployment gap, not a silent empty extraction.

So the answer to "Go or Python?" is **both, resolved per document at runtime**. Adding Python means writing one adapter and registering it; nothing in the ingestion pipeline changes, because nothing in it knows which implementation ran. Selection behaviour is covered by tests in `parser_test.go`.

This is the same discipline as `VectorStore` (§4.1) and the provider interfaces (§3.2): polyglot **by necessity, not by default**. A second language costs another toolchain, another dependency tree, another CI path and a cross-language debugging story. That cost is real, so it is paid only when a document demands it.

**Trigger to build the Python adapter:** the corpus contains scanned or photographed papers, or formulas must be extracted as LaTeX rather than as approximate text.

### 4.3 API versioning from day one

Everything under `/api/v1/`. When a Flutter client exists, `/api/v2/` can be introduced without breaking it. This costs nothing now.

### 4.4 Authorization belongs to the backend

The frontend renders permissions; it never decides them. Every request is authorised server-side against the user's course memberships and document ownership. The MVP ships a single-user stub, but the *checks* are in place from the first endpoint, because retrofitting authorization once a mobile client exists is expensive and dangerous.

---

## 5. Scope

### 5.1 In the MVP

1. Upload a past-paper PDF → async ingest → split into questions → classify to chapter → index.
2. Browse: courses → chapters → exams by year/term → individual questions.
3. Similar questions across years, from the question viewer.
4. Ask, with the query router, answered with document + page citations.
5. Analytics: chapter frequency across exams.
6. Two model providers behind one interface: Ollama (local) and one hosted.
7. Eval harness with a golden set and a CI regression gate.

### 5.2 Explicitly deferred to phase 2

Voice input · vision model for formula/figure extraction · bounding-box citation highlighting (page-level is sufficient) · full multi-user auth · the personal-notes / interview-prep corpus mode · Kubernetes · Flutter · the Qdrant adapter.

Each is real. Each is also how a six-week project becomes a six-month project that never ships. The working budget is roughly 50-60 hours across six weeks.

### 5.3 Copyright

Universities generally hold rights over their exam papers. The public demo uses only material the author is entitled to share; the product is designed for self-hosting, where the corpus stays on the institution's or the student's own infrastructure. This is stated plainly rather than ignored.

---

## 6. Open questions

- Does question segmentation need a layout model, or is a rule-based splitter over "Aufgabe N / Question N" headings sufficient for the first corpus? *Try rules first; measure with the eval set.*
- Is chapter classification better as keyword-first with LLM fallback, or LLM-first? *Measure both; the eval harness exists to answer exactly this.*
- German and English material in one corpus: one multilingual embedding model, or per-language models with a language column? *Start multilingual; revisit if recall on German queries lags.*
