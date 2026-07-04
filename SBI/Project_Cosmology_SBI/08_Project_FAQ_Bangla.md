# Project FAQ — তোমার সব প্রশ্নের উত্তর (বাংলায়)

> এই ফাইলে তুমি project নিয়ে যত প্রশ্ন করেছ, সব একসাথে গুছিয়ে বাংলায় উত্তর দেওয়া হলো। viva-র আগে এটা একবার পড়ে নিলেই পুরো ছবি মাথায় থাকবে। English term গুলো রাখা হয়েছে (report/presentation-এ ওগুলোই লাগবে)।

---

## পর্ব ১ — প্রজেক্ট কী, কেন, কোথা থেকে

### Q1. প্রজেক্টের motivation কী? কী সমস্যা সমাধান করছি, আর সেটা কেন গুরুত্বপূর্ণ?
তিন স্তরে বুঝি:
1. **বৈজ্ঞানিক সমস্যা (inverse problem):** আমরা মহাবিশ্বের মৌলিক সংখ্যা (H₀, Ωₘ, nₛ, Aₛ) জানতে চাই, কিন্তু মহাবিশ্বে experiment করা যায় না — শুধু observe করা যায় (যেমন P(k))। তাই উল্টো দিকে যেতে হয়: observation → parameters।
2. **পরিসংখ্যানগত সমস্যা:** forward (parameter→P(k)) সহজ, কিন্তু inverse (P(k)→parameter, uncertainty সহ) কঠিন। সাধারণ পদ্ধতি (MCMC + likelihood) দুই জায়গায় ভেঙে পড়ে — বাস্তব simulator-এর likelihood লেখা যায় না, আর প্রতিটা data-র জন্য ঘণ্টার পর ঘণ্টা লাগে।
3. **সমাধান:** **amortized, likelihood-free inference** (SBI/BayesFlow) — শুধু simulate করতে পারলেই হলো, আর একবার train করে যেকোনো নতুন data-র জন্য তাৎক্ষণিক উত্তর।

**কেন গুরুত্বপূর্ণ:** নতুন survey (Euclid, DESI) বিশাল data দিচ্ছে — প্রতিবার MCMC চালানো অসম্ভব; calibrated error bar ছাড়া "Hubble tension সত্যি কিনা" বলা যায় না। এটা আজকের cosmology-র frontier পদ্ধতি।

### Q2. এটা কি copy project? নতুন কী আছে? source কী?
**Copy/plagiarism না।** method (cosmology-তে SBI) আগে থেকেই আছে — কিন্তু course **ইচ্ছে করেই** সেটা চায় ("state-of-the-art বানানো লাগবে না")। **নতুন/তোমার অবদান:** নিজের implementation, priors, noise model, network design, **calibration analysis (SBC)**, আর H₀ degeneracy-র ব্যাখ্যা।
**Source:** (a) assignment = তোমার course-এর `sbi_project_info_2026.pdf`, Topic 5, lecturer Aayush; (b) method = SBI/NPE literature; (c) tool = CAMB (simulator) + BayesFlow (instructor Bürkner নিজেই co-author!) + Planck-2018 (fiducial মান)।

### Q3. আমাদের project ভালো কেন?
নতুনত্বের জন্য না (course novelty চায় না), বরং: **সঠিক method** (amortized NPE), **সঠিক tool** (BayesFlow — instructor-এর নিজের), **calibration-first** (SBC), **ground-truth দিয়ে যাচাই**, **physics বোঝা** (H₀ degeneracy), **reproducible + engineered** (Go demo), আর **পুরো syllabus ছোঁয়**। এই অক্ষগুলোতেই এই course মার্ক দেয়।

---

## পর্ব ২ — পদ্ধতি, data ও model

### Q4. কেন neural network? এটা কি image data?
**না, image না।** input হলো **১০০টা সংখ্যার একটা vector** (P(k) curve-এর ১০০ বিন্দু) — Excel-এর এক row-র মতো, ছবি না। neural network মানেই image না — এটা যেকোনো input→output সম্পর্ক শেখার **general যন্ত্র**। আমরা NN নিচ্ছি কারণ (১) সম্পর্কটা nonlinear, (২) আমাদের শুধু একটা guess না, পুরো **posterior** (uncertainty + correlation) দরকার।

### Q5. dataset কোথায়? কীভাবে বানানো?
**Download করিনি — নিজেরা বানিয়েছি** (এটাই simulation-based)। আছে `code/data.npz`-এ (১২,০০০ train + ২,০০০ val + ৩,০০০ test জোড়া)। বানানোর ৩ ধাপ: (১) prior থেকে random parameter নাও (Ch3), (২) simulator চালিয়ে P(k) বানাও, (৩) noise যোগ করো → (parameter, spectrum) জোড়া জমাও। প্রতিটা জোড়া = একটা flashcard (এক পিঠে ১০০ সংখ্যা প্রশ্ন, আরেক পিঠে ৪ সংখ্যা উত্তর)।

### Q6. data বানানোর পর কি data cleaning / feature selection লাগে?
**বেশিরভাগ লাগে না** — কারণ data simulator নিজে বানিয়েছে, তাই missing value/ভুল label নেই, প্রতিটা feature physical signal। যা করেছি: simulator-side validity (unphysical বাদ, ωc>0), scaling (log+standardize)। feature selection লাগে না — network নিজেই (summary network) feature বেছে নেয়। **তবে** আসল survey data দিলে আলাদা "cleaning" ফিরে আসে (galaxy bias, survey window — "simulation gap")।

### Q7. কোন model train করেছি, আর কেন এই model?
**Gaussian Neural Posterior Estimator** — একটা MLP যা posterior-এর **mean + full covariance** দেয়। কেন: posterior লাগে (point না), correlation/degeneracy ধরতে full covariance লাগে, আর এটা সবচেয়ে **সহজ-সঠিক** (দ্রুত, CPU-তে চলে, যাচাইযোগ্য — prototype-এর জন্য আদর্শ)। brief বলেছে "minimal working model"।

### Q8. অন্য model কেন না? আর কী কী model ব্যবহার করা যেত?
তিন স্তরে বিকল্প:
- **আরও ভালো density estimator:** MDN (multimodal), **Normalizing Flow** (যেকোনো আকৃতি — BayesFlow এটাই দেয়, আসল upgrade), Flow Matching।
- **অন্য SBI পদ্ধতি:** NLE (likelihood শেখা + MCMC), NRE (classifier), ABC (Ch10, neural ছাড়া)।
- **Classical:** MCMC (এখানে likelihood tractable বলে চলে, কিন্তু ধীর, amortized না); random forest/regression (শুধু point, posterior না — তাই বাদ)।
**সেরা upgrade = normalizing flow**, কারণ আমাদের দুই দুর্বলতা (বাঁকা banana + prior bound লিক) ঠিক সেটাই সারায়।

---

## পর্ব ৩ — course-এর সাথে যোগ

### Q9. শুধু প্রথম ৩ chapter ব্যবহার করেছি? বাকি chapter?
না — প্রথম ৩টা শুধু "যন্ত্রপাতি" (sampling, RNG, Monte Carlo)। আসল হৃদয় ৭–১২:
- Ch1 Monte Carlo (posterior summary, change of variables = lnAs), Ch2 RNG/seed, Ch3 prior sampling।
- Ch4 generative model + **PPC**, Ch6 **simulation study/coverage**, **Ch9 SBC** (মূল diagnostic)।
- Ch7 Bayes (prior×likelihood→posterior), Ch8 likelihood/MCMC (contrast)।
- **Ch10 ABC** (problem class), **Ch11 neural net**, **Ch12 NPE/BayesFlow** (method)।
একমাত্র দুর্বল যোগ = Ch5 (bootstrap) — viva-তে সৎভাবে বলবে।

---

## পর্ব ৪ — ফলাফল ও সংখ্যা বোঝা

### Q10. accuracy / prediction rate (১০০-র মধ্যে কত%) কি লাগে?
**না** — ওটা classification-এর ধারণা (ঠিক/ভুল)। আমাদের কাজ continuous সংখ্যা + uncertainty। বদলে আমরা দিই: **recovery r** (Ωₘ 0.97, nₛ 0.91, Aₛ 0.87, H₀ 0.30), **coverage/calibration** (৯৫% interval সত্য ধরে কিনা ≈ 0.96–1.00), **contraction**। জোর করে accuracy% দিলে ভুল বোঝাবে (H₀-এর r কম, কিন্তু model সঠিক)।

### Q11. baseline / limit / range কী?
- **range/limit = priors:** H₀[60,80], Ωₘ[0.1,0.5], nₛ[0.9,1.05], lnAs[2.5,3.5]।
- **baseline:** (a) prior নিজে (contraction এর চেয়ে কত ভালো মাপে), (b) linear regression (sanity check — R² Ωₘ=0.94, nₛ=0.99)।
- **আসল limit:** noise floor + H₀ degeneracy — physics-এর, কোনো model ভাঙতে পারবে না।

### Q12. কোন সংখ্যা certainty বোঝায়?
**± সংখ্যাটা (posterior standard deviation, σ)** — যত ছোট, তত নিশ্চিত। যেমন Ωₘ ±0.032 = খুব নিশ্চিত; H₀ ±6.2 = অনিশ্চিত। ০–১ স্কেলে চাইলে = **contraction** (১=অনেক শিখেছি, ০=কিছু না)। আর **coverage** আলাদা — ওই ± টা সৎ কিনা যাচাই করে।

### Q13. corr = −0.96 — এটা কি correlation? negative কেন?
হ্যাঁ, এটা H₀ ও Ωₘ-এর **correlation coefficient** (−১ থেকে +১; চিহ্ন=দিক, মান=কতটা শক্ত)। **Negative কেন:** spectrum-এর আকৃতি নির্ভর করে **Γ ≈ Ωₘ×h**-এর ওপর; Γ একই রাখতে H₀ বাড়লে Ωₘ **কমতেই হবে** → উল্টো দিক → negative। (যেমন নির্দিষ্ট bill দুজনে ভাগ করলে একজন বেশি দিলে আরেকজন কম।)

### Q14. recovery panel-এ কখনো সত্য মান ব্যান্ডের বাইরে — model কি খারাপ?
না, **স্বাভাবিক**। ±1σ মানে ৬৮% — সত্য মান ~৩ বারে ১ বার বাইরে থাকবেই। ৪টা parameter থাকলে গড়ে ~১.৩টা প্রতিবারই বাইরে। তাছাড়া প্রতি click-এ নতুন noise, আর nₛ/Aₛ correlated বলে একসাথে নড়ে। এক click দেখে বিচার না — **SBC/coverage** দিয়ে অনেকবার চালিয়ে বিচার হয় (সেটা ভালো এসেছে)।

---

## পর্ব ৫ — figure ও demo বোঝা

### Q15. P(k) plot — সাদা line কী, dot কী, কখনো line-এ কখনো বাইরে কেন, line এভাবে বাঁকে কেন?
- **সাদা line = সত্য (clean) P(k)** — noise ছাড়া মহাবিশ্বের আসল spectrum।
- **কমলা dot = মাপা (noisy) মান** — telescope যা পেত; **এই dot গুলোই network দেখে**, line না।
- **কখনো line-এ, কখনো দূরে:** noise (cosmic variance)। বাঁ দিকে (বড় scale) নমুনা কম → noise বেশি → dots ছড়ানো; ডান দিকে নমুনা অনেক → noise কম → dots line-এ।
- **line এভাবে বাঁকে কেন:** ওঠা = nₛ (tilt); চূড়া কোথায় = Ωₘ ও H₀; পুরো উচ্চতা = Aₛ। প্রতিটা বাঁক একটা parameter-এর স্বাক্ষর।

### Q16. P(k) plot-এ একদম বাঁয়ে y-axis-এর কাছে একলা নিচু একটা dot — স্বাভাবিক?
হ্যাঁ। ওটা সবচেয়ে ছোট k (বড় scale), যেখানে noise প্রায় **১০০%** (cosmic variance, cap করা)। এই draw-টায় মানটা প্রায় শূন্যে নেমে গেছে → log scale-এ একদম নিচে। bug না — সবচেয়ে noisy bin-এর স্বাভাবিক আচরণ। আবার click করলে অন্য জায়গায় লাফাবে।

### Q17. H₀–Ωₘ scatter — dots prior box-এর বাইরে চলে যায়, স্বাভাবিক?
এটা ছিল একটা ছোট খুঁত: posterior **Gaussian**, Gaussian-এর লেজ অসীম, তাই কিছু sample [60,80]-এর বাইরে পড়ত। কিন্তু uniform prior বলে বাইরের মান **অসম্ভব**, তাই সঠিক posterior হলো box-এ **truncated** — demo এখন বাইরের sample বাদ দেয় (ঠিক করা হয়েছে)। আসল BayesFlow (normalizing flow) bound মেনে চলত।

### Q18. degeneracy figure (banana) কী বলছে?
network বলছে: "H₀ আর Ωₘ আলাদা করে নিশ্চিত না, কিন্তু তাদের **সম্পর্ক** নিশ্চিত — একটা বাড়লে আরেকটা কমবে।" মেঘ ছড়ানো (আলাদা মান অনিশ্চিত) কিন্তু সরু (সম্পর্ক/মিশ্রণ নিশ্চিত)। কারণ spectrum মূলত **Γ=Ωₘ×h** মাপে, আলাদা করে না। এটা ভুল না — **সঠিক physics** (এজন্যই H₀-এর bar চওড়া)।

### Q19. "banana" মানে কী?
এটা technical শব্দ না — H₀–Ωₘ scatter-এর **আকৃতির ডাকনাম**। dots লম্বা, সরু, কাত, একটু **বাঁকা** মেঘ → দেখতে কলা (banana ফল)-র মতো। বাঁকা কেন: Γ=Ωₘ×h স্থির রাখলে Ωₘ=Γ/h একটা বাঁকা রেখা। (round blob=independent; সোজা ellipse=linear correlation; বাঁকা=banana=nonlinear degeneracy।)

### Q20. demo-তে আমরা ৪টা মান বসাই কেন — আমরা না inverse করছি?
৪টা মান বসানো = একটা **জানা পরীক্ষা** বানানো (truth)। আমরা ওগুলো দিয়ে forward-এ spectrum বানাই; **network ওই মান দেখেই না**, শুধু spectrum দেখে উল্টো করে অনুমান করে; তারপর তার অনুমান (teal) আমাদের সত্যের (gold) সাথে মিলিয়ে দেখি "inverse কাজ করল কিনা।" আসল জীবনে সত্য জানি না — শুধু real spectrum দিয়ে network-এ ভরসা করি।

### Q21. একই input-এ বারবার click করলে ভিন্ন output কেন?
প্রতি click-এ নতুন random **noise** (cosmic variance) যোগ হয় → observation একটু আলাদা → উত্তরও একটু নড়ে। সাথে posterior sample-ও random। বাস্তবেও মহাবিশ্ব দুবার মাপলে noise-এর জন্য একটু আলাদা মান আসত। উত্তর একটু নড়ে কিন্তু সত্যের কাছেই থাকে — ভালো method-এর লক্ষণ।

### Q22. Planck-2018 কী? এটা কি standard?
**Planck** = ESA-র space telescope যা Big Bang-এর afterglow (CMB) মেপেছে; ২০১৮-র ফল cosmological parameter-এর সবচেয়ে নিখুঁত মান দিয়েছে (H₀≈67.4, Ωₘ≈0.315, nₛ≈0.965, Aₛ≈2.1e-9)। **হ্যাঁ, এটাই standard/gold reference।** তাই এটাকে "সত্যিকারের মহাবিশ্ব" ধরে test করি। (ছোট টীকা: H₀ নিয়ে "Hubble tension" বিতর্ক — Planck বলে ~৬৭, কাছের supernova বলে ~৭৩।)

### Q23. loss curve-এর (fig3) বাঁ প্যানেলে লাল ও নীল লাইন এত আলাদা / overlap করে না কেন?
মূল কথা: এরা তুলনার দুটো লাইন **না** — এরা একই training-এর **দুটো আলাদা পর্ব (phase)**, একটার পর একটা সময়ে। তাই এক হওয়ার/overlap করার কথাই না।
- **🔴 লাল লাইন (Phase 1 — MSE pretraining):** training-এর প্রথম অংশ (step ০–২৪০০)। এই সময় শুধু **mean** ঠিক করছি (MSE loss), কিন্তু graph-এ **NLL** দেখাচ্ছি — যা তখন কমানোরই চেষ্টা করছি না। তাই NLL এলোমেলো হয়ে **উপরে ওঠে** (variance তখনো শেখানো হয়নি)। *উপমা: marathon-এর জন্য দৌড় প্র্যাকটিস করছ (MSE), কিন্তু graph-এ ওজন (NLL) দেখাচ্ছ — ওজন তখন লক্ষ্য না, তাই এদিক-ওদিক করে।*
- **┊ dashed লাইন:** এখানে Phase 1 → Phase 2-তে switch (loss বদলাই: MSE → NLL)।
- **🔵 নীল লাইন (Phase 2 — Gaussian NLL):** দ্বিতীয় অংশ (২৪০০ থেকে শেষ)। এখন আসল **NLL** কমাচ্ছি, তাই দ্রুত **নিচে নামে** ও সমতল হয় (converged)।
- **overlap করে না কেন:** (১) ভিন্ন সময় — লাল বাঁয়ে, নীল ডানে (একটার পর একটা); (২) ভিন্ন লক্ষ্য — লাল MSE, নীল NLL। তাই pattern আলাদা।
- **ডান প্যানেলে overlap করে কেন (তুলনা):** ওখানে দুই লাইন = **একই NLL, একই সময়ে, দুই ভিন্ন data** (train ও validation) → overlap = **overfitting নেই** ✅। বাঁ প্যানেলে দুই লাইন = ভিন্ন phase, ভিন্ন সময় → overlap করার কথাই না।

> এক লাইনে: **বাঁ প্যানেল = একই training-এর দুই ধাপ (আগে MSE, পরে NLL) — তাই আলাদা; ডান প্যানেল = একই NLL দুই data-তে একসাথে — তাই overlap (no overfit)।**

---

## পর্ব ৬ — engineering (তোমার Go skill)

### Q24. live demo (Go backend + frontend) কী যোগ করল?
একটা **pure Go server** (Python ছাড়াই) যা trained network-এর weight লোড করে দুটোই করে — forward (θ→P(k)) আর inverse (P(k)→posterior)। frontend-এ slider দিয়ে "মহাবিশ্ব" সেট করো → recovery bar + banana live দেখো। এটা **amortized inference** চোখে দেখানোর সেরা উপায় (MCMC-তে যা মিনিট লাগত, এখানে milliseconds)। তোমার backend skill-এর সরাসরি প্রয়োগ — presentation-এ বড় প্লাস। (চালাতে: `cd webdemo && go run .` → `http://localhost:8081`)

---

## 🎯 viva-র আগে এক নজরে মনে রাখো
1. সমস্যা = ঝাপসা P(k) থেকে ৪টা সংখ্যা **uncertainty সহ** বের করা (inverse, likelihood-free, amortized)।
2. data নিজে simulate করা; model = Gaussian NPE (upgrade = normalizing flow)।
3. ভালো = calibrated (SBC), H₀ degeneracy বোঝা; accuracy% না, **calibration**।
4. certainty = ± (σ); corr −0.96 = H₀–Ωₘ trade-off; banana = সেই degeneracy-র আকৃতি।
5. প্রতি click ভিন্ন = noise; ৪ মান বসাই = জানা পরীক্ষা; Planck = standard সত্য।

> কোনো প্রশ্নের উত্তর আরও সহজ/বিস্তারিত চাইলে বলো — যোগ করে দেব। 💪
