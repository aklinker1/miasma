<script lang="ts" setup>
import { computed } from "vue";
import { App, useAppListQuery } from "../composition/list-apps-query";

const props = defineProps<{
  showHidden?: boolean;
}>();
const showHidden = computed(() => props.showHidden);
const clusterIpAddress = location.hostname;

const { result, loading, error } = useAppListQuery({
  showHidden,
  clusterIpAddress,
});

const groups = computed<Array<[string, App[]]>>(() => {
  const apps = result.value?.apps;
  if (apps == null || apps.length === 0) return [["", []]];

  return Object.entries<any[]>(
    apps.reduce((grouped: Record<string, any>, app: any) => {
      const g = app.group?.toLowerCase() ?? "";
      grouped[g] ??= [];
      grouped[g].push(app);
      return grouped;
    }, {})
  );
});
</script>

<template>
  <!-- Loading -->
  <div v-if="loading && result == null" class="loading m-8" />

  <!-- Error -->
  <p v-else-if="error" class="text-error">{{ error }}</p>

  <!-- Apps Table -->
  <div v-if="result" class="space-y-8">
    <div
      class="w-full shadow-2xl rounded-lg"
      v-for="[group, apps] of groups"
      :key="group"
    >
      <table class="table w-full">
        <thead>
          <tr>
            <th></th>
            <th class="w-full">{{ group || "Apps" }}</th>
            <th class="text-center">Instances</th>
            <th></th>
          </tr>
        </thead>

        <!-- Main list -->
        <tbody>
          <app-list-item v-for="app of apps" :key="app.id" :app="app" />
        </tbody>
      </table>

      <!-- Empty message -->
      <p v-if="apps.length === 0" class="text-center w-full p-6">No apps</p>
    </div>
  </div>
</template>
