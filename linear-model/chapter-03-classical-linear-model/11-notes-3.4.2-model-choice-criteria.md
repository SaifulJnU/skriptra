# 3.4.2 — Model Choice Criteria

> **Sheet 5 is entirely this section.** AIC and BIC are computed by hand in at least one past paper, and appear as T/F statements in several. Get the formulas exactly right — including which $\hat\sigma^2$ to use.

---

## 1. The common structure

Every criterion has the same shape:

$$\text{criterion} = \underbrace{\text{(how badly it fits)}}_{\text{decreases with complexity}} + \underbrace{\text{(penalty for complexity)}}_{\text{increases with complexity}}$$

**Smaller is better** for AIC, BIC and Mallow's $C_p$. **Larger is better** for $\bar R^2$.

Notation for this section: $|M|$ = number of **regression parameters** in model $M$ (including the intercept). So $|M|=p$ for the full model.

---

## 2. Corrected (adjusted) coefficient of determination

$$\boxed{\;\bar R^2 = 1-\frac{n-1}{n-p}\left(1-R^2\right)\;}$$

**Larger is better.**

### Properties

- **Can decrease** when a variable is added — unlike $R^2$
- **Can be negative** (if the model is worse than the intercept-only model)
- Always $\bar R^2\leq R^2$
- The penalty grows with $p$ through the factor $\frac{n-1}{n-p}$

### ⚠️ The book's own warning

> *"At this point, we advise against its usage, since the 'penalty' for newly included covariates appears to be too small. It can be shown that $\bar R^2$ already increases when a variable with a **t-value greater than 1** is included in the model, implying we would include variables with a p-value of about 0.3."*

**Learn that sentence.** If asked to compare criteria, saying *"adjusted $R^2$ penalises too weakly — it admits variables with $|t|>1$, i.e. p-values around 0.3 — so AIC or BIC is preferable"* is a strong answer.

### Past-paper T/F

> 🔴 *Linear_model_exam_sheet, Block III(iv):* "Adjusted $R^2$ takes into account the number of covariates and **can never be negative**." → **FALSE.** It can.

---

## 3. ⭐ AIC — Akaike Information Criterion

**General definition:**
$$\text{AIC}=-2\,\ell(\hat{\boldsymbol\beta}_M,\hat\sigma^2)+2\,(|M|+1)$$

where $\ell$ is the maximised log-likelihood. *(The $+1$ counts $\sigma^2$ as a parameter.)*

**For a linear model with Gaussian errors, this reduces to:**

$$\boxed{\;\text{AIC}=n\log(\hat\sigma^2)+2(|M|+1)\;}$$

**🔴 with the ML variance estimator:**

$$\boxed{\;\hat\sigma^2=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n}\quad\text{— divide by } n,\ \textbf{NOT } n-p\;}$$

The book is explicit: *"Note that the ML estimator $\hat\sigma^2=\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}/n$ is considered in AIC and not the usual unbiased variance estimator."*

**Also note:** $\log$ means the **natural** logarithm. Sheet 5 says so explicitly — *"the book uses $\log(x)$ as the natural logarithm with base $e$ (and not base 10)."*

**Smaller AIC = better.**

---

## 4. ⭐ BIC — Bayesian Information Criterion

$$\text{BIC}=-2\,\ell(\hat{\boldsymbol\beta}_M,\hat\sigma^2)+\log(n)(|M|+1)$$

$$\boxed{\;\text{BIC}=n\log(\hat\sigma^2)+\log(n)\,(|M|+1)\;}$$

Same $\hat\sigma^2_{ML}$. **Smaller is better.** *(BIC$\times\frac12$ is sometimes called the Schwarz criterion.)*

### AIC vs BIC — the only difference is the penalty

| | Penalty per parameter | Behaviour |
|---|---|---|
| **AIC** | $2$ | constant |
| **BIC** | $\log(n)$ | **grows with sample size** |

$\log(n)>2 \iff n>e^2\approx7.39$. So **for essentially any real dataset, BIC penalises more heavily than AIC and therefore selects smaller models.**

> 🔴 **Linear_model_exam_sheet, Block III(iii):** *"AIC and BIC are both criteria for model selection where they penalize models for the number of parameters, but **AIC penalizes more heavily than BIC**."* → **FALSE.** It's the other way round (for $n>8$).
>
> **Mnemonic:** *BIC is the Bigger penalty. B for Bigger, B for BIC.*

| | AIC | BIC |
|---|---|---|
| Goal | best **prediction** | find the **true** model |
| Asymptotics | efficient for prediction | **consistent** for selection |
| Selects | larger models | smaller models |

---

## 5. Mallow's $C_p$

$$C_p=\frac{\sum_i(y_i-\hat y_i^M)^2}{\hat\sigma^2_{\text{full}}}-n+2|M|$$

Same idea: fit term plus complexity penalty. Uses $\hat\sigma^2$ from the **full** model as a reference. For Gaussian errors it's essentially equivalent to AIC.

**Know it exists and what it does. Don't memorise it in detail.**

---

## 6. Cross-validation

Directly estimates out-of-sample prediction error rather than approximating it with a penalty.

**$r$-fold cross-validation:**
1. Split the data into $r$ subsets of similar size
2. For each subset $j$: fit the model on the other $r-1$ subsets, predict subset $j$, record the squared errors
3. Average over all folds

**Leave-one-out ($r=n$)** has a closed form for linear models:

$$\text{CV}=\frac{1}{n}\sum_{i=1}^n\left(\frac{\hat\varepsilon_i}{1-h_{ii}}\right)^2$$

**Elegant fact:** you get leave-one-out CV without refitting the model $n$ times — the leverages $h_{ii}$ do the work.

**Advantages:** makes no distributional assumption; directly targets prediction.
**Disadvantages:** computationally heavier; results vary with the random split.

---

## 7. 📝 Worked example — Sheet 5, Exercise 1

Model 1: wage on age + 4 education dummies + health. $n=3000$, $p=|M|=7$, $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}=3819720$, $R^2=0.2685$.

### (a) Corrected coefficient of determination

$$\bar R^2=1-\frac{n-1}{n-p}(1-R^2)=1-\frac{2999}{2993}(1-0.2685)=1-\frac{2999}{2993}(0.7315)$$
$$=1-1.002005\times0.7315=1-0.73297=\boxed{0.2670}$$

Slightly below $R^2=0.2685$, as it must be.

### (b) AIC and BIC

**Step 1 — the ML variance estimate.**
$$\hat\sigma^2=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n}=\frac{3819720}{3000}=1273.240$$

**Step 2 — the log.**
$$\log(1273.240)=7.14929 \quad\text{(natural log!)}$$

**Step 3 — the fit term.**
$$n\log(\hat\sigma^2)=3000\times7.14929=21447.96$$

**Step 4 — AIC.** $|M|=7$, so $|M|+1=8$:
$$\text{AIC}=21447.96+2(8)=21447.96+16=\boxed{21463.96}$$

**Step 5 — BIC.** $\log(3000)=8.00637$:
$$\text{BIC}=21447.96+8.00637\times8=21447.96+64.05=\boxed{21512.01}$$

> ⚠️ **Three ways to get this wrong:** (i) dividing by $n-p$ instead of $n$; (ii) using $\log_{10}$; (iii) forgetting the $+1$ for $\sigma^2$. All three are easy and all three cost the mark.

---

## 8. 📝 Worked example — Sheet 5, Exercise 2 (model comparison)

Model 2: all available covariates. From the R output: `Residual standard error: 34 on 2983 degrees of freedom`, `Multiple R-squared: 0.3396, Adjusted R-squared: 0.3361`, and $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}=3448498$.

**First, get $p$:** $n-p=2983$ and $n=3000$, so $p=|M|=17$.

### (c) AIC and BIC

$$\hat\sigma^2=\frac{3448498}{3000}=1149.499,\qquad \log(1149.499)=7.04709$$
$$n\log(\hat\sigma^2)=3000\times7.04709=21141.25$$

$$\text{AIC}=21141.25+2(18)=21141.25+36=\boxed{21177.25}$$
$$\text{BIC}=21141.25+8.00637\times18=21141.25+144.11=\boxed{21285.36}$$

### (d) Which model do you prefer?

| Criterion | Model 1 (7 params) | Model 2 (17 params) | Winner |
|---|---|---|---|
| $\bar R^2$ (larger better) | 0.2670 | **0.3361** | Model 2 |
| **AIC** (smaller better) | 21463.96 | **21177.25** | Model 2 |
| **BIC** (smaller better) | 21512.01 | **21285.36** | Model 2 |

> **Model answer:** *All three criteria agree on Model 2. Its adjusted $R^2$ is substantially higher (0.336 vs 0.267), and both AIC and BIC are lower — by roughly 287 and 227 units respectively. Crucially, **BIC** also favours Model 2 despite carrying ten extra parameters and a $\log(n)=8.01$ penalty per parameter; since BIC penalises complexity most heavily, agreement from BIC is the strongest evidence that the improvement in fit is genuine and not merely a consequence of the larger model size. I would therefore prefer Model 2.*
>
> *Note also that the comparison is legitimate because both models are fitted to the same $n=3000$ observations — AIC and BIC values are only comparable across models fitted to identical data.*

**That last caveat is worth mentioning.** It's a real condition and it shows understanding.

---

## 9. Past-paper T/F on model choice

| Statement | Verdict | Why |
|---|---|---|
| "AIC penalises more heavily than BIC" | ❌ **FALSE** | BIC's $\log n$ penalty is bigger for $n>8$ |
| "Adjusted $R^2$ can never be negative" | ❌ **FALSE** | it can |
| "$R^2$ can decrease as more variables are added" | ❌ **FALSE** | $R^2$ is monotone increasing |
| "RSS may increase as more variables are added" | ❌ **FALSE** | SSE is monotone decreasing |
| "When an unrelated predictor is added, AIC **must** decrease" *(WS23/24 III(ii))* | ❌ **FALSE** | AIC will typically **increase** — the penalty outweighs the tiny fit gain. And "must" is far too strong either way |

**Notice the pattern:** several of these test the same underlying fact — $R^2$ and SSE are **monotone** in model size, which is exactly why the penalised criteria exist. One idea, four questions.

---

## 10. Model choice cheat sheet

$$\bar R^2 = 1-\frac{n-1}{n-p}(1-R^2) \qquad\text{(larger better; penalty too weak)}$$

$$\text{AIC}=n\log(\hat\sigma^2)+2(|M|+1) \qquad\text{(smaller better)}$$

$$\text{BIC}=n\log(\hat\sigma^2)+\log(n)(|M|+1) \qquad\text{(smaller better)}$$

$$\hat\sigma^2=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n}\quad\textbf{(ML — divide by }n\textbf{)},\qquad \log=\ln$$

**The $n\log(\hat\sigma^2)$ term is identical in AIC and BIC.** Compute it once, then just add the two different penalties. Saves a minute.

---

## 11. Key takeaways

1. Every criterion = **fit + complexity penalty**.
2. $\bar R^2$: larger better; **can decrease, can go negative**; penalty is too weak (admits $|t|>1$).
3. **AIC $=n\log(\hat\sigma^2)+2(|M|+1)$**, smaller better.
4. **BIC $=n\log(\hat\sigma^2)+\log(n)(|M|+1)$**, smaller better.
5. 🔴 **Use $\hat\sigma^2=\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}/n$ (ML), and natural log.**
6. 🔴 **BIC penalises MORE than AIC** for $n>8$ — *B for Bigger*. BIC ⟹ smaller models.
7. AIC ≈ prediction; BIC ≈ true-model identification.
8. Cross-validation estimates prediction error directly; LOO-CV $=\frac1n\sum\left(\frac{\hat\varepsilon_i}{1-h_{ii}}\right)^2$.
9. AIC/BIC are **only comparable across models fitted to the same data**.
