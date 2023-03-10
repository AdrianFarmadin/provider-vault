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

type CustomMetadataObservation struct {
}

type CustomMetadataParameters struct {

	// If true, all keys will require the cas parameter to be set on all write requests.
	// +kubebuilder:validation:Optional
	CasRequired *bool `json:"casRequired,omitempty" tf:"cas_required,omitempty"`

	// A map of arbitrary string to string valued user-provided metadata meant to describe the secret.
	// +kubebuilder:validation:Optional
	Data map[string]*string `json:"data,omitempty" tf:"data,omitempty"`

	// If set, specifies the length of time before a version is deleted.
	// +kubebuilder:validation:Optional
	DeleteVersionAfter *float64 `json:"deleteVersionAfter,omitempty" tf:"delete_version_after,omitempty"`

	// The number of versions to keep per key.
	// +kubebuilder:validation:Optional
	MaxVersions *float64 `json:"maxVersions,omitempty" tf:"max_versions,omitempty"`
}

type SecretV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Metadata associated with this secret read from Vault.
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Full path where the KV-V2 secret will be written.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type SecretV2Parameters struct {

	// This flag is required if cas_required is set to true on either the secret or the engine's config. In order for a write to be successful, cas must be set to the current version of the secret.
	// +kubebuilder:validation:Optional
	Cas *float64 `json:"cas,omitempty" tf:"cas,omitempty"`

	// Custom metadata to be set for the secret.
	// +kubebuilder:validation:Optional
	CustomMetadata []CustomMetadataParameters `json:"customMetadata,omitempty" tf:"custom_metadata,omitempty"`

	// JSON-encoded secret data to write.
	// +kubebuilder:validation:Required
	DataJSONSecretRef v1.SecretKeySelector `json:"dataJsonSecretRef" tf:"-"`

	// If set to true, permanently deletes all versions for the specified key.
	// +kubebuilder:validation:Optional
	DeleteAllVersions *bool `json:"deleteAllVersions,omitempty" tf:"delete_all_versions,omitempty"`

	// If set to true, disables reading secret from Vault; note: drift won't be detected.
	// +kubebuilder:validation:Optional
	DisableRead *bool `json:"disableRead,omitempty" tf:"disable_read,omitempty"`

	// Path where KV-V2 engine is mounted.
	// +kubebuilder:validation:Required
	Mount *string `json:"mount" tf:"mount,omitempty"`

	// Full name of the secret. For a nested secret, the name is the nested path excluding the mount and data prefix. For example, for a secret at 'kvv2/data/foo/bar/baz', the name is 'foo/bar/baz'
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// An object that holds option settings.
	// +kubebuilder:validation:Optional
	Options map[string]*string `json:"options,omitempty" tf:"options,omitempty"`
}

// SecretV2Spec defines the desired state of SecretV2
type SecretV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretV2Parameters `json:"forProvider"`
}

// SecretV2Status defines the observed state of SecretV2.
type SecretV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretV2 is the Schema for the SecretV2s API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretV2Spec   `json:"spec"`
	Status            SecretV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretV2List contains a list of SecretV2s
type SecretV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretV2 `json:"items"`
}

// Repository type metadata.
var (
	SecretV2_Kind             = "SecretV2"
	SecretV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretV2_Kind}.String()
	SecretV2_KindAPIVersion   = SecretV2_Kind + "." + CRDGroupVersion.String()
	SecretV2_GroupVersionKind = CRDGroupVersion.WithKind(SecretV2_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretV2{}, &SecretV2List{})
}
