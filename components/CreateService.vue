<script lang="ts" setup>
import { useRouter } from 'vue-router';

const router = useRouter();

const modal = ref<HTMLDialogElement>();

const name = ref('');
const image = ref('');
function reset() {
  name.value = '';
  image.value = '';
}

const isSubmitDisabled = computed(
  () => !name.value.trim() || !image.value.trim() || isLoading.value,
);

function showModal() {
  reset();
  modal.value?.showModal();
}

function dismissModal() {
  modal.value?.close();
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
      dismissModal();
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
  <button
    v-bind="$attrs"
    class="btn btn-outline hover:btn-primary gap-2"
    title="Create Service"
    @click="showModal()"
  >
    <div class="i-mdi-plus text-2xl" />
    <span>Create Service</span>
  </button>

  <dialog ref="modal" class="modal">
    <!-- Main form -->
    <form method="dialog" class="modal-box space-y-4" @submit.prevent.stop="createAppWithValues">
      <h3 class="font-bold text-lg">Create New Service</h3>

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

      <!-- Docker image -->
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
          <span class="label-text-alt text-error">{{ error.data.message ?? error.message }}</span>
        </label>
      </div>

      <div class="modal-action">
        <button
          type="button"
          class="btn"
          :class="{ disabled: isLoading }"
          @click.stop.prevent="dismissModal"
        >
          <span v-if="isLoading" class="loading loading-spinner" />
          Cancel
        </button>
        <button type="submit" class="btn btn-primary" :disabled="isSubmitDisabled">
          <span v-if="isLoading" class="loading loading-spinner" />
          Create
        </button>
      </div>
    </form>

    <!-- Click to dismiss -->
    <form method="dialog" class="modal-backdrop">
      <button>close</button>
    </form>
  </dialog>
</template>
