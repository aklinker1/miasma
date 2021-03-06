---
id: introduction
title: Introduction
slug: /intro
---

Miasma was designed as a simple way to manage applications running on my cluster of Raspberry Pis in my closet.

It is made of two components: [lightweight REST API server](/docs/server), and a [CLI tool](/docs/cli). The server is responsible for managing everything docker related, and CLI talks with the server, simplifying the interaction API calls to a CLI tool.

## Features

- Create, manage, and destroy apps from any computer
- Custom hostname and path routing
- Multiple architecture support (`arm/v7`, `arm64`, `amd64`)
- Clusters

### Future Work

- Web based dashboard as an alternative to the CLI
- Horizontal scaling

### Not Features

- **Built-in TLS/SSL** - this means no HTTPS
- **Complete docker service configuration** - some functionality, like advanced network configuration, is missing

:::caution Miasma was not designed for an Enterprise environment

If you're considering using Miasma to manage some kind of production environment in AWS or Digital Ocean, try it out in a locally first. Understand what it can and can't do before deciding if it's right for your stack. It's probably not.
:::
