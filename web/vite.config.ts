import { defineConfig, Plugin } from "vite";
import vue from "@vitejs/plugin-vue";
import icons from "unplugin-icons/vite";
import IconsResolver from "unplugin-icons/resolver";
import components from "unplugin-vue-components/vite";
import fetch from "node-fetch";

const urlLoader = (): Plugin => {
  const cache: Record<string, string> = {};

  return {
    name: "vite-plugin-url-loader",
    enforce: "pre",
    resolveId(id) {
      if (id.startsWith("url:")) return `virtual:${id}`;
    },
    async load(id) {
      if (!id.startsWith("virtual:url:")) return;
      const url = id.replace("virtual:url:", "");
      cache[url] ??= await fetch(url).then((r) => r.text());
      return cache[url];
    },
  };
};

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => ({
  define: {
    __API_URL__:
      mode === "development" ? `"http://localhost:3000/graphql"` : `"/graphql"`,
  },
  plugins: [
    urlLoader(),
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
