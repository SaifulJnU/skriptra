# Chapter 2: Classification — Practice True/False Questions

---

## Questions

**Q1.** In classification, the output y takes values in a discrete set {C_1,...,C_g} with g ≥ 2.

**Q2.** In the binary case (g = 2), the standard encodings used in this course are Y = {0,1} or Y = {−1,+1}.

**Q3.** In the multiclass case, the convention in this course is Y = {1, 2, ..., g} for g ≥ 3.

**Q4.** Classification models are defined to output discrete class labels directly, since that is the ultimate goal.

**Q5.** Scores/probabilities contain more information than class labels alone, which is why models output them instead of labels.

**Q6.** Once discrete class labels are produced, they can always be converted back to scores by applying an inverse transformation.

**Q7.** A scoring classifier uses g discriminant functions f_1,...,f_g and predicts the class with the maximum score.

**Q8.** In the binary case, a single discriminant function f(x) = f_1(x) − f_{−1}(x) is sufficient for a scoring classifier.

**Q9.** For a binary scoring classifier, the prediction rule is h(x) = sgn(f(x)).

**Q10.** The quantity |f(x)| in a binary scoring classifier is called the "confidence."

**Q11.** A probabilistic classifier constructs g probability functions that must sum to 1.

**Q12.** For a binary probabilistic classifier, the default threshold for class prediction is c = 0.

**Q13.** The holdout threshold for binary scoring classifiers (where h(x) = sgn(f(x))) is c = 0.

**Q14.** Probabilistic classifiers cannot be viewed as scoring classifiers.

**Q15.** Thresholding can convert probabilities to discrete class labels.

**Q16.** Thresholding can convert discrete class labels back to probabilities.

**Q17.** A decision boundary is a hypersurface that partitions the input space X into g decision regions.

**Q18.** Decision region X_k is defined as {x ∈ X : h(x) = k}.

**Q19.** The ties between decision regions are called decision boundaries.

**Q20.** In the binary case, the decision boundary for a probabilistic classifier is where f(x) = 0.5.

**Q21.** If the discriminant functions of a classifier can be written as g(f_k(x)) = w_k^T x + b_k via a monotone transformation, the classifier is called a linear classifier.

**Q22.** The decision boundary (tie) between two classes in a linear classifier is always a hyperplane.

**Q23.** A linear classifier can only produce linear decision boundaries in the original input space.

**Q24.** Using polynomial or basis function features, a linear classifier can represent non-linear boundaries in the original input space.

**Q25.** A sigmoid function is a bounded, differentiable function s : R → [0,1] with non-negative derivative everywhere.

**Q26.** The logistic function is defined as s(t) = 1 / (1 + e^t).

**Q27.** The logistic function satisfies: lim_{t→−∞} s(t) = 0 and lim_{t→+∞} s(t) = 1.

**Q28.** The derivative of the logistic function is ∂s(t)/∂t = s(t)(1 − s(t)).

**Q29.** The logistic function is symmetric about the point (0, 1/2).

**Q30.** The hyperbolic tangent (tanh) is an example of a sigmoid function.

**Q31.** In deep learning, sigmoid functions are used as activation functions.

**Q32.** The softmax function maps a g-dimensional real-valued score vector to a vector of g probabilities summing to 1.

**Q33.** The softmax function for class k is: π_k(x) = exp(f_k(x)) / Σ_j exp(f_j(x)).

**Q34.** For g = 2, the softmax reduces to the logistic function.

**Q35.** Unlike softmax, the argmax operator keeps information about non-maximal elements in a reversible way.

**Q36.** The generative approach to classification uses Bayes' theorem to compute π_k(x) = P(y=k|x).

**Q37.** In the generative approach, the discriminant functions are π_k(x) or log P(x|y=k) + log π_k.

**Q38.** LDA assumes each class has a multivariate Gaussian distribution with class-specific (unequal) covariances.

**Q39.** QDA assumes each class has a multivariate Gaussian distribution with equal covariances across all classes.

**Q40.** Naive Bayes assumes that features are conditionally independent given the class label y.

**Q41.** The conditional independence assumption in Naive Bayes means features are unconditionally independent of each other.

**Q42.** Logistic regression is a discriminant approach (not a generative one).

**Q43.** Discriminant approaches try to model the discriminant functions directly, often by loss minimization.

**Q44.** LDA produces a linear decision boundary because it assumes equal covariances across classes.

**Q45.** QDA produces a quadratic (non-linear) decision boundary because it allows unequal covariances.

---

## Answers

| Q | Answer | Key Reason |
|---|--------|------------|
| 1 | **TRUE** | Definition of classification |
| 2 | **TRUE** | Course convention for binary |
| 3 | **TRUE** | Course convention for multiclass |
| 4 | **FALSE** | Models output continuous scores/probabilities, not discrete labels |
| 5 | **TRUE** | Scores contain more information; labels → scores is irreversible |
| 6 | **FALSE** | Discrete classes CANNOT be converted back to scores |
| 7 | **TRUE** | Definition of scoring classifier |
| 8 | **TRUE** | Single f suffices for binary; uses {+1,−1} labels |
| 9 | **TRUE** | h(x) = sgn(f(x)) for binary scoring classifier |
| 10 | **TRUE** | |f(x)| = confidence |
| 11 | **TRUE** | π_1,...,π_g ∈ [0,1], Σ π_l = 1 |
| 12 | **FALSE** | Default threshold for **probabilistic** is c = **0.5** |
| 13 | **TRUE** | For scoring classifiers, threshold c = 0 (h = sgn(f)) |
| 14 | **FALSE** | Probabilistic classifiers CAN also be seen as scoring classifiers |
| 15 | **TRUE** | Thresholding: probabilities → discrete classes |
| 16 | **FALSE** | Discrete classes CANNOT be reverse-transformed to probabilities |
| 17 | **TRUE** | Definition of decision boundary |
| 18 | **TRUE** | Decision region definition |
| 19 | **TRUE** | Ties between regions = decision boundaries |
| 20 | **FALSE** | For binary probabilistic: decision boundary at π(x) = c (default c=0.5), NOT f(x) = 0.5 |
| 21 | **TRUE** | Definition of linear classifier |
| 22 | **TRUE** | (w_i − w_j)^T x + (b_i − b_j) = 0 is a hyperplane |
| 23 | **FALSE** | Can produce non-linear boundaries with derived/polynomial features |
| 24 | **TRUE** | Feature engineering allows non-linear boundaries |
| 25 | **TRUE** | Definition of sigmoid |
| 26 | **FALSE** | Logistic is s(t) = 1/(1 + e^{**−**t}), not e^t |
| 27 | **TRUE** | Standard limits of logistic function |
| 28 | **TRUE** | Key property: ∂s/∂t = s(t)(1−s(t)) |
| 29 | **TRUE** | Symmetric about (0, 1/2) |
| 30 | **TRUE** | tanh is listed as a sigmoid example |
| 31 | **TRUE** | Sigmoids as activation functions in deep learning |
| 32 | **TRUE** | Softmax maps R^g → Δ^g (probability simplex) |
| 33 | **TRUE** | Softmax formula |
| 34 | **TRUE** | Softmax is a generalization; reduces to logistic for g=2 |
| 35 | **FALSE** | It is **soft**max (not argmax) that keeps non-maximal info reversibly; argmax discards it |
| 36 | **TRUE** | Generative approach uses Bayes' theorem |
| 37 | **TRUE** | Discriminant functions in generative approach |
| 38 | **FALSE** | LDA has **equal** covariances Σ (same for all classes); QDA has unequal |
| 39 | **FALSE** | QDA has **unequal** covariances Σ_k; LDA has equal covariances |
| 40 | **TRUE** | Naive Bayes conditional independence assumption |
| 41 | **FALSE** | Conditional independence given y ≠ unconditional independence |
| 42 | **TRUE** | Logistic regression is discriminant |
| 43 | **TRUE** | Definition of discriminant approach |
| 44 | **TRUE** | Equal Σ → cancellation → linear boundary |
| 45 | **TRUE** | Unequal Σ_k → quadratic terms remain |

---

## Score Interpretation

| Score | Meaning |
|-------|---------|
| 40–45 | Excellent — Chapter 2 mastered |
| 34–39 | Good — review the FALSE answers carefully |
| 27–33 | Needs work — re-read summary sections 5–10 |
| < 27 | Re-study the full chapter 2 summary |
