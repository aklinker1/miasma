import {
  defaultTheme,
  defineUserConfig,
  HooksExposed,
  Page,
  Plugin,
} from "vuepress";
import { searchPlugin } from "@vuepress/plugin-search";
import { execSync } from "child_process";
import { renderSchema } from "graphql-markdown";
import { buildSchema, graphql, getIntrospectionQuery } from "graphql";
import fs from "fs/promises";
import path from "path";

type App = Parameters<HooksExposed["extendsPage"]>[1];

function generateCodePlugin(): Plugin {
  async function executeTemplate(
    page: Page,
    app: App,
    variables: Record<string, string>
  ) {
    for (const [key, mdValue] of Object.entries(variables)) {
      page.content = page.content.replace(`{{ ${key} }}`, mdValue);
      page.contentRendered = app.markdown.render(page.content);
    }
  }

  async function generateGraphqlSchema(): Promise<string> {
    console.log("Generating GraphQL schema docs...");

    // Merge schemas together
    const schemaDir = path.resolve(__dirname, "..", "api");
    const files = (await fs.readdir(schemaDir)).map((f) =>
      path.resolve(schemaDir, f)
    );
    const schemaStr = (
      await Promise.all(files.map((f) => fs.readFile(f, "utf-8")))
    ).join("\n");

    // Introspect
    const schema = buildSchema(schemaStr);
    const introspection = await graphql({
      schema,
      source: getIntrospectionQuery(),
    });

    // Render to markdown
    const content: string[] = [];
    const printer = (line: string) => {
      content.push(line);
    };
    renderSchema(introspection.data, {
      title: "Schema",
      headingLevel: 2,
      skipTableOfContents: true,
      printer,
    });

    // Additional processing
    return (
      content
        .join("\n")
        // Remove extra blank lines inside tables
        .replace(/<td>\n\n/gm, "<td>\n")
        .replace(/\n\n<\/td>/gm, "\n</td>")
    );
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
        const help = await generateCliHelp();
        await executeTemplate(page, app, { help });
      }
      if (page.path === "/reference/graphql.html") {
        const schema = await generateGraphqlSchema();
        await executeTemplate(page, app, { schema });
      }
    },
  };
}

export default defineUserConfig({
  lang: "en-US",
  title: "Miasma",
  description: "A Heroku-like, docker based PaaS with cluster and ARM support",
  define: {
    __CLI_VERSION__: require("../meta.json").cliVersion,
  },
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
  head: [["link", { rel: "icon", href: "/favicon.svg" }]],
  theme: defaultTheme({
    logo: "/nav-branding.svg",
    docsRepo: "https://github.com/aklinker1/miasma",
    docsDir: "docs",
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
            children: [
              "/reference/server-config.md",
              "/reference/cli.md",
              "/reference/graphql.md",
            ],
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
        "/guide/private-registries.md",
        {
          text: "Plugins",
          children: ["/plugins/index.md", "/plugins/traefik.md"],
        },
        "/faq.md",
        "/contributing.md",
      ],
      "/reference/": [
        "/reference/server-config.md",
        "/reference/cli.md",
        "/reference/graphql.md",
      ],
    },
  }),
});
