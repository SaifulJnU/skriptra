# 3.4.4 — Model Diagnosis

> **Every assumption from 3.1.2 gets checked here.** Highly examinable as both T/F statements and "what does this plot tell you" questions. The good news: it's Chapter 1.2's four questions, asked of residuals instead of raw data.

---

## 1. Why we need special residuals

Even under **perfect** assumptions ($\boldsymbol\varepsilon\sim N(\boldsymbol{0},\sigma^2\boldsymbol{I})$), the residuals are **not** iid normal:

$$\text{Cov}(\hat{\boldsymbol\varepsilon})=\sigma^2(\boldsymbol{I}-\boldsymbol{H}) \qquad\Longrightarrow\qquad \text{Var}(\hat\varepsilon_i)=\sigma^2(1-h_{ii})$$

- **Unequal variances** — depend on leverage $h_{ii}$
- **Correlated** — $\text{Cov}(\hat\varepsilon_i,\hat\varepsilon_j)=-\sigma^2h_{ij}\neq0$

> 🔴 *RCLM WS 22/23, Block III(i):* "Since the residuals are normally distributed, standardized residuals are also normally distributed."
> The premise itself is the giveaway — the whole reason standardisation exists is that raw residuals **don't** behave like the errors. High-leverage points have *smaller* residual variance, which artificially shrinks them and can hide outliers.

### The three residual types

| Type | Formula | Purpose |
|---|---|---|
| **Raw** | $\hat\varepsilon_i=y_i-\hat y_i$ | in the units of $y$; interpretable |
| **Standardised** | $r_i=\dfrac{\hat\varepsilon_i}{\hat\sigma\sqrt{1-h_{ii}}}$ | ≈ unit variance; comparable across observations |
| **Studentised (deleted)** | $t_i=\dfrac{\hat\varepsilon_i}{\hat\sigma_{(i)}\sqrt{1-h_{ii}}}$ | $\hat\sigma_{(i)}$ excludes obs $i$ ⟹ an outlier can't inflate its own denominator |

**Rule of thumb:** $|r_i|>2$ is notable, $|r_i|>3$ is a probable outlier.

> 💰 **Sheet 3, Ex 2(d)** asks for John's standardised residual, giving $h_{ii}=0.0016$:
> $$r=\frac{\hat\varepsilon}{\hat\sigma\sqrt{1-h_{ii}}}=\frac{-1.22}{35.724\times\sqrt{0.9984}}=\frac{-1.22}{35.695}=\boxed{-0.034}$$
> Far inside $\pm2$ — the model fits John very well.

---

## 2. ⭐ The four standard diagnostic plots

This grid is worth memorising verbatim — it answers "which plot detects what" instantly.

| # | Plot | Axes | Detects | Good looks like | Bad looks like |
|---|---|---|---|---|---|
| **1** | **Residuals vs Fitted** | $\hat\varepsilon_i$ vs $\hat y_i$ | **non-linearity**, **heteroscedasticity** | structureless horizontal band around 0 | curved/U pattern (non-linearity); **funnel/fan** (heteroscedasticity) |
| **2** | **QQ plot** | ordered $r_i$ vs normal quantiles | **non-normality** | points on the 45° line | S-shape (skew); ends curving away (heavy tails) |
| **3** | **Scale–Location** | $\sqrt{\lvert r_i\rvert}$ vs $\hat y_i$ | **heteroscedasticity** (specifically) | flat band, horizontal smoother | upward (or downward) trend |
| **4** | **Residuals vs Leverage** | $r_i$ vs $h_{ii}$, Cook's contours | **influential points** | no point beyond Cook's D $=0.5$ | points in the top/bottom right corner |

> 🔴 *Linear_model_exam_sheet, Block I(ii):* "Residual plots **cannot** be used to identify non-linearity in a regression model." → **FALSE.** That's exactly what plot 1 is for.

### More plots and what they add

| Plot | Detects |
|---|---|
| Residuals vs each covariate $x_j$ | non-linearity **in that specific covariate** — tells you *which* one to transform |
| Residuals vs time / index | **autocorrelation** (runs, cycles) |
| Histogram of residuals | non-normality, skew |
| Partial residual (component+residual) plot | the correct functional form for $x_j$ |

**Plot 1 tells you *something* is non-linear; residuals vs $x_j$ tells you *which* covariate.** That distinction is worth a mark.

---

## 3. The QQ plot — read it properly

Plot the **ordered standardised residuals** against the corresponding **theoretical normal quantiles**. If the residuals are normal, points fall on the diagonal.

| Pattern | Diagnosis |
|---|---|
| Points on the straight line | ✅ normality plausible |
| **S-shape** | skewness |
| Both ends **above** the line at right, **below** at left | heavy tails (more extremes than normal) |
| Both ends bending toward the line | light tails |
| A few isolated points off at the ends | outliers |

> 🔴 **RCLM WS 22/23, Block I(iv):** *"In Q-Q plots, the empirical quantiles are compared to the quantiles of the theoretical distribution. If the data follows the distribution, the points should closely follow a **horizontal line through the origin**."* → **FALSE.**
>
> The first half is right; the ending is wrong. Points should follow the **45° diagonal line** ($y=x$), **not a horizontal line**. A horizontal line would mean every empirical quantile is identical — i.e. no variation at all.
>
> **This is a classic "true first clause, false second clause" construction.** Read T/F statements to the very end.

---

## 4. Leverage, outliers and influence — three different things

Students conflate these constantly. Keep them separate.

| Concept | Meaning | Measured by | Unusual in |
|---|---|---|---|
| **Leverage** | unusual **covariate** values | $h_{ii}$ | $\boldsymbol{x}$ |
| **Outlier** | unusual **response** given $\boldsymbol{x}$ | $\lvert r_i\rvert$ large | $y$ |
| **Influence** | **actually changes** the fit | Cook's $D_i$ | both matter |

$$\text{high influence} \approx \text{high leverage} \times \text{large residual}$$

**The key asymmetry:** high leverage alone is *not* a problem. A point far out in $x$-space that sits exactly on the regression line is high-leverage and harmless — in fact it *improves* precision. It only becomes dangerous when it *also* has a large residual, because then it drags the line toward itself.

### Leverage

$$h_{ii}=\boldsymbol{x}_i'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_i,\qquad 0\leq h_{ii}\leq1,\qquad \sum_ih_{ii}=p$$

Average leverage $=p/n$. **Rule of thumb: $h_{ii}>2p/n$ is high.**

### Cook's distance

$$D_i=\frac{\sum_{j}(\hat y_j-\hat y_{j(i)})^2}{p\,\hat\sigma^2}=\frac{r_i^2}{p}\cdot\frac{h_{ii}}{1-h_{ii}}$$

**Read the second form.** It is literally (outlyingness) × (leverage) — both factors must be non-trivial for $D_i$ to be large.

**Meaning:** how much all the fitted values change when observation $i$ is deleted.
**Rule of thumb:** $D_i>0.5$ warrants attention; $D_i>1$ is influential.

> 🔴 *Linear_model_exam_sheet, Block III(i):* "Cook's distance helps identify points that might be unduly influencing the model's fit." → **TRUE.**
> 🔴 *Block III(ii):* "The hat matrix provides the leverages which help identify potential outliers." → **TRUE** per the key — though strictly leverage flags unusual **x**-values, and a low-leverage point can still be a $y$-outlier.

**What to do with an influential point:** never delete it reflexively. Check for a data error first. If it's genuine, report results **with and without** it, and consider whether the model form is wrong (an influential point often signals missing curvature or a missing covariate).

---

## 5. Diagnosing each assumption — the master grid

| Assumption | Check with | Failure looks like | Remedies |
|---|---|---|---|
| **A1 Linearity** | residuals vs fitted; residuals vs each $x_j$ | curved pattern | add $x^2$, $\log x$; add omitted covariate; transform $y$ |
| **A3 Homoscedasticity** | residuals vs fitted; scale–location | funnel / fan | log-transform $y$; weighted LS; robust SEs |
| **A4 Independence** | residuals vs time/index; Durbin–Watson | runs, cycles | time-series model; cluster-robust SEs; GLS |
| **A5 Full rank** | VIF; correlation matrix; software warnings | VIF $>5$–10 | drop or combine covariates; ridge; more data |
| **A6 Normality** | QQ plot; histogram of residuals | S-shape, heavy tails | transform $y$; rely on CLT for large $n$ |
| **Outliers/influence** | residuals vs leverage; Cook's D | isolated corner points | investigate; report with and without |

---

## 6. Multicollinearity diagnostics

$$\text{VIF}_j=\frac{1}{1-R_j^2}$$

where $R_j^2$ is from regressing $x_j$ on **all the other covariates**.

| VIF | Interpretation |
|---|---|
| $=1$ | $x_j$ **uncorrelated** with the others — ideal |
| $<5$ | fine |
| $5$–$10$ | moderate concern |
| $>10$ | serious ($R_j^2>0.9$) |

$\text{Var}(\hat\beta_j)$ is inflated by exactly the factor $\text{VIF}_j$ relative to an orthogonal design.

> 🔴 *Linear_model_exam_sheet, Block II(ii):* "If the VIF for all variables is close to 1, multicollinearity is likely a concern." → **FALSE.** VIF ≈ 1 means **no** collinearity. Perfect situation. The statement inverts the scale.

**Remember (from 3.1.2):** near-multicollinearity leaves $\hat{\boldsymbol\beta}$ **unbiased and still BLUE**. It only inflates variances. Wide CIs, insignificant $t$'s, unstable signs — but nothing is *wrong*, just imprecise.

---

## 7. 📝 Model answer: "Diagnose this model from the residual plots"

A template you can adapt:

> *I would examine four plots.*
>
> *(1) **Residuals versus fitted values.** Under a correctly specified model this should show a structureless horizontal band centred on zero. A systematic curve would indicate a violation of the linearity assumption — suggesting a missing polynomial term or transformation — while a funnel shape, where the spread widens with the fitted values, would indicate heteroscedasticity.*
>
> *(2) **A QQ plot of the standardised residuals** against normal quantiles, to assess normality. Points should lie close to the 45° line; an S-shape indicates skewness and departures at both ends indicate heavy tails. Since normality is required only for the exactness of the $t$- and $F$-tests, mild departures are tolerable when $n$ is large.*
>
> *(3) **A scale–location plot** ($\sqrt{|r_i|}$ against fitted values) as a more sensitive check for heteroscedasticity — a clear upward trend confirms non-constant variance.*
>
> *(4) **Residuals versus leverage**, with Cook's distance contours, to identify influential observations. Points with high leverage $h_{ii}$ **and** large standardised residuals, giving Cook's $D_i>0.5$, would be examined individually.*
>
> *If heteroscedasticity were present, OLS would remain unbiased and consistent but would no longer be BLUE, and the usual standard errors would be invalid; I would consider a log transformation of the response, weighted least squares, or heteroscedasticity-robust standard errors.*

**That last paragraph — connecting the diagnosis back to the consequences from 3.1.2 — is what earns the top mark.** Diagnosis alone is description; linking it to unbiasedness/efficiency/validity is analysis.

---

## 8. Key takeaways

1. Residuals are **not** iid even under perfect assumptions: $\text{Var}(\hat\varepsilon_i)=\sigma^2(1-h_{ii})$, and they're correlated. Hence standardisation.
2. $r_i=\dfrac{\hat\varepsilon_i}{\hat\sigma\sqrt{1-h_{ii}}}$; $|r_i|>2$ notable, $>3$ probable outlier.
3. **The four plots:** residuals-vs-fitted (linearity + heteroscedasticity) · QQ (normality) · scale–location (heteroscedasticity) · residuals-vs-leverage (influence).
4. 🔴 **QQ plot points follow the 45° DIAGONAL, not a horizontal line.**
5. **Leverage ≠ outlier ≠ influence.** Cook's $D_i=\frac{r_i^2}{p}\cdot\frac{h_{ii}}{1-h_{ii}}$ — outlyingness × leverage.
6. $h_{ii}>2p/n$ is high leverage; $D_i>0.5$ warrants attention.
7. **VIF $=1$ means NO collinearity.** Concern above 5–10.
8. Always link a diagnosis back to its **consequence** (unbiased? efficient? valid inference?) and name a **remedy**.
