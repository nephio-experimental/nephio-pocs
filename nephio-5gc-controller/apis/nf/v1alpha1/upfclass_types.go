/*
Copyright 2022 The Nephio Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	automationv1alpha1 "github.com/nephio-project/nephio-controller-poc/apis/automation/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// UPFClassSpec defines the desired state of UPFClass
type UPFClassSpec struct {
	PackageRef automationv1alpha1.PackageRevisionReference `json:"packageRef"`

	N3EndpointCount int `json:"n3EndpointCount"`
	N4EndpointCount int `json:"n4EndpointCount"`
	N6EndpointCount int `json:"n6EndpointCount"`
	N9EndpointCount int `json:"n9EndpointCount"`

	// +optional
	DNNs []string `json:"dnns"`
}

// UPFClassStatus defines the observed state of UPFClass
// TODO: we need a controller to validate that the packageRef is
// valid and that the underlying package is Ready
type UPFClassStatus struct {
	// Specifies whether the UPFClass is ready to be used
	Ready bool `json:"ready"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:path=upfclasses,scope=Cluster

// UPFClass is the Schema for the upfclasses API
type UPFClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UPFClassSpec   `json:"spec,omitempty"`
	Status UPFClassStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// UPFClassList contains a list of UPFClass
type UPFClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UPFClass `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UPFClass{}, &UPFClassList{})
}
