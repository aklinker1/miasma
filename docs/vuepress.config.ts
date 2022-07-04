import { defaultTheme, defineUserConfig, Page, Plugin } from "vuepress";
import { searchPlugin } from "@vuepress/plugin-search";
import { execSync } from "child_process";

function generateCodePlugin(): Plugin {
  async function executeTemplate(
    page: Page,
    variables: Record<string, string>
  ) {
    for (const [key, value] of Object.entries(variables)) {
      const templateKey = `{{ ${key} }}`;
      while (page.contentRendered.includes(templateKey))
        page.contentRendered = page.contentRendered.replace(templateKey, value);
    }
  }

  async function generateGraphqlSchema(): Promise<string> {
    console.log("Generating GraphQL schema docs...");
    return ":::danger\nTODO\n:::";
  }

  async function generateCliHelp(): Promise<string> {
    console.log("Generating CLI help docs...");
    return execSync("go run cmd/print-cli-docs/main.go", {
      cwd: "..",
      encoding: "utf-8",
    });
  }

  return {
    name: "GeneratedCode",

    async extendsPage(page, app) {
      if (page.path === "/reference/cli.html") {
        const help = app.markdown.render(await generateCliHelp());
        await executeTemplate(page, { help });
      }
      if (page.path === "/reference/graphql.html") {
        const schema = app.markdown.render(await generateGraphqlSchema());
        await executeTemplate(page, { schema });
      }
    },
  };
}

export default defineUserConfig({
  lang: "en-US",
  title: "Miasma",
  description: "A Heroku-like, docker based PaaS with cluster and ARM support",
  // @ts-expect-error: Untyped env var doesn't match /string/
  base: process.env.VUE_PRESS_BASE,
  plugins: [
    generateCodePlugin(),
    searchPlugin({
      locales: {
        "/": {
          placeholder: "Search ( / )",
        },
      },
    }),
  ],
  theme: defaultTheme({
    docsRepo: "https://github.com/aklinker1/miasma",
    docsDir: "docs/docs",
    editLink: true,
    repo: "https://github.com/aklinker1/miasma",
    navbar: [
      {
        text: "Guide",
        link: "/",
      },
      {
        text: "Reference",
        children: [
          {
            text: "Miasma",
            children: ["/reference/cli.md", "/reference/graphql.md"],
          },
        ],
      },
    ],
    sidebar: {
      "/": [
        "/index.md",
        "/guide/installation.md",
        "/guide/first-app.md",
        "/guide/port-management.md",
        "/guide/app-communication.md",
        "/guide/examples.md",
        "/authorization.md",
        {
          text: "Plugins",
          children: ["/plugins/index.md", "/plugins/traefik.md"],
        },
        "/faq.md",
        "/contributing.md",
      ],
      "/reference/": ["/reference/cli.md", "/reference/graphql.md"],
    },
  }),
});
