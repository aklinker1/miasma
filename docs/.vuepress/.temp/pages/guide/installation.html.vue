<template><div><h1 id="installation" tabindex="-1"><a class="header-anchor" href="#installation" aria-hidden="true">#</a> Installation</h1>
<nav class="table-of-contents"><ul><li><RouterLink to="#deploy-the-miasma-server">Deploy the Miasma Server</RouterLink><ul><li><RouterLink to="#add-more-nodes-optional">Add More Nodes (Optional)</RouterLink></li></ul></li><li><RouterLink to="#install-the-cli">Install the CLI</RouterLink></li></ul></nav>
<h2 id="deploy-the-miasma-server" tabindex="-1"><a class="header-anchor" href="#deploy-the-miasma-server" aria-hidden="true">#</a> Deploy the Miasma Server</h2>
<p>Run the following commands on the main machine you want to host applications on. If you are planning on adding more nodes to the cluster, this device will be the main manager node.</p>
<div class="custom-container tip"><p class="custom-container-title">Machine OS</p>
<p>Before starting, make sure the machine you're installing the server on is one of the <a href="https://hub.docker.com/r/aklinker1/miasma/tags" target="_blank" rel="noopener noreferrer">supported architectures<ExternalLinkIcon/></a>.</p>
<p>For example, <code v-pre>linux/arm64</code> is published, so if you want to install the server on a Raspberry Pi, it needs to be running a 64bit OS, like Ubuntu Server 64bit.</p>
</div>
<ol>
<li><a href="https://docs.docker.com/get-docker/" target="_blank" rel="noopener noreferrer">Install Docker<ExternalLinkIcon/></a>. On linux, I recommend using <code v-pre>get.docker.com</code>:<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code><span class="token function">curl</span> -fsSL https://get.docker.com -o get-docker.sh
<span class="token function">sh</span> get-docker.sh
</code></pre></div></li>
<li>Initialize the swarm (<strong>this is required</strong> even if you aren't planning on using multiple nodes)<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code><span class="token function">docker</span> swarm init
</code></pre></div></li>
<li>Run the server's Docker image<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code><span class="token comment"># Minimal configuration</span>
<span class="token function">docker</span> run -d <span class="token punctuation">\</span>
    --restart unless-stopped <span class="token punctuation">\</span>
    -p <span class="token number">3000</span>:3000 <span class="token punctuation">\</span>
    -v /var/run/docker.sock:/var/run/docker.sock <span class="token punctuation">\</span>
    -v <span class="token environment constant">$HOME</span>/.miasma:/data/miasma <span class="token punctuation">\</span>
    aklinker1/miasma
</code></pre></div><details class="custom-container details"><summary>What are all those parameters for?</summary>
<ul>
<li><strong><code v-pre>-d</code></strong></li>
<li><strong><code v-pre>--restart unless-stopped</code></strong> - Always restart the container (on reboot, crash, etc) until you explicitly call <code v-pre>docker container stop &lt;id&gt;</code></li>
<li><strong><code v-pre>-v /var/run/docker.sock:/var/run/docker.sock</code></strong> - Bind a special volume that lets Miasma talk to the docker daemon running on the host machine so it can do things like start/stop apps</li>
<li><strong><code v-pre>-v $HOME/.miasma:/data/miasma</code></strong> - Bind a volume so Miasma can persist it's configuration and remember what apps have been setup if the container is stopped and restarted</li>
<li><a href="https://hub.docker.com/r/aklinker1/miasma" target="_blank" rel="noopener noreferrer"><code v-pre>aklinker1/miasma</code><ExternalLinkIcon/></a> - The name of the docker image to run. To use a different version of the image, include a <code v-pre>:tag</code> suffix</li>
</ul>
</details>
</li>
</ol>
<p>And you're server is setup!</p>
<details class="custom-container details"><summary>Verify Setup</summary>
<p>Open up the GraphQL playground, <code v-pre>http://&lt;machine-ip&gt;:3000/playground</code>, and run the following query:</p>
<div class="language-graphql ext-graphql"><pre v-pre class="language-graphql"><code><span class="token punctuation">{</span>
  <span class="token object">health</span> <span class="token punctuation">{</span>
    <span class="token property">version</span>
    <span class="token property">dockerVersion</span>
    <span class="token object">cluster</span> <span class="token punctuation">{</span>
      <span class="token property">joinCommand</span>
      <span class="token property">createdAt</span>
    <span class="token punctuation">}</span>
  <span class="token punctuation">}</span>
<span class="token punctuation">}</span>
</code></pre></div><p>If everything is setup correctly, you should see a response like this:</p>
<div class="language-json ext-json"><pre v-pre class="language-json"><code><span class="token punctuation">{</span>
  <span class="token property">"data"</span><span class="token operator">:</span> <span class="token punctuation">{</span>
    <span class="token property">"health"</span><span class="token operator">:</span> <span class="token punctuation">{</span>
      <span class="token property">"version"</span><span class="token operator">:</span> <span class="token string">"1.1.0"</span><span class="token punctuation">,</span>
      <span class="token property">"dockerVersion"</span><span class="token operator">:</span> <span class="token string">"20.10.17"</span><span class="token punctuation">,</span>
      <span class="token property">"cluster"</span><span class="token operator">:</span> <span class="token punctuation">{</span>
        <span class="token property">"joinCommand"</span><span class="token operator">:</span> <span class="token string">"docker swarm join --token SWMTKN-1-1hgnz90jzs1xytc8chcxdlcqn1r1p38xtlus4mzyr7wx3lnuie-ejy87vthuum7hc36cw0xr4wjq 192.168.0.3:2377"</span><span class="token punctuation">,</span>
        <span class="token property">"createdAt"</span><span class="token operator">:</span> <span class="token string">"2022-05-05T07:11:01.82641567Z"</span>
      <span class="token punctuation">}</span>
    <span class="token punctuation">}</span>
  <span class="token punctuation">}</span>
<span class="token punctuation">}</span>
</code></pre></div></details>
<h3 id="add-more-nodes-optional" tabindex="-1"><a class="header-anchor" href="#add-more-nodes-optional" aria-hidden="true">#</a> Add More Nodes (Optional)</h3>
<ol>
<li>SSH into the machine to add</li>
</ol>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code><span class="token function">ssh</span> user@<span class="token operator">&lt;</span>machine-ip<span class="token operator">></span>
</code></pre></div><ol start="2">
<li>Install Docker (same as before)</li>
</ol>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code><span class="token function">curl</span> -fsSL https://get.docker.com -o get-docker.sh
<span class="token function">sh</span> get-docker.sh
</code></pre></div><ol start="3">
<li>Join the swarm using the join command output during <code v-pre>docker swarm init</code></li>
</ol>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code><span class="token function">docker</span> swarm <span class="token function">join</span> --token <span class="token operator">&lt;</span>some-token<span class="token operator">></span> <span class="token operator">&lt;</span>swarm-ip<span class="token operator">></span>:2377
</code></pre></div><h2 id="install-the-cli" tabindex="-1"><a class="header-anchor" href="#install-the-cli" aria-hidden="true">#</a> Install the CLI</h2>
<p>The CLI is a standalone binary, so just download the <a href="https://github.com/aklinker1/miasma/releases?q=cli" target="_blank" rel="noopener noreferrer">latest version from GitHub<ExternalLinkIcon/></a>, verify it's checksum, and add it to your path.</p>
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
<p>To upgrade to a newer version of the CLI, just run the same commands again using the newer version's tag.</p>
</div></template>

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
