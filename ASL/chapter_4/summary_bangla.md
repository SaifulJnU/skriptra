# অধ্যায় ৪: Risk Minimization I
### — বাংলায় সহজ ভাষায় সম্পূর্ণ ব্যাখ্যা

---

> **এই chapter-এর সবচেয়ে কঠিন অংশ:** Loss function-গুলোর পার্থক্য (L2 vs L1 vs Huber vs Quantile), optimal constant model (mean/median/quantile), pseudo-residual, এবং MLE ↔ Loss-এর সম্পর্ক (Gaussian↔L2, Laplace↔L1) — এই জিনিসগুলো মাথায় গেঁথে নাও।

---

## ১. Learning = Representation + Cost + Optimization

### গল্প: জামা বানানো

একটা ভালো model বানানো = জামা বানানোর মতো, তিনটা জিনিস লাগে:

$$\text{Learning} = \underbrace{\text{Representation}}_{\text{কী ধরনের model}} + \underbrace{\textbf{Cost}}_{\text{ভালো-খারাপ মাপা}} + \underbrace{\text{Optimization}}_{\text{সেরাটা খুঁজে বের করা}}$$

- **Representation:** hypothesis space $\mathcal{H}$ (যেমন linear model) — Chapter 3।
- **Cost:** কোন model ভালো কোনটা খারাপ — সেটা মাপার নিয়ম (**loss function**)।
- **Optimization:** সেরা model $\hat f$ কীভাবে খুঁজে পাব।

> **মূল কথা:** একই linear model space দিলেও কোন model বের হবে — সেটা **cost function-এর উপর নির্ভর করে।**

---

## ২. Loss — ভুল মাপার যন্ত্র

আমরা চাই $f(\mathbf{x}) \approx y$। কতটা ভুল হলো সেটা **point-wise** (প্রতিটা data point-এ আলাদা করে) মাপি **loss function** দিয়ে:

$$L : \mathcal{Y} \times \mathbb{R}^g \to \mathbb{R}_{\ge 0}$$

**দুটো শর্ত:**
1. **Non-negativity:** $L(y, \tilde y) \ge 0$ — ভুল কখনো ঋণাত্মক হয় না।
2. **Optimality:** $L(y, \tilde y) = 0 \iff y = \tilde y$ — ভুল শূন্য মানেই prediction একদম সঠিক।

**উদাহরণ:** squared error $L(y, f(\mathbf{x})) = (y - f(\mathbf{x}))^2$।

---

## ৩. Residual, Pseudo-Residual ও Loss-এর ধর্ম

### Residual
$$r := y - f(\mathbf{x}) \quad (\text{আসল মান} - \text{prediction})$$

### দুটো গুরুত্বপূর্ণ ধর্ম

- **Distance-based:** loss-টা শুধু residual দিয়ে লেখা যায়, $L = \psi(r)$, এবং $\psi(0) = 0$।
- **Translation-invariant:** $L(y + a, f(\mathbf{x}) + a) = L(y, f(\mathbf{x}))$।

> **🔑 মূল সত্য:** একটা loss **translation-invariant হয় তখনই, যখন সেটা distance-based**। (দুটো equivalent)

### Pseudo-Residual
Loss-এর $f$-এর সাপেক্ষে **ঋণাত্মক first derivative**:

$$\tilde r := -\frac{\partial L(y, f(\mathbf{x}))}{\partial f}$$

> **L2-loss-এর জন্য pseudo-residual আর সাধারণ residual এক হয়ে যায়** — এজন্যই নাম "pseudo-residual"।

### Loss Plot
$L(y, f(\mathbf{x}))$ বনাম residual $(y - f(\mathbf{x}))$-এর গ্রাফ। Pseudo-residual = ওই বিন্দুতে **tangent-এর slope**।

---

## ৪. Theoretical Risk Minimization

$\mathbb{P}_{xy}$ = data যে distribution থেকে আসে।

**Goal:** এমন $f$ খুঁজো যা **expected loss (risk)** সবচেয়ে কম করে:

$$\min_{f \in \mathcal{H}} \mathcal{R}(f) = \min_{f} \mathbb{E}[L(y, f(\mathbf{x}))]$$

### L2-loss → Conditional Expectation

Hypothesis space-এ কোনো বাধা না থাকলে, প্রতিটা $\mathbf{x}$-এর জন্য সেরা prediction:

$$\hat f(\mathbf{x}) = \mathbb{E}[y \mid \mathbf{x}]$$

কারণ: $\mathbb{E}[(y-c)^2] = \text{Var}[y] + (\mathbb{E}[y] - c)^2$, যা $c = \mathbb{E}[y]$-তে সবচেয়ে ছোট।

> **🔑 Squared loss-এর জন্য সেরা prediction = conditional expectation $\mathbb{E}[y\mid\mathbf{x}]$।**

### সমস্যা (Limitation)
- $\mathbb{P}_{xy}$ **অজানা।**
- Non-parametric estimate (যেমন kernel density) **high dimension-এ কাজ করে না** (curse of dimensionality)।
- কড়া assumption দিলে estimate করা যায় (LDA/QDA), কিন্তু ML সাধারণত **flexible** model নিয়ে কাজ করে।

---

## ৫. Empirical Risk Minimization (ERM)

$\mathbb{P}_{xy}$ অজানা বলে, data $\mathcal{D}$ দিয়ে risk-কে approximate করি **empirical risk** দিয়ে:

$$\mathcal{R}_{\text{emp}}(f) = \sum_{i=1}^n L\!\left(y^{(i)}, f(\mathbf{x}^{(i)})\right)$$

($\tfrac1n$ গুণ করলেও minimizer বদলায় না।)

> **Note:** $\mathcal{R}_{\text{emp}}$ ভালো approximation **শুধু তখনই** যখন data **unbiased, independent, এবং যথেষ্ট বড়**।

Learning মানে:
$$\hat f = \arg\min_{f \in \mathcal{H}} \mathcal{R}_{\text{emp}}(f)$$

> **🔑 Learning (প্রায়ই) মানে একটা optimization problem solve করা।**

### Loss বেছে নেওয়া কেন গুরুত্বপূর্ণ?
- **Statistical:** robustness, implicit error distribution নির্ধারণ করে।
- **Computational:**
  - **Smoothness:** কিছু optimizer (gradient method) smoothness দরকার।
  - **Convexity:** loss **convex** + $f$ **linear-in-$\theta$** হলে → প্রতিটা local minimum-ই global! Non-convex হলে → অনেক local minima (খারাপ!)।

---

## ৬. Regression Losses — সবগুলো একসাথে

| Loss | Formula | ধর্ম | Optimal constant model |
|------|---------|------|------------------------|
| **L2** | $(y-f)^2$ | Convex, differentiable, **outlier-sensitive** | **mean (গড়)** |
| **L1** | $\lvert y-f \rvert$ | **Robust**, $y=f$-এ differentiable **না** | **median** |
| **Quantile** | pinball | L1-এর extension, asymmetric | **$\alpha$-quantile** |
| **Huber** | L1+L2 piecewise | **differentiable + robust**, closed-form নেই | numerical |
| **$\epsilon$-insensitive** | modified L1 | ছোট error উপেক্ষা করে | numerical |
| **Log-barrier** | — | residual > $a$ নিষিদ্ধ | solution নাও থাকতে পারে |

### L2-loss (Squared)
$$L = (y - f)^2$$
- বড় residual-কে বেশি শাস্তি দেয় (residual ২ গুণ → loss ৪ গুণ) → **outlier সমস্যা করে।**
- Convex, differentiable। Residual = pseudo-residual।
- **Optimal constant = গড় $\bar y$।**

### L1-loss (Absolute)
$$L = \lvert y - f \rvert$$
- L2-এর চেয়ে **robust** (outlier কম সমস্যা করে)।
- $y = f$-এ **differentiable না** → optimization কঠিন।
- **Optimal constant = median।**

### Quantile / Pinball loss
$$L = \begin{cases}(1-\alpha)(f - y), & y < f \\ \alpha(y - f), & y \ge f\end{cases}$$
- L1-এর extension; $\alpha = 0.5$ দিলে L1।
- $\alpha < 0.5$ → over-estimation-কে শাস্তি; $\alpha > 0.5$ → under-estimation-কে শাস্তি।
- **Optimal constant = empirical $\alpha$-quantile।**

### Huber loss
$$L = \begin{cases}\tfrac12(y - f)^2, & \lvert y - f \rvert \le \delta \\ \delta \lvert y - f \rvert - \tfrac12\delta^2, & \text{নাহলে}\end{cases}$$
- L1 আর L2-এর মিশ্রণ → **differentiable + robust** (দুটোর সুবিধা একসাথে)।
- Closed-form নেই; solution L1 আর L2-এর **মাঝখানে**।

### $\epsilon$-insensitive loss
$$L = \begin{cases}0, & \lvert y - f \rvert \le \epsilon \\ \lvert y - f \rvert - \epsilon, & \text{নাহলে}\end{cases}$$
- L1-এর modification: $\epsilon$-এর কম error-এ কোনো শাস্তি নেই।

### Log-barrier loss
ছোট residual-এ L2-এর মতো; residual **$a$-এর বেশি একদম চায় না** ($\infty$ শাস্তি)। Solution থাকার নিশ্চয়তা নেই।

---

## ৭. Numerical Optimization

Closed-form (analytical) solution না থাকলে → **numerical optimization** ব্যবহার করি।

### Gradient Descent
**Negative gradient**-এর দিকে ধাপে ধাপে নামি (steepest descent):

$$\boldsymbol\theta^{[j+1]} = \boldsymbol\theta^{[j]} - \alpha^{[j]} \cdot \nabla_{\boldsymbol\theta}\, \mathcal{R}_{\text{emp}}(\boldsymbol\theta)$$

- $\alpha^{[j]}$ = **step-size / learning rate**।
- First-order iterative algorithm।
- **Stopping rule:** $\frac{\lVert \boldsymbol\theta^{[j+1]} - \boldsymbol\theta^{[j]} \rVert}{\lVert \boldsymbol\theta^{[j]} \rVert} < \varepsilon$ (যেমন $\varepsilon = 0.0001$)।

### Stochastic Gradient Descent (SGD)
- পুরো gradient হিসাব **ব্যয়বহুল** হলে ব্যবহার করি।
- **একটা random observation $i$** দিয়ে gradient approximate করে।
- Parameter sequence **stochastic** (প্রতি ধাপে random point-এর উপর নির্ভর করে)।

### Mini-batch Gradient Descent
- Full GD (সব point) আর SGD (একটা point)-এর মাঝামাঝি **trade-off**।
- একটা random subset $I \subset \{1, \ldots, n\}$ ব্যবহার করে।
- SGD সস্তা কিন্তু **বেশি noisy**; mini-batch ভারসাম্য রাখে।

---

## ৮. Maximum Likelihood (MLE) vs ERM — সবচেয়ে সুন্দর সংযোগ

Assumption: $y = f_{\text{true}}(\mathbf{x}) + \epsilon$, $\mathbb{E}[\epsilon] = 0$, $\epsilon \perp \mathbf{x}$।

**Maximum-likelihood principle** — likelihood সর্বোচ্চ করা = negative log-likelihood সর্বনিম্ন করা:

$$-\ell(\boldsymbol\theta) = -\sum_{i=1}^n \log p\!\left(y^{(i)} \mid \mathbf{x}^{(i)}, \boldsymbol\theta\right)$$

এখন একটা **নতুন loss** define করি:
$$L(y, f(\mathbf{x} \mid \boldsymbol\theta)) := -\log p(y \mid \mathbf{x}, \boldsymbol\theta)$$

> **🔑 তাহলে MLE estimator = loss-minimal estimator!** দুটো একই উত্তর দেয়।

### মূল সম্পর্ক
- **প্রতিটা error distribution → একটা equivalent loss** (একই point estimator)।
- Loss-এর multiplicative/additive constant বাদ দেওয়া যায় (minimizer বদলায় না)।
- **⚠️ উল্টোটা সবসময় হয় না:** প্রতিটা loss-এর জন্য error distribution থাকে **না** — **hinge loss** তার বিখ্যাত উদাহরণ।

### বিখ্যাত জোড়া (মুখস্থ রাখো)

| Error distribution | → Loss |
|--------------------|--------|
| **Gaussian** $\mathcal{N}(0, \sigma^2)$ | **L2-loss** |
| **Laplace** | **L1-loss** |

> **Gaussian error ⟺ L2**, **Laplace error ⟺ L1।**

### Empirical error distribution
- Model fit করার পর residual-গুলোর histogram = "empirical" error distribution।
- কিছু loss (Huber, $\epsilon$-insensitive) আসল error density-র সাথে মেলে না, কিন্তু intuitively loss নির্ধারণ করে **residual কীভাবে বিতরণ হবে**।

---

## ৯. Quick Revision Table — Quiz-এর আগে দেখো

| Statement | T/F | কারণ |
|-----------|-----|------|
| Learning = Representation + Cost + Optimization | TRUE | — |
| Loss function সবসময় ≥ 0 | TRUE | non-negativity |
| L = 0 ⟺ y = ỹ | TRUE | optimality |
| Translation-invariant ⟺ distance-based | TRUE | — |
| Pseudo-residual = positive derivative of loss | FALSE | **negative** derivative |
| L2-তে pseudo-residual = residual | TRUE | — |
| Squared loss-এর সেরা prediction = E[y\|x] | TRUE | conditional expectation |
| P_xy জানা, তাই theoretical risk সহজ | FALSE | P_xy **অজানা** |
| 1/n empirical risk-এর minimizer বদলায় | FALSE | বদলায় না |
| Convex loss + linear f → local min = global | TRUE | — |
| L2 outlier-sensitive | TRUE | residual×2 → loss×4 |
| L2-এর optimal constant = median | FALSE | **mean (গড়)** |
| L1-এর optimal constant = median | TRUE | — |
| L1 সব জায়গায় differentiable | FALSE | y=f-এ না |
| Quantile loss α=0.5 = L1 | TRUE | — |
| Quantile-এর optimal constant = α-quantile | TRUE | — |
| Huber differentiable + robust | TRUE | — |
| Huber-এর closed-form solution আছে | FALSE | numerical |
| ε-insensitive: ছোট error free | TRUE | — |
| Gradient descent positive gradient দিকে যায় | FALSE | **negative** gradient |
| Learning rate = step-size α | TRUE | — |
| SGD একটা random point ব্যবহার করে | TRUE | — |
| Mini-batch সব n point ব্যবহার করে | FALSE | subset |
| MLE = loss-minimal estimator | TRUE | L = −log p |
| প্রতিটা loss-এর error distribution আছে | FALSE | hinge loss ব্যতিক্রম |
| Gaussian error ⟺ L2 | TRUE | — |
| Laplace error ⟺ L1 | TRUE | — |

---

## ১০. Bonus: মনে রাখার Trick

### Optimal Constant Model — তিনটা "M"
```
L2 (squared) → Mean   (গড়)
L1 (absolute) → Median
Quantile     → α-quantile (Q_α)
Huber        → numerical (closed-form নেই)
```

### MLE ↔ Loss জোড়া
```
Gaussian (bell curve) → L2  (square, smooth)
Laplace  (sharp peak) → L1  (absolute, robust)
```
মনে রাখো: **Gauss-2 (L2), Laplace-1 (L1)।**

### L1 vs L2 — Robustness
```
L2: outlier-কে square করে → বড় শাস্তি → outlier টেনে নেয় (sensitive)
L1: outlier-কে absolute → কম শাস্তি → robust
Huber: ছোট error-এ L2, বড় error-এ L1 → best of both
```

### Gradient Descent পরিবার
```
Full GD:     সব n point   → নিখুঁত কিন্তু ধীর
SGD:         ১টা point     → দ্রুত কিন্তু noisy
Mini-batch:  subset I      → মাঝামাঝি (trade-off)
```
