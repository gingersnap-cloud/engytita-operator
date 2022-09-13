// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfigurations

import (
	cacheregionv1alpha1 "github.com/engytita/engytita-operator/api/v1alpha1"
	v1alpha1 "github.com/engytita/engytita-operator/api/v1alpha1"
	cachev1alpha1 "github.com/engytita/engytita-operator/pkg/applyconfigurations/cache/v1alpha1"
	applyconfigurationscacheregionv1alpha1 "github.com/engytita/engytita-operator/pkg/applyconfigurations/cacheregion/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=cache, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("Cache"):
		return &cachev1alpha1.CacheApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("CacheSpec"):
		return &cachev1alpha1.CacheSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("CacheStatus"):
		return &cachev1alpha1.CacheStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ServiceBinding"):
		return &cachev1alpha1.ServiceBindingApplyConfiguration{}

		// Group=cacheregion, Version=v1alpha1
	case cacheregionv1alpha1.SchemeGroupVersion.WithKind("CacheRegion"):
		return &applyconfigurationscacheregionv1alpha1.CacheRegionApplyConfiguration{}
	case cacheregionv1alpha1.SchemeGroupVersion.WithKind("CacheRegionSpec"):
		return &applyconfigurationscacheregionv1alpha1.CacheRegionSpecApplyConfiguration{}
	case cacheregionv1alpha1.SchemeGroupVersion.WithKind("CacheService"):
		return &applyconfigurationscacheregionv1alpha1.CacheServiceApplyConfiguration{}

	}
	return nil
}
