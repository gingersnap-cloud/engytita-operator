// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceBindingWorkloadReferenceApplyConfiguration represents an declarative configuration of the ServiceBindingWorkloadReference type for use
// with apply.
type ServiceBindingWorkloadReferenceApplyConfiguration struct {
	APIVersion *string           `json:"apiVersion,omitempty"`
	Kind       *string           `json:"kind,omitempty"`
	Name       *string           `json:"name,omitempty"`
	Selector   *v1.LabelSelector `json:"selector,omitempty"`
	Containers []string          `json:"containers,omitempty"`
}

// ServiceBindingWorkloadReferenceApplyConfiguration constructs an declarative configuration of the ServiceBindingWorkloadReference type for use with
// apply.
func ServiceBindingWorkloadReference() *ServiceBindingWorkloadReferenceApplyConfiguration {
	return &ServiceBindingWorkloadReferenceApplyConfiguration{}
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *ServiceBindingWorkloadReferenceApplyConfiguration) WithAPIVersion(value string) *ServiceBindingWorkloadReferenceApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *ServiceBindingWorkloadReferenceApplyConfiguration) WithKind(value string) *ServiceBindingWorkloadReferenceApplyConfiguration {
	b.Kind = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ServiceBindingWorkloadReferenceApplyConfiguration) WithName(value string) *ServiceBindingWorkloadReferenceApplyConfiguration {
	b.Name = &value
	return b
}

// WithSelector sets the Selector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Selector field is set to the value of the last call.
func (b *ServiceBindingWorkloadReferenceApplyConfiguration) WithSelector(value v1.LabelSelector) *ServiceBindingWorkloadReferenceApplyConfiguration {
	b.Selector = &value
	return b
}

// WithContainers adds the given value to the Containers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Containers field.
func (b *ServiceBindingWorkloadReferenceApplyConfiguration) WithContainers(values ...string) *ServiceBindingWorkloadReferenceApplyConfiguration {
	for i := range values {
		b.Containers = append(b.Containers, values[i])
	}
	return b
}
