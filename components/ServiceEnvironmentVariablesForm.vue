<script lang="ts" setup>
const props = defineProps<{
  env: Docker.ContainerEnv;
}>();
const emits = defineEmits<{
  (event: 'update:env', newEnv: Docker.ContainerEnv): void;
}>();

const {
  list,
  removeItem,
  updateItem: _updateItem,
} = useInputList({
  key: 'env',
  props,
  emits,
  emptyValue: '=',
  isEmpty: item => item.startsWith('='),
});

function updateKey(index: number, event: Event) {
  const newKey = (event.target as HTMLInputElement).value.trim();
  const [_, value] = list.value[index].split('=', 2);
  _updateItem(index, `${newKey}=${value}`);
}

function updateValue(index: number, event: Event) {
  const newValue = (event.target as HTMLInputElement).value;
  const [key] = list.value[index].split('=', 2);
  _updateItem(index, `${key}=${newValue}`);
}
</script>

<template>
  <div>
    <table class="table w-full table-compact shadow-2xl">
      <thead>
        <tr>
          <td class="min-w-[16rem]">Key</td>
          <td class="w-full">Value</td>
        </tr>
      </thead>
      <!-- Items -->
      <tbody>
        <tr v-for="(envVar, i) of list" :key="i">
          <td>
            <input
              type="text"
              placeholder="Key"
              class="input input-sm input-bordered focus:input-primary w-full min-w-[16rem] placeholder:opacity-50 font-mono"
              :value="envVar.split('=', 2)[0]"
              @input="event => updateKey(i, event)"
            />
          </td>
          <td>
            <div class="flex w-full gap-2">
              <input
                type="text"
                placeholder="Value"
                class="input input-sm input-bordered focus:input-primary min-w-[16rem] placeholder:opacity-50 font-mono flex-1"
                :value="envVar.split('=', 2)[1]"
                @input="event => updateValue(i, event)"
              />
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
