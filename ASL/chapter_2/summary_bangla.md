# অধ্যায় ২: Classification
### — বাংলায় সহজ ভাষায় সম্পূর্ণ ব্যাখ্যা

---

> **এই chapter-এর সবচেয়ে কঠিন অংশ:** Scoring vs Probabilistic classifier-এর পার্থক্য, LDA vs QDA, এবং "কোন দিক থেকে কোনদিকে যাওয়া সম্ভব" — এই জিনিসগুলো মাথায় গেঁথে নাও।

---

## ১. Classification Task — শুরু থেকে বুঝি

### গল্প: হাসপাতালের ডাক্তার

ধরো একজন ডাক্তার রোগীর test result দেখে সিদ্ধান্ত নিতে চান:
- রোগীর ক্যান্সার আছে কি নেই? → **Binary Classification** (২টা class)
- রোগীর রোগটা কোনটা: ডেঙ্গু, ম্যালেরিয়া, টাইফয়েড, বা ফ্লু? → **Multiclass Classification** (৪টা class)

এটাই classification। Input features দেখে output class predict করা।

### Encoding — এই course-এ কীভাবে class লেখা হয়

**Binary (g = 2) এর জন্য দুটো encoding চলে:**
- Y = {0, 1} — (0 = negative class, 1 = positive class)
- Y = {−1, +1} — (−1 = negative, +1 = positive)

**Multiclass (g ≥ 3) এর জন্য:**
- Y = {1, 2, ..., g}

---

## ২. কেন Model Class Output করে না? — এটা না বুঝলে সব গুলিয়ে যাবে

এখানেই বেশিরভাগ ছাত্র confused হয়। একটু সময় নাও।

### গল্প: ক্রিকেট ম্যাচের বেটিং

তুমি বলো:
- Version A: "বাংলাদেশ জিতবে।" — শুধু একটা সিদ্ধান্ত।
- Version B: "বাংলাদেশ জেতার সম্ভাবনা ৭২%, ভারতের ২৮%।" — score সহ সিদ্ধান্ত।

Version B অনেক বেশি useful কারণ:
1. তুমি threshold পরিবর্তন করতে পারো (৫০% না, ৬০% confidence হলেই bet করব)
2. Version B থেকে Version A বানানো সহজ (৭২% > ৫০% → বাংলাদেশ জিতবে)
3. কিন্তু Version A থেকে Version B বানানো **impossible** — "বাংলাদেশ জিতবে" থেকে ৭২% কে বের করবে?

এজন্য models f : X → R^g দিয়ে **score** output করে — discrete class না।

**তিনটা কারণ (slide থেকে):**
1. Continuous function-এর উপর optimization করা অনেক সহজ
2. Score/probability-তে বেশি information থাকে
3. Score → Class: সম্ভব; Class → Score: **impossible**

---

## ৩. দুই ধরনের Classifier

### ৩.১ Scoring Classifier

**কী করে:** g টা scoring function বানায়, সবচেয়ে বেশি score যে class-এর সেটা predict করে।

```
Input x → [f_1(x), f_2(x), ..., f_g(x)] → argmax → class k
         (scores for each class)
```

**Formula:**
$$h(\mathbf{x}) = \underset{k \in \{1,2,\ldots,g\}}{\arg\max} \; f_k(\mathbf{x})$$

**Binary case (g=2) গল্প:**

ধরো তোমার কাছে দুটো score আছে: f_1(x) = spam score, f_{-1}(x) = not-spam score।

তুমি f_1 − f_{-1} নাও → f(x)।
- f(x) > 0 মানে f_1 > f_{-1} মানে spam বেশি likely → predict spam (class +1)
- f(x) < 0 মানে f_1 < f_{-1} মানে not-spam বেশি likely → predict not-spam (class -1)

তাই: **h(x) = sgn(f(x))** (sign function)

**Confidence কী?**

|f(x)| = score-এর magnitude = কতটা confident।

```
f(x) = +5 → sgn = +1 (spam), confidence = 5 (very confident)
f(x) = +0.1 → sgn = +1 (spam), confidence = 0.1 (barely sure)
f(x) = -3 → sgn = -1 (not spam), confidence = 3
```

### ৩.২ Probabilistic Classifier

**কী করে:** প্রতিটা class-এর probability output করে।

$$\pi_1(\mathbf{x}), \pi_2(\mathbf{x}), \ldots, \pi_g(\mathbf{x}) \quad \text{যেখানে} \sum_l \pi_l(\mathbf{x}) = 1$$

**Formula:**
$$h(\mathbf{x}) = \underset{k}{\arg\max} \; \pi_k(\mathbf{x})$$

**Binary case:**
একটাই probability function π(x) দেওয়া হয় (P(class=1 | x))।

Threshold c দিয়ে class predict করো:
$$h(\mathbf{x}) = \mathbb{1}(\pi(\mathbf{x}) \geq c)$$

**Default threshold: c = 0.5** (সমান সম্ভাবনার চেয়ে বেশি হলেই class 1)

**Quiz traps:**
- "Default threshold for binary **probabilistic** classifier" → **c = 0.5**
- "Default threshold for binary **scoring** classifier" → **c = 0** (h = sgn(f), f > 0 মানেই class 1)
- "Probabilistic classifiers CANNOT be seen as scoring classifiers" → **FALSE** (পারে!)

---

## ৪. Probabilities, Scores, Classes — এদের সম্পর্ক

এটা এই chapter-এর সবচেয়ে tricky অংশ। ভালো করে পড়ো।

```
          Calibrating/Scaling
Probabilities ◄─────────────── Scores
      │                           │
      │ Thresholding    Thresholding
      ▼                           ▼
         Discrete Classes
    (often intrinsically produced by scores)
```

### কোন দিক থেকে কোনদিকে যাওয়া যায়?

| From → To | সম্ভব? | Method |
|-----------|--------|--------|
| Scores → Probabilities | ✅ YES | Calibrating/Scaling (sigmoid দিয়ে) |
| Scores → Discrete Classes | ✅ YES | Thresholding |
| Probabilities → Discrete Classes | ✅ YES | Thresholding |
| Probabilities → Scores | ✅ YES | Inverse (Calibrating/Scaling এর উল্টো) |
| **Discrete Classes → Scores** | ❌ **NO** | Impossible! |
| **Discrete Classes → Probabilities** | ❌ **NO** | Impossible! |

**গল্প দিয়ে বোঝাই:**

তুমি একটা পরীক্ষায় "Pass/Fail" result পেয়েছ। এখন কেউ যদি জিজ্ঞেস করে "তুমি কত নম্বর পেয়েছিলে?" — তুমি বলতে পারবে না! Pass/Fail থেকে exact number বের হয় না।

কিন্তু যদি তুমি নম্বর জানো (say, 65/100), তাহলে:
- 65/100 → "Pass" → সহজ (thresholding: ≥50 → Pass)
- 65/100 → probability 65% → সম্ভব (scaling)

**Quiz trap:** "Discrete classes can always be converted back to scores" → **FALSE**

---

## ৫. Decision Boundary — এটা বোঝা দরকার

### গল্প: মানচিত্রে দেশের সীমানা

বাংলাদেশ-ভারত সীমানার কথা ভাবো। সীমানার এক পাশে "বাংলাদেশ" (class 1), অন্য পাশে "ভারত" (class 2)।

Decision boundary হলো ঠিক সেই সীমানা — যেখানে model বলতে পারে না কোন class।

### Formal Definition

Input space X কে g টা **decision region** ভাগ করা হয়:

$$\mathcal{X}_k = \{\mathbf{x} \in \mathcal{X} : h(\mathbf{x}) = k\}$$

Decision region X_k = সব সেই points যেখানে model class k predict করে।

Decision boundary = এই regions-এর **সীমানা** — যেখানে দুটো class-এর score সমান।

**Binary case:**
- Scoring classifier: f(x) = 0 → decision boundary
- Probabilistic classifier: π(x) = 0.5 → decision boundary

**Multiclass case:**
$$\{x : f_i(x) = f_j(x) \text{ এবং } f_i(x) \geq f_k(x) \text{ সব } k \neq i,j \text{ এর জন্য}\}$$

---

## ৬. Linear Classifier — সরল রেখার সীমানা

### সংজ্ঞা

যদি discriminant function-কে একটা monotone (একদিকে বাড়তে থাকা) transformation g এর মাধ্যমে linear করা যায়:

$$g(f_k(\mathbf{x})) = \mathbf{w}_k^\top \mathbf{x} + b_k$$

তাহলে এটা **linear classifier।**

### কেন এই classifier "linear"?

দুটো class i এবং j এর মধ্যে tie (decision boundary):

$$f_i(\mathbf{x}) = f_j(\mathbf{x})$$
$$\Downarrow$$
$$\mathbf{w}_i^\top \mathbf{x} + b_i = \mathbf{w}_j^\top \mathbf{x} + b_j$$
$$\Downarrow$$
$$(\mathbf{w}_i - \mathbf{w}_j)^\top \mathbf{x} + (b_i - b_j) = 0$$

এটা একটা **hyperplane** — সরল রেখা (2D-তে), সমতল (3D-তে)।

### গুরুত্বপূর্ণ সত্য (অনেকে এটা জানে না)

**Linear classifier মানেই original input space-এ linear boundary না!**

যদি তুমি features engineer করো (polynomial features, basis functions), তাহলে:
- Feature space-এ: linear boundary
- **Original input space-এ: non-linear boundary**

**উদাহরণ:**
Input: x (১D)। তুমি feature বানালে: [x, x², x³]।
এই নতুন space-এ linear classifier → original space-এ curve (polynomial boundary)।

**Quiz trap:** "Linear classifiers can ONLY produce linear decision boundaries" → **FALSE**

---

## ৭. Sigmoid Functions — Score কে Probability-তে বদলানো

### সমস্যা

Scoring classifier-এর output f(x) যেকোনো real number হতে পারে: −∞ থেকে +∞।
কিন্তু probability [0, 1] range-এ থাকতে হবে।

**কীভাবে convert করব?** → **Sigmoid function** ব্যবহার করো।

### Sigmoid-এর সংজ্ঞা

একটা sigmoid function s : R → [0, 1] এর বৈশিষ্ট্য:
1. **Bounded:** [0, 1] range-এ থাকে
2. **Differentiable:** derivative আছে সর্বত্র
3. **Non-decreasing:** derivative ≥ 0 সব জায়গায়

### চারটা sigmoid উদাহরণ

| Sigmoid | Formula |
|---------|---------|
| Arctan | s(t) = arctan(t) |
| Tanh | s(t) = (e^t − e^{−t}) / (e^t + e^{−t}) |
| **Logistic** | s(t) = 1 / (1 + e^{−t}) |
| Probit | Normal CDF |

### Logistic Function — সবচেয়ে গুরুত্বপূর্ণ

$$s(t) = \frac{1}{1 + e^{-t}}$$

**Graph দেখতে কেমন:**
```
s(t)
1.0 |─────────────────────────────────────────
    |                                ╭─────────
0.5 |─────────────────────── ╭──────╯
    |                   ╭────╯
0.0 |────────────────────────────────────────→ t
                       0
```

**৩টা গুরুত্বপূর্ণ property (এগুলো quiz-এ আসে):**

**Property 1:** Limits
$$\lim_{t \to -\infty} s(t) = 0 \qquad \lim_{t \to +\infty} s(t) = 1$$

**Property 2:** Derivative
$$\frac{\partial s(t)}{\partial t} = s(t)(1 - s(t))$$

এটা অনেক সুন্দর — derivative নিজেকে দিয়েই লেখা যায়!

**Property 3:** Symmetry
$$s(t) \text{ is symmetrical about the point } (0, \tfrac{1}{2})$$

মানে: s(0) = 1/2, এবং s(t) + s(−t) = 1।

**কোথায় ব্যবহার:**
- Logistic Regression-এ probability output করতে
- Deep learning-এ activation function হিসেবে

**Quiz trap:** "Logistic function: s(t) = 1/(1+e^t)" → **FALSE** (e^{**−**t} হবে, e^t না!)

---

## ৮. Softmax — Multiclass-এর জন্য Logistic-এর বড় ভাই

### সমস্যা

Binary-তে একটা logistic function কাজ করে। কিন্তু ৩, ৪, ৫ class হলে?

→ **Softmax** ব্যবহার করো।

### Formula

$$\pi_k(\mathbf{x}) = \frac{\exp(f_k(\mathbf{x}))}{\sum_{j=1}^g \exp(f_j(\mathbf{x}))}$$

**গল্প দিয়ে বোঝাই:**

তোমার ৩টা class: {কুকুর, বিড়াল, পাখি}।
Scores: f_1 = 3 (কুকুর), f_2 = 1 (বিড়াল), f_3 = 0 (পাখি)।

```
exp(3) = 20.09
exp(1) = 2.72
exp(0) = 1.00
Sum    = 23.81

π_1 = 20.09 / 23.81 = 0.844 (কুকুর: 84.4%)
π_2 = 2.72  / 23.81 = 0.114 (বিড়াল: 11.4%)
π_3 = 1.00  / 23.81 = 0.042 (পাখি: 4.2%)
```

সব মিলিয়ে = 1.000 ✓

### Softmax vs Argmax — এখানে অনেকে ভুল করে

| | Argmax | Softmax |
|--|--------|---------|
| **Output** | শুধু সবচেয়ে বড়টার index | সবগুলোর probability |
| **Non-maximal info** | **হারিয়ে যায়** | **সংরক্ষিত থাকে** |
| **Reversible?** | ❌ No | ✅ Yes |

**"Soft" max কেন?**
কারণ এটা argmax-এর মতো কাজ করে কিন্তু "নরমভাবে" — non-maximal elements-এর তথ্যও রাখে।

**Quiz traps:**
- "Softmax generalizes logistic to multiclass" → **TRUE** ✓ (g=2 দিলে logistic পাওয়া যায়)
- "It is the argmax that keeps non-maximal information in a reversible way" → **FALSE** (softmax করে, argmax না!)

---

## ৯. Generative vs Discriminative — এটাই সবচেয়ে বড় conceptual trap

এটা বোঝার জন্য একটু ধৈর্য্য ধরো।

### দুটো দল

ধরো তুমি আমকে আর জামকে আলাদা করতে চাও।

**Generative approach (দল ১):** আমের বৈশিষ্ট্য মনে রাখো (আম কেমন দেখতে), জামের বৈশিষ্ট্য মনে রাখো (জাম কেমন দেখতে)। নতুন ফল দেখলে — এটা আমের মতো বেশি নাকি জামের মতো বেশি? Bayes' theorem দিয়ে সিদ্ধান্ত নাও।

**Discriminative approach (দল ২):** সরাসরি শিখে নাও — এই features দেখলে আম, ওই features দেখলে জাম। আমের/জামের পূর্ণ বৈশিষ্ট্য মনে রাখার দরকার নেই।

### Generative Approach — Bayes' Theorem

$$\pi_k(\mathbf{x}) = \mathbb{P}(y=k \mid \mathbf{x}) = \frac{\mathbb{P}(\mathbf{x} \mid y=k) \cdot \mathbb{P}(y=k)}{\mathbb{P}(\mathbf{x})}$$

ভেঙে বলি:
- P(x | y=k) = "class k-এর data কেমন দেখতে?" — এটা model করা হয়
- P(y=k) = class k কতটা common? (prior)
- P(x) = normalizing constant

Discriminant functions (যা দিয়ে class বেছে নিই):
- π_k(x) নিজেই, বা
- log P(x|y=k) + log π_k

### চারটা Classifier — LDA, QDA, Naive Bayes, Logistic Regression

এই চারটার পার্থক্য quiz-এ অবশ্যই আসবে।

#### LDA (Linear Discriminant Analysis)

**Assumption:** প্রতিটা class-এর data multivariate Gaussian distribution follow করে, কিন্তু **সব class-এর covariance matrix একই (Σ)।**

$$\mathbb{P}(\mathbf{x} \mid y=k) \sim \mathcal{N}(\boldsymbol{\mu}_k, \boldsymbol{\Sigma})$$

**গল্প:** ধরো ক্লাস ছেলে এবং মেয়েদের উচ্চতা-ওজন distribution। LDA বলে: ছেলে ও মেয়েদের distribution-এর shape একই (same Σ), শুধু center (μ_k) আলাদা।

**Result:** Decision boundary = **hyperplane (linear)**

কারণ: Σ সব class-এ same → quadratic terms cancel হয়ে যায় → linear থাকে।

#### QDA (Quadratic Discriminant Analysis)

**Assumption:** প্রতিটা class-এর **আলাদা আলাদা covariance matrix (Σ_k)**।

$$\mathbb{P}(\mathbf{x} \mid y=k) \sim \mathcal{N}(\boldsymbol{\mu}_k, \boldsymbol{\Sigma}_k)$$

**Result:** Decision boundary = **quadratic (non-linear)**

কারণ: Σ_k আলাদা → quadratic terms cancel হয় না।

#### LDA vs QDA একটা table-এ:

| | LDA | QDA |
|--|-----|-----|
| Covariance | Equal (Σ) for all classes | Different (Σ_k) per class |
| Decision boundary | **Linear (hyperplane)** | **Quadratic (curve)** |
| Type | Generative, Linear | Generative, Non-linear |

**Quiz trap (সবচেয়ে বেশি আসে):**
- "LDA assumes **unequal** covariances" → **FALSE** (equal!)
- "QDA assumes **equal** covariances" → **FALSE** (unequal!)
- "LDA produces a **non-linear** decision boundary" → **FALSE** (linear!)

#### Naive Bayes

**Assumption:** Features গুলো class দেওয়া সাপেক্ষে (conditionally) independent।

$$\mathbb{P}(\mathbf{x} \mid y=k) = \prod_{j=1}^p \mathbb{P}(x_j \mid y=k)$$

**গল্প:** ধরো spam email classify করতে চাও। Naive Bayes বলে: "email-এ 'free' শব্দ আছে কিনা, আর 'money' শব্দ আছে কিনা" — এই দুটো feature spam হওয়ার সাপেক্ষে independent।

বাস্তবে এটা সত্যি নাও হতে পারে ('free' এবং 'money' একসাথে আসতে পারে spam-এ), তাই এটা "Naive" — কারণ এই assumption-টা naive/সরলীকৃত।

**Quiz trap:**
- "Naive Bayes assumes features are unconditionally independent" → **FALSE**
- **Conditionally independent** — given class y → সঠিক।

#### Logistic Regression

- **Discriminant approach** (generative না!)
- Discriminant function সরাসরি learn করে
- Linear decision boundary
- Loss minimization ব্যবহার করে

**Quiz trap:** "Logistic Regression is a generative approach" → **FALSE** (discriminant!)

---

## ১০. পুরো Chapter ২ একটা বড় গল্পে

ধরো তুমি একটা চাকরির interview নিচ্ছ। ৩ ধরনের candidate: Junior, Mid-level, Senior।

**তুমি scoring classifier হলে:**
প্রতিটা candidate-কে ৩টা score দেবে: Junior score, Mid-level score, Senior score। সবচেয়ে বেশি score-এর category predict করবে।

**তুমি probabilistic classifier হলে:**
প্রতিটা candidate-এর জন্য probability বলবে: Junior 20%, Mid-level 65%, Senior 15%। Softmax দিয়ে scores → probabilities।

**Decision boundary:**
একটা imaginary line যেখানে তুমি বলতে পারো না কোন category। এর দুই পাশে আলাদা সিদ্ধান্ত।

**LDA হলে:**
Junior, Mid-level, Senior — তিনটা class-এর CV data same spread কিন্তু different center → linear boundary।

**QDA হলে:**
Junior CV-গুলো closely clustered, Senior CV-গুলো widely spread → different covariance → quadratic (curved) boundary।

**Naive Bayes হলে:**
"Programming skill" এবং "Communication skill" — class দেওয়া সাপেক্ষে independent assume করো।

**Logistic Regression হলে:**
সরাসরি শেখো: এই combination of features মানেই Senior। Generative distribution model করো না।

---

## ১১. Quick Revision Table — Quiz-এর আগে দেখো

| Statement | T/F | কারণ |
|-----------|-----|------|
| Models output discrete classes directly | FALSE | Score output করে |
| Discrete classes → probabilities: possible | FALSE | Irreversible! |
| Scores → probabilities: possible | TRUE | Sigmoid/calibration দিয়ে |
| Default threshold for binary probabilistic classifier: c = 0.5 | TRUE | — |
| Default threshold for binary scoring classifier: c = 0 | TRUE | h = sgn(f) |
| Probabilistic classifiers cannot be seen as scoring classifiers | FALSE | পারে |
| Linear classifiers can ONLY produce linear decision boundaries | FALSE | Feature engineering দিয়ে non-linear হয় |
| Logistic function: s(t) = 1/(1+e^t) | FALSE | e^{**−**t} হবে |
| Derivative of logistic: ∂s/∂t = s(t)(1−s(t)) | TRUE | — |
| Logistic is symmetric about (0, 1/2) | TRUE | — |
| Softmax reduces to logistic for g=2 | TRUE | — |
| Argmax keeps non-maximal info reversibly | FALSE | Softmax করে |
| LDA: equal covariances, linear boundary | TRUE | — |
| QDA: equal covariances, quadratic boundary | FALSE | **Un**equal covariances |
| LDA: unequal covariances | FALSE | LDA = **equal** |
| Naive Bayes: unconditional independence of features | FALSE | **Conditional** independence given y |
| Logistic Regression: generative approach | FALSE | **Discriminant** approach |
| Deep learning uses sigmoid as activation functions | TRUE | — |

---

## ১২. Bonus: যে জিনিসটা প্রায় সবাই ভুল করে — আরেকবার

### LDA vs QDA — চোখ বন্ধ করে মনে রাখার trick

**L**DA → **L**inear → **L**evel covariance (সব class-এ same Σ)
**Q**DA → **Q**uadratic → ভিন্ন covariance (Σ_k per class)

### Scoring vs Probabilistic Threshold

**মনে রাখো:**
- Scoring: f(x) compare করা হয় **0** এর সাথে (positive/negative side)
- Probabilistic: π(x) compare করা হয় **0.5** এর সাথে (majority probability)

### Softmax "Soft" কেন?

```
Input scores: [3, 1, 0]

Argmax:  [1, 0, 0]  ← শুধু winner রাখে, বাকি সব হারিয়ে যায়
Softmax: [0.84, 0.11, 0.04]  ← সবার information রাখে, reversible
```

Softmax "soft" কারণ এটা gentle — সবাইকে কিছু না কিছু দেয়। Argmax "hard" — winner-takes-all।
