-- Development seed: a small but realistic Linear Models corpus.
--
-- Deliberately includes the awkward states, because those are the ones that
-- break UIs: an unclassified question, a low-confidence classification, an exam
-- without solutions, and a document still mid-ingest.
--
--   docker exec -i <pg> psql -U skriptra -d skriptra -f /tmp/seed.sql
--
-- Embeddings here are synthetic unit vectors, not real ones. They make
-- similarity and hybrid search exercisable without a model running; the
-- ingestion pipeline replaces them with real vectors.

BEGIN;

TRUNCATE users, courses, chapters, documents, exams, questions,
         chunks, question_embeddings, ingest_jobs, course_members,
         conversations, messages RESTART IDENTITY CASCADE;

INSERT INTO users (id, display_name, email)
VALUES ('11111111-1111-1111-1111-111111111111', 'Saiful', 'saiful@example.com');

INSERT INTO courses (id, name, code, institution, language, created_by) VALUES
 ('22222222-2222-2222-2222-222222222222', 'Linear Models', 'STAT-412', 'TU Dortmund', 'en',
  '11111111-1111-1111-1111-111111111111'),
 ('22222222-2222-2222-2222-222222222299', 'Natural Language Processing', 'CS-508', 'TU Dortmund', 'en',
  '11111111-1111-1111-1111-111111111111');

INSERT INTO course_members (course_id, user_id, role) VALUES
 ('22222222-2222-2222-2222-222222222222', '11111111-1111-1111-1111-111111111111', 'owner'),
 ('22222222-2222-2222-2222-222222222299', '11111111-1111-1111-1111-111111111111', 'owner');

INSERT INTO chapters (id, course_id, number, title, topics) VALUES
 ('33333333-0000-0000-0000-000000000001', '22222222-2222-2222-2222-222222222222', 1, 'The Linear Model',                 ARRAY['design matrix','assumptions']),
 ('33333333-0000-0000-0000-000000000002', '22222222-2222-2222-2222-222222222222', 2, 'Least Squares Estimation',         ARRAY['OLS','Gauss-Markov','BLUE']),
 ('33333333-0000-0000-0000-000000000003', '22222222-2222-2222-2222-222222222222', 3, 'Inference and Hypothesis Testing', ARRAY['F-test','t-test','confidence region']),
 ('33333333-0000-0000-0000-000000000004', '22222222-2222-2222-2222-222222222222', 4, 'Model Diagnostics',               ARRAY['residuals','leverage','Cook''s distance']),
 ('33333333-0000-0000-0000-000000000005', '22222222-2222-2222-2222-222222222222', 5, 'Generalized Linear Models',       ARRAY['link function','logistic','Poisson']);

-- Six years of papers. 2022 winter and 2024 winter have no solutions.
INSERT INTO documents (id, course_id, uploaded_by, filename, kind, status, storage_key,
                       content_hash, size_bytes, page_count, year, term)
SELECT
  ('44444444-0000-0000-0000-' || lpad(row_number() OVER ()::text, 12, '0'))::uuid,
  '22222222-2222-2222-2222-222222222222',
  '11111111-1111-1111-1111-111111111111',
  'LiMo_' || y || '_' || CASE WHEN t = 'summer' THEN 'SS' ELSE 'WS' END || '.pdf',
  'exam', 'indexed',
  's/' || y || t,
  md5(y::text || t) || md5(t || y::text),
  1240000, 8, y, t::term
FROM (VALUES (2025,'summer'),(2025,'winter'),(2024,'summer'),(2024,'winter'),
             (2023,'summer'),(2023,'winter'),(2022,'summer'),(2022,'winter'),
             (2021,'summer'),(2021,'winter'),(2020,'summer'),(2020,'winter')) AS v(y,t);

-- A document still being processed, so the ingest UI has a live state.
INSERT INTO documents (id, course_id, filename, kind, status, storage_key, content_hash,
                       size_bytes, page_count, year, term)
VALUES ('44444444-0000-0000-0000-000000000099', '22222222-2222-2222-2222-222222222222',
        'LiMo_2019_WS_scan.pdf', 'exam', 'classifying', 's/2019ws',
        repeat('c', 64), 8900000, 7, 2019, 'winter');

INSERT INTO ingest_jobs (document_id, status, progress, stage_detail, questions_extracted)
VALUES ('44444444-0000-0000-0000-000000000099', 'classifying', 0.62, 'classifying question 12 of 31', 12);

INSERT INTO exams (id, course_id, document_id, year, term, title)
SELECT ('55555555-0000-0000-0000-' || lpad(row_number() OVER (ORDER BY d.year DESC, d.term)::text, 12, '0'))::uuid,
       d.course_id, d.id, d.year, d.term,
       d.year || ' — ' || initcap(d.term::text)
FROM documents d
WHERE d.course_id = '22222222-2222-2222-2222-222222222222' AND d.status = 'indexed';

-- Question stems reused across years on purpose, so "similar questions" has
-- something true to find rather than noise.
CREATE TEMP TABLE stems (ord int, chap int, topic text, marks numeric, txt text, conf real) ON COMMIT DROP;
INSERT INTO stems VALUES
 (1, 2, 'OLS',                 12, 'Derive the ordinary least squares estimator for beta in the linear model y = X*beta + epsilon and state the assumptions required for it to be unbiased.', 0.94),
 (2, 2, 'Gauss-Markov',        14, 'State and prove the Gauss-Markov theorem. Explain precisely what "best" means in the acronym BLUE.', 0.91),
 (3, 3, 'F-test',              10, 'Construct an F-test for the joint significance of two regression coefficients. Give the null distribution and the rejection region at level alpha = 0.05.', 0.93),
 (4, 3, 'Distribution theory', 12, 'Derive the distribution of the residual sum of squares under the normal linear model and use it to build a confidence interval for sigma squared.', 0.89),
 (5, 3, 't-test',               8, 'Explain the relationship between the t-test for a single coefficient and the corresponding F-test. Show that t squared equals F in this case.', 0.87),
 (6, 4, 'Residual analysis',   10, 'A researcher reports R squared = 0.94 but residual plots show clear curvature. Discuss what has gone wrong and which diagnostics you would run.', 0.85),
 (7, 4, 'Influence',            9, 'Define leverage and Cook''s distance. Explain how a point can have high leverage but low influence.', 0.90),
 (8, 5, 'Link functions',      11, 'Define the link function in a generalized linear model and derive the canonical link for the Poisson distribution.', 0.92),
 (9, 1, 'Assumptions',         10, 'State the assumptions of the classical linear model and explain the consequence of violating each one.', 0.88),
 (10,3, 'Partitioned regression', 13, 'Given the partitioned model y = X1*beta1 + X2*beta2 + epsilon, derive the Frisch-Waugh-Lovell result.', 0.58);

INSERT INTO questions (course_id, exam_id, document_id, number, ordinal, text, marks, source_page,
                       chapter_id, chapter_confidence, chapter_source, topic,
                       solution_text, solution_page)
SELECT e.course_id, e.id, e.document_id,
       s.ord::text, s.ord, s.txt, s.marks, s.ord,
       -- every 11th question is left unclassified: the UI must handle a
       -- question the classifier could not place
       CASE WHEN (e.year + s.ord) % 11 = 0 THEN NULL
            ELSE (SELECT id FROM chapters c WHERE c.course_id = e.course_id AND c.number = s.chap) END,
       CASE WHEN (e.year + s.ord) % 11 = 0 THEN NULL ELSE s.conf END,
       CASE WHEN (e.year + s.ord) % 11 = 0 THEN NULL ELSE 'llm' END,
       CASE WHEN (e.year + s.ord) % 11 = 0 THEN NULL ELSE s.topic END,
       CASE WHEN e.year IN (2024, 2022) AND e.term = 'winter' THEN NULL
            ELSE 'Start from the log-likelihood and set the score to zero; the normal equations follow directly, and the ML variance estimator divides by n rather than n - p, which is why it is biased downward.' END,
       CASE WHEN e.year IN (2024, 2022) AND e.term = 'winter' THEN NULL ELSE s.ord + 4 END
FROM exams e
CROSS JOIN stems s
WHERE e.course_id = '22222222-2222-2222-2222-222222222222'
  AND s.ord <= 5 + (e.year % 3);

-- Synthetic embeddings: one axis per chapter, so questions in the same chapter
-- are near-identical and different chapters are orthogonal. Enough to exercise
-- k-NN and RRF without a model running.
INSERT INTO chunks (course_id, document_id, question_id, chapter_id, ordinal, text, page,
                    embedding, embedding_model, embedding_dim, source_weight)
SELECT q.course_id, q.document_id, q.id, q.chapter_id, 1, q.text, q.source_page,
       ('[' || array_to_string(
           (SELECT array_agg(CASE WHEN i = COALESCE(ch.number, 6) THEN 1.0 ELSE 0.0 END)
            FROM generate_series(1, 768) i), ',') || ']')::vector,
       'seed-synthetic', 768, 1.0
FROM questions q
LEFT JOIN chapters ch ON ch.id = q.chapter_id
WHERE q.course_id = '22222222-2222-2222-2222-222222222222';

INSERT INTO question_embeddings (question_id, course_id, embedding, embedding_model, embedding_dim)
SELECT c.question_id, c.course_id, c.embedding, 'seed-synthetic', 768
FROM chunks c WHERE c.question_id IS NOT NULL;

COMMIT;

\echo ''
SELECT (SELECT count(*) FROM courses)   AS courses,
       (SELECT count(*) FROM chapters)  AS chapters,
       (SELECT count(*) FROM exams)     AS exams,
       (SELECT count(*) FROM questions) AS questions,
       (SELECT count(*) FROM chunks)    AS chunks;
