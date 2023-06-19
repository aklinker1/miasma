# Plugins

Plugins are simple docker services with a deep integration with the UI.

[[toc]]

## Traefik

:::warning
The Traefik plugin is not finished yet. These are just the plans and they are subject to change.
:::

The [Traefik](https://traefik.io) plugin adds an ingress router to your swarm, letting you setup custom domain and path routing for each of your apps. You don't need to do anything else to set it up, other than choosing routing rules for each of your services!

You can view Traefik's dashboard at `http://<server-ip>:8080/dashboard/` in a browser.

:::tip
The Traefik plugin is configured run on the Manager Node (placement of `node.role == manager`). You cannot change this.
:::

### Settings

- Traefik dashboard routing pattern

  You can configure a routing path for Traefik itself, instead of going to `http://<node-ip>:8080/dashboard/` for the dashboard.

- Miasma routing pattern

  Instead of going to `http://<node-ip>:3000` to access the miasma UI, you can setup a custom domain or routing pattern to access the dashboard.
