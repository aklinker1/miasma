<script lang="ts" setup>
const props = defineProps<{
  service: Docker.Service;
}>();

const service = toRef(props, 'service');

const status = useServiceStatus(service);

const name = computed(() => service.value.Spec?.Name ?? 'Unknown');
const runningTasks = computed(() => service.value?.ServiceStatus?.RunningTasks ?? 0);
const desiredTasks = computed(() => service.value?.ServiceStatus?.DesiredTasks ?? 0);

const urls = useServiceUrls(service);
</script>

<template>
  <tr class="hover">
    <th>
      <service-icon class="w-6 h-6 ml-2" :name="name" />
    </th>
    <td>
      <nuxt-link :to="`/services/${service.ID}`">
        <p class="text-lg">
          {{ name }}
        </p>
        <p
          class="text-sm uppercase"
          :class="{
            'text-success': status === 'running',
            'text-error font-medium': status === 'stopped',
            'text-warning': status === 'degraded',
          }"
        >
          {{ status }}
        </p>
      </nuxt-link>
    </td>
    <td>
      <p v-if="status !== 'stopped'" class="text-lg text-center">
        <span>{{ runningTasks }}</span>
        /
        <span>{{ desiredTasks }}</span>
      </p>
    </td>
    <td class="space-x-3 text-right">
      <open-service-button v-if="status !== 'stopped' && urls?.length" :available-at="urls" />
      <stop-service-button v-if="status !== 'stopped'" :service="service" />
      <start-service-button v-else :service="service" />
    </td>
  </tr>
</template>
