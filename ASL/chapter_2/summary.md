# Chapter 2: Classification — Summary

## 1. Classification Task Setup

Predict a discrete output:

$$y \in \mathcal{Y} = \{C_1, \ldots, C_g\}, \quad 2 \leq g < \infty$$

**Encoding conventions in this course:**
- Binary (g = 2): Y = {0, 1} or Y = {−1, +1}
- Multiclass (g ≥ 3): Y = {1, 2, ..., g}

---

## 2. Why Models Output Scores, Not Classes

Models are defined as f : X → R^g outputting **scores** (not discrete classes). Why?

1. **Optimization is easier** on continuous-valued functions
2. **Scores/probabilities contain more information** than labels alone
3. Scores → classes is easy (thresholding); classes → scores is **not possible** (irreversible)

We distinguish two types of classifiers: **scoring** and **probabilistic**.

---

## 3. Scoring Classifiers

Construct g discriminant/scoring functions f_1, ..., f_g : X → R.

**Prediction rule (multiclass):**

$$h(\mathbf{x}) = \underset{k \in \{1,2,\ldots,g\}}{\arg\max} \; f_k(\mathbf{x})$$

**Binary case (g = 2):**  
A single discriminant function f(x) = f_1(x) − f_{−1}(x) is sufficient.

$$h(\mathbf{x}) = 1 \iff f_1(\mathbf{x}) > f_{-1}(\mathbf{x}) \iff f(\mathbf{x}) > 0$$

- **h(x) = sgn(f(x))**
- **|f(x)|** is called the **confidence**

---

## 4. Probabilistic Classifiers

Construct g probability functions π_1, ..., π_g : X → [0, 1] with Σ π_l = 1.

**Prediction rule (multiclass):**

$$h(\mathbf{x}) = \underset{k \in \{1,2,\ldots,g\}}{\arg\max} \; \pi_k(\mathbf{x})$$

**Binary case (g = 2):**  
Usually a single function π(x) is output. Thresholding with c ∈ [0, 1]:

$$h(\mathbf{x}) := \mathbb{1}(\pi(\mathbf{x}) \geq c)$$

- Default threshold: **c = 0.5**
- Probabilistic classifiers can **also be seen as** scoring classifiers

---

## 5. Relationship: Probabilities ↔ Scores ↔ Discrete Classes

```
                  Calibrating/Scaling
Probabilities <─────────────────────── Scores
     │                                    ▲
     │ Thresholding              Thresholding
     ▼                                    │
Discrete Classes ───────────────────── (intrinsic)
```

Key rules:
- **Probabilities → Discrete Classes:** thresholding ✓
- **Scores → Discrete Classes:** thresholding ✓
- **Scores → Probabilities:** calibrating/scaling ✓
- **Discrete Classes → Scores or Probabilities:** NOT possible (irreversible!)

> Discrete classes are often intrinsically produced by scores, but can NOT be transferred back.

---

## 6. Decision Boundaries

A **decision boundary** is a hypersurface that partitions X into g **decision regions**:

$$\mathcal{X}_k = \{\mathbf{x} \in \mathcal{X} : h(\mathbf{x}) = k\}$$

The **ties** between regions are the decision boundaries:

**Multiclass general case:**
$$\{\mathbf{x} \in \mathcal{X} : \exists i \neq j \text{ s.t. } f_i(\mathbf{x}) = f_j(\mathbf{x}) \text{ and } f_i(\mathbf{x}), f_j(\mathbf{x}) \geq f_k(\mathbf{x}) \; \forall k \neq i,j\}$$

**Binary case:**
$$f(\mathbf{x}) = c$$

where c is the threshold (c = 0 for scoring classifiers, c = 0.5 for probabilistic classifiers).

---

## 7. Linear Classifiers

If discriminant functions can be written as (possibly via monotone transformation g: R → R):

$$g(f_k(\mathbf{x})) = \mathbf{w}_k^\top \mathbf{x} + b_k$$

then the classifier is a **linear classifier**.

The tie between two classes i and j becomes:

$$f_i(\mathbf{x}) = f_j(\mathbf{x}) \iff (\mathbf{w}_i - \mathbf{w}_j)^\top \mathbf{x} + (b_i - b_j) = 0$$

This is a **hyperplane** separating the two classes (with w_ij = w_i − w_j, b_ij = b_i − b_j).

> **Important:** Linear classifiers can produce **non-linear decision boundaries** in the **original input space** if derived features (polynomial, basis function expansions) are used.

---

## 8. Sigmoid Functions (Binary Scaling)

Any scoring model can become a probabilistic model using a sigmoid transformation:

$$\pi(\mathbf{x}) := s(f(\mathbf{x})) \in [0,1]$$

A **sigmoid** is a bounded, differentiable, real-valued function s : R → [0,1] with non-negative derivative.

| Sigmoid | Formula |
|---------|---------|
| Arctan | s(t) = arctan(t) |
| Hyperbolic tangent | s(t) = tanh(t) = (e^t − e^{−t}) / (e^t + e^{−t}) |
| **Logistic** | s(t) = 1 / (1 + e^{−t}) |
| Probit | s(t) = CDF of normal distribution |

### The Logistic Function (most important)

$$s(t) = \frac{1}{1 + e^{-t}}$$

Properties:
- lim_{t→−∞} s(t) = 0 and lim_{t→+∞} s(t) = 1
- Derivative: ∂s(t)/∂t = s(t)(1 − s(t))
- **Symmetrical about the point (0, 1/2)**
- Used in **logistic regression**
- In deep learning, sigmoids are used as **activation functions**

---

## 9. Softmax Function (Multiclass Scaling)

Generalizes the logistic to multiclass. Maps scores (f_1(x), ..., f_g(x)) to probabilities:

$$\pi_k(\mathbf{x}) = \frac{\exp(f_k(\mathbf{x}))}{\sum_{j=1}^g \exp(f_j(\mathbf{x}))}$$

Properties:
- π_k(x) ∈ [0, 1] for all k
- Σ π_k(x) = 1
- **Generalizes logistic**: for g = 2, softmax reduces to logistic function
- "Squashes" a g-dimensional vector to the same dimension with entries in [0,1] summing to 1
- Compared to argmax: **soft**max keeps information about non-maximal elements in a reversible way

---

## 10. Generative vs. Discriminative Approaches

### Generative Approach
Uses Bayes' theorem:

$$\pi_k(\mathbf{x}) = \mathbb{P}(y=k \mid \mathbf{x}) = \frac{\mathbb{P}(\mathbf{x} \mid y=k) \, \mathbb{P}(y=k)}{\mathbb{P}(\mathbf{x})} \propto \mathbb{P}(\mathbf{x} \mid y=k) \, \pi_k$$

Models P(x | y = k) to compute π_k(x).  
Discriminant functions: **π_k(x) or log P(x|y=k) + log π_k**

| Model | Type | Decision Boundary | Distribution Assumption |
|-------|------|-------------------|------------------------|
| **LDA** | Generative, Linear | Linear hyperplane | Multivariate Gaussian, **equal covariances** Σ |
| **QDA** | Generative, Non-linear | Quadratic | Multivariate Gaussian, **unequal covariances** Σ_k |
| **Naive Bayes** | Generative, Non-linear | Non-linear | Conditional independence of features given class |
| **Logistic Regression** | Discriminant, Linear | Linear | — (models discriminant function directly) |

### LDA (Linear Discriminant Analysis)
$$\mathbb{P}(\mathbf{x} \mid y=k) \sim \mathcal{N}(\boldsymbol{\mu}_k, \boldsymbol{\Sigma})$$
with **equal covariances** across all classes.

### QDA (Quadratic Discriminant Analysis)
$$\mathbb{P}(\mathbf{x} \mid y=k) \sim \mathcal{N}(\boldsymbol{\mu}_k, \boldsymbol{\Sigma}_k)$$
with **unequal covariances** Σ_k per class.

### Naive Bayes
Assumes features are **conditionally independent** given the class:

$$\mathbb{P}(\mathbf{x} \mid y=k) = \prod_{j=1}^p \mathbb{P}(x_j \mid y=k)$$

### Discriminant (Discriminative) Approaches
- Try to model the discriminant functions **directly**, often by loss minimization
- Example: Logistic Regression

---

## Quick-Check: Common True/False Traps

- Models output **scores**, not classes (by design — easier to optimize)
- Default threshold for binary probabilistic classifiers: **c = 0.5**
- Default threshold for scoring classifiers: **c = 0** (since h(x) = sgn(f(x)))
- **Discrete classes CANNOT be converted back to scores** (one-way)
- Linear classifiers can produce **non-linear** decision boundaries if feature engineering is used
- The softmax is a **generalization** of the logistic (not a completely different function)
- **LDA = equal covariances**, **QDA = unequal covariances** (easy to mix up!)
- **Naive Bayes** assumes conditional independence of features given class (NOT unconditional independence)
- Logistic regression is **discriminant** (not generative)
- |f(x)| in binary scoring classifiers is called **confidence**
- The softmax keeps non-maximal elements **in a reversible way** (unlike argmax)
- Probabilistic classifiers CAN also be seen as scoring classifiers
