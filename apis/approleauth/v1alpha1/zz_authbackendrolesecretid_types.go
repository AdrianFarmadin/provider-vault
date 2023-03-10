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

type AuthBackendRoleSecretIDObservation struct {

	// The unique ID used to access this SecretID.
	Accessor *string `json:"accessor,omitempty" tf:"accessor,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The wrapped SecretID accessor.
	WrappingAccessor *string `json:"wrappingAccessor,omitempty" tf:"wrapping_accessor,omitempty"`
}

type AuthBackendRoleSecretIDParameters struct {

	// Unique name of the auth backend to configure.
	// +crossplane:generate:reference:type=github.com/AdrianFarmadin/provider-vault/apis/auth/v1alpha1.Backend
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Reference to a Backend in auth to populate backend.
	// +kubebuilder:validation:Optional
	BackendRef *v1.Reference `json:"backendRef,omitempty" tf:"-"`

	// Selector for a Backend in auth to populate backend.
	// +kubebuilder:validation:Optional
	BackendSelector *v1.Selector `json:"backendSelector,omitempty" tf:"-"`

	// List of CIDR blocks that can log in using the SecretID.
	// +kubebuilder:validation:Optional
	CidrList []*string `json:"cidrList,omitempty" tf:"cidr_list,omitempty"`

	// JSON-encoded secret data to write.
	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Name of the role.
	// +crossplane:generate:reference:type=AuthBackendRole
	// +crossplane:generate:reference:extractor=github.com/AdrianFarmadin/provider-vault/config/common.RoleNameExtractor()
	// +kubebuilder:validation:Optional
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// Reference to a AuthBackendRole to populate roleName.
	// +kubebuilder:validation:Optional
	RoleNameRef *v1.Reference `json:"roleNameRef,omitempty" tf:"-"`

	// Selector for a AuthBackendRole to populate roleName.
	// +kubebuilder:validation:Optional
	RoleNameSelector *v1.Selector `json:"roleNameSelector,omitempty" tf:"-"`

	// The SecretID to be managed. If not specified, Vault auto-generates one.
	// +kubebuilder:validation:Optional
	SecretIDSecretRef *v1.SecretKeySelector `json:"secretIdSecretRef,omitempty" tf:"-"`

	// Use the wrapped secret-id accessor as the id of this resource. If false, a fresh secret-id will be regenerated whenever the wrapping token is expired or invalidated through unwrapping.
	// +kubebuilder:validation:Optional
	WithWrappedAccessor *bool `json:"withWrappedAccessor,omitempty" tf:"with_wrapped_accessor,omitempty"`

	// The TTL duration of the wrapped SecretID.
	// +kubebuilder:validation:Optional
	WrappingTTL *string `json:"wrappingTtl,omitempty" tf:"wrapping_ttl,omitempty"`
}

// AuthBackendRoleSecretIDSpec defines the desired state of AuthBackendRoleSecretID
type AuthBackendRoleSecretIDSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthBackendRoleSecretIDParameters `json:"forProvider"`
}

// AuthBackendRoleSecretIDStatus defines the observed state of AuthBackendRoleSecretID.
type AuthBackendRoleSecretIDStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthBackendRoleSecretIDObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendRoleSecretID is the Schema for the AuthBackendRoleSecretIDs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type AuthBackendRoleSecretID struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AuthBackendRoleSecretIDSpec   `json:"spec"`
	Status            AuthBackendRoleSecretIDStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendRoleSecretIDList contains a list of AuthBackendRoleSecretIDs
type AuthBackendRoleSecretIDList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthBackendRoleSecretID `json:"items"`
}

// Repository type metadata.
var (
	AuthBackendRoleSecretID_Kind             = "AuthBackendRoleSecretID"
	AuthBackendRoleSecretID_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthBackendRoleSecretID_Kind}.String()
	AuthBackendRoleSecretID_KindAPIVersion   = AuthBackendRoleSecretID_Kind + "." + CRDGroupVersion.String()
	AuthBackendRoleSecretID_GroupVersionKind = CRDGroupVersion.WithKind(AuthBackendRoleSecretID_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthBackendRoleSecretID{}, &AuthBackendRoleSecretIDList{})
}
