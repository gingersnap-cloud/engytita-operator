// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
)

// TLSConfigApplyConfiguration represents an declarative configuration of the TLSConfig type for use
// with apply.
type TLSConfigApplyConfiguration struct {
	SafeTLSConfigApplyConfiguration `json:",inline"`
	CAFile                          *string `json:"caFile,omitempty"`
	CertFile                        *string `json:"certFile,omitempty"`
	KeyFile                         *string `json:"keyFile,omitempty"`
}

// TLSConfigApplyConfiguration constructs an declarative configuration of the TLSConfig type for use with
// apply.
func TLSConfig() *TLSConfigApplyConfiguration {
	return &TLSConfigApplyConfiguration{}
}

// WithCA sets the CA field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CA field is set to the value of the last call.
func (b *TLSConfigApplyConfiguration) WithCA(value *SecretOrConfigMapApplyConfiguration) *TLSConfigApplyConfiguration {
	b.CA = value
	return b
}

// WithCert sets the Cert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Cert field is set to the value of the last call.
func (b *TLSConfigApplyConfiguration) WithCert(value *SecretOrConfigMapApplyConfiguration) *TLSConfigApplyConfiguration {
	b.Cert = value
	return b
}

// WithKeySecret sets the KeySecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KeySecret field is set to the value of the last call.
func (b *TLSConfigApplyConfiguration) WithKeySecret(value corev1.SecretKeySelector) *TLSConfigApplyConfiguration {
	b.KeySecret = &value
	return b
}

// WithServerName sets the ServerName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServerName field is set to the value of the last call.
func (b *TLSConfigApplyConfiguration) WithServerName(value string) *TLSConfigApplyConfiguration {
	b.ServerName = &value
	return b
}

// WithInsecureSkipVerify sets the InsecureSkipVerify field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InsecureSkipVerify field is set to the value of the last call.
func (b *TLSConfigApplyConfiguration) WithInsecureSkipVerify(value bool) *TLSConfigApplyConfiguration {
	b.InsecureSkipVerify = &value
	return b
}

// WithCAFile sets the CAFile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CAFile field is set to the value of the last call.
func (b *TLSConfigApplyConfiguration) WithCAFile(value string) *TLSConfigApplyConfiguration {
	b.CAFile = &value
	return b
}

// WithCertFile sets the CertFile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CertFile field is set to the value of the last call.
func (b *TLSConfigApplyConfiguration) WithCertFile(value string) *TLSConfigApplyConfiguration {
	b.CertFile = &value
	return b
}

// WithKeyFile sets the KeyFile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KeyFile field is set to the value of the last call.
func (b *TLSConfigApplyConfiguration) WithKeyFile(value string) *TLSConfigApplyConfiguration {
	b.KeyFile = &value
	return b
}
