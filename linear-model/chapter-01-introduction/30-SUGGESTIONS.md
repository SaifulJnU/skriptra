# Ch 1 — Suggestions: how to study this chapter

*বাংলা সংস্করণ নিচে আছে → [বাংলায় পড়ো](#অধ্যায়-১--কীভাবে-পড়বে-বাংলা)*

**Time budget: 60–90 minutes total. Once. Never again except the summary.**

Chapter 1 is worth roughly 1–2 marks on the exam, and those marks are always *disguised* — they show up as "choose an appropriate response variable" or "why is this a regression problem" inside a bigger question. Nobody has ever been asked "describe the Galton data" on this exam.

## What to actually do

1. **Read `01-notes-1.1-examples-of-applications.md` fast.** You are collecting vocabulary, not knowledge. Response, covariate, systematic component, random error, continuous vs binary vs count response.
2. **Read `02-notes-1.2-first-steps.md` properly.** This is the only part with exam value: scatter plots, what correlation does and doesn't tell you, and the idea that you *look at the data before you model it*. Diagnostics in Chapter 3.4 are the same idea wearing a suit.
3. **Read `03-notes-1.3-notation.md` twice, and write the notation table out by hand.** This is the highest-value 15 minutes in the entire chapter. Every formula in Chapter 3 is built from these symbols. If `x_i` vs `x_j` vs `x_ij` ever confuses you mid-exam, you lose minutes you don't have.
4. **Do `20-EXERCISES.md`** (15 min), check `21-SOLUTIONS.md`.
5. **Read `52-STORY-FOR-A-BABY.md`.** Then close everything and tell the story out loud.

## What to skip

- ❌ The specific datasets (Munich rent, credit scoring, malnutrition in Zambia). You need to *recognise* that regression applies to rent, credit default, and child growth — you do not need any numbers.
- ❌ The Galton historical detail. Know one sentence: *Galton invented "regression" describing children's heights regressing toward the mean.* That's the entire examinable content.
- ❌ Anything about the R packages or the book's website.

## The one thing to carry forward

Chapter 1's real job is to plant this sentence in your head:

> **observed = systematic + random**
> **y = f(x) + ε**

Everything for the next 100 pages is: *what shape is f, and what do we assume about ε?* When you hit a wall later, come back to that line.

## Self-check before moving on

Can you, without notes:

- [ ] Say what a response variable and a covariate are, and give an example of each?
- [ ] Say why we need the error term ε at all?
- [ ] Distinguish continuous / binary / count responses and name a model for each?
- [ ] Write `y_i = β₀ + β₁x_i + ε_i` and say what the index `i` ranges over?
- [ ] Say what a scatter plot is for, and one thing it can reveal that a correlation coefficient hides?

Five yeses → go to Chapter 2. Don't linger here.

---
---

# অধ্যায় ১ — কীভাবে পড়বে (বাংলা)

> টেকনিক্যাল শব্দ আর ফাইলের নাম ইংরেজিতেই রেখেছি — **পরীক্ষার উত্তর ইংরেজিতে লিখতে হবে**।

**সময় বরাদ্দ: মোট ৬০–৯০ মিনিট। একবারই। এরপর শুধু `10-SUMMARY.md`, আর কিছু না।**

অধ্যায় ১-এর দাম পরীক্ষায় মোটামুটি **১–২ নম্বর**, আর সেই নম্বরগুলো সবসময় **ছদ্মবেশে** আসে — বড় কোনো প্রশ্নের ভেতরে "choose an appropriate response variable" বা "why is this a regression problem" হিসেবে। "Galton-এর ডেটা বর্ণনা করো" — এমন প্রশ্ন এই পরীক্ষায় কোনোদিন আসেনি।

## আসলে যা করবে

1. **`01-notes-1.1-examples-of-applications.md` দ্রুত পড়ো।** এখানে তুমি **জ্ঞান** সংগ্রহ করছ না, **শব্দভাণ্ডার** সংগ্রহ করছ: response, covariate, systematic component, random error, আর continuous / binary / count response-এর পার্থক্য।
2. **`02-notes-1.2-first-steps.md` মন দিয়ে পড়ো।** পুরো অধ্যায়ে পরীক্ষার দিক থেকে দামি অংশ কেবল এটাই: scatter plot, correlation কী বলে **আর কী বলে না**, এবং মূল ধারণা — **মডেল বানানোর আগে ডেটার দিকে তাকাও**। অধ্যায় ৩.৪-এর diagnostics ঠিক এই একই ধারণা, শুধু স্যুট-টাই পরে এসেছে।
3. **`03-notes-1.3-notation.md` দু'বার পড়ো, আর নোটেশন টেবিলটা হাতে লিখে ফেলো।** পুরো অধ্যায়ের সবচেয়ে দামি ১৫ মিনিট এটাই। অধ্যায় ৩-এর প্রতিটা সূত্র এই প্রতীকগুলো দিয়েই বানানো। পরীক্ষার মাঝখানে `x_i`, `x_j` আর `x_ij` গুলিয়ে গেলে এমন মিনিট নষ্ট হবে যা তোমার হাতে নেই।
4. **`20-EXERCISES.md` করো** (১৫ মিনিট), তারপর `21-SOLUTIONS.md` মিলিয়ে দেখো।
5. **`52-STORY-FOR-A-BABY.md` পড়ো।** তারপর সব বন্ধ করে **গল্পটা মুখে বলো**।

## যা বাদ দেবে

- ❌ নির্দিষ্ট ডেটাসেটগুলো (Munich rent, credit scoring, Zambia-র অপুষ্টি)। তোমাকে শুধু **চিনতে** হবে যে বাড়িভাড়া, ঋণখেলাপি আর শিশুর বৃদ্ধি — এগুলোতে regression খাটে। **কোনো সংখ্যা মুখস্থ করার দরকার নেই।**
- ❌ Galton-এর ঐতিহাসিক খুঁটিনাটি। এক বাক্য জানলেই যথেষ্ট: *সন্তানের উচ্চতা গড়ের দিকে সরে আসা বোঝাতে গিয়ে Galton "regression" শব্দটা চালু করেন।* পরীক্ষায় আসার মতো বিষয়বস্তু এইটুকুই।
- ❌ R প্যাকেজ বা বইয়ের ওয়েবসাইট সংক্রান্ত কিছুই।

## যে একটা জিনিস সাথে নিয়ে যাবে

অধ্যায় ১-এর আসল কাজ হলো তোমার মাথায় এই বাক্যটা গেঁথে দেওয়া:

> **observed = systematic + random**
> **y = f(x) + ε**

পরের ১০০ পৃষ্ঠার পুরোটাই আসলে দুটো প্রশ্ন: ***f-এর আকৃতি কী, আর ε সম্পর্কে আমরা কী ধরে নিচ্ছি?*** পরে কোথাও আটকে গেলে এই লাইনটায় ফিরে এসো।

## এগোনোর আগে নিজেকে যাচাই করো

নোট না দেখে তুমি কি পারবে —

- [ ] response variable আর covariate কী বলতে, আর দুটোরই একটা করে উদাহরণ দিতে?
- [ ] error term ε আদৌ কেন লাগে সেটা বলতে?
- [ ] continuous / binary / count response আলাদা করতে, আর প্রতিটার জন্য একটা করে মডেলের নাম বলতে?
- [ ] `y_i = β₀ + β₁x_i + ε_i` লিখতে, আর `i` কোন পর্যন্ত চলে সেটা বলতে?
- [ ] scatter plot কী কাজে লাগে, আর correlation coefficient যা লুকিয়ে ফেলে এমন একটা জিনিস scatter plot কীভাবে দেখায় — বলতে?

**পাঁচটাই "হ্যাঁ" → সোজা অধ্যায় ২-এ যাও। এখানে ঝুলে থেকো না।**
