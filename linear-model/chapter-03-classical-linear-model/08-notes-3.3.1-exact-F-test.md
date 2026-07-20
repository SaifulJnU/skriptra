# 3.3.1 — The Exact F-Test

> **The single most computed thing in this exam.** Sheet 4 is three F-tests. Exam Summer 2025 Ex 3(d) is an F-test. The Example Exam has one. Learn the two practical formulas and when each applies.

---

## 1. The idea

Compare two nested models:

| | Model | Parameters | Residual SS |
|---|---|---|---|
| **Unrestricted** | $\boldsymbol{y}=\boldsymbol{X}\boldsymbol\beta+\boldsymbol\varepsilon$ | $p$ | $\text{SSE}=\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}$ |
| **Restricted** ($H_0$ imposed) | $\boldsymbol{C}\boldsymbol\beta=\boldsymbol{d}$ | $p-r$ | $\text{SSE}_{H_0}=\hat{\boldsymbol\varepsilon}'_{H_0}\hat{\boldsymbol\varepsilon}_{H_0}$ |

The restricted model has fewer free parameters, so it **cannot fit better**:

$$\text{SSE}_{H_0}\geq\text{SSE}\qquad\text{always}$$

**The question:** is the extra residual sum of squares you incur by imposing $H_0$ *large*, relative to the noise level?

- Small increase ⟹ the restriction costs little ⟹ $H_0$ is plausible ⟹ **don't reject**
- Large increase ⟹ the restriction badly damages the fit ⟹ **reject**

---

## 2. ⭐ Formula 1 — the SSE version (use this by default)

$$\boxed{\;F=\frac{(\text{SSE}_{H_0}-\text{SSE})/r}{\text{SSE}/(n-p)}=\frac{n-p}{r}\cdot\frac{\text{SSE}_{H_0}-\text{SSE}}{\text{SSE}}\;\overset{H_0}{\sim}\;F_{r,\,n-p}\;}$$

**Reading it as a ratio of two variance estimates:**

$$F=\frac{\text{extra residual SS per restriction}}{\text{residual variance }\hat\sigma^2}$$

The denominator is exactly $\hat\sigma^2=\text{SSE}/(n-p)$. So:

$$F=\frac{(\text{SSE}_{H_0}-\text{SSE})/r}{\hat\sigma^2}$$

Under $H_0$ both numerator and denominator estimate $\sigma^2$, so $F\approx1$. Values much larger than 1 are evidence against $H_0$.

**Use this formula when the question gives you SSE and $\text{SSE}_{H_0}$** — which it usually does.

---

## 3. Formula 2 — the $R^2$ version

$$\boxed{\;F=\frac{(R^2-R^2_{H_0})/r}{(1-R^2)/(n-p)}\;}$$

Since $R^2=1-\text{SSE}/\text{SST}$, dividing the SSE version through by SST gives this immediately.

**Special case — the test of significance of regression** ($H_0:\beta_1=\dots=\beta_k=0$). Then the restricted model is intercept-only, so $R^2_{H_0}=0$ and $r=k=p-1$:

$$\boxed{\;F=\frac{R^2/k}{(1-R^2)/(n-p)}=\frac{R^2}{1-R^2}\cdot\frac{n-p}{k}\;\overset{H_0}{\sim}\;F_{k,\,n-p}\;}$$

**This is the "F-statistic" line at the bottom of every R summary.**

> From Exam Summer 2025's output: `F-statistic: 43.62 on 4 and 501 DF`. Check: $k=4$ covariates, $n-p=501$. ✓
> And indeed $\frac{0.2583}{1-0.2583}\cdot\frac{501}{4}=43.62$ ✓

---

## 4. 📝 Worked example 1 — Exam Summer 2025, Ex 3(d) [1.5 pts]

> $H_0:\beta_{\text{crim}}=3\beta_{\text{rad}}-0.1,\ \beta_{\text{nox}}=-40$. Given $\hat{\boldsymbol\varepsilon}'_{H_0}\hat{\boldsymbol\varepsilon}_{H_0}=32333.15$ and $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}=31682.02$. Test at $\alpha=0.05$.

**Step 1 — identify the pieces.** From part (c): $r=2$. From the output: $n-p=501$.

**Step 2 — compute.**

$$F=\frac{(\text{SSE}_{H_0}-\text{SSE})/r}{\text{SSE}/(n-p)}=\frac{(32333.15-31682.02)/2}{31682.02/501}=\frac{651.13/2}{63.238}=\frac{325.565}{63.238}=\boxed{5.148}$$

**Step 3 — critical value.** From the paper's Table 2 (F-distributions with df$_2$ = 501), row $r=2$, column 0.95:

$$F_{2,501}(0.95)=3.0137$$

**Step 4 — decide.**

$$F=5.148>3.0137=F_{2,501}(0.95)$$

> **We reject $H_0$ at the 5% significance level.** The data provide significant evidence against the joint restriction — at least one of $\beta_{\text{crim}}=3\beta_{\text{rad}}-0.1$ and $\beta_{\text{nox}}=-40$ does not hold.

⚠️ **Note the phrasing of the conclusion.** Rejecting a *joint* hypothesis tells you at least one restriction fails; it does not tell you which.

---

## 5. 📝 Worked example 2 — Sheet 4, Ex 1 (test of significance of regression)

Model: wage on age + 4 education dummies + health. $n=3000$, $p=7$, $R^2=0.2685$.

**(a) Hypotheses.**
$$H_0:(\beta_1,\dots,\beta_6)'=\boldsymbol{0}\qquad H_1:(\beta_1,\dots,\beta_6)'\neq\boldsymbol{0}$$

> *In plain English:* $H_0$ says **none** of age, education or health has any effect on expected wage — the model explains nothing beyond the overall mean. $H_1$ says **at least one** of them has a non-zero effect.

**(b) Model under $H_0$.**
$$\text{wage}_i=\beta_0+\varepsilon_i$$
The intercept-only model; its fitted value is $\bar y$ for everyone.

**(c) Restriction matrix.**
$$\boldsymbol{C}=\begin{pmatrix}\boldsymbol{0}_{6\times1}&\boldsymbol{I}_6\end{pmatrix}\in\mathbb{R}^{6\times7},\qquad \boldsymbol{d}=\boldsymbol{0}_{6\times1},\qquad r=6$$

**(d) F-statistic.** Since $R^2_{H_0}=0$, use the $R^2$ version:

$$F=\frac{R^2/k}{(1-R^2)/(n-p)}=\frac{0.2685/6}{0.7315/2993}=\frac{0.04475}{0.00024440}=\boxed{183.10}$$

**(e) Decision.** Under $H_0$, $F\sim F_{6,\,2993}$. The sheet's table (Table 1) gives $F_{6,\infty}(0.95)=2.1016$ — with $n-p=2993$ the large-df row is the right one to use.

$$F=183.10\gg2.1016$$

> **Reject $H_0$ at $\alpha=0.05$.** Degrees of freedom: $(6,\ 2993)$.
>
> *Interpretation:* the covariates jointly have highly significant explanatory power for wage — the regression as a whole is significant. This does **not** imply every individual coefficient is non-zero.

---

## 6. 📝 Worked example 3 — Sheet 4, Ex 2 (composite test of a subvector)

$H_0:(\beta_1,\beta_6)'=\boldsymbol{0}$ — do age and health jointly matter, given education?

Given: $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}=3819720$, $\hat{\boldsymbol\varepsilon}'_{H_0}\hat{\boldsymbol\varepsilon}_{H_0}=3995721$.

**(c)** $\boldsymbol{C}=\begin{pmatrix}0&1&0&0&0&0&0\\0&0&0&0&0&0&1\end{pmatrix}$, $\boldsymbol{d}=\boldsymbol{0}$, $r=2$.

**(d)**
$$F=\frac{(3995721-3819720)/2}{3819720/2993}=\frac{176001/2}{1276.218}=\frac{88000.5}{1276.218}=\boxed{68.95}$$

**(e)** $F\sim F_{2,\,2993}$; critical value $\approx F_{2,\infty}(0.95)=2.9987$ from the sheet's table.

$$68.95\gg2.9987 \Longrightarrow \textbf{reject } H_0$$

> Age and health **jointly** have significant explanatory power for wage, even after controlling for education. Degrees of freedom: $(2,\ 2993)$.

---

## 7. 🔧 Getting $\text{SSE}_{H_0}$ yourself

Sometimes you must derive the restricted model rather than being handed its SSE. The method from 3.1.3:

**Substitute the restriction → collect terms in the free parameters → move parameter-free covariate terms to the left as an offset.**

> 🔴 **Exam Summer 2025, Ex 4(d)** [2 pts]: model $y_i=\beta_0+\beta_1x_{1i}+\beta_2x_{2i}+\varepsilon_i$, $H_0:\beta_1=\beta_2+1$.
>
> $$y_i=\beta_0+(\beta_2+1)x_{1i}+\beta_2x_{2i}+\varepsilon_i=\beta_0+\beta_2(x_{1i}+x_{2i})+x_{1i}+\varepsilon_i$$
> $$\boxed{\;y_i-x_{1i}=\beta_0+\beta_2(x_{1i}+x_{2i})+\varepsilon_i\;}$$
>
> Regress $\tilde y_i=y_i-x_{1i}$ on $\tilde x_i=x_{1i}+x_{2i}$ by OLS. The resulting residual sum of squares **is** $\text{SSE}_{H_0}$. The restricted model has 2 parameters vs 3, confirming $r=1$.

**More practice cases:**

| $H_0$ | Restricted model | $r$ |
|---|---|---|
| $\beta_1=0$ | drop $x_1$ | 1 |
| $\beta_1=\beta_2$ | regress on $(x_1+x_2)$ | 1 |
| $\beta_1=\beta_2+1$ | regress $y-x_1$ on $(x_1+x_2)$ | 1 |
| $\beta_1+\beta_2=1$ | regress $y-x_2$ on $(x_1-x_2)$ | 1 |
| $\beta_1=2\beta_2$ | regress on $(2x_1+x_2)$ | 1 |
| $\beta_1=\beta_2=0$ | drop both | 2 |

**Always check:** (parameters in unrestricted) − (parameters in restricted) = $r$. If it doesn't, you've made an error.

---

## 8. Reading the F-quantile tables

The exam supplies tables. **Two columns you must not confuse:**

| Table axis | What it is |
|---|---|
| Row label ($r$, or df1) | **numerator** df = number of restrictions |
| Column | the **quantile level** $1-\alpha$ |
| Table caption | **denominator** df = $n-p$ (often fixed) |

The F-test is **always one-sided in the upper tail** — you only reject for large $F$, because $\text{SSE}_{H_0}\geq\text{SSE}$ makes the numerator non-negative by construction.

> ⚠️ **So for $\alpha=0.05$ you use the 0.95 column, not 0.975.** This is different from the two-sided $t$-test, where $\alpha=0.05$ means the 0.975 quantile. **Mixing these up is one of the most common errors in this exam.**

| Test | $\alpha=0.05$ ⟹ use quantile |
|---|---|
| two-sided $t$-test | **0.975** |
| one-sided $t$-test | 0.95 |
| $F$-test | **0.95** |
| two-sided CI at 95% | **0.975** |
| two-sided CI at 99% | **0.995** |

---

## 9. Asymptotic validity

The exact F-test requires normal errors (A6). Without them, the book notes:

$$W=rF\overset{a}{\sim}\chi^2_r$$

so the $F_{r,n-p}$ distribution converges to $\chi^2_r/r$ as $n\to\infty$. **For large $n$ the F-test remains valid even without normality.** Worth one sentence if asked about robustness.

---

## 10. Common errors — check these every time

| Error | Fix |
|---|---|
| Using $r$ = number of betas mentioned | $r$ = number of **equations** |
| Numerator df = $p$ | numerator df = $r$ |
| Denominator df = $n-r$ | denominator df = $n-p$ (from the **unrestricted** model) |
| Dividing by $\text{SSE}_{H_0}$ | denominator uses the **unrestricted** SSE |
| Using the 0.975 F-quantile for $\alpha=0.05$ | F is one-sided: use **0.95** |
| Getting $F<0$ | impossible — you swapped $\text{SSE}_{H_0}$ and SSE |
| Concluding "all restrictions fail" | conclude "**at least one** fails" |

**Two sanity checks:**
- $F\geq0$ always. Negative means you swapped the SSEs.
- If $r=1$, then $\sqrt{F}$ should equal the $|t|$ from the corresponding t-test.

---

## 11. Key takeaways

1. **The F-test compares nested models** via the increase in residual sum of squares.
2. $$F=\frac{(\text{SSE}_{H_0}-\text{SSE})/r}{\text{SSE}/(n-p)}\sim F_{r,\,n-p}$$
3. $$F=\frac{(R^2-R^2_{H_0})/r}{(1-R^2)/(n-p)}\;;\quad\text{overall test: } F=\frac{R^2/k}{(1-R^2)/(n-p)}$$
4. **df:** numerator $=r$ (equations), denominator $=n-p$ (**unrestricted** model).
5. Get $\text{SSE}_{H_0}$ by **substituting the restriction and re-fitting** — offsets go to the left-hand side.
6. **F is one-sided:** $\alpha=0.05$ ⟹ **0.95** quantile. (Two-sided $t$ at 0.05 ⟹ 0.975.)
7. $F=t^2$ when $r=1$ — use it as a check.
8. Reject ⟹ "**at least one** restriction fails," never "all."
