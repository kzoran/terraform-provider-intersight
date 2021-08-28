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

// VirtualizationNetworkInterfaceAllOf Definition of the list of properties defined in 'virtualization.NetworkInterface', excluding properties defined in parent classes.
type VirtualizationNetworkInterfaceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Virtual machine network adaptor type. * `Unknown` - The type of the network adaptor type is unknown. * `E1000` - Emulated version of the Intel 82545EM Gigabit Ethernet NIC. * `SRIOV` - Representation of a virtual function (VF) on a physical NIC with SR-IOV support. * `VMXNET2` - VMXNET 2 (Enhanced) is available only for some guest operating systems on ESX/ESXi 3.5 and later. * `VMXNET3` - VMXNET 3 offers all the features available in VMXNET 2 and adds several new features. * `E1000E` - E1000E – emulates a newer real network adapter, the 1 Gbit Intel 82574, and is available for Windows 2012 and later. The E1000E needs virtual machine hardware version 8 or later. * `NE2K_PCI` - The Ne2000 network card uses two ring buffers for packet handling. These are circular buffers made of 256-byte pages that the chip's DMA logic will use to store received packets or to get received packets. * `PCnet` - The PCnet-PCI II is a PCI network adapter. It has built-in support for CRC checks and can automatically pad short packets to the minimum Ethernet length. * `RTL8139` - The RTL8139 is a fast Ethernet card that operates at 10/100 Mbps. It is compliant with PCI version 2.0/2.1 and it is known for reliability and superior performance. * `VirtIO` - VirtIO is a standardized interface which allows virtual machines access to simplified \"virtual\" devices, such as block devices, network adapters and consoles. Accessing devices through VirtIO on a guest VM improves performance over more traditional \"emulated\" devices, as VirtIO devices require only the bare minimum setup and configuration needed to send and receive data, while the host machine handles the majority of the setup and maintenance of the actual physical hardware. * `` - Default network adaptor type supported by the hypervisor.
	AdaptorType *string `json:"AdaptorType,omitempty"`
	// Virtual machine network bridge name.
	Bridge *string `json:"Bridge,omitempty"`
	// Connect the adaptor at virtual machine power on.
	ConnectAtPowerOn *bool `json:"ConnectAtPowerOn,omitempty"`
	// Enable the direct path I/O.
	DirectPathIo *bool `json:"DirectPathIo,omitempty"`
	// Virtual machine network mac address.
	MacAddress *string `json:"MacAddress,omitempty"`
	// Name of the network interface. This may be different from guest operating assigned.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationNetworkInterfaceAllOf VirtualizationNetworkInterfaceAllOf

// NewVirtualizationNetworkInterfaceAllOf instantiates a new VirtualizationNetworkInterfaceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationNetworkInterfaceAllOf(classId string, objectType string) *VirtualizationNetworkInterfaceAllOf {
	this := VirtualizationNetworkInterfaceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adaptorType string = "Unknown"
	this.AdaptorType = &adaptorType
	return &this
}

// NewVirtualizationNetworkInterfaceAllOfWithDefaults instantiates a new VirtualizationNetworkInterfaceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationNetworkInterfaceAllOfWithDefaults() *VirtualizationNetworkInterfaceAllOf {
	this := VirtualizationNetworkInterfaceAllOf{}
	var classId string = "virtualization.NetworkInterface"
	this.ClassId = classId
	var objectType string = "virtualization.NetworkInterface"
	this.ObjectType = objectType
	var adaptorType string = "Unknown"
	this.AdaptorType = &adaptorType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationNetworkInterfaceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterfaceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationNetworkInterfaceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationNetworkInterfaceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterfaceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationNetworkInterfaceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdaptorType returns the AdaptorType field value if set, zero value otherwise.
func (o *VirtualizationNetworkInterfaceAllOf) GetAdaptorType() string {
	if o == nil || o.AdaptorType == nil {
		var ret string
		return ret
	}
	return *o.AdaptorType
}

// GetAdaptorTypeOk returns a tuple with the AdaptorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterfaceAllOf) GetAdaptorTypeOk() (*string, bool) {
	if o == nil || o.AdaptorType == nil {
		return nil, false
	}
	return o.AdaptorType, true
}

// HasAdaptorType returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterfaceAllOf) HasAdaptorType() bool {
	if o != nil && o.AdaptorType != nil {
		return true
	}

	return false
}

// SetAdaptorType gets a reference to the given string and assigns it to the AdaptorType field.
func (o *VirtualizationNetworkInterfaceAllOf) SetAdaptorType(v string) {
	o.AdaptorType = &v
}

// GetBridge returns the Bridge field value if set, zero value otherwise.
func (o *VirtualizationNetworkInterfaceAllOf) GetBridge() string {
	if o == nil || o.Bridge == nil {
		var ret string
		return ret
	}
	return *o.Bridge
}

// GetBridgeOk returns a tuple with the Bridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterfaceAllOf) GetBridgeOk() (*string, bool) {
	if o == nil || o.Bridge == nil {
		return nil, false
	}
	return o.Bridge, true
}

// HasBridge returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterfaceAllOf) HasBridge() bool {
	if o != nil && o.Bridge != nil {
		return true
	}

	return false
}

// SetBridge gets a reference to the given string and assigns it to the Bridge field.
func (o *VirtualizationNetworkInterfaceAllOf) SetBridge(v string) {
	o.Bridge = &v
}

// GetConnectAtPowerOn returns the ConnectAtPowerOn field value if set, zero value otherwise.
func (o *VirtualizationNetworkInterfaceAllOf) GetConnectAtPowerOn() bool {
	if o == nil || o.ConnectAtPowerOn == nil {
		var ret bool
		return ret
	}
	return *o.ConnectAtPowerOn
}

// GetConnectAtPowerOnOk returns a tuple with the ConnectAtPowerOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterfaceAllOf) GetConnectAtPowerOnOk() (*bool, bool) {
	if o == nil || o.ConnectAtPowerOn == nil {
		return nil, false
	}
	return o.ConnectAtPowerOn, true
}

// HasConnectAtPowerOn returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterfaceAllOf) HasConnectAtPowerOn() bool {
	if o != nil && o.ConnectAtPowerOn != nil {
		return true
	}

	return false
}

// SetConnectAtPowerOn gets a reference to the given bool and assigns it to the ConnectAtPowerOn field.
func (o *VirtualizationNetworkInterfaceAllOf) SetConnectAtPowerOn(v bool) {
	o.ConnectAtPowerOn = &v
}

// GetDirectPathIo returns the DirectPathIo field value if set, zero value otherwise.
func (o *VirtualizationNetworkInterfaceAllOf) GetDirectPathIo() bool {
	if o == nil || o.DirectPathIo == nil {
		var ret bool
		return ret
	}
	return *o.DirectPathIo
}

// GetDirectPathIoOk returns a tuple with the DirectPathIo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterfaceAllOf) GetDirectPathIoOk() (*bool, bool) {
	if o == nil || o.DirectPathIo == nil {
		return nil, false
	}
	return o.DirectPathIo, true
}

// HasDirectPathIo returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterfaceAllOf) HasDirectPathIo() bool {
	if o != nil && o.DirectPathIo != nil {
		return true
	}

	return false
}

// SetDirectPathIo gets a reference to the given bool and assigns it to the DirectPathIo field.
func (o *VirtualizationNetworkInterfaceAllOf) SetDirectPathIo(v bool) {
	o.DirectPathIo = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *VirtualizationNetworkInterfaceAllOf) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterfaceAllOf) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterfaceAllOf) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *VirtualizationNetworkInterfaceAllOf) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationNetworkInterfaceAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationNetworkInterfaceAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationNetworkInterfaceAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationNetworkInterfaceAllOf) SetName(v string) {
	o.Name = &v
}

func (o VirtualizationNetworkInterfaceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdaptorType != nil {
		toSerialize["AdaptorType"] = o.AdaptorType
	}
	if o.Bridge != nil {
		toSerialize["Bridge"] = o.Bridge
	}
	if o.ConnectAtPowerOn != nil {
		toSerialize["ConnectAtPowerOn"] = o.ConnectAtPowerOn
	}
	if o.DirectPathIo != nil {
		toSerialize["DirectPathIo"] = o.DirectPathIo
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationNetworkInterfaceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationNetworkInterfaceAllOf := _VirtualizationNetworkInterfaceAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationNetworkInterfaceAllOf); err == nil {
		*o = VirtualizationNetworkInterfaceAllOf(varVirtualizationNetworkInterfaceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdaptorType")
		delete(additionalProperties, "Bridge")
		delete(additionalProperties, "ConnectAtPowerOn")
		delete(additionalProperties, "DirectPathIo")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "Name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationNetworkInterfaceAllOf struct {
	value *VirtualizationNetworkInterfaceAllOf
	isSet bool
}

func (v NullableVirtualizationNetworkInterfaceAllOf) Get() *VirtualizationNetworkInterfaceAllOf {
	return v.value
}

func (v *NullableVirtualizationNetworkInterfaceAllOf) Set(val *VirtualizationNetworkInterfaceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationNetworkInterfaceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationNetworkInterfaceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationNetworkInterfaceAllOf(val *VirtualizationNetworkInterfaceAllOf) *NullableVirtualizationNetworkInterfaceAllOf {
	return &NullableVirtualizationNetworkInterfaceAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationNetworkInterfaceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationNetworkInterfaceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
