import { Ref } from "vue";

export function useIsAppRunning(app: Ref<{ status: string } | undefined>) {
  return computed(() => app.value?.status === "running");
}
