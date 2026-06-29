# Simulation-Based Inference (SBI) — মাস্টার স্টাডি প্ল্যান

> **লক্ষ্য:** ১২টা টপিক পুরোপুরি বুঝে, presentation আর তার পরের প্রশ্নোত্তরে (viva) এমনভাবে উত্তর দেওয়া যেন teacher সন্তুষ্ট হয় এবং তুমি ১০০% মার্ক + ৯ ECTS পাও।

---

## ০. সবচেয়ে আগে: পরীক্ষা আসলে কীভাবে হয়? (এটা না বুঝলে পুরো প্ল্যান ভুল হবে)

কোর্সের official document দুটো (`00_organization.pdf`, `sbi_project_info_2026.pdf`) আমি পড়ে দেখেছি। সত্যিকারের নিয়মগুলো হলো:

- **কোনো classical exam নাই।** পুরো মার্ক আসে একটা **project** থেকে।
- Project = একটা real-world data analysis, যেখানে **BayesFlow** library দিয়ে **Amortized Bayesian Inference (ABI)** করতে হবে।
- দল হয় **১–৩ জনের** (teacher রা ২ জনের দল সবচেয়ে বেশি recommend করে)।
- **মার্ক ভাগ: ৫০% presentation + ৫০% report।**
- Presentation: group প্রতি একটাই, ১০ মিনিট (১–২ জন) বা ১২ মিনিট (৩ জন)। সময় শেষ হলে থামিয়ে দেওয়া হবে।
- **Presentation-এর পরে teacher প্রশ্ন করবে** — শুধু তোমার project নিয়ে না, **পুরো SBI কোর্সের যেকোনো টপিক (১–১২) নিয়েও** প্রশ্ন করতে পারে। ভালো উত্তর দিলে মার্ক **বাড়বে**, না পারলে মার্ক **কমবে**।
- ⚠️ **সবচেয়ে গুরুত্বপূর্ণ লাইন (document থেকে হুবহু):** *"If you fail the presentation, you will fail the entire course and then you don't need to submit the report."* — মানে presentation pass না করলে পুরো কোর্স fail। তাই presentation + প্রশ্নোত্তরই আসল যুদ্ধক্ষেত্র।

**এর মানে তোমার কৌশল ঠিক আছে:** শুধু নিজের ৪টা chapter না, **১২টা chapter-ই** ভালো করে বুঝতে হবে, কারণ প্রশ্ন যেকোনো জায়গা থেকে আসতে পারে। আমরা ঠিক সেটাই করবো।

### গুরুত্বপূর্ণ তারিখগুলো (২০২৬)
| তারিখ | কী |
|---|---|
| June 28 (midnight) | দল ও topic registration-এর শেষ তারিখ |
| In July | lecture/exercise-এ project নিয়ে প্রশ্ন করার সুযোগ |
| **July 19 (midnight)** | presentation slides জমা |
| **July 20–24** | presentation সপ্তাহ (তুমি শুধু নিজের session-এ থাকতে পারবে) |
| **August 23 (midnight)** | report জমা |

> মনে রেখো: এখনো teacher project topic দেয়নি। যেই মুহূর্তে topic পাবে, আমাকে জানাবে — তখন আমরা ওই topic-এর জন্য আলাদা গভীর প্ল্যান বানাবো (BayesFlow code সহ)। ততদিন আমরা ১২টা chapter-এর ভিত শক্ত করবো।

---

## ১. ১২টা টপিক — file-এর সাথে ম্যাপিং

Course-এর slide-এ ১২টা "major topic" দেওয়া আছে। file-গুলোর নম্বর একটু এদিক-ওদিক (১১ নম্বর নাই, কিন্তু টপিক ১২টাই)। নিচে পুরো ম্যাপিং:

| # | টপিক | ফাইল | তোমার ভাগ |
|---|---|---|---|
| 1 | Reasons for using simulations (Introduction) | `01_introduction_simulation_based_inference.pdf` | **তুমি** |
| 2 | Pseudo-randomness with computers | `02_pseudo_randomness.pdf` | **তুমি** |
| 3 | Sampling from simple distributions | `03_simple_sampling.pdf` | **তুমি** |
| 4 | Simulation-based tests | `04_simulation_based_tests.pdf` | **তুমি** |
| 5 | Resampling methods | `05_resampling.pdf` | teammate |
| 6 | Frequentist simulation studies | `06_simulation_studies_frequentist.pdf` | teammate |
| 7 | Introduction to Bayes | `07_introduction_bayes.pdf` | teammate |
| 8 | Likelihood-based Bayesian inference | `08_likelihood_based_bayes.pdf` | teammate |
| 9 | Simulation-based verification of Bayes (SBC) | `09_simulation_studies_bayes.pdf` | teammate |
| 10 | Approximate Bayesian Computation (ABC) | `10_approximate_bayesian_computation.pdf` | teammate |
| 11 | Neural networks & generative neural models | `12_deep_learning_v1.pdf` | teammate |
| 12 | Neural Bayesian Inference (NPE) | `13_neural_posterior_estimation.pdf` | teammate |

**তোমার পরিকল্পনা:** chapter ১–৪ আগে শেষ করবে (তোমার নিজের অংশ), তারপর ধীরে ধীরে ৫–১২। আমি ঠিক এই ক্রমেই বানাবো।

### একটা বড় ছবি (পুরো কোর্স কীভাবে একটা গল্প)
পুরো কোর্স আসলে একটা সিঁড়ি:
1. **কেন simulation?** (ch1) → integral/expectation হাতে হিসাব করা কঠিন, তাই আমরা random draw দিয়ে আনুমানিক করি (Monte Carlo)।
2. **Random কোথা থেকে আসে?** (ch2) → computer তো deterministic, তাই "pseudo-random" বানাতে হয়।
3. **কোনো নির্দিষ্ট distribution থেকে draw কীভাবে?** (ch3) → inverse-CDF, rejection, importance sampling।
4. **Draw দিয়ে hypothesis test** (ch4) এবং **uncertainty estimate** (ch5 resampling/bootstrap)।
5. **আমাদের method ঠিক কাজ করছে কিনা যাচাই** (ch6 frequentist simulation studies)।
6. এরপর **Bayesian দুনিয়া**: Bayes কী (ch7) → likelihood-based Bayes/MCMC (ch8) → Bayesian method যাচাই/SBC (ch9) → likelihood ছাড়া inference/ABC (ch10)।
7. শেষে **neural network** (ch11) দিয়ে ABC-কে replace করে **Neural Posterior Estimation** (ch12) — এটাই **BayesFlow** ও তোমার project-এর হৃদয়।

> এই "সিঁড়ির গল্প" viva-তে বললে teacher বুঝবে তুমি বিচ্ছিন্ন তথ্য মুখস্থ করোনি, পুরো ছবিটা ধরেছো। এটাই ১০০% মার্কের চাবি।

---

## ২. সাপ্তাহিক সময়সূচি (প্রতি chapter-এর schedule)

প্রতিটা chapter আমরা একই ৪-ধাপ পদ্ধতিতে শেষ করবো (নিচে দেখো)। আজ June 28। presentation week July 20। তাই হাতে ~৩ সপ্তাহ। নিরাপদ একটা ছন্দ:

| সপ্তাহ | কী করবে |
|---|---|
| **সপ্তাহ ১** (Jun 28 – Jul 5) | Chapter 1, 2, 3, 4 (তোমার নিজের অংশ) — গভীরভাবে |
| **সপ্তাহ ২** (Jul 6 – Jul 12) | Chapter 5, 6, 7, 8 |
| **সপ্তাহ ৩** (Jul 13 – Jul 19) | Chapter 9, 10, 11, 12 + পুরো revision + viva mock |
| presentation-এর আগে | প্রতিটা chapter-এর "Viva Q&A" ফাইল আরেকবার পড়া |

প্রতিদিন ১–২টা chapter ধরলে সহজেই কুলিয়ে যাবে। তাড়াহুড়া নাই, কিন্তু নিয়মিত।

### প্রতি chapter-এ আমার ৪-ধাপ teaching পদ্ধতি
প্রতিটা chapter folder-এ তুমি ৩টা/৪টা ফাইল পাবে:
1. **`01_Lesson_Bangla.md`** — পুরো chapter বাংলায়, গল্প/উদাহরণ/analogy সহ, একটা টপিকও বাদ না দিয়ে।
2. **`02_Viva_Questions_Bangla.md`** — teacher যেসব প্রশ্ন করতে পারে + প্রতিটার পরিষ্কার বাংলা উত্তর (corner case সহ)।
3. **`03_Exercise_and_Solution_Bangla.md`** — ওই chapter-এর exercise + কেন উত্তর এমন হলো তার ব্যাখ্যা।
4. মূল exercise ও solution ফাইল (PDF/notebook) — copy করে একই folder-এ রাখা, যাতে এক জায়গায় সব পাও।

---

## ৩. কীভাবে পড়লে সত্যিই মাথায় গাঁথবে (পদ্ধতি)

1. **আগে `01_Lesson` পড়ো** — তাড়াহুড়া না করে, প্রতিটা formula-র "কেন" বোঝো।
2. **নিজে কাগজে formula derive করার চেষ্টা করো** (যেমন Jacobian, LCG period)। হাতে লিখলে গাঁথে।
3. **`03_Exercise` চালাও** — code টা নিজে run করো, parameter পাল্টে দেখো কী হয়।
4. **`02_Viva_Questions` জোরে জোরে উত্তর দাও** — বইয়ের দিকে না তাকিয়ে। আটকে গেলে আবার lesson-এ ফেরো।
5. **প্রতি সপ্তাহ শেষে** আগের chapter-গুলোর viva প্রশ্ন mix করে নিজেকে test করো।
6. **আমার সাথে feedback loop:** কোনো জায়গা কঠিন লাগলে আমাকে বলো — আমি ওই অংশটা আরও সহজ করে, আরও উদাহরণ দিয়ে আবার লিখে দেবো।

---

## ৪. Presentation ও প্রশ্নোত্তরে ১০০% পাওয়ার টিপস (project document থেকে)

- Slide-এ প্রথম পাতায় title + সব member-এর নাম। **"THANK YOU" slide দেওয়া নিষেধ।**
- শেষ slide-এ একটা **TL;DR / take-home message** + contact info।
- প্রতিটা figure-এ caption ও label দিতে হবে। Slide ভর্তি করে ফেলো না — পরিষ্কার রাখো।
- Report-এ **explicit proper priors** লিখতে হবে, network architecture (কয় layer, activation), training budget, **diagnostics: convergence + Simulation-Based Calibration (SBC) + posterior contraction**, posterior predictive checks।
- প্রশ্নোত্তরের জন্য: **প্রতি chapter-এর "Viva Q&A" ফাইলই তোমার মূল অস্ত্র।** teacher general প্রশ্ন করলে এখান থেকেই আসবে।

---

## ৫. এখন কী আছে, পরে কী আসবে

- ✅ **এখন তৈরি:** এই প্ল্যান + **Chapter 1** ও **Chapter 2**-এর পূর্ণ folder (lesson + viva + exercise + solution)।
- ⏭️ **পরের ধাপ:** তুমি Chapter 1, 2 দেখে feedback দেবে (teaching style ঠিক আছে কিনা, আরও সহজ লাগবে কিনা)। তারপর আমি Chapter 3, 4 … করে ১২ পর্যন্ত একই কায়দায় বানাবো।
- 🎯 **project topic পেলে:** আলাদা একটা BayesFlow-focused প্ল্যান + code বানাবো।

> চলো শুরু করি। Chapter 1 আর 2 এর folder খুলে `01_Lesson_Bangla.md` দিয়ে শুরু করো। 💪
