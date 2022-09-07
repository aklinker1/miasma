---
title: Private Registries
---

# Private Registries

Miasma is based off Docker Swarm, and Docker Swarm supports private registries, so that means **Miasma supports private registries as well**!

### How does it work in Docker Swarm?

Normally, you log into the private registries on your cluster's manager nodes, and use the `--with-registry-auth` flag when creating a "service" (Swarm service = Miasma app).

See Docker's documentation for more details: <https://docs.docker.com/engine/swarm/services/#create-a-service-using-an-image-on-a-private-registry>

### How does it work in Miasma?

Currently, Miasma does not support a `--with-registry-auth` equivalent since it uses the Docker Engine API to interact with docker, not the CLI.

> I haven't figured out how to pass this flag, so if you have an idea, open a PR!

So instead of just logging into the private registry on your manager nodes, you need to do it on all nodes in your cluster.

```bash:no-line-numbers
ssh username@<machine-ip>
docker login myprivateregistry.com
exit

ssh username@<machine-2-ip>
docker login myprivateregistry.com
exit

ssh username@<machine-3-ip>
# ...
```

After you've logged into the registry on all the nodes, you can create an app from the CLI or UI using the `myprivateregistry.com/some/image:tag` image name:

```bash:no-line-numbers
miasma apps:create Example -i myprivateregistry.com/some/image:tag
```
