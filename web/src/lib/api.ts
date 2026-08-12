/**
 * The API surface, defined once as an interface.
 *
 * There are two implementations: a real HTTP client and a mock, and the
 * application cannot tell them apart. That is the same discipline the Go
 * backend applies to model providers: depend on the interface, choose the
 * implementation at the edge. It is what lets the UI be built and reviewed
 * before the server exists, without the UI being rewritten when it lands.
 */
import type {
  Answer,
  Chapter,
  Conversation,
  ConversationDetail,
  ChapterFrequencyResponse,
  Course,
  CourseDetail,
  Document,
  DocumentKind,
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
import { getAccessToken, refreshOnce } from "@/lib/session";

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

export interface UploadMeta {
  kind: DocumentKind;
  year?: number;
  term?: "summer" | "winter";
}

export interface UploadResult {
  id: string;
  status: string;
  /** True when the same bytes were already in the course, so nothing was queued. */
  deduplicated?: boolean;
  message?: string;
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

  createCourse(input: {
    name: string;
    code?: string;
    institution?: string;
    language?: "en" | "de";
  }): Promise<Course>;

  listConversations(courseId: string): Promise<Paged<Conversation>>;
  getConversation(conversationId: string): Promise<ConversationDetail>;
  deleteConversation(conversationId: string): Promise<void>;

  /**
   * Uploads a document and returns as soon as it is queued.
   *
   * Ingestion is asynchronous, so the caller polls `documentStatus` from here.
   * `onProgress` reports bytes sent, which matters because a scanned paper can
   * be tens of megabytes on a slow connection and a spinner with no movement is
   * indistinguishable from a hang.
   */
  uploadDocument(
    courseId: string,
    file: File,
    meta: UploadMeta,
    onProgress?: (fraction: number) => void,
  ): Promise<UploadResult>;

  /** Streams answer events. Returns when the stream completes or aborts. */
  ask(req: AskRequest, onEvent: (e: AskEvent) => void, signal?: AbortSignal): Promise<void>;

  /**
   * Generates a worked solution for a question that has none.
   *
   * The result is never stored as the official solution. It is a model's
   * attempt grounded in course material, and the UI must present it as such.
   */
  generateSolution(
    questionId: string,
    onEvent: (e: AskEvent) => void,
    signal?: AbortSignal,
  ): Promise<void>;
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

/**
 * Absolute URL for a stored PDF, optionally anchored to a page.
 *
 * Built from the same base as every other call rather than as a relative path.
 * A relative href resolves against the web origin, which in development is the
 * Vite dev server: inside its container "localhost" is the container itself,
 * so every citation link died on a proxy error. Going through BASE means the
 * link points wherever the API actually is.
 */
export function documentFileURL(documentId: string, page?: number): string {
  const url = `${BASE}/documents/${documentId}/file`;
  return page ? `${url}#page=${page}` : url;
}

/**
 * Every call carries the access token, and a 401 is retried once behind a
 * refresh.
 *
 * The retry is what makes a 15-minute access token invisible to the user.
 * Without it the app breaks every quarter of an hour, and the obvious
 * workaround, a long-lived token, is the thing short expiry exists to avoid.
 * Retried once only: a second 401 means the session is genuinely gone.
 */
async function request<T>(path: string, init?: RequestInit, retry = true): Promise<T> {
  const token = getAccessToken();
  const res = await fetch(`${BASE}${path}`, {
    ...init,
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
      ...init?.headers,
    },
  });

  if (res.status === 401 && retry) {
    const session = await refreshOnce();
    if (session) return request<T>(path, init, false);
  }

  if (!res.ok) {
    let code = "unknown";
    let message = res.statusText;
    try {
      const body = await res.json();
      code = body?.error?.code ?? code;
      message = body?.error?.message ?? message;
    } catch {
      // Non-JSON error body, keep the status text.
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

  createCourse: (input) =>
    request("/courses", { method: "POST", body: JSON.stringify(input) }),

  listConversations: (id) => request(`/courses/${id}/conversations`),
  getConversation: (id) => request(`/conversations/${id}`),
  deleteConversation: async (id) => {
    await request<void>(`/conversations/${id}`, { method: "DELETE" });
  },

  /**
   * Parses the `text/event-stream` response by hand rather than using
   * EventSource, because EventSource cannot issue a POST and the ask endpoint
   * needs a request body.
   */
  /**
   * XHR rather than fetch, purely for upload progress.
   *
   * fetch() still cannot report request-body progress in any shipping browser,
   * and a large scan on a slow connection needs a moving bar rather than an
   * indefinite spinner. Everything else in this client uses fetch.
   */
  uploadDocument(courseId, file, meta, onProgress) {
    return new Promise<UploadResult>((resolve, reject) => {
      const form = new FormData();
      form.append("file", file);
      form.append("kind", meta.kind);
      if (meta.year) form.append("year", String(meta.year));
      if (meta.term) form.append("term", meta.term);

      const xhr = new XMLHttpRequest();
      xhr.open("POST", `${BASE}/courses/${courseId}/documents`);
      xhr.withCredentials = true;
      const token = getAccessToken();
      if (token) xhr.setRequestHeader("Authorization", `Bearer ${token}`);

      xhr.upload.onprogress = (e) => {
        if (e.lengthComputable) onProgress?.(e.loaded / e.total);
      };

      xhr.onload = () => {
        let body: Record<string, unknown> = {};
        try {
          body = JSON.parse(xhr.responseText);
        } catch {
          // Fall through to the status-based message below.
        }

        // 202 queued, 200 deduplicated. Both are success, and the contract
        // distinguishes them so the UI can say "already uploaded" rather than
        // pretending work was started.
        if (xhr.status === 200 || xhr.status === 202) {
          resolve(body as unknown as UploadResult);
          return;
        }
        const err = body as { error?: { message?: string; code?: string } };
        reject(
          new ApiRequestError(
            xhr.status,
            err.error?.code ?? "upload_failed",
            err.error?.message ?? `Upload failed (${xhr.status})`,
          ),
        );
      };

      xhr.onerror = () =>
        reject(new ApiRequestError(0, "network", "Could not reach the server."));
      xhr.send(form);
    });
  },

  ask(req, onEvent, signal) {
    return streamSSE(`${BASE}/ask`, { ...req, stream: true }, onEvent, signal);
  },

  generateSolution(questionId, onEvent, signal) {
    return streamSSE(`${BASE}/questions/${questionId}/solution`, { stream: true }, onEvent, signal);
  },
};

/**
 * Parses a `text/event-stream` response by hand rather than using EventSource,
 * because EventSource cannot issue a POST and both endpoints need a body.
 */
async function streamSSE(
  url: string,
  body: unknown,
  onEvent: (e: AskEvent) => void,
  signal?: AbortSignal,
  retry = true,
): Promise<void> {
  let res: Response;
  try {
    const token = getAccessToken();
    res = await fetch(url, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
        Accept: "text/event-stream",
        ...(token ? { Authorization: `Bearer ${token}` } : {}),
      },
      body: JSON.stringify(body),
      signal,
    });
  } catch (err) {
    if ((err as Error)?.name === "AbortError") return;
    onEvent({ type: "error", message: "Could not reach the server." });
    return;
  }

  // The same one-shot refresh as request(). An expired token must not surface
  // as "something went wrong" mid-question.
  if (res.status === 401 && retry) {
    const session = await refreshOnce();
    if (session) return streamSSE(url, body, onEvent, signal, false);
  }

  if (!res.ok || !res.body) {
    // A non-streaming error body still carries the contract's error envelope,
    // so surface the server's message rather than a bare status code.
    let message = `Request failed (${res.status})`;
    try {
      const parsed = await res.json();
      message = parsed?.error?.message ?? message;
    } catch {
      // Non-JSON body; keep the status message.
    }
    onEvent({ type: "error", message });
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
}
