/**
 * Types mirroring `api/openapi.yaml`.
 *
 * The contract is the single source of truth. When the Go server lands, these
 * are regenerated from the spec rather than edited by hand, if you find
 * yourself changing a type here to make a component compile, change the spec.
 */

export type Term = "summer" | "winter";

export type DocumentKind = "exam" | "solution" | "notes" | "textbook" | "syllabus";

export type IngestStatus =
  | "queued"
  | "parsing"
  | "segmenting"
  | "classifying"
  | "embedding"
  | "indexed"
  | "failed";

/** The query router's decision. Only `explain` and `hybrid` invoke the model. */
export type QueryIntent = "enumerate" | "explain" | "similar" | "analyse" | "hybrid";

export interface PageMeta {
  page: number;
  pageSize: number;
  total: number;
  totalPages: number;
}

export interface Paged<T> {
  data: T[];
  meta: PageMeta;
}

export interface ApiError {
  error: {
    code: string;
    message: string;
    details?: { field?: string; issue?: string }[];
    requestId?: string;
  };
}

export interface User {
  id: string;
  displayName: string;
  email?: string;
}

export interface ProviderInfo {
  provider: string;
  model: string;
  /** True when inference never leaves this machine. */
  local?: boolean;
  dimensions?: number;
}

export interface Providers {
  llm: ProviderInfo;
  embedding: ProviderInfo;
}

export interface Course {
  id: string;
  name: string;
  code?: string;
  institution?: string;
  language?: "en" | "de";
  examCount: number;
  questionCount: number;
  createdAt?: string;
}

export interface CourseDetail extends Course {
  chapterCount?: number;
  documentCount?: number;
  yearRange?: { from: number; to: number };
}

export interface Chapter {
  id: string;
  number: number;
  title: string;
  topics?: string[];
  questionCount?: number;
}

/**
 * Nullable on a question, and carries a confidence score, a question is not
 * always confidently classifiable, and the UI shows that rather than
 * presenting a guess as fact.
 */
export interface ChapterRef {
  id: string;
  number: number;
  title: string;
  confidence?: number;
}

export interface Exam {
  id: string;
  courseId?: string;
  year: number;
  term: Term;
  title?: string;
  documentId?: string;
  hasSolutions?: boolean;
  questionCount: number;
}

export interface ExamDetail extends Exam {
  questions?: Question[];
}

export interface Question {
  id: string;
  examId?: string;
  number: string;
  text: string;
  marks?: number;
  /** 1-indexed page in the source PDF; drives citation deep links. */
  sourcePage: number;
  chapter?: ChapterRef;
  topic?: string;
  year?: number;
  term?: Term;
  hasSolution?: boolean;
}

export interface QuestionDetail extends Question {
  documentId?: string;
  solutionText?: string;
  solutionSourcePage?: number;
}

export interface SimilarQuestion {
  question: Question;
  score: number;
}

export interface Citation {
  documentId: string;
  documentTitle: string;
  documentKind?: DocumentKind;
  page: number;
  questionId?: string;
  questionNumber?: string;
  /** Pre-rendered by the server, e.g. "2025 Summer Exam · Q4 · Page 3". */
  label?: string;
}

export interface Usage {
  promptTokens?: number;
  completionTokens?: number;
  retrievedChunks?: number;
  latencyMs?: number;
  provider?: string;
  model?: string;
}

export interface Answer {
  conversationId: string;
  messageId: string;
  intent: QueryIntent;
  /** Markdown. */
  answer: string;
  sources: Citation[];
  /** Populated for `enumerate` and `hybrid` intents. */
  questions?: Question[];
  usage?: Usage;
}

/** A thread of questions and answers about one course. */
export interface Conversation {
  id: string;
  courseId: string;
  /** Taken from the first question asked, truncated at a word break. */
  title: string;
  messageCount: number;
  createdAt: string;
  updatedAt: string;
}

/**
 * One turn. Citations are stored on the message rather than resolved at read
 * time, so a past answer reads exactly as it was given even after the document
 * it cited has been deleted or re-ingested.
 */
export interface Message {
  id: string;
  role: "user" | "assistant";
  content: string;
  intent?: QueryIntent;
  sources: Citation[];
  usage?: Usage;
  createdAt: string;
}

export interface ConversationDetail extends Conversation {
  messages: Message[];
}

export interface RetrievalFilters {
  chapterIds?: string[];
  chapterNumbers?: number[];
  yearFrom?: number;
  yearTo?: number;
  documentKinds?: DocumentKind[];
  /** Weights the user's own uploads above shared material. */
  preferOwnDocuments?: boolean;
}

export interface SearchHit {
  chunkId: string;
  text: string;
  score: number;
  denseScore?: number;
  sparseScore?: number;
  citation: Citation;
}

export interface Document {
  id: string;
  courseId?: string;
  filename: string;
  kind: DocumentKind;
  status: IngestStatus;
  year?: number;
  term?: Term;
  pageCount?: number;
  sizeBytes?: number;
  contentHash?: string;
  uploadedAt?: string;
}

export interface DocumentStatus {
  documentId: string;
  status: IngestStatus;
  progress: number;
  stageDetail?: string;
  questionsExtracted?: number;
  error?: string;
}

export interface ChapterFrequency {
  chapter: ChapterRef;
  questionCount: number;
  share: number;
  examCount?: number;
  byYear?: { year: number; questionCount: number }[];
}

export interface ChapterFrequencyResponse {
  totalQuestions: number;
  data: ChapterFrequency[];
}

export interface QuestionFilters {
  chapterId?: string;
  chapterNumber?: number;
  yearFrom?: number;
  yearTo?: number;
  term?: Term;
  q?: string;
  sort?: "newest" | "oldest" | "chapter";
  page?: number;
  pageSize?: number;
}
