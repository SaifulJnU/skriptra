# Chapter 3: Hypothesis Spaces and Capacity — Summary

## 1. Capacity, Underfitting, Overfitting

The performance of a learner depends on its ability to:
1. **Minimize the training error**
2. **Generalize** well to new (unseen) data

- **Underfitting:** failure to obtain a sufficiently low training error.
- **Overfitting:** a large difference between training and test error.
- **Generalization gap:** the gap between training error and generalization (test) error.

The tendency to over/underfit is a function of the model's **capacity** — determined by the type of hypotheses it can learn.

- **Low capacity:** can only learn a few simple hypotheses → tends to underfit.
- **High capacity:** can learn many, possibly complex hypotheses → tends to overfit.
- **Optimal capacity:** test error is minimized when the model neither underfits nor overfits.

```
Error
  |\                          Generalization error
  | \                        /
  |  \    Underfit | Overfit/
  |   \___________ | ______/      <- generalization gap
  |               \|_____/  ____ Training error
  |________________|____________________ Capacity
              Optimal Capacity
```

---

## 2. Hypothesis Spaces

The **representation** of a learner is the space of allowed models, also called the **hypothesis space**. The hypothesis space $\mathcal{H}$ is a space of functions with a certain functional form:

$$\mathcal{H} := \{ f : \mathcal{X} \to \mathbb{R}^g \mid f \text{ has a specific form} \}$$

Often $f$ is parametrized by $\boldsymbol\theta \in \Theta$, written $f(\mathbf{x}) = f(\mathbf{x} \mid \boldsymbol\theta)$.

> **Note:** When explicitly talking about hard classifiers outputting a discrete class we write $h$ instead of $f$. The generic symbol $f$ subsumes discrete classes, scores, and probabilities.

### Examples of hypothesis spaces

| Model | Functional form |
|-------|-----------------|
| **Linear regression** | $f(\mathbf{x} \mid \theta_0, \boldsymbol\theta) = \mathbf{x}^\top\boldsymbol\theta + \theta_0$, with $\boldsymbol\theta \in \mathbb{R}^p, \theta_0 \in \mathbb{R}$ |
| **Separating hyperplanes** | $h(\mathbf{x} \mid \theta_0, \boldsymbol\theta) = \mathbb{1}(\mathbf{x}^\top\boldsymbol\theta - \theta_0 > 0)$ |
| **Decision trees** | $f(\mathbf{x}) = \sum_{i=1}^m c_i \mathbb{1}(\mathbf{x} \in Q_i)$ — recursively divides feature space into axis-aligned rectangles |
| **Ensemble methods** | $f(\mathbf{x} \mid \beta^{[l]}) = \sum_{l=1}^m \beta^{[l]} b^{[l]}(\mathbf{x})$ — aggregates several models (e.g. random forests, bagging, tree-based boosting) |
| **Neural networks** | $f(\mathbf{x}) = \tau \circ \phi \circ \sigma^{(h)} \circ \phi^{(h)} \circ \cdots \circ \sigma^{(1)} \circ \phi^{(1)}(\mathbf{x})$ |

### Neural networks — neuron computation
Each neuron performs a **weighted sum** of its inputs followed by a **non-linear transformation**:

$$\phi^{(j)}(\mathbf{z}) = \mathbf{w}_j^\top \mathbf{z} + b_j, \qquad \sigma^{(j)}\!\left(\phi^{(j)}(\mathbf{z})\right) = \sigma^{(j)}\!\left(\mathbf{w}_j^\top \mathbf{z} + b_j\right)$$

The network as a whole is a **nested composition** of such operations.

---

## 3. Overfitting

- The capacity ("complexity") of a model can be increased by **increasing the size of the hypothesis space**.
- This usually also increases the number of **learnable parameters**.
- Examples: increasing the degree of a polynomial, increasing tree depth, growing a neural network, adding predictors.
- As $\mathcal{H}$ grows, the tendency to overfit increases — the model may fit even the **random quirks** in the training data, failing to generalize.

### Worked examples

**Polynomial regression** (data from $y = 3x_1 + 2x_1^2 + x_1^5 + \epsilon$, $\epsilon \sim \mathcal{N}(0, 1.25)$):

| | Degree 1 | Degree 5 | Degree 13 |
|--|----------|----------|-----------|
| Capacity | Low (underfit) | Appropriate | High (overfit) |
| Training error (RMSE) | 3.87 | 1.23 | **0.48** |
| Test error (RMSE) | 4.11 | **1.55** | 148.5 |

**Decision trees** (`minsplit` = min samples in a node to split; smaller = more capacity):

| | minsplit 60 | minsplit 12 | minsplit 2 |
|--|-------------|-------------|------------|
| Capacity | Low (underfit) | Appropriate | High (overfit) |
| Training error (Misclass.) | 0.36 | 0.12 | **0.02** |
| Test error (Misclass.) | 0.40 | **0.32** | 0.35 |

**k-Nearest Neighbours** (smaller $k$ = more capacity):

| | k = 20 | k = 7 | k = 1 |
|--|--------|-------|-------|
| Capacity | Low (underfit) | Appropriate | High (overfit) |
| Training error (Misclass.) | 0.22 | 0.13 | **0** |
| Test error (Misclass.) | 0.40 | **0.25** | 0.33 |

> **Pattern:** Training error keeps **dropping** as capacity rises (lowest at highest capacity), but test error is **U-shaped** — minimized at intermediate capacity.

---

## 4. The Complexity of Hypothesis Spaces — VC Dimension

A general measure of the complexity of a function space is the **Vapnik–Chervonenkis (VC) dimension**.

- The **VC dimension** of a class of binary-valued functions $\mathcal{H} = \{h : \mathcal{X} \to \{0,1\}\}$ is the **largest number of points** (in some configuration) that can be **shattered** by members of $\mathcal{H}$. Written $VC_p(\mathcal{H})$, where $p$ = dimension of input space.
- A set of points is **shattered** by a class if a member can perfectly separate them **no matter how** we assign binary labels.

> **Note:** VC dimension $d$ does **not** mean *all* sets of size $d$ are shattered. It means there is **at least one** set of size $d$ that can be shattered, and **no** set of size $d+1$ that can.

### Key examples / theorems

| Hypothesis class (in $\mathbb{R}^p$) | VC dimension |
|--------------------------------------|--------------|
| Linear indicator functions in $\mathbb{R}^2$ ($h = \mathbb{1}(\mathbf{x}^\top\boldsymbol\theta - \theta_0 > 0)$) | **3** (can shatter 3 points, not 4) |
| Homogeneous halfspaces $h = \text{sign}(\mathbf{x}^\top\boldsymbol\theta)$ | exactly $p$ |
| Non-homogeneous halfspaces $h = \text{sign}(\mathbf{x}^\top\boldsymbol\theta + \theta_0)$ | exactly $p+1$ |
| Axis-aligned rectangles in $\mathbb{R}^2$ | **4** |
| Single-parametric threshold classifier $h = \mathbb{1}(x \ge \theta)$ | **1** |
| 1-nearest neighbour ($k=1$) | **infinite** |
| Single-parametric sine classifier $h = \mathbb{1}(\sin(\theta x) > 0)$ | **infinite** |

> **Caution:** VC dimension generally **increases** with the number of learnable parameters, but capacity **cannot** always be judged from parameter count alone — the sine classifier has **one** parameter yet **infinite** VC dimension.

### VC bound and PAC learning

The training error is an **optimistic** estimate of the generalization (test) error. For a classification model with VC dimension $d$, with probability $1-\delta$:

$$\mathbb{P}\!\left(\text{err}_{\text{test}} \le \text{err}_{\text{train}} + \underbrace{\sqrt{\tfrac{1}{|\mathcal{D}_{\text{train}}|}\left[d\left(\log\tfrac{2|\mathcal{D}_{\text{train}}|}{d} + 1\right) - \log\tfrac{\delta}{4}\right]}}_{\epsilon}\right) = 1-\delta$$

for $\delta \in [0,1]$, provided the training set is large enough ($d < |\mathcal{D}_{\text{train}}|$).

- The bound holds over **all** datasets of size $|\mathcal{D}_{\text{train}}|$ drawn from an **arbitrary** $\mathbb{P}_{xy}$.
- If $d$ is **finite**, both $\delta$ and $\epsilon$ can be made arbitrarily small by **increasing the sample size**.
- **Corollary:** if the training error is low, the test error is also (probably) low. An algorithm that minimizes training error reliably picks a "good" hypothesis → **Probably Approximately Correct (PAC)** algorithm.

> **Reality check:** VC bounds are **extremely loose and pessimistic** (they must hold for arbitrary $\mathbb{P}_{xy}$). Complex models like neural networks often perform far better than the bounds suggest. Tighter alternatives exist (e.g. **Rademacher complexity**), and the *effective* capacity also depends on the **optimizer**. A better estimate of generalization error is simply evaluating on the test set.

---

## 5. Bias–Variance Decomposition

Generalization error of an inducer $\mathcal{I}_{L,\mathcal{O}}$:

$$GE_n(\mathcal{I}_{L,\mathcal{O}}) = \mathbb{E}_{\mathcal{D}_n, xy}\left(L\left(y, \hat f_{\mathcal{D}_n}(\mathbf{x})\right)\right)$$

**Assumption:** data is generated by $y = f_{\text{true}}(\mathbf{x}) + \epsilon$, with $\epsilon \sim \mathcal{N}(0, \sigma^2)$ independent of $\mathbf{x}$ (so $y \sim \mathcal{N}(f_{\text{true}}(\mathbf{x}), \sigma^2)$).

Plugging in the **L2 loss** $L(y, f(\mathbf{x})) = (y - f(\mathbf{x}))^2$ and taking the expectation over $(\mathbf{x}, y) \sim \mathbb{P}_{xy}$ decomposes the error into **three** terms:

$$GE_n(\mathcal{I}_{L,\mathcal{O}}) = \underbrace{\sigma^2}_{\text{variance of data}} + \underbrace{\mathbb{E}_{xy}\!\left[\text{Var}_{\mathcal{D}_n}\!\left(\hat f_{\mathcal{D}_n}(\mathbf{x}) \mid \mathbf{x}, y\right)\right]}_{\text{variance of inducer}} + \underbrace{\mathbb{E}_{xy}\!\left[\mathbb{E}^2_{\mathcal{D}_n}\!\left(f_{\text{true}}(\mathbf{x}) - \hat f_{\mathcal{D}_n}(\mathbf{x}) \mid \mathbf{x}, y\right)\right]}_{\text{squared bias of inducer}}$$

### The three terms

1. **Variance of the data ($\sigma^2$):** the **noise** in the data. Also called **intrinsic / unavoidable / irreducible error**. No learner can ever get below this.
2. **Variance of the inducer:** how much the prediction varies w.r.t. the training data used. Captures the tendency to learn **random things** irrespective of the real signal → **overfitting**.
3. **Squared bias:** the learner's tendency to **consistently misclassify** certain instances → **underfitting**.

### Capacity ↔ bias/variance
- **High capacity** → **low bias**, **high variance**.
- **Low capacity** → **high bias**, **low variance**.

> Figure intuition — *high bias*: a model too rigid to fit a curved relationship. *High variance*: a flexible model that can in principle learn the truth, but outputs **wildly different** hypotheses for different training sets.

---

## 6. Learning Curves

### Error vs. capacity
As capacity increases:
- Error due to **variance increases**.
- Error due to **bias decreases**.
- In the **overfitting region**, generalization error rises because the increase in variance is much larger than the decrease in bias.

### Error vs. training-set size
As the training set grows:
- Error due to **variance vanishes**.
- Generalization error and training error **converge to the bias** of the algorithm (assuming noise is zero).
- The **generalization gap vanishes**.

### Diagnosis & remedies

| Symptom | Cause | Remedy |
|---------|-------|--------|
| High bias (underfitting) | Model too rigid | Make model **more flexible** (or choose another) — *reduce bias* |
| High variance (overfitting) | Model too flexible | Make model **less flexible (regularization)** or **add more data** — *reduce variance* |

---

## 7. ML as an Ill-Posed Problem

- A learning algorithm must perform well on **previously unseen** (test) data.
- Suppose we learn a boolean function over 4 boolean features. There are $2^{16} = 65536$ possible functions ($2^4 = 16$ possible feature combinations → $2^{16}$ labelings).
- Given a training set of 7 examples, **$2^9$ possible functions remain** consistent with the data.
- The unseen datapoints can have **any** labels → **machine learning is an ill-posed problem**.
- Without further assumptions, can an ML algorithm really do better than **random guessing**? Not in general.

---

## 8. No Free Lunch (NFL) Theorem

- For a **specific** problem (a specific distribution over target functions), some algorithms can beat others.
- **Averaged over all possible problems**, **no algorithm** is better than any other — including random guessing. This is the **No Free Lunch theorem**.
- An algorithm that does well on one set of problems must necessarily do **badly** on another. There is no "free lunch."

### Takeaways
- No algorithm is **universally** better than all others.
- A learning algorithm must be tailored to a **specific prior** — without assumptions about the problem, learning is impossible.
- **Very specific** assumptions → great on a narrow set of problems.
- **Very broad** assumptions → a "jack of all trades, master of none."

---

## Quick-Check: Common True/False Traps

- **Underfitting** = high training error; **overfitting** = large train-test gap.
- Test error is **U-shaped** in capacity; training error **monotonically decreases**.
- Higher capacity → **lower bias, higher variance** (not the other way around).
- VC dim $d$ = there exists **one** shatterable set of size $d$, and **none** of size $d+1$ (NOT all sets of size $d$).
- Homogeneous halfspaces in $\mathbb{R}^p$: VC = $p$; **non-homogeneous**: VC = $p+1$.
- Axis-aligned rectangles in $\mathbb{R}^2$: VC = **4**. Threshold classifier $\mathbb{1}(x\ge\theta)$: VC = **1**.
- A **one-parameter** sine classifier has **infinite** VC dimension — parameter count ≠ capacity.
- The **irreducible error** $\sigma^2$ can **never** be removed by any learner.
- VC bounds are **loose & pessimistic**; real models often do much better.
- **No Free Lunch:** averaged over *all* problems, no learner beats random guessing.
- Learning is **ill-posed** without prior assumptions.
