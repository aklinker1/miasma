<script lang="ts" setup>
const props = defineProps<{
  constraints: string[];
}>();
const emits = defineEmits<{
  (event: 'update:constraints', newConstraints: string[]): void;
}>();

const regex = /^(.*)(==|!=)(.*)$/;

const {
  list,
  removeItem,
  updateItem: _updateItem,
} = useInputList({
  key: 'constraints',
  props,
  emits,
  emptyValue: '==',
  isEmpty: item => item.startsWith('==') || item.startsWith('!='),
});

function updateKey(index: number, event: Event) {
  const newKey = (event.target as HTMLInputElement).value.trim();
  const [_, key, comparison, value] = list.value[index].match(regex) ?? [];
  _updateItem(index, `${newKey}${comparison}${value}`);
}

function updateComparison(index: number, event: Event) {
  const newComparison = (event.target as HTMLSelectElement).value;
  const [_, key, comparison, value] = list.value[index].match(regex) ?? [];
  console.log({ key, comparison, newComparison, value });
  _updateItem(index, `${key}${newComparison}${value}`);
}

function updateValue(index: number, event: Event) {
  const newValue = (event.target as HTMLInputElement).value;
  const [_, key, comparison, value] = list.value[index].match(regex) ?? [];
  _updateItem(index, `${key}${comparison}${newValue}`);
}
</script>

<template>
  <div>
    <table class="table w-full table-compact shadow-2xl">
      <thead>
        <tr>
          <th>Key</th>
          <th></th>
          <th class="w-full">Value</th>
        </tr>
      </thead>
      <!-- Items -->
      <tbody>
        <tr v-for="(constraint, i) of list" :key="i">
          <td>
            <input
              type="text"
              placeholder="Key"
              class="input input-sm input-bordered focus:input-primary w-full min-w-[16rem] placeholder:opacity-50 font-mono"
              :value="constraint.match(regex)![1]"
              @input="event => updateKey(i, event)"
            />
          </td>
          <td>
            <select
              class="select select-sm select-bordered"
              :value="constraint.match(regex)![2]"
              @input="event => updateComparison(i, event)"
              :disabled="!constraint.match(regex)![1]"
            >
              <option value="==">==</option>
              <option value="!=">!=</option>
            </select>
          </td>
          <td>
            <div class="flex w-full gap-2">
              <input
                type="text"
                placeholder="Value"
                class="input input-sm input-bordered focus:input-primary min-w-[16rem] placeholder:opacity-50 font-mono flex-1"
                :value="constraint.match(regex)![3]"
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
