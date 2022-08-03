<script lang="ts" setup>
import { useAppTasksQuery } from "../composition/app-tasks-query.js";

const props = defineProps<{
  appId: string;
}>();

const vars = computed(() => ({
  appId: props.appId,
}));
const { result } = useAppTasksQuery(vars);
</script>

<template>
  <ul>
    <task-list-item v-for="task of result?.tasks" :task="task" />
  </ul>
</template>
