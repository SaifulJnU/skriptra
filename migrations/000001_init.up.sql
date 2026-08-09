-- Lernova initial schema
--
-- Design notes worth reading before changing anything here:
--
--  * `questions` is a first-class table, not a view over chunks. The `enumerate`
--    and `analyse` query intents are SQL over this table. Retrieval cannot
--    answer "give me ALL chapter 3 questions" and is not asked to.
--
--  * `chunks.embedding` has a FIXED dimension. Changing the embedding model
--    changes the dimension and invalidates every stored vector, so the model
--    and dimension are recorded per row and re-embedding is an explicit,
--    resumable job. See 000002 if the dimension ever needs to change.
--
--  * Full-text uses the 'simple' configuration deliberately: the corpus is
--    mixed German and English, and a generated column cannot switch stemmer
--    per row (to_tsvector with a non-constant config is not immutable).

BEGIN;

CREATE EXTENSION IF NOT EXISTS vector;
CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ---------------------------------------------------------------- enums ----

CREATE TYPE term            AS ENUM ('summer', 'winter');
CREATE TYPE document_kind   AS ENUM ('exam', 'solution', 'notes', 'textbook', 'syllabus');
CREATE TYPE ingest_status   AS ENUM ('queued', 'parsing', 'segmenting', 'classifying',
                                     'embedding', 'indexed', 'failed');
CREATE TYPE query_intent    AS ENUM ('enumerate', 'explain', 'similar', 'analyse', 'hybrid');
CREATE TYPE course_role     AS ENUM ('owner', 'member');

-- ---------------------------------------------------------------- users ----

CREATE TABLE users (
    id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    email         citext,
    display_name  text NOT NULL,
    created_at    timestamptz NOT NULL DEFAULT now()
);

-- MVP ships a single-user stub, but authorization checks are wired from the
-- first endpoint. Retrofitting them once a mobile client exists is expensive.
CREATE TABLE course_members (
    course_id  uuid NOT NULL,
    user_id    uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role       course_role NOT NULL DEFAULT 'member',
    joined_at  timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (course_id, user_id)
);

-- -------------------------------------------------------------- courses ----

CREATE TABLE courses (
    id           uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name         text NOT NULL,
    code         text,
    institution  text,
    language     text NOT NULL DEFAULT 'en' CHECK (language IN ('en', 'de')),
    created_by   uuid REFERENCES users(id) ON DELETE SET NULL,
    created_at   timestamptz NOT NULL DEFAULT now()
);

ALTER TABLE course_members
    ADD CONSTRAINT course_members_course_fk
    FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE;

-- The chapter taxonomy. This is what makes "chapter 2" a filter instead of a
-- hope: derived once at course setup from a syllabus or textbook contents page.
CREATE TABLE chapters (
    id          uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    course_id   uuid NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    number      integer NOT NULL,
    title       text NOT NULL,
    topics      text[] NOT NULL DEFAULT '{}',
    created_at  timestamptz NOT NULL DEFAULT now(),
    UNIQUE (course_id, number)
);

-- ------------------------------------------------------------ documents ----

CREATE TABLE documents (
    id            uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    course_id     uuid NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    uploaded_by   uuid REFERENCES users(id) ON DELETE SET NULL,
    filename      text NOT NULL,
    kind          document_kind NOT NULL,
    status        ingest_status NOT NULL DEFAULT 'queued',
    storage_key   text NOT NULL,
    content_hash  char(64) NOT NULL,          -- sha-256, dedup key
    size_bytes    bigint NOT NULL,
    page_count    integer,
    year          integer,
    term          term,
    language      text,
    error         text,
    uploaded_at   timestamptz NOT NULL DEFAULT now(),
    indexed_at    timestamptz,
    -- the same paper uploaded eleven times is stored once
    UNIQUE (course_id, content_hash)
);

CREATE INDEX documents_course_status_idx ON documents (course_id, status);

-- Live ingestion progress, kept separate so frequent progress writes do not
-- churn the documents row that everything else reads.
CREATE TABLE ingest_jobs (
    document_id         uuid PRIMARY KEY REFERENCES documents(id) ON DELETE CASCADE,
    status              ingest_status NOT NULL DEFAULT 'queued',
    progress            real NOT NULL DEFAULT 0 CHECK (progress BETWEEN 0 AND 1),
    stage_detail        text,
    questions_extracted integer NOT NULL DEFAULT 0,
    attempts            integer NOT NULL DEFAULT 0,
    error               text,
    started_at          timestamptz,
    completed_at        timestamptz,
    updated_at          timestamptz NOT NULL DEFAULT now()
);

-- ---------------------------------------------------------------- exams ----

CREATE TABLE exams (
    id           uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    course_id    uuid NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    document_id  uuid REFERENCES documents(id) ON DELETE SET NULL,
    year         integer NOT NULL,
    term         term NOT NULL,
    title        text,
    created_at   timestamptz NOT NULL DEFAULT now(),
    UNIQUE (course_id, year, term)
);

CREATE INDEX exams_course_year_idx ON exams (course_id, year DESC, term);

-- ------------------------------------------------------------ questions ----

CREATE TABLE questions (
    id                   uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    course_id            uuid NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    exam_id              uuid REFERENCES exams(id) ON DELETE CASCADE,
    document_id          uuid NOT NULL REFERENCES documents(id) ON DELETE CASCADE,
    number               text NOT NULL,           -- as printed: "4", "4b"
    ordinal              integer NOT NULL,        -- sort order within the exam
    text                 text NOT NULL,
    marks                numeric(5,2),
    source_page          integer NOT NULL,        -- 1-indexed, drives citations

    chapter_id           uuid REFERENCES chapters(id) ON DELETE SET NULL,
    chapter_confidence   real CHECK (chapter_confidence BETWEEN 0 AND 1),
    chapter_source       text CHECK (chapter_source IN ('keyword', 'llm', 'manual')),
    topic                text,

    solution_text        text,
    solution_page        integer,
    solution_document_id uuid REFERENCES documents(id) ON DELETE SET NULL,

    search_tsv           tsvector GENERATED ALWAYS AS
                             (to_tsvector('simple', coalesce(text, ''))) STORED,
    created_at           timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX questions_course_chapter_idx ON questions (course_id, chapter_id);
CREATE INDEX questions_exam_idx           ON questions (exam_id, ordinal);
CREATE INDEX questions_document_idx       ON questions (document_id);
CREATE INDEX questions_tsv_idx            ON questions USING gin (search_tsv);
CREATE INDEX questions_trgm_idx           ON questions USING gin (text gin_trgm_ops);
-- surfaces low-confidence classifications for human correction
CREATE INDEX questions_low_confidence_idx ON questions (course_id, chapter_confidence)
    WHERE chapter_confidence IS NOT NULL AND chapter_confidence < 0.7;

-- --------------------------------------------------------------- chunks ----

-- 768 matches the default multilingual embedding model. Changing the model
-- means changing this dimension: add a migration, re-embed, then swap. The
-- model/dimension columns exist so a mismatch is a loud error, not silent
-- garbage similarity scores.
CREATE TABLE chunks (
    id              uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    course_id       uuid NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    document_id     uuid NOT NULL REFERENCES documents(id) ON DELETE CASCADE,
    question_id     uuid REFERENCES questions(id) ON DELETE CASCADE,
    chapter_id      uuid REFERENCES chapters(id) ON DELETE SET NULL,

    ordinal         integer NOT NULL,
    text            text NOT NULL,
    page            integer NOT NULL,
    token_count     integer,

    embedding       vector(768),
    embedding_model text NOT NULL,
    embedding_dim   integer NOT NULL,

    source_weight   real NOT NULL DEFAULT 1.0,   -- source-priority ranking
    search_tsv      tsvector GENERATED ALWAYS AS
                        (to_tsvector('simple', coalesce(text, ''))) STORED,
    created_at      timestamptz NOT NULL DEFAULT now(),

    CONSTRAINT chunks_dim_matches CHECK (embedding_dim = 768)
);

-- Filters are applied as WHERE clauses, never as post-filtering. At this corpus
-- size a selective chapter filter reduces the candidate set to a few thousand
-- rows, where an exact scan beats an approximate index on both speed and
-- recall — the filter IS the optimisation.
CREATE INDEX chunks_course_chapter_idx ON chunks (course_id, chapter_id);
CREATE INDEX chunks_document_idx       ON chunks (document_id, ordinal);
CREATE INDEX chunks_question_idx       ON chunks (question_id);
CREATE INDEX chunks_tsv_idx            ON chunks USING gin (search_tsv);

-- HNSW for unfiltered / weakly-filtered semantic search across a whole course.
CREATE INDEX chunks_embedding_hnsw_idx ON chunks
    USING hnsw (embedding vector_cosine_ops)
    WITH (m = 16, ef_construction = 64);

-- Question-level embeddings power "similar questions across years". Kept apart
-- from chunks because the unit of similarity is the whole question, not a
-- passage window.
CREATE TABLE question_embeddings (
    question_id     uuid PRIMARY KEY REFERENCES questions(id) ON DELETE CASCADE,
    course_id       uuid NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    embedding       vector(768) NOT NULL,
    embedding_model text NOT NULL,
    embedding_dim   integer NOT NULL,
    created_at      timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT question_embeddings_dim_matches CHECK (embedding_dim = 768)
);

CREATE INDEX question_embeddings_course_idx ON question_embeddings (course_id);
CREATE INDEX question_embeddings_hnsw_idx ON question_embeddings
    USING hnsw (embedding vector_cosine_ops)
    WITH (m = 16, ef_construction = 64);

-- -------------------------------------------------------- conversations ----

CREATE TABLE conversations (
    id          uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    course_id   uuid NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    user_id     uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title       text,
    created_at  timestamptz NOT NULL DEFAULT now(),
    updated_at  timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX conversations_user_idx ON conversations (user_id, updated_at DESC);

CREATE TABLE messages (
    id               uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    conversation_id  uuid NOT NULL REFERENCES conversations(id) ON DELETE CASCADE,
    role             text NOT NULL CHECK (role IN ('user', 'assistant')),
    content          text NOT NULL,
    intent           query_intent,
    -- citations are denormalised onto the message so an answer stays
    -- reproducible even if a source document is later deleted
    sources          jsonb NOT NULL DEFAULT '[]'::jsonb,
    usage            jsonb,
    created_at       timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX messages_conversation_idx ON messages (conversation_id, created_at);

COMMIT;
