<script lang="ts" setup>
const nodeId = useNodeIdPathParam();

const { data: node, error, refetch } = useDockerNodeQuery(nodeId);

useHead({
  title: node.value?.Description?.Hostname ?? 'Unknown',
});
</script>

<template>
  <error-display title="Failed to load node details" :error="error" :retry="refetch" />

  <node-details-form v-if="node" :node="node" />
</template>
