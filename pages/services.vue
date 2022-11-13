<script lang="ts" setup>
import { defineService } from '~~/utils/services';

useHead({
  title: 'Services',
});

const {
  mutate: _createService,
  isLoading: createPending,
  error: createError,
} = useDockerCreateServiceMutation();

function createService() {
  const service = defineService({
    TaskTemplate: {
      ContainerSpec: {
        Image: 'ealen/echo-server',
      },
    },
  });
  _createService(service);
}
</script>

<template>
  <!-- Title bar -->
  <div class="space-y-4">
    <div class="flex items-center gap-8">
      <service-breadcrumbs />

      <!-- Create Button -->
      <button
        class="btn btn-outline hover:btn-primary gap-2"
        :class="{ loading: createPending }"
        :disabled="createPending"
        @click="createService()"
      >
        <span>Create</span>
        <div class="i-mdi-plus text-2xl" />
      </button>
    </div>

    <error-display :error="createError" :retry="createService" title="Failed to create service" />

    <!-- Service Lists -->
    <nuxt-page />
  </div>
</template>
