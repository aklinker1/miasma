---
id: mongo
title: MongoDB
---

Create the [MongoDB](https://hub.docker.com/_/mongo/) instance:

```bash
miasma apps:create mongo -i mongo:4
miasma env:edit -a mongo
miasma apps:configure \
    --add-placement-constraint "node.labels.database == true" \
    --add-volume /dir/path/on/physical/machine:/data/db \
    --add-target-ports 27017 --add-published-ports 27017
```

:::note
Mongo can take a few minutes to startup, so be patient. If it isn't running, you can troubleshoot by SSHing onto the manager node and running `docker service ps mongo --no-trunc`
:::

Optionally add [Mongo Express](https://hub.docker.com/_/mongo-express/) to get a UI:

```bash
miasma apps:create mongo-express -i mongo-express
miasma env:edit -a mongo-express
miasma apps:configure --add-target-ports 8081
```
