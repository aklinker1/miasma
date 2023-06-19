# Advanced Usage

[[toc]]

## Using with `docker`

Since Miasma is a simple wrapper around Docker Swarm, you can use the `docker` CLI to manage services as well. When modifying a service created or updated through the CLI (or by any means other than Miasma), Miasma will not overwrite any settings not editable in it's UI.

Be careful modifying docker-compose based services, as the next time you spin them up with docker compose, your changes will not be saved.
