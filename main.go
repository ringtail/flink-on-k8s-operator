package main

import (
	"flag"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/golang/glog"

	apiv1 "k8s.io/api/core/v1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	clientset "k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	crdclientset "github.com/iSofiane/flink-on-k8s-operator/pkg/client/clientset/versioned"
	crdinformers "github.com/iSofiane/flink-on-k8s-operator/pkg/client/informers/externalversions"
	"github.com/iSofiane/flink-on-k8s-operator/pkg/controller/flinkapplication"
	"github.com/iSofiane/flink-on-k8s-operator/pkg/crd"
	facrd "github.com/iSofiane/flink-on-k8s-operator/pkg/crd/flinkapplication"
)

var (
	installCRDs       = flag.Bool("install-crds", true, "Whether to install CRDs")
	controllerThreads = flag.Int("controller-threads", 3, "Number of worker threads used by the controller.")
	namespace         = flag.String("namespace", apiv1.NamespaceAll, "The Kubernetes namespace to manage..")
	resyncInterval    = flag.Int("resync-interval", 30, "Informer resync interval in seconds.")
)

func main() {

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional)")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	stopCh := make(chan struct{})

	kubeClient, err := clientset.NewForConfig(config)
	if err != nil {
		glog.Fatal(err)
	}

	crdClient, err := crdclientset.NewForConfig(config)
	if err != nil {
		glog.Fatal(err)
	}

	apiExtensionsClient, err := apiextensionsclient.NewForConfig(config)
	if err != nil {
		glog.Fatal(err)
	}

	// Start: FlinkApplicaion installation
	if *installCRDs {
		glog.Info("Installing the FlinkApplication CRD")
		err = crd.CreateOrUpdateCRD(apiExtensionsClient, facrd.GetCRD())
		if err != nil {
			glog.Fatalf("failed to create or update CustomResourceDefinition %s: %v", facrd.FullName, err)
		}
	}

	var factoryOpts []crdinformers.SharedInformerOption
	if *namespace != apiv1.NamespaceAll {
		factoryOpts = append(factoryOpts, crdinformers.WithNamespace(*namespace))
	}

	factory := crdinformers.NewSharedInformerFactoryWithOptions(
		crdClient,
		// resyncPeriod. Every resyncPeriod, all resources in the cache will re-trigger events.
		time.Duration(*resyncInterval)*time.Second,
		factoryOpts...)

	applicationController := flinkapplication.NewController(
		crdClient, kubeClient, apiExtensionsClient, factory, *namespace)

	// Start the informer factory that in turn starts the informer.
	go factory.Start(stopCh)
	if err = applicationController.Start(*controllerThreads, stopCh); err != nil {
		glog.Fatal(err)
	}

	//
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh

	close(stopCh)

	glog.Info("Shutting down the Spark operator")
	applicationController.Stop()

}
