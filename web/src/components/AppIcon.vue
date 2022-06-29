<script lang="ts" setup>
import IDatabase from "~icons/mdi/database";
import ICube from "~icons/mdi/cube";
import IEarth from "~icons/mdi/earth";
import ITools from "~icons/mdi/tools";

const props = defineProps<{
  name: string;
}>();

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
</script>

<template>
  <icon />
</template>
