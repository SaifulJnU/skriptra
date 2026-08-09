import { useParams } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
import {
  Bar,
  BarChart,
  CartesianGrid,
  Cell,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from "recharts";
import { api } from "@/lib/client";
import { Card, ErrorState, PageHeader, Skeleton } from "@/components/ui";
import { formatPercent } from "@/lib/utils";

/**
 * The `analyse` intent, rendered.
 *
 * These numbers come from a SQL aggregate, not from a language model — they are
 * exact and instant. That distinction is the point of the whole architecture,
 * and it is why this page can state "187 questions" rather than "roughly 190".
 */
export default function Analytics() {
  const { courseId = "" } = useParams();
  const freq = useQuery({
    queryKey: ["chapter-frequency", courseId],
    queryFn: () => api.chapterFrequency(courseId),
  });

  if (freq.isLoading) {
    return (
      <div className="space-y-4">
        <Skeleton className="h-9 w-64" />
        <Skeleton className="h-80" />
      </div>
    );
  }
  if (freq.isError) return <ErrorState error={freq.error} onRetry={() => freq.refetch()} />;

  const data = freq.data!;
  const rows = [...data.data].sort((a, b) => b.questionCount - a.questionCount);
  const top = rows[0];
  const chartData = rows.map((r) => ({
    name: `Ch ${r.chapter.number}`,
    full: r.chapter.title,
    questions: r.questionCount,
    share: r.share,
  }));

  return (
    <div className="animate-in">
      <PageHeader
        title="Exam analytics"
        subtitle={`${data.totalQuestions} questions indexed across all papers`}
      />

      <Card className="mb-6 p-6">
        <p className="text-sm text-secondary">
          <strong className="font-semibold" style={{ color: "var(--text-primary)" }}>
            Chapter {top.chapter.number} — {top.chapter.title}
          </strong>{" "}
          is the most heavily tested material, accounting for{" "}
          <strong className="font-semibold tabular-nums" style={{ color: "var(--text-primary)" }}>
            {formatPercent(top.share)}
          </strong>{" "}
          of every indexed question and appearing in {top.examCount} separate exams. If revision
          time is short, this is where it goes.
        </p>
      </Card>

      <Card className="p-6">
        <h2 className="mb-6 text-[13px] font-semibold uppercase tracking-wider text-tertiary">
          Most frequently tested chapters
        </h2>
        <div className="h-[300px] w-full">
          <ResponsiveContainer width="100%" height="100%">
            <BarChart data={chartData} margin={{ top: 4, right: 8, left: -16, bottom: 4 }}>
              <CartesianGrid strokeDasharray="3 3" stroke="var(--border-subtle)" vertical={false} />
              <XAxis
                dataKey="name"
                tick={{ fontSize: 12, fill: "var(--text-tertiary)" }}
                axisLine={{ stroke: "var(--border-subtle)" }}
                tickLine={false}
              />
              <YAxis
                tick={{ fontSize: 12, fill: "var(--text-tertiary)" }}
                axisLine={false}
                tickLine={false}
              />
              <Tooltip
                cursor={{ fill: "var(--surface-sunken)" }}
                contentStyle={{
                  background: "var(--surface-raised)",
                  border: "1px solid var(--border-subtle)",
                  borderRadius: 12,
                  fontSize: 13,
                  color: "var(--text-primary)",
                }}
                formatter={(value: number) => [`${value} questions`, ""]}
                labelFormatter={(_, payload) => payload?.[0]?.payload?.full ?? ""}
              />
              <Bar dataKey="questions" radius={[6, 6, 0, 0]}>
                {chartData.map((_, i) => (
                  <Cell key={i} fill={i === 0 ? "var(--accent)" : "var(--border-subtle)"} />
                ))}
              </Bar>
            </BarChart>
          </ResponsiveContainer>
        </div>
      </Card>

      <Card className="mt-6 overflow-hidden">
        <table className="w-full text-sm">
          <thead>
            <tr className="border-b text-left text-xs uppercase tracking-wider text-tertiary">
              <th className="px-5 py-3 font-semibold">Chapter</th>
              <th className="px-5 py-3 text-right font-semibold">Questions</th>
              <th className="px-5 py-3 text-right font-semibold">Share</th>
              <th className="px-5 py-3 text-right font-semibold">Exams</th>
            </tr>
          </thead>
          <tbody className="divide-y">
            {rows.map((r) => (
              <tr key={r.chapter.id}>
                <td className="px-5 py-3">
                  <span className="mr-2 font-mono text-xs text-tertiary">
                    {String(r.chapter.number).padStart(2, "0")}
                  </span>
                  {r.chapter.title}
                </td>
                <td className="px-5 py-3 text-right tabular-nums">{r.questionCount}</td>
                <td className="px-5 py-3 text-right tabular-nums text-secondary">
                  {formatPercent(r.share, 1)}
                </td>
                <td className="px-5 py-3 text-right tabular-nums text-secondary">{r.examCount}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </Card>
    </div>
  );
}
