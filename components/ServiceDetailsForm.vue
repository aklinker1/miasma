<script lang="ts" setup>
import { MiasmaLabels } from '~~/utils/labels';
import isEqual from 'lodash.isequal';
import { clone } from '~~/utils/clone';
import { detailedDiff } from 'deep-object-diff';

const props = defineProps<{
  service: Docker.Service;
}>();

const getPrevName = (): string => props.service.Spec?.Name?.trim() ?? '';
const getPrevImage = (): string => props.service.Spec!.TaskTemplate!.ContainerSpec!.Image!;
const getPrevGroup = (): string => props.service.Spec?.Labels?.[MiasmaLabels.Group]?.trim() ?? '';
const getPrevEnv = (): Docker.ContainerEnv => {
  const currentEnv = toRaw(props.service).Spec?.TaskTemplate?.ContainerSpec?.Env;
  return currentEnv ? [...currentEnv] : [];
};
const getPrevPortConfigs = (): Docker.EndpointPortConfig[] => {
  const currentPorts = toRaw(props.service).Spec?.EndpointSpec?.Ports;
  return currentPorts ? [...currentPorts] : [];
};
const getPrevMounts = (): Docker.Mount[] => {
  const currentMounts = toRaw(props.service).Spec?.TaskTemplate?.ContainerSpec?.Mounts;
  return currentMounts ? [...currentMounts] : [];
};

const currentName = computed(() => getPrevName());
const name = ref(getPrevName());
const image = ref(getPrevImage());
const group = ref(getPrevGroup());
const env = ref(getPrevEnv());
const ports = ref(getPrevPortConfigs());
const mounts = ref(getPrevMounts());

function getNewSpec(base: Docker.ServiceSpec = toRaw(props.service).Spec!) {
  const newSpec: Docker.ServiceSpec = clone(base);

  newSpec.Name = name.value.trim();

  newSpec.TaskTemplate!.ContainerSpec!.Image = image.value.trim();

  const newGroup = group.value.trim();
  if (newGroup) newSpec.Labels![MiasmaLabels.Group] = newGroup;
  else delete newSpec.Labels![MiasmaLabels.Group];

  if (env.value.length) newSpec.TaskTemplate!.ContainerSpec!.Env = env.value;
  else delete newSpec.TaskTemplate!.ContainerSpec!.Env;

  const validPorts = ports.value.filter(port => !!port.PublishedPort && !!port.TargetPort);
  if (validPorts.length) {
    newSpec.EndpointSpec ??= {};
    newSpec.EndpointSpec.Ports = validPorts;
  } else if (newSpec?.EndpointSpec?.Ports) {
    delete newSpec.EndpointSpec.Ports;
  }

  if (mounts.value.length) {
    newSpec.TaskTemplate ??= {};
    newSpec.TaskTemplate.ContainerSpec ??= {};
    newSpec.TaskTemplate.ContainerSpec.Mounts = mounts.value;
  } else if (newSpec?.TaskTemplate?.ContainerSpec?.Mounts) {
    delete newSpec.TaskTemplate.ContainerSpec.Mounts;
  }

  return newSpec;
}

function discardChanges() {
  name.value = getPrevName();
  image.value = getPrevImage();
  group.value = getPrevGroup();
  env.value = getPrevEnv();
  ports.value = getPrevPortConfigs();
  mounts.value = getPrevMounts();

  // If you discard on an error, don't show error next time there is a change
  resetSave.value();
}

const newSpec = computed(() => getNewSpec());
const hasChanges = computed(() => !isEqual(clone(props.service.Spec), clone(newSpec.value)));

watch(
  [newSpec, hasChanges],
  ([newNewSpec, newHasChanges]) => {
    if (newHasChanges)
      console.debug('Changes:', {
        diff: detailedDiff(toRaw(clone(props.service.Spec!)), toRaw(clone(newNewSpec))),
        new: clone(newNewSpec),
        old: clone(props.service.Spec),
      });
  },
  { immediate: true },
);

const {
  mutate: _updateService,
  isLoading: isSaving,
  error: saveError,
  reset: resetSave,
} = useDockerUpdateServiceMutation();
const router = useRouter();

function saveChanges() {
  _updateService({ service: props.service, newSpec: getNewSpec() });
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
    <p>TODO</p>

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
