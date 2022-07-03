---
title: Create Your First App
---

# Create Your First App

If you haven't already, follow the install instructions to spin up a Miasma Server and install the CLI.

We'll assume you're running the server on you dev machine, and access everything via `localhost`.

[[toc]]

## Connect to the Miasma Server

Before managing any apps, you first need to tell the CLI where your server is at:

```bash:no-line-numbers
$ miasma connect localhost:3000

Connected to miasma!
  Miasma Server Version: 1.1.0
  Docker Version: 20.10.17
  Cluster Enabled? true
  Join Command: docker swarm join --token <some-token> 192.168.0.3:2377

✔ Connected to server
```

Nice! We're connected, so lets create an app.

## Creating an App

Let's deploy an example web app that just responds with information about the incoming request using the [`ealen/echo-server`](https://hub.docker.com/r/ealen/echo-server) image.

:::tip Why <code>ealen/echo-server</code>?
Because it runs on the port specified by the `PORT` environment variable. For more details, learn how Miasma [manages ports automatically](/guide/port-management).
:::

```bash:no-line-numbers{1}
$ miasma apps:create example-web-app -i ealen/echo-server

Creating example-web-app...
✔ example-web-app started
```

List the apps. We now have one app, and it's deployed to at `http://<server-ip>:3001`.

```:no-line-numbers
$ miasma apps

Apps
 ● example-web-app (http://localhost:3001, 1/1 running)

(1 total)
```

:::tip
If you see `0/1 running`, wait a few seconds and run `misama apps` again.
:::

Visit <http://localhost:3001> in your browser, and you should see a JSON response.

```json:no-line-numbers
{
    "http": {
        "baseUrl": "",
        "method": "GET",
        "originalUrl": "/some/path",
        "protocol": "http"
    },
    // ...
}
```

The app is up and running with a single CLI command!

---

#### Next Steps

- Read how Miasma [automatically manages ports](/guide/port-management)
- Learn how to [communicate between apps](/guide/app-communication)
- See more complex, common [example apps](/guide/examples)
- Learn about [plugins](/plugins)
- Checkout the [CLI reference](/reference/cli) for advanced usage
- Setup a server that [requires authorization](/authorization)
