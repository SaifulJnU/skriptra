/**
 * Mock implementation of `SkriptraApi`.
 *
 * The intent router below is a deliberately simplified version of the real one
 * that will live in Go. It is here so the UI is built against the *behaviour*
 * that matters, different question shapes take different paths and render
 * differently, rather than against a single generic "chat" response that would
 * hide the whole design.
 */
import type { AskEvent, AskRequest, SkriptraApi } from "@/lib/api";
import type {
  Citation,
  IngestStatus,
  Question,
  QueryIntent,
  SimilarQuestion,
} from "@/types/api";
import {
  COURSE_ID,
  chapterFrequency,
  chapters,
  courses,
  documents,
  exams,
  questions,
} from "./data";

const delay = (ms: number) => new Promise((r) => setTimeout(r, ms));

/** Documents uploaded during this session, so their status can advance. */
const mockIngest = new Map<string, { started: number; filename: string }>();

function paged<T>(data: T[], page = 1, pageSize = 20) {
  const start = (page - 1) * pageSize;
  return {
    data: data.slice(start, start + pageSize),
    meta: {
      page,
      pageSize,
      total: data.length,
      totalPages: Math.max(1, Math.ceil(data.length / pageSize)),
    },
  };
}

/** Maps "chapter two", "kapitel 2", "ch. 2" onto a chapter number. */
const WORD_NUMBERS: Record<string, number> = {
  one: 1, two: 2, three: 3, four: 4, five: 5,
  eins: 1, zwei: 2, drei: 3, vier: 4, fünf: 5,
};

export function resolveChapter(text: string): number | undefined {
  const t = text.toLowerCase();
  const digit = t.match(/(?:chapter|kapitel|ch\.?|kap\.?)\s*(\d+)/);
  if (digit) return Number(digit[1]);
  const word = t.match(/(?:chapter|kapitel)\s+([a-zäöü]+)/);
  if (word && WORD_NUMBERS[word[1]]) return WORD_NUMBERS[word[1]];
  // Fall back to matching a chapter by title.
  const byTitle = chapters.find((c) => t.includes(c.title.toLowerCase()));
  return byTitle?.number;
}

/**
 * The query router. Only `explain` and `hybrid` reach a language model; the
 * others are database queries and are therefore exact and instant.
 */
export function classifyIntent(text: string): QueryIntent {
  const t = text.toLowerCase();
  if (/\b(most|frequen|often|trend|distribution|how many times|statistic)\b/.test(t)) return "analyse";
  if (/\b(similar|like this|same kind|repeated|appeared before)\b/.test(t)) return "similar";
  if (/\b(all|list|every|give me the|show me the|which questions)\b/.test(t)) {
    return /\b(why|how|explain|derive|prove|meaning)\b/.test(t) ? "hybrid" : "enumerate";
  }
  return "explain";
}

function citationFor(q: Question): Citation {
  const exam = exams.find((e) => e.id === q.examId);
  return {
    documentId: exam?.documentId ?? "doc-unknown",
    documentTitle: `${exam?.year} ${exam?.term === "summer" ? "Summer" : "Winter"} Exam`,
    documentKind: "exam",
    page: q.sourcePage,
    questionId: q.id,
    questionNumber: q.number,
    label: `${exam?.year} ${exam?.term === "summer" ? "Summer" : "Winter"} Exam · Q${q.number} · Page ${q.sourcePage}`,
  };
}

/**
 * Canned explanations, keyed by topic.
 *
 * There is no model behind this yet. The honest thing for a mock to do is
 * answer what it actually recognises and say plainly when it does not,  * a mock that returns a confident essay for every input hides exactly the
 * failure mode (answering off-topic from the wrong passages) that the real
 * system has to be tested against.
 */
const EXPLANATIONS: { match: RegExp; text: string }[] = [
  {
    match: /\b(regression|linear model|what is the linear model)\b/,
    text: `In this course, **regression** means modelling a response *y* as a linear combination of predictors plus noise: y = Xβ + ε.\n\nThe word "linear" refers to linearity **in the parameters β**, not in the predictors, y = β₀ + β₁x + β₂x² is still a linear model. X is the design matrix, β is what you estimate, and ε carries everything the predictors do not explain.\n\nChapter 1 sets out the assumptions on ε (zero mean, constant variance, uncorrelated) that everything later depends on. Chapter 2 estimates β; Chapter 3 does inference on it.`,
  },
  {
    match: /\b(ols|least squares|gauss.?markov|blue)\b/,
    text: `**Ordinary least squares** chooses β̂ to minimise the residual sum of squares (y − Xβ)ᵀ(y − Xβ), giving β̂ = (XᵀX)⁻¹Xᵀy whenever XᵀX is invertible.\n\nThe **Gauss-Markov theorem** says that under the standard assumptions, errors with mean zero, constant variance and no correlation, this estimator is BLUE: the *best linear unbiased* estimator, where "best" means smallest variance among all linear unbiased estimators.\n\nNote what it does *not* require: normality. That assumption is only needed for the distributional results in Chapter 3.`,
  },
  {
    match: /\b(maximum likelihood|mle|likelihood)\b/,
    text: `Maximum likelihood is used here because the model specifies a full parametric distribution for the errors, once ε ~ N(0, σ²I) is assumed, the likelihood is completely determined by β and σ².\n\nUnder that normality assumption the ML estimator of β coincides exactly with the OLS estimator, which is why the two approaches agree on the coefficients. They part company on the variance: the ML estimator of σ² divides by *n* rather than *n − p*, and is therefore **biased downward**.\n\nThat bias is why the unbiased estimator s² = RSS/(n − p) is the one used for inference.`,
  },
  {
    match: /\b(f.?test|t.?test|hypothesis|significance|p.?value|confidence)\b/,
    text: `Hypothesis testing in the linear model compares a restricted model against an unrestricted one.\n\nThe **F-test** for q simultaneous restrictions uses F = [(RSS_r − RSS_u)/q] / [RSS_u/(n − p)], which is F(q, n − p) under the null. The **t-test** is the special case q = 1, and t² = F exactly in that case.\n\nBoth rely on normality of the errors, without it these are approximate, justified only by large-sample arguments.`,
  },
  {
    // "cook" must be anchored to Cook's distance, a bare keyword matched
    // "how do I cook biryani" and confidently returned a diagnostics lecture.
    match: /\b(residual|diagnostic|leverage|cook['’]?s distance|influential observation|outlier)\b/,
    text: `Diagnostics ask whether the fitted model's assumptions actually hold.\n\n**Leverage** h_ii measures how unusual an observation's *predictor* values are, it depends only on X, not on y. **Cook's distance** combines leverage with the size of the residual to measure actual influence on the fitted coefficients.\n\nThis is why a high-leverage point can have low influence: if its response happens to sit exactly where the model predicts, removing it changes nothing.`,
  },
  {
    match: /\b(glm|generalized linear|link function|logistic|poisson)\b/,
    text: `A **generalized linear model** relaxes two things at once: the response need not be normal, and the mean need not be linear in the predictors directly.\n\nInstead a **link function** g connects them: g(μ) = Xβ. The canonical link comes from writing the distribution in exponential-family form, log for the Poisson, logit for the binomial, identity for the normal, which recovers ordinary regression as a special case.`,
  },
];

function answerText(
  intent: QueryIntent,
  chapterNo: number | undefined,
  matched: Question[],
  question: string,
): string {
  switch (intent) {
    case "enumerate": {
      const ch = chapters.find((c) => c.number === chapterNo);
      return `Found **${matched.length} questions**${ch ? ` in Chapter ${ch.number}, ${ch.title}` : ""}, across ${new Set(matched.map((q) => q.year)).size} exam years.\n\nThey are listed below in reverse chronological order. This is an exhaustive result from the question index, not a sample.`;
    }
    case "analyse": {
      const top = [...chapterFrequency].sort((a, b) => b.questionCount - a.questionCount)[0];
      return `**Chapter ${top.chapter.number}, ${top.chapter.title}** is the most frequently tested, accounting for ${(top.share * 100).toFixed(0)}% of all indexed questions and appearing in ${top.examCount} of ${exams.length} exams.\n\nSee the chart on the Analytics tab for the full distribution and the year-by-year breakdown.`;
    }
    case "similar":
      return `Found ${matched.length} closely related questions across previous years. The strongest matches reuse the same derivation with different notation.`;
    default: {
      const hit = EXPLANATIONS.find((e) => e.match.test(question.toLowerCase()));
      if (hit) return hit.text;
      // Say so, rather than returning a confident answer to a different
      // question. The real system must refuse the same way when retrieval
      // comes back empty, that behaviour is measured by the eval harness.
      return `**No indexed passage covers that yet.**\n\nThis build is running on mock data, there is no model and no document index behind it. The mock only has worked explanations for the core Linear Models topics: regression and the linear model, OLS and Gauss-Markov, maximum likelihood, hypothesis testing, diagnostics, and GLMs.\n\nOnce the Go backend and the ingestion pipeline are running, this answer comes from real retrieval over uploaded papers, and when the corpus genuinely does not cover a question, it should still say so instead of inventing one.`;
    }
  }
}

export const mockApi: SkriptraApi = {
  async me() {
    await delay(80);
    return { id: "user-1", displayName: "Saiful", email: "saiful@example.com" };
  },

  async providers() {
    await delay(80);
    return {
      llm: { provider: "ollama", model: "llama3.1:8b", local: true },
      embedding: { provider: "ollama", model: "bge-m3", local: true, dimensions: 768 },
    };
  },

  async listCourses() {
    await delay(180);
    return paged(courses, 1, 20);
  },

  async getCourse(courseId) {
    await delay(140);
    const c = courses.find((x) => x.id === courseId) ?? courses[0];
    return {
      ...c,
      chapterCount: chapters.length,
      documentCount: documents.length,
      yearRange: { from: 2020, to: 2025 },
    };
  },

  async listChapters() {
    await delay(120);
    return chapters;
  },

  async listExams() {
    await delay(160);
    return paged(
      [...exams].sort((a, b) => b.year - a.year || a.term.localeCompare(b.term)),
      1,
      50,
    );
  },

  async getExam(examId) {
    await delay(160);
    const exam = exams.find((e) => e.id === examId) ?? exams[0];
    return { ...exam, questions: questions.filter((q) => q.examId === exam.id) };
  },

  async listQuestions(_courseId, f = {}) {
    await delay(200);
    let out = questions;
    if (f.chapterNumber !== undefined) out = out.filter((q) => q.chapter?.number === f.chapterNumber);
    if (f.chapterId) out = out.filter((q) => q.chapter?.id === f.chapterId);
    if (f.yearFrom) out = out.filter((q) => (q.year ?? 0) >= f.yearFrom!);
    if (f.yearTo) out = out.filter((q) => (q.year ?? 0) <= f.yearTo!);
    if (f.term) out = out.filter((q) => q.term === f.term);
    if (f.q) {
      const needle = f.q.toLowerCase();
      out = out.filter((q) => q.text.toLowerCase().includes(needle));
    }
    out = [...out].sort((a, b) =>
      f.sort === "oldest"
        ? (a.year ?? 0) - (b.year ?? 0)
        : f.sort === "chapter"
          ? (a.chapter?.number ?? 99) - (b.chapter?.number ?? 99)
          : (b.year ?? 0) - (a.year ?? 0),
    );
    return paged(out, f.page ?? 1, f.pageSize ?? 20);
  },

  async getQuestion(questionId) {
    await delay(140);
    const q = questions.find((x) => x.id === questionId) ?? questions[0];
    const exam = exams.find((e) => e.id === q.examId);
    return {
      ...q,
      documentId: exam?.documentId,
      solutionText: q.hasSolution
        ? "Start from the log-likelihood ℓ(β, σ²) = −(n/2)log(2πσ²) − (1/2σ²)(y − Xβ)ᵀ(y − Xβ). Differentiating with respect to β and setting the score to zero recovers the normal equations XᵀXβ = Xᵀy, so β̂_ML = β̂_OLS. Differentiating with respect to σ² gives σ̂²_ML = RSS/n, whose expectation is ((n − p)/n)σ², biased downward by a factor of (n − p)/n."
        : undefined,
      solutionSourcePage: q.hasSolution ? q.sourcePage + 4 : undefined,
    };
  },

  async similarQuestions(questionId, limit = 10) {
    await delay(220);
    const source = questions.find((q) => q.id === questionId);
    if (!source) return [];
    // Same chapter and topic scores highest, then same chapter, then neither.
    const scored: SimilarQuestion[] = questions
      .filter((q) => q.id !== questionId)
      .map((q) => {
        let score = 0.2;
        if (q.chapter?.number === source.chapter?.number) score += 0.35;
        if (q.topic && q.topic === source.topic) score += 0.4;
        if (q.text === source.text) score = 0.97;
        return { question: q, score: Math.min(0.99, score) };
      })
      .filter((s) => s.score > 0.5)
      .sort((a, b) => b.score - a.score);
    return scored.slice(0, limit);
  },

  async chapterFrequency() {
    await delay(200);
    return { totalQuestions: questions.length, data: chapterFrequency };
  },

  async listDocuments() {
    await delay(160);
    return paged(documents, 1, 50);
  },

  async documentStatus(documentId) {
    await delay(100);

    // A document uploaded in this session advances through the real stage
    // sequence on a timer.
    const job = mockIngest.get(documentId);
    if (job) {
      const elapsed = Date.now() - job.started;
      const stages: { at: number; status: IngestStatus; detail?: string }[] = [
        { at: 0, status: "queued" },
        { at: 1200, status: "parsing", detail: "extracting text" },
        { at: 3000, status: "segmenting", detail: "splitting into questions" },
        { at: 4800, status: "classifying", detail: "classifying question 4 of 6" },
        { at: 7000, status: "embedding", detail: "computing embeddings" },
        { at: 9000, status: "indexed" },
      ];
      const current = [...stages].reverse().find((s) => elapsed >= s.at)!;
      return {
        documentId,
        status: current.status,
        progress: Math.min(elapsed / 9000, 1),
        stageDetail: current.detail,
        questionsExtracted: current.status === "indexed" ? 6 : 0,
      };
    }

    const d = documents.find((x) => x.id === documentId);
    return {
      documentId,
      status: d?.status ?? "indexed",
      progress: d?.status === "indexed" ? 1 : 0.62,
      stageDetail: d?.status === "indexed" ? undefined : "classifying question 12 of 31",
      questionsExtracted: d?.status === "indexed" ? 7 : 12,
    };
  },

  async ask(req: AskRequest, onEvent: (e: AskEvent) => void, signal?: AbortSignal) {
    const intent = classifyIntent(req.question);
    const chapterNo = resolveChapter(req.question) ?? req.filters?.chapterNumbers?.[0];

    await delay(280);
    if (signal?.aborted) return;
    onEvent({ type: "intent", intent });

    const matched =
      intent === "enumerate" || intent === "hybrid"
        ? questions
            .filter((q) => (chapterNo ? q.chapter?.number === chapterNo : true))
            .sort((a, b) => (b.year ?? 0) - (a.year ?? 0))
        : questions
            .filter((q) => (chapterNo ? q.chapter?.number === chapterNo : true))
            .slice(0, 4);

    const text = answerText(intent, chapterNo, matched, req.question);

    // An answer that grounds nothing must cite nothing. Showing sources beside
    // "no indexed passage covers that" would be the exact dishonesty this
    // product exists to avoid.
    const grounded = !text.startsWith("**No indexed passage");
    const sources = grounded ? matched.slice(0, 5).map(citationFor) : [];

    await delay(320);
    if (signal?.aborted) return;
    // Sources arrive before generation so citations render while tokens stream.
    onEvent({
      type: "sources",
      sources,
      questions: intent === "enumerate" || intent === "hybrid" ? matched : undefined,
    });
    const tokens = text.match(/\S+\s*/g) ?? [];
    let streamed = "";
    for (const tok of tokens) {
      if (signal?.aborted) return;
      await delay(14);
      streamed += tok;
      onEvent({ type: "token", text: tok });
    }

    onEvent({
      type: "done",
      answer: {
        conversationId: req.conversationId ?? crypto.randomUUID(),
        messageId: crypto.randomUUID(),
        intent,
        answer: streamed,
        sources,
        questions: intent === "enumerate" || intent === "hybrid" ? matched : undefined,
        usage: {
          promptTokens: 1240 + matched.length * 30,
          completionTokens: tokens.length,
          retrievedChunks: sources.length,
          latencyMs: 900 + tokens.length * 14,
          provider: "ollama",
          model: "llama3.1:8b",
        },
      },
    });
  },

  // Walks the same stages the real pipeline reports, so the upload UI can be
  // built and reviewed without a backend, and the progress states are exercised
  // rather than assumed.
  async uploadDocument(_courseId: string, file: File, _meta, onProgress) {
    for (let p = 0; p <= 1; p += 0.2) {
      await delay(90);
      onProgress?.(Math.min(p, 1));
    }
    const id = crypto.randomUUID();
    mockIngest.set(id, { started: Date.now(), filename: file.name });
    return { id, status: "queued" };
  },

  async generateSolution(questionId: string, onEvent: (e: AskEvent) => void, signal?: AbortSignal) {
    const q = questions.find((x) => x.id === questionId);
    await delay(400);
    if (signal?.aborted) return;

    const sources = questions
      .filter((x) => x.id !== questionId && x.chapter?.number === q?.chapter?.number)
      .slice(0, 3)
      .map(citationFor);
    onEvent({ type: "sources", sources });

    const text = `Start from the model y = Xβ + ε and minimise the residual sum of squares S(β) = (y − Xβ)ᵀ(y − Xβ).\n\nDifferentiating with respect to β and setting the derivative to zero gives the normal equations XᵀXβ = Xᵀy, so β̂ = (XᵀX)⁻¹Xᵀy whenever XᵀX is invertible.\n\nUnbiasedness needs only E[ε] = 0 and that X is fixed: E[β̂] = β + (XᵀX)⁻¹XᵀE[ε] = β. Neither constant variance nor normality is required for this part.`;

    for (const tok of text.match(/\S+\s*/g) ?? []) {
      if (signal?.aborted) return;
      await delay(12);
      onEvent({ type: "token", text: tok });
    }

    onEvent({
      type: "done",
      answer: {
        conversationId: crypto.randomUUID(),
        messageId: crypto.randomUUID(),
        intent: "explain",
        answer: text,
        sources,
        usage: { retrievedChunks: sources.length, latencyMs: 1400, provider: "ollama", model: "llama3.2:3b" },
      },
    });
  },
};

export { COURSE_ID };
