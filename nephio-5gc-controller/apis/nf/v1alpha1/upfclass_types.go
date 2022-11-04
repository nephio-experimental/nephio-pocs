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

// UpfClassSpec defines the desired state of UpfClass
type UpfClassSpec struct {
	PackageRef automationv1alpha1.PackageRevisionReference `json:"packageRef"`

	N3Endpoints int `json:"n3endpoints"`
	N4Endpoints int `json:"n4endpoints"`
	N6Endpoints int `json:"n6endpoints"`
	N9Endpoints int `json:"n9endpoints"`

	// +optional
	Dnn []string `json:"dnn"`
}

// UpfClassStatus defines the observed state of UpfClass
// TODO: we need a controller to validate that the packageRef is
// valid and that the underlying package is Ready
type UpfClassStatus struct {
	// Specifies whether the UpfClass is ready to be used
	Ready bool `json:"ready"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:path=upfclasses,scope=Cluster

// UpfClass is the Schema for the upfclasses API
type UpfClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UpfClassSpec   `json:"spec,omitempty"`
	Status UpfClassStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// UpfClassList contains a list of UpfClass
type UpfClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UpfClass `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UpfClass{}, &UpfClassList{})
}
