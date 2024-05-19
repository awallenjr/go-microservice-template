# go-microservice-template
A simple starter template for microservices in GO.

## Install Docker and Kubernetes Components on Ubuntu server

### Kubernetes Installation

Kubernetes is an open-source container orchestration platform used for automating the deployment, scaling, and management of containerized applications.

Click here for the [Online tutorial I found very useful](https://www.cherryservers.com/blog/install-kubernetes-on-ubuntu)

## Uninstalling Docker and Kubernetes

### Uninstall Kubernetes

1. Reset Kubernetes cluster:

    ```
    sudo kubeadm reset -f
    ```

2. Clean up any residual directories and files:

    ```
    sudo rm -rf /etc/kubernetes /etc/cni /etc/containerd /var/lib/etcd /var/lib/kubelet ~/.kube /var/lib/dockershim /var/run/kubernetes
    ```

4. Remove kubelet, kubeadm, and kubectl packages:

    ```
    sudo apt-get purge -y kubelet kubeadm kubectl
    ```

5. Remove unused dependencies:

    ```
    sudo apt-get autoremove -y
    ```

### Uninstall Docker
1. Remove Docker packages:

    ```
    sudo apt-get purge -y docker-ce docker-ce-cli containerd.io
    ```

2. Remove Docker related configuration and files:

    ```
    sudo rm -rf /var/lib/docker /etc/docker
    ```

3. Remove Docker GPG key and repository:

    ```
    sudo rm -rf /etc/apt/sources.list.d/docker.list /etc/apt/trusted.gpg.d/docker.gpg
    ```

4. Remove unused dependencies:

    ```
    sudo apt-get autoremove -y
    ```

## Common Useful Commands

1. kubectl

**kubectl is the primary command-line tool for interacting with Kubernetes clusters.**

**Viewing Resources:**

    ```
    kubectl get pods - List all pods in the current namespace.
    kubectl get nodes - List all nodes in the cluster.
    kubectl get deployments - List all deployments.
    kubectl get services - List all services.
    ```

**Detailed Information:**

    ```
    kubectl describe pod <pod_name> - Get detailed information about a specific pod.
    kubectl describe node <node_name> - Get detailed information about a specific node.
    ```

**Logs:**

    ```
    kubectl logs <pod_name> - Retrieve the logs of a specific pod.
    kubectl logs -f <pod_name> - Stream the logs of a specific pod.
    ```

**Exec into a Pod:**

    ```
    kubectl exec -it <pod_name> -- /bin/bash - Execute an interactive shell session in a pod.
    ```

2. Monitoring and Metrics

**Kubernetes provides metrics through its Metrics API, and you can use tools to visualize and analyze these metrics.**

**Metrics Server:**

    ```
    kubectl top node - Show metrics for all nodes.
    kubectl top pod - Show metrics for all pods in the current namespace.
    ```

3. Cluster-wide Status and Health

    ```
    kubectl cluster-info - Display cluster info such as Kubernetes master and services.
    kubectl get events - View cluster events.
    ```

4. Advanced Troubleshooting

    ```
    kubectl describe - Get detailed information about Kubernetes resources.
    kubectl explain - Get help for a specific resource schema.
    ```