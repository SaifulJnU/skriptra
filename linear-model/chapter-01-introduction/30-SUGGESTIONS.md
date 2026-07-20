# Ch 1 — Suggestions: how to study this chapter

**Time budget: 60–90 minutes total. Once. Never again except the summary.**

Chapter 1 is worth roughly 1–2 marks on the exam, and those marks are always *disguised* — they show up as "choose an appropriate response variable" or "why is this a regression problem" inside a bigger question. Nobody has ever been asked "describe the Galton data" on this exam.

## What to actually do

1. **Read `01-notes-1.1-examples-of-applications.md` fast.** You are collecting vocabulary, not knowledge. Response, covariate, systematic component, random error, continuous vs binary vs count response.
2. **Read `02-notes-1.2-first-steps.md` properly.** This is the only part with exam value: scatter plots, what correlation does and doesn't tell you, and the idea that you *look at the data before you model it*. Diagnostics in Chapter 3.4 are the same idea wearing a suit.
3. **Read `03-notes-1.3-notation.md` twice, and write the notation table out by hand.** This is the highest-value 15 minutes in the entire chapter. Every formula in Chapter 3 is built from these symbols. If `x_i` vs `x_j` vs `x_ij` ever confuses you mid-exam, you lose minutes you don't have.
4. **Do `20-EXERCISES.md`** (15 min), check `21-SOLUTIONS.md`.
5. **Read `52-STORY-FOR-A-BABY.md`.** Then close everything and tell the story out loud.

## What to skip

- ❌ The specific datasets (Munich rent, credit scoring, malnutrition in Zambia). You need to *recognise* that regression applies to rent, credit default, and child growth — you do not need any numbers.
- ❌ The Galton historical detail. Know one sentence: *Galton invented "regression" describing children's heights regressing toward the mean.* That's the entire examinable content.
- ❌ Anything about the R packages or the book's website.

## The one thing to carry forward

Chapter 1's real job is to plant this sentence in your head:

> **observed = systematic + random**
> **y = f(x) + ε**

Everything for the next 100 pages is: *what shape is f, and what do we assume about ε?* When you hit a wall later, come back to that line.

## Self-check before moving on

Can you, without notes:

- [ ] Say what a response variable and a covariate are, and give an example of each?
- [ ] Say why we need the error term ε at all?
- [ ] Distinguish continuous / binary / count responses and name a model for each?
- [ ] Write `y_i = β₀ + β₁x_i + ε_i` and say what the index `i` ranges over?
- [ ] Say what a scatter plot is for, and one thing it can reveal that a correlation coefficient hides?

Five yeses → go to Chapter 2. Don't linger here.
