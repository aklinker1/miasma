<script lang="ts" setup>
const pluginName = usePluginNamePathParam();
</script>

<template>
  <div class="text-2xl font-stylized breadcrumbs grow overflow-x-hidden">
    <ul v-if="!pluginName">
      <li>Plugins</li>
    </ul>
    <ul v-else>
      <li>
        <nuxt-link
          :to="routes.nodes"
          class="link link-hover opacity-50 hover:opacity-100 active:opacity-100"
        >
          <span>Plugins</span>
        </nuxt-link>
      </li>
      <li class="space-x-2">
        <span>{{ pluginName }}</span>
      </li>
    </ul>
  </div>
</template>
