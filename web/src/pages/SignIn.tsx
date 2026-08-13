import { useState, type FormEvent } from "react";
import { Link, useLocation, useNavigate } from "react-router-dom";
import { Button, Card } from "@/components/ui";
import { AuthError, login, signup } from "@/lib/session";

/**
 * One component for both sign in and sign up.
 *
 * The two forms differ by a single field and a verb. Splitting them into two
 * files would duplicate the error handling, the redirect and the layout, and
 * the copies would drift.
 */
export default function SignIn({ mode }: { mode: "login" | "signup" }) {
  const navigate = useNavigate();
  const location = useLocation();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [displayName, setDisplayName] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [busy, setBusy] = useState(false);

  const isSignup = mode === "signup";

  async function onSubmit(e: FormEvent) {
    e.preventDefault();
    setError(null);
    setBusy(true);
    try {
      if (isSignup) {
        await signup(email, password, displayName || undefined);
      } else {
        await login(email, password);
      }
      // Back to whatever the guard interrupted, so a bookmarked deep link
      // survives signing in.
      const from = (location.state as { from?: string } | null)?.from ?? "/";
      navigate(from, { replace: true });
    } catch (err) {
      setError(
        err instanceof AuthError ? err.message : "Something went wrong. Try again.",
      );
    } finally {
      setBusy(false);
    }
  }

  return (
    <div className="mx-auto flex min-h-screen max-w-md flex-col justify-center px-6">
      <div className="mb-8">
        <h1 className="text-2xl font-semibold tracking-tight">Skriptra</h1>
        <p className="muted-text mt-1 text-sm">
          {isSignup
            ? "Create an account to upload your course material."
            : "Sign in to your courses."}
        </p>
      </div>

      <Card className="p-6">
        <form onSubmit={onSubmit} className="space-y-4">
          {isSignup && (
            <Field
              label="Name"
              value={displayName}
              onChange={setDisplayName}
              type="text"
              autoComplete="name"
              placeholder="Optional"
            />
          )}

          <Field
            label="Email"
            value={email}
            onChange={setEmail}
            type="email"
            autoComplete="email"
            required
          />

          <div>
            <Field
              label="Password"
              value={password}
              onChange={setPassword}
              type="password"
              // Tells a password manager to offer a new password on signup and
              // a saved one on login. Getting this wrong is why so many signup
              // forms fight the browser.
              autoComplete={isSignup ? "new-password" : "current-password"}
              required
              minLength={isSignup ? 10 : undefined}
            />
            {isSignup && (
              <p className="muted-text mt-1.5 text-xs">
                At least 10 characters. Length is what makes a password hard to
                guess, so there are no symbol rules.
              </p>
            )}
          </div>

          {error && (
            <p
              role="alert"
              className="rounded-md border border-red-500/30 bg-red-500/10 px-3 py-2 text-sm text-red-600 dark:text-red-400"
            >
              {error}
            </p>
          )}

          <Button type="submit" disabled={busy} className="w-full justify-center">
            {busy ? "Please wait..." : isSignup ? "Create account" : "Sign in"}
          </Button>
        </form>
      </Card>

      <p className="muted-text mt-6 text-center text-sm">
        {isSignup ? "Already have an account? " : "No account yet? "}
        <Link
          to={isSignup ? "/signin" : "/signup"}
          className="accent-text font-medium hover:underline"
        >
          {isSignup ? "Sign in" : "Create one"}
        </Link>
      </p>
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
