<script lang="ts" setup>
import { useRoute } from "vue-router";
import { useAppDetailsQuery } from "../composition/app-details-query";

const route = useRoute();

const appId = computed(() => route.params["appId"] as string);
const { app } = useAppDetailsQuery(appId);
</script>

<template>
  <div v-if="!app" class="py-24">
    <div class="loading" />
  </div>

  <template v-else>
    <app-controller
      class="w-full md:w-56 md:flex-shrink-0 md:float-left md:mr-12"
      :app="app"
    />
    <app-details-form class="md:flex-grow" :app="app" />
  </template>
</template>
