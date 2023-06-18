export default function () {
  const route = useRoute();
  return computed(() => (route.params.pluginName ?? '') as string);
}
