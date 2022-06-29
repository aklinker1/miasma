<script lang="ts" setup></script>

<template>
  <div class="flex">
    <h1 class="text-2xl mb-8 grow">Nodes</h1>

    <button
      class="btn btn-outline hover:btn-primary gap-2 disabled"
      title="Create App"
    >
      Add
      <i-mdi-plus class="w-6 h-6" />
    </button>
  </div>

  <node-grid />
</template>
