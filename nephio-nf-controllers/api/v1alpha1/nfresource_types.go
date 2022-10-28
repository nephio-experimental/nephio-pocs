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

// UpfNadResourceSpec defines the user NAD structure for UPF
type UpfNadResourceSpec struct {
	N3Cni    string `json:"n3cni,omitempty"`
	N3Master string `json:"n3master,omitempty"`
	N3Gw     string `json:"n3gw,omitempty"`
	N4Cni    string `json:"n4cni,omitempty"`
	N4Master string `json:"n4master,omitempty"`
	N4Gw     string `json:"n4gw,omitempty"`
	N6Cni    string `json:"n6cni,omitempty"`
	N6Master string `json:"n6master,omitempty"`
	N6Gw     string `json:"n6gw,omitempty"`
}

// NfResourceSpec defines the desired state of NfResource
type NfResourceSpec struct {
	ClusterName string             `json:"clustername,omitempty"`
	Namespace   string             `json:"namespace,omitempty"`
	UpfNad      UpfNadResourceSpec `json:"upfnad,omitempty"`
	// TODO(user): more to come
}

// NfResourceStatus defines the observed state of NfResource
type NfResourceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// NfResource is the Schema for the nfresources API
type NfResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NfResourceSpec   `json:"spec,omitempty"`
	Status NfResourceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NfResourceList contains a list of NfResource
type NfResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NfResource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NfResource{}, &NfResourceList{})
}
