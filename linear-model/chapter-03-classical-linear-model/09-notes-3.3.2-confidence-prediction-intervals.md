# 3.3.2 — Confidence Regions and Prediction Intervals

> **A confidence interval appears in every single past paper.** And the CI-vs-prediction-interval distinction is a classic conceptual question. Both are cheap marks once the formulas are automatic.

---

## 1. Confidence interval for a single coefficient ⭐

From $\dfrac{\hat\beta_j-\beta_j}{\widehat{\text{se}}(\hat\beta_j)}\sim t_{n-p}$, a $(1-\alpha)$ confidence interval is

$$\boxed{\;\hat\beta_j\ \pm\ t_{n-p}\!\left(1-\tfrac{\alpha}{2}\right)\cdot\widehat{\text{se}}(\hat\beta_j)\;}$$

| Confidence level | $\alpha$ | Quantile to look up |
|---|---|---|
| 90% | 0.10 | $t_{n-p}(0.95)$ |
| **95%** | 0.05 | $t_{n-p}(\mathbf{0.975})$ |
| **99%** | 0.01 | $t_{n-p}(\mathbf{0.995})$ |

> ⚠️ **The most common single error in this exam: using the 0.95 column for a 95% CI.** A two-sided interval splits $\alpha$ between two tails, so you need $1-\alpha/2$. Write "$1-\alpha/2$" next to the number every time.

### 🔑 The CI–test duality

$$\text{A } (1-\alpha) \text{ CI for } \beta_j \text{ contains } c \iff H_0:\beta_j=c \text{ is NOT rejected at level } \alpha$$

**In particular:** if the CI excludes 0, the coefficient is significant at level $\alpha$.

> 🔴 **Exam Summer 2025, Ex 1(k):** *"If a confidence interval for a slope coefficient $\beta_j$ does not contain zero, we **fail to reject** $H_0:\beta_j=0$."* → **FALSE.** Excluding zero means we **DO reject**. The statement has the logic backwards.
>
> 🔴 *WS 23/24, Block II(iii):* *"When the CI contains zero, we cannot reject the hypothesis that the coefficient is zero."* → **TRUE.** Same fact, stated correctly.
>
> **Memory hook:** *zero inside the net ⟹ zero is still a live possibility ⟹ don't reject. Zero outside the net ⟹ zero is ruled out ⟹ reject.*

### ⚠️ How to interpret a CI (markers watch for this)

> ✅ *"With 95% confidence, the interval $[0.507, 0.733]$ covers the true $\beta_1$."*
> ✅ *"If we repeated the sampling procedure many times, 95% of the intervals so constructed would contain the true $\beta_1$."*
> ❌ *"There is a 95% probability that $\beta_1$ lies in $[0.507, 0.733]$."*

$\beta_1$ is a **fixed unknown constant** — it either is or isn't in your computed interval. The randomness lives in the **interval**, not in $\beta_1$. Before you draw the sample, the interval has a 95% chance of covering; after you compute it, the probability is 0 or 1 and you just don't know which.

---

## 2. 📝 Worked example — Exam Summer 2025, Ex 3(b) [2 pts]

> *"Calculate a 99% confidence interval for $\beta_{\text{nox}}$. At the 1% significance level, would you reject $H_0:\beta_{\text{nox}}=-30$ against $H_1:\beta_{\text{nox}}\neq-30$?"*

**Given:** $\hat\beta_{\text{nox}}=-36.99122$, $\widehat{\text{se}}=5.25574$, $n-p=501$.

**Step 1 — quantile.** 99% CI ⟹ $\alpha=0.01$ ⟹ need $t_{501}(0.995)$. From the paper's Table 1, row 501, column 0.995: $\;t_{501}(0.995)=2.5857$.

**Step 2 — margin of error.**
$$2.5857\times5.25574=13.590$$

**Step 3 — interval.**
$$-36.99122\pm13.590 \;\Longrightarrow\; \boxed{[-50.581,\ -23.402]}$$

**Step 4 — the test.** Is $-30$ inside the interval?

$$-50.581 < -30 < -23.402 \qquad\checkmark \text{ yes}$$

> **Since the 99% confidence interval contains $-30$, we do not reject $H_0:\beta_{\text{nox}}=-30$ at the 1% significance level.**

*(Equivalent direct test: $t=\frac{-36.99122-(-30)}{5.25574}=\frac{-6.99122}{5.25574}=-1.330$, and $|-1.330|<2.5857$ ⟹ don't reject. Same answer, as it must be. Doing this as a check costs 20 seconds.)*

---

## 3. 📝 Worked example — Sheet 4, Ex 3(d)

95% CI for $\beta_1$ (age) in the wage model. $\hat\beta_1=0.62$, $\widehat{\text{se}}=0.0576$, $n-p=2993$.

With df this large the sheet supplies **normal** quantiles: $z_{0.975}=1.9608$.

$$0.62\pm1.9608\times0.0576=0.62\pm0.1129 \;\Longrightarrow\; \boxed{[0.507,\ 0.733]}$$

> The interval excludes 0, so $H_0:\beta_1=0$ is **rejected** at the 5% level — consistent with the t-test in Ex 3(c) ($t=10.76$). ✓
>
> *Interpretation:* with 95% confidence, each additional year of age is associated with an increase in expected hourly wage of between \$0.51 and \$0.73, holding education and health fixed.

---

## 4. Prediction: two different questions ⭐

You have a **new** covariate vector $\boldsymbol{x}_0$. There are **two** different things you might want, and confusing them is a classic exam trap.

### Question A: "What is the *average* $y$ for people like this?"

Target: $E(y_0)=\boldsymbol{x}_0'\boldsymbol\beta$ — a **fixed unknown number**.

$$\hat y_0=\boldsymbol{x}_0'\hat{\boldsymbol\beta},\qquad \text{Var}(\hat y_0)=\sigma^2\,\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0$$

$$\boxed{\;\text{CI for the mean: }\ \boldsymbol{x}_0'\hat{\boldsymbol\beta}\ \pm\ t_{n-p}(1-\tfrac{\alpha}{2})\cdot\hat\sigma\sqrt{\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0}\;}$$

### Question B: "What will *this particular person's* $y$ be?"

Target: $y_0=\boldsymbol{x}_0'\boldsymbol\beta+\varepsilon_0$ — a **random variable**, because it carries its own fresh error.

$$\text{Var}(y_0-\hat y_0)=\underbrace{\sigma^2}_{\text{new error }\varepsilon_0}+\underbrace{\sigma^2\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0}_{\text{uncertainty in }\hat{\boldsymbol\beta}}=\sigma^2\left(1+\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0\right)$$

$$\boxed{\;\text{Prediction interval: }\ \boldsymbol{x}_0'\hat{\boldsymbol\beta}\ \pm\ t_{n-p}(1-\tfrac{\alpha}{2})\cdot\hat\sigma\sqrt{1+\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0}\;}$$

### 🔑 The difference is exactly the "$1+$"

| | Confidence interval (mean) | Prediction interval (individual) |
|---|---|---|
| Target | $E(y_0)$ — fixed | $y_0$ — random |
| Under the root | $\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0$ | $\mathbf{1}+\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0$ |
| Sources of uncertainty | estimating $\boldsymbol\beta$ only | estimating $\boldsymbol\beta$ **plus** the new $\varepsilon_0$ |
| Width | narrower | **wider** — always |
| As $n\to\infty$ | shrinks to **zero** | shrinks to $\pm t\hat\sigma$, **never zero** |

> ### The one-sentence answer if asked "which is wider and why?"
>
> *"The prediction interval is wider, because it must account for two sources of uncertainty — the uncertainty in estimating $\boldsymbol\beta$ (shared with the confidence interval) **and** the irreducible random error $\varepsilon_0$ of the new observation. Even with infinite data the prediction interval cannot shrink below $\pm t\hat\sigma$, whereas the confidence interval for the mean shrinks to zero."*

**Intuition:** you can learn the *average* height of 40-year-olds to arbitrary precision with enough data. You can never predict *one specific* 40-year-old's height precisely, no matter how much data you have — because that person has their own randomness that no sample can tell you about.

---

## 5. 📝 Worked example — Sheet 4, Ex 3(e)

> *"Calculate the prediction interval for the wage of a 50-year-old man with an advanced degree but less than good health. Use $\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0=0.0035$."*

**Step 1 — build $\boldsymbol{x}_0$.** Model order: $(1,\ \text{age},\ \text{HS},\ \text{SC},\ \text{CG},\ \text{AD},\ \text{health.VG})$.

Age 50, Advanced Degree ⟹ AD dummy = 1, others 0. Health *less than good* ⟹ health.VG = 0.

$$\boldsymbol{x}_0=(1,\ 50,\ 0,\ 0,\ 0,\ 1,\ 0)'$$

**Step 2 — point prediction.**
$$\hat y_0=52.61+0.62(50)+62.63=52.61+31.00+62.63=\boxed{146.24}$$

**Step 3 — the standard error.** $\hat\sigma=35.724$ (from Sheet 3).

$$\hat\sigma\sqrt{1+0.0035}=35.724\times\sqrt{1.0035}=35.724\times1.001749=35.786$$

**Step 4 — assemble** with $z_{0.975}=1.9608$:

$$146.24\pm1.9608\times35.786=146.24\pm70.17$$

$$\boxed{[76.07,\ 216.41]}$$

**Compare the CI for the mean** at the same $\boldsymbol{x}_0$:

$$146.24\pm1.9608\times35.724\times\sqrt{0.0035}=146.24\pm4.14 \;\Longrightarrow\; [142.10,\ 150.38]$$

> **The prediction interval is ~17 times wider.** That's the whole lesson in one comparison: we know the *group average* wage quite precisely (±\$4), and an *individual's* wage barely at all (±\$70). The gap is $\sigma$ — genuine person-to-person variation that no amount of data removes.
>
> ⚠️ Notice how much this is driven by the "1": $\sqrt{1.0035}\approx1$ versus $\sqrt{0.0035}\approx0.059$.

---

## 6. Confidence intervals for a linear combination

For any fixed vector $\boldsymbol{a}$, since $\boldsymbol{a}'\hat{\boldsymbol\beta}\sim N(\boldsymbol{a}'\boldsymbol\beta,\ \sigma^2\boldsymbol{a}'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{a})$:

$$\boldsymbol{a}'\hat{\boldsymbol\beta}\pm t_{n-p}(1-\tfrac{\alpha}{2})\cdot\hat\sigma\sqrt{\boldsymbol{a}'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{a}}$$

**Useful for questions like "give a CI for the wage difference between Advanced Degree and HS Grad"** — set $\boldsymbol{a}$ to have $+1$ in the AD position and $-1$ in the HS position.

⚠️ This requires the **off-diagonal** elements too:
$$\text{Var}(\hat\beta_5-\hat\beta_2)=\text{Var}(\hat\beta_5)+\text{Var}(\hat\beta_2)-2\text{Cov}(\hat\beta_5,\hat\beta_2)$$
The covariance term is usually **not** zero. Don't just add the variances.

---

## 7. Simultaneous confidence regions (brief)

A joint $(1-\alpha)$ confidence region for the whole vector $\boldsymbol\beta$ is an **ellipsoid**:

$$\left\{\boldsymbol\beta:\ \frac{1}{p\hat\sigma^2}(\hat{\boldsymbol\beta}-\boldsymbol\beta)'\boldsymbol{X}'\boldsymbol{X}(\hat{\boldsymbol\beta}-\boldsymbol\beta)\leq F_{p,n-p}(1-\alpha)\right\}$$

**The point to know:** the joint region is **not** the rectangle formed by the individual CIs. Because $\hat\beta_i$ and $\hat\beta_j$ are correlated, the ellipse is tilted. So a pair of values can lie inside both individual intervals yet outside the joint region, and vice versa.

**Practical consequence:** testing $k$ coefficients one at a time at level $\alpha$ does **not** give an overall level $\alpha$ — that's the multiple-testing problem, and it's why the joint F-test exists rather than just running lots of t-tests.

---

## 8. Quantile lookup — the cheat table

| You want | Quantile | Note |
|---|---|---|
| 90% CI | $t(0.95)$ | |
| **95% CI** | $t(\mathbf{0.975})$ | most common |
| **99% CI** | $t(\mathbf{0.995})$ | Exam 2025 Ex 3(b) |
| Two-sided t-test, $\alpha=0.05$ | $t(0.975)$ | |
| Two-sided t-test, $\alpha=0.01$ | $t(0.995)$ | |
| One-sided t-test, $\alpha=0.05$ | $t(0.95)$ | |
| **F-test, $\alpha=0.05$** | $F(\mathbf{0.95})$ | **one-sided!** |
| F-test, $\alpha=0.01$ | $F(0.99)$ | |

> 🔑 **The rule:** $t$ and CI are **two-sided** ⟹ use $1-\alpha/2$. $F$ is **one-sided** ⟹ use $1-\alpha$.

Also: when $n-p$ is large (say $>200$), $t$ quantiles are essentially normal quantiles. The tables reflect this — the Sheet 4 table supplies $1.9608$, which is just $z_{0.975}$.

---

## 9. Key takeaways

1. $$\hat\beta_j\pm t_{n-p}(1-\tfrac{\alpha}{2})\cdot\widehat{\text{se}}(\hat\beta_j)$$ — **$1-\alpha/2$, not $1-\alpha$.**
2. **CI–test duality:** $c$ inside the CI ⟺ don't reject $H_0:\beta_j=c$. **Zero outside ⟹ significant.**
3. Interpret the CI as covering, not as a probability about $\beta_j$.
4. **CI for the mean:** $\hat\sigma\sqrt{\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0}$.
5. **Prediction interval:** $\hat\sigma\sqrt{\mathbf{1}+\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0}$.
6. **Prediction interval is always wider** — it carries the new observation's own error $\varepsilon_0$, which never vanishes.
7. Linear combinations need the **off-diagonal covariances** too.
8. **Two-sided ⟹ $1-\alpha/2$. F-test ⟹ $1-\alpha$.**
