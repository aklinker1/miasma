<script lang="ts" setup>
defineProps<{
  title?: string;
  error: H3Error<Docker.ErrorResponse> | null;
  retry: () => void;
}>();
</script>

<template>
  <div v-if="error" class="alert alert-error shadow-lg overflow-hidden">
    <div>
      <div class="i-mdi-information shrink-0 text-2xl" />
      <div class="space-y-1">
        <p class="font-bold">{{ title ?? error.message }}</p>
        <p class="text-sm">{{ error.data.message }}</p>
      </div>
    </div>
    <div class="flex-none">
      <button class="btn btn-sm" @click="retry()">Retry</button>
    </div>
  </div>
</template>
