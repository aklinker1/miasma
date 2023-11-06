<script lang="ts" setup>
const props = defineProps<{
  to: string;
  title: string;
  enabled: boolean;
  disabled?: boolean;
  icon: string;
}>();

const emit = defineEmits<{
  'update:enabled': [enabled: boolean];
}>();

const enabledVModel = useVModel(props, 'enabled', emit);
</script>

<template>
  <li
    class="w-full flex items-center rounded-lg bg-neutral bg-opacity-0 hover:bg-opacity-10 transition-all"
    :class="{
      'ring-1 ring-neutral': !enabledVModel,
      'ring ring-primary bg-primary bg-opacity-5': enabledVModel,
    }"
  >
    <nuxt-link :to="to" class="flex gap-6 flex-1 p-6 items-center">
      <div class="text-3xl" :class="icon" />

      <p class="flex-1 truncate">{{ title }}</p>
    </nuxt-link>

    <label class="flex">
      <span class="sr-only">Toggle Traefik</span>
      <input
        v-model="enabledVModel"
        type="checkbox"
        class="toggle checked:toggle-primary"
        :disabled="disabled"
      />
    </label>

    <nuxt-link :to="to" class="p-6">
      <div class="i-mdi-chevron-right text-2xl" />
    </nuxt-link>
  </li>
</template>
