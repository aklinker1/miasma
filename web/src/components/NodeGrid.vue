<script lang="ts" setup>
import { useNodesQuery } from "../composition/nodes-query";

const { result, loading, error } = useNodesQuery();
</script>

<template>
  <div class="flex gap-8 flex-wrap justify-center">
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
        class="w-[20rem] shrink-0"
      />
    </template>
  </div>
</template>
