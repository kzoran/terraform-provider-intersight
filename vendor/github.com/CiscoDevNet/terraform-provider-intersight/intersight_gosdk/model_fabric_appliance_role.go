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

// FabricApplianceRole Configuration object sent by user to create an appliance port.
type FabricApplianceRole struct {
	FabricTransceiverRole
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Port mode to be set on the appliance port. * `trunk` - Trunk Mode Switch Port Type. * `access` - Access Mode Switch Port Type.
	Mode *string `json:"Mode,omitempty"`
	// The 'name' of the System QoS Class. * `Best Effort` - QoS Priority for Best-effort traffic. * `FC` - QoS Priority for FC traffic. * `Platinum` - QoS Priority for Platinum traffic. * `Gold` - QoS Priority for Gold traffic. * `Silver` - QoS Priority for Silver traffic. * `Bronze` - QoS Priority for Bronze traffic.
	Priority                *string                                    `json:"Priority,omitempty"`
	EthNetworkControlPolicy *FabricEthNetworkControlPolicyRelationship `json:"EthNetworkControlPolicy,omitempty"`
	EthNetworkGroupPolicy   *FabricEthNetworkGroupPolicyRelationship   `json:"EthNetworkGroupPolicy,omitempty"`
	FlowControlPolicy       *FabricFlowControlPolicyRelationship       `json:"FlowControlPolicy,omitempty"`
	LinkControlPolicy       *FabricLinkControlPolicyRelationship       `json:"LinkControlPolicy,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _FabricApplianceRole FabricApplianceRole

// NewFabricApplianceRole instantiates a new FabricApplianceRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricApplianceRole(classId string, objectType string) *FabricApplianceRole {
	this := FabricApplianceRole{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	var fec string = "Auto"
	this.Fec = &fec
	var mode string = "trunk"
	this.Mode = &mode
	var priority string = "Best Effort"
	this.Priority = &priority
	return &this
}

// NewFabricApplianceRoleWithDefaults instantiates a new FabricApplianceRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricApplianceRoleWithDefaults() *FabricApplianceRole {
	this := FabricApplianceRole{}
	var classId string = "fabric.ApplianceRole"
	this.ClassId = classId
	var objectType string = "fabric.ApplianceRole"
	this.ObjectType = objectType
	var mode string = "trunk"
	this.Mode = &mode
	var priority string = "Best Effort"
	this.Priority = &priority
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricApplianceRole) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricApplianceRole) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricApplianceRole) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricApplianceRole) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricApplianceRole) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricApplianceRole) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *FabricApplianceRole) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricApplianceRole) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *FabricApplianceRole) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *FabricApplianceRole) SetMode(v string) {
	o.Mode = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *FabricApplianceRole) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricApplianceRole) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *FabricApplianceRole) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *FabricApplianceRole) SetPriority(v string) {
	o.Priority = &v
}

// GetEthNetworkControlPolicy returns the EthNetworkControlPolicy field value if set, zero value otherwise.
func (o *FabricApplianceRole) GetEthNetworkControlPolicy() FabricEthNetworkControlPolicyRelationship {
	if o == nil || o.EthNetworkControlPolicy == nil {
		var ret FabricEthNetworkControlPolicyRelationship
		return ret
	}
	return *o.EthNetworkControlPolicy
}

// GetEthNetworkControlPolicyOk returns a tuple with the EthNetworkControlPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricApplianceRole) GetEthNetworkControlPolicyOk() (*FabricEthNetworkControlPolicyRelationship, bool) {
	if o == nil || o.EthNetworkControlPolicy == nil {
		return nil, false
	}
	return o.EthNetworkControlPolicy, true
}

// HasEthNetworkControlPolicy returns a boolean if a field has been set.
func (o *FabricApplianceRole) HasEthNetworkControlPolicy() bool {
	if o != nil && o.EthNetworkControlPolicy != nil {
		return true
	}

	return false
}

// SetEthNetworkControlPolicy gets a reference to the given FabricEthNetworkControlPolicyRelationship and assigns it to the EthNetworkControlPolicy field.
func (o *FabricApplianceRole) SetEthNetworkControlPolicy(v FabricEthNetworkControlPolicyRelationship) {
	o.EthNetworkControlPolicy = &v
}

// GetEthNetworkGroupPolicy returns the EthNetworkGroupPolicy field value if set, zero value otherwise.
func (o *FabricApplianceRole) GetEthNetworkGroupPolicy() FabricEthNetworkGroupPolicyRelationship {
	if o == nil || o.EthNetworkGroupPolicy == nil {
		var ret FabricEthNetworkGroupPolicyRelationship
		return ret
	}
	return *o.EthNetworkGroupPolicy
}

// GetEthNetworkGroupPolicyOk returns a tuple with the EthNetworkGroupPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricApplianceRole) GetEthNetworkGroupPolicyOk() (*FabricEthNetworkGroupPolicyRelationship, bool) {
	if o == nil || o.EthNetworkGroupPolicy == nil {
		return nil, false
	}
	return o.EthNetworkGroupPolicy, true
}

// HasEthNetworkGroupPolicy returns a boolean if a field has been set.
func (o *FabricApplianceRole) HasEthNetworkGroupPolicy() bool {
	if o != nil && o.EthNetworkGroupPolicy != nil {
		return true
	}

	return false
}

// SetEthNetworkGroupPolicy gets a reference to the given FabricEthNetworkGroupPolicyRelationship and assigns it to the EthNetworkGroupPolicy field.
func (o *FabricApplianceRole) SetEthNetworkGroupPolicy(v FabricEthNetworkGroupPolicyRelationship) {
	o.EthNetworkGroupPolicy = &v
}

// GetFlowControlPolicy returns the FlowControlPolicy field value if set, zero value otherwise.
func (o *FabricApplianceRole) GetFlowControlPolicy() FabricFlowControlPolicyRelationship {
	if o == nil || o.FlowControlPolicy == nil {
		var ret FabricFlowControlPolicyRelationship
		return ret
	}
	return *o.FlowControlPolicy
}

// GetFlowControlPolicyOk returns a tuple with the FlowControlPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricApplianceRole) GetFlowControlPolicyOk() (*FabricFlowControlPolicyRelationship, bool) {
	if o == nil || o.FlowControlPolicy == nil {
		return nil, false
	}
	return o.FlowControlPolicy, true
}

// HasFlowControlPolicy returns a boolean if a field has been set.
func (o *FabricApplianceRole) HasFlowControlPolicy() bool {
	if o != nil && o.FlowControlPolicy != nil {
		return true
	}

	return false
}

// SetFlowControlPolicy gets a reference to the given FabricFlowControlPolicyRelationship and assigns it to the FlowControlPolicy field.
func (o *FabricApplianceRole) SetFlowControlPolicy(v FabricFlowControlPolicyRelationship) {
	o.FlowControlPolicy = &v
}

// GetLinkControlPolicy returns the LinkControlPolicy field value if set, zero value otherwise.
func (o *FabricApplianceRole) GetLinkControlPolicy() FabricLinkControlPolicyRelationship {
	if o == nil || o.LinkControlPolicy == nil {
		var ret FabricLinkControlPolicyRelationship
		return ret
	}
	return *o.LinkControlPolicy
}

// GetLinkControlPolicyOk returns a tuple with the LinkControlPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricApplianceRole) GetLinkControlPolicyOk() (*FabricLinkControlPolicyRelationship, bool) {
	if o == nil || o.LinkControlPolicy == nil {
		return nil, false
	}
	return o.LinkControlPolicy, true
}

// HasLinkControlPolicy returns a boolean if a field has been set.
func (o *FabricApplianceRole) HasLinkControlPolicy() bool {
	if o != nil && o.LinkControlPolicy != nil {
		return true
	}

	return false
}

// SetLinkControlPolicy gets a reference to the given FabricLinkControlPolicyRelationship and assigns it to the LinkControlPolicy field.
func (o *FabricApplianceRole) SetLinkControlPolicy(v FabricLinkControlPolicyRelationship) {
	o.LinkControlPolicy = &v
}

func (o FabricApplianceRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricTransceiverRole, errFabricTransceiverRole := json.Marshal(o.FabricTransceiverRole)
	if errFabricTransceiverRole != nil {
		return []byte{}, errFabricTransceiverRole
	}
	errFabricTransceiverRole = json.Unmarshal([]byte(serializedFabricTransceiverRole), &toSerialize)
	if errFabricTransceiverRole != nil {
		return []byte{}, errFabricTransceiverRole
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	if o.Priority != nil {
		toSerialize["Priority"] = o.Priority
	}
	if o.EthNetworkControlPolicy != nil {
		toSerialize["EthNetworkControlPolicy"] = o.EthNetworkControlPolicy
	}
	if o.EthNetworkGroupPolicy != nil {
		toSerialize["EthNetworkGroupPolicy"] = o.EthNetworkGroupPolicy
	}
	if o.FlowControlPolicy != nil {
		toSerialize["FlowControlPolicy"] = o.FlowControlPolicy
	}
	if o.LinkControlPolicy != nil {
		toSerialize["LinkControlPolicy"] = o.LinkControlPolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricApplianceRole) UnmarshalJSON(bytes []byte) (err error) {
	type FabricApplianceRoleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Port mode to be set on the appliance port. * `trunk` - Trunk Mode Switch Port Type. * `access` - Access Mode Switch Port Type.
		Mode *string `json:"Mode,omitempty"`
		// The 'name' of the System QoS Class. * `Best Effort` - QoS Priority for Best-effort traffic. * `FC` - QoS Priority for FC traffic. * `Platinum` - QoS Priority for Platinum traffic. * `Gold` - QoS Priority for Gold traffic. * `Silver` - QoS Priority for Silver traffic. * `Bronze` - QoS Priority for Bronze traffic.
		Priority                *string                                    `json:"Priority,omitempty"`
		EthNetworkControlPolicy *FabricEthNetworkControlPolicyRelationship `json:"EthNetworkControlPolicy,omitempty"`
		EthNetworkGroupPolicy   *FabricEthNetworkGroupPolicyRelationship   `json:"EthNetworkGroupPolicy,omitempty"`
		FlowControlPolicy       *FabricFlowControlPolicyRelationship       `json:"FlowControlPolicy,omitempty"`
		LinkControlPolicy       *FabricLinkControlPolicyRelationship       `json:"LinkControlPolicy,omitempty"`
	}

	varFabricApplianceRoleWithoutEmbeddedStruct := FabricApplianceRoleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricApplianceRoleWithoutEmbeddedStruct)
	if err == nil {
		varFabricApplianceRole := _FabricApplianceRole{}
		varFabricApplianceRole.ClassId = varFabricApplianceRoleWithoutEmbeddedStruct.ClassId
		varFabricApplianceRole.ObjectType = varFabricApplianceRoleWithoutEmbeddedStruct.ObjectType
		varFabricApplianceRole.Mode = varFabricApplianceRoleWithoutEmbeddedStruct.Mode
		varFabricApplianceRole.Priority = varFabricApplianceRoleWithoutEmbeddedStruct.Priority
		varFabricApplianceRole.EthNetworkControlPolicy = varFabricApplianceRoleWithoutEmbeddedStruct.EthNetworkControlPolicy
		varFabricApplianceRole.EthNetworkGroupPolicy = varFabricApplianceRoleWithoutEmbeddedStruct.EthNetworkGroupPolicy
		varFabricApplianceRole.FlowControlPolicy = varFabricApplianceRoleWithoutEmbeddedStruct.FlowControlPolicy
		varFabricApplianceRole.LinkControlPolicy = varFabricApplianceRoleWithoutEmbeddedStruct.LinkControlPolicy
		*o = FabricApplianceRole(varFabricApplianceRole)
	} else {
		return err
	}

	varFabricApplianceRole := _FabricApplianceRole{}

	err = json.Unmarshal(bytes, &varFabricApplianceRole)
	if err == nil {
		o.FabricTransceiverRole = varFabricApplianceRole.FabricTransceiverRole
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Mode")
		delete(additionalProperties, "Priority")
		delete(additionalProperties, "EthNetworkControlPolicy")
		delete(additionalProperties, "EthNetworkGroupPolicy")
		delete(additionalProperties, "FlowControlPolicy")
		delete(additionalProperties, "LinkControlPolicy")

		// remove fields from embedded structs
		reflectFabricTransceiverRole := reflect.ValueOf(o.FabricTransceiverRole)
		for i := 0; i < reflectFabricTransceiverRole.Type().NumField(); i++ {
			t := reflectFabricTransceiverRole.Type().Field(i)

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

type NullableFabricApplianceRole struct {
	value *FabricApplianceRole
	isSet bool
}

func (v NullableFabricApplianceRole) Get() *FabricApplianceRole {
	return v.value
}

func (v *NullableFabricApplianceRole) Set(val *FabricApplianceRole) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricApplianceRole) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricApplianceRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricApplianceRole(val *FabricApplianceRole) *NullableFabricApplianceRole {
	return &NullableFabricApplianceRole{value: val, isSet: true}
}

func (v NullableFabricApplianceRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricApplianceRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
