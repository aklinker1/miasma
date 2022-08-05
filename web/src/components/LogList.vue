<script lang="ts" setup>
import {
  AppLog,
  useLogsSubscription,
} from "../composition/logs-subscription.js";

const props = defineProps<{
  appId: string;
}>();

const { onResult } = useLogsSubscription(
  computed(() => ({
    appId: props.appId,
  }))
);
onResult((res) => {
  if (res.data) logs.value.push(res.data?.log);
});

const logs = ref<AppLog[]>([]);
</script>

<template>
  <div v-if="logs.length === 0" class="flex h-full">
    <button class="btn btn-ghost loading m-auto">Waiting for logs...</button>
  </div>
  <ul v-else class="p-4">
    <log-list-item v-for="log of logs" :key="log.timestamp" :log="log" />
  </ul>
</template>
