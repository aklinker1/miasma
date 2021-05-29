---
id: get-started
title: Get Started
slug: /
---

Miasma is setup in 2 steps:

1. Deploy the [server](/docs/server)
2. Install the [CLI](/docs/cli)

## Deploy the server

To deploy the server, you will need a device running on one of the supported architectures. Setting up this device is outside the scope of this tutorial.

:::note
My device and OS of choice a **Raspberry Pi 3b running 64bit Ubuntu Server**.
:::

This device will be the first manager node of the docker swarm. After your device is up and running, all that's left is to run the install script on that device!

```bash
$ bash | curl https://github.com/aklinker1/miasma/releases/current/install-server.sh
```

The install script might take a while if you don't have docker installed already. Once it finishes, the server should up and running on port `3000`! To verify the installation, we're going to ping the health endpoint of the REST API.

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

If everything worked correctly, you should see a `"swarm"` field in the JSON response, signifying that a swarm was initialized and the server is ready to go! You're now ready to [install the CLI](#install-the-cli)

### Manual Install

If you don't trust the install script, or it did not succeed, you may have to install the dependencies and manually start the server.

1. [Install Docker](https://docs.docker.com/get-docker/)
1. [Initialize the docker swarm](https://docs.docker.com/engine/swarm/swarm-tutorial/create-swarm/)
1. Start the server
   ```bash
   docker run -d \
     --restart unless-stopped \
     -p 3000:3000 \
     -v /var/run/docker.sock:/var/run/docker.sock \
     -v $HOME/.miasma:/data/miasma \
     aklinker1/miasma
   ```

> If you're unfamiliar with Docker, here's what's happening:
> 
> - **`-d`**: Start the server in the background in a daemon rather than the foreground
> - **`--restart unless-stopped`**: Restart the server unless it is manually stopped, ensuring the server restarts if it crashes or the device restarts
> - **`-p 3000:3000`**: Exposing port 3000 on the host, and mapping it to port 3000 in the container (the port the server runs on)
> - **`-v /var/run/docker.sock:/var/run/docker.sock`**: Binding a special volume to the container, allowing miasma to communicate with docker
> - **`-v $HOME/.miasma:/data/miasma`**: Bind another volume, this time the server's data directory, to the the `~/.miasma` directory of the user that ran the docker run command
> - **`aklinker1/miasma`**: The name of the image to run, in this case the latest stable version of the Miasma server. See [Docker Hub](https://hub.docker.com/r/aklinker1/miasma/tags) for additional tags that could be used

## Install the CLI

You can install the CLI on any computer you want to manage apps from. In most cases, that your main dev computer, but it can also be installed during CI.

1. Use the install script
   ```bash
   $ bash | curl https://github.com/aklinker1/miasma/releases/current/install-cli.sh
   ```
1. Connect to the server, using the IP address or hostname pointing to the machine the server is installed on
   ```bash
   $ miasma connect 192.168.1.0:3000
   Connected to miasma!
   192.168.1.0:3000 added to /home/user/.miasma.yaml
   ```

And you're done installing the CLI! Checkout the next page to create and manage your first app.