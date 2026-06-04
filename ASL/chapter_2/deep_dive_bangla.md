# 🤿 অধ্যায় ২: Classification — Deep Dive (গভীর ডুব)
## — বাংলায়, একদম noob-এর জন্য, বন্ধুর মতো করে শেখানো

---

> **হ্যালো বন্ধু!** 👋
> ধরে নিচ্ছি তুমি classification-এর কিছুই জানো না — কোনো সমস্যা নেই। আমি তোমার পাশে বসা এক বড় ভাই, যে চা খেতে খেতে পুরো জিনিসটা গল্পের মতো বুঝিয়ে দেবে।
>
> প্রতিটা concept-এর জন্য তুমি পাবে:
> - 🔊 **উচ্চারণ** (শব্দটা মুখে কীভাবে বলবে)
> - 👶 **গল্প** (noob হিসেবে বোঝা)
> - 🧮 **গণিত দিয়ে কীভাবে ভাববে**
> - 🧠 **কীভাবে মনে রাখবে** (টোটকা)
> - ✍️ শেষে **অঙ্ক + সমাধান**
>
> তাড়াহুড়ো নয়। এক চুমুক চা, এক concept। চলো ডুব দিই। 🤿

---

# 📚 সূচিপত্র

1. [Classification আসলে কী?](#১-classification-আসলে-কী)
2. [Model কেন score দেয়, class না?](#২-model-কেন-score-দেয়-class-না)
3. [Scoring Classifier](#৩-scoring-classifier)
4. [Probabilistic Classifier](#৪-probabilistic-classifier)
5. [Score ↔ Probability ↔ Class — একমুখী রাস্তা](#৫-score--probability--class--একমুখী-রাস্তা)
6. [Decision Boundary](#৬-decision-boundary)
7. [Linear Classifier](#৭-linear-classifier)
8. [Sigmoid ও Logistic Function](#৮-sigmoid-ও-logistic-function)
9. [Softmax Function](#৯-softmax-function)
10. [Generative vs Discriminative (LDA, QDA, Naive Bayes)](#১০-generative-vs-discriminative)
11. [🧠 মনে রাখার Master Table](#১১-মনে-রাখার-master-table)
12. [✍️ অঙ্ক ও সমাধান (Practice Problems)](#১২-অঙ্ক-ও-সমাধান)

---

# ১. Classification আসলে কী?

🔊 **উচ্চারণ:** *ক্ল্যাসিফিকেশন* (clas-uh-fih-KAY-shun)

👶 **গল্প:** তোমার সামনে একটা ফল। প্রশ্ন: "এটা আম, না কাঁঠাল, না লিচু?" — তুমি একটা **নির্দিষ্ট দল (class)** বেছে দিচ্ছ। এটাই classification। output টা সংখ্যা না (যেমন দাম), বরং একটা **শ্রেণি/লেবেল**।

🧮 **গণিত দিয়ে ভাবো:** output আসে একটা সসীম সেট থেকে —
$$y \in \mathcal{Y} = \{C_1, C_2, \ldots, C_g\}, \qquad 2 \le g < \infty$$

- $g = 2$ → **binary** (দুই দল): $\mathcal{Y} = \{0, 1\}$ অথবা $\{-1, +1\}$।
- $g \ge 3$ → **multiclass**: $\mathcal{Y} = \{1, 2, \ldots, g\}$।

🧠 **মনে রাখার টোটকা:** Regression = "**কত**?" (সংখ্যা)। Classification = "**কোনটা**?" (দল)। 🔢 vs 🏷️

---

# ২. Model কেন score দেয়, class না?

👶 **গল্প:** তুমি জিজ্ঞেস করলে "এটা কি বিড়াল?" Model সরাসরি "হ্যাঁ/না" না বলে বলে "আমি **৮৭% নিশ্চিত** বিড়াল"। এই ৮৭% হলো **score**। কেন এভাবে?

1. **Optimization সহজ** — continuous সংখ্যার উপর গণিত করা সহজ (derivative নেওয়া যায়)।
2. **Score-এ বেশি তথ্য** — শুধু "বিড়াল" বলার চেয়ে "৮৭% বিড়াল" অনেক বেশি কথা বলে।
3. Score → class করা **সহজ** (একটা threshold দাও); কিন্তু class → score করা **অসম্ভব** (একবার "বিড়াল" বলে দিলে আর % ফেরত পাবে না)।

🧮 **গণিত:** Model হলো $f : \mathcal{X} \to \mathbb{R}^g$ — অর্থাৎ প্রতিটা class-এর জন্য একটা বাস্তব সংখ্যা (score) বের করে।

🧠 **টোটকা:** "Score → Class এক ক্লিকে; Class → Score কখনো না।" একমুখী রাস্তা, like টুথপেস্ট টিউব থেকে বের হলে আর ঢোকে না। 🪥

---

# ৩. Scoring Classifier

🔊 **উচ্চারণ:** *স্কোরিং ক্ল্যাসিফায়ার*

👶 **গল্প:** প্রতিটা দলকে একটা "নম্বর" দাও। যে দল সবচেয়ে বেশি নম্বর পায়, সেই দল জেতে।

🧮 **গণিত (multiclass):** $g$টা scoring function $f_1, \ldots, f_g$। প্রেডিকশন =
$$h(\mathbf{x}) = \underset{k \in \{1,\ldots,g\}}{\arg\max}\; f_k(\mathbf{x})$$

🔊 $\arg\max$ = *"আর্গ ম্যাক্স"* = "যে $k$-এর জন্য মান সবচেয়ে বড়, সেই $k$"।

🧮 **Binary case ($g=2$):** একটা function-ই যথেষ্ট, $f(\mathbf{x}) = f_1(\mathbf{x}) - f_{-1}(\mathbf{x})$।
$$h(\mathbf{x}) = \operatorname{sgn}(f(\mathbf{x})), \qquad h(\mathbf{x}) = 1 \iff f(\mathbf{x}) > 0$$
- $|f(\mathbf{x})|$ = **confidence** (কতটা নিশ্চিত)। 🔊 *কনফিডেন্স*।

🧠 **টোটকা:** score-এর **চিহ্ন (sign)** দল ঠিক করে; score-এর **মান (size)** confidence ঠিক করে। চিহ্ন = সিদ্ধান্ত, মাপ = জোর।

---

# ৪. Probabilistic Classifier

🔊 **উচ্চারণ:** *প্রবাবিলিস্টিক ক্ল্যাসিফায়ার*

👶 **গল্প:** এখানে নম্বর গুলো সুন্দর করে **শতাংশে (০ থেকে ১)** দেওয়া হয়, আর সব দলের শতাংশ যোগ করলে **১** (১০০%) হয়।

🧮 **গণিত:** $\pi_1, \ldots, \pi_g : \mathcal{X} \to [0,1]$ with $\sum_l \pi_l = 1$।
$$h(\mathbf{x}) = \underset{k}{\arg\max}\; \pi_k(\mathbf{x})$$

🔊 $\pi$ = *"পাই"* (এখানে probability বোঝায়, ৩.১৪ নয়!)।

🧮 **Binary:** একটা $\pi(\mathbf{x})$, threshold $c$ দিয়ে: $h(\mathbf{x}) = \mathbb{1}(\pi(\mathbf{x}) \ge c)$। **Default $c = 0.5$**।

🧠 **টোটকা:** Probabilistic = Scoring + "সুন্দর জামা" (০–১ এ বাঁধা, যোগফল ১)। তাই **probabilistic-কে scoring হিসেবেও দেখা যায়**, কিন্তু উল্টোটা সবসময় নয়।

---

# ৫. Score ↔ Probability ↔ Class — একমুখী রাস্তা

🧮 **নিয়ম গুলো (এটা quiz-এ আসেই):**

| থেকে → যেতে | কীভাবে | সম্ভব? |
|------------|--------|--------|
| Probability → Class | Thresholding | ✅ |
| Score → Class | Thresholding | ✅ |
| Score → Probability | Calibrating / Scaling (sigmoid) | ✅ |
| **Class → Score বা Probability** | — | ❌ **অসম্ভব** |

🧠 **টোটকা — "নিচে নামা সহজ, উপরে ওঠা কঠিন":**
Probability/Score উপরে, Class নিচে। উপর থেকে নিচে (threshold) সহজ; নিচ থেকে উপরে (class → score) ওঠা যায় না। 🪜

---

# ৬. Decision Boundary

🔊 **উচ্চারণ:** *ডিসিশন বাউন্ডারি* (সিদ্ধান্ত-সীমানা)

👶 **গল্প:** মানচিত্রে দুই দেশের **সীমান্ত রেখা**। এক পাশে "বিড়াল দেশ", আরেক পাশে "কুকুর দেশ"। ঠিক সীমান্তের উপর দাঁড়ালে model দ্বিধায় পড়ে (tie)।

🧮 **গণিত:** decision region —
$$\mathcal{X}_k = \{\mathbf{x} \in \mathcal{X} : h(\mathbf{x}) = k\}$$
Binary boundary: $f(\mathbf{x}) = c$ (scoring-এ $c=0$, probabilistic-এ $c=0.5$)।

🧠 **টোটকা:** Boundary = যেখানে দুই দলের score **সমান** ($f_i = f_j$) — "টাই" এর জায়গা।

---

# ৭. Linear Classifier

🔊 **উচ্চারণ:** *লিনিয়ার ক্ল্যাসিফায়ার*

👶 **গল্প:** সীমান্তটা যদি একটা **সোজা রেখা/সমতল** হয়, তবে সেটা linear classifier।

🧮 **গণিত:** যদি (হয়তো একটা monotone transform $g$-এর পরে) লেখা যায়
$$g(f_k(\mathbf{x})) = \mathbf{w}_k^\top \mathbf{x} + b_k$$
তাহলে linear। দুই দলের tie =
$$(\mathbf{w}_i - \mathbf{w}_j)^\top \mathbf{x} + (b_i - b_j) = 0 \quad(\text{একটা hyperplane})$$

🔊 $\mathbf{w}$ = *"ওয়েট/w"*, $b$ = *"বায়াস/b"*, hyperplane = *"হাইপারপ্লেন"*।

> ⚠️ **মস্ত বড় ফাঁদ:** linear classifier **মূল input space-এ non-linear boundary** বানাতে পারে — যদি তুমি feature বাড়াও (polynomial, basis function)। "Linear" মানে parameter-এ linear, ছবিতে নয়!

🧠 **টোটকা:** "Linear in **weights**, not necessarily in the **picture**." 🎨

---

# ৮. Sigmoid ও Logistic Function

🔊 **উচ্চারণ:** *সিগময়েড* (SIG-moyd), *লজিস্টিক* (luh-JIS-tik)

👶 **গল্প:** Score যেকোনো সংখ্যা হতে পারে (−১০০০ থেকে +১০০০)। কিন্তু probability তো ০–১ এর মধ্যে দরকার। Sigmoid হলো একটা "S-আকৃতির চাপুনি যন্ত্র" — যেকোনো সংখ্যাকে চেপে ০–১ এ এনে ফেলে।

🧮 **গণিত — Logistic (সবচেয়ে গুরুত্বপূর্ণ):**
$$s(t) = \frac{1}{1 + e^{-t}}$$

বৈশিষ্ট্য:
- $t \to -\infty \Rightarrow s \to 0$; $\;t \to +\infty \Rightarrow s \to 1$।
- $s(0) = 0.5$ — **(0, ½) বিন্দুতে symmetric**।
- **Derivative (মুখস্থ করো!):** $\;\dfrac{\partial s}{\partial t} = s(t)\,(1 - s(t))$।

🧠 **টোটকা:**
- "**S** for **S**igmoid for **S**quash (চেপে ০–১)।"
- Derivative মনে রাখো: "$s$ গুণ (১ − $s$)" — নিজেকেই দিয়ে তৈরি, সুন্দর! 💡
- Logistic regression = score → logistic sigmoid → probability।

---

# ৯. Softmax Function

🔊 **উচ্চারণ:** *সফটম্যাক্স*

👶 **গল্প:** Logistic ছিল ২ দলের জন্য। ৩+ দল হলে? Softmax! এটা সব দলের score নিয়ে exponent বসায়, তারপর normalize করে — সব probability ০–১, যোগফল ১।

🧮 **গণিত:**
$$\pi_k(\mathbf{x}) = \frac{\exp(f_k(\mathbf{x}))}{\sum_{j=1}^g \exp(f_j(\mathbf{x}))}$$

বৈশিষ্ট্য:
- প্রতিটা $\pi_k \in [0,1]$, আর $\sum_k \pi_k = 1$।
- $g = 2$ হলে softmax → logistic-এ পরিণত হয় (**generalization**)।
- $\arg\max$-এর "নরম" রূপ: **soft**max হারানো দলগুলোর তথ্যও **reversible** ভাবে রাখে।

🧠 **টোটকা:** "Soft = নরম argmax।" Hard argmax শুধু winner বলে; softmax সবার ভোট-শতাংশ বলে। 🗳️
**Recipe: "exp করো → যোগ দিয়ে ভাগ করো।"**

---

# ১০. Generative vs Discriminative

🔊 **উচ্চারণ:** *জেনারেটিভ* (JEN-er-uh-tiv), *ডিসক্রিমিনেটিভ* (dis-KRIM-in-uh-tiv)

👶 **গল্প — দুই গোয়েন্দা:**
- **Generative** গোয়েন্দা প্রতিটা দল কেমন দেখতে শেখে ("বিড়াল সাধারণত এমন, কুকুর এমন"), তারপর Bayes দিয়ে উল্টো হিসাব করে।
- **Discriminative** গোয়েন্দা শুধু **সীমান্ত** শেখে ("এই রেখার বাঁয়ে বিড়াল, ডানে কুকুর") — দল কেমন দেখতে তা নিয়ে মাথা ঘামায় না।

🧮 **Generative — Bayes' theorem:**
$$\pi_k(\mathbf{x}) = \mathbb{P}(y=k \mid \mathbf{x}) = \frac{\mathbb{P}(\mathbf{x} \mid y=k)\,\mathbb{P}(y=k)}{\mathbb{P}(\mathbf{x})} \propto \mathbb{P}(\mathbf{x} \mid y=k)\,\pi_k$$

🔊 $\propto$ = *"প্রপোরশনাল টু"* = "সমানুপাতিক"।

| Model | ধরন | Boundary | Assumption |
|-------|-----|----------|------------|
| **LDA** | Generative, Linear | সরল | Gaussian, **সমান covariance** Σ |
| **QDA** | Generative, Non-linear | quadratic | Gaussian, **আলাদা covariance** Σₖ |
| **Naive Bayes** | Generative | Non-linear | feature গুলো class দেওয়া থাকলে **স্বাধীন** |
| **Logistic Regression** | **Discriminative**, Linear | সরল | — (সরাসরি boundary শেখে) |

🧮 **Naive Bayes assumption:**
$$\mathbb{P}(\mathbf{x} \mid y=k) = \prod_{j=1}^p \mathbb{P}(x_j \mid y=k)$$

🧠 **মনে রাখার টোটকা:**
- **LDA = সমান (eQuaL... না, L = Linear = equaL covariance)।** QDA = **Q**uadratic = আলাদা covariance। "L সোজা, Q বাঁকা।"
- **Naive** Bayes "naive" (বোকা) কারণ ধরে নেয় feature গুলো একে অপরের থেকে স্বাধীন — যা প্রায়ই বাস্তবে মিথ্যা, তবু কাজ করে।
- **Logistic Regression = Discriminative** (নাম "regression" হলেও এটা classification, আর generative নয়!) — সবচেয়ে বড় ফাঁদ।

---

# ১১. মনে রাখার Master Table

| বিষয় | সঠিক | ফাঁদ (ভুল) |
|------|------|------------|
| Model output | **score** ($f:\mathcal{X}\to\mathbb{R}^g$) | "সরাসরি class" |
| Binary scoring threshold | $c = 0$ ($\operatorname{sgn}$) | "0.5" |
| Binary probabilistic threshold | $c = 0.5$ | "0" |
| Class → Score | **অসম্ভব** | "thresholding দিয়ে করা যায়" |
| $|f(\mathbf{x})|$ | **confidence** | "probability" |
| Linear classifier | non-linear boundary দিতে **পারে** (feature বাড়ালে) | "সবসময় সোজা রেখা" |
| Softmax | logistic-এর **generalization** | "সম্পূর্ণ আলাদা জিনিস" |
| Logistic derivative | $s(1-s)$ | "$s^2$" বা অন্য কিছু |
| LDA | **সমান** covariance | "আলাদা covariance" |
| QDA | **আলাদা** covariance | "সমান covariance" |
| Naive Bayes | **conditional** independence | "unconditional independence" |
| Logistic Regression | **discriminative** | "generative" |

---

# ১১.৫ কেন / কেন নয় (Why & Why Not)

> এই অংশটা শুধু "কী" নয়, "**কেন**" বোঝার জন্য — quiz-এ এই reasoning-ই আসল নম্বর আনে।

✅❌ **Model কেন score দেয়, class নয়?** কারণ (১) continuous score-এ optimization সহজ (derivative নেওয়া যায়), (২) score-এ বেশি তথ্য, (৩) score→class করা যায় কিন্তু class→score **অসম্ভব** (অপরিবর্তনীয়)।

✅❌ **Class → Score কেন অসম্ভব?** একবার "বিড়াল" বলে দিলে কতটা নিশ্চিত ছিলে সেই % আর ফেরত পাওয়া যায় না — তথ্য হারিয়ে গেছে। তাই একমুখী।

✅❌ **Binary scoring-এ threshold কেন $c=0$, probabilistic-এ $c=0.5$?** Scoring-এ সিদ্ধান্ত হয় $\operatorname{sgn}(f)$ দিয়ে — চিহ্ন বদলায় ০-তে। Probabilistic-এ $\pi \in [0,1]$, মাঝবিন্দু ০.৫।

✅❌ **Linear classifier কি সবসময় সোজা boundary দেয়?** ❌ না। Feature engineering (polynomial/basis) করলে **মূল input space-এ non-linear** boundary দিতে পারে। "Linear" মানে weights-এ linear, ছবিতে নয়।

✅❌ **Softmax কি logistic-এর চেয়ে আলাদা জিনিস?** ❌ না — softmax হলো logistic-এর **generalization**; $g=2$ দিলে softmax → logistic-এ পরিণত হয়।

✅❌ **LDA কেন linear, QDA কেন quadratic?** LDA সব class-এ **সমান** covariance ধরে → quadratic term কাটাকাটি হয়ে যায় → linear boundary। QDA **আলাদা** covariance ধরে → quadratic term থেকে যায় → বাঁকা boundary।

✅❌ **Naive Bayes-এ "naive" কেন?** কারণ এটা ধরে নেয় feature গুলো class দেওয়া থাকলে **conditionally independent** — যা প্রায়ই বাস্তবে মিথ্যা, তবু কাজ করে।

✅❌ **Logistic Regression কেন discriminative, generative নয়?** কারণ এটা সরাসরি $\mathbb{P}(y\mid\mathbf{x})$ / discriminant function শেখে — $\mathbb{P}(\mathbf{x}\mid y)$ model করে Bayes দিয়ে উল্টো হিসাব করে না (যা generative করত)।

---

# ১২. অঙ্ক ও সমাধান

> এই অঙ্কগুলো পরীক্ষায় আসার মতো। আগে নিজে চেষ্টা করো, তারপর সমাধান দেখো। ✍️

---

### 🧩 সমস্যা ১ — Logistic sigmoid হিসাব

একটা binary scoring model একটা data point-এ score দিল $f(\mathbf{x}) = 2$। Logistic sigmoid দিয়ে probability $\pi(\mathbf{x})$ বের করো। Threshold $c = 0.5$ হলে prediction কী?

**সমাধান:**
$$\pi(\mathbf{x}) = s(2) = \frac{1}{1 + e^{-2}} = \frac{1}{1 + 0.1353} = \frac{1}{1.1353} \approx 0.88$$
যেহেতু $0.88 \ge 0.5$, prediction $h(\mathbf{x}) = 1$ (positive class)। ✅
আর scoring দৃষ্টিতে: $f = 2 > 0$ → $\operatorname{sgn}(2) = +1$, confidence $|f| = 2$। একই সিদ্ধান্ত।

---

### 🧩 সমস্যা ২ — Softmax হিসাব

৩টা class-এর score: $f_1 = 1,\; f_2 = 2,\; f_3 = 0$। Softmax probability বের করো। Prediction কী?

**সমাধান:** আগে exp:
$$e^1 = 2.718,\quad e^2 = 7.389,\quad e^0 = 1$$
যোগফল: $2.718 + 7.389 + 1 = 11.107$।
$$\pi_1 = \frac{2.718}{11.107} \approx 0.245,\quad \pi_2 = \frac{7.389}{11.107} \approx 0.665,\quad \pi_3 = \frac{1}{11.107} \approx 0.090$$
যোগফল $= 1$ ✅। সবচেয়ে বড় $\pi_2$, তাই $h(\mathbf{x}) = 2$। (সবচেয়ে বড় score-ই জেতে — softmax order বদলায় না।)

---

### 🧩 সমস্যা ৩ — Binary scoring: sign ও confidence

দুই class-এর score: $f_1(\mathbf{x}) = 3,\; f_{-1}(\mathbf{x}) = 5$। Combined function $f = f_1 - f_{-1}$ ব্যবহার করে prediction ও confidence বের করো।

**সমাধান:**
$$f(\mathbf{x}) = f_1 - f_{-1} = 3 - 5 = -2$$
$\operatorname{sgn}(-2) = -1$ → prediction = **−1 (negative class)**। Confidence $= |f| = 2$।
(মনে রাখো: চিহ্ন সিদ্ধান্ত, মাপ জোর।)

---

### 🧩 সমস্যা ৪ — Linear decision boundary

দুই class-এর linear discriminant: $f_1(\mathbf{x}) = 2x_1 + x_2$, $f_2(\mathbf{x}) = x_1 + 3x_2$. Decision boundary-এর সমীকরণ বের করো।

**সমাধান:** Boundary = যেখানে দুই score সমান:
$$f_1 = f_2 \Rightarrow 2x_1 + x_2 = x_1 + 3x_2 \Rightarrow x_1 - 2x_2 = 0$$
অর্থাৎ boundary হলো সরলরেখা $x_1 = 2x_2$ — একটা hyperplane। ✅

---

### 🧩 সমস্যা ৫ — Bayes / Naive Bayes

ধরো $\mathbb{P}(y=\text{spam}) = 0.4$, $\mathbb{P}(y=\text{ham}) = 0.6$। একটা mail-এ কিছু শব্দ দেখে পেলে:
$\mathbb{P}(\mathbf{x}\mid \text{spam}) = 0.05$, $\mathbb{P}(\mathbf{x}\mid \text{ham}) = 0.01$। mail টা spam নাকি ham?

**সমাধান:** Bayes (numerator-ই যথেষ্ট, কারণ denominator দুদিকেই সমান):
$$\text{spam score} \propto 0.05 \times 0.4 = 0.02$$
$$\text{ham score} \propto 0.01 \times 0.6 = 0.006$$
$0.02 > 0.006$ → **spam**। চাইলে normalize:
$$\mathbb{P}(\text{spam}\mid\mathbf{x}) = \frac{0.02}{0.02 + 0.006} = \frac{0.02}{0.026} \approx 0.77$$
অর্থাৎ ~৭৭% নিশ্চিত spam। ✅

---

### 🧩 সমস্যা ৬ — LDA নাকি QDA? (concept)

দুটো ছবি দেখানো হলো: একটায় দুই class-এর decision boundary সোজা রেখা, অন্যটায় বাঁকা (curve)। কোনটা LDA, কোনটা QDA? কারণ?

**সমাধান:**
- সোজা রেখা → **LDA** (সমান covariance ধরে, তাই linear boundary)।
- বাঁকা → **QDA** (প্রতি class-এ আলাদা covariance, তাই quadratic/বাঁকা boundary)।
🧠 "L = Linear (সোজা), Q = Quadratic (বাঁকা)।"

---

### 🧩 সমস্যা ৭ — Logistic derivative প্রয়োগ

$s(t) = \frac{1}{1+e^{-t}}$। দেখাও যে $s(0) = 0.5$ এবং ওই বিন্দুতে slope (derivative) = $0.25$।

**সমাধান:**
$$s(0) = \frac{1}{1 + e^{0}} = \frac{1}{1+1} = 0.5$$
Derivative: $s'(t) = s(t)(1 - s(t))$, তাই
$$s'(0) = 0.5 \times (1 - 0.5) = 0.5 \times 0.5 = 0.25$$
✅ (logistic curve-এর সবচেয়ে খাড়া জায়গা মাঝখানে, $t=0$ এ।)

---

> 🎓 **শেষ কথা, বন্ধু:**
> Classification মানে "কোনটা?" বাছা। Model **score** দেয় → দরকারে sigmoid/softmax দিয়ে **probability** বানাও → threshold দিয়ে **class** ঠিক করো। Boundary মানে যেখানে score সমান। আর সবচেয়ে বড় চারটা ফাঁদ মনে রাখো:
> **(১)** Class → Score অসম্ভব, **(২)** LDA সমান / QDA আলাদা covariance, **(৩)** Naive Bayes = conditional independence, **(৪)** Logistic Regression = discriminative।
>
> এই চারটা ধরলে chapter 2 তোমার হাতের মুঠোয়। 🤜🔥
