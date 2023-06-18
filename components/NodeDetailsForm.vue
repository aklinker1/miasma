<script lang="ts" setup>
const props = defineProps<{
  node: Docker.Node;
}>();

const {
  mutate: _updateNode,
  isLoading: isSaving,
  error: saveError,
  reset: resetSave,
} = useDockerUpdateNodeMutation();

const node = toRef(props, 'node');
const {
  latestModel: latestNode,
  discardChanges,
  hasChanges,
  labels,
} = useDeepEditable(
  computed(() => node.value.Spec ?? {}),
  {
    labels: model => Object.entries(model.Labels ?? {}),
  },
  (base, values): Docker.NodeSpec => {
    base.Labels = Object.fromEntries(values.labels);
    return base;
  },
  resetSave,
);

function saveChanges() {
  _updateNode({ node: node.value, newSpec: latestNode.value });
}
</script>

<template>
  <form @submit.prevent="saveChanges" @reset.prevent="discardChanges" class="space-y-4">
    <!-- Details -->
    <h1 class="text-2xl">Details</h1>
    <p>
      Hostname: <code>{{ node?.Description?.Hostname }}</code>
    </p>
    <p>
      Manager? <code>{{ JSON.stringify(node?.ManagerStatus) }}</code>
    </p>
    <p>
      Platform: <code>{{ JSON.stringify(node?.Description?.Platform) }}</code>
    </p>

    <div class="divider" />

    <!-- Labels -->
    <h2 class="text-xl">Labels</h2>
    <node-labels-input v-model:labels="labels" />

    <!-- Save bar -->
    <save-changes-alert
      :is-saving="isSaving"
      :visible="hasChanges"
      :error="saveError"
      type="submit"
      @discard="discardChanges"
    />
  </form>
</template>
