<script lang="ts" setup>
import { ButtonHTMLAttributes } from "vue";

defineProps<{
  isSaving: boolean;
  visible: boolean;
  type?: ButtonHTMLAttributes["type"];
}>();
</script>

<template>
  <div class="fixed inset-0 pointer-events-none">
    <div class="absolute p-16 bottom-0 right-0">
      <div
        class="alert bg-primary shadow-lg pointer-events-auto transition"
        :class="{
          'opacity-100 translate-y-0': visible,
          'opacity-50 translate-y-48': !visible,
        }"
      >
        <div class="text-primary-content">
          <span>You have unsaved changes</span>
        </div>
        <div class="flex-none">
          <button
            class="btn btn-primary text-primary-content"
            :class="{
              'opacity-0': isSaving,
            }"
            :disabled="isSaving"
            type="reset"
          >
            Discard
          </button>
          <button
            type="submit"
            class="btn btn"
            :class="{ loading: isSaving }"
            :disabled="isSaving"
          >
            Save
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
