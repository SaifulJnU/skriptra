# Chapter 1 — Exercise ও Solution (বাংলায় বুঝে বুঝে)

> Chapter 1-এর জন্য প্রাসঙ্গিক হলো **Exercise 1-এর Part 1: Logit-Normal Distribution** (change of variables + MCSE — ঠিক যা আমরা lesson-এ শিখলাম)। মূল ফাইল `01_exercise.pdf` ও solution `01_exercise_solution.pdf` / `01_exercise_solution.ipynb` এই folder-এ copy করা আছে। (Exercise 1-এর Part 2 = LCS generators, যা Chapter 2-এর — ওটার ব্যাখ্যা Chapter 2 folder-এ।)

---

## 📋 প্রশ্ন (Part 1: Logit-Normal Distribution)

`X ~ N(0,1)`। logit-normal transform: `Y = f(X) = 1 / (1 + e^{-X})` (এটা sigmoid / inverse-logit ফাংশন)।

- **(a)** দেখাও `Y ∈ (0,1)`। inverse transform `X = f⁻¹(Y)` বের করো।
- **(b)** change-of-variable (Jacobian) দিয়ে `Y`-এর density `p_Y(y)` বের করো।
- **(c)** `N(0,1)` থেকে ১০০০ sample নিয়ে transform করে histogram আঁকো।
- **(d)** (b)-এর analytical PDF overlay করো; `S = 10, 100, 1000, 10000`-এ fit কেমন বদলায় দেখাও।
- **(e)** প্রতি `S`-এ MCSE হিসাব করে MCSE vs S plot করো; সম্পর্ক নির্ণয় করো।

---

## ✅ (a) Domain আর Inverse — ধাপে ধাপে

### `Y ∈ (0,1)` কেন? (intuition আগে)
`Y = 1/(1+e^{-X})` হলো বিখ্যাত **sigmoid** ("S"-আকৃতির) ফাংশন। এটা যেকোনো বাস্তব সংখ্যা `X ∈ (−∞, ∞)`-কে চেপে এনে `(0,1)`-এর ভেতর বসায়।

প্রমাণ (সহজ যুক্তি):
- যেকোনো `X`-এর জন্য `e^{-X} > 0` সবসময় (exponential কখনো ঋণাত্মক বা শূন্য না)।
- তাই denominator `1 + e^{-X} > 1` → `Y = 1/(1+e^{-X}) < 1`।
- আবার `e^{-X} > 0` হওয়ায় `Y > 0`।
- প্রান্ত যাচাই: `X → −∞` হলে `e^{-X} → ∞`, তাই `Y → 0` (ছোঁয় না)। `X → +∞` হলে `e^{-X} → 0`, তাই `Y → 1` (ছোঁয় না)।
- ∴ `Y` কঠোরভাবে `(0,1)`-এর ভেতরে। ∎

### Inverse `X = f⁻¹(Y)` বের করা
`Y = 1/(1+e^{-X})` থেকে শুরু করে `X` বের করি:

```
1/Y = 1 + e^{-X}
e^{-X} = 1/Y − 1 = (1 − Y)/Y
−X = ln((1 − Y)/Y)
 X = ln(Y/(1 − Y))      ← উল্টে নিলে চিহ্ন বদলে যায়
```

$$\boxed{X = f^{-1}(Y) = \ln\!\left(\frac{Y}{1-Y}\right)}$$

> এটাই বিখ্যাত **logit** ফাংশন — sigmoid-এর ঠিক উল্টো। (নামেই বোঝা যায়: logit ↔ inverse-logit/sigmoid।)

---

## ✅ (b) Density বের করা (Jacobian দিয়ে) — Chapter-এর মূল দক্ষতা

আমরা lesson-এর সূত্র ব্যবহার করবো:
$$p_Y(y) = p_X\big(f^{-1}(y)\big)\cdot \left|\frac{d}{dy}f^{-1}(y)\right|$$

**ধাপ ১ — Jacobian হিসাব।** ধরি `g(y) = f⁻¹(y) = ln(y/(1−y))`। log-নিয়মে ভাঙি:
```
g(y) = ln(y) − ln(1 − y)
```
পদে পদে derivative:
```
g'(y) = 1/y − ( −1/(1−y) ) = 1/y + 1/(1−y)
```
common denominator-এ এনে:
```
g'(y) = [(1−y) + y] / [y(1−y)] = 1 / [y(1−y)]
```
যেহেতু `0<y<1`-এ এটা ধনাত্মক, absolute value একই:
$$\left|\frac{d}{dy}f^{-1}(y)\right| = \frac{1}{y(1-y)}$$

**ধাপ ২ — Normal density বসানো।** `p_X(x) = (1/√2π) exp(−x²/2)`, যেখানে `x = ln(y/(1−y))`। তাই:

$$\boxed{p_Y(y) = \frac{1}{\sqrt{2\pi}}\exp\!\left(-\frac{1}{2}\Big[\ln\!\big(\tfrac{y}{1-y}\big)\Big]^2\right)\cdot \frac{1}{y(1-y)}, \quad 0<y<1}$$

এটাই **logit-normal PDF**।

### 🔍 কেন উত্তরটা এই রূপে — অন্তর্দৃষ্টি
- প্রথম অংশ (`exp(...)`) = মূল normal-এর আকৃতি, কিন্তু এখন `y`-এর ভাষায় (logit বসিয়ে)।
- শেষ অংশ `1/[y(1−y)]` = **Jacobian** — sigmoid যেহেতু `0` ও `1`-এর কাছে খুব চ্যাপ্টা (অনেক `X` অল্প জায়গায় চাপে), ওখানে density বাড়াতে হয়; এই factor ঠিক সেটাই করে (`y→0` বা `y→1`-এ `1/[y(1−y)]` বড় হয়)।
- Jacobian বাদ দিলে integral 1 হতো না — উত্তর ভুল।

---

## ✅ (c)(d) Sampling করে যাচাই — কোড বুঝে

মূল ধারণা: **transform via draws** (Jacobian ছাড়া!) — `N(0,1)` থেকে draw নিয়ে শুধু sigmoid লাগাই, ওগুলোই logit-normal sample।

```python
import numpy as np
import matplotlib.pyplot as plt
np.random.seed(42)                      # reproducible (Chapter 1-এর seed ধারণা)

def f(x):  return 1.0 / (1 + np.exp(-x))         # sigmoid (inverse-logit)
def g(y):  return np.log(y / (1.0 - y))          # logit (inverse transform)

def pdf_normal(x):  return (1/np.sqrt(2*np.pi)) * np.exp(-0.5*x**2)
def pdf_logitnormal(y):                          # (b)-তে derive করা PDF
    x = g(y)
    return pdf_normal(x) * (1.0 / (y*(1.0 - y))) # ← শেষ অংশই Jacobian

# analytical curve
y_vals = np.linspace(0.001, 0.999, 500)          # 0 আর 1 ছোঁয় না (domain খোলা)
pdf_vals = pdf_logitnormal(y_vals)

for S in [10, 100, 1000, 10000]:
    x = np.random.randn(S)                       # N(0,1) draw
    y = f(x)                                      # transform → logit-normal draw
    plt.hist(y, bins=30, density=True, alpha=0.5) # density=True জরুরি, নাহলে PDF-এর সাথে scale মিলবে না
    plt.plot(y_vals, pdf_vals, 'r-')
    plt.title(f'S={S}'); plt.show()
```

### কোডের সূক্ষ্ম কিন্তু গুরুত্বপূর্ণ পয়েন্ট (viva-তে জিজ্ঞেস হতে পারে)
- **`density=True` কেন?** histogram-কে normalize করে (মোট ক্ষেত্রফল 1) যাতে analytical PDF curve-এর সাথে একই scale-এ থাকে। না দিলে histogram হবে raw count — curve-এর সাথে মিলবে না।
- **`linspace(0.001, 0.999)` কেন, `0`/`1` নয়?** কারণ domain `(0,1)` **খোলা**; `y=0` বা `y=1`-এ `1/[y(1−y)]` ভাগশূন্য (infinity) হয়ে যায়।
- **fit কেমন বদলায়:** `S=10`-এ histogram খুব এবড়োখেবড়ো, curve-এর সাথে মেলে না; `S` বাড়ার সাথে সাথে histogram ধীরে ধীরে curve-কে নিখুঁতভাবে ঢাকে — এটাই Law of Large Numbers-এর চাক্ষুষ রূপ।

---

## ✅ (e) MCSE হিসাব ও সম্পর্ক — Chapter-এর সবচেয়ে মার্কওয়ালা অংশ

এখানে আমরা `E[Y]` (mean of Y) estimate করছি, তার MCSE মাপছি:
`MCSE = (sample std of Y) / √S` — মানে `√(Var(Y)/S)`।

```python
sample_sizes = [10, 100, 1000, 10000]
mcse_values = []
for S in sample_sizes:
    y = f(np.random.randn(S))
    sample_std = np.std(y, ddof=1)       # ddof=1 → unbiased (sample) variance
    mcse = sample_std / np.sqrt(S)       # MCSE সূত্র
    mcse_values.append(mcse)

plt.plot(sample_sizes, mcse_values, 'o-')
plt.xscale('log'); plt.yscale('log')     # log-log হলে সম্পর্ক সরলরেখা হবে
plt.xlabel('S'); plt.ylabel('MCSE'); plt.show()
```

### ফলাফল ও ব্যাখ্যা (এটাই উত্তর)
- log-log plot-এ MCSE vs S একটা **নিচের দিকে নামা সরলরেখা**, slope ≈ **−0.5**।
- কারণ `MCSE = √(Var/S) ∝ S^{−1/2}` → `log(MCSE) = const − ½ log(S)`।
- **সম্পর্ক:** sample size ৪ গুণ বাড়ালে MCSE অর্ধেক হয়; ১০০ গুণ বাড়ালে ১০ ভাগের ১ ভাগ। অর্থাৎ **MCSE ∝ 1/√S**।

> 💡 **`ddof=1` কেন?** এটা sample variance-কে unbiased করে (n−1 দিয়ে ভাগ, n নয়)। ছোট `S`-এ পার্থক্য গুরুত্বপূর্ণ। viva-তে জিজ্ঞেস করলে এই এক লাইন বললেই হবে।

---

## 🎯 এই exercise থেকে viva-তে যা আসতে পারে
1. "Jacobian-টা physically কী করছে?" → `0`/`1`-এর কাছে sigmoid চ্যাপ্টা, তাই ওখানে density বাড়াতে `1/[y(1−y)]` লাগে।
2. "transform করতে Jacobian লাগলো, কিন্তু sample করতে লাগলো না কেন?" → density vs draws (Chapter 1, Q14)।
3. "MCSE plot-এর slope কেন −0.5?" → `MCSE ∝ S^{−1/2}`।
4. "`density=True` না দিলে কী হতো?" → histogram scale PDF-এর সাথে মিলত না।
5. "logit আর sigmoid-এর সম্পর্ক?" → একটা আরেকটার inverse।

> ✍️ পরামর্শ: notebook-টা (`01_exercise_solution.ipynb`) নিজে একবার run করো, `S` ও `seed` পাল্টে দেখো histogram ও MCSE কীভাবে নড়ে — হাতে-কলমে করলে viva-তে আত্মবিশ্বাস আসবে।
