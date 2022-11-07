/*
Copyright 2022.

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
    "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SmfSpec defines the desired state of Smf
type SmfSpec struct {
    SmfClassName    string  `json:"parent"`
    ClusterName     string  `json:"clustername"`
    Namespace       string  `json:"namespace"`
    N4              string  `json:"n4"`
    Sbi             SbiSpec `json:"sbi"`
    // +optional
    SnssaiInfo      v1.ObjectReference  `json:"snssaiinfo"`
}

// SmfStatus defines the observed state of Smf
type SmfStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Smf is the Schema for the smfs API
type Smf struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SmfSpec   `json:"spec,omitempty"`
	Status SmfStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SmfList contains a list of Smf
type SmfList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Smf `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Smf{}, &SmfList{})
}
