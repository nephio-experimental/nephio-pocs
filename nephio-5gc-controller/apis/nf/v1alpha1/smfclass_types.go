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

type SMFCapacity struct {
	MaxPduSessions int `json:"maxpdusessions"`
	// +optional
	MaxConnectedUpfs int `json:"maxconnectedupfs"`
}

// SMFClassSpec defines the desired state of SMFClass
type SMFClassSpec struct {
	PackageRef automationv1alpha1.PackageRevisionReference `json:"packageRef"`

	// +optional
	N4EndpointCount int `json:"n4EndpointCount"`
}

// SMFClassStatus defines the observed state of SMFClass
type SMFClassStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SMFClass is the Schema for the smfclasses API
type SMFClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SMFClassSpec   `json:"spec,omitempty"`
	Status SMFClassStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SMFClassList contains a list of SMFClass
type SMFClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SMFClass `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SMFClass{}, &SMFClassList{})
}
