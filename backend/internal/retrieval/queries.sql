-- Reference retrieval queries, one per query intent.
--
-- All four have been executed against PostgreSQL 17 + pgvector and verified to
-- return correct results. They are kept here as the canonical shapes; the Go
-- layer parameterises them, it does not reinvent them.
--
-- The central design claim: only ONE of these four intents is a vector search.
-- Routing to the right one is what makes the system correct rather than
-- plausible.

-- name: EnumerateQuestions :many
-- INTENT `enumerate` — "give me ALL chapter 3 questions".
-- Exhaustive, ordered, paginated. No vectors involved. Top-k retrieval cannot
-- answer this correctly and is not asked to.
SELECT q.id, q.number, q.text, q.marks, q.source_page,
       e.year, e.term,
       ch.id AS chapter_id, ch.number AS chapter_number, ch.title AS chapter_title,
       q.chapter_confidence, q.topic,
       (q.solution_text IS NOT NULL) AS has_solution
FROM questions q
JOIN exams    e  ON e.id  = q.exam_id
LEFT JOIN chapters ch ON ch.id = q.chapter_id
WHERE q.course_id = @course_id
  AND (@chapter_number::int  IS NULL OR ch.number = @chapter_number)
  AND (@year_from::int       IS NULL OR e.year   >= @year_from)
  AND (@year_to::int         IS NULL OR e.year   <= @year_to)
  AND (@term::term           IS NULL OR e.term    = @term)
ORDER BY e.year DESC, e.term, q.ordinal
LIMIT @page_size OFFSET @page_offset;

-- name: ChapterFrequency :many
-- INTENT `analyse` — "which chapters are tested most often".
-- A pure aggregate. The LLM is never in this path, so the numbers are exact and
-- the response is instant. This is what makes Lernova more than a chatbot.
SELECT ch.id, ch.number, ch.title,
       count(q.id)                                              AS question_count,
       count(DISTINCT q.exam_id)                                AS exam_count,
       COALESCE(count(q.id)::float / NULLIF(sum(count(q.id)) OVER (), 0), 0) AS share
FROM chapters ch
LEFT JOIN questions q ON q.chapter_id = ch.id
LEFT JOIN exams    e  ON e.id = q.exam_id
     AND (@year_from::int IS NULL OR e.year >= @year_from)
     AND (@year_to::int   IS NULL OR e.year <= @year_to)
WHERE ch.course_id = @course_id
GROUP BY ch.id, ch.number, ch.title
ORDER BY ch.number;

-- name: SimilarQuestions :many
-- INTENT `similar` — "questions like this one, across years".
-- Question-level k-NN, excluding self. Unit of similarity is the whole
-- question, which is why question_embeddings is separate from chunks.
SELECT q.id, q.number, q.text, q.source_page,
       e.year, e.term,
       ch.number AS chapter_number, ch.title AS chapter_title,
       1 - (qe.embedding <=> ref.embedding) AS score
FROM question_embeddings qe
JOIN questions q  ON q.id  = qe.question_id
JOIN exams     e  ON e.id  = q.exam_id
LEFT JOIN chapters ch ON ch.id = q.chapter_id
CROSS JOIN (
    SELECT embedding FROM question_embeddings WHERE question_id = @question_id
) ref
WHERE qe.course_id  = @course_id
  AND qe.question_id <> @question_id
  AND 1 - (qe.embedding <=> ref.embedding) >= @min_score
ORDER BY qe.embedding <=> ref.embedding
LIMIT @limit_n;

-- name: HybridSearch :many
-- INTENT `explain` — retrieve passages to ground an answer.
--
-- Dense (pgvector cosine) + sparse (Postgres full-text) fused with reciprocal
-- rank fusion, all in one statement and one round trip.
--
-- Note `filtered`: structured filters are applied as a WHERE clause BEFORE
-- ranking, not as post-filtering. This is the single most important line in the
-- file. A selective chapter filter cuts the candidate set to a few thousand
-- rows, where an exact scan is both faster and more accurate than an
-- approximate index — at this corpus size the filter IS the optimisation, which
-- is precisely why a dedicated vector database is not needed yet.
WITH params AS (
    SELECT @course_id::uuid       AS course_id,
           @query_text::text      AS q_text,
           @query_vector::vector  AS q_vec,
           @chapter_number::int   AS chapter_number,
           @year_from::int        AS year_from,
           @year_to::int          AS year_to,
           60::int                AS k          -- RRF damping constant
),
filtered AS (
    SELECT c.*
    FROM chunks c
    JOIN params p           ON c.course_id = p.course_id
    LEFT JOIN chapters  ch  ON ch.id = c.chapter_id
    LEFT JOIN documents d   ON d.id  = c.document_id
    WHERE (p.chapter_number IS NULL OR ch.number = p.chapter_number)
      AND (p.year_from      IS NULL OR d.year   >= p.year_from)
      AND (p.year_to        IS NULL OR d.year   <= p.year_to)
),
dense AS (
    SELECT f.id,
           row_number() OVER (ORDER BY f.embedding <=> p.q_vec) AS rnk,
           1 - (f.embedding <=> p.q_vec)                        AS score
    FROM filtered f CROSS JOIN params p
    ORDER BY f.embedding <=> p.q_vec
    LIMIT 50
),
sparse AS (
    SELECT f.id,
           row_number() OVER (
               ORDER BY ts_rank_cd(f.search_tsv, plainto_tsquery('simple', p.q_text)) DESC
           ) AS rnk,
           ts_rank_cd(f.search_tsv, plainto_tsquery('simple', p.q_text)) AS score
    FROM filtered f CROSS JOIN params p
    WHERE f.search_tsv @@ plainto_tsquery('simple', p.q_text)
    ORDER BY score DESC
    LIMIT 50
),
fused AS (
    SELECT COALESCE(d.id, s.id) AS chunk_id,
           COALESCE(1.0 / (p.k + d.rnk), 0) + COALESCE(1.0 / (p.k + s.rnk), 0) AS rrf,
           d.score AS dense_score,
           s.score AS sparse_score
    FROM dense d
    FULL OUTER JOIN sparse s ON s.id = d.id
    CROSS JOIN params p
)
SELECT c.id AS chunk_id, c.text, c.page,
       f.rrf * c.source_weight AS score,        -- source_weight = "prefer my own notes"
       f.dense_score, f.sparse_score,
       d.id AS document_id, d.filename, d.kind AS document_kind,
       q.id AS question_id, q.number AS question_number
FROM fused f
JOIN chunks    c ON c.id = f.chunk_id
JOIN documents d ON d.id = c.document_id
LEFT JOIN questions q ON q.id = c.question_id
ORDER BY f.rrf * c.source_weight DESC
LIMIT @limit_n;
