---
title: Installation
---

# Installation

[[toc]]

## Deploy the Miasma Server

Run the following commands on the main machine you want to host applications on. If you are planning on adding more nodes to the cluster, this device will be the main manager node.

:::tip Machine OS
Before starting, make sure the machine you're installing the server on is one of the [supported architectures](https://hub.docker.com/r/aklinker1/miasma/tags).

For example, `linux/arm64` is published, so if you want to install the server on a Raspberry Pi, it needs to be running a 64bit OS, like Ubuntu Server 64bit.
:::

1. [Install Docker](https://docs.docker.com/engine/install/).
1. Initialize the swarm (**this is required** even if you aren't planning on using multiple nodes)
   ```bash:no-line-numbers
   docker swarm init
   ```
1. Run the server's Docker image
   ```bash:no-line-numbers
   # Minimal configuration
   docker run -d \
       --restart unless-stopped \
       -p 3000:3000 \
       -v /var/run/docker.sock:/var/run/docker.sock \
       -v $HOME/.miasma:/data/miasma \
       aklinker1/miasma
   ```
   :::details What are all those parameters for?
   - **`-d`**
   - **`--restart unless-stopped`** - Always restart the container (on reboot, crash, etc) until you explicitly call `docker container stop <id>`
   - **`-v /var/run/docker.sock:/var/run/docker.sock`** - Bind a special volume that lets Miasma talk to the docker daemon running on the host machine so it can do things like start/stop apps
   - **`-v $HOME/.miasma:/data/miasma`** - Bind a volume so Miasma can persist it's configuration and remember what apps have been setup if the container is stopped and restarted
   - [`aklinker1/miasma`](https://hub.docker.com/r/aklinker1/miasma) - The name of the docker image to run. To use a different version of the image, include a `:tag` suffix
     :::

And you're server is setup!

:::details Verify Setup
Open up the GraphQL playground, `http://<machine-ip>:3000/playground`, and run the following query:

```graphql:no-line-numbers
{
  health {
    version
    dockerVersion
    cluster {
      joinCommand
      createdAt
    }
  }
}
```

If everything is setup correctly, you should see a response like this:

```json:no-line-numbers
{
  "data": {
    "health": {
      "version": "1.1.0",
      "dockerVersion": "20.10.17",
      "cluster": {
        "joinCommand": "docker swarm join --token SWMTKN-1-1hgnz90jzs1xytc8chcxdlcqn1r1p38xtlus4mzyr7wx3lnuie-ejy87vthuum7hc36cw0xr4wjq 192.168.0.3:2377",
        "createdAt": "2022-05-05T07:11:01.82641567Z"
      }
    }
  }
}
```

:::

### Add More Nodes (Optional)

1. SSH into the machine to add
   ```bash:no-line-numbers
   ssh user@<machine-ip>
   ```
2. [Install Docker](https://docs.docker.com/engine/install/) (same as before).
3. Join the swarm using the join command output during `docker swarm init`
   ```bash:no-line-numbers
   docker swarm join --token <some-token> <swarm-ip>:2377
   ```

## Install the CLI

<script setup>
  import { ref, computed, watch } from 'vue';

  const cliVersion = __CLI_VERSION__;
  const arches = {
    linux: [
      { display: 'amd64/x86_64', value: 'x86_64' },
      { display: 'arm64/aarch64', value: 'aarch64' },
      { display: 'armv6', value: 'armv6' },
      { display: 'armv7', value: 'armv7' },
      { display: 'ppc64le', value: 'ppc64le' },
      { display: 's390x', value: 's390x' },
    ],
    windows: [
      { display: "amd64/x86_64", value: "x86_64.exe" },
    ],
    darwin: [
      { display: "amd64/x86_64", value: "x86_64" },
      { display: "arm64/aarch64", value: "aarch64" },
    ],
  }
  const binaries = {
    linux: "/usr/local/bin/miasma",
    windows: "C:\\\"Program Files\"\\miasma\\miasma.exe",
    darwin: "/usr/local/bin/miasma",
  }
  
  const os = ref("linux");
  const arch = ref("amd64");
  const archOptions = computed(() => arches[os.value])
  const binary = computed(() => binaries[os.value])

  watch(os, (newOs) => {
    arch.value = arches[newOs][0].value
  }, { immediate: true })
</script>

The CLI is a standalone binary, so just download the [latest version from GitHub](https://github.com/aklinker1/miasma/releases?q=cli), verify it's checksum, and add it to your path.

<div>
  <label>
    <span><strong>OS:{{' '}}</strong></span>
    <select placeholder="OS" v-model="os">
      <option value="linux">Linux</option>
      <option value="darwin">Mac</option>
      <option value="windows">Windows</option>
    </select>
  </label>
  &emsp;
  <label>
    <span><strong>Arch:{{' '}}</strong></span>
    <select placeholder="Architecture" v-model="arch">
      <option v-for="o of archOptions" :key="o.value" :value="o.value">{{o.display}}</option>
    </select>
  </label>
</div>

<div class="language-bash ext-sh"><pre class="language-bash"><code><span class="token comment"># Download</span>
<span class="token function">curl</span> -L -O https://github.com/aklinker1/miasma/releases/download/cli-v{{cliVersion}}/miasma-cli-{{os}}-{{arch}}
<span class="token function">curl</span> -L -O https://github.com/aklinker1/miasma/releases/download/cli-v{{cliVersion}}/miasma-cli-{{os}}-{{arch}}.sha256
<span></span>
<span class="token comment"># Verify</span>
<span>sha256sum --check miasma-cli-{{os}}-{{arch}}.sha256</span>
<span></span>
<span class="token comment"># Add to Path</span>
<span class="token function">sudo mv</span> miasma-cli-{{os}}-{{arch}} {{binary}}
<template v-if="os !== 'windows'"><span class="token function">sudo chmod</span> +x {{binary}}</template>
</code></pre></div>

To upgrade to a newer version of the CLI, just run the same commands again using the newer version's tag.
