import { Fragment, type ReactNode } from "react";

/**
 * A small Markdown renderer for model output.
 *
 * Deliberately hand-written rather than pulling in react-markdown. The input is
 * not arbitrary Markdown, it is the narrow subset a model produces for a worked
 * solution: headings, bold, lists, and equations on their own lines. Handling
 * that subset is around a hundred lines and avoids shipping a parser plus a
 * sanitiser for syntax that never arrives.
 *
 * Nothing here uses dangerouslySetInnerHTML. Every node is a real React
 * element, so model output cannot inject markup no matter what it emits.
 */

/** Lines that are mostly symbols are displayed as equations rather than prose. */
function looksLikeEquation(line: string): boolean {
  const t = line.trim();
  if (t.length < 3 || t.length > 120) return false;
  if (/^[-*\d]/.test(t) && !/[=<>]/.test(t)) return false;

  const symbols = (t.match(/[=+\-*/^()[\]{}|∑∫√βεσμθλΣΩ<>]/g) ?? []).length;
  const letters = (t.match(/[a-zA-Z]/g) ?? []).length;
  // An equation is symbol-dense and short on connective words.
  return symbols >= 3 && symbols >= letters / 4 && t.split(/\s+/).length <= 18;
}

/** Renders bold, italic and inline code inside a single line. */
function inline(text: string, keyPrefix: string): ReactNode[] {
  const out: ReactNode[] = [];
  // One pass over the three markers, longest first so ** wins over *.
  const pattern = /(\*\*[^*]+\*\*|\*[^*]+\*|`[^`]+`)/g;
  let last = 0;
  let match: RegExpExecArray | null;
  let i = 0;

  while ((match = pattern.exec(text)) !== null) {
    if (match.index > last) out.push(text.slice(last, match.index));
    const token = match[0];

    if (token.startsWith("**")) {
      out.push(
        <strong key={`${keyPrefix}-b${i}`} className="font-semibold" style={{ color: "var(--text-primary)" }}>
          {token.slice(2, -2)}
        </strong>,
      );
    } else if (token.startsWith("`")) {
      out.push(
        <code
          key={`${keyPrefix}-c${i}`}
          className="rounded px-1 py-0.5 font-mono text-[0.9em]"
          style={{ background: "var(--surface-sunken)" }}
        >
          {token.slice(1, -1)}
        </code>,
      );
    } else {
      out.push(
        <em key={`${keyPrefix}-i${i}`}>{token.slice(1, -1)}</em>,
      );
    }

    last = match.index + token.length;
    i++;
  }
  if (last < text.length) out.push(text.slice(last));
  return out;
}

export function Markdown({ children, className }: { children: string; className?: string }) {
  const lines = children.split("\n");
  const blocks: ReactNode[] = [];

  let paragraph: string[] = [];
  let list: { ordered: boolean; items: string[] } | null = null;

  const flushParagraph = (key: string) => {
    if (!paragraph.length) return;
    blocks.push(
      <p key={key} className="mb-3 last:mb-0">
        {inline(paragraph.join(" "), key)}
      </p>,
    );
    paragraph = [];
  };

  const flushList = (key: string) => {
    if (!list) return;
    const Tag = list.ordered ? "ol" : "ul";
    blocks.push(
      <Tag
        key={key}
        className={`mb-3 space-y-1.5 pl-5 last:mb-0 ${list.ordered ? "list-decimal" : "list-disc"}`}
      >
        {list.items.map((item, i) => (
          <li key={i} className="pl-1">
            {inline(item, `${key}-${i}`)}
          </li>
        ))}
      </Tag>,
    );
    list = null;
  };

  lines.forEach((raw, index) => {
    const line = raw.trimEnd();
    const key = `b${index}`;

    if (!line.trim()) {
      flushParagraph(key);
      flushList(key);
      return;
    }

    const heading = line.match(/^(#{1,4})\s+(.*)$/);
    if (heading) {
      flushParagraph(key);
      flushList(key);
      const level = heading[1].length;
      blocks.push(
        <p
          key={key}
          className={`mb-2 mt-4 font-semibold first:mt-0 ${level <= 2 ? "text-[15px]" : "text-[14px]"}`}
          style={{ color: "var(--text-primary)" }}
        >
          {inline(heading[2], key)}
        </p>,
      );
      return;
    }

    const bullet = line.match(/^\s*[-*+]\s+(.*)$/);
    const numbered = line.match(/^\s*(\d+)[.)]\s+(.*)$/);

    if (bullet || numbered) {
      flushParagraph(key);
      const ordered = !!numbered;
      const text = bullet ? bullet[1] : numbered![2];
      if (!list || list.ordered !== ordered) {
        flushList(key);
        list = { ordered, items: [] };
      }
      list.items.push(text);
      return;
    }

    flushList(key);

    // An equation on its own line gets its own block, centred and monospaced,
    // so a derivation is scannable instead of buried in a paragraph.
    if (!paragraph.length && looksLikeEquation(line)) {
      blocks.push(
        <div
          key={key}
          className="mb-3 overflow-x-auto rounded-lg px-4 py-3 text-center font-mono text-[13.5px] last:mb-0"
          style={{ background: "var(--surface-sunken)", color: "var(--text-primary)" }}
        >
          {line.trim()}
        </div>,
      );
      return;
    }

    paragraph.push(line.trim());
  });

  flushParagraph("tail");
  flushList("tail-list");

  return <div className={className}>{blocks.map((b, i) => <Fragment key={i}>{b}</Fragment>)}</div>;
}
