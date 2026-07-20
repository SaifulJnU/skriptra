# Ch 2 — MIND MAP

```
                    ┌────────────────────────────────────┐
                    │   E(y | x) = f(x)                  │
                    │   regression models the MEAN       │
                    │   ⟹ always say "expected"          │
                    └─────────────────┬──────────────────┘
                                      │
                    LINEAR PREDICTOR  η = x'β
                    (the skeleton of every model here)
                                      │
                ┌─────────────────────┴─────────────────────┐
                │                                           │
        E(y) = η  DIRECTLY                        P(y=1) = h(η)  SQUASHED
        ┌───────────────┐                          ┌──────────────┐
        │  2.2 LINEAR   │                          │  2.3 LOGIT   │
        │     MODEL     │                          │              │
        │ y continuous  │                          │  y binary    │
        └───────┬───────┘                          └──────┬───────┘
                │                                         │
```

---

## Branch A — 2.2 The Linear Model

```
2.2 LINEAR MODEL
│
├── 2.2.1 SIMPLE  y = β₀ + β₁x + ε
│   │
│   ├── ESTIMATES
│   │   ├── β̂₁ = Cov(x,y)/Var(x) = r·(sᵧ/sₓ)
│   │   ├── β̂₀ = ȳ − β̂₁x̄
│   │   └── ⭐ line passes through (x̄, ȳ)
│   │       └── recovers a missing intercept from R output
│   │           (Example Exam: 46.61 − 0.90509×48.61 = 2.614)
│   │
│   ├── INTERPRET SLOPE  ⭐ the sentence you write 20×
│   │   └── "one-[unit] ↑ in x is ASSOCIATED WITH an estimated
│   │        β̂₁ [unit] change in the EXPECTED y,
│   │        HOLDING OTHER COVARIATES FIXED"
│   │       └── 4 checkable parts: unit-x · unit-y · "expected" · "associated"
│   │
│   ├── INTERPRET INTERCEPT  ⚠️ usually DON'T
│   │   ├── only if x=0 is IN RANGE and MEANINGFUL
│   │   ├── Sheet 1(c): age 0 ⟹ extrapolation ⟹ don't interpret
│   │   └── FIX: CENTRE the variable  (x − x̄) or (age − 48)
│   │
│   ├── RESIDUALS (with intercept)
│   │   ├── Σε̂ᵢ = 0        ← automatic, NOT evidence of good fit
│   │   ├── Σxᵢε̂ᵢ = 0      ← residuals ⊥ covariates
│   │   └── ŷ̄ = ȳ
│   │
│   └── POLYNOMIAL  y = β₀+β₁x+β₂x²+ε
│       ├── STILL A LINEAR MODEL (linear in β!)
│       ├── ⚠️ effect = β₁ + 2β₂x  ← NOT β₁
│       ├── turning point x* = −β̂₁/(2β̂₂)
│       └── β̂₂ < 0 ⟹ ∩ shape ⟹ maximum
│           (Sheet 2: 5.29/0.10 = 52.9 yrs)
│
└── 2.2.2 MULTIPLE  y = x'β + ε
    │
    ├── INTERPRETATION: PARTIAL effect
    │   └── ⚠️ partial ≠ marginal; can have OPPOSITE SIGNS
    │       └── 🔴 WS23/24 I(i): "β̂ⱼ>0 ⟹ r>0"  FALSE
    │
    ├── ⭐⭐⭐ DUMMY VARIABLES ⭐⭐⭐  (examined EVERY year)
    │   │
    │   ├── 🔴 RULE: c levels ⟹ c−1 dummies
    │   │   └── omitted level = REFERENCE CATEGORY
    │   │       └── = the level MISSING from the R output
    │   │
    │   ├── WHY c−1?  the DUMMY VARIABLE TRAP
    │   │   └── ΣDⱼ = 1 = intercept column
    │   │       └── columns linearly DEPENDENT
    │   │           └── rank(X) < p
    │   │               └── X'X SINGULAR
    │   │                   └── (X'X)⁻¹ doesn't exist
    │   │                       └── ❌ NO UNIQUE OLS
    │   │                           └── ⟹ Ch3 assumes rank(X)=p
    │   │
    │   ├── EVERY dummy coefficient compares TO THE REFERENCE
    │   ├── ⭐ non-reference comparison ⟹ SUBTRACT coefficients
    │   │   └── Advanced vs HS = 62.63 − 11.01 = 51.62
    │   └── shared characteristics CANCEL ⟹ compute Δ directly
    │
    └── INTERACTIONS   y = β₀+β₁x+β₂D+β₃(xD)+ε
        │
        ├── SPLIT INTO TWO LINES immediately:
        │   ├── D=0 :  intercept β₀        slope β₁
        │   └── D=1 :  intercept β₀+β₂     slope β₁+β₃
        │
        ├── ⭐ GEOMETRY: dummy MOVES the line, interaction TILTS it
        │   ├── no interaction → PARALLEL lines
        │   └── interaction    → NON-PARALLEL (may cross)
        │
        ├── ⚠️ β₁ = effect of x ONLY when D=0
        │   ⚠️ β₂ = effect of D ONLY at x=0
        │   └── true effects: ∂y/∂x = β₁+β₃D ; ∂y/∂D = β₂+β₃x
        │
        └── 🔴 any covariate types allowed (cat×cat too)
```

---

## Branch B — 2.3 The Logit Model

```
2.3 LOGIT   (GUARANTEED exam content)
│
├── SETUP:  y ∈ {0,1}  ⟹  E(y|x) = P(y=1) = π ∈ [0,1]
│   └── ⭐ THE MEAN IS A PROBABILITY ⟹ must be bounded
│       └── this ONE fact generates everything below
│
├── 🎯 ANSWER 1 — WHY LINEAR FAILS  (Exam25 Ex4a, 1 pt)
│   ├── (1) 🔴 x'β is UNBOUNDED ⟹ predicts 1.34 and −0.20  ← LEAD WITH THIS
│   ├── (2) Var(y) = π(1−π) ⟹ HETEROSCEDASTIC BY CONSTRUCTION
│   ├── (3) ε takes only TWO VALUES ⟹ can't be normal ⟹ no exact t/F
│   └── (4) constant marginal effect implausible near 0 and 1
│       └── FIX: wrap η in a squashing function h: ℝ → (0,1)
│
├── THREE EQUIVALENT FORMS  (be able to write all three)
│   ├── PROBABILITY  π = eη/(1+eη) = 1/(1+e⁻η)      ← the S-curve
│   ├── ODDS         π/(1−π) = eη                   ← multiplicative
│   └── ⭐ LOG-ODDS  log(π/(1−π)) = η = x'β         ← THE LINK
│       └── this is what makes it a GENERALISED LINEAR model
│           └── linear in g(π), not in π
│
├── 🎯 ANSWER 2 — INTERPRETATION  (Exam25 Ex1h)
│   │
│   ├── 🔴 β̂ⱼ is a change in LOG-ODDS — NOT in probability
│   │
│   ├── SCALE TABLE
│   │   ├── log-odds:    +β̂ⱼ            (exact, constant)
│   │   ├── ⭐ odds:      ×exp(β̂ⱼ)       ← THE ODDS RATIO, say this
│   │   └── probability: β̂ⱼ·π(1−π)      ← NOT constant, depends on π
│   │                     └── max at π=0.5, where it = 0.25·β̂ⱼ
│   │
│   └── only the SIGN transfers reliably (logistic is ↑ increasing)
│
├── PROBIT: same idea, Φ instead of logistic
│   └── ≈ identical fits; logit wins on interpretability (odds ratios)
│
└── ESTIMATION: MAXIMUM LIKELIHOOD, not OLS
    └── no closed form · numerical (IRLS) · asymptotic z/Wald/LR tests
        └── (that's the whole required depth — Ch 5 is out of scope)
```

---

## The compressed picture

```
                    WHAT KIND OF y?
                          │
        ┌─────────────────┴─────────────────┐
        │                                   │
   CONTINUOUS                            BINARY
        │                                   │
   E(y) = x'β                        P(y=1) = h(x'β)
   LINEAR MODEL                          LOGIT
        │                                   │
        │                          β̂ⱼ = Δ log-odds
        │                          exp(β̂ⱼ) = odds ratio
        │
   ┌────┴────┐
   │         │
CONTINUOUS  CATEGORICAL
   x           x
   │           │
 in as-is   c levels
 (or x²,    ⟹ c−1 DUMMIES
  log x)    ⟹ one REFERENCE
   │           │
   └─────┬─────┘
         │
    INTERACTIONS?
         │
    ┌────┴────┐
   no        yes
    │         │
 parallel  non-parallel
  lines      lines
    │         │
    └────┬────┘
         │
  ⚠️ variable in 2+ terms
     ⟹ DIFFERENTIATE
     ⟹ never quote one coefficient
```

**Chapter 2 in one sentence:** *the linear predictor $\boldsymbol{x}'\boldsymbol\beta$ is universal — what changes is how you build $\boldsymbol{x}$ (dummies, polynomials, interactions) and whether you use $\eta$ raw or squash it.*
