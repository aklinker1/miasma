<template><div><h1 id="contributing" tabindex="-1"><a class="header-anchor" href="#contributing" aria-hidden="true">#</a> Contributing</h1>
<p>Welcome to the contributor's guide! Here you'll find everything you need to know about contributing to Miasma.</p>
<nav class="table-of-contents"><ul><li><RouterLink to="#required-tools">Required Tools</RouterLink></li><li><RouterLink to="#project-structure">Project Structure</RouterLink></li><li><RouterLink to="#running-locally">Running Locally</RouterLink><ul><li><RouterLink to="#graphql-api">GraphQL API</RouterLink></li><li><RouterLink to="#web-ui">Web UI</RouterLink></li><li><RouterLink to="#docs">Docs</RouterLink></li><li><RouterLink to="#cli">CLI</RouterLink></li></ul></li><li><RouterLink to="#building-for-production">Building for Production</RouterLink><ul><li><RouterLink to="#docker-image">Docker Image</RouterLink></li><li><RouterLink to="#docs-1">Docs</RouterLink></li><li><RouterLink to="#cli-1">CLI</RouterLink></li></ul></li><li><RouterLink to="#contributors">Contributors</RouterLink></li></ul></nav>
<h2 id="required-tools" tabindex="-1"><a class="header-anchor" href="#required-tools" aria-hidden="true">#</a> Required Tools</h2>
<ul>
<li><a href="https://www.gnu.org/software/make/" target="_blank" rel="noopener noreferrer">GNU <code v-pre>make</code><ExternalLinkIcon/></a> to run scripts</li>
<li><a href="https://go.dev/doc/install" target="_blank" rel="noopener noreferrer">Go 1.19<ExternalLinkIcon/></a> to build the CLI</li>
<li><a href="https://nodejs.org/en/" target="_blank" rel="noopener noreferrer">Node 16 LTS<ExternalLinkIcon/></a> to run the UI in dev mode</li>
<li><a href="https://docs.docker.com/get-docker/" target="_blank" rel="noopener noreferrer">Docker<ExternalLinkIcon/></a> to build the server and production docker image</li>
</ul>
<h2 id="project-structure" tabindex="-1"><a class="header-anchor" href="#project-structure" aria-hidden="true">#</a> Project Structure</h2>
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
<div class="custom-container tip"><p class="custom-container-title">TIP</p>
<p>Hover over a folder or file to see it's description</p>
</div>
<h2 id="running-locally" tabindex="-1"><a class="header-anchor" href="#running-locally" aria-hidden="true">#</a> Running Locally</h2>
<p>To develop and run Miasma locally, it's best to run each &quot;part&quot; individually. This will give you the fastest start times and hot reloading (for most of the apps).</p>
<h3 id="graphql-api" tabindex="-1"><a class="header-anchor" href="#graphql-api" aria-hidden="true">#</a> GraphQL API</h3>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code><span class="token function">make</span> run
</code></pre></div><p>This runs the Miasma Server in development mode on port <code v-pre>3000</code>, excluding the web UI. Including the web UI in this process really slows down this command.</p>
<p><strong>This is the only dev command does not reload on save</strong>. You have to stop and start it manually.</p>
<p>The app is written with <a href="https://go.dev" target="_blank" rel="noopener noreferrer">Go<ExternalLinkIcon/></a> and <a href="https://gqlgen.com/" target="_blank" rel="noopener noreferrer">GQLGen<ExternalLinkIcon/></a>.</p>
<h4 id="changing-the-graphql-schema" tabindex="-1"><a class="header-anchor" href="#changing-the-graphql-schema" aria-hidden="true">#</a> Changing the GraphQL Schema</h4>
<p>To make changes to the schema, edit any of the schema files in the <code v-pre>api/</code> directory, and run:</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code><span class="token function">make</span> gen
</code></pre></div><p>This will regenerate all the GQLGen files and leave queries/mutations as <code v-pre>panic(&quot;not implemented&quot;)</code>.</p>
<h3 id="web-ui" tabindex="-1"><a class="header-anchor" href="#web-ui" aria-hidden="true">#</a> Web UI</h3>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code><span class="token function">make</span> ui
</code></pre></div><p>This runs the UI in watch mode on port <code v-pre>8080</code>.</p>
<p>Normally the UI just uses <code v-pre>/graphql</code> as the API endpoint, but during development when the API is ran separately, it uses <code v-pre>http://localhost:3000/graphql</code>.</p>
<div class="custom-container warning"><p class="custom-container-title">WARNING</p>
<p>Make sure you have the API running so the UI has something to communicate with.</p>
</div>
<p>The app is written in <a href="https://v2.vuepress.vuejs.org/" target="_blank" rel="noopener noreferrer">Vue<ExternalLinkIcon/></a> and bundled by <a href="https://vitejs.dev/" target="_blank" rel="noopener noreferrer">Vite<ExternalLinkIcon/></a>.</p>
<h3 id="docs" tabindex="-1"><a class="header-anchor" href="#docs" aria-hidden="true">#</a> Docs</h3>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code><span class="token function">make</span> doc
</code></pre></div><p>This runs the docs website on port <code v-pre>8081</code>.</p>
<p>The website is made using <a href="https://v2.vuepress.vuejs.org/" target="_blank" rel="noopener noreferrer">VuePress<ExternalLinkIcon/></a>, and published to <a href="https://pages.github.com/" target="_blank" rel="noopener noreferrer">GitHub pages<ExternalLinkIcon/></a>.</p>
<h3 id="cli" tabindex="-1"><a class="header-anchor" href="#cli" aria-hidden="true">#</a> CLI</h3>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code><span class="token comment"># Compile and install to your $PATH</span>
<span class="token function">make</span> cli

<span class="token comment"># Then use the CLI</span>
miasma-dev connect localhost:3000
miasma-dev apps
</code></pre></div><p>Make sure you've set the <code v-pre>$GOPATH</code> environment varaible and it's on your <code v-pre>$PATH</code>. If <code v-pre>$GOPATH</code> is not set, the command will fail. If <code v-pre>$GOPATH</code> isn't in your <code v-pre>$PATH</code>, then your shell won't be able to locate the <code v-pre>miasma-dev</code> command.</p>
<div class="custom-container warning"><p class="custom-container-title">WARNING</p>
<p>If you have the CLI installed already, <code v-pre>miasma-dev</code> will use the same <code v-pre>~/.miasma.yml</code> config file.</p>
</div>
<p>See the <a href="https://github.com/aklinker1/miasma/blob/74b1d25009432262112d1627c0bdd69d46826722/Makefile#L56-L61" target="_blank" rel="noopener noreferrer"><code v-pre>cli</code> target in the <code v-pre>Makefile</code><ExternalLinkIcon/></a> for more details.</p>
<p>The CLI is written with <a href="https://go.dev" target="_blank" rel="noopener noreferrer">Go<ExternalLinkIcon/></a>, <a href="https://cobra.dev/" target="_blank" rel="noopener noreferrer"><code v-pre>spf13/cobra</code><ExternalLinkIcon/></a>, and <a href="https://github.com/spf13/viper" target="_blank" rel="noopener noreferrer"><code v-pre>spf13/viper</code><ExternalLinkIcon/></a>.</p>
<h2 id="building-for-production" tabindex="-1"><a class="header-anchor" href="#building-for-production" aria-hidden="true">#</a> Building for Production</h2>
<h3 id="docker-image" tabindex="-1"><a class="header-anchor" href="#docker-image" aria-hidden="true">#</a> Docker Image</h3>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code><span class="token function">make</span> build
</code></pre></div><p>This builds the production docker image, <code v-pre>aklinker1/miasma</code>. This image includes the GraphQL API and Web UI.</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code><span class="token function">make</span> preview
</code></pre></div><p>This will build and run the production docker image. This is slower than running the GraphQL API and Web UI part individually while during development, but is useful to verify they are integrated correctly.</p>
<h3 id="docs-1" tabindex="-1"><a class="header-anchor" href="#docs-1" aria-hidden="true">#</a> Docs</h3>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code><span class="token function">make</span> docs
</code></pre></div><p>This will bundle the VuePress docs website to <code v-pre>docs/.vuepress/dist</code>.</p>
<h3 id="cli-1" tabindex="-1"><a class="header-anchor" href="#cli-1" aria-hidden="true">#</a> CLI</h3>
<div class="custom-container danger"><p class="custom-container-title">TODO</p>
</div>
<h2 id="contributors" tabindex="-1"><a class="header-anchor" href="#contributors" aria-hidden="true">#</a> Contributors</h2>
<br />
<a href="https://github.com/aklinker1/miasma/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=aklinker1/miasma" />
</a>
<p>Special thanks to myself for being the only contributor (thus far)! 🎉</p>
</div></template>

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
