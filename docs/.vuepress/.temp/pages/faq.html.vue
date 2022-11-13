<template><div><h1 id="frequently-asked-questions" tabindex="-1"><a class="header-anchor" href="#frequently-asked-questions" aria-hidden="true">#</a> Frequently Asked Questions</h1>
<nav class="table-of-contents"><ul><li><RouterLink to="#why-not-kubernetes">Why not Kubernetes?</RouterLink></li><li><RouterLink to="#why-not-just-use-docker-swarm-directly">Why not just use Docker Swarm directly?</RouterLink></li><li><RouterLink to="#what-about-portainer">What about Portainer?</RouterLink></li></ul></nav>
<h3 id="why-not-kubernetes" tabindex="-1"><a class="header-anchor" href="#why-not-kubernetes" aria-hidden="true">#</a> Why not Kubernetes?</h3>
<p>I don't need Kubernetes in my closest. I made Miasma as a little project to host test and IoT applications for my home. Nothing extremely complex, so solutions like Kubernetes are just over the top.</p>
<p>In the end, <strong>Miasma is just crazy simple</strong> and I learned a lot about Docker while working on it. So it was a win-win.</p>
<h3 id="why-not-just-use-docker-swarm-directly" tabindex="-1"><a class="header-anchor" href="#why-not-just-use-docker-swarm-directly" aria-hidden="true">#</a> Why not just use Docker Swarm directly?</h3>
<p>I don't want to SSH into my device every time I want to update the app version or change an environment variable. Ideally, apps should just update auto-magically (watchtower plugin incoming).</p>
<p>I also was really bad at managing traefik routing labels by hand, I always seemed to misspell something or mess something else up, meaning my routing was very flaky. Now, Miasma manages routes consistently and more easily then when I did it by hand.</p>
<p>Plus there was inconsistency in more than just routing, <code v-pre>.env</code> file formatting, <code v-pre>docker-compose.yml</code> format errors, etc. <strong>Lots of things were inconsistent and now aren't thanks to Miasma</strong>.</p>
<h3 id="what-about-portainer" tabindex="-1"><a class="header-anchor" href="#what-about-portainer" aria-hidden="true">#</a> What about Portainer?</h3>
<p>Portainer gives you access to every thing Docker can do on a UI, but I wanted to simplify things, most apps don't need all that configuration.</p>
<p>Having Miasma manage ports and routing for me is huge, it's so easy to mistype a traefik label and be stuck for hours.</p>
</div></template>
