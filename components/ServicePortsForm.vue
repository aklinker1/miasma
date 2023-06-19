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
  emptyValue: {
    Protocol: 'tcp',
    PublishMode: 'ingress',
  } as Docker.EndpointPortConfig,
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
function updateProtocol(index: number, event: Event) {
  const newProtocol = (event.target as HTMLSelectElement).value as 'tcp' | 'udp' | 'sctp';
  _updateItem(index, {
    ...list.value[index],
    Protocol: newProtocol,
  });
}
function updatePublishMode(index: number, event: Event) {
  const newPublishMode = (event.target as HTMLSelectElement).value as 'ingress' | 'host';
  _updateItem(index, {
    ...list.value[index],
    PublishMode: newPublishMode,
  });
}
</script>

<template>
  <div>
    <table class="table w-full table-compact shadow-2xl">
      <thead>
        <tr>
          <!-- <th>Protocol</th> -->
          <!-- <th>Publish Mode</th> -->
          <th>Target Port</th>
          <th>Published Port</th>
        </tr>
      </thead>
      <!-- Items -->
      <tbody>
        <tr v-for="(config, i) of list" :key="i">
          <!-- <td class="whitespace-nowrap pr-0">
            <select
              class="select select-sm select-bordered focus:select-primary placeholder:opacity-50 min-w-0"
              :value="config.Protocol ?? 'tcp'"
              @input="event => updateProtocol(i, event)"
            >
              <option value="tcp">TCP</option>
              <option value="udp">UDP</option>
              <option value="sctp">SCTP</option>
            </select>
          </td> -->
          <!-- <td class="whitespace-nowrap pr-0">
            <select
              class="select select-sm select-bordered focus:select-primary placeholder:opacity-50 min-w-0"
              :value="config.PublishMode ?? 'ingress'"
              @input="event => updatePublishMode(i, event)"
            >
              <option value="ingress">Ingress</option>
              <option value="host">Host</option>
            </select>
          </td> -->
          <td class="w-full pr-0">
            <div class="flex w-full gap-2 items-center">
              <label class="input-group">
                <span>
                  <div class="i-mdi-docker text-xl" />
                </span>
                <input
                  type="number"
                  placeholder="Target"
                  class="input input-sm input-bordered focus:input-primary placeholder:opacity-50 font-mono flex-1 min-w-0"
                  min="0"
                  max="65535"
                  :value="config.TargetPort"
                  autocomplete="off"
                  @input="event => updateTarget(i, event)"
                />
              </label>
              <div class="shrink-0 i-mdi-arrow-right text-xl m-1" />
            </div>
          </td>
          <td class="w-full">
            <div class="flex w-full gap-2 items-center">
              <label class="input-group">
                <span>
                  <div class="i-ri-server-fill text-xl" />
                </span>
                <input
                  type="number"
                  placeholder="Published"
                  class="input input-sm input-bordered focus:input-primary placeholder:opacity-50 font-mono flex-1 min-w-0"
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
