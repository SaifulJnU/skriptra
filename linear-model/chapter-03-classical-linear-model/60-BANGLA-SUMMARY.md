# অধ্যায় ৩ — ক্লাসিক্যাল লিনিয়ার মডেল (বাংলা সারাংশ)

> **এই অধ্যায়ই পরীক্ষা।** মোট নম্বরের প্রায় **৮৫%** এখান থেকে আসে। বাকি সব ওয়ার্ম-আপ।
>
> এই ফাইলটা ইংরেজি নোটগুলোর **বিকল্প নয়** — পরিপূরক। ইংরেজিতে পড়ার পর ধারণা পাকা করার জন্য, অথবা কোনো জায়গায় আটকে গেলে দ্রুত বোঝার জন্য ব্যবহার করো। **পরীক্ষার উত্তর অবশ্যই ইংরেজিতে লিখতে হবে**, তাই টেকনিক্যাল টার্মগুলো ইংরেজিতেই রেখে দিয়েছি।

---

## ৩.১ — মডেলের সংজ্ঞা

### মূল সমীকরণ

$$y_i = \beta_0+\beta_1x_{i1}+\dots+\beta_kx_{ik}+\varepsilon_i \qquad\Longleftrightarrow\qquad \boldsymbol{y}=\boldsymbol{X}\boldsymbol\beta+\boldsymbol\varepsilon$$

| প্রতীক | মাত্রা (dimension) | ব্যাখ্যা |
|---|---|---|
| $\boldsymbol{y}$ | $n\times1$ | response ভেক্টর |
| $\boldsymbol{X}$ | $n\times p$ | **design matrix** |
| $\boldsymbol\beta$ | $p\times1$ | প্যারামিটার |
| $\boldsymbol\varepsilon$ | $n\times1$ | error |

**$\boldsymbol{X}$ মনে রাখার নিয়ম:** প্রতি **পর্যবেক্ষণের জন্য একটি সারি**, প্রতি **প্যারামিটারের জন্য একটি কলাম**, এবং **প্রথম কলাম পুরোটাই ১** (intercept-এর জন্য)।

$$p = k+1 \qquad (k = \text{covariate সংখ্যা})$$

> ⚠️ **সবচেয়ে বড় নোটেশন ফাঁদ:** বইতে $p$ মানে **প্যারামিটার সংখ্যা** ($k+1$), কিন্তু কিছু পরীক্ষার প্রশ্নপত্রে $p$ মানে **covariate সংখ্যা**।
>
> 🛟 **নিরাপদ উপায়:** কখনোই "$n-p$" মুখস্থ করবে না। বরং মনে রাখো —
> **residual df = $n$ − (যতগুলো $\beta$ estimate করেছ, intercept সহ) = $n-k-1$**

---

### Hat matrix ও residual

$$\hat{\boldsymbol{y}}=\boldsymbol{H}\boldsymbol{y},\qquad \boldsymbol{H}=\boldsymbol{X}(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'$$

$\boldsymbol{H}$ symmetric, idempotent, এবং $\text{tr}(\boldsymbol{H})=p$।

**$n-p$ কোথা থেকে আসে?** $\boldsymbol{y}$ থাকে $n$-মাত্রার জায়গায়; $\hat{\boldsymbol{y}}$ আটকে থাকে $p$-মাত্রার subspace-এ; তাই residual শুধু বাকি **$n-p$** মাত্রায় নড়তে পারে। **ফিট করার পরে যত মাত্রা বাকি থাকে, সেটাই degrees of freedom।**

**Leverage:** $h_{ii}=\boldsymbol{x}_i'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_i$, গড় $p/n$, $h_{ii}>2p/n$ হলে বেশি।

### Intercept থাকলে যা স্বয়ংক্রিয়ভাবে সত্য

$$\sum_i\hat\varepsilon_i=0,\qquad \sum_i x_{ij}\hat\varepsilon_i=0,\qquad \bar{\hat y}=\bar y$$

> ⚠️ এগুলো **গঠনগতভাবেই** সত্য — ভালো মডেলের প্রমাণ **নয়**। খারাপ মডেলেও residual-এর যোগফল শূন্য হয়। তাই diagnostics-এ residual-এর **গড়** না দেখে **প্যাটার্ন** দেখা হয়।

---

## ৩.১.২ — অনুমান (Assumptions) ⭐

| # | অনুমান | গাণিতিক রূপ | কী দেয় |
|---|---|---|---|
| **A1** | Linearity / সঠিক specification | $E(\boldsymbol{y}\mid\boldsymbol{X})=\boldsymbol{X}\boldsymbol\beta$ | **unbiased** |
| **A2** | Error-এর গড় শূন্য | $E(\varepsilon_i)=0$ | unbiased |
| **A3** | **Homoscedasticity** | $\text{Var}(\varepsilon_i)=\sigma^2$ | efficiency |
| **A4** | **Autocorrelation নেই** | $\text{Cov}(\varepsilon_i,\varepsilon_j)=0$ | efficiency |
| **A5** | **Full column rank** | $\text{rank}(\boldsymbol{X})=p$ | **অস্তিত্ব ও অদ্বিতীয়তা** |
| **A6** | **Normality** (অতিরিক্ত) | $\boldsymbol\varepsilon\sim N(\boldsymbol0,\sigma^2\boldsymbol{I})$ | **exact** test |

A3 + A4 একসাথে: $\;\text{Cov}(\boldsymbol\varepsilon)=\sigma^2\boldsymbol{I}_n$

### 🔑 স্তরবিন্যাস — এটাই আসল কথা

```
A1, A2, A5  ───────►  β̂ আছে, অদ্বিতীয়, এবং UNBIASED
     │
     │  + A3, A4
     ▼
GAUSS–MARKOV ──────►  β̂ হলো BLUE
     │
     │  + A6 (normality)
     ▼
EXACT INFERENCE ───►  t-test, F-test, CI সব EXACT
                      এবং OLS = ML
```

**তিনটি বাক্য যা নম্বর এনে দেবে:**
1. **Unbiasedness-এর জন্য শুধু A1, A2, A5 লাগে** — homoscedasticity, independence বা normality লাগে **না**।
2. **BLUE-এর জন্য A1–A5 লাগে** — **normality লাগে না**।
3. **Exact test-এর জন্য ছয়টাই লাগে।**

### 🔴 অনুমান ভাঙলে কী হয় — এই ছকটা মুখস্থ করো

| যা ভেঙেছে | Unbiased? | BLUE? | se ঠিক? | Test বৈধ? |
|---|---|---|---|---|
| **A1** linearity | ❌ **না** | ❌ | ❌ | ❌ |
| **A3** homoscedasticity | ✅ হ্যাঁ | ❌ না | ❌ না | ❌ না |
| **A4** independence | ✅ হ্যাঁ | ❌ না | ❌ না | ❌ না |
| **A6** normality | ✅ হ্যাঁ | ✅ **হ্যাঁ** | ✅ হ্যাঁ | ⚠️ শুধু asymptotically |
| **Near-multicollinearity** | ✅ হ্যাঁ | ✅ **হ্যাঁ** | ✅ বৈধ কিন্তু **বড়** | ✅ বৈধ, power কম |

> 🔑 **শুধু A1 (ভুল specification) $\hat{\boldsymbol\beta}$-কে biased করে।** A3 ও A4 efficiency ও inference নষ্ট করে, bias নয়।

### 🔴 Heteroscedasticity-এর টেমপ্লেট উত্তর (মুখস্থ করো)

> *OLS estimator remains **unbiased** and **consistent**, since unbiasedness requires only correct specification, zero-mean errors and full rank. However it is no longer **efficient** — the Gauss–Markov theorem no longer applies, so OLS is not BLUE, and a weighted least squares estimator would have smaller variance. In addition, the usual standard errors are biased, so the resulting t-tests, F-tests and confidence intervals are invalid.*

*(Exam Summer 2025 Ex 4(e)-এ ঠিক এটাই চাওয়া হয়েছে — "revenue-এর variation employee সংখ্যার সাথে বাড়ে"। দুটো মূল শব্দ: **unbiased** ও **inefficient**।)*

### Perfect বনাম Near multicollinearity — গুলিয়ে ফেলো না

| | **Perfect** | **Near** |
|---|---|---|
| A5 ভাঙে? | ✅ হ্যাঁ | ❌ **না** |
| $(\boldsymbol{X}'\boldsymbol{X})^{-1}$ | নেই | আছে |
| $\hat{\boldsymbol\beta}$ | **identified নয়** | unbiased, BLUE |
| প্রভাব | সমাধান নেই | variance **ফুলে যায়** |
| পরিমাপ | — | **VIF** $=1/(1-R_j^2)$ |

> ⚠️ **VIF ≈ 1 মানে কোনো সমস্যা নেই** (একদম আদর্শ)। সমস্যা শুরু VIF > 5–10 থেকে। পরীক্ষায় এই স্কেল উল্টে দিয়ে প্রশ্ন করা হয়।

---

## ৩.১.৩ — Covariate-এর প্রভাব মডেল করা

### Dummy variable — 🔴 প্রতি বছর আসে

> **$c$ টি level ⟹ $c-1$ টি dummy।** বাদ পড়া level-টিই **reference category**।

**কেন $c-1$?** সব $c$ টি dummy-র যোগফল সবসময় ১ = intercept কলাম। মানে কলামগুলো **linearly dependent** ⟹ $\boldsymbol{X}'\boldsymbol{X}$ **singular** ⟹ **অদ্বিতীয় OLS সমাধান নেই**। একেই বলে **dummy variable trap** — এবং এই কারণেই A5 অনুমানটা দরকার।

**নিয়মগুলো:**
- প্রতিটি dummy coefficient **reference-এর সাথে তুলনা**
- দুটো non-reference level তুলনা করতে **coefficient বিয়োগ করো**
- Reference = R output-এ **যে level-টা নেই**

### Polynomial

$$y=\beta_0+\beta_1x+\beta_2x^2+\varepsilon \qquad\Longrightarrow\qquad \frac{\partial E(y)}{\partial x}=\beta_1+2\beta_2x$$

⚠️ **প্রভাব $\hat\beta_1$ নয়!** সর্বোচ্চ/সর্বনিম্ন বিন্দু: $\;x^*=-\hat\beta_1/(2\hat\beta_2)$

**কেন $(\text{age}-48)^2$ লেখা হয়, শুধু $\text{age}^2$ নয়?** দুটো কারণ:
1. **ব্যাখ্যাযোগ্যতা** — ৪৮ বছরে term-টা শূন্য হয়, তাই বাকি coefficient-গুলো একজন বাস্তব মানুষকে বর্ণনা করে (বয়স ০ নয়)
2. **কম multicollinearity** — ১৮–৮০ পরিসরে age ও age² এর correlation ০.৯৮-এর বেশি; centre করলে প্রায় uncorrelated হয়ে যায়

### Interaction

$$y=\beta_0+\beta_1x+\beta_2D+\beta_3(xD)+\varepsilon$$

| গ্রুপ | Intercept | Slope |
|---|---|---|
| $D=0$ | $\beta_0$ | $\beta_1$ |
| $D=1$ | $\beta_0+\beta_2$ | $\beta_1+\beta_3$ |

> **শুধু dummy = সমান্তরাল রেখা (সরে যায়)। Dummy + interaction = অসমান্তরাল রেখা (সরে যায় ও হেলে যায়)।**

⚠️ **একটি ভেরিয়েবল দুই জায়গায় থাকলে (polynomial বা interaction), তার coefficient একা ব্যাখ্যা করা যাবে না — differentiate করতে হবে।**

### Restricted model বানানো

$H_0$ কে মডেলে বসিয়ে নতুন মডেল বের করা। **পদ্ধতি: প্রতিস্থাপন → মুক্ত প্যারামিটার অনুযায়ী পদ সাজানো → প্যারামিটারবিহীন পদ বামে পাঠানো।**

উদাহরণ ($H_0:\beta_1=\beta_2+1$):
$$y_i=\beta_0+(\beta_2+1)x_{1i}+\beta_2x_{2i}+\varepsilon_i$$
$$\boxed{\;y_i-x_{1i}=\beta_0+\beta_2(x_{1i}+x_{2i})+\varepsilon_i\;}$$

$y_i-x_{1i}$ কে $(x_{1i}+x_{2i})$ এর উপর regress করো। ৩টার বদলে ২টা প্যারামিটার ⟹ **$r=1$**।

---

## ৩.২.১ — OLS ⭐ সবচেয়ে নিশ্চিত নম্বর

### যা minimise করা হয়

$$S(\boldsymbol\beta)=(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)$$

> 💰 **শুধু এই বাক্যটা লেখার জন্যই ১ নম্বর:** *"OLS chooses $\hat{\boldsymbol\beta}$ to minimise the residual sum of squares."*

### 🔴 চার লাইনের derivation — খালি কাগজে ১০ বার অনুশীলন করো

$$S(\boldsymbol\beta)=\boldsymbol{y}'\boldsymbol{y}-2\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{y}+\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{X}\boldsymbol\beta$$
*(মাঝের দুটো পদ এক হয়ে যায় কারণ $\boldsymbol{y}'\boldsymbol{X}\boldsymbol\beta$ একটি **scalar**, আর scalar তার নিজের transpose-এর সমান)*

$$\frac{\partial S}{\partial\boldsymbol\beta}=-2\boldsymbol{X}'\boldsymbol{y}+2\boldsymbol{X}'\boldsymbol{X}\boldsymbol\beta\overset{!}{=}\boldsymbol0$$

$$\Longrightarrow\ \boldsymbol{X}'\boldsymbol{X}\hat{\boldsymbol\beta}=\boldsymbol{X}'\boldsymbol{y}\quad\textbf{(normal equations)}\ \Longrightarrow\ \boxed{\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}}$$

$$\frac{\partial^2S}{\partial\boldsymbol\beta\partial\boldsymbol\beta'}=2\boldsymbol{X}'\boldsymbol{X}>0\ \Longrightarrow\ \text{অদ্বিতীয় minimum}$$

### জ্যামিতিক অর্থ

$\hat{\boldsymbol{y}}$ হলো $\boldsymbol{y}$-এর **orthogonal projection** ($\boldsymbol{X}$-এর column space-এ)। Residual লম্ব। এই একটি ছবি থেকেই আসে: normal equations, $\text{SST}=\text{SSE}+$explained SS (পিথাগোরাস), এবং $n-p$।

### Simple regression

$$\hat\beta_1=\frac{\widehat{\text{Cov}}(x,y)}{\widehat{\text{Var}}(x)}=r\frac{s_y}{s_x},\qquad \boxed{\hat\beta_0=\bar y-\hat\beta_1\bar x}$$

> 💰 রেখা সবসময় $(\bar x,\bar y)$ দিয়ে যায় — R output-এ হারানো intercept ১৫ সেকেন্ডে বের করা যায়।

### মূল ফলাফল

$$\hat{\boldsymbol\beta}=\boldsymbol\beta+(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol\varepsilon \qquad\text{(সব কিছুর ভিত্তি)}$$
$$E(\hat{\boldsymbol\beta})=\boldsymbol\beta,\qquad \boxed{\text{Cov}(\hat{\boldsymbol\beta})=\sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}}$$
$$\widehat{\text{se}}(\hat\beta_j)=\hat\sigma\sqrt{\left[(\boldsymbol{X}'\boldsymbol{X})^{-1}\right]_{jj}}\qquad\textbf{(শুধু diagonal!)}$$

> ⚠️ **সূচক ভুল করা যাবে না।** Matrix ছাপা হয় সারি ১…p হিসেবে, কিন্তু index হলো $\beta_0,\dots,\beta_k$। **$\beta_1$ হলো দ্বিতীয় diagonal element।** হিসাব শুরুর আগে matrix-এর পাশে $\beta_0,\beta_1,\dots$ লিখে নাও।

---

## ৩.২.২ — Error variance

$$\hat\sigma^2=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n-p}\ \textbf{(unbiased)} \qquad\qquad \hat\sigma^2_{ML}=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n}\ \textbf{(ML)}$$

### 🔴 কোনটা কোথায় — এই ছকটাই আসল

| কাজ | ভাজক |
|---|---|
| se, t-test, F-test, CI, prediction interval, standardised residual | **$n-p$** |
| R output-এর "residual standard error" | **$n-p$** |
| প্রশ্নে "REML" / "restricted maximum likelihood" | **$n-p$** |
| **AIC এবং BIC** | **$n$** |

**কেন $n-p$?** Residual-গুলো তো $p$ টা প্যারামিটার দিয়ে **ছোট করার জন্যই** optimise করা হয়েছে — তাই সেগুলো স্বাভাবিকভাবেই খুব ছোট। Normal equations $p$ টা কঠোর শর্ত চাপায়, তাই মাত্র $n-p$ টা residual সত্যিকারের স্বাধীন।

---

## ৩.২.৩ — Gauss–Markov ও BLUE ⭐ [৪ নম্বর]

> WS 23/24 Ex 2(a): *"Gauss–Markov theorem ও অনুমানগুলো বর্ণনা করো"* — marking key বলছে **"১ নম্বর প্রতিটি অনুমানের জন্য"**। এটা **মুখস্থের প্রশ্ন**। ছেড়ে দিও না।

### বিবৃতি

> A1–A5 সত্য হলে, OLS estimator $\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$ হলো $\boldsymbol\beta$-এর **Best Linear Unbiased Estimator (BLUE)**।

### BLUE — অক্ষর ধরে ধরে

| অক্ষর | অর্থ | সতর্কতা |
|---|---|---|
| **B** | Best = **সর্বনিম্ন variance** | শুধু L∩U শ্রেণির মধ্যে |
| **L** | Linear = $\hat{\boldsymbol\beta}=\boldsymbol{A}\boldsymbol{y}$ | এটা প্রতিযোগীদের উপর **শর্ত**, গুণ নয় |
| **U** | Unbiased = $E(\hat{\boldsymbol\beta})=\boldsymbol\beta$ | এটাও শর্ত |
| **E** | Estimator = নিয়ম, সংখ্যা নয় | — |

> 🔴 **Normality Gauss–Markov-এর অনুমান নয়!** এটা লিখলেই ৩ নম্বরের উত্তর ৪ নম্বর হয়ে যায়।
>
> 🔴 "linear" আর "unbiased" — দুটো শর্তের একটাও বাদ দিলে দাবিটা **মিথ্যা** হয়ে যায়। **Biased** estimator (ridge, lasso) কিংবা **non-linear** estimator-এর variance কম হতেই পারে।

### OLS = ML (normality-র অধীনে)

Gaussian log-likelihood-এ $\boldsymbol\beta$ শুধু $-\frac{1}{2\sigma^2}(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)$ পদে আছে, **ঋণাত্মক চিহ্ন সহ**। তাই maximise করা = SSE minimise করা।

$$\hat{\boldsymbol\beta}_{ML}=\hat{\boldsymbol\beta}_{LS} \qquad\text{কিন্তু}\qquad \hat\sigma^2_{ML}\neq\hat\sigma^2_{LS}$$

---

## ৩.৩ — Hypothesis testing ⭐⭐⭐ সবচেয়ে বেশি পরীক্ষিত

### t-test

$$\frac{\hat\beta_j-\beta_j}{\widehat{\text{se}}(\hat\beta_j)}\sim t_{n-p} \qquad\Longrightarrow\qquad t=\frac{\hat\beta_j-c}{\widehat{\text{se}}(\hat\beta_j)}$$

Significance test-এ $c=0$। **⚠️ অন্য ক্ষেত্রে $-c$ ভুলে যেও না।**

$$\text{reject if } |t|>t_{n-p}(1-\tfrac\alpha2)$$

### R output পূরণ করার সূত্র

$$t=\frac{\hat\beta_j}{\widehat{\text{se}}} \quad\Longleftrightarrow\quad \hat\beta_j=t\times\widehat{\text{se}} \quad\Longleftrightarrow\quad \widehat{\text{se}}=\frac{\hat\beta_j}{t}$$

তিনটার দুটো সবসময় দেওয়া থাকে।

### 🔴 $\boldsymbol{C}\boldsymbol\beta=\boldsymbol{d}$ বানানো

**তিন ধাপ:**
1. প্রতিটি restriction সাজাও: **সব $\beta$ বামে, ধ্রুবক ডানে**
2. প্রতি restriction-এর জন্য **$\boldsymbol{C}$-এর একটি সারি**; $\beta_0,\beta_1,\dots,\beta_k$ ক্রমে coefficient বসাও
3. $\boldsymbol{C}$ হবে $r\times p$; $r$ = **স্বাধীন সমীকরণের সংখ্যা**

> 🔴 **সবচেয়ে বড় ফাঁদ: $r$ = সমীকরণের সংখ্যা, উল্লেখিত $\beta$-র সংখ্যা নয়।**
>
> $H_0:\beta_1=-\beta_2+\beta_3$ — এখানে **তিনটি** $\beta$ আছে, কিন্তু **একটি** সমীকরণ ⟹ $r=1$, $r=3$ নয়।
> *(Exam Summer 2025 Ex 1(i) ঠিক এই ফাঁদটাই পাতে — উত্তর **FALSE**।)*

> ⚠️ **$\boldsymbol{C}$-তে $\beta_0$-এর কলাম ভুলবে না।** $\boldsymbol{C}$-এর কলাম সংখ্যা **$p$**, $k$ নয়।

---

## ৩.৩.১ — F-test

$$\boxed{\;F=\frac{(\text{SSE}_{H_0}-\text{SSE})/r}{\text{SSE}/(n-p)}\sim F_{r,\,n-p}\;}$$

$R^2$ দেওয়া থাকলে:
$$F=\frac{(R^2-R^2_{H_0})/r}{(1-R^2)/(n-p)};\qquad \text{সামগ্রিক test: } F=\frac{R^2/k}{(1-R^2)/(n-p)}$$

**ধারণা:** $H_0$ চাপালে fit কখনো ভালো হতে পারে না, তাই $\text{SSE}_{H_0}\geq\text{SSE}$। প্রশ্ন হলো — **restriction চাপানোর খরচ কি noise-এর তুলনায় বেশি?**

### 🔴 কোন version কখন

| প্রশ্নে যা দেওয়া | যা ব্যবহার করবে |
|---|---|
| $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}$ **এবং** $\hat{\boldsymbol\varepsilon}'_{H_0}\hat{\boldsymbol\varepsilon}_{H_0}$ | **SSE** version |
| শুধু $R^2$, আর $H_0$ = "সব slope শূন্য" | **$R^2$** version |

### 🔴 Quantile-এর নিয়ম — সবচেয়ে সাধারণ ভুল

| Test | $\alpha=0.05$ হলে |
|---|---|
| দুই-পার্শ্বীয় t-test | **0.975** |
| 95% CI | **0.975** |
| 99% CI | **0.995** |
| **F-test** | **0.95** ← এক-পার্শ্বীয়! |

> 🔑 **t ও CI দুই-পার্শ্বীয় ⟹ $1-\alpha/2$। F এক-পার্শ্বীয় ⟹ $1-\alpha$।**

**যাচাই:** $F\geq0$ সবসময়। $r=1$ হলে $\sqrt{F}=|t|$।

**সিদ্ধান্ত লেখার ভাষা:** reject করলে বলবে **"অন্তত একটি"** restriction ভাঙে — "সবগুলো" নয়। আর কখনো "accept $H_0$" লিখবে না, লিখবে **"fail to reject"**।

---

## ৩.৩.২ — Confidence ও Prediction Interval

### Coefficient-এর CI

$$\hat\beta_j\pm t_{n-p}\!\left(1-\tfrac\alpha2\right)\cdot\widehat{\text{se}}(\hat\beta_j)$$

### 🔑 CI–test দ্বৈততা

> **CI-তে $c$ থাকলে ⟹ $H_0:\beta_j=c$ reject করা যায় না।**
> **শূন্য CI-র বাইরে থাকলে ⟹ coefficient significant।**

*(Exam Summer 2025 Ex 1(k) এই যুক্তি উল্টে দিয়ে প্রশ্ন করে — উত্তর **FALSE**।)*

### 🔴 CI বনাম Prediction Interval — পার্থক্য শুধু "$1+$"

| | **CI (গড়ের জন্য)** | **Prediction Interval (একজনের জন্য)** |
|---|---|---|
| লক্ষ্য | $E(y_0)$ — নির্দিষ্ট | $y_0$ — এলোমেলো |
| মূলের ভেতরে | $\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0$ | $\mathbf{1}+\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0$ |
| প্রস্থ | সরু | **চওড়া** |
| $n\to\infty$ হলে | **শূন্যে** নেমে আসে | $\pm t\hat\sigma$-এ থামে, **কখনো শূন্য নয়** |

### 🔑 কেন "$1+$"?

নতুন একজন মানুষের **নিজস্ব $\varepsilon_0$** আছে — তার আলোচনার দক্ষতা, তার নিয়োগকর্তা, তার ভাগ্য।

> **তুমি যতই ডেটা জোগাড় করো, যাকে দেখোনি তার সম্পর্কে ডেটা কিছু বলবে না।**
>
> ৪০ বছর বয়সীদের **গড়** মজুরি যত খুশি নিখুঁতভাবে জানা যায়। কিন্তু **একজন নির্দিষ্ট** ৪০ বছর বয়সীর মজুরি কখনোই নিখুঁতভাবে জানা যাবে না। এই "১" সেই গাণিতিক স্বীকারোক্তি যে **ব্যক্তি ≠ গড়**।

**কোনটা চাইছে বুঝবে কীভাবে?** "একজন ৫০ বছর বয়সী পুরুষের মজুরি" — একবচন, একজন মানুষ ⟹ **prediction interval**। "৫০ বছর বয়সীদের গড়" ⟹ **CI**।

---

## ৩.৪.১ — Bias–Variance

$$E[(y_0-\hat f)^2]=\underbrace{\sigma^2}_{\text{অপরিহার্য}}+\underbrace{\text{Bias}^2}_{\text{ব্যবস্থাগত ভুল}}+\underbrace{\text{Variance}}_{\text{estimate-এর দোলা}}$$

**জটিলতা বাড়ালে ⟹ bias কমে, variance বাড়ে।** মোট ত্রুটি **U-আকৃতির**।

**Overfitting:** $p=n$ হলে প্রতিটি বিন্দুতে fit হয়, $R^2=1$, কিন্তু নতুন ডেটায় সম্পূর্ণ ব্যর্থ। **মডেল সংকেত নয়, গোলমাল মুখস্থ করেছে।**

> 🔑 **Bias বর্গ হয়ে ঢোকে, variance একরৈখিকভাবে** — তাই সামান্য bias মেনে নিয়ে অনেক variance কমানো লাভজনক হতে পারে। এ কারণেই BLUE শেষ কথা নয়।

---

## ৩.৪.২ — Model Choice Criteria

$$\bar R^2=1-\frac{n-1}{n-p}(1-R^2) \qquad\text{(বড় ভালো)}$$

$$\boxed{\text{AIC}=n\log(\hat\sigma^2)+2(|M|+1)} \qquad \boxed{\text{BIC}=n\log(\hat\sigma^2)+\log(n)(|M|+1)}$$

$$\hat\sigma^2=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{\mathbf{n}}\quad\textbf{(ML — } n \textbf{ দিয়ে ভাগ!)},\qquad \log=\ln$$

**দুটোতেই ছোট মান ভালো।**

### 🔴 তিনটি ফাঁদ

1. AIC/BIC-তে $n-p$ দিয়ে ভাগ করা — **ভুল**, $n$ দিয়ে ভাগ করতে হবে
2. $\log_{10}$ ব্যবহার করা — **natural log** লাগবে
3. **$+1$** ভুলে যাওয়া — $\sigma^2$-ও একটা প্যারামিটার

### 🔴 AIC না BIC কে বেশি শাস্তি দেয়?

$\log(n)>2$ যখন $n>7.39$ — অর্থাৎ **সব বাস্তব ডেটাসেটে**।

> **BIC বেশি শাস্তি দেয় ⟹ ছোট মডেল বাছে।**
> **মনে রাখার কৌশল: B for Bigger penalty.**

| | AIC | BIC |
|---|---|---|
| লক্ষ্য | সেরা **prediction** | **সত্যিকারের** মডেল খোঁজা |
| বাছে | বড় মডেল | ছোট মডেল |

> 💡 **সময় বাঁচানোর কৌশল:** $n\log(\hat\sigma^2)$ পদটা AIC ও BIC-তে **একই**। একবার হিসাব করে দুটো penalty যোগ করো।

> 🔑 **যদি BIC — সবচেয়ে কঠোর সমালোচক — ও বড় মডেলকেই পছন্দ করে, তাহলে উন্নতিটা সত্যিকারের।**

### 🔴 R² সম্পর্কিত ফাঁদ (প্রতি বছর আসে)

- $R^2$ ভেরিয়েবল যোগ করলে **কখনো কমে না** ⟹ "$R^2$ কমবে" = **FALSE**
- SSE ভেরিয়েবল যোগ করলে **কখনো বাড়ে না** ⟹ "RSS বাড়তে পারে" = **FALSE**
- $\bar R^2$ **ঋণাত্মক হতে পারে** ⟹ "কখনো ঋণাত্মক হয় না" = **FALSE**

---

## ৩.৪.৪ — Model Diagnosis

### কেন standardise করতে হয়?

নিখুঁত অনুমানের অধীনেও:
$$\text{Var}(\hat\varepsilon_i)=\sigma^2(1-h_{ii}) \qquad\text{— সমান নয়!}$$

**উচ্চ leverage-এর বিন্দুর residual কৃত্রিমভাবে ছোট হয়ে যায়** — ফলে সত্যিকারের outlier লুকিয়ে যেতে পারে। তাই:

$$r_i=\frac{\hat\varepsilon_i}{\hat\sigma\sqrt{1-h_{ii}}} \qquad |r_i|>2 \text{ লক্ষণীয়},\quad >3 \text{ সম্ভাব্য outlier}$$

### ⭐ চারটি প্লট — মুখস্থ করো

| প্লট | কী ধরে | খারাপ দেখায় |
|---|---|---|
| **Residuals vs Fitted** | **non-linearity**, heteroscedasticity | বাঁকা প্যাটার্ন; **ফানেল** আকৃতি |
| **QQ plot** | **non-normality** | S-আকৃতি; প্রান্তে বাঁক |
| **Scale–Location** | heteroscedasticity (বিশেষভাবে) | ঊর্ধ্বমুখী ঢাল |
| **Residuals vs Leverage** | **influential points** | Cook's D > 0.5 |

> 🔴 **QQ plot-এর বিন্দুগুলো ৪৫° কোণের তির্যক রেখা বরাবর থাকে — অনুভূমিক রেখা নয়।** (WS 22/23-এর একটি প্রশ্ন প্রথম অংশ সঠিক লিখে শেষ অংশে এই ভুলটা ঢুকিয়ে দেয়। **শেষ পর্যন্ত পড়ো।**)

### Leverage ≠ Outlier ≠ Influence

| ধারণা | মানে | পরিমাপ |
|---|---|---|
| **Leverage** | অস্বাভাবিক **covariate** মান | $h_{ii}$ |
| **Outlier** | অস্বাভাবিক **response** | $|r_i|$ বড় |
| **Influence** | fit **সত্যিই বদলে দেয়** | Cook's $D_i$ |

$$D_i=\frac{r_i^2}{p}\cdot\frac{h_{ii}}{1-h_{ii}} \qquad = \text{(outlyingness)}\times\text{(leverage)}$$

> 🔑 **উচ্চ leverage একা কোনো সমস্যা নয়।** দূরে থাকা একটা বিন্দু যদি ঠিক রেখার উপরেই বসে, সে বরং **নির্ভুলতা বাড়ায়**। বিপদ তখনই, যখন তার residual-ও বড় — তখন সে পুরো রেখাটা নিজের দিকে টেনে নেয়।

---

# 🎯 পরীক্ষার আগের রাতের চেকলিস্ট

- [ ] **Residual df = $n$ − (β-র সংখ্যা, intercept সহ) = $n-k-1$**
- [ ] **$r$ = সমীকরণের সংখ্যা**, উল্লেখিত β-র সংখ্যা নয়
- [ ] **t ও CI: $1-\alpha/2$। F: $1-\alpha$।**
- [ ] **AIC/BIC-তে $n$ দিয়ে ভাগ, $\ln$ ব্যবহার, $+1$ যোগ**
- [ ] **BIC-এর penalty বেশি** (B for Bigger)
- [ ] **$c$ level ⟹ $c-1$ dummy**
- [ ] **Prediction interval-এ "$1+$" আছে**
- [ ] **Heteroscedasticity ⟹ unbiased কিন্তু inefficient**
- [ ] **Normality Gauss–Markov-এর জন্য লাগে না**
- [ ] **$R^2$ ও SSE মডেল আকারে একমুখী** (কখনো উল্টো যায় না)
- [ ] **VIF ≈ 1 মানে সমস্যা নেই**
- [ ] **QQ plot = তির্যক রেখা**
- [ ] **"fail to reject" লিখবে, "accept" নয়**
- [ ] **"associated with" লিখবে, "causes" নয়**
- [ ] **"holding all other covariates fixed" — প্রতিবার লিখবে**
- [ ] **সংখ্যার আগে সূত্র লিখবে** (method-এর জন্য নম্বর আছে)
- [ ] **৩ দশমিক ঘর পর্যন্ত round করবে**

---

**৬০ মিনিটে ৬০ নম্বর = ১ নম্বরে ১ মিনিট।** গতি আসে চেনা থেকে, বুদ্ধি থেকে নয়। এই ফোল্ডারের প্রতিটা প্রশ্ন তুমি আগে দেখেছ — পরীক্ষার হলে সেটাই তোমার সবচেয়ে বড় সুবিধা।
