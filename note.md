0. Basic structure and workflow definitions for `container` terminology/jargon:

- Definition:

  > - Container is an environment for executing processes with configurable isolation and resource limitations.
  > - The term "configurable" refers to the meaning that the resource requirements are absolutely controlled by the configure themself.

- `OCI`: stands for "Open Container Initiative" and is an independent organization that delivers a schema of the standard of a qualified container.

- `Container`: can infer as a box that contains a bunch of items, all of these items being instances of the processes that are isolated from the host system's procedures.

> - Isolation:
>   | Namespaces | Filesystem |
>   |:-|:-|
>   | uts / mount / pid / ipc / net / user / cgroup | chroot / pivot_root |

1. Simple structure of a `container`:

- `Namespaces`: created with `syscalls`, limited what was the process (container) can see,
  including: (also can be understood in the way the restrictions you applied to the container
  and forced it to follow)

  > - Unix Timesharing System
  > - Process IDs
  > - Mounts
  > - Network
  > - User IDs
  > - InterProcess Communications

- `Chroot`: Changing Root
- `Cgroups`: Control Groups
