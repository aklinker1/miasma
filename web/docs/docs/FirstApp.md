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

For any API calls to the apps we create, we'll be using [`httpie`](https://httpie.io/) instead of `curl` because it formats JSON output, making responses easier to understand. You can use `curl`, Postman, Insomnia or any other tool to make the requests, but since we're only going to be doing GETs, the examples should be clear regardless of the tool.

:::note
For this entire example, the Miasma server will be located at `192.168.1.0`. It will most likely be different for you.
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

For our first app, we're going to deploy a simple web server that returns JSON info about the request it received. We'll be using [`ealen/echo-server` from Docker Hub](https://hub.docker.com/r/ealen/echo-server) for this app:

```bash
$ miasma apps:create echo -i ealen/echo-server
Creating echo...
Starting...
Deployed at http://192.168.1.0:3001
Done!

$ misama apps
List apps:
 - echo
(1 total)
```

As you can see, we now have one app, and it's deployed to at <http://192.168.1.0:3001>. If you then make a request to that IP and port, you'll get a response from the app!

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

So with absolutely no configuration, we were able to get a web app up and running!

However, it is worth nothing that **`ealen/echo-server` was used for a very specific reason**: it's server respects `PORT` enviornment variable.

Miasma, like Heroku, runs all applications with the `PORT` environment variable set. This variable, like Heroku, is managed by Misama by default. Miasma starts at port `3001` and goes up to `4000` as more apps are added.

Here is a major difference between Heroku and Miasma, this is the default behavior, **but unlike Heroku, the target port can be changed**.

Lets create another app using another popular echo image: [`nginxdemos/hello:plain-text`](https://hub.docker.com/r/nginxdemos/hello:plain-text/). The key difference between this server and `ealen/echo-server` is that it does not respect the `PORT`, and instead exposes port 80.

```bash
$ miasma apps:create echo-2 -i nginxdemos/hello:plain-text
Creating echo-2...
Starting...
Deployed at http://192.168.1.0:3002
Done!

$ misama apps
List apps:
 - echo
 - echo-2
(2 total)

$ http 192.168.1.0:3002

http: error: ConnectionError: HTTPConnectionPool(host='192.168.1.0', port=3002): Max retries exceeded with url: / (Caused by NewConnectionError('<urllib3.connection.HTTPConnection object at 0x7f6a6b2caa60>: Failed to establish a new connection: [Errno 111] Connection refused')) while doing GET request to URL: http://192.168.1.0:3002/
```

As you can see, we get a 404 instead of a response. This is because the `PORT` variable isn't pointing to the correct port, so the server is never getting hit, and we get a `404 Not Found`.

It's very easy to fix this. Simply configure the app to use a target port of 80!

```bash
$ miasma apps:configure -a echo-2 --add-target-ports 80
Updating echo-2...
Done!
```

Then, if we make the same request, we get a response!

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

:::info
`miasma apps:configure --add-target-ports 80` is a really important shortcut to know. Most of the time when dockerizing a web app, the base image is nginx, just like this example.

There are tutorials for making nginx respect the `PORT` variable, but it complcates your docker files, and it's easier in the long run to just set the target port rather than fight nginx.
:::

Now that we have two echo servers running, lets stop one, they both don't need to be running. Notice how when we call stop, we don't have to pass the `-a` flag. Miasma remembers what app a command was last run for, so you can exclude that flag when you're working on a single app.

```bash
$ miasma apps:stop
Stopping echo-2...
Done!

$ http 192.168.1.0:3002

http: error: ConnectionError: HTTPConnectionPool(host='192.168.1.0', port=3002): Max retries exceeded with url: / (Caused by NewConnectionError('<urllib3.connection.HTTPConnection object at 0x7f6a6b2caa60>: Failed to establish a new connection: [Errno 111] Connection refused')) while doing GET request to URL: http://192.168.1.0:3002/
```

And if you want to restart it:

```bash
$ miasma apps:start
Starting echo-2...
Done!
```

And that's it! Before you start deploying your own apps, lets quickly clean up these two test apps so you're starting with a clean slate!

```bash
$ miasma apps:destroy -a echo
Destroying echo...
Done!

$ miasma apps:destroy -a echo-2
Destroying echo-2...
Done!
```
