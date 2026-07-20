# Ch 1 — MIND MAP

```
                        ┌─────────────────────────────────┐
                        │   y = f(x) + ε                  │
                        │   observed = signal + noise     │
                        └────────────┬────────────────────┘
                                     │
              ┌──────────────────────┼──────────────────────┐
              │                      │                      │
       ┌──────▼──────┐        ┌──────▼──────┐        ┌──────▼──────┐
       │  1.1 WHAT   │        │  1.2 LOOK   │        │ 1.3 WRITE   │
       │ regression  │        │  at data    │        │  notation   │
       │    is for   │        │ before you  │        │             │
       └──────┬──────┘        │    model    │        └──────┬──────┘
              │               └──────┬──────┘               │
              │                      │                      │
    ┌─────────┴─────────┐   ┌────────┴────────┐   ┌─────────┴─────────┐
    │                   │   │                 │   │                   │
    ▼                   ▼   ▼                 ▼   ▼                   ▼
```

---

## Branch 1 — WHAT regression is for (1.1)

```
1.1 APPLICATIONS
│
├── ORIGIN
│   └── Galton 1885 · heights · "regression to mediocrity"
│       └── y = β₀ + β₁x + ε · fitted by least squares: min Σ(yᵢ − β₀ − β₁xᵢ)²
│
├── THE VOCABULARY
│   ├── y  → response / dependent / regressand / target
│   ├── x  → covariate / explanatory / independent / regressor / predictor
│   ├── f(x) → systematic component
│   └── ε  → random error
│
├── ⭐ THE MODEL CLASS IS CHOSEN BY THE TYPE OF y
│   ├── continuous ────────► CLASSICAL LINEAR MODEL   (Ch 3) ← this course
│   ├── binary 0/1 ────────► LOGIT / PROBIT           (Ch 2.3)
│   ├── count 0,1,2,… ─────► POISSON                  (Ch 5)
│   └── ordered categories ► ORDINAL                  (Ch 6)
│
├── ⭐ "LINEAR" = LINEAR IN β, NOT IN x
│   ├── ✅ β₀ + β₁x + β₂x²        (polynomial)
│   ├── ✅ β₀ + β₁log(x)
│   ├── ✅ β₀ + β₁x₁ + β₂x₂ + β₃x₁x₂  (interaction)
│   ├── ✅ y = exp(x'β + ε)  →  log y = x'β + ε
│   └── ❌ β₀ + x^β₁
│
└── THREE GOALS (don't confuse them)
    ├── DESCRIPTION  → sign/size/significance of βⱼ  → Ch 3.3
    ├── PREDICTION   → ŷ good on NEW data            → Ch 3.4
    └── CAUSALITY    → needs random assignment       → beyond course
        └── say "associated with", never "causes"
```

---

## Branch 2 — LOOK at the data (1.2)

```
1.2 FIRST STEPS
│
├── 1.2.1 ONE VARIABLE AT A TIME
│   ├── continuous → histogram · boxplot · mean vs median · s²
│   │   └── ⭐ mean > median ⟹ RIGHT-SKEWED
│   │       └── wage, rent, income, price
│   │           └── CONSEQUENCES: non-normal ε · heteroscedasticity
│   │               └── REMEDY: model log(y)
│   │                   └── interpretation becomes % → see log table
│   └── categorical → frequency table · bar chart
│       └── count the levels! c levels ⟹ c−1 dummies (Ch 3.1.3)
│
├── 1.2.2 TWO VARIABLES AT A TIME
│   │
│   ├── SCATTER PLOT — ask 4 questions
│   │   ├── (1) is there a trend?
│   │   ├── (2) is it STRAIGHT?   → if not: x², log x   (Ch 3.1.3)
│   │   ├── (3) is spread CONSTANT? → if not: heteroscedasticity (Ch 4.1.3)
│   │   └── (4) OUTLIERS?          → leverage, Cook's D (Ch 3.4.4)
│   │       ⭐ these same 4 questions reappear as RESIDUAL PLOTS in 3.4.4
│   │
│   ├── CORRELATION  r = Cov(x,y)/(sₓsᵧ) ∈ [−1,1]
│   │   ├── ⭐ R² = r²  (SIMPLE regression only, with intercept)
│   │   │   └── sign of r = sign of β̂₁
│   │   └── THREE LIMITS
│   │       ├── only sees STRAIGHT lines   (y = x² ⟹ r = 0)
│   │       ├── not CAUSAL                 (confounders)
│   │       └── ANSCOMBE: same stats, different pictures
│   │
│   └── categorical x vs continuous y → SIDE-BY-SIDE BOXPLOTS
│
└── PHILOSOPHY OF THE BRANCH
    └── "The picture tells you what model to build."
```

---

## Branch 3 — WRITE it down (1.3)

```
1.3 NOTATION
│
├── COUNTS
│   ├── n = observations,  i = 1…n
│   ├── k = covariates,    j = 1…k
│   └── ⚠️ p = k+1 = PARAMETERS  (book) ... but some papers use p = k
│       └── 🛟 SAFE RULE: residual df = n − (number of β's) = n − k − 1
│
├── MATRIX FORM      y = Xβ + ε
│   ├── y : n×1
│   ├── X : n×p   ← ONE ROW PER OBSERVATION, ONE COLUMN PER PARAMETER
│   │              ← FIRST COLUMN ALL ONES (the intercept)
│   ├── β : p×1
│   └── ε : n×1
│       └── dimension check: (n×p)(p×1) = n×1 ✓
│
├── DERIVED OBJECTS (all show up in Ch 3)
│   ├── X'X            p×p   symmetric, invertible ⟺ full rank
│   ├── (X'X)⁻¹        p×p
│   ├── β̂ = (X'X)⁻¹X'y p×1
│   ├── H = X(X'X)⁻¹X' n×n   the HAT matrix
│   └── x₀'(X'X)⁻¹x₀   SCALAR (prediction intervals)
│
├── ⭐ THE HAT RULE
│   ├── NO hat → true, unknown, FIXED     (β, ε, σ²)
│   └── hat    → from data, RANDOM        (β̂, ε̂, σ̂²)
│       └── unbiasedness: E(β̂) = β   ← the hat vanishes under E
│
├── SUMS OF SQUARES
│   ├── SSE = ε̂'ε̂ = Σε̂ᵢ²                    (residual)
│   ├── explained SS = Σ(ŷᵢ − ȳ)²
│   ├── SST = Σ(yᵢ − ȳ)²
│   ├── SST = SSE + explained SS
│   ├── R² = 1 − SSE/SST
│   └── ⚠️ NEVER write "SSR" — ambiguous
│
└── ⭐ THE ONE MATRIX IDENTITY TO MEMORISE
    └── Cov(Az) = A · Cov(z) · A'
        └── gives Cov(β̂) = σ²(X'X)⁻¹
            └── gives every standard error, t-test and CI in the course
```

---

## The compressed version (if you only remember one picture)

```
                    y = f(x) + ε
                         │
        ┌────────────────┴────────────────┐
        │                                 │
   TYPE OF y                        SHAPE OF f
   picks the MODEL                  picks the DESIGN MATRIX
        │                                 │
   continuous → LM                   x, x², log x,
   binary     → logit                dummies, interactions
   count      → Poisson              (all still LINEAR in β)
        │                                 │
        └────────────────┬────────────────┘
                         │
                    y = Xβ + ε
                         │
                  minimise ε̂'ε̂
                         │
                β̂ = (X'X)⁻¹X'y          ← Chapter 3.2
                         │
              how confident may I be?     ← Chapter 3.3
                         │
              which model is best?        ← Chapter 3.4
```

**Everything in this course hangs off that last column.**
