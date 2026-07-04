# Live Inference Demo — Go backend + web frontend

An interactive demo of the project: **drag the knobs of the Universe → see its noisy
power-spectrum fingerprint → watch the trained neural posterior estimator recover the
parameters with calibrated uncertainty**, including the live H₀–Ωₘ degeneracy.

The Go server does **everything in pure standard library** — it loads the trained
network weights from `model.json` and performs *both* the forward simulation
(θ → P(k)) and the inverse amortized inference (P(k) → posterior). **No Python at
runtime.** This is the engineering value-add described in `../05_Engineering_Value_Add_Golang.md`
(value-add #2), and a memorable presentation demo of "amortized inference = one instant forward pass."

## Files
- `main.go` — HTTP server: forward simulator (BBKS) + NPE forward pass + Gaussian posterior + sampling.
- `index.html` — self-contained frontend (sliders + canvas plots, no external libraries).
- `model.json` — exported trained weights + standardization + simulator constants (regenerate from `../code/` if you retrain).
- `go.mod` — module file (Go ≥ 1.21).

## Run
```bash
# install Go once: https://go.dev/dl/   (or: sudo apt install golang)
cd webdemo
go run .
# open http://localhost:8081
```

## What you see
1. **Set the true Universe** — 4 sliders (H₀, Ωₘ, nₛ, ln 10¹⁰Aₛ) + a Planck-2018 preset.
2. **P(k) plot** — the noisy spectrum the network receives (red) vs the clean truth (white).
3. **Recovered vs true** — per-parameter bars: gold = truth, teal = posterior mean ± 1σ. Watch Ωₘ/nₛ/Aₛ land tight on truth while **H₀ stays wide**.
4. **H₀–Ωₘ scatter** — the live "banana" with its correlation; the degeneracy you can *see*.

> Note: the uniform prior makes parameter values outside its box physically impossible, so the
> scatter shows the Gaussian posterior **truncated to the prior support** (samples are rejection-sampled
> to stay in-box). The reported correlation is computed analytically from the posterior covariance, so it
> matches the report's value exactly and is unaffected by the display-side truncation.

## How it maps to the project
- Forward simulator = the `theta → P(k)` model (Chapters 1–3).
- The network = the amortized neural posterior estimator (Chapters 11–12 / BayesFlow analogue).
- The wide H₀ + banana = the shape-parameter degeneracy (the headline science result).

## Regenerating model.json (after retraining)
```bash
cd ../code
python3 -c "import json,numpy as np,simulator as S; S.pk_clean(67.4,.315,.965,2.1e-9); \
M=np.load('model.npz'); D=np.load('data.npz'); \
json.dump(dict(W1=M['arr_0'].tolist(),b1=M['arr_1'].tolist(),W2=M['arr_2'].tolist(),b2=M['arr_3'].tolist(),\
W3=M['arr_4'].tolist(),b3=M['arr_5'].tolist(),logmu=M['logmu'].tolist(),logsd=M['logsd'].tolist(),\
th_mu=M['th_mu'].tolist(),th_sd=M['th_sd'].tolist(),floor=float(M['floor']),k=D['k'].tolist(),\
sigma_rel=D['sigma_rel'].tolist(),NORM=float(S._NORM),AS0=float(S.AS0),OMEGA_B_H2=float(S.OMEGA_B_H2),\
prior=dict(H0=[60,80],Om=[.1,.5],ns=[.9,1.05],lnAs=[2.5,3.5])), open('../webdemo/model.json','w'))"
```

> Note: this is the **prototype** network (Gaussian NPE + BBKS simulator). When you port the
> core to CAMB + BayesFlow, re-export the weights (or expose a `/run` that calls the BayesFlow
> model) and the same frontend works unchanged.
