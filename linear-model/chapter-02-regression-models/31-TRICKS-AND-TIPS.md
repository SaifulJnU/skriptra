# Ch 2 — TRICKS & TIPS

---

## 1. The dummy-counting reflex

Read the question, and for each categorical covariate immediately write in the margin:

```
education: 5 levels → 4 dummies
birthplace: 2 levels → 1 dummy
health:    2 levels → 1 dummy
```

Then total the parameters:

$$p = \underbrace{1}_{\text{intercept}} + \underbrace{(\#\text{continuous})}_{} + \underbrace{\textstyle\sum(c_m-1)}_{\text{categoricals}} + \underbrace{(\#\text{interactions})}_{}$$

**This one habit gives you three things at once:** the model equation, the number of parameters, and the residual df $n-p$ for every test later in the paper. Do it before you write anything else.

---

## 2. Spot the reference category in two seconds

> **The level missing from the R output is the reference.**

R output shows `education2. HS Grad`, `education3. Some College`, `education4. College Grad`, `education5. Advanced Degree`. Where's `education1`? Missing ⟹ **`< HS Grad` is the reference.**

R's default is the **first level** of the factor (alphabetical, or the order you specified). The `1.` prefix in that dataset makes it obvious.

---

## 3. Comparing two non-reference categories: subtract

To compare category A to category B (neither is the reference), and everything else is held equal:

$$\Delta = \hat\beta_A - \hat\beta_B$$

**Never** use $\hat\beta_A$ alone — that compares A to the *reference*.

| Comparison | Calculation |
|---|---|
| Advanced Degree vs `< HS Grad` (reference) | $62.63 - 0 = 62.63$ |
| Advanced Degree vs HS Grad | $62.63 - 11.01 = 51.62$ |
| Some College vs College Grad | $23.16-37.97 = -14.81$ |

---

## 4. The "cancel what's shared" shortcut

When comparing two hypothetical individuals, **anything they have in common cancels**. Don't compute two full predictions — compute the difference directly.

> *Sheet 1(c): 50-year-old at level 3 vs 50-year-old at level 5.*
> Both 50 ⟹ the intercept and both age terms cancel.
> $\Delta = \hat\beta_5-\hat\beta_3 = 64.99-24.17 = 40.82$. **One subtraction.**

> *Sheet 1(d): 40-year-old at level 4 vs 20-year-old at level 1.*
> Only the intercept cancels.
> $\Delta = \hat\beta_1(40-20)+(\hat\beta_4 - 0) = 0.56869(20)+39.767 = 51.14$.

Saves 90 seconds and removes two chances to make an arithmetic slip. In a 60-minute exam, that's real.

---

## 5. Interaction models: split into two lines immediately

The moment you see an interaction with a dummy, write the two-row table. Don't try to interpret the four coefficients in the abstract.

$$y = \beta_0+\beta_1x+\beta_2D+\beta_3(xD)$$

| | Intercept | Slope |
|---|---|---|
| $D=0$ | $\beta_0$ | $\beta_1$ |
| $D=1$ | $\beta_0+\beta_2$ | $\beta_1+\beta_3$ |

**Everything you'll be asked reads straight off this table:**
- "Give the slope for smokers" → row 2, column 2
- "What does the interaction do geometrically?" → the slopes differ ⟹ **non-parallel lines**
- "Where do the lines cross?" → set the two equations equal and solve

**Mnemonic:** *the dummy moves the line, the interaction tilts it.*

---

## 6. When a variable appears twice, differentiate

If a covariate shows up in more than one term — as $x$ and $x^2$, or in an interaction — **never** quote a single coefficient as "the effect." Take the derivative:

| Model | Effect of $x$ |
|---|---|
| $\beta_0+\beta_1x$ | $\beta_1$ (constant) |
| $\beta_0+\beta_1x+\beta_2x^2$ | $\beta_1+2\beta_2x$ |
| $\beta_0+\beta_1x+\beta_2D+\beta_3xD$ | $\beta_1+\beta_3D$ |
| $\beta_0+\beta_1x+\beta_2z+\beta_3xz$ | $\beta_1+\beta_3z$ |

**Quadratic turning point:** $\;x^* = -\dfrac{\hat\beta_1}{2\hat\beta_2}$

- $\hat\beta_2 < 0$ ⟹ downward parabola ⟹ $x^*$ is a **maximum** (the usual case for age/wage)
- $\hat\beta_2 > 0$ ⟹ upward parabola ⟹ $x^*$ is a **minimum**

*(Sheet 2: $-5.29/(2\times-0.05) = 52.9$ years.)*

---

## 7. Logit interpretation in one line

$$\boxed{\;\text{"multiplies the odds by } \exp(\hat\beta_j)\text{"}\;}$$

Memorise that phrase. It is exact, it is correct, and it is what markers want.

**Mental exponentials** (no calculator needed for small $\beta$):

For small $|\beta|$, $\exp(\beta)\approx 1+\beta$:

| $\hat\beta$ | $\exp(\hat\beta)$ | Quick reading |
|---|---|---|
| $0.01$ | $1.010$ | +1% odds |
| $0.05$ | $1.051$ | +5% odds |
| $0.10$ | $1.105$ | +10% odds |
| $0.028$ | $1.028$ | **+2.8% odds** ✓ |
| $0.69$ | $2.00$ | odds **double** |
| $1.00$ | $2.72$ | — |
| $-0.69$ | $0.50$ | odds **halve** |
| $-0.85$ | $0.427$ | odds down ~57% |

**Two anchors worth memorising:** $\log 2 = 0.693$ and $e = 2.718$. From those you can approximate almost anything.

**Marginal effect on probability:** $\hat\beta_j\pi(1-\pi)$, maximised at $\pi=0.5$ where $\pi(1-\pi)=0.25$. So the **biggest possible** probability effect is $0.25\hat\beta_j$. Useful sanity check.

---

## 8. The four-part interpretation checklist

Before moving on from any interpretation question, verify:

- [ ] **Unit of $x$** stated ("one **year**", "one **month**")
- [ ] **Unit of $y$** stated ("**dollars per hour**", "**percent**", "**the odds**")
- [ ] **"Expected"** or **"on average"** present
- [ ] **"Holding all other covariates fixed"** present (multiple regression)
- [ ] **"Associated with"**, never "causes"
- [ ] For dummies: **named the reference category** you're comparing to

Six ticks, full marks, twenty seconds.

---

## 9. Sanity checks that catch errors fast

| Check | What it catches |
|---|---|
| Is the predicted $y$ plausible? (a \$400/hr wage isn't) | arithmetic slip |
| Do the dummy coefficients increase monotonically with education? | mis-assigned coefficients |
| Is the quadratic turning point inside the data range? | sign error in $\hat\beta_2$ |
| Do I have $c-1$ dummies, not $c$? | the classic |
| Does my parameter count match the R output's df? | missing/extra term |
| Is a logit "probability" between 0 and 1? | used the wrong scale |

**The R-output df check is the strongest one.** If the output says "Residual standard error: 34 on 2983 degrees of freedom" and $n=3000$, then $p = 3000-2983 = 17$ parameters — so 16 covariate terms plus an intercept. Count your terms; if they don't total 16, you've miscounted a categorical variable's levels.

---

## 10. Writing model equations fast (exam template)

Have this skeleton memorised so you can produce it under pressure:

> **Define** $D^{A}_i = \begin{cases}1 & \text{if } i \text{ is } A\\ 0 & \text{otherwise}\end{cases}$, $\;D^B_i = \dots$
>
> **Model:** $y_i = \beta_0+\beta_1x_i+\beta_2D^A_i+\beta_3D^B_i+\varepsilon_i,\quad i=1,\dots,n$
>
> **with** $\varepsilon_i$ iid, $E(\varepsilon_i)=0$, $\text{Var}(\varepsilon_i)=\sigma^2$.
>
> **Reference category:** [the omitted level].

Four lines. Every mark in an Exercise-2-style question is somewhere in those four lines. Practise writing them until it's automatic — you should be able to produce the whole thing in under two minutes.
