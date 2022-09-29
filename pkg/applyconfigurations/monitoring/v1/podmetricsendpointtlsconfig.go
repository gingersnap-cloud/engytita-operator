// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
)

// PodMetricsEndpointTLSConfigApplyConfiguration represents an declarative configuration of the PodMetricsEndpointTLSConfig type for use
// with apply.
type PodMetricsEndpointTLSConfigApplyConfiguration struct {
	SafeTLSConfigApplyConfiguration `json:",inline"`
}

// PodMetricsEndpointTLSConfigApplyConfiguration constructs an declarative configuration of the PodMetricsEndpointTLSConfig type for use with
// apply.
func PodMetricsEndpointTLSConfig() *PodMetricsEndpointTLSConfigApplyConfiguration {
	return &PodMetricsEndpointTLSConfigApplyConfiguration{}
}

// WithCA sets the CA field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CA field is set to the value of the last call.
func (b *PodMetricsEndpointTLSConfigApplyConfiguration) WithCA(value *SecretOrConfigMapApplyConfiguration) *PodMetricsEndpointTLSConfigApplyConfiguration {
	b.CA = value
	return b
}

// WithCert sets the Cert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Cert field is set to the value of the last call.
func (b *PodMetricsEndpointTLSConfigApplyConfiguration) WithCert(value *SecretOrConfigMapApplyConfiguration) *PodMetricsEndpointTLSConfigApplyConfiguration {
	b.Cert = value
	return b
}

// WithKeySecret sets the KeySecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KeySecret field is set to the value of the last call.
func (b *PodMetricsEndpointTLSConfigApplyConfiguration) WithKeySecret(value corev1.SecretKeySelector) *PodMetricsEndpointTLSConfigApplyConfiguration {
	b.KeySecret = &value
	return b
}

// WithServerName sets the ServerName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServerName field is set to the value of the last call.
func (b *PodMetricsEndpointTLSConfigApplyConfiguration) WithServerName(value string) *PodMetricsEndpointTLSConfigApplyConfiguration {
	b.ServerName = &value
	return b
}

// WithInsecureSkipVerify sets the InsecureSkipVerify field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InsecureSkipVerify field is set to the value of the last call.
func (b *PodMetricsEndpointTLSConfigApplyConfiguration) WithInsecureSkipVerify(value bool) *PodMetricsEndpointTLSConfigApplyConfiguration {
	b.InsecureSkipVerify = &value
	return b
}
