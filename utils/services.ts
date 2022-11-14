import merge from 'lodash.merge';
import { MiasmaLabels } from './labels';

export function defineService(overrides: DeepPartial<Docker.ServiceSpec>): Docker.ServiceSpec {
  const defaultSpec: Docker.ServiceSpec = {
    Mode: {
      Replicated: {
        Replicas: 0,
      },
    },
  };
  const service = merge(defaultSpec, overrides);

  // Append or overwrite required fields
  service.Labels ??= {};
  service.Labels[MiasmaLabels.InstanceCount] = '1';

  service.Networks ??= [];
  // Communicate with other miasma services
  service.Networks.push({ Target: 'miasma-default' });

  return service;
}
