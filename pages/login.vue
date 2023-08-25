<script lang="ts" setup>
import { useMutation } from 'vue-query';
import { useRoute, useRouter } from 'vue-router';

const route = useRoute();
const redirect = route.query['redirect'] as string | undefined;

const { data, isLoading } = useHealthQuery();

const docker = useDocker();
const test = useMutation(docker.getSystemInfo);

const router = useRouter();
const authHeader = useAuthHeader();
function login(cookieValue: string) {
  authHeader.value = cookieValue;
  test.mutate(undefined, {
    onSuccess() {
      router.push(redirect ?? routes.services);
    },
  });
}
</script>

<template>
  <div class="h-full max-w-screen-sm mx-auto flex">
    <div class="my-auto w-full shadow-2xl p-8 rounded-box flex flex-col">
      <!-- Loading auth type -->
      <div v-if="isLoading" class="mx-auto text-primary loading loading-spinner loading-lg" />

      <!-- Token -->
      <token-login-form v-else-if="data?.auth === 'token'" :login-test="test" @submit="login" />

      <!-- Basic -->
      <basic-login-form v-else-if="data?.auth === 'basic'" :login-test="test" @submit="login" />

      <!-- Unknown -->
      <p v-else class="text-error font-medium text-center">Unknown auth type: {{ data?.auth }}</p>
    </div>
  </div>
</template>
