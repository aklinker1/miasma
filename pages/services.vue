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

      <div class="join">
        <!-- Create Button -->
        <create-service class="join-item" />

        <!-- Filter -->
        <filter-button button-class="rounded-l-none" />
      </div>
    </div>

    <!-- Service Lists -->
    <nuxt-page />
  </div>
</template>
