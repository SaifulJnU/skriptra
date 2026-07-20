# Ch 3 — FORMULA DECISION GUIDE

> **The question this file answers:** *"I'm looking at an exam question. Which formula do I reach for, and how did I know?"*
>
> Every entry has the same six fields:
> **① Formula ② USE WHEN ③ DON'T USE WHEN ④ WHY ⑤ 🔍 TRIGGER PHRASES in the question ⑥ ⚠️ Errors**
>
> Read the 🔍 lines the night before the exam. They are the fastest part of this folder.

---

# ⚡ THE 30-SECOND TRIAGE

Read the question. Find the phrase. Jump to the section.

| The question says…                                            | Go to                                                      |
| -------------------------------------------------------------- | ---------------------------------------------------------- |
| "estimate the coefficients" / "derive" / "explain OLS"         | **F1**                                               |
| "std. error", "t value", or`[[A]]` in an R table             | **F2, F3**                                           |
| "residual standard error" / "$\hat\sigma$"                   | **F4**                                               |
| "AIC" / "BIC"                                                  | **F5** ⚠️ *different $\hat\sigma^2$!*          |
| "confidence interval for$\beta_j$"                           | **F6**                                               |
| "test$H_0:\beta_j=$ something"                               | **F7**                                               |
| "joint hypothesis" / two or more restrictions / "$C\beta=d$" | **F8, F9**                                           |
| "predict the wage of a 50-year-old…"                          | **F10** or **F11** — *read carefully which* |
| "$R^2$" / "corrected coefficient of determination"           | **F12, F13**                                         |
| "residual" / "standardised residual"                           | **F14**                                              |
| "unbiased?" "efficient?" "BLUE?"                               | **F15** — no formula, a decision tree               |

---

# F1 — OLS estimator

### ① Formula

$$
\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}
$$

$$
\text{from minimising } S(\boldsymbol\beta)=(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)
$$

### ② USE WHEN

- Asked to **derive** or **explain** OLS (2 marks, appears most years)
- Computing $\hat{\boldsymbol\beta}$ by hand from a small dataset
- Any question mentioning "method of least squares"

### ③ DON'T USE WHEN

- ❌ $\boldsymbol{X}$ is **not full rank** — $(\boldsymbol{X}'\boldsymbol{X})^{-1}$ doesn't exist. Say "not identified," don't compute.
- ❌ The response is **binary** — use logit and maximum likelihood.
- ❌ You're asked for **weighted** least squares (heteroscedastic errors) — that's $(\boldsymbol{X}'\boldsymbol{W}\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{W}\boldsymbol{y}$, Chapter 4.

### ④ WHY

Squared error is differentiable, so setting $\partial S/\partial\boldsymbol\beta=0$ gives the linear **normal equations** $\boldsymbol{X}'\boldsymbol{X}\hat{\boldsymbol\beta}=\boldsymbol{X}'\boldsymbol{y}$, solvable in closed form. Geometrically it's the orthogonal projection of $\boldsymbol{y}$ onto the column space of $\boldsymbol{X}$.

### ⑤ 🔍 TRIGGERS

> *"Explain the method of ordinary least squares"* · *"show the steps necessary to obtain $\hat\beta_0,\dots,\hat\beta_k$"* · *"derive the solution"* · *"it is not necessary to explicitly calculate them"* · *"estimate the coefficients"*

**The tell:** if the question says *"the mathematical approach suffices"* or *"it is not necessary to calculate,"* they want the **four-line derivation**, not numbers. Write: expand → differentiate → normal equations → invert → second-order check.

### ⑥ ⚠️ ERRORS

- Forgetting to say **"minimise the residual sum of squares"** — that sentence alone is 1 of the 2 marks
- Not explaining why the middle terms combine into $-2\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{y}$ (scalar = its own transpose)
- Omitting the second-derivative check
- Forgetting the intercept column in $\boldsymbol{X}$

---

# F2 — Covariance of $\hat{\boldsymbol\beta}$ and standard errors

### ① Formula

$$
\text{Cov}(\hat{\boldsymbol\beta})=\sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1} \qquad\Longrightarrow\qquad \widehat{\text{se}}(\hat\beta_j)=\hat\sigma\sqrt{\left[(\boldsymbol{X}'\boldsymbol{X})^{-1}\right]_{jj}}
$$

### ② USE WHEN

- You're handed the matrix $(\boldsymbol{X}'\boldsymbol{X})^{-1}$ and asked for standard errors
- You need $\widehat{\text{se}}$ before a t-test or CI and the R table doesn't give it

### ③ DON'T USE WHEN

- ❌ **Heteroscedasticity or autocorrelation is present** — this formula is then **wrong**. Say so; that's often the actual answer.
- ❌ You already have `Std. Error` in the R output — just read it off.

### ④ WHY

From $\hat{\boldsymbol\beta}=\boldsymbol\beta+(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol\varepsilon$ and $\text{Cov}(\boldsymbol{Az})=\boldsymbol{A}\text{Cov}(\boldsymbol{z})\boldsymbol{A}'$ with $\text{Cov}(\boldsymbol\varepsilon)=\sigma^2\boldsymbol{I}$.

### ⑤ 🔍 TRIGGERS

> *"Calculate the estimated standard errors of the regression coefficients"* · *"Use $(\boldsymbol{X}'\boldsymbol{X})^{-1}=\dots$"* · *"Hint: you do not need the whole matrix"*

**That hint is the giveaway** — it means "you only need the **diagonal**."

### ⑥ ⚠️ ERRORS

- 🔴 **Off-by-one on the index.** The matrix is printed as rows 1…p but indexes $\beta_0,\dots,\beta_k$. **$\beta_1$ is the SECOND diagonal element.** This is the single most common slip in Sheet 3(e).
- Forgetting the **square root**
- Using $\hat\sigma^2$ instead of $\hat\sigma$ (i.e. forgetting to sqrt the variance first)
- Using an **off**-diagonal element

---

# F3 — The R-output identity

### ① Formula

$$
t=\frac{\hat\beta_j}{\widehat{\text{se}}(\hat\beta_j)} \qquad\Longleftrightarrow\qquad \hat\beta_j=t\times\widehat{\text{se}} \qquad\Longleftrightarrow\qquad \widehat{\text{se}}=\frac{\hat\beta_j}{t}
$$

### ② USE WHEN

**Any** "fill in the missing value in this R table" question. Two of the three are always given.

### ③ DON'T USE WHEN

- ❌ The missing value is the **residual standard error** — that's F4
- ❌ The missing value is the **intercept** and you're given $\bar x,\bar y$ — use $\hat\beta_0=\bar y-\hat\beta_1\bar x$ (faster, F16)

### ④ WHY

The `t value` column is *defined* as estimate/std.error.

### ⑤ 🔍 TRIGGERS

> *"Reproduce the values of [[A]], [[B]], [[C]]"* · *"there are missing parts in the R output"* · any table with `A`, `B`, `C` placeholders

### ⑥ ⚠️ ERRORS

- **Losing the minus sign.** If $t<0$ then $\hat\beta<0$. Check signs against the other columns.
- Rounding too early — the paper says round to **3 decimals** at the end

---

# F4 — Error variance

### ① Formula

$$
\hat\sigma^2=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n-p}\quad\textbf{(unbiased)} \qquad\qquad \hat\sigma^2_{ML}=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n}\quad\textbf{(ML)}
$$

### ② USE WHICH — 🔴 the whole point of this entry

| Purpose                                                                              | Denominator       |
| ------------------------------------------------------------------------------------ | ----------------- |
| standard errors, t-tests, F-tests, CIs, prediction intervals, standardised residuals | **$n-p$** |
| "residual standard error" in R output                                                | **$n-p$** |
| **AIC and BIC**                                                                | **$n$**   |
| question says "REML" or "restricted maximum likelihood"                              | **$n-p$** |
| question says "maximum likelihood estimate of$\sigma^2$"                           | **$n$**   |

### ③ DON'T USE WHEN

- ❌ $n=p$ — undefined, and correctly so
- ❌ Mixing them up. This is the trap.

### ④ WHY

The normal equations impose $p$ exact linear constraints on the residuals, so only $n-p$ are free. Formally $E(\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon})=\sigma^2\text{tr}(\boldsymbol{I}-\boldsymbol{H})=\sigma^2(n-p)$. ML doesn't correct for this, so $\hat\sigma^2_{ML}$ is biased low.

### ⑤ 🔍 TRIGGERS

> *"Residual standard error: [[D]] on 501 degrees of freedom"* · *"Use $\hat\varepsilon'\hat\varepsilon=3819720$"* · *"Calculate the Restricted Maximum Likelihood estimation for $\sigma^2$"*

**When they hand you $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}$ as a bare number, they are about to ask for $\hat\sigma^2$, an F-test, or AIC/BIC.** Work out which before dividing.

### ⑥ ⚠️ ERRORS

- 🔴 Using $n-p$ in AIC/BIC
- Forgetting the **square root** for the *standard error* (R reports $\hat\sigma$, not $\hat\sigma^2$)
- Getting $p$ wrong — cross-check against the "on ___ degrees of freedom" line

---

# F5 — AIC and BIC

### ① Formula

$$
\text{AIC}=n\log(\hat\sigma^2_{ML})+2(|M|+1) \qquad \text{BIC}=n\log(\hat\sigma^2_{ML})+\log(n)(|M|+1)
$$

$$
\hat\sigma^2_{ML}=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n},\qquad \log=\ln,\qquad |M|=\text{number of regression parameters}
$$

### ② USE WHEN

- Explicitly asked to compute AIC/BIC
- **Comparing two models** — especially non-nested ones
- Asked "how would you decide whether to add these variables?" (name it as method 1)

### ③ DON'T USE WHEN

- ❌ Models fitted to **different data** or different $n$ — values aren't comparable
- ❌ Models with **different response transformations** ($y$ vs $\log y$) — not comparable without a Jacobian correction
- ❌ You want a *significance test* of nested models — use the **F-test** (F9) instead

### ④ WHY

$-2\times$log-likelihood measures misfit; the penalty prices complexity. AIC targets prediction; BIC's $\log n$ penalty makes it consistent for finding the true model.

### ⑤ 🔍 TRIGGERS

> *"Calculate the AIC and the BIC"* · *"Which model do you prefer? Justify"* · *"the book uses $\log(x)$ as the natural logarithm"* · *"how would you decide whether you should do so? State two methods"*

### ⑥ ⚠️ ERRORS

- 🔴 Dividing by $n-p$ instead of $n$
- 🔴 Using $\log_{10}$ instead of $\ln$
- 🔴 Forgetting the **$+1$** (σ² counts as a parameter)
- 🔴 Saying **"AIC penalises more heavily than BIC"** — it's the reverse. *B for Bigger.*

### 💡 SHORTCUT

$n\log(\hat\sigma^2)$ is **identical** in both. Compute it once, then add the two penalties.

---

# F6 — Confidence interval for a coefficient

### ① Formula

$$
\hat\beta_j\ \pm\ t_{n-p}\!\left(1-\tfrac{\alpha}{2}\right)\cdot\widehat{\text{se}}(\hat\beta_j)
$$

### ② USE WHEN

- "Calculate a 95%/99% confidence interval for $\beta_j$"
- Asked to test $H_0:\beta_j=c$ **and** they've given you a CI — check whether $c$ is inside

### ③ DON'T USE WHEN

- ❌ You want an interval for a **prediction** — that's F10/F11
- ❌ You want an interval for a **combination** like $\beta_5-\beta_2$ — you need the off-diagonal covariance too
- ❌ Heteroscedasticity is present — the CI is invalid

### ④ WHY

$\frac{\hat\beta_j-\beta_j}{\widehat{\text{se}}}\sim t_{n-p}$; invert the two-sided acceptance region.

### ⑤ 🔍 TRIGGERS

> *"Calculate a 99% confidence interval"* · *"centered around the estimated coefficient"* · *"Note: you can find the correct quantile on the last page"*

**That last phrase means: go to the table. Pick the right column.**

### ⑥ ⚠️ ERRORS

- 🔴 **Using the $1-\alpha$ column instead of $1-\alpha/2$.** 95% CI ⟹ **0.975**. 99% CI ⟹ **0.995**. This is the #1 error in this exam.
- Using the F quantile instead of the t quantile
- Saying "there is a 95% probability $\beta_j$ lies in the interval" — wrong interpretation, marks lost

---

# F7 — t-test for a single coefficient

### ① Formula

$$
t=\frac{\hat\beta_j-c}{\widehat{\text{se}}(\hat\beta_j)}\overset{H_0}{\sim}t_{n-p};\qquad \text{reject if } |t|>t_{n-p}(1-\tfrac\alpha2)
$$

For a significance test, $c=0$.

### ② USE WHEN

- **One** restriction on **one** coefficient
- "Is this variable significant?"
- "Would you reject $H_0:\beta_{\text{nox}}=-30$?"

### ③ DON'T USE WHEN

- ❌ The hypothesis involves **two or more** coefficients (e.g. $\beta_1=\beta_2$) — use the F-test with a $\boldsymbol{C}$ matrix (F8/F9)
- ❌ **Several** restrictions at once — F-test
- ❌ You want to know if a *group* of dummies matters — F-test on the whole group

### ④ WHY

$\hat\beta_j$ is normal; $\hat\sigma$ is estimated, giving a $t$ rather than a $z$.

### ⑤ 🔍 TRIGGERS

> *"Test of significance"* · *"significantly different from 0"* · *"Which effects are significantly different from 0 at $\alpha=0.05$?"* · *"How is the test used in this scenario called?"* (→ "a t-test / test of significance")

### ⑥ ⚠️ ERRORS

- 🔴 **Forgetting the $-c$** when $H_0$ isn't $\beta_j=0$
- Using $1-\alpha$ instead of $1-\alpha/2$
- Writing "we accept $H_0$" — write **"fail to reject"**
- Concluding "the effect is zero" — write "not significantly different from zero"

### 💡 SHORTCUT

For the significance test you can read the p-value straight from `Pr(>|t|)`. Reject at $\alpha$ ⟺ p-value $<\alpha$. Faster than any table lookup.

---

# F8 — Building $\boldsymbol{C}\boldsymbol\beta=\boldsymbol{d}$

### ① Method (not a formula)

1. Rearrange each restriction: **all $\beta$'s on the left, constants on the right**
2. One **row of $\boldsymbol{C}$ per restriction**; read coefficients in the order $\beta_0,\beta_1,\dots,\beta_k$
3. $\boldsymbol{C}$ is $r\times p$; $\boldsymbol{d}$ is $r\times1$; $r=$ number of **independent equations**

### ② USE WHEN

- Any hypothesis with two or more coefficients, or two or more restrictions
- Asked explicitly to "express $H_0$ as $\boldsymbol{C}\boldsymbol\beta=\boldsymbol{d}$"

### ③ DON'T USE WHEN

- ❌ Single simple $H_0:\beta_j=0$ — a t-test is faster (though $\boldsymbol{C}$ still works)

### ④ WHY

It's the general framework: every linear hypothesis, however phrased, is $r$ linear equations in $\boldsymbol\beta$.

### ⑤ 🔍 TRIGGERS

> *"joint null hypothesis"* · *"Express the null hypothesis as $C\beta=d$"* · *"Specify the restriction matrix $C$"* · *"What is the number of linearly independent restrictions $r$?"* · *"composite test of a subvector"*

### ⑥ ⚠️ ERRORS

- 🔴 **Forgetting the $\beta_0$ column.** $\boldsymbol{C}$ has **$p$** columns, not $k$.
- 🔴 **Miscounting $r$.** $r$ = number of **EQUATIONS**, never the number of betas mentioned.
  > $H_0:\beta_1=-\beta_2+\beta_3$ mentions **three** betas but is **ONE** equation ⟹ $r=1$. *(Exam Summer 2025 Ex 1(i) is exactly this trap.)*
  >
- Not rearranging first — $\beta_1=3\beta_4-0.1$ must become $\beta_1-3\beta_4=-0.1$ before you can read off the row

---

# F9 — The F-test

### ① Formula

$$
F=\frac{(\text{SSE}_{H_0}-\text{SSE})/r}{\text{SSE}/(n-p)}\overset{H_0}{\sim}F_{r,\,n-p}
$$

or, when $R^2$'s are given,

$$
F=\frac{(R^2-R^2_{H_0})/r}{(1-R^2)/(n-p)}; \qquad \text{overall test: } F=\frac{R^2/k}{(1-R^2)/(n-p)}
$$

### ② USE WHICH VERSION

| They give you                                                                                                                                         | Use                                               |
| ----------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------- |
| $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}$ **and** $\hat{\boldsymbol\varepsilon}'_{H_0}\hat{\boldsymbol\varepsilon}_{H_0}$ | the**SSE** version                          |
| only$R^2$, and $H_0$ is "all slopes = 0"                                                                                                          | the**$R^2$** version with $R^2_{H_0}=0$ |
| the R output's`F-statistic:` line                                                                                                                   | it's already computed — just read it             |

### ③ DON'T USE WHEN

- ❌ Models are **not nested** — use AIC/BIC instead
- ❌ You want to know *which* coefficient is non-zero — F only says "at least one"

### ④ WHY

Imposing $H_0$ can only worsen the fit. The F-statistic asks whether the extra residual SS per restriction is large relative to the noise level $\hat\sigma^2$. Under $H_0$ both are estimates of $\sigma^2$, so $F\approx1$.

### ⑤ 🔍 TRIGGERS

> *"Calculate the F-statistic"* · *"Let $\hat\varepsilon'_{H_0}\hat\varepsilon_{H_0}=32333.15$"* · *"decide whether you would reject the null hypothesis"* · *"How many degrees of freedom does the distribution have?"*

**Being handed a second, larger sum of squares is a certain signal that an F-test is coming.**

### ⑥ ⚠️ ERRORS

- 🔴 **Using the 0.975 quantile.** F is **one-sided**: $\alpha=0.05$ ⟹ **0.95** column.
- 🔴 Denominator df $=n-p$ from the **unrestricted** model, not $n-r$
- 🔴 Numerator df $=r$, not $p$
- Dividing by $\text{SSE}_{H_0}$ in the denominator (must be the unrestricted SSE)
- Getting $F<0$ — you swapped the two SSEs
- Concluding "all restrictions fail" — say "**at least one**"

### 💡 CHECKS

- $F\geq0$ always
- If $r=1$, then $\sqrt{F}=|t|$ from the equivalent t-test

---

# F10 vs F11 — 🔴 CONFIDENCE vs PREDICTION INTERVAL

**This pair is the most confusable in the course. The difference is one "+1".**

### F10 — CI for the **mean** response

$$
\boldsymbol{x}_0'\hat{\boldsymbol\beta}\pm t_{n-p}(1-\tfrac\alpha2)\cdot\hat\sigma\sqrt{\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0}
$$

### F11 — PREDICTION interval for an **individual**

$$
\boldsymbol{x}_0'\hat{\boldsymbol\beta}\pm t_{n-p}(1-\tfrac\alpha2)\cdot\hat\sigma\sqrt{\mathbf{1}+\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0}
$$

### ② HOW TO TELL WHICH

| Question says                                          | Use                        |
| ------------------------------------------------------ | -------------------------- |
| "the**average** wage of 50-year-olds"            | **F10** (mean)       |
| "the**expected** wage for this group"            | **F10**              |
| "**the** wage of a 50-year-old man" (one person) | **F11** (prediction) |
| "**predict** the wage of…"                      | **F11**              |
| the words**"prediction interval"**               | **F11**              |

**Default heuristic:** *one specific individual ⟹ prediction interval. A group average ⟹ confidence interval.*

### ④ WHY THE +1

A new individual carries **their own fresh error $\varepsilon_0$**, with variance $\sigma^2$. That's the "1". The mean has no such term — averages don't have individual randomness.

**Consequence:** as $n\to\infty$ the CI shrinks to **zero width**; the prediction interval shrinks to $\pm t\hat\sigma$ and **never** to zero.

### ⑤ 🔍 TRIGGERS

> *"Calculate the prediction interval for the wage of a 50 year old man with an advanced degree"* · *"Use $x_0'(X'X)^{-1}x_0=0.0035$"*

**Being handed the scalar $\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0$ means one of these two is coming.** Then just read whether it's an individual or a mean.

### ⑥ ⚠️ ERRORS

- 🔴 **Forgetting the +1** — turns a ±70 interval into a ±4 one
- Building $\boldsymbol{x}_0$ wrong — **remember the leading 1** for the intercept, and set the right dummies (a reference-category person has **all** dummies 0)
- Using the wrong quantile column

---

# F12 — $R^2$

### ① Formula

$$
R^2=1-\frac{\text{SSE}}{\text{SST}}=\frac{\text{explained SS}}{\text{SST}};\qquad R^2=r_{xy}^2 \text{ (simple regression only)}
$$

### ② USE WHEN

- Interpreting goodness of fit
- Converting between $R^2$ and correlation in **simple** regression
- Feeding the $R^2$ version of the F-test

### ③ DON'T USE WHEN

- ❌ **Comparing models of different size** — $R^2$ always favours the bigger one. Use $\bar R^2$, AIC or BIC.
- ❌ $R^2=r^2$ in **multiple** regression — there $R^2=\text{corr}(y,\hat y)^2$
- ❌ Model has no intercept — the decomposition breaks down

### ⑤ 🔍 TRIGGERS

> *"What is the correlation between wage and age?"* given $R^2$ → $r=\text{sign}(\hat\beta_1)\sqrt{R^2}$
> *"Interpret your results and provide one possible explanation"* → low $R^2$ ≠ no relationship; suspect **non-linearity** or **omitted variables**

### ⑥ ⚠️ ERRORS

- Forgetting the **sign** comes from $\hat\beta_1$
- Claiming $R^2$ can decrease when covariates are added — it **can't**

---

# F13 — Adjusted $R^2$

### ① Formula

$$
\bar R^2=1-\frac{n-1}{n-p}(1-R^2)
$$

### ② USE WHEN

- Explicitly asked for the "corrected coefficient of determination"
- Comparing models of **different size** (but prefer AIC/BIC)

### ③ DON'T USE WHEN

- ❌ As your *primary* selection criterion — the book explicitly advises against it: the penalty is too weak, admitting variables with $|t|>1$ (p ≈ 0.3)

### ⑤ 🔍 TRIGGERS

> *"Calculate the corrected coefficient of determination"* · *"Take the corrected coefficient of determination, the AIC and the BIC into account"*

### ⑥ ⚠️ ERRORS

- Using $n-1$ in the denominator (it's $n-p$)
- Claiming it can't be negative — **it can**

---

# F14 — Residuals

### ① Formula

$$
\hat\varepsilon_i=y_i-\hat y_i \qquad\qquad r_i=\frac{\hat\varepsilon_i}{\hat\sigma\sqrt{1-h_{ii}}}
$$

### ② USE WHICH

- **Raw** — when asked to "calculate and interpret the residual" (keeps the units of $y$)
- **Standardised** — when asked to judge whether a point is an **outlier** ($|r_i|>2$)

### ④ WHY STANDARDISE

$\text{Var}(\hat\varepsilon_i)=\sigma^2(1-h_{ii})$ is **not constant**. High-leverage points have artificially small residuals, hiding outliers. Dividing by $\hat\sigma\sqrt{1-h_{ii}}$ equalises them.

### ⑤ 🔍 TRIGGERS

> *"What is his residual? Interpret this residual with one sentence."* · *"What is his standardized residual? (Hint: you can use $h_{ii}=0.0016$)"*

**Being given $h_{ii}$ is the signal for the standardised version.**

### ⑥ ⚠️ ERRORS

- Sign convention: $\hat\varepsilon=y-\hat y$ (**observed minus fitted**). Positive = earns more than predicted.
- Forgetting the $\sqrt{1-h_{ii}}$
- Using $\hat\sigma^2$ where $\hat\sigma$ belongs

---

# F15 — "Unbiased? Efficient? Valid?" — the decision tree

**No formula. A flowchart. Worth full marks when a question describes a violated assumption.**

```
What has gone wrong?
│
├── Wrong functional form / omitted relevant variable (A1)
│      └─► BIASED. Everything downstream is wrong.
│
├── Heteroscedasticity (A3)  or  Autocorrelation (A4)
│      └─► UNBIASED ✅ and CONSISTENT ✅
│          NOT BLUE ❌ (inefficient — WLS/GLS is better)
│          Standard errors WRONG ❌ ⟹ t/F/CI INVALID ❌
│
├── Non-normality (A6)
│      └─► UNBIASED ✅  BLUE ✅  se's fine ✅
│          Tests only ASYMPTOTICALLY valid ⚠️
│
├── PERFECT multicollinearity (A5 fails)
│      └─► β̂ NOT IDENTIFIED — X'X singular, no unique solution
│
└── NEAR multicollinearity (A5 holds)
       └─► UNBIASED ✅  BLUE ✅  se's VALID but INFLATED ⚠️
           ⟹ wide CIs, insignificant t's, unstable signs
```

### ⑤ 🔍 TRIGGERS

> *"Which impact does this phenomenon have on the OLS estimate in terms of **bias and efficiency**?"* — the words *bias* and *efficiency* mean: **answer "unbiased but inefficient"** and then explain.
>
> *"the variation in revenue grows as the number of employees grows"* → heteroscedasticity
> *"the residuals are not independent"* → autocorrelation
> *"highly correlated explanatory variables"* → near-multicollinearity

### ⑥ ⚠️ ERRORS

- Saying heteroscedasticity **biases** $\hat{\boldsymbol\beta}$ — it doesn't
- Saying normality is needed for **Gauss–Markov** — it isn't
- Confusing near- and perfect multicollinearity

---

# F16 — Small helpers worth having ready

| Formula                                                              | Use when                               | Trigger                                          |
| -------------------------------------------------------------------- | -------------------------------------- | ------------------------------------------------ |
| $\hat\beta_0=\bar y-\hat\beta_1\bar x$                             | recover a missing intercept            | *"an average of 48.61 goals and 46.61 points"* |
| $\hat\beta_1=\widehat{\text{Cov}}/\widehat{\text{Var}}=r\,s_y/s_x$ | simple regression by hand              | raw data or summary stats given                  |
| turning point$=-\hat\beta_1/(2\hat\beta_2)$                        | quadratic model                        | *"at what age is wage maximised?"*             |
| effect$=\hat\beta_1+2\hat\beta_2x$                                 | quadratic model                        | *"the effect of age at age 30"*                |
| effect$=\hat\beta_1+\hat\beta_3D$                                  | interaction model                      | *"the slope for smokers"*                      |
| $\hat\beta_A-\hat\beta_B$                                          | compare two non-reference dummy levels | *"difference between level 3 and level 5"*     |
| $F=t^2$                                                            | cross-check a single-restriction F     | —                                               |
| $h_{ii}>2p/n$                                                      | high leverage                          | *"identify influential observations"*          |
| $D_i=\frac{r_i^2}{p}\cdot\frac{h_{ii}}{1-h_{ii}}$                  | Cook's distance                        | *"unduly influencing the fit"*                 |
| $\text{VIF}_j=1/(1-R_j^2)$                                         | multicollinearity                      | VIF mentioned                                    |
| $\exp(\hat\beta_j)$                                                | logit odds ratio                       | logit model                                      |

---

# 🚦 THE MASTER FLOWCHART

```
                    READ THE QUESTION
                            │
        ┌───────────────────┼───────────────────┐
        │                   │                   │
   "DERIVE" or          "CALCULATE"       "EXPLAIN / TRUE-FALSE"
   "EXPLAIN OLS"             │                   │
        │                    │                   │
       F1              what's missing?        F15 tree
                             │              + the notes files
        ┌────────────────────┼────────────────────┐
        │                    │                    │
   a table cell        an interval           a test
        │                    │                    │
    F3 / F16      ┌──────────┴──────────┐   ┌─────┴─────┐
                  │                     │   │           │
             for a βⱼ            for a NEW x₀  1 restriction?
                  │                     │   │           │
                 F6              individual? ─┤    YES → F7 (t)
                                  │      │   │    NO  → F8 + F9 (F)
                              YES │      │ NO
                                  F11    F10
```

**And the two rules that survive everything:**

> 🔑 **Residual df = $n$ − (number of $\beta$'s you estimated, including the intercept).**
>
> 🔑 **Two-sided (t, CI) ⟹ use $1-\alpha/2$. One-sided (F) ⟹ use $1-\alpha$.**
