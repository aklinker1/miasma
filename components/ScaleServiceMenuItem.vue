<script lang="ts" setup>
import { clamp } from '@vueuse/core';

const props = defineProps<{
  desiredTasks: number;
  runningTasks: number;
  service: Docker.Service;
}>();

const scaleTo = ref(props.desiredTasks);
function increment() {
  scaleTo.value = clamp(scaleTo.value + 1, 1, Infinity);
}
function decrement() {
  scaleTo.value = clamp(scaleTo.value - 1, 1, Infinity);
}

const { mutate: updateService, isLoading } = useDockerUpdateServiceMutation();
function scale() {
  if (disabled.value) return;

  const newSpec = props.service.Spec ? clone(props.service.Spec) : {};
  newSpec.Mode ??= {};
  newSpec.Mode.Replicated ??= {};
  newSpec.Mode.Replicated.Replicas = scaleTo.value;
  newSpec.Labels ??= {};
  newSpec.Labels[MiasmaLabels.InstanceCount] = String(scaleTo.value);

  updateService(
    { service: props.service, newSpec },
    {
      onSuccess: dismissModal,
    },
  );
}

const disabled = computed(() => !scaleTo.value);

const modal = ref<HTMLDialogElement>();
function showModal() {
  scaleTo.value = props.desiredTasks;
  modal.value?.showModal();
}
function dismissModal() {
  modal.value?.close();
}
</script>

<template>
  <li v-bind="$attrs" @click="showModal">
    <div>
      <div class="i-mdi-pencil text-2xl"></div>
      <span>{{ runningTasks ?? 0 }} / {{ desiredTasks }} instances</span>
    </div>
  </li>

  <dialog ref="modal" class="modal">
    <form method="dialog" class="modal-box w-80" @submit.prevent="scale">
      <h3 class="font-bold text-lg">Scale Horizontally</h3>

      <div class="flex justify-center">
        <div class="join mx-auto mt-4">
          <button
            type="button"
            class="join-item btn btn-lg btn-square btn-primary"
            @click="decrement"
          >
            <span class="i-mdi-minus text-2xl" />
          </button>
          <input
            type="number"
            class="join-item input input-lg input-bordered focus:input-primary text-center text-2xl min-w-0 w-full"
            min="1"
            v-model="scaleTo"
          />
          <button
            type="button"
            class="join-item btn btn-lg btn-square btn-primary"
            @click="increment"
          >
            <span class="i-mdi-plus text-2xl" />
          </button>
        </div>
      </div>

      <div class="modal-action">
        <!-- if there is a button in form, it will close the modal -->
        <button type="button" class="btn" @click="dismissModal">Cancel</button>
        <button type="button" class="btn btn-primary gap-2" :disabled="disabled" @click="scale">
          <span v-if="isLoading" class="loading loading-spinner" />
          Apply
        </button>
      </div>
    </form>

    <!-- Click to dismiss -->
    <form method="dialog" class="modal-backdrop">
      <button>close</button>
    </form>
  </dialog>
</template>

<style scoped>
/* Chrome, Safari, Edge, Opera */
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
  -webkit-appearance: none;
  margin: 0;
}

/* Firefox */
input[type='number'] {
  -moz-appearance: textfield;
}
</style>
