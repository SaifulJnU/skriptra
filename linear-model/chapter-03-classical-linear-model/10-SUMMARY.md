# Ch 3 — SUMMARY (the whole exam on two pages)

## The model

$$\boldsymbol{y}=\boldsymbol{X}\boldsymbol\beta+\boldsymbol\varepsilon,\qquad \boldsymbol{X}\in\mathbb{R}^{n\times p},\ p=k+1$$

$\boldsymbol{X}$: one row per observation, one column per parameter, **first column all ones**.

$$\hat{\boldsymbol{y}}=\boldsymbol{H}\boldsymbol{y},\quad \boldsymbol{H}=\boldsymbol{X}(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}',\quad \text{tr}(\boldsymbol{H})=p,\quad h_{ii}=\boldsymbol{x}_i'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_i$$

**With an intercept (automatic, not diagnostic):** $\sum\hat\varepsilon_i=0$, $\sum x_{ij}\hat\varepsilon_i=0$, $\bar{\hat y}=\bar y$.

$$\text{SST}=\text{SSE}+\text{explained SS},\qquad R^2=1-\frac{\text{SSE}}{\text{SST}}$$

> 🔴 **$R^2$ never decreases and SSE never increases when a covariate is added.** Four past-paper T/F statements test this one fact.

---

## 3.1.2 — Assumptions

| # | Assumption | Buys you |
|---|---|---|
| A1 | linearity / correct specification | **unbiased** |
| A2 | $E(\varepsilon_i)=0$ | unbiased |
| A3 | homoscedasticity $\text{Var}(\varepsilon_i)=\sigma^2$ | efficiency |
| A4 | no autocorrelation $\text{Cov}(\varepsilon_i,\varepsilon_j)=0$ | efficiency |
| A5 | $\text{rank}(\boldsymbol{X})=p$ | **existence + uniqueness** |
| A6 | normality (extra) | **exact** tests |

A3 + A4 $\;\Longleftrightarrow\;\text{Cov}(\boldsymbol\varepsilon)=\sigma^2\boldsymbol{I}_n$

### 🔑 The tiers

```
A1,A2,A5 ──► UNBIASED   +A3,A4 ──► BLUE   +A6 ──► EXACT tests, OLS=ML
```

### 🔴 The consequence grid

| Violated | Unbiased? | BLUE? | se valid? | Tests valid? |
|---|---|---|---|---|
| A1 linearity | ❌ | ❌ | ❌ | ❌ |
| A3 / A4 | ✅ | ❌ | ❌ | ❌ |
| A6 normality | ✅ | ✅ | ✅ | ⚠️ asymptotic only |
| Near-collinearity | ✅ | ✅ | ✅ but **inflated** | ✅ low power |

**Only A1 biases $\hat{\boldsymbol\beta}$.**

### Template answer for A3/A4 violations

> *OLS remains **unbiased** and consistent, since unbiasedness needs only correct specification, zero-mean errors and full rank. But it is no longer **efficient** — Gauss–Markov no longer applies, so OLS is not BLUE and WLS/GLS would have smaller variance. The usual standard errors are also biased, so t-tests, F-tests and CIs are invalid.*

**Perfect vs near multicollinearity:** perfect breaks A5 (not identified). Near leaves everything **unbiased and BLUE**, only inflating variance. **VIF ≈ 1 means NO problem**; concern above 5–10.

---

## 3.1.3 — Building $\boldsymbol{X}$

- **$c$ levels ⟹ $c-1$ dummies.** All $c$ + intercept ⟹ singular $\boldsymbol{X}'\boldsymbol{X}$ (the *dummy variable trap*). Reference = the level **missing** from R output. Compare two non-reference levels by **subtracting**.
- **Polynomial:** effect $=\beta_1+2\beta_2x$, turning point $-\hat\beta_1/(2\hat\beta_2)$. **Centre** ($\text{age}-48$) for interpretability + less collinearity.
- **Interaction:** dummy shifts, interaction tilts ⟹ **non-parallel lines**. $\partial y/\partial x=\beta_1+\beta_3D$.
- **⚠️ Variable in 2+ terms ⟹ differentiate. Never quote one coefficient.**
- **Restricted model:** substitute → collect → move parameter-free terms left as an offset.
  $H_0:\beta_1=\beta_2+1$ ⟹ regress $(y-x_1)$ on $(x_1+x_2)$, $r=1$.

---

## 3.2 — Estimation

### ⭐ The derivation (2 marks, practise 10×)

$$S(\boldsymbol\beta)=(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)=\boldsymbol{y}'\boldsymbol{y}-2\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{y}+\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{X}\boldsymbol\beta$$
$$\frac{\partial S}{\partial\boldsymbol\beta}=-2\boldsymbol{X}'\boldsymbol{y}+2\boldsymbol{X}'\boldsymbol{X}\boldsymbol\beta\overset{!}{=}\boldsymbol{0}\ \Longrightarrow\ \boldsymbol{X}'\boldsymbol{X}\hat{\boldsymbol\beta}=\boldsymbol{X}'\boldsymbol{y}$$
$$\boxed{\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}},\qquad \frac{\partial^2S}{\partial\boldsymbol\beta\partial\boldsymbol\beta'}=2\boldsymbol{X}'\boldsymbol{X}>0$$

*(Middle terms combine because $\boldsymbol{y}'\boldsymbol{X}\boldsymbol\beta$ is a scalar = its own transpose.)*

**Simple regression:** $\hat\beta_1=\widehat{\text{Cov}}/\widehat{\text{Var}}$, $\;\hat\beta_0=\bar y-\hat\beta_1\bar x$ ← recovers missing intercepts.

### Properties

$$\hat{\boldsymbol\beta}=\boldsymbol\beta+(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol\varepsilon\quad\text{(workhorse)}$$
$$E(\hat{\boldsymbol\beta})=\boldsymbol\beta,\qquad \boxed{\text{Cov}(\hat{\boldsymbol\beta})=\sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}},\qquad \widehat{\text{se}}(\hat\beta_j)=\hat\sigma\sqrt{[(\boldsymbol{X}'\boldsymbol{X})^{-1}]_{jj}}$$

> ⚠️ **Diagonal only — and label the rows $\beta_0,\beta_1,\dots$ first.** $\beta_1$ is the **second** diagonal element.
> ⚠️ Coefficients are generally **correlated**: "$\hat\beta_0$ and $\hat\beta_1$ are always uncorrelated" is **FALSE** (true only if $\bar x=0$).

### Error variance 🔴

$$\hat\sigma^2=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n-p}\ \text{(unbiased)}\qquad \hat\sigma^2_{ML}=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n}$$

| Use | Denominator |
|---|---|
| se, t, F, CI, PI, standardised residuals, R's "residual standard error", **REML** | **$n-p$** |
| **AIC, BIC** | **$n$** |

### Gauss–Markov [4 marks]

> Under **A1–A5**, $\hat{\boldsymbol\beta}$ is **BLUE**: among all estimators **linear** in $\boldsymbol{y}$ and **unbiased**, it has minimum variance.

🔴 **Normality is NOT required.** Say so — that's the 4th mark.
🔴 Drop "linear" or "unbiased" and the claim becomes **false** (ridge is biased with lower variance).

**Under A6:** $\hat{\boldsymbol\beta}\sim N(\boldsymbol\beta,\sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1})$ exactly, and $\hat{\boldsymbol\beta}_{ML}=\hat{\boldsymbol\beta}_{LS}$ (but $\hat\sigma^2_{ML}\neq\hat\sigma^2_{LS}$).

---

## 3.3 — Testing

### t-test

$$t=\frac{\hat\beta_j-c}{\widehat{\text{se}}(\hat\beta_j)}\sim t_{n-p};\qquad \text{reject if }|t|>t_{n-p}(1-\tfrac\alpha2)$$

**R-output identity:** $\;t=\hat\beta/\widehat{\text{se}}\;\Longleftrightarrow\;\hat\beta=t\times\widehat{\text{se}}\;\Longleftrightarrow\;\widehat{\text{se}}=\hat\beta/t$

### $\boldsymbol{C}\boldsymbol\beta=\boldsymbol{d}$

One **row per restriction**, one **column per parameter (including $\beta_0$)**. $\boldsymbol{C}$ is $r\times p$.

> 🔴 **$r$ = number of EQUATIONS, not number of betas mentioned.**
> $H_0:\beta_1=-\beta_2+\beta_3$ has three betas but **$r=1$**.

### F-test

$$\boxed{F=\frac{(\text{SSE}_{H_0}-\text{SSE})/r}{\text{SSE}/(n-p)}\sim F_{r,\,n-p}}\qquad F=\frac{(R^2-R^2_{H_0})/r}{(1-R^2)/(n-p)}$$

Overall test: $\;F=\dfrac{R^2/k}{(1-R^2)/(n-p)}\sim F_{k,\,n-p}$

**Checks:** $F\geq0$ always; if $r=1$ then $\sqrt{F}=|t|$.
**Conclusion language:** "at least one restriction fails," never "all." And "**fail to reject**," never "accept."

### Intervals

$$\text{CI for }\beta_j:\quad \hat\beta_j\pm t_{n-p}(1-\tfrac\alpha2)\,\widehat{\text{se}}(\hat\beta_j)$$

$$\text{CI for the mean at }\boldsymbol{x}_0:\quad \boldsymbol{x}_0'\hat{\boldsymbol\beta}\pm t\,\hat\sigma\sqrt{\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0}$$

$$\text{PREDICTION interval:}\quad \boldsymbol{x}_0'\hat{\boldsymbol\beta}\pm t\,\hat\sigma\sqrt{\mathbf{1}+\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0}$$

> 🔑 **The "$1+$" is the new observation's own error $\varepsilon_0$.** The prediction interval is always wider and **never shrinks to zero**, however large $n$ is. *Individuals are not averages.*

**CI–test duality:** $c$ inside the CI ⟺ don't reject $H_0:\beta_j=c$. **Zero outside ⟹ significant.**

### 🔴 Quantile rule

| | $\alpha=0.05$ |
|---|---|
| two-sided t-test, 95% CI | **0.975** |
| 99% CI | **0.995** |
| **F-test** | **0.95** ← one-sided |

> **t and CI are two-sided ⟹ $1-\alpha/2$. F is one-sided ⟹ $1-\alpha$.**

---

## 3.4 — Model choice & diagnostics

$$E[(y_0-\hat f)^2]=\underbrace{\sigma^2}_{\text{irreducible}}+\text{Bias}^2+\text{Variance}$$

More complexity ⟹ less bias, more variance. Total error is **U-shaped**. Bias enters **squared**, variance **linearly** — so a little bias can pay, which is why BLUE isn't the last word.

$$\bar R^2=1-\frac{n-1}{n-p}(1-R^2)\qquad\text{(can decrease; CAN be negative; penalty too weak)}$$

$$\boxed{\text{AIC}=n\log(\hat\sigma^2)+2(|M|+1)}\qquad \boxed{\text{BIC}=n\log(\hat\sigma^2)+\log(n)(|M|+1)}$$

$$\hat\sigma^2=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{\mathbf{n}}\ \textbf{(ML)},\qquad \log=\ln,\qquad \text{smaller is better}$$

> 🔴 **Three AIC/BIC traps:** divide by $n$ (not $n-p$) · natural log · don't forget the **$+1$** for $\sigma^2$.
> 🔴 **BIC penalises MORE than AIC** for $n>8$. *B for Bigger.* BIC ⟹ smaller models.
> 💡 $n\log(\hat\sigma^2)$ is identical in both — compute once.
> AIC ≈ prediction; BIC ≈ true model. Only comparable across models on the **same data**.

### Diagnostics

$$r_i=\frac{\hat\varepsilon_i}{\hat\sigma\sqrt{1-h_{ii}}}\qquad\text{because }\text{Var}(\hat\varepsilon_i)=\sigma^2(1-h_{ii})\text{ is NOT constant}$$

| Plot | Detects |
|---|---|
| Residuals vs fitted | **non-linearity** (curve), **heteroscedasticity** (funnel) |
| **QQ plot** | non-normality — points follow the **45° diagonal**, not a horizontal line |
| Scale–location | heteroscedasticity specifically |
| Residuals vs leverage | influential points (Cook's D) |

**Leverage ≠ outlier ≠ influence:** $h_{ii}$ (unusual $x$) · $|r_i|$ (unusual $y$) · $D_i=\frac{r_i^2}{p}\cdot\frac{h_{ii}}{1-h_{ii}}$ (changes the fit).
High leverage alone is harmless — it only hurts when paired with a large residual.

---

# 🎯 THE NIGHT-BEFORE CARD

- [ ] **Residual df = $n$ − (number of $\beta$'s incl. intercept) = $n-k-1$**
- [ ] **$r$ = number of EQUATIONS**
- [ ] **t & CI ⟹ $1-\alpha/2$; F ⟹ $1-\alpha$**
- [ ] **AIC/BIC: divide by $n$, natural log, $+1$**
- [ ] **BIC penalises more** (B for Bigger)
- [ ] **$c$ levels ⟹ $c-1$ dummies**
- [ ] **Prediction interval has the "$1+$"**
- [ ] **Heteroscedasticity ⟹ unbiased but inefficient**
- [ ] **Normality not needed for Gauss–Markov**
- [ ] **$R^2$ ↑ and SSE ↓ monotonically in model size**
- [ ] **VIF ≈ 1 = no problem**
- [ ] **QQ plot = 45° diagonal**
- [ ] **se uses the DIAGONAL; $\beta_1$ is the 2nd element**
- [ ] "fail to reject" · "associated with" · "holding others fixed"
- [ ] **Formula before numbers.** Round to 3 decimals.

**60 marks in 60 minutes = 1 minute per mark.** Speed comes from recognition, not cleverness.
