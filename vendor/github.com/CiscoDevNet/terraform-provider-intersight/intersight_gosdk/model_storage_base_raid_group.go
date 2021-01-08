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

// StorageBaseRaidGroup A RAID group unifies a number of physical disks into one logical unit and distributes data across those disks. RAID groups can be configured with a particular protection level depending on the performance, capacity, and redundancy needs of the environment.
type StorageBaseRaidGroup struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The RAID level associated with the group of disks. * `Unknown` - Default unknown RAID type. * `RAID0` - RAID0 splits (\"stripes\") data evenly across two or more disks, without parity information. * `RAID1` - RAID1 requires a minimum of two disks to work, and provides data redundancy and failover. * `RAID4` - RAID4 stripes block level data and dedicates a disk to parity. * `RAID5` - RAID5  distributes striping and parity at a block level. * `RAID6` - RAID6 level operates like RAID5 with distributed parity and striping. * `RAID10` - RAID10 requires a minimum of four disks in the array. It stripes across disks for higher performance, and mirrors for redundancy. * `RAIDDP` - RAIDDP uses up to two spare disks to replace and reconstruct the data from up to two simultaneously failed disks within the RAID group. * `RAIDTEC` - With RAIDTEC protection, use up to three spare disks to replace and reconstruct the data from up to three simultaneously failed disks within the RAID group.
	Level *string `json:"Level,omitempty"`
	// Human readable name of the RAID group.
	Name                 *string                     `json:"Name,omitempty"`
	StorageUtilization   NullableStorageBaseCapacity `json:"StorageUtilization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseRaidGroup StorageBaseRaidGroup

// NewStorageBaseRaidGroup instantiates a new StorageBaseRaidGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseRaidGroup(classId string, objectType string) *StorageBaseRaidGroup {
	this := StorageBaseRaidGroup{}
	this.ClassId = classId
	this.ObjectType = objectType
	var level string = "Unknown"
	this.Level = &level
	return &this
}

// NewStorageBaseRaidGroupWithDefaults instantiates a new StorageBaseRaidGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseRaidGroupWithDefaults() *StorageBaseRaidGroup {
	this := StorageBaseRaidGroup{}
	var classId string = "storage.HitachiParityGroup"
	this.ClassId = classId
	var objectType string = "storage.HitachiParityGroup"
	this.ObjectType = objectType
	var level string = "Unknown"
	this.Level = &level
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseRaidGroup) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseRaidGroup) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseRaidGroup) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseRaidGroup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseRaidGroup) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseRaidGroup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *StorageBaseRaidGroup) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseRaidGroup) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *StorageBaseRaidGroup) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *StorageBaseRaidGroup) SetLevel(v string) {
	o.Level = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseRaidGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseRaidGroup) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseRaidGroup) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseRaidGroup) SetName(v string) {
	o.Name = &v
}

// GetStorageUtilization returns the StorageUtilization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageBaseRaidGroup) GetStorageUtilization() StorageBaseCapacity {
	if o == nil || o.StorageUtilization.Get() == nil {
		var ret StorageBaseCapacity
		return ret
	}
	return *o.StorageUtilization.Get()
}

// GetStorageUtilizationOk returns a tuple with the StorageUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageBaseRaidGroup) GetStorageUtilizationOk() (*StorageBaseCapacity, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageUtilization.Get(), o.StorageUtilization.IsSet()
}

// HasStorageUtilization returns a boolean if a field has been set.
func (o *StorageBaseRaidGroup) HasStorageUtilization() bool {
	if o != nil && o.StorageUtilization.IsSet() {
		return true
	}

	return false
}

// SetStorageUtilization gets a reference to the given NullableStorageBaseCapacity and assigns it to the StorageUtilization field.
func (o *StorageBaseRaidGroup) SetStorageUtilization(v StorageBaseCapacity) {
	o.StorageUtilization.Set(&v)
}

// SetStorageUtilizationNil sets the value for StorageUtilization to be an explicit nil
func (o *StorageBaseRaidGroup) SetStorageUtilizationNil() {
	o.StorageUtilization.Set(nil)
}

// UnsetStorageUtilization ensures that no value is present for StorageUtilization, not even an explicit nil
func (o *StorageBaseRaidGroup) UnsetStorageUtilization() {
	o.StorageUtilization.Unset()
}

func (o StorageBaseRaidGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Level != nil {
		toSerialize["Level"] = o.Level
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.StorageUtilization.IsSet() {
		toSerialize["StorageUtilization"] = o.StorageUtilization.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseRaidGroup) UnmarshalJSON(bytes []byte) (err error) {
	type StorageBaseRaidGroupWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The RAID level associated with the group of disks. * `Unknown` - Default unknown RAID type. * `RAID0` - RAID0 splits (\"stripes\") data evenly across two or more disks, without parity information. * `RAID1` - RAID1 requires a minimum of two disks to work, and provides data redundancy and failover. * `RAID4` - RAID4 stripes block level data and dedicates a disk to parity. * `RAID5` - RAID5  distributes striping and parity at a block level. * `RAID6` - RAID6 level operates like RAID5 with distributed parity and striping. * `RAID10` - RAID10 requires a minimum of four disks in the array. It stripes across disks for higher performance, and mirrors for redundancy. * `RAIDDP` - RAIDDP uses up to two spare disks to replace and reconstruct the data from up to two simultaneously failed disks within the RAID group. * `RAIDTEC` - With RAIDTEC protection, use up to three spare disks to replace and reconstruct the data from up to three simultaneously failed disks within the RAID group.
		Level *string `json:"Level,omitempty"`
		// Human readable name of the RAID group.
		Name               *string                     `json:"Name,omitempty"`
		StorageUtilization NullableStorageBaseCapacity `json:"StorageUtilization,omitempty"`
	}

	varStorageBaseRaidGroupWithoutEmbeddedStruct := StorageBaseRaidGroupWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageBaseRaidGroupWithoutEmbeddedStruct)
	if err == nil {
		varStorageBaseRaidGroup := _StorageBaseRaidGroup{}
		varStorageBaseRaidGroup.ClassId = varStorageBaseRaidGroupWithoutEmbeddedStruct.ClassId
		varStorageBaseRaidGroup.ObjectType = varStorageBaseRaidGroupWithoutEmbeddedStruct.ObjectType
		varStorageBaseRaidGroup.Level = varStorageBaseRaidGroupWithoutEmbeddedStruct.Level
		varStorageBaseRaidGroup.Name = varStorageBaseRaidGroupWithoutEmbeddedStruct.Name
		varStorageBaseRaidGroup.StorageUtilization = varStorageBaseRaidGroupWithoutEmbeddedStruct.StorageUtilization
		*o = StorageBaseRaidGroup(varStorageBaseRaidGroup)
	} else {
		return err
	}

	varStorageBaseRaidGroup := _StorageBaseRaidGroup{}

	err = json.Unmarshal(bytes, &varStorageBaseRaidGroup)
	if err == nil {
		o.MoBaseMo = varStorageBaseRaidGroup.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Level")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "StorageUtilization")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableStorageBaseRaidGroup struct {
	value *StorageBaseRaidGroup
	isSet bool
}

func (v NullableStorageBaseRaidGroup) Get() *StorageBaseRaidGroup {
	return v.value
}

func (v *NullableStorageBaseRaidGroup) Set(val *StorageBaseRaidGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseRaidGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseRaidGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseRaidGroup(val *StorageBaseRaidGroup) *NullableStorageBaseRaidGroup {
	return &NullableStorageBaseRaidGroup{value: val, isSet: true}
}

func (v NullableStorageBaseRaidGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseRaidGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
