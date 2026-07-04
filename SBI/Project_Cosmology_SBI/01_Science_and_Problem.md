# 01 — The Science & the Problem

> Goal of this file: give you (and your teammates) enough cosmology to write a confident Introduction, answer viva questions, and understand *why* the model is built the way it is. You do **not** need to be an astrophysicist — you need ~6 solid concepts. Each has a 💡 বাংলায় intuition note.

---

## 1. The big picture: what is the matter power spectrum `P(k)`?

The Universe is not perfectly smooth — matter clumps into galaxies, clusters, filaments, and voids (the "cosmic web"). Cosmologists describe the lumpiness with the **matter density contrast**:

$$\delta(\mathbf{x}) = \frac{\rho(\mathbf{x}) - \bar\rho}{\bar\rho}$$

i.e. "how much denser than average is this point." Take its Fourier transform `δ(k)`. The **matter power spectrum** `P(k)` is the variance of those Fourier modes:

$$\langle \delta(\mathbf{k})\,\delta^*(\mathbf{k}') \rangle = (2\pi)^3\, \delta_D(\mathbf{k}-\mathbf{k}')\, P(k)$$

In plain words: **`P(k)` tells you how much structure (clumpiness) exists at each spatial scale.** Here `k` is the **wavenumber** (units `h/Mpc`); large `k` = small scales (galaxies), small `k` = huge scales (superclusters).

> 💡 **বাংলায়:** মহাবিশ্বে পদার্থ সমানভাবে ছড়ানো না — কোথাও গুচ্ছ (galaxy), কোথাও ফাঁকা (void)। `P(k)` হলো একটা "scale অনুযায়ী আঁশটেপনার মানচিত্র": প্রতিটা scale-এ কতটা গুঁড়ি-গুঁড়ি structure আছে তার পরিমাপ। ছোট scale (বড় k) = galaxy আকারের গুচ্ছ; বড় scale (ছোট k) = বিশাল supercluster।

### The shape of `P(k)` (you'll plot this in the Intro figure)
- **Large scales (small k):** `P(k) ∝ k^{nₛ}` — rises, set by the primordial spectrum.
- **A turnover** near the **matter–radiation equality scale** `k_eq ∝ Ωₘh²` — the single most information-rich feature.
- **Small scales (large k):** falls roughly `∝ k^{nₛ-4}` with logarithmic corrections (the "transfer function" suppresses small-scale growth that happened during the radiation era).
- **Baryon Acoustic Oscillations (BAO):** small wiggles imprinted by sound waves in the early baryon–photon plasma; their amplitude depends on the baryon fraction.

> 💡 **বাংলায়:** P(k)-এর curve-টা একটা পাহাড়ের মতো: বাঁ দিকে ওঠে, একটা চূড়া (turnover, `k_eq`) থেকে ডান দিকে নামে, আর গায়ে ছোট ছোট ঢেউ (BAO)। এই curve-এর **আকৃতিই** parameter-গুলোর তথ্য বহন করে।

---

## 2. The four parameters θ = {H₀, Ωₘ, nₛ, Aₛ} — what each one *does* to `P(k)`

This is the heart of the project: each parameter changes the curve in a characteristic way, and the network learns to read those changes backwards.

| Parameter | Meaning | Typical value | Effect on `P(k)` |
|---|---|---|---|
| **H₀** | Hubble constant (expansion rate today), `h = H₀/100` | ~67–70 km/s/Mpc | Sets units & enters `k_eq ∝ Ωₘh²`; shifts the turnover. **Weakly constrained alone** (see §4). |
| **Ωₘ** | Total matter density fraction | ~0.31 | Moves the turnover scale `k_eq`; more matter → turnover at larger k, more small-scale power. |
| **nₛ** | Scalar spectral index (primordial "tilt") | ~0.965 | Tilts the whole spectrum: `nₛ<1` means slightly less power on small scales (a gentle clockwise tilt). |
| **Aₛ** | Primordial amplitude of fluctuations | ~2.1×10⁻⁹ | Overall **vertical** amplitude of `P(k)` — bigger Aₛ → more power at all scales. |

> 💡 **বাংলায়:** ভাবো P(k) curve-টা একটা গ্রাফ, আর ৪টা parameter ৪টা knob:
> - **Aₛ** = পুরো curve-কে উপরে-নিচে তোলা (amplitude)।
> - **nₛ** = curve-টাকে কাত করা (tilt)।
> - **Ωₘ** = চূড়া (turnover) বাঁয়ে-ডানে সরানো।
> - **H₀** = scale/unit ঠিক করা; কিন্তু এর effect দুর্বল ও Ωₘ-এর সাথে জড়ানো (পরে §4)।
> Network-টা শেখে: curve দেখে এই ৪টা knob-এর সম্ভাব্য অবস্থান (posterior) বলা।

### Fixed (nuisance) quantities — to keep the model minimal
- **Baryon density** `ω_b = Ω_b h² ≈ 0.0224` — **held fixed** (well-measured by CMB; freeing it adds little for this scope).
- **Neutrino density** `ω_ν` — **neglected / fixed** (massless neutrinos default in CAMB).
- **Cold dark matter density** is then *derived*: `ω_c = Ωₘ h² − ω_b − ω_ν`. (This is why CAMB needs `omch2` computed from your sampled H₀ and Ωₘ — see `02` §3.)

> 💡 **বাংলায়:** CAMB physical density (`ω = Ω h²`) চায়, সরাসরি `Ωₘ` না। তাই প্রতিবার sampled H₀, Ωₘ থেকে `ω_c = Ωₘh² − 0.0224` হিসাব করে CAMB-কে দিতে হবে। এটাই project-এর একটা গুরুত্বপূর্ণ technical detail (report-এ লিখতে হবে)।

---

## 3. Why infer parameters *from* `P(k)`? (the scientific motivation for the Intro)

The large-scale structure of the Universe is one of our two great cosmological data sources (the other is the CMB). The matter power spectrum is **the** primary summary statistic of that structure: galaxy surveys (SDSS, DESI, Euclid) measure something close to it. Fitting `P(k)` constrains the **ΛCDM** model and the parameters above — e.g. how fast the Universe expands (H₀), how much matter it contains (Ωₘ), and the nature of the primordial fluctuations (nₛ, Aₛ). Precise, *calibrated* constraints are how we test whether ΛCDM is right and chase tensions like the famous **"Hubble tension"** (different methods disagree on H₀).

> 💡 **বাংলায়:** P(k) হলো মহাবিশ্বের structure-এর মূল "summary statistic" — survey-রা এটাই মাপে। এটা থেকে parameter বের করা মানে আসলে মহাবিশ্বের মৌলিক ধর্ম (কত দ্রুত প্রসারণ, কত matter) মাপা। Intro-তে এই motivation + একটা P(k) curve-এর figure দিলে শক্ত শুরু হবে।

---

## 4. ⭐ The H₀ degeneracy — your "expert insight" (worth real marks)

A subtle but important point that graders love. We express `k` in `h/Mpc` and `P(k)` in `(Mpc/h)³`. In these "h-units," the **shape** of the linear spectrum depends almost entirely on the *physical* densities `ω_m = Ωₘh²` and `ω_b = Ω_b h²` — **not on H₀ separately**. Consequences:

- **H₀ is only weakly constrained** by the shape of the linear `P(k)` alone.
- **H₀ and Ωₘ are degenerate:** many (H₀, Ωₘ) pairs that keep `Ωₘh²` roughly constant give nearly identical spectra.
- The real handle on the absolute scale comes from the **turnover position `k_eq ∝ Ωₘh²`** and the BAO scale.

**Why this matters for your grade:** When your posterior shows a broad/degenerate H₀ and a banana-shaped H₀–Ωₘ contour, that is **not a bug — it's correct physics.** Predicting it in advance and explaining it in the report/viva demonstrates genuine understanding, exactly the kind of thing that turns a 1.7 into a 1.0.

> 💡 **বাংলায়:** এটাই তোমার "expert চমক"। h-unit-এ P(k)-এর আকৃতি মূলত `Ωₘh²`-এর ওপর নির্ভর করে, একা H₀-এর ওপর না। তাই H₀-এর posterior চওড়া হবে আর Ωₘ-এর সাথে কলা-আকৃতির (banana) degeneracy দেখাবে — এটা ভুল না, এটাই সঠিক বিজ্ঞান। report/viva-তে এটা আগে থেকে বলে দিলে teacher বুঝবে তুমি সত্যিই বুঝেছ।

---

## 5. Why this is a *simulation-based inference* problem (the SBI framing)

You could, in principle, write a Gaussian likelihood here (the noise is Gaussian). But the project's purpose is to learn the **amortized, likelihood-free** workflow:

- We have a **simulator** `θ → P(k) → noisy P̂(k)` (CAMB + noise). It's easy to *sample from* but we treat the mapping as a black box.
- We want the **posterior** `p(θ | P̂(k))` — what parameters could have produced this spectrum, with full uncertainty.
- **Amortized NPE** trains a neural network once on many simulations so that inference for any new spectrum is instant. This is the ABC → neural posterior estimation story of Chapters 10–12.

The forward/inverse picture:

```
        PRIOR p(θ)                 SIMULATOR (CAMB + noise)
   θ = {H0, Ωm, ns, As}  ───────────────────────────────────►   P̂(k)  (noisy spectrum)
        ▲                                                          │
        │            NEURAL POSTERIOR  q(θ | P̂(k))  ◄──────────────┘
        └──────────────  (what BayesFlow learns)  ─────────────────
```

> 💡 **বাংলায়:** আমাদের কাছে একটা simulator আছে (parameter → P(k)), কিন্তু উল্টোটা (P(k) → parameter, posterior সহ) কঠিন। BayesFlow একবার অনেক simulation-এ শিখে নেয়, তারপর যেকোনো নতুন P(k)-র জন্য সাথে সাথে posterior দেয় — একেই বলে amortized। এটাই ch10-12-এর মূল গল্পের বাস্তব প্রয়োগ।

---

## 6. Glossary (keep handy for viva)

- **ΛCDM** — the standard cosmological model: cold dark matter + cosmological constant Λ (dark energy).
- **Wavenumber `k`** — inverse scale; `h/Mpc` units.
- **`h`** — `H₀/100`; absorbs the Hubble uncertainty in distance units.
- **Transfer function** — encodes how primordial fluctuations evolved into today's `P(k)` (the small-scale suppression).
- **`k_eq`** — matter–radiation equality scale; the turnover; `∝ Ωₘh²`.
- **BAO** — baryon acoustic oscillations; the wiggles.
- **Cosmic variance** — irreducible uncertainty because we have only one Universe / finite volume; sets the noise floor (see `02` §4).
- **σ₈** — a common derived amplitude measure (RMS fluctuation in 8 Mpc/h spheres); related to Aₛ. You can report it as a derived quantity for bonus.
- **Redshift `z`** — cosmic time proxy; we work at `z=0` (today).

➡️ Next: **`02_Implementation_Guide.md`** — the actual code path.
