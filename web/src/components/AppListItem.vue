<script lang="ts" setup>
import IDatabase from "~icons/mdi/database";
import ICube from "~icons/mdi/cube";
import IEarth from "~icons/mdi/earth";
import ITools from "~icons/mdi/tools";
import { useStartAppMutation } from "../composition/start-app-mutation";
import { useStopAppMutation } from "../composition/stop-app-mutation";

const props = defineProps<{
  id: string;
  name: string;
  status: string;
  running?: number;
  total?: number;
  simpleRoute?: string;
}>();

const isRunning = computed(() => props.status === "running");
const isStopped = computed(() => props.status === "stopped");

const openLink = computed(() => {
  if (isRunning.value) {
    if (props.simpleRoute) return `https://${props.simpleRoute}`;
  }
  return undefined;
});

const iconKeywords = {
  database: ["database", "postgres", "sql", "mongo", "redis"],
  website: ["web", "www", ".com"],
  tool: ["daemon", "worker", "helper", "scheduler"],
};

const icon = computed(() => {
  const name = props.name.toLowerCase();
  if (iconKeywords.database.find((w) => name.includes(w))) return IDatabase;
  if (iconKeywords.website.find((w) => name.includes(w))) return IEarth;
  if (iconKeywords.tool.find((w) => name.includes(w))) return ITools;
  return ICube;
});

const a = useStopAppMutation();
const { mutate: stopApp, loading: stoppingApp } = useStopAppMutation();
const { mutate: startApp, loading: startingApp } = useStartAppMutation();
</script>

<template>
  <tr class="hover">
    <th>
      <app-icon class="w-6 h-6 ml-2" :name="name" />
    </th>
    <td>
      <router-link :to="`/apps/${name}`">
        <p class="text-lg">
          {{ name }}
        </p>
        <p
          class="text-sm text-success uppercase"
          :class="{
            'text-success': isRunning,
            'text-error font-medium': isStopped,
            'text-warning': total != null && total !== running,
          }"
        >
          {{ status }}
        </p>
      </router-link>
    </td>
    <td>
      <p v-if="total" class="text-lg text-center">
        <span>{{ running ?? 0 }}</span>
        /
        <span>{{ total }}</span>
      </p>
    </td>
    <td class="space-x-3 text-right">
      <a
        v-if="openLink"
        class="btn btn-circle btn-outline"
        title="Open in new tab"
        target="_blank"
        :href="openLink"
      >
        <i-mdi-open-in-new class="w-5 h-5" />
      </a>
      <button
        v-if="isRunning"
        class="btn btn-circle btn-outline hover:btn-error"
        :class="{
          'loading disabled': stoppingApp,
        }"
        title="Stop"
        @click="stopApp({ id })"
      >
        <i-mdi-stop v-if="!stoppingApp" class="w-5 h-5" />
      </button>
      <button
        v-else
        class="btn btn-circle btn-outline hover:btn-success"
        :class="{
          'loading disabled': startingApp,
        }"
        title="Start"
        @click="startApp({ id })"
      >
        <i-mdi-play v-if="!startingApp" class="w-5 h-5" />
      </button>
    </td>
  </tr>
</template>
