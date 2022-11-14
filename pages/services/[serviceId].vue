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

    <div v-if="service">
      <service-controller
        class="w-full md:w-56 md:flex-shrink-0 md:float-left md:mr-12"
        :service="service"
      />
      <service-details-form class="md:flex-grow" :service="service" />
    </div>
  </div>
</template>
