---
title: Contributing
---

# Contributing

Welcome to the contributor's guide! Here you'll find everything you need to know about contributing to Miasma.

[[toc]]

## Required Tools

- [GNU `make`](https://www.gnu.org/software/make/) to run scripts
- [Go 1.18](https://go.dev/doc/install) to build the CLI
- [Node 16 LTS](https://nodejs.org/en/) to run the UI in dev mode
- [Docker](https://docs.docker.com/get-docker/) to build the server and production docker image

## Project Structure

<script setup>
  import { h } from 'vue';
  const FileIcon = (props) => h(
    "span",
    {
      title: props.description ?? "No description",
      style: {
        display: "inline-flex",
        'align-items': "center",
        height: '19px',
        color: 'var()',
        cursor: 'pointer',
      },
    },
    [
      props.icon == "folder"
        ? h("svg", {
            width: "22",
            height: "22",
            viewBox: "0 0 24 24",
            fill: "none",
            xmlns: "http://www.w3.org/2000/svg"
          }, [
            h("path", {
              d: "M10 4H4C2.89 4 2 4.89 2 6V18C2 19.097 2.903 20 4 20H20C21.097 20 22 19.097 22 18V8C22 7.46957 21.7893 6.96086 21.4142 6.58579C21.0391 6.21071 20.5304 6 20 6H12L10 4Z",
              fill: "currentColor"
            })
          ])
        : h(
          "img",
          {
            src: `https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/main/icons/${props.icon ?? props.name.split(".").pop()?.toLowerCase()}.svg`,
            width: '22',
          },
      ),
      " " + props.name,
    ]
  )
</script>

<pre  style="background-color: var(--c-tip-bg);">
<FileIcon name="miasma/" icon="folder" />
  <FileIcon name="api" icon="folder-api" />
    <FileIcon
      name="*.graphqls"
      description="GraphQL schema definitions for the API"
      icon="graphql"
    />
  <FileIcon
    name="cmd/"
    description="Go executable entrypoints"
    icon="folder"
  />
    <FileIcon
      name="*/"
      description="Binary name"
      icon="folder"
    />
      <FileIcon
        name="main.go"
        description="Main packages for go executables"
      />
  <FileIcon
    name="docs/"
    description="VuePress documentation website"
    icon="folder-docs"
  />
  <FileIcon
    name="internal/"
    description="Go package for Miasma types, like GraphQL models"
    icon="folder"
  />
    <FileIcon
      name="server/"
      description="Go package for server only code"
      icon="folder"
    />
    <FileIcon
      name="cli/"
      description="Go package for CLI only code"
      icon="folder"
    />
    <FileIcon
      name="utils/"
      description="Go package for shared utilities"
      icon="folder-utils"
    />
  <FileIcon
    name="web/"
    description="Vue based Web UI"
    icon="folder-public"
  />
  <FileIcon
    name="Dockerfile"
    description="Production dockerfile that builds both the server and UI"
    icon="docker"
  />
  <FileIcon
    name="Dockerfile.dev"
    description="Dockerfile that only builds the server, and excludes the UI for development"
    icon="docker"
  />
  <FileIcon
    name="Makefile"
    description="Contains project scripts for building and running all the different parts of Miasma"
  />
  <FileIcon
    name="meta.json"
    description="Contains metadata about the current version of Miasma"
  />
</pre>

:::tip
Hover over a folder or file to see it's description
:::

## Running Locally

To develop and run Miasma locally, it's best to run each "part" individually. This will give you the fastest start times and hot reloading (for most of the apps).

### GraphQL API

```bash:no-line-numbers
make run
```

This runs the Miasma Server in development mode on port `3000`, excluding the web UI. Including the web UI in this process really slows down this command.

**This is the only dev command does not reload on save**. You have to stop and start it manually.

The app is written with [Go](https://go.dev) and [GQLGen](https://gqlgen.com/).

#### Changing the GraphQL Schema

To make changes to the schema, edit any of the schema files in the `api/` directory, and run:

```bash:no-line-numbers
make gen
```

This will regenerate all the GQLGen files and leave queries/mutations as `panic("not implemented")`.

### Web UI

```bash:no-line-numbers
make ui
```

This runs the UI in watch mode on port `8080`.

Normally the UI just uses `/graphql` as the API endpoint, but during development when the API is ran separately, it uses `http://localhost:3000/graphql`.

:::warning
Make sure you have the API running so the UI has something to communicate with.
:::

The app is written in [Vue](https://v2.vuepress.vuejs.org/) and bundled by [Vite](https://vitejs.dev/).

### Docs

```bash:no-line-numbers
make doc
```

This runs the docs website on port `8081`.

The website is made using [VuePress](https://v2.vuepress.vuejs.org/), and published to [GitHub pages](https://pages.github.com/).

### CLI

```bash:no-line-numbers
# Compile and install to your $PATH
make cli

# Then use the CLI
miasma-dev connect localhost:3000
miasma-dev apps
```

Make sure you've set the `$GOPATH` environment varaible and it's on your `$PATH`. If `$GOPATH` is not set, the command will fail. If `$GOPATH` isn't in your `$PATH`, then your shell won't be able to locate the `miasma-dev` command.

:::warning
If you have the CLI installed already, `miasma-dev` will use the same `~/.miasma.yml` config file.
:::

See the [`cli` target in the `Makefile`](https://github.com/aklinker1/miasma/blob/74b1d25009432262112d1627c0bdd69d46826722/Makefile#L56-L61) for more details.

The CLI is written with [Go](https://go.dev), [`spf13/cobra`](https://cobra.dev/), and [`spf13/viper`](https://github.com/spf13/viper).

## Building for Production

### Docker Image

```bash:no-line-numbers
make build
```

This builds the production docker image, `aklinker1/miasma`. This image includes the GraphQL API and Web UI.

```bash:no-line-numbers
make preview
```

This will build and run the production docker image. This is slower than running the GraphQL API and Web UI part individually while during development, but is useful to verify they are integrated correctly.

### Docs

```bash:no-line-numbers
make docs
```

This will bundle the VuePress docs website to `docs/docs/.vuepress/dist`.

### CLI

:::danger TODO
:::

## Contributors

<br />

<a href="https://github.com/aklinker1/miasma/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=aklinker1/miasma" />
</a>

Special thanks to myself for being the only contributor (thus far)! :tada:
