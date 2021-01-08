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

// FabricPcMember PcMember object is to establish the relationship between port parameters and pcId.
type FabricPcMember struct {
	FabricPortBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Port Channel Identifier for the collection of ports.
	PcId                 *int64                        `json:"PcId,omitempty"`
	PortPolicy           *FabricPortPolicyRelationship `json:"PortPolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricPcMember FabricPcMember

// NewFabricPcMember instantiates a new FabricPcMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricPcMember(classId string, objectType string) *FabricPcMember {
	this := FabricPcMember{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricPcMemberWithDefaults instantiates a new FabricPcMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricPcMemberWithDefaults() *FabricPcMember {
	this := FabricPcMember{}
	var classId string = "fabric.PcMember"
	this.ClassId = classId
	var objectType string = "fabric.PcMember"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricPcMember) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricPcMember) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricPcMember) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricPcMember) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricPcMember) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricPcMember) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPcId returns the PcId field value if set, zero value otherwise.
func (o *FabricPcMember) GetPcId() int64 {
	if o == nil || o.PcId == nil {
		var ret int64
		return ret
	}
	return *o.PcId
}

// GetPcIdOk returns a tuple with the PcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPcMember) GetPcIdOk() (*int64, bool) {
	if o == nil || o.PcId == nil {
		return nil, false
	}
	return o.PcId, true
}

// HasPcId returns a boolean if a field has been set.
func (o *FabricPcMember) HasPcId() bool {
	if o != nil && o.PcId != nil {
		return true
	}

	return false
}

// SetPcId gets a reference to the given int64 and assigns it to the PcId field.
func (o *FabricPcMember) SetPcId(v int64) {
	o.PcId = &v
}

// GetPortPolicy returns the PortPolicy field value if set, zero value otherwise.
func (o *FabricPcMember) GetPortPolicy() FabricPortPolicyRelationship {
	if o == nil || o.PortPolicy == nil {
		var ret FabricPortPolicyRelationship
		return ret
	}
	return *o.PortPolicy
}

// GetPortPolicyOk returns a tuple with the PortPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPcMember) GetPortPolicyOk() (*FabricPortPolicyRelationship, bool) {
	if o == nil || o.PortPolicy == nil {
		return nil, false
	}
	return o.PortPolicy, true
}

// HasPortPolicy returns a boolean if a field has been set.
func (o *FabricPcMember) HasPortPolicy() bool {
	if o != nil && o.PortPolicy != nil {
		return true
	}

	return false
}

// SetPortPolicy gets a reference to the given FabricPortPolicyRelationship and assigns it to the PortPolicy field.
func (o *FabricPcMember) SetPortPolicy(v FabricPortPolicyRelationship) {
	o.PortPolicy = &v
}

func (o FabricPcMember) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricPortBase, errFabricPortBase := json.Marshal(o.FabricPortBase)
	if errFabricPortBase != nil {
		return []byte{}, errFabricPortBase
	}
	errFabricPortBase = json.Unmarshal([]byte(serializedFabricPortBase), &toSerialize)
	if errFabricPortBase != nil {
		return []byte{}, errFabricPortBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.PcId != nil {
		toSerialize["PcId"] = o.PcId
	}
	if o.PortPolicy != nil {
		toSerialize["PortPolicy"] = o.PortPolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricPcMember) UnmarshalJSON(bytes []byte) (err error) {
	type FabricPcMemberWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Port Channel Identifier for the collection of ports.
		PcId       *int64                        `json:"PcId,omitempty"`
		PortPolicy *FabricPortPolicyRelationship `json:"PortPolicy,omitempty"`
	}

	varFabricPcMemberWithoutEmbeddedStruct := FabricPcMemberWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricPcMemberWithoutEmbeddedStruct)
	if err == nil {
		varFabricPcMember := _FabricPcMember{}
		varFabricPcMember.ClassId = varFabricPcMemberWithoutEmbeddedStruct.ClassId
		varFabricPcMember.ObjectType = varFabricPcMemberWithoutEmbeddedStruct.ObjectType
		varFabricPcMember.PcId = varFabricPcMemberWithoutEmbeddedStruct.PcId
		varFabricPcMember.PortPolicy = varFabricPcMemberWithoutEmbeddedStruct.PortPolicy
		*o = FabricPcMember(varFabricPcMember)
	} else {
		return err
	}

	varFabricPcMember := _FabricPcMember{}

	err = json.Unmarshal(bytes, &varFabricPcMember)
	if err == nil {
		o.FabricPortBase = varFabricPcMember.FabricPortBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PcId")
		delete(additionalProperties, "PortPolicy")

		// remove fields from embedded structs
		reflectFabricPortBase := reflect.ValueOf(o.FabricPortBase)
		for i := 0; i < reflectFabricPortBase.Type().NumField(); i++ {
			t := reflectFabricPortBase.Type().Field(i)

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

type NullableFabricPcMember struct {
	value *FabricPcMember
	isSet bool
}

func (v NullableFabricPcMember) Get() *FabricPcMember {
	return v.value
}

func (v *NullableFabricPcMember) Set(val *FabricPcMember) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricPcMember) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricPcMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricPcMember(val *FabricPcMember) *NullableFabricPcMember {
	return &NullableFabricPcMember{value: val, isSet: true}
}

func (v NullableFabricPcMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricPcMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
