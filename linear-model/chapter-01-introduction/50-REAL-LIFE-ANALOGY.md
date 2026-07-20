# Ch 1 — REAL-LIFE ANALOGIES

> When a formula feels arbitrary, it's because you've lost the picture. Here are the pictures.

---

## 1. The master analogy: a badly tuned radio

You're listening to a radio station. You hear two things at once:

- **the music** — structured, repeating, meaningful → **the systematic component $f(x)$**
- **the static** — hiss, crackle, no pattern → **the random error $\varepsilon$**

$$\text{what you hear} = \text{music} + \text{static}$$
$$y = f(x) + \varepsilon$$

Everything in this course is about **separating the music from the static.**

This analogy pays off repeatedly:

| Radio | Regression |
|---|---|
| You can't hear the pure music directly; you only hear music+static | You never observe $f(x)$; you only observe $y$ |
| A good tuner extracts more music | A good model has higher $R^2$ |
| Some static is irreducible | $\sigma^2 > 0$; even a perfect model has error |
| If you "tune" so aggressively you start hearing patterns in the static | **Overfitting** — Chapter 3.4 |
| Turning up the volume doesn't improve the music-to-static ratio | More data reduces variance of $\hat\beta$, not $\sigma^2$ |

**And the crucial one:** if you tune the radio to fit *this particular burst of static perfectly*, the setting will be terrible for the next song. That is exactly why Chapter 3.4 penalises complexity with AIC and BIC, and why $R^2$ alone is a bad model-selection criterion. Fitting the static is the entire problem.

---

## 2. Least squares: the tug-of-war rope

Imagine each data point is a child standing in a field, and you're holding a long straight rope. Each child grabs the rope with a rubber band tied from their head to the rope directly above them.

- A child far from the rope stretches their band a lot → **pulls hard**
- A child near the rope barely stretches → **pulls gently**

You let go and the rope settles. Where does it settle? At the position where **total stored energy in all the rubber bands is minimal.**

Rubber-band energy is proportional to (stretch)². So the rope settles exactly where

$$\sum_i(\text{vertical distance})^2 = \sum_i(y_i - \hat y_i)^2$$

is **minimised**. That is the least squares line.

**Two things this explains instantly:**

1. **Why squares and not absolute values?** Squaring makes distant points pull disproportionately hard. It also makes the problem smooth and differentiable, so there's a clean closed-form answer $\hat{\boldsymbol\beta} = (\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$. Absolute deviations (LAD) give a different line, with no closed form. *This is the answer to Exam Summer 2025, Exercise 1(b) — minimising absolute deviations does **not** give the same estimates. FALSE.*
2. **Why outliers are dangerous.** One child standing very far away stretches their band enormously and drags the whole rope toward themselves. That's **leverage and influence** — Cook's distance, Chapter 3.4.4.

---

## 3. Multiple regression: the confounded firefighters

The correlation between **number of firefighters** and **property damage** is strongly positive. Do firefighters cause damage?

No. **Fire size** causes both. Fire size is the **confounder**.

Now here's the analogy for what multiple regression *does about it*:

> Imagine sorting all the fires into bins by size — small fires here, medium there, huge over there. Now, *within each bin*, ask: do more firefighters mean more damage?
>
> Within a bin, fire size is roughly constant. And within a bin, more firefighters almost certainly means **less** damage.

Multiple regression does this sorting mathematically and continuously, rather than in bins. That is precisely what "**holding all other covariates fixed**" means, and it's why that phrase is worth a mark every time you write it.

**It also explains a trap:** the *marginal* correlation between firefighters and damage is positive, but the *partial* coefficient $\hat\beta_{\text{firefighters}}$ in a model that includes fire size is negative. **A positive coefficient does not imply a positive correlation** — Trap 6b, straight from the WS 23/24 paper.

---

## 4. The error term: the recipe you don't have

You're trying to predict how good a cake tastes from three things: grams of sugar, minutes in the oven, and oven temperature.

Your model:
$$\text{tastiness} = \beta_0 + \beta_1\text{sugar} + \beta_2\text{time} + \beta_3\text{temp} + \varepsilon$$

What's in $\varepsilon$?

- **Omitted variables:** flour quality, egg freshness, humidity in the kitchen, altitude, whether the baker was in a good mood
- **Measurement error:** your oven thermometer is off by 5°C; "tastiness" was rated by a person on a vague 1–10 scale
- **Genuine randomness:** bake the identical cake twice and it won't taste identical
- **Misspecification:** maybe the real effect of sugar is curved, not straight

$\varepsilon$ is **not "mistakes."** It's *everything real that your model doesn't contain.* It will never be zero, and a model with $\varepsilon = 0$ would be a model that has memorised your kitchen rather than learned about baking.

---

## 5. Why the type of $y$ picks the model: the container

You have something to store. What container do you need?

| What you're storing | Container |
|---|---|
| Water (flows, any amount, ±) | a **bucket** — continuous linear model |
| A light switch (only on or off) | a **switch box** — logit/probit |
| Apples (0, 1, 2, 3… never 2.5, never −1) | a **crate** — Poisson |
| T-shirt sizes (S < M < L < XL) | a **labelled rack** — ordinal |

You would never store water in a crate. And a linear model applied to a binary $y$ is exactly that mistake: the linear model happily predicts $\hat y = 1.34$ or $\hat y = -0.2$, and a probability cannot be either of those. That's the core of Exam Summer 2025 Exercise 4(a), and Chapter 2.3 gives the fix.

**Covariates, by contrast, are just ingredients you pour in.** Categorical ingredients get chopped into dummy variables first, but you don't change the container for them. **Only the thing being stored decides the container.**

---

## 6. $R^2$: the jigsaw puzzle

$y$ varies. Total variation = $\text{SST} = \sum(y_i - \bar y)^2$. Think of it as a 1000-piece puzzle.

- Your model correctly places some pieces → **explained SS**
- The rest are still scattered on the table → **SSE**

$$R^2 = \frac{\text{pieces you placed}}{\text{total pieces}}$$

$R^2 = 0.038$ means you placed 38 pieces out of 1000. Not great.

**Why this analogy earns you marks:** adding *any* new covariate can only place more pieces or leave the count unchanged — it can never *un*-place a piece. So **$R^2$ never decreases when you add a variable**, even a useless one (like dummies for the weekday someone was born on). That's why $R^2$ is worthless for comparing models of different sizes, and why $\bar R^2$, AIC and BIC exist. Chapter 3.4 in a sentence.

> *This makes Exam Summer 2025 Exercise 1(c) instant: "adding dummies for the weekday a person was born on can be expected to **lower** $R^2$" → **FALSE**. $R^2$ can never go down. (Adjusted $R^2$ can, and probably would.)*

---

## 7. Hats: the archer and the arrows

$\beta$ is the **bullseye** — a fixed point on the target. It doesn't move. You just can't see it.

$\hat\beta$ is **where your arrow landed.** Shoot again with a fresh sample and it lands somewhere else. $\hat\beta$ is random; $\beta$ is not.

- **Unbiased** ($E(\hat\beta) = \beta$): your arrows are centred on the bullseye — the *average* landing point is exactly right, even if no single arrow is.
- **Low variance:** your arrows are tightly grouped.
- **BLUE:** among all archers who shoot in a straight-line style (linear estimators) and are centred on the bullseye (unbiased), you have the tightest grouping.

This picture makes the **bias–variance tradeoff** (Chapter 3.4.1) obvious: an archer whose arrows are tightly grouped but 2 cm to the left may land closer to the bullseye *on average* than an unbiased archer who sprays wildly. Sometimes a little bias buys a lot of precision. That is the entire justification for ridge regression, lasso, and model selection.

---

## 8. Log transformation: switching from euros to percent

Two salary rises:

- Ali earns €1,000/month and gets a €500 rise
- Bilal earns €10,000/month and gets a €500 rise

In **euros**, identical. In **lived experience**, wildly different — Ali's life changed, Bilal barely noticed.

For skewed, positive quantities — wages, rents, prices, populations — **the natural unit of change is percent, not absolute.** Taking $\log(y)$ is exactly the change of ruler that converts "additive in euros" into "additive in percent."

$$\log(y) = \beta_0 + \beta_1 x \quad\Longleftrightarrow\quad y = e^{\beta_0}\cdot e^{\beta_1 x}$$

Additive on the log scale = **multiplicative** on the original scale. And because percentage changes are symmetric in a way euro changes aren't, the errors become better behaved: less skew, more constant variance.

That is *why* the log fix works for heteroscedasticity — you're not applying a trick, you're finally measuring in the right units.

---

## 9. The four scatter-plot questions: a doctor's examination

When a doctor examines you, she runs a fixed checklist regardless of what you came in for. Same with a scatter plot:

| Doctor asks | You ask |
|---|---|
| Is there a heartbeat? | **Is there a trend?** |
| Is the rhythm regular? | **Is the trend straight?** |
| Is blood pressure stable throughout? | **Is the spread constant?** |
| Any unusual lumps? | **Any outliers?** |

You run this checklist *before* the model (Chapter 1.2) on $y$ vs $x$, and *after* the model (Chapter 3.4.4) on $\hat\varepsilon$ vs $\hat y$. **Same four questions, twice.** Learning them once means Chapter 3.4.4 costs you almost nothing.

---

## The analogy summary card

| Concept | Picture |
|---|---|
| $y = f(x) + \varepsilon$ | music + static on a radio |
| Least squares | rubber bands pulling a rope; minimum total energy |
| Squaring, not absolute value | distant children pull disproportionately hard |
| Outliers / leverage | one child standing far away dragging the rope |
| Confounding | firefighters and fire size |
| "Holding others fixed" | comparing only within same-size fires |
| $\varepsilon$ | the parts of the recipe you don't have |
| Type of $y$ → model | you don't store water in a crate |
| $R^2$ | fraction of the jigsaw you've placed |
| $R^2$ never decreases | you can't un-place a puzzle piece |
| $\beta$ vs $\hat\beta$ | bullseye vs where the arrow landed |
| Bias–variance | tight-but-offset grouping can beat wide-but-centred |
| $\log(y)$ | switching the ruler from euros to percent |
| Residual plots | the doctor's fixed checklist, run twice |
