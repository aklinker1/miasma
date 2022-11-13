<script lang="ts" setup>
const nodeId = useNodeIdPathParam();
const { data: service } = useDockerServiceQuery(nodeId);
</script>

<template>
  <div class="text-2xl breadcrumbs grow overflow-x-hidden">
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
        <service-icon v-if="service" :name="service.Spec?.Name ?? ''" />
        <span v-if="service">{{ service.Spec?.Name ?? 'Unknown' }}</span>
      </li>
    </ul>
  </div>
</template>
