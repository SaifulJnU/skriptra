# 3.3 — Hypothesis Testing: the t-test and the General Linear Hypothesis

> **The most heavily examined section in the course.** Every past paper has multiple questions here. Two skills: the **t-test** (easy, always appears) and **building $\boldsymbol{C}\boldsymbol\beta=\boldsymbol{d}$** (surprises people every year — it shouldn't surprise you).

---

## 1. The sampling distribution — where all tests come from

Under A1–A6:

$$\hat{\boldsymbol\beta}\sim N\!\left(\boldsymbol\beta,\ \sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}\right)$$

For a single coefficient:

$$\hat\beta_j\sim N\!\left(\beta_j,\ \sigma^2\left[(\boldsymbol{X}'\boldsymbol{X})^{-1}\right]_{jj}\right) \qquad\Longrightarrow\qquad \frac{\hat\beta_j-\beta_j}{\sigma\sqrt{[(\boldsymbol{X}'\boldsymbol{X})^{-1}]_{jj}}}\sim N(0,1)$$

But $\sigma$ is unknown. Replace it by $\hat\sigma$, and because $\frac{(n-p)\hat\sigma^2}{\sigma^2}\sim\chi^2_{n-p}$ independently of $\hat{\boldsymbol\beta}$, the standard normal becomes a **$t$-distribution**:

$$\boxed{\;\frac{\hat\beta_j-\beta_j}{\widehat{\text{se}}(\hat\beta_j)}\sim t_{n-p}\;}\qquad \widehat{\text{se}}(\hat\beta_j)=\hat\sigma\sqrt{\left[(\boldsymbol{X}'\boldsymbol{X})^{-1}\right]_{jj}}$$

> **Why $t$ and not $z$?** Because you estimated $\sigma$. That extra uncertainty makes the tails heavier. As $n-p\to\infty$, $t\to N(0,1)$ — which is why with $n=3000$ the exercise sheets happily hand you normal quantiles like $1.9608$ instead of $t$ quantiles.

---

## 2. The t-test for a single coefficient

### Test of significance: $H_0:\beta_j=0$

$$H_0:\beta_j=0 \quad\text{vs}\quad H_1:\beta_j\neq0$$

$$\boxed{\;t_j=\frac{\hat\beta_j}{\widehat{\text{se}}(\hat\beta_j)}\;\overset{H_0}{\sim}\;t_{n-p}\;}$$

**Reject $H_0$ at level $\alpha$ if** $\;|t_j|>t_{n-p}(1-\alpha/2)$.

### General $H_0:\beta_j=c$

$$t=\frac{\hat\beta_j-c}{\widehat{\text{se}}(\hat\beta_j)}\overset{H_0}{\sim}t_{n-p}$$

Same critical value. **Don't forget the $-c$** — a favourite trap.

### Interpreting the R output columns

| Column | What it is |
|---|---|
| `Estimate` | $\hat\beta_j$ |
| `Std. Error` | $\widehat{\text{se}}(\hat\beta_j)$ |
| `t value` | $\hat\beta_j/\widehat{\text{se}}(\hat\beta_j)$ |
| `Pr(>|t|)` | two-sided p-value $=2\cdot P(T_{n-p}>|t_j|)$ |

**Reject at level $\alpha$ ⟺ p-value $<\alpha$.**

Significance stars: `***` $<0.001$, `**` $<0.01$, `*` $<0.05$, `.` $<0.1$.

> 💰 **This gives you three ways to fill any hole in a regression table:**
> $$\hat\beta=t\times\widehat{\text{se}},\qquad \widehat{\text{se}}=\frac{\hat\beta}{t},\qquad t=\frac{\hat\beta}{\widehat{\text{se}}}$$

### 📝 Worked: Exam Summer 2025, Ex 3(a) [2.5 pts]

```
             Estimate Std. Error t value  Pr(>|t|)
(Intercept)  48.38458   [[A]]     13.591  < 2e-16
crim         -0.25959   0.05302   -4.896  1.32e-06
nox         -36.99122   5.25574   [[B]]   6.44e-12
dis           [[C]]     0.26423   -3.796  0.000165
rad          -0.06165   0.05983   -1.030  0.303290
Residual standard error: [[D]] on 501 degrees of freedom
sum(lm$residuals^2) = 31682.02
```

$$\textbf{[[A]]}=\frac{\hat\beta_0}{t}=\frac{48.38458}{13.591}=\boxed{3.560}$$
$$\textbf{[[B]]}=\frac{\hat\beta_{\text{nox}}}{\widehat{\text{se}}}=\frac{-36.99122}{5.25574}=\boxed{-7.038}$$
$$\textbf{[[C]]}=t\times\widehat{\text{se}}=-3.796\times0.26423=\boxed{-1.003}$$
$$\textbf{[[D]]}=\hat\sigma=\sqrt{\frac{31682.02}{501}}=\sqrt{63.238}=\boxed{7.952}$$

**Fitted regression formula:**
$$\widehat{\text{medv}}=48.385-0.260\,\text{crim}-36.991\,\text{nox}-1.003\,\text{dis}-0.062\,\text{rad}$$

*(Note the write-up asks for the fitted formula too. Don't skip it — it's part of the 2.5 points.)*

### 📝 Worked: Sheet 4, Ex 3 — t-test with a given $(\boldsymbol{X}'\boldsymbol{X})^{-1}$

Model: wage on age + 4 education dummies + health, $n=3000$, $p=7$, $\hat\beta_1=0.62$ (age).

From Sheet 3: $\hat\sigma^2=\dfrac{3819720}{2993}=1276.218$, so $\hat\sigma=35.724$.
From the given matrix: $[(\boldsymbol{X}'\boldsymbol{X})^{-1}]_{22}=0.26\times10^{-5}=2.6\times10^{-6}$.

$$\widehat{\text{se}}(\hat\beta_1)=35.724\times\sqrt{2.6\times10^{-6}}=35.724\times0.0016125=0.0576$$

$$t=\frac{0.62}{0.0576}=10.76$$

$|t|=10.76>1.9608=$ the 0.975 quantile ⟹ **reject $H_0:\beta_1=0$** at $\alpha=0.05$. The effect of age on wage is significantly different from zero.

> ⚠️ **Read the $(\boldsymbol{X}'\boldsymbol{X})^{-1}$ index carefully.** The matrix is indexed $0,\dots,6$ for $\beta_0,\dots,\beta_6$, but printed as rows $1,\dots,7$. **The age entry is row 2, column 2** of the printed matrix (second diagonal element), because row 1 is the intercept. Off-by-one here is a guaranteed lost mark. Sheet 3(e) is precisely testing whether you can find the right diagonal element.

---

## 3. ⭐ The general linear hypothesis $\boldsymbol{C}\boldsymbol\beta=\boldsymbol{d}$

Most hypotheses aren't "$\beta_j=0$." They can be several restrictions at once, or involve combinations of coefficients. All of them fit one framework:

$$\boxed{\;H_0:\boldsymbol{C}\boldsymbol\beta=\boldsymbol{d}\quad\text{vs}\quad H_1:\boldsymbol{C}\boldsymbol\beta\neq\boldsymbol{d}\;}$$

| Object | Dimension | Meaning |
|---|---|---|
| $\boldsymbol{C}$ | $r\times p$ | restriction matrix — **one row per restriction** |
| $\boldsymbol\beta$ | $p\times1$ | the full parameter vector, **including $\beta_0$** |
| $\boldsymbol{d}$ | $r\times1$ | the values the restrictions equal |
| $r$ | scalar | number of **linearly independent** restrictions $=\text{rank}(\boldsymbol{C})$ |

### 🔧 The three-step recipe

**Step 1.** Write each restriction with **everything on the left, constants on the right.**
**Step 2.** For each restriction, read off the coefficient of each $\beta_j$ in order $\beta_0,\beta_1,\dots,\beta_k$. That's one row of $\boldsymbol{C}$. The constant goes into $\boldsymbol{d}$.
**Step 3.** $r$ = number of rows (assuming they're independent).

**Never forget the $\beta_0$ column.** $\boldsymbol{C}$ has $p$ columns, not $k$.

### Worked examples — practise until instant

Take $\boldsymbol\beta=(\beta_0,\beta_1,\beta_2,\beta_3,\beta_4)'$, so $p=5$.

**(a) $H_0:\beta_2=0$** (a single significance test)
$$\boldsymbol{C}=\begin{pmatrix}0&0&1&0&0\end{pmatrix},\quad \boldsymbol{d}=(0),\quad r=1$$

**(b) $H_0:\beta_1=\beta_2$** → rewrite as $\beta_1-\beta_2=0$
$$\boldsymbol{C}=\begin{pmatrix}0&1&-1&0&0\end{pmatrix},\quad \boldsymbol{d}=(0),\quad r=1$$

**(c) $H_0:\beta_1=\beta_2+1$** → rewrite as $\beta_1-\beta_2=1$
$$\boldsymbol{C}=\begin{pmatrix}0&1&-1&0&0\end{pmatrix},\quad \boldsymbol{d}=(1),\quad r=1$$

**(d) $H_0:\beta_1=\beta_2=\beta_3=\beta_4=0$** (test of overall significance)
$$\boldsymbol{C}=\begin{pmatrix}0&1&0&0&0\\0&0&1&0&0\\0&0&0&1&0\\0&0&0&0&1\end{pmatrix}=\begin{pmatrix}\boldsymbol{0}&\boldsymbol{I}_4\end{pmatrix},\quad \boldsymbol{d}=\boldsymbol{0}_{4\times1},\quad r=4$$

**(e) $H_0:\beta_1=-\beta_2+\beta_3$** → rewrite as $\beta_1+\beta_2-\beta_3=0$
$$\boldsymbol{C}=\begin{pmatrix}0&1&1&-1&0\end{pmatrix},\quad \boldsymbol{d}=(0),\quad \boxed{r=1}$$

> 🔴 **Exam Summer 2025, Ex 1(i):** *"The F-statistic for testing $H_0:\beta_1=-\beta_2+\beta_3$ in a model with $k\geq3$ predictors plus an intercept has an $F$-distribution with $(3,\ n-k-1)$ degrees of freedom under $H_0$."* → **FALSE.**
>
> **Why:** this is **ONE** equation, hence **one** restriction: $r=1$, not 3. The correct distribution is $F_{1,\,n-k-1}$. The "3" is bait — it counts the *betas mentioned*, not the *restrictions imposed*.
>
> 🔑 **Rule: $r$ = the number of EQUATIONS (independent rows of $\boldsymbol{C}$), never the number of parameters appearing in them.**

### 📝 Worked: Exam Summer 2025, Ex 3(c) [2 pts]

> $H_0:\ \beta_{\text{crim}}=3\beta_{\text{rad}}-0.1,\quad \beta_{\text{nox}}=-40$, with $\boldsymbol\beta=(\beta_0,\beta_{\text{crim}},\beta_{\text{nox}},\beta_{\text{dis}},\beta_{\text{rad}})'$.

**Rearrange:**
$$\beta_{\text{crim}}-3\beta_{\text{rad}}=-0.1$$
$$\beta_{\text{nox}}=-40$$

**Read off the coefficients in the order $(\beta_0,\beta_{\text{crim}},\beta_{\text{nox}},\beta_{\text{dis}},\beta_{\text{rad}})$:**

$$\boxed{\;\boldsymbol{C}=\begin{pmatrix}0&1&0&0&-3\\0&0&1&0&0\end{pmatrix},\qquad \boldsymbol{d}=\begin{pmatrix}-0.1\\-40\end{pmatrix},\qquad r=2\;}$$

**Check:** $\boldsymbol{C}$ is $2\times5$ ✓ ($r\times p$). The two rows are clearly linearly independent, so $r=\text{rank}(\boldsymbol{C})=2$. ✓

### 📝 Worked: Sheet 4, Ex 1(c) and 2(c)

Model with $\boldsymbol\beta=(\beta_0,\dots,\beta_6)'$, $p=7$.

**Ex 1:** $H_0:(\beta_1,\beta_2,\beta_3,\beta_4,\beta_5,\beta_6)'=\boldsymbol{0}$ — the **test of significance of regression**.

$$\boldsymbol{C}=\begin{pmatrix}\boldsymbol{0}_{6\times1}&\boldsymbol{I}_6\end{pmatrix}=\begin{pmatrix}
0&1&0&0&0&0&0\\
0&0&1&0&0&0&0\\
0&0&0&1&0&0&0\\
0&0&0&0&1&0&0\\
0&0&0&0&0&1&0\\
0&0&0&0&0&0&1
\end{pmatrix},\quad \boldsymbol{d}=\boldsymbol{0}_{6\times1},\quad r=6$$

**Model under $H_0$:** $\;\text{wage}_i=\beta_0+\varepsilon_i$ — the intercept-only model, whose fitted value is $\bar y$.

**Ex 2:** $H_0:\begin{pmatrix}\beta_1\\\beta_6\end{pmatrix}=\boldsymbol{0}$ — a **composite test of a subvector**.

$$\boldsymbol{C}=\begin{pmatrix}0&1&0&0&0&0&0\\0&0&0&0&0&0&1\end{pmatrix},\quad \boldsymbol{d}=\begin{pmatrix}0\\0\end{pmatrix},\quad r=2$$

**Model under $H_0$:** $\;\text{wage}_i=\beta_0+\beta_2\text{HS}_i+\beta_3\text{SC}_i+\beta_4\text{CG}_i+\beta_5\text{AD}_i+\varepsilon_i$ — drop age and health.

### Stating $H_1$ properly

$$H_1:\boldsymbol{C}\boldsymbol\beta\neq\boldsymbol{d}$$

**In plain English** (Sheet 4 asks for this explicitly): *"at least one of the restrictions fails."* **Not** "all of them fail."

> For $H_0:\beta_1=\dots=\beta_6=0$: $H_1$ is *"at least one $\beta_j$, $j=1,\dots,6$, is non-zero"* — i.e. at least one covariate has explanatory power for wage. Rejecting does **not** tell you which one.

---

## 4. Test statistics for the general hypothesis

$$\boxed{\;F=\frac{1}{r}(\boldsymbol{C}\hat{\boldsymbol\beta}-\boldsymbol{d})'\left[\hat\sigma^2\boldsymbol{C}(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{C}'\right]^{-1}(\boldsymbol{C}\hat{\boldsymbol\beta}-\boldsymbol{d})\;\overset{H_0}{\sim}\;F_{r,\,n-p}\;}$$

You will essentially never compute this form by hand. Section 3.3.1 gives the practical SSE-based version. But **know that it exists and what it measures**: how far $\boldsymbol{C}\hat{\boldsymbol\beta}$ is from $\boldsymbol{d}$, scaled by how uncertain that quantity is.

### 🔑 The $t$–$F$ relationship

When $r=1$:

$$\boxed{\;F=t^2\;}\qquad\text{and}\qquad t^2_{n-p}\sim F_{1,\,n-p}$$

**The two tests are identical for a single restriction.** They always give the same decision and the same p-value.

> 🔴 *RCLM WS 22/23, Block I(iii):* "For $H_0:\beta_j=0$, under $H_0$, $t_j^2\sim F_{1,n-p}$." → **TRUE.**

**Useful check:** if you compute an F for one restriction and $\sqrt{F}$ doesn't match the t-value in the output, one of them is wrong.

---

## 5. Degrees of freedom — the recurring trap

| Test | Distribution under $H_0$ |
|---|---|
| $t$-test for one coefficient | $t_{n-p}$ |
| $F$-test for $r$ restrictions | $F_{r,\,n-p}$ |
| Overall $F$ (all $k$ slopes zero) | $F_{k,\,n-k-1}=F_{p-1,\,n-p}$ |

With $k$ covariates plus an intercept, $p=k+1$, so $n-p=n-k-1$. **Always.**

**Past-paper degrees-of-freedom statements:**

| Statement | Verdict |
|---|---|
| $t_j\sim t_{n-k-1}$ where $k+1$ = #parameters *(Exam 2025 Ex 1g)* | ✅ **TRUE** |
| $t_j\sim t_{n-p-1}$ where $p$ = #covariates | ✅ TRUE *(same number)* |
| Full-model F test: $F\sim F_{k,\,n-k-1}$ *(WS23/24 I(iii))* | ✅ **TRUE** |
| $F\sim F_{p+1,\,n-p}$ for $H_0:\boldsymbol{C}\boldsymbol\beta=\boldsymbol{d}$ *(WS22/23 II(iv))* | ❌ **FALSE** — first df is $r$, not $p+1$ |
| Significance-of-regression: $F\sim F_{p-1,\,n-p}$ *(WS22/23 III(ii))* | ✅ TRUE in book notation ($p-1=k$ restrictions) |

> 🛟 **Your defence, every time:** count the betas you estimated in the **unrestricted** model → that's $p$ → second df is $n-p$. Count the **equations** in $H_0$ → that's $r$ → first df is $r$.

---

## 6. Decision, p-values, and language

### The decision rule

$$\text{Reject } H_0 \text{ at level } \alpha \iff F>F_{r,n-p}(1-\alpha) \iff |t|>t_{n-p}(1-\alpha/2) \iff p\text{-value}<\alpha$$

### ⚠️ Say it correctly

| ❌ Don't write | ✅ Write |
|---|---|
| "We accept $H_0$" | "We **fail to reject** $H_0$" / "We **cannot reject** $H_0$" |
| "$H_0$ is true" | "There is insufficient evidence against $H_0$" |
| "The effect is zero" | "The effect is **not significantly different from zero** at the 5% level" |
| "p is the probability $H_0$ is true" | "p is the probability of a test statistic this extreme **if $H_0$ were true**" |

**Failing to reject is not evidence of no effect** — it may just mean low power (small $n$, high variance, multicollinearity). Markers do notice this language.

### ⚠️ Individual vs joint significance

A subtle and genuinely examinable point:

- **Individually insignificant, jointly significant** — happens under multicollinearity. Two correlated covariates each have inflated standard errors, so neither $t$-test rejects, but the $F$-test for both together does. *Neither is needed **given** the other, but at least one is needed.*
- **Individually significant, jointly insignificant** — rarer, but possible.

> **The exam framing:** "Rejecting $H_0:\beta_5=\beta_6=0$ does not imply both coefficients differ from zero; it implies **at least one** does." The book makes exactly this remark in its kitchen-quality example.

---

## 7. Key takeaways

1. $\dfrac{\hat\beta_j-\beta_j}{\widehat{\text{se}}(\hat\beta_j)}\sim t_{n-p}$; $\;\widehat{\text{se}}(\hat\beta_j)=\hat\sigma\sqrt{[(\boldsymbol{X}'\boldsymbol{X})^{-1}]_{jj}}$ — **diagonal element, right index**.
2. $t=\hat\beta_j/\widehat{\text{se}}$ for significance; $t=(\hat\beta_j-c)/\widehat{\text{se}}$ in general. **Don't forget the $-c$.**
3. **$\hat\beta=t\times\widehat{\text{se}}$** and its rearrangements fill any hole in an R table.
4. **$\boldsymbol{C}$ is $r\times p$: one row per restriction, one column per parameter — including $\beta_0$.**
5. 🔴 **$r$ = number of EQUATIONS, not number of betas mentioned.** $\beta_1=-\beta_2+\beta_3$ is $r=1$.
6. $F\sim F_{r,\,n-p}$ under $H_0$; for $r=1$, $F=t^2$.
7. Under $H_0$, restate the **model equation** — you'll need it for $\text{SSE}_{H_0}$ in 3.3.1.
8. $H_1$ = "**at least one** restriction fails." Rejecting doesn't identify which.
9. Say "**fail to reject**," never "accept."
