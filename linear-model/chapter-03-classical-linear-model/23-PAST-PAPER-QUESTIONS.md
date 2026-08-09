# Ch 3 — PAST-PAPER QUESTIONS (all five papers)

*বাংলা সংস্করণ নিচে আছে → [বাংলায় পড়ো](#অধ্যায়-৩--বিগত-বছরের-প্রশ্ন-বাংলা)*

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

---
---

# অধ্যায় ৩ — বিগত বছরের প্রশ্ন (বাংলা)

> টেকনিক্যাল শব্দ, ফাইলের নাম, সূত্র আর পরীক্ষার হুবহু উদ্ধৃতি ইংরেজিতেই রেখেছি — **পরীক্ষার উত্তর ইংরেজিতে লিখতে হবে**।
>
> **S25** = Exam Summer 2025 · **LMES** = Linear_model_exam_sheet · **W23** = WiSe 2023/24 · **W22** = RCLM WS 22/23 · **EX20** = Example Exam LiMo 2020
>
> ⚠️ LMES আর W23-এ $p$ = **covariate** (তাই $\text{df}=n-p-1$); W22 আর বইয়ে $p$ = **parameter** (তাই $\text{df}=n-p$)। একই সংখ্যা। **Intercept সহ beta গোনো, আর $n$ থেকে বাদ দাও।**

**এটাই সেই ফাইল। প্রতিটা পেপারের প্রায় ৮৫% এখানেই থাকে।** নিচের সেকশনগুলো বইয়ের ক্রম অনুসরণ করে; প্রতিটা প্রশ্নে সোর্স পেপার লেখা আছে যাতে তুমি PDF-এর সাথে মিলিয়ে নিতে পারো।

---

# ৩.১ — মডেলের সংজ্ঞা আর assumption

## প্রশ্ন ১ — $\boldsymbol{X}'\boldsymbol{X}$ কখন invertible?

> **LMES, Ex 2(b) [২ নম্বর].** *"Write down the matrix $\boldsymbol{X}'\boldsymbol{X}$. Under which conditions is the matrix $\boldsymbol{X}'\boldsymbol{X}$ invertible?"* (simple linear regression)

### সমাধান

$$\boldsymbol{X}'\boldsymbol{X}=\begin{pmatrix}1&1&\cdots&1\\x_1&x_2&\cdots&x_n\end{pmatrix}\begin{pmatrix}1&x_1\\1&x_2\\\vdots&\vdots\\1&x_n\end{pmatrix}=\begin{pmatrix}n&\sum_{i}x_i\\\sum_i x_i&\sum_i x_i^2\end{pmatrix}\quad\text{[১ নম্বর]}$$

**শর্ত:** $\boldsymbol{X}$-এর **full column rank** থাকতে হবে, $\text{rank}(\boldsymbol{X})=p$ — সমতুল্যভাবে $\boldsymbol{X}'\boldsymbol{X}$ non-singular হতে হবে। **[১ নম্বর]**

Simple ক্ষেত্রে এটা ঠিক তখনই ভাঙে যখন সব $x_i$ একই (তখন $\det=n\sum x_i^2-(\sum x_i)^2=n\sum(x_i-\bar x)^2=0$): $x$-এ কোনো ভিন্নতা না থাকলে **চিহ্নিত করার মতো কোনো slope-ই নেই**।

## প্রশ্ন ২ 🔴 — Rank নিয়ে TRUE/FALSE, চারবার

| পেপার | বক্তব্য | রায় | কেন |
|---|---|---|---|
| **LMES 1a(i)** | *"$\text{rank}(\boldsymbol{X}'\boldsymbol{X})=p$, with the number of **variables** $p$"* | ❌ **FALSE** | ওই পেপারে $p$ = covariate, তাই rank হবে $p+1$ — **বাক্যের ভেতরের সংজ্ঞাটাই** এটাকে ভুল বানিয়েছে |
| **W22 1a(i)** | *"$\text{rk}(\boldsymbol{X}'\boldsymbol{X})=p$"* (খালি) | ✅ **TRUE** | বইয়ের নোটেশন, $p$ = parameter |
| **W22 1b(ii)** | *"$\text{rk}(\boldsymbol{X}'\boldsymbol{X})=k$"* | ❌ **FALSE** | একটা কম — intercept ভুলে গেছে |
| **W22 1c(iv)** | *"$\text{rk}(\boldsymbol{X}'\boldsymbol{X})=n$"* | ❌ **FALSE** | $\boldsymbol{X}'\boldsymbol{X}$ হলো $p\times p$; rank কখনো ছোট মাত্রাটা ছাড়াতে পারে না |

> 🔴 **দেখতে একরকম দুটো বক্তব্য, উল্টো উত্তর — কারণ একটায় বাক্যের ভেতরেই $p$-এর সংজ্ঞা দেওয়া।** উত্তর দেওয়ার আগে পেপারের নিজের সংজ্ঞাটা পড়ো।

## প্রশ্ন ৩ 🔴 — Rank ঘাটতির সাথে অপ্রাসঙ্গিক শর্ত জুড়ে দেওয়া

> **S25, Ex 1(d).** *"When the design matrix $\boldsymbol{X}$ does not have full column rank, the OLS estimates still exist and are unique as long as the error variance is constant."*

### **FALSE.**

Full rank না থাকলে $\boldsymbol{X}'\boldsymbol{X}$ singular, আর normal equation $\boldsymbol{X}'\boldsymbol{X}\hat{\boldsymbol\beta}=\boldsymbol{X}'\boldsymbol{y}$-এর **অসীম সংখ্যক** সমাধান থাকে। Homoscedasticity **efficiency** নিয়ে, identification-এর সাথে এর কোনোই সম্পর্ক নেই।

> 🛟 **"as long as…" গঠনটা।** প্রতিটা এমন শর্ত সম্পর্কে জিজ্ঞেস করো: *এই শর্তের সাথে দাবিটার আদৌ কোনো সম্পর্ক আছে?* এখানে নেই — আর ওটাই পুরো ফাঁদ।

## প্রশ্ন ৪ — Intercept থাকলে যা আপনা-আপনি সত্য

> **W22, Ex 1a(ii).** *"The average of the predicted values is equal to the average of the observed response."*

### **TRUE** — যখন মডেলে intercept আছে।

Intercept-এর normal equation থেকে $\sum_i\hat\varepsilon_i=0$ আসে, ফলে $\bar{\hat y}=\bar y$, আর fitted রেখা $(\bar x,\bar y)$ দিয়ে যায়।

> ⚠️ **গঠনগতভাবে সত্য, ভালো fit-এর কারণে নয়।** নিচের প্রশ্ন ৩২-এ এর উল্টো ফাঁদটা দেখো।

---

# ৩.২ — Estimation

## প্রশ্ন ৫ 🔴🔴 — OLS estimator derive করা

**পাঁচটার মধ্যে তিনটা পেপারে এসেছে।** একবার শেখো; প্রতি বছর তুলে নাও।

> **S25, Ex 4(b) [২ নম্বর].** *"Explain the method of ordinary least squares (OLS) estimation and show the steps necessary to obtain the estimators $\hat\beta_0,\dots,\hat\beta_k$. It is not necessary to explicitly calculate them; it suffices to show the mathematical approach."*
>
> **LMES, Ex 2(c) [৩ নম্বর].** *"Derive the least square estimators in matrix form."*
>
> **W23, Ex 2(b) [২ নম্বর].** *"Explain how the ordinary least squares (OLS) method is used to estimate the coefficients $\beta_0$ and $\beta_1$."*

### সমাধান

$$\boldsymbol\varepsilon=\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta,\qquad \text{RSS}=\boldsymbol\varepsilon'\boldsymbol\varepsilon=(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)=\boldsymbol{y}'\boldsymbol{y}-2\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{y}+\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{X}\boldsymbol\beta\quad\text{[১ নম্বর]}$$

$$\frac{\partial\,\text{RSS}}{\partial\boldsymbol\beta}=-2\boldsymbol{X}'\boldsymbol{y}+2\boldsymbol{X}'\boldsymbol{X}\boldsymbol\beta\overset{!}{=}\boldsymbol{0}\quad\text{[১ নম্বর]}$$

$$\boldsymbol{X}'\boldsymbol{X}\hat{\boldsymbol\beta}=\boldsymbol{X}'\boldsymbol{y}\ \Longrightarrow\ \boxed{\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}}\quad\text{[১ নম্বর]}$$

Inverse নিতে $\text{rank}(\boldsymbol{X})=p$ লাগে। Hessian $2\boldsymbol{X}'\boldsymbol{X}$ positive definite, তাই এটা **minimum**।

> **W23-এর marking key, হুবহু:** *"1 point for correctly state that RSS needs to be minimized. And 1 point for correctly derive the solution."*
>
> 🔑 **অন্তরকলন করার আগে বলো তুমি কী minimise করছ।** অর্ধেক নম্বর শুধু objective-টার নাম বলার জন্য — যারা সরাসরি বীজগণিতে ঝাঁপ দেয়, তারা জানা জিনিসের নম্বরই হারায়।

## প্রশ্ন ৬ — হাতে $\hat{\boldsymbol\beta}$, RSS আর $\hat\sigma^2$ বের করা

> **W23, Ex 3(a)+(b) [৩+২ নম্বর].** $X=(20,30,33,40,15)$, $Y=(7,9,8,11,5)$, $n=5$, দেওয়া $(\boldsymbol{X}'\boldsymbol{X})^{-1}=\begin{pmatrix}2.079&-0.068\\-0.068&0.0024\end{pmatrix}$।

### সমাধান

$$\boldsymbol{X}'\boldsymbol{X}=\begin{pmatrix}5&138\\138&4214\end{pmatrix},\quad \boldsymbol{X}'\boldsymbol{Y}=\begin{pmatrix}40\\1189\end{pmatrix},\quad \hat{\boldsymbol\beta}=\begin{pmatrix}2.079&-0.068\\-0.068&0.0024\end{pmatrix}\begin{pmatrix}40\\1189\end{pmatrix}=\begin{pmatrix}2.210\\0.209\end{pmatrix}$$

$$\hat{\boldsymbol{y}}=\boldsymbol{X}\hat{\boldsymbol\beta}=(6.40,\ 8.50,\ 9.13,\ 10.60,\ 5.35)'$$

$$\hat{\boldsymbol\varepsilon}=\boldsymbol{y}-\hat{\boldsymbol{y}}=(0.594,\ 0.496,\ -1.132,\ 0.398,\ -0.356)'$$

$$\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}=2.169,\qquad \hat\sigma^2=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n-2}=\frac{2.169}{3}=\boxed{0.723}$$

> **হর $n-2$**, কারণ দুটো parameter estimate করা হয়েছে। বইয়ের নোটেশনে এটাই $n-p$।

## প্রশ্ন ৭ 🔴 — Gauss–Markov, ৪ নম্বরের জন্য

> **W23, Ex 2(a) [৪ নম্বর].** *"Briefly describe the main contents of the Gauss–Markov Theorem and the assumptions."*
> **অফিসিয়াল key: *"1 point for every assumption."***

### সমাধান — তালিকা করে লেখো, অনুচ্ছেদ নয়

> এই assumption-গুলোর অধীনে —
> **(i)** মডেল সঠিকভাবে specified, $E(\boldsymbol{y})=\boldsymbol{X}\boldsymbol\beta$;
> **(ii)** $E(\boldsymbol\varepsilon)=\boldsymbol{0}$;
> **(iii)** homoscedasticity, সব $i$-এর জন্য $\text{Var}(\varepsilon_i)=\sigma^2$;
> **(iv)** uncorrelated error, $i\neq j$-এর জন্য $\text{Cov}(\varepsilon_i,\varepsilon_j)=0$ — (iii) আর (iv) একসাথে মানে $\text{Cov}(\boldsymbol\varepsilon)=\sigma^2\boldsymbol{I}_n$;
> **(v)** $\text{rank}(\boldsymbol{X})=p$,
>
> OLS estimator $\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$ হলো **BLUE** — $\boldsymbol{y}$-তে **linear** এবং $\boldsymbol\beta$-এর জন্য **unbiased** সব estimator-এর মধ্যে এর variance সবচেয়ে কম।
>
> **Error-এর normality *লাগে না*।**

> 🔴 Key **প্রতি assumption-এ নম্বর** দেয়, তাই এক এক করে লেখো। আর normality-র বাক্যটা সত্যিকারের নম্বর, কারণ দুটো আলাদা T/F ঠিক ওটাই পরীক্ষা করে (নিচে প্রশ্ন ৮)।

## প্রশ্ন ৮ 🔴 — BLUE, তিনভাবে

| পেপার | বক্তব্য | রায় |
|---|---|---|
| **S25 1(e)** | *"A BLUE is 'best' in the sense that there is no other **linear unbiased** estimator with a lower variance."* | ✅ **TRUE** — সঠিকভাবে শর্তযুক্ত |
| **W23 1a(ii)** | *"The LS estimator is BLUE **if and only if** the error term is expected to be zero and has constant variance."* | ❌ **FALSE** — "iff" বড্ড কড়া, আর তালিকায় uncorrelated error, correct specification ও full rank নেই |
| **S25 1(l)** | *"The OLS estimator is equivalent to the ML estimator under i.i.d. normal errors."* | ✅ **TRUE** — $\boldsymbol\beta$-এর জন্য। ⚠️ কিন্তু $\hat\sigma^2_{ML}\neq\hat\sigma^2_{LS}$ |

> 🔑 BLUE-র বক্তব্য থেকে **"linear"** বা **"unbiased"** বাদ দিলেই সেটা মিথ্যা হয়ে যায় — ridge biased অথচ কম variance-এর।

## প্রশ্ন ৯ 🔴 — অপ্রয়োজনীয় ভেরিয়েবল যোগ করলে unbiasedness

> **S25, Ex 1(a).** *"Adding a variable which is not correlated with the dependent variable will not affect the unbiasedness of the OLS estimator, but it may affect its variance."*

### **TRUE.**

অপ্রাসঙ্গিক একটা regressor রাখলে $\hat{\boldsymbol\beta}$ unbiased-ই থাকে (মডেল তখনও সঠিকভাবে specified — ওর সত্যিকারের coefficient কেবল শূন্য), কিন্তু একটা degree of freedom খরচ হয়, আর নতুন ভেরিয়েবলটা আগের কোনোটার সাথে correlated হলে variance **বেড়ে** যায়।

**বাদ দেওয়ার সাথে তুলনা করো:** একটা *প্রাসঙ্গিক* covariate বাদ দিলে, যেটা অন্তর্ভুক্তগুলোর সাথে correlated, তখন $\hat{\boldsymbol\beta}$ **biased হয়**। অপ্রয়োজনীয় জিনিস ঢোকালে **নিখুঁততা** কমে; দরকারি জিনিস বাদ দিলে **সঠিকতা** যায়।

## প্রশ্ন ১০ — Omitted-variable bias, সংখ্যা দিয়ে

> **EX20, Ex 1(c)(d)(e) [৪+৬+৪+৩ নম্বর].** মডেল: `points ~ goals`, $\hat\beta_{\text{goals}}=0.90509$। দুটো সহায়ক মডেল দেওয়া: `goals.received ~ goals` দেয় $-0.5850$ (significant), আর `points ~ goals.received` দেয় $-0.9096$ (significant)। গড়: $\overline{\text{goals}}=48.61$, $\overline{\text{points}}=46.61$।

### (c) তিনি কী নিয়ে চিন্তিত? [৪ নম্বর]

**$\hat{\boldsymbol\beta}$-এর bias** — অর্থাৎ estimator-এর প্রত্যাশিত মান সত্যিকারের parameter নয়। Unbiased মানে $E(\hat{\boldsymbol\beta})=\boldsymbol\beta$; biased মানে $E(\hat{\boldsymbol\beta})\neq\boldsymbol\beta$।

### (d) তিনি কি ঠিক? [৬ নম্বর]

**হ্যাঁ।** `goals.received` (i) `points`-কে প্রভাবিত করে বলে ধরা যায় — দ্বিতীয় output-এ coefficient significant — এবং (ii) **`goals`-এর সাথে correlated**, কারণ প্রথম output দেখাচ্ছে `goals` উল্লেখযোগ্যভাবে `goals.received` predict করে। **দুটো শর্তই লাগে।** যে omitted variable কোনো অন্তর্ভুক্ত covariate-এর সাথেই correlated নয়, সেটা estimate-কে unbiased রাখে; এখানে দুটোই সত্য, তাই মূল মডেলের estimate biased।

### (e) Bias কত? আর full মডেলের $\hat\beta_0$ বের করো। [৪+৩ নম্বর]

$$\text{bias}=\hat\beta_{\text{goals}\to\text{goals.rec}}\times\beta_{\text{goals.rec}}=(-0.585)\times(-0.45763)=\boxed{0.268}$$

$$\beta_{\text{goals}}=0.90509-0.268=\boxed{0.637}$$

দুটো গড়ই ৪৮.৬১ (প্রতিটা গোল একটা দল দেয়, আরেকটা দল খায়), তাই:

$$\hat\beta_0=46.61-\left[(0.637-0.45763)\times48.61\right]=46.61-8.719=\boxed{37.891}$$

> 🔑 **Bias হলো দুটো পথের গুণফল:** omitted variable-টা অন্তর্ভুক্তটার সাথে কীভাবে নড়ে, গুণ omitted variable-এর নিজের সত্যিকারের প্রভাব। Bias-এর চিহ্ন = দুটো চিহ্নের গুণফল — এখানে ঋণাত্মক × ঋণাত্মক = **ধনাত্মক**, তাই সরল estimate-টা **বেশি বড় ছিল**।

---

# ৩.৩ — Testing আর interval

## প্রশ্ন ১১ 🔴🔴 — R output-এর ফাঁকা ঘর পূরণ

**সবচেয়ে বেশিবার পুনরাবৃত্ত হিসাবের প্রশ্ন। S25, LMES আর EX20 — তিনটাতেই আছে।**

> **S25, Ex 3(a) [২.৫ নম্বর].** `[[A]]`–`[[D]]` বের করো আর fitted regression formula লেখো।
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
> সাথে $\sum\hat\varepsilon_i^2=31682.02$।

### সমাধান

সবকিছু আসে **একটাই সম্পর্ক** থেকে, $t=\hat\beta/\widehat{\text{se}}$, তিনভাবে সাজিয়ে:

$$\text{[[A]]}=\frac{\hat\beta}{t}=\frac{48.38458}{13.591}=\boxed{3.560}$$

$$\text{[[B]]}=\frac{\hat\beta}{\widehat{\text{se}}}=\frac{-36.99122}{5.25574}=\boxed{-7.038}$$

$$\text{[[C]]}=t\times\widehat{\text{se}}=-3.796\times0.26423=\boxed{-1.003}$$

$$\text{[[D]]}=\hat\sigma=\sqrt{\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n-p}}=\sqrt{\frac{31682.02}{501}}=\sqrt{63.238}=\boxed{7.952}$$

$$\widehat{\text{medv}}=48.385-0.260\,\text{crim}-36.991\,\text{nox}-1.003\,\text{dis}-0.062\,\text{rad}$$

> 🔴 **[[D]] হলো standard *error*** — বর্গমূল নাও। ভুলে গেলে ৭.৯৫২-এর বদলে ৬৩.২৩৮ আসবে।
> 🔴 এখানে $n=506$: ৫০১ df $+\ p=5$ parameter। **df-এর লাইনটা সবসময় $n$ ফিরিয়ে দেয়।**

### একই দক্ষতা, আরও দুইবার

> **LMES, Ex 3(b) [৪ নম্বর].** $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}=22.84961$, $n=50$, ৭টা covariate, $\bar R^2=0.6981$।
>
> $$\text{(A)}\ t=\frac{-0.298}{0.045}=-6.587\qquad \text{(B)}\ \hat\sigma=\sqrt{\frac{22.849}{42}}=\sqrt{0.544}=0.738$$
> $$\text{(C)}\ R^2=1-\frac{n-p-1}{n-1}(1-\bar R^2)=1-\frac{42}{49}(0.3019)=0.741$$
> $$\text{(D)}\ F=\frac{R^2}{1-R^2}\cdot\frac{n-p-1}{p}=\frac{0.7412}{0.2588}\cdot\frac{42}{7}=17.186$$

> **EX20, Ex 1(a) [৪+৪+৪ নম্বর].** `points ~ goals`, $n=18$, $\overline{\text{goals}}=48.61$, $\overline{\text{points}}=46.61$, $\hat\beta_{\text{goals}}=0.90509$, $t_{\text{goals}}=9.562$, $\widehat{\text{se}}_0=4.81560$।
>
> $$\text{A}=\hat\beta_0=\bar y-\hat\beta_1\bar x=46.61-0.90509(48.61)=46.61-43.996=\boxed{2.614}$$
> $$\text{B}=\widehat{\text{se}}_1=\frac{0.90509}{9.562}=\boxed{0.095}\qquad \text{C}=t_0=\frac{2.614}{4.8156}=\boxed{0.543}$$
>
> **(b) [৬ নম্বর]** p-value $D$ হলো $2\big(1-F_{t_{16}}(9.562)\big)$ — R-এ: `(1 - pt(q = 9.562, df = 16)) * 2`। Degrees of freedom $=n-p-1=18-1-1=16$। **২ দিয়ে গুণ করা হয় কারণ test দুই-পাশের** আর $t$-distribution প্রতিসম, তাই দুই লেজই হিসাবে আসে।

> 🔑 **তিনটা পেপার, একটাই টুলবক্স:** $t=\hat\beta/\widehat{\text{se}}$ · $\hat\beta_0=\bar y-\hat\beta_1\bar x$ · $\hat\sigma^2=\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}/(n-p)$ · $\bar R^2\leftrightarrow R^2$ · $F\leftrightarrow R^2$। **এর বাইরে কোনোদিন কিছু লাগেনি।**

## প্রশ্ন ১২ — পূর্ণাঙ্গ t-test, নম্বরের জন্য লেখা

> **LMES, Ex 3(d) [২ নম্বর].** *"Conduct a significance test for `Murder`. Clearly write the null and alternative hypotheses, your test statistic, degrees of freedom, critical value, and draw the correct conclusion. $\alpha=0.05$."*

### সমাধান — প্রতিবার এই কাঠামোটাই লেখো

| ধাপ | |
|---|---|
| $H_0$ | $\beta_{\text{Murder}}=0$ |
| $H_1$ | $\beta_{\text{Murder}}\neq0$ |
| Distribution | $t$ |
| Statistic | $t=\dfrac{-0.298}{0.045}=-6.587$ |
| df | $n-p-1=50-7-1=42$ |
| Critical value | $t_{0.975,42}=2.0180$ |
| সিদ্ধান্ত | $|-6.587|>2.0180$ ⟹ **$H_0$ বাতিল** |
| উপসংহার | ৫% স্তরে `Murder` life expectancy-র সাথে উল্লেখযোগ্যভাবে যুক্ত |

> Key-তে লেখা: *"if the student is confused to use other degrees of freedom or other significance level, −0.5."* **Quantile হলো $1-\alpha/2=0.975$, $0.95$ নয়।**
>
> **W23, Ex 3(c)** একই কাঠামো $n=5$-এ: $\widehat{\text{se}}(\hat\beta_1)=\sqrt{\hat\sigma^2\left[(\boldsymbol{X}'\boldsymbol{X})^{-1}\right]_{22}}=\sqrt{0.723\times0.0024}=0.042$, $t=0.20977/0.04224=4.966$, df $=3$, $t_{0.975,3}=3.182$ ⟹ বাতিল।
>
> ⚠️ W23-এর key-তে variance ভেক্টর ছাপা আছে $(1.504, 0.084)$; আসলে $0.0024\times0.723=0.0017$, আর key-র নিজের পরের লাইনই ($\text{sd}=0.042=\sqrt{0.0017}$) প্রমাণ করে $0.084$ একটা টাইপো। **Standard error ০.০৪২-ই সঠিক।**

## প্রশ্ন ১৩ — Confidence interval, তিনটা পেপারে

> **S25, Ex 3(b) [২ নম্বর].** $\beta_{\text{nox}}$-এর ৯৯% CI; তারপর ১% স্তরে $H_0:\beta_{\text{nox}}=-30$ পরীক্ষা করো।

$$-36.99122\pm t_{501}(0.995)\times5.25574=-36.99122\pm2.5857(5.25574)=-36.991\pm13.590$$

$$\boxed{[-50.581,\ -23.401]}$$

**$-30$ interval-এর ভেতরে ⟹ ১% স্তরে $H_0:\beta_{\text{nox}}=-30$ বাতিল করা যায় না।**

> 🔴 ৯৯% ⟹ **০.৯৯৫** quantile। আর খেয়াল করো, উত্তর "fail to reject" — যদিও **শূন্যের** বিপরীতে nox প্রচণ্ড significant। **null-এর মানটাই গুরুত্বপূর্ণ, ভেরিয়েবলটা নয়।**

> **LMES, Ex 3(g) [২ নম্বর].** `HS.Grad`-এর ৯৫% CI: $0.0584\pm2.0180(0.0242)=[0.0095,\ 0.1073]$। Output-এর significance-এর সাথে সামঞ্জস্যপূর্ণ, কারণ **শূন্য এর ভেতরে নেই** ⟹ ৫%-এ $H_0:\beta=0$ বাতিল।
>
> **W23, Ex 3(d) [১ নম্বর].** $0.20977\pm3.182(0.04224)=[0.0753,\ 0.3442]$।
>
> **W22, Ex 2(f) [৩.৫ নম্বর].** $\beta_{12}$ (windspeed)-এর ৯৯% CI: $-42.244\pm2.5761(7.938)=-42.244\pm20.449=[-62.693,\ -21.795]$।

> 🔑 **CI–test duality থেকে ফ্রি নম্বর আসে:** প্রশ্ন যদি হয় "is this consistent with the regression table?", উত্তর সবসময় null-এর মানটা ভেতরে আছে কিনা তা নিয়ে।

## প্রশ্ন ১৪ 🔴🔴 — $\boldsymbol{C}$ আর $\boldsymbol{d}$ বানানো

> **S25, Ex 3(c) [২ নম্বর].** *"Test the joint null $H_0:\beta_{\text{crim}}=3\beta_{\text{rad}}-0.1,\ \beta_{\text{nox}}=-40$. Express as $\boldsymbol{C}\boldsymbol\beta=\boldsymbol{d}$ with $\boldsymbol\beta=(\beta_0,\beta_{\text{crim}},\beta_{\text{nox}},\beta_{\text{dis}},\beta_{\text{rad}})'$. What is the number of linearly independent restrictions $r$?"*

### সমাধান

**আগে সাজাও** — parameter বাঁয়ে, ধ্রুবক ডানে:

$$\beta_{\text{crim}}-3\beta_{\text{rad}}=-0.1,\qquad \beta_{\text{nox}}=-40$$

$$\boldsymbol{C}=\begin{pmatrix}0&1&0&0&-3\\0&0&1&0&0\end{pmatrix},\qquad \boldsymbol{d}=\begin{pmatrix}-0.1\\-40\end{pmatrix},\qquad \boxed{r=2}$$

> 🔴 $\boldsymbol{C}$ হলো $r\times p$ — $\beta_0$-এর কলামটা শূন্য হলেও **থাকতেই হবে**।
> 🔴 ধ্রুবক $-0.1$ যায় $\boldsymbol{d}$-তে, কখনো $\boldsymbol{C}$-তে নয়।

### উল্টো দিক

> **W22, Ex 2(d) [১ নম্বর].** $\boldsymbol{C}$-তে তিনটা সারি ১১, ১২, ১৩ অবস্থান বাছে আর $\boldsymbol{d}=\boldsymbol{0}_3$ — hypothesis লেখো।
>
> $$H_0:\beta_{\text{temp}}=\beta_{\text{hum}}=\beta_{\text{windspeed}}=0\qquad\text{বনাম}\qquad H_1:\text{অন্তত একটা}\neq0$$
>
> $\boldsymbol{C}$-কে **কলাম ধরে ধরে** পড়ো: প্রতিটা ১ চিহ্নিত করছে ওই সারি কোন parameter-কে বাঁধছে। তারপর কলামের অবস্থান R output-এর coefficient-এর ক্রমের সাথে মেলাও।

## প্রশ্ন ১৫ 🔴🔴 — Restriction গোনা

> **S25, Ex 1(i).** *"The F-statistic for testing $H_0:\beta_1=-\beta_2+\beta_3$ in a linear model with $k\geq3$ predictors plus an intercept has an F-distribution with $(3,\,n-k-1)$ degrees of freedom under $H_0$."*

### **FALSE.**

$\beta_1+\beta_2-\beta_3=0$ হিসেবে লেখো আর **সমান চিহ্ন গোনো: একটা**। তাই $r=1$ আর $F\sim F_{1,\,n-k-1}$।

ওই "৩" গুনছে **উল্লিখিত beta**, যেটা কখনোই $r$-এর মানে নয়।

> **সম্পর্কিত df ফাঁদ:** **W22 1b(iv)** *"$F\sim F_{p+1,n-p}$ for $H_0:\boldsymbol{C}\boldsymbol\beta=\boldsymbol{d}$"* → **FALSE** (লব হলো $r$)। **W22 1c(ii)** *"$F\sim F_{p-1,n-p}$ for the overall test"* → **FALSE** (লব হলো $k$, আর $p-1=k$ শুধু বইয়ের নোটেশনে — পেপারটা অসামঞ্জস্যপূর্ণ, তাই তার নিজের সংজ্ঞা পড়ো)। **W23 1a(iii)** *"full-model F test $\sim F_{k,n-k-1}$"* → ✅ **TRUE**। **S25 1(g)** *"$t_j\sim t_{n-k-1}$ where $k+1$ is the number of estimated parameters"* → ✅ **TRUE**।

## প্রশ্ন ১৬ — $\Delta$SSE থেকে F-test

> **S25, Ex 3(d) [১.৫ নম্বর].** $\hat{\boldsymbol\varepsilon}'_{H_0}\hat{\boldsymbol\varepsilon}_{H_0}=32333.15$, unrestricted $\text{SSE}=31682.02$, $r=2$, ৫০১ df।

$$F=\frac{(\text{SSE}_{H_0}-\text{SSE})/r}{\text{SSE}/(n-p)}=\frac{(32333.15-31682.02)/2}{31682.02/501}=\frac{325.565}{63.238}=\boxed{5.148}$$

$5.148>F_{2,501}(0.95)=3.0137$ ⟹ **$H_0$ বাতিল**। দুটো restriction-এর **অন্তত একটা** ভাঙে।

> **W22, Ex 2(e) [৮ নম্বর].** $\Delta\text{SSE}=26731881$, $r=3$, estimated model standard error $120.3$ তাই $\hat\sigma^2=120.3^2=14472.09$:
>
> $$F=\frac{\Delta\text{SSE}/r}{\hat\sigma^2}=\frac{26731881/3}{14472.09}=\frac{8910627}{14472.09}=615.72$$
>
> Quantile: $F_{3,\,17366}(0.99)=3.7827$। $615.72\gg3.783$ ⟹ **বাতিল**। **ব্যাখ্যা:** বাকি covariate দেওয়া থাকলে, temperature, humidity আর wind speed-এর অন্তত একটা ঘণ্টাপ্রতি বাইক ভাড়ার সাথে উল্লেখযোগ্যভাবে যুক্ত।

> 🔴 **Quantile নিয়ম:** F-test **এক-পাশের** ($\text{SSE}_{H_0}\geq\text{SSE}$ সবসময়), তাই $1-\alpha$ ব্যবহার করো — ৫%-এ $0.95$, ১%-এ $0.99$। $1-\alpha/2$ নয়।
> ✅ **যাচাই:** $F\geq0$ সবসময়; $r=1$ হলে $\sqrt F=|t|$।

## প্রশ্ন ১৭ — Regression-এর সামগ্রিক significance test

> **W22, Ex 2(c) [৭ নম্বর].** $n=17379$, ১৩টা parameter, $R^2=0.5604$, $\alpha=0.01$। Test statistic, quantile, সিদ্ধান্ত, ব্যাখ্যা।

$$F=\frac{R^2/k}{(1-R^2)/(n-p)}=\frac{0.5604/12}{0.4396/17366}=\frac{0.04670}{0.00002531}=\boxed{1844.8}\ \sim F_{12,\,17366}$$

Quantile $F_{12,17366}(0.99)=2.1858$। যেহেতু $1844.8\gg2.186$, **$H_0:\beta_1=\dots=\beta_{12}=0$ বাতিল**।

**ব্যাখ্যা:** অন্তত একটা covariate ঘণ্টাপ্রতি ভাড়া করা বাইকের সংখ্যার সাথে উল্লেখযোগ্যভাবে যুক্ত — মডেলটার সামগ্রিকভাবে ব্যাখ্যা করার ক্ষমতা আছে।

> **LMES, Ex 3(e) [৩ নম্বর]** একই test, অন্য নোটেশনে: $F=\frac{R^2}{1-R^2}\cdot\frac{n-p-1}{p}=\frac{0.7412}{0.2588}\cdot\frac{42}{7}=17.186$, df $(7,42)$, $F_{7,42,0.95}=2.2371$ ⟹ বাতিল।

> 🔴 **"অন্তত একটা", কখনো "সবগুলো" নয়।** দুটো অফিসিয়াল key-তেই ঠিক এই কথাটা আছে।

## প্রশ্ন ১৮ 🔴 — Confidence interval-এর যুক্তি

| পেপার | বক্তব্য | রায় |
|---|---|---|
| **S25 1(k)** | *"If a CI for $\beta_j$ does not contain zero, we **fail to reject** $H_0:\beta_j=0$."* | ❌ **FALSE** — শূন্য বাইরে ⟹ আমরা **বাতিল করি** |
| **W23 1b(iii)** | *"When the CI contains zero, we cannot reject the hypothesis that the coefficient is zero."* | ✅ **TRUE** |

> **মনে রাখার কৌশল:** *জালের ভেতরে শূন্য ⟹ শূন্য এখনো সম্ভব ⟹ বাতিল কোরো না।*

## প্রশ্ন ১৯ — Hypothesis-কে restricted মডেলে রূপান্তর

> **S25, Ex 4(d) [২ নম্বর].** *"For $y_i=\beta_0+\beta_1x_{1,i}+\beta_2x_{2,i}+\varepsilon_i$ you want to test $H_0:\beta_1=\beta_2+1$ via an F-test… Incorporate the null hypothesis into the model equation to obtain a restricted model that you can estimate by OLS directly."*

### সমাধান

$\beta_1=\beta_2+1$ বসাও:

$$y_i=\beta_0+(\beta_2+1)x_{1,i}+\beta_2x_{2,i}+\varepsilon_i=\beta_0+x_{1,i}+\beta_2(x_{1,i}+x_{2,i})+\varepsilon_i$$

Parameter-হীন term-টা বাঁ পাশে সরাও:

$$\boxed{\ \underbrace{(y_i-x_{1,i})}_{\text{নতুন response}}=\beta_0+\beta_2\underbrace{(x_{1,i}+x_{2,i})}_{\text{নতুন covariate}}+\varepsilon_i\ }$$

Intercept সহ **$(y-x_1)$-কে $(x_1+x_2)$-এর উপর regress করো**; এর SSE-ই $\text{SSE}_{H_0}$। এখানে $r=1$।

> 🔑 **রেসিপি: বসাও → গুছাও → parameter নেই এমন সবকিছু বাঁ পাশে offset হিসেবে সরাও।** তারপর গোনো কয়টা parameter হারালে: ওটাই $r$।

---

# ৩.৪ — Model choice আর diagnostics

## প্রশ্ন ২০ 🔴 — AIC আর BIC বের করা

> **LMES, Ex 3(f) [৪ নম্বর].** *"Calculate the AIC and BIC of the model. In general, how do you choose a model using AIC and/or BIC? What is the difference between the two?"* — $\text{RSS}=22.84961$, $n=50$, ৭টা covariate।

### সমাধান

$$\hat\sigma^2_{ML}=\frac{\text{RSS}}{n}=\frac{22.84961}{50}=0.45699,\qquad n\log\hat\sigma^2_{ML}=50\ln(0.45699)=-39.157$$

$$\text{AIC}=-39.157+2(7+1)=\boxed{-23.157}\qquad \text{BIC}=-39.157+\ln(50)(7+1)=-39.157+31.296=\boxed{-7.861}$$

*(অফিসিয়াল key দেয় −23.154 আর −7.858 — শুধু rounding-এর পার্থক্য)*

**কীভাবে বাছবে:** দুটোতেই ছোট মানে ভালো। **এরা দ্বিমত করতে পারে**, আর key স্পষ্ট করে বলে যে শুধু "AIC ছোট বলে AIC বাছলাম" বলা চলবে না — **দুটো সংখ্যা একে অপরের সাথে তুলনীয় নয়**, কেবল একই criterion-এর মধ্যে একই ডেটায় বিভিন্ন মডেলের তুলনা চলে।

**পার্থক্য:** *"the BIC penalizes complex models much more than the AIC. Thus, the resulting 'best' models are typically more parsimonious when using the BIC rather than the AIC."*

> 🔴 **এক সূত্রে তিনটা ফাঁদ।** Key প্রথমটা বানান করে বলে দেয়: *"in the regression book page 148, the ML estimator of the variance $\hat\sigma^2=\text{RSS}/n$ is considered in AIC and BIC, and **not** the usual unbiased variance estimator."*
> ① $\boldsymbol{n}$ দিয়ে ভাগ করো, $n-p$ নয় · ② $\log=\ln$ · ③ penalty হলো $(|M|+1)$ — ওই $+1$ হলো $\sigma^2$।

## প্রশ্ন ২১ — হাতে ধরিয়ে দেওয়া AIC/BIC পড়া

> **W22, Ex 2(g) [১ নম্বর].** *"Suppose AIC = 215827.9 and BIC = 215936.6 for the model in this exercise. What does this tell you about the model choice?"*

### সমাধান

**একা একা: কিছুই না।** Information criteria কেবল **একই ডেটায় fit করা আরেকটা মডেলের তুলনায়** অর্থবহ — এর কোনো পরম স্কেল নেই, আর এখানকার মানগুলো $n\log\hat\sigma^2$ দিয়েই প্রভাবিত, যা সব প্রার্থী মডেলে অভিন্ন। বাছতে হলে প্রতিদ্বন্দ্বী specification-এর মান লাগবে, তারপর ছোটটা নিতে হবে।

এখানে BIC > AIC হওয়াটাও fit সম্পর্কে কিছু বলে না: $n=17379$ হলে $\log(n)=9.76>2$, তাই একই মডেলের জন্য BIC-র penalty অনিবার্যভাবেই বড়।

> **W23, Ex 4(c) [২ নম্বর]** হলো দুই মডেলওয়ালা সংস্করণ: $\text{AIC}_1=2363.324$ বনাম $\text{AIC}_2=2274.354$ ⟹ **মডেল ২**, কারণ যোগ করা quadratic term-এর RSS কমানোটা তার penalty-র চেয়ে **বেশি**।

## প্রশ্ন ২২ 🔴 — $R^2$, adjusted $R^2$ আর একমুখিতা

> **LMES, Ex 3(h) [১ নম্বর].** *"How do you interpret the $R^2$ term? Why is there a difference between $R^2$ and adjusted $R^2$?"*

### সমাধান

$R^2$ হলো response-এর যতটুকু ভিন্নতা মডেল ব্যাখ্যা করে তার অনুপাত — এখানে *"about 74.12% variation of life expectancy is explained by the change of the covariates."*

**Covariate যোগ করলে $R^2$ সবসময় বাড়ে**, তাই এটা ভিন্ন আকারের মডেল তুলনা করতে পারে না। Adjusted $R^2$ covariate-সংখ্যার জন্য একটা penalty বাদ দেয়; নতুন ভেরিয়েবল তার degree of freedom-এর দাম উসুল করলে **তবেই** এটা বাড়ে, আর — $R^2$-এর বিপরীতে — fit খারাপ হলে এটা **ঋণাত্মকও হতে পারে**।

### এই একটাই তথ্যের চারটা T/F রূপ

| পেপার | বক্তব্য | রায় |
|---|---|---|
| **S25 1(c)** | *"adding dummies for the weekday a person was born on can be expected to **lower** the $R^2$"* | ❌ **FALSE** |
| **W23 1c(i)** | *"RSS **may increase** as more variables are added"* | ❌ **FALSE** |
| **LMES 1b(iii)** | *"The coefficient of determination **can decrease** as more variables are added"* | ❌ **FALSE** |
| **LMES 1c(iv)** | *"Adjusted $R^2$ … **can never be negative**"* | ❌ **FALSE** |

> **কেন:** একটা কলাম যোগ করলে column space বড় হয়, তাই $\boldsymbol{y}$-এর projection কেবল কাছেই আসতে পারে। *(বসানো jigsaw টুকরো তো আর খোলা যায় না।)*
> 🔴 **চারটা পেপার, একটাই তথ্য।** আর চতুর্থটা সেই ব্যতিক্রম যেটা নিয়মটা প্রমাণ করে — $\bar R^2$, AIC আর BIC **খারাপ হতে পারে**; §৩.৪-এর পুরো উদ্দেশ্যই তো এটা।

## প্রশ্ন ২৩ 🔴 — AIC আর BIC নিয়ে T/F

| পেপার | বক্তব্য | রায় |
|---|---|---|
| **LMES 1c(iii)** | *"AIC and BIC both penalize the number of parameters, but **AIC penalizes more heavily** than BIC"* | ❌ **FALSE** — $n>7.4$ হলে $\log(n)>2$, তাই **BIC** বেশি শাস্তি দেয়। *B for Bigger.* |
| **W23 1c(ii)** | *"When a predictor unrelated to the response is added, the AIC **must decrease**"* | ❌ **FALSE** — সাধারণত **বাড়ে**; আর "must" দুই দিকেই বড্ড কড়া |

## প্রশ্ন ২৪ — Information criteria আসলে *কীসের জন্য*

> **S25, Ex 4(c) [২ নম্বর].** *"Explain what the intuition behind information criteria is. Explain why their design makes them a good tool for model selection."*

### সমাধান

একটা information criterion **fit-এর মান** আর **মডেলের জটিলতা**-কে একটামাত্র সংখ্যায় ভারসাম্য করে:

$$\text{IC}=\underbrace{n\log(\hat\sigma^2_{ML})}_{\text{fit — covariate বাড়লে কমে}}+\underbrace{c\cdot(|M|+1)}_{\text{penalty — covariate বাড়লে বাড়ে}}$$

$R^2$ বা RSS-এর মতো খাঁটি fit-এর মাপকাঠি মডেলের আকারের সাথে **একমুখীভাবে** ভালো হতে থাকে, তাই সেগুলো optimise করলে সবসময় সবচেয়ে বড় মডেলটাই বেছে নেওয়া হয় আর **overfit** হয়: বাড়তি parameter-গুলো এমন noise-এর পেছনে ছোটে যা নতুন ডেটায় আর ফিরবে না। Penalty term-টা যোগ করা covariate-কে **নিজের দাম উসুল করতে বাধ্য করে** — সেটা তখনই থাকে যখন $n\log\hat\sigma^2$-কে তার খরচের চেয়ে বেশি কমায়।

এটাই **bias–variance trade-off**-কে কার্যকর রূপ দেওয়া: বেশি জটিলতা মানে কম bias কিন্তু বেশি variance, আর criterion-টা U-আকৃতির মোট ত্রুটির সর্বনিম্ন বিন্দু খুঁজে দেয়। **AIC** ($c=2$) পূর্বাভাসের নিশানা নেয়; **BIC** ($c=\log n$) কড়া শাস্তি দিয়ে সত্যিকারের মডেল চিহ্নিত করার নিশানা নেয়, ফলে আরও মিতব্যয়ী নির্বাচন আসে।

## প্রশ্ন ২৫ — Term যোগ করা উচিত কিনা কীভাবে ঠিক করবে?

> **S25, Ex 2(d) [২ নম্বর].** *"A colleague suggests adding many more functions of age, e.g. $(\text{age}-48)^3$, $\text{age}^4$ or $\log(\text{age})$ because it will improve the fit. How would you decide whether you should do so? **State two methods.**"*

### সমাধান — যেকোনো দুটো বলো

1. **Information criteria** — বাড়তি term সহ আর ছাড়া AIC/BIC বের করো, ছোট মানেরটা রাখো। Fit তো এমনিতেই ভালো হবে, আর ঠিক এজন্যই penalty-যুক্ত criterion লাগে।
2. **যোগ করা term-গুলোর উপর F-test** — general linear hypothesis $H_0:$ নতুন সব coefficient $=0$, যেখানে $r$ = যোগ করা term-এর সংখ্যা।
3. Cross-validation / আলাদা রাখা ডেটায় predictive MSE।
4. Adjusted $R^2$, বা Mallow's $C_p$।
5. Diagnostic plot দেখে যাচাই করা যে residual-এর প্যাটার্ন সত্যিই ভালো হলো কিনা।

> **W23, Ex 4(d) [১ নম্বর]** একই জিনিস জিজ্ঞেস করে আর key মেনে নেয়: *"the F-test, test for significance of the polynomial term, Mallows CP value, predicted MSE, or simply observe the diagnostic plot."*

## প্রশ্ন ২৬ — চারটা diagnostic plot পড়া

> **W23, Ex 4(a) [৩ নম্বর].** `mpg ~ horsepower`, diagnostic plot দেওয়া। *"Briefly discuss three problems that the model may have."*
> **অফিসিয়াল key: প্রতিটায় ১ নম্বর, ব্যাখ্যা ছাড়া শুধু নাম বললে ০.৫ কাটা।**

### সমাধান

1. **Non-linearity** — residuals vs fitted-এ আকারহীন ব্যান্ডের বদলে স্পষ্ট **বাঁকানো (U-আকৃতির)** প্যাটার্ন, তাই সরলরেখার specification ভুল (A1 লঙ্ঘিত)।
2. **Heteroscedasticity** — fitted value বাড়ার সাথে residual-এর বিস্তার **চওড়া হচ্ছে** (funnel), তাই $\text{Var}(\varepsilon_i)$ স্থির নয় (A3 লঙ্ঘিত)।
3. **Non-normality** — QQ plot-এ বিন্দুগুলো লেজের দিকে **৪৫° কর্ণ থেকে সরে যাচ্ছে**, যা skew/ভারী লেজ নির্দেশ করে (A6 প্রশ্নবিদ্ধ)।

> 🔴 **নাম বলো *এবং* কী দেখলে সেটাও বলো।** শুধু "heteroscedasticity" লিখলে আধা নম্বর কাটা; "heteroscedasticity — fitted value-এর সাথে residual-এর বিস্তার চওড়া হচ্ছে" লিখলে পুরো নম্বর।

> **W23, Ex 4(b) [২ নম্বর].** *$\text{horsepower}^2$ যোগ করলে কি সমস্যা মেটে?* — এটা **কেবল non-linearity**-র সমাধান করে। Quadratic mean function-এর আকার বদলায়; অসম error variance বা non-normal error নিয়ে কিছুই করে না — ওগুলোর জন্য $y$-এর রূপান্তর, weighted least squares, বা robust standard error লাগবে।

## প্রশ্ন ২৭ 🔴 — Standardised residual

> **W22, Ex 1c(i).** *"Since the residuals are normally distributed, standardized residuals are also normally distributed."*

### **FALSE** (আর কেন, সেটাও লেখো)।

নিখুঁত assumption-এর অধীনেও $\text{Var}(\hat\varepsilon_i)=\sigma^2(1-h_{ii})$ **স্থির নয়**, আর $\text{Cov}(\hat\varepsilon_i,\hat\varepsilon_j)=-\sigma^2h_{ij}\neq0$। $\hat\sigma\sqrt{1-h_{ii}}$ দিয়ে standardise করার *পুরো কারণই* হলো কাঁচা residual heteroscedastic আর correlated। আর standardised residual-ও হুবহু normal নয়, কারণ $\hat\sigma$ একই ডেটা থেকে estimate করা।

## প্রশ্ন ২৮ 🔴 — যে QQ-plot বাক্যটা শেষে গিয়ে ভাঙে

> **W22, Ex 1a(iv).** *"In Q-Q plots, the empirical quantiles are compared to the quantiles of the theoretical distribution. If the data follows the distribution, the points should closely follow a **horizontal line through the origin**."*

### **FALSE.**

প্রথম খণ্ড ঠিক, দ্বিতীয় খণ্ড ভুল: বিন্দুগুলো **৪৫° কর্ণ** $y=x$ অনুসরণ করবে। Horizontal রেখা মানে হতো প্রতিটা empirical quantile অভিন্ন — অর্থাৎ **কোনো ভিন্নতাই নেই**।

> 🛟 **এই কোর্সের আদর্শ ফাঁদ-গঠন।** শেষ শব্দ পর্যন্ত পড়ো।

## প্রশ্ন ২৯ 🔴 — Heteroscedasticity: bias না efficiency?

> **S25, Ex 4(e) [১ নম্বর].** *"The variation in revenue grows as the number of employees grows… Which impact does this phenomenon have on the OLS estimate in terms of bias and efficiency?"*

### সমাধান

ঘটনাটা হলো **heteroscedasticity** (A3 লঙ্ঘিত)।

> **Bias:** নেই। OLS **unbiased** আর consistent থাকে — unbiasedness-এর জন্য কেবল correct specification, শূন্য-গড় error আর full rank লাগে।
>
> **Efficiency:** হারায়। Gauss–Markov আর খাটে না, তাই $\hat{\boldsymbol\beta}$ **BLUE নয়**; weighted/generalised least squares-এর variance কম হতো।
>
> **এবং:** স্বাভাবিক standard error-গুলো biased, তাই t-test, F-test আর CI **অবৈধ** — যদিও point estimate ঠিকই আছে।

> **প্রশ্নটা নিজেই দুটো শব্দ বলে দিচ্ছে** — "bias and efficiency" — তাই ঠিক ওই দুই পরিভাষাতেই উত্তর দাও। **W23 1b(ii)**, *"correlated residuals do not affect consistency but may affect efficiency"*, একই যুক্তিতে ✅ **TRUE**। **LMES 1b(iv)**, *"traditional standard errors remain valid under heteroskedasticity"*, ❌ **FALSE**।

## প্রশ্ন ৩০ 🔴 — Multicollinearity, চারভাবে

| পেপার | বক্তব্য | রায় |
|---|---|---|
| **S25 1(j)** | *"Multicollinearity can inflate the variance of the OLS coefficient estimators."* | ✅ **TRUE** |
| **LMES 1b(i)** | *"…can inflate the variance but does not bias the estimates themselves."* | ✅ **TRUE** |
| **W23 1b(i)** | *"Highly correlated explanatory variables may lead to a **reduction** in the standard errors."* | ❌ **FALSE** — বাড়ে, কমে না |
| **LMES 1b(ii)** | *"If VIF for all variables is close to 1, multicollinearity is **likely a concern**."* | ❌ **FALSE** — VIF ≈ ১ মানে **কোনো** collinearity নেই |

> **Perfect আর near আলাদা রাখো:** perfect collinearity A5 ভাঙে (একেবারেই চিহ্নিত হয় না); near collinearity-তে $\hat{\boldsymbol\beta}$ **unbiased আর তখনও BLUE**, শুধু অনিখুঁত।

## প্রশ্ন ৩১ — Leverage, outlier, influence

| পেপার | বক্তব্য | রায় |
|---|---|---|
| **LMES 1c(i)** | *"Cook's distance helps identify points that might be unduly influencing the model's fit."* | ✅ **TRUE** |
| **LMES 1c(ii)** | *"The hat matrix provides the leverages which help identify potential outliers."* | ✅ **TRUE** *(অফিসিয়াল key অনুযায়ী — যদিও কড়াভাবে বললে leverage অস্বাভাবিক $\boldsymbol{x}$ ধরে, অস্বাভাবিক $y$ নয়)* |
| **LMES 1a(ii)** | *"Residual plots **cannot** be used to identify non-linearity."* | ❌ **FALSE** — residuals-vs-fitted ঠিক এই কাজটাই করে |

$$h_{ii}\ (\text{অস্বাভাবিক }\boldsymbol{x})\qquad |r_i|\ (\text{অস্বাভাবিক }y)\qquad D_i=\frac{r_i^2}{p}\cdot\frac{h_{ii}}{1-h_{ii}}\ (\text{দুটোই})$$

**একা উঁচু leverage ক্ষতিকর নয়** — বড় residual-এর সাথে জুটলে তবেই ক্ষতি করে।

## প্রশ্ন ৩২ — Residual-এর গড়

> **W23, Ex 1b(iv).** *"In a well-fitted linear regression model, the mean of the residuals should be close to zero."* → ✅ **TRUE** *(অফিসিয়াল key)*

কিন্তু বুঝে রাখো **এটা কেন দুর্বল বক্তব্য**: intercept থাকলে $\sum\hat\varepsilon_i=0$ **গঠনগতভাবেই** সত্য, ভালো মডেলের মতোই জঘন্য মডেলেও। তাই diagnostics সবসময় residual-এর **প্যাটার্ন** পড়ে, কখনো তার গড় নয়।

## প্রশ্ন ৩৩ — বিবিধ

> **W22, Ex 1c(iii).** *"Interaction can only be computed between two continuous or between a continuous and a categorical independent variable."* → ❌ **FALSE** — দুটো **categorical** ভেরিয়েবলের মধ্যেও interaction হয় (দুটো dummy-র গুণফল), যা ঠিক S25-এর wage মডেলের সম্প্রসারণ।
>
> **W22, Ex 1a(iii).** *"For $H_0:\beta_j=0$, under $H_0$, $t_j^2\sim F_{1,\,n-p}$."* → ✅ **TRUE** — $r=1$-এর সেই $\sqrt{F}=|t|$ সম্পর্ক, বর্গ করা। $r=1$ হলেই এটা যাচাই হিসেবে ব্যবহার করো।

---

# পাঁচ পেপারের প্যাটার্ন

| প্রশ্নের ধরন | S25 | LMES | W23 | W22 | EX20 | যখন আসে তখনকার নম্বর |
|---|---|---|---|---|---|---|
| R output-এর ফাঁকা পূরণ | ✅ | ✅ | | | ✅ | ২.৫–১২ |
| OLS derive করা | ✅ | ✅ | ✅ | | | ২–৩ |
| Gauss–Markov / BLUE | ✅ | | ✅ | | | ৪ |
| $\boldsymbol{C},\boldsymbol{d}$ বানানো, $r$ গোনা | ✅ | | | ✅ | | ১–২ |
| F-test (সামগ্রিক বা আংশিক) | ✅ | ✅ | | ✅ | | ১.৫–৮ |
| Confidence interval | ✅ | ✅ | ✅ | ✅ | | ১–৩.৫ |
| পূর্ণাঙ্গ t-test লেখা | | ✅ | ✅ | | | ২ |
| AIC / BIC | | ✅ | ✅ | ✅ | | ১–৪ |
| Diagnostics / residual plot | | ✅ | ✅ | ✅ | | ১–৩ |
| Heteroscedasticity-র পরিণতি | ✅ | ✅ | ✅ | | | ১ |
| Dummy / interaction | ✅ | | | ✅ | ✅ | ৩–২৫ |
| Logit | ✅ | | | | | ১–১.৫ |

**এর প্রতিটাই অন্তত দুটো পেপারে আছে।** এই কোর্সে এমন কোনো প্রশ্নের ধরন নেই যেটা তুমি পরীক্ষার হলে ঢোকার আগে দেখোনি।

> **এরপর:** `../99-exam-vault/30-TRUE-FALSE-DRILL.md`-এর T/F আইটেমগুলো drill করো, তারপর ৬০ মিনিটে বই বন্ধ রেখে একটা পুরো পেপার দাও।
