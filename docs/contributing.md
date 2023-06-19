---
title: Contributing
---

# Contributing

Welcome to the contributor's guide! Here you'll find everything you need to know about contributing to Miasma.

<a href="https://github.com/aklinker1/miasma/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=aklinker1/miasma" />
</a>

[[toc]]

## Required Tools

- [Node 18 LTS](https://nodejs.org/)
- [PNPM](https://pnpm.io)

## Tech Stack

Miasma uses [Nuxt](https://nuxt.com) to create a Server-side-rendered application (SSR).

The CSS framework is [TailwindCSS](https://tailwindcss.com/) + [DaisyUI](https://daisyui.com/).

The UI uses a API proxy (`/api/docker`) for communicating with the docker engine.

```bash
# Start Miasma in dev mode
pnpm dev

# Build the Nuxt application
pnpm build

# Build the docker image that someone would install from docker hub
pnpm docker:build

# Run the built docker image
pnpm docker:run
```

## Docs

```bash
pnpm docs:dev
```

The website is made using [VitePress](https://v2.vuepress.vuejs.org/), and published to [GitHub pages](https://pages.github.com/).
