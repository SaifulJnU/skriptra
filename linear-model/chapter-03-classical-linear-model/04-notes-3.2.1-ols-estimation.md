# 3.2.1 — Estimation of the Regression Coefficients (OLS)

> **The single most reliable source of marks in this exam.** It appears as a "derive/explain" question in at least two of your five past papers, and the marking key tells you exactly how it's scored:
>
> *WS 23/24, Ex 2(b) key:* **"1 point for correctly stating that RSS needs to be minimized. And 1 point for correctly derive the solution."**
>
> Two points, fully specified, always there. Learn to produce this on blank paper in four minutes.

---

## 1. The criterion

We choose $\boldsymbol\beta$ to make the fitted values as close as possible to the observed values, measuring closeness by the **sum of squared residuals**:

$$\boxed{\;S(\boldsymbol\beta)=\sum_{i=1}^n\left(y_i-\boldsymbol{x}_i'\boldsymbol\beta\right)^2=(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)\;}$$

$$\hat{\boldsymbol\beta}=\arg\min_{\boldsymbol\beta}S(\boldsymbol\beta)$$

**Say this out loud in the exam:** *"The method of ordinary least squares chooses $\hat{\boldsymbol\beta}$ to minimise the residual sum of squares $S(\boldsymbol\beta)=(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)$."* **That sentence alone is worth 1 point.**

### Why squares? (one line each, in case they ask)

- **Differentiable** ⟹ closed-form solution exists
- **Geometric** ⟹ squared Euclidean distance ⟹ orthogonal projection
- **Probabilistic** ⟹ under normal errors, equals maximum likelihood
- **Symmetric** ⟹ over- and under-prediction penalised equally; large errors penalised disproportionately

> 🔴 *Exam Summer 2025, Ex 1(b):* "Minimising the sum of **absolute** deviations gives the same coefficient estimates as minimising squared deviations." → **FALSE.** Least absolute deviations (LAD) targets the conditional **median** rather than the mean, has no closed-form solution, and generally produces different estimates. It is more robust to outliers.

---

## 2. ⭐ THE DERIVATION — memorise this page

### Step 1 — Expand the objective

$$S(\boldsymbol\beta)=(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)$$

$$=\boldsymbol{y}'\boldsymbol{y}-\boldsymbol{y}'\boldsymbol{X}\boldsymbol\beta-\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{y}+\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{X}\boldsymbol\beta$$

**Key observation:** $\boldsymbol{y}'\boldsymbol{X}\boldsymbol\beta$ is a $1\times1$ **scalar**, and a scalar equals its own transpose:

$$\boldsymbol{y}'\boldsymbol{X}\boldsymbol\beta=(\boldsymbol{y}'\boldsymbol{X}\boldsymbol\beta)'=\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{y}$$

So the two middle terms are identical and combine:

$$\boxed{\;S(\boldsymbol\beta)=\boldsymbol{y}'\boldsymbol{y}-2\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{y}+\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{X}\boldsymbol\beta\;}$$

*(Write the "scalar = its own transpose" remark explicitly — it shows you know why the 2 appears.)*

### Step 2 — Differentiate

Two matrix-calculus rules (Appendix A.8 of the book):

$$\frac{\partial(\boldsymbol{a}'\boldsymbol\beta)}{\partial\boldsymbol\beta}=\boldsymbol{a}\qquad\qquad \frac{\partial(\boldsymbol\beta'\boldsymbol{A}\boldsymbol\beta)}{\partial\boldsymbol\beta}=2\boldsymbol{A}\boldsymbol\beta\ \ (\boldsymbol{A}\text{ symmetric})$$

$\boldsymbol{X}'\boldsymbol{X}$ **is** symmetric, so:

$$\frac{\partial S(\boldsymbol\beta)}{\partial\boldsymbol\beta}=-2\boldsymbol{X}'\boldsymbol{y}+2\boldsymbol{X}'\boldsymbol{X}\boldsymbol\beta$$

### Step 3 — First-order condition

Set the gradient to zero:

$$-2\boldsymbol{X}'\boldsymbol{y}+2\boldsymbol{X}'\boldsymbol{X}\hat{\boldsymbol\beta}=\boldsymbol{0}$$

$$\boxed{\;\boldsymbol{X}'\boldsymbol{X}\hat{\boldsymbol\beta}=\boldsymbol{X}'\boldsymbol{y}\;}\qquad\textbf{the NORMAL EQUATIONS}$$

### Step 4 — Solve

If $\text{rank}(\boldsymbol{X})=p$ then $\boldsymbol{X}'\boldsymbol{X}$ is invertible, and:

$$\boxed{\;\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}\;}$$

### Step 5 — Confirm it's a minimum

$$\frac{\partial^2S(\boldsymbol\beta)}{\partial\boldsymbol\beta\,\partial\boldsymbol\beta'}=2\boldsymbol{X}'\boldsymbol{X}$$

which is **positive definite** when $\boldsymbol{X}$ has full column rank. Therefore $S$ is strictly convex and $\hat{\boldsymbol\beta}$ is the **unique global minimum**. ✓

*(Mention this. It's one line and it demonstrates you know a stationary point isn't automatically a minimum.)*

---

## 3. The four-line version (for when time is short)

$$S(\boldsymbol\beta)=(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)=\boldsymbol{y}'\boldsymbol{y}-2\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{y}+\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{X}\boldsymbol\beta$$
$$\frac{\partial S}{\partial\boldsymbol\beta}=-2\boldsymbol{X}'\boldsymbol{y}+2\boldsymbol{X}'\boldsymbol{X}\boldsymbol\beta\overset{!}{=}\boldsymbol{0}$$
$$\Longrightarrow\ \boldsymbol{X}'\boldsymbol{X}\hat{\boldsymbol\beta}=\boldsymbol{X}'\boldsymbol{y}\ \Longrightarrow\ \hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$$
$$\frac{\partial^2S}{\partial\boldsymbol\beta\partial\boldsymbol\beta'}=2\boldsymbol{X}'\boldsymbol{X}>0 \ \Longrightarrow\ \text{minimum, unique since } \text{rank}(\boldsymbol{X})=p$$

**Practise writing exactly this, timed, ten times.**

---

## 4. Model answer for the exam [2 pts]

> **Exam Summer 2025, Ex 4(b):** *"Explain the method of ordinary least squares and show the steps necessary to obtain $\hat\beta_0,\dots,\hat\beta_k$. It is not necessary to calculate them explicitly; the mathematical approach suffices."*

> **The idea.** The residual for observation $i$ is $\hat\varepsilon_i=y_i-\hat y_i=y_i-\boldsymbol{x}_i'\boldsymbol\beta$. Ordinary least squares chooses the parameter vector that makes these residuals collectively as small as possible, measured by the **sum of squared residuals**
> $$S(\boldsymbol\beta)=\sum_{i=1}^n(y_i-\boldsymbol{x}_i'\boldsymbol\beta)^2=(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta).$$
> Squaring ensures positive and negative deviations do not cancel, penalises large deviations more heavily, and makes the objective differentiable so a closed-form solution exists.
>
> **The steps.**
> 1. Expand: $S(\boldsymbol\beta)=\boldsymbol{y}'\boldsymbol{y}-2\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{y}+\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{X}\boldsymbol\beta$, using that $\boldsymbol{y}'\boldsymbol{X}\boldsymbol\beta$ is a scalar and hence equals its transpose $\boldsymbol\beta'\boldsymbol{X}'\boldsymbol{y}$.
> 2. Differentiate with respect to $\boldsymbol\beta$: $\;\partial S/\partial\boldsymbol\beta=-2\boldsymbol{X}'\boldsymbol{y}+2\boldsymbol{X}'\boldsymbol{X}\boldsymbol\beta$.
> 3. Set the derivative to zero, giving the **normal equations** $\boldsymbol{X}'\boldsymbol{X}\hat{\boldsymbol\beta}=\boldsymbol{X}'\boldsymbol{y}$.
> 4. Provided $\text{rank}(\boldsymbol{X})=p$, $\boldsymbol{X}'\boldsymbol{X}$ is invertible and $\;\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$.
> 5. The second derivative $2\boldsymbol{X}'\boldsymbol{X}$ is positive definite, so this stationary point is the unique global minimum.
>
> The individual estimates $\hat\beta_0,\dots,\hat\beta_k$ are the components of this vector.

---

## 5. The simple regression special case

Sometimes faster to work component-wise. Minimise $\sum(y_i-\beta_0-\beta_1x_i)^2$:

$$\frac{\partial S}{\partial\beta_0}=-2\sum(y_i-\beta_0-\beta_1x_i)=0 \;\Longrightarrow\; \sum\hat\varepsilon_i=0$$
$$\frac{\partial S}{\partial\beta_1}=-2\sum x_i(y_i-\beta_0-\beta_1x_i)=0 \;\Longrightarrow\; \sum x_i\hat\varepsilon_i=0$$

Solving:

$$\boxed{\;\hat\beta_1=\frac{\sum(x_i-\bar x)(y_i-\bar y)}{\sum(x_i-\bar x)^2}=\frac{\widehat{\text{Cov}}(x,y)}{\widehat{\text{Var}}(x)}=r_{xy}\frac{s_y}{s_x}\;}\qquad \boxed{\;\hat\beta_0=\bar y-\hat\beta_1\bar x\;}$$

> 💰 **$\hat\beta_0=\bar y-\hat\beta_1\bar x$ is a free mark.** *Example Exam LiMo 2020* gives mean goals $=48.61$, mean points $=46.61$, $\hat\beta_1=0.90509$, and asks for the missing intercept **A**:
> $$\hat\beta_0=46.61-0.90509\times48.61=46.61-43.997=\boxed{2.613}$$
> Fifteen seconds, no data needed.

---

## 6. Geometry — the picture that explains everything

$$\hat{\boldsymbol{y}}=\boldsymbol{X}\hat{\boldsymbol\beta}=\boldsymbol{H}\boldsymbol{y},\qquad \boldsymbol{H}=\boldsymbol{X}(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'$$

```
              y
              │╲
              │ ╲  ε̂  (residual — perpendicular)
              │  ╲
   ───────────┴───╲──────────────  column space of X
                   ŷ = Hy            (dimension p)
```

$\boldsymbol{y}$ lives in $\mathbb{R}^n$. The column space of $\boldsymbol{X}$ is a $p$-dimensional subspace — every value $\boldsymbol{X}\boldsymbol\beta$ that the model can possibly produce. $\hat{\boldsymbol{y}}$ is the **closest point in that subspace to $\boldsymbol{y}$**, and the shortest route is perpendicular.

**Three things fall out of this picture immediately:**

1. **The normal equations are an orthogonality statement.** $\boldsymbol{X}'\hat{\boldsymbol\varepsilon}=\boldsymbol{0}$ says the residual is perpendicular to every column of $\boldsymbol{X}$ — which is exactly what "shortest distance" means.
2. **$\text{SST}=\text{explained SS}+\text{SSE}$ is Pythagoras.** The right angle kills the cross-term.
3. **Residual df $=n-p$.** $\boldsymbol{y}$ has $n$ free dimensions; $\hat{\boldsymbol{y}}$ is pinned into $p$ of them; the residual can only move in the remaining $n-p$.

If you remember one picture from this course, remember this one. It makes six separate facts obvious.

---

## 7. Computing $\hat{\boldsymbol\beta}$ by hand (small examples)

The Example Exam and some sheets ask you to actually compute. Recipe for $n$ small, $p=2$:

**Step 1** — build $\boldsymbol{X}$ (don't forget the ones column) and $\boldsymbol{y}$.

**Step 2** — compute
$$\boldsymbol{X}'\boldsymbol{X}=\begin{pmatrix}n&\sum x_i\\\sum x_i&\sum x_i^2\end{pmatrix},\qquad \boldsymbol{X}'\boldsymbol{y}=\begin{pmatrix}\sum y_i\\\sum x_iy_i\end{pmatrix}$$

**Step 3** — invert the $2\times2$ using $\begin{pmatrix}a&b\\c&d\end{pmatrix}^{-1}=\frac{1}{ad-bc}\begin{pmatrix}d&-b\\-c&a\end{pmatrix}$:

$$(\boldsymbol{X}'\boldsymbol{X})^{-1}=\frac{1}{n\sum x_i^2-(\sum x_i)^2}\begin{pmatrix}\sum x_i^2&-\sum x_i\\-\sum x_i&n\end{pmatrix}$$

**Step 4** — multiply: $\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$.

**Step 5** — sanity check with $\hat\beta_0=\bar y-\hat\beta_1\bar x$. If it doesn't match, you have an arithmetic error. **Always do this check** — it costs ten seconds and catches most slips.

---

## 8. Properties of $\hat{\boldsymbol\beta}$ (full treatment in 3.2.3)

$$\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$$

**Linear in $\boldsymbol{y}$:** $\hat{\boldsymbol\beta}=\boldsymbol{A}\boldsymbol{y}$ with $\boldsymbol{A}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'$ a fixed matrix. *(This is the "L" in BLUE.)*

**Unbiased:** substitute $\boldsymbol{y}=\boldsymbol{X}\boldsymbol\beta+\boldsymbol\varepsilon$:

$$\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'(\boldsymbol{X}\boldsymbol\beta+\boldsymbol\varepsilon)=\underbrace{(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{X}}_{=\boldsymbol{I}}\boldsymbol\beta+(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol\varepsilon$$

$$\boxed{\;\hat{\boldsymbol\beta}=\boldsymbol\beta+(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol\varepsilon\;}$$

Taking expectations with $E(\boldsymbol\varepsilon)=\boldsymbol{0}$:

$$E(\hat{\boldsymbol\beta})=\boldsymbol\beta \qquad\checkmark$$

> **That boxed line is the workhorse of Section 3.2.3.** Both unbiasedness and the covariance matrix come straight from it. Memorise it.

**Covariance:** apply $\text{Cov}(\boldsymbol{Az})=\boldsymbol{A}\,\text{Cov}(\boldsymbol{z})\,\boldsymbol{A}'$ with $\boldsymbol{A}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'$ and $\text{Cov}(\boldsymbol\varepsilon)=\sigma^2\boldsymbol{I}$:

$$\text{Cov}(\hat{\boldsymbol\beta})=\boldsymbol{A}\sigma^2\boldsymbol{I}\boldsymbol{A}'=\sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{X}(\boldsymbol{X}'\boldsymbol{X})^{-1}$$

$$\boxed{\;\text{Cov}(\hat{\boldsymbol\beta})=\sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}\;}$$

**This is the formula behind every standard error, t-test and confidence interval in the course.** The standard error of $\hat\beta_j$ is

$$\widehat{\text{se}}(\hat\beta_j)=\hat\sigma\sqrt{\left[(\boldsymbol{X}'\boldsymbol{X})^{-1}\right]_{jj}}$$

— the square root of $\hat\sigma^2$ times the $j$-th **diagonal** element of $(\boldsymbol{X}'\boldsymbol{X})^{-1}$.

> 💰 **Sheet 3, Ex 2(e)** hands you the whole $(\boldsymbol{X}'\boldsymbol{X})^{-1}$ matrix and says *"Hint: you do not need the whole matrix."* Exactly — you need only the **diagonal**.

### Why the off-diagonals matter too

$\text{Cov}(\hat\beta_i,\hat\beta_j)=\sigma^2[(\boldsymbol{X}'\boldsymbol{X})^{-1}]_{ij}$, generally **non-zero**. Estimated coefficients are correlated with each other.

> 🔴 *Exam Summer 2025, Ex 1(f):* "In simple linear regression, $\hat\beta_1$ is **always** uncorrelated with $\hat\beta_0$." → **FALSE.**
>
> For simple regression, $\text{Cov}(\hat\beta_0,\hat\beta_1)=\dfrac{-\sigma^2\bar x}{\sum(x_i-\bar x)^2}$, which is zero **only if $\bar x=0$** — i.e. only if $x$ happens to be centred. The word "always" is what makes it false. *(Another argument for centring: it decorrelates the intercept from the slope.)*

---

## 9. Key takeaways

1. **OLS minimises $S(\boldsymbol\beta)=(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)'(\boldsymbol{y}-\boldsymbol{X}\boldsymbol\beta)$.** Saying this is worth 1 point.
2. **Derivation:** expand (scalar = its transpose ⟹ the 2) → differentiate → normal equations $\boldsymbol{X}'\boldsymbol{X}\hat{\boldsymbol\beta}=\boldsymbol{X}'\boldsymbol{y}$ → invert → check $2\boldsymbol{X}'\boldsymbol{X}$ is positive definite.
3. $\hat{\boldsymbol\beta}=(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol{y}$ — requires full rank.
4. **Simple regression:** $\hat\beta_1=\widehat{\text{Cov}}/\widehat{\text{Var}}$, $\hat\beta_0=\bar y-\hat\beta_1\bar x$ ← recovers missing intercepts.
5. **Geometry:** orthogonal projection onto the column space. Explains normal equations, Pythagoras, and $n-p$.
6. $\hat{\boldsymbol\beta}=\boldsymbol\beta+(\boldsymbol{X}'\boldsymbol{X})^{-1}\boldsymbol{X}'\boldsymbol\varepsilon$ — the workhorse identity.
7. $\text{Cov}(\hat{\boldsymbol\beta})=\sigma^2(\boldsymbol{X}'\boldsymbol{X})^{-1}$; $\widehat{\text{se}}(\hat\beta_j)=\hat\sigma\sqrt{[(\boldsymbol{X}'\boldsymbol{X})^{-1}]_{jj}}$ — **diagonal only**.
8. Coefficients are **correlated** unless the design says otherwise. "$\hat\beta_0$ and $\hat\beta_1$ are always uncorrelated" is FALSE.
