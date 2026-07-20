# 2.2.2 — Multiple Linear Regression (and Dummy Variables)

> **This is the most examined section in Chapter 2.** Dummy variables appear in every past paper. Give it the time.

---

## 1. The model

$$\boxed{\;y_i = \beta_0 + \beta_1x_{i1} + \beta_2x_{i2} + \dots + \beta_kx_{ik} + \varepsilon_i = \boldsymbol{x}_i'\boldsymbol\beta + \varepsilon_i\;}$$

In matrix form: $\boldsymbol{y} = \boldsymbol{X}\boldsymbol\beta + \boldsymbol\varepsilon$, with $\boldsymbol{X}$ of dimension $n\times p$, $p = k+1$.

---

## 2. 🔑 What changes about interpretation

$$\hat\beta_j = \text{estimated change in expected } y \text{ for a one-unit increase in } x_j, \textbf{ holding all other covariates fixed}$$

That bolded phrase is the entire difference from simple regression, and it is worth marks every time you write it. Equivalent phrasings all acceptable:

- "…**holding all other covariates constant**"
- "…**ceteris paribus**"
- "…**controlling for the other variables**"
- "…**for two individuals who differ by one unit in $x_j$ but are otherwise identical**"

### Why "partial" effects differ from "marginal" effects

$\hat\beta_j$ is a **partial** effect: it compares people who are identical except in $x_j$.
The simple correlation between $x_j$ and $y$ is a **marginal** effect: it compares people who differ in $x_j$ *and in everything correlated with it*.

**These can have opposite signs.**

> 🔴 **WS 23/24, Block I(i):** *"In linear regression, if the coefficient of a variable is positive, then there must be a positive correlation between that variable and the response."* → **FALSE.**
>
> Firefighters and property damage: marginal correlation strongly positive, partial effect (controlling for fire size) negative. This is the confounding story from Chapter 1, now with a formula attached.

---

## 3. Dummy variables — the core skill

### The problem

Covariates like `education` (5 levels) or `place of birth` (2 levels) are **categorical**. You cannot put "College Grad" into an equation. And you must not code them as 1,2,3,4,5 — that would force the model to assume the gap from level 1→2 equals the gap from 4→5, and that the levels are *equally spaced*, which is nonsense.

### The solution

> **A categorical covariate with $c$ levels is represented by $c-1$ binary dummy variables. The omitted level is the reference (baseline) category.**

Each dummy is defined:

$$D_j = \begin{cases} 1 & \text{if the observation is in category } j\\ 0 & \text{otherwise}\end{cases}$$

### ⚠️ Why $c-1$ and not $c$ — the dummy variable trap

If you include all $c$ dummies **and** an intercept, then for every observation

$$D_1 + D_2 + \dots + D_c = 1 = \text{the intercept column}$$

The columns of $\boldsymbol{X}$ are **linearly dependent**. Therefore:

- $\text{rank}(\boldsymbol{X}) < p$ — not full column rank
- $\boldsymbol{X}'\boldsymbol{X}$ is **singular** — $(\boldsymbol{X}'\boldsymbol{X})^{-1}$ does not exist
- **OLS has no unique solution**

This is called the **dummy variable trap** (or perfect multicollinearity), and it is the concrete reason Chapter 3 assumes $\text{rank}(\boldsymbol{X}) = p$.

> 🔴 **Exam Summer 2025, Ex 1(d):** *"When the design matrix $X$ does not have full column rank, the OLS estimates still exist and are unique as long as the error variance is constant."* → **FALSE.** Uniqueness requires full rank. Homoscedasticity is irrelevant to this. (Estimates in the sense of *some* solution exist via generalised inverse, but they are **not unique**, and $\hat{\boldsymbol\beta}$ is not identified.)

**Past-paper T/F on the count:**
- *Linear_model_exam_sheet*, I(iv): "with $k$ levels, $k-1$ dummies are created" → **TRUE**
- *WS 23/24*, I(iv): "for $m$ categories we need $m$ dummy variables" → **FALSE**

---

## 4. 📝 Worked example — Sheet 1, Exercise 2

**Setup.** Model wage on age and education. Education has **5 levels**:

1. `< HS Grad`
2. `HS Grad`
3. `Some College`
4. `College Grad`
5. `Advanced Degree`

**Step 1 — define the dummies.** $5$ levels ⟹ $4$ dummies. Take level 1 as reference:

$$D_2 = \mathbb{1}\{\text{HS Grad}\},\quad D_3 = \mathbb{1}\{\text{Some College}\},\quad D_4 = \mathbb{1}\{\text{College Grad}\},\quad D_5 = \mathbb{1}\{\text{Advanced Degree}\}$$

**Step 2 — write the model.**

$$\text{wage}_i = \beta_0 + \beta_1\text{age}_i + \beta_2 D_{2i} + \beta_3D_{3i} + \beta_4D_{4i} + \beta_5D_{5i} + \varepsilon_i$$

**Step 3 — the fitted output (Sheet 1).**

| Term | Estimate |
|---|---|
| (Intercept) | 60.336 |
| age | 0.569 |
| education2. HS Grad | 11.439 |
| education3. Some College | 24.167 |
| education4. College Grad | 39.767 |
| education5. Advanced Degree | 64.987 |

### (b) Which is the reference category?

> **`< HS Grad` (level 1)** — it is the level with no dummy in the output. R omits the first level alphabetically/by factor order by default.

**How to spot the reference in any R output:** the level that *doesn't appear* is the reference. If `education` has 5 levels and only 4 appear, the missing one is baseline.

### (b) Interpret the education effects

> Holding age fixed, men with a **high school degree** earn on average an estimated **\$11.44** more per hour than men with **less than a high school degree** (the reference category).
>
> The estimated premium relative to the same reference is **\$24.17** for *Some College*, **\$39.77** for *College Grad*, and **\$64.99** for *Advanced Degree*.
>
> The coefficients increase monotonically, indicating that expected wage rises steadily with education level.

**Three things to notice:**
1. Every dummy coefficient is a comparison **to the reference**, never to the level below it.
2. The **units** are the units of $y$ (dollars per hour).
3. "Holding age fixed" appears. Always.

### (c) Wage difference: 50-year-old with level 3 vs 50-year-old with level 5

Both are 50, so the age terms **cancel**. The difference is purely the two dummy coefficients:

$$\hat\beta_5 - \hat\beta_3 = 64.987 - 24.167 = \boxed{40.82}$$

> The man with an Advanced Degree is estimated to earn \$40.82 per hour more.

> 🔑 **The key move:** to compare two non-reference categories, **subtract their coefficients**. Do not use the raw coefficient of one of them — that compares to the reference, not to each other.

### (d) 40-year-old with level 4 vs 20-year-old with level 1

Now **both** age and education differ. Two options: build both predictions, or take the difference directly.

**Direct difference:**
$$\Delta = \hat\beta_1(40-20) + (\hat\beta_4 - 0) = 0.56869\times 20 + 39.76677 = 11.374 + 39.767 = \boxed{51.14}$$

(The 20-year-old is in the reference category, so his education contribution is 0.)

**Check by full prediction:**
- 40-yr, College Grad: $60.336 + 0.56869(40) + 39.767 = 60.336+22.748+39.767 = 122.85$
- 20-yr, < HS Grad: $60.336 + 0.56869(20) + 0 = 60.336 + 11.374 = 71.71$
- Difference: $122.85 - 71.71 = 51.14$ ✓

---

## 5. 📝 Worked example — Exam Summer 2025, Exercise 2(a) [3 points]

**Question.** Model hourly wage on **age**, **education** (no degree / high school / college), **place of birth** (inside US / outside US). Include an intercept. Formulate it so it can be estimated by OLS.

**Full-mark answer:**

Define dummy variables:

$$D^{\text{HS}}_i = \begin{cases}1 & \text{person } i \text{ has a high school degree}\\0&\text{else}\end{cases}\qquad
D^{\text{Col}}_i = \begin{cases}1&\text{person } i \text{ has a college/university degree}\\0&\text{else}\end{cases}$$

$$D^{\text{out}}_i = \begin{cases}1 & \text{person } i \text{ was born outside the US}\\0&\text{born inside the US}\end{cases}$$

Model:

$$\text{wage}_i = \beta_0 + \beta_1\,\text{age}_i + \beta_2 D^{\text{HS}}_i + \beta_3D^{\text{Col}}_i + \beta_4D^{\text{out}}_i + \varepsilon_i,\qquad i=1,\dots,n$$

with $\varepsilon_i$ iid, $E(\varepsilon_i)=0$, $\text{Var}(\varepsilon_i)=\sigma^2$.

**Reference category:** a person with **no degree, born inside the US**.

**Why this earns all three marks:**

| Requirement | Where it's satisfied |
|---|---|
| All three covariates present | age, education, birthplace ✓ |
| Intercept present | $\beta_0$ ✓ |
| Education (3 levels) → **2** dummies, not 3 | $D^{\text{HS}}, D^{\text{Col}}$ ✓ |
| Birthplace (2 levels) → **1** dummy | $D^{\text{out}}$ ✓ |
| Dummies **explicitly defined** | the braces ✓ |
| Reference category stated | ✓ |
| Estimable by OLS (full rank) | 5 parameters, no redundancy ✓ |
| Error term specified | ✓ |

> **The most common way students lose marks here:** writing "education" as a single variable without defining dummies, or defining 3 dummies for 3 levels. Count the levels, subtract one, define them explicitly with braces, and name the reference. Every time.

---

## 6. Interaction effects — Sheet 2, Exercise 2

Sometimes the *effect of one covariate depends on the value of another*. That is an **interaction**, coded as a **product term**.

### The model

Let $H_i = \mathbb{1}\{\text{health} \geq \text{Very Good}\}$:

$$\text{wage}_i = \beta_0 + \beta_1\text{age}_i + \beta_2 H_i + \beta_3(\text{age}_i \times H_i) + \varepsilon_i$$

### Reading it: split by group

**Group $H=0$ (health ≤ Good):**
$$\hat{\text{wage}} = \hat\beta_0 + \hat\beta_1\text{age}$$
→ intercept $\hat\beta_0$, slope $\hat\beta_1$

**Group $H=1$ (health ≥ Very Good):**
$$\hat{\text{wage}} = (\hat\beta_0+\hat\beta_2) + (\hat\beta_1+\hat\beta_3)\text{age}$$
→ intercept $\hat\beta_0+\hat\beta_2$, slope $\hat\beta_1+\hat\beta_3$

| Coefficient | Meaning |
|---|---|
| $\beta_2$ | difference in **intercept** between groups |
| $\beta_3$ | difference in **slope** between groups ← *this is the interaction* |

### Sheet 2(b) worked

$$\widehat{\text{wage}} = 78.66 + 0.51\cdot\text{age} - 1.81\cdot H + 0.43\cdot\text{age}\cdot H$$

| Group | Intercept | Slope |
|---|---|---|
| health ≤ Good ($H=0$) | $78.66$ | $0.51$ |
| health ≥ Very Good ($H=1$) | $78.66 - 1.81 = \boxed{76.85}$ | $0.51+0.43 = \boxed{0.94}$ |

**Interpretation:** healthier men start at a slightly *lower* fitted wage in youth, but their wage rises with age nearly **twice as fast** (\$0.94/yr vs \$0.51/yr). The lines **cross**.

### Sheet 2(c): "What does the interaction do from a geometric perspective?"

**Model answer:**

> **Without** an interaction, the dummy shifts the line up or down but the slope is shared — the two fitted lines are **parallel**. The health effect is a constant \$16.90 at every age.
>
> **With** an interaction, each group gets its own slope, so the lines are **no longer parallel** — they may converge, diverge, or cross. The effect of health on wage now **depends on age**, and equivalently the effect of age on wage depends on health.

> 🔑 **The one-line summary to memorise:**
> **Dummy alone = parallel lines (shift). Dummy + interaction = non-parallel lines (shift *and* tilt).**

### ⚠️ The interaction interpretation trap

With an interaction present, $\hat\beta_1 = 0.51$ is **not** "the effect of age." It is the effect of age **only in the reference group** ($H=0$). Likewise $\hat\beta_2 = -1.81$ is the health effect **only at age 0** — which here isn't even meaningful.

$$\frac{\partial E(\text{wage})}{\partial\text{age}} = \beta_1 + \beta_3 H \qquad \frac{\partial E(\text{wage})}{\partial H} = \beta_2 + \beta_3\text{age}$$

**Same rule as polynomials:** *when a variable appears in more than one term, interpret its coefficients jointly, never in isolation.*

**Hierarchy rule:** if you include an interaction $x_1x_2$, always include the **main effects** $x_1$ and $x_2$ too. Omitting a main effect forces an arbitrary constraint on the model.

---

## 7. Model building summary — the recipe

Given a word problem, produce the model in four steps:

1. **Identify $y$** and check its type.
2. **List covariates.** For each: continuous, binary, or categorical with $c$ levels?
3. **Code them:**
   - continuous → in as-is (consider $x^2$ or $\log x$ if the relationship curves)
   - binary → one dummy
   - categorical with $c$ levels → **$c-1$ dummies**, name the reference
4. **Write the equation** with an intercept, explicitly define every dummy, and add $+\varepsilon_i$ with its assumptions.

**Parameter count check:** $p = 1 + (\text{continuous terms}) + \sum_{\text{categorical}}(c-1) + (\text{interaction terms})$.

For the Exam 2025 wage model: $p = 1 + 1 + 2 + 1 = 5$, so residual df $= n-5$.

---

## 8. Key takeaways

1. $\hat\beta_j$ = partial effect, **holding all other covariates fixed**. Say the phrase.
2. Partial effect and marginal correlation can have **opposite signs**.
3. **$c$ levels ⟹ $c-1$ dummies.** All $c$ plus an intercept ⟹ singular $\boldsymbol{X}'\boldsymbol{X}$ ⟹ no unique OLS.
4. Every dummy coefficient compares to the **reference category**. To compare two non-reference levels, **subtract their coefficients**.
5. The reference category is the level **missing** from the R output.
6. **Interaction = product term = non-parallel lines.** $\beta_2$ shifts the intercept, $\beta_3$ shifts the slope.
7. **A variable appearing in multiple terms (polynomial or interaction) must have its coefficients interpreted jointly.**
