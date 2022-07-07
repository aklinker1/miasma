<script lang="ts" setup>
import { App } from "../composition/list-apps-query";
import { useIsAppRunning } from "../composition/is-app-running";

const props = defineProps<{
  app: App;
}>();

const isRunning = useIsAppRunning(computed(() => props.app));
</script>

<template>
  <tr class="hover">
    <th>
      <app-icon class="w-6 h-6 ml-2" :name="app.name" />
    </th>
    <td>
      <router-link :to="`/apps/${app.id}`">
        <p class="text-lg">
          {{ app.name }}
        </p>
        <p
          class="text-sm text-success uppercase"
          :class="{
            'text-success': isRunning,
            'text-error font-medium': !isRunning,
            'text-warning':
              app.instances?.total != null &&
              app.instances.total !== app.instances.running,
          }"
        >
          {{ app.status }}
        </p>
      </router-link>
    </td>
    <td>
      <p v-if="app.instances?.total" class="text-lg text-center">
        <span>{{ app.instances.running }}</span>
        /
        <span>{{ app.instances.total }}</span>
      </p>
    </td>
    <td class="space-x-3 text-right">
      <open-app-button v-if="isRunning" :available-at="app.availableAt" />
      <stop-app-button v-if="isRunning" :app-id="app.id" />
      <start-app-button v-else :app-id="app.id" />
    </td>
  </tr>
</template>
