<script lang="ts" setup>
useHead({
  title: 'Nodes',
});
</script>

<template>
  <div class="space-y-4">
    <div class="flex justify-between items-center">
      <node-breadcrumbs />
      <add-node />
    </div>

    <!-- Node List -->
    <nuxt-page />
  </div>
</template>
