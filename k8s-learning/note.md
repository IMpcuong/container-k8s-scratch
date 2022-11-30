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

- `kubect apply` or `kubectl create --save-config` examples:

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

- `kubectl get` or `kubectl describe` services/system's current status:

  ```bash
  kubectl describe nodes minikube
  kubectl describe node kube-system

  kubectl get all -o wide
  kubectl get pods -n default
  kubectl get ns # kubectl get namespaces
  kubectl get all -n kube-system

  # Error: Metrics API not available.
  # --> Needs to install the required metrics to execute this command.
  kubectl top node
  kubectl top pod -A
  ```
