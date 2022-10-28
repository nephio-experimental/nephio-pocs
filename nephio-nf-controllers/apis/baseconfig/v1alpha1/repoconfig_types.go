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

// RepoConfigSpec defines the desired state of RepoConfig
type RepoConfigSpec struct {
	CatalogRepo     string            `json:"catalogrepo"`
	MgmtRepo        string            `json:"mgmtrepo"`
	DeployRepos     map[string]string `json:"deployrepos"`
	CatalogBasePkgs map[string]string `json:"catalogbasepkgs"`
	DeployBasePkgs  map[string]string `json:"deploybasepkgs"`
}

// RepoConfigStatus defines the observed state of RepoConfig
type RepoConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:path=repoconfigs,scope=Cluster

// RepoConfig is the Schema for the repoconfigs API
type RepoConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RepoConfigSpec   `json:"spec,omitempty"`
	Status RepoConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RepoConfigList contains a list of RepoConfig
type RepoConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RepoConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RepoConfig{}, &RepoConfigList{})
}
