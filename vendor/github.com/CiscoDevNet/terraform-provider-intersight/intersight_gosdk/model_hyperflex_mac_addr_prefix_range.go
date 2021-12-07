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

// HyperflexMacAddrPrefixRange A MAC address prefix range. The range is inclusive and comprised of a start and end MAC addresses. A single address can be specified by setting it as the start and end of the range.
type HyperflexMacAddrPrefixRange struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The end MAC address prefix of a MAC address prefix range in the form of 00:25:B5:XX.
	EndAddr *string `json:"EndAddr,omitempty"`
	// The start MAC address prefix of a MAC address prefix range in the form of 00:25:B5:XX.
	StartAddr            *string `json:"StartAddr,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexMacAddrPrefixRange HyperflexMacAddrPrefixRange

// NewHyperflexMacAddrPrefixRange instantiates a new HyperflexMacAddrPrefixRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexMacAddrPrefixRange(classId string, objectType string) *HyperflexMacAddrPrefixRange {
	this := HyperflexMacAddrPrefixRange{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexMacAddrPrefixRangeWithDefaults instantiates a new HyperflexMacAddrPrefixRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexMacAddrPrefixRangeWithDefaults() *HyperflexMacAddrPrefixRange {
	this := HyperflexMacAddrPrefixRange{}
	var classId string = "hyperflex.MacAddrPrefixRange"
	this.ClassId = classId
	var objectType string = "hyperflex.MacAddrPrefixRange"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexMacAddrPrefixRange) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexMacAddrPrefixRange) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexMacAddrPrefixRange) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexMacAddrPrefixRange) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexMacAddrPrefixRange) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexMacAddrPrefixRange) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEndAddr returns the EndAddr field value if set, zero value otherwise.
func (o *HyperflexMacAddrPrefixRange) GetEndAddr() string {
	if o == nil || o.EndAddr == nil {
		var ret string
		return ret
	}
	return *o.EndAddr
}

// GetEndAddrOk returns a tuple with the EndAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexMacAddrPrefixRange) GetEndAddrOk() (*string, bool) {
	if o == nil || o.EndAddr == nil {
		return nil, false
	}
	return o.EndAddr, true
}

// HasEndAddr returns a boolean if a field has been set.
func (o *HyperflexMacAddrPrefixRange) HasEndAddr() bool {
	if o != nil && o.EndAddr != nil {
		return true
	}

	return false
}

// SetEndAddr gets a reference to the given string and assigns it to the EndAddr field.
func (o *HyperflexMacAddrPrefixRange) SetEndAddr(v string) {
	o.EndAddr = &v
}

// GetStartAddr returns the StartAddr field value if set, zero value otherwise.
func (o *HyperflexMacAddrPrefixRange) GetStartAddr() string {
	if o == nil || o.StartAddr == nil {
		var ret string
		return ret
	}
	return *o.StartAddr
}

// GetStartAddrOk returns a tuple with the StartAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexMacAddrPrefixRange) GetStartAddrOk() (*string, bool) {
	if o == nil || o.StartAddr == nil {
		return nil, false
	}
	return o.StartAddr, true
}

// HasStartAddr returns a boolean if a field has been set.
func (o *HyperflexMacAddrPrefixRange) HasStartAddr() bool {
	if o != nil && o.StartAddr != nil {
		return true
	}

	return false
}

// SetStartAddr gets a reference to the given string and assigns it to the StartAddr field.
func (o *HyperflexMacAddrPrefixRange) SetStartAddr(v string) {
	o.StartAddr = &v
}

func (o HyperflexMacAddrPrefixRange) MarshalJSON() ([]byte, error) {
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
	if o.EndAddr != nil {
		toSerialize["EndAddr"] = o.EndAddr
	}
	if o.StartAddr != nil {
		toSerialize["StartAddr"] = o.StartAddr
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexMacAddrPrefixRange) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexMacAddrPrefixRangeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The end MAC address prefix of a MAC address prefix range in the form of 00:25:B5:XX.
		EndAddr *string `json:"EndAddr,omitempty"`
		// The start MAC address prefix of a MAC address prefix range in the form of 00:25:B5:XX.
		StartAddr *string `json:"StartAddr,omitempty"`
	}

	varHyperflexMacAddrPrefixRangeWithoutEmbeddedStruct := HyperflexMacAddrPrefixRangeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexMacAddrPrefixRangeWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexMacAddrPrefixRange := _HyperflexMacAddrPrefixRange{}
		varHyperflexMacAddrPrefixRange.ClassId = varHyperflexMacAddrPrefixRangeWithoutEmbeddedStruct.ClassId
		varHyperflexMacAddrPrefixRange.ObjectType = varHyperflexMacAddrPrefixRangeWithoutEmbeddedStruct.ObjectType
		varHyperflexMacAddrPrefixRange.EndAddr = varHyperflexMacAddrPrefixRangeWithoutEmbeddedStruct.EndAddr
		varHyperflexMacAddrPrefixRange.StartAddr = varHyperflexMacAddrPrefixRangeWithoutEmbeddedStruct.StartAddr
		*o = HyperflexMacAddrPrefixRange(varHyperflexMacAddrPrefixRange)
	} else {
		return err
	}

	varHyperflexMacAddrPrefixRange := _HyperflexMacAddrPrefixRange{}

	err = json.Unmarshal(bytes, &varHyperflexMacAddrPrefixRange)
	if err == nil {
		o.MoBaseComplexType = varHyperflexMacAddrPrefixRange.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EndAddr")
		delete(additionalProperties, "StartAddr")

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

type NullableHyperflexMacAddrPrefixRange struct {
	value *HyperflexMacAddrPrefixRange
	isSet bool
}

func (v NullableHyperflexMacAddrPrefixRange) Get() *HyperflexMacAddrPrefixRange {
	return v.value
}

func (v *NullableHyperflexMacAddrPrefixRange) Set(val *HyperflexMacAddrPrefixRange) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexMacAddrPrefixRange) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexMacAddrPrefixRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexMacAddrPrefixRange(val *HyperflexMacAddrPrefixRange) *NullableHyperflexMacAddrPrefixRange {
	return &NullableHyperflexMacAddrPrefixRange{value: val, isSet: true}
}

func (v NullableHyperflexMacAddrPrefixRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexMacAddrPrefixRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
