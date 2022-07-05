<script lang="ts" setup>
import { AppDetails } from "../composition/app-details-query";
import { useInternalValue } from "../composition/internal-value";

const props = defineProps<{
  name: string;
  image: string;
  group: string;
}>();
const emit = defineEmits<{
  (event: "update:name", newName: string): void;
  (event: "update:image", newImage: string): void;
  (event: "update:group", newGroup: string): void;
}>();

const internalName = useInternalValue<"name">("name", props, emit);
const internalImage = useInternalValue<"image">("image", props, emit);
const internalGroup = useInternalValue<"group">("group", props, emit);
</script>

<template>
  <div class="form-control">
    <label class="input-group">
      <span>Name</span>
      <input
        class="input input-lg input-bordered w-full"
        :class="{ 'input-error': !internalName.trim() }"
        v-model="internalName"
        placeholder="Enter a name..."
      />
    </label>
  </div>

  <div class="flex gap-4">
    <div class="form-control flex-grow">
      <label class="input-group">
        <span>Image</span>
        <input
          class="input input-bordered w-full"
          :class="{ 'input-error': !internalImage.trim() }"
          v-model="internalImage"
          placeholder="Enter an image..."
        />
      </label>
    </div>
    <div class="form-control flex-grow">
      <label class="input-group">
        <span>Group</span>
        <input
          class="input input-bordered w-full"
          v-model="internalGroup"
          placeholder="..."
        />
      </label>
    </div>
  </div>
</template>
