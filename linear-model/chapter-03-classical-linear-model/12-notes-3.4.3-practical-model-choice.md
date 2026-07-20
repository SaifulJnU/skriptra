# 3.4.3 — Practical Use of Model Choice Criteria

> Short section. You have a criterion (AIC, BIC…). But with $k$ covariates there are $2^k$ possible models. How do you actually search?

---

## 1. The combinatorial problem

With $k$ covariates, the number of possible subsets is $2^k$:

| $k$ | Number of models |
|---|---|
| 5 | 32 |
| 10 | 1,024 |
| 20 | 1,048,576 |
| 40 | $\approx1.1\times10^{12}$ |

**Complete enumeration** (fit all $2^k$, pick the best) is exact but only feasible for small $k$ — roughly $k\leq20$ with efficient algorithms.

For larger $k$ we use **stepwise** procedures: greedy searches that are fast but not guaranteed to find the global optimum.

---

## 2. Forward selection

1. Start with the **null model** (intercept only).
2. For each covariate not yet in the model, compute the criterion if it were added.
3. Add the one giving the **largest improvement**.
4. Repeat until no addition improves the criterion.

**Pros:** fast; works when $k>n$; starts simple.
**Cons:** a variable added early can't be removed later; misses variables that only matter **in combination** with another.

---

## 3. Backward elimination

1. Start with the **full model** (all covariates).
2. For each covariate, compute the criterion if it were removed.
3. Remove the one whose removal gives the **largest improvement**.
4. Repeat until no removal improves the criterion.

**Pros:** sees each variable in the context of all others, so it handles correlated groups better than forward.
**Cons:** requires $n>p$ to fit the full model at all; computationally heavier at the start.

---

## 4. Stepwise (hybrid)

At each step consider **both** adding and removing. This lets a variable that became redundant after later additions be dropped again. R's `step()` does this by default.

---

## 5. ⚠️ The warnings — this is the examinable part

Stepwise selection is standard practice and also widely criticised. If asked to discuss it, mention at least two of these:

| Problem | Explanation |
|---|---|
| **No global optimum** | Greedy search; forward and backward can end at **different** models |
| **Post-selection inference is invalid** | The p-values and CIs from the final model are computed **as if the model were chosen in advance**. They are too small / too narrow, because you selected on the same data |
| **Instability** | Small changes in the data can produce a completely different selected model, especially with correlated covariates |
| **Multiple testing** | Searching many models inflates the chance of including a covariate by luck |
| **Ignores subject knowledge** | A theoretically essential variable can be dropped for being marginally insignificant |

> **The strongest single sentence:** *"Because the model was selected using the same data used for inference, the reported p-values and confidence intervals of the final model are not valid — they understate uncertainty. Selection and inference should ideally use separate data, or the selection step should be accounted for explicitly."*

**Alternatives worth naming:** regularisation (ridge, lasso — Chapter 4, which does selection and estimation simultaneously and more stably), cross-validation over a small set of theory-driven candidate models, or simply fixing the model in advance on substantive grounds.

---

## 6. Practical guidance

| Situation | Suggested approach |
|---|---|
| $k$ small ($\leq15$), $n$ large | complete enumeration + BIC |
| $k$ moderate | backward elimination + AIC |
| $k>n$ | forward selection or lasso |
| Prediction is the goal | AIC or cross-validation |
| Identifying relevant variables | BIC |
| Strong subject-matter theory | fix the model in advance; don't search |

**Always report that a search was performed.** A model presented as if it were pre-specified, when it was actually selected from hundreds, is misleading — and this is a real and common failure in applied work, not just an exam nicety.

---

## 7. Key takeaways

1. $2^k$ candidate models — complete enumeration only for small $k$.
2. **Forward:** start empty, add. **Backward:** start full, remove. **Stepwise:** both.
3. Greedy searches give **no guarantee of the global optimum**; forward and backward can disagree.
4. 🔑 **Post-selection inference is invalid** — p-values from a selected model are too optimistic.
5. Stepwise selection is **unstable** under correlated covariates.
6. AIC/CV for prediction; BIC for identification; theory beats both when you have it.
