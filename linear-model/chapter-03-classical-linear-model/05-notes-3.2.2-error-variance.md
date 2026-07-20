# 3.2.2 — Estimation of the Error Variance

> Short section, but it feeds **everything**: standard errors, t-tests, F-tests, confidence intervals, AIC, BIC. And it contains one distinction — **ML vs unbiased** — that the exercise sheets test directly.

---

## 1. What we're estimating

$\sigma^2=\text{Var}(\varepsilon_i)$ — the variance of the true errors, assumed constant across observations (A3).

We can't observe $\varepsilon_i$. We observe $\hat\varepsilon_i$. So we estimate $\sigma^2$ from the residuals.

---

## 2. The two estimators

### The unbiased (REML / "standard") estimator ⭐

$$\boxed{\;\hat\sigma^2=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n-p}=\frac{\sum_{i=1}^n\hat\varepsilon_i^2}{n-p}=\frac{\text{SSE}}{n-p}\;}$$

**This is what R reports** (as the square root), what standard errors use, and what t-tests and confidence intervals use.

$$E(\hat\sigma^2)=\sigma^2 \qquad\checkmark \text{ unbiased}$$

The square root $\hat\sigma$ is the **residual standard error**:

$$\hat\sigma=\sqrt{\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n-p}}$$

> In R output: `Residual standard error: 34 on 2983 degrees of freedom` — that's $\hat\sigma=34$ and $n-p=2983$.

### The maximum likelihood estimator

$$\boxed{\;\hat\sigma^2_{ML}=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n}\;}$$

$$E(\hat\sigma^2_{ML})=\frac{n-p}{n}\sigma^2 \qquad\Longrightarrow\qquad \textbf{biased downwards}$$

**This is the version used in AIC and BIC.** The Fahrmeir book is explicit about it:

> *"Note that the ML estimator $\hat\sigma^2=\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}/n$ is considered in AIC and not the usual unbiased variance estimator."*

### 🔴 Which one where — this table is the exam content

| Used for | Estimator | Denominator |
|---|---|---|
| Standard errors, $t$-tests, CIs, prediction intervals | $\hat\sigma^2$ | $n-p$ |
| $F$-test | $\hat\sigma^2$ | $n-p$ |
| Residual standard error in R output | $\hat\sigma$ | $n-p$ |
| Standardised residuals | $\hat\sigma$ | $n-p$ |
| **AIC and BIC** | $\hat\sigma^2_{ML}$ | **$n$** |

> ⚠️ **Sheet 5, Exercise 1(b)** asks you to compute AIC and BIC. Using $n-p$ instead of $n$ there gives the wrong answer. Since the two differ by a factor $n/(n-p)$, with $n=3000$ and $p=7$ the difference is small but real — and in a small-$n$ problem it's large.

---

## 3. Why $n-p$ and not $n$?

Three ways of seeing it. Know at least the first two.

### (a) Degrees of freedom — the countable argument

You started with $n$ observations. You **used up $p$ of them** estimating $\hat\beta_0,\dots,\hat\beta_k$. Only $n-p$ independent pieces of information remain to estimate the spread.

The clearest evidence: the $p$ normal equations $\boldsymbol{X}'\hat{\boldsymbol\varepsilon}=\boldsymbol{0}$ are $p$ **exact linear constraints** on the residuals. Given any $n-p$ of them, the remaining $p$ are determined. So there are genuinely only $n-p$ free residuals.

**The extreme case makes it vivid.** If $n=p$, the model fits every point exactly: $\hat{\boldsymbol\varepsilon}=\boldsymbol{0}$, $R^2=1$, and $\hat\sigma^2=0/0$ — undefined. And rightly so: with no leftover information you cannot say anything about noise. *(This is also the cleanest statement of overfitting.)*

**Familiar special case:** the sample variance $s^2=\frac{1}{n-1}\sum(y_i-\bar y)^2$ is exactly this with $p=1$ — the model $y_i=\beta_0+\varepsilon_i$, one parameter estimated ($\bar y$), so $n-1$.

### (b) The trace argument — the formal one

$$\hat{\boldsymbol\varepsilon}=(\boldsymbol{I}-\boldsymbol{H})\boldsymbol{y}=(\boldsymbol{I}-\boldsymbol{H})\boldsymbol\varepsilon$$

*(the second equality because $(\boldsymbol{I}-\boldsymbol{H})\boldsymbol{X}=\boldsymbol{0}$)*

$$E(\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon})=E(\boldsymbol\varepsilon'(\boldsymbol{I}-\boldsymbol{H})\boldsymbol\varepsilon)=\sigma^2\,\text{tr}(\boldsymbol{I}-\boldsymbol{H})=\sigma^2(n-p)$$

using $\text{tr}(\boldsymbol{H})=\text{rank}(\boldsymbol{H})=p$. Dividing by $n-p$ therefore gives exactly $\sigma^2$. ∎

### (c) Geometry

The residual vector is confined to the $(n-p)$-dimensional space orthogonal to the column space of $\boldsymbol{X}$. Its squared length is $\sigma^2$ times a $\chi^2$ variable with $n-p$ degrees of freedom:

$$\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{\sigma^2}=\frac{(n-p)\hat\sigma^2}{\sigma^2}\sim\chi^2_{n-p}$$

*(This requires normality, A6, and is what makes the $t$- and $F$-distributions appear in Section 3.3.)*

---

## 4. 📝 Worked example — Sheet 3, Exercise 2(b)

> *"Calculate the Restricted Maximum Likelihood estimation for the model variance $\sigma^2$. Use $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}=3819720$."*

**Setup.** The model is
$$\text{wage}=\beta_0+\beta_1\text{age}+\beta_2\text{HSGrad}+\beta_3\text{SomeCollege}+\beta_4\text{CollegeGrad}+\beta_5\text{AdvDegree}+\beta_6\text{health.VeryGood}+\varepsilon$$

so $p=7$ parameters, and the `Wage` dataset has $n=3000$.

**"Restricted maximum likelihood" (REML) here means the unbiased estimator** — the one that corrects for the degrees of freedom used in estimating $\boldsymbol\beta$:

$$\hat\sigma^2_{\text{REML}}=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n-p}=\frac{3819720}{3000-7}=\frac{3819720}{2993}=\boxed{1276.218}$$

$$\hat\sigma=\sqrt{1276.218}=\boxed{35.724}$$

**Compare the plain ML estimate:**
$$\hat\sigma^2_{ML}=\frac{3819720}{3000}=1273.240$$

Slightly smaller, as expected from the downward bias. Here the difference is 0.23% — negligible for interpretation, but it matters for AIC/BIC and it matters conceptually.

> **Naming note.** "ML" divides by $n$. "REML"/"restricted ML"/"the usual unbiased estimator" divides by $n-p$. If a question says REML, use $n-p$. If it says ML — or if you're computing AIC/BIC — use $n$.

---

## 5. What $\hat\sigma$ feeds into

```
                    σ̂² = ε̂'ε̂/(n−p)
                          │
        ┌─────────────────┼─────────────────┐
        │                 │                 │
   se(β̂ⱼ) = σ̂√[(X'X)⁻¹]ⱼⱼ   F-statistic   standardised
        │                 │              residuals
        │                 │              rᵢ = ε̂ᵢ/(σ̂√(1−hᵢᵢ))
   ┌────┴────┐            │
   │         │            │
t-test    CI for βⱼ    prediction
   │         │          intervals
   └────┬────┘
        │
   test decisions
```

**Practical consequence: an error in $\hat\sigma$ propagates into every inferential quantity in the paper.** If you compute it once at the start and it's wrong, you lose marks in three separate exercises. Compute it carefully, and sanity-check it: $\hat\sigma$ should be on the same scale as the spread of $y$.

> **Quick sanity check with R output.** If the output says `Residual standard error: 34`, and your computed $\hat\sigma=35.7$, you're in the right neighbourhood. If you get 1276, you forgot the square root.

---

## 6. 📝 Worked example — Exam Summer 2025, Ex 3(a), value [[D]]

> R output: `Residual standard error: [[D]] on 501 degrees of freedom`, and `sum(lm$residuals^2) = 31682.02`.

$$\hat\sigma=\sqrt{\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n-p}}=\sqrt{\frac{31682.02}{501}}=\sqrt{63.238}=\boxed{7.952}$$

**Note:** the output hands you $n-p=501$ directly. You don't need to work out $n$ or $p$ separately — though as a check: the model is `medv ~ crim + nox + dis + rad`, so $p=5$, and $n=501+5=506$ (the Boston housing dataset — correct ✓).

---

## 7. Key takeaways

1. **Unbiased:** $\hat\sigma^2=\dfrac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n-p}$ — used for standard errors, $t$, $F$, CIs, standardised residuals. **This is R's "residual standard error" (squared).**
2. **ML:** $\hat\sigma^2_{ML}=\dfrac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n}$ — biased downward, **used in AIC and BIC**.
3. **Why $n-p$:** you spent $p$ degrees of freedom estimating $\boldsymbol\beta$; the normal equations impose $p$ exact constraints on the residuals; formally $\text{tr}(\boldsymbol{I}-\boldsymbol{H})=n-p$.
4. If $n=p$ the fit is perfect and $\sigma^2$ is inestimable — the cleanest picture of overfitting.
5. $\dfrac{(n-p)\hat\sigma^2}{\sigma^2}\sim\chi^2_{n-p}$ under normality — the source of the $t$ and $F$ distributions.
6. **$\hat\sigma$ feeds everything.** Get it right once; check it against the R output; don't forget the square root.
