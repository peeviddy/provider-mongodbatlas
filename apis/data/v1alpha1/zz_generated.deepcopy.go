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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsObservation) DeepCopyInto(out *AwsObservation) {
	*out = *in
	if in.ExternalID != nil {
		in, out := &in.ExternalID, &out.ExternalID
		*out = new(string)
		**out = **in
	}
	if in.IAMAssumedRoleArn != nil {
		in, out := &in.IAMAssumedRoleArn, &out.IAMAssumedRoleArn
		*out = new(string)
		**out = **in
	}
	if in.IAMUserArn != nil {
		in, out := &in.IAMUserArn, &out.IAMUserArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsObservation.
func (in *AwsObservation) DeepCopy() *AwsObservation {
	if in == nil {
		return nil
	}
	out := new(AwsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsParameters) DeepCopyInto(out *AwsParameters) {
	*out = *in
	if in.RoleID != nil {
		in, out := &in.RoleID, &out.RoleID
		*out = new(string)
		**out = **in
	}
	if in.TestS3Bucket != nil {
		in, out := &in.TestS3Bucket, &out.TestS3Bucket
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsParameters.
func (in *AwsParameters) DeepCopy() *AwsParameters {
	if in == nil {
		return nil
	}
	out := new(AwsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectionsObservation) DeepCopyInto(out *CollectionsObservation) {
	*out = *in
	if in.DataSources != nil {
		in, out := &in.DataSources, &out.DataSources
		*out = make([]DataSourcesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectionsObservation.
func (in *CollectionsObservation) DeepCopy() *CollectionsObservation {
	if in == nil {
		return nil
	}
	out := new(CollectionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectionsParameters) DeepCopyInto(out *CollectionsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectionsParameters.
func (in *CollectionsParameters) DeepCopy() *CollectionsParameters {
	if in == nil {
		return nil
	}
	out := new(CollectionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataProcessRegionObservation) DeepCopyInto(out *DataProcessRegionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataProcessRegionObservation.
func (in *DataProcessRegionObservation) DeepCopy() *DataProcessRegionObservation {
	if in == nil {
		return nil
	}
	out := new(DataProcessRegionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataProcessRegionParameters) DeepCopyInto(out *DataProcessRegionParameters) {
	*out = *in
	if in.CloudProvider != nil {
		in, out := &in.CloudProvider, &out.CloudProvider
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataProcessRegionParameters.
func (in *DataProcessRegionParameters) DeepCopy() *DataProcessRegionParameters {
	if in == nil {
		return nil
	}
	out := new(DataProcessRegionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataSourcesObservation) DeepCopyInto(out *DataSourcesObservation) {
	*out = *in
	if in.DefaultFormat != nil {
		in, out := &in.DefaultFormat, &out.DefaultFormat
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.StoreName != nil {
		in, out := &in.StoreName, &out.StoreName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataSourcesObservation.
func (in *DataSourcesObservation) DeepCopy() *DataSourcesObservation {
	if in == nil {
		return nil
	}
	out := new(DataSourcesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataSourcesParameters) DeepCopyInto(out *DataSourcesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataSourcesParameters.
func (in *DataSourcesParameters) DeepCopy() *DataSourcesParameters {
	if in == nil {
		return nil
	}
	out := new(DataSourcesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Lake) DeepCopyInto(out *Lake) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Lake.
func (in *Lake) DeepCopy() *Lake {
	if in == nil {
		return nil
	}
	out := new(Lake)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Lake) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LakeList) DeepCopyInto(out *LakeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Lake, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LakeList.
func (in *LakeList) DeepCopy() *LakeList {
	if in == nil {
		return nil
	}
	out := new(LakeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LakeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LakeObservation) DeepCopyInto(out *LakeObservation) {
	*out = *in
	if in.Hostnames != nil {
		in, out := &in.Hostnames, &out.Hostnames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.StorageDatabases != nil {
		in, out := &in.StorageDatabases, &out.StorageDatabases
		*out = make([]StorageDatabasesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StorageStores != nil {
		in, out := &in.StorageStores, &out.StorageStores
		*out = make([]StorageStoresObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LakeObservation.
func (in *LakeObservation) DeepCopy() *LakeObservation {
	if in == nil {
		return nil
	}
	out := new(LakeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LakeParameters) DeepCopyInto(out *LakeParameters) {
	*out = *in
	if in.Aws != nil {
		in, out := &in.Aws, &out.Aws
		*out = make([]AwsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DataProcessRegion != nil {
		in, out := &in.DataProcessRegion, &out.DataProcessRegion
		*out = make([]DataProcessRegionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LakeParameters.
func (in *LakeParameters) DeepCopy() *LakeParameters {
	if in == nil {
		return nil
	}
	out := new(LakeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LakeSpec) DeepCopyInto(out *LakeSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LakeSpec.
func (in *LakeSpec) DeepCopy() *LakeSpec {
	if in == nil {
		return nil
	}
	out := new(LakeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LakeStatus) DeepCopyInto(out *LakeStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LakeStatus.
func (in *LakeStatus) DeepCopy() *LakeStatus {
	if in == nil {
		return nil
	}
	out := new(LakeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageDatabasesObservation) DeepCopyInto(out *StorageDatabasesObservation) {
	*out = *in
	if in.Collections != nil {
		in, out := &in.Collections, &out.Collections
		*out = make([]CollectionsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MaxWildcardCollections != nil {
		in, out := &in.MaxWildcardCollections, &out.MaxWildcardCollections
		*out = new(int64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Views != nil {
		in, out := &in.Views, &out.Views
		*out = make([]ViewsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageDatabasesObservation.
func (in *StorageDatabasesObservation) DeepCopy() *StorageDatabasesObservation {
	if in == nil {
		return nil
	}
	out := new(StorageDatabasesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageDatabasesParameters) DeepCopyInto(out *StorageDatabasesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageDatabasesParameters.
func (in *StorageDatabasesParameters) DeepCopy() *StorageDatabasesParameters {
	if in == nil {
		return nil
	}
	out := new(StorageDatabasesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageStoresObservation) DeepCopyInto(out *StorageStoresObservation) {
	*out = *in
	if in.AdditionalStorageClasses != nil {
		in, out := &in.AdditionalStorageClasses, &out.AdditionalStorageClasses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.Delimiter != nil {
		in, out := &in.Delimiter, &out.Delimiter
		*out = new(string)
		**out = **in
	}
	if in.IncludeTags != nil {
		in, out := &in.IncludeTags, &out.IncludeTags
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageStoresObservation.
func (in *StorageStoresObservation) DeepCopy() *StorageStoresObservation {
	if in == nil {
		return nil
	}
	out := new(StorageStoresObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageStoresParameters) DeepCopyInto(out *StorageStoresParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageStoresParameters.
func (in *StorageStoresParameters) DeepCopy() *StorageStoresParameters {
	if in == nil {
		return nil
	}
	out := new(StorageStoresParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ViewsObservation) DeepCopyInto(out *ViewsObservation) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Pipeline != nil {
		in, out := &in.Pipeline, &out.Pipeline
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ViewsObservation.
func (in *ViewsObservation) DeepCopy() *ViewsObservation {
	if in == nil {
		return nil
	}
	out := new(ViewsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ViewsParameters) DeepCopyInto(out *ViewsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ViewsParameters.
func (in *ViewsParameters) DeepCopy() *ViewsParameters {
	if in == nil {
		return nil
	}
	out := new(ViewsParameters)
	in.DeepCopyInto(out)
	return out
}
