import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function termLabel(term?: string) {
  return term === "summer" ? "Summer" : term === "winter" ? "Winter" : "";
}

export function examLabel(year?: number, term?: string) {
  return year ? `${year} ${termLabel(term)}`.trim() : "";
}

/** Confidence below 0.7 is surfaced for review rather than shown as fact. */
export const LOW_CONFIDENCE = 0.7;

export function formatPercent(v: number, digits = 0) {
  return `${(v * 100).toFixed(digits)}%`;
}
