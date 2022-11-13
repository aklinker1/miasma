<script lang="ts" setup>
const serviceId = useServiceIdPathParam();

const {
  data: service,
  error,
  refetch,
} = useDockerServiceQuery(computed(() => serviceId.value ?? ''));

useHead({
  title: service.value?.Spec?.Name ?? 'Unknown',
});
</script>

<template>
  <div class="space-y-4">
    <error-display title="Failed to load service details" :error="error" :retry="refetch" />

    <template v-if="service">
      <service-controller :service="service" />
      <service-details-form :service="service" />
    </template>
  </div>
</template>
