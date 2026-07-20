# Ch 1 — EXERCISES

**Attempt everything before opening `21-SOLUTIONS.md`.** Time yourself: this set should take 35–45 minutes.

---

## Part A — TRUE / FALSE (exam Exercise 1 style)

Answer TRUE or FALSE **and write one line of justification.** In the real exam you don't write the justification, but here it's how you learn.

**A1.** In a regression model, the error term $\varepsilon_i$ and the residual $\hat\varepsilon_i$ are the same thing.

**A2.** The relation $y = \exp(\beta_0 + \beta_1x_1 + \dots + \beta_kx_k + \varepsilon)$ cannot be analysed within the linear regression framework.

**A3.** A model containing $x^2$ as a covariate is not a linear model.

**A4.** If the Pearson correlation between $x$ and $y$ is zero, then $x$ and $y$ are independent.

**A5.** In a simple linear regression with an intercept, the coefficient of determination equals the squared Pearson correlation between $x$ and $y$.

**A6.** The design matrix $\boldsymbol{X}$ of a model with $n$ observations, $k$ covariates and an intercept has dimension $n \times k$.

**A7.** The type of the response variable determines which class of regression model is appropriate.

**A8.** In the model $\log(y) = \beta_0 + \beta_1 x + \varepsilon$, the coefficient $\beta_1$ is interpreted as the percentage change in $y$ associated with a one-percent increase in $x$.

**A9.** $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}$ is a scalar equal to the sum of squared residuals.

**A10.** Two datasets with the same mean, variance and correlation must look the same when plotted.

---

## Part B — Short answer

**B1.** (2 points) Explain in your own words what the error term $\varepsilon$ represents and give three distinct reasons why it is needed in a regression model.

**B2.** (2 points) A researcher finds a strong positive correlation between the number of firefighters sent to a fire and the amount of property damage. She concludes that firefighters cause damage. Explain what is wrong, name the statistical phenomenon, and say how regression addresses it.

**B3.** (3 points) For each of the following studies, state (i) the response variable, (ii) two sensible covariates, (iii) the type of the response, and (iv) the appropriate model class.

- (a) Predicting a used car's selling price
- (b) Predicting whether a customer will default on a loan
- (c) Predicting how many insurance claims a policyholder files in a year

**B4.** (2 points) The distribution of hourly wages in the `Wage` dataset is strongly right-skewed. State two consequences this may have for a linear model fitted to raw wage, and state the standard remedy.

**B5.** (1 point) Write down the general linear model in matrix form and state the dimension of every object in it.

---

## Part C — Calculation

**C1.** (2 points) A simple linear regression of `wage` on `age` gives $R^2 = 0.038$ and $\hat\beta_1 = 0.71$.
(a) What is the correlation between wage and age?
(b) Interpret this value in one sentence.
(c) Give one plausible explanation for why the value is so small, given what you know from the exercise sheets.

**C2.** (3 points) You have $n = 5$ observations and want to fit $y_i = \beta_0 + \beta_1 x_{i1} + \beta_2 x_{i2} + \varepsilon_i$. The data are:

| $i$ | $y_i$ | $x_{i1}$ | $x_{i2}$ |
|---|---|---|---|
| 1 | 3 | 1 | 0 |
| 2 | 5 | 2 | 0 |
| 3 | 4 | 1 | 1 |
| 4 | 8 | 3 | 1 |
| 5 | 7 | 2 | 1 |

(a) Write down the design matrix $\boldsymbol{X}$ explicitly.
(b) State the dimensions of $\boldsymbol{X}$, $\boldsymbol\beta$, $\boldsymbol{X}'\boldsymbol{X}$ and $\boldsymbol{X}'\boldsymbol{y}$.
(c) What is the residual degrees of freedom of this model?

**C3.** (2 points) For a fitted model you are told $\text{SST} = 1000$ and $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon} = 260$.
(a) Compute $R^2$.
(b) Compute the explained sum of squares.

**C4.** (2 points) Given the four rules of expectation and variance, show that if $E(\hat\beta_1) = \beta_1$ and $\text{Var}(\hat\beta_1) = \sigma^2 c$ for some constant $c$, then for $\tilde\beta = 2\hat\beta_1 - \beta_1$ we have $E(\tilde\beta) = \beta_1$ and $\text{Var}(\tilde\beta) = 4\sigma^2 c$. Comment on which of $\hat\beta_1$ and $\tilde\beta$ you would prefer and why.

---

## Part D — Exam-realistic mini-question

**D1.** (4 points, Exam Exercise 2 style)
A city government wants to understand what drives monthly net rent for apartments. Available variables: living area (m²), year of construction, district (12 districts), and kitchen quality (standard / premium / none).

(a) Identify the response variable and its type. Which model class is appropriate? (1 pt)
(b) Before fitting anything, name three plots you would produce and say what you would look for in each. (2 pts)
(c) The histogram of net rent is strongly right-skewed. What would you do, and how would this change the interpretation of your coefficients? (1 pt)
