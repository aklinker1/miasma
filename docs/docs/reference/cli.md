---
title: CLI
description: CLI commands and documentation
---

The Misama CLI should be installed on any computer that you would deploy/manage application from. It's very similar to Heroku's CLI, but also takes inspiration from the Docker CLI.

Checkout the [Get Started](/docs#install-cli) page to learn how to install the CLI

## `miasma`

```
Manage and deploy dockerized applications to a docker swarm

Usage:
  miasma [flags]
  miasma [command]

Available Commands:
  apps           List apps
  apps:configure Update an application's properties
  apps:create    Create and deploy a new application
  apps:destroy   Destroy an existing application
  apps:edit      Update an app's display properties
  apps:reload    Reload an application or start it if it's not already running
  apps:start     Start an application
  apps:stop      Stop a running application
  apps:upgrade   Pull the latest version of the application's image and reload the app
  connect        Connect to a Miasma Server
  env            Print all the environment variables for an application
  env:edit       Edit the environment variables for the application
  plugins        List plugins
  plugins:add    Install and start a pre-defined plugin
  plugins:remove Stop and remove an installed plugin
  traefik        Get routing rules for an app
  traefik:set    Set routing rules for an app

Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
  -v, --version         Print the CLI version

Use "miasma [command] --help" for more information about a command.
```

### `apps`

```
List apps

Usage:
  miasma apps [flags]

Flags:
  -A, --all   List all apps, including hidden ones

Global Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
```

### `apps:configure`

```
Update an application's properties such as target ports. See the list of flags for all the properties
that can be set for an application.

It is worth noting that for properties that are lists, there is no add or remove. Instead, include
all the values for an array property you would like to change:

  miasma apps:configure --app app-name --ports 80,22

Only the properties specified in the flags will update be updated. To remove a propterty, pass in an empty string for the value:

  miasma apps:configure --app app-name --ports ""

Usage:
  miasma apps:configure [flags]

Flags:
      --add-placement-constraint strings   Add to the list of constraints specifying which nodes can run the app
      --add-published-ports int32Slice     Add to the list of ports the app can be accessed through (default [])
      --add-target-ports int32Slice        Add to the list of ports that the app is listening to inside the container (default [])
      --add-volume strings                 Add a bound volume to the host machine
  -a, --app string                         The app to perform the action on
      --rm-placement-constraint strings    Remove from the list of constraints specifying which nodes can run the app
      --rm-published-ports int32Slice      Remove from the list of ports the app can be accessed through (default [])
      --rm-target-ports int32Slice         Remove from the list of ports that the app is listening to inside the container (default [])
      --rm-volume strings                  Remove a bound volume to the host machine

Global Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
```

### `apps:create`

```
Create and deploy a new application

Usage:
  miasma apps:create [flags]

Flags:
      --hidden         Whether or not the app is hidden
  -i, --image string   Application's docker image that is ran

Global Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
```

### `apps:destroy`

```
Destroy an existing application. If the app is running, it is stopped first

Usage:
  miasma apps:destroy [flags]

Flags:
  -a, --app string   The app to perform the action on

Global Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
```

### `apps:edit`

```
Update an app's properties such as name and group. See the list of flags for all the properties
that can be set for an application.

  miasma apps:edit --app app-name --group some-group

Usage:
  miasma apps:edit [flags]

Flags:
  -a, --app string     The app to perform the action on
  -g, --group string   Change the app's group
      --hidden         Make the app hidden
  -n, --name string    Change the app's name
      --visible        Remove the hidden flag from the app

Global Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
```

### `apps:reload`

```
Reload an application or start it if it's not already running

Usage:
  miasma apps:reload [flags]

Flags:
  -a, --app string   The app to perform the action on

Global Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
```

### `apps:start`

```
Start an application

Usage:
  miasma apps:start [flags]

Flags:
  -a, --app string   The app to perform the action on

Global Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
```

### `apps:stop`

```
Stop a running application

Usage:
  miasma apps:stop [flags]

Flags:
  -a, --app string   The app to perform the action on

Global Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
```

### `apps:upgrade`

```
Pull the latest version of the application's image and reload the app. If a new image is passed, the app is updated to use that image instead of the current one

Usage:
  miasma apps:upgrade [flags]

Flags:
  -a, --app string     The app to perform the action on
  -i, --image string   The image the app should use instead of the current image

Global Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
```

### `connect`

```
Connect to a Miasma Server

Usage:
  miasma connect [flags]

Examples:
miasma connect localhost:3000

Global Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
```

### `env`

```
Print all the environment variables for an application

Usage:
  miasma env [flags]

Flags:
  -a, --app string   The app to perform the action on

Global Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
```

### `env:edit`

```
Edit the environment variables for the application

Usage:
  miasma env:edit [flags]

Flags:
  -a, --app string   The app to perform the action on

Global Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
```

### `plugins`

```
List plugins

Usage:
  miasma plugins [flags]

Global Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
```

### `plugins:add`

```
Install one of the pre-defined plugins: Traefik

Plugins are simple, pre-configured, reusable applications. Every plugin can also be defined as an
app.

Usage:
  miasma plugins:add [flags]

Global Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
```

### `plugins:remove`

```
Stop and remove an installed plugins: Traefik

Plugins are simple, pre-configured, reusable applications. Every plugin can also be defined as an
app.

Usage:
  miasma plugins:remove [flags]

Global Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
```

### `traefik`

```
Get routing rules for an app

Usage:
  miasma traefik [flags]

Flags:
  -a, --app string   The app to perform the action on

Global Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
```

### `traefik:set`

```
Set routing rules for an app

Usage:
  miasma traefik:set [flags]

Flags:
  -a, --app string    The app to perform the action on
      --host string   The hostname of the app. EX: test.home.io
      --path string   The path at the host the app will live at. EX: /api
      --rule string   Custom traefik routing rule

Global Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
```
