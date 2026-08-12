BEGIN;

-- Back to the single-dictionary index. The column is generated, so the content
-- is rebuilt from the text either way and nothing is lost but recall.
ALTER TABLE chunks DROP COLUMN search_tsv;
ALTER TABLE chunks ADD COLUMN search_tsv tsvector
    GENERATED ALWAYS AS (to_tsvector('simple', coalesce(text, ''))) STORED;
CREATE INDEX chunks_tsv_idx ON chunks USING gin (search_tsv);

ALTER TABLE questions DROP COLUMN search_tsv;
ALTER TABLE questions ADD COLUMN search_tsv tsvector
    GENERATED ALWAYS AS (to_tsvector('simple', coalesce(text, ''))) STORED;
CREATE INDEX questions_tsv_idx ON questions USING gin (search_tsv);

COMMIT;
