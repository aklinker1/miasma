import { defaultTheme, defineUserConfig } from "vuepress";
import fs from "fs/promises";
import { searchPlugin } from "@vuepress/plugin-search";

async function generateGraphqlDocs() {
  await fs.writeFile("docs/generated/graphql-schema.md", "");
}

async function generateCliDocs() {
  await fs.writeFile("docs/generated/cli-help.md", "");
}

export default defineUserConfig({
  lang: "en-US",
  title: "Miasma",
  description: "A Heroku-like, docker based PaaS with cluster and ARM support",
  plugins: [
    searchPlugin({
      locales: {
        "/": {
          placeholder: "Search ( / )",
        },
      },
    }),
  ],
  async onInitialized() {
    await generateGraphqlDocs();
    await generateCliDocs();
  },
  theme: defaultTheme({
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
        {
          text: "Plugins",
          children: ["/plugins/index.md", "/plugins/traefik.md"],
        },
        "/examples.md",
        "/faq.md",
        "/contributing.md",
      ],
    },
  }),
});
