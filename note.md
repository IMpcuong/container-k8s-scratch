0. Basic structure and workflow definitions for `container` terminology/jargon: (Linux Container for more specifically)

- Definition:

  > - Container is an environment for executing processes with configurable isolation and resource limitations.
  > - The term "configurable" refers to the meaning that the resource requirements are absolutely controlled by the configure themself.

- `OCI`: stands for "Open Container Initiative" and is an independent organization that delivers a schema of the standard of a qualified container.

- `Container`: can infer as a box that contains a bunch of items, all of these items being instances of the processes that are isolated from the host system's procedures.

> - Isolation: everything inside a container is separated from the environment outside of the box.
>   That means all of the isolated processes are only visible with each other, obviously, inbound of the container's border itself.
>
>   | Namespaces                                    | Filesystem          |
>   | :-------------------------------------------- | :------------------ |
>   | uts / mount / pid / ipc / net / user / cgroup | chroot / pivot_root |

> - Resources control: `CGroups` := Control Groups in the Linux OS.
>
>   | CGgroups                      |
>   | :---------------------------- |
>   | CPU, Memory, IO, PID, Devices |

> - Security:
>
>   - `SELinux` := Security-Enhanced Linux, that provides a mechanism for supporting access control security policies.
>     Including mandatory access control, such as start/stop web services (Apache, Nginx, ...)
>   - `Capabilities` := Linux administrator's privileges partition for minimum capabilities of usage.
>
>   | Linux Security Module      | Capabilities                  | Seccomp |
>   | :------------------------- | :---------------------------- | :------ |
>   | SELinux / Smack / AppArmor | CAP_SYS_ADMIN / CAP_NET_ADMIN |         |

1. Simple structure of a `container`:

- `Namespaces`: created with `syscalls`, limited what was the process (container) can see,
  including: (also can be understood in the way the restrictions you applied to the container
  and forced it to follow)

  > - Unix Timesharing System
  > - Mounts
  > - CGroups
  > - Process IDs
  > - Network
  > - UTS
  > - User IDs
  > - IPC ~ Inter-Process Communications

  NOTE:

  - It's `a mechanism` to `provide processes` with different boundary/view on different system resources.
  - Working with namespace:
    - Create a new namespace: `unshare` ~ also run a process inside the new container/box had just been created.
    - Enter a namespace: `nsenter` ~ adding new processes to an existing namespace.

- `Chroot`: Changing Root
- `CGroups`: Control Groups
