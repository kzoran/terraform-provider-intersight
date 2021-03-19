/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-03-10T06:51:24Z.
 *
 * API version: 1.0.9-3942
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// InventoryInventoryMo Complex type representing the inventory MO.
type InventoryInventoryMo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The UCS DN of the MO for which the latest inventory to be fetched. If this property is empty and moId property has the Moid of the MO to be updated, the Moid will be used. If this property is empty and moId is also empty, all the MOs of the given moType will be updated.
	MoDn *string `json:"MoDn,omitempty"`
	// The MO id of an MO for which the latest inventory to be fetched. If this property is empty and moDn property has the UCS DN of the MO to be updated, the DN will be used. If this property is empty and moDn is also empty, all the MOs of the given moType will be updated.
	MoId *string `json:"MoId,omitempty"`
	// The type of the MO for which the latest inventory to be fetched.
	MoType               *string `json:"MoType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryInventoryMo InventoryInventoryMo

// NewInventoryInventoryMo instantiates a new InventoryInventoryMo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryInventoryMo(classId string, objectType string) *InventoryInventoryMo {
	this := InventoryInventoryMo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewInventoryInventoryMoWithDefaults instantiates a new InventoryInventoryMo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryInventoryMoWithDefaults() *InventoryInventoryMo {
	this := InventoryInventoryMo{}
	var classId string = "inventory.InventoryMo"
	this.ClassId = classId
	var objectType string = "inventory.InventoryMo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *InventoryInventoryMo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *InventoryInventoryMo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *InventoryInventoryMo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *InventoryInventoryMo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *InventoryInventoryMo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *InventoryInventoryMo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMoDn returns the MoDn field value if set, zero value otherwise.
func (o *InventoryInventoryMo) GetMoDn() string {
	if o == nil || o.MoDn == nil {
		var ret string
		return ret
	}
	return *o.MoDn
}

// GetMoDnOk returns a tuple with the MoDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryInventoryMo) GetMoDnOk() (*string, bool) {
	if o == nil || o.MoDn == nil {
		return nil, false
	}
	return o.MoDn, true
}

// HasMoDn returns a boolean if a field has been set.
func (o *InventoryInventoryMo) HasMoDn() bool {
	if o != nil && o.MoDn != nil {
		return true
	}

	return false
}

// SetMoDn gets a reference to the given string and assigns it to the MoDn field.
func (o *InventoryInventoryMo) SetMoDn(v string) {
	o.MoDn = &v
}

// GetMoId returns the MoId field value if set, zero value otherwise.
func (o *InventoryInventoryMo) GetMoId() string {
	if o == nil || o.MoId == nil {
		var ret string
		return ret
	}
	return *o.MoId
}

// GetMoIdOk returns a tuple with the MoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryInventoryMo) GetMoIdOk() (*string, bool) {
	if o == nil || o.MoId == nil {
		return nil, false
	}
	return o.MoId, true
}

// HasMoId returns a boolean if a field has been set.
func (o *InventoryInventoryMo) HasMoId() bool {
	if o != nil && o.MoId != nil {
		return true
	}

	return false
}

// SetMoId gets a reference to the given string and assigns it to the MoId field.
func (o *InventoryInventoryMo) SetMoId(v string) {
	o.MoId = &v
}

// GetMoType returns the MoType field value if set, zero value otherwise.
func (o *InventoryInventoryMo) GetMoType() string {
	if o == nil || o.MoType == nil {
		var ret string
		return ret
	}
	return *o.MoType
}

// GetMoTypeOk returns a tuple with the MoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryInventoryMo) GetMoTypeOk() (*string, bool) {
	if o == nil || o.MoType == nil {
		return nil, false
	}
	return o.MoType, true
}

// HasMoType returns a boolean if a field has been set.
func (o *InventoryInventoryMo) HasMoType() bool {
	if o != nil && o.MoType != nil {
		return true
	}

	return false
}

// SetMoType gets a reference to the given string and assigns it to the MoType field.
func (o *InventoryInventoryMo) SetMoType(v string) {
	o.MoType = &v
}

func (o InventoryInventoryMo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MoDn != nil {
		toSerialize["MoDn"] = o.MoDn
	}
	if o.MoId != nil {
		toSerialize["MoId"] = o.MoId
	}
	if o.MoType != nil {
		toSerialize["MoType"] = o.MoType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InventoryInventoryMo) UnmarshalJSON(bytes []byte) (err error) {
	type InventoryInventoryMoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The UCS DN of the MO for which the latest inventory to be fetched. If this property is empty and moId property has the Moid of the MO to be updated, the Moid will be used. If this property is empty and moId is also empty, all the MOs of the given moType will be updated.
		MoDn *string `json:"MoDn,omitempty"`
		// The MO id of an MO for which the latest inventory to be fetched. If this property is empty and moDn property has the UCS DN of the MO to be updated, the DN will be used. If this property is empty and moDn is also empty, all the MOs of the given moType will be updated.
		MoId *string `json:"MoId,omitempty"`
		// The type of the MO for which the latest inventory to be fetched.
		MoType *string `json:"MoType,omitempty"`
	}

	varInventoryInventoryMoWithoutEmbeddedStruct := InventoryInventoryMoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varInventoryInventoryMoWithoutEmbeddedStruct)
	if err == nil {
		varInventoryInventoryMo := _InventoryInventoryMo{}
		varInventoryInventoryMo.ClassId = varInventoryInventoryMoWithoutEmbeddedStruct.ClassId
		varInventoryInventoryMo.ObjectType = varInventoryInventoryMoWithoutEmbeddedStruct.ObjectType
		varInventoryInventoryMo.MoDn = varInventoryInventoryMoWithoutEmbeddedStruct.MoDn
		varInventoryInventoryMo.MoId = varInventoryInventoryMoWithoutEmbeddedStruct.MoId
		varInventoryInventoryMo.MoType = varInventoryInventoryMoWithoutEmbeddedStruct.MoType
		*o = InventoryInventoryMo(varInventoryInventoryMo)
	} else {
		return err
	}

	varInventoryInventoryMo := _InventoryInventoryMo{}

	err = json.Unmarshal(bytes, &varInventoryInventoryMo)
	if err == nil {
		o.MoBaseComplexType = varInventoryInventoryMo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MoDn")
		delete(additionalProperties, "MoId")
		delete(additionalProperties, "MoType")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInventoryInventoryMo struct {
	value *InventoryInventoryMo
	isSet bool
}

func (v NullableInventoryInventoryMo) Get() *InventoryInventoryMo {
	return v.value
}

func (v *NullableInventoryInventoryMo) Set(val *InventoryInventoryMo) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryInventoryMo) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryInventoryMo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryInventoryMo(val *InventoryInventoryMo) *NullableInventoryInventoryMo {
	return &NullableInventoryInventoryMo{value: val, isSet: true}
}

func (v NullableInventoryInventoryMo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryInventoryMo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
