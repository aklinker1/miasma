---
id: traefik
title: Traefik Routing
---

Traefik (pronounced "traffic") is a modern ingress router used to define hostname and path routing. Checkout their documentation to learn more:

<https://traefik.io>

To get started, install the plugin via the CLI:

```bash
$ miasma plugins:install traefik
Installing traefik...
Done!
```

The traefik plugin requires additional setup via docker, it's placement constraints don't match any nodes in the swarm. It is looking for a node with the label: `traefik=true`. We need to add this label to a manager node. **It is very important that Traefik is put on the manager node, otherwise routing will not work**.

1. SSH into your manager node
   ```bash
   $ shh user@192.168.1.0
   ```
1. Verify that traefik is not running, you should see `0/1` replicas:
   ```bash
   $ docker service ls
   ID                  NAME                MODE                REPLICAS            IMAGE                                       PORTS
   lyvg0isec8h4        plugin-traefik      replicated          0/1                 traefik:v2.4                                *:80->80/tcp,   *:4000->8080/tcp
   ...
   ```
1. List all the nodes in the swarm
   ```bash
   $ docker node ls
   ID                            HOSTNAME             STATUS   AVAILABILITY   MANAGER STATUS      ENGINE VERSION
   5b3svkys1cs8qlzgsd6aaeosf *   docker1-rpi3b-32gb   Ready    Active         Leader              19.03.13
   u4latjalzvg1dn9ps4niud50b     docker2-rpi4-128gb   Ready    Active                             19.03.12
   dhbbq3fhy2txdtj6jmixmoi08     docker3-rpi4-64gb    Ready    Active                             19.03.12
   ```
1. The management node we're looking for has a `MANAGER STATUS` of "Leader". You may have 1 or multiple leaders, any are ok for the next step
1. Using a "Leader" node ID (in this case `5b3svkys1cs8qlzgsd6aaeosf`), add the label to the node. After the label is added, docker swarm will automatically start the traefik service now that there is a valid node to put it.
   ```bash
   $ docker node update 5b3svkys1cs8qlzgsd6aaeosf --label-add traefik=true
   ```
1. Verify that the service running traefik is now working, you should see replicas `1/1`. It might take a minute or two to pull down the image and start up
   ```bash
   $ docker service ls
   ID             NAME             MODE         REPLICAS   IMAGE          PORTS
   lyvg0isec8h4   plugin-traefik   replicated   1/1        traefik:v2.4   *:80->80/tcp,   *:4000->8080/tcp
   ...
   ```

:::note
Eventually this step will be built into the Miasma CLI, so you won't have to SSH or wonder what node to add the label to.
:::

After adding the label, you can view Traefik's dashboard at `https://<swarm-ip>:4000` in a browser.

## Internal Network Routing

If you're just using Miasma inside a local home network, there are no restrictions on what the hostnames are.

For this walk-through, we're going to assume you've chosen `home.io` to be your primary hostname you want to route these 5 apps accordingly:

- `api-1` &rarr; `api1.home.io`
- `web-1` &rarr; `home.io`
- `api-2` &rarr; `web2.home.io/api`
- `web-2` &rarr; `web2.home.io`

As you can see, we're going to be able to map to the apex domain, subdomains, and even paths.

### Router/DNS Setup

Most Routers have the ability to set custom dns maps. The goal here is to redirect all the domains we plan on using to any of the IP addresses the docker swarm is running on.

Traefik exposes port `80` externally (the default web, non-https port), and it's the port that all requests you want routed should be made to go through. Because of this, the DNS mapping is very straight forward, you can exclude any port numbering.

Because of [Docker Swarm's routing mesh](https://docs.docker.com/engine/swarm/ingress/), the domain names can be mapped to any IP address that is apart of the swarm. In this case `192.168.1.1` is one of the IP addresses. It doesn't have to the the IP on the manager node that got the `traefik` label.

- `api1.home.io` &rarr; `192.168.1.1`
- `home.io` &rarr; `192.168.1.1`
- `web2.home.io` &rarr; `192.168.1.1`

How you set this up depends on your router. Generally, the router can be configured at <http://192.168.0.1> or <http://0.0.0.0>.

### Setting up the routes

From here on, it's easy to setup routing for our 4 apps using the Misama CLI.

```bash
$ miasma router:set -a api-1 --hostname api1.home.io
Update routing for api-1...
Done!

$ miasma router:set -a web-1 --hostname home.io
Update routing for web-1...
Done!

$ miasma router:set -a web-2 --hostname web2.home.io
Update routing for web-2...
Done!

$ miasma router:set -a api-2 --hostname web2.home.io --path /api
Update routing for api-2...
Done!
```

And our routes are setup!

:::info
After adding or updating routes, traefik automatically picks up on them. This process can take up to 2 minutes, so don't restart the app wondering why the route is not working immediately.

Give it a minute, and watch the HTTP Routers on the Traefik dashboard (<http://swarm-ip:4000/dashboard/#/http/routers>) to see when it's been registered.
:::
