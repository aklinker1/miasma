# Migrate to V2

If you're coming from version 1, there's one major difference: no more CLI. Instead, everything is done through the UI directly.

## Restart Docker Container

Miasma is still just a docker container, but the data volume is no longer needed - Miasma stores everything it needs inside docker services directly. The rest stays the same

```sh
# Get the ID of the currently running miasma instance and stop it
docker container ls
docker stop <container-id>

# Pull the latest image, and restart the container
docker pull aklinker1/miasma
docker run -d \
    --restart unless-stopped \
    -p 3000:3000 \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v $HOME/.miasma:/data/miasma \ // [!code --]
    aklinker1/miasma
```

## Migrate Apps... Sike!

You shouldn't have to do anything to migrate your v1 apps to v2. As long as you don't kill the swarm, they're just docker services that will keep running, and they will show up in v2 with no additional steps.

## Missing Features

Some features are missing from V2:

- Traefik Plugin (ingress router + automatic HTTPS certs)

I will be adding these with time. I'll probably forget to update this page, so check [Introduction](./introduction#future-work) for a more up-to-date list of all features.
