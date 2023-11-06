<script lang="ts" setup>
const route = useRoute();

const plugin = computed(() => {
  if (route.path === '/plugins/traefik') return 'Traefik';
  return undefined;
});
const icon = computed(() => {
  if (plugin.value === 'Traefik') return 'i-devicon-traefikproxy';
  return undefined;
});
</script>

<template>
  <div class="text-2xl font-stylized breadcrumbs grow overflow-x-hidden">
    <ul v-if="!plugin">
      <li>Plugins</li>
    </ul>
    <ul v-else>
      <li>
        <nuxt-link
          :to="routes.plugins"
          class="link link-hover opacity-50 hover:opacity-100 active:opacity-100"
        >
          <span>Plugins</span>
        </nuxt-link>
      </li>
      <li class="flex gap-2 items-center">
        <div class="text-2xl" :class="icon" />
        <span>{{ plugin }}</span>
      </li>
    </ul>
  </div>
</template>
