<script lang="ts" setup>
const props = defineProps<{
  ports: Docker.EndpointPortConfig[];
}>();
const emits = defineEmits<{
  (event: 'update:ports', newPorts: Docker.EndpointPortConfig[]): void;
}>();

const {
  list,
  removeItem,
  updateItem: _updateItem,
} = useInputList({
  key: 'ports',
  props,
  emits,
  emptyValue: {} as Docker.EndpointPortConfig,
  isEmpty: item => item.PublishedPort == null && item.TargetPort == null,
});

function updateTarget(index: number, event: Event) {
  const newTargetPort = (event.target as HTMLInputElement).value;
  _updateItem(index, {
    ...list.value[index],
    TargetPort: Number(newTargetPort),
  });
}

function updatePublished(index: number, event: Event) {
  const newPublishedPort = (event.target as HTMLInputElement).value;
  _updateItem(index, {
    ...list.value[index],
    PublishedPort: Number(newPublishedPort),
  });
}
</script>

<template>
  <div>
    <table class="table w-full table-compact shadow-2xl">
      <thead>
        <tr>
          <td class="w-8">Protocol</td>
          <td class="min-w-[12rem]">Target Port</td>
          <td />
          <td class="min-w-[12rem]">Published Port</td>
        </tr>
      </thead>
      <!-- Items -->
      <tbody>
        <tr v-for="(config, i) of list" :key="i">
          <td>
            <select
              class="select select-sm select-bordered focus:select-primary w-full min-w-[8rem] placeholder:opacity-50"
              :value="config.Protocol ?? 'tcp'"
              disabled
              title="Not implemented yet"
            >
              <option value="tcp">TCP</option>
            </select>
          </td>
          <td>
            <div class="form-control">
              <label class="input-group input-group-lg">
                <span>
                  <div class="i-mdi-docker text-xl" />
                </span>
                <input
                  type="number"
                  placeholder="Target"
                  class="input input-sm input-bordered focus:input-primary w-full placeholder:opacity-50 font-mono"
                  min="0"
                  max="65535"
                  :value="config.TargetPort"
                  autocomplete="off"
                  @input="event => updateTarget(i, event)"
                />
              </label>
            </div>
          </td>
          <td class="align-middle w-3">
            <div class="i-mdi-arrow-right text-xl m-1" />
          </td>
          <td>
            <div class="flex w-full gap-2">
              <label class="input-group input-group-lg">
                <span>
                  <div class="i-ri-server-fill text-xl" />
                </span>
                <input
                  type="number"
                  placeholder="Published"
                  class="input input-sm input-bordered focus:input-primary placeholder:opacity-50 font-mono flex-1"
                  min="0"
                  max="65535"
                  :value="config.PublishedPort"
                  autocomplete="off"
                  @input="event => updatePublished(i, event)"
                />
              </label>
              <div
                v-if="i !== list.length - 1"
                role="button"
                class="btn btn-ghost hover:btn-error btn-circle btn-sm"
                title="Delete variable"
                @click="removeItem(i)"
              >
                <div class="i-mdi-close text-xl" />
              </div>
            </div>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
