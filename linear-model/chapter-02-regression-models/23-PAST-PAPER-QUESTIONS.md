# Ch 2 — PAST-PAPER QUESTIONS (all five papers)

*বাংলা সংস্করণ নিচে আছে → [বাংলায় পড়ো](#অধ্যায়-২--বিগত-বছরের-প্রশ্ন-বাংলা)*

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

---
---

# অধ্যায় ২ — বিগত বছরের প্রশ্ন (বাংলা)

> টেকনিক্যাল শব্দ, ফাইলের নাম, সূত্র আর পরীক্ষার হুবহু উদ্ধৃতি ইংরেজিতেই রেখেছি — **পরীক্ষার উত্তর ইংরেজিতে লিখতে হবে**।
>
> **S25** = Exam Summer 2025 · **LMES** = Linear_model_exam_sheet · **W23** = WiSe 2023/24 · **W22** = RCLM WS 22/23 · **EX20** = Example Exam LiMo 2020

**অধ্যায় ২-এর দাম তার ~১৩% ভাগের চেয়ে অনেক বেশি, কারণ EX20 পুরো একটা ২৫ নম্বরের প্রশ্ন এতে দিয়েছে।** Dummy coding, interaction আর logit interpretation — **ব্যতিক্রম ছাড়াই প্রতিটা পেপারে** আছে।

| বিষয় | যেসব পেপারে |
|---|---|
| Dummy coding, $c-1$ নিয়ম | S25, LMES, W23, W22, EX20 |
| Interaction | EX20 (২৫ নম্বর), W22 |
| Coefficient-এর interpretation | LMES, W23, EX20 |
| Logit | S25 (দুইবার) |
| Log/polynomial রূপান্তর | S25, W23, W22 |

---

# পর্ব ক — মডেল বানানো আর dummy coding

## প্রশ্ন ১ 🔴 — দুটো categorical covariate দিয়ে মডেল

> **S25, Ex 2(a) [৩ নম্বর].** *"You want to examine the effect of a person's age, education (no degree, high school degree or college/university degree) and place of birth (inside the US or outside the US) on their hourly wage… Provide an adequate linear model equation that incorporates all of the above explanatory variables, including an intercept. Formulate it such that you can estimate this exact equation by OLS."*

### সমাধান

$$\text{wage}_i=\beta_0+\beta_1\,\text{age}_i+\beta_2 D^{\text{HS}}_i+\beta_3 D^{\text{college}}_i+\beta_4 D^{\text{outside}}_i+\varepsilon_i,\qquad i=1,\dots,n$$

যেখানে

$$D^{\text{HS}}_i=\begin{cases}1&\text{যদি }i\text{-এর high school degree থাকে}\\0&\text{নয়তো}\end{cases}\qquad
D^{\text{college}}_i=\begin{cases}1&\text{যদি college/university degree থাকে}\\0&\text{নয়তো}\end{cases}$$

$$D^{\text{outside}}_i=\begin{cases}1&\text{যদি }i\text{ US-এর বাইরে জন্মায়}\\0&\text{নয়তো}\end{cases}$$

**Reference category: *no degree* আর *born inside the US*।** Error $\varepsilon_i$ i.i.d., $E(\varepsilon_i)=0$, $\text{Var}(\varepsilon_i)=\sigma^2$।

$$p = 1+1+(3-1)+(2-1)=5$$

> **"formulate it such that you can estimate this exact equation by OLS" — এই বাক্যটা প্রশ্নে কেন আছে:** এটা তোমাকে বলছে **dummy-গুলো লিখে ফেলতে**। `education`-কে একটা প্রতীক হিসেবে লিখলে সেটা OLS দিয়ে estimate করা যায় না — categorical ভেরিয়েবল কোনো সংখ্যা নয়। **প্রশ্নকর্তা এই বাক্যেই পদ্ধতিটা হাতে তুলে দিচ্ছেন।**
>
> 🔴 তিনটা education level-এর জন্য তিনটা dummy **এবং** একটা intercept ⟹ কলামগুলো যোগ হয়ে intercept-এর কলাম হয়ে যায় ⟹ $\boldsymbol{X}'\boldsymbol{X}$ singular ⟹ **OLS-এর কোনো unique সমাধান নেই**। এটাই *dummy variable trap*, আর এ কারণেই $c-1$ নিয়মটা আছে।

---

## প্রশ্ন ২ 🔴🔴 — $c-1$ নিয়ম TRUE/FALSE হিসেবে

**দুটো পেপারে, একবার এদিক আর একবার ওদিক করে** — এই কোর্সের সবচেয়ে নিয়মিত পুনরাবৃত্ত T/F।

> **LMES, Ex 1a(iv).** *"In a regression model with a categorical predictor of $k$ levels, $k-1$ dummy variables are created to represent this predictor."* → ✅ **TRUE** *(অফিসিয়াল key: `FFTT`)*
>
> **W23, Ex 1a(iv).** *"For nominal variables with $m$ categories, we need to add $m$ dummy variables to represent this predictor when constructing the linear model."* → ❌ **FALSE** *(অফিসিয়াল key: `FFTF`)*

**একই তথ্য, উল্টো উত্তর — কারণ দ্বিতীয়টায় "−1" বাদ পড়েছে।** শেষ শব্দ পর্যন্ত পড়ো।

---

## প্রশ্ন ৩ — R output থেকে design পুনরুদ্ধার

> **W22, Ex 2** প্রেক্ষাপট। Output-এ আছে `season2, season3, season4`, `yr1`, `daytimemorning, daytimemidday, daytimeafternoon, daytimeevening`, `holiday1`, সাথে `temp`, `hum`, `windspeed`।

### যা তোমাকে পড়ে ফেলতে হবে

| ভেরিয়েবল | Level | ছাপা dummy | Reference (যে level **নেই** output-এ) |
|---|---|---|---|
| `season` | ৪ | ৩ | **season1 = winter** |
| `yr` | ২ | ১ | **yr0 = 2011** |
| `daytime` | ৫ | ৪ | **night** |
| `holiday` | ২ | ১ | **holiday0** |

$$p = 1 + 3 + 1 + 4 + 1 + 3\ \text{continuous} = 13$$

> 🔑 **Reference category মানেই সেই level যেটা output-এ দেখা যায় না।** ছাপা প্রতিটা coefficient তার **থেকে** পার্থক্য। দুটো **non-reference** level তুলনা করতে — যেমন summer বনাম spring — **বিয়োগ করো**: $18.977-37.108=-18.131$ বাইক প্রতি ঘণ্টায়।

---

## প্রশ্ন ৪ — Centred quadratic term কেন

> **S25, Ex 2(b) [২ নম্বর].** *"Explain briefly why adding the following regressor to the equation might be a sensible idea: $\text{gage2}_i := (\text{age}_i-48)^2$."*
>
> **S25, Ex 2(c) [১ নম্বর].** *"Would you expect the OLS estimate of the coefficient for $\text{gage2}_i$ to be positive or negative? Explain your choice."*

### সমাধান (b)

Wage-এর উপর age-এর প্রভাব **যুক্তিসঙ্গতভাবে linear নয়**: কর্মজীবনের শুরু ও মাঝামাঝি সময়ে বেতন বাড়ে, তারপর সমান হয়ে যায়, আর অবসরের কাছাকাছি কমে। Age-এ linear মডেল পুরো পরিসরে একটাই স্থির প্রভাব চাপিয়ে দেয়, ফলে মডেলটা **misspecified** হবে (A1 লঙ্ঘন — একমাত্র assumption যার ব্যর্থতা $\hat{\boldsymbol\beta}$-কে *biased* করে)।

একটা squared term যোগ করলে fitted curve বাঁকতে পারে, অথচ **এটা linear model-ই থেকে যায়** — গুরুত্বপূর্ণ হলো parameter-এ linear হওয়া, ভেরিয়েবলে নয়।

**৪৮-এ centre করা** (মোটামুটি কর্মজীবনের মাঝামাঝি) দুটো কাজ করে: $\text{age}$ আর $\text{age}^2$-এর মধ্যে collinearity কমায়, আর $\beta_1$-কে **৪৮ বছর বয়সে** marginal effect হিসেবে ব্যাখ্যাযোগ্য করে — অর্থহীন ০ বছর বয়সে নয়।

### সমাধান (c)

**নেগেটিভ।** Wage কর্মজীবনের মাঝামাঝি সর্বোচ্চ হয়ে দুই দিকেই কমে, তাই parabola-টা **নিচের দিকে** খুলতে হবে, অর্থাৎ $\hat\beta_{\text{gage2}}<0$। ৪৮-এ centre করায় fitted peak-টা ৪৮-এর কাছাকাছি বসবে।

> 🔴 Age দুটো term-এ থাকায় marginal effect হলো $\dfrac{\partial E(\text{wage})}{\partial\text{age}}=\beta_1+2\beta_2(\text{age}-48)$ — **কখনো একা $\beta_1$ উদ্ধৃত কোরো না**। Turning point $\text{age}=48-\hat\beta_1/(2\hat\beta_2)$-এ।

---

# পর্ব খ — Interpretation

## প্রশ্ন ৫ 🔴 — Log-রূপান্তরিত covariate-এর ব্যাখ্যা

> **LMES, Ex 3(c) [১ নম্বর].** *"What is the interpretation of the coefficients of the independent variable `Population`?"* — মডেল `Life.Exp ~ log(Population) + …`, $\hat\beta_{\log(\text{Pop})}=0.2707$।

### সমাধান — অফিসিয়াল key

> *"As population increases by 1%, the average life expectancy will increase by $0.01\times0.27=0.0027$ units, given other conditions unchanged."*

**Key-তে স্পষ্ট লেখা: "on average" বা "given other conditions unchanged" বাদ পড়লে আধা নম্বর কাটা।**

> 🔑 **Log-covariate নিয়ম:** $x$-এ **১% বৃদ্ধিতে** $\hat\beta/100$। ওই ১০০-র ভাগটাই পুরো প্রশ্ন।
>
> আর এর উল্টো ফাঁদ: **W23 Ex 1c(iv)** — *"In a linear regression model applying logarithmic transformation, the coefficients should be interpreted as the percentage change in the response for a 1% change in the predictor."* → ❌ **FALSE**। ওই elasticity-র পাঠের জন্য **দুই পাশেই** log লাগে। এখানে শুধু covariate-এ log, তাই response নড়ে **একক**-এ, শতাংশে নয়।

---

## প্রশ্ন ৬ — শূন্য থেকে intercept আর slope-এর ব্যাখ্যা

> **W23, Ex 3(a) [৩ নম্বর].** পাঁচটা জোড়া: $X=(20,30,33,40,15)$, $Y=(7,9,8,11,5)$. *"Build a regression model of consumption $Y$ on income $X$ with intercept and estimate the intercept and slope. How do you interpret the estimated intercept and slope?"* হিন্ট দেওয়া: $(\boldsymbol{X}'\boldsymbol{X})^{-1}=\begin{pmatrix}2.079&-0.068\\-0.068&0.0024\end{pmatrix}$

### সমাধান

$$\boldsymbol{X}'\boldsymbol{X}=\begin{pmatrix}n&\sum x_i\\\sum x_i&\sum x_i^2\end{pmatrix}=\begin{pmatrix}5&138\\138&4214\end{pmatrix},\qquad \boldsymbol{X}'\boldsymbol{Y}=\begin{pmatrix}\sum y_i\\\sum x_iy_i\end{pmatrix}=\begin{pmatrix}40\\1189\end{pmatrix}$$

$$\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{Y}=\begin{pmatrix}2.079&-0.068\\-0.068&0.0024\end{pmatrix}\begin{pmatrix}40\\1189\end{pmatrix}=\begin{pmatrix}2.210\\0.209\end{pmatrix}$$

$$\boxed{\hat y=2.210+0.209\,x}$$

**ব্যাখ্যা।** *Slope:* income-এ এক এককের বৃদ্ধি প্রত্যাশিত consumption-এ আনুমানিক **০.২০৯ একক** বৃদ্ধির সাথে যুক্ত। *Intercept:* $x=0$-তে প্রত্যাশিত consumption ২.২১০ — কিন্তু $x=0$ **পর্যবেক্ষিত পরিসরের বাইরে** (১৫ থেকে ৪০), তাই এটা extrapolation, এর বিষয়গত ব্যাখ্যা দেওয়া উচিত নয়।

> Key বলে একই উত্তর $\hat\beta_1=\frac{\sum(x_i-\bar x)(y_i-\bar y)}{\sum(x_i-\bar x)^2}$ আর $\hat\beta_0=\bar y-\hat\beta_1\bar x$ থেকেও আসে। **যেটা দ্রুত সেটাই ব্যবহার করো** — inverse হাতে দেওয়া থাকলে matrix পথে মাত্র দুটো গুণ।

---

## প্রশ্ন ৭ 🔴 — Coefficient-এর চিহ্ন বনাম correlation-এর চিহ্ন

> **W23, Ex 1a(i).** *"In linear regression, if the coefficient of a variable is positive, then there must be a positive correlation between that variable and the response variable."*

### **FALSE.**

$\hat\beta_j$ হলো একটা **partial** effect — $y$-এর সাথে সম্পর্ক, **বাকি সব covariate স্থির রেখে**। কিন্তু $x_j$ আর $y$-এর marginal correlation ওই বাকি covariate-গুলোকে পুরোপুরি উপেক্ষা করে। Covariate-গুলো নিজেদের মধ্যে correlated হলে দুটোর **চিহ্ন উল্টো** হতে পারে (regression-এর ভাষায় Simpson's paradox)।

> এজন্যই প্রতিটা interpretation বাক্য **"holding all other covariates fixed"** দিয়ে শেষ হতে হবে। এটা ভরাট করার কথা নয় — **এটাই দাবিটা**।

---

## প্রশ্ন ৮ — একটা dummy-র p-value এত বড় কেন

> **W22, Ex 3 [১.৫ নম্বর].** `weathersit4`-এর estimate $-127.920$, std. error $71.699$, $t=-1.784$, $p=0.0744$. *"Explain why the p-value for `weathersit4` is comparatively high. Hint: Also consider Table 1."*

### সমাধান

Table 1-এ category-র সংখ্যা দেওয়া আছে: `weathersit` = 1: ১১,৪১৩ · 2: ৪,৫৪৪ · 3: ১,৪১৯ · **4: ৩**।

**Category 4-এ মাত্র ৩টা observation।** একটা dummy-র coefficient-এর standard error নির্ভর করে কতগুলো observation সেটাকে চিহ্নিত করছে তার উপর; $n_4=3$ হলে estimate-টা ভয়ানক অনিশ্চিত — `weathersit2`-এর $2.193$-এর বিপরীতে এখানে $\widehat{\text{se}}=71.699$, প্রায় **৩৩ গুণ বড়**। তাই estimate বড় হওয়া সত্ত্বেও ($-127.920$, মডেলের সবচেয়ে বড় coefficient) $|t|=1.784$, যা $t_{0.975}\approx1.96$-এর নিচে — এজন্যই p-value বড়।

> 🔑 **বড় coefficient-এর সাথে বড় p-value মানে সাধারণত ছোট subgroup, ছোট effect নয়।** "Not significant" ≠ "কোনো প্রভাব নেই" — এখানে এর মানে *তিনটা observation দিয়ে আমরা বলতেই পারছি না*।

---

# পর্ব গ — Interaction (EX20-এর ২৫ নম্বরের প্রশ্ন)

> **EX20, Exercise 2 [২৫ নম্বর].** তিনটা গ্রুপের master's thesis-এর final grade $X$ (১ = naturalists, ২ = engineers, ৩ = humanities), $n=30$; response $Y$ = প্রথম চাকরি পর্যন্ত সপ্তাহ।

(c)-তে দেওয়া estimate:

| Term | Estimate |
|---|---|
| (Intercept) | 1.9505 |
| `Group2` | −1.1278 |
| `Group3` | 3.1035 |
| `X` (final grade) | 5.0195 |
| `Group2:X` | −3.0054 |
| `Group3:X` | 2.9523 |

## প্রশ্ন ৯(a) [৪ নম্বর] — Interaction সহ মডেল লেখা

$$y_i=\beta_0+\beta_1x_i+\beta_2\text{Group2}_i+\beta_3\text{Group3}_i+\beta_4(\text{Group2}_i\!\cdot\!x_i)+\beta_5(\text{Group3}_i\!\cdot\!x_i)+\varepsilon_i$$

$$\text{Group2}_i=\begin{cases}1&\text{engineering}\\0&\text{নয়তো}\end{cases}\qquad \text{Group3}_i=\begin{cases}1&\text{humanities}\\0&\text{নয়তো}\end{cases}$$

Reference = **Group 1, naturalists**। Design matrix, $n\times6$:

$$\boldsymbol{X}=\begin{pmatrix}
1&x_1&1&0&x_1&0\\
1&x_2&0&1&0&x_2\\
1&x_3&0&0&0&0\\
\vdots&\vdots&\vdots&\vdots&\vdots&\vdots
\end{pmatrix}$$

*(দেখানো সারি: একজন engineer, একজন humanities শিক্ষার্থী, একজন naturalist)*

## প্রশ্ন ৯(b) [৬ নম্বর] — গ্রুপভিত্তিক সমীকরণ, আর effect coding

| গ্রুপ | সমীকরণ | একত্রিত রূপ |
|---|---|---|
| ১ naturalists | $\hat\beta_0+\hat\beta_1x$ | একত্র করার দরকার নেই |
| ২ engineers | $\hat\beta_0+\hat\beta_1x+\hat\beta_{G2}+\hat\beta_{G2:x}x$ | $\gamma_0+\gamma_1x$, $\ \gamma_0=\hat\beta_0+\hat\beta_{G2}$, $\gamma_1=\hat\beta_1+\hat\beta_{G2:x}$ |
| ৩ humanities | $\hat\beta_0+\hat\beta_1x+\hat\beta_{G3}+\hat\beta_{G3:x}x$ | $\alpha_0+\alpha_1x$, $\ \alpha_0=\hat\beta_0+\hat\beta_{G3}$, $\alpha_1=\hat\beta_1+\hat\beta_{G3:x}$ |

**Effect coding-এ** $\text{Group}J_i\in\{1,-1,0\}$ (naturalists-দের জন্য $-1$): **intercept হয়ে যায় grand mean** — সব গ্রুপ মিলিয়ে response-এর গড় — আর প্রতিটা গ্রুপ parameter হয় **সেই গ্রুপের গড় আর grand mean-এর পার্থক্য**, কোনো reference গ্রুপ থেকে পার্থক্য নয়।

## প্রশ্ন ৯(c) [৪ নম্বর] — Humanities-এর রেখা পুনরুদ্ধার

$$\text{slope}=\hat\beta_1+\hat\beta_{G3:x}=5.0195+2.9523=\boxed{7.9718}$$

$$x=1\text{-এ মান}:\quad 1.9505+5.0195+3.1035+2.9523=13.0258$$

$$\hat y_{G3}=13.0258+7.9718(x-1)\qquad\text{অথবা}\qquad \boxed{\hat y_{G3}=5.054+7.9718\,x}$$

## প্রশ্ন ৯(d) [৬ নম্বর] — তিনটা রেখারই ব্যাখ্যা

| গ্রুপ | Grade 1.0-এ | প্রতি **০.১** grade point-এ (এক দশমিক, অর্থাৎ সামান্য খারাপ) |
|---|---|---|
| ১ naturalists | ≈ **৬ সপ্তাহ** | $+0.502$ সপ্তাহ |
| ২ engineers | ≈ **২.৫ সপ্তাহ** | $+0.201$ সপ্তাহ $(\hat\beta_1+\hat\beta_{G2:x})$ |
| ৩ humanities | ≈ **১২.৫ সপ্তাহ** | $+0.797$ সপ্তাহ $(\hat\beta_1+\hat\beta_{G3:x})$ |

> **এককের দিকে খেয়াল করো।** Key ব্যাখ্যা করেছে grade-এর প্রতি **দশমিক বিন্দু** ধরে, অর্থাৎ প্রতি ০.১ — তাই একক-প্রতি slope ৫.০২০ / ২.০১৪ / ৭.৯৭২-কে ১০ দিয়ে ভাগ করা হয়েছে। **প্রশ্নের ভেরিয়েবল আসলে কত ধাপে বাড়ছে, সেটা পড়ো।**

## প্রশ্ন ৯(e) [২ নম্বর] — সবচেয়ে বড় প্রভাব

**গ্রুপ ৩, humanities** — সবচেয়ে খাড়া slope (৭.৯৭২), চিত্রে সরাসরি সবচেয়ে খাড়া রেখা হিসেবেই দেখা যায়।

## প্রশ্ন ৯(f) [৩ নম্বর] — Grade 2.0-তে পূর্বাভাস

$$\hat y_{G1}=1.9505+5.0195(2.0)=\boxed{6.970}$$

$$\hat y_{G2}=1.9505+5.0195(2.0)+(-1.1278)+(-3.0054)(2.0)=11.9895-1.1278-6.0108=\boxed{4.851}$$

$$\hat y_{G3}=1.9505+5.0195(2.0)+3.1035+2.9523(2.0)=\boxed{20.998}$$

> ⚠️ **অফিসিয়াল EX20 key-তে Group 2-এর জন্য ছাপা আছে 16.8725। সেটা key-র নিজের হিসাবের ভুল** — দুটো ঋণাত্মক term বিয়োগ না করে যোগ করা হয়েছে। Key-র **নিজের** (d) অংশের সাথে মিলিয়ে দেখো: engineers grade 1.0-তে ≈২.৫ সপ্তাহে শুরু করে প্রতি দশমিকে ≈০.২০১ বাড়ে, তাই grade 2.0-তে হওয়ার কথা $2.5+2.01\approx4.5$-এর কাছাকাছি। **সামঞ্জস্যপূর্ণ উত্তর ৪.৮৫১।** এই প্রশ্ন আবার করলে ছাপা সংখ্যাটাকে নয়, **হিসাবটাকে বিশ্বাস করো** — আর কাজ দেখিয়ে লেখো, নম্বর তো ওখানেই।

---

# পর্ব ঘ — Logit model

## প্রশ্ন ১০ 🔴 — Binary $y$-এর জন্য linear model কেন ব্যর্থ

> **S25, Ex 4(a) [১ নম্বর].** *"Explain why a linear regression model is not appropriate to model a binary dependent variable and why you should instead use a logit, probit or similar model."*

### সমাধান

Binary $y$-এর জন্য $E(y\mid\boldsymbol{x})=P(y=1)=\pi\in[0,1]$। Linear model দেয় $\pi=\boldsymbol{x}'\boldsymbol\beta$, যা চার জায়গায় ব্যর্থ:

1. 🔴 **Fitted value-এর কোনো সীমা নেই** — $\boldsymbol{x}'\hat{\boldsymbol\beta}$ ০-এর নিচে বা ১-এর উপরে যেতে পারে, আর probability তা পারে না। *(এটা দিয়েই শুরু করো; একাই নম্বরটা এনে দিতে পারে।)*
2. **গঠনগতভাবেই heteroscedasticity:** $\text{Var}(y\mid\boldsymbol{x})=\pi(1-\pi)$ নির্ভর করে $\boldsymbol{x}$-এর উপর, তাই A3 আপনা-আপনি ভাঙে।
3. **Error normal হতেই পারে না:** $\boldsymbol{x}$ দেওয়া থাকলে $\varepsilon$ মাত্র দুটো মান নেয় — $1-\pi$ আর $-\pi$।
4. **স্থির marginal effect অবাস্তব** — $\pi$ ০ বা ১-এর কাছে গেলে $x$-এর একই পরিবর্তন আর $\pi$-কে ঠেলতে পারে না।

**সমাধান:** $\pi=h(\boldsymbol{x}'\boldsymbol\beta)$ মডেল করো, যেখানে response function linear predictor-কে $(0,1)$-এ চেপে ঢোকায় — logit-এর জন্য logistic $h(\eta)=e^\eta/(1+e^\eta)$, probit-এর জন্য $\Phi(\eta)$। দুটোই **maximum likelihood** দিয়ে fit হয়, OLS দিয়ে নয়।

## প্রশ্ন ১১ 🔴🔴 — Logit interpretation-এর ফাঁদ

> **S25, Ex 1(h).** *"Consider a logit model where you regress $y_i$ onto $1,x_{1,i},\dots,x_{k,i}$ to obtain $\hat\beta_0,\dots,\hat\beta_k$. An increase of 1 in $x_{j,i}$ is interpreted as an increase of $\hat\beta_j$ in $P(y_i=1)$."*

### **FALSE.**

$\hat\beta_j$ হলো **log-odds**-এর উপর প্রভাব, probability-র উপর নয়:

$$\log\frac{\pi}{1-\pi}=\boldsymbol{x}'\boldsymbol\beta\ \Longrightarrow\ x_j\text{-এ}+1\ \Longrightarrow\ \text{log-odds-এ}+\hat\beta_j\ \Longrightarrow\ \text{odds}\times\exp(\hat\beta_j)$$

**Probability** স্কেলে প্রভাব $\hat\beta_j\,\pi(1-\pi)$, যা **$\pi$-এর উপর নির্ভর করে** — তাই প্রত্যেক ব্যক্তির জন্য আলাদা, $\pi=0.5$-এর কাছে সবচেয়ে বড়, প্রান্তে প্রায় শূন্য।

> 🔑 **শুধু $\hat\beta_j$-এর চিহ্নটাই স্কেল বদলালেও অবিকৃত থাকে।** তিনটা স্কেল, তিনটা আলাদা বক্তব্য — **প্রতিবার তুমি কোন স্কেলে আছ সেটার নাম বলো।**

| স্কেল | $x_j$-এ $+1$-এর প্রভাব |
|---|---|
| log-odds | $+\hat\beta_j$ (সঠিক, স্থির) |
| **odds** | $\times\exp(\hat\beta_j)$ — **odds ratio** |
| probability | $\hat\beta_j\pi(1-\pi)$ — **স্থির নয়** |

---

## অধ্যায় ২-এর স্কোরকার্ড

| প্রশ্নের ধরন | পেপার | সাধারণ নম্বর | তোমার লক্ষ্য |
|---|---|---|---|
| Dummy দিয়ে মডেল বানানো | S25, EX20 | ৩–৪ | **ঠান্ডা মাথায়, ৩ মিনিট** |
| Coefficient interpret করা | LMES, W23, EX20 | ১–৬ | **ঠান্ডা মাথায়, দুটো জাদু-বাক্য সহ** |
| Interaction: গ্রুপের রেখা বের করা | EX20 | ৪–৬ | **ঠান্ডা মাথায়** |
| Logit: linear কেন নয় / $\hat\beta$-এর মানে | S25 ×২ | ১–২ | **অক্ষরে অক্ষরে** |
| Dummy-গোনার T/F | LMES, W23 | ০.৫ করে | **স্বয়ংক্রিয়** |

**পুরো পেপারে সবচেয়ে দ্রুত পাওয়া নম্বর এগুলোই।** প্রশ্ন ৬ আর ৯(f) ছাড়া এখানে কোথাও ক্যালকুলেটর লাগে না।
