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

// VnicLunAllOf Definition of the list of properties defined in 'vnic.Lun', excluding properties defined in parent classes.
type VnicLunAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Specifies LUN is bootable.
	Bootable *bool `json:"Bootable,omitempty"`
	// The Identifier of the LUN.
	LunId                *int64 `json:"LunId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicLunAllOf VnicLunAllOf

// NewVnicLunAllOf instantiates a new VnicLunAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicLunAllOf(classId string, objectType string) *VnicLunAllOf {
	this := VnicLunAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVnicLunAllOfWithDefaults instantiates a new VnicLunAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicLunAllOfWithDefaults() *VnicLunAllOf {
	this := VnicLunAllOf{}
	var classId string = "vnic.Lun"
	this.ClassId = classId
	var objectType string = "vnic.Lun"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicLunAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicLunAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicLunAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicLunAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicLunAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicLunAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBootable returns the Bootable field value if set, zero value otherwise.
func (o *VnicLunAllOf) GetBootable() bool {
	if o == nil || o.Bootable == nil {
		var ret bool
		return ret
	}
	return *o.Bootable
}

// GetBootableOk returns a tuple with the Bootable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLunAllOf) GetBootableOk() (*bool, bool) {
	if o == nil || o.Bootable == nil {
		return nil, false
	}
	return o.Bootable, true
}

// HasBootable returns a boolean if a field has been set.
func (o *VnicLunAllOf) HasBootable() bool {
	if o != nil && o.Bootable != nil {
		return true
	}

	return false
}

// SetBootable gets a reference to the given bool and assigns it to the Bootable field.
func (o *VnicLunAllOf) SetBootable(v bool) {
	o.Bootable = &v
}

// GetLunId returns the LunId field value if set, zero value otherwise.
func (o *VnicLunAllOf) GetLunId() int64 {
	if o == nil || o.LunId == nil {
		var ret int64
		return ret
	}
	return *o.LunId
}

// GetLunIdOk returns a tuple with the LunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLunAllOf) GetLunIdOk() (*int64, bool) {
	if o == nil || o.LunId == nil {
		return nil, false
	}
	return o.LunId, true
}

// HasLunId returns a boolean if a field has been set.
func (o *VnicLunAllOf) HasLunId() bool {
	if o != nil && o.LunId != nil {
		return true
	}

	return false
}

// SetLunId gets a reference to the given int64 and assigns it to the LunId field.
func (o *VnicLunAllOf) SetLunId(v int64) {
	o.LunId = &v
}

func (o VnicLunAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Bootable != nil {
		toSerialize["Bootable"] = o.Bootable
	}
	if o.LunId != nil {
		toSerialize["LunId"] = o.LunId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicLunAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicLunAllOf := _VnicLunAllOf{}

	if err = json.Unmarshal(bytes, &varVnicLunAllOf); err == nil {
		*o = VnicLunAllOf(varVnicLunAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Bootable")
		delete(additionalProperties, "LunId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicLunAllOf struct {
	value *VnicLunAllOf
	isSet bool
}

func (v NullableVnicLunAllOf) Get() *VnicLunAllOf {
	return v.value
}

func (v *NullableVnicLunAllOf) Set(val *VnicLunAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicLunAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicLunAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicLunAllOf(val *VnicLunAllOf) *NullableVnicLunAllOf {
	return &NullableVnicLunAllOf{value: val, isSet: true}
}

func (v NullableVnicLunAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicLunAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
