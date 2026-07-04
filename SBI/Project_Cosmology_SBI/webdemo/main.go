// SBI Cosmology — live inference demo backend (pure Go stdlib, no Python at runtime).
//
// Loads the trained neural-posterior-estimator weights from model.json and serves
// two things in one forward pass: (1) the FORWARD simulator theta -> noisy P(k),
// and (2) the INVERSE amortized posterior P(k) -> p(theta | P(k)).
//
//	go run .         # then open http://localhost:8081
package main

import (
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Model struct {
	W1 [][]float64 `json:"W1"`
	B1 []float64   `json:"b1"`
	W2 [][]float64 `json:"W2"`
	B2 []float64   `json:"b2"`
	W3 [][]float64 `json:"W3"`
	B3 []float64   `json:"b3"`

	Logmu []float64 `json:"logmu"`
	Logsd []float64 `json:"logsd"`
	ThMu  []float64 `json:"th_mu"`
	ThSd  []float64 `json:"th_sd"`
	Floor float64   `json:"floor"`

	K        []float64 `json:"k"`
	SigmaRel []float64 `json:"sigma_rel"`
	NORM     float64   `json:"NORM"`
	AS0      float64   `json:"AS0"`
	OmegaBH2 float64   `json:"OMEGA_B_H2"`

	Prior map[string][]float64 `json:"prior"`
}

var M Model

// ---------- forward simulator (BBKS + Sugiyama, identical to simulator.py) ----------
func transfer(k, Om, Ob, h float64) float64 {
	Gamma := Om * h * math.Exp(-Ob*(1.0+math.Sqrt(2.0*h)/Om))
	if Gamma < 1e-8 {
		Gamma = 1e-8
	}
	q := k / Gamma
	if q < 1e-8 {
		q = 1e-8
	}
	L := math.Log(1.0+2.34*q) / (2.34 * q)
	C := 1.0 + 3.89*q + math.Pow(16.1*q, 2) + math.Pow(5.46*q, 3) + math.Pow(6.71*q, 4)
	return L * math.Pow(C, -0.25)
}

func pkClean(H0, Om, ns, As float64) []float64 {
	h := H0 / 100.0
	Ob := M.OmegaBH2 / (h * h)
	out := make([]float64, len(M.K))
	for i, k := range M.K {
		T := transfer(k, Om, Ob, h)
		out[i] = (As / M.AS0) * math.Pow(k, ns) * T * T * M.NORM
	}
	return out
}

func addNoise(pk []float64, rng *rand.Rand) []float64 {
	out := make([]float64, len(pk))
	for i := range pk {
		out[i] = pk[i] + M.SigmaRel[i]*pk[i]*rng.NormFloat64()
	}
	return out
}

// ---------- neural posterior estimator (forward pass + Gaussian posterior) ----------
func vecMat(x []float64, W [][]float64, b []float64) []float64 {
	out := make([]float64, len(b))
	for j := range b {
		s := b[j]
		for i := range x {
			s += x[i] * W[i][j]
		}
		out[j] = s
	}
	return out
}
func tanhv(x []float64) []float64 {
	for i := range x {
		x[i] = math.Tanh(x[i])
	}
	return x
}

func preprocess(pk []float64) []float64 {
	x := make([]float64, len(pk))
	for i := range pk {
		v := pk[i]
		if v < M.Floor {
			v = M.Floor
		}
		x[i] = (math.Log10(v) - M.Logmu[i]) / M.Logsd[i]
	}
	return x
}

// returns posterior mean (physical), per-param std, and physical Cholesky factor
func posterior(x []float64) (mean, std []float64, Lp [4][4]float64) {
	o := vecMat(tanhv(vecMat(tanhv(vecMat(x, M.W1, M.B1)), M.W2, M.B2)), M.W3, M.B3)
	ld := make([]float64, 4)
	for i := 0; i < 4; i++ {
		ld[i] = math.Max(-7, math.Min(3, o[4+i]))
	}
	off := o[8:14]
	var Ls [4][4]float64
	for i := 0; i < 4; i++ {
		Ls[i][i] = math.Exp(ld[i])
	}
	Ls[1][0], Ls[2][0], Ls[2][1] = off[0], off[1], off[2]
	Ls[3][0], Ls[3][1], Ls[3][2] = off[3], off[4], off[5]
	mean = make([]float64, 4)
	std = make([]float64, 4)
	for i := 0; i < 4; i++ {
		mean[i] = o[i]*M.ThSd[i] + M.ThMu[i]
		var s float64
		for j := 0; j <= i; j++ {
			Lp[i][j] = Ls[i][j] * M.ThSd[i]
			s += Lp[i][j] * Lp[i][j]
		}
		std[i] = math.Sqrt(s)
	}
	return
}

// Posterior samples of (H0, Om). The uniform prior makes values outside its
// support impossible, so the correct posterior is the Gaussian TRUNCATED to the
// prior box; we therefore reject out-of-bounds draws (rejection sampling).
func sampleHO(mean []float64, Lp [4][4]float64, n int, rng *rand.Rand) [][2]float64 {
	h0lo, h0hi := M.Prior["H0"][0], M.Prior["H0"][1]
	omlo, omhi := M.Prior["Om"][0], M.Prior["Om"][1]
	out := make([][2]float64, 0, n)
	for tries := 0; len(out) < n && tries < n*40; tries++ {
		var eps [4]float64
		for j := 0; j < 4; j++ {
			eps[j] = rng.NormFloat64()
		}
		var v [4]float64
		for i := 0; i < 4; i++ {
			v[i] = mean[i]
			for j := 0; j <= i; j++ {
				v[i] += Lp[i][j] * eps[j]
			}
		}
		if v[0] < h0lo || v[0] > h0hi || v[1] < omlo || v[1] > omhi {
			continue // outside prior support → physically impossible, reject
		}
		out = append(out, [2]float64{v[0], v[1]})
	}
	return out
}

// ---------- HTTP ----------
type ParamOut struct {
	Name string  `json:"name"`
	True float64 `json:"true"`
	Mean float64 `json:"mean"`
	Std  float64 `json:"std"`
	Min  float64 `json:"min"`
	Max  float64 `json:"max"`
}
type RunReq struct{ H0, Om, Ns, LnAs float64 }
type RunResp struct {
	K        []float64    `json:"k"`
	Spectrum []float64    `json:"spectrum"`
	Clean    []float64    `json:"clean"`
	Params   []ParamOut   `json:"params"`
	Samples  [][2]float64 `json:"samples"`
	Corr     float64      `json:"corr"`
}

// Exact H0-Om posterior correlation from the Cholesky factor (covariance = Lp Lp^T).
// For a lower-triangular Lp: Cov00=Lp00^2, Cov11=Lp10^2+Lp11^2, Cov01=Lp00*Lp10,
// so corr = Lp10 / sqrt(Lp10^2 + Lp11^2). This matches the report's value and is
// independent of display-side sample truncation.
func corrHO(Lp [4][4]float64) float64 {
	return Lp[1][0] / math.Sqrt(Lp[1][0]*Lp[1][0]+Lp[1][1]*Lp[1][1]+1e-12)
}

func runHandler(w http.ResponseWriter, r *http.Request) {
	var req RunReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	As := 1e-10 * math.Exp(req.LnAs)
	clean := pkClean(req.H0, req.Om, req.Ns, As)
	noisy := addNoise(clean, rng)
	mean, std, Lp := posterior(preprocess(noisy))
	samples := sampleHO(mean, Lp, 500, rng)
	names := []string{"H0", "Om", "ns", "lnAs"}
	truth := []float64{req.H0, req.Om, req.Ns, req.LnAs}
	keys := []string{"H0", "Om", "ns", "lnAs"}
	params := make([]ParamOut, 4)
	for i := 0; i < 4; i++ {
		pr := M.Prior[keys[i]]
		params[i] = ParamOut{names[i], truth[i], mean[i], std[i], pr[0], pr[1]}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RunResp{M.K, noisy, clean, params, samples, corrHO(Lp)})
}

func main() {
	raw, err := os.ReadFile("model.json")
	if err != nil {
		log.Fatal("model.json not found: ", err)
	}
	if err := json.Unmarshal(raw, &M); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/run", runHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	log.Println("SBI cosmology demo running →  http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
