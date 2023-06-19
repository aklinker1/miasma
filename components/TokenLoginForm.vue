<script lang="ts" setup>
import { UseMutationReturnType } from 'vue-query';

const props = defineProps<{
  loginTest: Omit<UseMutationReturnType<any, unknown, void, any>, 'mutate' | 'mutateAsync'>;
}>();

const emit = defineEmits<{
  (event: 'submit', cookieValue: string): void;
}>();

const token = ref('');

const submitDisabled = computed(() => !token.value);
function submit() {
  if (submitDisabled.value) return;

  emit('submit', `Bearer ${token.value}`);
}
</script>

<template>
  <form class="flex flex-col gap-4" @submit.prevent.stop="submit">
    <h3>Login&ensp;<span class="opacity-50">&bull;</span>&ensp;Token Required</h3>
    <label class="join group">
      <span class="join-item btn btn-square group-focus-within:text-primary">
        <span class="i-mdi-key text-2xl" />
      </span>
      <input
        class="join-item input input-bordered group-focus-within:input-primary w-full min-w-0"
        :class="{
          'input-error': loginTest.error.value,
        }"
        placeholder="Enter a token..."
        autocomplete="current-password"
        v-model="token"
      />
    </label>
    <p v-if="loginTest.error.value" class="text-error">Token is not correct</p>
    <button
      type="submit"
      class="btn btn-primary self-end"
      :disabled="submitDisabled || loginTest.isLoading.value"
    >
      <span v-if="loginTest.isLoading.value" class="loading loading-spinner" />
      Login
    </button>
  </form>
</template>
