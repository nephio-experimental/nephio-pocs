/*
Copyright 2022 Authors of Project Nephio.

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


type InterfaceConfig struct {
	Name   string   `json:"name"`
	IpAddr []string `json:"ipAddr"`
	GwAddr []string `json:"gwAddr"`
}

type UpfCapacity struct {
	UplinkThroughput   string `json:"uplinkThroughput"`
	DownlinkThroughput string `json:"downlinkThroughput"`
}

type N6InterfaceConfig struct {
	Dnn        string          `json:"dnn"`
	Interface  InterfaceConfig `json:"interface"`
	IpAddrPool string          `json:"ipAddrPool"`
}

// UpfDeploySpec specifies config parameters for UPF
type UpfDeploySpec struct {
	ImagePaths   map[string]string   `json:"imagePaths,omitempty"`
	Capacity     UpfCapacity         `json:"capacity,omitempty"`
	N3Interfaces []InterfaceConfig   `json:"n3Interfaces,omitempty"`
	N4Interfaces []InterfaceConfig   `json:"n4Interfaces,omitempty"`
	N6Interfaces []N6InterfaceConfig `json:"n6Interfaces,omitempty"`
	// +optional
	N9Interfaces []InterfaceConfig `json:"n9Interfaces,omitempty"`
}


type UpfDeployStatus struct {
	ComputeStatus   string      `json:"computestatus,omitempty"`
	ComputeUpTime   metav1.Time `json:"computeuptime,omitempty"`
	OperationStatus string      `json:"operationstatus,omitempty"`
	OperationUpTime metav1.Time `json:"operationuptime,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// UpfDeploy is the Schema for the upfdeploys API
type UpfDeploy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UpfDeploySpec   `json:"spec,omitempty"`
	Status UpfDeployStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// UpfDeployList contains a list of UpfDeploy
type UpfDeployList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UpfDeploy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UpfDeploy{}, &UpfDeployList{})
}
