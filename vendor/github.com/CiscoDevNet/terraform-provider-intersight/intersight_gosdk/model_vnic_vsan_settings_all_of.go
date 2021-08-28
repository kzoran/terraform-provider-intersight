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

// VnicVsanSettingsAllOf Definition of the list of properties defined in 'vnic.VsanSettings', excluding properties defined in parent classes.
type VnicVsanSettingsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Default VLAN of the virtual interface in Standalone Rack server. Setting the value to 0 is equivalent to None and will not associate any Default VLAN to the traffic on the virtual interface (0-4094).
	DefaultVlanId *int64 `json:"DefaultVlanId,omitempty"`
	// VSAN ID of the virtual interface in FI attached server (1-4094).
	Id                   *int64 `json:"Id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicVsanSettingsAllOf VnicVsanSettingsAllOf

// NewVnicVsanSettingsAllOf instantiates a new VnicVsanSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicVsanSettingsAllOf(classId string, objectType string) *VnicVsanSettingsAllOf {
	this := VnicVsanSettingsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var defaultVlanId int64 = 0
	this.DefaultVlanId = &defaultVlanId
	var id int64 = 1
	this.Id = &id
	return &this
}

// NewVnicVsanSettingsAllOfWithDefaults instantiates a new VnicVsanSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicVsanSettingsAllOfWithDefaults() *VnicVsanSettingsAllOf {
	this := VnicVsanSettingsAllOf{}
	var classId string = "vnic.VsanSettings"
	this.ClassId = classId
	var objectType string = "vnic.VsanSettings"
	this.ObjectType = objectType
	var defaultVlanId int64 = 0
	this.DefaultVlanId = &defaultVlanId
	var id int64 = 1
	this.Id = &id
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicVsanSettingsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicVsanSettingsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicVsanSettingsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicVsanSettingsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicVsanSettingsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicVsanSettingsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultVlanId returns the DefaultVlanId field value if set, zero value otherwise.
func (o *VnicVsanSettingsAllOf) GetDefaultVlanId() int64 {
	if o == nil || o.DefaultVlanId == nil {
		var ret int64
		return ret
	}
	return *o.DefaultVlanId
}

// GetDefaultVlanIdOk returns a tuple with the DefaultVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVsanSettingsAllOf) GetDefaultVlanIdOk() (*int64, bool) {
	if o == nil || o.DefaultVlanId == nil {
		return nil, false
	}
	return o.DefaultVlanId, true
}

// HasDefaultVlanId returns a boolean if a field has been set.
func (o *VnicVsanSettingsAllOf) HasDefaultVlanId() bool {
	if o != nil && o.DefaultVlanId != nil {
		return true
	}

	return false
}

// SetDefaultVlanId gets a reference to the given int64 and assigns it to the DefaultVlanId field.
func (o *VnicVsanSettingsAllOf) SetDefaultVlanId(v int64) {
	o.DefaultVlanId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VnicVsanSettingsAllOf) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVsanSettingsAllOf) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VnicVsanSettingsAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *VnicVsanSettingsAllOf) SetId(v int64) {
	o.Id = &v
}

func (o VnicVsanSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DefaultVlanId != nil {
		toSerialize["DefaultVlanId"] = o.DefaultVlanId
	}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicVsanSettingsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicVsanSettingsAllOf := _VnicVsanSettingsAllOf{}

	if err = json.Unmarshal(bytes, &varVnicVsanSettingsAllOf); err == nil {
		*o = VnicVsanSettingsAllOf(varVnicVsanSettingsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DefaultVlanId")
		delete(additionalProperties, "Id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicVsanSettingsAllOf struct {
	value *VnicVsanSettingsAllOf
	isSet bool
}

func (v NullableVnicVsanSettingsAllOf) Get() *VnicVsanSettingsAllOf {
	return v.value
}

func (v *NullableVnicVsanSettingsAllOf) Set(val *VnicVsanSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicVsanSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicVsanSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicVsanSettingsAllOf(val *VnicVsanSettingsAllOf) *NullableVnicVsanSettingsAllOf {
	return &NullableVnicVsanSettingsAllOf{value: val, isSet: true}
}

func (v NullableVnicVsanSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicVsanSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
