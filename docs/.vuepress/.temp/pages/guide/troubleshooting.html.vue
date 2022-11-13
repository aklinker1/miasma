<template><div><h1 id="troubleshooting" tabindex="-1"><a class="header-anchor" href="#troubleshooting" aria-hidden="true">#</a> Troubleshooting</h1>
<h2 id="app-not-starting" tabindex="-1"><a class="header-anchor" href="#app-not-starting" aria-hidden="true">#</a> App Not Starting</h2>
<p>There are several reasons why your app may not be starting:</p>
<h3 id="_1-there-is-not-a-suitable-machine-in-your-cluster-that-can-run-the-application-s-docker-image" tabindex="-1"><a class="header-anchor" href="#_1-there-is-not-a-suitable-machine-in-your-cluster-that-can-run-the-application-s-docker-image" aria-hidden="true">#</a> 1. There is not a suitable machine in your cluster that can run the application's docker image</h3>
<p>Docker images are built for specific CPU architectures. For example, an image might only support <code v-pre>x86_64</code>, but all your machines in the cluster are Raspberry Pis, which are all <code v-pre>armv7</code>. In this case, your app could not start because it cannot find a suitable node to run on, because none have the <code v-pre>x86_64</code> architecture.</p>
<p>To see if that's the case, you can look at the &quot;tasks&quot; for your app from the web dashboard. If so, you will see an error like this:</p>
<blockquote>
<p>1/1: no suitable node (unsupported platform on X nodes)</p>
</blockquote>
<h3 id="_2-app-is-misconfigured" tabindex="-1"><a class="header-anchor" href="#_2-app-is-misconfigured" aria-hidden="true">#</a> 2. App is misconfigured</h3>
<p>Your app could be configured in a way that is preventing the app from starting. An example of this would be binding volumes that don't exist.</p>
<p>Once again, to see if this is the problem, look at the app's &quot;tasks&quot; on the web dashboard. If this is the case, an error will clearly be shown stating what the problem is.</p>
<h3 id="_3-the-app-is-crashing-on-startup" tabindex="-1"><a class="header-anchor" href="#_3-the-app-is-crashing-on-startup" aria-hidden="true">#</a> 3. The app is crashing on startup</h3>
<p>This could be because you haven't provided required environment variables, or because there's a bug. In either case, you can look at the app's logs to figure out why it's crashing and fix it.</p>
<p>Viewing logs inside Miasma is coming soon. For now, you can use the Docker CLI to see them:</p>
<div class="language-bash ext-sh"><pre v-pre class="language-bash"><code><span class="token function">ssh</span> username@<span class="token operator">&lt;</span>manger-node-ip<span class="token operator">></span>
<span class="token function">docker</span> <span class="token function">service</span> <span class="token function">ls</span> <span class="token comment"># To get the service name/id</span>
<span class="token function">docker</span> <span class="token function">service</span> logs <span class="token operator">&lt;</span>service-name-or-id<span class="token operator">></span> --follow
</code></pre></div><h3 id="_4-ports-are-misconfigured" tabindex="-1"><a class="header-anchor" href="#_4-ports-are-misconfigured" aria-hidden="true">#</a> 4. Ports are misconfigured</h3>
<p>If the app is running, but you can't access/connect to it, the ports are likely misconfigured, or the <code v-pre>$PORT</code>/<code v-pre>$PORT_1</code>, <code v-pre>$PORT_2</code>, <code v-pre>$PORT_3</code>, etc environment variables are not respected.</p>
<p>See the <RouterLink to="/guide/port-management.html">port management docs</RouterLink> for more details on configuring the ports used by your application.</p>
</div></template>
