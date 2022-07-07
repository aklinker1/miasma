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

## `PORT`

**Default: `3000`**

The port the miasma server runs on. It can be any valid port, but it is recommended to stay away from 3001-4000, where apps are published.
