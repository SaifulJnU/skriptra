-- Question format, as distinct from question topic.
--
-- Students revise by format as often as by topic: "give me the true/false
-- questions", "which ones are proofs". Without this the enumerate path could
-- only filter by chapter and year, so a request for true/false questions in
-- chapter 2 silently returned every chapter 2 question, which is a worse answer
-- than none.
--
-- Derived from the question wording at ingest, not from the chapter taxonomy,
-- because format and topic are independent: a chapter can be examined by proof
-- one year and by multiple choice the next.

BEGIN;

CREATE TYPE question_type AS ENUM (
    'true_false',
    'multiple_choice',
    'proof',
    'derivation',
    'computation',
    'discussion',
    'unknown'
);

ALTER TABLE questions
    ADD COLUMN question_type question_type NOT NULL DEFAULT 'unknown';

-- Filtering is always scoped to a course, and usually to a chapter as well.
CREATE INDEX questions_course_type_idx ON questions (course_id, question_type);

COMMIT;
