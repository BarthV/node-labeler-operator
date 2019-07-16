package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MergeSpec defines what need to be added to selected nodes
type MergeSpec struct {
	// NOTE : We cannot directly depends from metav1.ObjectMeta & v1.NodeSpec because
	// it would bring too many other unwanted fields and fill CRD with crap.
	// So we we need to mock~copy labels/annotations/taints/... defs in this project.

	// Annotations is an unstructured key value map stored with a resource that may be
	// set by external tools to store and retrieve arbitrary metadata. They are not
	// queryable and should be preserved when modifying objects.
	// More info: http://kubernetes.io/docs/user-guide/annotations
	// +optional
	Annotations map[string]string `json:"annotations,omitempty" protobuf:"bytes,1,rep,name=annotations"`

	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication controllers
	// and services.
	// More info: http://kubernetes.io/docs/user-guide/labels
	// +optional
	Labels map[string]string `json:"labels,omitempty" protobuf:"bytes,2,rep,name=labels"`

	// If specified, the node's taints.
	// +optional
	Taints []v1.Taint `json:"taints,omitempty" protobuf:"bytes,3,opt,name=taints"`
}

// LabelerSpec defines the desired state of Labeler
// +k8s:openapi-gen=true
type LabelerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// NodeSelector sets how the targets will be selected.
	v1.NodeSelector `json:",inline"`

	// MergeSpec sets everything that will be added to selected nodes
	Merge MergeSpec `json:"merge,omitempty"`

	// Size is how many nodes to label (not implemented yet).
	// Size int `json:"Size,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Labeler is the Schema for the labelers API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type Labeler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec LabelerSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LabelerList contains a list of Labeler
type LabelerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Labeler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Labeler{}, &LabelerList{})
}
