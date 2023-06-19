import { Ref } from 'vue';

export interface AppGroup {
  name?: string;
  services: Docker.Service[];
}

export default function (services: Ref<Docker.Service[] | undefined>): Ref<AppGroup[]> {
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

    const groups: AppGroup[] = Object.entries(groupMap).map(([group, groupServices]) => ({
      name: group.trim() || undefined,
      services: groupServices,
    }));

    // Return a single empty group so we can show "no services" inside a table
    if (groups.length === 0) return [{ name: '', services: [] }];
    return groups;
  });
}
