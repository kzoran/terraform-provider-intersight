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

// VirtualizationVmwareVirtualSwitchAllOf Definition of the list of properties defined in 'virtualization.VmwareVirtualSwitch', excluding properties defined in parent classes.
type VirtualizationVmwareVirtualSwitchAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// If forgedTransmits property value is set to reject, outbound frames with a source MAC address different from the one set on the adapter are dropped. If property value is set to accept, the switch does not perform filtering and permits all outbound frames. * `Reject` - Indicates that the security policy is rejected. * `Accept` - Indicates that the security policy is accepted.
	ForgedTransmits *string `json:"ForgedTransmits,omitempty"`
	// If macAddressChanges property value is set to reject and the MAC address of the adapter is changed to a value other than the one specified in .vmx configuration file, the switch drops all inbound frames to the adapter. If property value is set to accept and the MAC address is changed, the switch allows frames to the new MAC address to pass. * `Reject` - Indicates that the security policy is rejected. * `Accept` - Indicates that the security policy is accepted.
	MacAddressChanges *string `json:"MacAddressChanges,omitempty"`
	// Maximum transmission unit configured on a virtual switch.
	Mtu                   *int64                                         `json:"Mtu,omitempty"`
	NicTeamingAndFailover NullableVirtualizationVmwareTeamingAndFailover `json:"NicTeamingAndFailover,omitempty"`
	// Number of networks available on this virtual switch.
	NumNetworks *int64 `json:"NumNetworks,omitempty"`
	// Number of physical network interfaces connected with this virtual switch.
	NumPhysicalNetworkInterfaces *int64 `json:"NumPhysicalNetworkInterfaces,omitempty"`
	// If promiscuousMode property value is set to reject, the virtual switch forwards only frames that are addressed to the adapter. If property value is set to accept, the virtual switch forwards all frames to the adapter in compliance with the active VLAN policy for the port to which it is connected. * `Reject` - Indicates that the security policy is rejected. * `Accept` - Indicates that the security policy is accepted.
	PromiscuousMode      *string                               `json:"PromiscuousMode,omitempty"`
	Host                 *VirtualizationVmwareHostRelationship `json:"Host,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareVirtualSwitchAllOf VirtualizationVmwareVirtualSwitchAllOf

// NewVirtualizationVmwareVirtualSwitchAllOf instantiates a new VirtualizationVmwareVirtualSwitchAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareVirtualSwitchAllOf(classId string, objectType string) *VirtualizationVmwareVirtualSwitchAllOf {
	this := VirtualizationVmwareVirtualSwitchAllOf{}
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

// NewVirtualizationVmwareVirtualSwitchAllOfWithDefaults instantiates a new VirtualizationVmwareVirtualSwitchAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareVirtualSwitchAllOfWithDefaults() *VirtualizationVmwareVirtualSwitchAllOf {
	this := VirtualizationVmwareVirtualSwitchAllOf{}
	var classId string = "virtualization.VmwareVirtualSwitch"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareVirtualSwitch"
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
func (o *VirtualizationVmwareVirtualSwitchAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualSwitchAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareVirtualSwitchAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareVirtualSwitchAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualSwitchAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareVirtualSwitchAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetForgedTransmits returns the ForgedTransmits field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualSwitchAllOf) GetForgedTransmits() string {
	if o == nil || o.ForgedTransmits == nil {
		var ret string
		return ret
	}
	return *o.ForgedTransmits
}

// GetForgedTransmitsOk returns a tuple with the ForgedTransmits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualSwitchAllOf) GetForgedTransmitsOk() (*string, bool) {
	if o == nil || o.ForgedTransmits == nil {
		return nil, false
	}
	return o.ForgedTransmits, true
}

// HasForgedTransmits returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualSwitchAllOf) HasForgedTransmits() bool {
	if o != nil && o.ForgedTransmits != nil {
		return true
	}

	return false
}

// SetForgedTransmits gets a reference to the given string and assigns it to the ForgedTransmits field.
func (o *VirtualizationVmwareVirtualSwitchAllOf) SetForgedTransmits(v string) {
	o.ForgedTransmits = &v
}

// GetMacAddressChanges returns the MacAddressChanges field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualSwitchAllOf) GetMacAddressChanges() string {
	if o == nil || o.MacAddressChanges == nil {
		var ret string
		return ret
	}
	return *o.MacAddressChanges
}

// GetMacAddressChangesOk returns a tuple with the MacAddressChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualSwitchAllOf) GetMacAddressChangesOk() (*string, bool) {
	if o == nil || o.MacAddressChanges == nil {
		return nil, false
	}
	return o.MacAddressChanges, true
}

// HasMacAddressChanges returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualSwitchAllOf) HasMacAddressChanges() bool {
	if o != nil && o.MacAddressChanges != nil {
		return true
	}

	return false
}

// SetMacAddressChanges gets a reference to the given string and assigns it to the MacAddressChanges field.
func (o *VirtualizationVmwareVirtualSwitchAllOf) SetMacAddressChanges(v string) {
	o.MacAddressChanges = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualSwitchAllOf) GetMtu() int64 {
	if o == nil || o.Mtu == nil {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualSwitchAllOf) GetMtuOk() (*int64, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualSwitchAllOf) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *VirtualizationVmwareVirtualSwitchAllOf) SetMtu(v int64) {
	o.Mtu = &v
}

// GetNicTeamingAndFailover returns the NicTeamingAndFailover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareVirtualSwitchAllOf) GetNicTeamingAndFailover() VirtualizationVmwareTeamingAndFailover {
	if o == nil || o.NicTeamingAndFailover.Get() == nil {
		var ret VirtualizationVmwareTeamingAndFailover
		return ret
	}
	return *o.NicTeamingAndFailover.Get()
}

// GetNicTeamingAndFailoverOk returns a tuple with the NicTeamingAndFailover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareVirtualSwitchAllOf) GetNicTeamingAndFailoverOk() (*VirtualizationVmwareTeamingAndFailover, bool) {
	if o == nil {
		return nil, false
	}
	return o.NicTeamingAndFailover.Get(), o.NicTeamingAndFailover.IsSet()
}

// HasNicTeamingAndFailover returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualSwitchAllOf) HasNicTeamingAndFailover() bool {
	if o != nil && o.NicTeamingAndFailover.IsSet() {
		return true
	}

	return false
}

// SetNicTeamingAndFailover gets a reference to the given NullableVirtualizationVmwareTeamingAndFailover and assigns it to the NicTeamingAndFailover field.
func (o *VirtualizationVmwareVirtualSwitchAllOf) SetNicTeamingAndFailover(v VirtualizationVmwareTeamingAndFailover) {
	o.NicTeamingAndFailover.Set(&v)
}

// SetNicTeamingAndFailoverNil sets the value for NicTeamingAndFailover to be an explicit nil
func (o *VirtualizationVmwareVirtualSwitchAllOf) SetNicTeamingAndFailoverNil() {
	o.NicTeamingAndFailover.Set(nil)
}

// UnsetNicTeamingAndFailover ensures that no value is present for NicTeamingAndFailover, not even an explicit nil
func (o *VirtualizationVmwareVirtualSwitchAllOf) UnsetNicTeamingAndFailover() {
	o.NicTeamingAndFailover.Unset()
}

// GetNumNetworks returns the NumNetworks field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualSwitchAllOf) GetNumNetworks() int64 {
	if o == nil || o.NumNetworks == nil {
		var ret int64
		return ret
	}
	return *o.NumNetworks
}

// GetNumNetworksOk returns a tuple with the NumNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualSwitchAllOf) GetNumNetworksOk() (*int64, bool) {
	if o == nil || o.NumNetworks == nil {
		return nil, false
	}
	return o.NumNetworks, true
}

// HasNumNetworks returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualSwitchAllOf) HasNumNetworks() bool {
	if o != nil && o.NumNetworks != nil {
		return true
	}

	return false
}

// SetNumNetworks gets a reference to the given int64 and assigns it to the NumNetworks field.
func (o *VirtualizationVmwareVirtualSwitchAllOf) SetNumNetworks(v int64) {
	o.NumNetworks = &v
}

// GetNumPhysicalNetworkInterfaces returns the NumPhysicalNetworkInterfaces field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualSwitchAllOf) GetNumPhysicalNetworkInterfaces() int64 {
	if o == nil || o.NumPhysicalNetworkInterfaces == nil {
		var ret int64
		return ret
	}
	return *o.NumPhysicalNetworkInterfaces
}

// GetNumPhysicalNetworkInterfacesOk returns a tuple with the NumPhysicalNetworkInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualSwitchAllOf) GetNumPhysicalNetworkInterfacesOk() (*int64, bool) {
	if o == nil || o.NumPhysicalNetworkInterfaces == nil {
		return nil, false
	}
	return o.NumPhysicalNetworkInterfaces, true
}

// HasNumPhysicalNetworkInterfaces returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualSwitchAllOf) HasNumPhysicalNetworkInterfaces() bool {
	if o != nil && o.NumPhysicalNetworkInterfaces != nil {
		return true
	}

	return false
}

// SetNumPhysicalNetworkInterfaces gets a reference to the given int64 and assigns it to the NumPhysicalNetworkInterfaces field.
func (o *VirtualizationVmwareVirtualSwitchAllOf) SetNumPhysicalNetworkInterfaces(v int64) {
	o.NumPhysicalNetworkInterfaces = &v
}

// GetPromiscuousMode returns the PromiscuousMode field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualSwitchAllOf) GetPromiscuousMode() string {
	if o == nil || o.PromiscuousMode == nil {
		var ret string
		return ret
	}
	return *o.PromiscuousMode
}

// GetPromiscuousModeOk returns a tuple with the PromiscuousMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualSwitchAllOf) GetPromiscuousModeOk() (*string, bool) {
	if o == nil || o.PromiscuousMode == nil {
		return nil, false
	}
	return o.PromiscuousMode, true
}

// HasPromiscuousMode returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualSwitchAllOf) HasPromiscuousMode() bool {
	if o != nil && o.PromiscuousMode != nil {
		return true
	}

	return false
}

// SetPromiscuousMode gets a reference to the given string and assigns it to the PromiscuousMode field.
func (o *VirtualizationVmwareVirtualSwitchAllOf) SetPromiscuousMode(v string) {
	o.PromiscuousMode = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualSwitchAllOf) GetHost() VirtualizationVmwareHostRelationship {
	if o == nil || o.Host == nil {
		var ret VirtualizationVmwareHostRelationship
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualSwitchAllOf) GetHostOk() (*VirtualizationVmwareHostRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualSwitchAllOf) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given VirtualizationVmwareHostRelationship and assigns it to the Host field.
func (o *VirtualizationVmwareVirtualSwitchAllOf) SetHost(v VirtualizationVmwareHostRelationship) {
	o.Host = &v
}

func (o VirtualizationVmwareVirtualSwitchAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.Mtu != nil {
		toSerialize["Mtu"] = o.Mtu
	}
	if o.NicTeamingAndFailover.IsSet() {
		toSerialize["NicTeamingAndFailover"] = o.NicTeamingAndFailover.Get()
	}
	if o.NumNetworks != nil {
		toSerialize["NumNetworks"] = o.NumNetworks
	}
	if o.NumPhysicalNetworkInterfaces != nil {
		toSerialize["NumPhysicalNetworkInterfaces"] = o.NumPhysicalNetworkInterfaces
	}
	if o.PromiscuousMode != nil {
		toSerialize["PromiscuousMode"] = o.PromiscuousMode
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareVirtualSwitchAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationVmwareVirtualSwitchAllOf := _VirtualizationVmwareVirtualSwitchAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationVmwareVirtualSwitchAllOf); err == nil {
		*o = VirtualizationVmwareVirtualSwitchAllOf(varVirtualizationVmwareVirtualSwitchAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ForgedTransmits")
		delete(additionalProperties, "MacAddressChanges")
		delete(additionalProperties, "Mtu")
		delete(additionalProperties, "NicTeamingAndFailover")
		delete(additionalProperties, "NumNetworks")
		delete(additionalProperties, "NumPhysicalNetworkInterfaces")
		delete(additionalProperties, "PromiscuousMode")
		delete(additionalProperties, "Host")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVmwareVirtualSwitchAllOf struct {
	value *VirtualizationVmwareVirtualSwitchAllOf
	isSet bool
}

func (v NullableVirtualizationVmwareVirtualSwitchAllOf) Get() *VirtualizationVmwareVirtualSwitchAllOf {
	return v.value
}

func (v *NullableVirtualizationVmwareVirtualSwitchAllOf) Set(val *VirtualizationVmwareVirtualSwitchAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareVirtualSwitchAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareVirtualSwitchAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareVirtualSwitchAllOf(val *VirtualizationVmwareVirtualSwitchAllOf) *NullableVirtualizationVmwareVirtualSwitchAllOf {
	return &NullableVirtualizationVmwareVirtualSwitchAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareVirtualSwitchAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareVirtualSwitchAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
