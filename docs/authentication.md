# Authentication

By default, Miasma does not require authentication to access the UI. Authentication is disabled when the `MIASMA_AUTH` environment variable is not set, or set to an unknown type.

Supported authentication methods:

[[toc]]

## Token

When using token based authentication, you are required to login with a token when accessing Miasma's UI.

Set the `MIASMA_AUTH=token:<your-token>` when starting the docker image, like so:

```sh
$ docker run -d \
    --restart unless-stopped \
    -p 3000:3000 \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -e MIASMA_AUTH=token:1234 \
    aklinker1/miasma
```

Here, you'll need to enter `1234` on the UI to access the dashboard.

## Basic

When using basic authentication, you'll be required to enter a username and password.

Set the `MIASMA_AUTH=basic:<username1>:<password1>\n<username2>:<password2>\n...` environment variable.

```sh
$ docker run -d \
    --restart unless-stopped \
    -p 3000:3000 \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -e MIASMA_AUTH=basic:user1:password1\nuser2:password2 \
    aklinker1/miasma
```

Here, you'll have to either enter `user1` and `password1` or `user2` and `password2`

As of now, there are no access settings to prevent certain users from accessing certain services or settings.
