# Ch 3 — PHILOSOPHY

---

## 1. What an assumption actually is: a receipt, not a fact

There's a temptation to read A1–A6 as a checklist of things that are simply *true or false about the world*, waiting to be verified like a fact-check. That's a subtly wrong picture.

An assumption in this chapter is closer to a **receipt for a specific promise.** "Under A1–A5, OLS is BLUE" is not a description of reality — it's a conditional statement: *if the world cooperates in these specific ways, here is exactly what you're entitled to conclude.* Nothing more, nothing less. The theorem is not making a claim about the world; it's making a claim about **what follows logically from a set of premises**, and it is *your* job, as the analyst, to argue that the premises are plausible for your particular data.

This reframing matters because it changes what a "violated assumption" question is really asking. It is never asking "is this bad?" in some vague moral sense. It's asking: **"which specific promise did you lose, and which theorem's premises no longer hold?"** That's why the correct answer to almost every assumption-violation question in this course has the same shape: *unbiased still, because [premise X still holds]; not BLUE anymore, because [premise Y failed]; standard errors invalid, because [the formula assumed premise Y].* You're not making a value judgement. You're tracing which receipts you can still cash in.

> **The habit this builds:** stop asking "is my model good?" Start asking "which specific guarantees does my model currently have, and which has it forfeited?" Those are different questions, and only the second one is answerable with precision.

---

## 2. The three-tier structure and the seduction of "good enough"

Notice something about the tier structure: A1, A2, A5 alone get you unbiasedness — arguably the *most important* property, since a biased estimator is systematically wrong in a predictable direction, forever, no matter how much data you collect. A3 and A4 only buy you *efficiency* — a smaller variance, but not a different expected value.

This creates a genuine and underappreciated asymmetry: **an estimator that's unbiased but inefficient is still, in a deep sense, honest.** Run the experiment a million times and it centres on the truth. An estimator that's efficient but biased is a different, more dangerous kind of wrong — it's *confidently, precisely, and reproducibly* wrong. Small variance around the wrong number is a worse epistemic position than large variance around the right one, because the small variance masks the error and inspires false confidence.

This is why heteroscedasticity (violating A3) is treated by the course as a genuine but *manageable* problem — you lose efficiency and your standard errors are wrong, both fixable with weighted least squares or robust standard errors — while misspecification (violating A1) is treated as much more serious. **Only A1 costs you the thing that can't be patched after the fact: the estimator's aim.** Everything else is a question of precision around a target you're still hitting on average.

> **The general lesson:** when you're choosing which imperfection to tolerate in a model, ask whether it costs you *accuracy* or merely *precision*. These are not equally forgivable, and this chapter's assumption structure is quietly teaching you to distinguish them.

---

## 3. Gauss–Markov: optimality is always optimality *within a rulebook*

"BLUE" — Best Linear Unbiased Estimator — sounds like an absolute superlative. It isn't. It's optimal *among linear, unbiased estimators.* Two adjectives, both silently restricting the competition.

This is one of the most important intellectual moves in all of statistics, and it's easy to walk past it in this course because the theorem is presented as a triumphant conclusion ("OLS is BLUE!") rather than as what it structurally is: **a statement that is true only because the comparison class was defined narrowly enough to make it true.**

Ridge regression, which you'll meet in Chapter 4, is not "worse" than OLS in some absolute sense — it simply **declined to enter the competition** by accepting bias in exchange for lower variance, and can beat OLS on total mean squared error precisely because MSE doesn't care about the linear/unbiased rulebook at all. Gauss–Markov never claimed otherwise. Its silence on biased estimators isn't an oversight; it's the boundary of what the theorem was ever trying to say.

> **This generalises far beyond regression.** Any time you hear "X is the best possible Y," the honest next question is: *best among what comparison class, judged by what criterion?* Optimality claims are never free-floating. They're always relative to a rulebook someone chose, and the choice of rulebook is itself a decision with consequences — usually made for reasons of tractability or interpretability, not because it's the only reasonable rulebook available.

---

## 4. Hypothesis testing: what a "restriction" really represents

Section 3.3 dresses up a very old and very human activity in matrix notation: **you have a complicated story ($k$ free parameters), and you want to know whether a simpler story ($k-r$ free parameters) is good enough.**

This is worth sitting with, because it reframes what "$r$" actually is. $r$ isn't a count of variables or coefficients — it's a measure of **how much simpler your competing story is**, in units of "independent claims removed." $H_0:\beta_1=-\beta_2+\beta_3$ removes exactly *one* degree of freedom from the model, no matter how many Greek letters are involved in expressing that removal, because it's asserting exactly one linear relationship among the parameters. The model still has all the same parameters in it — you've just tied one relationship between them down.

**And here is the deeper idea the F-test formalises:** simplicity should not be assumed for free. The onus of proof sits with the *simpler* model. You only reject the simpler story when the loss in explanatory power ($\text{SSE}_{H_0}-\text{SSE}$) is large relative to the noise inherent in the data ($\text{SSE}/(n-p)$). This is a quantitative version of Occam's razor with an explicit exchange rate: how much extra unexplained variation are you willing to tolerate, per unit of simplicity gained, before you insist on keeping the complexity? The F-statistic *is* that exchange rate, computed.

> **A researcher who never uses hypothesis tests, and one who rejects every simplification on sight, are making the same mistake in opposite directions: both have decided the exchange rate without ever calculating it.**

---

## 5. Confidence and prediction intervals: two entirely different kinds of not-knowing

It's tempting to think of "uncertainty" as one undifferentiated fog that shrinks as you gather more data. Sections 3.3.2 draws a sharp and philosophically important line through that fog.

**Estimation uncertainty** — how well do we know $\boldsymbol\beta$? — is *epistemic*. It exists because we're inferring a fixed, true quantity from a finite, noisy sample. Collect more data and this uncertainty genuinely, mathematically shrinks toward zero. In principle, with infinite data, you would know $\boldsymbol\beta$ exactly.

**Individual variation** — what will *this specific* new observation's $\varepsilon_0$ turn out to be? — is *aleatory* (from the Latin for "dice"). It isn't a gap in your knowledge that better data collection closes. It's genuine, irreducible randomness in the process that generates individual outcomes, and it exists independent of how much you know about $\boldsymbol\beta$. Even a researcher who has watched infinite data and knows $\boldsymbol\beta$ with perfect precision still cannot tell you exactly what one new person will earn, weigh, or score — only the *distribution* they're drawn from.

The prediction interval's "+1" is this distinction made visible in a formula. It is the chapter's way of insisting: **don't confuse "I don't know yet" with "it isn't determined yet."** The former is fixable by more data. The latter is not fixable by *anything* — it's a structural feature of a world where individuals differ from their group averages, and no amount of statistical sophistication makes that variation disappear. This is a genuinely humbling and important idea disguised as an algebra exercise: **however good your model, individuals will still surprise you, forever, by an amount that has a name ($\sigma$) and does not go to zero.**

---

## 6. Model choice: information criteria as a formalised bet about the future

AIC and BIC look like competing formulas, but they're really answers to two different metaphysical questions, and the difference is worth taking seriously rather than treating as a technicality.

**AIC is built to minimise expected prediction error on *new*, unseen data.** It's not asking "which model is true?" — a question it doesn't even really believe has a clean answer — but rather "which model, deployed forward in time on data it hasn't seen yet, will make the smallest average mistakes?" It's a forward-looking, pragmatic, engineering question.

**BIC is built (via a Bayesian derivation) to approximate which model has the highest posterior probability of *being* the data-generating process**, under the assumption that one of your candidate models actually is correct. It's asking a metaphysically bolder question: "which of these is the truth?" — and because it's chasing truth rather than mere predictive performance, it becomes pickier as evidence accumulates: with more data, weak effects that AIC will happily keep (because they help predict a little bit) BIC discards (because they don't survive the higher bar of "is this a real feature of the world, or noise that happened to help fit this particular sample?").

**Notice this is the same philosophical fork you find throughout the philosophy of science:** instrumentalism (models are tools; judge them by usefulness) versus realism (models are attempted descriptions of the world; judge them by truth). This course doesn't ask you to resolve that 300-year-old debate — but it quietly hands you two formulas, one for each stance, and expects you to know which one you're implicitly endorsing when you pick a model for a given purpose. Building a forecasting tool? You're an instrumentalist; lean AIC. Trying to establish which risk factors are real, for a scientific paper? You're chasing truth; lean BIC.

> **The deeper point:** "which model is best" is not a question with one correct answer waiting to be computed. It's a question that only becomes well-posed once you've said what you want the model *for* — and AIC versus BIC is that choice, formalised into two lines of algebra that look almost identical and mean something quite different.

---

## 7. Diagnostics: the chapter's honest admission that no test proves your model is correct

Here is the discomfort that Section 3.4.4 is built to manage, and it's worth stating plainly: **nothing in this entire chapter can prove your model is correctly specified.** Every t-test, F-test, AIC comparison, and confidence interval you've computed was computed *conditional on* the model being roughly right. If A1 fails — if the true relationship isn't linear, if you've omitted an interaction that actually matters — every downstream number can be precise, well-calculated, and completely misleading, all at once.

Diagnostics don't fix this. What they offer is narrower and more honest: **the ability to notice specific, recognisable failure signatures** — a curve in a residual plot suggesting missing nonlinearity, a funnel suggesting heteroscedasticity, a point with impossible leverage and a huge residual suggesting an outlier that's silently steering your entire fit. This is pattern recognition, not proof. A residual plot that looks clean is not a certificate of correctness — it only means the *specific* failure modes you know to look for aren't visibly present. There could be a mistake you don't have the vocabulary to see in a scatter plot.

> **This is the chapter's quiet final lesson, and it's a genuinely mature one:** statistical inference doesn't hand you certainty. It hands you a set of specific, falsifiable promises (the assumptions), tools to check whether the *visible* signatures of those promises' failure are present (diagnostics), and precise, honest arithmetic for what follows *if* those promises hold (everything else in the chapter). The humility is built into the structure. Respect it, and you'll write better exam answers — because "how would I know if this assumption failed?" is a more sophisticated question than "is this assumption true?", and it's the question this chapter is actually equipping you to answer.

---

## What Chapter 3 is really teaching

On the surface: how to estimate $\boldsymbol\beta$ and test hypotheses about it.

Underneath: **a discipline of tracking exactly which promises you're relying on, at every step, and refusing to claim more certainty than those promises actually license.** Every trap in this course's past papers — the "iff" that's too strong, the "always" that should be "usually," the CI-vs-test confusion, the AIC/BIC mix-up — is, underneath the arithmetic, a version of the same failure: **claiming a stronger guarantee than the premises actually earned.**

If you take one thing from this chapter into the rest of your statistical life, take this: **before stating any conclusion, ask what you assumed to get there, and say the conclusion no more confidently than the assumptions deserve.**
