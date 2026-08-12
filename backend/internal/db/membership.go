package db

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/skriptra/skriptra/backend/internal/domain"
)

// Course membership.
//
// The table and its role enum were in the first migration, because retrofitting
// authorization once a mobile client exists is expensive. This is where it
// starts being enforced.

// CreateCourse creates a course and makes its creator the owner.
//
// Both rows in one transaction. A course with no members is unreachable by
// anyone, including the person who just created it, so a crash between the two
// inserts must not be able to leave one behind.
func (s *Store) CreateCourse(
	ctx context.Context, userID uuid.UUID, name, code, institution, language string,
) (*domain.Course, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	c := domain.Course{
		Name:        strings.TrimSpace(name),
		Code:        strings.TrimSpace(code),
		Institution: strings.TrimSpace(institution),
		Language:    language,
	}
	if c.Language != "de" {
		c.Language = "en"
	}

	if err := tx.QueryRow(ctx, `
		INSERT INTO courses (name, code, institution, language, created_by)
		VALUES ($1, nullif($2, ''), nullif($3, ''), $4, $5)
		RETURNING id, created_at`,
		c.Name, c.Code, c.Institution, c.Language, userID).Scan(&c.ID, &c.CreatedAt); err != nil {
		return nil, err
	}

	if _, err := tx.Exec(ctx, `
		INSERT INTO course_members (course_id, user_id, role)
		VALUES ($1, $2, 'owner')`, c.ID, userID); err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	return &c, nil
}

// IsCourseMember reports whether a user may see a course at all.
//
// One indexed primary-key lookup, which is why it can afford to run on every
// course-scoped request rather than being cached into a token where it would
// go stale the moment membership changed.
func (s *Store) IsCourseMember(ctx context.Context, courseID, userID uuid.UUID) (bool, error) {
	var exists bool
	err := s.pool.QueryRow(ctx, `
		SELECT EXISTS (
			SELECT 1 FROM course_members WHERE course_id = $1 AND user_id = $2
		)`, courseID, userID).Scan(&exists)
	return exists, err
}

// CourseIDForExam and CourseIDForDocument let the membership check run on
// endpoints addressed by a child identifier. Without them, a question id
// leaked from another user's course would read straight through.
func (s *Store) CourseIDForExam(ctx context.Context, examID uuid.UUID) (uuid.UUID, error) {
	var id uuid.UUID
	err := s.pool.QueryRow(ctx, `SELECT course_id FROM exams WHERE id = $1`, examID).Scan(&id)
	return id, norm(err)
}

func (s *Store) CourseIDForDocument(ctx context.Context, documentID uuid.UUID) (uuid.UUID, error) {
	var id uuid.UUID
	err := s.pool.QueryRow(ctx, `SELECT course_id FROM documents WHERE id = $1`, documentID).Scan(&id)
	return id, norm(err)
}
