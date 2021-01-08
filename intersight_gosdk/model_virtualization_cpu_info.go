/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-01-04T05:15:49Z.
 *
 * API version: 1.0.9-3144
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// VirtualizationCpuInfo The details of the CPUs on this platform.
type VirtualizationCpuInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Number of cores per CPU, as reported by the manufacturer.
	Cores *int64 `json:"Cores,omitempty"`
	// The vendor provided description of the CPU. For example, Intel Xeon E5-2640 v3 @ 2.60GHz.
	Description *string `json:"Description,omitempty"`
	// Number of CPU sockets available.
	Sockets *int64 `json:"Sockets,omitempty"`
	// Speed of the CPUs in Hertz. For example, 2593749663.
	Speed *int64 `json:"Speed,omitempty"`
	// Manufacturer of the CPU . For example, Intel.
	Vendor               *string `json:"Vendor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationCpuInfo VirtualizationCpuInfo

// NewVirtualizationCpuInfo instantiates a new VirtualizationCpuInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationCpuInfo(classId string, objectType string) *VirtualizationCpuInfo {
	this := VirtualizationCpuInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationCpuInfoWithDefaults instantiates a new VirtualizationCpuInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationCpuInfoWithDefaults() *VirtualizationCpuInfo {
	this := VirtualizationCpuInfo{}
	var classId string = "virtualization.CpuInfo"
	this.ClassId = classId
	var objectType string = "virtualization.CpuInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationCpuInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationCpuInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationCpuInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationCpuInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationCpuInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationCpuInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCores returns the Cores field value if set, zero value otherwise.
func (o *VirtualizationCpuInfo) GetCores() int64 {
	if o == nil || o.Cores == nil {
		var ret int64
		return ret
	}
	return *o.Cores
}

// GetCoresOk returns a tuple with the Cores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCpuInfo) GetCoresOk() (*int64, bool) {
	if o == nil || o.Cores == nil {
		return nil, false
	}
	return o.Cores, true
}

// HasCores returns a boolean if a field has been set.
func (o *VirtualizationCpuInfo) HasCores() bool {
	if o != nil && o.Cores != nil {
		return true
	}

	return false
}

// SetCores gets a reference to the given int64 and assigns it to the Cores field.
func (o *VirtualizationCpuInfo) SetCores(v int64) {
	o.Cores = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VirtualizationCpuInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCpuInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VirtualizationCpuInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VirtualizationCpuInfo) SetDescription(v string) {
	o.Description = &v
}

// GetSockets returns the Sockets field value if set, zero value otherwise.
func (o *VirtualizationCpuInfo) GetSockets() int64 {
	if o == nil || o.Sockets == nil {
		var ret int64
		return ret
	}
	return *o.Sockets
}

// GetSocketsOk returns a tuple with the Sockets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCpuInfo) GetSocketsOk() (*int64, bool) {
	if o == nil || o.Sockets == nil {
		return nil, false
	}
	return o.Sockets, true
}

// HasSockets returns a boolean if a field has been set.
func (o *VirtualizationCpuInfo) HasSockets() bool {
	if o != nil && o.Sockets != nil {
		return true
	}

	return false
}

// SetSockets gets a reference to the given int64 and assigns it to the Sockets field.
func (o *VirtualizationCpuInfo) SetSockets(v int64) {
	o.Sockets = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *VirtualizationCpuInfo) GetSpeed() int64 {
	if o == nil || o.Speed == nil {
		var ret int64
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCpuInfo) GetSpeedOk() (*int64, bool) {
	if o == nil || o.Speed == nil {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *VirtualizationCpuInfo) HasSpeed() bool {
	if o != nil && o.Speed != nil {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int64 and assigns it to the Speed field.
func (o *VirtualizationCpuInfo) SetSpeed(v int64) {
	o.Speed = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *VirtualizationCpuInfo) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCpuInfo) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *VirtualizationCpuInfo) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *VirtualizationCpuInfo) SetVendor(v string) {
	o.Vendor = &v
}

func (o VirtualizationCpuInfo) MarshalJSON() ([]byte, error) {
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
	if o.Cores != nil {
		toSerialize["Cores"] = o.Cores
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Sockets != nil {
		toSerialize["Sockets"] = o.Sockets
	}
	if o.Speed != nil {
		toSerialize["Speed"] = o.Speed
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationCpuInfo) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationCpuInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Number of cores per CPU, as reported by the manufacturer.
		Cores *int64 `json:"Cores,omitempty"`
		// The vendor provided description of the CPU. For example, Intel Xeon E5-2640 v3 @ 2.60GHz.
		Description *string `json:"Description,omitempty"`
		// Number of CPU sockets available.
		Sockets *int64 `json:"Sockets,omitempty"`
		// Speed of the CPUs in Hertz. For example, 2593749663.
		Speed *int64 `json:"Speed,omitempty"`
		// Manufacturer of the CPU . For example, Intel.
		Vendor *string `json:"Vendor,omitempty"`
	}

	varVirtualizationCpuInfoWithoutEmbeddedStruct := VirtualizationCpuInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationCpuInfoWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationCpuInfo := _VirtualizationCpuInfo{}
		varVirtualizationCpuInfo.ClassId = varVirtualizationCpuInfoWithoutEmbeddedStruct.ClassId
		varVirtualizationCpuInfo.ObjectType = varVirtualizationCpuInfoWithoutEmbeddedStruct.ObjectType
		varVirtualizationCpuInfo.Cores = varVirtualizationCpuInfoWithoutEmbeddedStruct.Cores
		varVirtualizationCpuInfo.Description = varVirtualizationCpuInfoWithoutEmbeddedStruct.Description
		varVirtualizationCpuInfo.Sockets = varVirtualizationCpuInfoWithoutEmbeddedStruct.Sockets
		varVirtualizationCpuInfo.Speed = varVirtualizationCpuInfoWithoutEmbeddedStruct.Speed
		varVirtualizationCpuInfo.Vendor = varVirtualizationCpuInfoWithoutEmbeddedStruct.Vendor
		*o = VirtualizationCpuInfo(varVirtualizationCpuInfo)
	} else {
		return err
	}

	varVirtualizationCpuInfo := _VirtualizationCpuInfo{}

	err = json.Unmarshal(bytes, &varVirtualizationCpuInfo)
	if err == nil {
		o.MoBaseComplexType = varVirtualizationCpuInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cores")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Sockets")
		delete(additionalProperties, "Speed")
		delete(additionalProperties, "Vendor")

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

type NullableVirtualizationCpuInfo struct {
	value *VirtualizationCpuInfo
	isSet bool
}

func (v NullableVirtualizationCpuInfo) Get() *VirtualizationCpuInfo {
	return v.value
}

func (v *NullableVirtualizationCpuInfo) Set(val *VirtualizationCpuInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationCpuInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationCpuInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationCpuInfo(val *VirtualizationCpuInfo) *NullableVirtualizationCpuInfo {
	return &NullableVirtualizationCpuInfo{value: val, isSet: true}
}

func (v NullableVirtualizationCpuInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationCpuInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
