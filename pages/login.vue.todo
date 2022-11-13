<script lang="ts" setup>
import { useRoute, useRouter } from "vue-router";
import { useHealthQuery } from "../composition/health-query";
import { setAccessToken } from "../utils/apollo-client";

const route = useRoute();
const redirect = route.params["redirect"] as string | undefined;

const token = ref("");
const validToken = computed(() => token.value !== "");

const healthQuery = useHealthQuery();

const checkingToken = ref(false);
const error = ref<string>();

const router = useRouter();
async function login() {
  try {
    error.value = undefined;
    setAccessToken(token.value);
    checkingToken.value = true;
    await healthQuery.refetch();
    checkingToken.value = false;
    router.push(redirect ?? "/");
  } catch (err) {
    await new Promise((res) => setTimeout(res, 0.5e3));
    console.log(err);
    checkingToken.value = false;
    error.value = (err as Error).message;
  }
}
</script>

<template>
  <div class="flex p-16 h-screen fixed inset-0 overflow-x-auto">
    <form
      @submit.prevent="login"
      class="p-8 m-auto shadow-2xl flex flex-col gap-4 border-2 border-base-200 rounded-lg min-w-max"
    >
      <h1 class="text-xl text-primary">Login</h1>
      <div class="form-control w-full max-w-xs">
        <div class="input-group">
          <input
            type="password"
            placeholder="Enter access token"
            class="input input-bordered w-full max-w-xs float-right"
            v-model="token"
          />
          <button
            type="submit"
            :disabled="!validToken || checkingToken"
            class="btn btn-primary"
            :class="{
              loading: checkingToken,
            }"
          >
            Log In
          </button>
        </div>
        <label class="label" v-if="error">
          <span class="label-text-alt text-error">{{ error }}</span>
        </label>
      </div>
    </form>
  </div>
</template>
