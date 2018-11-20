# flink-on-k8s-operator
Kubernetes operator for specifying and managing the lifecycle of Apache Flink applications on Kubernetes.

## Project Status

**Project status:** *alpha*

* Define the FlinkApplication CRD template => Done
* Generate the FlinkApllication clientset  => Done
* Complete the FlinkApplication CRD...
* Define the ScheduledFlinkApplication CRD...
* Relies on k8s [garbage collecttion](https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/)...
* Implement flinkctl as a kubectl plugin...

## Installation

```bash
$ mkdir -p $GOPATH/src/github.com/iSofiane/flink-on-k8s-operator
$ cd $GOPATH/src/github.com/iSofiane/flink-on-k8s-operator
$ git clone https://github.com/iSofiane/flink-on-k8s-operator.git
$ go install
```

## Running a demo
```bash
$ flink-on-k8s-operator &
$ kubectl get crd
$ kubectl apply -f examples/test.yaml
$ ...
$ onAdd called
```

For a bet experience, you may want to use two terminals.
