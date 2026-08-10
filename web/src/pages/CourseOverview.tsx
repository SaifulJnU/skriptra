import { useState } from "react";
import { Link, useParams } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
import { ArrowRight, FileText, ListTree, Sparkles, Upload } from "lucide-react";
import { api } from "@/lib/client";
import { Button, Card, ErrorState, PageHeader, Skeleton } from "@/components/ui";
import UploadDialog from "@/components/UploadDialog";
import { formatPercent } from "@/lib/utils";

export default function CourseOverview() {
  const { courseId = "" } = useParams();
  const [uploading, setUploading] = useState(false);

  const course = useQuery({
    queryKey: ["course", courseId],
    queryFn: () => api.getCourse(courseId),
  });
  const chapters = useQuery({
    queryKey: ["chapters", courseId],
    queryFn: () => api.listChapters(courseId),
  });
  const freq = useQuery({
    queryKey: ["chapter-frequency", courseId],
    queryFn: () => api.chapterFrequency(courseId),
  });
  const docs = useQuery({
    queryKey: ["documents", courseId],
    queryFn: () => api.listDocuments(courseId),
  });

  if (course.isError) return <ErrorState error={course.error} onRetry={() => course.refetch()} />;

  const ingesting = docs.data?.data.filter((d) => d.status !== "indexed" && d.status !== "failed");
  const top = freq.data ? [...freq.data.data].sort((a, b) => b.questionCount - a.questionCount)[0] : undefined;

  return (
    <div className="animate-in">
      {uploading && <UploadDialog courseId={courseId} onClose={() => setUploading(false)} />}
      <PageHeader
        title={course.data?.name ?? <Skeleton className="h-8 w-52" />}
        subtitle={
          course.data
            ? [
                course.data.code,
                course.data.institution,
                course.data.yearRange &&
                  `${course.data.yearRange.from}, ${course.data.yearRange.to}`,
              ]
                .filter(Boolean)
                .join(" · ")
            : undefined
        }
        actions={
          <>
            <Button variant="outline" onClick={() => setUploading(true)}>
              <Upload size={15} /> Upload
            </Button>
            <Link to={`/courses/${courseId}/ask`}>
              <Button>
                <Sparkles size={15} /> Ask
              </Button>
            </Link>
          </>
        }
      />

      {ingesting && ingesting.length > 0 && (
        <Card className="mb-6 flex items-center gap-4 p-4">
          <span className="h-2 w-2 shrink-0 animate-pulse rounded-full accent-bg" />
          <div className="min-w-0 flex-1">
            <p className="text-sm font-medium">
              Indexing {ingesting.length} document{ingesting.length > 1 ? "s" : ""}
            </p>
            <p className="mt-0.5 truncate text-xs text-tertiary">
              {ingesting.map((d) => d.filename).join(", ")}, questions become searchable as each
              finishes.
            </p>
          </div>
        </Card>
      )}

      <div className="mb-8 grid gap-4 sm:grid-cols-3">
        <Stat label="Exams indexed" value={course.data?.examCount} icon={<FileText size={14} />} />
        <Stat label="Questions" value={course.data?.questionCount} icon={<ListTree size={14} />} />
        <Stat label="Chapters" value={course.data?.chapterCount} />
      </div>

      {top && (
        <Card className="mb-8 p-5">
          <p className="text-sm text-secondary">
            Most tested:{" "}
            <strong className="font-semibold" style={{ color: "var(--text-primary)" }}>
              Chapter {top.chapter.number}, {top.chapter.title}
            </strong>{" "}
            at {formatPercent(top.share)} of all questions.{" "}
            <Link
              to={`/courses/${courseId}/analytics`}
              className="font-medium accent-text hover:underline"
            >
              See full breakdown
            </Link>
          </p>
        </Card>
      )}

      <section>
        <div className="mb-4 flex items-center justify-between">
          <h2 className="text-[13px] font-semibold uppercase tracking-wider text-tertiary">
            Chapters
          </h2>
          <Link
            to={`/courses/${courseId}/questions`}
            className="flex items-center gap-1 text-sm text-secondary transition hover:accent-text"
          >
            All questions <ArrowRight size={14} />
          </Link>
        </div>

        {chapters.isLoading && (
          <div className="space-y-2.5">
            {[0, 1, 2, 3].map((i) => (
              <Skeleton key={i} className="h-16" />
            ))}
          </div>
        )}

        <div className="space-y-2.5">
          {chapters.data?.map((c) => (
            <Link key={c.id} to={`/courses/${courseId}/questions?chapter=${c.number}`}>
              <Card className="flex items-center gap-4 p-4 transition hover:shadow-md hover:shadow-black/[0.04] dark:hover:shadow-black/20">
                <span className="w-8 shrink-0 font-mono text-sm text-tertiary">
                  {String(c.number).padStart(2, "0")}
                </span>
                <div className="min-w-0 flex-1">
                  <p className="text-sm font-medium">{c.title}</p>
                  {c.topics && c.topics.length > 0 && (
                    <p className="mt-0.5 truncate text-xs text-tertiary">{c.topics.join(" · ")}</p>
                  )}
                </div>
                <span className="shrink-0 text-sm tabular-nums text-tertiary">
                  {c.questionCount}
                </span>
                <ArrowRight size={15} className="shrink-0 text-tertiary" />
              </Card>
            </Link>
          ))}
        </div>
      </section>
    </div>
  );
}

function Stat({
  label,
  value,
  icon,
}: {
  label: string;
  value?: number;
  icon?: React.ReactNode;
}) {
  return (
    <Card className="p-5">
      <div className="flex items-center gap-1.5 text-xs text-tertiary">
        {icon}
        {label}
      </div>
      <p className="mt-2 text-[26px] font-semibold leading-none tabular-nums">
        {value ?? <span className="shimmer inline-block h-6 w-12 rounded" />}
      </p>
    </Card>
  );
}
