<template><div><h1 id="graphql-api" tabindex="-1"><a class="header-anchor" href="#graphql-api" aria-hidden="true">#</a> GraphQL API</h1>
<p>The Miasma server ships with a GraphQL API for accessing and editing apps. Once your server has been started, you can access it at:</p>
<p><code v-pre>http://&lt;server-ip&gt;:3000/graphql</code></p>
<nav class="table-of-contents"><ul><li><RouterLink to="#playground">Playground</RouterLink></li><li><RouterLink to="#introspection">Introspection</RouterLink></li><li><RouterLink to="#changesets">Changesets</RouterLink></li><li><RouterLink to="#schema">Schema</RouterLink><ul><li><RouterLink to="#query">Query</RouterLink></li><li><RouterLink to="#mutation">Mutation</RouterLink></li><li><RouterLink to="#objects">Objects</RouterLink></li><li><RouterLink to="#inputs">Inputs</RouterLink></li><li><RouterLink to="#enums">Enums</RouterLink></li><li><RouterLink to="#scalars">Scalars</RouterLink></li></ul></li></ul></nav>
<h2 id="playground" tabindex="-1"><a class="header-anchor" href="#playground" aria-hidden="true">#</a> Playground</h2>
<p>The API includes a GraphiQL playground that lets you experiment with the API from the browser. No additional tools necessary! You can access it at:</p>
<p><code v-pre>http://&lt;server-ip&gt;:3000/playground</code></p>
<h2 id="introspection" tabindex="-1"><a class="header-anchor" href="#introspection" aria-hidden="true">#</a> Introspection</h2>
<p>The API also supports introspection so you can get docs inside a HTTP client like Insomnia or Postman.</p>
<h2 id="changesets" tabindex="-1"><a class="header-anchor" href="#changesets" aria-hidden="true">#</a> Changesets</h2>
<p>To perform partial updates of objects, the GraphQL API contains some input types with the <code v-pre>Changes</code> suffix, like <code v-pre>AppChanges</code>.</p>
<p>When using these types, the API will only update the fields provided in the JSON. That means excluding a field will not set it to <code v-pre>null</code>. To set a field to <code v-pre>null</code>, you would have to include the field as <code v-pre>null</code> for the API to make that change.</p>
<p>For example, if we want to change the app name, remove the group, and leave everything else as is:</p>
<div class="language-graphql ext-graphql"><pre v-pre class="language-graphql"><code><span class="token comment"># Query</span>
<span class="token keyword">mutation</span> <span class="token definition-mutation function">updateApp</span><span class="token punctuation">(</span><span class="token variable variable-input">$id</span><span class="token punctuation">:</span> <span class="token scalar">ID</span><span class="token operator">!</span><span class="token punctuation">,</span> <span class="token variable variable-input">$changes</span><span class="token punctuation">:</span> <span class="token class-name">AppChanges</span><span class="token operator">!</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
    <span class="token property-query property-mutation">editApp</span><span class="token punctuation">(</span><span class="token attr-name">id</span><span class="token punctuation">:</span> <span class="token variable variable-input">$id</span><span class="token punctuation">,</span> <span class="token attr-name">changes</span><span class="token punctuation">:</span> <span class="token variable variable-input">$changes</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
        <span class="token operator">...</span>
    <span class="token punctuation">}</span>
<span class="token punctuation">}</span>
</code></pre></div><div class="language-json ext-json"><pre v-pre class="language-json"><code><span class="token comment">// Variables</span>
<span class="token punctuation">{</span>
    <span class="token property">"id"</span><span class="token operator">:</span> <span class="token string">"..."</span><span class="token punctuation">,</span>
    <span class="token property">"changes"</span><span class="token operator">:</span> <span class="token punctuation">{</span>
        <span class="token property">"name"</span><span class="token operator">:</span> <span class="token string">"New Name"</span><span class="token punctuation">,</span>
        <span class="token property">"group"</span><span class="token operator">:</span> <span class="token null keyword">null</span><span class="token punctuation">,</span>
    <span class="token punctuation">}</span>
<span class="token punctuation">}</span>
</code></pre></div><h2 id="schema" tabindex="-1"><a class="header-anchor" href="#schema" aria-hidden="true">#</a> Schema</h2>
<h3 id="query" tabindex="-1"><a class="header-anchor" href="#query" aria-hidden="true">#</a> Query</h3>
<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>health</strong></td>
<td valign="top"><a href="#health">Health</a></td>
<td>
Get the server's health and version information
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>listApps</strong></td>
<td valign="top">[<a href="#app">App</a>!]!</td>
<td>
List the running apps
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">page</td>
<td valign="top"><a href="#int">Int</a></td>
<td>
The page to start at for pagination, the first page is 1.
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">size</td>
<td valign="top"><a href="#int">Int</a></td>
<td>
Number of apps to return per page.
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">showHidden</td>
<td valign="top"><a href="#boolean">Boolean</a></td>
<td>
Whether or not to includes apps that are marked as hidden.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>getApp</strong></td>
<td valign="top"><a href="#app">App</a>!</td>
<td>
Grab an app by it's ID
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">id</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>getAppTasks</strong></td>
<td valign="top">[<a href="#apptask">AppTask</a>!]!</td>
<td>
List of tasks for an app, up the 5 most recent
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">id</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>listPlugins</strong></td>
<td valign="top">[<a href="#plugin">Plugin</a>!]!</td>
<td>
List all the available plugins for Miasma
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>getPlugin</strong></td>
<td valign="top"><a href="#plugin">Plugin</a>!</td>
<td>
Grab a plugin by it's name
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">name</td>
<td valign="top"><a href="#pluginname">PluginName</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>nodes</strong></td>
<td valign="top">[<a href="#node">Node</a>!]!</td>
<td>
List the nodes that are apart of the cluster
</td>
</tr>
</tbody>
</table>
<h3 id="mutation" tabindex="-1"><a class="header-anchor" href="#mutation" aria-hidden="true">#</a> Mutation</h3>
<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>createApp</strong></td>
<td valign="top"><a href="#app">App</a>!</td>
<td>
Create and start a new app.
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">input</td>
<td valign="top"><a href="#appinput">AppInput</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>editApp</strong></td>
<td valign="top"><a href="#app">App</a>!</td>
<td>
Edit app configuration.
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">id</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">changes</td>
<td valign="top"><a href="#appchanges">AppChanges</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>deleteApp</strong></td>
<td valign="top"><a href="#app">App</a>!</td>
<td>
Stop and delete an app.
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">id</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>startApp</strong></td>
<td valign="top"><a href="#app">App</a>!</td>
<td>
Start a stopped app.
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">id</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>stopApp</strong></td>
<td valign="top"><a href="#app">App</a>!</td>
<td>
Stop a running app.
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">id</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>restartApp</strong></td>
<td valign="top"><a href="#app">App</a>!</td>
<td>
Stop and restart an app.
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">id</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>upgradeApp</strong></td>
<td valign="top"><a href="#app">App</a>!</td>
<td>
Pull the latest version of the app's image and then restart.
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">id</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>enablePlugin</strong></td>
<td valign="top"><a href="#plugin">Plugin</a>!</td>
<td>
Enable one of Miasma's plugins.
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">name</td>
<td valign="top"><a href="#pluginname">PluginName</a>!</td>
<td>
The name of the plugin to enable.
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">config</td>
<td valign="top"><a href="#map">Map</a></td>
<td>
Any plugin specific configuration.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>disablePlugin</strong></td>
<td valign="top"><a href="#plugin">Plugin</a>!</td>
<td>
Disable one of Miasma's plugins.
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">name</td>
<td valign="top"><a href="#pluginname">PluginName</a>!</td>
<td>
The name of the plugin to disable.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>setAppEnv</strong></td>
<td valign="top"><a href="#map">Map</a></td>
<td>
Set an app's environnement variables
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">appId</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">newEnv</td>
<td valign="top"><a href="#map">Map</a></td>
<td>
A map of variable names to their values. Docker only supports UPPER_SNAKE_CASE variable names
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>setAppRoute</strong></td>
<td valign="top"><a href="#route">Route</a></td>
<td>
Set's an app's route.
<p>Only available when the 'router' plugin is enabled</p>
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">appId</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">route</td>
<td valign="top"><a href="#routeinput">RouteInput</a></td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>removeAppRoute</strong></td>
<td valign="top"><a href="#route">Route</a></td>
<td>
Removes an app's route.
<p>Only available when the 'router' plugin is enabled</p>
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">appId</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
</tbody>
</table>
<h3 id="objects" tabindex="-1"><a class="header-anchor" href="#objects" aria-hidden="true">#</a> Objects</h3>
<h4 id="app" tabindex="-1"><a class="header-anchor" href="#app" aria-hidden="true">#</a> App</h4>
<p>Managed application</p>
<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>createdAt</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>updatedAt</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>name</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>
The app name. Different from the docker service name, which is the name but lower case and all spaces replaced with dashes
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>system</strong></td>
<td valign="top"><a href="#boolean">Boolean</a>!</td>
<td>
Whether or not the application is managed by the system. You cannot edit or delete system apps.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>group</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>
A string used to group the app
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>image</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>
The image and tag the application runs.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>imageDigest</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>
The currently running image digest (hash). Used internally when running
applications instead of the tag because the when a new image is pushed, the
tag stays the same but the digest changes.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>autoUpgrade</strong></td>
<td valign="top"><a href="#boolean">Boolean</a>!</td>
<td>
Whether or not the app should automatically upgrade when a newer version of it's image is available. Defaults to `true` when creating an app
<p>App upgrades are automatically checked according the the <code v-pre>AUTO_UPDATE_CRON</code> expression.</p>
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>hidden</strong></td>
<td valign="top"><a href="#boolean">Boolean</a>!</td>
<td>
Whether or not the app is returned during regular requests.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>route</strong></td>
<td valign="top"><a href="#route">Route</a></td>
<td>
If the app has a route and the traefik plugin is enabled, this is it's config.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>simpleRoute</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>
If the app has a route and the traefik plugin is enabled, this is a simple representation of it.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>availableAt</strong></td>
<td valign="top">[<a href="#string">String</a>!]!</td>
<td>
A list of URLs the application can be accessed at, including the `simpleRoute`, and all the published ports
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">clusterIpAddress</td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>env</strong></td>
<td valign="top"><a href="#map">Map</a></td>
<td>
The environment variables configured for this app.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>status</strong></td>
<td valign="top"><a href="#runtimestatus">RuntimeStatus</a>!</td>
<td>
Whether or not the application is running, or stopped.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>instances</strong></td>
<td valign="top"><a href="#appinstances">AppInstances</a>!</td>
<td>
The number of instances running vs what should be running.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>targetPorts</strong></td>
<td valign="top">[<a href="#int">Int</a>!]</td>
<td>
The ports that the app is listening to inside the container. If no target
ports are specified, then the container should respect the `PORT` env var.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>publishedPorts</strong></td>
<td valign="top">[<a href="#int">Int</a>!]</td>
<td>
The ports that you access the app through in the swarm. This field can, and
should be left empty. Miasma automatically manages assigning published ports
between 3001-4999. If you need to specify a port, make sure it's outside that
range or the port has not been taken. Plugins have set ports starting with
4000, so avoid 4000-4020 if you want to add a plugin at a later date.
<p>If these ports are ever cleared, the app will continue using the same ports it
was published to before, so that the ports don't change unnecessarily. If you
removed it to clear a port for another app/plugin, make sure to restart the
app and a new, random port will be allocated for the app, freeing the old
port.</p>
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>placement</strong></td>
<td valign="top">[<a href="#string">String</a>!]</td>
<td>
The placement constraints specifying which nodes the app will be ran on. Any
valid value for the [`--constraint` flag](https://docs.docker.com/engine/swarm/services/#placement-constraints)
is valid item in this list.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>volumes</strong></td>
<td valign="top">[<a href="#boundvolume">BoundVolume</a>!]</td>
<td>
Volume bindings for the app.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>networks</strong></td>
<td valign="top">[<a href="#string">String</a>!]</td>
<td>
A list of other apps that the service communicates with using their service
name and docker's internal DNS. Services don't have to be two way; only the
service that accesses the other needs the other network added.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>command</strong></td>
<td valign="top">[<a href="#string">String</a>!]</td>
<td>
Custom docker command. This is an array of arguments starting with the binary that is being executed
</td>
</tr>
</tbody>
</table>
<h4 id="appinstances" tabindex="-1"><a class="header-anchor" href="#appinstances" aria-hidden="true">#</a> AppInstances</h4>
<p>Contains information about how many instances of the app are running vs supposed to be running</p>
<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>running</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>total</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td></td>
</tr>
</tbody>
</table>
<h4 id="apptask" tabindex="-1"><a class="header-anchor" href="#apptask" aria-hidden="true">#</a> AppTask</h4>
<p>Tasks define the desired state of on app. If you're familiar with docker, this returns the result of <code v-pre>docker service ps</code></p>
<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>message</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>state</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>desiredState</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>timestamp</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>appId</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>app</strong></td>
<td valign="top"><a href="#app">App</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>nodeId</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>node</strong></td>
<td valign="top"><a href="#node">Node</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>name</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>error</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>exitCode</strong></td>
<td valign="top"><a href="#int">Int</a></td>
<td></td>
</tr>
</tbody>
</table>
<h4 id="boundvolume" tabindex="-1"><a class="header-anchor" href="#boundvolume" aria-hidden="true">#</a> BoundVolume</h4>
<p>Docker volume configuration</p>
<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>target</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>
The path inside the container that the data is served from.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>source</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>
The volume name or directory on the host that the data is stored in.
</td>
</tr>
</tbody>
</table>
<h4 id="clusterinfo" tabindex="-1"><a class="header-anchor" href="#clusterinfo" aria-hidden="true">#</a> ClusterInfo</h4>
<p>Contains useful information about the cluster.</p>
<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>
The Docker swarm ID
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>joinCommand</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>
The command to run on other machines to join the cluster
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>createdAt</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td>
When the cluster was initialized
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>updatedAt</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td>
When the cluster was last updated
</td>
</tr>
</tbody>
</table>
<h4 id="health" tabindex="-1"><a class="header-anchor" href="#health" aria-hidden="true">#</a> Health</h4>
<p>Server health and version information</p>
<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>version</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>
Miasma server's current version.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>dockerVersion</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>
The version of docker running on the host, or null if docker is not running.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>cluster</strong></td>
<td valign="top"><a href="#clusterinfo">ClusterInfo</a></td>
<td>
The cluster versioning and information, or `null` if not apart of a cluster.
</td>
</tr>
</tbody>
</table>
<h4 id="log" tabindex="-1"><a class="header-anchor" href="#log" aria-hidden="true">#</a> Log</h4>
<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>message</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>timestamp</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td></td>
</tr>
</tbody>
</table>
<h4 id="node" tabindex="-1"><a class="header-anchor" href="#node" aria-hidden="true">#</a> Node</h4>
<p>Details about a machine in the cluster.</p>
<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>
The docker node's ID.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>os</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>
The OS the node is running
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>architecture</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>
The CPU architecture of the node. Services are automatically placed on nodes based on their image's supported architectures and the nodes' architectures.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>hostname</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>
The machines hostname, as returned by the `hostname` command.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>ip</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>
The IP address the node joined the cluster as.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>status</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>
`unknown`, `down`, `ready`, or `disconnected`. See Docker's [API docs](https://docs.docker.com/engine/api/v1.41/#operation/NodeInspect).
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>statusMessage</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>
The node's status message, usually present when when the status is not `ready`.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>labels</strong></td>
<td valign="top"><a href="#map">Map</a>!</td>
<td>
The node's labels, mostly used to place apps on specific nodes.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>services</strong></td>
<td valign="top">[<a href="#app">App</a>!]!</td>
<td>
List of apps running on the machine
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">showHidden</td>
<td valign="top"><a href="#boolean">Boolean</a></td>
<td>
Same as `listApps`'s argument. When `true`, hidden apps will be returned
</td>
</tr>
</tbody>
</table>
<h4 id="plugin" tabindex="-1"><a class="header-anchor" href="#plugin" aria-hidden="true">#</a> Plugin</h4>
<p>Plugins are apps with deeper integrations with Miasma.</p>
<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>name</strong></td>
<td valign="top"><a href="#pluginname">PluginName</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>enabled</strong></td>
<td valign="top"><a href="#boolean">Boolean</a>!</td>
<td>
Whether or not the plugin has been enabled.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>config</strong></td>
<td valign="top"><a href="#map">Map</a>!</td>
<td>
Plugin's configuration.
</td>
</tr>
</tbody>
</table>
<h4 id="route" tabindex="-1"><a class="header-anchor" href="#route" aria-hidden="true">#</a> Route</h4>
<p>Rules around where an app can be accessed from.</p>
<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>appId</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>createdAt</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>updatedAt</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>host</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>
The URL's hostname, ex: 'example.com' or 'google.com'.
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>path</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>
A custom path at the end of the URL: ex: '/search' or '/console'
</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>traefikRule</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>
A custom Traefik rule instead of just a host and path, ex: '(Host(domain1.com) || Host(domain2.com)'
<p>See <a href="https://doc.traefik.io/traefik/routing/routers/#rule" target="_blank" rel="noopener noreferrer">Traefik's docs<ExternalLinkIcon/></a> for usage and complex examples.</p>
</td>
</tr>
</tbody>
</table>
<h4 id="subscription" tabindex="-1"><a class="header-anchor" href="#subscription" aria-hidden="true">#</a> Subscription</h4>
<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>appLogs</strong></td>
<td valign="top"><a href="#log">Log</a>!</td>
<td>
Returns the latest log one at a time.
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">id</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>
The ID of the app you want to listen for logs from.
</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">excludeStdout</td>
<td valign="top"><a href="#boolean">Boolean</a></td>
<td></td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">excludeStderr</td>
<td valign="top"><a href="#boolean">Boolean</a></td>
<td></td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">initialCount</td>
<td valign="top"><a href="#int">Int</a></td>
<td>
The subscription will load this number of logs from the past initially before listening for future logs.
</td>
</tr>
</tbody>
</table>
<h3 id="inputs" tabindex="-1"><a class="header-anchor" href="#inputs" aria-hidden="true">#</a> Inputs</h3>
<h4 id="appchanges" tabindex="-1"><a class="header-anchor" href="#appchanges" aria-hidden="true">#</a> AppChanges</h4>
<p><a href="#Changeset">Changeset</a> input type for <a href="#app">App</a>.</p>
<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>name</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>image</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>group</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>hidden</strong></td>
<td valign="top"><a href="#boolean">Boolean</a></td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>targetPorts</strong></td>
<td valign="top">[<a href="#int">Int</a>!]</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>publishedPorts</strong></td>
<td valign="top">[<a href="#int">Int</a>!]</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>placement</strong></td>
<td valign="top">[<a href="#string">String</a>!]</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>volumes</strong></td>
<td valign="top">[<a href="#boundvolumeinput">BoundVolumeInput</a>!]</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>networks</strong></td>
<td valign="top">[<a href="#string">String</a>!]</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>command</strong></td>
<td valign="top">[<a href="#string">String</a>!]</td>
<td></td>
</tr>
</tbody>
</table>
<h4 id="appinput" tabindex="-1"><a class="header-anchor" href="#appinput" aria-hidden="true">#</a> AppInput</h4>
<p>Input type for <a href="#app">App</a>.</p>
<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>name</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>image</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>autoUpgrade</strong></td>
<td valign="top"><a href="#boolean">Boolean</a></td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>group</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>hidden</strong></td>
<td valign="top"><a href="#boolean">Boolean</a></td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>targetPorts</strong></td>
<td valign="top">[<a href="#int">Int</a>!]</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>publishedPorts</strong></td>
<td valign="top">[<a href="#int">Int</a>!]</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>placement</strong></td>
<td valign="top">[<a href="#string">String</a>!]</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>volumes</strong></td>
<td valign="top">[<a href="#boundvolumeinput">BoundVolumeInput</a>!]</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>networks</strong></td>
<td valign="top">[<a href="#string">String</a>!]</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>command</strong></td>
<td valign="top">[<a href="#string">String</a>!]</td>
<td></td>
</tr>
</tbody>
</table>
<h4 id="boundvolumeinput" tabindex="-1"><a class="header-anchor" href="#boundvolumeinput" aria-hidden="true">#</a> BoundVolumeInput</h4>
<p>Input type for <a href="#boundvolume">BoundVolume</a>.</p>
<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>target</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>source</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
</tbody>
</table>
<h4 id="routeinput" tabindex="-1"><a class="header-anchor" href="#routeinput" aria-hidden="true">#</a> RouteInput</h4>
<p>Input type for <a href="#route">Route</a>.</p>
<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>host</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>path</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>traefikRule</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td></td>
</tr>
</tbody>
</table>
<h3 id="enums" tabindex="-1"><a class="header-anchor" href="#enums" aria-hidden="true">#</a> Enums</h3>
<h4 id="pluginname" tabindex="-1"><a class="header-anchor" href="#pluginname" aria-hidden="true">#</a> PluginName</h4>
<p>Unique identifier for plugins</p>
<table>
<thead>
<th align="left">Value</th>
<th align="left">Description</th>
</thead>
<tbody>
<tr>
<td valign="top"><strong>TRAEFIK</strong></td>
<td>
The name of the [Traefik](https://doc.traefik.io/traefik/) ingress router plugin
</td>
</tr>
</tbody>
</table>
<h4 id="runtimestatus" tabindex="-1"><a class="header-anchor" href="#runtimestatus" aria-hidden="true">#</a> RuntimeStatus</h4>
<table>
<thead>
<th align="left">Value</th>
<th align="left">Description</th>
</thead>
<tbody>
<tr>
<td valign="top"><strong>RUNNING</strong></td>
<td></td>
</tr>
<tr>
<td valign="top"><strong>STOPPED</strong></td>
<td></td>
</tr>
</tbody>
</table>
<h3 id="scalars" tabindex="-1"><a class="header-anchor" href="#scalars" aria-hidden="true">#</a> Scalars</h3>
<h4 id="boolean" tabindex="-1"><a class="header-anchor" href="#boolean" aria-hidden="true">#</a> Boolean</h4>
<p>The <code v-pre>Boolean</code> scalar type represents <code v-pre>true</code> or <code v-pre>false</code>.</p>
<h4 id="id" tabindex="-1"><a class="header-anchor" href="#id" aria-hidden="true">#</a> ID</h4>
<p>The <code v-pre>ID</code> scalar type represents a unique identifier, often used to refetch an object or as key for a cache. The ID type appears in a JSON response as a String; however, it is not intended to be human-readable. When expected as an input type, any string (such as <code v-pre>&quot;4&quot;</code>) or integer (such as <code v-pre>4</code>) input value will be accepted as an ID.</p>
<h4 id="int" tabindex="-1"><a class="header-anchor" href="#int" aria-hidden="true">#</a> Int</h4>
<p>The <code v-pre>Int</code> scalar type represents non-fractional signed whole numeric values. Int can represent values between -(2^31) and 2^31 - 1.</p>
<h4 id="map" tabindex="-1"><a class="header-anchor" href="#map" aria-hidden="true">#</a> Map</h4>
<p>A JSON map of key-value pairs. Values can be any type.</p>
<h4 id="string" tabindex="-1"><a class="header-anchor" href="#string" aria-hidden="true">#</a> String</h4>
<p>The <code v-pre>String</code> scalar type represents textual data, represented as UTF-8 character sequences. The String type is most often used by GraphQL to represent free-form human-readable text.</p>
<h4 id="time" tabindex="-1"><a class="header-anchor" href="#time" aria-hidden="true">#</a> Time</h4>
<p>ISO 8601 date time string.</p>
</div></template>
