/*
Copyright 2021 The KubeVela Authors.

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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ApplicationRevisionSpec is the spec of ApplicationRevision
type ApplicationRevisionSpec struct {
	// Application records the snapshot of the created/modified Application
	Application *Application `json:"application"`

	// ComponentDefinitions records the snapshot of the componentDefinitions related with the created/modified Application
	ComponentDefinitions []*ComponentDefinition `json:"componentDefinitions,omitempty"`

	// WorkloadDefinitions records the snapshot of the workloadDefinitions related with the created/modified Application
	WorkloadDefinitions []*WorkloadDefinition `json:"workloadDefinitions,omitempty"`

	// TraitDefinitions records the snapshot of the traitDefinitions related with the created/modified Application
	TraitDefinitions []*TraitDefinition `json:"traitDefinitions,omitempty"`

	// Components records the rendered components from Application, it will contains the whole K8s CR of workload in it.
	Components []*Component `json:"components"`

	// ApplicationConfiguration records the rendered applicationConfiguration from Application,
	// it will contains the whole K8s CR of trait and the reference component in it.
	ApplicationConfiguration *ApplicationConfiguration `json:"applicationConfiguration"`
}

// +kubebuilder:object:root=true

// ApplicationRevision is the Schema for the ApplicationRevision API
type ApplicationRevision struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ApplicationRevisionSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationRevisionList contains a list of ApplicationRevision
type ApplicationRevisionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationRevision `json:"items"`
}
