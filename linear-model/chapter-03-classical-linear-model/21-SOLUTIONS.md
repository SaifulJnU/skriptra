# Ch 3 — SOLUTIONS

> Every numeric answer below was verified in Python before being written down. Work the exercise yourself first — reading a solution before attempting it teaches recognition, not ability.

---

## Part A — Estimation by hand

**Data:** $x=(1,2,3,4,5,6)$, $y=(3.1,4.0,5.2,6.8,8.1,9.0)$, $n=6$.

$$\bar x = 3.5,\qquad \bar y = 6.0333$$

**A1.**
$$S_{xy}=\sum(x_i-\bar x)(y_i-\bar y)=21.7,\qquad S_{xx}=\sum(x_i-\bar x)^2=17.5$$
$$\hat\beta_1=\frac{S_{xy}}{S_{xx}}=\frac{21.7}{17.5}=\boxed{1.240}$$
$$\hat\beta_0=\bar y-\hat\beta_1\bar x=6.0333-1.240(3.5)=\boxed{1.693}$$

**A2.** Fitted values $\hat y_i=1.693+1.240x_i$:

| $x$ | 1 | 2 | 3 | 4 | 5 | 6 |
|---|---|---|---|---|---|---|
| $\hat y$ | 2.933 | 4.173 | 5.413 | 6.653 | 7.893 | 9.133 |
| $\hat\varepsilon$ | 0.167 | −0.173 | −0.213 | 0.147 | 0.207 | −0.133 |

$$\sum\hat\varepsilon_i = 0.167-0.173-0.213+0.147+0.207-0.133 \approx 0.000\ \checkmark$$

*(This is guaranteed whenever the model has an intercept — it's algebra, not evidence of a good fit.)*

**A3.**
$$\text{SSE}=\sum\hat\varepsilon_i^2=\boxed{0.185}\qquad \text{SST}=\sum(y_i-\bar y)^2=\boxed{27.093}$$
$$R^2=1-\frac{\text{SSE}}{\text{SST}}=1-\frac{0.185}{27.093}=\boxed{0.9932}$$

*(A near-perfect line — sanity check: the data was constructed to lie almost exactly on $y\approx1.7+1.24x$, and $R^2\approx0.99$ confirms it visually before you even plot it.)*

**A4.** $p=2$ (intercept + slope), residual df $=n-p=4$.
$$\hat\sigma^2=\frac{\text{SSE}}{n-p}=\frac{0.185}{4}=\boxed{0.0463},\qquad \hat\sigma=0.2153$$
$$\widehat{\text{se}}(\hat\beta_1)=\frac{\hat\sigma}{\sqrt{S_{xx}}}=\frac{0.2153}{\sqrt{17.5}}=\frac{0.2153}{4.183}=\boxed{0.0515}$$

**A5.**
$$t=\frac{\hat\beta_1-0}{\widehat{\text{se}}(\hat\beta_1)}=\frac{1.240}{0.0515}=\boxed{24.10}$$
$$|24.10| > t_4(0.975)=2.776 \Longrightarrow \textbf{reject } H_0$$
> *There is strong evidence that $x$ has a non-zero effect on expected $y$.*

**A6.**
$$1.240\pm2.776(0.0515) = 1.240\pm0.143 \Longrightarrow \boxed{[1.097,\ 1.383]}$$
Excludes 0 ⟹ consistent with rejecting $H_0$ in A5 ✓ — the CI and the t-test are the same statement told two ways, so they must agree. If they hadn't, that would be a sign of an arithmetic slip, not a paradox.

---

## Part B — Joint hypothesis testing

**B1.** $H_0:\beta_{\text{team size}}=\beta_{\text{code age}}=0$ (both coefficients are zero — neither variable helps once test coverage and reviewers are already in the model). $H_1:$ at least one of the two is non-zero.

**B2.** Unrestricted model has an intercept + 4 covariates ⟹ $p=5$. Residual df $=n-p=40-5=35$.

**B3.**
$$F=\frac{(\text{SSE}_{H_0}-\text{SSE})/r}{\text{SSE}/(n-p)}=\frac{(1020-850)/2}{850/35}=\frac{85}{24.286}=\boxed{3.500}$$

**B4.** $F\sim F_{2,35}$ under $H_0$. Critical value $F_{2,35}(0.95)=3.267$.
$$3.500 > 3.267 \Longrightarrow \textbf{reject } H_0$$
> *At the 5% level, team size and code age jointly have significant explanatory power for defect count, even after controlling for test coverage and reviewer count. This does not mean both are individually significant — only that at least one is.*

**B5.** $F$ can never be negative. $\text{SSE}_{H_0}\geq\text{SSE}$ always, because the restricted model is a special case of the unrestricted one — imposing a restriction can only make the fit equal or worse, never better (you're forcing the optimiser to search a smaller space). So the numerator of $F$ is a ratio of non-negative quantities, and $F\geq0$ by construction. A negative $F$ means the two SSEs were swapped, or the restricted model was fit incorrectly (e.g. dropped the intercept too, or used different data).

---

## Part C — Model choice

**C1.**
$$R^2_A=1-\frac{980}{1500}=\boxed{0.3467}\qquad R^2_B=1-\frac{860}{1500}=\boxed{0.4267}$$
$$\bar R^2_A=1-\frac{199}{196}(1-0.3467)=\boxed{0.3367}\qquad \bar R^2_B=1-\frac{199}{193}(1-0.4267)=\boxed{0.4088}$$

*(As expected, $R^2$ rose going from A to B — it can't do otherwise when you add covariates — and $\bar R^2$ rose too, meaning the extra 3 parameters earned their keep on the adjusted-fit criterion.)*

**C2.**
$$\hat\sigma^2_{ML,A}=\frac{980}{200}=4.900,\quad \log(4.900)=1.5892 \quad\Rightarrow\quad n\log(\hat\sigma^2)=317.85$$
$$\text{AIC}_A=317.85+2(5)=\boxed{327.85}\qquad \text{BIC}_A=317.85+\log(200)(5)=317.85+5.298(5)=\boxed{344.34}$$

$$\hat\sigma^2_{ML,B}=\frac{860}{200}=4.300,\quad \log(4.300)=1.4586\quad\Rightarrow\quad n\log(\hat\sigma^2)=291.72$$
$$\text{AIC}_B=291.72+2(8)=\boxed{307.72}\qquad \text{BIC}_B=291.72+5.298(8)=\boxed{334.11}$$

**C3.** AIC prefers B ($307.72 < 327.85$). BIC prefers B ($334.11 < 344.34$). $\bar R^2$ prefers B too. **No disagreement** — all three penalised criteria agree Model B is worth its extra parameters.

**C4.** $\bar R^2$'s penalty (via $(n-1)/(n-p)$) is known to be too weak — the book notes it keeps rising for any added variable with $|t|>1$, i.e. fairly weak evidence. **BIC's penalty is derived from an approximation to the model's posterior probability and is the harshest of the three** ($\log(200)=5.30$ per parameter here, versus AIC's flat 2). If even BIC — charging roughly $5.3\times3=15.9$ points just for Model B's 3 extra parameters — still prefers B by a margin of $344.34-334.11=10.2$ points net, that's a stronger signal that the improvement in fit is real and not just an artefact of a lenient penalty.

---

## Part D — Diagnostics and short concepts

**D1.**
$$r_i=\frac{\hat\varepsilon_i}{\hat\sigma\sqrt{1-h_{ii}}}=\frac{-45}{60\sqrt{1-0.35}}=\frac{-45}{60\times0.8062}=\frac{-45}{48.37}=\boxed{-0.930}$$

**D2.**
$$D_i=\frac{r_i^2}{p}\cdot\frac{h_{ii}}{1-h_{ii}}=\frac{(-0.930)^2}{5}\cdot\frac{0.35}{0.65}=\frac{0.865}{5}\times0.5385=0.173\times0.5385=\boxed{0.093}$$

Not large — $r_i$ is well inside $\pm2$, and $D_i=0.093$ is nowhere near the common rule-of-thumb concern threshold (around 1, or $4/n$ for a formal cutoff). **This point is unremarkable**, despite the moderately high leverage — because a large leverage only becomes dangerous when it's paired with a large residual, and here the residual is small.

**D3.**
(a) **FALSE.** High leverage with a *small* residual is exactly the harmless case — the point sits far out in covariate space but close to the fitted line, which often *improves* the precision of the fit (it anchors the regression line). Leverage is only dangerous combined with a *large* residual, which is what Cook's distance captures jointly.

(b) **FALSE.** $\hat\beta_2$ is a **partial** effect — the effect of $x_2$ holding all other covariates fixed. The simple/marginal correlation ignores every other covariate. These can have opposite signs whenever $x_2$ is correlated with an omitted or held-fixed covariate that also affects $y$ (the classic case: a suppressor or confounding variable). Partial and marginal are different questions and can disagree.

**D4.** $\sum\hat\varepsilon_i=0$ is not evidence of anything — it is a **guaranteed algebraic consequence** of including an intercept in the model (it follows directly from the first normal equation $\boldsymbol{X}'\hat{\boldsymbol\varepsilon}=\boldsymbol{0}$). A badly misspecified model — wrong functional form, an omitted quadratic term, whatever — will *still* satisfy this exactly. What actually catches misspecification is looking at the **pattern** of residuals, not their average: a plot of residuals against fitted values (or against a covariate) that shows a curve, a funnel, or any systematic shape is the real diagnostic. The mean being zero is automatic; the *shape* is informative.
