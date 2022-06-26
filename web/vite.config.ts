import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import icons from "unplugin-icons/vite";
import IconsResolver from "unplugin-icons/resolver";
import components from "unplugin-vue-components/vite";

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => ({
  define: {
    __API_URL__:
      mode === "development" ? `"http://localhost:3000/graphql"` : `"/graphql"`,
  },
  plugins: [
    vue(),
    icons(),
    components({
      resolvers: [IconsResolver()],
      dts: "src/@types/components.d.ts",
    }),
  ],
  server: {
    port: 8080,
  },
}));
