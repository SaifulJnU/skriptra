# Ch 2 — EXERCISES

**Attempt everything before opening solutions.** Target time: 70 minutes.

---

## Part A — TRUE / FALSE

**A1.** In a multiple linear regression, if $\hat\beta_j > 0$ then the correlation between $x_j$ and $y$ must be positive.

**A2.** A categorical covariate with $m$ categories requires $m$ dummy variables when an intercept is included.

**A3.** In a logit model, an increase of 1 in $x_{j}$ is interpreted as an increase of $\hat\beta_j$ in $P(y_i=1)$.

**A4.** In a logit model, $\exp(\hat\beta_j)$ gives the multiplicative change in the odds of $y=1$ per one-unit increase in $x_j$.

**A5.** In a model with an interaction term $x_1x_2$, the coefficient $\hat\beta_1$ can be interpreted as the effect of $x_1$ on $y$.

**A6.** Adding an interaction term between a continuous covariate and a dummy makes the two group-specific regression lines non-parallel.

**A7.** The average of the fitted values equals the average of the observed responses, provided the model contains an intercept.

**A8.** A linear probability model can produce fitted values outside the interval $[0,1]$.

**A9.** Interactions can only be computed between two continuous covariates, or between a continuous and a categorical covariate.

**A10.** In the model $y = \beta_0+\beta_1x+\beta_2x^2+\varepsilon$, the estimated effect on $y$ of a one-unit increase in $x$ is $\hat\beta_1$.

---

## Part B — Model building

**B1.** (3 points) A university wants to model students' final exam scores. Available covariates:
- hours studied per week (continuous)
- degree programme (Statistics / Data Science / Mathematics / Economics)
- whether the student attended the tutorial (yes/no)

(a) Define all dummy variables explicitly and write the model equation with an intercept.
(b) State the reference category.
(c) How many parameters are estimated? What are the residual degrees of freedom if $n = 240$?

**B2.** (3 points) An insurer models annual medical cost on `age`, `smoker` (yes/no), and `BMI category` (underweight / normal / overweight / obese). The insurer suspects the effect of age is **stronger for smokers**.

(a) Write the model equation, including the appropriate interaction.
(b) Give the intercept and the age-slope separately for smokers and non-smokers.
(c) Explain in one sentence what the interaction coefficient means.

**B3.** (2 points) Explain why including all $c$ dummy variables for a $c$-level categorical covariate *together with an intercept* makes the OLS estimator undefined. Refer explicitly to $\boldsymbol{X}$ and $\boldsymbol{X}'\boldsymbol{X}$.

---

## Part C — Interpretation from output

**C1.** (4 points) A regression of `wage` (dollars/hour) gives:

| Term | Estimate |
|---|---|
| (Intercept) | 52.61 |
| age | 0.62 |
| education: HS Grad | 11.01 |
| education: Some College | 23.16 |
| education: College Grad | 37.97 |
| education: Advanced Degree | 62.63 |
| health: ≥ Very Good | 9.13 |

(a) How many levels does `education` have? What is the reference category?
(b) Interpret the coefficient on `age`.
(c) Interpret the coefficient on `health: ≥ Very Good`.
(d) Compute the predicted wage for a 37-year-old College Graduate in very good health.
(e) His actual wage is \$121.43. Compute and interpret his residual.
(f) What is the estimated wage difference between an Advanced Degree holder and a HS Grad of the same age and health?

**C2.** (3 points) From Sheet 2, the fitted model with interaction is
$$\widehat{\text{wage}} = 78.66 + 0.51\cdot\text{age} - 1.81\cdot H + 0.43\cdot\text{age}\cdot H$$
where $H=1$ if health ≥ Very Good.

(a) Give the intercept and slope for each health group.
(b) At what age do the two fitted lines cross?
(c) Is it correct to say "being in very good health reduces expected wage by \$1.81"? Explain.

**C3.** (2 points) A quadratic model gives $\hat\beta_0 = -10.43$, $\hat\beta_1 = 5.29$, $\hat\beta_2 = -0.05$ for wage on age and age².
(a) Compute the estimated effect on wage of a marginal increase in age, at age 30 and at age 65.
(b) At what age is expected wage maximised?

---

## Part D — Logit

**D1.** (3 points) Explain why a linear regression model is not appropriate to model a binary dependent variable, and why you should instead use a logit, probit or similar model. Give at least three distinct reasons.

**D2.** (3 points) A logit model for loan default gives $\hat\beta_{\text{duration}} = 0.028$ (duration in months) and $\hat\beta_{\text{ownsHome}} = -0.85$.

(a) Interpret $\hat\beta_{\text{duration}}$ on the odds scale.
(b) Interpret $\hat\beta_{\text{ownsHome}}$ on the odds scale.
(c) A colleague says "owning a home reduces the default probability by 0.85." Correct him precisely.
(d) For a customer currently at $\hat\pi = 0.5$, what is the approximate marginal effect of one extra month of duration on the **probability** of default?

**D3.** (2 points) Write down the logit model in all three equivalent forms (probability, odds, log-odds), and state which of the three makes it a *generalised linear* model and why.

---

## Part E — Exam-realistic

**E1.** (3 points, Exam Summer 2025 Ex 2(a) style)
You want to examine the effect of a person's **age**, **education** (no degree / high school degree / college or university degree) and **place of birth** (inside the US / outside the US) on their **hourly wage**, using a linear model estimated by OLS on a random sample of employed US residents.

Provide an adequate linear model equation incorporating all explanatory variables, including an intercept, formulated so that you can estimate this exact equation by OLS.

**E2.** (2 points, Exam Summer 2025 Ex 2(b)–(c) style)
A colleague suggests adding the regressor $\widetilde{\text{age}}^2_i := (\text{age}_i - 48)^2$ to the equation in E1.

(a) Explain briefly why this might be a sensible idea.
(b) Would you expect the OLS estimate of its coefficient to be positive or negative? Explain.
(c) Why subtract 48 rather than simply using $\text{age}_i^2$?
