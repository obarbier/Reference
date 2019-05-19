# Learn Kubernetes Basics
### Objective
- Deploy a containerized application on a cluster
- Scale the deployment
- Update the containerized application with a new software version
- Debug the containerized application
- Learn about application Deployments
- Deploy on Kubernetes with kubectl
- Learn about Kubernetes Pods
- Learn about Kubernetes Nodes
- Troubleshoot deployed applications

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
``` kubectl run kubernetes-bootcamp --image=gcr.io/google-samples/kubernetes-bootcamp:v1 --port=8080 ```
  The run command creates a new deployment. We need to provide the deployment name and app image location (include the full repository url for images hosted outside Docker hub). We want to run the app on a specific port so we add the --port parameter. This performed a few things for you:
  - searched for a suitable node where an instance of the application could be run (we have only 1 available node)
  - scheduled the application to run on that Node
  - configured the cluster to reschedule the instance on a new Node when needed

2. To get the deployment we just created. ``kubectl get deployments`` and to get Pod information use ``kubectl get pods``. To get events use ``kubectl get events``. And finaly to view kubectl configuration use ``kubectl config view``

#### Accessing Using Proxy
Pods that are running inside Kubernetes are running on a private, isolated network. By default they are visible from other pods and services within the same kubernetes cluster, but not outside that network. When we use **kubectl**, we're interacting through an API endpoint to communicate with our application.
The kubectl command can create a proxy that will forward communications into the cluster-wide, private network. The proxy can be terminated by pressing control-C and won't show any output while its running.

Example
In terminal 1 run ``kubectl proxy``

run the following in terminal 2

``
export POD_NAME=$(kubectl get pods -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')
echo Name of the Pod: $POD_NAME
``

``
curl http://localhost:8001/api/v1/namespaces/default/pods/$POD_NAME/proxy/
``

#### Create a Service
By default, the Pod is only accessible by its internal IP address within the Kubernetes cluster. To make the **hello-node** Container accessible from outside the Kubernetes virtual network, you have to expose the Pod as a Kubernetes Service.

Example

``kubectl expose deployment hello-node --type=LoadBalancer --port=8080``

The **--type=LoadBalancer** flag indicates that you want to expose your Service outside of the cluster. One can view created services using command ``kubectl get services``

### how to apply all of the K8s objects into the k8s cluster
Reference: [Stack Overflow](https://stackoverflow.com/questions/48015637/kubernetes-kubectl-run-vs-create-and-apply)
There are several ways to do this job.
- Using Generators (Run, Expose)
- Using Imperative way (Create)
- Using Declarative way (Apply)

All of the above ways have a different purpose and simplicity. For instance, If you want to check quickly whether the container is working as you desired then you might use **Generators**.

If you want to version control the k8s object then it's better to use **declarative** way which helps us to determine the accuracy of data in k8s objects.

Deployment, ReplicaSet and Pods are different layers which solve different problems.All of these concepts provide flexibility to k8s.

- Pods: It makes sure that related containers are together and provide efficiency.
- ReplicaSet: It makes sure that k8s cluster has desirable replicas of the pods
- Deployment: It makes sure that you can have different version of Pods and provide the capability to rollback to the previous version

Lastly, It depends on use case how you want to use these concepts or methodology. It's not about which is good or which is bad.

### Viewing Pods and Nodes

<img src="https://d33wubrfki0l68.cloudfront.net/fe03f68d8ede9815184852ca2a4fd30325e5d15a/98064/docs/tutorials/kubernetes-basics/public/images/module_03_pods.svg">
A **Pod**  is the atomic unit on the Kubernetes platform. It is a group of one or more application containers (such as Docker or rkt) and includes shared storage (volumes), IP address and information about how to run them. A Pod models an application-specific "logical host" and can contain different application containers which are relatively tightly coupled. When we create a Deployment on Kubernetes, that Deployment creates Pods with containers inside them (as opposed to creating containers directly).

A Pod always runs on a Node. A Node is a worker machine in Kubernetes and may be either a virtual or a physical machine, depending on the cluster. Each Node is managed by the Master. A Node can have multiple pods, and the Kubernetes master automatically handles scheduling the pods across the Nodes in the cluster. The Master's automatic scheduling takes into account the available resources on each Node.

<img src="https://d33wubrfki0l68.cloudfront.net/5cb72d407cbe2755e581b6de757e0d81760d5b86/a9df9/docs/tutorials/kubernetes-basics/public/images/module_03_nodes.svg">

Every Kubernetes Node runs at least:

- Kubelet, a process responsible for communication between the Kubernetes Master and the Node; it manages the Pods and the containers running on a machine.
- A container runtime (like Docker, rkt) responsible for pulling the container image from a registry, unpacking the container, and running the application.
Containers should only be scheduled together in a single Pod if they are tightly coupled and need to share resources such as disk.
