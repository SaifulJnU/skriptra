# Ch 1 — SUMMARY (one page)

## The one sentence

$$\boxed{\;y = \underbrace{f(x)}_{\text{systematic}} + \underbrace{\varepsilon}_{\text{random}}\;}$$

Regression separates signal from noise. Everything else is detail about the shape of $f$ and the assumptions on $\varepsilon$.

---

## Vocabulary

| Term | Symbol | Meaning |
|---|---|---|
| Response | $y$ | what we explain |
| Covariate | $x$ | what we explain it with |
| Systematic component | $f(x)$, $\boldsymbol{x}'\boldsymbol\beta$ | the structured part |
| Error | $\varepsilon$ | **true**, unobservable |
| Residual | $\hat\varepsilon = y - \hat y$ | **observed**, computable |
| Parameter | $\beta$ | true, unknown, fixed |
| Estimate | $\hat\beta$ | from data, random |

**Hat = from data = random. No hat = truth = unknown.**

---

## The model class is chosen by the type of $y$

| Type of $y$ | Model |
|---|---|
| continuous | classical linear model (Ch 3) |
| binary 0/1 | **logit / probit** (Ch 2.3) |
| count | Poisson (Ch 5) |
| ordered categories | ordinal (Ch 6) |

Covariate type affects *how* you code it (dummies, polynomials), never the model class.

---

## "Linear" = linear in $\beta$, not in $x$

✅ $y = \beta_0 + \beta_1 x + \beta_2 x^2 + \varepsilon$ (polynomial)
✅ $y = \beta_0 + \beta_1\log x + \varepsilon$
✅ $y = \beta_0 + \beta_1 x_1 + \beta_2 x_2 + \beta_3 x_1x_2 + \varepsilon$ (interaction)
✅ $y = \exp(\boldsymbol{x}'\boldsymbol\beta + \varepsilon)$ → take logs → linear
❌ $y = \beta_0 + x^{\beta_1} + \varepsilon$

---

## Look before you model

**Univariate:** histogram, boxplot, mean vs median.
Right-skewed $y$ (wage, rent, income) → consider $\log(y)$.

**Bivariate:** scatter plot. Ask four questions:
1. trend? 2. straight? 3. constant spread? 4. outliers?

**Correlation:**
$$r_{xy} = \frac{\widehat{\text{Cov}}(x,y)}{s_x s_y} \in [-1,1]$$

Three limits: only linear · not causal · Anscombe (identical stats, different pictures).

**$R^2 = r^2$ in simple linear regression.** Sign from $\hat\beta_1$.
*(Sheet 3 Ex 1: $R^2 = 0.038 \Rightarrow r = +0.195$.)*

---

## Log interpretation table (memorise)

| Model | 1-unit ↑ in $x$ means |
|---|---|
| $y \sim x$ | $\beta_1$ units in $y$ |
| $\log y \sim x$ | $100\beta_1$ **%** in $y$ |
| $y \sim \log x$ | (1% ↑ in $x$) → $\beta_1/100$ units in $y$ |
| $\log y \sim \log x$ | (1% ↑ in $x$) → $\beta_1$ **%** in $y$ — elasticity |

---

## Notation

$$\boldsymbol{y} = \boldsymbol{X}\boldsymbol\beta + \boldsymbol\varepsilon, \qquad \boldsymbol{X} \in \mathbb{R}^{n\times p},\ p = k+1$$

- $\boldsymbol{X}$: one row per observation, one column per parameter, **first column all ones**
- $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon} = \sum\hat\varepsilon_i^2 = \text{SSE}$
- $\text{SST} = \text{SSE} + \text{explained SS}$, and $R^2 = 1 - \text{SSE}/\text{SST}$
- $\text{Cov}(\boldsymbol{Az}) = \boldsymbol{A}\,\text{Cov}(\boldsymbol z)\,\boldsymbol{A}'$

**⚠️ The `p` trap:** book has $p = k+1$ (parameters); some exam papers use $p = k$ (covariates). **Always fall back on: residual df = $n$ − (number of betas including intercept) = $n-k-1$.**

---

## Three goals, three different jobs

| Goal | You care about | Drives |
|---|---|---|
| Description | sign, size, significance of $\beta_j$ | Ch 3.3 |
| Prediction | $\hat y$ on **new** data | Ch 3.4 |
| Causality | needs near-random assignment | beyond this course |

Say **"associated with"**, never "causes."
