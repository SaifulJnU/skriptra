# Chapter 4 — Viva প্রশ্নোত্তর (Simulation-Based Tests)

> ⭐ = খুব সম্ভাব্য / গুরুত্বপূর্ণ। জোরে জোরে নিজে উত্তর দাও।

---

### ⭐ Q1. Frequentist দৃষ্টিতে probability মানে কী?
**উত্তর:** Relative frequency-র সীমা — "একই পরীক্ষা বহুবার করলে কত ভাগ ক্ষেত্রে ঘটে"। তাই "৩০% দিনে বৃষ্টি" অর্থপূর্ণ, কিন্তু "আগামীকাল ৩০% বৃষ্টি" frequentist-এ অস্বস্তিকর (একবারের ঘটনা) — ওটা বরং Bayesian বিশ্বাসের মাত্রা।

---

### ⭐ Q2. H₀ আর H₁ কী? কেন H₁ সরাসরি প্রমাণ না করে H₀ বাতিল করি?
**উত্তর:** H₁ = substantive hypothesis (যা দেখাতে চাই, যেমন effect আছে); H₀ = তার যৌক্তিক উল্টো (effect নেই)। আমরা H₀ বাতিল করি কারণ H₀ একটা **নির্দিষ্ট** দাবি — এর থেকে data simulate করা যায়। H₁ অস্পষ্ট (effect কত বড়?), তাই সরাসরি simulate করা যায় না। নির্দিষ্ট H₀-কে চ্যালেঞ্জ করাই সহজ ও পরিষ্কার।

---

### ⭐⭐ Q3. p-value-র সংজ্ঞা দাও।
**উত্তর:** H₀ সত্যি ধরে নিয়ে, observed test statistic বা তার চেয়ে আরও চরম (H₁-এর দিকে) মান পাওয়ার probability। মূল দিক: (১) H₀ ধরে নেয়, (২) নির্দিষ্ট test statistic-এর সাপেক্ষে, (৩) observed-কে H₀-এর অধীন distribution-এর সাথে তুলনা করে, (৪) "চরম"-এর দিক H₁ ঠিক করে।

---

### ⭐⭐⭐ Q4. p-value যা **নয়** — ব্যাখ্যা করো। (খুব সম্ভাব্য)
**উত্তর:** p-value "H₀ সত্যি হওয়ার probability" **নয়**, "H₁ ভুল হওয়ার probability"-ও নয়। এটা শুধু "H₀ সত্যি ধরে নিলে এত-চরম (বা তার বেশি) data দেখার probability"। ছোট p মানে data H₀-এর সাথে বেমানান, কিন্তু H₀-এর probability সম্পর্কে সরাসরি কিছু বলে না। (২০১৬-তে ASA এই ভুল বোঝাবুঝি নিয়ে আনুষ্ঠানিক বিবৃতি দেয়।)

---

### ⭐⭐ Q5. Simulation-based test-এর পুরো পদ্ধতি ধাপে ধাপে বলো।
**উত্তর:**
1. research question → `H₀` ও `H₁`।
2. test statistic `T(y)` বাছো (mean, SD, skewness, coefficient…)।
3. H₀-কে generative model `M₀`-এ লেখো (data simulate করা যায় এমন)।
4. `M₀` থেকে `S` বার dataset simulate করে প্রতিবার `T(y⁽ˢ⁾)` হিসাব করো → sampling distribution।
5. `p = (1/S)Σ𝟙(T(y⁽ˢ⁾) ≥ T(y_obs))` (এক-পাশের) — observed কতটা চরম।
6. `p<α` হলে H₀ বাতিল।

---

### ⭐ Q6. simulation-based p-value-র সূত্রটা কীভাবে Chapter 1-এর সাথে যুক্ত?
**উত্তর:** `p = (1/S)Σ𝟙(T(y⁽ˢ⁾)≥T(y_obs))` আসলে indicator function-এর গড় — মানে একটা **Monte Carlo expectation** (Chapter 1)। probability = expectation of indicator। তাই p-value মানে "H₀-এর দুনিয়ায় observed-বা-চরম ঘটনার Monte Carlo estimate"।

---

### ⭐⭐ Q7. কোন ক্ষেত্রে simulation অপরিহার্য, কোথায় লাগে না — উদাহরণ দাও।
**উত্তর:** mean-এর sampling distribution analytic-ভাবে জানা (`Normal(μ,σ/√N)`), তাই simulation লাগে না। কিন্তু **skewness**-এর finite-sample sampling distribution-এর কোনো closed-form নেই — সেখানে simulation **একমাত্র** পথ। সাধারণভাবে: জটিল test statistic বা ছোট নমুনায় analytic distribution না থাকলে simulation অপরিহার্য।

---

### ⭐⭐ Q8. Nuisance parameter কী? এটা simulation-based test-এ কী সমস্যা তৈরি করে?
**উত্তর:** Nuisance parameter (`ψ`) = এমন parameter যা আমাদের আগ্রহের বিষয় নয়, কিন্তু model-এ আছে এবং H₀ তার মান ঠিক করে দেয় না (যেমন σ অজানা)। সমস্যা: `M₀` থেকে simulate করতে সব parameter লাগে, কিন্তু nuisance-এর মান জানা নেই — তাহলে কোন মান দিয়ে simulate করব? sampling distribution তো nuisance-এর ওপর নির্ভর করে।

---

### ⭐⭐ Q9. Pivotal statistic কীভাবে nuisance সমস্যা সমাধান করে? উদাহরণ দাও।
**উত্তর:** Pivotal statistic-এর sampling distribution nuisance থেকে **স্বাধীন**। উদাহরণ: `Z=(ȳ−μ_G)/σ̂` — data-র নিজের `σ̂` দিয়ে ভাগ করায় scale বাতিল হয়, তাই `Z`-এর distribution `σ`-এর ওপর নির্ভর করে না। ফলে nuisance σ না জেনেও test করা যায়। (এটাই ক্লাসিক z/t-test-এর ভিত।)

---

### ⭐ Q10. Generative model কী? এমন model-এর উদাহরণ দাও যা generative নয়।
**উত্তর:** Generative model = যে statistical model থেকে data simulate করা যায়, `y~M(θ,ψ)`। উদাহরণ যা generative নয়: শুধু একটা discriminative/conditional সিদ্ধান্ত নিয়ম বা একটা test statistic-এর সূত্র যা data বানায় না — যেমন কেবল একটা decision boundary, যা `P(y|x)` ঠিক করলেও পূর্ণ data-generating process দেয় না। (মূল কথা: generate করতে পারলেই generative।)

---

### ⭐ Q11. Simulation-based confidence interval কীভাবে বানাও?
**উত্তর:** `M₀` থেকে simulate-করা test statistic-এর সেট `{T(y⁽ˢ⁾)}`-এর empirical quantile নাও — দুই-পাশের ৯৫% CI-এর জন্য `[Q_{2.5%}, Q_{97.5%}]`। `T(y_obs)` এই interval-এর বাইরে কিনা দেখা = দুই-পাশের p-value `<α` কিনা দেখার সমতুল্য।

---

### Q12. Model-based test statistic-এ predictor মান simulate না করে আসল data থেকে নেওয়া হয় কেন?
**উত্তর:** কারণ test-টা **response-এর randomness** নিয়ে — "H₀ সত্যি হলে এই predictor design-এ কেমন response আসত?"। predictor (design matrix) fixed ধরাই স্বাভাবিক ও বাস্তবসম্মত (আমরা ওই নির্দিষ্ট ব্যক্তিদের নিয়েই কাজ করছি)। predictor-ও simulate করলে অপ্রাসঙ্গিক বাড়তি variation ঢুকত।

---

### ⭐⭐ Q13. Simulation-based test-এর তিনটা সীমাবদ্ধতা বলো।
**উত্তর:** (১) `M₀`-এর **পূর্ণ specification** লাগে, কিন্তু sampling distribution প্রায়ই nuisance `ψ`-এর ওপর নির্ভর করে যা H₀ দেয় না। (২) **Confidence interval সংজ্ঞায়িত করা কঠিন** — H₀-তে true `θ*` ধরতে হয়, অথচ CI-এর জন্য observed data-র আলোকে যুক্তিসঙ্গত মান জানতে হয়। (৩) Model-based statistic **ব্যয়বহুল** — ধীর model `S` বার ফিট করা ভয়ানক ধীর।

---

### Q14. এই সীমাবদ্ধতাগুলো পরের কোন chapter-এর প্রেরণা দেয়?
**উত্তর:** খরচ ও nuisance সমস্যা পরে **Bayesian inference** (ch7–9) এবং বিশেষত **amortized neural inference / BayesFlow** (ch12)-এর প্রেরণা — যেখানে একবার network "train" করে বহুবার দ্রুত inference করা যায়, প্রতিবার নতুন করে হাজারবার ফিট করতে হয় না।

---

### Q15. (চিন্তা) p ≈ 0.15 (skewness) পেলে উপসংহার কী?
**উত্তর:** `p=0.15 > 0.05`, তাই H₀ (`γ=0`, symmetric) বাতিল করার মতো যথেষ্ট প্রমাণ নেই। মানে data থেকে skewness-এর শক্ত প্রমাণ পাওয়া যায়নি — কিন্তু এটা "distribution নিশ্চিত symmetric" প্রমাণ **করে না** (প্রমাণের অভাব ≠ অনুপস্থিতির প্রমাণ; ছোট N=10-এ power কম)।

---

## 🎯 দ্রুত flashcard
- frequentist prob = বারবার করলে কত ভাগ
- H₁ দেখাতে নির্দিষ্ট H₀ বাতিল করি
- **p = H₀ ধরে observed-বা-চরম পাওয়ার prob; H₀ সত্যির prob নয়**
- sim p-value = `(1/S)Σ𝟙(T(yˢ)≥T(y_obs))` (= Monte Carlo of indicator)
- skewness → analytic নেই → simulation অপরিহার্য
- nuisance → pivotal `Z=(ȳ−μ)/σ̂` দিয়ে সামলাও
- CI = sampling dist-এর quantile
- model-based: predictor fixed, response simulate
- সীমা: full spec, CI কঠিন, খরচ বেশি → Bayesian/BayesFlow-এর প্রেরণা
