# Ch 2 — SOLUTIONS

---

## Part A — TRUE / FALSE

**A1. FALSE.** 🔴 *(WS 23/24, Block I(i))* $\hat\beta_j$ is a **partial** effect (holding other covariates fixed); the correlation is a **marginal** association. With confounding they can have opposite signs — the firefighters/fire-size example. *(In **simple** regression with one covariate, sign$(\hat\beta_1)=$ sign$(r)$ — but the statement doesn't say simple.)*

**A2. FALSE.** 🔴 *(WS 23/24, Block I(iv))* It requires $m-1$ dummies. Including all $m$ plus an intercept makes $\boldsymbol{X}$ rank-deficient.

**A3. FALSE.** 🔴 *(Exam Summer 2025, Ex 1(h))* $\hat\beta_j$ is the change in **log-odds**. The effect on the probability is $\hat\beta_j\pi(1-\pi)$, which is not constant.

**A4. TRUE.** $\exp(\hat\beta_j)$ is the **odds ratio**: odds are multiplied by $\exp(\hat\beta_j)$ per one-unit increase in $x_j$, holding others fixed.

**A5. FALSE.** With an interaction, $\partial E(y)/\partial x_1 = \beta_1+\beta_3x_2$. $\hat\beta_1$ is the effect of $x_1$ **only when $x_2 = 0$** (i.e. in the reference group).

**A6. TRUE.** Without an interaction the dummy shifts the intercept only ⟹ parallel lines. The interaction gives each group its own slope ⟹ non-parallel.

**A7. TRUE.** $\bar{\hat y}=\bar y$ follows from $\sum\hat\varepsilon_i = 0$, which holds whenever the model contains an intercept. *(RCLM WS22/23, Block I(ii).)*

**A8. TRUE.** This is the fatal objection to the linear probability model: $\boldsymbol{x}'\boldsymbol\beta$ is unbounded.

**A9. FALSE.** 🔴 *(RCLM WS22/23, Block III(iii))* Interactions can also be formed between **two categorical** covariates (products of dummies). No restriction on types.

**A10. FALSE.** The effect is $\hat\beta_1 + 2\hat\beta_2x$ — it depends on $x$. Interpreting $\hat\beta_1$ alone is wrong.

---

## Part B — Model building

### B1 (3 pts)

**(a)** `degree programme` has 4 levels ⟹ **3 dummies**. `tutorial` has 2 levels ⟹ **1 dummy**.

$$D^{\text{DS}}_i=\mathbb{1}\{\text{Data Science}\},\quad D^{\text{Ma}}_i=\mathbb{1}\{\text{Mathematics}\},\quad D^{\text{Ec}}_i=\mathbb{1}\{\text{Economics}\}$$
$$D^{\text{Tut}}_i = \mathbb{1}\{\text{attended the tutorial}\}$$

$$\text{score}_i = \beta_0+\beta_1\text{hours}_i+\beta_2D^{\text{DS}}_i+\beta_3D^{\text{Ma}}_i+\beta_4D^{\text{Ec}}_i+\beta_5D^{\text{Tut}}_i+\varepsilon_i$$

with $\varepsilon_i$ iid, $E(\varepsilon_i)=0$, $\text{Var}(\varepsilon_i)=\sigma^2$.

**(b)** Reference: a **Statistics student who did not attend the tutorial**.

**(c)** $p = 6$ parameters ($\beta_0,\dots,\beta_5$). Residual df $= n-p = 240-6 = \boxed{234}$.

### B2 (3 pts)

**(a)** `BMI category` has 4 levels ⟹ 3 dummies (reference: *normal*). `smoker` ⟹ 1 dummy $S_i$.

$$\text{cost}_i = \beta_0+\beta_1\text{age}_i+\beta_2S_i+\beta_3D^{\text{under}}_i+\beta_4D^{\text{over}}_i+\beta_5D^{\text{obese}}_i+\beta_6(\text{age}_i\times S_i)+\varepsilon_i$$

**(b)**

| | Intercept | Age slope |
|---|---|---|
| Non-smoker ($S=0$) | $\beta_0$ | $\beta_1$ |
| Smoker ($S=1$) | $\beta_0+\beta_2$ | $\beta_1+\beta_6$ |

*(Both hold at the reference BMI category.)*

**(c)** $\beta_6$ is the **difference in the age slope** between smokers and non-smokers: the additional increase in expected annual cost per extra year of age for a smoker relative to a non-smoker. The insurer's hypothesis is $\beta_6 > 0$.

### B3 (2 pts)

Let the $c$ dummies be $D_1,\dots,D_c$. Each observation belongs to exactly one category, so for every $i$:

$$D_{i1}+D_{i2}+\dots+D_{ic} = 1$$

The right-hand side is exactly the **intercept column** of $\boldsymbol{X}$. Hence the intercept column is a **linear combination** of the dummy columns: the columns of $\boldsymbol{X}$ are **linearly dependent**.

Consequently $\text{rank}(\boldsymbol{X}) < p$, so $\text{rank}(\boldsymbol{X}'\boldsymbol{X}) < p$ and $\boldsymbol{X}'\boldsymbol{X}$ is **singular** — its inverse does not exist. The normal equations $\boldsymbol{X}'\boldsymbol{X}\hat{\boldsymbol\beta} = \boldsymbol{X}'\boldsymbol{y}$ then have **infinitely many solutions**, so $\hat{\boldsymbol\beta} = (\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$ is undefined and the individual coefficients are **not identified**.

*(This is the **dummy variable trap** / perfect multicollinearity, and it is why the classical linear model assumes $\text{rank}(\boldsymbol{X}) = p$.)*

---

## Part C — Interpretation

### C1 (4 pts)

**(a)** Four education dummies appear ⟹ education has **5 levels**. Reference: the level not shown, **`< HS Grad`**.

**(b)** *Holding education and health fixed, a one-year increase in age is associated with an estimated increase of **\$0.62** in expected hourly wage.*

**(c)** *Holding age and education fixed, men in very good or better health earn on average an estimated **\$9.13** more per hour than men in good or worse health (the reference).*

**(d)**
$$\widehat{\text{wage}} = 52.61 + 0.62(37) + 37.97 + 9.13 = 52.61+22.94+37.97+9.13 = \boxed{\$122.65}$$

**(e)**
$$\hat\varepsilon = 121.43 - 122.65 = \boxed{-\$1.22}$$
*John earns \$1.22 per hour **less** than the model predicts for a man of his age, education and health. The residual is very small relative to the wage scale, so the model fits him well.*

**(f)** Same age and health ⟹ those terms cancel. Subtract the two dummy coefficients:
$$62.63 - 11.01 = \boxed{\$51.62}$$

### C2 (3 pts)

**(a)**

| Group | Intercept | Slope |
|---|---|---|
| health ≤ Good ($H=0$) | $78.66$ | $0.51$ |
| health ≥ Very Good ($H=1$) | $78.66-1.81 = 76.85$ | $0.51+0.43 = 0.94$ |

**(b)** Set the two lines equal:
$$78.66+0.51a = 76.85+0.94a \;\Longrightarrow\; 1.81 = 0.43a \;\Longrightarrow\; a = \boxed{4.21 \text{ years}}$$

*Interpretation:* the crossing point lies far **outside the data range** (the sample contains men aged roughly 18–80). So within the observed range the healthier group always has the higher fitted wage; the negative intercept difference is an artefact of extrapolating both lines back to age 0 and should not be interpreted.

**(c)** **No.** With the interaction present, $-1.81$ is the health effect **only at age 0**, which is meaningless here. The health effect is
$$\frac{\partial\widehat{\text{wage}}}{\partial H} = -1.81 + 0.43\cdot\text{age}$$
At age 40 it is $-1.81+17.2 = +\$15.39$; at age 60, $+\$24.99$. The correct statement is that **the health effect increases with age**, and is positive at every age above ≈ 4.2.

### C3 (2 pts)

**(a)** $\dfrac{\partial\widehat{\text{wage}}}{\partial\text{age}} = \hat\beta_1+2\hat\beta_2\text{age} = 5.29 - 0.10\cdot\text{age}$

| Age | Marginal effect |
|---|---|
| 30 | $5.29-3.00 = +\$2.29$ per year |
| 65 | $5.29-6.50 = -\$1.21$ per year |

At 30 an extra year of age raises expected wage; at 65 it lowers it.

**(b)** Set the derivative to zero:
$$5.29-0.10\,\text{age} = 0 \;\Longrightarrow\; \text{age}^* = \frac{5.29}{0.10} = \boxed{52.9 \text{ years}}$$

Since $\hat\beta_2 = -0.05 < 0$ the parabola opens downward, so this is a **maximum**. ✓ (Consistent with Sheet 2, Figure 2.)

---

## Part D — Logit

### D1 (3 pts)

> For a binary $y_i\in\{0,1\}$ we have $E(y_i\mid\boldsymbol{x}_i)=P(y_i=1\mid\boldsymbol{x}_i)=\pi_i$, so the regression function is a **probability** and must lie in $[0,1]$.
>
> **(1) Fitted values are not valid probabilities.** The linear predictor $\boldsymbol{x}_i'\boldsymbol\beta$ is unbounded, so for extreme covariate values the linear model produces fitted probabilities below 0 or above 1, which are meaningless.
>
> **(2) Heteroscedasticity is unavoidable.** For a Bernoulli response, $\text{Var}(y_i)=\pi_i(1-\pi_i)$, which depends on $\boldsymbol{x}_i$. The homoscedasticity assumption is violated **by construction**, so OLS is no longer BLUE and the usual standard errors are invalid.
>
> **(3) Errors cannot be normally distributed.** Given $\boldsymbol{x}_i$, $\varepsilon_i = y_i-\boldsymbol{x}_i'\boldsymbol\beta$ takes only **two** values, so the exact $t$- and $F$-tests, which rest on normality, have no justification.
>
> **(4) Constant marginal effects are implausible.** A linear model imposes the same change in probability per unit of $x_j$ everywhere, but effects must attenuate near 0 and 1.
>
> The **logit** (or **probit**) model instead specifies $P(y_i=1)=h(\boldsymbol{x}_i'\boldsymbol\beta)$ with $h$ strictly increasing from $\mathbb{R}$ onto $(0,1)$ — the logistic CDF for logit, $\Phi$ for probit. Fitted values are then always valid probabilities, marginal effects automatically shrink near the boundaries, and the Bernoulli likelihood is modelled correctly. Estimation is by **maximum likelihood** rather than OLS.

### D2 (3 pts)

**(a)** $\exp(0.028) = 1.028$. Each additional **month** of loan duration **multiplies the odds of default by 1.028**, i.e. increases the odds by about **2.8%**, holding other covariates fixed.

**(b)** $\exp(-0.85) = 0.427$. Home owners have odds of default about **0.43 times** those of non-owners — a reduction of roughly **57%** in the odds — holding other covariates fixed.

**(c)** The statement confuses the **log-odds** scale with the **probability** scale.
> $\hat\beta = -0.85$ is a change in the **log-odds**, not in the probability. Correctly: owning a home **multiplies the odds** of default by $\exp(-0.85)=0.427$. The change in the *probability* is $\hat\beta\,\pi(1-\pi)$, which depends on the customer's current $\pi$ and is therefore different for every customer — no single number describes it. Only the **sign** (home ownership lowers default risk) transfers directly.

**(d)** At $\hat\pi = 0.5$: $\;\hat\beta\,\pi(1-\pi) = 0.028\times0.5\times0.5 = 0.028\times0.25 = \boxed{0.007}$

One extra month raises the default probability by about **0.7 percentage points** — and this is the **largest** the marginal effect can be, since $\pi(1-\pi)$ is maximised at $\pi=0.5$.

### D3 (2 pts)

$$\text{(probability)}\quad \pi_i = \frac{\exp(\boldsymbol{x}_i'\boldsymbol\beta)}{1+\exp(\boldsymbol{x}_i'\boldsymbol\beta)}$$
$$\text{(odds)}\quad \frac{\pi_i}{1-\pi_i} = \exp(\boldsymbol{x}_i'\boldsymbol\beta)$$
$$\text{(log-odds)}\quad \log\!\left(\frac{\pi_i}{1-\pi_i}\right) = \boldsymbol{x}_i'\boldsymbol\beta$$

The **log-odds** form makes it a *generalised linear model*: after applying the **link function** $g(\pi)=\log\frac{\pi}{1-\pi}$, the transformed mean is a **linear** function of the covariates. The model is linear — not in $\pi$, but in $g(\pi)$. The linear predictor $\eta=\boldsymbol{x}'\boldsymbol\beta$ is the shared skeleton; the link is what adapts it to a bounded response.

---

## Part E

### E1 (3 pts)

Define:

$$D^{\text{HS}}_i=\begin{cases}1&\text{person }i\text{ has a high school degree}\\0&\text{otherwise}\end{cases}
\qquad
D^{\text{Col}}_i=\begin{cases}1&\text{person }i\text{ has a college/university degree}\\0&\text{otherwise}\end{cases}$$

$$D^{\text{out}}_i=\begin{cases}1&\text{person }i\text{ was born outside the US}\\0&\text{person }i\text{ was born inside the US}\end{cases}$$

Model:

$$\text{wage}_i = \beta_0+\beta_1\text{age}_i+\beta_2D^{\text{HS}}_i+\beta_3D^{\text{Col}}_i+\beta_4D^{\text{out}}_i+\varepsilon_i,\qquad i=1,\dots,n$$

with $\varepsilon_i$ iid, $E(\varepsilon_i)=0$, $\text{Var}(\varepsilon_i)=\sigma^2$.

**Reference category:** a person with **no degree, born inside the US**. Education (3 levels) contributes **2** dummies and place of birth (2 levels) contributes **1**, so $\boldsymbol{X}$ has full column rank $p=5$ and the model is estimable by OLS.

### E2 (2 pts)

**(a)** The effect of age on wage is unlikely to be **linear**. Earnings typically rise steeply in early career, flatten in mid career, and decline near retirement. Adding a quadratic term lets the model capture this **curvature** — the fitted age–wage profile becomes a parabola rather than a straight line — while the model remains **linear in the parameters** and can still be estimated by OLS.

**(b)** **Negative.** A negative coefficient makes the parabola open **downward**, producing a hump-shaped age–wage profile with a maximum near age 48, which matches the empirical pattern (and the exercise-sheet scatter plot). A positive coefficient would imply wage is *minimised* around 48 and rises without bound at both young and old ages, which is implausible.

**(c)** Two reasons:

1. **Interpretability.** With $(\text{age}-48)^2$, the term vanishes at age 48, so the remaining coefficients describe a person at age 48 — a meaningful, in-sample reference point. With raw $\text{age}^2$ the coefficients refer to age 0, which is meaningless (see the intercept discussion in 2.2.1).
2. **Reduced multicollinearity.** Over a positive range like 18–80, $\text{age}$ and $\text{age}^2$ are very strongly correlated, which inflates the standard errors of both coefficients. **Centring** at a value near the middle of the range makes $(\text{age}-48)$ and $(\text{age}-48)^2$ close to uncorrelated, giving more stable and precisely estimated coefficients.

---

## Marking guide

| Section | Points |
|---|---|
| A (T/F) | 10 |
| B (model building) | 8 |
| C (interpretation) | 9 |
| D (logit) | 8 |
| E (exam-realistic) | 5 |
| **Total** | **40** |

| Score | Verdict |
|---|---|
| 34–40 | Chapter 2 done. Go to Chapter 3 today. |
| 26–33 | Redo Part B and Part D. Re-read `03-notes-2.2.2` and `04-notes-2.3`. |
| < 26 | Re-read all four notes files. The dummy rule and the logit interpretation are non-negotiable — they are guaranteed marks and you are currently leaving them on the table. |
