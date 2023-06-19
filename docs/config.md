# Config

Miasma is configured through _environment variables_. Here is a list of all environment variables and their options.

[[toc]]

## `VITE_AUTH`

Setup authentication. See [Authentication](./authentication) for usage details.

Here are the valid options:

- `VITE_AUTH=token:<token>`
- `VITE_AUTH=basic:<user1>:<pass1>\n<user2>:<pass2>\n...`

<!-- ## `VITE_AUTO_UPGRADE_CRON`

**Default: `"@daily"`**

According to this CRON expression, the latest version of app images are pulled. If a newer version of the image exists, the app is upgraded to use that new image.

Internally, Miasma uses [`robfig/cron/v3`](https://pkg.go.dev/github.com/robfig/cron) to evaluate the CRON expression. Any expression supported by this library can be used for `AUTO_UPGRADE_CRON`.

By default, all apps automatically upgrade on this schedule. To prevent an app from upgrading automatically, you can set `autoUpgrade` to `false` on that app. -->
