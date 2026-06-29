# Chapter 1 — Introduction to Simulations (বাংলায় পুরো পাঠ)

> এই chapter-টা পুরো কোর্সের **ভিত্তি**। এখানে আমরা শিখবো: simulation মানে কী, draw দিয়ে কীভাবে যেকোনো হিসাব (expectation) আনুমানিক করা যায়, সেই আনুমানিকের ভুল কতটুকু (MCSE), আর একটা random variable-কে transform করলে তার density কীভাবে বদলায় (change of variables + Jacobian)। প্রতিটা ধারণা গল্প দিয়ে বুঝবো, একটাও বাদ দেব না।

---

## 🎯 শুরুর আগে: এক লাইনে পুরো chapter

> **"যে integral হাতে কষা যায় না, সেটা আমরা random draw গুনে গুনে আনুমানিক করি — এটাই simulation/Monte Carlo। আর draw-কে transform করলে নতুন distribution পাওয়া যায়।"**

---

## ১. "Random" Sampling বা Simulation আসলে কী?

### গল্প দিয়ে শুরু
ধরো তোমার কাছে একটা distribution আছে — যেমন standard normal `N(0,1)` (সেই বিখ্যাত ঘণ্টা-আকৃতির bell curve)। তুমি চাও এই distribution থেকে কিছু "draw" বা "sample" নিতে — মানে এমন কিছু সংখ্যা, যেগুলোর histogram যদি আঁকো, সেটা যেন ওই bell curve-এর আকার নেয়।

আনুষ্ঠানিক সংজ্ঞা (slide থেকে):
> আমরা `S`টা মান `x⁽¹⁾, x⁽²⁾, …, x⁽ˢ⁾` এমনভাবে বের করি, যাতে তাদের histogram, target distribution-এর density-র সাথে **মিলে যায়** — এবং draw-এর সংখ্যা `S` যত অসীমের দিকে যায়, মিলটা তত নিখুঁত হয়।

এই বের করা মানগুলোকে বলে **draws** বা **samples**।

### 🍪 কুকি-জারের analogy
একটা বড় কাচের জারে হাজার হাজার রঙিন বল আছে — কিছু লাল, কিছু নীল, কিছু সবুজ। জারের ভেতরের আসল অনুপাত = **distribution** (যেটা তুমি সরাসরি দেখতে পাও না)। তুমি চোখ বন্ধ করে একটা একটা করে বল তুলছো — প্রতিটা তোলা বল = একটা **draw**। যত বেশি বল তুলবে, তোমার হাতে জমা বলগুলোর রঙের অনুপাত তত বেশি আসল অনুপাতের কাছাকাছি যাবে। এটাই sampling।

### একটা খুব গুরুত্বপূর্ণ সত্য: computer আসলে "random" না!
- Computer হলো **deterministic** যন্ত্র — একই input দিলে সবসময় একই output দেয়। তাহলে সে "random" সংখ্যা বানায় কীভাবে?
- উত্তর: সে আসলে **pseudo-random** ("ছদ্ম-random") সংখ্যা বানায়, একটা চালাক algorithm (pseudo-random number generator, PRNG) দিয়ে। সংখ্যাগুলো দেখতে random লাগে, কিন্তু আসলে একটা নিয়মে বানানো। (কীভাবে — সেটাই Chapter 2।)
- **Seed:** যেহেতু এটা একটা নিয়মে বানানো, তুমি যদি একই শুরুর সংখ্যা (seed) দাও, তাহলে **হুবহু একই** "random" সংখ্যাগুলো আবার পাবে। তাই simulation **reproducible** (পুনরায় হুবহু চালানো যায়) — গবেষণায় এটা সোনার মতো দামি, কারণ অন্য কেউ তোমার ফলাফল যাচাই করতে পারে।

> 💡 **Viva-corner:** "একটা simulation reproducible করো কীভাবে?" → উত্তর: **seed fix করে** (যেমন Python-এ `np.random.seed(42)`)। এক বাক্যে এই উত্তরটা মুখস্থ রাখো।

---

## ২. Draw বাড়লে histogram কীভাবে density-র দিকে যায় (Law of Large Numbers-এর ছবি)

Slide-এ একটা সুন্দর জিনিস দেখানো হয়েছে: standard normal থেকে draw নিয়ে histogram আঁকা হয়েছে ভিন্ন ভিন্ন সংখ্যায় —

- **১০ draw:** histogram এবড়োখেবড়ো, bell curve-এর সাথে মেলে না তেমন। (কম তথ্য → বেশি noise)
- **১০০ draw:** আকৃতি একটু bell-এর মতো হচ্ছে।
- **১০০০ draw:** বেশ ভালো মিলছে।
- **১০০০০ draw:** প্রায় নিখুঁতভাবে bell curve-কে ঢেকে ফেলছে।

### 🎯 এর শিক্ষা
draw যত বাড়ে, তোমার আনুমানিক (histogram/estimate) তত আসল density-র কাছে যায়। এটাই **Law of Large Numbers**-এর চাক্ষুষ রূপ। কিন্তু খেয়াল করো — এটা **ধীরে** উন্নত হয়: ১০ থেকে ১০০-তে অনেক লাফ, কিন্তু ১০০০ থেকে ১০০০০-এ improvement তুলনায় কম। ঠিক কত ধীরে? সেটা আসছে MCSE-তে (`1/√S` হারে)।

---

## ৩. Draw দিয়ে Expectation (গড়/প্রত্যাশা) আনুমানিক করা — Monte Carlo-র হৃদয়

### কেন আমরা draw চাই? — আসল উদ্দেশ্য
Slide-এ একটা চমৎকার লাইন আছে:
> **প্রায় প্রতিটা গুরুত্বপূর্ণ রাশি (quantity of interest) আসলে একটা expectation।**

গণিতে expectation এমন দেখতে:

$$\mathbb{E}_x[f(x)] = \int f(x)\, p(x)\, dx$$

এখানে `p(x)` হলো density, আর `f(x)` হলো যে জিনিসটার গড় তুমি চাও।

**উদাহরণ দিয়ে বুঝি `f` কী:**
- গড় (mean) চাইলে: `f(x) = x` → `E[x]`
- variance-এর অংশ চাইলে: `f(x) = x²` → `E[x²]`
- কোনো ঘটনার probability চাইলে: `f(x) = 1` যদি শর্ত মেটে, নাহলে `0` (indicator function) → `E[f(x)] = P(শর্ত)`
- যেমন `P(X > 2)` = `E[ 𝟙(X>2) ]`

মানে — mean, variance, probability, tail — সব আসলে একই জিনিস: **কোনো না কোনো `f`-এর expectation**।

### সমস্যাটা কোথায়?
ওই integral `∫ f(x) p(x) dx` অনেক সময় **হাতে কষা অসম্ভব** (no closed-form)। বিশেষ করে high-dimensional হলে। তাহলে?

### Monte Carlo সমাধান (অসাধারণ সহজ আইডিয়া)
distribution `p(x)` থেকে `S`টা independent draw `{x⁽ˢ⁾}` নাও, তারপর শুধু **গড় করো**:

$$\frac{1}{S}\sum_{s=1}^{S} f(x^{(s)}) \;\approx\; \mathbb{E}_x[f(x)]$$

মানে: **কঠিন integral → সহজ যোগফল।** integral-এর জায়গায় শুধু কিছু সংখ্যা বসিয়ে গড় নিলেই হলো!

### 🍲 রান্নার লবণ-চাখা analogy
পুরো এক হাঁড়ি তরকারির গড় লবণ মাপতে তুমি কি পুরোটা খেয়ে ফেলবে? না। তুমি ভালো করে নাড়ো (mixing — যাতে sample representative হয়), তারপর এক চামচ তুলে চাখো। এক চামচ = কয়েকটা draw। একবার চাখলে আন্দাজ মোটামুটি; কয়েকবার বিভিন্ন জায়গা থেকে চাখলে আন্দাজ আরও পাকা। পুরো হাঁড়ি (পুরো integral) না খেয়েও তুমি গড়টা পেয়ে গেলে — এটাই Monte Carlo।

---

## ৪. কতটা নিশ্চিত হওয়া যায়? — Monte Carlo Standard Error (MCSE)

### সমস্যা: আমার আন্দাজটা কি একদম ঠিক?
Monte Carlo estimate তো একটা **আনুমানিক** — random draw দিয়ে বানানো, তাই এতে একটু **ভুল/অনিশ্চয়তা** থাকবেই। প্রশ্ন: কতটুকু?

Slide বলছে — যথেষ্ট বড় `S`-এর জন্য, তোমার estimate প্রায় normal distribution মেনে চলে (Central Limit Theorem-এর কারণে):

$$\frac{1}{S}\sum_{s=1}^{S} f(x^{(s)}) \;\sim\; \text{Normal}\!\left(\mathbb{E}_x[f(x)],\; \sqrt{\tfrac{\text{Var}_x(f(x))}{S}}\right)$$

ওই standard deviation-টাকেই বলে **Monte Carlo Standard Error (MCSE)**:

$$\boxed{\text{MCSE} = \sqrt{\frac{\text{Var}_x(f(x))}{S}}}$$

### 🔑 সবচেয়ে গুরুত্বপূর্ণ অন্তর্দৃষ্টি: `1/√S` সম্পর্ক
লক্ষ্য করো denominator-এ `√S`। মানে:
- MCSE ∝ `1 / √S`
- error **অর্ধেক** করতে চাইলে draw **৪ গুণ** বাড়াতে হবে (কারণ `√4 = 2`)।
- error **১০ ভাগের ১ ভাগ** করতে চাইলে draw **১০০ গুণ** লাগবে।

> 💡 এজন্যই উপরের histogram-গুলোতে ১০০০ → ১০০০০ এ improvement কম মনে হয়েছিল — কারণ error কমে `1/√S` হারে, যা ধীর। এটা Monte Carlo-র "দাম": বেশি নিখুঁততার জন্য অনেক বেশি draw লাগে।

> 💡 **Viva-corner:** "MCSE আর sample size-এর সম্পর্ক?" → **MCSE কমে `1/√S` হারে।** practical-এ: log(MCSE) vs log(S) আঁকলে একটা সরলরেখা পাবে যার slope ≈ **−0.5**। (Exercise 1(e) ঠিক এটাই করায়।)

---

## ৫. Change of Variables — একটা variable transform করলে density কী হয়?

### প্রশ্নটা (slide থেকে)
ধরো তোমার কাছে `x`-এর density `p_x(x)` জানা। এখন `y = exp(x)` বানালে, `y`-এর density `p_y(y)` কী হবে?

### 🤔 কেন এটা সরল ভাগ-গুণ না?
সহজে মনে হতে পারে: "x-এর density জায়গায় y বসিয়ে দিলেই তো হলো!" — কিন্তু **না**। কারণ transform করলে অক্ষটা (axis) টানা-চাপা (stretch/squeeze) হয়। density মানে "প্রতি একক জায়গায় কত probability ঘনত্ব"। জায়গাটাই যদি টানা-চাপা হয়, ঘনত্বও বদলাবে।

### 🪀 রাবার-ব্যান্ড analogy
একটা রাবার-ব্যান্ডের ওপর সমান দূরত্বে দাগ আঁকা। এবার ব্যান্ডটা টানলে — যেখানে বেশি টানলে, দাগগুলো দূরে দূরে (density কম); যেখানে চাপ দিলে, দাগ ঘন (density বেশি)। transform `f` ঠিক এই টানা-চাপাটা করে, আর **Jacobian** `|f'(y)|` মাপে "এই বিন্দুতে কতটা টানা/চাপা হলো"।

### Jacobian Adjustment (গাণিতিক নিয়ম)
ধরা যাক `x = f(y)` continuously differentiable। তখন "integration by substitution" নিয়ম বলে:

$$\int_{\Omega_x} p_x(x)\, dx = \int_{\Omega_y} p_x(f(y))\, |f'(y)|\, dy$$

যেহেতু একটা density-র মোট integral সবসময় `1` (`∫ p_x(x) dx = 1`), তাই `y`-এর density হলো:

$$\boxed{p_y(y) = p_x(f(y))\, |f'(y)|}$$

- এখানে `f = g⁻¹` — মানে সাধারণত আমরা শুরু করি `y = g(x)` থেকে, তারপর উল্টো করে `x = f(y)` বের করি।
- `|f'(y)|`-এ **absolute value** কেন? কারণ density কখনো negative হতে পারে না; transform যদি উল্টো দিকে যায় (decreasing), derivative negative হবে, কিন্তু আমরা শুধু "কতটা টানা-চাপা" সেই মাপটা চাই — দিক না।
- **High dimension-এ:** `|f'(y)|`-এর জায়গায় বসে **Jacobian matrix-এর determinant**-এর absolute value, `|det J_f(y)|`।

### 🌰 উদাহরণ: Lognormal distribution (slide-এর হিসাব)
ধরো `x ~ N(0,1)`, তাই:

$$p_x(x) = \frac{1}{\sqrt{2\pi}} \exp\!\left(-\tfrac{1}{2}x^2\right)$$

এবার `y = g(x) = exp(x)`। উল্টো করে: `x = f(y) = log(y)`।
Derivative: `dx/dy = f'(y) = 1/y`, মানে `dx = (1/y) dy`।

নিয়ম বসিয়ে:

$$p_y(y) = \frac{1}{\sqrt{2\pi}} \exp\!\left(-\tfrac{1}{2}(\log y)^2\right)\cdot \frac{1}{y}$$

এটাই **standard lognormal distribution**-এর density! খেয়াল করো ওই শেষের `1/y`-টাই হলো Jacobian — ওটা বাদ দিলে উত্তর ভুল হতো (পুরো জিনিসটা আর density থাকতো না, integral 1 হতো না)।

> 💡 **Viva-corner:** "Jacobian বাদ দিলে কী হয়?" → density আর valid থাকে না, `∫ p_y(y) dy ≠ 1` হয়ে যায়। Jacobian হলো সেই "সংশোধন কারক" যা probability conserve করে।

---

## ৬. Change of Variables **via random draws** — একটা দারুণ শর্টকাট (এবং Chapter-এর সবচেয়ে practical আইডিয়া)

এখন সবচেয়ে মজার অংশ। উপরের Jacobian-এর হিসাব **density** চাইলে দরকার। কিন্তু যদি তোমার শুধু **draw/sample** দরকার হয় (density-র সূত্র না), তাহলে:

> তুমি যদি `p(x)` থেকে draw `x⁽ˢ⁾` পেয়ে থাকো, তাহলে প্রতিটা draw-কে শুধু `f` দিয়ে রূপান্তর করো — `f(x⁽ˢ⁾)` — আর এগুলোই হবে `p(f(x))` distribution থেকে draw।

$$x^{(s)} \sim p(x) \quad\Longrightarrow\quad f(x^{(s)}) \sim p(f(x))$$

আর সবচেয়ে দারুণ কথা (slide থেকে):
> **No Jacobian adjustment needed!** (draw transform করতে Jacobian লাগে না!)

### 🤯 কেন এত আশ্চর্য সহজ?
- **Density** নিয়ে কাজ করলে: জায়গার টানা-চাপা হিসাব রাখতে হয় → Jacobian লাগে।
- **Draw/sample** নিয়ে কাজ করলে: প্রতিটা বিন্দু এমনিতেই transform হয়ে নতুন জায়গায় চলে যায়, ঘনত্বের হিসাব আপনাআপনি ঠিক হয়ে যায় → কিছু করতে হয় না!

**উদাহরণ (slide-এর ছবি):** বাঁয়ে `x ~ N(0,1)`-এর histogram (symmetric bell)। প্রতিটা draw-এ `exp()` লাগালে ডানে `exp(x)`-এর histogram — যা ডানদিকে লম্বা লেজওয়ালা (right-skewed) lognormal আকৃতি। তুমি density-র সূত্র না জেনেও শুধু draw transform করে lognormal থেকে sample পেয়ে গেলে।

> 💡 **এটাই পুরো "simulation-based" দর্শনের মূল মন্ত্র:** "যা হিসাব করা কঠিন, তা draw দিয়ে সারো।" Density-র জটিল সূত্রে না গিয়ে, draw transform করেই কাজ চলে। এই সরলতাই পরে BayesFlow-এর মতো method-কে সম্ভব করে।

---

## ৭. পুরো Chapter এক নজরে (revision card)

| ধারণা | মূল কথা | সূত্র/মন্ত্র |
|---|---|---|
| Draw/Sample | distribution থেকে তোলা মান | histogram → density (S বাড়লে) |
| Pseudo-random | computer-এর "নকল" random | seed fix → reproducible |
| Monte Carlo | integral → গড় | `E[f] ≈ (1/S)Σ f(xˢ)` |
| MCSE | estimate-এর error | `√(Var/S)`, কমে `1/√S` হারে |
| Change of variables (density) | transform-এ density বদলায় | `p_y(y)=p_x(f(y))·|f'(y)|` |
| Jacobian | টানা-চাপার সংশোধন | high-dim → `|det J|` |
| Transform via draws | draw-এ Jacobian লাগে না | `f(xˢ) ~ p(f(x))` |

---

## ৮. নিজেকে যাচাই করো (পরের ফাইলে যাওয়ার আগে)
নিচের প্রশ্নগুলোর উত্তর মুখে মুখে দিতে পারলে বুঝবে chapter হজম হয়েছে:
1. Computer তো deterministic, তাহলে random সংখ্যা কীভাবে বানায়? seed-এর ভূমিকা কী?
2. `E[f(x)]`-কে draw দিয়ে কীভাবে আনুমানিক করবে? কেন এটা কাজ করে?
3. MCSE-র সূত্র কী, আর error অর্ধেক করতে draw কত গুণ লাগে?
4. `y = g(x)` হলে `p_y(y)` বের করার Jacobian সূত্র কী? `|f'(y)|`-এ absolute value কেন?
5. draw transform করতে Jacobian কেন লাগে না, অথচ density transform করতে লাগে?

> এই প্রশ্নগুলোর বিস্তারিত উত্তর + আরও সম্ভাব্য viva প্রশ্ন আছে `02_Viva_Questions_Bangla.md` ফাইলে। আর hands-on derivation + code আছে `03_Exercise_and_Solution_Bangla.md`-এ।
