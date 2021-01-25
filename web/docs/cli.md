## `miasma`

```
Manage and deploy dockerized applications to a docker swarm

Usage:
  miasma [flags]
  miasma [command]

Available Commands:
  apps            List apps
  apps:configure  Update an application's properties
  apps:create     Create and deploy a new application
  apps:destroy    Destroy an existing application
  apps:start      Start an application
  apps:stop       Stop a running application
  plugins         List plugins
  plugins:install Install and start a pre-defined plugin
  plugins:remove  Stop and remove an installed plugin

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
Update an application's properties such as target ports, networks, etc. See the list of
flags for all the properties that can be set for an application.

It is worth noting that for properties that are lists, there is no add or remove. Instead, include
all the values for an array property you would like to change:

  miasma app:configure --app app-name --ports 80,22
	
Only the properties specified in the flags will update be updated. To remove a propterty, pass in an empty string for the value:

  miasma app:configure --app app-name --ports ""

Usage:
  miasma apps:configure [flags]

Flags:
  -a, --app string   The app to perform the action on

Global Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
```

### `apps:create`

```
Create and deploy a new application

Usage:
  miasma apps:create [flags]

Flags:
  -h, --hidden         Whether or not the app is hidden
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

### `plugins`

```
List plugins

Usage:
  miasma plugins [flags]

Global Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
```

### `plugins:install`

```
Install one of the pre-defined plugins: PostgreSQL, Mongo
	
Plugins are simple, pre-configured, reusable applications. Every plugin can also be defined as an
app.

Usage:
  miasma plugins:install [flags]

Global Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
```

### `plugins:remove`

```
Stop and remove an installed plugins: PostgreSQL, Mongo
	
Plugins are simple, pre-configured, reusable applications. Every plugin can also be defined as an
app.

Usage:
  miasma plugins:remove [flags]

Global Flags:
      --config string   config file (default is $HOME/.miasma.yaml)
```

