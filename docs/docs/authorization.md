---
title: Authorization
---

# Authorization

Miasma supports the following types of authorization:

- None (default)
- Hard-coded Access Token

## None

Miasma will not require any authorization by default. This is ideal for testing Miasma out on a single device or for a local network that you're OK with allowing access to everyone.

## Hard-coded Access Token

Require all requests to include a hard coded access token. This is ideal for a local network or a cloud hosted instance of Miasma.

:::warning Token Requirements
There are no requirements for length or complexity. If you don't want to be hacked, make sure the token is sufficiently long.
:::

When starting up the server, you can set the `ACCESS_TOKEN` environment variable:

```bash:no-line-numbers
$ cat .env
ACCESS_TOKEN=<token>

$ docker run --env-file .env \
    -d \
    --restart unless-stopped \
    -p 3000:3000 \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v $HOME/.miasma:/data/miasma \
    aklinker1/miasma
```

All requests to the API will then require an authorization header:

```text:no-line-numbers
Authorization: Bearer <token>
```

For the CLI, you'll need to reconnect and pass the token via the `--auth` flag:

```bash:no-line-numbers
miasma connect <ip:port> --auth <token>
```

And finally, the UI will require you to log in using the same token.
