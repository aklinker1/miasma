<script lang="ts" setup>
import { useStartAppMutation } from "../composition/start-app-mutation";

defineProps<{ appId: string }>();

const { mutate: startApp, loading } = useStartAppMutation();
</script>

<template>
  <div
    class="btn btn-circle btn-outline hover:btn-success"
    :class="{
      'loading disabled': loading,
    }"
    title="Start"
    @click.prevent="startApp({ id: appId })"
  >
    <i-mdi-play v-if="!loading" class="w-5 h-5" />
  </div>
</template>
