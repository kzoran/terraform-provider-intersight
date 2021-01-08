/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-01-04T05:15:49Z.
 *
 * API version: 1.0.9-3144
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// IamLdapBaseProperties Base settings of LDAP required while configuring LDAP policy.
type IamLdapBaseProperties struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Role and locale information of the user.
	Attribute *string `json:"Attribute,omitempty"`
	// Base Distinguished Name (DN). Starting point from where server will search for users and groups.
	BaseDn *string `json:"BaseDn,omitempty"`
	// Distinguished Name (DN) of the user, that is used to authenticate against LDAP servers.
	BindDn *string `json:"BindDn,omitempty"`
	// Authentication method to access LDAP servers. * `LoginCredentials` - Requires the user credentials. If the bind process fails, then user is denied access. * `Anonymous` - Requires no username and password. If this option is selected and the LDAP server is configured for Anonymous logins, then the user gains access. * `ConfiguredCredentials` - Requires a known set of credentials to be specified for the initial bind process. If the initial bind process succeeds, then the distinguished name (DN) of the user name is queried and re-used for the re-binding process. If the re-binding process fails, then the user is denied access.
	BindMethod *string `json:"BindMethod,omitempty"`
	// The IPv4 domain that all users must be in.
	Domain *string `json:"Domain,omitempty"`
	// If enabled, the endpoint encrypts all information it sends to the LDAP server.
	EnableEncryption *bool `json:"EnableEncryption,omitempty"`
	// If enabled, user authorization is also done at the group level for LDAP users not in the local user database.
	EnableGroupAuthorization *bool `json:"EnableGroupAuthorization,omitempty"`
	// Criteria to identify entries in search requests.
	Filter *string `json:"Filter,omitempty"`
	// Groups to which an LDAP entry belongs.
	GroupAttribute *string `json:"GroupAttribute,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// Search depth to look for a nested LDAP group in an LDAP group map.
	NestedGroupSearchDepth *int64 `json:"NestedGroupSearchDepth,omitempty"`
	// The password of the user for initial bind process. It can be any string that adheres to the following constraints. It can have character except spaces, tabs, line breaks. It cannot be more than 254 characters.
	Password *string `json:"Password,omitempty"`
	// LDAP authentication timeout duration, in seconds.
	Timeout              *int64 `json:"Timeout,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamLdapBaseProperties IamLdapBaseProperties

// NewIamLdapBaseProperties instantiates a new IamLdapBaseProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamLdapBaseProperties(classId string, objectType string) *IamLdapBaseProperties {
	this := IamLdapBaseProperties{}
	this.ClassId = classId
	this.ObjectType = objectType
	var bindMethod string = "LoginCredentials"
	this.BindMethod = &bindMethod
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	var nestedGroupSearchDepth int64 = 128
	this.NestedGroupSearchDepth = &nestedGroupSearchDepth
	var timeout int64 = 0
	this.Timeout = &timeout
	return &this
}

// NewIamLdapBasePropertiesWithDefaults instantiates a new IamLdapBaseProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamLdapBasePropertiesWithDefaults() *IamLdapBaseProperties {
	this := IamLdapBaseProperties{}
	var classId string = "iam.LdapBaseProperties"
	this.ClassId = classId
	var objectType string = "iam.LdapBaseProperties"
	this.ObjectType = objectType
	var bindMethod string = "LoginCredentials"
	this.BindMethod = &bindMethod
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	var nestedGroupSearchDepth int64 = 128
	this.NestedGroupSearchDepth = &nestedGroupSearchDepth
	var timeout int64 = 0
	this.Timeout = &timeout
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamLdapBaseProperties) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamLdapBaseProperties) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamLdapBaseProperties) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamLdapBaseProperties) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamLdapBaseProperties) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamLdapBaseProperties) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *IamLdapBaseProperties) GetAttribute() string {
	if o == nil || o.Attribute == nil {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapBaseProperties) GetAttributeOk() (*string, bool) {
	if o == nil || o.Attribute == nil {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *IamLdapBaseProperties) HasAttribute() bool {
	if o != nil && o.Attribute != nil {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *IamLdapBaseProperties) SetAttribute(v string) {
	o.Attribute = &v
}

// GetBaseDn returns the BaseDn field value if set, zero value otherwise.
func (o *IamLdapBaseProperties) GetBaseDn() string {
	if o == nil || o.BaseDn == nil {
		var ret string
		return ret
	}
	return *o.BaseDn
}

// GetBaseDnOk returns a tuple with the BaseDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapBaseProperties) GetBaseDnOk() (*string, bool) {
	if o == nil || o.BaseDn == nil {
		return nil, false
	}
	return o.BaseDn, true
}

// HasBaseDn returns a boolean if a field has been set.
func (o *IamLdapBaseProperties) HasBaseDn() bool {
	if o != nil && o.BaseDn != nil {
		return true
	}

	return false
}

// SetBaseDn gets a reference to the given string and assigns it to the BaseDn field.
func (o *IamLdapBaseProperties) SetBaseDn(v string) {
	o.BaseDn = &v
}

// GetBindDn returns the BindDn field value if set, zero value otherwise.
func (o *IamLdapBaseProperties) GetBindDn() string {
	if o == nil || o.BindDn == nil {
		var ret string
		return ret
	}
	return *o.BindDn
}

// GetBindDnOk returns a tuple with the BindDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapBaseProperties) GetBindDnOk() (*string, bool) {
	if o == nil || o.BindDn == nil {
		return nil, false
	}
	return o.BindDn, true
}

// HasBindDn returns a boolean if a field has been set.
func (o *IamLdapBaseProperties) HasBindDn() bool {
	if o != nil && o.BindDn != nil {
		return true
	}

	return false
}

// SetBindDn gets a reference to the given string and assigns it to the BindDn field.
func (o *IamLdapBaseProperties) SetBindDn(v string) {
	o.BindDn = &v
}

// GetBindMethod returns the BindMethod field value if set, zero value otherwise.
func (o *IamLdapBaseProperties) GetBindMethod() string {
	if o == nil || o.BindMethod == nil {
		var ret string
		return ret
	}
	return *o.BindMethod
}

// GetBindMethodOk returns a tuple with the BindMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapBaseProperties) GetBindMethodOk() (*string, bool) {
	if o == nil || o.BindMethod == nil {
		return nil, false
	}
	return o.BindMethod, true
}

// HasBindMethod returns a boolean if a field has been set.
func (o *IamLdapBaseProperties) HasBindMethod() bool {
	if o != nil && o.BindMethod != nil {
		return true
	}

	return false
}

// SetBindMethod gets a reference to the given string and assigns it to the BindMethod field.
func (o *IamLdapBaseProperties) SetBindMethod(v string) {
	o.BindMethod = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *IamLdapBaseProperties) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapBaseProperties) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *IamLdapBaseProperties) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *IamLdapBaseProperties) SetDomain(v string) {
	o.Domain = &v
}

// GetEnableEncryption returns the EnableEncryption field value if set, zero value otherwise.
func (o *IamLdapBaseProperties) GetEnableEncryption() bool {
	if o == nil || o.EnableEncryption == nil {
		var ret bool
		return ret
	}
	return *o.EnableEncryption
}

// GetEnableEncryptionOk returns a tuple with the EnableEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapBaseProperties) GetEnableEncryptionOk() (*bool, bool) {
	if o == nil || o.EnableEncryption == nil {
		return nil, false
	}
	return o.EnableEncryption, true
}

// HasEnableEncryption returns a boolean if a field has been set.
func (o *IamLdapBaseProperties) HasEnableEncryption() bool {
	if o != nil && o.EnableEncryption != nil {
		return true
	}

	return false
}

// SetEnableEncryption gets a reference to the given bool and assigns it to the EnableEncryption field.
func (o *IamLdapBaseProperties) SetEnableEncryption(v bool) {
	o.EnableEncryption = &v
}

// GetEnableGroupAuthorization returns the EnableGroupAuthorization field value if set, zero value otherwise.
func (o *IamLdapBaseProperties) GetEnableGroupAuthorization() bool {
	if o == nil || o.EnableGroupAuthorization == nil {
		var ret bool
		return ret
	}
	return *o.EnableGroupAuthorization
}

// GetEnableGroupAuthorizationOk returns a tuple with the EnableGroupAuthorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapBaseProperties) GetEnableGroupAuthorizationOk() (*bool, bool) {
	if o == nil || o.EnableGroupAuthorization == nil {
		return nil, false
	}
	return o.EnableGroupAuthorization, true
}

// HasEnableGroupAuthorization returns a boolean if a field has been set.
func (o *IamLdapBaseProperties) HasEnableGroupAuthorization() bool {
	if o != nil && o.EnableGroupAuthorization != nil {
		return true
	}

	return false
}

// SetEnableGroupAuthorization gets a reference to the given bool and assigns it to the EnableGroupAuthorization field.
func (o *IamLdapBaseProperties) SetEnableGroupAuthorization(v bool) {
	o.EnableGroupAuthorization = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *IamLdapBaseProperties) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapBaseProperties) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *IamLdapBaseProperties) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *IamLdapBaseProperties) SetFilter(v string) {
	o.Filter = &v
}

// GetGroupAttribute returns the GroupAttribute field value if set, zero value otherwise.
func (o *IamLdapBaseProperties) GetGroupAttribute() string {
	if o == nil || o.GroupAttribute == nil {
		var ret string
		return ret
	}
	return *o.GroupAttribute
}

// GetGroupAttributeOk returns a tuple with the GroupAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapBaseProperties) GetGroupAttributeOk() (*string, bool) {
	if o == nil || o.GroupAttribute == nil {
		return nil, false
	}
	return o.GroupAttribute, true
}

// HasGroupAttribute returns a boolean if a field has been set.
func (o *IamLdapBaseProperties) HasGroupAttribute() bool {
	if o != nil && o.GroupAttribute != nil {
		return true
	}

	return false
}

// SetGroupAttribute gets a reference to the given string and assigns it to the GroupAttribute field.
func (o *IamLdapBaseProperties) SetGroupAttribute(v string) {
	o.GroupAttribute = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *IamLdapBaseProperties) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapBaseProperties) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *IamLdapBaseProperties) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *IamLdapBaseProperties) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetNestedGroupSearchDepth returns the NestedGroupSearchDepth field value if set, zero value otherwise.
func (o *IamLdapBaseProperties) GetNestedGroupSearchDepth() int64 {
	if o == nil || o.NestedGroupSearchDepth == nil {
		var ret int64
		return ret
	}
	return *o.NestedGroupSearchDepth
}

// GetNestedGroupSearchDepthOk returns a tuple with the NestedGroupSearchDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapBaseProperties) GetNestedGroupSearchDepthOk() (*int64, bool) {
	if o == nil || o.NestedGroupSearchDepth == nil {
		return nil, false
	}
	return o.NestedGroupSearchDepth, true
}

// HasNestedGroupSearchDepth returns a boolean if a field has been set.
func (o *IamLdapBaseProperties) HasNestedGroupSearchDepth() bool {
	if o != nil && o.NestedGroupSearchDepth != nil {
		return true
	}

	return false
}

// SetNestedGroupSearchDepth gets a reference to the given int64 and assigns it to the NestedGroupSearchDepth field.
func (o *IamLdapBaseProperties) SetNestedGroupSearchDepth(v int64) {
	o.NestedGroupSearchDepth = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *IamLdapBaseProperties) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapBaseProperties) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *IamLdapBaseProperties) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *IamLdapBaseProperties) SetPassword(v string) {
	o.Password = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *IamLdapBaseProperties) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapBaseProperties) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *IamLdapBaseProperties) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *IamLdapBaseProperties) SetTimeout(v int64) {
	o.Timeout = &v
}

func (o IamLdapBaseProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Attribute != nil {
		toSerialize["Attribute"] = o.Attribute
	}
	if o.BaseDn != nil {
		toSerialize["BaseDn"] = o.BaseDn
	}
	if o.BindDn != nil {
		toSerialize["BindDn"] = o.BindDn
	}
	if o.BindMethod != nil {
		toSerialize["BindMethod"] = o.BindMethod
	}
	if o.Domain != nil {
		toSerialize["Domain"] = o.Domain
	}
	if o.EnableEncryption != nil {
		toSerialize["EnableEncryption"] = o.EnableEncryption
	}
	if o.EnableGroupAuthorization != nil {
		toSerialize["EnableGroupAuthorization"] = o.EnableGroupAuthorization
	}
	if o.Filter != nil {
		toSerialize["Filter"] = o.Filter
	}
	if o.GroupAttribute != nil {
		toSerialize["GroupAttribute"] = o.GroupAttribute
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.NestedGroupSearchDepth != nil {
		toSerialize["NestedGroupSearchDepth"] = o.NestedGroupSearchDepth
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Timeout != nil {
		toSerialize["Timeout"] = o.Timeout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamLdapBaseProperties) UnmarshalJSON(bytes []byte) (err error) {
	type IamLdapBasePropertiesWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Role and locale information of the user.
		Attribute *string `json:"Attribute,omitempty"`
		// Base Distinguished Name (DN). Starting point from where server will search for users and groups.
		BaseDn *string `json:"BaseDn,omitempty"`
		// Distinguished Name (DN) of the user, that is used to authenticate against LDAP servers.
		BindDn *string `json:"BindDn,omitempty"`
		// Authentication method to access LDAP servers. * `LoginCredentials` - Requires the user credentials. If the bind process fails, then user is denied access. * `Anonymous` - Requires no username and password. If this option is selected and the LDAP server is configured for Anonymous logins, then the user gains access. * `ConfiguredCredentials` - Requires a known set of credentials to be specified for the initial bind process. If the initial bind process succeeds, then the distinguished name (DN) of the user name is queried and re-used for the re-binding process. If the re-binding process fails, then the user is denied access.
		BindMethod *string `json:"BindMethod,omitempty"`
		// The IPv4 domain that all users must be in.
		Domain *string `json:"Domain,omitempty"`
		// If enabled, the endpoint encrypts all information it sends to the LDAP server.
		EnableEncryption *bool `json:"EnableEncryption,omitempty"`
		// If enabled, user authorization is also done at the group level for LDAP users not in the local user database.
		EnableGroupAuthorization *bool `json:"EnableGroupAuthorization,omitempty"`
		// Criteria to identify entries in search requests.
		Filter *string `json:"Filter,omitempty"`
		// Groups to which an LDAP entry belongs.
		GroupAttribute *string `json:"GroupAttribute,omitempty"`
		// Indicates whether the value of the 'password' property has been set.
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
		// Search depth to look for a nested LDAP group in an LDAP group map.
		NestedGroupSearchDepth *int64 `json:"NestedGroupSearchDepth,omitempty"`
		// The password of the user for initial bind process. It can be any string that adheres to the following constraints. It can have character except spaces, tabs, line breaks. It cannot be more than 254 characters.
		Password *string `json:"Password,omitempty"`
		// LDAP authentication timeout duration, in seconds.
		Timeout *int64 `json:"Timeout,omitempty"`
	}

	varIamLdapBasePropertiesWithoutEmbeddedStruct := IamLdapBasePropertiesWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamLdapBasePropertiesWithoutEmbeddedStruct)
	if err == nil {
		varIamLdapBaseProperties := _IamLdapBaseProperties{}
		varIamLdapBaseProperties.ClassId = varIamLdapBasePropertiesWithoutEmbeddedStruct.ClassId
		varIamLdapBaseProperties.ObjectType = varIamLdapBasePropertiesWithoutEmbeddedStruct.ObjectType
		varIamLdapBaseProperties.Attribute = varIamLdapBasePropertiesWithoutEmbeddedStruct.Attribute
		varIamLdapBaseProperties.BaseDn = varIamLdapBasePropertiesWithoutEmbeddedStruct.BaseDn
		varIamLdapBaseProperties.BindDn = varIamLdapBasePropertiesWithoutEmbeddedStruct.BindDn
		varIamLdapBaseProperties.BindMethod = varIamLdapBasePropertiesWithoutEmbeddedStruct.BindMethod
		varIamLdapBaseProperties.Domain = varIamLdapBasePropertiesWithoutEmbeddedStruct.Domain
		varIamLdapBaseProperties.EnableEncryption = varIamLdapBasePropertiesWithoutEmbeddedStruct.EnableEncryption
		varIamLdapBaseProperties.EnableGroupAuthorization = varIamLdapBasePropertiesWithoutEmbeddedStruct.EnableGroupAuthorization
		varIamLdapBaseProperties.Filter = varIamLdapBasePropertiesWithoutEmbeddedStruct.Filter
		varIamLdapBaseProperties.GroupAttribute = varIamLdapBasePropertiesWithoutEmbeddedStruct.GroupAttribute
		varIamLdapBaseProperties.IsPasswordSet = varIamLdapBasePropertiesWithoutEmbeddedStruct.IsPasswordSet
		varIamLdapBaseProperties.NestedGroupSearchDepth = varIamLdapBasePropertiesWithoutEmbeddedStruct.NestedGroupSearchDepth
		varIamLdapBaseProperties.Password = varIamLdapBasePropertiesWithoutEmbeddedStruct.Password
		varIamLdapBaseProperties.Timeout = varIamLdapBasePropertiesWithoutEmbeddedStruct.Timeout
		*o = IamLdapBaseProperties(varIamLdapBaseProperties)
	} else {
		return err
	}

	varIamLdapBaseProperties := _IamLdapBaseProperties{}

	err = json.Unmarshal(bytes, &varIamLdapBaseProperties)
	if err == nil {
		o.MoBaseComplexType = varIamLdapBaseProperties.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Attribute")
		delete(additionalProperties, "BaseDn")
		delete(additionalProperties, "BindDn")
		delete(additionalProperties, "BindMethod")
		delete(additionalProperties, "Domain")
		delete(additionalProperties, "EnableEncryption")
		delete(additionalProperties, "EnableGroupAuthorization")
		delete(additionalProperties, "Filter")
		delete(additionalProperties, "GroupAttribute")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "NestedGroupSearchDepth")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Timeout")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableIamLdapBaseProperties struct {
	value *IamLdapBaseProperties
	isSet bool
}

func (v NullableIamLdapBaseProperties) Get() *IamLdapBaseProperties {
	return v.value
}

func (v *NullableIamLdapBaseProperties) Set(val *IamLdapBaseProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableIamLdapBaseProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableIamLdapBaseProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamLdapBaseProperties(val *IamLdapBaseProperties) *NullableIamLdapBaseProperties {
	return &NullableIamLdapBaseProperties{value: val, isSet: true}
}

func (v NullableIamLdapBaseProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamLdapBaseProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
