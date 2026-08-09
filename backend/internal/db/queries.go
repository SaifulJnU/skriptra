package db

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/skriptra/skriptra/backend/internal/domain"
)

// ---------------------------------------------------------------- courses --

func (s *Store) ListCourses(ctx context.Context, page, pageSize int) ([]domain.Course, int, error) {
	var total int
	if err := s.pool.QueryRow(ctx, `SELECT count(*) FROM courses`).Scan(&total); err != nil {
		return nil, 0, err
	}

	rows, err := s.pool.Query(ctx, `
		SELECT c.id, c.name, coalesce(c.code, ''), coalesce(c.institution, ''),
		       c.language, c.created_at,
		       (SELECT count(*) FROM exams e     WHERE e.course_id = c.id),
		       (SELECT count(*) FROM questions q WHERE q.course_id = c.id)
		FROM courses c
		ORDER BY c.created_at DESC, c.name
		LIMIT $1 OFFSET $2`, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	out := []domain.Course{}
	for rows.Next() {
		var c domain.Course
		if err := rows.Scan(&c.ID, &c.Name, &c.Code, &c.Institution, &c.Language,
			&c.CreatedAt, &c.ExamCount, &c.QuestionCount); err != nil {
			return nil, 0, err
		}
		out = append(out, c)
	}
	return out, total, rows.Err()
}

func (s *Store) GetCourse(ctx context.Context, id uuid.UUID) (*domain.CourseDetail, error) {
	var d domain.CourseDetail
	var yearFrom, yearTo *int

	err := s.pool.QueryRow(ctx, `
		SELECT c.id, c.name, coalesce(c.code, ''), coalesce(c.institution, ''),
		       c.language, c.created_at,
		       (SELECT count(*) FROM exams     e WHERE e.course_id = c.id),
		       (SELECT count(*) FROM questions q WHERE q.course_id = c.id),
		       (SELECT count(*) FROM chapters  ch WHERE ch.course_id = c.id),
		       (SELECT count(*) FROM documents d WHERE d.course_id = c.id),
		       (SELECT min(e.year) FROM exams e WHERE e.course_id = c.id),
		       (SELECT max(e.year) FROM exams e WHERE e.course_id = c.id)
		FROM courses c WHERE c.id = $1`, id).
		Scan(&d.ID, &d.Name, &d.Code, &d.Institution, &d.Language, &d.CreatedAt,
			&d.ExamCount, &d.QuestionCount, &d.ChapterCount, &d.DocumentCount,
			&yearFrom, &yearTo)
	if err != nil {
		return nil, norm(err)
	}
	if yearFrom != nil && yearTo != nil {
		d.YearRange = &domain.YearRange{From: *yearFrom, To: *yearTo}
	}
	return &d, nil
}

func (s *Store) ListChapters(ctx context.Context, courseID uuid.UUID) ([]domain.Chapter, error) {
	rows, err := s.pool.Query(ctx, `
		SELECT ch.id, ch.number, ch.title, ch.topics,
		       (SELECT count(*) FROM questions q WHERE q.chapter_id = ch.id)
		FROM chapters ch
		WHERE ch.course_id = $1
		ORDER BY ch.number`, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := []domain.Chapter{}
	for rows.Next() {
		var c domain.Chapter
		if err := rows.Scan(&c.ID, &c.Number, &c.Title, &c.Topics, &c.QuestionCount); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

// ------------------------------------------------------------------ exams --

func (s *Store) ListExams(ctx context.Context, courseID uuid.UUID, page, pageSize int) ([]domain.Exam, int, error) {
	var total int
	if err := s.pool.QueryRow(ctx, `SELECT count(*) FROM exams WHERE course_id = $1`, courseID).Scan(&total); err != nil {
		return nil, 0, err
	}

	rows, err := s.pool.Query(ctx, `
		SELECT e.id, e.course_id, e.year, e.term, coalesce(e.title, ''), e.document_id,
		       (SELECT count(*) FROM questions q WHERE q.exam_id = e.id),
		       EXISTS (SELECT 1 FROM questions q
		               WHERE q.exam_id = e.id AND q.solution_text IS NOT NULL)
		FROM exams e
		WHERE e.course_id = $1
		ORDER BY e.year DESC, e.term
		LIMIT $2 OFFSET $3`, courseID, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	out := []domain.Exam{}
	for rows.Next() {
		var e domain.Exam
		if err := rows.Scan(&e.ID, &e.CourseID, &e.Year, &e.Term, &e.Title,
			&e.DocumentID, &e.QuestionCount, &e.HasSolutions); err != nil {
			return nil, 0, err
		}
		out = append(out, e)
	}
	return out, total, rows.Err()
}

func (s *Store) GetExam(ctx context.Context, examID uuid.UUID) (*domain.ExamDetail, error) {
	var d domain.ExamDetail
	err := s.pool.QueryRow(ctx, `
		SELECT e.id, e.course_id, e.year, e.term, coalesce(e.title, ''), e.document_id,
		       (SELECT count(*) FROM questions q WHERE q.exam_id = e.id),
		       EXISTS (SELECT 1 FROM questions q
		               WHERE q.exam_id = e.id AND q.solution_text IS NOT NULL)
		FROM exams e WHERE e.id = $1`, examID).
		Scan(&d.ID, &d.CourseID, &d.Year, &d.Term, &d.Title, &d.DocumentID,
			&d.QuestionCount, &d.HasSolutions)
	if err != nil {
		return nil, norm(err)
	}

	qs, _, err := s.ListQuestions(ctx, d.CourseID, domain.QuestionFilters{
		Page: 1, PageSize: 200, Sort: "chapter",
	}, &examID)
	if err != nil {
		return nil, err
	}
	d.Questions = qs
	return &d, nil
}

// -------------------------------------------------------------- questions --

const questionSelect = `
	SELECT q.id, q.exam_id, q.number, q.text, q.marks, q.source_page,
	       ch.id, ch.number, ch.title, q.chapter_confidence,
	       coalesce(q.topic, ''), e.year, e.term,
	       (q.solution_text IS NOT NULL)
	FROM questions q
	LEFT JOIN exams    e  ON e.id  = q.exam_id
	LEFT JOIN chapters ch ON ch.id = q.chapter_id`

func scanQuestions(rows pgx.Rows) ([]domain.Question, error) {
	out := []domain.Question{}
	for rows.Next() {
		var q domain.Question
		var chID *uuid.UUID
		var chNum *int
		var chTitle *string
		var conf *float64
		var year *int
		var term *domain.Term

		if err := rows.Scan(&q.ID, &q.ExamID, &q.Number, &q.Text, &q.Marks, &q.SourcePage,
			&chID, &chNum, &chTitle, &conf, &q.Topic, &year, &term, &q.HasSolution); err != nil {
			return nil, err
		}
		if chID != nil && chNum != nil && chTitle != nil {
			q.Chapter = &domain.ChapterRef{ID: *chID, Number: *chNum, Title: *chTitle, Confidence: conf}
		}
		q.Year = year
		if term != nil {
			q.Term = *term
		}
		out = append(out, q)
	}
	return out, rows.Err()
}

// ListQuestions is the `enumerate` path: exhaustive, ordered and paginated.
//
// Filters are composed as SQL predicates rather than applied after the fact —
// a request for every Chapter 3 question must return every Chapter 3 question,
// which top-k retrieval structurally cannot guarantee.
func (s *Store) ListQuestions(ctx context.Context, courseID uuid.UUID, f domain.QuestionFilters, examID *uuid.UUID) ([]domain.Question, int, error) {
	where := []string{"q.course_id = $1"}
	args := []any{courseID}
	add := func(clause string, v any) {
		args = append(args, v)
		where = append(where, fmt.Sprintf(clause, len(args)))
	}

	if examID != nil {
		add("q.exam_id = $%d", *examID)
	}
	if f.ChapterID != nil {
		add("q.chapter_id = $%d", *f.ChapterID)
	}
	if f.ChapterNumber != nil {
		add("ch.number = $%d", *f.ChapterNumber)
	}
	if f.YearFrom != nil {
		add("e.year >= $%d", *f.YearFrom)
	}
	if f.YearTo != nil {
		add("e.year <= $%d", *f.YearTo)
	}
	if f.Term != nil {
		add("e.term = $%d", *f.Term)
	}
	if f.Query != "" {
		add("q.search_tsv @@ plainto_tsquery('simple', $%d)", f.Query)
	}

	clause := " WHERE " + joinAnd(where)

	var total int
	countSQL := `SELECT count(*) FROM questions q
	             LEFT JOIN exams e ON e.id = q.exam_id
	             LEFT JOIN chapters ch ON ch.id = q.chapter_id` + clause
	if err := s.pool.QueryRow(ctx, countSQL, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	order := "e.year DESC, e.term, q.ordinal"
	switch f.Sort {
	case "oldest":
		order = "e.year ASC, e.term, q.ordinal"
	case "chapter":
		order = "ch.number NULLS LAST, e.year DESC, q.ordinal"
	}

	args = append(args, f.PageSize, (f.Page-1)*f.PageSize)
	sql := questionSelect + clause +
		fmt.Sprintf(" ORDER BY %s LIMIT $%d OFFSET $%d", order, len(args)-1, len(args))

	rows, err := s.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	qs, err := scanQuestions(rows)
	return qs, total, err
}

func (s *Store) GetQuestion(ctx context.Context, id uuid.UUID) (*domain.QuestionDetail, error) {
	var d domain.QuestionDetail
	var chID *uuid.UUID
	var chNum *int
	var chTitle *string
	var conf *float64
	var year *int
	var term *domain.Term
	var solText *string

	err := s.pool.QueryRow(ctx, `
		SELECT q.id, q.exam_id, q.number, q.text, q.marks, q.source_page,
		       ch.id, ch.number, ch.title, q.chapter_confidence,
		       coalesce(q.topic, ''), e.year, e.term,
		       q.document_id, q.solution_text, q.solution_page
		FROM questions q
		LEFT JOIN exams    e  ON e.id  = q.exam_id
		LEFT JOIN chapters ch ON ch.id = q.chapter_id
		WHERE q.id = $1`, id).
		Scan(&d.ID, &d.ExamID, &d.Number, &d.Text, &d.Marks, &d.SourcePage,
			&chID, &chNum, &chTitle, &conf, &d.Topic, &year, &term,
			&d.DocumentID, &solText, &d.SolutionSourcePage)
	if err != nil {
		return nil, norm(err)
	}
	if chID != nil && chNum != nil && chTitle != nil {
		d.Chapter = &domain.ChapterRef{ID: *chID, Number: *chNum, Title: *chTitle, Confidence: conf}
	}
	d.Year = year
	if term != nil {
		d.Term = *term
	}
	if solText != nil {
		d.SolutionText = *solText
		d.HasSolution = true
	}
	return &d, nil
}

// SimilarQuestions is the `similar` path: question-level k-NN excluding self.
func (s *Store) SimilarQuestions(ctx context.Context, questionID uuid.UUID, limit int, minScore float64) ([]domain.SimilarQuestion, error) {
	rows, err := s.pool.Query(ctx, `
		WITH ref AS (
			SELECT embedding, course_id FROM question_embeddings WHERE question_id = $1
		)
		SELECT q.id, q.exam_id, q.number, q.text, q.marks, q.source_page,
		       ch.id, ch.number, ch.title, q.chapter_confidence,
		       coalesce(q.topic, ''), e.year, e.term,
		       (q.solution_text IS NOT NULL),
		       1 - (qe.embedding <=> ref.embedding) AS score
		FROM question_embeddings qe
		JOIN ref ON qe.course_id = ref.course_id
		JOIN questions q  ON q.id  = qe.question_id
		LEFT JOIN exams    e  ON e.id  = q.exam_id
		LEFT JOIN chapters ch ON ch.id = q.chapter_id
		WHERE qe.question_id <> $1
		  AND 1 - (qe.embedding <=> ref.embedding) >= $2
		ORDER BY qe.embedding <=> ref.embedding
		LIMIT $3`, questionID, minScore, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := []domain.SimilarQuestion{}
	for rows.Next() {
		var q domain.Question
		var chID *uuid.UUID
		var chNum *int
		var chTitle *string
		var conf *float64
		var year *int
		var term *domain.Term
		var score float64

		if err := rows.Scan(&q.ID, &q.ExamID, &q.Number, &q.Text, &q.Marks, &q.SourcePage,
			&chID, &chNum, &chTitle, &conf, &q.Topic, &year, &term, &q.HasSolution, &score); err != nil {
			return nil, err
		}
		if chID != nil && chNum != nil && chTitle != nil {
			q.Chapter = &domain.ChapterRef{ID: *chID, Number: *chNum, Title: *chTitle, Confidence: conf}
		}
		q.Year = year
		if term != nil {
			q.Term = *term
		}
		out = append(out, domain.SimilarQuestion{Question: q, Score: score})
	}
	return out, rows.Err()
}

// -------------------------------------------------------------- analytics --

// ChapterFrequency is the `analyse` path: a pure aggregate, no model involved,
// so the numbers are exact rather than plausible.
func (s *Store) ChapterFrequency(ctx context.Context, courseID uuid.UUID, yearFrom, yearTo *int) (*domain.ChapterFrequencyResponse, error) {
	rows, err := s.pool.Query(ctx, `
		SELECT ch.id, ch.number, ch.title,
		       count(q.id)               AS question_count,
		       count(DISTINCT q.exam_id) AS exam_count
		FROM chapters ch
		LEFT JOIN questions q ON q.chapter_id = ch.id
		LEFT JOIN exams     e ON e.id = q.exam_id
		     AND ($2::int IS NULL OR e.year >= $2)
		     AND ($3::int IS NULL OR e.year <= $3)
		WHERE ch.course_id = $1
		GROUP BY ch.id, ch.number, ch.title
		ORDER BY ch.number`, courseID, yearFrom, yearTo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resp := &domain.ChapterFrequencyResponse{Data: []domain.ChapterFrequency{}}
	for rows.Next() {
		var f domain.ChapterFrequency
		if err := rows.Scan(&f.Chapter.ID, &f.Chapter.Number, &f.Chapter.Title,
			&f.QuestionCount, &f.ExamCount); err != nil {
			return nil, err
		}
		resp.Data = append(resp.Data, f)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Total is every question in the course, including unclassified ones —
	// shares must not silently sum to 100% while questions are missing.
	if err := s.pool.QueryRow(ctx,
		`SELECT count(*) FROM questions WHERE course_id = $1`, courseID).Scan(&resp.TotalQuestions); err != nil {
		return nil, err
	}
	if resp.TotalQuestions > 0 {
		for i := range resp.Data {
			resp.Data[i].Share = float64(resp.Data[i].QuestionCount) / float64(resp.TotalQuestions)
		}
	}

	byYear, err := s.chapterByYear(ctx, courseID)
	if err != nil {
		return nil, err
	}
	for i := range resp.Data {
		resp.Data[i].ByYear = byYear[resp.Data[i].Chapter.ID]
	}
	return resp, nil
}

func (s *Store) chapterByYear(ctx context.Context, courseID uuid.UUID) (map[uuid.UUID][]domain.YearCount, error) {
	rows, err := s.pool.Query(ctx, `
		SELECT q.chapter_id, e.year, count(*)
		FROM questions q
		JOIN exams e ON e.id = q.exam_id
		WHERE q.course_id = $1 AND q.chapter_id IS NOT NULL
		GROUP BY q.chapter_id, e.year
		ORDER BY e.year`, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := map[uuid.UUID][]domain.YearCount{}
	for rows.Next() {
		var id uuid.UUID
		var yc domain.YearCount
		if err := rows.Scan(&id, &yc.Year, &yc.QuestionCount); err != nil {
			return nil, err
		}
		out[id] = append(out[id], yc)
	}
	return out, rows.Err()
}

// -------------------------------------------------------------- documents --

func (s *Store) ListDocuments(ctx context.Context, courseID uuid.UUID, page, pageSize int) ([]domain.Document, int, error) {
	var total int
	if err := s.pool.QueryRow(ctx, `SELECT count(*) FROM documents WHERE course_id = $1`, courseID).Scan(&total); err != nil {
		return nil, 0, err
	}

	rows, err := s.pool.Query(ctx, `
		SELECT id, course_id, filename, kind, status, year, term,
		       page_count, size_bytes, content_hash, uploaded_at
		FROM documents WHERE course_id = $1
		ORDER BY uploaded_at DESC
		LIMIT $2 OFFSET $3`, courseID, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	out := []domain.Document{}
	for rows.Next() {
		var d domain.Document
		var term *domain.Term
		if err := rows.Scan(&d.ID, &d.CourseID, &d.Filename, &d.Kind, &d.Status,
			&d.Year, &term, &d.PageCount, &d.SizeBytes, &d.ContentHash, &d.UploadedAt); err != nil {
			return nil, 0, err
		}
		if term != nil {
			d.Term = *term
		}
		out = append(out, d)
	}
	return out, total, rows.Err()
}

func (s *Store) DocumentStatus(ctx context.Context, id uuid.UUID) (*domain.DocumentStatus, error) {
	var st domain.DocumentStatus
	var stage, errMsg *string
	err := s.pool.QueryRow(ctx, `
		SELECT d.id, d.status,
		       coalesce(j.progress, CASE WHEN d.status = 'indexed' THEN 1 ELSE 0 END),
		       j.stage_detail, coalesce(j.questions_extracted, 0), j.error
		FROM documents d
		LEFT JOIN ingest_jobs j ON j.document_id = d.id
		WHERE d.id = $1`, id).
		Scan(&st.DocumentID, &st.Status, &st.Progress, &stage, &st.QuestionsExtracted, &errMsg)
	if err != nil {
		return nil, norm(err)
	}
	if stage != nil {
		st.StageDetail = *stage
	}
	if errMsg != nil {
		st.Error = *errMsg
	}
	return &st, nil
}

func joinAnd(parts []string) string {
	out := parts[0]
	for _, p := range parts[1:] {
		out += " AND " + p
	}
	return out
}
