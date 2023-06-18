<script lang="ts" setup>
defineProps<{
  title?: string;
  error: H3Error<Docker.ErrorResponse> | null;
  retry: () => void;
}>();
</script>

<template>
  <!-- <div class="alert shadow-lg">
  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-info shrink-0 w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
  <div>
    <h3 class="font-bold">New message!</h3>
    <div class="text-xs">You have 1 unread message</div>
  </div>
  <button class="btn btn-sm">See</button>
</div> -->

  <div v-if="error" class="alert alert-error shadow-lg overflow-hidden">
    <div class="i-mdi-information shrink-0 text-2xl" />
    <div class="space-y-1">
      <p class="font-bold">{{ title ?? error.message }}</p>
      <p class="text-sm">{{ error.data.message }}</p>
    </div>
    <button class="btn btn-sm" @click="retry()">Retry</button>
  </div>
</template>
