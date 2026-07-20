# 2.1 — Introduction: what a regression model actually specifies

> **Purpose:** to state precisely what a regression model claims, before any particular model is chosen. This section is short but it fixes the framing for everything that follows.

---

## 1. The general regression problem

You have:
- a **response** $y$
- **covariates** $x_1, \dots, x_k$, collected as $\boldsymbol{x} = (1, x_1,\dots,x_k)'$
- $n$ observations $(y_i, \boldsymbol{x}_i)$, $i=1,\dots,n$

You want to describe how the **distribution of $y$ depends on $\boldsymbol{x}$**.

Read that carefully. Not "the value of $y$" — the **distribution** of $y$. For a given $\boldsymbol{x}$ there is not one $y$, there is a whole spread of possible $y$'s. A 40-year-old college graduate does not have *a* wage; there is a distribution of wages among such people.

---

## 2. 🔑 The central idea: regression models the conditional mean

Almost all of this course models **one feature** of that conditional distribution: its **mean**.

$$\boxed{\;E(y \mid \boldsymbol{x}) = f(\boldsymbol{x})\;}$$

Combined with additive errors:

$$y = f(\boldsymbol{x}) + \varepsilon, \qquad E(\varepsilon\mid\boldsymbol{x}) = 0$$

These two statements are **equivalent**. Take expectations of the second:

$$E(y\mid\boldsymbol{x}) = E(f(\boldsymbol{x})\mid\boldsymbol{x}) + E(\varepsilon\mid\boldsymbol{x}) = f(\boldsymbol{x}) + 0 = f(\boldsymbol{x})$$

> **This equivalence is worth internalising.** The assumption "$E(\varepsilon) = 0$" is not an extra restriction you're hoping is true — it is *what makes $f(\boldsymbol{x})$ the mean function*. If $E(\varepsilon) = 3$, then $f(\boldsymbol{x})$ isn't the mean, and you could just move the 3 into $\beta_0$. **In a model with an intercept, $E(\varepsilon)=0$ is essentially free**, because the intercept absorbs any constant shift.

### Why this matters for interpretation

Every coefficient interpretation in this course contains the word **"expected"** or **"on average"** for this reason:

> *"A one-year increase in age is associated with a \$0.71 increase in the **expected** hourly wage."*

Not "your wage goes up by \$0.71." The model says nothing about *your* wage. It says something about the **average** wage of people like you.

Leaving out "expected"/"on average" is a small thing that markers notice.

---

## 3. What the model does *not* specify (yet)

$E(y\mid\boldsymbol{x}) = f(\boldsymbol{x})$ pins down the mean. It says nothing about:

| Feature | Where it gets specified |
|---|---|
| The **variance** of $y$ given $\boldsymbol{x}$ | Homoscedasticity assumption: $\text{Var}(\varepsilon_i) = \sigma^2$ (Ch 3.1) |
| The **shape** of the conditional distribution | Normality assumption: $\varepsilon \sim N(0,\sigma^2)$ (Ch 3.1) |
| **Dependence** between observations | Independence assumption (Ch 3.1) |
| The **functional form** of $f$ | This chapter (linear? logit? etc.) |

**This is a useful map of the whole course.** Chapter 2 chooses the *shape of $f$*. Chapter 3.1 nails down the assumptions on $\varepsilon$. Chapters 3.2–3.3 exploit those assumptions. Chapter 3.4 checks whether they held.

> **Beyond mean regression:** the book notes (2.9) that you could model other features — quantile regression models the conditional *median* or other quantiles; GAMLSS models the variance too. Out of scope, but worth one sentence of awareness: **"classical regression models the conditional mean"** is a *choice*, not a law.

---

## 4. Two decisions define any regression model

Every model in this book — and every model in the exam — is fully specified by answering two questions:

### Decision 1: What is the shape of $f$?

| Shape | Model | Where |
|---|---|---|
| $\boldsymbol{x}'\boldsymbol\beta$, a straight sum | **linear model** | 2.2, Ch 3 |
| $h(\boldsymbol{x}'\boldsymbol\beta)$ for some squashing function $h$ | **generalised linear model** (incl. logit) | 2.3, Ch 5 |
| $\beta_0 + f_1(x_1) + \dots + f_k(x_k)$, smooth unknown functions | additive model | 2.6 — out of scope |

Note that even the logit model has $\boldsymbol{x}'\boldsymbol\beta$ hiding inside it. That inner quantity gets its own name:

$$\eta = \boldsymbol{x}'\boldsymbol\beta = \beta_0 + \beta_1x_1 + \dots + \beta_kx_k \qquad \textbf{the linear predictor}$$

**The linear predictor is the common skeleton of every model in this course.** Linear model: $E(y) = \eta$. Logit model: $P(y=1) = h(\eta)$ where $h$ squashes into $[0,1]$. Same skeleton, different wrapper.

### Decision 2: What kind of thing is $y$?

Covered in Chapter 1, restated because it drives everything:

| $y$ | Model | $E(y\mid\boldsymbol{x})$ must live in |
|---|---|---|
| continuous | linear | $(-\infty,\infty)$ — no constraint |
| binary | **logit/probit** | $[0,1]$ — it's a probability! |
| count | Poisson | $[0,\infty)$ — non-negative |

That third column is the whole reason non-linear models exist. **If the mean of $y$ is constrained, the model must respect the constraint.** $\boldsymbol{x}'\boldsymbol\beta$ respects no constraints — it ranges over all of $\mathbb{R}$. So for binary $y$ you must wrap it in something that squashes. That's Section 2.3, and that's the answer to "why not linear for binary $y$."

---

## 5. Systematic vs random, restated precisely

$$\underbrace{y}_{\text{observed}} = \underbrace{f(\boldsymbol{x})}_{\text{systematic}} + \underbrace{\varepsilon}_{\text{random}}$$

| | Systematic component | Random component |
|---|---|---|
| Symbol | $f(\boldsymbol{x})$, $\eta$, $\boldsymbol{x}'\boldsymbol\beta$ | $\varepsilon$ |
| Depends on covariates? | Yes | No (that's the point) |
| Same for two people with identical $\boldsymbol{x}$? | **Yes, identical** | No, differs |
| What we estimate | $\boldsymbol\beta$ | $\sigma^2$ (its spread) |
| Reducible? | — | Partly: add covariates. Never fully. |

> **The sharpest way to say it:** two individuals with *exactly the same covariate values* have exactly the same systematic component. Any difference in their observed $y$ is, by definition, error. This is why $\sigma^2$ measures "how much people who look identical to the model still differ."

---

## 6. Key takeaways

1. **Regression models the conditional mean:** $E(y\mid\boldsymbol{x}) = f(\boldsymbol{x})$.
2. $y = f(\boldsymbol{x}) + \varepsilon$ with $E(\varepsilon)=0$ is an equivalent statement — and with an intercept in the model, $E(\varepsilon)=0$ costs nothing.
3. Therefore every interpretation says **"expected"** or **"on average."**
4. The **linear predictor** $\eta = \boldsymbol{x}'\boldsymbol\beta$ is the skeleton of every model here. Linear model uses it raw; logit wraps it in a squashing function.
5. A model is defined by two decisions: **shape of $f$** and **type of $y$**.
6. If the mean of $y$ is constrained (a probability, a count), the model must be built to respect that constraint. This is the entire motivation for Section 2.3.
