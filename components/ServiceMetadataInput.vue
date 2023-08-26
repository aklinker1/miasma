<script lang="ts" setup>
const props = defineProps<{
  currentName: string;
  name: string;
  image: string;
  group: string;
  hidden: boolean;
}>();

const emit = defineEmits<{
  (event: 'update:name', newName: string): void;
  (event: 'update:image', newImage: string): void;
  (event: 'update:group', newGroup: string): void;
  (event: 'update:hidden', newHidden: boolean): void;
}>();

const internalName = useInternalValue<'name'>('name', props, emit);
const internalImage = useInternalValue<'image'>('image', props, emit);
const internalGroup = useInternalValue<'group'>('group', props, emit);
const internalHidden = useInternalValue<'hidden'>('hidden', props, emit);

const isShowingNameWarning = computed(() => internalName.value.trim() !== props.currentName);
</script>

<template>
  <!-- Service Name -->
  <div class="form-control">
    <label class="input-group group">
      <span class="group-focus-within:text-primary">Name</span>
      <input
        class="input input-lg input-bordered focus:input-primary w-full"
        :class="{ 'input-error': !internalName.trim() }"
        v-model="internalName"
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
          :class="{ 'input-error': !internalImage.trim() }"
          v-model="internalImage"
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
          v-model="internalGroup"
          placeholder=""
        />
      </label>
    </div>
  </div>

  <!-- Hidden -->
  <label class="flex py-2 items-center gap-4">
    <input type="checkbox" class="checkbox checked:checkbox-primary" v-model="internalHidden" />
    <span>Hidden</span>
  </label>
</template>
