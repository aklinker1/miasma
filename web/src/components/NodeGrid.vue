<script lang="ts" setup>
import { useNodesQuery } from "../composition/nodes-query";

const { result } = useNodesQuery();
</script>

<template>
  <div class="grid grid-flow-row">
    <node-grid-item
      v-for="node of result?.nodes"
      :key="node.id"
      :node="node"
      class="max-w-[20rem]"
    />
  </div>
</template>
