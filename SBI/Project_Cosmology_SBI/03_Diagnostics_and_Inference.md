# 03 — Diagnostics & Inference (the highest-value section)

> The rubric explicitly lists **convergence + Simulation-Based Calibration + posterior contraction**, plus inference on real data and posterior predictive checks. This is also where the *Bayesian-calibration-focused* graders look hardest. Do this thoroughly and you secure the report's 50%.
>
> 💡 **বাংলায়:** এই অংশটাই সবচেয়ে বেশি মার্কের। rubric সরাসরি চায়: convergence, SBC, posterior contraction, real-data inference, posterior predictive check। teacher রা (calibration-focused group) এখানেই সবচেয়ে কড়া নজর দেয়। এটা ভালো করলেই report-এর ৫০% নিশ্চিত।

---

## 1. Convergence (the easy one — but don't skip)

Plot **training and validation loss vs epoch**. A healthy run: both decrease and plateau, with the validation loss tracking the training loss (no large gap = no overfitting).

- If val loss diverges upward → overfitting → add regularization / more sims / early stopping.
- Report the final epoch, where you stopped, and why.

> 💡 **বাংলায়:** train ও validation loss-এর curve আঁকো। দুটোই নেমে সমতল হলে ভালো; validation অনেক উপরে থাকলে overfitting (তখন বেশি sim বা regularization দরকার)। এটা সহজ কিন্তু report-এ অবশ্যই থাকতে হবে।

---

## 2. ⭐ Simulation-Based Calibration (SBC) — your core diagnostic (this is Ch9!)

**The question SBC answers:** "Is my posterior *calibrated*?" i.e. when I say a 90% credible interval, does the truth fall inside it 90% of the time?

**How it works (rank statistic):**
1. Draw many test pairs `(θ*, x)` from the simulator (θ* from prior, x from forward model). These are held-out.
2. For each, draw `L` posterior samples and compute the **rank** of the true `θ*` among them (for each parameter separately).
3. If the posterior is perfectly calibrated, those ranks are **uniformly distributed**.

**What to show:**
- **Rank histograms** (one per parameter) — should look flat/uniform.
  - ∪-shape (ranks pile at the ends) → posterior **too narrow** (overconfident).
  - ∩-shape (ranks pile in the middle) → posterior **too wide** (underconfident).
  - tilt → bias.
- **ECDF-difference plots** with simultaneous confidence bands (the modern SBC) — staying inside the band = calibrated.

```python
# BayesFlow provides these diagnostics directly, e.g.:
bf.diagnostics.calibration_histogram(estimates=post_samples, targets=theta_true)
bf.diagnostics.calibration_ecdf(estimates=post_samples, targets=theta_true)   # SBC-ECDF
```

> 💡 **বাংলায়:** SBC জিজ্ঞেস করে — "আমার posterior কি সৎ?" পদ্ধতি: অনেক (সত্য θ*, x) জোড়া নাও, প্রতিটার posterior-এ সত্য θ*-এর rank বের করো। calibrated হলে rank গুলো uniform হবে। ∪-আকৃতি = posterior খুব সরু (overconfident); ∩-আকৃতি = খুব চওড়া। এটা ঠিক Ch9-এর বিষয় — viva-তে এই সংযোগ বললে বড় প্লাস।

> 🎯 **Grade tip:** Even if your posterior is wide (especially H₀), **passing SBC is the win.** Calibration > tightness. Lead the diagnostics section with SBC.

---

## 3. ⭐ Posterior contraction — "did the data teach us anything?"

**Definition:** how much narrower the posterior is than the prior, per parameter:

$$\text{contraction} = 1 - \frac{\mathrm{Var}(\text{posterior})}{\mathrm{Var}(\text{prior})}$$

- ≈ 1 → posterior much tighter than prior → **highly informative** data.
- ≈ 0 → posterior ≈ prior → data uninformative for that parameter.

**Expected pattern (and a talking point):** strong contraction for **Aₛ** (amplitude) and **Ωₘ**, decent for **nₛ**, and **weak contraction for H₀** — exactly the degeneracy story (file `01` §4). Showing this *and explaining why* = expert marks.

Also report the **posterior z-score** `z = (θ_est − θ*)/posterior_sd`; across many test sims it should be ≈ `N(0,1)`. The **z-score vs contraction** scatter ("posterior shrinkage plot") is a great single figure: well-behaved points cluster at high contraction, |z| small.

```python
contraction = 1 - post_var / prior_var          # per parameter, averaged over test set
```

> 💡 **বাংলায়:** contraction মানে posterior কতটা prior-এর চেয়ে সরু হলো — মানে data কতটা শেখাল। Aₛ, Ωₘ-এ বেশি contraction আশা করো, H₀-তে কম (degeneracy-র কারণে)। এটা দেখানো + কারণ ব্যাখ্যা করা = expert মার্ক। z-score ≈ N(0,1) হওয়াও দেখাও।

---

## 4. Recovery / coverage check

- **Recovery plot:** estimated vs true parameter for the held-out test set (scatter around the diagonal), with credible-interval error bars. Tight scatter on the diagonal = good recovery.
- **Coverage:** fraction of test cases where the true value falls in the X% credible interval should ≈ X%. (Closely related to SBC; report at least the 50%/90% empirical coverage.)

```python
bf.diagnostics.recovery(estimates=post_samples, targets=theta_true)
```

> 💡 **বাংলায়:** recovery plot — সত্য vs estimated parameter; diagonal-এর কাছে থাকলে ভালো। coverage — ৯০% credible interval-এ সত্য মান কি ৯০% ক্ষেত্রে পড়ছে? এটাও calibration-এরই অংশ।

---

## 5. Inference on "real" data (rubric: fit to real data)

True galaxy-survey `P(k)` needs galaxy-bias + redshift-space + nonlinear modeling — **out of scope**. So use the standard, defensible proxy and **state it clearly**:

**Recommended approach — Planck-2018 fiducial "observation":**
1. Take the published **Planck 2018** best-fit ΛCDM values (≈ `H₀=67.4, Ωₘ=0.315, nₛ=0.965, Aₛ=2.1e-9`; ref arXiv:1807.06209).
2. Generate the linear `P(k)` at those values with your simulator, add one noise realization → treat as the **observed** spectrum `x_obs`.
3. Run your amortized posterior on `x_obs`.
4. **Show the posterior brackets the Planck values** (corner plot with truth lines), and discuss the H₀–Ωₘ banana.

This is honest (a controlled "real-world-like" test with known ground truth), demonstrates the pipeline end-to-end, and lets you *quantify* success. Optionally, as a stretch, overlay a **published linear P(k)** (e.g. a Planck-derived spectrum) if you can obtain one — but keep the fiducial test as the primary, clearly-labeled result.

> 💡 **বাংলায়:** আসল galaxy survey data-তে galaxy bias ইত্যাদি লাগে — scope-এর বাইরে। তাই standard পদ্ধতি: Planck-2018-এর মান দিয়ে একটা noisy P(k) বানিয়ে সেটাকেই "observed" ধরো, posterior চালাও, দেখাও posterior Planck মানকে ঘিরে রেখেছে। সৎ, ground-truth জানা, আর pipeline পুরোটা দেখায়। report-এ স্পষ্ট বলো এটা fiducial/controlled test।

---

## 6. Posterior predictive checks (PPC)

The Bayesian "does my fitted model reproduce the data?" check (ties to Ch4 simulation-based tests):
1. Draw θ from the posterior (given `x_obs`).
2. For each, simulate a spectrum `P(k)` (+noise).
3. Overlay these posterior-predictive spectra on `x_obs`.

**Good fit:** `x_obs` sits comfortably inside the spread of predicted spectra at all k. Systematic misses at some k → model misspecification.

> 💡 **বাংলায়:** PPC — posterior থেকে θ নিয়ে আবার P(k) simulate করো, observed-এর সাথে overlay করো। observed যদি predicted-দের ছড়ানোর ভেতরে আরাম করে বসে → model ঠিক fit করেছে। এটা Ch4-এর posterior predictive check-এর সরাসরি প্রয়োগ।

---

## 7. The diagnostics checklist (paste into the report as a table)

| Diagnostic | What it proves | Pass criterion | Figure |
|---|---|---|---|
| Train/val loss curve | training converged, no overfit | both plateau, small gap | loss vs epoch |
| **SBC rank / ECDF** | posterior calibrated | flat ranks / inside ECDF band | rank hist + ECDF |
| **Posterior contraction** | data is informative | high for Aₛ/Ωₘ, low for H₀ (explained) | contraction bar / shrinkage plot |
| z-score | unbiased, right width | ≈ N(0,1) | z histogram |
| Recovery / coverage | accurate point + interval | scatter on diagonal; 90%≈90% | recovery scatter |
| Real-data (Planck fiducial) | works on realistic input | posterior brackets truth | corner plot w/ truth lines |
| PPC | model reproduces data | obs within predicted spread | overlaid spectra |

> 🎯 If every row of this table has a passing figure with a caption, your diagnostics section is essentially perfect.

➡️ Next: **`04_Report_and_Presentation_Blueprint.md`** — turn all this into the graded deliverables.
