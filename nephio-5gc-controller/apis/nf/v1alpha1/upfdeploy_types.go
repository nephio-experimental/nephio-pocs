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
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type InterfaceConfig struct {
	Name      string `json:"name"`
	IP        string `json:"ip"`
	GatewayIP string `json:"gatewayIP"`
}

type UPFCapacity struct {
	UplinkThroughput   resource.Quantity `json:"uplinkThroughput"`
	DownlinkThroughput resource.Quantity `json:"downlinkThroughput"`
}

type N6InterfaceConfig struct {
	// Dnn is the Data Network Name
	DNN       string          `json:"dnn"`
	Interface InterfaceConfig `json:"interface"`
	UEIPPool  string          `json:"ueIPPool"`
}

// UPFDeploymentSpec specifies config parameters for UPF
type UPFDeploymentSpec struct {
	Capacity     UPFCapacity         `json:"capacity,omitempty"`
	N3Interfaces []InterfaceConfig   `json:"n3Interfaces,omitempty"`
	N4Interfaces []InterfaceConfig   `json:"n4Interfaces,omitempty"`
	N6Interfaces []N6InterfaceConfig `json:"n6Interfaces,omitempty"`
	// +optional
	N9Interfaces []InterfaceConfig `json:"n9Interfaces,omitempty"`
}

type UPFDeploymentStatus struct {
	ComputeStatus   string      `json:"computestatus,omitempty"`
	ComputeUpTime   metav1.Time `json:"computeuptime,omitempty"`
	OperationStatus string      `json:"operationstatus,omitempty"`
	OperationUpTime metav1.Time `json:"operationuptime,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// UPFDeployment is the Schema for the upfdeployments API
type UPFDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UPFDeploymentSpec   `json:"spec,omitempty"`
	Status UPFDeploymentStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// UPFDeploymentList contains a list of UPFDeployment
type UPFDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UPFDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UPFDeployment{}, &UPFDeploymentList{})
}
