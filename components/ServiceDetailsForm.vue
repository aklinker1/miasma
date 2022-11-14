<script lang="ts" setup>
import { MiasmaLabels } from '~~/utils/labels';
import isEqual from 'lodash.isequal';
import { clone } from '~~/utils/clone';

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

const currentName = computed(() => getPrevName());
const name = ref(getPrevName());
const image = ref(getPrevImage());
const group = ref(getPrevGroup());
const env = ref(getPrevEnv());
const ports = ref(getPrevPortConfigs());

function getNewSpec(base: Docker.ServiceSpec = toRaw(props.service).Spec!) {
  const newSpec: Docker.ServiceSpec = clone(base);

  newSpec.Name = name.value.trim();

  newSpec.TaskTemplate!.ContainerSpec!.Image = image.value.trim();

  const newGroup = group.value.trim();
  if (newGroup) newSpec.Labels![MiasmaLabels.Group] = newGroup;
  else delete newSpec.Labels![MiasmaLabels.Group];

  if (env.value.length) newSpec.TaskTemplate!.ContainerSpec!.Env = env.value;
  else delete newSpec.TaskTemplate!.ContainerSpec!.Env;

  if (ports.value.length) {
    newSpec.EndpointSpec ??= {};
    newSpec.EndpointSpec.Ports = ports.value;
  } else if (newSpec?.EndpointSpec?.Ports) {
    delete newSpec.EndpointSpec.Ports;
  }

  return newSpec;
}

function discardChanges() {
  name.value = getPrevName();
  image.value = getPrevImage();
  group.value = getPrevGroup();
  env.value = getPrevEnv();

  // If you discard on an error, don't show error next time there is a change
  resetSave.value();
  resetRename.value();
}

const newSpec = computed(() => getNewSpec());
const hasChanged = computed(() => !isEqual(toRaw(newSpec.value), toRaw(props.service).Spec));

const {
  mutate: _updateService,
  isLoading: isSaving,
  error: saveError,
  reset: resetSave,
} = useDockerUpdateServiceMutation((_, prevSpec) => getNewSpec(prevSpec));
const {
  mutate: _renameService,
  isLoading: isRenaming,
  error: renameError,
  reset: resetRename,
} = useDockerRenameServiceMutation();
const router = useRouter();

function saveChanges() {
  if (name.value !== currentName.value) {
    _renameService(
      {
        prevService: toRaw(props.service),
        newSpec: getNewSpec(),
      },
      {
        onSuccess({ ID }) {
          router.replace(`/services/${ID}`);
        },
      },
    );
  } else {
    _updateService(props.service);
  }
}
</script>

<template>
  <form @submit.prevent="saveChanges" @reset.prevent="discardChanges" class="space-y-4">
    <!-- Details -->
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
    <h2 class="text-xl">Volumes</h2>
    <p>TODO</p>

    <div class="divider" />

    <!-- Constraints -->
    <h2 class="text-xl">Constraints</h2>
    <p>TODO</p>

    <!-- TODO -->

    <!-- Save bar -->
    <save-changes-alert
      :is-saving="isSaving || isRenaming"
      :visible="hasChanged"
      :error="saveError ?? renameError"
      type="submit"
      @discard="discardChanges"
    />
  </form>
</template>
