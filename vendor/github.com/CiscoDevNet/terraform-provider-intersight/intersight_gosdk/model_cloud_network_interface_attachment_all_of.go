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

// CloudNetworkInterfaceAttachmentAllOf Definition of the list of properties defined in 'cloud.NetworkInterfaceAttachment', excluding properties defined in parent classes.
type CloudNetworkInterfaceAttachmentAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType   string                           `json:"ObjectType"`
	AccessConfig NullableCloudNetworkAccessConfig `json:"AccessConfig,omitempty"`
	// The internally generated identity of this network interface attachment.
	Identity *string `json:"Identity,omitempty"`
	// Set to true, if IP forwarding is enabled on this interface.
	IpForwardingEnabled *bool `json:"IpForwardingEnabled,omitempty"`
	// The MAC (Media Access Control) address of the network interface.
	MacAddress *string `json:"MacAddress,omitempty"`
	// The identity of the network to which this network interface attachment belongs.
	NetworkId *string `json:"NetworkId,omitempty"`
	// User friendly name of the network to which this network interface attachment belongs.
	NetworkName *string `json:"NetworkName,omitempty"`
	// The device index of the network interface attachment in the virtual machine.
	NicIndex       *int64                `json:"NicIndex,omitempty"`
	PrivateAddress []CloudNetworkAddress `json:"PrivateAddress,omitempty"`
	PublicAddress  []CloudNetworkAddress `json:"PublicAddress,omitempty"`
	SecurityGroups []string              `json:"SecurityGroups,omitempty"`
	// The identity of this network interface attachment's subnet.
	SubNetworkId *string `json:"SubNetworkId,omitempty"`
	// User friendly name of this network interface attachment's subnet.
	SubNetworkName       *string `json:"SubNetworkName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudNetworkInterfaceAttachmentAllOf CloudNetworkInterfaceAttachmentAllOf

// NewCloudNetworkInterfaceAttachmentAllOf instantiates a new CloudNetworkInterfaceAttachmentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudNetworkInterfaceAttachmentAllOf(classId string, objectType string) *CloudNetworkInterfaceAttachmentAllOf {
	this := CloudNetworkInterfaceAttachmentAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudNetworkInterfaceAttachmentAllOfWithDefaults instantiates a new CloudNetworkInterfaceAttachmentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudNetworkInterfaceAttachmentAllOfWithDefaults() *CloudNetworkInterfaceAttachmentAllOf {
	this := CloudNetworkInterfaceAttachmentAllOf{}
	var classId string = "cloud.NetworkInterfaceAttachment"
	this.ClassId = classId
	var objectType string = "cloud.NetworkInterfaceAttachment"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudNetworkInterfaceAttachmentAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudNetworkInterfaceAttachmentAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudNetworkInterfaceAttachmentAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudNetworkInterfaceAttachmentAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccessConfig returns the AccessConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudNetworkInterfaceAttachmentAllOf) GetAccessConfig() CloudNetworkAccessConfig {
	if o == nil || o.AccessConfig.Get() == nil {
		var ret CloudNetworkAccessConfig
		return ret
	}
	return *o.AccessConfig.Get()
}

// GetAccessConfigOk returns a tuple with the AccessConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudNetworkInterfaceAttachmentAllOf) GetAccessConfigOk() (*CloudNetworkAccessConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessConfig.Get(), o.AccessConfig.IsSet()
}

// HasAccessConfig returns a boolean if a field has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) HasAccessConfig() bool {
	if o != nil && o.AccessConfig.IsSet() {
		return true
	}

	return false
}

// SetAccessConfig gets a reference to the given NullableCloudNetworkAccessConfig and assigns it to the AccessConfig field.
func (o *CloudNetworkInterfaceAttachmentAllOf) SetAccessConfig(v CloudNetworkAccessConfig) {
	o.AccessConfig.Set(&v)
}

// SetAccessConfigNil sets the value for AccessConfig to be an explicit nil
func (o *CloudNetworkInterfaceAttachmentAllOf) SetAccessConfigNil() {
	o.AccessConfig.Set(nil)
}

// UnsetAccessConfig ensures that no value is present for AccessConfig, not even an explicit nil
func (o *CloudNetworkInterfaceAttachmentAllOf) UnsetAccessConfig() {
	o.AccessConfig.Unset()
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *CloudNetworkInterfaceAttachmentAllOf) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *CloudNetworkInterfaceAttachmentAllOf) SetIdentity(v string) {
	o.Identity = &v
}

// GetIpForwardingEnabled returns the IpForwardingEnabled field value if set, zero value otherwise.
func (o *CloudNetworkInterfaceAttachmentAllOf) GetIpForwardingEnabled() bool {
	if o == nil || o.IpForwardingEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IpForwardingEnabled
}

// GetIpForwardingEnabledOk returns a tuple with the IpForwardingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) GetIpForwardingEnabledOk() (*bool, bool) {
	if o == nil || o.IpForwardingEnabled == nil {
		return nil, false
	}
	return o.IpForwardingEnabled, true
}

// HasIpForwardingEnabled returns a boolean if a field has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) HasIpForwardingEnabled() bool {
	if o != nil && o.IpForwardingEnabled != nil {
		return true
	}

	return false
}

// SetIpForwardingEnabled gets a reference to the given bool and assigns it to the IpForwardingEnabled field.
func (o *CloudNetworkInterfaceAttachmentAllOf) SetIpForwardingEnabled(v bool) {
	o.IpForwardingEnabled = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *CloudNetworkInterfaceAttachmentAllOf) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *CloudNetworkInterfaceAttachmentAllOf) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *CloudNetworkInterfaceAttachmentAllOf) GetNetworkId() string {
	if o == nil || o.NetworkId == nil {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) GetNetworkIdOk() (*string, bool) {
	if o == nil || o.NetworkId == nil {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) HasNetworkId() bool {
	if o != nil && o.NetworkId != nil {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *CloudNetworkInterfaceAttachmentAllOf) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetNetworkName returns the NetworkName field value if set, zero value otherwise.
func (o *CloudNetworkInterfaceAttachmentAllOf) GetNetworkName() string {
	if o == nil || o.NetworkName == nil {
		var ret string
		return ret
	}
	return *o.NetworkName
}

// GetNetworkNameOk returns a tuple with the NetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) GetNetworkNameOk() (*string, bool) {
	if o == nil || o.NetworkName == nil {
		return nil, false
	}
	return o.NetworkName, true
}

// HasNetworkName returns a boolean if a field has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) HasNetworkName() bool {
	if o != nil && o.NetworkName != nil {
		return true
	}

	return false
}

// SetNetworkName gets a reference to the given string and assigns it to the NetworkName field.
func (o *CloudNetworkInterfaceAttachmentAllOf) SetNetworkName(v string) {
	o.NetworkName = &v
}

// GetNicIndex returns the NicIndex field value if set, zero value otherwise.
func (o *CloudNetworkInterfaceAttachmentAllOf) GetNicIndex() int64 {
	if o == nil || o.NicIndex == nil {
		var ret int64
		return ret
	}
	return *o.NicIndex
}

// GetNicIndexOk returns a tuple with the NicIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) GetNicIndexOk() (*int64, bool) {
	if o == nil || o.NicIndex == nil {
		return nil, false
	}
	return o.NicIndex, true
}

// HasNicIndex returns a boolean if a field has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) HasNicIndex() bool {
	if o != nil && o.NicIndex != nil {
		return true
	}

	return false
}

// SetNicIndex gets a reference to the given int64 and assigns it to the NicIndex field.
func (o *CloudNetworkInterfaceAttachmentAllOf) SetNicIndex(v int64) {
	o.NicIndex = &v
}

// GetPrivateAddress returns the PrivateAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudNetworkInterfaceAttachmentAllOf) GetPrivateAddress() []CloudNetworkAddress {
	if o == nil {
		var ret []CloudNetworkAddress
		return ret
	}
	return o.PrivateAddress
}

// GetPrivateAddressOk returns a tuple with the PrivateAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudNetworkInterfaceAttachmentAllOf) GetPrivateAddressOk() (*[]CloudNetworkAddress, bool) {
	if o == nil || o.PrivateAddress == nil {
		return nil, false
	}
	return &o.PrivateAddress, true
}

// HasPrivateAddress returns a boolean if a field has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) HasPrivateAddress() bool {
	if o != nil && o.PrivateAddress != nil {
		return true
	}

	return false
}

// SetPrivateAddress gets a reference to the given []CloudNetworkAddress and assigns it to the PrivateAddress field.
func (o *CloudNetworkInterfaceAttachmentAllOf) SetPrivateAddress(v []CloudNetworkAddress) {
	o.PrivateAddress = v
}

// GetPublicAddress returns the PublicAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudNetworkInterfaceAttachmentAllOf) GetPublicAddress() []CloudNetworkAddress {
	if o == nil {
		var ret []CloudNetworkAddress
		return ret
	}
	return o.PublicAddress
}

// GetPublicAddressOk returns a tuple with the PublicAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudNetworkInterfaceAttachmentAllOf) GetPublicAddressOk() (*[]CloudNetworkAddress, bool) {
	if o == nil || o.PublicAddress == nil {
		return nil, false
	}
	return &o.PublicAddress, true
}

// HasPublicAddress returns a boolean if a field has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) HasPublicAddress() bool {
	if o != nil && o.PublicAddress != nil {
		return true
	}

	return false
}

// SetPublicAddress gets a reference to the given []CloudNetworkAddress and assigns it to the PublicAddress field.
func (o *CloudNetworkInterfaceAttachmentAllOf) SetPublicAddress(v []CloudNetworkAddress) {
	o.PublicAddress = v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudNetworkInterfaceAttachmentAllOf) GetSecurityGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudNetworkInterfaceAttachmentAllOf) GetSecurityGroupsOk() (*[]string, bool) {
	if o == nil || o.SecurityGroups == nil {
		return nil, false
	}
	return &o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) HasSecurityGroups() bool {
	if o != nil && o.SecurityGroups != nil {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []string and assigns it to the SecurityGroups field.
func (o *CloudNetworkInterfaceAttachmentAllOf) SetSecurityGroups(v []string) {
	o.SecurityGroups = v
}

// GetSubNetworkId returns the SubNetworkId field value if set, zero value otherwise.
func (o *CloudNetworkInterfaceAttachmentAllOf) GetSubNetworkId() string {
	if o == nil || o.SubNetworkId == nil {
		var ret string
		return ret
	}
	return *o.SubNetworkId
}

// GetSubNetworkIdOk returns a tuple with the SubNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) GetSubNetworkIdOk() (*string, bool) {
	if o == nil || o.SubNetworkId == nil {
		return nil, false
	}
	return o.SubNetworkId, true
}

// HasSubNetworkId returns a boolean if a field has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) HasSubNetworkId() bool {
	if o != nil && o.SubNetworkId != nil {
		return true
	}

	return false
}

// SetSubNetworkId gets a reference to the given string and assigns it to the SubNetworkId field.
func (o *CloudNetworkInterfaceAttachmentAllOf) SetSubNetworkId(v string) {
	o.SubNetworkId = &v
}

// GetSubNetworkName returns the SubNetworkName field value if set, zero value otherwise.
func (o *CloudNetworkInterfaceAttachmentAllOf) GetSubNetworkName() string {
	if o == nil || o.SubNetworkName == nil {
		var ret string
		return ret
	}
	return *o.SubNetworkName
}

// GetSubNetworkNameOk returns a tuple with the SubNetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) GetSubNetworkNameOk() (*string, bool) {
	if o == nil || o.SubNetworkName == nil {
		return nil, false
	}
	return o.SubNetworkName, true
}

// HasSubNetworkName returns a boolean if a field has been set.
func (o *CloudNetworkInterfaceAttachmentAllOf) HasSubNetworkName() bool {
	if o != nil && o.SubNetworkName != nil {
		return true
	}

	return false
}

// SetSubNetworkName gets a reference to the given string and assigns it to the SubNetworkName field.
func (o *CloudNetworkInterfaceAttachmentAllOf) SetSubNetworkName(v string) {
	o.SubNetworkName = &v
}

func (o CloudNetworkInterfaceAttachmentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AccessConfig.IsSet() {
		toSerialize["AccessConfig"] = o.AccessConfig.Get()
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.IpForwardingEnabled != nil {
		toSerialize["IpForwardingEnabled"] = o.IpForwardingEnabled
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.NetworkId != nil {
		toSerialize["NetworkId"] = o.NetworkId
	}
	if o.NetworkName != nil {
		toSerialize["NetworkName"] = o.NetworkName
	}
	if o.NicIndex != nil {
		toSerialize["NicIndex"] = o.NicIndex
	}
	if o.PrivateAddress != nil {
		toSerialize["PrivateAddress"] = o.PrivateAddress
	}
	if o.PublicAddress != nil {
		toSerialize["PublicAddress"] = o.PublicAddress
	}
	if o.SecurityGroups != nil {
		toSerialize["SecurityGroups"] = o.SecurityGroups
	}
	if o.SubNetworkId != nil {
		toSerialize["SubNetworkId"] = o.SubNetworkId
	}
	if o.SubNetworkName != nil {
		toSerialize["SubNetworkName"] = o.SubNetworkName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudNetworkInterfaceAttachmentAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCloudNetworkInterfaceAttachmentAllOf := _CloudNetworkInterfaceAttachmentAllOf{}

	if err = json.Unmarshal(bytes, &varCloudNetworkInterfaceAttachmentAllOf); err == nil {
		*o = CloudNetworkInterfaceAttachmentAllOf(varCloudNetworkInterfaceAttachmentAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccessConfig")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "IpForwardingEnabled")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "NetworkId")
		delete(additionalProperties, "NetworkName")
		delete(additionalProperties, "NicIndex")
		delete(additionalProperties, "PrivateAddress")
		delete(additionalProperties, "PublicAddress")
		delete(additionalProperties, "SecurityGroups")
		delete(additionalProperties, "SubNetworkId")
		delete(additionalProperties, "SubNetworkName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudNetworkInterfaceAttachmentAllOf struct {
	value *CloudNetworkInterfaceAttachmentAllOf
	isSet bool
}

func (v NullableCloudNetworkInterfaceAttachmentAllOf) Get() *CloudNetworkInterfaceAttachmentAllOf {
	return v.value
}

func (v *NullableCloudNetworkInterfaceAttachmentAllOf) Set(val *CloudNetworkInterfaceAttachmentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudNetworkInterfaceAttachmentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudNetworkInterfaceAttachmentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudNetworkInterfaceAttachmentAllOf(val *CloudNetworkInterfaceAttachmentAllOf) *NullableCloudNetworkInterfaceAttachmentAllOf {
	return &NullableCloudNetworkInterfaceAttachmentAllOf{value: val, isSet: true}
}

func (v NullableCloudNetworkInterfaceAttachmentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudNetworkInterfaceAttachmentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
