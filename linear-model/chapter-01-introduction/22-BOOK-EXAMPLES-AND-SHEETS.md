# Ch 1 — BOOK EXAMPLES & EXERCISES (intuition first, then solve)

> **On "book exercises":** the Fahrmeir book has **no end-of-chapter exercise sets** — I searched the full PDF. It teaches through worked **Examples** (1.1–1.10). Chapter 1 has no dedicated sheet either; the one course exercise that lives here is **Sheet 3, Exercise 1** (the $R^2$–correlation question).
>
> This is a short file, because Chapter 1 is worth ~2 marks. **Read the intuitions, do the one calculation, move on.**

---

# PART A — THE BOOK'S EXAMPLES

## Examples 1.1–1.4 — The application zoo

| Example | Response $y$ | Type | Model class |
|---|---|---|---|
| **1.1** Munich Rent Index | net rent (€) | continuous | linear model |
| **1.2** Malnutrition in Zambia | child Z-score | continuous | (geo)additive |
| **1.3** Patent Opposition | opposed? yes/no | **binary** | **logit** |
| **1.4** Forest Health Status | damage level | ordered categorical | ordinal |

### 🤔 INTUITION

*Why does the book open with four datasets before any mathematics?*

To make one point land before you can argue with it: **the type of $y$ decides the model class, and nothing else does.**

Test yourself on the awkward part: Example 1.1 has *categorical covariates* (district, kitchen quality) and still uses an ordinary linear model. Example 1.3 has *continuous covariates* and cannot.

> 🔑 **Covariates are ingredients you pour in — chop the categorical ones into dummies and proceed. Only the thing being *predicted* decides which container you need.**
>
> You would not store water in a crate. A linear model on a binary $y$ predicts 1.34 and $-0.20$, and a probability can be neither.

---

## Examples 1.5–1.6 — Univariate distributions

**Situation.** Histograms and boxplots of rent, of area, of Z-scores, before any model.

### 🤔 INTUITION

*What could a histogram of $y$ possibly tell you about a model you haven't fitted?*

More than you'd think. If rent is strongly **right-skewed** — a long tail of expensive flats — then:
- the errors will probably inherit that skew ⟹ **A6 (normality) in trouble**
- for positive skewed quantities, spread usually grows with level ⟹ **A3 (homoscedasticity) in trouble**
- a few very large values will be **high-leverage**

**Three assumption warnings from one picture, before fitting anything.**

*And the fix follows from the diagnosis:* if the natural unit of change is **percentage** rather than absolute — €500 means something different to a €1,000 flat and a €10,000 flat — then model $\log(y)$. You're not applying a trick; you're finally measuring in the right units.

**Quick test you can do in your head:** mean > median ⟹ right-skewed.

---

## Examples 1.7–1.9 — Scatter plots and categorical covariates

### 🤔 INTUITION — the four questions

Every scatter plot gets the same interrogation, regardless of the subject:

1. **Is there a trend?**
2. **Is it straight?** — if not, you need $x^2$ or $\log x$ (§3.1.3)
3. **Is the spread constant?** — if it fans, that's heteroscedasticity (§3.1.2)
4. **Are there outliers?** — leverage and Cook's distance (§3.4.4)

> 🔑 **The payoff:** these are *exactly* the four questions you'll ask of **residual plots** in Section 3.4.4. Same checklist, run twice — once on the raw data before modelling, once on $\hat\varepsilon$ after.
>
> Learn them once here and Chapter 3.4.4 costs you almost nothing.

For a **categorical** covariate (Example 1.9), the equivalent picture is **side-by-side boxplots** — one box per level. Boxes at clearly different heights means dummies for that variable will earn their keep.

---

# PART B — THE ONE EXERCISE

## 📄 Sheet 3, Exercise 1 — $R^2$ and correlation

> *"The estimated model in Exercise 1 on Sheet 1 resulted in an $R^2$ of 0.038. What is the correlation between an individual's wage and his age? Interpret your results and provide one possible explanation."*

### 🤔 INTUITION — think before computing

**Part 1: why can you get $r$ from $R^2$ at all?**

In general you can't. But this is a **simple** regression — one covariate, one intercept. There, and only there, $R^2=r_{xy}^2$.

*Why?* $R^2$ is the fraction of variance in $y$ explained by $\hat y$. With one covariate, $\hat y$ is just a linear rescaling of $x$ — and correlation is invariant to linear rescaling. So "how much of $y$ does $\hat y$ explain" and "how strongly do $x$ and $y$ correlate" are the same question.

> ⚠️ **In multiple regression this breaks.** There $R^2=\text{corr}(y,\hat y)^2$, not the squared correlation with any single covariate.

**Part 2: square-rooting loses the sign. Where do you get it back?**

From $\hat\beta_1$. In simple regression $\hat\beta_1=r\cdot s_y/s_x$, and $s_y/s_x>0$, so **$\hat\beta_1$ and $r$ always share a sign**. Sheet 1 gave $\hat\beta_1=+0.71$ ⟹ $r>0$.

**Part 3 — the real question. $R^2=0.038$ is tiny. Does that mean age doesn't matter?**

**No.** And working out *why not* is the whole exercise. Three distinct possibilities:

1. **The relationship isn't linear.** And here you already know it isn't — Sheet 2 shows the age–wage scatter is clearly a **hump**: rising, peaking near 53, then falling. A straight line through a hump has almost no slope, because the rise and the fall cancel. **Pearson correlation and straight-line $R^2$ are blind to curvature.** They aren't measuring "is there a relationship" — they're measuring "is there a *linear* one."

2. **Omitted variables.** Age is a minor determinant of wage compared to education, job class, health. Sheet 5's full model reaches $R^2=0.34$.

3. **Wage is intrinsically noisy.** Individual variation is large no matter what you include.

> 🔑 **The transferable lesson: a low $R^2$ is not evidence of no relationship. It's an invitation to look harder — usually at the functional form or at what you left out.**

### ✍️ SOLUTION

**(a)** In simple linear regression with an intercept, $R^2=r_{xy}^2$, so

$$|r|=\sqrt{0.038}=0.1949$$

The sign follows $\hat\beta_1=+0.71>0$:

$$\boxed{r=+0.195}$$

**(b) Interpretation.** *There is a **weak positive linear** association between age and wage: older men in this sample tend to earn slightly more, but age explains only **3.8%** of the variation in wage.*

**(c) Explanation.** *The most likely reason is that **the relationship is not linear**. The scatter plot in Sheet 2 shows a clearly quadratic, hump-shaped pattern — wage rises through early and mid career, peaks around age 53, then declines. Pearson correlation and the $R^2$ of a straight-line fit measure only **linear** association, so the rising and falling portions largely cancel and the true strength of the age–wage relationship is badly understated.*

*A second contributing reason is **omitted variables**: education, job class and health are far stronger determinants of wage than age, and a model including them reaches $R^2=0.34$.*

---

# 🎯 THE INTUITIONS, COLLECTED

| # | Intuition | From |
|---|---|---|
| 1 | 🔑 **Only the type of $y$ chooses the model class.** Covariates are just ingredients | Ex 1.1–1.4 |
| 2 | You can't store water in a crate — a linear model on binary $y$ predicts 1.34 | Ex 1.3 |
| 3 | A histogram of $y$ warns you about **three** assumptions before you fit anything | Ex 1.5–1.6 |
| 4 | mean > median ⟹ right-skewed ⟹ consider $\log(y)$ | Ex 1.5 |
| 5 | Logging isn't a trick — it's switching the ruler from **euros to percent** | Ex 1.5 |
| 6 | 🔑 **The four scatter-plot questions reappear as the four residual plots** in 3.4.4 | Ex 1.7 |
| 7 | Categorical covariate ⟹ side-by-side boxplots ⟹ will dummies help? | Ex 1.9 |
| 8 | $R^2=r^2$ **only** in simple regression; the sign comes from $\hat\beta_1$ | Sheet 3(1) |
| 9 | 🔑 **A low $R^2$ measures the absence of a *linear* relationship, not of a relationship** | Sheet 3(1) |
| 10 | A straight line through a hump has near-zero slope — the halves cancel | Sheet 3(1) |
