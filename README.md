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

## Development Setup

```bash
# Create a swarm locally
docker swarm init
# Create an network for miasma to run on
docker network create --driver overlay miasma-default
```

# TODO

- [ ] Edit Service Ports
- [ ] Edit Service Environment
- [ ] Edit Service Volumes
- [ ] Edit Service Logs
- [ ] Edit Service Tasks
- [ ] Edit Node Labels
