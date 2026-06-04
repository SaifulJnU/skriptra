# অধ্যায় ৩: Hypothesis Spaces and Capacity
## — Opus-এর বিস্তারিত বাংলা ব্যাখ্যা (একদম শিশুর জন্য)

---

> **এই নোটটা কাদের জন্য?**
> তোমার জন্য, যে capacity, overfitting, VC dimension আর bias-variance আগে গভীরভাবে বোঝোনি। ধরে নিচ্ছি তুমি Chapter 1 ও 2 পড়ে এসেছ — মানে regression, classification, loss, hypothesis space-এর প্রাথমিক ধারণা আছে। এই অধ্যায়ের নতুন concept গুলো (capacity, overfitting/underfitting, VC dimension, shattering, bias-variance decomposition, learning curves, No Free Lunch) — সব ভেঙে ভেঙে বলব।

> **কীভাবে পড়বে?**
> ১) প্রতিটা section পড়ার পর একটু থেমে নিজেকে জিজ্ঞেস করো: "এটা কি ক্লাস ৮-এর ছাত্রকেও বোঝাতে পারব?"
> ২) Quiz-trap গুলো লাল কালি দিয়ে highlight করো — পরীক্ষায় ১০টার মধ্যে ৭টা সরাসরি ওখান থেকে আসবে।
> ৩) **VC dimension** এবং **Bias-Variance** — এই দুই জায়গায় সবচেয়ে বেশি ভুল হয়। সেগুলোর জন্য আলাদা সময় নাও।

---

# 📚 সূচিপত্র (Table of Contents)

1. [Capacity — Model কতটা "বুদ্ধিমান"](#১-capacity--model-কতটা-বুদ্ধিমান)
2. [Underfitting ও Overfitting](#২-underfitting-ও-overfitting)
3. [Hypothesis Space আসলে কী?](#৩-hypothesis-space-আসলে-কী)
4. [Hypothesis Space-এর উদাহরণ](#৪-hypothesis-space-এর-উদাহরণ)
5. [Overfitting কীভাবে ঘটে — তিনটা উদাহরণ](#৫-overfitting-কীভাবে-ঘটে--তিনটা-উদাহরণ)
6. [VC Dimension — Complexity মাপা](#৬-vc-dimension--complexity-মাপা)
7. [Shattering ভালো করে বোঝা](#৭-shattering-ভালো-করে-বোঝা)
8. [VC Dimension-এর উদাহরণগুলো](#৮-vc-dimension-এর-উদাহরণগুলো)
9. [VC Bound ও PAC Learning](#৯-vc-bound-ও-pac-learning)
10. [Bias-Variance Decomposition](#১০-bias-variance-decomposition)
11. [তিনটা Error Term বিস্তারিত](#১১-তিনটা-error-term-বিস্তারিত)
12. [Learning Curves](#১২-learning-curves)
13. [ML একটা Ill-Posed Problem](#১৩-ml-একটা-ill-posed-problem)
14. [No Free Lunch Theorem](#১৪-no-free-lunch-theorem)
15. [সব মিলিয়ে একটা বড় গল্প](#১৫-সব-মিলিয়ে-একটা-বড়-গল্প)
16. [Master Quiz-Trap Table](#১৬-master-quiz-trap-table)
17. [Golden Memorization Rules](#১৭-golden-memorization-rules)
18. [পরীক্ষার আগে করণীয়](#১৮-পরীক্ষার-আগে-করণীয়)

---

# ১. Capacity — Model কতটা "বুদ্ধিমান"

## 🎓 গল্প: তিন ধরনের ছাত্র

ধরো পরীক্ষার আগে তিনজন ছাত্র পড়ছে:

**ছাত্র A — অলস (Low Capacity):**
শুধু একটা সূত্র মুখস্থ করেছে। বাসার practice problem-ও পারে না, পরীক্ষাতেও পারে না। তার "বুদ্ধির ক্ষমতা" কম — জটিল কিছু ধরতে পারে না।

**ছাত্র B — চালাক কিন্তু সুষম (Right Capacity):**
বুঝে বুঝে পড়েছে। যথেষ্ট জটিল সমস্যা ধরতে পারে, আবার অতিরিক্ত মুখস্থও করেনি। Practice-ও পারে, নতুন প্রশ্নও পারে।

**ছাত্র C — অতি-মুখস্থবাজ (High Capacity):**
সব practice problem হুবহু মুখস্থ করে ফেলেছে — এমনকি প্রশ্নের ছাপার ভুলগুলোও! বাসায় ১০০% পারে, কিন্তু পরীক্ষায় একটু ঘুরিয়ে প্রশ্ন দিলেই গুবলেট।

## 📖 Capacity-এর সংজ্ঞা

**Capacity** = একটা model কী ধরনের ও কত জটিল hypothesis শিখতে পারে তার ক্ষমতা।

- একটা learner-এর performance নির্ভর করে দুটো জিনিসের উপর:
  1. **Training error কমানো** (বাসার practice পারা)
  2. **Generalize করা** — নতুন (unseen) data-তে ভালো করা (পরীক্ষায় পারা)

- **Low capacity:** শুধু কয়েকটা **সহজ** hypothesis শিখতে পারে।
- **High capacity:** অনেক, এমনকি **জটিল** hypothesis শিখতে পারে।

## 🎯 মূল কথা

> Model কতটা over বা underfit করবে — সেটা পুরোপুরি তার **capacity**-র উপর নির্ভর করে। আর test error সবচেয়ে কম হয় যখন model-এর **ঠিক ঠিক (optimal) capacity** থাকে — না বেশি, না কম।

---

# ২. Underfitting ও Overfitting

## 📖 সংজ্ঞা (এগুলো হুবহু মনে রাখো)

**Underfitting:**
> যথেষ্ট কম training error অর্জন করতে **না পারা।**
> মানে — model এতই সরল যে training data-ই ঠিকমতো শিখতে পারছে না। (ছাত্র A)

**Overfitting:**
> Training error আর test error-এর মধ্যে **বড় পার্থক্য।**
> মানে — model training-এ দারুণ কিন্তু নতুন data-তে খারাপ। (ছাত্র C)

**Generalization Gap:**
> Training error আর generalization (test) error-এর মধ্যেকার **ফাঁক**।

## 📊 বিখ্যাত Capacity গ্রাফ

```
Error
  │ \                              ___ Generalization (test) error
  │  \                            /
  │   \  Underfitting │ Overfitting
  │    \    zone       │   zone   /
  │     \______________│________ /     } ← Generalization Gap
  │                    │\______ /  ____ Training error
  │____________________│_____________________→ Capacity
                  Optimal Capacity
```

এই গ্রাফ থেকে যা শিখবে:
- **বাঁ দিকে (low capacity):** training error ও test error দুটোই বেশি → **underfitting**।
- **ডান দিকে (high capacity):** training error প্রায় শূন্য, কিন্তু test error বাড়ছে → **overfitting**, gap বড়।
- **মাঝখানে (optimal):** test error সবচেয়ে কম।

## ⚠️ Quiz Traps

- "Test error highest capacity-তে minimum" → **FALSE** (optimal/মাঝামাঝি capacity-তে minimum)
- "Underfitting মানে train-test gap বড়" → **FALSE** (সেটা overfitting; underfitting = training error-ই বেশি)
- "Low capacity model জটিল hypothesis শিখতে পারে" → **FALSE** (শুধু সহজগুলো)

---

# ৩. Hypothesis Space আসলে কী?

## 🧩 গল্প: দর্জির কাছে জামার নকশা

ধরো তুমি দর্জির কাছে গেছ। দর্জির কাছে কিছু নির্দিষ্ট **নকশার সেট** আছে — সে এই সেটের মধ্যেই জামা বানাতে পারবে। তুমি যত পছন্দই করো, সে তার নকশা-সেটের বাইরে যেতে পারবে না।

**Hypothesis space** ঠিক এই "নকশার সেট" — model যেসব function বেছে নিতে পারে তাদের সমষ্টি।

## 📖 Formal Definition

একটা learner-এর **representation** = তার **hypothesis space** $\mathcal{H}$।

$$\mathcal{H} := \{ f : \mathcal{X} \to \mathbb{R}^g \mid f \text{ এর একটা নির্দিষ্ট form আছে} \}$$

**ভেঙে বুঝি:**
- $\mathcal{X}$ = input space (feature গুলো)
- $\mathbb{R}^g$ = output (score/probability)
- "নির্দিষ্ট form" = function-টা কেমন দেখতে (linear, polynomial, tree ইত্যাদি)

প্রায়ই $f$ একটা parameter $\boldsymbol\theta \in \Theta$ দিয়ে লেখা হয়:
$$f(\mathbf{x}) = f(\mathbf{x} \mid \boldsymbol\theta)$$

## 📝 $f$ বনাম $h$ — একটা ছোট Note

> যখন স্পষ্টভাবে **hard classifier** (discrete class output করে) বোঝাই, তখন $f$-এর বদলে $h$ লিখি। কিন্তু সাধারণভাবে আমরা $f$ লিখি — যা discrete class, score, ও probability — সবকিছু একসাথে বোঝায়।

---

# ৪. Hypothesis Space-এর উদাহরণ

এই পাঁচটা হলো course-এর গুরুত্বপূর্ণ উদাহরণ। প্রতিটার "form" বোঝো।

## ① Linear Regression
$$f(\mathbf{x} \mid \theta_0, \boldsymbol\theta) = \mathbf{x}^\top\boldsymbol\theta + \theta_0, \quad \boldsymbol\theta \in \mathbb{R}^p, \; \theta_0 \in \mathbb{R}$$
সরল রেখা (বা hyperplane) দিয়ে data fit করে।

## ② Separating Hyperplane (Classification)
$$h(\mathbf{x} \mid \theta_0, \boldsymbol\theta) = \mathbb{1}(\mathbf{x}^\top\boldsymbol\theta - \theta_0 > 0)$$
একটা সরল রেখা দিয়ে দুই class আলাদা করে।

## ③ Decision Trees
$$f(\mathbf{x}) = \sum_{i=1}^m c_i \mathbb{1}(\mathbf{x} \in Q_i)$$
Feature space-কে **axis-aligned rectangle** (অক্ষ-সমান্তরাল আয়তক্ষেত্র) এ recursively ভাগ করে। প্রতিটা region $Q_i$-তে একটা constant value $c_i$।

## ④ Ensemble Methods
$$f(\mathbf{x} \mid \beta^{[l]}) = \sum_{l=1}^m \beta^{[l]} b^{[l]}(\mathbf{x})$$
অনেকগুলো model-এর prediction একসাথে যোগ করে (aggregate)। উদাহরণ:
- Random forests
- Bagging
- Tree-based boosting

## ⑤ Neural Networks
$$f(\mathbf{x}) = \tau \circ \phi \circ \sigma^{(h)} \circ \phi^{(h)} \circ \cdots \circ \sigma^{(1)} \circ \phi^{(1)}(\mathbf{x})$$

অনেকগুলো layer-এর nested composition ($h+1$টা layer)।

**প্রতিটা neuron দুই ধাপে কাজ করে:**

ধাপ ১ — **Weighted sum** (input-গুলোর ভারিত যোগফল):
$$\phi^{(j)}(\mathbf{z}) = \mathbf{w}_j^\top \mathbf{z} + b_j$$

ধাপ ২ — **Non-linear transformation** ($\sigma$):
$$\sigma^{(j)}\!\left(\mathbf{w}_j^\top \mathbf{z} + b_j\right)$$

পুরো network = এরকম operation-এর **nested composition**।

---

# ৫. Overfitting কীভাবে ঘটে — তিনটা উদাহরণ

## কীভাবে capacity বাড়ানো হয়?

- Model-এর capacity ("complexity") বাড়ানো হয় **hypothesis space বড় করে।**
- এতে সাধারণত **learnable parameter-এর সংখ্যাও** বাড়ে।
- উদাহরণ: polynomial-এর degree বাড়ানো, tree-এর depth বাড়ানো, neural network বড় করা, আরো predictor যোগ করা।
- $\mathcal{H}$ যত বড় হয়, overfit করার প্রবণতা তত বাড়ে — model এমনকি training data-র **random quirk (এলোমেলো ভুল)** ও শিখে ফেলে, ফলে generalize করতে পারে না।

## 📊 উদাহরণ ১: Polynomial Regression

Data simulate করা হয়েছে: $y = 3x_1 + 2x_1^2 + x_1^5 + \epsilon$, যেখানে $\epsilon \sim \mathcal{N}(0, 1.25)$।

| | Degree 1 | Degree 5 | Degree 13 |
|--|----------|----------|-----------|
| অবস্থা | Underfit (low cap.) | ✅ Just right | Overfit (high cap.) |
| Training error (RMSE) | 3.87 | 1.23 | **0.48** |
| Test error (RMSE) | 4.11 | **1.55** | 148.5 |

দেখো — Degree 13-এর training error সবচেয়ে কম (0.48), কিন্তু test error **বিশাল (148.5)**! এটাই overfitting।

## 📊 উদাহরণ ২: Decision Trees

`minsplit` = একটা node split করতে কমপক্ষে কত sample লাগবে। **minsplit ছোট হলে tree বেশি গভীর হয় → capacity বেশি।**

| | minsplit 60 | minsplit 12 | minsplit 2 |
|--|-------------|-------------|------------|
| অবস্থা | Underfit | ✅ Just right | Overfit |
| Training error (Misclass.) | 0.36 | 0.12 | **0.02** |
| Test error (Misclass.) | 0.40 | **0.32** | 0.35 |

## 📊 উদাহরণ ৩: k-Nearest Neighbours

**k ছোট হলে capacity বেশি** (k=1 মানে প্রতিটা point নিজের নিকটতম একটা প্রতিবেশী দেখে — খুবই wiggly boundary)।

| | k = 20 | k = 7 | k = 1 |
|--|--------|-------|-------|
| অবস্থা | Underfit | ✅ Just right | Overfit |
| Training error (Misclass.) | 0.22 | 0.13 | **0** |
| Test error (Misclass.) | 0.40 | **0.25** | 0.33 |

k=1-এ training error **শূন্য** (প্রতিটা point নিজেই নিজের প্রতিবেশী!), কিন্তু test error বেশি।

## 🔑 সবচেয়ে গুরুত্বপূর্ণ Pattern

> Capacity বাড়লে **training error সবসময় কমতে থাকে** (highest capacity-তে সবচেয়ে কম)। কিন্তু **test error U-shaped** — মাঝামাঝি capacity-তে সবচেয়ে কম, দুই প্রান্তে বেশি।

## ⚠️ Quiz Traps

- "Capacity বাড়লে test error monotonically কমে" → **FALSE** (U-shaped)
- "k-NN-এ বড় k = বেশি capacity" → **FALSE** (ছোট k = বেশি)
- "Decision tree-এ বড় minsplit = বেশি capacity" → **FALSE** (ছোট minsplit = বেশি)
- "Polynomial-এর degree বাড়ালে capacity কমে" → **FALSE** (বাড়ে)

---

# ৬. VC Dimension — Complexity মাপা

## 🤔 প্রশ্ন: একটা hypothesis space কতটা "শক্তিশালী"?

আমরা চাই একটা সংখ্যা দিয়ে বলতে — এই function class কতটা জটিল pattern ধরতে পারে। এই সংখ্যাটাই **Vapnik–Chervonenkis (VC) dimension।**

## 📖 সংজ্ঞা

Binary-valued function-এর একটা class $\mathcal{H} = \{h : \mathcal{X} \to \{0,1\}\}$-এর **VC dimension** হলো:

> **সবচেয়ে বড় সংখ্যক point** (কোনো এক configuration-এ) যাদের $\mathcal{H}$-এর সদস্যরা **shatter** করতে পারে।

লেখা হয় $VC_p(\mathcal{H})$, যেখানে $p$ = input space-এর dimension।

---

# ৭. Shattering ভালো করে বোঝা

## 📖 Shatter মানে কী?

> একটা point-এর set "shattered" হয় যদি class-এর কোনো একটা function ওই point-গুলোকে **যেকোনো binary label assignment**-এর জন্য perfectly আলাদা করতে পারে।

## 🎨 গল্প: ৩টা বিন্দু আর একটা সরল রেখা

ℝ²-তে ৩টা বিন্দু আঁকো (একই সরলরেখায় না)। এখন প্রতিটা বিন্দুকে লাল বা নীল রং দাও — মোট $2^3 = 8$টা সম্ভাব্য রং-বিন্যাস।

একটা সরল রেখা (linear classifier) দিয়ে কি **প্রতিটা** বিন্যাসে লাল-নীল আলাদা করা যায়? — **হ্যাঁ!** তাই ৩টা বিন্দু shatter করা যায়।

কিন্তু ৪টা বিন্দু? — কিছু বিন্যাসে (যেমন XOR-এর মতো diagonal pattern) একটা সরল রেখা দিয়ে আলাদা করা **যায় না**। তাই ℝ²-এ linear classifier ৪ বিন্দু shatter করতে **পারে না**।

ফলে: $VC_2(\mathcal{H}) = 3$।

## ⚠️ সবচেয়ে বড় Conceptual Trap

> VC dimension = $d$ মানে **সব** $d$-সাইজ set shatter হয় — এটা **ভুল!**
>
> সঠিক মানে: **অন্তত একটা** $d$-সাইজ set shatter হয়, এবং **কোনো** $(d+1)$-সাইজ set shatter হয় না।

মানে এমন একটা set থাকলেই হলো — সব set না।

---

# ৮. VC Dimension-এর উদাহরণগুলো

## 📋 মুখস্থ রাখার মূল Table

| Hypothesis class (ℝᵖ-তে) | VC dimension |
|--------------------------|--------------|
| ℝ²-তে linear indicator function | **3** |
| Homogeneous halfspace $\text{sign}(\mathbf{x}^\top\boldsymbol\theta)$ | exactly **p** |
| Non-homogeneous halfspace $\text{sign}(\mathbf{x}^\top\boldsymbol\theta + \theta_0)$ | exactly **p+1** |
| ℝ²-তে axis-aligned rectangle | **4** |
| Threshold classifier $\mathbb{1}(x \ge \theta)$ | **1** |
| 1-Nearest Neighbour (k=1) | **infinite (∞)** |
| Sine classifier $\mathbb{1}(\sin(\theta x) > 0)$ | **infinite (∞)** |

## 🔍 কয়েকটা বিস্তারিত

### Homogeneous halfspace (origin দিয়ে যাওয়া hyperplane)
$$h(\mathbf{x}) = \text{sign}(\mathbf{x}^\top\boldsymbol\theta), \quad VC = p$$
**Proof-এর intuition:** standard basis vector $\mathbf{e}^{(1)}, \ldots, \mathbf{e}^{(p)}$ নাও। যেকোনো labeling $y^{(i)}$-এর জন্য $\boldsymbol\theta = (y^{(1)}, \ldots, y^{(p)})^\top$ বসালেই $h(\mathbf{e}^{(i)}) = \text{sign}(y^{(i)}) = y^{(i)}$ — তাই $p$ point shatter হয়। আর $p+1$ vector ℝᵖ-তে সবসময় linearly dependent → shatter হয় না।

### Non-homogeneous halfspace (bias সহ)
$$h(\mathbf{x}) = \text{sign}(\mathbf{x}^\top\boldsymbol\theta + \theta_0), \quad VC = p+1$$
**Trick:** $\tilde{\mathbf{x}} = (1, x_1, \ldots, x_p)$ আর $\tilde{\boldsymbol\theta} = (\theta_0, \theta_1, \ldots, \theta_p)$ নিলে যেকোনো affine function ℝᵖ-তে = homogeneous linear function ℝ^{p+1}-তে। তাই VC = $(p+1)$।

### Axis-aligned rectangle (ℝ²-এ)
VC = **4**। ৪টা point (চার দিকে) shatter করা যায়, কিন্তু ৫টা না — কারণ ৫টা point-এর মধ্যে leftmost/rightmost/top/bottom বেছে class 1 দিলে, মাঝের ৫ম point-ও rectangle-এর ভেতরে পড়ে যায়, আলাদা করা যায় না।

### 🚨 সবচেয়ে বড় Trap: Sine Classifier
$$h(x) = \mathbb{1}(\sin(\theta x) > 0)$$
এর মাত্র **একটা** parameter $\theta$, তবু VC dimension **infinite!** কারণ frequency $\theta$ যথেষ্ট বড় নিলে যেকোনো সংখ্যক point shatter করা যায়।

> **🔑 শিক্ষা:** VC dimension সাধারণত parameter সংখ্যার সাথে বাড়ে, **কিন্তু শুধু parameter সংখ্যা দিয়ে capacity বিচার করা যায় না!** Sine হলো জ্বলন্ত উদাহরণ — ১ parameter, ∞ capacity।

## ⚠️ Quiz Traps

- "VC dim d মানে সব d-point shatter হয়" → **FALSE** (অন্তত একটা)
- "Non-homogeneous halfspace-এর VC = p" → **FALSE** (**p+1**)
- "Sine classifier-এর VC = 1 কারণ এক parameter" → **FALSE** (∞)
- "Parameter সংখ্যা দিয়ে সবসময় capacity বলা যায়" → **FALSE**
- "k=1 NN-এর VC finite" → **FALSE** (∞)

---

# ৯. VC Bound ও PAC Learning

## 📖 মূল ধারণা

Training error হলো generalization (test) error-এর একটা **optimistic (আশাবাদী)** estimate — মানে training error সবসময় test error-এর চেয়ে কম দেখায় (rose-tinted glasses)।

VC dimension $d$ ব্যবহার করে test error-এর একটা **probabilistic upper bound** দেওয়া যায়। $1-\delta$ probability-তে:

$$\mathbb{P}\!\left(\text{err}_{\text{test}} \le \text{err}_{\text{train}} + \underbrace{\sqrt{\tfrac{1}{|\mathcal{D}_{\text{train}}|}\left[d\left(\log\tfrac{2|\mathcal{D}_{\text{train}}|}{d} + 1\right) - \log\tfrac{\delta}{4}\right]}}_{\epsilon}\right) = 1-\delta$$

(শর্ত: training set যথেষ্ট বড়, $d < |\mathcal{D}_{\text{train}}|$)

## 🔍 কী বলছে এই bound?

- এই bound **সব** ($|\mathcal{D}_{\text{train}}|$ সাইজের) dataset-এর উপর ধরে, যা **arbitrary** $\mathbb{P}_{xy}$ থেকে আসা।
- $d$ **finite** হলে — sample size বাড়িয়ে $\delta$ আর $\epsilon$ **দুটোই** যত খুশি ছোট করা যায়।
- **Corollary:** training error কম হলে test error-ও (probably) কম। মানে training error minimize করা মানেই hypothesis set থেকে একটা "ভালো" hypothesis বেছে নেওয়া।

## 📖 PAC — Probably Approximately Correct

> যে algorithm training error minimize করে নির্ভরযোগ্যভাবে একটা "ভালো" hypothesis বেছে নেয়, তাকে বলে **Probably Approximately Correct (PAC)** algorithm।

- **Probably** → $1-\delta$ probability-তে
- **Approximately Correct** → $\epsilon$ পরিমাণ ভুলের মধ্যে

## ⚠️ বাস্তবতা (গুরুত্বপূর্ণ)

> VC bound গুলো **খুবই loose আর pessimistic** (অতি-হতাশাবাদী)। কারণ এগুলোকে **arbitrary** $\mathbb{P}_{xy}$-এর জন্য ধরতে হয়, তাই bound কড়া করতে বিশাল training set লাগে।
>
> বাস্তবে neural network-এর মতো জটিল model এই bound-এর চেয়ে **অনেক ভালো** করে। তাই:
> - **Rademacher complexity** = বিকল্প, প্রায়ই tighter bound দেয়।
> - আসল (effective) capacity শুধু hypothesis space না, **optimizer**-এর উপরও নির্ভর করে।
> - Generalization error-এর সবচেয়ে ভালো estimate = সরাসরি **test set-এ evaluate করা**।

## ⚠️ Quiz Traps

- "VC bound tight ও optimistic" → **FALSE** (loose ও pessimistic)
- "VC bound এক নির্দিষ্ট distribution-এর জন্য" → **FALSE** (arbitrary সব distribution)
- "Training error হলো pessimistic estimate" → **FALSE** (optimistic)

---

# ১০. Bias-Variance Decomposition

এটা এই chapter-এর হৃদয়। ধীরে পড়ো।

## 📖 Setup ও Assumption

Inducer $\mathcal{I}_{L,\mathcal{O}}$-এর generalization error:
$$GE_n(\mathcal{I}_{L,\mathcal{O}}) = \mathbb{E}_{\mathcal{D}_n, xy}\left(L\left(y, \hat f_{\mathcal{D}_n}(\mathbf{x})\right)\right)$$

**ধরে নিই (assumption):** data তৈরি হয়েছে
$$y = f_{\text{true}}(\mathbf{x}) + \epsilon, \quad \epsilon \sim \mathcal{N}(0, \sigma^2), \; \mathbf{x}\text{-এর থেকে independent}$$

মানে $y \sim \mathcal{N}(f_{\text{true}}(\mathbf{x}), \sigma^2)$ — সত্যিকার function-এর চারপাশে গাউসিয়ান noise।

## 🧮 L2-loss দিয়ে ভাঙা

L2-loss $L(y, f(\mathbf{x})) = (y - f(\mathbf{x}))^2$ বসিয়ে, $(\mathbf{x}, y) \sim \mathbb{P}_{xy}$-এর উপর expectation নিলে error **তিনটা term**-এ ভেঙে যায়:

$$GE_n = \underbrace{\sigma^2}_{\text{(১) data variance}} + \underbrace{\mathbb{E}_{xy}\!\left[\text{Var}_{\mathcal{D}_n}\!\left(\hat f_{\mathcal{D}_n}(\mathbf{x})\right)\right]}_{\text{(২) inducer variance}} + \underbrace{\mathbb{E}_{xy}\!\left[\mathbb{E}^2_{\mathcal{D}_n}\!\left(f_{\text{true}}(\mathbf{x}) - \hat f_{\mathcal{D}_n}(\mathbf{x})\right)\right]}_{\text{(৩) squared bias}}$$

> এই তিন-term ভাঙনটা পরীক্ষায় চাইবেই। তিনটা term মনে রাখো: **data variance + inducer variance + squared bias**।

---

# ১১. তিনটা Error Term বিস্তারিত

## ① Variance of the Data — $\sigma^2$

- এটা data-র মধ্যেকার **noise**।
- একে বলা হয় **intrinsic / unavoidable / irreducible error** (অন্তর্নিহিত / অপরিহার্য / অপরিবর্তনীয় ভুল)।
- **🔑 যত ভালো learner-ই ব্যবহার করো না কেন, এই error-এর নিচে কখনো যেতে পারবে না।**

**গল্প:** তুমি একটা থার্মোমিটার দিয়ে জ্বর মাপছ, কিন্তু থার্মোমিটারে ±০.৫° random ত্রুটি আছে। তুমি যত চালাকই হও, এই ০.৫° ভুল দূর করতে পারবে না — এটা data-তেই আছে।

## ② Variance of the Inducer

- $\text{Var}_{\mathcal{D}_n}(\hat f_{\mathcal{D}_n}(\mathbf{x}))$ — একটা fixed test point-এ।
- মানে: **training data বদলালে prediction কতটা বদলায়?**
- Learner-এর **random জিনিস শেখার** প্রবণতা — যা real signal-এর সাথে সম্পর্কহীন → **overfitting**।

**গল্প:** ছাত্র C (মুখস্থবাজ) — তাকে আলাদা আলাদা practice set দিলে সে প্রতিবার একদম **আলাদা** "নিয়ম" বানিয়ে ফেলে। অস্থির, অনির্ভরযোগ্য → high variance।

## ③ Squared Bias

- $\mathbb{E}^2_{\mathcal{D}_n}(f_{\text{true}}(\mathbf{x}) - \hat f_{\mathcal{D}_n}(\mathbf{x}))$ — একটা fixed test point-এ।
- Learner-এর **ধারাবাহিকভাবে (consistently) ভুল করার** প্রবণতা → **underfitting**।

**গল্প:** ছাত্র A (অলস) — সে সবসময় একই ধরনের ভুল করে, কারণ তার model এত সরল যে সত্যিকার pattern ধরতেই পারে না। স্থির কিন্তু ভুল → high bias।

## 📊 Capacity-এর সাথে সম্পর্ক (অবশ্যই মনে রাখো)

| Capacity | Bias | Variance |
|----------|------|----------|
| **High** | **Low** | **High** |
| **Low** | **High** | **Low** |

> এটাই বিখ্যাত **Bias-Variance Trade-off** — একটা কমালে অন্যটা বাড়ে।

## 🎨 ছবির intuition

- **High bias:** model এত শক্ত যে data-র বাঁকা (curved) সম্পর্ক ধরতে পারে না — সরল রেখা টেনে দেয়।
- **High variance, no bias:** নীতিগতভাবে সত্যিকার pattern শিখতে পারে, কিন্তু বাস্তবে আলাদা আলাদা training set-এর জন্য **একদম আলাদা আলাদা** hypothesis output করে।

## ⚠️ Quiz Traps

- "High capacity model-এর high bias, low variance" → **FALSE** (low bias, high variance)
- "σ² (data variance) দূর করা যায়" → **FALSE** (irreducible)
- "Inducer variance underfitting বোঝায়" → **FALSE** (overfitting)
- "Squared bias overfitting বোঝায়" → **FALSE** (underfitting)

---

# ১২. Learning Curves

দুই ধরনের learning curve আছে — গুলিয়ে ফেলো না।

## 📈 Curve ১: Error vs Capacity

Capacity (model-এর জটিলতা) বাড়ালে:
- Variance-এর error **বাড়ে** ↑
- Bias-এর error **কমে** ↓
- **Overfitting region-এ:** generalization error **বাড়ে**, কারণ variance বৃদ্ধি >> bias হ্রাস।

```
Error
  │\         generalization error
  │ \       /↗ (variance বাড়ছে)
  │  \_____/
  │  bias↓  variance↑
  │________________________→ capacity
       optimal capacity
```

## 📈 Curve ২: Error vs Training-set Size

Training set (data-র পরিমাণ) বড় করলে:
- Variance-এর error **মিলিয়ে যায়** (vanishes)।
- Generalization error আর training error দুটোই **algorithm-এর bias-এ গিয়ে মেলে** (noise শূন্য ধরলে)।
- **Generalization gap-ও মিলিয়ে যায়।**

```
Error
  │\___ generalization error
  │    \____
  │    ____ ⟶ দুটো মিলে যায় (bias-এ)
  │___/ training error
  │________________________→ training set size
```

## 🛠️ Diagnosis ও সমাধান

| লক্ষণ | কারণ | সমাধান |
|-------|------|--------|
| High bias (underfit) | Model খুব শক্ত | Model **বেশি flexible** করো (বা অন্য model নাও) — *bias কমাও* |
| High variance (overfit) | Model খুব flexible | Model **কম flexible (regularization)** বা **আরো data** যোগ করো — *variance কমাও* |

## ⚠️ Quiz Traps

- "Underfitting কমাতে regularization বাড়াও" → **FALSE** (regularization variance কমায়; bias কমাতে model flexible করো)
- "Training set বড় হলে variance error বাড়ে" → **FALSE** (মিলিয়ে যায়)
- "Capacity বাড়লে bias বাড়ে" → **FALSE** (bias কমে, variance বাড়ে)

---

# ১৩. ML একটা Ill-Posed Problem

## 🧩 গল্প: ৭টা উদাহরণ দিয়ে নিয়ম শেখা

ধরো তুমি ৪টা boolean feature ($x_1, x_2, x_3, x_4$) থেকে একটা boolean output $y$ শিখতে চাও।

- ৪টা boolean feature → $2^4 = 16$টা সম্ভাব্য feature combination।
- প্রতিটাকে 0 বা 1 দিতে পারো → মোট $2^{16} = 65536$টা সম্ভাব্য function!

এখন তোমাকে মাত্র **৭টা** training example দেওয়া হলো।

## 🤔 সমস্যা

- ৭টা example দেখার পরেও বাকি $16-7 = 9$টা combination-এর label অজানা → **$2^9 = 512$টা function এখনো consistent** থাকে।
- Unseen datapoint-গুলোর **যেকোনো** label হতে পারে।

> এজন্যই **machine learning একটা ill-posed problem** — শুধু training data দিয়ে unique উত্তর নির্ধারণ করা যায় না।

## ❓ মূল প্রশ্ন

কোনো অতিরিক্ত assumption ছাড়া, একটা ML algorithm কি সত্যিই **random guessing**-এর চেয়ে ভালো করতে পারে? — সাধারণভাবে **না**।

---

# ১৪. No Free Lunch Theorem

## 📖 মূল বক্তব্য

- একটা **নির্দিষ্ট** problem-এর জন্য (target function-এর একটা নির্দিষ্ট distribution), কিছু algorithm অন্যদের চেয়ে ভালো করতে **পারে**।
- কিন্তু **সব সম্ভাব্য problem-এর উপর গড় নিলে**, **কোনো** algorithm অন্য কারো চেয়ে ভালো না — এমনকি random guessing-এর চেয়েও না।

> এটাই বিখ্যাত **No Free Lunch (NFL) theorem**।

- NFL-এর মূল কথা: যে algorithm একদল problem-এ ভালো করে, তাকে অন্য একদল problem-এ **খারাপ করতেই হবে**। "Free lunch" বলে কিছু নেই — কোথাও না কোথাও দাম দিতেই হবে।

## 🎯 Takeaway (গুরুত্বপূর্ণ)

- কোনো algorithm **সর্বজনীনভাবে (universally)** অন্য সব algorithm-এর চেয়ে ভালো না।
- একটা learning algorithm-কে একটা **নির্দিষ্ট prior**-এর সাথে মানিয়ে নিতে হয়। **Problem সম্পর্কে assumption না করলে ML-ই সম্ভব না!**
- **খুব নির্দিষ্ট** assumption → ছোট একদল problem-এ দারুণ পারফর্ম।
- **খুব broad** assumption → "jack of all trades, master of none" (সব কাজের কাজী, কোনো কাজেই দক্ষ না)।

## ⚠️ Quiz Traps

- "এমন একটা algorithm আছে যা সব problem-এ সবার চেয়ে ভালো" → **FALSE**
- "Assumption ছাড়া ML সম্ভব ও কার্যকর" → **FALSE**
- "সব problem-এর গড়ে best algorithm random guessing-কে হারায়" → **FALSE** (সমান)

---

# ১৫. সব মিলিয়ে একটা বড় গল্প

ধরো তুমি একটা company-তে নতুন employee-দের performance predict করার model বানাচ্ছ।

**Hypothesis space (Section ৩-৪):**
তুমি ঠিক করলে model কেমন হবে — linear? decision tree? neural network? এটাই তোমার "নকশার সেট"।

**Capacity ও Overfitting (Section ১-২, ৫):**
- খুব সরল model (linear, একটা feature) → underfit, কিছুই ঠিকমতো শেখে না।
- খুব জটিল model (degree-13 polynomial / k=1 NN / minsplit=2 tree) → training data হুবহু মুখস্থ করে, কিন্তু নতুন employee-তে ব্যর্থ → overfit।
- মাঝামাঝি → just right।

**VC dimension (Section ৬-৯):**
তোমার model-এর "তাত্ত্বিক ক্ষমতা" কত? linear হলে $p+1$, কিন্তু sine-এর মতো হলে ∞। বেশি VC = বেশি overfit-এর ঝুঁকি। VC bound বলে training error কম হলে test error-ও probably কম (PAC) — কিন্তু bound খুবই loose।

**Bias-Variance (Section ১০-১১):**
- সরল model → consistently ভুল (high bias)।
- জটিল model → training set বদলালেই উত্তর বদলায় (high variance)।
- আর data-তে যে noise ($\sigma^2$) আছে, সেটা কখনোই দূর হবে না।

**Learning curves (Section ১২):**
আরো data জোগাড় করলে variance কমবে, gap মিলিয়ে যাবে। কিন্তু bias কমাতে হলে model নিজেই flexible করতে হবে।

**No Free Lunch (Section ১৩-১৪):**
কোনো একটা "জাদুকরী" model নেই যা সব company, সব employee-র জন্য সেরা হবে। তোমার নির্দিষ্ট সমস্যা সম্পর্কে assumption করেই model বেছে নিতে হবে।

---

# ১৬. Master Quiz-Trap Table

| Statement | T/F | কারণ |
|-----------|-----|------|
| Underfitting = যথেষ্ট কম training error না পারা | TRUE | সংজ্ঞা |
| Overfitting = বড় train-test gap | TRUE | সংজ্ঞা |
| Low capacity model জটিল hypothesis শেখে | FALSE | শুধু সহজগুলো |
| Test error highest capacity-তে minimum | FALSE | optimal capacity-তে |
| Capacity বাড়লে training error কমতে থাকে | TRUE | — |
| Capacity বাড়লে test error monotonically কমে | FALSE | U-shaped |
| Polynomial degree বাড়ালে capacity কমে | FALSE | বাড়ে |
| k-NN: ছোট k = বেশি capacity | TRUE | — |
| Decision tree: বড় minsplit = বেশি capacity | FALSE | ছোট minsplit = বেশি |
| Decision tree axis-aligned rectangle বানায় | TRUE | — |
| Ensemble = অনেক model একসাথে | TRUE | — |
| VC dimension = function space-এর complexity | TRUE | — |
| Shatter = যেকোনো labeling perfectly আলাদা করা | TRUE | — |
| VC dim d → সব d-point shatter হয় | FALSE | অন্তত একটা set |
| ℝ²-তে linear indicator: VC = 3 | TRUE | — |
| Homogeneous halfspace: VC = p | TRUE | — |
| Non-homogeneous halfspace: VC = p | FALSE | **p+1** |
| Axis-aligned rectangle (ℝ²): VC = 4 | TRUE | — |
| Threshold classifier 1(x≥θ): VC = 1 | TRUE | — |
| Sine classifier: VC = 1 (এক parameter) | FALSE | **infinite** |
| k=1 NN: VC infinite | TRUE | — |
| Parameter সংখ্যা দিয়ে সবসময় capacity বিচার করা যায় | FALSE | sine উদাহরণ |
| Training error = optimistic estimate | TRUE | — |
| VC bound এক distribution-এর জন্য | FALSE | arbitrary সব |
| Finite d → sample বাড়িয়ে δ,ε ছোট করা যায় | TRUE | — |
| Training-error-min algorithm = PAC | TRUE | — |
| VC bound tight ও optimistic | FALSE | loose ও pessimistic |
| Rademacher complexity = tighter বিকল্প | TRUE | — |
| Data: y = f_true(x) + ε, ε~N(0,σ²) | TRUE | assumption |
| তিন term: data var + inducer var + bias² | TRUE | — |
| σ² (data variance) দূর করা যায় | FALSE | irreducible |
| Inducer variance ↔ overfitting | TRUE | — |
| Squared bias ↔ underfitting | TRUE | — |
| High capacity → high bias, low variance | FALSE | low bias, high variance |
| Capacity↑ → variance↑, bias↓ | TRUE | — |
| Training set↑ → variance error মিলিয়ে যায় | TRUE | — |
| Underfit কমাতে regularization বাড়াও | FALSE | model flexible করো |
| Overfit কমাতে regularization বা আরো data | TRUE | — |
| ML ill-posed কারণ unseen label যেকোনো হতে পারে | TRUE | — |
| NFL: সব problem-এর গড়ে কোনো algorithm সেরা না | TRUE | — |
| এমন algorithm আছে যা সব problem-এ সেরা | FALSE | NFL |
| Assumption ছাড়া ML সম্ভব | FALSE | ill-posed |

---

# ১৭. Golden Memorization Rules

### Rule 1: Underfit vs Overfit
- **Under**fit → training error-ই বেশি (ছাত্র A, অলস)
- **Over**fit → বড় train-test gap (ছাত্র C, মুখস্থবাজ)

### Rule 2: Capacity কোনদিকে বেশি (উল্টোগুলো মনে রাখো)
```
Polynomial:    বড় degree     = বেশি capacity
Neural net:    বড় network    = বেশি capacity
Decision tree: ছোট minsplit   = বেশি capacity   ← উল্টো!
k-NN:          ছোট k          = বেশি capacity   ← উল্টো!
```

### Rule 3: Error vs Capacity
- Training error → সবসময় কমে (monotone ↓)
- Test error → **U-shaped** (minimum মাঝখানে)

### Rule 4: VC Dimension সংখ্যা
- Homogeneous halfspace → **p**
- Non-homogeneous (bias সহ) → **p+1**
- Axis-aligned rectangle (ℝ²) → **4**
- Threshold classifier → **1**
- Sine / 1-NN → **∞**
> উল্টে বসালে FALSE।

### Rule 5: VC Dimension সংজ্ঞা
"VC dim d মানে সব d-set shatter হয়" → **সবসময় FALSE**। অন্তত একটা set shatter হলেই হলো।

### Rule 6: Parameter ≠ Capacity
Sine classifier-এর ১ parameter কিন্তু ∞ VC। "Parameter সংখ্যা দিয়ে capacity বিচার করা যায়" → **FALSE**।

### Rule 7: VC Bound প্রকৃতি
VC bound **loose, pessimistic**, **arbitrary distribution**-এর জন্য। Training error হলো **optimistic** estimate।

### Rule 8: তিনটা Error Term
```
GE = σ²       + Variance      + Bias²
   (noise,      (overfit,        (underfit,
   অপরিবর্তনীয়)  training বদলালে    ধারাবাহিক ভুল)
                prediction অস্থির)
```

### Rule 9: Bias-Variance Trade-off
- High capacity → **low bias, high variance**
- Low capacity → **high bias, low variance**
> উল্টো লিখলে FALSE।

### Rule 10: সমাধান
- Underfit (high bias) → model **flexible** করো
- Overfit (high variance) → **regularization** বা **আরো data**

### Rule 11: No Free Lunch
"সব problem-এর গড়ে কোনো algorithm সেরা না" → **TRUE**। "একটা universal best algorithm আছে" → **FALSE**।

### Rule 12: Ill-Posed
Assumption/prior ছাড়া ML অসম্ভব — কারণ অনেক function training data-র সাথে consistent থাকে।

---

# ১৮. পরীক্ষার আগে করণীয়

1. **এই full document একবার পড়ো** (১.৫–২ ঘণ্টা)
2. **Section ১৬-এর Master Quiz-Trap Table** ৩ বার পড়ো
3. **Section ৮-এর VC Dimension Table** মুখস্থ করো (সংখ্যাগুলো!)
4. **Section ১১-এর তিনটা error term** ও bias-variance trade-off গেঁথে নাও
5. **Section ১৭-এর ১২টা Golden Rules** মনে রাখো
6. **`true_false_quiz.md`** solve করো — ৩০ মিনিটে ৫১টা প্রশ্ন
7. **ভুল উত্তরগুলো** আবার সংশ্লিষ্ট section-এ পড়ো

## 🏆 শেষ কথা

Chapter 3 concept-ভিত্তিক, মুখস্থের জিনিস কম। ৫টা মূল idea বুঝলে যেকোনো প্রশ্নের উত্তর বের করতে পারবে:
1. **Capacity** → over/underfit নির্ধারণ করে; test error U-shaped
2. **VC dimension** → complexity মাপে; parameter সংখ্যা ≠ capacity
3. **Bias-Variance** → high capacity = low bias + high variance
4. **Learning curves** → আরো data variance কমায়, gap মিলিয়ে দেয়
5. **No Free Lunch** → assumption ছাড়া learning অসম্ভব

**মনে রাখো:** এই chapter বোঝার জিনিস। কেন overfit হয়, কেন bias-variance trade-off — এগুলো concept দিয়ে বুঝলে কখনো ভুলবে না।

পরীক্ষায় ভালো করার জন্য শুভ কামনা! 🎓

---

*— Opus-এর তরফ থেকে*
