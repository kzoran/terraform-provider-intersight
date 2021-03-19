/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-03-10T06:51:24Z.
 *
 * API version: 1.0.9-3942
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// SdwanVmanageAccountPolicyAllOf Definition of the list of properties defined in 'sdwan.VmanageAccountPolicy', excluding properties defined in parent classes.
type SdwanVmanageAccountPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// VManage server hostname (FQDN) that the acccount holds information for.
	EndpointAddress *string `json:"EndpointAddress,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// Local password for authenticating with the vManage server.
	Password *string `json:"Password,omitempty"`
	// VManage Port number on which the application is running.
	Port *int64 `json:"Port,omitempty"`
	// Local username for authenticating with the vManage server.
	Username     *string                               `json:"Username,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to sdwanProfile resources.
	Profiles             []SdwanProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SdwanVmanageAccountPolicyAllOf SdwanVmanageAccountPolicyAllOf

// NewSdwanVmanageAccountPolicyAllOf instantiates a new SdwanVmanageAccountPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdwanVmanageAccountPolicyAllOf(classId string, objectType string) *SdwanVmanageAccountPolicyAllOf {
	this := SdwanVmanageAccountPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	var port int64 = 8443
	this.Port = &port
	return &this
}

// NewSdwanVmanageAccountPolicyAllOfWithDefaults instantiates a new SdwanVmanageAccountPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdwanVmanageAccountPolicyAllOfWithDefaults() *SdwanVmanageAccountPolicyAllOf {
	this := SdwanVmanageAccountPolicyAllOf{}
	var classId string = "sdwan.VmanageAccountPolicy"
	this.ClassId = classId
	var objectType string = "sdwan.VmanageAccountPolicy"
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	var port int64 = 8443
	this.Port = &port
	return &this
}

// GetClassId returns the ClassId field value
func (o *SdwanVmanageAccountPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SdwanVmanageAccountPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SdwanVmanageAccountPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SdwanVmanageAccountPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SdwanVmanageAccountPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SdwanVmanageAccountPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEndpointAddress returns the EndpointAddress field value if set, zero value otherwise.
func (o *SdwanVmanageAccountPolicyAllOf) GetEndpointAddress() string {
	if o == nil || o.EndpointAddress == nil {
		var ret string
		return ret
	}
	return *o.EndpointAddress
}

// GetEndpointAddressOk returns a tuple with the EndpointAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanVmanageAccountPolicyAllOf) GetEndpointAddressOk() (*string, bool) {
	if o == nil || o.EndpointAddress == nil {
		return nil, false
	}
	return o.EndpointAddress, true
}

// HasEndpointAddress returns a boolean if a field has been set.
func (o *SdwanVmanageAccountPolicyAllOf) HasEndpointAddress() bool {
	if o != nil && o.EndpointAddress != nil {
		return true
	}

	return false
}

// SetEndpointAddress gets a reference to the given string and assigns it to the EndpointAddress field.
func (o *SdwanVmanageAccountPolicyAllOf) SetEndpointAddress(v string) {
	o.EndpointAddress = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *SdwanVmanageAccountPolicyAllOf) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanVmanageAccountPolicyAllOf) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *SdwanVmanageAccountPolicyAllOf) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *SdwanVmanageAccountPolicyAllOf) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SdwanVmanageAccountPolicyAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanVmanageAccountPolicyAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SdwanVmanageAccountPolicyAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SdwanVmanageAccountPolicyAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SdwanVmanageAccountPolicyAllOf) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanVmanageAccountPolicyAllOf) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *SdwanVmanageAccountPolicyAllOf) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *SdwanVmanageAccountPolicyAllOf) SetPort(v int64) {
	o.Port = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SdwanVmanageAccountPolicyAllOf) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanVmanageAccountPolicyAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SdwanVmanageAccountPolicyAllOf) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SdwanVmanageAccountPolicyAllOf) SetUsername(v string) {
	o.Username = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *SdwanVmanageAccountPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanVmanageAccountPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *SdwanVmanageAccountPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *SdwanVmanageAccountPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SdwanVmanageAccountPolicyAllOf) GetProfiles() []SdwanProfileRelationship {
	if o == nil {
		var ret []SdwanProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SdwanVmanageAccountPolicyAllOf) GetProfilesOk() (*[]SdwanProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *SdwanVmanageAccountPolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []SdwanProfileRelationship and assigns it to the Profiles field.
func (o *SdwanVmanageAccountPolicyAllOf) SetProfiles(v []SdwanProfileRelationship) {
	o.Profiles = v
}

func (o SdwanVmanageAccountPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EndpointAddress != nil {
		toSerialize["EndpointAddress"] = o.EndpointAddress
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
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

func (o *SdwanVmanageAccountPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSdwanVmanageAccountPolicyAllOf := _SdwanVmanageAccountPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varSdwanVmanageAccountPolicyAllOf); err == nil {
		*o = SdwanVmanageAccountPolicyAllOf(varSdwanVmanageAccountPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EndpointAddress")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Username")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSdwanVmanageAccountPolicyAllOf struct {
	value *SdwanVmanageAccountPolicyAllOf
	isSet bool
}

func (v NullableSdwanVmanageAccountPolicyAllOf) Get() *SdwanVmanageAccountPolicyAllOf {
	return v.value
}

func (v *NullableSdwanVmanageAccountPolicyAllOf) Set(val *SdwanVmanageAccountPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSdwanVmanageAccountPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSdwanVmanageAccountPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdwanVmanageAccountPolicyAllOf(val *SdwanVmanageAccountPolicyAllOf) *NullableSdwanVmanageAccountPolicyAllOf {
	return &NullableSdwanVmanageAccountPolicyAllOf{value: val, isSet: true}
}

func (v NullableSdwanVmanageAccountPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdwanVmanageAccountPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
