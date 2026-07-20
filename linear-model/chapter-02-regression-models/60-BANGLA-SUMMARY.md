# অধ্যায় ২ — রিগ্রেশন মডেল (২.১–২.৩) — বাংলা সারাংশ

> পরীক্ষার প্রায় **১৩%**। কিন্তু এখানকার নম্বর সবচেয়ে **সস্তা** — interpretation আর logit-এর দুটো নির্দিষ্ট উত্তর জানা থাকলেই হয়ে যায়।
>
> টেকনিক্যাল শব্দ ইংরেজিতেই রেখেছি, কারণ **পরীক্ষার উত্তর ইংরেজিতে লিখতে হবে**।

---

## ২.১ — রিগ্রেশন আসলে কী মডেল করে?

$$\boxed{\;E(y\mid\boldsymbol{x})=f(\boldsymbol{x})\;}\qquad\Longleftrightarrow\qquad y=f(\boldsymbol{x})+\varepsilon,\ E(\varepsilon)=0$$

> 🔑 **রিগ্রেশন $y$-এর মান নয়, $y$-এর গড় (conditional mean) মডেল করে।**

একই covariate থাকা দুজন মানুষের $y$ আলাদা হয় — তাই মডেল বলে "এরকম মানুষদের **গড়ে** কত"।

**এই কারণেই প্রতিটি interpretation-এ "expected" বা "on average" শব্দটা থাকতে হয়।** বাদ দিলে নম্বর কাটা যায়।

### Linear predictor — সব মডেলের কঙ্কাল

$$\eta=\boldsymbol{x}'\boldsymbol\beta=\beta_0+\beta_1x_1+\dots+\beta_kx_k$$

- **Linear model:** $E(y)=\eta$ (সরাসরি)
- **Logit model:** $P(y=1)=h(\eta)$ ($\eta$-কে চেপে $[0,1]$-এ আনা হয়)

**একই কঙ্কাল, আলাদা মোড়ক।**

---

## ২.২.১ — Simple Linear Regression

$$y_i=\beta_0+\beta_1x_i+\varepsilon_i$$

$$\hat\beta_1=\frac{\widehat{\text{Cov}}(x,y)}{\widehat{\text{Var}}(x)}=r\frac{s_y}{s_x},\qquad \boxed{\hat\beta_0=\bar y-\hat\beta_1\bar x}$$

### 💰 রেখা সবসময় $(\bar x,\bar y)$ দিয়ে যায়

R output-এ intercept হারিয়ে গেলে ১৫ সেকেন্ডে বের করা যায়:
$$\hat\beta_0=46.61-0.90509\times48.61=\boxed{2.613}$$

### ⭐ Slope ব্যাখ্যার টেমপ্লেট (২০ বার লিখতে হবে)

> **"A one-[unit] increase in [x] is associated with an estimated [β̂₁] [unit] change in the expected [y], holding all other covariates fixed."**

**পরীক্ষক চারটে জিনিস খোঁজে:**
1. $x$-এর **একক** ("one **year**")
2. $y$-এর **একক** ("**\$**0.71 per hour")
3. **"expected"** বা "on average"
4. **"associated with"** — কখনো "causes" নয়

Multiple regression হলে পঞ্চম: **"holding all other covariates fixed"**।

### ⚠️ Intercept কখন ব্যাখ্যা করবে না

$\hat\beta_0$ = $x=0$ হলে প্রত্যাশিত $y$।

**Sheet 1(1c)-এর উত্তর:** *বয়স ০ ডেটার পরিসরের **অনেক বাইরে** এবং **অর্থহীন** (নবজাতকের ঘণ্টাপ্রতি মজুরি নেই)। এটা ব্যাখ্যা করা মানে **extrapolation**। Intercept রেখাটাকে ঠিক জায়গায় বসায়, কিন্তু কোনো অর্থ বহন করে না।*

**সমাধান — centring:** $x$-এর বদলে $(x-\bar x)$ বা $(\text{age}-48)$ ব্যবহার করলে $\beta_0$ মানে দাঁড়ায় "৪৮ বছর বয়সে প্রত্যাশিত মজুরি" — অর্থবহ ও ডেটার ভেতরে।

### Intercept থাকলে যা স্বয়ংক্রিয়

$$\sum\hat\varepsilon_i=0,\qquad \sum x_i\hat\varepsilon_i=0,\qquad \bar{\hat y}=\bar y$$

---

## ২.২.২ — Multiple Regression ও Dummy Variable ⭐⭐⭐

### Partial effect

$\hat\beta_j$ = **অন্য সব covariate স্থির রেখে** $x_j$ এক একক বাড়লে প্রত্যাশিত $y$-এর পরিবর্তন।

> ⚠️ **Partial effect আর marginal correlation-এর চিহ্ন উল্টোও হতে পারে।**
>
> উদাহরণ: দমকলকর্মীর সংখ্যা ও ক্ষয়ক্ষতির correlation **ধনাত্মক**। কিন্তু **আগুনের আকার** স্থির রেখে দেখলে effect **ঋণাত্মক**।
>
> *(WS 23/24 Block I(i): "coefficient ধনাত্মক হলে correlation-ও ধনাত্মক হতে হবে" → **FALSE**।)*

---

### 🔴 THE DUMMY RULE — প্রতি বছর আসে

> **$c$ টি level ⟹ $c-1$ টি dummy। বাদ পড়া level = reference category।**

### কেন $c-1$? — দেয়ালে দাগের গল্প

পাঁচ বন্ধুর উচ্চতা মাপবে। একজনের মাথার উচ্চতায় দেয়ালে **দাগ** দাও। এখন বাকি সবাইকে **দাগ থেকে কতটা উপরে** বলে প্রকাশ করো।

```
আমিনার দাগ ─────────────  ০   ← REFERENCE
বিলাল      ─────────────  +১১
চেন        ─────────────  +২৪
দারা       ─────────────  +৪০
এরিক       ─────────────  +৬৫
```

**পাঁচজন, চারটে সংখ্যা।** যে **নিজেই দাগ**, তার জন্য কোনো সংখ্যা লাগে না।

**সবগুলোকে সংখ্যা দিলে কী হয়?**

| উত্তর ১ | উত্তর ২ |
|---|---|
| intercept=১০, আমিনা=০, চেন=৫ | intercept=০, আমিনা=১০, চেন=১৫ |

দুটোই **একই মডেল** বর্ণনা করে! চেন আমিনার চেয়ে ৫ বেশি — দুটোতেই।

> 🔑 **যে প্রশ্নের লক্ষ লক্ষ উত্তর, তার কোনো উত্তর নেই।**
>
> এটাই **dummy variable trap**: কলামগুলো linearly dependent ⟹ $\boldsymbol{X}'\boldsymbol{X}$ **singular** ⟹ $(\boldsymbol{X}'\boldsymbol{X})^{-1}$ নেই ⟹ **অদ্বিতীয় OLS সমাধান নেই**।
>
> **গণিত তোমাকে তুলনার ভিত্তি না বলে তুলনামূলক প্রশ্ন করতে দেয় না।** এটা ত্রুটি নয়, বৈশিষ্ট্য।

### তিনটি ব্যবহারিক নিয়ম

1. প্রতিটি dummy coefficient **reference-এর সাথে** তুলনা — কখনো পাশের level-এর সাথে নয়
2. দুটো non-reference level তুলনা করতে **coefficient বিয়োগ করো**
3. Reference = R output-এ **যে level-টা নেই**

### 💡 "যা মিল, তা কাটে" — সময় বাঁচানোর কৌশল

দুজন মানুষ তুলনা করার সময় **যা তাদের মধ্যে মিল, তা বাতিল হয়ে যায়**। দুটো prediction হিসাব করো না — **পার্থক্য সরাসরি বের করো**।

> *Sheet 1(c): দুজনেই ৫০ বছর, level 3 বনাম level 5।*
> Intercept ও age — দুটোই মিলে যায়।
> $$\hat\beta_5-\hat\beta_3=64.99-24.17=\boxed{\$40.82}$$
> **একটা বিয়োগ। ছটা হিসাবের বদলে।**

> *Sheet 1(d): ৪০ বছর/level 4 বনাম ২০ বছর/level 1।*
> শুধু intercept মেলে।
> $$0.56869\times20+39.767=\boxed{\$51.14}$$

---

### Interaction — দুটো এসকেলেটর

$$y=\beta_0+\beta_1x+\beta_2D+\beta_3(xD)+\varepsilon$$

| গ্রুপ | Intercept | Slope |
|---|---|---|
| $D=0$ | $\beta_0$ | $\beta_1$ |
| $D=1$ | $\beta_0+\beta_2$ | $\beta_1+\beta_3$ |

**Interaction ছাড়া:** দুটো এসকেলেটর **একই গতিতে** চলে, একটা শুধু উঁচুতে শুরু হয়। ফাঁক **কখনো বদলায় না** ⟹ **সমান্তরাল রেখা**।

**Interaction সহ:** এসকেলেটর দুটো **আলাদা গতিতে** চলে। ফাঁক বদলায়, রেখা **ক্রস করতে পারে** ⟹ **অসমান্তরাল রেখা**।

> **Dummy রেখা সরায়। Interaction রেখা হেলায়।**

### ⚠️ Interaction থাকলে main effect আর "the effect" নয়

$$\frac{\partial E(y)}{\partial x}=\beta_1+\beta_3D,\qquad \frac{\partial E(y)}{\partial D}=\beta_2+\beta_3x$$

**Sheet 2-এ:** স্বাস্থ্যের coefficient $-1.81$। এটা "খারাপ স্বাস্থ্যের সুবিধা" **নয়** — এটা **বয়স ০-তে** স্বাস্থ্যের effect, যেখানে কেউ দাঁড়িয়ে নেই।

৪০ বছরে effect: $-1.81+0.43(40)=+\$15.39$। ৬০ বছরে: $+\$24.99$।

> 🔑 *"এসকেলেটর B কতটা এগিয়ে?"* — এর কোনো একক উত্তর নেই। **উত্তর নির্ভর করে তুমি কখন জিজ্ঞেস করছ তার উপর।**

### একই নিয়ম Polynomial-এও

$$y=\beta_0+\beta_1x+\beta_2x^2 \qquad\Longrightarrow\qquad \frac{\partial E(y)}{\partial x}=\beta_1+2\beta_2x$$

**ছোঁড়া বলের উপমা:** $\hat\beta_1$ = **ছোঁড়ার গতি**, $\hat\beta_2$ = **মাধ্যাকর্ষণ**।

*"বলটা কত জোরে যাচ্ছে?"* — **কখন?** ছোঁড়ার সময় দ্রুত ও ঊর্ধ্বমুখী; চূড়ায় শূন্য; নামার সময় ঋণাত্মক।

$$\text{সর্বোচ্চ বিন্দু: } x^*=-\frac{\hat\beta_1}{2\hat\beta_2}=\frac{5.29}{0.10}=52.9 \text{ বছর}$$

**$\hat\beta_2<0$ তুমি ডেটা না দেখেই বলতে পারো** — কর্মজীবন, বলের মতোই, নেমে আসে।

> **⚠️ সাধারণ নিয়ম: একটি ভেরিয়েবল একাধিক পদে থাকলে, তার coefficient একা ব্যাখ্যা করা যাবে না — differentiate করতে হবে।**

---

## ২.৩ — Logit Model ⭐ নিশ্চিতভাবে আসবে

### সমস্যাটা কোথায়?

Binary $y\in\{0,1\}$ হলে:
$$E(y\mid\boldsymbol{x})=P(y=1\mid\boldsymbol{x})=\pi \qquad\Longrightarrow\qquad \pi \text{ অবশ্যই } [0,1]\text{-এ থাকবে}$$

কিন্তু $\boldsymbol{x}'\boldsymbol\beta$ একটা **সরলরেখা** — রেখা $\pm\infty$ পর্যন্ত যায়।

> 🔑 **তুমি সীমাবদ্ধ জিনিস দেয়ালবিহীন পাত্রে রাখার চেষ্টা করছ।**
>
> বাথটাবের উপমা: কল ১০০ বছর খুলে রাখলেও টাব "পূর্ণ"-এর বেশি ভরবে না। কিন্তু **সরলরেখা জানেই না যে মেঝে আর কানা বলে কিছু আছে**।

### 🎯 উত্তর ১ — কেন linear model কাজ করে না [Exam 2025 Ex 4(a), ১ নম্বর]

**চারটি কারণ। প্রথমটা দিয়ে শুরু করবে।**

1. 🔴 **Prediction $[0,1]$-এর বাইরে যায়** — মডেল ১.৩৪ বা $-০.২০$ সম্ভাবনা দেয়, যা অর্থহীন। ← **এটাই মূল কারণ**
2. $\text{Var}(y)=\pi(1-\pi)$ — $\boldsymbol{x}$-এর উপর নির্ভরশীল ⟹ **গঠনগতভাবেই heteroscedastic**
3. $\varepsilon$ মাত্র **দুটি মান** নেয় ⟹ normal হতেই পারে না ⟹ exact t/F test-এর ভিত্তি নেই
4. সর্বত্র **একই marginal effect** অবাস্তব — ০.৫০ থেকে ০.৫৫ সহজ, ০.৯৯ থেকে ১.০৪ অসম্ভব

### সমাধান — খাঁচা থেকে মুক্তি, দুই ধাপে

$$\pi\in[0,1] \xrightarrow{\ \text{odds}\ } (0,\infty) \xrightarrow{\ \log\ } (-\infty,\infty)$$

**Odds** সম্ভাবনাকে $[0,\infty)$-তে নিয়ে যায়। **Log** সেটাকে পুরো বাস্তব রেখায় ছড়িয়ে দেয়। এখন linear predictor স্বচ্ছন্দে বসতে পারে।

### তিনটি সমতুল্য রূপ — তিনটাই লিখতে পারতে হবে

$$\text{(সম্ভাবনা)}\quad \pi=\frac{e^{\eta}}{1+e^{\eta}}$$
$$\text{(odds)}\quad \frac{\pi}{1-\pi}=e^{\eta}$$
$$\text{(log-odds)}\quad \boxed{\log\frac{\pi}{1-\pi}=\eta=\boldsymbol{x}'\boldsymbol\beta} \quad\leftarrow \textbf{এটাই link}$$

**তৃতীয় রূপটাই একে "generalised LINEAR model" বানায়** — মডেল linear, তবে $\pi$-তে নয়, $\log\frac{\pi}{1-\pi}$-তে।

### 🎯 উত্তর ২ — $\hat\beta_j$ কী বোঝায় [Exam 2025 Ex 1(h)]

> 🔴 **প্রশ্ন:** "$x_j$ এক বাড়লে $P(y=1)$ $\hat\beta_j$ পরিমাণ বাড়ে" → **FALSE**

**কেন?** $\hat\beta_j$ **ঢোকার সময়ের** সংখ্যা (log-odds scale), **বেরোনোর সময়ের** নয় (probability scale)। মেশিন মাঝপথে সব চেপে দেয়।

| Scale | $x_j$ এক বাড়লে |
|---|---|
| **log-odds** | $+\hat\beta_j$ (ধ্রুবক, সঠিক) |
| ⭐ **odds** | $\times\exp(\hat\beta_j)$ ← **odds ratio, এটাই বলবে** |
| **probability** | $\hat\beta_j\,\pi(1-\pi)$ ← **ধ্রুবক নয়!** |

$\pi(1-\pi)$ সর্বোচ্চ হয় $\pi=0.5$-এ (মান ০.২৫), আর প্রান্তের দিকে শূন্যের দিকে যায় — ঠিক বাথটাবের মতো।

> 🔑 **যদি $\hat\beta_j$ সত্যিই ধ্রুবক probability পরিবর্তন হতো, তাহলে তুমি আবার সেই ভাঙা linear model-এ ফিরে যেতে। Link-এর পুরো উদ্দেশ্যই হলো probability effect ধ্রুবক না হওয়া।**

**শুধু চিহ্ন (sign) নির্ভরযোগ্যভাবে বহাল থাকে** — logistic function সর্বদা বর্ধমান।

### মুখস্থ বাক্য

> **"multiplies the odds by $\exp(\hat\beta_j)$"**

| $\hat\beta$ | $\exp(\hat\beta)$ | অর্থ |
|---|---|---|
| $0.69$ | $2.00$ | odds **দ্বিগুণ** |
| $-0.69$ | $0.50$ | odds **অর্ধেক** |
| $0.028$ | $1.028$ | odds **২.৮% বাড়ে** |

**Probit:** একই ধারণা, logistic-এর বদলে $\Phi$। ফলাফল প্রায় অভিন্ন; logit জেতে **ব্যাখ্যাযোগ্যতায়** (odds ratio)।

**দুটোই maximum likelihood দিয়ে fit হয়, OLS দিয়ে নয়।**

---

# 🎯 পরীক্ষার আগের চেকলিস্ট

- [ ] **$c$ level ⟹ $c-1$ dummy** (সব $c$ দিলে $\boldsymbol{X}'\boldsymbol{X}$ singular)
- [ ] **Reference = output-এ যে level নেই**
- [ ] **Non-reference তুলনা ⟹ coefficient বিয়োগ**
- [ ] **যা মিল তা কাটে** — পার্থক্য সরাসরি বের করো
- [ ] **Interaction ⟹ অসমান্তরাল রেখা**; দুই সারির ছক লেখো
- [ ] **ভেরিয়েবল দুই পদে ⟹ differentiate**, একটা coefficient বলবে না
- [ ] **Polynomial: $\hat\beta_2<0$ = মাধ্যাকর্ষণ**; চূড়া $-\hat\beta_1/(2\hat\beta_2)$
- [ ] 🔴 **Logit $\hat\beta$ = log-odds; $\exp(\hat\beta)$ = odds ratio; probability effect ধ্রুবক নয়**
- [ ] **Logit-এর ৪টি কারণ** — "$[0,1]$-এর বাইরে যায়" দিয়ে শুরু
- [ ] **"expected" / "on average"** লিখেছ?
- [ ] **"associated with"**, "causes" নয়
- [ ] **"holding all other covariates fixed"**
- [ ] **একক লিখেছ?** (dollars per hour, percent)
- [ ] **Reference category-র নাম বলেছ?**
