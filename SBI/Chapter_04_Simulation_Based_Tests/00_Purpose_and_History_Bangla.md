# Chapter 4 — এই chapter এভাবে কেন সাজানো? উদ্দেশ্য + গল্প

> 👉 **lesson-এর আগে এটা পড়ো।** Chapter 1–3 ছিল "যন্ত্রপাতি বানানো" (draw বানাও, distribution থেকে sample নাও)। Chapter 4-এ প্রথমবার সেই যন্ত্র দিয়ে **আসল বৈজ্ঞানিক প্রশ্নের উত্তর** দেওয়া শুরু — "আমার অনুমানটা কি সত্যি, নাকি কাকতালীয়?"

---

## ১. এই chapter কোন প্রশ্নের উত্তর দেয়?

বিজ্ঞানে সবচেয়ে সাধারণ প্রশ্ন: *"আমি data-তে একটা pattern দেখলাম — এটা কি আসল effect, নাকি শুধু random চান্স?"* উদাহরণ: "ছাত্ররা কি গড়ে বেশি বুদ্ধিমান?", "ওষুধটা কি সত্যিই কাজ করে?"

ঐতিহ্যবাহী পরিসংখ্যানে এর উত্তর দিতে জটিল গাণিতিক সূত্র (t-test, z-test ইত্যাদির closed-form) মুখস্থ করতে হয়। কিন্তু Chapter 4-এর মূল বার্তা:

> **"সূত্র মুখস্থ করো না। যদি H₀ (null hypothesis) সত্যি হতো, তাহলে কেমন data আসত — সেটা হাজারবার simulate করো, আর দেখো তোমার আসল observation সেই simulate-করা দুনিয়ায় কতটা অস্বাভাবিক। এটাই p-value, এটাই test।"**

মানে: **কঠিন গণিতের জায়গায় simulation।** এটাই পুরো কোর্সের "simulation-based" নামের প্রথম সরাসরি প্রয়োগ।

### 🎲 analogy
ধরো তুমি সন্দেহ করছ একটা পাশা কারচুপি-করা (বেশি ৬ পড়ে)। গণিতের সূত্রে না গিয়ে: একটা **ন্যায্য** পাশা (এটাই H₀) ১০০০০ বার ছুঁড়ে দেখো কতবার এত ৬ পড়ে। যদি তোমার সন্দেহজনক পাশার ফলাফল ওই ১০০০০ ন্যায্য-পরীক্ষার মধ্যে খুব বিরল (যেমন ২%) হয়, তাহলে বলবে "এটা কাকতালীয় না, পাশাটা সত্যিই কারচুপি।" ওই ২%-ই p-value।

---

## ২. প্রতিটা টপিক **ঠিক এই ক্রমে** কেন? (design-এর যুক্তি)

**ধাপ ১ — Frequentist দৃষ্টিভঙ্গি ও hypothesis (H₀ vs H₁)**
আগে দর্শন: frequentist-এ probability মানে "বারবার করলে কত ভাগ"। তারপর research question-কে দুই hypothesis-এ রূপ দেওয়া — H₁ (যা প্রমাণ করতে চাই) ও H₀ (তার উল্টো, যাকে আমরা বাতিল করার চেষ্টা করি)। ভিত্তি না বুঝলে test অর্থহীন, তাই এটা প্রথম।

**ধাপ ২ — p-value-র সংজ্ঞা (semi-formal)**
test-এর আত্মা। "H₀ সত্যি ধরে নিয়ে, observed test statistic বা তার চেয়ে চরম মান পাওয়ার probability।" এই ধারণা পরিষ্কার না হলে বাকি সব ভুল বোঝা যাবে।

**ধাপ ৩ — Generative model under H₀ + বারবার simulate**
এখানেই simulation ঢোকে: H₀-কে এমন একটা model-এ লেখো যা থেকে data বানানো যায় (generative model `M₀`)। তারপর হাজারবার data simulate করে test statistic-এর **sampling distribution** বানাও। আসল observation-কে এর সাথে তুলনা করে **simulation-based p-value**।

**ধাপ ৪ — উদাহরণে গভীরে: SD ও Skewness**
সহজ ক্ষেত্রে (mean) সূত্র জানা, কিন্তু **skewness**-এর মতো জটিল test statistic-এর কোনো analytic sampling distribution নেই — সেখানেই simulation-এর আসল শক্তি। তাই এই উদাহরণগুলো দিয়ে দেখানো হয় "কখন simulation অপরিহার্য"।

**ধাপ ৫ — Nuisance parameter ও Pivotal statistic**
বাস্তব সমস্যা: H₀ সব parameter ঠিক করে দেয় না (যেমন σ অজানা — এটা "nuisance")। চালাকি: এমন test statistic বানাও (যেমন `Z=(ȳ−μ)/σ̂`) যা nuisance parameter থেকে **স্বাধীন** (pivotal)। এটা না বুঝলে অনেক বাস্তব test ভেঙে পড়ে।

**ধাপ ৬ — Confidence interval (simulation দিয়ে)**
p-value হলো "হ্যাঁ/না", কিন্তু আমরা প্রায়ই একটা **পরিসর** চাই। sampling distribution-এর quantile থেকে simulation-based confidence interval।

**ধাপ ৭ — Model-based test statistic + সীমাবদ্ধতা**
জটিল প্রশ্নে test statistic নিজেই একটা fit-করা model (যেমন logistic regression-এর coefficient)। শেষে সততার সাথে simulation-based test-এর **তিনটা সীমাবদ্ধতা** — কারণ ভালো বিজ্ঞানী টুলের দুর্বলতাও জানে।

### 🪜 এক বাক্যে design
> **দর্শন (H₀/H₁) → p-value কী → H₀ simulate করে sampling distribution → জটিল statistic-এ এটাই একমাত্র পথ → nuisance সামলাও pivotal দিয়ে → interval বানাও → model-based statistic ও সীমাবদ্ধতা।**

---

## ৩. কেন এই chapter গুরুত্বপূর্ণ (পরের সাথে যোগ)

| Chapter 4-এর ধারণা | পরে যেখানে লাগে |
|---|---|
| Generative model `M₀` (data simulate) | পুরো simulation study (ch6), Bayesian model (ch7–9) |
| Sampling distribution (simulate করে) | bootstrap (ch5), frequentist calibration (ch6) |
| p-value, test statistic | যেকোনো hypothesis যাচাই |
| Confidence interval via quantiles | bootstrap CI (ch5), posterior interval-এর সমান্তরাল |
| Model-based test (fit করা statistic) | যেকোনো জটিল model-এ inference |
| Limitations (nuisance, খরচ) | কেন পরে Bayesian/neural পদ্ধতি দরকার তার প্রেরণা |

> এই chapter আসলে পুরো কোর্সের **frequentist অর্ধেকের শিখর** — আর এর সীমাবদ্ধতাগুলোই পরে Bayesian দুনিয়ায় (ch7+) যাওয়ার যুক্তি তৈরি করে।

---

## ৪. একটু ইতিহাস ও গল্প 📜

### 👨‍🌾 Fisher আর চা-পরীক্ষা (p-value-র জন্ম)
আধুনিক hypothesis testing ও p-value-র জনক **Ronald A. Fisher** (১৯২০-৩০-এর দশক)। বিখ্যাত গল্প — **"Lady Tasting Tea"**: এক ভদ্রমহিলা দাবি করলেন তিনি চাখলেই বলতে পারেন চায়ে আগে দুধ ঢালা হয়েছিল না চা। Fisher একটা পরীক্ষা সাজালেন — ৮ কাপ (৪টা দুধ-আগে, ৪টা চা-আগে) এলোমেলো করে দিয়ে দেখলেন তিনি কতটা ঠিক বলেন। তারপর হিসাব করলেন: *যদি তিনি শুধু আন্দাজ করতেন (এটাই H₀), তাহলে এত কাপ ঠিক বলার probability কত?* — এই probability-ই **p-value**। এটাই permutation/randomization test-এরও আদিরূপ। (Chapter 4-এর পুরো দর্শন সরাসরি এখান থেকেই।)

### 🃏 আবার সেই Monte Carlo সংযোগ
Fisher-এর যুগে computer ছিল না, তাই সব হাতে গুনতে হতো (তাই গাণিতিক সূত্রের ওপর নির্ভরতা)। কিন্তু computer আসার পর (Chapter 1-এর Ulam/von Neumann-এর Monte Carlo) ধারণাটা বদলে গেল — **"সূত্র না জানলেও H₀ থেকে হাজারবার simulate করে p-value পেয়ে যাও।"** Chapter 4 আসলে Fisher-এর ১৯২০-এর আইডিয়া + Monte Carlo-র কম্পিউটেশনাল শক্তির মিলন।

### ⚖️ p-value নিয়ে আধুনিক বিতর্ক (viva-তে চমক)
p-value প্রায়ই **ভুল বোঝা হয়।** এটা "H₀ সত্যি হওয়ার probability" **নয়** — বরং "H₀ সত্যি ধরে নিলে এত-চরম data পাওয়ার probability"। এই ভুল বোঝাবুঝি এত বেড়েছিল যে ২০১৬ সালে **American Statistical Association (ASA)** আনুষ্ঠানিক বিবৃতি দিয়ে p-value-র সঠিক ব্যবহার ব্যাখ্যা করে। (teacher এই সূক্ষ্মতা জিজ্ঞেস করতে পারেন — নিচের viva ফাইলে বিস্তারিত আছে।)

---

## ৫. গল্পের শিক্ষা
> **"H₀ যদি সত্যি হতো, তাহলে দুনিয়াটা কেমন দেখাত — সেই কাল্পনিক দুনিয়া simulate করো, আর তোমার আসল observation সেখানে কতটা বিরল দেখো। বিরল হলে → H₀ বাতিল।"** Fisher থেকে আজকের BayesFlow — এই মূল চিন্তাটাই বহন করে চলেছে।

---

## 🎯 viva-তে এই ফাইল থেকে যা কাজে লাগবে
- **"Chapter 4 কেন দরকার?"** → আগের যন্ত্র (draw/sample) দিয়ে প্রথম আসল inference — "effect সত্যি না কাকতালীয়?"
- **"simulation-based test-এর মূল আইডিয়া?"** → H₀-এর generative model থেকে হাজারবার simulate, observed statistic-এর সাথে তুলনা।
- **"p-value-র ইতিহাস?"** → Fisher, Lady Tasting Tea, randomization test, ১৯২০-৩০-এর দশক।
- **"p-value কী নয়?"** → H₀ সত্যি হওয়ার probability নয়; ASA-2016 বিবৃতি।

> এবার মূল পাঠে → `01_Lesson_Bangla.md`।
