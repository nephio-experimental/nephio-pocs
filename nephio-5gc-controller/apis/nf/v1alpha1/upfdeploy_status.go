package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type UpfDeployStatus struct {
	ComputeStatus   string      `json:"computestatus,omitempty"`
	ComputeUpTime   metav1.Time `json:"computeuptime,omitempty"`
	OperationStatus string      `json:"operationstatus,omitempty"`
	OperationUpTime metav1.Time `json:"operationuptime,omitempty"`
}
