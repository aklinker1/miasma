<script lang="ts" setup>
const props = defineProps<{
  service: Docker.Service;
}>();

const {
  mutate: _updateService,
  isLoading: isSaving,
  error: saveError,
  reset: resetSave,
} = useDockerUpdateServiceMutation();

const service = toRef(props, 'service');
const {
  latestModel,
  discardChanges,
  hasChanges,
  constraints,
  env,
  group,
  image,
  mounts,
  name,
  ports,
} = useDeepEditable(
  computed(() => service.value.Spec ?? {}),
  {
    name: model => model?.Name?.trim() ?? '',
    image: model => model?.TaskTemplate?.ContainerSpec?.Image,
    group: model => model?.Labels?.[MiasmaLabels.Group]?.trim() ?? '',
    env: model => model?.TaskTemplate?.ContainerSpec?.Env ?? [],
    ports: model => model?.EndpointSpec?.Ports ?? [],
    mounts: model => model?.TaskTemplate?.ContainerSpec?.Mounts ?? [],
    constraints: model => model?.TaskTemplate?.Placement?.Constraints ?? [],
  },
  (base, values): Docker.ServiceSpec => {
    base.Name = values.name.trim();

    base.TaskTemplate!.ContainerSpec!.Image = values.image.trim();

    const newGroup = values.group.trim();
    if (newGroup) base.Labels![MiasmaLabels.Group] = newGroup;
    else delete base.Labels![MiasmaLabels.Group];

    if (values.env.length) base.TaskTemplate!.ContainerSpec!.Env = values.env;
    else delete base.TaskTemplate!.ContainerSpec!.Env;

    if (values.ports.length) {
      base.EndpointSpec ??= {};
      base.EndpointSpec.Ports = values.ports;
    } else if (base?.EndpointSpec?.Ports) {
      delete base.EndpointSpec.Ports;
    }

    if (values.mounts.length) {
      base.TaskTemplate ??= {};
      base.TaskTemplate.ContainerSpec ??= {};
      base.TaskTemplate.ContainerSpec.Mounts = values.mounts;
    } else if (base?.TaskTemplate?.ContainerSpec?.Mounts) {
      delete base.TaskTemplate.ContainerSpec.Mounts;
    }

    if (values.constraints.length) {
      base.TaskTemplate ??= {};
      base.TaskTemplate.Placement ??= {};
      base.TaskTemplate.Placement.Constraints = values.constraints;
    } else if (base?.TaskTemplate?.Placement?.Constraints) {
      delete base.TaskTemplate.Placement.Constraints;
    }

    return base;
  },
  resetSave,
);

const currentName = computed(() => service.value.Spec?.Name ?? '');

function saveChanges() {
  _updateService({ service: service.value, newSpec: latestModel.value });
}
</script>

<template>
  <form @submit.prevent="saveChanges" @reset.prevent="discardChanges" class="space-y-4">
    <!-- Details -->
    <h1 class="text-2xl">Details</h1>
    <service-metadata-input
      :current-name="currentName"
      v-model:name="name"
      v-model:image="image"
      v-model:group="group"
    />

    <div class="divider" />

    <!-- Networks -->
    <h2 class="text-xl">Networking</h2>
    <service-ports-form v-model:ports="ports" />

    <div class="divider" />

    <!-- Config -->
    <h2 class="text-xl">Environment</h2>
    <service-environment-variables-form v-model:env="env" />

    <div class="divider" />

    <!-- Volumes -->
    <h2 class="text-xl">Mounts</h2>
    <service-mounts-form v-model:mounts="mounts" />

    <div class="divider" />

    <!-- Constraints -->
    <h2 class="text-xl">Constraints</h2>
    <service-constraints-form v-model:constraints="constraints" />

    <!-- TODO -->

    <!-- Save bar -->
    <save-changes-alert
      :is-saving="isSaving"
      :visible="hasChanges"
      :error="saveError"
      type="submit"
      @discard="discardChanges"
    />
  </form>
</template>
