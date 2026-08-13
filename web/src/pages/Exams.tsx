import { Link, useParams } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
import { CheckCircle2, FileText, MinusCircle } from "lucide-react";
import { api } from "@/lib/client";
import { Card, ErrorState, PageHeader, Skeleton } from "@/components/ui";
import { termLabel } from "@/lib/utils";
import type { Exam } from "@/types/api";

export default function Exams() {
  const { courseId = "" } = useParams();
  const exams = useQuery({
    queryKey: ["exams", courseId],
    queryFn: () => api.listExams(courseId),
  });

  // Grouped by year so the browser reads the way a student thinks about it.
  const byYear = (exams.data?.data ?? []).reduce<Record<number, Exam[]>>((acc, e) => {
    (acc[e.year] ??= []).push(e);
    return acc;
  }, {});

  return (
    <div className="animate-in">
      <PageHeader
        title="Exams"
        subtitle={
          exams.data
            ? `${exams.data.meta.total} papers indexed`
            : "Past papers, newest first"
        }
      />

      {exams.isLoading && (
        <div className="space-y-3">
          {[0, 1, 2].map((i) => (
            <Skeleton key={i} className="h-24" />
          ))}
        </div>
      )}
      {exams.isError && <ErrorState error={exams.error} onRetry={() => exams.refetch()} />}

      <div className="space-y-8">
        {Object.entries(byYear)
          .sort(([a], [b]) => Number(b) - Number(a))
          .map(([year, list]) => (
            <section key={year}>
              <h2 className="mb-3 flex items-baseline gap-3">
                <span className="text-lg font-semibold tabular-nums">{year}</span>
                <span className="text-xs text-tertiary">
                  {list.reduce((n, e) => n + e.questionCount, 0)} questions
                </span>
              </h2>
              <div className="grid gap-3 sm:grid-cols-2">
                {list.map((e) => (
                  <Link key={e.id} to={`/courses/${courseId}/questions?year=${e.year}`}>
                    <Card className="flex items-center gap-4 p-4 transition hover:shadow-md hover:shadow-black/[0.04] dark:hover:shadow-black/20">
                      <span className="flex h-10 w-10 shrink-0 items-center justify-center rounded-lg surface-sunken">
                        <FileText size={17} className="text-secondary" />
                      </span>
                      <div className="min-w-0 flex-1">
                        <p className="text-sm font-medium">
                          {e.year} {termLabel(e.term)}
                        </p>
                        <p className="mt-0.5 text-xs text-tertiary">
                          {e.questionCount} questions
                        </p>
                      </div>
                      {e.hasSolutions ? (
                        <span className="flex items-center gap-1.5 text-xs text-emerald-600 dark:text-emerald-400">
                          <CheckCircle2 size={13} /> solutions
                        </span>
                      ) : (
                        <span className="flex items-center gap-1.5 text-xs text-tertiary">
                          <MinusCircle size={13} /> no solutions
                        </span>
                      )}
                    </Card>
                  </Link>
                ))}
              </div>
            </section>
          ))}
      </div>
    </div>
  );
}
