/*
Copyright 2024 The ORC Authors.

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
package common

import (
	applyconfigv1 "k8s.io/client-go/applyconfigurations/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/k-orc/openstack-resource-controller/internal/util/ssa"
)

type metaApplyConfig struct {
	applyconfigv1.TypeMetaApplyConfiguration   `json:",inline"`
	applyconfigv1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
}

func metaApplyConfigFromObject(obj client.Object) metaApplyConfig {
	gvk := obj.GetObjectKind().GroupVersionKind()

	applyConfig := metaApplyConfig{}

	// Type meta
	applyConfig.
		WithAPIVersion(gvk.GroupVersion().String()).
		WithKind(gvk.Kind)

	// Object meta
	applyConfig.
		WithName(obj.GetName()).
		WithNamespace(obj.GetNamespace()).
		WithUID(obj.GetUID()) // For safety: ensure we don't accidentally create a new object if we race with delete

	return applyConfig
}

// setFinalizer sets a finalizer on the object in its own SSA transaction.
func GetFinalizerPatch(obj client.Object, finalizer string) client.Patch {
	applyConfig := metaApplyConfigFromObject(obj)
	applyConfig.WithFinalizers(finalizer)
	return ssa.ApplyConfigPatch(applyConfig)
}

func RemoveFinalizerPatch(obj client.Object) client.Patch {
	applyConfig := metaApplyConfigFromObject(obj)
	return ssa.ApplyConfigPatch(applyConfig)
}
