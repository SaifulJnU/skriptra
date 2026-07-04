# 05 — Where Your Go / Backend Skills Add Real Value (optional, but smart)

> Honest framing first, because it protects your grade.

## The honest constraint
The **graded core is Python**: BayesFlow (the NPE engine) and CAMB (the simulator) are Python/Fortran. The neural inference itself **cannot** be done in Go. And the project rubric explicitly warns: *"keep to a minimal working model so that you do not become distracted by too many moving parts."* So:

- ❌ Don't try to rewrite the ML in Go. It won't earn project marks and risks the "too many moving parts" trap.
- ✅ Use Go where backend engineering *genuinely* helps: **around** the Python core — data generation, reproducibility, and serving. Present it as **engineering rigor** (a differentiator in the Reflection section, the presentation, and the viva), not as the science.
- ✅ Only after the Python core works end-to-end.

> 💡 **বাংলায়:** ML অংশটা Python-ই থাকবে (BayesFlow/CAMB) — Go দিয়ে inference করা যাবে না, আর rubric বলেছে "জিনিস বেশি জটিল করো না।" তাই Go-কে **মূল কোডের চারপাশে** infrastructure হিসেবে ব্যবহার করো (data generation, reproducibility, serving) — যেটা তোমার শক্তি, আর report/viva-তে "engineering quality" হিসেবে আলাদা দাগ কাটবে। কিন্তু core Python কাজ করার **পরে**।

---

## Value-add #1 — Concurrent simulation orchestrator (best fit) ⭐
Generating 10⁴–10⁵ CAMB spectra is an **embarrassingly parallel batch job** — exactly what Go's concurrency model (goroutines, worker pools, channels) is built for.

**Design:**
- A Go **coordinator** samples the prior (or reads a parameter manifest), fans parameter sets out to a pool of **Python CAMB workers** (via subprocess / a thin local socket), collects results, and writes **sharded** output (`.npz`/Parquet).
- Adds the things a backend engineer naturally thinks of and a physicist often doesn't:
  - **Worker pool with bounded concurrency** (= #CPU cores) → max throughput without thrashing.
  - **Retries + resume** → a crash at sim 47,000 doesn't restart from zero.
  - **Dedup by parameter hash** → never simulate the same θ twice.
  - **Manifest / metadata store** → seed, prior ranges, git commit, per-shard checksums → full **reproducibility** (a Ch2 value, and a rubric-friendly one).

**Why it's legit:** CAMB still does the physics (Python/Fortran); Go just orchestrates — idiomatic and honest. This is a real **Track A (Simulator & Data)** contribution that makes the dataset faster, reliable, and reproducible.

```
   prior samples ──► [Go coordinator] ──► job channel ──► ┌─ CAMB worker (py) ─┐
                          ▲                                ├─ CAMB worker (py) ─┤──► results channel
                     manifest/checksums                   └─ CAMB worker (py) ─┘        │
                          ▼                                                              ▼
                    reproducible run record  ◄───────────────  sharded dataset (.npz/parquet)
```

> 💡 **বাংলায়:** ১০⁴–১০⁵ simulation বানানো একটা বড় parallel batch — Go-র goroutine/worker-pool এর জন্য আদর্শ। Go coordinator parameter ভাগ করে অনেক Python CAMB worker-এ পাঠায়, ফল জমা করে, retry/resume করে, parameter-hash দিয়ে dedup করে, আর একটা manifest (seed, range, checksum) লেখে → পুরো reproducible। physics CAMB-ই করে, Go শুধু orchestrate করে — একদম legit।

---

## Value-add #2 — Amortized-inference serving API (the "wow" demo) ⭐
This one *showcases the project's whole point*: "amortized" means inference is instant. Make that tangible with a service.

**Design:**
- A small **Go gRPC/REST gateway**: `POST` a `P(k)` spectrum → returns posterior samples / credible intervals as JSON.
- Behind it, either: (a) a thin Python inference worker holding the trained BayesFlow model, with Go handling the API/concurrency/validation/rate-limiting; or (b) export the trained flow to **ONNX** and run inference directly from Go (more ambitious).

**Why it's gold:** it turns "amortized inference" from a claim into a **live demo slide** — "send a spectrum, get a calibrated posterior in milliseconds." That viscerally demonstrates the advantage over MCMC (which would take minutes–hours per query). Plays directly to your 3 years of backend/API experience.

```
  client ──POST /infer {pk:[...]}──► [Go API gateway] ──► [Py inference worker / ONNX] ──► posterior samples (JSON)
                                       (validation,                (trained BayesFlow
                                        concurrency,                 model, forward pass)
                                        batching)
```

> 💡 **বাংলায়:** এটাই project-এর মূল সৌন্দর্য চোখে দেখানোর সুযোগ — amortized মানে inference তাৎক্ষণিক। একটা Go API বানাও: P(k) পাঠাও → milliseconds-এ posterior ফেরত। presentation-এ live demo slide হিসেবে দুর্দান্ত, আর MCMC-র (যা প্রতি query-তে মিনিট/ঘণ্টা নিত) সাথে পার্থক্যটা জলজ্যান্ত দেখায়। তোমার backend/API অভিজ্ঞতা এখানে সরাসরি কাজে লাগবে।

---

## Value-add #3 — Reproducibility & experiment infrastructure (lightweight)
Backend hygiene that quietly lifts the whole project and gives easy Reflection-section points:
- **Config-driven runs** (one YAML/TOML → prior ranges, budget, seeds, network sizes).
- **Deterministic seeding** threaded through the whole pipeline (data gen → training → eval).
- **Run registry / tiny status dashboard** (a Go HTTP service showing sim progress, dataset stats, training status).

> 💡 **বাংলায়:** ছোট কিন্তু কার্যকর — config-চালিত run, সব জায়গায় deterministic seed, আর একটা ছোট Go dashboard (progress/status)। reflection section-এ সহজ মার্ক, আর পুরো pipeline পরিষ্কার থাকে।

---

## What to skip
- ❌ Rewriting noise generation / preprocessing in Go — NumPy is already fast; marginal gain, not worth the moving parts.
- ❌ Any Go in the actual gradient training loop — that's Keras/JAX/PyTorch territory.

---

## How to present it for marks (without breaking the "minimal" rule)
- **Report:** a short "Implementation & reproducibility" paragraph + one line in **Reflection** ("we engineered a concurrent Go pipeline for reproducible, resumable simulation and a serving API demonstrating amortized inference"). Keep it out of the 10-page *science* budget — it's supporting, not central.
- **Presentation:** at most **one** slide (or a single sentence on the training slide) — the serving demo is the highest-impact way to spend it.
- **Viva:** if asked about scaling, reproducibility, or "what would you need for real surveys," you have a strong, differentiated answer most groups won't.
- **Always say:** the scientific result does **not** depend on Go; it's infrastructure and demonstration. This keeps you safely on the right side of "minimal working model."

---

## Recommended scope decision
| If you have… | Do |
|---|---|
| Limited time | **#1 orchestrator only** (most natural fit, real benefit, low risk) |
| Some spare time | **#1 + #2 serving demo** (the serving API is the best marks-per-hour) |
| Comfortable & ahead | add **#3** opportunistically |

**My recommendation:** get the **Python core fully working and calibrated first** (that's the grade). Then add the **Go simulation orchestrator (#1)** — it's the cleanest fit for your skills and genuinely improves the dataset. If you're ahead of schedule, the **serving API (#2)** is the single most impressive add for the presentation. Treat everything here as a bonus layer that demonstrates engineering maturity, never as a substitute for the SBI science.

> 💡 **বাংলায়:** আগে Python core কাজ করাও ও calibrate করো (এটাই মার্ক)। তারপর #1 (Go orchestrator) — তোমার skill-এর সাথে সবচেয়ে মানানসই ও সত্যিকারের উপকারী। সময় থাকলে #2 (serving API) — presentation-এ সবচেয়ে বেশি প্রভাব ফেলবে। সবই bonus layer, science-এর বিকল্প নয়।
