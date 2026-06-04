# Chapter 4: Risk Minimization I — Practice True/False Questions

---

## Questions

**Q1.** Learning can be decomposed into three components: representation, cost, and optimization.

**Q2.** The cost (loss) function is what distinguishes good models from bad ones.

**Q3.** Given a fixed hypothesis space of linear models, the model returned by learning is always the same regardless of the cost function.

**Q4.** A loss function maps L : Y × ℝᵍ → ℝ≥0 and must be non-negative.

**Q5.** A loss function satisfies L(y, ỹ) = 0 if and only if y = ỹ.

**Q6.** The point-wise squared error loss is L(y, f(x)) = (y − f(x))².

**Q7.** The residual is defined as r = y − f(x).

**Q8.** A loss is called distance-based if it can be written purely in terms of the residual.

**Q9.** A loss is translation-invariant if L(y + a, f(x) + a) = L(y, f(x)) for all a.

**Q10.** A loss is translation-invariant if and only if it is distance-based.

**Q11.** The pseudo-residual is the positive first derivative of the loss with respect to f.

**Q12.** For the L2-loss, the pseudo-residuals coincide with the ordinary residuals.

**Q13.** In a loss plot, the pseudo-residual corresponds to the slope of the tangent at the given residual.

**Q14.** The theoretical risk R(f) is the expected loss E[L(y, f(x))] over (x, y) ~ P_xy.

**Q15.** For the L2-loss with an unrestricted hypothesis space, the optimal prediction is the conditional expectation E[y | x].

**Q16.** The identity E[(y − c)²] = Var[y] + (E[y] − c)² shows the risk is minimized at c = E[y].

**Q17.** Minimizing the theoretical risk R(f) directly is generally easy because P_xy is known.

**Q18.** Estimating P_xy non-parametrically scales well to high dimensions.

**Q19.** The empirical risk R_emp(f) = Σ L(y⁽ⁱ⁾, f(x⁽ⁱ⁾)) approximates the theoretical risk using the data.

**Q20.** The factor 1/n in the averaged empirical risk changes the location of the minimizer.

**Q21.** R_emp is a good approximation of R only if the data is an unbiased, independent, and large enough sample.

**Q22.** Empirical risk minimization means learning is (often) an optimization problem.

**Q23.** If the loss is convex in its second argument and f is linear in θ, then every local minimum of R_emp(θ) is also global.

**Q24.** If the loss is non-convex, R_emp(θ) is guaranteed to have a single global minimum.

**Q25.** The L2-loss is sensitive to outliers because doubling a residual quadruples the loss.

**Q26.** The L2-loss is convex and differentiable.

**Q27.** The optimal constant model under the L2-loss is the median of the observed outcomes.

**Q28.** The optimal constant model under the L2-loss is the mean (average) of the observed outcomes.

**Q29.** The L1-loss is more robust to outliers than the L2-loss.

**Q30.** The L1-loss is differentiable everywhere, including at y = f(x).

**Q31.** The optimal constant model under the L1-loss is the median of the observed outcomes.

**Q32.** The quantile (pinball) loss is an extension of the L1-loss, and equals L1 when α = 0.5.

**Q33.** The optimal constant model under the quantile loss is the empirical α-quantile of the data.

**Q34.** A quantile loss with α > 0.5 penalizes under-estimation more than over-estimation.

**Q35.** The Huber loss is a piecewise combination of L1 and L2 loss.

**Q36.** The Huber loss is both differentiable and robust.

**Q37.** The Huber loss has a closed-form optimal constant model.

**Q38.** The constant model fitted with the Huber loss lies between the L1 and L2 solutions.

**Q39.** The ε-insensitive loss accepts errors below ε without any penalty.

**Q40.** The log-barrier loss guarantees the risk-minimization problem always has a solution.

**Q41.** Gradient descent updates parameters in the direction of the positive gradient.

**Q42.** The negative gradient is the direction of steepest descent.

**Q43.** The step-size α in gradient descent is also called the learning rate.

**Q44.** Stochastic gradient descent (SGD) approximates the gradient using a single random observation.

**Q45.** Mini-batch gradient descent uses all n observations to compute the gradient.

**Q46.** SGD is computationally cheaper than full gradient descent but noisier.

**Q47.** Mini-batch gradient descent is a trade-off between full gradient descent and SGD.

**Q48.** Under maximum likelihood, defining the loss as L = −log p(y | x, θ) makes the MLE identical to the loss-minimal estimator.

**Q49.** For every loss function, there exists a corresponding error distribution.

**Q50.** The hinge loss is an example of a loss that does NOT correspond to any error distribution.

**Q51.** Assuming Gaussian errors leads to the L2-loss via the negative log-likelihood.

**Q52.** Assuming Laplace-distributed errors leads to the L1-loss.

**Q53.** Multiplicative and additive constants in the loss change the minimizing parameter.

**Q54.** A loss function intuitively defines how the residuals will be distributed after fitting.

---

## Answers

| Q | Answer | Key Reason |
|---|--------|------------|
| 1 | **TRUE** | Learning = Representation + Cost + Optimization |
| 2 | **TRUE** | Cost/loss distinguishes good vs bad models |
| 3 | **FALSE** | The returned model **depends** on the cost function |
| 4 | **TRUE** | L : Y × ℝᵍ → ℝ≥0, non-negative |
| 5 | **TRUE** | Optimality requirement |
| 6 | **TRUE** | Point-wise squared error |
| 7 | **TRUE** | r = y − f(x) |
| 8 | **TRUE** | Distance-based = function of residual |
| 9 | **TRUE** | Translation-invariance definition |
| 10 | **TRUE** | Translation-invariant ⟺ distance-based |
| 11 | **FALSE** | **Negative** first derivative |
| 12 | **TRUE** | For L2, pseudo-residuals = residuals |
| 13 | **TRUE** | Pseudo-residual = slope of tangent |
| 14 | **TRUE** | Risk = expected loss |
| 15 | **TRUE** | L2 optimal = conditional expectation |
| 16 | **TRUE** | Minimal at c = E[y] |
| 17 | **FALSE** | P_xy is **unknown** → not feasible |
| 18 | **FALSE** | Does **not** scale (curse of dimensionality) |
| 19 | **TRUE** | Empirical risk definition |
| 20 | **FALSE** | 1/n does **not** change the minimizer |
| 21 | **TRUE** | Requires unbiased, independent, large sample |
| 22 | **TRUE** | ERM ⇒ optimization problem |
| 23 | **TRUE** | Convex L + linear f → local min = global |
| 24 | **FALSE** | Non-convex → possibly **multiple** local minima |
| 25 | **TRUE** | Residual ×2 → loss ×4 |
| 26 | **TRUE** | L2 is convex & differentiable |
| 27 | **FALSE** | L2 optimal constant = **mean**, not median |
| 28 | **TRUE** | Optimal constant = mean ȳ |
| 29 | **TRUE** | L1 is more robust |
| 30 | **FALSE** | L1 not differentiable at y = f(x) |
| 31 | **TRUE** | L1 optimal constant = median |
| 32 | **TRUE** | Quantile extends L1; α=0.5 → L1 |
| 33 | **TRUE** | Optimal constant = empirical α-quantile |
| 34 | **TRUE** | α > 0.5 penalizes under-estimation |
| 35 | **TRUE** | Huber = piecewise L1 + L2 |
| 36 | **TRUE** | Huber: differentiable + robust |
| 37 | **FALSE** | Huber has **no** closed form (numerical) |
| 38 | **TRUE** | Huber solution lies between L1 and L2 |
| 39 | **TRUE** | ε-insensitive: errors below ε free |
| 40 | **FALSE** | **No guarantee** of a solution |
| 41 | **FALSE** | **Negative** gradient direction |
| 42 | **TRUE** | Negative gradient = steepest descent |
| 43 | **TRUE** | Step-size α = learning rate |
| 44 | **TRUE** | SGD uses one random observation |
| 45 | **FALSE** | Mini-batch uses a **subset** I ⊂ {1,…,n} |
| 46 | **TRUE** | SGD: cheap but noisy |
| 47 | **TRUE** | Mini-batch = trade-off |
| 48 | **TRUE** | MLE = loss-minimal estimator |
| 49 | **FALSE** | NOT every loss has an error distribution (hinge) |
| 50 | **TRUE** | Hinge loss = counterexample |
| 51 | **TRUE** | Gaussian errors ⟺ L2-loss |
| 52 | **TRUE** | Laplace errors ⟺ L1-loss |
| 53 | **FALSE** | Constants do **not** change the minimizer |
| 54 | **TRUE** | Loss defines residual distribution |

---

## Score Interpretation

| Score | Meaning |
|-------|---------|
| 49–54 | Excellent — Chapter 4 mastered |
| 41–48 | Good — review the FALSE answers carefully |
| 32–40 | Needs work — re-read summary sections 5–8 |
| < 32 | Re-study the full chapter 4 summary |
