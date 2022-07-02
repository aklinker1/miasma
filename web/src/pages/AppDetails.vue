<script lang="ts" setup>
import { useRoute } from "vue-router";

const route = useRoute();

const appId = computed(() => route.params["appId"] as string);
</script>

<template>
  <p>{{ appId }}</p>
</template>
