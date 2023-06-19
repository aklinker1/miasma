namespace Docker {
  type ContainerEnv = DeepRequired<Docker.TaskSpec>['ContainerSpec']['Env'];
}
