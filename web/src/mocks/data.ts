/**
 * Mock corpus, shaped exactly like the API contract.
 *
 * This is not filler. It is a realistic Linear Models corpus, because a mock
 * that is too clean hides the states that matter: unclassified questions, low
 * classification confidence, documents still ingesting, exams without
 * solutions. Those states are designed for here so they are not retrofitted
 * later.
 */
import type {
  Chapter,
  ChapterFrequency,
  Course,
  Document,
  Exam,
  Question,
} from "@/types/api";

export const COURSE_ID = "22222222-2222-2222-2222-222222222222";
const NLP_ID = "22222222-2222-2222-2222-222222222299";

export const courses: Course[] = [
  {
    id: COURSE_ID,
    name: "Linear Models",
    code: "STAT-412",
    institution: "TU Dortmund",
    language: "en",
    examCount: 12,
    questionCount: 436,
    createdAt: "2026-08-01T10:00:00Z",
  },
  {
    id: NLP_ID,
    name: "Natural Language Processing",
    code: "CS-508",
    institution: "TU Dortmund",
    language: "en",
    examCount: 8,
    questionCount: 281,
    createdAt: "2026-08-03T09:00:00Z",
  },
];

export const chapters: Chapter[] = [
  { id: "ch-1", number: 1, title: "The Linear Model", topics: ["design matrix", "assumptions"], questionCount: 52 },
  { id: "ch-2", number: 2, title: "Least Squares Estimation", topics: ["OLS", "Gauss-Markov", "BLUE"], questionCount: 78 },
  { id: "ch-3", number: 3, title: "Inference and Hypothesis Testing", topics: ["F-test", "t-test", "confidence regions"], questionCount: 187 },
  { id: "ch-4", number: 4, title: "Model Diagnostics", topics: ["residuals", "leverage", "Cook's distance"], questionCount: 71 },
  { id: "ch-5", number: 5, title: "Generalized Linear Models", topics: ["link function", "logistic", "Poisson"], questionCount: 48 },
];

const YEARS: { year: number; term: "summer" | "winter"; solutions: boolean }[] = [
  { year: 2025, term: "summer", solutions: true },
  { year: 2025, term: "winter", solutions: true },
  { year: 2024, term: "summer", solutions: true },
  { year: 2024, term: "winter", solutions: false },
  { year: 2023, term: "summer", solutions: true },
  { year: 2023, term: "winter", solutions: true },
  { year: 2022, term: "summer", solutions: false },
  { year: 2022, term: "winter", solutions: true },
  { year: 2021, term: "summer", solutions: true },
  { year: 2021, term: "winter", solutions: false },
  { year: 2020, term: "summer", solutions: true },
  { year: 2020, term: "winter", solutions: true },
];

export const exams: Exam[] = YEARS.map((y, i) => ({
  id: `exam-${y.year}-${y.term}`,
  courseId: COURSE_ID,
  year: y.year,
  term: y.term,
  title: `${y.year} ${y.term === "summer" ? "Summer" : "Winter"}`,
  documentId: `doc-${y.year}-${y.term}`,
  hasSolutions: y.solutions,
  questionCount: 5 + (i % 3),
}));

/** Question stems reused across years, deliberately, so "similar questions" has something true to find. */
const STEMS: { text: string; chapter: number; topic: string; marks: number; confidence?: number }[] = [
  { text: "Derive the ordinary least squares estimator for β in the linear model y = Xβ + ε and state the assumptions required for it to be unbiased.", chapter: 2, topic: "OLS", marks: 12 },
  { text: "State and prove the Gauss-Markov theorem. Explain precisely what 'best' means in the acronym BLUE.", chapter: 2, topic: "Gauss-Markov", marks: 14 },
  { text: "Construct an F-test for the joint significance of two regression coefficients. Give the null distribution and the rejection region at level α = 0.05.", chapter: 3, topic: "F-test", marks: 10 },
  { text: "Derive the distribution of the residual sum of squares under the normal linear model and use it to build a confidence interval for σ².", chapter: 3, topic: "Distribution theory", marks: 12 },
  { text: "Explain the relationship between the t-test for a single coefficient and the corresponding F-test. Show that t² = F in this case.", chapter: 3, topic: "t-test", marks: 8 },
  { text: "A researcher reports R² = 0.94 but residual plots show clear curvature. Discuss what has gone wrong and which diagnostics you would run.", chapter: 4, topic: "Residual analysis", marks: 10 },
  { text: "Define leverage and Cook's distance. Explain how a point can have high leverage but low influence.", chapter: 4, topic: "Influence", marks: 9 },
  { text: "Define the link function in a generalized linear model and derive the canonical link for the Poisson distribution.", chapter: 5, topic: "Link functions", marks: 11 },
  { text: "State the assumptions of the classical linear model and explain the consequence of violating each one.", chapter: 1, topic: "Assumptions", marks: 10 },
  { text: "Discuss the effect of multicollinearity on the variance of the OLS estimator and describe one remedy.", chapter: 2, topic: "Multicollinearity", marks: 9 },
  { text: "Given the partitioned model y = X₁β₁ + X₂β₂ + ε, derive the Frisch-Waugh-Lovell result.", chapter: 3, topic: "Partitioned regression", marks: 13, confidence: 0.58 },
  { text: "Compare the interpretation of coefficients in a logistic regression against a linear probability model.", chapter: 5, topic: "Logistic regression", marks: 10 },
];

export const questions: Question[] = exams.flatMap((exam, ei) =>
  Array.from({ length: exam.questionCount }, (_, qi) => {
    const stem = STEMS[(ei * 3 + qi) % STEMS.length];
    const chapter = chapters.find((c) => c.number === stem.chapter)!;
    // Every fifth question is left unclassified on purpose, the UI has to
    // handle a question the classifier could not place.
    const unclassified = (ei + qi) % 11 === 0;
    return {
      id: `q-${exam.year}-${exam.term}-${qi + 1}`,
      examId: exam.id,
      number: String(qi + 1),
      text: stem.text,
      marks: stem.marks,
      sourcePage: qi + 1,
      chapter: unclassified
        ? undefined
        : {
            id: chapter.id,
            number: chapter.number,
            title: chapter.title,
            confidence: stem.confidence ?? 0.82 + ((ei * 7 + qi * 3) % 17) / 100,
          },
      topic: unclassified ? undefined : stem.topic,
      year: exam.year,
      term: exam.term,
      hasSolution: exam.hasSolutions,
    } satisfies Question;
  }),
);

export const documents: Document[] = [
  ...exams.map<Document>((e) => ({
    id: e.documentId!,
    courseId: COURSE_ID,
    filename: `LiMo_${e.year}_${e.term === "summer" ? "SS" : "WS"}.pdf`,
    kind: "exam",
    status: "indexed",
    year: e.year,
    term: e.term,
    pageCount: 8,
    sizeBytes: 1_240_000,
    uploadedAt: "2026-08-05T12:00:00Z",
  })),
  {
    id: "doc-notes-1",
    courseId: COURSE_ID,
    filename: "Lecture_Notes_Ch3.pdf",
    kind: "notes",
    status: "indexed",
    pageCount: 64,
    sizeBytes: 4_100_000,
    uploadedAt: "2026-08-06T08:30:00Z",
  },
  {
    // Mid-ingest, so the upload UI has a real in-progress state to render.
    id: "doc-ingesting",
    courseId: COURSE_ID,
    filename: "LiMo_2019_WS_scan.pdf",
    kind: "exam",
    status: "classifying",
    year: 2019,
    term: "winter",
    pageCount: 7,
    sizeBytes: 8_900_000,
    uploadedAt: "2026-08-09T17:55:00Z",
  },
];

export const chapterFrequency: ChapterFrequency[] = chapters.map((c) => {
  const count = questions.filter((q) => q.chapter?.number === c.number).length;
  return {
    chapter: { id: c.id, number: c.number, title: c.title },
    questionCount: count,
    share: count / questions.length,
    examCount: new Set(
      questions.filter((q) => q.chapter?.number === c.number).map((q) => q.examId),
    ).size,
    byYear: [...new Set(exams.map((e) => e.year))]
      .sort()
      .map((year) => ({
        year,
        questionCount: questions.filter(
          (q) => q.year === year && q.chapter?.number === c.number,
        ).length,
      })),
  };
});
