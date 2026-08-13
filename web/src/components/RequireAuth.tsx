import { useEffect, useState } from "react";
import { Navigate, Outlet, useLocation } from "react-router-dom";
import { usingMocks } from "@/lib/client";
import { getUser, refreshOnce, setSession } from "@/lib/session";

/**
 * Route guard.
 *
 * The access token is held in memory, so a page reload starts with no session
 * even when the user is signed in. The HttpOnly refresh cookie is what
 * survives, so the first thing the app does is try to exchange it. Until that
 * answers, the app must not render either the signed-in view or the sign-in
 * page: showing the login form for a moment to a user who is signed in is the
 * flicker that makes an app feel broken.
 *
 * This is a convenience, not a security boundary. It decides what to render,
 * never what is allowed. Every endpoint is enforced server side, and the
 * client can only ever get this wrong in its own favour by showing a page that
 * then fails to load data.
 */
export default function RequireAuth() {
  const location = useLocation();
  const [state, setState] = useState<"checking" | "in" | "out">(
    getUser() ? "in" : "checking",
  );

  useEffect(() => {
    if (state !== "checking") return;

    // Mock mode has no server to refresh against. The adapter exists so the
    // interface can be built and reviewed without a backend, and a login wall
    // in front of it would defeat that.
    if (usingMocks) {
      setSession({
        accessToken: "mock",
        tokenType: "Bearer",
        expiresIn: 3600,
        user: { id: "11111111-1111-1111-1111-111111111111", displayName: "Saiful" },
      });
      setState("in");
      return;
    }

    let cancelled = false;

    refreshOnce().then((session) => {
      if (!cancelled) setState(session ? "in" : "out");
    });

    return () => {
      cancelled = true;
    };
  }, [state]);

  if (state === "checking") {
    return (
      <div className="flex min-h-screen items-center justify-center">
        <span className="muted-text text-sm">Loading...</span>
      </div>
    );
  }

  if (state === "out") {
    // The attempted path is carried through so signing in returns the user
    // where they were going rather than dumping them on the dashboard.
    return <Navigate to="/signin" replace state={{ from: location.pathname + location.search }} />;
  }

  return <Outlet />;
}
