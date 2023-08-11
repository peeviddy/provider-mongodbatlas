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

type APIKeyObservation struct {
	APIKeyID *string `json:"apiKeyId,omitempty" tf:"api_key_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`
}

type APIKeyParameters struct {

	// +kubebuilder:validation:Required
	Description *string `json:"description" tf:"description,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-mongodbatlas/apis/mongodbatlas/v1alpha1.Project
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-mongodbatlas/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in mongodbatlas to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	RoleNames []*string `json:"roleNames" tf:"role_names,omitempty"`
}

// APIKeySpec defines the desired state of APIKey
type APIKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIKeyParameters `json:"forProvider"`
}

// APIKeyStatus defines the observed state of APIKey.
type APIKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APIKey is the Schema for the APIKeys API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type APIKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APIKeySpec   `json:"spec"`
	Status            APIKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIKeyList contains a list of APIKeys
type APIKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIKey `json:"items"`
}

// Repository type metadata.
var (
	APIKey_Kind             = "APIKey"
	APIKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APIKey_Kind}.String()
	APIKey_KindAPIVersion   = APIKey_Kind + "." + CRDGroupVersion.String()
	APIKey_GroupVersionKind = CRDGroupVersion.WithKind(APIKey_Kind)
)

func init() {
	SchemeBuilder.Register(&APIKey{}, &APIKeyList{})
}
