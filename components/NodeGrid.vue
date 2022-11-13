<script lang="ts" setup>
const { data, isLoading, error } = useDockerNodesQuery();
</script>

<template>
  <div>
    <!-- Loading -->
    <div v-if="isLoading" class="flex py-32 items-center justify-center">
      <cube-spinner class="m-auto" />
    </div>

    <!-- Error -->
    <p v-else-if="error" class="text-error">{{ error }}</p>

    <div v-else class="node-grid">
      <node-grid-item v-for="node of data" :key="node.ID" :node="node" />
    </div>
  </div>
</template>

<style scoped>
.node-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(17rem, 1fr));
  gap: 2rem;
}
</style>
