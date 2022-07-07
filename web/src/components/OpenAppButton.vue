<script lang="ts" setup>
defineProps<{
  availableAt: string[];
}>();
</script>

<template>
  <a
    v-if="availableAt.length === 1"
    class="btn btn-circle btn-outline"
    title="Open in new tab"
    target="_blank"
    :href="availableAt[0]"
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
        <li v-for="link of availableAt">
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
</template>
