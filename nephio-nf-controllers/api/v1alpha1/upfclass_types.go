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

// UpfclassSpec defines the desired state of Upfclass
type UpfClassSpec struct {
	//Controller         string   `json:"controller"`
	UplinkThroughput   string `json:"uplinkThroughput"`
	DownlinkThroughput string `json:"downlinkThroughput"`
	N3Endpoints        int    `json:"n3endpoints"`
	N4Endpoints        int    `json:"n4endpoints"`
	N6Endpoints        int    `json:"n6endpoints"`
	N9Endpoints        int    `json:"n9endpoints"`
	// +optional
	Dnn []string `json:"dnn"`
}

// UpfclassStatus defines the observed state of Upfclass
type UpfClassStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:path=upfclasses,scope=Cluster

// Upfclass is the Schema for the upfclasses API
type UpfClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UpfClassSpec   `json:"spec,omitempty"`
	Status UpfClassStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// UpfclassList contains a list of Upfclass
type UpfClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UpfClass `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UpfClass{}, &UpfClassList{})
}
