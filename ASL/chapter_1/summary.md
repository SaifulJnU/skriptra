# Chapter 1: Introduction and Formalization — Summary

## 1. What is Machine Learning?

> "A computer program is said to **learn** from experience E with respect to some task T and some performance measure P, if its performance on T, as measured by P, improves with experience E."  
> — Tom Mitchell, 1998

**Hierarchy:** Artificial Intelligence ⊃ Machine Learning ⊃ Deep Learning

---

## 2. The Data

Given a dataset:

$$\mathcal{D} = \{(\mathbf{x}^{(1)}, y^{(1)}), \ldots, (\mathbf{x}^{(n)}, y^{(n)})\} \subset (\mathcal{X} \times \mathcal{Y})^n$$

- **X** = input space, dim(X) = p (usually X ⊂ R^p)
- **Y** = output / target space
- **(x^(i), y^(i))** = the i-th observation
- **x_j** = the j-th feature vector across all observations

---

## 3. Data Generating Process

- We assume a joint probability distribution **P_xy** on X × Y
- **P_xy is typically unknown and very complicated** in practice
- Data is assumed drawn **i.i.d.** from p(x, y)
- Often parameterized: p(x, y | θ) with θ ∈ Θ
- Note: `p(x|θ)` is frequentist notation — the `|` is for readability, NOT Bayesian conditioning

---

## 4. Supervised Task Types

Three types based on the output space Y:

| Task | Output Y | Notes |
|------|----------|-------|
| **Regression** | Y ⊂ R^g, 1 ≤ g < ∞ | g=1 → univariate; g>1 → multi-target |
| **Classification** | Y = {C_1,...,C_g}, g ≥ 2 | g=2 → binary; g>2 → multiclass |
| **Density estimation** | predict p(y\|x) on Y | — |

### Regression details
- Predict a **continuous/metric** output y ∈ R
- f(x) = prediction = **score**
- **Residual** r = y − f(x) or r^(i) = y^(i) − f(x^(i))

### Classification details
- Predict class y ∈ {C_1,...,C_g}
- Goal: predict class labels OR class membership probabilities π_1,...,π_g ∈ [0,1]

---

## 5. Supervised Models

A model is a function:

$$f : \mathcal{X} \to \mathbb{R}^g$$

- Output = **scores** (g real values per input)
- For regression: scores are taken as predictions directly
- For classification: scores are transformed to classes/probabilities
- **ŷ := f(x)** is the (model) prediction
- **H** = hypothesis space = the space of all allowed models

---

## 6. Inducing Algorithm

$$\mathcal{I}_{L,O} : (\mathcal{X} \times \mathcal{Y})^n \to \mathcal{H}$$

- Maps dataset D to a model f ∈ H
- **H** = representation / hypothesis space
- **L : Y × R^p → R≥0** = loss function (measures cost/risk)
- **O** = optimizer

Running I is called **training** or **fitting**. The fitted model is written f̂(x) = I_{L,O}(D).

### The Three Components (Domingos 2012)

> **Learning = Representation + Cost function + Optimization**

| Component | Role | Examples |
|-----------|------|---------|
| **Representation** | What models can be learned (H) | Neighbors, Linear functions, Decision trees, Neural networks, Graphical models |
| **Cost function** | Distinguishes good from bad models | Squared error, Misclassification, Likelihood, Information gain |
| **Optimization** | Efficiently searches H | Gradient descent, Quadratic programming, Combinatorial optimization, Genetic algorithms |

### Example: Linear Regression
- **Representation:** H = {f(x) = θ^T x̃ | θ ∈ R^{p+1}}, where x̃ = (1, x)^T
- **Cost:** SSE = Σ(y^(i) − f(x^(i)))² = ||y − Xθ||²
- **Optimization:** SSE minimized **analytically** via derivation w.r.t. θ (no numerical optimization needed)

---

## 7. Generalization

**The fundamental goal of ML is to generalize beyond training data.**

### Generalization Error of a fitted model f̂ = I_{L,O}(D):

$$GE(\hat{f}) = \mathbb{E}_{(\mathbf{x},y) \sim \mathbb{P}_{xy}} \left[ L(y, \hat{f}(\mathbf{x})) \right]$$

- Training data D (and thus f̂) is **fixed**
- Test observations (x, y) ~ P_xy are random
- **Cannot be computed exactly** because P_xy is unknown

### Inner Loss vs Outer Loss

| | Inner Loss | Outer Loss |
|--|-----------|-----------|
| **When** | During model fitting | When assessing performance afterwards |
| **Purpose** | Optimized to find f̂ | Measures true generalization |
| **Relationship** | Desired to match outer loss | Often given by application |
| **Issue** | Outer loss can be hard to optimize numerically | — |

**Example:** Logistic regression minimizes binomial loss (inner), but we may evaluate by AUC or misclassification rate (outer).

---

## 8. Estimating Generalization Error

### Train Error (biased estimator)

$$\widehat{GE}_{\mathcal{D}}(\hat{f}) = \frac{1}{|\mathcal{D}|} \sum_{(\mathbf{x},y) \in \mathcal{D}} L(y, \hat{f}(\mathbf{x}))$$

- **Overly optimistic (biased)** — model was fitted on this data
- Training error **underestimates** true generalization error

### Holdout (Test Error)

Split D = D_train ∪ D_test, fit on D_train, evaluate on D_test:

$$\widehat{GE}_{\mathcal{D}_{test}}(\hat{f}) = \frac{1}{|\mathcal{D}_{test}|} \sum_{(\mathbf{x},y) \in \mathcal{D}_{test}} L(y, \hat{f}(\mathbf{x}))$$

- This procedure is called **holdout**
- Variance can be reduced with **iterated holdout** or **cross-validation**

### Key Relationship: Train vs Test Error

> **Test error ≥ Training error** (typically)

- If hypothesis is fixed before seeing data: expected training error = expected test error
- But in ML: the inducer **minimizes error on training data** → training error is optimistic
- Therefore: **training error is a biased (optimistic) estimate of true performance**

### Generalization Error of a Learning Algorithm

$$GE_n(\mathcal{I}_{L,O}) = \mathbb{E}_{\mathcal{D}_n \sim \mathbb{P}_{xy}^n, (\mathbf{x},y) \sim \mathbb{P}_{xy}} \left[ L\left(y, \hat{f}_{\mathcal{D}_n}(\mathbf{x})\right) \right]$$

- Both training data D_n and test observations (x,y) are **random variables**
- Expectation taken over both D_n ~ P_xy^{i.i.d.} and (x,y) ~ P_xy

---

## 9. ML Synonyms (Know These!)

| Group | Synonyms |
|-------|----------|
| Algorithm | inducer, inducing algorithm, learning algorithm, learner |
| Process | learning, training, inducing, fitting |
| Data point | example, instance, observation |
| Input variable | feature, attribute, covariate, input variable, predictor |
| Output variable | output, target, dependent variable, outcome, response |
| Loss | cost function, costs, risk |
| Class | label, class, categories |

---

## Quick-Check: Common True/False Traps

- The test error is typically **greater than or equal to** the training error (NOT less than)
- Training error is an **optimistic** estimate (NOT pessimistic)
- **Density estimation** IS a supervised task
- The inducer maps D → f ∈ H (NOT D → H)
- For regression g=1 means **univariate** response (g>1 is multi-target)
- **Binary classification**: g=2; **Multiclass**: g>2 (g≥3)
- The `|` in p(x|θ) is frequentist notation for readability, **NOT** Bayesian conditioning (unless explicitly stated)
- The outer loss is used to **assess** performance (NOT to train the model)
- SSE in linear regression can be minimized **analytically** — no numerical optimization needed
