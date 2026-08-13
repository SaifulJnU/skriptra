BEGIN;

-- Make the sparse half of hybrid search actually work.
--
-- The index was built with the `simple` dictionary, which does no stemming and
-- removes no stop words. Paired with plainto_tsquery, which ANDs every term,
-- the query for "Why is the ordinary least squares estimator unbiased?" became
--
--     'why' & 'is' & 'the' & 'ordinary' & 'least' & 'squares' & 'estimator' & 'unbiased'
--
-- and demanded that the passage contain the literal words "why", "is" and
-- "the". No exam question does, so the full-text side matched nothing, for any
-- natural-language question, and every retrieval was silently dense-only. The
-- evaluation harness found it: sparse scored 0.000 on all twelve hits of a
-- query whose top answer shares five words with it.
--
-- Both languages are indexed because the corpus mixes them: a German paper and
-- an English one sit in the same course, and often in the same document. The
-- two vectors are concatenated rather than chosen between, so "Schaetzer" and
-- "estimator" are both findable without knowing in advance which language a
-- chunk is in. English stemming of German text produces harmless noise, and
-- vice versa; missing one language entirely does not.
ALTER TABLE chunks DROP COLUMN search_tsv;
ALTER TABLE chunks ADD COLUMN search_tsv tsvector
    GENERATED ALWAYS AS (
        to_tsvector('english', coalesce(text, '')) ||
        to_tsvector('german',  coalesce(text, ''))
    ) STORED;
CREATE INDEX chunks_tsv_idx ON chunks USING gin (search_tsv);

ALTER TABLE questions DROP COLUMN search_tsv;
ALTER TABLE questions ADD COLUMN search_tsv tsvector
    GENERATED ALWAYS AS (
        to_tsvector('english', coalesce(text, '')) ||
        to_tsvector('german',  coalesce(text, ''))
    ) STORED;
CREATE INDEX questions_tsv_idx ON questions USING gin (search_tsv);

COMMIT;
