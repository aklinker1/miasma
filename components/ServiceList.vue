<script lang="ts" setup>
const props = defineProps<{
  showHidden?: boolean;
}>();
const showHidden = computed(() => props.showHidden);

const includeStatus = ref(true);
const { data: services, error, refetch, isLoading } = useDockerServicesQuery(includeStatus);

const groups = useServiceGroups(services);
</script>

<template>
  <!-- Loading -->
  <div v-if="isLoading" class="flex py-32 items-center justify-center">
    <cube-spinner />
  </div>

  <!-- Error -->
  <error-display :error="error" :retry="refetch" />

  <!-- Services Table -->
  <div v-if="services" class="space-y-8">
    <div class="w-full shadow-2xl rounded-lg" v-for="group of groups" :key="group.name">
      <table class="table w-full">
        <thead>
          <tr>
            <th></th>
            <th class="w-full">{{ group.name || '' }}</th>
            <th class="text-center">Instances</th>
            <th></th>
          </tr>
        </thead>

        <!-- Main list -->
        <tbody>
          <service-list-item
            v-for="service of group.services"
            :key="service.ID"
            :service="service"
          />
        </tbody>
      </table>

      <!-- Empty message -->
      <p v-if="services.length === 0" class="text-center w-full p-6">No services</p>
    </div>
  </div>
</template>
