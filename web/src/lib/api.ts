/**
 * The API surface, defined once as an interface.
 *
 * There are two implementations — a real HTTP client and a mock — and the
 * application cannot tell them apart. That is the same discipline the Go
 * backend applies to model providers: depend on the interface, choose the
 * implementation at the edge. It is what lets the UI be built and reviewed
 * before the server exists, without the UI being rewritten when it lands.
 */
import type {
  Answer,
  Chapter,
  ChapterFrequencyResponse,
  Course,
  CourseDetail,
  Document,
  DocumentStatus,
  ExamDetail,
  Exam,
  Paged,
  Providers,
  Question,
  QuestionDetail,
  QuestionFilters,
  RetrievalFilters,
  SimilarQuestion,
  User,
} from "@/types/api";

/** One frame of a streamed answer, mirroring the SSE event names in the spec. */
export type AskEvent =
  | { type: "intent"; intent: Answer["intent"] }
  | { type: "sources"; sources: Answer["sources"]; questions?: Question[] }
  | { type: "token"; text: string }
  | { type: "done"; answer: Answer }
  | { type: "error"; message: string };

export interface AskRequest {
  courseId: string;
  question: string;
  conversationId?: string;
  filters?: RetrievalFilters;
}

export interface SkriptraApi {
  me(): Promise<User>;
  providers(): Promise<Providers>;

  listCourses(): Promise<Paged<Course>>;
  getCourse(courseId: string): Promise<CourseDetail>;
  listChapters(courseId: string): Promise<Chapter[]>;
  listExams(courseId: string): Promise<Paged<Exam>>;
  getExam(examId: string): Promise<ExamDetail>;

  listQuestions(courseId: string, filters?: QuestionFilters): Promise<Paged<Question>>;
  getQuestion(questionId: string): Promise<QuestionDetail>;
  similarQuestions(questionId: string, limit?: number): Promise<SimilarQuestion[]>;

  chapterFrequency(courseId: string): Promise<ChapterFrequencyResponse>;

  listDocuments(courseId: string): Promise<Paged<Document>>;
  documentStatus(documentId: string): Promise<DocumentStatus>;

  /** Streams answer events. Returns when the stream completes or aborts. */
  ask(req: AskRequest, onEvent: (e: AskEvent) => void, signal?: AbortSignal): Promise<void>;
}

/** Thrown for any non-2xx response, carrying the server's stable error code. */
export class ApiRequestError extends Error {
  constructor(
    public status: number,
    public code: string,
    message: string,
  ) {
    super(message);
    this.name = "ApiRequestError";
  }
}

const BASE = import.meta.env.VITE_API_BASE_URL ?? "/api/v1";

async function request<T>(path: string, init?: RequestInit): Promise<T> {
  const res = await fetch(`${BASE}${path}`, {
    ...init,
    headers: { "Content-Type": "application/json", ...init?.headers },
  });

  if (!res.ok) {
    let code = "unknown";
    let message = res.statusText;
    try {
      const body = await res.json();
      code = body?.error?.code ?? code;
      message = body?.error?.message ?? message;
    } catch {
      // Non-JSON error body — keep the status text.
    }
    throw new ApiRequestError(res.status, code, message);
  }
  return res.json() as Promise<T>;
}

function query(params: Record<string, unknown | undefined>): string {
  const q = new URLSearchParams();
  for (const [k, v] of Object.entries(params)) {
    if (v !== undefined && v !== null && v !== "") q.set(k, String(v));
  }
  const s = q.toString();
  return s ? `?${s}` : "";
}

export const httpApi: SkriptraApi = {
  me: () => request("/me"),
  providers: () => request("/providers"),

  listCourses: () => request("/courses"),
  getCourse: (id) => request(`/courses/${id}`),
  listChapters: async (id) => (await request<{ data: Chapter[] }>(`/courses/${id}/chapters`)).data,
  listExams: (id) => request(`/courses/${id}/exams`),
  getExam: (id) => request(`/exams/${id}`),

  listQuestions: (id, f = {}) => request(`/courses/${id}/questions${query({ ...f })}`),
  getQuestion: (id) => request(`/questions/${id}`),
  similarQuestions: async (id, limit = 10) =>
    (await request<{ data: SimilarQuestion[] }>(`/questions/${id}/similar${query({ limit })}`)).data,

  chapterFrequency: (id) => request(`/courses/${id}/analytics/chapter-frequency`),

  listDocuments: (id) => request(`/courses/${id}/documents`),
  documentStatus: (id) => request(`/documents/${id}/status`),

  /**
   * Parses the `text/event-stream` response by hand rather than using
   * EventSource, because EventSource cannot issue a POST and the ask endpoint
   * needs a request body.
   */
  async ask(req, onEvent, signal) {
    const res = await fetch(`${BASE}/ask`, {
      method: "POST",
      headers: { "Content-Type": "application/json", Accept: "text/event-stream" },
      body: JSON.stringify({ ...req, stream: true }),
      signal,
    });

    if (!res.ok || !res.body) {
      onEvent({ type: "error", message: `Request failed (${res.status})` });
      return;
    }

    const reader = res.body.getReader();
    const decoder = new TextDecoder();
    let buffer = "";

    while (true) {
      const { done, value } = await reader.read();
      if (done) break;
      buffer += decoder.decode(value, { stream: true });

      // SSE frames are separated by a blank line.
      const frames = buffer.split("\n\n");
      buffer = frames.pop() ?? "";

      for (const frame of frames) {
        let event = "message";
        const dataLines: string[] = [];
        for (const line of frame.split("\n")) {
          if (line.startsWith("event:")) event = line.slice(6).trim();
          else if (line.startsWith("data:")) dataLines.push(line.slice(5).trim());
        }
        if (!dataLines.length) continue;

        const payload = JSON.parse(dataLines.join("\n"));
        switch (event) {
          case "intent":
            onEvent({ type: "intent", intent: payload.intent });
            break;
          case "sources":
            onEvent({ type: "sources", sources: payload.sources, questions: payload.questions });
            break;
          case "token":
            onEvent({ type: "token", text: payload.text });
            break;
          case "done":
            onEvent({ type: "done", answer: payload });
            break;
          case "error":
            onEvent({ type: "error", message: payload.message ?? "Stream error" });
            break;
        }
      }
    }
  },
};
