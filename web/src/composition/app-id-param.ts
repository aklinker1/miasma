import { useRoute } from "vue-router";

export function useAppIdParam() {
  const route = useRoute();
  return computed(() => route.params.appId as string | undefined);
}
