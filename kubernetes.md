# Learn Kubernetes Basics
### Objective
- Deploy a containerized application on a cluster
- Scale the deployment
- Update the containerized application with a new software version
- Debug the containerized application

### Using Minikube to Create a Cluster
**Kubernetes coordinates a highly available cluster of computers that are connected to work as a single unit.** The abstractions in Kubernetes allow you to deploy containerized applications to a cluster without tying them specifically to individual machines. To make use of this new model of deployment, applications need to be packaged in a way that decouples them from individual hosts: they need to be containerized. Containerized applications are more flexible and available than in past deployment models, where applications were installed directly onto specific machines as packages deeply integrated into the host. **Kubernetes automates the distribution and scheduling of application containers across a cluster in a more efficient way.** Kubernetes is an open-source platform and is production-ready.

A Kubernetes cluster consists of two types of resources:
- The **Master** coordinates the cluster. The master coordinates all activities in your cluster, such as scheduling applications, maintaining applications' desired state, scaling applications, and rolling out new updates.
- **Nodes** are the workers that run applications. Each node has a Kubelet, which is an agent for managing the node and communicating with the Kubernetes master.

<img src="https://d33wubrfki0l68.cloudfront.net/99d9808dcbf2880a996ed50d308a186b5900cec9/40b94/docs/tutorials/kubernetes-basics/public/images/module_01_cluster.svg">

#### Minikube
 Minikube is a lightweight Kubernetes implementation that creates a VM on your local machine and deploys a simple cluster containing only one node.

 1. ``minikube version`` ensure that minikube is installed and ``minikube start`` start the cluster.
 2. Get cluster details by running ``kubectl cluster-info`` and to view the nodes on the cluster use ``kubectl get nodes``

 Example output

 `Kubernetes master is running at https://192.168.99.100:8443
KubeDNS is running at https://192.168.99.100:8443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
`

### Using kubectl to Create a Deployment
Once you have a running Kubernetes cluster, you can deploy your containerized applications on top of it. To do so, you create a Kubernetes Deployment configuration.  The Deployment instructs Kubernetes how to create and update instances of your application. Once you've created a Deployment, the Kubernetes master schedules mentioned application instances onto individual Nodes in the cluster. Once the application instances are created, a Kubernetes Deployment Controller continuously monitors those instances. If the Node hosting an instance goes down or is deleted, the Deployment controller replaces the instance with an instance on another Node in the cluster. This provides a self-healing mechanism to address machine failure or maintenance.

<img src="https://d33wubrfki0l68.cloudfront.net/152c845f25df8e69dd24dd7b0836a289747e258a/4a1d2/docs/tutorials/kubernetes-basics/public/images/module_02_first_app.svg">

Applications need to be packaged into one of the supported container formats in order to be deployed on Kubernetes. You can create and manaCreate a Deploymentge a Deployment by using the Kubernetes command line interface, Kubectl. Kubectl uses the Kubernetes API to interact with the cluster.
#### Create a Deployment
1. A Kubernetes Pod is a group of one or more Containers, tied together for the purposes of administration and networking. The Pod in this tutorial has only one Container. A Kubernetes Deployment checks on the health of your Pod and restarts the Pod’s Container if it terminates. Deployments are the recommended way to manage the creation and scaling of Pods.
``` kubectl create deployment hello-node --image=gcr.io/hello-minikube-zero-install/hello-node ```
2. To get the deployment we just created. ``kubectl get deployments`` and to get Pod information use ``kubectl get pods``. To get events use ``kubectl get events``. And finaly to view kubectl configuration use ``kubectl config view``

#### Create a Service
By default, the Pod is only accessible by its internal IP address within the Kubernetes cluster. To make the **hello-node** Container accessible from outside the Kubernetes virtual network, you have to expose the Pod as a Kubernetes Service.

Example

``kubectl expose deployment hello-node --type=LoadBalancer --port=8080``

The **--type=LoadBalancer** flag indicates that you want to expose your Service outside of the cluster. One can view created services using command ``kubectl get services``
