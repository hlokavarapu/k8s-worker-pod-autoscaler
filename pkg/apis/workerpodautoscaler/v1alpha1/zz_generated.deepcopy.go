// +build !ignore_autogenerated

/*
Copyright 2019 Practo Authors.

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
func (in *WorkerPodAutoScaler) DeepCopyInto(out *WorkerPodAutoScaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerPodAutoScaler.
func (in *WorkerPodAutoScaler) DeepCopy() *WorkerPodAutoScaler {
	if in == nil {
		return nil
	}
	out := new(WorkerPodAutoScaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerPodAutoScaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerPodAutoScalerList) DeepCopyInto(out *WorkerPodAutoScalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkerPodAutoScaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerPodAutoScalerList.
func (in *WorkerPodAutoScalerList) DeepCopy() *WorkerPodAutoScalerList {
	if in == nil {
		return nil
	}
	out := new(WorkerPodAutoScalerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerPodAutoScalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerPodAutoScalerSpec) DeepCopyInto(out *WorkerPodAutoScalerSpec) {
	*out = *in
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		*out = new(int32)
		**out = **in
	}
	if in.TargetMessagesPerWorker != nil {
		in, out := &in.TargetMessagesPerWorker, &out.TargetMessagesPerWorker
		*out = new(int32)
		**out = **in
	}
	if in.SecondsToProcessOneJob != nil {
		in, out := &in.SecondsToProcessOneJob, &out.SecondsToProcessOneJob
		*out = new(float32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerPodAutoScalerSpec.
func (in *WorkerPodAutoScalerSpec) DeepCopy() *WorkerPodAutoScalerSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerPodAutoScalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerPodAutoScalerStatus) DeepCopyInto(out *WorkerPodAutoScalerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerPodAutoScalerStatus.
func (in *WorkerPodAutoScalerStatus) DeepCopy() *WorkerPodAutoScalerStatus {
	if in == nil {
		return nil
	}
	out := new(WorkerPodAutoScalerStatus)
	in.DeepCopyInto(out)
	return out
}
