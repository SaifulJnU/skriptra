import { useEffect, useState } from "react";
import { Link, NavLink, Outlet, useMatch, useParams } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
import {
  BarChart3,
  BookOpen,
  Cpu,
  FileText,
  LayoutGrid,
  ListTree,
  Moon,
  Sparkles,
  Sun,
} from "lucide-react";
import { api, usingMocks } from "@/lib/client";
import { cn } from "@/lib/utils";

function ThemeToggle() {
  const [dark, setDark] = useState(() =>
    document.documentElement.classList.contains("dark"),
  );
  useEffect(() => {
    document.documentElement.classList.toggle("dark", dark);
  }, [dark]);
  return (
    <button
      onClick={() => setDark((d) => !d)}
      aria-label={dark ? "Switch to light theme" : "Switch to dark theme"}
      className="rounded-lg p-2 text-secondary transition hover:surface-sunken"
    >
      {dark ? <Sun size={17} /> : <Moon size={17} />}
    </button>
  );
}

/**
 * Shows which model is actually serving this deployment.
 *
 * Worth the header space: a student self-hosting with Ollama can see at a
 * glance that nothing is leaving their machine, and it makes the provider
 * abstraction visible rather than merely claimed.
 */
function ProviderPill() {
  const { data } = useQuery({ queryKey: ["providers"], queryFn: () => api.providers() });
  if (!data) return null;
  return (
    <div className="hidden items-center gap-1.5 rounded-full border px-2.5 py-1 text-xs text-secondary sm:flex">
      <Cpu size={13} className={data.llm.local ? "text-emerald-500" : "text-amber-500"} />
      <span className="font-medium">{data.llm.model}</span>
      <span className="text-tertiary">{data.llm.local ? "local" : "hosted"}</span>
    </div>
  );
}

const navItem = (active: boolean) =>
  cn(
    "flex items-center gap-2.5 rounded-lg px-3 py-2 text-sm transition",
    active ? "accent-soft-bg accent-text font-medium" : "text-secondary hover:surface-sunken",
  );

function CourseNav({ courseId }: { courseId: string }) {
  const { data: course } = useQuery({
    queryKey: ["course", courseId],
    queryFn: () => api.getCourse(courseId),
  });
  const { data: chapters } = useQuery({
    queryKey: ["chapters", courseId],
    queryFn: () => api.listChapters(courseId),
  });

  const links = [
    { to: `/courses/${courseId}`, label: "Overview", icon: LayoutGrid, end: true },
    { to: `/courses/${courseId}/exams`, label: "Exams", icon: FileText },
    { to: `/courses/${courseId}/questions`, label: "Questions", icon: ListTree },
    { to: `/courses/${courseId}/ask`, label: "Ask", icon: Sparkles },
    { to: `/courses/${courseId}/analytics`, label: "Analytics", icon: BarChart3 },
  ];

  return (
    <>
      <div className="px-3 pb-3">
        <p className="truncate text-sm font-semibold">{course?.name ?? "Loading…"}</p>
        <p className="mt-0.5 text-xs text-tertiary">
          {course ? `${course.examCount} exams · ${course.questionCount} questions` : " "}
        </p>
      </div>

      <nav className="space-y-0.5">
        {links.map(({ to, label, icon: Icon, end }) => (
          <NavLink key={to} to={to} end={end} className={({ isActive }) => navItem(isActive)}>
            <Icon size={16} strokeWidth={1.8} />
            {label}
          </NavLink>
        ))}
      </nav>

      {chapters && chapters.length > 0 && (
        <div className="mt-7">
          <p className="px-3 pb-2 text-[11px] font-semibold uppercase tracking-wider text-tertiary">
            Chapters
          </p>
          <div className="space-y-0.5">
            {chapters.map((c) => (
              <Link
                key={c.id}
                to={`/courses/${courseId}/questions?chapter=${c.number}`}
                className="flex items-center gap-2.5 rounded-lg px-3 py-1.5 text-[13px] text-secondary transition hover:surface-sunken"
              >
                <span className="w-4 shrink-0 font-mono text-[11px] text-tertiary">
                  {String(c.number).padStart(2, "0")}
                </span>
                <span className="truncate">{c.title}</span>
                <span className="ml-auto text-[11px] text-tertiary">{c.questionCount}</span>
              </Link>
            ))}
          </div>
        </div>
      )}
    </>
  );
}

export default function AppShell() {
  const { courseId } = useParams();
  const questionMatch = useMatch("/questions/:questionId");
  // The question viewer is reached from a course but has no courseId in its
  // path; fall back to the first course so the sidebar does not blank out.
  const { data: courses } = useQuery({ queryKey: ["courses"], queryFn: () => api.listCourses() });
  const activeCourseId = courseId ?? (questionMatch ? courses?.data[0]?.id : undefined);

  return (
    <div className="min-h-screen">
      <header className="sticky top-0 z-30 border-b backdrop-blur-xl" style={{ background: "color-mix(in oklch, var(--surface) 82%, transparent)" }}>
        <div className="mx-auto flex h-14 max-w-[1400px] items-center gap-4 px-5">
          <Link to="/" className="flex items-center gap-2.5">
            <span className="flex h-7 w-7 items-center justify-center rounded-lg accent-bg text-white">
              <BookOpen size={15} strokeWidth={2.2} />
            </span>
            <span className="text-[15px] font-semibold tracking-[-0.01em]">Lernova</span>
          </Link>

          <div className="ml-auto flex items-center gap-2">
            {usingMocks && (
              <span className="hidden rounded-full border border-dashed px-2.5 py-1 text-xs text-tertiary md:inline">
                mock data
              </span>
            )}
            <ProviderPill />
            <ThemeToggle />
            <div className="ml-1 flex h-7 w-7 items-center justify-center rounded-full surface-sunken text-xs font-semibold">
              S
            </div>
          </div>
        </div>
      </header>

      <div className="mx-auto flex max-w-[1400px] gap-8 px-5 py-8">
        {activeCourseId && (
          <aside className="hidden w-60 shrink-0 lg:block">
            <div className="sticky top-[5.5rem]">
              <CourseNav courseId={activeCourseId} />
            </div>
          </aside>
        )}
        <main className="min-w-0 flex-1">
          <Outlet />
        </main>
      </div>
    </div>
  );
}
