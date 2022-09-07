---
title: Troubleshooting
---

# Troubleshooting

## App Not Starting

There are several reasons why your app may not be starting:

### 1. There is not a suitable machine in your cluster that can run the application's docker image

Docker images are built for specific CPU architectures. For example, an image might only support `x86_64`, but all your machines in the cluster are Raspberry Pis, which are all `armv7`. In this case, your app could not start because it cannot find a suitable node to run on, because none have the `x86_64` architecture.

To see if that's the case, you can look at the "tasks" for your app from the web dashboard. If so, you will see an error like this:

> 1/1: no suitable node (unsupported platform on X nodes)

### 2. App is misconfigured

Your app could be configured in a way that is preventing the app from starting. An example of this would be binding volumes that don't exist.

Once again, to see if this is the problem, look at the app's "tasks" on the web dashboard. If this is the case, an error will clearly be shown stating what the problem is.

### 3. The app is crashing on startup

This could be because you haven't provided required environment variables, or because there's a bug. In either case, you can look at the app's logs to figure out why it's crashing and fix it.

Viewing logs inside Miasma is coming soon. For now, you can use the Docker CLI to see them:

```bash:no-line-numbers
ssh username@<manger-node-ip>
docker service ls # To get the service name/id
docker service logs <service-name-or-id> --follow
```

### 4. Ports are misconfigured

If the app is running, but you can't access/connect to it, the ports are likely misconfigured, or the `$PORT`/`$PORT_1`, `$PORT_2`, `$PORT_3`, etc environment variables are not respected.

See the [port management docs](./port-management.md) for more details on configuring the ports used by your application.
