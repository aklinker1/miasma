---
id: get-started
title: Get Started
slug: /
---

Miasma is setup in a couple easy steps:

1. Deploy the [server](/docs/server)
2. (Optional) Add more nodes to the cluster
3. Install the [CLI](/docs/cli) on your dev computer

## Deploy the server

To deploy the server, you will need at least one device running on one of the supported architectures. Setting up this device is outside the scope of this tutorial. This device will be the main node of the cluster if you choose to add more nodes.

:::note
My device and OS of choice a **Raspberry Pi 3b (or higher) running 64bit Ubuntu Server**.
:::

After your device is up and running, run the install script from that device:

```bash
$ curl -o- https://raw.githubusercontent.com/aklinker1/miasma/main/scripts/install-server.sh | bash
```

Once it finishes, the Miasma server is up and running on port `3000`! The script will also print a "join command" used in the next step to add more nodes to the cluster.

### Manual Install

If you don't trust the install script, or it did not succeed, you may have to install the dependencies and manually start the server on your main node:

1. [Install Docker](https://docs.docker.com/get-docker/)

1. Initialize the docker swarm (see [Docker's docs](https://docs.docker.com/engine/swarm/swarm-tutorial/create-swarm/) for more details)

   ```bash
   $ docker swarm init
   ```
   
1. Start the Miasma server

   ```bash
   $ docker run -d \
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

1. Ping the health endpoint of the server to make sure it's up and running.

   ```bash
   $ curl localhost:3000/api/health
   {
      "dockerVersion": "19.03.13-4484c46",
      "swarm": {
         "createdAt": "2020-11-15 21:44:20.898473606 +0000 UTC",
         "id": "yk39px7m8pql5apr6qx8cekak",
         "joinCommand": "docker swarm join --token <some-long-string> <miasma-ip:port>",
         "updatedAt": "2021-02-13 16:31:13.583995176 +0000 UTC"
      },
      "version": "<some-version>"
   }
   ```

   > The health endpoint also returns the "join command" used in the next step to add more nodes to the cluster.

## Add More Nodes

Adding more nodes (or devices) to the cluster is simple. Using the join command from the server install step, run the command on any devices you want to add to the cluster:

```bash
$ docker swarm join --token <some-token> <main-node-ip:port>
```

## Install the CLI

You should install the CLI on any computer you want to manage apps from, or during CI.

First, use the install script to install the CLI on your `$PATH`:

```bash
$ curl -o- https://raw.githubusercontent.com/aklinker1/miasma/main/scripts/install-cli.sh | bash
```

Finally, connect the CLI to the miasma server running on the main node:

```bash
$ miasma connect 192.168.1.0:3000
Connected to miasma!
```

---

**Great work, that's the end of the setup!** You're ready to deploy your first app.
