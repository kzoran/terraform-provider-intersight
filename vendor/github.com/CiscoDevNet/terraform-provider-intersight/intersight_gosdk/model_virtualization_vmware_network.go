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

// VirtualizationVmwareNetwork Details of VMware Virtual Network.
type VirtualizationVmwareNetwork struct {
	VirtualizationBaseVirtualNetwork
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// If forgedTransmits property value is set to reject, outbound frames with source MAC address different from the one set on the adapter are dropped. If property value is set to accept, no filtering is performed and all outbound frames are passed. * `Reject` - Indicates that the security policy is rejected. * `Accept` - Indicates that the security policy is accepted.
	ForgedTransmits *string `json:"ForgedTransmits,omitempty"`
	// If macAddressChanges property value is set to reject and the MAC address of the adapter is changed to a value other than the one specified in .vmx configuration file, all inbound frames are dropped. If property value is set to accept and the MAC address is changed, inbound frames to the new MAC address are received. * `Reject` - Indicates that the security policy is rejected. * `Accept` - Indicates that the security policy is accepted.
	MacAddressChanges     *string                                        `json:"MacAddressChanges,omitempty"`
	NicTeamingAndFailover NullableVirtualizationVmwareTeamingAndFailover `json:"NicTeamingAndFailover,omitempty"`
	// If promiscuousMode property value is set to reject, incoming traffic only targeted to that network will be visible. If property value is set to accept, objects defined within the network can see all incoming traffic on the virtual switch based on the VLAN policy. * `Reject` - Indicates that the security policy is rejected. * `Accept` - Indicates that the security policy is accepted.
	PromiscuousMode *string `json:"PromiscuousMode,omitempty"`
	// VLAN id with which the network is associated. A value of 0 specifies that port is not associated with a VLAN.
	VlanId               *int64                                         `json:"VlanId,omitempty"`
	Host                 *VirtualizationVmwareHostRelationship          `json:"Host,omitempty"`
	VirtualSwitch        *VirtualizationVmwareVirtualSwitchRelationship `json:"VirtualSwitch,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareNetwork VirtualizationVmwareNetwork

// NewVirtualizationVmwareNetwork instantiates a new VirtualizationVmwareNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareNetwork(classId string, objectType string) *VirtualizationVmwareNetwork {
	this := VirtualizationVmwareNetwork{}
	this.ClassId = classId
	this.ObjectType = objectType
	var forgedTransmits string = "Reject"
	this.ForgedTransmits = &forgedTransmits
	var macAddressChanges string = "Reject"
	this.MacAddressChanges = &macAddressChanges
	var promiscuousMode string = "Reject"
	this.PromiscuousMode = &promiscuousMode
	return &this
}

// NewVirtualizationVmwareNetworkWithDefaults instantiates a new VirtualizationVmwareNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareNetworkWithDefaults() *VirtualizationVmwareNetwork {
	this := VirtualizationVmwareNetwork{}
	var classId string = "virtualization.VmwareNetwork"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareNetwork"
	this.ObjectType = objectType
	var forgedTransmits string = "Reject"
	this.ForgedTransmits = &forgedTransmits
	var macAddressChanges string = "Reject"
	this.MacAddressChanges = &macAddressChanges
	var promiscuousMode string = "Reject"
	this.PromiscuousMode = &promiscuousMode
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareNetwork) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareNetwork) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareNetwork) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareNetwork) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareNetwork) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareNetwork) SetObjectType(v string) {
	o.ObjectType = v
}

// GetForgedTransmits returns the ForgedTransmits field value if set, zero value otherwise.
func (o *VirtualizationVmwareNetwork) GetForgedTransmits() string {
	if o == nil || o.ForgedTransmits == nil {
		var ret string
		return ret
	}
	return *o.ForgedTransmits
}

// GetForgedTransmitsOk returns a tuple with the ForgedTransmits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareNetwork) GetForgedTransmitsOk() (*string, bool) {
	if o == nil || o.ForgedTransmits == nil {
		return nil, false
	}
	return o.ForgedTransmits, true
}

// HasForgedTransmits returns a boolean if a field has been set.
func (o *VirtualizationVmwareNetwork) HasForgedTransmits() bool {
	if o != nil && o.ForgedTransmits != nil {
		return true
	}

	return false
}

// SetForgedTransmits gets a reference to the given string and assigns it to the ForgedTransmits field.
func (o *VirtualizationVmwareNetwork) SetForgedTransmits(v string) {
	o.ForgedTransmits = &v
}

// GetMacAddressChanges returns the MacAddressChanges field value if set, zero value otherwise.
func (o *VirtualizationVmwareNetwork) GetMacAddressChanges() string {
	if o == nil || o.MacAddressChanges == nil {
		var ret string
		return ret
	}
	return *o.MacAddressChanges
}

// GetMacAddressChangesOk returns a tuple with the MacAddressChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareNetwork) GetMacAddressChangesOk() (*string, bool) {
	if o == nil || o.MacAddressChanges == nil {
		return nil, false
	}
	return o.MacAddressChanges, true
}

// HasMacAddressChanges returns a boolean if a field has been set.
func (o *VirtualizationVmwareNetwork) HasMacAddressChanges() bool {
	if o != nil && o.MacAddressChanges != nil {
		return true
	}

	return false
}

// SetMacAddressChanges gets a reference to the given string and assigns it to the MacAddressChanges field.
func (o *VirtualizationVmwareNetwork) SetMacAddressChanges(v string) {
	o.MacAddressChanges = &v
}

// GetNicTeamingAndFailover returns the NicTeamingAndFailover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareNetwork) GetNicTeamingAndFailover() VirtualizationVmwareTeamingAndFailover {
	if o == nil || o.NicTeamingAndFailover.Get() == nil {
		var ret VirtualizationVmwareTeamingAndFailover
		return ret
	}
	return *o.NicTeamingAndFailover.Get()
}

// GetNicTeamingAndFailoverOk returns a tuple with the NicTeamingAndFailover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareNetwork) GetNicTeamingAndFailoverOk() (*VirtualizationVmwareTeamingAndFailover, bool) {
	if o == nil {
		return nil, false
	}
	return o.NicTeamingAndFailover.Get(), o.NicTeamingAndFailover.IsSet()
}

// HasNicTeamingAndFailover returns a boolean if a field has been set.
func (o *VirtualizationVmwareNetwork) HasNicTeamingAndFailover() bool {
	if o != nil && o.NicTeamingAndFailover.IsSet() {
		return true
	}

	return false
}

// SetNicTeamingAndFailover gets a reference to the given NullableVirtualizationVmwareTeamingAndFailover and assigns it to the NicTeamingAndFailover field.
func (o *VirtualizationVmwareNetwork) SetNicTeamingAndFailover(v VirtualizationVmwareTeamingAndFailover) {
	o.NicTeamingAndFailover.Set(&v)
}

// SetNicTeamingAndFailoverNil sets the value for NicTeamingAndFailover to be an explicit nil
func (o *VirtualizationVmwareNetwork) SetNicTeamingAndFailoverNil() {
	o.NicTeamingAndFailover.Set(nil)
}

// UnsetNicTeamingAndFailover ensures that no value is present for NicTeamingAndFailover, not even an explicit nil
func (o *VirtualizationVmwareNetwork) UnsetNicTeamingAndFailover() {
	o.NicTeamingAndFailover.Unset()
}

// GetPromiscuousMode returns the PromiscuousMode field value if set, zero value otherwise.
func (o *VirtualizationVmwareNetwork) GetPromiscuousMode() string {
	if o == nil || o.PromiscuousMode == nil {
		var ret string
		return ret
	}
	return *o.PromiscuousMode
}

// GetPromiscuousModeOk returns a tuple with the PromiscuousMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareNetwork) GetPromiscuousModeOk() (*string, bool) {
	if o == nil || o.PromiscuousMode == nil {
		return nil, false
	}
	return o.PromiscuousMode, true
}

// HasPromiscuousMode returns a boolean if a field has been set.
func (o *VirtualizationVmwareNetwork) HasPromiscuousMode() bool {
	if o != nil && o.PromiscuousMode != nil {
		return true
	}

	return false
}

// SetPromiscuousMode gets a reference to the given string and assigns it to the PromiscuousMode field.
func (o *VirtualizationVmwareNetwork) SetPromiscuousMode(v string) {
	o.PromiscuousMode = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *VirtualizationVmwareNetwork) GetVlanId() int64 {
	if o == nil || o.VlanId == nil {
		var ret int64
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareNetwork) GetVlanIdOk() (*int64, bool) {
	if o == nil || o.VlanId == nil {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *VirtualizationVmwareNetwork) HasVlanId() bool {
	if o != nil && o.VlanId != nil {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int64 and assigns it to the VlanId field.
func (o *VirtualizationVmwareNetwork) SetVlanId(v int64) {
	o.VlanId = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *VirtualizationVmwareNetwork) GetHost() VirtualizationVmwareHostRelationship {
	if o == nil || o.Host == nil {
		var ret VirtualizationVmwareHostRelationship
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareNetwork) GetHostOk() (*VirtualizationVmwareHostRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *VirtualizationVmwareNetwork) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given VirtualizationVmwareHostRelationship and assigns it to the Host field.
func (o *VirtualizationVmwareNetwork) SetHost(v VirtualizationVmwareHostRelationship) {
	o.Host = &v
}

// GetVirtualSwitch returns the VirtualSwitch field value if set, zero value otherwise.
func (o *VirtualizationVmwareNetwork) GetVirtualSwitch() VirtualizationVmwareVirtualSwitchRelationship {
	if o == nil || o.VirtualSwitch == nil {
		var ret VirtualizationVmwareVirtualSwitchRelationship
		return ret
	}
	return *o.VirtualSwitch
}

// GetVirtualSwitchOk returns a tuple with the VirtualSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareNetwork) GetVirtualSwitchOk() (*VirtualizationVmwareVirtualSwitchRelationship, bool) {
	if o == nil || o.VirtualSwitch == nil {
		return nil, false
	}
	return o.VirtualSwitch, true
}

// HasVirtualSwitch returns a boolean if a field has been set.
func (o *VirtualizationVmwareNetwork) HasVirtualSwitch() bool {
	if o != nil && o.VirtualSwitch != nil {
		return true
	}

	return false
}

// SetVirtualSwitch gets a reference to the given VirtualizationVmwareVirtualSwitchRelationship and assigns it to the VirtualSwitch field.
func (o *VirtualizationVmwareNetwork) SetVirtualSwitch(v VirtualizationVmwareVirtualSwitchRelationship) {
	o.VirtualSwitch = &v
}

func (o VirtualizationVmwareNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseVirtualNetwork, errVirtualizationBaseVirtualNetwork := json.Marshal(o.VirtualizationBaseVirtualNetwork)
	if errVirtualizationBaseVirtualNetwork != nil {
		return []byte{}, errVirtualizationBaseVirtualNetwork
	}
	errVirtualizationBaseVirtualNetwork = json.Unmarshal([]byte(serializedVirtualizationBaseVirtualNetwork), &toSerialize)
	if errVirtualizationBaseVirtualNetwork != nil {
		return []byte{}, errVirtualizationBaseVirtualNetwork
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ForgedTransmits != nil {
		toSerialize["ForgedTransmits"] = o.ForgedTransmits
	}
	if o.MacAddressChanges != nil {
		toSerialize["MacAddressChanges"] = o.MacAddressChanges
	}
	if o.NicTeamingAndFailover.IsSet() {
		toSerialize["NicTeamingAndFailover"] = o.NicTeamingAndFailover.Get()
	}
	if o.PromiscuousMode != nil {
		toSerialize["PromiscuousMode"] = o.PromiscuousMode
	}
	if o.VlanId != nil {
		toSerialize["VlanId"] = o.VlanId
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
	}
	if o.VirtualSwitch != nil {
		toSerialize["VirtualSwitch"] = o.VirtualSwitch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareNetwork) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVmwareNetworkWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// If forgedTransmits property value is set to reject, outbound frames with source MAC address different from the one set on the adapter are dropped. If property value is set to accept, no filtering is performed and all outbound frames are passed. * `Reject` - Indicates that the security policy is rejected. * `Accept` - Indicates that the security policy is accepted.
		ForgedTransmits *string `json:"ForgedTransmits,omitempty"`
		// If macAddressChanges property value is set to reject and the MAC address of the adapter is changed to a value other than the one specified in .vmx configuration file, all inbound frames are dropped. If property value is set to accept and the MAC address is changed, inbound frames to the new MAC address are received. * `Reject` - Indicates that the security policy is rejected. * `Accept` - Indicates that the security policy is accepted.
		MacAddressChanges     *string                                        `json:"MacAddressChanges,omitempty"`
		NicTeamingAndFailover NullableVirtualizationVmwareTeamingAndFailover `json:"NicTeamingAndFailover,omitempty"`
		// If promiscuousMode property value is set to reject, incoming traffic only targeted to that network will be visible. If property value is set to accept, objects defined within the network can see all incoming traffic on the virtual switch based on the VLAN policy. * `Reject` - Indicates that the security policy is rejected. * `Accept` - Indicates that the security policy is accepted.
		PromiscuousMode *string `json:"PromiscuousMode,omitempty"`
		// VLAN id with which the network is associated. A value of 0 specifies that port is not associated with a VLAN.
		VlanId        *int64                                         `json:"VlanId,omitempty"`
		Host          *VirtualizationVmwareHostRelationship          `json:"Host,omitempty"`
		VirtualSwitch *VirtualizationVmwareVirtualSwitchRelationship `json:"VirtualSwitch,omitempty"`
	}

	varVirtualizationVmwareNetworkWithoutEmbeddedStruct := VirtualizationVmwareNetworkWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareNetworkWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareNetwork := _VirtualizationVmwareNetwork{}
		varVirtualizationVmwareNetwork.ClassId = varVirtualizationVmwareNetworkWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareNetwork.ObjectType = varVirtualizationVmwareNetworkWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareNetwork.ForgedTransmits = varVirtualizationVmwareNetworkWithoutEmbeddedStruct.ForgedTransmits
		varVirtualizationVmwareNetwork.MacAddressChanges = varVirtualizationVmwareNetworkWithoutEmbeddedStruct.MacAddressChanges
		varVirtualizationVmwareNetwork.NicTeamingAndFailover = varVirtualizationVmwareNetworkWithoutEmbeddedStruct.NicTeamingAndFailover
		varVirtualizationVmwareNetwork.PromiscuousMode = varVirtualizationVmwareNetworkWithoutEmbeddedStruct.PromiscuousMode
		varVirtualizationVmwareNetwork.VlanId = varVirtualizationVmwareNetworkWithoutEmbeddedStruct.VlanId
		varVirtualizationVmwareNetwork.Host = varVirtualizationVmwareNetworkWithoutEmbeddedStruct.Host
		varVirtualizationVmwareNetwork.VirtualSwitch = varVirtualizationVmwareNetworkWithoutEmbeddedStruct.VirtualSwitch
		*o = VirtualizationVmwareNetwork(varVirtualizationVmwareNetwork)
	} else {
		return err
	}

	varVirtualizationVmwareNetwork := _VirtualizationVmwareNetwork{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareNetwork)
	if err == nil {
		o.VirtualizationBaseVirtualNetwork = varVirtualizationVmwareNetwork.VirtualizationBaseVirtualNetwork
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ForgedTransmits")
		delete(additionalProperties, "MacAddressChanges")
		delete(additionalProperties, "NicTeamingAndFailover")
		delete(additionalProperties, "PromiscuousMode")
		delete(additionalProperties, "VlanId")
		delete(additionalProperties, "Host")
		delete(additionalProperties, "VirtualSwitch")

		// remove fields from embedded structs
		reflectVirtualizationBaseVirtualNetwork := reflect.ValueOf(o.VirtualizationBaseVirtualNetwork)
		for i := 0; i < reflectVirtualizationBaseVirtualNetwork.Type().NumField(); i++ {
			t := reflectVirtualizationBaseVirtualNetwork.Type().Field(i)

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

type NullableVirtualizationVmwareNetwork struct {
	value *VirtualizationVmwareNetwork
	isSet bool
}

func (v NullableVirtualizationVmwareNetwork) Get() *VirtualizationVmwareNetwork {
	return v.value
}

func (v *NullableVirtualizationVmwareNetwork) Set(val *VirtualizationVmwareNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareNetwork(val *VirtualizationVmwareNetwork) *NullableVirtualizationVmwareNetwork {
	return &NullableVirtualizationVmwareNetwork{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
