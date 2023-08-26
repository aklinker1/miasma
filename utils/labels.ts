export enum MiasmaLabels {
  /**
   * The service's group name.
   */
  Group = 'miasma.group',
  /**
   * How many instances should be desired. This label persists the value while the service is
   * stopped and scaled down to 0.
   */
  InstanceCount = 'miasma.instance-count',
  /**
   * Whether or not the service should be shown on the main services list.
   */
  Hidden = 'miasma.hidden',
  /**
   * Whether or not the service is a system app controlled by Miasma.
   */
  System = 'miasma.system',
}
