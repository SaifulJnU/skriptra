# Ch 2 — REAL-LIFE ANALOGIES

---

## 1. Dummy variables: the group photo and the "compared to whom?" problem

Five friends line up to be measured against a wall: Amina, Bilal, Chen, Dara, Erik.

You want to report their heights. You could give five absolute numbers. Or you could pick one person — say **Amina** — mark her height on the wall, and report everyone else as a **difference from Amina's mark**.

```
        Amina's mark ─────────────────────  0   ← the REFERENCE
        Bilal        ─────────────────────  +11
        Chen         ─────────────────────  +24
        Dara         ─────────────────────  +40
        Erik         ─────────────────────  +65
```

**Five people, four numbers.** You never need a number for Amina — she *is* the mark. That's why $c$ levels give $c-1$ dummies.

**Now the three consequences, all of them exam content:**

**(a) Why not five numbers?** If you also wrote "Amina: 0" you'd be saying the same thing twice — the mark and the zero carry identical information. In matrix terms, the column of ones (the intercept) is already Amina's mark, so a fifth dummy is redundant. The computer can't tell which of infinitely many splits between "the mark" and "Amina's offset" you meant. **No unique solution.** That's the dummy variable trap.

**(b) "Erik is 65 taller" — than *whom*?** Than **Amina**, always. Never than Bilal. To get Erik vs Bilal you **subtract**: $65-11 = 54$. This is the single most common dummy error.

**(c) Move the mark and every number changes — but no relationship does.** Mark Chen's height instead and all the numbers shift by 24. Erik is now $+41$, Amina is $-24$. **Every pairwise difference is identical.** The choice of reference category is arbitrary and affects no substantive conclusion — which is exactly why examiners ask you to *state* which reference you chose. It's a labelling convention, not a finding.

---

## 2. Interaction: the two escalators

Picture a shopping centre with two escalators side by side, both going up.

**No interaction (parallel lines):**
Escalator B starts one floor higher than A, but both move at the same speed. The gap between them is **always one floor**, forever. That's a dummy variable on its own: a constant shift.

```
    height
      │           ╱ B
      │         ╱ ╱ A     gap is constant
      │       ╱ ╱
      │     ╱ ╱
      └──────────── time
```

**With interaction (non-parallel lines):**
Now escalator B starts one floor *lower* but moves **twice as fast**. Early on, A is ahead. They cross. After that, B pulls away and keeps pulling.

```
    height
      │            ╱ B
      │          ╱ ╱ A     gap CHANGES
      │        ╱╱          they CROSS
      │      ╳╱
      └──────────── time
```

**Now read Sheet 2 directly off this picture.** Healthier men "start lower" ($-1.81$) but "ride faster" ($+0.43$ per year on top of $0.51$). The lines cross at age 4.2 — before anybody in the dataset exists — so in practice healthier men are always ahead, and *increasingly* so with age.

**And here's the trap the picture makes obvious:** asking "how far ahead is B?" has no single answer. It depends **when you ask**. That's why $\hat\beta_2 = -1.81$ isn't "the health effect" — it's the health effect *at the very bottom of the escalator* (age 0), which is a place nobody stands.

---

## 3. The polynomial: throwing a ball

$$\text{wage} = -10.43 + 5.29\,\text{age} - 0.05\,\text{age}^2$$

This is the same equation as a ball thrown in the air. It goes up, slows, stops, comes down.

| Ball | Wage |
|---|---|
| Initial upward speed | $\hat\beta_1 = 5.29$ |
| Gravity pulling down | $\hat\beta_2 = -0.05$ |
| Height at the top | maximum wage |
| Time at the top | age 52.9 |

**And the interpretation trap becomes obvious.** If someone asks "how fast is the ball moving?", you cannot answer without asking **when**. At launch: fast, upward. At the peak: zero. On the way down: fast, downward.

$$\text{speed} = \hat\beta_1 + 2\hat\beta_2\cdot\text{age} = 5.29 - 0.10\,\text{age}$$

At age 30: $+\$2.29$/year — still climbing.
At age 52.9: $\$0$ — the peak.
At age 65: $-\$1.21$/year — descending.

**So "$\hat\beta_1 = 5.29$ means wage rises \$5.29 per year" is exactly as wrong as "the ball moves at its launch speed the whole time."** It's the speed at the instant of launch — at age zero — and nowhere else.

**A negative $\hat\beta_2$ is gravity.** That's why you can predict its sign in the exam without seeing any data: careers, like balls, come back down.

---

## 4. Why linear regression fails for binary y: the bathtub with no walls

You want to predict **how full a bathtub is**, as a fraction from 0 (empty) to 1 (full).

A linear model is a **ramp**: for every minute of running the tap, add 5%. Straight line, forever.

Run it 30 minutes → 150% full. Run it backwards 5 minutes → −25% full.

**The ramp doesn't know the tub has a bottom and a rim.** That's objection (1), and it's fatal.

The **logit** model replaces the ramp with the way water *actually* behaves in a tub with walls:

```
  fullness
     1 ─────────────────────────  ← ceiling, approached, never crossed
                          ╭──────
                        ╭─╯
                     ╭──╯
                  ╭──╯            ← steepest in the middle
               ╭──╯
     0 ────────╯─────────────────  ← floor
              time / x'β
```

**And this picture explains all four objections at once:**

- **(1) Bounded:** the curve flattens as it approaches 0 and 1. Never escapes.
- **(4) Non-constant effect:** one extra minute matters enormously when the tub is half full and almost nothing when it's nearly overflowing. That's $\pi(1-\pi)$ — maximal at 0.5, vanishing at the ends.
- **(2) Heteroscedasticity:** uncertainty about "is it full?" is greatest in the middle and near zero at the extremes. When the tub is 99% full, you're confident. Same $\pi(1-\pi)$ shape. *The variance and the marginal effect have the same formula, because they come from the same geometry.*
- **(3) Non-normal errors:** at any moment the tub either overflowed or it didn't. Two outcomes, not a bell curve.

---

## 5. Odds: the gambler's translation

Probability and odds are the same information in two languages.

| Probability | Odds | Gambler says |
|---|---|---|
| 0.50 | 1 | "evens" |
| 0.75 | 3 | "3 to 1 on" |
| 0.90 | 9 | "9 to 1 on" |
| 0.99 | 99 | "99 to 1 on" |
| 0.25 | 1/3 | "3 to 1 against" |

**Why bother?** Because probabilities are trapped in $[0,1]$ and **odds are not** — odds run from 0 to $\infty$. Take the log and you get $-\infty$ to $+\infty$: the whole real line, which is exactly where a linear predictor lives.

$$\pi \in [0,1] \xrightarrow{\ \text{odds}\ } (0,\infty) \xrightarrow{\ \log\ } (-\infty,\infty)$$

**That's the entire construction of the logit link.** Two moves to free the response from its cage.

**And now the interpretation makes sense.** Look at the odds column: going from 0.90 to 0.99 is only 9 percentage points of probability, but the odds go from 9 to 99 — an elevenfold jump. Probability changes get *compressed* near the boundaries; odds changes don't. **That's why $\hat\beta_j$ is constant on the odds scale and not on the probability scale.** The model is linear where the geometry is uniform.

$$\exp(\hat\beta_j) = 2 \;\Longrightarrow\; \text{"this doubles your odds"}$$

Same statement whether you started at evens or at 99-to-1. That's the promise a linear model on the log-odds scale makes, and it's a reasonable one.

---

## 6. Partial vs marginal: the height and the shoes

You measure everyone entering a nightclub and find: **people wearing high heels are taller.**

Marginal association: strongly positive. Obviously.

Now sort everyone by their **barefoot height** into narrow bins, and look within each bin. Within a bin, does wearing heels still predict measured height? Yes, by exactly the heel height — that's the real, causal, partial effect.

But suppose instead you'd found something else: within each barefoot-height bin, the heel-wearers were *shorter* on measured height. That would be strange... unless shorter people wear taller heels, and you'd binned on the *wrong* variable.

**The point of the analogy:** what you condition on determines what you learn. $\hat\beta_j$ answers "*among people identical in every other included covariate*, how does $y$ differ?" Change the set of covariates and you change the question — which is why $\hat\beta_j$ can flip sign when you add a variable, and why "holding all other covariates fixed" isn't decoration. **It names the comparison being made.**

---

## 7. Centring: choosing where to put the origin

A map of your city can put $(0,0)$ anywhere. Put it at the North Pole and every building has coordinates like $(4,832,116\text{ m}, 998,412\text{ m})$ — technically correct, humanly useless, and the numbers are so large and so correlated that small errors get amplified.

Put $(0,0)$ at the **city centre** and everything reads $(2.3\text{ km}, -1.1\text{ km})$ — instantly meaningful.

**Nothing about the city changed. Distances between buildings are identical.** You moved the origin.

That's centring. $(\text{age}-48)$ instead of $\text{age}$ puts the origin at a **person who actually exists**, in the middle of the data:

- $\hat\beta_0$ becomes "expected wage at age 48" — interpretable, instead of "expected wage at age 0" — absurd.
- $(\text{age}-48)$ and $(\text{age}-48)^2$ become nearly **uncorrelated**, whereas raw age and age² over the range 18–80 are correlated above 0.98. Less collinearity ⟹ **smaller standard errors** ⟹ more precise estimates.

The fitted curve, the predictions, and the $R^2$ are all completely unchanged. You just chose better coordinates.

---

## 8. The link function: the adapter plug

You have a European appliance (the linear predictor: outputs any real number) and a British socket (the response: only accepts values in $[0,1]$).

You don't rebuild the appliance. You don't rewire the house. **You put an adapter between them.**

| Response lives in | Adapter (inverse link) |
|---|---|
| $(-\infty,\infty)$ | none needed — identity |
| $(0,1)$ | logistic / $\Phi$ — **logit / probit** |
| $(0,\infty)$ | $\exp$ — **Poisson** |

**This is the whole idea of generalised linear models**, and it's why the linear predictor $\boldsymbol{x}'\boldsymbol\beta$ never changes across the book. One appliance, different adapters. Everything you learn about building $\boldsymbol{X}$ — dummies, polynomials, interactions — transfers unchanged to every model in the rest of the textbook.

Which is a good reason to get Chapter 2.2 right.

---

## Analogy summary card

| Concept | Picture |
|---|---|
| Dummy variables | heights marked against one person on a wall |
| Why $c-1$ | you don't need a number for the person who *is* the mark |
| Reference category | move the mark: numbers change, differences don't |
| Comparing non-reference levels | subtract the two marks |
| No interaction | two escalators, same speed, constant gap |
| Interaction | different speeds — gap changes, lines cross |
| Why $\hat\beta_2$ isn't "the health effect" | "how far ahead is B?" depends **when** you ask |
| Polynomial | a thrown ball; $\hat\beta_2 < 0$ is gravity |
| Why not $\hat\beta_1$ alone | speed at launch ≠ speed throughout |
| Linear model on binary $y$ | a ramp in a bathtub with no walls |
| $\pi(1-\pi)$ | why the middle is where everything happens |
| Odds | the gambler's uncaged scale, $0$ to $\infty$ |
| Log-odds | uncaged twice — now it's the whole real line |
| Partial vs marginal | heels at the nightclub — what you binned on |
| Centring | moving the map origin to the city centre |
| Link function | an adapter plug between appliance and socket |
