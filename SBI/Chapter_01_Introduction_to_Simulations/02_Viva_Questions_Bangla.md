# Chapter 1 — Viva প্রশ্নোত্তর (Introduction to Simulations)

> teacher presentation-এর পরে এমন প্রশ্ন করতে পারে। প্রতিটা উত্তর জোরে জোরে নিজে বলার অভ্যাস করো — বই না দেখে। ⭐ = খুব সম্ভাব্য / গুরুত্বপূর্ণ।

---

### ⭐ Q1. Simulation-based inference-এ "simulation" বলতে আসলে কী বোঝায়?
**উত্তর:** কোনো distribution থেকে অনেকগুলো random draw (sample) তৈরি করা, যাতে ওই draw-গুলো দিয়ে distribution সম্পর্কিত কঠিন হিসাব (যেমন expectation, probability) আনুমানিক করা যায়। মূল দর্শন: **"যা হাতে কষা কঠিন, তা draw গুনে আনুমানিক করো।"**

---

### ⭐ Q2. Computer তো deterministic যন্ত্র, তাহলে "random" সংখ্যা বানায় কীভাবে?
**উত্তর:** সত্যিকারের random না — সে **pseudo-random** সংখ্যা বানায় একটা deterministic algorithm (PRNG) দিয়ে। সংখ্যাগুলো random-এর মতো আচরণ করে (unpredictable মনে হয়), কিন্তু আসলে একটা নিয়মে তৈরি। কীভাবে — সেটা Chapter 2 (LCG, Mersenne Twister ইত্যাদি)।

---

### ⭐ Q3. Seed কী এবং কেন দরকার?
**উত্তর:** Seed হলো PRNG-এর শুরুর মান। একই seed দিলে হুবহু একই "random" sequence পাওয়া যায়। এতে simulation **reproducible** হয় — অন্য কেউ তোমার ফলাফল হুবহু যাচাই করতে পারে। বৈজ্ঞানিক কাজে এটা অপরিহার্য। (Python: `np.random.seed(42)`)

---

### ⭐ Q4. draw-এর সংখ্যা বাড়ালে histogram-এর কী হয়? কোন theorem এর পেছনে?
**উত্তর:** `S` বাড়লে draw-এর histogram ধীরে ধীরে আসল density-র আকৃতি নেয়। ১০ draw-এ এবড়োখেবড়ো, ১০০০০-এ প্রায় নিখুঁত। এর পেছনে **Law of Large Numbers** — sample average আসল expectation-এর দিকে converge করে।

---

### ⭐ Q5. কেন আমরা বলি "প্রায় সব quantity of interest একটা expectation"?
**উত্তর:** কারণ —
- mean = `E[x]` (এখানে `f(x)=x`)
- variance-এর উপাদান = `E[x²]`
- কোনো ঘটনার probability = `E[𝟙(ঘটনা)]` (indicator function-এর expectation, যেমন `P(X>2)=E[𝟙(X>2)]`)

তাই যদি আমরা যেকোনো `f`-এর expectation আনুমানিক করতে পারি, কার্যত সব কিছুই আনুমানিক করতে পারি। সাধারণ রূপ: `E[f(x)] = ∫ f(x) p(x) dx`।

---

### ⭐ Q6. Monte Carlo estimator কী এবং কেন কাজ করে?
**উত্তর:** `p(x)` থেকে `S`টা independent draw `x⁽ˢ⁾` নিয়ে আমরা গড় নিই:
`(1/S) Σ f(x⁽ˢ⁾) ≈ E[f(x)]`।
কাজ করে কারণ Law of Large Numbers অনুযায়ী sample average, true expectation-এ converge করে। সুবিধা: **কঠিন integral → সহজ যোগফল**, এমনকি high-dimension-এও কাজ করে।

---

### ⭐ Q7. MCSE কী? এর সূত্র লেখো।
**উত্তর:** Monte Carlo Standard Error — Monte Carlo estimate-এর অনিশ্চয়তা (standard deviation)। যথেষ্ট বড় `S`-এ estimate প্রায় `Normal(E[f], MCSE)` distribution মেনে চলে (CLT), যেখানে
`MCSE = √( Var_x(f(x)) / S )`।

---

### ⭐⭐ Q8. MCSE আর sample size-এর সম্পর্ক কী? error অর্ধেক করতে কত draw লাগবে?
**উত্তর:** `MCSE ∝ 1/√S`। তাই error **অর্ধেক** করতে draw **৪ গুণ** (`√4=2`), error **১০ ভাগের ১ ভাগ** করতে **১০০ গুণ** draw লাগবে। log-log plot-এ (log MCSE vs log S) slope ≈ **−0.5**। এটাই Monte Carlo-র "মূল্য" — নিখুঁততা ব্যয়বহুল।

---

### Q9. Monte Carlo error কমানোর উপায় কী কী?
**উত্তর:** দুইভাবে: (১) `S` বাড়ানো — কিন্তু ধীর (`1/√S`)। (২) `Var(f(x))` কমানো — variance reduction technique দিয়ে, যেমন **importance sampling**, control variates, antithetic variates। দ্বিতীয়টা প্রায়ই বেশি কার্যকর। (Importance sampling আসছে Chapter 3-এ।)

---

### ⭐ Q10. Change of variables দরকার কেন? `y=g(x)` হলে density কেন সরাসরি বসিয়ে দেওয়া যায় না?
**উত্তর:** কারণ transform অক্ষকে টানা-চাপা (stretch/squeeze) করে। density = প্রতি একক জায়গায় probability ঘনত্ব; জায়গাই বদলালে ঘনত্বও বদলায়। তাই **Jacobian** `|f'(y)|` দিয়ে সেই টানা-চাপার সংশোধন করতে হয়। নাহলে probability conserve হয় না।

---

### ⭐⭐ Q11. Change-of-variable সূত্রটা লেখো এবং প্রতিটা অংশ ব্যাখ্যা করো।
**উত্তর:** `p_y(y) = p_x(f(y)) · |f'(y)|`, যেখানে `x = f(y) = g⁻¹(y)`।
- `p_x(f(y))`: পুরোনো density, নতুন বিন্দুতে মূল্যায়ন।
- `|f'(y)|`: Jacobian — local stretch/squeeze factor।
- absolute value: density negative হতে পারে না, তাই দিক বাদ দিয়ে শুধু মাত্রা।
- High dimension-এ `|f'(y)|` → `|det J_f(y)|` (Jacobian matrix-এর determinant)।

---

### Q12. একটা concrete উদাহরণ দাও যেখানে change of variables লাগে।
**উত্তর:** Lognormal। `x ~ N(0,1)`, `y = exp(x)`। উল্টো: `x = log(y)`, `f'(y)=1/y`। তাই
`p_y(y) = (1/√2π) exp(−½(log y)²) · (1/y)` — standard lognormal। ওই `1/y`-ই Jacobian; বাদ দিলে উত্তর ভুল।

---

### ⭐⭐ Q13. Jacobian বাদ দিলে কী ভুল হয়?
**উত্তর:** ফলাফল আর valid probability density থাকে না — তার মোট integral 1 হয় না (`∫ p_y dy ≠ 1`)। Jacobian হলো সেই সংশোধন কারক যা মোট probability সংরক্ষণ (conserve) করে।

---

### ⭐⭐⭐ Q14. (খুব গুরুত্বপূর্ণ) draw transform করতে Jacobian লাগে না কেন, অথচ density transform করতে লাগে?
**উত্তর:**
- **Density** এ কাজ করলে আমরা "প্রতি একক জায়গায় ঘনত্ব" নিয়ে কাজ করছি; transform জায়গা টানা-চাপা করে, তাই Jacobian দিয়ে ঘনত্ব ঠিক করতে হয়।
- **Draw** এ কাজ করলে প্রতিটা বিন্দু এমনিতেই `f` দিয়ে নতুন জায়গায় চলে যায়; অনেক বিন্দু একসাথে নিলে ঘনত্ব আপনাআপনিই ঠিক হয়ে যায়। তাই `x⁽ˢ⁾ ~ p(x)` হলে `f(x⁽ˢ⁾) ~ p(f(x))` — **কোনো Jacobian ছাড়াই**। এটাই simulation-এর সরলতা: density-র জটিল সূত্র এড়িয়ে শুধু draw transform করে কাজ সারা যায়।

---

### Q15. (চিন্তার প্রশ্ন) আমার কাছে একটা অদ্ভুত distribution `Z = X²` (যেখানে `X~N(0,1)`) থেকে sample দরকার, কিন্তু এর density-র সূত্র মনে নেই। কী করবে?
**উত্তর:** density-র সূত্র (এটা chi-square with 1 df) দরকারই নেই। শুধু `N(0,1)` থেকে অনেক draw নাও, প্রতিটাকে square করো — `x⁽ˢ⁾²` ই হলো target distribution থেকে valid sample (transform via draws, no Jacobian)। এটাই simulation-চিন্তার সৌন্দর্য।

---

### Q16. এই chapter-এর ধারণাগুলো পুরো কোর্সে কীভাবে কাজে লাগে?
**উত্তর:** Monte Carlo + expectation আনুমানিক করা পুরো কোর্সের ভিত — bootstrap (ch5), simulation studies (ch6, ch9), ABC (ch10), সবই draw-নির্ভর। change of variables/Jacobian পরে normalizing flow ও neural posterior estimation (ch12, BayesFlow)-এ কেন্দ্রীয়, কারণ flow আসলে শিখে-নেওয়া invertible transform যার density Jacobian দিয়ে হিসাব হয়।

---

## 🎯 দ্রুত মুখস্থ ("flashcard") — presentation-এর ঠিক আগে চোখ বুলাও
- Simulation = "কঠিন হিসাব draw দিয়ে সারো"
- pseudo-random + seed → reproducible
- `E[f] ≈ (1/S)Σf(xˢ)` (Monte Carlo)
- `MCSE = √(Var/S)`, কমে `1/√S` হারে, log-log slope −0.5
- density transform: `p_y = p_x(f(y))·|f'(y)|`; high-dim → `|det J|`
- draw transform: Jacobian লাগে না — `f(xˢ) ~ p(f(x))`
