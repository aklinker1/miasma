---
title: Traefik Ingress Router
---

## Enabling

[Traefik](https://traefik.io) (pronounced "traffic") is a modern ingress router used to define hostname and path routing. To get started, add the plugin via the CLI:

```bash:no-line-numbers{1}
$ miasma plugins:enable TRAEFIK
Enabling TRAEFIK...
Done!
```

You can view Traefik's dashboard at `http://<server-ip>:8080` in a browser.

:::tip
The Traefik plugin is configured run on the Manager Node (placement of `node.role == manager`).
:::

## Example Usage

Let's say you want to host 4 apps on the `home.io` domain in the following locations:

| App Name | Hosted At                 |
| -------- | ------------------------- |
| `web-1`  | <http://home.io>          |
| `api-1`  | <http://api1.home.io>     |
| `web-2`  | <http://web2.home.io>     |
| `api-2`  | <http://web2.home.io/api> |

As you can see, we're going to be able to map to the apex domain, subdomains, and even paths.

### Configuring Domain Names

You'll need to setup at least one domain name to use the traefik plugin. For this example, it's just `home.io`, but you can add as many domains as you'd like.

If you're hosting Miasma in the cloud, you might already have a domain name. If not, or if you're hosting it inside your local network, you'll need to setup DNS settings to point devices toward your cluster when making requests to the domain name.

- When hosted in the cloud, you'll need to update your DNS settings for your domain
- When hosted in a local network, you can use any domain name you'd like, and configure the custom DNS record on a device (via the OS's DNS settings) or at the network level (by configuring your router)

Just point your domain names to the IP address of the cluster, and you're good to go. The traefik plugin runs on port 80, so you don't need to specify a port.

Here are the custom DNS mappings you need to setup for this example:

- `api1.home.io` &rarr; `192.168.1.0`
- `home.io` &rarr; `192.168.1.0`
- `web2.home.io` &rarr; `192.168.1.0`

### Setting Up the Routes

From here on, it's easy to setup the routes for our 4 apps using the Misama CLI.

```bash:no-line-numbers{1,4,7,10}
$ miasma traefik:set -a web-1 --host home.io
Done!

$ miasma traefik:set -a api-1 --host api1.home.io
Done!

$ miasma traefik:set -a web-2 --host web2.home.io
Done!

$ miasma traefik:set -a api-2 --host web2.home.io --path /api
Done!
```

And the routes are setup! Give it a minute, and watch the HTTP Routers on the Traefik dashboard (<http://swarm-ip:8080/dashboard/#/http/routers>) to see when they've been registered.

:::tip
After adding or updating routes, Traefik will automatically discover them. This process can take up to 2 minutes, so don't restart the app wondering why the route is not working immediately.
:::

## HTTPS & TLS Support

If you're using Miasma with public IP addresses that can be accessed from the internet like any other web app, you'll want to make sure your apps use HTTPS.

The Traefik plugin can be configured to use HTTPS. Behind the scenes, Traefik is managing and auto-renewing certificates via [LetsEncrypt](https://letsencrypt.org/).

```bash:no-line-numbers
# Make the directory where you're certs will be stored
mkdir /root/letsencrypt

# Enable the plugin with required config
miasma plugins:enable TRAEFIK --plugin-config '{
    "enableHttps": true,
    "certsEmail": "<your-email>",
    "certsDir": "/root/letsencrypt"
}'
```

:::tip
If you've already enabled the plugin, disable it first and re-enable it with the additional config.
:::

All three fields are required to enable HTTPS.

- `enableHttps` - Self explanitory. Set to `true` to enable HTTPS
- `certsEmail` (required when `enableHttps=true`) - The email you'd like to use for the LetsEncrypt certificates
- `certsDir` (required when `enableHttps=true`) - The path on your manager node to where you want to store certificates. It must exist and be an absolute path

You don't have to specify any domains. Certs are generated automatically for all domains configured on running apps.

And that's it!

:::warning
Self managed certs are not supported at this time. Feel free to open a PR!
:::
