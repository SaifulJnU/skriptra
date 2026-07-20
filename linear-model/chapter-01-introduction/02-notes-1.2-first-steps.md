# 1.2 — First Steps: Looking Before Modelling

> **Purpose:** before you fit anything, you look. This section teaches what to look at. It is the ancestor of Section 3.4.4 (Model Diagnosis), which *is* heavily examined.

---

## 1.2.1 Univariate distributions of the variables

Before any model: examine **each variable on its own**.

### For a continuous variable

| Tool | What it shows | What you're hunting for |
|---|---|---|
| **Histogram** | shape of the distribution | skewness, multiple peaks, gaps |
| **Boxplot** | median, quartiles, outliers | extreme values, asymmetry |
| **Mean vs median** | centre | if mean ≫ median → right-skewed |
| **Variance / SD** | spread | scale, comparability |

**Key numbers (know these definitions cold):**

$$\bar{y} = \frac{1}{n}\sum_{i=1}^{n} y_i \qquad s^2 = \frac{1}{n-1}\sum_{i=1}^{n}(y_i - \bar{y})^2$$

Note the $n-1$. It appears because we used up one degree of freedom estimating $\bar{y}$. **This is the same logic that gives $n-p$ later in the course** — every parameter you estimate costs one degree of freedom. Notice it now; it will save you in 3.2.2.

### For a categorical variable

Frequency table + bar chart. Count the levels — because a categorical variable with $c$ levels becomes $c-1$ dummy variables in Section 3.1.3, and miscounting is a classic exam error.

### Why skewness matters (this is the examinable bit)

Income, wages, rents and house prices are **right-skewed**: a long tail of large values. If $y$ is strongly right-skewed:

- The normality assumption on $\varepsilon$ is likely violated
- Variance often grows with the mean → **heteroscedasticity**
- **Fix: model $\log(y)$ instead of $y$**

The log transformation compresses the long right tail toward symmetry. This is why you constantly see `log(rent)`, `logwage`, `log(price)` in applied work.

> **Exam link:** Exam Summer 2025, Ex 4(e): *"the variation in revenue grows as the number of employees grows."* That's heteroscedasticity, spotted exactly the way this section teaches — by looking. The log transform is a standard remedy worth mentioning.

**Interpretation after logging** (memorise, it's a free mark):

| Model | Interpretation of $\beta_1$ |
|---|---|
| $y = \beta_0 + \beta_1 x$ | 1-unit ↑ in $x$ → $\beta_1$-unit change in $y$ |
| $\log(y) = \beta_0 + \beta_1 x$ | 1-unit ↑ in $x$ → approx. $100\beta_1$ **percent** change in $y$ |
| $y = \beta_0 + \beta_1\log(x)$ | 1-**percent** ↑ in $x$ → approx. $\beta_1/100$ unit change in $y$ |
| $\log(y) = \beta_0 + \beta_1\log(x)$ | 1-percent ↑ in $x$ → $\beta_1$ **percent** change in $y$ (elasticity) |

> ⚠️ **WS 23/24 exam, Block III(iv):** *"In a linear regression model applying logarithmic transformation, the coefficients should be interpreted as the percentage change in the response for a 1% change in the predictor."* → **FALSE**, because that's only true in the **log-log** model. The statement doesn't say which variable was logged. See `32-TRAPS.md`.

---

## 1.2.2 Graphical association analysis

Now look at variables **in pairs**.

### The scatter plot

Plot $y$ (vertical) against $x$ (horizontal), one point per observation. You are asking four questions:

1. **Is there a trend?** Up, down, or none.
2. **Is the trend straight?** If it curves, you need $x^2$ or $\log(x)$ — Section 3.1.3.
3. **Is the spread constant?** If the cloud fans out, you have heteroscedasticity.
4. **Are there outliers?** Points far from the pattern.

> The exercise sheets do exactly this. Sheet 2 gives you a scatter plot of `wage` against `age` and says *"the effect of age seems to be rather quadratic than linear"* — then asks you to fit a second-order polynomial. That decision came from a picture, not a test.

### Correlation, and its limits

**Covariance** measures how two variables move together:

$$\widehat{\text{Cov}}(x,y) = \frac{1}{n-1}\sum_{i=1}^{n}(x_i - \bar{x})(y_i - \bar{y})$$

Problem: its size depends on the units. Covariance in €·m² is uninterpretable.

**Pearson correlation** fixes this by standardising:

$$r_{xy} = \frac{\widehat{\text{Cov}}(x,y)}{s_x \, s_y} = \frac{\sum (x_i - \bar{x})(y_i - \bar{y})}{\sqrt{\sum (x_i-\bar{x})^2}\sqrt{\sum(y_i-\bar{y})^2}} \in [-1, 1]$$

| $r$ | Meaning |
|---|---|
| $+1$ | perfect positive **linear** relationship |
| $0$ | no **linear** relationship |
| $-1$ | perfect negative **linear** relationship |

### The three things correlation cannot do — all examinable

**(1) It only sees straight lines.**
Data on a perfect parabola ($y = x^2$, $x$ symmetric about 0) has $r = 0$. There is a perfect deterministic relationship and correlation reports nothing. *Always plot.*

**(2) It doesn't imply causation.**
Ice cream sales correlate with drownings. Temperature causes both. The technical name is a **confounder**, and the entire point of *multiple* regression is to control for confounders by including them as covariates.

**(3) It's not the whole story about strength.**
Anscombe's quartet: four datasets with identical means, variances, correlations and regression lines, but completely different pictures. One is linear, one is curved, one has an outlier dragging the line, one is a single influential point. This is the single best argument for looking at plots, and it's the intellectual ancestor of **leverage and Cook's distance** in Section 3.4.4.

### The link you must remember: $r^2 = R^2$

> **In a simple linear regression with one covariate and an intercept, the squared Pearson correlation between $x$ and $y$ equals the coefficient of determination $R^2$.**
>
> $$R^2 = r_{xy}^2$$

**This is directly examined.** Exercise Sheet 3, Exercise 1:

> *"The estimated model resulted in an $R^2$ of 0.038. What is the correlation between wage and age?"*

Answer: $r = \sqrt{0.038} = 0.195$. And the **sign** comes from the sign of $\hat\beta_1$, which was $+0.71$ — so $r = +0.195$.

The interpretation: age explains only 3.8% of the variation in wage. That's tiny — because wage depends far more on education, job class and health, and because (from Sheet 2) the age effect isn't even linear. A weak $R^2$ is not proof of no relationship; it's an invitation to look harder.

### For a categorical $x$ and a continuous $y$

Use **side-by-side boxplots** — one box per category. If the boxes sit at clearly different heights, that category matters. This is the visual version of the dummy-variable coefficients you'll interpret in 3.1.3.

---

## The philosophy of this section in one line

> **The model does not tell you what to look for. The picture tells you what model to build.**

Chapter 3.4.4 (Model Diagnosis) is this same section applied to residuals instead of raw data. Same four questions — trend, curvature, constant spread, outliers — asked of $\hat\varepsilon$ rather than $y$. If you understand 1.2, you already understand 3.4.4.

---

## Key takeaways

1. Look at each variable alone (histogram, boxplot), then in pairs (scatter plot, correlation).
2. **Right-skewed responses → consider $\log(y)$.** Know all four log-interpretation cases.
3. Correlation measures **linear** association only, is scale-free, and lives in $[-1,1]$.
4. **Correlation ≠ causation**; confounders are the reason multiple regression exists.
5. **$R^2 = r^2$ in simple linear regression** — the sign comes from $\hat\beta_1$.
6. Anscombe's quartet: identical statistics, different pictures. Always plot.
