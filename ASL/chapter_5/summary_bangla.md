# অধ্যায় ৫: Linear Regression — Least Squares ও Normal Equations
### — বাংলায় সহজ ভাষায় সম্পূর্ণ ব্যাখ্যা

---

> **এই chapter-এর মূল কথা:** Chapter 4-এ আমরা risk minimization-এর সাধারণ ধারণা শিখেছি। এখানে সেটাকে **linear model**-এর জন্য concrete করা হয়েছে — L2 loss-এর জন্য হাতে-কলমে (analytically) সমাধান বের করা যায়: $\hat{\boldsymbol\theta} = (\mathbf{X}^\top\mathbf{X})^{-1}\mathbf{X}^\top\mathbf{y}$। এই formula, normal equations, আর expectation-এর সংজ্ঞা — মাথায় গেঁথে নাও।

---

## ১. Gaussian noise সহ Linear Model

ধরে নিই data এভাবে তৈরি হয়:

$$y = \beta_0 + \mathbf{x}^\top \boldsymbol\beta + \varepsilon, \qquad \varepsilon \sim \mathcal{N}(0, \sigma^2)$$

- $\beta_0$ = **intercept** (bias, যেখানে রেখা y-অক্ষ কাটে)।
- $\boldsymbol\beta$ = **slope** (প্রতিটা feature-এর জন্য একটা coefficient)।
- $\varepsilon$ = **noise / error** — এখানে ধরা হয়েছে **Gaussian**, mean ০, variance $\sigma^2$।

সব parameter একসাথে: $\boldsymbol\theta = (\beta_0, \boldsymbol\beta^\top)^\top$।

> Gaussian noise ধরলেই **L2 loss** maximum-likelihood loss হয়ে যায় (Chapter 4: Gaussian ⟺ L2)।

---

## ২. Regression function = Conditional Expectation

$\mathbb{E}[\varepsilon] = 0$ হওয়ায়:

$$\mathbb{E}[y \mid \mathbf{x}] = \beta_0 + \mathbf{x}^\top \boldsymbol\beta = f(\mathbf{x})$$

> **মূল কথা:** আমরা যে function $f(\mathbf{x})$ শিখতে চাই, সেটা আসলে $y$-এর **conditional mean** (x দেওয়া থাকলে y-এর গড়)।

---

## ৩. Augmented (Homogeneous) Notation

Intercept-টাকে dot product-এর ভেতরে ঢোকানোর জন্য feature ও parameter-এর শুরুতে একটা $1$ যোগ করি:

$$\tilde{\mathbf{x}} = (1, \mathbf{x}^\top)^\top, \qquad \tilde{\boldsymbol\beta} = (\beta_0, \boldsymbol\beta^\top)^\top$$

তাহলে model সংক্ষেপে:

$$f(\mathbf{x}) = \tilde{\mathbf{x}}^\top \tilde{\boldsymbol\beta}$$

সব $n$টা observation সাজিয়ে পাই **design matrix** $\mathbf{X}$ ($n \times (p+1)$, প্রথম column পুরোটা ১)। পুরো dataset-এর prediction = $\mathbf{X}\boldsymbol\theta$।

---

## ৪. Least-Squares Objective

L2 loss দিয়ে empirical risk = residual-এর বর্গের যোগফল, যা Euclidean norm দিয়ে লেখা যায়:

$$\lVert \mathbf{y} - \mathbf{X}\boldsymbol\theta \rVert_2^2 = (\mathbf{y} - \mathbf{X}\boldsymbol\theta)^\top (\mathbf{y} - \mathbf{X}\boldsymbol\theta)$$

এটাই **Ordinary Least Squares (OLS)**। এটাকে $\boldsymbol\theta$-এর সাপেক্ষে minimize করি।

---

## ৫. Analytical Solution — Normal Equations

Chapter 4-এর কিছু loss (Huber, ε-insensitive) numerical optimization লাগত। কিন্তু OLS-এর **closed-form (হাতে-কলমে) সমাধান** আছে।

Gradient শূন্য বসিয়ে:

$$\nabla_{\boldsymbol\theta}\, (\mathbf{y} - \mathbf{X}\boldsymbol\theta)^\top(\mathbf{y} - \mathbf{X}\boldsymbol\theta) = -2\,\mathbf{X}^\top(\mathbf{y} - \mathbf{X}\boldsymbol\theta) \stackrel{!}{=} 0$$

এতে পাই **normal equations**:

$$\mathbf{X}^\top \mathbf{X}\, \boldsymbol\theta = \mathbf{X}^\top \mathbf{y}$$

সমাধান করে ($\mathbf{X}^\top\mathbf{X}$ invertible ধরে):

$$\boxed{\;\hat{\boldsymbol\theta} = (\mathbf{X}^\top \mathbf{X})^{-1} \mathbf{X}^\top \mathbf{y}\;}$$

**মনে রাখার শর্ত:**
- $\mathbf{X}^\top\mathbf{X}$ invertible হয় তখনই, যখন $\mathbf{X}$-এর **full column rank** থাকে (feature গুলো linearly independent)।
- Problem-টা **convex** (L2 + linear model) → এই একটাই সমাধানই **global minimum**।
- Feature গুলো collinear হলে $\mathbf{X}^\top\mathbf{X}$ singular → inverse থাকে না → এজন্যই **regularization** (ridge: $(\mathbf{X}^\top\mathbf{X} + \lambda \mathbf{I})^{-1}\mathbf{X}^\top\mathbf{y}$) দরকার হয়।

---

## ৬. Expectation কী? Sample mean → Theoretical mean

Sample (গড়) mean বড় $n$-এ গিয়ে **আসল theoretical mean**-এর কাছে যায় (Law of Large Numbers):

$$\frac{1}{n}\sum_{i=1}^n x_i = \bar{x} \;\longrightarrow\; \mathbb{E}[X]$$

Theoretical mean (expectation)-এর সংজ্ঞা:

- **Discrete RV:** $\;\mathbb{E}[X] = \sum_{j} j \cdot \mathbb{P}(X = j)$
- **Continuous RV** (density $f$): $\;\mathbb{E}[X] = \int_{\mathbb{R}} x\, f(x)\, dx$
- **Joint (vector):** $\;\mathbb{E}[(\mathbf{x}, y)] = \int_{\mathcal{X} \times \mathcal{Y}} (\mathbf{x}, y)\, f(\mathbf{x}, y)\, d(\mathbf{x}, y)$

> **কেন গুরুত্বপূর্ণ:** empirical risk একটা sample average যা আসল (theoretical) risk $\mathbb{E}[L(y, f(\mathbf{x}))]$-কে estimate করে।

---

## ৭. Absolute Loss (L1) — সংক্ষিপ্ত recap

নতুন observation-এর জন্য absolute error:

$$\lvert y_i^{\text{new}} - \hat f(\mathbf{x}_i^{\text{new}}) \rvert$$

এটাই **L1 / absolute loss**, $g(x) = \lvert x \rvert$:

- $0$-তে একটা **কোণা (kink)** → $0$-তে **differentiable নয়**।
- L2-এর চেয়ে **outlier-এ বেশি robust** (error linearly বাড়ে, quadratically না)।
- Optimal constant model = **median** (L2-তে ছিল mean)।

---

## ব্যাখ্যা: Quiz-এ যেসব ফাঁদ আসে

- Model: $y = \beta_0 + \mathbf{x}^\top\boldsymbol\beta + \varepsilon$, noise **Gaussian** $\mathcal{N}(0,\sigma^2)$।
- $\mathbb{E}[\varepsilon]=0$ বলে regression function = **conditional expectation** $\mathbb{E}[y\mid\mathbf{x}]=f(\mathbf{x})$।
- Augmented: $\tilde{\mathbf{x}}=(1,\mathbf{x}^\top)^\top$, $f(\mathbf{x})=\tilde{\mathbf{x}}^\top\tilde{\boldsymbol\beta}$ — শুরুর $1$ intercept ধরে রাখে।
- OLS objective = $\lVert \mathbf{y}-\mathbf{X}\boldsymbol\theta\rVert_2^2$।
- **Normal equations:** $\mathbf{X}^\top\mathbf{X}\boldsymbol\theta = \mathbf{X}^\top\mathbf{y}$।
- **Closed-form:** $\hat{\boldsymbol\theta} = (\mathbf{X}^\top\mathbf{X})^{-1}\mathbf{X}^\top\mathbf{y}$ — iteration লাগে না।
- $\mathbf{X}^\top\mathbf{X}$ **invertible** দরকার (full column rank); convex → **global** minimum।
- Sample mean $\bar x \to \mathbb{E}[X]$; discrete $\sum j\,\mathbb{P}(X=j)$, continuous $\int x f(x)\,dx$।
- **L1 / absolute loss**: robust, $0$-তে **differentiable নয়**, optimal constant = **median**।
- **Gaussian noise ⟺ L2-loss** → OLS-ই Gaussian errors-এর অধীনে maximum-likelihood fit।
