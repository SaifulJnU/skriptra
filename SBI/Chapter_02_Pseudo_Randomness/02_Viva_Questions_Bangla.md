# Chapter 2 — Viva প্রশ্নোত্তর (Pseudo-randomness)

> ⭐ = খুব সম্ভাব্য / গুরুত্বপূর্ণ। প্রতিটা জোরে জোরে বলে অভ্যাস করো।

---

### ⭐ Q1. "Pseudo-random" মানে কী? কেন computer সত্যিকারের random বানাতে পারে না?
**উত্তর:** Computer deterministic — একই input → একই output। তাই সে সত্যিকারের random না, বরং একটা deterministic algorithm (PRNG) দিয়ে এমন sequence বানায় যা random-এর মতো দেখায় (পরের মান আগেরগুলো থেকে অনুমান করা কঠিন)। একে বলে pseudo-random। এর জন্য চাই **পূর্ণসংখ্যার ওপর একটা chaotic function**।

---

### ⭐ Q2. Modulo operator কী, আর generator-এ এর ভূমিকা কী?
**উত্তর:** `mod_m(b)` = `b`-কে `m` দিয়ে ভাগ করলে ভাগশেষ। এটা ফলাফলকে সবসময় `{0,…,m−1}`-এর ভেতরে আবদ্ধ রাখে (ঘড়ির মতো চক্রাকার)। এই "wrap-around"-ই sequence-কে সীমাবদ্ধ ও chaotic রাখে, তাই প্রায় সব classic generator (LCG, MWC) modulo-র ওপর দাঁড়ানো।

---

### ⭐⭐ Q3. Linear Congruential Generator-এর সূত্র ও প্যারামিটার লেখো।
**উত্তর:** `x_{j+1} = mod_m(a·x_j + c)`। চারটা প্যারামিটার:
- `x₀` = seed (শুরুর মান)
- `a` = multiplier
- `c` = increment
- `m` = modulus (`m > x₀, a, c`)

কাজ: আগের সংখ্যাকে `a` দিয়ে গুণ, `c` যোগ, `m` দিয়ে ভাগশেষ — সেটাই পরের সংখ্যা।

---

### ⭐ Q4. Period মানে কী? কেন বড় period চাই?
**উত্তর:** Period `μ` = কত ঘর পর sequence হুবহু পুনরাবৃত্তি করে (`xⱼ = x_{j+μ}`)। ছোট period মানে generator অল্প কিছু সংখ্যা ঘুরিয়ে-ফিরিয়ে দেয় → simulation-এ নকল pattern ও bias। তাই যত বড় period তত ভালো।

---

### ⭐⭐⭐ Q5. প্রমাণ করো প্রতিটা LCG-এর period `≤ m`। (pigeonhole)
**উত্তর:** প্রতিটা পদ `mod_m(...)`, তাই মাত্র `m`টা সম্ভাব্য মান (`0` থেকে `m−1`)। প্রথম `m+1`টা পদের মধ্যে অন্তত দুটো অবশ্যই সমান হবে (pigeonhole: `m+1` পদ, খোপ `m`টা)। recurrence deterministic হওয়ায় একবার কোনো মান ফিরে এলে পুরো sequence পুনরাবৃত্তি হয়। ∴ `μ ≤ m`। ∎

---

### ⭐ Q6. Full period মানে কী, আর কখন পাওয়া যায়?
**উত্তর:** Full period = period ঠিক `m`-এর সমান (সর্বোচ্চ সম্ভব, প্রতিটা মান ঠিক একবার আসে)। `m = 2^k`-এর জন্য **Hull–Dobell** শর্ত: (১) `c` ও `m` coprime, (২) `a−1`, `m`-এর প্রতিটা মৌলিক উৎপাদক দিয়ে বিভাজ্য, (৩) `m`, 4 দিয়ে বিভাজ্য হলে `a−1`-ও 4 দিয়ে বিভাজ্য।

---

### ⭐ Q7. `m`-কে 2-এর ঘাত (`2³¹`/`2³²`) নেওয়ার সুবিধা কী?
**উত্তর:** দুটো: (১) বড় `m` → বড় period। (২) binary-তে 2-এর ঘাত দিয়ে modulo খুব দ্রুত — শুধু শেষের bit-গুলো ফেলে দিলেই হয়, ব্যয়বহুল division লাগে না।

---

### Q8. MWC (Multiply-With-Carry) generator LCG থেকে কীভাবে আলাদা?
**উত্তর:** LCG-এ increment `c` স্থির; MWC-তে increment প্রতিবার বদলায় — modulo-র সময় যে অংশ "carry" (ভাগফল `⌊(ax+c)/m⌋`) হয়, সেটাই পরের increment হয়। ফলে অনেক বড় period (যেমন `~10^{18.2}`) ও ভালো পরিসংখ্যানগত গুণ।

---

### ⭐⭐ Q9. Mersenne Twister কী এবং কেন এত জনপ্রিয়?
**উত্তর:** Matsumoto–Nishimura (1997)-এর generator, NumPy/R/Python-এর default। ৬২৪টা মান দিয়ে শুরু করে bitwise XOR/shift/AND অপারেশনে পরের সংখ্যা বানায়। জনপ্রিয় কারণ: **বিশাল period (`~10^6001`)** এবং উচ্চ-মাত্রিক uniformity (৬২৩ dimension পর্যন্ত independent)।

---

### ⭐⭐ Q10. chi-square test দিয়ে কীভাবে যাচাই করো generator uniform কিনা?
**উত্তর:** `M`টা observation নিয়ে `K`টা category-তে ফেলি। H₀: প্রতিটা category-র probability `pₛ` (uniform হলে সব সমান)। test statistic:
`V = Σ (Yₛ − M pₛ)² / (M pₛ)`, যেখানে `Yₛ` = category `s`-এ observed count, `M pₛ` = expected count। H₀ সত্যি হলে `V ~ χ²` with `K−1` df। `V` ছোট → uniform মনে হয়; `V` বড় → uniformity বাতিল। শর্ত: `M pₛ ≥ 5`।

---

### Q11. chi-square test-এ `M pₛ ≥ 5` শর্ত আর "`M` খুব বড়ও না" — কেন?
**উত্তর:** `M pₛ ≥ 5` না হলে chi-square approximation দুর্বল (expected count খুব ছোট হলে distribution আর χ²-এর মতো থাকে না)। আবার `M` অতি বড় হলে: (১) computing সময় বেশি, (২) অতি সামান্য, বাস্তবে অগুরুত্বপূর্ণ পার্থক্যও "statistically significant" হয়ে যায়। তাই ভারসাম্য, যেমন `M = 10000`।

---

### Q12. k-tuple uniformity test কী যাচাই করে?
**উত্তর:** শুধু একক সংখ্যা না, **পরপর সংখ্যার ক্রমও** uniform/independent কিনা। sequence-কে `D`টা bin-এ scale করে (`yᵢ=⌊D xᵢ/m⌋`), তারপর `k`-tuple (`k=1,2,3`) নিয়ে chi-square (`K=Dᵏ`, `pₛ=1/Dᵏ`)। এতে লুকানো serial pattern (একটার পর আরেকটা নির্দিষ্ট মান আসার প্রবণতা) ধরা পড়ে।

---

### ⭐ Q13. Sobol sequence কী, আর random-এর চেয়ে কখন ভালো?
**উত্তর:** Sobol হলো quasi-random / low-discrepancy sequence — deterministic ভাবে বিন্দু এমনভাবে বসায় যাতে পুরো জায়গা সমানভাবে ঢাকে (random-এর মতো cluster/gap নেই)। Monte Carlo integration-এ এতে estimate দ্রুত converge করে (Quasi-Monte Carlo), `1/√S`-এর চেয়ে ভালো হারে। তাই integration-এর কাজে random-এর চেয়ে ভালো।

---

### Q14. (চিন্তা) একটা LCG বানালাম যার period মাত্র 4। simulation-এ ব্যবহার করলে কী সমস্যা?
**উত্তর:** মাত্র ৪টা সংখ্যা চক্রাকারে ফিরবে, তাই draw-গুলো independent না — প্রবল serial correlation ও নকল pattern। Monte Carlo estimate biased হবে, MCSE-র হিসাবও ভুল হবে (কারণ ওটা independence ধরে নেয়)। তাই full বা বড় period অপরিহার্য।

---

### Q15. এই chapter পুরো কোর্সে কীভাবে কাজে লাগে?
**উত্তর:** সব simulation-এর একদম নিচের ভিত্তি uniform pseudo-random সংখ্যা। Chapter 3-এ এই uniform থেকেই inverse-CDF/rejection দিয়ে যেকোনো distribution বানানো হবে। reproducibility (seed), generator-এর গুণ (period, uniformity) — পুরো কোর্সের প্রতিটা simulation-এর নির্ভরযোগ্যতা এর ওপর দাঁড়িয়ে।

---

## 🎯 দ্রুত flashcard
- pseudo-random = deterministic chaotic নিয়ম; seed → reproducible
- LCG: `x_{j+1}=mod_m(a x_j+c)`; period `≤ m` (pigeonhole)
- full period: Hull–Dobell; `m=2^k` দ্রুত modulo
- MWC: carry-ভিত্তিক variable increment
- Mersenne Twister: default, period `~10^6001`
- χ² test: `V=Σ(Yₛ−Mpₛ)²/(Mpₛ)`, df `K−1`, `Mpₛ≥5`
- Sobol: low-discrepancy, QMC দ্রুত converge
