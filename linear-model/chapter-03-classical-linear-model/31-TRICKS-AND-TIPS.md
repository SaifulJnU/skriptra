# Ch 3 — TRICKS & TIPS

---

## 1. The parameter-counting reflex — do this before anything else

Before you touch a formula, write down $p$ = total number of $\beta$'s including the intercept. Every single thing in this chapter — degrees of freedom for t-tests, F-tests, AIC/BIC, standard errors — needs $p$ or $n-p$ as an input. Get $p$ wrong once at the top of the page and every downstream number is wrong, even if every subsequent step is executed perfectly.

**Fastest way to get $p$ right: use the R output.** If R reports "residual standard error on $df$ degrees of freedom" and you know $n$, then $p=n-df$. This is a free, built-in answer key — use it to check your own count of dummies and covariates.

---

## 2. The three-tier assumption ladder — memorise the shape, not the list

$$A1,A2,A5 \to \text{unbiased} \qquad +A3,A4 \to \text{BLUE} \qquad +A6 \to \text{exact tests, OLS=ML}$$

Don't memorise "A1 through A6" as an unordered list — memorise the **staircase**. Then any assumption question becomes: *which step does the violated assumption sit on, and what do you lose by falling off it?*

**Shortcut for the ubiquitous "what happens if A3/A4 is violated" question** — one sentence, reusable every time:

> *OLS remains unbiased and consistent (that only needs A1, A2, A5), but it's no longer BLUE — Gauss–Markov requires A3 and A4 too — and the usual $\sigma^2(\boldsymbol X'\boldsymbol X)^{-1}$ formula for $\text{Cov}(\hat{\boldsymbol\beta})$ is wrong, so every se, t-test, F-test and CI built on it is invalid.*

Have that sentence ready verbatim. It's worth 3–4 marks whenever it appears, and it appears almost every year.

---

## 3. Restriction-counting: count equals signs, not betas

The single highest-value habit in Section 3.3. Given any $H_0$, before doing anything else:

1. Move everything to one side so it reads "... $= 0$".
2. **Count the equals signs.** That's $r$.

$$H_0:\beta_1=-\beta_2+\beta_3 \;\Rightarrow\; \beta_1+\beta_2-\beta_3=0 \;\Rightarrow\; \textbf{one equation} \;\Rightarrow\; r=1$$

Three betas appear. $r$ is still 1. This single confusion is worth more lost marks across past papers than any other Chapter 3 trap — see Trap B1 in `32-TRAPS.md`.

---

## 4. Building $\boldsymbol{C}$ without fear

$\boldsymbol{C}$ is $r\times p$: **one row per restriction, one column per parameter — including $\beta_0$, even when its coefficient is 0.**

Fast recipe:
1. Write the parameter list across the top: $\beta_0,\beta_1,\dots,\beta_k$.
2. For each restriction (each row), write the coefficient multiplying each parameter once the equation is in "$=$ constant" form.
3. $\boldsymbol{d}$ is whatever's left on the other side.

$$H_0:\beta_1=\beta_2+1 \;\Rightarrow\; \beta_1-\beta_2=1 \;\Rightarrow\; \boldsymbol{C}=(0,1,-1,0,\dots),\quad d=1$$

---

## 5. The $F$–$t$ cross-check that catches half your arithmetic errors

If $r=1$ (testing a single coefficient or a single linear combination), you can test it with **either** a t-test or an F-test, and they must agree exactly:

$$F = t^2$$

**Use this as a free check.** Computed both by accident, or want to verify a t-test answer? Square it and see if it matches an F-computation of the same restriction. If it doesn't, you have an arithmetic error somewhere — found before you hand in the paper, not after.

---

## 6. Quantile lookup — write the rule down before you touch the table

$$\textbf{t-test, CI} \Rightarrow 1-\alpha/2 \qquad \textbf{F-test} \Rightarrow 1-\alpha$$

**Why this isn't arbitrary, so you never forget it under pressure:** a t-test / CI is two-sided — $\beta_j$ could be too high *or* too low — so the $\alpha$ risk splits into two tails, each getting $\alpha/2$. An F-test only ever rejects in the **upper** tail, because $\text{SSE}_{H_0}\geq\text{SSE}$ always (restricting can't improve fit), so $F\geq0$ and large values of $F$ are the only evidence against $H_0$. One tail, one full $\alpha$.

**Physically write "$1-\alpha/2=$" or "$1-\alpha=$" next to your test type before opening the table.** It takes two seconds and eliminates the single most common numerical slip in Section 3.3.

---

## 7. AIC/BIC in one breath, without re-deriving anything

$$\hat\sigma^2_{ML}=\frac{\text{SSE}}{n}\ (\textbf{always }n,\textbf{ never }n-p) \qquad \log = \ln \qquad \text{penalty} = (\text{const})\times(|M|+1)$$

$$\text{AIC}=n\log(\hat\sigma^2_{ML})+2(|M|+1)\qquad \text{BIC}=n\log(\hat\sigma^2_{ML})+\log(n)(|M|+1)$$

**Speed trick:** the term $n\log(\hat\sigma^2_{ML})$ is identical for both. Compute it **once**, then just add the two different penalties. Saves a full logarithm calculation.

**Sanity check before you commit to an answer:** is $\log(n) > 2$? For any $n>7.4$ (i.e. essentially every dataset you'll see), yes — so **BIC's penalty is always the bigger one**, and BIC always prefers models at least as small as AIC's choice. If your computed BIC is *smaller* than your AIC for the same model... that's fine, they're on different absolute scales (different additive constants), don't compare AIC of one model to BIC of another. **Only compare AIC-to-AIC and BIC-to-BIC across models.**

---

## 8. CI vs prediction interval: one word decides it

Read the question for **singular vs plural / group vs individual language**:

| Phrase in the question | Interval |
|---|---|
| "the **average** wage of…" / "the expected value for…" / "on average" | CI (no "+1") |
| "**a** 50-year-old man" / "**this** house" / "**one** new observation" | Prediction interval (**+1**) |

If genuinely ambiguous, compute both and comment that the prediction interval is the relevant one for a single new case — examiners very rarely leave this truly ambiguous, but stating the distinction explicitly earns the interpretation mark even if you picked the "wrong" one.

**Physical check on your final numbers:** the prediction interval must **always** be wider — if you compute a narrower one, you dropped the "+1" or the "1" under the square root.

---

## 9. Diagnostics — three quantities, three different jobs, don't blend them

| Symbol | Question it answers | Large value means |
|---|---|---|
| $h_{ii}$ (leverage) | Is this point's **$x$** unusual? | Point sits far from the centre of covariate space |
| $r_i$ (standardised residual) | Is this point's **$y$** surprising, given its $x$? | Poor fit for this point specifically |
| $D_i$ (Cook's distance) | Does removing this point **change the fitted line**? | Both of the above, multiplied together |

**One-line mnemonic:** *leverage is about where you stand, residual is about how far off you are, Cook's distance is about how much you'd be missed.* A point can have huge leverage and be totally harmless (small residual — it's exactly where the line predicts) or small leverage and still show up as an outlier (large residual near the centre of the data).

---

## 10. The exam-day formula card (write this from memory before you start)

$$\hat{\boldsymbol\beta}=(\boldsymbol X'\boldsymbol X)^{-1}\boldsymbol X'\boldsymbol y \qquad \text{Cov}(\hat{\boldsymbol\beta})=\sigma^2(\boldsymbol X'\boldsymbol X)^{-1} \qquad \hat\sigma^2=\frac{\hat{\boldsymbol\varepsilon}'\hat{\boldsymbol\varepsilon}}{n-p}$$
$$t=\frac{\hat\beta_j-c}{\widehat{\text{se}}(\hat\beta_j)}\sim t_{n-p} \qquad F=\frac{(\text{SSE}_{H_0}-\text{SSE})/r}{\text{SSE}/(n-p)}\sim F_{r,n-p}$$
$$\text{CI: }\hat\beta_j\pm t_{n-p}(1-\alpha/2)\widehat{\text{se}}(\hat\beta_j) \qquad \text{PI: }\boldsymbol x_0'\hat{\boldsymbol\beta}\pm t\,\hat\sigma\sqrt{1+\boldsymbol x_0'(\boldsymbol X'\boldsymbol X)^{-1}\boldsymbol x_0}$$
$$\text{AIC}=n\log(\hat\sigma^2_{ML})+2(|M|+1) \qquad \text{BIC}=n\log(\hat\sigma^2_{ML})+\log(n)(|M|+1)$$

If you can write these ten lines cold, in under three minutes, before the exam clock starts on your practice runs, you're not fighting notation during the real thing — you're just plugging in numbers.
