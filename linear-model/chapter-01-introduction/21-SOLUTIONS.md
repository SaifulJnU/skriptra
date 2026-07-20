# Ch 1 — SOLUTIONS

---

## Part A — TRUE / FALSE

**A1. FALSE.** $\varepsilon_i = y_i - \boldsymbol{x}_i'\boldsymbol\beta$ uses the **true unknown** $\boldsymbol\beta$ and is unobservable. $\hat\varepsilon_i = y_i - \boldsymbol{x}_i'\hat{\boldsymbol\beta}$ uses the **estimate** and is computable. Assumptions are made about $\varepsilon$; diagnostics are performed on $\hat\varepsilon$.

**A2. FALSE.** Take logarithms: $\log(y) = \beta_0 + \beta_1x_1 + \dots + \beta_kx_k + \varepsilon$, which is an ordinary linear model with response $\log(y)$.
> *This is literally a past-paper statement (RCLM WS 22/23, Block II(i)). Expect it again.*

**A3. FALSE.** "Linear" refers to linearity in the **parameters**. $y = \beta_0 + \beta_1x + \beta_2x^2 + \varepsilon$ is linear in $(\beta_0,\beta_1,\beta_2)$ — just define $x_2 := x^2$ and it's an ordinary multiple regression.

**A4. FALSE.** Zero correlation means no **linear** association. $Y = X^2$ with $X$ symmetric about 0 has $r = 0$ but $X$ and $Y$ are perfectly dependent. (Converse *is* true: independence ⟹ zero correlation.)

**A5. TRUE.** $R^2 = r_{xy}^2$ holds in simple linear regression with an intercept. It does **not** generalise to multiple regression, where $R^2$ equals the squared correlation between $y$ and $\hat y$.

**A6. FALSE.** Dimension is $n \times (k+1) = n \times p$. The extra column is the vector of ones for the intercept. Forgetting the intercept column is one of the most common errors in this course.

**A7. TRUE.** Continuous → linear model; binary → logit/probit; count → Poisson; ordinal → ordinal models.

**A8. FALSE.** That is the interpretation for the **log-log** model. In the **log-linear** model $\log y \sim x$, a one-**unit** increase in $x$ is associated with an approximately $100\beta_1$ **percent** change in $y$.
> *WS 23/24 Block III(iv) is this exact trap.*

**A9. TRUE.** $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}$ is $(1\times n)(n \times 1) = 1\times 1$, a scalar, equal to $\sum_i \hat\varepsilon_i^2 = $ SSE.

**A10. FALSE.** Anscombe's quartet: four datasets with identical means, variances, correlations and fitted lines but completely different structures (linear, curved, outlier-driven, single-leverage-point). Always plot.

---

## Part B — Short answer

### B1 (2 pts)

$\varepsilon$ collects **everything affecting $y$ that is not captured by the systematic component** $\boldsymbol{x}'\boldsymbol\beta$.

Three distinct reasons it is needed:

1. **Omitted variables.** We can never include every relevant covariate. Wage depends on motivation, luck, negotiating skill, none of which are in the dataset.
2. **Measurement error.** $y$ is recorded imperfectly (rounding, self-reporting, instrument noise).
3. **Genuine randomness / individual variation.** Even two identical individuals would not have identical outcomes. The world is not deterministic at the level we observe it.

*(A fourth acceptable answer: **model misspecification** — the true $f$ may not be exactly linear.)*

### B2 (2 pts)

The conclusion is wrong: correlation does not imply causation. Both variables are driven by a third variable — **the size of the fire**. Big fires cause both more firefighters to be dispatched *and* more property damage. This third variable is a **confounder**.

**How regression addresses it:** include the confounder as a covariate. In a multiple regression of damage on (firefighters, fire size), $\beta_{\text{firefighters}}$ is interpreted *holding fire size fixed* — and would very likely turn negative, which is the sensible answer. This is exactly why multiple regression exists: to isolate a partial effect. (Caveat: this only works for confounders you can measure and include.)

### B3 (3 pts)

| Study | Response | Two covariates | Type of $y$ | Model class |
|---|---|---|---|---|
| (a) Used car price | selling price (€) | mileage, age of car | continuous (right-skewed → consider log) | **classical linear model** |
| (b) Loan default | default yes/no | loan amount, credit history | **binary** | **logit / probit** |
| (c) Insurance claims | number of claims/year | driver age, annual mileage | **count** (non-negative integer) | **Poisson regression** |

### B4 (2 pts)

Two consequences of right-skewed $y$:

1. **The normality assumption on $\varepsilon$ is likely violated** — the residuals will inherit the right skew, invalidating the exact $t$- and $F$-tests (though large-$n$ asymptotics help).
2. **Heteroscedasticity** — for skewed positive variables the variance typically grows with the mean, so $\text{Var}(\varepsilon_i)$ is not constant. OLS remains unbiased but is no longer BLUE, and the usual standard errors are wrong.

*(Also acceptable: a few very large values become high-leverage/influential points.)*

**Standard remedy:** model $\log(y)$ instead of $y$. This compresses the right tail toward symmetry and often stabilises the variance. Interpretation changes: $\beta_j$ becomes an approximate **percentage** effect ($\times 100$).

### B5 (1 pt)

$$\boldsymbol{y} = \boldsymbol{X}\boldsymbol\beta + \boldsymbol\varepsilon$$

| Object | Dimension |
|---|---|
| $\boldsymbol{y}$ | $n \times 1$ |
| $\boldsymbol{X}$ | $n \times p$, with $p = k+1$; first column is $\boldsymbol{1}$ |
| $\boldsymbol\beta$ | $p \times 1$ |
| $\boldsymbol\varepsilon$ | $n \times 1$ |

Check: $(n\times p)(p\times 1) = n\times 1$ ✓

---

## Part C — Calculation

### C1 (2 pts)

**(a)** In simple linear regression, $R^2 = r^2$, so $|r| = \sqrt{0.038} = 0.1949$. The sign follows the sign of $\hat\beta_1 = +0.71 > 0$, so

$$\boxed{r = +0.195}$$

**(b)** There is a **weak positive linear** association between age and wage: older men in this sample tend to earn slightly more, but age explains only $3.8\%$ of the variation in wage.

**(c)** Two good explanations:
- **The relationship is not linear.** Sheet 2 shows the scatter plot is clearly **quadratic** — wage rises through early career, peaks around age 45–50, then flattens or declines. Pearson correlation and a straight-line $R^2$ are blind to this curvature, so they understate the true association.
- **Omitted variables.** Education, job class, health and marital status matter far more than age. Sheet 5 shows a model with all covariates reaches $R^2 = 0.34$.

### C2 (3 pts)

**(a)**
$$\boldsymbol{X} = \begin{pmatrix}
1 & 1 & 0\\
1 & 2 & 0\\
1 & 1 & 1\\
1 & 3 & 1\\
1 & 2 & 1
\end{pmatrix}$$

Note the leading column of ones for the intercept.

**(b)**

| Object | Dimension |
|---|---|
| $\boldsymbol{X}$ | $5 \times 3$ |
| $\boldsymbol\beta$ | $3 \times 1$ |
| $\boldsymbol{X}'\boldsymbol{X}$ | $3 \times 3$ |
| $\boldsymbol{X}'\boldsymbol{y}$ | $3 \times 1$ |

**(c)** Three parameters estimated ($\beta_0,\beta_1,\beta_2$), so residual df $= n - p = 5 - 3 = \boxed{2}$.

### C3 (2 pts)

**(a)** $R^2 = 1 - \dfrac{\text{SSE}}{\text{SST}} = 1 - \dfrac{260}{1000} = \boxed{0.74}$

**(b)** Explained SS $= \text{SST} - \text{SSE} = 1000 - 260 = \boxed{740}$

Check: $740/1000 = 0.74 = R^2$ ✓

### C4 (2 pts)

$$E(\tilde\beta) = E(2\hat\beta_1 - \beta_1) = 2E(\hat\beta_1) - \beta_1 = 2\beta_1 - \beta_1 = \beta_1 \quad\checkmark$$

(using $E(aX+b) = aE(X)+b$ with $a=2$, $b = -\beta_1$, since $\beta_1$ is a constant)

$$\text{Var}(\tilde\beta) = \text{Var}(2\hat\beta_1 - \beta_1) = 2^2\,\text{Var}(\hat\beta_1) = 4\sigma^2 c$$

(using $\text{Var}(aX+b) = a^2\text{Var}(X)$ — the constant shift contributes nothing)

**Comment:** both estimators are unbiased, but $\tilde\beta$ has four times the variance. We prefer $\hat\beta_1$ because among unbiased estimators we want the **smallest variance** — that is precisely the criterion in the Gauss–Markov theorem, where the OLS estimator is shown to be **BLUE** (Best Linear Unbiased Estimator), "best" meaning minimum variance in the class of linear unbiased estimators.

> This tiny exercise is the whole idea behind Section 3.2.3. Unbiasedness alone is not enough — you can always find silly unbiased estimators. Variance is the tiebreaker.

---

## Part D

### D1 (4 pts)

**(a) (1 pt)** Response: **monthly net rent (€)**. Type: **continuous** (and positive, likely right-skewed). Appropriate class: the **classical linear model**, estimated by OLS.

**(b) (2 pts)** Three plots:

1. **Histogram (or boxplot) of net rent** — check for skewness and outliers; decide whether a log transformation is needed.
2. **Scatter plot of net rent against living area** — check whether the trend is linear or curved, and whether the spread is constant (fanning out ⟹ heteroscedasticity). Sensible to also plot against year of construction.
3. **Side-by-side boxplots of net rent by district (and by kitchen quality)** — check whether the categorical variables shift the level of rent, which tells you whether dummy variables for them will earn their keep.

**(c) (1 pt)** Fit the model with **$\log(\text{net rent})$** as the response. This reduces skewness and often stabilises the error variance.

Interpretation change: coefficients are no longer in euros. In $\log(\text{rent}) = \beta_0 + \beta_1\cdot\text{area} + \dots$, a one-m² increase in living area is associated with an approximately $100\beta_1$ **percent** change in rent, holding the other covariates fixed. Effects become **multiplicative** on the original scale rather than additive.

---

## How to mark yourself

| Score | Verdict |
|---|---|
| 25–30 | Chapter 1 is done. Move to Chapter 2 today. |
| 18–24 | Re-read `03-notes-1.3-notation.md` and redo Part C. |
| < 18 | Re-read all three notes files, then retry. Do not proceed — Chapter 3 will be unreadable. |

Total available: 10 (A) + 10 (B) + 9 (C) + 4 (D) = **33 points**.
