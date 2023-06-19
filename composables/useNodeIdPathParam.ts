export default function () {
  const route = useRoute();
  return computed(() => (route.params.nodeId ?? '') as string);
}
