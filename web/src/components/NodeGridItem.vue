<script lang="ts" setup>
import { Node } from "../composition/nodes-query";

const props = defineProps<{
  node: Node;
}>();

const labels = computed(() => {
  return Object.entries(props.node.labels).map(([key, value]) => {
    if (value) return `${key}=${value}`;
    return key;
  });
});
</script>

<template>
  <div class="flex flex-col rounded-lg shadow-2xl overflow-hidden">
    <div class="py-3 px-4 bg-base-200 space-y-1">
      <p class="text-xl font-medium text-primary text-opacity-100">
        {{ node.hostname }}
      </p>
      <p class="text-sm uppercase">
        {{ node.ip }}&ensp;&bull;&ensp;{{ node.os }} {{ node.architecture }}
      </p>
      <p class="space-x-2">
        <span v-for="label of labels" :key="label" class="badge font-medium">
          {{ label }}
        </span>
      </p>
    </div>

    <ul class="menu rounded-box p-2">
      <li v-for="service of node.services" :key="service.name">
        <router-link :to="`/apps/${service.name}`">
          <app-icon class="w-6 h-6" :name="service.name" />
          {{ service.name }}
        </router-link>
      </li>
    </ul>
  </div>
</template>
