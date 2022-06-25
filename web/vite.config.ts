import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => ({
  define: {
    __API_URL__:
      mode === "development" ? `"http://localhost:3000/graphql"` : `"/graphql"`,
  },
  plugins: [vue()],
  server: {
    port: 8080,
  },
}));
