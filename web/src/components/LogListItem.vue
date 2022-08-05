<script lang="ts" setup>
import { AppLog } from "../composition/logs-subscription.js";

defineProps<{
  log: AppLog;
}>();
</script>

<template>
  <li class="font-mono whitespace-nowrap">
    <span class="text-opacity-50 text-base-content">{{ log.timestamp }}</span
    >&ensp;<span>{{ log.message }}</span>
  </li>
</template>
