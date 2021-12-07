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

// ComputePersistentMemoryModule Persistent Memory Module on which an operation has to be performed.
type ComputePersistentMemoryModule struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Socket ID of the Persistent Memory Module on the server.
	SocketId *string `json:"SocketId,omitempty"`
	// Socket Memory ID of the Persistent Memory Module on the server.
	SocketMemoryId       *string `json:"SocketMemoryId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputePersistentMemoryModule ComputePersistentMemoryModule

// NewComputePersistentMemoryModule instantiates a new ComputePersistentMemoryModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputePersistentMemoryModule(classId string, objectType string) *ComputePersistentMemoryModule {
	this := ComputePersistentMemoryModule{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewComputePersistentMemoryModuleWithDefaults instantiates a new ComputePersistentMemoryModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputePersistentMemoryModuleWithDefaults() *ComputePersistentMemoryModule {
	this := ComputePersistentMemoryModule{}
	var classId string = "compute.PersistentMemoryModule"
	this.ClassId = classId
	var objectType string = "compute.PersistentMemoryModule"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputePersistentMemoryModule) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputePersistentMemoryModule) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputePersistentMemoryModule) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputePersistentMemoryModule) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputePersistentMemoryModule) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputePersistentMemoryModule) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSocketId returns the SocketId field value if set, zero value otherwise.
func (o *ComputePersistentMemoryModule) GetSocketId() string {
	if o == nil || o.SocketId == nil {
		var ret string
		return ret
	}
	return *o.SocketId
}

// GetSocketIdOk returns a tuple with the SocketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePersistentMemoryModule) GetSocketIdOk() (*string, bool) {
	if o == nil || o.SocketId == nil {
		return nil, false
	}
	return o.SocketId, true
}

// HasSocketId returns a boolean if a field has been set.
func (o *ComputePersistentMemoryModule) HasSocketId() bool {
	if o != nil && o.SocketId != nil {
		return true
	}

	return false
}

// SetSocketId gets a reference to the given string and assigns it to the SocketId field.
func (o *ComputePersistentMemoryModule) SetSocketId(v string) {
	o.SocketId = &v
}

// GetSocketMemoryId returns the SocketMemoryId field value if set, zero value otherwise.
func (o *ComputePersistentMemoryModule) GetSocketMemoryId() string {
	if o == nil || o.SocketMemoryId == nil {
		var ret string
		return ret
	}
	return *o.SocketMemoryId
}

// GetSocketMemoryIdOk returns a tuple with the SocketMemoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePersistentMemoryModule) GetSocketMemoryIdOk() (*string, bool) {
	if o == nil || o.SocketMemoryId == nil {
		return nil, false
	}
	return o.SocketMemoryId, true
}

// HasSocketMemoryId returns a boolean if a field has been set.
func (o *ComputePersistentMemoryModule) HasSocketMemoryId() bool {
	if o != nil && o.SocketMemoryId != nil {
		return true
	}

	return false
}

// SetSocketMemoryId gets a reference to the given string and assigns it to the SocketMemoryId field.
func (o *ComputePersistentMemoryModule) SetSocketMemoryId(v string) {
	o.SocketMemoryId = &v
}

func (o ComputePersistentMemoryModule) MarshalJSON() ([]byte, error) {
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
	if o.SocketId != nil {
		toSerialize["SocketId"] = o.SocketId
	}
	if o.SocketMemoryId != nil {
		toSerialize["SocketMemoryId"] = o.SocketMemoryId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputePersistentMemoryModule) UnmarshalJSON(bytes []byte) (err error) {
	type ComputePersistentMemoryModuleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Socket ID of the Persistent Memory Module on the server.
		SocketId *string `json:"SocketId,omitempty"`
		// Socket Memory ID of the Persistent Memory Module on the server.
		SocketMemoryId *string `json:"SocketMemoryId,omitempty"`
	}

	varComputePersistentMemoryModuleWithoutEmbeddedStruct := ComputePersistentMemoryModuleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varComputePersistentMemoryModuleWithoutEmbeddedStruct)
	if err == nil {
		varComputePersistentMemoryModule := _ComputePersistentMemoryModule{}
		varComputePersistentMemoryModule.ClassId = varComputePersistentMemoryModuleWithoutEmbeddedStruct.ClassId
		varComputePersistentMemoryModule.ObjectType = varComputePersistentMemoryModuleWithoutEmbeddedStruct.ObjectType
		varComputePersistentMemoryModule.SocketId = varComputePersistentMemoryModuleWithoutEmbeddedStruct.SocketId
		varComputePersistentMemoryModule.SocketMemoryId = varComputePersistentMemoryModuleWithoutEmbeddedStruct.SocketMemoryId
		*o = ComputePersistentMemoryModule(varComputePersistentMemoryModule)
	} else {
		return err
	}

	varComputePersistentMemoryModule := _ComputePersistentMemoryModule{}

	err = json.Unmarshal(bytes, &varComputePersistentMemoryModule)
	if err == nil {
		o.MoBaseComplexType = varComputePersistentMemoryModule.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SocketId")
		delete(additionalProperties, "SocketMemoryId")

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

type NullableComputePersistentMemoryModule struct {
	value *ComputePersistentMemoryModule
	isSet bool
}

func (v NullableComputePersistentMemoryModule) Get() *ComputePersistentMemoryModule {
	return v.value
}

func (v *NullableComputePersistentMemoryModule) Set(val *ComputePersistentMemoryModule) {
	v.value = val
	v.isSet = true
}

func (v NullableComputePersistentMemoryModule) IsSet() bool {
	return v.isSet
}

func (v *NullableComputePersistentMemoryModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputePersistentMemoryModule(val *ComputePersistentMemoryModule) *NullableComputePersistentMemoryModule {
	return &NullableComputePersistentMemoryModule{value: val, isSet: true}
}

func (v NullableComputePersistentMemoryModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputePersistentMemoryModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
