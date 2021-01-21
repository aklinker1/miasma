# Miasma

Miasma is a docker-swarm based PasS (Platform as a Service) tool for applications to a Raspberry Pi Cluster.

## Get Started

To get started, install docker on the device you plan on being the main node (pi) of the docker swarm. This node will be in charge of orchestration.

```bash

```

# CLI

The CLI manages a file system structure:

- `apps/`
  - `app-name/`
    - **`docker-compose.yml`** - Generated file for each swarm command
    - **`meta.yml`** - Used to generate the docker compose before each swarm command, and store all application info
    - **`.env`** - Stores all env variables for the application
- `plugins/`
  - `plugin-name/`
    - **`docker-compose.yml`** - Generated file for each swarm command
    - **`meta.yml`** - Contains information about the 
    - **`.env`** - Stores all env variables necessary for a connected application to function

## `app`

Manage applications running on the swarm

### `app:list`

List applications

```
miasma app:list
```

### `app:create`

Create an app for a given image

```
miasma app:create <app-name> --image <registry/image>
```

> Generates the following data structure, then executes the docker-compose file:
> 
> - `apps/`
>   - `app-name/`
>     - **`docker-compose.yml`** - Generated file for each swarm command
>     - **`meta.yml`** - Used to generate the docker compose before each swarm command, and store all application info
>     - **`.env`** - Stores all env variables for the application
> 

### `app:update`

Pulls and deploys the `:miasma` tag for the app's image

```
miasma app:update <app-name>
```

### `app:destroy`

Destroy the application. Not recoverable.

```bash
miasma app:destroy <app-name>
```

> This kills the service in docker and moves the compose file to the trash

### `app:hide`

Hide's the application from all the lists and the dashboard.

```bash
miasma app:hide <app-name>
```

## `env`

Manage the environment variables for the application

### `env:set`

```bash
miasma env:set -a|--app <app-name> <ENV_VAR>=<ENV_VALUE>
```

### `env:remove`

```bash
miasma env:remove -a|--app <app-name> <ENV_VAR>
```

### `env:list`

```bash
miasma env:list -a|--app <app-name>
```

## `plugin`

Manage plugins tied to given applications. Plugins provide additional access to resources such as `PostgreSQL` and `MongoDB`.

### `plugin:install`

Install a plugin for use.

```bash
miasma plugin:remove <plugin-name>
```

### `plugin:add`

Add a plugin to link the application to another resource.

```bash
miasma plugin:add -a|--app <app-name> <postgres|mongo> <plugin-name>
```

> Simply add the application to the the plugin's network

### `plugin:remove`

Remove a plugin from an application

```bash
miasma plugin:remove -a|--app <app-name> <plugin-name>
```
