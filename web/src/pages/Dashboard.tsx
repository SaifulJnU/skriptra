import { useState } from "react";
import { Link } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
import { ArrowUpRight, FileText, ListTree, Upload } from "lucide-react";
import { api } from "@/lib/client";
import { Button, Card, EmptyState, ErrorState, PageHeader, Skeleton } from "@/components/ui";
import UploadDialog from "@/components/UploadDialog";

export default function Dashboard() {
  const [uploadCourse, setUploadCourse] = useState<string | null>(null);
  const courses = useQuery({ queryKey: ["courses"], queryFn: () => api.listCourses() });
  const docs = useQuery({
    queryKey: ["documents", courses.data?.data[0]?.id],
    queryFn: () => api.listDocuments(courses.data!.data[0].id),
    enabled: !!courses.data?.data[0],
  });

  return (
    <div className="animate-in">
      {uploadCourse && (
        <UploadDialog courseId={uploadCourse} onClose={() => setUploadCourse(null)} />
      )}
      <PageHeader
        title="Your courses"
        subtitle="Upload past papers once, then search, compare and ask across every year."
        actions={
          <Button
            onClick={() => setUploadCourse(courses.data?.data[0]?.id ?? null)}
            disabled={!courses.data?.data.length}
            title={courses.data?.data.length ? undefined : "Create a course first"}
          >
            <Upload size={15} /> Upload paper
          </Button>
        }
      />

      {courses.isLoading && (
        <div className="grid gap-4 sm:grid-cols-2 xl:grid-cols-3">
          {[0, 1, 2].map((i) => (
            <Skeleton key={i} className="h-[152px]" />
          ))}
        </div>
      )}

      {courses.isError && <ErrorState error={courses.error} onRetry={() => courses.refetch()} />}

      {courses.data && courses.data.data.length === 0 && (
        <EmptyState
          title="No courses yet"
          description="Create a course and upload its past exam papers. Skriptra will split them into questions and index them by chapter."
          action={<Button>Create a course</Button>}
        />
      )}

      {courses.data && courses.data.data.length > 0 && (
        <div className="grid gap-4 sm:grid-cols-2 xl:grid-cols-3">
          {courses.data.data.map((c) => (
            <Link key={c.id} to={`/courses/${c.id}`} className="group">
              <Card className="h-full p-5 transition hover:shadow-lg hover:shadow-black/[0.04] dark:hover:shadow-black/20">
                <div className="flex items-start justify-between gap-3">
                  <div className="min-w-0">
                    <h2 className="truncate text-[15px] font-semibold">{c.name}</h2>
                    <p className="mt-0.5 text-xs text-tertiary">
                      {[c.code, c.institution].filter(Boolean).join(" · ")}
                    </p>
                  </div>
                  <ArrowUpRight
                    size={16}
                    className="shrink-0 text-tertiary transition group-hover:accent-text"
                  />
                </div>

                <div className="mt-6 flex items-center gap-6">
                  <div>
                    <div className="flex items-baseline gap-1.5">
                      <span className="text-[22px] font-semibold tabular-nums leading-none">
                        {c.examCount}
                      </span>
                      <FileText size={13} className="text-tertiary" />
                    </div>
                    <p className="mt-1 text-xs text-tertiary">exams</p>
                  </div>
                  <div>
                    <div className="flex items-baseline gap-1.5">
                      <span className="text-[22px] font-semibold tabular-nums leading-none">
                        {c.questionCount}
                      </span>
                      <ListTree size={13} className="text-tertiary" />
                    </div>
                    <p className="mt-1 text-xs text-tertiary">questions</p>
                  </div>
                </div>
              </Card>
            </Link>
          ))}
        </div>
      )}

      <section className="mt-12">
        <h2 className="mb-4 text-[13px] font-semibold uppercase tracking-wider text-tertiary">
          Recent activity
        </h2>
        {docs.isLoading && <Skeleton className="h-40" />}
        {docs.data && (
          <Card className="divide-y overflow-hidden">
            {docs.data.data.slice(0, 5).map((d) => (
              <div key={d.id} className="flex items-center gap-4 px-5 py-3.5">
                <FileText size={16} className="shrink-0 text-tertiary" />
                <div className="min-w-0 flex-1">
                  <p className="truncate text-sm font-medium">{d.filename}</p>
                  <p className="mt-0.5 text-xs text-tertiary capitalize">
                    {d.kind}
                    {d.pageCount ? ` · ${d.pageCount} pages` : ""}
                  </p>
                </div>
                {d.status === "indexed" ? (
                  <span className="text-xs text-tertiary">indexed</span>
                ) : (
                  <span className="flex items-center gap-2 text-xs accent-text">
                    <span className="h-1.5 w-1.5 animate-pulse rounded-full accent-bg" />
                    {d.status}
                  </span>
                )}
              </div>
            ))}
          </Card>
        )}
      </section>
    </div>
  );
}
