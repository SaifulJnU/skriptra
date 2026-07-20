# Ch 1 — TRICKS & TIPS

---

## 1. The "count the betas" rule

Never memorise `n−p` or `n−k−1` as symbols. Do this instead:

> **Look at the model equation. Count how many $\beta$'s appear. Subtract from $n$.**

$$y = \beta_0 + \beta_1x_1 + \beta_2x_2 + \beta_3x_3 + \varepsilon$$
Four betas → residual df $= n - 4$.

Works under every convention, every paper, every time. **This single habit is worth 2–4 marks across a typical paper.**

---

## 2. Dimension-checking as an error detector

Before you compute anything with matrices, write the dimensions underneath:

$$\underset{p\times 1}{\hat{\boldsymbol\beta}} = \underset{p\times p}{(\boldsymbol{X}'\boldsymbol{X})^{-1}}\ \underset{p\times n}{\boldsymbol{X}'}\ \underset{n\times 1}{\boldsymbol{y}}$$

If inner dimensions don't match, the formula is wrong — **before** you've spent five minutes on arithmetic. Under exam time pressure this is the cheapest insurance you can buy.

**Instant scalar-detector:** anything of the form $\boldsymbol{a}'\boldsymbol{M}\boldsymbol{a}$ with $\boldsymbol{a}$ a column vector is a **scalar**. So $\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0$ — which looks terrifying in the prediction-interval formula — is just a number, and the exam will hand it to you (Sheet 4 Ex 3(e): "use $\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0 = 0.0035$").

---

## 3. The interpretation sentence template

Every "interpret this coefficient" question is 1–2 free marks in about 20 seconds if you have this memorised:

> **"Holding all other covariates fixed, a one-[unit] increase in [covariate] is associated with an estimated [$\hat\beta_j$] [unit] change in the expected [response]."**

Fill in the brackets. Then adapt:

| Situation | Template |
|---|---|
| Continuous $x$ | "…a one-year increase in **age** is associated with an estimated **0.71 dollar** increase in expected hourly **wage**, holding education fixed." |
| Dummy variable | "…**HS graduates** earn on average an estimated **11.44 dollars** more per hour than the reference category **< HS Grad**, holding age fixed." |
| $\log(y)$ model | "…a one-unit increase in $x$ is associated with an estimated **$100\hat\beta_j$ percent** change in $y$." |
| Intercept | "…the expected response when **all covariates equal zero**" — and immediately say whether that's meaningful. |

**Two words that earn and lose marks:**
- ✅ say **"associated with"** — never "causes"
- ✅ say **"holding all other covariates fixed"** — this phrase alone is often worth a mark in multiple regression

---

## 4. Why $\hat\beta_0$ often should *not* be interpreted

Sheet 1, Exercise 1(c) asks this directly. The answer:

> $\hat\beta_0$ is the expected wage for a man of **age 0**, which is outside the range of the data and biologically meaningless. Interpreting it would be **extrapolation**. The intercept is needed to position the regression line correctly, but it carries no substantive meaning here.

**Rule:** the intercept is interpretable only when $x = 0$ is (i) inside or near the observed data range and (ii) substantively meaningful. Otherwise say "it anchors the line but should not be interpreted."

**Trick to make it interpretable:** **centre** the covariate. Replace $x$ by $(x - \bar{x})$, or as the exam does, $(\text{age} - 48)$. Then $\beta_0$ = expected response at the *average* (or reference) value — meaningful, and inside the data range.

> This is exactly why Exam Summer 2025 writes $(\text{age}_i - 48)^2$ rather than $\text{age}_i^2$. Centring at 48 makes the coefficients interpretable and reduces collinearity between the linear and quadratic terms.

---

## 5. $R^2 \leftrightarrow r$ conversions (simple regression only)

$$r = \text{sign}(\hat\beta_1)\cdot\sqrt{R^2}$$

Mental square roots you should be able to do without a calculator:

| $R^2$ | $\sqrt{R^2}$ |
|---|---|
| 0.01 | 0.10 |
| 0.04 | 0.20 |
| 0.09 | 0.30 |
| 0.16 | 0.40 |
| 0.25 | 0.50 |
| 0.36 | 0.60 |
| 0.49 | 0.70 |
| 0.64 | 0.80 |

For anything else, interpolate: $\sqrt{0.038} \approx$ a bit under $\sqrt{0.04} = 0.20$, so $\approx 0.195$. ✓

⚠️ **Only in simple regression.** In multiple regression, $R^2 = \text{corr}(y,\hat y)^2$, not the squared correlation with any single covariate.

---

## 6. Skew detection in one second

> **mean > median ⟹ right-skewed (long right tail)**
> **mean < median ⟹ left-skewed**
> **mean ≈ median ⟹ roughly symmetric**

If they hand you summary statistics and the mean is well above the median, you have licence to say "right-skewed, consider a log transformation" — a fast, correct, mark-earning observation.

---

## 7. Which plot detects what — memorise this grid

You'll reuse it verbatim in Section 3.4.4.

| Plot | Detects |
|---|---|
| Histogram of $y$ | skewness, multimodality |
| Scatter $y$ vs $x$ | non-linearity, heteroscedasticity, outliers |
| Boxplots of $y$ by category | group differences (⟹ dummies will help) |
| Residuals vs fitted | non-linearity (curved band), heteroscedasticity (funnel) |
| QQ plot of residuals | non-normality (S-shape or heavy tails) |
| Scale–location | heteroscedasticity specifically |
| Residuals vs leverage / Cook's D | influential observations |

---

## 8. Exam-day micro-tactics from Chapter 1

- **Write the formula before the numbers.** The paper says "provide sufficient reasons." A formula with a wrong number scores far better than a right number alone.
- **Round to 3 decimals** — stated explicitly on the Example Exam paper.
- **Never leave a TRUE/FALSE blank.** Negative marking is floored at zero *per block*, so a guess can only help.
- **Underline the units** in any interpretation ("dollars per hour", "percent", "euros per m²"). Examiners look for them.
- If asked "why is this a regression problem," the answer is always some form of: *there is a systematic relationship contaminated by random variation, and we want to separate the two.*
