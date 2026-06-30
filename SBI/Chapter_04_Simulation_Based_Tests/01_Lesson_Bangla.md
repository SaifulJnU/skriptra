# Chapter 4 — Simulation-Based Tests (বাংলায় পুরো পাঠ)

> এই chapter-এ প্রথমবার আমরা simulation দিয়ে **আসল বৈজ্ঞানিক সিদ্ধান্ত** নিই: "data-তে দেখা effect কি সত্যি, নাকি random চান্স?" মূল কৌশল: H₀ সত্যি হলে কেমন data আসত, সেটা হাজারবার simulate করে observed-এর সাথে তুলনা। একটাও ধাপ বাদ দেব না।

---

## 🎯 এক লাইনে পুরো chapter
> **"H₀ যদি সত্যি হতো — সেই দুনিয়া হাজারবার simulate করো; তোমার আসল observation সেখানে কতটা বিরল, সেটাই p-value; খুব বিরল হলে H₀ বাতিল।"**

---

## ১. Frequentist দৃষ্টিভঙ্গি — probability মানে কী?

frequentist-এ **probability = relative frequency-র সীমা** ("বারবার করলে কত ভাগ ঘটে")।
- ✅ অর্থপূর্ণ: "৩০% দিনে Dortmund-এ বৃষ্টি হয়" (বহু দিনের অনুপাত)।
- ❓ frequentist-এ অস্বস্তিকর: "আগামীকাল ৩০% probability-তে বৃষ্টি" — কারণ "আগামীকাল" একবারই ঘটে, বারবার না; এটা বরং Bayesian ভাষা (বিশ্বাসের মাত্রা)।

> এই পার্থক্যটা মনে রাখো — পুরো frequentist test "বহুবার পরীক্ষা করলে কী হতো" ধারণার ওপর দাঁড়ানো, আর সেটাই simulation দিয়ে আমরা নকল করি।

---

## ২. উদাহরণ যা পুরো chapter জুড়ে চলবে: ছাত্রদের বুদ্ধিমত্তা

**Research question:** ছাত্ররা কি গড়ে সাধারণ জনগণের চেয়ে বেশি বুদ্ধিমান?
- সাধারণ জনগণ: IQ normal, mean `μ_G=100`, SD `σ_G=15` (সংজ্ঞা অনুযায়ী)।
- ছাত্র: mean `μ_S` **অজানা** (এটাই প্রশ্ন), SD `σ_S=15` ধরা (সরলীকরণ), normal ধরা।

## ৩. Statistical hypothesis (H₀ vs H₁)

research question-কে দুই hypothesis-এ রূপ দিই:
$$H_0: \mu_S \le \mu_G \quad\text{vs.}\quad H_1: \mu_S > \mu_G$$
- `H₁` = **substantive hypothesis** (যা আমরা দেখাতে চাই — ছাত্ররা বেশি বুদ্ধিমান)।
- `H₀` = তার **যৌক্তিক উল্টো** (counter-hypothesis)।
- আমরা **H₀ বাতিল করে** H₁-কে সমর্থন করি (null hypothesis significance test, NHST)।

> 🤔 কেন সরাসরি H₁ প্রমাণ না করে H₀ বাতিল করি? কারণ "কোনো effect নেই" (H₀) একটা **নির্দিষ্ট** দাবি — এর থেকে data simulate করা যায়। কিন্তু "effect আছে" (H₁) অস্পষ্ট (কত বড় effect?)। তাই আমরা নির্দিষ্ট H₀-কে চ্যালেঞ্জ করি।

## ৪. Data ও test statistic
৫০ জন ছাত্রের নমুনায় empirical mean `ȳ_obs = 104.5` (এটাই `μ_S`-এর estimate)। আমাদের **test statistic** `T(y) = ȳ` (mean)। প্রশ্ন: `104.5` কি `100`-এর থেকে "যথেষ্ট বড়", নাকি random ওঠানামা?

## ৫. ⭐ p-value — chapter-এর আত্মা

**সংজ্ঞা (semi-formal):** p-value হলো **H₀ সত্যি ধরে নিয়ে**, observed test statistic বা তার চেয়ে **আরও চরম** (H₁-এর দিকে) মান পাওয়ার probability।

মূল ৪টা দিক (মুখস্থ রাখো):
1. p-value **H₀ সত্যি ধরে নেয়**।
2. এটা একটা **নির্দিষ্ট test statistic**-এর সাপেক্ষে।
3. এটা observed মানকে H₀-এর অধীন তার **তাত্ত্বিক distribution**-এর সাথে তুলনা করে।
4. "চরম" মানে কী, তা `H₁`-এর দিক ঠিক করে (এখানে H₁: `μ_S>μ_G`, তাই "চরম" = **বড়** মান)।

> ⚠️ **যা p-value নয়:** এটা "H₀ সত্যি হওয়ার probability" **নয়**। এটা "H₀ সত্যি হলে এমন (বা আরও চরম) data দেখার probability"। এই পার্থক্য viva-তে সোনা।

## ৬. ⭐⭐ Generative model under H₀ — এখানেই simulation ঢোকে

H₀-কে এমনভাবে লিখি যা থেকে **data বানানো যায়** (generative model `M₀`)। H₀-এর সীমানায় (`μ_S=μ_G=100`):
$$y_i \sim \text{Normal}(100, 15), \quad i=1,\ldots,N$$
এটাই H₀-এর generative model — এখান থেকে আমরা যত খুশি কৃত্রিম dataset বানাতে পারি!

### বারবার simulate করার পদ্ধতি
`S` বার (যেমন `S=10000`) পুনরাবৃত্তি করো:
1. `y_i⁽ˢ⁾ ~ Normal(100,15)` draw করো (`i=1…N`, এখানে N=50)।
2. ওই কৃত্রিম dataset-এর mean `ȳ⁽ˢ⁾` হিসাব করো।

ফল: `S`টা simulate-করা mean `{ȳ⁽¹⁾,…,ȳ⁽ˢ⁾}` — এটাই **sampling distribution of the mean under H₀**। (slide-এ এর histogram `100`-এর চারপাশে সুন্দর bell।)

## ৭. ⭐ Simulation-based p-value

observed `ȳ_obs`-কে ওই simulate-করা distribution-এর সাথে তুলনা: H₀-এর দুনিয়ায় কত ভাগ simulate-করা mean `ȳ_obs`-এর সমান বা বেশি?
$$p = \frac{1}{S}\sum_{s=1}^{S} \mathbb{1}\big(T(y^{(s)}) \ge T(y_{\text{obs}})\big)$$
মানে: **simulate-করা মানগুলোর মধ্যে কত ভাগ observed-এর চেয়ে চরম, সেটাই p-value** (indicator function-এর গড় — Chapter 1-এর "probability = expectation of indicator"-এর সরাসরি প্রয়োগ!)।

উদাহরণে `p ≈ 0.02`। মানে: H₀ সত্যি হলে `104.5` বা তার বেশি mean পাওয়া মাত্র ২% ক্ষেত্রে ঘটে — খুব বিরল। তাই H₀ বাতিল, ছাত্ররা সত্যিই বেশি বুদ্ধিমান (সাধারণত `p<0.05` হলে বাতিল)।

### সংজ্ঞাগুলো গুছিয়ে (slide থেকে)
- **Test statistic** `T(y)`: data-র যেকোনো function (যেমন mean, SD, skewness)।
- **Sampling distribution**: `y~M` হলে `T(y)`-এর distribution। simulation-based version: প্রতিটা simulate-করা dataset-এ `T(y⁽ˢ⁾)` হিসাব করে histogram।
- **Generative model**: যে statistical model থেকে data simulate করা যায়, `y ~ M(θ,ψ)` — `θ` আগ্রহের parameter, `ψ` nuisance parameter।

## ৮. কখন simulation **অপরিহার্য**: SD ও Skewness উদাহরণ

### Mean-এর ক্ষেত্রে simulation লাগে না (কিন্তু শেখায়)
mean-এর sampling distribution analytic-ভাবে জানা: `ȳ ~ Normal(μ, σ/√N)`। তাই এখানে simulation দরকার নেই — কিন্তু ধারণা বোঝাতে ভালো উদাহরণ।

### Standard Deviation (SD)
**Research question:** ছাত্রদের IQ-র SD কি `15`-এর চেয়ে ছোট? `H₀: σ≥15 vs H₁: σ<15`। `N=10`-এ observed `σ̂_obs=8.4`।
- পদ্ধতি: `Normal(100,15)` থেকে বহুবার `N=10`-এর dataset simulate করে প্রতিবার `σ̂⁽ˢ⁾` হিসাব করো → SD-র sampling distribution।
- `p = P(σ̂⁽ˢ⁾ ≤ 8.4) ≈ 0.03`। বিরল → H₀ বাতিল।
- (নোট: normality-তে variance estimator scaled χ² মানে — analytic-ও সম্ভব, কিন্তু simulation সহজ।)

### ⭐ Skewness — এখানেই simulation-এর আসল শক্তি
**Research question:** IQ distribution কি symmetric নাকি একদিকে হেলানো? `H₀: γ=0 vs H₁: γ≠0`। observed `γ̂_obs=0.84`।
$$\hat\gamma(y) = \frac{\sum(y_i-\bar y)^3/N}{\left(\sum(y_i-\bar y)^2/N\right)^{3/2}}$$
- **মূল কথা:** finite sample-এ skewness-এর sampling distribution **analytic নয়** — কোনো সূত্র নেই! তাই **simulation-ই একমাত্র পথ।**
- পদ্ধতি: `Normal(100,15)` থেকে বহুবার simulate করে `γ̂⁽ˢ⁾`-র distribution বানাও (symmetric, `0`-কেন্দ্রিক)।
- দুই-পাশের test (H₁: `≠`), তাই `p = P(|γ̂⁽ˢ⁾| ≥ |0.84|) ≈ 0.15`। বড় (>0.05) → H₀ বাতিল করা যায় না (skewness-এর শক্ত প্রমাণ নেই)।

> 💡 **এই chapter-এর মূল বিক্রয়যুক্তি:** যেখানে গাণিতিক সূত্র নেই (skewness, জটিল model), সেখানেও simulation দিয়ে test করা যায়। এটাই "simulation-based"-এর শক্তি।

## ৯. ⭐ Nuisance parameter ও Pivotal statistic (সূক্ষ্ম কিন্তু জরুরি)

**সমস্যা:** `M₀(θ,ψ)` থেকে simulate করতে **সব** parameter জানা লাগে — আগ্রহের `θ` (H₀ ঠিক করে দেয়) আর **nuisance** `ψ` (H₀ ঠিক করে দেয় না)। যেমন σ যদি অজানা হয়, তাহলে কোন σ দিয়ে simulate করব?

**চালাকি — Pivotal statistic:** এমন test statistic বানাও যার sampling distribution **nuisance থেকে স্বাধীন**। উদাহরণ:
$$Z(y) = \frac{\bar y - \mu_G}{\hat\sigma}$$
এখানে data-র নিজস্ব `σ̂` দিয়ে ভাগ করায় `Z`-এর distribution `σ_G`-এর ওপর **নির্ভর করে না** (slide দেখায়: `σ=15` ও `σ=5`-এ `ȳ`-র distribution আলাদা, কিন্তু `Z`-এর distribution প্রায় একই)।

> 🎯 **অন্তর্দৃষ্টি:** কাঁচা mean `ȳ` σ-এর ওপর নির্ভরশীল (σ বড় হলে ছড়ানো)। কিন্তু σ̂ দিয়ে ভাগ করলে scale বাতিল হয়ে যায় — তাই `Z` যেকোনো σ-তে একই আচরণ করে। এটাই pivotal-এর সৌন্দর্য: nuisance না জেনেও test করা যায়। (এই `Z`-ই ক্লাসিক t/z-test-এর ভিত।)

## ১০. Confidence Interval (CI) — simulation দিয়ে

p-value "হ্যাঁ/না" দেয়, কিন্তু আমরা প্রায়ই একটা **পরিসর** চাই। 
**সংজ্ঞা:** `[θ_l(y), θ_u(y)]` একটা (calibrated) `1−α` confidence interval, যদি `y~M₀`-এর জন্য সত্যিকারের `θ*` এই interval-এ `1−α` probability-তে পড়ে।

**Simulation-based পদ্ধতি:** simulate-করা test statistic `{T(y⁽ˢ⁾)}`-এর **empirical quantile** নাও:
- দুই-পাশের: `[Q_{α/2}, Q_{1−α/2}]` (যেমন ৯৫%-এর জন্য ২.৫% ও ৯৭.৫% quantile)।
- এক-পাশের: `[Q_0, Q_{1−α}]` বা `[Q_α, Q_1]`।
- `T(y_obs)` এই interval-এর বাইরে কিনা দেখা = দুই-পাশের p-value `>α` কিনা দেখার সমতুল্য।

> ⚠️ একটা সূক্ষ্মতা (slide-এর question): sampling-distribution interval আর confidence interval ঠিক একই জিনিস নয় — CI-এর জন্য জানতে হয় observed data-র আলোকে কোন parameter মান যুক্তিসঙ্গত, যা frequentist-এ জটিল। (এই টানাপোড়েনই পরে Bayesian credible interval-এর প্রেরণা।)

## ১১. Model-based test statistic (জটিল প্রশ্নে)

জটিল গবেষণায় test statistic নিজেই একটা **fit-করা model**। উদাহরণ — logistic regression:
$$y_i \sim \text{Bernoulli}(\pi_i),\quad \pi_i = \text{logistic}\Big(\sum_k b_k x_{ik}\Big)$$
coefficient `b_k`-র closed-form estimator নেই (MLE iterative, আর তার sampling distribution শুধু **asymptotically** normal — বড় `N`-এ)। ছোট `N`-এ? → **simulation-based test।**

**উদাহরণ (heart disease, N=20):** age (`b₁`)-এর effect test করতে — `H₀: b₁=0`. পদ্ধতি:
1. H₀-এর model `M₀` ফিট করো (age বাদ দিয়ে): `disease ~ logistic(b₀+b₂·sex)`। nuisance `b₀,b₂` আসল data-র estimate-এ fix করো।
2. `M₀` থেকে বহুবার response simulate করো (**predictor মান আসল data থেকেই নাও**, simulate করো না — কারণ আমরা response-এর randomness নিয়ে test করছি, predictor design fixed)।
3. প্রতিবার full model ফিট করে `b̂₁⁽ˢ⁾` রাখো → sampling distribution → p-value।
উদাহরণে `p ≈ 0.11` (age-এর effect পরিসংখ্যানগতভাবে স্পষ্ট নয়)।

## ১২. ⚠️ Simulation-based test-এর সীমাবদ্ধতা (slide-এর summary — viva-তে আসে)

1. **পূর্ণ specification লাগে:** `M₀`-এর সব parameter জানা লাগে, কিন্তু sampling distribution প্রায়ই nuisance `ψ`-এর ওপর নির্ভর করে, যা H₀ দেয় না। (pivotal statistic দিয়ে আংশিক সমাধান।)
2. **Confidence interval সংজ্ঞায়িত করা কঠিন:** CI-এর জন্য H₀-তে "সত্যিকারের `θ*`" ধরতে হয়, কিন্তু interval-এর জন্য জানতে হয় observed data-র আলোকে কোন মান যুক্তিসঙ্গত — এই দ্বন্দ্ব frequentist-এ অস্বস্তিকর।
3. **Model-based statistic ব্যয়বহুল:** একটা ধীর model `S=10000` বার ফিট করা ভয়ানক ধীর।

> 💡 এই সীমাবদ্ধতাগুলোই (বিশেষত খরচ ও nuisance) পরে **Bayesian inference** (ch7+) এবং **amortized neural inference/BayesFlow** (ch12)-এর প্রেরণা — যেখানে একবার "train" করে বহুবার দ্রুত inference করা যায়।

---

## ১৩. পুরো Chapter এক নজরে (revision card)

| ধারণা | মূল কথা |
|---|---|
| Frequentist probability | বারবার করলে কত ভাগ |
| H₀ vs H₁ | H₁ দেখাতে H₀ বাতিল করি |
| p-value | H₀ ধরে observed-বা-চরম পাওয়ার prob; H₀ সত্যির prob **নয়** |
| Generative model M₀ | H₀-কে data-simulate-যোগ্য রূপে লেখা |
| Sim-based p-value | `(1/S)Σ𝟙(T(yˢ)≥T(y_obs))` |
| Sampling distribution | simulate-করা `T(yˢ)`-র histogram |
| Skewness উদাহরণ | analytic নেই → simulation **অপরিহার্য** |
| Nuisance + pivotal | `Z=(ȳ−μ)/σ̂` nuisance-মুক্ত |
| Confidence interval | sampling dist-এর quantile `[Q_{α/2},Q_{1−α/2}]` |
| Model-based statistic | fit-করা coefficient; predictor fixed রাখো |
| সীমাবদ্ধতা | full spec, CI কঠিন, খরচ বেশি |

---

## ১৪. নিজেকে যাচাই করো
1. p-value-র সংজ্ঞা বলো — আর এটা যা **নয়** সেটাও।
2. simulation-based p-value-র সূত্র ও পদ্ধতি বলো (generative model থেকে শুরু করে)।
3. mean-এর জন্য simulation লাগে না, কিন্তু skewness-এর জন্য লাগে — কেন?
4. nuisance parameter কী? pivotal statistic কীভাবে সমাধান দেয়? `Z`-এর উদাহরণ।
5. simulation-based confidence interval কীভাবে বানাও?
6. model-based test-এ predictor মান simulate না করে আসল data থেকে নেওয়া হয় কেন?
7. simulation-based test-এর তিনটা সীমাবদ্ধতা বলো।

> বিস্তারিত viva উত্তর `02_Viva_Questions_Bangla.md`-এ; Exercise 3 (Poisson regression, overdispersion, negative binomial) বুঝে বুঝে `03_Exercise_and_Solution_Bangla.md`-এ।
