# Ch 2 — SUMMARY (one page)

## The frame

$$E(y\mid\boldsymbol{x}) = f(\boldsymbol{x}) \qquad\Longleftrightarrow\qquad y = f(\boldsymbol{x})+\varepsilon,\; E(\varepsilon)=0$$

Regression models the **conditional mean**. Hence every interpretation says *"expected"* / *"on average."*

**Linear predictor** $\eta = \boldsymbol{x}'\boldsymbol\beta$ — the skeleton of every model here.
Linear model: $E(y)=\eta$. Logit: $P(y=1)=h(\eta)$ with $h$ squashing into $(0,1)$.

---

## 2.2.1 Simple linear regression

$$y_i = \beta_0+\beta_1x_i+\varepsilon_i$$

$$\hat\beta_1 = \frac{\widehat{\text{Cov}}(x,y)}{\widehat{\text{Var}}(x)} = r\frac{s_y}{s_x} \qquad \hat\beta_0 = \bar y - \hat\beta_1\bar x$$

- **Line passes through $(\bar x,\bar y)$** ← recovers missing intercepts from R output
- $R^2 = r^2$, sign of $r$ = sign of $\hat\beta_1$ (simple regression only)
- With intercept: $\sum\hat\varepsilon_i=0$, $\sum x_i\hat\varepsilon_i=0$, $\bar{\hat y}=\bar y$

**Slope sentence:** *"A one-[unit] increase in [x] is associated with an estimated [β̂₁] [unit] change in the expected [y], holding other covariates fixed."*

**Don't interpret $\hat\beta_0$** unless $x=0$ is meaningful and in range → **centre** the variable instead.

**Polynomial:** still linear in $\beta$. Effect $= \beta_1+2\beta_2x$ — **not** $\beta_1$. Peak at $-\hat\beta_1/(2\hat\beta_2)$.
*(Sheet 2: $5.29/0.10 = 52.9$ years.)*

---

## 2.2.2 Multiple regression + dummies ⭐

$$y_i = \boldsymbol{x}_i'\boldsymbol\beta+\varepsilon_i$$

$\hat\beta_j$ = **partial** effect, *holding all other covariates fixed*.
⚠️ Partial effect and marginal correlation can have **opposite signs**.

### 🔴 THE DUMMY RULE

> **$c$ levels ⟹ $c-1$ dummies. One level is the reference.**

All $c$ dummies + intercept ⟹ columns sum to the intercept ⟹ $\boldsymbol{X}$ not full rank ⟹ $\boldsymbol{X}'\boldsymbol{X}$ singular ⟹ **no unique OLS**. (The *dummy variable trap*.)

- Every dummy coefficient compares **to the reference**
- To compare two non-reference levels: **subtract their coefficients**
- The reference is the level **missing from the R output**

**Sheet 1:** 5 education levels → 4 dummies, reference `< HS Grad`.
Level 5 vs level 3 at same age: $64.99-24.17 = 40.82$.
40yr/L4 vs 20yr/L1: $0.56869(20)+39.767 = 51.14$.

### Interactions

$$y = \beta_0+\beta_1x+\beta_2D+\beta_3(x\cdot D)+\varepsilon$$

| $D=0$ | $D=1$ |
|---|---|
| intercept $\beta_0$, slope $\beta_1$ | intercept $\beta_0+\beta_2$, slope $\beta_1+\beta_3$ |

> **Dummy alone = parallel lines. Dummy + interaction = non-parallel lines.**

⚠️ **Variable in multiple terms ⟹ interpret coefficients jointly, never alone.**

---

## 2.3 Logit model ⭐ (guaranteed exam content)

Binary $y$ ⟹ $E(y\mid\boldsymbol{x}) = P(y=1) = \pi \in [0,1]$.

### Why linear fails — 4 reasons

1. 🔴 **Predictions escape $[0,1]$** (lead with this)
2. $\text{Var}(y)=\pi(1-\pi)$ ⟹ **heteroscedasticity by construction**
3. Errors take only **two values** ⟹ cannot be normal
4. Constant marginal effect is implausible near the boundaries

### Three forms

$$\pi = \frac{e^{\eta}}{1+e^{\eta}} \qquad \frac{\pi}{1-\pi}=e^{\eta} \qquad \log\frac{\pi}{1-\pi}=\eta = \boldsymbol{x}'\boldsymbol\beta$$

### 🔴 Interpretation

| Scale | Effect of +1 in $x_j$ |
|---|---|
| log-odds | $+\hat\beta_j$ (exact, constant) |
| **odds** | $\times\exp(\hat\beta_j)$ — the **odds ratio** |
| probability | $\hat\beta_j\pi(1-\pi)$ — **not constant** |

**NOT** "increases $P(y=1)$ by $\hat\beta_j$." That's the Exam 2025 trap, and it's FALSE.
Only the **sign** transfers reliably.

Probit = same with $\Phi$. Both fitted by **maximum likelihood**, not OLS.

---

## Model-building recipe

1. Identify $y$, check its type → picks the model class
2. For each covariate: continuous / binary / categorical ($c$ levels)?
3. Code: continuous as-is (or $x^2$, $\log x$) · binary → 1 dummy · categorical → $c-1$ dummies
4. Write equation with intercept, **define every dummy explicitly**, name the reference, add $\varepsilon_i$ with assumptions

$$p = 1 + (\text{continuous}) + \textstyle\sum(c-1) + (\text{interactions}), \qquad \text{df} = n-p$$
