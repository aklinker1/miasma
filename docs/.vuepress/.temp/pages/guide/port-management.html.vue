<template><div><h1 id="port-management" tabindex="-1"><a class="header-anchor" href="#port-management" aria-hidden="true">#</a> Port Management</h1>
<p><a href="https://devcenter.heroku.com/articles/runtime-principles#web-servers" target="_blank" rel="noopener noreferrer">Like Heroku<ExternalLinkIcon/></a>, Miasma will automatically manage ports for you. But unlike Heroku, you can override the app's <strong>target</strong> and <strong>published</strong> ports.</p>
<details class="custom-container details"><summary>About Docker Ports</summary>
<p>Docker (and thus Miasma) has two types of ports:</p>
<ol>
<li><strong>Target</strong>: the port inside the container</li>
<li><strong>Published</strong>: the publicly exposed port you would use to interact with the app</li>
</ol>
<p>For more info, see Docker's <a href="https://docs.docker.com/compose/compose-file/compose-file-v3/#ports" target="_blank" rel="noopener noreferrer">official documentation<ExternalLinkIcon/></a></p>
</details>
<p>By default, a random target port will be selected and provided to the app via the <code v-pre>PORT</code> environment variable. The app should respect this environment variable, and run on the provided port.</p>
<h2 id="published-ports" tabindex="-1"><a class="header-anchor" href="#published-ports" aria-hidden="true">#</a> Published Ports</h2>
<p>Usually you don't need to set the published port of an application unless you want to, or the image requires it.</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code><span class="token comment"># Add a port</span>
miasma apps:edit -a postgres --add-published-ports <span class="token number">5432</span>

<span class="token comment"># Go back to a randomized port</span>
miasma apps:edit -a postgres --rm-published-ports <span class="token number">5432</span>
</code></pre></div><h2 id="target-ports" tabindex="-1"><a class="header-anchor" href="#target-ports" aria-hidden="true">#</a> Target Ports</h2>
<p>Sometimes it's difficult to configure a docker image to respect the <code v-pre>$PORT</code> environment variable. If you can't change which port the application runs on, or it's difficult to do, you can just tell the app to target a specific port instead:</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code><span class="token comment"># Add a port</span>
miasma apps:edit -a example-website --add-target-ports <span class="token number">80</span>

<span class="token comment"># Go back to a randomized port</span>
miasma apps:edit -a postgres --rm-target-ports <span class="token number">80</span>
</code></pre></div><h2 id="expose-multiple-ports" tabindex="-1"><a class="header-anchor" href="#expose-multiple-ports" aria-hidden="true">#</a> Expose Multiple Ports</h2>
<p>If your image uses multiple ports, and you want to be able to access both of them, you can just add multiple target or published ports. Behind the scenes, multiple target ports will be randomly generated for the container, and provided to the app via the <code v-pre>PORT_1</code> (same as <code v-pre>PORT</code>), <code v-pre>PORT_2</code>, <code v-pre>PORT_3</code>, etc environment variables.</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code>miasma apps:edit -a example --add-published-ports <span class="token number">8081</span> --add-published-ports <span class="token number">8082</span>
</code></pre></div><p>To use hard coded target ports, you can also set those:</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code>miasma apps:edit -a example --add-target-ports <span class="token number">8081</span> --add-target-ports <span class="token number">8082</span>
</code></pre></div><div class="custom-container warning"><p class="custom-container-title">Order Matters</p>
<p>In the example above, 8081 was added before 8082 in both cases, so the ports end up connected like <code v-pre>8081:8081</code> and <code v-pre>8082:8082</code>.</p>
<p>If we had instead ran:</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code>miasma apps:edit -a example --add-published-ports <span class="token number">8082</span> --add-published-ports <span class="token number">8081</span>
</code></pre></div><p>We would end up with <code v-pre>8081:8082</code> and <code v-pre>8082:8081</code> instead.</p>
</div>
</div></template>
