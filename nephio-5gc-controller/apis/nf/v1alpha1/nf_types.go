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

type Endpoint struct {

	// NetworkInstance identifies the layer 3 address space for IPAM
	// +optional
	NetworkInstance *string `json:"networkInstance,omitempty"`

	// NetworkName identifies the specific network to use for IPs
	// +optional
	NetworkName *string `json:"networkName,omitempty"`
}

type Pool struct {
	// NetworkInstance identifies the layer 3 address space for IPAM
	// +optional
	NetworkInstance *string `json:"networkIntance,omitempty"`

	// NetworkName identifies the specific network to use for IPs
	// +optional
	NetworkName *string `json:"networkName,omitempty"`

	// PrefixSize identifies the size of the pool needed
	// +optional
	PrefixSize *string `json:"prefixSize,omitempty"`
}
