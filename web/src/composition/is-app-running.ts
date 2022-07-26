import { Ref } from "vue";
import { RuntimeStatus } from "./app-status-query";

export function useIsAppRunning(
  app: Ref<{ status: RuntimeStatus } | undefined>
) {
  return computed(() => app.value?.status === "RUNNING");
}
