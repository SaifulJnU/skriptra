# Ch 3 — BOOK EXAMPLES & EXERCISE SHEETS (intuition first, then solve)

> ## 📌 First, an honest answer about "book exercises"
>
> **The Fahrmeir book has no end-of-chapter exercise sets.** I searched the whole PDF — no "Exercises" sections, no problem numbers. It teaches through **worked Examples** embedded in the text (Example 3.1, 3.2, …), each analysing a real dataset.
>
> So "the chapter exercises for this course" are really **two** things, and both are covered below:
>
> | Source | What it is | Exam value |
> |---|---|---|
> | **Book Examples 3.1–3.7** | worked analyses of the Munich Rent Index and supermarket data | ⭐⭐ conceptual — shows you *why* each technique exists |
> | **TU Dortmund Sheets 3, 4, 5** | your actual course exercises, on the `Wage` data | ⭐⭐⭐⭐⭐ these ARE the exam, split into pieces |
>
> ---
>
> ## 🧠 How to use this file — the rule
>
> **Every item below has an `🤔 INTUITION` block BEFORE the `✍️ SOLUTION` block.**
>
> Read the intuition. **Close the file. Try it.** Then read the solution.
>
> If you read the solution first you will learn nothing — you'll recognise the steps and mistake recognition for ability. That feeling of "yes, obviously" while reading someone else's working is the single most common way students fail exams they thought they'd prepared for.

---
---

# PART A — THE BOOK'S WORKED EXAMPLES

These aren't questions with answers. They're **demonstrations of why a technique is needed**. For each, I give the situation, the intuition, and what the book concludes.

---

## Example 3.1 — Munich Rent Index: Heteroscedastic Variances

**Situation.** Rent plotted against living area. The scatter is not a uniform band — it **fans out**: small flats have rents tightly clustered, large flats have rents spread across a huge range.

### 🤔 INTUITION — think before reading on

*Why would variance grow with size?*

Ask yourself what determines the rent of a 30 m² flat versus a 200 m² flat. The small flat is a small flat — there isn't much room for it to be luxurious or shabby. The big flat could be a run-down family apartment in the suburbs or a penthouse. **Size creates room for other factors to matter.**

So the spread isn't a data problem. It's the world telling you that the *unexplained* part of rent is genuinely larger for large flats.

**Now: what does that break?** Go back to the assumption list. Which one says "the spread is the same everywhere"?

### ✍️ WHAT THE BOOK SHOWS

It breaks **A3, homoscedasticity**: $\text{Var}(\varepsilon_i)=\sigma^2$ is false; the variance depends on $\text{area}_i$.

**Consequences** (the sentence you'll reuse all exam):
- $\hat{\boldsymbol\beta}$ is still **unbiased** and consistent
- $\hat{\boldsymbol\beta}$ is **no longer BLUE** — weighted least squares does better
- $\text{Cov}(\hat{\boldsymbol\beta})=\sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}$ is **wrong** ⟹ standard errors, $t$-tests, $F$-tests, CIs all invalid

**Remedies:** model $\log(\text{rent})$; weighted least squares; robust standard errors.

> 🔗 **This example IS Exam Summer 2025, Ex 4(e)** — "the variation in revenue grows as the number of employees grows." Same structure, different nouns. The book showed you the answer.

---

## Example 3.3 — Munich Rent Index: Variable Transformation

**Situation.** The effect of living area on rent is clearly **not** linear. The book uses the transformation $f(\text{area}_i) = 1/\text{area}_i$.

### 🤔 INTUITION

*Why would rent-per-flat be non-linear in area, and why $1/\text{area}$ specifically?*

Think about **rent per square metre**. A 40 m² flat and an 80 m² flat: the bigger one costs more in total, but usually *less per square metre*. There are fixed costs — a kitchen, a bathroom, a front door — that every flat needs regardless of size. Those get spread over more m² in a big flat.

If total rent = (fixed part) + (per-m² part × area), then **rent per m²** = fixed/area + constant. That's a $1/\text{area}$ shape.

**The key realisation:** you didn't need a fancy method. You needed to *think about what generates the data*, and the functional form fell out.

### ✍️ WHAT THE BOOK SHOWS

$$\text{rent}_i = \beta_0 + \beta_1\cdot\frac{1}{\text{area}_i} + \varepsilon_i$$

Still an **ordinary linear model** — linear in $\beta$. Set $x_i := 1/\text{area}_i$ and run OLS unchanged.

> 🔑 **The transferable lesson:** *"linear model" constrains how parameters enter, not how covariates are shaped.* You can apply any function to a covariate before it becomes a column of $\boldsymbol{X}$.

---

## Example 3.4 — Munich Rent Index: Polynomial Regression

**Situation.** An alternative to transformation: use $\text{area}, \text{area}^2, \text{area}^3,\dots$

### 🤔 INTUITION

*Transformation vs polynomial — when would you pick each?*

- **Transformation** ($\log$, $1/x$, $\sqrt{x}$): you have a *reason* to expect that shape. One extra parameter. Interpretable.
- **Polynomial**: you *don't* know the shape and want flexibility. More parameters. Harder to interpret, and behaves badly at the edges of the data.

*And the danger?* Each extra power is another parameter — more variance, more overfitting risk. A degree-8 polynomial will wiggle through your data beautifully and predict catastrophically.

### ✍️ WHAT THE BOOK SHOWS

Both approaches work; the book plots **AIC as a function of polynomial degree** and it forms the classic **U-shape** — falling as real curvature is captured, then rising as the penalty for extra parameters overtakes the fit gain.

> 🔗 **This is Exam Summer 2025, Ex 2(d)** — the colleague who wants to add $(\text{age}-48)^3$, $\text{age}^4$, $\log(\text{age})$. The book's AIC-vs-degree plot is literally the answer: use an information criterion to find the bottom of the U.

---

## Example 3.6 — Munich Rent Index: Effect Coding

**Situation.** An alternative to dummy coding for categorical covariates.

### 🤔 INTUITION

*Dummy coding compares everything to one reference level. Is that the only option?*

No — and asking why reveals something. With dummy coding, the reference level is privileged and $\beta_0$ means "the reference group's mean." With **effect coding** (reference gets $-1$ instead of $0$), $\beta_0$ becomes the **grand mean** and each coefficient is a deviation from it.

**Nothing about the model changes** — same fitted values, same $R^2$, same predictions. Only the *labels on the coefficients* change.

### ✍️ WHAT TO KNOW

| Coding | $\beta_0$ means | $\beta_j$ means |
|---|---|---|
| **Dummy** (R default, exam default) | reference group mean | difference from **reference** |
| **Effect** | grand mean | difference from **grand mean** |
| **No intercept, all $c$ dummies** | — | that group's **mean** |

Use dummy coding unless told otherwise. Know the others exist.

---

## Example 3.7 — Munich Rent Index: Interaction with Quality of Kitchen ⭐

**Situation.** Does kitchen quality (none / standard / premium) affect rent, and does its effect depend on living area?

**This example contains a fully worked F-test.** Study it.

### 🤔 INTUITION — do this before reading

Kitchen quality has 3 levels ⟹ 2 dummies ⟹ **2 coefficients**, say $\beta_5$ (standard kitchen) and $\beta_6$ (premium kitchen).

Now, two *different* questions:

**Q1: "Does kitchen quality matter at all?"**
That's asking whether **both** coefficients are zero. Two restrictions. $H_0:\beta_5=\beta_6=0$, $r=2$.

**Q2: "Is a premium kitchen worth more than a standard one?"**
That's asking whether the two coefficients **differ from each other**. One restriction. $H_0:\beta_5-\beta_6=0$, $r=1$.

> 🔑 **Notice: same two coefficients, completely different tests.** The number of restrictions comes from the *question you're asking*, not from how many parameters appear. This is exactly the trap in Exam Summer 2025 Ex 1(i).

### ✍️ WHAT THE BOOK FINDS

**Q1 — does kitchen quality matter?**
$$F = 82.22, \qquad F_{2,4551}(0.95)=3.00$$
$$82.22 > 3.00 \Longrightarrow \textbf{reject } H_0$$

> Kitchen quality has a significant influence on average net rent.

**The book's own caution, worth quoting in an exam:**
> *"Note that this test does not necessarily imply that both regression coefficients are different from zero. The null hypothesis is rejected when **at least one** coefficient significantly differs from zero."*

**Q2 — is premium different from standard?**
$$F = 2.23,\qquad F_{1,4551}(0.95)=3.84,\qquad p = 0.13$$
$$2.23 < 3.84 \Longrightarrow \textbf{do not reject}$$

> The difference between a standard and a premium kitchen is **not significant**.

### 🔑 THE LESSON

Kitchen quality matters **as a whole** (Q1 rejects), but the **premium/standard distinction** doesn't (Q2 doesn't reject). The real effect is *having a proper kitchen at all*.

**That's a genuinely useful modelling conclusion reached by two F-tests** — and it's the model for how to answer "should I keep this variable / can I simplify this categorical variable?" questions.

---
---

# PART B — THE EXERCISE SHEETS (your actual course exercises)

**Sheets 3, 4 and 5 use one running model.** They are a single exam paper split into three weeks.

$$\text{wage}=\beta_0+\beta_1\text{age}+\beta_2\text{HSGrad}+\beta_3\text{SomeCollege}+\beta_4\text{CollegeGrad}+\beta_5\text{AdvDegree}+\beta_6\text{health.VeryGood}+\varepsilon$$

$$\hat{\boldsymbol\beta}'=(52.61,\ 0.62,\ 11.01,\ 23.16,\ 37.97,\ 62.63,\ 9.13),\qquad n=3000,\quad p=7$$

---

## 📄 SHEET 3 — Estimation

### Ex 2(a) — "Assuming normally distributed errors, what is the ML estimate for $\boldsymbol\beta$?"

#### 🤔 INTUITION

This looks like it wants a calculation. **It doesn't.** It's checking one piece of understanding.

Ask: *where does $\boldsymbol\beta$ appear in the Gaussian log-likelihood?*

$$\ell(\boldsymbol\beta,\sigma^2)=-\frac n2\log(2\pi\sigma^2)-\frac{1}{2\sigma^2}(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)$$

Only in the last term. With a **minus** sign. So maximising over $\boldsymbol\beta$ = **minimising** that quadratic form = least squares.

**No calculation required. The answer was already computed by OLS.**

#### ✍️ SOLUTION

$$\hat{\boldsymbol\beta}_{ML}=\hat{\boldsymbol\beta}_{LS}=(52.61,\ 0.62,\ 11.01,\ 23.16,\ 37.97,\ 62.63,\ 9.13)'$$

*Under normally distributed errors the ML and LS estimators of $\boldsymbol\beta$ coincide, because $\boldsymbol\beta$ enters the log-likelihood only through $-\frac{1}{2\sigma^2}(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)$.*

> 🔗 Exam Summer 2025 Ex 1(l) tests exactly this as TRUE/FALSE.

---

### Ex 2(b) — REML estimate of $\sigma^2$, given $\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}=3819720$

#### 🤔 INTUITION

*Why doesn't the question just say "divide by $n$"?*

Because there are **two** variance estimators and the word "restricted" picks one.

Think about what you did to get the residuals: you fitted 7 parameters. Those 7 parameters were chosen *to make the residuals small*. So the residuals are systematically too small — they've been optimised against. Dividing by $n$ would inherit that optimism.

**The correction is to divide by the number of residuals that were genuinely free to vary: $n-p$.**

> **The tell in the question:** "restricted maximum likelihood" ⟹ $n-p$. Plain "maximum likelihood" ⟹ $n$. And AIC/BIC always use $n$.

#### ✍️ SOLUTION

$$\hat\sigma^2_{\text{REML}}=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n-p}=\frac{3819720}{3000-7}=\frac{3819720}{2993}=\boxed{1276.218}$$

$$\hat\sigma=\sqrt{1276.218}=\boxed{35.724}$$

*(Compare $\hat\sigma^2_{ML}=3819720/3000=1273.240$ — smaller, as the downward bias predicts.)*

---

### Ex 2(c) — Predict John: 37, College Grad, very good health

#### 🤔 INTUITION

*The only real work is building $\boldsymbol{x}_0$ correctly.*

Walk the coefficient vector in order and ask "does John have this?"

| Term | John? | Contributes |
|---|---|---|
| intercept | always | $52.61$ |
| age | 37 | $0.62\times37$ |
| HS Grad | ❌ | $0$ |
| Some College | ❌ | $0$ |
| **College Grad** | ✅ | $37.97$ |
| Advanced Degree | ❌ | $0$ |
| **health ≥ Very Good** | ✅ | $9.13$ |

**The trap:** dummies are **mutually exclusive**. College Grad = 1 means every *other* education dummy = 0. Students sometimes add $\beta_2+\beta_3+\beta_4$ "because he passed through those levels." He didn't — dummy coding isn't cumulative.

#### ✍️ SOLUTION

$$\hat y_0 = 52.61+0.62(37)+37.97+9.13 = 52.61+22.94+37.97+9.13 = \boxed{\$122.65}$$

---

### Ex 2(d) — Residual and standardised residual ($h_{ii}=0.0016$)

#### 🤔 INTUITION

**Part 1, the raw residual:** just observed − fitted. Sign matters: positive = earns *more* than predicted.

**Part 2 — and here's the question worth thinking about:** *why can't we just say "$-1.22$ is small, so John isn't an outlier"?*

Because **small compared to what?** $-1.22$ dollars. Is that small? Depends on the typical size of a residual, which is $\hat\sigma\approx36$. So yes, tiny.

But there's a second, subtler issue. Residual variances are **not all equal**:
$$\text{Var}(\hat\varepsilon_i)=\sigma^2(1-h_{ii})$$
A person with unusual covariates (**high leverage**, $h_{ii}$ near 1) has the regression line dragged *toward* them, so their residual is artificially **small**. They could be a genuine outlier and the raw residual would hide it.

**Standardising divides that effect out**, putting everyone on a comparable scale where $|r_i|>2$ means the same thing for everybody.

*The hint "$h_{ii}=0.0016$" is the signal that they want the standardised version.*

#### ✍️ SOLUTION

$$\hat\varepsilon = y-\hat y = 121.43-122.65=\boxed{-\$1.22}$$

> *John earns \$1.22 per hour **less** than the model predicts for a man of his age, education and health. The model fits him very well.*

$$r=\frac{\hat\varepsilon}{\hat\sigma\sqrt{1-h_{ii}}}=\frac{-1.22}{35.724\times\sqrt{1-0.0016}}=\frac{-1.22}{35.724\times0.99920}=\frac{-1.22}{35.695}=\boxed{-0.034}$$

> *Far inside $\pm2$ — John is not an outlier. His leverage is also tiny ($0.0016$ vs an average of $p/n=7/3000=0.0023$), so he is an entirely typical observation.*

---

### Ex 2(e) — Standard errors from $(\boldsymbol{X}'\boldsymbol{X})^{-1}$ ["Hint: you do not need the whole matrix"]

#### 🤔 INTUITION

*Why is a 7×7 matrix given when we only want 7 numbers?*

Because $\text{Cov}(\hat{\boldsymbol\beta})=\sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}$ is a **whole covariance matrix**:
- **diagonal** = variances of each $\hat\beta_j$ ← this is what standard errors need
- **off-diagonal** = covariances *between* coefficients ← needed only for combinations like $\hat\beta_5-\hat\beta_2$

The hint is telling you: **you only need the diagonal.**

> ⚠️ **And now the trap that costs a mark every year.** The matrix is *printed* as rows 1…7. But it *indexes* $\beta_0,\beta_1,\dots,\beta_6$.
>
> **$\beta_1$ (age) is the SECOND diagonal element**, the $0.26\times10^{-5}$, not the first.
>
> Before touching a calculator: **write $\beta_0,\beta_1,\dots,\beta_6$ down the side of the matrix.** Ten seconds, one mark saved.

#### ✍️ SOLUTION

$$\widehat{\text{se}}(\hat\beta_j)=\hat\sigma\sqrt{\left[(\boldsymbol{X}'\boldsymbol{X})^{-1}\right]_{jj}},\qquad \hat\sigma=35.724$$

For $\beta_1$ (age), the diagonal entry is $0.26\times10^{-5}=2.6\times10^{-6}$:

$$\widehat{\text{se}}(\hat\beta_1)=35.724\times\sqrt{2.6\times10^{-6}}=35.724\times0.0016125=\boxed{0.0576}$$

*(Same recipe for the others: take the $j$-th diagonal entry, square-root it, multiply by $\hat\sigma$.)*

---

## 📄 SHEET 4 — Testing

### Ex 1 — Test of significance of regression

#### 🤔 INTUITION — the most important intuition on this sheet

$H_0:(\beta_1,\dots,\beta_6)'=\boldsymbol{0}$.

*What is this actually asking?*

It's asking: **"Is this model worth anything at all?"**

If all six slopes are zero, the model reduces to $\text{wage}_i=\beta_0+\varepsilon_i$ — predict $\bar y$ for everybody, ignore all covariates. So the test compares *your model* against *guessing the mean*.

**Why can't we just run six t-tests?** Two reasons, both examinable:
1. **Multiple testing.** Six tests at 5% each gives roughly a 26% chance of at least one false rejection.
2. They answer a different question. Six t-tests ask "is *this one* needed given the others?" The F-test asks "is *any* of them needed?" Under multicollinearity, every t-test can fail while the F-test resoundingly rejects.

**Why does $R^2$ appear in the formula?** Because $R^2$ *is* the comparison against the mean-only model — it's literally the fraction of variance the covariates explain beyond $\bar y$. So a test of "do the covariates help" naturally reads off $R^2$.

#### ✍️ SOLUTION

**(a)** $H_1:(\beta_1,\dots,\beta_6)'\neq\boldsymbol{0}$.
> $H_0$: none of age, education or health has any effect on expected wage.
> $H_1$: **at least one** of them does.

**(b)** Under $H_0$: $\;\text{wage}_i=\beta_0+\varepsilon_i$.

**(c)** $\boldsymbol{C}=(\boldsymbol{0}_{6\times1}\ \ \boldsymbol{I}_6)\in\mathbb{R}^{6\times7}$, $\;\boldsymbol{d}=\boldsymbol{0}_{6\times1}$, $\;r=6$.

**(d)** With $R^2_{H_0}=0$:
$$F=\frac{R^2/r}{(1-R^2)/(n-p)}=\frac{0.2685/6}{0.7315/2993}=\frac{0.044750}{0.00024440}=\boxed{183.10}$$

**(e)** $F\sim F_{6,\,2993}$ under $H_0$. Critical value $\approx F_{6,\infty}(0.95)=2.1016$.

$$183.10 \gg 2.1016 \Longrightarrow \textbf{reject } H_0 \text{ at } \alpha=0.05$$

> *Degrees of freedom $(6,\ 2993)$. The covariates **jointly** have highly significant explanatory power for wage. This does **not** imply every individual coefficient is non-zero.*

> ⚠️ $\alpha=0.05$ ⟹ the **0.95** column. F is one-sided. (For a two-sided t-test you'd want 0.975 — different rule, same $\alpha$.)

---

### Ex 2 — Composite test of a subvector

#### 🤔 INTUITION

$H_0:(\beta_1,\beta_6)'=\boldsymbol{0}$ — do **age and health together** matter, *given that education is already in the model*?

*Why is this a different, better question than testing them separately?*

Because it's a **conditional** question. It doesn't ask "does age matter?" — it asks "does age add anything **once education is accounted for**?" That's the question you actually care about when deciding whether to keep variables.

**And why do they give you two sums of squares this time, instead of $R^2$?**

Because the restricted model here is **not** the intercept-only model, so $R^2_{H_0}\neq0$ and you can't use the shortcut. You need the actual restricted fit.

> 🔑 **General signal: being handed a second, larger sum of squares means "F-test, SSE version." Being handed only $R^2$ with an all-slopes-zero hypothesis means "F-test, $R^2$ version."**

Note $3995721 > 3819720$ — as it must be. Imposing a restriction can never improve the fit.

#### ✍️ SOLUTION

**(c)** $\boldsymbol{C}=\begin{pmatrix}0&1&0&0&0&0&0\\0&0&0&0&0&0&1\end{pmatrix}$, $\;\boldsymbol{d}=\boldsymbol{0}$, $\;r=2$.

Model under $H_0$: $\;\text{wage}_i=\beta_0+\beta_2\text{HS}+\beta_3\text{SC}+\beta_4\text{CG}+\beta_5\text{AD}+\varepsilon_i$.

**(d)**
$$F=\frac{(\text{SSE}_{H_0}-\text{SSE})/r}{\text{SSE}/(n-p)}=\frac{(3995721-3819720)/2}{3819720/2993}=\frac{88000.5}{1276.218}=\boxed{68.95}$$

**(e)** $F\sim F_{2,\,2993}$; critical value $\approx F_{2,\infty}(0.95)=2.9987$.

$$68.95\gg2.9987\Longrightarrow\textbf{reject } H_0$$

> *Age and health **jointly** have significant explanatory power for wage even after controlling for education. Degrees of freedom $(2,\ 2993)$.*

---

### Ex 3 — t-test, CI, and prediction interval for $\beta_1$

#### 🤔 INTUITION — parts (a)–(d)

$H_0:\beta_1=0$, a **single** restriction on a **single** coefficient. That's a **t-test**.

*Could you use an F-test instead?* Yes — and you'd get $F=t^2$ and the identical decision. The t-test is just faster and gives you the **sign**, which F throws away.

For the CI in (d): *why does it have to agree with the test in (c)?*

Because they are the **same statement**. The CI is the set of values $c$ for which $H_0:\beta_1=c$ would not be rejected. So "CI excludes 0" and "reject $H_0:\beta_1=0$" are two phrasings of one fact. **If they disagree, you've made an arithmetic error** — that's a free self-check.

#### ✍️ SOLUTION (a)–(d)

**(a)** A **t-test** (test of significance).

**(b)** $\;t=\dfrac{\hat\beta_1}{\widehat{\text{se}}(\hat\beta_1)}=\dfrac{0.62}{0.0576}=\boxed{10.76}$

**(c)** $t\sim t_{n-p}=t_{2993}$; with df this large the table supplies the normal quantile $1.9608$.
$$|10.76|>1.9608\Longrightarrow\textbf{reject } H_0$$
> *Age has a significant effect on expected hourly wage at the 5% level.*

**(d)**
$$0.62\pm1.9608\times0.0576=0.62\pm0.1129\Longrightarrow\boxed{[0.507,\ 0.733]}$$
> *Excludes 0 ⟹ consistent with (c) ✓. With 95% confidence, each additional year of age is associated with between \$0.51 and \$0.73 more expected hourly wage, holding education and health fixed.*

---

### Ex 3(e) — Prediction interval ⭐ the conceptual heart of the sheet

#### 🤔 INTUITION — think hard here

The question: *predict the wage of a 50-year-old man with an advanced degree but less than good health.*

**Before any formula, ask: how well can I possibly do?**

Two completely different questions hide behind the word "predict":

**Q1: "What's the average wage of men like this?"**
With 3000 observations you can pin the average down tightly. More data ⟹ better. In the limit, you'd know it exactly.

**Q2: "What will THIS man earn?"**
Here's the thing: **even if you knew $\boldsymbol\beta$ perfectly**, you still wouldn't know his wage. He has his own $\varepsilon_0$ — his negotiating skill, his employer, his luck. **No amount of data tells you about a person you haven't met.**

So there's an **irreducible floor** on Q2 that doesn't exist for Q1.

$$\underbrace{\sigma^2}_{\text{his own randomness — never goes away}}+\underbrace{\sigma^2\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0}_{\text{our uncertainty about }\boldsymbol\beta\text{ — shrinks with }n}$$

**That first term is the "$1+$" in the formula.** It isn't a technicality — it's the mathematical statement that individuals are not averages.

**How to tell which the question wants:** *"the wage of a 50 year old man"* — singular, a person. **Prediction interval.**

**Watch the numbers when you're done:** $\sqrt{1+0.0035}\approx1.0017$ versus $\sqrt{0.0035}\approx0.059$. The "1" dominates completely — nearly all the uncertainty about one man is *his own randomness*, not our ignorance of $\boldsymbol\beta$. **With $n=3000$ we know the model well; we still barely know him.**

#### ✍️ SOLUTION

**Step 1 — build $\boldsymbol{x}_0$.** Age 50; Advanced Degree ⟹ that dummy is 1, all other education dummies 0; *less than good health* ⟹ health dummy 0.

$$\boldsymbol{x}_0=(1,\ 50,\ 0,\ 0,\ 0,\ 1,\ 0)'$$

**Step 2 — point prediction.**
$$\hat y_0=52.61+0.62(50)+62.63=52.61+31.00+62.63=\boxed{146.24}$$

**Step 3 — the standard error, with the "$1+$".**
$$\hat\sigma\sqrt{1+\boldsymbol{x}_0'(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{x}_0}=35.724\times\sqrt{1.0035}=35.724\times1.00175=35.786$$

**Step 4 — assemble.**
$$146.24\pm1.9608\times35.786=146.24\pm70.17$$
$$\boxed{[76.07,\ 216.41]}$$

**Compare — the CI for the mean at the same $\boldsymbol{x}_0$** (drop the 1):
$$146.24\pm1.9608\times35.724\times\sqrt{0.0035}=146.24\pm4.14=[142.10,\ 150.38]$$

> 🔑 **The prediction interval is ~17× wider.** We know the *group average* to ±\$4. We know *him* to ±\$70. That gap is $\sigma$ — real person-to-person variation, and it is the same whether you have 3,000 observations or 3,000,000.

---

## 📄 SHEET 5 — Model choice

### Ex 1(a) — Corrected coefficient of determination

#### 🤔 INTUITION

*Why does $R^2$ need correcting?*

Because $R^2$ **cannot decrease** when you add a covariate — not even a useless one. Adding a column expands the space $\hat{\boldsymbol{y}}$ can live in, so the projection can only get closer to $\boldsymbol{y}$.

**A criterion that always prefers the bigger model cannot choose a model.**

The correction multiplies the unexplained fraction $(1-R^2)$ by $\frac{n-1}{n-p}$, which **grows with $p$**. Add a parameter and you're penalised; the penalty is only worth paying if $R^2$ rises enough to offset it.

*Sanity check before computing:* $\bar R^2 < R^2$ always (when $p>1$).

#### ✍️ SOLUTION

$$\bar R^2=1-\frac{n-1}{n-p}(1-R^2)=1-\frac{2999}{2993}(1-0.2685)=1-1.002005\times0.7315$$
$$=1-0.73297=\boxed{0.2670}$$

Slightly below $R^2=0.2685$ ✓

> ⚠️ The book explicitly warns this penalty is **too weak** — $\bar R^2$ rises whenever a variable with $|t|>1$ is added, i.e. p-values around 0.3. Prefer AIC/BIC.

---

### Ex 1(b) — AIC and BIC

#### 🤔 INTUITION

*Adjusted $R^2$ already penalises complexity. Why do we need AIC and BIC too?*

Because **the size of the penalty is a choice, and it should follow from something.** Adjusted $R^2$'s penalty is arbitrary. AIC's comes from information theory (estimated expected out-of-sample prediction loss); BIC's comes from an approximation to the posterior probability of the model. **They're derived, not invented.**

*And why do AIC and BIC differ?* Because they answer different questions:
- **AIC** (penalty $2$ per parameter): "which model predicts best?"
- **BIC** (penalty $\log n$ per parameter): "which model is the true one?"

$\log(3000)=8.01 \gg 2$, so **BIC penalises four times harder** and picks smaller models. *B for Bigger penalty.*

**Three traps before you compute:**
1. 🔴 Use $\hat\sigma^2_{ML}=\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}/\mathbf{n}$ — **not** $n-p$. AIC comes from the *likelihood*, and the ML estimator is what maximises the likelihood.
2. 🔴 $\log$ = **natural** log. The sheet says so explicitly.
3. 🔴 The penalty is $(|M|+\mathbf{1})$ — **$\sigma^2$ counts as a parameter too**.

💡 **And a shortcut:** $n\log(\hat\sigma^2)$ is identical in both. Compute it once.

#### ✍️ SOLUTION

$$\hat\sigma^2_{ML}=\frac{3819720}{3000}=1273.240,\qquad \log(1273.240)=7.14929$$
$$n\log(\hat\sigma^2)=3000\times7.14929=21447.96$$

With $|M|=p=7$, so $|M|+1=8$:

$$\text{AIC}=21447.96+2(8)=\boxed{21463.96}$$
$$\text{BIC}=21447.96+\log(3000)\times8=21447.96+8.00637\times8=21447.96+64.05=\boxed{21512.01}$$

---

### Ex 2 — The full model, and choosing between them

#### 🤔 INTUITION

**First: where does $p$ come from?** The output says `Residual standard error: 34 on 2983 degrees of freedom`. With $n=3000$:
$$p=3000-2983=17$$

> 💡 **Learn this move.** The R output *always* tells you $p$ via the df line. Use it to cross-check your parameter count — if you counted the dummies wrong, this catches it immediately.

**Second, and this is the real thinking:** Model 2 has **10 more parameters**. Of course it fits better — $R^2$ rose from 0.269 to 0.340, and it had to rise.

**So the question is not "does it fit better?" It's "does it fit better than 10 extra parameters' worth?"**

That's precisely what the penalties price. And here's how to read the result:

> **If BIC — the harshest critic — still prefers the bigger model, the improvement is real.** BIC charges $8.01$ per parameter, so those 10 extra parameters cost 80 BIC points before they earn anything. If BIC still comes out ahead, they earned their keep several times over.

**Third — a condition people forget:** AIC and BIC are only comparable across models fitted to **the same data**, same $n$, same response scale. Both models here use the same 3000 observations and untransformed wage. ✓ (If one used $\log(\text{wage})$, the comparison would be meaningless without a correction.)

#### ✍️ SOLUTION

**(c)** $\hat\sigma^2_{ML}=\dfrac{3448498}{3000}=1149.499$, $\;\log(1149.499)=7.04709$, $\;n\log(\hat\sigma^2)=21141.25$

With $|M|=17$, $|M|+1=18$:
$$\text{AIC}=21141.25+2(18)=\boxed{21177.25}$$
$$\text{BIC}=21141.25+8.00637\times18=21141.25+144.11=\boxed{21285.36}$$

**(d)** Comparison:

| Criterion | Model 1 ($p=7$) | Model 2 ($p=17$) | Winner |
|---|---|---|---|
| $R^2$ | 0.2685 | 0.3396 | *(uninformative — must rise)* |
| $\bar R^2$ (larger better) | 0.2670 | **0.3361** | Model 2 |
| **AIC** (smaller better) | 21463.96 | **21177.25** | Model 2 |
| **BIC** (smaller better) | 21512.01 | **21285.36** | Model 2 |

> **I prefer Model 2.** All three penalised criteria agree. Adjusted $R^2$ rises substantially (0.267 → 0.336), and AIC and BIC fall by about 287 and 227 points respectively.
>
> The decisive evidence is **BIC**: it charges $\log(3000)=8.01$ per extra parameter, so Model 2's ten additional parameters cost it roughly 80 BIC points before contributing anything — and it still wins by 227. The improvement in fit is therefore far larger than could be explained by model size alone.
>
> The comparison is legitimate because both models are fitted to the same $n=3000$ observations with the same untransformed response.

---
---

# 🎯 THE INTUITIONS, COLLECTED

If you remember nothing else from this file:

| # | Intuition | Where it came from |
|---|---|---|
| 1 | Fanning scatter = variance grows with the covariate = **A3 broken** = unbiased but inefficient | Ex 3.1, Exam 2025 Ex 4(e) |
| 2 | Think about **what generates the data** and the functional form appears | Ex 3.3 |
| 3 | Flexibility costs variance; AIC-vs-complexity is **U-shaped** | Ex 3.4, Exam 2025 Ex 2(d) |
| 4 | **The number of restrictions comes from the question, not the parameters** | Ex 3.7, Exam 2025 Ex 1(i) |
| 5 | Rejecting a joint $H_0$ means **at least one** restriction fails | Ex 3.7 (book's own caution) |
| 6 | $\boldsymbol\beta$ enters the Gaussian likelihood only through the SSE ⟹ **ML = LS** | Sheet 3(a) |
| 7 | Residuals were optimised against ⟹ divide by $n-p$, not $n$ | Sheet 3(b) |
| 8 | Dummies are **mutually exclusive**, not cumulative | Sheet 3(c) |
| 9 | Standardise because high leverage **shrinks** residuals and hides outliers | Sheet 3(d) |
| 10 | Standard errors need the **diagonal**; label the rows $\beta_0…\beta_k$ first | Sheet 3(e) |
| 11 | The overall F-test asks **"is this model better than guessing the mean?"** | Sheet 4(1) |
| 12 | Six t-tests ≠ one F-test: multiple testing, and different questions | Sheet 4(1) |
| 13 | Restricting can never improve fit ⟹ $\text{SSE}_{H_0}\geq\text{SSE}$ ⟹ $F\geq0$ | Sheet 4(2) |
| 14 | CI and t-test are **the same statement** — use each to check the other | Sheet 4(3) |
| 15 | 🔑 **The "$1+$" is the fact that individuals are not averages.** It never shrinks | Sheet 4(3e) |
| 16 | $R^2$ can't fall ⟹ can't select ⟹ penalised criteria exist | Sheet 5(1a) |
| 17 | AIC = prediction, BIC = truth. **B for Bigger penalty** | Sheet 5(1b) |
| 18 | **If BIC still prefers the bigger model, the gain is real** | Sheet 5(2d) |
| 19 | The R output's df line **hands you $p$** — use it to check your dummy count | Sheet 5(2) |

---

## How to drill this file

**Pass 1 (now).** Read intuitions only. Skip every solution. You're building the *"what kind of question is this?"* reflex.

**Pass 2 (Day 9).** Cover the solutions. Work each one on paper. Check.

**Pass 3 (Day 14).** Solutions covered, **timed**. Sheet 4 in 25 minutes, Sheet 5 in 20.

**Pass 4 (Day 20).** Read only the 19-row table above. If any row makes you think *"wait, why?"*, that's the file section to reopen.
