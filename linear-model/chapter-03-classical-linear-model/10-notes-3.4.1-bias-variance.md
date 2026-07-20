# 3.4.1 — Bias, Variance and Prediction Quality

> The conceptual foundation for everything in 3.4. Short, and rarely computed — but it's the *reason* AIC and BIC exist, and an "explain the intuition" question here is worth 2 points.

---

## 1. The problem $R^2$ can't solve

$R^2$ **never decreases** when you add a covariate. Add dummies for the weekday someone was born on and $R^2$ goes *up*. So:

> **A criterion that always favours the biggest model cannot be used to choose a model.**

We need something that asks a different question: not *"how well does this fit the data I have?"* but *"how well would this predict data I haven't seen?"*

---

## 2. The decomposition

Suppose the true model is $y=f(\boldsymbol{x})+\varepsilon$ with $\text{Var}(\varepsilon)=\sigma^2$, and we predict a new observation at $\boldsymbol{x}_0$ using $\hat f(\boldsymbol{x}_0)$ fitted on our sample. The expected squared prediction error is

$$\boxed{\;E\left[(y_0-\hat f(\boldsymbol{x}_0))^2\right]=\underbrace{\sigma^2}_{\text{irreducible}}+\underbrace{\left[\text{Bias}(\hat f(\boldsymbol{x}_0))\right]^2}_{\text{systematic error}}+\underbrace{\text{Var}(\hat f(\boldsymbol{x}_0))}_{\text{estimation noise}}\;}$$

| Term | What it is | Can we reduce it? |
|---|---|---|
| $\sigma^2$ | the new observation's own randomness | **No.** Ever. |
| $\text{Bias}^2$ | the model is systematically wrong | Yes — add relevant covariates / better functional form |
| $\text{Variance}$ | estimates wobble from sample to sample | Yes — **fewer** parameters, or more data |

---

## 3. 🔑 The tradeoff

**Bias and variance move in opposite directions as model complexity changes.**

```
error
  │
  │╲                                      ╱  ← total prediction error
  │ ╲                                   ╱      (the U-curve)
  │  ╲                                ╱
  │   ╲                            ╱
  │     ╲___                   ╱  ← VARIANCE (grows with complexity)
  │         ╲___          ╱
  │              ╲___ ╱
  │            ╱      ╲___
  │        ╱               ╲______  ← BIAS² (falls with complexity)
  │────────────────────────────────
  │- - - - - - - - - - - - - - - -   ← σ² (irreducible floor)
  └────────────────────────────────── model complexity
            ↑
        the sweet spot
```

| Model | Bias | Variance | Symptom |
|---|---|---|---|
| **Too simple** (underfit) | **high** | low | misses real structure; systematic residual patterns |
| **Just right** | low | moderate | ← what we want |
| **Too complex** (overfit) | low | **high** | fits noise; great $R^2$, terrible out-of-sample |

### The extreme cases make it vivid

**Maximum simplicity:** $\hat y_i=\bar y$ for everyone. Zero variance in the *slopes* (there are none), but enormous bias — you're ignoring all covariates.

**Maximum complexity:** $p=n$. The model passes through every point exactly. $R^2=1$, $\text{SSE}=0$, and $\hat\sigma^2=0/0$ is undefined. **Zero bias, infinite variance.** Show it one new observation and it fails completely — it memorised your sample rather than learning the pattern.

> **The key insight for the exam:** *adding a variable always reduces (or leaves unchanged) the bias, but always increases the variance.* Whether it's worth it depends on which effect is larger — and that is exactly the trade AIC and BIC price.

---

## 4. Why unbiasedness isn't the whole story

Gauss–Markov gives you the **best unbiased** estimator. But look at the decomposition: **bias enters squared, variance enters linearly.** A little bias can buy a lot of variance reduction and lower the total.

This is why:
- **Omitting a weakly relevant covariate** can *improve* prediction — you accept a small bias for a meaningful variance reduction
- **Ridge and lasso** (Chapter 4) deliberately introduce bias to shrink variance
- **BLUE is a constrained optimum**, not an unconstrained one. "Best among unbiased" ≠ "best."

> 🔴 This is what makes *Exam Summer 2025, Ex 1(e)* subtle. "A BLUE is best in the sense that there is no other **linear unbiased** estimator with lower variance" → **TRUE**, precisely because of the two qualifiers. Drop them and it's false — biased estimators can beat OLS on mean squared error.

---

## 5. Connecting to the exam's actual questions

> 🔴 **Exam Summer 2025, Ex 1(a):** *"Adding a variable which is not correlated with the dependent variable will not affect the unbiasedness of the OLS estimator, but it may affect its variance."* → **TRUE.**
>
> **Why:** including an irrelevant covariate doesn't bias the others (its true coefficient is zero, so nothing is omitted). But it costs a degree of freedom and, if it's correlated with the included covariates, inflates their variances. **Unbiased but less efficient — the bias–variance tradeoff in one sentence.**

> 🔴 **Exam Summer 2025, Ex 1(c):** *"Adding dummies for the weekday a person was born on can be expected to lower $R^2$."* → **FALSE.** $R^2$ can only rise. But **adjusted $R^2$, AIC and BIC would all get worse**, correctly flagging that these variables add variance without reducing bias. That contrast — $R^2$ up, adjusted criteria down — *is* the point of Section 3.4.

> 🔴 **Exam Summer 2025, Ex 2(d)** [2 pts]: *"A colleague suggests adding many more functions of age — $(\text{age}-48)^3$, $\text{age}^4$, $\log(\text{age})$ — because it will improve the fit. How would you decide whether to do so? State two methods."*
>
> **Model answer:** *The colleague is right that the in-sample fit will improve — $R^2$ cannot decrease when covariates are added — but this is not a reason to include them, since increasing model complexity raises the variance of the estimates and risks **overfitting**, i.e. modelling noise rather than signal. Two methods to decide:*
>
> 1. ***Information criteria.** Compare AIC (or BIC) between the models. These add a penalty for the number of parameters to the goodness-of-fit term, so a variable is only worth including if it improves fit by more than the penalty. Choose the model with the smaller value.*
> 2. ***An F-test for the additional terms.** Test $H_0$: the coefficients of all added terms are jointly zero, using $F=\frac{(\text{SSE}_{H_0}-\text{SSE})/r}{\text{SSE}/(n-p)}$. If we cannot reject $H_0$, the extra terms do not significantly improve the model.*
>
> *(Also acceptable: **cross-validation** — compare out-of-sample prediction error; or **adjusted $R^2$**, though its penalty is generally regarded as too weak.)*

---

## 6. Two distinct goals, two different answers

| Goal | Criterion | Tends to select |
|---|---|---|
| **Best prediction** | AIC, cross-validation, Mallow's $C_p$ | somewhat **larger** models |
| **Identifying the true model** | BIC | **smaller**, more parsimonious models |

BIC's penalty ($\log n$ per parameter) grows with sample size while AIC's (2 per parameter) doesn't. So BIC is **consistent** for model selection — it picks the true model with probability → 1 as $n\to\infty$ — while AIC is **asymptotically efficient** for prediction.

**Neither is "correct."** They answer different questions. If asked which to use, say what you're optimising for.

---

## 7. Key takeaways

1. $$E[(y_0-\hat f)^2]=\sigma^2+\text{Bias}^2+\text{Variance}$$
2. $\sigma^2$ is **irreducible**. No model, no amount of data removes it.
3. **More complexity ⟹ less bias, more variance.** Total error is U-shaped.
4. Overfitting = fitting the noise. Perfect in-sample fit ($p=n$) means zero out-of-sample value.
5. **Bias enters squared, variance linearly** — so accepting a little bias can pay. This is why BLUE isn't the last word.
6. $R^2$ measures in-sample fit only and always favours bigger models — hence adjusted $R^2$, AIC, BIC, cross-validation.
7. **AIC ≈ prediction; BIC ≈ finding the true model.** BIC penalises harder and picks smaller models.
