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
)

// FabricSwitchProfileAllOf Definition of the list of properties defined in 'fabric.SwitchProfile', excluding properties defined in parent classes.
type FabricSwitchProfileAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType       string                      `json:"ObjectType"`
	ConfigChanges    NullablePolicyConfigChange  `json:"ConfigChanges,omitempty"`
	AssignedSwitch   *NetworkElementRelationship `json:"AssignedSwitch,omitempty"`
	AssociatedSwitch *NetworkElementRelationship `json:"AssociatedSwitch,omitempty"`
	// An array of relationships to fabricConfigChangeDetail resources.
	ConfigChangeDetails []FabricConfigChangeDetailRelationship `json:"ConfigChangeDetails,omitempty"`
	ConfigResult        *FabricConfigResultRelationship        `json:"ConfigResult,omitempty"`
	// An array of relationships to workflowWorkflowInfo resources.
	RunningWorkflows     []WorkflowWorkflowInfoRelationship      `json:"RunningWorkflows,omitempty"`
	SwitchClusterProfile *FabricSwitchClusterProfileRelationship `json:"SwitchClusterProfile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricSwitchProfileAllOf FabricSwitchProfileAllOf

// NewFabricSwitchProfileAllOf instantiates a new FabricSwitchProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricSwitchProfileAllOf(classId string, objectType string) *FabricSwitchProfileAllOf {
	this := FabricSwitchProfileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricSwitchProfileAllOfWithDefaults instantiates a new FabricSwitchProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricSwitchProfileAllOfWithDefaults() *FabricSwitchProfileAllOf {
	this := FabricSwitchProfileAllOf{}
	var classId string = "fabric.SwitchProfile"
	this.ClassId = classId
	var objectType string = "fabric.SwitchProfile"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricSwitchProfileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricSwitchProfileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricSwitchProfileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricSwitchProfileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricSwitchProfileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricSwitchProfileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigChanges returns the ConfigChanges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSwitchProfileAllOf) GetConfigChanges() PolicyConfigChange {
	if o == nil || o.ConfigChanges.Get() == nil {
		var ret PolicyConfigChange
		return ret
	}
	return *o.ConfigChanges.Get()
}

// GetConfigChangesOk returns a tuple with the ConfigChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSwitchProfileAllOf) GetConfigChangesOk() (*PolicyConfigChange, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigChanges.Get(), o.ConfigChanges.IsSet()
}

// HasConfigChanges returns a boolean if a field has been set.
func (o *FabricSwitchProfileAllOf) HasConfigChanges() bool {
	if o != nil && o.ConfigChanges.IsSet() {
		return true
	}

	return false
}

// SetConfigChanges gets a reference to the given NullablePolicyConfigChange and assigns it to the ConfigChanges field.
func (o *FabricSwitchProfileAllOf) SetConfigChanges(v PolicyConfigChange) {
	o.ConfigChanges.Set(&v)
}

// SetConfigChangesNil sets the value for ConfigChanges to be an explicit nil
func (o *FabricSwitchProfileAllOf) SetConfigChangesNil() {
	o.ConfigChanges.Set(nil)
}

// UnsetConfigChanges ensures that no value is present for ConfigChanges, not even an explicit nil
func (o *FabricSwitchProfileAllOf) UnsetConfigChanges() {
	o.ConfigChanges.Unset()
}

// GetAssignedSwitch returns the AssignedSwitch field value if set, zero value otherwise.
func (o *FabricSwitchProfileAllOf) GetAssignedSwitch() NetworkElementRelationship {
	if o == nil || o.AssignedSwitch == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.AssignedSwitch
}

// GetAssignedSwitchOk returns a tuple with the AssignedSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSwitchProfileAllOf) GetAssignedSwitchOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.AssignedSwitch == nil {
		return nil, false
	}
	return o.AssignedSwitch, true
}

// HasAssignedSwitch returns a boolean if a field has been set.
func (o *FabricSwitchProfileAllOf) HasAssignedSwitch() bool {
	if o != nil && o.AssignedSwitch != nil {
		return true
	}

	return false
}

// SetAssignedSwitch gets a reference to the given NetworkElementRelationship and assigns it to the AssignedSwitch field.
func (o *FabricSwitchProfileAllOf) SetAssignedSwitch(v NetworkElementRelationship) {
	o.AssignedSwitch = &v
}

// GetAssociatedSwitch returns the AssociatedSwitch field value if set, zero value otherwise.
func (o *FabricSwitchProfileAllOf) GetAssociatedSwitch() NetworkElementRelationship {
	if o == nil || o.AssociatedSwitch == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.AssociatedSwitch
}

// GetAssociatedSwitchOk returns a tuple with the AssociatedSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSwitchProfileAllOf) GetAssociatedSwitchOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.AssociatedSwitch == nil {
		return nil, false
	}
	return o.AssociatedSwitch, true
}

// HasAssociatedSwitch returns a boolean if a field has been set.
func (o *FabricSwitchProfileAllOf) HasAssociatedSwitch() bool {
	if o != nil && o.AssociatedSwitch != nil {
		return true
	}

	return false
}

// SetAssociatedSwitch gets a reference to the given NetworkElementRelationship and assigns it to the AssociatedSwitch field.
func (o *FabricSwitchProfileAllOf) SetAssociatedSwitch(v NetworkElementRelationship) {
	o.AssociatedSwitch = &v
}

// GetConfigChangeDetails returns the ConfigChangeDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSwitchProfileAllOf) GetConfigChangeDetails() []FabricConfigChangeDetailRelationship {
	if o == nil {
		var ret []FabricConfigChangeDetailRelationship
		return ret
	}
	return o.ConfigChangeDetails
}

// GetConfigChangeDetailsOk returns a tuple with the ConfigChangeDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSwitchProfileAllOf) GetConfigChangeDetailsOk() (*[]FabricConfigChangeDetailRelationship, bool) {
	if o == nil || o.ConfigChangeDetails == nil {
		return nil, false
	}
	return &o.ConfigChangeDetails, true
}

// HasConfigChangeDetails returns a boolean if a field has been set.
func (o *FabricSwitchProfileAllOf) HasConfigChangeDetails() bool {
	if o != nil && o.ConfigChangeDetails != nil {
		return true
	}

	return false
}

// SetConfigChangeDetails gets a reference to the given []FabricConfigChangeDetailRelationship and assigns it to the ConfigChangeDetails field.
func (o *FabricSwitchProfileAllOf) SetConfigChangeDetails(v []FabricConfigChangeDetailRelationship) {
	o.ConfigChangeDetails = v
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise.
func (o *FabricSwitchProfileAllOf) GetConfigResult() FabricConfigResultRelationship {
	if o == nil || o.ConfigResult == nil {
		var ret FabricConfigResultRelationship
		return ret
	}
	return *o.ConfigResult
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSwitchProfileAllOf) GetConfigResultOk() (*FabricConfigResultRelationship, bool) {
	if o == nil || o.ConfigResult == nil {
		return nil, false
	}
	return o.ConfigResult, true
}

// HasConfigResult returns a boolean if a field has been set.
func (o *FabricSwitchProfileAllOf) HasConfigResult() bool {
	if o != nil && o.ConfigResult != nil {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given FabricConfigResultRelationship and assigns it to the ConfigResult field.
func (o *FabricSwitchProfileAllOf) SetConfigResult(v FabricConfigResultRelationship) {
	o.ConfigResult = &v
}

// GetRunningWorkflows returns the RunningWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSwitchProfileAllOf) GetRunningWorkflows() []WorkflowWorkflowInfoRelationship {
	if o == nil {
		var ret []WorkflowWorkflowInfoRelationship
		return ret
	}
	return o.RunningWorkflows
}

// GetRunningWorkflowsOk returns a tuple with the RunningWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSwitchProfileAllOf) GetRunningWorkflowsOk() (*[]WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.RunningWorkflows == nil {
		return nil, false
	}
	return &o.RunningWorkflows, true
}

// HasRunningWorkflows returns a boolean if a field has been set.
func (o *FabricSwitchProfileAllOf) HasRunningWorkflows() bool {
	if o != nil && o.RunningWorkflows != nil {
		return true
	}

	return false
}

// SetRunningWorkflows gets a reference to the given []WorkflowWorkflowInfoRelationship and assigns it to the RunningWorkflows field.
func (o *FabricSwitchProfileAllOf) SetRunningWorkflows(v []WorkflowWorkflowInfoRelationship) {
	o.RunningWorkflows = v
}

// GetSwitchClusterProfile returns the SwitchClusterProfile field value if set, zero value otherwise.
func (o *FabricSwitchProfileAllOf) GetSwitchClusterProfile() FabricSwitchClusterProfileRelationship {
	if o == nil || o.SwitchClusterProfile == nil {
		var ret FabricSwitchClusterProfileRelationship
		return ret
	}
	return *o.SwitchClusterProfile
}

// GetSwitchClusterProfileOk returns a tuple with the SwitchClusterProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSwitchProfileAllOf) GetSwitchClusterProfileOk() (*FabricSwitchClusterProfileRelationship, bool) {
	if o == nil || o.SwitchClusterProfile == nil {
		return nil, false
	}
	return o.SwitchClusterProfile, true
}

// HasSwitchClusterProfile returns a boolean if a field has been set.
func (o *FabricSwitchProfileAllOf) HasSwitchClusterProfile() bool {
	if o != nil && o.SwitchClusterProfile != nil {
		return true
	}

	return false
}

// SetSwitchClusterProfile gets a reference to the given FabricSwitchClusterProfileRelationship and assigns it to the SwitchClusterProfile field.
func (o *FabricSwitchProfileAllOf) SetSwitchClusterProfile(v FabricSwitchClusterProfileRelationship) {
	o.SwitchClusterProfile = &v
}

func (o FabricSwitchProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigChanges.IsSet() {
		toSerialize["ConfigChanges"] = o.ConfigChanges.Get()
	}
	if o.AssignedSwitch != nil {
		toSerialize["AssignedSwitch"] = o.AssignedSwitch
	}
	if o.AssociatedSwitch != nil {
		toSerialize["AssociatedSwitch"] = o.AssociatedSwitch
	}
	if o.ConfigChangeDetails != nil {
		toSerialize["ConfigChangeDetails"] = o.ConfigChangeDetails
	}
	if o.ConfigResult != nil {
		toSerialize["ConfigResult"] = o.ConfigResult
	}
	if o.RunningWorkflows != nil {
		toSerialize["RunningWorkflows"] = o.RunningWorkflows
	}
	if o.SwitchClusterProfile != nil {
		toSerialize["SwitchClusterProfile"] = o.SwitchClusterProfile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricSwitchProfileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricSwitchProfileAllOf := _FabricSwitchProfileAllOf{}

	if err = json.Unmarshal(bytes, &varFabricSwitchProfileAllOf); err == nil {
		*o = FabricSwitchProfileAllOf(varFabricSwitchProfileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigChanges")
		delete(additionalProperties, "AssignedSwitch")
		delete(additionalProperties, "AssociatedSwitch")
		delete(additionalProperties, "ConfigChangeDetails")
		delete(additionalProperties, "ConfigResult")
		delete(additionalProperties, "RunningWorkflows")
		delete(additionalProperties, "SwitchClusterProfile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricSwitchProfileAllOf struct {
	value *FabricSwitchProfileAllOf
	isSet bool
}

func (v NullableFabricSwitchProfileAllOf) Get() *FabricSwitchProfileAllOf {
	return v.value
}

func (v *NullableFabricSwitchProfileAllOf) Set(val *FabricSwitchProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricSwitchProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricSwitchProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricSwitchProfileAllOf(val *FabricSwitchProfileAllOf) *NullableFabricSwitchProfileAllOf {
	return &NullableFabricSwitchProfileAllOf{value: val, isSet: true}
}

func (v NullableFabricSwitchProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricSwitchProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
