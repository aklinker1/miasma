---
title: Installation
---

# Installation

Miasma is setup in two easy steps:

1. Deploy the Miasma Server
1. Install the Miasma CLI on your dev computer

## Deploy the server

To deploy the server, you will need at least one device running on one of the supported architectures. Setting up this device is outside the scope of this tutorial. This device will be the main node of the cluster if you choose to add more nodes.

:::tip
My device and OS of choice a **Raspberry Pi 3b (or higher) running 64bit Ubuntu Server**.
:::

After your device is up and running, run the install script from that device:

```bash:no-line-numbers
curl -o- https://raw.githubusercontent.com/aklinker1/miasma/main/scripts/install-server.sh | bash
```

Once it finishes, the Miasma server is up and running on port `3000`! The script will also print a "join command" used in the next step to add more nodes to the cluster.

### Manual Install

If you don't trust the install script, or it did not succeed, you may have to install the dependencies and manually start the server on your main node:

1. [Install Docker](https://docs.docker.com/get-docker/)

1. Initialize the docker swarm (see [Docker's docs](https://docs.docker.com/engine/swarm/swarm-tutorial/create-swarm/) for more details)

   ```bash:no-line-numbers
   docker swarm init
   ```

1. Start the Miasma server

   ```bash:no-line-numbers
   docker run -d \
       --restart unless-stopped \
       -p 3000:3000 \
       -v /var/run/docker.sock:/var/run/docker.sock \
       -v $HOME/.miasma:/data/miasma \
       aklinker1/miasma
   ```

   > If you're unfamiliar with Docker, here's what's happening:
   >
   > - `-d`: Start the server in a background daemon rather than the foreground so the process doesn't block your terminal
   > - `--restart unless-stopped`: Restart the server unless you tell it to stop, ensuring the server restarts if it crashes or the device restarts
   > - `-p 3000:3000`: Exposes port `3000` on the main node, and mapping it to port `3000` in the container (the port the server runs on)
   > - `-v /var/run/docker.sock:/var/run/docker.sock`: Bind a special volume to the container, allowing Miasma to communicate with docker on your behalf
   > - `-v $HOME/.miasma:/data/miasma`: Bind another volume, this time the server's config directory
   > - `aklinker1/miasma`: The name of the latest stable version of the [Miasma server docker image](https://hub.docker.com/r/aklinker1/miasma/tags)

1. Open the GraphiQL playground to make sure the server is up and running: `http://<server-ip>:3000/playground`

## Install the CLI

:::danger
TODO: Right now you have to build the CLI from source
:::

<!---

You should install the CLI on any computer you want to manage apps from, or during CI.

First, use the install script to install the CLI on your `$PATH`:

```bash:no-line-numbers
curl -o- https://raw.githubusercontent.com/aklinker1/miasma/main/scripts/install-cli.sh | bash
```

Finally, connect the CLI to the miasma server running on the main node:

```bash:no-line-numbers{1}
$ miasma connect 192.168.1.0:3000
Join cluster:

  docker swarm join --token <some-token> <server-ip:port>

Connected to miasma!
```

--->

## Add More Nodes (Optional)

Adding more nodes (or devices) to the cluster is simple. Run the join command printed while connecting the CLI to the server on any devices you want to add to the cluster:

```bash:no-line-numbers
docker swarm join --token <some-token> <server-ip:port>
```

---

**Great work, that's the end of the setup!** You're ready to deploy your first app.
