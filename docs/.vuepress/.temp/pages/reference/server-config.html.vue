<template><div><h1 id="server-configuration" tabindex="-1"><a class="header-anchor" href="#server-configuration" aria-hidden="true">#</a> Server Configuration</h1>
<p>The server can be configured through environment variables. The following is a list of all the environment variables and what they do.</p>
<nav class="table-of-contents"><ul><li><RouterLink to="#access-token">ACCESS_TOKEN</RouterLink></li><li><RouterLink to="#auto-upgrade-cron">AUTO_UPGRADE_CRON</RouterLink></li><li><RouterLink to="#port">PORT</RouterLink></li></ul></nav>
<div class="custom-container tip"><p class="custom-container-title">TIP</p>
<p>All environment variables are optional, their default values are listed</p>
</div>
<h2 id="access-token" tabindex="-1"><a class="header-anchor" href="#access-token" aria-hidden="true">#</a> <code v-pre>ACCESS_TOKEN</code></h2>
<p><strong>Default: <code v-pre>&quot;&quot;</code></strong></p>
<p>When set to something other than an empty string, it enables server wide authorization. See the <a href="/authorization">authorization docs</a> for more details</p>
<h2 id="auto-upgrade-cron" tabindex="-1"><a class="header-anchor" href="#auto-upgrade-cron" aria-hidden="true">#</a> <code v-pre>AUTO_UPGRADE_CRON</code></h2>
<p><strong>Default: <code v-pre>&quot;@daily&quot;</code></strong></p>
<p>According to this CRON expression, the latest version of app images are pulled. If a newer version of the image exists, the app is upgraded to use that new image.</p>
<p>Internally, Miasma uses <a href="https://pkg.go.dev/github.com/robfig/cron" target="_blank" rel="noopener noreferrer"><code v-pre>robfig/cron/v3</code><ExternalLinkIcon/></a> to evaluate the CRON expression. Any expression supported by this library can be used for <code v-pre>AUTO_UPGRADE_CRON</code>.</p>
<p>By default, all apps automatically upgrade on this schedule. To prevent an app from upgrading automatically, you can set <code v-pre>autoUpgrade</code> to <code v-pre>false</code> on that app.</p>
<h2 id="port" tabindex="-1"><a class="header-anchor" href="#port" aria-hidden="true">#</a> <code v-pre>PORT</code></h2>
<p><strong>Default: <code v-pre>3000</code></strong></p>
<p>The port the Miasma server runs on. It can be any valid port, but it is recommended to stay away from 3001-4000, where apps are published.</p>
</div></template>
