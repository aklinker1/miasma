<script lang="ts" setup>
import { ButtonHTMLAttributes } from 'vue';

defineProps<{
  isSaving: boolean;
  visible: boolean;
  error: H3Error<Docker.ErrorResponse> | null;
  type?: ButtonHTMLAttributes['type'];
}>();
</script>

<template>
  <div class="fixed inset-0 pointer-events-none">
    <div class="absolute p-16 bottom-0 right-0 z-50">
      <div
        class="alert shadow-lg pointer-events-auto transition"
        :class="{
          'opacity-100 translate-y-0': visible,
          'opacity-50 translate-y-48': !visible,
          'bg-primary': !error,
          'alert-error': error,
        }"
      >
        <div class="text-primary-content">
          <p v-if="!error">You have unsaved changes</p>
          <div v-else class="space-y-1">
            <p class="font-bold">Failed to save changes</p>
            <p class="text-sm max-w-xs">{{ error.data.message }}</p>
          </div>
        </div>
        <div class="flex-none">
          <button
            class="btn btn-primary"
            :class="{ 'opacity-0': isSaving, 'btn-error text-error-content': !!error }"
            :disabled="isSaving"
            type="reset"
          >
            Discard
          </button>
          <button type="submit" class="btn" :class="{ loading: isSaving }" :disabled="isSaving">
            Save
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
