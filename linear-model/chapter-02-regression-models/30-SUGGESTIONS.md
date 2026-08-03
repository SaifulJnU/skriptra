# Ch 2 — Suggestions: how to study this chapter

*বাংলা সংস্করণ নিচে আছে → [বাংলায় পড়ো](#অধ্যায়-২--কীভাবে-পড়বে-বাংলা)*

**Time budget: 4–5 hours (Days 1–2 of your plan).**
**Scope: Sections 2.1, 2.2 and 2.3 only.** Sections 2.4–2.10 are outside your reading course — skip them entirely.

---

## Why this chapter matters more than its page count suggests

Chapter 2 is the book's *tour* chapter: it shows you the models without the machinery. That makes it easy to underrate. Don't. Two things here are directly examined and both are cheap marks:

1. **Section 2.2 gives you the interpretation skills.** Slope, intercept, dummy variables, and what "holding others fixed" means. The exercise sheets (Sheet 1, Sheet 2, Sheet 5) are almost entirely interpretation, and interpretation questions appear in every past paper. These are the fastest marks on the whole exam.

2. **Section 2.3 (the Logit model) is guaranteed to appear.** Look at your papers:
   - Exam Summer 2025, Ex 1(h): a T/F on interpreting logit coefficients
   - Exam Summer 2025, Ex 4(a): *"Explain why a linear regression model is not appropriate for a binary dependent variable"* — 1 full point
   - It's in the WS papers too

   You are **not** asked to derive the logit model, estimate it, or do maximum likelihood for it. You are asked two questions only: *why not linear?* and *what does $\hat\beta_j$ mean?* Learn those two answers word-perfect and you cannot lose these marks.

---

## The order to work in

| Step | File | Time | Goal |
|---|---|---|---|
| 1 | `01-notes-2.1-introduction.md` | 20 min | The general framing: $E(y\mid\boldsymbol{x})$ |
| 2 | `02-notes-2.2.1-simple-linear-regression.md` | 45 min | Interpret slope and intercept perfectly |
| 3 | `03-notes-2.2.2-multiple-linear-regression.md` | 90 min | **Dummy variables.** This is the heaviest section here |
| 4 | `04-notes-2.3-logit-model.md` | 60 min | Two answers, word-perfect |
| 5 | `20-EXERCISES.md` → `21-SOLUTIONS.md` | 60 min | Then redo Sheets 1 and 2 |
| 6 | `10-SUMMARY.md`, `40-MIND-MAP.md`, `52-STORY-FOR-A-BABY.md` | 30 min | Lock it in |

---

## The single most important thing in this chapter

**Dummy variable coding.** If you learn one thing from Chapter 2, learn this:

> A categorical covariate with **$c$ levels** becomes **$c-1$ dummy variables**. One level is left out and becomes the **reference category**. Each dummy's coefficient is the difference *from the reference*, holding everything else fixed.

Why $c-1$ and not $c$? Because including all $c$ dummies plus an intercept makes the design matrix **singular** — the dummies sum to the intercept column. No unique OLS solution. This is called the **dummy variable trap**, and it is the reason Chapter 3's full-rank assumption exists.

**Past-paper evidence this is examined every single year:**
- *Linear_model_exam_sheet*, Block I(iv): "*k* levels ⟹ *k*−1 dummies" → TRUE
- *WS 23/24*, Block I(iv): "for *m* categories we need *m* dummies" → **FALSE**
- *Exam Summer 2025*, Ex 2(a): build a wage model with education (3 levels) and birthplace (2 levels) — 3 points
- *Sheet 1*, Ex 2: define dummies for 5 education levels, identify the reference category

That's four different papers testing one idea. Learn it once, cash it every year.

---

## What to skip

- ❌ Sections 2.4 (mixed models), 2.5 (nonparametric), 2.6 (additive), 2.7 (GAM), 2.8 (geoadditive), 2.9 (quantile/GAMLSS), 2.10 (nutshell summary of all of them). **Out of scope.**
- ❌ Any *estimation* of the logit model — no maximum likelihood, no Newton–Raphson, no derivations. Chapter 5 territory.
- ❌ The probit model beyond one sentence ("same idea, uses the normal CDF instead of the logistic — very similar results").
- ❌ The book's specific datasets and numbers.

---

## Self-check before Chapter 3

Can you, cold, with no notes:

- [ ] Write $E(y\mid\boldsymbol{x}) = \boldsymbol{x}'\boldsymbol\beta$ and say why the model is about the **mean** of $y$?
- [ ] Interpret $\hat\beta_1$ in a simple regression, in one sentence, with units?
- [ ] Say when $\hat\beta_0$ should **not** be interpreted, and why?
- [ ] Take "education has 5 levels" and write down the dummies, name the reference category, and interpret two coefficients?
- [ ] Compute the wage gap between two hypothetical people from a fitted model (Sheet 1, Ex 2(c) and 2(d))?
- [ ] Write the model with an interaction term and explain **geometrically** what the interaction does?
- [ ] Give **three** reasons a linear model fails for binary $y$?
- [ ] Write the logit model both ways — as $P(y=1)$ and as $\log\frac{p}{1-p}$?
- [ ] Say exactly what $\hat\beta_j$ means in a logit model, and what it does **not** mean?

Nine yeses → Chapter 3. Anything less, go back to that specific file.

---

## A warning about pace

Chapters 1 and 2 together are ~15% of the exam. Chapter 3 is ~85%. If you are on Day 4 and still in Chapter 2, you are behind — move on, and come back to Chapter 2's interpretation drills during Week 3 when you're revising. **Chapter 3 needs your freshest attention, not your leftovers.**

---
---

# অধ্যায় ২ — কীভাবে পড়বে (বাংলা)

> টেকনিক্যাল শব্দ আর ফাইলের নাম ইংরেজিতেই রেখেছি — **পরীক্ষার উত্তর ইংরেজিতে লিখতে হবে**।

**সময় বরাদ্দ: ৪–৫ ঘণ্টা (তোমার প্ল্যানের দিন ১–২)।**
**পরিধি: শুধু সেকশন ২.১, ২.২ আর ২.৩।** সেকশন ২.৪–২.১০ তোমার reading course-এর বাইরে — **পুরোপুরি বাদ**।

---

## পৃষ্ঠাসংখ্যা দেখে এই অধ্যায়কে ছোট ভেবো না

অধ্যায় ২ হলো বইয়ের **ঘুরে দেখানোর** অধ্যায় — যন্ত্রপাতি ছাড়াই মডেলগুলো দেখিয়ে দেয়। এজন্যই একে কম দাম দেওয়া সহজ। **দিও না।** এখানে দুটো জিনিস সরাসরি পরীক্ষায় আসে, আর দুটোই **সস্তায় পাওয়া নম্বর**:

1. **সেকশন ২.২ তোমাকে interpretation-এর দক্ষতা দেয়।** Slope, intercept, dummy variables, আর "holding others fixed" কথাটার মানে। Exercise sheet-গুলো (Sheet 1, Sheet 2, Sheet 5) প্রায় পুরোটাই interpretation, আর প্রতিটা past paper-এ interpretation-এর প্রশ্ন আছে। **পুরো পরীক্ষায় সবচেয়ে দ্রুত পাওয়া নম্বর এগুলোই।**

2. **সেকশন ২.৩ (Logit model) আসবেই — নিশ্চিত।** নিজের পেপারগুলো দেখো:
   - Exam Summer 2025, Ex 1(h): logit coefficient interpret করা নিয়ে একটা T/F
   - Exam Summer 2025, Ex 4(a): *"Explain why a linear regression model is not appropriate for a binary dependent variable"* — পুরো ১ নম্বর
   - WS-এর পেপারগুলোতেও আছে

   তোমাকে logit model **derive** করতে, **estimate** করতে বা এর maximum likelihood করতে **বলা হবে না**। মাত্র দুটো প্রশ্নই করা হয়: *linear কেন নয়?* আর *$\hat\beta_j$-এর মানে কী?* এই দুটো উত্তর **অক্ষরে অক্ষরে** শিখে রাখলে এই নম্বর হারানো অসম্ভব।

---

## যে ক্রমে কাজ করবে

| ধাপ | ফাইল | সময় | লক্ষ্য |
|---|---|---|---|
| ১ | `01-notes-2.1-introduction.md` | ২০ মিনিট | সাধারণ কাঠামো: $E(y\mid\boldsymbol{x})$ |
| ২ | `02-notes-2.2.1-simple-linear-regression.md` | ৪৫ মিনিট | Slope আর intercept নিখুঁতভাবে interpret করা |
| ৩ | `03-notes-2.2.2-multiple-linear-regression.md` | ৯০ মিনিট | **Dummy variables.** এখানকার সবচেয়ে ভারী অংশ |
| ৪ | `04-notes-2.3-logit-model.md` | ৬০ মিনিট | দুটো উত্তর, অক্ষরে অক্ষরে |
| ৫ | `20-EXERCISES.md` → `21-SOLUTIONS.md` | ৬০ মিনিট | তারপর Sheet 1 আর 2 আবার করো |
| ৬ | `10-SUMMARY.md`, `40-MIND-MAP.md`, `52-STORY-FOR-A-BABY.md` | ৩০ মিনিট | পাকা করে নাও |

---

## এই অধ্যায়ের সবচেয়ে গুরুত্বপূর্ণ জিনিস

**Dummy variable coding।** অধ্যায় ২ থেকে যদি একটাই জিনিস শেখো, এটাই শেখো:

> $c$ **সংখ্যক level**-এর একটা categorical covariate হয়ে যায় $c-1$ **টা dummy variable**। একটা level বাদ পড়ে, সেটাই হয় **reference category**। প্রতিটা dummy-র coefficient হলো **reference থেকে পার্থক্য**, বাকি সব স্থির রেখে।

$c$ কেন, $c-1$ কেন নয়? কারণ সবগুলো $c$ dummy **এবং** intercept একসাথে রাখলে design matrix **singular** হয়ে যায় — dummy-গুলো যোগ হয়ে intercept column-টাই বানিয়ে ফেলে। তখন OLS-এর কোনো **unique** সমাধান থাকে না। একেই বলে **dummy variable trap**, আর এই কারণেই অধ্যায় ৩-এ full-rank assumption-টা আছে।

**প্রতি বছর এটা আসে — past paper-এর প্রমাণ:**
- *Linear_model_exam_sheet*, Block I(iv): "*k* levels ⟹ *k*−1 dummies" → **TRUE**
- *WS 23/24*, Block I(iv): "for *m* categories we need *m* dummies" → **FALSE**
- *Exam Summer 2025*, Ex 2(a): education (৩ level) আর birthplace (২ level) দিয়ে wage model বানাও — ৩ নম্বর
- *Sheet 1*, Ex 2: ৫টা education level-এর dummy লেখো, reference category চিহ্নিত করো

**চারটা আলাদা পেপার, একটাই ধারণা পরীক্ষা করছে।** একবার শেখো, প্রতি বছর নম্বর তোলো।

---

## যা বাদ দেবে

- ❌ সেকশন ২.৪ (mixed models), ২.৫ (nonparametric), ২.৬ (additive), ২.৭ (GAM), ২.৮ (geoadditive), ২.৯ (quantile/GAMLSS), ২.১০ (সবগুলোর সংক্ষিপ্তসার)। **সিলেবাসের বাইরে।**
- ❌ Logit model-এর **estimation** সংক্রান্ত যেকোনো কিছু — maximum likelihood নয়, Newton–Raphson নয়, derivation নয়। ওটা অধ্যায় ৫-এর এলাকা।
- ❌ Probit model — এক বাক্যের বেশি নয় ("একই ধারণা, logistic-এর বদলে normal CDF ব্যবহার করে — ফলাফল প্রায় একই")।
- ❌ বইয়ের নির্দিষ্ট ডেটাসেট আর সংখ্যাগুলো।

---

## অধ্যায় ৩-এ যাওয়ার আগে নিজেকে যাচাই করো

নোট না দেখে, ঠান্ডা মাথায় তুমি কি পারবে —

- [ ] $E(y\mid\boldsymbol{x}) = \boldsymbol{x}'\boldsymbol\beta$ লিখতে, আর মডেলটা কেন $y$-এর **গড়** নিয়ে সেটা বলতে?
- [ ] Simple regression-এ $\hat\beta_1$ এক বাক্যে, **একক সহ** interpret করতে?
- [ ] $\hat\beta_0$ কখন interpret করা **উচিত নয়** আর কেন — বলতে?
- [ ] "education-এর ৫টা level" থেকে dummy-গুলো লিখতে, reference category-র নাম বলতে, আর দুটো coefficient interpret করতে?
- [ ] Fitted model থেকে দুজন কাল্পনিক মানুষের wage-এর পার্থক্য বের করতে (Sheet 1, Ex 2(c) আর 2(d))?
- [ ] Interaction term সহ মডেল লিখতে, আর interaction **জ্যামিতিকভাবে** কী করে সেটা ব্যাখ্যা করতে?
- [ ] Binary $y$-এর জন্য linear model কেন ব্যর্থ হয় — **তিনটা** কারণ বলতে?
- [ ] Logit model দুইভাবেই লিখতে — $P(y=1)$ হিসেবে আর $\log\frac{p}{1-p}$ হিসেবে?
- [ ] Logit model-এ $\hat\beta_j$-এর মানে ঠিক কী, আর কী **নয়** — বলতে?

**নয়টাই "হ্যাঁ" → অধ্যায় ৩।** এর কম হলে ঠিক সেই ফাইলটায় ফিরে যাও।

---

## গতি নিয়ে একটা সতর্কবার্তা

অধ্যায় ১ আর ২ মিলে পরীক্ষার **~১৫%**। অধ্যায় ৩ একাই **~৮৫%**। দিন ৪-এ এসেও যদি তুমি অধ্যায় ২-এ আটকে থাকো, **তুমি পিছিয়ে আছ** — এগিয়ে যাও, আর অধ্যায় ২-এর interpretation drill-গুলো সপ্তাহ ৩-এর রিভিশনে ফিরে এসে করো। **অধ্যায় ৩ তোমার সবচেয়ে তরতাজা মনোযোগ চায়, উচ্ছিষ্ট নয়।**
