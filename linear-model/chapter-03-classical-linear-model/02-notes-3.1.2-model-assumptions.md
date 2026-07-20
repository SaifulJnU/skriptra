# 3.1.2 — Discussion of Model Assumptions

> **Learn this section as a list you can recite.** WS 23/24 Exercise 2(a) asks you to describe the Gauss–Markov theorem "and the assumptions" for **4 points**, with the marking key stating *"1 point for every assumption."* This is the most literally memorisable set of marks in the entire paper.

---

## 1. The assumptions of the classical linear model

$$\boldsymbol{y}=\boldsymbol{X}\boldsymbol\beta+\boldsymbol\varepsilon$$

| # | Assumption | Formal statement | What it buys you |
|---|---|---|---|
| **A1** | **Linearity** (correct specification) | $E(\boldsymbol{y}\mid\boldsymbol{X}) = \boldsymbol{X}\boldsymbol\beta$ | $\hat{\boldsymbol\beta}$ is **unbiased** |
| **A2** | **Zero mean errors** | $E(\varepsilon_i)=0$ for all $i$ | unbiasedness (free if intercept present) |
| **A3** | **Homoscedasticity** | $\text{Var}(\varepsilon_i)=\sigma^2$ for all $i$ | correct standard errors; **efficiency** |
| **A4** | **No autocorrelation** | $\text{Cov}(\varepsilon_i,\varepsilon_j)=0$, $i\neq j$ | correct standard errors; **efficiency** |
| **A5** | **Full column rank** | $\text{rank}(\boldsymbol{X})=p$, $\boldsymbol{X}$ non-stochastic | $\hat{\boldsymbol\beta}$ **exists and is unique** |
| **A6** | **Normality** *(extra)* | $\boldsymbol\varepsilon\sim N(\boldsymbol{0},\sigma^2\boldsymbol{I})$ | **exact** $t$- and $F$-tests, exact CIs |

**A3 and A4 combine into one compact statement:**

$$\boxed{\;\text{Cov}(\boldsymbol\varepsilon)=E(\boldsymbol\varepsilon\boldsymbol\varepsilon') = \sigma^2\boldsymbol{I}_n\;}$$

Constant diagonal (homoscedastic), zero off-diagonal (uncorrelated). One matrix, two assumptions.

**With A6, the full model statement is:**

$$\boldsymbol{y}\sim N(\boldsymbol{X}\boldsymbol\beta,\ \sigma^2\boldsymbol{I}_n)$$

---

## 2. 🔑 The layered structure — this is the key insight

The assumptions are **not** a flat list. They come in tiers, and each tier buys you strictly more. **Knowing which assumption buys which property is what separates a 2-mark answer from a 4-mark answer.**

```
   A1, A2, A5  ────────────►  β̂ EXISTS, is UNIQUE, and is UNBIASED
      │
      │  + A3, A4 (Cov(ε) = σ²I)
      ▼
   GAUSS–MARKOV  ───────────►  β̂ is BLUE (minimum variance among
      │                         linear unbiased estimators)
      │                         AND σ̂² = ε̂'ε̂/(n−p) is unbiased
      │
      │  + A6 (normality)
      ▼
   EXACT INFERENCE  ─────────►  β̂ ~ N(β, σ²(X'X)⁻¹) EXACTLY
                               t-tests, F-tests, CIs are EXACT
                               OLS = ML estimator
```

**Three sentences that pay:**

> **Unbiasedness needs only A1, A2, A5. It does not need homoscedasticity, independence, or normality.**
>
> **BLUE needs A1–A5. It does not need normality.**
>
> **Exact tests need all six. Without A6 the tests are only asymptotically valid.**

This layering answers a large family of exam questions at once. Whenever a T/F statement says "if [assumption] fails, then [property] is lost," check which tier the property lives in.

---

## 3. Each assumption in detail

### A1 — Linearity / correct specification

$E(y_i\mid\boldsymbol{x}_i) = \boldsymbol{x}_i'\boldsymbol\beta$. The systematic part really is a linear combination of the *included* covariates.

**Remember:** linear in $\boldsymbol\beta$, not in $\boldsymbol{x}$. Squares, logs, dummies and interactions are all fine (Section 3.1.3).

**If it fails:** $\hat{\boldsymbol\beta}$ is **biased** and the model is misspecified. This is the most serious failure, because no amount of extra data fixes it.

**Two ways it fails:**
- **Wrong functional form** — true relationship is curved, you fitted a line. *Detected by:* residuals-vs-fitted plot showing a systematic curve.
- **Omitted variable bias** — a relevant covariate is left out **and** it's correlated with an included one. Then $\hat\beta_j$ absorbs part of the omitted variable's effect.
  > This is the confounding story with a formula: bias $\propto$ (effect of omitted variable) × (correlation between omitted and included).
  >
  > ⚠️ **Note the "and."** *Exam Summer 2025, Ex 1(a):* adding a variable **uncorrelated with the response** doesn't affect unbiasedness but may affect variance → **TRUE**.

**Fix:** add polynomial terms, transform variables, include the missing covariate.

---

### A2 — $E(\varepsilon_i)=0$

**If the model has an intercept, this is essentially free.** Any constant shift $E(\varepsilon)=c\neq0$ gets absorbed: $\beta_0^{\text{new}} = \beta_0+c$. Only the intercept is affected, and it's rarely of interest.

The assumption becomes substantive as $E(\varepsilon_i\mid\boldsymbol{x}_i)=0$ — errors uncorrelated with the covariates. **That** is what fails under omitted variable bias, and it's really A1 in disguise.

---

### A3 — Homoscedasticity: $\text{Var}(\varepsilon_i)=\sigma^2$

Every observation has the **same** error variance, regardless of $\boldsymbol{x}_i$.

**Violation = heteroscedasticity.** Typical causes:
- Response is a positive, skewed quantity (income, revenue, rent) where spread grows with level
- Aggregated data (city averages have smaller variance than individual observations)
- Binary response: $\text{Var}=\pi(1-\pi)$ — heteroscedastic **by construction** (Section 2.3)

**Consequences — memorise this pair:**

| | Effect |
|---|---|
| $\hat{\boldsymbol\beta}$ | ✅ still **unbiased** and consistent (A1, A2, A5 untouched) |
| $\hat{\boldsymbol\beta}$ | ❌ no longer **BLUE** — a weighted estimator (GLS/WLS) has lower variance |
| $\widehat{\text{se}}(\hat\beta_j)$ | ❌ **biased** — usual formula $\hat\sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}$ is wrong |
| $t$- and $F$-tests, CIs | ❌ **invalid** — wrong size, wrong coverage |

> 🔴 **Exam Summer 2025, Ex 4(e)** [1 point]: *"the variation in revenue grows as the number of employees grows. Which impact does this have on the OLS estimate in terms of bias and efficiency?"*
>
> **Model answer:** *This is heteroscedasticity — $\text{Var}(\varepsilon_i)$ is not constant but increases with the covariate. The OLS estimator remains **unbiased** (and consistent), since unbiasedness requires only correct specification and $E(\varepsilon)=0$. However it is no longer **efficient**: the Gauss–Markov theorem no longer applies, so OLS is not BLUE, and a weighted least squares estimator would have smaller variance. In addition, the usual standard errors are biased, so the resulting $t$-tests, $F$-tests and confidence intervals are invalid.*
>
> That paragraph is a full-mark answer. The two words the marker is looking for are **unbiased** and **inefficient**.

> 🔴 *Linear_model_exam_sheet, Block II(iv):* "In the presence of heteroskedasticity, traditional standard errors remain valid for hypothesis testing and CIs." → **FALSE**.

**Detected by:** residuals-vs-fitted plot showing a **funnel/fan**; scale–location plot with an upward trend.
**Fixed by:** transform $y$ (log), weighted least squares, or robust (Huber–White) standard errors.

---

### A4 — No autocorrelation: $\text{Cov}(\varepsilon_i,\varepsilon_j)=0$

Errors of different observations are uncorrelated.

**Typically violated in:** time series (today's shock resembles yesterday's), spatial data (neighbouring regions), clustered data (students within a school, repeated measures on the same person).

**Consequences: identical in structure to heteroscedasticity.**

| | Effect |
|---|---|
| $\hat{\boldsymbol\beta}$ | ✅ unbiased, consistent |
| $\hat{\boldsymbol\beta}$ | ❌ not BLUE — inefficient |
| Standard errors | ❌ biased (usually **too small** with positive autocorrelation ⟹ over-confident, spuriously significant results) |

> 🔴 *WS 23/24, Block II(ii):* "When the residuals are not independent, their correlation does not affect the **consistency** of the coefficient estimates, but it may affect the **efficiency**." → **TRUE.** Exactly the pattern above. Learn this sentence — it's a template answer for both A3 and A4 violations.

**Detected by:** residuals plotted against time/index showing runs or cycles; Durbin–Watson test.

---

### A5 — Full column rank: $\text{rank}(\boldsymbol{X})=p$

The columns of $\boldsymbol{X}$ are linearly independent. Equivalently $\boldsymbol{X}'\boldsymbol{X}$ is **non-singular** (invertible, positive definite).

**Necessary condition:** $n\geq p$. You cannot estimate more parameters than you have observations.

**How it fails — perfect (exact) multicollinearity:**
- The **dummy variable trap** — all $c$ dummies plus an intercept
- A covariate that is an exact linear combination of others (e.g. including `age`, `age at hire`, and `years employed` when the third is the difference of the first two)
- Including the same variable twice under different names

**If it fails:** $(\boldsymbol{X}'\boldsymbol{X})^{-1}$ doesn't exist, the normal equations have **infinitely many** solutions, and $\hat{\boldsymbol\beta}$ is **not identified**.

> 🔴 *Exam Summer 2025, Ex 1(d):* "When $\boldsymbol{X}$ does not have full column rank, the OLS estimates still exist and are unique as long as the error variance is constant." → **FALSE.** Not unique; and homoscedasticity is irrelevant to identification. Watch for "as long as…" clauses that attach an unrelated condition.

**Rank facts for T/F questions** (book convention, $p$ = parameters):

| Claim | Verdict |
|---|---|
| $\text{rank}(\boldsymbol{X}'\boldsymbol{X})=p$ | ✅ TRUE |
| $\text{rank}(\boldsymbol{X}'\boldsymbol{X})=k$ | ❌ FALSE (one short — forgot the intercept) |
| $\text{rank}(\boldsymbol{X}'\boldsymbol{X})=n$ | ❌ FALSE — $\boldsymbol{X}'\boldsymbol{X}$ is $p\times p$; **rank cannot exceed the smaller dimension** |

### ⚠️ Near-multicollinearity — different problem, don't confuse them

| | **Perfect** multicollinearity | **Near** multicollinearity |
|---|---|---|
| Columns are | exactly dependent | highly but not perfectly correlated |
| Violates A5? | ✅ Yes | ❌ **No** — A5 still holds |
| $(\boldsymbol{X}'\boldsymbol{X})^{-1}$ | doesn't exist | exists but is **ill-conditioned** |
| $\hat{\boldsymbol\beta}$ | not identified | **unbiased**, still BLUE |
| Standard errors | — | **inflated** ⟹ wide CIs, insignificant $t$'s |
| Diagnostic | — | **VIF** $=1/(1-R_j^2)$ |

> 🔴 *Exam Summer 2025, Ex 1(j):* "Multicollinearity can inflate the variance of the OLS coefficient estimators." → **TRUE.**
> 🔴 *Linear_model_exam_sheet, II(i):* "Multicollinearity can inflate the variance but does not bias the estimates." → **TRUE.**
> 🔴 *WS 23/24, II(i):* "When highly correlated explanatory variables are present, this may lead to a **reduction** in the standard errors." → **FALSE** — it *inflates* them.
> 🔴 *Linear_model_exam_sheet, II(ii):* "If VIF for all variables is close to 1, multicollinearity is likely a concern." → **FALSE.** VIF ≈ 1 means **no** collinearity. Concern starts around VIF > 5 or 10.

**Memory hook:** *near-collinearity doesn't lie to you, it just refuses to give you a precise answer.* Unbiased, but imprecise.

---

### A6 — Normality: $\boldsymbol\varepsilon\sim N(\boldsymbol{0},\sigma^2\boldsymbol{I})$

The **extra** assumption. Everything up to and including BLUE holds without it.

**What it buys:**

$$\hat{\boldsymbol\beta}\sim N\!\left(\boldsymbol\beta,\ \sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}\right)\ \textbf{exactly}$$

$$\frac{\hat\beta_j-\beta_j}{\widehat{\text{se}}(\hat\beta_j)}\sim t_{n-p}\ \textbf{exactly}, \qquad F\sim F_{r,n-p}\ \textbf{exactly}$$

And: **OLS coincides with the maximum likelihood estimator.**

> 🔴 *Exam Summer 2025, Ex 1(l):* "The OLS estimator is equivalent to the ML estimator under iid normal errors." → **TRUE.**
>
> *Why:* the Gaussian log-likelihood is
> $$\ell(\boldsymbol\beta,\sigma^2) = -\tfrac{n}{2}\log(2\pi\sigma^2)-\tfrac{1}{2\sigma^2}(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)$$
> $\boldsymbol\beta$ appears **only** in the final quadratic term, with a minus sign. Maximising over $\boldsymbol\beta$ is therefore exactly minimising $(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)$ — the least squares criterion. **Same estimator, two justifications.**
>
> *This is also the answer to Sheet 3, Ex 2(a): "Assuming normally distributed errors, what is the ML estimate for $\beta$?" → **the same as the LS estimate**, $\hat{\boldsymbol\beta}_{ML}=\hat{\boldsymbol\beta}_{LS}$.*

**If normality fails:** point estimates and BLUE are untouched. Only the **exactness** of tests is lost. For large $n$ the Central Limit Theorem rescues you — $t$- and $F$-tests remain **asymptotically** valid. This is why mild non-normality is usually tolerable and severe skew is not.

**Detected by:** QQ plot of (standardised) residuals; histogram of residuals.
**Fixed by:** transform $y$; or rely on large-sample results.

---

## 4. The consequence table — learn this grid

| Assumption violated | $\hat{\boldsymbol\beta}$ unbiased? | BLUE? | se's valid? | Tests valid? |
|---|---|---|---|---|
| **A1** linearity | ❌ **NO** | ❌ | ❌ | ❌ |
| **A3** homoscedasticity | ✅ yes | ❌ no | ❌ no | ❌ no |
| **A4** independence | ✅ yes | ❌ no | ❌ no | ❌ no |
| **A5** full rank | — *(doesn't exist)* | — | — | — |
| **A6** normality | ✅ yes | ✅ **yes** | ✅ yes | ⚠️ only asymptotically |
| Near-multicollinearity | ✅ yes | ✅ **yes** | ✅ valid but **large** | ✅ valid but low power |

**Read the columns.** Only A1 (and A5, catastrophically) destroys unbiasedness. A3 and A4 cost efficiency and valid inference. A6 costs only exactness.

---

## 5. Model answer: "Describe the Gauss–Markov theorem and its assumptions" [4 pts]

> **Statement.** Consider the linear model $\boldsymbol{y}=\boldsymbol{X}\boldsymbol\beta+\boldsymbol\varepsilon$. If
> 1. the model is correctly specified and linear in the parameters, i.e. $E(\boldsymbol{y}\mid\boldsymbol{X})=\boldsymbol{X}\boldsymbol\beta$;
> 2. the errors have zero mean, $E(\varepsilon_i)=0$ for all $i$;
> 3. the errors are **homoscedastic**, $\text{Var}(\varepsilon_i)=\sigma^2$ for all $i$;
> 4. the errors are **uncorrelated**, $\text{Cov}(\varepsilon_i,\varepsilon_j)=0$ for $i\neq j$ — with (3) and (4) written jointly as $\text{Cov}(\boldsymbol\varepsilon)=\sigma^2\boldsymbol{I}_n$;
> 5. $\boldsymbol{X}$ is non-stochastic with full column rank $\text{rank}(\boldsymbol{X})=p$;
>
> then the OLS estimator $\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$ is the **Best Linear Unbiased Estimator (BLUE)** of $\boldsymbol\beta$.
>
> **Meaning of BLUE.** Among all estimators that are (i) **linear** in $\boldsymbol{y}$ and (ii) **unbiased** for $\boldsymbol\beta$, the OLS estimator has the **smallest variance** — more precisely, for any other linear unbiased $\tilde{\boldsymbol\beta}$, the matrix $\text{Cov}(\tilde{\boldsymbol\beta})-\text{Cov}(\hat{\boldsymbol\beta})$ is positive semi-definite, so $\text{Var}(\boldsymbol{a}'\hat{\boldsymbol\beta})\leq\text{Var}(\boldsymbol{a}'\tilde{\boldsymbol\beta})$ for every $\boldsymbol{a}$.
>
> **Note.** Normality of the errors is **not** required for Gauss–Markov. It is needed only for the exactness of $t$- and $F$-tests and confidence intervals.

That last note is what turns a 3-mark answer into a 4-mark one.

---

## 6. Key takeaways

1. **Six assumptions:** linearity · $E(\varepsilon)=0$ · homoscedasticity · no autocorrelation · full rank · (normality).
2. $\text{Cov}(\boldsymbol\varepsilon)=\sigma^2\boldsymbol{I}$ packs A3 and A4 into one statement.
3. 🔑 **The tiers:** A1,A2,A5 ⟹ unbiased. +A3,A4 ⟹ BLUE. +A6 ⟹ exact inference and OLS = ML.
4. **Heteroscedasticity and autocorrelation:** unbiased ✅, efficient ❌, standard errors ❌. Learn this sentence — it answers a whole family of questions.
5. **Only misspecification (A1) biases $\hat{\boldsymbol\beta}$.**
6. **Perfect vs near multicollinearity:** perfect breaks A5 and identification; near inflates variance but leaves everything unbiased and BLUE. VIF ≈ 1 means **no** problem.
7. **Normality buys exactness only.** Without it, tests are asymptotically valid.
