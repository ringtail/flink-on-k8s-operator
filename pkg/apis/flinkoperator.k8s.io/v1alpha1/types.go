package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FlinkApplication represents a Flinkk application running on and using Kubernetes as a cluster manager.
type FlinkApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              FlinkApplicationSpec   `json:"spec"`
	Status            FlinkApplicationStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FlinkApplicationList carries a list of FlinkApplication objects.
type FlinkApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlinkApplication `json:"items,omitempty"`
}

// FlinkApplicationSpec describes the specification of a Spark application using Kubernetes as a cluster manager.
// It carries every pieces of information a spark-submit command takes and recognizes.
type FlinkApplicationSpec struct {

	// Image is the container image for the driver, executor, and init-container. Any custom container images for the
	// driver, executor, or init-container takes precedence over this.
	// Optional.
	Image *string `json:"image,omitempty"`

	// Jobmanager is the jobmanager specification
	Jobmanager JobmanagerSpec `json:"jobmanager"`

	// TaskManager is the executor specification.
	Taskmanager TaskmanagerSpec `json:"taskmanager"`
}

// JobmanagerSpec is specification of the jobmanager.
type JobmanagerSpec struct {
	// ServiceAccount is the name of the Kubernetes service account used by the driver pod
	// when requesting executor pods from the API server.
	ServiceAccount *string `json:"serviceAccount,omitempty"`
	// JavaOptions is a string of extra JVM options to pass to the driver. For instance,
	// GC settings or other logging.
	JavaOptions *string `json:"javaOptions,omitempty"`
}

// TaskmanagerSpec is specification of the taskmanager.
type TaskmanagerSpec struct {
	// Instances is the number of taskmanager instances.
	// Optional.
	Instances *int32 `json:"instances,omitempty"`
}

// FlinkApplicationStatus is the status for a FlinkApplication resource
type FlinkApplicationStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}
