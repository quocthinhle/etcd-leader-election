# Leader Election with ETCD

## Why leader election?
- Now a days, microservices are very popular. In microservices, we have multiple instances of a service running.
Leader election is the idea of giving one process a massive power over the otherss.
This is really helpful, one process with the power includes assign work, schedule jobs, managing data, etc.

## How to implement leader election!
- We can implement leader election using etcd sw (set and watch) with the help of lease mechanism.
  - etcd provides a lease mechanism during leader election. When a node is elected leader,
    it writes a key in etcd to indicate its leadership and sets a time-to-live (TTL) for this key.
  - The leader must keep refreshing this key within the TTL. If the key expires (for example, if the leader fails),
    a new leader election is triggered, allowing another node to claim leadership.


## Getting started
- Init devbox project
```aiignore
devbox init
```

- Attach to devbox shell
```aiignore
devbox shell
```

- Setup the KinD cluster
```aiignore
cd ./infras/kind
# Create the KinD cluster with 4 nodes: 1 master and 3 workers
task kind-01:create-cluster
# For testing purpose, forward the ports to host with LoadBalancer service.
task kind-02:cloud-provider
```

- Start the etcd stateful set
```aiignore
cd ./infras/etcd
kubectl apply -f Namespace.yaml
kubectl apply -f Headless.yaml
kubectl apply -f StatefulSet.yaml
kubectl apply -f LoadBalancer.yaml
```

- Start the leader election application
```
cd ./infras/svc/leaderelection
kubectl apply -f Namespace.yaml
kubectl apply -f ConfigMap.yaml
kubectl apply -f Deployment.yaml
```

- Check the logs of the leader election application
```aiignore
# Switch to correct namespace
kubens svc-workload
kubectl get pods
kubectl logs -f <pod-name>
```

- Manually kill the leader pod and check the logs
```aiignore
# Kill the leader pod
kubectl delete pod <pod-name>
# Check the logs of the new leader
```