import { useRef, useState } from "react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { FileUp, Plus, Trash2 } from "lucide-react";
import { api } from "@/lib/client";
import { ApiRequestError, type ProposedChapter } from "@/lib/api";
import { Button, Card } from "@/components/ui";

/**
 * Course setup: telling Skriptra what the course actually covers.
 *
 * This is the first thing a new course needs, and until it exists nothing else
 * works properly. Classification scores a question against chapter vocabulary,
 * so a course with no chapters leaves every question unclassified, which in turn
 * empties the chapter filter and the analytics.
 *
 * Two steps on purpose. The extraction is a proposal, not a decision: a contents
 * page is messy, rules find most chapters and miss some, and only the course
 * owner knows whether a line is a chapter or a heading inside one. The review
 * step is also where topics gain the exam vocabulary that a bare title does not
 * carry, and that is where most of the classification quality comes from.
 */
export default function CourseSetup({
  courseId,
  onDone,
  onClose,
}: {
  courseId: string;
  onDone?: (classified: number) => void;
  onClose?: () => void;
}) {
  const queryClient = useQueryClient();
  const fileRef = useRef<HTMLInputElement>(null);

  const [chapters, setChapters] = useState<ProposedChapter[] | null>(null);
  const [source, setSource] = useState<"rules" | "llm" | null>(null);
  const [error, setError] = useState<string | null>(null);

  const extract = useMutation({
    mutationFn: (file: File) => api.extractOutline(courseId, file),
    onSuccess: (proposal) => {
      setChapters(proposal.chapters);
      setSource(proposal.source);
      setError(null);
    },
    onError: (err) =>
      setError(
        err instanceof ApiRequestError
          ? err.message
          : "Could not read that file. Is the server running?",
      ),
  });

  const save = useMutation({
    mutationFn: () => api.saveChapters(courseId, chapters ?? []),
    onSuccess: (result) => {
      queryClient.invalidateQueries({ queryKey: ["chapters", courseId] });
      queryClient.invalidateQueries({ queryKey: ["course", courseId] });
      queryClient.invalidateQueries({ queryKey: ["questions"] });
      onDone?.(result.questionsClassified);
    },
    onError: (err) =>
      setError(err instanceof ApiRequestError ? err.message : "Could not save the chapters."),
  });

  function patch(index: number, next: Partial<ProposedChapter>) {
    setChapters((prev) =>
      (prev ?? []).map((ch, i) => (i === index ? { ...ch, ...next } : ch)),
    );
  }

  // ------------------------------------------------------------- step one --
  if (!chapters) {
    return (
      <Card className="p-6">
        <h2 className="text-base font-semibold">Add the course outline</h2>
        <p className="muted-text mt-1.5 text-sm leading-relaxed">
          Upload the syllabus, or the contents page of the course textbook.
          Skriptra reads the chapter list from it and uses that to sort every
          question you upload. Without it, questions stay unclassified and the
          chapter filter and analytics have nothing to work with.
        </p>

        <input
          ref={fileRef}
          type="file"
          accept=".pdf,.docx,.png,.jpg,.jpeg,.heic,.tiff"
          className="hidden"
          onChange={(e) => {
            const file = e.target.files?.[0];
            if (file) extract.mutate(file);
            // Reset, so choosing the same file twice after a failure still
            // fires a change event.
            e.target.value = "";
          }}
        />

        {error && (
          <p
            role="alert"
            className="mt-4 rounded-md border border-red-500/30 bg-red-500/10 px-3 py-2 text-sm text-red-600 dark:text-red-400"
          >
            {error}
          </p>
        )}

        <div className="mt-5 flex items-center gap-2">
          <Button onClick={() => fileRef.current?.click()} disabled={extract.isPending}>
            <FileUp size={15} />
            {extract.isPending ? "Reading..." : "Choose a file"}
          </Button>
          <Button
            variant="ghost"
            onClick={() => {
              // Nothing to extract from, so start with one empty row and type
              // it in. A student who has no syllabus file still knows their
              // chapters.
              setChapters([{ number: 1, title: "", topics: [] }]);
              setSource(null);
            }}
          >
            Enter them by hand
          </Button>
        </div>
      </Card>
    );
  }

  // ------------------------------------------------------------- step two --
  return (
    <Card className="p-6">
      <div className="flex items-start justify-between gap-4">
        <div>
          <h2 className="text-base font-semibold">Review the chapters</h2>
          <p className="muted-text mt-1.5 text-sm leading-relaxed">
            {source === "rules" && "Read from the contents page. "}
            {source === "llm" && "Read by the model, so check it carefully. "}
            Fix anything wrong and add the terms your exams actually use. Those
            terms are what questions are matched against, so a few good ones
            here are worth more than a tidy title.
          </p>
        </div>
        {onClose && (
          <Button variant="ghost" size="sm" onClick={onClose}>
            Cancel
          </Button>
        )}
      </div>

      <ul className="mt-5 space-y-3">
        {chapters.map((ch, i) => (
          <li key={i} className="flex items-start gap-2">
            <input
              type="number"
              min={1}
              max={99}
              value={ch.number}
              onChange={(e) => patch(i, { number: Number(e.target.value) })}
              aria-label="Chapter number"
              className="w-14 shrink-0 rounded-md border px-2 py-2 text-sm surface-sunken"
            />
            <div className="min-w-0 flex-1 space-y-1.5">
              <input
                value={ch.title}
                onChange={(e) => patch(i, { title: e.target.value })}
                placeholder="Chapter title"
                aria-label="Chapter title"
                className="w-full rounded-md border px-3 py-2 text-sm surface-sunken"
              />
              <input
                value={ch.topics.join(", ")}
                onChange={(e) =>
                  patch(i, {
                    topics: e.target.value
                      .split(",")
                      .map((t) => t.trim())
                      .filter(Boolean),
                  })
                }
                placeholder="Topics, comma separated: OLS, Gauss-Markov, BLUE"
                aria-label="Topics"
                className="w-full rounded-md border px-3 py-2 text-[13px] surface-sunken"
              />
            </div>
            <button
              onClick={() => setChapters(chapters.filter((_, j) => j !== i))}
              aria-label={`Remove chapter ${ch.number}`}
              className="mt-2 rounded p-1.5 text-tertiary hover:text-red-500"
            >
              <Trash2 size={14} />
            </button>
          </li>
        ))}
      </ul>

      <button
        onClick={() =>
          setChapters([
            ...chapters,
            { number: (chapters.at(-1)?.number ?? 0) + 1, title: "", topics: [] },
          ])
        }
        className="mt-3 flex items-center gap-1.5 text-sm font-medium accent-text hover:underline"
      >
        <Plus size={14} /> Add a chapter
      </button>

      {error && (
        <p
          role="alert"
          className="mt-4 rounded-md border border-red-500/30 bg-red-500/10 px-3 py-2 text-sm text-red-600 dark:text-red-400"
        >
          {error}
        </p>
      )}

      <div className="mt-6 flex justify-end gap-2">
        <Button variant="outline" onClick={() => setChapters(null)}>
          Start over
        </Button>
        <Button
          onClick={() => {
            setError(null);
            if (chapters.some((ch) => !ch.title.trim())) {
              setError("Every chapter needs a title.");
              return;
            }
            save.mutate();
          }}
          disabled={save.isPending || chapters.length === 0}
        >
          {save.isPending ? "Saving and sorting..." : `Save ${chapters.length} chapters`}
        </Button>
      </div>
    </Card>
  );
}
