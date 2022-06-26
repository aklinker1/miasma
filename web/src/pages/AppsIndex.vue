<script lang="ts" setup></script>

<template>
  <div class="flex">
    <div class="text-2xl breadcrumbs mb-8 grow">
      <ul>
        <li>Apps</li>
      </ul>
    </div>

    <button class="btn btn-outline hover:btn-primary gap-2" title="Create App">
      Create
      <i-mdi-plus class="w-6 h-6" />
    </button>
  </div>

  <div>
    <router-view />
  </div>
</template>
