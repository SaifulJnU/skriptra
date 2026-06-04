# Chapter 5: Linear Regression — Least Squares & Normal Equations — Practice True/False Questions

---

## Questions

**Q1.** The linear model is written as y = β₀ + xᵀβ + ε.

**Q2.** In the model y = β₀ + xᵀβ + ε, the term β₀ is the intercept (bias).

**Q3.** Lecture 5 assumes the noise term ε follows a Gaussian distribution N(0, σ²).

**Q4.** The parameter vector is collected as θ = (β₀, βᵀ)ᵀ.

**Q5.** Because E[ε] = 0, the conditional expectation E[y | x] equals β₀ + xᵀβ = f(x).

**Q6.** The regression function f(x) is the conditional expectation of y given x.

**Q7.** In augmented notation, x̃ = (1, xᵀ)ᵀ adds a leading 1 to the feature vector.

**Q8.** With augmented notation, the model can be written as f(x) = x̃ᵀβ̃.

**Q9.** The leading 1 in x̃ exists so that the intercept β₀ can be absorbed into the dot product.

**Q10.** The least-squares objective is ‖y − Xθ‖₂² = (y − Xθ)ᵀ(y − Xθ).

**Q11.** Ordinary Least Squares minimizes the sum of absolute residuals.

**Q12.** Setting the gradient of the least-squares objective to zero yields the normal equations XᵀX θ = Xᵀy.

**Q13.** The closed-form least-squares estimator is θ̂ = (XᵀX)⁻¹Xᵀy.

**Q14.** The OLS solution requires numerical optimization because no analytical form exists.

**Q15.** The matrix XᵀX must be invertible for the standard OLS formula to apply.

**Q16.** XᵀX is invertible if and only if X has full column rank.

**Q17.** Because the L2 loss with a linear model is convex, the OLS solution is the global minimum.

**Q18.** When features are perfectly collinear, XᵀX becomes singular and the inverse fails to exist.

**Q19.** The gradient of (y − Xθ)ᵀ(y − Xθ) with respect to θ is −2Xᵀ(y − Xθ).

**Q20.** The sample mean (1/n)Σxᵢ = x̄ converges to the theoretical mean E[X] as n grows.

**Q21.** For a discrete random variable, E[X] = Σ j·P(X = j).

**Q22.** For a continuous random variable with density f, E[X] = ∫ x f(x) dx.

**Q23.** The empirical risk is a sample average that estimates the theoretical (expected) risk.

**Q24.** The absolute (L1) loss is g(x) = |x|.

**Q25.** The absolute loss |y − f(x)| is differentiable everywhere, including at 0.

**Q26.** The absolute (L1) loss is more robust to outliers than the squared (L2) loss.

**Q27.** The optimal constant model under the absolute loss is the median.

**Q28.** Under Gaussian noise, the least-squares (OLS) fit coincides with the maximum-likelihood estimate.

**Q29.** Adding a positive λI to XᵀX (ridge regularization) can restore invertibility when XᵀX is singular.

**Q30.** In the design matrix X for an intercept model, the first column consists entirely of ones.

---

## Answers

| Q | Answer | Key Reason |
|---|--------|------------|
| 1 | **TRUE** | Standard linear model form |
| 2 | **TRUE** | β₀ = intercept / bias |
| 3 | **TRUE** | ε ~ N(0, σ²) (top of the side notes) |
| 4 | **TRUE** | θ = (β₀, βᵀ)ᵀ |
| 5 | **TRUE** | E[ε]=0 ⇒ E[y|x] = β₀ + xᵀβ = f(x) |
| 6 | **TRUE** | f(x) is the conditional mean |
| 7 | **TRUE** | x̃ = (1, xᵀ)ᵀ |
| 8 | **TRUE** | f(x) = x̃ᵀβ̃ |
| 9 | **TRUE** | Leading 1 absorbs the intercept |
| 10 | **TRUE** | OLS objective as squared norm |
| 11 | **FALSE** | OLS minimizes **squared** residuals (that is L1, not OLS) |
| 12 | **TRUE** | Normal equations XᵀX θ = Xᵀy |
| 13 | **TRUE** | θ̂ = (XᵀX)⁻¹Xᵀy |
| 14 | **FALSE** | OLS has a **closed-form analytical** solution |
| 15 | **TRUE** | Need XᵀX invertible |
| 16 | **TRUE** | Invertible ⟺ full column rank |
| 17 | **TRUE** | Convex ⇒ stationary point is global min |
| 18 | **TRUE** | Collinearity ⇒ XᵀX singular |
| 19 | **TRUE** | Gradient = −2Xᵀ(y − Xθ) |
| 20 | **TRUE** | LLN: x̄ → E[X] |
| 21 | **TRUE** | Discrete expectation |
| 22 | **TRUE** | Continuous expectation |
| 23 | **TRUE** | Empirical risk estimates theoretical risk |
| 24 | **TRUE** | L1 loss g(x) = absolute value of x |
| 25 | **FALSE** | L1 is **not differentiable** at 0 (kink) |
| 26 | **TRUE** | L1 is more robust than L2 |
| 27 | **TRUE** | Optimal constant under L1 = median |
| 28 | **TRUE** | Gaussian noise ⟺ L2 ⟺ MLE = OLS |
| 29 | **TRUE** | Ridge restores invertibility |
| 30 | **TRUE** | First column of X is all ones |

---

## Score Interpretation

| Score | Meaning |
|-------|---------|
| 28–30 | Excellent — Chapter 5 mastered |
| 23–27 | Good — review the FALSE answers carefully |
| 18–22 | Needs work — re-read summary sections 4–7 |
| < 18 | Re-study the full chapter 5 summary |
