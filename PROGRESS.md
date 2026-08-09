# PROGRESS — read this first

**If you are Claude resuming this project in a new session: read this file, then `docs/00-DESIGN.md`, then continue from "Next step" below. Do not re-plan; the decisions are already made.**

Last updated: 9 August 2026

---

## What this is

**Skriptra** — *Intelligent Learning, Grounded in Knowledge.*

A self-hostable, provider-agnostic exam-intelligence platform. Students upload past exam papers, solutions and notes for a course; anyone in that course can then browse questions by chapter and year, find similar questions across years, and ask natural-language questions that are answered **with page-level citations**.

Origin story: at TU Dortmund, getting past papers means paying ~5€ at the library and physically returning them. In 2026.

It is the Phase 1 portfolio project for a backend/cloud engineer job hunt in Germany. Positioning comes from the CV: 3 years at Shikho (Series-A EdTech, 3M+ users) building **exam, MCQ and analytics systems** in Go/K8s/AWS. This project is the same problem domain with an LLM retrieval layer on top.

---

## Build order (the agreed strategy)

Frontend first — but as a **vertical slice against a frozen API contract**, not a throwaway mockup.

```
1. Product structure        (pages, components, navigation, states)
2. UI/UX                    (the six real screens)
3. Frontend architecture    (routing, state, API client, types, error/loading/empty)
4. API contracts            <- FROZEN before backend work starts
5. Go backend               (API -> Postgres -> NATS -> workers -> PDF -> embeddings -> retrieval -> LLM)
6. Evaluation harness       (prove retrieval actually works)
7. Production deployment    (Docker -> CI/CD -> hosting -> monitoring -> backups -> secrets)
8. Flutter / mobile         (only once the API is stable)
```

Contracts are defined early (step 4 artefact written up front) so the frontend never has to be redesigned when the backend lands.

---

## First milestone (definition of done)

> A user can open the app -> select a course -> browse exams -> open a question -> see similar questions -> ask an AI question -> get an answer with citations.

Once that flow works end to end, the product skeleton exists.

---

## Status

| # | Task | Status |
|---|---|---|
| 1 | Repo scaffold + PROGRESS.md | DONE |
| 2 | Design doc (`docs/00-DESIGN.md`) | DONE |
| 3 | v1 API contract (`api/openapi.yaml`) | DONE |
| 4 | Data model + migrations (`migrations/`) | DONE |
| 5 | Go backend skeleton + provider interfaces | DONE |
| 6 | docker-compose local stack | DONE |
| 7 | React frontend against mock data | DONE — all six screens render, typecheck clean |
| 8 | **Go API implementation** | **DONE** — all read endpoints + `/ask` + `/search`, serving real Postgres |
| 9 | **Ingestion pipeline** | **Built.** Parse, segment, classify verified on a real PDF. Full run through embeddings not yet executed. |
| 10 | **Eval harness** | **Built.** 16 golden cases, metrics, regression gate. Baseline not yet recorded. |
| 11 | CI workflow | Built. Backend, migrations up/down/up, frontend. Eval job written but commented out. |

### What is verified, and what is not

Verified by execution:

- Parse, segment and classify run correctly against `testdata/sample_exam.pdf`, a genuine rendered PDF. Correct numbering, marks, page tracking, heading stripping, and the Gauss-Markov question lands in chapter 2.
- 21 Go tests pass across ingest, router and parser selection.
- All read endpoints, `/ask` routing and analytics against a live seeded database.

**Not yet executed:** a full upload through the worker to indexed questions. That path needs an embedding model, and the Ollama image was still downloading when the session ended. Everything it depends on is individually tested; the wiring between them is not.

To finish that verification:

```bash
docker run -d --name skriptra-ollama -p 11434:11434 ollama/ollama
docker exec skriptra-ollama ollama pull nomic-embed-text
```

`nomic-embed-text` is 768-dimensional, which matches the schema. Then start `cmd/worker`, POST the sample PDF to `/api/v1/courses/{id}/documents`, and poll `/documents/{id}/status` until `indexed`.

Once that works, record the eval baseline:

```bash
cd backend && go run ./cmd/eval -update
```

### Backend, verified against a live database

Started Postgres 17 + pgvector, applied the schema, seeded `dev/seed.sql` (12 exams, 72 questions), ran `cmd/api` and exercised it:

- `GET /courses`, `/chapters`, `/exams`, `/questions`, `/analytics/chapter-frequency` all return correct real data.
- `POST /ask` — "Give me all Chapter 3 questions from the last five years" routed to `enumerate` and returned **24** questions, versus 32 for all years. The year range was resolved from the natural-language phrase, not passed in.
- `POST /ask` — "Which chapters are tested most often?" routed to `analyse` and answered from a SQL aggregate with no model in the path.
- `POST /ask` — an `explain` question with Ollama stopped returned `503 provider_unavailable` rather than a fabricated answer.

Run it yourself:

```bash
docker run -d --name skriptra-pg -e POSTGRES_USER=skriptra -e POSTGRES_PASSWORD=skriptra -e POSTGRES_DB=skriptra -p 55433:5432 pgvector/pgvector:pg17
```

### Verified, not assumed

- Migrations apply, roll back and re-apply cleanly on PostgreSQL 17 + pgvector.
- All four retrieval intents were executed against real Postgres with seed data and returned correct results. The reference SQL is `backend/internal/retrieval/queries.sql`.
- `go build ./...` and `go vet ./...` pass.
- `tsc --noEmit` passes; the app runs and every screen renders with no console errors.
- The query router was exercised in the browser: "all Chapter 3 questions" routed to `enumerate` and returned 23 questions exhaustively; "why is maximum likelihood used" routed to `explain` and streamed a cited answer.

---

## How to run it

Node is **not** installed on this machine and does not need to be — everything runs through Docker.

```bash
docker run -d --name skriptra-web -v D:\skriptra\web:/app -w /app -p 5173:5173 -e CHOKIDAR_USEPOLLING=true node:22-alpine npx vite --host 0.0.0.0
```

Then open http://localhost:5173. The app runs entirely on the mock adapter, so no backend is required yet.

---

## Next step

**Task 9 — the ingestion pipeline.** Everything downstream of it is built and proven, so this is now the only thing standing between the system and real documents:

1. `POST /courses/{id}/documents` — accept an upload, content-hash it, write the row, publish `document.uploaded` to NATS.
2. Worker: consume the job, extract text and page numbers via the `ingest.Chain` (register a `go-fitz` parser as the first implementation).
3. Segment pages into questions — try rules first (`Aufgabe N` / `Question N` headings) and measure before reaching for anything cleverer.
4. Classify each question against the chapter taxonomy; store confidence.
5. Embed chunks and questions through the `Embedder` interface; write to `chunks` and `question_embeddings`.

**Before starting, do the two-minute check:** open a real past paper and try to select the text. It selects → the Go path is correct. It does not → it is a scan, and the Python OCR adapter gets registered behind the existing `ingest.Chain` seam.

Then task 10, the eval harness, which is the highest-value remaining item for interviews.

---

## Frontend architecture notes

- **Mock layer is an adapter, not MSW.** `src/lib/api.ts` defines the `SkriptraApi` interface plus a real HTTP implementation; `src/mocks/mockApi.ts` implements the same interface; `src/lib/client.ts` picks between them on `VITE_USE_MOCKS`. Same discipline as the Go provider interfaces — no service worker to configure, and components never know which one they got.
- Types in `src/types/api.ts` mirror `api/openapi.yaml`. When the server lands, generate them from the spec instead of editing by hand.
- Filters live in the URL (`?chapter=3&year=2025`) so filtered lists are shareable and survive refresh.
- The mock corpus deliberately includes unclassified questions, low-confidence classifications and a mid-ingest document, so those states are designed for rather than retrofitted.

---

## Screens to build (task 7 detail)

| Screen | Route | What it must prove |
|---|---|---|
| Dashboard | `/` | Course cards with exam + question counts; recent activity |
| Course | `/courses/:id` | Tabs: Overview / Exams / Questions / Analytics. Chapter list is the primary navigation. |
| Exam browser | `/courses/:id/exams` | Grouped by year and term (2025 Summer, 2025 Winter, ...) |
| Question viewer | `/questions/:id` | Question text, chapter, topic, source page — **plus "Similar questions" across years.** This screen is the product. |
| Ask | `/courses/:id/ask` | Streaming answer + a Sources list of `Exam · Question · Page`, each clickable to the PDF page |
| Analytics | `/courses/:id/analytics` | "Most frequently tested chapters" bar chart — this is what makes it more than a chatbot |

Upload lives as a modal/drawer, not a screen, with async ingest status polling via `GET /api/v1/documents/{id}/status`.

---

## Decisions already locked (do not relitigate)

- **Name:** Skriptra.
- **Vector store:** PostgreSQL + **pgvector**, behind a `VectorStore` interface. Qdrant adapter is phase 2, kept specifically to produce a benchmark comparison for interviews. Rationale in `docs/00-DESIGN.md` — at ~200k chunks the chapter filter *is* the optimisation, and filtered-HNSW recall (Qdrant's real advantage) does not bite until far larger corpora.
- **Backend: Go only.** No Python. Embeddings and LLM calls are HTTP, retrieval is SQL, and PDF text + page coordinates come from `go-fitz`. Python is warranted solely for OCR and layout analysis, which are phase 2 — the extraction interface is a seam so adding a sidecar later is additive. Rationale in `docs/00-DESIGN.md` §4.2.
- **Messaging:** NATS for ingestion jobs. **Cache:** Redis.
- **Model independence:** `LLM` and `Embedder` interfaces; implementation chosen by env config only. No `if production` branching anywhere.
- **API versioning:** everything under `/api/v1/` from day one.
- **Authorization lives in the backend**, never the frontend — mobile clients come later.
- **Same architecture local and prod.** Difference is infrastructure and config, not code.

## Explicitly deferred (phase 2 — do not build these yet)

Python parser sidecar (OCR, layout, formulas → LaTeX) · vision model for formula extraction · bounding-box citation highlighting (page-level is enough) · multi-user auth beyond a stub · the personal-notes/interview-prep corpus mode · Kubernetes · Flutter · Qdrant adapter.

**Voice input is built** (browser SpeechRecognition, `web/src/lib/useVoiceInput.ts`) — it needed no model and no backend, so it did not belong in the deferred list.

**Before building ingestion:** open a real past paper and try to select the text. If it selects, the PDF is digital and Go-only extraction is correct. If nothing selects, it is a scan and OCR becomes necessary — at which point the Python sidecar gets built behind the existing seam.

Every one is a real feature. Every one is also how this becomes a six-month project that never ships. The budget is ~50-60 working hours across six weeks.

---

## Context for whoever resumes this

The owner is a new father (daughter born 31 July 2026) with a Linear Models exam on ~19-20 August 2026. His own life plan freezes portfolio work until 21 August. He chose to start early anyway — that was raised once and is his call; do not raise it repeatedly.

He works in ~90-minute blocks and gets overwhelmed by parallel task lists. **Give him one ranked next step at a time, never a menu.**
