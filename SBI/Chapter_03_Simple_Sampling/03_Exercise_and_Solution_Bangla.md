# Chapter 3 — Exercise ও Solution (বাংলায় বুঝে বুঝে)

> Chapter 3-এর জন্য প্রাসঙ্গিক হলো **Exercise 2** (exponential sampling, Poisson via exponential, importance sampling for Cauchy tail)। মূল ফাইল `SOURCE_02_exercise.pdf` ও solution `SOURCE_02_exercise_solution.ipynb` এই folder-এ copy করা আছে।

---

## 📋 প্রশ্ন (Exercise 2)
1. **Exponential sampling:** (a) CDF থেকে quantile function বের করে, Uniform(0,1) দিয়ে Exp(λ=5) sample করো। (b) Poisson ও exponential-এর সম্পর্ক জানো। (c) exponential draw দিয়ে Poisson(5) sample করো।
2. **Importance sampling:** Cauchy distribution-এর tail probability `θ=P(X>2)` estimate করো — (a) সাধারণ Monte Carlo, (b) standard-normal proposal দিয়ে IS, (c) truncated-Cauchy proposal দিয়ে IS, (d) সত্যিকারের মানের সাথে তুলনা।

---

## ✅ 1(a) Exponential-এর quantile function — inverse-CDF প্রয়োগ

এটাই Chapter 3-এর কেন্দ্রীয় কৌশলের সরাসরি প্রয়োগ।

**ধাপ ১ — CDF বের করো** (density `f(x)=λe^{−λx}`, `x≥0`):
$$F(x)=\int_0^x \lambda e^{-\lambda t}\,dt = 1-e^{-\lambda x}$$

**ধাপ ২ — CDF উল্টে quantile** (`p=F(x)` ধরে `x` বের করি):
```
p = 1 − e^{−λx}
e^{−λx} = 1 − p
−λx = ln(1−p)
x = −(1/λ) ln(1−p)
```
$$\boxed{Q(p)=F^{-1}(p)=-\frac{1}{\lambda}\ln(1-p)}$$

**কোড:**
```python
import random, math
def sample_exponential(rate=5):
    p = random.uniform(0, 1)
    return -(1/rate) * math.log(1 - p)      # quantile function
```
**কেন কাজ করে:** Chapter 3-এর উপপাদ্য — Uniform draw-এ `F⁻¹` লাগালে target distribution থেকে sample। histogram আঁকলে `5e^{−5x}` curve-এর সাথে মিলবে।

> 💡 ছোট্ট নোট: `1−p`-ও Uniform(0,1), তাই অনেকে `−(1/λ)ln(p)` লেখে — একই ফল। viva-তে জিজ্ঞেস করলে এটা বলতে পারো।

---

## ✅ 1(b) Poisson ও Exponential-এর সম্পর্ক

**মূল কথা:** যদি ঘটনাগুলোর **মধ্যবর্তী সময় (waiting time)** i.i.d. Exponential(λ) হয়, তাহলে নির্দিষ্ট একক সময়ে **ঘটনার সংখ্যা** Poisson(λ)। গাণিতিকভাবে — `Eᵢ~Exp(λ)`, partial sum `Sₖ=ΣEᵢ` হলে:
$$P(S_k \le x < S_{k+1}) = \frac{(\lambda x)^k}{k!}e^{-\lambda x}$$
যা ঠিক Poisson(λx)-এর pmf।

> 🚌 analogy: বাস আসার মধ্যে গড় ব্যবধান exponential; ১ ঘণ্টায় কয়টা বাস এলো — সেটা Poisson। দুটো একই প্রক্রিয়ার দুই দৃষ্টিকোণ।

---

## ✅ 1(c) Exponential draw দিয়ে Poisson(5) sample

(b)-র সম্পর্ক কাজে লাগাই: `x=1` ধরে, exponential draw যোগ করতে থাকি যতক্ষণ যোগফল `1` ছাড়িয়ে না যায়; কয়টা লাগল (শেষেরটা বাদে), সেটাই Poisson(5) sample।

```python
def sample_poisson(mean=5):
    cumulative_sum = 0.0
    count = 0
    while cumulative_sum <= 1.0:                    # threshold x=1
        cumulative_sum += sample_exponential(rate=mean)
        count += 1
    return count - 1     # শেষ draw threshold পার করিয়ে দিল, তাই −1
```
**ব্যাখ্যা:** `N = max{k : Sₖ ≤ 1}` দেয় `Poisson(5)`। বহু sample-এর histogram আঁকলে Poisson(5)-এর আকৃতি পাবে। এটা **distribution-সম্পর্ক দিয়ে sampling**-এর সুন্দর উদাহরণ (Chapter 3-এর পর্ব ১০-এর চেতনা)।

---

## ✅ 2. Importance Sampling — Cauchy-র tail `P(X>2)`

Cauchy distribution-এর লেজ খুব মোটা; standard Cauchy: `p(x)=1/(π(1+x²))`, CDF `F(x)=½+(1/π)arctan(x)`, quantile `Q(p)=tan(π(p−½))`। সত্যিকারের মান: `θ=1−F(2)≈0.1476`।

### (a) সাধারণ Monte Carlo
```python
S = 1_000_000
u = np.random.uniform(0,1,S)
x = np.tan(np.pi*(u-0.5))           # Cauchy draw (inverse-CDF)
theta_mc = np.mean(x > 2)            # P(X>2) ≈ indicator-এর গড়
```
**ফল:** ঠিক মানের কাছাকাছি, কিন্তু `X>2` বিরল বলে variance বেশি — অনেক (10⁶) sample লাগল। এটাই সমস্যা।

### (b) Importance sampling, proposal = standard Normal
```python
S = 10000
x = np.random.normal(0,1,S)                          # q = Normal(0,1) থেকে draw
w = cauchy_pdf(x) / normal_pdf(x)                    # weight r = p/q
indicator = (x > 2).astype(float)                    # f(x)=𝟙(x>2)
theta = np.sum(indicator*w) / np.sum(w)              # self-normalized IS
```
**সতর্কতা (গুরুত্বপূর্ণ):** Normal-এর লেজ Cauchy-র চেয়ে **অনেক পাতলা**। তাই `x>2`-এর অঞ্চলে Normal প্রায় sample-ই দেয় না, আর সেখানে `p/q` বিশাল হয়ে যায় → weight অস্থির, estimate noisy। **এটাই "proposal-এর লেজ পাতলা হলে IS ভেঙে পড়ে"-র জীবন্ত উদাহরণ।**

### (c) Importance sampling, proposal = truncated Cauchy on [2, large]
```python
# q = Cauchy কিন্তু শুধু [2, b]-তে সীমাবদ্ধ (normalize করা)
# এখানে সব draw সরাসরি x>2 অঞ্চলে পড়ে
```
**কেন এটা সবচেয়ে ভালো:** proposal ঠিক যেখানে আমরা আগ্রহী (`x>2`) সেখানেই কেন্দ্রীভূত, আর target-এর মতোই Cauchy-আকৃতির — তাই weight স্থিতিশীল, অল্প sample-এই নির্ভুল।

### (d) তুলনা
| পদ্ধতি | sample লাগল | মান | মন্তব্য |
|---|---|---|---|
| (a) সাধারণ MC | ~10⁶ | ≈0.1476 | ঠিক, কিন্তু অপচয়ী |
| (b) Normal proposal | 10⁴ | কাছাকাছি কিন্তু **noisy** | পাতলা লেজ → weight অস্থির |
| (c) truncated Cauchy | 10⁴ | নির্ভুল ও স্থিতিশীল | ভালো-মেলানো proposal |
| সত্য মান | — | **0.1476** | `1−F(2)` |

> 🎯 **মূল শিক্ষা:** importance sampling-এর সাফল্য পুরোপুরি **proposal `q` কতটা ভালো মেলে** তার ওপর নির্ভর। ভালো proposal (c) → কম sample-এ নির্ভুল; খারাপ proposal (b) → অস্থির।

---

## 🎯 এই exercise থেকে viva-তে যা আসতে পারে
1. "exponential-এর quantile derive করো" → CDF `1−e^{−λx}` উল্টে `−(1/λ)ln(1−p)`।
2. "Poisson আর exponential-এর সম্পর্ক?" → waiting time exponential হলে count Poisson।
3. "Cauchy tail-এ সাধারণ MC-র সমস্যা কী?" → rare event, বেশি variance।
4. "Normal proposal খারাপ কেন, truncated Cauchy ভালো কেন?" → লেজের পুরুত্ব ও proposal-target মিল।
5. "self-normalized estimator-এ ভাগ কেন?" → unknown normalizing constant কাটাতে + weight normalize।

> ✍️ পরামর্শ: `SOURCE_02_exercise_solution.ipynb` নিজে run করো; (b)-তে seed পাল্টে দেখো estimate কত নড়ে (অস্থিরতা), (c)-তে কত স্থিতিশীল — হাতে-কলমে দেখলে IS-এর মূল পাঠ গেঁথে যাবে।
