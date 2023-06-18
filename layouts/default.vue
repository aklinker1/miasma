<script lang="ts" setup>
import { routes } from '~/utils/routes';

const route = useRoute();

const isServices = computed(() => route.path.startsWith(routes.services));
const isNodes = computed(() => route.path.startsWith(routes.nodes));
const isPlugins = computed(() => route.path.startsWith(routes.plugins));
const isLogin = computed(() => route.path.startsWith(routes.login));

const { error, refetch } = useDockerSwarmInfoQuery();
</script>

<template>
  <drawer-scaffold nav-title="Miasma">
    <template #nav-items v-if="!isLogin">
      <li>
        <nuxt-link
          :to="routes.services"
          class="gap-2"
          :class="{ 'active text-primary': isServices }"
        >
          <div class="i-mdi-cube text-2xl" />
          <span>Services</span>
        </nuxt-link>
      </li>
      <li>
        <nuxt-link :to="routes.nodes" class="gap-2" :class="{ 'active text-primary': isNodes }">
          <div class="i-ri-server-fill text-2xl" />
          <span>Nodes</span>
        </nuxt-link>
      </li>
      <li>
        <nuxt-link :to="routes.plugins" class="gap-2" :class="{ 'active text-primary': isPlugins }">
          <div class="i-mdi-puzzle text-2xl" />
          <span>Plugins</span>
        </nuxt-link>
      </li>
    </template>

    <template #drawer-items v-if="!isLogin">
      <li>
        <nuxt-link
          :to="routes.services"
          class="gap-2"
          :class="{ 'active text-primary': isServices }"
        >
          <div class="i-mdi-cube text-2xl" />
          <span>Services</span>
        </nuxt-link>
      </li>
      <li>
        <nuxt-link :to="routes.nodes" class="gap-2" :class="{ 'active text-primary': isNodes }">
          <div class="i-ri-server-fill text-2xl" />
          <span>Nodes</span>
        </nuxt-link>
      </li>
      <li>
        <nuxt-link :to="routes.plugins" class="gap-2" :class="{ 'active text-primary': isPlugins }">
          <div class="i-mdi-puzzle text-2xl" />
          <span>Plugins</span>
        </nuxt-link>
      </li>
    </template>

    <template #content>
      <error-display title="Problem with Docker Swarm" :error="error" :retry="refetch" />

      <router-view />
    </template>
  </drawer-scaffold>
</template>
