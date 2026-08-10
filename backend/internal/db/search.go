package db

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/skriptra/skriptra/backend/internal/domain"
)

// HybridSearch is the `explain` retrieval path.
//
// Dense (pgvector cosine) and sparse (Postgres full-text) are ranked
// separately and fused with reciprocal rank fusion, all in one statement and
// one round trip.
//
// The important detail is `filtered`: structured filters are applied as a
// WHERE clause BEFORE ranking, not as post-filtering. At this corpus size a
// selective chapter filter cuts the candidate set to a few thousand rows,
// where an exact scan beats an approximate index on both latency and recall, // the filter is the optimisation, which is why a dedicated vector database
// buys nothing yet.
func (s *Store) HybridSearch(
	ctx context.Context,
	courseID uuid.UUID,
	queryText string,
	queryVector []float32,
	f domain.RetrievalFilters,
	limit int,
) ([]domain.SearchHit, error) {
	var chapterNumber *int
	if len(f.ChapterNumbers) > 0 {
		chapterNumber = &f.ChapterNumbers[0]
	}

	rows, err := s.pool.Query(ctx, `
WITH params AS (
    SELECT $1::uuid AS course_id, $2::text AS q_text, $3::vector AS q_vec,
           $4::int AS chapter_number, $5::int AS year_from, $6::int AS year_to,
           60::int AS k
),
filtered AS (
    SELECT c.*
    FROM chunks c
    JOIN params p          ON c.course_id = p.course_id
    LEFT JOIN chapters  ch ON ch.id = c.chapter_id
    LEFT JOIN documents d  ON d.id  = c.document_id
    WHERE (p.chapter_number IS NULL OR ch.number = p.chapter_number)
      AND (p.year_from      IS NULL OR d.year   >= p.year_from)
      AND (p.year_to        IS NULL OR d.year   <= p.year_to)
),
dense AS (
    SELECT f.id, row_number() OVER (ORDER BY f.embedding <=> p.q_vec) AS rnk,
           1 - (f.embedding <=> p.q_vec) AS score
    FROM filtered f CROSS JOIN params p
    WHERE f.embedding IS NOT NULL
    ORDER BY f.embedding <=> p.q_vec
    LIMIT 50
),
sparse AS (
    SELECT f.id,
           row_number() OVER (
               ORDER BY ts_rank_cd(f.search_tsv, plainto_tsquery('simple', p.q_text)) DESC
           ) AS rnk,
           ts_rank_cd(f.search_tsv, plainto_tsquery('simple', p.q_text)) AS score
    FROM filtered f CROSS JOIN params p
    WHERE f.search_tsv @@ plainto_tsquery('simple', p.q_text)
    ORDER BY score DESC
    LIMIT 50
),
fused AS (
    SELECT COALESCE(d.id, s.id) AS chunk_id,
           COALESCE(1.0 / (p.k + d.rnk), 0) + COALESCE(1.0 / (p.k + s.rnk), 0) AS rrf,
           COALESCE(d.score, 0) AS dense_score,
           COALESCE(s.score, 0) AS sparse_score
    FROM dense d
    FULL OUTER JOIN sparse s ON s.id = d.id
    CROSS JOIN params p
)
SELECT c.id, c.text, c.page,
       f.rrf * c.source_weight AS score,
       f.dense_score, f.sparse_score,
       d.id, d.filename, d.kind,
       q.id, coalesce(q.number, ''),
       e.year, e.term
FROM fused f
JOIN chunks    c ON c.id = f.chunk_id
JOIN documents d ON d.id = c.document_id
LEFT JOIN questions q ON q.id = c.question_id
LEFT JOIN exams     e ON e.id = q.exam_id
ORDER BY f.rrf * c.source_weight DESC
LIMIT $7`,
		courseID, queryText, vectorLiteral(queryVector),
		chapterNumber, f.YearFrom, f.YearTo, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := []domain.SearchHit{}
	for rows.Next() {
		var h domain.SearchHit
		var qID *uuid.UUID
		var qNum string
		var year *int
		var term *domain.Term

		if err := rows.Scan(&h.ChunkID, &h.Text, &h.Citation.Page, &h.Score,
			&h.DenseScore, &h.SparseScore,
			&h.Citation.DocumentID, &h.Citation.DocumentTitle, &h.Citation.DocumentKind,
			&qID, &qNum, &year, &term); err != nil {
			return nil, err
		}
		h.Citation.QuestionID = qID
		h.Citation.QuestionNumber = qNum
		h.Citation.Label = citationLabel(h.Citation.DocumentTitle, year, term, qNum, h.Citation.Page)
		out = append(out, h)
	}
	return out, rows.Err()
}

// citationLabel renders the string the UI shows, server-side, so every client
// (web today, mobile later) displays citations identically.
func citationLabel(filename string, year *int, term *domain.Term, questionNumber string, page int) string {
	title := filename
	if year != nil {
		season := ""
		if term != nil {
			if *term == domain.TermSummer {
				season = " Summer"
			} else {
				season = " Winter"
			}
		}
		title = fmt.Sprintf("%d%s Exam", *year, season)
	}
	if questionNumber != "" {
		return fmt.Sprintf("%s · Q%s · Page %d", title, questionNumber, page)
	}
	return fmt.Sprintf("%s · Page %d", title, page)
}
