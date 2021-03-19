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

// ConnectorSshConfigAllOf Definition of the list of properties defined in 'connector.SshConfig', excluding properties defined in parent classes.
type ConnectorSshConfigAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A jump host for establishing a connection to a server. Plugin will first establish a connection to this server, then create a tunneled connection to the target host.
	JumpHost *string `json:"JumpHost,omitempty"`
	// Password to use in the connection credentials (If empty the private key will be used).
	Password *string `json:"Password,omitempty"`
	// The private key to use in the connection credentials (Optional if password is given).
	Pkey *string `json:"Pkey,omitempty"`
	// The remote server to connect to.
	Target *string `json:"Target,omitempty"`
	// Username for the remote connection.
	User                 *string `json:"User,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorSshConfigAllOf ConnectorSshConfigAllOf

// NewConnectorSshConfigAllOf instantiates a new ConnectorSshConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorSshConfigAllOf(classId string, objectType string) *ConnectorSshConfigAllOf {
	this := ConnectorSshConfigAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorSshConfigAllOfWithDefaults instantiates a new ConnectorSshConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorSshConfigAllOfWithDefaults() *ConnectorSshConfigAllOf {
	this := ConnectorSshConfigAllOf{}
	var classId string = "connector.SshConfig"
	this.ClassId = classId
	var objectType string = "connector.SshConfig"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorSshConfigAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorSshConfigAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorSshConfigAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorSshConfigAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorSshConfigAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorSshConfigAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetJumpHost returns the JumpHost field value if set, zero value otherwise.
func (o *ConnectorSshConfigAllOf) GetJumpHost() string {
	if o == nil || o.JumpHost == nil {
		var ret string
		return ret
	}
	return *o.JumpHost
}

// GetJumpHostOk returns a tuple with the JumpHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshConfigAllOf) GetJumpHostOk() (*string, bool) {
	if o == nil || o.JumpHost == nil {
		return nil, false
	}
	return o.JumpHost, true
}

// HasJumpHost returns a boolean if a field has been set.
func (o *ConnectorSshConfigAllOf) HasJumpHost() bool {
	if o != nil && o.JumpHost != nil {
		return true
	}

	return false
}

// SetJumpHost gets a reference to the given string and assigns it to the JumpHost field.
func (o *ConnectorSshConfigAllOf) SetJumpHost(v string) {
	o.JumpHost = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ConnectorSshConfigAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshConfigAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ConnectorSshConfigAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ConnectorSshConfigAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetPkey returns the Pkey field value if set, zero value otherwise.
func (o *ConnectorSshConfigAllOf) GetPkey() string {
	if o == nil || o.Pkey == nil {
		var ret string
		return ret
	}
	return *o.Pkey
}

// GetPkeyOk returns a tuple with the Pkey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshConfigAllOf) GetPkeyOk() (*string, bool) {
	if o == nil || o.Pkey == nil {
		return nil, false
	}
	return o.Pkey, true
}

// HasPkey returns a boolean if a field has been set.
func (o *ConnectorSshConfigAllOf) HasPkey() bool {
	if o != nil && o.Pkey != nil {
		return true
	}

	return false
}

// SetPkey gets a reference to the given string and assigns it to the Pkey field.
func (o *ConnectorSshConfigAllOf) SetPkey(v string) {
	o.Pkey = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ConnectorSshConfigAllOf) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshConfigAllOf) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ConnectorSshConfigAllOf) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *ConnectorSshConfigAllOf) SetTarget(v string) {
	o.Target = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ConnectorSshConfigAllOf) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshConfigAllOf) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ConnectorSshConfigAllOf) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *ConnectorSshConfigAllOf) SetUser(v string) {
	o.User = &v
}

func (o ConnectorSshConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.JumpHost != nil {
		toSerialize["JumpHost"] = o.JumpHost
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Pkey != nil {
		toSerialize["Pkey"] = o.Pkey
	}
	if o.Target != nil {
		toSerialize["Target"] = o.Target
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorSshConfigAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConnectorSshConfigAllOf := _ConnectorSshConfigAllOf{}

	if err = json.Unmarshal(bytes, &varConnectorSshConfigAllOf); err == nil {
		*o = ConnectorSshConfigAllOf(varConnectorSshConfigAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "JumpHost")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Pkey")
		delete(additionalProperties, "Target")
		delete(additionalProperties, "User")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorSshConfigAllOf struct {
	value *ConnectorSshConfigAllOf
	isSet bool
}

func (v NullableConnectorSshConfigAllOf) Get() *ConnectorSshConfigAllOf {
	return v.value
}

func (v *NullableConnectorSshConfigAllOf) Set(val *ConnectorSshConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorSshConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorSshConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorSshConfigAllOf(val *ConnectorSshConfigAllOf) *NullableConnectorSshConfigAllOf {
	return &NullableConnectorSshConfigAllOf{value: val, isSet: true}
}

func (v NullableConnectorSshConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorSshConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
