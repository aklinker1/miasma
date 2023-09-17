<script setup lang="ts">
import { ref } from 'vue';

const { canToggle: canToggleTraefik, isEnabled: isTraefikEnabled } = useTraefikPlugin();
</script>

<template>
  <ul class="flex flex-col gap-8">
    <plugin-list-item
      :to="routes.traefikPlugin"
      v-model:enabled="isTraefikEnabled"
      :disabled="!canToggleTraefik"
      style="--p: var(--traefik)"
      icon="i-devicon-traefikproxy"
      title="Traefik (Reverse Proxy)"
    />
  </ul>
</template>
