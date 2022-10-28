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

// NfDeploySite defines per site user intent for NF deployment
type NfDeloySite struct {
	Id           string `json:"id,omitempty" yaml:"id,omitempty"`
	LocationName string `json:"locationName,omitempty" yaml:"locationName,omitempty"`
	ClusterName  string `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`
	NfKind       string `json:"nfKind,omitempty" yaml:"nfKind,omitempty"`
	NfClassName  string `json:"nfClassName,omitempty" yaml:"nfClassName,omitempty"`
	NfVendor     string `json:"nfVendor,omitempty" yaml:"nfVendor,omitempty"`
	NfVersion    string `json:"nfVersion,omitempty" yaml:"nfVersion,omitempty"`
	NfNamespace  string `json:"nfNamespace,omitempty" yaml:"nfNamespace,omitempty"`
}

// NfDeploymentSpec defines the desired state of NfDeployment
type NfDeploymentSpec struct {
	Sites []NfDeloySite `json:"sites,omitempty" yaml:"sites,omitempty"`
}

// NfDeploymentStatus defines the observed state of NfDeployment
type NfDeploymentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// NfDeployment is the Schema for the nfdeployments API
type NfDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NfDeploymentSpec   `json:"spec,omitempty"`
	Status NfDeploymentStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NfDeploymentList contains a list of NfDeployment
type NfDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NfDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NfDeployment{}, &NfDeploymentList{})
}
