# অধ্যায় ৪: Risk Minimization I
## — Opus-এর বিস্তারিত বাংলা ব্যাখ্যা (একদম শিশুর জন্য)

---

> **এই নোটটা কাদের জন্য?**
> তোমার জন্য, যে loss function, risk minimization, gradient descent আর maximum likelihood আগে গভীরভাবে বোঝোনি। ধরে নিচ্ছি তুমি Chapter 1–3 পড়ে এসেছ — মানে regression, hypothesis space, bias-variance জানো। এই অধ্যায়ের নতুন concept গুলো (loss, residual, pseudo-residual, empirical risk, ৬ ধরনের regression loss, gradient descent পরিবার, MLE↔loss) — সব ভেঙে ভেঙে বলব।

> **কীভাবে পড়বে?**
> ১) প্রতিটা section পড়ার পর থেমে নিজেকে জিজ্ঞেস করো: "এটা ক্লাস ৮-এর ছাত্রকেও বোঝাতে পারব?"
> ২) Quiz-trap গুলো লাল কালি দিয়ে highlight করো।
> ৩) **৬টা Loss-এর optimal constant model** আর **MLE↔Loss জোড়া** — এই দুই জায়গায় সবচেয়ে বেশি ভুল হয়। আলাদা সময় নাও।

---

# 📚 সূচিপত্র (Table of Contents)

1. [Learning-এর তিনটা অংশ](#১-learning-এর-তিনটা-অংশ)
2. [Loss — ভুল মাপার যন্ত্র](#২-loss--ভুল-মাপার-যন্ত্র)
3. [Residual ও Pseudo-Residual](#৩-residual-ও-pseudo-residual)
4. [Distance-based ও Translation-invariant](#৪-distance-based-ও-translation-invariant)
5. [Theoretical Risk Minimization](#৫-theoretical-risk-minimization)
6. [L2-loss → Conditional Expectation](#৬-l2-loss--conditional-expectation)
7. [Empirical Risk Minimization (ERM)](#৭-empirical-risk-minimization-erm)
8. [Loss বেছে নেওয়া কেন গুরুত্বপূর্ণ](#৮-loss-বেছে-নেওয়া-কেন-গুরুত্বপূর্ণ)
9. [L2-loss বিস্তারিত](#৯-l2-loss-বিস্তারিত)
10. [L1-loss বিস্তারিত](#১০-l1-loss-বিস্তারিত)
11. [Quantile (Pinball) loss](#১১-quantile-pinball-loss)
12. [Huber loss](#১২-huber-loss)
13. [ε-insensitive ও Log-barrier loss](#১৩-ε-insensitive-ও-log-barrier-loss)
14. [Gradient Descent](#১৪-gradient-descent)
15. [SGD ও Mini-batch](#১৫-sgd-ও-mini-batch)
16. [Maximum Likelihood vs ERM](#১৬-maximum-likelihood-vs-erm)
17. [সব মিলিয়ে একটা বড় গল্প](#১৭-সব-মিলিয়ে-একটা-বড়-গল্প)
18. [Master Quiz-Trap Table](#১৮-master-quiz-trap-table)
19. [Golden Memorization Rules](#১৯-golden-memorization-rules)
20. [পরীক্ষার আগে করণীয়](#২০-পরীক্ষার-আগে-করণীয়)

---

# ১. Learning-এর তিনটা অংশ

## 🧵 গল্প: দর্জির দোকান

ভালো model বানানো ঠিক জামা বানানোর মতো — তিনটা জিনিস লাগে:

$$\text{Learning} = \underbrace{\text{Representation}}_{\text{কোন নকশা?}} + \underbrace{\textbf{Cost}}_{\text{কোনটা ভালো?}} + \underbrace{\text{Optimization}}_{\text{কীভাবে বানাব?}}$$

**① Representation:** কী ধরনের model — linear? tree? neural net? এটাই hypothesis space $\mathcal{H}$ (Chapter 3-এ পড়েছ)।

**② Cost:** কোন model ভালো কোনটা খারাপ — এটা মাপার নিয়ম। একেই বলে **cost / loss function**। (এই chapter-এর মূল বিষয়)

**③ Optimization:** $\mathcal{H}$-এর মধ্যে সেরা model $\hat f$ কীভাবে খুঁজে বের করব। (এই chapter-এর দ্বিতীয় বিষয়)

## 🎯 মূল প্রশ্ন (Scenario)

> একই linear model space দিলে, learning algorithm কোন model ফেরত দেবে?
> **উত্তর: এটা পুরোপুরি cost function-এর উপর নির্ভর করে।**

একই data, একই linear space — কিন্তু L2 loss দিলে এক রেখা, L1 দিলে আরেক রেখা, quantile দিলে আরেকটা। **Cost-ই নির্ধারণ করে কোন রেখা।**

---

# ২. Loss — ভুল মাপার যন্ত্র

## 📖 আমরা কী চাই?

আমরা চাই prediction $f(\mathbf{x})$ আসল target $y$-এর কাছাকাছি হোক: $y \approx f(\mathbf{x})$।

কিন্তু "কাছাকাছি" মানে কী? কতটা ভুল হলো সেটা **point-wise** (প্রতিটা data point-এ আলাদা করে) মাপতে হবে। এই মাপার যন্ত্রই **loss function**:

$$L : \mathcal{Y} \times \mathbb{R}^g \to \mathbb{R}_{\ge 0}$$

ভেঙে বুঝি:
- Input: আসল target $y$ আর prediction $f(\mathbf{x})$
- Output: একটা **non-negative** সংখ্যা — কতটা ভুল হলো

## 📏 দুটো অপরিহার্য শর্ত

**শর্ত ১ — Non-negativity:**
$$L(y, \tilde y) \ge 0 \quad \forall y, \tilde y$$
ভুলের পরিমাণ কখনো ঋণাত্মক হতে পারে না (−৫ unit ভুল বলে কিছু নেই)।

**শর্ত ২ — Optimality:**
$$L(y, \tilde y) = 0 \iff y = \tilde y$$
ভুল ঠিক শূন্য হবে **তখনই, যখন** prediction একদম নিখুঁত।

## 📊 উদাহরণ: Squared Error

$$L(y, f(\mathbf{x})) = (y - f(\mathbf{x}))^2$$

prediction আর আসল মানের পার্থক্যকে square করা। দূরত্ব যত বেশি, square তত বড়।

---

# ৩. Residual ও Pseudo-Residual

## 📖 Residual কী?

$$r := y - f(\mathbf{x}) \qquad (\text{আসল মান} - \text{prediction})$$

প্রতিটা data point-এর জন্য: $r^{(i)} := y^{(i)} - f(\mathbf{x}^{(i)})$।

Regression loss সাধারণত **শুধু residual-এর উপর** নির্ভর করে — মানে কতটা দূরে আছি সেটাই মুখ্য।

## 🔑 Pseudo-Residual — একটু কঠিন কিন্তু গুরুত্বপূর্ণ

**Pseudo-residual** = loss-এর $f$-এর সাপেক্ষে **ঋণাত্মক (negative) first derivative**:

$$\tilde r := -\frac{\partial L(y, f(\mathbf{x}))}{\partial f}$$

**কেন "negative"?** কারণ আমরা loss **কমাতে** চাই — derivative-এর উল্টো দিকে গেলে loss কমে।

## ✨ সবচেয়ে সুন্দর জিনিস

L2-loss নাও: $L = (y - f)^2$ (অথবা $0.5(y-f)^2$)।

$$\tilde r = -\frac{\partial\, 0.5(y-f)^2}{\partial f} = (y - f) = r$$

> **L2-loss-এর জন্য pseudo-residual আর সাধারণ residual একদম এক!** — এজন্যই নাম "pseudo-residual" (residual-এর মতো, কিন্তু আসলে derivative থেকে আসা)।

## 📈 Loss Plot

$L(y, f(\mathbf{x}))$ (y-অক্ষ) বনাম residual $y - f(\mathbf{x})$ (x-অক্ষ) — এই গ্রাফকে বলে **loss plot**।

> Pseudo-residual = ওই বিন্দুতে **tangent (স্পর্শক)-এর slope**।

L2-এর loss plot একটা **parabola** (U-আকৃতি) — মাঝখানে (residual=0) সবচেয়ে নিচু।

---

# ৪. Distance-based ও Translation-invariant

দুটো গুরুত্বপূর্ণ ধর্ম, যা প্রায়ই quiz-এ আসে।

## 📖 Distance-based

একটা loss **distance-based** যদি:
- শুধু residual দিয়ে লেখা যায়: $L(y, f(\mathbf{x})) = \psi(r)$ কোনো $\psi : \mathbb{R} \to \mathbb{R}$-এর জন্য,
- এবং residual শূন্য হলে loss শূন্য: $\psi(0) = 0$।

## 📖 Translation-invariant

একটা loss **translation-invariant** যদি $y$ আর $f(\mathbf{x})$ উভয়ে একই $a$ যোগ করলে loss না বদলায়:

$$L(y + a, f(\mathbf{x}) + a) = L(y, f(\mathbf{x})), \quad a \in \mathbb{R}$$

## 🔑 মূল সত্য (Theorem)

> একটা loss **translation-invariant হয় তখনই, যখন সেটা distance-based** — দুটো equivalent!

**কেন? (intuition):**
- **distance-based → translation-invariant:** loss শুধু $r = y - f$ দিয়ে লেখা। $y$ আর $f$ দুটোতে $a$ যোগ করলে $r$ একই থাকে ($(y+a)-(f+a) = y-f$), তাই loss একই।
- **translation-invariant → distance-based:** residual নিজেই translation-invariant, তাই শুধু residual-নির্ভর loss-ও translation-invariant।

## ⚠️ Quiz Traps
- "Pseudo-residual = loss-এর positive derivative" → **FALSE** (negative)
- "L2-তে pseudo-residual আর residual আলাদা" → **FALSE** (এক)
- "Translation-invariant আর distance-based আলাদা ধর্ম" → **FALSE** (equivalent)

---

# ৫. Theoretical Risk Minimization

## 📖 Risk কী?

$\mathbb{P}_{xy}$ = data যে joint distribution থেকে আসে।

**Goal:** এমন $f$ খুঁজো যা **গড় ভুল (expected loss = risk)** সবচেয়ে কম করে:

$$\min_{f \in \mathcal{H}} \mathcal{R}(f) = \min_{f \in \mathcal{H}} \mathbb{E}[L(y, f(\mathbf{x}))] = \min_{f} \int_{\mathcal{X} \times \mathcal{Y}} L(y, f(\mathbf{x})) \, d\mathbb{P}_{xy}$$

**পার্থক্য বোঝো:**
- **Loss** = একটা point-এ ভুল।
- **Risk** = সব সম্ভাব্য $(\mathbf{x}, y)$-এর উপর গড় ভুল।

---

# ৬. L2-loss → Conditional Expectation

এটা একটা সুন্দর গাণিতিক ফলাফল। ধীরে পড়ো।

## 🧮 প্রশ্ন

$\mathcal{Y} = \mathbb{R}$, L2-loss, আর hypothesis space-এ **কোনো বাধা নেই** (সব function allowed)। তাহলে সেরা $\hat f$ কী?

## 🔍 সমাধান

Law of total expectation দিয়ে:
$$\mathcal{R}(f) = \mathbb{E}_x\!\left[\mathbb{E}_{y|x}\!\left[(y - f(\mathbf{x}))^2 \mid \mathbf{x}\right]\right]$$

যেহেতু কোনো বাধা নেই, প্রতিটা $\mathbf{x}$-এর জন্য আলাদা করে minimize করি:
$$\hat f(\mathbf{x}) = \arg\min_c \mathbb{E}_{y|x}\!\left[(y - c)^2 \mid \mathbf{x}\right]$$

এখন এই পরিচয় ব্যবহার করি:
$$\mathbb{E}[(y - c)^2] = \underbrace{\mathbb{E}[(y-c)^2] - (\mathbb{E}[y] - c)^2}_{\text{Var}[y]} + (\mathbb{E}[y] - c)^2 = \text{Var}[y] + (\mathbb{E}[y] - c)^2$$

$\text{Var}[y]$ তো $c$-এর উপর নির্ভর করে না। তাই minimize হয় যখন $(\mathbb{E}[y] - c)^2 = 0$, মানে $c = \mathbb{E}[y]$।

## 🔑 ফলাফল

> **Squared loss-এর জন্য সেরা prediction = conditional expectation $\mathbb{E}[y \mid \mathbf{x}]$।**

## ⚠️ সমস্যা (Limitation)

কিন্তু বাস্তবে এটা সরাসরি করা যায় না:
- **$\mathbb{P}_{xy}$ অজানা।** (জানা থাকলে তো সরাসরি optimal prediction বানাতাম!)
- Non-parametric estimate (kernel density) **high dimension-এ ভেঙে পড়ে** (curse of dimensionality)।
- কড়া distributional assumption দিলে estimate করা যায় (যেমন LDA: $x|y \sim \mathcal{N}_p(\mu_k, \Sigma)$, QDA: $\Sigma_k$), কিন্তু ML সাধারণত **আরো flexible** model নিয়ে কাজ করে।

---

# ৭. Empirical Risk Minimization (ERM)

## 💡 সমাধান: data দিয়ে approximate করো

$\mathbb{P}_{xy}$ অজানা, কিন্তু আমাদের কাছে data $\mathcal{D}$ আছে (i.i.d. drawn from $\mathbb{P}_{xy}$)। তাই risk-কে data দিয়ে approximate করি — **empirical risk**:

$$\mathcal{R}_{\text{emp}}(f) = \sum_{i=1}^n L\!\left(y^{(i)}, f(\mathbf{x}^{(i)})\right)$$

$$\bar{\mathcal{R}}_{\text{emp}}(f) = \frac{1}{n}\sum_{i=1}^n L\!\left(y^{(i)}, f(\mathbf{x}^{(i)})\right)$$

> $\tfrac1n$ গুণফল optimization-এ কোনো পার্থক্য করে না (minimizer একই থাকে), তাই বেশিরভাগ সময় $\mathcal{R}_{\text{emp}}$ ব্যবহার করি।

## ⚠️ গুরুত্বপূর্ণ Note

> $\mathcal{R}_{\text{emp}}$ ভালো approximation **শুধু তখনই** যখন $\mathcal{D}$ একটা **unbiased, independent, এবং যথেষ্ট বড়** sample। নাহলে empirical risk আর theoretical risk মিলবে না।

## 📖 Learning = Optimization

Learning দাঁড়ায়:
$$\hat f = \arg\min_{f \in \mathcal{H}} \mathcal{R}_{\text{emp}}(f)$$

$f$ যদি $\boldsymbol\theta$ দিয়ে parametrized হয় (যেমন linear model $f(\boldsymbol\theta) = \boldsymbol\theta^\top \mathbf{x}$):
$$\hat{\boldsymbol\theta} = \arg\min_{\boldsymbol\theta \in \Theta} \mathcal{R}_{\text{emp}}(\boldsymbol\theta)$$

> **🔑 Learning (প্রায়ই) মানে এই optimization problem solve করা — ML আর optimization-এর গভীর সংযোগ।**

---

# ৮. Loss বেছে নেওয়া কেন গুরুত্বপূর্ণ

দুই দিক থেকে loss-এর পছন্দ গুরুত্বপূর্ণ:

## ① Statistical Properties
Loss-এর পছন্দ $f$-এর statistical ধর্ম নির্ধারণ করে — যেমন:
- **Robustness** (outlier কতটা সমস্যা করবে)
- **Implicit error distribution** (পরে Section ১৬-এ দেখব)

## ② Computational / Optimization Complexity

Optimization কতটা কঠিন, সেটা মূলত loss-এর উপর নির্ভর করে:

**Smoothness:** কিছু optimization method (gradient method) smoothness দরকার। L1 non-smooth, তাই কঠিন।

**Uni- vs Multi-modality:**
> যদি $L(y, f(\mathbf{x}))$ তার ২য় argument-এ **convex** হয়, **আর** $f(\mathbf{x} \mid \boldsymbol\theta)$ $\boldsymbol\theta$-তে **linear** হয়, তাহলে $\mathcal{R}_{\text{emp}}(\boldsymbol\theta)$ convex — **প্রতিটা local minimum-ই global minimum!**
>
> কিন্তু $L$ **non-convex** হলে → $\mathcal{R}_{\text{emp}}$-এর **একাধিক local minima** থাকতে পারে (খারাপ! optimizer আটকে যেতে পারে)।

## ⚠️ Quiz Traps
- "Theoretical risk সহজে minimize করা যায় কারণ P_xy জানা" → **FALSE** (অজানা)
- "1/n empirical risk-এর minimizer বদলে দেয়" → **FALSE** (বদলায় না)
- "Non-convex loss-এও একটাই global minimum থাকে" → **FALSE** (একাধিক হতে পারে)

---

# ৯. L2-loss বিস্তারিত

$$L(y, f(\mathbf{x})) = (y - f(\mathbf{x}))^2 \quad \text{বা} \quad 0.5(y - f(\mathbf{x}))^2$$

## 📖 ধর্ম

- **Outlier-sensitive:** residual ২ গুণ হলে loss **৪ গুণ** হয় (square-এর কারণে)। তাই $y$-তে outlier থাকলে model টেনে নিয়ে যায় — সমস্যা।
- **Convex ও differentiable:** gradient minimization-এ কোনো সমস্যা নেই (মসৃণ parabola)।
- **Residual = Pseudo-residual:** $-\partial(0.5(y-f)^2)/\partial f = y - f = r$।

## 🧮 Optimal Constant Model

প্রশ্ন: $\mathcal{H} = \{f(\mathbf{x}) = \theta \mid \theta \in \mathbb{R}\}$ (constant model) হলে, L2-এর সেরা $\theta$ কী?

$$\frac{\partial \mathcal{R}_{\text{emp}}(\theta)}{\partial \theta} = 2\sum_{i=1}^n (y^{(i)} - \theta) \stackrel{!}{=} 0$$
$$\sum_{i=1}^n y^{(i)} - n\theta = 0 \implies \boxed{\hat\theta = \frac{1}{n}\sum_{i=1}^n y^{(i)} = \bar y}$$

> **🔑 L2-এর optimal constant = গড় (mean) $\bar y$।**

---

# ১০. L1-loss বিস্তারিত

$$L(y, f(\mathbf{x})) = \lvert y - f(\mathbf{x}) \rvert$$

## 📖 ধর্ম

- **Robust:** L2-এর চেয়ে outlier-এর প্রতি অনেক কম সংবেদনশীল (square করে না, তাই outlier কম টানে)।
- **Convex কিন্তু $y = f(\mathbf{x})$-এ differentiable না** — ওখানে একটা ধারালো কোণ (V-আকৃতি)। তাই optimization **কঠিন** হয়ে যায়।

## 🧮 Optimal Constant Model

$$\hat\theta = \arg\min_{\theta \in \mathbb{R}} \sum_{i=1}^n \lvert y^{(i)} - \theta \rvert = \text{median}(y^{(i)})$$

> **🔑 L1-এর optimal constant = median।** (Proof: exercise)

## 💡 কেন L1 robust?

ধরো data: {1, 2, 3, 4, **100**} (100 একটা outlier)।
- **Mean (L2):** (1+2+3+4+100)/5 = **22** — outlier টেনে নিয়ে গেছে!
- **Median (L1):** **3** — outlier-এর কোনো প্রভাব নেই!

এজন্যই L1 robust।

## ⚠️ Quiz Traps
- "L2-এর optimal constant = median" → **FALSE** (mean)
- "L1-এর optimal constant = mean" → **FALSE** (median)
- "L1 সব জায়গায় differentiable" → **FALSE** (y=f-এ না)

---

# ১১. Quantile (Pinball) loss

$$L(y, f(\mathbf{x})) = \begin{cases}(1-\alpha)(f(\mathbf{x}) - y), & y < f(\mathbf{x}) \\ \alpha(y - f(\mathbf{x})), & y \ge f(\mathbf{x})\end{cases}, \quad \alpha \in (0, 1)$$

## 📖 ধর্ম

- **L1-এর extension:** $\alpha = 0.5$ দিলে L1-loss পাওয়া যায় (scaling সহ)।
- **Asymmetric (অসম):** positive আর negative residual-কে আলাদা ওজন দেয়।
  - $\alpha < 0.5$ → **over-estimation**-কে বেশি শাস্তি।
  - $\alpha > 0.5$ → **under-estimation** (positive residual, $y > f$)-কে বেশি শাস্তি।
- **Pinball loss** নামেও পরিচিত।

## 🧮 Optimal Constant Model

> **🔑 Quantile loss-এর optimal constant = empirical $\alpha$-quantile $Q_\alpha(\{y^{(i)}\})$।**

মানে $\alpha = 0.5$ → median, $\alpha = 0.25$ → প্রথম quartile, ইত্যাদি।

## 💡 উদাহরণ

$\alpha = 0.75$ হলে under-estimation (যখন আমরা কম predict করি)-কে বেশি শাস্তি দেয়। তাই model উপরের দিকে টানে — ৭৫তম percentile-এ বসে।

---

# ১২. Huber loss

$$L(y, f(\mathbf{x})) = \begin{cases}\tfrac12(y - f(\mathbf{x}))^2, & \lvert y - f(\mathbf{x}) \rvert \le \delta \\ \delta \lvert y - f(\mathbf{x}) \rvert - \tfrac12\delta^2, & \text{নাহলে}\end{cases}$$

## 📖 ধর্ম — দুই জগতের সেরা

- **Piecewise combination of L1 and L2:**
  - ছোট residual ($\le \delta$)-এ → **L2-এর মতো** (মসৃণ, differentiable)।
  - বড় residual ($> \delta$)-এ → **L1-এর মতো** (linear, robust)।
- **Convex, differentiable, robust** — L1 আর L2-এর সুবিধা একসাথে! (L2-এর smoothness + L1-এর robustness)

## 🧮 Optimal Constant Model

> **Closed-form solution নেই!** Numerical optimization দরকার।
> Huber-এর constant model L1 আর L2-এর solution-এর **মাঝখানে** পড়ে।

$\delta$ ছোট → L1-এর কাছাকাছি; $\delta$ বড় → L2-এর কাছাকাছি।

---

# ১৩. ε-insensitive ও Log-barrier loss

## ① ε-insensitive Loss

$$L(y, f(\mathbf{x})) = \begin{cases}0, & \lvert y - f(\mathbf{x}) \rvert \le \epsilon \\ \lvert y - f(\mathbf{x}) \rvert - \epsilon, & \text{নাহলে}\end{cases}, \quad \epsilon \in \mathbb{R}_+$$

- **L1-এর modification:** $\epsilon$-এর কম error-এ **কোনো শাস্তি নেই** (একটা "সহনশীল অঞ্চল")।
- Convex, কিন্তু $y - f \in \{-\epsilon, \epsilon\}$-এ differentiable না। Closed-form নেই।
- **Support Vector Regression (SVR)**-এ ব্যবহৃত হয়।

## ② Log-barrier Loss

$$L(y, f(\mathbf{x})) = \begin{cases}-a^2 \log\!\left(1 - \left(\tfrac{\lvert y-f\rvert}{a}\right)^2\right), & \lvert y - f \rvert \le a \\ \infty, & \lvert y - f \rvert > a\end{cases}$$

- ছোট residual-এ **L2-এর মতো** আচরণ করে।
- ব্যবহার করি যখন চাই **residual যেন কখনোই $a$-এর বেশি না হয়** (barrier = দেয়াল, $\infty$ শাস্তি)।
- **Risk minimization-এর solution থাকার কোনো নিশ্চয়তা নেই।**

## 📊 সব Loss-এর তুলনা

```
y-f → 0 থেকে দূরে গেলে:
L2:          সবচেয়ে দ্রুত বাড়ে (parabola)
L1:          সরলরেখায় বাড়ে
Quantile:    অসম সরলরেখা
Huber:       কাছে L2, দূরে L1
ε-insensitive: ε পর্যন্ত flat (0), তারপর L1
Log-barrier: a-তে দেয়াল (∞)
```

---

# ১৪. Gradient Descent

## 🏔️ গল্প: পাহাড় থেকে নামা

ধরো তুমি কুয়াশায় ঢাকা পাহাড়ের চূড়ায়, নিচে নামতে চাও। কিছু দেখছ না। কী করবে? — পায়ের নিচে যেদিকে সবচেয়ে খাড়া ঢাল, সেদিকে এক পা এগোও। বারবার এটা করলে শেষমেশ উপত্যকায় (minimum) পৌঁছবে।

এটাই **gradient descent** — loss-এর "পাহাড়" থেকে নামা।

## 📖 কেন দরকার?

Closed-form (analytical) solution না থাকলে (যেমন Huber, ε-insensitive) → **numerical optimization** দরকার।

## 🧮 Update Rule

**Negative gradient**-এর দিকে ধাপ নাও (steepest descent = সবচেয়ে খাড়া নিচের দিক):

$$\boldsymbol\theta^{[j+1]} = \boldsymbol\theta^{[j]} - \alpha^{[j]} \cdot \nabla_{\boldsymbol\theta}\, \mathcal{R}_{\text{emp}}(\boldsymbol\theta)\big|_{\boldsymbol\theta = \boldsymbol\theta^{[j]}}$$

- $\alpha^{[j]}$ = **step-size**, এই context-এ একে **learning rate** বলে (যেমন $\alpha = 0.1$)।
- **First-order** iterative algorithm (শুধু first derivative লাগে)।

**Stopping rule (কখন থামবে):**
$$\frac{\lVert \boldsymbol\theta^{[j+1]} - \boldsymbol\theta^{[j]} \rVert}{\lVert \boldsymbol\theta^{[j]} \rVert} < \varepsilon \quad (\text{যেমন } \varepsilon = 0.0001)$$
মানে parameter আর তেমন বদলাচ্ছে না → পৌঁছে গেছি।

## 🔗 Pseudo-residual দিয়ে Update

Chain rule দিয়ে update rule-টা pseudo-residual দিয়ে লেখা যায়:

$$\boldsymbol\theta^{[j+1]} \leftarrow \boldsymbol\theta^{[j]} + \alpha^{[j]} \frac{1}{n}\sum_{i=1}^n \tilde r^{(i)} \cdot \nabla_{\boldsymbol\theta} f(\mathbf{x}^{(i)} \mid \boldsymbol\theta)$$

(এখানে pseudo-residual $\tilde r^{(i)}$ চলে এসেছে — এজন্যই pseudo-residual গুরুত্বপূর্ণ!)

## ⚠️ Quiz Traps
- "Gradient descent positive gradient-এর দিকে যায়" → **FALSE** (negative)
- "Learning rate আর step-size আলাদা জিনিস" → **FALSE** (একই)

---

# ১৫. SGD ও Mini-batch

## 🐢 সমস্যা: Full Gradient Descent ধীর

Full gradient হিসাব করতে **প্রতিটা** data point evaluate করতে হয় ($\sum_{i=1}^n$)। $n$ বড় হলে (১০ লক্ষ point) প্রতি ধাপ ব্যয়বহুল।

## ⚡ Stochastic Gradient Descent (SGD)

- Gradient-এর একটা stochastic approximation।
- $\sum_i \nabla_{\boldsymbol\theta} L$ **ব্যয়বহুল** হলে ব্যবহার করি।
- **মাত্র একটা random observation $i$** দিয়ে gradient approximate করে:

$$\boldsymbol\theta^{[j+1]} \leftarrow \boldsymbol\theta^{[j]} - \alpha^{[j]} \nabla_{\boldsymbol\theta} L\!\left(y^{(i)}, f(\mathbf{x}^{(i)} \mid \boldsymbol\theta)\right)$$

- Parameter sequence $\{\boldsymbol\theta^{[1]}, \boldsymbol\theta^{[2]}, \ldots\}$ **stochastic** (এলোমেলো) — কারণ প্রতি ধাপে random point।

## ⚖️ Mini-batch Gradient Descent

- Full GD (সব point) আর SGD (১টা point)-এর মাঝামাঝি **trade-off**।
- একটা random subset $I \subset \{1, 2, \ldots, n\}$ ব্যবহার করে:

$$\boldsymbol\theta^{[j+1]} \leftarrow \boldsymbol\theta^{[j]} - \alpha^{[j]} \sum_{i \in I} \nabla_{\boldsymbol\theta} L\!\left(y^{(i)}, f(\mathbf{x}^{(i)} \mid \boldsymbol\theta)\right)$$

## 📊 তিনটার তুলনা

| পদ্ধতি | কত point | গতি | Noise |
|--------|---------|-----|-------|
| **Full GD** | সব $n$ | ধীর | কম (নিখুঁত) |
| **SGD** | ১টা | দ্রুত | **বেশি noisy** |
| **Mini-batch** | subset $I$ | মাঝামাঝি | মাঝামাঝি |

> SGD computationally সস্তা কিন্তু **noisy**; mini-batch ভারসাম্য রাখে।

## ⚠️ Quiz Trap
- "Mini-batch সব n observation ব্যবহার করে" → **FALSE** (subset)

---

# ১৬. Maximum Likelihood vs ERM

এই chapter-এর সবচেয়ে সুন্দর সংযোগ। মন দিয়ে পড়ো।

## 📖 Maximum Likelihood দৃষ্টিকোণ

Regression-কে আরেকভাবে দেখি। ধরে নিই:
$$y = f_{\text{true}}(\mathbf{x}) + \epsilon, \quad \epsilon \sim p, \;\; \mathbb{E}[\epsilon] = 0, \;\; \epsilon \perp \mathbf{x}$$

তাহলে $y$ একটা distribution follow করে যার mean $f_{\text{true}}(\mathbf{x})$, density $p(y \mid \mathbf{x}, \boldsymbol\theta)$।

উদাহরণ: $\epsilon \sim \mathcal{N}(0, \sigma^2)$ হলে $y \sim \mathcal{N}(f_{\text{true}}(\mathbf{x}), \sigma^2)$।

## 🧮 Likelihood ও Negative Log-likelihood

Data i.i.d. ধরে, **maximum-likelihood principle** — likelihood সর্বোচ্চ করা:

$$\mathcal{L}(\boldsymbol\theta) = \prod_{i=1}^n p\!\left(y^{(i)} \mid \mathbf{x}^{(i)}, \boldsymbol\theta\right)$$

Product-এর বদলে log নিয়ে sum-এ পরিণত করি (সহজ), আর negative করি (maximize → minimize):

$$-\ell(\boldsymbol\theta) = -\sum_{i=1}^n \log p\!\left(y^{(i)} \mid \mathbf{x}^{(i)}, \boldsymbol\theta\right)$$

## ✨ নতুন Loss সংজ্ঞা

এখন একটা **নতুন loss** define করি:
$$L(y, f(\mathbf{x} \mid \boldsymbol\theta)) := -\log p(y \mid \mathbf{x}, \boldsymbol\theta)$$

তাহলে empirical risk:
$$\mathcal{R}_{\text{emp}}(\boldsymbol\theta) = \sum_{i=1}^n L\!\left(y^{(i)}, f(\mathbf{x}^{(i)} \mid \boldsymbol\theta)\right)$$

> **🔑 তাহলে MLE estimator $\hat{\boldsymbol\theta}$ (likelihood maximize করে) = loss-minimal estimator $\hat{\boldsymbol\theta}$ (risk minimize করে)। দুটো একই!**

## 🔄 মূল সম্পর্ক (খুব গুরুত্বপূর্ণ)

- **প্রতিটা error distribution → একটা equivalent loss function** (একই point estimator $\boldsymbol\theta$ দেয়)।
- Loss-এর **multiplicative/additive constant** বাদ দেওয়া যায় (minimizer বদলায় না)।
- **⚠️ উল্টোটা সবসময় হয় না!** প্রতিটা loss-এর জন্য error distribution থাকে **না**। **Hinge loss** তার বিখ্যাত উদাহরণ (এর কোনো corresponding error density নেই)।

## 🎯 বিখ্যাত জোড়া (অবশ্যই মুখস্থ করো)

### Gaussian errors → L2-loss
$\epsilon \sim \mathcal{N}(0, \sigma^2)$ হলে:
$$\mathcal{L}(\boldsymbol\theta) \propto \exp\!\left(-\frac{\sum_{i=1}^n (y^{(i)} - f(\mathbf{x}^{(i)} \mid \boldsymbol\theta))^2}{2\sigma^2}\right)$$
$$\implies -\ell(\boldsymbol\theta) \propto \sum_{i=1}^n (y^{(i)} - f(\mathbf{x}^{(i)} \mid \boldsymbol\theta))^2 \quad \text{(L2-loss!)}$$

### Laplace errors → L1-loss
$\epsilon$ Laplace distribution ($\frac{1}{2b}\exp(-\tfrac{\lvert x-\mu\rvert}{b})$) follow করলে:
$$-\ell(\boldsymbol\theta) \propto \sum_{i=1}^n \lvert y^{(i)} - f(\mathbf{x}^{(i)} \mid \boldsymbol\theta) \rvert \quad \text{(L1-loss!)}$$

> **🔑 মনে রাখো: Gaussian ⟺ L2, Laplace ⟺ L1।**

## 📊 Empirical Error Distribution

- Model fit করার পর residual-গুলোর **histogram** = "empirical" error distribution।
- L2 দিয়ে fit করলে residual-গুলো Gaussian-এর মতো দেখায়; L1 দিয়ে করলে Laplace-এর মতো (ধারালো চূড়া)।
- কিছু loss (Huber, ε-insensitive) আসল error density-র সাথে মেলে না, কিন্তু intuitively loss নির্ধারণ করে **residual কীভাবে বিতরণ হবে**।

## ⚠️ Quiz Traps
- "প্রতিটা loss-এর একটা error distribution আছে" → **FALSE** (hinge loss ব্যতিক্রম)
- "Gaussian error → L1-loss" → **FALSE** (L2!)
- "Laplace error → L2-loss" → **FALSE** (L1!)
- "Loss-এর constant minimizer বদলে দেয়" → **FALSE** (বদলায় না)

---

# ১৭. সব মিলিয়ে একটা বড় গল্প

ধরো তুমি বাড়ির দাম predict করার model বানাচ্ছ।

**Representation (Section ১):** তুমি ঠিক করলে linear model ব্যবহার করবে — $\hat f = \boldsymbol\theta^\top \mathbf{x}$।

**Cost — কোন loss? (Section ২, ৬):**
- যদি data পরিষ্কার হয় → **L2** (mean ভিত্তিক, efficient)।
- যদি কিছু "ভুতুড়ে দাম" (outlier) থাকে → **L1 বা Huber** (robust)।
- যদি তুমি চাও ৯০% ক্ষেত্রে দামের চেয়ে কম predict না করতে → **Quantile loss** ($\alpha = 0.9$)।

**Theoretical vs Empirical (Section ৫, ৭):** আসল $\mathbb{P}_{xy}$ অজানা, তাই তোমার data দিয়ে empirical risk minimize করবে।

**Optimization (Section ১৪-১৫):**
- L2 হলে closed-form (mean) আছে।
- Huber হলে closed-form নেই → gradient descent।
- Data বিশাল হলে → SGD বা mini-batch।

**MLE দৃষ্টিকোণ (Section ১৬):** তুমি আসলে অজান্তেই একটা assumption করছ — L2 মানে তুমি ধরে নিচ্ছ error Gaussian, L1 মানে Laplace। Loss বেছে নেওয়া = error distribution বেছে নেওয়া!

---

# ১৮. Master Quiz-Trap Table

| Statement | T/F | কারণ |
|-----------|-----|------|
| Learning = Representation + Cost + Optimization | TRUE | — |
| একই space-এ returned model cost-এর উপর নির্ভর করে | TRUE | — |
| Loss সবসময় ≥ 0 | TRUE | non-negativity |
| L = 0 ⟺ y = ỹ | TRUE | optimality |
| Residual r = y − f(x) | TRUE | — |
| Pseudo-residual = positive derivative of loss | FALSE | **negative** |
| L2-তে pseudo-residual = residual | TRUE | — |
| Loss plot-এ pseudo-residual = tangent-এর slope | TRUE | — |
| Translation-invariant ⟺ distance-based | TRUE | — |
| Risk = expected loss E[L(y,f(x))] | TRUE | — |
| Squared loss-এর সেরা prediction = E[y\|x] | TRUE | conditional expectation |
| Theoretical risk সহজ কারণ P_xy জানা | FALSE | P_xy **অজানা** |
| Non-parametric P_xy estimate high-dim-এ scale করে | FALSE | curse of dimensionality |
| R_emp ভালো approximation: unbiased, independent, বড় sample | TRUE | — |
| 1/n empirical risk-এর minimizer বদলায় | FALSE | বদলায় না |
| ERM মানে learning = optimization | TRUE | — |
| Convex loss + linear f → local min = global | TRUE | — |
| Non-convex loss-এ একটাই global min | FALSE | একাধিক local minima |
| L2 outlier-sensitive (×2 → ×4) | TRUE | — |
| L2 convex ও differentiable | TRUE | — |
| L2-এর optimal constant = mean | TRUE | — |
| L2-এর optimal constant = median | FALSE | mean |
| L1 robust, y=f-এ differentiable না | TRUE | — |
| L1-এর optimal constant = median | TRUE | — |
| Quantile loss α=0.5 = L1 | TRUE | — |
| Quantile-এর optimal constant = α-quantile | TRUE | — |
| Quantile α>0.5 → under-estimation শাস্তি | TRUE | — |
| Huber = piecewise L1 + L2 | TRUE | — |
| Huber differentiable + robust | TRUE | — |
| Huber-এর closed-form solution আছে | FALSE | numerical |
| Huber solution L1 ও L2-এর মাঝে | TRUE | — |
| ε-insensitive: ছোট error free | TRUE | — |
| Log-barrier: solution থাকার নিশ্চয়তা | FALSE | নেই |
| Gradient descent negative gradient দিকে | TRUE | — |
| Negative gradient = steepest descent | TRUE | — |
| Learning rate = step-size α | TRUE | — |
| SGD একটা random point ব্যবহার করে | TRUE | — |
| SGD সস্তা কিন্তু noisy | TRUE | — |
| Mini-batch সব n point ব্যবহার করে | FALSE | subset |
| Mini-batch = full GD ও SGD-এর trade-off | TRUE | — |
| MLE = loss-minimal estimator (L = −log p) | TRUE | — |
| প্রতিটা error distribution → একটা loss | TRUE | — |
| প্রতিটা loss → একটা error distribution | FALSE | hinge loss ব্যতিক্রম |
| Gaussian error ⟺ L2-loss | TRUE | — |
| Laplace error ⟺ L1-loss | TRUE | — |
| Loss-এর constant minimizer বদলে দেয় | FALSE | বদলায় না |

---

# ১৯. Golden Memorization Rules

### Rule 1: তিনটা অংশ
Learning = **Representation + Cost + Optimization**। Cost-ই নির্ধারণ করে কোন model বের হবে।

### Rule 2: Loss-এর দুই শর্ত
**Non-negative** আর **L=0 ⟺ y=ỹ**।

### Rule 3: Pseudo-residual
**Negative** derivative of loss w.r.t. $f$। L2-তে = সাধারণ residual। "Positive" বললে FALSE।

### Rule 4: Optimal Constant Model — তিনটা "M" + একটা Q
```
L2 (squared)  → Mean (গড়)
L1 (absolute) → Median
Quantile      → α-Quantile (Q_α)
Huber         → numerical (closed-form নেই)
```
> উল্টে বসালে (L2→median, L1→mean) **FALSE**।

### Rule 5: Robustness
```
L2:    outlier-sensitive (square করে বড় শাস্তি)
L1:    robust
Huber: differentiable (L2) + robust (L1) = best of both
```

### Rule 6: L1 Differentiability
L1 **$y = f$-এ differentiable না** (ধারালো কোণ)। "L1 সব জায়গায় differentiable" → FALSE।

### Rule 7: Convexity → Global Min
Convex loss + linear-in-θ model → প্রতিটা **local min = global min**। Non-convex → multiple minima।

### Rule 8: Theoretical Risk-এর সমস্যা
$\mathbb{P}_{xy}$ **অজানা** → directly minimize করা যায় না → তাই **empirical risk**। "P_xy জানা" → FALSE।

### Rule 9: Gradient Descent
**Negative** gradient (steepest descent)। Step-size α = **learning rate**।

### Rule 10: GD পরিবার
```
Full GD:    সব n point    (ধীর, নিখুঁত)
SGD:        ১টা point      (দ্রুত, noisy)
Mini-batch: subset I       (trade-off)
```
"Mini-batch সব point ব্যবহার করে" → FALSE।

### Rule 11: MLE ↔ Loss
$L = -\log p(y\mid\mathbf{x},\boldsymbol\theta)$ → MLE = loss-minimal estimator।

### Rule 12: Error Distribution ↔ Loss জোড়া
```
Gaussian (bell)  ⟺ L2
Laplace (peak)   ⟺ L1
```
মনে রাখো: **Gauss-2, Laplace-1**।

### Rule 13: এক দিকেই কাজ করে
প্রতিটা error distribution → loss (হ্যাঁ)। প্রতিটা loss → error distribution (**না!**, hinge loss ব্যতিক্রম)।

---

# ২০. পরীক্ষার আগে করণীয়

1. **এই full document একবার পড়ো** (১.৫–২ ঘণ্টা)
2. **Section ১৮-এর Master Quiz-Trap Table** ৩ বার পড়ো
3. **Section ৬টা Loss-এর optimal constant model** মুখস্থ করো (mean/median/quantile/numerical)
4. **Section ১৬-এর MLE↔Loss জোড়া** (Gaussian↔L2, Laplace↔L1) গেঁথে নাও
5. **Section ১৯-এর ১৩টা Golden Rules** মনে রাখো
6. **`true_false_quiz.md`** solve করো — ৩০ মিনিটে ৫৪টা প্রশ্ন
7. **ভুল উত্তরগুলো** আবার সংশ্লিষ্ট section-এ পড়ো

## 🏆 শেষ কথা

Chapter 4 logical, মুখস্থের জিনিস কম। ৫টা মূল idea বুঝলে যেকোনো প্রশ্নের উত্তর বের করতে পারবে:
1. **Loss নির্ধারণ করে কোন model বের হবে** (Representation + Cost + Optimization)
2. **Optimal constant:** L2→mean, L1→median, Quantile→α-quantile
3. **Robustness:** L2 sensitive, L1/Huber robust; Huber = differentiable + robust
4. **Gradient descent:** negative gradient; SGD/mini-batch হলো data-এর subset
5. **MLE ↔ Loss:** Gaussian↔L2, Laplace↔L1; কিন্তু এক দিকেই (hinge-এর error distribution নেই)

**মনে রাখো:** এই chapter বোঝার জিনিস। কেন L2 mean দেয়, কেন L1 robust, কেন Gaussian↔L2 — concept দিয়ে বুঝলে কখনো ভুলবে না।

পরীক্ষায় ভালো করার জন্য শুভ কামনা! 🎓

---

*— Opus-এর তরফ থেকে*
