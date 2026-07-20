# Ch 1 — PHILOSOPHY: why the subject is built this way

> You don't need this for marks. You need it so the marks come cheaply, because a subject you find *reasonable* is far easier to remember than one you find arbitrary.

---

## 1. The founding admission: "I cannot predict you"

Every regression model opens with an act of intellectual humility.

$$y = f(x) + \varepsilon$$

That $+\varepsilon$ is a statistician saying, in writing, before doing any work: **"Whatever I build, I will be wrong about every individual case, and I am going to quantify how wrong rather than pretend otherwise."**

Compare this with the alternative traditions:

- **Physics** (classically) hoped for $y = f(x)$, full stop. Given the initial conditions, the planet goes *there*.
- **Machine learning** often chases $\hat y \approx y$ as hard as possible, treating $\varepsilon$ as a nuisance to be minimised.
- **Statistics** treats $\varepsilon$ as **an object of study in its own right.** Its variance $\sigma^2$ gets an estimator. Its distribution gets assumptions. Its estimated values get plotted and interrogated.

This is why the course spends so much time on things that look like bookkeeping — degrees of freedom, standard errors, distributions of test statistics. They're not bookkeeping. **They are the discipline of knowing how much you don't know.** In a world of confident predictions, that is the rarer and more valuable skill.

---

## 2. Why "regression" is a misleading name — and why that's instructive

Galton found that exceptionally tall parents had children who were tall but *less* exceptional. He called it "regression toward mediocrity" and thought he'd found a law of heredity.

He hadn't. He'd found a **statistical artefact that appears whenever two variables are imperfectly correlated.** The tallest parents in a sample are tall partly because of genuine height and partly because of *luck* — good measurement, good nutrition, being at the top of their own random variation. Luck doesn't transmit. So the children keep the genuine part and lose the lucky part.

The same phenomenon explains why the best-performing fund this year underperforms next year, why the sophomore album disappoints, and why "being shouted at after a bad landing improves the next landing" (Kahneman's flight instructors) is an illusion.

**The lesson, and it's the deepest one in Chapter 1:** a pattern in data can be entirely real, perfectly reproducible, and still not mean what it obviously seems to mean. The whole apparatus of this course — confounding, partial effects, hypothesis tests, model diagnostics — exists to protect you from your own pattern-recognition.

And the field is *named after* the first time someone got fooled. That's an appropriate monument.

---

## 3. Why squares? The most-questioned choice in statistics

Why minimise $\sum(y_i - \hat y_i)^2$ rather than $\sum|y_i - \hat y_i|$?

There are four answers, each true, and they converge suspiciously well:

**(a) Mathematical.** Squares are differentiable everywhere. Set the derivative to zero, get a linear system, solve it in closed form: $\hat{\boldsymbol\beta} = (\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$. Absolute values have a kink at zero; there's no closed form, only iterative algorithms. In 1805, when Legendre published least squares, this difference was decisive.

**(b) Geometric.** Squared error is squared Euclidean distance. Minimising it means **projecting $\boldsymbol{y}$ orthogonally onto the space spanned by the columns of $\boldsymbol{X}$.** The residual vector is perpendicular to every covariate — which is why residuals are uncorrelated with fitted values, and why the hat matrix $\boldsymbol{H} = \boldsymbol{X}(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'$ is a projection matrix. All of Chapter 3's structure is Pythagoras in $n$ dimensions.

**(c) Probabilistic.** If $\varepsilon \sim N(0,\sigma^2)$, then **maximising the likelihood is exactly equivalent to minimising the sum of squares.** The Gaussian density has $e^{-(\cdot)^2}$ in it; taking logs turns the exponent into a sum of squares with a minus sign. This is why *Exam Summer 2025, Exercise 1(l)* — "OLS is equivalent to the ML estimator under iid normal errors" — is **TRUE**, and why *Sheet 3, Exercise 2(a)* can ask for the ML estimate of $\beta$ and the answer is simply "the same as the least-squares estimate."

**(d) Decision-theoretic.** Squared error is the loss function under which the optimal prediction is the **conditional mean** $E(y|x)$. Absolute error targets the conditional **median** instead — which is what quantile regression does (Chapter 10). So "least squares" is really a *choice about what question you're asking*: mean or median.

**The philosophical point:** four independent lines of reasoning — algebra, geometry, probability, decision theory — arrive at the same estimator. That kind of convergence is what makes a method canonical rather than merely convenient.

**The philosophical caveat:** the convergence is *conditional* on the Gaussian assumption and the mean-squared-error loss. Change either and least squares stops being special. Squares aren't sacred; they're the answer to a particular, very reasonable question.

---

## 4. The trade you make when you choose a model class

Notice what happens when you say "the type of $y$ determines the model class."

You are **imposing structure you did not observe.** You never saw "the true model." You decided, on the basis of what kind of thing $y$ is, to restrict yourself to a family of possible answers.

Why restrict yourself at all? Why not let the data speak entirely freely?

Because **a model that can fit anything explains nothing.** With $n$ data points you can always draw a curve through all of them exactly, achieving $R^2 = 1$ and $\hat\varepsilon = 0$. That curve has learned the static, not the music. It will predict the next observation terribly.

This is the **bias–variance tradeoff**, and it's not a technicality — it's the central epistemological fact of the subject:

> **Every assumption you impose is a bet.** If the bet is right, you get precision you didn't earn from the data. If it's wrong, you get confident, systematic error.
>
> Assuming linearity is a bet. Assuming normal errors is a bet. Choosing which covariates to include is a bet.

Chapter 3.1.2 (discussion of assumptions) is you inspecting your bets. Chapter 3.4.4 (diagnostics) is you checking whether they paid off. Chapter 3.4.2 (AIC, BIC) is you pricing the bet: *how much complexity is this improvement in fit actually worth?*

**Nothing in this course is assumption-free, and the honest response to that is not to seek assumption-free methods — there are none — but to state your assumptions precisely enough that someone can check them.** That is what the whole formal apparatus is for.

---

## 5. Why "association, not causation" is more than a slogan

The regression coefficient $\hat\beta_j$ answers a specific question:

> *Among observations that happen to have the same values of the other covariates, how does $y$ differ between those with different $x_j$?*

That is a question about **comparison**. Causation asks a different question:

> *If I intervened and changed $x_j$ for a given unit, what would happen to its $y$?*

These coincide only when the assignment of $x_j$ is unrelated to everything else that affects $y$ — which is what randomisation buys you and what observational data almost never provides.

The firefighter example is funny because the confounder is obvious. The dangerous cases are the ones where it isn't: does education cause higher wages, or do the people who pursue education differ in ways (ability, family resources, patience) that would have raised their wages anyway?

**Multiple regression is a partial defence.** Include the confounder and you compare like with like. But you can only include confounders you (a) thought of and (b) measured. There is no statistical test that tells you what you forgot.

Hence the discipline: **write "associated with."** Not out of timidity, but because it is the *precisely correct* description of what you computed. Saying "causes" is claiming something your data cannot support. In this course it costs you a mark. In a policy report, it costs someone else something worse.

---

## 6. The strange status of the error term

Here's something worth sitting with.

$\varepsilon$ is defined as "everything not captured by $\boldsymbol{x}'\boldsymbol\beta$." But that means **$\varepsilon$ is not a property of the world — it's a property of your model.** Add a covariate and $\varepsilon$ changes. It is the shadow cast by your ignorance, and it moves when you move.

Yet we make strong assumptions about it: mean zero, constant variance, independent, normally distributed.

Is that legitimate? Assuming properties of your own ignorance?

Partly, yes, and for a good reason. If $\varepsilon$ contains *many small independent influences*, the Central Limit Theorem says their sum will be approximately normal, roughly symmetric, and centred. The normality assumption isn't arbitrary — it's a bet that **no single omitted variable dominates.**

Which gives you a real interpretation of assumption failure:

- **Non-normal, skewed residuals** ⟹ probably one big omitted effect, or the wrong scale for $y$
- **Heteroscedastic residuals** ⟹ the *amount* of what you're missing depends on where you are in covariate space
- **Curved residual pattern** ⟹ you got the shape of $f$ wrong, and the leftover structure has nowhere to go but into $\varepsilon$

**Diagnostics are not a compliance exercise. They are how the model tells you what you left out.** Chapter 3.4.4 is the model talking back, and Chapter 1.2 taught you the language.

---

## 7. Why this course is one idea, ten times

Here is the shape of everything ahead:

| Chapter | The question |
|---|---|
| 1 | What is the systematic part and what is noise? |
| 2 | What shapes can the systematic part take? |
| 3.1 | What exactly are we assuming? |
| 3.2 | Given the assumptions, what's the best estimate — and what does "best" mean? |
| 3.3 | Given the assumptions, how sure am I allowed to be? |
| 3.4 | Which set of assumptions should I have made in the first place? |

Read that column downward. It's a single argument, told in order, and every chapter is the previous chapter asked more carefully.

The reason people find this subject hard is that they try to learn six chapters. There is one chapter. It's called *"separate signal from noise, then be honest about the uncertainty,"* and it's told six times at increasing resolution.

---

## 8. A closing thought for someone who says they're bad at maths

The mathematics in this course is genuinely not difficult — it's arithmetic, a derivative, and matrix multiplication. What's difficult is something else, and it's worth naming so you don't mistake it for a maths problem:

**The difficulty is holding two things in mind at once: the true unknown world ($\beta$, $\varepsilon$, $\sigma^2$) and your sample-based shadow of it ($\hat\beta$, $\hat\varepsilon$, $\hat\sigma^2$) — and never letting them touch.**

That's not a mathematical skill. It's a discipline of attention. Every hat in every formula is a reminder of which world you're in. Students who "can't do the maths" are almost always students who have let the two worlds blur — who think a $p$-value is the probability the hypothesis is true, or that a confidence interval contains $\beta$ with probability 0.95 *after* it's been computed.

Keep the two worlds separate and the formulas will start looking like consequences rather than commandments.

That's the whole trick. The rest is arithmetic.
