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

// PolicyAbstractConfigProfile AbstractConfigProfile is an abstract base type for all config actions on a profile.
type PolicyAbstractConfigProfile struct {
	PolicyAbstractProfile
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign.
	Action        *string                     `json:"Action,omitempty"`
	ConfigContext NullablePolicyConfigContext `json:"ConfigContext,omitempty"`
	// An array of relationships to policyAbstractPolicy resources.
	PolicyBucket         []PolicyAbstractPolicyRelationship `json:"PolicyBucket,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyAbstractConfigProfile PolicyAbstractConfigProfile

// NewPolicyAbstractConfigProfile instantiates a new PolicyAbstractConfigProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAbstractConfigProfile(classId string, objectType string) *PolicyAbstractConfigProfile {
	this := PolicyAbstractConfigProfile{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "instance"
	this.Type = &type_
	var action string = "No-op"
	this.Action = &action
	return &this
}

// NewPolicyAbstractConfigProfileWithDefaults instantiates a new PolicyAbstractConfigProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAbstractConfigProfileWithDefaults() *PolicyAbstractConfigProfile {
	this := PolicyAbstractConfigProfile{}
	var action string = "No-op"
	this.Action = &action
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyAbstractConfigProfile) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigProfile) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyAbstractConfigProfile) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PolicyAbstractConfigProfile) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigProfile) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyAbstractConfigProfile) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PolicyAbstractConfigProfile) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigProfile) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PolicyAbstractConfigProfile) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *PolicyAbstractConfigProfile) SetAction(v string) {
	o.Action = &v
}

// GetConfigContext returns the ConfigContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyAbstractConfigProfile) GetConfigContext() PolicyConfigContext {
	if o == nil || o.ConfigContext.Get() == nil {
		var ret PolicyConfigContext
		return ret
	}
	return *o.ConfigContext.Get()
}

// GetConfigContextOk returns a tuple with the ConfigContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyAbstractConfigProfile) GetConfigContextOk() (*PolicyConfigContext, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigContext.Get(), o.ConfigContext.IsSet()
}

// HasConfigContext returns a boolean if a field has been set.
func (o *PolicyAbstractConfigProfile) HasConfigContext() bool {
	if o != nil && o.ConfigContext.IsSet() {
		return true
	}

	return false
}

// SetConfigContext gets a reference to the given NullablePolicyConfigContext and assigns it to the ConfigContext field.
func (o *PolicyAbstractConfigProfile) SetConfigContext(v PolicyConfigContext) {
	o.ConfigContext.Set(&v)
}

// SetConfigContextNil sets the value for ConfigContext to be an explicit nil
func (o *PolicyAbstractConfigProfile) SetConfigContextNil() {
	o.ConfigContext.Set(nil)
}

// UnsetConfigContext ensures that no value is present for ConfigContext, not even an explicit nil
func (o *PolicyAbstractConfigProfile) UnsetConfigContext() {
	o.ConfigContext.Unset()
}

// GetPolicyBucket returns the PolicyBucket field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyAbstractConfigProfile) GetPolicyBucket() []PolicyAbstractPolicyRelationship {
	if o == nil {
		var ret []PolicyAbstractPolicyRelationship
		return ret
	}
	return o.PolicyBucket
}

// GetPolicyBucketOk returns a tuple with the PolicyBucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyAbstractConfigProfile) GetPolicyBucketOk() (*[]PolicyAbstractPolicyRelationship, bool) {
	if o == nil || o.PolicyBucket == nil {
		return nil, false
	}
	return &o.PolicyBucket, true
}

// HasPolicyBucket returns a boolean if a field has been set.
func (o *PolicyAbstractConfigProfile) HasPolicyBucket() bool {
	if o != nil && o.PolicyBucket != nil {
		return true
	}

	return false
}

// SetPolicyBucket gets a reference to the given []PolicyAbstractPolicyRelationship and assigns it to the PolicyBucket field.
func (o *PolicyAbstractConfigProfile) SetPolicyBucket(v []PolicyAbstractPolicyRelationship) {
	o.PolicyBucket = v
}

func (o PolicyAbstractConfigProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractProfile, errPolicyAbstractProfile := json.Marshal(o.PolicyAbstractProfile)
	if errPolicyAbstractProfile != nil {
		return []byte{}, errPolicyAbstractProfile
	}
	errPolicyAbstractProfile = json.Unmarshal([]byte(serializedPolicyAbstractProfile), &toSerialize)
	if errPolicyAbstractProfile != nil {
		return []byte{}, errPolicyAbstractProfile
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.ConfigContext.IsSet() {
		toSerialize["ConfigContext"] = o.ConfigContext.Get()
	}
	if o.PolicyBucket != nil {
		toSerialize["PolicyBucket"] = o.PolicyBucket
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyAbstractConfigProfile) UnmarshalJSON(bytes []byte) (err error) {
	type PolicyAbstractConfigProfileWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign.
		Action        *string                     `json:"Action,omitempty"`
		ConfigContext NullablePolicyConfigContext `json:"ConfigContext,omitempty"`
		// An array of relationships to policyAbstractPolicy resources.
		PolicyBucket []PolicyAbstractPolicyRelationship `json:"PolicyBucket,omitempty"`
	}

	varPolicyAbstractConfigProfileWithoutEmbeddedStruct := PolicyAbstractConfigProfileWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPolicyAbstractConfigProfileWithoutEmbeddedStruct)
	if err == nil {
		varPolicyAbstractConfigProfile := _PolicyAbstractConfigProfile{}
		varPolicyAbstractConfigProfile.ClassId = varPolicyAbstractConfigProfileWithoutEmbeddedStruct.ClassId
		varPolicyAbstractConfigProfile.ObjectType = varPolicyAbstractConfigProfileWithoutEmbeddedStruct.ObjectType
		varPolicyAbstractConfigProfile.Action = varPolicyAbstractConfigProfileWithoutEmbeddedStruct.Action
		varPolicyAbstractConfigProfile.ConfigContext = varPolicyAbstractConfigProfileWithoutEmbeddedStruct.ConfigContext
		varPolicyAbstractConfigProfile.PolicyBucket = varPolicyAbstractConfigProfileWithoutEmbeddedStruct.PolicyBucket
		*o = PolicyAbstractConfigProfile(varPolicyAbstractConfigProfile)
	} else {
		return err
	}

	varPolicyAbstractConfigProfile := _PolicyAbstractConfigProfile{}

	err = json.Unmarshal(bytes, &varPolicyAbstractConfigProfile)
	if err == nil {
		o.PolicyAbstractProfile = varPolicyAbstractConfigProfile.PolicyAbstractProfile
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "ConfigContext")
		delete(additionalProperties, "PolicyBucket")

		// remove fields from embedded structs
		reflectPolicyAbstractProfile := reflect.ValueOf(o.PolicyAbstractProfile)
		for i := 0; i < reflectPolicyAbstractProfile.Type().NumField(); i++ {
			t := reflectPolicyAbstractProfile.Type().Field(i)

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

type NullablePolicyAbstractConfigProfile struct {
	value *PolicyAbstractConfigProfile
	isSet bool
}

func (v NullablePolicyAbstractConfigProfile) Get() *PolicyAbstractConfigProfile {
	return v.value
}

func (v *NullablePolicyAbstractConfigProfile) Set(val *PolicyAbstractConfigProfile) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAbstractConfigProfile) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAbstractConfigProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAbstractConfigProfile(val *PolicyAbstractConfigProfile) *NullablePolicyAbstractConfigProfile {
	return &NullablePolicyAbstractConfigProfile{value: val, isSet: true}
}

func (v NullablePolicyAbstractConfigProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAbstractConfigProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
