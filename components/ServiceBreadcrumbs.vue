<script lang="ts" setup>
const serviceId = useServiceIdPathParam();
const { data: service } = useDockerServiceQuery(serviceId);

const serviceName = computed(() => service.value?.Spec?.Name ?? 'Unknown');
</script>

<template>
  <div class="text-2xl breadcrumbs grow overflow-x-hidden">
    <ul v-if="!serviceId">
      <li>Services</li>
    </ul>
    <ul v-else>
      <li>
        <nuxt-link
          to="/services"
          class="link link-hover opacity-50 hover:opacity-100 active:opacity-100"
        >
          <span>Services</span>
        </nuxt-link>
      </li>
      <li class="space-x-2">
        <service-icon :name="serviceName" />
        <span>{{ serviceName }}</span>
      </li>
    </ul>
  </div>
</template>
