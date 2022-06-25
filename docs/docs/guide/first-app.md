---
title: Create Your First App
---

# Create Your First App

In this walk-through, we're going to cover:

- Basic CLI usage to create, update, and delete apps
- How ports are automatically allocated and how to override that behavior

For this walk through, the main node and Miasma server will be located at `192.168.1.0:3000`. The IP address will probably be different for you.

:::details HTTPie
This walk through will use [`httpie`](https://httpie.io/) to make requests to web apps we deploy because it formats JSON output, making responses easier to read.

You can use any HTTP client to follow along (`curl`, Insomnia, Postman, etc). We'll just be doing GET requests to the apps we create, so the examples should be easy to understand regardless of the method used.
:::

### Connect to Your Miasma Server

Before managing any apps, you first need to tell the CLI where your server is at. When you're not connected to a server, you'll see something like this when you run a command.

```bash:no-line-numbers{1}
$ misama apps
Miasma CLI is not connected to a server yet. Run 'miasma connect <ip:port>'
```

Run the connect command, then list the apps again via `miasma apps`:

```bash:no-line-numbers{1,4}
$ miasma connect 192.168.1.0:3000
Connected to miasma!

$ misama apps
List apps:
(0 total)
```

Nice! We're connected, but we don't have any apps yet... Lets change that!

### Creating an App

For our first app, we're going to deploy a simple web app that just responds with information about the request it recieved. We'll be using the [`ealen/echo-server`](https://hub.docker.com/r/ealen/echo-server) image for this app.

:::tip Why <code>ealen/echo-server</code>?
Because it runs on the port specified by the `PORT` environment variable.

We'll talk about ports a little further down. By default, Miasma expects your application to run on the port provided by this environment variable.
:::

```bash:no-line-numbers{1}
$ miasma apps:create example-web-app -i ealen/echo-server
Creating example-web-app...
Starting...
Done!
```

List the apps again. We now have one app, and it's deployed to at <http://192.168.1.0:3001>.

```:no-line-numbers{1}
$ misama apps
List apps:
 - example-web-app (:3001, 1/1)
(1 total)
```

:::tip
If you see `0/1`, that means 0 out of the 1 instances are running. You might have to wait a few seconds before an instance is started.
:::

:::danger
If it still doesn't start after a few seconds, make sure the image supports at least one of your nodes' architecture. It won't start the app if no nodes can run an image.
:::

Make a request to port `3001` and you'll get a response from our echo server:

```bash:no-line-numbers{1}
$ http http://192.168.1.0:3001/some/path
HTTP/1.1 200 OK
Connection: keep-alive
Content-Length: 566
Content-Type: application/json; charset=utf-8
Date: Sun, 14 Feb 2021 01:49:09 GMT
ETag: W/"236-LPIRPcp4mNwblDJqXzAWPaIm5cI"
Keep-Alive: timeout=5

{
    "http": {
        "baseUrl": "",
        "method": "GET",
        "originalUrl": "/some/path",
        "protocol": "http"
    },
    ...
}
```

The app is up and running with a single CLI command!

### Ports

[Like Heroku](https://devcenter.heroku.com/articles/runtime-principles#web-servers), Miasma will automatically manage ports for you. But unlike Heroku, you can override the app's **target** and **published** ports.

:::details About Docker Ports
Docker (and thus Miasma) has two types of ports:

1. **Target**: the port inside the container
2. **Published**: the publicly exposed port you would use to interact with the app

For more info, see Docker's [official documentation](https://docs.docker.com/compose/compose-file/compose-file-v3/#ports)
:::

In our `example-web-app`, Miasma automatically configured the app to use a published port of `3001`. Behind the scenes, it also generated a random target port, then ran the app with that target port as the `PORT` environment variable.

Most of the time, you don't need to override the published port, let Miasma manage that one. On the other hand, overriding the target port is very useful when the image you're using doesn't respect the `PORT` environment variable.

### Hosting a Website

A common way to host a website in Docker is with the [nginx](https://hub.docker.com/_/nginx/) image. **This image does not respect the `PORT` environment variable** and it is difficult to configure so that it does. Instead, it's much easier to tell Miasma to target port `80`, Nginx's default port.

Our website will use an Nginx based hello world image, [`nginxdemos/hello:plain-text`](https://hub.docker.com/r/nginxdemos/hello:plain-text/):

```bash:no-line-numbers{1,6,11}
$ miasma apps:create example-website -i nginxdemos/hello:plain-text
Creating example-website...
Done!

$ misama apps
List apps:
 - example-web-app (:3001, 1/1)
 - example-website (:3002, 1/1)
(2 total)

$ http 192.168.1.0:3002

http: error: ConnectionError: HTTPConnectionPool(host='192.168.1.0', port=3002): Max retries exceeded with url: / (Caused by NewConnectionError('<urllib3.connection.HTTPConnection object at 0x7f6a6b2caa60>: Failed to establish a new connection: [Errno 111] Connection refused')) while doing GET request to URL: http://192.168.1.0:3002/
```

As you can see, we fail to connect to nginx. Here's whats happening:

```text:no-line-numbers
192.168.1.0:3002 ───┬─── docker:$PORT ─── nothing
                    ╰─── docker:80    ─── nginx
```

Miasma deployed the image assuming it will have something running on the `PORT` environment variable. So when we made the request, it didn't make the request to port `80`, but to the port specified by the environment variable.

Instead, we need to tell Miasma that we should actually be targeting port `80` inside the container, not some random number. We do that using the `apps:edit` command and `--add-target-ports` flag:

```bash
$ miasma apps:edit -a example-website --add-target-ports 80
Updating example-website...
Done!
```

Now when we make the same request, we get a response!

```bash:no-line-numbers{1}
$ http 192.168.1.0:3002
HTTP/1.1 200 OK
Cache-Control: no-cache
Connection: keep-alive
Content-Length: 141
Content-Type: text/plain
Date: Sun, 14 Feb 2021 02:28:36 GMT
Expires: Sun, 14 Feb 2021 02:28:35 GMT
Server: nginx/1.13.8

Server address: 10.0.0.121:80
Server name: 41048f8da61a
Date: 14/Feb/2021:02:28:36 +0000
URI: /
Request ID: e4a54530d02d6069d23cd8cc396332ae
```

### Stopping Apps

Now that we have two apps running, lets stop `example-website`.

```bash:no-line-numbers{1}
$ miasma apps:stop
Stopping example-website...
Done!
```

:::tip
Here when we called `apps:stop` without the `-a` flag, but Miasma still used `example-website`.

This is because Miasma remembers which app was last used in the previous command, and uses it in the next command if the `-a|-app` flag is excluded.
:::

Now when we make a request, it's not connecting because the app is not running, not because of ports

```bash:no-line-numbers{1}
$ http 192.168.1.0:3002

http: error: ConnectionError: HTTPConnectionPool(host='192.168.1.0', port=3002): Max retries exceeded with url: / (Caused by NewConnectionError('<urllib3.connection.HTTPConnection object at 0x7f6a6b2caa60>: Failed to establish a new connection: [Errno 111] Connection refused')) while doing GET request to URL: http://192.168.1.0:3002/
```

### Starting Apps

If you want to start it back up:

```bash:no-line-numbers{1}
$ miasma apps:start
Starting example-website...
Done!
```

Apps are automatically started when they get created. If an app is updated, it will automatically restart if it was running before, otherwise it will remain stopped.

### Deleting Apps

Before you start deploying your own apps, lets quickly clean up these two example apps so you're starting with a clean slate:

```bash:no-line-numbers{1,5}
$ miasma apps:delete -a example-web-app
Destroying example-web-app...
Done!

$ miasma apps:delete -a example-website
Destroying example-website...
Done!
```

<br />

---

<br />

You now know the basics of managing apps. Next steps?

- Look through the [Example Apps](/examples.md)
- Learn about [plugins](/plugins/index.md)
- Checkout the [CLI reference](/docs/cli) for advanced usage
