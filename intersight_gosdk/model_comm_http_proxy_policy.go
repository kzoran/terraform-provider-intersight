/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4437
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// CommHttpProxyPolicy A policy specifying the HTTP proxy settings to be used by the HyperFlex installation process and HyperFlex storage controller VMs. This policy is required when the internet access of your servers including CIMC and HyperFlex storage controller VMs is secured by a HTTP proxy.
type CommHttpProxyPolicy struct {
	CommAbstractHttpProxyPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles      []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommHttpProxyPolicy CommHttpProxyPolicy

// NewCommHttpProxyPolicy instantiates a new CommHttpProxyPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommHttpProxyPolicy(classId string, objectType string) *CommHttpProxyPolicy {
	this := CommHttpProxyPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCommHttpProxyPolicyWithDefaults instantiates a new CommHttpProxyPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommHttpProxyPolicyWithDefaults() *CommHttpProxyPolicy {
	this := CommHttpProxyPolicy{}
	var classId string = "comm.HttpProxyPolicy"
	this.ClassId = classId
	var objectType string = "comm.HttpProxyPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CommHttpProxyPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CommHttpProxyPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CommHttpProxyPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CommHttpProxyPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CommHttpProxyPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CommHttpProxyPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommHttpProxyPolicy) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommHttpProxyPolicy) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *CommHttpProxyPolicy) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *CommHttpProxyPolicy) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

func (o CommHttpProxyPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCommAbstractHttpProxyPolicy, errCommAbstractHttpProxyPolicy := json.Marshal(o.CommAbstractHttpProxyPolicy)
	if errCommAbstractHttpProxyPolicy != nil {
		return []byte{}, errCommAbstractHttpProxyPolicy
	}
	errCommAbstractHttpProxyPolicy = json.Unmarshal([]byte(serializedCommAbstractHttpProxyPolicy), &toSerialize)
	if errCommAbstractHttpProxyPolicy != nil {
		return []byte{}, errCommAbstractHttpProxyPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CommHttpProxyPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type CommHttpProxyPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// An array of relationships to hyperflexClusterProfile resources.
		ClusterProfiles []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	}

	varCommHttpProxyPolicyWithoutEmbeddedStruct := CommHttpProxyPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCommHttpProxyPolicyWithoutEmbeddedStruct)
	if err == nil {
		varCommHttpProxyPolicy := _CommHttpProxyPolicy{}
		varCommHttpProxyPolicy.ClassId = varCommHttpProxyPolicyWithoutEmbeddedStruct.ClassId
		varCommHttpProxyPolicy.ObjectType = varCommHttpProxyPolicyWithoutEmbeddedStruct.ObjectType
		varCommHttpProxyPolicy.ClusterProfiles = varCommHttpProxyPolicyWithoutEmbeddedStruct.ClusterProfiles
		*o = CommHttpProxyPolicy(varCommHttpProxyPolicy)
	} else {
		return err
	}

	varCommHttpProxyPolicy := _CommHttpProxyPolicy{}

	err = json.Unmarshal(bytes, &varCommHttpProxyPolicy)
	if err == nil {
		o.CommAbstractHttpProxyPolicy = varCommHttpProxyPolicy.CommAbstractHttpProxyPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterProfiles")

		// remove fields from embedded structs
		reflectCommAbstractHttpProxyPolicy := reflect.ValueOf(o.CommAbstractHttpProxyPolicy)
		for i := 0; i < reflectCommAbstractHttpProxyPolicy.Type().NumField(); i++ {
			t := reflectCommAbstractHttpProxyPolicy.Type().Field(i)

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

type NullableCommHttpProxyPolicy struct {
	value *CommHttpProxyPolicy
	isSet bool
}

func (v NullableCommHttpProxyPolicy) Get() *CommHttpProxyPolicy {
	return v.value
}

func (v *NullableCommHttpProxyPolicy) Set(val *CommHttpProxyPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableCommHttpProxyPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableCommHttpProxyPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommHttpProxyPolicy(val *CommHttpProxyPolicy) *NullableCommHttpProxyPolicy {
	return &NullableCommHttpProxyPolicy{value: val, isSet: true}
}

func (v NullableCommHttpProxyPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommHttpProxyPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
