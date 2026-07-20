# 1.3 — Notational Remarks

> **This is the highest-value 15 minutes in Chapter 1.** Every formula in Chapter 3 is assembled from these symbols. Write this table out by hand once. Seriously.

---

## 1. The core symbols

| Symbol | Reads as | Dimension | Meaning |
|---|---|---|---|
| $n$ | "n" | scalar | number of observations |
| $k$ | "k" | scalar | number of **covariates** |
| $p$ | "p" | scalar | number of **parameters** = $k+1$ (covariates + intercept) |
| $i$ | index | — | observation index, $i = 1,\dots,n$ |
| $j$ | index | — | covariate index, $j = 1,\dots,k$ |
| $y_i$ | "y-i" | scalar | response for observation $i$ |
| $x_{ij}$ | "x-i-j" | scalar | value of covariate $j$ for observation $i$ |
| $\boldsymbol{x}_i$ | "x-i" (bold) | $p \times 1$ | **row** of the design matrix for observation $i$: $(1, x_{i1},\dots,x_{ik})'$ |
| $\boldsymbol{y}$ | "y" (bold) | $n \times 1$ | vector of all responses |
| $\boldsymbol{X}$ | "X" | $n \times p$ | **design matrix** |
| $\boldsymbol{\beta}$ | "beta" | $p \times 1$ | true parameter vector $(\beta_0,\dots,\beta_k)'$ |
| $\boldsymbol{\varepsilon}$ | "epsilon" | $n \times 1$ | true error vector |
| $\hat{\boldsymbol{\beta}}$ | "beta hat" | $p \times 1$ | estimated parameters |
| $\hat{\boldsymbol{y}}$ | "y hat" | $n \times 1$ | fitted values $\boldsymbol{X}\hat{\boldsymbol\beta}$ |
| $\hat{\boldsymbol{\varepsilon}}$ | "epsilon hat" | $n \times 1$ | residuals $\boldsymbol{y} - \hat{\boldsymbol{y}}$ |
| $\sigma^2$ | "sigma squared" | scalar | true error variance |
| $\hat\sigma^2$ | "sigma hat squared" | scalar | estimated error variance |

---

## 2. 🚨 THE `p` TRAP — read this three times

**The Fahrmeir textbook defines $p = k+1$.** Quote from the book (Section 3.1):

> *"If we combine the covariates and the unknown parameters each into $p = (k+1)$ dimensional vectors…"*

So in **book notation**:

| Quantity | Book formula |
|---|---|
| Number of parameters | $p$ |
| $\text{rank}(\boldsymbol{X}) = \text{rank}(\boldsymbol{X}'\boldsymbol{X})$ | $p$ |
| Residual degrees of freedom | $n - p$ |
| t-test distribution | $t_j \sim t_{n-p}$ |
| F-test distribution | $F \sim F_{r,\,n-p}$ |
| Corrected $R^2$ | $\bar{R}^2 = 1 - \frac{n-1}{n-p}(1-R^2)$ |

**But some of your past exam papers use $p$ = number of covariates**, and then correctly write $t_j \sim t_{n-p-1}$. Evidence from the answer keys you have:

- *Linear_model_exam_sheet*, Block I(iii): "under $H_0$, $t_j \sim t_{n-p-1}$" → marked **TRUE** ⟹ there, $p$ = covariates.
- *Linear_model_exam_sheet*, Block I(i): "$\text{rank}(X'X) = p$, with the number of variables $p$" → marked **FALSE** ⟹ consistent, since rank is $p+1$ when $p$ = covariates.
- *RCLMWS2223*, Block I(i): "$\text{rk}(X'X) = p$" (no gloss) → **TRUE** in book notation.

**Both papers are internally consistent. They just disagree with each other.**

### The bulletproof habit

Never memorise the string "n−p". Memorise the **sentence**:

> **Residual df = $n$ − (number of $\beta$'s you estimated, counting the intercept).**

Then, in the exam:
1. Count the betas. Call it $B$.
2. Residual df is $n - B$.
3. Look at how *the paper in front of you* defines its symbols, and write your answer in *their* letters.

With $k$ covariates and an intercept, $B = k+1$, so residual df $= n-k-1$. This is always true and never ambiguous. Write $n-k-1$ if you're unsure — it's unambiguous in both conventions.

---

## 3. The design matrix $\boldsymbol{X}$ — build it once, properly

For $n$ observations and $k$ covariates plus an intercept:

$$
\boldsymbol{X} = \begin{pmatrix}
1 & x_{11} & x_{12} & \cdots & x_{1k} \\
1 & x_{21} & x_{22} & \cdots & x_{2k} \\
\vdots & \vdots & \vdots & \ddots & \vdots \\
1 & x_{n1} & x_{n2} & \cdots & x_{nk}
\end{pmatrix} \in \mathbb{R}^{n \times p}, \quad p = k+1
$$

**Read it as:** one **row per observation**, one **column per parameter**.

- The **first column is all ones** — this is the intercept. It is not decoration; it is the covariate whose value is always 1, so its coefficient $\beta_0$ is the baseline.
- Column $j+1$ holds covariate $j$ across all observations.

### Model in matrix form

$$\boldsymbol{y} = \boldsymbol{X}\boldsymbol{\beta} + \boldsymbol{\varepsilon}$$

$$
\underbrace{\begin{pmatrix} y_1 \\ y_2 \\ \vdots \\ y_n\end{pmatrix}}_{n\times 1}
=
\underbrace{\begin{pmatrix} 1 & x_{11} & \cdots & x_{1k} \\ 1 & x_{21} & \cdots & x_{2k} \\ \vdots & & & \vdots \\ 1 & x_{n1} & \cdots & x_{nk}\end{pmatrix}}_{n \times p}
\underbrace{\begin{pmatrix}\beta_0 \\ \beta_1 \\ \vdots \\ \beta_k\end{pmatrix}}_{p \times 1}
+
\underbrace{\begin{pmatrix}\varepsilon_1 \\ \varepsilon_2 \\ \vdots \\ \varepsilon_n\end{pmatrix}}_{n\times 1}
$$

> **This exact display was worth 1 point in the *Linear_model_exam_sheet* paper, Exercise 2(a):** *"Express the model in matrix form, clearly specifying terms Y, X, β and ε."* Practise writing it in under 60 seconds, including the dimension labels. Free mark.

### Dimension sanity check (do this reflexively)

$$(n\times p)(p\times 1) = (n \times 1) \;\checkmark$$

Inner dimensions must match; outer dimensions give the result. If your dimensions don't work, your formula is wrong — this catches most matrix errors before you waste time.

Quick reference for the formulas you'll meet:

| Expression | Dimensions | Result |
|---|---|---|
| $\boldsymbol{X}'\boldsymbol{X}$ | $(p\times n)(n \times p)$ | $p \times p$ — square, symmetric, invertible iff full rank |
| $(\boldsymbol{X}'\boldsymbol{X})^{-1}$ | — | $p \times p$ |
| $\boldsymbol{X}'\boldsymbol{y}$ | $(p\times n)(n\times 1)$ | $p \times 1$ |
| $\hat{\boldsymbol\beta} = (\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$ | $(p\times p)(p\times 1)$ | $p\times 1$ ✓ |
| $\boldsymbol{H} = \boldsymbol{X}(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'$ | $(n\times p)(p\times p)(p \times n)$ | $n \times n$ — the **hat matrix** |
| $\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0$ | $(1\times p)(p\times p)(p\times 1)$ | **scalar** |

That last one shows up in prediction intervals (Sheet 4, Ex 3(e)) and looks intimidating until you notice it's just a number.

---

## 4. Transpose, and the two ways to write a sum

$$\boldsymbol{a}'\boldsymbol{b} = \sum_{i=1}^n a_i b_i \qquad\text{(a scalar)}$$

Therefore:

$$\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon} = \sum_{i=1}^{n}\hat\varepsilon_i^2 = \text{SSE} = \text{RSS}$$

**Learn to read $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}$ instantly as "sum of squared residuals."** Your exercise sheets hand you this quantity as a raw number ("Use $\hat\varepsilon'\hat\varepsilon = 3819720$") and expect you to plug it into AIC, BIC, $\hat\sigma^2$ and the F-statistic without hesitating.

Naming warning — the same quantity has several names in the wild:

| Name | Abbreviation | Formula |
|---|---|---|
| Sum of Squared Errors / Residual Sum of Squares | **SSE**, RSS | $\sum \hat\varepsilon_i^2 = \hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}$ |
| Explained / Regression Sum of Squares | **SSR**, ESS | $\sum(\hat y_i - \bar y)^2$ |
| Total Sum of Squares | **SST**, TSS | $\sum(y_i - \bar y)^2$ |

⚠️ **SSR is ambiguous in the literature** — some texts use it for "Sum of Squares *Regression*", others for "Sum of Squares *Residual*". **Never write SSR on the exam.** Write SSE for residuals and "explained SS" for the other, or just write the sum out. The decomposition is:

$$\text{SST} = \text{SSE} + \text{explained SS}, \qquad R^2 = 1 - \frac{\text{SSE}}{\text{SST}} = \frac{\text{explained SS}}{\text{SST}}$$

---

## 5. Greek letters and the hat convention

| Greek | Name | Used for |
|---|---|---|
| $\beta$ | beta | regression coefficients |
| $\varepsilon$ | epsilon | error term |
| $\sigma$ | sigma | standard deviation ($\sigma^2$ = variance) |
| $\mu$ | mu | mean / expected value |
| $\alpha$ | alpha | significance level (0.05, 0.01) |
| $\lambda$ | lambda | penalty parameter (ridge, lasso — Ch 4) |
| $\eta$ | eta | linear predictor $\eta = \boldsymbol{x}'\boldsymbol\beta$ |
| $\pi$ | pi | a probability, e.g. $P(y=1)$ in the logit model |

**The hat rule, stated once and for all:**

> **Greek letter without a hat = the true, unknown, fixed quantity in the population.**
> **Greek letter with a hat = a number computed from your sample, which is random and would change with a new sample.**

This distinction *is* statistics. $\beta_1$ is a fixed unknown constant. $\hat\beta_1$ is a random variable — it has an expectation ($E(\hat\beta_1) = \beta_1$, i.e. unbiasedness) and a variance (which Gauss–Markov says is as small as possible among linear unbiased estimators). Everything in Chapter 3.2.3 and 3.3 is about the *distribution of the hatted things*.

If a question says "show the estimator is unbiased," it's asking you to show $E(\hat\beta) = \beta$ — the hat disappears when you take the expectation. That's the whole shape of the proof.

---

## 6. Expectation, variance, covariance — the four rules you actually use

For constants $a, b$ and random variables $X, Y$:

$$E(aX + b) = aE(X) + b$$
$$\text{Var}(aX + b) = a^2\text{Var}(X)$$
$$\text{Var}(X+Y) = \text{Var}(X) + \text{Var}(Y) + 2\text{Cov}(X,Y)$$
$$\text{Cov}(X,Y) = E(XY) - E(X)E(Y)$$

Matrix versions (used constantly in 3.2.3):

$$E(\boldsymbol{A}\boldsymbol{z}) = \boldsymbol{A}E(\boldsymbol{z}) \qquad \text{Cov}(\boldsymbol{A}\boldsymbol{z}) = \boldsymbol{A}\,\text{Cov}(\boldsymbol{z})\,\boldsymbol{A}'$$

That second one — **$A$ on the left, $A$ transposed on the right** — is the engine behind

$$\text{Cov}(\hat{\boldsymbol\beta}) = \sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}$$

which you will use for every standard error, every t-test and every confidence interval in this course. If you learn one matrix identity, learn that one.

---

## Key takeaways

1. **$n$** observations, **$k$** covariates, **$p = k+1$** parameters *in the book's notation* — but always fall back on *"count the betas."*
2. $\boldsymbol{X}$ is $n \times p$: **one row per observation, one column per parameter, first column all ones.**
3. Check dimensions on every matrix expression. It catches errors for free.
4. $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon} = \sum\hat\varepsilon_i^2 = $ SSE. Recognise it instantly.
5. **Hat = computed from data and random. No hat = true and unknown.**
6. $\text{Cov}(\boldsymbol{Az}) = \boldsymbol{A}\,\text{Cov}(\boldsymbol{z})\,\boldsymbol{A}'$ — the identity behind everything in Chapter 3.
