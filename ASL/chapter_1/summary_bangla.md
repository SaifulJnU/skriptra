# অধ্যায় ১: Introduction and Formalization
### — বাংলায় সহজ ভাষায় সম্পূর্ণ ব্যাখ্যা

---

> **পড়ার আগে একটা কথা:** এই নোটটা তোমার জন্য লেখা হয়েছে যেন তুমি মাথা দিয়ে বুঝতে পারো, মুখস্থ করে পরীক্ষায় পাশ করার জন্য না। প্রতিটা concept-এর পেছনে একটা গল্প আছে — সেই গল্পটা বোঝো, তাহলে True/False প্রশ্ন দেখলেই মাথায় উত্তর চলে আসবে।

---

## ১. Machine Learning আসলে কী? — একটা গল্প দিয়ে শুরু করি

ধরো তোমার একটা ছোট বোন আছে, বয়স ৩ বছর। তুমি তাকে কুকুর চেনাতে চাও।

তুমি কী করবে? বইয়ের পাতায় কুকুরের ছবি দেখিয়ে বলবে: "এটা কুকুর।" তারপর রাস্তায় একটা কুকুর দেখলে বলবে: "এটাও কুকুর।" আরেকটা ছবি দেখিয়ে: "এটাও কুকুর।"

কিছুদিন পর তোমার বোন নিজেই নতুন কুকুর দেখলে চিনে ফেলতে পারে — কারণ সে **অভিজ্ঞতা (Experience) থেকে শিখেছে।**

Machine Learning-ও ঠিক এটাই করে। একটা computer program অভিজ্ঞতা থেকে শেখে এবং নতুন পরিস্থিতিতে সঠিক কাজ করতে পারে।

### Tom Mitchell-এর সংজ্ঞা (১৯৯৮) — এটা মুখস্থ করো

> *"A computer program is said to **learn** from experience **E** with respect to some task **T** and some performance measure **P**, if its performance on **T**, as measured by **P**, improves with experience **E**."*

বাংলায় বললে:
- **E (Experience):** প্রোগ্রাম যা যা দেখেছে / শিখেছে (training data)
- **T (Task):** যে কাজটা করতে হবে (যেমন: spam email চেনা)
- **P (Performance):** কতটা ভালো করছে সেটা মাপার উপায় (যেমন: কতটা % সঠিক)

**উদাহরণ:**
- E = ১০,০০০ email (spam/not spam labeled)
- T = নতুন email spam কিনা বলা
- P = সঠিক classification-এর হার (accuracy)

যদি বেশি email দেখার পর accuracy বাড়ে → সে **learn** করেছে।

---

### AI > ML > Deep Learning — এটা কেন জানা দরকার?

অনেক ছাত্র এই তিনটাকে গুলিয়ে ফেলে।

চিন্তা করো তিনটা concentric circle (একটার মধ্যে আরেকটা):

```
┌─────────────────────────────────┐
│   Artificial Intelligence (AI)  │
│  ┌────────────────────────────┐ │
│  │  Machine Learning (ML)     │ │
│  │  ┌──────────────────────┐  │ │
│  │  │   Deep Learning (DL) │  │ │
│  │  └──────────────────────┘  │ │
│  └────────────────────────────┘ │
└─────────────────────────────────┘
```

- **AI** সবচেয়ে বড় — যেকোনো "বুদ্ধিমান" কাজ করানো (chess খেলা, মানুষের মতো কথা বলা, ইত্যাদি)
- **ML** — AI-র একটা অংশ — data থেকে শেখার পদ্ধতি
- **DL** — ML-এর একটা অংশ — neural network দিয়ে শেখা

**Quiz trap:** "Deep Learning is a subset of Artificial Intelligence" → TRUE  
"Machine Learning is a subset of Deep Learning" → **FALSE** (উল্টো!)

---

## ২. Data — এটা বোঝা সবকিছুর ভিত্তি

### Dataset কী?

ধরো তুমি ঢাকায় বাড়ির দাম predict করতে চাও। তুমি ১০০টা বাড়ির তথ্য সংগ্রহ করলে:

| বাড়ি নং | আয়তন (sqft) | বেডরুম | দাম (লাখ টাকা) |
|---------|------------|--------|--------------|
| ১ | ১২০০ | ৩ | ৮০ |
| ২ | ৮০০ | ২ | ৫৫ |
| ৩ | ১৫০০ | ৪ | ১১০ |
| ... | ... | ... | ... |

এটাই dataset। Mathematically:

$$\mathcal{D} = \{(\mathbf{x}^{(1)}, y^{(1)}), (\mathbf{x}^{(2)}, y^{(2)}), \ldots, (\mathbf{x}^{(n)}, y^{(n)})\}$$

এখানে:
- **x^(i)** = i-তম বাড়ির features (আয়তন, বেডরুম সংখ্যা) → **input**
- **y^(i)** = i-তম বাড়ির দাম → **output/target**
- **n** = মোট বাড়ির সংখ্যা (এখানে ১০০)

### গুরুত্বপূর্ণ notation (এগুলো গুলিয়ে ফেলো না!)

| Symbol | মানে | উদাহরণ |
|--------|-----|--------|
| **X** (বড় হাতে) | Input space — সব সম্ভাব্য input-এর জগৎ | সব সম্ভাব্য বাড়ির তথ্য |
| **Y** (বড় হাতে) | Output space — সব সম্ভাব্য output-এর জগৎ | সব সম্ভাব্য দাম |
| **x^(i)** | i-তম observation-এর input vector | i-তম বাড়ির features |
| **y^(i)** | i-তম observation-এর output | i-তম বাড়ির দাম |
| **x_j** | j-তম feature সব observation-এর জন্য | সব বাড়ির আয়তন |
| **p** | Feature-এর সংখ্যা (dimension) | এখানে p=2 (আয়তন + বেডরুম) |

---

## ৩. Data Generating Process — এখানেই বেশিরভাগ ছাত্র আটকে যায়

এটা একটু abstract — কিন্তু একটা গল্প দিয়ে সহজ করি।

### গল্প: ঢাকার বাড়ির বাজার

ঢাকায় হাজার হাজার বাড়ি আছে। প্রতিটা বাড়ি তৈরি হয়েছে কিছু কারণে:
- কোন এলাকায় (দাম বেশি/কম)
- কত বড় (sqft বেশি → দাম বেশি)
- কত পুরনো

এই পুরো "বাজারের নিয়ম" কে বলা হয় **P_xy** — একটা probability distribution যা X × Y-এর উপর defined।

**সবচেয়ে গুরুত্বপূর্ণ কথা: P_xy সাধারণত আমরা জানি না।**

আমরা শুধু সেই বাজার থেকে কিছু sample দেখতে পাই (আমাদের dataset)। পুরো বাজারের সব তথ্য আমাদের কাছে নেই।

### i.i.d. assumption

মনে করো তুমি ১০০টা বাড়ি random-ভাবে বেছে নিয়েছ। এখানে দুটো assumption আছে:

1. **Independent:** একটা বাড়ির দাম অন্য বাড়ির selection-কে affect করেনি
2. **Identically distributed:** সবগুলো বাড়ি same "বাজারের নিয়ম" (P_xy) থেকে এসেছে

এটাকে বলে **i.i.d.** = **i**ndependently and **i**dentically **d**istributed।

### p(x|θ) — এই notation নিয়ে অনেকে confused হয়

Slide-এ একটা important note আছে: `p(x|θ)`-তে `|` মানে **Bayesian conditioning না।**

এটা শুধু readability-র জন্য লেখা হয়। মানে হলো "parameter θ দিয়ে defined distribution থেকে x-এর probability।"

সহজ ভাষায়: **এই course frequentist perspective থেকে পড়ানো হচ্ছে।**

---

## ৪. Supervised Learning-এর তিন ধরনের Task

এটা বোঝার জন্য শুধু output space Y-এর দিকে তাকাও।

### ৪.১ Regression — যখন output একটা number

**উদাহরণ:** বাড়ির দাম predict করা।
- Output: y ∈ R (যেকোনো real number হতে পারে)
- **Residual** = actual - predicted = y − f(x)

```
actual দাম = ৮০ লাখ
predicted দাম = ৭৫ লাখ
residual = ৮০ − ৭৫ = ৫ লাখ (আমরা ৫ লাখ কম বললাম)
```

**Notation:**
- g = 1 → univariate response (একটা output)
- g > 1 → multi-target regression (একাধিক output — এই course-এ শুধু g=1 দেখব)

### ৪.২ Classification — যখন output একটা category

**উদাহরণ:** email spam কিনা বলা, রোগ আছে কিনা বলা।
- Output: y ∈ {C_1, C_2, ..., C_g} (নির্দিষ্ট কিছু ক্লাস)
- **g = 2:** Binary classification (spam/not spam)
- **g > 2:** Multiclass classification (৩ বা তার বেশি class)

**Quiz trap:** "Binary classification requires g > 2" → **FALSE** (g = 2)
"Multiclass classification requires g ≥ 3" → **TRUE** (g > 2 মানেই g ≥ 3)

### ৪.৩ Density Estimation

**উদাহরণ:** একটা input x দিলে p(y|x) পুরো distribution predict করা।

**Quiz trap:** "Density estimation is NOT a supervised task" → **FALSE** — এটা supervised task-ই।

---

## ৫. Model এবং Hypothesis Space — এটা অনেকেই গুলিয়ে ফেলে

### Model কী?

একটা model হলো একটা function:

$$f : \mathcal{X} \to \mathbb{R}^g$$

এটা input নেয়, output দেয়। কিন্তু আউটপুট সবসময় **score** (real number) — সরাসরি class label না।

**কেন score? class label কেন না?**
→ Optimization সহজ হয় continuous value-র উপর করলে।
→ Score → class এ convert করা সহজ, কিন্তু class → score করা impossible।

### Hypothesis Space (H) কী?

ধরো তুমি একটা curve আঁকতে চাও data দেখে। তুমি ঠিক করলে:
- শুধু সরল রেখা (straight line) আঁকব → H = {সব সম্ভাব্য সরল রেখা}
- অথবা curve আঁকব → H = {সব সম্ভাব্য polynomial curve}

**H = সব সম্ভাব্য model যেগুলো থেকে আমরা বেছে নিতে পারি।**

**Quiz trap:** "The hypothesis space H contains all possible datasets" → **FALSE**
H-তে থাকে model (functions), dataset না।

---

## ৬. Inducing Algorithm — Machine Learning-এর Engine

এটা ML-এর সবচেয়ে গুরুত্বপূর্ণ concept।

### সহজ ভাষায়:

Inducer একটা "recipe" যা data দেখে সেরা model বের করে।

$$\mathcal{I}_{L,O} : (\mathcal{X} \times \mathcal{Y})^n \to \mathcal{H}$$

মানে: dataset D নিয়ে একটা model f ∈ H বানাও।

### তিনটা উপাদান — এটা সবচেয়ে বেশি quiz-এ আসে

**Domingos (2012)-এর বিখ্যাত equation:**

> **Learning = Representation + Cost function + Optimization**

#### উপাদান ১: Representation (Hypothesis Space H)

তুমি কী ধরনের model শিখতে চাও সেটা define করা।

| Representation | মানে |
|---------------|-----|
| Linear functions | সরল রেখা দিয়ে model |
| Decision trees | if-else দিয়ে decision |
| Neural networks | brain-এর মতো connected nodes |
| Neighbors | কাছের data point দেখে সিদ্ধান্ত |

#### উপাদান ২: Cost Function (Loss Function L)

"এই model কতটা খারাপ?" — এটা measure করে।

| Loss function | ব্যবহার |
|--------------|--------|
| Squared error | regression |
| Misclassification | classification |
| Likelihood | probabilistic model |

#### উপাদান ৩: Optimization

"সবচেয়ে কম cost-এর model কীভাবে খুঁজব?"

| Method | মানে |
|--------|-----|
| Gradient descent | পাহাড় থেকে নামার মতো সবচেয়ে নিচে যাও |
| Quadratic programming | mathematical optimization |
| Genetic algorithms | evolution-এর মতো ভালো solution বেছে নাও |

### উদাহরণ: Linear Regression

- **Representation:** H = {f(x) = θ^T x̃ | θ ∈ R^{p+1}} → সরল রেখার family
- **Cost:** SSE = Σ(y^(i) − f(x^(i)))² → কত ভুল হচ্ছে তার যোগফল
- **Optimization:** Analytically solve করা যায় (derivative = 0 করে)

**Quiz trap:** "SSE in linear regression requires numerical optimization like gradient descent" → **FALSE**  
SSE analytically minimize করা যায় — কোনো iterative method দরকার নেই!

---

## ৭. Generalization — এটাই ML-এর আসল লক্ষ্য

### গল্প: পরীক্ষার প্রস্তুতি

ধরো তোমার দুই বন্ধু আছে — রাফি এবং সাকিব।

**রাফি** প্রতিটা past exam-এর প্রশ্ন মুখস্থ করে। সে past exam-এ ১০০% পাবে।

**সাকিব** concept বোঝে। Past exam-এ হয়তো ৮৫% পাবে, কিন্তু নতুন প্রশ্নেও ভালো করবে।

Real পরীক্ষায় (নতুন প্রশ্ন) — কে ভালো করবে? **সাকিব।**

Machine Learning-এ এই "real exam performance"-কে বলা হয় **Generalization Error।**

### Generalization Error এর সংজ্ঞা

$$GE(\hat{f}) = \mathbb{E}_{(\mathbf{x},y) \sim \mathbb{P}_{xy}} \left[ L(y, \hat{f}(\mathbf{x})) \right]$$

ভেঙে বলি:
- f̂ = আমাদের trained model (fixed — training শেষ হয়ে গেছে)
- (x, y) ~ P_xy = নতুন, unseen data (random)
- L = loss function
- E = expected value (average over all possible new data)

**মানে:** নতুন data দেখলে model গড়ে কতটা ভুল করবে।

### কেন GE compute করা যায় না?

কারণ P_xy আমরা জানি না! আমাদের কাছে শুধু কিছু sample আছে, পুরো distribution না।

---

## ৮. Inner Loss vs Outer Loss — এখানে ৯০% ছাত্র ভুল করে

এটা quiz-এ প্রায় নিশ্চিত আসবে। মনোযোগ দাও।

### গল্প দিয়ে বোঝাই

ধরো তুমি একটা university-তে ভর্তি হতে চাও। তারা বলল:

- **ভর্তি পরীক্ষা (MCQ):** এটা দিয়ে তোমাকে select করা হবে।
- **আসল performance:** University পরে দেখবে তুমি কতটা ভালো student।

**Inner Loss = ভর্তি পরীক্ষা** (এটা optimize করে model বানানো হয়)
**Outer Loss = আসল performance** (এটা দিয়ে পরে model evaluate করা হয়)

| | Inner Loss | Outer Loss |
|--|-----------|-----------|
| **কখন** | Model training-এর সময় | Training শেষে evaluate করতে |
| **উদ্দেশ্য** | f̂ খুঁজে বের করা | f̂-এর generalization মাপা |
| **কে দেয়?** | Optimizer optimize করে | Application দেয় (বাইরে থেকে) |

**উদাহরণ:**
- Logistic Regression: inner loss = **binomial/Bernoulli loss** (optimize করে)
- কিন্তু আমরা চাই model কতটা % সঠিক classify করে — এটা **outer loss = misclassification rate**

**কেন inner ≠ outer?**  
Outer loss numerically optimize করা কঠিন হতে পারে (যেমন: misclassification rate এর derivative নেই)।

**Quiz traps:**
- "Inner loss is used to assess model performance AFTER training" → **FALSE** (outer loss করে সেটা)
- "Outer loss is optimized during model fitting" → **FALSE** (inner loss optimize হয়)
- "It is always possible to use outer loss as inner loss" → **FALSE** (কঠিন/impossible হতে পারে numerically)

---

## ৯. Generalization Error Estimate করা — Training Error কেন বিশ্বাসযোগ্য না?

### গল্প: নিজের exam নিজে দেওয়া

ধরো তুমি নিজেই তোমার নিজের পরীক্ষা নাও — প্রশ্নও তুমি, উত্তরও তুমি। তুমি নিজে ১০০% পাবে।

কিন্তু এই ১০০% কি তোমার real knowledge বলে? **না।**

এটাই **Training Error** এর সমস্যা। Model তার নিজের training data-তে test করলে সে জানে উত্তর — সে "মুখস্থ" করে ফেলেছে।

### Training Error এর গণনা

$$\widehat{GE}_{\mathcal{D}}(\hat{f}) = \frac{1}{|\mathcal{D}|} \sum_{(\mathbf{x},y) \in \mathcal{D}} L(y, \hat{f}(\mathbf{x}))$$

**এটা biased (optimistic) কারণ:**
- Inducer specifically training data-র error minimize করে f̂ বানিয়েছে
- তাই training data-তে test করলে error কম দেখাবে — **বাস্তব error এর চেয়ে ভালো দেখাবে**

### Holdout Method — সঠিক উপায়

Dataset দুই ভাগ করো:

```
পুরো Dataset D
      │
   ┌──┴──┐
D_train   D_test
(train)   (evaluate)
```

1. **D_train** দিয়ে model fit করো → f̂ পাও
2. **D_test** দিয়ে evaluate করো (model এটা আগে কখনো দেখেনি)

$$\widehat{GE}_{\mathcal{D}_{test}}(\hat{f}) = \frac{1}{|\mathcal{D}_{test}|} \sum_{(\mathbf{x},y) \in \mathcal{D}_{test}} L(y, \hat{f}(\mathbf{x}))$$

এই procedure-কে বলা হয় **holdout।**

### Train Error vs Test Error — সবচেয়ে গুরুত্বপূর্ণ সম্পর্ক

> **Test error ≥ Training error (সাধারণত)**

**কেন?**
- Model training data-এর উপর optimize হয়েছে → training data-তে ভালো করবেই
- Test data নতুন → একটু বেশি ভুল হবে

**Quiz trap (সবচেয়ে বেশি আসে):**
- "Test error is typically less than training error" → **FALSE** (বেশি, কম না)
- "Training error is a pessimistic estimate of generalization error" → **FALSE** (optimistic/biased, pessimistic না)
- "Training error is an optimistic estimate" → **TRUE** ✓

**Special case:** যদি hypothesis training data দেখার আগেই fixed থাকে → expected training error = expected test error। কিন্তু ML-এ hypothesis data দেখে বানানো হয় → এই equality আর থাকে না।

### Cross-Validation

Holdout-এর একটা সমস্যা: একটা নির্দিষ্ট split-এ variance বেশি।

Cross-validation এই সমস্যা কমায়: data-কে k ভাগ করো, k বার train-test করো, average নাও।

---

## ১০. Generalization Error of a Learning Algorithm — একটু গভীরে

এতক্ষণ দেখলাম fixed model f̂-এর GE। এখন দেখব পুরো learning algorithm-এর GE।

### পার্থক্য বোঝো

| | GE of Model f̂ | GE of Algorithm I |
|--|-------------|----------------|
| **f̂** | Fixed (trained once) | Random (depends on training data) |
| **Training data** | Fixed | Random (drawn from P_xy) |
| **Test data** | Random | Random |

Algorithm-level GE:

$$GE_n(\mathcal{I}_{L,O}) = \mathbb{E}_{\mathcal{D}_n \sim \mathbb{P}_{xy}^n, (\mathbf{x},y) \sim \mathbb{P}_{xy}} \left[ L\left(y, \hat{f}_{\mathcal{D}_n}(\mathbf{x})\right) \right]$$

**এখানে উভয়ই random:** training data D_n এবং test observation (x, y)।

---

## ১১. ML Synonyms — এগুলো না জানলে quiz-এ ধরা খাবে

এই table-টা মুখস্থ করো:

| Group | Synonyms (সব একই জিনিস) |
|-------|------------------------|
| Algorithm | inducer, inducing algorithm, learning algorithm, learner |
| Process | learning, training, inducing, fitting |
| Data point | example, instance, observation |
| Input | feature, attribute, covariate, predictor, input variable |
| Output | output, target, response, outcome, dependent variable |
| Loss | cost function, costs, risk |
| Class | label, class, category |

**Quiz trap:** "Feature and predictor refer to different things" → **FALSE** (synonyms!)
"Cost function and risk are different" → **FALSE** (synonyms!)

---

## ১২. সব মিলিয়ে একটা বড় গল্প

এখন পুরো Chapter 1 একটা গল্পে বলি।

তুমি ঢাকার বাড়ির দাম predict করতে চাও।

**Step 1 — Data সংগ্রহ (Dataset):**
১০০টা বাড়ির তথ্য নিলে: x = (আয়তন, বেডরুম), y = দাম।
ধরে নিলে এগুলো P_xy থেকে i.i.d. draw হয়েছে।

**Step 2 — Model বেছে নাও (Hypothesis Space):**
সিদ্ধান্ত নিলে linear model ব্যবহার করবে: H = সব linear functions।

**Step 3 — Loss বেছে নাও (Cost Function):**
SSE = (actual − predicted)² যোগ করব।

**Step 4 — Train করো (Optimization):**
SSE minimize করো analytically → OLS estimator পাওয়া গেল।

**Step 5 — Evaluate (Generalization Error Estimate):**
D_test-এ test করো। Training error না — সেটা optimistic।
Test error দেখো — এটা real-world performance-এর ভালো estimate।

**লক্ষ্য:** নতুন বাড়ি দেখলে সঠিক দাম বলতে পারা — **Generalization।**

---

## ১৩. Quick Revision — Quiz-এর আগে এই পয়েন্টগুলো দেখো

| Statement | True/False | কেন |
|-----------|-----------|-----|
| ML hierarchy: DL ⊂ ML ⊂ AI | TRUE | — |
| P_xy is typically known | FALSE | Unknown, complicated |
| p(x\|θ) implies Bayesian treatment | FALSE | Frequentist notation |
| Binary classification: g = 2 | TRUE | — |
| Multiclass: g > 2 | TRUE | — |
| Density estimation is NOT supervised | FALSE | এটা supervised |
| Residual r = f(x) − y | FALSE | r = y − f(x) |
| H contains datasets | FALSE | H contains models |
| SSE needs numerical optimization | FALSE | Analytically solvable |
| Inner loss = outer loss always | FALSE | Often different |
| Test error ≥ Training error | TRUE | Key result |
| Training error is pessimistic | FALSE | Optimistic/biased |
| In algorithm GE, both D and (x,y) are random | TRUE | — |
| Feature ≠ Predictor | FALSE | Synonyms |
| Holdout: train on D_test, evaluate on D_train | FALSE | উল্টো! |
