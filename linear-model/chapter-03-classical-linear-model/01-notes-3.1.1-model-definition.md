# 3.1.1 — Model Definition, Parameters, Estimation and Residuals

---

## 1. The model, built from scratch

The book constructs the classical linear model from two assumptions about $f$ and $\varepsilon$:

**Assumption 1 — the systematic component is a linear combination of covariates.**

$$f(x_1,\dots,x_k) = \beta_0+\beta_1x_1+\dots+\beta_kx_k$$

Collecting into $p = k+1$ dimensional vectors $\boldsymbol{x}=(1,x_1,\dots,x_k)'$ and $\boldsymbol\beta=(\beta_0,\dots,\beta_k)'$:

$$f(\boldsymbol{x}) = \boldsymbol{x}'\boldsymbol\beta$$

The leading 1 in $\boldsymbol{x}$ is what puts the intercept in.

**Assumption 2 — errors are additive.**

$$y = \boldsymbol{x}'\boldsymbol\beta + \varepsilon$$

For each observation $i=1,\dots,n$:

$$\boxed{\;y_i = \beta_0+\beta_1x_{i1}+\dots+\beta_kx_{ik}+\varepsilon_i = \boldsymbol{x}_i'\boldsymbol\beta+\varepsilon_i\;}$$

---

## 2. Matrix form — write this fluently

$$\boxed{\;\boldsymbol{y} = \boldsymbol{X}\boldsymbol\beta+\boldsymbol\varepsilon\;}$$

$$
\underbrace{\begin{pmatrix}y_1\\y_2\\\vdots\\y_n\end{pmatrix}}_{n\times1}
=
\underbrace{\begin{pmatrix}
1 & x_{11} & x_{12} & \cdots & x_{1k}\\
1 & x_{21} & x_{22} & \cdots & x_{2k}\\
\vdots & \vdots & \vdots & \ddots & \vdots\\
1 & x_{n1} & x_{n2} & \cdots & x_{nk}
\end{pmatrix}}_{n\times p,\;\; p=k+1}
\underbrace{\begin{pmatrix}\beta_0\\\beta_1\\\vdots\\\beta_k\end{pmatrix}}_{p\times1}
+
\underbrace{\begin{pmatrix}\varepsilon_1\\\varepsilon_2\\\vdots\\\varepsilon_n\end{pmatrix}}_{n\times1}
$$

**$\boldsymbol{X}$ is the design matrix: one row per observation, one column per parameter, first column all ones.**

> 💰 **This display is worth 1 point in the *Linear_model_exam_sheet* paper, Exercise 2(a).** Practise until you can produce it — with dimension labels — in 60 seconds.

---

## 3. Three quantities you must never confuse

| | Formula | Known? | Random? |
|---|---|---|---|
| **Parameter** $\beta_j$ | — | ❌ unknown | ❌ fixed constant |
| **Estimator** $\hat\beta_j$ | $(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$ | ✅ computable | ✅ random (depends on sample) |
| **Error** $\varepsilon_i$ | $y_i-\boldsymbol{x}_i'\boldsymbol\beta$ | ❌ unobservable | ✅ random |
| **Residual** $\hat\varepsilon_i$ | $y_i-\boldsymbol{x}_i'\hat{\boldsymbol\beta}$ | ✅ computable | ✅ random |

**The hat rule:** *hat = computed from your sample = random. No hat = the truth = unknown and fixed.*

---

## 4. Fitted values, residuals, and the hat matrix

$$\hat{\boldsymbol{y}} = \boldsymbol{X}\hat{\boldsymbol\beta} = \underbrace{\boldsymbol{X}(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'}_{=:\;\boldsymbol{H}}\boldsymbol{y} = \boldsymbol{H}\boldsymbol{y}$$

$\boldsymbol{H}$ is the **hat matrix** (it "puts the hat on $\boldsymbol{y}$"), also called the **projection matrix** or **prediction matrix**. It is $n\times n$.

$$\hat{\boldsymbol\varepsilon} = \boldsymbol{y}-\hat{\boldsymbol{y}} = (\boldsymbol{I}-\boldsymbol{H})\boldsymbol{y}$$

### Properties of $\boldsymbol{H}$ — each one gets used later

| Property | Statement | Where it matters |
|---|---|---|
| **Symmetric** | $\boldsymbol{H}'=\boldsymbol{H}$ | variance derivations |
| **Idempotent** | $\boldsymbol{H}\boldsymbol{H}=\boldsymbol{H}$ | projecting twice = projecting once |
| **Trace = rank = $p$** | $\text{tr}(\boldsymbol{H}) = \sum_i h_{ii} = p$ | **gives the $n-p$ in $\hat\sigma^2$** |
| $\boldsymbol{I}-\boldsymbol{H}$ also symmetric idempotent | $\text{tr}(\boldsymbol{I}-\boldsymbol{H})=n-p$ | ← this *is* the residual df |
| $\boldsymbol{H}\boldsymbol{X}=\boldsymbol{X}$ | projection leaves $\boldsymbol{X}$ alone | $\boldsymbol{X}'\hat{\boldsymbol\varepsilon}=\boldsymbol{0}$ |

**Geometric meaning:** $\boldsymbol{H}$ projects $\boldsymbol{y}$ orthogonally onto the column space of $\boldsymbol{X}$. $\hat{\boldsymbol{y}}$ is the point in that space closest to $\boldsymbol{y}$; $\hat{\boldsymbol\varepsilon}$ is the perpendicular from $\boldsymbol{y}$ down to it.

> **This picture explains where $n-p$ comes from, and it's worth understanding rather than memorising.** $\boldsymbol{y}$ lives in $n$ dimensions. The fitted values are pinned to a $p$-dimensional subspace. So the residual vector is free to move in only the remaining $n-p$ dimensions. **Degrees of freedom = dimensions left over after fitting.**

### Leverage

$$h_{ii} = \text{the } i\text{-th diagonal element of } \boldsymbol{H} = \boldsymbol{x}_i'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_i$$

$h_{ii}$ is the **leverage** of observation $i$ — how much influence it has on its own fitted value. Facts:

- $0 \leq h_{ii}\leq 1$ and $\sum_i h_{ii}=p$, so the **average leverage is $p/n$**
- Rule of thumb: $h_{ii} > 2p/n$ is high leverage
- High leverage = unusual **covariate** values (far from $\bar{\boldsymbol{x}}$), regardless of $y$
- Used in standardised residuals and Cook's distance — Section 3.4.4

> 🔴 **RCLM WS 22/23, Block III(ii):** *"The 'Hat Matrix' provides the leverages which help identify potential outliers."* → **TRUE** (per the answer key). Careful with the wording though — leverage identifies unusual **x**-values, which is a specific kind of outlier. A point can be an outlier in $y$ with low leverage.

---

## 5. The normal equations and their consequences

Minimising the residual sum of squares gives the **normal equations** (full derivation in `04-notes-3.2.1`):

$$\boldsymbol{X}'\boldsymbol{X}\hat{\boldsymbol\beta} = \boldsymbol{X}'\boldsymbol{y} \qquad\Longleftrightarrow\qquad \boldsymbol{X}'\hat{\boldsymbol\varepsilon} = \boldsymbol{0}$$

That right-hand form says: **the residual vector is orthogonal to every column of $\boldsymbol{X}$.**

### Consequences (all guaranteed when the model contains an intercept)

| Consequence | Reason | Examinable as |
|---|---|---|
| $\sum_i\hat\varepsilon_i = 0$ | orthogonality to the **ones column** | "residuals have mean zero" |
| $\sum_i x_{ij}\hat\varepsilon_i=0$ for all $j$ | orthogonality to column $j$ | residuals uncorrelated with covariates |
| $\bar{\hat y} = \bar y$ | follows from the first | 🔴 RCLM WS 22/23 I(ii) → **TRUE** |
| $\sum_i\hat y_i\hat\varepsilon_i = 0$ | $\hat{\boldsymbol y}$ is in the column space | residuals ⊥ fitted values |

> ⚠️ **Important nuance.** $\sum\hat\varepsilon_i = 0$ is **automatic**, not evidence of a good model. A terrible model still has residuals summing to zero. When *WS 23/24 Block II(iv)* says *"in a well-fitted model, the mean of the residuals should be close to zero"* and marks it **TRUE**, understand that it's true by construction — it tells you nothing diagnostic.
>
> **The residual plots in 3.4.4 look at the *pattern* of residuals, not their mean, precisely because the mean is rigged.**

### Contrast: errors vs residuals

| | $\varepsilon_i$ (true) | $\hat\varepsilon_i$ (residual) |
|---|---|---|
| Mean | $E(\varepsilon_i)=0$ **assumed** | $\sum\hat\varepsilon_i=0$ **guaranteed** |
| Variance | $\sigma^2$, **constant** (assumed) | $\sigma^2(1-h_{ii})$, **not constant** |
| Covariance | $\text{Cov}(\varepsilon_i,\varepsilon_j)=0$ (assumed) | $-\sigma^2h_{ij}\neq0$ — **correlated!** |
| Observable | No | Yes |

$$\text{Cov}(\hat{\boldsymbol\varepsilon}) = \sigma^2(\boldsymbol{I}-\boldsymbol{H})$$

> 🔴 **This is the source of a nasty exam statement.** *RCLM WS 22/23, Block III(i):* "Since the residuals are normally distributed, standardized residuals are also normally distributed."
>
> The whole reason we **need** standardised residuals is that raw residuals have unequal variances $\sigma^2(1-h_{ii})$ and are correlated — even when the true errors are perfectly iid normal. Standardising divides out the $\sqrt{1-h_{ii}}$:
> $$r_i = \frac{\hat\varepsilon_i}{\hat\sigma\sqrt{1-h_{ii}}}$$
> These have approximately equal variance, but they are still not exactly normal (because $\hat\sigma$ is estimated from the same data) and still not independent. **Treat the statement as FALSE** and explain why in one line.

---

## 6. Sums of squares and $R^2$

$$\underbrace{\sum_i(y_i-\bar y)^2}_{\text{SST (total)}} = \underbrace{\sum_i(\hat y_i - \bar y)^2}_{\text{explained SS}} + \underbrace{\sum_i(y_i-\hat y_i)^2}_{\text{SSE (residual)}}$$

**This decomposition requires the model to contain an intercept.** It is Pythagoras in $n$ dimensions: the orthogonality $\hat{\boldsymbol{y}}\perp\hat{\boldsymbol\varepsilon}$ kills the cross-term.

$$\boxed{\;R^2 = \frac{\text{explained SS}}{\text{SST}} = 1-\frac{\text{SSE}}{\text{SST}} = 1-\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{\sum(y_i-\bar y)^2}\;}$$

### Four facts about $R^2$ that get examined

1. $0\leq R^2\leq 1$ (with an intercept).
2. **$R^2$ never decreases when a covariate is added** — even a useless one. Adding a column can only expand the column space, so the projection can only get closer to $\boldsymbol{y}$.
   > 🔴 *Exam Summer 2025, Ex 1(c):* "adding dummies for the weekday a person was born on can be expected to **lower** $R^2$" → **FALSE.** $R^2$ can only rise (or stay equal). *Adjusted* $R^2$ would likely fall.
   > 🔴 *WS 23/24, Block III(i):* "RSS may increase as more variables are added" → **FALSE.** SSE can only decrease or stay the same.
   > 🔴 *Linear_model_exam_sheet, Block II(iii):* "The coefficient of determination can decrease as more variables are added" → **FALSE** (per the key).
3. In **simple** regression with an intercept, $R^2 = r_{xy}^2$. In multiple regression, $R^2 = \text{corr}(y,\hat y)^2$.
4. $R^2$ measures **in-sample linear fit**, nothing about prediction, correctness or causation.

**Corrected (adjusted) $R^2$** — the fix, covered fully in 3.4.2:

$$\bar R^2 = 1-\frac{n-1}{n-p}(1-R^2)$$

This one **can** decrease when you add a variable, and **can be negative**.
> 🔴 *Linear_model_exam_sheet, Block III(iv):* "Adjusted $R^2$ … can never be negative" → **FALSE.**

---

## 7. Key takeaways

1. $\boldsymbol{y}=\boldsymbol{X}\boldsymbol\beta+\boldsymbol\varepsilon$, $\boldsymbol{X}$ is $n\times p$ with $p=k+1$, first column ones.
2. $\hat{\boldsymbol{y}}=\boldsymbol{H}\boldsymbol{y}$ with $\boldsymbol{H}=\boldsymbol{X}(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'$: symmetric, idempotent, $\text{tr}(\boldsymbol{H})=p$.
3. **$\text{tr}(\boldsymbol{I}-\boldsymbol{H}) = n-p$ is where residual degrees of freedom come from.** Dimensions left after fitting.
4. Leverage $h_{ii}=\boldsymbol{x}_i'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_i$; average $p/n$; flags unusual **covariate** values.
5. Normal equations ⟹ $\boldsymbol{X}'\hat{\boldsymbol\varepsilon}=\boldsymbol{0}$ ⟹ $\sum\hat\varepsilon_i=0$, $\bar{\hat y}=\bar y$ — **automatic, not diagnostic**.
6. $\text{Cov}(\hat{\boldsymbol\varepsilon})=\sigma^2(\boldsymbol{I}-\boldsymbol{H})$: residuals are **heteroscedastic and correlated even under perfect assumptions**. Hence standardisation.
7. $\text{SST}=\text{explained SS}+\text{SSE}$; $R^2=1-\text{SSE}/\text{SST}$; **$R^2$ and SSE are monotone in model size** — this generates at least three T/F statements per paper.
