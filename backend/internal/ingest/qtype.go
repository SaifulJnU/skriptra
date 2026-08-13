package ingest

import (
	"regexp"
	"strings"
)

// QuestionType is the format of a question, independent of its topic.
//
// Students revise by format as much as by chapter: "give me the true/false
// questions" is a normal request, and answering it with every question in the
// chapter is worse than answering nothing.
type QuestionType string

const (
	TypeTrueFalse      QuestionType = "true_false"
	TypeMultipleChoice QuestionType = "multiple_choice"
	TypeProof          QuestionType = "proof"
	TypeDerivation     QuestionType = "derivation"
	TypeComputation    QuestionType = "computation"
	TypeDiscussion     QuestionType = "discussion"
	TypeUnknown        QuestionType = "unknown"
)

// Rules again, for the same reasons as segmentation: exam wording for these
// formats is highly conventional, in both English and German, and a regex is
// deterministic, free and testable where a model call is none of those.
//
// Order matters. A question can say "prove or disprove, true or false" and the
// more specific format wins.
var typeRules = []struct {
	typ     QuestionType
	pattern *regexp.Regexp
}{
	// The instruction form: "decide whether the following are TRUE or FALSE".
	{TypeTrueFalse, regexp.MustCompile(`(?i)\b(true or false|true/false|wahr oder falsch|richtig oder falsch|state whether .{0,40}(true|false)|decide whether .{0,40}(true|false))\b`)},

	// The answer-box form. A real paper prints the instruction once, then puts
	// only empty TRUE/FALSE boxes under each of twelve statements. Matching the
	// instruction alone typed the parent exercise and left all twelve children
	// unknown, so a request for the true/false questions returned nothing.
	{TypeTrueFalse, regexp.MustCompile(`(?i)[\x{25A1}\x{2610}\x{2751}]\s*(?:TRUE|WAHR|RICHTIG)\b`)},
	{TypeTrueFalse, regexp.MustCompile(`(?i)\bTRUE\b[\s\x{25A1}\x{2610}\x{2751}\[\]]{0,12}\bFALSE\b`)},
	{TypeMultipleChoice, regexp.MustCompile(`(?i)\b(which of the following|select all that apply|choose the correct|tick the|welche der folgenden|kreuzen sie)\b`)},
	{TypeProof, regexp.MustCompile(`(?i)\b(prove|proof|disprove|show that|beweisen|zeigen sie, dass)\b`)},
	{TypeDerivation, regexp.MustCompile(`(?i)\b(derive|derivation|obtain an expression|herleiten|leiten sie)\b`)},
	{TypeComputation, regexp.MustCompile(`(?i)\b(compute|calculate|evaluate|find the value|determine the|berechnen|bestimmen sie)\b`)},
	{TypeDiscussion, regexp.MustCompile(`(?i)\b(discuss|explain|comment on|interpret|describe|justify|erklären|diskutieren|begründen)\b`)},
}

// ClassifyType assigns a format to a question.
//
// Returns unknown rather than guessing when nothing matches, on the same
// principle as chapter classification: an honest gap is recoverable, a
// confident mislabel silently corrupts every filtered result built on it.
func ClassifyType(text string) QuestionType {
	// Only the opening of a question carries the instruction. Later sentences
	// often say "explain your answer", which would drag almost everything into
	// discussion.
	head := text
	if len(head) > 220 {
		head = head[:220]
	}

	// Answer boxes are printed after the statement, not before it, so the
	// true/false patterns are matched against the whole text while the
	// instruction-verb patterns below stay anchored to the opening.
	for _, rule := range typeRules {
		if rule.typ == TypeTrueFalse && rule.pattern.MatchString(text) {
			return TypeTrueFalse
		}
	}

	for _, rule := range typeRules {
		if rule.pattern.MatchString(head) {
			return rule.typ
		}
	}
	return TypeUnknown
}

// ParseQuestionType maps the words a user might type onto a type.
// Returns unknown when the query does not name a format at all.
func ParseQuestionType(query string) QuestionType {
	q := strings.ToLower(query)
	switch {
	case regexp.MustCompile(`(?i)\b(true[ /-]?(or)?[ /-]?false|t/f|wahr[ /-]?falsch)\b`).MatchString(q):
		return TypeTrueFalse
	case strings.Contains(q, "multiple choice"), strings.Contains(q, "mcq"),
		strings.Contains(q, "multiple-choice"):
		return TypeMultipleChoice
	case regexp.MustCompile(`\b(proofs?|prove)\b`).MatchString(q):
		return TypeProof
	case regexp.MustCompile(`\b(derivations?|derive)\b`).MatchString(q):
		return TypeDerivation
	case regexp.MustCompile(`\b(computations?|calculations?|numerical)\b`).MatchString(q):
		return TypeComputation
	case regexp.MustCompile(`\b(discussions?|essay)\b`).MatchString(q):
		return TypeDiscussion
	}
	return TypeUnknown
}

// Label renders a type for display.
func (t QuestionType) Label() string {
	switch t {
	case TypeTrueFalse:
		return "true/false"
	case TypeMultipleChoice:
		return "multiple choice"
	case TypeProof:
		return "proof"
	case TypeDerivation:
		return "derivation"
	case TypeComputation:
		return "computation"
	case TypeDiscussion:
		return "discussion"
	}
	return "unclassified"
}
