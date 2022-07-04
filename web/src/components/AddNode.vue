<script lang="ts" setup>
import "prismjs";
import "prismjs/components/prism-bash";
import "url:https://raw.githubusercontent.com/PrismJS/prism-themes/master/themes/prism-atom-dark.css";
import Prism from "vue-prism-component";
import { useHealthQuery } from "../composition/health-query";

const open = ref(false);

const options = computed(() => ({ pollInterval: 2e3, enabled: open.value }));
const { result } = useHealthQuery(options);

const initCommand = "docker swarm init";

const code = computed(() => {
  const lines = [
    "# Install Docker",
    "curl -fsSL https://get.docker.com -o get-docker.sh",
    "sh get-docker.sh",
    "",
    "# Join the cluster",
    result.value?.health.cluster?.joinCommand,
  ];
  return lines.join("\n").trim();
});
</script>

<template>
  <!-- The button to open modal -->
  <label
    for="add-node-modal"
    class="btn btn-outline hover:btn-primary gap-2 disabled"
    title="Create App"
  >
    Add
    <i-mdi-plus class="w-6 h-6" />
  </label>

  <!-- Put this part before </body> tag -->
  <teleport to="body">
    <input
      type="checkbox"
      id="add-node-modal"
      class="modal-toggle"
      v-model="open"
    />
    <label
      ref="modalBackground"
      for="add-node-modal"
      class="modal cursor-pointer"
    >
      <label class="modal-box relative" for="">
        <div class="flex flex-col gap-4">
          <h3 class="text-lg font-bold">Add Node to Cluster</h3>

          <template v-if="result?.health == null">
            <div class="loading" />
          </template>

          <template v-else-if="result.health.cluster == null">
            <p>SSH into you're main node and initialize the swarm:</p>
            <prism language="bash">{{ initCommand }}</prism>
          </template>

          <template v-else>
            <p>
              To add a node, SSH into the machine and run the below commands.
            </p>

            <div>
              <prism language="bash">{{ code }}</prism>
              <label class="label">
                <span class="label-text-alt"
                  >Having problems? Checkout the
                  <a
                    class="link link-primary"
                    target="_blank"
                    href="https://aklinker1.github.io/miasma/guide/installation.html#add-more-nodes-optional"
                    >full documentation</a
                  ></span
                >
              </label>
            </div>
          </template>

          <!-- Close -->
          <label class="btn self-end" for="add-node-modal">Close</label>
        </div>
      </label>
    </label>
  </teleport>
</template>
