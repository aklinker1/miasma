<script lang="ts" setup>
defineProps<{ service: Docker.Service }>();

const { mutate: stopApp, isLoading } = useDockerStopServiceMutation();
</script>

<template>
  <button
    class="btn btn-circle btn-outline hover:btn-error"
    :class="{ loading: isLoading }"
    :disabled="isLoading"
    title="Stop"
    @click.prevent="stopApp(service)"
  >
    <div v-if="!isLoading" class="i-mdi-stop text-2xl" />
  </button>
</template>
