<script lang="ts" setup>
import { useRouter } from 'vue-router';

const router = useRouter();

const modalBackground = ref<HTMLElement>();

const name = ref('');
const image = ref('');
function reset() {
  name.value = '';
  image.value = '';
}

const isSubmitDisabled = computed(
  () => !name.value.trim() || !image.value.trim() || isLoading.value,
);

function dismiss() {
  modalBackground.value?.click();
}

const { isLoading, error, mutate: createService } = useDockerCreateServiceMutation();

async function createAppWithValues() {
  const spec = defineService({
    Name: name.value.trim(),
    TaskTemplate: {
      ContainerSpec: {
        Image: image.value.trim(),
      },
    },
  });
  createService(spec, {
    onSuccess(service) {
      dismiss();
      if (!service.ID) {
        console.warn('Created service did not have an ID');
      }
      router.push(routes.service(service.ID));
    },
  });
}
</script>

<template>
  <!-- The button to open modal -->
  <label
    for="create-service-modal"
    class="btn btn-outline hover:btn-primary gap-2 modal-button"
    title="Create Service"
    @click="reset"
  >
    <span>Create</span>
    <div class="i-mdi-plus text-2xl" />
  </label>

  <!-- Put this part before </body> tag -->
  <teleport to="body">
    <input type="checkbox" id="create-service-modal" class="modal-toggle" />
    <label ref="modalBackground" for="create-service-modal" class="modal cursor-pointer">
      <label class="modal-box relative" for="">
        <form @submit.prevent="createAppWithValues" class="flex flex-col gap-4">
          <h3 class="text-lg font-bold">Create New Service</h3>

          <!-- Service Name -->
          <div class="form-control w-full group">
            <label class="input-group">
              <span><service-icon class="group-focus-within:text-primary" :name="name" /></span>
              <input
                type="text"
                placeholder="Service Name"
                class="input input-bordered focus:input-primary w-full placeholder:opacity-50"
                v-model="name"
              />
            </label>
          </div>

          <!-- Docker Image -->
          <div class="form-control w-full group">
            <label class="input-group">
              <span class="px-3">
                <span class="i-mdi-docker text-2xl group-focus-within:text-primary" />
              </span>
              <input
                type="text"
                placeholder="Docker Image"
                class="input input-bordered focus:input-primary w-full placeholder:opacity-50"
                v-model="image"
              />
            </label>

            <!-- Error Label -->
            <label v-if="error" class="label">
              <span class="label-text-alt text-error">{{ error }}</span>
            </label>
          </div>

          <!-- Submit -->
          <button type="submit" class="btn btn-primary self-end" :disabled="isSubmitDisabled">
            <span v-if="isLoading" class="loading loading-spinner" />
            Create
          </button>
        </form>
      </label>
    </label>
  </teleport>
</template>
