# Ch 1 — FORMULA DECISION GUIDE

> **The question this file answers:** *"I'm looking at an exam question. Which formula do I reach for, and how did I know?"*
>
> Every entry has the same six fields:
> **① Formula ② USE WHEN ③ DON'T USE WHEN ④ WHY ⑤ 🔍 TRIGGER PHRASES in the question ⑥ ⚠️ Errors**
>
> Chapter 1 has fewer formulas than Chapters 2 and 3 — it's mostly conceptual scaffolding — but the few it has (correlation, $R^2$, the log-transform table, and the model-class decision) show up as quick warm-up marks almost every year, and losing them is entirely avoidable.

---

# ⚡ THE 30-SECOND TRIAGE

| The question says… | Go to |
|---|---|
| "what type of model would you use for $y$?" | **F1** |
| "correlation between $x$ and $y$" | **F2** |
| "$R^2$" in a **simple** regression context | **F3** |
| "interpret a coefficient where $y$ or $x$ (or both) are logged" | **F4** |
| "how many parameters / degrees of freedom" | **F5** |
| "is this still a linear model?" | **F6** |

---

# F1 — Choosing the model class

### ① Method (not a formula)

Look at the **type of $y$** — that alone picks the model family.

| $y$ is… | Model |
|---|---|
| continuous, roughly unbounded | classical linear model (Ch 3) |
| binary, $\{0,1\}$ | logit / probit (Ch 2.3) |
| a count, $\{0,1,2,\dots\}$ | Poisson (Ch 5 — out of scope, know it exists) |
| ordered categories | ordinal regression (Ch 6 — out of scope) |

### ② USE WHEN

- The question describes a dataset and asks "which type of regression model is appropriate, and why?"
- You need to justify *why* linear regression is inappropriate for a binary or count outcome

### ③ DON'T USE WHEN

- ❌ The question is about covariate type, not response type — that never changes the model **class**, only how you build $\boldsymbol{X}$ (dummies, polynomials)

### ④ WHY

Ordinary least squares assumes an unbounded, continuous response with additive, roughly symmetric noise. A bounded or discrete response breaks that structurally, not just numerically — see Ch 2's logit derivation for exactly how.

### ⑤ 🔍 TRIGGERS

> *"the outcome is whether or not the patient recovered"* → logit
> *"the number of customer complaints per month"* → count/Poisson — but for this course, flag it and explain **why linear/logit both don't fit** (unbounded-below problem)
> *"model the price of a house"* → continuous, classical linear model

### ⑥ ⚠️ ERRORS

- Picking a model class based on the **covariates** instead of the **response**
- Forgetting to justify the choice — "because $y$ is binary and probabilities must lie in $[0,1]$" is the sentence that earns the mark, not just naming logit

---

# F2 — Correlation coefficient

### ① Formula

$$
r_{xy} = \frac{\widehat{\text{Cov}}(x,y)}{s_x\,s_y} \in [-1,1]
$$

### ② USE WHEN

- Asked directly for the correlation between two variables
- Converting between $R^2$ and $r$ in a **simple** (one-covariate) regression

### ③ DON'T USE WHEN

- ❌ You want to claim causation — correlation never licenses that word
- ❌ The relationship is **non-linear** — $r$ can be near zero even for a strong non-linear relationship (the Anscombe warning)
- ❌ There are **multiple** covariates — $r$ is pairwise only; use $R^2$ or partial correlations instead

### ④ WHY

$r$ standardises the covariance by both variables' spread, making it scale-free and bounded — that boundedness is exactly what makes "$r=0.87$" comparable across completely different units.

### ⑤ 🔍 TRIGGERS

> *"What is the correlation between age and wage?"* (often given only $R^2$ — recover $r$ via F3)
> *"the scatter plot shows almost no linear pattern, yet you suspect a relationship — explain"* → Anscombe: check for **non-linear** structure before concluding "no relationship"

### ⑥ ⚠️ ERRORS

- Forgetting the **sign** — $r$ takes the same sign as $\hat\beta_1$ in simple regression, and the square root in F3 doesn't supply it automatically
- Treating $r=0$ as "no relationship" instead of "no **linear** relationship"

---

# F3 — $R^2$ in simple regression, and recovering $r$

### ① Formula

$$
R^2 = r_{xy}^2 \qquad\text{(simple/one-covariate regression only)} \qquad\Longrightarrow\qquad r = \text{sign}(\hat\beta_1)\sqrt{R^2}
$$

### ② USE WHEN

- Given $R^2$ from a simple regression and asked for the correlation (or vice versa)
- Sanity-checking a fitted simple regression against a scatter plot

### ③ DON'T USE WHEN

- ❌ **Multiple** regression — there, $R^2 = \text{corr}(y,\hat y)^2$, not the square of any single pairwise correlation
- ❌ The model has **no intercept** — the usual $R^2$ decomposition ($\text{SST}=\text{SSE}+\text{explained SS}$) doesn't hold cleanly without one

### ④ WHY

With one covariate, the fitted line's explanatory power and the linear association between $x$ and $y$ are the same piece of information seen from two formulas — $R^2$ is literally the squared correlation.

### ③ 🔍 TRIGGERS

> *"$R^2=0.038$. What is the correlation between the covariate and the response?"* → $\sqrt{0.038}=0.195$, sign from $\hat\beta_1$'s sign *(this exact number appears in Sheet 3 Ex 1 — verified: $\sqrt{0.038}=0.19494$)*

### ⑥ ⚠️ ERRORS

- Taking the square root and **forgetting the sign** — a negative slope means a negative correlation, even though $\sqrt{R^2}$ is always reported positive by your calculator
- Applying this shortcut in a multiple-regression question — the single most common misuse

---

# F4 — Interpreting logged variables

### ① The table (memorise, don't derive under pressure)

| Model | A one-unit ↑ in $x$ means… |
|---|---|
| $y \sim x$ | $\hat\beta_1$ units change in expected $y$ |
| $\log(y) \sim x$ | approximately $100\hat\beta_1\%$ change in expected $y$ |
| $y \sim \log(x)$ | a **1% increase** in $x$ is associated with $\hat\beta_1/100$ units change in $y$ |
| $\log(y) \sim \log(x)$ | a **1% increase** in $x$ is associated with $\hat\beta_1\%$ change in $y$ — this is an **elasticity** |

### ② USE WHEN

- Either $y$, $x$, or both were transformed with a natural log before fitting
- Right-skewed variables (wage, rent, price, income) — the question describes fitting the model on the log scale

### ③ DON'T USE WHEN

- ❌ Neither variable is logged — just use the plain slope interpretation
- ❌ The transform is something other than $\log$ (e.g. $\sqrt{x}$, $1/x$) — those need their own derivative-based interpretation, not this table

### ④ WHY

For small $\hat\beta_1$, $\log(1+\hat\beta_1)\approx\hat\beta_1$ is the underlying approximation that turns "a $\hat\beta_1$ change in $\log y$" into "roughly a $100\hat\beta_1\%$ change in $y$." It's the same logic that makes $\exp(\hat\beta_j)$ the odds ratio in the logit model (Ch 2) — percentage-scale reasoning keeps reappearing throughout the course precisely because logging removes multiplicative structure and turns it additive.

### ⑤ 🔍 TRIGGERS

> *"the model is fit on $\log(\text{rent})$"* · *"interpret $\hat\beta_1$ in percentage terms"* · *"an elasticity of..."*

### ⑥ ⚠️ ERRORS

- Reporting $\hat\beta_1$ itself as "the percentage change" without multiplying by 100 for the $\log(y)\sim x$ case
- Forgetting that the approximation is only good for **small** $\hat\beta_1$ — for large coefficients the exact statement is $(\exp(\hat\beta_1)-1)\times100\%$, but the linear approximation is what's normally expected unless told otherwise
- Confusing which side ($x$ or $y$) was logged — re-read the model equation before answering

---

# F5 — Counting parameters and degrees of freedom

### ① Formula

$$
p = 1\ (\text{intercept}) + (\#\text{continuous covariates}) + \textstyle\sum(c_m - 1)\ (\text{categoricals}) + (\#\text{interactions})
$$

$$
\text{residual df} = n - p
$$

### ② USE WHEN

- Building any model equation from a verbal description, before doing anything else
- Cross-checking an R output's "degrees of freedom" line against your own covariate count

### ③ DON'T USE WHEN

- ❌ Nothing — do this **every time** you set up a model. It's the single habit that prevents downstream errors in Chapters 2 and 3.

### ④ WHY

Every categorical variable with $c$ levels needs $c-1$ dummy columns (else the design matrix becomes singular — the *dummy variable trap*, Ch 2), and every fitted parameter "uses up" one degree of freedom from the $n$ observations.

### ⑤ 🔍 TRIGGERS

> Any question that hands you a verbal model description ("wage depends on age, education with 5 levels, and health with 2 levels") before asking anything else — count $p$ **first**, on the margin of the page, before reading further.

### ⑥ ⚠️ ERRORS

- 🔴 **The $p$-notation trap**: the book uses $p=k+1$ (parameters, incl. intercept); some past exam papers use $p=k$ (covariates only), causing $t_{n-p}$ vs $t_{n-p-1}$ ambiguity. **Defence: never quote a bare symbol — write "$n$ minus the number of $\beta$'s including the intercept" in free-response answers.**
- Forgetting the intercept counts as one parameter even though it's "automatic"
- Undercounting a categorical variable's dummies by one (using $c$ instead of $c-1$)

---

# F6 — "Is this still a linear model?"

### ① The test (not a formula)

**"Linear" means linear in $\boldsymbol\beta$, not linear in $x$.**

| Model | Still linear (in $\beta$)? |
|---|---|
| $y=\beta_0+\beta_1x+\beta_2x^2+\varepsilon$ | ✅ yes — polynomial |
| $y=\beta_0+\beta_1\log x+\varepsilon$ | ✅ yes |
| $y=\beta_0+\beta_1x_1+\beta_2x_2+\beta_3x_1x_2+\varepsilon$ | ✅ yes — interaction |
| $y=\exp(\boldsymbol{x}'\boldsymbol\beta+\varepsilon)$ | ✅ yes, after taking logs |
| $y=\beta_0+x^{\beta_1}+\varepsilon$ | ❌ no — $\beta_1$ is inside a nonlinear function |

### ② USE WHEN

- Asked to justify whether OLS is still applicable to a transformed or curved-looking model

### ③ DON'T USE WHEN

- ❌ The nonlinearity is in $\boldsymbol\beta$ itself (e.g. $x^{\beta_1}$, $\exp(\beta_1 x)$ where $\beta_1$ sits inside the exponent as a multiplier of a non-constant) — those genuinely require nonlinear least squares, outside this course's scope

### ④ WHY

OLS's entire closed-form solution $(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$ comes from the objective function being **quadratic in $\boldsymbol\beta$**, which requires $\boldsymbol\beta$ to enter only through linear combinations (dot products) inside the model. Once $\boldsymbol\beta$ sits inside a nonlinear transformation of itself, the derivative of the sum of squares is no longer linear in $\boldsymbol\beta$, and there's no closed-form solution.

### ⑤ 🔍 TRIGGERS

> *"is $y=\beta_0+\beta_1x+\beta_2x^2$ a linear model?"* · *"can OLS still be applied here?"*

### ⑥ ⚠️ ERRORS

- Confusing "linear relationship between $x$ and $y$" (a curve is not linear) with "linear model" (a polynomial curve is still a linear model) — these are different meanings of the word "linear" and the exam exploits exactly this ambiguity

---

# 🚦 THE MASTER FLOWCHART

```
                    READ THE QUESTION
                            │
        ┌───────────────────┼───────────────────┐
        │                   │                   │
   "what model         "correlation /       "interpret a
    class...?"          R^2...?"             log coefficient"
        │                   │                   │
       F1              simple reg?             F4
                        │        │
                      YES        NO
                        │        │
                    F2/F3   use R² directly, not r²
```

**And the two habits that carry into every later chapter:**

> 🔑 **Count $p$ — intercept + continuous + $\sum(c-1)$ + interactions — before writing anything else.**
>
> 🔑 **"Linear" means linear in $\boldsymbol\beta$. A curved relationship in $x$ is still a linear model.**
