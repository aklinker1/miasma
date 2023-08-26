<script lang="ts" setup>
useHead({
  title: 'Nodes',
});
</script>

<template>
  <div class="space-y-4">
    <div class="flex justify-between items-center">
      <node-breadcrumbs />

      <div class="join">
        <!-- Add -->
        <add-node class="join-item" />

        <!-- Filter -->
        <filter-button button-class="rounded-l-none" />
      </div>
    </div>

    <!-- Node List -->
    <nuxt-page />
  </div>
</template>
