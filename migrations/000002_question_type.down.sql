BEGIN;

DROP INDEX IF EXISTS questions_course_type_idx;
ALTER TABLE questions DROP COLUMN IF EXISTS question_type;
DROP TYPE IF EXISTS question_type;

COMMIT;
