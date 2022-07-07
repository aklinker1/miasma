<script lang="ts" setup>
import { AppDetails } from "../composition/app-details-query";
import { useDeleteAppMutation } from "../composition/delete-app-mutation";
import { useEditAppMutation } from "../composition/edit-app-mutation";
import { useIsAppRunning } from "../composition/is-app-running";
import { useRestartAppMutation } from "../composition/restart-app-mutation";
import { useStartAppMutation } from "../composition/start-app-mutation";
import { useStopAppMutation } from "../composition/stop-app-mutation";
import { useAppStatusQuery } from "../composition/app-status-query";
import { useRouter } from "vue-router";

const props = defineProps<{
  app: AppDetails;
}>();

const status = useAppStatusQuery(computed(() => props.app.id));

const isRunning = useIsAppRunning(status);
const isStartingUp = computed(
  () =>
    status.value?.instances != null &&
    status.value.instances.total !== status.value.instances.running
);

const { mutate: startApp, loading: isStarting } = useStartAppMutation();
const { mutate: stopApp, loading: isStopping } = useStopAppMutation();
const { mutate: restartApp, loading: isRestarting } = useRestartAppMutation();
const {
  mutate: deleteApp,
  loading: isDeleting,
  onDone: onDeleted,
} = useDeleteAppMutation();
const { mutate: editApp, loading: isPullingLatest } = useEditAppMutation();
function pullLatest() {
  editApp({ id: props.app.id, changes: { image: props.app.image } });
}

const router = useRouter();
onDeleted(router.back);

const isUpdating = computed(
  () =>
    isStarting.value ||
    isStopping.value ||
    isRestarting.value ||
    isDeleting.value ||
    isPullingLatest.value
);
</script>

<template>
  <ul class="menu bg-base-300 p-2 rounded-box shadow-2xl">
    <!-- Status -->
    <li class="menu-title"><span>Status</span></li>
    <li v-if="status" class="font-medium">
      <span
        class="bg-opacity-0 cursor-auto uppercase"
        :class="{
          'text-success': isRunning,
          'text-error': !isRunning,
          'text-warning': isStartingUp,
        }"
      >
        <i-mdi-warning v-if="isStartingUp" />
        <i-mdi-check-circle v-else />
        {{ status.status }}
      </span>
    </li>
    <li
      v-if="status?.instances"
      :class="{ 'disabled pointer-events-none': isUpdating || true }"
    >
      <span
        ><i-mdi-pencil /> {{ status.instances.running }} /
        {{ status.instances.total }} instances</span
      >
    </li>

    <!-- URLs -->
    <template v-if="app.availableAt.length > 0">
      <li class="menu-title"><span>URLs</span></li>
      <li
        v-for="link of app.availableAt"
        :key="link"
        class="text-ellipsis overflow-hidden"
      >
        <a class="link hover:link-primary w-full" :href="link" target="_blank"
          ><i-mdi-open-in-new class="shrink-0" />
          {{ link.replace(/https?:\/\//, "") }}</a
        >
      </li>
      <div class="divider-horizontal" />
    </template>

    <!-- Actions -->
    <li class="menu-title"><span>Actions</span></li>
    <li
      :class="{ 'disabled pointer-events-none': isUpdating || isRunning }"
      @click="startApp(app)"
    >
      <span
        class="hover:text-success bg-success bg-opacity-0 hover:bg-opacity-10 active:bg-opacity-100 active:text-success-content"
        ><i-mdi-play />Start App</span
      >
    </li>
    <li
      :class="{ 'disabled pointer-events-none': isUpdating || !isRunning }"
      @click="stopApp(app)"
    >
      <span
        class="hover:text-error bg-error bg-opacity-0 hover:bg-opacity-10 active:bg-opacity-100 active:text-error-content"
        ><i-mdi-stop /> Stop App</span
      >
    </li>
    <li
      :class="{ 'disabled pointer-events-none': isUpdating || !isRunning }"
      @click="restartApp(app)"
    >
      <span><i-mdi-restart />Restart App</span>
    </li>
    <li
      :class="{ 'disabled pointer-events-none': isUpdating || true }"
      @click="pullLatest"
    >
      <span><i-mdi-cloud-download />Pull Latest Image</span>
    </li>
    <li
      :class="{ 'disabled pointer-events-none': isUpdating }"
      @click="deleteApp(app)"
    >
      <span
        class="hover:text-error bg-error bg-opacity-0 hover:bg-opacity-10 active:bg-opacity-100 active:text-error-content"
        ><i-mdi-trash />Delete App</span
      >
    </li>
  </ul>
</template>
