# 2.2.1 — The Simple Linear Regression Model

> **Purpose:** one covariate, one slope, one intercept. Everything here generalises, so get the interpretation *exactly* right now — it's the same sentence forever after.

---

## 1. The model

$$\boxed{\;y_i = \beta_0 + \beta_1 x_i + \varepsilon_i, \qquad i = 1,\dots,n\;}$$

with $E(\varepsilon_i) = 0$, so that

$$E(y_i \mid x_i) = \beta_0 + \beta_1 x_i$$

| Symbol | Name | Meaning |
|---|---|---|
| $\beta_0$ | intercept | expected $y$ when $x = 0$ |
| $\beta_1$ | slope | change in expected $y$ per one-unit increase in $x$ |
| $\varepsilon_i$ | error | everything else |

**In matrix form** (worth 1 point in the *Linear_model_exam_sheet* paper):

$$
\begin{pmatrix} y_1 \\ y_2 \\ \vdots \\ y_n\end{pmatrix} =
\begin{pmatrix} 1 & x_1 \\ 1 & x_2 \\ \vdots & \vdots \\ 1 & x_n \end{pmatrix}
\begin{pmatrix}\beta_0\\\beta_1\end{pmatrix} +
\begin{pmatrix}\varepsilon_1\\\varepsilon_2\\\vdots\\\varepsilon_n\end{pmatrix}
$$

$\boldsymbol{X}$ is $n\times 2$, $\boldsymbol\beta$ is $2\times 1$. Here $k=1$, $p=2$, residual df $= n-2$.

---

## 2. The least squares estimates (closed form)

For the *simple* model there are explicit formulas. Know them — they occasionally let you reconstruct a missing R output without touching a matrix.

$$\boxed{\;\hat\beta_1 = \frac{\sum_{i}(x_i-\bar x)(y_i - \bar y)}{\sum_i (x_i - \bar x)^2} = \frac{\widehat{\text{Cov}}(x,y)}{\widehat{\text{Var}}(x)} = r_{xy}\cdot\frac{s_y}{s_x}\;}$$

$$\boxed{\;\hat\beta_0 = \bar y - \hat\beta_1\bar x\;}$$

### Three consequences you should be able to state instantly

**(a) The fitted line passes through $(\bar x, \bar y)$.**
Rearranging $\hat\beta_0 = \bar y - \hat\beta_1\bar x$ gives $\bar y = \hat\beta_0 + \hat\beta_1\bar x$. The centre of gravity of the data is always on the line.

> **This is a free mark generator.** *Example Exam LiMo 2020* gives you: mean goals = 48.61, mean points = 46.61, and $\hat\beta_1 = 0.90509$, then asks for the missing intercept **A**.
> $$\hat\beta_0 = 46.61 - 0.90509 \times 48.61 = 46.61 - 43.997 = 2.613$$
> Done in fifteen seconds, no data required.

**(b) The sign of $\hat\beta_1$ equals the sign of $r_{xy}$** (in *simple* regression only — since $s_y/s_x > 0$).

**(c) $R^2 = r_{xy}^2$** (again, simple regression with intercept only).

---

## 3. 🎯 Interpreting the slope — the sentence you will write twenty times

$$\hat\beta_1 = \text{estimated change in the } \textbf{expected } y \text{ for a } \textbf{one-unit} \text{ increase in } x$$

### The template

> **"A one-[unit of $x$] increase in [$x$] is associated with an estimated [$\hat\beta_1$] [unit of $y$] change in the expected [$y$]."**

### Worked from Sheet 1

Model: $\text{wage}_i = \beta_0 + \beta_1\cdot\text{age}_i + \varepsilon_i$, with $\hat\beta_0 = 81.70$, $\hat\beta_1 = 0.71$.

> **"A one-year increase in age is associated with an estimated increase of \$0.71 in expected hourly wage."**

Check the four things every marker looks for:

| Element | Present? |
|---|---|
| The **unit** of $x$ ("one **year**") | ✅ |
| The **unit** of $y$ ("**\$**0.71 per hour") | ✅ |
| The word **"expected"** or "on average" | ✅ |
| **"Associated with"**, not "causes" | ✅ |

Missing any of these loses fractions of marks that add up.

**In multiple regression add a fifth element:** *"holding all other covariates fixed."* Not needed here (there's only one covariate) but never wrong to say.

---

## 4. ⚠️ Interpreting the intercept — and when not to

$$\hat\beta_0 = \text{estimated expected } y \text{ when } x = 0$$

For Sheet 1: $\hat\beta_0 = 81.70$ would be *"the expected hourly wage of a man of age 0."*

**Sheet 1 Exercise 1(c) asks exactly this: "Why should $\hat\beta_0$ not be interpreted here?"**

Model answer:

> Because $x = 0$ (age zero) lies **far outside the observed range** of the data — the `Wage` dataset contains men aged roughly 18 to 80 — and is **substantively meaningless**: a newborn has no hourly wage. Interpreting $\hat\beta_0$ would be **extrapolation** beyond the data, where we have no evidence the linear relationship holds. The intercept is still necessary to position the regression line correctly, but it carries no substantive meaning.

### The rule

Interpret $\hat\beta_0$ **only if** both:
1. $x = 0$ lies inside (or very near) the observed range of $x$, **and**
2. $x=0$ is a meaningful state of the world.

### The fix: centring

Replace $x$ by $(x - \bar x)$ or $(x - c)$ for a meaningful constant $c$:

$$y_i = \beta_0^* + \beta_1(x_i - 48) + \varepsilon_i$$

Now $\beta_0^*$ = expected wage at **age 48**, which is meaningful and inside the data. $\beta_1$ is unchanged.

> **This is exactly why Exam Summer 2025 writes $(\text{age}_i - 48)^2$ instead of $\text{age}_i^2$.** Centring makes coefficients interpretable and reduces the correlation between the linear and quadratic terms (less multicollinearity). If you're ever asked "why subtract 48?", that's the answer.

---

## 5. Fitted values and residuals

$$\hat y_i = \hat\beta_0 + \hat\beta_1 x_i \qquad\qquad \hat\varepsilon_i = y_i - \hat y_i$$

**Interpretation of a residual** — worth a mark in Sheet 3, Ex 2(d):

- $\hat\varepsilon_i > 0$: this person earns **more** than the model predicts for someone with their characteristics
- $\hat\varepsilon_i < 0$: this person earns **less** than predicted
- $|\hat\varepsilon_i|$ large: the model fits this person poorly

> Sheet 3 Ex 2(d): John's predicted wage is \$122.65, actual \$121.43, so $\hat\varepsilon = -1.22$.
> **"John earns \$1.22 per hour less than the model predicts for a man with his age, education and health — the model fits him very well."**

### Two algebraic properties (true whenever the model contains an intercept)

$$\sum_i \hat\varepsilon_i = 0 \qquad\text{and}\qquad \sum_i x_i\hat\varepsilon_i = 0$$

These are the **normal equations** — they fall directly out of setting the derivatives of the squared-error sum to zero. They say: the residuals have mean zero, and are **uncorrelated with the covariate**. Geometrically, the residual vector is orthogonal to every column of $\boldsymbol{X}$.

**Consequences worth knowing:**
- $\bar{\hat y} = \bar y$ — the average fitted value equals the average observed response.
  > *RCLM WS 22/23, Block I(ii): "The average of the predicted values is equal to the average of the observed response."* → **TRUE**, provided the model has an intercept.
- The mean of the residuals being ≈ 0 is **not evidence the model is good** — it is guaranteed by construction.
  > *WS 23/24, Block II(iv): "In a well-fitted model, the mean of the residuals should be close to zero."* → marked **TRUE** on that key, but understand *why*: it's automatic, not diagnostic.

---

## 6. Non-linear relationships in a "linear" model

Restating Chapter 1's point in the setting where you'll use it.

### Polynomial regression — Sheet 2

Sheet 2 shows the wage-vs-age scatter and says the effect "seems rather quadratic than linear." The model:

$$\text{wage}_i = \beta_0 + \beta_1\text{age}_i + \beta_2\text{age}_i^2 + \varepsilon_i$$

Still a **linear model** (linear in $\beta$). Just set $x_1 = \text{age}$, $x_2 = \text{age}^2$ and run ordinary OLS.

Sheet 2's estimates: $\hat\beta_0 = -10.43$, $\hat\beta_1 = 5.29$, $\hat\beta_2 = -0.05$.

**Interpretation (this is the marked part):**

> $\hat\beta_2 = -0.05 < 0$, so the fitted curve is a **downward-opening parabola**: wage increases with age at first, reaches a maximum, then declines. This matches the economic story — earnings rise through early and mid career, peak, then fall as people move to part-time work or lower-paid late-career roles.

**Where is the peak?** Differentiate and set to zero:

$$\frac{\partial \hat{\text{wage}}}{\partial\text{age}} = \hat\beta_1 + 2\hat\beta_2\,\text{age} = 0 \;\Longrightarrow\; \text{age}^* = -\frac{\hat\beta_1}{2\hat\beta_2}$$

$$\text{age}^* = -\frac{5.29}{2(-0.05)} = \frac{5.29}{0.10} = 52.9 \text{ years}$$

Which is right where the curve in Sheet 2's Figure 2 turns over. ✓

> 🔑 **The crucial interpretation warning.** In $y = \beta_0+\beta_1x+\beta_2x^2$, you **cannot interpret $\hat\beta_1$ alone** as "the effect of $x$." The effect of $x$ now *depends on the value of $x$*:
> $$\frac{\partial E(y)}{\partial x} = \beta_1 + 2\beta_2 x$$
> Saying "a one-year increase in age raises wage by \$5.29" is **wrong** and will be marked wrong. The effect at age 30 is $5.29 + 2(-0.05)(30) = \$2.29$; at age 60 it is $5.29 - 6.00 = -\$0.71$.
>
> **Rule: when a variable appears more than once in the model (as $x$ and $x^2$, or in an interaction), its coefficients must be interpreted together.**

### Other linearisable forms

| Relationship | Linear form | Note |
|---|---|---|
| $y = \beta_0 + \beta_1\log(x) + \varepsilon$ | already linear | diminishing returns |
| $y = \exp(\beta_0+\beta_1x+\varepsilon)$ | $\log y = \beta_0+\beta_1x+\varepsilon$ | error **inside** — log works |
| $y = \exp(\beta_0+\beta_1x)+\varepsilon$ | **not** linearisable | error **outside** — needs NLS |

---

## 7. Key takeaways

1. $y_i = \beta_0+\beta_1x_i+\varepsilon_i$; $\boldsymbol{X}$ is $n\times2$; residual df $= n-2$.
2. $\hat\beta_1 = \widehat{\text{Cov}}(x,y)/\widehat{\text{Var}}(x)$ and $\hat\beta_0 = \bar y - \hat\beta_1\bar x$.
3. **The line always passes through $(\bar x,\bar y)$** — use it to recover missing intercepts.
4. **Slope sentence:** unit of $x$ · unit of $y$ · "expected" · "associated with".
5. **Don't interpret $\hat\beta_0$** unless $x=0$ is meaningful and in range. Centring fixes this.
6. With an intercept: $\sum\hat\varepsilon_i = 0$, $\sum x_i\hat\varepsilon_i = 0$, $\bar{\hat y} = \bar y$.
7. **Polynomial terms are still a linear model** — but their coefficients must be interpreted **jointly**, never alone. Turning point at $-\hat\beta_1/(2\hat\beta_2)$.
