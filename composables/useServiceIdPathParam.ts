export default function () {
  const route = useRoute();
  return computed<string>(() => (route.params.serviceId ?? '') as string);
}
