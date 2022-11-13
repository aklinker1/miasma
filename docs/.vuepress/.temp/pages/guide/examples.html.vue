<template><div><h1 id="example-apps" tabindex="-1"><a class="header-anchor" href="#example-apps" aria-hidden="true">#</a> Example Apps</h1>
<p>Here are some example apps setup and configured through the <RouterLink to="/reference/cli.html">Miasma CLI</RouterLink>.</p>
<nav class="table-of-contents"><ul><li><RouterLink to="#docker-registry">Docker Registry</RouterLink></li><li><RouterLink to="#mongodb">MongoDB</RouterLink><ul><li><RouterLink to="#mongo-express">Mongo Express</RouterLink></li></ul></li><li><RouterLink to="#redis">Redis</RouterLink></li><li><RouterLink to="#postgres">Postgres</RouterLink></li><li><RouterLink to="#nginx-based-website">Nginx Based Website</RouterLink></li></ul></nav>
<h2 id="docker-registry" tabindex="-1"><a class="header-anchor" href="#docker-registry" aria-hidden="true">#</a> Docker Registry</h2>
<p>Host an insecure <a href="https://hub.docker.com/_/registry/" target="_blank" rel="noopener noreferrer">docker registry<ExternalLinkIcon/></a>.</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code><span class="token comment"># Create the app</span>
miasma apps:create docker-registry -i registry:2.7

<span class="token comment"># Configure it's ports, node placement, and mounted volume for data storage</span>
miasma apps:edit -a docker-registry <span class="token punctuation">\</span>
    --add-target-ports <span class="token number">5000</span> <span class="token punctuation">\</span>
    --add-placement-constraint <span class="token string">"node.labels.database == true"</span> <span class="token punctuation">\</span>
    --add-volume /dir/path/on/physical/machine:/var/lib/registry

<span class="token comment"># Give it a route</span>
miasma route:set --hostname docker.hostname.com
</code></pre></div><p>Because this is an insecure registry, you have to <a href="https://docs.docker.com/registry/insecure/" target="_blank" rel="noopener noreferrer">configure docker to allow it<ExternalLinkIcon/></a> by editing the Docker daemon config:</p>
<div class="language-json ext-json line-numbers-mode"><pre v-pre class="language-json"><code><span class="token comment">// /etc/docker/daemon.json</span>
<span class="token punctuation">{</span>
  <span class="token property">"insecure-registries"</span><span class="token operator">:</span> <span class="token punctuation">[</span><span class="token string">"docker.hostname.com"</span><span class="token punctuation">]</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Then restart the computer so the changes take effect.</p>
<p>If you want to use the registry as a source for Miasma apps, each node also needs to update their daemon config and restart.</p>
<h2 id="mongodb" tabindex="-1"><a class="header-anchor" href="#mongodb" aria-hidden="true">#</a> MongoDB</h2>
<p>Create the <a href="https://hub.docker.com/_/mongo/" target="_blank" rel="noopener noreferrer">MongoDB<ExternalLinkIcon/></a> instance:</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code>miasma apps:create mongo -i mongo:4
miasma env:edit -a mongo
miasma apps:edit <span class="token punctuation">\</span>
    --add-placement-constraint <span class="token string">"node.labels.database == true"</span> <span class="token punctuation">\</span>
    --add-volume <span class="token string">"/dir/path/on/physical/machine"</span>:/data/db <span class="token punctuation">\</span>
    --add-target-ports <span class="token number">27017</span> --add-published-ports <span class="token number">27017</span>
</code></pre></div><div class="custom-container tip"><p class="custom-container-title">TIP</p>
<p>Mongo can take a few minutes to startup, so be patient. If it isn't running, you can troubleshoot by SSHing onto the manager node and running <code v-pre>docker service ps miasma-mongo --no-trunc</code></p>
</div>
<h3 id="mongo-express" tabindex="-1"><a class="header-anchor" href="#mongo-express" aria-hidden="true">#</a> Mongo Express</h3>
<p>Optionally, you can also setup <a href="https://hub.docker.com/_/mongo-express/" target="_blank" rel="noopener noreferrer">Mongo Express<ExternalLinkIcon/></a> to get a UI for inspecting your mongo app:</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code>miasma apps:create mongo-express -i mongo-express
miasma env:edit -a mongo-express
miasma apps:edit --add-target-ports <span class="token number">8081</span>
</code></pre></div><h2 id="redis" tabindex="-1"><a class="header-anchor" href="#redis" aria-hidden="true">#</a> Redis</h2>
<p>Create the <a href="https://hub.docker.com/_/redis/" target="_blank" rel="noopener noreferrer">Redis<ExternalLinkIcon/></a> instance. Nothing special here, Redis respects the <code v-pre>PORT</code> environment variable.</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code>miasma apps:create redis -i redis:6-alpine
</code></pre></div><h2 id="postgres" tabindex="-1"><a class="header-anchor" href="#postgres" aria-hidden="true">#</a> Postgres</h2>
<p>Create the <a href="https://hub.docker.com/_/postgres/" target="_blank" rel="noopener noreferrer">PostgreSQL<ExternalLinkIcon/></a> database. The Postgres docker image can be quite finicky with it's port, so it's best to just set the target and published ports to <code v-pre>5432</code>, Postgres's default.</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code>miasma apps:create postgres -i postgres:13-alpine
miasma env:edit -a postgres
miasma apps:edit <span class="token punctuation">\</span>
    --add-placement-constraint <span class="token string">"node.labels.database == true"</span> <span class="token punctuation">\</span>
    --add-volume <span class="token string">"/dir/path/on/physical/machine"</span>:/var/lib/postgresql/data <span class="token punctuation">\</span>
    --add-target-ports <span class="token number">5432</span> --add-published-ports <span class="token number">5432</span>
</code></pre></div><h2 id="nginx-based-website" tabindex="-1"><a class="header-anchor" href="#nginx-based-website" aria-hidden="true">#</a> Nginx Based Website</h2>
<p>A common way to host a static website in Docker is with the <a href="https://hub.docker.com/_/nginx/" target="_blank" rel="noopener noreferrer"><code v-pre>nginx</code><ExternalLinkIcon/></a> image.</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code>miasma apps:create example-website -i nginxdemos/hello:plain-text <span class="token comment"># or your custom image</span>
<span class="token comment"># Nginx runs on port 80, and is difficult to change, so just target that port instead</span>
miasma apps:edit -a example-website --add-target-ports <span class="token number">80</span>
</code></pre></div></div></template>
