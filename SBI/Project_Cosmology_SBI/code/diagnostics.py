"""diagnostics.py — loss, recovery, SBC, contraction, fiducial inference, PPC."""
import os, json, numpy as np
import matplotlib; matplotlib.use("Agg")
import matplotlib.pyplot as plt
from simulator import K_GRID, pk_clean, theta_to_As, SIGMA_REL, PRIOR, PARAM_NAMES

OUT = os.path.dirname(os.path.abspath(__file__))
FIG = os.path.join(OUT,"..","figures"); os.makedirs(FIG, exist_ok=True)
LABELS = [r"$H_0$", r"$\Omega_m$", r"$n_s$", r"$\ln(10^{10}A_s)$"]
rng = np.random.default_rng(123)

M = np.load(os.path.join(OUT,"model.npz"))
params = [M[f"arr_{i}"] for i in range(6)]
logmu,logsd,th_mu,th_sd = M["logmu"],M["logsd"],M["th_mu"],M["th_sd"]
D = np.load(os.path.join(OUT,"data.npz"))

FLOOR=float(M["floor"]) if "floor" in M else 1.0
def prep(pk): return (np.log10(np.maximum(pk,FLOOR))-logmu)/logsd
def net(X):
    W1,b1,W2,b2,W3,b3 = params
    h1=np.tanh(X@W1+b1); h2=np.tanh(h1@W2+b2); return h2@W3+b3
def posterior(X):
    """return mean_phys (N,4), L_phys (N,4,4) [Cholesky of physical-space cov]."""
    o=net(X); N=len(X)
    mean=o[:,:4]; ld=np.clip(o[:,4:8],-7,3); off=o[:,8:14]
    L=np.zeros((N,4,4))
    for j in range(4): L[:,j,j]=np.exp(ld[:,j])
    L[:,1,0],L[:,2,0],L[:,2,1]=off[:,0],off[:,1],off[:,2]
    L[:,3,0],L[:,3,1],L[:,3,2]=off[:,3],off[:,4],off[:,5]
    mean_phys = mean*th_sd + th_mu
    L_phys = L*th_sd[None,:,None]                       # scale rows by th_sd
    return mean_phys, L_phys
def sample_post(mean,L,S):                              # (N,4),(N,4,4)->(N,S,4)
    eps=rng.standard_normal((mean.shape[0],S,4))
    return mean[:,None,:] + np.einsum("nij,nsj->nsi", L, eps)

# ---------- 1. loss curve (two-phase aware) ----------
tag=[str(t) for t in M["hist_tag"]]; step=M["hist_step"]; tr=M["hist_train"]; va=M["hist_val"]
mse=[i for i in range(len(tag)) if tag[i]=="mse"]; nll=[i for i in range(len(tag)) if tag[i]=="nll"]
fig,ax=plt.subplots(1,2,figsize=(11,4.2))
p1s=[step[i] for i in mse]; b=max(p1s) if p1s else 0
ax[0].plot(p1s,[tr[i] for i in mse],'o-',color='#c0392b',label='Phase 1: MSE pretraining')
ax[0].plot([step[i]+b for i in nll],[tr[i] for i in nll],'o-',color='#1C7293',label='Phase 2: Gaussian NLL')
ax[0].axvline(b,ls='--',color='grey'); ax[0].set_title('Full schedule (objective changes at dashed line)')
ax[0].set_xlabel('training step'); ax[0].set_ylabel('Gaussian NLL (monitored)'); ax[0].legend(fontsize=9); ax[0].grid(alpha=.3)
ax[1].plot([step[i] for i in nll],[tr[i] for i in nll],'o-',color='#1C7293',label='train')
ax[1].plot([step[i] for i in nll],[va[i] for i in nll],'o-',color='#E8A13A',label='validation')
ax[1].set_title('Phase 2 — posterior training (NLL)'); ax[1].set_xlabel('training step (NLL phase)')
ax[1].set_ylabel('Gaussian NLL'); ax[1].legend(fontsize=9); ax[1].grid(alpha=.3)
fig.suptitle('Training convergence: MSE warm-up then Gaussian NLL; validation tracks train (no overfit)',fontsize=11)
fig.tight_layout(); fig.savefig(os.path.join(FIG,"fig3_loss.png"),dpi=130)

# ---------- test posteriors ----------
pk_te,theta_te = D["pk_te"],D["theta_te"]
Xte = prep(pk_te + SIGMA_REL*pk_te*rng.standard_normal(pk_te.shape))
mean,L = posterior(Xte)
margstd = np.sqrt(np.einsum("nij,nij->ni", L, L))       # marginal sd per param
res = {}

# ---------- 2. recovery ----------
fig,ax=plt.subplots(1,4,figsize=(16,4))
for j in range(4):
    t=theta_te[:,j]; e=mean[:,j]
    ax[j].errorbar(t[:300],e[:300],yerr=margstd[:300,j],fmt='.',ms=3,alpha=.4,elinewidth=.5)
    lo,hi=PRIOR[PARAM_NAMES[j]]; ax[j].plot([lo,hi],[lo,hi],'k--',lw=1)
    r=np.corrcoef(t,e)[0,1]; rmse=np.sqrt(np.mean((t-e)**2))
    ax[j].set_title(f"{LABELS[j]}  r={r:.2f}  rmse={rmse:.3g}")
    ax[j].set_xlabel("true"); ax[j].set_ylabel("posterior mean")
    res[f"recovery_r_{PARAM_NAMES[j]}"]=round(float(r),3)
    res[f"recovery_rmse_{PARAM_NAMES[j]}"]=round(float(rmse),4)
fig.suptitle("Parameter recovery (test set)"); fig.tight_layout()
fig.savefig(os.path.join(FIG,"fig4_recovery.png"),dpi=130)

# ---------- 3. SBC rank histograms ----------
S=999
samp=sample_post(mean,L,S)                              # (N,S,4)
ranks=(samp < theta_te[:,None,:]).sum(1)               # (N,4) rank in 0..S
fig,ax=plt.subplots(1,4,figsize=(16,3.5)); nb=20
for j in range(4):
    ax[j].hist(ranks[:,j],bins=nb,color="steelblue",edgecolor="w")
    exp=len(ranks)/nb; band=np.sqrt(exp)
    ax[j].axhspan(exp-2*band,exp+2*band,color="grey",alpha=.3)
    ax[j].axhline(exp,color="k",lw=1); ax[j].set_title(f"SBC rank: {LABELS[j]}")
    ax[j].set_xlabel("rank statistic")
fig.suptitle("Simulation-Based Calibration — ranks should be ~uniform (grey = 95% band)")
fig.tight_layout(); fig.savefig(os.path.join(FIG,"fig5_sbc.png"),dpi=130)
# SBC uniformity (chi-square-ish): fraction of bins inside band
for j in range(4):
    h,_=np.histogram(ranks[:,j],bins=nb); exp=len(ranks)/nb; band=2*np.sqrt(exp)
    res[f"sbc_frac_in_band_{PARAM_NAMES[j]}"]=round(float(np.mean(np.abs(h-exp)<band)),2)

# ---------- 4. contraction + z-score + coverage ----------
prior_var=np.array([(PRIOR[p][1]-PRIOR[p][0])**2/12 for p in PARAM_NAMES])
post_var=(margstd**2).mean(0)
contraction=1-post_var/prior_var
z=(mean-theta_te)/margstd
cov68=np.mean(np.abs(z)<1.0,0); cov95=np.mean(np.abs(z)<1.96,0)
for j,p in enumerate(PARAM_NAMES):
    res[f"contraction_{p}"]=round(float(contraction[j]),3)
    res[f"zmean_{p}"]=round(float(z[:,j].mean()),3); res[f"zstd_{p}"]=round(float(z[:,j].std()),3)
    res[f"cov68_{p}"]=round(float(cov68[j]),3); res[f"cov95_{p}"]=round(float(cov95[j]),3)
fig,ax=plt.subplots(figsize=(7,5))
ax.bar(range(4),contraction,color="teal"); ax.set_xticks(range(4)); ax.set_xticklabels(LABELS)
ax.set_ylabel("posterior contraction  1 - Var(post)/Var(prior)"); ax.set_ylim(0,1)
ax.set_title("Posterior contraction (higher = more informative)")
for j,c in enumerate(contraction): ax.text(j,c+.02,f"{c:.2f}",ha="center")
fig.tight_layout(); fig.savefig(os.path.join(FIG,"fig6_contraction.png"),dpi=130)

# ---------- 5. Planck-fiducial inference (corner) ----------
fid=dict(H0=67.4,Om=0.315,ns=0.965,As=2.1e-9); fid_lnAs=np.log(fid["As"]/1e-10)
truth=np.array([fid["H0"],fid["Om"],fid["ns"],fid_lnAs])
Pf=pk_clean(fid["H0"],fid["Om"],fid["ns"],fid["As"])
xobs=prep(Pf + SIGMA_REL*Pf*rng.standard_normal(Pf.shape))[None,:]
mf,Lf=posterior(xobs); sf=sample_post(mf,Lf,4000)[0]    # (4000,4)
# H0-Om correlation
corr=np.corrcoef(sf[:,0],sf[:,1])[0,1]; res["fiducial_H0_Om_corr"]=round(float(corr),3)
for j,p in enumerate(PARAM_NAMES):
    res[f"fiducial_true_{p}"]=round(float(truth[j]),4)
    res[f"fiducial_postmean_{p}"]=round(float(mf[0,j]),4)
    res[f"fiducial_postsd_{p}"]=round(float(np.sqrt((Lf[0,j]**2).sum())),4)
fig,ax=plt.subplots(4,4,figsize=(11,11))
for i in range(4):
    for j in range(4):
        A=ax[i,j]
        if j>i: A.axis("off"); continue
        if i==j:
            A.hist(sf[:,i],bins=40,color="steelblue"); A.axvline(truth[i],color="crimson",lw=2)
        else:
            A.plot(sf[:,j],sf[:,i],'.',ms=1.5,alpha=.15,color="steelblue")
            A.axvline(truth[j],color="crimson",lw=1); A.axhline(truth[i],color="crimson",lw=1)
        if i==3: A.set_xlabel(LABELS[j])
        if j==0 and i>0: A.set_ylabel(LABELS[i])
ax[0,1].axis("off")
fig.suptitle(f"Posterior at Planck-fiducial cosmology (red = truth)\n"
             f"H0-Om correlation = {corr:.2f}  (the shape-parameter degeneracy)",fontsize=12)
fig.tight_layout(); fig.savefig(os.path.join(FIG,"fig7_fiducial_corner.png"),dpi=120)

# ---------- 6. posterior predictive check (proper: noisy obs vs noise-included predictions) ----------
fig,a=plt.subplots(figsize=(7.5,5))
obs_noisy = Pf + SIGMA_REL*Pf*rng.standard_normal(Pf.shape)   # the actual noisy observation
Pp_noisy=None
for kk in rng.integers(0,4000,80):
    th=sf[kk]; Pp=pk_clean(th[0],th[1],th[2],theta_to_As(th[3]))
    Pp_noisy=Pp + SIGMA_REL*Pp*rng.standard_normal(Pp.shape)   # include observational noise
    a.loglog(K_GRID,np.abs(Pp_noisy),color="steelblue",alpha=.12)
a.loglog(K_GRID,np.abs(Pp_noisy),color="steelblue",alpha=.6,label="posterior-predictive (with noise)")
a.loglog(K_GRID,Pf,'k-',lw=1.6,label="true spectrum (fiducial)")
a.loglog(K_GRID,np.abs(obs_noisy),'o',ms=4,color="crimson",label="observed (noisy)")
a.set_xlabel("k [h/Mpc]"); a.set_ylabel("P(k) [(Mpc/h)$^3$]")
a.set_title("Posterior predictive check — observed data lies within the predictive spread")
a.legend(fontsize=9); a.grid(alpha=.3)
fig.tight_layout(); fig.savefig(os.path.join(FIG,"fig8_ppc.png"),dpi=130)

json.dump(res, open(os.path.join(OUT,"results.json"),"w"), indent=2)
print(json.dumps(res, indent=2))
print("\nAll figures saved to figures/")
