---
title: Port Management
---

# Port Management

[Like Heroku](https://devcenter.heroku.com/articles/runtime-principles#web-servers), Miasma will automatically manage ports for you. But unlike Heroku, you can override the app's **target** and **published** ports.

:::details About Docker Ports
Docker (and thus Miasma) has two types of ports:

1. **Target**: the port inside the container
2. **Published**: the publicly exposed port you would use to interact with the app

For more info, see Docker's [official documentation](https://docs.docker.com/compose/compose-file/compose-file-v3/#ports)
:::

By default, a random target port will be selected and provided to the app via the `PORT` environment variable. The app should respect this environment variable, and run on the provided port.

## Published Ports

Usually you don't need to set the published port of an application unless you want to, or the image requires it.

```bash:no-line-numbers
# Add a port
miasma apps:edit -a postgres --add-published-ports 5432

# Go back to a randomized port
miasma apps:edit -a postgres --rm-published-ports 5432
```

## Target Ports

Sometimes it's difficult to configure a docker image to respect the `$PORT` environment variable. If you can't change which port the application runs on, or it's difficult to do, you can just tell the app to target a specific port instead:

```bash:no-line-numbers
# Add a port
miasma apps:edit -a example-website --add-target-ports 80

# Go back to a randomized port
miasma apps:edit -a postgres --rm-target-ports 80
```

## Expose Multiple Ports

If your image uses multiple ports, and you want to be able to access both of them, you can just add multiple published ports. Behind the scenes, multiple target ports will be randomly generated for the container, and provided to the app via the `PORT_1` (same as `PORT`), `PORT_2`, `PORT_3`, etc environment variables.

```bash:no-line-numbers
miasma apps:edit -a example --add-published-ports 8081 --add-published-ports 8082
```

To use hard coded target ports, you can also set those:

```bash:no-line-numbers
miasma apps:edit -a example --add-published-ports 8081 --add-published-ports 8082
```

:::warning Order Matters
In the example above, 8081 was added before 8082 in both cases, so the ports end up connected like `8081:8081` and `8082:8082`.

If we had instead ran:

```bash:no-line-numbers
miasma apps:edit -a example --add-published-ports 8082 --add-published-ports 8081
```

We would end up with `8081:8082` and `8082:8081` instead.
:::
