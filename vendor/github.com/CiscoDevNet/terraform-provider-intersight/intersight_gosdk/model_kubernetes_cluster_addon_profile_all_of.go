/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4437
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// KubernetesClusterAddonProfileAllOf Definition of the list of properties defined in 'kubernetes.ClusterAddonProfile', excluding properties defined in parent classes.
type KubernetesClusterAddonProfileAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string            `json:"ObjectType"`
	Addons     []KubernetesAddon `json:"Addons,omitempty"`
	// Name of the cluster addon profile.
	Name                 *string                               `json:"Name,omitempty"`
	AssociatedCluster    *KubernetesClusterRelationship        `json:"AssociatedCluster,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesClusterAddonProfileAllOf KubernetesClusterAddonProfileAllOf

// NewKubernetesClusterAddonProfileAllOf instantiates a new KubernetesClusterAddonProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesClusterAddonProfileAllOf(classId string, objectType string) *KubernetesClusterAddonProfileAllOf {
	this := KubernetesClusterAddonProfileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesClusterAddonProfileAllOfWithDefaults instantiates a new KubernetesClusterAddonProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesClusterAddonProfileAllOfWithDefaults() *KubernetesClusterAddonProfileAllOf {
	this := KubernetesClusterAddonProfileAllOf{}
	var classId string = "kubernetes.ClusterAddonProfile"
	this.ClassId = classId
	var objectType string = "kubernetes.ClusterAddonProfile"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesClusterAddonProfileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesClusterAddonProfileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesClusterAddonProfileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesClusterAddonProfileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesClusterAddonProfileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesClusterAddonProfileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAddons returns the Addons field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesClusterAddonProfileAllOf) GetAddons() []KubernetesAddon {
	if o == nil {
		var ret []KubernetesAddon
		return ret
	}
	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterAddonProfileAllOf) GetAddonsOk() (*[]KubernetesAddon, bool) {
	if o == nil || o.Addons == nil {
		return nil, false
	}
	return &o.Addons, true
}

// HasAddons returns a boolean if a field has been set.
func (o *KubernetesClusterAddonProfileAllOf) HasAddons() bool {
	if o != nil && o.Addons != nil {
		return true
	}

	return false
}

// SetAddons gets a reference to the given []KubernetesAddon and assigns it to the Addons field.
func (o *KubernetesClusterAddonProfileAllOf) SetAddons(v []KubernetesAddon) {
	o.Addons = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesClusterAddonProfileAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterAddonProfileAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesClusterAddonProfileAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesClusterAddonProfileAllOf) SetName(v string) {
	o.Name = &v
}

// GetAssociatedCluster returns the AssociatedCluster field value if set, zero value otherwise.
func (o *KubernetesClusterAddonProfileAllOf) GetAssociatedCluster() KubernetesClusterRelationship {
	if o == nil || o.AssociatedCluster == nil {
		var ret KubernetesClusterRelationship
		return ret
	}
	return *o.AssociatedCluster
}

// GetAssociatedClusterOk returns a tuple with the AssociatedCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterAddonProfileAllOf) GetAssociatedClusterOk() (*KubernetesClusterRelationship, bool) {
	if o == nil || o.AssociatedCluster == nil {
		return nil, false
	}
	return o.AssociatedCluster, true
}

// HasAssociatedCluster returns a boolean if a field has been set.
func (o *KubernetesClusterAddonProfileAllOf) HasAssociatedCluster() bool {
	if o != nil && o.AssociatedCluster != nil {
		return true
	}

	return false
}

// SetAssociatedCluster gets a reference to the given KubernetesClusterRelationship and assigns it to the AssociatedCluster field.
func (o *KubernetesClusterAddonProfileAllOf) SetAssociatedCluster(v KubernetesClusterRelationship) {
	o.AssociatedCluster = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesClusterAddonProfileAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterAddonProfileAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesClusterAddonProfileAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesClusterAddonProfileAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o KubernetesClusterAddonProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Addons != nil {
		toSerialize["Addons"] = o.Addons
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.AssociatedCluster != nil {
		toSerialize["AssociatedCluster"] = o.AssociatedCluster
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesClusterAddonProfileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesClusterAddonProfileAllOf := _KubernetesClusterAddonProfileAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesClusterAddonProfileAllOf); err == nil {
		*o = KubernetesClusterAddonProfileAllOf(varKubernetesClusterAddonProfileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Addons")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "AssociatedCluster")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesClusterAddonProfileAllOf struct {
	value *KubernetesClusterAddonProfileAllOf
	isSet bool
}

func (v NullableKubernetesClusterAddonProfileAllOf) Get() *KubernetesClusterAddonProfileAllOf {
	return v.value
}

func (v *NullableKubernetesClusterAddonProfileAllOf) Set(val *KubernetesClusterAddonProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesClusterAddonProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesClusterAddonProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesClusterAddonProfileAllOf(val *KubernetesClusterAddonProfileAllOf) *NullableKubernetesClusterAddonProfileAllOf {
	return &NullableKubernetesClusterAddonProfileAllOf{value: val, isSet: true}
}

func (v NullableKubernetesClusterAddonProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesClusterAddonProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
