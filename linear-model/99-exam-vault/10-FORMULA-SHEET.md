# 10 — FORMULA SHEET

> **Notation:** book (Fahrmeir) convention throughout. $k$ = covariates, $p = k+1$ = parameters including the intercept, $n$ = observations.
> In free-response answers write **$n-k-1$** — it is unambiguous under both conventions. See `chapter-01-introduction/03-notes-1.3-notation.md` §2.

---

## §0 — The reproduction order

**This is the point of the file.** Blank page, this order, no notes. Three times on Day 20.

1. The model and the hat matrix
2. OLS: the objective, the derivative, the solution
3. Moments of $\hat{\boldsymbol\beta}$, and $\widehat{\text{se}}$
4. The two $\hat\sigma^2$
5. t-test, $\boldsymbol{C}\boldsymbol\beta=\boldsymbol{d}$, F-test
6. The three intervals
7. $R^2$, $\bar R^2$, AIC, BIC
8. Standardised residuals, leverage, Cook's D
9. The logit model, three forms

Target: **under 15 minutes** by the third attempt.

---

## §1 — The model

$$\boldsymbol{y}=\boldsymbol{X}\boldsymbol\beta+\boldsymbol\varepsilon,\qquad \boldsymbol{y}\in\mathbb{R}^{n},\quad \boldsymbol{X}\in\mathbb{R}^{n\times p},\quad \boldsymbol\beta\in\mathbb{R}^{p},\quad p=k+1$$

$\boldsymbol{X}$: one **row per observation**, one **column per parameter**, first column all ones.

$$E(\boldsymbol\varepsilon)=\boldsymbol{0},\qquad \text{Cov}(\boldsymbol\varepsilon)=\sigma^2\boldsymbol{I}_n,\qquad \text{rank}(\boldsymbol{X})=p$$

### Hat matrix

$$\hat{\boldsymbol{y}}=\boldsymbol{H}\boldsymbol{y},\qquad \boldsymbol{H}=\boldsymbol{X}(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'$$

$$\boldsymbol{H}'=\boldsymbol{H},\qquad \boldsymbol{H}\boldsymbol{H}=\boldsymbol{H},\qquad \text{tr}(\boldsymbol{H})=p,\qquad h_{ii}=\boldsymbol{x}_i'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_i$$

$$\hat{\boldsymbol\varepsilon}=\boldsymbol{y}-\hat{\boldsymbol{y}}=(\boldsymbol{I}-\boldsymbol{H})\boldsymbol{y}$$

### Automatic with an intercept — construction, not diagnosis

$$\sum_i\hat\varepsilon_i=0,\qquad \sum_i x_{ij}\hat\varepsilon_i=0,\qquad \bar{\hat y}=\bar y,\qquad \text{line passes through }(\bar x,\bar y)$$

---

## §2 — OLS

### The derivation (2 marks, practise 10×)

$$S(\boldsymbol\beta)=(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)=\boldsymbol{y}'\boldsymbol{y}-2\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{y}+\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{X}\boldsymbol\beta$$

$$\frac{\partial S}{\partial\boldsymbol\beta}=-2\boldsymbol{X}'\boldsymbol{y}+2\boldsymbol{X}'\boldsymbol{X}\boldsymbol\beta\overset{!}{=}\boldsymbol{0}\quad\Longrightarrow\quad \underbrace{\boldsymbol{X}'\boldsymbol{X}\hat{\boldsymbol\beta}=\boldsymbol{X}'\boldsymbol{y}}_{\text{normal equations}}$$

$$\boxed{\;\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}\;}\qquad \frac{\partial^2 S}{\partial\boldsymbol\beta\,\partial\boldsymbol\beta'}=2\boldsymbol{X}'\boldsymbol{X}>0\ \Rightarrow\ \text{minimum}$$

*Say out loud while writing:* "minimise the residual sum of squares · differentiate · set to zero · normal equations · invert, which needs full rank · second derivative positive definite so it's a minimum."

### Simple regression — recovers missing R output

$$\hat\beta_1=\frac{\widehat{\text{Cov}}(x,y)}{\widehat{\text{Var}}(x)}=r\frac{s_y}{s_x},\qquad \hat\beta_0=\bar y-\hat\beta_1\bar x,\qquad R^2=r^2$$

---

## §3 — Moments and standard errors

$$\hat{\boldsymbol\beta}=\boldsymbol\beta+(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol\varepsilon\qquad\text{(the workhorse — everything below follows from it)}$$

$$E(\hat{\boldsymbol\beta})=\boldsymbol\beta,\qquad \boxed{\text{Cov}(\hat{\boldsymbol\beta})=\sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}},\qquad \widehat{\text{se}}(\hat\beta_j)=\hat\sigma\sqrt{\left[(\boldsymbol{X}'\boldsymbol{X})^{-1}\right]_{jj}}$$

> ⚠️ **Diagonal element only.** Label the rows $\beta_0,\beta_1,\dots$ before reading: $\beta_1$ is the **second** diagonal entry.
> ⚠️ Coefficients are generally **correlated**. In simple regression $\text{Cov}(\hat\beta_0,\hat\beta_1)=\dfrac{-\sigma^2\bar x}{\sum(x_i-\bar x)^2}$ — zero **only if $\bar x=0$**.

### Gauss–Markov [4 marks]

> Under **A1–A5**, $\hat{\boldsymbol\beta}$ is **BLUE**: among all estimators that are **linear in $\boldsymbol{y}$** and **unbiased**, it has minimum variance.

**Normality is not required** — say so explicitly, that is the fourth mark.

Under **A6** additionally: $\hat{\boldsymbol\beta}\sim N\!\left(\boldsymbol\beta,\ \sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}\right)$ exactly, and $\hat{\boldsymbol\beta}_{ML}=\hat{\boldsymbol\beta}_{LS}$ (but $\hat\sigma^2_{ML}\neq\hat\sigma^2_{LS}$).

### The assumption tiers

```
A1, A2, A5  ──►  UNBIASED
  + A3, A4  ──►  BLUE
      + A6  ──►  EXACT tests, OLS = ML
```

| # | Assumption | Buys you |
|---|---|---|
| A1 | linearity / correct specification | unbiasedness |
| A2 | $E(\varepsilon_i)=0$ | unbiasedness |
| A3 | homoscedasticity $\text{Var}(\varepsilon_i)=\sigma^2$ | efficiency |
| A4 | no autocorrelation $\text{Cov}(\varepsilon_i,\varepsilon_j)=0$ | efficiency |
| A5 | $\text{rank}(\boldsymbol{X})=p$ | existence + uniqueness |
| A6 | normality | exact tests |

**Only A1 biases $\hat{\boldsymbol\beta}$.**

---

## §4 — The two error variances 🔴

$$\hat\sigma^2=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n-p}\quad\textbf{(unbiased)}\qquad\qquad \hat\sigma^2_{ML}=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n}\quad\textbf{(ML)}$$

| Use | Denominator |
|---|---|
| se, t, F, CI, PI, standardised residuals, R's "residual standard error", REML | $n-p$ |
| **AIC, BIC** | $\boldsymbol{n}$ |

R reports the residual **standard error** $\hat\sigma$ — take the square root.

---

## §5 — Testing

### t-test

$$t=\frac{\hat\beta_j-c}{\widehat{\text{se}}(\hat\beta_j)}\ \sim\ t_{n-p}\quad\text{under }H_0:\beta_j=c;\qquad \text{reject if }|t|>t_{n-p}(1-\tfrac{\alpha}{2})$$

**R-output identity:** $t=\hat\beta/\widehat{\text{se}}\iff\hat\beta=t\cdot\widehat{\text{se}}\iff\widehat{\text{se}}=\hat\beta/t$ — this fills two-thirds of every missing-output question.

### General linear hypothesis

$$H_0:\boldsymbol{C}\boldsymbol\beta=\boldsymbol{d}\quad\text{vs}\quad H_1:\boldsymbol{C}\boldsymbol\beta\neq\boldsymbol{d},\qquad \boldsymbol{C}\in\mathbb{R}^{r\times p}$$

One **row per restriction**, one **column per parameter including $\beta_0$**.

> 🔴 **$r$ = the number of EQUATIONS**, not the number of betas mentioned.
> $H_0:\beta_1=-\beta_2+\beta_3$ mentions three betas but is **one** equation ⟹ $r=1$.
> Always rearrange to "everything on the left = constant on the right" **before** reading off the row.

### F-test

$$\boxed{\;F=\frac{(\text{SSE}_{H_0}-\text{SSE})/r}{\text{SSE}/(n-p)}\ \sim\ F_{r,\,n-p}\;}\qquad F=\frac{(R^2-R^2_{H_0})/r}{(1-R^2)/(n-p)}$$

Both SSE and df in the **denominator** come from the **unrestricted** model.

**Overall model test** ($H_0:\beta_1=\dots=\beta_k=0$, so $r=k$):

$$F=\frac{R^2/k}{(1-R^2)/(n-p)}\ \sim\ F_{k,\,n-p}$$

**Sanity checks:** $F\geq0$ always (negative ⟹ you swapped the SSEs). If $r=1$ then $\sqrt{F}=|t|$.

**Language:** reject ⟹ "**at least one** restriction fails," never "all." And "**fail to reject**," never "accept."

### 🔴 The quantile rule

| Situation | Quantile at $\alpha=0.05$ |
|---|---|
| two-sided t-test, 95% CI | $1-\alpha/2 = 0.975$ |
| 99% CI | $0.995$ |
| **F-test** | $1-\alpha = 0.95$ |

> **t and CI are two-sided ⟹ split $\alpha$. F rejects only in the upper tail ⟹ don't split.**

---

## §6 — The three intervals

$$\text{CI for }\beta_j:\qquad \hat\beta_j\ \pm\ t_{n-p}(1-\tfrac{\alpha}{2})\cdot\widehat{\text{se}}(\hat\beta_j)$$

$$\text{CI for the MEAN at }\boldsymbol{x}_0:\qquad \boldsymbol{x}_0'\hat{\boldsymbol\beta}\ \pm\ t_{n-p}(1-\tfrac{\alpha}{2})\cdot\hat\sigma\sqrt{\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0}$$

$$\text{PREDICTION interval for a NEW }y_0:\qquad \boldsymbol{x}_0'\hat{\boldsymbol\beta}\ \pm\ t_{n-p}(1-\tfrac{\alpha}{2})\cdot\hat\sigma\sqrt{\mathbf{1}+\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0}$$

> 🔑 The **$1+$** is the new observation's own error $\varepsilon_0$. The prediction interval is always wider and **never shrinks to zero** however large $n$ gets. *Individuals are not averages.*
>
> **Which one is wanted?** "*the* wage of *a* 50-year-old man" (singular, one person) ⟹ **prediction**. "the *average* wage of 50-year-old men" ⟹ **CI for the mean**.

**CI–test duality:** $c$ inside the CI $\iff$ do **not** reject $H_0:\beta_j=c$. Zero **outside** ⟹ significant.

---

## §7 — Model choice

$$\text{SST}=\text{SSE}+\text{SSR},\qquad R^2=1-\frac{\text{SSE}}{\text{SST}}$$

> 🔴 Adding a covariate: $R^2$ can only **rise or stay**; SSE can only **fall or stay**. Never the reverse.

$$\bar R^2=1-\frac{n-1}{n-p}\left(1-R^2\right)\qquad\text{(can decrease; CAN be negative; penalty is weak)}$$

$$\boxed{\;\text{AIC}=n\log(\hat\sigma^2_{ML})+2(|M|+1)\;}\qquad \boxed{\;\text{BIC}=n\log(\hat\sigma^2_{ML})+\log(n)\,(|M|+1)\;}$$

with $\hat\sigma^2_{ML}=\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}/\boldsymbol{n}$, $\log=\ln$, and **smaller is better**.

> 🔴 **Three traps in one formula:** divide by $n$ (not $n-p$) · natural log · don't drop the **$+1$** (that's $\sigma^2$, a parameter too).
> 🔴 $\log(n)>2$ once $n>7.4$, so **BIC penalises more** and picks **smaller** models. *B for Bigger penalty.*
> 💡 $n\log(\hat\sigma^2_{ML})$ is identical in both — compute it once, then add the two different penalties.
> AIC ≈ best prediction · BIC ≈ the true model. Comparable only across models fitted to the **same data**.

### Bias–variance

$$E\left[(y_0-\hat f)^2\right]=\underbrace{\sigma^2}_{\text{irreducible}}+\text{Bias}^2+\text{Variance}$$

More complexity ⟹ less bias, more variance; total error is **U-shaped**. Bias enters **squared**, variance **linearly** — a little bias can pay, which is why BLUE is not the last word.

---

## §8 — Diagnostics

$$\text{Var}(\hat{\boldsymbol\varepsilon})=\sigma^2(\boldsymbol{I}-\boldsymbol{H})\ \Longrightarrow\ \text{Var}(\hat\varepsilon_i)=\sigma^2(1-h_{ii})\ \textbf{not constant},\quad \text{Cov}(\hat\varepsilon_i,\hat\varepsilon_j)=-\sigma^2h_{ij}\neq0$$

$$\text{standardised residual}\quad r_i=\frac{\hat\varepsilon_i}{\hat\sigma\sqrt{1-h_{ii}}}\qquad\text{— this is *why* we standardise}$$

$$\text{Cook's distance}\quad D_i=\frac{r_i^2}{p}\cdot\frac{h_{ii}}{1-h_{ii}}$$

| Concept | Unusual in | Measured by |
|---|---|---|
| Leverage | $\boldsymbol{x}$ | $h_{ii}$ |
| Outlier | $y$ | $\lvert r_i\rvert$ |
| Influence | both jointly | $D_i$ |

**High leverage alone is harmless** — it hurts only when paired with a large residual.

| R plot | Detects |
|---|---|
| Residuals vs fitted | non-linearity (curve), heteroscedasticity (funnel) |
| **QQ plot** | non-normality — points follow the **45° diagonal**, not a horizontal line |
| Scale–location | heteroscedasticity specifically |
| Residuals vs leverage | influential points (Cook's D) |

### Multicollinearity

Perfect ⟹ breaks A5, $\boldsymbol{X}'\boldsymbol{X}$ singular, **not identified**.
Near ⟹ still **unbiased and still BLUE**, only **inflated** variances. **VIF ≈ 1 means no problem**; concern above 5–10.

---

## §9 — Chapter 2 carry-ins

### Dummy coding

> **$c$ levels ⟹ $c-1$ dummies.** The omitted level is the **reference** — it's the one missing from the R output. Every dummy coefficient is a difference **from the reference**; to compare two non-reference levels, **subtract** their coefficients.

All $c$ dummies plus an intercept ⟹ columns sum to the intercept column ⟹ $\boldsymbol{X}$ not full rank ⟹ no unique OLS. (*The dummy variable trap* — and the reason A5 exists.)

$$p = 1 + \#\{\text{continuous}\} + \sum(c-1) + \#\{\text{interactions}\},\qquad \text{df}=n-p$$

### Non-linear terms

**Polynomial:** effect $=\beta_1+2\beta_2 x$ — **not** $\beta_1$. Turning point at $-\hat\beta_1/(2\hat\beta_2)$.
**Interaction:** $y=\beta_0+\beta_1x+\beta_2D+\beta_3(x\cdot D)$; at $D=1$ the intercept is $\beta_0+\beta_2$ and the slope $\beta_1+\beta_3$. Dummy alone ⟹ **parallel** lines; plus interaction ⟹ **non-parallel**.

> ⚠️ **A variable appearing in two or more terms must be differentiated. Never quote one coefficient alone.**

### Logit

$$\pi=\frac{e^{\eta}}{1+e^{\eta}},\qquad \frac{\pi}{1-\pi}=e^{\eta},\qquad \log\frac{\pi}{1-\pi}=\eta=\boldsymbol{x}'\boldsymbol\beta$$

| Scale | Effect of $+1$ in $x_j$ |
|---|---|
| log-odds | $+\hat\beta_j$ — exact, constant |
| **odds** | $\times\exp(\hat\beta_j)$ — the **odds ratio** |
| probability | $\hat\beta_j\,\pi(1-\pi)$ — **not** constant |

**Why linear regression fails for binary $y$:** ① predictions escape $[0,1]$ · ② $\text{Var}(y)=\pi(1-\pi)$ ⟹ heteroscedastic by construction · ③ errors take only two values ⟹ cannot be normal · ④ a constant marginal effect is implausible near the boundaries.

Fitted by **maximum likelihood**, not OLS. Probit is the same with $\Phi$.

---

## §10 — Sentence templates (free marks)

**Slope:** *"A one-[unit] increase in [x] is associated with an estimated [$\hat\beta_1$] [unit] change in the expected [y], holding all other covariates fixed."*

**Dummy:** *"[Level] earns on average [$\hat\beta_j$] [unit] more than [reference level], holding all other covariates fixed."*

**Odds ratio:** *"A one-unit increase in [x] multiplies the odds of [event] by $\exp(\hat\beta_j)$ = [value], holding all other covariates fixed."*

**Test decision:** *"Since $|t| = [\cdot] > t_{n-p}(0.975) = [\cdot]$, we reject $H_0$ at the 5% level; [x] is significantly associated with [y]."*

**A3/A4 violated:** *"OLS remains unbiased and consistent, since unbiasedness needs only correct specification, zero-mean errors and full rank. But it is no longer efficient — Gauss–Markov no longer applies, so OLS is not BLUE and WLS/GLS would have smaller variance. The usual standard errors are also biased, so t-tests, F-tests and CIs are invalid."*

Words that earn marks: **"associated with"** (not "causes") · **"expected"** / "on average" · **"holding all other covariates fixed"** · **"fail to reject"** (not "accept") · **"at least one"** (not "all").
