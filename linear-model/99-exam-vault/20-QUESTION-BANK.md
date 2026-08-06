# 20 — QUESTION BANK

**20 questions in the four exam blocks · answers and marking notes at the bottom**

> **How to use this.** Work a block on paper, timed, **formula before numbers**, rounding to 3 decimals. Then mark yourself against the key — and mark like Groll, not like a friend: a correct number with no formula loses method marks.
>
> A running data set is used for the computational block so you can chain the questions the way Sheets 3–5 do.

| Block | Points on the real paper | Covered here |
|---|---|---|
| Exercise 1 — TRUE/FALSE | 6 | → `30-TRUE-FALSE-DRILL.md` (72 statements) |
| Exercise 2 — model building | 6–35 | Q1–Q5 |
| Exercise 3 — computation | 6–25 | Q6–Q13 |
| Exercise 4 — explain / derive | 8 | Q14–Q20 |

---

# BLOCK 2 — Model building

**Q1.** A study models monthly `wage` (€) on `age` (years), `education` (3 levels: *none*, *school*, *degree*) and `birthplace` (2 levels: *domestic*, *foreign*). Write the model equation, define every dummy explicitly, name the reference categories, and state $p$.

**Q2.** In the model of Q1, wages are believed to rise with age and then flatten. Modify the model, write the marginal effect of age, and give the age at which the effect turns.

**Q3.** Still in Q1's setting, the return to education is believed to differ between domestic and foreign workers. Write the model with the necessary interaction terms and state how many parameters it now has.

**Q4.** In the fitted model $\widehat{\text{wage}} = 1800 + 45\cdot\text{age} + 620\cdot D_{\text{degree}} + 310\cdot D_{\text{school}} - 240\cdot D_{\text{foreign}}$, compute the expected wage gap between a 40-year-old domestic worker with a degree and a 25-year-old foreign worker with school education.

**Q5.** Write the restricted model implied by $H_0:\beta_1=\beta_2+1$ in $y_i=\beta_0+\beta_1x_{i1}+\beta_2x_{i2}+\varepsilon_i$, in a form you could actually fit, and state $r$.

---

# BLOCK 3 — Computation

**The running output.** A linear model is fitted to $n=50$ observations with an intercept and $k=4$ covariates.

```
Coefficients:
             Estimate  Std. Error  t value
(Intercept)    12.400       2.500    (a)
age             2.145       0.412    (b)
educ            (c)         0.048   -3.250
region          1.800       (d)       4.500

Residual standard error: (e) on (f) degrees of freedom
Multiple R-squared: 0.680,  Adjusted R-squared: (g)
```

Additionally: $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}=1250$, and $t_{45}(0.975)=2.014$, $F_{2,45}(0.95)=3.204$, $F_{4,45}(0.95)=2.579$.

**Q6.** Fill in (a) through (g).

**Q7.** Construct a 95% confidence interval for $\beta_{\text{age}}$ and state whether $H_0:\beta_{\text{age}}=0$ is rejected at the 5% level, giving your reason.

**Q8.** Test $H_0:\beta_{\text{age}}=1.5$ at the 5% level.

**Q9.** Compute the overall F statistic and carry out the overall model test.

**Q10.** A restricted model imposing $H_0:\beta_{\text{educ}}=\beta_{\text{region}}$ **and** $\beta_{\text{age}}=0$ gives $R^2_{H_0}=0.620$. Write $\boldsymbol{C}$ and $\boldsymbol{d}$, state $r$, compute $F$ and decide.

**Q11.** Compute AIC and BIC for this model. *(Use $\ln 25 = 3.219$, $\ln 50 = 3.912$.)*

**Q12.** A competing model with 2 covariates has $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}=1420$. Compute its AIC and BIC *(use $\ln 28.4 = 3.346$)* and say which model each criterion selects.

**Q13.** In a simple regression of $y$ on $x$ you are told $\hat\beta_1=0.750$, $\bar x=12$, $\bar y=30$. Find $\hat\beta_0$, and state the one property of OLS with an intercept that makes this possible.

---

# BLOCK 4 — Explain and derive

**Q14.** Derive $\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$. [2 marks]

**Q15.** State the Gauss–Markov theorem, listing the assumptions it requires. [4 marks]

**Q16.** Explain why a linear regression model is not appropriate for a binary dependent variable. [2 marks]

**Q17.** The errors in a fitted model are heteroscedastic. Discuss the consequences for the bias and the efficiency of the OLS estimator, and for inference. [3 marks]

**Q18.** Explain the difference between a confidence interval for $E(y\mid\boldsymbol{x}_0)$ and a prediction interval for a new observation $y_0$, and say which is wider and why. [2 marks]

**Q19.** Explain why $R^2$ cannot be used to choose between models with different numbers of covariates, and name two criteria that can. [2 marks]

**Q20.** Interpret $\hat\beta_j=0.693$ in a logit model for the probability of loan default, where $x_j$ is years of employment. State precisely what it does **not** mean. [2 marks]

---
---

# ANSWER KEY

## Q1

$$\text{wage}_i=\beta_0+\beta_1\text{age}_i+\beta_2D_{i,\text{school}}+\beta_3D_{i,\text{degree}}+\beta_4D_{i,\text{foreign}}+\varepsilon_i$$

with $D_{i,\text{school}}=1$ if worker $i$ has school education and 0 otherwise; $D_{i,\text{degree}}=1$ if degree, 0 otherwise; $D_{i,\text{foreign}}=1$ if foreign-born, 0 otherwise. **Reference categories: *none* (education) and *domestic* (birthplace).** $\varepsilon_i$ i.i.d. with $E(\varepsilon_i)=0$, $\text{Var}(\varepsilon_i)=\sigma^2$.

$$p=1+1+(3-1)+(2-1)=5$$

> **Marking note:** defining the dummies explicitly and naming the reference are separately credited. Writing "education" as a single variable is the standard way to lose 2 of 3 marks.

## Q2

Add a quadratic term: $\ \text{wage}_i=\beta_0+\beta_1\text{age}_i+\beta_2\text{age}_i^2+\dots+\varepsilon_i$

$$\frac{\partial E(\text{wage})}{\partial\text{age}}=\beta_1+2\beta_2\,\text{age}\qquad\text{turning point at }\ \text{age}^*=-\frac{\hat\beta_1}{2\hat\beta_2}$$

Still **linear in the parameters**, so it is still a linear model. Flattening then falling requires $\hat\beta_2<0$. Centring age (e.g. $\text{age}-48$) improves interpretability and reduces collinearity with $\text{age}^2$.

> 🔴 Quoting $\beta_1$ alone as "the effect of age" is the trap. A variable in two terms must be **differentiated**.

## Q3

Interact **both** education dummies with the foreign dummy:

$$\text{wage}_i=\beta_0+\beta_1\text{age}_i+\beta_2D_{\text{school}}+\beta_3D_{\text{degree}}+\beta_4D_{\text{foreign}}+\beta_5(D_{\text{school}}\!\cdot\!D_{\text{foreign}})+\beta_6(D_{\text{degree}}\!\cdot\!D_{\text{foreign}})+\varepsilon_i$$

$p=7$. The return to a degree is $\beta_3$ for domestic workers and $\beta_3+\beta_6$ for foreign workers — i.e. the dummy shifts, the interaction **tilts**.

## Q4

$$(1800+45(40)+620)-(1800+45(25)+310-240)=4220-2995=\boxed{1225\ \text{€}}$$

Or directly on the differences: $45(15)+620-310+240=675+310+240=1225$.

## Q5

Substitute $\beta_1=\beta_2+1$:

$$y_i=\beta_0+(\beta_2+1)x_{i1}+\beta_2x_{i2}+\varepsilon_i\ \Longrightarrow\ \underbrace{(y_i-x_{i1})}_{\text{new response}}=\beta_0+\beta_2\underbrace{(x_{i1}+x_{i2})}_{\text{new covariate}}+\varepsilon_i$$

**Regress $(y-x_1)$ on $(x_1+x_2)$**, and $\boldsymbol{r}=1$ — one equation, even though two betas appear.

## Q6

| | Value | Working |
|---|---|---|
| (a) | $4.960$ | $12.400/2.500$ |
| (b) | $5.206$ | $2.145/0.412$ |
| (c) | $-0.156$ | $t\times\widehat{\text{se}}=-3.250\times0.048$ |
| (d) | $0.400$ | $\hat\beta/t=1.800/4.500$ |
| (e) | $5.271$ | $\hat\sigma=\sqrt{1250/45}=\sqrt{27.778}$ |
| (f) | $45$ | $n-p=50-5$ |
| (g) | $0.652$ | $1-\frac{49}{45}(1-0.680)=1-1.089(0.320)$ |

> 🔴 (e) is a **standard error** — take the root. 🔴 (f) counts **all five** betas including the intercept.

## Q7

$$\hat\beta_{\text{age}}\pm t_{45}(0.975)\,\widehat{\text{se}}=2.145\pm2.014(0.412)=2.145\pm0.830=[\,1.315,\ 2.975\,]$$

**Zero lies outside the interval ⟹ reject $H_0:\beta_{\text{age}}=0$ at the 5% level.** Age is significantly associated with wage.

> 🔴 The quantile is $1-\alpha/2=0.975$, not $0.95$. 🔴 Zero **outside** means **reject** — the direction is examined every year.

## Q8

$$t=\frac{\hat\beta_j-c}{\widehat{\text{se}}}=\frac{2.145-1.5}{0.412}=\frac{0.645}{0.412}=1.566$$

$|1.566|<t_{45}(0.975)=2.014$ ⟹ **fail to reject** $H_0:\beta_{\text{age}}=1.5$ at the 5% level.

> 🔴 The $-c$ is the whole question. And note the answer flips relative to Q7 — significantly different from 0, not significantly different from 1.5. Consistent with the CI, which contains 1.5.

## Q9

$$F=\frac{R^2/k}{(1-R^2)/(n-p)}=\frac{0.680/4}{0.320/45}=\frac{0.170}{0.007111}=23.906\ \sim F_{4,45}$$

$23.906>F_{4,45}(0.95)=2.579$ ⟹ **reject** $H_0:\beta_1=\beta_2=\beta_3=\beta_4=0$. **At least one** covariate is significantly associated with the response.

## Q10

Parameter order $(\beta_0,\beta_{\text{age}},\beta_{\text{educ}},\beta_{\text{region}})$ — here shown for the four named parameters, with a zero column for any remaining covariate:

$$\boldsymbol{C}=\begin{pmatrix}0&0&1&-1\\0&1&0&0\end{pmatrix},\qquad \boldsymbol{d}=\begin{pmatrix}0\\0\end{pmatrix},\qquad r=2$$

$$F=\frac{(R^2-R^2_{H_0})/r}{(1-R^2)/(n-p)}=\frac{(0.680-0.620)/2}{0.320/45}=\frac{0.030}{0.007111}=4.219\ \sim F_{2,45}$$

$4.219>F_{2,45}(0.95)=3.204$ ⟹ **reject $H_0$**. **At least one** of the two restrictions fails — we cannot say which.

> 🔴 $r=2$ because there are **two equations**. 🔴 The F quantile is $0.95$, not $0.975$. 🔴 "At least one," never "both."

## Q11

$$\hat\sigma^2_{ML}=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n}=\frac{1250}{50}=25\qquad n\log\hat\sigma^2_{ML}=50(3.219)=160.950$$

$$\text{AIC}=160.950+2(4+1)=\boxed{170.950}\qquad \text{BIC}=160.950+3.912(4+1)=160.950+19.560=\boxed{180.510}$$

> 🔴 Divide by $\boldsymbol{n}=50$, **not** $n-p=45$. 🔴 Natural log. 🔴 The penalty term is $|M|+1=5$ — the $+1$ is $\sigma^2$.

## Q12

$$\hat\sigma^2_{ML}=\frac{1420}{50}=28.4,\qquad n\log\hat\sigma^2_{ML}=50(3.346)=167.300$$

$$\text{AIC}=167.300+2(3)=173.300\qquad \text{BIC}=167.300+3.912(3)=167.300+11.736=179.036$$

| | 4-covariate model | 2-covariate model | Selects |
|---|---|---|---|
| AIC | **170.950** | 173.300 | the **larger** model |
| BIC | 180.510 | **179.036** | the **smaller** model |

**They disagree — and that is the expected direction.** BIC's penalty per parameter is $\log(50)=3.912$ against AIC's $2$, so BIC punishes the two extra covariates harder. *B for Bigger penalty ⟹ smaller models.*

## Q13

$$\hat\beta_0=\bar y-\hat\beta_1\bar x=30-0.750(12)=30-9=\boxed{21}$$

Possible because **with an intercept the fitted line passes through $(\bar x,\bar y)$** — equivalently $\bar{\hat y}=\bar y$ and $\sum\hat\varepsilon_i=0$. This is the standard route to a missing intercept in an R output question.

## Q14 [2 marks]

$$S(\boldsymbol\beta)=(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)=\boldsymbol{y}'\boldsymbol{y}-2\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{y}+\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{X}\boldsymbol\beta$$

$$\frac{\partial S}{\partial\boldsymbol\beta}=-2\boldsymbol{X}'\boldsymbol{y}+2\boldsymbol{X}'\boldsymbol{X}\boldsymbol\beta\overset{!}{=}\boldsymbol{0}\ \Longrightarrow\ \boldsymbol{X}'\boldsymbol{X}\hat{\boldsymbol\beta}=\boldsymbol{X}'\boldsymbol{y}\ \Longrightarrow\ \hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$$

Inversion requires $\text{rank}(\boldsymbol{X})=p$. Since $\partial^2S/\partial\boldsymbol\beta\partial\boldsymbol\beta'=2\boldsymbol{X}'\boldsymbol{X}$ is positive definite, this is a minimum.

> **Marking key (WS 23/24, verbatim):** *"1 point for correctly stating that RSS needs to be minimized. And 1 point for correctly deriving the solution."* Say what you are minimising **before** you differentiate.

## Q15 [4 marks]

> Under the assumptions **(i)** the model is correctly specified / linear, **(ii)** $E(\boldsymbol\varepsilon)=\boldsymbol{0}$, **(iii)** homoscedastic errors $\text{Var}(\varepsilon_i)=\sigma^2$, **(iv)** uncorrelated errors $\text{Cov}(\varepsilon_i,\varepsilon_j)=0$ for $i\neq j$ — jointly $\text{Cov}(\boldsymbol\varepsilon)=\sigma^2\boldsymbol{I}_n$ — and **(v)** $\text{rank}(\boldsymbol{X})=p$, the OLS estimator $\hat{\boldsymbol\beta}$ is **BLUE**: among all estimators that are **linear in $\boldsymbol{y}$** and **unbiased** for $\boldsymbol\beta$, it has minimum variance.
>
> **Normality of the errors is *not* required.**

> **Marking key notes *"1 point for every assumption"*** — so **list them**, don't prose them. And the normality sentence is a genuine mark: it is what the question is testing. Dropping either "linear" or "unbiased" makes the claim false, since biased estimators (e.g. ridge) can have smaller variance.

## Q16 [2 marks]

1. **Fitted values are not confined to $[0,1]$** — $\boldsymbol{x}'\hat{\boldsymbol\beta}$ is unbounded, so the model can predict probabilities below 0 or above 1, which are meaningless. *(Lead with this.)*
2. **Heteroscedasticity by construction:** $\text{Var}(y\mid\boldsymbol{x})=\pi(1-\pi)$ depends on $\boldsymbol{x}$, so A3 fails automatically.
3. **Errors cannot be normal:** for given $\boldsymbol{x}$, $\varepsilon$ takes only the two values $1-\pi$ and $-\pi$.
4. A **constant** marginal effect is implausible — near $\pi=0$ or $1$ the same change in $x$ cannot keep moving the probability.

The fix: model $\pi=h(\boldsymbol{x}'\boldsymbol\beta)$ with $h$ squashing the linear predictor into $(0,1)$ — the logit link.

## Q17 [3 marks]

> **Bias:** OLS remains **unbiased** and consistent. Unbiasedness requires only correct specification, zero-mean errors and full rank — none of which heteroscedasticity violates.
>
> **Efficiency:** OLS is **no longer efficient**. Gauss–Markov requires homoscedastic errors, so $\hat{\boldsymbol\beta}$ is no longer BLUE; weighted or generalised least squares would achieve smaller variance.
>
> **Inference:** the usual variance formula $\sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}$ is wrong, so the reported standard errors are biased and **t-tests, F-tests and confidence intervals are invalid** — even though the point estimates are fine.

> 🔴 "Heteroscedasticity biases the estimates" is FALSE and is the single most-repeated wrong answer on this topic.

## Q18 [2 marks]

$$\text{CI (mean): }\ \boldsymbol{x}_0'\hat{\boldsymbol\beta}\pm t\,\hat\sigma\sqrt{\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0}\qquad \text{PI: }\ \boldsymbol{x}_0'\hat{\boldsymbol\beta}\pm t\,\hat\sigma\sqrt{\mathbf{1}+\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0}$$

The CI covers the **mean response** $E(y\mid\boldsymbol{x}_0)$ — a fixed unknown number — and reflects only the uncertainty in $\hat{\boldsymbol\beta}$. The PI covers a **single new observation** $y_0=\boldsymbol{x}_0'\boldsymbol\beta+\varepsilon_0$ and must additionally carry that observation's own error $\varepsilon_0$, which contributes the extra $\mathbf{1}$ under the root.

**The prediction interval is therefore always wider**, and as $n\to\infty$ its width tends to $\pm t\hat\sigma$ rather than to zero: estimation error vanishes, individual noise does not.

## Q19 [2 marks]

$R^2$ is **monotone in model size** — adding any covariate expands the column space, so the projection moves closer to $\boldsymbol{y}$ and $R^2$ can only rise (SSE only fall). Selecting on $R^2$ therefore always picks the largest model regardless of whether the extra covariates carry information, i.e. it rewards **overfitting**.

Criteria that penalise complexity instead: **AIC** $=n\log\hat\sigma^2_{ML}+2(|M|+1)$ and **BIC** $=n\log\hat\sigma^2_{ML}+\log(n)(|M|+1)$. *(Adjusted $R^2$ and Mallow's $C_p$ also acceptable — but note $\bar R^2$'s penalty is weak.)*

## Q20 [2 marks]

> **"A one-year increase in employment is associated with an estimated increase of 0.693 in the log-odds of default, holding all other covariates fixed. Equivalently, the odds of default are multiplied by $\exp(0.693)=2.000$ — they double."**

**What it does not mean:** it is **not** an increase of 0.693 in the probability of default. On the probability scale the effect is $\hat\beta_j\,\pi(1-\pi)$, which **depends on $\pi$** and is therefore not constant across individuals. Only the **sign** of $\hat\beta_j$ transfers unambiguously from the log-odds scale to the probability scale.

> 🔴 This is the Exam 2025 trap in interpretation form. Naming the wrong scale costs the mark even when the arithmetic is right.

---

## After marking

Log every miss as one of four types — the type tells you what to fix:

| Type | Fix |
|---|---|
| **Didn't know the formula** | → `10-FORMULA-SHEET.md`, blank-page reproduction |
| **Knew the formula, picked the wrong one** | → the relevant `33-FORMULA-DECISION-GUIDE.md` |
| **Right method, arithmetic slip** | → slow down on quantiles and denominators; they cause most of these |
| **Right answer, no working** | → **formula before numbers.** This one costs real marks on the day |
