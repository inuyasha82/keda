// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleTriggers) DeepCopyInto(out *ScaleTriggers) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleTriggers.
func (in *ScaleTriggers) DeepCopy() *ScaleTriggers {
	if in == nil {
		return nil
	}
	out := new(ScaleTriggers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledObject) DeepCopyInto(out *ScaledObject) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledObject.
func (in *ScaledObject) DeepCopy() *ScaledObject {
	if in == nil {
		return nil
	}
	out := new(ScaledObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScaledObject) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledObjectList) DeepCopyInto(out *ScaledObjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScaledObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledObjectList.
func (in *ScaledObjectList) DeepCopy() *ScaledObjectList {
	if in == nil {
		return nil
	}
	out := new(ScaledObjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScaledObjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledObjectSpec) DeepCopyInto(out *ScaledObjectSpec) {
	*out = *in
	out.ScaleTargetRef = in.ScaleTargetRef
	if in.PollingInterval != nil {
		in, out := &in.PollingInterval, &out.PollingInterval
		*out = new(int32)
		**out = **in
	}
	if in.CooldownPeriod != nil {
		in, out := &in.CooldownPeriod, &out.CooldownPeriod
		*out = new(int32)
		**out = **in
	}
	if in.MinReplicaCount != nil {
		in, out := &in.MinReplicaCount, &out.MinReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicaCount != nil {
		in, out := &in.MaxReplicaCount, &out.MaxReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]ScaleTriggers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledObjectSpec.
func (in *ScaledObjectSpec) DeepCopy() *ScaledObjectSpec {
	if in == nil {
		return nil
	}
	out := new(ScaledObjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledObjectStatus) DeepCopyInto(out *ScaledObjectStatus) {
	*out = *in
	if in.LastActiveTime != nil {
		in, out := &in.LastActiveTime, &out.LastActiveTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledObjectStatus.
func (in *ScaledObjectStatus) DeepCopy() *ScaledObjectStatus {
	if in == nil {
		return nil
	}
	out := new(ScaledObjectStatus)
	in.DeepCopyInto(out)
	return out
}