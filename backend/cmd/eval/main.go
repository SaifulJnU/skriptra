// Command eval runs the retrieval evaluation against a live API and fails when
// quality regresses below the committed baseline.
//
//	go run ./cmd/eval                 # measure and gate
//	go run ./cmd/eval -update         # accept current numbers as the new floor
//
// It drives the real HTTP endpoint rather than calling packages directly, so it
// measures the system a user actually gets, including routing, filters and the
// prompt. That includes authentication: the harness signs in like any other
// client, and the account it uses must be a member of the course under test,
// or every case comes back as a 404 that looks like an empty corpus.
//
// Credentials come from EVAL_EMAIL and EVAL_PASSWORD, defaulting to the
// development account in dev/seed.sql.
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
		email    = flag.String("email", envOr("EVAL_EMAIL", "saiful@example.com"), "account to run as")
		password = flag.String("password", envOr("EVAL_PASSWORD", "skriptra-dev-password"), "account password")
	)
	flag.Parse()

	client := &apiClient{base: *apiURL, email: *email, password: *password}
	if err := client.signIn(); err != nil {
		fatal("signing in as %s: %v", *email, err)
	}

	g, err := eval.LoadGolden(*golden)
	if err != nil {
		fatal("loading golden set: %v", err)
	}

	results := make([]eval.Result, 0, len(g.Cases))
	for _, c := range g.Cases {
		got, err := client.ask(g.CourseID.String(), c.Question)
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
	Intent  string `json:"intent"`
	Answer  string `json:"answer"`
	Sources []struct {
		QuestionNumber string `json:"questionNumber"`
	} `json:"sources"`
	Questions []struct {
		Number  string `json:"number"`
		Chapter *struct {
			Number int `json:"number"`
		} `json:"chapter"`
	} `json:"questions"`
}

// apiClient holds the session the harness runs under.
//
// Access tokens last fifteen minutes and a full run over the golden set can
// take longer than that, because the explain cases each wait on a model. The
// client therefore signs in again on a 401 and retries the case once, rather
// than failing a run half way through for a reason that has nothing to do with
// retrieval quality.
type apiClient struct {
	base     string
	email    string
	password string
	token    string
}

func (c *apiClient) signIn() error {
	body, _ := json.Marshal(map[string]string{"email": c.email, "password": c.password})

	res, err := http.Post(c.base+"/auth/login", "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("login returned %d, is the account seeded and the API running?", res.StatusCode)
	}

	var session struct {
		AccessToken string `json:"accessToken"`
	}
	if err := json.NewDecoder(res.Body).Decode(&session); err != nil {
		return err
	}
	if session.AccessToken == "" {
		return fmt.Errorf("login returned no access token")
	}
	c.token = session.AccessToken
	return nil
}

func (c *apiClient) ask(courseID, question string) (eval.Retrieved, error) {
	return c.askOnce(courseID, question, true)
}

func (c *apiClient) askOnce(courseID, question string, retry bool) (eval.Retrieved, error) {
	body, _ := json.Marshal(map[string]any{
		"courseId": courseID, "question": question, "stream": false,
	})

	req, err := http.NewRequest(http.MethodPost, c.base+"/ask", bytes.NewReader(body))
	if err != nil {
		return eval.Retrieved{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)

	started := time.Now()
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return eval.Retrieved{}, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized && retry {
		if err := c.signIn(); err != nil {
			return eval.Retrieved{}, fmt.Errorf("re-authenticating mid-run: %w", err)
		}
		return c.askOnce(courseID, question, false)
	}

	if res.StatusCode == http.StatusNotFound {
		// Membership, not a missing route. A 404 here means the eval account is
		// not a member of the course under test, which would otherwise be
		// scored as a corpus that answers nothing.
		return eval.Retrieved{}, fmt.Errorf(
			"API returned 404 for course %s, is %s a member of it?", courseID, c.email)
	}
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
