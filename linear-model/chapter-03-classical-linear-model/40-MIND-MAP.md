# Ch 3 — MIND MAP

```
                    ┌─────────────────────────────────────────┐
                    │   y = Xβ + ε                            │
                    │   THE model this whole chapter builds    │
                    │   inference machinery around             │
                    └─────────────────────┬─────────────────────┘
                                          │
                ┌─────────────────────────┼─────────────────────────┐
                │                         │                         │
          3.1 DEFINE IT           3.2 ESTIMATE IT           3.3 TEST IT
       (what has to be true)     (get β̂, know its           (is β̂ or a
                                   properties)                 model good?)
                │                         │                         │
                │                         │                   3.4 CHOOSE
                │                         │                   BETWEEN MODELS
                │                         │                         │
```

---

## Branch A — 3.1 Model Definition & Assumptions

```
3.1 THE MODEL
│
├── 3.1.1 y = Xβ + ε
│   ├── X: n×p, one row per obs, one column per parameter
│   ├── p = k+1 (book convention — WATCH for p=k in other papers)
│   └── ŷ = Hy,  H = X(X'X)⁻¹X'  (the hat matrix — projects y onto col(X))
│
├── 3.1.2 ⭐⭐⭐ THE ASSUMPTIONS (examined every year)
│   │
│   ├── A1  linearity / correct specification  ─┐
│   ├── A2  E(ε)=0                              ├─► UNBIASED
│   ├── A5  rank(X)=p                           ┘   (+ existence/uniqueness)
│   │
│   ├── A3  Var(εᵢ)=σ²  (homoscedastic)         ─┐
│   ├── A4  Cov(εᵢ,εⱼ)=0 (no autocorrelation)    ┴─► + above ⟹ BLUE
│   │
│   └── A6  ε ~ Normal                          ──► + above ⟹ EXACT
│                                                     t/F/CI, OLS=ML
│   🔴 ONLY A1 biases β̂. A3/A4 cost EFFICIENCY, not bias.
│   🔴 Normality is NEVER required for Gauss–Markov — only for exactness.
│
└── 3.1.3 BUILDING X (imported from Ch 2, used constantly here)
    ├── dummies: c levels → c−1 columns
    ├── polynomials: effect = β₁+2β₂x, differentiate if x appears twice
    ├── interactions: non-parallel lines
    └── RESTRICTED MODEL under H₀: substitute → collect → this is 3.3's engine
```

---

## Branch B — 3.2 Parameter Estimation

```
3.2 ESTIMATION
│
├── 3.2.1 THE DERIVATION (practise from memory — 2 easy marks)
│   S(β) = (y−Xβ)'(y−Xβ)
│   ∂S/∂β = −2X'y + 2X'Xβ = 0
│   ⟹  X'Xβ̂ = X'y   (the normal equations)
│   ⟹  β̂ = (X'X)⁻¹X'y                    ← needs A5 (invertibility)
│
├── 3.2.2 ERROR VARIANCE — the ONE distinction that decides half the marks
│   │
│   ├── σ̂² = ε̂'ε̂/(n−p)     UNBIASED  ← t, F, CI, PI, se, REML, R's "resid SE"
│   └── σ̂²_ML = ε̂'ε̂/n        ML       ← AIC, BIC ONLY
│       🔴 mixing these up is the single most common numeric error in Ch 3
│
└── 3.2.3 PROPERTIES — Gauss–Markov [usually a 4-mark question]
    │
    ├── β̂ = β + (X'X)⁻¹X'ε                    (the workhorse identity)
    ├── E(β̂) = β                               unbiased (needs A1,A2,A5)
    ├── Cov(β̂) = σ²(X'X)⁻¹                     ⭐ THE formula
    │   └── se(β̂ⱼ) = σ̂·√[(X'X)⁻¹]ⱼⱼ            DIAGONAL only
    │       └── ⚠️ label rows β₀,β₁,... first — β₁ is the 2nd entry
    │
    ├── 🏆 GAUSS–MARKOV THEOREM
    │   "Under A1–A5, β̂ is BLUE: minimum variance among
    │    all estimators LINEAR in y and UNBIASED"
    │   🔴 normality NOT required — say so, it's the 4th mark
    │   🔴 drop "linear" or "unbiased" ⟹ claim is FALSE (ridge: biased, lower var)
    │
    └── UNDER A6 (normal errors):
        β̂ ~ N(β, σ²(X'X)⁻¹) exactly
        β̂_ML = β̂_LS  (but σ̂²_ML ≠ σ̂²_LS — different denominators!)
```

---

## Branch C — 3.3 Hypothesis Testing

```
3.3 TESTING
│
├── ⭐ THE ROUTING QUESTION: how many independent EQUATIONS in H₀?
│   │
│   ├── EXACTLY ONE ──► t-TEST
│   │   t = (β̂ⱼ−c)/se(β̂ⱼ)  ~  t_{n−p}
│   │   reject if |t| > t_{n−p}(1−α/2)   ← TWO-SIDED
│   │
│   └── TWO OR MORE ──► F-TEST
│       Cβ = d,  C is r×p (r = #restrictions, NOT #betas mentioned! 🔴)
│       F = [(SSEH₀−SSE)/r] / [SSE/(n−p)]  ~  F_{r,n−p}
│       reject if F > F_{r,n−p}(1−α)     ← ONE-SIDED (F≥0 always)
│
├── 🔴 THE #1 TRAP: β₁=−β₂+β₃ has THREE betas but r=1
│   (rewrite as β₁+β₂−β₃=0, count the "=" signs)
│
├── OVERALL F-TEST (special case, H₀: all slopes=0)
│   F = [R²/k] / [(1−R²)/(n−p)]  ~ F_{k,n−p}
│   asks: "is this model better than guessing ȳ?"
│
├── CONFIDENCE & PREDICTION INTERVALS
│   │
│   ├── CI for βⱼ:        β̂ⱼ ± t(1−α/2)·se(β̂ⱼ)
│   ├── CI for the MEAN at x₀:   x₀'β̂ ± t·σ̂·√[x₀'(X'X)⁻¹x₀]
│   └── PREDICTION for ONE NEW y₀: x₀'β̂ ± t·σ̂·√[𝟏+x₀'(X'X)⁻¹x₀]
│       🔑 the "+1" = ε₀'s own variance — NEVER shrinks, however big n gets
│       🔑 individuals ≠ averages — that's the whole idea in one interval
│
├── CI–TEST DUALITY
│   c inside CI  ⟺  fail to reject H₀:βⱼ=c
│   🔴 "CI excludes 0" ⟹ REJECT (not "fail to reject" — backwards trap)
│
└── LANGUAGE DISCIPLINE
    "fail to reject" not "accept" · "at least one" not "all"
    · rejecting joint H₀ ⇏ every individual β is nonzero
```

---

## Branch D — 3.4 Model Choice & Diagnostics

```
3.4 MODEL CHOICE
│
├── 3.4.1 BIAS–VARIANCE TRADEOFF
│   E[(y₀−f̂)²] = σ² (irreducible) + Bias² + Variance
│   more complexity ⟹ ↓bias, ↑variance ⟹ U-shaped total error
│   (bias enters SQUARED, variance LINEARLY — a little bias can pay off)
│
├── 3.4.2 THE THREE CRITERIA
│   │
│   ├── R̄² = 1 − [(n−1)/(n−p)](1−R²)
│   │   can decrease · CAN go negative · penalty criticised as too weak
│   │
│   ├── AIC = n·log(σ̂²_ML) + 2(|M|+1)        "predicts best"
│   │
│   └── BIC = n·log(σ̂²_ML) + log(n)(|M|+1)   "is the true model"
│       🔑 log(n) > 2 for n > 7.4 ⟹ BIC ALWAYS penalises harder
│       🔑 "B for Bigger penalty" ⟹ picks SMALLER models
│
│   🔴 always: σ̂²_ML = ε̂'ε̂/n (never n−p) · natural log · the "+1" for σ²
│   🔑 if BIC still likes the bigger model, the gain is REAL
│   ⚠️ only comparable: same data, same n, same response SCALE
│
├── 3.4.3 R² CAN'T FALL — the monotonicity fact tested 4× in past papers
│   adding ANY covariate: R² ↑ (or same), SSE ↓ (or same) — always
│   ⟹ THAT'S WHY you need penalised criteria to select a model at all
│
└── 3.4.4 DIAGNOSIS — the 4 questions from Ch 1's scatterplot, reborn
    │
    ├── residuals vs fitted    → non-linearity (curve), heteroscedasticity (funnel)
    ├── QQ plot                → non-normality; points on the 45° DIAGONAL
    │                             (🔴 NOT a horizontal line)
    ├── scale–location         → heteroscedasticity specifically
    └── residuals vs leverage  → influential points, via Cook's D
        │
        ├── hᵢᵢ (leverage)     unusual x — WHERE you stand
        ├── rᵢ (std. residual) unusual y — HOW FAR OFF you are
        └── Dᵢ (Cook's D)      both combined — HOW MUCH you'd be missed
            high leverage ALONE = harmless (only dangerous + large residual)
```

---

## The compressed picture — how the four sections chain together

```
   3.1 DEFINE                3.2 ESTIMATE              3.3 TEST                3.4 CHOOSE
   ─────────                 ────────────              ────────                ──────────
   y=Xβ+ε                    β̂=(X'X)⁻¹X'y              t-test (1 restriction)   R̄², AIC, BIC
   A1–A6 assumptions    ──►  σ̂² (÷n−p) vs             or F-test (r≥1)      ──► pick smallest
   determine WHAT              σ̂²_ML (÷n)         ──►  CI / PI                  penalised value
   you're allowed to claim    Gauss–Markov: BLUE         diagnose the fit        check residual plots
                              under A1–A5                                       to catch violated
                                                                                 assumptions ↺ (loops
                                                                                 back to 3.1)
```

**Chapter 3 in one sentence:** *the assumptions in 3.1 are a menu of promises — Section 3.2 tells you what each promise buys (unbiasedness, then BLUE, then exactness), Section 3.3 lets you act on those promises to test hypotheses, and Section 3.4 is what to do when you have several candidate promises (models) and need to pick one — closing the loop by diagnosing whether the promises actually held.*
