<script lang="ts" setup>
import { useStartAppMutation } from "../composition/start-app-mutation";
import { useStopAppMutation } from "../composition/stop-app-mutation";
import { App } from "../composition/list-apps-query";

const props = defineProps<{
  app: App;
}>();

const isRunning = computed(() => props.app.status === "running");
const isStopped = computed(() => props.app.status === "stopped");

const { mutate: stopApp, loading: stoppingApp } = useStopAppMutation();
const { mutate: startApp, loading: startingApp } = useStartAppMutation();
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
            'text-error font-medium': isStopped,
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
      <template v-if="isRunning && app.availableAt != null">
        <a
          v-if="app.availableAt.length === 1"
          class="btn btn-circle btn-outline"
          title="Open in new tab"
          target="_blank"
          :href="app.availableAt[0]"
        >
          <i-mdi-open-in-new class="w-5 h-5" />
        </a>
        <div v-else class="overflow-visible btn-circle inline">
          <div class="dropdown dropdown-end">
            <label tabindex="0" class="btn btn-circle btn-outline">
              <i-mdi-open-in-new class="w-5 h-5" />
            </label>
            <ul
              tabindex="0"
              class="dropdown-content menu p-2 shadow bg-base-200 rounded-box min-w-[13rem]"
            >
              <li v-for="link of app.availableAt">
                <a
                  :href="link"
                  target="_blank"
                  class="link link-primary link-hover"
                  >{{ link }}</a
                >
              </li>
            </ul>
          </div>
        </div>
        <!-- TODO: Support multiple links in a dropdown -->
        <!-- <a
          v-else
          class="btn btn-circle btn-outline"
          title="Open in new tab"
          target="_blank"
          :href="availableAt![0]"
        >
          <i-mdi-open-in-new class="w-5 h-5" />
        </a> -->
      </template>
      <button
        v-if="isRunning"
        class="btn btn-circle btn-outline hover:btn-error"
        :class="{
          'loading disabled': stoppingApp,
        }"
        title="Stop"
        @click="stopApp(app)"
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
        @click="startApp(app)"
      >
        <i-mdi-play v-if="!startingApp" class="w-5 h-5" />
      </button>
    </td>
  </tr>
</template>
