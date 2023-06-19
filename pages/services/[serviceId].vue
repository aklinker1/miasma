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
  <div class="service-grid">
    <error-display
      title="Failed to load service details"
      :error="error"
      :retry="refetch"
      data-grid="error"
    />

    <template v-if="service">
      <service-controller :service="service" data-grid="side" />
      <service-details-form :service="service" data-grid="content" />
    </template>
  </div>
</template>

<style scoped>
.service-grid {
  display: grid;
  grid-template-columns: 1fr;
  grid-auto-flow: row;
  gap: 2rem;
}
*[data-grid='side'] {
  width: 100%;
  max-width: 28rem;
  justify-self: center;
}

@media (min-width: 768px) {
  .service-grid {
    grid-template-rows: auto 1fr;
    grid-template-columns: 14rem auto;
    grid-template-areas:
      'error error'
      'side content';
  }

  *[data-grid='error'] {
    grid-area: error;
  }
  *[data-grid='side'] {
    grid-area: side;
    align-self: flex-start;
    position: sticky;
    top: 2rem;
  }
  *[data-grid='content'] {
    grid-area: content;
  }
}
</style>
