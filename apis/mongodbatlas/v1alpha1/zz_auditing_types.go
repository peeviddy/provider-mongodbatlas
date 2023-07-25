/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AuditingObservation struct {
	ConfigurationType *string `json:"configurationType,omitempty" tf:"configuration_type,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AuditingParameters struct {

	// +kubebuilder:validation:Optional
	AuditAuthorizationSuccess *bool `json:"auditAuthorizationSuccess,omitempty" tf:"audit_authorization_success,omitempty"`

	// +kubebuilder:validation:Optional
	AuditFilter *string `json:"auditFilter,omitempty" tf:"audit_filter,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +crossplane:generate:reference:type=Project
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-mongodbatlas/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`
}

// AuditingSpec defines the desired state of Auditing
type AuditingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuditingParameters `json:"forProvider"`
}

// AuditingStatus defines the observed state of Auditing.
type AuditingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuditingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Auditing is the Schema for the Auditings API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type Auditing struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AuditingSpec   `json:"spec"`
	Status            AuditingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuditingList contains a list of Auditings
type AuditingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Auditing `json:"items"`
}

// Repository type metadata.
var (
	Auditing_Kind             = "Auditing"
	Auditing_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Auditing_Kind}.String()
	Auditing_KindAPIVersion   = Auditing_Kind + "." + CRDGroupVersion.String()
	Auditing_GroupVersionKind = CRDGroupVersion.WithKind(Auditing_Kind)
)

func init() {
	SchemeBuilder.Register(&Auditing{}, &AuditingList{})
}
