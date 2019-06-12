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
- Learn about a Service in Kubernetes
- Understand how labels and LabelSelector objects relate to a Service
- Expose an application outside a Kubernetes cluster using a Service


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

<img src="https://d33wubrfki0l68.cloudfronat.net/fe03f68d8ede9815184852ca2a4fd30325e5d15a/98064/docs/tutorials/kubernetes-basics/public/images/module_03_pods.svg">
A **Pod**  is the atomic unit on the Kubernetes platform. It is a group of one or more application containers (such as Docker or rkt) and includes shared storage (volumes), IP address and information about how to run them. A Pod models an application-specific "logical host" and can contain different application containers which are relatively tightly coupled. When we create a Deployment on Kubernetes, that Deployment creates Pods with containers inside them (as opposed to creating containers directly).

A Pod always runs on a Node. A **Node** is a worker machine in Kubernetes and may be either a virtual or a physical machine, depending on the cluster. Each Node is managed by the Master. A Node can have multiple pods, and the Kubernetes master automatically handles scheduling the pods across the Nodes in the cluster. The Master's automatic scheduling takes into account the available resources on each Node.

<img src="https://d33wubrfki0l68.cloudfront.net/5cb72d407cbe2755e581b6de757e0d81760d5b86/a9df9/docs/tutorials/kubernetes-basics/public/images/module_03_nodes.svg">

Every Kubernetes Node runs at least:

- Kubelet, a process responsible for communication between the Kubernetes Master and the Node; it manages the Pods and the containers running on a machine.
- A container runtime (like Docker, rkt) responsible for pulling the container image from a registry, unpacking the container, and running the application.
Containers should only be scheduled together in a single Pod if they are tightly coupled and need to share resources such as disk.

The most common operations can be done with the following kubectl commands:
  - **kubectl get** - list resources
  - **kubectl describe** - show detailed information about a resource
  - **kubectl logs** - print the logs from a container in a pod
  - **kubectl exec** - execute a command on a container in a pod

### Example

Let’s verify that the application we deployed in the above scenario is running. We’ll use the **kubectl get** command and look for existing Pods:
```
olivier@obarbier:~$ kubectl get pods
NAME                                  READY   STATUS    RESTARTS   AGE
hello-node-78cd77d68f-hbjcj           1/1     Running   1          5d23h
kubernetes-bootcamp                   1/1     Running   1          5d23h
kubernetes-bootcamp-b94cb9bff-tf9sg   1/1     Running   1          5d23h
 ```
 Now to view container inside a pod and what image are used to build those containers we used **kubectl describe pods**. We see details about the Pod’s container: IP address, the ports used and a list of events related to the lifecycle of the Pod. An example of one of the container is
```
olivier@obarbier:~$ kubectl describe pods
Name:               hello-node-78cd77d68f-hbjcj
Namespace:          default
Priority:           0
PriorityClassName:  <none>
Node:               minikube/10.0.2.15
Start Time:         Sun, 19 May 2019 12:00:06 -0400
Labels:             app=hello-node
                    pod-template-hash=78cd77d68f
Annotations:        <none>
Status:             Running
IP:                 172.17.0.2
Controlled By:      ReplicaSet/hello-node-78cd77d68f
Containers:
  hello-node:
    Container ID:   docker://96bb1b59ada08edd8b24c128b6eb2a4c58bfa475a2cb16027c05d6da32438e2f
    Image:          gcr.io/hello-minikube-zero-install/hello-node
    Image ID:       docker-pullable://gcr.io/hello-minikube-zero-install/hello-node@sha256:9cf82733f7278ae7ae899d432f8d3b3bb0fcb54e673c67496a9f76bb58f30a1c
    Port:           <none>
    Host Port:      <none>
    State:          Running
      Started:      Sat, 25 May 2019 11:22:41 -0400
    Last State:     Terminated
      Reason:       Error
      Exit Code:    255
      Started:      Sun, 19 May 2019 12:01:25 -0400
      Finished:     Sat, 25 May 2019 11:21:54 -0400
    Ready:          True
    Restart Count:  1
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-mz9wk (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             True
  ContainersReady   True
  PodScheduled      True
Volumes:
  default-token-mz9wk:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-mz9wk
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute for 300s
                 node.kubernetes.io/unreachable:NoExecute for 300s
Events:
  Type    Reason          Age    From               Message
  ----    ------          ----   ----               -------
  Normal  SandboxChanged  6m28s  kubelet, minikube  Pod sandbox changed, it will be killed and re-created.
  Normal  Pulling         6m26s  kubelet, minikube  Pulling image "gcr.io/hello-minikube-zero-install/hello-node"
  Normal  Pulled          6m25s  kubelet, minikube  Successfully pulled image "gcr.io/hello-minikube-zero-install/hello-node"
  Normal  Created         6m24s  kubelet, minikube  Created container hello-node
  Normal  Started         6m23s  kubelet, minikube  Started container hello-node
```
Anything that the application would normally send to STDOUT becomes logs for the container within the Pod. We can retrieve these logs using the **kubectl logs** command:
```
olivier@obarbier:~$ kubectl logs $POD_NAME
Kubernetes Bootcamp App Started At: 2019-05-25T15:22:41.153Z | Running On:  kubernetes-bootcamp-b94cb9bff-tf9sg

Running On: kubernetes-bootcamp-b94cb9bff-tf9sg | Total Requests: 1 | App Uptime: 1983.44 seconds | Log Time: 2019-05-25T15:55:44.593Z
```
We can execute commands directly on the container once the Pod is up and running. For this, we use the **exec** command and use the name of the Pod as a parameter. example:

```
olivier@obarbier:~$ kubectl exec $POD_NAME env
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
HOSTNAME=kubernetes-bootcamp-b94cb9bff-tf9sg
KUBERNETES_SERVICE_HOST=10.96.0.1
KUBERNETES_PORT=tcp://10.96.0.1:443
KUBERNETES_PORT_443_TCP_PROTO=tcp
KUBERNETES_PORT_443_TCP_PORT=443
KUBERNETES_PORT_443_TCP_ADDR=10.96.0.1
HELLO_NODE_PORT_8080_TCP=tcp://10.106.103.123:8080
HELLO_NODE_SERVICE_PORT=8080
HELLO_NODE_PORT=tcp://10.106.103.123:8080
HELLO_NODE_PORT_8080_TCP_PORT=8080
KUBERNETES_PORT_443_TCP=tcp://10.96.0.1:443
HELLO_NODE_SERVICE_HOST=10.106.103.123
HELLO_NODE_PORT_8080_TCP_ADDR=10.106.103.123
KUBERNETES_SERVICE_PORT=443
KUBERNETES_SERVICE_PORT_HTTPS=443
HELLO_NODE_PORT_8080_TCP_PROTO=tcp
NPM_CONFIG_LOGLEVEL=info
NODE_VERSION=6.3.1
HOME=/root
```
Or we can start a bash session like so (similarly to docker)
```
olivier@obarbier:~$ kubectl exec -ti $POD_NAME bash
root@kubernetes-bootcamp-b94cb9bff-tf9sg:/#
```
### Using a Service to Expose Your App
**Problem Statement**: Each Pod in a Kubernetes cluster has a unique IP address, even Pods on the same Node, so there needs to be a way of automatically reconciling changes among Pods so that your applications continue to function.(**Side thought**: [What would be some changes one need to do at pod level?](underConstruction.md) <sup>[1 -Pods](https://kubernetes.io/docs/concepts/workloads/pods/pod-overview/)</sup>)

value  |  Description
--|--
**Pending**  |  The Pod has been accepted by the Kubernetes system, but one or more of the Container images has not been created. This includes time before being scheduled as well as time spent downloading images over the network, which could take a while.
**Running**  |  The Pod has been bound to a node, and all of the Containers have been created. At least one Container is still running, or is in the process of starting or restarting.
**Succeeded**  |  All Containers in the Pod have terminated in success, and will not be restarted.
**Failed**  |  All Containers in the Pod have terminated, and at least one Container has terminated in failure. That is, the Container either exited with non-zero status or was terminated by the system
**Unknown**  |  For some reason the state of the Pod could not be obtained, typically due to an error in communicating with the host of the Pod.
Table| 1


Kubernetes Pods are mortal. Pods in fact have a lifecycle<sub>Table 1</sub> . When a worker node dies, the Pods running on the Node are also lost. A ReplicaSet might then dynamically drive the cluster back to desired state via creation of new Pods to keep your application running. A Service in Kubernetes is an abstraction which defines a logical set of Pods and a policy by which to access them. Services enable a loose coupling between dependent Pods. A Service is defined using YAML (preferred) or JSON, like all Kubernetes objects. The set of Pods targeted by a Service is usually determined by a LabelSelector (see below for why you might want a Service without including selector in the spec). A Service in Kubernetes is an abstraction which defines a logical set of Pods and a policy by which to access them. Services enable a loose coupling between dependent Pods. A Service is defined using YAML (preferred) or JSON, like all Kubernetes objects. The set of Pods targeted by a Service is usually determined by a LabelSelector . Although each Pod has a unique IP address, those IPs are not exposed outside the cluster without a Service. Services allow your applications to receive traffic. Services can be exposed in different ways by specifying a type in the ServiceSpec:
- ClusterIP (default) - Exposes the Service on an internal IP in the cluster. This type makes the Service only reachable from within the cluster.
- NodePort - Exposes the Service on the same port of each selected Node in the cluster using NAT. Makes a Service accessible from outside the cluster using <NodeIP>:<NodePort>. Superset of ClusterIP.
- LoadBalancer - Creates an external load balancer in the current cloud (if supported) and assigns a fixed, external IP to the Service. Superset of NodePort.
- ExternalName - Exposes the Service using an arbitrary name (specified by externalName in the spec) by returning a CNAME record with the name. No proxy is used. This type requires v1.7 or higher of kube-dns.

Additionally, note that there are some use cases with Services that involve not defining selector in the spec. A Service created without selector will also not create the corresponding Endpoints object. This allows users to manually map a Service to specific endpoints. Another possibility why there may be no selector is you are strictly using type: ExternalName.



<img src="https://d33wubrfki0l68.cloudfront.net/cc38b0f3c0fd94e66495e3a4198f2096cdecd3d5/ace10/docs/tutorials/kubernetes-basics/public/images/module_04_services.svg">

A Service routes traffic across a set of Pods. Services are the abstraction that allow pods to die and replicate in Kubernetes without impacting your application. Discovery and routing among dependent Pods (such as the frontend and backend components in an application) is handled by Kubernetes Services.

Services match a set of Pods using labels and selectors, a grouping primitive that allows logical operation on objects in Kubernetes. Labels are key/value pairs attached to objects and can be used in any number of ways:
- Designate objects for development, test, and production
- Embed version tags
- Classify an object using tags

<img src="https://d33wubrfki0l68.cloudfront.net/b964c59cdc1979dd4e1904c25f43745564ef6bee/f3351/docs/tutorials/kubernetes-basics/public/images/module_04_labels.svg">

Labels can be attached to objects at creation time or later on.you can create a Service at the same time you create a Deployment by using `--expose` in kubectl.

Example:

From previous module we have
```
olivier@obarbier:~$ kubectl get services
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   5m25s
olivier@obarbier:~$ kubectl get pods
NAME                                  READY   STATUS    RESTARTS   AGE
kubernetes-bootcamp-b94cb9bff-2z7l9   1/1     Running   0          86s
olivier@obarbier:~$ kubectl get deployments
NAME                  READY   UP-TO-DATE   AVAILABLE   AGE
kubernetes-bootcamp   1/1     1            1           2m14s

```

To create a service let's run the following command

``kubectl expose deployment/kubernetes-bootcamp --type="NodePort" --port 8080
``
1. Expose Service by specifying the deployment resources
2. identity type to expose
3. select port to access service.

We have now a running Service called kubernetes-bootcamp. Here we see that the Service received a unique cluster-IP, an internal port and an external-IP (the IP of the Node).
```
olivier@obarbier:~$ kubectl get services
NAME                  TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
kubernetes            ClusterIP   10.96.0.1       <none>        443/TCP          14m
kubernetes-bootcamp   NodePort    10.96.237.241   <none>        8080:30314/TCP   4m59s

```
```
olivier@obarbier:~$ kubectl describe services/kubernetes-bootcamp
Name:                     kubernetes-bootcamp
Namespace:                default
Labels:                   run=kubernetes-bootcamp
Annotations:              <none>
Selector:                 run=kubernetes-bootcamp
Type:                     NodePort
IP:                       10.96.237.241
Port:                     <unset>  8080/TCP
TargetPort:               8080/TCP
NodePort:                 <unset>  30314/TCP
Endpoints:                172.17.0.5:8080
Session Affinity:         None
External Traffic Policy:  Cluster
Events:                   <none>
```

How do we use labels Now

The Deployment created automatically a label for our Pod. i.e.
```
olivier@obarbier:~$ kubectl describe deployments
Name:                   kubernetes-bootcamp
Namespace:              default
CreationTimestamp:      Tue, 11 Jun 2019 19:41:36 -0400
Labels:                 run=kubernetes-bootcamp
Annotations:            deployment.kubernetes.io/revision: 1
Selector:               run=kubernetes-bootcamp
Replicas:               1 desired | 1 updated | 1 total | 1 available | 0 unavailable
StrategyType:           RollingUpdate
MinReadySeconds:        0
RollingUpdateStrategy:  25% max unavailable, 25% max surge
Pod Template:
  Labels:  run=kubernetes-bootcamp
  Containers:
   kubernetes-bootcamp:
    Image:        gcr.io/google-samples/kubernetes-bootcamp:v1
    Port:         8080/TCP
    Host Port:    0/TCP
    Environment:  <none>
    Mounts:       <none>
  Volumes:        <none>
Conditions:
  Type           Status  Reason
  ----           ------  ------
  Available      True    MinimumReplicasAvailable
  Progressing    True    NewReplicaSetAvailable
OldReplicaSets:  <none>
NewReplicaSet:   kubernetes-bootcamp-b94cb9bff (1/1 replicas created)
Events:
  Type    Reason             Age   From                   Message
  ----    ------             ----  ----                   -------
  Normal  ScalingReplicaSet  15m   deployment-controller  Scaled up replica set kubernetes-bootcamp-b94cb9bff to 1
```
et’s use this label to query our list of Pods. We’ll use the kubectl get pods command with -l as a parameter, followed by the label values. You can do the same to list the existing services
```
olivier@obarbier:~$ kubectl get pods -l run=kubernetes-bootcamp
NAME                                  READY   STATUS    RESTARTS   AGE
kubernetes-bootcamp-b94cb9bff-2z7l9   1/1     Running   0          17m
olivier@obarbier:~$ kubectl get services -l run=kubernetes-bootcamp
NAME                  TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
kubernetes-bootcamp   NodePort   10.96.237.241   <none>        8080:30314/TCP   13m
```

Now let's change labels
```
olivier@obarbier:~$ export POD_NAME=$(kubectl get pods -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')
olivier@obarbier:~$ echo Name of the Pod: $POD_NAME
Name of the Pod: kubernetes-bootcamp-b94cb9bff-2z7l9
olivier@obarbier:~$ kubectl label pod $POD_NAME app=v1
pod/kubernetes-bootcamp-b94cb9bff-2z7l9 labeled
olivier@obarbier:~$ kubectl describe pods $POD_NAME
Name:               kubernetes-bootcamp-b94cb9bff-2z7l9
Namespace:          default
Priority:           0
PriorityClassName:  <none>
Node:               minikube/10.0.2.15
Start Time:         Tue, 11 Jun 2019 19:41:36 -0400
Labels:             app=v1
                    pod-template-hash=b94cb9bff
                    run=kubernetes-bootcamp
Annotations:        <none>
Status:             Running
IP:                 172.17.0.5
Controlled By:      ReplicaSet/kubernetes-bootcamp-b94cb9bff
Containers:
  kubernetes-bootcamp:
    Container ID:   docker://b44ce34e044e75ab753824f50e13a3bf30d873be85cb34e1c3a7db14e5ea402b
    Image:          gcr.io/google-samples/kubernetes-bootcamp:v1
    Image ID:       docker-pullable://gcr.io/google-samples/kubernetes-bootcamp@sha256:0d6b8ee63bb57c5f5b6156f446b3bc3b3c143d233037f3a2f00e279c8fcc64af
    Port:           8080/TCP
    Host Port:      0/TCP
    State:          Running
      Started:      Tue, 11 Jun 2019 19:42:13 -0400
    Ready:          True
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-xcmdg (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             True
  ContainersReady   True
  PodScheduled      True
Volumes:
  default-token-xcmdg:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-xcmdg
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute for 300s
                 node.kubernetes.io/unreachable:NoExecute for 300s
Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  20m   default-scheduler  Successfully assigned default/kubernetes-bootcamp-b94cb9bff-2z7l9 to minikube
  Normal  Pulling    20m   kubelet, minikube  Pulling image "gcr.io/google-samples/kubernetes-bootcamp:v1"
  Normal  Pulled     20m   kubelet, minikube  Successfully pulled image "gcr.io/google-samples/kubernetes-bootcamp:v1"
  Normal  Created    20m   kubelet, minikube  Created container kubernetes-bootcamp
  Normal  Started    20m   kubelet, minikube  Started container kubernetes-bootcamp
olivier@obarbier:~$ kubectl get pods -l app=v1
NAME                                  READY   STATUS    RESTARTS   AGE
kubernetes-bootcamp-b94cb9bff-2z7l9   1/1     Running   0          21m
```
How to then delete a service

`kubectl delete service -l run=kubernetes-bootcamp`

```
olivier@obarbier:~$ kubectl delete service -l run=kubernetes-bootcamp
service "kubernetes-bootcamp" deleted
olivier@obarbier:~$ kubectl get services
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   27m
```
trying to connect using IP and port result in
```
olivier@obarbier:~$ curl $(minikube ip):$NODE_PORT
curl: (7) Failed to connect to 192.168.99.100 port 80: Connection refused
```

however this is still Running
```
olivier@obarbier:~$ kubectl exec -ti $POD_NAME curl localhost:8080
Hello Kubernetes bootcamp! | Running on: kubernetes-bootcamp-b94cb9bff-2z7l9 | v=1
```



## Reference
1. [Kubernetes Blog](https://kubernetes.io/blog/)
