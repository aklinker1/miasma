<script lang="ts" setup>
import { useNodesQuery } from "../composition/nodes-query";

const { result, loading, error } = useNodesQuery();
</script>

<template>
  <div class="node-grid">
    <!-- Loading -->
    <div v-if="loading && result == null" class="loading m-8" />

    <!-- Error -->
    <p v-else-if="error" class="text-error">{{ error }}</p>

    <!-- Nodes -->
    <template v-else>
      <node-grid-item
        v-for="node of result?.nodes"
        :key="node.id"
        :node="node"
      />
    </template>
  </div>
</template>

<style scoped>
.node-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(17rem, 1fr));
  gap: 2rem;
}
</style>
