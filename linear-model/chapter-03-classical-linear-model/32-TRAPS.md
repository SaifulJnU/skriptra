# Ch 3 — POSSIBLE TRAPS

> 🔴 = appeared **verbatim** in one of your five past papers. Read this file on Day 15, Day 20, and the morning of the exam.
>
> **The meta-pattern first, because it saves more marks than any individual trap:**
>
> | Construction | Example | Defence |
> |---|---|---|
> | **True first clause, false second** | *"In QQ plots the empirical quantiles are compared to theoretical quantiles [✓]. Points should follow a **horizontal line** [✗]."* | **Read to the very end.** Don't stop when it starts sounding right. |
> | **Irrelevant "as long as…" clause** | *"OLS estimates are unique as long as **the error variance is constant**"* | Ask: does that condition have anything to do with the claim? |
> | **"if and only if" / "must" / "always"** | *"BLUE **if and only if** zero mean and constant variance"* | These are almost always too strong. Check for missing assumptions. |
> | **Correct fact, wrong direction** | *"AIC penalises more heavily than BIC"* · *"CI excludes zero ⟹ fail to reject"* | Know which way round, not just that a relationship exists. |
> | **Right concept, wrong scale** | *"$\hat\beta_j$ increases $P(y=1)$ by $\hat\beta_j$"* | Ask: which scale is this quantity measured on? |

---

# GROUP A — Degrees of freedom & notation

## TRAP A1 🔴 — The `p` convention

| | Book (Fahrmeir) | Some papers |
|---|---|---|
| $p$ = | **parameters** $=k+1$ | **covariates** $=k$ |
| t-test df | $n-p$ | $n-p-1$ |

**Evidence from your own answer keys:**
- *Linear_model_exam_sheet* I(i): *"rank$(X'X)=p$, with the number of variables $p$"* → **FALSE**; I(iii): *"$t_j\sim t_{n-p-1}$"* → **TRUE**. That paper uses $p$ = covariates.
- *RCLM WS22/23* I(i): *"rk$(X'X)=p$"* bare → **TRUE** in book notation.

> 🛟 **Defence:** never memorise a symbol string. **Count the $\beta$'s you estimated, including the intercept, and subtract from $n$.** Write $n-k-1$ in free-response answers — unambiguous under both conventions.

## TRAP A2 🔴 — Rank claims

| Claim | Verdict (book notation) |
|---|---|
| $\text{rk}(\boldsymbol{X}'\boldsymbol{X})=p$ | ✅ TRUE |
| $\text{rk}(\boldsymbol{X}'\boldsymbol{X})=k$ | ❌ FALSE — one short, forgot the intercept |
| $\text{rk}(\boldsymbol{X}'\boldsymbol{X})=n$ | ❌ FALSE — **rank can never exceed the smaller dimension**, and $\boldsymbol{X}'\boldsymbol{X}$ is $p\times p$ with $p\ll n$ |

## TRAP A3 🔴 — F-test degrees of freedom

- *"$F\sim F_{p+1,\,n-p}$ for $H_0:\boldsymbol{C}\boldsymbol\beta=\boldsymbol{d}$"* (WS22/23 II(iv)) → **FALSE.** First df is **$r$**, not $p+1$.
- *"Full-model F test: $F\sim F_{k,\,n-k-1}$"* (WS23/24 I(iii)) → ✅ **TRUE**.
- *"$t_j\sim t_{n-k-1}$ where $k+1$ = number of parameters"* (Exam 2025 1(g)) → ✅ **TRUE**.

**Rule:** numerator df $=r$ (equations). Denominator df $=n-p$ from the **unrestricted** model.

---

# GROUP B — Counting restrictions

## TRAP B1 🔴🔴 — The biggest one in Section 3.3

> **Exam Summer 2025, Ex 1(i):** *"The F-statistic for $H_0:\beta_1=-\beta_2+\beta_3$ in a model with $k\geq3$ predictors plus an intercept has an F-distribution with $(3,\,n-k-1)$ degrees of freedom."*
>
> ## **FALSE.**

$\beta_1=-\beta_2+\beta_3$ is **ONE equation** ⟹ $r=1$ ⟹ $F\sim F_{1,\,n-k-1}$.

The "3" counts the **betas mentioned**. That is never what $r$ means.

> 🔑 **$r$ = the number of independent EQUATIONS in $H_0$ = the number of rows of $\boldsymbol{C}$.**
>
> Rewrite as $\beta_1+\beta_2-\beta_3=0$ and count the equals signs. **One.**

## TRAP B2 — Forgetting the $\beta_0$ column in $\boldsymbol{C}$

$\boldsymbol{C}$ is $r\times \mathbf{p}$, not $r\times k$. Every row starts with the coefficient on $\beta_0$ — usually 0, but the column must be there or the dimensions don't conform.

## TRAP B3 — Not rearranging before reading off

$\beta_{\text{crim}}=3\beta_{\text{rad}}-0.1$ must first become $\beta_{\text{crim}}-3\beta_{\text{rad}}=-0.1$. Only then can you read the row $(0,1,0,0,-3)$ and $d=-0.1$.

---

# GROUP C — Which $\hat\sigma^2$?

## TRAP C1 🔴 — AIC/BIC use $n$, everything else uses $n-p$

$$\hat\sigma^2_{ML}=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n}\ \textbf{for AIC/BIC} \qquad\qquad \hat\sigma^2=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n-p}\ \textbf{for everything else}$$

The book states it explicitly: *"the ML estimator $\hat\sigma^2=\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}/n$ is considered in AIC and **not** the usual unbiased variance estimator."*

**Word triggers:** "REML" / "restricted maximum likelihood" ⟹ $n-p$. "Maximum likelihood estimate of $\sigma^2$" ⟹ $n$. AIC/BIC ⟹ always $n$.

## TRAP C2 — $\log_{10}$ instead of $\ln$

Sheet 5 says it outright: *"the book uses $\log(x)$ as the natural logarithm with base $e$ (and not base 10)."* If they felt the need to warn you, people get it wrong.

## TRAP C3 — Forgetting the $+1$

$$\text{AIC}=n\log(\hat\sigma^2)+2(|M|\mathbf{+1})$$

The $+1$ is $\sigma^2$ — it's a parameter too.

## TRAP C4 — Forgetting the square root

R reports the residual **standard error** $\hat\sigma$, not $\hat\sigma^2$. If your answer is 1276 and the output says 34, you skipped the root.

---

# GROUP D — Quantiles

## TRAP D1 🔴 — $1-\alpha$ vs $1-\alpha/2$ (the most common numerical error in this exam)

| Situation | Quantile |
|---|---|
| 95% CI | **0.975** |
| 99% CI | **0.995** |
| Two-sided t-test, $\alpha=0.05$ | **0.975** |
| **F-test, $\alpha=0.05$** | **0.95** |

> 🔑 **t and CI are two-sided ⟹ split $\alpha$ ⟹ $1-\alpha/2$.**
> **F only rejects in the upper tail (since $\text{SSE}_{H_0}\geq\text{SSE}$) ⟹ one-sided ⟹ $1-\alpha$.**

**Habit:** write "$1-\alpha/2 = 0.975$" beside the number before you look it up.

## TRAP D2 — Reading the wrong table axis

Row = numerator df ($r$). Column = quantile level. Caption = denominator df ($n-p$, often fixed).

---

# GROUP E — Assumptions and their consequences

## TRAP E1 🔴 — "Heteroscedasticity biases the estimates"

**FALSE.** OLS stays **unbiased and consistent**. What you lose: **efficiency** (not BLUE) and **valid standard errors** (hence invalid t/F/CI).

> *Exam 2025 Ex 4(e)* asks for "bias and efficiency" — the two words tell you the answer: **unbiased but inefficient**.
> *WS 23/24 II(ii)*: "correlated residuals don't affect consistency but may affect efficiency" → ✅ **TRUE** — same structure.

## TRAP E2 🔴 — "Normality is required for BLUE"

**FALSE.** Gauss–Markov needs only A1–A5. Normality buys the **exactness** of tests and makes OLS = ML.

> *WS 23/24 I(ii):* *"the LS estimator is BLUE **if and only if** the error term is expected to be zero and has constant variance"* → **FALSE.** The "iff" is wrong, and the list omits **uncorrelated errors**, **correct specification** and **full rank**.

## TRAP E3 🔴 — "Rank deficiency is fine if the error variance is constant"

> *Exam 2025 Ex 1(d):* → **FALSE.** Without full rank, $\boldsymbol{X}'\boldsymbol{X}$ is singular and the normal equations have **infinitely many** solutions. Homoscedasticity is **completely irrelevant** to identification. Watch for "as long as…" clauses that attach an unrelated condition.

## TRAP E4 🔴 — Multicollinearity, in four flavours

| Statement | Verdict |
|---|---|
| "Multicollinearity can inflate the variance of the OLS estimators" *(Exam25 1j)* | ✅ **TRUE** |
| "…inflates variance but does not bias the estimates" *(LMES II(i))* | ✅ **TRUE** |
| "Correlated explanatory variables may lead to a **reduction** in standard errors" *(WS23/24 II(i))* | ❌ **FALSE** — inflates |
| "If VIF ≈ 1 for all variables, multicollinearity is likely a concern" *(LMES II(ii))* | ❌ **FALSE** — VIF ≈ 1 means **no** collinearity |

**And keep perfect vs near separate:** perfect breaks A5 (not identified); near leaves $\hat{\boldsymbol\beta}$ **unbiased and still BLUE**, only imprecise.

## TRAP E5 🔴 — "Minimising absolute deviations gives the same estimates"

> *Exam 2025 Ex 1(b):* → **FALSE.** LAD targets the conditional **median**, has no closed form, and gives different estimates. It's more robust to outliers.

## TRAP E6 🔴 — "OLS = ML under normal errors"

> *Exam 2025 Ex 1(l):* → ✅ **TRUE** for $\boldsymbol\beta$. $\boldsymbol\beta$ enters the Gaussian log-likelihood only through $-\frac{1}{2\sigma^2}(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)$, so maximising = minimising SSE.
>
> ⚠️ But $\hat\sigma^2_{ML}\neq\hat\sigma^2_{LS}$ — different denominators.

## TRAP E7 🔴 — "$\hat\beta_0$ and $\hat\beta_1$ are always uncorrelated"

> *Exam 2025 Ex 1(f):* → **FALSE.**
> $$\text{Cov}(\hat\beta_0,\hat\beta_1)=\frac{-\sigma^2\bar x}{\sum(x_i-\bar x)^2}$$
> Zero **only if $\bar x=0$**. The word **"always"** is what breaks it.

---

# GROUP F — Monotonicity of $R^2$ and SSE

## TRAP F1 🔴 — Four papers, one fact

$$\textbf{Adding a covariate: } R^2 \textbf{ can only rise (or stay). SSE can only fall (or stay).}$$

| Statement | Verdict |
|---|---|
| "Adding weekday-of-birth dummies can be expected to **lower** $R^2$" *(Exam25 1c)* | ❌ **FALSE** |
| "RSS **may increase** as more variables are added" *(WS23/24 III(i))* | ❌ **FALSE** |
| "The coefficient of determination **can decrease** as variables are added" *(LMES II(iii))* | ❌ **FALSE** |
| "Adjusted $R^2$ … **can never be negative**" *(LMES III(iv))* | ❌ **FALSE** — it can |

**Why:** adding a column expands the column space, so the projection can only get closer to $\boldsymbol{y}$. *(You can't un-place a jigsaw piece.)*

**But $\bar R^2$, AIC and BIC CAN worsen** — that's the entire point of Section 3.4.

## TRAP F2 🔴 — "AIC penalises more heavily than BIC"

> *LMES III(iii):* → **FALSE.** $\log(n)>2$ for $n>7.4$, so **BIC penalises more** and selects **smaller** models.
>
> **Mnemonic: B for Bigger penalty.**

## TRAP F3 🔴 — "Adding an unrelated predictor must decrease AIC"

> *WS 23/24 III(ii):* → **FALSE.** An irrelevant variable barely improves fit while costing a full penalty, so AIC typically **increases**. And "must" is far too strong in either direction.

---

# GROUP G — Testing and intervals

## TRAP G1 🔴 — CI logic backwards

> *Exam 2025 Ex 1(k):* *"If a CI for $\beta_j$ does not contain zero, we **fail to reject** $H_0:\beta_j=0$."* → **FALSE.** Excluding zero means we **DO reject**.
>
> *WS 23/24 II(iii):* *"When the CI contains zero, we cannot reject."* → ✅ **TRUE.** Same fact, stated correctly.

**Memory hook:** *zero inside the net ⟹ zero still possible ⟹ don't reject.*

## TRAP G2 — Forgetting the $-c$

$H_0:\beta_j=c$ needs $t=\dfrac{\hat\beta_j-c}{\widehat{\text{se}}}$. Students routinely compute $\hat\beta_j/\widehat{\text{se}}$ out of habit.

## TRAP G3 🔴 — Forgetting the "$1+$" in the prediction interval

$$\text{CI (mean): } \hat\sigma\sqrt{\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0} \qquad \text{PI: } \hat\sigma\sqrt{\mathbf{1}+\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0}$$

In Sheet 4(3e) this is the difference between **±\$4** and **±\$70**. A 17-fold error.

**How to tell which is wanted:** *"the wage of a 50 year old man"* — singular, one person ⟹ **prediction interval**. *"the average wage of…"* ⟹ CI.

## TRAP G4 — "We accept $H_0$"

Write **"fail to reject"** / "cannot reject" / "not significantly different from zero." Failing to reject may just mean low power.

## TRAP G5 — "Rejecting a joint $H_0$ means all restrictions fail"

**No — at least one.** The book says so explicitly in the kitchen-quality example: *"The null hypothesis is rejected when **at least one** coefficient significantly differs from zero."*

## TRAP G6 — Denominator uses the wrong SSE or the wrong df

$$F=\frac{(\text{SSE}_{H_0}-\text{SSE})/r}{\underbrace{\text{SSE}}_{\text{UNRESTRICTED}}/\underbrace{(n-p)}_{\text{UNRESTRICTED}}}$$

**Sanity checks:** $F\geq0$ always (negative ⟹ you swapped the SSEs). If $r=1$, $\sqrt{F}$ must equal $|t|$.

---

# GROUP H — Residuals and diagnostics

## TRAP H1 🔴 — "Standardised residuals are normal because residuals are"

> *RCLM WS22/23 III(i):* Treat as **FALSE** and explain why.
>
> Even under **perfect** assumptions, $\text{Var}(\hat\varepsilon_i)=\sigma^2(1-h_{ii})$ is **not constant** and $\text{Cov}(\hat\varepsilon_i,\hat\varepsilon_j)=-\sigma^2h_{ij}\neq0$. The whole *reason* for standardising is that raw residuals are heteroscedastic and correlated. And even standardised ones aren't exactly normal, since $\hat\sigma$ comes from the same data.

## TRAP H2 🔴 — QQ plot follows a HORIZONTAL line

> *RCLM WS22/23 I(iv):* → **FALSE.** Points follow the **45° diagonal** ($y=x$). A horizontal line would mean every empirical quantile is identical — no variation at all.
>
> **Classic true-first-clause construction: the sentence starts correctly and breaks at the end.**

## TRAP H3 🔴 — "Residual plots cannot identify non-linearity"

> *LMES I(ii):* → **FALSE.** Residuals-vs-fitted is *precisely* the non-linearity detector.

## TRAP H4 — Conflating leverage, outliers and influence

| Concept | Unusual in | Measured by |
|---|---|---|
| **Leverage** | $\boldsymbol{x}$ | $h_{ii}$ |
| **Outlier** | $y$ | $|r_i|$ |
| **Influence** | both, jointly | $D_i=\frac{r_i^2}{p}\cdot\frac{h_{ii}}{1-h_{ii}}$ |

**High leverage alone is harmless** — a distant point sitting exactly on the line *improves* precision. It's dangerous only with a large residual too.

## TRAP H5 — "Mean of residuals ≈ 0 shows the model is good"

$\sum\hat\varepsilon_i=0$ is **guaranteed by construction** whenever there's an intercept. A terrible model satisfies it too. That's why diagnostics look at the **pattern**, not the mean.

*(WS23/24 II(iv) marks the statement TRUE — but understand it's true automatically, not diagnostically.)*

---

# GROUP I — Chapter 4 leakage

## TRAP I1 🔴 — The ridge estimator

> *RCLM WS22/23 II(iii):* *"The ridge estimator adds a constant $\lambda$ to every component of $\boldsymbol{X}'\boldsymbol{X}$ to obtain a regularized and invertible matrix."*
>
> **FALSE** as stated — ridge adds $\lambda$ to the **diagonal only**: $(\boldsymbol{X}'\boldsymbol{X}+\lambda\boldsymbol{I})^{-1}\boldsymbol{X}'\boldsymbol{y}$. "Every component" would add $\lambda$ to all $p^2$ entries.
>
> That one sentence is the entire depth required. Ridge is Chapter 4 (out of scope), but one T/F can still appear.

---

# ⚡ THE MORNING-OF-THE-EXAM CHECKLIST

**Degrees of freedom**
- [ ] Count the $\beta$'s (incl. intercept); df $=n-k-1$
- [ ] $r$ = number of **equations**
- [ ] F: numerator $r$, denominator $n-p$ from the **unrestricted** model

**Quantiles**
- [ ] t & CI ⟹ $1-\alpha/2$ · F ⟹ $1-\alpha$

**Variance**
- [ ] AIC/BIC ⟹ divide by $n$, natural log, $+1$
- [ ] Everything else ⟹ $n-p$
- [ ] Don't forget the square root

**Assumptions**
- [ ] Heteroscedasticity/autocorrelation ⟹ **unbiased, inefficient, invalid se**
- [ ] Only A1 (misspecification) biases $\hat{\boldsymbol\beta}$
- [ ] Normality NOT needed for Gauss–Markov
- [ ] BLUE = min variance among **linear** AND **unbiased**
- [ ] VIF ≈ 1 = no problem

**Monotonicity**
- [ ] $R^2$ can't fall, SSE can't rise
- [ ] $\bar R^2$ **can** be negative
- [ ] **BIC penalises more** (B for Bigger)

**Testing**
- [ ] Zero outside the CI ⟹ **reject**
- [ ] Prediction interval has the **"$1+$"**
- [ ] Don't forget the $-c$
- [ ] $F\geq0$; if $r=1$, $\sqrt{F}=|t|$
- [ ] "at least one" · "fail to reject"

**Diagnostics**
- [ ] QQ plot = **45° diagonal**
- [ ] Leverage ≠ outlier ≠ influence
- [ ] se uses the **diagonal**; $\beta_1$ is the **2nd** element

**Language**
- [ ] "associated with" · "holding all other covariates fixed" · "expected"
- [ ] **Formula before numbers** · round to 3 decimals
- [ ] **Never leave a TRUE/FALSE blank** — negative marking is floored at zero per block

**And read every T/F statement to its final word.**
