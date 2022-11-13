namespace Docker {
  type Port = {
    /**
     * Host IP address that the container's port is mapped to
     */
    'IP'?: string;
    /**
     * Port on the container
     */
    'PrivatePort'?: number;
    /**
     * Port exposed on the host
     */
    'PublicPort'?: number;
    'Type'?: "tcp" | "udp" | "sctp";
  };

  type MountPoint = {
    /**
     * The mount type:
     * 
     * - `bind` a mount of a file or directory from the host into the container.
     * - `volume` a docker volume with the given `Name`.
     * - `tmpfs` a `tmpfs`.
     * - `npipe` a named pipe from the host into the container.
     */
    'Type'?: "bind" | "volume" | "tmpfs" | "npipe";
    /**
     * Name is the name reference to the underlying data defined by `Source`
     * e.g., the volume name.
     */
    'Name'?: string;
    /**
     * Source location of the mount.
     * 
     * For volumes, this contains the storage location of the volume (within
     * `/var/lib/docker/volumes/`). For bind-mounts, and `npipe`, this contains
     * the source (host) part of the bind-mount. For `tmpfs` mount points, this
     * field is empty.
     */
    'Source'?: string;
    /**
     * Destination is the path relative to the container root (`/`) where
     * the `Source` is mounted inside the container.
     */
    'Destination'?: string;
    /**
     * Driver is the volume driver used to create the volume (if it is a volume).
     */
    'Driver'?: string;
    /**
     * Mode is a comma separated list of options supplied by the user when
     * creating the bind/volume mount.
     * 
     * The default is platform-specific (`"z"` on Linux, empty on Windows).
     */
    'Mode'?: string;
    /**
     * Whether the mount is mounted writable (read-write).
     */
    'RW'?: boolean;
    /**
     * Propagation describes how mounts are propagated from the host into the
     * mount point, and vice-versa. Refer to the [Linux kernel documentation](https://www.kernel.org/doc/Documentation/filesystems/sharedsubtree.txt)
     * for details. This field is not used on Windows.
     */
    'Propagation'?: string;
  };

  type DeviceMapping = {
    'PathOnHost'?: string;
    'PathInContainer'?: string;
    'CgroupPermissions'?: string;
  };

  type DeviceRequest = {
    'Driver'?: string;
    'Count'?: number;
    'DeviceIDs'?: Array<string>;
    /**
     * A list of capabilities; an OR list of AND lists of capabilities.
     */
    'Capabilities'?: Array<Array<string>>;
    /**
     * Driver-specific options, specified as a key/value pairs. These options
     * are passed directly to the driver.
     */
    'Options'?: Record<string, string>;
  };

  type ThrottleDevice = {
    /**
     * Device path
     */
    'Path'?: string;
    /**
     * Rate
     */
    'Rate'?: number;
  };

  type Mount = {
    /**
     * Container path.
     */
    'Target'?: string;
    /**
     * Mount source (e.g. a volume name, a host path).
     */
    'Source'?: string;
    /**
     * The mount type. Available types:
     * 
     * - `bind` Mounts a file or directory from the host into the container. Must exist prior to creating the container.
     * - `volume` Creates a volume with the given name and options (or uses a pre-existing volume with the same name and options). These are **not** removed when the container is removed.
     * - `tmpfs` Create a tmpfs with the given options. The mount source cannot be specified for tmpfs.
     * - `npipe` Mounts a named pipe from the host into the container. Must exist prior to creating the container.
     */
    'Type'?: "bind" | "volume" | "tmpfs" | "npipe";
    /**
     * Whether the mount should be read-only.
     */
    'ReadOnly'?: boolean;
    /**
     * The consistency requirement for the mount: `default`, `consistent`, `cached`, or `delegated`.
     */
    'Consistency'?: string;
    /**
     * Optional configuration for the `bind` type.
     */
    'BindOptions'?: {
      /**
       * A propagation mode with the value `[r]private`, `[r]shared`, or `[r]slave`.
       */
      'Propagation'?: "private" | "rprivate" | "shared" | "rshared" | "slave" | "rslave";
      /**
       * Disable recursive bind mount.
       */
      'NonRecursive'?: boolean;
    };
    /**
     * Optional configuration for the `volume` type.
     */
    'VolumeOptions'?: {
      /**
       * Populate volume with data from the target.
       */
      'NoCopy'?: boolean;
      /**
       * User-defined key/value metadata.
       */
      'Labels'?: Record<string, string>;
      /**
       * Map of driver specific options
       */
      'DriverConfig'?: {
        /**
         * Name of the driver to use to create the volume.
         */
        'Name'?: string;
        /**
         * key/value map of driver specific options.
         */
        'Options'?: Record<string, string>;
      };
    };
    /**
     * Optional configuration for the `tmpfs` type.
     */
    'TmpfsOptions'?: {
      /**
       * The size for the tmpfs mount in bytes.
       */
      'SizeBytes'?: number;
      /**
       * The permission mode for the tmpfs mount in an integer.
       */
      'Mode'?: number;
    };
  };

  type RestartPolicy = {
    /**
     * - Empty string means not to restart
     * - `no` Do not automatically restart
     * - `always` Always restart
     * - `unless-stopped` Restart always except when the user has manually stopped the container
     * - `on-failure` Restart only when the container exit code is non-zero
     */
    'Name'?: "" | "no" | "always" | "unless-stopped" | "on-failure";
    /**
     * If `on-failure` is used, the number of times to retry before giving up.
     */
    'MaximumRetryCount'?: number;
  };

  type Resources = {
    /**
     * An integer value representing this container's relative CPU weight
     * versus other containers.
     */
    'CpuShares'?: number;
    /**
     * Memory limit in bytes.
     */
    'Memory'?: number;
    /**
     * Path to `cgroups` under which the container's `cgroup` is created. If
     * the path is not absolute, the path is considered to be relative to the
     * `cgroups` path of the init process. Cgroups are created if they do not
     * already exist.
     */
    'CgroupParent'?: string;
    /**
     * Block IO weight (relative weight).
     */
    'BlkioWeight'?: number;
    /**
     * Block IO weight (relative device weight) in the form:
     * 
     * ```
     * [{"Path": "device_path", "Weight": weight}]
     * ```
     */
    'BlkioWeightDevice'?: Array<{
      'Path'?: string;
      'Weight'?: number;
    }>;
    /**
     * Limit read rate (bytes per second) from a device, in the form:
     * 
     * ```
     * [{"Path": "device_path", "Rate": rate}]
     * ```
     */
    'BlkioDeviceReadBps'?: Array<ThrottleDevice>;
    /**
     * Limit write rate (bytes per second) to a device, in the form:
     * 
     * ```
     * [{"Path": "device_path", "Rate": rate}]
     * ```
     */
    'BlkioDeviceWriteBps'?: Array<ThrottleDevice>;
    /**
     * Limit read rate (IO per second) from a device, in the form:
     * 
     * ```
     * [{"Path": "device_path", "Rate": rate}]
     * ```
     */
    'BlkioDeviceReadIOps'?: Array<ThrottleDevice>;
    /**
     * Limit write rate (IO per second) to a device, in the form:
     * 
     * ```
     * [{"Path": "device_path", "Rate": rate}]
     * ```
     */
    'BlkioDeviceWriteIOps'?: Array<ThrottleDevice>;
    /**
     * The length of a CPU period in microseconds.
     */
    'CpuPeriod'?: number;
    /**
     * Microseconds of CPU time that the container can get in a CPU period.
     */
    'CpuQuota'?: number;
    /**
     * The length of a CPU real-time period in microseconds. Set to 0 to
     * allocate no time allocated to real-time tasks.
     */
    'CpuRealtimePeriod'?: number;
    /**
     * The length of a CPU real-time runtime in microseconds. Set to 0 to
     * allocate no time allocated to real-time tasks.
     */
    'CpuRealtimeRuntime'?: number;
    /**
     * CPUs in which to allow execution (e.g., `0-3`, `0,1`).
     */
    'CpusetCpus'?: string;
    /**
     * Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only
     * effective on NUMA systems.
     */
    'CpusetMems'?: string;
    /**
     * A list of devices to add to the container.
     */
    'Devices'?: Array<DeviceMapping>;
    /**
     * a list of cgroup rules to apply to the container
     */
    'DeviceCgroupRules'?: Array<string>;
    /**
     * A list of requests for devices to be sent to device drivers.
     */
    'DeviceRequests'?: Array<DeviceRequest>;
    /**
     * Kernel memory limit in bytes.
     * 
     * <p><br /></p>
     * 
     * > **Deprecated**: This field is deprecated as the kernel 5.4 deprecated
     * > `kmem.limit_in_bytes`.
     */
    'KernelMemory'?: number;
    /**
     * Hard limit for kernel TCP buffer memory (in bytes).
     */
    'KernelMemoryTCP'?: number;
    /**
     * Memory soft limit in bytes.
     */
    'MemoryReservation'?: number;
    /**
     * Total memory limit (memory + swap). Set as `-1` to enable unlimited
     * swap.
     */
    'MemorySwap'?: number;
    /**
     * Tune a container's memory swappiness behavior. Accepts an integer
     * between 0 and 100.
     */
    'MemorySwappiness'?: number;
    /**
     * CPU quota in units of 10<sup>-9</sup> CPUs.
     */
    'NanoCpus'?: number;
    /**
     * Disable OOM Killer for the container.
     */
    'OomKillDisable'?: boolean;
    /**
     * Run an init inside the container that forwards signals and reaps
     * processes. This field is omitted if empty, and the default (as
     * configured on the daemon) is used.
     */
    'Init'?: boolean;
    /**
     * Tune a container's PIDs limit. Set `0` or `-1` for unlimited, or `null`
     * to not change.
     */
    'PidsLimit'?: number;
    /**
     * A list of resource limits to set in the container. For example:
     * 
     * ```
     * {"Name": "nofile", "Soft": 1024, "Hard": 2048}
     * ```
     */
    'Ulimits'?: Array<{
      /**
       * Name of ulimit
       */
      'Name'?: string;
      /**
       * Soft limit
       */
      'Soft'?: number;
      /**
       * Hard limit
       */
      'Hard'?: number;
    }>;
    /**
     * The number of usable CPUs (Windows only).
     * 
     * On Windows Server containers, the processor resource controls are
     * mutually exclusive. The order of precedence is `CPUCount` first, then
     * `CPUShares`, and `CPUPercent` last.
     */
    'CpuCount'?: number;
    /**
     * The usable percentage of the available CPUs (Windows only).
     * 
     * On Windows Server containers, the processor resource controls are
     * mutually exclusive. The order of precedence is `CPUCount` first, then
     * `CPUShares`, and `CPUPercent` last.
     */
    'CpuPercent'?: number;
    /**
     * Maximum IOps for the container system drive (Windows only)
     */
    'IOMaximumIOps'?: number;
    /**
     * Maximum IO in bytes per second for the container system drive
     * (Windows only).
     */
    'IOMaximumBandwidth'?: number;
  };

  type Limit = {
    'NanoCPUs'?: number;
    'MemoryBytes'?: number;
    /**
     * Limits the maximum number of PIDs in the container. Set `0` for unlimited.
     */
    'Pids'?: number;
  };

  type ResourceObject = {
    'NanoCPUs'?: number;
    'MemoryBytes'?: number;
    'GenericResources'?: GenericResources;
  };

  type GenericResources = Array<{
    'NamedResourceSpec'?: {
      'Kind'?: string;
      'Value'?: string;
    };
    'DiscreteResourceSpec'?: {
      'Kind'?: string;
      'Value'?: number;
    };
  }>;

  type HealthConfig = {
    /**
     * The test to perform. Possible values are:
     * 
     * - `[]` inherit healthcheck from image or parent image
     * - `["NONE"]` disable healthcheck
     * - `["CMD", args...]` exec arguments directly
     * - `["CMD-SHELL", command]` run command with system's default shell
     */
    'Test'?: Array<string>;
    /**
     * The time to wait between checks in nanoseconds. It should be 0 or at
     * least 1000000 (1 ms). 0 means inherit.
     */
    'Interval'?: number;
    /**
     * The time to wait before considering the check to have hung. It should
     * be 0 or at least 1000000 (1 ms). 0 means inherit.
     */
    'Timeout'?: number;
    /**
     * The number of consecutive failures needed to consider a container as
     * unhealthy. 0 means inherit.
     */
    'Retries'?: number;
    /**
     * Start period for the container to initialize before starting
     * health-retries countdown in nanoseconds. It should be 0 or at least
     * 1000000 (1 ms). 0 means inherit.
     */
    'StartPeriod'?: number;
  };

  type Health = {
    /**
     * Status is one of `none`, `starting`, `healthy` or `unhealthy`
     * 
     * - "none"      Indicates there is no healthcheck
     * - "starting"  Starting indicates that the container is not yet ready
     * - "healthy"   Healthy indicates that the container is running correctly
     * - "unhealthy" Unhealthy indicates that the container has a problem
     */
    'Status'?: "none" | "starting" | "healthy" | "unhealthy";
    /**
     * FailingStreak is the number of consecutive failures
     */
    'FailingStreak'?: number;
    /**
     * Log contains the last few results (oldest first)
     */
    'Log'?: Array<HealthcheckResult>;
  };

  type HealthcheckResult = {
    /**
     * Date and time at which this check started in
     * [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
     */
    'Start'?: string;
    /**
     * Date and time at which this check ended in
     * [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
     */
    'End'?: string;
    /**
     * ExitCode meanings:
     * 
     * - `0` healthy
     * - `1` unhealthy
     * - `2` reserved (considered unhealthy)
     * - other values: error running probe
     */
    'ExitCode'?: number;
    /**
     * Output from last check
     */
    'Output'?: string;
  };

  type HostConfig = Resources & {
    /**
     * A list of volume bindings for this container. Each volume binding
     * is a string in one of these forms:
     * 
     * - `host-src:container-dest[:options]` to bind-mount a host path
     *   into the container. Both `host-src`, and `container-dest` must
     *   be an _absolute_ path.
     * - `volume-name:container-dest[:options]` to bind-mount a volume
     *   managed by a volume driver into the container. `container-dest`
     *   must be an _absolute_ path.
     * 
     * `options` is an optional, comma-delimited list of:
     * 
     * - `nocopy` disables automatic copying of data from the container
     *   path to the volume. The `nocopy` flag only applies to named volumes.
     * - `[ro|rw]` mounts a volume read-only or read-write, respectively.
     *   If omitted or set to `rw`, volumes are mounted read-write.
     * - `[z|Z]` applies SELinux labels to allow or deny multiple containers
     *   to read and write to the same volume.
     *     - `z`: a _shared_ content label is applied to the content. This
     *       label indicates that multiple containers can share the volume
     *       content, for both reading and writing.
     *     - `Z`: a _private unshared_ label is applied to the content.
     *       This label indicates that only the current container can use
     *       a private volume. Labeling systems such as SELinux require
     *       proper labels to be placed on volume content that is mounted
     *       into a container. Without a label, the security system can
     *       prevent a container's processes from using the content. By
     *       default, the labels set by the host operating system are not
     *       modified.
     * - `[[r]shared|[r]slave|[r]private]` specifies mount
     *   [propagation behavior](https://www.kernel.org/doc/Documentation/filesystems/sharedsubtree.txt).
     *   This only applies to bind-mounted volumes, not internal volumes
     *   or named volumes. Mount propagation requires the source mount
     *   point (the location where the source directory is mounted in the
     *   host operating system) to have the correct propagation properties.
     *   For shared volumes, the source mount point must be set to `shared`.
     *   For slave volumes, the mount must be set to either `shared` or
     *   `slave`.
     */
    'Binds'?: Array<string>;
    /**
     * Path to a file where the container ID is written
     */
    'ContainerIDFile'?: string;
    /**
     * The logging configuration for this container
     */
    'LogConfig'?: {
      'Type'?: "json-file" | "syslog" | "journald" | "gelf" | "fluentd" | "awslogs" | "splunk" | "etwlogs" | "none";
      'Config'?: Record<string, string>;
    };
    /**
     * Network mode to use for this container. Supported standard values
     * are: `bridge`, `host`, `none`, and `container:<name|id>`. Any
     * other value is taken as a custom network's name to which this
     * container should connect to.
     */
    'NetworkMode'?: string;
    'PortBindings'?: PortMap;
    'RestartPolicy'?: RestartPolicy;
    /**
     * Automatically remove the container when the container's process
     * exits. This has no effect if `RestartPolicy` is set.
     */
    'AutoRemove'?: boolean;
    /**
     * Driver that this container uses to mount volumes.
     */
    'VolumeDriver'?: string;
    /**
     * A list of volumes to inherit from another container, specified in
     * the form `<container name>[:<ro|rw>]`.
     */
    'VolumesFrom'?: Array<string>;
    /**
     * Specification for mounts to be added to the container.
     */
    'Mounts'?: Array<Mount>;
    /**
     * A list of kernel capabilities to add to the container. Conflicts
     * with option 'Capabilities'.
     */
    'CapAdd'?: Array<string>;
    /**
     * A list of kernel capabilities to drop from the container. Conflicts
     * with option 'Capabilities'.
     */
    'CapDrop'?: Array<string>;
    /**
     * cgroup namespace mode for the container. Possible values are:
     * 
     * - `"private"`: the container runs in its own private cgroup namespace
     * - `"host"`: use the host system's cgroup namespace
     * 
     * If not specified, the daemon default is used, which can either be `"private"`
     * or `"host"`, depending on daemon version, kernel support and configuration.
     */
    'CgroupnsMode'?: "private" | "host";
    /**
     * A list of DNS servers for the container to use.
     */
    'Dns'?: Array<string>;
    /**
     * A list of DNS options.
     */
    'DnsOptions'?: Array<string>;
    /**
     * A list of DNS search domains.
     */
    'DnsSearch'?: Array<string>;
    /**
     * A list of hostnames/IP mappings to add to the container's `/etc/hosts`
     * file. Specified in the form `["hostname:IP"]`.
     */
    'ExtraHosts'?: Array<string>;
    /**
     * A list of additional groups that the container process will run as.
     */
    'GroupAdd'?: Array<string>;
    /**
     * IPC sharing mode for the container. Possible values are:
     * 
     * - `"none"`: own private IPC namespace, with /dev/shm not mounted
     * - `"private"`: own private IPC namespace
     * - `"shareable"`: own private IPC namespace, with a possibility to share it with other containers
     * - `"container:<name|id>"`: join another (shareable) container's IPC namespace
     * - `"host"`: use the host system's IPC namespace
     * 
     * If not specified, daemon default is used, which can either be `"private"`
     * or `"shareable"`, depending on daemon version and configuration.
     */
    'IpcMode'?: string;
    /**
     * Cgroup to use for the container.
     */
    'Cgroup'?: string;
    /**
     * A list of links for the container in the form `container_name:alias`.
     */
    'Links'?: Array<string>;
    /**
     * An integer value containing the score given to the container in
     * order to tune OOM killer preferences.
     */
    'OomScoreAdj'?: number;
    /**
     * Set the PID (Process) Namespace mode for the container. It can be
     * either:
     * 
     * - `"container:<name|id>"`: joins another container's PID namespace
     * - `"host"`: use the host's PID namespace inside the container
     */
    'PidMode'?: string;
    /**
     * Gives the container full access to the host.
     */
    'Privileged'?: boolean;
    /**
     * Allocates an ephemeral host port for all of a container's
     * exposed ports.
     * 
     * Ports are de-allocated when the container stops and allocated when
     * the container starts. The allocated port might be changed when
     * restarting the container.
     * 
     * The port is selected from the ephemeral port range that depends on
     * the kernel. For example, on Linux the range is defined by
     * `/proc/sys/net/ipv4/ip_local_port_range`.
     */
    'PublishAllPorts'?: boolean;
    /**
     * Mount the container's root filesystem as read only.
     */
    'ReadonlyRootfs'?: boolean;
    /**
     * A list of string values to customize labels for MLS systems, such
     * as SELinux.
     */
    'SecurityOpt'?: Array<string>;
    /**
     * Storage driver options for this container, in the form `{"size": "120G"}`.
     */
    'StorageOpt'?: Record<string, string>;
    /**
     * A map of container directories which should be replaced by tmpfs
     * mounts, and their corresponding mount options. For example:
     * 
     * ```
     * { "/run": "rw,noexec,nosuid,size=65536k" }
     * ```
     */
    'Tmpfs'?: Record<string, string>;
    /**
     * UTS namespace to use for the container.
     */
    'UTSMode'?: string;
    /**
     * Sets the usernamespace mode for the container when usernamespace
     * remapping option is enabled.
     */
    'UsernsMode'?: string;
    /**
     * Size of `/dev/shm` in bytes. If omitted, the system uses 64MB.
     */
    'ShmSize'?: number;
    /**
     * A list of kernel parameters (sysctls) to set in the container.
     * For example:
     * 
     * ```
     * {"net.ipv4.ip_forward": "1"}
     * ```
     */
    'Sysctls'?: Record<string, string>;
    /**
     * Runtime to use with this container.
     */
    'Runtime'?: string;
    /**
     * Initial console size, as an `[height, width]` array. (Windows only)
     */
    'ConsoleSize'?: Array<number>;
    /**
     * Isolation technology of the container. (Windows only)
     */
    'Isolation'?: "default" | "process" | "hyperv";
    /**
     * The list of paths to be masked inside the container (this overrides
     * the default set of paths).
     */
    'MaskedPaths'?: Array<string>;
    /**
     * The list of paths to be set as read-only inside the container
     * (this overrides the default set of paths).
     */
    'ReadonlyPaths'?: Array<string>;
  };

  type ContainerConfig = {
    /**
     * The hostname to use for the container, as a valid RFC 1123 hostname.
     */
    'Hostname'?: string;
    /**
     * The domain name to use for the container.
     */
    'Domainname'?: string;
    /**
     * The user that commands are run as inside the container.
     */
    'User'?: string;
    /**
     * Whether to attach to `stdin`.
     */
    'AttachStdin'?: boolean;
    /**
     * Whether to attach to `stdout`.
     */
    'AttachStdout'?: boolean;
    /**
     * Whether to attach to `stderr`.
     */
    'AttachStderr'?: boolean;
    /**
     * An object mapping ports to an empty object in the form:
     * 
     * `{"<port>/<tcp|udp|sctp>": {}}`
     */
    'ExposedPorts'?: Record<string, unknown>;
    /**
     * Attach standard streams to a TTY, including `stdin` if it is not closed.
     */
    'Tty'?: boolean;
    /**
     * Open `stdin`
     */
    'OpenStdin'?: boolean;
    /**
     * Close `stdin` after one attached client disconnects
     */
    'StdinOnce'?: boolean;
    /**
     * A list of environment variables to set inside the container in the
     * form `["VAR=value", ...]`. A variable without `=` is removed from the
     * environment, rather than to have an empty value.
     */
    'Env'?: Array<string>;
    /**
     * Command to run specified as a string or an array of strings.
     */
    'Cmd'?: Array<string>;
    'Healthcheck'?: HealthConfig;
    /**
     * Command is already escaped (Windows only)
     */
    'ArgsEscaped'?: boolean;
    /**
     * The name (or reference) of the image to use when creating the container,
     * or which was used when the container was created.
     */
    'Image'?: string;
    /**
     * An object mapping mount point paths inside the container to empty
     * objects.
     */
    'Volumes'?: Record<string, unknown>;
    /**
     * The working directory for commands to run in.
     */
    'WorkingDir'?: string;
    /**
     * The entry point for the container as a string or an array of strings.
     * 
     * If the array consists of exactly one empty string (`[""]`) then the
     * entry point is reset to system default (i.e., the entry point used by
     * docker when there is no `ENTRYPOINT` instruction in the `Dockerfile`).
     */
    'Entrypoint'?: Array<string>;
    /**
     * Disable networking for the container.
     */
    'NetworkDisabled'?: boolean;
    /**
     * MAC address of the container.
     */
    'MacAddress'?: string;
    /**
     * `ONBUILD` metadata that were defined in the image's `Dockerfile`.
     */
    'OnBuild'?: Array<string>;
    /**
     * User-defined key/value metadata.
     */
    'Labels'?: Record<string, string>;
    /**
     * Signal to stop a container as a string or unsigned integer.
     */
    'StopSignal'?: string;
    /**
     * Timeout to stop a container in seconds.
     */
    'StopTimeout'?: number;
    /**
     * Shell for when `RUN`, `CMD`, and `ENTRYPOINT` uses a shell.
     */
    'Shell'?: Array<string>;
  };

  type NetworkingConfig = {
    /**
     * A mapping of network name to endpoint configuration for that network.
     */
    'EndpointsConfig'?: Record<string, EndpointSettings>;
  };

  type NetworkSettings = {
    /**
     * Name of the network'a bridge (for example, `docker0`).
     */
    'Bridge'?: string;
    /**
     * SandboxID uniquely represents a container's network stack.
     */
    'SandboxID'?: string;
    /**
     * Indicates if hairpin NAT should be enabled on the virtual interface.
     */
    'HairpinMode'?: boolean;
    /**
     * IPv6 unicast address using the link-local prefix.
     */
    'LinkLocalIPv6Address'?: string;
    /**
     * Prefix length of the IPv6 unicast address.
     */
    'LinkLocalIPv6PrefixLen'?: number;
    'Ports'?: PortMap;
    /**
     * SandboxKey identifies the sandbox
     */
    'SandboxKey'?: string;
    'SecondaryIPAddresses'?: Array<Address>;
    'SecondaryIPv6Addresses'?: Array<Address>;
    /**
     * EndpointID uniquely represents a service endpoint in a Sandbox.
     * 
     * <p><br /></p>
     * 
     * > **Deprecated**: This field is only propagated when attached to the
     * > default "bridge" network. Use the information from the "bridge"
     * > network inside the `Networks` map instead, which contains the same
     * > information. This field was deprecated in Docker 1.9 and is scheduled
     * > to be removed in Docker 17.12.0
     */
    'EndpointID'?: string;
    /**
     * Gateway address for the default "bridge" network.
     * 
     * <p><br /></p>
     * 
     * > **Deprecated**: This field is only propagated when attached to the
     * > default "bridge" network. Use the information from the "bridge"
     * > network inside the `Networks` map instead, which contains the same
     * > information. This field was deprecated in Docker 1.9 and is scheduled
     * > to be removed in Docker 17.12.0
     */
    'Gateway'?: string;
    /**
     * Global IPv6 address for the default "bridge" network.
     * 
     * <p><br /></p>
     * 
     * > **Deprecated**: This field is only propagated when attached to the
     * > default "bridge" network. Use the information from the "bridge"
     * > network inside the `Networks` map instead, which contains the same
     * > information. This field was deprecated in Docker 1.9 and is scheduled
     * > to be removed in Docker 17.12.0
     */
    'GlobalIPv6Address'?: string;
    /**
     * Mask length of the global IPv6 address.
     * 
     * <p><br /></p>
     * 
     * > **Deprecated**: This field is only propagated when attached to the
     * > default "bridge" network. Use the information from the "bridge"
     * > network inside the `Networks` map instead, which contains the same
     * > information. This field was deprecated in Docker 1.9 and is scheduled
     * > to be removed in Docker 17.12.0
     */
    'GlobalIPv6PrefixLen'?: number;
    /**
     * IPv4 address for the default "bridge" network.
     * 
     * <p><br /></p>
     * 
     * > **Deprecated**: This field is only propagated when attached to the
     * > default "bridge" network. Use the information from the "bridge"
     * > network inside the `Networks` map instead, which contains the same
     * > information. This field was deprecated in Docker 1.9 and is scheduled
     * > to be removed in Docker 17.12.0
     */
    'IPAddress'?: string;
    /**
     * Mask length of the IPv4 address.
     * 
     * <p><br /></p>
     * 
     * > **Deprecated**: This field is only propagated when attached to the
     * > default "bridge" network. Use the information from the "bridge"
     * > network inside the `Networks` map instead, which contains the same
     * > information. This field was deprecated in Docker 1.9 and is scheduled
     * > to be removed in Docker 17.12.0
     */
    'IPPrefixLen'?: number;
    /**
     * IPv6 gateway address for this network.
     * 
     * <p><br /></p>
     * 
     * > **Deprecated**: This field is only propagated when attached to the
     * > default "bridge" network. Use the information from the "bridge"
     * > network inside the `Networks` map instead, which contains the same
     * > information. This field was deprecated in Docker 1.9 and is scheduled
     * > to be removed in Docker 17.12.0
     */
    'IPv6Gateway'?: string;
    /**
     * MAC address for the container on the default "bridge" network.
     * 
     * <p><br /></p>
     * 
     * > **Deprecated**: This field is only propagated when attached to the
     * > default "bridge" network. Use the information from the "bridge"
     * > network inside the `Networks` map instead, which contains the same
     * > information. This field was deprecated in Docker 1.9 and is scheduled
     * > to be removed in Docker 17.12.0
     */
    'MacAddress'?: string;
    /**
     * Information about all networks that the container is connected to.
     */
    'Networks'?: Record<string, EndpointSettings>;
  };

  type Address = {
    /**
     * IP address.
     */
    'Addr'?: string;
    /**
     * Mask length of the IP address.
     */
    'PrefixLen'?: number;
  };

  type PortMap = Record<string, Array<PortBinding>>;

  type PortBinding = {
    /**
     * Host IP address that the container's port is mapped to.
     */
    'HostIp'?: string;
    /**
     * Host port number that the container's port is mapped to.
     */
    'HostPort'?: string;
  };

  type GraphDriverData = {
    /**
     * Name of the storage driver.
     */
    'Name'?: string;
    /**
     * Low-level storage metadata, provided as key/value pairs.
     * 
     * This information is driver-specific, and depends on the storage-driver
     * in use, and should be used for informational purposes only.
     */
    'Data'?: Record<string, string>;
  };

  type ImageInspect = {
    /**
     * ID is the content-addressable ID of an image.
     * 
     * This identifier is a content-addressable digest calculated from the
     * image's configuration (which includes the digests of layers used by
     * the image).
     * 
     * Note that this digest differs from the `RepoDigests` below, which
     * holds digests of image manifests that reference the image.
     */
    'Id'?: string;
    /**
     * List of image names/tags in the local image cache that reference this
     * image.
     * 
     * Multiple image tags can refer to the same imagem and this list may be
     * empty if no tags reference the image, in which case the image is
     * "untagged", in which case it can still be referenced by its ID.
     */
    'RepoTags'?: Array<string>;
    /**
     * List of content-addressable digests of locally available image manifests
     * that the image is referenced from. Multiple manifests can refer to the
     * same image.
     * 
     * These digests are usually only available if the image was either pulled
     * from a registry, or if the image was pushed to a registry, which is when
     * the manifest is generated and its digest calculated.
     */
    'RepoDigests'?: Array<string>;
    /**
     * ID of the parent image.
     * 
     * Depending on how the image was created, this field may be empty and
     * is only set for images that were built/created locally. This field
     * is empty if the image was pulled from an image registry.
     */
    'Parent'?: string;
    /**
     * Optional message that was set when committing or importing the image.
     */
    'Comment'?: string;
    /**
     * Date and time at which the image was created, formatted in
     * [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
     */
    'Created'?: string;
    /**
     * The ID of the container that was used to create the image.
     * 
     * Depending on how the image was created, this field may be empty.
     */
    'Container'?: string;
    'ContainerConfig'?: ContainerConfig;
    /**
     * The version of Docker that was used to build the image.
     * 
     * Depending on how the image was created, this field may be empty.
     */
    'DockerVersion'?: string;
    /**
     * Name of the author that was specified when committing the image, or as
     * specified through MAINTAINER (deprecated) in the Dockerfile.
     */
    'Author'?: string;
    'Config'?: ContainerConfig;
    /**
     * Hardware CPU architecture that the image runs on.
     */
    'Architecture'?: string;
    /**
     * CPU architecture variant (presently ARM-only).
     */
    'Variant'?: string;
    /**
     * Operating System the image is built to run on.
     */
    'Os'?: string;
    /**
     * Operating System version the image is built to run on (especially
     * for Windows).
     */
    'OsVersion'?: string;
    /**
     * Total size of the image including all layers it is composed of.
     */
    'Size'?: number;
    /**
     * Total size of the image including all layers it is composed of.
     * 
     * In versions of Docker before v1.10, this field was calculated from
     * the image itself and all of its parent images. Docker v1.10 and up
     * store images self-contained, and no longer use a parent-chain, making
     * this field an equivalent of the Size field.
     * 
     * This field is kept for backward compatibility, but may be removed in
     * a future version of the API.
     */
    'VirtualSize'?: number;
    'GraphDriver'?: GraphDriverData;
    /**
     * Information about the image's RootFS, including the layer IDs.
     */
    'RootFS': {
      'Type'?: string;
      'Layers'?: Array<string>;
    };
    /**
     * Additional metadata of the image in the local cache. This information
     * is local to the daemon, and not part of the image itself.
     */
    'Metadata'?: {
      /**
       * Date and time at which the image was last tagged in
       * [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
       * 
       * This information is only available if the image was tagged locally,
       * and omitted otherwise.
       */
      'LastTagTime'?: string;
    };
  };

  type ImageSummary = {
    /**
     * ID is the content-addressable ID of an image.
     * 
     * This identifier is a content-addressable digest calculated from the
     * image's configuration (which includes the digests of layers used by
     * the image).
     * 
     * Note that this digest differs from the `RepoDigests` below, which
     * holds digests of image manifests that reference the image.
     */
    'Id'?: string;
    /**
     * ID of the parent image.
     * 
     * Depending on how the image was created, this field may be empty and
     * is only set for images that were built/created locally. This field
     * is empty if the image was pulled from an image registry.
     */
    'ParentId'?: string;
    /**
     * List of image names/tags in the local image cache that reference this
     * image.
     * 
     * Multiple image tags can refer to the same imagem and this list may be
     * empty if no tags reference the image, in which case the image is
     * "untagged", in which case it can still be referenced by its ID.
     */
    'RepoTags'?: Array<string>;
    /**
     * List of content-addressable digests of locally available image manifests
     * that the image is referenced from. Multiple manifests can refer to the
     * same image.
     * 
     * These digests are usually only available if the image was either pulled
     * from a registry, or if the image was pushed to a registry, which is when
     * the manifest is generated and its digest calculated.
     */
    'RepoDigests'?: Array<string>;
    /**
     * Date and time at which the image was created as a Unix timestamp
     * (number of seconds sinds EPOCH).
     */
    'Created'?: number;
    /**
     * Total size of the image including all layers it is composed of.
     */
    'Size'?: number;
    /**
     * Total size of image layers that are shared between this image and other
     * images.
     * 
     * This size is not calculated by default. `-1` indicates that the value
     * has not been set / calculated.
     */
    'SharedSize'?: number;
    /**
     * Total size of the image including all layers it is composed of.
     * 
     * In versions of Docker before v1.10, this field was calculated from
     * the image itself and all of its parent images. Docker v1.10 and up
     * store images self-contained, and no longer use a parent-chain, making
     * this field an equivalent of the Size field.
     * 
     * This field is kept for backward compatibility, but may be removed in
     * a future version of the API.
     */
    'VirtualSize'?: number;
    /**
     * User-defined key/value metadata.
     */
    'Labels'?: Record<string, string>;
    /**
     * Number of containers using this image. Includes both stopped and running
     * containers.
     * 
     * This size is not calculated by default, and depends on which API endpoint
     * is used. `-1` indicates that the value has not been set / calculated.
     */
    'Containers'?: number;
  };

  type AuthConfig = {
    'username'?: string;
    'password'?: string;
    'email'?: string;
    'serveraddress'?: string;
  };

  type ProcessConfig = {
    'privileged'?: boolean;
    'user'?: string;
    'tty'?: boolean;
    'entrypoint'?: string;
    'arguments'?: Array<string>;
  };

  type Volume = {
    /**
     * Name of the volume.
     */
    'Name'?: string;
    /**
     * Name of the volume driver used by the volume.
     */
    'Driver'?: string;
    /**
     * Mount path of the volume on the host.
     */
    'Mountpoint'?: string;
    /**
     * Date/Time the volume was created.
     */
    'CreatedAt'?: string;
    /**
     * Low-level details about the volume, provided by the volume driver.
     * Details are returned as a map with key/value pairs:
     * `{"key":"value","key2":"value2"}`.
     * 
     * The `Status` field is optional, and is omitted if the volume driver
     * does not support this feature.
     */
    'Status'?: Record<string, {
    }>;
    /**
     * User-defined key/value metadata.
     */
    'Labels'?: Record<string, string>;
    /**
     * The level at which the volume exists. Either `global` for cluster-wide,
     * or `local` for machine level.
     */
    'Scope'?: "local" | "global";
    /**
     * The driver specific options used when creating the volume.
     */
    'Options'?: Record<string, string>;
    /**
     * Usage details about the volume. This information is used by the
     * `GET /system/df` endpoint, and omitted in other endpoints.
     */
    'UsageData': {
      /**
       * Amount of disk space used by the volume (in bytes). This information
       * is only available for volumes created with the `"local"` volume
       * driver. For volumes created with other volume drivers, this field
       * is set to `-1` ("not available")
       */
      'Size'?: number;
      /**
       * The number of containers referencing this volume. This field
       * is set to `-1` if the reference-count is not available.
       */
      'RefCount'?: number;
    };
  };

  type VolumeCreateOptions = {
    /**
     * The new volume's name. If not specified, Docker generates a name.
     */
    'Name'?: string;
    /**
     * Name of the volume driver to use.
     */
    'Driver'?: string;
    /**
     * A mapping of driver options and values. These options are
     * passed directly to the driver and are driver specific.
     */
    'DriverOpts'?: Record<string, string>;
    /**
     * User-defined key/value metadata.
     */
    'Labels'?: Record<string, string>;
  };

  type Network = {
    'Name'?: string;
    'Id'?: string;
    'Created'?: string;
    'Scope'?: string;
    'Driver'?: string;
    'EnableIPv6'?: boolean;
    'IPAM'?: IPAM;
    'Internal'?: boolean;
    'Attachable'?: boolean;
    'Ingress'?: boolean;
    'Containers'?: Record<string, NetworkContainer>;
    'Options'?: Record<string, string>;
    'Labels'?: Record<string, string>;
  };

  type IPAM = {
    /**
     * Name of the IPAM driver to use.
     */
    'Driver'?: string;
    /**
     * List of IPAM configuration options, specified as a map:
     * 
     * ```
     * {"Subnet": <CIDR>, "IPRange": <CIDR>, "Gateway": <IP address>, "AuxAddress": <device_name:IP address>}
     * ```
     */
    'Config'?: Array<IPAMConfig>;
    /**
     * Driver-specific options, specified as a map.
     */
    'Options'?: Record<string, string>;
  };

  type IPAMConfig = {
    'Subnet'?: string;
    'IPRange'?: string;
    'Gateway'?: string;
    'AuxiliaryAddresses'?: Record<string, string>;
  };

  type NetworkContainer = {
    'Name'?: string;
    'EndpointID'?: string;
    'MacAddress'?: string;
    'IPv4Address'?: string;
    'IPv6Address'?: string;
  };

  type BuildInfo = {
    'id'?: string;
    'stream'?: string;
    'error'?: string;
    'errorDetail'?: ErrorDetail;
    'status'?: string;
    'progress'?: string;
    'progressDetail'?: ProgressDetail;
    'aux'?: ImageID;
  };

  type BuildCache = {
    /**
     * Unique ID of the build cache record.
     */
    'ID'?: string;
    /**
     * ID of the parent build cache record.
     */
    'Parent'?: string;
    /**
     * Cache record type.
     */
    'Type'?: "internal" | "frontend" | "source.local" | "source.git.checkout" | "exec.cachemount" | "regular";
    /**
     * Description of the build-step that produced the build cache.
     */
    'Description'?: string;
    /**
     * Indicates if the build cache is in use.
     */
    'InUse'?: boolean;
    /**
     * Indicates if the build cache is shared.
     */
    'Shared'?: boolean;
    /**
     * Amount of disk space used by the build cache (in bytes).
     */
    'Size'?: number;
    /**
     * Date and time at which the build cache was created in
     * [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
     */
    'CreatedAt'?: string;
    /**
     * Date and time at which the build cache was last used in
     * [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
     */
    'LastUsedAt'?: string;
    'UsageCount'?: number;
  };

  type ImageID = {
    'ID'?: string;
  };

  type CreateImageInfo = {
    'id'?: string;
    'error'?: string;
    'status'?: string;
    'progress'?: string;
    'progressDetail'?: ProgressDetail;
  };

  type PushImageInfo = {
    'error'?: string;
    'status'?: string;
    'progress'?: string;
    'progressDetail'?: ProgressDetail;
  };

  type ErrorDetail = {
    'code'?: number;
    'message'?: string;
  };

  type ProgressDetail = {
    'current'?: number;
    'total'?: number;
  };

  type ErrorResponse = {
    /**
     * The error message.
     */
    'message'?: string;
  };

  type IdResponse = {
    /**
     * The id of the newly created object.
     */
    'Id'?: string;
  };

  type EndpointSettings = {
    'IPAMConfig'?: EndpointIPAMConfig;
    'Links'?: Array<string>;
    'Aliases'?: Array<string>;
    /**
     * Unique ID of the network.
     */
    'NetworkID'?: string;
    /**
     * Unique ID for the service endpoint in a Sandbox.
     */
    'EndpointID'?: string;
    /**
     * Gateway address for this network.
     */
    'Gateway'?: string;
    /**
     * IPv4 address.
     */
    'IPAddress'?: string;
    /**
     * Mask length of the IPv4 address.
     */
    'IPPrefixLen'?: number;
    /**
     * IPv6 gateway address.
     */
    'IPv6Gateway'?: string;
    /**
     * Global IPv6 address.
     */
    'GlobalIPv6Address'?: string;
    /**
     * Mask length of the global IPv6 address.
     */
    'GlobalIPv6PrefixLen'?: number;
    /**
     * MAC address for the endpoint on this network.
     */
    'MacAddress'?: string;
    /**
     * DriverOpts is a mapping of driver options and values. These options
     * are passed directly to the driver and are driver specific.
     */
    'DriverOpts'?: Record<string, string>;
  };

  type EndpointIPAMConfig = {
    'IPv4Address'?: string;
    'IPv6Address'?: string;
    'LinkLocalIPs'?: Array<string>;
  };

  type PluginMount = {
    'Name'?: string;
    'Description'?: string;
    'Settable'?: Array<string>;
    'Source'?: string;
    'Destination'?: string;
    'Type'?: string;
    'Options'?: Array<string>;
  };

  type PluginDevice = {
    'Name'?: string;
    'Description'?: string;
    'Settable'?: Array<string>;
    'Path'?: string;
  };

  type PluginEnv = {
    'Name'?: string;
    'Description'?: string;
    'Settable'?: Array<string>;
    'Value'?: string;
  };

  type PluginInterfaceType = {
    'Prefix'?: string;
    'Capability'?: string;
    'Version'?: string;
  };

  type PluginPrivilege = {
    'Name'?: string;
    'Description'?: string;
    'Value'?: Array<string>;
  };

  type Plugin = {
    'Id'?: string;
    'Name'?: string;
    /**
     * True if the plugin is running. False if the plugin is not running, only installed.
     */
    'Enabled'?: boolean;
    /**
     * Settings that can be modified by users.
     */
    'Settings': {
      'Mounts'?: Array<PluginMount>;
      'Env'?: Array<string>;
      'Args'?: Array<string>;
      'Devices'?: Array<PluginDevice>;
    };
    /**
     * plugin remote reference used to push/pull the plugin
     */
    'PluginReference'?: string;
    /**
     * The config of a plugin.
     */
    'Config': {
      /**
       * Docker Version used to create the plugin
       */
      'DockerVersion'?: string;
      'Description'?: string;
      'Documentation'?: string;
      /**
       * The interface between Docker and the plugin
       */
      'Interface': {
        'Types'?: Array<PluginInterfaceType>;
        'Socket'?: string;
        /**
         * Protocol to use for clients connecting to the plugin.
         */
        'ProtocolScheme'?: "" | "moby.plugins.http/v1";
      };
      'Entrypoint'?: Array<string>;
      'WorkDir'?: string;
      'User'?: {
        'UID'?: number;
        'GID'?: number;
      };
      'Network': {
        'Type'?: string;
      };
      'Linux': {
        'Capabilities'?: Array<string>;
        'AllowAllDevices'?: boolean;
        'Devices'?: Array<PluginDevice>;
      };
      'PropagatedMount'?: string;
      'IpcHost'?: boolean;
      'PidHost'?: boolean;
      'Mounts'?: Array<PluginMount>;
      'Env'?: Array<PluginEnv>;
      'Args': {
        'Name'?: string;
        'Description'?: string;
        'Settable'?: Array<string>;
        'Value'?: Array<string>;
      };
      'rootfs'?: {
        'type'?: string;
        'diff_ids'?: Array<string>;
      };
    };
  };

  type ObjectVersion = {
    'Index'?: number;
  };

  type NodeSpec = {
    /**
     * Name for the node.
     */
    'Name'?: string;
    /**
     * User-defined key/value metadata.
     */
    'Labels'?: Record<string, string>;
    /**
     * Role of the node.
     */
    'Role'?: "worker" | "manager";
    /**
     * Availability of the node.
     */
    'Availability'?: "active" | "pause" | "drain";
  };

  type Node = {
    'ID'?: string;
    'Version'?: ObjectVersion;
    /**
     * Date and time at which the node was added to the swarm in
     * [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
     */
    'CreatedAt'?: string;
    /**
     * Date and time at which the node was last updated in
     * [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
     */
    'UpdatedAt'?: string;
    'Spec'?: NodeSpec;
    'Description'?: NodeDescription;
    'Status'?: NodeStatus;
    'ManagerStatus'?: ManagerStatus;
  };

  type NodeDescription = {
    'Hostname'?: string;
    'Platform'?: Platform;
    'Resources'?: ResourceObject;
    'Engine'?: EngineDescription;
    'TLSInfo'?: TLSInfo;
  };

  type Platform = {
    /**
     * Architecture represents the hardware architecture (for example,
     * `x86_64`).
     */
    'Architecture'?: string;
    /**
     * OS represents the Operating System (for example, `linux` or `windows`).
     */
    'OS'?: string;
  };

  type EngineDescription = {
    'EngineVersion'?: string;
    'Labels'?: Record<string, string>;
    'Plugins'?: Array<{
      'Type'?: string;
      'Name'?: string;
    }>;
  };

  type TLSInfo = {
    /**
     * The root CA certificate(s) that are used to validate leaf TLS
     * certificates.
     */
    'TrustRoot'?: string;
    /**
     * The base64-url-safe-encoded raw subject bytes of the issuer.
     */
    'CertIssuerSubject'?: string;
    /**
     * The base64-url-safe-encoded raw public key bytes of the issuer.
     */
    'CertIssuerPublicKey'?: string;
  };

  type NodeStatus = {
    'State'?: NodeState;
    'Message'?: string;
    /**
     * IP address of the node.
     */
    'Addr'?: string;
  };

  type NodeState = "unknown" | "down" | "ready" | "disconnected";

  type ManagerStatus = {
    'Leader'?: boolean;
    'Reachability'?: Reachability;
    /**
     * The IP address and port at which the manager is reachable.
     */
    'Addr'?: string;
  };

  type Reachability = "unknown" | "unreachable" | "reachable";

  type SwarmSpec = {
    /**
     * Name of the swarm.
     */
    'Name'?: string;
    /**
     * User-defined key/value metadata.
     */
    'Labels'?: Record<string, string>;
    /**
     * Orchestration configuration.
     */
    'Orchestration'?: {
      /**
       * The number of historic tasks to keep per instance or node. If
       * negative, never remove completed or failed tasks.
       */
      'TaskHistoryRetentionLimit'?: number;
    };
    /**
     * Raft configuration.
     */
    'Raft'?: {
      /**
       * The number of log entries between snapshots.
       */
      'SnapshotInterval'?: number;
      /**
       * The number of snapshots to keep beyond the current snapshot.
       */
      'KeepOldSnapshots'?: number;
      /**
       * The number of log entries to keep around to sync up slow followers
       * after a snapshot is created.
       */
      'LogEntriesForSlowFollowers'?: number;
      /**
       * The number of ticks that a follower will wait for a message from
       * the leader before becoming a candidate and starting an election.
       * `ElectionTick` must be greater than `HeartbeatTick`.
       * 
       * A tick currently defaults to one second, so these translate
       * directly to seconds currently, but this is NOT guaranteed.
       */
      'ElectionTick'?: number;
      /**
       * The number of ticks between heartbeats. Every HeartbeatTick ticks,
       * the leader will send a heartbeat to the followers.
       * 
       * A tick currently defaults to one second, so these translate
       * directly to seconds currently, but this is NOT guaranteed.
       */
      'HeartbeatTick'?: number;
    };
    /**
     * Dispatcher configuration.
     */
    'Dispatcher'?: {
      /**
       * The delay for an agent to send a heartbeat to the dispatcher.
       */
      'HeartbeatPeriod'?: number;
    };
    /**
     * CA configuration.
     */
    'CAConfig'?: {
      /**
       * The duration node certificates are issued for.
       */
      'NodeCertExpiry'?: number;
      /**
       * Configuration for forwarding signing requests to an external
       * certificate authority.
       */
      'ExternalCAs'?: Array<{
        /**
         * Protocol for communication with the external CA (currently
         * only `cfssl` is supported).
         */
        'Protocol'?: "cfssl";
        /**
         * URL where certificate signing requests should be sent.
         */
        'URL'?: string;
        /**
         * An object with key/value pairs that are interpreted as
         * protocol-specific options for the external CA driver.
         */
        'Options'?: Record<string, string>;
        /**
         * The root CA certificate (in PEM format) this external CA uses
         * to issue TLS certificates (assumed to be to the current swarm
         * root CA certificate if not provided).
         */
        'CACert'?: string;
      }>;
      /**
       * The desired signing CA certificate for all swarm node TLS leaf
       * certificates, in PEM format.
       */
      'SigningCACert'?: string;
      /**
       * The desired signing CA key for all swarm node TLS leaf certificates,
       * in PEM format.
       */
      'SigningCAKey'?: string;
      /**
       * An integer whose purpose is to force swarm to generate a new
       * signing CA certificate and key, if none have been specified in
       * `SigningCACert` and `SigningCAKey`
       */
      'ForceRotate'?: number;
    };
    /**
     * Parameters related to encryption-at-rest.
     */
    'EncryptionConfig'?: {
      /**
       * If set, generate a key and use it to lock data stored on the
       * managers.
       */
      'AutoLockManagers'?: boolean;
    };
    /**
     * Defaults for creating tasks in this cluster.
     */
    'TaskDefaults'?: {
      /**
       * The log driver to use for tasks created in the orchestrator if
       * unspecified by a service.
       * 
       * Updating this value only affects new tasks. Existing tasks continue
       * to use their previously configured log driver until recreated.
       */
      'LogDriver'?: {
        /**
         * The log driver to use as a default for new tasks.
         */
        'Name'?: string;
        /**
         * Driver-specific options for the selectd log driver, specified
         * as key/value pairs.
         */
        'Options'?: Record<string, string>;
      };
    };
  };

  type ClusterInfo = {
    /**
     * The ID of the swarm.
     */
    'ID'?: string;
    'Version'?: ObjectVersion;
    /**
     * Date and time at which the swarm was initialised in
     * [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
     */
    'CreatedAt'?: string;
    /**
     * Date and time at which the swarm was last updated in
     * [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
     */
    'UpdatedAt'?: string;
    'Spec'?: SwarmSpec;
    'TLSInfo'?: TLSInfo;
    /**
     * Whether there is currently a root CA rotation in progress for the swarm
     */
    'RootRotationInProgress'?: boolean;
    /**
     * DataPathPort specifies the data path port number for data traffic.
     * Acceptable port range is 1024 to 49151.
     * If no port is set or is set to 0, the default port (4789) is used.
     */
    'DataPathPort'?: number;
    /**
     * Default Address Pool specifies default subnet pools for global scope
     * networks.
     */
    'DefaultAddrPool'?: Array<string>;
    /**
     * SubnetSize specifies the subnet size of the networks created from the
     * default subnet pool.
     */
    'SubnetSize'?: number;
  };

  type JoinTokens = {
    /**
     * The token workers can use to join the swarm.
     */
    'Worker'?: string;
    /**
     * The token managers can use to join the swarm.
     */
    'Manager'?: string;
  };

  type Swarm = ClusterInfo & {
    'JoinTokens'?: JoinTokens;
  };

  type TaskSpec = {
    /**
     * Plugin spec for the service.  *(Experimental release only.)*
     * 
     * <p><br /></p>
     * 
     * > **Note**: ContainerSpec, NetworkAttachmentSpec, and PluginSpec are
     * > mutually exclusive. PluginSpec is only used when the Runtime field
     * > is set to `plugin`. NetworkAttachmentSpec is used when the Runtime
     * > field is set to `attachment`.
     */
    'PluginSpec'?: {
      /**
       * The name or 'alias' to use for the plugin.
       */
      'Name'?: string;
      /**
       * The plugin image reference to use.
       */
      'Remote'?: string;
      /**
       * Disable the plugin once scheduled.
       */
      'Disabled'?: boolean;
      'PluginPrivilege'?: Array<PluginPrivilege>;
    };
    /**
     * Container spec for the service.
     * 
     * <p><br /></p>
     * 
     * > **Note**: ContainerSpec, NetworkAttachmentSpec, and PluginSpec are
     * > mutually exclusive. PluginSpec is only used when the Runtime field
     * > is set to `plugin`. NetworkAttachmentSpec is used when the Runtime
     * > field is set to `attachment`.
     */
    'ContainerSpec'?: {
      /**
       * The image name to use for the container
       */
      'Image'?: string;
      /**
       * User-defined key/value data.
       */
      'Labels'?: Record<string, string>;
      /**
       * The command to be run in the image.
       */
      'Command'?: Array<string>;
      /**
       * Arguments to the command.
       */
      'Args'?: Array<string>;
      /**
       * The hostname to use for the container, as a valid
       * [RFC 1123](https://tools.ietf.org/html/rfc1123) hostname.
       */
      'Hostname'?: string;
      /**
       * A list of environment variables in the form `VAR=value`.
       */
      'Env'?: Array<string>;
      /**
       * The working directory for commands to run in.
       */
      'Dir'?: string;
      /**
       * The user inside the container.
       */
      'User'?: string;
      /**
       * A list of additional groups that the container process will run as.
       */
      'Groups'?: Array<string>;
      /**
       * Security options for the container
       */
      'Privileges'?: {
        /**
         * CredentialSpec for managed service account (Windows only)
         */
        'CredentialSpec'?: {
          /**
           * Load credential spec from a Swarm Config with the given ID.
           * The specified config must also be present in the Configs
           * field with the Runtime property set.
           * 
           * <p><br /></p>
           * 
           * 
           * > **Note**: `CredentialSpec.File`, `CredentialSpec.Registry`,
           * > and `CredentialSpec.Config` are mutually exclusive.
           */
          'Config'?: string;
          /**
           * Load credential spec from this file. The file is read by
           * the daemon, and must be present in the `CredentialSpecs`
           * subdirectory in the docker data directory, which defaults
           * to `C:\ProgramData\Docker\` on Windows.
           * 
           * For example, specifying `spec.json` loads
           * `C:\ProgramData\Docker\CredentialSpecs\spec.json`.
           * 
           * <p><br /></p>
           * 
           * > **Note**: `CredentialSpec.File`, `CredentialSpec.Registry`,
           * > and `CredentialSpec.Config` are mutually exclusive.
           */
          'File'?: string;
          /**
           * Load credential spec from this value in the Windows
           * registry. The specified registry value must be located in:
           * 
           * `HKLM\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Virtualization\Containers\CredentialSpecs`
           * 
           * <p><br /></p>
           * 
           * 
           * > **Note**: `CredentialSpec.File`, `CredentialSpec.Registry`,
           * > and `CredentialSpec.Config` are mutually exclusive.
           */
          'Registry'?: string;
        };
        /**
         * SELinux labels of the container
         */
        'SELinuxContext'?: {
          /**
           * Disable SELinux
           */
          'Disable'?: boolean;
          /**
           * SELinux user label
           */
          'User'?: string;
          /**
           * SELinux role label
           */
          'Role'?: string;
          /**
           * SELinux type label
           */
          'Type'?: string;
          /**
           * SELinux level label
           */
          'Level'?: string;
        };
      };
      /**
       * Whether a pseudo-TTY should be allocated.
       */
      'TTY'?: boolean;
      /**
       * Open `stdin`
       */
      'OpenStdin'?: boolean;
      /**
       * Mount the container's root filesystem as read only.
       */
      'ReadOnly'?: boolean;
      /**
       * Specification for mounts to be added to containers created as part
       * of the service.
       */
      'Mounts'?: Array<Mount>;
      /**
       * Signal to stop the container.
       */
      'StopSignal'?: string;
      /**
       * Amount of time to wait for the container to terminate before
       * forcefully killing it.
       */
      'StopGracePeriod'?: number;
      'HealthCheck'?: HealthConfig;
      /**
       * A list of hostname/IP mappings to add to the container's `hosts`
       * file. The format of extra hosts is specified in the
       * [hosts(5)](http://man7.org/linux/man-pages/man5/hosts.5.html)
       * man page:
       * 
       *     IP_address canonical_hostname [aliases...]
       */
      'Hosts'?: Array<string>;
      /**
       * Specification for DNS related configurations in resolver configuration
       * file (`resolv.conf`).
       */
      'DNSConfig'?: {
        /**
         * The IP addresses of the name servers.
         */
        'Nameservers'?: Array<string>;
        /**
         * A search list for host-name lookup.
         */
        'Search'?: Array<string>;
        /**
         * A list of internal resolver variables to be modified (e.g.,
         * `debug`, `ndots:3`, etc.).
         */
        'Options'?: Array<string>;
      };
      /**
       * Secrets contains references to zero or more secrets that will be
       * exposed to the service.
       */
      'Secrets'?: Array<{
        /**
         * File represents a specific target that is backed by a file.
         */
        'File'?: {
          /**
           * Name represents the final filename in the filesystem.
           */
          'Name'?: string;
          /**
           * UID represents the file UID.
           */
          'UID'?: string;
          /**
           * GID represents the file GID.
           */
          'GID'?: string;
          /**
           * Mode represents the FileMode of the file.
           */
          'Mode'?: number;
        };
        /**
         * SecretID represents the ID of the specific secret that we're
         * referencing.
         */
        'SecretID'?: string;
        /**
         * SecretName is the name of the secret that this references,
         * but this is just provided for lookup/display purposes. The
         * secret in the reference will be identified by its ID.
         */
        'SecretName'?: string;
      }>;
      /**
       * Configs contains references to zero or more configs that will be
       * exposed to the service.
       */
      'Configs'?: Array<{
        /**
         * File represents a specific target that is backed by a file.
         * 
         * <p><br /><p>
         * 
         * > **Note**: `Configs.File` and `Configs.Runtime` are mutually exclusive
         */
        'File'?: {
          /**
           * Name represents the final filename in the filesystem.
           */
          'Name'?: string;
          /**
           * UID represents the file UID.
           */
          'UID'?: string;
          /**
           * GID represents the file GID.
           */
          'GID'?: string;
          /**
           * Mode represents the FileMode of the file.
           */
          'Mode'?: number;
        };
        /**
         * Runtime represents a target that is not mounted into the
         * container but is used by the task
         * 
         * <p><br /><p>
         * 
         * > **Note**: `Configs.File` and `Configs.Runtime` are mutually
         * > exclusive
         */
        'Runtime'?: {
        };
        /**
         * ConfigID represents the ID of the specific config that we're
         * referencing.
         */
        'ConfigID'?: string;
        /**
         * ConfigName is the name of the config that this references,
         * but this is just provided for lookup/display purposes. The
         * config in the reference will be identified by its ID.
         */
        'ConfigName'?: string;
      }>;
      /**
       * Isolation technology of the containers running the service.
       * (Windows only)
       */
      'Isolation'?: "default" | "process" | "hyperv";
      /**
       * Run an init inside the container that forwards signals and reaps
       * processes. This field is omitted if empty, and the default (as
       * configured on the daemon) is used.
       */
      'Init'?: boolean;
      /**
       * Set kernel namedspaced parameters (sysctls) in the container.
       * The Sysctls option on services accepts the same sysctls as the
       * are supported on containers. Note that while the same sysctls are
       * supported, no guarantees or checks are made about their
       * suitability for a clustered environment, and it's up to the user
       * to determine whether a given sysctl will work properly in a
       * Service.
       */
      'Sysctls'?: Record<string, string>;
      /**
       * A list of kernel capabilities to add to the default set
       * for the container.
       */
      'CapabilityAdd'?: Array<string>;
      /**
       * A list of kernel capabilities to drop from the default set
       * for the container.
       */
      'CapabilityDrop'?: Array<string>;
      /**
       * A list of resource limits to set in the container. For example: `{"Name": "nofile", "Soft": 1024, "Hard": 2048}`"
       */
      'Ulimits'?: Array<{
        /**
         * Name of ulimit
         */
        'Name'?: string;
        /**
         * Soft limit
         */
        'Soft'?: number;
        /**
         * Hard limit
         */
        'Hard'?: number;
      }>;
    };
    /**
     * Read-only spec type for non-swarm containers attached to swarm overlay
     * networks.
     * 
     * <p><br /></p>
     * 
     * > **Note**: ContainerSpec, NetworkAttachmentSpec, and PluginSpec are
     * > mutually exclusive. PluginSpec is only used when the Runtime field
     * > is set to `plugin`. NetworkAttachmentSpec is used when the Runtime
     * > field is set to `attachment`.
     */
    'NetworkAttachmentSpec'?: {
      /**
       * ID of the container represented by this task
       */
      'ContainerID'?: string;
    };
    /**
     * Resource requirements which apply to each individual container created
     * as part of the service.
     */
    'Resources'?: {
      /**
       * Define resources limits.
       */
      'Limits'?: Limit;
      /**
       * Define resources reservation.
       */
      'Reservations'?: ResourceObject;
    };
    /**
     * Specification for the restart policy which applies to containers
     * created as part of this service.
     */
    'RestartPolicy'?: {
      /**
       * Condition for restart.
       */
      'Condition'?: "none" | "on-failure" | "any";
      /**
       * Delay between restart attempts.
       */
      'Delay'?: number;
      /**
       * Maximum attempts to restart a given container before giving up
       * (default value is 0, which is ignored).
       */
      'MaxAttempts'?: number;
      /**
       * Windows is the time window used to evaluate the restart policy
       * (default value is 0, which is unbounded).
       */
      'Window'?: number;
    };
    'Placement'?: {
      /**
       * An array of constraint expressions to limit the set of nodes where
       * a task can be scheduled. Constraint expressions can either use a
       * _match_ (`==`) or _exclude_ (`!=`) rule. Multiple constraints find
       * nodes that satisfy every expression (AND match). Constraints can
       * match node or Docker Engine labels as follows:
       * 
       * node attribute       | matches                        | example
       * ---------------------|--------------------------------|-----------------------------------------------
       * `node.id`            | Node ID                        | `node.id==2ivku8v2gvtg4`
       * `node.hostname`      | Node hostname                  | `node.hostname!=node-2`
       * `node.role`          | Node role (`manager`/`worker`) | `node.role==manager`
       * `node.platform.os`   | Node operating system          | `node.platform.os==windows`
       * `node.platform.arch` | Node architecture              | `node.platform.arch==x86_64`
       * `node.labels`        | User-defined node labels       | `node.labels.security==high`
       * `engine.labels`      | Docker Engine's labels         | `engine.labels.operatingsystem==ubuntu-14.04`
       * 
       * `engine.labels` apply to Docker Engine labels like operating system,
       * drivers, etc. Swarm administrators add `node.labels` for operational
       * purposes by using the [`node update endpoint`](#operation/NodeUpdate).
       */
      'Constraints'?: Array<string>;
      /**
       * Preferences provide a way to make the scheduler aware of factors
       * such as topology. They are provided in order from highest to
       * lowest precedence.
       */
      'Preferences'?: Array<{
        'Spread'?: {
          /**
           * label descriptor, such as `engine.labels.az`.
           */
          'SpreadDescriptor'?: string;
        };
      }>;
      /**
       * Maximum number of replicas for per node (default value is 0, which
       * is unlimited)
       */
      'MaxReplicas'?: number;
      /**
       * Platforms stores all the platforms that the service's image can
       * run on. This field is used in the platform filter for scheduling.
       * If empty, then the platform filter is off, meaning there are no
       * scheduling restrictions.
       */
      'Platforms'?: Array<Platform>;
    };
    /**
     * A counter that triggers an update even if no relevant parameters have
     * been changed.
     */
    'ForceUpdate'?: number;
    /**
     * Runtime is the type of runtime specified for the task executor.
     */
    'Runtime'?: string;
    /**
     * Specifies which networks the service should attach to.
     */
    'Networks'?: Array<NetworkAttachmentConfig>;
    /**
     * Specifies the log driver to use for tasks created from this spec. If
     * not present, the default one for the swarm will be used, finally
     * falling back to the engine default if not specified.
     */
    'LogDriver'?: {
      'Name'?: string;
      'Options'?: Record<string, string>;
    };
  };

  type TaskState = "new" | "allocated" | "pending" | "assigned" | "accepted" | "preparing" | "ready" | "starting" | "running" | "complete" | "shutdown" | "failed" | "rejected" | "remove" | "orphaned";

  type Task = {
    /**
     * The ID of the task.
     */
    'ID'?: string;
    'Version'?: ObjectVersion;
    'CreatedAt'?: string;
    'UpdatedAt'?: string;
    /**
     * Name of the task.
     */
    'Name'?: string;
    /**
     * User-defined key/value metadata.
     */
    'Labels'?: Record<string, string>;
    'Spec'?: TaskSpec;
    /**
     * The ID of the service this task is part of.
     */
    'ServiceID'?: string;
    'Slot'?: number;
    /**
     * The ID of the node that this task is on.
     */
    'NodeID'?: string;
    'AssignedGenericResources'?: GenericResources;
    'Status'?: {
      'Timestamp'?: string;
      'State'?: TaskState;
      'Message'?: string;
      'Err'?: string;
      'ContainerStatus'?: {
        'ContainerID'?: string;
        'PID'?: number;
        'ExitCode'?: number;
      };
    };
    'DesiredState'?: TaskState;
    /**
     * If the Service this Task belongs to is a job-mode service, contains
     * the JobIteration of the Service this Task was created for. Absent if
     * the Task was created for a Replicated or Global Service.
     */
    'JobIteration'?: ObjectVersion;
  };

  type ServiceSpec = {
    /**
     * Name of the service.
     */
    'Name'?: string;
    /**
     * User-defined key/value metadata.
     */
    'Labels'?: Record<string, string>;
    'TaskTemplate'?: TaskSpec;
    /**
     * Scheduling mode for the service.
     */
    'Mode'?: {
      'Replicated'?: {
        'Replicas'?: number;
      };
      'Global'?: {
      };
      /**
       * The mode used for services with a finite number of tasks that run
       * to a completed state.
       */
      'ReplicatedJob'?: {
        /**
         * The maximum number of replicas to run simultaneously.
         */
        'MaxConcurrent'?: number;
        /**
         * The total number of replicas desired to reach the Completed
         * state. If unset, will default to the value of `MaxConcurrent`
         */
        'TotalCompletions'?: number;
      };
      /**
       * The mode used for services which run a task to the completed state
       * on each valid node.
       */
      'GlobalJob'?: {
      };
    };
    /**
     * Specification for the update strategy of the service.
     */
    'UpdateConfig'?: {
      /**
       * Maximum number of tasks to be updated in one iteration (0 means
       * unlimited parallelism).
       */
      'Parallelism'?: number;
      /**
       * Amount of time between updates, in nanoseconds.
       */
      'Delay'?: number;
      /**
       * Action to take if an updated task fails to run, or stops running
       * during the update.
       */
      'FailureAction'?: "continue" | "pause" | "rollback";
      /**
       * Amount of time to monitor each updated task for failures, in
       * nanoseconds.
       */
      'Monitor'?: number;
      /**
       * The fraction of tasks that may fail during an update before the
       * failure action is invoked, specified as a floating point number
       * between 0 and 1.
       */
      'MaxFailureRatio'?: number;
      /**
       * The order of operations when rolling out an updated task. Either
       * the old task is shut down before the new task is started, or the
       * new task is started before the old task is shut down.
       */
      'Order'?: "stop-first" | "start-first";
    };
    /**
     * Specification for the rollback strategy of the service.
     */
    'RollbackConfig'?: {
      /**
       * Maximum number of tasks to be rolled back in one iteration (0 means
       * unlimited parallelism).
       */
      'Parallelism'?: number;
      /**
       * Amount of time between rollback iterations, in nanoseconds.
       */
      'Delay'?: number;
      /**
       * Action to take if an rolled back task fails to run, or stops
       * running during the rollback.
       */
      'FailureAction'?: "continue" | "pause";
      /**
       * Amount of time to monitor each rolled back task for failures, in
       * nanoseconds.
       */
      'Monitor'?: number;
      /**
       * The fraction of tasks that may fail during a rollback before the
       * failure action is invoked, specified as a floating point number
       * between 0 and 1.
       */
      'MaxFailureRatio'?: number;
      /**
       * The order of operations when rolling back a task. Either the old
       * task is shut down before the new task is started, or the new task
       * is started before the old task is shut down.
       */
      'Order'?: "stop-first" | "start-first";
    };
    /**
     * Specifies which networks the service should attach to.
     */
    'Networks'?: Array<NetworkAttachmentConfig>;
    'EndpointSpec'?: EndpointSpec;
  };

  type EndpointPortConfig = {
    'Name'?: string;
    'Protocol'?: "tcp" | "udp" | "sctp";
    /**
     * The port inside the container.
     */
    'TargetPort'?: number;
    /**
     * The port on the swarm hosts.
     */
    'PublishedPort'?: number;
    /**
     * The mode in which port is published.
     * 
     * <p><br /></p>
     * 
     * - "ingress" makes the target port accessible on every node,
     *   regardless of whether there is a task for the service running on
     *   that node or not.
     * - "host" bypasses the routing mesh and publish the port directly on
     *   the swarm node where that service is running.
     */
    'PublishMode'?: "ingress" | "host";
  };

  type EndpointSpec = {
    /**
     * The mode of resolution to use for internal load balancing between tasks.
     */
    'Mode'?: "vip" | "dnsrr";
    /**
     * List of exposed ports that this service is accessible on from the
     * outside. Ports can only be provided if `vip` resolution mode is used.
     */
    'Ports'?: Array<EndpointPortConfig>;
  };

  type Service = {
    'ID'?: string;
    'Version'?: ObjectVersion;
    'CreatedAt'?: string;
    'UpdatedAt'?: string;
    'Spec'?: ServiceSpec;
    'Endpoint'?: {
      'Spec'?: EndpointSpec;
      'Ports'?: Array<EndpointPortConfig>;
      'VirtualIPs'?: Array<{
        'NetworkID'?: string;
        'Addr'?: string;
      }>;
    };
    /**
     * The status of a service update.
     */
    'UpdateStatus'?: {
      'State'?: "updating" | "paused" | "completed";
      'StartedAt'?: string;
      'CompletedAt'?: string;
      'Message'?: string;
    };
    /**
     * The status of the service's tasks. Provided only when requested as
     * part of a ServiceList operation.
     */
    'ServiceStatus'?: {
      /**
       * The number of tasks for the service currently in the Running state.
       */
      'RunningTasks'?: number;
      /**
       * The number of tasks for the service desired to be running.
       * For replicated services, this is the replica count from the
       * service spec. For global services, this is computed by taking
       * count of all tasks for the service with a Desired State other
       * than Shutdown.
       */
      'DesiredTasks'?: number;
      /**
       * The number of tasks for a job that are in the Completed state.
       * This field must be cross-referenced with the service type, as the
       * value of 0 may mean the service is not in a job mode, or it may
       * mean the job-mode service has no tasks yet Completed.
       */
      'CompletedTasks'?: number;
    };
    /**
     * The status of the service when it is in one of ReplicatedJob or
     * GlobalJob modes. Absent on Replicated and Global mode services. The
     * JobIteration is an ObjectVersion, but unlike the Service's version,
     * does not need to be sent with an update request.
     */
    'JobStatus'?: {
      /**
       * JobIteration is a value increased each time a Job is executed,
       * successfully or otherwise. "Executed", in this case, means the
       * job as a whole has been started, not that an individual Task has
       * been launched. A job is "Executed" when its ServiceSpec is
       * updated. JobIteration can be used to disambiguate Tasks belonging
       * to different executions of a job.  Though JobIteration will
       * increase with each subsequent execution, it may not necessarily
       * increase by 1, and so JobIteration should not be used to
       */
      'JobIteration'?: ObjectVersion;
      /**
       * The last time, as observed by the server, that this job was
       * started.
       */
      'LastExecution'?: string;
    };
  };

  type ImageDeleteResponseItem = {
    /**
     * The image ID of an image that was untagged
     */
    'Untagged'?: string;
    /**
     * The image ID of an image that was deleted
     */
    'Deleted'?: string;
  };

  type ServiceUpdateResponse = {
    /**
     * Optional warning messages
     */
    'Warnings'?: Array<string>;
  };

  type ContainerSummary = {
    /**
     * The ID of this container
     */
    'Id'?: string;
    /**
     * The names that this container has been given
     */
    'Names'?: Array<string>;
    /**
     * The name of the image used when creating this container
     */
    'Image'?: string;
    /**
     * The ID of the image that this container was created from
     */
    'ImageID'?: string;
    /**
     * Command to run when starting the container
     */
    'Command'?: string;
    /**
     * When the container was created
     */
    'Created'?: number;
    /**
     * The ports exposed by this container
     */
    'Ports'?: Array<Port>;
    /**
     * The size of files that have been created or changed by this container
     */
    'SizeRw'?: number;
    /**
     * The total size of all the files in this container
     */
    'SizeRootFs'?: number;
    /**
     * User-defined key/value metadata.
     */
    'Labels'?: Record<string, string>;
    /**
     * The state of this container (e.g. `Exited`)
     */
    'State'?: string;
    /**
     * Additional human-readable status of this container (e.g. `Exit 0`)
     */
    'Status'?: string;
    'HostConfig'?: {
      'NetworkMode'?: string;
    };
    /**
     * A summary of the container's network settings
     */
    'NetworkSettings'?: {
      'Networks'?: Record<string, EndpointSettings>;
    };
    'Mounts'?: Array<MountPoint>;
  };

  type Driver = {
    /**
     * Name of the driver.
     */
    'Name'?: string;
    /**
     * Key/value map of driver-specific options.
     */
    'Options'?: Record<string, string>;
  };

  type SecretSpec = {
    /**
     * User-defined name of the secret.
     */
    'Name'?: string;
    /**
     * User-defined key/value metadata.
     */
    'Labels'?: Record<string, string>;
    /**
     * Base64-url-safe-encoded ([RFC 4648](https://tools.ietf.org/html/rfc4648#section-5))
     * data to store as secret.
     * 
     * This field is only used to _create_ a secret, and is not returned by
     * other endpoints.
     */
    'Data'?: string;
    /**
     * Name of the secrets driver used to fetch the secret's value from an
     * external secret store.
     */
    'Driver'?: Driver;
    /**
     * Templating driver, if applicable
     * 
     * Templating controls whether and how to evaluate the config payload as
     * a template. If no driver is set, no templating is used.
     */
    'Templating'?: Driver;
  };

  type Secret = {
    'ID'?: string;
    'Version'?: ObjectVersion;
    'CreatedAt'?: string;
    'UpdatedAt'?: string;
    'Spec'?: SecretSpec;
  };

  type ConfigSpec = {
    /**
     * User-defined name of the config.
     */
    'Name'?: string;
    /**
     * User-defined key/value metadata.
     */
    'Labels'?: Record<string, string>;
    /**
     * Base64-url-safe-encoded ([RFC 4648](https://tools.ietf.org/html/rfc4648#section-5))
     * config data.
     */
    'Data'?: string;
    /**
     * Templating driver, if applicable
     * 
     * Templating controls whether and how to evaluate the config payload as
     * a template. If no driver is set, no templating is used.
     */
    'Templating'?: Driver;
  };

  type Config = {
    'ID'?: string;
    'Version'?: ObjectVersion;
    'CreatedAt'?: string;
    'UpdatedAt'?: string;
    'Spec'?: ConfigSpec;
  };

  type ContainerState = {
    /**
     * String representation of the container state. Can be one of "created",
     * "running", "paused", "restarting", "removing", "exited", or "dead".
     */
    'Status'?: "created" | "running" | "paused" | "restarting" | "removing" | "exited" | "dead";
    /**
     * Whether this container is running.
     * 
     * Note that a running container can be _paused_. The `Running` and `Paused`
     * booleans are not mutually exclusive:
     * 
     * When pausing a container (on Linux), the freezer cgroup is used to suspend
     * all processes in the container. Freezing the process requires the process to
     * be running. As a result, paused containers are both `Running` _and_ `Paused`.
     * 
     * Use the `Status` field instead to determine if a container's state is "running".
     */
    'Running'?: boolean;
    /**
     * Whether this container is paused.
     */
    'Paused'?: boolean;
    /**
     * Whether this container is restarting.
     */
    'Restarting'?: boolean;
    /**
     * Whether this container has been killed because it ran out of memory.
     */
    'OOMKilled'?: boolean;
    'Dead'?: boolean;
    /**
     * The process ID of this container
     */
    'Pid'?: number;
    /**
     * The last exit code of this container
     */
    'ExitCode'?: number;
    'Error'?: string;
    /**
     * The time when this container was last started.
     */
    'StartedAt'?: string;
    /**
     * The time when this container last exited.
     */
    'FinishedAt'?: string;
    'Health'?: Health;
  };

  type ContainerWaitResponse = {
    /**
     * Exit code of the container
     */
    'StatusCode'?: number;
    'Error'?: ContainerWaitExitError;
  };

  type ContainerWaitExitError = {
    /**
     * Details of an error
     */
    'Message'?: string;
  };

  type SystemVersion = {
    'Platform': {
      'Name'?: string;
    };
    /**
     * Information about system components
     */
    'Components'?: Array<{
      /**
       * Name of the component
       */
      'Name'?: string;
      /**
       * Version of the component
       */
      'Version'?: string;
      /**
       * Key/value pairs of strings with additional information about the
       * component. These values are intended for informational purposes
       * only, and their content is not defined, and not part of the API
       * specification.
       * 
       * These messages can be printed by the client as information to the user.
       */
      'Details'?: {
      };
    }>;
    /**
     * The version of the daemon
     */
    'Version'?: string;
    /**
     * The default (and highest) API version that is supported by the daemon
     */
    'ApiVersion'?: string;
    /**
     * The minimum API version that is supported by the daemon
     */
    'MinAPIVersion'?: string;
    /**
     * The Git commit of the source code that was used to build the daemon
     */
    'GitCommit'?: string;
    /**
     * The version Go used to compile the daemon, and the version of the Go
     * runtime in use.
     */
    'GoVersion'?: string;
    /**
     * The operating system that the daemon is running on ("linux" or "windows")
     */
    'Os'?: string;
    /**
     * The architecture that the daemon is running on
     */
    'Arch'?: string;
    /**
     * The kernel version (`uname -r`) that the daemon is running on.
     * 
     * This field is omitted when empty.
     */
    'KernelVersion'?: string;
    /**
     * Indicates if the daemon is started with experimental features enabled.
     * 
     * This field is omitted when empty / false.
     */
    'Experimental'?: boolean;
    /**
     * The date and time that the daemon was compiled.
     */
    'BuildTime'?: string;
  };

  type SystemInfo = {
    /**
     * Unique identifier of the daemon.
     * 
     * <p><br /></p>
     * 
     * > **Note**: The format of the ID itself is not part of the API, and
     * > should not be considered stable.
     */
    'ID'?: string;
    /**
     * Total number of containers on the host.
     */
    'Containers'?: number;
    /**
     * Number of containers with status `"running"`.
     */
    'ContainersRunning'?: number;
    /**
     * Number of containers with status `"paused"`.
     */
    'ContainersPaused'?: number;
    /**
     * Number of containers with status `"stopped"`.
     */
    'ContainersStopped'?: number;
    /**
     * Total number of images on the host.
     * 
     * Both _tagged_ and _untagged_ (dangling) images are counted.
     */
    'Images'?: number;
    /**
     * Name of the storage driver in use.
     */
    'Driver'?: string;
    /**
     * Information specific to the storage driver, provided as
     * "label" / "value" pairs.
     * 
     * This information is provided by the storage driver, and formatted
     * in a way consistent with the output of `docker info` on the command
     * line.
     * 
     * <p><br /></p>
     * 
     * > **Note**: The information returned in this field, including the
     * > formatting of values and labels, should not be considered stable,
     * > and may change without notice.
     */
    'DriverStatus'?: Array<Array<string>>;
    /**
     * Root directory of persistent Docker state.
     * 
     * Defaults to `/var/lib/docker` on Linux, and `C:\ProgramData\docker`
     * on Windows.
     */
    'DockerRootDir'?: string;
    'Plugins'?: PluginsInfo;
    /**
     * Indicates if the host has memory limit support enabled.
     */
    'MemoryLimit'?: boolean;
    /**
     * Indicates if the host has memory swap limit support enabled.
     */
    'SwapLimit'?: boolean;
    /**
     * Indicates if the host has kernel memory limit support enabled.
     * 
     * <p><br /></p>
     * 
     * > **Deprecated**: This field is deprecated as the kernel 5.4 deprecated
     * > `kmem.limit_in_bytes`.
     */
    'KernelMemory'?: boolean;
    /**
     * Indicates if the host has kernel memory TCP limit support enabled.
     * 
     * Kernel memory TCP limits are not supported when using cgroups v2, which
     * does not support the corresponding `memory.kmem.tcp.limit_in_bytes` cgroup.
     */
    'KernelMemoryTCP'?: boolean;
    /**
     * Indicates if CPU CFS(Completely Fair Scheduler) period is supported by
     * the host.
     */
    'CpuCfsPeriod'?: boolean;
    /**
     * Indicates if CPU CFS(Completely Fair Scheduler) quota is supported by
     * the host.
     */
    'CpuCfsQuota'?: boolean;
    /**
     * Indicates if CPU Shares limiting is supported by the host.
     */
    'CPUShares'?: boolean;
    /**
     * Indicates if CPUsets (cpuset.cpus, cpuset.mems) are supported by the host.
     * 
     * See [cpuset(7)](https://www.kernel.org/doc/Documentation/cgroup-v1/cpusets.txt)
     */
    'CPUSet'?: boolean;
    /**
     * Indicates if the host kernel has PID limit support enabled.
     */
    'PidsLimit'?: boolean;
    /**
     * Indicates if OOM killer disable is supported on the host.
     */
    'OomKillDisable'?: boolean;
    /**
     * Indicates IPv4 forwarding is enabled.
     */
    'IPv4Forwarding'?: boolean;
    /**
     * Indicates if `bridge-nf-call-iptables` is available on the host.
     */
    'BridgeNfIptables'?: boolean;
    /**
     * Indicates if `bridge-nf-call-ip6tables` is available on the host.
     */
    'BridgeNfIp6tables'?: boolean;
    /**
     * Indicates if the daemon is running in debug-mode / with debug-level
     * logging enabled.
     */
    'Debug'?: boolean;
    /**
     * The total number of file Descriptors in use by the daemon process.
     * 
     * This information is only returned if debug-mode is enabled.
     */
    'NFd'?: number;
    /**
     * The  number of goroutines that currently exist.
     * 
     * This information is only returned if debug-mode is enabled.
     */
    'NGoroutines'?: number;
    /**
     * Current system-time in [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt)
     * format with nano-seconds.
     */
    'SystemTime'?: string;
    /**
     * The logging driver to use as a default for new containers.
     */
    'LoggingDriver'?: string;
    /**
     * The driver to use for managing cgroups.
     */
    'CgroupDriver'?: "cgroupfs" | "systemd" | "none";
    /**
     * The version of the cgroup.
     */
    'CgroupVersion'?: "1" | "2";
    /**
     * Number of event listeners subscribed.
     */
    'NEventsListener'?: number;
    /**
     * Kernel version of the host.
     * 
     * On Linux, this information obtained from `uname`. On Windows this
     * information is queried from the <kbd>HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion\\</kbd>
     * registry value, for example _"10.0 14393 (14393.1198.amd64fre.rs1_release_sec.170427-1353)"_.
     */
    'KernelVersion'?: string;
    /**
     * Name of the host's operating system, for example: "Ubuntu 16.04.2 LTS"
     * or "Windows Server 2016 Datacenter"
     */
    'OperatingSystem'?: string;
    /**
     * Version of the host's operating system
     * 
     * <p><br /></p>
     * 
     * > **Note**: The information returned in this field, including its
     * > very existence, and the formatting of values, should not be considered
     * > stable, and may change without notice.
     */
    'OSVersion'?: string;
    /**
     * Generic type of the operating system of the host, as returned by the
     * Go runtime (`GOOS`).
     * 
     * Currently returned values are "linux" and "windows". A full list of
     * possible values can be found in the [Go documentation](https://golang.org/doc/install/source#environment).
     */
    'OSType'?: string;
    /**
     * Hardware architecture of the host, as returned by the Go runtime
     * (`GOARCH`).
     * 
     * A full list of possible values can be found in the [Go documentation](https://golang.org/doc/install/source#environment).
     */
    'Architecture'?: string;
    /**
     * The number of logical CPUs usable by the daemon.
     * 
     * The number of available CPUs is checked by querying the operating
     * system when the daemon starts. Changes to operating system CPU
     * allocation after the daemon is started are not reflected.
     */
    'NCPU'?: number;
    /**
     * Total amount of physical memory available on the host, in bytes.
     */
    'MemTotal'?: number;
    /**
     * Address / URL of the index server that is used for image search,
     * and as a default for user authentication for Docker Hub and Docker Cloud.
     */
    'IndexServerAddress'?: string;
    'RegistryConfig'?: RegistryServiceConfig;
    'GenericResources'?: GenericResources;
    /**
     * HTTP-proxy configured for the daemon. This value is obtained from the
     * [`HTTP_PROXY`](https://www.gnu.org/software/wget/manual/html_node/Proxies.html) environment variable.
     * Credentials ([user info component](https://tools.ietf.org/html/rfc3986#section-3.2.1)) in the proxy URL
     * are masked in the API response.
     * 
     * Containers do not automatically inherit this configuration.
     */
    'HttpProxy'?: string;
    /**
     * HTTPS-proxy configured for the daemon. This value is obtained from the
     * [`HTTPS_PROXY`](https://www.gnu.org/software/wget/manual/html_node/Proxies.html) environment variable.
     * Credentials ([user info component](https://tools.ietf.org/html/rfc3986#section-3.2.1)) in the proxy URL
     * are masked in the API response.
     * 
     * Containers do not automatically inherit this configuration.
     */
    'HttpsProxy'?: string;
    /**
     * Comma-separated list of domain extensions for which no proxy should be
     * used. This value is obtained from the [`NO_PROXY`](https://www.gnu.org/software/wget/manual/html_node/Proxies.html)
     * environment variable.
     * 
     * Containers do not automatically inherit this configuration.
     */
    'NoProxy'?: string;
    /**
     * Hostname of the host.
     */
    'Name'?: string;
    /**
     * User-defined labels (key/value metadata) as set on the daemon.
     * 
     * <p><br /></p>
     * 
     * > **Note**: When part of a Swarm, nodes can both have _daemon_ labels,
     * > set through the daemon configuration, and _node_ labels, set from a
     * > manager node in the Swarm. Node labels are not included in this
     * > field. Node labels can be retrieved using the `/nodes/(id)` endpoint
     * > on a manager node in the Swarm.
     */
    'Labels'?: Array<string>;
    /**
     * Indicates if experimental features are enabled on the daemon.
     */
    'ExperimentalBuild'?: boolean;
    /**
     * Version string of the daemon.
     * 
     * > **Note**: the [standalone Swarm API](/swarm/swarm-api/)
     * > returns the Swarm version instead of the daemon  version, for example
     * > `swarm/1.2.8`.
     */
    'ServerVersion'?: string;
    /**
     * URL of the distributed storage backend.
     * 
     * 
     * The storage backend is used for multihost networking (to store
     * network and endpoint information) and by the node discovery mechanism.
     * 
     * <p><br /></p>
     * 
     * > **Deprecated**: This field is only propagated when using standalone Swarm
     * > mode, and overlay networking using an external k/v store. Overlay
     * > networks with Swarm mode enabled use the built-in raft store, and
     * > this field will be empty.
     */
    'ClusterStore'?: string;
    /**
     * The network endpoint that the Engine advertises for the purpose of
     * node discovery. ClusterAdvertise is a `host:port` combination on which
     * the daemon is reachable by other hosts.
     * 
     * <p><br /></p>
     * 
     * > **Deprecated**: This field is only propagated when using standalone Swarm
     * > mode, and overlay networking using an external k/v store. Overlay
     * > networks with Swarm mode enabled use the built-in raft store, and
     * > this field will be empty.
     */
    'ClusterAdvertise'?: string;
    /**
     * List of [OCI compliant](https://github.com/opencontainers/runtime-spec)
     * runtimes configured on the daemon. Keys hold the "name" used to
     * reference the runtime.
     * 
     * The Docker daemon relies on an OCI compliant runtime (invoked via the
     * `containerd` daemon) as its interface to the Linux kernel namespaces,
     * cgroups, and SELinux.
     * 
     * The default runtime is `runc`, and automatically configured. Additional
     * runtimes can be configured by the user and will be listed here.
     */
    'Runtimes'?: Record<string, Runtime>;
    /**
     * Name of the default OCI runtime that is used when starting containers.
     * 
     * The default can be overridden per-container at create time.
     */
    'DefaultRuntime'?: string;
    'Swarm'?: SwarmInfo;
    /**
     * Indicates if live restore is enabled.
     * 
     * If enabled, containers are kept running when the daemon is shutdown
     * or upon daemon start if running containers are detected.
     */
    'LiveRestoreEnabled'?: boolean;
    /**
     * Represents the isolation technology to use as a default for containers.
     * The supported values are platform-specific.
     * 
     * If no isolation value is specified on daemon start, on Windows client,
     * the default is `hyperv`, and on Windows server, the default is `process`.
     * 
     * This option is currently not used on other platforms.
     */
    'Isolation'?: "default" | "hyperv" | "process";
    /**
     * Name and, optional, path of the `docker-init` binary.
     * 
     * If the path is omitted, the daemon searches the host's `$PATH` for the
     * binary and uses the first result.
     */
    'InitBinary'?: string;
    'ContainerdCommit'?: Commit;
    'RuncCommit'?: Commit;
    'InitCommit'?: Commit;
    /**
     * List of security features that are enabled on the daemon, such as
     * apparmor, seccomp, SELinux, user-namespaces (userns), and rootless.
     * 
     * Additional configuration options for each security feature may
     * be present, and are included as a comma-separated list of key/value
     * pairs.
     */
    'SecurityOptions'?: Array<string>;
    /**
     * Reports a summary of the product license on the daemon.
     * 
     * If a commercial license has been applied to the daemon, information
     * such as number of nodes, and expiration are included.
     */
    'ProductLicense'?: string;
    /**
     * List of custom default address pools for local networks, which can be
     * specified in the daemon.json file or dockerd option.
     * 
     * Example: a Base "10.10.0.0/16" with Size 24 will define the set of 256
     * 10.10.[0-255].0/24 address pools.
     */
    'DefaultAddressPools'?: Array<{
      /**
       * The network address in CIDR format
       */
      'Base'?: string;
      /**
       * The network pool size
       */
      'Size'?: number;
    }>;
    /**
     * List of warnings / informational messages about missing features, or
     * issues related to the daemon configuration.
     * 
     * These messages can be printed by the client as information to the user.
     */
    'Warnings'?: Array<string>;
  };

  type PluginsInfo = {
    /**
     * Names of available volume-drivers, and network-driver plugins.
     */
    'Volume'?: Array<string>;
    /**
     * Names of available network-drivers, and network-driver plugins.
     */
    'Network'?: Array<string>;
    /**
     * Names of available authorization plugins.
     */
    'Authorization'?: Array<string>;
    /**
     * Names of available logging-drivers, and logging-driver plugins.
     */
    'Log'?: Array<string>;
  };

  type RegistryServiceConfig = {
    /**
     * List of IP ranges to which nondistributable artifacts can be pushed,
     * using the CIDR syntax [RFC 4632](https://tools.ietf.org/html/4632).
     * 
     * Some images (for example, Windows base images) contain artifacts
     * whose distribution is restricted by license. When these images are
     * pushed to a registry, restricted artifacts are not included.
     * 
     * This configuration override this behavior, and enables the daemon to
     * push nondistributable artifacts to all registries whose resolved IP
     * address is within the subnet described by the CIDR syntax.
     * 
     * This option is useful when pushing images containing
     * nondistributable artifacts to a registry on an air-gapped network so
     * hosts on that network can pull the images without connecting to
     * another server.
     * 
     * > **Warning**: Nondistributable artifacts typically have restrictions
     * > on how and where they can be distributed and shared. Only use this
     * > feature to push artifacts to private registries and ensure that you
     * > are in compliance with any terms that cover redistributing
     * > nondistributable artifacts.
     */
    'AllowNondistributableArtifactsCIDRs'?: Array<string>;
    /**
     * List of registry hostnames to which nondistributable artifacts can be
     * pushed, using the format `<hostname>[:<port>]` or `<IP address>[:<port>]`.
     * 
     * Some images (for example, Windows base images) contain artifacts
     * whose distribution is restricted by license. When these images are
     * pushed to a registry, restricted artifacts are not included.
     * 
     * This configuration override this behavior for the specified
     * registries.
     * 
     * This option is useful when pushing images containing
     * nondistributable artifacts to a registry on an air-gapped network so
     * hosts on that network can pull the images without connecting to
     * another server.
     * 
     * > **Warning**: Nondistributable artifacts typically have restrictions
     * > on how and where they can be distributed and shared. Only use this
     * > feature to push artifacts to private registries and ensure that you
     * > are in compliance with any terms that cover redistributing
     * > nondistributable artifacts.
     */
    'AllowNondistributableArtifactsHostnames'?: Array<string>;
    /**
     * List of IP ranges of insecure registries, using the CIDR syntax
     * ([RFC 4632](https://tools.ietf.org/html/4632)). Insecure registries
     * accept un-encrypted (HTTP) and/or untrusted (HTTPS with certificates
     * from unknown CAs) communication.
     * 
     * By default, local registries (`127.0.0.0/8`) are configured as
     * insecure. All other registries are secure. Communicating with an
     * insecure registry is not possible if the daemon assumes that registry
     * is secure.
     * 
     * This configuration override this behavior, insecure communication with
     * registries whose resolved IP address is within the subnet described by
     * the CIDR syntax.
     * 
     * Registries can also be marked insecure by hostname. Those registries
     * are listed under `IndexConfigs` and have their `Secure` field set to
     * `false`.
     * 
     * > **Warning**: Using this option can be useful when running a local
     * > registry, but introduces security vulnerabilities. This option
     * > should therefore ONLY be used for testing purposes. For increased
     * > security, users should add their CA to their system's list of trusted
     * > CAs instead of enabling this option.
     */
    'InsecureRegistryCIDRs'?: Array<string>;
    'IndexConfigs'?: Record<string, IndexInfo>;
    /**
     * List of registry URLs that act as a mirror for the official
     * (`docker.io`) registry.
     */
    'Mirrors'?: Array<string>;
  };

  type IndexInfo = {
    /**
     * Name of the registry, such as "docker.io".
     */
    'Name'?: string;
    /**
     * List of mirrors, expressed as URIs.
     */
    'Mirrors'?: Array<string>;
    /**
     * Indicates if the registry is part of the list of insecure
     * registries.
     * 
     * If `false`, the registry is insecure. Insecure registries accept
     * un-encrypted (HTTP) and/or untrusted (HTTPS with certificates from
     * unknown CAs) communication.
     * 
     * > **Warning**: Insecure registries can be useful when running a local
     * > registry. However, because its use creates security vulnerabilities
     * > it should ONLY be enabled for testing purposes. For increased
     * > security, users should add their CA to their system's list of
     * > trusted CAs instead of enabling this option.
     */
    'Secure'?: boolean;
    /**
     * Indicates whether this is an official registry (i.e., Docker Hub / docker.io)
     */
    'Official'?: boolean;
  };

  type Runtime = {
    /**
     * Name and, optional, path, of the OCI executable binary.
     * 
     * If the path is omitted, the daemon searches the host's `$PATH` for the
     * binary and uses the first result.
     */
    'path'?: string;
    /**
     * List of command-line arguments to pass to the runtime when invoked.
     */
    'runtimeArgs'?: Array<string>;
  };

  type Commit = {
    /**
     * Actual commit ID of external tool.
     */
    'ID'?: string;
    /**
     * Commit ID of external tool expected by dockerd as set at build time.
     */
    'Expected'?: string;
  };

  type SwarmInfo = {
    /**
     * Unique identifier of for this node in the swarm.
     */
    'NodeID'?: string;
    /**
     * IP address at which this node can be reached by other nodes in the
     * swarm.
     */
    'NodeAddr'?: string;
    'LocalNodeState'?: LocalNodeState;
    'ControlAvailable'?: boolean;
    'Error'?: string;
    /**
     * List of ID's and addresses of other managers in the swarm.
     */
    'RemoteManagers'?: Array<PeerNode>;
    /**
     * Total number of nodes in the swarm.
     */
    'Nodes'?: number;
    /**
     * Total number of managers in the swarm.
     */
    'Managers'?: number;
    'Cluster'?: ClusterInfo;
  };

  type LocalNodeState = "" | "inactive" | "pending" | "active" | "error" | "locked";

  type PeerNode = {
    /**
     * Unique identifier of for this node in the swarm.
     */
    'NodeID'?: string;
    /**
     * IP address and ports at which this node can be reached.
     */
    'Addr'?: string;
  };

  type NetworkAttachmentConfig = {
    /**
     * The target network for attachment. Must be a network name or ID.
     */
    'Target'?: string;
    /**
     * Discoverable alternate names for the service on this network.
     */
    'Aliases'?: Array<string>;
    /**
     * Driver attachment options for the network target.
     */
    'DriverOpts'?: Record<string, string>;
  };

  type EventActor = {
    /**
     * The ID of the object emitting the event
     */
    'ID'?: string;
    /**
     * Various key/value attributes of the object, depending on its type.
     */
    'Attributes'?: Record<string, string>;
  };

  type EventMessage = {
    /**
     * The type of object emitting the event
     */
    'Type'?: "builder" | "config" | "container" | "daemon" | "image" | "network" | "node" | "plugin" | "secret" | "service" | "volume";
    /**
     * The type of event
     */
    'Action'?: string;
    'Actor'?: EventActor;
    /**
     * Scope of the event. Engine events are `local` scope. Cluster (Swarm)
     * events are `swarm` scope.
     */
    'scope'?: "local" | "swarm";
    /**
     * Timestamp of event
     */
    'time'?: number;
    /**
     * Timestamp of event, with nanosecond accuracy
     */
    'timeNano'?: number;
  };

  type OCIDescriptor = {
    /**
     * The media type of the object this schema refers to.
     */
    'mediaType'?: string;
    /**
     * The digest of the targeted content.
     */
    'digest'?: string;
    /**
     * The size in bytes of the blob.
     */
    'size'?: number;
  };

  type OCIPlatform = {
    /**
     * The CPU architecture, for example `amd64` or `ppc64`.
     */
    'architecture'?: string;
    /**
     * The operating system, for example `linux` or `windows`.
     */
    'os'?: string;
    /**
     * Optional field specifying the operating system version, for example on
     * Windows `10.0.19041.1165`.
     */
    'os.version'?: string;
    /**
     * Optional field specifying an array of strings, each listing a required
     * OS feature (for example on Windows `win32k`).
     */
    'os.features'?: Array<string>;
    /**
     * Optional field specifying a variant of the CPU, for example `v7` to
     * specify ARMv7 when architecture is `arm`.
     */
    'variant'?: string;
  };

  type DistributionInspect = {
    'Descriptor'?: OCIDescriptor;
    /**
     * An array containing all platforms supported by the image.
     */
    'Platforms'?: Array<OCIPlatform>;
  };

  /**
   * `GET /containers/json`
   * 
   * Returns a list of containers. For details on the format, see the
   * [inspect endpoint](#operation/ContainerInspect).
   * 
   * Note that it uses a different, smaller representation of a container
   * than inspecting a single container. For example, the list of linked
   * containers is not propagated .
   * 
   * Code 200: no error
   */
  type GetContainerListResponse200 = Array<ContainerSummary>;

  /**
   * `GET /containers/json`
   * 
   * Returns a list of containers. For details on the format, see the
   * [inspect endpoint](#operation/ContainerInspect).
   * 
   * Note that it uses a different, smaller representation of a container
   * than inspecting a single container. For example, the list of linked
   * containers is not propagated .
   * 
   * Code 400: bad parameter
   */
  type GetContainerListResponse400 = ErrorResponse;

  /**
   * `GET /containers/json`
   * 
   * Returns a list of containers. For details on the format, see the
   * [inspect endpoint](#operation/ContainerInspect).
   * 
   * Note that it uses a different, smaller representation of a container
   * than inspecting a single container. For example, the list of linked
   * containers is not propagated .
   * 
   * Code 500: server error
   */
  type GetContainerListResponse500 = ErrorResponse;

  /**
   * `POST /containers/create`
   * 
   * Code 201: Container created successfully
   */
  type PostContainerCreateResponse201 = {
    /**
     * The ID of the created container
     */
    'Id'?: string;
    /**
     * Warnings encountered when creating the container
     */
    'Warnings'?: Array<string>;
  };

  /**
   * `POST /containers/create`
   * 
   * Code 400: bad parameter
   */
  type PostContainerCreateResponse400 = ErrorResponse;

  /**
   * `POST /containers/create`
   * 
   * Code 404: no such image
   */
  type PostContainerCreateResponse404 = ErrorResponse;

  /**
   * `POST /containers/create`
   * 
   * Code 409: conflict
   */
  type PostContainerCreateResponse409 = ErrorResponse;

  /**
   * `POST /containers/create`
   * 
   * Code 500: server error
   */
  type PostContainerCreateResponse500 = ErrorResponse;

  /**
   * `GET /containers/{id}/json`
   * 
   * Return low-level information about a container.
   * 
   * Code 200: no error
   */
  type GetContainerInspectResponse200 = {
    /**
     * The ID of the container
     */
    'Id'?: string;
    /**
     * The time the container was created
     */
    'Created'?: string;
    /**
     * The path to the command being run
     */
    'Path'?: string;
    /**
     * The arguments to the command being run
     */
    'Args'?: Array<string>;
    'State'?: ContainerState;
    /**
     * The container's image ID
     */
    'Image'?: string;
    'ResolvConfPath'?: string;
    'HostnamePath'?: string;
    'HostsPath'?: string;
    'LogPath'?: string;
    'Name'?: string;
    'RestartCount'?: number;
    'Driver'?: string;
    'Platform'?: string;
    'MountLabel'?: string;
    'ProcessLabel'?: string;
    'AppArmorProfile'?: string;
    /**
     * IDs of exec instances that are running in the container.
     */
    'ExecIDs'?: Array<string>;
    'HostConfig'?: HostConfig;
    'GraphDriver'?: GraphDriverData;
    /**
     * The size of files that have been created or changed by this
     * container.
     */
    'SizeRw'?: number;
    /**
     * The total size of all the files in this container.
     */
    'SizeRootFs'?: number;
    'Mounts'?: Array<MountPoint>;
    'Config'?: ContainerConfig;
    'NetworkSettings'?: NetworkSettings;
  };

  /**
   * `GET /containers/{id}/json`
   * 
   * Return low-level information about a container.
   * 
   * Code 404: no such container
   */
  type GetContainerInspectResponse404 = ErrorResponse;

  /**
   * `GET /containers/{id}/json`
   * 
   * Return low-level information about a container.
   * 
   * Code 500: server error
   */
  type GetContainerInspectResponse500 = ErrorResponse;

  /**
   * `GET /containers/{id}/top`
   * 
   * On Unix systems, this is done by running the `ps` command. This endpoint
   * is not supported on Windows.
   * 
   * Code 200: no error
   */
  type GetContainerTopResponse200 = {
    /**
     * The ps column titles
     */
    'Titles'?: Array<string>;
    /**
     * Each process running in the container, where each is process
     * is an array of values corresponding to the titles.
     */
    'Processes'?: Array<Array<string>>;
  };

  /**
   * `GET /containers/{id}/top`
   * 
   * On Unix systems, this is done by running the `ps` command. This endpoint
   * is not supported on Windows.
   * 
   * Code 404: no such container
   */
  type GetContainerTopResponse404 = ErrorResponse;

  /**
   * `GET /containers/{id}/top`
   * 
   * On Unix systems, this is done by running the `ps` command. This endpoint
   * is not supported on Windows.
   * 
   * Code 500: server error
   */
  type GetContainerTopResponse500 = ErrorResponse;

  /**
   * `GET /containers/{id}/logs`
   * 
   * Get `stdout` and `stderr` logs from a container.
   * 
   * Note: This endpoint works only for containers with the `json-file` or
   * `journald` logging driver.
   * 
   * Code 200: logs returned as a stream in response body.
   * For the stream format, [see the documentation for the attach endpoint](#operation/ContainerAttach).
   * Note that unlike the attach endpoint, the logs endpoint does not
   * upgrade the connection and does not set Content-Type.
   */
  type GetContainerLogsResponse200 = string;

  /**
   * `GET /containers/{id}/logs`
   * 
   * Get `stdout` and `stderr` logs from a container.
   * 
   * Note: This endpoint works only for containers with the `json-file` or
   * `journald` logging driver.
   * 
   * Code 404: no such container
   */
  type GetContainerLogsResponse404 = ErrorResponse;

  /**
   * `GET /containers/{id}/logs`
   * 
   * Get `stdout` and `stderr` logs from a container.
   * 
   * Note: This endpoint works only for containers with the `json-file` or
   * `journald` logging driver.
   * 
   * Code 500: server error
   */
  type GetContainerLogsResponse500 = ErrorResponse;

  /**
   * `GET /containers/{id}/changes`
   * 
   * Returns which files in a container's filesystem have been added, deleted,
   * or modified. The `Kind` of modification can be one of:
   * 
   * - `0`: Modified
   * - `1`: Added
   * - `2`: Deleted
   * 
   * Code 200: The list of changes
   */
  type GetContainerChangesResponse200 = Array<{
    /**
     * Path to file that has changed
     */
    'Path'?: string;
    /**
     * Kind of change
     */
    'Kind'?: 0 | 1 | 2;
  }>;

  /**
   * `GET /containers/{id}/changes`
   * 
   * Returns which files in a container's filesystem have been added, deleted,
   * or modified. The `Kind` of modification can be one of:
   * 
   * - `0`: Modified
   * - `1`: Added
   * - `2`: Deleted
   * 
   * Code 404: no such container
   */
  type GetContainerChangesResponse404 = ErrorResponse;

  /**
   * `GET /containers/{id}/changes`
   * 
   * Returns which files in a container's filesystem have been added, deleted,
   * or modified. The `Kind` of modification can be one of:
   * 
   * - `0`: Modified
   * - `1`: Added
   * - `2`: Deleted
   * 
   * Code 500: server error
   */
  type GetContainerChangesResponse500 = ErrorResponse;

  /**
   * `GET /containers/{id}/export`
   * 
   * Export the contents of a container as a tarball.
   * 
   * Code 404: no such container
   */
  type GetContainerExportResponse404 = ErrorResponse;

  /**
   * `GET /containers/{id}/export`
   * 
   * Export the contents of a container as a tarball.
   * 
   * Code 500: server error
   */
  type GetContainerExportResponse500 = ErrorResponse;

  /**
   * `GET /containers/{id}/stats`
   * 
   * This endpoint returns a live stream of a container’s resource usage
   * statistics.
   * 
   * The `precpu_stats` is the CPU statistic of the *previous* read, and is
   * used to calculate the CPU usage percentage. It is not an exact copy
   * of the `cpu_stats` field.
   * 
   * If either `precpu_stats.online_cpus` or `cpu_stats.online_cpus` is
   * nil then for compatibility with older daemons the length of the
   * corresponding `cpu_usage.percpu_usage` array should be used.
   * 
   * On a cgroup v2 host, the following fields are not set
   * * `blkio_stats`: all fields other than `io_service_bytes_recursive`
   * * `cpu_stats`: `cpu_usage.percpu_usage`
   * * `memory_stats`: `max_usage` and `failcnt`
   * Also, `memory_stats.stats` fields are incompatible with cgroup v1.
   * 
   * To calculate the values shown by the `stats` command of the docker cli tool
   * the following formulas can be used:
   * * used_memory = `memory_stats.usage - memory_stats.stats.cache`
   * * available_memory = `memory_stats.limit`
   * * Memory usage % = `(used_memory / available_memory) * 100.0`
   * * cpu_delta = `cpu_stats.cpu_usage.total_usage - precpu_stats.cpu_usage.total_usage`
   * * system_cpu_delta = `cpu_stats.system_cpu_usage - precpu_stats.system_cpu_usage`
   * * number_cpus = `lenght(cpu_stats.cpu_usage.percpu_usage)` or `cpu_stats.online_cpus`
   * * CPU usage % = `(cpu_delta / system_cpu_delta) * number_cpus * 100.0`
   * 
   * Code 200: no error
   */
  type GetContainerStatsResponse200 = {
  };

  /**
   * `GET /containers/{id}/stats`
   * 
   * This endpoint returns a live stream of a container’s resource usage
   * statistics.
   * 
   * The `precpu_stats` is the CPU statistic of the *previous* read, and is
   * used to calculate the CPU usage percentage. It is not an exact copy
   * of the `cpu_stats` field.
   * 
   * If either `precpu_stats.online_cpus` or `cpu_stats.online_cpus` is
   * nil then for compatibility with older daemons the length of the
   * corresponding `cpu_usage.percpu_usage` array should be used.
   * 
   * On a cgroup v2 host, the following fields are not set
   * * `blkio_stats`: all fields other than `io_service_bytes_recursive`
   * * `cpu_stats`: `cpu_usage.percpu_usage`
   * * `memory_stats`: `max_usage` and `failcnt`
   * Also, `memory_stats.stats` fields are incompatible with cgroup v1.
   * 
   * To calculate the values shown by the `stats` command of the docker cli tool
   * the following formulas can be used:
   * * used_memory = `memory_stats.usage - memory_stats.stats.cache`
   * * available_memory = `memory_stats.limit`
   * * Memory usage % = `(used_memory / available_memory) * 100.0`
   * * cpu_delta = `cpu_stats.cpu_usage.total_usage - precpu_stats.cpu_usage.total_usage`
   * * system_cpu_delta = `cpu_stats.system_cpu_usage - precpu_stats.system_cpu_usage`
   * * number_cpus = `lenght(cpu_stats.cpu_usage.percpu_usage)` or `cpu_stats.online_cpus`
   * * CPU usage % = `(cpu_delta / system_cpu_delta) * number_cpus * 100.0`
   * 
   * Code 404: no such container
   */
  type GetContainerStatsResponse404 = ErrorResponse;

  /**
   * `GET /containers/{id}/stats`
   * 
   * This endpoint returns a live stream of a container’s resource usage
   * statistics.
   * 
   * The `precpu_stats` is the CPU statistic of the *previous* read, and is
   * used to calculate the CPU usage percentage. It is not an exact copy
   * of the `cpu_stats` field.
   * 
   * If either `precpu_stats.online_cpus` or `cpu_stats.online_cpus` is
   * nil then for compatibility with older daemons the length of the
   * corresponding `cpu_usage.percpu_usage` array should be used.
   * 
   * On a cgroup v2 host, the following fields are not set
   * * `blkio_stats`: all fields other than `io_service_bytes_recursive`
   * * `cpu_stats`: `cpu_usage.percpu_usage`
   * * `memory_stats`: `max_usage` and `failcnt`
   * Also, `memory_stats.stats` fields are incompatible with cgroup v1.
   * 
   * To calculate the values shown by the `stats` command of the docker cli tool
   * the following formulas can be used:
   * * used_memory = `memory_stats.usage - memory_stats.stats.cache`
   * * available_memory = `memory_stats.limit`
   * * Memory usage % = `(used_memory / available_memory) * 100.0`
   * * cpu_delta = `cpu_stats.cpu_usage.total_usage - precpu_stats.cpu_usage.total_usage`
   * * system_cpu_delta = `cpu_stats.system_cpu_usage - precpu_stats.system_cpu_usage`
   * * number_cpus = `lenght(cpu_stats.cpu_usage.percpu_usage)` or `cpu_stats.online_cpus`
   * * CPU usage % = `(cpu_delta / system_cpu_delta) * number_cpus * 100.0`
   * 
   * Code 500: server error
   */
  type GetContainerStatsResponse500 = ErrorResponse;

  /**
   * `POST /containers/{id}/resize`
   * 
   * Resize the TTY for a container.
   * 
   * Code 404: no such container
   */
  type PostContainerResizeResponse404 = ErrorResponse;

  /**
   * `POST /containers/{id}/resize`
   * 
   * Resize the TTY for a container.
   * 
   * Code 500: cannot resize container
   */
  type PostContainerResizeResponse500 = ErrorResponse;

  /**
   * `POST /containers/{id}/start`
   * 
   * Code 404: no such container
   */
  type PostContainerStartResponse404 = ErrorResponse;

  /**
   * `POST /containers/{id}/start`
   * 
   * Code 500: server error
   */
  type PostContainerStartResponse500 = ErrorResponse;

  /**
   * `POST /containers/{id}/stop`
   * 
   * Code 404: no such container
   */
  type PostContainerStopResponse404 = ErrorResponse;

  /**
   * `POST /containers/{id}/stop`
   * 
   * Code 500: server error
   */
  type PostContainerStopResponse500 = ErrorResponse;

  /**
   * `POST /containers/{id}/restart`
   * 
   * Code 404: no such container
   */
  type PostContainerRestartResponse404 = ErrorResponse;

  /**
   * `POST /containers/{id}/restart`
   * 
   * Code 500: server error
   */
  type PostContainerRestartResponse500 = ErrorResponse;

  /**
   * `POST /containers/{id}/kill`
   * 
   * Send a POSIX signal to a container, defaulting to killing to the
   * container.
   * 
   * Code 404: no such container
   */
  type PostContainerKillResponse404 = ErrorResponse;

  /**
   * `POST /containers/{id}/kill`
   * 
   * Send a POSIX signal to a container, defaulting to killing to the
   * container.
   * 
   * Code 409: container is not running
   */
  type PostContainerKillResponse409 = ErrorResponse;

  /**
   * `POST /containers/{id}/kill`
   * 
   * Send a POSIX signal to a container, defaulting to killing to the
   * container.
   * 
   * Code 500: server error
   */
  type PostContainerKillResponse500 = ErrorResponse;

  /**
   * `POST /containers/{id}/update`
   * 
   * Change various configuration options of a container without having to
   * recreate it.
   * 
   * Code 200: The container has been updated.
   */
  type PostContainerUpdateResponse200 = {
    'Warnings'?: Array<string>;
  };

  /**
   * `POST /containers/{id}/update`
   * 
   * Change various configuration options of a container without having to
   * recreate it.
   * 
   * Code 404: no such container
   */
  type PostContainerUpdateResponse404 = ErrorResponse;

  /**
   * `POST /containers/{id}/update`
   * 
   * Change various configuration options of a container without having to
   * recreate it.
   * 
   * Code 500: server error
   */
  type PostContainerUpdateResponse500 = ErrorResponse;

  /**
   * `POST /containers/{id}/rename`
   * 
   * Code 404: no such container
   */
  type PostContainerRenameResponse404 = ErrorResponse;

  /**
   * `POST /containers/{id}/rename`
   * 
   * Code 409: name already in use
   */
  type PostContainerRenameResponse409 = ErrorResponse;

  /**
   * `POST /containers/{id}/rename`
   * 
   * Code 500: server error
   */
  type PostContainerRenameResponse500 = ErrorResponse;

  /**
   * `POST /containers/{id}/pause`
   * 
   * Use the freezer cgroup to suspend all processes in a container.
   * 
   * Traditionally, when suspending a process the `SIGSTOP` signal is used,
   * which is observable by the process being suspended. With the freezer
   * cgroup the process is unaware, and unable to capture, that it is being
   * suspended, and subsequently resumed.
   * 
   * Code 404: no such container
   */
  type PostContainerPauseResponse404 = ErrorResponse;

  /**
   * `POST /containers/{id}/pause`
   * 
   * Use the freezer cgroup to suspend all processes in a container.
   * 
   * Traditionally, when suspending a process the `SIGSTOP` signal is used,
   * which is observable by the process being suspended. With the freezer
   * cgroup the process is unaware, and unable to capture, that it is being
   * suspended, and subsequently resumed.
   * 
   * Code 500: server error
   */
  type PostContainerPauseResponse500 = ErrorResponse;

  /**
   * `POST /containers/{id}/unpause`
   * 
   * Resume a container which has been paused.
   * 
   * Code 404: no such container
   */
  type PostContainerUnpauseResponse404 = ErrorResponse;

  /**
   * `POST /containers/{id}/unpause`
   * 
   * Resume a container which has been paused.
   * 
   * Code 500: server error
   */
  type PostContainerUnpauseResponse500 = ErrorResponse;

  /**
   * `POST /containers/{id}/attach`
   * 
   * Attach to a container to read its output or send it input. You can attach
   * to the same container multiple times and you can reattach to containers
   * that have been detached.
   * 
   * Either the `stream` or `logs` parameter must be `true` for this endpoint
   * to do anything.
   * 
   * See the [documentation for the `docker attach` command](/engine/reference/commandline/attach/)
   * for more details.
   * 
   * ### Hijacking
   * 
   * This endpoint hijacks the HTTP connection to transport `stdin`, `stdout`,
   * and `stderr` on the same socket.
   * 
   * This is the response from the daemon for an attach request:
   * 
   * ```
   * HTTP/1.1 200 OK
   * Content-Type: application/vnd.docker.raw-stream
   * 
   * [STREAM]
   * ```
   * 
   * After the headers and two new lines, the TCP connection can now be used
   * for raw, bidirectional communication between the client and server.
   * 
   * To hint potential proxies about connection hijacking, the Docker client
   * can also optionally send connection upgrade headers.
   * 
   * For example, the client sends this request to upgrade the connection:
   * 
   * ```
   * POST /containers/16253994b7c4/attach?stream=1&stdout=1 HTTP/1.1
   * Upgrade: tcp
   * Connection: Upgrade
   * ```
   * 
   * The Docker daemon will respond with a `101 UPGRADED` response, and will
   * similarly follow with the raw stream:
   * 
   * ```
   * HTTP/1.1 101 UPGRADED
   * Content-Type: application/vnd.docker.raw-stream
   * Connection: Upgrade
   * Upgrade: tcp
   * 
   * [STREAM]
   * ```
   * 
   * ### Stream format
   * 
   * When the TTY setting is disabled in [`POST /containers/create`](#operation/ContainerCreate),
   * the stream over the hijacked connected is multiplexed to separate out
   * `stdout` and `stderr`. The stream consists of a series of frames, each
   * containing a header and a payload.
   * 
   * The header contains the information which the stream writes (`stdout` or
   * `stderr`). It also contains the size of the associated frame encoded in
   * the last four bytes (`uint32`).
   * 
   * It is encoded on the first eight bytes like this:
   * 
   * ```go
   * header := [8]byte{STREAM_TYPE, 0, 0, 0, SIZE1, SIZE2, SIZE3, SIZE4}
   * ```
   * 
   * `STREAM_TYPE` can be:
   * 
   * - 0: `stdin` (is written on `stdout`)
   * - 1: `stdout`
   * - 2: `stderr`
   * 
   * `SIZE1, SIZE2, SIZE3, SIZE4` are the four bytes of the `uint32` size
   * encoded as big endian.
   * 
   * Following the header is the payload, which is the specified number of
   * bytes of `STREAM_TYPE`.
   * 
   * The simplest way to implement this protocol is the following:
   * 
   * 1. Read 8 bytes.
   * 2. Choose `stdout` or `stderr` depending on the first byte.
   * 3. Extract the frame size from the last four bytes.
   * 4. Read the extracted size and output it on the correct output.
   * 5. Goto 1.
   * 
   * ### Stream format when using a TTY
   * 
   * When the TTY setting is enabled in [`POST /containers/create`](#operation/ContainerCreate),
   * the stream is not multiplexed. The data exchanged over the hijacked
   * connection is simply the raw data from the process PTY and client's
   * `stdin`.
   * 
   * Code 400: bad parameter
   */
  type PostContainerAttachResponse400 = ErrorResponse;

  /**
   * `POST /containers/{id}/attach`
   * 
   * Attach to a container to read its output or send it input. You can attach
   * to the same container multiple times and you can reattach to containers
   * that have been detached.
   * 
   * Either the `stream` or `logs` parameter must be `true` for this endpoint
   * to do anything.
   * 
   * See the [documentation for the `docker attach` command](/engine/reference/commandline/attach/)
   * for more details.
   * 
   * ### Hijacking
   * 
   * This endpoint hijacks the HTTP connection to transport `stdin`, `stdout`,
   * and `stderr` on the same socket.
   * 
   * This is the response from the daemon for an attach request:
   * 
   * ```
   * HTTP/1.1 200 OK
   * Content-Type: application/vnd.docker.raw-stream
   * 
   * [STREAM]
   * ```
   * 
   * After the headers and two new lines, the TCP connection can now be used
   * for raw, bidirectional communication between the client and server.
   * 
   * To hint potential proxies about connection hijacking, the Docker client
   * can also optionally send connection upgrade headers.
   * 
   * For example, the client sends this request to upgrade the connection:
   * 
   * ```
   * POST /containers/16253994b7c4/attach?stream=1&stdout=1 HTTP/1.1
   * Upgrade: tcp
   * Connection: Upgrade
   * ```
   * 
   * The Docker daemon will respond with a `101 UPGRADED` response, and will
   * similarly follow with the raw stream:
   * 
   * ```
   * HTTP/1.1 101 UPGRADED
   * Content-Type: application/vnd.docker.raw-stream
   * Connection: Upgrade
   * Upgrade: tcp
   * 
   * [STREAM]
   * ```
   * 
   * ### Stream format
   * 
   * When the TTY setting is disabled in [`POST /containers/create`](#operation/ContainerCreate),
   * the stream over the hijacked connected is multiplexed to separate out
   * `stdout` and `stderr`. The stream consists of a series of frames, each
   * containing a header and a payload.
   * 
   * The header contains the information which the stream writes (`stdout` or
   * `stderr`). It also contains the size of the associated frame encoded in
   * the last four bytes (`uint32`).
   * 
   * It is encoded on the first eight bytes like this:
   * 
   * ```go
   * header := [8]byte{STREAM_TYPE, 0, 0, 0, SIZE1, SIZE2, SIZE3, SIZE4}
   * ```
   * 
   * `STREAM_TYPE` can be:
   * 
   * - 0: `stdin` (is written on `stdout`)
   * - 1: `stdout`
   * - 2: `stderr`
   * 
   * `SIZE1, SIZE2, SIZE3, SIZE4` are the four bytes of the `uint32` size
   * encoded as big endian.
   * 
   * Following the header is the payload, which is the specified number of
   * bytes of `STREAM_TYPE`.
   * 
   * The simplest way to implement this protocol is the following:
   * 
   * 1. Read 8 bytes.
   * 2. Choose `stdout` or `stderr` depending on the first byte.
   * 3. Extract the frame size from the last four bytes.
   * 4. Read the extracted size and output it on the correct output.
   * 5. Goto 1.
   * 
   * ### Stream format when using a TTY
   * 
   * When the TTY setting is enabled in [`POST /containers/create`](#operation/ContainerCreate),
   * the stream is not multiplexed. The data exchanged over the hijacked
   * connection is simply the raw data from the process PTY and client's
   * `stdin`.
   * 
   * Code 404: no such container
   */
  type PostContainerAttachResponse404 = ErrorResponse;

  /**
   * `POST /containers/{id}/attach`
   * 
   * Attach to a container to read its output or send it input. You can attach
   * to the same container multiple times and you can reattach to containers
   * that have been detached.
   * 
   * Either the `stream` or `logs` parameter must be `true` for this endpoint
   * to do anything.
   * 
   * See the [documentation for the `docker attach` command](/engine/reference/commandline/attach/)
   * for more details.
   * 
   * ### Hijacking
   * 
   * This endpoint hijacks the HTTP connection to transport `stdin`, `stdout`,
   * and `stderr` on the same socket.
   * 
   * This is the response from the daemon for an attach request:
   * 
   * ```
   * HTTP/1.1 200 OK
   * Content-Type: application/vnd.docker.raw-stream
   * 
   * [STREAM]
   * ```
   * 
   * After the headers and two new lines, the TCP connection can now be used
   * for raw, bidirectional communication between the client and server.
   * 
   * To hint potential proxies about connection hijacking, the Docker client
   * can also optionally send connection upgrade headers.
   * 
   * For example, the client sends this request to upgrade the connection:
   * 
   * ```
   * POST /containers/16253994b7c4/attach?stream=1&stdout=1 HTTP/1.1
   * Upgrade: tcp
   * Connection: Upgrade
   * ```
   * 
   * The Docker daemon will respond with a `101 UPGRADED` response, and will
   * similarly follow with the raw stream:
   * 
   * ```
   * HTTP/1.1 101 UPGRADED
   * Content-Type: application/vnd.docker.raw-stream
   * Connection: Upgrade
   * Upgrade: tcp
   * 
   * [STREAM]
   * ```
   * 
   * ### Stream format
   * 
   * When the TTY setting is disabled in [`POST /containers/create`](#operation/ContainerCreate),
   * the stream over the hijacked connected is multiplexed to separate out
   * `stdout` and `stderr`. The stream consists of a series of frames, each
   * containing a header and a payload.
   * 
   * The header contains the information which the stream writes (`stdout` or
   * `stderr`). It also contains the size of the associated frame encoded in
   * the last four bytes (`uint32`).
   * 
   * It is encoded on the first eight bytes like this:
   * 
   * ```go
   * header := [8]byte{STREAM_TYPE, 0, 0, 0, SIZE1, SIZE2, SIZE3, SIZE4}
   * ```
   * 
   * `STREAM_TYPE` can be:
   * 
   * - 0: `stdin` (is written on `stdout`)
   * - 1: `stdout`
   * - 2: `stderr`
   * 
   * `SIZE1, SIZE2, SIZE3, SIZE4` are the four bytes of the `uint32` size
   * encoded as big endian.
   * 
   * Following the header is the payload, which is the specified number of
   * bytes of `STREAM_TYPE`.
   * 
   * The simplest way to implement this protocol is the following:
   * 
   * 1. Read 8 bytes.
   * 2. Choose `stdout` or `stderr` depending on the first byte.
   * 3. Extract the frame size from the last four bytes.
   * 4. Read the extracted size and output it on the correct output.
   * 5. Goto 1.
   * 
   * ### Stream format when using a TTY
   * 
   * When the TTY setting is enabled in [`POST /containers/create`](#operation/ContainerCreate),
   * the stream is not multiplexed. The data exchanged over the hijacked
   * connection is simply the raw data from the process PTY and client's
   * `stdin`.
   * 
   * Code 500: server error
   */
  type PostContainerAttachResponse500 = ErrorResponse;

  /**
   * `GET /containers/{id}/attach/ws`
   * 
   * Code 400: bad parameter
   */
  type GetContainerAttachWebsocketResponse400 = ErrorResponse;

  /**
   * `GET /containers/{id}/attach/ws`
   * 
   * Code 404: no such container
   */
  type GetContainerAttachWebsocketResponse404 = ErrorResponse;

  /**
   * `GET /containers/{id}/attach/ws`
   * 
   * Code 500: server error
   */
  type GetContainerAttachWebsocketResponse500 = ErrorResponse;

  /**
   * `POST /containers/{id}/wait`
   * 
   * Block until a container stops, then returns the exit code.
   * 
   * Code 200: The container has exit.
   */
  type PostContainerWaitResponse200 = ContainerWaitResponse;

  /**
   * `POST /containers/{id}/wait`
   * 
   * Block until a container stops, then returns the exit code.
   * 
   * Code 400: bad parameter
   */
  type PostContainerWaitResponse400 = ErrorResponse;

  /**
   * `POST /containers/{id}/wait`
   * 
   * Block until a container stops, then returns the exit code.
   * 
   * Code 404: no such container
   */
  type PostContainerWaitResponse404 = ErrorResponse;

  /**
   * `POST /containers/{id}/wait`
   * 
   * Block until a container stops, then returns the exit code.
   * 
   * Code 500: server error
   */
  type PostContainerWaitResponse500 = ErrorResponse;

  /**
   * `DELETE /containers/{id}`
   * 
   * Code 400: bad parameter
   */
  type DeleteContainerDeleteResponse400 = ErrorResponse;

  /**
   * `DELETE /containers/{id}`
   * 
   * Code 404: no such container
   */
  type DeleteContainerDeleteResponse404 = ErrorResponse;

  /**
   * `DELETE /containers/{id}`
   * 
   * Code 409: conflict
   */
  type DeleteContainerDeleteResponse409 = ErrorResponse;

  /**
   * `DELETE /containers/{id}`
   * 
   * Code 500: server error
   */
  type DeleteContainerDeleteResponse500 = ErrorResponse;

  /**
   * `HEAD /containers/{id}/archive`
   * 
   * A response header `X-Docker-Container-Path-Stat` is returned, containing
   * a base64 - encoded JSON object with some filesystem header information
   * about the path.
   * 
   * Code 400: Bad parameter
   */
  type HeadContainerArchiveInfoResponse400 = ErrorResponse;

  /**
   * `HEAD /containers/{id}/archive`
   * 
   * A response header `X-Docker-Container-Path-Stat` is returned, containing
   * a base64 - encoded JSON object with some filesystem header information
   * about the path.
   * 
   * Code 404: Container or path does not exist
   */
  type HeadContainerArchiveInfoResponse404 = ErrorResponse;

  /**
   * `HEAD /containers/{id}/archive`
   * 
   * A response header `X-Docker-Container-Path-Stat` is returned, containing
   * a base64 - encoded JSON object with some filesystem header information
   * about the path.
   * 
   * Code 500: Server error
   */
  type HeadContainerArchiveInfoResponse500 = ErrorResponse;

  /**
   * `GET /containers/{id}/archive`
   * 
   * Get a tar archive of a resource in the filesystem of container id.
   * 
   * Code 400: Bad parameter
   */
  type GetContainerArchiveResponse400 = ErrorResponse;

  /**
   * `GET /containers/{id}/archive`
   * 
   * Get a tar archive of a resource in the filesystem of container id.
   * 
   * Code 404: Container or path does not exist
   */
  type GetContainerArchiveResponse404 = ErrorResponse;

  /**
   * `GET /containers/{id}/archive`
   * 
   * Get a tar archive of a resource in the filesystem of container id.
   * 
   * Code 500: server error
   */
  type GetContainerArchiveResponse500 = ErrorResponse;

  /**
   * `PUT /containers/{id}/archive`
   * 
   * Upload a tar archive to be extracted to a path in the filesystem of container id.
   * `path` parameter is asserted to be a directory. If it exists as a file, 400 error
   * will be returned with message "not a directory".
   * 
   * Code 400: Bad parameter
   */
  type PutPutContainerArchiveResponse400 = ErrorResponse;

  /**
   * `PUT /containers/{id}/archive`
   * 
   * Upload a tar archive to be extracted to a path in the filesystem of container id.
   * `path` parameter is asserted to be a directory. If it exists as a file, 400 error
   * will be returned with message "not a directory".
   * 
   * Code 403: Permission denied, the volume or container rootfs is marked as read-only.
   */
  type PutPutContainerArchiveResponse403 = ErrorResponse;

  /**
   * `PUT /containers/{id}/archive`
   * 
   * Upload a tar archive to be extracted to a path in the filesystem of container id.
   * `path` parameter is asserted to be a directory. If it exists as a file, 400 error
   * will be returned with message "not a directory".
   * 
   * Code 404: No such container or path does not exist inside the container
   */
  type PutPutContainerArchiveResponse404 = ErrorResponse;

  /**
   * `PUT /containers/{id}/archive`
   * 
   * Upload a tar archive to be extracted to a path in the filesystem of container id.
   * `path` parameter is asserted to be a directory. If it exists as a file, 400 error
   * will be returned with message "not a directory".
   * 
   * Code 500: Server error
   */
  type PutPutContainerArchiveResponse500 = ErrorResponse;

  /**
   * `POST /containers/prune`
   * 
   * Code 200: No error
   */
  type PostContainerPruneResponse200 = {
    /**
     * Container IDs that were deleted
     */
    'ContainersDeleted'?: Array<string>;
    /**
     * Disk space reclaimed in bytes
     */
    'SpaceReclaimed'?: number;
  };

  /**
   * `POST /containers/prune`
   * 
   * Code 500: Server error
   */
  type PostContainerPruneResponse500 = ErrorResponse;

  /**
   * `GET /images/json`
   * 
   * Returns a list of images on the server. Note that it uses a different, smaller representation of an image than inspecting a single image.
   * 
   * Code 200: Summary image data for the images matching the query
   */
  type GetImageListResponse200 = Array<ImageSummary>;

  /**
   * `GET /images/json`
   * 
   * Returns a list of images on the server. Note that it uses a different, smaller representation of an image than inspecting a single image.
   * 
   * Code 500: server error
   */
  type GetImageListResponse500 = ErrorResponse;

  /**
   * `POST /build`
   * 
   * Build an image from a tar archive with a `Dockerfile` in it.
   * 
   * The `Dockerfile` specifies how the image is built from the tar archive. It is typically in the archive's root, but can be at a different path or have a different name by specifying the `dockerfile` parameter. [See the `Dockerfile` reference for more information](/engine/reference/builder/).
   * 
   * The Docker daemon performs a preliminary validation of the `Dockerfile` before starting the build, and returns an error if the syntax is incorrect. After that, each instruction is run one-by-one until the ID of the new image is output.
   * 
   * The build is canceled if the client drops the connection by quitting or being killed.
   * 
   * Code 400: Bad parameter
   */
  type PostImageBuildResponse400 = ErrorResponse;

  /**
   * `POST /build`
   * 
   * Build an image from a tar archive with a `Dockerfile` in it.
   * 
   * The `Dockerfile` specifies how the image is built from the tar archive. It is typically in the archive's root, but can be at a different path or have a different name by specifying the `dockerfile` parameter. [See the `Dockerfile` reference for more information](/engine/reference/builder/).
   * 
   * The Docker daemon performs a preliminary validation of the `Dockerfile` before starting the build, and returns an error if the syntax is incorrect. After that, each instruction is run one-by-one until the ID of the new image is output.
   * 
   * The build is canceled if the client drops the connection by quitting or being killed.
   * 
   * Code 500: server error
   */
  type PostImageBuildResponse500 = ErrorResponse;

  /**
   * `POST /build/prune`
   * 
   * Code 200: No error
   */
  type PostBuildPruneResponse200 = {
    'CachesDeleted'?: Array<string>;
    /**
     * Disk space reclaimed in bytes
     */
    'SpaceReclaimed'?: number;
  };

  /**
   * `POST /build/prune`
   * 
   * Code 500: Server error
   */
  type PostBuildPruneResponse500 = ErrorResponse;

  /**
   * `POST /images/create`
   * 
   * Create an image by either pulling it from a registry or importing it.
   * 
   * Code 404: repository does not exist or no read access
   */
  type PostImageCreateResponse404 = ErrorResponse;

  /**
   * `POST /images/create`
   * 
   * Create an image by either pulling it from a registry or importing it.
   * 
   * Code 500: server error
   */
  type PostImageCreateResponse500 = ErrorResponse;

  /**
   * `GET /images/{name}/json`
   * 
   * Return low-level information about an image.
   * 
   * Code 200: No error
   */
  type GetImageInspectResponse200 = ImageInspect;

  /**
   * `GET /images/{name}/json`
   * 
   * Return low-level information about an image.
   * 
   * Code 404: No such image
   */
  type GetImageInspectResponse404 = ErrorResponse;

  /**
   * `GET /images/{name}/json`
   * 
   * Return low-level information about an image.
   * 
   * Code 500: Server error
   */
  type GetImageInspectResponse500 = ErrorResponse;

  /**
   * `GET /images/{name}/history`
   * 
   * Return parent layers of an image.
   * 
   * Code 200: List of image layers
   */
  type GetImageHistoryResponse200 = Array<{
    'Id'?: string;
    'Created'?: number;
    'CreatedBy'?: string;
    'Tags'?: Array<string>;
    'Size'?: number;
    'Comment'?: string;
  }>;

  /**
   * `GET /images/{name}/history`
   * 
   * Return parent layers of an image.
   * 
   * Code 404: No such image
   */
  type GetImageHistoryResponse404 = ErrorResponse;

  /**
   * `GET /images/{name}/history`
   * 
   * Return parent layers of an image.
   * 
   * Code 500: Server error
   */
  type GetImageHistoryResponse500 = ErrorResponse;

  /**
   * `POST /images/{name}/push`
   * 
   * Push an image to a registry.
   * 
   * If you wish to push an image on to a private registry, that image must
   * already have a tag which references the registry. For example,
   * `registry.example.com/myimage:latest`.
   * 
   * The push is cancelled if the HTTP connection is closed.
   * 
   * Code 404: No such image
   */
  type PostImagePushResponse404 = ErrorResponse;

  /**
   * `POST /images/{name}/push`
   * 
   * Push an image to a registry.
   * 
   * If you wish to push an image on to a private registry, that image must
   * already have a tag which references the registry. For example,
   * `registry.example.com/myimage:latest`.
   * 
   * The push is cancelled if the HTTP connection is closed.
   * 
   * Code 500: Server error
   */
  type PostImagePushResponse500 = ErrorResponse;

  /**
   * `POST /images/{name}/tag`
   * 
   * Tag an image so that it becomes part of a repository.
   * 
   * Code 400: Bad parameter
   */
  type PostImageTagResponse400 = ErrorResponse;

  /**
   * `POST /images/{name}/tag`
   * 
   * Tag an image so that it becomes part of a repository.
   * 
   * Code 404: No such image
   */
  type PostImageTagResponse404 = ErrorResponse;

  /**
   * `POST /images/{name}/tag`
   * 
   * Tag an image so that it becomes part of a repository.
   * 
   * Code 409: Conflict
   */
  type PostImageTagResponse409 = ErrorResponse;

  /**
   * `POST /images/{name}/tag`
   * 
   * Tag an image so that it becomes part of a repository.
   * 
   * Code 500: Server error
   */
  type PostImageTagResponse500 = ErrorResponse;

  /**
   * `DELETE /images/{name}`
   * 
   * Remove an image, along with any untagged parent images that were
   * referenced by that image.
   * 
   * Images can't be removed if they have descendant images, are being
   * used by a running container or are being used by a build.
   * 
   * Code 200: The image was deleted successfully
   */
  type DeleteImageDeleteResponse200 = Array<ImageDeleteResponseItem>;

  /**
   * `DELETE /images/{name}`
   * 
   * Remove an image, along with any untagged parent images that were
   * referenced by that image.
   * 
   * Images can't be removed if they have descendant images, are being
   * used by a running container or are being used by a build.
   * 
   * Code 404: No such image
   */
  type DeleteImageDeleteResponse404 = ErrorResponse;

  /**
   * `DELETE /images/{name}`
   * 
   * Remove an image, along with any untagged parent images that were
   * referenced by that image.
   * 
   * Images can't be removed if they have descendant images, are being
   * used by a running container or are being used by a build.
   * 
   * Code 409: Conflict
   */
  type DeleteImageDeleteResponse409 = ErrorResponse;

  /**
   * `DELETE /images/{name}`
   * 
   * Remove an image, along with any untagged parent images that were
   * referenced by that image.
   * 
   * Images can't be removed if they have descendant images, are being
   * used by a running container or are being used by a build.
   * 
   * Code 500: Server error
   */
  type DeleteImageDeleteResponse500 = ErrorResponse;

  /**
   * `GET /images/search`
   * 
   * Search for an image on Docker Hub.
   * 
   * Code 200: No error
   */
  type GetImageSearchResponse200 = Array<{
    'description'?: string;
    'is_official'?: boolean;
    'is_automated'?: boolean;
    'name'?: string;
    'star_count'?: number;
  }>;

  /**
   * `GET /images/search`
   * 
   * Search for an image on Docker Hub.
   * 
   * Code 500: Server error
   */
  type GetImageSearchResponse500 = ErrorResponse;

  /**
   * `POST /images/prune`
   * 
   * Code 200: No error
   */
  type PostImagePruneResponse200 = {
    /**
     * Images that were deleted
     */
    'ImagesDeleted'?: Array<ImageDeleteResponseItem>;
    /**
     * Disk space reclaimed in bytes
     */
    'SpaceReclaimed'?: number;
  };

  /**
   * `POST /images/prune`
   * 
   * Code 500: Server error
   */
  type PostImagePruneResponse500 = ErrorResponse;

  /**
   * `POST /auth`
   * 
   * Validate credentials for a registry and, if available, get an identity
   * token for accessing the registry without password.
   * 
   * Code 200: An identity token was generated successfully.
   */
  type PostSystemAuthResponse200 = {
    /**
     * The status of the authentication
     */
    'Status'?: string;
    /**
     * An opaque token used to authenticate a user after a successful login
     */
    'IdentityToken'?: string;
  };

  /**
   * `POST /auth`
   * 
   * Validate credentials for a registry and, if available, get an identity
   * token for accessing the registry without password.
   * 
   * Code 500: Server error
   */
  type PostSystemAuthResponse500 = ErrorResponse;

  /**
   * `GET /info`
   * 
   * Code 200: No error
   */
  type GetSystemInfoResponse200 = SystemInfo;

  /**
   * `GET /info`
   * 
   * Code 500: Server error
   */
  type GetSystemInfoResponse500 = ErrorResponse;

  /**
   * `GET /version`
   * 
   * Returns the version of Docker that is running and various information about the system that Docker is running on.
   * 
   * Code 200: no error
   */
  type GetSystemVersionResponse200 = SystemVersion;

  /**
   * `GET /version`
   * 
   * Returns the version of Docker that is running and various information about the system that Docker is running on.
   * 
   * Code 500: server error
   */
  type GetSystemVersionResponse500 = ErrorResponse;

  /**
   * `GET /_ping`
   * 
   * This is a dummy endpoint you can use to test if the server is accessible.
   * 
   * Code 200: no error
   */
  type GetSystemPingResponse200 = string;

  /**
   * `GET /_ping`
   * 
   * This is a dummy endpoint you can use to test if the server is accessible.
   * 
   * Code 500: server error
   */
  type GetSystemPingResponse500 = ErrorResponse;

  /**
   * `HEAD /_ping`
   * 
   * This is a dummy endpoint you can use to test if the server is accessible.
   * 
   * Code 200: no error
   */
  type HeadSystemPingHeadResponse200 = string;

  /**
   * `HEAD /_ping`
   * 
   * This is a dummy endpoint you can use to test if the server is accessible.
   * 
   * Code 500: server error
   */
  type HeadSystemPingHeadResponse500 = ErrorResponse;

  /**
   * `POST /commit`
   * 
   * Code 201: no error
   */
  type PostImageCommitResponse201 = IdResponse;

  /**
   * `POST /commit`
   * 
   * Code 404: no such container
   */
  type PostImageCommitResponse404 = ErrorResponse;

  /**
   * `POST /commit`
   * 
   * Code 500: server error
   */
  type PostImageCommitResponse500 = ErrorResponse;

  /**
   * `GET /events`
   * 
   * Stream real-time events from the server.
   * 
   * Various objects within Docker report events when something happens to them.
   * 
   * Containers report these events: `attach`, `commit`, `copy`, `create`, `destroy`, `detach`, `die`, `exec_create`, `exec_detach`, `exec_start`, `exec_die`, `export`, `health_status`, `kill`, `oom`, `pause`, `rename`, `resize`, `restart`, `start`, `stop`, `top`, `unpause`, `update`, and `prune`
   * 
   * Images report these events: `delete`, `import`, `load`, `pull`, `push`, `save`, `tag`, `untag`, and `prune`
   * 
   * Volumes report these events: `create`, `mount`, `unmount`, `destroy`, and `prune`
   * 
   * Networks report these events: `create`, `connect`, `disconnect`, `destroy`, `update`, `remove`, and `prune`
   * 
   * The Docker daemon reports these events: `reload`
   * 
   * Services report these events: `create`, `update`, and `remove`
   * 
   * Nodes report these events: `create`, `update`, and `remove`
   * 
   * Secrets report these events: `create`, `update`, and `remove`
   * 
   * Configs report these events: `create`, `update`, and `remove`
   * 
   * The Builder reports `prune` events
   * 
   * Code 200: no error
   */
  type GetSystemEventsResponse200 = EventMessage;

  /**
   * `GET /events`
   * 
   * Stream real-time events from the server.
   * 
   * Various objects within Docker report events when something happens to them.
   * 
   * Containers report these events: `attach`, `commit`, `copy`, `create`, `destroy`, `detach`, `die`, `exec_create`, `exec_detach`, `exec_start`, `exec_die`, `export`, `health_status`, `kill`, `oom`, `pause`, `rename`, `resize`, `restart`, `start`, `stop`, `top`, `unpause`, `update`, and `prune`
   * 
   * Images report these events: `delete`, `import`, `load`, `pull`, `push`, `save`, `tag`, `untag`, and `prune`
   * 
   * Volumes report these events: `create`, `mount`, `unmount`, `destroy`, and `prune`
   * 
   * Networks report these events: `create`, `connect`, `disconnect`, `destroy`, `update`, `remove`, and `prune`
   * 
   * The Docker daemon reports these events: `reload`
   * 
   * Services report these events: `create`, `update`, and `remove`
   * 
   * Nodes report these events: `create`, `update`, and `remove`
   * 
   * Secrets report these events: `create`, `update`, and `remove`
   * 
   * Configs report these events: `create`, `update`, and `remove`
   * 
   * The Builder reports `prune` events
   * 
   * Code 400: bad parameter
   */
  type GetSystemEventsResponse400 = ErrorResponse;

  /**
   * `GET /events`
   * 
   * Stream real-time events from the server.
   * 
   * Various objects within Docker report events when something happens to them.
   * 
   * Containers report these events: `attach`, `commit`, `copy`, `create`, `destroy`, `detach`, `die`, `exec_create`, `exec_detach`, `exec_start`, `exec_die`, `export`, `health_status`, `kill`, `oom`, `pause`, `rename`, `resize`, `restart`, `start`, `stop`, `top`, `unpause`, `update`, and `prune`
   * 
   * Images report these events: `delete`, `import`, `load`, `pull`, `push`, `save`, `tag`, `untag`, and `prune`
   * 
   * Volumes report these events: `create`, `mount`, `unmount`, `destroy`, and `prune`
   * 
   * Networks report these events: `create`, `connect`, `disconnect`, `destroy`, `update`, `remove`, and `prune`
   * 
   * The Docker daemon reports these events: `reload`
   * 
   * Services report these events: `create`, `update`, and `remove`
   * 
   * Nodes report these events: `create`, `update`, and `remove`
   * 
   * Secrets report these events: `create`, `update`, and `remove`
   * 
   * Configs report these events: `create`, `update`, and `remove`
   * 
   * The Builder reports `prune` events
   * 
   * Code 500: server error
   */
  type GetSystemEventsResponse500 = ErrorResponse;

  /**
   * `GET /system/df`
   * 
   * Code 200: no error
   */
  type GetSystemDataUsageResponse200 = {
    'LayersSize'?: number;
    'Images'?: Array<ImageSummary>;
    'Containers'?: Array<ContainerSummary>;
    'Volumes'?: Array<Volume>;
    'BuildCache'?: Array<BuildCache>;
  };

  /**
   * `GET /system/df`
   * 
   * Code 500: server error
   */
  type GetSystemDataUsageResponse500 = ErrorResponse;

  /**
   * `GET /images/{name}/get`
   * 
   * Get a tarball containing all images and metadata for a repository.
   * 
   * If `name` is a specific name and tag (e.g. `ubuntu:latest`), then only that image (and its parents) are returned. If `name` is an image ID, similarly only that image (and its parents) are returned, but with the exclusion of the `repositories` file in the tarball, as there were no image names referenced.
   * 
   * ### Image tarball format
   * 
   * An image tarball contains one directory per image layer (named using its long ID), each containing these files:
   * 
   * - `VERSION`: currently `1.0` - the file format version
   * - `json`: detailed layer information, similar to `docker inspect layer_id`
   * - `layer.tar`: A tarfile containing the filesystem changes in this layer
   * 
   * The `layer.tar` file contains `aufs` style `.wh..wh.aufs` files and directories for storing attribute changes and deletions.
   * 
   * If the tarball defines a repository, the tarball should also include a `repositories` file at the root that contains a list of repository and tag names mapped to layer IDs.
   * 
   * ```json
   * {
   *   "hello-world": {
   *     "latest": "565a9d68a73f6706862bfe8409a7f659776d4d60a8d096eb4a3cbce6999cc2a1"
   *   }
   * }
   * ```
   * 
   * Code 200: no error
   */
  type GetImageGetResponse200 = string;

  /**
   * `GET /images/{name}/get`
   * 
   * Get a tarball containing all images and metadata for a repository.
   * 
   * If `name` is a specific name and tag (e.g. `ubuntu:latest`), then only that image (and its parents) are returned. If `name` is an image ID, similarly only that image (and its parents) are returned, but with the exclusion of the `repositories` file in the tarball, as there were no image names referenced.
   * 
   * ### Image tarball format
   * 
   * An image tarball contains one directory per image layer (named using its long ID), each containing these files:
   * 
   * - `VERSION`: currently `1.0` - the file format version
   * - `json`: detailed layer information, similar to `docker inspect layer_id`
   * - `layer.tar`: A tarfile containing the filesystem changes in this layer
   * 
   * The `layer.tar` file contains `aufs` style `.wh..wh.aufs` files and directories for storing attribute changes and deletions.
   * 
   * If the tarball defines a repository, the tarball should also include a `repositories` file at the root that contains a list of repository and tag names mapped to layer IDs.
   * 
   * ```json
   * {
   *   "hello-world": {
   *     "latest": "565a9d68a73f6706862bfe8409a7f659776d4d60a8d096eb4a3cbce6999cc2a1"
   *   }
   * }
   * ```
   * 
   * Code 500: server error
   */
  type GetImageGetResponse500 = ErrorResponse;

  /**
   * `GET /images/get`
   * 
   * Get a tarball containing all images and metadata for several image
   * repositories.
   * 
   * For each value of the `names` parameter: if it is a specific name and
   * tag (e.g. `ubuntu:latest`), then only that image (and its parents) are
   * returned; if it is an image ID, similarly only that image (and its parents)
   * are returned and there would be no names referenced in the 'repositories'
   * file for this image ID.
   * 
   * For details on the format, see the [export image endpoint](#operation/ImageGet).
   * 
   * Code 200: no error
   */
  type GetImageGetAllResponse200 = string;

  /**
   * `GET /images/get`
   * 
   * Get a tarball containing all images and metadata for several image
   * repositories.
   * 
   * For each value of the `names` parameter: if it is a specific name and
   * tag (e.g. `ubuntu:latest`), then only that image (and its parents) are
   * returned; if it is an image ID, similarly only that image (and its parents)
   * are returned and there would be no names referenced in the 'repositories'
   * file for this image ID.
   * 
   * For details on the format, see the [export image endpoint](#operation/ImageGet).
   * 
   * Code 500: server error
   */
  type GetImageGetAllResponse500 = ErrorResponse;

  /**
   * `POST /images/load`
   * 
   * Load a set of images and tags into a repository.
   * 
   * For details on the format, see the [export image endpoint](#operation/ImageGet).
   * 
   * Code 500: server error
   */
  type PostImageLoadResponse500 = ErrorResponse;

  /**
   * `POST /containers/{id}/exec`
   * 
   * Run a command inside a running container.
   * 
   * Code 201: no error
   */
  type PostContainerExecResponse201 = IdResponse;

  /**
   * `POST /containers/{id}/exec`
   * 
   * Run a command inside a running container.
   * 
   * Code 404: no such container
   */
  type PostContainerExecResponse404 = ErrorResponse;

  /**
   * `POST /containers/{id}/exec`
   * 
   * Run a command inside a running container.
   * 
   * Code 409: container is paused
   */
  type PostContainerExecResponse409 = ErrorResponse;

  /**
   * `POST /containers/{id}/exec`
   * 
   * Run a command inside a running container.
   * 
   * Code 500: Server error
   */
  type PostContainerExecResponse500 = ErrorResponse;

  /**
   * `POST /exec/{id}/start`
   * 
   * Starts a previously set up exec instance. If detach is true, this endpoint
   * returns immediately after starting the command. Otherwise, it sets up an
   * interactive session with the command.
   * 
   * Code 404: No such exec instance
   */
  type PostExecStartResponse404 = ErrorResponse;

  /**
   * `POST /exec/{id}/start`
   * 
   * Starts a previously set up exec instance. If detach is true, this endpoint
   * returns immediately after starting the command. Otherwise, it sets up an
   * interactive session with the command.
   * 
   * Code 409: Container is stopped or paused
   */
  type PostExecStartResponse409 = ErrorResponse;

  /**
   * `POST /exec/{id}/resize`
   * 
   * Resize the TTY session used by an exec instance. This endpoint only works
   * if `tty` was specified as part of creating and starting the exec instance.
   * 
   * Code 400: bad parameter
   */
  type PostExecResizeResponse400 = ErrorResponse;

  /**
   * `POST /exec/{id}/resize`
   * 
   * Resize the TTY session used by an exec instance. This endpoint only works
   * if `tty` was specified as part of creating and starting the exec instance.
   * 
   * Code 404: No such exec instance
   */
  type PostExecResizeResponse404 = ErrorResponse;

  /**
   * `POST /exec/{id}/resize`
   * 
   * Resize the TTY session used by an exec instance. This endpoint only works
   * if `tty` was specified as part of creating and starting the exec instance.
   * 
   * Code 500: Server error
   */
  type PostExecResizeResponse500 = ErrorResponse;

  /**
   * `GET /exec/{id}/json`
   * 
   * Return low-level information about an exec instance.
   * 
   * Code 200: No error
   */
  type GetExecInspectResponse200 = {
    'CanRemove'?: boolean;
    'DetachKeys'?: string;
    'ID'?: string;
    'Running'?: boolean;
    'ExitCode'?: number;
    'ProcessConfig'?: ProcessConfig;
    'OpenStdin'?: boolean;
    'OpenStderr'?: boolean;
    'OpenStdout'?: boolean;
    'ContainerID'?: string;
    /**
     * The system process ID for the exec process.
     */
    'Pid'?: number;
  };

  /**
   * `GET /exec/{id}/json`
   * 
   * Return low-level information about an exec instance.
   * 
   * Code 404: No such exec instance
   */
  type GetExecInspectResponse404 = ErrorResponse;

  /**
   * `GET /exec/{id}/json`
   * 
   * Return low-level information about an exec instance.
   * 
   * Code 500: Server error
   */
  type GetExecInspectResponse500 = ErrorResponse;

  /**
   * `GET /volumes`
   * 
   * Code 200: Summary volume data that matches the query
   */
  type GetVolumeListResponse200 = {
    /**
     * List of volumes
     */
    'Volumes'?: Array<Volume>;
    /**
     * Warnings that occurred when fetching the list of volumes.
     */
    'Warnings'?: Array<string>;
  };

  /**
   * `GET /volumes`
   * 
   * Code 500: Server error
   */
  type GetVolumeListResponse500 = ErrorResponse;

  /**
   * `POST /volumes/create`
   * 
   * Code 201: The volume was created successfully
   */
  type PostVolumeCreateResponse201 = Volume;

  /**
   * `POST /volumes/create`
   * 
   * Code 500: Server error
   */
  type PostVolumeCreateResponse500 = ErrorResponse;

  /**
   * `GET /volumes/{name}`
   * 
   * Code 200: No error
   */
  type GetVolumeInspectResponse200 = Volume;

  /**
   * `GET /volumes/{name}`
   * 
   * Code 404: No such volume
   */
  type GetVolumeInspectResponse404 = ErrorResponse;

  /**
   * `GET /volumes/{name}`
   * 
   * Code 500: Server error
   */
  type GetVolumeInspectResponse500 = ErrorResponse;

  /**
   * `DELETE /volumes/{name}`
   * 
   * Instruct the driver to remove the volume.
   * 
   * Code 404: No such volume or volume driver
   */
  type DeleteVolumeDeleteResponse404 = ErrorResponse;

  /**
   * `DELETE /volumes/{name}`
   * 
   * Instruct the driver to remove the volume.
   * 
   * Code 409: Volume is in use and cannot be removed
   */
  type DeleteVolumeDeleteResponse409 = ErrorResponse;

  /**
   * `DELETE /volumes/{name}`
   * 
   * Instruct the driver to remove the volume.
   * 
   * Code 500: Server error
   */
  type DeleteVolumeDeleteResponse500 = ErrorResponse;

  /**
   * `POST /volumes/prune`
   * 
   * Code 200: No error
   */
  type PostVolumePruneResponse200 = {
    /**
     * Volumes that were deleted
     */
    'VolumesDeleted'?: Array<string>;
    /**
     * Disk space reclaimed in bytes
     */
    'SpaceReclaimed'?: number;
  };

  /**
   * `POST /volumes/prune`
   * 
   * Code 500: Server error
   */
  type PostVolumePruneResponse500 = ErrorResponse;

  /**
   * `GET /networks`
   * 
   * Returns a list of networks. For details on the format, see the
   * [network inspect endpoint](#operation/NetworkInspect).
   * 
   * Note that it uses a different, smaller representation of a network than
   * inspecting a single network. For example, the list of containers attached
   * to the network is not propagated in API versions 1.28 and up.
   * 
   * Code 200: No error
   */
  type GetNetworkListResponse200 = Array<Network>;

  /**
   * `GET /networks`
   * 
   * Returns a list of networks. For details on the format, see the
   * [network inspect endpoint](#operation/NetworkInspect).
   * 
   * Note that it uses a different, smaller representation of a network than
   * inspecting a single network. For example, the list of containers attached
   * to the network is not propagated in API versions 1.28 and up.
   * 
   * Code 500: Server error
   */
  type GetNetworkListResponse500 = ErrorResponse;

  /**
   * `GET /networks/{id}`
   * 
   * Code 200: No error
   */
  type GetNetworkInspectResponse200 = Network;

  /**
   * `GET /networks/{id}`
   * 
   * Code 404: Network not found
   */
  type GetNetworkInspectResponse404 = ErrorResponse;

  /**
   * `GET /networks/{id}`
   * 
   * Code 500: Server error
   */
  type GetNetworkInspectResponse500 = ErrorResponse;

  /**
   * `DELETE /networks/{id}`
   * 
   * Code 403: operation not supported for pre-defined networks
   */
  type DeleteNetworkDeleteResponse403 = ErrorResponse;

  /**
   * `DELETE /networks/{id}`
   * 
   * Code 404: no such network
   */
  type DeleteNetworkDeleteResponse404 = ErrorResponse;

  /**
   * `DELETE /networks/{id}`
   * 
   * Code 500: Server error
   */
  type DeleteNetworkDeleteResponse500 = ErrorResponse;

  /**
   * `POST /networks/create`
   * 
   * Code 201: No error
   */
  type PostNetworkCreateResponse201 = {
    /**
     * The ID of the created network.
     */
    'Id'?: string;
    'Warning'?: string;
  };

  /**
   * `POST /networks/create`
   * 
   * Code 403: operation not supported for pre-defined networks
   */
  type PostNetworkCreateResponse403 = ErrorResponse;

  /**
   * `POST /networks/create`
   * 
   * Code 404: plugin not found
   */
  type PostNetworkCreateResponse404 = ErrorResponse;

  /**
   * `POST /networks/create`
   * 
   * Code 500: Server error
   */
  type PostNetworkCreateResponse500 = ErrorResponse;

  /**
   * `POST /networks/{id}/connect`
   * 
   * Code 403: Operation not supported for swarm scoped networks
   */
  type PostNetworkConnectResponse403 = ErrorResponse;

  /**
   * `POST /networks/{id}/connect`
   * 
   * Code 404: Network or container not found
   */
  type PostNetworkConnectResponse404 = ErrorResponse;

  /**
   * `POST /networks/{id}/connect`
   * 
   * Code 500: Server error
   */
  type PostNetworkConnectResponse500 = ErrorResponse;

  /**
   * `POST /networks/{id}/disconnect`
   * 
   * Code 403: Operation not supported for swarm scoped networks
   */
  type PostNetworkDisconnectResponse403 = ErrorResponse;

  /**
   * `POST /networks/{id}/disconnect`
   * 
   * Code 404: Network or container not found
   */
  type PostNetworkDisconnectResponse404 = ErrorResponse;

  /**
   * `POST /networks/{id}/disconnect`
   * 
   * Code 500: Server error
   */
  type PostNetworkDisconnectResponse500 = ErrorResponse;

  /**
   * `POST /networks/prune`
   * 
   * Code 200: No error
   */
  type PostNetworkPruneResponse200 = {
    /**
     * Networks that were deleted
     */
    'NetworksDeleted'?: Array<string>;
  };

  /**
   * `POST /networks/prune`
   * 
   * Code 500: Server error
   */
  type PostNetworkPruneResponse500 = ErrorResponse;

  /**
   * `GET /plugins`
   * 
   * Returns information about installed plugins.
   * 
   * Code 200: No error
   */
  type GetPluginListResponse200 = Array<Plugin>;

  /**
   * `GET /plugins`
   * 
   * Returns information about installed plugins.
   * 
   * Code 500: Server error
   */
  type GetPluginListResponse500 = ErrorResponse;

  /**
   * `GET /plugins/privileges`
   * 
   * Code 200: no error
   */
  type GetGetPluginPrivilegesResponse200 = Array<PluginPrivilege>;

  /**
   * `GET /plugins/privileges`
   * 
   * Code 500: server error
   */
  type GetGetPluginPrivilegesResponse500 = ErrorResponse;

  /**
   * `POST /plugins/pull`
   * 
   * Pulls and installs a plugin. After the plugin is installed, it can be
   * enabled using the [`POST /plugins/{name}/enable` endpoint](#operation/PostPluginsEnable).
   * 
   * Code 500: server error
   */
  type PostPluginPullResponse500 = ErrorResponse;

  /**
   * `GET /plugins/{name}/json`
   * 
   * Code 200: no error
   */
  type GetPluginInspectResponse200 = Plugin;

  /**
   * `GET /plugins/{name}/json`
   * 
   * Code 404: plugin is not installed
   */
  type GetPluginInspectResponse404 = ErrorResponse;

  /**
   * `GET /plugins/{name}/json`
   * 
   * Code 500: server error
   */
  type GetPluginInspectResponse500 = ErrorResponse;

  /**
   * `DELETE /plugins/{name}`
   * 
   * Code 200: no error
   */
  type DeletePluginDeleteResponse200 = Plugin;

  /**
   * `DELETE /plugins/{name}`
   * 
   * Code 404: plugin is not installed
   */
  type DeletePluginDeleteResponse404 = ErrorResponse;

  /**
   * `DELETE /plugins/{name}`
   * 
   * Code 500: server error
   */
  type DeletePluginDeleteResponse500 = ErrorResponse;

  /**
   * `POST /plugins/{name}/enable`
   * 
   * Code 404: plugin is not installed
   */
  type PostPluginEnableResponse404 = ErrorResponse;

  /**
   * `POST /plugins/{name}/enable`
   * 
   * Code 500: server error
   */
  type PostPluginEnableResponse500 = ErrorResponse;

  /**
   * `POST /plugins/{name}/disable`
   * 
   * Code 404: plugin is not installed
   */
  type PostPluginDisableResponse404 = ErrorResponse;

  /**
   * `POST /plugins/{name}/disable`
   * 
   * Code 500: server error
   */
  type PostPluginDisableResponse500 = ErrorResponse;

  /**
   * `POST /plugins/{name}/upgrade`
   * 
   * Code 404: plugin not installed
   */
  type PostPluginUpgradeResponse404 = ErrorResponse;

  /**
   * `POST /plugins/{name}/upgrade`
   * 
   * Code 500: server error
   */
  type PostPluginUpgradeResponse500 = ErrorResponse;

  /**
   * `POST /plugins/create`
   * 
   * Code 500: server error
   */
  type PostPluginCreateResponse500 = ErrorResponse;

  /**
   * `POST /plugins/{name}/push`
   * 
   * Push a plugin to the registry.
   * 
   * Code 404: plugin not installed
   */
  type PostPluginPushResponse404 = ErrorResponse;

  /**
   * `POST /plugins/{name}/push`
   * 
   * Push a plugin to the registry.
   * 
   * Code 500: server error
   */
  type PostPluginPushResponse500 = ErrorResponse;

  /**
   * `POST /plugins/{name}/set`
   * 
   * Code 404: Plugin not installed
   */
  type PostPluginSetResponse404 = ErrorResponse;

  /**
   * `POST /plugins/{name}/set`
   * 
   * Code 500: Server error
   */
  type PostPluginSetResponse500 = ErrorResponse;

  /**
   * `GET /nodes`
   * 
   * Code 200: no error
   */
  type GetNodeListResponse200 = Array<Node>;

  /**
   * `GET /nodes`
   * 
   * Code 500: server error
   */
  type GetNodeListResponse500 = ErrorResponse;

  /**
   * `GET /nodes`
   * 
   * Code 503: node is not part of a swarm
   */
  type GetNodeListResponse503 = ErrorResponse;

  /**
   * `GET /nodes/{id}`
   * 
   * Code 200: no error
   */
  type GetNodeInspectResponse200 = Node;

  /**
   * `GET /nodes/{id}`
   * 
   * Code 404: no such node
   */
  type GetNodeInspectResponse404 = ErrorResponse;

  /**
   * `GET /nodes/{id}`
   * 
   * Code 500: server error
   */
  type GetNodeInspectResponse500 = ErrorResponse;

  /**
   * `GET /nodes/{id}`
   * 
   * Code 503: node is not part of a swarm
   */
  type GetNodeInspectResponse503 = ErrorResponse;

  /**
   * `DELETE /nodes/{id}`
   * 
   * Code 404: no such node
   */
  type DeleteNodeDeleteResponse404 = ErrorResponse;

  /**
   * `DELETE /nodes/{id}`
   * 
   * Code 500: server error
   */
  type DeleteNodeDeleteResponse500 = ErrorResponse;

  /**
   * `DELETE /nodes/{id}`
   * 
   * Code 503: node is not part of a swarm
   */
  type DeleteNodeDeleteResponse503 = ErrorResponse;

  /**
   * `POST /nodes/{id}/update`
   * 
   * Code 400: bad parameter
   */
  type PostNodeUpdateResponse400 = ErrorResponse;

  /**
   * `POST /nodes/{id}/update`
   * 
   * Code 404: no such node
   */
  type PostNodeUpdateResponse404 = ErrorResponse;

  /**
   * `POST /nodes/{id}/update`
   * 
   * Code 500: server error
   */
  type PostNodeUpdateResponse500 = ErrorResponse;

  /**
   * `POST /nodes/{id}/update`
   * 
   * Code 503: node is not part of a swarm
   */
  type PostNodeUpdateResponse503 = ErrorResponse;

  /**
   * `GET /swarm`
   * 
   * Code 200: no error
   */
  type GetSwarmInspectResponse200 = Swarm;

  /**
   * `GET /swarm`
   * 
   * Code 404: no such swarm
   */
  type GetSwarmInspectResponse404 = ErrorResponse;

  /**
   * `GET /swarm`
   * 
   * Code 500: server error
   */
  type GetSwarmInspectResponse500 = ErrorResponse;

  /**
   * `GET /swarm`
   * 
   * Code 503: node is not part of a swarm
   */
  type GetSwarmInspectResponse503 = ErrorResponse;

  /**
   * `POST /swarm/init`
   * 
   * Code 200: no error
   */
  type PostSwarmInitResponse200 = string;

  /**
   * `POST /swarm/init`
   * 
   * Code 400: bad parameter
   */
  type PostSwarmInitResponse400 = ErrorResponse;

  /**
   * `POST /swarm/init`
   * 
   * Code 500: server error
   */
  type PostSwarmInitResponse500 = ErrorResponse;

  /**
   * `POST /swarm/init`
   * 
   * Code 503: node is already part of a swarm
   */
  type PostSwarmInitResponse503 = ErrorResponse;

  /**
   * `POST /swarm/join`
   * 
   * Code 400: bad parameter
   */
  type PostSwarmJoinResponse400 = ErrorResponse;

  /**
   * `POST /swarm/join`
   * 
   * Code 500: server error
   */
  type PostSwarmJoinResponse500 = ErrorResponse;

  /**
   * `POST /swarm/join`
   * 
   * Code 503: node is already part of a swarm
   */
  type PostSwarmJoinResponse503 = ErrorResponse;

  /**
   * `POST /swarm/leave`
   * 
   * Code 500: server error
   */
  type PostSwarmLeaveResponse500 = ErrorResponse;

  /**
   * `POST /swarm/leave`
   * 
   * Code 503: node is not part of a swarm
   */
  type PostSwarmLeaveResponse503 = ErrorResponse;

  /**
   * `POST /swarm/update`
   * 
   * Code 400: bad parameter
   */
  type PostSwarmUpdateResponse400 = ErrorResponse;

  /**
   * `POST /swarm/update`
   * 
   * Code 500: server error
   */
  type PostSwarmUpdateResponse500 = ErrorResponse;

  /**
   * `POST /swarm/update`
   * 
   * Code 503: node is not part of a swarm
   */
  type PostSwarmUpdateResponse503 = ErrorResponse;

  /**
   * `GET /swarm/unlockkey`
   * 
   * Code 200: no error
   */
  type GetSwarmUnlockkeyResponse200 = {
    /**
     * The swarm's unlock key.
     */
    'UnlockKey'?: string;
  };

  /**
   * `GET /swarm/unlockkey`
   * 
   * Code 500: server error
   */
  type GetSwarmUnlockkeyResponse500 = ErrorResponse;

  /**
   * `GET /swarm/unlockkey`
   * 
   * Code 503: node is not part of a swarm
   */
  type GetSwarmUnlockkeyResponse503 = ErrorResponse;

  /**
   * `POST /swarm/unlock`
   * 
   * Code 500: server error
   */
  type PostSwarmUnlockResponse500 = ErrorResponse;

  /**
   * `POST /swarm/unlock`
   * 
   * Code 503: node is not part of a swarm
   */
  type PostSwarmUnlockResponse503 = ErrorResponse;

  /**
   * `GET /services`
   * 
   * Code 200: no error
   */
  type GetServiceListResponse200 = Array<Service>;

  /**
   * `GET /services`
   * 
   * Code 500: server error
   */
  type GetServiceListResponse500 = ErrorResponse;

  /**
   * `GET /services`
   * 
   * Code 503: node is not part of a swarm
   */
  type GetServiceListResponse503 = ErrorResponse;

  /**
   * `POST /services/create`
   * 
   * Code 201: no error
   */
  type PostServiceCreateResponse201 = {
    /**
     * The ID of the created service.
     */
    'ID'?: string;
    /**
     * Optional warning message
     */
    'Warning'?: string;
  };

  /**
   * `POST /services/create`
   * 
   * Code 400: bad parameter
   */
  type PostServiceCreateResponse400 = ErrorResponse;

  /**
   * `POST /services/create`
   * 
   * Code 403: network is not eligible for services
   */
  type PostServiceCreateResponse403 = ErrorResponse;

  /**
   * `POST /services/create`
   * 
   * Code 409: name conflicts with an existing service
   */
  type PostServiceCreateResponse409 = ErrorResponse;

  /**
   * `POST /services/create`
   * 
   * Code 500: server error
   */
  type PostServiceCreateResponse500 = ErrorResponse;

  /**
   * `POST /services/create`
   * 
   * Code 503: node is not part of a swarm
   */
  type PostServiceCreateResponse503 = ErrorResponse;

  /**
   * `GET /services/{id}`
   * 
   * Code 200: no error
   */
  type GetServiceInspectResponse200 = Service;

  /**
   * `GET /services/{id}`
   * 
   * Code 404: no such service
   */
  type GetServiceInspectResponse404 = ErrorResponse;

  /**
   * `GET /services/{id}`
   * 
   * Code 500: server error
   */
  type GetServiceInspectResponse500 = ErrorResponse;

  /**
   * `GET /services/{id}`
   * 
   * Code 503: node is not part of a swarm
   */
  type GetServiceInspectResponse503 = ErrorResponse;

  /**
   * `DELETE /services/{id}`
   * 
   * Code 404: no such service
   */
  type DeleteServiceDeleteResponse404 = ErrorResponse;

  /**
   * `DELETE /services/{id}`
   * 
   * Code 500: server error
   */
  type DeleteServiceDeleteResponse500 = ErrorResponse;

  /**
   * `DELETE /services/{id}`
   * 
   * Code 503: node is not part of a swarm
   */
  type DeleteServiceDeleteResponse503 = ErrorResponse;

  /**
   * `POST /services/{id}/update`
   * 
   * Code 200: no error
   */
  type PostServiceUpdateResponse200 = ServiceUpdateResponse;

  /**
   * `POST /services/{id}/update`
   * 
   * Code 400: bad parameter
   */
  type PostServiceUpdateResponse400 = ErrorResponse;

  /**
   * `POST /services/{id}/update`
   * 
   * Code 404: no such service
   */
  type PostServiceUpdateResponse404 = ErrorResponse;

  /**
   * `POST /services/{id}/update`
   * 
   * Code 500: server error
   */
  type PostServiceUpdateResponse500 = ErrorResponse;

  /**
   * `POST /services/{id}/update`
   * 
   * Code 503: node is not part of a swarm
   */
  type PostServiceUpdateResponse503 = ErrorResponse;

  /**
   * `GET /services/{id}/logs`
   * 
   * Get `stdout` and `stderr` logs from a service. See also
   * [`/containers/{id}/logs`](#operation/ContainerLogs).
   * 
   * **Note**: This endpoint works only for services with the `local`,
   * `json-file` or `journald` logging drivers.
   * 
   * Code 200: logs returned as a stream in response body
   */
  type GetServiceLogsResponse200 = string;

  /**
   * `GET /services/{id}/logs`
   * 
   * Get `stdout` and `stderr` logs from a service. See also
   * [`/containers/{id}/logs`](#operation/ContainerLogs).
   * 
   * **Note**: This endpoint works only for services with the `local`,
   * `json-file` or `journald` logging drivers.
   * 
   * Code 404: no such service
   */
  type GetServiceLogsResponse404 = ErrorResponse;

  /**
   * `GET /services/{id}/logs`
   * 
   * Get `stdout` and `stderr` logs from a service. See also
   * [`/containers/{id}/logs`](#operation/ContainerLogs).
   * 
   * **Note**: This endpoint works only for services with the `local`,
   * `json-file` or `journald` logging drivers.
   * 
   * Code 500: server error
   */
  type GetServiceLogsResponse500 = ErrorResponse;

  /**
   * `GET /services/{id}/logs`
   * 
   * Get `stdout` and `stderr` logs from a service. See also
   * [`/containers/{id}/logs`](#operation/ContainerLogs).
   * 
   * **Note**: This endpoint works only for services with the `local`,
   * `json-file` or `journald` logging drivers.
   * 
   * Code 503: node is not part of a swarm
   */
  type GetServiceLogsResponse503 = ErrorResponse;

  /**
   * `GET /tasks`
   * 
   * Code 200: no error
   */
  type GetTaskListResponse200 = Array<Task>;

  /**
   * `GET /tasks`
   * 
   * Code 500: server error
   */
  type GetTaskListResponse500 = ErrorResponse;

  /**
   * `GET /tasks`
   * 
   * Code 503: node is not part of a swarm
   */
  type GetTaskListResponse503 = ErrorResponse;

  /**
   * `GET /tasks/{id}`
   * 
   * Code 200: no error
   */
  type GetTaskInspectResponse200 = Task;

  /**
   * `GET /tasks/{id}`
   * 
   * Code 404: no such task
   */
  type GetTaskInspectResponse404 = ErrorResponse;

  /**
   * `GET /tasks/{id}`
   * 
   * Code 500: server error
   */
  type GetTaskInspectResponse500 = ErrorResponse;

  /**
   * `GET /tasks/{id}`
   * 
   * Code 503: node is not part of a swarm
   */
  type GetTaskInspectResponse503 = ErrorResponse;

  /**
   * `GET /tasks/{id}/logs`
   * 
   * Get `stdout` and `stderr` logs from a task.
   * See also [`/containers/{id}/logs`](#operation/ContainerLogs).
   * 
   * **Note**: This endpoint works only for services with the `local`,
   * `json-file` or `journald` logging drivers.
   * 
   * Code 200: logs returned as a stream in response body
   */
  type GetTaskLogsResponse200 = string;

  /**
   * `GET /tasks/{id}/logs`
   * 
   * Get `stdout` and `stderr` logs from a task.
   * See also [`/containers/{id}/logs`](#operation/ContainerLogs).
   * 
   * **Note**: This endpoint works only for services with the `local`,
   * `json-file` or `journald` logging drivers.
   * 
   * Code 404: no such task
   */
  type GetTaskLogsResponse404 = ErrorResponse;

  /**
   * `GET /tasks/{id}/logs`
   * 
   * Get `stdout` and `stderr` logs from a task.
   * See also [`/containers/{id}/logs`](#operation/ContainerLogs).
   * 
   * **Note**: This endpoint works only for services with the `local`,
   * `json-file` or `journald` logging drivers.
   * 
   * Code 500: server error
   */
  type GetTaskLogsResponse500 = ErrorResponse;

  /**
   * `GET /tasks/{id}/logs`
   * 
   * Get `stdout` and `stderr` logs from a task.
   * See also [`/containers/{id}/logs`](#operation/ContainerLogs).
   * 
   * **Note**: This endpoint works only for services with the `local`,
   * `json-file` or `journald` logging drivers.
   * 
   * Code 503: node is not part of a swarm
   */
  type GetTaskLogsResponse503 = ErrorResponse;

  /**
   * `GET /secrets`
   * 
   * Code 200: no error
   */
  type GetSecretListResponse200 = Array<Secret>;

  /**
   * `GET /secrets`
   * 
   * Code 500: server error
   */
  type GetSecretListResponse500 = ErrorResponse;

  /**
   * `GET /secrets`
   * 
   * Code 503: node is not part of a swarm
   */
  type GetSecretListResponse503 = ErrorResponse;

  /**
   * `POST /secrets/create`
   * 
   * Code 201: no error
   */
  type PostSecretCreateResponse201 = IdResponse;

  /**
   * `POST /secrets/create`
   * 
   * Code 409: name conflicts with an existing object
   */
  type PostSecretCreateResponse409 = ErrorResponse;

  /**
   * `POST /secrets/create`
   * 
   * Code 500: server error
   */
  type PostSecretCreateResponse500 = ErrorResponse;

  /**
   * `POST /secrets/create`
   * 
   * Code 503: node is not part of a swarm
   */
  type PostSecretCreateResponse503 = ErrorResponse;

  /**
   * `GET /secrets/{id}`
   * 
   * Code 200: no error
   */
  type GetSecretInspectResponse200 = Secret;

  /**
   * `GET /secrets/{id}`
   * 
   * Code 404: secret not found
   */
  type GetSecretInspectResponse404 = ErrorResponse;

  /**
   * `GET /secrets/{id}`
   * 
   * Code 500: server error
   */
  type GetSecretInspectResponse500 = ErrorResponse;

  /**
   * `GET /secrets/{id}`
   * 
   * Code 503: node is not part of a swarm
   */
  type GetSecretInspectResponse503 = ErrorResponse;

  /**
   * `DELETE /secrets/{id}`
   * 
   * Code 404: secret not found
   */
  type DeleteSecretDeleteResponse404 = ErrorResponse;

  /**
   * `DELETE /secrets/{id}`
   * 
   * Code 500: server error
   */
  type DeleteSecretDeleteResponse500 = ErrorResponse;

  /**
   * `DELETE /secrets/{id}`
   * 
   * Code 503: node is not part of a swarm
   */
  type DeleteSecretDeleteResponse503 = ErrorResponse;

  /**
   * `POST /secrets/{id}/update`
   * 
   * Code 400: bad parameter
   */
  type PostSecretUpdateResponse400 = ErrorResponse;

  /**
   * `POST /secrets/{id}/update`
   * 
   * Code 404: no such secret
   */
  type PostSecretUpdateResponse404 = ErrorResponse;

  /**
   * `POST /secrets/{id}/update`
   * 
   * Code 500: server error
   */
  type PostSecretUpdateResponse500 = ErrorResponse;

  /**
   * `POST /secrets/{id}/update`
   * 
   * Code 503: node is not part of a swarm
   */
  type PostSecretUpdateResponse503 = ErrorResponse;

  /**
   * `GET /configs`
   * 
   * Code 200: no error
   */
  type GetConfigListResponse200 = Array<Config>;

  /**
   * `GET /configs`
   * 
   * Code 500: server error
   */
  type GetConfigListResponse500 = ErrorResponse;

  /**
   * `GET /configs`
   * 
   * Code 503: node is not part of a swarm
   */
  type GetConfigListResponse503 = ErrorResponse;

  /**
   * `POST /configs/create`
   * 
   * Code 201: no error
   */
  type PostConfigCreateResponse201 = IdResponse;

  /**
   * `POST /configs/create`
   * 
   * Code 409: name conflicts with an existing object
   */
  type PostConfigCreateResponse409 = ErrorResponse;

  /**
   * `POST /configs/create`
   * 
   * Code 500: server error
   */
  type PostConfigCreateResponse500 = ErrorResponse;

  /**
   * `POST /configs/create`
   * 
   * Code 503: node is not part of a swarm
   */
  type PostConfigCreateResponse503 = ErrorResponse;

  /**
   * `GET /configs/{id}`
   * 
   * Code 200: no error
   */
  type GetConfigInspectResponse200 = Config;

  /**
   * `GET /configs/{id}`
   * 
   * Code 404: config not found
   */
  type GetConfigInspectResponse404 = ErrorResponse;

  /**
   * `GET /configs/{id}`
   * 
   * Code 500: server error
   */
  type GetConfigInspectResponse500 = ErrorResponse;

  /**
   * `GET /configs/{id}`
   * 
   * Code 503: node is not part of a swarm
   */
  type GetConfigInspectResponse503 = ErrorResponse;

  /**
   * `DELETE /configs/{id}`
   * 
   * Code 404: config not found
   */
  type DeleteConfigDeleteResponse404 = ErrorResponse;

  /**
   * `DELETE /configs/{id}`
   * 
   * Code 500: server error
   */
  type DeleteConfigDeleteResponse500 = ErrorResponse;

  /**
   * `DELETE /configs/{id}`
   * 
   * Code 503: node is not part of a swarm
   */
  type DeleteConfigDeleteResponse503 = ErrorResponse;

  /**
   * `POST /configs/{id}/update`
   * 
   * Code 400: bad parameter
   */
  type PostConfigUpdateResponse400 = ErrorResponse;

  /**
   * `POST /configs/{id}/update`
   * 
   * Code 404: no such config
   */
  type PostConfigUpdateResponse404 = ErrorResponse;

  /**
   * `POST /configs/{id}/update`
   * 
   * Code 500: server error
   */
  type PostConfigUpdateResponse500 = ErrorResponse;

  /**
   * `POST /configs/{id}/update`
   * 
   * Code 503: node is not part of a swarm
   */
  type PostConfigUpdateResponse503 = ErrorResponse;

  /**
   * `GET /distribution/{name}/json`
   * 
   * Return image digest and platform information by contacting the registry.
   * 
   * Code 200: descriptor and platform information
   */
  type GetDistributionInspectResponse200 = DistributionInspect;

  /**
   * `GET /distribution/{name}/json`
   * 
   * Return image digest and platform information by contacting the registry.
   * 
   * Code 401: Failed authentication or no image found
   */
  type GetDistributionInspectResponse401 = ErrorResponse;

  /**
   * `GET /distribution/{name}/json`
   * 
   * Return image digest and platform information by contacting the registry.
   * 
   * Code 500: Server error
   */
  type GetDistributionInspectResponse500 = ErrorResponse;

  /**
   * `POST /session`
   * 
   * Start a new interactive session with a server. Session allows server to
   * call back to the client for advanced capabilities.
   * 
   * ### Hijacking
   * 
   * This endpoint hijacks the HTTP connection to HTTP2 transport that allows
   * the client to expose gPRC services on that connection.
   * 
   * For example, the client sends this request to upgrade the connection:
   * 
   * ```
   * POST /session HTTP/1.1
   * Upgrade: h2c
   * Connection: Upgrade
   * ```
   * 
   * The Docker daemon responds with a `101 UPGRADED` response follow with
   * the raw stream:
   * 
   * ```
   * HTTP/1.1 101 UPGRADED
   * Connection: Upgrade
   * Upgrade: h2c
   * ```
   * 
   * Code 400: bad parameter
   */
  type PostSessionResponse400 = ErrorResponse;

  /**
   * `POST /session`
   * 
   * Start a new interactive session with a server. Session allows server to
   * call back to the client for advanced capabilities.
   * 
   * ### Hijacking
   * 
   * This endpoint hijacks the HTTP connection to HTTP2 transport that allows
   * the client to expose gPRC services on that connection.
   * 
   * For example, the client sends this request to upgrade the connection:
   * 
   * ```
   * POST /session HTTP/1.1
   * Upgrade: h2c
   * Connection: Upgrade
   * ```
   * 
   * The Docker daemon responds with a `101 UPGRADED` response follow with
   * the raw stream:
   * 
   * ```
   * HTTP/1.1 101 UPGRADED
   * Connection: Upgrade
   * Upgrade: h2c
   * ```
   * 
   * Code 500: server error
   */
  type PostSessionResponse500 = ErrorResponse;

}

