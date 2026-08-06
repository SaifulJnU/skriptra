# 30 — TRUE / FALSE DRILL

**72 statements · 18 blocks of 4 · answers and reasons at the bottom**

> **How to use this.** Answer a whole block before scrolling. Write **T** or **F** *and one clause of reason*. If you can't write the reason, mark it ⚠️ even when the verdict is right — a right answer for the wrong reason will fail on the next variant.
>
> **Never leave one blank in the real exam.** Negative marking is floored at zero **per block**, so a guess has positive expected value and no downside. If you are sure of 3 of 4 and guessing the last, answer all 4.
>
> **And read every statement to its final word.** The most common construction in your past papers is a *true first clause followed by a false second clause*.

Notation: $k$ covariates, $p = k+1$ parameters, $n$ observations.

---

## Block 1 — Model and design matrix

1. In $\boldsymbol{y}=\boldsymbol{X}\boldsymbol\beta+\boldsymbol\varepsilon$ with an intercept and $k$ covariates, $\boldsymbol{X}$ has $p=k+1$ columns.
2. $\text{rank}(\boldsymbol{X}'\boldsymbol{X})=n$.
3. When the model contains an intercept, the first column of $\boldsymbol{X}$ is a column of ones.
4. Adding a further column to $\boldsymbol{X}$ can reduce the rank of $\boldsymbol{X}$.

## Block 2 — Degrees of freedom

5. In a model with $k$ covariates and an intercept, the residual degrees of freedom are $n-k-1$.
6. For $H_0:\boldsymbol{C}\boldsymbol\beta=\boldsymbol{d}$ with $\boldsymbol{C}$ of dimension $r\times p$, the test statistic follows $F_{p+1,\,n-p}$.
7. The denominator degrees of freedom of the F statistic come from the **restricted** model.
8. For the overall model test, $F\sim F_{k,\,n-k-1}$.

## Block 3 — Counting restrictions

9. $H_0:\beta_1=-\beta_2+\beta_3$ imposes $r=3$ restrictions.
10. $\boldsymbol{C}$ has one row per restriction and one column per parameter, including $\beta_0$.
11. Testing $\beta_1=\beta_2$ **and** $\beta_3=0$ jointly gives $r=2$.
12. A restriction must be rearranged so all parameters sit on one side before its row of $\boldsymbol{C}$ can be read off.

## Block 4 — OLS

13. $\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$ minimises the sum of squared residuals.
14. The OLS estimate is unique as long as the error variance is constant.
15. Minimising the sum of **absolute** deviations yields the same estimates as OLS.
16. That $\partial^2S/\partial\boldsymbol\beta\partial\boldsymbol\beta'=2\boldsymbol{X}'\boldsymbol{X}$ is positive definite confirms the stationary point is a minimum.

## Block 5 — Unbiasedness

17. $E(\hat{\boldsymbol\beta})=\boldsymbol\beta$ requires normally distributed errors.
18. Heteroscedasticity biases the OLS estimator.
19. Omitting a relevant covariate that is correlated with the included ones biases $\hat{\boldsymbol\beta}$.
20. Autocorrelated errors leave OLS unbiased but affect its efficiency.

## Block 6 — Gauss–Markov

21. Under A1–A5, OLS has the smallest variance among **all** unbiased estimators.
22. Normality of the errors is required for the Gauss–Markov theorem to hold.
23. The LS estimator is BLUE **if and only if** the errors have zero mean and constant variance.
24. There exist biased estimators whose variance is smaller than that of OLS.

## Block 7 — Estimating $\sigma^2$

25. $\hat\sigma^2=\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}/(n-p)$ is unbiased for $\sigma^2$.
26. AIC is computed using $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}/(n-p)$.
27. The maximum likelihood estimator of $\sigma^2$ has denominator $n$.
28. The quantity R labels "residual standard error" is $\hat\sigma^2$.

## Block 8 — $R^2$ and SSE

29. Adding a covariate can decrease $R^2$.
30. The residual sum of squares may increase as more variables are added.
31. Adjusted $R^2$ can be negative.
32. Adjusted $R^2$ can decrease when a covariate is added.

## Block 9 — AIC and BIC

33. AIC penalises model complexity more heavily than BIC.
34. For $n>8$, BIC tends to select smaller models than AIC.
35. Adding an irrelevant predictor must decrease AIC.
36. AIC values are comparable across models fitted to different data sets.

## Block 10 — The t-test

37. To test $H_0:\beta_j=c$, the statistic is $\hat\beta_j/\widehat{\text{se}}(\hat\beta_j)$.
38. In an R regression table, $\hat\beta = t\times\widehat{\text{se}}$.
39. $\hat\beta_0$ and $\hat\beta_1$ are always uncorrelated.
40. $\widehat{\text{se}}(\hat\beta_j)$ uses the $j$th **diagonal** element of $(\boldsymbol{X}'\boldsymbol{X})^{-1}$.

## Block 11 — The F-test

41. The F statistic is negative when the restricted model fits better than the unrestricted one.
42. When $r=1$, $\sqrt{F}=|t|$.
43. Rejecting a joint null hypothesis means every restriction in it fails.
44. Given $r$ and $n-p$, $F$ can be computed from $R^2$ and $R^2_{H_0}$ alone.

## Block 12 — Quantiles

45. A two-sided t-test at $\alpha=0.05$ uses the $0.975$ quantile.
46. An F-test at $\alpha=0.05$ uses the $0.975$ quantile.
47. A 99% confidence interval uses the $0.995$ quantile.
48. The F-test rejects only in the upper tail.

## Block 13 — Intervals

49. If a CI for $\beta_j$ does **not** contain zero, we fail to reject $H_0:\beta_j=0$.
50. When the CI contains zero, we cannot reject $H_0:\beta_j=0$.
51. At the same $\boldsymbol{x}_0$, the prediction interval is wider than the confidence interval for the mean.
52. As $n\to\infty$ the width of the prediction interval tends to zero.

## Block 14 — Residuals

53. With an intercept in the model, the residuals sum to zero by construction.
54. A residual mean of approximately zero is evidence that the model fits well.
55. $\text{Var}(\hat\varepsilon_i)=\sigma^2$ for every $i$.
56. Even when all model assumptions hold, the estimated residuals are correlated with one another.

## Block 15 — Diagnostic plots

57. In a QQ plot the points should follow a horizontal line.
58. A plot of residuals against fitted values cannot reveal non-linearity.
59. A funnel shape in residuals versus fitted values suggests heteroscedasticity.
60. A high-leverage observation always distorts the fitted line.

## Block 16 — Leverage and influence

61. $h_{ii}$ measures how unusual observation $i$ is in the **covariates**.
62. $\text{tr}(\boldsymbol{H})=p$.
63. Cook's distance combines the standardised residual with the leverage.
64. An outlier in $y$ and a high-leverage point are the same thing.

## Block 17 — Multicollinearity and dummies

65. Near-multicollinearity inflates the variances but does not bias $\hat{\boldsymbol\beta}$.
66. If VIF $\approx1$ for all variables, multicollinearity is likely to be a concern.
67. A categorical covariate with $m$ categories requires $m$ dummy variables alongside an intercept.
68. The reference category is the level that does **not** appear in the R output.

## Block 18 — The logit model

69. A linear model is inappropriate for binary $y$ partly because fitted values can fall outside $[0,1]$.
70. In a logit model, $\hat\beta_j$ is the increase in $P(y=1)$ for a one-unit rise in $x_j$.
71. $\exp(\hat\beta_j)$ is the odds ratio associated with a one-unit increase in $x_j$.
72. The logit model is estimated by ordinary least squares.

---
---

# ANSWER KEY

> 🔴 marks a statement whose structure appeared **verbatim** in one of your five past papers. Cross-references point at the full treatment in the chapter `32-TRAPS.md` files.

## Block 1

| # | | Reason |
|---|---|---|
| 1 | **T** | One column per **parameter**: intercept + $k$ covariates. |
| 2 | **F** | 🔴 $\boldsymbol{X}'\boldsymbol{X}$ is $p\times p$, and rank never exceeds the smaller dimension. Full rank means $p$. *(Trap A2)* |
| 3 | **T** | That column is what makes $\beta_0$ an intercept. |
| 4 | **F** | Adding a column can only leave the rank alone or raise it — the old columns are all still there. |

## Block 2

| # | | Reason |
|---|---|---|
| 5 | **T** | Count the $\beta$'s including the intercept: $n-p=n-k-1$. *(Trap A1)* |
| 6 | **F** | 🔴 Numerator df is $\boldsymbol{r}$, the number of restrictions — not $p+1$. *(Trap A3)* |
| 7 | **F** | Both the SSE and the df in the denominator come from the **unrestricted** model. *(Trap G6)* |
| 8 | **T** | 🔴 Overall test has $r=k$ restrictions and $n-p=n-k-1$ residual df. |

## Block 3

| # | | Reason |
|---|---|---|
| 9 | **F** | 🔴🔴 **The single biggest trap in 3.3.** Rewrite as $\beta_1+\beta_2-\beta_3=0$ and count equals signs: **one**. The "3" counts betas mentioned, which is never what $r$ means. *(Trap B1)* |
| 10 | **T** | $\boldsymbol{C}$ is $r\times p$. The $\beta_0$ column is usually zeros but must be present or dimensions don't conform. *(Trap B2)* |
| 11 | **T** | Two independent equations ⟹ $r=2$. |
| 12 | **T** | $\beta_{\text{crim}}=3\beta_{\text{rad}}-0.1$ must become $\beta_{\text{crim}}-3\beta_{\text{rad}}=-0.1$ first. *(Trap B3)* |

## Block 4

| # | | Reason |
|---|---|---|
| 13 | **T** | That is the definition of the least squares estimator. |
| 14 | **F** | 🔴 Uniqueness needs **full rank**. Homoscedasticity is completely irrelevant to identification — classic irrelevant "as long as…" clause. *(Trap E3)* |
| 15 | **F** | 🔴 LAD targets the conditional **median**, has no closed form, gives different estimates, and is more robust to outliers. *(Trap E5)* |
| 16 | **T** | Positive definite Hessian ⟹ minimum. Worth stating in the derivation — it is part of the marking key. |

## Block 5

| # | | Reason |
|---|---|---|
| 17 | **F** | Unbiasedness needs only A1, A2 and A5. Normality buys **exact tests**, nothing else here. |
| 18 | **F** | 🔴 OLS stays **unbiased and consistent**. What you lose is **efficiency** and **valid standard errors**. *(Trap E1)* |
| 19 | **T** | Omitted-variable bias — this is a violation of A1, and A1 is the only one that biases $\hat{\boldsymbol\beta}$. |
| 20 | **T** | 🔴 Same structure as 18, stated correctly this time. |

## Block 6

| # | | Reason |
|---|---|---|
| 21 | **F** | Minimum variance among estimators that are **linear in $\boldsymbol{y}$ AND unbiased**. Drop "linear" and the claim is false. |
| 22 | **F** | 🔴 Gauss–Markov needs only A1–A5. Saying so explicitly is worth a mark. *(Trap E2)* |
| 23 | **F** | 🔴 Two faults: the "**iff**" is too strong, and the list omits uncorrelated errors, correct specification and full rank. *(Trap E2)* |
| 24 | **T** | Ridge, for instance — biased, lower variance. This is exactly why "unbiased" must stay in the BLUE statement. |

## Block 7

| # | | Reason |
|---|---|---|
| 25 | **T** | The $n-p$ denominator is what makes it unbiased. |
| 26 | **F** | 🔴 AIC and BIC use the **ML** estimator, denominator $\boldsymbol{n}$. The book says so explicitly. *(Trap C1)* |
| 27 | **T** | $\hat\sigma^2_{ML}=\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}/n$. |
| 28 | **F** | It is $\hat\sigma$ — a **standard error**, so already square-rooted. If your answer is 1276 and the output says 34, you skipped the root. *(Trap C4)* |

## Block 8

| # | | Reason |
|---|---|---|
| 29 | **F** | 🔴 Adding a column expands the column space, so the projection can only get **closer** to $\boldsymbol{y}$. *(Trap F1)* |
| 30 | **F** | 🔴 Same fact, other direction: SSE can only fall or stay. |
| 31 | **T** | 🔴 $\bar R^2$ **can** go negative — the penalty is unbounded below. A paper claimed it can't; that was FALSE. |
| 32 | **T** | And this is the entire point of §3.4 — $\bar R^2$, AIC and BIC **can** worsen where $R^2$ cannot. |

## Block 9

| # | | Reason |
|---|---|---|
| 33 | **F** | 🔴 Backwards. $\log(n)>2$ once $n>7.4$, so **BIC penalises more**. *B for Bigger penalty.* *(Trap F2)* |
| 34 | **T** | Heavier penalty ⟹ smaller selected models. |
| 35 | **F** | 🔴 An irrelevant variable barely improves fit while costing a full penalty, so AIC typically **rises**. And "must" is too strong either way. *(Trap F3)* |
| 36 | **F** | Both contain $n\log(\hat\sigma^2)$ computed on the data at hand — only comparable within the **same** data set. |

## Block 10

| # | | Reason |
|---|---|---|
| 37 | **F** | Missing the $-c$: $t=(\hat\beta_j-c)/\widehat{\text{se}}$. Students compute $\hat\beta_j/\widehat{\text{se}}$ out of habit. *(Trap G2)* |
| 38 | **T** | The rearrangement that fills missing R output. |
| 39 | **F** | 🔴 $\text{Cov}(\hat\beta_0,\hat\beta_1)=-\sigma^2\bar x/\sum(x_i-\bar x)^2$, zero **only if $\bar x=0$**. The word "always" is what breaks it. *(Trap E7)* |
| 40 | **T** | Diagonal only — and label the rows first, since $\beta_1$ is the **second** entry. |

## Block 11

| # | | Reason |
|---|---|---|
| 41 | **F** | The restricted model can never fit better: $\text{SSE}_{H_0}\geq\text{SSE}$ always, so $F\geq0$. A negative $F$ means you swapped the SSEs. |
| 42 | **T** | A useful sanity check whenever $r=1$. |
| 43 | **F** | **At least one** fails. The book states this explicitly. *(Trap G5)* |
| 44 | **T** | $F=\dfrac{(R^2-R^2_{H_0})/r}{(1-R^2)/(n-p)}$ — the second formula exists precisely for when you're given $R^2$ instead of SSE. |

## Block 12

| # | | Reason |
|---|---|---|
| 45 | **T** | Two-sided ⟹ split $\alpha$ ⟹ $1-\alpha/2$. |
| 46 | **F** | 🔴 **The most common numerical error in this exam.** F is one-sided ⟹ $1-\alpha=0.95$. *(Trap D1)* |
| 47 | **T** | $1-0.01/2=0.995$. |
| 48 | **T** | And that is *why* it uses $1-\alpha$ rather than $1-\alpha/2$. |

## Block 13

| # | | Reason |
|---|---|---|
| 49 | **F** | 🔴 Backwards. Zero **outside** the interval ⟹ we **do reject**. *(Trap G1)* |
| 50 | **T** | 🔴 Same fact stated correctly. *Zero inside the net ⟹ zero still possible ⟹ don't reject.* |
| 51 | **T** | 🔴 The extra "$\mathbf{1}+$" under the root is the new observation's own error $\varepsilon_0$. *(Trap G3)* |
| 52 | **F** | It converges to $\pm t\hat\sigma$, not to zero. Estimation error vanishes; the individual's own noise never does. *Individuals are not averages.* |

## Block 14

| # | | Reason |
|---|---|---|
| 53 | **T** | Guaranteed by the intercept's normal equation. |
| 54 | **F** | It holds **by construction** whenever there is an intercept — a terrible model satisfies it too. Diagnostics read the **pattern**, not the mean. *(Trap H5)* |
| 55 | **F** | $\text{Var}(\hat\varepsilon_i)=\sigma^2(1-h_{ii})$, which varies with $i$. This is the whole reason standardised residuals exist. |
| 56 | **T** | 🔴 $\text{Cov}(\hat\varepsilon_i,\hat\varepsilon_j)=-\sigma^2h_{ij}\neq0$ — even under perfect assumptions. *(Trap H1)* |

## Block 15

| # | | Reason |
|---|---|---|
| 57 | **F** | 🔴 The **45° diagonal** ($y=x$). A horizontal line would mean every empirical quantile is identical — no variation at all. Classic true-first-clause construction. *(Trap H2)* |
| 58 | **F** | 🔴 Residuals-vs-fitted is *precisely* the non-linearity detector — a curve in the cloud. *(Trap H3)* |
| 59 | **T** | Spread growing with the fitted value is the textbook funnel. |
| 60 | **F** | High leverage **alone is harmless** — a distant point sitting exactly on the line actually improves precision. It is dangerous only alongside a large residual. *(Trap H4)* |

## Block 16

| # | | Reason |
|---|---|---|
| 61 | **T** | $h_{ii}=\boldsymbol{x}_i'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_i$ — a function of $\boldsymbol{x}$ only, no $y$ in it. |
| 62 | **T** | The $p$ dimensions the projection uses up. |
| 63 | **T** | $D_i=\dfrac{r_i^2}{p}\cdot\dfrac{h_{ii}}{1-h_{ii}}$ — unusual $y$ **times** unusual $\boldsymbol{x}$. |
| 64 | **F** | Outlier = unusual in $y$ ($\lvert r_i\rvert$). Leverage = unusual in $\boldsymbol{x}$ ($h_{ii}$). Influence = both together. Three different things. *(Trap H4)* |

## Block 17

| # | | Reason |
|---|---|---|
| 65 | **T** | 🔴 Near-collinearity leaves $\hat{\boldsymbol\beta}$ unbiased and still **BLUE** — only imprecise. *(Trap E4)* |
| 66 | **F** | 🔴 Backwards. VIF $\approx1$ means **no** collinearity; concern starts around 5–10. *(Trap E4)* |
| 67 | **F** | 🔴 $m-1$ dummies. All $m$ plus an intercept ⟹ columns sum to the intercept ⟹ singular $\boldsymbol{X}'\boldsymbol{X}$ — the **dummy variable trap**. |
| 68 | **T** | And every other coefficient is a difference **from** it. |

## Block 18

| # | | Reason |
|---|---|---|
| 69 | **T** | 🔴 Lead with this one. The other three: $\text{Var}(y)=\pi(1-\pi)$ so heteroscedastic by construction · errors take only two values so cannot be normal · constant marginal effects implausible near the boundaries. |
| 70 | **F** | 🔴 **The Exam 2025 trap.** $\hat\beta_j$ is the effect on the **log-odds**. On the probability scale the effect is $\hat\beta_j\pi(1-\pi)$ — not constant. Only the **sign** transfers reliably. |
| 71 | **T** | The odds are multiplied by $\exp(\hat\beta_j)$. |
| 72 | **F** | Maximum likelihood. There is no closed form — but you are never asked to derive it. |

---

## Score yourself

| Score | Read this |
|---|---|
| **65–72** | You're ready for Exercise 1. Re-drill only the ones you got right without a reason. |
| **55–64** | Solid. Go to the specific `32-TRAPS.md` sections cited on your misses. |
| **45–54** | Re-read all three `32-TRAPS.md` files today, then redo this drill tomorrow. |
| **under 45** | You're guessing. Go back to `chapter-03-classical-linear-model/10-SUMMARY.md` before drilling again — the drill can't teach, it can only test. |

**The reasons matter more than the verdicts.** Every statement here has three or four variants the examiner could write; only the reason generalises.
