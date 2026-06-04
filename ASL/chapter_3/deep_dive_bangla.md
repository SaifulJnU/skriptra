# 🤿 অধ্যায় ৩: Hypothesis Spaces & Capacity — Deep Dive
## — বাংলায়, একদম noob-এর জন্য, বন্ধুর মতো করে শেখানো

---

> **হ্যালো বন্ধু!** 👋 এই অধ্যায়টা ASL-এর **হৃৎপিণ্ড** — capacity, overfitting, VC dimension, bias-variance। এখানে concept বোঝা সবচেয়ে জরুরি, কারণ quiz-এ "কেন/কেন নয়" টাইপ প্রশ্ন বেশি আসে।
>
> প্রতিটা concept-এ: 🔊 উচ্চারণ · 👶 গল্প · 🧮 গণিত-চিহ্ন বোঝা · ✅❌ কেন/কেন নয় · 🧠 টোটকা · শেষে ✍️ অঙ্ক+সমাধান।

---

# 📚 সূচিপত্র
1. [Capacity, Underfitting, Overfitting](#১-capacity-underfitting-overfitting)
2. [Hypothesis Space](#২-hypothesis-space)
3. [Overfitting কীভাবে বাড়ে](#৩-overfitting-কীভাবে-বাড়ে)
4. [VC Dimension](#৪-vc-dimension)
5. [VC Bound ও PAC Learning](#৫-vc-bound-ও-pac-learning)
6. [Bias–Variance Decomposition](#৬-biasvariance-decomposition)
7. [Learning Curves](#৭-learning-curves)
8. [ML একটা Ill-Posed Problem](#৮-ml-একটা-ill-posed-problem)
9. [No Free Lunch Theorem](#৯-no-free-lunch-theorem)
10. [🧠 Master মনে রাখার Table](#১০-master-মনে-রাখার-table)
11. [✍️ অঙ্ক ও সমাধান](#১১-অঙ্ক-ও-সমাধান)

---

# ১. Capacity, Underfitting, Overfitting

🔊 **উচ্চারণ:** *ক্যাপাসিটি, আন্ডারফিটিং, ওভারফিটিং*

👶 **গল্প:** ছাত্রের কথা ভাবো —
- **Underfit:** এত কম পড়েছে যে train পরীক্ষাতেও ফেল। (model খুব সরল)
- **Overfit:** পুরো প্রশ্নব্যাংক **মুখস্থ** করেছে, কিন্তু নতুন প্রশ্নে ফেল। (model খুব নমনীয়, noise মুখস্থ)
- **Just right:** নিয়ম বুঝেছে, train ও test দুটোতেই ভালো।

🧮 **গণিত-চিহ্ন:**
- **Underfitting:** training error যথেষ্ট কম করতে ব্যর্থ।
- **Overfitting:** train ও test error-এর মধ্যে **বড় ব্যবধান**।
- **Generalization gap:** $= \text{test error} - \text{train error}$।
- **Capacity** = model কত ধরনের/কত জটিল hypothesis শিখতে পারে।

✅❌ **Capacity বাড়ালেই কি ভালো?** ❌ না! কম capacity → underfit (high bias)। বেশি capacity → overfit (high variance)। **Optimal** মাঝখানে — যেখানে test error সবচেয়ে কম।

🧠 **টোটকা:** "Train error সবসময় নামে; Test error **U-আকৃতি** (মাঝে সর্বনিম্ন)।"

---

# ২. Hypothesis Space

🔊 **উচ্চারণ:** *হাইপোথিসিস স্পেস*

👶 **গল্প:** $\mathcal{H}$ = সব অনুমোদিত model-এর "ব্যাগ"। শেখা = ব্যাগ থেকে সেরা model বেছে নেওয়া।

🧮 **গণিত-চিহ্ন:**
$$\mathcal{H} := \{ f : \mathcal{X} \to \mathbb{R}^g \mid f \text{-এর নির্দিষ্ট রূপ আছে} \}$$
প্রায়ই $\boldsymbol\theta$ দিয়ে parameterized: $f(\mathbf{x} \mid \boldsymbol\theta)$।

| Model | রূপ (functional form) |
|-------|------------------------|
| Linear regression | $f(\mathbf{x}) = \mathbf{x}^\top\boldsymbol\theta + \theta_0$ |
| Separating hyperplane | $h(\mathbf{x}) = \mathbb{1}(\mathbf{x}^\top\boldsymbol\theta - \theta_0 > 0)$ |
| Decision tree | $f(\mathbf{x}) = \sum_i c_i \mathbb{1}(\mathbf{x} \in Q_i)$ |
| Neural net | nested composition $\tau \circ \phi \circ \sigma \cdots$ |

🧮 **Neuron-এর হিসাব:** weighted sum + non-linear:
$$\sigma^{(j)}\!\big(\mathbf{w}_j^\top \mathbf{z} + b_j\big)$$

✅❌ **$f$ নাকি $h$?** discrete class বের করলে $h$ লেখো; $f$ হলো generic (score/probability/class সবই বোঝায়)।

🧠 **টোটকা:** "$\mathcal{H}$ = model-এর ব্যাগ; বড় ব্যাগ = বেশি capacity।"

---

# ৩. Overfitting কীভাবে বাড়ে

👶 **গল্প:** $\mathcal{H}$ যত বড় করবে (polynomial-এর degree বাড়াও, tree গভীর করো, NN বড় করো), model তত noise-ও মুখস্থ করতে পারবে।

🧮 **উদাহরণ — Polynomial regression** (সত্য: $y = 3x + 2x^2 + x^5 + \epsilon$):

| | Degree 1 | Degree 5 | Degree 13 |
|--|----------|----------|-----------|
| Capacity | কম (underfit) | মানানসই | বেশি (overfit) |
| Train RMSE | 3.87 | 1.23 | **0.48** |
| Test RMSE | 4.11 | **1.55** | 148.5 |

👉 খেয়াল করো: Train error সবচেয়ে কম degree 13-এ (0.48), কিন্তু Test error সেখানে **বিস্ফোরিত** (148.5)! Test সবচেয়ে কম degree 5-এ।

✅❌ **Train error 0 = সেরা model?** ❌ একদম না! Train error 0 মানে প্রায়ই **overfit** — noise মুখস্থ। (k-NN এ $k=1$ দিলে train error 0, কিন্তু test খারাপ।)

🧠 **টোটকা:** "Train error সর্বনিম্ন = সাবধান, overfit হতে পারে।"

---

# ৪. VC Dimension

🔊 **উচ্চারণ:** *ভিসি ডাইমেনশন* (Vapnik–Chervonenkis)

👶 **গল্প:** Capacity-কে সংখ্যা দিয়ে মাপার একটা উপায়। প্রশ্ন: "সর্বোচ্চ কয়টা point আমি **যেকোনোভাবে label করলেও** নিখুঁত আলাদা করতে পারব?" — সেই সর্বোচ্চ সংখ্যাই VC dimension।

🧮 **চিহ্ন:** $VC_p(\mathcal{H})$। একটা set **shattered** হয় যদি $\mathcal{H}$-এর কোনো member সব সম্ভাব্য label বিন্যাসে point গুলো নিখুঁত আলাদা করতে পারে।

✅❌ **VC = $d$ মানে কি সব $d$-point set shatter হয়?** ❌ **না!** মানে **অন্তত একটা** $d$-আকারের set shatter হয়, আর **কোনো** $d+1$-আকারের set হয় না। (বড় ফাঁদ!)

| Hypothesis class (in $\mathbb{R}^p$) | VC |
|---|---|
| $\mathbb{R}^2$-এ linear classifier | **3** |
| Homogeneous halfspace $\text{sign}(\mathbf{x}^\top\boldsymbol\theta)$ | $p$ |
| Non-homogeneous $\text{sign}(\mathbf{x}^\top\boldsymbol\theta + \theta_0)$ | $p+1$ |
| $\mathbb{R}^2$-এ axis-aligned rectangle | **4** |
| Threshold $\mathbb{1}(x \ge \theta)$ | **1** |
| 1-NN | **∞** |
| Sine $\mathbb{1}(\sin(\theta x) > 0)$ | **∞** |

✅❌ **Parameter কম = capacity কম?** ❌ সবসময় না! Sine classifier-এর **মাত্র ১টা** parameter, তবু VC **অসীম**। তাই parameter গুনে capacity বিচার করা যায় না।

🧠 **টোটকা:** "VC = সর্বোচ্চ shatter-যোগ্য point সংখ্যা; non-homo halfspace = $p+1$।"

---

# ৫. VC Bound ও PAC Learning

🧮 **গণিত-চিহ্ন:** VC dimension $d$ হলে, $1-\delta$ সম্ভাবনায়:
$$\text{err}_{\text{test}} \le \text{err}_{\text{train}} + \underbrace{\sqrt{\tfrac{1}{|\mathcal{D}_{\text{train}}|}\big[d(\log\tfrac{2|\mathcal{D}_{\text{train}}|}{d}+1) - \log\tfrac{\delta}{4}\big]}}_{\epsilon}$$

👶 **মানে:** test error ≤ train error + একটা "জরিমানা" $\epsilon$। $\epsilon$ ছোট হয় যদি **data বাড়াও** আর $d$ ছোট থাকে।

✅❌ **VC bound কি ব্যবহারিকভাবে কাজে লাগে?** ❌ খুব একটা না — এটা **ভয়ানক loose ও pessimistic** (যেকোনো $\mathbb{P}_{xy}$-র জন্য ধরতে হয়)। Neural network বাস্তবে bound-এর চেয়ে অনেক ভালো করে।
✅ **PAC (Probably Approximately Correct):** $d$ সসীম হলে data বাড়িয়ে $\delta, \epsilon$ যত ছোট চাই করা যায় → train error কম মানে test error-ও (সম্ভবত) কম।

🧠 **টোটকা:** "VC bound = loose; আসল GE মাপতে test set-ই ভালো।"

---

# ৬. Bias–Variance Decomposition

🔊 **উচ্চারণ:** *বায়াস-ভেরিয়েন্স ডিকম্পোজিশন*

👶 **গল্প:** total error = তিনটা টুকরোর যোগফল। ধরে নাও data: $y = f_{\text{true}}(\mathbf{x}) + \epsilon$, $\epsilon \sim \mathcal{N}(0, \sigma^2)$।

🧮 **গণিত-চিহ্ন (L2 loss):**
$$GE = \underbrace{\sigma^2}_{\text{data-র noise}} + \underbrace{\text{Var}(\hat f)}_{\text{inducer-এর variance}} + \underbrace{\text{Bias}^2(\hat f)}_{\text{inducer-এর bias}^2}$$

তিনটা টুকরো:
1. **$\sigma^2$ (irreducible error):** data-র নিজস্ব noise। **কোনো** model এর নিচে নামতে পারে না।
2. **Variance:** train data বদলালে prediction কতটা নড়ে → noise শেখার প্রবণতা → **overfitting**।
3. **Bias²:** model-এর ধারাবাহিকভাবে ভুল করার প্রবণতা → **underfitting**।

✅❌ **High capacity → কী?** ✅ **low bias, high variance**। Low capacity → **high bias, low variance**। **কেন?** নমনীয় model সত্যের কাছে যেতে পারে (low bias) কিন্তু data বদলালে অস্থির (high variance)।

✅❌ **$\sigma^2$ কি কমানো যায়?** ❌ **কখনো না** — এটা data-র অন্তর্নিহিত noise, irreducible।

🧠 **টোটকা:** "Bias = কম পড়া (underfit); Variance = মুখস্থ (overfit); $\sigma^2$ = চিরস্থায়ী noise।"

---

# ৭. Learning Curves

👶 **গল্প:** দুটো ছবি বোঝো —

**(ক) Error vs Capacity:**
- Capacity বাড়লে variance-error **বাড়ে**, bias-error **কমে**।
- Overfit অঞ্চলে GE বাড়ে, কারণ variance-এর বৃদ্ধি bias-এর হ্রাসের চেয়ে অনেক বড়।

**(খ) Error vs Training-set size ($n$):**
- $n$ বাড়লে variance-error **মিলিয়ে যায়**।
- GE ও train error এসে **bias-এ মিলে যায়** (noise শূন্য ধরলে)।
- Generalization gap **মুছে যায়**।

✅❌ **কখন বেশি data সাহায্য করে?** ✅ যখন সমস্যা **high variance** (overfitting) — বেশি data variance কমায়। ❌ **high bias** (underfitting) হলে বেশি data তেমন সাহায্য করে না — তখন model **আরও নমনীয়** করতে হবে।

| উপসর্গ | কারণ | প্রতিকার |
|--------|------|----------|
| High bias (underfit) | model খুব শক্ত | model নমনীয় করো |
| High variance (overfit) | model খুব নমনীয় | regularization বা বেশি data |

🧠 **টোটকা:** "বেশি data → variance মারে, bias নয়।"

---

# ৮. ML একটা Ill-Posed Problem

🔊 **উচ্চারণ:** *ইল-পোজড প্রবলেম*

👶 **গল্প:** ৪টা boolean feature-এর উপর function শিখছ। মোট সম্ভাব্য function = $2^{16} = 65536$। তোমার হাতে ৭টা example থাকলে এখনো $2^9$টা function data-র সাথে মেলে! অদেখা point-গুলোর label **যেকোনো** কিছু হতে পারে।

✅❌ **Data দেখেই কি নিশ্চিত শেখা যায়?** ❌ না — অনুমান (assumption/prior) ছাড়া অদেখা data-তে random guessing-এর চেয়ে ভালো করা **অসম্ভব**। তাই ML **ill-posed**।

🧠 **টোটকা:** "Assumption ছাড়া শেখা = এলোমেলো আন্দাজ।"

---

# ৯. No Free Lunch Theorem

🔊 **উচ্চারণ:** *নো ফ্রি লাঞ্চ* (NFL)

👶 **গল্প:** "ফ্রি দুপুরের খাবার নেই" — কোনো একটা algorithm **সব** সমস্যায় সেরা হতে পারে না।

✅❌ **একটা সমস্যায় কি A, B-কে হারাতে পারে?** ✅ হ্যাঁ, **নির্দিষ্ট** সমস্যায়। ❌ কিন্তু **সব সম্ভাব্য সমস্যায় গড় করলে** কোনো algorithm অন্যটার চেয়ে ভালো না — random guessing সহ সবাই সমান।

**মূল শিক্ষা:**
- কোনো algorithm **সর্বজনীনভাবে** সেরা নয়।
- ভালো করতে হলে সমস্যার জন্য সঠিক **prior/assumption** চাই।
- খুব নির্দিষ্ট assumption → অল্প সমস্যায় দুর্দান্ত। খুব বিস্তৃত → "সব কাজের কাজি, কোনোটিতেই ওস্তাদ না"।

🧠 **টোটকা:** "সব সমস্যায় গড়ে কেউ জেতে না — assumption-ই শক্তি।"

---

# ১০. Master মনে রাখার Table

| বিষয় | সঠিক ✅ | ফাঁদ ❌ |
|------|---------|---------|
| Underfit | high **train** error | "train-test gap" |
| Overfit | বড় **train-test gap** | "শুধু high train error" |
| Train error vs capacity | একটানা **কমে** | "U-আকৃতি" |
| Test error vs capacity | **U-আকৃতি** | "একটানা কমে" |
| High capacity | low bias, **high variance** | উল্টানো |
| VC = $d$ | **অন্তত একটা** set shatter | "সব $d$-set shatter" |
| Non-homo halfspace | VC = **$p+1$** | "$p$" |
| Sine classifier | VC = **∞** (১ parameter!) | "VC = 1" |
| $\sigma^2$ | **irreducible** | "data বাড়ালে শূন্য" |
| বেশি data | **variance** কমায় | "bias কমায়" |
| VC bound | loose ও pessimistic | "tight ও practical" |
| NFL | গড়ে কেউ জেতে না | "একটা algo সবসময় সেরা" |

---

# ১১. অঙ্ক ও সমাধান

---

### 🧩 সমস্যা ১ — Underfit নাকি Overfit?
Model A: train error 0.36, test 0.40। Model B: train 0.02, test 0.35। কোনটা underfit, কোনটা overfit?

**সমাধান:**
- A: train error-ই বেশি (0.36) → **underfit** (gap ছোট, কিন্তু train খারাপ)।
- B: train প্রায় শূন্য কিন্তু test অনেক বেশি (gap = 0.33) → **overfit**।

---

### 🧩 সমস্যা ২ — Generalization gap
Train error 0.12, test error 0.32। Gap কত? এটা কীসের ইঙ্গিত?

**সমাধান:** Gap $= 0.32 - 0.12 = 0.20$। বড় gap → **overfitting** (high variance)। প্রতিকার: regularization বা বেশি data।

---

### 🧩 সমস্যা ৩ — VC dimension
$\mathbb{R}^3$-এ non-homogeneous halfspace $h = \text{sign}(\mathbf{x}^\top\boldsymbol\theta + \theta_0)$-এর VC dimension কত?

**সমাধান:** Non-homogeneous halfspace in $\mathbb{R}^p$ → VC $= p+1 = 3+1 = \mathbf{4}$।
(Homogeneous হলে হতো $p=3$।)

---

### 🧩 সমস্যা ৪ — Bias-Variance চেনা
দুটো model: (ক) সরলরেখা দিয়ে একটা বাঁকা সম্পর্ক fit করার চেষ্টা; (খ) এমন নমনীয় model যা প্রতিবার আলাদা train set-এ **সম্পূর্ণ ভিন্ন** curve দেয়। কোনটায় high bias, কোনটায় high variance?

**সমাধান:**
- (ক) **high bias** — model খুব শক্ত, বাঁক ধরতে পারে না (underfit)।
- (খ) **high variance** — train set বদলালেই prediction বুনোভাবে বদলায় (overfit)।

---

### 🧩 সমস্যা ৫ — কেন train error 0 খারাপ হতে পারে?
k-NN এ $k=1$ দিলে train error 0 কিন্তু test error 0.33। কেন? (why & why not)

**সমাধান:** $k=1$ এ প্রতিটা train point নিজের প্রতিবেশী নিজেই → train-এ ১০০% সঠিক (error 0)। কিন্তু এতে **noise সহ সব মুখস্থ** হয়ে যায় (high variance, overfit), তাই অদেখা test-এ খারাপ। ✅ কম capacity ভালো না (underfit), ❌ অতি capacity-ও না (overfit) — মাঝামাঝি $k$ (যেমন 7) সেরা।

---

### 🧩 সমস্যা ৬ — No Free Lunch প্রয়োগ
"আমার নতুন algorithm প্রতিটা dataset-এ সেরা হবে" — দাবিটা কি সম্ভব? কেন/কেন নয়?

**সমাধান:** ❌ সম্ভব নয়। NFL theorem অনুযায়ী, সব সম্ভাব্য সমস্যায় গড় করলে কোনো algorithm অন্যটার চেয়ে ভালো নয়। একটা সেটে ভালো করলে অন্য সেটে খারাপ করতেই হবে। ভালো performance আসে সমস্যার সাথে মানানসই **assumption** থেকে, সর্বজনীন শ্রেষ্ঠত্ব থেকে নয়।

---

> 🎓 **শেষ কথা:** অধ্যায় ৩-এর মূল গল্প — **capacity-র ভারসাম্য**। কম = underfit (high bias), বেশি = overfit (high variance), মাঝে = সেরা। চারটা সোনা: **(১)** Test error U-আকৃতি, train একটানা নামে; **(২)** VC = অন্তত একটা set shatter (সব নয়), parameter ≠ capacity; **(৩)** $\sigma^2$ irreducible, বেশি data শুধু variance মারে; **(৪)** NFL — assumption ছাড়া কেউ জেতে না। 🍼💪
