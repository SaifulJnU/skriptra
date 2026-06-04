# Chapter 4: Risk Minimization I — Summary

## 1. Learning = Representation + Cost + Optimization

A learning algorithm consists of three components:

$$\text{Learning} = \underbrace{\text{Representation}}_{\text{hypothesis space } \mathcal{H}} + \underbrace{\textbf{Cost}}_{\text{how good is a model?}} + \underbrace{\text{Optimization}}_{\text{how do we get there?}}$$

- **Representation:** the hypothesis space $\mathcal{H}$ (e.g. linear models) — covered in Chapter 3.
- **Cost:** how we distinguish good models from bad ones (the **cost / loss function**).
- **Optimization:** how we actually find the best model $\hat f \in \mathcal{H}$.

> **Scenario:** Given the hypothesis space of linear models, *which* model is returned? **It depends on how the cost function is specified.**

---

## 2. Losses: Measuring Errors Point-wise

We search for $f \in \mathcal{H}$ that predicts outputs $f(\mathbf{x})$ close to the real targets $y$: $\; y \approx f(\mathbf{x})$.

The **loss function** measures the "goodness" of a prediction **point-wise**:

$$L : \mathcal{Y} \times \mathbb{R}^g \to \mathbb{R}_{\ge 0}$$

**Requirements:**
- **Non-negativity:** $L(y, \tilde y) \ge 0$ for all $y, \tilde y \in \mathcal{Y}$
- **Optimality:** $L(y, \tilde y) = 0 \iff y = \tilde y$

**Example (point-wise squared error):** $L(y, f(\mathbf{x})) = (y - f(\mathbf{x}))^2$.

---

## 3. Residuals, Pseudo-Residuals & Loss Properties

Regression losses usually depend only on the **residual**:

$$r := y - f(\mathbf{x}), \qquad r^{(i)} := y^{(i)} - f(\mathbf{x}^{(i)})$$

- A loss is **distance-based** if it can be written as $L(y, f(\mathbf{x})) = \psi(r)$ for some $\psi : \mathbb{R} \to \mathbb{R}$, and it is zero iff $r = 0$ (i.e. $\psi(0) = 0$).
- A loss is **translation-invariant** if $L(y + a, f(\mathbf{x}) + a) = L(y, f(\mathbf{x}))$ for $a \in \mathbb{R}$.

> **Key fact:** A loss is **translation-invariant iff it is distance-based**.

### Pseudo-residuals
The **pseudo-residual** is the negative first derivative of the loss w.r.t. $f$:

$$\tilde r := -\frac{\partial L(y, f(\mathbf{x}))}{\partial f}, \qquad \tilde r^{(i)} := -\frac{\partial L(y^{(i)}, f(\mathbf{x}^{(i)}))}{\partial f}$$

> For the **L2-loss**, pseudo-residuals **coincide** with the (ordinary) residuals — hence the name. ($-\partial(y-f)^2/\partial f = 2(y-f)$, or $r$ for the $0.5$-scaled version.)

### Loss plot
The **loss plot** shows the point-wise error $L(y, f(\mathbf{x}))$ vs. the residual $y - f(\mathbf{x})$. The pseudo-residual corresponds to the **slope of the tangent** at $(y - f(\mathbf{x}), L(y, f(\mathbf{x})))$.

---

## 4. (Theoretical) Risk Minimization

Let $\mathbb{P}_{xy}$ be the joint distribution that generates the data.

**Goal:** find $f \in \mathcal{H}$ minimizing the **expected loss (risk)** over $(\mathbf{x}, y) \sim \mathbb{P}_{xy}$:

$$\min_{f \in \mathcal{H}} \mathcal{R}(f) = \min_{f \in \mathcal{H}} \mathbb{E}[L(y, f(\mathbf{x}))] = \min_{f \in \mathcal{H}} \int_{\mathcal{X} \times \mathcal{Y}} L(y, f(\mathbf{x})) \, d\mathbb{P}_{xy}$$

### L2-loss → conditional expectation
With $\mathcal{Y} = \mathbb{R}$ and unrestricted $\mathcal{H}$, minimizing the risk point-wise for each $\mathbf{x}$:

$$\hat f(\mathbf{x}) = \arg\min_c \mathbb{E}_{y|x}\!\left[(y - c)^2 \mid \mathbf{x} = \mathbf{x}\right] = \mathbb{E}_{y|x}[y \mid \mathbf{x} = \mathbf{x}]$$

This follows from $\mathbb{E}[(y-c)^2] = \text{Var}[y] + (\mathbb{E}[y] - c)^2$, minimal at $c = \mathbb{E}[y]$.

> **For squared loss, the best prediction is the conditional expectation $\mathbb{E}[y \mid \mathbf{x}]$.**

### Limitation
Minimizing $\mathcal{R}(f)$ directly is generally not feasible:
- $\mathbb{P}_{xy}$ is **unknown**.
- Estimating it non-parametrically (e.g. kernel density estimation) does **not scale** to high dimensions (curse of dimensionality).
- We *can* estimate $\mathbb{P}_{xy}$ efficiently with rigorous distributional assumptions (e.g. discriminant analysis: LDA/QDA), but ML usually studies **more flexible** models.

---

## 5. Empirical Risk Minimization (ERM)

Assume $\mathcal{D}$ is drawn i.i.d. from $\mathbb{P}_{xy}$. Approximate the risk using the data via the **empirical risk**:

$$\mathcal{R}_{\text{emp}}(f) = \sum_{i=1}^n L\!\left(y^{(i)}, f(\mathbf{x}^{(i)})\right), \qquad \bar{\mathcal{R}}_{\text{emp}}(f) = \frac{1}{n}\sum_{i=1}^n L\!\left(y^{(i)}, f(\mathbf{x}^{(i)})\right)$$

(The factor $\tfrac{1}{n}$ does not affect optimization.)

> **Note:** $\mathcal{R}_{\text{emp}}$ is a good approximation of $\mathcal{R}$ only if $\mathcal{D}$ is an **unbiased, independent, and large enough** sample from $\mathbb{P}_{xy}$.

Learning then amounts to **empirical risk minimization**:

$$\hat f = \arg\min_{f \in \mathcal{H}} \mathcal{R}_{\text{emp}}(f), \qquad \hat{\boldsymbol\theta} = \arg\min_{\boldsymbol\theta \in \Theta} \mathcal{R}_{\text{emp}}(\boldsymbol\theta)$$

> Learning (often) **means solving an optimization problem** → tight connection between ML and optimization.

### Why the choice of loss matters
- **Statistical properties:** the loss determines robustness and the implicit error distribution.
- **Computational/optimization complexity:**
  - **Smoothness:** some optimizers (e.g. gradient methods) require smoothness.
  - **Uni-/multimodality:** if $L$ is **convex** in its 2nd argument and $f(\mathbf{x}\mid\boldsymbol\theta)$ is **linear** in $\boldsymbol\theta$, then $\mathcal{R}_{\text{emp}}(\boldsymbol\theta)$ is convex and **every local minimum is global**. If $L$ is **not convex**, there may be multiple local minima (bad!).

---

## 6. Regression Losses

| Loss | Formula | Properties | Optimal constant model |
|------|---------|------------|------------------------|
| **L2 (squared)** | $(y-f)^2$ or $0.5(y-f)^2$ | Convex, differentiable; **sensitive to outliers** (residual ×2 → loss ×4) | **mean** $\bar y$ |
| **L1 (absolute)** | $\lvert y-f \rvert$ | Convex, **robust**, **not differentiable** at $y=f$ | **median** |
| **Quantile (pinball)** | see below | Extension of L1; asymmetric | empirical **$\alpha$-quantile** $Q_\alpha$ |
| **Huber** | see below | Convex, **differentiable + robust**; no closed form | numerical only |
| **$\epsilon$-insensitive** | see below | Convex, not differentiable; ignores small errors | numerical only |
| **Log-barrier** | see below | L2-like for small $r$; forbids $\lvert r \rvert > a$ | may have no solution |

### L2-loss
$$L(y, f) = (y - f)^2$$
Tries to reduce large residuals → outliers in $y$ become problematic. Residuals and pseudo-residuals **coincide**: $-\partial(0.5(y-f)^2)/\partial f = y - f = r$. Optimal constant: $\hat\theta = \tfrac1n\sum y^{(i)} = \bar y$.

### L1-loss
$$L(y, f) = \lvert y - f \rvert$$
More robust than L2 (outliers less problematic), but not differentiable at $y = f$ (optimization harder). Optimal constant: $\hat\theta = \text{median}(y^{(i)})$.

### Quantile / Pinball loss
$$L(y, f) = \begin{cases}(1-\alpha)(f - y), & y < f \\ \alpha(y - f), & y \ge f\end{cases}, \quad \alpha \in (0,1)$$
- Extension of L1 ($\alpha = 0.5$ gives L1, up to scaling).
- Weights positive/negative residuals differently: $\alpha < 0.5$ penalizes over-estimation, $\alpha > 0.5$ penalizes under-estimation.
- Optimal constant = empirical **$\alpha$-quantile** $Q_\alpha(\{y^{(i)}\})$.

### Huber loss
$$L(y, f) = \begin{cases}\tfrac12(y - f)^2, & \lvert y - f \rvert \le \delta \\ \delta \lvert y - f \rvert - \tfrac12\delta^2, & \text{otherwise}\end{cases}$$
Piecewise combination of L1 and L2 → **differentiable + robust**. No closed-form optimal constant (needs numerical optimization); it lies **between** the L1 and L2 solutions.

### $\epsilon$-insensitive loss
$$L(y, f) = \begin{cases}0, & \lvert y - f \rvert \le \epsilon \\ \lvert y - f \rvert - \epsilon, & \text{otherwise}\end{cases}, \quad \epsilon \in \mathbb{R}_+$$
Modification of L1: errors below $\epsilon$ accepted without penalty. Convex, not differentiable, no closed form. (Used in support vector regression.)

### Log-barrier loss
$$L(y, f) = \begin{cases}-a^2 \log\!\left(1 - \left(\tfrac{\lvert y-f\rvert}{a}\right)^2\right), & \lvert y - f \rvert \le a \\ \infty, & \lvert y - f \rvert > a\end{cases}$$
Behaves like L2 for small residuals; used when we want **no residuals larger than $a$ at all**. No guarantee the risk-minimization problem has a solution.

---

## 7. Numerical Optimization

When there is **no closed-form** (analytical) solution to $\min_{\boldsymbol\theta} \bar{\mathcal{R}}_{\text{emp}}(\boldsymbol\theta)$, we use **numerical optimization**.

### Gradient Descent
Take steps in the direction of the **negative gradient** (direction of steepest descent):

$$\boldsymbol\theta^{[j+1]} = \boldsymbol\theta^{[j]} - \alpha^{[j]} \cdot \nabla_{\boldsymbol\theta}\, \mathcal{R}_{\text{emp}}(\boldsymbol\theta)\big|_{\boldsymbol\theta = \boldsymbol\theta^{[j]}}$$

- $\alpha^{[j]}$ = **step-size / learning rate** (fixed, line-search, ...).
- First-order iterative algorithm.
- Stopping rule (example): $\dfrac{\lVert \boldsymbol\theta^{[j+1]} - \boldsymbol\theta^{[j]} \rVert}{\lVert \boldsymbol\theta^{[j]} \rVert} < \varepsilon$ (e.g. $\varepsilon = 0.0001$).

Using the chain rule, the update can be written with **pseudo-residuals**:

$$\boldsymbol\theta^{[j+1]} \leftarrow \boldsymbol\theta^{[j]} + \alpha^{[j]} \frac{1}{n}\sum_{i=1}^n \tilde r^{(i)} \cdot \nabla_{\boldsymbol\theta} f(\mathbf{x}^{(i)} \mid \boldsymbol\theta)\big|_{\boldsymbol\theta = \boldsymbol\theta^{[j]}}$$

### Stochastic Gradient Descent (SGD)
- A stochastic approximation of gradient descent, used when $\sum_i \nabla_{\boldsymbol\theta} L$ is **expensive** (every summand must be evaluated).
- Approximates the gradient using just **one random observation $i$**:
$$\boldsymbol\theta^{[j+1]} \leftarrow \boldsymbol\theta^{[j]} - \alpha^{[j]} \nabla_{\boldsymbol\theta} L\!\left(y^{(i)}, f(\mathbf{x}^{(i)} \mid \boldsymbol\theta)\right)$$
- The parameter sequence is **stochastic** (depends on the randomly drawn observation each step).

### Mini-batch Gradient Descent
- A **trade-off** between full gradient descent (all observations) and SGD (one observation).
- Uses a set of randomly drawn observations $I \subset \{1, \ldots, n\}$:
$$\boldsymbol\theta^{[j+1]} \leftarrow \boldsymbol\theta^{[j]} - \alpha^{[j]} \sum_{i \in I} \nabla_{\boldsymbol\theta} L\!\left(y^{(i)}, f(\mathbf{x}^{(i)} \mid \boldsymbol\theta)\right)$$
- SGD is computationally cheap but **noisier**; mini-batch balances cost and noise.

---

## 8. Maximum Likelihood Estimation (MLE) vs. ERM

Approach regression from the **maximum likelihood** perspective. Assume:

$$y = f_{\text{true}}(\mathbf{x}) + \epsilon, \quad \epsilon \sim p, \;\; \mathbb{E}[\epsilon] = 0, \;\; \epsilon \perp \mathbf{x}$$

Then $y$ follows a distribution with mean $f_{\text{true}}(\mathbf{x})$, density $p(y \mid \mathbf{x}, \boldsymbol\theta)$.

**Maximum-likelihood principle** — maximize the likelihood, or minimize the negative log-likelihood:

$$\mathcal{L}(\boldsymbol\theta) = \prod_{i=1}^n p\!\left(y^{(i)} \mid \mathbf{x}^{(i)}, \boldsymbol\theta\right), \qquad -\ell(\boldsymbol\theta) = -\sum_{i=1}^n \log p\!\left(y^{(i)} \mid \mathbf{x}^{(i)}, \boldsymbol\theta\right)$$

Define a **new loss**: $\; L(y, f(\mathbf{x} \mid \boldsymbol\theta)) := -\log p(y \mid \mathbf{x}, \boldsymbol\theta)$.

> Then the ML estimator $\hat{\boldsymbol\theta}$ (maximizing $\mathcal{L}$) is **identical** to the loss-minimal $\hat{\boldsymbol\theta}$ (minimizing $\mathcal{R}_{\text{emp}}$).

### Key correspondence
- For **every error distribution** we can derive an **equivalent loss function** giving the same point estimator for $\boldsymbol\theta$.
- Multiplicative/additive constants in the loss can be dropped (they don't change the minimizer).
- **⚠️ The reverse does NOT always hold:** not every loss corresponds to an error distribution — the **hinge loss** is a prominent counterexample.

### Famous correspondences

| Error distribution | Density | Equivalent loss |
|--------------------|---------|-----------------|
| **Gaussian** $\epsilon \sim \mathcal{N}(0, \sigma^2)$ | $\propto \exp(-\tfrac{(y-f)^2}{2\sigma^2})$ | **L2-loss** $\sum (y^{(i)} - f)^2$ |
| **Laplace** $\frac{1}{2b}\exp(-\tfrac{\lvert x-\mu\rvert}{b})$ | $\propto \exp(-\tfrac{\lvert y-f\rvert}{b})$ | **L1-loss** $\sum \lvert y^{(i)} - f \rvert$ |

> **Gaussian errors ⟺ L2-loss**, **Laplace errors ⟺ L1-loss.**

### Empirical error distributions
- We can plot the "empirical" error distribution = histogram of residuals after fitting w.r.t. a given loss.
- Some losses (Huber, $\epsilon$-insensitive) do **not** correspond to "real" error densities, but intuitively a loss still **defines how residuals will be distributed**.

---

## Quick-Check: Common True/False Traps

- **Learning = Representation + Cost + Optimization** (three components).
- Loss requirements: **non-negativity** and **$L=0 \iff y=\tilde y$**.
- A loss is **translation-invariant iff distance-based**.
- **Pseudo-residual = negative derivative of loss w.r.t. $f$**; for L2 it **equals** the residual.
- For **squared loss**, the optimal (theoretical) prediction is the **conditional expectation** $\mathbb{E}[y\mid\mathbf{x}]$.
- Optimal **constant** model: L2 → **mean**, L1 → **median**, quantile → **$\alpha$-quantile**.
- L2 is **sensitive to outliers**; L1/Huber/quantile are more **robust**.
- L1 is **not differentiable** at $0$; Huber is **differentiable AND robust**.
- Convex loss + linear-in-$\boldsymbol\theta$ model → **every local min is global**.
- The **learning rate** is the step-size $\alpha$ in gradient descent.
- **SGD** uses one random point; **mini-batch** uses a subset; full GD uses all points.
- **Gaussian errors ⟺ L2**, **Laplace errors ⟺ L1** (via negative log-likelihood).
- Every error distribution → a loss, but **NOT** every loss → an error distribution (**hinge loss**).
