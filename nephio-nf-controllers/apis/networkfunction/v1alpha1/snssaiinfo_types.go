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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SnssaiBlk struct {
}

type DnsInfo struct {
    Ipv4    string  `json:"ipv4"`
}

type DnnInfo struct {
    Dnn     string  `json:"dnn"`
    Dns     DnsInfo `json:dns"`
}

type SnssaiInfo struct {
    Sst     int      `json:"sst"`
    Sd      int      `json:"sd"`
    Dnns    []DnnInfo  `json:"dnns"`
}

// SnssaiInfoSpec defines the desired state of SnssaiInfo
type SnssaiInfoSpec struct {
    // very preliminary
    Snssai  []SnssaiInfo  `json:"snssai"`
}

// SnssaiInfoStatus defines the observed state of SnssaiInfo
type SnssaiInfoStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SnssaiInfo is the Schema for the snssaiinfoes API
type SnssaiInfo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SnssaiInfoSpec   `json:"spec,omitempty"`
	Status SnssaiInfoStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SnssaiInfoList contains a list of SnssaiInfo
type SnssaiInfoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnssaiInfo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SnssaiInfo{}, &SnssaiInfoList{})
}
