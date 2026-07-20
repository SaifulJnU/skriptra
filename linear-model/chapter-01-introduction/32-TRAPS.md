# Ch 1 — POSSIBLE TRAPS

> These are the *specific* ways marks are lost. Several are lifted verbatim from your past papers. Read this file again on Day 15 and Day 20.

---

## TRAP 1 — "Linear" means linear in $x$ ❌

**The trap statement:** *"A model containing $x^2$ is not a linear model."*

**Reality:** linear refers to the **parameters**. $y = \beta_0 + \beta_1x + \beta_2x^2 + \varepsilon$ is a perfectly ordinary linear model — set $x_2 := x^2$ and estimate it by OLS exactly as usual.

**Test:** if you can write the model as $\boldsymbol{y} = \boldsymbol{X}\boldsymbol\beta + \boldsymbol\varepsilon$ for *some* design matrix $\boldsymbol{X}$ built from known functions of the covariates, it is a linear model.

❌ Not linear: $y = \beta_0 + x^{\beta_1} + \varepsilon$ · $y = \beta_0 + \frac{1}{\beta_1 + x} + \varepsilon$

---

## TRAP 2 — The multiplicative model 🔴 *appeared verbatim in RCLM WS 22/23*

**Statement:** *"The relation $y = \exp(\beta_0 + \beta_1x_1 + \dots + \beta_kx_k + \varepsilon)$ cannot be analysed within the linear regression framework."*

**Answer: FALSE.** Take logs:
$$\log(y) = \beta_0 + \beta_1x_1 + \dots + \beta_kx_k + \varepsilon$$
An ordinary linear model with response $\log y$.

**The subtlety worth a bonus mark:** notice the $\varepsilon$ is *inside* the exponential. That's what makes the log trick work. If the model were $y = \exp(\boldsymbol{x}'\boldsymbol\beta) + \varepsilon$ (additive error **outside**), logging would **not** linearise it — you'd need non-linear least squares. Read where the error sits.

---

## TRAP 3 — Forgetting the intercept column in $\boldsymbol{X}$

**The trap:** stating $\boldsymbol{X}$ is $n \times k$.

**Reality:** $\boldsymbol{X}$ is $n \times (k+1)$. The first column is all ones.

**Downstream damage:** if you get this wrong, you also get wrong: $\text{rank}(\boldsymbol{X})$, the dimension of $(\boldsymbol{X}'\boldsymbol{X})^{-1}$, and every degrees-of-freedom calculation in the paper. One error, five lost marks.

**Related past-paper statements:**
- *"$\text{rk}(X'X) = p$"* — TRUE in book notation ($p$ = parameters), FALSE if $p$ = covariates.
- *"$\text{rk}(X'X) = k$"* — FALSE (that's one short).
- *"$\text{rk}(X'X) = n$"* — FALSE. $\boldsymbol{X}'\boldsymbol{X}$ is $p\times p$, so its rank is at most $p$, and $p \ll n$ normally. **A matrix's rank cannot exceed its smaller dimension.**

---

## TRAP 4 — The `p` convention 🔴 *the biggest one in this course*

Covered fully in `03-notes-1.3-notation.md`, repeated here because it matters:

| | Book (Fahrmeir) | Some exam papers |
|---|---|---|
| $p$ means | parameters = $k+1$ | covariates = $k$ |
| $t$-test df | $n-p$ | $n-p-1$ |
| $\text{rank}(\boldsymbol{X}'\boldsymbol{X})$ | $p$ | $p+1$ |

**Evidence from your own answer keys:**
- *Linear_model_exam_sheet* Block I: "rank$(X'X) = p$, with the number of variables $p$" → **FALSE**, and "$t_j \sim t_{n-p-1}$" → **TRUE**. That paper uses $p$ = covariates.
- *RCLMWS2223* Block I: "rk$(X'X) = p$" bare → **TRUE** in book notation.

**Defence:** read the paper's own gloss ("with the number of variables $p$"). If there is no gloss, assume the **book convention** ($p$ = parameters), because that's the course textbook. And in free-response answers, write $n - k - 1$ — unambiguous in both worlds.

---

## TRAP 5 — Confusing $\varepsilon$ with $\hat\varepsilon$

| | $\varepsilon_i$ | $\hat\varepsilon_i$ |
|---|---|---|
| Definition | $y_i - \boldsymbol{x}_i'\boldsymbol\beta$ | $y_i - \boldsymbol{x}_i'\hat{\boldsymbol\beta}$ |
| Observable? | **No** | **Yes** |
| Independent? | Assumed yes | **No** — residuals are correlated with each other |
| Constant variance? | Assumed yes ($\sigma^2$) | **No** — $\text{Var}(\hat\varepsilon_i) = \sigma^2(1-h_{ii})$ |
| Sum to zero? | Not necessarily | **Yes**, if the model has an intercept |

**The consequence that gets examined:** *"Since the residuals are normally distributed, standardized residuals are also normally distributed."* (RCLM WS 22/23, Block III(i)) — the whole reason standardisation is needed is that raw residuals have **unequal** variances $\sigma^2(1-h_{ii})$. See Chapter 3.4.4 for the full treatment.

---

## TRAP 6 — Correlation traps

**6a. "Zero correlation implies independence."** FALSE. Zero correlation = no *linear* association. $Y = X^2$ is the standard counterexample.
*(Independence ⟹ zero correlation is true. The converse is not.)*

**6b. "Positive coefficient implies positive correlation."** 🔴 **FALSE — and this appeared in WS 23/24, Block I(i).**
In **multiple** regression, $\hat\beta_j$ is a *partial* effect, holding other covariates fixed. The **marginal** correlation between $x_j$ and $y$ can have the opposite sign. This is Simpson's paradox territory and it happens routinely with confounded covariates.
*In simple regression with one covariate, sign$(\hat\beta_1) = $ sign$(r_{xy})$ — that part is true, but only there.*

**6c. "$R^2 = r^2$ always."** FALSE. Only in **simple** linear regression with an intercept.

---

## TRAP 7 — Log interpretation 🔴 *appeared in WS 23/24*

**Statement:** *"In a linear regression model applying logarithmic transformation, the coefficients should be interpreted as the percentage change in the response for a 1% change in the predictor."*

**Answer: FALSE.** That interpretation belongs to the **log-log** model only. The statement doesn't specify which variable was logged.

Keep the four cases straight:

| Model | 1-unit ↑ in $x$ | 1-**percent** ↑ in $x$ |
|---|---|---|
| $y \sim x$ | $\beta_1$ units in $y$ | — |
| $\log y \sim x$ | $100\beta_1$ % in $y$ | — |
| $y \sim \log x$ | — | $\beta_1/100$ units in $y$ |
| $\log y \sim \log x$ | — | $\beta_1$ % in $y$ |

**Mnemonic:** *the log is on the side that becomes a percentage.*

---

## TRAP 8 — Interpreting the intercept when you shouldn't

Saying "the expected wage of a man of age 0 is \$81.70" is not wrong arithmetic — it's wrong *judgement*, and examiners mark it down. Always add: *"however, age 0 lies outside the observed data range and is not meaningful, so $\hat\beta_0$ should not be substantively interpreted."*

---

## TRAP 9 — Causal language

Writing "an extra year of education **causes** a \$X increase in wage" loses marks. Regression on observational data identifies **association**, not causation, unless the covariate was randomly assigned.

Safe verbs: *is associated with · corresponds to · the model estimates · on average.*
Dangerous verbs: *causes · leads to · results in · produces.*

---

## TRAP 10 — SSR ambiguity

"SSR" means "Sum of Squares Regression" in some books and "Sum of Squares Residual" in others. **Never write SSR.** Write **SSE** for $\sum\hat\varepsilon_i^2$ and either "explained SS" or $\sum(\hat y_i - \bar y)^2$ for the other. Zero risk of a marker misreading you.

---

## TRAP 11 — Assuming a low $R^2$ means "no relationship"

$R^2 = 0.038$ for wage on age (Sheet 3). That is **not** evidence of no relationship. It could mean:
- the relationship is **non-linear** (and it is — the effect of age is quadratic, per Sheet 2)
- important covariates are **omitted** (Sheet 5 reaches $R^2 = 0.34$ with the full model)
- the outcome is intrinsically noisy

$R^2$ measures *linear explained variation in this model*, nothing more.

---

## TRAP 12 — Model class chosen by covariate type

Nobody states this outright, but people reason this way under pressure: "my covariates are categorical, so I need a special model." **No.** Categorical covariates become dummy variables and go into an ordinary linear model. It is **only the response variable's type** that changes the model class.

---

## Rapid-fire trap checklist (read the morning of the exam)

- [ ] Linear = linear in **β**
- [ ] $\exp$ model → take logs → still linear
- [ ] $\boldsymbol{X}$ has $k+1$ columns
- [ ] Count the betas for df
- [ ] $\varepsilon \neq \hat\varepsilon$
- [ ] $r=0$ ≠ independent
- [ ] Positive $\hat\beta_j$ ≠ positive correlation (multiple regression)
- [ ] $R^2 = r^2$ only in **simple** regression
- [ ] Log interpretation: check **which** variable is logged
- [ ] Don't interpret $\hat\beta_0$ if $x=0$ is meaningless
- [ ] Say "associated with," never "causes"
- [ ] Never write "SSR"
- [ ] Low $R^2$ ≠ no relationship
- [ ] Only **$y$'s** type picks the model class
