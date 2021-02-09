# Miasma

Miasma is a docker-swarm based PasS (Platform as a Service) tool for applications to a Raspberry Pi Cluster.

Miasma consists of two parts:

- `miasma-server` - Docker container ran on the main node of the docker swarm, in charge of orchestrating the applications running in the swarm
- `miasma-cli` - Used on any computer you want to manage and deploy applications

## Get Started

### Miasma Server

> Eventually this process will just be an install script you curl down and run in a single line

First, install docker on the device you plan on being the main node (pi) and initialize the swarm. This node will be running the miasma server.

```bash
# Install Docker
# TODO...

# Initialize the swarm
docker swarm init
```

Take note of the worker join token, it will be used later on to add more Pis to the swarm.

Finally, spin up the `miasma-server`:

```bash
docker run -d \
    --restart unless-stopped \
    -p 3000:3000 \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v $HOME/.miasma:/data/miasma \
    aklinker1/miasma
```

> `-d` starts the container in a daemon
>
> `--restart unless-stopped` will make sure miasma is restarted on error or if the pi is power cycled. If you stop this image
>
> `-p 3000:3000` exposes the server's ports. By default, it runs on port 3000
>
> `-v /var/run/docker.sock:/var/run/docker.sock` gives miasma access to docker so it can manage the swarm
>
> `-v $HOME/.miasma:/data/miasma` sets the directory where all the data is stored to `~/.miasma`

At this point, you can move on to installing the CLI on your other computers.

If you have more devices you would like to add to the swarm, install `docker` on the other devices (the same way you did on the main node), then run the below command to join the swarm.

```bash
docker swarm join --token <worker-join-token> <miasma-server-ip:3000>
```

> You can do this at any point after starting the miasma server. Miasma works with any number of nodes and will automatically balance the applications accross the available nodes as more are added to the swarm

That's it! Once a node is apart of the swarm, the basic setup is done.

### Miasma CLI

```bash
# Install and make sure it's working
go get -u github.com/aklinker1/miasma
miasma version

# Connect to the server
miasma connect <miasma-server-ip:3000>
```

After you've connected the CLI to a miasma instance, you can start deploying applications!
