<script lang="ts" setup>
const nodeId = useNodeIdPathParam();
const { data: node } = useDockerNodeQuery(nodeId, {
  enabled: computed(() => !!nodeId.value),
});
</script>

<template>
  <div class="text-2xl font-stylized breadcrumbs grow overflow-x-hidden">
    <ul v-if="!nodeId">
      <li>Nodes</li>
    </ul>
    <ul v-else>
      <li>
        <nuxt-link
          to="/nodes"
          class="link link-hover opacity-50 hover:opacity-100 active:opacity-100"
        >
          <span>Nodes</span>
        </nuxt-link>
      </li>
      <li class="space-x-2">
        <span v-if="node">{{ node.Description?.Hostname ?? node.ID ?? 'Unknown' }}</span>
      </li>
    </ul>
  </div>
</template>
