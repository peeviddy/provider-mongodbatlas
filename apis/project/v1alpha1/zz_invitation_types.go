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

type InvitationObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InvitationID *string `json:"invitationId,omitempty" tf:"invitation_id,omitempty"`

	InviterUsername *string `json:"inviterUsername,omitempty" tf:"inviter_username,omitempty"`
}

type InvitationParameters struct {

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
	Roles []*string `json:"roles" tf:"roles,omitempty"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

// InvitationSpec defines the desired state of Invitation
type InvitationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InvitationParameters `json:"forProvider"`
}

// InvitationStatus defines the observed state of Invitation.
type InvitationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InvitationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Invitation is the Schema for the Invitations API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlas}
type Invitation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InvitationSpec   `json:"spec"`
	Status            InvitationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InvitationList contains a list of Invitations
type InvitationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Invitation `json:"items"`
}

// Repository type metadata.
var (
	Invitation_Kind             = "Invitation"
	Invitation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Invitation_Kind}.String()
	Invitation_KindAPIVersion   = Invitation_Kind + "." + CRDGroupVersion.String()
	Invitation_GroupVersionKind = CRDGroupVersion.WithKind(Invitation_Kind)
)

func init() {
	SchemeBuilder.Register(&Invitation{}, &InvitationList{})
}
