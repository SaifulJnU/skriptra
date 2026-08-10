/**
 * Small presentational primitives.
 *
 * Loading, empty and error states are first-class components rather than
 * inline afterthoughts, they are the states a user actually hits on a slow
 * connection or an empty course, and designing them once here keeps every page
 * honest about them.
 */
import type { ReactNode } from "react";
import { AlertCircle, Inbox } from "lucide-react";
import { cn, LOW_CONFIDENCE } from "@/lib/utils";
import type { ChapterRef } from "@/types/api";

export function Card({ className, children }: { className?: string; children: ReactNode }) {
  return <div className={cn("card", className)}>{children}</div>;
}

export function Badge({
  children,
  tone = "neutral",
  className,
}: {
  children: ReactNode;
  tone?: "neutral" | "accent" | "warn" | "success";
  className?: string;
}) {
  const tones = {
    neutral: "surface-sunken text-secondary",
    accent: "accent-soft-bg accent-text",
    warn: "bg-amber-500/12 text-amber-700 dark:text-amber-400",
    success: "bg-emerald-500/12 text-emerald-700 dark:text-emerald-400",
  };
  return (
    <span
      className={cn(
        "inline-flex items-center gap-1 rounded-full px-2 py-0.5 text-xs font-medium",
        tones[tone],
        className,
      )}
    >
      {children}
    </span>
  );
}

/**
 * Renders a question's chapter, including the case where the classifier could
 * not place it, and flags low-confidence assignments instead of presenting a
 * guess as fact.
 */
export function ChapterBadge({ chapter }: { chapter?: ChapterRef }) {
  if (!chapter) {
    return (
      <Badge tone="warn">
        <AlertCircle size={12} /> Unclassified
      </Badge>
    );
  }
  const low = chapter.confidence !== undefined && chapter.confidence < LOW_CONFIDENCE;
  return (
    <Badge tone={low ? "warn" : "accent"}>
      Ch {chapter.number} · {chapter.title}
      {low && <span className="opacity-70">· needs review</span>}
    </Badge>
  );
}

export function Skeleton({ className }: { className?: string }) {
  return <div className={cn("shimmer rounded-md", className)} />;
}

export function EmptyState({
  title,
  description,
  action,
  icon,
}: {
  title: string;
  description?: string;
  action?: ReactNode;
  icon?: ReactNode;
}) {
  return (
    <div className="flex flex-col items-center justify-center rounded-[var(--radius-card)] border border-dashed px-6 py-16 text-center">
      <div className="mb-3 text-tertiary">{icon ?? <Inbox size={28} strokeWidth={1.5} />}</div>
      <h3 className="text-[15px] font-semibold">{title}</h3>
      {description && (
        <p className="mt-1.5 max-w-sm text-sm text-secondary leading-relaxed">{description}</p>
      )}
      {action && <div className="mt-5">{action}</div>}
    </div>
  );
}

export function ErrorState({ error, onRetry }: { error: unknown; onRetry?: () => void }) {
  const message = error instanceof Error ? error.message : "Something went wrong.";
  return (
    <div className="rounded-[var(--radius-card)] border border-red-500/25 bg-red-500/5 px-5 py-4">
      <div className="flex items-start gap-3">
        <AlertCircle size={18} className="mt-0.5 shrink-0 text-red-500" />
        <div>
          <p className="text-sm font-medium">Could not load this</p>
          <p className="mt-1 text-sm text-secondary">{message}</p>
          {onRetry && (
            <button
              onClick={onRetry}
              className="mt-3 rounded-lg border px-3 py-1.5 text-sm font-medium transition hover:surface-sunken"
            >
              Try again
            </button>
          )}
        </div>
      </div>
    </div>
  );
}

export function Button({
  children,
  variant = "primary",
  size = "md",
  className,
  ...props
}: {
  variant?: "primary" | "ghost" | "outline";
  size?: "sm" | "md";
} & React.ButtonHTMLAttributes<HTMLButtonElement>) {
  const variants = {
    primary: "accent-bg text-white hover:opacity-90 disabled:opacity-40",
    outline: "border hover:surface-sunken",
    ghost: "hover:surface-sunken",
  };
  const sizes = { sm: "px-2.5 py-1.5 text-[13px]", md: "px-3.5 py-2 text-sm" };
  return (
    <button
      className={cn(
        "inline-flex items-center justify-center gap-1.5 rounded-lg font-medium transition disabled:cursor-not-allowed",
        variants[variant],
        sizes[size],
        className,
      )}
      {...props}
    >
      {children}
    </button>
  );
}

export function PageHeader({
  title,
  subtitle,
  actions,
}: {
  title: ReactNode;
  subtitle?: ReactNode;
  actions?: ReactNode;
}) {
  return (
    <div className="mb-7 flex flex-wrap items-end justify-between gap-4">
      <div>
        <h1 className="text-[26px] font-semibold leading-tight tracking-[-0.02em]">{title}</h1>
        {subtitle && <p className="mt-1.5 text-sm text-secondary">{subtitle}</p>}
      </div>
      {actions && <div className="flex items-center gap-2">{actions}</div>}
    </div>
  );
}
