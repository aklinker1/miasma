<script lang="ts" setup>
import { useNodesQuery } from "../composition/nodes-query";

const { result } = useNodesQuery();
</script>

<template>
  <div class="flex gap-8 flex-wrap items-center justify-center">
    <node-grid-item
      v-for="node of result?.nodes"
      :key="node.id"
      :node="node"
      class="w-[20rem] shrink-0"
    />
  </div>
</template>
