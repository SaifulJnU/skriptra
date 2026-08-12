package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/skriptra/skriptra/backend/internal/ingest"
)

// Chapter taxonomy writes.
//
// Nothing created chapters before this: the five in the development course came
// from the seed, so a course created through the product could never classify
// anything. Classification scores a question against chapter vocabulary, and
// with no chapters there is nothing to score against.

// ReplaceChapters sets the taxonomy for a course.
//
// A replace rather than an append, because the taxonomy is a single coherent
// statement of what the course covers. Merging two versions of it would leave a
// course holding chapter 3 from one syllabus and chapter 3 from another.
//
// Existing chapter rows are updated in place rather than deleted and reinserted,
// so questions already classified keep their chapter_id and their assignment
// survives an edit to a title or a topic list. A chapter that disappears from
// the new taxonomy has its questions unclassified rather than left pointing at
// something the course no longer claims to cover.
func (s *Store) ReplaceChapters(ctx context.Context, courseID uuid.UUID, chapters []ingest.Chapter) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	numbers := make([]int, 0, len(chapters))
	for _, ch := range chapters {
		numbers = append(numbers, ch.Number)
	}

	// Questions first: the FK is ON DELETE SET NULL, but doing it explicitly
	// keeps the intent visible and covers the case where a chapter is kept but
	// renumbered.
	if _, err := tx.Exec(ctx, `
		UPDATE questions SET chapter_id = NULL, chapter_confidence = NULL
		WHERE course_id = $1
		  AND chapter_id IN (
		      SELECT id FROM chapters WHERE course_id = $1 AND number <> ALL($2::int[])
		  )`, courseID, numbers); err != nil {
		return err
	}

	if _, err := tx.Exec(ctx,
		`DELETE FROM chapters WHERE course_id = $1 AND number <> ALL($2::int[])`,
		courseID, numbers); err != nil {
		return err
	}

	for _, ch := range chapters {
		if _, err := tx.Exec(ctx, `
			INSERT INTO chapters (course_id, number, title, topics)
			VALUES ($1, $2, $3, $4)
			ON CONFLICT (course_id, number)
			DO UPDATE SET title = EXCLUDED.title, topics = EXCLUDED.topics`,
			courseID, ch.Number, ch.Title, ch.Topics); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

// UnclassifiedQuestions returns the questions in a course that have no chapter,
// so a newly saved taxonomy can be applied to material uploaded before it
// existed.
//
// This is the common case, not an edge one: a student uploads the papers they
// have, then finds the syllabus. Leaving those questions permanently
// unclassified would mean the taxonomy only ever applied to future uploads.
func (s *Store) UnclassifiedQuestions(ctx context.Context, courseID uuid.UUID) ([]QuestionText, error) {
	rows, err := s.pool.Query(ctx, `
		SELECT id, text FROM questions
		WHERE course_id = $1 AND chapter_id IS NULL
		ORDER BY created_at`, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := []QuestionText{}
	for rows.Next() {
		var q QuestionText
		if err := rows.Scan(&q.ID, &q.Text); err != nil {
			return nil, err
		}
		out = append(out, q)
	}
	return out, rows.Err()
}

// QuestionText is the minimum needed to classify: an identifier and the words.
type QuestionText struct {
	ID   uuid.UUID
	Text string
}

// SetQuestionChapter records a classification against the chapter with the
// given number in that course.
func (s *Store) SetQuestionChapter(ctx context.Context, questionID uuid.UUID, courseID uuid.UUID, number int, confidence float64) error {
	_, err := s.pool.Exec(ctx, `
		UPDATE questions
		SET chapter_id = (SELECT id FROM chapters WHERE course_id = $2 AND number = $3),
		    chapter_confidence = $4
		WHERE id = $1`, questionID, courseID, number, confidence)
	return err
}
