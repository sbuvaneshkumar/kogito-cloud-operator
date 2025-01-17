// Copyright 2019 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Builds) DeepCopyInto(out *Builds) {
	*out = *in
	if in.New != nil {
		in, out := &in.New, &out.New
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Pending != nil {
		in, out := &in.Pending, &out.Pending
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Running != nil {
		in, out := &in.Running, &out.Running
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Complete != nil {
		in, out := &in.Complete, &out.Complete
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Failed != nil {
		in, out := &in.Failed, &out.Failed
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Cancelled != nil {
		in, out := &in.Cancelled, &out.Cancelled
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Builds.
func (in *Builds) DeepCopy() *Builds {
	if in == nil {
		return nil
	}
	out := new(Builds)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataIndexCondition) DeepCopyInto(out *DataIndexCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataIndexCondition.
func (in *DataIndexCondition) DeepCopy() *DataIndexCondition {
	if in == nil {
		return nil
	}
	out := new(DataIndexCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Deployments) DeepCopyInto(out *Deployments) {
	*out = *in
	if in.Ready != nil {
		in, out := &in.Ready, &out.Ready
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Starting != nil {
		in, out := &in.Starting, &out.Starting
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Stopped != nil {
		in, out := &in.Stopped, &out.Stopped
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Failed != nil {
		in, out := &in.Failed, &out.Failed
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Deployments.
func (in *Deployments) DeepCopy() *Deployments {
	if in == nil {
		return nil
	}
	out := new(Deployments)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Env) DeepCopyInto(out *Env) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Env.
func (in *Env) DeepCopy() *Env {
	if in == nil {
		return nil
	}
	out := new(Env)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitSource) DeepCopyInto(out *GitSource) {
	*out = *in
	if in.URI != nil {
		in, out := &in.URI, &out.URI
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitSource.
func (in *GitSource) DeepCopy() *GitSource {
	if in == nil {
		return nil
	}
	out := new(GitSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanConnectionProperties) DeepCopyInto(out *InfinispanConnectionProperties) {
	*out = *in
	out.Credentials = in.Credentials
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanConnectionProperties.
func (in *InfinispanConnectionProperties) DeepCopy() *InfinispanConnectionProperties {
	if in == nil {
		return nil
	}
	out := new(InfinispanConnectionProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaConnectionProperties) DeepCopyInto(out *KafkaConnectionProperties) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaConnectionProperties.
func (in *KafkaConnectionProperties) DeepCopy() *KafkaConnectionProperties {
	if in == nil {
		return nil
	}
	out := new(KafkaConnectionProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoApp) DeepCopyInto(out *KogitoApp) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoApp.
func (in *KogitoApp) DeepCopy() *KogitoApp {
	if in == nil {
		return nil
	}
	out := new(KogitoApp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoApp) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoAppBuildObject) DeepCopyInto(out *KogitoAppBuildObject) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]Env, len(*in))
		copy(*out, *in)
	}
	if in.GitSource != nil {
		in, out := &in.GitSource, &out.GitSource
		*out = new(GitSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Webhooks != nil {
		in, out := &in.Webhooks, &out.Webhooks
		*out = make([]WebhookSecret, len(*in))
		copy(*out, *in)
	}
	out.ImageS2I = in.ImageS2I
	out.ImageRuntime = in.ImageRuntime
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoAppBuildObject.
func (in *KogitoAppBuildObject) DeepCopy() *KogitoAppBuildObject {
	if in == nil {
		return nil
	}
	out := new(KogitoAppBuildObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoAppList) DeepCopyInto(out *KogitoAppList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KogitoApp, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoAppList.
func (in *KogitoAppList) DeepCopy() *KogitoAppList {
	if in == nil {
		return nil
	}
	out := new(KogitoAppList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoAppList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoAppServiceObject) DeepCopyInto(out *KogitoAppServiceObject) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoAppServiceObject.
func (in *KogitoAppServiceObject) DeepCopy() *KogitoAppServiceObject {
	if in == nil {
		return nil
	}
	out := new(KogitoAppServiceObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoAppSpec) DeepCopyInto(out *KogitoAppSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]Env, len(*in))
		copy(*out, *in)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(KogitoAppBuildObject)
		(*in).DeepCopyInto(*out)
	}
	in.Service.DeepCopyInto(&out.Service)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoAppSpec.
func (in *KogitoAppSpec) DeepCopy() *KogitoAppSpec {
	if in == nil {
		return nil
	}
	out := new(KogitoAppSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoAppStatus) DeepCopyInto(out *KogitoAppStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Deployments.DeepCopyInto(&out.Deployments)
	in.Builds.DeepCopyInto(&out.Builds)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoAppStatus.
func (in *KogitoAppStatus) DeepCopy() *KogitoAppStatus {
	if in == nil {
		return nil
	}
	out := new(KogitoAppStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoDataIndex) DeepCopyInto(out *KogitoDataIndex) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoDataIndex.
func (in *KogitoDataIndex) DeepCopy() *KogitoDataIndex {
	if in == nil {
		return nil
	}
	out := new(KogitoDataIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoDataIndex) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoDataIndexList) DeepCopyInto(out *KogitoDataIndexList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KogitoDataIndex, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoDataIndexList.
func (in *KogitoDataIndexList) DeepCopy() *KogitoDataIndexList {
	if in == nil {
		return nil
	}
	out := new(KogitoDataIndexList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoDataIndexList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoDataIndexSpec) DeepCopyInto(out *KogitoDataIndexSpec) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Infinispan = in.Infinispan
	out.Kafka = in.Kafka
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoDataIndexSpec.
func (in *KogitoDataIndexSpec) DeepCopy() *KogitoDataIndexSpec {
	if in == nil {
		return nil
	}
	out := new(KogitoDataIndexSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoDataIndexStatus) DeepCopyInto(out *KogitoDataIndexStatus) {
	*out = *in
	in.DeploymentStatus.DeepCopyInto(&out.DeploymentStatus)
	in.ServiceStatus.DeepCopyInto(&out.ServiceStatus)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]DataIndexCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DependenciesStatus != nil {
		in, out := &in.DependenciesStatus, &out.DependenciesStatus
		*out = make([]DataIndexDependenciesStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoDataIndexStatus.
func (in *KogitoDataIndexStatus) DeepCopy() *KogitoDataIndexStatus {
	if in == nil {
		return nil
	}
	out := new(KogitoDataIndexStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMap) DeepCopyInto(out *ResourceMap) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMap.
func (in *ResourceMap) DeepCopy() *ResourceMap {
	if in == nil {
		return nil
	}
	out := new(ResourceMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make([]ResourceMap, len(*in))
		copy(*out, *in)
	}
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = make([]ResourceMap, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretCredentialsType) DeepCopyInto(out *SecretCredentialsType) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretCredentialsType.
func (in *SecretCredentialsType) DeepCopy() *SecretCredentialsType {
	if in == nil {
		return nil
	}
	out := new(SecretCredentialsType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookSecret) DeepCopyInto(out *WebhookSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookSecret.
func (in *WebhookSecret) DeepCopy() *WebhookSecret {
	if in == nil {
		return nil
	}
	out := new(WebhookSecret)
	in.DeepCopyInto(out)
	return out
}
