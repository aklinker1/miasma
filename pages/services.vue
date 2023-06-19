<script lang="ts" setup>
useHead({
  title: 'Services',
});
</script>

<template>
  <!-- Title bar -->
  <div class="space-y-4">
    <div class="flex items-center gap-8">
      <service-breadcrumbs />

      <!-- Create Button -->
      <create-service />
    </div>

    <!-- Service Lists -->
    <nuxt-page />
  </div>
</template>
