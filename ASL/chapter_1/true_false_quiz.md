# Chapter 1: Practice True/False Questions

**Instructions:** Answer each statement TRUE or FALSE, then check against the answers below.  
Quiz format: 30 statements in 30 minutes — aim to answer each in under 60 seconds.

---

## Questions

**Q1.** A computer program is said to learn if its performance on task T improves with experience E, as measured by performance measure P.

**Q2.** Deep Learning is a subset of Machine Learning, which is itself a subset of Artificial Intelligence.

**Q3.** In supervised learning, the dataset D consists of input-output pairs (x^(i), y^(i)) drawn from X × Y.

**Q4.** The joint probability distribution P_xy is typically known and easy to compute in practice.

**Q5.** Data in supervised learning is assumed to be drawn i.i.d. from the joint distribution p(x, y).

**Q6.** Writing p(x|θ) implies a Bayesian treatment where we condition on θ as a random variable.

**Q7.** Multi-target regression refers to the case where g > 1 in Y ⊂ R^g.

**Q8.** Binary classification is a special case of classification where g = 2.

**Q9.** Multiclass classification requires g > 2 classes.

**Q10.** Density estimation is not considered a supervised learning task.

**Q11.** In regression, the prediction f(x) is also referred to as a score.

**Q12.** The residual is defined as r = f(x) − y.

**Q13.** A model f : X → R^g maps inputs to scores (real values), not directly to class labels.

**Q14.** The hypothesis space H contains all possible datasets D.

**Q15.** The inducing algorithm maps a dataset D ∈ (X×Y)^n to a model f ∈ H.

**Q16.** The loss function L : Y × R^p → R≥0 can output negative values.

**Q17.** According to Domingos (2012), Learning = Representation + Cost function + Optimization.

**Q18.** The hypothesis space defines the set of classifiers the model can possibly learn.

**Q19.** In linear regression, the SSE can be minimized analytically without numerical optimization.

**Q20.** The SSE for linear regression is defined as SSE = Σ(y^(i) − f(x^(i)))² = ||y − Xθ||².

**Q21.** The fundamental goal of machine learning is to memorize the training data as accurately as possible.

**Q22.** The generalization error GE(f̂) is the expected loss over new observations drawn from P_xy.

**Q23.** The generalization error of a fitted model can always be computed exactly.

**Q24.** In the expression for generalization error, the training data D (and thus f̂) is fixed, while the test observations are random.

**Q25.** The inner loss is used to assess model performance after training.

**Q26.** The outer loss is the loss that is optimized during model fitting.

**Q27.** It is always possible to use the outer loss as the inner loss during optimization.

**Q28.** The training error is an optimistic (biased) estimate of the generalization error.

**Q29.** The test error is typically less than or equal to the training error.

**Q30.** The holdout method involves training on D_train and evaluating on D_test.

**Q31.** If a hypothesis is fixed before seeing the data, the expected training error equals the expected test error.

**Q32.** In machine learning, because the inducer minimizes error on training data, the training error typically underestimates the true generalization error.

**Q33.** In the generalization error of a learning algorithm, both the training data and test observations are treated as random variables.

**Q34.** "Feature" and "predictor" refer to different things in the ML context.

**Q35.** "Cost function", "loss", and "risk" are used interchangeably in this course.

**Q36.** The terms "learning", "training", "inducing", and "fitting" are synonymous in this course.

**Q37.** The optimizer O in the inducing algorithm is part of the hypothesis space.

**Q38.** Cross-validation is a more systematic resampling approach than the basic holdout method.

**Q39.** Gradient descent, quadratic programming, and genetic algorithms are all examples of optimization methods used in ML.

**Q40.** The design matrix X in linear regression has dimensions n × (p+1) when an intercept is included.

---

## Answers

| Q | Answer | Key Reason |
|---|--------|------------|
| 1 | **TRUE** | Exact Mitchell definition |
| 2 | **TRUE** | DL ⊂ ML ⊂ AI |
| 3 | **TRUE** | Definition of D |
| 4 | **FALSE** | P_xy is "typically unknown and very complicated" |
| 5 | **TRUE** | Standard i.i.d. assumption |
| 6 | **FALSE** | Frequentist notation for readability; does NOT imply Bayesian treatment |
| 7 | **TRUE** | Multi-target = g > 1 |
| 8 | **TRUE** | Binary: g = 2 |
| 9 | **TRUE** | Multiclass: g > 2 |
| 10 | **FALSE** | Density estimation IS listed as a supervised task |
| 11 | **TRUE** | "A prediction f(x) is also referred to as score" |
| 12 | **FALSE** | Residual r = y − f(x), NOT f(x) − y |
| 13 | **TRUE** | Models output scores/probabilities, not discrete labels |
| 14 | **FALSE** | H contains all possible models, not datasets |
| 15 | **TRUE** | Definition of inducing algorithm |
| 16 | **FALSE** | L maps to R≥0 (non-negative) |
| 17 | **TRUE** | Domingos 2012 quote |
| 18 | **TRUE** | Definition of hypothesis space / representation |
| 19 | **TRUE** | SSE has an analytical (closed-form) solution via OLS |
| 20 | **TRUE** | SSE = ||y − Xθ||² |
| 21 | **FALSE** | Goal is to **generalize** beyond training data |
| 22 | **TRUE** | GE(f̂) = E_{(x,y)~P_xy}[L(y, f̂(x))] |
| 23 | **FALSE** | Cannot be computed exactly because P_xy is usually unknown |
| 24 | **TRUE** | f̂ is fixed (trained); test observations are random |
| 25 | **FALSE** | Inner loss is used **during** training (fitting), not after |
| 26 | **FALSE** | **Outer** loss assesses performance; **inner** loss is optimized |
| 27 | **FALSE** | Outer loss is often numerically hard to optimize; approximations needed |
| 28 | **TRUE** | Inducer minimizes on training data → optimistic |
| 29 | **FALSE** | Test error is typically **greater than or equal to** training error |
| 30 | **TRUE** | Definition of holdout |
| 31 | **TRUE** | Only when hypothesis is fixed in advance (pre-data) |
| 32 | **TRUE** | Key result from the lecture |
| 33 | **TRUE** | Both D_n and (x,y) are random in the algorithm-level GE |
| 34 | **FALSE** | "Feature", "attribute", "covariate", "predictor" are synonyms |
| 35 | **TRUE** | All listed as synonyms in the course |
| 36 | **TRUE** | All listed as synonyms |
| 37 | **FALSE** | O (optimizer) is separate from H (hypothesis space) |
| 38 | **TRUE** | Cross-validation is more systematic than basic holdout |
| 39 | **TRUE** | All listed as optimization methods |
| 40 | **TRUE** | n observations × (p features + 1 intercept column) |

---

## Score Interpretation

| Score | Meaning |
|-------|---------|
| 36–40 | Excellent — Chapter 1 mastered |
| 30–35 | Good — review the FALSE answers carefully |
| 24–29 | Needs work — re-read summary, focus on definitions |
| < 24 | Re-study the full chapter summary before re-attempting |
