//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Archive) DeepCopyInto(out *Archive) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Archive.
func (in *Archive) DeepCopy() *Archive {
	if in == nil {
		return nil
	}
	out := new(Archive)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Archive) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArchiveList) DeepCopyInto(out *ArchiveList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Archive, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArchiveList.
func (in *ArchiveList) DeepCopy() *ArchiveList {
	if in == nil {
		return nil
	}
	out := new(ArchiveList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArchiveList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArchiveObservation) DeepCopyInto(out *ArchiveObservation) {
	*out = *in
	if in.ArchiveID != nil {
		in, out := &in.ArchiveID, &out.ArchiveID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.PartitionFields != nil {
		in, out := &in.PartitionFields, &out.PartitionFields
		*out = make([]PartitionFieldsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArchiveObservation.
func (in *ArchiveObservation) DeepCopy() *ArchiveObservation {
	if in == nil {
		return nil
	}
	out := new(ArchiveObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArchiveParameters) DeepCopyInto(out *ArchiveParameters) {
	*out = *in
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.CollName != nil {
		in, out := &in.CollName, &out.CollName
		*out = new(string)
		**out = **in
	}
	if in.Criteria != nil {
		in, out := &in.Criteria, &out.Criteria
		*out = make([]CriteriaParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DBName != nil {
		in, out := &in.DBName, &out.DBName
		*out = new(string)
		**out = **in
	}
	if in.PartitionFields != nil {
		in, out := &in.PartitionFields, &out.PartitionFields
		*out = make([]PartitionFieldsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Paused != nil {
		in, out := &in.Paused, &out.Paused
		*out = new(bool)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.ProjectIDRef != nil {
		in, out := &in.ProjectIDRef, &out.ProjectIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ProjectIDSelector != nil {
		in, out := &in.ProjectIDSelector, &out.ProjectIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SyncCreation != nil {
		in, out := &in.SyncCreation, &out.SyncCreation
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArchiveParameters.
func (in *ArchiveParameters) DeepCopy() *ArchiveParameters {
	if in == nil {
		return nil
	}
	out := new(ArchiveParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArchiveSpec) DeepCopyInto(out *ArchiveSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArchiveSpec.
func (in *ArchiveSpec) DeepCopy() *ArchiveSpec {
	if in == nil {
		return nil
	}
	out := new(ArchiveSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArchiveStatus) DeepCopyInto(out *ArchiveStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArchiveStatus.
func (in *ArchiveStatus) DeepCopy() *ArchiveStatus {
	if in == nil {
		return nil
	}
	out := new(ArchiveStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CriteriaObservation) DeepCopyInto(out *CriteriaObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CriteriaObservation.
func (in *CriteriaObservation) DeepCopy() *CriteriaObservation {
	if in == nil {
		return nil
	}
	out := new(CriteriaObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CriteriaParameters) DeepCopyInto(out *CriteriaParameters) {
	*out = *in
	if in.DateField != nil {
		in, out := &in.DateField, &out.DateField
		*out = new(string)
		**out = **in
	}
	if in.DateFormat != nil {
		in, out := &in.DateFormat, &out.DateFormat
		*out = new(string)
		**out = **in
	}
	if in.ExpireAfterDays != nil {
		in, out := &in.ExpireAfterDays, &out.ExpireAfterDays
		*out = new(float64)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CriteriaParameters.
func (in *CriteriaParameters) DeepCopy() *CriteriaParameters {
	if in == nil {
		return nil
	}
	out := new(CriteriaParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionFieldsObservation) DeepCopyInto(out *PartitionFieldsObservation) {
	*out = *in
	if in.FieldType != nil {
		in, out := &in.FieldType, &out.FieldType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionFieldsObservation.
func (in *PartitionFieldsObservation) DeepCopy() *PartitionFieldsObservation {
	if in == nil {
		return nil
	}
	out := new(PartitionFieldsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionFieldsParameters) DeepCopyInto(out *PartitionFieldsParameters) {
	*out = *in
	if in.FieldName != nil {
		in, out := &in.FieldName, &out.FieldName
		*out = new(string)
		**out = **in
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionFieldsParameters.
func (in *PartitionFieldsParameters) DeepCopy() *PartitionFieldsParameters {
	if in == nil {
		return nil
	}
	out := new(PartitionFieldsParameters)
	in.DeepCopyInto(out)
	return out
}
