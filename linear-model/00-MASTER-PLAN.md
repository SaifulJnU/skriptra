# 00 — MASTER PLAN: RC Linear Models → 60/60

**Course:** Reading Course Linear Models, TU Dortmund, Faculty of Statistics (Prof. Dr. Andreas Groll)
**Textbook:** Fahrmeir, Kneib, Lang, Marx — *Regression: Models, Methods and Applications* (Springer, 2013)
**Scope:** Chapter 1 · Chapter 2.1–2.3 · Chapter 3.1–3.4
**Exam:** 60 minutes · 60 points · non-programmable calculator + dictionary allowed
**Your timeline:** 3 weeks
**Your coach's promise:** if you finish this folder, there is no question in this exam you have not already seen.

---

## 1. What the exam actually is

I have read all five past papers you gave me. The structure is astonishingly stable:

| Block | Typical points | What it is |
|---|---|---|
| Exercise 1 | 6 | TRUE/FALSE, 12 statements (or 3 blocks of 4). **Negative marking within a block, floored at 0.** |
| Exercise 2 | 6–35 | Model building: write the equation, dummy variables, interpret coefficients |
| Exercise 3 | 6–25 | Computation: fill missing R output, CI, t-test, F-test, AIC/BIC |
| Exercise 4 | 8 | Explain/derive: OLS, Gauss–Markov, information criteria, logit, heteroscedasticity |

**60 minutes for 60 points = 1 minute per point.** This is the single most important fact about this exam. You do not have time to derive anything you have not already derived at home. Speed comes from recognition, not intelligence.

### Marks by chapter (my estimate from the five papers)

| Chapter | Share of marks |
|---|---|
| 1 — Introduction | ~2% |
| 2.1–2.3 — Regression Models (incl. Logit) | ~13% |
| **3.1 — Model Definition** | **~20%** |
| **3.2 — Parameter Estimation** | **~25%** |
| **3.3 — Testing & Confidence Intervals** | **~25%** |
| **3.4 — Model Choice & Diagnostics** | **~15%** |

Chapter 3 is the exam. Chapters 1 and 2 are the vocabulary you need to read Chapter 3.

---

## 2. Folder map

```
linear-model/
├── 00-MASTER-PLAN.md                 ← you are here
├── chapter-01-introduction/
├── chapter-02-regression-models/
├── chapter-03-classical-linear-model/
├── 99-exam-vault/                    ← formula sheet, question bank, T/F drill (used in Week 3, Days 15 & 20)
├── chatgpt-chat-outcomes.md          ← your earlier notes (archived)
├── liner-model-book.pdf
├── exercise-sheets/                  ← Sheets 1–5 (PDF)
└── prev-year-questions/              ← 5 past papers (PDF)
```

Every chapter folder has the same numbered spine, so you always know where you are. The number tells you the *kind* of file, not the order you read it in:

| File | Purpose | When to read it |
|---|---|---|
| `01…13-notes-*.md` | Subsection-by-subsection teaching notes | First pass, slowly |
| `10-SUMMARY.md` | One-page compression of the whole chapter | Second pass; then daily |
| `20-EXERCISES.md` | Problems in exam style | After first pass |
| `21-SOLUTIONS.md` | Full worked solutions | Only after you've attempted |
| `22-BOOK-EXAMPLES-AND-SHEETS.md` | Worked book examples + the sheet questions they map to | Alongside the notes |
| `23-PAST-PAPER-QUESTIONS.md` | **Every question from all 5 past papers that belongs to this chapter, with full solutions** | After the chapter, and again in Week 3 |
| `30-SUGGESTIONS.md` | How to study *this* chapter, what to skip | Before you start the chapter |
| `31-TRICKS-AND-TIPS.md` | Shortcuts, mental arithmetic, exam speed | Week 2–3 |
| `32-TRAPS.md` | The specific ways students lose marks here | Week 3, repeatedly |
| `33-FORMULA-DECISION-GUIDE.md` | *Which* formula, given what the question hands you | Week 2–3, and in the last 48 hours |
| `40-MIND-MAP.md` | The whole chapter as one tree | Whenever you feel lost |
| `50-REAL-LIFE-ANALOGY.md` | Concrete intuition | When a formula feels arbitrary |
| `51-PHILOSOPHY.md` | Why the subject is built this way | When you want it to *stick* |
| `52-STORY-FOR-A-BABY.md` | The whole chapter explained to a 5-year-old | Final check: can you tell the story? |
| `60-BANGLA-SUMMARY.md` | The chapter in Bangla | When the English is the bottleneck, not the maths |
| `70…72-visual-*.html` | Three single-idea diagrams — open in a browser | With the matching notes file |
| `80-ANIMATED-NARRATION.html` | The chapter as a click-through animation | End of the chapter, as a recap |

Chapter 3 has no `80-ANIMATED-NARRATION.html` yet — use `40-MIND-MAP.md` for that recap instead.

**The baby-story file is not a joke.** If you can tell the story without notes, you understand the chapter. If you can't, you have memorised formulas. Use it as your test.

---

## 3. The three-week schedule

Assume ~3 hours/day. Adjust the clock, not the order.

### Week 1 — Build the machine (Days 1–7)

| Day | Focus | Deliverable |
|---|---|---|
| 1 | Ch 1 (all), Ch 2.1–2.2 notes | Can define response/covariate, write a simple LM, interpret β̂₁ |
| 2 | Ch 2.3 Logit + Ch 2 exercises & solutions | Can say in 4 sentences why linear regression fails for binary y |
| 3 | Ch 3.1.1 + 3.1.2 (model, assumptions, residuals) | Can list all 5 assumptions from memory |
| 4 | Ch 3.1.3 — **dummy variables, polynomials, interactions** | Can write the wage model with education + birthplace cold |
| 5 | Ch 3.2.1 — OLS derivation | Can derive β̂ = (X′X)⁻¹X′y on a blank page in 4 minutes |
| 6 | Ch 3.2.2 + 3.2.3 — σ̂², Gauss–Markov, BLUE | Can state Gauss–Markov with all assumptions |
| 7 | **Consolidation.** Redo Sheets 1–3. Ch 1–3.2 summaries | No new material. Repair only. |

### Week 2 — Inference and selection (Days 8–14)

| Day | Focus | Deliverable |
|---|---|---|
| 8 | Ch 3.3 — t-test, general linear hypothesis Cβ = d | Can build C for any restriction in 60 seconds |
| 9 | Ch 3.3.1 — exact F-test, both formulas | Can compute F from SSE/SSE_H₀ *and* from R² |
| 10 | Ch 3.3.2 — confidence intervals, prediction intervals | Know the difference cold; know which is wider and why |
| 11 | Ch 3.4.1 + 3.4.2 — bias–variance, AIC, BIC, Mallow's Cp, R̄² | Can compute AIC/BIC from ε̂′ε̂ and n |
| 12 | Ch 3.4.3 + 3.4.4 — selection procedures, diagnostics, residual plots | Can name what each of the 4 standard R plots detects |
| 13 | Full past paper: **Exam Summer 2025**, closed book, 60 min | A score. Write it down. |
| 14 | **Consolidation.** Mark it brutally. Redo Sheets 4–5 | List of your top 5 weaknesses |

### Week 3 — Become fast and trap-proof (Days 15–21)

| Day | Focus | Deliverable |
|---|---|---|
| 15 | All `32-TRAPS.md` files + `99-exam-vault/30-TRUE-FALSE-DRILL.md` | 60+ T/F statements answered with reasons |
| 16 | Full past paper: **WiSe 2023/24**, closed book, 60 min | Compare to Day 13 score |
| 17 | Weakness repair (your Day 14 list) + `31-TRICKS-AND-TIPS.md` | Weaknesses closed |
| 18 | Full past paper: **RCLM WS22/23**, closed book, 60 min | Third data point |
| 19 | Full past paper: **Example Exam LiMo 2020** (the 35-point R-output monster) | Comfort with fill-in-the-blank R output |
| 20 | `99-exam-vault/10-FORMULA-SHEET.md` — write it out from memory 3× | A blank-page reproduction |
| 21 | Light. Re-read all 3 `10-SUMMARY.md` + all 3 `52-STORY-FOR-A-BABY.md`. Sleep. | Calm |

---

## 4. Scoring strategy inside the 60 minutes

1. **Minute 0–2: triage.** Read every exercise. Mark each as GREEN (I can do this now), AMBER (I can do this with thought), RED (I'm unsure). Do all GREEN first. Never let a RED question eat 15 minutes.
2. **TRUE/FALSE is scored per block with negative marking floored at zero.** Inside a block of 4: if you're sure of 3 and guessing 1, answer all 4 — the floor protects you. If you're sure of 0 and guessing 4, you still can't go below zero for that block, so **always answer everything**. Blank is never better than a guess here.
3. **"Provide sufficient reasons for your solutions."** This appears verbatim on the paper. A correct number with no working loses marks. A wrong number with correct working keeps most of them. **Always write the formula before the numbers.**
4. **Round to 3 decimals** — the Example Exam says so explicitly. Do it everywhere.
5. **Interpretation questions are free marks.** "Interpret the slope" is 1–2 points and takes 20 seconds if you have the sentence template memorised (see `chapter-02.../31-TRICKS-AND-TIPS.md`). Never skip these.
6. **When you're stuck on a computation, write down the formula and what you'd plug in.** Partial credit is real in this exam.

---

## 5. The single biggest notation trap in this course — read this now

The textbook and the exam papers use `p` differently. This has cost students marks in every past paper.

**Fahrmeir book convention (used in your lecture material):**
- `k` = number of **covariates**
- `p = k + 1` = number of **parameters** (covariates + intercept)
- Therefore: `rank(X) = p`, `t_j ~ t_{n−p}`, `F ~ F_{r, n−p}`, `R̄² = 1 − (n−1)/(n−p)·(1−R²)`

**Some exam papers instead use `p` = number of covariates**, and then write `t_j ~ t_{n−p−1}`.

**Both are the same number.** With k covariates plus an intercept, residual degrees of freedom is always:

> **df = n − (number of estimated β's) = n − k − 1 = n − p** (book convention)

**Rule for the exam:** never memorise "n−p" or "n−k−1" as a symbol string. Memorise: *count the betas you estimated, including the intercept, and subtract that from n.* Then read the paper's own notation and write it in their symbols. This one habit will save you 2–4 marks.

---

## 6. How to use me during these three weeks

Ask me for any of these at any time and I'll produce it on the spot:

- **"Quiz me on Chapter 3.3"** — I'll fire exam-style questions and mark you.
- **"Explain X three ways"** — derivation, intuition, and exam-answer version.
- **"Mark this"** — paste your attempt at a past-paper question, I'll grade it like Groll would.
- **"Make me 20 more T/F on model selection"** — extra drilling on demand.
- **"I have 40 minutes today, what do I do?"** — I'll pick the highest-value block.

---

## 7. The honest bit

You told me you're weak at maths. Here is what that actually means for this course: it means the first four days will feel worse than they should, and then it will click, because this subject is **one idea repeated at increasing levels of generality**:

> Fit a line by making squared errors as small as possible; then ask how confident you're allowed to be about it.

That's it. Everything — OLS, BLUE, t-tests, F-tests, AIC — is a variation on that sentence. You are not learning ten things. You are learning one thing ten times.

Start with `chapter-01-introduction/30-SUGGESTIONS.md`.
