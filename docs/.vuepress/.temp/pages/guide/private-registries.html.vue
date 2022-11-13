<template><div><h1 id="private-registries" tabindex="-1"><a class="header-anchor" href="#private-registries" aria-hidden="true">#</a> Private Registries</h1>
<p>Miasma is based off Docker Swarm, and Docker Swarm supports private registries, so that means <strong>Miasma supports private registries as well</strong>!</p>
<h3 id="how-does-it-work-in-docker-swarm" tabindex="-1"><a class="header-anchor" href="#how-does-it-work-in-docker-swarm" aria-hidden="true">#</a> How does it work in Docker Swarm?</h3>
<p>Normally, you log into the private registries on your cluster's manager nodes, and use the <code v-pre>--with-registry-auth</code> flag when creating a &quot;service&quot; (Swarm service = Miasma app).</p>
<p>See Docker's documentation for more details: <a href="https://docs.docker.com/engine/swarm/services/#create-a-service-using-an-image-on-a-private-registry" target="_blank" rel="noopener noreferrer">https://docs.docker.com/engine/swarm/services/#create-a-service-using-an-image-on-a-private-registry<ExternalLinkIcon/></a></p>
<h3 id="how-does-it-work-in-miasma" tabindex="-1"><a class="header-anchor" href="#how-does-it-work-in-miasma" aria-hidden="true">#</a> How does it work in Miasma?</h3>
<p>Currently, Miasma does not support a <code v-pre>--with-registry-auth</code> equivalent since it uses the Docker Engine API to interact with docker, not the CLI.</p>
<blockquote>
<p>I haven't figured out how to pass this flag, so if you have an idea, open a PR!</p>
</blockquote>
<p>So instead of just logging into the private registry on your manager nodes, you need to do it on all nodes in your cluster.</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code><span class="token function">ssh</span> username@<span class="token operator">&lt;</span>machine-ip<span class="token operator">></span>
<span class="token function">docker</span> login myprivateregistry.com
<span class="token builtin class-name">exit</span>

<span class="token function">ssh</span> username@<span class="token operator">&lt;</span>machine-2-ip<span class="token operator">></span>
<span class="token function">docker</span> login myprivateregistry.com
<span class="token builtin class-name">exit</span>

<span class="token function">ssh</span> username@<span class="token operator">&lt;</span>machine-3-ip<span class="token operator">></span>
<span class="token comment"># ...</span>
</code></pre></div><p>After you've logged into the registry on all the nodes, you can create an app from the CLI or UI using the <code v-pre>myprivateregistry.com/some/image:tag</code> image name:</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code>miasma apps:create Example -i myprivateregistry.com/some/image:tag
</code></pre></div></div></template>
