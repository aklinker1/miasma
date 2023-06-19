# Get Started

[[toc]]

## Prerequisites

Before installing Miasma, you'll need a machine with [`docker` installed](https://docs.docker.com/get-docker/). My device of choice is a _Raspberry Pi 4 with Ubuntu Server_, but Miasma support almost all hardware and operating systems.

:::info
If you choose to [add more nodes](#add-more-nodes), this machine will be the ["manager node"](https://docs.docker.com/engine/swarm/manage-nodes/) of the swarm.
:::

If you can `ssh` into the device and run `docker --version`, you're good to go!

## Initialize Docker Swarm

Miasma is based around docker swarm, so you'll need to initialize the swarm on your manager node.

Make sure you're `ssh`ed into the machine, and run the following command:

```sh
$ docker swarm init
```

## Start Miasma

Miasma is a simple docker image, and can be started with `docker run`. Here's the minimial configuration required:

```sh
$ docker run -d \
    --restart unless-stopped \
    -p 3000:3000 \
    -v /var/run/docker.sock:/var/run/docker.sock \
    aklinker1/miasma
```

:::details What are all those parameters for?

- `-d`

  Runs the container in the background without blocking your SSH session

- `--restart unless-stopped`

  Always restart the container (on reboot, crash, etc) until you explicitly call `docker container stop <id>`

- `-v /var/run/docker.sock:/var/run/docker.sock`

  Bind a special volume that lets Miasma talk to Docker running on the host machine so it can do things like start/stop services

- [`aklinker1/miasma`](https://hub.docker.com/r/aklinker1/miasma)

  The name of the docker image to run. To use a different version of the image, include a `:tag` suffix

:::

---

And that's it, you're done! You should be able to visit `http://<node-ip>:3000` and see the Miasma UI.
