import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import tailwindcss from "@tailwindcss/vite";
import { fileURLToPath, URL } from "node:url";

export default defineConfig({
  plugins: [react(), tailwindcss()],
  resolve: {
    // Must mirror `paths` in tsconfig.json — TypeScript resolves the alias for
    // the editor, Vite resolves it for the bundle, and they are separate.
    alias: { "@": fileURLToPath(new URL("./src", import.meta.url)) },
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
