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

// FiveGCoreTopologySpec defines the desired state of FiveGCoreTopology
type FiveGCoreTopologySpec struct {
	// UPFs lists different UPF configurations needed in this topology
	UPFs []UPFSpec `json:"upfs,omitempty"`
}

// FiveGCoreTopologyStatus defines the observed state of FiveGCoreTopology
type FiveGCoreTopologyStatus struct {
	// UPFStatuses lists the deployment status of each UPF configuration
	UPFStatuses []UPFStatus `json:"upfStatuses,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// FiveGCoreTopology is the Schema for the fivegcoretopologies API
type FiveGCoreTopology struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FiveGCoreTopologySpec   `json:"spec,omitempty"`
	Status FiveGCoreTopologyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// FiveGCoreTopologyList contains a list of FiveGCoreTopology
type FiveGCoreTopologyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FiveGCoreTopology `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FiveGCoreTopology{}, &FiveGCoreTopologyList{})
}
