package db

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/skriptra/skriptra/backend/internal/ingest"
)

// This file implements ingest.Store. The interface is declared by the pipeline,
// which states what it needs; this type happens to satisfy it.

// CreateDocument records an upload, deduplicating by content hash.
//
// Returns existing=true when the same bytes are already in the course. Students
// upload the same paper repeatedly, and indexing it eleven times would inflate
// every count and skew the analytics the product is supposed to make
// trustworthy.
func (s *Store) CreateDocument(ctx context.Context, courseID uuid.UUID, filename, kind, storageKey, contentHash string, sizeBytes int64, year *int, term *string) (uuid.UUID, bool, error) {
	var id uuid.UUID
	err := s.pool.QueryRow(ctx, `
		SELECT id FROM documents WHERE course_id = $1 AND content_hash = $2`,
		courseID, contentHash).Scan(&id)
	if err == nil {
		return id, true, nil
	}
	if err != pgx.ErrNoRows {
		return uuid.Nil, false, err
	}

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return uuid.Nil, false, err
	}
	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx, `
		INSERT INTO documents (course_id, filename, kind, status, storage_key,
		                       content_hash, size_bytes, year, term)
		VALUES ($1, $2, $3::document_kind, 'queued', $4, $5, $6, $7, $8::term)
		RETURNING id`,
		courseID, filename, kind, storageKey, contentHash, sizeBytes, year, term).Scan(&id)
	if err != nil {
		return uuid.Nil, false, err
	}

	if _, err := tx.Exec(ctx,
		`INSERT INTO ingest_jobs (document_id, status) VALUES ($1, 'queued')`, id); err != nil {
		return uuid.Nil, false, err
	}
	return id, false, tx.Commit(ctx)
}

func (s *Store) SetStatus(ctx context.Context, documentID uuid.UUID, status string, progress float64, detail string) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx,
		`UPDATE documents SET status = $2::ingest_status WHERE id = $1`, documentID, status); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx, `
		INSERT INTO ingest_jobs (document_id, status, progress, stage_detail, started_at, updated_at)
		VALUES ($1, $2::ingest_status, $3, $4, now(), now())
		ON CONFLICT (document_id) DO UPDATE
		SET status = EXCLUDED.status, progress = EXCLUDED.progress,
		    stage_detail = EXCLUDED.stage_detail, updated_at = now()`,
		documentID, status, progress, detail); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (s *Store) FailDocument(ctx context.Context, documentID uuid.UUID, reason string) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx,
		`UPDATE documents SET status = 'failed', error = $2 WHERE id = $1`, documentID, reason); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx, `
		UPDATE ingest_jobs SET status = 'failed', error = $2, completed_at = now(), updated_at = now()
		WHERE document_id = $1`, documentID, reason); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (s *Store) ChaptersFor(ctx context.Context, courseID uuid.UUID) ([]ingest.Chapter, error) {
	rows, err := s.pool.Query(ctx,
		`SELECT number, title, topics FROM chapters WHERE course_id = $1 ORDER BY number`, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := []ingest.Chapter{}
	for rows.Next() {
		var c ingest.Chapter
		if err := rows.Scan(&c.Number, &c.Title, &c.Topics); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

// ReplaceQuestions makes re-ingestion idempotent: the previous extraction for
// this document is deleted and rewritten in one transaction, so a re-run after
// a parser improvement never leaves two copies behind.
func (s *Store) ReplaceQuestions(ctx context.Context, documentID uuid.UUID, questions []ingest.StoredQuestion) ([]uuid.UUID, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	var courseID uuid.UUID
	var year *int
	var term *string
	if err := tx.QueryRow(ctx,
		`SELECT course_id, year, term FROM documents WHERE id = $1`, documentID).
		Scan(&courseID, &year, &term); err != nil {
		return nil, err
	}

	if _, err := tx.Exec(ctx, `DELETE FROM questions WHERE document_id = $1`, documentID); err != nil {
		return nil, err
	}

	// An exam row is derived from the document's year and term, so questions
	// can be browsed by sitting rather than by file.
	var examID *uuid.UUID
	if year != nil && term != nil {
		var id uuid.UUID
		// Both parameters are used twice with different types: $3 as the
		// integer year and inside a string concatenation, $4 as the term enum
		// and as text. Postgres cannot deduce one type for each, so every use
		// is cast explicitly.
		if err := tx.QueryRow(ctx, `
			INSERT INTO exams (course_id, document_id, year, term, title)
			VALUES ($1, $2, $3::int, $4::term, $3::text || ' ' || initcap($4::text))
			ON CONFLICT (course_id, year, term)
			DO UPDATE SET document_id = EXCLUDED.document_id
			RETURNING id`, courseID, documentID, *year, *term).Scan(&id); err != nil {
			return nil, err
		}
		examID = &id
	}

	ids := make([]uuid.UUID, len(questions))
	for i, q := range questions {
		var chapterID *uuid.UUID
		if q.ChapterNumber != 0 {
			var id uuid.UUID
			if err := tx.QueryRow(ctx,
				`SELECT id FROM chapters WHERE course_id = $1 AND number = $2`,
				courseID, q.ChapterNumber).Scan(&id); err == nil {
				chapterID = &id
			}
		}

		if err := tx.QueryRow(ctx, `
			INSERT INTO questions (course_id, exam_id, document_id, number, ordinal, text,
			                       marks, source_page, chapter_id, chapter_confidence,
			                       chapter_source, topic)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
			RETURNING id`,
			courseID, examID, documentID, q.Number, q.Ordinal, q.Text,
			q.Marks, q.SourcePage, chapterID, q.Confidence,
			nullIfEmpty(q.Source), nullIfEmpty(q.Topic)).Scan(&ids[i]); err != nil {
			return nil, fmt.Errorf("insert question %s: %w", q.Number, err)
		}
	}
	return ids, tx.Commit(ctx)
}

func (s *Store) ReplaceChunks(ctx context.Context, documentID uuid.UUID, chunks []ingest.StoredChunk) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var courseID uuid.UUID
	if err := tx.QueryRow(ctx, `SELECT course_id FROM documents WHERE id = $1`, documentID).
		Scan(&courseID); err != nil {
		return err
	}

	if _, err := tx.Exec(ctx, `DELETE FROM chunks WHERE document_id = $1`, documentID); err != nil {
		return err
	}

	// Question ids in document order, so a chunk can be linked to the question
	// it came from without the pipeline needing to know database ids.
	rows, err := tx.Query(ctx,
		`SELECT id FROM questions WHERE document_id = $1 ORDER BY ordinal`, documentID)
	if err != nil {
		return err
	}
	var questionIDs []uuid.UUID
	for rows.Next() {
		var id uuid.UUID
		if err := rows.Scan(&id); err != nil {
			rows.Close()
			return err
		}
		questionIDs = append(questionIDs, id)
	}
	rows.Close()

	for _, c := range chunks {
		var questionID *uuid.UUID
		if c.QuestionIndex >= 0 && c.QuestionIndex < len(questionIDs) {
			questionID = &questionIDs[c.QuestionIndex]
		}
		var chapterID *uuid.UUID
		if c.ChapterNumber != 0 {
			var id uuid.UUID
			if err := tx.QueryRow(ctx,
				`SELECT id FROM chapters WHERE course_id = $1 AND number = $2`,
				courseID, c.ChapterNumber).Scan(&id); err == nil {
				chapterID = &id
			}
		}

		if _, err := tx.Exec(ctx, `
			INSERT INTO chunks (course_id, document_id, question_id, chapter_id, ordinal,
			                    text, page, embedding, embedding_model, embedding_dim)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8::vector,$9,$10)`,
			courseID, documentID, questionID, chapterID, c.Ordinal,
			c.Text, c.Page, vectorLiteral(c.Embedding), "configured", len(c.Embedding)); err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
}

func (s *Store) SetQuestionEmbeddings(ctx context.Context, courseID uuid.UUID, ids []uuid.UUID, vectors [][]float32, model string) error {
	if len(ids) != len(vectors) {
		return fmt.Errorf("have %d question ids but %d vectors", len(ids), len(vectors))
	}
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for i, id := range ids {
		if _, err := tx.Exec(ctx, `
			INSERT INTO question_embeddings (question_id, course_id, embedding, embedding_model, embedding_dim)
			VALUES ($1, $2, $3::vector, $4, $5)
			ON CONFLICT (question_id) DO UPDATE
			SET embedding = EXCLUDED.embedding, embedding_model = EXCLUDED.embedding_model`,
			id, courseID, vectorLiteral(vectors[i]), model, len(vectors[i])); err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
}

func (s *Store) FinishDocument(ctx context.Context, documentID uuid.UUID, pageCount, questionCount int) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, `
		UPDATE documents SET status = 'indexed', page_count = $2, indexed_at = now(), error = NULL
		WHERE id = $1`, documentID, pageCount); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx, `
		UPDATE ingest_jobs
		SET status = 'indexed', progress = 1, stage_detail = NULL,
		    questions_extracted = $2, completed_at = now(), updated_at = now()
		WHERE document_id = $1`, documentID, questionCount); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func nullIfEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
