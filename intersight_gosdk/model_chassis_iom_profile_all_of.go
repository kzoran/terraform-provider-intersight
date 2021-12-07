/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4929
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// ChassisIomProfileAllOf Definition of the list of properties defined in 'chassis.IomProfile', excluding properties defined in parent classes.
type ChassisIomProfileAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string                     `json:"ObjectType"`
	ConfigChanges NullablePolicyConfigChange `json:"ConfigChanges,omitempty"`
	// IOM in chassis for which IOM profile is applicable. or which is attached to a Fabric Interconnect managed by Intersight. * `IOMA` - IOM on left side of chassis. * `IOMB` - IOM on right side of chassis.
	IomEntity            *string                               `json:"IomEntity,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	Profile              *ChassisProfileRelationship           `json:"Profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChassisIomProfileAllOf ChassisIomProfileAllOf

// NewChassisIomProfileAllOf instantiates a new ChassisIomProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChassisIomProfileAllOf(classId string, objectType string) *ChassisIomProfileAllOf {
	this := ChassisIomProfileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var iomEntity string = "IOMA"
	this.IomEntity = &iomEntity
	return &this
}

// NewChassisIomProfileAllOfWithDefaults instantiates a new ChassisIomProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChassisIomProfileAllOfWithDefaults() *ChassisIomProfileAllOf {
	this := ChassisIomProfileAllOf{}
	var classId string = "chassis.IomProfile"
	this.ClassId = classId
	var objectType string = "chassis.IomProfile"
	this.ObjectType = objectType
	var iomEntity string = "IOMA"
	this.IomEntity = &iomEntity
	return &this
}

// GetClassId returns the ClassId field value
func (o *ChassisIomProfileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ChassisIomProfileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ChassisIomProfileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ChassisIomProfileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ChassisIomProfileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ChassisIomProfileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigChanges returns the ConfigChanges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisIomProfileAllOf) GetConfigChanges() PolicyConfigChange {
	if o == nil || o.ConfigChanges.Get() == nil {
		var ret PolicyConfigChange
		return ret
	}
	return *o.ConfigChanges.Get()
}

// GetConfigChangesOk returns a tuple with the ConfigChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisIomProfileAllOf) GetConfigChangesOk() (*PolicyConfigChange, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigChanges.Get(), o.ConfigChanges.IsSet()
}

// HasConfigChanges returns a boolean if a field has been set.
func (o *ChassisIomProfileAllOf) HasConfigChanges() bool {
	if o != nil && o.ConfigChanges.IsSet() {
		return true
	}

	return false
}

// SetConfigChanges gets a reference to the given NullablePolicyConfigChange and assigns it to the ConfigChanges field.
func (o *ChassisIomProfileAllOf) SetConfigChanges(v PolicyConfigChange) {
	o.ConfigChanges.Set(&v)
}

// SetConfigChangesNil sets the value for ConfigChanges to be an explicit nil
func (o *ChassisIomProfileAllOf) SetConfigChangesNil() {
	o.ConfigChanges.Set(nil)
}

// UnsetConfigChanges ensures that no value is present for ConfigChanges, not even an explicit nil
func (o *ChassisIomProfileAllOf) UnsetConfigChanges() {
	o.ConfigChanges.Unset()
}

// GetIomEntity returns the IomEntity field value if set, zero value otherwise.
func (o *ChassisIomProfileAllOf) GetIomEntity() string {
	if o == nil || o.IomEntity == nil {
		var ret string
		return ret
	}
	return *o.IomEntity
}

// GetIomEntityOk returns a tuple with the IomEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisIomProfileAllOf) GetIomEntityOk() (*string, bool) {
	if o == nil || o.IomEntity == nil {
		return nil, false
	}
	return o.IomEntity, true
}

// HasIomEntity returns a boolean if a field has been set.
func (o *ChassisIomProfileAllOf) HasIomEntity() bool {
	if o != nil && o.IomEntity != nil {
		return true
	}

	return false
}

// SetIomEntity gets a reference to the given string and assigns it to the IomEntity field.
func (o *ChassisIomProfileAllOf) SetIomEntity(v string) {
	o.IomEntity = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ChassisIomProfileAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisIomProfileAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ChassisIomProfileAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ChassisIomProfileAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *ChassisIomProfileAllOf) GetProfile() ChassisProfileRelationship {
	if o == nil || o.Profile == nil {
		var ret ChassisProfileRelationship
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisIomProfileAllOf) GetProfileOk() (*ChassisProfileRelationship, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *ChassisIomProfileAllOf) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given ChassisProfileRelationship and assigns it to the Profile field.
func (o *ChassisIomProfileAllOf) SetProfile(v ChassisProfileRelationship) {
	o.Profile = &v
}

func (o ChassisIomProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigChanges.IsSet() {
		toSerialize["ConfigChanges"] = o.ConfigChanges.Get()
	}
	if o.IomEntity != nil {
		toSerialize["IomEntity"] = o.IomEntity
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profile != nil {
		toSerialize["Profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChassisIomProfileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varChassisIomProfileAllOf := _ChassisIomProfileAllOf{}

	if err = json.Unmarshal(bytes, &varChassisIomProfileAllOf); err == nil {
		*o = ChassisIomProfileAllOf(varChassisIomProfileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigChanges")
		delete(additionalProperties, "IomEntity")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChassisIomProfileAllOf struct {
	value *ChassisIomProfileAllOf
	isSet bool
}

func (v NullableChassisIomProfileAllOf) Get() *ChassisIomProfileAllOf {
	return v.value
}

func (v *NullableChassisIomProfileAllOf) Set(val *ChassisIomProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableChassisIomProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableChassisIomProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChassisIomProfileAllOf(val *ChassisIomProfileAllOf) *NullableChassisIomProfileAllOf {
	return &NullableChassisIomProfileAllOf{value: val, isSet: true}
}

func (v NullableChassisIomProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChassisIomProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
