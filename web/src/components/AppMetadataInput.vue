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
  <div class="space-y-2 p-2">
    <div class="form-control flex-1">
      <label class="label">
        <span class="label-text">App Name</span>
      </label>
      <input
        class="input input-lg input-bordered"
        v-model="internalName"
        placeholder="Enter a name..."
      />
    </div>

    <div class="form-control">
      <label class="label">
        <span class="label-text">Image</span>
      </label>
      <input
        class="input input-bordered w-full"
        v-model="internalImage"
        placeholder="Enter an image..."
      />
    </div>
    <div class="form-control">
      <label class="label">
        <span class="label-text">Group</span>
      </label>
      <input
        class="input input-bordered w-full"
        v-model="internalGroup"
        placeholder="Enter a group..."
      />
    </div>
  </div>
</template>
