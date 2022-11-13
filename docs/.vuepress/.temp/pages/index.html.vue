<template><div><h1 id="introduction" tabindex="-1"><a class="header-anchor" href="#introduction" aria-hidden="true">#</a> Introduction</h1>
<p>Miasma was designed as a simple way to manage applications running on my cluster of Raspberry Pis in my closet.</p>
<p><img src="https://github.com/aklinker1/miasma/raw/main/.github/assets/ui.png" alt="UI"></p>
<p>It has two components: a <strong>lightweight server</strong>, and a <strong>CLI tool</strong>. The server is responsible for managing everything docker related, while CLI talks with the server via it's GraphQL API.</p>
<h2 id="features" tabindex="-1"><a class="header-anchor" href="#features" aria-hidden="true">#</a> Features</h2>
<ul>
<li>Create, manage, and destroy apps</li>
<li>Custom hostname and path routing</li>
<li>Multiple architecture support (<code v-pre>arm/v7</code>, <code v-pre>arm64</code>, <code v-pre>amd64</code>)</li>
<li>Cluster/multiple device support</li>
<li>Automatically managed HTTPS certificates</li>
</ul>
<h3 id="future-work" tabindex="-1"><a class="header-anchor" href="#future-work" aria-hidden="true">#</a> Future Work</h3>
<ul>
<li>Web dashboard (in progress)</li>
<li>Horizontal scaling</li>
</ul>
<h2 id="not-features" tabindex="-1"><a class="header-anchor" href="#not-features" aria-hidden="true">#</a> Not Features</h2>
<p>Miasma does not, and will never, provide complete docker service configuration. It provides a flexible, but opinionated way of managing applications.</p>
<div class="custom-container warning"><p class="custom-container-title">⚠️&emsp;Miasma was not designed for an enterprise environment</p>
<p>If you're considering using Miasma to manage some kind of production environment in the cloud, try it out locally first. Understand what it can and can't do before deciding if it's right for your stack.</p>
</div>
</div></template>
