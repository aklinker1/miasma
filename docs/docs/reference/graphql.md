---
title: GraphQL API
---

# GraphQL API

The Miasma server ships with a GraphQL API for accessing and editing apps. Once you're server has been started, you can access it at:

`http://<server-ip>:3000/graphql`

[[toc]]

## Playground

The API includes a GraphiQL playground that lets you experiment with the API from the browser. No additional tools necessary! You can access it at:

`http://<server-ip>:3000/playground`

## Introspection

The API also supports introspection so you can get docs inside a HTTP client like Insomnia or Postman.

## Changesets

To perform partial updates of objects, the GraphQL API contains some input types with the `Changes` suffix, like `AppChanges`.

When using these types, the API will only update the fields provided in the JSON. That means excluding a field will not set it to `null`. To set a field to `null`, you would have to include the field as `null` for the API to make that change.

For example, if we want to change the app name, remove the group, and leave everything else as is:

```graphql:no-line-numbers
# Query
mutation updateApp($id: ID!, $changes: AppChanges!) {
    editApp(id: $id, changes: $changes) {
        ...
    }
}
```
```json:no-line-numbers
// Variables
{
    "id": "...",
    "changes": {
        "name": "New Name",
        "group": null,
    }
}
```

{{ schema }}
