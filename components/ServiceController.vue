<script lang="ts" setup>
import useDockerPullLatestMutation from '~/composables/useDockerPullLatestMutation';
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
const { mutate: pullLatest, isLoading: isPullingLatest } = useDockerPullLatestMutation();

const urls = useServiceUrls(service);

const isUpdating = computed(
  () =>
    isStarting.value ||
    isStopping.value ||
    // isRestarting.value ||
    isDeleting.value ||
    isPullingLatest.value,
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
    <li v-if="status" class="font-medium pointer-events-none">
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
      <div>
        <div class="i-mdi-pencil text-2xl"></div>
        <span>{{ service.ServiceStatus?.RunningTasks ?? 0 }} / {{ desiredTasks }} instances</span>
      </div>
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
    <li :class="{ 'disabled pointer-events-none': isUpdating }" @click="pullLatest(service)">
      <span>
        <div v-if="!isPullingLatest" class="i-mdi-cloud-download text-2xl" />
        <div v-else class="loading loading-spinner" />
        <span>Pull Latest Image</span>
      </span>
    </li>
    <li :class="{ 'disabled pointer-events-none': isUpdating }" @click="deleteService()">
      <span
        class="hover:text-error bg-error bg-opacity-0 hover:bg-opacity-10 active:bg-opacity-100 active:text-error-content"
      >
        <div v-if="!isDeleting" class="i-mdi-trash text-2xl" />
        <div v-else class="loading loading-spinner" />
        <span>Delete Service</span>
      </span>
    </li>
  </ul>
</template>
