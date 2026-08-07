# Ch 2 — PAST-PAPER QUESTIONS (all five papers)

> **S25** = Exam Summer 2025 · **LMES** = Linear_model_exam_sheet · **W23** = WiSe 2023/24 · **W22** = RCLM WS 22/23 · **EX20** = Example Exam LiMo 2020

**Chapter 2 is worth far more than its ~13% share suggests, because EX20 devoted an entire 25-point exercise to it.** Dummy coding, interactions and logit interpretation appear on every paper without exception.

| Topic | Appears in |
|---|---|
| Dummy coding, $c-1$ rule | S25, LMES, W23, W22, EX20 |
| Interactions | EX20 (25 pts), W22 |
| Interpretation of coefficients | LMES, W23, EX20 |
| Logit | S25 (twice) |
| Log/polynomial transformations | S25, W23, W22 |

---

# PART A — Model building and dummy coding

## Q1 🔴 — Build a model with two categorical covariates

> **S25, Ex 2(a) [3 Points].** *"You want to examine the effect of a person's age, education (no degree, high school degree or college/university degree) and place of birth (inside the US or outside the US) on their hourly wage… Provide an adequate linear model equation that incorporates all of the above explanatory variables, including an intercept. Formulate it such that you can estimate this exact equation by OLS."*

### Solution

$$\text{wage}_i=\beta_0+\beta_1\,\text{age}_i+\beta_2 D^{\text{HS}}_i+\beta_3 D^{\text{college}}_i+\beta_4 D^{\text{outside}}_i+\varepsilon_i,\qquad i=1,\dots,n$$

with

$$D^{\text{HS}}_i=\begin{cases}1&\text{if person }i\text{ has a high school degree}\\0&\text{else}\end{cases}\qquad
D^{\text{college}}_i=\begin{cases}1&\text{if person }i\text{ has a college/university degree}\\0&\text{else}\end{cases}$$

$$D^{\text{outside}}_i=\begin{cases}1&\text{if person }i\text{ was born outside the US}\\0&\text{else}\end{cases}$$

**Reference categories: *no degree* and *born inside the US*.** Errors $\varepsilon_i$ i.i.d. with $E(\varepsilon_i)=0$, $\text{Var}(\varepsilon_i)=\sigma^2$.

$$p = 1+1+(3-1)+(2-1)=5$$

> **Why "formulate it such that you can estimate this exact equation by OLS" is in the question:** it is telling you to *write out the dummies*. Writing `education` as one symbol is not estimable by OLS — a categorical variable is not a number. That phrase is the examiner handing you the method.
>
> 🔴 Three dummies for three education levels **plus** an intercept ⟹ the columns sum to the intercept column ⟹ $\boldsymbol{X}'\boldsymbol{X}$ singular ⟹ **no unique OLS solution**. That is the *dummy variable trap*, and it is the reason the $c-1$ rule exists.

---

## Q2 🔴🔴 — The $c-1$ rule as TRUE/FALSE

Asked on **two** papers, once each way round — this is the single most reliably repeated T/F in the course.

> **LMES, Ex 1a(iv).** *"In a regression model with a categorical predictor of $k$ levels, $k-1$ dummy variables are created to represent this predictor."* → ✅ **TRUE** *(official key: `FFTT`)*
>
> **W23, Ex 1a(iv).** *"For nominal variables with $m$ categories, we need to add $m$ dummy variables to represent this predictor when constructing the linear model."* → ❌ **FALSE** *(official key: `FFTF`)*

**Same fact, opposite verdicts, because the second one dropped the "−1".** Read to the final word.

---

## Q3 — Recovering the design from an R output

> **W22, Ex 2** context. The output lists `season2, season3, season4`, `yr1`, `daytimemorning, daytimemidday, daytimeafternoon, daytimeevening`, `holiday1`, plus `temp`, `hum`, `windspeed`.

### What you must be able to read off

| Variable | Levels | Dummies printed | Reference (the level **missing** from the output) |
|---|---|---|---|
| `season` | 4 | 3 | **season1 = winter** |
| `yr` | 2 | 1 | **yr0 = 2011** |
| `daytime` | 5 | 4 | **night** |
| `holiday` | 2 | 1 | **holiday0** |

$$p = 1 + 3 + 1 + 4 + 1 + 3\ \text{continuous} = 13$$

> 🔑 **The reference category is always the level that does not appear.** Every printed coefficient is a difference *from* it. To compare two **non-reference** levels — say summer against spring — **subtract**: $18.977-37.108=-18.131$ bikes per hour.

---

## Q4 — Why a centred quadratic term

> **S25, Ex 2(b) [2 Points].** *"Explain briefly why adding the following regressor to the equation might be a sensible idea: $\text{gage2}_i := (\text{age}_i-48)^2$."*
>
> **S25, Ex 2(c) [1 Point].** *"Would you expect the OLS estimate of the coefficient for $\text{gage2}_i$ to be positive or negative? Explain your choice."*

### Solution (b)

The effect of age on wage is **not plausibly linear**: wages rise through early and mid career, plateau, and fall near retirement. A model linear in age forces one constant effect over the whole range and will be **misspecified** (a violation of A1, the only assumption whose failure *biases* $\hat{\boldsymbol\beta}$).

Adding a squared term lets the fitted curve bend while **remaining a linear model** — linear in the parameters is what matters, not linear in the variables.

**Centring at 48** (roughly mid-career) does two things: it reduces the collinearity between $\text{age}$ and $\text{age}^2$, and it makes $\beta_1$ interpretable as the marginal effect **at age 48** rather than at the meaningless age 0.

### Solution (c)

**Negative.** Wage peaks in mid career and declines either side, so the parabola must open **downwards**, i.e. $\hat\beta_{\text{gage2}}<0$. With the centring at 48, the fitted peak sits near 48.

> 🔴 With age in two terms, the marginal effect is $\dfrac{\partial E(\text{wage})}{\partial\text{age}}=\beta_1+2\beta_2(\text{age}-48)$ — **never quote $\beta_1$ alone**. Turning point at $\text{age}=48-\hat\beta_1/(2\hat\beta_2)$.

---

# PART B — Interpretation

## Q5 🔴 — Interpreting a log-transformed covariate

> **LMES, Ex 3(c) [1 Point].** *"What is the interpretation of the coefficients of the independent variable `Population`?"* — the model is `Life.Exp ~ log(Population) + …` with $\hat\beta_{\log(\text{Pop})}=0.2707$.

### Solution — the official key

> *"As population increases by 1%, the average life expectancy will increase by $0.01\times0.27=0.0027$ units, given other conditions unchanged."*

**The key explicitly deducts half a mark if "on average" or "given other conditions unchanged" is missing.**

> 🔑 **Log-covariate rule:** $\hat\beta/100$ per **1% increase** in $x$. That factor of 100 is the whole question.
>
> And the mirror-image trap: **W23 Ex 1c(iv)** — *"In a linear regression model applying logarithmic transformation, the coefficients should be interpreted as the percentage change in the response for a 1% change in the predictor."* → ❌ **FALSE**. That elasticity reading needs logs on **both** sides. Here only the covariate is logged, so the response moves in **units**, not percent.

---

## Q6 — Interpreting intercept and slope from scratch

> **W23, Ex 3(a) [3 Points].** Five pairs: $X=(20,30,33,40,15)$, $Y=(7,9,8,11,5)$. *"Build a regression model of consumption $Y$ on income $X$ with intercept and estimate the intercept and slope. How do you interpret the estimated intercept and slope?"* Hint given: $(\boldsymbol{X}'\boldsymbol{X})^{-1}=\begin{pmatrix}2.079&-0.068\\-0.068&0.0024\end{pmatrix}$

### Solution

$$\boldsymbol{X}'\boldsymbol{X}=\begin{pmatrix}n&\sum x_i\\\sum x_i&\sum x_i^2\end{pmatrix}=\begin{pmatrix}5&138\\138&4214\end{pmatrix},\qquad \boldsymbol{X}'\boldsymbol{Y}=\begin{pmatrix}\sum y_i\\\sum x_iy_i\end{pmatrix}=\begin{pmatrix}40\\1189\end{pmatrix}$$

$$\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{Y}=\begin{pmatrix}2.079&-0.068\\-0.068&0.0024\end{pmatrix}\begin{pmatrix}40\\1189\end{pmatrix}=\begin{pmatrix}2.210\\0.209\end{pmatrix}$$

$$\boxed{\hat y=2.210+0.209\,x}$$

**Interpretation.** *Slope:* a one-unit increase in income is associated with an estimated **0.209-unit** increase in expected consumption. *Intercept:* the expected consumption at income $x=0$ is 2.210 — but $x=0$ is **outside the observed range** (15 to 40), so this is an extrapolation and should not be interpreted substantively.

> The key notes the same answer follows from $\hat\beta_1=\frac{\sum(x_i-\bar x)(y_i-\bar y)}{\sum(x_i-\bar x)^2}$ and $\hat\beta_0=\bar y-\hat\beta_1\bar x$. **Use whichever is faster** — with the inverse handed to you, the matrix route is two multiplications.

---

## Q7 🔴 — Sign of a coefficient vs sign of a correlation

> **W23, Ex 1a(i).** *"In linear regression, if the coefficient of a variable is positive, then there must be a positive correlation between that variable and the response variable."*

### **FALSE.**

$\hat\beta_j$ is a **partial** effect — the association with $y$ **holding all other covariates fixed**. The marginal correlation between $x_j$ and $y$ ignores those other covariates entirely. When covariates are correlated with each other, the two can carry **opposite signs** (Simpson's paradox in regression form).

> This is why every interpretation sentence must end with **"holding all other covariates fixed."** It is not padding — it is the claim.

---

## Q8 — Why one dummy's p-value is huge

> **W22, Ex 3 [1.5 Points].** `weathersit4` has estimate $-127.920$, std. error $71.699$, $t=-1.784$, $p=0.0744$. *"Explain why the p-value for `weathersit4` is comparatively high. Hint: Also consider Table 1."*

### Solution

Table 1 gives the category frequencies: `weathersit` = 1: 11,413 · 2: 4,544 · 3: 1,419 · **4: 3**.

**There are only 3 observations in category 4.** The standard error of a dummy's coefficient depends on how many observations identify it; with $n_4=3$ the estimate is extremely imprecise — $\widehat{\text{se}}=71.699$ against $2.193$ for `weathersit2`, roughly **33 times larger**. So even a large estimated effect ($-127.920$, the biggest coefficient in the model) yields $|t|=1.784$, below $t_{0.975}\approx1.96$, hence the large p-value.

> 🔑 **A big coefficient with a big p-value usually means a small subgroup, not a small effect.** "Not significant" ≠ "no effect" — here it means *we cannot tell*, on three observations.

---

# PART C — Interactions (EX20's 25-point exercise)

> **EX20, Exercise 2 [25 Points].** Final grade $X$ of master's theses for three groups (1 = naturalists, 2 = engineers, 3 = humanities), $n=30$; response $Y$ = weeks until first employment.

Estimates given in part (c):

| Term | Estimate |
|---|---|
| (Intercept) | 1.9505 |
| `Group2` | −1.1278 |
| `Group3` | 3.1035 |
| `X` (final grade) | 5.0195 |
| `Group2:X` | −3.0054 |
| `Group3:X` | 2.9523 |

## Q9(a) [4 P.] — Specify the model with interactions

$$y_i=\beta_0+\beta_1x_i+\beta_2\text{Group2}_i+\beta_3\text{Group3}_i+\beta_4(\text{Group2}_i\!\cdot\!x_i)+\beta_5(\text{Group3}_i\!\cdot\!x_i)+\varepsilon_i$$

$$\text{Group2}_i=\begin{cases}1&\text{engineering}\\0&\text{else}\end{cases}\qquad \text{Group3}_i=\begin{cases}1&\text{humanities}\\0&\text{else}\end{cases}$$

Reference = **Group 1, naturalists**. Design matrix, $n\times6$:

$$\boldsymbol{X}=\begin{pmatrix}
1&x_1&1&0&x_1&0\\
1&x_2&0&1&0&x_2\\
1&x_3&0&0&0&0\\
\vdots&\vdots&\vdots&\vdots&\vdots&\vdots
\end{pmatrix}$$

*(rows shown: an engineer, a humanities student, a naturalist)*

## Q9(b) [6 P.] — Group-wise equations, and effect coding

| Group | Equation | Combined |
|---|---|---|
| 1 naturalists | $\hat\beta_0+\hat\beta_1x$ | no combination needed |
| 2 engineers | $\hat\beta_0+\hat\beta_1x+\hat\beta_{G2}+\hat\beta_{G2:x}x$ | $\gamma_0+\gamma_1x$, $\ \gamma_0=\hat\beta_0+\hat\beta_{G2}$, $\gamma_1=\hat\beta_1+\hat\beta_{G2:x}$ |
| 3 humanities | $\hat\beta_0+\hat\beta_1x+\hat\beta_{G3}+\hat\beta_{G3:x}x$ | $\alpha_0+\alpha_1x$, $\ \alpha_0=\hat\beta_0+\hat\beta_{G3}$, $\alpha_1=\hat\beta_1+\hat\beta_{G3:x}$ |

**Under effect coding** $\text{Group}J_i\in\{1,-1,0\}$ ($-1$ for the naturalists): the **intercept becomes the grand mean** — the average response across all groups — and each group parameter becomes the **difference between that group's mean and the grand mean**, rather than a difference from a reference group.

## Q9(c) [4 P.] — Recover the humanities line

$$\text{slope}=\hat\beta_1+\hat\beta_{G3:x}=5.0195+2.9523=\boxed{7.9718}$$

$$\text{value at }x=1:\quad 1.9505+5.0195+3.1035+2.9523=13.0258$$

$$\hat y_{G3}=13.0258+7.9718(x-1)\qquad\text{or}\qquad \boxed{\hat y_{G3}=5.054+7.9718\,x}$$

## Q9(d) [6 P.] — Interpret all three lines

| Group | At grade 1.0 | Per **0.1** grade point (one decimal, i.e. slightly worse) |
|---|---|---|
| 1 naturalists | ≈ **6 weeks** | $+0.502$ weeks |
| 2 engineers | ≈ **2.5 weeks** | $+0.201$ weeks $(\hat\beta_1+\hat\beta_{G2:x})$ |
| 3 humanities | ≈ **12.5 weeks** | $+0.797$ weeks $(\hat\beta_1+\hat\beta_{G3:x})$ |

> **Note the units.** The key interprets per **decimal point** of grade, i.e. per 0.1 — so the per-unit slopes 5.020 / 2.014 / 7.972 are divided by 10. Read what the question's variable actually steps by.

## Q9(e) [2 P.] — Largest influence

**Group 3, humanities** — steepest slope (7.972), visible directly as the steepest line in the figure.

## Q9(f) [3 P.] — Predictions at grade 2.0

$$\hat y_{G1}=1.9505+5.0195(2.0)=\boxed{6.970}$$

$$\hat y_{G2}=1.9505+5.0195(2.0)+(-1.1278)+(-3.0054)(2.0)=11.9895-1.1278-6.0108=\boxed{4.851}$$

$$\hat y_{G3}=1.9505+5.0195(2.0)+3.1035+2.9523(2.0)=\boxed{20.998}$$

> ⚠️ **The official EX20 key prints 16.8725 for Group 2. That is an arithmetic slip in the key — it adds the two negative terms instead of subtracting them.** Check it against the key's *own* part (d): engineers start at ≈2.5 weeks at grade 1.0 and gain ≈0.201 weeks per decimal, so at grade 2.0 they should be near $2.5+2.01\approx4.5$. **4.851 is the consistent answer.** If you reproduce this question, trust the arithmetic, not the printed number — and show your working, which is what earns the marks anyway.

---

# PART D — The logit model

## Q10 🔴 — Why a linear model fails for binary $y$

> **S25, Ex 4(a) [1 Point].** *"Explain why a linear regression model is not appropriate to model a binary dependent variable and why you should instead use a logit, probit or similar model."*

### Solution

For binary $y$, $E(y\mid\boldsymbol{x})=P(y=1)=\pi\in[0,1]$. A linear model gives $\pi=\boldsymbol{x}'\boldsymbol\beta$, which fails on four counts:

1. 🔴 **Fitted values are unbounded** — $\boldsymbol{x}'\hat{\boldsymbol\beta}$ can fall below 0 or above 1, and a probability cannot. *(Lead with this one; it alone can carry the mark.)*
2. **Heteroscedasticity by construction:** $\text{Var}(y\mid\boldsymbol{x})=\pi(1-\pi)$ depends on $\boldsymbol{x}$, so A3 fails automatically.
3. **Errors cannot be normal:** given $\boldsymbol{x}$, $\varepsilon$ takes only the two values $1-\pi$ and $-\pi$.
4. **A constant marginal effect is implausible** near the boundaries — the same change in $x$ cannot keep pushing $\pi$ once it is near 0 or 1.

**The fix:** model $\pi=h(\boldsymbol{x}'\boldsymbol\beta)$ with a response function squashing the linear predictor into $(0,1)$ — the logistic $h(\eta)=e^\eta/(1+e^\eta)$ for logit, $\Phi(\eta)$ for probit. Both are fitted by **maximum likelihood**, not OLS.

## Q11 🔴🔴 — The logit interpretation trap

> **S25, Ex 1(h).** *"Consider a logit model where you regress $y_i$ onto $1,x_{1,i},\dots,x_{k,i}$ to obtain $\hat\beta_0,\dots,\hat\beta_k$. An increase of 1 in $x_{j,i}$ is interpreted as an increase of $\hat\beta_j$ in $P(y_i=1)$."*

### **FALSE.**

$\hat\beta_j$ is the effect on the **log-odds**, not on the probability:

$$\log\frac{\pi}{1-\pi}=\boldsymbol{x}'\boldsymbol\beta\ \Longrightarrow\ +1\text{ in }x_j\ \Longrightarrow\ +\hat\beta_j\text{ in log-odds}\ \Longrightarrow\ \text{odds}\times\exp(\hat\beta_j)$$

On the **probability** scale the effect is $\hat\beta_j\,\pi(1-\pi)$, which **depends on $\pi$** and is therefore different for every individual — largest near $\pi=0.5$, near zero at the extremes.

> 🔑 **Only the sign of $\hat\beta_j$ transfers unambiguously between scales.** Three scales, three different statements — name the scale you are on, every time.

| Scale | Effect of $+1$ in $x_j$ |
|---|---|
| log-odds | $+\hat\beta_j$ (exact, constant) |
| **odds** | $\times\exp(\hat\beta_j)$ — the **odds ratio** |
| probability | $\hat\beta_j\pi(1-\pi)$ — **not constant** |

---

## Chapter 2 scorecard

| Question type | Papers | Typical marks | Your target |
|---|---|---|---|
| Build model with dummies | S25, EX20 | 3–4 | **cold, 3 minutes** |
| Interpret a coefficient | LMES, W23, EX20 | 1–6 | **cold, with the two magic phrases** |
| Interaction: recover a group's line | EX20 | 4–6 | **cold** |
| Logit: why not linear / what does $\hat\beta$ mean | S25 ×2 | 1–2 | **word-perfect** |
| Dummy-count T/F | LMES, W23 | 0.5 each | **automatic** |

**These are the fastest marks on the paper.** Nothing here requires a calculator except Q6 and Q9(f).
