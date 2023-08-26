<script lang="ts" setup>
defineProps<{
  buttonClass: string;
}>();

const showSystem = useShowSystemServices();
const showHidden = useShowHiddenServices();
</script>

<template>
  <div class="dropdown dropdown-end">
    <label
      tabindex="0"
      class="btn btn-outline hover:btn-primary"
      :class="buttonClass"
      title="Filter"
    >
      <span class="i-mdi-filter text-2xl" />
    </label>

    <div tabindex="0" class="dropdown-content z-[1] p-4 w-max bg-base-200 shadow-lg mt-2 space-y-2">
      <h6 class="opacity-50 font-bold">Visibility</h6>

      <!-- System -->
      <label class="flex gap-2 items-center">
        <input
          type="checkbox"
          class="checkbox checkbox-sm checked:checkbox-primary"
          v-model="showSystem"
        />
        <span>Show system services</span>
      </label>

      <!-- Hidden -->
      <label class="flex gap-2 items-center">
        <input
          type="checkbox"
          class="checkbox checkbox-sm checked:checkbox-primary"
          v-model="showHidden"
        />
        <span>Show hidden services</span>
      </label>
    </div>
  </div>
</template>
