<script lang="ts" setup>
const props = defineProps<{
  currentName: string;
}>();

const name = defineModel<string>('name', { required: true });
const image = defineModel<string>('image', { required: true });
const group = defineModel<string>('group', { required: true });
const hidden = defineModel<boolean>('hidden', { required: true });
const traefikRule = defineModel<string>('traefikRule', { required: true });

const isShowingNameWarning = computed(() => name.value.trim() !== props.currentName);

const { isEnabled: isTraefikEnabled } = useTraefikPlugin();
</script>

<template>
  <!-- Service Name -->
  <div class="form-control">
    <label class="input-group group">
      <span class="group-focus-within:text-primary">Name</span>
      <input
        class="input input-lg input-bordered focus:input-primary w-full"
        :class="{ 'input-error': !name.trim() }"
        v-model="name"
        placeholder="Enter a name..."
      />
    </label>
    <label v-if="isShowingNameWarning" class="label">
      <span class="label-text text-warning">
        Changing an service name will delete and create a new Docker service with a different ID
      </span>
    </label>
  </div>

  <div class="flex gap-4">
    <!-- Docker Image -->
    <div class="form-control flex-grow">
      <label class="input-group group">
        <span class="group-focus-within:text-primary">Image</span>
        <input
          class="input input-bordered focus:input-primary w-full"
          :class="{ 'input-error': !image.trim() }"
          v-model="image"
          placeholder="Enter an image..."
        />
      </label>
    </div>

    <!-- Group -->
    <div class="form-control flex-grow">
      <label class="input-group group" title="Group related services together on the dashboard">
        <span class="group-focus-within:text-primary">Group</span>
        <input
          class="input input-bordered focus:input-primary w-full"
          v-model="group"
          placeholder=""
        />
      </label>
    </div>
  </div>

  <!-- Traefik -->
  <label
    v-if="isTraefikEnabled"
    class="input-group group"
    title="Group related services together on the dashboard"
    style="--p: var(--traefik)"
  >
    <span class="group-focus-within:text-primary"><div class="i-devicon-traefikproxy" /></span>
    <input
      class="input input-bordered focus:input-primary w-full"
      v-model="traefikRule"
      placeholder="Enter a domain or Traefik rule"
    />
  </label>

  <!-- Hidden -->
  <label class="flex py-2 items-center gap-4">
    <input type="checkbox" class="checkbox checked:checkbox-primary" v-model="hidden" />
    <span>Hidden</span>
  </label>
</template>
