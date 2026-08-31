import react from "@vitejs/plugin-react";
import { defineConfig } from "vitest/config";

// One config at the workspace root, because the workspace is one component and
// one install. The `web` component's runner invokes vitest here, not per package.
export default defineConfig({
  plugins: [react()],
  server: {
    // packages/app reads docs/openapi.json, which sits above this directory.
    fs: { allow: [".."] },
  },
  test: {
    environment: "jsdom",
    include: ["packages/*/src/**/*.test.{ts,tsx}"],
  },
});
