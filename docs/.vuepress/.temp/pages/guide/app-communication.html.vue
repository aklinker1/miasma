<template><div><h1 id="app-communication" tabindex="-1"><a class="header-anchor" href="#app-communication" aria-hidden="true">#</a> App Communication</h1>
<p>When multiple apps are running, they can communicate without leaving the cluster's network via <code v-pre>&lt;protocol&gt;://&lt;app-name&gt;:&lt;target-port&gt;</code>.</p>
<p>For example, lets say you have a Postgres database and a web app that needs to access it.</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code>miasma apps:create postgres -i postgres-alpine
miasma apps:edit -a postgres --add-target-ports <span class="token number">5432</span>

miasma apps:create web-app -i <span class="token operator">&lt;</span>my-custom-image<span class="token operator">></span>
</code></pre></div><p>You should use <code v-pre>postgres://postgres:5432</code> when connecting to the database.</p>
<div class="custom-container tip"><p class="custom-container-title">App name transformation</p>
<p>The hostname of an app will the the name, but lower-cased and with dashes instead of spaces:</p>
<ul>
<li><code v-pre>Postgres 1</code> → <code v-pre>postgres-1</code></li>
<li><code v-pre>Personal Website</code> → <code v-pre>personal-website</code></li>
<li><code v-pre>Mongo</code> → <code v-pre>mongo</code></li>
</ul>
</div>
</div></template>
