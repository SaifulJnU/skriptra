package ingest

import "testing"

func TestClassifyType(t *testing.T) {
	cases := []struct {
		text string
		want QuestionType
	}{
		{"True or false: the OLS estimator is unbiased under heteroscedasticity. Justify your answer.", TypeTrueFalse},
		{"State whether the following statement is true or false and explain why.", TypeTrueFalse},
		{"Wahr oder falsch: der Schaetzer ist erwartungstreu.", TypeTrueFalse},

		{"Which of the following statements about leverage is correct?", TypeMultipleChoice},
		{"Welche der folgenden Aussagen ist richtig?", TypeMultipleChoice},

		{"State and prove the Gauss-Markov theorem.", TypeProof},
		{"Show that t squared equals F in this case.", TypeProof},

		{"Derive the ordinary least squares estimator for beta.", TypeDerivation},
		{"Leiten Sie den Kleinste-Quadrate-Schaetzer her.", TypeDerivation},

		{"Compute the residual sum of squares for the fitted model in Table 1.", TypeComputation},
		{"Calculate the leverage of observation 7.", TypeComputation},

		{"Discuss what has gone wrong and which diagnostics you would run.", TypeDiscussion},

		{"Write your matriculation number on every sheet.", TypeUnknown},
	}

	for _, tc := range cases {
		if got := ClassifyType(tc.text); got != tc.want {
			t.Errorf("ClassifyType(%.45q) = %q, want %q", tc.text, got, tc.want)
		}
	}
}

// A question ending "explain your answer" is still a true/false question. Only
// the opening carries the instruction, which is why the head is what gets
// matched.
func TestClassifyTypePrefersTheOpeningInstruction(t *testing.T) {
	got := ClassifyType("True or false: the Gauss-Markov theorem requires normality. Explain your answer and discuss the consequences.")
	if got != TypeTrueFalse {
		t.Errorf("got %q, want true_false despite the trailing 'explain'", got)
	}
}

func TestParseQuestionTypeFromUserQuery(t *testing.T) {
	cases := []struct {
		query string
		want  QuestionType
	}{
		{"true false question from chapter 2", TypeTrueFalse},
		{"give me all true/false questions", TypeTrueFalse},
		{"show me the T/F questions", TypeTrueFalse},
		{"any multiple choice questions in chapter 3", TypeMultipleChoice},
		{"list the proofs", TypeProof},
		{"all derivations from last year", TypeDerivation},
		// No format named: the filter must not be invented.
		{"all questions from chapter 2", TypeUnknown},
		{"why is MLE used here", TypeUnknown},
	}

	for _, tc := range cases {
		if got := ParseQuestionType(tc.query); got != tc.want {
			t.Errorf("ParseQuestionType(%q) = %q, want %q", tc.query, got, tc.want)
		}
	}
}
