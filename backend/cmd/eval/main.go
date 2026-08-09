// Command eval runs the retrieval evaluation against a live API and fails when
// quality regresses below the committed baseline.
//
//	go run ./cmd/eval                 # measure and gate
//	go run ./cmd/eval -update         # accept current numbers as the new floor
//
// It drives the real HTTP endpoint rather than calling packages directly, so it
// measures the system a user actually gets, including routing, filters and the
// prompt.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/skriptra/skriptra/backend/internal/eval"
)

func main() {
	var (
		apiURL   = flag.String("api", envOr("EVAL_API_URL", "http://localhost:8080/api/v1"), "API base URL")
		golden   = flag.String("golden", "../eval/golden.json", "golden dataset")
		baseline = flag.String("baseline", "../eval/baseline.json", "baseline metrics")
		update   = flag.Bool("update", false, "write current results as the new baseline")
		verbose  = flag.Bool("v", false, "print every case")
	)
	flag.Parse()

	g, err := eval.LoadGolden(*golden)
	if err != nil {
		fatal("loading golden set: %v", err)
	}

	results := make([]eval.Result, 0, len(g.Cases))
	for _, c := range g.Cases {
		got, err := ask(*apiURL, g.CourseID.String(), c.Question)
		if err != nil {
			fatal("case %s: %v", c.ID, err)
		}
		r := eval.Score(c, got)
		results = append(results, r)

		if *verbose || !r.IntentCorrect || r.RecallAt5 < 1 || !r.RefusalOK {
			status := "ok"
			if !r.IntentCorrect || r.RecallAt5 < 1 || !r.RefusalOK {
				status = "FAIL"
			}
			fmt.Printf("%-4s %-28s intent=%-9s recall@5=%.2f %s\n",
				status, c.ID, got.Intent, r.RecallAt5, r.Notes)
		}
	}

	report := eval.Aggregate(results)
	fmt.Printf("\n%s\n", report)

	if *update {
		if err := eval.WriteBaseline(*baseline, report); err != nil {
			fatal("writing baseline: %v", err)
		}
		fmt.Printf("baseline updated: %s\n", *baseline)
		return
	}

	b, err := eval.LoadBaseline(*baseline)
	if err != nil {
		// No baseline yet is not a failure on a first run; it is an instruction.
		fmt.Printf("no baseline at %s. Run with -update to record one.\n", *baseline)
		return
	}

	if regressions := report.Regressions(*b); len(regressions) > 0 {
		fmt.Fprintf(os.Stderr, "\nREGRESSION\n")
		for _, r := range regressions {
			fmt.Fprintf(os.Stderr, "  %s\n", r)
		}
		fmt.Fprintf(os.Stderr, "\nIf this change is a deliberate trade-off, re-run with -update.\n")
		os.Exit(1)
	}
	fmt.Println("no regression against baseline")
}

// askResponse mirrors the non-streamed /ask body.
type askResponse struct {
	Intent    string `json:"intent"`
	Answer    string `json:"answer"`
	Sources   []struct {
		QuestionNumber string `json:"questionNumber"`
	} `json:"sources"`
	Questions []struct {
		Number  string `json:"number"`
		Chapter *struct {
			Number int `json:"number"`
		} `json:"chapter"`
	} `json:"questions"`
}

func ask(base, courseID, question string) (eval.Retrieved, error) {
	body, _ := json.Marshal(map[string]any{
		"courseId": courseID, "question": question, "stream": false,
	})

	started := time.Now()
	res, err := http.Post(base+"/ask", "application/json", bytes.NewReader(body))
	if err != nil {
		return eval.Retrieved{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return eval.Retrieved{}, fmt.Errorf("API returned %d", res.StatusCode)
	}

	var parsed askResponse
	if err := json.NewDecoder(res.Body).Decode(&parsed); err != nil {
		return eval.Retrieved{}, err
	}

	out := eval.Retrieved{Intent: parsed.Intent, Latency: time.Since(started)}

	// Refusal is detected from the answer text because it is a property of the
	// response the user sees, not a separate field. If the wording changes,
	// this must change with it, which is the correct coupling.
	lower := strings.ToLower(parsed.Answer)
	out.Refused = strings.Contains(lower, "no indexed passage") ||
		strings.Contains(lower, "nothing in the uploaded material")

	for _, s := range parsed.Sources {
		if s.QuestionNumber != "" {
			out.QuestionNumbers = append(out.QuestionNumbers, s.QuestionNumber)
		}
	}
	for _, q := range parsed.Questions {
		out.QuestionNumbers = append(out.QuestionNumbers, q.Number)
		if out.ChapterNumber == nil && q.Chapter != nil {
			n := q.Chapter.Number
			out.ChapterNumber = &n
		}
	}
	return out, nil
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func fatal(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
