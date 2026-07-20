# 2.3 — Regression with Binary Response: The Logit Model

> **Guaranteed exam content.** Every past paper touches this. But the demands are narrow: you need **two answers word-perfect** and nothing more. No derivations, no estimation, no maximum likelihood.
>
> **Answer 1:** why a linear model fails for binary $y$.
> **Answer 2:** what $\hat\beta_j$ means in a logit model — and what it does *not* mean.

---

## 1. The setting

$y_i \in \{0,1\}$. Examples: credit default (yes/no), patent opposed (yes/no), disease present (yes/no), passed exam (yes/no).

For a binary variable, the mean **is** the probability:

$$E(y_i\mid\boldsymbol{x}_i) = 1\cdot P(y_i=1\mid\boldsymbol{x}_i) + 0\cdot P(y_i=0\mid\boldsymbol{x}_i) = P(y_i=1\mid\boldsymbol{x}_i) =: \pi_i$$

So regression on a binary response is **modelling a probability**. And a probability must lie in $[0,1]$.

This single fact generates everything below.

---

## 2. 🎯 ANSWER 1 — Why the linear model fails (Exam Summer 2025, Ex 4(a), 1 point)

The naive **linear probability model** would be

$$P(y_i=1\mid\boldsymbol{x}_i) = \beta_0 + \beta_1x_{i1}+\dots+\beta_kx_{ik} = \boldsymbol{x}_i'\boldsymbol\beta$$

There are **four** things wrong with it. Memorise at least three; give them in this order.

### (1) 🔴 Predictions fall outside $[0,1]$ — the fatal one

$\boldsymbol{x}'\boldsymbol\beta$ is a straight line. Lines are unbounded: they run to $+\infty$ and $-\infty$. So for extreme covariate values the model predicts probabilities like $1.34$ or $-0.20$.

**A probability cannot be 1.34. The model is producing impossible numbers.** Lead with this one — it is the answer the marker is looking for.

### (2) Heteroscedasticity is guaranteed, not accidental

If $y_i \sim \text{Bernoulli}(\pi_i)$, then

$$\text{Var}(y_i) = \pi_i(1-\pi_i)$$

Since $\pi_i$ depends on $\boldsymbol{x}_i$, so does the variance. The homoscedasticity assumption $\text{Var}(\varepsilon_i)=\sigma^2$ is **violated by construction** — you cannot fix it by choosing better covariates.

*(Note the variance is largest at $\pi = 0.5$, where it is $0.25$, and shrinks to 0 at the extremes.)*

### (3) The errors cannot possibly be normal

For a given $\boldsymbol{x}_i$, $y_i$ takes only two values, so $\varepsilon_i = y_i - \boldsymbol{x}_i'\boldsymbol\beta$ takes only **two values**:

$$\varepsilon_i = \begin{cases} 1 - \boldsymbol{x}_i'\boldsymbol\beta & \text{with prob } \pi_i\\ -\boldsymbol{x}_i'\boldsymbol\beta & \text{with prob } 1-\pi_i\end{cases}$$

A two-point distribution is about as far from normal as it gets. So the exact $t$- and $F$-tests have no basis.

### (4) A constant marginal effect is implausible

The linear model says: a one-unit increase in $x_j$ changes the probability by $\beta_j$, **always** — whether you're at $\pi = 0.5$ or $\pi = 0.99$.

That's substantively wrong. Moving someone's default probability from 0.50 to 0.55 is easy; from 0.99 to 1.04 is impossible. **Effects should shrink as you approach the boundaries.**

### The fix

Wrap the linear predictor in a function that squashes $\mathbb{R}$ into $(0,1)$:

$$\pi_i = h(\eta_i) = h(\boldsymbol{x}_i'\boldsymbol\beta), \qquad h : \mathbb{R}\to(0,1)$$

**Logit** uses the logistic CDF; **probit** uses the standard normal CDF. Both fix all four problems. Both are estimated by **maximum likelihood**, not OLS.

### ✍️ Model exam answer (write approximately this)

> A linear model is inappropriate for a binary response because $E(y_i\mid\boldsymbol{x}_i) = P(y_i=1\mid\boldsymbol{x}_i)$ is a **probability** and must lie in $[0,1]$, whereas the linear predictor $\boldsymbol{x}_i'\boldsymbol\beta$ is unbounded and will produce fitted values below 0 or above 1. In addition, a Bernoulli response has $\text{Var}(y_i)=\pi_i(1-\pi_i)$, which depends on $\boldsymbol{x}_i$, so the **homoscedasticity assumption is necessarily violated**, and the errors take only two values and therefore **cannot be normally distributed**, invalidating the usual $t$- and $F$-tests.
>
> The logit (or probit) model solves this by modelling $P(y_i=1) = h(\boldsymbol{x}_i'\boldsymbol\beta)$ with $h$ a strictly increasing function mapping $\mathbb{R}$ onto $(0,1)$ — the logistic CDF for logit, the standard normal CDF for probit. Fitted probabilities are then guaranteed to be valid, and marginal effects automatically shrink near 0 and 1. Estimation is by maximum likelihood.

---

## 3. The logit model — three equivalent forms

**You must be able to write all three and move between them.**

### Form A — the probability (response scale)

$$\boxed{\;\pi_i = P(y_i = 1\mid\boldsymbol{x}_i) = \frac{\exp(\boldsymbol{x}_i'\boldsymbol\beta)}{1+\exp(\boldsymbol{x}_i'\boldsymbol\beta)} = \frac{1}{1+\exp(-\boldsymbol{x}_i'\boldsymbol\beta)}\;}$$

This is the **logistic function**. Its graph is the S-curve (sigmoid): flat near 0 on the far left, steep in the middle, flat near 1 on the far right. It never touches 0 or 1.

### Form B — the odds

$$\frac{\pi_i}{1-\pi_i} = \exp(\boldsymbol{x}_i'\boldsymbol\beta) = \exp(\beta_0)\cdot\exp(\beta_1x_{i1})\cdots\exp(\beta_kx_{ik})$$

**Odds** = (probability of success)/(probability of failure). $\pi = 0.75$ ⟹ odds $= 3$, "3 to 1 on."

Notice: on the odds scale the model is **multiplicative**.

### Form C — the log-odds / logit (link scale) ⭐

$$\boxed{\;\text{logit}(\pi_i) = \log\!\left(\frac{\pi_i}{1-\pi_i}\right) = \beta_0 + \beta_1x_{i1}+\dots+\beta_kx_{ik}\;}$$

**This is the form that makes it a "generalised *linear* model."** The model is linear — just not in $\pi$, in $\log(\pi/(1-\pi))$.

The function $\text{logit}(\pi) = \log\frac{\pi}{1-\pi}$ is called the **link function**. It maps $(0,1)\to\mathbb{R}$; the logistic function is its inverse, mapping back.

$$\pi \in (0,1) \xrightarrow{\ \text{logit}\ } \eta \in \mathbb{R} \xrightarrow{\ \text{logistic}\ } \pi \in (0,1)$$

> **Why the exam cares:** your ChatGPT notes flagged $\log\frac{p}{1-p}$ as the reason you need logarithms. This is it. Form C is the only place logs are structurally required in this course.

---

## 4. 🎯 ANSWER 2 — Interpreting $\hat\beta_j$ (Exam Summer 2025, Ex 1(h))

> 🔴 **The exact past-paper statement:**
> *"Consider a logit model where you regress $y_i$ onto $1, x_{1,i},\dots,x_{k,i}$ to obtain $\hat\beta_0,\dots,\hat\beta_k$. An increase of 1 in $x_{j,i}$ is interpreted as an increase of $\hat\beta_j$ in $P(y_i=1)$."*
>
> ## **FALSE.**

### Why it's false

$\hat\beta_j$ is a change on the **log-odds** scale, **not** on the probability scale.

$$\text{logit}(\pi) = \beta_0+\beta_1x_1+\dots \;\Longrightarrow\; \frac{\partial\,\text{logit}(\pi)}{\partial x_j} = \beta_j$$

The effect **on the probability itself** is

$$\frac{\partial\pi}{\partial x_j} = \beta_j\,\pi(1-\pi)$$

which **depends on $\pi$**, hence on all the covariates. It is largest when $\pi = 0.5$ (where it equals $0.25\beta_j$) and shrinks toward 0 as $\pi$ approaches 0 or 1. **There is no single number describing "the effect on the probability."** That's precisely the flexibility the logit model was built to provide.

### The three correct interpretations

| Scale | Statement |
|---|---|
| **Log-odds** (exact, linear) | A one-unit increase in $x_j$ increases the **log-odds** of $y=1$ by $\hat\beta_j$, holding other covariates fixed. |
| **Odds** (exact, multiplicative) ⭐ | A one-unit increase in $x_j$ **multiplies the odds** of $y=1$ by $\exp(\hat\beta_j)$, holding other covariates fixed. |
| **Probability** (not constant) | The effect on $P(y=1)$ is $\hat\beta_j\pi(1-\pi)$ — it **depends on where you are** and cannot be summarised by one number. |

**$\exp(\hat\beta_j)$ is called the odds ratio.** It's the most quotable interpretation:

| $\hat\beta_j$ | $\exp(\hat\beta_j)$ | Meaning |
|---|---|---|
| $0$ | $1$ | no effect — odds unchanged |
| $+0.69$ | $\approx 2$ | odds **double** |
| $-0.69$ | $\approx 0.5$ | odds **halve** |
| $+1$ | $2.72$ | odds multiply by 2.72 |

### The sign *is* reliable

One thing that *does* carry over: $\hat\beta_j > 0$ means increasing $x_j$ **increases** $P(y=1)$, and $\hat\beta_j<0$ decreases it. Because the logistic function is strictly increasing, the sign is unambiguous. Only the *magnitude* fails to transfer.

### ✍️ Model exam answer

> **FALSE.** In a logit model the linear predictor equals the **log-odds**, not the probability: $\log\frac{\pi_i}{1-\pi_i} = \boldsymbol{x}_i'\boldsymbol\beta$. Hence a one-unit increase in $x_{j,i}$ increases the log-odds by $\hat\beta_j$, equivalently **multiplies the odds by $\exp(\hat\beta_j)$**. The effect on the probability itself is $\hat\beta_j\pi_i(1-\pi_i)$, which depends on $\pi_i$ and therefore on all covariate values, so it is **not** a constant increase of $\hat\beta_j$.

---

## 5. Probit — one paragraph is enough

$$P(y_i=1) = \Phi(\boldsymbol{x}_i'\boldsymbol\beta)$$

where $\Phi$ is the standard normal CDF. Same S-shape, same guarantees, same motivation.

| | Logit | Probit |
|---|---|---|
| Link | $\log\frac{\pi}{1-\pi}$ | $\Phi^{-1}(\pi)$ |
| Inverse | logistic CDF | normal CDF $\Phi$ |
| Tails | slightly heavier | slightly lighter |
| Interpretation | **odds ratios** — clean | no clean equivalent |
| Coefficients | ≈ $1.6\times$ probit's | — |

**Fitted probabilities from the two are nearly indistinguishable in practice.** Logit is preferred mainly because $\exp(\hat\beta_j)$ has the clean odds-ratio meaning.

---

## 6. Estimation — what you need to say and no more

**OLS is not used.** There is no closed-form solution. The logit model is fitted by **maximum likelihood**, maximising

$$L(\boldsymbol\beta) = \prod_{i=1}^n \pi_i^{y_i}(1-\pi_i)^{1-y_i}$$

numerically (Newton–Raphson / IRLS). Inference uses **asymptotic** normality of the ML estimator, so tests are $z$-tests and likelihood-ratio / Wald / score tests rather than exact $t$- and $F$-tests.

That paragraph is the complete extent of what your reading course requires. Do not go further — Chapter 5 is out of scope.

---

## 7. Key takeaways

1. Binary $y$ ⟹ $E(y\mid\boldsymbol{x}) = P(y=1) = \pi$ ⟹ **must lie in $[0,1]$.**
2. **Four reasons linear fails:** predictions outside $[0,1]$ (the fatal one) · guaranteed heteroscedasticity $\pi(1-\pi)$ · two-valued, non-normal errors · implausibly constant marginal effects.
3. **Three forms:** $\pi = \frac{e^{\eta}}{1+e^\eta}$ · odds $=e^\eta$ · $\log\frac{\pi}{1-\pi} = \eta$.
4. 🔴 **$\hat\beta_j$ is a change in LOG-ODDS, not in probability.** $\exp(\hat\beta_j)$ = **odds ratio**.
5. Effect on probability is $\hat\beta_j\pi(1-\pi)$ — non-constant. Only the **sign** transfers reliably.
6. Probit ≈ logit; logit wins on interpretability.
7. Estimated by **maximum likelihood**, not OLS.
