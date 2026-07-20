# Ch 2 — Suggestions: how to study this chapter

**Time budget: 4–5 hours (Days 1–2 of your plan).**
**Scope: Sections 2.1, 2.2 and 2.3 only.** Sections 2.4–2.10 are outside your reading course — skip them entirely.

---

## Why this chapter matters more than its page count suggests

Chapter 2 is the book's *tour* chapter: it shows you the models without the machinery. That makes it easy to underrate. Don't. Two things here are directly examined and both are cheap marks:

1. **Section 2.2 gives you the interpretation skills.** Slope, intercept, dummy variables, and what "holding others fixed" means. The exercise sheets (Sheet 1, Sheet 2, Sheet 5) are almost entirely interpretation, and interpretation questions appear in every past paper. These are the fastest marks on the whole exam.

2. **Section 2.3 (the Logit model) is guaranteed to appear.** Look at your papers:
   - Exam Summer 2025, Ex 1(h): a T/F on interpreting logit coefficients
   - Exam Summer 2025, Ex 4(a): *"Explain why a linear regression model is not appropriate for a binary dependent variable"* — 1 full point
   - It's in the WS papers too

   You are **not** asked to derive the logit model, estimate it, or do maximum likelihood for it. You are asked two questions only: *why not linear?* and *what does $\hat\beta_j$ mean?* Learn those two answers word-perfect and you cannot lose these marks.

---

## The order to work in

| Step | File | Time | Goal |
|---|---|---|---|
| 1 | `01-notes-2.1-introduction.md` | 20 min | The general framing: $E(y\mid\boldsymbol{x})$ |
| 2 | `02-notes-2.2.1-simple-linear-regression.md` | 45 min | Interpret slope and intercept perfectly |
| 3 | `03-notes-2.2.2-multiple-linear-regression.md` | 90 min | **Dummy variables.** This is the heaviest section here |
| 4 | `04-notes-2.3-logit-model.md` | 60 min | Two answers, word-perfect |
| 5 | `20-EXERCISES.md` → `21-SOLUTIONS.md` | 60 min | Then redo Sheets 1 and 2 |
| 6 | `10-SUMMARY.md`, `40-MIND-MAP.md`, `52-STORY-FOR-A-BABY.md` | 30 min | Lock it in |

---

## The single most important thing in this chapter

**Dummy variable coding.** If you learn one thing from Chapter 2, learn this:

> A categorical covariate with **$c$ levels** becomes **$c-1$ dummy variables**. One level is left out and becomes the **reference category**. Each dummy's coefficient is the difference *from the reference*, holding everything else fixed.

Why $c-1$ and not $c$? Because including all $c$ dummies plus an intercept makes the design matrix **singular** — the dummies sum to the intercept column. No unique OLS solution. This is called the **dummy variable trap**, and it is the reason Chapter 3's full-rank assumption exists.

**Past-paper evidence this is examined every single year:**
- *Linear_model_exam_sheet*, Block I(iv): "*k* levels ⟹ *k*−1 dummies" → TRUE
- *WS 23/24*, Block I(iv): "for *m* categories we need *m* dummies" → **FALSE**
- *Exam Summer 2025*, Ex 2(a): build a wage model with education (3 levels) and birthplace (2 levels) — 3 points
- *Sheet 1*, Ex 2: define dummies for 5 education levels, identify the reference category

That's four different papers testing one idea. Learn it once, cash it every year.

---

## What to skip

- ❌ Sections 2.4 (mixed models), 2.5 (nonparametric), 2.6 (additive), 2.7 (GAM), 2.8 (geoadditive), 2.9 (quantile/GAMLSS), 2.10 (nutshell summary of all of them). **Out of scope.**
- ❌ Any *estimation* of the logit model — no maximum likelihood, no Newton–Raphson, no derivations. Chapter 5 territory.
- ❌ The probit model beyond one sentence ("same idea, uses the normal CDF instead of the logistic — very similar results").
- ❌ The book's specific datasets and numbers.

---

## Self-check before Chapter 3

Can you, cold, with no notes:

- [ ] Write $E(y\mid\boldsymbol{x}) = \boldsymbol{x}'\boldsymbol\beta$ and say why the model is about the **mean** of $y$?
- [ ] Interpret $\hat\beta_1$ in a simple regression, in one sentence, with units?
- [ ] Say when $\hat\beta_0$ should **not** be interpreted, and why?
- [ ] Take "education has 5 levels" and write down the dummies, name the reference category, and interpret two coefficients?
- [ ] Compute the wage gap between two hypothetical people from a fitted model (Sheet 1, Ex 2(c) and 2(d))?
- [ ] Write the model with an interaction term and explain **geometrically** what the interaction does?
- [ ] Give **three** reasons a linear model fails for binary $y$?
- [ ] Write the logit model both ways — as $P(y=1)$ and as $\log\frac{p}{1-p}$?
- [ ] Say exactly what $\hat\beta_j$ means in a logit model, and what it does **not** mean?

Nine yeses → Chapter 3. Anything less, go back to that specific file.

---

## A warning about pace

Chapters 1 and 2 together are ~15% of the exam. Chapter 3 is ~85%. If you are on Day 4 and still in Chapter 2, you are behind — move on, and come back to Chapter 2's interpretation drills during Week 3 when you're revising. **Chapter 3 needs your freshest attention, not your leftovers.**
