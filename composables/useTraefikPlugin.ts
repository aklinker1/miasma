export default function () {
  const {
    data: services,
    isLoading: isLodaingServices,
    refetch: refetchServices,
  } = useDockerServicesQuery({
    filters: {
      name: ['miasma-traefik-plugin'],
    },
  });
  const service = computed(() => services.value?.[0]);

  const { mutate: createService, isLoading: isCreatingService } = useDockerCreateServiceMutation();
  const enable = async () => {
    console.log('Enabling traefik plugin...');
    createService(buildTraefikService({}), {
      onSettled: () => refetchServices.value(),
    });
  };

  const { mutate: deleteService, isLoading: isDeletingService } = useDockerDeleteServiceMutation();
  const disable = async () => {
    console.log('Disabling traefik plugin...');
    if (service.value != null) {
      deleteService(service.value);
    }
  };

  const canToggle = computed(
    () => !isLodaingServices.value && !isCreatingService.value && !isDeletingService.value,
  );

  const isEnabled = computed({
    get() {
      return !!service.value;
    },
    set(enabled) {
      if (enabled) enable();
      else disable();
    },
  });

  return {
    canToggle,
    isEnabled,
  };
}
