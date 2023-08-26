<script lang="ts" setup>
import { TaskFilters } from '~/composables/useDockerTasksQuery';

const props = defineProps<{
  node: Docker.Node;
}>();

const labels = computed(() => {
  return Object.entries(props.node.Spec?.Labels ?? {}).map(([key, value]) => {
    if (value) return `${key}=${value}`;
    return key;
  });
});

const taskFilters = ref<TaskFilters>({
  'desired-state': ['running'],
  node: [props.node.ID!],
});
const { data: tasks, isLoading } = useDockerTasksQuery(taskFilters);

const applyVisibilityFilter = useApplyVisibilityFilter();
const visibleTasks = computed(() => {
  return applyVisibilityFilter(tasks.value ?? [], task => task.Spec?.ContainerSpec?.Labels);
});
</script>

<template>
  <div class="flex flex-col rounded-lg shadow-2xl overflow-hidden">
    <div class="py-3 px-4 bg-base-200 space-y-1">
      <!-- Disable link while node details page is not implemented -->
      <nuxt-link :to="`/nodes/${node.ID}`" class="text-xl link link-primary link-hover block">
        {{ node.Description?.Hostname }}
      </nuxt-link>
      <p class="text-sm uppercase">
        <!-- IP -->
        <span>{{ node.Status?.Addr }}</span>
        <span>&ensp;</span>
        <span class="text-base-content text-opacity-30">&bull;</span>
        <span>&ensp;</span>
        <!-- OS -->
        <span>{{ node.Description?.Platform?.OS }}</span>
        {{ ' ' }}
        <!-- Architechure -->
        <span>{{ node.Description?.Platform?.Architecture }}</span>
      </p>
      <ul class="flex gap-2 flex-wrap">
        <li v-for="label of labels" :key="label" class="badge badge-md font-mono">
          {{ label }}
        </li>
      </ul>
    </div>

    <div v-if="isLoading" class="h-full p-2 min-h-12 flex box-content">
      <p class="text-center opacity-50 m-auto">Loading...</p>
    </div>
    <div v-else-if="!visibleTasks?.length" class="h-full p-2 min-h-12 flex box-content">
      <p class="text-center opacity-50 m-auto">0 containers</p>
    </div>
    <ul v-else class="menu rounded-box p-2">
      <node-grid-task-item v-for="task of visibleTasks" :key="task.ID" :task="task" />
    </ul>
  </div>
</template>
