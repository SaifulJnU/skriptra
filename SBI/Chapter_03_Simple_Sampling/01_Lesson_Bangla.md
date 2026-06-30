# Chapter 3 — Sampling from Simple Distributions (বাংলায় পুরো পাঠ)

> এই chapter-টা একটা **টুলবক্স**: Uniform(0,1) থেকে শুরু করে যেকোনো distribution (discrete, continuous, multivariate) থেকে কীভাবে draw বানানো যায়। শেষে **importance sampling** — যেখান থেকে sample করা যায় না সেখান থেকেও আনুমানিক করার চালাকি। একটাও কৌশল বাদ দেব না।

---

## 🎯 এক লাইনে পুরো chapter
> **"হাতে শুধু Uniform(0,1); সেটাকে নানা কৌশলে (interval ভাগ, CDF উল্টানো, distribution-সম্পর্ক, Cholesky, reweighting) রূপান্তর করে যেকোনো distribution থেকে draw বানাই।"**

---

# পর্ব A — Discrete distribution থেকে sampling

মূল কৌশল সবগুলোর একই: **uniform integer interval `[0, m]`-কে টুকরো করে ভাগ করো; RNG যে টুকরোয় পড়ে, সেই টুকরোর মানটাই উত্তর।**

## ১. Bernoulli distribution (সবচেয়ে সরল — হ্যাঁ/না)
একটা ঘটনা ঘটল কি ঘটল না — `x ∈ {0,1}`।
$$p(x\mid \pi) = \pi^x (1-\pi)^{1-x}$$
এখানে `π` = ঘটনার probability (এটাই mean)।

**কীভাবে sample করি:** integer interval `I=[0,m]`-কে দুই ভাগ `I₀, I₁`-এ ভাগ করি, যাতে `I₁`-এর আকার `≈ π·|I|`। RNG যদি `I₁`-এ পড়ে → `x=1`, নাহলে `x=0`। (যেমন `π=0.2` হলে interval-এর ২০% হলো `I₁`।)

> 🪙 analogy: একটা অসম মুদ্রা — ২০% মাথা আসে। Uniform(0,1) draw `< 0.2` হলে "মাথা" (1), নাহলে "লেজ" (0)।

## ২. Binomial distribution (Bernoulli-র যোগফল)
`N` বার চেষ্টায় কতবার সফল — `x ∈ {0,…,N}`।
$$p(x\mid N,\pi) = \binom{N}{x}\pi^x(1-\pi)^{N-x}$$
**কীভাবে sample করি:** **`N`টা independent Bernoulli(π) draw নাও, যোগ করো।** ব্যস। (Binomial মানেই "N বার মুদ্রা ছুঁড়ে মোট মাথার সংখ্যা"।)

## ৩. Categorical distribution (একাধিক বিকল্প)
`K`টা সম্ভাব্য ফলাফলের একটা — `x ∈ {1,…,K}`, প্রতিটার probability `πₖ` (`Σπₖ=1`)।
$$p(x\mid \pi) = \pi_x$$
**কীভাবে sample করি:** interval-কে `K` ভাগে ভাগ করো, ভাগ `k`-এর আকার `≈ πₖ·|I|`। RNG যে ভাগে পড়ে, সেটাই `x`। (Bernoulli-র সাধারণ রূপ — ২-এর বদলে `K` টুকরো।)

> 🎲 analogy: একটা অসম পাশা যার ৬ মুখের probability আলাদা। `[0,1]`-কে ৬ টুকরো করো (টুকরোর মাপ = probability), draw যে টুকরোয় পড়ে সেটাই ফল।

## ৪. Poisson distribution (গণনা — কতবার ঘটল)
নির্দিষ্ট সময়ে কোনো ঘটনা কতবার ঘটল — `x ∈ {0,1,2,…}`।
$$p(x\mid\mu) = \frac{\mu^x}{x!}\exp(-\mu)$$
`μ` = গড় (mean)।
**কীভাবে sample করি (slide-এর পদ্ধতি):** Poisson-এর support অসীম, তাই **truncate** করি — কোনো বড় `K_max`-এর পর probability `0` ধরি। তারপর `K_max+1` category-র একটা **categorical** বানাই যেখানে `πₖ = p(k|μ)` আর শেষ category সব বাকিটা শুষে নেয়: `π_{K_max} = 1 − Σ`। মানে Poisson → categorical-এ নামিয়ে আনা।

> (Exercise 2-এ আরেকটা সুন্দর পদ্ধতি আছে: exponential draw যোগ করে Poisson — নিচে `03_Exercise` ফাইলে।)

## ৫. Negative Binomial distribution (ব্যর্থতার গণনা / overdispersed count)
নির্দিষ্ট সংখ্যক সাফল্য (`φ`) পাওয়ার আগে কতগুলো ব্যর্থতা — `x ∈ {0,1,…}`।
$$p(x\mid\pi,\phi)=\binom{x+\phi-1}{x}(1-\pi)^x\pi^\phi$$
**কীভাবে sample করি (integer `φ`):** Bernoulli(π) draw নিতে থাকো যতক্ষণ না `φ`টা সাফল্য আসে; মাঝে যত ব্যর্থতা, সেটাই `x`।

**Mean-parameterization:** বাস্তবে `μ` (mean) ও `φ` (overdispersion/shape) দিয়ে লেখা সুবিধাজনক। তখন `π = φ/(φ+μ)`-এ রূপান্তর করে উপরের পদ্ধতিতেই sample করি। (Poisson-এর চেয়ে এর variance বড় — "overdispersed", যা Chapter 4-এর exercise-এ লাগবে।)

---

# পর্ব B — Continuous distribution থেকে sampling

## ৬. Continuous Uniform — সেতুবন্ধন
`[a,b]`-এ সমান density: `p(x|a,b)=1/(b−a)`।
**কীভাবে পাই:** Chapter 2-এর discrete uniform `z ∈ [0,m]` নিয়ে — `y := z/m` দেয় (প্রায়) continuous Uniform(0,1); তারপর `x := (b−a)y + a` দেয় Uniform(a,b)। এই `Uniform(0,1)`-ই বাকি সব continuous sampling-এর কাঁচামাল।

## ৭. CDF আর Quantile — সর্বজনীন অস্ত্রের ভিত্তি

### CDF (Cumulative Distribution Function) — মনে করিয়ে দেওয়া
$$F(\tilde x) = P(x\le \tilde x) = \int_{-\infty}^{\tilde x} p(x)\,dx$$
মানে "`x`, `x̃`-এর চেয়ে ছোট হওয়ার probability"। `F` সবসময় `0` থেকে `1`-এ ওঠা একটা S-আকৃতির বক্ররেখা।
- CCDF (complement): `1−F(x̃) = P(x>x̃)` — উপরের লেজের probability।

### Quantile function — CDF-এর উল্টো
`Q_x(π) = F⁻¹(π)`। মানে: "যে মানের নিচে `π` ভাগ (যেমন ৯০%) data পড়ে।" CDF বলে "এই মান পর্যন্ত কত%", quantile বলে "এত% হতে গেলে কোন মান"।
- একটা সুন্দর ধর্ম: monotonic transform `f`-এর জন্য `Q_{f(x)}(π) = f(Q_x(π))` — quantile transform-এর সাথে "সঙ্গী" (equivariant)। (কারণ monotonic f ক্রম বদলায় না, তাই percentile-ও বদলায় না।)

## ৮. ⭐ Inverse-CDF Sampling — chapter-এর কেন্দ্রীয় কৌশল

### মূল উপপাদ্য (কেন কাজ করে)
**Proposition:** যদি `X`-এর CDF `F` continuous ও strictly increasing হয়, তাহলে `F(X)` হলো **Uniform(0,1)**!
(প্রমাণের সার: `Y=F(X)` ধরলে `P(Y≤y)=P(F(X)≤y)=P(X≤F⁻¹(y))=F(F⁻¹(y))=y` — যা ঠিক Uniform(0,1)-এর CDF।)

### এর উল্টোটাই হলো sampling-এর চাবি
যদি `F(X)` uniform হয়, তাহলে **উল্টো করে** — uniform নিয়ে `F⁻¹` লাগালে `X` পাওয়া যায়:
$$\pi \sim \text{Uniform}(0,1) \quad\Longrightarrow\quad Q_x(\pi)=F^{-1}(\pi) \sim p(x)$$

> 🎯 **এক বাক্যে:** *"Uniform(0,1) draw নাও, quantile function `F⁻¹` লাগাও — যেকোনো continuous distribution থেকে sample পেয়ে গেলে!"* তাত্ত্বিকভাবে সব continuous sampling এই একটা কৌশলে নেমে আসে।

### 🪜 সিঁড়ির analogy
CDF হলো একটা সিঁড়ি যা `0` থেকে `1`-এ ওঠে। তুমি `[0,1]`-এ একটা উচ্চতা random ভাবে বাছলে (uniform π), তারপর জিজ্ঞেস করলে "এই উচ্চতায় সিঁড়ির কোন ধাপ?" — সেই ধাপের `x`-মানই তোমার sample। যেখানে সিঁড়ি খাড়া (density বেশি), সেখানে বেশি sample পড়ে; যেখানে ঢালু (density কম), কম পড়ে — ঠিক যেমন হওয়া উচিত।

### উদাহরণ
- **Logistic:** CDF `π=1/(1+exp(−(x−μ)/σ))` → quantile `x=μ+σ·log(π/(1−π))`। সহজ, সরাসরি কাজ করে।
- **Normal:** CDF-এ error function `erf`, quantile-এ `erf⁻¹` — closed-form সহজ না, কিন্তু সংখ্যাগতভাবে (`qnorm`) করা যায়।

## ৯. Normal sampling-এর বিকল্প: CLT-approximation
যেহেতু Normal-এর quantile কঠিন, একটা চালাকি (Central Limit Theorem):
> অনেকগুলো i.i.d. random variable-এর যোগফল Normal-এর দিকে যায়।

`uᵢ ~ Uniform(0,1)` নিয়ে:
$$x := \frac{\left(\sum_{i=1}^{K}u_i\right) - K/2}{\sqrt{K/12}} \;\approx\; \text{Normal}(0,1)$$
(এখানে `K/2` হলো `K`টা uniform-এর mean, `K/12` হলো variance — তাই standardize করছি।) `K` বড় হলে আনুমানিক ভালো হয়; slide-এ `K=3,6,12`-এ histogram ক্রমে bell-এর কাছে যায়। `K=12` জনপ্রিয় (তখন `√(K/12)=1`, সূত্র সরল)।

## ১০. এক distribution থেকে আরেকটা — সম্পর্কের জাল
এই অংশটা viva-তে খুব আসে: **"X distribution থেকে Y কীভাবে বানাবে?"**

| Distribution | কীভাবে বানাই (সম্পর্ক) |
|---|---|
| **Lognormal** | `z~Normal(μ,σ)` নাও, `x=exp(z)`। (positive, right-skewed) |
| **χ² (chi-square, ν df)** | `z₁,…,z_ν ~ Normal(0,1)`, `x=Σ zₖ²`। (normal-এর বর্গের যোগ) |
| **Student-t (ν df)** | `u~χ²(ν)` ও `z~Normal(0,1)`, `x=μ+σ·z·√(ν/u)`। (normal-কে এলোমেলো scale দিয়ে ভাগ → মোটা লেজ) |
| **Skew-normal** | দুটো normal `z₁,z₂`; যেখানে `z₂>α z₁` সেখানে `z₁`-এর চিহ্ন উল্টে, `x=ξ+ω z₁`। (একদিকে হেলানো) |

> 💡 **কেন Student-t-এর লেজ মোটা?** কারণ তাকে একটা random (χ²) সংখ্যা দিয়ে ভাগ করা হয়; মাঝে মাঝে ভাজক ছোট হলে মান অনেক বড় হয়ে যায় → বেশি extreme value → মোটা লেজ। `ν` বড় হলে t → Normal।

## ১১. Multivariate distribution — একসাথে অনেক সম্পর্কিত variable

### Multivariate Normal (Cholesky কৌশল)
$$p(x\mid\mu,\Sigma) = \frac{1}{\sqrt{(2\pi)^K|\Sigma|}}\exp\!\left(-\tfrac12 (x-\mu)^T\Sigma^{-1}(x-\mu)\right)$$
**কীভাবে sample করি:**
1. `z₁,…,z_K ~ Normal(0,1)` (independent), vector `z` বানাও।
2. covariance `Σ`-কে **Cholesky factor** করো: `Σ = L Lᵀ`।
3. `x = μ + L z`।

> 🧩 **অন্তর্দৃষ্টি:** independent normal-গুলোর কোনো correlation নেই (গোল মেঘ)। `L` দিয়ে গুণ করলে তাদের মধ্যে ঠিক `Σ`-এর correlation বসে যায় (মেঘটা টানা-হেলানো ডিম্বাকৃতি হয়), আর `μ` কেন্দ্র সরায়। `L` যেন correlation-এর "ছাঁচ"।

### Multivariate Student-t
একইভাবে: `u~χ²(ν)`, `z~MultiNormal(0,Σ)`, তারপর `x=μ+z·√(ν/u)`। (univariate t-এর বহুমাত্রিক রূপ — সব dimension একই random scale পায়।)

---

# পর্ব C — Importance Sampling (chapter-এর মুকুট) 👑

## ১২. সমস্যা: target `p(x)` থেকে sample করা যাচ্ছে না
ধরো `E_p[f(x)] = ∫ f(x)p(x)dx` চাই, কিন্তু `p` থেকে draw নেওয়া কঠিন (বা rare-event বলে সরাসরি Monte Carlo অকেজো)। কী করব?

## ১৩. চালাকি: distribution বদলে নাও
integral-টা একটা **proposal** distribution `q` (যেখান থেকে sample করা সহজ)-এর সাপেক্ষে লিখি:
$$\mathbb{E}_{p}[f(x)] = \int f(x)\,\underbrace{\frac{p(x)}{q(x)}}_{r(x)}\,q(x)\,dx = \mathbb{E}_q[f(x)\,r(x)]$$
এখানে `r(x)=p(x)/q(x)` = **importance ratio/weight**।
মানে: `q` থেকে sample নাও, কিন্তু প্রতিটা sample-কে `r(x)` দিয়ে "ওজন" দাও — যেখানে `p` বড় কিন্তু `q` ছোট, সেখানকার sample-কে বেশি গুরুত্ব দাও।

> ⚖️ analogy: তুমি ভুল ভিড় (q) থেকে জরিপ করছ, কিন্তু জানো আসল জনসংখ্যায় (p) কোন দল কম-বেশি। প্রতিটা উত্তরকে সেই অনুপাতে weight দিয়ে সংশোধন করছ।

## ১৪. Self-normalized importance sampling (বাস্তবে যেটা ব্যবহার হয়)
**সমস্যা:** প্রায়ই `p(x)`-এর normalizing constant `Z_p` জানা থাকে না (শুধু unnormalized `p̃(x)` জানি — Bayesian-এ খুব সাধারণ)। তখন weight দিয়েই normalize করি:
$$\mathbb{E}_{p}[f(x)] \approx \frac{\sum_{s=1}^{S} f(x^{(s)})\,r(x^{(s)})}{\sum_{s=1}^{S} r(x^{(s)})} = \sum_{s=1}^{S} f(x^{(s)})\,w^{(s)}, \quad x^{(s)}\sim q$$
যেখানে normalized weight `w⁽ˢ⁾ = r(x⁽ˢ⁾) / Σ r`. (ভাগ করায় unknown constant কেটে যায় — দারুণ!)

## ১৫. Importance resampling
weight `w⁽ˢ⁾` দিয়ে সরাসরি weighted গড় না নিয়ে, **weight অনুযায়ী resample** করতে পারি (multinomial draw): বেশি weight-ওয়ালা draw বেশিবার বাছা পড়ে। ফলে resampled draw `x̃⁽ˢ⁾` প্রায় `p` থেকে আসা equal-weight sample হয়ে যায় — পরে যেকোনো হিসাবে সরাসরি ব্যবহারযোগ্য।

## ১৬. ⚠️ বড় সতর্কতা: weight অস্থির হতে পারে
যদি `q`, `p`-র সাথে ভালো না মেলে (বিশেষত `q`-এর লেজ `p`-র চেয়ে পাতলা হলে), তখন দু-একটা sample-এর weight বিশাল হয়ে যায়, বাকিগুলো প্রায় শূন্য → estimate কয়েকটা sample-এর ওপর নির্ভরশীল, variance বিস্ফোরিত। **নিয়ম: `q`-এর লেজ `p`-র চেয়ে মোটা (বা সমান) হওয়া উচিত, আর `q` যেন `p`-র সব গুরুত্বপূর্ণ অঞ্চল ঢাকে।**

> 💡 **Viva-corner:** "importance sampling কখন ভেঙে পড়ে?" → যখন proposal `q` target `p`-র সাথে খারাপ মেলে, বিশেষ করে `q`-এর লেজ পাতলা হলে — তখন weight বিস্ফোরিত হয়, কার্যকর sample সংখ্যা কমে যায়।

---

## ১৭. পুরো Chapter এক নজরে (revision card)

| কৌশল | কীসের জন্য | মূল মন্ত্র |
|---|---|---|
| Interval ভাগ | discrete (Bernoulli/Categorical) | uniform কোন টুকরোয় পড়ল |
| Bernoulli যোগ | Binomial | N বার মুদ্রা |
| Truncate→categorical | Poisson | বড় K-এর পর কেটে দাও |
| **Inverse-CDF** | যেকোনো continuous | `F⁻¹(Uniform)` |
| CLT-approx / Box–Muller | Normal | uniform যোগ |
| Distribution-সম্পর্ক | lognormal, χ², t, skew | `exp`, বর্গ-যোগ, অনুপাত |
| Cholesky `μ+Lz` | Multivariate Normal | correlation বসানো |
| **Importance sampling** | `p` থেকে পারি না | `q` থেকে নাও, `r=p/q` দিয়ে weight |
| Self-normalized | unknown `Z_p` | weight দিয়েই normalize |

---

## ১৮. নিজেকে যাচাই করো
1. uniform interval ভাগ করে Bernoulli/Categorical কীভাবে sample করো?
2. Binomial আর Poisson সহজ distribution থেকে কীভাবে বানাও?
3. Inverse-CDF sampling-এর উপপাদ্য ও পদ্ধতি বলো — কেন `F(X)` uniform হয়?
4. Normal-এর quantile কঠিন কেন? দুটো বিকল্প পদ্ধতি বলো।
5. Multivariate Normal-এ Cholesky `L` কী করে?
6. Importance sampling-এর সূত্র লেখো; `r(x)` কী; self-normalization কেন দরকার; কখন ভেঙে পড়ে?

> বিস্তারিত viva উত্তর `02_Viva_Questions_Bangla.md`-এ; Exercise 2 (exponential→Poisson, importance sampling Cauchy) বুঝে বুঝে `03_Exercise_and_Solution_Bangla.md`-এ।
