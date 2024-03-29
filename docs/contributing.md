---
title: Contributing
---

# Contributing

<a href="https://github.com/aklinker1/miasma/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=aklinker1/miasma" />
</a>

Welcome to the contributor's guide! Here you'll find everything you need to know about contributing to Miasma.

[[toc]]

## Required Tools

- [Node 18 LTS](https://nodejs.org/)
- [PNPM](https://pnpm.io)

## Tech Stack

Miasma uses [Nuxt](https://nuxt.com) to create a server-side-rendered web application (SSR).

The CSS framework is [TailwindCSS](https://tailwindcss.com/) + [DaisyUI](https://daisyui.com/).

The UI uses an API proxy (`/api/docker`) for communicating with the docker engine.

```bash
# Create a swarm on your development computer
docker swarm init

# Start Miasma in dev mode
pnpm dev

# Build the Nuxt application
pnpm build

# Build the docker image that someone would install from docker hub
pnpm docker:build

# Run the built docker image
pnpm docker:run
```

### Docker Socket Customization

When running `pnpm dev` or `pnpm docker:run`, you might run into 500 status code errors when attempting to connect to docker. By default, it tries to connect to `/var/run/docker.sock`. But the docker socket isn't always hosted there, like on Windows. To fix this, add a `.env` file to the project root, and fill out the `MIASMA_DOCKER_SOCKET` variable.

For example, if you use Colima on a Macbook, it hosts the docker socket at `$HOME/.colima/default/docker.sock`. So the `.env` would look like this:

```txt
MIASMA_DOCKER_SOCKET=/Users/<username>/.colima/default/docker.sock
```

## Documentation Website

The docs are located under `docs/` directory, and can be ran locally with:

```bash
pnpm docs:dev
```

The website is made using [VitePress](https://v2.vuepress.vuejs.org/), and published to [GitHub pages](https://pages.github.com/).

## Publish a New Release

> Maintainers only

The project uses conventional commits, so when PRs are merged, make sure to update their titles or commit titles to a conventional commit.

Then, to preform a release, simply dispatch the ["Publish Docker Image"](https://github.com/aklinker1/miasma/actions/workflows/publish-docker.yml) GitHub Action.

It will bump the version, build the Docker Image, push it to Docker Hub, and create a release with release notes on GitHub.
