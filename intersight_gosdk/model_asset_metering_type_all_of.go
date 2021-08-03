/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-07-21T16:37:30Z.
 *
 * API version: 1.0.9-4403
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// AssetMeteringTypeAllOf Definition of the list of properties defined in 'asset.MeteringType', excluding properties defined in parent classes.
type AssetMeteringTypeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Metric type used to calculate metering for the device sent from the IB Contract. example  Node, vMemory, vCPU. * `None` - A default value to catch cases where metric type is not correctly detected. * `Node` - The metering of the device is on the basis of Power state. * `Storage` - The metering of the device is on the basis of used Storage. * `vMemory` - The metering of the device is on the basis of VM Memory. * `vCPU` - The metering of the device is on the basis of vCPU. * `vStorage` - The metering of the device is on the basis of used virtual Storage. * `Switch` - The metering of the device is on the basis of Switch.
	Name *string `json:"Name,omitempty"`
	// Metric unit used to calculate metering for the device sent from the IB Contract. example  Node, GiB, Cores. * `None` - A default value to catch cases where metric unit is not correctly detected. * `Node` - It is applicable for Node Metric type. * `GiB` - It is applicable for VMemory, vStorage and Storage Metric types. * `TiB` - It is applicable for VMemory, vStorage and Storage Metric types. * `Cores` - It is applicable for vCPU Metric type. * `Switch` - It is applicable for Switch Metric type. * `Port` - It is applicable for Switch Metric type.
	Unit                 *string `json:"Unit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetMeteringTypeAllOf AssetMeteringTypeAllOf

// NewAssetMeteringTypeAllOf instantiates a new AssetMeteringTypeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetMeteringTypeAllOf(classId string, objectType string) *AssetMeteringTypeAllOf {
	this := AssetMeteringTypeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var name string = "None"
	this.Name = &name
	var unit string = "None"
	this.Unit = &unit
	return &this
}

// NewAssetMeteringTypeAllOfWithDefaults instantiates a new AssetMeteringTypeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetMeteringTypeAllOfWithDefaults() *AssetMeteringTypeAllOf {
	this := AssetMeteringTypeAllOf{}
	var classId string = "asset.MeteringType"
	this.ClassId = classId
	var objectType string = "asset.MeteringType"
	this.ObjectType = objectType
	var name string = "None"
	this.Name = &name
	var unit string = "None"
	this.Unit = &unit
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetMeteringTypeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetMeteringTypeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetMeteringTypeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetMeteringTypeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetMeteringTypeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetMeteringTypeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AssetMeteringTypeAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetMeteringTypeAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AssetMeteringTypeAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AssetMeteringTypeAllOf) SetName(v string) {
	o.Name = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *AssetMeteringTypeAllOf) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetMeteringTypeAllOf) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *AssetMeteringTypeAllOf) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *AssetMeteringTypeAllOf) SetUnit(v string) {
	o.Unit = &v
}

func (o AssetMeteringTypeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Unit != nil {
		toSerialize["Unit"] = o.Unit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetMeteringTypeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetMeteringTypeAllOf := _AssetMeteringTypeAllOf{}

	if err = json.Unmarshal(bytes, &varAssetMeteringTypeAllOf); err == nil {
		*o = AssetMeteringTypeAllOf(varAssetMeteringTypeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Unit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetMeteringTypeAllOf struct {
	value *AssetMeteringTypeAllOf
	isSet bool
}

func (v NullableAssetMeteringTypeAllOf) Get() *AssetMeteringTypeAllOf {
	return v.value
}

func (v *NullableAssetMeteringTypeAllOf) Set(val *AssetMeteringTypeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetMeteringTypeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetMeteringTypeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetMeteringTypeAllOf(val *AssetMeteringTypeAllOf) *NullableAssetMeteringTypeAllOf {
	return &NullableAssetMeteringTypeAllOf{value: val, isSet: true}
}

func (v NullableAssetMeteringTypeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetMeteringTypeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}