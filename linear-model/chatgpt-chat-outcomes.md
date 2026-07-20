# ChatGPT Chat Outcomes — RC Linear Models Exam Prep

> Archive of a prior ChatGPT conversation about math prerequisites, exam coverage, and repeated topics for the "RC Linear Models" exam (5 ECTS). Reading course scope: **Chapters 1 – 3.4**.

---

## Part 1 — What math do I actually need?

No advanced pure math or proofs required. Level = first-year university calculus, linear algebra, probability/statistics.

| # | Area | Priority | What you need |
|---|------|----------|----------------|
| 1 | Algebra | ⭐⭐⭐⭐⭐ | Linear equations, rearranging formulas, fractions, exponents, summation notation (Σ), basic logs |
| 2 | Functions | ⭐⭐⭐⭐ | Linear, quadratic, exponential, logarithmic. `log(x)` matters for logistic regression's `log(p/(1−p))` |
| 3 | Matrices & Vectors | ⭐⭐⭐⭐⭐ | Vector, matrix, matrix multiplication, transpose. Needed for `Y = Xβ + ε` |
| 4 | Basic Calculus | ⭐⭐⭐ | Derivatives, partial derivatives (for minimizing `(y − ŷ)²`). No hard integration |
| 5 | Probability | ⭐⭐⭐⭐⭐ | Random variable, `E(X)`, `Var(X)`, `Cov(X,Y)`, independence |
| 6 | Statistics | ⭐⭐⭐⭐⭐ | Mean, median, variance, SD, sample vs population, normal distribution (`ε ~ N(0, σ²)`) |
| 7 | Hypothesis Testing | ⭐⭐⭐⭐⭐ | H₀, H₁, p-value, significance level, confidence interval, t-test, F-test |
| 8 | Least Squares | ⭐⭐⭐⭐⭐ | Why `Σ(yᵢ − ŷᵢ)²` is minimized |
| 9 | Correlation | ⭐⭐⭐⭐ | Covariance, Pearson correlation, interpreting sign/strength |
| 10 | Basic Graphs | ⭐⭐⭐ | Scatter plots, regression line, residual plots, histograms, QQ plots |

**Example calculation that recurs throughout:** `y = β₀ + β₁x`; with β₀ = 2, β₁ = 5, x = 3 → `y = 2 + 5(3) = 17`.

### Not needed
❌ Differential equations · complex numbers · Fourier transforms · topology · abstract algebra · real analysis · multivariable integration

### Difficulty by chapter

| Chapter | Math level | Difficulty |
|---------|-----------|------------|
| Chapter 1 | ⭐ | Easy |
| Chapter 2.1–2.3 | ⭐⭐ | Easy–Medium |
| Chapter 3.1 | ⭐⭐⭐ | Medium |
| Chapter 3.2 | ⭐⭐⭐⭐ | Medium–Hard |
| Chapter 3.3 | ⭐⭐⭐⭐ | Hard (statistics) |
| Chapter 3.4 | ⭐⭐⭐ | Medium |

### Suggested learning order (from scratch)
1. Basic algebra
2. Functions and graphs
3. Mean, variance, standard deviation
4. Probability basics
5. Matrices and vectors
6. Regression line
7. Least squares
8. Hypothesis testing
9. Confidence intervals
10. Model selection and diagnostics

---

## Part 2 — Coverage analysis of the previous-year exam (Exam 1)

**Short answer:** studying only the previous exam is *not* enough — but mastering everything needed to solve it covers roughly **85–95%** of what is likely to appear, assuming an unchanged syllabus.

### Chapter mapping

| Chapter | Coverage | Where it shows up |
|---------|----------|-------------------|
| 1 — Introduction | ⭐ (light) | Choosing variables, understanding regression, interpreting data. No heavy theory |
| 2.1–2.3 | Yes | Ex 1(h) logit interpretation; Ex 4(a) why logit over linear. Binary response, probability in [0,1], why linear regression fails. **No derivation of logit required** |
| 3.1 | A LOT | Model definition, assumptions, regression equation, residuals, intercept, dummy variables, design matrix → Ex 2, Ex 4(d), Ex 1 |
| 3.2 | Huge | OLS, least squares, parameter estimation, BLUE, variance, matrix rank, MLE → Ex 4(b), Ex 3 |
| 3.3 | Very heavy | t-test, F-test, CI, joint hypothesis, restrictions, p-values → Ex 3 almost entirely |
| 3.4 | Very important | Model selection, information criteria, adjusted R², diagnostics → Ex 2(d), Ex 4(c), Ex 1(c) |

### Repeatedly tested (⭐⭐⭐⭐⭐)
- **OLS** — intuition, objective `min Σeᵢ²`, assumptions, properties
- **Regression equation** — writing `y = β₀ + β₁x₁ + …` with dummy variables and quadratic terms
- **Confidence interval** — must calculate
- **t-test** — must calculate
- **F-test** — must calculate
- **Joint hypothesis** — `Cβ = d` (surprises many students)
- **Model selection** — AIC, BIC, adjusted R², bias–variance tradeoff
- **Interpretation** — sign of coefficients and why
- **Regression output** — fill in missing values from R output, residual variance, t-value, CI

### Not in Exam 1 but could still appear
Multiple regression assumptions (linearity, homoscedasticity, normality, independence) · hat / projection matrix · residual properties · leverage · Cook's distance · prediction interval vs confidence interval · ANOVA table · R² derivation · residual plots · QQ plots

### Probably not needed
GLM · Poisson · mixed models · GAM · quantile regression — beyond the basic intuition in Chapter 2.

### Two-week study priority

**Priority 1 (must know):** linear model equation · OLS estimation · OLS assumptions · dummy variables · coefficient interpretation · confidence intervals · t-tests · F-tests · joint hypotheses · R regression output · model selection (AIC, BIC, adj. R²) · model diagnostics · logit intuition

**Priority 2:** matrix notation · BLUE · Gauss–Markov theorem · MLE relationship · multicollinearity · bias vs variance · homoscedasticity vs heteroscedasticity

**Priority 3:** proofs · matrix derivations · detailed algebra behind OLS

**Overall impression:** the exam is **application-oriented, not proof-oriented**. Expect to interpret regression results, perform standard calculations, write regression models, explain concepts, and justify modeling choices.

---

## Part 3 — Comparison of Exam 1 vs Exam 2: repetition analysis

**The two exams are ~80–90% similar** — same concepts, different data and wording.

### Topics that repeat (very high probability)

| # | Topic | Exam 1 | Exam 2 |
|---|-------|--------|--------|
| 1 | TRUE/FALSE conceptual block ⭐⭐⭐⭐⭐ | BLUE, OLS, multicollinearity, CI, F-test, logit, dummies, MLE, matrix rank | BLUE, OLS, dummies, CI, residuals, multicollinearity, RSS, AIC, trend |
| 2 | OLS ⭐⭐⭐⭐⭐ | Explain mathematically | Explain and derive |
| 3 | Gauss–Markov / BLUE ⭐⭐⭐⭐⭐ | TRUE/FALSE | Full question: explain the theorem |
| 4 | Confidence interval ⭐⭐⭐⭐⭐ | 99% CI | 95% CI |
| 5 | t-test ⭐⭐⭐⭐⭐ | Yes | Yes — H₀, H₁, statistic, df, conclusion |
| 6 | F-test ⭐⭐⭐⭐⭐ | Joint hypothesis | Model comparison |
| 7 | Model selection ⭐⭐⭐⭐⭐ | Information criteria | AIC, BIC, alternative methods |
| 8 | Regression interpretation ⭐⭐⭐⭐⭐ | Everywhere | Everywhere |
| 9 | Dummy variables ⭐⭐⭐⭐ | Construct regression | TRUE/FALSE |
| 10 | Quadratic term ⭐⭐⭐⭐ | `(age − 48)²` | `horsepower²` |
| 11 | Diagnostics ⭐⭐⭐⭐ | Model diagnosis | Residual plots, heteroscedasticity, normality, nonlinearity |
| 12 | Logit ⭐⭐⭐⭐ | Why not linear regression | TRUE/FALSE |

### New in Exam 2 only
- **Regression calculation by hand** (big): `β̂ = (X′X)⁻¹X′Y`, compute coefficients, RSS, variance, R²
- **ESS**: know `TSS = ESS + RSS`
- **Matrix calculation**: build X, X′X, invert, solve for β
- **Diagnostic plots**: identify heteroscedasticity, QQ plot, nonlinear pattern

### In Exam 1 only
- **Restricted model** / `Cβ = d` (harder)
- **Missing R output** — fill in `[[A]]`, `[[B]]`, `[[C]]`, `[[D]]`
- **Joint hypothesis** as matrix restrictions

### Approximate exam composition
- ~20% conceptual knowledge (True/False, assumptions, interpretations)
- ~40% calculations (OLS, CI, t-tests, F-tests, regression output)
- ~40% modeling and reasoning (writing models, choosing variables, model selection, diagnostics, justification)

### Prediction for the 2026 exam
- ✅ TRUE/FALSE block — definitely
- ✅ OLS derivation or explanation — very likely
- ✅ Write a regression model with dummy variables — very likely
- ✅ Compute a confidence interval — almost certain
- ✅ Perform a t-test or F-test — almost certain
- ✅ Interpret R output — very likely
- ✅ Discuss AIC/BIC or another model selection method — very likely
- ✅ Diagnose a regression model from residual plots — likely
- ✅ One short question on the logit model — likely

---

## Appendix — Previous-year exam (Exam 1), full text

### Exercise 1 (6 Points) — TRUE or FALSE

a) In a multiple linear regression model, adding a variable which is not correlated with the dependent variable will not affect the unbiasedness of the OLS estimator, but it may affect its variance. ☐ TRUE ☐ FALSE

b) In a linear regression model, minimizing the sum of the absolute deviations results in the same coefficient estimates as minimizing the sum of the squared deviations. ☐ TRUE ☐ FALSE

c) In a linear regression of people's income onto their years of professional experience, adding dummies for the weekday a person was born on as regressors can be expected to lower the R². ☐ TRUE ☐ FALSE

d) When the design matrix X does not have full column rank, the OLS estimates still exist and are unique as long as the error variance is constant. ☐ TRUE ☐ FALSE

e) A best linear unbiased estimator (BLUE) is "best" in the sense that there is no other linear unbiased estimator with a lower variance. ☐ TRUE ☐ FALSE

f) In a simple linear regression, the slope estimator β̂₁ is always uncorrelated with the intercept estimator β̂₀. ☐ TRUE ☐ FALSE

g) When testing H₀: βⱼ = 0 versus H₁: βⱼ ≠ 0 in a linear regression model, the t-statistic for βⱼ follows a t-distribution with n − k − 1 degrees of freedom under H₀, where k + 1 is the number of estimated parameters. ☐ TRUE ☐ FALSE

h) Consider a logit model where you regress yᵢ onto 1, x₁ᵢ, x₂ᵢ, …, x_kᵢ to obtain β̂₀, β̂₁, β̂₂, …, β̂_k. An increase of 1 in x_jᵢ is interpreted as an increase of β̂ⱼ in P(yᵢ = 1). ☐ TRUE ☐ FALSE

i) The F-statistic for testing H₀: β₁ = −β₂ + β₃ in a linear model with k ≥ 3 predictors plus an intercept has an F-distribution with (3, n − k − 1) degrees of freedom under H₀. ☐ TRUE ☐ FALSE

j) Multicollinearity in the design matrix X can inflate the variance of the OLS coefficient estimators. ☐ TRUE ☐ FALSE

k) If a confidence interval for a slope coefficient βⱼ does not contain zero, we fail to reject H₀: βⱼ = 0. ☐ TRUE ☐ FALSE

l) The OLS estimator in a linear regression is equivalent to the maximum likelihood estimator under the assumption of independent and identically normally distributed errors. ☐ TRUE ☐ FALSE

### Exercise 2 (8 Points)

You want to examine the effect of a person's age, education (no degree, high school degree or college/university degree) and place of birth (inside the US or outside the US) on their hourly wage. You sampled randomly from all employed people in the US (both full-time and part-time) and want to use a linear regression model estimated by OLS.

a) **(3 Points)** Provide an adequate linear model equation that incorporates all of the above explanatory variables, including an intercept. Formulate it such that you can estimate this exact equation by OLS.

b) **(2 Points)** Explain briefly why adding the following regressor might be a sensible idea: `age_g2ᵢ := (ageᵢ − 48)²`, where ageᵢ is the age of person i.

c) **(1 Point)** Would you expect the OLS estimate of the coefficient for `age_g2ᵢ` to be positive or negative? Explain your choice.

d) **(2 Points)** A colleague suggests adding many more functions of ageᵢ, e.g. `(ageᵢ − 48)³`, `ageᵢ⁴` or `log(ageᵢ)` because it will improve the fit. How would you decide whether you should do so or not? State two methods.

### Exercise 3 (8 Points) — R output

```r
> lm <- lm(medv ~ crim + nox + dis + rad)
> summary(lm)

Call:
lm(formula = medv ~ crim + nox + dis + rad)

Residuals:
    Min      1Q  Median      3Q     Max
-16.802  -5.152  -2.127   2.604  31.084

Coefficients:
             Estimate Std. Error t value Pr(>|t|)
(Intercept)  48.38458    [[A]]    13.591  < 2e-16 ***
crim         -0.25959   0.05302   -4.896 1.32e-06 ***
nox         -36.99122   5.25574   [[B]]   6.44e-12 ***
dis           [[C]]     0.26423   -3.796 0.000165 ***
rad          -0.06165   0.05983   -1.030 0.303290
---
Signif. codes:  0 '***' 0.001 '**' 0.01 '*' 0.05 '.' 0.1 ' ' 1

Residual standard error: [[D]] on 501 degrees of freedom
Multiple R-squared: 0.2583,  Adjusted R-squared: 0.2524
F-statistic: 43.62 on 4 and 501 DF,  p-value: < 2.2e-16

> sum(lm$residuals^2)
[1] 31682.02
```

a) **(2.5 Points)** Reproduce the values of [[A]], [[B]], [[C]] and [[D]]. Write down the fitted regression formula of the model.

b) **(2 Points)** Calculate a 99% confidence interval for β_nox, centered around β̂_nox. At the 1% significance level, would you reject H₀: β_nox = −30 against H₁: β_nox ≠ −30?

c) **(2 Points)** Test the joint null hypothesis H₀: β_crim = 3β_rad − 0.1, β_nox = −40. Express it as `Cβ = d`, i.e. identify C and d with β = (β₀, β_crim, β_nox, β_dis, β_rad)′. What is the number of linearly independent restrictions r?

d) **(1.5 Points)** Let `ε̂′_H₀ ε̂_H₀ = 32333.15`. Calculate the F-statistic and decide whether you would reject the null hypothesis from (c) at the 5% significance level.

### Exercise 4 (8 Points)

a) **(1 Point)** Explain why a linear regression model is not appropriate to model a binary dependent variable and why you should instead use a logit, probit or similar model.

b) **(2 Points)** Given `yᵢ = β₀ + β₁x₁ᵢ + … + β_k x_kᵢ + εᵢ`, explain the method of ordinary least squares (OLS) estimation and show the steps necessary to obtain β̂₀, β̂₁, …, β̂_k. It is not necessary to explicitly calculate them; the mathematical approach suffices.

c) **(2 Points)** Explain the intuition behind information criteria, and why their design makes them a good tool for model selection. Use any information criterion as an example.

d) **(2 Points)** For `yᵢ = β₀ + β₁x₁ᵢ + β₂x₂ᵢ + εᵢ`, you want to test H₀: β₁ = β₂ + 1 via an F-test, by estimating the model unrestricted and then under H₀ to obtain SSE and SSE_H₀. Incorporate the null hypothesis into the model equation to obtain a restricted model estimable by OLS directly.

e) **(1 Point)** Analyzing yearly revenue vs. number of employees, you notice the variation in revenue grows as the number of employees grows. Which impact does this phenomenon have on the OLS estimate in terms of bias and efficiency?

### Quantile tables

**Table 1 — Quantiles for some t-distributions**

| df \ quantile | 0.9 | 0.95 | 0.975 | 0.99 | 0.995 |
|---|---|---|---|---|---|
| 3 | 1.6377 | 2.3534 | 3.1824 | 4.5407 | 5.8409 |
| 7 | 1.4149 | 1.8946 | 2.3646 | 2.9980 | 3.4995 |
| 12 | 1.3562 | 1.7823 | 2.1788 | 2.6810 | 3.0545 |
| 13 | 1.3502 | 1.7709 | 2.1604 | 2.6503 | 3.0123 |
| 25 | 1.3163 | 1.7081 | 2.0595 | 2.4851 | 2.7874 |
| 41 | 1.3020 | 1.6819 | 2.0180 | 2.4184 | 2.6980 |
| 49 | 1.2990 | 1.6766 | 2.0095 | 2.4048 | 2.6799 |
| 50 | 1.2987 | 1.6759 | 2.0086 | 2.4033 | 2.6778 |
| 100 | 1.2901 | 1.6602 | 1.9840 | 2.3642 | 2.6259 |
| 250 | 1.2849 | 1.6510 | 1.9695 | 2.3414 | 2.5956 |
| 501 | 1.2832 | 1.6479 | 1.9647 | 2.3338 | 2.5857 |

**Table 2 — Quantiles for some F-distributions with fixed df = 501**

| r \ quantile | 0.9 | 0.95 | 0.975 | 0.99 |
|---|---|---|---|---|
| r = 2 | 2.3132 | 3.0137 | 3.7162 | 4.6478 |
| 3 | 2.0948 | 2.6227 | 3.1422 | 3.8209 |
| 4 | 1.9561 | 2.3897 | 2.8114 | 3.3568 |
| 5 | 1.8588 | 2.2320 | 2.5918 | 3.0539 |

---

## Next steps (as stated by Saiful)

1. Share all related resources in this folder: the book PDF, previous years' exam questions, and solutions if available.
2. Build a preparation plan together.
3. Goal: understand **every** topic that might appear, from every angle — derivation, theory, exercise, and any other question style — so that any question on a given topic can be answered confidently.
4. Target: full marks (course is 5 ECTS).
