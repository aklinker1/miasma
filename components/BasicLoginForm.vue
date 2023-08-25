<script lang="ts" setup>
import { UseMutationReturnType } from 'vue-query';

const props = defineProps<{
  loginTest: Omit<UseMutationReturnType<any, unknown, void, any>, 'mutate' | 'mutateAsync'>;
}>();

const emit = defineEmits<{
  (event: 'submit', cookieValue: string): void;
}>();

const username = ref('');
const password = ref('');

const submitDisabled = computed(() => !username.value.trim() || !password.value.trim());
function submit() {
  if (submitDisabled.value) return;

  const base64 = `Basic ${btoa(username.value + ':' + password.value)}`;
  emit('submit', base64);
}
</script>

<template>
  <form class="flex flex-col gap-4" @submit.prevent.stop="submit">
    <h3>Login&ensp;<span class="opacity-50">&bull;</span>&ensp;Basic</h3>

    <label class="join group">
      <span class="join-item btn btn-square group-focus-within:text-primary">
        <span class="i-mdi-account text-2xl" />
      </span>
      <input
        class="join-item input input-bordered group-focus-within:input-primary w-full min-w-0"
        placeholder="Username"
        autocomplete="username"
        v-model="username"
      />
    </label>

    <label class="join group">
      <span class="join-item btn btn-square group-focus-within:text-primary">
        <span class="i-mdi-key text-2xl" />
      </span>
      <input
        class="join-item input input-bordered group-focus-within:input-primary w-full min-w-0"
        :class="{
          'input-error': loginTest.error.value,
        }"
        type="password"
        placeholder="Password"
        autocomplete="current-password"
        v-model="password"
      />
    </label>

    <p v-if="loginTest.error.value" class="text-error">Username or password is not correct</p>

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
