<script lang="ts" setup>
import { useAppIdParam } from "../composition/app-id-param";
import { useAppDetailsQuery } from "../composition/app-details-query";

const appId = useAppIdParam();
const { app } = useAppDetailsQuery(appId);
</script>

<template>
  <div class="text-2xl breadcrumbs mb-8 grow">
    <ul v-if="!app">
      <li>Apps</li>
    </ul>
    <ul v-else>
      <li><router-link to="/apps">Apps</router-link></li>
      <li class="space-x-2">
        <app-icon :name="app.name" />
        <span>{{ app.name }}</span>
      </li>
    </ul>
  </div>
</template>
