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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SMFDeploymentSpec defines the desired state of SMFDeployment
type SMFDeploymentSpec struct {
	Capacity     SMFCapacity       `json:"capacity"`
	N4Interfaces []InterfaceConfig `json:"n4Interfaces"`
}

// SMFDeploymentStatus defines the observed state of SMFDeployment
type SMFDeploymentStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SMFDeployment is the Schema for the smfdeployments API
type SMFDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SMFDeploymentSpec   `json:"spec,omitempty"`
	Status SMFDeploymentStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SMFDeploymentList contains a list of SMFDeployment
type SMFDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SMFDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SMFDeployment{}, &SMFDeploymentList{})
}
