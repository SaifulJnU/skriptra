package api

import (
	"strings"
	"testing"

	"github.com/skriptra/skriptra/backend/internal/domain"
	"github.com/skriptra/skriptra/backend/internal/provider"
)

func TestPriorTurnsMapsRoles(t *testing.T) {
	got := priorTurns([]domain.Message{
		{Role: domain.RoleUser, Content: "What is the Gauss-Markov theorem?"},
		{Role: domain.RoleAssistant, Content: "It states that OLS is BLUE under the classical assumptions."},
		{Role: domain.RoleUser, Content: "Why?"},
	})

	if len(got) != 3 {
		t.Fatalf("got %d messages, want 3", len(got))
	}
	want := []provider.Role{provider.RoleUser, provider.RoleAssistant, provider.RoleUser}
	for i, w := range want {
		if got[i].Role != w {
			t.Fatalf("message %d: got role %q, want %q", i, got[i].Role, w)
		}
	}
}

// Replaying a refusal teaches the model that refusing is the expected shape of
// an answer, and it then refuses questions the corpus does cover.
func TestPriorTurnsDropsRefusals(t *testing.T) {
	got := priorTurns([]domain.Message{
		{Role: domain.RoleUser, Content: "Who teaches this course?"},
		{Role: domain.RoleAssistant, Content: "**No indexed passage covers that.**\n\nNothing in the uploaded material answers the question."},
		{Role: domain.RoleUser, Content: "What about the F-test?"},
	})

	for _, m := range got {
		if m.Role == provider.RoleAssistant {
			t.Fatalf("refusal was replayed: %q", m.Content)
		}
	}
	if len(got) != 2 {
		t.Fatalf("got %d messages, want the 2 user turns", len(got))
	}
}

// A previous answer can run to 800 tokens. Three of them would crowd out the
// retrieved passages the current answer has to be grounded in.
func TestPriorTurnsTruncatesLongAnswers(t *testing.T) {
	long := strings.Repeat("Schätzer ", 400)
	got := priorTurns([]domain.Message{
		{Role: domain.RoleUser, Content: "Explain it"},
		{Role: domain.RoleAssistant, Content: long},
	})

	answer := got[1].Content
	if len([]rune(answer)) > 610 {
		t.Fatalf("answer not truncated: %d runes", len([]rune(answer)))
	}
	if !strings.HasSuffix(answer, "...") {
		t.Fatal("truncated answer should be marked with an ellipsis")
	}
	// Byte slicing would split a multibyte rune here.
	if strings.Contains(answer, "�") {
		t.Fatal("truncation broke a multibyte character")
	}
}

func TestPriorTurnsSkipsEmpty(t *testing.T) {
	got := priorTurns([]domain.Message{
		{Role: domain.RoleUser, Content: "   "},
		{Role: domain.RoleUser, Content: "Real question"},
	})
	if len(got) != 1 {
		t.Fatalf("got %d messages, want 1", len(got))
	}
}
