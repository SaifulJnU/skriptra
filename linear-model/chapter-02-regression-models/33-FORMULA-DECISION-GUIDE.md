# Ch 2 — FORMULA DECISION GUIDE

> **The question this file answers:** *"I'm looking at an exam question. Which formula do I reach for, and how did I know?"*
>
> Every entry has the same six fields:
> **① Formula ② USE WHEN ③ DON'T USE WHEN ④ WHY ⑤ 🔍 TRIGGER PHRASES in the question ⑥ ⚠️ Errors**
>
> Read the 🔍 lines the night before the exam.

---

# ⚡ THE 30-SECOND TRIAGE

| The question says… | Go to |
|---|---|
| "simple regression by hand" / given $\bar x,\bar y$ or raw data | **F1** |
| "interpret $\hat\beta_1$" (plain, no interaction/polynomial) | **F2** |
| the covariate is a **category** with several levels | **F3, F4** |
| comparing **two non-reference** categories | **F5** |
| a covariate appears **twice** (e.g. $x$ and $x^2$, or in an interaction) | **F6, F7** |
| "interaction term" / non-parallel lines | **F7** |
| "why can't we use linear regression here" for binary $y$ | **F8** |
| $\exp(\hat\beta_j)$ / "odds ratio" / logit interpretation | **F9** |
| "probability effect" of a logit coefficient | **F10** |

---

# F1 — Simple linear regression by hand

### ① Formula

$$
\hat\beta_1=\frac{\widehat{\text{Cov}}(x,y)}{\widehat{\text{Var}}(x)}=r\frac{s_y}{s_x} \qquad\qquad \hat\beta_0=\bar y-\hat\beta_1\bar x
$$

### ② USE WHEN

- Raw $(x,y)$ pairs or summary statistics ($\bar x,\bar y,s_x,s_y,r$) are given and you're asked to fit or reconstruct a simple regression
- An R output is **missing the intercept** but gives you $\bar x, \bar y$, and the slope

### ③ DON'T USE WHEN

- ❌ There is more than one covariate — this is the **simple**-regression-only shortcut; use $(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$ for multiple regression (Ch 3, F1)
- ❌ You're not given enough summary statistics to compute $S_{xy}/S_{xx}$ or $r,s_x,s_y$

### ④ WHY

The line of best fit always passes through $(\bar x,\bar y)$ — that single geometric fact is exactly what lets you recover $\hat\beta_0$ once you know the slope and the two means.

### ⑤ 🔍 TRIGGERS

> *"the average was 48.61 [x] and 46.61 [y]"* · a table missing the `(Intercept)` row · *"reconstruct the fitted line"*

### ⑥ ⚠️ ERRORS

- Trying to use this shortcut in a **multiple**-regression setting — it only works with one covariate
- Forgetting the line **must** pass through the means — if your recovered intercept doesn't reproduce $\bar y$ at $x=\bar x$, you've made an arithmetic slip

---

# F2 — The plain slope interpretation sentence

### ① Formula (a sentence, not an equation)

> *"A one-[unit] increase in [$x$] is associated with an estimated $\hat\beta_1$ [unit] change in the expected [$y$], holding all other covariates fixed."*

### ② USE WHEN

- Asked to "interpret $\hat\beta_1$" and $x$ appears **exactly once** in the model, with no interaction and no polynomial term

### ③ DON'T USE WHEN

- ❌ $x$ appears in more than one term (polynomial, interaction) — go to F6/F7, differentiate first
- ❌ The intercept is being interpreted and $x=0$ is not meaningful/in-range — don't interpret it; centre the covariate instead
- ❌ $y$ or $x$ has been logged — use F4 from the Chapter 1 guide instead

### ④ WHY

Regression models the **conditional mean**, so every interpretation must say "expected" or "on average" — you never observe the exact deterministic change, only the average one, and only after holding the rest of $\boldsymbol{x}$ fixed since $\hat\beta_1$ is a **partial** effect in a multiple regression.

### ⑤ 🔍 TRIGGERS

> *"interpret the coefficient on age"* · *"what does $\hat\beta_1=0.62$ mean?"*

### ⑥ ⚠️ ERRORS

- Dropping "holding other covariates fixed" in a **multiple** regression — the coefficient's meaning genuinely changes without that clause
- Saying "causes" instead of "associated with"
- Missing the **unit** of $x$ or $y$ — always name them

---

# F3 — Dummy variable coding

### ① Rule

$$
c\ \text{levels} \Longrightarrow c-1\ \text{dummy columns} + \text{one reference level (coded implicitly as all-zeros)}
$$

### ② USE WHEN

- A covariate is **categorical** with more than 2 levels (education, region, job class, …)
- Building $\boldsymbol{X}$ from a verbal description that includes a category

### ③ DON'T USE WHEN

- ❌ The covariate is **binary already** (e.g. sex, yes/no) — that's just 1 dummy, no special rule needed
- ❌ You're tempted to include **all $c$ levels plus the intercept** — that's the dummy variable trap: $\boldsymbol{X}'\boldsymbol{X}$ becomes singular, no unique OLS solution exists

### ④ WHY

If every level got its own dummy *and* you kept the intercept, the dummy columns would sum to exactly the intercept column (since every observation is in exactly one category) — a perfect linear dependency, so $\text{rank}(\boldsymbol{X})<p$ and $(\boldsymbol{X}'\boldsymbol{X})^{-1}$ doesn't exist.

### ⑤ 🔍 TRIGGERS

> *"education has 5 categories"* · *"health status: poor / fair / good / very good / excellent"* · any R table showing dummy names like `education2.HS Grad`

**The reference level is the one MISSING from the R output.**

### ⑥ ⚠️ ERRORS

- 🔴 Using $c$ dummies instead of $c-1$
- Forgetting to **name the reference category** explicitly when writing the model equation — this loses an interpretation mark even if the coefficients are right
- Treating dummies as **cumulative** ("College Grad" does NOT also get the "Some College" coefficient — they're mutually exclusive)

---

# F4 — Interpreting a single dummy coefficient

### ① Formula

$$
\hat\beta_j = \text{expected difference in } y \text{ between level } j \text{ and the reference level, other covariates held fixed}
$$

### ② USE WHEN

- Asked to interpret one dummy coefficient directly

### ③ DON'T USE WHEN

- ❌ Comparing **two non-reference** levels to each other — that's F5, not this
- ❌ There's an interaction between this dummy and another covariate — the "effect" is then not a single number (F7)

### ④ WHY

By construction, the reference category's dummy is implicitly 0 across the board, so $\boldsymbol\beta_0$ already absorbs its mean — every other dummy's coefficient is measured relative to that baseline.

### ⑤ 🔍 TRIGGERS

> *"interpret the coefficient for Advanced Degree"* — **always says "compared to the reference (< HS Grad)"** in a correct answer

### ⑥ ⚠️ ERRORS

- Interpreting $\hat\beta_j$ as an absolute level of $y$ instead of a **difference from the reference**
- Omitting which category is the reference

---

# F5 — Comparing two non-reference dummy levels

### ① Formula

$$
\Delta = \hat\beta_A - \hat\beta_B
$$

### ② USE WHEN

- The question asks to compare two categories, **neither of which is the reference**

### ③ DON'T USE WHEN

- ❌ One of the two categories **is** the reference — then just read $\hat\beta_j$ directly (F4), no subtraction needed

### ④ WHY

Both $\hat\beta_A$ and $\hat\beta_B$ are measured relative to the **same** reference, so that reference cancels out algebraically when you subtract — what's left is exactly the A-vs-B difference.

### ⑤ 🔍 TRIGGERS

> *"what is the difference in expected wage between College Grad and Some College?"* · any "compare group X to group Y" where neither is the omitted level

### ⑥ ⚠️ ERRORS

- 🔴 Using $\hat\beta_A$ **alone** as "the difference" — that compares A to the *reference*, not to B. This is the single most common dummy-variable error.
- Sign confusion — check whether the question wants "A minus B" or "B minus A"

---

# F6 — A covariate appearing in two terms (polynomial)

### ① Formula

$$
y=\beta_0+\beta_1x+\beta_2x^2+\varepsilon \qquad\Longrightarrow\qquad \text{effect of }x = \frac{\partial y}{\partial x}=\hat\beta_1+2\hat\beta_2x
$$

$$
\text{turning point: } x^*=-\frac{\hat\beta_1}{2\hat\beta_2}
$$

### ② USE WHEN

- $x$ appears both as $x$ and $x^2$ (or higher powers) in the model
- Asked for "the effect of $x$ **at** a specific value" or "the age at which wage is maximised"

### ③ DON'T USE WHEN

- ❌ Only $x$ (no $x^2$) appears — the effect is just $\hat\beta_1$, constant, use F2
- ❌ Asked to interpret $\hat\beta_1$ **alone** as "the effect of $x$" — this is never correct once $x^2$ is in the model

### ④ WHY

The true marginal effect of $x$ on $E(y)$ is the derivative of the whole expression with respect to $x$, and once $x$ appears in more than one term, that derivative is no longer a single constant — it's a function of $x$ itself.

### ⑤ 🔍 TRIGGERS

> *"at what age is expected wage maximised?"* · *"what is the effect of age at age 30?"* · any model with an `age` and `age^2` (or `I(age^2)`) term in the R output

### ⑥ ⚠️ ERRORS

- 🔴 Quoting $\hat\beta_1$ alone as "the effect of $x$" — the classic trap
- Sign of $\hat\beta_2$: **negative** ⟹ downward parabola ⟹ turning point is a **maximum** (the usual wage/age story); **positive** ⟹ minimum
- Plugging into the turning-point formula without checking it's inside the observed data range — outside the range, extrapolation warnings apply

---

# F7 — Interaction terms

### ① Formula

$$
y=\beta_0+\beta_1x+\beta_2D+\beta_3(xD)+\varepsilon
$$

$$
\frac{\partial y}{\partial x}=\hat\beta_1+\hat\beta_3D \qquad\qquad \frac{\partial y}{\partial D}=\hat\beta_2+\hat\beta_3x
$$

| | Intercept | Slope |
|---|---|---|
| $D=0$ | $\beta_0$ | $\beta_1$ |
| $D=1$ | $\beta_0+\beta_2$ | $\beta_1+\beta_3$ |

### ② USE WHEN

- A model includes a **product term** between two covariates (dummy×continuous, dummy×dummy, or continuous×continuous)
- Asked whether "the effect of $x$ is the same across groups"

### ③ DON'T USE WHEN

- ❌ No product term is present — go back to F2 for a plain slope
- ❌ Asked for the effect of $x$ **without specifying** the value of the interacting variable — you must be given or asked to hold a specific value/group, since the effect is no longer a single number

### ④ WHY

An interaction term is the model's way of saying the world doesn't add up separably — the effect of one covariate genuinely **depends** on the value of another, and the algebra reflects that by making the partial derivative a function of the second variable rather than a constant.

### ⑤ 🔍 TRIGGERS

> *"does the effect of age differ by health status?"* · *"give the slope for [group]"* · any R output with a term like `age:healthVeryGood` or `x1:x2`

### ⑥ ⚠️ ERRORS

- Quoting $\hat\beta_1$ as "the effect of $x$" when an interaction is present — it's only the effect **when $D=0$**
- Quoting $\hat\beta_2$ as "the effect of $D$" — it's only the effect **when $x=0$**, which may not be meaningful (see Ch1 F4 on centring)
- Describing an interaction model as producing **parallel** lines — interaction is precisely what makes them **non-parallel**

---

# F8 — Why linear regression fails for a binary response

### ① The four reasons (a list, not a formula)

1. 🔴 **Predicted probabilities can fall outside $[0,1]$** — lead with this one
2. $\text{Var}(y\mid\boldsymbol{x})=\pi(1-\pi)$ depends on $\boldsymbol{x}$ ⟹ **heteroscedasticity by construction**, not by accident
3. $\varepsilon$ can only take **two values** (for a given $\boldsymbol{x}$) ⟹ can never be normally distributed
4. A **constant** marginal effect is implausible near the boundaries — the same $\Delta x$ can't move a probability from 0.98 to 1.08

### ② USE WHEN

- Asked to justify **why** a logit/probit model is used instead of OLS for a binary outcome

### ③ DON'T USE WHEN

- ❌ The response is continuous — none of this applies; use the ordinary linear model

### ④ WHY

All four reasons trace back to one structural fact: $E(y\mid\boldsymbol{x})=P(y=1\mid\boldsymbol{x})$ is a **probability**, and probabilities are bounded, but the linear predictor $\boldsymbol{x}'\boldsymbol\beta$ is not.

### ⑤ 🔍 TRIGGERS

> *"why is a linear model not appropriate here?"* · *"the response is whether the applicant was approved"*

### ⑥ ⚠️ ERRORS

- Giving only one reason when the question implies several ("give reasons," plural)
- Not leading with the boundedness argument — it's the most fundamental and most-rewarded point

---

# F9 — Logit interpretation: the odds ratio

### ① Formula

$$
\log\frac{\pi}{1-\pi}=\boldsymbol{x}'\boldsymbol\beta \qquad\Longrightarrow\qquad \frac{\pi}{1-\pi}\ \text{is multiplied by}\ \exp(\hat\beta_j)\ \text{for a one-unit increase in } x_j
$$

### ② USE WHEN

- Asked to interpret a logit coefficient "in terms of the odds"
- Given a fitted logit model and asked "how does a one-unit change in $x_j$ affect the outcome?"

### ③ DON'T USE WHEN

- ❌ Asked specifically for the effect on **probability** — that's F10, and it's a different (non-constant) number
- ❌ The model is a plain linear model — no odds concept applies there

### ④ WHY

$\hat\beta_j$ enters the model additively on the **log-odds** scale, so exponentiating converts an additive change into a **multiplicative** one on the odds scale — this is just the algebra of $\log$ and $\exp$ being inverses, but it's the reason odds ratios are *the* portable, constant-across-$\pi$ way to describe a logit effect.

### ⑤ 🔍 TRIGGERS

> *"interpret $\hat\beta_j$ in terms of the odds"* · *"by what factor do the odds change...?"* · $\exp(\hat\beta_j)$ appears in an R output or is asked for directly

### ⑥ ⚠️ ERRORS

- 🔴 Saying $\hat\beta_j$ "increases $P(y=1)$ by $\hat\beta_j$" — **false**, that confuses the log-odds scale with the probability scale (a repeated exam T/F trap)
- Forgetting to **exponentiate** — the odds ratio is $\exp(\hat\beta_j)$, not $\hat\beta_j$ itself
- Misreading $\exp(\hat\beta_j)>1$ as "increases the probability by that much" — it increases the **odds** multiplicatively; the probability change depends on where you start (F10)

---

# F10 — Logit interpretation: the (non-constant) probability effect

### ① Formula

$$
\frac{\partial \pi}{\partial x_j} = \hat\beta_j\,\pi(1-\pi)
$$

### ② USE WHEN

- Explicitly asked for the effect on the **probability scale**, not the odds scale
- Asked for the **maximum possible** probability effect of a coefficient

### ③ DON'T USE WHEN

- ❌ The question just says "interpret the coefficient" with no scale specified — default to the odds-ratio interpretation (F9), which is exact and constant; only use this one when probability is explicitly requested, since it requires picking a value of $\pi$ to evaluate at

### ④ WHY

This is the calculus derivative of the logistic function $\pi=e^\eta/(1+e^\eta)$ with respect to $\eta$, chained through $\eta=\boldsymbol{x}'\boldsymbol\beta$ — and because $\pi(1-\pi)$ is not constant, neither is this effect: it depends on **where** on the S-curve you evaluate it.

### ⑤ 🔍 TRIGGERS

> *"what is the maximum possible change in probability associated with a one-unit change in $x_j$?"* → $\pi(1-\pi)$ is maximised at $\pi=0.5$, where it equals $0.25$, so the answer is $0.25\hat\beta_j$

### ⑥ ⚠️ ERRORS

- Treating this as a **constant** effect like the odds ratio — it isn't, and stating a single number without specifying $\pi$ is incomplete
- Forgetting that $\pi(1-\pi)$ peaks at $\pi=0.5$ (value $0.25$) and shrinks toward the boundaries

---

# 🚦 THE MASTER FLOWCHART

```
                    READ THE QUESTION
                            │
        ┌───────────────────┼───────────────────┐
        │                   │                   │
   response is           response is         "interpret
   CONTINUOUS             BINARY               β̂ⱼ"
        │                   │                   │
   does x appear         F8 (why not      does xⱼ appear
   more than once?        linear?)         more than once
        │                    │              or interact?
   ┌────┴────┐           F9 (odds)              │
   NO        YES          F10 (prob.)      ┌─────┴─────┐
   │          │                            NO          YES
  F2      polynomial?                       │           │
           │      │                        F2      is it a dummy?
          YES    interaction               (F4 if dummy)  │
           │      │                                   ┌────┴────┐
          F6     F7                                  YES        NO
                                                   compare to    F6/F7
                                                   reference?
                                                       │
                                                  yes→F4  no→F5
```

**And the two habits that carry through every question in this chapter:**

> 🔑 **If a covariate appears in more than one term — polynomial or interaction — never quote a single coefficient as "the effect." Differentiate first.**
>
> 🔑 **Logit coefficients live on the log-odds scale. Exponentiate for odds; multiply by $\pi(1-\pi)$ for probability. Never read $\hat\beta_j$ directly as a probability change.**
