# 🍼 ML ও Statistics-এর পরিভাষা — একদম শিশুর মতো করে বাংলায়

> **কাদের জন্য?** তোমার জন্য, যে ASL ক্লাসে *likelihood, variance, distribution, hypothesis, confidence* শব্দগুলো শুনে মাথা নাড়ো অথচ ভেতরে ভেতরে ঘাবড়ে যাও। 😅
> এই ফাইল প্রতিটা গুরুত্বপূর্ণ শব্দ **শূন্য থেকে** শেখাবে, সাথে থাকবে:
> - 🔊 **উচ্চারণ** (English শব্দ + বাংলা উচ্চারণ)
> - 👶 **শিশুর মতো ব্যাখ্যা** (বাস্তব গল্প দিয়ে)
> - 🧮 **আসল গণিত** (কোনো লুকোচুরি নেই)
> - 🎯 **একাধিক উদাহরণ**

> **কীভাবে পড়বে:** তাড়াহুড়ো নয়। এক শব্দ পড়ো, চোখ বন্ধ করো, নিজেকে গল্পটা আবার বলো। তারপর পরেরটায় যাও।

---

# 📚 সূচিপত্র

1. [Probability (সম্ভাবনা)](#১-probability-সম্ভাবনা)
2. [Random Variable](#২-random-variable)
3. [Distribution (বণ্টন)](#৩-distribution-বণ্টন)
4. [Mean / Expectation (গড়)](#৪-mean--expectation-গড়)
5. [Variance (ভেদাঙ্ক)](#৫-variance-ভেদাঙ্ক)
6. [Standard Deviation](#৬-standard-deviation)
7. [Probability vs Likelihood — সবচেয়ে বড় গোলমাল](#৭-probability-vs-likelihood--সবচেয়ে-বড়-গোলমাল)
8. [Likelihood (সম্ভাব্যতা-স্কোর)](#৮-likelihood)
9. [Maximum Likelihood Estimation (MLE)](#৯-maximum-likelihood-estimation-mle)
10. [Hypothesis ও Hypothesis Space](#১০-hypothesis-ও-hypothesis-space)
11. [Parameter](#১১-parameter)
12. [Estimator ও Estimate](#১২-estimator-ও-estimate)
13. [Bias (পক্ষপাত)](#১৩-bias-পক্ষপাত)
14. [Confidence ও Confidence Interval](#১৪-confidence-ও-confidence-interval)
15. [Conditional Probability (শর্তসাপেক্ষ সম্ভাবনা)](#১৫-conditional-probability)
16. [i.i.d.](#১৬-iid)
17. [Loss ও Risk](#১৭-loss-ও-risk)
18. [বোনাস মিনি-অভিধান](#১৮-বোনাস-মিনি-অভিধান)
19. [গ্রিক অক্ষর ও চিহ্ন — কীভাবে বলবে](#১৯-গ্রিক-অক্ষর-ও-চিহ্ন)

---

# ১. Probability (সম্ভাবনা)

🔊 **উচ্চারণ:** *prob-uh-BIL-uh-tee* · বাংলায়: "প্রোব্যাবিলিটি"

👶 **গল্প:** তুমি একটা কয়েন ছুঁড়লে। কতটা নিশ্চিত যে *head* পড়বে? "অর্ধেক-অর্ধেক।" এই "কতটা নিশ্চিত" সংখ্যাটা, **0 (কখনো না)** থেকে **1 (সবসময়)** এর মধ্যে — এটাই probability।

🧮 **গণিত:**
$$0 \le \mathbb{P}(A) \le 1, \qquad \mathbb{P}(\text{নিশ্চিত ঘটনা}) = 1, \quad \mathbb{P}(\text{অসম্ভব ঘটনা}) = 0$$

🎯 **উদাহরণ:**
- ন্যায্য কয়েন: $\mathbb{P}(\text{head}) = \tfrac{1}{2} = 0.5$
- ছক্কায় 4 পড়া: $\mathbb{P}(X=4) = \tfrac{1}{6} \approx 0.167$
- আজ বৃষ্টি (weather app): $\mathbb{P}(\text{বৃষ্টি}) = 0.8$ → খুব সম্ভব।

> 🔑 Probability = "ঘটনা **ঘটার আগে** তার সম্ভাবনা"।

---

# ২. Random Variable

🔊 **উচ্চারণ:** *RAN-duhm VAIR-ee-uh-bul* · বাংলায়: "র‍্যান্ডম ভেরিয়েবল"

👶 **গল্প:** এমন একটা বাক্স যা *আগে থেকে নিশ্চিত না-জানা একটা সংখ্যা* বের করে দেয়। ছক্কা ছুঁড়লে → 1–6 এর *কোনো একটা* পাবে, কিন্তু কোনটা সেটা random। এই সংখ্যা-মেশিনই random variable, সাধারণত $X$ নামে ডাকা হয়।

🧮 **গণিত:** একটা function $X : \text{outcomes} \to \mathbb{R}$। দুই রকম:
- **Discrete** — গোনা যায় এমন মান (ছক্কা: 1,2,3,4,5,6)।
- **Continuous** — যেকোনো বাস্তব মান (একজনের উচ্চতা: 170.4 cm, 170.41 cm, ...)।

🎯 **উদাহরণ:**
- $X$ = ছক্কার সংখ্যা (discrete)।
- $Y$ = আগামীকালের তাপমাত্রা (continuous)।
- $Z$ = আজ তুমি কয়টা email পেলে (discrete)।

---

# ৩. Distribution (বণ্টন)

🔊 **উচ্চারণ:** *dis-truh-BYOO-shun* · বাংলায়: "ডিস্ট্রিবিউশন"

👶 **গল্প:** Distribution হলো সেই **সম্পূর্ণ রেসিপি** যা বলে দেয় random variable *কোন মানগুলো বেশি পছন্দ করে, আর কতটা*। ভাবো টেবিলের উপর ১০০ দানা চিনি ছিটালে — কোথায় জমে উঠছে সেটাই "distribution"।

🧮 **গণিত:** discrete $X$-এর জন্য: সম্ভাবনার একটা তালিকা $\mathbb{P}(X = x)$ যাদের যোগফল ১। Continuous $X$-এর জন্য: একটা **density** $f(x) \ge 0$ যেখানে $\int f(x)\,dx = 1$।

### যে distribution গুলো অবশ্যই জানতে হবে (উচ্চারণ সহ):

| নাম | 🔊 উচ্চারণ | আকৃতি / ব্যবহার | Density বা P |
|------|-----------|------------------|--------------|
| **Bernoulli** | *ber-NOO-lee* (বার্নুলি) | হ্যাঁ/না, কয়েন ছোঁড়া | $\mathbb{P}(X=1)=p,\ \mathbb{P}(X=0)=1-p$ |
| **Uniform** | *YOO-nuh-form* (ইউনিফর্ম) | সব মান সমান সম্ভাব্য | $f(x)=\tfrac{1}{b-a}$ on $[a,b]$ |
| **Gaussian / Normal** | *GOW-see-an* (গাউসিয়ান) | ঘণ্টা-আকৃতি 🔔 | $f(x)=\tfrac{1}{\sqrt{2\pi\sigma^2}}e^{-\frac{(x-\mu)^2}{2\sigma^2}}$ |
| **Laplace** | *luh-PLOSS* (লাপ্লাস) | ধারালো চূড়া, মোটা লেজ | $f(x)=\tfrac{1}{2b}e^{-\frac{|x-\mu|}{b}}$ |

🎯 **উদাহরণ:**
- প্রাপ্তবয়স্কদের উচ্চতা → **Gaussian** (বেশিরভাগ গড়ের কাছে, খুব কম খুব লম্বা/খাটো)।
- একটা কয়েন ছোঁড়া → **Bernoulli**।
- লটারি নম্বর 1–100 → **Uniform**।

> 🔑 ASL-এ: **Gaussian noise ⟺ L2 loss**, **Laplace noise ⟺ L1 loss**। এই জোড়া মুখস্থ করো!

---

# ৪. Mean / Expectation (গড়)

🔊 **উচ্চারণ:** *meen* / *ek-spek-TAY-shun* · বাংলায়: "মিন / এক্সপেক্টেশন"
🔊 চিহ্ন $\mathbb{E}[X]$ পড়বে **"the expected value of X"** বা "E of X" (ই অফ এক্স)।

👶 **গল্প:** **ভারসাম্য বিন্দু** — distribution-টা যদি ওজনের একটা সিসঅ (seesaw) হতো, mean হলো সেই জায়গা যেখানে আঙুল রাখলে সমান থাকে। এটা "অসীমবার করলে যে গড় পেতে" তা।

🧮 **গণিত:**
- Sample (data থেকে): $\;\bar{x} = \dfrac{1}{n}\sum_{i=1}^n x_i$
- Discrete (আসল): $\;\mathbb{E}[X] = \sum_j x_j\,\mathbb{P}(X=x_j)$
- Continuous (আসল): $\;\mathbb{E}[X] = \int x\, f(x)\, dx$

🎯 **উদাহরণ:**
- ছক্কা: $\mathbb{E}[X] = \tfrac{1+2+3+4+5+6}{6} = 3.5$ (তুমি কখনো 3.5 *ফেলতে* পারবে না — mean সম্ভব মান হতেই হবে এমন নয়!)।
- পরীক্ষার নম্বর {70, 80, 90}: $\bar{x} = 80$।
- কয়েন (head=1, tail=0), ন্যায্য: $\mathbb{E}[X] = 0.5$।

> 🔑 যত বেশি data জোগাড় করবে, **sample mean $\bar x$ → আসল mean $\mathbb{E}[X]$**-এর দিকে যায় (Law of Large Numbers)।

---

# ৫. Variance (ভেদাঙ্ক)

🔊 **উচ্চারণ:** *VAIR-ee-unss* · বাংলায়: "ভেরিয়েন্স"
🔊 চিহ্ন: $\operatorname{Var}(X)$ বা $\sigma^2$ (পড়বে **"sigma squared"**, সিগমা স্কয়ার্ড)।

👶 **গল্প:** "সংখ্যাগুলো কতটা **ছড়ানো / এলোমেলো**?" দুই ক্লাসেরই গড় ৮০ নম্বর। ক্লাস A: সবাই পেল ৭৯–৮১ (টানটান, ছোট variance)। ক্লাস B: অর্ধেক ৬০, অর্ধেক ১০০ (বুনো, বড় variance)। একই mean, কিন্তু *ছড়ানো* সম্পূর্ণ আলাদা।

🧮 **গণিত:** mean থেকে **বর্গ-দূরত্বের** গড়:
$$\operatorname{Var}(X) = \mathbb{E}\big[(X - \mathbb{E}[X])^2\big] = \mathbb{E}[X^2] - (\mathbb{E}[X])^2$$
Sample version: $\;s^2 = \dfrac{1}{n}\sum_{i=1}^n (x_i - \bar{x})^2$

🎯 **উদাহরণ:**
- {80, 80, 80}: variance = 0 (কোনো ছড়ানো নেই)।
- {60, 80, 100}: mean 80, variance $= \tfrac{(-20)^2+0^2+20^2}{3} = \tfrac{800}{3} \approx 266.7$।
- বর্গ কেন? যাতে ধনাত্মক ও ঋণাত্মক ফাঁক একে অপরকে কাটতে না পারে, আর বড় ফাঁক বেশি শাস্তি পায়।

> 🔑 ASL risk minimization-এ আসা দরকারি সূত্র: $\;\mathbb{E}[(y-c)^2] = \operatorname{Var}(y) + (\mathbb{E}[y]-c)^2$, সবচেয়ে ছোট হয় যখন $c = \mathbb{E}[y]$।

---

# ৬. Standard Deviation

🔊 **উচ্চারণ:** *STAN-derd dee-vee-AY-shun* · বাংলায়: "স্ট্যান্ডার্ড ডিভিয়েশন"
🔊 চিহ্ন: $\sigma$ (পড়বে **"sigma"**, সিগমা)।

👶 **গল্প:** Variance থাকে *বর্গ* এককে (যেমন "নম্বর²" — অদ্ভুত!)। তার বর্গমূল নিলে আবার স্বাভাবিক এককে (নম্বর) ফিরে আসো। এটাই standard deviation = "mean থেকে সাধারণ দূরত্ব"।

🧮 **গণিত:**
$$\sigma = \sqrt{\operatorname{Var}(X)}$$

🎯 **উদাহরণ:**
- {60, 80, 100}: variance ≈ 266.7 → $\sigma = \sqrt{266.7} \approx 16.3$ নম্বর।
- উচ্চতায় $\sigma = 7$ cm → বেশিরভাগ মানুষ গড়ের ±7 cm এর মধ্যে।

> 🔑 আংগুলের হিসাব (Gaussian): ~৬৮% data থাকে mean থেকে **1σ**-এর মধ্যে, ~৯৫% থাকে **2σ**-এর মধ্যে।

---

# ৭. Probability vs Likelihood — সবচেয়ে বড় গোলমাল

ছাত্ররা এই জিনিসটাই সবচেয়ে বেশি গুলিয়ে ফেলে। ধীরে পড়ো। 🐢

👶 **গল্প:** একই কয়েন, দুটো আলাদা প্রশ্ন।

- **Probability:** *"আমি জানি কয়েনটা ন্যায্য (p = 0.5)। ৩ বার ছুঁড়ে ৩টাই head পড়ার chance কত?"* → তুমি **model/parameter** ঠিক রাখলে, **data** নিয়ে প্রশ্ন।
- **Likelihood:** *"আমি দেখলাম ৩ বারে ৩টাই head। তাহলে p = 0.5 কতটা বিশ্বাসযোগ্য? নাকি p = 0.9?"* → তুমি **data** ঠিক রাখলে, **parameter** নিয়ে প্রশ্ন।

🧮 **গণিত (একই সূত্র, পার্থক্য "কোনটা স্থির"):**
$$\underbrace{\mathbb{P}(\text{data} \mid \theta)}_{\text{probability: data বদলায়, } \theta \text{ স্থির}} \qquad\qquad \underbrace{\mathcal{L}(\theta) = \mathbb{P}(\text{data} \mid \theta)}_{\text{likelihood: } \theta\text{ বদলায়, data স্থির}}$$

> 🔑 **এক বাক্যে:** *Probability* জানা model থেকে data ভবিষ্যদ্বাণী করে; *Likelihood* জানা data দিয়ে model-কে নম্বর দেয়।

---

# ৮. Likelihood

🔊 **উচ্চারণ:** *LYKE-lee-hood* · বাংলায়: "লাইকলিহুড"
🔊 চিহ্ন: $\mathcal{L}(\theta)$ পড়বে **"likelihood of theta"** (লাইকলিহুড অফ থিটা)।

👶 **গল্প:** তুমি ভেজা মাটি দেখলে 🌧️। *প্রতিটা ব্যাখ্যা কতটা সম্ভব?* "বৃষ্টি হয়েছে" বেশি নম্বর পায়; "ড্রাগন হাঁচি দিয়েছে" কম নম্বর। Likelihood = তুমি যা দেখেছ তার ভিত্তিতে **প্রতিটা সম্ভাব্য ব্যাখ্যা (parameter)-এর জন্য একটা নম্বর**।

🧮 **গণিত:** স্বাধীন data point $x^{(1)}, \dots, x^{(n)}$-এর জন্য likelihood = তাদের probability/density-র **গুণফল**:
$$\mathcal{L}(\theta) = \prod_{i=1}^n p(x^{(i)} \mid \theta)$$
ছোট ছোট সংখ্যার গুণফল কুৎসিত হয় বলে আমরা সাধারণত **log-likelihood** নিই (× কে + এ বদলে দেয়):
$$\ell(\theta) = \log \mathcal{L}(\theta) = \sum_{i=1}^n \log p(x^{(i)} \mid \theta)$$

🎯 **উদাহরণ (কয়েন, তুমি দেখলে H, H, T):**
- $p = 0.5$ হলে: $\mathcal{L} = 0.5 \times 0.5 \times 0.5 = 0.125$
- $p = 0.9$ হলে: $\mathcal{L} = 0.9 \times 0.9 \times 0.1 = 0.081$
- $p = 0.66$ হলে: $\mathcal{L} = 0.66 \times 0.66 \times 0.34 \approx 0.148$ ← **সবচেয়ে বেশি!**
- তাই data "ভোট দেয়" $p \approx 2/3$-এর পক্ষে (যা ৩-এ ২টা head-এর সাথে মেলে 🎉)।

---

# ৯. Maximum Likelihood Estimation (MLE)

🔊 **উচ্চারণ:** *MAX-ih-mum LYKE-lee-hood ess-tih-MAY-shun* · বাংলায়: "ম্যাক্সিমাম লাইকলিহুড এস্টিমেশন"

👶 **গল্প:** *সব* সম্ভাব্য ব্যাখ্যার মধ্যে, **যেটা তোমার দেখা data-কে সবচেয়ে বিশ্বাসযোগ্য করে সেটাই বেছে নাও**। উপরের কয়েনের উদাহরণে সেটা ছিল $p = 2/3$।

🧮 **গণিত:** যে parameter likelihood সর্বোচ্চ করে সেটা বেছে নাও (= **negative** log-likelihood সর্বনিম্ন করা):
$$\hat\theta = \arg\max_\theta \mathcal{L}(\theta) = \arg\min_\theta \Big(-\sum_{i=1}^n \log p(x^{(i)} \mid \theta)\Big)$$

🎯 **উদাহরণ / সংযোগ (ASL সোনা):**
- $n$ বারে $k$টা head → MLE হলো $\hat p = k/n$ (শুধু ভগ্নাংশ!)।
- **Gaussian noise** → MLE = **least squares (L2 loss)**।
- **Laplace noise** → MLE = **L1 loss**।

> 🔑 "$\arg\max$" (পড়ো *"আর্গ ম্যাক্স"*) মানে **"যে input এটাকে সবচেয়ে বড় করে"** — সবচেয়ে বড় মান নয়, বরং *কোথায়* সেটা ঘটে।

---

# ১০. Hypothesis ও Hypothesis Space

🔊 **উচ্চারণ:** *hy-POTH-uh-sis* (বহুবচন *hy-POTH-uh-seez*) · বাংলায়: "হাইপোথিসিস"

👶 **গল্প:**
- একটা **hypothesis** $f$ = একটা *অনুমান* / প্রার্থী model। ("হয়তো দাম = ১০০০ × আয়তন।")
- একটা **hypothesis space** $\mathcal{H}$ = তুমি যত অনুমান বিবেচনা করতে রাজি, তার **পুরো ব্যাগ**। ("সব সরলরেখা।")

শেখা মানে = ব্যাগ $\mathcal{H}$-এ হাত ঢুকিয়ে সেরা অনুমান $\hat f$ বের করা।

🧮 **গণিত:**
$$\hat f \in \mathcal{H}, \qquad \mathcal{H} = \{\, f(\mathbf{x}) = \theta_0 + \theta_1 x \;:\; \theta_0, \theta_1 \in \mathbb{R} \,\} \;\;(\text{যেমন সব রেখা})$$

🎯 **উদাহরণ:**
- $\mathcal{H}$ = সব সরলরেখা → সরল, **underfit** করতে পারে।
- $\mathcal{H}$ = সব আঁকাবাঁকা degree-20 polynomial → নমনীয়, **overfit** করতে পারে।
- বড় ব্যাগ = বেশি ক্ষমতা কিন্তু noise মুখস্থ করার ঝুঁকিও বেশি।

> 🔑 Statistics-এ "hypothesis" মানে *যাচাই করার দাবি*-ও হতে পারে (যেমন "কয়েনটা ন্যায্য") — context বলে দেবে কোন অর্থ।

---

# ১১. Parameter

🔊 **উচ্চারণ:** *puh-RAM-uh-ter* · বাংলায়: "প্যারামিটার"
🔊 চিহ্ন: $\theta$ (পড়বে **"theta"**, থিটা)।

👶 **গল্প:** model বদলাতে যে **নব (knob)** ঘোরাও। একটা রেখা $y = \theta_0 + \theta_1 x$-এর দুটো নব: কোথা থেকে শুরু ($\theta_0$) আর কতটা খাড়া ($\theta_1$)।

🧮 **গণিত:** $\boldsymbol\theta = (\theta_0, \theta_1, \dots)$ — একটা vector-এ জড়ো করা।

🎯 **উদাহরণ:**
- রেখা: parameter = intercept + slope।
- Gaussian: parameter = mean $\mu$ + variance $\sigma^2$।
- কয়েন: parameter = $p$ (head পড়ার chance)।

---

# ১২. Estimator ও Estimate

🔊 **উচ্চারণ:** *ESS-tih-may-ter* / *ESS-tih-mut* · বাংলায়: "এস্টিমেটর / এস্টিমেট"

👶 **গল্প:**
- **Estimator** = data থেকে parameter অনুমানের *রেসিপি/সূত্র* (যেমন "গড় নাও")।
- **Estimate** = data বসিয়ে পাওয়া *আসল সংখ্যা* (যেমন "80.3")।
- **hat** $\hat{\theta}$ (পড়ো *"theta hat"*, থিটা হ্যাট) মানে "আসল $\theta$-এর আমাদের অনুমান"।

🧮 **গণিত:**
$$\hat\mu = \frac{1}{n}\sum_{i=1}^n x_i \quad(\text{estimator}); \qquad \hat\mu = 80.3 \quad(\text{estimate})$$

🎯 **উদাহরণ:** "$\hat p = k/n$" রেসিপিটা estimator; ১০ বারে ৭টা head পেলে estimate $\hat p = 0.7$।

---

# ১৩. Bias (পক্ষপাত)

🔊 **উচ্চারণ:** *BY-uhss* · বাংলায়: "বায়াস"

👶 **গল্প:** একটা ওজন মাপার যন্ত্র যা **সবসময় ২ কেজি বেশি দেখায়** — সেটা *biased*। Statistics-এ bias = "estimator গড়ে সত্য থেকে কতটা দূরে"।

🧮 **গণিত:**
$$\operatorname{Bias}(\hat\theta) = \mathbb{E}[\hat\theta] - \theta$$
Unbiased মানে $\mathbb{E}[\hat\theta] = \theta$ (গড়ে *সঠিক*)।

🎯 **উদাহরণ:**
- Sample mean $\bar x$ আসল mean-এর একটা **unbiased** estimator।
- প্যাটার্ন ধরার জন্য model খুব সরল → **high bias** (underfitting)।
- ⚠️ খেয়াল করো: "bias" শব্দটা model-এর **intercept** $\beta_0$ / ধ্রুবক 1 term-ও বোঝায় — আলাদা অর্থ, একই শব্দ!

---

# ১৪. Confidence ও Confidence Interval

🔊 **উচ্চারণ:** *KON-fih-dunss IN-ter-vul* · বাংলায়: "কনফিডেন্স ইন্টারভাল"

👶 **গল্প:** "গড় উচ্চতা ১৭০ cm" (একটা নড়বড়ে অনুমান) বলার বদলে তুমি বলো "আমি **৯৫% নিশ্চিত** এটা **১৬৮ আর ১৭২ cm**-এর মধ্যে।" এই পরিসরটাই confidence interval — অনিশ্চয়তা নিয়ে সততা।

🧮 **গণিত (mean-এর জন্য ৯৫% CI, মোটামুটি):**
$$\bar{x} \pm 1.96 \cdot \frac{\sigma}{\sqrt{n}}$$
$\dfrac{\sigma}{\sqrt{n}}$ অংশটা **standard error** — এটা **$n$ বাড়লে ছোট হয়** (বেশি data → টানটান, বেশি নিশ্চিত পরিসর)।

🎯 **উদাহরণ:**
- $\bar x = 170$, $\sigma = 10$, $n = 100$ → $170 \pm 1.96 \cdot \tfrac{10}{10} = 170 \pm 1.96$ → প্রায় **[168.0, 172.0]**।
- একই কিন্তু $n = 10000$ → পরিসর কমে ±0.196 → **অনেক টানটান**।

> 🔑 **সাবধানে পড়ো:** "৯৫% confidence" মানে *পদ্ধতিটা* বহুবার করলে ৯৫% বার আসল মান ধরে — "এই একটা পরিসরে সত্য থাকার ৯৫% chance" নয়।

---

# ১৫. Conditional Probability

🔊 **উচ্চারণ:** *kun-DISH-un-ul* · বাংলায়: "কন্ডিশনাল প্রবাবিলিটি" (শর্তসাপেক্ষ সম্ভাবনা)
🔊 চিহ্ন: $\mathbb{P}(A \mid B)$ পড়বে **"probability of A given B"** (`|` মানে "given/দেওয়া থাকলে")।

👶 **গল্প:** "বৃষ্টির chance" বনাম "আকাশ ধূসর **থাকলে** বৃষ্টির chance"। $B$ জানলে $A$ নিয়ে তোমার বিশ্বাস বদলে যায়।

🧮 **গণিত:**
$$\mathbb{P}(A \mid B) = \frac{\mathbb{P}(A \cap B)}{\mathbb{P}(B)}$$

🎯 **উদাহরণ:**
- $\mathbb{P}(\text{রোগ}) = 1\%$, কিন্তু $\mathbb{P}(\text{রোগ} \mid \text{পজিটিভ টেস্ট}) = 80\%$ — টেস্ট chance-টাকে *update* করে।
- ASL-এ: $\mathbb{E}[y \mid \mathbf{x}]$ = "$\mathbf{x}$ জানা **থাকলে** প্রত্যাশিত $y$" — regression যা শিখতে চায়!

---

# ১৬. i.i.d.

🔊 **উচ্চারণ:** অক্ষর ধরে ধরে *"eye-eye-dee"* · বাংলায়: "আই-আই-ডি"
(পুরো নাম: **i**ndependent and **i**dentically **d**istributed)

👶 **গল্প:** প্রতিটা data point (১) **স্বাধীনভাবে** তোলা — একটা আরেকটাকে প্রভাবিত করে না, যেমন আলাদা আলাদা কয়েন ছোঁড়া — আর (২) **একই** distribution থেকে — প্রতিবার একই কয়েন।

🧮 **গণিত:** $x^{(1)}, \dots, x^{(n)} \overset{\text{i.i.d.}}{\sim} \mathbb{P}$। এজন্যই likelihood একটা পরিষ্কার গুণফল $\prod_i p(x^{(i)})$ হয়।

🎯 **উদাহরণ:**
- ✅ ন্যায্য কয়েনের ১০০ বার ছোঁড়া = i.i.d.।
- ❌ প্রতিদিনের তাপমাত্রা = i.i.d. **নয়** (আজ গতকালের উপর নির্ভর করে)।

---

# ১৭. Loss ও Risk

🔊 **উচ্চারণ:** *loss* / *risk* · বাংলায়: "লস / রিস্ক"

👶 **গল্প:**
- **Loss** $L$ = **একটা** ভুল prediction-এর শাস্তি। ("তুমি বললে ৯০, সত্যি ছিল ১০০ → আউচ।")
- **Risk** $\mathcal{R}$ = সবকিছুর উপর শাস্তির **গড়**।

🧮 **গণিত:**
- Point-wise loss: $L(y, f(\mathbf{x})) \ge 0$, যেমন squared $\;(y - f(\mathbf{x}))^2$।
- True risk: $\mathcal{R}(f) = \mathbb{E}[L(y, f(\mathbf{x}))]$।
- Empirical risk (data থেকে): $\;\mathcal{R}_{\text{emp}}(f) = \tfrac{1}{n}\sum_{i=1}^n L(y^{(i)}, f(\mathbf{x}^{(i)}))$।

🎯 **উদাহরণ:**
- Squared loss (L2): বড় ভুল *অনেক* বেশি লাগে (outlier-এ স্পর্শকাতর)।
- Absolute loss (L1): ভুল সমানুপাতে লাগে (robust)।
- শেখা = **empirical risk minimize করা**।

---

# ১৮. বোনাস মিনি-অভিধান

| শব্দ | 🔊 উচ্চারণ | এক লাইনে শিশু-অর্থ |
|------|-----------|---------------------|
| **Covariance** | *koh-VAIR-ee-unss* | দুটো জিনিস কি একসাথে ওঠে/নামে? |
| **Correlation** | *kor-uh-LAY-shun* | covariance-কে **[−1, 1]**-এ মাপা |
| **Residual** | *reh-ZID-yoo-ul* | সত্য আর prediction-এর ফাঁক, $y - f(x)$ |
| **Gradient** | *GRAY-dee-unt* | "উপরে ওঠার তির"; আমরা এর *উল্টোদিকে* হাঁটি |
| **Convex** | *KON-veks* | বাটি-আকৃতি 🥣 → একটাই তলা, optimize সহজ |
| **Overfitting** | *OH-ver-fit-ing* | প্যাটার্নের বদলে noise মুখস্থ করা |
| **i.i.d. sample** | *eye-eye-dee* | পরিষ্কার, ন্যায্য data তোলা |
| **Density** | *DEN-suh-tee* | continuous variable-এর curve-এর উচ্চতা, $f(x)$ |
| **Posterior** | *poss-TEER-ee-or* | data দেখার *পরে* আপডেট হওয়া বিশ্বাস |
| **Prior** | *PRY-or* | data দেখার *আগের* বিশ্বাস |

---

# ১৯. গ্রিক অক্ষর ও চিহ্ন — কীভাবে বলবে

| চিহ্ন | 🔊 নাম | সাধারণত বোঝায় |
|--------|--------|----------------|
| $\theta$ | **theta** (থিটা) | parameter |
| $\mu$ | **mu** (মিউ) | mean |
| $\sigma$ | **sigma** (সিগমা) | standard deviation |
| $\sigma^2$ | **sigma squared** | variance |
| $\beta$ | **beta** (বিটা) | regression coefficient |
| $\varepsilon$ | **epsilon** (এপসিলন) | noise / error |
| $\lambda$ | **lambda** (ল্যাম্বডা) | regularization strength |
| $\alpha$ | **alpha** (আলফা) | learning rate / level |
| $\hat{\theta}$ | **theta hat** | $\theta$-এর *estimate* |
| $\bar{x}$ | **x bar** | sample mean |
| $\mathbb{E}[X]$ | **E of X / expected value** | mean |
| $\mathcal{L}$ | **likelihood** (বা loss $L$) | context-নির্ভর! |
| $\mathbb{P}$ | **P / probability** | probability |
| $\sum$ | **sum / sigma-sum** | যোগ করা |
| $\prod$ | **product / pi-product** | গুণ করা |
| $\int$ | **integral** | continuous "যোগ" |
| $\nabla$ | **nabla / del / gradient** | derivative-এর vector |
| $\sim$ | **"distributed as"** | $X \sim \mathcal{N}$ = "X Normal মেনে চলে" |
| $\mid$ | **"given"** | conditioning |
| $\propto$ | **"proportional to"** | ধ্রুবক বাদে সমান |
| $\arg\max$ | **"arg max"** | যে input সর্বোচ্চ করে |

---

> 🎓 **শেষ শিশু-উপদেশ:** যখনই কোনো ভয়ংকর সূত্র আসে, তিনটা প্রশ্ন করো:
> ১) *কোন অক্ষরটা data, আর কোনটা নব (parameter)?*
> ২) *এটা কি কেন্দ্র মাপছে (mean), ছড়ানো মাপছে (variance), নাকি বিশ্বাস মাপছে (probability/likelihood)?*
> ৩) *আমাকে কি data ভবিষ্যদ্বাণী করতে বলা হচ্ছে (probability), নাকি model-কে নম্বর দিতে (likelihood)?*
>
> এই তিনটার উত্তর দিতে পারলে ASL-এর ৯০% ভয় চলে যাবে। তুমি পারবে। 🍼💪
