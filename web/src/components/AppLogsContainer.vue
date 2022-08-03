<script lang="ts" setup>
defineProps<{
  appId: string;
}>();

const open = ref(false);

enum Tab {
  LOGS,
  TASKS,
}

const tab = ref(Tab.LOGS);
</script>

<template>
  <span for="app-logs-modal" @click="open = true">
    <i-mdi-console />
    View Logs & Tasks
  </span>

  <teleport to="body">
    <input
      type="checkbox"
      id="app-logs-modal"
      class="modal-toggle"
      v-model="open"
    />
    <label for="app-logs-modal" class="modal cursor-pointer">
      <label class="modal-box relative space-y-4 max-w-none w-fit" for="">
        <div class="tabs mx-auto">
          <a
            class="tab"
            :class="{
              'tab-active text-primary font-medium': tab === Tab.LOGS,
            }"
            @click="tab = Tab.LOGS"
            >Logs</a
          >
          <a
            class="tab"
            :class="{
              'tab-active text-primary font-medium': tab === Tab.TASKS,
            }"
            @click="tab = Tab.TASKS"
            >Tasks</a
          >
        </div>
        <div
          class="bg-base-300 rounded-lg overflow-scroll flex flex-col-reverse w-[640px] h-[480px]"
        >
          <transition-group>
            <log-list v-if="tab === Tab.LOGS" :app-id="appId" />
            <task-list v-else :app-id="appId" />
          </transition-group>
        </div>
      </label>
    </label>
  </teleport>
</template>
