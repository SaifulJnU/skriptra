# Chapter 2 — Exercise ও Solution (বাংলায় বুঝে বুঝে)

> Chapter 2-এর জন্য প্রাসঙ্গিক হলো **Exercise 1-এর Part 2: LCS Generators**। মূল ফাইল `01_exercise.pdf` ও solution `01_exercise_solution.ipynb` এই folder-এ copy করা আছে।

---

## 📋 প্রশ্ন (Part 2: LCS Generators)

- **(a)** LCG algorithm-এর একটা basic implementation করো। web search করে `a`, `c`, `m`-এর মধ্যে সেই সম্পর্ক বের করো যা **full period** (period `= m`) দেয়।
- **(b)** `m = 2³²`-এর জন্য full-period দেওয়া `a`, `c` নাও, ১০০০০ sample বানিয়ে (normalized) histogram আঁকো — চোখে uniformity যাচাই করো।

---

## ✅ (a) LCG implementation + full-period শর্ত

### Basic implementation
```python
def lcg(x, a, c, m):
    return (a * x + c) % m          # ঠিক slide-এর সূত্র: mod_m(a x + c)

def find_period(a, c, m, x0):       # period মাপার সহজ উপায়
    seen = []
    x = x0
    while x not in seen:            # যতক্ষণ না কোনো মান পুনরাবৃত্তি হয়
        seen.append(x)
        x = lcg(x, a, c, m)
    return len(seen)                # যতগুলো ভিন্ন মান এসেছিল = period
```
**`find_period` কীভাবে কাজ করে:** মান বানাতে থাকি যতক্ষণ না আগে-দেখা কোনো মান ফিরে আসে। যেহেতু recurrence deterministic, প্রথম পুনরাবৃত্তিই পুরো cycle-এর শেষ — তাই `seen`-এর দৈর্ঘ্যই period।

### Full-period শর্ত (Hull–Dobell Theorem) — এটাই (a)-এর মূল উত্তর
একটা LCG full period (`= m`) দেয় **যদি এবং কেবল যদি**:
1. **`c` ও `m` coprime** — অর্থাৎ `gcd(c, m) = 1`।
2. **`a − 1`, `m`-এর প্রতিটা মৌলিক উৎপাদক (prime factor) দিয়ে বিভাজ্য।**
3. **`m`, 4 দিয়ে বিভাজ্য হলে `a − 1`-ও 4 দিয়ে বিভাজ্য।**

### ছোট উদাহরণ দিয়ে যাচাই (`m = 16 = 2⁴`)
```python
m  = 2**4   # 16
c  = 3      # gcd(3,16)=1 → শর্ত ১ ঠিক (c, m coprime)
a  = 5      # a−1 = 4. 16-এর একমাত্র prime factor 2; 4, 2 দিয়ে বিভাজ্য → শর্ত ২ ঠিক।
            #          16, 4 দিয়ে বিভাজ্য, আর 4 ও 4 দিয়ে বিভাজ্য → শর্ত ৩ ঠিক।
x0 = 3
print(find_period(a, c, m, x0))   # → 16  (full period!)
```
**কেন `a=5, c=3` কাজ করল (মিলিয়ে দেখো):**
- `m = 16`, এর একমাত্র prime factor `2`। `a−1 = 4`, যা `2` দিয়ে বিভাজ্য ✔ (শর্ত ২)।
- `16` চার দিয়ে বিভাজ্য, তাই `a−1=4`-ও চার দিয়ে বিভাজ্য হতে হবে — `4/4=1` ✔ (শর্ত ৩)।
- `gcd(3,16)=1` ✔ (শর্ত ১)। তিনটাই মেলে → period = 16 = m।

> 💡 মজা করে একটা শর্ত ভাঙো (যেমন `c=2`, তখন `gcd(2,16)=2≠1`) → period কমে যাবে। নিজে চালিয়ে দেখো, viva-তে বললে teacher খুশি হবে।

---

## ✅ (b) `m = 2³²`-এ full-period generator + uniformity

বাস্তবে বহুল-ব্যবহৃত একটা full-period parameter সেট (Numerical Recipes-এর) — Hull–Dobell তিন শর্তই মানে:
```python
import matplotlib.pyplot as plt

m    = 2**32
a    = 1_664_525
c    = 1_013_904_223
seed = 156

N = 100_000
x = seed
samples = []
for _ in range(N):
    x = lcg(x, a, c, m)
    samples.append(x / m)          # ← [0,1)-এ normalize: পূর্ণসংখ্যাকে m দিয়ে ভাগ

plt.hist(samples, bins=50, density=True, alpha=0.7, edgecolor='black')
plt.xlabel("Value"); plt.ylabel("Density")
plt.title("LCG samples (normalized to [0,1])")
plt.show()
```

### ব্যাখ্যা — কী ঘটছে ও কেন উত্তর এমন
- **`x / m` কেন?** generator পূর্ণসংখ্যা দেয় `{0,…,m−1}`-এ; `m` দিয়ে ভাগ করে `[0,1)`-এ আনি — Chapter 2-এর "discrete uniform → `[0,maxinteger]`-কে ভাগ করে scale" ধারণার বাস্তব রূপ। এখন এগুলো প্রায় continuous Uniform(0,1)।
- **histogram প্রায় সমতল (flat)** হবে — প্রতিটা bin-এ প্রায় সমান উচ্চতা। এটাই চোখে-দেখা uniformity-র প্রমাণ।
- **`density=True` কেন?** histogram-কে normalize করে যাতে মোট ক্ষেত্রফল 1 — uniform হলে প্রতিটা bar উচ্চতা ≈ 1 (কারণ `[0,1)`-এ uniform density = 1)।
- full period থাকায় প্রথম ১০০০০০ সংখ্যায় কোনো পুনরাবৃত্তি নেই, তাই sample-গুলো ভালোভাবে পুরো `[0,1)` ঢাকে।

> ⚠️ **সতর্কতা (viva-তে বললে বাড়তি নম্বর):** এই histogram শুধু **1-D uniformity** দেখায়। LCG বিখ্যাত একটা দুর্বলতা: পরপর সংখ্যাগুলো higher dimension-এ plot করলে (যেমন `(xᵢ, xᵢ₊₁)` জোড়া) সব বিন্দু কয়েকটা সমান্তরাল hyperplane-এ পড়ে (Marsaglia-র "lattice" সমস্যা)। তাই 1-D histogram ভালো দেখালেও LCG সবসময় ভালো generator না — এজন্যই বাস্তবে Mersenne Twister ব্যবহৃত হয়।

---

## 🎯 এই exercise থেকে viva-তে যা আসতে পারে
1. "full period-এর তিন শর্ত বলো" → Hull–Dobell (coprime `c,m`; `a−1` prime factor দিয়ে বিভাজ্য; 4-শর্ত)।
2. "`a=5,c=3,m=16` কেন full period দেয়?" → তিন শর্ত মিলিয়ে দেখাও (উপরে করা আছে)।
3. "`x/m` কেন করলে?" → integer draw-কে `[0,1)` uniform-এ normalize।
4. "1-D histogram uniform দেখালেই কি generator ভালো?" → না; higher-dimension lattice সমস্যা থাকতে পারে; তাই Mersenne Twister।
5. "period মাপলে কীভাবে?" → মান বানাও যতক্ষণ না পুনরাবৃত্তি; গোনা ভিন্ন মানের সংখ্যাই period।

> ✍️ পরামর্শ: `01_exercise_solution.ipynb`-এ এই কোডগুলো আছে। নিজে `a, c, m, seed` পাল্টে period ও histogram কীভাবে বদলায় দেখো — বিশেষ করে একটা Hull–Dobell শর্ত ভেঙে period ছোট হওয়াটা চোখে দেখো।
