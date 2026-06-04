# অধ্যায় ৩: Hypothesis Spaces and Capacity
### — বাংলায় সহজ ভাষায় সম্পূর্ণ ব্যাখ্যা

---

> **এই chapter-এর সবচেয়ে কঠিন অংশ:** VC dimension (shattering), Bias-Variance decomposition, এবং কোন capacity-তে কী হয় (bias↓ variance↑) — এই জিনিসগুলো মাথায় গেঁথে নাও। আর No Free Lunch theorem-টা concept হিসেবে বুঝে রাখো।

---

## ১. Capacity, Underfitting, Overfitting — শুরু থেকে বুঝি

### গল্প: পরীক্ষার প্রস্তুতি

ধরো তিনজন ছাত্র পরীক্ষা দিচ্ছে:
- **ছাত্র A (Underfitting):** কিছুই ভালো করে পড়েনি। বাসায় practice problem-ও পারে না, পরীক্ষাতেও পারে না। → **Training error বেশি।**
- **ছাত্র B (Just right):** বুঝে বুঝে পড়েছে। practice-ও পারে, নতুন প্রশ্নও পারে। → **ভালো generalization।**
- **ছাত্র C (Overfitting):** সব practice problem মুখস্থ করে ফেলেছে। বাসায় ১০০% পারে, কিন্তু পরীক্ষায় নতুন প্রশ্ন দিলে গুবলেট। → **Training error কম, কিন্তু test error বেশি।**

### সংজ্ঞা

- **Underfitting:** training error যথেষ্ট কম করতে না পারা।
- **Overfitting:** training error আর test error-এর মধ্যে বড় পার্থক্য।
- **Generalization gap:** training error আর generalization (test) error-এর মধ্যেকার ফাঁক।

### Capacity কী?

Model over/underfit করবে কিনা সেটা নির্ভর করে তার **capacity**-এর উপর — মানে সে কী ধরনের hypothesis শিখতে পারে।

- **Low capacity:** কয়েকটা সহজ hypothesis শিখতে পারে → underfit করে।
- **High capacity:** অনেক, জটিল hypothesis শিখতে পারে → overfit করে।
- **Optimal capacity:** যখন model না underfit করে না overfit করে — তখন test error সবচেয়ে কম।

```
Error
  |\                        Generalization error (test)
  | \                      /
  |  \   Underfit | Overfit
  |   \__________ | _____/   ← generalization gap
  |              \|____/ ___ Training error
  |_______________|________________ Capacity
            Optimal Capacity
```

---

## ২. Hypothesis Space — Model-এর "অনুমোদিত আকৃতি"

একটা learner-এর **representation** = তার hypothesis space = যেসব model সে বেছে নিতে পারে তাদের set।

$$\mathcal{H} := \{ f : \mathcal{X} \to \mathbb{R}^g \mid f \text{ এর একটা নির্দিষ্ট form আছে} \}$$

প্রায়ই $f$ একটা parameter $\boldsymbol\theta$ দিয়ে লেখা হয়: $f(\mathbf{x}) = f(\mathbf{x} \mid \boldsymbol\theta)$।

> **Note:** যখন hard classifier (discrete class output) বোঝাই তখন $f$-এর বদলে $h$ লিখি। সাধারণভাবে $f$ দিয়েই class/score/probability সব বোঝানো হয়।

### উদাহরণ

| Model | Form |
|-------|------|
| **Linear regression** | $f(\mathbf{x}) = \mathbf{x}^\top\boldsymbol\theta + \theta_0$ |
| **Separating hyperplane** | $h(\mathbf{x}) = \mathbb{1}(\mathbf{x}^\top\boldsymbol\theta - \theta_0 > 0)$ |
| **Decision tree** | $f(\mathbf{x}) = \sum_i c_i \mathbb{1}(\mathbf{x} \in Q_i)$ — feature space-কে axis-aligned rectangle-এ ভাগ করে |
| **Ensemble** | $f(\mathbf{x}) = \sum_l \beta^{[l]} b^{[l]}(\mathbf{x})$ — অনেক model-এর prediction একসাথে (random forest, bagging, boosting) |
| **Neural network** | অনেকগুলো layer-এর nested composition |

### Neural network-এর একটা neuron

প্রতিটা neuron দুই ধাপে কাজ করে:
1. **Weighted sum:** $\phi^{(j)}(\mathbf{z}) = \mathbf{w}_j^\top \mathbf{z} + b_j$
2. **Non-linear transform:** $\sigma^{(j)}(\mathbf{w}_j^\top \mathbf{z} + b_j)$

পুরো network = এরকম operation-এর nested composition।

---

## ৩. Overfitting — কীভাবে হয়

- Model-এর capacity বাড়ানো হয় **hypothesis space বড় করে।**
- এতে সাধারণত learnable parameter-এর সংখ্যাও বাড়ে।
- উদাহরণ: polynomial-এর degree বাড়ানো, tree-এর depth বাড়ানো, network বড় করা, predictor যোগ করা।
- $\mathcal{H}$ যত বড় হয়, overfit-এর প্রবণতা তত বাড়ে — model এমনকি training data-র **random ভুলগুলোও** শিখে ফেলে।

### তিনটা উদাহরণ (table মুখস্থ না, pattern বুঝো)

**Polynomial regression:**

| | Degree 1 | Degree 5 | Degree 13 |
|--|----------|----------|-----------|
| | Underfit | ✅ Just right | Overfit |
| Train error | 3.87 | 1.23 | **0.48** (সবচেয়ে কম!) |
| Test error | 4.11 | **1.55** (সবচেয়ে কম!) | 148.5 (বিশাল!) |

**Decision tree (minsplit ছোট = capacity বেশি):**

| | minsplit 60 | minsplit 12 | minsplit 2 |
|--|-------------|-------------|------------|
| | Underfit | ✅ Just right | Overfit |
| Train error | 0.36 | 0.12 | **0.02** |
| Test error | 0.40 | **0.32** | 0.35 |

**k-NN (k ছোট = capacity বেশি):**

| | k = 20 | k = 7 | k = 1 |
|--|--------|-------|-------|
| | Underfit | ✅ Just right | Overfit |
| Train error | 0.22 | 0.13 | **0** |
| Test error | 0.40 | **0.25** | 0.33 |

> **🔑 মূল Pattern:** capacity বাড়লে **training error সবসময় কমতে থাকে** (highest capacity-তে সবচেয়ে কম), কিন্তু **test error U-shaped** — মাঝামাঝি capacity-তে সবচেয়ে কম।

**Quiz traps:**
- "capacity বাড়লে test error monotonically কমে" → **FALSE** (U-shaped!)
- "k-NN-এ বড় k মানে বেশি capacity" → **FALSE** (ছোট k = বেশি capacity)
- "decision tree-তে বড় minsplit = বেশি capacity" → **FALSE** (ছোট minsplit = বেশি)

---

## ৪. VC Dimension — Hypothesis Space কতটা জটিল

Function space-এর complexity মাপার একটা general পদ্ধতি = **Vapnik–Chervonenkis (VC) dimension।**

### Shatter মানে কী?

একটা point-এর set কে যদি কোনো class-এর function **যেকোনো** label assignment-এর জন্য perfectly আলাদা করতে পারে, তাহলে বলি set-টা **shattered**।

### VC dimension-এর সংজ্ঞা

$\mathcal{H}$-এর VC dimension = **সবচেয়ে বড় সংখ্যক point** (কোনো এক configuration-এ) যাদের shatter করা যায়।

> **খুব গুরুত্বপূর্ণ:** VC dimension = d মানে **সব** d-size set shatter হয় **না**। মানে — **অন্তত একটা** d-size set shatter হয়, এবং **কোনো** (d+1)-size set shatter হয় না।

### মুখস্থ রাখার table

| Hypothesis class | VC dimension |
|------------------|--------------|
| ℝ²-তে linear indicator function | **3** (৩ point shatter করে, ৪ পারে না) |
| Homogeneous halfspace $\text{sign}(\mathbf{x}^\top\boldsymbol\theta)$ | exactly **p** |
| Non-homogeneous halfspace $\text{sign}(\mathbf{x}^\top\boldsymbol\theta + \theta_0)$ | exactly **p+1** |
| ℝ²-তে axis-aligned rectangle | **4** |
| Threshold classifier $\mathbb{1}(x \ge \theta)$ | **1** |
| 1-Nearest Neighbour (k=1) | **infinite** |
| Sine classifier $\mathbb{1}(\sin(\theta x) > 0)$ | **infinite** |

> **🔑 সবচেয়ে বড় trap:** VC dimension সাধারণত parameter সংখ্যার সাথে বাড়ে, কিন্তু **শুধু parameter সংখ্যা দিয়ে capacity বিচার করা যায় না!** Sine classifier-এর মাত্র **একটা** parameter, তবু VC dimension **infinite**।

### VC Bound আর PAC

Training error হলো generalization error-এর একটা **optimistic** (আশাবাদী) estimate। VC dimension d হলে, $1-\delta$ probability-তে:

$$\text{err}_{\text{test}} \le \text{err}_{\text{train}} + \epsilon$$

যেখানে $\epsilon$ depends on d, sample size, আর $\delta$।

- এই bound **arbitrary** (যেকোনো) $\mathbb{P}_{xy}$-এর জন্য ধরে।
- d **finite** হলে, sample size বাড়িয়ে $\delta$ আর $\epsilon$ দুটোই যত খুশি ছোট করা যায়।
- **Corollary:** training error কম হলে test error-ও (probably) কম। এমন algorithm-কে বলে **Probably Approximately Correct (PAC)**।

> **বাস্তবতা:** VC bound গুলো **খুবই loose আর pessimistic** (কারণ যেকোনো distribution-এর জন্য ধরতে হয়)। Neural network-এর মতো জটিল model এই bound-এর চেয়ে অনেক ভালো করে। তাই **Rademacher complexity**-এর মতো tighter বিকল্পও আছে। আসল capacity optimizer-এর উপরও নির্ভর করে। সবচেয়ে ভালো estimate = test set-এ evaluate করা।

---

## ৫. Bias–Variance Decomposition — এই chapter-এর হৃদয়

### Assumption

ধরে নিই data তৈরি হয়েছে: $y = f_{\text{true}}(\mathbf{x}) + \epsilon$, যেখানে noise $\epsilon \sim \mathcal{N}(0, \sigma^2)$ এবং $\mathbf{x}$-এর থেকে independent।

### L2-loss দিয়ে ভাঙলে — তিনটা term

$$GE = \underbrace{\sigma^2}_{\text{(১) data variance}} + \underbrace{\text{Var}(\hat f)}_{\text{(২) inducer variance}} + \underbrace{\text{Bias}^2}_{\text{(৩) squared bias}}$$

**(১) Variance of the data ($\sigma^2$):**
- Data-র মধ্যেকার **noise**।
- একে বলে **intrinsic / unavoidable / irreducible error**।
- **যত ভালো learner-ই হোক, এর নিচে কখনো যাওয়া যাবে না।**

**(২) Variance of the inducer:**
- Training data বদলালে prediction কতটা বদলায়।
- Learner কতটা **random জিনিস শেখে** real signal বাদ দিয়ে → **overfitting**।

**(৩) Squared bias:**
- Learner কতটা **ধারাবাহিকভাবে ভুল** করে → **underfitting**।

### Capacity-এর সাথে সম্পর্ক

| Capacity | Bias | Variance |
|----------|------|----------|
| **High** | **Low** | **High** |
| **Low** | **High** | **Low** |

> **গল্প:** High bias = এত শক্ত model যে curve-টাই ধরতে পারে না। High variance = flexible model যা সত্যিটা শিখতে পারে, কিন্তু প্রতিবার আলাদা training set দিলে **একদম আলাদা আলাদা** উত্তর দেয়।

**Quiz trap:** "High capacity model-এর high bias আর low variance" → **FALSE** (উল্টো: low bias, high variance)

---

## ৬. Learning Curves

### Error vs Capacity
Capacity বাড়লে:
- Variance-এর error **বাড়ে**।
- Bias-এর error **কমে**।
- Overfitting region-এ generalization error **বাড়ে** কারণ variance বৃদ্ধি >> bias হ্রাস।

### Error vs Training-set size
Training set বড় হলে:
- Variance-এর error **মিলিয়ে যায়**।
- Generalization error আর training error দুটোই **bias-এ গিয়ে মেলে** (noise শূন্য ধরলে)।
- **Generalization gap মিলিয়ে যায়।**

### সমাধান

| সমস্যা | কারণ | সমাধান |
|--------|------|--------|
| High bias (underfit) | Model খুব শক্ত | Model **বেশি flexible** করো (bias কমাও) |
| High variance (overfit) | Model খুব flexible | Model **কম flexible (regularization)** বা **আরো data** যোগ করো |

**Quiz trap:** "Underfitting কমাতে আরো regularization যোগ করো" → **FALSE** (regularization variance কমায়, bias কমায় না)

---

## ৭. ML একটা Ill-Posed Problem

- Learning algorithm-কে **আগে দেখেনি এমন (test)** data-তে ভালো করতে হবে।
- ধরো ৪টা boolean feature-এর উপর একটা boolean function শিখছ। সম্ভাব্য function = $2^{16} = 65536$টা।
- ৭টা example দেখার পরেও **$2^9$টা function** এখনো consistent থাকে।
- Unseen point-গুলোর **যেকোনো** label হতে পারে → **ML একটা ill-posed problem**।
- কোনো assumption ছাড়া, ML কি **random guess**-এর চেয়ে ভালো করতে পারে? — সাধারণভাবে না।

---

## ৮. No Free Lunch (NFL) Theorem

- একটা **নির্দিষ্ট** problem-এ কিছু algorithm অন্যদের চেয়ে ভালো করতে পারে।
- কিন্তু **সব সম্ভাব্য problem-এর গড়ে**, **কোনো** algorithm অন্য কারো চেয়ে ভালো না — random guessing সহ। এটাই **No Free Lunch theorem**।
- যে algorithm একদল problem-এ ভালো করে, তাকে অন্য একদল problem-এ **খারাপ** করতেই হবে। "Free lunch" বলে কিছু নেই।

### মূল কথা
- কোনো algorithm **সর্বজনীনভাবে** সেরা না।
- Learning algorithm-কে একটা **নির্দিষ্ট prior**-এর সাথে মানিয়ে নিতে হয় — assumption ছাড়া learning অসম্ভব।
- **খুব নির্দিষ্ট** assumption → ছোট একদল problem-এ দারুণ।
- **খুব broad** assumption → "jack of all trades, master of none"।

**Quiz trap:** "এমন একটা algorithm আছে যা সব problem-এ সবার চেয়ে ভালো" → **FALSE** (NFL অনুযায়ী নেই)

---

## ৯. Quick Revision Table — Quiz-এর আগে দেখো

| Statement | T/F | কারণ |
|-----------|-----|------|
| Underfitting = high training error | TRUE | — |
| Overfitting = বড় train-test gap | TRUE | — |
| Low capacity model জটিল hypothesis শেখে | FALSE | শুধু সহজগুলো |
| Test error highest capacity-তে minimum | FALSE | Optimal (মাঝামাঝি) capacity-তে |
| Capacity বাড়লে test error monotonically কমে | FALSE | U-shaped |
| k-NN-এ ছোট k = বেশি capacity | TRUE | — |
| VC dim d মানে সব d-point shatter হয় | FALSE | অন্তত একটা set |
| Homogeneous halfspace: VC = p | TRUE | — |
| Non-homogeneous halfspace: VC = p | FALSE | **p+1** |
| Axis-aligned rectangle (ℝ²): VC = 4 | TRUE | — |
| Threshold classifier: VC = 1 | TRUE | — |
| Sine classifier: VC = 1 (এক parameter) | FALSE | **infinite!** |
| Parameter সংখ্যা দিয়ে সবসময় capacity বিচার করা যায় | FALSE | Sine example |
| VC bound tight ও optimistic | FALSE | loose ও pessimistic |
| σ² (data variance) দূর করা যায় | FALSE | irreducible error |
| High capacity → high bias, low variance | FALSE | low bias, high variance |
| Capacity↑ → variance↑, bias↓ | TRUE | — |
| Training set↑ → variance error মিলিয়ে যায় | TRUE | — |
| Underfit কমাতে regularization বাড়াও | FALSE | model flexible করো |
| ML assumption ছাড়াও কাজ করে | FALSE | ill-posed |
| NFL: সব problem-এর গড়ে কোনো algorithm সেরা না | TRUE | — |

---

## ১০. Bonus: মনে রাখার Trick

### Bias vs Variance চোখ বন্ধ করে
- **Bias** → **B**ekar শক্ত model → curve ধরতে পারে না → **underfit**
- **Variance** → **V**ery flexible → প্রতিবার আলাদা উত্তর → **overfit**
- Capacity বাড়াও → Bias কমে, Variance বাড়ে (একটা বাড়লে অন্যটা কমে — trade-off)

### Capacity বেশি কোনদিকে?
```
Polynomial:  বড় degree     = বেশি capacity
Decision tree: ছোট minsplit = বেশি capacity   ← উল্টো!
k-NN:         ছোট k         = বেশি capacity   ← উল্টো!
```

### তিনটা error term
```
GE = σ²        +  Variance     +  Bias²
   (noise,        (overfit,        (underfit,
   অপরিবর্তনীয়)   training বদলালে   ধারাবাহিক ভুল)
                  prediction বদলায়)
```
