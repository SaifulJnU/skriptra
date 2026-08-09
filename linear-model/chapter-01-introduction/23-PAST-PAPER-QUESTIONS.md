# Ch 1 — PAST-PAPER QUESTIONS (all five papers)

*বাংলা সংস্করণ নিচে আছে → [বাংলায় পড়ো](#অধ্যায়-১--বিগত-বছরের-প্রশ্ন-বাংলা)*

> **Source codes used throughout the three `23-` files:**
> **S25** = Exam Summer 2025 · **LMES** = Linear_model_exam_sheet · **W23** = WiSe 2023/24 · **W22** = RCLM WS 22/23 · **EX20** = Example Exam LiMo 2020
>
> ⚠️ **Notation warning.** LMES and W23 use $p$ = **number of covariates**, so they write $\text{df}=n-p-1$. W22 and the book use $p$ = **number of parameters** = $k+1$, so $\text{df}=n-p$. **Both are the same number.** Count the betas you estimated, including the intercept, and subtract from $n$.

---

## The honest summary

**Chapter 1 has never been examined on its own.** Not one question in five papers asks about Galton, about the rent or credit-scoring data sets, or about the difference between a regressand and a response.

What Chapter 1 *does* supply is the **notation** that every other question is written in — and there, it earns marks indirectly on every single paper. Below are the only past-paper items where Chapter 1 material is the thing actually being tested.

If you're revising and short of time: read this file once, then never come back. Spend the hour on Chapter 3.

---

## Q1 — Dimensions of the design matrix

> **W22, Ex 2(a) [1 Point].** *"How does the first column of the design matrix $\boldsymbol{X}$ look like in this specific example? Also name the dimensions."*
>
> Context: the `bike` data, $n=17{,}379$ observations, R output listing an intercept plus 12 further coefficients.

### Solution

The first column is a **column of ones**, $\boldsymbol{1}_n$ — it is what multiplies $\beta_0$ and so makes it an intercept.

$$\boldsymbol{X}\in\mathbb{R}^{n\times p}=\mathbb{R}^{17379\times13}$$

Count the columns from the **R output**: 1 intercept + 3 season dummies + 1 yr dummy + 4 daytime dummies + 1 holiday dummy + temp + hum + windspeed = **13**.

> **Method:** the number of columns is always "how many coefficient lines does R print." Never try to count from the variable list — categorical variables print $c-1$ lines each.

---

## Q2 — Dimensions of $(\boldsymbol{X}'\boldsymbol{X})^{-1}$

Asked in **two** papers, in identical words:

> **W22, Ex 2(b) [1 Point]** and **LMES, Ex 3(a) [1 Point].** *"What are the dimensions (i.e. number of rows and columns) of $(\boldsymbol{X}'\boldsymbol{X})^{-1}$ in this example?"*

### Solution

$\boldsymbol{X}$ is $n\times p$, so $\boldsymbol{X}'\boldsymbol{X}$ is $(p\times n)(n\times p) = p\times p$, and the inverse has the **same dimensions**:

| Paper | Setting | Answer |
|---|---|---|
| **W22** | 13 coefficient lines | $\boldsymbol{13\times13}$ |
| **LMES** | 7 covariates + intercept | $\boldsymbol{8\times8}$ |

LMES's own key spells the counting out: *"Find $n=50$ from the question, find $p=7$ from the regression table"* — then the matrix is $8\times8$, because their $p$ counts covariates only and the intercept must be added back.

> 🔴 **The trap is $n$.** $(\boldsymbol{X}'\boldsymbol{X})^{-1}$ is **never** $n\times n$ — the $n$'s cancel in the product. A related T/F, "$\text{rk}(\boldsymbol{X}'\boldsymbol{X})=n$", appears in **W22 Ex 1c(iv)** and is **FALSE** for exactly this reason: rank can never exceed the smaller dimension, and $p\ll n$.

---

## Q3 — The model in matrix form

> **LMES, Ex 2(a) [1 Point].** *"Consider the simple linear regression model $Y_i=\beta_0+\beta_1x_i+\varepsilon_i$ where $\varepsilon_i\sim N(0,\sigma^2)$ are independent. Express the model in matrix form, clearly specifying $\boldsymbol{Y}$, $\boldsymbol{X}$, $\boldsymbol\beta$ and $\boldsymbol\varepsilon$."*

### Solution — the official key, verbatim in structure

$$\boldsymbol{Y}=\boldsymbol{X}\boldsymbol\beta+\boldsymbol\varepsilon:\qquad
\begin{pmatrix}Y_1\\Y_2\\\vdots\\Y_n\end{pmatrix}
=\begin{pmatrix}1&x_1\\1&x_2\\\vdots&\vdots\\1&x_n\end{pmatrix}
\begin{pmatrix}\beta_0\\\beta_1\end{pmatrix}
+\begin{pmatrix}\varepsilon_1\\\varepsilon_2\\\vdots\\\varepsilon_n\end{pmatrix}$$

**Dimensions:** $\boldsymbol{Y}$ is $n\times1$, $\boldsymbol{X}$ is $n\times2$, $\boldsymbol\beta$ is $2\times1$, $\boldsymbol\varepsilon$ is $n\times1$.

> **1 free mark, 40 seconds.** Write the column of ones. Every year somebody writes $\boldsymbol{X}$ as a single column of $x_i$ and loses it.

---

## Q4 — Is this even a linear model?

> **W22, Ex 1b(i) [T/F].** *"The relation $y=\exp(\beta_0+\beta_1x_1+\dots+\beta_kx_k+\varepsilon)$ cannot be analysed within the linear regression framework."*

### **FALSE.**

Take logs: $\log y=\beta_0+\beta_1x_1+\dots+\beta_kx_k+\varepsilon$ — a linear model in the transformed response.

> 🔑 **"Linear" means linear in the parameters $\boldsymbol\beta$, not linear in the variables.** This one sentence is the entire content of Chapter 1's contribution to the exam, and it also covers polynomials ($x^2$ is fine), logs, and interactions. What would break linearity is a parameter inside a non-linear function, e.g. $y=\beta_0 x^{\beta_1}+\varepsilon$.

---

## What Chapter 1 actually buys you

Four small items across five papers, worth roughly **1–2 marks**. But the notation underneath them is loaded into every Chapter 3 question:

| Symbol | You must never hesitate on |
|---|---|
| $\boldsymbol{X}$ | $n\times p$: one **row per observation**, one **column per parameter** |
| $\boldsymbol{X}'\boldsymbol{X}$ | $p\times p$ — and so is its inverse |
| $p$ vs $k$ | Read the paper's own definition, then **count the betas** |
| $x_i$ vs $x_j$ vs $x_{ij}$ | observation $i$ · covariate $j$ · covariate $j$ of observation $i$ |
| "linear" | in the **parameters** |

→ Full treatment in `03-notes-1.3-notation.md` §2. Chapter 2's file is where the marks start.

---
---

# অধ্যায় ১ — বিগত বছরের প্রশ্ন (বাংলা)

> টেকনিক্যাল শব্দ, ফাইলের নাম, সূত্র আর পরীক্ষার হুবহু উদ্ধৃতি ইংরেজিতেই রেখেছি — **পরীক্ষার উত্তর ইংরেজিতে লিখতে হবে**।
>
> **সোর্স কোড:** **S25** = Exam Summer 2025 · **LMES** = Linear_model_exam_sheet · **W23** = WiSe 2023/24 · **W22** = RCLM WS 22/23 · **EX20** = Example Exam LiMo 2020
>
> ⚠️ **নোটেশন সতর্কতা।** LMES আর W23-এ $p$ = **covariate-এর সংখ্যা**, তাই তারা লেখে $\text{df}=n-p-1$। W22 আর বইয়ে $p$ = **parameter-এর সংখ্যা** $=k+1$, তাই $\text{df}=n-p$। **দুটোই একই সংখ্যা।** যতগুলো beta estimate করেছ (intercept সহ) গোনো, আর $n$ থেকে বাদ দাও।

---

## সোজা কথা

**অধ্যায় ১ কখনো একা পরীক্ষায় আসেনি।** পাঁচটা পেপারের একটা প্রশ্নেও Galton, বাড়িভাড়া বা credit-scoring ডেটাসেট, কিংবা regressand আর response-এর পার্থক্য জিজ্ঞেস করা হয়নি।

অধ্যায় ১ যা দেয় তা হলো **নোটেশন** — বাকি প্রতিটা প্রশ্ন সেই নোটেশনেই লেখা, আর সেখান থেকে প্রতিটা পেপারে পরোক্ষভাবে নম্বর আসে। নিচে শুধু সেই আইটেমগুলো আছে যেখানে অধ্যায় ১-এর বিষয়বস্তুই সরাসরি পরীক্ষা করা হয়েছে।

রিভিশনে সময় কম থাকলে: **এই ফাইলটা একবার পড়ো, আর ফিরে এসো না।** ঘণ্টাটা অধ্যায় ৩-এ দাও।

---

## প্রশ্ন ১ — Design matrix-এর dimension

> **W22, Ex 2(a) [১ নম্বর].** *"How does the first column of the design matrix $\boldsymbol{X}$ look like in this specific example? Also name the dimensions."*
>
> প্রেক্ষাপট: `bike` ডেটা, $n=17{,}379$ observation, R output-এ intercept সহ আরও ১২টা coefficient।

### সমাধান

প্রথম কলামটা **এক-এর কলাম**, $\boldsymbol{1}_n$ — এটাই $\beta_0$-কে গুণ করে, আর এজন্যই সেটা intercept হয়।

$$\boldsymbol{X}\in\mathbb{R}^{n\times p}=\mathbb{R}^{17379\times13}$$

কলাম গোনো **R output থেকে**: ১ intercept + ৩ season dummy + ১ yr dummy + ৪ daytime dummy + ১ holiday dummy + temp + hum + windspeed = **১৩**।

> **পদ্ধতি:** কলামের সংখ্যা মানে সবসময় "R কয়টা coefficient-এর লাইন ছাপিয়েছে"। **ভেরিয়েবলের তালিকা থেকে কখনো গুনো না** — প্রতিটা categorical ভেরিয়েবল $c-1$ টা করে লাইন ছাপে।

---

## প্রশ্ন ২ — $(\boldsymbol{X}'\boldsymbol{X})^{-1}$-এর dimension

**দুটো পেপারে হুবহু একই ভাষায় এসেছে:**

> **W22, Ex 2(b) [১ নম্বর]** এবং **LMES, Ex 3(a) [১ নম্বর].** *"What are the dimensions (i.e. number of rows and columns) of $(\boldsymbol{X}'\boldsymbol{X})^{-1}$ in this example?"*

### সমাধান

$\boldsymbol{X}$ হলো $n\times p$, তাই $\boldsymbol{X}'\boldsymbol{X}$ হয় $(p\times n)(n\times p)=p\times p$, আর inverse-এর dimension **একই**:

| পেপার | প্রেক্ষাপট | উত্তর |
|---|---|---|
| **W22** | ১৩টা coefficient লাইন | $\boldsymbol{13\times13}$ |
| **LMES** | ৭টা covariate + intercept | $\boldsymbol{8\times8}$ |

LMES-এর নিজের key-তে গোনাটা বলে দেওয়া আছে: *"Find $n=50$ from the question, find $p=7$ from the regression table"* — তারপর ম্যাট্রিক্সটা $8\times8$, কারণ ওদের $p$ শুধু covariate গোনে, intercept-টা যোগ করে নিতে হয়।

> 🔴 **ফাঁদটা হলো $n$।** $(\boldsymbol{X}'\boldsymbol{X})^{-1}$ **কখনোই** $n\times n$ নয় — গুণফলে $n$ কেটে যায়। এর সাথে মিল রেখে **W22 Ex 1c(iv)**-এ একটা T/F আছে, "$\text{rk}(\boldsymbol{X}'\boldsymbol{X})=n$", আর সেটা ঠিক এই কারণেই **FALSE**: rank কখনো ছোট মাত্রাটাকে ছাড়াতে পারে না, আর $p\ll n$।

---

## প্রশ্ন ৩ — মডেলটা matrix আকারে

> **LMES, Ex 2(a) [১ নম্বর].** *"Consider the simple linear regression model $Y_i=\beta_0+\beta_1x_i+\varepsilon_i$ where $\varepsilon_i\sim N(0,\sigma^2)$ are independent. Express the model in matrix form, clearly specifying $\boldsymbol{Y}$, $\boldsymbol{X}$, $\boldsymbol\beta$ and $\boldsymbol\varepsilon$."*

### সমাধান — অফিসিয়াল key-র গঠন অনুযায়ী

$$\boldsymbol{Y}=\boldsymbol{X}\boldsymbol\beta+\boldsymbol\varepsilon:\qquad
\begin{pmatrix}Y_1\\Y_2\\\vdots\\Y_n\end{pmatrix}
=\begin{pmatrix}1&x_1\\1&x_2\\\vdots&\vdots\\1&x_n\end{pmatrix}
\begin{pmatrix}\beta_0\\\beta_1\end{pmatrix}
+\begin{pmatrix}\varepsilon_1\\\varepsilon_2\\\vdots\\\varepsilon_n\end{pmatrix}$$

**Dimension:** $\boldsymbol{Y}$ হলো $n\times1$, $\boldsymbol{X}$ হলো $n\times2$, $\boldsymbol\beta$ হলো $2\times1$, $\boldsymbol\varepsilon$ হলো $n\times1$।

> **১টা ফ্রি নম্বর, ৪০ সেকেন্ড।** এক-এর কলামটা লিখতে ভুলো না। প্রতি বছর কেউ না কেউ $\boldsymbol{X}$-কে শুধু $x_i$-এর একটা কলাম হিসেবে লেখে আর নম্বরটা হারায়।

---

## প্রশ্ন ৪ — এটা কি আদৌ linear model?

> **W22, Ex 1b(i) [T/F].** *"The relation $y=\exp(\beta_0+\beta_1x_1+\dots+\beta_kx_k+\varepsilon)$ cannot be analysed within the linear regression framework."*

### **FALSE.**

দুই পাশে log নাও: $\log y=\beta_0+\beta_1x_1+\dots+\beta_kx_k+\varepsilon$ — রূপান্তরিত response-এ এটা একটা linear model।

> 🔑 **"Linear" মানে parameter $\boldsymbol\beta$-তে linear, ভেরিয়েবলে linear নয়।** এই একটা বাক্যই পরীক্ষায় অধ্যায় ১-এর পুরো অবদান, আর এটাই polynomial ($x^2$ চলবে), log আর interaction — সবগুলো কভার করে। **linearity ভাঙবে তখনই**, যখন parameter কোনো non-linear ফাংশনের ভেতরে ঢুকে যাবে, যেমন $y=\beta_0 x^{\beta_1}+\varepsilon$।

---

## অধ্যায় ১ আসলে কী দেয়

পাঁচটা পেপারে চারটা ছোট আইটেম, মোট প্রায় **১–২ নম্বর**। কিন্তু এর নিচের নোটেশনটা অধ্যায় ৩-এর প্রতিটা প্রশ্নে লোড হয়ে আছে:

| প্রতীক | যেটাতে কখনো দ্বিধা করবে না |
|---|---|
| $\boldsymbol{X}$ | $n\times p$: প্রতি **observation-এ এক সারি**, প্রতি **parameter-এ এক কলাম** |
| $\boldsymbol{X}'\boldsymbol{X}$ | $p\times p$ — এবং এর inverse-ও তাই |
| $p$ বনাম $k$ | পেপারের নিজের সংজ্ঞাটা পড়ো, তারপর **beta গোনো** |
| $x_i$ বনাম $x_j$ বনাম $x_{ij}$ | observation $i$ · covariate $j$ · observation $i$-এর covariate $j$ |
| "linear" | **parameter**-এ linear |

→ পুরো আলোচনা `03-notes-1.3-notation.md`-এর §২-এ। **নম্বর আসা শুরু হয় অধ্যায় ২-এর ফাইল থেকে।**
