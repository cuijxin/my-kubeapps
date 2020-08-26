// +build !ignore_autogenerated

/*
Copyright 2020 Bitnami.

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
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppRepository) DeepCopyInto(out *AppRepository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppRepository.
func (in *AppRepository) DeepCopy() *AppRepository {
	if in == nil {
		return nil
	}
	out := new(AppRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppRepository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppRepositoryAuth) DeepCopyInto(out *AppRepositoryAuth) {
	*out = *in
	if in.Header != nil {
		in, out := &in.Header, &out.Header
		*out = new(AppRepositoryAuthHeader)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomCA != nil {
		in, out := &in.CustomCA, &out.CustomCA
		*out = new(AppRepositoryCustomCA)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppRepositoryAuth.
func (in *AppRepositoryAuth) DeepCopy() *AppRepositoryAuth {
	if in == nil {
		return nil
	}
	out := new(AppRepositoryAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppRepositoryAuthHeader) DeepCopyInto(out *AppRepositoryAuthHeader) {
	*out = *in
	in.SecretKeyRef.DeepCopyInto(&out.SecretKeyRef)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppRepositoryAuthHeader.
func (in *AppRepositoryAuthHeader) DeepCopy() *AppRepositoryAuthHeader {
	if in == nil {
		return nil
	}
	out := new(AppRepositoryAuthHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppRepositoryCustomCA) DeepCopyInto(out *AppRepositoryCustomCA) {
	*out = *in
	in.SecretKeyRef.DeepCopyInto(&out.SecretKeyRef)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppRepositoryCustomCA.
func (in *AppRepositoryCustomCA) DeepCopy() *AppRepositoryCustomCA {
	if in == nil {
		return nil
	}
	out := new(AppRepositoryCustomCA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppRepositoryList) DeepCopyInto(out *AppRepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppRepository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppRepositoryList.
func (in *AppRepositoryList) DeepCopy() *AppRepositoryList {
	if in == nil {
		return nil
	}
	out := new(AppRepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppRepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppRepositorySpec) DeepCopyInto(out *AppRepositorySpec) {
	*out = *in
	in.Auth.DeepCopyInto(&out.Auth)
	in.SyncJobPodTemplate.DeepCopyInto(&out.SyncJobPodTemplate)
	if in.DockerRegistrySecrets != nil {
		in, out := &in.DockerRegistrySecrets, &out.DockerRegistrySecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppRepositorySpec.
func (in *AppRepositorySpec) DeepCopy() *AppRepositorySpec {
	if in == nil {
		return nil
	}
	out := new(AppRepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppRepositoryStatus) DeepCopyInto(out *AppRepositoryStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppRepositoryStatus.
func (in *AppRepositoryStatus) DeepCopy() *AppRepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(AppRepositoryStatus)
	in.DeepCopyInto(out)
	return out
}
