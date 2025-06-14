/*
Copyright 2020 The Kubernetes Authors.

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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AegisOpsTemplate describe a ops template for alert
type AegisOpsTemplate struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// AegisOpsTemplateSpec defines the ops template content.
	// +optional
	Spec AegisOpsTemplateSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// AegisOpsTemplateStatus defines the template status.
	// +optional
	Status AegisOpsTemplateStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// AegisOpsTemplateSpec defines the ops template content.
type AegisOpsTemplateSpec struct {
	// +optional
	GenerateName string `json:"generateName,omitempty" protobuf:"bytes,1,rep,name=generateName"`

	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,2,rep,name=namespace"`

	// +optional
	Manifest string `json:"manifest,omitempty" protobuf:"bytes,3,opt,name=manifest"`
}

// AegisOpsTemplateStatus defines the alert/ops status.
type AegisOpsTemplateStatus struct {
	// Status is the template status.
	// +optional
	Status string `json:"status,omitempty" protobuf:"bytes,1,rep,name=status"`

	// +optional
	ExecuteStatus ExecuteStatus `json:"executeStatus,omitempty" protobuf:"bytes,2,rep,name=executeStatus"`
}

type ExecuteStatus struct {
	// +optional
	Succeeded int32 `json:"succeeded,omitempty" protobuf:"bytes,1,rep,name=succeeded"`

	// +optional
	Failed int32 `json:"failed,omitempty" protobuf:"bytes,2,rep,name=failed"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AegisOpsTemplateList is a list of AegisOpsTemplate items.
type AegisOpsTemplateList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard list metadata.
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items is a list of AegisOpsTemplate objects.
	Items []AegisOpsTemplate `json:"items" protobuf:"bytes,2,rep,name=items"`
}
