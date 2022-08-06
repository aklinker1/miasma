<img src="./.github/assets/readme_branding.png" alt="Miasma - The cloud for your closet" height="50%">

<p align="center">
  <a href="https://github.com/aklinker1/miasma/actions/workflows/validate.yml">
    <img src="https://github.com/aklinker1/miasma/actions/workflows/validate.yml/badge.svg" alt="Validate" />
  </a>
  <a href="https://hub.docker.com/r/aklinker1/miasma">
    <img src="https://img.shields.io/badge/Docker Hub-aklinker1%2Fmiasma-success?logo=docker&logoColor=aaa" alt="Docker Hub" />
  </a>
  <a href="https://github.com/aklinker1/miasma/releases">
    <img src="https://img.shields.io/badge/CLI-Latest-success?logo=github&logoColor=aaa" alt="Install the CLI" />
  </a>
  <a href="https://aklinker1.github.io/miasma">
    <img src="https://img.shields.io/badge/Documentation-blue" alt="Documentation" />
  </a>
</p>

Miasma is a self-hosted, docker based, Heroku-like PaaS with cluster and ARM support.

###### 😖 What does that even mean?

- ***Self-hosted***: You deploy the [`aklinker1/miasma` docker image](https://hub.docker.com/r/aklinker1/miasma) on your own infrastructure, and own everything
- ***Docker based***: Apps are based off of [Docker images](https://docs.docker.com/develop/)
- ***Heroku-like [PaaS](https://en.wikipedia.org/wiki/Platform_as_a_service)***: Deploy and manage apps from a [UI (work in progress)](./.github/assets/ui.png) or CLI, similar to Heroku
- [***Cluster***](https://en.wikipedia.org/wiki/Computer_cluster): Scale the platform horizontally by using multiple computers/devices to run your apps
- ***ARM support***: In addition to the standard CPU architectures, Miasma is [published with ARM support](https://hub.docker.com/r/aklinker1/miasma/tags)

> Checkout the [full feature list](https://aklinker1.github.io/miasma/#features)!

#### 🚪 Why is it the "cloud for your closet"?

I originally wrote Miasma as a simple way to deploy apps to some Raspberry Pis in my closet on my local network!

Since then, features have been added to make it an *almost* production ready service. I'm hosting <https://aklinker1.io> (an echo server) using Miasma on Digital Ocean, and will be using it for some personal services in the future!

#### ✨ What makes Miasma unique?

- The only self hosted PaaS that supports both **clusters** and **ARM devices**
- Most apps can be deployed with zero config

### 📘 Docs

Head over to the docs to learn more or get started: <https://aklinker1.github.io/miasma>

### 👋 Contributing

<a href="https://github.com/aklinker1/miasma/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=aklinker1/miasma" />
</a>

Check out the [contributing docs](https://aklinker1.github.io/miasma/docs/contributing) to get your local development environment setup.
