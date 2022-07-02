---
title: Introduction
---

# Introduction

Miasma was designed as a simple way to manage applications running on my cluster of Raspberry Pis in my closet.

![UI](https://github.com/aklinker1/miasma/raw/main/.github/assets/ui.png)

It has two components: a **lightweight server**, and a **CLI tool**. The server is responsible for managing everything docker related, while CLI talks with the server via it's GraphQL API.

## Features

- Create, manage, and destroy apps in docker containers
- Custom hostname and path routing
- Multiple architecture support (`arm/v7`, `arm64`, `amd64`)
- Cluster/multiple device support

## Future Work

- Web dashboard as an alternative to the CLI
- Horizontal scaling

## Not Features

- **Built-in TLS/SSL** - this means no HTTPS
- **Complete docker service configuration** - some functionality, like advanced network and volume configuration, are missing

:::warning ⚠️&emsp;Miasma was not designed for an enterprise environment

If you're considering using Miasma to manage some kind of production environment in the cloud, try it out locally first. Understand what it can and can't do before deciding if it's right for your stack.

**It's probably not**.
:::
