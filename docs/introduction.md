# Introduction

## Background

Miasma is a simple UI wrapping the docker swarm APIs. Miasma was designed as a simple way to manage applications running on a cluster of Raspberry Pis in my closet.

![UI](https://github.com/aklinker1/miasma/raw/main/.github/assets/ui.png)

## Features

- Create, manage, and delete services
- Supports all major CPU architectures (`arm/v7`, `arm64`, `amd64`)
- Add multiple devices to form a cluster
- Horizontal scaling

### Future Work

- Built-in hostname and path routing via Traefik
- Automatically managed HTTPS certificates
- Watchtower plugin for upgrading services

## Not Features

Miasma does not, and will never, provide complete docker service configuration. See Portainer if this is what you're looking for. It provides a simple and opinionated way of managing applications.

:::warning ⚠️&ensp;Miasma was not designed for an enterprise environment
If you're considering using Miasma to manage some kind of production environment in the cloud, try it out locally first. Understand what it can and can't do before deciding if it's right for your stack.
:::
