# Chapter 0b: Practice True/False Questions

**Topic:** Notation & Definitions (Bischl, Moosbauer, Groll — `chapter_0b_Notation_slides_SoSe2026.pdf`)

**Instructions:** Answer each statement TRUE or FALSE, then check against the answers below.  
Quiz format: 30 statements in 30 minutes — aim to answer each in under 60 seconds.

---

## Questions

**Q1.** Calligraphic 𝒳 and 𝒴 denote the input space and the output/target space, while lowercase x and y denote individual values within those spaces.

**Q2.** The feature vector **x** is written in bold because it collects several feature values, e.g. **x** = (x₁, …, x_p)ᵀ ∈ 𝒳.

**Q3.** The symbols "target", "label", "output", and "response" all refer to y.

**Q4.** ℙ_xy, the joint probability distribution over 𝒳 × 𝒴, is typically known to the modeller.

**Q5.** Blackboard ℙ denotes a probability distribution, whereas lowercase p denotes a density/pdf.

**Q6.** In p(x | θ) the bar "|" indicates Bayesian conditioning on θ as a random variable.

**Q7.** Because this course is frequentist, p(x | θ) can equivalently be written p_θ(x) or p(x ; θ).

**Q8.** In x⁽ⁱ⁾ the superscript (i) indicates which feature is being referenced.

**Q9.** The dataset is written 𝒟 = {(x⁽¹⁾, y⁽¹⁾), …, (x⁽ⁿ⁾, y⁽ⁿ⁾)} and contains n observations.

**Q10.** 𝒟_train and 𝒟_test are usually a disjoint split of 𝒟 (no overlapping observations).

**Q11.** f(x) returns a discrete class label directly, while h(x) returns a real-valued score.

**Q12.** The parameter space Θ is the set of all admissible parameter vectors θ.

**Q13.** The hypothesis space ℋ is the set of all admissible models f.

**Q14.** The residual in regression is defined as ε = f(x) − y.

**Q15.** For binary classification with 𝒴 = {−1, 1}, the margin is y · f(x), and a positive margin means a correct classification.

**Q16.** A margin of −2 indicates a confident, correct prediction.

**Q17.** The posterior probability π_k(x) = ℙ(y = k | x) is the class-k probability after seeing x.

**Q18.** The prior probability π_k = ℙ(y = k) depends on the specific feature vector x.

**Q19.** ℒ(θ) denotes the likelihood and ℓ(θ) denotes the log-likelihood, with ℓ = log ℒ.

**Q20.** The hat in θ̂ indicates the true, unknown parameter value.

**Q21.** θ denotes the true (unknown) parameter, while θ̂ denotes the estimate learned from data.

**Q22.** When an intercept is included, the design matrix **X** has dimensions n × (p + 1) and its first column is all ones.

**Q23.** In the subscript/superscript convention, x_j (subscript) refers to the j-th feature (a column) across all observations.

**Q24.** The entry x_j⁽ⁱ⁾ is the j-th feature of the i-th observation (a single cell).

**Q25.** The intercept trick lets us write f(x) = θᵀx instead of f(x) = θ₀ + θᵀx by appending a constant-1 feature.

**Q26.** Under the coding 𝒴 = {0, 1}, the decision rule is h(x) = 𝟙(π(x) ≥ 0.5).

**Q27.** The indicator function 𝟙(·) returns 1 when its condition is true and 0 otherwise.

**Q28.** Under the coding 𝒴 = {−1, +1}, the decision rule is h(x) = sign(f(x)), and |f(x)| expresses the confidence.

**Q29.** The {0, 1} coding is the probability-based view (π, threshold 0.5), while the {−1, +1} coding is the geometric/score-based view (f, sign + confidence).

**Q30.** The superscript ᵀ denotes the transpose, which swaps rows and columns.

---

## Answers

| Q | Answer | Key Reason |
|---|--------|------------|
| 1 | **TRUE** | Calligraphic = space, lowercase = a single value |
| 2 | **TRUE** | Bold x = feature vector (x₁,…,x_p)ᵀ |
| 3 | **TRUE** | All four are synonyms for y |
| 4 | **FALSE** | ℙ_xy is generally **unknown** |
| 5 | **TRUE** | ℙ = distribution, p = density/pdf |
| 6 | **FALSE** | Frequentist notation only; **not** Bayesian conditioning |
| 7 | **TRUE** | "|" here means "determined by θ" → p_θ(x) / p(x;θ) |
| 8 | **FALSE** | Superscript (i) = which **observation** (row), not feature |
| 9 | **TRUE** | Definition of 𝒟 with n observations |
| 10 | **TRUE** | Usually a disjoint union (no overlap) |
| 11 | **FALSE** | Reversed: f = score (real), h = discrete class |
| 12 | **TRUE** | Θ = set of all possible θ |
| 13 | **TRUE** | ℋ = set of all admissible models f |
| 14 | **FALSE** | Residual is ε = y − f(x), not f(x) − y |
| 15 | **TRUE** | margin = y·f(x); positive → correct |
| 16 | **FALSE** | Negative margin → **incorrect** classification |
| 17 | **TRUE** | Posterior uses x (after evidence) |
| 18 | **FALSE** | Prior does **not** depend on x (before evidence) |
| 19 | **TRUE** | ℒ = likelihood, ℓ = log-likelihood = log ℒ |
| 20 | **FALSE** | Hat = learned/estimated, not the true value |
| 21 | **TRUE** | θ = true unknown; θ̂ = data-based estimate |
| 22 | **TRUE** | Intercept version is n × (p+1), first column all 1s |
| 23 | **TRUE** | Subscript j = j-th feature (column) |
| 24 | **TRUE** | x_j⁽ⁱ⁾ = i-th observation's j-th feature |
| 25 | **TRUE** | Append constant-1 feature → f(x) = θᵀx |
| 26 | **TRUE** | {0,1}: h(x) = 𝟙(π(x) ≥ 0.5) |
| 27 | **TRUE** | Definition of the indicator function |
| 28 | **TRUE** | {−1,+1}: h(x) = sign(f(x)), confidence = |f(x)| |
| 29 | **TRUE** | {0,1} → probability; {−1,+1} → score |
| 30 | **TRUE** | ᵀ = transpose (rows ↔ columns) |

---

## Score Interpretation

| Score | Meaning |
|-------|---------|
| 27–30 | Excellent — Chapter 0b notation mastered |
| 22–26 | Good — review the FALSE answers carefully |
| 17–21 | Needs work — re-read the symbol dictionary and traps |
| < 17  | Re-study the full chapter summary before re-attempting |

---

> **Top traps to remember:** (1) superscript (i) = observation, subscript j = feature; (2) "|" here is frequentist readability, **not** Bayesian conditioning; (3) residual ε = y − f(x); (4) f = score, h = discrete class; (5) θ̂ = estimate, not truth; (6) {0,1} → probability (π, 0.5), {−1,+1} → score (sign + confidence).
