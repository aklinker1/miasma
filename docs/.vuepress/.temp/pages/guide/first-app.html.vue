<template><div><h1 id="create-your-first-app" tabindex="-1"><a class="header-anchor" href="#create-your-first-app" aria-hidden="true">#</a> Create Your First App</h1>
<p>If you haven't already, follow the install instructions to spin up a Miasma Server and install the CLI.</p>
<p>We'll assume you're running the server on you dev machine, and access everything via <code v-pre>localhost</code>.</p>
<nav class="table-of-contents"><ul><li><RouterLink to="#connect-to-the-miasma-server">Connect to the Miasma Server</RouterLink></li><li><RouterLink to="#creating-an-app">Creating an App</RouterLink></li></ul></nav>
<h2 id="connect-to-the-miasma-server" tabindex="-1"><a class="header-anchor" href="#connect-to-the-miasma-server" aria-hidden="true">#</a> Connect to the Miasma Server</h2>
<p>Before managing any apps, you first need to tell the CLI where your server is at:</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code>$ miasma connect localhost:3000

Connected to miasma<span class="token operator">!</span>
  Miasma Server Version: <span class="token number">1.1</span>.0
  Docker Version: <span class="token number">20.10</span>.17
  Cluster Enabled? <span class="token boolean">true</span>
  Join Command: <span class="token function">docker</span> swarm <span class="token function">join</span> --token <span class="token operator">&lt;</span>some-token<span class="token operator">></span> <span class="token number">192.168</span>.0.3:2377

✔ Connected to server
</code></pre></div><p>Nice! We're connected, so lets create an app.</p>
<h2 id="creating-an-app" tabindex="-1"><a class="header-anchor" href="#creating-an-app" aria-hidden="true">#</a> Creating an App</h2>
<p>Let's deploy an example web app that just responds with information about the incoming request using the <a href="https://hub.docker.com/r/ealen/echo-server" target="_blank" rel="noopener noreferrer"><code v-pre>ealen/echo-server</code><ExternalLinkIcon/></a> image.</p>
<div class="custom-container tip"><p class="custom-container-title">Why <code>ealen/echo-server</code>?</p>
<p>Because it runs on the port specified by the <code v-pre>PORT</code> environment variable. For more details, learn how Miasma <a href="/guide/port-management">manages ports automatically</a>.</p>
</div>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code>$ miasma apps:create example-web-app -i ealen/echo-server

Creating example-web-app<span class="token punctuation">..</span>.
✔ example-web-app started
</code></pre><div class="highlight-lines"><div class="highlight-line">&nbsp;</div><br><br><br></div></div><p>List the apps. We now have one app, and it's deployed to at <code v-pre>http://&lt;server-ip&gt;:3001</code>.</p>
<div class="language-text ext-text"><pre v-pre class="language-text"><code>$ miasma apps

Apps
 ● example-web-app (http://localhost:3001, 1/1 running)

(1 total)
</code></pre></div><div class="custom-container tip"><p class="custom-container-title">TIP</p>
<p>If you see <code v-pre>0/1 running</code>, wait a few seconds and run <code v-pre>misama apps</code> again.</p>
</div>
<p>Visit <a href="http://localhost:3001" target="_blank" rel="noopener noreferrer">http://localhost:3001<ExternalLinkIcon/></a> in your browser, and you should see a JSON response.</p>
<div class="language-json ext-json"><pre v-pre class="language-json"><code><span class="token punctuation">{</span>
    <span class="token property">"http"</span><span class="token operator">:</span> <span class="token punctuation">{</span>
        <span class="token property">"baseUrl"</span><span class="token operator">:</span> <span class="token string">""</span><span class="token punctuation">,</span>
        <span class="token property">"method"</span><span class="token operator">:</span> <span class="token string">"GET"</span><span class="token punctuation">,</span>
        <span class="token property">"originalUrl"</span><span class="token operator">:</span> <span class="token string">"/some/path"</span><span class="token punctuation">,</span>
        <span class="token property">"protocol"</span><span class="token operator">:</span> <span class="token string">"http"</span>
    <span class="token punctuation">}</span><span class="token punctuation">,</span>
    <span class="token comment">// ...</span>
<span class="token punctuation">}</span>
</code></pre></div><p>The app is up and running with a single CLI command!</p>
<hr>
<h4 id="next-steps" tabindex="-1"><a class="header-anchor" href="#next-steps" aria-hidden="true">#</a> Next Steps</h4>
<ul>
<li>Read how Miasma <a href="/guide/port-management">automatically manages ports</a></li>
<li>Learn how to <a href="/guide/app-communication">communicate between apps</a></li>
<li>See more complex, common <a href="/guide/examples">example apps</a></li>
<li>Learn about <a href="/plugins">plugins</a></li>
<li>Checkout the <a href="/reference/cli">CLI reference</a> for advanced usage</li>
<li>Enable <a href="/authorization">authorization</a></li>
</ul>
</div></template>
