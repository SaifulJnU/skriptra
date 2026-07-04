"""
npe.py — amortized Gaussian Neural Posterior Estimator (full covariance).

MLP maps standardized log P(k) -> mean + Cholesky factor of a 4-D Gaussian
posterior over (H0, Om, ns, lnAs). Trained in two phases:
  Phase 1: MSE on the mean (avoids the variance-inflation trap),
  Phase 2: full Gaussian NLL (learns calibrated covariance, incl. correlations).
Inputs standardized using NOISY training statistics so cosmic-variance-dominated
low-k bins are down-weighted rather than amplified.
Backend: autograd + Adam. CPU-friendly.
"""
import os, time, numpy as np
import autograd.numpy as anp
from autograd import grad

OUT = os.path.dirname(os.path.abspath(__file__))
D = np.load(os.path.join(OUT, "data.npz"))
th_mu, th_sd = D["th_mu"], D["th_sd"]
sigma_rel = D["sigma_rel"]
pk_tr, theta_tr = D["pk_tr"], D["theta_tr"]
pk_va, theta_va = D["pk_va"], D["theta_va"]
NK = pk_tr.shape[1]; P = 4
rng = np.random.default_rng(7)
FLOOR = 1.0

def make_noisy(pk): return pk + sigma_rel*pk*rng.standard_normal(pk.shape)
# standardization from NOISY training data (robust to noisy low-k bins)
_xn = np.log10(np.maximum(make_noisy(pk_tr), FLOOR))
logmu, logsd = _xn.mean(0), _xn.std(0)

def preprocess_x(pk): return (np.log10(np.maximum(pk, FLOOR)) - logmu) / logsd
def standardize_y(theta): return (theta - th_mu) / th_sd

Xva = preprocess_x(make_noisy(pk_va)); Yva = standardize_y(theta_va)

H = 128
def init_params(seed=0):
    r = np.random.default_rng(seed)
    def xav(nin, nout): return r.standard_normal((nin, nout))*np.sqrt(1.0/nin)
    p = [xav(NK,H), np.zeros(H), xav(H,H), np.zeros(H), xav(H,14), np.zeros(14)]
    p[5][4:8] = np.log(0.6)            # init posterior sd ~0.6 (std-space)
    return p

def net(params, X):
    W1,b1,W2,b2,W3,b3 = params
    h1 = anp.tanh(anp.dot(X,W1)+b1); h2 = anp.tanh(anp.dot(h1,W2)+b2)
    return anp.dot(h2,W3)+b3

def mse(params, X, Y):
    return anp.mean((net(params,X)[:, :4] - Y)**2)

def nll(params, X, Y):
    o = net(params, X)
    mean = o[:, :4]; ld = anp.clip(o[:, 4:8], -7.0, 3.0); off = o[:, 8:14]
    d = Y - mean
    L00,L11,L22,L33 = [anp.exp(ld[:,i]) for i in range(4)]
    l10,l20,l21,l30,l31,l32 = [off[:,i] for i in range(6)]
    z0 = d[:,0]/L00
    z1 = (d[:,1]-l10*z0)/L11
    z2 = (d[:,2]-l20*z0-l21*z1)/L22
    z3 = (d[:,3]-l30*z0-l31*z1-l32*z2)/L33
    quad = z0**2+z1**2+z2**2+z3**2
    logdet = 2.0*(ld[:,0]+ld[:,1]+ld[:,2]+ld[:,3])
    return anp.mean(0.5*quad + 0.5*logdet) + 0.5*P*anp.log(2*anp.pi)

dmse, dnll = grad(mse), grad(nll)

def adam(params, dfun, steps, lr, hist, tag):
    m=[np.zeros_like(p) for p in params]; v=[np.zeros_like(p) for p in params]
    b1_,b2_,eps=0.9,0.999,1e-8; t0=time.time()
    for s in range(1,steps+1):
        idx=rng.integers(0,len(pk_tr),256); pk=pk_tr[idx]
        xb=preprocess_x(make_noisy(pk)); yb=standardize_y(theta_tr[idx])
        g=dfun(params,xb,yb); lr_s=lr*(0.5**(s/ (steps*0.6)))
        for i in range(len(params)):
            m[i]=b1_*m[i]+(1-b1_)*g[i]; v[i]=b2_*v[i]+(1-b2_)*g[i]**2
            params[i]-=lr_s*(m[i]/(1-b1_**s))/(np.sqrt(v[i]/(1-b2_**s))+eps)
        if s%400==0 or s==1:
            tr=float(nll(params,xb,yb)); va=float(nll(params,Xva,Yva))
            hist["step"].append((tag,s)); hist["train"].append(tr); hist["val"].append(va)
            print(f"[{tag}] step {s:5d}  NLL train {tr:7.3f}  val {va:7.3f} ({time.time()-t0:4.1f}s)",flush=True)

def train():
    params=init_params(0); hist={"step":[],"train":[],"val":[]}
    print("Phase 1: MSE mean pretraining"); adam(params,dmse,2500,2e-3,hist,"mse")
    print("Phase 2: full Gaussian NLL");    adam(params,dnll,7000,1e-3,hist,"nll")
    steps=[s for _,s in hist["step"]]
    np.savez(os.path.join(OUT,"model.npz"), *params, logmu=logmu,logsd=logsd,
             th_mu=th_mu,th_sd=th_sd,floor=FLOOR,
             hist_step=steps,hist_train=hist["train"],hist_val=hist["val"],
             hist_tag=[t for t,_ in hist["step"]])
    print("DONE saved model.npz",flush=True)

if __name__=="__main__": train()
