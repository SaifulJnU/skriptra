import { useRef, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
import { AlertTriangle, ArrowLeft, FileText, Lightbulb, Loader2, Sparkles, Square } from "lucide-react";
import { api } from "@/lib/client";
import type { AskEvent } from "@/lib/api";
import { documentFileURL } from "@/lib/api";
import type { Citation } from "@/types/api";
import { Button, Card, ChapterBadge, ErrorState, Skeleton } from "@/components/ui";
import { Markdown } from "@/components/Markdown";
import { examLabel, formatPercent } from "@/lib/utils";

/**
 * Offers a generated worked solution when the course has not published one.
 *
 * Presented as clearly distinct from an official solution throughout: a
 * different heading, a standing caveat, and citations to the material it was
 * grounded in. A student revising from this will act on it in an exam, so the
 * distinction has to survive skim-reading.
 */
function GeneratedSolution({ questionId }: { questionId: string }) {
  const [text, setText] = useState("");
  const [sources, setSources] = useState<Citation[]>([]);
  const [streaming, setStreaming] = useState(false);
  const [error, setError] = useState<string>();
  const [started, setStarted] = useState(false);
  const abortRef = useRef<AbortController | null>(null);

  async function generate() {
    setStarted(true);
    setStreaming(true);
    setError(undefined);
    setText("");
    setSources([]);

    const controller = new AbortController();
    abortRef.current = controller;

    await api.generateSolution(
      questionId,
      (e: AskEvent) => {
        switch (e.type) {
          case "sources":
            setSources(e.sources);
            break;
          case "token":
            setText((prev) => prev + e.text);
            break;
          case "error":
            // A failed generation grounds nothing, so its citations go too.
            setError(e.message);
            setSources([]);
            setText("");
            break;
        }
      },
      controller.signal,
    );
    setStreaming(false);
  }

  if (!started) {
    return (
      <Card className="mt-6 border-dashed p-5">
        <div className="flex flex-wrap items-center justify-between gap-4">
          <div>
            <p className="text-sm font-medium">No official solution for this paper</p>
            <p className="mt-1 text-sm text-secondary">
              Skriptra can work through it using this course&apos;s material.
            </p>
          </div>
          <Button onClick={generate}>
            <Sparkles size={15} /> Generate a solution
          </Button>
        </div>
      </Card>
    );
  }

  return (
    <section className="mt-6">
      <h2 className="mb-3 flex items-center gap-2 text-[13px] font-semibold uppercase tracking-wider text-tertiary">
        <Sparkles size={14} /> Generated solution
      </h2>

      <Card className="p-6">
        <div className="mb-4 flex items-start gap-2.5 rounded-lg bg-amber-500/10 px-3.5 py-2.5">
          <AlertTriangle size={15} className="mt-0.5 shrink-0 text-amber-600 dark:text-amber-400" />
          <p className="text-xs leading-relaxed text-secondary">
            Generated from course material, not an official mark scheme. Check it against your
            lecture notes before relying on it.
          </p>
        </div>

        {error ? (
          <p className="text-sm text-red-500">{error}</p>
        ) : text ? (
          <Markdown className="text-[14.5px] leading-[1.75] text-secondary">{text}</Markdown>
        ) : (
          streaming && (
            <span className="inline-flex items-center gap-2 text-sm text-tertiary">
              <Loader2 size={14} className="animate-spin" /> Working through it...
            </span>
          )
        )}

        {sources.length > 0 && (
          <div className="mt-5 border-t pt-4">
            <p className="mb-2 text-[11px] font-semibold uppercase tracking-wider text-tertiary">
              Grounded in
            </p>
            <div className="space-y-1">
              {sources.map((s, i) => (
                <p key={i} className="text-xs text-tertiary">
                  {s.label}
                </p>
              ))}
            </div>
          </div>
        )}

        <div className="mt-5 flex gap-2">
          {streaming ? (
            <Button variant="outline" size="sm" onClick={() => abortRef.current?.abort()}>
              <Square size={13} /> Stop
            </Button>
          ) : (
            <Button variant="outline" size="sm" onClick={generate}>
              Regenerate
            </Button>
          )}
        </div>
      </Card>
    </section>
  );
}

/**
 * The question viewer: the screen that carries the product.
 *
 * A question is only useful in context: which chapter it belongs to, where it
 * came from, whether it has been asked before. "Similar questions" is the
 * feature that turns a folder of PDFs into a study tool, so it sits beside the
 * question rather than behind a tab.
 */
export default function QuestionViewer() {
  const { questionId = "" } = useParams();

  const question = useQuery({
    queryKey: ["question", questionId],
    queryFn: () => api.getQuestion(questionId),
  });

  const similar = useQuery({
    queryKey: ["similar", questionId],
    queryFn: () => api.similarQuestions(questionId, 6),
    enabled: !!question.data,
  });

  if (question.isLoading) {
    return (
      <div className="space-y-4">
        <Skeleton className="h-8 w-56" />
        <Skeleton className="h-44" />
        <Skeleton className="h-64" />
      </div>
    );
  }

  if (question.isError) {
    return <ErrorState error={question.error} onRetry={() => question.refetch()} />;
  }

  const q = question.data!;

  return (
    <div className="animate-in">
      <button
        onClick={() => history.back()}
        className="mb-5 inline-flex items-center gap-1.5 text-sm text-secondary transition hover:accent-text"
      >
        <ArrowLeft size={15} /> Back
      </button>

      <div className="grid gap-8 lg:grid-cols-[minmax(0,1fr)_300px]">
        <div className="min-w-0">
          <div className="mb-4 flex flex-wrap items-center gap-3">
            <h1 className="text-xl font-semibold tracking-[-0.01em]">
              {examLabel(q.year, q.term)} Exam · Question {q.number}
            </h1>
            <ChapterBadge chapter={q.chapter} />
          </div>

          <Card className="p-6">
            <p className="text-[15.5px] leading-[1.75]">{q.text}</p>
            {q.marks && (
              <p className="mt-5 border-t pt-4 text-xs text-tertiary">{q.marks} marks</p>
            )}
          </Card>

          {q.solutionText ? (
            <section className="mt-6">
              <h2 className="mb-3 flex items-center gap-2 text-[13px] font-semibold uppercase tracking-wider text-tertiary">
                <Lightbulb size={14} /> Worked solution
              </h2>
              <Card className="p-6">
                <p className="text-[14.5px] leading-[1.75] text-secondary">{q.solutionText}</p>
                {q.solutionSourcePage && (
                  <p className="mt-5 border-t pt-4 text-xs text-tertiary">
                    Source: solution sheet, page {q.solutionSourcePage}
                  </p>
                )}
              </Card>
            </section>
          ) : (
            <GeneratedSolution questionId={q.id} />
          )}

          <section className="mt-10">
            <h2 className="mb-3 text-[13px] font-semibold uppercase tracking-wider text-tertiary">
              Similar questions
            </h2>

            {similar.isLoading && (
              <div className="space-y-2.5">
                {[0, 1, 2].map((i) => (
                  <Skeleton key={i} className="h-20" />
                ))}
              </div>
            )}

            {similar.data?.length === 0 && (
              <Card className="border-dashed p-5">
                <p className="text-sm text-tertiary">
                  Nothing closely related found in the indexed years.
                </p>
              </Card>
            )}

            <div className="space-y-2.5">
              {similar.data?.map(({ question: s, score }) => (
                <Link key={s.id} to={`/questions/${s.id}`} className="block">
                  <Card className="p-4 transition hover:shadow-md hover:shadow-black/[0.04] dark:hover:shadow-black/20">
                    <div className="mb-1.5 flex items-center gap-3 text-xs">
                      <span className="font-medium text-secondary">
                        {examLabel(s.year, s.term)} · Q{s.number}
                      </span>
                      <span
                        className="ml-auto rounded-full px-2 py-0.5 font-medium tabular-nums accent-soft-bg accent-text"
                        title="Cosine similarity of question embeddings"
                      >
                        {formatPercent(score)}
                      </span>
                    </div>
                    <p className="line-clamp-2 text-[13.5px] leading-relaxed text-secondary">
                      {s.text}
                    </p>
                  </Card>
                </Link>
              ))}
            </div>
          </section>
        </div>

        <aside className="space-y-5">
          <Card className="p-5">
            <h3 className="mb-4 text-[13px] font-semibold uppercase tracking-wider text-tertiary">
              Details
            </h3>
            <dl className="space-y-3.5 text-sm">
              <div>
                <dt className="text-xs text-tertiary">Chapter</dt>
                <dd className="mt-0.5">
                  {q.chapter ? `${q.chapter.number}. ${q.chapter.title}` : "Unclassified"}
                </dd>
              </div>
              {q.topic && (
                <div>
                  <dt className="text-xs text-tertiary">Topic</dt>
                  <dd className="mt-0.5">{q.topic}</dd>
                </div>
              )}
              <div>
                <dt className="text-xs text-tertiary">Source page</dt>
                <dd className="mt-0.5">Page {q.sourcePage}</dd>
              </div>
              {q.chapter?.confidence !== undefined && (
                <div>
                  <dt className="text-xs text-tertiary">Classification confidence</dt>
                  <dd className="mt-1">
                    <div className="flex items-center gap-2">
                      <div className="h-1.5 flex-1 overflow-hidden rounded-full surface-sunken">
                        <div
                          className="h-full rounded-full accent-bg"
                          style={{ width: `${q.chapter.confidence * 100}%` }}
                        />
                      </div>
                      <span className="text-xs tabular-nums text-tertiary">
                        {formatPercent(q.chapter.confidence)}
                      </span>
                    </div>
                  </dd>
                </div>
              )}
            </dl>
          </Card>

          <a
            href={q.documentId ? documentFileURL(q.documentId, q.sourcePage) : "#"}
            target="_blank"
            rel="noreferrer"
            className="flex items-center gap-3 rounded-[var(--radius-card)] border p-4 transition hover:surface-sunken"
          >
            <FileText size={17} className="shrink-0 text-tertiary" />
            <div className="min-w-0">
              <p className="text-sm font-medium">Open source PDF</p>
              <p className="mt-0.5 text-xs text-tertiary">Jumps to page {q.sourcePage}</p>
            </div>
          </a>
        </aside>
      </div>
    </div>
  );
}
