# 🤿 অধ্যায় ১: Introduction & Formalization — Deep Dive (গভীর ডুব)
## — বাংলায়, একদম noob-এর জন্য, বন্ধুর মতো করে শেখানো

---

> **হ্যালো বন্ধু!** 👋 এই অধ্যায়টা ASL-এর **ভিত্তি (foundation)**। এখানে কোনো কঠিন অঙ্ক নেই, কিন্তু **সংজ্ঞা আর notation** এত precise যে quiz-এ এক শব্দ এদিক-ওদিক হলেই ফাঁদে পড়বে।
>
> প্রতিটা concept-এ পাবে:
> - 🔊 **উচ্চারণ**
> - 👶 **গল্প** (noob ব্যাখ্যা)
> - 🧮 **গণিত-চিহ্ন বোঝা** (প্রতিটা symbol ভেঙে)
> - ✅❌ **কেন / কেন নয়** (why & why not)
> - 🧠 **মনে রাখার টোটকা**
> - ✍️ শেষে **অঙ্ক + সমাধান**

---

# 📚 সূচিপত্র
1. [Machine Learning আসলে কী?](#১-machine-learning-আসলে-কী)
2. [Data — চিহ্নগুলো বোঝা](#২-data--চিহ্নগুলো-বোঝা)
3. [Data Generating Process](#৩-data-generating-process)
4. [তিন ধরনের Supervised Task](#৪-তিন-ধরনের-supervised-task)
5. [Model কী?](#৫-model-কী)
6. [Inducing Algorithm ও তিন উপাদান](#৬-inducing-algorithm-ও-তিন-উপাদান)
7. [Generalization](#৭-generalization)
8. [Inner Loss vs Outer Loss](#৮-inner-loss-vs-outer-loss)
9. [Generalization Error মাপা: Train vs Test](#৯-generalization-error-মাপা-train-vs-test)
10. [ML Synonyms](#১০-ml-synonyms)
11. [🧠 Master মনে রাখার Table](#১১-master-মনে-রাখার-table)
12. [✍️ অঙ্ক ও সমাধান](#১২-অঙ্ক-ও-সমাধান)

---

# ১. Machine Learning আসলে কী?

🔊 **উচ্চারণ:** *মেশিন লার্নিং*

👶 **গল্প:** তুমি বাচ্চাকে অনেকগুলো বিড়াল-কুকুরের ছবি দেখালে। কেউ নিয়ম শেখায়নি, তবু সে নিজে নিজে পার্থক্য ধরে ফেলল। মেশিনও তেমনি **অভিজ্ঞতা (data) থেকে নিজে নিজে উন্নতি করে** — এটাই ML।

🧮 **গণিত-চিহ্ন (Tom Mitchell-এর সংজ্ঞা):**
> একটা program **শেখে** experience $E$ থেকে, task $T$-এর জন্য, performance measure $P$ অনুযায়ী — যদি $E$ বাড়ার সাথে $T$-তে তার $P$ **উন্নত হয়**।

মনে রাখো: **T**ask, **P**erformance, **E**xperience — "TPE"।

✅❌ **কেন AI ⊃ ML ⊃ DL?** AI সবচেয়ে বড় ছাতা (যেকোনো "বুদ্ধিমান" আচরণ), ML তার ভেতরে (data থেকে শেখা), Deep Learning ML-এর ভেতরে (neural network দিয়ে শেখা)। **কেন উল্টো নয়?** কারণ সব ML-ই deep নয় (যেমন linear regression), কিন্তু সব DL-ই ML।

🧠 **টোটকা:** "TPE — performance improves with experience।"

---

# ২. Data — চিহ্নগুলো বোঝা

👶 **গল্প:** Data হলো শেখার "পাঠ্যবই" — অনেকগুলো উদাহরণ, প্রতিটায় একটা প্রশ্ন (input) আর তার উত্তর (output)।

🧮 **গণিত-চিহ্ন (ভেঙে ভেঙে):**
$$\mathcal{D} = \{(\mathbf{x}^{(1)}, y^{(1)}), \ldots, (\mathbf{x}^{(n)}, y^{(n)})\} \subset (\mathcal{X} \times \mathcal{Y})^n$$

- $\mathcal{D}$ (পড়ো *"ডি"*) = পুরো dataset।
- $\mathbf{x}^{(i)}$ = $i$-তম observation-এর **input** (বোল্ড মানে vector — অনেকগুলো feature)।
- $y^{(i)}$ = তার **output / target**।
- $\mathcal{X}$ = input space, মাত্রা $p$ (সাধারণত $\mathcal{X} \subset \mathbb{R}^p$)।
- $\mathcal{Y}$ = output space।
- উপরে কোঠায় সংখ্যা $^{(i)}$ = কোন observation; নিচে $x_j$ = কোন feature।

✅❌ **কেন $\mathbf{x}^{(i)}$ (উপরে) আর $x_j$ (নিচে) আলাদা?** উপরের index = **কোন data point** (সারি/row), নিচের index = **কোন feature** (কলাম/column)। এই দুটো গুলিয়ে ফেলা #১ ভুল।

🧠 **টোটকা:** "উপরে = কে (which point), নিচে = কী (which feature)।"

---

# ৩. Data Generating Process

🔊 **উচ্চারণ:** *ডেটা জেনারেটিং প্রসেস*

👶 **গল্প:** ধরো প্রকৃতির একটা গোপন "যন্ত্র" আছে যা $(\mathbf{x}, y)$ জোড়া তৈরি করে। আমরা সেই যন্ত্র দেখি না, শুধু তার বের করা data দেখি।

🧮 **গণিত-চিহ্ন:**
- একটা joint distribution $\mathbb{P}_{xy}$ আছে $\mathcal{X} \times \mathcal{Y}$-এর উপর।
- Data তোলা হয় **i.i.d.** (আই-আই-ডি: independent + identically distributed)।
- প্রায়ই parameterized: $p(x, y \mid \theta)$, $\theta \in \Theta$।

✅❌ **`p(x | θ)`-এর `|` কি Bayesian conditioning?** ❌ **না!** এটা frequentist notation — শুধু পড়ার সুবিধার্থে। ✅ এটা মানে "$\theta$ parameter দিয়ে নির্ধারিত distribution", "given $\theta$" এর Bayesian অর্থে নয় (যদি না স্পষ্ট বলা থাকে)। এটা একটা ক্লাসিক quiz-ফাঁদ।

🧠 **টোটকা:** "$\mathbb{P}_{xy}$ অজানা ও জটিল; data তার i.i.d. নমুনা।"

---

# ৪. তিন ধরনের Supervised Task

👶 **গল্প:** Output space $\mathcal{Y}$ কেমন, তার উপর task-এর ধরন নির্ভর করে।

🧮 **গণিত-চিহ্ন:**

| Task | Output $\mathcal{Y}$ | মানে |
|------|------|------|
| **Regression** | $\mathcal{Y} \subset \mathbb{R}^g$, $1 \le g < \infty$ | $g=1$ → univariate; $g>1$ → multi-target |
| **Classification** | $\mathcal{Y} = \{C_1,\ldots,C_g\}$, $g \ge 2$ | $g=2$ → binary; $g>2$ → multiclass |
| **Density estimation** | $\mathcal{Y}$-এর উপর $p(y\mid x)$ অনুমান | — |

- **Regression detail:** continuous output, $f(\mathbf{x})$ = prediction = **score**। **Residual** $r = y - f(\mathbf{x})$।
- **Classification detail:** class label বা membership probability $\pi_1,\ldots,\pi_g \in [0,1]$।

✅❌ **Density estimation কি supervised?** ✅ **হ্যাঁ** (এই কোর্সে এটাকে supervised task হিসেবে ধরা হয়)। **কেন noob রা ভুল করে?** কারণ "density estimation" শুনলে unsupervised মনে হয় — কিন্তু এখানে নয়।

✅❌ **$g=1$ মানে কী?** ✅ univariate (একটা output)। ❌ binary নয় — সেটা classification-এর $g=2$।

🧠 **টোটকা:** Regression = "কত?" (সংখ্যা), Classification = "কোনটা?" (দল), Density = "বণ্টন কেমন?"

---

# ৫. Model কী?

🧮 **গণিত-চিহ্ন:** একটা model হলো function
$$f : \mathcal{X} \to \mathbb{R}^g$$

- Output = **scores** (প্রতি input-এ $g$টা বাস্তব সংখ্যা)।
- Regression-এ: score-ই সরাসরি prediction।
- Classification-এ: score → class/probability তে রূপান্তর।
- $\hat y := f(\mathbf{x})$ = prediction। $\mathcal{H}$ = **hypothesis space** (সব অনুমোদিত model-এর সেট)।

✅❌ **Model class দেয় নাকি score?** ✅ **score**। **কেন class নয়?** কারণ score-এ বেশি তথ্য আর optimization সহজ; score → class করা যায়, উল্টোটা নয়।

🧠 **টোটকা:** "$\hat y$ = f(x), output সবসময় score।"

---

# ৬. Inducing Algorithm ও তিন উপাদান

🔊 **উচ্চারণ:** *ইনডিউসিং অ্যালগরিদম*

👶 **গল্প:** Inducer হলো সেই "শিক্ষক-যন্ত্র" যা data $\mathcal{D}$ নিয়ে একটা শেখা model $\hat f$ বের করে দেয়।

🧮 **গণিত-চিহ্ন:**
$$\mathcal{I}_{L,O} : (\mathcal{X} \times \mathcal{Y})^n \to \mathcal{H}$$

- $L : \mathcal{Y} \times \mathbb{R}^p \to \mathbb{R}_{\ge 0}$ = **loss function**।
- $O$ = **optimizer**।
- এটা চালানো = **training / fitting**। ফল: $\hat f(\mathbf{x}) = \mathcal{I}_{L,O}(\mathcal{D})$।

✅❌ **Inducer কি $\mathcal{D} \to \mathcal{H}$ map করে?** ❌ **না!** এটা $\mathcal{D} \to f \in \mathcal{H}$ map করে — একটা **নির্দিষ্ট model**, পুরো space নয়। (ফাঁদ!)

**🍕 Learning-এর তিন উপাদান (Domingos):**
$$\text{Learning} = \text{Representation} + \text{Cost} + \text{Optimization}$$

| উপাদান | কাজ | উদাহরণ |
|--------|-----|--------|
| **Representation** | কোন model শেখা যাবে ($\mathcal{H}$) | linear, tree, neural net |
| **Cost** | ভালো-খারাপ আলাদা করা | squared error, likelihood |
| **Optimization** | $\mathcal{H}$-এ সেরাটা খোঁজা | gradient descent, QP |

🧠 **টোটকা:** "RCO — Representation + Cost + Optimization।"

---

# ৭. Generalization

🔊 **উচ্চারণ:** *জেনারালাইজেশন*

👶 **গল্প:** পরীক্ষায় আসল কৃতিত্ব = পড়া প্রশ্ন মুখস্থ নয়, বরং **নতুন প্রশ্নে** ভালো করা। ML-এর আসল লক্ষ্যও তাই — training data ছাড়িয়ে নতুন data-তে ভালো করা।

🧮 **গণিত-চিহ্ন (Generalization Error):**
$$GE(\hat f) = \mathbb{E}_{(\mathbf{x},y) \sim \mathbb{P}_{xy}}\big[L(y, \hat f(\mathbf{x}))\big]$$

- $\hat f$ স্থির (training data fixed)। Test point $(\mathbf{x}, y)$ random।
- ⚠️ **হুবহু হিসাব করা যায় না** — কারণ $\mathbb{P}_{xy}$ অজানা।

✅❌ **GE কি ঠিকঠাক বের করা যায়?** ❌ না, কারণ $\mathbb{P}_{xy}$ অজানা; তাই আমরা test set দিয়ে **আনুমানিক** মাপি।

🧠 **টোটকা:** "GE = প্রত্যাশিত loss নতুন data-তে; মাপা যায় না, অনুমান করা যায়।"

---

# ৮. Inner Loss vs Outer Loss

👶 **গল্প:** **Inner loss** = পড়ার সময় তুমি যেটা minimize করো (সহজে অঙ্ক করা যায়)। **Outer loss** = পরীক্ষক যেটা দিয়ে তোমাকে আসলে মাপে (application-এর চাহিদা)।

🧮 **চিহ্ন:**
- Inner = model fit করার সময় optimize করা loss।
- Outer = পরে performance মাপার loss।

✅❌ **Inner আর outer কি একই হতে হবে?** ❌ না। উদাহরণ: logistic regression **binomial loss** (inner) minimize করে, কিন্তু আমরা **AUC বা misclassification rate** (outer) দিয়ে মাপি। **কেন আলাদা?** কারণ outer loss প্রায়ই সরাসরি optimize করা কঠিন (যেমন AUC differentiable নয়)।

🧠 **টোটকা:** "Inner = train করতে, Outer = মাপতে।"

---

# ৯. Generalization Error মাপা: Train vs Test

🧮 **Train error (biased):**
$$\widehat{GE}_{\mathcal{D}}(\hat f) = \frac{1}{|\mathcal{D}|}\sum_{(\mathbf{x},y)\in\mathcal{D}} L(y, \hat f(\mathbf{x}))$$

🧮 **Holdout / Test error:** $\mathcal{D} = \mathcal{D}_{\text{train}} \cup \mathcal{D}_{\text{test}}$, train-এ fit, test-এ মাপো।

> **মূল সম্পর্ক:** $\;\text{Test error} \ge \text{Train error}$ (সাধারণত)।

✅❌ **কেন train error optimistic (কম দেখায়)?** কারণ model **ওই train data-র উপরেই** fit হয়েছে — সে train data "চেনে"। তাই train error আসল performance-কে **কম করে দেখায় (underestimate)**, biased। ❌ **কেন pessimistic নয়?** কারণ model train data-তে advantage পায়, ভুল কম দেখায়, বেশি নয়।

✅❌ **কখন train error = test error (প্রত্যাশায়)?** যদি hypothesis **data দেখার আগেই** fix করা থাকে। কিন্তু ML-এ inducer train data-তে error minimize করে → তাই optimistic।

🧠 **টোটকা:** "Train চেনা প্রশ্ন (সহজ), Test অচেনা প্রশ্ন (কঠিন) → Test ≥ Train।"

---

# ১০. ML Synonyms

👶 একই জিনিসের অনেক নাম — quiz-এ এক নাম দিয়ে অন্যটা ধরিয়ে দেয়।

| দল | প্রতিশব্দ |
|----|----------|
| Algorithm | inducer, learner, learning algorithm |
| Process | learning, training, fitting, inducing |
| Data point | example, instance, observation |
| Input | feature, attribute, covariate, predictor |
| Output | target, label, response, outcome, dependent variable |
| Loss | cost, risk |
| Class | label, category |

🧠 **টোটকা:** "Feature = covariate = predictor; Target = response = outcome।"

---

# ১১. Master মনে রাখার Table

| বিষয় | সঠিক ✅ | ফাঁদ ❌ |
|------|---------|---------|
| Test vs Train error | Test **≥** Train | "Test < Train" |
| Train error | **optimistic** (biased) | "pessimistic" |
| Density estimation | **supervised** task | "unsupervised" |
| Inducer map করে | $\mathcal{D} \to f \in \mathcal{H}$ | "$\mathcal{D} \to \mathcal{H}$" |
| Regression $g=1$ | **univariate** | "binary" |
| Classification | binary $g=2$, multiclass $g>2$ | উল্টানো |
| `p(x\|θ)`-এর `\|` | frequentist (পড়ার সুবিধা) | "Bayesian conditioning" |
| Outer loss | performance **মাপতে** | "model train করতে" |
| Linear reg. SSE | **analytically** minimize | "শুধু numerical" |
| Learning = | **R + C + O** | শুধু একটা |

---

# ১২. অঙ্ক ও সমাধান

> আগে নিজে চেষ্টা করো, তারপর সমাধান দেখো। ✍️

---

### 🧩 সমস্যা ১ — Residual হিসাব
Data point: সত্য $y = 10$, model prediction $f(\mathbf{x}) = 7$। Residual বের করো।

**সমাধান:** $r = y - f(\mathbf{x}) = 10 - 7 = 3$। (model ৩ কম বলেছে।)

---

### 🧩 সমস্যা ২ — Train error হিসাব (L2 loss)
৩টা train point: predictions $\{5, 8, 10\}$, সত্য $\{6, 8, 13\}$। Squared loss দিয়ে গড় train error বের করো।

**সমাধান:** residual: $6-5=1,\; 8-8=0,\; 13-10=3$।
Squared: $1, 0, 9$। যোগ = $10$। গড় = $10/3 \approx 3.33$।

---

### 🧩 সমস্যা ৩ — Task ধরা
(ক) বাড়ির দাম (টাকায়) ভবিষ্যদ্বাণী। (খ) email spam/ham। (গ) ছবিতে digit ০–৯ চেনা। প্রতিটার task ও $g$ বলো।

**সমাধান:**
- (ক) **Regression**, output continuous ($g=1$, univariate)।
- (খ) **Classification**, binary, $g=2$।
- (গ) **Classification**, multiclass, $g=10$।

---

### 🧩 সমস্যা ৪ — কেন test ≥ train? (ব্যাখ্যা)
এক বন্ধু বলল "আমার model train-এ ২% error, test-এ ৯% — model খারাপ?" ব্যাখ্যা করো কী ঘটছে।

**সমাধান:** এটা স্বাভাবিক — train error সবসময় optimistic (model train data চেনে)। ২% vs ৯% বড় ব্যবধান = **overfitting**-এর ইঙ্গিত (generalization gap বড়)। model train data-র random quirk মুখস্থ করেছে। প্রতিকার: regularization বা বেশি data।

---

### 🧩 সমস্যা ৫ — তিন উপাদান চেনা (linear regression)
Linear regression-এর Representation, Cost, Optimization আলাদা করে লেখো।

**সমাধান:**
- **Representation:** $\mathcal{H} = \{f(\mathbf{x}) = \boldsymbol\theta^\top \tilde{\mathbf{x}} \mid \boldsymbol\theta \in \mathbb{R}^{p+1}\}$।
- **Cost:** SSE $= \sum (y^{(i)} - f(\mathbf{x}^{(i)}))^2 = \lVert \mathbf{y} - \mathbf{X}\boldsymbol\theta\rVert^2$।
- **Optimization:** SSE **analytically** minimize ($\theta$-এর সাপেক্ষে derivative = 0; numerical লাগে না)।

---

> 🎓 **শেষ কথা:** অধ্যায় ১ = ভাষা ও নিয়ম শেখা। চারটা সোনার সূত্র: **(১)** Test ≥ Train, train optimistic; **(২)** Inducer দেয় একটা model, পুরো space নয়; **(৩)** Learning = R+C+O; **(৪)** `p(x|θ)`-এর `|` Bayesian নয়। এগুলো ধরলেই বাকি সব অধ্যায় সহজ হবে। 🍼💪
