<script lang="ts" setup>
import { MiasmaLabels } from '~~/utils/labels';

const props = defineProps<{
  service: Docker.Service;
}>();

const service = toRef(props, 'service');

const status = useServiceStatus(service);
const desiredTasks = computed(() => service.value.Spec?.Labels?.[MiasmaLabels.InstanceCount] ?? 1);

const { mutate: startService, isLoading: isStarting } = useDockerStartServiceMutation();
const { mutate: stopService, isLoading: isStopping } = useDockerStopServiceMutation();
// const { mutate: restartService, isLoading: isRestarting } = useRestartServiceMutation();
const { mutate: _deleteService, isLoading: isDeleting } = useDockerDeleteServiceMutation();
// const { mutate: editService, isLoading: isPullingLatest } = useDockerUpdateServiceMutation();
// function pullLatest() {
//   editService({ id: props.service.id, changes: { image: props.service.image } });
// }

const urls = useServiceUrls(service);

const isUpdating = computed(
  () =>
    isStarting.value ||
    isStopping.value ||
    // isRestarting.value ||
    isDeleting.value,
  // isPullingLatest.value,
);

const router = useRouter();
function deleteService() {
  _deleteService(service.value, { onSuccess: router.back });
}
</script>

<template>
  <ul class="menu bg-base-300 p-2 rounded-box shadow-2xl">
    <!-- Status -->
    <li class="menu-title"><span>Status</span></li>
    <li v-if="status" class="font-medium">
      <span
        class="bg-opacity-0 cursor-auto uppercase"
        :class="{
          'text-success': status === 'running',
          'text-error': status === 'stopped',
          'text-warning': status === 'degraded',
        }"
      >
        <div
          class="text-2xl"
          :class="{
            'i-mdi-warning': status === 'degraded',
            'i-mdi-check-circle': status !== 'degraded',
          }"
        />
        {{ status }}
      </span>
    </li>
    <li v-if="desiredTasks" :class="{ 'disabled pointer-events-none': isUpdating || true }">
      <span>
        <div class="i-mdi-pencil text-2xl" />
        <span>{{ service.ServiceStatus?.RunningTasks ?? 0 }}</span>
        /
        <span>{{ desiredTasks }} instances</span>
      </span>
    </li>
    <!-- <li>
      <service-logs-container :service-id="service.id" />
    </li> -->

    <!-- URLs -->
    <template v-if="urls?.length">
      <li class="menu-title"><span>URLs</span></li>
      <li v-for="url of urls" :key="url" class="text-ellipsis overflow-hidden">
        <a class="link hover:link-primary w-full" :href="url" target="_blank">{{
          url.replace(/https?:\/\//, '')
        }}</a>
      </li>
      <div class="divider-horizontal" />
    </template>

    <!-- Actions -->
    <li class="menu-title"><span>Actions</span></li>
    <li
      :class="{ 'disabled pointer-events-none': isUpdating || status !== 'stopped' }"
      @click="startService(service)"
    >
      <span
        class="hover:text-success bg-success bg-opacity-0 hover:bg-opacity-10 active:bg-opacity-100 active:text-success-content"
      >
        <div class="i-mdi-play text-2xl" />
        <span>Start Service</span>
      </span>
    </li>
    <li
      :class="{ 'disabled pointer-events-none': isUpdating || status === 'stopped' }"
      @click="stopService(service)"
    >
      <span
        class="hover:text-error bg-error bg-opacity-0 hover:bg-opacity-10 active:bg-opacity-100 active:text-error-content"
      >
        <div class="i-mdi-stop text-2xl" />
        <span>Stop Service</span>
      </span>
    </li>
    <!-- <li
      :class="{ 'disabled pointer-events-none': isUpdating || status !== 'running' }"
      @click="restartService(service)"
    >
      <span>
        <div class="i-mdi-restart text-2xl" />
        <span>Restart Service</span>
      </span>
    </li> -->
    <!-- <li :class="{ 'disabled pointer-events-none': isUpdating || true }" @click="pullLatest">
      <span>
        <div class="i-mdi-cloud-download text-2xl" />
        <span>Pull Latest Image</span>
      </span>
    </li> -->
    <li :class="{ 'disabled pointer-events-none': isUpdating }" @click="deleteService()">
      <span
        class="hover:text-error bg-error bg-opacity-0 hover:bg-opacity-10 active:bg-opacity-100 active:text-error-content"
      >
        <div class="i-mdi-trash text-2xl" />
        <span>Delete Service</span>
      </span>
    </li>
  </ul>
</template>