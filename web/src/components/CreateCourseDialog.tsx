import { useState, type FormEvent } from "react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";
import { api } from "@/lib/client";
import { ApiRequestError } from "@/lib/api";
import { Button } from "@/components/ui";

/**
 * Creating a course is the first thing a new account has to do.
 *
 * Membership is what grants access to anything, so an account that belongs to
 * no course sees an empty product with no way forward. This dialog is that way
 * forward, and the caller becomes the course owner.
 */
export default function CreateCourseDialog({ onClose }: { onClose: () => void }) {
  const navigate = useNavigate();
  const queryClient = useQueryClient();

  const [name, setName] = useState("");
  const [code, setCode] = useState("");
  const [institution, setInstitution] = useState("");
  const [language, setLanguage] = useState<"en" | "de">("en");
  const [error, setError] = useState<string | null>(null);

  const create = useMutation({
    mutationFn: () =>
      api.createCourse({
        name: name.trim(),
        code: code.trim() || undefined,
        institution: institution.trim() || undefined,
        language,
      }),
    onSuccess: (course) => {
      queryClient.invalidateQueries({ queryKey: ["courses"] });
      // Straight into the new course, because the next thing to do is upload a
      // paper, and that lives there.
      navigate(`/courses/${course.id}`);
      onClose();
    },
    onError: (err) =>
      setError(
        err instanceof ApiRequestError
          ? err.message
          : "Could not create the course. Is the server running?",
      ),
  });

  function onSubmit(e: FormEvent) {
    e.preventDefault();
    setError(null);
    if (!name.trim()) {
      setError("A course needs a name.");
      return;
    }
    create.mutate();
  }

  return (
    <div
      className="fixed inset-0 z-50 flex items-center justify-center bg-black/40 p-4"
      onClick={onClose}
    >
      <div
        className="w-full max-w-md rounded-[var(--radius-card)] border surface p-6 shadow-xl"
        onClick={(e) => e.stopPropagation()}
        role="dialog"
        aria-modal="true"
        aria-labelledby="create-course-title"
      >
        <h2 id="create-course-title" className="text-lg font-semibold">
          New course
        </h2>
        <p className="muted-text mt-1 text-sm">
          You will be its owner. Upload past papers once they are in.
        </p>

        <form onSubmit={onSubmit} className="mt-5 space-y-4">
          <Field
            label="Name"
            value={name}
            onChange={setName}
            placeholder="Linear Models"
            autoFocus
            required
          />
          <div className="grid grid-cols-2 gap-3">
            <Field label="Code" value={code} onChange={setCode} placeholder="STAT-412" />
            <label className="block">
              <span className="mb-1.5 block text-sm font-medium">Language</span>
              <select
                value={language}
                onChange={(e) => setLanguage(e.target.value as "en" | "de")}
                className="w-full rounded-md border px-3 py-2 text-sm surface-sunken"
                // Explicit colours: the inherited ones made options unreadable
                // against the native dropdown background in dark mode.
                style={{ color: "var(--text)" }}
              >
                <option value="en">English</option>
                <option value="de">German</option>
              </select>
            </label>
          </div>
          <Field
            label="Institution"
            value={institution}
            onChange={setInstitution}
            placeholder="TU Dortmund"
          />

          {error && (
            <p
              role="alert"
              className="rounded-md border border-red-500/30 bg-red-500/10 px-3 py-2 text-sm text-red-600 dark:text-red-400"
            >
              {error}
            </p>
          )}

          <div className="flex justify-end gap-2 pt-1">
            <Button type="button" variant="outline" onClick={onClose}>
              Cancel
            </Button>
            <Button type="submit" disabled={create.isPending}>
              {create.isPending ? "Creating..." : "Create course"}
            </Button>
          </div>
        </form>
      </div>
    </div>
  );
}

function Field({
  label,
  value,
  onChange,
  ...props
}: {
  label: string;
  value: string;
  onChange: (v: string) => void;
} & Omit<React.InputHTMLAttributes<HTMLInputElement>, "onChange" | "value">) {
  return (
    <label className="block">
      <span className="mb-1.5 block text-sm font-medium">{label}</span>
      <input
        {...props}
        value={value}
        onChange={(e) => onChange(e.target.value)}
        className="w-full rounded-md border px-3 py-2 text-sm outline-none surface-sunken focus:ring-2 focus:ring-[var(--accent)]"
      />
    </label>
  );
}
