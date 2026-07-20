# Ch 3 — EXERCISES (fresh practice set, 38 points)

> These are **new** scenarios, not the Sheet 3/4/5 questions (those live in `22-BOOK-EXAMPLES-AND-SHEETS.md` with intuition-first walkthroughs). Do this set **closed-book, on paper, timed at 38 minutes** — 1 minute per point, the real exam pace. Then check `21-SOLUTIONS.md`.

---

## Part A — Estimation by hand (12 points)

A tiny dataset, $n=6$, one covariate:

| $x$ | 1 | 2 | 3 | 4 | 5 | 6 |
|---|---|---|---|---|---|---|
| $y$ | 3.1 | 4.0 | 5.2 | 6.8 | 8.1 | 9.0 |

**A1** (3 pts) Compute $\hat\beta_0$ and $\hat\beta_1$ by hand using $\hat\beta_1=S_{xy}/S_{xx}$, $\hat\beta_0=\bar y-\hat\beta_1\bar x$.

**A2** (2 pts) Compute the fitted values and residuals. Verify $\sum\hat\varepsilon_i \approx 0$.

**A3** (2 pts) Compute $\text{SSE}$, $\text{SST}$, and $R^2$.

**A4** (2 pts) Compute $\hat\sigma^2$ (the unbiased estimator) and $\widehat{\text{se}}(\hat\beta_1)$.

**A5** (2 pts) Test $H_0:\beta_1=0$ against $H_1:\beta_1\neq0$ at $\alpha=0.05$. Use $t_{4}(0.975)=2.776$.

**A6** (1 pt) Construct the 95% CI for $\beta_1$ and confirm it is consistent with your test decision in A5.

---

## Part B — Joint hypothesis testing (8 points)

A researcher models software defect count on 4 covariates (team size, code age, test coverage, number of reviewers) with an intercept, $n=40$. The unrestricted SSE is $850$. She wants to test whether **team size and code age together** can be dropped from the model ($r=2$ restrictions). Refitting without those two covariates gives $\text{SSE}_{H_0}=1020$.

**B1** (1 pt) State $H_0$ and $H_1$ in words.

**B2** (1 pt) Give $p$ (total parameters in the unrestricted model) and the residual degrees of freedom.

**B3** (3 pts) Compute the F-statistic.

**B4** (2 pts) The critical value is $F_{2,35}(0.95)=3.267$. State your decision and conclusion in one sentence, using the correct hedge language ("fail to reject" / "at least one").

**B5** (1 pt) Suppose instead a colleague tells you $F=-1.4$ came out of their calculation. Without redoing any arithmetic, explain why this must be wrong.

---

## Part C — Model choice: AIC, BIC, adjusted $R^2$ (10 points)

Two nested models are fit to the same $n=200$ observations, same response, same scale:

| Model | $p$ (incl. intercept) | SSE | SST |
|---|---|---|---|
| A | 4 | 980 | 1500 |
| B | 7 | 860 | 1500 |

**C1** (3 pts) Compute $R^2$ and $\bar R^2$ for both models.

**C2** (4 pts) Compute AIC and BIC for both models. (Use $\hat\sigma^2_{ML}=\text{SSE}/n$, natural log, penalty $2(p+1)$ for AIC and $\log(n)(p+1)$ for BIC.)

**C3** (2 pts) Which model does each criterion prefer? Is there disagreement?

**C4** (1 pt) Model B has 3 more parameters than Model A. In one sentence, explain why BIC agreeing with AIC here is stronger evidence than $\bar R^2$ agreeing.

---

## Part D — Diagnostics and short concepts (8 points)

**D1** (2 pts) An observation has raw residual $\hat\varepsilon_i=-45$, $\hat\sigma=60$, leverage $h_{ii}=0.35$, and the model has $p=5$ parameters. Compute the standardised residual $r_i$.

**D2** (2 pts) Using the same numbers, compute Cook's distance $D_i$. Is this observation likely to be flagged as influential (typical rule of thumb: $D_i > 4/n$, or simply "notably larger than the rest")?

**D3** (2 pts, T/F + one-line justification each)

(a) "A point with high leverage but a small residual is necessarily a problem for the fit."

(b) "If $\hat\beta_2>0$ in a multiple regression, then the marginal (simple, two-variable) correlation between $x_2$ and $y$ must also be positive."

**D4** (2 pts) A colleague says: "I checked — the mean of my residuals is exactly zero, so my model must be well specified." Explain in 2–3 sentences why this reasoning is flawed, and name the diagnostic that would actually catch misspecification.

---

## Self-check before opening the solutions

- [ ] Did you write the formula **before** plugging in numbers, every time?
- [ ] Did you round to 3 decimals only at the end?
- [ ] Did every T/F answer include a reason, not just a verdict?
- [ ] Did B5 use logic, not arithmetic?
