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
       d.year || ' ' || initcap(d.term::text)
FROM documents d
WHERE d.course_id = '22222222-2222-2222-2222-222222222222' AND d.status = 'indexed';

-- Question stems reused across years on purpose, so "similar questions" has
-- something true to find rather than noise.
CREATE TEMP TABLE stems (ord int, chap int, topic text, marks numeric, txt text, conf real, sol text) ON COMMIT DROP;
INSERT INTO stems VALUES
 (1, 2, 'OLS', 12, 'Derive the ordinary least squares estimator for beta in the linear model y = X*beta + epsilon and state the assumptions required for it to be unbiased.', 0.94,
  'Minimise S(b) = (y - Xb)^T (y - Xb). Differentiating and setting the derivative to zero gives the normal equations X^T X b = X^T y, so b_hat = (X^T X)^-1 X^T y whenever X^T X is invertible, which requires X to have full column rank. Unbiasedness needs only E[e] = 0 with X fixed: E[b_hat] = b + (X^T X)^-1 X^T E[e] = b. Neither constant variance nor normality is required for this part.'),
 (2, 2, 'Gauss-Markov', 14, 'State and prove the Gauss-Markov theorem. Explain precisely what "best" means in the acronym BLUE.', 0.91,
  'Under E[e] = 0, Var(e) = sigma^2 I and X of full column rank, OLS has the smallest variance among all linear unbiased estimators. Take any linear unbiased estimator b~ = Cy and write C = (X^T X)^-1 X^T + D. Unbiasedness forces DX = 0, and then Var(b~) = sigma^2 (X^T X)^-1 + sigma^2 D D^T, exceeding Var(b_hat) by a positive semi-definite matrix. "Best" means minimum variance within the class of linear unbiased estimators, not among all estimators. Normality is not required.'),
 (3, 3, 'F-test', 10, 'Construct an F-test for the joint significance of two regression coefficients. Give the null distribution and the rejection region at level alpha = 0.05.', 0.93,
  'Fit the unrestricted model and the model with both coefficients set to zero. With q = 2 restrictions, F = [(RSS_r - RSS_u)/q] / [RSS_u/(n - p)]. Under the null with normal errors this is F(q, n - p). Reject when F exceeds the upper 5 per cent point. The test weighs improvement in fit against the degrees of freedom spent, which is why a joint test can reject where two separate t-tests do not.'),
 (4, 3, 'Distribution theory', 12, 'Derive the distribution of the residual sum of squares under the normal linear model and use it to build a confidence interval for sigma squared.', 0.89,
  'Write RSS = y^T M y with M = I - X (X^T X)^-1 X^T, symmetric and idempotent of rank n - p. For normal errors RSS/sigma^2 is chi-squared on n - p degrees of freedom and independent of b_hat. Inverting the pivot gives [RSS / chi2_upper, RSS / chi2_lower] on n - p degrees of freedom. The interval is not symmetric, because the chi-squared distribution is not.'),
 (5, 3, 't-test', 8, 'Explain the relationship between the t-test for a single coefficient and the corresponding F-test. Show that t squared equals F in this case.', 0.87,
  'For one restriction t = (b_j - b_j0) / se(b_j), a t on n - p degrees of freedom. The F-statistic with q = 1 is the squared numerator over the same variance estimate, so F = t^2 algebraically. The distributions agree: the square of a t on n - p degrees of freedom is F(1, n - p). They differ only in that t is signed and supports one-sided alternatives, while F does not.'),
 (6, 4, 'Residual analysis', 10, 'A researcher reports R squared = 0.94 but residual plots show clear curvature. Discuss what has gone wrong and which diagnostics you would run.', 0.85,
  'R squared measures variance explained, not correctness of functional form. Curvature in residuals against fitted values means the mean function is misspecified, so the fit is high but biased. Plot residuals against each predictor to find the offending term, try a quadratic or a transformation, and use a RESET test. Misspecification also invalidates the standard errors, so the reported significance cannot be trusted either.'),
 (7, 4, 'Influence', 9, 'Define leverage and Cook''s distance. Explain how a point can have high leverage but low influence.', 0.90,
  'Leverage h_ii is the i-th diagonal of H = X (X^T X)^-1 X^T. It depends only on the predictors and measures how unusual an observation is in X-space. Cook''s distance combines leverage with the residual to measure the actual change in the fitted coefficients when the point is dropped. A point far out in X-space whose response lies exactly on the fitted surface has a near-zero residual, so removing it changes almost nothing: high leverage, low influence.'),
 (8, 5, 'Link functions', 11, 'Define the link function in a generalized linear model and derive the canonical link for the Poisson distribution.', 0.92,
  'A GLM has a random component from the exponential family, a linear predictor eta = X*beta, and a link g with g(mu) = eta. Writing the Poisson density in exponential-family form gives natural parameter theta = log(mu), so the canonical link is the log and the model is log(mu) = X*beta. The canonical link makes observed and expected information coincide, simplifying fitting, and maps a positive mean onto the whole real line so the linear predictor is unconstrained.'),
 (9, 1, 'Assumptions', 10, 'State the assumptions of the classical linear model and explain the consequence of violating each one.', 0.88,
  'Linearity in the parameters: a wrong mean function biases the estimates and more data does not help. Zero-mean errors with X exogenous: violation biases b_hat, the endogeneity problem. Constant variance: OLS stays unbiased but is no longer efficient and the usual standard errors are wrong, so use robust errors or GLS. Uncorrelated errors: same consequence, common in time series and clustered data. Full column rank of X: without it the estimator is not unique. Normality: not needed for unbiasedness or Gauss-Markov, only for exact t and F inference in small samples, where large samples lean on the CLT instead.'),
 (10, 3, 'Partitioned regression', 13, 'Given the partitioned model y = X1*beta1 + X2*beta2 + epsilon, derive the Frisch-Waugh-Lovell result.', 0.58,
  'Let M1 = I - X1 (X1^T X1)^-1 X1^T, the residual maker for X1. FWL states that b2_hat from the full regression equals the coefficient from regressing M1*y on M1*X2, and the residuals coincide. It follows from the normal equations by pre-multiplying by M1 to annihilate the X1 term. The reading is that a multiple regression coefficient is a simple regression on the part of the predictor orthogonal to everything else, which is what controlling for a variable means.');

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
       -- The solution belongs to the stem, so it answers the question asked.
       -- One shared paragraph made every question display a confident answer to
       -- a different question, under a heading claiming an official source.
       CASE WHEN e.year IN (2024, 2022) AND e.term = 'winter' THEN NULL ELSE s.sol END,
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
