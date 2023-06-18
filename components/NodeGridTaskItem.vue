<script lang="ts" setup>
import { routes } from '~/utils/routes';

const props = defineProps<{
  task: Docker.Task;
}>();

const { data: service } = useDockerServiceQuery(computed(() => props.task.ServiceID ?? ''));
</script>

<template>
  <li>
    <nuxt-link :to="routes.service(task.ServiceID)">
      <service-icon v-if="service?.Spec?.Name" :name="service.Spec.Name" />
      {{ service?.Spec?.Name ?? 'Unknown' }}
    </nuxt-link>
  </li>
</template>
