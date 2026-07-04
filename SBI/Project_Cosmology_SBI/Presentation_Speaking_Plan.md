# Presentation Speaking Plan — 12 min, 3 speakers
### Saiful · Masum · Rehan   |   deck: `Presentation_Cosmology_SBI.pdf` (11 slides)

> Times are **sharp** — you get cut off at 12:00. This plan totals **~11:30** (buffer built in). Rehearse twice with a timer. **Saiful** anchors the results + live demo + fields the hard Q&A (you did the work). **Masum** opens and closes. **Rehan** owns the "how we built it" block.

---

## ⏱️ Timing at a glance
| Block | Speaker | Slides | Time |
|---|---|---|---|
| Setup / why | **Masum** | 1, 2, 3 | 2:35 |
| How we built it | **Rehan** | 4, 5, 6 | 3:30 |
| Results + LIVE DEMO | **Saiful** | 7, 8, 9, **DEMO**, 10 | 4:50 |
| Close | **Masum** | 11 | 0:30 |
| **Total** | | | **~11:25** |

**Demo placement:** right after **Slide 9** (the degeneracy) — the live banana + recovery bars make the static results come alive at the emotional peak, then you land limitations and the take-home.

---

## 🎤 Slide-by-slide script (what to actually say)

### ▶ MASUM — opening block (2:35)

**Slide 1 · Title (0:25)**
- "Good morning. We're Saiful, Masum, and Rehan. Our project: inferring cosmological parameters from the matter power spectrum using simulation-based inference."
- "We infer four numbers about the Universe — H₀, Ωₘ, nₛ, Aₛ — from a single noisy spectrum."

**Slide 2 · Why this problem matters (1:05)**
- The Universe is clumpy; **P(k)** measures how much structure exists at each scale; galaxy surveys measure it.
- We want the parameters *behind* the spectrum — an **inverse problem**.
- Realistic models have no tractable likelihood, and we want it **fast** → that's why simulation-based inference.

**Slide 3 · The four parameters shape P(k) (1:05)**
- Walk the figure: **Aₛ** = height, **nₛ** = tilt, **Ωₘ** = turnover position, **H₀** = weak effect.
- "So the shape is a fingerprint — and our job is to read it backwards."
- **Hand off:** "Rehan will explain how we built the pipeline that does that."

### ▶ REHAN — build block (3:30)

**Slide 4 · Method: amortized NPE (1:05)**
- Forward: simulator turns parameters into a spectrum. Inverse: a neural network turns a spectrum into a posterior.
- **Amortized** = train once, then inference for any new spectrum is one instant forward pass. **Likelihood-free** = we only need to simulate.

**Slide 5 · Simulator, priors & noise (1:20)**
- Linear P(k) on 100 log-spaced k-bins (CAMB; here a fast BBKS stand-in).
- **Read the priors** (uniform, proper): H₀∈[60,80], Ωₘ∈[0.1,0.5], nₛ∈[0.9,1.05], ln(10¹⁰Aₛ)∈[2.5,3.5].
- **Cosmic-variance noise** σ=P√(2/N): noisy at large scales, precise at small scales.
- We feed **log + standardized** P(k) to the network.

**Slide 6 · Network & training (1:05)**
- MLP → mean + full covariance of a 4-D Gaussian posterior (captures correlations).
- **Two-phase** training (MSE warm-up → Gaussian NLL); Adam; trains on CPU in ~20 s; validation tracks training → no overfit.
- **Hand off:** "Now Saiful will show whether we can *trust* it — and demo it live."

### ▶ SAIFUL — results + demo block (4:50)  ⭐ your anchor block

**Slide 7 · Calibration — is the posterior trustworthy? (1:05)**
- Lead with this: "In a Bayesian project, calibration matters more than a single number."
- **SBC**: ranks ~uniform for nₛ/Aₛ; mild ∪ for Ωₘ/H₀ (slight, honest). **Coverage** ≈ nominal, z-score std ≈ 0.95.
- **Contraction** high for Ωₘ/nₛ/Aₛ, ≈0 for H₀.

**Slide 8 · Parameter recovery (0:40)**
- Estimated vs true: Ωₘ (r=0.97), nₛ (0.91), Aₛ (0.87) land on the diagonal; **H₀ scatters** (r=0.30).
- "Three are nailed; H₀ is special — here's why."

**Slide 9 · The H₀–Ωₘ degeneracy (1:05)**
- The banana: **correlation −0.96**. The spectrum's shape depends on the **combination Γ≈Ωₘ·h**, so H₀ and Ωₘ trade off — H₀ can't be pinned down alone.
- **Key line:** "A broad but calibrated H₀ is the *correct* scientific result, not a failure."

**▶▶ LIVE DEMO (1:20)** — *see the demo script below*

**Slide 10 · Model check & limitations (0:40)**
- PPC: the noisy observation sits within the predictive spread → the model reproduces the data.
- Honest limitations: Gaussian posterior (a flow would improve tails); BBKS vs CAMB; idealized noise; H₀ weak by design.
- **Hand back:** "Masum will wrap up."

### ▶ MASUM — close (0:30)

**Slide 11 · Take-home (0:30)**
- "Three things: (1) a calibrated, amortized posterior for ΛCDM parameters from P(k); (2) Ωₘ, nₛ, Aₛ recovered and calibrated; (3) H₀ shows the expected degeneracy — correct physics."
- Stop. **No thank-you slide.** Go straight to questions.

---

## 🖥️ LIVE DEMO script (Saiful, ~80 sec) — rehearse the clicks!

> Have it **already running** (`go run .`, browser at localhost:8080) on the title slide so there's zero loading time. Alt-tab to it after Slide 9.

1. **(10s)** "This is the whole project, live. The sliders set a *true* Universe; the network never sees them — it only sees the spectrum."
2. **(15s)** Click **Reset to Planck-2018** → "These are the real Universe's values." Point to the **P(k)** plot: "the noisy fingerprint the AI receives — noisy on the left, clean on the right."
3. **(20s)** Point to **Recovered vs True**: "Ωₘ, nₛ, Aₛ — tight bands, gold truth inside. **H₀ — wide.**"
4. **(20s)** Point to the **H₀–Ωₘ scatter**: "There's the degeneracy, live — the banana. H₀ and Ωₘ trade off."
5. **(15s)** Click **Observe & Infer** again: "Each click is a fresh noisy observation — the answer wobbles slightly but stays calibrated. That's honest uncertainty." Alt-tab back to Slide 10.

**Fallback if Wi-Fi/laptop fails:** a screenshot of the demo is in the deck appendix / `figures/` — show that and narrate the same 4 points. **Never debug live.**

---

## 🛡️ Q&A strategy (after the talk — can be ANY of the 12 chapters)
- **Saiful** fields the technical + any-chapter questions (you have the chapter folders — revise the `02_Viva_Questions` files). Default answerer.
- **Masum** takes science/motivation questions (what is P(k), why these parameters, Planck).
- **Rehan** takes simulator/training questions (priors, noise model, network, overfitting).
- If unsure, it's fine to say "good question — [name] worked closely on that" and pass. **Don't bluff.**
- Three answers to have ready cold: (1) *why H₀ is degenerate* (Γ=Ωₘh), (2) *what SBC proves* (calibration), (3) *what amortized means* (train once, infer instantly).

---

## ✅ Pre-talk checklist
- [ ] Each speaker rehearsed their block to time (use a phone timer).
- [ ] Demo running before you start; browser zoomed so the back row can read it.
- [ ] Slides submitted as PDF by **July 19** (separate from presenting!).
- [ ] One laptop is "primary," one is "backup" with the deck + demo.
- [ ] Practice the 2 handoffs out loud ("…now Rehan", "…now Saiful", "…Masum will wrap up").
- [ ] Saiful skimmed all `Chapter_xx/02_Viva_Questions_Bangla.md` the night before.
