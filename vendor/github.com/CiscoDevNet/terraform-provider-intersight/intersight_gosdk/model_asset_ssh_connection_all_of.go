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
)

// AssetSshConnectionAllOf Definition of the list of properties defined in 'asset.SshConnection', excluding properties defined in parent classes.
type AssetSshConnectionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The DNS hostname or IP Address, either IPv4 or IPv6, to be used to connect to the managed target.
	ManagementAddress *string `json:"ManagementAddress,omitempty"`
	// The port number to be used to connect to the managed target. Valid values are 1 - 65535. If not provided, a default port of 22 is used to establish the SSH connection to the given target.
	Port                 *int64 `json:"Port,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetSshConnectionAllOf AssetSshConnectionAllOf

// NewAssetSshConnectionAllOf instantiates a new AssetSshConnectionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetSshConnectionAllOf(classId string, objectType string) *AssetSshConnectionAllOf {
	this := AssetSshConnectionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetSshConnectionAllOfWithDefaults instantiates a new AssetSshConnectionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetSshConnectionAllOfWithDefaults() *AssetSshConnectionAllOf {
	this := AssetSshConnectionAllOf{}
	var classId string = "asset.SshConnection"
	this.ClassId = classId
	var objectType string = "asset.SshConnection"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetSshConnectionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetSshConnectionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetSshConnectionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetSshConnectionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetSshConnectionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetSshConnectionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetManagementAddress returns the ManagementAddress field value if set, zero value otherwise.
func (o *AssetSshConnectionAllOf) GetManagementAddress() string {
	if o == nil || o.ManagementAddress == nil {
		var ret string
		return ret
	}
	return *o.ManagementAddress
}

// GetManagementAddressOk returns a tuple with the ManagementAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSshConnectionAllOf) GetManagementAddressOk() (*string, bool) {
	if o == nil || o.ManagementAddress == nil {
		return nil, false
	}
	return o.ManagementAddress, true
}

// HasManagementAddress returns a boolean if a field has been set.
func (o *AssetSshConnectionAllOf) HasManagementAddress() bool {
	if o != nil && o.ManagementAddress != nil {
		return true
	}

	return false
}

// SetManagementAddress gets a reference to the given string and assigns it to the ManagementAddress field.
func (o *AssetSshConnectionAllOf) SetManagementAddress(v string) {
	o.ManagementAddress = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *AssetSshConnectionAllOf) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSshConnectionAllOf) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *AssetSshConnectionAllOf) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *AssetSshConnectionAllOf) SetPort(v int64) {
	o.Port = &v
}

func (o AssetSshConnectionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ManagementAddress != nil {
		toSerialize["ManagementAddress"] = o.ManagementAddress
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetSshConnectionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetSshConnectionAllOf := _AssetSshConnectionAllOf{}

	if err = json.Unmarshal(bytes, &varAssetSshConnectionAllOf); err == nil {
		*o = AssetSshConnectionAllOf(varAssetSshConnectionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ManagementAddress")
		delete(additionalProperties, "Port")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetSshConnectionAllOf struct {
	value *AssetSshConnectionAllOf
	isSet bool
}

func (v NullableAssetSshConnectionAllOf) Get() *AssetSshConnectionAllOf {
	return v.value
}

func (v *NullableAssetSshConnectionAllOf) Set(val *AssetSshConnectionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetSshConnectionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetSshConnectionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetSshConnectionAllOf(val *AssetSshConnectionAllOf) *NullableAssetSshConnectionAllOf {
	return &NullableAssetSshConnectionAllOf{value: val, isSet: true}
}

func (v NullableAssetSshConnectionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetSshConnectionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
