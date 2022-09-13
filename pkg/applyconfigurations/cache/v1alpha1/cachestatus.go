// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// CacheStatusApplyConfiguration represents an declarative configuration of the CacheStatus type for use
// with apply.
type CacheStatusApplyConfiguration struct {
	ServiceBinding *ServiceBindingApplyConfiguration `json:"binding,omitempty"`
}

// CacheStatusApplyConfiguration constructs an declarative configuration of the CacheStatus type for use with
// apply.
func CacheStatus() *CacheStatusApplyConfiguration {
	return &CacheStatusApplyConfiguration{}
}

// WithServiceBinding sets the ServiceBinding field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceBinding field is set to the value of the last call.
func (b *CacheStatusApplyConfiguration) WithServiceBinding(value *ServiceBindingApplyConfiguration) *CacheStatusApplyConfiguration {
	b.ServiceBinding = value
	return b
}