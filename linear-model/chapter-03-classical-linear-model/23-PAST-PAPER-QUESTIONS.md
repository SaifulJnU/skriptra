# Ch 3 — PAST-PAPER QUESTIONS (all five papers)

> **S25** = Exam Summer 2025 · **LMES** = Linear_model_exam_sheet · **W23** = WiSe 2023/24 · **W22** = RCLM WS 22/23 · **EX20** = Example Exam LiMo 2020
>
> ⚠️ LMES and W23 use $p$ = **covariates** (so $\text{df}=n-p-1$); W22 and the book use $p$ = **parameters** (so $\text{df}=n-p$). Same number. **Count the betas including the intercept and subtract from $n$.**

**This is the file. Roughly 85% of every paper lives here.** Sections below follow the book's order; the source paper is cited on every question so you can cross-check against the PDF.

---

# 3.1 — MODEL DEFINITION AND ASSUMPTIONS

## Q1 — When is $\boldsymbol{X}'\boldsymbol{X}$ invertible?

> **LMES, Ex 2(b) [2 Points].** *"Write down the matrix $\boldsymbol{X}'\boldsymbol{X}$. Under which conditions is the matrix $\boldsymbol{X}'\boldsymbol{X}$ invertible?"* (simple linear regression)

### Solution

$$\boldsymbol{X}'\boldsymbol{X}=\begin{pmatrix}1&1&\cdots&1\\x_1&x_2&\cdots&x_n\end{pmatrix}\begin{pmatrix}1&x_1\\1&x_2\\\vdots&\vdots\\1&x_n\end{pmatrix}=\begin{pmatrix}n&\sum_{i}x_i\\\sum_i x_i&\sum_i x_i^2\end{pmatrix}\quad\text{[1 pt]}$$

**Condition:** $\boldsymbol{X}$ must have **full column rank**, $\text{rank}(\boldsymbol{X})=p$ — equivalently $\boldsymbol{X}'\boldsymbol{X}$ is non-singular. **[1 pt]**

In the simple case this fails exactly when all $x_i$ are identical (then $\det=n\sum x_i^2-(\sum x_i)^2=n\sum(x_i-\bar x)^2=0$): with no variation in $x$ there is no slope to identify.

## Q2 🔴 — Rank as TRUE/FALSE, four times

| Paper | Statement | Verdict | Why |
|---|---|---|---|
| **LMES 1a(i)** | *"$\text{rank}(\boldsymbol{X}'\boldsymbol{X})=p$, with the number of **variables** $p$"* | ❌ **FALSE** | In that paper $p$ = covariates, so the rank is $p+1$ — the definition given in the sentence is what breaks it |
| **W22 1a(i)** | *"$\text{rk}(\boldsymbol{X}'\boldsymbol{X})=p$"* (bare) | ✅ **TRUE** | Book notation, $p$ = parameters |
| **W22 1b(ii)** | *"$\text{rk}(\boldsymbol{X}'\boldsymbol{X})=k$"* | ❌ **FALSE** | one short — forgot the intercept |
| **W22 1c(iv)** | *"$\text{rk}(\boldsymbol{X}'\boldsymbol{X})=n$"* | ❌ **FALSE** | $\boldsymbol{X}'\boldsymbol{X}$ is $p\times p$; rank can never exceed the smaller dimension |

> 🔴 **Two identical-looking statements, opposite answers, because one of them defines $p$ inside the sentence.** Read the definition the paper gives before you answer.

## Q3 🔴 — Rank deficiency with an irrelevant escape clause

> **S25, Ex 1(d).** *"When the design matrix $\boldsymbol{X}$ does not have full column rank, the OLS estimates still exist and are unique as long as the error variance is constant."*

### **FALSE.**

Without full rank, $\boldsymbol{X}'\boldsymbol{X}$ is singular and the normal equations $\boldsymbol{X}'\boldsymbol{X}\hat{\boldsymbol\beta}=\boldsymbol{X}'\boldsymbol{y}$ have **infinitely many** solutions. Homoscedasticity is about **efficiency**, and has nothing whatever to do with identification.

> 🛟 **The "as long as…" construction.** Ask of every such clause: *does this condition have any bearing on the claim?* Here it does not, and that is the whole trap.

## Q4 — Automatic properties with an intercept

> **W22, Ex 1a(ii).** *"The average of the predicted values is equal to the average of the observed response."*

### **TRUE** — when the model contains an intercept.

$\sum_i\hat\varepsilon_i=0$ follows from the intercept's normal equation, hence $\bar{\hat y}=\bar y$, and the fitted line passes through $(\bar x,\bar y)$.

> ⚠️ **True by construction, not by good fit.** See the mirror trap at Q23 below.

---

# 3.2 — ESTIMATION

## Q5 🔴🔴 — Derive the OLS estimator

**Asked on three of the five papers.** Learn this once; collect it every year.

> **S25, Ex 4(b) [2 Points].** *"Explain the method of ordinary least squares (OLS) estimation and show the steps necessary to obtain the estimators $\hat\beta_0,\dots,\hat\beta_k$. It is not necessary to explicitly calculate them; it suffices to show the mathematical approach."*
>
> **LMES, Ex 2(c) [3 Points].** *"Derive the least square estimators in matrix form."*
>
> **W23, Ex 2(b) [2 Points].** *"Explain how the ordinary least squares (OLS) method is used to estimate the coefficients $\beta_0$ and $\beta_1$."*

### Solution

$$\boldsymbol\varepsilon=\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta,\qquad \text{RSS}=\boldsymbol\varepsilon'\boldsymbol\varepsilon=(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)=\boldsymbol{y}'\boldsymbol{y}-2\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{y}+\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{X}\boldsymbol\beta\quad\text{[1 pt]}$$

$$\frac{\partial\,\text{RSS}}{\partial\boldsymbol\beta}=-2\boldsymbol{X}'\boldsymbol{y}+2\boldsymbol{X}'\boldsymbol{X}\boldsymbol\beta\overset{!}{=}\boldsymbol{0}\quad\text{[1 pt]}$$

$$\boldsymbol{X}'\boldsymbol{X}\hat{\boldsymbol\beta}=\boldsymbol{X}'\boldsymbol{y}\ \Longrightarrow\ \boxed{\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}}\quad\text{[1 pt]}$$

Inversion requires $\text{rank}(\boldsymbol{X})=p$. The Hessian $2\boldsymbol{X}'\boldsymbol{X}$ is positive definite, so this is a **minimum**.

> **W23's marking key, verbatim:** *"1 point for correctly state that RSS needs to be minimized. And 1 point for correctly derive the solution."*
>
> 🔑 **Say what you are minimising before you differentiate.** Half the marks are for naming the objective — students who dive straight into the algebra lose a mark they already knew.

## Q6 — Compute $\hat{\boldsymbol\beta}$, RSS and $\hat\sigma^2$ by hand

> **W23, Ex 3(a)+(b) [3+2 Points].** $X=(20,30,33,40,15)$, $Y=(7,9,8,11,5)$, $n=5$, with $(\boldsymbol{X}'\boldsymbol{X})^{-1}=\begin{pmatrix}2.079&-0.068\\-0.068&0.0024\end{pmatrix}$ given.

### Solution

$$\boldsymbol{X}'\boldsymbol{X}=\begin{pmatrix}5&138\\138&4214\end{pmatrix},\quad \boldsymbol{X}'\boldsymbol{Y}=\begin{pmatrix}40\\1189\end{pmatrix},\quad \hat{\boldsymbol\beta}=\begin{pmatrix}2.079&-0.068\\-0.068&0.0024\end{pmatrix}\begin{pmatrix}40\\1189\end{pmatrix}=\begin{pmatrix}2.210\\0.209\end{pmatrix}$$

$$\hat{\boldsymbol{y}}=\boldsymbol{X}\hat{\boldsymbol\beta}=(6.40,\ 8.50,\ 9.13,\ 10.60,\ 5.35)'$$

$$\hat{\boldsymbol\varepsilon}=\boldsymbol{y}-\hat{\boldsymbol{y}}=(0.594,\ 0.496,\ -1.132,\ 0.398,\ -0.356)'$$

$$\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}=2.169,\qquad \hat\sigma^2=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n-2}=\frac{2.169}{3}=\boxed{0.723}$$

> **Denominator $n-2$**, because two parameters were estimated. This is $n-p$ in book notation.

## Q7 🔴 — Gauss–Markov, for 4 marks

> **W23, Ex 2(a) [4 Points].** *"Briefly describe the main contents of the Gauss–Markov Theorem and the assumptions."*
> **Official key: *"1 point for every assumption."***

### Solution — write it as a list, not as prose

> Under the assumptions
> **(i)** the model is correctly specified, $E(\boldsymbol{y})=\boldsymbol{X}\boldsymbol\beta$;
> **(ii)** $E(\boldsymbol\varepsilon)=\boldsymbol{0}$;
> **(iii)** homoscedasticity, $\text{Var}(\varepsilon_i)=\sigma^2$ for all $i$;
> **(iv)** uncorrelated errors, $\text{Cov}(\varepsilon_i,\varepsilon_j)=0$ for $i\neq j$ — (iii) and (iv) together being $\text{Cov}(\boldsymbol\varepsilon)=\sigma^2\boldsymbol{I}_n$;
> **(v)** $\text{rank}(\boldsymbol{X})=p$,
>
> the OLS estimator $\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$ is **BLUE** — among all estimators **linear in $\boldsymbol{y}$** and **unbiased** for $\boldsymbol\beta$, it has minimum variance.
>
> **Normality of the errors is *not* required.**

> 🔴 The key pays **per assumption**, so enumerate them. And the normality sentence is a real mark, because two separate T/F items test exactly that (Q8 below).

## Q8 🔴 — BLUE, three ways

| Paper | Statement | Verdict |
|---|---|---|
| **S25 1(e)** | *"A BLUE is 'best' in the sense that there is no other **linear unbiased** estimator with a lower variance."* | ✅ **TRUE** — correctly qualified |
| **W23 1a(ii)** | *"The LS estimator is BLUE **if and only if** the error term is expected to be zero and has constant variance."* | ❌ **FALSE** — the "iff" is too strong, and the list omits uncorrelated errors, correct specification and full rank |
| **S25 1(l)** | *"The OLS estimator is equivalent to the ML estimator under i.i.d. normal errors."* | ✅ **TRUE** — for $\boldsymbol\beta$. ⚠️ but $\hat\sigma^2_{ML}\neq\hat\sigma^2_{LS}$ |

> 🔑 Drop **"linear"** or **"unbiased"** from the BLUE statement and it becomes false — ridge is biased with smaller variance.

## Q9 🔴 — Unbiasedness under a useless extra variable

> **S25, Ex 1(a).** *"Adding a variable which is not correlated with the dependent variable will not affect the unbiasedness of the OLS estimator, but it may affect its variance."*

### **TRUE.**

Including an irrelevant regressor leaves $\hat{\boldsymbol\beta}$ unbiased (the model is still correctly specified — its true coefficient is simply zero), but costs a degree of freedom and, if the new variable is correlated with existing ones, **inflates** the variances.

**Contrast with omission:** leaving out a *relevant* covariate that is correlated with the included ones **does** bias $\hat{\boldsymbol\beta}$. Inclusion of junk costs precision; omission of substance costs correctness.

## Q10 — Omitted-variable bias, worked numerically

> **EX20, Ex 1(c)(d)(e) [4+6+4+3 Points].** Model: `points ~ goals`, $\hat\beta_{\text{goals}}=0.90509$. Two auxiliary models are supplied: `goals.received ~ goals` gives $-0.5850$ (significant), and `points ~ goals.received` gives $-0.9096$ (significant). Averages: $\overline{\text{goals}}=48.61$, $\overline{\text{points}}=46.61$.

### (c) What is she worried about? [4 P.]

**Bias in $\hat{\boldsymbol\beta}$** — that the estimator's expectation is not the true parameter. Unbiased means $E(\hat{\boldsymbol\beta})=\boldsymbol\beta$; biased means $E(\hat{\boldsymbol\beta})\neq\boldsymbol\beta$.

### (d) Is she right? [6 P.]

**Yes.** `goals.received` (i) plausibly influences `points` — the second output shows a significant coefficient — and (ii) is **correlated with `goals`**, since the first output shows `goals` significantly predicts `goals.received`. **Both conditions are needed.** An omitted variable that is uncorrelated with every included covariate leaves the estimates unbiased; here both hold, so the main model's estimate is biased.

### (e) How large is the bias? And find $\hat\beta_0$ for the full model. [4+3 P.]

$$\text{bias}=\hat\beta_{\text{goals}\to\text{goals.rec}}\times\beta_{\text{goals.rec}}=(-0.585)\times(-0.45763)=\boxed{0.268}$$

$$\beta_{\text{goals}}=0.90509-0.268=\boxed{0.637}$$

Both averages are 48.61 (every goal is scored by one team and received by another), so:

$$\hat\beta_0=46.61-\left[(0.637-0.45763)\times48.61\right]=46.61-8.719=\boxed{37.891}$$

> 🔑 **The bias is the product of two paths:** how the omitted variable moves with the included one, times the omitted variable's own true effect. Sign of the bias = product of the two signs — here negative × negative = **positive**, so the naive estimate was **too large**.

---

# 3.3 — TESTING AND INTERVALS

## Q11 🔴🔴 — Fill the missing R output

**The single most repeated computational question. It appears on S25, LMES and EX20.**

> **S25, Ex 3(a) [2.5 Points].** Reproduce `[[A]]`–`[[D]]` and write the fitted regression formula.
>
> ```
>              Estimate   Std. Error  t value
> (Intercept)  48.38458   [[A]]       13.591
> crim         -0.25959   0.05302     -4.896
> nox         -36.99122   5.25574     [[B]]
> dis          [[C]]      0.26423     -3.796
> rad          -0.06165   0.05983     -1.030
> Residual standard error: [[D]] on 501 degrees of freedom
> ```
> with $\sum\hat\varepsilon_i^2=31682.02$.

### Solution

Everything comes from **one identity**, $t=\hat\beta/\widehat{\text{se}}$, rearranged three ways:

$$\text{[[A]]}=\frac{\hat\beta}{t}=\frac{48.38458}{13.591}=\boxed{3.560}$$

$$\text{[[B]]}=\frac{\hat\beta}{\widehat{\text{se}}}=\frac{-36.99122}{5.25574}=\boxed{-7.038}$$

$$\text{[[C]]}=t\times\widehat{\text{se}}=-3.796\times0.26423=\boxed{-1.003}$$

$$\text{[[D]]}=\hat\sigma=\sqrt{\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n-p}}=\sqrt{\frac{31682.02}{501}}=\sqrt{63.238}=\boxed{7.952}$$

$$\widehat{\text{medv}}=48.385-0.260\,\text{crim}-36.991\,\text{nox}-1.003\,\text{dis}-0.062\,\text{rad}$$

> 🔴 **[[D]] is a standard *error*** — take the root. Skipping it gives 63.238 instead of 7.952.
> 🔴 $n=506$ here: $501$ df $+\ p=5$ parameters. The df line always lets you recover $n$.

### The same skill, two more times

> **LMES, Ex 3(b) [4 Points].** $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}=22.84961$, $n=50$, 7 covariates, $\bar R^2=0.6981$.
>
> $$\text{(A)}\ t=\frac{-0.298}{0.045}=-6.587\qquad \text{(B)}\ \hat\sigma=\sqrt{\frac{22.849}{42}}=\sqrt{0.544}=0.738$$
> $$\text{(C)}\ R^2=1-\frac{n-p-1}{n-1}(1-\bar R^2)=1-\frac{42}{49}(0.3019)=0.741$$
> $$\text{(D)}\ F=\frac{R^2}{1-R^2}\cdot\frac{n-p-1}{p}=\frac{0.7412}{0.2588}\cdot\frac{42}{7}=17.186$$

> **EX20, Ex 1(a) [4+4+4 Points].** `points ~ goals`, $n=18$, $\overline{\text{goals}}=48.61$, $\overline{\text{points}}=46.61$, $\hat\beta_{\text{goals}}=0.90509$, $t_{\text{goals}}=9.562$, $\widehat{\text{se}}_0=4.81560$.
>
> $$\text{A}=\hat\beta_0=\bar y-\hat\beta_1\bar x=46.61-0.90509(48.61)=46.61-43.996=\boxed{2.614}$$
> $$\text{B}=\widehat{\text{se}}_1=\frac{0.90509}{9.562}=\boxed{0.095}\qquad \text{C}=t_0=\frac{2.614}{4.8156}=\boxed{0.543}$$
>
> **(b) [6 P.]** The p-value $D$ is $2\big(1-F_{t_{16}}(9.562)\big)$ — R: `(1 - pt(q = 9.562, df = 16)) * 2`. Degrees of freedom $=n-p-1=18-1-1=16$. **The factor 2 is because the test is two-sided** and the $t$-distribution is symmetric, so both tails count.

> 🔑 **Three papers, one toolkit:** $t=\hat\beta/\widehat{\text{se}}$ · $\hat\beta_0=\bar y-\hat\beta_1\bar x$ · $\hat\sigma^2=\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}/(n-p)$ · $\bar R^2\leftrightarrow R^2$ · $F\leftrightarrow R^2$. Nothing else has ever been needed.

## Q12 — Full t-test, written out for marks

> **LMES, Ex 3(d) [2 Points].** *"Conduct a significance test for `Murder`. Clearly write the null and alternative hypotheses, your test statistic, degrees of freedom, critical value, and draw the correct conclusion. $\alpha=0.05$."*

### Solution — reproduce this skeleton every time

| Step | |
|---|---|
| $H_0$ | $\beta_{\text{Murder}}=0$ |
| $H_1$ | $\beta_{\text{Murder}}\neq0$ |
| Distribution | $t$ |
| Statistic | $t=\dfrac{-0.298}{0.045}=-6.587$ |
| df | $n-p-1=50-7-1=42$ |
| Critical value | $t_{0.975,42}=2.0180$ |
| Decision | $|-6.587|>2.0180$ ⟹ **reject $H_0$** |
| Conclusion | `Murder` is significantly associated with life expectancy at the 5% level |

> The key states: *"if the student is confused to use other degrees of freedom or other significance level, −0.5."* **The quantile is $1-\alpha/2=0.975$, not $0.95$.**
>
> **W23, Ex 3(c)** is the same skeleton on $n=5$: $\widehat{\text{se}}(\hat\beta_1)=\sqrt{\hat\sigma^2\left[(\boldsymbol{X}'\boldsymbol{X})^{-1}\right]_{22}}=\sqrt{0.723\times0.0024}=0.042$, $t=0.20977/0.04224=4.966$, df $=3$, $t_{0.975,3}=3.182$ ⟹ reject.
>
> ⚠️ The W23 key prints the variance vector as $(1.504, 0.084)$; $0.0024\times0.723=0.0017$, and the key's own next line ($\text{sd}=0.042=\sqrt{0.0017}$) confirms $0.084$ is a typo. The standard error 0.042 is correct.

## Q13 — Confidence intervals, three papers

> **S25, Ex 3(b) [2 Points].** 99% CI for $\beta_{\text{nox}}$; then test $H_0:\beta_{\text{nox}}=-30$ at the 1% level.

$$-36.99122\pm t_{501}(0.995)\times5.25574=-36.99122\pm2.5857(5.25574)=-36.991\pm13.590$$

$$\boxed{[-50.581,\ -23.401]}$$

**$-30$ lies inside the interval ⟹ do not reject $H_0:\beta_{\text{nox}}=-30$** at the 1% level.

> 🔴 99% ⟹ the **0.995** quantile. And note the answer is "fail to reject" even though nox is wildly significant against **zero** — the null value matters, not the variable.

> **LMES, Ex 3(g) [2 Points].** 95% CI for `HS.Grad`: $0.0584\pm2.0180(0.0242)=[0.0095,\ 0.1073]$. Consistent with the output's significance because **zero is not contained** ⟹ reject $H_0:\beta=0$ at 5%.
>
> **W23, Ex 3(d) [1 Point].** $0.20977\pm3.182(0.04224)=[0.0753,\ 0.3442]$.
>
> **W22, Ex 2(f) [3.5 Points].** 99% CI for $\beta_{12}$ (windspeed): $-42.244\pm2.5761(7.938)=-42.244\pm20.449=[-62.693,\ -21.795]$.

> 🔑 **CI–test duality is worth free marks:** if the question asks "is this consistent with the regression table?", the answer is always about whether the null value sits inside.

## Q14 🔴🔴 — Build $\boldsymbol{C}$ and $\boldsymbol{d}$

> **S25, Ex 3(c) [2 Points].** *"Test the joint null $H_0:\beta_{\text{crim}}=3\beta_{\text{rad}}-0.1,\ \beta_{\text{nox}}=-40$. Express as $\boldsymbol{C}\boldsymbol\beta=\boldsymbol{d}$ with $\boldsymbol\beta=(\beta_0,\beta_{\text{crim}},\beta_{\text{nox}},\beta_{\text{dis}},\beta_{\text{rad}})'$. What is the number of linearly independent restrictions $r$?"*

### Solution

**Rearrange first** — parameters left, constants right:

$$\beta_{\text{crim}}-3\beta_{\text{rad}}=-0.1,\qquad \beta_{\text{nox}}=-40$$

$$\boldsymbol{C}=\begin{pmatrix}0&1&0&0&-3\\0&0&1&0&0\end{pmatrix},\qquad \boldsymbol{d}=\begin{pmatrix}-0.1\\-40\end{pmatrix},\qquad \boxed{r=2}$$

> 🔴 $\boldsymbol{C}$ is $r\times p$ — the $\beta_0$ column is zeros but **must be there**.
> 🔴 The constant $-0.1$ goes into $\boldsymbol{d}$, never into $\boldsymbol{C}$.

### The reverse direction

> **W22, Ex 2(d) [1 Point].** Given $\boldsymbol{C}$ = three rows selecting positions 11, 12, 13 and $\boldsymbol{d}=\boldsymbol{0}_3$, state the hypotheses.
>
> $$H_0:\beta_{\text{temp}}=\beta_{\text{hum}}=\beta_{\text{windspeed}}=0\qquad\text{vs}\qquad H_1:\text{at least one}\neq0$$
>
> Read $\boldsymbol{C}$ **column by column**: each 1 marks the parameter that row constrains. Then match the column position against the R output's coefficient order.

## Q15 🔴🔴 — Counting restrictions

> **S25, Ex 1(i).** *"The F-statistic for testing $H_0:\beta_1=-\beta_2+\beta_3$ in a linear model with $k\geq3$ predictors plus an intercept has an F-distribution with $(3,\,n-k-1)$ degrees of freedom under $H_0$."*

### **FALSE.**

Rewrite as $\beta_1+\beta_2-\beta_3=0$ and **count the equals signs: one**. So $r=1$ and $F\sim F_{1,\,n-k-1}$.

The "3" counts the **betas mentioned**, which is never what $r$ means.

> **Related df traps:** **W22 1b(iv)** *"$F\sim F_{p+1,n-p}$ for $H_0:\boldsymbol{C}\boldsymbol\beta=\boldsymbol{d}$"* → **FALSE** (numerator is $r$). **W22 1c(ii)** *"$F\sim F_{p-1,n-p}$ for the overall test"* → **FALSE** (numerator is $k$, and $p-1=k$ only in book notation — the paper is inconsistent, so read its own definitions). **W23 1a(iii)** *"full-model F test $\sim F_{k,n-k-1}$"* → ✅ **TRUE**. **S25 1(g)** *"$t_j\sim t_{n-k-1}$ where $k+1$ is the number of estimated parameters"* → ✅ **TRUE**.

## Q16 — F-test from $\Delta$SSE

> **S25, Ex 3(d) [1.5 Points].** $\hat{\boldsymbol\varepsilon}'_{H_0}\hat{\boldsymbol\varepsilon}_{H_0}=32333.15$, unrestricted $\text{SSE}=31682.02$, $r=2$, 501 df.

$$F=\frac{(\text{SSE}_{H_0}-\text{SSE})/r}{\text{SSE}/(n-p)}=\frac{(32333.15-31682.02)/2}{31682.02/501}=\frac{325.565}{63.238}=\boxed{5.148}$$

$5.148>F_{2,501}(0.95)=3.0137$ ⟹ **reject $H_0$**. **At least one** of the two restrictions fails.

> **W22, Ex 2(e) [8 Points].** $\Delta\text{SSE}=26731881$, $r=3$, estimated model standard error $120.3$ so $\hat\sigma^2=120.3^2=14472.09$:
>
> $$F=\frac{\Delta\text{SSE}/r}{\hat\sigma^2}=\frac{26731881/3}{14472.09}=\frac{8910627}{14472.09}=615.72$$
>
> Quantile: $F_{3,\,17366}(0.99)=3.7827$. $615.72\gg3.783$ ⟹ **reject**. **Interpretation:** at least one of temperature, humidity and wind speed is significantly associated with hourly bike rentals, given the other covariates.

> 🔴 **Quantile rule:** the F-test is **one-sided** ($\text{SSE}_{H_0}\geq\text{SSE}$ always), so use $1-\alpha$ — $0.95$ at 5%, $0.99$ at 1%. Not $1-\alpha/2$.
> ✅ **Sanity checks:** $F\geq0$ always; if $r=1$ then $\sqrt F=|t|$.

## Q17 — Overall test for significance of regression

> **W22, Ex 2(c) [7 Points].** $n=17379$, 13 parameters, $R^2=0.5604$, $\alpha=0.01$. Test statistic, quantile, decision, interpretation.

$$F=\frac{R^2/k}{(1-R^2)/(n-p)}=\frac{0.5604/12}{0.4396/17366}=\frac{0.04670}{0.00002531}=\boxed{1844.8}\ \sim F_{12,\,17366}$$

Quantile $F_{12,17366}(0.99)=2.1858$. Since $1844.8\gg2.186$, **reject** $H_0:\beta_1=\dots=\beta_{12}=0$.

**Interpretation:** at least one covariate is significantly associated with the hourly count of rental bikes — the model as a whole has explanatory power.

> **LMES, Ex 3(e) [3 Points]** is the same test in the other notation: $F=\frac{R^2}{1-R^2}\cdot\frac{n-p-1}{p}=\frac{0.7412}{0.2588}\cdot\frac{42}{7}=17.186$, df $(7,42)$, $F_{7,42,0.95}=2.2371$ ⟹ reject.

> 🔴 **"At least one," never "all."** Both official keys use exactly that phrase.

## Q18 🔴 — Confidence-interval logic

| Paper | Statement | Verdict |
|---|---|---|
| **S25 1(k)** | *"If a CI for $\beta_j$ does not contain zero, we **fail to reject** $H_0:\beta_j=0$."* | ❌ **FALSE** — zero outside ⟹ we **do** reject |
| **W23 1b(iii)** | *"When the CI contains zero, we cannot reject the hypothesis that the coefficient is zero."* | ✅ **TRUE** |

> **Memory hook:** *zero inside the net ⟹ zero still possible ⟹ don't reject.*

## Q19 — Turn a hypothesis into a restricted model

> **S25, Ex 4(d) [2 Points].** *"For $y_i=\beta_0+\beta_1x_{1,i}+\beta_2x_{2,i}+\varepsilon_i$ you want to test $H_0:\beta_1=\beta_2+1$ via an F-test… Incorporate the null hypothesis into the model equation to obtain a restricted model that you can estimate by OLS directly."*

### Solution

Substitute $\beta_1=\beta_2+1$:

$$y_i=\beta_0+(\beta_2+1)x_{1,i}+\beta_2x_{2,i}+\varepsilon_i=\beta_0+x_{1,i}+\beta_2(x_{1,i}+x_{2,i})+\varepsilon_i$$

Move the parameter-free term to the left:

$$\boxed{\ \underbrace{(y_i-x_{1,i})}_{\text{new response}}=\beta_0+\beta_2\underbrace{(x_{1,i}+x_{2,i})}_{\text{new covariate}}+\varepsilon_i\ }$$

**Regress $(y-x_1)$ on $(x_1+x_2)$** with an intercept; its SSE is $\text{SSE}_{H_0}$. Here $r=1$.

> 🔑 **The recipe: substitute → collect → move anything without a parameter to the left as an offset.** Then count the parameters you lost: that's $r$.

---

# 3.4 — MODEL CHOICE AND DIAGNOSTICS

## Q20 🔴 — Compute AIC and BIC

> **LMES, Ex 3(f) [4 Points].** *"Calculate the AIC and BIC of the model. In general, how do you choose a model using AIC and/or BIC? What is the difference between the two?"* — $\text{RSS}=22.84961$, $n=50$, 7 covariates.

### Solution

$$\hat\sigma^2_{ML}=\frac{\text{RSS}}{n}=\frac{22.84961}{50}=0.45699,\qquad n\log\hat\sigma^2_{ML}=50\ln(0.45699)=-39.157$$

$$\text{AIC}=-39.157+2(7+1)=\boxed{-23.157}\qquad \text{BIC}=-39.157+\ln(50)(7+1)=-39.157+31.296=\boxed{-7.861}$$

*(the official key gives −23.154 and −7.858 — rounding only)*

**How to choose:** smaller is better, for both. **They may disagree**, and the key is explicit that you must not simply say "choose AIC because AIC is smaller" — the two numbers are **not comparable with each other**, only across models within the same criterion and the same data.

**The difference:** *"the BIC penalizes complex models much more than the AIC. Thus, the resulting 'best' models are typically more parsimonious when using the BIC rather than the AIC."*

> 🔴 **Three traps in one formula.** The key spells the first one out: *"in the regression book page 148, the ML estimator of the variance $\hat\sigma^2=\text{RSS}/n$ is considered in AIC and BIC, and **not** the usual unbiased variance estimator."*
> ① divide by $\boldsymbol{n}$, not $n-p$ · ② $\log=\ln$ · ③ the penalty is $(|M|+1)$ — the $+1$ is $\sigma^2$.

## Q21 — Reading AIC/BIC you're handed

> **W22, Ex 2(g) [1 Point].** *"Suppose AIC = 215827.9 and BIC = 215936.6 for the model in this exercise. What does this tell you about the model choice?"*

### Solution

**On their own: nothing.** Information criteria are only meaningful **relative to another model fitted to the same data** — there is no absolute scale, and the values here are dominated by $n\log\hat\sigma^2$, which is common to all candidate models. To choose, you need the corresponding values for a competing specification and pick the smaller.

That BIC exceeds AIC here carries no information about fit either: with $n=17379$, $\log(n)=9.76>2$, so BIC's penalty is necessarily larger for the same model.

> **W23, Ex 4(c) [2 Points]** is the version with two models: $\text{AIC}_1=2363.324$ vs $\text{AIC}_2=2274.354$ ⟹ **model 2**, because the reduction in RSS from the added quadratic term **outweighs** the penalty it costs.

## Q22 🔴 — $R^2$, adjusted $R^2$ and monotonicity

> **LMES, Ex 3(h) [1 Point].** *"How do you interpret the $R^2$ term? Why is there a difference between $R^2$ and adjusted $R^2$?"*

### Solution

$R^2$ is the proportion of the variation in the response explained by the model — here *"about 74.12% variation of life expectancy is explained by the change of the covariates."*

**$R^2$ always increases when a covariate is added**, so it cannot compare models of different sizes. Adjusted $R^2$ subtracts a penalty for the number of covariates; it rises **only** if the new variable earns its degree of freedom, and — unlike $R^2$ — it **can be negative** when the fit is poor.

### The four T/F variants of this one fact

| Paper | Statement | Verdict |
|---|---|---|
| **S25 1(c)** | *"adding dummies for the weekday a person was born on can be expected to **lower** the $R^2$"* | ❌ **FALSE** |
| **W23 1c(i)** | *"RSS **may increase** as more variables are added"* | ❌ **FALSE** |
| **LMES 1b(iii)** | *"The coefficient of determination **can decrease** as more variables are added"* | ❌ **FALSE** |
| **LMES 1c(iv)** | *"Adjusted $R^2$ … **can never be negative**"* | ❌ **FALSE** |

> **Why:** adding a column expands the column space, so the projection of $\boldsymbol{y}$ can only move closer. *(You can't un-place a jigsaw piece.)*
> 🔴 **Four papers, one fact.** And the fourth is the exception that proves it — $\bar R^2$, AIC and BIC **can** worsen; that is the entire point of §3.4.

## Q23 🔴 — AIC and BIC as T/F

| Paper | Statement | Verdict |
|---|---|---|
| **LMES 1c(iii)** | *"AIC and BIC both penalize the number of parameters, but **AIC penalizes more heavily** than BIC"* | ❌ **FALSE** — $\log(n)>2$ for $n>7.4$, so **BIC** penalises more. *B for Bigger.* |
| **W23 1c(ii)** | *"When a predictor unrelated to the response is added, the AIC **must decrease**"* | ❌ **FALSE** — it typically **rises**; and "must" is too strong either way |

## Q24 — What are information criteria *for*

> **S25, Ex 4(c) [2 Points].** *"Explain what the intuition behind information criteria is. Explain why their design makes them a good tool for model selection."*

### Solution

An information criterion balances **goodness of fit** against **model complexity** in a single number:

$$\text{IC}=\underbrace{n\log(\hat\sigma^2_{ML})}_{\text{fit — falls with more covariates}}+\underbrace{c\cdot(|M|+1)}_{\text{penalty — rises with more covariates}}$$

Pure fit measures like $R^2$ or the RSS improve **monotonically** with model size, so optimising them always selects the largest model and **overfits**: the extra parameters chase noise that will not recur in new data. The penalty term makes an added covariate pay for itself — it is kept only if it reduces $n\log\hat\sigma^2$ by more than it costs.

This is the **bias–variance trade-off** made operational: more complexity means less bias but more variance, and the criterion locates the minimum of the U-shaped total error. **AIC** ($c=2$) targets predictive performance; **BIC** ($c=\log n$) penalises harder and targets identifying the true model, yielding more parsimonious selections.

## Q25 — How would you decide whether to add terms?

> **S25, Ex 2(d) [2 Points].** *"A colleague suggests adding many more functions of age, e.g. $(\text{age}-48)^3$, $\text{age}^4$ or $\log(\text{age})$ because it will improve the fit. How would you decide whether you should do so? **State two methods.**"*

### Solution — name any two

1. **Information criteria** — compute AIC and/or BIC with and without the extra terms and keep the specification with the smaller value. Fit alone will always improve, which is exactly why a penalised criterion is needed.
2. **An F-test on the added terms** — the general linear hypothesis $H_0:$ all new coefficients $=0$, with $r$ = number of terms added.
3. Cross-validation / predictive MSE on held-out data.
4. Adjusted $R^2$, or Mallow's $C_p$.
5. Inspect the diagnostic plots for whether the residual pattern actually improves.

> **W23, Ex 4(d) [1 Point]** asks the same thing and the key accepts: *"the F-test, test for significance of the polynomial term, Mallows CP value, predicted MSE, or simply observe the diagnostic plot."*

## Q26 — Read the four diagnostic plots

> **W23, Ex 4(a) [3 Points].** `mpg ~ horsepower`, diagnostic plots supplied. *"Briefly discuss three problems that the model may have."*
> **Official key: 1 point each, minus 0.5 if named without explanation.**

### Solution

1. **Non-linearity** — residuals vs fitted show a clear **curved (U-shaped)** pattern rather than a formless band, so the straight-line specification is wrong (A1 violated).
2. **Heteroscedasticity** — the spread of the residuals **widens** with the fitted values (funnel), so $\text{Var}(\varepsilon_i)$ is not constant (A3 violated).
3. **Non-normality** — in the QQ plot the points **depart from the 45° diagonal** in the tails, indicating skew/heavy tails (A6 questionable).

> 🔴 **Name it *and* say what you saw.** "Heteroscedasticity" alone costs half a mark; "heteroscedasticity — the residual spread widens with fitted values" gets the full point.

> **W23, Ex 4(b) [2 Points].** *Does adding $\text{horsepower}^2$ fix it?* — It addresses **the non-linearity only**. A quadratic changes the shape of the mean function; it does nothing about non-constant error variance or non-normal errors, which would need a transformation of $y$, weighted least squares, or robust standard errors.

## Q27 🔴 — Standardised residuals

> **W22, Ex 1c(i).** *"Since the residuals are normally distributed, standardized residuals are also normally distributed."*

### **FALSE** (and explain why).

Even under perfect assumptions, $\text{Var}(\hat\varepsilon_i)=\sigma^2(1-h_{ii})$ is **not constant** across $i$, and $\text{Cov}(\hat\varepsilon_i,\hat\varepsilon_j)=-\sigma^2h_{ij}\neq0$. The whole *reason* for standardising by $\hat\sigma\sqrt{1-h_{ii}}$ is that raw residuals are heteroscedastic and correlated. And even standardised residuals are not exactly normal, since $\hat\sigma$ is estimated from the same data.

## Q28 🔴 — The QQ-plot sentence that breaks at the end

> **W22, Ex 1a(iv).** *"In Q-Q plots, the empirical quantiles are compared to the quantiles of the theoretical distribution. If the data follows the distribution, the points should closely follow a **horizontal line through the origin**."*

### **FALSE.**

First clause correct, second wrong: the points should follow the **45° diagonal** $y=x$. A horizontal line would mean every empirical quantile is identical — no variation at all.

> 🛟 **The archetypal construction in this course.** Read to the final word.

## Q29 🔴 — Heteroscedasticity: bias or efficiency?

> **S25, Ex 4(e) [1 Point].** *"The variation in revenue grows as the number of employees grows… Which impact does this phenomenon have on the OLS estimate in terms of bias and efficiency?"*

### Solution

The phenomenon is **heteroscedasticity** (A3 violated).

> **Bias:** none. OLS remains **unbiased** and consistent — unbiasedness needs only correct specification, zero-mean errors and full rank.
>
> **Efficiency:** lost. Gauss–Markov no longer applies, so $\hat{\boldsymbol\beta}$ is **not BLUE**; weighted/generalised least squares would have smaller variance.
>
> **And:** the usual standard errors are biased, so t-tests, F-tests and CIs are **invalid** — even though the point estimates are fine.

> **The question names the two words it wants** — "bias and efficiency" — so answer in exactly those terms. **W23 1b(ii)**, *"correlated residuals do not affect consistency but may affect efficiency"*, is ✅ **TRUE** by the same logic. **LMES 1b(iv)**, *"traditional standard errors remain valid under heteroskedasticity"*, is ❌ **FALSE**.

## Q30 🔴 — Multicollinearity, four ways

| Paper | Statement | Verdict |
|---|---|---|
| **S25 1(j)** | *"Multicollinearity can inflate the variance of the OLS coefficient estimators."* | ✅ **TRUE** |
| **LMES 1b(i)** | *"…can inflate the variance but does not bias the estimates themselves."* | ✅ **TRUE** |
| **W23 1b(i)** | *"Highly correlated explanatory variables may lead to a **reduction** in the standard errors."* | ❌ **FALSE** — inflation, not reduction |
| **LMES 1b(ii)** | *"If VIF for all variables is close to 1, multicollinearity is **likely a concern**."* | ❌ **FALSE** — VIF ≈ 1 means **no** collinearity |

> **Keep perfect and near separate:** perfect collinearity breaks A5 (not identified at all); near collinearity leaves $\hat{\boldsymbol\beta}$ **unbiased and still BLUE**, merely imprecise.

## Q31 — Leverage, outliers, influence

| Paper | Statement | Verdict |
|---|---|---|
| **LMES 1c(i)** | *"Cook's distance helps identify points that might be unduly influencing the model's fit."* | ✅ **TRUE** |
| **LMES 1c(ii)** | *"The hat matrix provides the leverages which help identify potential outliers."* | ✅ **TRUE** *(per the official key — though strictly leverage flags unusual $\boldsymbol{x}$, not unusual $y$)* |
| **LMES 1a(ii)** | *"Residual plots **cannot** be used to identify non-linearity."* | ❌ **FALSE** — that is precisely what residuals-vs-fitted does |

$$h_{ii}\ (\text{unusual }\boldsymbol{x})\qquad |r_i|\ (\text{unusual }y)\qquad D_i=\frac{r_i^2}{p}\cdot\frac{h_{ii}}{1-h_{ii}}\ (\text{both})$$

**High leverage alone is harmless** — it only hurts alongside a large residual.

## Q32 — Residual mean

> **W23, Ex 1b(iv).** *"In a well-fitted linear regression model, the mean of the residuals should be close to zero."* → ✅ **TRUE** *(official key)*

But understand **why it is a weak statement**: with an intercept, $\sum\hat\varepsilon_i=0$ holds **by construction**, for a terrible model as much as a good one. Diagnostics therefore read the **pattern** of the residuals, never their mean.

## Q33 — Miscellaneous

> **W22, Ex 1c(iii).** *"Interaction can only be computed between two continuous or between a continuous and a categorical independent variable."* → ❌ **FALSE** — two **categorical** variables can interact too (the product of two dummies), which is exactly the S25 wage-model extension.
>
> **W22, Ex 1a(iii).** *"For $H_0:\beta_j=0$, under $H_0$, $t_j^2\sim F_{1,\,n-p}$."* → ✅ **TRUE** — the $r=1$ identity $\sqrt{F}=|t|$, squared. Use it as a sanity check whenever $r=1$.

---

# THE FIVE-PAPER PATTERN

| Question type | S25 | LMES | W23 | W22 | EX20 | Marks when it appears |
|---|---|---|---|---|---|---|
| Fill missing R output | ✅ | ✅ | | | ✅ | 2.5–12 |
| Derive OLS | ✅ | ✅ | ✅ | | | 2–3 |
| Gauss–Markov / BLUE | ✅ | | ✅ | | | 4 |
| Build $\boldsymbol{C},\boldsymbol{d}$, count $r$ | ✅ | | | ✅ | | 1–2 |
| F-test (overall or partial) | ✅ | ✅ | | ✅ | | 1.5–8 |
| Confidence interval | ✅ | ✅ | ✅ | ✅ | | 1–3.5 |
| t-test written out in full | | ✅ | ✅ | | | 2 |
| AIC / BIC | | ✅ | ✅ | ✅ | | 1–4 |
| Diagnostics / residual plots | | ✅ | ✅ | ✅ | | 1–3 |
| Heteroscedasticity consequences | ✅ | ✅ | ✅ | | | 1 |
| Dummies / interactions | ✅ | | | ✅ | ✅ | 3–25 |
| Logit | ✅ | | | | | 1–1.5 |

**Every single one of these appears in at least two papers.** There is no question type in this course that you cannot have seen before walking in.

> **Next:** drill the T/F items in `../99-exam-vault/30-TRUE-FALSE-DRILL.md`, then sit a full paper closed-book at 60 minutes.
