# Ch 3 — Suggestions: how to study the chapter that is the exam

*বাংলা সংস্করণ নিচে আছে → [বাংলায় পড়ো](#অধ্যায়-৩--কীভাবে-পড়বে-বাংলা)*

**Time budget: 10 days of your 21 (Days 3–12).**
**This chapter is ~85% of your marks. Everything else is warm-up.**

---

## The shape of the chapter

| Section | Topic | Exam weight | Difficulty |
|---|---|---|---|
| **3.1** Model Definition | assumptions, dummies, transformations | ~20% | ⭐⭐⭐ |
| **3.2** Parameter Estimation | OLS, $\hat\sigma^2$, Gauss–Markov, BLUE | ~25% | ⭐⭐⭐⭐ |
| **3.3** Testing & Intervals | t-test, F-test, $C\beta=d$, CIs | ~25% | ⭐⭐⭐⭐⭐ |
| **3.4** Model Choice | AIC, BIC, $\bar R^2$, diagnostics | ~15% | ⭐⭐⭐ |

**3.3 is the hardest and the most heavily examined.** Budget accordingly: 3 full days on it, not 1.

---

## File order

| # | File | Time | Non-negotiable outcome |
|---|---|---|---|
| 1 | `01-notes-3.1.1-model-definition.md` | 60 min | Write $\boldsymbol{y}=\boldsymbol{X}\boldsymbol\beta+\boldsymbol\varepsilon$ with all dimensions in 60 sec |
| 2 | `02-notes-3.1.2-model-assumptions.md` | 90 min | **List all assumptions from memory**, and what breaks when each fails |
| 3 | `03-notes-3.1.3-covariate-effects.md` | 90 min | Dummies, polynomials, interactions, transformations |
| 4 | `04-notes-3.2.1-ols-estimation.md` | 120 min | **Derive $\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$ on a blank page in 4 min** |
| 5 | `05-notes-3.2.2-error-variance.md` | 45 min | $\hat\sigma^2$ — ML vs unbiased, and when each is used |
| 6 | `06-notes-3.2.3-properties-gauss-markov.md` | 90 min | State Gauss–Markov with all assumptions; explain BLUE word by word |
| 7 | `07-notes-3.3-hypothesis-testing.md` | 120 min | **Build $\boldsymbol{C}$ and $\boldsymbol{d}$ for any restriction in 60 sec** |
| 8 | `08-notes-3.3.1-exact-F-test.md` | 120 min | Both F formulas; know which inputs each needs |
| 9 | `09-notes-3.3.2-confidence-prediction-intervals.md` | 90 min | CI vs prediction interval — which is wider and why |
| 10 | `10-notes-3.4.1-bias-variance.md` | 45 min | Explain the tradeoff without formulas |
| 11 | `11-notes-3.4.2-model-choice-criteria.md` | 90 min | **Compute AIC and BIC from $\hat\varepsilon'\hat\varepsilon$ and $n$** |
| 12 | `12-notes-3.4.3-practical-model-choice.md` | 30 min | Forward/backward selection, cross-validation |
| 13 | `13-notes-3.4.4-model-diagnosis.md` | 90 min | Name what each residual plot detects |
| 14 | `20-EXERCISES.md` → `21-SOLUTIONS.md` | 180 min | Then redo Sheets 3, 4, 5 |

---

## The four things you must be able to do cold

If you can do these four, you pass comfortably. If you can't, nothing else saves you.

### ① Derive OLS on a blank page

$$\text{minimise } S(\boldsymbol\beta) = (\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta) \;\Longrightarrow\; \hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$$

**Appears in:** Exam Summer 2025 Ex 4(b) [2 pts], WS 23/24 Ex 2(b) [2 pts]. The WS 23/24 marking key literally says: *"1 point for correctly stating that RSS needs to be minimized. And 1 point for correctly deriving the solution."*

**Practise it 10 times.** Timed. Blank paper. It is the single most reliable 2 marks in the paper.

### ② Fill in a missing R output

Given a regression table with holes, reconstruct estimate / std. error / t-value / residual standard error.

**Appears in:** Exam Summer 2025 Ex 3(a) [2.5 pts], Example Exam LiMo Ex 1 [most of 35 pts].

The whole skill is three relationships:
$$t = \frac{\hat\beta}{\widehat{\text{se}}},\qquad \hat\sigma^2 = \frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n-p},\qquad \hat\beta_0=\bar y-\hat\beta_1\bar x$$

### ③ Build $\boldsymbol{C}$ and $\boldsymbol{d}$, then compute $F$

Turn a verbal hypothesis into $\boldsymbol{C}\boldsymbol\beta=\boldsymbol{d}$, count the restrictions $r$, compute $F$, compare to the quantile, decide.

**Appears in:** Exam Summer 2025 Ex 3(c)+(d) [3.5 pts], Sheet 4 (all three exercises), and every other paper.

This surprises people every year. It shouldn't surprise you.

### ④ Confidence interval + test decision

$$\hat\beta_j \pm t_{n-p}(1-\alpha/2)\cdot\widehat{\text{se}}(\hat\beta_j)$$

then answer "would you reject $H_0:\beta_j = c$?"

**Appears in:** every single paper.

---

## Study tactics specific to this chapter

**Do the derivations by hand, on paper, repeatedly.** Reading a derivation and reproducing one are completely different skills, and only the second is examined. The OLS derivation should become muscle memory.

**Build your own formula sheet as you go.** Don't just use mine (`99-exam-vault/10-FORMULA-SHEET.md`) — writing it yourself *is* the encoding. Then check yours against mine.

**Do every exercise sheet twice.** Sheets 3, 4 and 5 are essentially a past exam split into pieces, using the same running `Wage` model throughout. They are the closest thing you have to the real paper.

**Always write the formula before the numbers.** The marking keys explicitly award points for method. A right number with no formula can score less than a wrong number with the right formula.

**Track the notation.** This chapter is where the $p$ vs $k$ confusion does real damage — it propagates from degrees of freedom into every quantile lookup and every test decision. Re-read `chapter-01-introduction/03-notes-1.3-notation.md` §2 if you ever hesitate.

---

## What to skip

- ❌ **Section 3.5** (Bibliographic Notes and Proofs) — outside your scope entirely.
- ❌ The full algebraic proof of the Gauss–Markov theorem. **Know the statement, the assumptions, and what BLUE means.** WS 23/24 Ex 2(a) asks you to *describe* it for 4 points, with the key noting *"1 point for every assumption"* — that's a listing question, not a proof question.
- ❌ Deriving the F-statistic's distribution from quadratic forms. Know $F\sim F_{r,\,n-p}$ under $H_0$, and why $r$ and $n-p$ are what they are.
- ❌ Ridge / lasso / boosting / Bayesian — Chapter 4, out of scope. *(One caveat: a WS 22/23 T/F mentions the ridge estimator. Know one sentence — see `32-TRAPS.md`.)*
- ❌ Memorising Mallow's $C_p$ in detail. Know it exists and penalises complexity like AIC.

---

## Self-check before you start past papers (end of Day 12)

- [ ] Derive OLS on blank paper in under 4 minutes
- [ ] List all classical linear model assumptions, with what each one buys you
- [ ] State Gauss–Markov and explain each letter of BLUE
- [ ] Write $\hat\sigma^2$ in both ML and unbiased forms, and say which one AIC uses
- [ ] Given a verbal hypothesis, produce $\boldsymbol{C}$, $\boldsymbol{d}$ and $r$ in under a minute
- [ ] Compute $F$ from SSE and SSE$_{H_0}$ **and** from $R^2$
- [ ] Compute a CI for any $\hat\beta_j$ and decide any $H_0:\beta_j=c$
- [ ] Explain why a prediction interval is wider than a confidence interval
- [ ] Compute AIC and BIC from $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}$, $n$ and $|M|$
- [ ] Say what each of the four standard R diagnostic plots detects
- [ ] Fill missing values in an R regression table

**Eleven boxes. Tick them all before Day 13's mock.** Then Week 3 is about getting *fast*, not getting *competent* — which is exactly where you want to be.

---
---

# অধ্যায় ৩ — কীভাবে পড়বে (বাংলা)

> টেকনিক্যাল শব্দ আর ফাইলের নাম ইংরেজিতেই রেখেছি — **পরীক্ষার উত্তর ইংরেজিতে লিখতে হবে**।

**সময় বরাদ্দ: তোমার ২১ দিনের মধ্যে ১০ দিন (দিন ৩–১২)।**
**এই অধ্যায়ই তোমার নম্বরের ~৮৫%। বাকি সব কেবল গা-গরম।**

---

## অধ্যায়ের গড়ন

| সেকশন | বিষয় | পরীক্ষায় ওজন | কঠিনতা |
|---|---|---|---|
| **৩.১** Model Definition | assumptions, dummies, transformations | ~২০% | ⭐⭐⭐ |
| **৩.২** Parameter Estimation | OLS, $\hat\sigma^2$, Gauss–Markov, BLUE | ~২৫% | ⭐⭐⭐⭐ |
| **৩.৩** Testing & Intervals | t-test, F-test, $C\beta=d$, CIs | ~২৫% | ⭐⭐⭐⭐⭐ |
| **৩.৪** Model Choice | AIC, BIC, $\bar R^2$, diagnostics | ~১৫% | ⭐⭐⭐ |

**৩.৩ সবচেয়ে কঠিন, আবার সবচেয়ে বেশি পরীক্ষায় আসে।** সেভাবেই সময় ভাগ করো: এর পেছনে **পুরো ৩ দিন**, ১ দিন নয়।

---

## ফাইলের ক্রম

| # | ফাইল | সময় | যে ফলাফলে ছাড় নেই |
|---|---|---|---|
| ১ | `01-notes-3.1.1-model-definition.md` | ৬০ মিনিট | সব dimension সহ $\boldsymbol{y}=\boldsymbol{X}\boldsymbol\beta+\boldsymbol\varepsilon$ ৬০ সেকেন্ডে লেখা |
| ২ | `02-notes-3.1.2-model-assumptions.md` | ৯০ মিনিট | **সব assumption মুখস্থ বলা**, আর প্রতিটা ভাঙলে কী নষ্ট হয় |
| ৩ | `03-notes-3.1.3-covariate-effects.md` | ৯০ মিনিট | Dummies, polynomials, interactions, transformations |
| ৪ | `04-notes-3.2.1-ols-estimation.md` | ১২০ মিনিট | **সাদা কাগজে ৪ মিনিটে $\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$ derive করা** |
| ৫ | `05-notes-3.2.2-error-variance.md` | ৪৫ মিনিট | $\hat\sigma^2$ — ML বনাম unbiased, কোনটা কখন লাগে |
| ৬ | `06-notes-3.2.3-properties-gauss-markov.md` | ৯০ মিনিট | সব assumption সহ Gauss–Markov বলা; BLUE-র প্রতিটা অক্ষর ব্যাখ্যা করা |
| ৭ | `07-notes-3.3-hypothesis-testing.md` | ১২০ মিনিট | **যেকোনো restriction-এর জন্য ৬০ সেকেন্ডে $\boldsymbol{C}$ আর $\boldsymbol{d}$ বানানো** |
| ৮ | `08-notes-3.3.1-exact-F-test.md` | ১২০ মিনিট | F-এর দুটো সূত্রই; কোনটার জন্য কী ইনপুট লাগে |
| ৯ | `09-notes-3.3.2-confidence-prediction-intervals.md` | ৯০ মিনিট | CI বনাম prediction interval — কোনটা চওড়া, আর কেন |
| ১০ | `10-notes-3.4.1-bias-variance.md` | ৪৫ মিনিট | সূত্র ছাড়াই tradeoff ব্যাখ্যা করা |
| ১১ | `11-notes-3.4.2-model-choice-criteria.md` | ৯০ মিনিট | **$\hat\varepsilon'\hat\varepsilon$ আর $n$ থেকে AIC ও BIC বের করা** |
| ১২ | `12-notes-3.4.3-practical-model-choice.md` | ৩০ মিনিট | Forward/backward selection, cross-validation |
| ১৩ | `13-notes-3.4.4-model-diagnosis.md` | ৯০ মিনিট | কোন residual plot কী ধরে — নাম বলা |
| ১৪ | `20-EXERCISES.md` → `21-SOLUTIONS.md` | ১৮০ মিনিট | তারপর Sheet 3, 4, 5 আবার করো |

---

## যে চারটা জিনিস ঠান্ডা মাথায় পারতেই হবে

এই চারটা পারলে তুমি আরামে পাশ। না পারলে আর কিছুই তোমাকে বাঁচাবে না।

### ① সাদা কাগজে OLS derive করা

$$\text{minimise } S(\boldsymbol\beta) = (\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta) \;\Longrightarrow\; \hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$$

**যেখানে এসেছে:** Exam Summer 2025 Ex 4(b) [২ নম্বর], WS 23/24 Ex 2(b) [২ নম্বর]। WS 23/24-এর marking key-তে আক্ষরিক লেখা আছে: *"1 point for correctly stating that RSS needs to be minimized. And 1 point for correctly deriving the solution."*

**১০ বার প্র্যাকটিস করো।** ঘড়ি ধরে। সাদা কাগজে। **পুরো পেপারে এটাই সবচেয়ে নিশ্চিত ২ নম্বর।**

### ② R output-এর ফাঁকা ঘর পূরণ করা

ফাঁকওয়ালা একটা regression table দেওয়া থাকবে — estimate / std. error / t-value / residual standard error পুনর্গঠন করতে হবে।

**যেখানে এসেছে:** Exam Summer 2025 Ex 3(a) [২.৫ নম্বর], Example Exam LiMo Ex 1 [৩৫ নম্বরের বেশিরভাগ]।

পুরো দক্ষতাটা আসলে তিনটা সম্পর্ক:
$$t = \frac{\hat\beta}{\widehat{\text{se}}},\qquad \hat\sigma^2 = \frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n-p},\qquad \hat\beta_0=\bar y-\hat\beta_1\bar x$$

### ③ $\boldsymbol{C}$ আর $\boldsymbol{d}$ বানিয়ে $F$ বের করা

কথায় লেখা hypothesis-কে $\boldsymbol{C}\boldsymbol\beta=\boldsymbol{d}$-তে রূপান্তর করো, restriction-এর সংখ্যা $r$ গোনো, $F$ বের করো, quantile-এর সাথে তুলনা করো, সিদ্ধান্ত দাও।

**যেখানে এসেছে:** Exam Summer 2025 Ex 3(c)+(d) [৩.৫ নম্বর], Sheet 4 (তিনটা exercise-ই), আর বাকি প্রতিটা পেপারে।

**প্রতি বছর এটা সবাইকে চমকে দেয়। তোমাকে যেন না দেয়।**

### ④ Confidence interval + test-এর সিদ্ধান্ত

$$\hat\beta_j \pm t_{n-p}(1-\alpha/2)\cdot\widehat{\text{se}}(\hat\beta_j)$$

তারপর উত্তর দাও: "would you reject $H_0:\beta_j = c$?"

**যেখানে এসেছে:** প্রতিটা পেপারে, ব্যতিক্রম নেই।

---

## এই অধ্যায়ের জন্য বিশেষ কৌশল

**Derivation-গুলো হাতে, কাগজে, বারবার করো।** একটা derivation **পড়া** আর সেটা **নিজে লেখা** — সম্পূর্ণ আলাদা দুটো দক্ষতা, আর পরীক্ষায় আসে কেবল দ্বিতীয়টা। OLS-এর derivation হাতের মাংসপেশিতে বসে যাওয়া উচিত।

**পড়তে পড়তেই নিজের formula sheet বানাও।** শুধু আমারটা (`99-exam-vault/10-FORMULA-SHEET.md`) ব্যবহার কোরো না — **নিজে লেখাটাই আসল মুখস্থ হওয়া**। লেখা শেষে আমারটার সাথে মিলিয়ে নাও।

**প্রতিটা exercise sheet দু'বার করো।** Sheet 3, 4 আর 5 আসলে একটা past exam-কেই টুকরো করে সাজানো, পুরোটা জুড়ে সেই একই চলমান `Wage` model। **আসল পেপারের সবচেয়ে কাছাকাছি জিনিস এগুলোই।**

**সবসময় সংখ্যার আগে সূত্রটা লেখো।** Marking key-তে স্পষ্ট করে **পদ্ধতির জন্য** নম্বর দেওয়া আছে। সূত্র ছাড়া সঠিক সংখ্যা, সঠিক সূত্র সহ ভুল সংখ্যার চেয়েও কম নম্বর পেতে পারে।

**নোটেশনের দিকে খেয়াল রাখো।** $p$ বনাম $k$-এর গোলমাল এই অধ্যায়েই আসল ক্ষতি করে — degrees of freedom থেকে শুরু হয়ে সেটা প্রতিটা quantile খোঁজায় আর প্রতিটা test-এর সিদ্ধান্তে ছড়িয়ে পড়ে। এক মুহূর্তও দ্বিধা হলে `chapter-01-introduction/03-notes-1.3-notation.md`-এর §২ আবার পড়ো।

---

## যা বাদ দেবে

- ❌ **সেকশন ৩.৫** (Bibliographic Notes and Proofs) — পুরোপুরি সিলেবাসের বাইরে।
- ❌ Gauss–Markov theorem-এর পূর্ণ বীজগাণিতিক প্রমাণ। **statement, assumption-গুলো, আর BLUE-র মানে জানলেই হবে।** WS 23/24 Ex 2(a)-তে ৪ নম্বরের জন্য এটা *describe* করতে বলা হয়েছে, key-তে লেখা *"1 point for every assumption"* — অর্থাৎ এটা **তালিকা লেখার প্রশ্ন, প্রমাণের প্রশ্ন নয়**।
- ❌ Quadratic form থেকে F-statistic-এর distribution derive করা। $H_0$-এর অধীনে $F\sim F_{r,\,n-p}$ — এটুকু আর $r$ ও $n-p$ কেন ওরকম, সেটুকুই যথেষ্ট।
- ❌ Ridge / lasso / boosting / Bayesian — অধ্যায় ৪, সিলেবাসের বাইরে। *(একটা সতর্কতা: WS 22/23-এর একটা T/F-এ ridge estimator-এর উল্লেখ আছে। এক বাক্য জেনে রাখো — `32-TRAPS.md` দেখো।)*
- ❌ Mallow's $C_p$ খুঁটিয়ে মুখস্থ করা। জেনে রাখো যে এটা আছে আর AIC-র মতোই জটিলতাকে শাস্তি দেয়।

---

## Past paper শুরু করার আগে নিজেকে যাচাই করো (দিন ১২-র শেষে)

- [ ] সাদা কাগজে ৪ মিনিটের কমে OLS derive করা
- [ ] classical linear model-এর সব assumption বলা, আর প্রতিটা তোমাকে কী এনে দেয় সেটাও
- [ ] Gauss–Markov বলা আর BLUE-র প্রতিটা অক্ষর ব্যাখ্যা করা
- [ ] $\hat\sigma^2$ ML আর unbiased — দুই রূপেই লেখা, আর AIC কোনটা ব্যবহার করে সেটা বলা
- [ ] কথায় লেখা একটা hypothesis থেকে এক মিনিটের কমে $\boldsymbol{C}$, $\boldsymbol{d}$ আর $r$ বের করা
- [ ] SSE আর SSE$_{H_0}$ থেকে **এবং** $R^2$ থেকে — দুইভাবেই $F$ বের করা
- [ ] যেকোনো $\hat\beta_j$-এর CI বের করা আর যেকোনো $H_0:\beta_j=c$-এর সিদ্ধান্ত দেওয়া
- [ ] prediction interval কেন confidence interval-এর চেয়ে চওড়া — ব্যাখ্যা করা
- [ ] $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}$, $n$ আর $|M|$ থেকে AIC ও BIC বের করা
- [ ] R-এর চারটা standard diagnostic plot-এর প্রতিটা কী ধরে — বলা
- [ ] R-এর regression table-এর ফাঁকা ঘরগুলো পূরণ করা

**এগারোটা ঘর। দিন ১৩-র mock-এর আগে সবগুলো টিক দাও।** তারপর সপ্তাহ ৩ হবে **দ্রুত** হওয়ার জন্য, **দক্ষ** হওয়ার জন্য নয় — আর ঠিক সেখানেই তোমার থাকা উচিত।
