<script lang="ts" setup>
import { useRoute } from "vue-router";

const route = useRoute();
const isApps = computed(() => route.path.startsWith("/apps"));
const isPlugins = computed(() => route.path.startsWith("/plugins"));
const isNodes = computed(() => route.path.startsWith("/nodes"));
</script>

<template>
  <drawer-scaffold nav-title="Miasma">
    <template #nav-items>
      <li>
        <router-link
          to="/apps"
          class="gap-2"
          :class="{
            active: isApps,
          }"
          ><i-mdi-cube />Apps</router-link
        >
      </li>
      <li>
        <router-link
          to="/plugins"
          class="gap-2"
          :class="{
            active: isPlugins,
          }"
          ><i-mdi-toolbox />Plugins</router-link
        >
      </li>
      <li>
        <router-link
          to="/nodes"
          class="gap-2"
          :class="{
            active: isNodes,
          }"
          ><i-ri-server-fill />Nodes</router-link
        >
      </li>
    </template>

    <template #drawer-items>
      <li>
        <router-link
          to="/apps"
          class="gap-2"
          :class="{
            active: isApps,
          }"
          ><i-mdi-cube />Apps</router-link
        >
      </li>
      <li>
        <router-link
          to="/plugins"
          class="gap-2"
          :class="{
            active: isPlugins,
          }"
          ><i-mdi-toolbox />Plugins</router-link
        >
      </li>
      <li>
        <router-link
          to="/nodes"
          class="gap-2"
          :class="{
            active: isNodes,
          }"
          ><i-ri-server-fill />Nodes</router-link
        >
      </li>
    </template>

    <template #content>
      <router-view />
    </template>
  </drawer-scaffold>
</template>
