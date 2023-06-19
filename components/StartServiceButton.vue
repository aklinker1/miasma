<script lang="ts" setup>
defineProps<{
  service: Docker.Service;
}>();

const { mutate, isLoading } = useDockerStartServiceMutation();
</script>

<template>
  <button
    class="btn btn-circle btn-outline hover:btn-success"
    :class="{ loading: isLoading }"
    :disabled="isLoading"
    title="Start"
    @click.prevent="mutate(service)"
  >
    <div v-if="!isLoading" class="i-mdi-play text-2xl" />
  </button>
</template>
