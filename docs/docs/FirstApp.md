---
id: first-app
title: Create Your First App
description: Simple walk through for deploying some hello world docker images to your new misama server
---

In this walk-through, we're going to cover:

- Basic CLI Usage
- Creating apps
- Destroying apps
- Port customization and rules

For this walk through, the main node and Miasma server will be located at `192.168.1.0:3000`. The IP address will be different for you.

:::note
This walk through will use [`httpie`](https://httpie.io/) to make requests to the apps we create because it formats JSON output, making responses easier to read.

You can use any HTTP client to follow along (`curl`, Insomnia, Postman, etc). We'll just be doing GET requests to the apps, so the examples should be easy to understand regardless of the HTTP client used
:::

### Connect to a Miasma Server

Make sure you're connected to a miasma server. To test this, you can run any command.

```bash
$ misama apps
Miasma CLI is not connected to a server yet. Run 'miasma connect <ip:port>'
```

If you're not connected, go ahead and run the `connect` command, followed by the `apps` command again.

```bash
$ miasma connect 192.168.1.0:3000
Connected to miasma!
192.168.1.0:3000 added to /home/user/.miasma.yaml

$ misama apps
List apps:
(0 total)
```

Nice! We're connected, but we don't have any apps yet... Lets change that!

### Creating an App

For our first app, we're going to deploy a simple web app that returns JSON info about the request it received. We'll be using [`ealen/echo-server` from Docker Hub](https://hub.docker.com/r/ealen/echo-server) for this app.

We're using it because this image respects the `PORT` environment variable. We'll talk about ports more in the next section.

```bash
$ miasma apps:create example-web-app -i ealen/echo-server
Creating example-web-app...
Starting...
Done!
```

List the apps again. We now have one app, and it's deployed to at <http://192.168.1.0:3001>.

```
$ misama apps
List apps:
 - example-web-app (:3001, 1/1)
(1 total)
```

:::note
If you see `0/1`, that means 0 out of the 1 instances are running. You might have to wait a few seconds before an instance is started.
:::

Make a request to port `3000` and you'll get a response from our echo server:

```bash
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

[Like Heroku](https://devcenter.heroku.com/articles/runtime-principles#web-servers), Miasma will automatically manage ports for you. But unlike Heroku, you can override the app's target and published ports.

:::info Docker Ports
Docker (and thus Miasma) has two types of ports:

1. **Target**: the port inside the container
2. **Published**: the publicly exposed port

For more info, see Docker's [official documentation](https://docs.docker.com/compose/compose-file/compose-file-v3/#ports)
:::

In our `example-web-app`, Miasma automatically configured the app to use a published port of `3001`. Behind the scenes, it also generated a random target port, then ran the app with that target port as the `PORT` environment variable.

Most of the time, you don't need to override the published port, let Miasma manage that one. On the other hand, overriding the target port is very useful, especially when hosting a website.

### Hosting a Website

Most websites in Docker are hosted using [Nginx](https://hub.docker.com/_/nginx/), which always runs on port `80`. It can be configured to run on a port specified by the environment, but not easily. Instead, it makes more sense to just point toward port `80`.

Our website will use an Nginx based hello world image: [`nginxdemos/hello:plain-text`](https://hub.docker.com/r/nginxdemos/hello:plain-text/)

```bash
$ miasma apps:create example-website -i nginxdemos/hello:plain-text
Creating example-website...
Starting...
Deployed at http://192.168.1.0:3002
Done!

$ misama apps
List apps:
 - example-web-app (:3001, 1/1)
 - example-website (:3002, 1/1)
(2 total)

$ http 192.168.1.0:3002

http: error: ConnectionError: HTTPConnectionPool(host='192.168.1.0', port=3002): Max retries exceeded with url: / (Caused by NewConnectionError('<urllib3.connection.HTTPConnection object at 0x7f6a6b2caa60>: Failed to establish a new connection: [Errno 111] Connection refused')) while doing GET request to URL: http://192.168.1.0:3002/
```

As you can see, our request doesn't even make it to the app and fails to establish a connection.

Here's whats happening:

```text
192.168.1.0:3002 ───┬─── docker:$PORT ─── nothing
                    ╰─── docker:80    ─── nginx
```

Our request is taking the top flow (`192.168.1.0:3002` &rarr; `docker:$PORT` &rarr; `nothing`), when we need it to take the bottom (`192.168.1.0:3002` &rarr; `docker:80` &rarr; `nginx`), so we configure the target port to be `80`.

```bash
$ miasma apps:configure -a example-website --add-target-ports 80
Updating example-website...
Done!
```

Now, when we make the same request, we get a response!

```bash
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

Now that we have two echo servers running, lets stop `example-website`.

```bash
$ miasma apps:stop
Stopping example-website...
Done!
```

:::info
Notice when we call `apps:stop` without the `-a` flag, Miasma remembers which app was last used. Generally, you can exclude the flag when you're working with a single app.
:::

Now when we make a request, it's not connecting because the app is not running, not because of ports

```
$ http 192.168.1.0:3002

http: error: ConnectionError: HTTPConnectionPool(host='192.168.1.0', port=3002): Max retries exceeded with url: / (Caused by NewConnectionError('<urllib3.connection.HTTPConnection object at 0x7f6a6b2caa60>: Failed to establish a new connection: [Errno 111] Connection refused')) while doing GET request to URL: http://192.168.1.0:3002/
```

### Starting Apps

If you want to start it back up:

```bash
$ miasma apps:start
Starting example-website...
Done!
```

### Destroying Apps

Before you start deploying your own apps, lets quickly clean up these two test apps so you're starting with a clean slate:

```bash
$ miasma apps:destroy -a example-web-app
Destroying example-web-app...
Done!

$ miasma apps:destroy -a example-website
Destroying example-website...
Done!
```
