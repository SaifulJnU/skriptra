# 06 — Related Papers & Why Our Project Is Good

> These are **real, citable references** for your Introduction, Method, and Diagnostics sections. I can't bundle the PDFs (they're external/copyrighted — grab them free from arXiv via the links), but the citations and IDs below are ready to drop into your References.
>
> ✅ The **foundational** IDs (SBI, BayesFlow, SBC, CAMB, Planck) are verified. For the **cosmology-application** papers, double-check the exact arXiv number on arxiv.org before final citation (good scholarly hygiene anyway).
>
> 💡 **বাংলায়:** এগুলো আসল, cite-যোগ্য paper। PDF সরাসরি attach করতে পারছি না (external/copyright), কিন্তু link দিয়েছি — arXiv থেকে free নামাও। নিচের citation গুলো report-এর References-এ সরাসরি বসাও।

---

## A. Foundational SBI / method (cite in Introduction + Method)
1. **Cranmer, Brehmer & Louppe (2020)** — *The frontier of simulation-based inference.* PNAS 117(48). arXiv:**1911.01429**. → the standard overview of the whole field; perfect for the Intro's "what is SBI."
2. **Papamakarios & Murray (2016)** — *Fast ε-free Inference of Simulation Models with Bayesian Conditional Density Estimation.* NeurIPS. arXiv:**1605.06376**. → origin of **neural posterior estimation** (the method class you use).
3. **Greenberg, Nonnenmacher & Macke (2019)** — *Automatic Posterior Transformation for Likelihood-Free Inference.* ICML. arXiv:**1905.07488**. → modern NPE (SNPE-C/APT); the lineage behind amortized posteriors.

## B. The tool you use — BayesFlow (cite in Method; ⭐ note the author!)
4. **Radev, Mertens, Voss, Ardizzone & Köthe (2020)** — *BayesFlow: Learning Complex Stochastic Models With Invertible Neural Networks.* IEEE TNNLS. arXiv:**2003.06281**. → the core BayesFlow method.
5. **Radev, Schmitt, Schumacher, Elsemüller, Pratz, Schälte, Köthe & Bürkner (2023)** — *BayesFlow: Amortized Bayesian Workflows With Neural Networks.* JOSS. arXiv:**2306.16015**. → ⭐ **your course instructor Paul‑Christian Bürkner is a co‑author** — citing this signals you used the group's own tool correctly.

## C. Calibration — your diagnostics (cite in Diagnostics; this is Ch9)
6. **Talts, Betancourt, Simpson, Vehtari & Gelman (2018)** — *Validating Bayesian Inference Algorithms with Simulation-Based Calibration.* arXiv:**1804.06788**. → the **SBC** paper; your single most important diagnostic reference.
7. *(optional)* **Lemos, Coogan, Hezaveh & Perreault‑Levasseur (2023)** — *Sampling-based accuracy testing of posterior estimators for general inference (TARP).* arXiv:**2302.03026** *(verify ID)*. → a modern coverage test, complements SBC.

## D. SBI applied to cosmology — closest to your project (cite in Intro as prior work)
8. **Alsing, Charnock, Feeney & Wandelt (2019)** — *Fast likelihood-free cosmology with neural density estimators and active learning.* MNRAS 488, 4440. arXiv:**1903.00007** (code: `pydelfi`). → ⭐ **the most directly relevant paper**: NPE/DELFI for cosmological parameters, achieving inference with O(10³) sims. Your project is a scaled, BayesFlow-based cousin of this.
9. **Hahn et al. (2023)** — *SimBIG: A Forward Modeling Approach to Analyzing Galaxy Clustering.* arXiv:**2211.00723** *(verify ID)*. → SBI on the galaxy power spectrum / clustering — the "grown-up" version of exactly your problem (real survey data, bias, nonlinearity).
10. *(application examples, verify IDs)* **DES Y3 SBI weak lensing** (arXiv:**2405.10881**); **EFTofLSS meets SBI: σ₈ from biased tracers** (arXiv:**2310.03741**). → show SBI is now mainstream in real cosmological analyses.

## E. Simulator & data (cite in Data + Model)
11. **Lewis, Challinor & Lasenby (2000)** — *Efficient Computation of CMB Anisotropies in Closed FRW Models.* ApJ 538, 473. arXiv:**astro-ph/9911177**. → the **CAMB** code you use to generate `P(k)`.
12. **Planck Collaboration (2020)** — *Planck 2018 results. VI. Cosmological parameters.* A&A 641, A6. arXiv:**1807.06209**. → source of your fiducial parameter values (also in your project brief).

> **Minimal set if you only cite ~6:** #1 (Cranmer), #5 (BayesFlow/Bürkner), #6 (SBC), #8 (Alsing/pydelfi), #11 (CAMB), #12 (Planck). That covers method, tool, calibration, cosmology-SBI, simulator, and data.

---

## Why OUR project is good (the honest, defensible case)

Not because it's novel science — a course project isn't graded on novelty (file `04`/the brief say so). It's good because it is **well-chosen, well-validated, well-understood, and well-engineered** — exactly the axes this course rewards:

1. **Right method, right tool.** Amortized NPE is the *current frontier* of cosmological inference (refs #8–10), and you implement it in **BayesFlow** — the very library your instructor co-authored (#5). You're not reinventing; you're using the field's best practice on a clean problem.

2. **Calibration-first, not accuracy-first.** Most beginner projects chase a tight posterior. You lead with **SBC** (#6) and **posterior contraction** — which is precisely what a Bayesian, calibration-focused group values, and what real SBI cosmology papers insist on. *Proving* trustworthiness > a pretty point estimate.

3. **Controlled validation with known ground truth.** By testing on simulated/Planck‑fiducial spectra where the truth is known, you can *quantify* that the method works before trusting it — mirroring how the field validates (Alsing's O(10³)-sim demos, SBC checks). This is methodologically correct, not a shortcut.

4. **Physical literacy.** You anticipate and explain the **H₀–Ωₘ degeneracy** from h-units — understanding *why* a parameter is poorly constrained. That separates "ran a library" from "understands the inference."

5. **Reproducible & engineered.** Fixed seeds, a manifest, and the optional **Go simulation orchestrator / serving API** (file `05`) give a level of reproducibility and engineering most course groups don't reach — and a memorable presentation demo of *amortization in action*.

6. **Course mastery on display.** The project cleanly exercises the whole syllabus — Monte Carlo (Ch1), sampling priors (Ch3), posterior predictive checks (Ch4), SBC (Ch9), likelihood-free inference (Ch10), neural nets (Ch11), NPE/BayesFlow (Ch12) — which is gold in the post-presentation Q&A.

> **One-line positioning (use it in your intro/conclusion):**
> *"We build a calibrated, amortized neural posterior estimator for ΛCDM parameters from the linear matter power spectrum — a controlled, reproducible BayesFlow implementation of the simulation-based inference approach now used at the frontier of cosmology (e.g. Alsing et al. 2019, SimBIG), with calibration validated by SBC."*

> 💡 **বাংলায়:** আমাদের project "ভালো" নতুনত্বের জন্য নয় (course novelty চায় না), বরং কারণ এটা **সঠিক method (amortized NPE), সঠিক tool (BayesFlow — instructor-এর নিজের), calibration-first (SBC), ground-truth দিয়ে যাচাই, physics বোঝা (H₀ degeneracy), reproducible + engineered, আর পুরো syllabus ছোঁয়।** এই অক্ষগুলোতেই এই course মার্ক দেয়। উপরের এক-লাইনের positioning টা intro/conclusion-এ ব্যবহার করো।

---

## How to use these in the report
- **Introduction:** #1 (SBI overview), #8–10 (SBI in cosmology = motivation + prior work), #12 (why these parameters).
- **Statistical model / Data:** #11 (CAMB), #12 (Planck fiducial).
- **Approximator / Method:** #2–5 (NPE + BayesFlow).
- **Diagnostics:** #6 (SBC), #7 (coverage).
- Keep References on a separate page (excluded from the 10-page text limit).
