# Ch 2 — PHILOSOPHY

---

## 1. Why regression models the *mean*, and what that quietly costs you

$E(y\mid\boldsymbol{x})$ looks like an obvious choice. It isn't — it's a decision with consequences.

Think about what you're doing when you summarise a whole conditional distribution by one number. For 40-year-old college graduates there is a *spread* of wages: some earn \$40/hr, some \$300/hr. The model reports one number and calls it the answer.

**Why the mean and not the median?** Because of the loss function. Squared-error loss is minimised by the conditional mean; absolute-error loss is minimised by the conditional median. **The estimator you use encodes what kind of mistake you consider bad.**

And these differ most exactly where it matters most. For right-skewed variables like income, the mean sits well above the median, dragged up by a long tail. "The expected wage of a 40-year-old graduate" may describe almost nobody — it's the average of a crowd, most of whom are below it.

This is why **quantile regression** exists (Chapter 10, out of scope, but worth knowing about). If you care about the *bottom* of the wage distribution — poverty policy, say — the conditional mean is the wrong summary and no amount of careful estimation will fix that. You've answered a well-posed question that isn't the one you had.

> **The lesson:** "$E(y\mid\boldsymbol{x})$" is not a neutral starting point. It's an answer to *"what feature of the distribution do I care about?"*, and the answer was chosen for you before you saw any data.

---

## 2. Dummy variables and the impossibility of a view from nowhere

Here is something quietly profound about how categorical variables work.

You **cannot** estimate "the effect of having a college degree" in absolute terms. It is not a thing the data contains. All you can ever estimate is "the effect of having a college degree **relative to some other category**."

And which category? **You choose.** Arbitrarily. The choice changes every number in your output and changes no relationship whatsoever.

This is not a limitation of regression. It is a fact about what "effect" *means*. There is no absolute wage-value of a college degree floating free in the world; there is only *the difference between having one and not having one*, and "not having one" has to be specified. Effects are inherently **comparative**.

The mathematics enforces this. Try to include all $c$ dummies plus the intercept and the design matrix becomes singular — the model literally has no unique answer. **The algebra refuses to let you ask a question with no comparison baseline.** The dummy variable trap isn't a technical nuisance to route around; it is the model telling you that your question is malformed.

> A great deal of confused public argument — about pay gaps, treatment effects, policy impacts — is people asserting comparative claims while suppressing the comparison. Regression can't do that. It has to name the reference category, or it breaks.

That's a design feature worth admiring.

---

## 3. Interactions: the end of the separable world

Without interactions, the model makes a strong and often false claim:

> *The effect of age is the same for everyone. The effect of health is the same at every age.*

That's a claim of **separability** — that the world decomposes into independent additive contributions you can consider one at a time.

It's a beautiful assumption. It's also frequently wrong. Medicine: a drug's effect depends on the patient's genotype. Education: a teaching method's effect depends on prior preparation. Economics: a stimulus's effect depends on the state of the credit market.

An interaction term is the admission that **the world doesn't factorise**.

And notice the price. Without interactions, "the effect of $x_1$" is one number. With them, it's a *function*: $\beta_1 + \beta_3x_2$. You've traded a simple sentence for a conditional one — *"it depends"* — and you now need $x_2$ to say anything at all.

That is why the interpretation traps in this chapter are so persistent. Students want $\hat\beta_1$ to be "the effect of $x_1$" because a single number is what we *want* an effect to be. The model is telling you that in an interactive world, **there is no such number**, and it's not being difficult — it's being accurate.

> **The deeper point:** every additive model is a bet that effects don't interact. Chapter 3's diagnostics can sometimes catch when that bet fails. But mostly you have to *think* about whether separability is plausible for your problem, before the data ever arrives. No test substitutes for that.

---

## 4. The logit model: what to do when your answer must live in a box

The linear model is generous. $\boldsymbol{x}'\boldsymbol\beta$ can be anything. Push a covariate far enough and the prediction goes anywhere on the real line.

For a probability, that generosity is a defect. Some quantities are **structurally constrained**: probabilities in $[0,1]$, counts in $\{0,1,2,\dots\}$, durations in $(0,\infty)$, proportions in $[0,1]$. A model that can violate its own quantity's definition isn't slightly imprecise — it's producing statements that don't parse.

There are two ways to respond, and the difference between them is a genuine methodological fork.

**Response A: constrain the parameters.** Restrict $\boldsymbol\beta$ so predictions stay in bounds. This fails immediately: for any non-trivial linear function on an unbounded covariate range, no $\boldsymbol\beta$ keeps you inside $[0,1]$ everywhere. And it makes the parameter space an awkward, data-dependent region.

**Response B: transform the scale.** Leave $\boldsymbol\beta$ unconstrained and change *where* the linearity lives. Find a function mapping $(0,1)$ onto all of $\mathbb{R}$, and be linear **there**.

Response B is what generalised linear models do, and its elegance is worth pausing on. **The constraint doesn't disappear — it gets absorbed into the geometry.** You never have to enforce $\pi\in(0,1)$, because the logistic function cannot produce anything else. Validity becomes automatic rather than checked.

> **The general principle, which outlives this course:** *when a quantity is constrained, don't fight the constraint — find the coordinate system in which it's invisible.* Model log-variance instead of variance. Model log-odds instead of probability. Model log-price instead of price. The constraint is real; the difficulty was an artefact of the coordinates.

---

## 5. The link function is a claim about how the world adds up

There's a subtlety in Section 2.3 that's easy to slide past.

Choosing the logit link doesn't merely *permit* valid probabilities. It **asserts something about how effects combine**.

Specifically: it says effects are **additive on the log-odds scale**, hence **multiplicative on the odds scale**, hence **neither** on the probability scale.

Is that true? Sometimes. It's an empirical claim, not a mathematical necessity. The probit link makes a *different* claim (additive on the scale of a latent normal variable). They fit almost identically in practice, which tells you something slightly deflating: **the data usually cannot distinguish between these claims.** We choose logit largely because $\exp(\hat\beta_j)$ is a sentence humans can say.

That's an honest and slightly uncomfortable thing to notice. Some modelling choices are driven by evidence. Others are driven by interpretability, convention, and what your software defaults to. **It is worth knowing which of your choices are which**, and being able to say so out loud is a mark of someone who understands the model rather than operates it.

---

## 6. Why "holding all other covariates fixed" is a metaphysical claim, not a caveat

The phrase sounds like boilerplate. It isn't. It describes a **hypothetical comparison that may not exist in your data at all.**

$\hat\beta_{\text{education}}$ compares two people identical in age, health, region, marital status, job class — differing *only* in education. Does such a pair exist? In a large sample with few covariates, plenty. In a small sample with many covariates, possibly **none**. The estimate is then an extrapolation into a region of covariate space you never observed.

Worse: sometimes the comparison is **incoherent**. What is "the effect of education, holding job class fixed," when education is *how you get into a job class*? You've conditioned on a **mediator** — a variable on the causal path — and the resulting coefficient answers a question nobody asked. This is a real and common error, and no diagnostic will flag it, because nothing is wrong with the arithmetic.

> **The uncomfortable truth:** which covariates you include is the single most consequential decision in a regression, it changes every coefficient, it is not determined by the data, and it cannot be validated by any test in Chapter 3.
>
> Chapter 3.4 will offer AIC, BIC and cross-validation, and they will help you choose among models **for prediction**. They will not tell you whether your coefficient means what you think it means. That requires knowing something about the world.

**This is the honest boundary of the subject**, and it's better to meet it here — in a chapter about how to build $\boldsymbol{X}$ — than to discover it later while over-trusting a $p$-value.

---

## 7. What Chapter 2 is really teaching

On the surface: some models.

Underneath: **the linear predictor $\boldsymbol{x}'\boldsymbol\beta$ is a universal chassis**, and everything else is bodywork.

- Curved relationship? Put $x^2$ in $\boldsymbol{X}$.
- Categorical covariate? Put dummies in $\boldsymbol{X}$.
- Effects that depend on each other? Put products in $\boldsymbol{X}$.
- Bounded response? Wrap $\eta$ in a link.

Four different-looking problems, one solution: **change what goes into $\boldsymbol{X}$, or change what comes out of $\eta$.** The estimation machinery underneath — which is all of Chapter 3 — never changes.

This is why the course is structured as it is, and why Chapter 3 can be so long without being repetitive. **You learn to build $\boldsymbol{X}$ once, in Chapter 2. You learn what to do with $\boldsymbol{X}$ once, in Chapter 3. Every model in the remaining 500 pages of the book is a combination of those two skills.**

If Chapter 2 feels easy, that's not a sign it's unimportant. It's a sign the chassis is well-designed.

---

## 8. A note on why the traps are traps

Look at what the recurring exam traps have in common:

| Trap | The wish behind it |
|---|---|
| "$\hat\beta_1$ is the effect of $x_1$" (with interaction) | I want effects to be single numbers |
| "$\hat\beta_1$ is the effect of $x$" (with $x^2$) | I want effects to be single numbers |
| "$\hat\beta_j$ increases $P(y=1)$ by $\hat\beta_j$" | I want the coefficient to be on the scale I care about |
| "positive coefficient ⟹ positive correlation" | I want partial and marginal effects to agree |
| "$m$ categories ⟹ $m$ dummies" | I want each category to have its own absolute number |

**Every one of them is the same wish: that the world be simpler than the model says it is.**

That's worth knowing about yourself, because it means these errors won't go away through repetition alone. They come back under time pressure, when you want the simple answer. The defence isn't memorising twelve rules — it's holding one habit:

> **When a variable appears in more than one place, or on a transformed scale, stop and differentiate.**

One habit, five traps closed. And it works because it replaces the wish with a calculation.
