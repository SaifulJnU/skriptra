# 3.2.3 — Properties of the Estimators: Gauss–Markov and BLUE

> **WS 23/24, Exercise 2(a) is worth 4 points and says: "Briefly describe the main contents of the Gaussian Markov-Theorem and the assumptions."** The marking key: *"1 point for every assumption."*
>
> This is a **recitation question**. Four marks for a list you can memorise in twenty minutes. Do not leave them on the table.

---

## 1. The workhorse identity

Everything in this section flows from one substitution. Start with $\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$ and put $\boldsymbol{y}=\boldsymbol{X}\boldsymbol\beta+\boldsymbol\varepsilon$:

$$\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'(\boldsymbol{X}\boldsymbol\beta+\boldsymbol\varepsilon)=\underbrace{(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{X}}_{\boldsymbol{I}}\boldsymbol\beta+(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol\varepsilon$$

$$\boxed{\;\hat{\boldsymbol\beta}=\boldsymbol\beta+(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol\varepsilon\;}$$

**Read it:** *the estimate equals the truth plus a random term driven entirely by the errors.* Everything below is a statement about that random term.

---

## 2. The four properties

### (1) Linearity

$$\hat{\boldsymbol\beta}=\boldsymbol{A}\boldsymbol{y},\qquad \boldsymbol{A}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}' \text{ (a fixed, non-random matrix)}$$

$\hat{\boldsymbol\beta}$ is a **linear function of the response vector**. This is the **L** in BLUE — and it's a *restriction*, not a virtue. It defines the class of competitors we're comparing against.

**Requires:** A5 (full rank), $\boldsymbol{X}$ non-stochastic.

### (2) Unbiasedness

$$E(\hat{\boldsymbol\beta})=\boldsymbol\beta+(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\underbrace{E(\boldsymbol\varepsilon)}_{=\boldsymbol{0}}=\boldsymbol\beta$$

$$\boxed{\;E(\hat{\boldsymbol\beta})=\boldsymbol\beta\;}$$

**Meaning:** over repeated samples, the estimates centre on the truth. Not that any single estimate is correct — that the *procedure* has no systematic error.

**Requires only:** A1 (correct specification), A2 ($E(\boldsymbol\varepsilon)=\boldsymbol{0}$), A5 (full rank).

> 🔑 **Not required:** homoscedasticity, independence, or normality. This is why "OLS is still unbiased under heteroscedasticity" is true, and it's the most useful single fact in the assumptions section.

### (3) Covariance matrix

Apply $\text{Cov}(\boldsymbol{Az})=\boldsymbol{A}\,\text{Cov}(\boldsymbol{z})\,\boldsymbol{A}'$ to the workhorse identity ($\boldsymbol\beta$ is constant, so it drops out):

$$\text{Cov}(\hat{\boldsymbol\beta})=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\cdot\underbrace{\text{Cov}(\boldsymbol\varepsilon)}_{\sigma^2\boldsymbol{I}}\cdot\boldsymbol{X}(\boldsymbol{X}'\boldsymbol{X})^{-1}$$
$$=\sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}\underbrace{\boldsymbol{X}'\boldsymbol{X}(\boldsymbol{X}'\boldsymbol{X})^{-1}}_{\boldsymbol{I}}$$

$$\boxed{\;\text{Cov}(\hat{\boldsymbol\beta})=\sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}\;}$$

**Requires:** A3 **and** A4 (i.e. $\text{Cov}(\boldsymbol\varepsilon)=\sigma^2\boldsymbol{I}$). If either fails, this formula is **wrong**, and so is every standard error computed from it.

**Reading the matrix:**

$$\text{Var}(\hat\beta_j)=\sigma^2\left[(\boldsymbol{X}'\boldsymbol{X})^{-1}\right]_{jj}\qquad \text{Cov}(\hat\beta_i,\hat\beta_j)=\sigma^2\left[(\boldsymbol{X}'\boldsymbol{X})^{-1}\right]_{ij}$$

$$\widehat{\text{se}}(\hat\beta_j)=\hat\sigma\sqrt{\left[(\boldsymbol{X}'\boldsymbol{X})^{-1}\right]_{jj}}$$

### (4) Distribution — under normality only

Adding A6 ($\boldsymbol\varepsilon\sim N(\boldsymbol{0},\sigma^2\boldsymbol{I})$), since $\hat{\boldsymbol\beta}$ is a **linear** function of a normal vector:

$$\boxed{\;\hat{\boldsymbol\beta}\sim N\!\left(\boldsymbol\beta,\ \sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}\right)\;}$$

**Exactly**, for any $n$. This is what makes the $t$- and $F$-tests of Section 3.3 exact rather than approximate.

*(Without A6, the Central Limit Theorem still gives asymptotic normality for large $n$.)*

---

## 3. ⭐ The Gauss–Markov Theorem

### Statement

> Under assumptions **A1, A2, A3, A4, A5** — that is, in the linear model $\boldsymbol{y}=\boldsymbol{X}\boldsymbol\beta+\boldsymbol\varepsilon$ with $E(\boldsymbol\varepsilon)=\boldsymbol{0}$, $\text{Cov}(\boldsymbol\varepsilon)=\sigma^2\boldsymbol{I}_n$, and $\boldsymbol{X}$ non-stochastic of full column rank $p$ — the OLS estimator
> $$\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$$
> is the **Best Linear Unbiased Estimator (BLUE)** of $\boldsymbol\beta$.
>
> Formally: for any other estimator $\tilde{\boldsymbol\beta}=\boldsymbol{B}\boldsymbol{y}$ that is linear in $\boldsymbol{y}$ and unbiased for $\boldsymbol\beta$,
> $$\text{Cov}(\tilde{\boldsymbol\beta})-\text{Cov}(\hat{\boldsymbol\beta}) \text{ is positive semi-definite}$$
> equivalently, $\text{Var}(\boldsymbol{a}'\hat{\boldsymbol\beta})\leq\text{Var}(\boldsymbol{a}'\tilde{\boldsymbol\beta})$ for every fixed vector $\boldsymbol{a}$ — in particular $\text{Var}(\hat\beta_j)\leq\text{Var}(\tilde\beta_j)$ for each $j$.

### 🔑 BLUE, one letter at a time

| Letter | Word | Means | Watch out |
|---|---|---|---|
| **B** | **Best** | **minimum variance** | "Best" is *only* about variance, and *only* within the L∩U class |
| **L** | **Linear** | $\hat{\boldsymbol\beta}=\boldsymbol{A}\boldsymbol{y}$ | a **restriction on competitors**, not a merit |
| **U** | **Unbiased** | $E(\hat{\boldsymbol\beta})=\boldsymbol\beta$ | also a restriction |
| **E** | **Estimator** | a rule, not a number | it's the *procedure* that's optimal |

> 🔴 *Exam Summer 2025, Ex 1(e):* "A BLUE is 'best' in the sense that there is no other **linear unbiased** estimator with a lower variance." → **TRUE.** The statement correctly includes both qualifiers.
>
> If a statement drops either qualifier — e.g. "no other estimator has lower variance" — it becomes **FALSE**. There are plenty of **biased** estimators with lower variance (ridge, lasso, James–Stein), and plenty of **non-linear** ones. That's the whole bias–variance tradeoff of Section 3.4.1.

### 🔑 What Gauss–Markov does **not** require

**Normality is not needed.** This is the most commonly examined subtlety in the whole section.

> 🔴 *WS 23/24, Block I(ii):* "The least squares estimator is BLUE **if and only if** the error term is expected to be zero and has constant variance (homoskedasticity)." → **FALSE** (per the key).
>
> **Why:** the "if and only if" is wrong, and the assumption list is incomplete. BLUE also requires **uncorrelated errors** ($\text{Cov}(\varepsilon_i,\varepsilon_j)=0$), **correct specification**, and **full column rank**. Zero mean plus constant variance alone is not sufficient.
>
> **Lesson:** in T/F questions about Gauss–Markov, check the assumption list for *completeness*, and be suspicious of "if and only if."

### Sketch of the proof (know the shape, not the algebra)

Let $\tilde{\boldsymbol\beta}=\boldsymbol{B}\boldsymbol{y}$ be any linear estimator. Write $\boldsymbol{B}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'+\boldsymbol{D}$ for some $\boldsymbol{D}$.

- **Unbiasedness forces $\boldsymbol{D}\boldsymbol{X}=\boldsymbol{0}$.** *(Since $E(\tilde{\boldsymbol\beta})=\boldsymbol{B}\boldsymbol{X}\boldsymbol\beta=\boldsymbol\beta$ for all $\boldsymbol\beta$ requires $\boldsymbol{B}\boldsymbol{X}=\boldsymbol{I}$.)*
- Then the cross-terms in the covariance vanish, and

$$\text{Cov}(\tilde{\boldsymbol\beta})=\sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}+\sigma^2\boldsymbol{D}\boldsymbol{D}'=\text{Cov}(\hat{\boldsymbol\beta})+\sigma^2\boldsymbol{D}\boldsymbol{D}'$$

$\boldsymbol{D}\boldsymbol{D}'$ is always positive semi-definite, so any deviation from OLS **adds** variance. Equality holds only when $\boldsymbol{D}=\boldsymbol{0}$, i.e. $\tilde{\boldsymbol\beta}=\hat{\boldsymbol\beta}$. ∎

*(You will not be asked to reproduce this. Knowing that the proof is "any competitor's covariance = OLS's covariance + a PSD term" is enough, and mentioning it shows genuine understanding.)*

---

## 4. ✍️ Full-mark answer [4 pts]

> **The Gauss–Markov theorem.**
>
> *Setting:* the linear model $\boldsymbol{y}=\boldsymbol{X}\boldsymbol\beta+\boldsymbol\varepsilon$ with $\boldsymbol{y}\in\mathbb{R}^n$, $\boldsymbol{X}\in\mathbb{R}^{n\times p}$, $\boldsymbol\beta\in\mathbb{R}^p$.
>
> *Assumptions:*
> 1. **Correct specification / linearity in the parameters:** $E(\boldsymbol{y}\mid\boldsymbol{X})=\boldsymbol{X}\boldsymbol\beta$.
> 2. **Zero-mean errors:** $E(\varepsilon_i)=0$ for all $i$.
> 3. **Homoscedasticity:** $\text{Var}(\varepsilon_i)=\sigma^2$, the same for all $i$.
> 4. **No autocorrelation:** $\text{Cov}(\varepsilon_i,\varepsilon_j)=0$ for $i\neq j$.
>    *(3 and 4 together: $\text{Cov}(\boldsymbol\varepsilon)=\sigma^2\boldsymbol{I}_n$.)*
> 5. **Full column rank and non-stochastic design:** $\text{rank}(\boldsymbol{X})=p$, so $(\boldsymbol{X}'\boldsymbol{X})^{-1}$ exists.
>
> *Conclusion:* the OLS estimator $\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$ is the **Best Linear Unbiased Estimator** of $\boldsymbol\beta$: among all estimators that are linear in $\boldsymbol{y}$ and unbiased for $\boldsymbol\beta$, it has the smallest variance. Formally, for any competing linear unbiased $\tilde{\boldsymbol\beta}$, the difference $\text{Cov}(\tilde{\boldsymbol\beta})-\text{Cov}(\hat{\boldsymbol\beta})$ is positive semi-definite, so $\text{Var}(\hat\beta_j)\leq\text{Var}(\tilde\beta_j)$ for every $j$. Moreover $\text{Cov}(\hat{\boldsymbol\beta})=\sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}$ and $\hat\sigma^2=\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}/(n-p)$ is unbiased for $\sigma^2$.
>
> **Remark.** Normality of the errors is **not** required for Gauss–Markov. It is needed only for the *exactness* of the $t$- and $F$-tests and confidence intervals; without it those results hold only asymptotically. Conversely, "best" is restricted to the linear-unbiased class: **biased** estimators (e.g. ridge) or **non-linear** estimators can have smaller mean squared error.

**That final remark is what separates 3 marks from 4.** Include it.

---

## 5. What breaks when assumptions fail

| Violated | Unbiased? | BLUE? | $\text{Cov}(\hat{\boldsymbol\beta})=\sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}$? | Exact tests? |
|---|---|---|---|---|
| A1 linearity | ❌ | ❌ | ❌ | ❌ |
| A3 homoscedasticity | ✅ | ❌ | ❌ | ❌ |
| A4 independence | ✅ | ❌ | ❌ | ❌ |
| A5 full rank | *undefined* | — | — | — |
| A6 normality | ✅ | ✅ | ✅ | ⚠️ asymptotic only |

**The template sentence for A3/A4 violations** (reusable, worth marks in at least two papers):

> *The OLS estimator remains **unbiased** and **consistent**, since unbiasedness requires only correct specification, zero-mean errors and full rank. However it is no longer **efficient** — the Gauss–Markov theorem no longer applies, so OLS is not BLUE, and a generalised/weighted least squares estimator has smaller variance. The usual covariance formula $\sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}$ is also incorrect, so standard errors are biased and the resulting $t$-tests, $F$-tests and confidence intervals are invalid.*

---

## 6. OLS and maximum likelihood

Under A6 the log-likelihood is

$$\ell(\boldsymbol\beta,\sigma^2)=-\frac{n}{2}\log(2\pi\sigma^2)-\frac{1}{2\sigma^2}(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)$$

$\boldsymbol\beta$ appears **only** in the final quadratic form, with a negative coefficient. Maximising over $\boldsymbol\beta$ is therefore identical to **minimising** $(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)$ — the least-squares criterion.

$$\boxed{\;\hat{\boldsymbol\beta}_{ML}=\hat{\boldsymbol\beta}_{LS}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}\;}$$

But the variance estimators **differ**:

$$\hat\sigma^2_{ML}=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n}\quad\text{(biased)} \qquad\qquad \hat\sigma^2_{LS/REML}=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n-p}\quad\text{(unbiased)}$$

> 🔴 *Exam Summer 2025, Ex 1(l):* "The OLS estimator is equivalent to the ML estimator under iid normal errors." → **TRUE** (for $\boldsymbol\beta$).
>
> 💰 *Sheet 3, Ex 2(a):* "Assuming normally distributed errors, what is the ML estimate for $\boldsymbol\beta$?" → **Identical to the least-squares estimate**, $\hat{\boldsymbol\beta}'_{ML}=(52.61, 0.62, 11.01, 23.16, 37.97, 62.63, 9.13)$. The question is testing whether you know they coincide.

**Bonus property under normality:** the ML estimator attains the Cramér–Rao lower bound, so $\hat{\boldsymbol\beta}$ is not merely BLUE — it is **efficient among all unbiased estimators**, linear or not. Normality upgrades "best linear unbiased" to "best unbiased."

---

## 7. Consistency

$$\hat{\boldsymbol\beta}\overset{p}{\longrightarrow}\boldsymbol\beta \quad\text{as } n\to\infty$$

provided $\frac{1}{n}\boldsymbol{X}'\boldsymbol{X}$ converges to a positive definite matrix. Since $\text{Cov}(\hat{\boldsymbol\beta})=\sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}\to\boldsymbol{0}$, the estimator concentrates on the truth.

**Consistency survives heteroscedasticity and autocorrelation.** It does **not** survive misspecification.

| Property | Nature |
|---|---|
| **Unbiasedness** | finite-sample: correct *on average* at any $n$ |
| **Consistency** | asymptotic: converges to the truth as $n\to\infty$ |

They're independent properties — an estimator can be biased but consistent (e.g. $\hat\sigma^2_{ML}$), or unbiased but inconsistent.

---

## 8. Key takeaways

1. **Workhorse:** $\hat{\boldsymbol\beta}=\boldsymbol\beta+(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol\varepsilon$. Everything follows.
2. $E(\hat{\boldsymbol\beta})=\boldsymbol\beta$ needs **only** A1, A2, A5 — not homoscedasticity, independence or normality.
3. $\text{Cov}(\hat{\boldsymbol\beta})=\sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}$ needs A3 **and** A4.
4. $\hat{\boldsymbol\beta}\sim N(\boldsymbol\beta,\sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1})$ needs A6, and makes tests **exact**.
5. **Gauss–Markov: A1–A5 ⟹ OLS is BLUE.** Five assumptions, one point each.
6. **BLUE = minimum variance among LINEAR and UNBIASED estimators.** Both qualifiers matter; drop either and the claim is false.
7. **Normality is NOT a Gauss–Markov assumption.** Say so explicitly — it's the 4th mark.
8. **A3/A4 violated ⟹ unbiased but inefficient, standard errors invalid.** Memorise the template sentence.
9. Under normality, $\hat{\boldsymbol\beta}_{ML}=\hat{\boldsymbol\beta}_{LS}$, but $\hat\sigma^2_{ML}\neq\hat\sigma^2_{LS}$.
