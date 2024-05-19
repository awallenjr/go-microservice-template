# go-microservice-template
A simple starter template for microservices in GO.

## Install Docker and Kubernetes Components on Ubuntu server

### Docker Installation

Docker is a containerization platform that simplifies the packaging and distribution of applications. We will install Docker using the official Docker repository.

1. Update the apt package index and install packages to allow apt to use a repository over HTTPS, and curl for downloading the packages:

    ```bash
    sudo apt-get update
    sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common
    ```

2. Add Docker’s official GPG key:

    ```bash
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
    ```

3. Add the Docker repository to APT sources:

    ```bash
    echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
    ```

4. Install Docker CE (Community Edition):

    ```bashs
    sudo apt-get update
    sudo apt-get install -y docker-ce
    ```

5. Verify that Docker CE is installed correctly by running the hello-world image:

    ```bash
    sudo docker run hello-world
    ```

### Kubernetes Installation

Kubernetes is an open-source container orchestration platform used for automating the deployment, scaling, and management of containerized applications.

Click here for the [Online tutorial I found very useful](https://www.cherryservers.com/blog/install-kubernetes-on-ubuntu)

1. Disable swap
    ```bash
    sudo swapoff -
    # Keep the swap off after a system reboot
    sudo sed -i '/ swap / s/^/#/' /etc/fstab
    ```

2. Add the Kubernetes repository:

    ```bash
    sudo apt-get update && sudo apt-get install -y apt-transport-https ca-certificates curl
    sudo mkdir /etc/apt/keyrings

    # https://kubernetes.io/blog/2023/08/15/pkgs-k8s-io-introduction/#how-to-migrate
    curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.28/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg
    echo "deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.28/deb/ /" | sudo tee /etc/apt/sources.list.d/kubernetes.list
    ```

3. Update the apt package index:

    ```bash
    sudo apt-get update
    ```

4. Install Kubernetes components (kubelet, kubeadm, kubectl):

    ```bash
    sudo apt-get install -y kubelet kubeadm kubectl
    ```

    - `kubelet`: The primary node agent that runs on each node in the cluster.
    - `kubeadm`: A tool used to bootstrap Kubernetes clusters.
    - `kubectl`: A command-line utility used to interact with the Kubernetes cluster.

5. Verify that the Kubernetes components are installed:

    ```bash
    kubeadm version
    kubectl version --short --client
    ```

## Uninstalling Docker and Kubernetes

### Uninstall Kubernetes

1. Reset Kubernetes cluster:

    ```bash
    sudo kubeadm reset -f
    ```

2. Clean up any residual directories and files:

    ```bash
    sudo rm -rf /etc/kubernetes /etc/cni /etc/containerd /var/lib/etcd /var/lib/kubelet ~/.kube /var/lib/dockershim /var/run/kubernetes
    ```

4. Remove kubelet, kubeadm, and kubectl packages:

    ```bash
    sudo apt-get purge -y kubelet kubeadm kubectl
    ```

5. Remove unused dependencies:

    ```bash
    sudo apt-get autoremove -y
    ```

### Uninstall Docker
1. Remove Docker packages:

    ```bash
    sudo apt-get purge -y docker-ce docker-ce-cli containerd.io
    ```

2. Remove Docker related configuration and files:

    ```bash
    sudo rm -rf /var/lib/docker /etc/docker
    ```

3. Remove Docker GPG key and repository:

    ```bash
    sudo rm -rf /etc/apt/sources.list.d/docker.list /etc/apt/trusted.gpg.d/docker.gpg
    ```

4. Remove unused dependencies:

    ```bash
    sudo apt-get autoremove -y
    ```