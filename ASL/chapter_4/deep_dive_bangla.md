# 🤿 অধ্যায় ৪: Risk Minimization I — Deep Dive
## — বাংলায়, একদম noob-এর জন্য, math + অঙ্কে জোর

---

> **হ্যালো বন্ধু!** 👋 এই অধ্যায়ে loss, risk, gradient descent, MLE — সব **গণিতময়**। তাই এই deep dive-এ আমরা প্রতিটা সূত্র **ভেঙে ভেঙে** বুঝব, উচ্চারণ শিখব, আর অনেক **অঙ্ক কষব**।
>
> প্রতিটা concept-এ: 🔊 উচ্চারণ · 👶 গল্প · 🧮 সূত্র ভেঙে বোঝা · ✅❌ কেন/কেন নয় · 🧠 টোটকা · শেষে ✍️ অঙ্ক+সমাধান।
> (আরও বিস্তারিত গল্প চাইলে দেখো `summary_bangla_by_opus.md`; এই ফাইল math + problem-কেন্দ্রিক।)

---

# 📚 সূচিপত্র
1. [Learning = R + C + O](#১-learning--r--c--o)
2. [Loss — ভুল মাপা](#২-loss--ভুল-মাপা)
3. [Residual ও Pseudo-Residual](#৩-residual-ও-pseudo-residual)
4. [Risk ও Empirical Risk](#৪-risk-ও-empirical-risk)
5. [Regression Loss-গুলো](#৫-regression-loss-গুলো)
6. [Optimal Constant Model](#৬-optimal-constant-model)
7. [Gradient Descent পরিবার](#৭-gradient-descent-পরিবার)
8. [MLE ↔ Loss](#৮-mle--loss)
9. [🧠 Master Table](#৯-master-table)
10. [✍️ অঙ্ক ও সমাধান](#১০-অঙ্ক-ও-সমাধান)

---

# ১. Learning = R + C + O

🔊 **উচ্চারণ:** *রিপ্রেজেন্টেশন + কস্ট + অপটিমাইজেশন*

🧮 **সূত্র:** $\text{Learning} = \text{Representation} + \text{Cost} + \text{Optimization}$

✅❌ **একই linear space দিলে কি একই model আসে?** ❌ না — **cost function** ঠিক করে কোন model আসবে। L2 দিলে এক রেখা, L1 দিলে আরেক।

🧠 **টোটকা:** "R = কী model, C = ভালো-খারাপ, O = কীভাবে খুঁজি।"

---

# ২. Loss — ভুল মাপা

🔊 **উচ্চারণ:** *লস ফাংশন*

🧮 **সূত্র ভেঙে:** $L : \mathcal{Y} \times \mathbb{R}^g \to \mathbb{R}_{\ge 0}$ — দুটো জিনিস (সত্য $y$, prediction $\tilde y$) নিয়ে একটা **non-negative** সংখ্যা দেয়।

দুই শর্ত:
1. **Non-negativity:** $L(y, \tilde y) \ge 0$।
2. **Optimality:** $L(y, \tilde y) = 0 \iff y = \tilde y$।

✅❌ **Loss ঋণাত্মক হতে পারে?** ❌ না — ভুল কখনো negative হয় না। **কেন $=0$ শুধু $y=\tilde y$ তে?** কারণ শূন্য ভুল মানেই নিখুঁত prediction।

🧠 **টোটকা:** "Point-wise = প্রতি point-এ আলাদা ভুল।"

---

# ৩. Residual ও Pseudo-Residual

🔊 **উচ্চারণ:** *রেসিডুয়াল, সুডো-রেসিডুয়াল*

🧮 **সূত্র:**
- Residual: $r = y - f(\mathbf{x})$ (সত্য − prediction)।
- Pseudo-residual: $\tilde r = -\dfrac{\partial L(y, f(\mathbf{x}))}{\partial f}$ (loss-এর $f$-সাপেক্ষে **ঋণাত্মক** derivative)।

দুই ধর্ম:
- **Distance-based:** $L = \psi(r)$ (শুধু residual-এর function), $\psi(0)=0$।
- **Translation-invariant:** $L(y+a, f+a) = L(y, f)$।
- 🔑 **translation-invariant ⟺ distance-based।**

✅❌ **L2-তে residual আর pseudo-residual কি এক?** ✅ **হ্যাঁ** — এজন্যই নাম "pseudo"। দেখো: $-\partial (0.5(y-f)^2)/\partial f = (y-f) = r$।

🧠 **টোটকা:** "Pseudo-residual = loss-plot-এ tangent-এর slope।"

---

# ৪. Risk ও Empirical Risk

🔊 **উচ্চারণ:** *রিস্ক, এম্পিরিক্যাল রিস্ক*

🧮 **সূত্র:**
- True risk: $\mathcal{R}(f) = \mathbb{E}[L(y, f(\mathbf{x}))]$ — প্রত্যাশিত (গড়) loss।
- Empirical risk: $\mathcal{R}_{\text{emp}}(f) = \frac{1}{n}\sum_{i=1}^n L(y^{(i)}, f(\mathbf{x}^{(i)}))$ — data থেকে গড় loss।

✅❌ **True risk সরাসরি minimize করা যায়?** ❌ না — $\mathbb{P}_{xy}$ অজানা। তাই **empirical** risk minimize করি (ERM)।

✅❌ **L2-তে best (theoretical) prediction কী?** ✅ **conditional expectation** $\mathbb{E}[y\mid\mathbf{x}]$। কারণ $\mathbb{E}[(y-c)^2] = \text{Var}(y) + (\mathbb{E}[y]-c)^2$, সবচেয়ে ছোট হয় $c = \mathbb{E}[y]$ তে।

✅❌ **Loss convex + model linear → ?** ✅ প্রতিটা local minimum-ই **global** minimum। Non-convex হলে একাধিক local minimum (খারাপ!)।

🧠 **টোটকা:** "$\frac{1}{n}$ minimizer বদলায় না (শুধু scaling)।"

---

# ৫. Regression Loss-গুলো

🔊 **উচ্চারণ:** *এল-টু, এল-ওয়ান, কোয়ান্টাইল/পিনবল, হিউবার, এপসিলন-ইনসেনসিটিভ, লগ-ব্যারিয়ার*

| Loss | সূত্র | ধর্ম |
|------|------|------|
| **L2** | $(y-f)^2$ | convex, differentiable, **outlier-এ স্পর্শকাতর** |
| **L1** | $\lvert y-f\rvert$ | convex, **robust**, ০-তে differentiable নয় |
| **Quantile** | asymmetric (নিচে দেখো) | L1-এর সম্প্রসারণ |
| **Huber** | L1+L2 মিশ্রণ | **differentiable + robust** |
| **ε-insensitive** | $\epsilon$-এর নিচে ভুল ফ্রি | convex, differentiable নয় |
| **Log-barrier** | $\lvert r\rvert > a$ নিষিদ্ধ | সমাধান নাও থাকতে পারে |

🧮 **Quantile (pinball):** $L = (1-\alpha)(f-y)$ যদি $y<f$, নাহলে $\alpha(y-f)$।

✅❌ **কেন L2 outlier-এ খারাপ?** কারণ residual ২ গুণ হলে loss **৪ গুণ** (বর্গ) → বড় ভুল প্রবলভাবে টানে।
✅❌ **Huber-এর closed-form optimal constant আছে?** ❌ না — numerical লাগে; এটা L1 ও L2 সমাধানের **মাঝে** থাকে।

🧠 **টোটকা:** "L2 = বর্গ (স্পর্শকাতর), L1 = মান (robust), Huber = দুইয়ের সেরা।"

---

# ৬. Optimal Constant Model

👶 **গল্প:** শুধু একটা ধ্রুবক সংখ্যা দিয়ে সব predict করলে, কোন সংখ্যা সেরা? — loss-এর উপর নির্ভর করে।

🧮 **মুখস্থ করো:**
| Loss | Optimal constant |
|------|------------------|
| **L2** | **mean** $\bar y$ |
| **L1** | **median** |
| **Quantile ($\alpha$)** | **$\alpha$-quantile** $Q_\alpha$ |

✅❌ **L2 = mean কেন?** কারণ mean বর্গ-দূরত্বের যোগফল সর্বনিম্ন করে। **L1 = median কেন?** কারণ median পরম-দূরত্বের যোগফল সর্বনিম্ন করে (outlier কম প্রভাব ফেলে)।

🧠 **টোটকা:** "L2→mean, L1→median, Quantile→quantile।"

---

# ৭. Gradient Descent পরিবার

🔊 **উচ্চারণ:** *গ্রেডিয়েন্ট ডিসেন্ট, এসজিডি, মিনি-ব্যাচ*

🧮 **সূত্র:** $\boldsymbol\theta^{[j+1]} = \boldsymbol\theta^{[j]} - \alpha^{[j]} \nabla_{\boldsymbol\theta}\mathcal{R}_{\text{emp}}(\boldsymbol\theta)$

- $\alpha$ = **step-size / learning rate**।
- **GD:** সব $n$টা point দিয়ে gradient।
- **SGD:** **একটা** random point দিয়ে আনুমানিক gradient।
- **Mini-batch:** একটা **subset** $I \subset \{1,\ldots,n\}$ দিয়ে।

✅❌ **negative gradient কেন?** কারণ negative gradient = **steepest descent** (সবচেয়ে খাড়া নিচের দিক)। positive দিকে গেলে error বাড়বে।
✅❌ **SGD কেন ব্যবহার করি?** যখন পুরো gradient হিসাব **ব্যয়বহুল** — SGD সস্তা কিন্তু **noisy**; mini-batch হলো cost আর noise-এর **ভারসাম্য**।

🧠 **টোটকা:** "GD সব, SGD একটা, Mini-batch কিছু।"

---

# ৮. MLE ↔ Loss

🔊 **উচ্চারণ:** *ম্যাক্সিমাম লাইকলিহুড*

🧮 **মূল সংযোগ:** loss সংজ্ঞায়িত করো $L(y, f) := -\log p(y \mid \mathbf{x}, \boldsymbol\theta)$ — তাহলে MLE estimator = loss-minimal estimator।

| Error distribution | সমতুল্য Loss |
|--------------------|--------------|
| **Gaussian** $\mathcal{N}(0,\sigma^2)$ | **L2** |
| **Laplace** | **L1** |

✅❌ **প্রতিটা error distribution → একটা loss?** ✅ হ্যাঁ। **প্রতিটা loss → একটা error distribution?** ❌ **না!** Hinge loss একটা পাল্টা-উদাহরণ। (বড় ফাঁদ!)

🧠 **টোটকা:** "Gaussian↔L2, Laplace↔L1; loss→distribution সবসময় নয়।"

---

# ৯. Master Table

| বিষয় | সঠিক ✅ | ফাঁদ ❌ |
|------|---------|---------|
| Pseudo-residual | **−** derivative | "+ derivative" |
| L2 residual = pseudo? | **হ্যাঁ** | "না" |
| translation-invariant | ⟺ distance-based | — |
| L2 best (theory) | $\mathbb{E}[y\mid\mathbf{x}]$ | "median" |
| L2 optimal constant | **mean** | "median" |
| L1 optimal constant | **median** | "mean" |
| L1 differentiable at 0? | **না** | "হ্যাঁ" |
| convex L + linear f | local = **global** min | "একাধিক global" |
| GD direction | **negative** gradient | "positive" |
| SGD | **একটা** random point | "সব point" |
| Gaussian error | ⟺ **L2** | "L1" |
| every loss → distribution? | **না** (hinge) | "হ্যাঁ" |

---

# ১০. অঙ্ক ও সমাধান

---

### 🧩 সমস্যা ১ — Pseudo-residual (L2)
$L(y,f) = 0.5(y-f)^2$। $\tilde r = -\partial L/\partial f$ বের করো এবং দেখাও এটা residual-এর সমান।

**সমাধান:** $\dfrac{\partial}{\partial f}\,0.5(y-f)^2 = 0.5 \cdot 2(y-f)\cdot(-1) = -(y-f)$।
তাই $\tilde r = -[-(y-f)] = y - f = r$। ✅ pseudo-residual = residual।

---

### 🧩 সমস্যা ২ — Optimal constant (L2 vs L1)
Data: $\{2, 4, 4, 100\}$। (ক) L2-এর optimal constant কত? (খ) L1-এর? (গ) কোনটা outlier-এ ভালো?

**সমাধান:**
- (ক) L2 → mean $= (2+4+4+100)/4 = 110/4 = 27.5$।
- (খ) L1 → median $= (4+4)/2 = 4$।
- (গ) 100 একটা outlier। median (4) data-র সিংহভাগের কাছাকাছি, mean (27.5) outlier-এ টেনে গেছে → **L1 robust**, তাই outlier-এ ভালো। ✅

---

### 🧩 সমস্যা ৩ — Empirical risk (L2)
Predictions $\{3, 5\}$, সত্য $\{4, 1\}$। Squared loss দিয়ে $\bar{\mathcal{R}}_{\text{emp}}$ বের করো।

**সমাধান:** residual: $4-3=1,\; 1-5=-4$। Squared: $1, 16$। যোগ = $17$। গড় = $17/2 = 8.5$।

---

### 🧩 সমস্যা ৪ — Gradient descent এক ধাপ
$\mathcal{R}_{\text{emp}}(\theta) = (\theta - 3)^2$, শুরু $\theta^{[0]} = 0$, learning rate $\alpha = 0.1$। এক ধাপ পরে $\theta$ কত?

**সমাধান:** $\nabla = 2(\theta-3)$। $\theta^{[0]}=0$ এ $\nabla = 2(0-3) = -6$।
$\theta^{[1]} = 0 - 0.1 \times (-6) = 0.6$। (minimum $\theta=3$-এর দিকে এগোচ্ছে ✅।)

---

### 🧩 সমস্যা ৫ — কোন loss? (why & why not)
তোমার data-তে কিছু চরম outlier আছে। L2 নাকি L1 বাছবে? কেন?

**সমাধান:** **L1** (বা Huber)। কারণ L2-তে residual বর্গ হয় → outlier loss-কে প্রবলভাবে টানে, model বিকৃত হয়। L1 পরম-মান নেয় → outlier-এর প্রভাব কম (robust)। ❌ L2 এখানে খারাপ কারণ এটা outlier-এ অতিরিক্ত স্পর্শকাতর। (Huber হলে differentiable + robust দুটোই পাও।)

---

### 🧩 সমস্যা ৬ — MLE ↔ loss
তোমার noise Gaussian ধরা হলো। MLE করলে কার্যত কোন loss minimize হচ্ছে? কেন?

**সমাধান:** **L2 loss**। কারণ Gaussian density-তে $\exp(-\frac{(y-f)^2}{2\sigma^2})$ আছে; এর negative log নিলে পাই $\frac{(y-f)^2}{2\sigma^2}$ + ধ্রুবক → অর্থাৎ squared error (L2)। তাই **Gaussian error ⟺ L2 ⟺ OLS = MLE**। (Laplace হলে হতো L1।)

---

> 🎓 **শেষ কথা:** অধ্যায় ৪ = "ভুল মাপো (loss) → গড় ভুল কমাও (risk) → হাতে বা gradient দিয়ে সমাধান করো।" সোনার সূত্র: **(১)** L2→mean→Gaussian, L1→median→Laplace; **(২)** pseudo-residual = −derivative, L2-তে = residual; **(৩)** convex+linear → global min; **(৪)** GD সব / SGD একটা / mini-batch কিছু; **(৫)** every loss → distribution **নয়** (hinge)। 🍼💪
