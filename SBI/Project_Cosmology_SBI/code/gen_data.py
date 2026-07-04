"""gen_data.py — generate train/val/test datasets + a sanity figure."""
import numpy as np, os
import matplotlib; matplotlib.use("Agg")
import matplotlib.pyplot as plt
from simulator import (K_GRID, sample_prior, simulate_clean, pk_clean, theta_to_As,
                       PARAM_NAMES, PRIOR, SIGMA_REL)

OUT = os.path.dirname(os.path.abspath(__file__))
FIG = os.path.join(OUT, "..", "figures"); os.makedirs(FIG, exist_ok=True)
rng = np.random.default_rng(42)

N_TRAIN, N_VAL, N_TEST = 12000, 2000, 3000
theta_tr = sample_prior(N_TRAIN, rng); pk_tr = simulate_clean(theta_tr)
theta_va = sample_prior(N_VAL,  rng);  pk_va = simulate_clean(theta_va)
theta_te = sample_prior(N_TEST, rng);  pk_te = simulate_clean(theta_te)

# standardization stats from TRAIN (log10 P)
logmu = np.log10(pk_tr).mean(0); logsd = np.log10(pk_tr).std(0)
# parameter standardization (from prior ranges -> use train mean/std)
th_mu = theta_tr.mean(0); th_sd = theta_tr.std(0)

np.savez_compressed(os.path.join(OUT, "data.npz"),
    theta_tr=theta_tr, pk_tr=pk_tr, theta_va=theta_va, pk_va=pk_va,
    theta_te=theta_te, pk_te=pk_te, logmu=logmu, logsd=logsd,
    th_mu=th_mu, th_sd=th_sd, k=K_GRID, sigma_rel=SIGMA_REL)
print("saved data.npz  train/val/test =", N_TRAIN, N_VAL, N_TEST)

# ---- sanity figure: effect of each parameter on P(k) ----
fid = dict(H0=67.4, Om=0.315, ns=0.965, As=2.1e-9)
fig, ax = plt.subplots(2, 2, figsize=(11, 8))
specs = [("H0",[60,67.4,75],"km/s/Mpc"), ("Om",[0.20,0.315,0.45],""),
         ("ns",[0.92,0.965,1.02],""), ("As",[1.2e-9,2.1e-9,3.0e-9],"")]
for a,(name,vals,unit) in zip(ax.ravel(), specs):
    for v in vals:
        p = dict(fid); p[name]=v
        P = pk_clean(p["H0"],p["Om"],p["ns"],p["As"])
        lab = f"{name}={v:.3g}" if name!="As" else f"As={v:.1e}"
        a.loglog(K_GRID, P, label=lab)
    a.set_title(f"Effect of {name}"); a.set_xlabel("k  [h/Mpc]")
    a.set_ylabel("P(k)  [(Mpc/h)$^3$]"); a.legend(fontsize=8); a.grid(alpha=.3)
fig.suptitle("Forward simulator: how each parameter shapes the linear P(k)", fontsize=13)
fig.tight_layout(); fig.savefig(os.path.join(FIG,"fig1_simulator_effects.png"), dpi=130)
print("saved fig1_simulator_effects.png")

# ---- example noisy observation ----
fig2, a = plt.subplots(figsize=(7,5))
P = pk_clean(**{k:fid[k] for k in ["H0","Om","ns","As"]})
noisy = P + SIGMA_REL*P*rng.standard_normal(P.shape)
a.loglog(K_GRID, P, 'k-', lw=2, label="clean P(k) (Planck fiducial)")
a.loglog(K_GRID, np.abs(noisy), '.', ms=4, color="crimson", label="noisy observation")
a.set_xlabel("k  [h/Mpc]"); a.set_ylabel("P(k)  [(Mpc/h)$^3$]")
a.set_title("One simulated observation (cosmic-variance noise)"); a.legend(); a.grid(alpha=.3)
fig2.tight_layout(); fig2.savefig(os.path.join(FIG,"fig2_observation.png"), dpi=130)
print("saved fig2_observation.png")
