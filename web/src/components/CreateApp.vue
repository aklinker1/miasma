<script lang="ts" setup>
import { useRouter } from "vue-router";
import {
  AppInput,
  useCreateAppMutation,
} from "../composition/create-app-mutation";

const router = useRouter();

const modalBackground = ref<HTMLElement>();

const name = ref("");
const image = ref("");

const isSubmitDisabled = computed(
  () => !name.value.trim() || !image.value.trim() || loading.value
);

function dismiss() {
  modalBackground.value?.click();
}

const { loading, error, mutate: createApp } = useCreateAppMutation();

async function createAppWithValues() {
  const app: AppInput = {
    name: name.value.trim(),
    image: image.value.trim(),
  };
  const result = await createApp({ app });
  const newId = result?.data?.app.id;
  dismiss();
  if (!newId) {
    console.warn("Created app did not have an ID");
  }
  router.push(`/apps/${newId}`);
}
</script>

<template>
  <!-- The button to open modal -->
  <label
    for="create-app-modal"
    class="btn btn-outline hover:btn-primary gap-2 modal-button"
    title="Create App"
    >Create
    <i-mdi-plus class="w-6 h-6" />
  </label>

  <!-- Put this part before </body> tag -->
  <teleport to="body">
    <input type="checkbox" id="create-app-modal" class="modal-toggle" />
    <label
      ref="modalBackground"
      for="create-app-modal"
      class="modal cursor-pointer"
    >
      <label class="modal-box relative" for="">
        <form @submit.prevent="createAppWithValues" class="flex flex-col gap-4">
          <h3 class="text-lg font-bold">Create New App</h3>

          <!-- App Name -->
          <div class="form-control w-full group">
            <label class="input-group">
              <span
                ><app-icon class="group-focus-within:text-primary" :name="name"
              /></span>
              <input
                type="text"
                placeholder="App Name"
                class="input input-bordered focus:input-primary w-full"
                v-model="name"
              />
            </label>
          </div>

          <!-- Docker Image -->
          <div class="form-control w-full group">
            <label class="input-group">
              <span
                ><i-mdi-docker class="group-focus-within:text-primary"
              /></span>
              <input
                type="text"
                placeholder="Docker Image"
                class="input input-bordered focus:input-primary w-full"
                v-model="image"
              />
            </label>

            <!-- Error Label -->
            <label v-if="error" class="label">
              <span class="label-text-alt text-error">{{ error }}</span>
            </label>
          </div>

          <!-- Submit -->
          <button
            type="submit"
            class="btn btn-primary self-end"
            :class="{ loading }"
            :disabled="isSubmitDisabled"
          >
            Create
          </button>
        </form>
      </label>
    </label>
  </teleport>
</template>
