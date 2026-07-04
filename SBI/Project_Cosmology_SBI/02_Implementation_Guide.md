# 02 — Implementation Guide (step-by-step, with code)

> This is the build path: environment → simulator (CAMB) → priors → noise → preprocessing → BayesFlow networks → training. Code is **skeleton-level** and faithful to BayesFlow ≥ 2.0 structure; exact argument names can shift between minor versions, so keep the [BayesFlow docs](https://bayesflow.org/) open and adapt. We'll refine all of this together when we implement.
>
> 💡 **বাংলায়:** এই ফাইলটা actual কোডের রোডম্যাপ। প্রতিটা block-এর নিচে `💡 বাংলায়` দিয়ে "কেন এটা করছি" বলা আছে। API-র নাম version-ভেদে একটু বদলাতে পারে — docs মিলিয়ে নিও।

---

## 1. Environment setup (Colab / Kaggle)

```bash
pip install camb bayesflow getdist
# BayesFlow uses Keras 3 (multi-backend). Pick a backend BEFORE importing bayesflow:
```
```python
import os
os.environ["KERAS_BACKEND"] = "jax"     # or "torch" / "tensorflow"; set on a GPU runtime
import numpy as np, camb, bayesflow as bf
```
- On Colab/Kaggle: enable the **GPU runtime** for training (CAMB itself runs on CPU).
- Set a global **seed** everywhere for reproducibility (Ch2!): `np.random.seed(0)` and the framework seed.

> 💡 **বাংলায়:** CAMB CPU-তে চলে (P(k) বানাতে), কিন্তু network train করতে GPU দরকার — তাই Colab/Kaggle-এ GPU runtime চালু করো। seed fix করা বাধ্যতামূলক (reproducibility = Ch2-এর শিক্ষা, আর report-এ ভালো দেখায়)।

---

## 2. The k-grid (fixed once, used everywhere)

```python
N_K = 100
k_grid = np.logspace(-3, 0, N_K)        # 100 log-spaced wavenumbers in [1e-3, 1] h/Mpc
log_k  = np.log10(k_grid)
```
Log-spacing matches the project spec and the physics (features are roughly log-uniform). **Every** simulated spectrum is evaluated on this exact grid so the network always sees a fixed-length, aligned input.

---

## 3. The simulator: θ → clean `P(k)` via CAMB

```python
OMEGA_B_H2 = 0.0224          # fixed baryon density (ω_b)
def simulate_pk(H0, Om, ns, As, kgrid=k_grid):
    h = H0 / 100.0
    omch2 = Om * h**2 - OMEGA_B_H2          # derive cold dark matter physical density
    if omch2 <= 0:                          # guard: invalid cosmology
        return None
    pars = camb.set_params(H0=H0, ombh2=OMEGA_B_H2, omch2=omch2, ns=ns, As=As)
    pars.set_matter_power(redshifts=[0.0], kmax=2.0)
    pars.NonLinear = camb.model.NonLinear_none      # LINEAR spectrum (project requirement)
    results = camb.get_results(pars)
    # Evaluate on OUR fixed grid via the interpolator (guarantees identical k for all sims):
    PK = results.get_matter_power_interpolator(nonlinear=False, hubble_units=True, k_hunit=True)
    pk = PK.P(0.0, kgrid)                    # P(k) at z=0, shape (N_K,), units (Mpc/h)^3
    return pk
```

Key points to put in the report:
- `hubble_units=True, k_hunit=True` → `k` in `h/Mpc`, `P` in `(Mpc/h)³` (the h-unit convention that drives the H₀ degeneracy, file `01` §4).
- `NonLinear_none` → **linear** spectrum (keeps the simulator fast and matches scope).
- The `omch2 <= 0` guard handles the low-Ωₘ / low-H₀ corner.

> 💡 **বাংলায়:** এই function-টাই তোমার "likelihood/simulator" — parameter দিলে P(k) curve ফেরত দেয়। `omch2 = Ωₘh² − ω_b` হিসাবটা না করলে CAMB ভুল হবে। interpolator দিয়ে আমাদের নিজের k_grid-এ মূল্যায়ন করছি যাতে সব simulation একই grid-এ থাকে।

**Sanity check first (Week 1):** plot `simulate_pk` for Planck values `(H0=67.4, Om=0.315, ns=0.965, As=2.1e-9)` and for a few perturbed values; confirm the turnover, tilt, and amplitude move as file `01` §2 predicts. This becomes your **Intro figure**.

---

## 4. Priors (explicit & proper — the rubric demands this verbatim)

```python
def sample_prior(batch_size, rng=np.random):
    H0 = rng.uniform(60.0, 80.0,  batch_size)      # km/s/Mpc
    Om = rng.uniform(0.10, 0.50,  batch_size)      # total matter fraction
    ns = rng.uniform(0.90, 1.05,  batch_size)      # spectral index
    lnAs = rng.uniform(2.5, 3.5,  batch_size)      # ln(10^10 As)  → sample log-amplitude
    As = 1e-10 * np.exp(lnAs)                       # As ∈ ~[1.2e-9, 3.3e-9]
    return np.stack([H0, Om, ns, lnAs], axis=-1)   # keep lnAs as the inference variable
```

**Why these choices (explain in report):**
- **Uniform, wide-but-physical** ranges bracketing Planck 2018 values → "proper" (integrate to 1) and weakly informative.
- **Sample `ln(10¹⁰Aₛ)` instead of `Aₛ`** because Aₛ spans orders of magnitude and enters the amplitude ~linearly; the log makes the prior and the learning well-scaled (this is itself a Ch1 change-of-variables point!).
- The ranges keep `ω_c = Ωₘh² − ω_b > 0` across essentially the whole box (check: min `Ωₘh²` ≈ 0.10·0.6² = 0.036 > 0.0224 ✓).

> 💡 **বাংলায়:** prior = parameter-গুলোর "আগাম বিশ্বাস" — কোন মানগুলো সম্ভব। আমরা Planck মানের চারপাশে wide uniform নিচ্ছি (physical + proper)। Aₛ-এর বদলে `ln(10¹⁰Aₛ)` sample করছি কারণ Aₛ অনেক ছোট সংখ্যা (orders of magnitude জুড়ে) — log নিলে scale ভালো হয় (এটা Ch1-এর change-of-variables!)। ⚠️ report-এ এই priors হুবহু লিখতে হবে — rubric চায়।

---

## 5. The noise model (cosmic variance) — turning clean `P(k)` into an "observation"

```python
V = 1e9                                  # survey volume (Mpc/h)^3
# log-spaced bin widths Δk_i:
dk = np.gradient(k_grid)
N_modes = k_grid**2 * dk * V / (2*np.pi**2)      # number of Fourier modes per k-bin

def add_noise(pk_clean, rng=np.random):
    sigma = pk_clean * np.sqrt(2.0 / N_modes)    # cosmic-variance error per bin
    return pk_clean + sigma * rng.standard_normal(pk_clean.shape)
```

**Physics (for report/viva):** each k-bin contains `N_i ≈ k²Δk·V/(2π²)` independent Fourier modes; the power averaged over them has a relative error `√(2/N_i)` (Gaussian approximation to the χ²-with-`N_i`-modes distribution). Small k (few modes) → noisy; large k (many modes) → precise. `V` sets the overall noise floor.

- **Baseline alternative** (mention as a simpler variant): `P̂ = P·(1+ε)`, `ε ~ N(0, σ²)` with a fixed fractional σ.
- **Design choice that helps training:** precompute the **clean** `P(k)` (expensive CAMB) once, then add **fresh noise each batch** during training. The network sees the same cosmology under many noise realizations → better robustness, effectively free data augmentation.

> 💡 **বাংলায়:** আমরা শুধু একটা Universe দেখি, তাই finite volume-এর কারণে অনিবার্য "cosmic variance" noise থাকে। প্রতিটা k-bin-এ যত বেশি mode (`N_i`), তত কম noise — তাই বড় k নিখুঁত, ছোট k noisy। চালাকি: clean P(k) একবার বানিয়ে রাখো (CAMB ব্যয়বহুল), কিন্তু noise প্রতি batch-এ নতুন করে যোগ করো — এতে network বেশি robust হয় (free data augmentation)।

---

## 6. Precompute the training set (offline)

```python
N_SIM = 50_000                            # in [1e4, 1e5]; start smaller, scale up
theta = sample_prior(N_SIM)
pk_clean = np.zeros((N_SIM, N_K))
keep = np.ones(N_SIM, bool)
for i in range(N_SIM):
    out = simulate_pk(theta[i,0], theta[i,1], theta[i,2], 1e-10*np.exp(theta[i,3]))
    if out is None: keep[i] = False
    else: pk_clean[i] = out
theta, pk_clean = theta[keep], pk_clean[keep]
np.savez_compressed("cosmo_sims.npz", theta=theta, pk_clean=pk_clean,
                    k_grid=k_grid, N_modes=N_modes)
```
- **Offline regime** (Ch12 term): generate once, train many epochs on the cached set. Necessary because CAMB is the bottleneck.
- 50k linear spectra ≈ minutes-to-an-hour on CPU; run it once, reuse. Save to Google Drive so Colab restarts don't lose it.
- Hold out ~10–20% as a **validation/test** set (never trained on) for diagnostics.

> 💡 **বাংলায়:** offline training মানে — আগে একবার সব P(k) বানিয়ে file-এ রাখো, তারপর সেই cached data দিয়ে বহুবার train করো (Ch12-এর offline regime)। CAMB ধীর বলে এটা জরুরি। Drive-এ save করো যাতে Colab restart-এ হারিয়ে না যায়।

---

## 7. Preprocessing (the adapter) — critical, or training fails

`P(k)` spans many orders of magnitude → feed **standardized log P(k)**:

```python
log_pk = np.log10(pk_clean)                       # compress dynamic range
# standardize per k-bin using TRAIN stats:
mu, sd = log_pk.mean(0), log_pk.std(0)
def preprocess_pk(pk):                             # apply at train & inference time
    return (np.log10(pk) - mu) / sd
```
In BayesFlow you encode this in an **Adapter**, which also standardizes the inference variables (θ). Schematically:

```python
adapter = (bf.adapters.Adapter()
           .to_array()
           .as_set(False)                          # P(k) is an ORDERED curve, not a set
           .apply(include="pk", forward=lambda x: (np.log10(x)-mu)/sd)   # log+standardize
           .standardize(include=["H0","Om","ns","lnAs"])                 # scale params
           .concatenate(["H0","Om","ns","lnAs"], into="inference_variables")
           .concatenate(["pk"], into="summary_variables"))
```

> ⚠️ **Do NOT treat `P(k)` as a permutation-invariant set / DeepSet.** The ordering along `k` carries the physics (turnover, tilt, BAO). Use an order-aware summary network (§8).

> 💡 **বাংলায়:** raw P(k) বিশাল রেঞ্জ জুড়ে (10³ থেকে 10⁵...), তাই সরাসরি দিলে network শিখতে পারে না। `log10` নিয়ে compress করি, তারপর standardize (mean 0, std 1) করি। parameter-গুলোও standardize করি। ⚠️ P(k)-কে "set" ধরো না (DeepSet ব্যবহার করো না) — কারণ k-এর ক্রমেই (turnover, tilt, BAO) আসল তথ্য।

---

## 8. The networks (Approximator = summary + inference)

### 8a. Summary network — compresses the 100-point spectrum to a small embedding
The spectrum is a fixed-length **ordered 1-D curve**. Two good options:
- **Baseline (robust, easy):** an **MLP** summary network (since input length is fixed at 100). BayesFlow's `MLP`-style summary or a small custom Keras MLP → output dim ~ 8–16.
- **Upgrade (for extra marks):** a **1-D CNN** over the k-axis — convolutions naturally capture *local* shape features (turnover, BAO wiggles). Use a custom Keras `Sequential` of `Conv1D` layers as `summary_network`.

```python
summary_net = bf.networks.MLP(widths=[128, 128, 64])      # baseline; outputs an embedding
# (Upgrade: a custom keras Conv1D stack passed as summary_network)
```

### 8b. Inference network — the normalizing flow that represents the posterior
4 parameters → low-dimensional posterior. A **coupling flow** is fast and stable:

```python
inference_net = bf.networks.CouplingFlow()        # baseline normalizing flow
# Alternative (more expressive, slower): bf.networks.FlowMatching()
```

### 8c. Approximator — glue them together
```python
approximator = bf.approximators.ContinuousApproximator(
    summary_network   = summary_net,
    inference_network = inference_net,
    adapter           = adapter,
)
```

**Report must specify** (rubric): number of layers, widths, activations, any dropout/regularization, embedding dim, flow type & number of coupling layers. Write these down as you finalize.

> 💡 **বাংলায়:** দুটো network: (১) **summary network** ১০০-পয়েন্ট P(k)-কে ছোট একটা vector-এ চাপে (curve-এর সারমর্ম); (২) **inference network** (normalizing flow) সেই সারমর্ম থেকে posterior `q(θ|x)` বানায়। CouplingFlow সহজ ও দ্রুত (৪টা parameter-এ যথেষ্ট)। Conv1D summary network নিলে BAO/turnover-এর local feature ভালো ধরবে — বাড়তি মার্ক। report-এ layer/width/activation সব লিখতে হবে।

---

## 9. Training

```python
data = np.load("cosmo_sims.npz")
# A simulator that yields {params..., pk} dicts; add fresh noise per batch:
def online_batch(batch_size, rng=np.random):
    idx = rng.integers(0, len(data["theta"]), batch_size)
    th  = data["theta"][idx]
    pk  = add_noise(data["pk_clean"][idx], rng)        # fresh noise each draw
    return dict(H0=th[:,0], Om=th[:,1], ns=th[:,2], lnAs=th[:,3], pk=pk)

history = approximator.fit(
    simulator   = online_batch,    # or a BayesFlow Simulator wrapping prior+forward
    epochs      = 100,
    batch_size  = 512,
    # optimizer Adam + cosine-decay LR schedule are BayesFlow-friendly defaults
)
```

- **Regime:** clean spectra precomputed (offline) + **noise added online per batch** → best of both.
- **Budget to report:** epochs (~50–150 until val loss plateaus), batch size (256–512), steps/epoch, optimizer (**Adam**), **learning-rate schedule** (cosine decay), early stopping on validation loss.
- **Loss:** NPE maximizes the (approximate) posterior log-density of the true θ — BayesFlow handles this internally; just track train/val loss.

> 💡 **বাংলায়:** train করার সময় cached clean P(k) থেকে batch নিই, প্রতি batch-এ নতুন noise যোগ করি। Adam optimizer + cosine LR schedule, validation loss না কমা পর্যন্ত epoch চালাই (early stopping)। report-এ epoch/batch/optimizer/LR সব লিখতে হবে (rubric)।

---

## 10. Inference for one spectrum (after training)

```python
# x_obs: a single noisy spectrum (shape (N_K,)) — simulated test or Planck-fiducial
post = approximator.sample(conditions={"pk": x_obs[None, :]}, num_samples=2000)
# post → 2000 posterior draws of {H0, Om, ns, lnAs}; convert lnAs→As, derive σ8 if you like
```
These draws are everything: medians = point estimates, quantiles = credible intervals, pairwise scatter = the corner plot (where you'll *see* the H₀–Ωₘ banana). Diagnostics & real-data use of these draws → file `03`.

> 💡 **বাংলায়:** train শেষে যেকোনো P(k)-র জন্য সাথে সাথে ২০০০টা posterior draw পাও (amortized মানেই এই গতি)। এই draw থেকেই median (point estimate), credible interval, আর corner plot (H₀–Ωₘ degeneracy চোখে দেখা) — সব।

---

## 11. Common pitfalls (each has cost you real marks if missed)

| Pitfall | Fix |
|---|---|
| Raw `P(k)` to the network | always log + standardize (§7) |
| `Aₛ` sampled linearly | sample `ln(10¹⁰Aₛ)` (§4) |
| `omch2 ≤ 0` at low Ωₘ/H₀ | guard / reject (§3) |
| DeepSet on the spectrum | use order-aware summary net (§8a) |
| Different k-grid per sim | always evaluate on fixed `k_grid` via interpolator (§3) |
| Losing sims on Colab restart | save `.npz` to Google Drive (§6) |
| Reporting a suspiciously tight H₀ | expect & explain the degeneracy (file `01` §4) |
| No held-out set | reserve 10–20% for diagnostics (§6) |

➡️ Next: **`03_Diagnostics_and_Inference.md`** — proving the posterior is trustworthy (the highest-value section).
