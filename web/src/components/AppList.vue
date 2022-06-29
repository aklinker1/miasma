<script lang="ts" setup>
import { computed } from "vue";
import { useAppListQuery } from "../composition/list-apps-query";

const props = defineProps<{
  showHidden?: boolean;
}>();
const showHidden = computed(() => props.showHidden);

const { result } = useAppListQuery({ showHidden });

const groups = computed(() => {
  const apps = result.value?.apps;
  if (apps == null) return null;

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
  <div v-if="result && result.apps" class="space-y-8">
    <div
      class="overflow-x-auto w-full shadow-2xl"
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

        <tbody>
          <app-list-item
            v-for="app of apps"
            :key="app.id"
            :id="app.id"
            :name="app.name"
            :status="app.status"
            :total="app.instances.total"
            :running="app.instances.running"
          />
        </tbody>
      </table>
    </div>
  </div>
</template>
