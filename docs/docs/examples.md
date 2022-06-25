---
title: Example Apps
---

# Example Apps

Here are some example apps setup and configured through the [Miasma CLI](/reference/cli.md).

[[toc]]

## Docker Registry

Host an insecure [docker registry](https://hub.docker.com/_/registry/).

```bash:no-line-numbers
# Create the app
miasma apps:create docker-registry -i registry:2.7

# Configure it's ports, node placement, and mounted volume for data storage
miasma apps:edit -a docker-registry \
    --add-target-ports 5000 \
    --add-placement-constraint "node.labels.database == true" \
    --add-volume /dir/path/on/physical/machine:/var/lib/registry

# Give it a route
miasma route:set --hostname docker.hostname.com
```

Because this is an insecure registry, you have to [configure docker to allow it](https://docs.docker.com/registry/insecure/) by editing the Docker daemon config:

```json
// /etc/docker/daemon.json
{
  "insecure-registries": ["docker.hostname.com"]
}
```

Then restart the computer so the changes take effect.

If you want to use the registry as a source for Miasma apps, each node also needs to update their daemon config and restart.

## MongoDB

Create the [MongoDB](https://hub.docker.com/_/mongo/) instance:

```bash:no-line-numbers
miasma apps:create mongo -i mongo:4
miasma env:edit -a mongo
miasma apps:edit \
    --add-placement-constraint "node.labels.database == true" \
    --add-volume "/dir/path/on/physical/machine":/data/db \
    --add-target-ports 27017 --add-published-ports 27017
```

:::tip
Mongo can take a few minutes to startup, so be patient. If it isn't running, you can troubleshoot by SSHing onto the manager node and running `docker service ps miasma-mongo --no-trunc`
:::

### Mongo Express

Optionally, you can also setup [Mongo Express](https://hub.docker.com/_/mongo-express/) to get a UI for inspecting your mongo app:

```bash:no-line-numbers
miasma apps:create mongo-express -i mongo-express
miasma env:edit -a mongo-express
miasma apps:edit --add-target-ports 8081
```

## Redis
 
Create the [Redis](https://hub.docker.com/_/redis/) instance. Nothing special here, Redis respects the `PORT` environment variable.

```bash:no-line-numbers
miasma apps:create redis -i redis:6-alpine
```

## Postgres

Create the [PostgreSQL](https://hub.docker.com/_/postgres/) database. The Postgres docker image can be quite finicky with it's port, so it's best to just set the target and published ports to `5432`, Postgres's default.

```bash:no-line-numbers
miasma apps:create postgres -i postgres:13-alpine
miasma env:edit -a postgres
miasma apps:edit \
    --add-placement-constraint "node.labels.database == true" \
    --add-volume "/dir/path/on/physical/machine":/var/lib/postgresql/data \
    --add-target-ports 5432 --add-published-ports 5432
```
