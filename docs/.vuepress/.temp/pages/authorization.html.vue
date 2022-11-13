<template><div><h1 id="authorization" tabindex="-1"><a class="header-anchor" href="#authorization" aria-hidden="true">#</a> Authorization</h1>
<p>Miasma supports the following types of authorization:</p>
<ul>
<li>None (default)</li>
<li>Hard-coded Access Token</li>
</ul>
<h2 id="none" tabindex="-1"><a class="header-anchor" href="#none" aria-hidden="true">#</a> None</h2>
<p>Miasma will not require any authorization by default. This is ideal for testing Miasma out on a single device or for a local network that you're OK with allowing access to everyone.</p>
<h2 id="hard-coded-access-token" tabindex="-1"><a class="header-anchor" href="#hard-coded-access-token" aria-hidden="true">#</a> Hard-coded Access Token</h2>
<p>Require all requests to include a hard coded access token. This is ideal for a local network or a cloud hosted instance of Miasma.</p>
<div class="custom-container warning"><p class="custom-container-title">Token Requirements</p>
<p>There are no requirements for length or complexity. If you don't want to be hacked, make sure the token is sufficiently long.</p>
</div>
<p>When starting up the server, you can set the <code v-pre>ACCESS_TOKEN</code> environment variable:</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code>$ <span class="token function">cat</span> .env
<span class="token assign-left variable">ACCESS_TOKEN</span><span class="token operator">=</span><span class="token operator">&lt;</span>token<span class="token operator">></span>

$ <span class="token function">docker</span> run --env-file .env <span class="token punctuation">\</span>
    -d <span class="token punctuation">\</span>
    --restart unless-stopped <span class="token punctuation">\</span>
    -p <span class="token number">3000</span>:3000 <span class="token punctuation">\</span>
    -v /var/run/docker.sock:/var/run/docker.sock <span class="token punctuation">\</span>
    -v <span class="token environment constant">$HOME</span>/.miasma:/data/miasma <span class="token punctuation">\</span>
    aklinker1/miasma
</code></pre></div><p>All requests to the API will then require an authorization header:</p>
<div class="language-text ext-text"><pre v-pre class="language-text"><code>Authorization: Bearer &lt;token>
</code></pre></div><p>For the CLI, you'll need to reconnect and pass the token via the <code v-pre>--auth</code> flag:</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code>miasma connect <span class="token operator">&lt;</span>ip:port<span class="token operator">></span> --auth <span class="token operator">&lt;</span>token<span class="token operator">></span>
</code></pre></div><p>And finally, the UI will require you to log in using the same token.</p>
</div></template>
