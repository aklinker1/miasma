import { Ref } from 'vue';
import { MiasmaLabels } from '~~/utils/labels';

export interface AppGroup {
  name?: string;
  services: Docker.Service[];
}

export default function useGroupedApps(
  services: Ref<Docker.Service[] | undefined>,
): Ref<AppGroup[]> {
  return computed(() => {
    const groupMap = (services.value ?? []).reduce<Record<string, Docker.Service[]>>(
      (map, service) => {
        const groupName = service.Spec?.Labels?.[MiasmaLabels.Group]?.toUpperCase() ?? '';
        map[groupName] ??= [];
        map[groupName]?.push(service);
        return map;
      },
      {},
    );

    const groups: AppGroup[] = Object.entries(groupMap).map(([group, groupApps]) => ({
      name: group.trim() || undefined,
      services: groupApps,
    }));

    // Return a single empty group so we can show "no services" inside a table
    if (groups.length === 0) return [{ name: '', services: [] }];
    return groups;
  });
}
