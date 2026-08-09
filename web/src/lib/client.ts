import type { LernovaApi } from "./api";
import { httpApi } from "./api";
import { mockApi } from "@/mocks/mockApi";

/**
 * The one place that decides which implementation the app talks to.
 *
 * Set `VITE_USE_MOCKS=false` once the Go server is running. Nothing else in the
 * application changes — every component depends on `LernovaApi`, never on
 * fetch, and never on the mock.
 */
const useMocks = (import.meta.env.VITE_USE_MOCKS ?? "true") !== "false";

export const api: LernovaApi = useMocks ? mockApi : httpApi;
export const usingMocks = useMocks;
