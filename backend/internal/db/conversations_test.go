package db

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func TestConversationTitle(t *testing.T) {
	long := strings.Repeat("word ", 40)

	cases := []struct {
		name string
		in   string
		want string
	}{
		{"short question kept whole",
			"Why is the OLS estimator unbiased?",
			"Why is the OLS estimator unbiased?"},
		{"line breaks and runs of spaces collapse",
			"Explain\n\nthe   Gauss-Markov   theorem",
			"Explain the Gauss-Markov theorem"},
		{"empty stays empty", "", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := conversationTitle(tc.in); got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}

	t.Run("long question truncates at a word break", func(t *testing.T) {
		got := conversationTitle(long)
		if !strings.HasSuffix(got, "...") {
			t.Fatalf("expected an ellipsis, got %q", got)
		}
		if utf8.RuneCountInString(got) > 84 {
			t.Fatalf("title too long: %d runes", utf8.RuneCountInString(got))
		}
		if strings.HasSuffix(strings.TrimSuffix(got, "..."), " ") {
			t.Fatalf("trailing space before the ellipsis: %q", got)
		}
	})

	// The corpus is bilingual, so truncation has to be rune-safe. Slicing bytes
	// would split an umlaut and produce an invalid sequence in the sidebar.
	t.Run("multibyte text stays valid", func(t *testing.T) {
		got := conversationTitle(strings.Repeat("Schätzer für Störgrößen ", 10))
		if !utf8.ValidString(got) {
			t.Fatalf("invalid UTF-8: %q", got)
		}
	})
}
