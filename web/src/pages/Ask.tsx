import { useEffect, useRef, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { ArrowUp, Database, FileText, Loader2, Mic, Sparkles, Square } from "lucide-react";
import { api, usingMocks } from "@/lib/client";
import { useVoiceInput } from "@/lib/useVoiceInput";
import type { AskEvent } from "@/lib/api";
import type { Citation, QueryIntent, Question, Usage } from "@/types/api";
import { Card, PageHeader } from "@/components/ui";
import { examLabel } from "@/lib/utils";

/**
 * Explains what the router decided, in the user's language.
 *
 * Showing the intent is not debug output: it is how the product earns trust.
 * "Looked up 187 questions" and "read 5 passages and wrote an answer" are very
 * different claims, and a student revising for an exam deserves to know which
 * one they just got.
 */
const INTENT_COPY: Record<QueryIntent, { label: string; detail: string }> = {
  enumerate: { label: "Exhaustive lookup", detail: "Answered from the question index, complete, not a sample" },
  explain: { label: "Explained from sources", detail: "Retrieved passages, then generated an answer" },
  similar: { label: "Similarity search", detail: "Matched against question embeddings" },
  analyse: { label: "Statistics", detail: "Computed directly from the database, no model involved" },
  hybrid: { label: "Lookup + explanation", detail: "Selected questions, then explained them" },
};

const SUGGESTIONS = [
  "Give me all Chapter 3 questions from the last five years",
  "Why is maximum likelihood used in this question?",
  "Which chapters are tested most often?",
  "Has a question about the Gauss-Markov theorem appeared before?",
];

interface Turn {
  question: string;
  intent?: QueryIntent;
  text: string;
  sources: Citation[];
  questions?: Question[];
  usage?: Usage;
  streaming: boolean;
  error?: string;
}

export default function Ask() {
  const { courseId = "" } = useParams();
  const [input, setInput] = useState("");
  const [turns, setTurns] = useState<Turn[]>([]);
  const abortRef = useRef<AbortController | null>(null);
  const bottomRef = useRef<HTMLDivElement>(null);

  const busy = turns.some((t) => t.streaming);

  // Dictation fills the composer rather than sending on its own, a
  // mis-transcribed question should be correctable before it is asked.
  const voice = useVoiceInput();
  useEffect(() => {
    if (voice.transcript) setInput(voice.transcript);
  }, [voice.transcript]);

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth", block: "end" });
  }, [turns]);

  async function send(question: string) {
    if (!question.trim() || busy) return;
    setInput("");

    const index = turns.length;
    setTurns((prev) => [...prev, { question, text: "", sources: [], streaming: true }]);

    const controller = new AbortController();
    abortRef.current = controller;

    const patch = (fn: (t: Turn) => Turn) =>
      setTurns((prev) => prev.map((t, i) => (i === index ? fn(t) : t)));

    await api.ask(
      { courseId, question },
      (e: AskEvent) => {
        switch (e.type) {
          case "intent":
            patch((t) => ({ ...t, intent: e.intent }));
            break;
          case "sources":
            patch((t) => ({ ...t, sources: e.sources, questions: e.questions }));
            break;
          case "token":
            patch((t) => ({ ...t, text: t.text + e.text }));
            break;
          case "done":
            patch((t) => ({ ...t, streaming: false, usage: e.answer.usage }));
            break;
          case "error":
            patch((t) => ({ ...t, streaming: false, error: e.message }));
            break;
        }
      },
      controller.signal,
    );

    patch((t) => ({ ...t, streaming: false }));
  }

  return (
    <div className="animate-in flex min-h-[calc(100vh-9rem)] flex-col">
      <PageHeader
        title="Ask"
        subtitle="Ask by voice or text. Every answer cites the paper and page it came from."
      />

      {usingMocks && (
        <div className="mb-6 rounded-[var(--radius-card)] border border-dashed px-4 py-3">
          <p className="text-xs text-secondary">
            <strong className="font-semibold">Mock data.</strong> No model and no document index
            are connected yet, lookups and statistics are real queries over sample data, but
            explanations come from a small set of prepared answers and will say so when a question
            falls outside them.
          </p>
        </div>
      )}

      <div className="flex-1 space-y-8">
        {turns.length === 0 && (
          <div className="pt-4">
            <div className="mb-5 flex items-center gap-2 text-sm text-secondary">
              <Sparkles size={15} className="accent-text" />
              Try one of these
            </div>
            <div className="grid gap-2.5 sm:grid-cols-2">
              {SUGGESTIONS.map((s) => (
                <button
                  key={s}
                  onClick={() => send(s)}
                  className="rounded-[var(--radius-card)] border px-4 py-3.5 text-left text-[13.5px] leading-relaxed text-secondary transition hover:surface-sunken"
                >
                  {s}
                </button>
              ))}
            </div>
          </div>
        )}

        {turns.map((turn, i) => (
          <article key={i} className="animate-in">
            <h2 className="mb-4 text-[17px] font-semibold leading-snug tracking-[-0.01em]">
              {turn.question}
            </h2>

            {turn.intent && (
              <div className="mb-4 flex flex-wrap items-center gap-2">
                <span className="inline-flex items-center gap-1.5 rounded-full accent-soft-bg px-2.5 py-1 text-xs font-medium accent-text">
                  <Database size={12} />
                  {INTENT_COPY[turn.intent].label}
                </span>
                <span className="text-xs text-tertiary">{INTENT_COPY[turn.intent].detail}</span>
              </div>
            )}

            {turn.error ? (
              <Card className="border-red-500/25 bg-red-500/5 p-5">
                <p className="text-sm">{turn.error}</p>
              </Card>
            ) : (
              <div className="text-[15px] leading-[1.8] text-secondary">
                {turn.text.split("\n\n").map((para, pi) => (
                  <p
                    key={pi}
                    className={
                      pi === turn.text.split("\n\n").length - 1 && turn.streaming ? "caret" : ""
                    }
                    dangerouslySetInnerHTML={{ __html: renderInline(para) }}
                  />
                ))}
                {turn.streaming && !turn.text && (
                  <span className="inline-flex items-center gap-2 text-sm text-tertiary">
                    <Loader2 size={14} className="animate-spin" /> Retrieving…
                  </span>
                )}
              </div>
            )}

            {turn.questions && turn.questions.length > 0 && (
              <div className="mt-5 space-y-2.5">
                <p className="text-[13px] font-semibold uppercase tracking-wider text-tertiary">
                  {turn.questions.length} questions
                </p>
                {turn.questions.slice(0, 8).map((q) => (
                  <Link key={q.id} to={`/questions/${q.id}`} className="block">
                    <Card className="p-4 transition hover:surface-sunken">
                      <div className="mb-1.5 text-xs font-medium text-secondary">
                        {examLabel(q.year, q.term)} · Q{q.number} · Page {q.sourcePage}
                      </div>
                      <p className="line-clamp-2 text-[13.5px] leading-relaxed text-secondary">
                        {q.text}
                      </p>
                    </Card>
                  </Link>
                ))}
                {turn.questions.length > 8 && (
                  <Link
                    to={`/courses/${courseId}/questions`}
                    className="inline-block pt-1 text-sm font-medium accent-text hover:underline"
                  >
                    See all {turn.questions.length} →
                  </Link>
                )}
              </div>
            )}

            {turn.sources.length > 0 && (
              <div className="mt-6">
                <p className="mb-2.5 text-[13px] font-semibold uppercase tracking-wider text-tertiary">
                  Sources
                </p>
                <div className="divide-y rounded-[var(--radius-card)] border">
                  {turn.sources.map((s, si) => (
                    <a
                      key={si}
                      href={`/api/v1/documents/${s.documentId}/file#page=${s.page}`}
                      target="_blank"
                      rel="noreferrer"
                      className="flex items-center gap-3 px-4 py-2.5 text-sm transition hover:surface-sunken"
                    >
                      <FileText size={14} className="shrink-0 text-tertiary" />
                      <span className="text-secondary">{s.label}</span>
                    </a>
                  ))}
                </div>
              </div>
            )}

            {turn.usage && (
              <p className="mt-4 text-xs text-tertiary">
                {turn.usage.retrievedChunks} passages retrieved · {turn.usage.completionTokens}{" "}
                tokens · {turn.usage.latencyMs}ms · {turn.usage.provider}/{turn.usage.model}
              </p>
            )}
          </article>
        ))}
        <div ref={bottomRef} />
      </div>

      <div className="sticky bottom-0 -mx-1 mt-8 pb-6 pt-4">
        <div
          className="rounded-2xl border p-2 shadow-lg shadow-black/[0.04] dark:shadow-black/30"
          style={{ background: "var(--surface-raised)" }}
        >
          <form
            onSubmit={(e) => {
              e.preventDefault();
              send(input);
            }}
            className="flex items-end gap-2"
          >
            <textarea
              value={input}
              onChange={(e) => setInput(e.target.value)}
              onKeyDown={(e) => {
                if (e.key === "Enter" && !e.shiftKey) {
                  e.preventDefault();
                  send(input);
                }
              }}
              rows={1}
              placeholder={voice.listening ? "Listening…" : "Ask about this course…"}
              className="max-h-40 flex-1 resize-none bg-transparent px-3 py-2.5 text-[15px] outline-none placeholder:text-tertiary"
            />

            {voice.supported && (
              <button
                type="button"
                onClick={() => (voice.listening ? voice.stop() : voice.start())}
                disabled={busy}
                aria-label={voice.listening ? "Stop dictation" : "Ask by voice"}
                aria-pressed={voice.listening}
                title={voice.listening ? "Stop dictation" : "Ask by voice"}
                className={`flex h-9 w-9 items-center justify-center rounded-xl transition disabled:opacity-30 ${
                  voice.listening
                    ? "bg-red-500 text-white"
                    : "text-secondary hover:surface-sunken"
                }`}
              >
                <Mic size={16} className={voice.listening ? "animate-pulse" : ""} />
              </button>
            )}

            {busy ? (
              <button
                type="button"
                onClick={() => abortRef.current?.abort()}
                className="flex h-9 w-9 items-center justify-center rounded-xl surface-sunken transition hover:opacity-80"
                aria-label="Stop generating"
              >
                <Square size={14} />
              </button>
            ) : (
              <button
                type="submit"
                disabled={!input.trim()}
                className="flex h-9 w-9 items-center justify-center rounded-xl accent-bg text-white transition hover:opacity-90 disabled:opacity-30"
                aria-label="Send"
              >
                <ArrowUp size={16} />
              </button>
            )}
          </form>
        </div>

        {(voice.error || voice.listening) && (
          <p
            className={`mt-2 px-2 text-xs ${voice.error ? "text-red-500" : "text-tertiary"}`}
            role="status"
          >
            {voice.error ?? "Listening, speak your question, then press the mic again to stop."}
          </p>
        )}
        {!voice.supported && (
          <p className="mt-2 px-2 text-xs text-tertiary">
            Voice input needs Chrome, Edge or Safari.
          </p>
        )}
      </div>
    </div>
  );
}

/** Minimal inline markdown, bold only. Sources render as real links, not markdown. */
function renderInline(text: string): string {
  return text
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/\*\*(.+?)\*\*/g, "<strong class='font-semibold' style='color:var(--text-primary)'>$1</strong>")
    .replace(/\*(.+?)\*/g, "<em>$1</em>");
}
