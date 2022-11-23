// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// DBSyncerDeploymentSpecApplyConfiguration represents an declarative configuration of the DBSyncerDeploymentSpec type for use
// with apply.
type DBSyncerDeploymentSpecApplyConfiguration struct {
	Resources *ResourcesApplyConfiguration `json:"resources,omitempty"`
}

// DBSyncerDeploymentSpecApplyConfiguration constructs an declarative configuration of the DBSyncerDeploymentSpec type for use with
// apply.
func DBSyncerDeploymentSpec() *DBSyncerDeploymentSpecApplyConfiguration {
	return &DBSyncerDeploymentSpecApplyConfiguration{}
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *DBSyncerDeploymentSpecApplyConfiguration) WithResources(value *ResourcesApplyConfiguration) *DBSyncerDeploymentSpecApplyConfiguration {
	b.Resources = value
	return b
}
