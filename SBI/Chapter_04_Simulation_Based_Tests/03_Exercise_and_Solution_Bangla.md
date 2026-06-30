# Chapter 4 — Exercise ও Solution (বাংলায় বুঝে বুঝে)

> Chapter 4-এর জন্য প্রাসঙ্গিক হলো **Exercise 3** — biochemistry PhD ছাত্রদের article production নিয়ে Poisson ও Negative Binomial regression-এ simulation-based test। মূল ফাইল `SOURCE_03_exercise.pdf`, solution `SOURCE_03_exercise_solution.ipynb`, ও data `SOURCE_bioChemists.csv` এই folder-এ আছে।
>
> ⚠️ অংশ (d) bootstrap (Chapter 5) আর (e,f) cross-validation/loss — teacher নিজেই বলেছেন এগুলো পরে lecture-এ আসবে। তাই এখানে **মূলত (a), (b), (c)**-তে মন দেব (এগুলোই Chapter 4-এর simulation-based test), আর (d) সংক্ষেপে ছোঁব Chapter 5-এর সেতু হিসেবে।

---

## 📋 প্রসঙ্গ ও data
`bioChemists.csv` — ৯১৫ জন biochem PhD ছাত্র। variable:
- `art`: শেষ ৩ বছরে প্রকাশিত article সংখ্যা (count → response)
- `phd`: PhD বিভাগের মর্যাদা (prestige)
- `ment`: mentor-এর article সংখ্যা

**Model (Poisson regression):**
$$\text{art}_i \sim \text{Poisson}(\lambda_i),\quad \lambda_i = \exp(b_0 + b_1\,\text{phd}_i + b_2\,\text{ment}_i)$$
(`exp` ব্যবহার করা হয় যাতে `λ` সবসময় positive থাকে — count-এর জন্য জরুরি।)

---

## ✅ (a) Simulation-based test: `H₀: b₁=0 vs H₁: b₁≠0`

প্রশ্ন: PhD বিভাগের prestige (`phd`)-এর কি article সংখ্যায় কোনো effect আছে?

এটা Chapter 4-এর **model-based test statistic** পদ্ধতির হুবহু প্রয়োগ:

**ধাপ ১ — আসল data-তে full model ফিট, observed statistic রাখো:**
```python
import statsmodels.formula.api as smf
import statsmodels.api as sm
import numpy as np

fit_poisson = smf.glm("art ~ phd + ment", data=df, family=sm.families.Poisson()).fit()
b0, b1, b2 = fit_poisson.params["Intercept"], fit_poisson.params["phd"], fit_poisson.params["ment"]
b1_obs = b1   # observed test statistic
```

**ধাপ ২ — H₀-এর generative model `M₀` বানাও (phd বাদ দিয়ে):**
H₀ বলছে `b₁=0`, তাই null model-এ `phd` নেই। nuisance `b₀,b₂` আসল data-র estimate-এ fix করি:
$$\text{art}_i^{(s)} \sim \text{Poisson}(\lambda_i),\quad \lambda_i = \exp(\hat b_0 + \hat b_2\,\text{ment}_i)$$

**ধাপ ৩ — বহুবার simulate, প্রতিবার full model ফিট, `b̂₁⁽ˢ⁾` রাখো:**
```python
lambda_ = np.exp(b0 + b2 * df["ment"])     # H₀ অনুযায়ী λ (phd ছাড়া)
S = 1000
b1_sim = np.zeros(S)
for s in range(S):
    art_sim = np.random.poisson(lambda_)              # H₀ থেকে নতুন response
    data_sim = df.copy(); data_sim["art"] = art_sim   # predictor আসল data-রই (fixed!)
    fit_sim = smf.glm("art ~ phd + ment", data=data_sim,
                      family=sm.families.Poisson()).fit(disp=0)
    b1_sim[s] = fit_sim.params["phd"]                  # null distribution of b1
```

**ধাপ ৪ — দুই-পাশের p-value:**
```python
p_value = np.mean(np.abs(b1_sim) >= np.abs(b1_obs))
```
**ব্যাখ্যা:** `b1_sim` হলো "phd-র সত্যিকারের কোনো effect না থাকলে b̂₁ যেমন এলোমেলো হতো" — তার distribution (`0`-কেন্দ্রিক)। observed `b1_obs` যদি এই distribution-এ বিরল হয় (`p<0.05`), তাহলে phd-র effect আছে।

> 🔑 **লক্ষ্য করো (Chapter 4-এর মূল পাঠ):** predictor (`phd`, `ment`) আমরা simulate করি **না** — শুধু response `art` simulate করি। কারণ test-টা response-এর randomness নিয়ে, design fixed।

---

## ✅ অন্তর্বর্তী ধারণা: Overdispersion `φ` (কেন (b) দরকার)

Poisson model-এর একটা কঠোর অনুমান: **mean = variance** (`Var(Yᵢ)=E(Yᵢ)=λᵢ`)। কিন্তু বাস্তব count data-তে প্রায়ই variance > mean — একে বলে **overdispersion**। মাপার statistic:
$$\hat\phi = \frac{1}{n-p}\sum_{i=1}^n \frac{(\text{art}_i-\hat\lambda_i)^2}{\hat\lambda_i}$$
- `φ̂ ≈ 1`: Poisson ঠিক আছে।
- `φ̂ > 1`: overdispersed → Poisson **variability কম আঁচ করছে** → test statistic স্ফীত, p-value ভুলভাবে ছোট (false positive ঝুঁকি)।

```python
mu_hat = fit_poisson.fittedvalues
phi = np.sum((df["art"] - mu_hat)**2 / mu_hat) / fit_poisson.df_resid
```

---

## ✅ (b) Simulation-based test: `H₀: φ=1 vs H₁: φ>1`

প্রশ্ন: data কি সত্যিই overdispersed (Poisson-এর অনুমান ভাঙছে)?

পদ্ধতি (a)-র মতোই, কিন্তু এবার test statistic = `φ̂`। **পূর্ণ Poisson model** (phd সহ) থেকে simulate করি (কারণ H₀: data সত্যিই Poisson):
```python
lambda_full = fit_poisson.fittedvalues
phi_sim = np.zeros(S)
for s in range(S):
    art_sim = np.random.poisson(lambda_full)
    data_sim = df.copy(); data_sim["art"] = art_sim
    fit_sim = smf.glm("art ~ phd + ment", data=data_sim, family=sm.families.Poisson()).fit(disp=0)
    mu_sim = fit_sim.fittedvalues
    phi_sim[s] = np.sum((art_sim - mu_sim)**2 / mu_sim) / fit_sim.df_resid

p_value_phi = np.mean(phi_sim >= phi)     # এক-পাশের (H₁: φ>1)
```
**ব্যাখ্যা:** `phi_sim` হলো "data সত্যিই Poisson হলে `φ̂` যতটা ১-এর আশেপাশে ওঠানামা করত"। observed `φ̂` যদি এর চেয়ে অনেক বড় হয় (`p` ছোট) → data overdispersed, Poisson অপর্যাপ্ত → Negative Binomial দরকার। এটাই (c)-তে নিয়ে যায়।

> 💡 এটা Chapter 4-এর **model-checking**-এর সুন্দর উদাহরণ: simulation দিয়ে যাচাই করা যে model-এর মূল অনুমান data-তে টেকে কিনা।

---

## ✅ (c) Negative Binomial regression-এ একই test (`H₀: b₁=0`)

overdispersion ধরা পড়ায়, Poisson-এর বদলে **Negative Binomial** regression ব্যবহার করি — এর একটা বাড়তি dispersion parameter `θ` আছে, যা variance-কে mean-এর চেয়ে বড় হতে দেয় (Chapter 3-এর neg-binomial মনে করো)।

```python
fit_nb = smf.glm("art ~ phd + ment", data=df,
                 family=sm.families.NegativeBinomial()).fit()
b1_obs_nb = fit_nb.params["phd"]
# তারপর (a)-র মতোই: H₀ model (phd ছাড়া) থেকে simulate, প্রতিবার NB ফিট, b̂₁ রাখো, p-value
```
**কেন উত্তর বদলায়:** Poisson variability কম আঁচ করায় test statistic স্ফীত ছিল, p-value কৃত্রিমভাবে ছোট। NB সঠিক variability ধরে, তাই p-value সাধারণত **বড়** হয় — মানে phd-র effect Poisson-এ যতটা "significant" মনে হয়েছিল, বাস্তবে ততটা শক্ত নয়। 

> 🎯 **মূল শিক্ষা:** ভুল model (overdispersion উপেক্ষা) → ভুল আত্মবিশ্বাসী উপসংহার। সঠিক model বাছা inference-এর জন্য অপরিহার্য। simulation-based test দুটোই করতে দেয় — effect পরীক্ষা **এবং** model-এর অনুমান পরীক্ষা।

---

## ✅ (d) সংক্ষেপে: `θ`-এর bootstrap distribution (Chapter 5-এর সেতু)

(d) চায় NB-র dispersion `θ`-র একটা bootstrapped distribution + 95% interval (BPI = Bootstrap Percentile Interval, CBCI = bias-corrected)। এটা আসলে **Chapter 5 (Resampling/Bootstrap)**-এর বিষয়:
- **Bootstrap ধারণা:** আসল data থেকে বারবার "replacement সহ" resample করে (একই আকারের নকল dataset), প্রতিবার `θ̂` হিসাব করো → `θ̂`-এর distribution।
- **95% BPI:** ওই distribution-এর ২.৫% ও ৯৭.৫% percentile।
- (e,f: predictive simulation দিয়ে CI, আর log-likelihood loss দিয়ে in-sample vs out-of-sample তুলনা — overfitting দেখায়; পরের lecture-এ আসবে।)

> এই অংশটা Chapter 5 করার সময় আমরা বিস্তারিত করব। এখন শুধু জেনে রাখো: **simulation-based test (H₀ থেকে simulate) আর bootstrap (data থেকে resample) — দুটো ভিন্ন কিন্তু আত্মীয় ধারণা।** পার্থক্যটা viva-তে আসতে পারে (Q নিচে)।

---

## 🎯 এই exercise থেকে viva-তে যা আসতে পারে
1. "model-based simulation test-এ predictor simulate করো না কেন?" → response-এর randomness নিয়ে test, design fixed।
2. "overdispersion কী, `φ>1` মানে কী?" → variance>mean; Poisson variability কম আঁচ করে, p-value কৃত্রিমভাবে ছোট।
3. "Poisson থেকে NB-তে গেলে p-value বাড়ল কেন?" → NB সঠিক variability ধরে, তাই effect-এর প্রমাণ কম শক্ত মনে হয়।
4. **"simulation-based test আর bootstrap-এর পার্থক্য?"** → test: **H₀-এর model** থেকে simulate (null distribution); bootstrap: **observed data** থেকে resample (estimator-এর uncertainty)। একটা H₀ ধরে, আরেকটা data ধরে।
5. "(a)-তে null model-এ phd বাদ, কিন্তু (b)-তে full model — কেন?" → (a)-র H₀ "phd-র effect নেই" (তাই বাদ), (b)-র H₀ "data Poisson" (তাই full Poisson model থেকে simulate)।

> ✍️ পরামর্শ: `SOURCE_03_exercise_solution.ipynb` নিজে run করো (`pip install statsmodels` লাগবে)। `S` বাড়িয়ে-কমিয়ে p-value কতটা স্থির হয় দেখো, আর Poisson vs NB-র p-value পাশাপাশি রেখে পার্থক্যটা চোখে দেখো।
