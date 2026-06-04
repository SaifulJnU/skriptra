# 🤿 অধ্যায় ৫: Linear Regression — Least Squares & Normal Equations — Deep Dive
## — বাংলায়, একদম noob-এর জন্য, math + অঙ্কে জোর

---

> **হ্যালো বন্ধু!** 👋 এই অধ্যায়ের প্রাণ একটাই জাদুর সূত্র: $\hat{\boldsymbol\theta} = (\mathbf{X}^\top\mathbf{X})^{-1}\mathbf{X}^\top\mathbf{y}$। এই deep dive-এ আমরা এটা **কোথা থেকে এলো** ধাপে ধাপে বুঝব, matrix-চিহ্নগুলো ভাঙব, আর অঙ্ক কষব।
>
> প্রতিটা concept-এ: 🔊 উচ্চারণ · 👶 গল্প · 🧮 সূত্র ভেঙে · ✅❌ কেন/কেন নয় · 🧠 টোটকা · শেষে ✍️ অঙ্ক+সমাধান।
> (গল্পভিত্তিক বিস্তারিত চাইলে দেখো `summary_bangla_by_opus.md`; এই ফাইল math + problem-কেন্দ্রিক।)

---

# 📚 সূচিপত্র
1. [Gaussian noise সহ Linear Model](#১-gaussian-noise-সহ-linear-model)
2. [f(x) = E[y|x] কেন](#২-fx--eyx-কেন)
3. [Augmented Notation](#৩-augmented-notation)
4. [Least-Squares Objective](#৪-least-squares-objective)
5. [Normal Equations — derivation](#৫-normal-equations--derivation)
6. [OLS Closed-form Solution](#৬-ols-closed-form-solution)
7. [Invertibility — কখন ভাঙে](#৭-invertibility--কখন-ভাঙে)
8. [Expectation কী](#৮-expectation-কী)
9. [🧠 Master Table](#৯-master-table)
10. [✍️ অঙ্ক ও সমাধান](#১০-অঙ্ক-ও-সমাধান)

---

# ১. Gaussian noise সহ Linear Model

🔊 **উচ্চারণ:** *লিনিয়ার মডেল, গাউসিয়ান নয়েজ*

🧮 **সূত্র ভেঙে:**
$$y = \beta_0 + \mathbf{x}^\top\boldsymbol\beta + \varepsilon, \qquad \varepsilon \sim \mathcal{N}(0, \sigma^2)$$
- $\beta_0$ = intercept (base মান)। $\boldsymbol\beta$ = slope (feature-এর ওজন)।
- $\varepsilon$ = noise, Gaussian, গড় ০।
- সব parameter: $\boldsymbol\theta = (\beta_0, \boldsymbol\beta^\top)^\top$।

✅❌ **Gaussian ধরা কেন গুরুত্বপূর্ণ?** কারণ **Gaussian noise ⟺ L2 loss** → তাই square-error দিয়ে fit করা মানেই MLE করা, আর তার সুন্দর closed-form সমাধান আছে।

🧠 **টোটকা:** "Gaussian + linear = least squares।"

---

# ২. f(x) = E[y|x] কেন

🧮 **derivation:** $\mathbb{E}[\varepsilon]=0$ হওয়ায়
$$\mathbb{E}[y\mid\mathbf{x}] = \beta_0 + \mathbf{x}^\top\boldsymbol\beta + \underbrace{\mathbb{E}[\varepsilon]}_{=0} = f(\mathbf{x})$$

✅❌ **আমরা যা শিখছি সেটা কী?** ✅ $y$-এর **conditional mean**। কারণ noise গড়ে ০, তাই গড় নিলে শুধু রেখাটা থাকে। (Chapter 4-এর "L2 best = $\mathbb{E}[y\mid\mathbf{x}]$"-এর সাথে মিলে যায়।)

🧠 **টোটকা:** "Regression রেখা = প্রতি x-এ y-এর গড়।"

---

# ৩. Augmented Notation

🔊 **উচ্চারণ:** *অগমেন্টেড নোটেশন*

🧮 **সূত্র:** feature ও parameter-এর সামনে $1$ বসাও:
$$\tilde{\mathbf{x}} = (1, \mathbf{x}^\top)^\top, \quad \tilde{\boldsymbol\beta} = (\beta_0, \boldsymbol\beta^\top)^\top \;\Rightarrow\; f(\mathbf{x}) = \tilde{\mathbf{x}}^\top\tilde{\boldsymbol\beta}$$

✅❌ **ওই $1$ কেন?** যাতে intercept $\beta_0$ dot product-এর ভেতরে ঢুকে যায় — আলাদা term টানতে হয় না। Design matrix $\mathbf{X}$-এর **প্রথম column পুরোটা ১**।

🧠 **টোটকা:** "শুরুর ১ = intercept-এর জায়গা।"

---

# ৪. Least-Squares Objective

🔊 **উচ্চারণ:** *লিস্ট স্কয়ারস*

🧮 **সূত্র ভেঙে:**
$$\lVert \mathbf{y} - \mathbf{X}\boldsymbol\theta\rVert_2^2 = (\mathbf{y} - \mathbf{X}\boldsymbol\theta)^\top(\mathbf{y} - \mathbf{X}\boldsymbol\theta)$$
- বাম: vector-টার দৈর্ঘ্যের বর্গ। ডান: একই জিনিস dot product আকারে।
- এটাই residual-এর বর্গের যোগফল (SSE)।

✅❌ **এটা কি L1 (পরম-মান)?** ❌ না — **squared** residual। OLS = Ordinary Least **Squares**।

🧠 **টোটকা:** "OLS = বর্গ-residual-এর যোগ minimize।"

---

# ৫. Normal Equations — derivation

🔊 **উচ্চারণ:** *নরমাল ইকুয়েশনস*

🧮 **ধাপে ধাপে:**

**ধাপ ১** — খুলি: $(\mathbf{y}-\mathbf{X}\boldsymbol\theta)^\top(\mathbf{y}-\mathbf{X}\boldsymbol\theta) = \mathbf{y}^\top\mathbf{y} - 2\boldsymbol\theta^\top\mathbf{X}^\top\mathbf{y} + \boldsymbol\theta^\top\mathbf{X}^\top\mathbf{X}\boldsymbol\theta$।

**ধাপ ২** — derivative: $\nabla_{\boldsymbol\theta} = -2\mathbf{X}^\top\mathbf{y} + 2\mathbf{X}^\top\mathbf{X}\boldsymbol\theta = -2\mathbf{X}^\top(\mathbf{y}-\mathbf{X}\boldsymbol\theta)$।

**ধাপ ৩** — শূন্য বসাই: $\mathbf{X}^\top\mathbf{X}\boldsymbol\theta = \mathbf{X}^\top\mathbf{y}$ ← **normal equations**।

✅❌ **নাম "normal" কেন?** কারণ residual $(\mathbf{y}-\mathbf{X}\boldsymbol\theta)$ সব feature column-এর সাথে **লম্ব (orthogonal/normal)** হয়ে যায়।

🧠 **টোটকা:** "gradient = 0 → $\mathbf{X}^\top\mathbf{X}\boldsymbol\theta = \mathbf{X}^\top\mathbf{y}$।"

---

# ৬. OLS Closed-form Solution

🧮 **সূত্র:**
$$\boxed{\;\hat{\boldsymbol\theta} = (\mathbf{X}^\top\mathbf{X})^{-1}\mathbf{X}^\top\mathbf{y}\;}$$

✅❌ **iteration লাগে?** ❌ না — সরাসরি matrix অংকেই উত্তর (**closed-form / analytical**)। **কেন এটা global minimum?** কারণ L2 + linear = **convex** problem; convex-এ একটাই তলা।

🧠 **টোটকা:** "OLS = এক ধাপে উত্তর, gradient descent লাগে না।"

---

# ৭. Invertibility — কখন ভাঙে

🧮 **শর্ত:** $(\mathbf{X}^\top\mathbf{X})^{-1}$ তখনই আছে, যখন $\mathbf{X}^\top\mathbf{X}$ invertible ⟺ $\mathbf{X}$-এর **full column rank** (feature গুলো linearly independent, $n \ge p+1$)।

✅❌ **কখন ভাঙে?** ✅ feature গুলো **collinear** হলে (একটা আরেকটার গুণিতক) → $\mathbf{X}^\top\mathbf{X}$ singular → inverse নেই।
✅❌ **সমাধান?** **Ridge:** $(\mathbf{X}^\top\mathbf{X} + \lambda\mathbf{I})^{-1}\mathbf{X}^\top\mathbf{y}$ — diagonal-এ $\lambda$ যোগ করলে সবসময় invertible।

🧠 **টোটকা:** "Collinear ভাঙে, ridge জোড়া লাগায়।"

---

# ৮. Expectation কী

🧮 **সূত্র:**
- Sample mean → theoretical: $\frac{1}{n}\sum x_i = \bar x \to \mathbb{E}[X]$ (LLN)।
- Discrete: $\mathbb{E}[X] = \sum_j j\,\mathbb{P}(X=j)$।
- Continuous: $\mathbb{E}[X] = \int x f(x)\,dx$।

✅❌ **empirical risk আসলে কী?** একটা sample average যা true (theoretical) risk $\mathbb{E}[L]$-কে estimate করে — তাই expectation-এর সংজ্ঞা জানা জরুরি।

🧠 **টোটকা:** "Sample mean বড় $n$-এ theoretical mean।"

---

# ৯. Master Table

| বিষয় | সঠিক ✅ | ফাঁদ ❌ |
|------|---------|---------|
| Noise | Gaussian $\mathcal{N}(0,\sigma^2)$ | "assumption নেই" |
| $f(\mathbf{x})$ | $\mathbb{E}[y\mid\mathbf{x}]$ | "শুধু guess" |
| Intercept | augmented $1$ | "আলাদা টানতেই হবে" |
| Objective | $\lVert\mathbf{y}-\mathbf{X}\boldsymbol\theta\rVert_2^2$ | "sum of absolute" |
| Normal eq. | $\mathbf{X}^\top\mathbf{X}\boldsymbol\theta=\mathbf{X}^\top\mathbf{y}$ | "$\mathbf{X}\boldsymbol\theta=\mathbf{y}$" |
| Solution | $(\mathbf{X}^\top\mathbf{X})^{-1}\mathbf{X}^\top\mathbf{y}$ | "numerical লাগে" |
| Invertibility | full column rank | "সবসময় invertible" |
| Minimum | convex → global | "local-এ আটকায়" |
| Gaussian↔loss | ⟺ L2 ⟺ OLS=MLE | "↔ L1" |

---

# ১০. অঙ্ক ও সমাধান

---

### 🧩 সমস্যা ১ — Augmented vector
$\mathbf{x} = (2, 5)^\top$। augmented $\tilde{\mathbf{x}}$ লেখো। যদি $\tilde{\boldsymbol\beta} = (1, 3, -2)^\top$ হয়, $f(\mathbf{x})$ কত?

**সমাধান:** $\tilde{\mathbf{x}} = (1, 2, 5)^\top$।
$f(\mathbf{x}) = \tilde{\mathbf{x}}^\top\tilde{\boldsymbol\beta} = 1\cdot1 + 2\cdot3 + 5\cdot(-2) = 1 + 6 - 10 = -3$।

---

### 🧩 সমস্যা ২ — Normal equations সাজানো
$\mathbf{X} = \begin{bmatrix}1 & 1\\ 1 & 2\\ 1 & 3\end{bmatrix}$, $\mathbf{y} = (1, 2, 2)^\top$। $\mathbf{X}^\top\mathbf{X}$ ও $\mathbf{X}^\top\mathbf{y}$ বের করো।

**সমাধান:**
$$\mathbf{X}^\top\mathbf{X} = \begin{bmatrix}3 & 6\\ 6 & 14\end{bmatrix}, \qquad \mathbf{X}^\top\mathbf{y} = \begin{bmatrix}1+2+2\\ 1+4+6\end{bmatrix} = \begin{bmatrix}5\\ 11\end{bmatrix}$$
(যাচাই: $\sum 1 = 3$, $\sum x = 6$, $\sum x^2 = 1+4+9 = 14$ ✅)

---

### 🧩 সমস্যা ৩ — OLS সমাধান (সমস্যা ২ চালিয়ে)
উপরের normal equations $\begin{bmatrix}3&6\\6&14\end{bmatrix}\boldsymbol\theta = \begin{bmatrix}5\\11\end{bmatrix}$ সমাধান করো।

**সমাধান:** $2\times2$ inverse: $\det = 3\cdot14 - 6\cdot6 = 42-36 = 6$।
$$(\mathbf{X}^\top\mathbf{X})^{-1} = \frac{1}{6}\begin{bmatrix}14 & -6\\ -6 & 3\end{bmatrix}$$
$$\hat{\boldsymbol\theta} = \frac{1}{6}\begin{bmatrix}14 & -6\\ -6 & 3\end{bmatrix}\begin{bmatrix}5\\11\end{bmatrix} = \frac{1}{6}\begin{bmatrix}70-66\\ -30+33\end{bmatrix} = \frac{1}{6}\begin{bmatrix}4\\3\end{bmatrix} = \begin{bmatrix}0.667\\ 0.5\end{bmatrix}$$
অর্থাৎ intercept $\approx 0.667$, slope $= 0.5$। রেখা: $\hat y = 0.667 + 0.5x$।

---

### 🧩 সমস্যা ৪ — কেন iteration লাগে না? (why & why not)
এক বন্ধু OLS-এর জন্য gradient descent চালাচ্ছে। দরকার আছে কি?

**সমাধান:** ❌ দরকার নেই (ছোট/মাঝারি সমস্যায়)। OLS-এর **closed-form** সমাধান আছে — সরাসরি $(\mathbf{X}^\top\mathbf{X})^{-1}\mathbf{X}^\top\mathbf{y}$ দিলেই হয়। ✅ কখন GD লাগতে পারে? যখন data এত বিশাল যে $\mathbf{X}^\top\mathbf{X}$ inverse করা ব্যয়বহুল — তখন iterative পদ্ধতি কাজে দেয়।

---

### 🧩 সমস্যা ৫ — Invertibility ভাঙে
দুটো feature: $x_2 = 2x_1$ (একটা আরেকটার দ্বিগুণ)। OLS কি কাজ করবে? কেন/কেন নয়? সমাধান কী?

**সমাধান:** ❌ কাজ করবে না — feature দুটো **collinear**, তাই $\mathbf{X}$ full column rank নয়, $\mathbf{X}^\top\mathbf{X}$ singular, inverse নেই। ✅ সমাধান: একটা redundant feature বাদ দাও, অথবা **ridge** ($+\lambda\mathbf{I}$) ব্যবহার করো যা matrix-টা আবার invertible করে।

---

### 🧩 সমস্যা ৬ — Gaussian ↔ L2
কেন Gaussian noise ধরলে OLS-ই maximum-likelihood সমাধান? সংক্ষেপে ব্যাখ্যা করো।

**সমাধান:** Gaussian density-তে $\exp(-\frac{(y-f)^2}{2\sigma^2})$ আছে। সব point-এর likelihood-এর negative log নিলে পাই $\sum \frac{(y^{(i)}-f)^2}{2\sigma^2}$ + ধ্রুবক — যা square-error (L2) minimize করার সমান। তাই **Gaussian ⟺ L2 ⟺ OLS = MLE**।

---

> 🎓 **শেষ কথা:** অধ্যায় ৫ = Chapter 4-এর risk minimization-কে linear model-এ concrete করা। সোনার সূত্র: **(১)** $\hat{\boldsymbol\theta} = (\mathbf{X}^\top\mathbf{X})^{-1}\mathbf{X}^\top\mathbf{y}$ (মুখস্থ!); **(২)** normal equations = gradient 0; **(৩)** closed-form, convex → global, iteration লাগে না; **(৪)** $\mathbf{X}^\top\mathbf{X}$ invertible দরকার, collinearity ভাঙে, ridge জোড়ে; **(৫)** Gaussian↔L2↔OLS=MLE। 🍼💪
