// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors All rights reserved.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func init() {
	if err := api.Scheme.AddGeneratedDeepCopyFuncs(
		DeepCopy_v1_CrossVersionObjectReference,
		DeepCopy_v1_HorizontalPodAutoscaler,
		DeepCopy_v1_HorizontalPodAutoscalerList,
		DeepCopy_v1_HorizontalPodAutoscalerSpec,
		DeepCopy_v1_HorizontalPodAutoscalerStatus,
		DeepCopy_v1_Scale,
		DeepCopy_v1_ScaleSpec,
		DeepCopy_v1_ScaleStatus,
	); err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}

func DeepCopy_v1_CrossVersionObjectReference(in CrossVersionObjectReference, out *CrossVersionObjectReference, c *conversion.Cloner) error {
	out.Kind = in.Kind
	out.Name = in.Name
	out.APIVersion = in.APIVersion
	return nil
}

func DeepCopy_v1_HorizontalPodAutoscaler(in HorizontalPodAutoscaler, out *HorizontalPodAutoscaler, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := api_v1.DeepCopy_v1_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_v1_HorizontalPodAutoscalerSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := DeepCopy_v1_HorizontalPodAutoscalerStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1_HorizontalPodAutoscalerList(in HorizontalPodAutoscalerList, out *HorizontalPodAutoscalerList, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]HorizontalPodAutoscaler, len(in))
		for i := range in {
			if err := DeepCopy_v1_HorizontalPodAutoscaler(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_v1_HorizontalPodAutoscalerSpec(in HorizontalPodAutoscalerSpec, out *HorizontalPodAutoscalerSpec, c *conversion.Cloner) error {
	if err := DeepCopy_v1_CrossVersionObjectReference(in.ScaleTargetRef, &out.ScaleTargetRef, c); err != nil {
		return err
	}
	if in.MinReplicas != nil {
		in, out := in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = *in
	} else {
		out.MinReplicas = nil
	}
	out.MaxReplicas = in.MaxReplicas
	if in.TargetCPUUtilizationPercentage != nil {
		in, out := in.TargetCPUUtilizationPercentage, &out.TargetCPUUtilizationPercentage
		*out = new(int32)
		**out = *in
	} else {
		out.TargetCPUUtilizationPercentage = nil
	}
	return nil
}

func DeepCopy_v1_HorizontalPodAutoscalerStatus(in HorizontalPodAutoscalerStatus, out *HorizontalPodAutoscalerStatus, c *conversion.Cloner) error {
	if in.ObservedGeneration != nil {
		in, out := in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = *in
	} else {
		out.ObservedGeneration = nil
	}
	if in.LastScaleTime != nil {
		in, out := in.LastScaleTime, &out.LastScaleTime
		*out = new(unversioned.Time)
		if err := unversioned.DeepCopy_unversioned_Time(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.LastScaleTime = nil
	}
	out.CurrentReplicas = in.CurrentReplicas
	out.DesiredReplicas = in.DesiredReplicas
	if in.CurrentCPUUtilizationPercentage != nil {
		in, out := in.CurrentCPUUtilizationPercentage, &out.CurrentCPUUtilizationPercentage
		*out = new(int32)
		**out = *in
	} else {
		out.CurrentCPUUtilizationPercentage = nil
	}
	return nil
}

func DeepCopy_v1_Scale(in Scale, out *Scale, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := api_v1.DeepCopy_v1_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_v1_ScaleSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := DeepCopy_v1_ScaleStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1_ScaleSpec(in ScaleSpec, out *ScaleSpec, c *conversion.Cloner) error {
	out.Replicas = in.Replicas
	return nil
}

func DeepCopy_v1_ScaleStatus(in ScaleStatus, out *ScaleStatus, c *conversion.Cloner) error {
	out.Replicas = in.Replicas
	out.Selector = in.Selector
	return nil
}
