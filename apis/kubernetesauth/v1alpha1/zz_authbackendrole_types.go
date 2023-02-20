/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AuthBackendRoleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AuthBackendRoleParameters struct {

	// Configures how identity aliases are generated. Valid choices are: serviceaccount_uid, serviceaccount_name
	// +kubebuilder:validation:Optional
	AliasNameSource *string `json:"aliasNameSource,omitempty" tf:"alias_name_source,omitempty"`

	// Optional Audience claim to verify in the JWT.
	// +kubebuilder:validation:Optional
	Audience *string `json:"audience,omitempty" tf:"audience,omitempty"`

	// Unique name of the kubernetes backend to configure.
	// +crossplane:generate:reference:type=github.com/AdrianFarmadin/provider-vault/apis/auth/v1alpha1.Backend
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Reference to a Backend in auth to populate backend.
	// +kubebuilder:validation:Optional
	BackendRef *v1.Reference `json:"backendRef,omitempty" tf:"-"`

	// Selector for a Backend in auth to populate backend.
	// +kubebuilder:validation:Optional
	BackendSelector *v1.Selector `json:"backendSelector,omitempty" tf:"-"`

	// List of service account names able to access this role. If set to `["*"]` all names are allowed, both this and bound_service_account_namespaces can not be "*".
	// +kubebuilder:validation:Required
	BoundServiceAccountNames []*string `json:"boundServiceAccountNames" tf:"bound_service_account_names,omitempty"`

	// List of namespaces allowed to access this role. If set to `["*"]` all namespaces are allowed, both this and bound_service_account_names can not be set to "*".
	// +kubebuilder:validation:Required
	BoundServiceAccountNamespaces []*string `json:"boundServiceAccountNamespaces" tf:"bound_service_account_namespaces,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Name of the role.
	// +kubebuilder:validation:Required
	RoleName *string `json:"roleName" tf:"role_name,omitempty"`

	// Specifies the blocks of IP addresses which are allowed to use the generated token
	// +kubebuilder:validation:Optional
	TokenBoundCidrs []*string `json:"tokenBoundCidrs,omitempty" tf:"token_bound_cidrs,omitempty"`

	// Generated Token's Explicit Maximum TTL in seconds
	// +kubebuilder:validation:Optional
	TokenExplicitMaxTTL *float64 `json:"tokenExplicitMaxTtl,omitempty" tf:"token_explicit_max_ttl,omitempty"`

	// The maximum lifetime of the generated token
	// +kubebuilder:validation:Optional
	TokenMaxTTL *float64 `json:"tokenMaxTtl,omitempty" tf:"token_max_ttl,omitempty"`

	// If true, the 'default' policy will not automatically be added to generated tokens
	// +kubebuilder:validation:Optional
	TokenNoDefaultPolicy *bool `json:"tokenNoDefaultPolicy,omitempty" tf:"token_no_default_policy,omitempty"`

	// The maximum number of times a token may be used, a value of zero means unlimited
	// +kubebuilder:validation:Optional
	TokenNumUses *float64 `json:"tokenNumUses,omitempty" tf:"token_num_uses,omitempty"`

	// Generated Token's Period
	// +kubebuilder:validation:Optional
	TokenPeriod *float64 `json:"tokenPeriod,omitempty" tf:"token_period,omitempty"`

	// Generated Token's Policies
	// +kubebuilder:validation:Optional
	TokenPolicies []*string `json:"tokenPolicies,omitempty" tf:"token_policies,omitempty"`

	// The initial ttl of the token to generate in seconds
	// +kubebuilder:validation:Optional
	TokenTTL *float64 `json:"tokenTtl,omitempty" tf:"token_ttl,omitempty"`

	// The type of token to generate, service or batch
	// +kubebuilder:validation:Optional
	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`
}

// AuthBackendRoleSpec defines the desired state of AuthBackendRole
type AuthBackendRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthBackendRoleParameters `json:"forProvider"`
}

// AuthBackendRoleStatus defines the observed state of AuthBackendRole.
type AuthBackendRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthBackendRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendRole is the Schema for the AuthBackendRoles API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type AuthBackendRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AuthBackendRoleSpec   `json:"spec"`
	Status            AuthBackendRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendRoleList contains a list of AuthBackendRoles
type AuthBackendRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthBackendRole `json:"items"`
}

// Repository type metadata.
var (
	AuthBackendRole_Kind             = "AuthBackendRole"
	AuthBackendRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthBackendRole_Kind}.String()
	AuthBackendRole_KindAPIVersion   = AuthBackendRole_Kind + "." + CRDGroupVersion.String()
	AuthBackendRole_GroupVersionKind = CRDGroupVersion.WithKind(AuthBackendRole_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthBackendRole{}, &AuthBackendRoleList{})
}
