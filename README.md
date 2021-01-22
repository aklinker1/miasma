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
    - **`meta.yml`** - Used to generate the docker compose before each swarm command, and store all application info
    - **`.env`** - Stores all env variables for the application
- `plugins/`
  - `plugin-name/`
    - **`meta.yml`** - Contains information about the 
    - **`.env`** - Stores all env variables necessary for a connected application to function

## `apps`

Applications management

### `apps`

List applications

```
miasma apps
```

### `apps:create`

Create an app for a given image

```
miasma apps:create <app-name> -i|--image <registry/image:tag>
```

### `apps:update`

Pulls and deploys the tag for the app's image

```
miasma apps:update <app-name>
```

### `apps:configure`

Configure the application's deployment information

```
miasma apps:configure <app-name> [...flags]
```

- `-h|--hide <true|false>`: Set if the application is hidden
- `-n|--networks <app-name,plugin-name,...>`: Update the list networks the application has access to. Each application/plugin gets their own network with the same name.
- `-p|--placement <labelName=value,otherKey=otherValue>`: Comma separated list of all placement rules for what node the application goes on

### `apps:destroy`

Destroy the application

```bash
miasma apps:destroy <app-name>
```

> This operation cannot be undone

## `env`

Environment variable management

### `env`

List an application's environment variables

```bash
miasma env -a|--app <app-name>
```

### `env:edit`

Open your `editor` for an interactive way of updating an app's environment variables

```bash
miasma env:edit -a|--app <app-name>
```

### `env:set`

Set environment variables for an app without using an editor

```bash
miasma env:set -a|--app <app-name> <key1=value1> <key2=value2> ...
```

### `env:remove`

Remove environment variables from an app

```bash
miasma env:remove -a|--app <app-name> <key1> <key2> ...
```

## `plugin`

Manage plugins tied to given applications. Plugins provide additional access to resources such as `PostgreSQL` and `MongoDB`.

> Plugins are just special apps. Custom plugins can be created by simply modifying application config, and putting them on the same network

### `plugin:install`

Install one of the bundled plugins

```bash
miasma plugin:remove <plugin-name>
```

### `plugin:add`

Give an application access to a plugin

```bash
miasma plugin:add <postgres|mongo> -a|--app <app-name>
```

> Simply add the application to the the plugin's network

### `plugin:remove`

Remove an application's access from a plugin

```bash
miasma plugin:remove <plugin-name> -a|--app <app-name>
```
