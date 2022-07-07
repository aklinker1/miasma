---
title: Server Configuration
---

# Server Configuration

The server can be configured through environment variables. The following is a list of all the environment variables and what they do.

[[toc]]

:::tip
All environment variables are optional, their default values are listed
:::

## `ACCESS_TOKEN`

**Default: `""`**

When set to something other than an empty string, it enables server wide authorization. See the [authorization docs](/authorization) for more details

## `AUTO_UPGRADE_CRON`

**Default: `"@daily"`**

According to this CRON expression, the latest version of app images are pulled. If a newer version of the image exists, the app is upgraded to use that new image.

Internally, Miasma uses [`robfig/cron/v3`](https://pkg.go.dev/github.com/robfig/cron) to evaluate the CRON expression. Any expression supported by this library can be used for `AUTO_UPGRADE_CRON`.

By default, all apps automatically upgrade on this schedule. To prevent an app from upgrading automatically, you can set `autoUpgrade` to `false` on that app.

## `PORT`

**Default: `3000`**

The port the miasma server runs on. It can be any valid port, but it is recommended to stay away from 3001-4000, where apps are published.
