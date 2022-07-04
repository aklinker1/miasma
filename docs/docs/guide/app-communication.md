---
title: App Communication
---

# App Communication

When multiple apps are running, they can communicate without leaving the cluster's network via `<protocol>://<app-name>:<target-port>`.

For example, lets say you have a Postgres database and a web app that needs to access it.

```bash:no-line-numbers
miasma apps:create postgres -i postgres-alpine
miasma apps:edit -a postgres --add-target-ports 5432

miasma apps:create web-app -i <my-custom-image>
```

You should use `postgres://postgres:5432` when connecting to the database. 

:::tip App name transformation
The hostname of an app will the the name, but lower-cased and with dashes instead of spaces:

- `Postgres 1` &rarr; `postgres-1`
- `Personal Website` &rarr; `personal-website`
- `Mongo` &rarr; `mongo`
:::
