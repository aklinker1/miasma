---
title: Installation
---

# Installation

[[toc]]

## Deploy the Miasma Server

Run the following commands on the main machine you want to host applications on. If you are planning on adding more nodes to the cluster, this device will be the main manager node.

:::tip Machine OS
Before starting, make sure the machine you're installing the server on is one of the [supported architectures](https://hub.docker.com/r/aklinker1/miasma/tags).

For example, `linux/arm64` is published, so if you want to install the server on a Raspberry Pi, it needs to be running a 64bit OS, like Ubuntu Server 64bit.
:::

1. Install Docker
   ```bash:no-line-numbers
   curl -fsSL https://get.docker.com -o get-docker.sh
   sh get-docker.sh
   ```
1. Initialize the swarm (**this is required** even if you aren't planning on using multiple nodes)
   ```bash:no-line-numbers
   docker swarm init
   ```
1. Run the server's Docker image
   ```bash:no-line-numbers
   # Minimal configuration
   docker run -d \
       --restart unless-stopped \
       -p 3000:3000 \
       -v /var/run/docker.sock:/var/run/docker.sock \
       -v $HOME/.miasma/database:/data/miasma \
       aklinker1/miasma
   ```
   :::details What are all those parameters for?
   - **`-d`**
   - **`--restart unless-stopped`** - Always restart the container (on reboot, crash, etc) until you explicitly call `docker container stop <id>`
   - **`-v /var/run/docker.sock:/var/run/docker.sock`** - Bind a special volume that lets Miasma talk to the docker daemon running on the host machine so it can do things like start/stop apps 
   - **`-v $HOME/.miasma:/data/miasma`** - Bind a volume so Miasma can persist it's configuration and remember what apps have been setup if the container is stopped and restarted
   - [`aklinker1/miasma`](https://hub.docker.com/r/aklinker1/miasma) - The name of the docker image to run. To use a different version of the image, include a `:tag` suffix
   :::

And you're server is setup!

To verify the server is setup properly, open up the GraphQL playground, `http://<machine-ip>:3000/playground`, and run the following query:

```graphql:no-line-numbers
{
  health {
    version
    dockerVersion
    cluster {
      joinCommand
      createdAt
    }
  }
}
```


If everything is setup correctly, you should see a response like this:
```json:no-line-numbers
{
  "data": {
    "health": {
      "version": "1.1.0",
      "dockerVersion": "20.10.17",
      "cluster": {
        "joinCommand": "docker swarm join --token SWMTKN-1-1hgnz90jzs1xytc8chcxdlcqn1r1p38xtlus4mzyr7wx3lnuie-ejy87vthuum7hc36cw0xr4wjq 192.168.0.3:2377",
        "createdAt": "2022-05-05T07:11:01.82641567Z"
      }
    }
  }
}
```

### Add More Nodes (Optional)

1. SSH into the machine to add
  ```bash:no-line-numbers
  ssh user@<machine-ip>
  ```
2. Install Docker (same as before)
  ```bash:no-line-numbers
  curl -fsSL https://get.docker.com -o get-docker.sh
  sh get-docker.sh
  ```
3. Join the swarm using the join command from the health query
  ```bash:no-line-numbers
  docker swarm join --token <some-token> <machine-ip>:2377
  ```

Depending on the network your machines are located in, you might need to use a different IP address that what is returned from the health query.

## Install the CLI

:::warning TODO
Right now you have to build the CLI from source. See [contributing docs](/contributing) to get started, then run `make cli BINARY=miasma` to build `miasma` instead of `miasma-dev`
:::
