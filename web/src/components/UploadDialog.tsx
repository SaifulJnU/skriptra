import { useCallback, useEffect, useRef, useState } from "react";
import { useQueryClient } from "@tanstack/react-query";
import { AlertCircle, CheckCircle2, FileText, Loader2, UploadCloud, X } from "lucide-react";
import { api } from "@/lib/client";
import { ApiRequestError, type UploadMeta } from "@/lib/api";
import type { DocumentKind, IngestStatus } from "@/types/api";
import { Button } from "@/components/ui";

/**
 * Upload a past paper and watch it get indexed.
 *
 * Ingestion is asynchronous and takes tens of seconds, so this stays open and
 * reports each stage. A dialog that closed on 202 would leave the user staring
 * at unchanged counts with no idea whether anything happened.
 */

const KINDS: { value: DocumentKind; label: string; hint: string }[] = [
  { value: "exam", label: "Exam paper", hint: "Split into questions and classified by chapter" },
  { value: "solution", label: "Solutions", hint: "Worked answers for a paper" },
  { value: "notes", label: "Lecture notes", hint: "Indexed as searchable passages" },
  { value: "textbook", label: "Textbook", hint: "Indexed as searchable passages" },
  { value: "syllabus", label: "Syllabus", hint: "Used to build the chapter taxonomy" },
];

// Ordered, so progress reads as a sequence rather than a jumping label.
const STAGES: { key: IngestStatus; label: string }[] = [
  { key: "queued", label: "Queued" },
  { key: "parsing", label: "Extracting text" },
  { key: "segmenting", label: "Finding questions" },
  { key: "classifying", label: "Assigning chapters" },
  { key: "embedding", label: "Building the index" },
  { key: "indexed", label: "Done" },
];

export default function UploadDialog({
  courseId,
  onClose,
}: {
  courseId: string;
  onClose: () => void;
}) {
  const queryClient = useQueryClient();
  const [file, setFile] = useState<File | null>(null);
  const [kind, setKind] = useState<DocumentKind>("exam");
  const [year, setYear] = useState<string>("");
  const [term, setTerm] = useState<"" | "summer" | "winter">("");
  const [dragging, setDragging] = useState(false);

  const [sending, setSending] = useState(false);
  const [sendProgress, setSendProgress] = useState(0);
  const [documentId, setDocumentId] = useState<string | null>(null);
  const [status, setStatus] = useState<IngestStatus | null>(null);
  const [stageDetail, setStageDetail] = useState<string>();
  const [questionsFound, setQuestionsFound] = useState<number>();
  const [error, setError] = useState<string>();
  const [deduplicated, setDeduplicated] = useState(false);

  const inputRef = useRef<HTMLInputElement>(null);
  const pollRef = useRef<number | null>(null);

  // Escape closes, but not mid-upload: cancelling a request already in flight
  // would leave a half-written document behind.
  useEffect(() => {
    const onKey = (e: KeyboardEvent) => {
      if (e.key === "Escape" && !sending) onClose();
    };
    window.addEventListener("keydown", onKey);
    return () => window.removeEventListener("keydown", onKey);
  }, [sending, onClose]);

  useEffect(() => {
    return () => {
      if (pollRef.current) window.clearInterval(pollRef.current);
    };
  }, []);

  const finish = useCallback(() => {
    // Counts, chapter totals and the document list all change once a document
    // is indexed, so the whole course is refetched rather than guessing which
    // queries were affected.
    void queryClient.invalidateQueries();
  }, [queryClient]);

  const poll = useCallback(
    (id: string) => {
      pollRef.current = window.setInterval(async () => {
        try {
          const s = await api.documentStatus(id);
          setStatus(s.status);
          setStageDetail(s.stageDetail);
          setQuestionsFound(s.questionsExtracted);

          if (s.status === "indexed" || s.status === "failed") {
            if (pollRef.current) window.clearInterval(pollRef.current);
            pollRef.current = null;
            setSending(false);
            if (s.status === "failed") setError(s.error ?? "Indexing failed.");
            else finish();
          }
        } catch {
          // A dropped poll is not fatal; the next tick retries.
        }
      }, 1500);
    },
    [finish],
  );

  async function submit() {
    if (!file) return;
    setError(undefined);
    setSending(true);
    setSendProgress(0);

    const meta: UploadMeta = { kind };
    if (year) meta.year = Number(year);
    if (term) meta.term = term;

    try {
      const res = await api.uploadDocument(courseId, file, meta, setSendProgress);
      setDocumentId(res.id);

      if (res.deduplicated) {
        // Nothing was queued, so polling would spin forever on a terminal
        // status that never changes.
        setDeduplicated(true);
        setStatus("indexed");
        setSending(false);
        finish();
        return;
      }
      setStatus("queued");
      poll(res.id);
    } catch (err) {
      setSending(false);
      setError(
        err instanceof ApiRequestError ? err.message : "Something went wrong during upload.",
      );
    }
  }

  const stageIndex = status ? STAGES.findIndex((s) => s.key === status) : -1;
  const done = status === "indexed";

  return (
    <div
      className="fixed inset-0 z-50 flex items-center justify-center p-4"
      style={{ background: "color-mix(in oklch, black 45%, transparent)" }}
      onClick={() => !sending && onClose()}
    >
      <div
        role="dialog"
        aria-modal="true"
        aria-label="Upload a document"
        className="card w-full max-w-lg p-6"
        onClick={(e) => e.stopPropagation()}
      >
        <div className="mb-5 flex items-start justify-between gap-4">
          <div>
            <h2 className="text-[17px] font-semibold">Upload a document</h2>
            <p className="mt-1 text-sm text-secondary">
              PDFs with a text layer. Scans need OCR, which is not enabled yet.
            </p>
          </div>
          <button
            onClick={onClose}
            disabled={sending}
            aria-label="Close"
            className="rounded-lg p-1.5 text-tertiary transition hover:surface-sunken disabled:opacity-40"
          >
            <X size={17} />
          </button>
        </div>

        {!documentId && (
          <>
            <div
              onDragOver={(e) => {
                e.preventDefault();
                setDragging(true);
              }}
              onDragLeave={() => setDragging(false)}
              onDrop={(e) => {
                e.preventDefault();
                setDragging(false);
                const f = e.dataTransfer.files?.[0];
                if (f) setFile(f);
              }}
              onClick={() => inputRef.current?.click()}
              className={`flex cursor-pointer flex-col items-center justify-center rounded-[var(--radius-card)] border border-dashed px-6 py-9 text-center transition ${
                dragging ? "accent-soft-bg" : "hover:surface-sunken"
              }`}
            >
              <input
                ref={inputRef}
                type="file"
                accept="application/pdf,.pdf"
                className="hidden"
                onChange={(e) => setFile(e.target.files?.[0] ?? null)}
              />
              {file ? (
                <>
                  <FileText size={22} className="mb-2 accent-text" />
                  <p className="text-sm font-medium">{file.name}</p>
                  <p className="mt-1 text-xs text-tertiary">
                    {(file.size / 1024 / 1024).toFixed(1)} MB
                  </p>
                </>
              ) : (
                <>
                  <UploadCloud size={22} className="mb-2 text-tertiary" />
                  <p className="text-sm font-medium">Drop a PDF here, or click to choose</p>
                  <p className="mt-1 text-xs text-tertiary">Maximum 50 MB</p>
                </>
              )}
            </div>

            <div className="mt-5 space-y-4">
              <div>
                <label className="mb-1.5 block text-xs font-medium text-secondary">
                  What is it?
                </label>
                <select
                  value={kind}
                  onChange={(e) => setKind(e.target.value as DocumentKind)}
                  className="w-full rounded-lg border bg-transparent px-3 py-2 text-sm"
                >
                  {KINDS.map((k) => (
                    <option key={k.value} value={k.value}>
                      {k.label}
                    </option>
                  ))}
                </select>
                <p className="mt-1.5 text-xs text-tertiary">
                  {KINDS.find((k) => k.value === kind)?.hint}
                </p>
              </div>

              {(kind === "exam" || kind === "solution") && (
                <div className="grid grid-cols-2 gap-3">
                  <div>
                    <label className="mb-1.5 block text-xs font-medium text-secondary">Year</label>
                    <input
                      type="number"
                      value={year}
                      onChange={(e) => setYear(e.target.value)}
                      placeholder="2025"
                      min={1990}
                      max={2100}
                      className="w-full rounded-lg border bg-transparent px-3 py-2 text-sm"
                    />
                  </div>
                  <div>
                    <label className="mb-1.5 block text-xs font-medium text-secondary">Term</label>
                    <select
                      value={term}
                      onChange={(e) => setTerm(e.target.value as typeof term)}
                      className="w-full rounded-lg border bg-transparent px-3 py-2 text-sm"
                    >
                      <option value="">Not sure</option>
                      <option value="summer">Summer</option>
                      <option value="winter">Winter</option>
                    </select>
                  </div>
                </div>
              )}
              <p className="text-xs text-tertiary">
                Year and term let questions be browsed by sitting. You can leave them blank.
              </p>
            </div>
          </>
        )}

        {sending && sendProgress < 1 && !documentId && (
          <div className="mt-5">
            <div className="mb-1.5 flex justify-between text-xs text-secondary">
              <span>Uploading</span>
              <span className="tabular-nums">{Math.round(sendProgress * 100)}%</span>
            </div>
            <div className="h-1.5 overflow-hidden rounded-full surface-sunken">
              <div
                className="h-full rounded-full accent-bg transition-all"
                style={{ width: `${sendProgress * 100}%` }}
              />
            </div>
          </div>
        )}

        {documentId && !error && (
          <div className="mt-1 space-y-3">
            {deduplicated ? (
              <div className="flex items-start gap-3 rounded-lg surface-sunken px-4 py-3">
                <CheckCircle2 size={17} className="mt-0.5 shrink-0 text-emerald-500" />
                <div>
                  <p className="text-sm font-medium">Already in this course</p>
                  <p className="mt-0.5 text-xs text-secondary">
                    The same file was uploaded before, so nothing was indexed twice.
                  </p>
                </div>
              </div>
            ) : (
              STAGES.slice(0, -1).map((stage, i) => {
                const state = i < stageIndex ? "done" : i === stageIndex ? "active" : "todo";
                return (
                  <div key={stage.key} className="flex items-center gap-3">
                    <span className="flex h-5 w-5 shrink-0 items-center justify-center">
                      {state === "done" ? (
                        <CheckCircle2 size={16} className="text-emerald-500" />
                      ) : state === "active" ? (
                        <Loader2 size={15} className="animate-spin accent-text" />
                      ) : (
                        <span className="h-1.5 w-1.5 rounded-full surface-sunken" />
                      )}
                    </span>
                    <span
                      className={`text-sm ${state === "todo" ? "text-tertiary" : "text-secondary"}`}
                    >
                      {stage.label}
                    </span>
                    {state === "active" && stageDetail && (
                      <span className="text-xs text-tertiary">{stageDetail}</span>
                    )}
                  </div>
                );
              })
            )}

            {done && !deduplicated && (
              <div className="flex items-start gap-3 rounded-lg surface-sunken px-4 py-3">
                <CheckCircle2 size={17} className="mt-0.5 shrink-0 text-emerald-500" />
                <div>
                  <p className="text-sm font-medium">
                    Indexed {questionsFound ? `${questionsFound} questions` : "successfully"}
                  </p>
                  <p className="mt-0.5 text-xs text-secondary">
                    They are searchable now, and included in the analytics.
                  </p>
                </div>
              </div>
            )}
          </div>
        )}

        {error && (
          <div className="mt-5 flex items-start gap-3 rounded-lg border border-red-500/25 bg-red-500/5 px-4 py-3">
            <AlertCircle size={17} className="mt-0.5 shrink-0 text-red-500" />
            <div>
              <p className="text-sm font-medium">Upload failed</p>
              <p className="mt-0.5 text-xs text-secondary">{error}</p>
            </div>
          </div>
        )}

        <div className="mt-6 flex justify-end gap-2">
          {done || error ? (
            <Button onClick={onClose}>Close</Button>
          ) : (
            <>
              <Button variant="outline" onClick={onClose} disabled={sending}>
                Cancel
              </Button>
              <Button onClick={submit} disabled={!file || sending}>
                {sending ? (
                  <>
                    <Loader2 size={15} className="animate-spin" /> Working
                  </>
                ) : (
                  <>
                    <UploadCloud size={15} /> Upload
                  </>
                )}
              </Button>
            </>
          )}
        </div>
      </div>
    </div>
  );
}
