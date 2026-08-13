package db

import (
	"context"
	"encoding/json"
	"strings"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/skriptra/skriptra/backend/internal/domain"
)

// Conversation persistence.
//
// The `conversations` and `messages` tables existed from the first migration
// and nothing wrote to them: `/ask` minted a fresh uuid per request and threw
// it away. The contract documents `conversationId` as "omit to start a new
// conversation", which promises that passing one continues it, so the server
// was quietly lying about a documented feature. This file makes it true.

// StartOrContinueConversation resolves the thread a turn belongs to.
//
// A supplied id is verified to belong to this user and this course before it
// is accepted. Without that check, passing someone else's conversation id
// would append to their thread and read back their history, which is the
// classic insecure-direct-object-reference. The identity is a stub today, the
// check is not, and it is the part that is expensive to retrofit.
//
// An id that does not resolve returns ErrNotFound rather than silently
// starting a new thread, so a client bug surfaces instead of scattering
// orphaned single-message conversations.
func (s *Store) StartOrContinueConversation(
	ctx context.Context, courseID, userID uuid.UUID, existing *uuid.UUID, firstQuestion string,
) (uuid.UUID, error) {
	if existing != nil {
		var id uuid.UUID
		err := s.pool.QueryRow(ctx, `
			SELECT id FROM conversations
			WHERE id = $1 AND user_id = $2 AND course_id = $3`,
			*existing, userID, courseID).Scan(&id)
		if err != nil {
			return uuid.Nil, norm(err)
		}
		return id, nil
	}

	var id uuid.UUID
	err := s.pool.QueryRow(ctx, `
		INSERT INTO conversations (course_id, user_id, title)
		VALUES ($1, $2, $3)
		RETURNING id`, courseID, userID, conversationTitle(firstQuestion)).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

// AppendMessage writes one turn and bumps the thread's updated_at so the
// history list orders by genuine recency rather than by when the thread was
// opened.
func (s *Store) AppendMessage(ctx context.Context, conversationID uuid.UUID, m domain.Message) (uuid.UUID, error) {
	sources := m.Sources
	if sources == nil {
		sources = []domain.Citation{}
	}
	sourcesJSON, err := json.Marshal(sources)
	if err != nil {
		return uuid.Nil, err
	}

	var usageJSON []byte
	if m.Usage != nil {
		if usageJSON, err = json.Marshal(m.Usage); err != nil {
			return uuid.Nil, err
		}
	}

	// The intent column is an enum, so an empty string is not a valid value.
	var intent *string
	if m.Intent != "" {
		v := string(m.Intent)
		intent = &v
	}

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback(ctx)

	var id uuid.UUID
	if err := tx.QueryRow(ctx, `
		INSERT INTO messages (conversation_id, role, content, intent, sources, usage)
		VALUES ($1, $2, $3, $4::query_intent, $5, $6)
		RETURNING id`,
		conversationID, m.Role, m.Content, intent, sourcesJSON, usageJSON).Scan(&id); err != nil {
		return uuid.Nil, err
	}

	if _, err := tx.Exec(ctx,
		`UPDATE conversations SET updated_at = now() WHERE id = $1`, conversationID); err != nil {
		return uuid.Nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

// ListConversations returns a user's threads for one course, most recently
// active first.
func (s *Store) ListConversations(
	ctx context.Context, courseID, userID uuid.UUID, page, pageSize int,
) ([]domain.Conversation, int, error) {
	var total int
	if err := s.pool.QueryRow(ctx, `
		SELECT count(*) FROM conversations
		WHERE course_id = $1 AND user_id = $2`, courseID, userID).Scan(&total); err != nil {
		return nil, 0, err
	}

	rows, err := s.pool.Query(ctx, `
		SELECT c.id, c.course_id, coalesce(c.title, ''), c.created_at, c.updated_at,
		       (SELECT count(*) FROM messages m WHERE m.conversation_id = c.id)
		FROM conversations c
		WHERE c.course_id = $1 AND c.user_id = $2
		ORDER BY c.updated_at DESC
		LIMIT $3 OFFSET $4`, courseID, userID, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	out := []domain.Conversation{}
	for rows.Next() {
		var c domain.Conversation
		if err := rows.Scan(&c.ID, &c.CourseID, &c.Title,
			&c.CreatedAt, &c.UpdatedAt, &c.MessageCount); err != nil {
			return nil, 0, err
		}
		out = append(out, c)
	}
	return out, total, rows.Err()
}

// GetConversation returns a thread with every message in it, oldest first.
func (s *Store) GetConversation(ctx context.Context, id, userID uuid.UUID) (*domain.ConversationDetail, error) {
	var d domain.ConversationDetail
	err := s.pool.QueryRow(ctx, `
		SELECT c.id, c.course_id, coalesce(c.title, ''), c.created_at, c.updated_at,
		       (SELECT count(*) FROM messages m WHERE m.conversation_id = c.id)
		FROM conversations c
		WHERE c.id = $1 AND c.user_id = $2`, id, userID).Scan(
		&d.ID, &d.CourseID, &d.Title, &d.CreatedAt, &d.UpdatedAt, &d.MessageCount)
	if err != nil {
		return nil, norm(err)
	}

	msgs, err := s.messages(ctx, id, 0)
	if err != nil {
		return nil, err
	}
	d.Messages = msgs
	return &d, nil
}

// RecentMessages returns the tail of a thread, oldest first, for use as model
// context on a follow-up question.
func (s *Store) RecentMessages(ctx context.Context, conversationID uuid.UUID, limit int) ([]domain.Message, error) {
	return s.messages(ctx, conversationID, limit)
}

// messages reads a thread. A limit of zero means all of it; a positive limit
// takes the most recent N and returns them in chronological order, which is
// the order a model needs them in.
func (s *Store) messages(ctx context.Context, conversationID uuid.UUID, limit int) ([]domain.Message, error) {
	query := `
		SELECT id, role, content, coalesce(intent::text, ''), sources, usage, created_at
		FROM messages
		WHERE conversation_id = $1
		ORDER BY created_at, id`
	args := []any{conversationID}
	if limit > 0 {
		query = `
			SELECT * FROM (
				SELECT id, role, content, coalesce(intent::text, ''), sources, usage, created_at
				FROM messages
				WHERE conversation_id = $1
				ORDER BY created_at DESC, id DESC
				LIMIT $2
			) t ORDER BY 7, 1`
		args = append(args, limit)
	}

	rows, err := s.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := []domain.Message{}
	for rows.Next() {
		var m domain.Message
		var intent string
		var sourcesJSON, usageJSON []byte
		if err := rows.Scan(&m.ID, &m.Role, &m.Content, &intent,
			&sourcesJSON, &usageJSON, &m.CreatedAt); err != nil {
			return nil, err
		}
		m.Intent = domain.QueryIntent(intent)
		m.Sources = []domain.Citation{}
		if len(sourcesJSON) > 0 {
			if err := json.Unmarshal(sourcesJSON, &m.Sources); err != nil {
				return nil, err
			}
		}
		if len(usageJSON) > 0 {
			var u domain.Usage
			if err := json.Unmarshal(usageJSON, &u); err != nil {
				return nil, err
			}
			m.Usage = &u
		}
		out = append(out, m)
	}
	return out, rows.Err()
}

// DeleteConversation removes a thread and its messages, which cascade.
func (s *Store) DeleteConversation(ctx context.Context, id, userID uuid.UUID) error {
	tag, err := s.pool.Exec(ctx,
		`DELETE FROM conversations WHERE id = $1 AND user_id = $2`, id, userID)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

// conversationTitle derives a thread label from its opening question.
//
// Truncation is on a rune boundary and at a word break, because the corpus is
// bilingual and cutting mid-rune on a German umlaut produces a broken
// character in the sidebar.
func conversationTitle(question string) string {
	title := strings.Join(strings.Fields(question), " ")
	const max = 80
	if utf8.RuneCountInString(title) <= max {
		return title
	}

	runes := []rune(title)[:max]
	if i := strings.LastIndex(string(runes), " "); i > max/2 {
		return strings.TrimRight(string(runes)[:i], " ,.;:") + "..."
	}
	return string(runes) + "..."
}
