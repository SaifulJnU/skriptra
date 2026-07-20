# Ch 2 — POSSIBLE TRAPS

> Every trap below is taken from, or directly generalises, a statement in your five past papers. 🔴 marks ones that appeared verbatim.

---

## TRAP 1 🔴 — "Positive coefficient ⟹ positive correlation"

**Paper:** WS 23/24, Block I(i). **Answer: FALSE.**

$\hat\beta_j$ is a **partial** effect (others held fixed). Correlation is a **marginal** association (nothing held fixed). Confounding makes them differ, and they can have **opposite signs**.

**Canonical example:** firefighters vs property damage. Marginal correlation strongly positive; partial effect (controlling for fire size) negative.

⚠️ **The exception that makes this tricky:** in *simple* regression with one covariate, sign$(\hat\beta_1)$ **does** equal sign$(r_{xy})$, because $\hat\beta_1 = r\cdot s_y/s_x$ and $s_y/s_x>0$. The trap statement doesn't say "simple," so it's false in general.

---

## TRAP 2 🔴 — "$m$ categories ⟹ $m$ dummies"

**Paper:** WS 23/24, Block I(iv). **Answer: FALSE.** It's $m-1$.
**Paper:** Linear_model_exam_sheet, Block I(iv): "$k$ levels ⟹ $k-1$ dummies." **Answer: TRUE.**

They test the same fact in both directions. Read the number carefully.

**The one exception:** if you *drop the intercept*, you can include all $m$ dummies. Then each coefficient is that category's **mean** rather than a difference from a reference. Valid, occasionally useful, but not what's being asked unless stated.

---

## TRAP 3 🔴 — Logit coefficient = change in probability

**Paper:** Exam Summer 2025, Ex 1(h). **Answer: FALSE.**

> $\hat\beta_j$ = change in **LOG-ODDS**
> $\exp(\hat\beta_j)$ = multiplicative change in **ODDS**
> $\hat\beta_j\pi(1-\pi)$ = change in **PROBABILITY** — and it's not constant

**Memory hook:** *the logit model is linear in the log-odds. That's the whole point of the link function. If the coefficient were a probability change, you wouldn't need a link at all — you'd be back to the broken linear probability model.*

Only the **sign** transfers between scales.

---

## TRAP 4 — Interpreting a main effect when an interaction is present

$$y=\beta_0+\beta_1x+\beta_2D+\beta_3(xD)$$

❌ "$\hat\beta_1$ is the effect of $x$."
✅ "$\hat\beta_1$ is the effect of $x$ **in the reference group** ($D=0$). The effect of $x$ is $\beta_1+\beta_3D$."

❌ "$\hat\beta_2$ is the effect of $D$."
✅ "$\hat\beta_2$ is the effect of $D$ **at $x=0$**. The effect of $D$ is $\beta_2+\beta_3x$."

> **Sheet 2 makes this vivid:** the health coefficient is $-1.81$, which naively reads as "good health lowers wage." It doesn't — it's the health effect *at age 0*. At age 40 the effect is $-1.81+0.43(40) = +\$15.39$.

**Same trap for polynomials:** in $\beta_0+\beta_1x+\beta_2x^2$, the effect is $\beta_1+2\beta_2x$, never $\beta_1$.

**General rule:** *a variable appearing in more than one term never has its effect given by a single coefficient.*

---

## TRAP 5 🔴 — Rank deficiency and OLS existence

**Paper:** Exam Summer 2025, Ex 1(d): *"When $\boldsymbol{X}$ does not have full column rank, the OLS estimates still exist and are unique as long as the error variance is constant."*
**Answer: FALSE.**

Two things wrong:
1. **Uniqueness requires full column rank.** Without it, $\boldsymbol{X}'\boldsymbol{X}$ is singular, $(\boldsymbol{X}'\boldsymbol{X})^{-1}$ doesn't exist, and the normal equations have **infinitely many** solutions.
2. **Homoscedasticity is completely irrelevant** to whether $\hat{\boldsymbol\beta}$ is identified. The statement dangles an unrelated condition — a classic exam construction. Watch for "as long as…" clauses that have nothing to do with the claim.

**Related rank statements from RCLM WS 22/23:**

| Statement | Book convention ($p$ = parameters) |
|---|---|
| $\text{rk}(\boldsymbol{X}'\boldsymbol{X}) = p$ | ✅ TRUE |
| $\text{rk}(\boldsymbol{X}'\boldsymbol{X}) = k$ | ❌ FALSE (one short — forgot the intercept) |
| $\text{rk}(\boldsymbol{X}'\boldsymbol{X}) = n$ | ❌ FALSE — $\boldsymbol{X}'\boldsymbol{X}$ is $p\times p$; **rank can never exceed the smaller dimension**, and $p\ll n$ |

---

## TRAP 6 🔴 — "Interactions only between continuous × categorical"

**Paper:** RCLM WS 22/23, Block III(iii). **Answer: FALSE.**

Interactions can be formed between **any** two covariates, including **two categorical** ones (product of two dummies) and **two continuous** ones. There is no restriction whatsoever — an interaction is just a product column in $\boldsymbol{X}$.

---

## TRAP 7 — Interpreting the intercept when it's meaningless

Saying "the expected wage at age 0 is \$81.70" is arithmetically fine and judgementally wrong. Always add: *"however, age 0 is outside the data range and not meaningful, so $\hat\beta_0$ should not be substantively interpreted."*

**Sheet 1 Ex 1(c) asks this directly.** It's a whole sub-question. Don't fumble it.

---

## TRAP 8 — Forgetting "holding all other covariates fixed"

In multiple regression this phrase is often **explicitly worth part of a mark**. Writing "a one-year increase in age raises wage by \$0.62" without it describes a *marginal* relationship, which is not what $\hat\beta_j$ estimates.

Say it every single time. It costs four words.

---

## TRAP 9 — Using the raw dummy coefficient for a non-reference comparison

❌ "An Advanced Degree holder earns \$62.63 more than a HS Grad."
✅ "…than someone with **less than a high school degree** (the reference)."
✅ Advanced vs HS Grad $= 62.63-11.01 = \$51.62$.

**Always name what you're comparing to.**

---

## TRAP 10 — Thinking the linear probability model is simply "not allowed"

A subtler point worth a bonus mark: the LPM is not *illegal*, it's *inadequate*. It's occasionally used (with robust standard errors) as a quick approximation when all fitted probabilities happen to land safely inside $[0,1]$. But it cannot **guarantee** valid probabilities, its errors are necessarily heteroscedastic and non-normal, and its constant marginal effects are substantively wrong near the boundaries.

Frame your answer as *"inappropriate because…"* rather than *"forbidden."* More accurate, and it reads as understanding rather than recitation.

---

## TRAP 11 — Confusing logit with probit, or with log-linear

Three different things that all involve logs or S-curves:

| Model | Response | Equation |
|---|---|---|
| **Logit** | binary | $\log\frac{\pi}{1-\pi} = \boldsymbol{x}'\boldsymbol\beta$ |
| **Probit** | binary | $\Phi^{-1}(\pi) = \boldsymbol{x}'\boldsymbol\beta$ |
| **Log-linear** | **continuous, positive** | $\log(y) = \boldsymbol{x}'\boldsymbol\beta+\varepsilon$ |

The log-linear model is an **ordinary linear model** fitted by OLS with a transformed response. The logit is a **GLM** fitted by maximum likelihood. They are not related, despite both having "log" in them. Don't let the word blur them.

---

## TRAP 12 — "Linear model" excludes curves

Repeating from Chapter 1 because it recurs here in a new costume: $y = \beta_0+\beta_1\text{age}+\beta_2\text{age}^2+\varepsilon$ **is** a linear model. Linear in $\boldsymbol\beta$. Estimated by ordinary OLS. The *fitted curve* is a parabola; the *model* is linear.

---

## TRAP 13 — Miscounting parameters, then miscounting df

One miscounted categorical variable cascades through the whole paper:

```
wrong number of dummies
   → wrong p
      → wrong residual df (n − p)
         → wrong t quantile
            → wrong confidence interval
               → wrong test decision
```

**Defence:** cross-check against the R output. If it says "on 2983 degrees of freedom" and $n=3000$, then $p=17$. Count your terms and make them agree **before** you use any quantile.

---

## Rapid-fire checklist (read the morning of the exam)

- [ ] $c$ levels ⟹ $c-1$ dummies
- [ ] Reference = the level **missing** from the output
- [ ] Non-reference comparison ⟹ **subtract** coefficients
- [ ] Variable in two terms ⟹ **differentiate**, don't quote one coefficient
- [ ] Interaction ⟹ **non-parallel** lines; split into two rows
- [ ] Logit $\hat\beta$ = **log-odds**; $\exp(\hat\beta)$ = **odds ratio**; probability effect isn't constant
- [ ] Full rank is required for uniqueness — homoscedasticity is irrelevant to it
- [ ] Interactions work between **any** covariate types
- [ ] Don't interpret $\hat\beta_0$ if $x=0$ is meaningless
- [ ] Say **"holding all other covariates fixed"**
- [ ] Log-linear ≠ logit
- [ ] Cross-check $p$ against the R output's residual df
