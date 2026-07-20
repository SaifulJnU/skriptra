# 1.1 — Examples of Applications

> **Purpose of this section:** to convince you that one framework — regression — answers a startling range of questions, and to give you the vocabulary those questions are phrased in.

---

## 1. The origin: Galton, 1885

Sir Francis Galton collected the heights of 928 adult children and their 205 sets of parents. (He multiplied women's heights by 1.08 to make them comparable to men's — an early example of a *transformation*, which you'll meet properly in 3.1.2.)

He asked: **how does a child's height depend on the parents' average height?**

Two facts he noticed, and both matter:

1. **The relationship is real but not deterministic.** Tall parents tend to have tall children — but not always, and not by a fixed amount. There is a *systematic* part and a *random* part.
2. **The relationship pulls toward the average.** Very tall parents have children who are tall, but *less* tall than the parents. Galton called this "regression toward mediocrity." The name stuck and now describes the entire field, which is a historical accident — modern regression has nothing to do with regressing toward the mean.

Galton wrote the first regression model:

$$y = \beta_0 + \beta_1 x + \varepsilon$$

- $y$ = child's height (**response variable**, also called dependent variable, target, regressand)
- $x$ = parents' average height (**covariate**, also called explanatory variable, independent variable, predictor, regressor)
- $\beta_0 + \beta_1 x$ = the **systematic component** — the part explained by $x$
- $\varepsilon$ = the **random error** / **noise** — everything else

Galton eyeballed $\beta_0$ and $\beta_1$. We use **least squares**: choose the $\beta$'s that minimise

$$\sum_{i=1}^{n}(y_i - \beta_0 - \beta_1 x_i)^2$$

That sum is the only thing you need to remember from this section.

---

## 2. The application zoo

The book runs through several datasets. You need the *pattern*, not the data.

| Application | Response $y$ | Type of $y$ | Typical covariates | Model class |
|---|---|---|---|---|
| **Munich rent index** | net rent (€) | continuous | living area, year built, location, kitchen quality | linear model (Ch 3) |
| **Malnutrition in Zambia** | child's Z-score | continuous | age, mother's BMI, region | additive / geoadditive (Ch 9) |
| **Credit scoring** | default: yes/no | **binary** | loan amount, duration, account status | **logit / probit (Ch 2.3, Ch 5)** |
| **Patent opposition** | opposed: yes/no | binary | patent characteristics | logit |
| **Insurance claims** | number of claims | **count** | driver age, region | Poisson (Ch 5) |
| **Forest health** | damage level (ordinal) | ordered categorical | tree age, soil, canopy | ordinal models (Ch 6) |

### The one rule this table teaches

> **The type of the response variable decides the model class.**

- Continuous, roughly symmetric → **classical linear model**
- Binary (0/1) → **logit or probit**
- Count (0, 1, 2, 3, …) → **Poisson regression**
- Ordered categories → ordinal regression

Covariate types matter for *how you put them in the model* (Section 3.1.3), but they never change the model class. Only $y$ does that.

**Exam relevance:** this rule is directly examinable. Exam Summer 2025, Exercise 4(a) asks *"Explain why a linear regression model is not appropriate to model a binary dependent variable."* You are answering this table.

---

## 3. What questions can regression answer?

Three distinct goals, and the book is careful to separate them because they have different consequences:

### (a) Description / association
*"Is there a relationship between living area and rent, and how strong?"*
You care about the sign, size and significance of individual $\beta_j$'s.

### (b) Prediction
*"What rent should I expect for a 70 m² flat built in 1975?"*
You care about $\hat{y}$ being close to $y$ on **new** data. You don't care which covariate did the work. This goal drives Chapter 3.4 (model choice, AIC, cross-validation).

### (c) Causal effect
*"If I increase education by one level, what happens to wage?"*
This is the hardest, and regression alone rarely delivers it — you need the covariate to be as-good-as-randomly assigned. The book is careful here and so should you be.

> **Exam-safe phrasing:** when asked to interpret a coefficient, say *"holding all other covariates fixed, a one-unit increase in $x_j$ is associated with a change of $\hat\beta_j$ in the expected value of $y$."* Say **"associated with"**, not **"causes"**. Free marks for the right verb.

---

## 4. Vocabulary you must own after this section

| Term | Meaning | Synonyms you'll see |
|---|---|---|
| **Response variable** | what we're explaining, $y$ | dependent variable, regressand, target, outcome |
| **Covariate** | what we explain it with, $x$ | explanatory variable, independent variable, regressor, predictor, feature |
| **Systematic component** | $f(x)$ — the structured part | signal, mean function, predictor (linear predictor $\eta$) |
| **Random error** | $\varepsilon$ — the unexplained part | noise, disturbance, error term |
| **Observation / unit** | one row of data, indexed $i = 1,\dots,n$ | case, individual, subject |
| **Fitted value** | $\hat{y}_i$ — model's prediction for observation $i$ | prediction |
| **Residual** | $\hat\varepsilon_i = y_i - \hat{y}_i$ | — |
| **Parameter** | $\beta_j$ — unknown, to be estimated | coefficient, regression coefficient |
| **Estimate** | $\hat\beta_j$ — the number we computed from data | — |

**Critical distinction, examinable indirectly everywhere:**

> $\varepsilon_i$ is the **true, unobservable error**. $\hat\varepsilon_i$ is the **observable residual**.
> $\beta_j$ is the **true, unknown parameter**. $\hat\beta_j$ is the **estimate**.
>
> Assumptions are made about $\varepsilon$ and $\beta$ (the truth). Diagnostics are done on $\hat\varepsilon$ (what we can see). The whole of Chapter 3.4.4 is: *we can't check assumptions about $\varepsilon$ directly, so we check $\hat\varepsilon$ and hope.*

Every hat means "estimated from data." Students who lose track of hats lose marks in Chapter 3.

---

## 5. Why "linear" is much less restrictive than it sounds

A pre-emptive warning, because this catches people:

**"Linear model" means linear in the *parameters* $\beta$, not linear in the covariates $x$.**

All of these are linear models:

$$y = \beta_0 + \beta_1 x + \beta_2 x^2 + \varepsilon \quad \text{(quadratic in } x\text{, linear in } \beta)$$
$$y = \beta_0 + \beta_1 \log(x) + \varepsilon$$
$$y = \beta_0 + \beta_1 x_1 + \beta_2 x_2 + \beta_3 x_1 x_2 + \varepsilon \quad \text{(interaction)}$$

This one is **not**:

$$y = \beta_0 + x^{\beta_1} + \varepsilon$$

because $\beta_1$ sits in an exponent.

And this one is **not linear as written**, but *becomes* linear after a log transformation:

$$y = \exp(\beta_0 + \beta_1 x_1 + \dots + \beta_k x_k + \varepsilon) \;\Longrightarrow\; \log(y) = \beta_0 + \beta_1 x_1 + \dots + \beta_k x_k + \varepsilon$$

> ⚠️ **This exact multiplicative model appeared as a TRUE/FALSE statement in the WS 22/23 exam**: *"The relation y = exp(β₀ + β₁x₁ + ⋯ + βₖxₖ + ε) cannot be analysed within the linear regression framework."* → **FALSE**. Take logs and it's an ordinary linear model. See `32-TRAPS.md`.

---

## 6. Key takeaways

1. Regression = separating **systematic** from **random**: $y = f(x) + \varepsilon$.
2. Parameters are chosen by **least squares** — minimising the sum of squared deviations.
3. **The type of $y$ chooses the model class.** Continuous → linear; binary → logit; count → Poisson.
4. Three goals — description, prediction, causality — and they are not the same.
5. **"Linear" refers to the parameters.** Squares, logs and interactions are all fair game.
6. Hats matter: $\varepsilon$ vs $\hat\varepsilon$, $\beta$ vs $\hat\beta$.
