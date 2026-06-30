# Chapter 3 — Viva প্রশ্নোত্তর (Simple Sampling)

> ⭐ = খুব সম্ভাব্য / গুরুত্বপূর্ণ। জোরে জোরে নিজে উত্তর দেওয়ার অভ্যাস করো।

---

### ⭐ Q1. discrete distribution থেকে sample করার সাধারণ কৌশল কী?
**উত্তর:** uniform integer interval `[0,m]`-কে টুকরো করে ভাগ করি, প্রতিটা টুকরোর আকার সেই category-র probability-র সমানুপাতিক। RNG যে টুকরোয় পড়ে, সেটাই sample। Bernoulli = ২ টুকরো, Categorical = `K` টুকরো।

---

### Q2. Binomial আর Poisson সহজ distribution থেকে কীভাবে বানাও?
**উত্তর:** **Binomial(N,π)** = `N`টা independent Bernoulli(π) draw-এর যোগফল। **Poisson(μ)** = support truncate করে (`K_max`-এর পর `0`) একটা `K_max+1` category-র categorical বানাই, যেখানে `πₖ=p(k|μ)` আর শেষ category বাকিটা শুষে নেয়। (বিকল্প: exponential draw যোগ করে Poisson।)

---

### ⭐⭐ Q3. Inverse-CDF (inverse transform) sampling কী এবং কেন কাজ করে?
**উত্তর:** যদি CDF `F` continuous ও strictly increasing হয়, তাহলে `F(X) ~ Uniform(0,1)`। উল্টো করে: `π~Uniform(0,1)` নিয়ে `X=F⁻¹(π)=Q(π)` দিলে `X~p(x)`। প্রমাণ: `P(F(X)≤y)=P(X≤F⁻¹(y))=F(F⁻¹(y))=y` — যা Uniform(0,1)-এর CDF। তাই Uniform draw-এ quantile function লাগালেই target থেকে sample।

---

### ⭐ Q4. CDF আর Quantile function-এর সম্পর্ক এক বাক্যে?
**উত্তর:** Quantile = CDF-এর inverse (`Q=F⁻¹`)। CDF বলে "এই মান পর্যন্ত কত শতাংশ data", quantile বলে "এত শতাংশ হতে কোন মান লাগে"।

---

### Q5. Inverse-CDF sampling-এর একটা সরল উদাহরণ দাও।
**উত্তর:** Logistic distribution। CDF `π=1/(1+exp(−(x−μ)/σ))`, তাই quantile `x=μ+σ·log(π/(1−π))`। `π~Uniform(0,1)` নিয়ে এই সূত্রে বসালেই logistic sample। (Exercise 2-এ exponential: `Q(p)=−(1/λ)ln(1−p)`।)

---

### ⭐ Q6. Normal distribution-এর quantile সরাসরি ব্যবহার করা কঠিন কেন? বিকল্প কী?
**উত্তর:** Normal-এর CDF-এ error function `erf` থাকে, quantile-এ `erf⁻¹` — closed-form সরল সূত্র নেই (সংখ্যাগত হিসাব লাগে)। বিকল্প: (১) **CLT-approx** — অনেক uniform যোগ করে standardize: `x=(Σuᵢ−K/2)/√(K/12)`; (২) **Box–Muller** — দুটো uniform থেকে নিখুঁত দুটো standard normal।

---

### ⭐⭐ Q7. একটা distribution থেকে আরেকটা বানানোর কয়েকটা সম্পর্ক বলো।
**উত্তর:**
- Lognormal: `z~Normal(μ,σ)`, `x=exp(z)`।
- χ²(ν): `x=Σₖ zₖ²` (ν টা standard normal-এর বর্গের যোগ)।
- Student-t(ν): `x=μ+σz√(ν/u)`, `z~Normal(0,1)`, `u~χ²(ν)`।
- Multivariate Normal: `x=μ+Lz`, `Σ=LLᵀ` (Cholesky)।

---

### Q8. Student-t-এর লেজ Normal-এর চেয়ে মোটা কেন?
**উত্তর:** Student-t বানানো হয় normal-কে একটা random scale `√(ν/u)` দিয়ে (`u~χ²`)। মাঝে মাঝে `u` ছোট হলে scale বড় হয়ে যায় → বেশি extreme value → মোটা লেজ। `ν` বাড়লে scale-এর variation কমে, t → Normal।

---

### ⭐ Q9. Multivariate Normal sample করতে Cholesky factor `L` কী ভূমিকা রাখে?
**উত্তর:** independent standard normal vector `z`-এর কোনো correlation নেই। `x=μ+Lz`-এ `L` (যেখানে `Σ=LLᵀ`) ঠিক `Σ`-এর correlation/covariance বসিয়ে দেয়, আর `μ` কেন্দ্র সরায়। `L` যেন correlation-এর ছাঁচ — গোল মেঘকে টানা-হেলানো ডিম্বাকৃতিতে বদলায়।

---

### ⭐⭐ Q10. Importance sampling কী এবং কখন দরকার?
**উত্তর:** target `p(x)` থেকে সরাসরি sample করা কঠিন (বা rare-event বলে সাধারণ Monte Carlo অকেজো) হলে, একটা সহজ proposal `q(x)` থেকে sample নিই, আর প্রতিটাকে weight `r(x)=p(x)/q(x)` দিয়ে সংশোধন করি:
`E_p[f] = E_q[f(x)·r(x)]`। মানে ভুল distribution থেকে নিয়ে, ওজন দিয়ে ঠিক করা।

---

### ⭐ Q11. Self-normalized importance sampling কেন দরকার?
**উত্তর:** প্রায়ই target-এর normalizing constant `Z_p` অজানা (শুধু unnormalized `p̃` জানি — Bayesian-এ সাধারণ)। তখন weight-এর যোগফল দিয়ে ভাগ করে normalize করি:
`E_p[f] ≈ Σ f(xˢ)wˢ`, যেখানে `wˢ=r(xˢ)/Σr`। ভাগ করায় unknown `Z_p` কেটে যায়।

---

### ⭐⭐ Q12. Importance sampling কখন ভেঙে পড়ে? কীভাবে এড়াবে?
**উত্তর:** যখন proposal `q` target `p`-র সাথে খারাপ মেলে — বিশেষ করে `q`-এর লেজ `p`-র চেয়ে পাতলা হলে। তখন দু-একটা sample-এর weight বিশাল, বাকিরা ~০ → variance বিস্ফোরিত, কার্যকর sample সংখ্যা (ESS) কমে যায়। এড়াতে: `q`-এর লেজ `p`-র সমান/মোটা রাখো, আর `q` যেন `p`-র সব গুরুত্বপূর্ণ অঞ্চল ঢাকে।

---

### Q13. Importance resampling কী?
**উত্তর:** weight `wˢ` দিয়ে সরাসরি weighted গড় না নিয়ে, weight অনুযায়ী draw-গুলোকে multinomial resample করা। বেশি weight-ওয়ালা বেশিবার বাছা পড়ে, ফলে resampled draw প্রায় `p` থেকে আসা equal-weight sample — পরে সরাসরি ব্যবহারযোগ্য (যেমন particle filter-এ)।

---

### Q14. (চিন্তা) tail probability `P(X>2)` Cauchy-র জন্য সাধারণ Monte Carlo-তে অসুবিধা কী, importance sampling কীভাবে সাহায্য করে?
**উত্তর:** `X>2` ঘটনা তুলনায় বিরল, তাই সাধারণ MC-তে অল্প কিছু sample ওই অঞ্চলে পড়ে → estimate-এর variance বেশি। proposal হিসেবে `2`-এর ডানে কেন্দ্রীভূত distribution (যেমন truncated Cauchy on `[2,∞)`) নিলে প্রায় সব sample প্রাসঙ্গিক অঞ্চলে পড়ে → অনেক কম variance-এ নির্ভুল estimate। (Exercise 2-এর মূল শিক্ষা।)

---

### Q15. এই chapter পুরো কোর্সে কীভাবে কাজে লাগে?
**উত্তর:** সব simulation-এর draw এখান থেকেই আসে। generative model (ch4, ch6) data simulate করতে এই sampling লাগে; importance sampling পরে Bayesian posterior approximation, ABC (ch10), model comparison-এ কেন্দ্রীয়; Cholesky-ভিত্তিক multivariate sampling correlated parameter-এ লাগে।

---

## 🎯 দ্রুত flashcard
- discrete: uniform interval টুকরো করো
- Binomial = ΣBernoulli; Poisson = truncated categorical
- **inverse-CDF: `F⁻¹(Uniform)` → যেকোনো continuous**; কারণ `F(X)~Uniform`
- Normal quantile কঠিন → CLT-approx / Box–Muller
- lognormal=exp(N), χ²=Σz², t=N·√(ν/χ²), MVN=μ+Lz (Cholesky)
- **importance sampling: `q` থেকে নাও, weight `r=p/q`**; self-normalize (unknown Z); লেজ পাতলা হলে ভেঙে পড়ে
