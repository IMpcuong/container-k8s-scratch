0. Basic architecture:

- `Pod`: is the basic building block of Kubernetes - the smallest and simplest unit in the Kubernetes object model that you create or deploy.
  A `Pod` represents a running process on your cluster.

- The `one-container-per-Pod` model is the most common Kubernetes use case.
  In this case, you can think of a `Pod` as a wrapper around a single container, and Kubernetes manages the Pods rather than the containers directly.

- `minikube` cluster installation: `minikube start --driver=docker --cpus=2 --memory=2000m --disk-size="10000mb"`.

1. `kubectl` command's examples:

- Comamnd's instructions:

```bash
kubectl <command> --help
kubectl options
```

- `kubectl run` used to build image as a pod in an imperative way (not using declaraitve file):

  ```bash
  kubectl run coffee-container --image=ansilh/demo-coffee --restart=Never
  ```

- `kubectl apply` or `kubectl create --save-config` examples:

  - Usage: `kubectl apply (-f FILENAME | -k DIRECTORY) [options]`

  ```bash
  # Apply the configuration in pod.json to a pod.
  kubectl apply -f ./pod.json
  kubectl apply -f ./pod.yaml
  kubectl apply -f pod.yaml view-last-applied
  kubectl apply -f pod.yaml set-last-applied
  kubectl apply -f pod.yaml edit-last-applied

  # Apply resources from a directory containing kustomization.yaml - e.g. dir/kustomization.yaml.
  kubectl apply -k dir/

  # Apply the JSON passed into stdin to a pod
  cat pod.json | kubectl apply -f -

  # Note: --prune is still in Alpha-version.
  # Apply the configuration in manifest.yaml that matches label app=nginx and delete all other resources that are not in the file and match label app=nginx.
  kubectl apply --prune -f manifest.yaml -l app=nginx

  # Apply the configuration in manifest.yaml and delete all the other config maps that are not in the file.
  kubectl apply --prune -f manifest.yaml --all --prune-whitelist=core/v1/ConfigMap
  ```

- `kubectl delete` examples:

  - Usage: `kubectl delete ([-f FILENAME] | [-k DIRECTORY] | TYPE [(NAME | -l label | --all)]) [options]`

  ```bash
  # Delete a pod using the type and name specified in pod.json.
  kubectl delete -f ./pod.json

  # Delete resources from a directory containing kustomization.yaml - e.g. dir/kustomization.yaml.
  kubectl delete -k dir

  # Delete a pod based on the type and name in the JSON passed into stdin.
  cat pod.json | kubectl delete -f -

  # Delete pods and services with same names "baz" and "foo".
  kubectl delete pod,service baz foo

  # Delete pods and services with label name=myLabel.
  kubectl delete pods,services -l name=myLabel

  # Delete a pod with minimal delay.
  kubectl delete pod foo --now

  # Force delete a pod on a dead node.
  kubectl delete pod foo --force

  # Delete all pods.
  kubectl delete pods --all
  ```

- `kubectl exec` login to specific pod (the wrapper of container):

  ```bash
  kubectl exec -it centos-pod -- /bin/bash
  ```

- `kubectl label` or `kubectl annotate` used to labelize each pod to make them becoming an unique identifier:

  ```bash
  # Add custom label to pod.
  kubectl label pod coffee-container app=frontend

  # Remove label from pod.
  kubectl label pod coffee-container app-

  # Bind custom annotation into pod.
  # The metadata in an annotation can be small or large, structured or unstructured, and can be included characters which is not permitted by labels.
  kubectl annotate pod coffee-app url=https://hub.docker.com/r/IMpcuong/test-image

  # Decouple the existed annotation from pod.
  kubectl annotate pod coffee-app url-

  # New applied labels or annotations checker.
  kubectl describe pod coffee-app
  ```

- `kubectl taint` disables the given node from consuming pods and only accepts pods that contain the corresponding taint.

  ```bash
  # Format pattern: `key=value:Effect`.
  #
  # `Effect := [NoSchedule, PreferNoSchedule, NoExecute]`.
  # + NoSchedule ~ Pods are not going to be scheduled.
  # + PreferNoSchedule ~ This is a “preference” or “soft” version of NoSchedule; the system will try to avoid placing a pod that does not tolerate the taint on the node, but it is not required.
  # + NoExecute ~ pod will be evicted from the node (if it is already running on the node), and will not be scheduled onto the node (if it is not yet running on the node).
  kubectl taint node minikube node-role.kubernetes.io/control-plane="":NoSchedule
  kubectl taint node minikube node-role.kubernetes.io/control-plane="":NoExecute
  ```

- `kubectl explain` the `man-page` for Kubernetes' specifications and components:

  ```bash
  kubectl explain Pod
  kubectl explain Pod.spec
  kubectl explain Pod.spec.containers
  ```

- `kubectl get` or `kubectl describe` services/system's current status:

  ```bash
  kubectl describe nodes minikube
  kubectl describe node kube-system

  kubectl get all -o wide
  kubetcl get pods -o wide --show-labels
  kubectl get pods -n default
  kubectl get ns # kubectl get namespaces
  kubectl get all -n kube-system

  # Selecting a pod using its own unique label.
  kubectl get pods --selector=app=frontend

  # Error: Metrics API not available.
  # --> Needs to install the required metrics to execute this command.
  kubectl top node
  kubectl top pod -A
  ```
