<script lang="ts" setup>
import { useAppIdParam } from "../composition/app-id-param";
import { useAppDetailsQuery } from "../composition/app-details-query";

const appId = useAppIdParam();
const { app } = useAppDetailsQuery(appId);
</script>

<template>
  <div class="text-2xl breadcrumbs grow">
    <ul v-if="!app">
      <li>Apps</li>
    </ul>
    <ul v-else>
      <li>
        <router-link
          to="/apps"
          class="link link-hover opacity-50 hover:opacity-100 active:opacity-100"
          >Apps</router-link
        >
      </li>
      <li class="space-x-2">
        <app-icon :name="app.name" />
        <span>{{ app.name }}</span>
      </li>
    </ul>
  </div>
</template>
