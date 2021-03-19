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

// HyperflexFilePath File path information for this snapshot.
type HyperflexFilePath struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                         `json:"ObjectType"`
	DsInfo     NullableHyperflexDatastoreInfo `json:"DsInfo,omitempty"`
	// Relative file path within the datastore.
	RelativeFilePath     *string `json:"RelativeFilePath,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexFilePath HyperflexFilePath

// NewHyperflexFilePath instantiates a new HyperflexFilePath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexFilePath(classId string, objectType string) *HyperflexFilePath {
	this := HyperflexFilePath{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexFilePathWithDefaults instantiates a new HyperflexFilePath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexFilePathWithDefaults() *HyperflexFilePath {
	this := HyperflexFilePath{}
	var classId string = "hyperflex.FilePath"
	this.ClassId = classId
	var objectType string = "hyperflex.FilePath"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexFilePath) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexFilePath) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexFilePath) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexFilePath) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexFilePath) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexFilePath) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDsInfo returns the DsInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexFilePath) GetDsInfo() HyperflexDatastoreInfo {
	if o == nil || o.DsInfo.Get() == nil {
		var ret HyperflexDatastoreInfo
		return ret
	}
	return *o.DsInfo.Get()
}

// GetDsInfoOk returns a tuple with the DsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexFilePath) GetDsInfoOk() (*HyperflexDatastoreInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.DsInfo.Get(), o.DsInfo.IsSet()
}

// HasDsInfo returns a boolean if a field has been set.
func (o *HyperflexFilePath) HasDsInfo() bool {
	if o != nil && o.DsInfo.IsSet() {
		return true
	}

	return false
}

// SetDsInfo gets a reference to the given NullableHyperflexDatastoreInfo and assigns it to the DsInfo field.
func (o *HyperflexFilePath) SetDsInfo(v HyperflexDatastoreInfo) {
	o.DsInfo.Set(&v)
}

// SetDsInfoNil sets the value for DsInfo to be an explicit nil
func (o *HyperflexFilePath) SetDsInfoNil() {
	o.DsInfo.Set(nil)
}

// UnsetDsInfo ensures that no value is present for DsInfo, not even an explicit nil
func (o *HyperflexFilePath) UnsetDsInfo() {
	o.DsInfo.Unset()
}

// GetRelativeFilePath returns the RelativeFilePath field value if set, zero value otherwise.
func (o *HyperflexFilePath) GetRelativeFilePath() string {
	if o == nil || o.RelativeFilePath == nil {
		var ret string
		return ret
	}
	return *o.RelativeFilePath
}

// GetRelativeFilePathOk returns a tuple with the RelativeFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexFilePath) GetRelativeFilePathOk() (*string, bool) {
	if o == nil || o.RelativeFilePath == nil {
		return nil, false
	}
	return o.RelativeFilePath, true
}

// HasRelativeFilePath returns a boolean if a field has been set.
func (o *HyperflexFilePath) HasRelativeFilePath() bool {
	if o != nil && o.RelativeFilePath != nil {
		return true
	}

	return false
}

// SetRelativeFilePath gets a reference to the given string and assigns it to the RelativeFilePath field.
func (o *HyperflexFilePath) SetRelativeFilePath(v string) {
	o.RelativeFilePath = &v
}

func (o HyperflexFilePath) MarshalJSON() ([]byte, error) {
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
	if o.DsInfo.IsSet() {
		toSerialize["DsInfo"] = o.DsInfo.Get()
	}
	if o.RelativeFilePath != nil {
		toSerialize["RelativeFilePath"] = o.RelativeFilePath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexFilePath) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexFilePathWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                         `json:"ObjectType"`
		DsInfo     NullableHyperflexDatastoreInfo `json:"DsInfo,omitempty"`
		// Relative file path within the datastore.
		RelativeFilePath *string `json:"RelativeFilePath,omitempty"`
	}

	varHyperflexFilePathWithoutEmbeddedStruct := HyperflexFilePathWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexFilePathWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexFilePath := _HyperflexFilePath{}
		varHyperflexFilePath.ClassId = varHyperflexFilePathWithoutEmbeddedStruct.ClassId
		varHyperflexFilePath.ObjectType = varHyperflexFilePathWithoutEmbeddedStruct.ObjectType
		varHyperflexFilePath.DsInfo = varHyperflexFilePathWithoutEmbeddedStruct.DsInfo
		varHyperflexFilePath.RelativeFilePath = varHyperflexFilePathWithoutEmbeddedStruct.RelativeFilePath
		*o = HyperflexFilePath(varHyperflexFilePath)
	} else {
		return err
	}

	varHyperflexFilePath := _HyperflexFilePath{}

	err = json.Unmarshal(bytes, &varHyperflexFilePath)
	if err == nil {
		o.MoBaseComplexType = varHyperflexFilePath.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DsInfo")
		delete(additionalProperties, "RelativeFilePath")

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

type NullableHyperflexFilePath struct {
	value *HyperflexFilePath
	isSet bool
}

func (v NullableHyperflexFilePath) Get() *HyperflexFilePath {
	return v.value
}

func (v *NullableHyperflexFilePath) Set(val *HyperflexFilePath) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexFilePath) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexFilePath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexFilePath(val *HyperflexFilePath) *NullableHyperflexFilePath {
	return &NullableHyperflexFilePath{value: val, isSet: true}
}

func (v NullableHyperflexFilePath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexFilePath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
