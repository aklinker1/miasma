<script lang="ts" setup>
const props = defineProps<{
  name: string;
}>();

const iconKeywords = {
  database: ['database', 'postgres', 'sql', 'mongo', 'redis'],
  website: ['web', 'www', '.com'],
  tool: ['daemon', 'worker', 'helper', 'scheduler'],
};

const icon = computed<string>(() => {
  const name = props.name.toLowerCase();
  if (iconKeywords.database.find(w => name.includes(w))) return 'i-mdi-database';
  if (iconKeywords.website.find(w => name.includes(w))) return 'i-mdi-earch';
  if (iconKeywords.tool.find(w => name.includes(w))) return 'i-mdi-tools';
  return 'i-mdi-cube';
});
</script>

<template>
  <div class="text-2xl" :class="icon" />
</template>
