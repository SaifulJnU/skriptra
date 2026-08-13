import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import tailwindcss from "@tailwindcss/vite";
import path from "node:path";

export default defineConfig({
  plugins: [react(), tailwindcss()],
  resolve: {
    // Must mirror `paths` in tsconfig.json. TypeScript resolves the alias for
    // the editor, Vite resolves it for the bundle, and they are separate.
    //
    // Deliberately process.cwd() rather than import.meta.url. When this file is
    // edited while the dev server is running, Vite reloads the config by
    // bundling it into node_modules/.vite-temp, and import.meta.url then points
    // at that temp file. The alias would resolve to a directory that does not
    // exist and every "@/" import would fail until the server was restarted by
    // hand. Vite always runs from the project root, so cwd is stable across a
    // reload.
    alias: { "@": path.resolve(process.cwd(), "src") },
  },
  server: {
    port: 5173,
    // Windows bind mounts do not deliver inotify events to Linux containers,
    // so the dev server never sees a file change without polling.
    watch: { usePolling: true },
    proxy: {
      "/api": {
        target: process.env.VITE_API_PROXY ?? "http://localhost:8080",
        changeOrigin: true,
      },
    },
  },
});
