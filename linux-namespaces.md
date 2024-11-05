# Linux Namespaces in Kubernetes

Linux namespaces are a fundamental feature that enable container isolation by partitioning system resources to appear as if they are separate instances. Here are the main types of namespaces used in Kubernetes:

## 1. Network Namespace (net)
- Isolates network-related resources:
  - Network interfaces
  - IP routing tables
  - Firewall rules
  - Port numbers
  - `/proc/net` directory
```bash
# Example: Create a new network namespace
ip netns add my_netns
# List network namespaces
ip netns list
```

## 2. Mount Namespace (mnt)
- Isolates filesystem mount points
- Each container gets its own root filesystem view
- Enables containers to have different filesystem layouts
```bash
# Example: Run process in new mount namespace
unshare --mount /bin/bash
# View mount points
mount
```

## 3. PID Namespace (pid)
- Isolates process IDs
- Processes in different namespaces can have same PID
- Container processes can't see host processes
```bash
# Example: Create new PID namespace
unshare --pid --fork /bin/bash
# View processes
ps aux
```

## 4. UTS Namespace (uts)
- Isolates hostname and domain name
- Allows each container to have its own hostname
```bash
# Example: Create new UTS namespace
unshare --uts /bin/bash
# Change hostname
hostname container1
```

## 5. IPC Namespace (ipc)
- Isolates InterProcess Communication resources
- Separates shared memory segments, semaphores, message queues
```bash
# Example: Create new IPC namespace
unshare --ipc /bin/bash
# View IPC resources
ipcs
```

## 6. User Namespace (user)
- Maps container user IDs to different host user IDs
- Allows privileged operations inside container without host privileges
```bash
# Example: Create new user namespace
unshare --user /bin/bash
# View user mapping
cat /proc/$$/uid_map
```

## 7. Cgroup Namespace (cgroup)
- Isolates cgroup root directory
- Container processes see their cgroup as root
```bash
# Example: Create new cgroup namespace
unshare --cgroup /bin/bash
# View cgroups
cat /proc/self/cgroup
```

## 8. Time Namespace (time)
- Isolates system clocks
- Allows containers to have different system times
- Newest namespace type
```bash
# Example: Create new time namespace (requires Linux 5.6+)
unshare --time /bin/bash
# View system time
date
```

## Usage in Kubernetes

Kubernetes uses these namespaces in combination to create isolated containers:

1. **Pod Level**: Containers within a pod share some namespaces:
   - Network namespace (enabling localhost communication)
   - UTS namespace
   - IPC namespace

2. **Container Level**: Each container has its own:
   - Mount namespace (for filesystem isolation)
   - PID namespace (for process isolation)
   - User namespace (optional, for security)

## Common Operations

```bash
# View process namespace information
ls -l /proc/$$/ns

# Enter a process namespace
nsenter --target $PID --net --mount

# Create container with specific namespaces
docker run --pid=host --net=host nginx

# View namespace information in container
lsns
```

## Security Considerations

1. Root in user namespace != root on host
2. Network namespace isolation is crucial for security
3. Mount namespace prevents access to host filesystem
4. PID namespace hides host processes from containers

These namespaces form the foundation of container isolation in Kubernetes, ensuring that containers run securely and independently while still being able to communicate when needed.
