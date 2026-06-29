# Chapter 2 — Creating Pseudo-randomness with Computers (বাংলায় পুরো পাঠ)

> Chapter 1-এ আমরা ধরে নিয়েছিলাম "computer থেকে random draw পাওয়া যায়"। কিন্তু computer তো deterministic! এই chapter-এ শিখবো সেই random সংখ্যাগুলো **আসলে কীভাবে তৈরি হয়** — modulo, Linear Congruential Generator (LCG), Mersenne Twister — এবং কীভাবে যাচাই করি যে সেগুলো সত্যিই "random/uniform" (chi-square test)। শেষে দেখবো Sobol sequence, যা ইচ্ছে করেই "খুব বেশি random না" বানিয়ে কাজ আরও ভালো করে।

---

## 🎯 এক লাইনে পুরো chapter
> **"Computer সত্যিকারের random বানাতে পারে না, তাই সে একটা গোলমেলে (chaotic) গাণিতিক নিয়মে এমন সংখ্যা বানায় যা দেখতে random লাগে — আর আমরা chi-square test দিয়ে যাচাই করি সেগুলো সত্যিই uniform কিনা।"**

---

## ১. Randomness আসলে কী? আর computer-এর সমস্যা কোথায়?

### দার্শনিক প্রশ্ন
"Random" মানে কী? — যে সংখ্যার পরের মান আগেরগুলো দেখে **অনুমান করা যায় না**। সত্যিকারের randomness আসে প্রকৃতি থেকে (তেজস্ক্রিয় ক্ষয়, তাপীয় noise)।

### computer-এর দ্বন্দ্ব
Computer **deterministic** — একই input → একই output। তাহলে সে অনুমান-অযোগ্য সংখ্যা বানাবে কীভাবে? সমাধান: এমন একটা **চালাক, এলোমেলো (chaotic) ফাংশন** ব্যবহার করো যেটা এমন sequence বানায় যার পরের মান আগেরগুলো থেকে আন্দাজ করা কঠিন। দেখতে random লাগে — তাই নাম **pseudo-random** ("ছদ্ম-random")।

slide-এর সংক্ষেপ:
> PRNG = এমন algorithm যার পরের মান আগের মান থেকে unpredictable। "অজানা চোখে" random লাগে। মূলত আমাদের চাই **chaotic function**, আর যেহেতু computer discrete (পূর্ণসংখ্যা নিয়ে কাজ করে), তাই **পূর্ণসংখ্যার ওপর chaotic function** চাই।

### 🎲 analogy
জাদুকরের তাস-শাফলিং মনে করো। দর্শকের চোখে তাসের ক্রম এলোমেলো, কিন্তু জাদুকর ঠিক জানেন কোন তাস কোথায় — কারণ তিনি একটা গোপন নিয়মে শাফল করেছেন। PRNG হলো সেই জাদুকর: নিয়ম আছে, কিন্তু বাইরের চোখে random।

---

## ২. Discrete Uniform Distribution (ভিত্তি)

সব random number generator-এর লক্ষ্য প্রথমে **uniform** সংখ্যা বানানো (যেখানে প্রতিটা মান সমান সম্ভাবনার)। বাকি সব distribution এই uniform থেকে বানানো যায় (Chapter 3)।

discrete uniform density:
$$f(x\mid N) = \begin{cases} \frac{1}{N}, & x = x_1, x_2, \ldots, x_N \\ 0, & \text{নাহলে} \end{cases}$$

মানে `N`টা সম্ভাব্য মানের প্রতিটার probability সমান `1/N` (যেমন ন্যায্য পাশা: প্রতি মুখ `1/6`)।

**কীভাবে বানায়:** `[0, maxinteger]` ব্যবধিকে `N`টা সমান ভাগে ভাগ করে, প্রতিটা ভাগকে একটা `xᵢ`-এর সাথে মেলানো হয়। (অর্থাৎ বড় পূর্ণসংখ্যা বানিয়ে, তাকে scale/ভাগ করে কাঙ্ক্ষিত পরিসরে আনা হয়।)

---

## ৩. Modulo Operator — সব generator-এর কাঁচামাল

### Modulo মানে কী (সহজ ভাষায়)
`mod_m(b)` = `b`-কে `m` দিয়ে ভাগ করলে যে **ভাগশেষ** থাকে। যেমন `mod₁₂(27) = 3` (২৭÷১২ = ২, ভাগশেষ ৩)।

### 🕐 ঘড়ির analogy (সবচেয়ে সহজ)
ঘড়ি হলো `mod 12`। এখন ১০টা বাজে, ৫ ঘণ্টা পরে = ৩টা (১৫ নয়, কারণ ১৫ mod ১২ = ৩)। ঘড়ির কাঁটা ১২-এ পৌঁছে আবার ০ থেকে শুরু করে — এই "চক্রাকারে ফিরে আসা"-ই modulo-র মূল চরিত্র, আর এটাই generator-কে আবদ্ধ পরিসরে রাখে।

### আনুষ্ঠানিক
`a, b ∈ ℤ`, `m ∈ ℕ`। সমতুল্যতা সম্পর্ক `a ≡ b ⟺ (a − b), m দিয়ে নিঃশেষে বিভাজ্য`। এটা পূর্ণসংখ্যাকে **residue class**-এ ভাগ করে। প্রতিনিধি হিসেবে আমরা `{0, 1, 2, …, m−1}` নিই, এই সেটকে বলে `ℤ_m`।

---

## ৪. Linear Congruential Generator (LCG / LCS) — সবচেয়ে ক্লাসিক generator

### সংজ্ঞা
একটা Linear Congruential Sequence পুনরাবৃত্তিমূলকভাবে (recursively) সংজ্ঞায়িত:

$$x_{j+1} := \text{mod}_m(a\,x_j + c)$$

চারটা উপাদান:
- `x₀` = শুরুর মান (**seed**)
- `a` = multiplier (গুণক)
- `c` = increment (যোগফল)
- `m` = modulus, যেখানে `m > x₀, a, c`

**কাজের ধরন:** আগের সংখ্যাকে `a` দিয়ে গুণ করো, `c` যোগ করো, তারপর `m` দিয়ে ভাগশেষ নাও। এই ভাগশেষই পরের সংখ্যা। তারপর আবার একই কাজ — চক্রাকারে।

### 🌰 উদাহরণ (slide থেকে)
**উদাহরণ ১:** `a=5, c=7, x₀=4, m=13`
```
x₁ = mod₁₃(5·4+7) = mod₁₃(27) = 1
x₂ = mod₁₃(5·1+7) = mod₁₃(12) = 12
x₃ = mod₁₃(5·12+7)= mod₁₃(67) = 2
x₄ = mod₁₃(5·2+7) = mod₁₃(17) = 4 = x₀   ← আবার শুরুতে ফিরে এলো!
x₅ = x₁ ...                                ← এখান থেকে পুনরাবৃত্তি
```
এখানে sequence `4` ঘর পরপর নিজেকে পুনরাবৃত্তি করছে → **period = 4**।

**উদাহরণ ২:** `a=5, c=7, x₀=4, m=8`
```
x₁=3, x₂=6, x₃=5, x₄=0, x₅=7, x₆=2, x₇=1, x₈=4=x₀, x₉=x₁...
```
এখানে period = **8** = `m` → **full period** (সবচেয়ে ভালো!)।

> খেয়াল করো: একই `a, c, x₀` কিন্তু ভিন্ন `m`-এ period পুরো বদলে গেল। প্যারামিটার পছন্দ কত গুরুত্বপূর্ণ — সেটাই মূল শিক্ষা।

---

## ৫. Period of LCG — কেন এটা গুরুত্বপূর্ণ

### Period মানে কী
sequence `{xⱼ}`-এর period `μ` হলো সেই সংখ্যা যেখানে `xⱼ = x_{j+μ}` সব `j`-এর জন্য। মানে কত ঘর পর sequence নিজেকে হুবহু পুনরাবৃত্তি করে।

### ⭐ Theorem: প্রতিটা LCG-এর period `≤ m`
**প্রমাণ (pigeonhole / কবুতর-খোপ নীতি — viva-তে চাইতে পারে):**
- প্রতিটা `x_{j+1} = mod_m(...)` সবসময় `{0,1,…,m−1}`-এর ভেতরে, মানে মাত্র **`m`টা সম্ভাব্য মান**।
- তাই প্রথম `m+1`টা পদের মধ্যে অন্তত দুটো অবশ্যই সমান হবে (কারণ `m+1`টা পদ কিন্তু খোপ মাত্র `m`টা — pigeonhole)।
- আর যেহেতু পরের পদ সম্পূর্ণভাবে বর্তমান পদ দিয়ে নির্ধারিত (deterministic recurrence), একবার কোনো মান পুনরাবৃত্তি হলে পুরো sequence-ই পুনরাবৃত্তি হবে।
- ∴ period `μ ≤ m`। ∎

### কেন বড় period চাই?
period ছোট হলে generator অল্প কিছু সংখ্যা ঘুরিয়ে ফিরিয়ে দেয় → simulation-এ "নকল" pattern তৈরি হয়, ফলাফল biased হয়। তাই আমরা **যত বড় সম্ভব period** চাই।

---

## ৬. Modulus `m` কীভাবে বাছবো

slide-এর পরামর্শ:
- **`m` খুব বড় নাও** → বড় period সম্ভব।
- **`m`-কে 2-এর ঘাত নাও**, যেমন `m = 2³¹` বা `2³²` → কারণ binary system-এ 2-এর ঘাত দিয়ে ভাগ/modulo খুব দ্রুত হয় (শুধু শেষের bit-গুলো ফেলে দিলেই হয়, ব্যয়বহুল division লাগে না)।

> 💡 **Full period-এর শর্ত (Hull–Dobell theorem — exercise-এ লাগবে):** `m = 2^k` হলে full period (`= m`) পেতে হয় (১) `c` ও `m` coprime (গসাগু 1), (২) `a−1`, `m`-এর প্রতিটা মৌলিক উৎপাদক দিয়ে বিভাজ্য, (৩) `m`, 4 দিয়ে বিভাজ্য হলে `a−1`-ও 4 দিয়ে বিভাজ্য। (বিস্তারিত exercise solution-এ।)

---

## ৭. Multiply-With-Carry (MWC) Generator — LCG-এর উন্নত ভাই

LCG-এর সহজ একটা সম্প্রসারণ (Marsaglia, 1996), যা আধুনিক test-এও ভালো করে:

$$x_{j+1} = \text{mod}_m(a x_j + c_j), \quad c_{j+1} = \left\lfloor \frac{a x_j + c_j}{m} \right\rfloor$$

পার্থক্য: এখানে `c` স্থির না — প্রতিবার increment বদলায়, কারণ modulo-র সময় যে অংশটা "carry" (ভাগফল) হয়, সেটাই পরের increment হয়। (যেমন যোগ করতে গিয়ে হাতে রাখা সংখ্যা।)

ফল: `a = 698769069, m = 2³²` নিলে period প্রায় `2^{50.4} ≈ 10^{18.2}` — বিশাল!

---

## ৮. Mersenne Twister — বাস্তবে সবচেয়ে বেশি ব্যবহৃত

Matsumoto ও Nishimura (1997)। NumPy, R, Python সহ অনেক জায়গায় default generator।

মূল ধারণা (বিস্তারিত bit-অপারেশন মুখস্থ লাগবে না, কিন্তু এগুলো জানো):
- **৬২৪টা শুরুর মান** (seed = ৬২৪টা সংখ্যা) দিয়ে শুরু।
- পরের মান হিসাব হয় **bitwise XOR (⊕)**, **shift**, ও **AND (∧)** অপারেশন দিয়ে — পূর্ণসংখ্যার binary রূপের ওপর কাজ করে। (XOR: bit আলাদা হলে 1, একই হলে 0; AND: দুটোই 1 হলে 1।)
- **period অবিশ্বাস্য বড়:** প্রায় `10^{6001}`!
- generated সংখ্যাগুলো **৬২৩ মাত্রায় (dimensions) independent** — মানে ৬২৪তম সংখ্যা ১ম সংখ্যার ওপর নির্ভর করে, কিন্তু ২ থেকে ৬২৩ নম্বরের ওপর না।

> 💡 **Viva-corner:** "কোন generator বাস্তবে সবচেয়ে বেশি ব্যবহৃত হয় এবং কেন?" → **Mersenne Twister** — কারণ বিশাল period (`~10^6001`) ও উচ্চ-মাত্রিক uniformity। (নাম মনে রাখো, এক লাইন বললেই হবে।)

---

## ৯. Randomness যাচাই — সংখ্যাগুলো কি সত্যিই uniform/random? (Chi-square Test)

আমরা generator বানালাম — কিন্তু এর সংখ্যা কি সত্যিই uniform? এটা যাচাই করতে **chi-square (χ²) test**।

### মূল ধারণা
`M`টা স্বাধীন observation নাও, যেগুলো `K`টা ভিন্ন category-তে পড়তে পারে। আমরা hypothesis test করি:
> **H₀:** একটা observation category `s`-এ পড়ার probability ঠিক `pₛ` (`1 ≤ s ≤ K`)।

(uniformity test-এ সব `pₛ` সমান হওয়ার কথা।)

**শর্ত:** `M` যথেষ্ট বড় হতে হবে যাতে `M·pₛ ≥ 5` প্রতিটা category-তে (নাহলে chi-square approximation ভালো কাজ করে না)।

### Test statistic
`Yₛ` = category `s`-এ পড়া observation-এর সংখ্যা। তাহলে:

$$V = \sum_{s=1}^{K} \frac{(Y_s - M p_s)^2}{M p_s}$$

মানে: প্রতিটা category-তে **যা দেখলাম (`Yₛ`) vs যা প্রত্যাশিত (`M pₛ`)** — পার্থক্যের বর্গ, প্রত্যাশিত দিয়ে ভাগ, সব যোগ। H₀ সত্যি হলে `V` প্রায় **chi-square distribution মেনে চলে `K−1` degrees of freedom-এ**।

- `V` ছোট → observed ≈ expected → generator uniform মনে হচ্ছে (H₀ রাখি)।
- `V` বড় → observed অনেক আলাদা → generator সম্ভবত uniform না (H₀ বাতিল)।

**`M`-এর ভারসাম্য:** `M·pₛ ≥ 5` মেটাতে যথেষ্ট বড়, কিন্তু খুব বড়ও না (computing time বাঁচাতে, আর খুব সামান্য পার্থক্যও যেন "significant" হয়ে না যায়), যেমন `M = 10000`।

### 🎯 Uniformity test (k-tuples) — শুধু এক-একটা না, ক্রমও যাচাই
ভালো generator-এ শুধু individual সংখ্যা না, **পরপর সংখ্যাও** independent ও uniform হওয়া উচিত। তাই sequence-এর `k`টা পরপর মানের tuple যাচাই করি:
- প্রথমে scale: `yᵢ = ⌊D·xᵢ/m⌋` যাতে `0 ≤ yᵢ < D` (D-টা bin বানালাম)।
- `k=1`: একক মান `(yᵢ)` — category সংখ্যা `K = D`।
- `k=2`: জোড়া `(y_{2j−1}, y_{2j})` — `K = D²`।
- `k=3`: ত্রয়ী — `K = D³`। প্রতিটায় `pₛ = 1/Dᵏ` ধরে chi-square।

এতে ধরা পড়ে যদি generator-এ লুকানো pattern থাকে (যেমন একটা মানের পরে আরেকটা প্রায়ই আসে)।

---

## ১০. Sobol Sequence — "কম random" করে আরও ভালো ফল

একটা চমকপ্রদ মোড়: কখনো কখনো আমরা **ইচ্ছে করেই খুব বেশি random চাই না**!

### সমস্যা: সত্যিকারের random-এ "গুচ্ছ" আর "ফাঁক" থাকে
genuinely random বিন্দু একটা বর্গক্ষেত্রে ছড়ালে, কিছু জায়গায় বিন্দু **গুচ্ছ (cluster)** হয়, কিছু জায়গা **ফাঁকা** থাকে। মানে জায়গাটা সমানভাবে ঢাকা পড়ে না।

### সমাধান: Sobol sequence (quasi-random / low-discrepancy)
Sobol sequence **নির্ধারিতভাবে (deterministically)** বিন্দু এমনভাবে বসায় যাতে পুরো জায়গা **সমানভাবে** ঢাকে — গুচ্ছ বা ফাঁক প্রায় থাকে না। slide-এর ছবিতে: বাঁয়ে random (এবড়োখেবড়ো, ফাঁকওয়ালা), ডানে Sobol (সুন্দর সমান বিন্যাস)।

### কেন কাজে লাগে?
Monte Carlo integration (Chapter 1)-এ যদি বিন্দু সমানভাবে ছড়ানো থাকে, integral-এর estimate **দ্রুত** convergence করে (error কমে `1/√S`-এর চেয়েও ভালো হারে)। একে বলে **Quasi-Monte Carlo**।

> 💡 **Viva-corner:** "random আর Sobol-এর পার্থক্য?" → random-এ cluster/gap থাকে; Sobol (low-discrepancy) জায়গা সমানভাবে ঢাকে, তাই Quasi-Monte Carlo integration-এ দ্রুত converge করে।

---

## ১১. পুরো Chapter এক নজরে (revision card)

| ধারণা | মূল কথা |
|---|---|
| Pseudo-random | deterministic computer-এ chaotic নিয়মে "random-মতো" সংখ্যা |
| Seed | শুরুর মান; fix করলে reproducible |
| Modulo | ভাগশেষ; ঘড়ির মতো চক্রাকার; generator-এর কাঁচামাল |
| LCG | `x_{j+1}=mod_m(a x_j + c)`; ৪ প্যারামিটার `x₀,a,c,m` |
| Period | কত ঘর পর পুনরাবৃত্তি; **`≤ m`** (pigeonhole); বড় চাই |
| Full period | `period = m`; Hull–Dobell শর্ত |
| `m` পছন্দ | বড় + 2-এর ঘাত (দ্রুত modulo) |
| MWC | variable increment (carry); বড় period |
| Mersenne Twister | বাস্তবে default; period `~10^6001` |
| Chi-square test | observed vs expected; `V=Σ(Yₛ−Mpₛ)²/(Mpₛ)`; df `K−1`; `Mpₛ≥5` |
| Sobol | quasi-random, low-discrepancy; সমান বিন্যাস; QMC দ্রুত converge |

---

## ১২. নিজেকে যাচাই করো
1. computer pseudo-random সংখ্যা বানায় কীভাবে? "chaotic function over integers" মানে কী?
2. LCG-এর সূত্র ও ৪টা প্যারামিটার বলো। modulo-র ভূমিকা কী?
3. কেন প্রতিটা LCG-এর period `≤ m`? pigeonhole যুক্তিটা বলো।
4. `m`-কে 2-এর ঘাত নেওয়ার সুবিধা কী?
5. Mersenne Twister কেন বিখ্যাত?
6. chi-square test দিয়ে uniformity কীভাবে যাচাই করো? `V` বড় হলে কী বুঝবে?
7. Sobol sequence কেন কখনো random-এর চেয়ে ভালো?

> বিস্তারিত উত্তর `02_Viva_Questions_Bangla.md`-এ; LCG full-period exercise + code `03_Exercise_and_Solution_Bangla.md`-এ।
