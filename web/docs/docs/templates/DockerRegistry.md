---
id: docker-registry
title: Docker Registry
---

Host an insecure docker registry.

Docker Hub has gone crazy with their pulling enforcement, so custom registries for small/private images make sense.

```bash
# Create the app
miasma apps:create docker-registry -i registry:2.7

# Configure it's ports and placement onto node with logs of storage space
miasma apps:configure -a docker-registry \
    --add-target-ports 5000 \
    --add-placement "node.labels.database == true"

# Give it a route
miasma route:set --hostname docker.hostname.com

# 
miasma volume:add \
    --target /var/lib/registry/ \
    --source /physical/path/to/docker/docker-registry/data
```

Because this is an insecure registry, you have to configure docker to allow it by editing the daemon config.

```json title="/etc/docker/daemon.json"
{
  "insecure-registries": ["docker.hostname.com"]
}
```

Then restart the computer so the changes take effect.

If you want to use the registry as a source for apps in miasma, each node also needs to update their daemon config and restart.
