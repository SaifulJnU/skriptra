/**
 * Access token storage and refresh.
 *
 * The access token lives in a module variable, not in localStorage. Anything in
 * localStorage is readable by any script the page ever loads, so one
 * compromised dependency reads the token. Keeping it in memory means a page
 * reload loses it, which is exactly what the HttpOnly refresh cookie is for:
 * the app asks /auth/refresh on boot and gets a new one. The refresh token
 * itself is never visible to this code at all.
 */
import type { User } from "@/types/api";

const BASE = import.meta.env.VITE_API_BASE_URL ?? "/api/v1";

let accessToken: string | null = null;
let currentUser: User | null = null;
let listeners: Array<(u: User | null) => void> = [];

export interface Session {
  accessToken: string;
  tokenType: string;
  expiresIn: number;
  user: User;
}

export function getAccessToken(): string | null {
  return accessToken;
}

export function getUser(): User | null {
  return currentUser;
}

export function setSession(s: Session | null): void {
  accessToken = s?.accessToken ?? null;
  currentUser = s?.user ?? null;
  for (const fn of listeners) fn(currentUser);
}

export function onSessionChange(fn: (u: User | null) => void): () => void {
  listeners.push(fn);
  return () => {
    listeners = listeners.filter((l) => l !== fn);
  };
}

/**
 * Exchanges the refresh cookie for a new access token.
 *
 * `credentials: "include"` is what sends the cookie at all. Without it the
 * browser omits it cross-origin, which is exactly the case in development
 * where the Vite dev server and the API are different origins.
 */
export async function refreshSession(): Promise<Session | null> {
  try {
    const res = await fetch(`${BASE}/auth/refresh`, {
      method: "POST",
      credentials: "include",
    });
    if (!res.ok) {
      setSession(null);
      return null;
    }
    const session = (await res.json()) as Session;
    setSession(session);
    return session;
  } catch {
    // Offline or the API is down. Neither is a signed-out state, but there is
    // no token to work with either, so the caller sends the user to sign in.
    setSession(null);
    return null;
  }
}

/**
 * Serialises concurrent refreshes.
 *
 * On a cold load several queries fire at once and all of them get a 401. Each
 * retrying independently would rotate the refresh token several times over, and
 * rotation invalidates the previous one, so all but the first would fail and
 * log the user out. One in-flight refresh, shared by every waiter.
 */
let inFlight: Promise<Session | null> | null = null;

export function refreshOnce(): Promise<Session | null> {
  if (!inFlight) {
    inFlight = refreshSession().finally(() => {
      inFlight = null;
    });
  }
  return inFlight;
}

export async function login(email: string, password: string): Promise<Session> {
  return authRequest("/auth/login", { email, password });
}

export async function signup(
  email: string,
  password: string,
  displayName?: string,
): Promise<Session> {
  return authRequest("/auth/signup", { email, password, displayName });
}

export async function logout(): Promise<void> {
  try {
    await fetch(`${BASE}/auth/logout`, { method: "POST", credentials: "include" });
  } finally {
    // Cleared locally whatever the server said. A user who clicked sign out
    // must end up signed out even if the request failed.
    setSession(null);
  }
}

async function authRequest(path: string, body: unknown): Promise<Session> {
  let res: Response;
  try {
    res = await fetch(`${BASE}${path}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
      body: JSON.stringify(body),
    });
  } catch {
    // fetch rejects rather than resolving when the request never completed:
    // the API is down, or the browser refused the response. A CORS refusal
    // arrives here with no status and no body, so reporting it as a failed
    // sign-in would blame the password for a server problem.
    throw new AuthError(
      "network",
      "Could not reach the server. Is the API running?",
    );
  }

  const payload = await res.json().catch(() => ({}));
  if (!res.ok) {
    const err = payload as { error?: { code?: string; message?: string } };
    throw new AuthError(
      err.error?.code ?? "auth_failed",
      err.error?.message ?? "Could not sign you in.",
    );
  }

  const session = payload as Session;
  setSession(session);
  return session;
}

export class AuthError extends Error {
  constructor(
    public code: string,
    message: string,
  ) {
    super(message);
    this.name = "AuthError";
  }
}
