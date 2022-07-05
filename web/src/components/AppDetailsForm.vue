<script lang="ts" setup>
import { AppDetails } from "../composition/app-details-query";
import {
  AppChanges,
  useEditAppMutation,
} from "../composition/edit-app-mutation";

const props = defineProps<{
  app: AppDetails;
}>();

const name = ref(props.app.name);
const image = ref(props.app.image);
const group = ref(props.app.group ?? "");

function getAppChanges() {
  const c: AppChanges = {};

  const cleanName = name.value.trim();
  if (cleanName !== props.app.name) c.name = cleanName;

  const cleanImage = image.value.trim();
  if (cleanImage !== props.app.image) c.image = cleanImage;

  const cleanGroup = group.value.trim() || null;
  if (cleanGroup !== props.app.group) c.group = cleanGroup;

  return c;
}

function discardChanges() {
  name.value = props.app.name;
  image.value = props.app.image;
  group.value = props.app.group ?? "";
}

const changes = computed(getAppChanges);
const hasChanged = computed(() => Object.keys(changes.value).length > 0);

const { mutate: editApp, loading: isEditing } = useEditAppMutation();

function saveChanges() {
  editApp({ id: props.app.id, changes: changes.value });
}
</script>

<template>
  <form @submit.prevent="saveChanges" @reset.prevent="discardChanges">
    <!-- Header -->
    <app-metadata-input
      v-model:name="name"
      v-model:image="image"
      v-model:group="group"
      :app="app"
    />

    <!-- TODO: Environment Variables -->
    <!-- TODO: Ports -->
    <!-- TODO: Volumes -->
    <!-- TODO: Placement -->

    <!-- Save bar -->
    <save-changes-alert
      :is-saving="isEditing"
      :visible="hasChanged"
      type="submit"
      @discard="discardChanges"
    />
  </form>
</template>
