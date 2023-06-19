<script lang="ts" setup>
const props = defineProps<{
  mounts: Docker.Mount[];
}>();
const emits = defineEmits<{
  (event: 'update:mounts', newMounts: Docker.Mount[]): void;
}>();

const {
  list,
  removeItem,
  updateItem: _updateItem,
} = useInputList({
  key: 'mounts',
  props,
  emits,
  emptyValue: { Type: 'bind' } as Docker.Mount,
  isEmpty: item => !item.Target && !item.Source,
});

function updateTarget(index: number, event: Event) {
  const newTargetPort = (event.target as HTMLInputElement).value;
  _updateItem(index, {
    ...list.value[index],
    Target: newTargetPort,
  });
}

function updateSource(index: number, event: Event) {
  const newPublishedPort = (event.target as HTMLInputElement).value;
  _updateItem(index, {
    ...list.value[index],
    Source: newPublishedPort,
  });
}
</script>

<template>
  <div>
    <table class="table w-full table-compact shadow-2xl">
      <thead>
        <tr>
          <!-- <th class="w-8">Type</th> -->
          <th class="min-w-[12rem]">Target Path</th>
          <th />
          <th class="min-w-[12rem]">Source Path</th>
        </tr>
      </thead>
      <!-- Items -->
      <tbody>
        <tr v-for="(config, i) of list" :key="i">
          <!-- <td>
            <select
              class="select select-sm select-bordered focus:select-primary w-full min-w-[8rem] placeholder:opacity-50"
              :value="config.Type ?? 'bind'"
              disabled
              title="Not implemented yet"
            >
              <option value="bind">bind</option>
              <option value="volume">volume</option>
              <option value="tmpfs">tmpfs</option>
              <option value="npipe">npipe</option>
            </select>
          </td> -->
          <td>
            <div class="form-control">
              <label class="input-group input-group-lg">
                <span>
                  <div class="i-mdi-docker text-xl" />
                </span>
                <input
                  placeholder="Target"
                  class="input input-sm input-bordered focus:input-primary w-full placeholder:opacity-50 font-mono"
                  min="0"
                  max="65535"
                  :value="config.Target"
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
                  placeholder="Source"
                  class="input input-sm input-bordered focus:input-primary placeholder:opacity-50 font-mono flex-1"
                  min="0"
                  max="65535"
                  :value="config.Source"
                  autocomplete="off"
                  @input="event => updateSource(i, event)"
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
