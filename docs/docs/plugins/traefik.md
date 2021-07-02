---
id: traefik
title: Traefik Routing
---

Traefik (pronounced "traffic") is a modern ingress router used to define hostname and path routing. Checkout their documentation to learn more:

<https://traefik.io>

But the good new is, you don't need to know anything more to setup ingress routing for your apps! To get started, install the plugin via the CLI:

```bash
$ miasma plugins:add traefik
Installing traefik...
Done!
```

You can view Traefik's dashboard at `http://<swarm-ip>:4000` in a browser.

## Add Your First Routes

If you're just using Miasma inside a local home network, there are no restrictions on what the hostnames are.

For this walk-through, we're going to assume you've chosen `home.io` to be the primary hostname you want to route these 5 apps accordingly:

- `web-1` &rarr; `home.io`
- `api-1` &rarr; `api1.home.io`
- `web-2` &rarr; `web2.home.io`
- `api-2` &rarr; `web2.home.io/api`

As you can see, we're going to be able to map to the apex domain, subdomains, and even paths.

### Router/DNS Setup

Most Routers have the ability to set custom dns maps. The goal here is to redirect all the domains we plan on using to any of the IP addresses the docker swarm is running on.

Traefik exposes port `80` externally (the default web, non-https port), and it's the port that all requests you want routed should be made to go through. Because of this, the DNS mapping is very straight forward, you can exclude any port numbering. Just use the IP of any node in your cluster

:::note IP Addresses
Because of [Docker Swarm's routing mesh](https://docs.docker.com/engine/swarm/ingress/), the domain names can be mapped to any IP address that is apart of the swarm. It doesn't have to the the IP of the node Traefik is running on
:::

Here are the custom DNS mappings you need to setup for this example:

- `api1.home.io` &rarr; `192.168.1.1`
- `home.io` &rarr; `192.168.1.1`
- `web2.home.io` &rarr; `192.168.1.1`

How you set this up depends on your router. Generally, the router can be configured at <http://192.168.0.1> or <http://0.0.0.0>

### Setting Up the Routes

From here on, it's easy to setup the routes for our 4 apps using the Misama CLI.

```bash
$ miasma traefik:set -a web-1 --host home.io
Done!

$ miasma traefik:set -a api-1 --host api1.home.io
Done!

$ miasma traefik:set -a web-2 --host web2.home.io
Done!

$ miasma traefik:set -a api-2 --host web2.home.io --path /api
Done!
```

And the routes are setup! Give it a minute, and watch the HTTP Routers on the Traefik dashboard (<http://swarm-ip:4000/dashboard/#/http/routers>) to see when they've been registered.

:::info
After adding or updating routes, traefik automatically picks up on them. This process can take up to 2 minutes, so don't restart the app wondering why the route is not working immediately.
:::
