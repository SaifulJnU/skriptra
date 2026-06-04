# অধ্যায় ৫: Linear Regression — Least Squares ও Normal Equations
## — Opus-এর বিস্তারিত বাংলা ব্যাখ্যা (একদম শিশুর জন্য)

---

> **এই নোটটা কাদের জন্য?**
> তোমার জন্য, যে এখনো বুঝে উঠতে পারোনি কেন linear regression-এর একটা "সরাসরি formula" আছে অথচ অন্য loss-গুলোর নেই। ধরে নিচ্ছি তুমি Chapter 1–4 পড়ে এসেছ — মানে regression, hypothesis space, loss, empirical risk, gradient descent জানো। এই অধ্যায়ের নতুন/গভীর জিনিসগুলো (conditional expectation = f(x), augmented notation, least-squares objective, **normal equations**, OLS closed-form, expectation-এর সংজ্ঞা, L1 recap) — সব ভেঙে ভেঙে বলব।

> **কীভাবে পড়বে?**
> ১) প্রতিটা section পড়ার পর থেমে নিজেকে জিজ্ঞেস করো: "এটা ক্লাস ৮-এর ছাত্রকেও বোঝাতে পারব?"
> ২) **Normal equations** আর **$\hat{\boldsymbol\theta} = (\mathbf{X}^\top\mathbf{X})^{-1}\mathbf{X}^\top\mathbf{y}$** — এই দুটো একদম মুখস্থ করে নাও, derivation সহ।
> ৩) Quiz-trap গুলো লাল কালি দিয়ে highlight করো।

---

# 📚 সূচিপত্র (Table of Contents)

1. [Linear model — পুরো ছবিটা](#১-linear-model--পুরো-ছবিটা)
2. [Gaussian noise কী এবং কেন](#২-gaussian-noise-কী-এবং-কেন)
3. [কেন f(x) = E[y|x]](#৩-কেন-fx--eyx)
4. [Augmented notation — intercept লুকানোর কৌশল](#৪-augmented-notation--intercept-লুকানোর-কৌশল)
5. [Design matrix X](#৫-design-matrix-x)
6. [Least-squares objective](#৬-least-squares-objective)
7. [Normal equations — ধাপে ধাপে derivation](#৭-normal-equations--ধাপে-ধাপে-derivation)
8. [OLS closed-form solution](#৮-ols-closed-form-solution)
9. [কখন formula কাজ করে না (invertibility)](#৯-কখন-formula-কাজ-করে-না-invertibility)
10. [Expectation কী — sample থেকে theoretical](#১০-expectation-কী--sample-থেকে-theoretical)
11. [Absolute loss (L1) recap](#১১-absolute-loss-l1-recap)
12. [সব মিলিয়ে একটা বড় গল্প](#১২-সব-মিলিয়ে-একটা-বড়-গল্প)
13. [Master Quiz-Trap Table](#১৩-master-quiz-trap-table)
14. [Golden Memorization Rules](#১৪-golden-memorization-rules)

---

# ১. Linear model — পুরো ছবিটা

ধরো তুমি বাড়ির দাম (y) ভবিষ্যদ্বাণী করতে চাও বাড়ির আয়তন, ঘরের সংখ্যা ইত্যাদি (x) দিয়ে। সবচেয়ে সহজ ধারণা — দাম feature-গুলোর একটা **সরলরৈখিক (linear) যোগফল**, প্লাস কিছু এলোমেলো ভুল:

$$y = \beta_0 + \mathbf{x}^\top \boldsymbol\beta + \varepsilon$$

- $\beta_0$ → শুরুর মান (কোনো feature ০ হলেও একটা base দাম)। একে **intercept** বলে।
- $\boldsymbol\beta$ → প্রতিটা feature কত গুরুত্বপূর্ণ (slope/coefficient)।
- $\varepsilon$ → যা model ধরতে পারে না, সেই **এলোমেলো noise**।

সব parameter একসাথে রাখি: $\boldsymbol\theta = (\beta_0, \boldsymbol\beta^\top)^\top$।

---

# ২. Gaussian noise কী এবং কেন

side notes-এর একদম উপরে লেখা: $\varepsilon \sim \mathcal{N}(0, \sigma^2)$।

মানে noise-টা **ঘণ্টা আকৃতির (bell curve)** distribution মেনে চলে — গড় ০ (কখনো একটু বেশি, কখনো একটু কম, কিন্তু গড়ে শূন্য), আর ছড়ানোর মাত্রা $\sigma^2$।

> 🔑 **কেন এটা গুরুত্বপূর্ণ?** Chapter 4-এ শিখেছ: **Gaussian noise ⟺ L2 loss**। তাই Gaussian ধরা মানেই আমরা square-error (L2) দিয়ে fit করছি — আর সেটারই সুন্দর closed-form সমাধান আছে।

---

# ৩. কেন f(x) = E[y|x]

$y$-এর গড় নিই, $\mathbf{x}$ fixed রেখে। যেহেতু noise-এর গড় শূন্য ($\mathbb{E}[\varepsilon]=0$), noise-টা উবে যায়:

$$\mathbb{E}[y \mid \mathbf{x}] = \beta_0 + \mathbf{x}^\top \boldsymbol\beta + \underbrace{\mathbb{E}[\varepsilon]}_{=0} = \beta_0 + \mathbf{x}^\top \boldsymbol\beta = f(\mathbf{x})$$

> 🎯 **সরল ভাষায়:** আমরা যে রেখা $f(\mathbf{x})$ শিখছি, সেটা আসলে "প্রতিটা x-এর জন্য y-এর গড় কোথায়" — সেই **conditional mean**। Chapter 4 বলেছিল L2-এর best prediction = $\mathbb{E}[y\mid\mathbf{x}]$ — এখানে সেটাই মিলে গেল।

---

# ৪. Augmented notation — intercept লুকানোর কৌশল

$\beta_0$ আলাদাভাবে টানতে গেলে অংক নোংরা হয়। চালাকি: feature vector-এর সামনে একটা $1$ বসিয়ে দাও।

$$\tilde{\mathbf{x}} = (1, \mathbf{x}^\top)^\top, \qquad \tilde{\boldsymbol\beta} = (\beta_0, \boldsymbol\beta^\top)^\top$$

এখন:

$$f(\mathbf{x}) = \tilde{\mathbf{x}}^\top \tilde{\boldsymbol\beta} = 1\cdot\beta_0 + x_1\beta_1 + \dots + x_p\beta_p$$

দেখো — ওই $1$-টা $\beta_0$-কে dot product-এর ভেতরে টেনে এনেছে। আর কোনো আলাদা intercept term লাগছে না। 🎩

---

# ৫. Design matrix X

প্রতিটা observation একটা row। সব row সাজালে পাই **design matrix** $\mathbf{X}$, আকার $n \times (p+1)$:

- **প্রথম column পুরোটা ১** (intercept-এর জন্য)।
- বাকি column-গুলো feature মান।

পুরো dataset-এর সব prediction একসাথে: $\mathbf{X}\boldsymbol\theta$ (একটা $n$-length vector)।

---

# ৬. Least-squares objective

আমরা চাই prediction $\mathbf{X}\boldsymbol\theta$ আসল $\mathbf{y}$-এর কাছাকাছি। "কাছাকাছি" মাপি residual-এর বর্গের যোগফল দিয়ে:

$$\lVert \mathbf{y} - \mathbf{X}\boldsymbol\theta \rVert_2^2 = (\mathbf{y} - \mathbf{X}\boldsymbol\theta)^\top (\mathbf{y} - \mathbf{X}\boldsymbol\theta)$$

বাম পাশ = "vector-টার দৈর্ঘ্যের বর্গ"। ডান পাশ = একই জিনিস dot product আকারে (transpose × নিজে)। এটাই **Ordinary Least Squares (OLS)**।

---

# ৭. Normal equations — ধাপে ধাপে derivation

minimum পেতে হলে gradient = 0 বসাই।

**ধাপ ১** — objective খুলি:
$$(\mathbf{y} - \mathbf{X}\boldsymbol\theta)^\top(\mathbf{y} - \mathbf{X}\boldsymbol\theta) = \mathbf{y}^\top\mathbf{y} - 2\boldsymbol\theta^\top\mathbf{X}^\top\mathbf{y} + \boldsymbol\theta^\top\mathbf{X}^\top\mathbf{X}\boldsymbol\theta$$

**ধাপ ২** — $\boldsymbol\theta$-এর সাপেক্ষে derivative:
$$\nabla_{\boldsymbol\theta} = -2\mathbf{X}^\top\mathbf{y} + 2\mathbf{X}^\top\mathbf{X}\boldsymbol\theta = -2\,\mathbf{X}^\top(\mathbf{y} - \mathbf{X}\boldsymbol\theta)$$

**ধাপ ৩** — শূন্য বসাই:
$$-2\,\mathbf{X}^\top(\mathbf{y} - \mathbf{X}\boldsymbol\theta) = 0 \;\Longrightarrow\; \mathbf{X}^\top\mathbf{X}\boldsymbol\theta = \mathbf{X}^\top\mathbf{y}$$

এই শেষ সমীকরণটাই **normal equations**। নাম "normal" কারণ residual $(\mathbf{y}-\mathbf{X}\boldsymbol\theta)$ সব feature column-এর সাথে **লম্ব (orthogonal/normal)** হয়ে যায়।

---

# ৮. OLS closed-form solution

normal equations থেকে $\boldsymbol\theta$ আলাদা করি ($\mathbf{X}^\top\mathbf{X}$-এর inverse দিয়ে):

$$\boxed{\;\hat{\boldsymbol\theta} = (\mathbf{X}^\top \mathbf{X})^{-1} \mathbf{X}^\top \mathbf{y}\;}$$

> 🌟 **এটাই অধ্যায়ের হীরা।** কোনো gradient descent, কোনো iteration লাগে না — সরাসরি matrix অংকেই উত্তর। একে বলে **analytical / closed-form** solution। (Chapter 4-এর Huber, ε-insensitive loss-এ এমন সুবিধা ছিল না।)

কেন এটাই global minimum? কারণ L2 loss + linear model = **convex** problem। Convex-এ একটাই নিচু জায়গা — তাই এই stationary point-ই global minimum।

---

# ৯. কখন formula কাজ করে না (invertibility)

$(\mathbf{X}^\top\mathbf{X})^{-1}$ তখনই বের করা যায়, যখন $\mathbf{X}^\top\mathbf{X}$ **invertible**।

- Invertible হয় ⟺ $\mathbf{X}$-এর **full column rank** (feature গুলো linearly independent, এবং $n \ge p+1$)।
- দুটো feature যদি হুবহু একই/অনুপাতে থাকে (**collinear**), তখন $\mathbf{X}^\top\mathbf{X}$ **singular** → inverse নেই।
- **সমাধান — Ridge regularization:** diagonal-এ একটু $\lambda$ যোগ করো:
$$\hat{\boldsymbol\theta} = (\mathbf{X}^\top\mathbf{X} + \lambda\mathbf{I})^{-1}\mathbf{X}^\top\mathbf{y}$$
এতে matrix সবসময় invertible হয়ে যায়। (পরের chapter-এর regularization-এর সাথে যোগসূত্র।)

---

# ১০. Expectation কী — sample থেকে theoretical

এই পুরো গল্প "গড়"-এর উপর দাঁড়িয়ে। side notes মনে করিয়ে দেয়:

**Sample mean** (হাতের data থেকে) বড় $n$-এ গিয়ে **আসল theoretical mean**-এর কাছে যায় (Law of Large Numbers):

$$\frac{1}{n}\sum_{i=1}^n x_i = \bar{x} \;\longrightarrow\; \mathbb{E}[X]$$

**Theoretical mean (expectation)** কীভাবে সংজ্ঞায়িত:

| ধরন | সূত্র | মানে |
|------|------|------|
| **Discrete RV** | $\mathbb{E}[X] = \sum_{j} j \cdot \mathbb{P}(X=j)$ | প্রতিটা মান × তার সম্ভাবনা, যোগ |
| **Continuous RV** | $\mathbb{E}[X] = \int_{\mathbb{R}} x\, f(x)\, dx$ | density দিয়ে integral |
| **Joint (vector)** | $\int_{\mathcal{X}\times\mathcal{Y}} (\mathbf{x}, y)\, f(\mathbf{x}, y)\, d(\mathbf{x},y)$ | x ও y একসাথে |

> 🔑 empirical risk আসলে একটা sample average — যা theoretical risk $\mathbb{E}[L(y, f(\mathbf{x}))]$-কে estimate করে। এজন্যই expectation-এর সংজ্ঞা জানা দরকার।

---

# ১১. Absolute loss (L1) recap

নতুন data point-এর ভুল মাপতে absolute error:

$$\lvert y_i^{\text{new}} - \hat f(\mathbf{x}_i^{\text{new}}) \rvert$$

এটাই **L1 / absolute loss**, $g(x) = \lvert x \rvert$ — একটা **V আকৃতির** graph, ০-তে ধারালো কোণা।

- ০-তে **differentiable নয়** (কোণা আছে) → optimization একটু কঠিন।
- **Outlier-এ বেশি robust** (error linearly বাড়ে, L2-এর মতো quadratically না)।
- Optimal constant model = **median** (L2-তে ছিল mean)।

side notes-এর শেষ ছবিতে: একটা মসৃণ regression curve $\hat f$ এলোমেলো data-র মাঝখান দিয়ে গেছে; query point $x^{\text{new}}$-এ prediction = $\hat f(x^{\text{new}})$, আর residual = ওই বিন্দুতে আসল $y$ ও curve-এর মধ্যেকার উল্লম্ব দূরত্ব।

---

# ১২. সব মিলিয়ে একটা বড় গল্প

১) আমরা ধরলাম দুনিয়াটা **linear + Gaussian noise**।
২) তাই শেখার লক্ষ্য রেখা = **conditional mean** $f(\mathbf{x}) = \mathbb{E}[y\mid\mathbf{x}]$।
৩) Intercept সামলাতে **augmented notation** ($\tilde{\mathbf{x}}$) ব্যবহার করলাম।
৪) ভালো fit মাপতে **least-squares** objective নিলাম।
৫) gradient = 0 বসিয়ে **normal equations** পেলাম।
৬) সমাধান করে পেলাম জাদুর formula **$\hat{\boldsymbol\theta} = (\mathbf{X}^\top\mathbf{X})^{-1}\mathbf{X}^\top\mathbf{y}$** — কোনো iteration ছাড়াই।
৭) শর্ত: $\mathbf{X}^\top\mathbf{X}$ invertible (full rank); নাহলে **ridge** দিয়ে ঠিক করি।
৮) পুরোটাই দাঁড়িয়ে **expectation** (গড়)-এর ধারণার উপর।

---

# ১৩. Master Quiz-Trap Table

| বিষয় | সঠিক (TRUE) | ফাঁদ (FALSE) |
|------|-------------|--------------|
| Noise | Gaussian $\mathcal{N}(0,\sigma^2)$ | "uniform/কোনো assumption নেই" |
| $f(\mathbf{x})$ | conditional expectation $\mathbb{E}[y\mid\mathbf{x}]$ | "শুধু একটা guess" |
| Intercept | augmented $1$ দিয়ে absorb | "intercept আলাদা টানতেই হবে" |
| Objective | $\lVert\mathbf{y}-\mathbf{X}\boldsymbol\theta\rVert_2^2$ | "sum of absolute residuals" |
| Normal eq. | $\mathbf{X}^\top\mathbf{X}\boldsymbol\theta=\mathbf{X}^\top\mathbf{y}$ | "$\mathbf{X}\boldsymbol\theta=\mathbf{y}$" |
| Solution | $(\mathbf{X}^\top\mathbf{X})^{-1}\mathbf{X}^\top\mathbf{y}$ (closed-form) | "numerical optimization লাগে" |
| Invertibility | full column rank দরকার | "সবসময় invertible" |
| Minimum | convex → global | "local minimum-এ আটকে যেতে পারে" |
| L1 loss | ০-তে differentiable নয়, robust, median | "differentiable everywhere", "mean" |
| Gaussian↔loss | Gaussian ⟺ L2 ⟺ OLS = MLE | "Gaussian ⟺ L1" |

---

# ১৪. Golden Memorization Rules

1. **$\hat{\boldsymbol\theta} = (\mathbf{X}^\top\mathbf{X})^{-1}\mathbf{X}^\top\mathbf{y}$** — ঘুম থেকে উঠেই বলতে পারতে হবে।
2. **Normal equations** = gradient ০ বসানোর ফল: $\mathbf{X}^\top\mathbf{X}\boldsymbol\theta = \mathbf{X}^\top\mathbf{y}$।
3. **Gaussian noise ⟺ L2 ⟺ OLS = Maximum Likelihood।**
4. **$f(\mathbf{x}) = \mathbb{E}[y\mid\mathbf{x}]$** কারণ $\mathbb{E}[\varepsilon]=0$।
5. **Closed-form আছে** OLS-এর — iteration লাগে না (L2 + linear = convex)।
6. **Invertible দরকার** $\mathbf{X}^\top\mathbf{X}$; collinearity ভাঙে, ridge ঠিক করে।
7. **L1:** robust, median, ০-তে কোণা (differentiable নয়)।
8. **Sample mean → theoretical mean** (LLN); discrete = Σ, continuous = ∫।
