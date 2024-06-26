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

type HardcodedAttributeMapperInitParameters struct {

	// The name of the LDAP attribute to set.
	// Name of the LDAP attribute
	AttributeName *string `json:"attributeName,omitempty" tf:"attribute_name,omitempty"`

	// The value to set to the LDAP attribute. You can hardcode any value like 'foo'.
	// Value of the LDAP attribute. You can hardcode any value like 'foo'
	AttributeValue *string `json:"attributeValue,omitempty" tf:"attribute_value,omitempty"`

	// The ID of the LDAP user federation provider to attach this mapper to.
	// The ldap user federation provider to attach this mapper to.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/ldap/v1alpha1.UserFederation
	LdapUserFederationID *string `json:"ldapUserFederationId,omitempty" tf:"ldap_user_federation_id,omitempty"`

	// Reference to a UserFederation in ldap to populate ldapUserFederationId.
	// +kubebuilder:validation:Optional
	LdapUserFederationIDRef *v1.Reference `json:"ldapUserFederationIdRef,omitempty" tf:"-"`

	// Selector for a UserFederation in ldap to populate ldapUserFederationId.
	// +kubebuilder:validation:Optional
	LdapUserFederationIDSelector *v1.Selector `json:"ldapUserFederationIdSelector,omitempty" tf:"-"`

	// Display name of this mapper when displayed in the console.
	// Display name of the mapper when displayed in the console.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The realm that this LDAP mapper will exist in.
	// The realm in which the ldap user federation provider exists.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`
}

type HardcodedAttributeMapperObservation struct {

	// The name of the LDAP attribute to set.
	// Name of the LDAP attribute
	AttributeName *string `json:"attributeName,omitempty" tf:"attribute_name,omitempty"`

	// The value to set to the LDAP attribute. You can hardcode any value like 'foo'.
	// Value of the LDAP attribute. You can hardcode any value like 'foo'
	AttributeValue *string `json:"attributeValue,omitempty" tf:"attribute_value,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the LDAP user federation provider to attach this mapper to.
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationID *string `json:"ldapUserFederationId,omitempty" tf:"ldap_user_federation_id,omitempty"`

	// Display name of this mapper when displayed in the console.
	// Display name of the mapper when displayed in the console.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The realm that this LDAP mapper will exist in.
	// The realm in which the ldap user federation provider exists.
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`
}

type HardcodedAttributeMapperParameters struct {

	// The name of the LDAP attribute to set.
	// Name of the LDAP attribute
	// +kubebuilder:validation:Optional
	AttributeName *string `json:"attributeName,omitempty" tf:"attribute_name,omitempty"`

	// The value to set to the LDAP attribute. You can hardcode any value like 'foo'.
	// Value of the LDAP attribute. You can hardcode any value like 'foo'
	// +kubebuilder:validation:Optional
	AttributeValue *string `json:"attributeValue,omitempty" tf:"attribute_value,omitempty"`

	// The ID of the LDAP user federation provider to attach this mapper to.
	// The ldap user federation provider to attach this mapper to.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/ldap/v1alpha1.UserFederation
	// +kubebuilder:validation:Optional
	LdapUserFederationID *string `json:"ldapUserFederationId,omitempty" tf:"ldap_user_federation_id,omitempty"`

	// Reference to a UserFederation in ldap to populate ldapUserFederationId.
	// +kubebuilder:validation:Optional
	LdapUserFederationIDRef *v1.Reference `json:"ldapUserFederationIdRef,omitempty" tf:"-"`

	// Selector for a UserFederation in ldap to populate ldapUserFederationId.
	// +kubebuilder:validation:Optional
	LdapUserFederationIDSelector *v1.Selector `json:"ldapUserFederationIdSelector,omitempty" tf:"-"`

	// Display name of this mapper when displayed in the console.
	// Display name of the mapper when displayed in the console.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The realm that this LDAP mapper will exist in.
	// The realm in which the ldap user federation provider exists.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	// +kubebuilder:validation:Optional
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`
}

// HardcodedAttributeMapperSpec defines the desired state of HardcodedAttributeMapper
type HardcodedAttributeMapperSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HardcodedAttributeMapperParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider HardcodedAttributeMapperInitParameters `json:"initProvider,omitempty"`
}

// HardcodedAttributeMapperStatus defines the observed state of HardcodedAttributeMapper.
type HardcodedAttributeMapperStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HardcodedAttributeMapperObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HardcodedAttributeMapper is the Schema for the HardcodedAttributeMappers API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type HardcodedAttributeMapper struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.attributeName) || (has(self.initProvider) && has(self.initProvider.attributeName))",message="spec.forProvider.attributeName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.attributeValue) || (has(self.initProvider) && has(self.initProvider.attributeValue))",message="spec.forProvider.attributeValue is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   HardcodedAttributeMapperSpec   `json:"spec"`
	Status HardcodedAttributeMapperStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HardcodedAttributeMapperList contains a list of HardcodedAttributeMappers
type HardcodedAttributeMapperList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HardcodedAttributeMapper `json:"items"`
}

// Repository type metadata.
var (
	HardcodedAttributeMapper_Kind             = "HardcodedAttributeMapper"
	HardcodedAttributeMapper_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HardcodedAttributeMapper_Kind}.String()
	HardcodedAttributeMapper_KindAPIVersion   = HardcodedAttributeMapper_Kind + "." + CRDGroupVersion.String()
	HardcodedAttributeMapper_GroupVersionKind = CRDGroupVersion.WithKind(HardcodedAttributeMapper_Kind)
)

func init() {
	SchemeBuilder.Register(&HardcodedAttributeMapper{}, &HardcodedAttributeMapperList{})
}
