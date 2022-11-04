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

type UpfN3 struct {
	Endpoints []Endpoint `json:"endpoints"`
}

type UpfN4 struct {
	Endpoints []Endpoint `json:"endpoints"`
}

type N6Endpoint struct {
	Dnn         string   `json:"dnn"`
	IpEndpoints Endpoint `json:"ipendpoints"`
	// UE address pool
	IpAddrPool string `json:"ipaddrpool"`
}

type UpfN6 struct {
	Endpoints []N6Endpoint `json:"endpoints"`
}

type UpfN9 struct {
	Endpoints []Endpoint `json:"endpoints"`
}

// UpfSpec defines the desired state of Upf
type UpfSpec struct {
	UpfClassName       string `json:"upfClassName"`
	UplinkThroughput   string `json:"uplinkThroughput"`
	DownlinkThroughput string `json:"downlinkThroughput"`
	N3                 UpfN3  `json:"n3"`
	N4                 UpfN4  `json:"n4"`
	N6                 UpfN6  `json:"n6"`
	// +optional
	N9 *UpfN9 `json:"n9,omitempty"`
}

// UpfStatus defines the observed state of Upf
type UpfStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Upf is the Schema for the upfs API
type Upf struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UpfSpec   `json:"spec,omitempty"`
	Status UpfStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// UpfList contains a list of Upf
type UpfList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Upf `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Upf{}, &UpfList{})
}
