import { Link, useParams, useSearchParams } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
import { api } from "@/lib/client";
import {
  Button,
  Card,
  ChapterBadge,
  EmptyState,
  ErrorState,
  PageHeader,
  Skeleton,
} from "@/components/ui";
import { examLabel } from "@/lib/utils";

/**
 * The `enumerate` surface. Filters live in the URL so a filtered list is
 * shareable and survives a refresh, students send each other links like
 * "chapter 3, last three years".
 */
export default function Questions() {
  const { courseId = "" } = useParams();
  const [params, setParams] = useSearchParams();

  const chapterNumber = params.get("chapter") ? Number(params.get("chapter")) : undefined;
  const year = params.get("year") ? Number(params.get("year")) : undefined;
  const sort = (params.get("sort") ?? "newest") as "newest" | "oldest" | "chapter";

  const chapters = useQuery({
    queryKey: ["chapters", courseId],
    queryFn: () => api.listChapters(courseId),
  });

  const questions = useQuery({
    queryKey: ["questions", courseId, chapterNumber, year, sort],
    queryFn: () =>
      api.listQuestions(courseId, {
        chapterNumber,
        yearFrom: year,
        yearTo: year,
        sort,
        pageSize: 50,
      }),
  });

  function setParam(key: string, value?: string) {
    const next = new URLSearchParams(params);
    if (value === undefined) next.delete(key);
    else next.set(key, value);
    setParams(next, { replace: true });
  }

  const activeChapter = chapters.data?.find((c) => c.number === chapterNumber);

  return (
    <div className="animate-in">
      <PageHeader
        title={activeChapter ? `Chapter ${activeChapter.number}, ${activeChapter.title}` : "Questions"}
        subtitle={
          questions.data
            ? `${questions.data.meta.total} questions${year ? ` from ${year}` : " across all years"}`
            : undefined
        }
        actions={
          <Link to={`/courses/${courseId}/ask`}>
            <Button variant="outline" size="sm">
              Ask about these
            </Button>
          </Link>
        }
      />

      <div className="mb-6 flex flex-wrap items-center gap-2">
        <button
          onClick={() => setParam("chapter", undefined)}
          className={`rounded-full border px-3 py-1.5 text-[13px] transition ${
            chapterNumber === undefined ? "accent-bg border-transparent text-white" : "hover:surface-sunken"
          }`}
        >
          All chapters
        </button>
        {chapters.data?.map((c) => (
          <button
            key={c.id}
            onClick={() => setParam("chapter", String(c.number))}
            className={`rounded-full border px-3 py-1.5 text-[13px] transition ${
              chapterNumber === c.number
                ? "accent-bg border-transparent text-white"
                : "hover:surface-sunken"
            }`}
          >
            {c.number}. {c.title}
          </button>
        ))}

        <select
          value={sort}
          onChange={(e) => setParam("sort", e.target.value)}
          className="ml-auto rounded-lg border bg-transparent px-2.5 py-1.5 text-[13px]"
        >
          <option value="newest">Newest first</option>
          <option value="oldest">Oldest first</option>
          <option value="chapter">By chapter</option>
        </select>
      </div>

      {questions.isLoading && (
        <div className="space-y-3">
          {[0, 1, 2, 3].map((i) => (
            <Skeleton key={i} className="h-28" />
          ))}
        </div>
      )}
      {questions.isError && (
        <ErrorState error={questions.error} onRetry={() => questions.refetch()} />
      )}

      {questions.data?.data.length === 0 && (
        <EmptyState
          title="No questions match these filters"
          description="Try a different chapter or clear the year filter."
          action={
            <Button variant="outline" onClick={() => setParams(new URLSearchParams())}>
              Clear filters
            </Button>
          }
        />
      )}

      <div className="space-y-3">
        {questions.data?.data.map((q) => (
          <Link key={q.id} to={`/questions/${q.id}`} className="block">
            <Card className="p-5 transition hover:shadow-md hover:shadow-black/[0.04] dark:hover:shadow-black/20">
              <div className="mb-2.5 flex flex-wrap items-center gap-2.5 text-xs text-tertiary">
                <span className="font-medium text-secondary">
                  {examLabel(q.year, q.term)} · Q{q.number}
                </span>
                <span>Page {q.sourcePage}</span>
                {q.marks && <span>{q.marks} marks</span>}
                <span className="ml-auto">
                  <ChapterBadge chapter={q.chapter} />
                </span>
              </div>
              <p className="text-[14.5px] leading-relaxed text-secondary">{q.text}</p>
            </Card>
          </Link>
        ))}
      </div>
    </div>
  );
}
