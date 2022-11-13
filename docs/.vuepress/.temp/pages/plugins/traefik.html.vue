<template><div><h1 id="traefik-plugin" tabindex="-1"><a class="header-anchor" href="#traefik-plugin" aria-hidden="true">#</a> Traefik Plugin</h1>
<p><a href="https://traefik.io" target="_blank" rel="noopener noreferrer">Traefik<ExternalLinkIcon/></a> (pronounced &quot;traffic&quot;) is a modern ingress router used to define hostname and path routing. To get started, add the plugin via the CLI:</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code>$ miasma plugins:enable TRAEFIK

Enabling TRAEFIK<span class="token punctuation">..</span>.
Done<span class="token operator">!</span>
</code></pre><div class="highlight-lines"><div class="highlight-line">&nbsp;</div><br><br><br></div></div><p>You can view Traefik's dashboard at <code v-pre>http://&lt;server-ip&gt;:8080/dashboard/</code> in a browser.</p>
<div class="custom-container tip"><p class="custom-container-title">TIP</p>
<p>The Traefik plugin is configured run on the Manager Node (placement of <code v-pre>node.role == manager</code>).</p>
</div>
<h2 id="example-usage" tabindex="-1"><a class="header-anchor" href="#example-usage" aria-hidden="true">#</a> Example Usage</h2>
<p>Let's say you want to host 4 apps on the <code v-pre>home.io</code> domain in the following locations:</p>
<table>
<thead>
<tr>
<th>App Name</th>
<th>Hosted At</th>
</tr>
</thead>
<tbody>
<tr>
<td><code v-pre>web-1</code></td>
<td><a href="http://home.io" target="_blank" rel="noopener noreferrer">http://home.io<ExternalLinkIcon/></a></td>
</tr>
<tr>
<td><code v-pre>api-1</code></td>
<td><a href="http://api1.home.io" target="_blank" rel="noopener noreferrer">http://api1.home.io<ExternalLinkIcon/></a></td>
</tr>
<tr>
<td><code v-pre>web-2</code></td>
<td><a href="http://web2.home.io" target="_blank" rel="noopener noreferrer">http://web2.home.io<ExternalLinkIcon/></a></td>
</tr>
<tr>
<td><code v-pre>api-2</code></td>
<td><a href="http://web2.home.io/api" target="_blank" rel="noopener noreferrer">http://web2.home.io/api<ExternalLinkIcon/></a></td>
</tr>
</tbody>
</table>
<p>As you can see, we're going to be able to map to the apex domain, subdomains, and even paths.</p>
<h3 id="dns-records" tabindex="-1"><a class="header-anchor" href="#dns-records" aria-hidden="true">#</a> DNS Records</h3>
<p>You'll need to setup some A records to get started.</p>
<ul>
<li><code v-pre>api1.home.io</code> → <code v-pre>192.168.1.0</code></li>
<li><code v-pre>home.io</code> → <code v-pre>192.168.1.0</code></li>
<li><code v-pre>web2.home.io</code> → <code v-pre>192.168.1.0</code></li>
</ul>
<div class="custom-container tip"><p class="custom-container-title">TIP</p>
<p>For local networks, some routers include the ability to map hostnames to local IP address.</p>
</div>
<h3 id="setting-up-the-routes" tabindex="-1"><a class="header-anchor" href="#setting-up-the-routes" aria-hidden="true">#</a> Setting Up the Routes</h3>
<p>From here on, it's easy to setup the routes for our 4 apps using the Misama CLI.</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code>miasma traefik:set -a web-1 --host home.io
miasma traefik:set -a api-1 --host api1.home.io
miasma traefik:set -a web-2 --host web2.home.io
miasma traefik:set -a api-2 --host web2.home.io --path /api
</code></pre></div><p>And the routes are setup! Give it a minute, and watch the HTTP Routers on the Traefik dashboard (<code v-pre>http://&lt;server-ip&gt;:8080/dashboard/#/http/routers</code>) to see when they've been registered.</p>
<div class="custom-container tip"><p class="custom-container-title">TIP</p>
<p>After adding or updating routes, Traefik will automatically discover them. This process can take up to 2 minutes, so don't restart the apps wondering why their routes aren't working.</p>
</div>
<h2 id="https-tls-support" tabindex="-1"><a class="header-anchor" href="#https-tls-support" aria-hidden="true">#</a> HTTPS &amp; TLS Support</h2>
<p>If you're using Miasma with public IP addresses that can be accessed from the internet like any other web app, you'll want to make sure your apps use HTTPS.</p>
<p>The Traefik plugin can be configured to use HTTPS. Behind the scenes, Traefik is managing and auto-renewing certificates via <a href="https://letsencrypt.org/" target="_blank" rel="noopener noreferrer">LetsEncrypt<ExternalLinkIcon/></a>.</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code><span class="token comment"># Make the directory where you're certs will be stored</span>
<span class="token function">mkdir</span> /root/letsencrypt

<span class="token comment"># Enable the plugin with required config</span>
miasma plugins:enable TRAEFIK --plugin-config <span class="token string">'{
    "enableHttps": true,
    "certsEmail": "&lt;your-email>",
    "certsDir": "/root/letsencrypt"
}'</span>
</code></pre></div><div class="custom-container tip"><p class="custom-container-title">TIP</p>
<p>If you've already enabled the plugin, disable it first and re-enable it with the additional config.</p>
</div>
<p>All three fields are required to enable HTTPS.</p>
<ul>
<li><code v-pre>enableHttps</code> - Self explanitory. Set to <code v-pre>true</code> to enable HTTPS</li>
<li><code v-pre>certsEmail</code> (required when <code v-pre>enableHttps=true</code>) - The email you'd like to use for the LetsEncrypt certificates</li>
<li><code v-pre>certsDir</code> (required when <code v-pre>enableHttps=true</code>) - The path on your manager node to where you want to store certificates. It must exist and be an absolute path</li>
</ul>
<p>You don't have to specify any domains. Certs are generated automatically for all domains configured on running apps.</p>
<p>And that's it!</p>
<div class="custom-container warning"><p class="custom-container-title">WARNING</p>
<p>Self managed certs are not supported at this time. Feel free to open a PR!</p>
</div>
</div></template>
