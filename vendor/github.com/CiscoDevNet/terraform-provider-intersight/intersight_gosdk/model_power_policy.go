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

// PowerPolicy Power Management policy models a configuration that can be applied to Chassis or Server to manage Power Related Features.
type PowerPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Sets the Allocated Power Budget of the System (in Watts). This field is only supported for Cisco UCS X series Chassis.
	AllocatedBudget *int64 `json:"AllocatedBudget,omitempty"`
	// Sets the Power Profiling of the Server. This field is only supported for Cisco UCS X series servers. * `Enabled` - Set the value to Enabled. * `Disabled` - Set the value to Disabled.
	PowerProfiling *string `json:"PowerProfiling,omitempty"`
	// Sets the Power Restore State of the Server. This field is only supported for Cisco UCS X series servers. * `AlwaysOff` - Set the Power Restore Mode to Off. * `AlwaysOn` - Set the Power Restore Mode to On. * `LastState` - Set the Power Restore Mode to LastState.
	PowerRestoreState *string `json:"PowerRestoreState,omitempty"`
	// Sets the Power Redundancy of the System. N+2 mode is only supported for Cisco UCS X series Chassis. * `Grid` - Grid Mode requires two power sources. If one source fails, the surviving PSUs connected to the other source provides power to the chassis. * `NotRedundant` - Power Manager turns on the minimum number of PSUs required to support chassis power requirements. No Redundant PSUs are maintained. * `N+1` - Power Manager turns on the minimum number of PSUs required to support chassis power requirements plus one additional PSU for redundancy. * `N+2` - Power Manager turns on the minimum number of PSUs required to support chassis power requirements plus two additional PSU for redundancy. This Mode is only supported for UCS X series Chassis.
	RedundancyMode *string                               `json:"RedundancyMode,omitempty"`
	Organization   *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PowerPolicy PowerPolicy

// NewPowerPolicy instantiates a new PowerPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerPolicy(classId string, objectType string) *PowerPolicy {
	this := PowerPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var allocatedBudget int64 = 0
	this.AllocatedBudget = &allocatedBudget
	var powerProfiling string = "Enabled"
	this.PowerProfiling = &powerProfiling
	var powerRestoreState string = "AlwaysOff"
	this.PowerRestoreState = &powerRestoreState
	var redundancyMode string = "Grid"
	this.RedundancyMode = &redundancyMode
	return &this
}

// NewPowerPolicyWithDefaults instantiates a new PowerPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerPolicyWithDefaults() *PowerPolicy {
	this := PowerPolicy{}
	var classId string = "power.Policy"
	this.ClassId = classId
	var objectType string = "power.Policy"
	this.ObjectType = objectType
	var allocatedBudget int64 = 0
	this.AllocatedBudget = &allocatedBudget
	var powerProfiling string = "Enabled"
	this.PowerProfiling = &powerProfiling
	var powerRestoreState string = "AlwaysOff"
	this.PowerRestoreState = &powerRestoreState
	var redundancyMode string = "Grid"
	this.RedundancyMode = &redundancyMode
	return &this
}

// GetClassId returns the ClassId field value
func (o *PowerPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PowerPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PowerPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PowerPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PowerPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PowerPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAllocatedBudget returns the AllocatedBudget field value if set, zero value otherwise.
func (o *PowerPolicy) GetAllocatedBudget() int64 {
	if o == nil || o.AllocatedBudget == nil {
		var ret int64
		return ret
	}
	return *o.AllocatedBudget
}

// GetAllocatedBudgetOk returns a tuple with the AllocatedBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPolicy) GetAllocatedBudgetOk() (*int64, bool) {
	if o == nil || o.AllocatedBudget == nil {
		return nil, false
	}
	return o.AllocatedBudget, true
}

// HasAllocatedBudget returns a boolean if a field has been set.
func (o *PowerPolicy) HasAllocatedBudget() bool {
	if o != nil && o.AllocatedBudget != nil {
		return true
	}

	return false
}

// SetAllocatedBudget gets a reference to the given int64 and assigns it to the AllocatedBudget field.
func (o *PowerPolicy) SetAllocatedBudget(v int64) {
	o.AllocatedBudget = &v
}

// GetPowerProfiling returns the PowerProfiling field value if set, zero value otherwise.
func (o *PowerPolicy) GetPowerProfiling() string {
	if o == nil || o.PowerProfiling == nil {
		var ret string
		return ret
	}
	return *o.PowerProfiling
}

// GetPowerProfilingOk returns a tuple with the PowerProfiling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPolicy) GetPowerProfilingOk() (*string, bool) {
	if o == nil || o.PowerProfiling == nil {
		return nil, false
	}
	return o.PowerProfiling, true
}

// HasPowerProfiling returns a boolean if a field has been set.
func (o *PowerPolicy) HasPowerProfiling() bool {
	if o != nil && o.PowerProfiling != nil {
		return true
	}

	return false
}

// SetPowerProfiling gets a reference to the given string and assigns it to the PowerProfiling field.
func (o *PowerPolicy) SetPowerProfiling(v string) {
	o.PowerProfiling = &v
}

// GetPowerRestoreState returns the PowerRestoreState field value if set, zero value otherwise.
func (o *PowerPolicy) GetPowerRestoreState() string {
	if o == nil || o.PowerRestoreState == nil {
		var ret string
		return ret
	}
	return *o.PowerRestoreState
}

// GetPowerRestoreStateOk returns a tuple with the PowerRestoreState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPolicy) GetPowerRestoreStateOk() (*string, bool) {
	if o == nil || o.PowerRestoreState == nil {
		return nil, false
	}
	return o.PowerRestoreState, true
}

// HasPowerRestoreState returns a boolean if a field has been set.
func (o *PowerPolicy) HasPowerRestoreState() bool {
	if o != nil && o.PowerRestoreState != nil {
		return true
	}

	return false
}

// SetPowerRestoreState gets a reference to the given string and assigns it to the PowerRestoreState field.
func (o *PowerPolicy) SetPowerRestoreState(v string) {
	o.PowerRestoreState = &v
}

// GetRedundancyMode returns the RedundancyMode field value if set, zero value otherwise.
func (o *PowerPolicy) GetRedundancyMode() string {
	if o == nil || o.RedundancyMode == nil {
		var ret string
		return ret
	}
	return *o.RedundancyMode
}

// GetRedundancyModeOk returns a tuple with the RedundancyMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPolicy) GetRedundancyModeOk() (*string, bool) {
	if o == nil || o.RedundancyMode == nil {
		return nil, false
	}
	return o.RedundancyMode, true
}

// HasRedundancyMode returns a boolean if a field has been set.
func (o *PowerPolicy) HasRedundancyMode() bool {
	if o != nil && o.RedundancyMode != nil {
		return true
	}

	return false
}

// SetRedundancyMode gets a reference to the given string and assigns it to the RedundancyMode field.
func (o *PowerPolicy) SetRedundancyMode(v string) {
	o.RedundancyMode = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *PowerPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *PowerPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *PowerPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *PowerPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *PowerPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o PowerPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AllocatedBudget != nil {
		toSerialize["AllocatedBudget"] = o.AllocatedBudget
	}
	if o.PowerProfiling != nil {
		toSerialize["PowerProfiling"] = o.PowerProfiling
	}
	if o.PowerRestoreState != nil {
		toSerialize["PowerRestoreState"] = o.PowerRestoreState
	}
	if o.RedundancyMode != nil {
		toSerialize["RedundancyMode"] = o.RedundancyMode
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PowerPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type PowerPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Sets the Allocated Power Budget of the System (in Watts). This field is only supported for Cisco UCS X series Chassis.
		AllocatedBudget *int64 `json:"AllocatedBudget,omitempty"`
		// Sets the Power Profiling of the Server. This field is only supported for Cisco UCS X series servers. * `Enabled` - Set the value to Enabled. * `Disabled` - Set the value to Disabled.
		PowerProfiling *string `json:"PowerProfiling,omitempty"`
		// Sets the Power Restore State of the Server. This field is only supported for Cisco UCS X series servers. * `AlwaysOff` - Set the Power Restore Mode to Off. * `AlwaysOn` - Set the Power Restore Mode to On. * `LastState` - Set the Power Restore Mode to LastState.
		PowerRestoreState *string `json:"PowerRestoreState,omitempty"`
		// Sets the Power Redundancy of the System. N+2 mode is only supported for Cisco UCS X series Chassis. * `Grid` - Grid Mode requires two power sources. If one source fails, the surviving PSUs connected to the other source provides power to the chassis. * `NotRedundant` - Power Manager turns on the minimum number of PSUs required to support chassis power requirements. No Redundant PSUs are maintained. * `N+1` - Power Manager turns on the minimum number of PSUs required to support chassis power requirements plus one additional PSU for redundancy. * `N+2` - Power Manager turns on the minimum number of PSUs required to support chassis power requirements plus two additional PSU for redundancy. This Mode is only supported for UCS X series Chassis.
		RedundancyMode *string                               `json:"RedundancyMode,omitempty"`
		Organization   *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varPowerPolicyWithoutEmbeddedStruct := PowerPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPowerPolicyWithoutEmbeddedStruct)
	if err == nil {
		varPowerPolicy := _PowerPolicy{}
		varPowerPolicy.ClassId = varPowerPolicyWithoutEmbeddedStruct.ClassId
		varPowerPolicy.ObjectType = varPowerPolicyWithoutEmbeddedStruct.ObjectType
		varPowerPolicy.AllocatedBudget = varPowerPolicyWithoutEmbeddedStruct.AllocatedBudget
		varPowerPolicy.PowerProfiling = varPowerPolicyWithoutEmbeddedStruct.PowerProfiling
		varPowerPolicy.PowerRestoreState = varPowerPolicyWithoutEmbeddedStruct.PowerRestoreState
		varPowerPolicy.RedundancyMode = varPowerPolicyWithoutEmbeddedStruct.RedundancyMode
		varPowerPolicy.Organization = varPowerPolicyWithoutEmbeddedStruct.Organization
		varPowerPolicy.Profiles = varPowerPolicyWithoutEmbeddedStruct.Profiles
		*o = PowerPolicy(varPowerPolicy)
	} else {
		return err
	}

	varPowerPolicy := _PowerPolicy{}

	err = json.Unmarshal(bytes, &varPowerPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varPowerPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AllocatedBudget")
		delete(additionalProperties, "PowerProfiling")
		delete(additionalProperties, "PowerRestoreState")
		delete(additionalProperties, "RedundancyMode")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullablePowerPolicy struct {
	value *PowerPolicy
	isSet bool
}

func (v NullablePowerPolicy) Get() *PowerPolicy {
	return v.value
}

func (v *NullablePowerPolicy) Set(val *PowerPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerPolicy(val *PowerPolicy) *NullablePowerPolicy {
	return &NullablePowerPolicy{value: val, isSet: true}
}

func (v NullablePowerPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
