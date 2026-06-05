# 🍼 অধ্যায় ০b: Notation & Definitions — চিহ্নের বর্ণমালা (বাংলায়, একদম শিশুর জন্য)

> **কাদের জন্য?** তোমার জন্য, যে ASL-এর slide-এ $\mathcal{X}, \mathbb{P}_{xy}, \pi_k(\mathbf{x}), \hat{\boldsymbol\theta}$ দেখে ভয় পাও। 😅 এই অধ্যায়ে **কোনো নতুন তত্ত্ব নেই** — শুধু কোর্সে ব্যবহৃত প্রতিটা চিহ্ন (notation) কী, **কীভাবে পড়বে (উচ্চারণ)**, কী **অর্থ**, কোথায় **ব্যবহার (use case)**, আর একটা **উদাহরণ**।
>
> **Source:** `docs/chapter_0b_Notation_slides_SoSe2026.pdf` (Bischl, Moosbauer, Groll)।
>
> 💡 প্রতিটা notation এই ৫ ভাবে দেখব: 🔊 **পড়া** · 📖 **অর্থ** · 🎯 **use case** · ✍️ **উদাহরণ**।

---

# 📚 সূচিপত্র
1. [Space ও Variable: 𝒳, 𝒴, x, y](#১-space-ও-variable)
2. [Distribution: ℙxy, p(x,y), p(x|θ)](#২-distribution)
3. [Observation ও Dataset: x⁽ⁱ⁾, 𝒟, train/test](#৩-observation-ও-dataset)
4. [Model: f, h, θ, ℋ](#৪-model)
5. [Residual ও Margin: ε, y·f(x)](#৫-residual-ও-margin)
6. [Probability: πk(x), πk, posterior vs prior](#৬-probability-posterior-vs-prior)
7. [Likelihood ও Hat: ℒ(θ), ℓ(θ), θ̂](#৭-likelihood-ও-hat)
8. [Design Matrix: X, xⱼ, y](#৮-design-matrix)
9. [Binary Label Coding: {0,1} vs {−1,1}](#৯-binary-label-coding)
10. [🧠 এক নজরে চিহ্ন-অভিধান](#১০-এক-নজরে-চিহ্ন-অভিধান)
11. [✍️ অঙ্ক ও সমাধান](#১১-অঙ্ক-ও-সমাধান)

---

# ১. Space ও Variable

### $\mathcal{X}$ — input space
🔊 **পড়া:** "ক্যালিগ্রাফিক X" / "input space"।
📖 **অর্থ:** সব সম্ভাব্য input থাকে যেখানে — $p$-মাত্রার (p-dimensional)। সাধারণত $\mathcal{X} = \mathbb{R}^p$, তবে categorical feature-ও থাকতে পারে।
🎯 **use case:** feature কোন জগতে বাস করে তা বোঝাতে।
✍️ **উদাহরণ:** বাড়ির [আয়তন, ঘর-সংখ্যা] → $\mathcal{X} = \mathbb{R}^2$।

### $\mathcal{Y}$ — target/output space
🔊 **পড়া:** "ক্যালিগ্রাফিক Y" / "target space"।
📖 **অর্থ:** output কোন জগতে থাকে। যেমন $\mathcal{Y}=\mathbb{R}$ (regression), $\{0,1\}$ বা $\{-1,1\}$ (binary), $\{1,\ldots,g\}$ (multiclass)।
✍️ **উদাহরণ:** spam/ham → $\mathcal{Y}=\{0,1\}$।

### $\mathbf{x}$ — feature vector
🔊 **পড়া:** "বোল্ড x" / "feature vector"।
📖 **অর্থ:** একটা input-এর সব feature একসাথে: $\mathbf{x} = (x_1, \ldots, x_p)^\top \in \mathcal{X}$। (বোল্ড = vector; $^\top$ = transpose, "ট্রান্সপোজ"।)
✍️ **উদাহরণ:** $\mathbf{x} = (1200,\ 3)^\top$ (১২০০ বর্গফুট, ৩ ঘর)।

### $y$ — target / label / output / response
🔊 **পড়া:** "y" / "target"।
📖 **অর্থ:** আসল উত্তর, $y \in \mathcal{Y}$। (এর অনেক নাম: target = label = output = response।)
✍️ **উদাহরণ:** ওই বাড়ির দাম $y = 4300$।

🧠 **টোটকা:** "বড় caligraphic = জগত (space), ছোট = একটা মান।"

---

# ২. Distribution

### $\mathbb{P}_{xy}$ — joint probability distribution
🔊 **পড়া:** "P x y" / "joint distribution on 𝒳 × 𝒴"।
📖 **অর্থ:** প্রকৃতির সেই গোপন যন্ত্র যা $(\mathbf{x}, y)$ জোড়া তৈরি করে। সাধারণত **অজানা**।
🎯 **use case:** data কোথা থেকে আসে তার তাত্ত্বিক উৎস।

### $p(\mathbf{x}, y)$ বা $p(\mathbf{x}, y \mid \boldsymbol\theta)$ — joint pdf
🔊 **পড়া:** "p of x comma y" / "joint pdf"।
📖 **অর্থ:** ওই distribution-এর density function।

> ⚠️ **বড় ফাঁদ — `|` চিহ্ন:** এই কোর্স **frequentist** দৃষ্টিতে। `|`-এর পরে parameter থাকলে সেটা **শুধু পড়ার সুবিধা**, Bayesian conditioning **নয়**! অর্থাৎ $p(\mathbf{x}\mid\boldsymbol\theta)$ আসলে $p_{\boldsymbol\theta}(\mathbf{x})$ বা $p(\mathbf{x}; \boldsymbol\theta)$।
> ✍️ মনে রাখো: এখানে `|` মানে "এই $\theta$ দিয়ে নির্ধারিত", "given (শর্ত)" নয়।

🧠 **টোটকা:** "ℙ (blackboard) = distribution; p (ছোট) = density/pdf।"

---

# ৩. Observation ও Dataset

### $(\mathbf{x}^{(i)}, y^{(i)})$ — i-th observation / instance
🔊 **পড়া:** "x super i, y super i" / "i-th observation"।
📖 **অর্থ:** $i$-নম্বর data point (একটা সারি/row)। উপরের কোঠার $(i)$ = **কোন observation**।
✍️ **উদাহরণ:** $(\mathbf{x}^{(3)}, y^{(3)})$ = ৩-নম্বর বাড়ি ও তার দাম।

### $\mathcal{D}$ — dataset
🔊 **পড়া:** "ক্যালিগ্রাফিক D" / "dataset"।
📖 **অর্থ:** সব observation একসাথে: $\mathcal{D} = \{(\mathbf{x}^{(1)}, y^{(1)}), \ldots, (\mathbf{x}^{(n)}, y^{(n)})\}$, $n$টা observation।

### $\mathcal{D}_{\text{train}}, \mathcal{D}_{\text{test}}$ — train ও test data
🔊 **পড়া:** "D train, D test"।
📖 **অর্থ:** শেখার জন্য আর যাচাইয়ের জন্য আলাদা ভাগ। প্রায়ই $\mathcal{D} = \mathcal{D}_{\text{train}} \,\dot\cup\, \mathcal{D}_{\text{test}}$ (disjoint union — কোনো overlap নেই)।
🎯 **use case:** model train data-তে শেখে, test data-তে পরীক্ষা হয়।

🧠 **টোটকা:** "উপরে $(i)$ = কোন point; $n$ = মোট কয়টা।"

---

# ৪. Model

### $f(\mathbf{x})$ বা $f(\mathbf{x}\mid\boldsymbol\theta)$ — prediction function (model)
🔊 **পড়া:** "f of x" / "model"।
📖 **অর্থ:** data থেকে শেখা function, output $\in \mathbb{R}$ বা $\mathbb{R}^g$ (score)। কখনো $\theta$ লেখা বাদ দেওয়া হয়।
✍️ **উদাহরণ:** $f(\mathbf{x}) = 0.5\,x + 2$।

### $h(\mathbf{x})$ — discrete prediction (classification)
🔊 **পড়া:** "h of x"।
📖 **অর্থ:** classification-এ **নির্দিষ্ট class** (discrete) বের করে, $h(\mathbf{x}) \in \mathcal{Y}$।
✍️ **পার্থক্য:** $f$ = score (যেমন 0.87); $h$ = চূড়ান্ত class (যেমন "spam")।

### $\boldsymbol\theta \in \Theta$ — model parameters
🔊 **পড়া:** "theta" (থিটা); $\Theta$ = "বড় theta / parameter space"।
📖 **অর্থ:** model-এর নব (knob)। $\Theta$ = সব সম্ভাব্য $\theta$-এর সেট।
✍️ **উদাহরণ:** রেখা $f = \theta_0 + \theta_1 x$-এ $\boldsymbol\theta = (\theta_0, \theta_1)$।

### $\mathcal{H}$ — hypothesis space
🔊 **পড়া:** "ক্যালিগ্রাফিক H" / "hypothesis space"।
📖 **অর্থ:** $f$ যেখানে বাস করে — সব অনুমোদিত model-এর সেট; $f$-এর রূপ সীমিত করে।

🧠 **টোটকা:** "$f$ = score, $h$ = class, $\theta$ = নব, $\mathcal{H}$ = model-এর ব্যাগ।"

---

# ৫. Residual ও Margin

### $\varepsilon$ — residual (regression)
🔊 **পড়া:** "epsilon" (এপসিলন)।
📖 **অর্থ:** আসল − prediction: $\varepsilon = y - f(\mathbf{x})$, বা $\varepsilon^{(i)} = y^{(i)} - f(\mathbf{x}^{(i)})$।
✍️ **উদাহরণ:** সত্য 10, prediction 7 → $\varepsilon = 3$।

### $y\,f(\mathbf{x})$ — margin (binary classification, 𝒴 = {−1,1})
🔊 **পড়া:** "y times f of x" / "margin"।
📖 **অর্থ:** label আর score-এর গুণফল। **ধনাত্মক margin = সঠিক** শ্রেণিবিভাগ, **ঋণাত্মক = ভুল**।
✍️ **উদাহরণ:** $y=+1$, $f(\mathbf{x})=2$ → margin $=+2$ (সঠিক, আত্মবিশ্বাসী)। $y=+1$, $f=-2$ → margin $=-2$ (ভুল)।

🧠 **টোটকা:** "Regression-এ residual, classification-এ margin।"

---

# ৬. Probability: posterior vs prior

### $\pi_k(\mathbf{x}) = \mathbb{P}(y=k \mid \mathbf{x})$ — posterior probability
🔊 **পড়া:** "pi k of x" (পাই-কে) / "posterior probability for class k"।
📖 **অর্থ:** $\mathbf{x}$ **দেখার পরে** class $k$ হওয়ার সম্ভাবনা। Binary-তে সংক্ষেপে $\pi(\mathbf{x}) = \mathbb{P}(y=1 \mid \mathbf{x})$।
✍️ **উদাহরণ:** এই email দেখে spam হওয়ার সম্ভাবনা $\pi(\mathbf{x}) = 0.9$।

### $\pi_k = \mathbb{P}(y=k)$ — prior probability
🔊 **পড়া:** "pi k" / "prior probability for class k"।
📖 **অর্থ:** $\mathbf{x}$ **দেখার আগে** class $k$ হওয়ার সাধারণ সম্ভাবনা। Binary-তে $\pi = \mathbb{P}(y=1)$।
✍️ **উদাহরণ:** সব email-এর ৪০% spam → $\pi = 0.4$।

> 🔑 **পার্থক্য:** posterior = $\mathbf{x}$ জানার **পরে** (with evidence); prior = জানার **আগে** (background)।

🧠 **টোটকা:** "Posterior-এ $\mathbf{x}$ আছে (after); prior-এ $\mathbf{x}$ নেই (before)।"

---

# ৭. Likelihood ও Hat

### $\mathcal{L}(\boldsymbol\theta)$ ও $\ell(\boldsymbol\theta)$ — likelihood ও log-likelihood
🔊 **পড়া:** "caligraphic L of theta" (likelihood); "small ell of theta" (log-likelihood)।
📖 **অর্থ:** parameter $\theta$ কতটা বিশ্বাসযোগ্য, data দেখে। $\ell = \log\mathcal{L}$ (গুণফলকে যোগফলে বদলায়)।
🎯 **use case:** MLE — যে $\theta$ likelihood সর্বোচ্চ করে।

### হ্যাট $\;\hat{\,}\,$ — learned/estimated
🔊 **পড়া:** "hat" (হ্যাট)। $\hat f$ = "f hat", $\hat{\boldsymbol\theta}$ = "theta hat"।
📖 **অর্থ:** data থেকে **শেখা/অনুমান করা** version। $\hat f, \hat h, \hat\pi_k(\mathbf{x}), \hat\pi(\mathbf{x}), \hat{\boldsymbol\theta}$।
✍️ **পার্থক্য:** $\theta$ = আসল (অজানা) parameter; $\hat\theta$ = data থেকে আমাদের অনুমান।

> 📝 **Remark:** random variable $\mathbf{x}, y$ ছোট হাতের অক্ষরে লেখা হয় (সাধারণ variable-এর মতো) — context থেকে বোঝা যাবে কোনটা random।

🧠 **টোটকা:** "হ্যাট মানেই 'শেখা/অনুমান করা', আসল নয়।"

---

# ৮. Design Matrix

### $\mathbf{X}$ — design matrix
🔊 **পড়া:** "bold X" / "design matrix"।
📖 **অর্থ:** সব observation সারি-সারি সাজানো একটা matrix ($n$ সারি)। দুই রূপ:
- **intercept ছাড়া:** $n \times p$ (শুধু feature)।
- **intercept সহ:** $n \times (p+1)$ — **প্রথম column পুরোটা ১**।

🎯 **use case:** intercept trick — constant-1 feature যোগ করলে $f(\mathbf{x}) = \boldsymbol\theta^\top\mathbf{x}$ লেখা যায়, $f(\mathbf{x}) = \theta_0 + \boldsymbol\theta^\top\mathbf{x}$-এর বদলে (notation সহজ হয়)।

### $\mathbf{x}_j$ — j-th feature vector (column)
🔊 **পড়া:** "x sub j"।
📖 **অর্থ:** $j$-নম্বর **feature** সব observation জুড়ে (একটা column): $\mathbf{x}_j = (x_j^{(1)}, \ldots, x_j^{(n)})^\top$।

### $\mathbf{y}$ — target vector
📖 **অর্থ:** সব target একসাথে: $\mathbf{y} = (y^{(1)}, \ldots, y^{(n)})^\top$।

> 🔑 **মস্ত ফাঁদ — উপরে নাকি নিচে index?**
> - $\mathbf{x}^{(i)}$ (উপরে) = **কোন observation** (সারি/row)।
> - $\mathbf{x}_j$ (নিচে) = **কোন feature** (column)।
> - একসাথে $x_j^{(i)}$ = $i$-তম observation-এর $j$-তম feature (একটা ঘর/cell)।

🧠 **টোটকা:** "উপরে = সারি (which point), নিচে = কলাম (which feature)।"

---

# ৯. Binary Label Coding

👶 **গল্প:** Binary classification-এ দুটো জনপ্রিয় coding আছে — কোনটা ব্যবহার করবে নির্ভর করে model-এর ধরনের উপর।

### $\mathcal{Y} = \{0, 1\}$ — probability/likelihood-ভিত্তিক
📖 এখানে model করে $\pi(\mathbf{x})$ = class 1-এর posterior probability। সিদ্ধান্ত:
$$h(\mathbf{x}) = \mathbb{1}(\pi(\mathbf{x}) \ge 0.5)$$
🔊 $\mathbb{1}(\cdot)$ = "indicator function" — ভেতরের শর্ত সত্য হলে ১, নাহলে ০।
🎯 **use case:** logistic regression, probability-ভিত্তিক model।

### $\mathcal{Y} = \{-1, +1\}$ — geometric/loss-ভিত্তিক
📖 এখানে model করে $f(\mathbf{x})$ = একটা বাস্তব score। সিদ্ধান্ত:
$$h(\mathbf{x}) = \operatorname{sign}(f(\mathbf{x}))$$
আর $|f(\mathbf{x})|$ = predicted class-এর **confidence**।
🎯 **use case:** SVM, margin/loss-ভিত্তিক model।

> 🔑 **মনে রাখো:** **{0,1} → probability** ($\pi$, threshold 0.5); **{−1,+1} → score** ($f$, sign + confidence)।

🧠 **টোটকা:** "Probability হলে {0,1}; geometry/margin হলে {−1,+1}।"

---

# ১০. এক নজরে চিহ্ন-অভিধান

| চিহ্ন | 🔊 পড়া | অর্থ |
|------|--------|------|
| $\mathcal{X}, \mathcal{Y}$ | ক্যালিগ্রাফিক X/Y | input/output space (জগত) |
| $\mathbf{x}$ | বোল্ড x | feature vector |
| $y$ | y | target/label |
| $\mathbb{P}_{xy}$ | P x y | joint distribution |
| $p(\mathbf{x},y)$ | p of x,y | joint pdf |
| $p(\mathbf{x}\mid\boldsymbol\theta)$ | p of x given theta | $\theta$ দিয়ে নির্ধারিত density (**Bayesian নয়!**) |
| $(\mathbf{x}^{(i)},y^{(i)})$ | x super i | i-th observation |
| $\mathcal{D}$ | ক্যালিগ্রাফিক D | dataset |
| $f(\mathbf{x})$ | f of x | model (score) |
| $h(\mathbf{x})$ | h of x | discrete class prediction |
| $\boldsymbol\theta, \Theta$ | theta, বড় theta | parameter, parameter space |
| $\mathcal{H}$ | ক্যালিগ্রাফিক H | hypothesis space |
| $\varepsilon$ | epsilon | residual $y - f(\mathbf{x})$ |
| $y\,f(\mathbf{x})$ | y f of x | margin |
| $\pi_k(\mathbf{x})$ | pi k of x | posterior $\mathbb{P}(y=k\mid\mathbf{x})$ |
| $\pi_k$ | pi k | prior $\mathbb{P}(y=k)$ |
| $\mathcal{L}(\theta), \ell(\theta)$ | L, ell of theta | likelihood, log-likelihood |
| $\hat{\,}$ | hat | শেখা/অনুমান করা |
| $\mathbf{X}$ | বোল্ড X | design matrix |
| $\mathbf{x}_j$ | x sub j | j-th feature (column) |
| $x_j^{(i)}$ | x sub j super i | i-th point-এর j-th feature |
| $\mathbb{1}(\cdot)$ | indicator | শর্ত সত্য→1, নাহলে→0 |
| $\operatorname{sign}(\cdot)$ | sign | চিহ্ন (+/−) |
| $^\top$ | transpose | সারি↔কলাম উল্টানো |

---

# ১১. অঙ্ক ও সমাধান

---

### 🧩 সমস্যা ১ — উপরে নাকি নিচে index?
$x_2^{(5)}$ মানে কী?

**সমাধান:** **৫-নম্বর observation**-এর **২-নম্বর feature**-এর মান। (উপরে = কোন point, নিচে = কোন feature।)

---

### 🧩 সমস্যা ২ — Residual
$y^{(i)} = 25$, $f(\mathbf{x}^{(i)}) = 19$। $\varepsilon^{(i)}$ কত?

**সমাধান:** $\varepsilon^{(i)} = y^{(i)} - f(\mathbf{x}^{(i)}) = 25 - 19 = 6$।

---

### 🧩 সমস্যা ৩ — Margin (binary, {−1,1})
(ক) $y = +1$, $f(\mathbf{x}) = 3$। (খ) $y = -1$, $f(\mathbf{x}) = 2$। margin বের করো ও বলো সঠিক না ভুল।

**সমাধান:**
- (ক) margin $= y f(\mathbf{x}) = (+1)(3) = +3$ → **ধনাত্মক → সঠিক** (আত্মবিশ্বাসী)।
- (খ) margin $= (-1)(2) = -2$ → **ঋণাত্মক → ভুল**।

---

### 🧩 সমস্যা ৪ — posterior vs prior
"সব email-এর ৩০% spam" আর "এই নির্দিষ্ট email spam হওয়ার সম্ভাবনা ৯৫%" — কোনটা prior, কোনটা posterior? notation-এ লেখো।

**সমাধান:**
- ৩০% = **prior** $\pi = \mathbb{P}(y=\text{spam}) = 0.3$ ($\mathbf{x}$ ছাড়া)।
- ৯৫% = **posterior** $\pi(\mathbf{x}) = \mathbb{P}(y=\text{spam}\mid\mathbf{x}) = 0.95$ ($\mathbf{x}$ দেখে)।

---

### 🧩 সমস্যা ৫ — Binary coding ও সিদ্ধান্ত
(ক) $\mathcal{Y}=\{0,1\}$, $\pi(\mathbf{x}) = 0.4$। $h(\mathbf{x})$ কত? (খ) $\mathcal{Y}=\{-1,1\}$, $f(\mathbf{x}) = -1.5$। $h(\mathbf{x})$ ও confidence কত?

**সমাধান:**
- (ক) $h(\mathbf{x}) = \mathbb{1}(0.4 \ge 0.5) = \mathbb{1}(\text{মিথ্যা}) = 0$।
- (খ) $h(\mathbf{x}) = \operatorname{sign}(-1.5) = -1$; confidence $= |f(\mathbf{x})| = 1.5$।

---

### 🧩 সমস্যা ৬ — `|` চিহ্ন (why & why not)
$p(\mathbf{x}\mid\boldsymbol\theta)$-তে `|` কি Bayesian conditioning? ব্যাখ্যা করো।

**সমাধান:** ❌ না। এই কোর্স frequentist; এখানে `|` শুধু **পড়ার সুবিধার** চিহ্ন — মানে "$\theta$ দিয়ে নির্ধারিত density"। ✅ সঠিকভাবে এটাকে $p_{\boldsymbol\theta}(\mathbf{x})$ বা $p(\mathbf{x}; \boldsymbol\theta)$ পড়া উচিত। (Bayesian হলে $\theta$-কেও random ধরা হতো, যা এখানে হচ্ছে না।)

---

### 🧩 সমস্যা ৭ — Intercept trick
$f(\mathbf{x}) = \theta_0 + \theta_1 x_1 + \theta_2 x_2$-কে design matrix trick দিয়ে $\boldsymbol\theta^\top\tilde{\mathbf{x}}$ আকারে লেখো। $\tilde{\mathbf{x}}$ কী?

**সমাধান:** feature-এর সামনে constant-1 যোগ করো: $\tilde{\mathbf{x}} = (1, x_1, x_2)^\top$, $\boldsymbol\theta = (\theta_0, \theta_1, \theta_2)^\top$। তাহলে $f(\mathbf{x}) = \boldsymbol\theta^\top\tilde{\mathbf{x}}$। feature space তখন $(p+1)$-মাত্রার।

---

> 🎓 **শেষ কথা:** অধ্যায় ০b = ASL-এর **বর্ণমালা**। তিনটা সোনা: **(১)** উপরে index = observation, নিচে = feature; **(২)** `|` এখানে Bayesian নয় (frequentist পড়ার সুবিধা); **(৩)** binary-তে {0,1}→probability ($\pi$, 0.5), {−1,+1}→score ($f$, sign+confidence)। এই চিহ্নগুলো চিনলে বাকি সব অধ্যায় পড়া সহজ হয়ে যাবে। 🍼💪
