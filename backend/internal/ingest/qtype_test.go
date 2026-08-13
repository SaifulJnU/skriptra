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

// A real paper prints the instruction once and then gives each statement only a
// pair of answer boxes. Before this was handled, one exercise was typed
// true/false and the twelve statements under it were all unknown, so asking for
// the true/false questions returned the wrapper and none of the questions.
func TestClassifyTypeAnswerBoxes(t *testing.T) {
	cases := []struct {
		name string
		text string
	}{
		{"box glyph", "In a multiple linear regression model, adding a variable never decreases R squared. □ TRUE □ FALSE"},
		{"bare pair", "The OLS estimator is unbiased under the classical assumptions. TRUE FALSE"},
		{"german boxes", "Der Schaetzer ist erwartungstreu. □ WAHR □ FALSCH"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ClassifyType(tc.text); got != TypeTrueFalse {
				t.Fatalf("got %q, want %q", got, TypeTrueFalse)
			}
		})
	}
}

// The boxes must not swallow ordinary prose that happens to use both words.
func TestClassifyTypeAnswerBoxesNoFalsePositive(t *testing.T) {
	text := "Explain why a test can be true in the population but the estimator " +
		"still gives a false impression in a small sample, and discuss the consequences."
	if got := ClassifyType(text); got == TypeTrueFalse {
		t.Fatalf("prose classified as true/false")
	}
}
