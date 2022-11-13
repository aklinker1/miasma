<script lang="ts" setup>
const route = useRoute();

const isApps = computed(() => route.path.startsWith('/services'));
const isNodes = computed(() => route.path.startsWith('/nodes'));
const isLogin = computed(() => route.path.startsWith('/login'));

const { error, refetch } = useDockerSwarmInfoQuery();
</script>

<template>
  <drawer-scaffold nav-title="Miasma">
    <template #nav-items v-if="!isLogin">
      <li>
        <nuxt-link to="/services" class="gap-2" :class="{ active: isApps }">
          <div class="i-mdi-cube text-2xl" />
          <span>Services</span>
        </nuxt-link>
      </li>
      <li>
        <nuxt-link to="/nodes" class="gap-2" :class="{ active: isNodes }">
          <div class="i-ri-server-fill text-2xl" />
          <span>Nodes</span>
        </nuxt-link>
      </li>
    </template>

    <template #drawer-items v-if="!isLogin">
      <li>
        <nuxt-link to="/services" class="gap-2" :class="{ active: isApps }">
          <div class="i-mdi-cube text-2xl" />
          <span>Services</span>
        </nuxt-link>
      </li>
      <li>
        <nuxt-link to="/nodes" class="gap-2" :class="{ active: isNodes }">
          <div class="i-ri-server-fill text-2xl" />
          <span>Nodes</span>
        </nuxt-link>
      </li>
    </template>

    <template #content>
      <error-display title="Problem with Docker Swarm" :error="error" :retry="refetch" />

      <router-view />
    </template>
  </drawer-scaffold>
</template>
