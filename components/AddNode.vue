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

function showModal() {
  reset();
  modal.value?.showModal();
}

const { data: swarmInfo } = useDockerSwarmInfoQuery();
const { data: nodes } = useDockerNodesQuery();

const joinCommand = computed(() => {
  const token = swarmInfo.value?.JoinTokens?.Worker ?? '<token>';
  const ip =
    nodes.value?.find(node => node.ManagerStatus?.Reachability === 'reachable')?.ManagerStatus
      ?.Addr ?? '<manager-ip>';
  return `docker swarm join \\\n  --token ${token} \\\n  ${ip}`;
});
</script>

<template>
  <!-- The button to open modal -->
  <button
    class="btn btn-outline hover:btn-primary gap-2"
    title="Create Service"
    @click="showModal()"
  >
    <div class="i-mdi-plus text-2xl" />
    <span>Add Node</span>
  </button>

  <dialog ref="modal" class="modal">
    <!-- Main form -->
    <form method="dialog" class="modal-box space-y-4">
      <h3 class="font-bold text-lg">Add a Node</h3>

      <p>SSH into your machine and follow the steps below.</p>

      <ul class="steps steps-vertical">
        <li class="step">
          <a class="link link-primary" href="https://docs.docker.com/get-docker/" target="_blank"
            >Install Docker</a
          >
        </li>
        <li class="step">
          <p class="text-left">Join the Swarm</p>
        </li>
      </ul>

      <div class="p-4 bg-neutral text-neutral-content text-sm overflow-x-auto rounded-box">
        <pre>{{ joinCommand }}</pre>
      </div>

      <div class="modal-action">
        <button class="btn btn-primary">Done</button>
      </div>
    </form>

    <!-- Click to dismiss -->
    <form method="dialog" class="modal-backdrop">
      <button>close</button>
    </form>
  </dialog>
</template>
