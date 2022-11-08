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

type UPFN3 struct {
	Endpoints []Endpoint `json:"endpoints"`
}

type UPFN4 struct {
	Endpoints []Endpoint `json:"endpoints"`
}

type N6Endpoint struct {
	DNN        string   `json:"dnn"`
	IPEndpoint Endpoint `json:"ipEndpoint"`
	// UE address pool
	UEIPPool string `json:"ueIPPool"`
}

type UPFN6 struct {
	Endpoints []N6Endpoint `json:"endpoints"`
}

type UPFN9 struct {
	Endpoints []Endpoint `json:"endpoints"`
}

// UPFSpec defines the desired state of UPF
type UPFSpec struct {
	UPFClassName       string `json:"upfClassName"`
	UplinkThroughput   string `json:"uplinkThroughput"`
	DownlinkThroughput string `json:"downlinkThroughput"`
	N3                 UPFN3  `json:"n3"`
	N4                 UPFN4  `json:"n4"`
	N6                 UPFN6  `json:"n6"`
	// +optional
	N9 *UPFN9 `json:"n9,omitempty"`
}

// UPFStatus defines the observed state of UPF
type UPFStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// UPF is the Schema for the upfs API
type UPF struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UPFSpec   `json:"spec,omitempty"`
	Status UPFStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// UPFList contains a list of UPF
type UPFList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UPF `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UPF{}, &UPFList{})
}
