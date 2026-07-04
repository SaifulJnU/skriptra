"""
simulator.py — forward model for the SBI cosmology project.

Generates the LINEAR matter power spectrum P(k) at z=0 from cosmological
parameters theta = {H0, Omega_m, n_s, A_s} using the BBKS (Bardeen, Bond,
Kaiser & Szalay 1986) transfer function with the Sugiyama (1995) shape
parameter Gamma. This is a standard *analytic* approximation to the linear
P(k) used as a fast, dependency-free stand-in for CAMB in this self-contained
run. (In the production project the same interface is backed by CAMB.)

Units: k in h/Mpc, P(k) in (Mpc/h)^3.
The shape is controlled by Gamma ~ Omega_m * h, so the linear spectrum is
mainly sensitive to the COMBINATION Omega_m * h  ->  H0 and Omega_m are
degenerate along constant Gamma (the classic shape-parameter degeneracy).
"""
import numpy as np

# ---- fixed grid (100 log-spaced wavenumbers in [1e-3, 1] h/Mpc) ----
N_K = 100
K_GRID = np.logspace(-3.0, 0.0, N_K)        # h/Mpc
DK = np.gradient(K_GRID)

# ---- fixed / derived constants ----
OMEGA_B_H2 = 0.0224                          # fixed physical baryon density
AS0 = 2.1e-9                                 # reference amplitude (Planck-ish)
_NORM = None                                 # set on first call for nice magnitudes

def _transfer_bbks(k, Om, Ob, h):
    """BBKS transfer function with Sugiyama (1995) Gamma. k in h/Mpc."""
    Gamma = Om * h * np.exp(-Ob * (1.0 + np.sqrt(2.0 * h) / Om))   # shape parameter
    q = k / np.maximum(Gamma, 1e-8)
    q = np.maximum(q, 1e-8)
    L = np.log(1.0 + 2.34 * q) / (2.34 * q)
    C = 1.0 + 3.89 * q + (16.1 * q) ** 2 + (5.46 * q) ** 3 + (6.71 * q) ** 4
    return L * C ** (-0.25)

def pk_clean(H0, Om, ns, As, k=K_GRID):
    """Clean linear P(k) in (Mpc/h)^3. Vectorized over scalar params."""
    global _NORM
    h = H0 / 100.0
    Ob = OMEGA_B_H2 / h ** 2                  # Omega_b = omega_b / h^2 (fixed phys. density)
    T = _transfer_bbks(k, Om, Ob, h)
    P = (As / AS0) * (k ** ns) * T ** 2
    if _NORM is None:
        # normalize so a fiducial spectrum peaks at ~2.5e4 (Mpc/h)^3 (realistic)
        Tf = _transfer_bbks(k, 0.315, OMEGA_B_H2 / 0.674**2, 0.674)
        Pf = (k ** 0.965) * Tf ** 2
        _NORM = 2.5e4 / Pf.max()
    return _NORM * P

# ---- priors (explicit, proper) ----
PRIOR = {
    "H0":   (60.0, 80.0),     # km/s/Mpc
    "Om":   (0.10, 0.50),     # total matter fraction
    "ns":   (0.90, 1.05),     # spectral index
    "lnAs": (2.5, 3.5),       # ln(1e10 * A_s)  -> A_s in ~[1.2e-9, 3.3e-9]
}
PARAM_NAMES = ["H0", "Om", "ns", "lnAs"]

def sample_prior(n, rng):
    H0   = rng.uniform(*PRIOR["H0"],   n)
    Om   = rng.uniform(*PRIOR["Om"],   n)
    ns   = rng.uniform(*PRIOR["ns"],   n)
    lnAs = rng.uniform(*PRIOR["lnAs"], n)
    return np.stack([H0, Om, ns, lnAs], axis=-1)

def theta_to_As(lnAs):
    return 1e-10 * np.exp(lnAs)

# ---- cosmic-variance noise ----
V_SURVEY = 1e9                                # (Mpc/h)^3
N_MODES = K_GRID ** 2 * DK * V_SURVEY / (2.0 * np.pi ** 2)
SIGMA_REL = np.sqrt(2.0 / np.maximum(N_MODES, 1e-3))     # relative error per k-bin
SIGMA_REL = np.minimum(SIGMA_REL, 1.0)        # cap the few cosmic-variance-dominated low-k bins

def add_noise(pk, rng):
    return pk + (SIGMA_REL * pk) * rng.standard_normal(pk.shape)

def simulate_clean(theta):
    """theta: (N,4) -> clean P(k): (N, N_K)."""
    out = np.zeros((len(theta), N_K))
    for i, (H0, Om, ns, lnAs) in enumerate(theta):
        out[i] = pk_clean(H0, Om, ns, theta_to_As(lnAs))
    return out

if __name__ == "__main__":
    rng = np.random.default_rng(0)
    th = sample_prior(5, rng)
    P = simulate_clean(th)
    print("k range:", K_GRID[0], K_GRID[-1])
    print("P(k) sample peak values:", P.max(axis=1).round(1))
    print("noise rel range:", SIGMA_REL.min().round(4), SIGMA_REL.max().round(4))
