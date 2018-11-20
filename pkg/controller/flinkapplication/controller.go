package flinkapplication

import (
	"fmt"
	//"reflect"
	//"strconv"
	"time"

	"github.com/golang/glog"

	apiv1 "k8s.io/api/core/v1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientset "k8s.io/client-go/kubernetes"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"

	"github.com/iSofiane/flink-on-k8s-operator/pkg/apis/flinkoperator.k8s.io/v1alpha1"
	crdclientset "github.com/iSofiane/flink-on-k8s-operator/pkg/client/clientset/versioned"
	crdscheme "github.com/iSofiane/flink-on-k8s-operator/pkg/client/clientset/versioned/scheme"
	crdinformers "github.com/iSofiane/flink-on-k8s-operator/pkg/client/informers/externalversions"
	crdlisters "github.com/iSofiane/flink-on-k8s-operator/pkg/client/listers/flinkoperator.k8s.io/v1alpha1"
)

var (
	keyFunc = cache.DeletionHandlingMetaNamespaceKeyFunc
)

// Controller is the controller of the FlinkApplication
type Controller struct {
	crdClient        crdclientset.Interface
	kubeClient       clientset.Interface
	extensionsClient apiextensionsclient.Interface

	queue       workqueue.RateLimitingInterface
	cacheSynced cache.InformerSynced

	lister   crdlisters.FlinkApplicationLister
	recorder record.EventRecorder
}

// NewController creates a new Controller.
func NewController(
	crdClient crdclientset.Interface,
	kubeClient clientset.Interface,
	extensionsClient apiextensionsclient.Interface,
	informerFactory crdinformers.SharedInformerFactory,
	namespace string) *Controller {
	crdscheme.AddToScheme(scheme.Scheme)

	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.V(2).Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{
		Interface: kubeClient.CoreV1().Events(namespace),
	})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, apiv1.EventSource{Component: "flink-operator"})

	return newFlinkApplicationController(
		crdClient,
		kubeClient,
		extensionsClient,
		informerFactory,
		recorder,
		namespace)
}

func newFlinkApplicationController(
	crdClient crdclientset.Interface,
	kubeClient clientset.Interface,
	extensionsClient apiextensionsclient.Interface,
	informerFactory crdinformers.SharedInformerFactory,
	eventRecorder record.EventRecorder,
	namespace string) *Controller {
	queue := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(),
		"flink-application-controller")

	controller := &Controller{
		crdClient:        crdClient,
		kubeClient:       kubeClient,
		extensionsClient: extensionsClient,
		recorder:         eventRecorder,
		queue:            queue,
	}

	informer := informerFactory.Flinkoperator().V1alpha1().FlinkApplications()
	informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.onAdd,
		//UpdateFunc: controller.onUpdate,
		DeleteFunc: controller.onDelete,
	})
	controller.cacheSynced = informer.Informer().HasSynced
	controller.lister = informer.Lister()

	return controller
}

// Start starts the Controller by registering a watcher for SparkApplication objects.
func (c *Controller) Start(workers int, stopCh <-chan struct{}) error {
	glog.Info("Starting the SparkApplication controller")

	if !cache.WaitForCacheSync(stopCh, c.cacheSynced) {
		return fmt.Errorf("timed out waiting for cache to sync")
	}

	glog.Info("Starting the workers of the SparkApplication controller")
	for i := 0; i < workers; i++ {
		// runWorker will loop until "something bad" happens. Until will then rekick
		// the worker after one second.
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	return nil
}

// Stop stops the controller.
func (c *Controller) Stop() {
	glog.Info("Stopping the SparkApplication controller")
	c.queue.ShutDown()
}

// runWorker runs a single controller worker.
func (c *Controller) runWorker() {
	defer utilruntime.HandleCrash()
	for c.processNextItem() {
	}
}

func (c *Controller) processNextItem() bool {
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)

	err := c.syncFlinkApplication(key.(string))
	if err == nil {
		// Successfully processed the key or the key was not found so tell the queue to stop tracking
		// history for your key. This will reset things like failure counts for per-item rate limiting.
		c.queue.Forget(key)
		return true
	}

	// There was a failure so be sure to report it. This method allows for pluggable error handling
	// which can be used for things like cluster-monitoring
	utilruntime.HandleError(fmt.Errorf("failed to sync SparkApplication %q: %v", key, err))
	// Since we failed, we should requeue the item to work on later.  This method will add a backoff
	// to avoid hot-looping on particular items (they're probably still not going to work right away)
	// and overall controller protection (everything I've done is broken, this controller needs to
	// calm down or it can starve other useful work) cases.
	c.queue.AddRateLimited(key)

	return true
}

func (c *Controller) syncFlinkApplication(key string) error {
	fmt.Println("syncFlinkApplication called")
	return nil
}

// Callback function called when a new SparkApplication object gets created.
func (c *Controller) onAdd(obj interface{}) {
	fmt.Println("onAdd called")
	app := obj.(*v1alpha1.FlinkApplication)
	c.enqueue(app)
}

func (c *Controller) onUpdate(obj interface{}) {
	fmt.Println("onupdate called")
}

// Callback function called when a new SparkApplication object gets created.
func (c *Controller) onDelete(obj interface{}) {
	fmt.Println("onDelete called")
	c.dequeue(obj)
}

func (c *Controller) enqueue(obj interface{}) {
	key, err := keyFunc(obj)
	if err != nil {
		glog.Errorf("failed to get key for %v: %v", obj, err)
		return
	}

	c.queue.AddRateLimited(key)
}

func (c *Controller) dequeue(obj interface{}) {
	key, err := keyFunc(obj)
	if err != nil {
		glog.Errorf("failed to get key for %v: %v", obj, err)
		return
	}

	c.queue.Forget(key)
	c.queue.Done(key)
}
