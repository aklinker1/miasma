<script lang="ts" setup>
import { useStopAppMutation } from "../composition/stop-app-mutation";

defineProps<{ appId: string }>();

const { mutate: stopApp, loading } = useStopAppMutation();
</script>

<template>
  <div
    class="btn btn-circle btn-outline hover:btn-error"
    :class="{
      'loading disabled': loading,
    }"
    title="Stop"
    @click.prevent="stopApp({ id: appId })"
  >
    <i-mdi-stop v-if="!loading" class="w-5 h-5" />
  </div>
</template>
