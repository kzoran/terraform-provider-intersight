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
	"reflect"
	"strings"
)

// StorageHitachiArrayUtilization Storage space consumption of Hitachi Virtual Storage.
type StorageHitachiArrayUtilization struct {
	StorageBaseCapacity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Ratio of mapped sectors within a volume versus the amount of physical space the data occupies after data compression and deduplication. The data reduction ratio does not include thin provisioning savings. For example, a data reduction ratio of 5.0 means that for every 5 MB the host writes to the array, 1 MB is stored on the array's flash modules.
	DataReduction *float32 `json:"DataReduction,omitempty"`
	// Total provisioned storage capacity in Hitachi Virtual Storage, represented in bytes.
	Provisioned          *int64 `json:"Provisioned,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiArrayUtilization StorageHitachiArrayUtilization

// NewStorageHitachiArrayUtilization instantiates a new StorageHitachiArrayUtilization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiArrayUtilization(classId string, objectType string) *StorageHitachiArrayUtilization {
	this := StorageHitachiArrayUtilization{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiArrayUtilizationWithDefaults instantiates a new StorageHitachiArrayUtilization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiArrayUtilizationWithDefaults() *StorageHitachiArrayUtilization {
	this := StorageHitachiArrayUtilization{}
	var classId string = "storage.HitachiArrayUtilization"
	this.ClassId = classId
	var objectType string = "storage.HitachiArrayUtilization"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiArrayUtilization) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiArrayUtilization) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiArrayUtilization) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiArrayUtilization) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiArrayUtilization) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiArrayUtilization) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDataReduction returns the DataReduction field value if set, zero value otherwise.
func (o *StorageHitachiArrayUtilization) GetDataReduction() float32 {
	if o == nil || o.DataReduction == nil {
		var ret float32
		return ret
	}
	return *o.DataReduction
}

// GetDataReductionOk returns a tuple with the DataReduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiArrayUtilization) GetDataReductionOk() (*float32, bool) {
	if o == nil || o.DataReduction == nil {
		return nil, false
	}
	return o.DataReduction, true
}

// HasDataReduction returns a boolean if a field has been set.
func (o *StorageHitachiArrayUtilization) HasDataReduction() bool {
	if o != nil && o.DataReduction != nil {
		return true
	}

	return false
}

// SetDataReduction gets a reference to the given float32 and assigns it to the DataReduction field.
func (o *StorageHitachiArrayUtilization) SetDataReduction(v float32) {
	o.DataReduction = &v
}

// GetProvisioned returns the Provisioned field value if set, zero value otherwise.
func (o *StorageHitachiArrayUtilization) GetProvisioned() int64 {
	if o == nil || o.Provisioned == nil {
		var ret int64
		return ret
	}
	return *o.Provisioned
}

// GetProvisionedOk returns a tuple with the Provisioned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiArrayUtilization) GetProvisionedOk() (*int64, bool) {
	if o == nil || o.Provisioned == nil {
		return nil, false
	}
	return o.Provisioned, true
}

// HasProvisioned returns a boolean if a field has been set.
func (o *StorageHitachiArrayUtilization) HasProvisioned() bool {
	if o != nil && o.Provisioned != nil {
		return true
	}

	return false
}

// SetProvisioned gets a reference to the given int64 and assigns it to the Provisioned field.
func (o *StorageHitachiArrayUtilization) SetProvisioned(v int64) {
	o.Provisioned = &v
}

func (o StorageHitachiArrayUtilization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseCapacity, errStorageBaseCapacity := json.Marshal(o.StorageBaseCapacity)
	if errStorageBaseCapacity != nil {
		return []byte{}, errStorageBaseCapacity
	}
	errStorageBaseCapacity = json.Unmarshal([]byte(serializedStorageBaseCapacity), &toSerialize)
	if errStorageBaseCapacity != nil {
		return []byte{}, errStorageBaseCapacity
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DataReduction != nil {
		toSerialize["DataReduction"] = o.DataReduction
	}
	if o.Provisioned != nil {
		toSerialize["Provisioned"] = o.Provisioned
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageHitachiArrayUtilization) UnmarshalJSON(bytes []byte) (err error) {
	type StorageHitachiArrayUtilizationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Ratio of mapped sectors within a volume versus the amount of physical space the data occupies after data compression and deduplication. The data reduction ratio does not include thin provisioning savings. For example, a data reduction ratio of 5.0 means that for every 5 MB the host writes to the array, 1 MB is stored on the array's flash modules.
		DataReduction *float32 `json:"DataReduction,omitempty"`
		// Total provisioned storage capacity in Hitachi Virtual Storage, represented in bytes.
		Provisioned *int64 `json:"Provisioned,omitempty"`
	}

	varStorageHitachiArrayUtilizationWithoutEmbeddedStruct := StorageHitachiArrayUtilizationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageHitachiArrayUtilizationWithoutEmbeddedStruct)
	if err == nil {
		varStorageHitachiArrayUtilization := _StorageHitachiArrayUtilization{}
		varStorageHitachiArrayUtilization.ClassId = varStorageHitachiArrayUtilizationWithoutEmbeddedStruct.ClassId
		varStorageHitachiArrayUtilization.ObjectType = varStorageHitachiArrayUtilizationWithoutEmbeddedStruct.ObjectType
		varStorageHitachiArrayUtilization.DataReduction = varStorageHitachiArrayUtilizationWithoutEmbeddedStruct.DataReduction
		varStorageHitachiArrayUtilization.Provisioned = varStorageHitachiArrayUtilizationWithoutEmbeddedStruct.Provisioned
		*o = StorageHitachiArrayUtilization(varStorageHitachiArrayUtilization)
	} else {
		return err
	}

	varStorageHitachiArrayUtilization := _StorageHitachiArrayUtilization{}

	err = json.Unmarshal(bytes, &varStorageHitachiArrayUtilization)
	if err == nil {
		o.StorageBaseCapacity = varStorageHitachiArrayUtilization.StorageBaseCapacity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DataReduction")
		delete(additionalProperties, "Provisioned")

		// remove fields from embedded structs
		reflectStorageBaseCapacity := reflect.ValueOf(o.StorageBaseCapacity)
		for i := 0; i < reflectStorageBaseCapacity.Type().NumField(); i++ {
			t := reflectStorageBaseCapacity.Type().Field(i)

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

type NullableStorageHitachiArrayUtilization struct {
	value *StorageHitachiArrayUtilization
	isSet bool
}

func (v NullableStorageHitachiArrayUtilization) Get() *StorageHitachiArrayUtilization {
	return v.value
}

func (v *NullableStorageHitachiArrayUtilization) Set(val *StorageHitachiArrayUtilization) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiArrayUtilization) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiArrayUtilization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiArrayUtilization(val *StorageHitachiArrayUtilization) *NullableStorageHitachiArrayUtilization {
	return &NullableStorageHitachiArrayUtilization{value: val, isSet: true}
}

func (v NullableStorageHitachiArrayUtilization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiArrayUtilization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
