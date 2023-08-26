/**
 * Returns a function that can be used to apply the current visibility filters to an array.
 *
 * @example
 * const { data: services } = useDockerServicesQuery();
 * const applyVisibilityFilter = useApplyVisibilityFilter();
 *
 * const visibleServices = computed(
 *   () => applyVisibilityFilter(services.value, service => service.Spec?.Labels),
 * )
 */
export function useApplyVisibilityFilter() {
  const showHidden = useShowHiddenServices();
  const showSystem = useShowSystemServices();

  return <T>(items: T[], getLabels: (t: T) => Record<string, string> | undefined): T[] => {
    return items.filter(item => {
      const labels = getLabels(item);
      if (!showHidden.value && labels?.[MiasmaLabels.Hidden] != null) return false;
      if (!showSystem.value && labels?.[MiasmaLabels.System] != null) return false;
      return true;
    });
  };
}
