# Ch 2 — BOOK EXAMPLES & EXERCISE SHEETS (intuition first, then solve)

> **Note on "book exercises":** the Fahrmeir book has **no end-of-chapter exercise sets** — I searched the whole PDF. It teaches through worked **Examples** (2.1, 2.2, …) on the Munich Rent Index and patent data. Your course's real exercises are **Sheets 1 and 2**, which map exactly onto this chapter.
>
> **Rule for this file: read the `🤔 INTUITION`, close the file, try it, THEN read the `✍️ SOLUTION`.** Reading solutions feels like learning and isn't.

---
---

# PART A — THE BOOK'S EXAMPLES

## Examples 2.1 & 2.2 — Munich Rent Index: Simple Linear Regression

**Situation.** Net rent regressed on living area.

### 🤔 INTUITION

*Before fitting anything: what do you expect the intercept to mean here?*

$\hat\beta_0$ = expected rent of a flat with **zero square metres**. That flat doesn't exist. Yet the number will be positive and non-trivial, because the line has to pass through $(\bar x,\bar y)$ and it tilts up.

**So the intercept is doing a job (positioning the line) without carrying a meaning.** This is the general situation, not an exception, and it's why Sheet 1 asks about it directly.

*Second question: would you expect a straight line to fit?*

Think about rent **per square metre**. Big flats usually cost less per m² — fixed costs (kitchen, bathroom, door) get spread over more area. That means total rent rises with area but **decelerates**. A straight line will systematically over-predict at the extremes and under-predict in the middle.

### ✍️ WHAT THE BOOK DOES

Fits the simple model, then in Example 2.2 introduces a **transformation** of living area to capture the curvature — while keeping an ordinary linear model, because the model is linear in $\boldsymbol\beta$ regardless of what you do to $x$ first.

---

## Example 2.3 — Rent in Average vs Good Locations

**Situation.** A binary location indicator added to the model.

### 🤔 INTUITION

*What does adding a 0/1 variable do to the picture, geometrically?*

It adds a constant to the prediction for one group and not the other. The fitted line for "good location" is the average-location line **shifted vertically** by $\hat\beta_{\text{location}}$.

**Two parallel lines.** Same slope, different height. The location premium is the same €X for a 30 m² flat and a 200 m² flat.

*Is that plausible?* Probably not — you'd expect a good location to be worth *more* in absolute terms on a bigger flat. Which is exactly why the book goes on to Example 2.5.

### ✍️ WHAT TO TAKE

> **A dummy on its own shifts the line. It cannot tilt it.**

---

## Example 2.5 — Interaction Between Living Area and Location ⭐

### 🤔 INTUITION

Example 2.3's model forces one number to describe the location premium at every size. To let the premium **grow with size**, you need the location effect to depend on area — i.e. a **product term**.

$$\text{rent}=\beta_0+\beta_1\text{area}+\beta_2\text{loc}+\beta_3(\text{area}\times\text{loc})+\varepsilon$$

Now split by group and read it off:

| Group | Intercept | Slope |
|---|---|---|
| average location (loc=0) | $\beta_0$ | $\beta_1$ |
| good location (loc=1) | $\beta_0+\beta_2$ | $\beta_1+\beta_3$ |

**The lines are no longer parallel.**

*And here's the consequence students miss:* $\beta_2$ is now the location premium **at area = 0** — a flat that doesn't exist. It is no longer "the location effect." The location effect is $\beta_2+\beta_3\cdot\text{area}$, a **function**, not a number.

### ✍️ THE RULE

> **Dummy alone = parallel lines (shift). Dummy + interaction = non-parallel lines (shift AND tilt).**
>
> **Once a variable appears in two terms, differentiate. Never quote one coefficient.**

---

## Examples 2.7 & 2.8 — Patent Opposition (the logit examples)

**Situation.** Response: was the patent opposed? Yes/no. Binary.

### 🤔 INTUITION — the whole of Section 2.3 in one thought

The mean of a 0/1 variable **is** the probability it equals 1. So modelling $E(y\mid\boldsymbol{x})$ means modelling a probability. And a probability **lives in $[0,1]$**.

Now look at $\boldsymbol{x}'\boldsymbol\beta$. It's a straight line in $k$ dimensions. Lines run to $\pm\infty$.

**You are trying to store something bounded in a container with no walls.**

Everything else follows from that single mismatch:
- predictions escape $[0,1]$ ← the fatal one
- $\text{Var}(y)=\pi(1-\pi)$ depends on $\boldsymbol{x}$ ⟹ heteroscedastic **by construction**
- $\varepsilon$ takes only two values ⟹ can't be normal
- a constant marginal effect is impossible near the boundaries

*So what's the fix?* Not to constrain $\boldsymbol\beta$ — that can't work. Instead, **change the scale on which the model is linear.** Find a function mapping $(0,1)$ onto all of $\mathbb{R}$, and be linear *there*.

Odds take $[0,1]$ to $[0,\infty)$. Log takes $[0,\infty)$ to $(-\infty,\infty)$. **Two moves and the response is free.**

$$\log\frac{\pi}{1-\pi}=\boldsymbol{x}'\boldsymbol\beta$$

### ✍️ THE CONSEQUENCE FOR INTERPRETATION

If the model is linear in the **log-odds**, then $\hat\beta_j$ is a change in **log-odds** — and $\exp(\hat\beta_j)$ multiplies the **odds**.

It is *not* a change in probability. The probability effect is $\hat\beta_j\pi(1-\pi)$, which depends on where you are.

> 🔑 **Ask yourself why that has to be true.** If $\hat\beta_j$ *were* a constant probability change, you'd be back to the linear probability model, and the fitted values would escape $[0,1]$ again. **The whole point of the link is that the probability effect can't be constant.**

---
---

# PART B — THE EXERCISE SHEETS

## 📄 SHEET 1 — Simple and multiple regression

### Ex 1 — Simple linear regression of wage on age

Given $\hat\beta_0=81.70$, $\hat\beta_1=0.71$.

#### 🤔 INTUITION

**(b) Interpreting the slope.** The question is really "can you write a precise sentence?" Four things must appear: the **unit of $x$**, the **unit of $y$**, the word **"expected"**, and **"associated with"** (not "causes"). Miss one and you lose fractions that add up.

**(c) Why not interpret $\hat\beta_0$?** *Ask: what covariate value does the intercept describe?* Age zero. Is that in the data? No — the `Wage` dataset runs roughly 18–80. Is it meaningful? A newborn has no hourly wage.

> **The general rule this teaches:** the intercept is interpretable only when $x=0$ is both **in range** and **meaningful**. Otherwise it positions the line and nothing more. *(And if you want it to mean something — centre the variable. That's why the exam writes $(\text{age}-48)$.)*

#### ✍️ SOLUTION

**(a)** $\;\text{wage}_i=\beta_0+\beta_1\text{age}_i+\varepsilon_i,\quad i=1,\dots,n$, with $\varepsilon_i$ iid, $E(\varepsilon_i)=0$, $\text{Var}(\varepsilon_i)=\sigma^2$.

**(b)** *A one-year increase in age is associated with an estimated increase of **\$0.71** in expected hourly wage.*

**(c)** *Because $\hat\beta_0$ is the expected wage at **age 0**, which lies far outside the observed range of the data and is substantively meaningless. Interpreting it would be extrapolation into a region where we have no evidence the linear relationship holds. The intercept is still needed to position the regression line, but carries no substantive meaning.*

---

### Ex 2 — Multiple regression with education dummies ⭐

#### 🤔 INTUITION — the most important on this sheet

**(a) Why 4 dummies for 5 levels?**

Picture marking one person's height on a wall and reporting everyone else **relative to the mark**. Five people, four numbers — you never need a number for the person who *is* the mark.

*What if you insisted on five?* Then "intercept = 10, level1 = 0" and "intercept = 0, level1 = 10" describe **identical** models. Infinitely many parameter sets give the same fit. The computer can't choose, because the question has no unique answer. That's $\boldsymbol{X}'\boldsymbol{X}$ being **singular** — the **dummy variable trap**.

> 🔑 **The algebra refuses to let you ask a comparative question without naming the comparison.** That's a feature.

**(b) Which is the reference?** The level that **doesn't appear** in the output. Output shows education2–5; education1 (`< HS Grad`) is missing ⟹ it's the baseline.

**(c) The shortcut worth internalising.** Two men, both 50. Comparing level 3 to level 5.

*Before computing: what cancels?* The intercept (both have it) and the age term (both are 50). Everything shared vanishes.

**So don't build two predictions. Subtract the two coefficients.** One operation instead of six, and five fewer chances to slip.

⚠️ And note: you must **subtract**, because each coefficient is measured from the *reference*, not from each other.

**(d) What doesn't cancel?** Only the intercept. Age differs *and* education differs, so both contribute.

#### ✍️ SOLUTION

**(a)** 5 levels ⟹ **4 dummies**, reference `< HS Grad`:
$$D_2=\mathbb{1}\{\text{HS Grad}\},\ D_3=\mathbb{1}\{\text{Some College}\},\ D_4=\mathbb{1}\{\text{College Grad}\},\ D_5=\mathbb{1}\{\text{Advanced Degree}\}$$
$$\text{wage}_i=\beta_0+\beta_1\text{age}_i+\beta_2D_{2i}+\beta_3D_{3i}+\beta_4D_{4i}+\beta_5D_{5i}+\varepsilon_i$$

**(b)** Reference: **`< HS Grad`**.
> *Holding age fixed, high school graduates earn on average an estimated \$11.44 more per hour than men with less than a high school degree. The estimated premium over the same reference is \$24.17 for Some College, \$39.77 for College Grad and \$64.99 for Advanced Degree. The coefficients increase monotonically, so expected wage rises steadily with education.*

**(c)** Both aged 50 ⟹ intercept and age cancel:
$$\hat\beta_5-\hat\beta_3=64.98656-24.16700=\boxed{\$40.82}$$

**(d)** Only the intercept cancels; the 20-year-old is in the reference category so contributes 0 for education:
$$\hat\beta_1(40-20)+\hat\beta_4=0.56869\times20+39.76677=11.374+39.767=\boxed{\$51.14}$$

*Check by full prediction:* $122.85-71.71=51.14$ ✓

---

## 📄 SHEET 2 — Polynomials and interactions

### Ex 1 — Polynomial regression

$\hat\beta_0=-10.43$, $\hat\beta_1=5.29$, $\hat\beta_2=-0.05$.

#### 🤔 INTUITION

**Why does this deserve a quadratic at all?** Because the scatter plot said so. Sheet 2 opens with *"the effect of age seems to be rather quadratic than linear"* — a decision made from a **picture**, not a test. That's Chapter 1.2 doing its job.

**Now, before interpreting: is this still a linear model?**

Yes. Set $x_1=\text{age}$, $x_2=\text{age}^2$ and run ordinary OLS. **Linear refers to $\boldsymbol\beta$.** The *curve* is a parabola; the *model* is linear.

**And the interpretation trap — think of a thrown ball.**

$$\text{wage}=-10.43+5.29\,\text{age}-0.05\,\text{age}^2$$

is the same equation as a ball in flight. $\hat\beta_1$ is the **launch speed**; $\hat\beta_2$ is **gravity**.

*If someone asks "how fast is the ball moving?" — you cannot answer without asking "when?"* At launch it's fast and rising; at the top it's zero; on the way down it's negative.

> 🔑 **So "$\hat\beta_1=5.29$ means wage rises \$5.29 per year" is exactly as wrong as "the ball travels at its launch speed the whole time."** It's the effect at age zero and nowhere else.
>
> The effect is $\;\partial\hat y/\partial\text{age}=\hat\beta_1+2\hat\beta_2\,\text{age}$.

**And you can predict the sign of $\hat\beta_2$ without data:** careers, like balls, come back down. **Negative.**

#### ✍️ SOLUTION

**(a)** $\;\text{wage}_i=\beta_0+\beta_1\text{age}_i+\beta_2\text{age}_i^2+\varepsilon_i$ — still a linear model, estimable by OLS.

**(b)** *$\hat\beta_2=-0.05<0$, so the fitted curve is a downward-opening parabola: wage rises with age at first, reaches a maximum, then declines. This matches the economic pattern of earnings rising through early and mid career, peaking, then falling as workers move to part-time or lower-paid late-career roles.*

Turning point:
$$\text{age}^*=-\frac{\hat\beta_1}{2\hat\beta_2}=-\frac{5.29}{2(-0.05)}=\frac{5.29}{0.10}=\boxed{52.9\text{ years}}$$

Consistent with the curve in Figure 2 ✓

Marginal effects: at age 30, $5.29-0.10(30)=+\$2.29$/yr. At age 65, $5.29-6.50=-\$1.21$/yr.

---

### Ex 2 — Interaction between age and health ⭐

$$\widehat{\text{wage}}=78.66+0.51\,\text{age}-1.81\,H+0.43\,\text{age}\cdot H$$

#### 🤔 INTUITION

**Two escalators.** Both go up. Without an interaction they run at the same speed — one just starts higher, and the gap **never changes**. With an interaction they run at **different speeds**: the gap changes, and they may cross.

**(b) The move:** don't try to interpret four coefficients in the abstract. **Set $H=0$, then $H=1$, and read off two lines.**

**(c) The geometric answer.** Without the interaction, the dummy only shifts the intercept ⟹ **parallel** lines ⟹ the health effect is one constant number at every age. With it, each group gets its own slope ⟹ **non-parallel** ⟹ the health effect **depends on age**.

**And the trap:** $-1.81$ is *not* "the health penalty." It's the health effect **at age 0** — the very bottom of the escalator, where nobody stands. The real effect is $-1.81+0.43\,\text{age}$: at 40 that's $+\$15.39$; at 60, $+\$24.99$.

> 🔑 *"How far ahead is escalator B?"* has no single answer. **It depends when you ask.** Same structure as the ball. **Same rule: variable in two terms ⟹ differentiate.**

#### ✍️ SOLUTION

**(a)** $H_i=\mathbb{1}\{\text{health}\geq\text{Very Good}\}$:
$$\text{wage}_i=\beta_0+\beta_1\text{age}_i+\beta_2H_i+\beta_3(\text{age}_i\times H_i)+\varepsilon_i$$

**(b)**

| Group | Intercept | Slope |
|---|---|---|
| health ≤ Good ($H=0$) | $78.66$ | $0.51$ |
| health ≥ Very Good ($H=1$) | $78.66-1.81=\boxed{76.85}$ | $0.51+0.43=\boxed{0.94}$ |

*Healthier men have a slightly lower fitted intercept but their wage rises with age nearly **twice as fast** (\$0.94/yr vs \$0.51/yr).*

**(c)** *Without an interaction the dummy shifts the line vertically but the slope is shared, so the two fitted lines are **parallel** and the health effect is a constant \$16.90 at every age. With an interaction each group has its own slope, so the lines are **no longer parallel** — they may converge, diverge or cross. The effect of health on wage now depends on age, and equivalently the effect of age depends on health.*

*The lines cross at $1.81/0.43=4.21$ years — far outside the data range, so within the observed ages the healthier group is always ahead, and increasingly so.*

---

# 🎯 THE INTUITIONS, COLLECTED

| # | Intuition | From |
|---|---|---|
| 1 | The intercept **positions** the line; it only **means** something if $x=0$ is in range and sensible | Ex 2.1, Sheet 1(1c) |
| 2 | Fixed costs spread over more area ⟹ rent is **concave** in size | Ex 2.2/3.3 |
| 3 | **A dummy shifts. It cannot tilt.** | Ex 2.3 |
| 4 | You never need a number for the person who **is** the mark ⟹ $c-1$ dummies | Sheet 1(2a) |
| 5 | All $c$ dummies ⟹ **infinitely many** equivalent answers ⟹ singular $\boldsymbol{X}'\boldsymbol{X}$ | Sheet 1(2a) |
| 6 | The reference is the level **missing** from the output | Sheet 1(2b) |
| 7 | **Shared characteristics cancel** ⟹ subtract coefficients, don't build two predictions | Sheet 1(2c) |
| 8 | The scatter plot chooses the model, not a test | Sheet 2(1) |
| 9 | 🔑 **A polynomial is a thrown ball.** $\hat\beta_1$ = launch speed, $\hat\beta_2$ = gravity | Sheet 2(1) |
| 10 | "How fast is the ball going?" needs a **when** ⟹ differentiate | Sheet 2(1), 2(2) |
| 11 | $\hat\beta_2<0$ is predictable from theory — careers come back down | Exam 2025 Ex 2(c) |
| 12 | **Two escalators:** interaction = different speeds = non-parallel lines | Sheet 2(2) |
| 13 | With an interaction, a main effect is the effect **at zero** of the other variable | Sheet 2(2) |
| 14 | 🔑 A probability is **bounded**; $\boldsymbol{x}'\boldsymbol\beta$ is **not**. That mismatch generates all four objections | Ex 2.7/2.8 |
| 15 | Odds uncage $[0,1]\to[0,\infty)$; log uncages $\to\mathbb{R}$. **Two moves** | Ex 2.7/2.8 |
| 16 | $\hat\beta_j$ **can't** be a constant probability change — that's the broken model you escaped | Ex 2.8 |
