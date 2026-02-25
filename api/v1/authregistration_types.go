/*
This file was generated with "make generate-deepcopy".
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	// ConditionCompleted indicates that the registration process has finished.
	ConditionCompleted = "Completed"
	// ConditionCredentialsPublished indicates whether generated credentials were written to a Secret.
	ConditionCredentialsPublished = "CredentialsPublished"
)

// AuthProtocol is the protocol used for the registration.
type AuthProtocol string

const (
	AuthProtocolCAS   AuthProtocol = "CAS"
	AuthProtocolOIDC  AuthProtocol = "OIDC"
	AuthProtocolOAuth AuthProtocol = "OAUTH"
)

// AuthRegistrationSpec defines the desired state of AuthRegistration
type AuthRegistrationSpec struct {
	// Protocol defines which authentication protocol should be used for registration.
	// Supported values are CAS, OIDC and OAUTH.
	// +kubebuilder:validation:Enum=CAS;OIDC;OAUTH
	Protocol AuthProtocol `json:"protocol"`

	// Consumer is the identifier of the consuming application (service name).
	// +kubebuilder:validation:MinLength=1
	Consumer string `json:"consumer"`

	// SecretRef references an optional Secret name where generated credentials should be stored.
	// If this field is omitted, the controller creates and manages a Secret automatically.
	// +optional
	SecretRef *string `json:"secretRef,omitempty"`

	// LogoutURL defines an optional logout URL for single logout integrations.
	// +kubebuilder:validation:Format=uri
	// +optional
	LogoutURL *string `json:"logoutURL,omitempty"`

	// Params contains additional optional protocol-specific parameters.
	// +optional
	Params map[string]string `json:"params,omitempty"`
}

// AuthRegistrationStatus defines the observed state of AuthRegistration.
type AuthRegistrationStatus struct {
	// ResolvedSecretRef is the effective Secret name used to store credentials.
	// It points either to spec.secretRef or to a controller-generated Secret when spec.secretRef is not set.
	// +optional
	ResolvedSecretRef string `json:"resolvedSecretRef,omitempty"`

	// conditions represent the current state of the AuthRegistration resource.
	// Each condition has a unique type and reflects a specific state transition.
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=ar
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Protocol",type="string",JSONPath=".spec.protocol",description="Requested auth protocol"
// +kubebuilder:printcolumn:name="Consumer",type="string",JSONPath=".spec.consumer",description="Registered consumer/service name"
// +kubebuilder:printcolumn:name="Secret",type="string",JSONPath=".status.resolvedSecretRef",description="Effective Secret used for credentials"
// +kubebuilder:printcolumn:name="CredsPublished",type="string",JSONPath=".status.conditions[?(@.type == 'CredentialsPublished')].status",description="Whether credentials were published to a Secret"
// +kubebuilder:printcolumn:name="Completed",type="string",JSONPath=".status.conditions[?(@.type == 'Completed')].status",description="Whether the registration has been completed"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="The age of the resource"

// AuthRegistration is the Schema for the authregistrations API
type AuthRegistration struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of AuthRegistration
	// +required
	Spec AuthRegistrationSpec `json:"spec"`

	// status defines the observed state of AuthRegistration
	// +optional
	Status AuthRegistrationStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// AuthRegistrationList contains a list of AuthRegistration
type AuthRegistrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []AuthRegistration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AuthRegistration{}, &AuthRegistrationList{})
}
