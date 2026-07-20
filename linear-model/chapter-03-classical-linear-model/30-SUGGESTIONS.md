# Ch 3 — Suggestions: how to study the chapter that is the exam

**Time budget: 10 days of your 21 (Days 3–12).**
**This chapter is ~85% of your marks. Everything else is warm-up.**

---

## The shape of the chapter

| Section | Topic | Exam weight | Difficulty |
|---|---|---|---|
| **3.1** Model Definition | assumptions, dummies, transformations | ~20% | ⭐⭐⭐ |
| **3.2** Parameter Estimation | OLS, $\hat\sigma^2$, Gauss–Markov, BLUE | ~25% | ⭐⭐⭐⭐ |
| **3.3** Testing & Intervals | t-test, F-test, $C\beta=d$, CIs | ~25% | ⭐⭐⭐⭐⭐ |
| **3.4** Model Choice | AIC, BIC, $\bar R^2$, diagnostics | ~15% | ⭐⭐⭐ |

**3.3 is the hardest and the most heavily examined.** Budget accordingly: 3 full days on it, not 1.

---

## File order

| # | File | Time | Non-negotiable outcome |
|---|---|---|---|
| 1 | `01-notes-3.1.1-model-definition.md` | 60 min | Write $\boldsymbol{y}=\boldsymbol{X}\boldsymbol\beta+\boldsymbol\varepsilon$ with all dimensions in 60 sec |
| 2 | `02-notes-3.1.2-model-assumptions.md` | 90 min | **List all assumptions from memory**, and what breaks when each fails |
| 3 | `03-notes-3.1.3-covariate-effects.md` | 90 min | Dummies, polynomials, interactions, transformations |
| 4 | `04-notes-3.2.1-ols-estimation.md` | 120 min | **Derive $\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$ on a blank page in 4 min** |
| 5 | `05-notes-3.2.2-error-variance.md` | 45 min | $\hat\sigma^2$ — ML vs unbiased, and when each is used |
| 6 | `06-notes-3.2.3-properties-gauss-markov.md` | 90 min | State Gauss–Markov with all assumptions; explain BLUE word by word |
| 7 | `07-notes-3.3-hypothesis-testing.md` | 120 min | **Build $\boldsymbol{C}$ and $\boldsymbol{d}$ for any restriction in 60 sec** |
| 8 | `08-notes-3.3.1-exact-F-test.md` | 120 min | Both F formulas; know which inputs each needs |
| 9 | `09-notes-3.3.2-confidence-prediction-intervals.md` | 90 min | CI vs prediction interval — which is wider and why |
| 10 | `10-notes-3.4.1-bias-variance.md` | 45 min | Explain the tradeoff without formulas |
| 11 | `11-notes-3.4.2-model-choice-criteria.md` | 90 min | **Compute AIC and BIC from $\hat\varepsilon'\hat\varepsilon$ and $n$** |
| 12 | `12-notes-3.4.3-practical-model-choice.md` | 30 min | Forward/backward selection, cross-validation |
| 13 | `13-notes-3.4.4-model-diagnosis.md` | 90 min | Name what each residual plot detects |
| 14 | `20-EXERCISES.md` → `21-SOLUTIONS.md` | 180 min | Then redo Sheets 3, 4, 5 |

---

## The four things you must be able to do cold

If you can do these four, you pass comfortably. If you can't, nothing else saves you.

### ① Derive OLS on a blank page

$$\text{minimise } S(\boldsymbol\beta) = (\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta) \;\Longrightarrow\; \hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$$

**Appears in:** Exam Summer 2025 Ex 4(b) [2 pts], WS 23/24 Ex 2(b) [2 pts]. The WS 23/24 marking key literally says: *"1 point for correctly stating that RSS needs to be minimized. And 1 point for correctly deriving the solution."*

**Practise it 10 times.** Timed. Blank paper. It is the single most reliable 2 marks in the paper.

### ② Fill in a missing R output

Given a regression table with holes, reconstruct estimate / std. error / t-value / residual standard error.

**Appears in:** Exam Summer 2025 Ex 3(a) [2.5 pts], Example Exam LiMo Ex 1 [most of 35 pts].

The whole skill is three relationships:
$$t = \frac{\hat\beta}{\widehat{\text{se}}},\qquad \hat\sigma^2 = \frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n-p},\qquad \hat\beta_0=\bar y-\hat\beta_1\bar x$$

### ③ Build $\boldsymbol{C}$ and $\boldsymbol{d}$, then compute $F$

Turn a verbal hypothesis into $\boldsymbol{C}\boldsymbol\beta=\boldsymbol{d}$, count the restrictions $r$, compute $F$, compare to the quantile, decide.

**Appears in:** Exam Summer 2025 Ex 3(c)+(d) [3.5 pts], Sheet 4 (all three exercises), and every other paper.

This surprises people every year. It shouldn't surprise you.

### ④ Confidence interval + test decision

$$\hat\beta_j \pm t_{n-p}(1-\alpha/2)\cdot\widehat{\text{se}}(\hat\beta_j)$$

then answer "would you reject $H_0:\beta_j = c$?"

**Appears in:** every single paper.

---

## Study tactics specific to this chapter

**Do the derivations by hand, on paper, repeatedly.** Reading a derivation and reproducing one are completely different skills, and only the second is examined. The OLS derivation should become muscle memory.

**Build your own formula sheet as you go.** Don't just use mine (`99-exam-vault/10-FORMULA-SHEET.md`) — writing it yourself *is* the encoding. Then check yours against mine.

**Do every exercise sheet twice.** Sheets 3, 4 and 5 are essentially a past exam split into pieces, using the same running `Wage` model throughout. They are the closest thing you have to the real paper.

**Always write the formula before the numbers.** The marking keys explicitly award points for method. A right number with no formula can score less than a wrong number with the right formula.

**Track the notation.** This chapter is where the $p$ vs $k$ confusion does real damage — it propagates from degrees of freedom into every quantile lookup and every test decision. Re-read `chapter-01-introduction/03-notes-1.3-notation.md` §2 if you ever hesitate.

---

## What to skip

- ❌ **Section 3.5** (Bibliographic Notes and Proofs) — outside your scope entirely.
- ❌ The full algebraic proof of the Gauss–Markov theorem. **Know the statement, the assumptions, and what BLUE means.** WS 23/24 Ex 2(a) asks you to *describe* it for 4 points, with the key noting *"1 point for every assumption"* — that's a listing question, not a proof question.
- ❌ Deriving the F-statistic's distribution from quadratic forms. Know $F\sim F_{r,\,n-p}$ under $H_0$, and why $r$ and $n-p$ are what they are.
- ❌ Ridge / lasso / boosting / Bayesian — Chapter 4, out of scope. *(One caveat: a WS 22/23 T/F mentions the ridge estimator. Know one sentence — see `32-TRAPS.md`.)*
- ❌ Memorising Mallow's $C_p$ in detail. Know it exists and penalises complexity like AIC.

---

## Self-check before you start past papers (end of Day 12)

- [ ] Derive OLS on blank paper in under 4 minutes
- [ ] List all classical linear model assumptions, with what each one buys you
- [ ] State Gauss–Markov and explain each letter of BLUE
- [ ] Write $\hat\sigma^2$ in both ML and unbiased forms, and say which one AIC uses
- [ ] Given a verbal hypothesis, produce $\boldsymbol{C}$, $\boldsymbol{d}$ and $r$ in under a minute
- [ ] Compute $F$ from SSE and SSE$_{H_0}$ **and** from $R^2$
- [ ] Compute a CI for any $\hat\beta_j$ and decide any $H_0:\beta_j=c$
- [ ] Explain why a prediction interval is wider than a confidence interval
- [ ] Compute AIC and BIC from $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}$, $n$ and $|M|$
- [ ] Say what each of the four standard R diagnostic plots detects
- [ ] Fill missing values in an R regression table

**Eleven boxes. Tick them all before Day 13's mock.** Then Week 3 is about getting *fast*, not getting *competent* — which is exactly where you want to be.
