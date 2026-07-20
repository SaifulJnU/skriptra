# 3.1.3 — Modelling the Effects of Covariates

> **The section that answers Exercise 2 on the exam paper.** Chapter 2.2.2 introduced dummies and interactions; this section is the systematic treatment, plus transformations.
>
> If you've read `chapter-02.../03-notes-2.2.2` and `31-TRICKS-AND-TIPS.md`, much of §2–§4 will be revision. Read it anyway — the exam-mark density here is very high.

---

## 1. The master principle

> **The linear model is linear in $\boldsymbol\beta$. Anything you can compute from the covariates before fitting can be a column of $\boldsymbol{X}$.**

So the modelling question is never *"can the linear model handle this?"* It is **"what columns do I put in $\boldsymbol{X}$?"**

| Situation | Columns to add |
|---|---|
| continuous covariate, straight effect | $x$ |
| continuous covariate, curved effect | $x,\ x^2$ (maybe $x^3$) |
| diminishing returns / multiplicative effect | $\log(x)$ |
| skewed positive response | model $\log(y)$ instead |
| categorical covariate, $c$ levels | $c-1$ dummies |
| effect of one covariate depends on another | product term $x_1x_2$ |
| group-specific slopes | dummy + dummy×covariate |

---

## 2. Continuous covariates

### Linear effect

$$E(y)=\dots+\beta_jx_j+\dots \qquad \frac{\partial E(y)}{\partial x_j}=\beta_j$$

Constant effect everywhere. **Interpretation:** *"holding all other covariates fixed, a one-unit increase in $x_j$ is associated with an estimated $\hat\beta_j$-unit change in expected $y$."*

### Polynomial effect

$$E(y)=\dots+\beta_1x+\beta_2x^2+\dots \qquad \frac{\partial E(y)}{\partial x}=\beta_1+2\beta_2x$$

⚠️ **The effect depends on $x$.** Never quote $\hat\beta_1$ alone.

**Turning point:** $x^*=-\dfrac{\hat\beta_1}{2\hat\beta_2}$ — a maximum if $\hat\beta_2<0$, a minimum if $\hat\beta_2>0$.

**Centred version** (what the exam uses):

$$E(y)=\dots+\beta_1\,\text{age}+\beta_2(\text{age}-48)^2+\dots$$

**Why centre at 48?** Two reasons, both examinable:
1. **Interpretability** — the quadratic term vanishes at age 48, so the remaining coefficients describe a person at a real, in-sample age rather than at age 0.
2. **Reduced multicollinearity** — over a range like 18–80, $\text{age}$ and $\text{age}^2$ correlate above 0.98, inflating both standard errors. $(\text{age}-48)$ and $(\text{age}-48)^2$ are close to uncorrelated.

> 🔴 **Exam Summer 2025, Ex 2(b)** [2 pts]: *why is adding $(\text{age}_i-48)^2$ sensible?*
> → *The effect of age on wage is unlikely to be linear: earnings typically rise in early career, peak in mid career and decline near retirement. A quadratic term lets the model capture this **curvature** (a hump-shaped age–earnings profile) while remaining **linear in the parameters** and hence still estimable by OLS.*
>
> 🔴 **Ex 2(c)** [1 pt]: *positive or negative?*
> → ***Negative.** A negative coefficient makes the parabola open downward, giving the expected hump with a maximum near age 48. A positive coefficient would imply wage is **minimised** around 48 and rises without bound for both very young and very old workers, which is implausible.*

### Logarithmic effects

| Model | Effect of a change in $x$ |
|---|---|
| $y=\beta_0+\beta_1x$ | +1 unit in $x$ ⟹ $\beta_1$ units in $y$ |
| $\log y=\beta_0+\beta_1x$ | +1 unit in $x$ ⟹ $\approx 100\beta_1$ **%** in $y$ |
| $y=\beta_0+\beta_1\log x$ | +1 **%** in $x$ ⟹ $\approx\beta_1/100$ units in $y$ |
| $\log y=\beta_0+\beta_1\log x$ | +1 **%** in $x$ ⟹ $\beta_1$ **%** in $y$ (**elasticity**) |

**Mnemonic:** *the log is on the side that becomes a percentage.*

> 🔴 *WS 23/24, Block III(iv)* tests this by stating the log-log interpretation without saying which variable was logged → **FALSE**.

**When to log the response:**
- $y$ is positive and right-skewed (wage, rent, revenue, price)
- variance grows with the mean (heteroscedasticity)
- effects are plausibly **multiplicative** rather than additive

Recall from Chapter 1: $y=\exp(\boldsymbol{x}'\boldsymbol\beta+\varepsilon)$ becomes $\log y=\boldsymbol{x}'\boldsymbol\beta+\varepsilon$ — an ordinary linear model.
> 🔴 *RCLM WS 22/23, Block II(i):* "$y=\exp(\beta_0+\beta_1x_1+\dots+\varepsilon)$ cannot be analysed within the linear regression framework." → **FALSE.**
>
> ⚠️ The error must be **inside** the exponential. If $y=\exp(\boldsymbol{x}'\boldsymbol\beta)+\varepsilon$, logging does not linearise it.

---

## 3. Categorical covariates — dummy coding

### The rule

> **A categorical covariate with $c$ levels ⟹ $c-1$ dummy variables. The omitted level is the reference category.**

$$D^{(j)}_i=\begin{cases}1&\text{if observation }i\text{ is in level }j\\0&\text{otherwise}\end{cases}$$

### Why $c-1$: the dummy variable trap

All $c$ dummies sum to 1 for every observation — exactly the intercept column. The columns of $\boldsymbol{X}$ are linearly dependent ⟹ $\text{rank}(\boldsymbol{X})<p$ ⟹ $\boldsymbol{X}'\boldsymbol{X}$ singular ⟹ **no unique OLS solution**. This is why A5 exists.

### Interpretation

$\hat\beta^{(j)}$ = estimated difference in expected $y$ between level $j$ and the **reference level**, holding all other covariates fixed.

**To compare two non-reference levels: subtract their coefficients.**

### Worked: Sheet 1

Education, 5 levels, reference `< HS Grad`:

| Level | $\hat\beta$ | vs reference | vs HS Grad |
|---|---|---|---|
| `< HS Grad` | — (ref) | 0 | $-11.44$ |
| `HS Grad` | 11.439 | +11.44 | 0 |
| `Some College` | 24.167 | +24.17 | +12.73 |
| `College Grad` | 39.767 | +39.77 | +28.33 |
| `Advanced Degree` | 64.987 | +64.99 | +53.55 |

**Spotting the reference in R output:** it's the level that **doesn't appear**.

### Alternative codings (know they exist)

| Coding | Setup | $\hat\beta_j$ means |
|---|---|---|
| **Dummy / treatment** ⭐ | drop one level | difference from **reference** |
| **Effect coding** | reference gets $-1$ | difference from the **grand mean** |
| **No intercept, all $c$ dummies** | drop $\beta_0$ | the **level mean** itself |

Dummy coding is R's default and what the exam expects. Mention the others only if asked.

---

## 4. Interactions

$$y=\beta_0+\beta_1x+\beta_2D+\beta_3(x\cdot D)+\varepsilon$$

| Group | Intercept | Slope |
|---|---|---|
| $D=0$ | $\beta_0$ | $\beta_1$ |
| $D=1$ | $\beta_0+\beta_2$ | $\beta_1+\beta_3$ |

> **Dummy alone = parallel lines (shift). Dummy + interaction = non-parallel lines (shift and tilt).**

$$\frac{\partial E(y)}{\partial x}=\beta_1+\beta_3D \qquad\qquad \frac{\partial E(y)}{\partial D}=\beta_2+\beta_3x$$

⚠️ $\hat\beta_1$ is the effect of $x$ **only in the reference group**. $\hat\beta_2$ is the effect of $D$ **only at $x=0$**.

**Types allowed:** any × any.
> 🔴 *RCLM WS 22/23, III(iii):* "Interaction can only be computed between two continuous or between a continuous and a categorical variable." → **FALSE.** Categorical × categorical (product of dummies) is perfectly standard.

**Hierarchy principle:** if $x_1x_2$ is in the model, keep $x_1$ and $x_2$ as main effects. Dropping a main effect imposes an arbitrary constraint that is almost never what you mean.

---

## 5. 📝 The full exam template

**Exam Summer 2025, Ex 2(a)** [3 pts] — wage on age, education (3 levels), place of birth (2 levels).

**Step 1 — inventory:**

```
age         : continuous          → 1 column
education   : 3 levels            → 2 dummies
birthplace  : 2 levels            → 1 dummy
                          p = 1 + 1 + 2 + 1 = 5
```

**Step 2 — define the dummies explicitly:**

$$D^{\text{HS}}_i=\begin{cases}1&i\text{ has a high school degree}\\0&\text{otherwise}\end{cases}\qquad
D^{\text{Col}}_i=\begin{cases}1&i\text{ has a college/university degree}\\0&\text{otherwise}\end{cases}$$

$$D^{\text{out}}_i=\begin{cases}1&i\text{ was born outside the US}\\0&i\text{ was born inside the US}\end{cases}$$

**Step 3 — write the model:**

$$\text{wage}_i=\beta_0+\beta_1\text{age}_i+\beta_2D^{\text{HS}}_i+\beta_3D^{\text{Col}}_i+\beta_4D^{\text{out}}_i+\varepsilon_i,\quad i=1,\dots,n$$

with $\varepsilon_i$ iid, $E(\varepsilon_i)=0$, $\text{Var}(\varepsilon_i)=\sigma^2$.

**Step 4 — state the reference:** a person with **no degree, born inside the US**.

**Step 5 — justify estimability:** $\boldsymbol{X}$ has $p=5$ linearly independent columns (full rank), so $(\boldsymbol{X}'\boldsymbol{X})^{-1}$ exists and OLS is well defined.

### The marking checklist

- [ ] All covariates present
- [ ] Intercept present
- [ ] $c-1$ dummies per categorical variable — **count them**
- [ ] Dummies **defined explicitly** with braces
- [ ] Reference category named
- [ ] Error term with assumptions
- [ ] Index range $i=1,\dots,n$

Seven items, three points. Miss the dummy count and you lose most of them.

---

## 6. Restricted models — the 3.3 preview

Section 3.3 will ask you to estimate a model **under $H_0$**. That means substituting the restriction into the model equation and re-expressing it so OLS can be run directly.

> 🔴 **Exam Summer 2025, Ex 4(d)** [2 pts]: *Model $y_i=\beta_0+\beta_1x_{1i}+\beta_2x_{2i}+\varepsilon_i$; test $H_0:\beta_1=\beta_2+1$. Incorporate $H_0$ into the model equation to obtain a restricted model estimable by OLS.*

**Method — substitute, then collect terms in the free parameters:**

$$y_i=\beta_0+(\beta_2+1)x_{1i}+\beta_2x_{2i}+\varepsilon_i$$
$$y_i=\beta_0+\beta_2x_{1i}+x_{1i}+\beta_2x_{2i}+\varepsilon_i$$
$$y_i-x_{1i}=\beta_0+\beta_2(x_{1i}+x_{2i})+\varepsilon_i$$

**Restricted model:**

$$\boxed{\;\tilde y_i=\beta_0+\beta_2\tilde x_i+\varepsilon_i,\qquad \tilde y_i:=y_i-x_{1i},\quad \tilde x_i:=x_{1i}+x_{2i}\;}$$

Regress the **offset response** $y_i-x_{1i}$ on the **combined covariate** $x_{1i}+x_{2i}$. Two parameters instead of three ⟹ **$r=1$ restriction**. The SSE from this fit is $\text{SSE}_{H_0}$, which feeds straight into the F-statistic.

**The general recipe:**
1. Substitute the restriction to eliminate one parameter
2. Expand and group terms by the **remaining free** parameters
3. Move any covariate terms with **no free parameter** to the left-hand side (they become an offset)
4. The result is an ordinary linear model — count its parameters to confirm $r$

**Practise this with:** $\beta_1=\beta_2$ (⟹ regress on $x_1+x_2$) · $\beta_1=0$ (⟹ drop $x_1$) · $\beta_1+\beta_2=1$ · $\beta_1=2\beta_2$.

---

## 7. Key takeaways

1. Modelling = **choosing the columns of $\boldsymbol{X}$**. Everything stays linear in $\boldsymbol\beta$.
2. **Polynomial:** effect $=\beta_1+2\beta_2x$; turning point $-\hat\beta_1/(2\hat\beta_2)$; **centre** for interpretability and less collinearity.
3. **Logs:** four cases; *the log is on the side that becomes a percentage*.
4. **$c$ levels ⟹ $c-1$ dummies**, reference = the omitted level = the one missing from R output. Compare non-reference levels by **subtracting**.
5. **Interaction ⟹ non-parallel lines.** Main effects become conditional; differentiate.
6. **The Exercise-2 template is seven checkable items.** Rehearse it until it takes two minutes.
7. **Restricted models:** substitute, collect, move offsets left. This is the bridge into the F-test.
