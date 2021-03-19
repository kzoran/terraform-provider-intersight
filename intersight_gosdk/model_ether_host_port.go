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

// EtherHostPort Model object contains the details of host port available on IO card or fabric extender.
type EtherHostPort struct {
	EtherPhysicalPortBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Fabric extender identifier for this port.
	ModuleId *int64 `json:"ModuleId,omitempty"`
	// Host Port Speed of IO card or fabric extender.
	Speed                *string                              `json:"Speed,omitempty"`
	EquipmentIoCardBase  *EquipmentIoCardBaseRelationship     `json:"EquipmentIoCardBase,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EtherHostPort EtherHostPort

// NewEtherHostPort instantiates a new EtherHostPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEtherHostPort(classId string, objectType string) *EtherHostPort {
	this := EtherHostPort{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEtherHostPortWithDefaults instantiates a new EtherHostPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEtherHostPortWithDefaults() *EtherHostPort {
	this := EtherHostPort{}
	var classId string = "ether.HostPort"
	this.ClassId = classId
	var objectType string = "ether.HostPort"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EtherHostPort) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EtherHostPort) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EtherHostPort) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EtherHostPort) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EtherHostPort) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EtherHostPort) SetObjectType(v string) {
	o.ObjectType = v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *EtherHostPort) GetModuleId() int64 {
	if o == nil || o.ModuleId == nil {
		var ret int64
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherHostPort) GetModuleIdOk() (*int64, bool) {
	if o == nil || o.ModuleId == nil {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *EtherHostPort) HasModuleId() bool {
	if o != nil && o.ModuleId != nil {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given int64 and assigns it to the ModuleId field.
func (o *EtherHostPort) SetModuleId(v int64) {
	o.ModuleId = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *EtherHostPort) GetSpeed() string {
	if o == nil || o.Speed == nil {
		var ret string
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherHostPort) GetSpeedOk() (*string, bool) {
	if o == nil || o.Speed == nil {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *EtherHostPort) HasSpeed() bool {
	if o != nil && o.Speed != nil {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given string and assigns it to the Speed field.
func (o *EtherHostPort) SetSpeed(v string) {
	o.Speed = &v
}

// GetEquipmentIoCardBase returns the EquipmentIoCardBase field value if set, zero value otherwise.
func (o *EtherHostPort) GetEquipmentIoCardBase() EquipmentIoCardBaseRelationship {
	if o == nil || o.EquipmentIoCardBase == nil {
		var ret EquipmentIoCardBaseRelationship
		return ret
	}
	return *o.EquipmentIoCardBase
}

// GetEquipmentIoCardBaseOk returns a tuple with the EquipmentIoCardBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherHostPort) GetEquipmentIoCardBaseOk() (*EquipmentIoCardBaseRelationship, bool) {
	if o == nil || o.EquipmentIoCardBase == nil {
		return nil, false
	}
	return o.EquipmentIoCardBase, true
}

// HasEquipmentIoCardBase returns a boolean if a field has been set.
func (o *EtherHostPort) HasEquipmentIoCardBase() bool {
	if o != nil && o.EquipmentIoCardBase != nil {
		return true
	}

	return false
}

// SetEquipmentIoCardBase gets a reference to the given EquipmentIoCardBaseRelationship and assigns it to the EquipmentIoCardBase field.
func (o *EtherHostPort) SetEquipmentIoCardBase(v EquipmentIoCardBaseRelationship) {
	o.EquipmentIoCardBase = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EtherHostPort) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherHostPort) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EtherHostPort) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EtherHostPort) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EtherHostPort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEtherPhysicalPortBase, errEtherPhysicalPortBase := json.Marshal(o.EtherPhysicalPortBase)
	if errEtherPhysicalPortBase != nil {
		return []byte{}, errEtherPhysicalPortBase
	}
	errEtherPhysicalPortBase = json.Unmarshal([]byte(serializedEtherPhysicalPortBase), &toSerialize)
	if errEtherPhysicalPortBase != nil {
		return []byte{}, errEtherPhysicalPortBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ModuleId != nil {
		toSerialize["ModuleId"] = o.ModuleId
	}
	if o.Speed != nil {
		toSerialize["Speed"] = o.Speed
	}
	if o.EquipmentIoCardBase != nil {
		toSerialize["EquipmentIoCardBase"] = o.EquipmentIoCardBase
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EtherHostPort) UnmarshalJSON(bytes []byte) (err error) {
	type EtherHostPortWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Fabric extender identifier for this port.
		ModuleId *int64 `json:"ModuleId,omitempty"`
		// Host Port Speed of IO card or fabric extender.
		Speed               *string                              `json:"Speed,omitempty"`
		EquipmentIoCardBase *EquipmentIoCardBaseRelationship     `json:"EquipmentIoCardBase,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varEtherHostPortWithoutEmbeddedStruct := EtherHostPortWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEtherHostPortWithoutEmbeddedStruct)
	if err == nil {
		varEtherHostPort := _EtherHostPort{}
		varEtherHostPort.ClassId = varEtherHostPortWithoutEmbeddedStruct.ClassId
		varEtherHostPort.ObjectType = varEtherHostPortWithoutEmbeddedStruct.ObjectType
		varEtherHostPort.ModuleId = varEtherHostPortWithoutEmbeddedStruct.ModuleId
		varEtherHostPort.Speed = varEtherHostPortWithoutEmbeddedStruct.Speed
		varEtherHostPort.EquipmentIoCardBase = varEtherHostPortWithoutEmbeddedStruct.EquipmentIoCardBase
		varEtherHostPort.RegisteredDevice = varEtherHostPortWithoutEmbeddedStruct.RegisteredDevice
		*o = EtherHostPort(varEtherHostPort)
	} else {
		return err
	}

	varEtherHostPort := _EtherHostPort{}

	err = json.Unmarshal(bytes, &varEtherHostPort)
	if err == nil {
		o.EtherPhysicalPortBase = varEtherHostPort.EtherPhysicalPortBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ModuleId")
		delete(additionalProperties, "Speed")
		delete(additionalProperties, "EquipmentIoCardBase")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectEtherPhysicalPortBase := reflect.ValueOf(o.EtherPhysicalPortBase)
		for i := 0; i < reflectEtherPhysicalPortBase.Type().NumField(); i++ {
			t := reflectEtherPhysicalPortBase.Type().Field(i)

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

type NullableEtherHostPort struct {
	value *EtherHostPort
	isSet bool
}

func (v NullableEtherHostPort) Get() *EtherHostPort {
	return v.value
}

func (v *NullableEtherHostPort) Set(val *EtherHostPort) {
	v.value = val
	v.isSet = true
}

func (v NullableEtherHostPort) IsSet() bool {
	return v.isSet
}

func (v *NullableEtherHostPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEtherHostPort(val *EtherHostPort) *NullableEtherHostPort {
	return &NullableEtherHostPort{value: val, isSet: true}
}

func (v NullableEtherHostPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEtherHostPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
