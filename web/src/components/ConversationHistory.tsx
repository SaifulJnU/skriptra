import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { MessageSquare, Plus, Trash2 } from "lucide-react";
import { api } from "@/lib/client";
import { cn } from "@/lib/utils";

/**
 * Past threads for a course.
 *
 * The `conversations` and `messages` tables were in the first migration and
 * nothing wrote to them: every answer minted a conversation id and discarded
 * it. This is the other half of making that id mean something.
 */
export default function ConversationHistory({
  courseId,
  activeId,
  onSelect,
  onNew,
}: {
  courseId: string;
  activeId: string | null;
  onSelect: (id: string) => void;
  onNew: () => void;
}) {
  const queryClient = useQueryClient();

  const { data, isLoading } = useQuery({
    queryKey: ["conversations", courseId],
    queryFn: () => api.listConversations(courseId),
    enabled: Boolean(courseId),
  });

  const remove = useMutation({
    mutationFn: (id: string) => api.deleteConversation(id),
    onSuccess: (_, id) => {
      queryClient.invalidateQueries({ queryKey: ["conversations", courseId] });
      // Deleting the thread you are reading has to clear the transcript too,
      // or the page shows messages that no longer exist anywhere.
      if (id === activeId) onNew();
    },
  });

  const conversations = data?.data ?? [];

  return (
    <div className="space-y-3">
      <div className="flex items-center justify-between">
        <h2 className="text-xs font-semibold uppercase tracking-wide text-tertiary">History</h2>
        <button
          onClick={onNew}
          className="flex items-center gap-1 text-xs font-medium accent-text hover:underline"
        >
          <Plus size={12} />
          New
        </button>
      </div>

      {isLoading && <p className="muted-text text-xs">Loading...</p>}

      {!isLoading && conversations.length === 0 && (
        <p className="muted-text text-xs leading-relaxed">
          Nothing yet. Questions you ask here are saved so you can come back to
          them.
        </p>
      )}

      <ul className="space-y-0.5">
        {conversations.map((c) => (
          <li key={c.id} className="group relative">
            <button
              onClick={() => onSelect(c.id)}
              className={cn(
                "flex w-full items-start gap-2 rounded-md px-2 py-2 pr-7 text-left text-[13px] leading-snug hover:surface-sunken",
                c.id === activeId && "surface-sunken font-medium",
              )}
            >
              <MessageSquare size={13} className="mt-0.5 shrink-0 text-tertiary" />
              <span className="line-clamp-2">{c.title || "Untitled"}</span>
            </button>

            <button
              onClick={() => remove.mutate(c.id)}
              aria-label={`Delete conversation: ${c.title}`}
              className="absolute right-1 top-1.5 rounded p-1 text-tertiary opacity-0 transition-opacity hover:text-red-500 focus:opacity-100 group-hover:opacity-100"
            >
              <Trash2 size={12} />
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
}
