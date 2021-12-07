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

// TechsupportmanagementCollectionControlPolicyAllOf Definition of the list of properties defined in 'techsupportmanagement.CollectionControlPolicy', excluding properties defined in parent classes.
type TechsupportmanagementCollectionControlPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Deployment type defines whether the policy is associated with a SaaS or Appliance account. * `None` - Service deployment type None. * `SaaS` - Service deployment type SaaS. * `Appliance` - Service deployment type Appliance.
	DeploymentType *string `json:"DeploymentType,omitempty"`
	// Enable or Disable techsupport collection for a specific account. * `Enable` - Enable techsupport collection. * `Disable` - Disable techsupport collection.
	TechSupportCollection *string                 `json:"TechSupportCollection,omitempty"`
	Account               *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _TechsupportmanagementCollectionControlPolicyAllOf TechsupportmanagementCollectionControlPolicyAllOf

// NewTechsupportmanagementCollectionControlPolicyAllOf instantiates a new TechsupportmanagementCollectionControlPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechsupportmanagementCollectionControlPolicyAllOf(classId string, objectType string) *TechsupportmanagementCollectionControlPolicyAllOf {
	this := TechsupportmanagementCollectionControlPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var techSupportCollection string = "Enable"
	this.TechSupportCollection = &techSupportCollection
	return &this
}

// NewTechsupportmanagementCollectionControlPolicyAllOfWithDefaults instantiates a new TechsupportmanagementCollectionControlPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechsupportmanagementCollectionControlPolicyAllOfWithDefaults() *TechsupportmanagementCollectionControlPolicyAllOf {
	this := TechsupportmanagementCollectionControlPolicyAllOf{}
	var classId string = "techsupportmanagement.CollectionControlPolicy"
	this.ClassId = classId
	var objectType string = "techsupportmanagement.CollectionControlPolicy"
	this.ObjectType = objectType
	var techSupportCollection string = "Enable"
	this.TechSupportCollection = &techSupportCollection
	return &this
}

// GetClassId returns the ClassId field value
func (o *TechsupportmanagementCollectionControlPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementCollectionControlPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TechsupportmanagementCollectionControlPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TechsupportmanagementCollectionControlPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementCollectionControlPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TechsupportmanagementCollectionControlPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeploymentType returns the DeploymentType field value if set, zero value otherwise.
func (o *TechsupportmanagementCollectionControlPolicyAllOf) GetDeploymentType() string {
	if o == nil || o.DeploymentType == nil {
		var ret string
		return ret
	}
	return *o.DeploymentType
}

// GetDeploymentTypeOk returns a tuple with the DeploymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementCollectionControlPolicyAllOf) GetDeploymentTypeOk() (*string, bool) {
	if o == nil || o.DeploymentType == nil {
		return nil, false
	}
	return o.DeploymentType, true
}

// HasDeploymentType returns a boolean if a field has been set.
func (o *TechsupportmanagementCollectionControlPolicyAllOf) HasDeploymentType() bool {
	if o != nil && o.DeploymentType != nil {
		return true
	}

	return false
}

// SetDeploymentType gets a reference to the given string and assigns it to the DeploymentType field.
func (o *TechsupportmanagementCollectionControlPolicyAllOf) SetDeploymentType(v string) {
	o.DeploymentType = &v
}

// GetTechSupportCollection returns the TechSupportCollection field value if set, zero value otherwise.
func (o *TechsupportmanagementCollectionControlPolicyAllOf) GetTechSupportCollection() string {
	if o == nil || o.TechSupportCollection == nil {
		var ret string
		return ret
	}
	return *o.TechSupportCollection
}

// GetTechSupportCollectionOk returns a tuple with the TechSupportCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementCollectionControlPolicyAllOf) GetTechSupportCollectionOk() (*string, bool) {
	if o == nil || o.TechSupportCollection == nil {
		return nil, false
	}
	return o.TechSupportCollection, true
}

// HasTechSupportCollection returns a boolean if a field has been set.
func (o *TechsupportmanagementCollectionControlPolicyAllOf) HasTechSupportCollection() bool {
	if o != nil && o.TechSupportCollection != nil {
		return true
	}

	return false
}

// SetTechSupportCollection gets a reference to the given string and assigns it to the TechSupportCollection field.
func (o *TechsupportmanagementCollectionControlPolicyAllOf) SetTechSupportCollection(v string) {
	o.TechSupportCollection = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *TechsupportmanagementCollectionControlPolicyAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementCollectionControlPolicyAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *TechsupportmanagementCollectionControlPolicyAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *TechsupportmanagementCollectionControlPolicyAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o TechsupportmanagementCollectionControlPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DeploymentType != nil {
		toSerialize["DeploymentType"] = o.DeploymentType
	}
	if o.TechSupportCollection != nil {
		toSerialize["TechSupportCollection"] = o.TechSupportCollection
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TechsupportmanagementCollectionControlPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTechsupportmanagementCollectionControlPolicyAllOf := _TechsupportmanagementCollectionControlPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varTechsupportmanagementCollectionControlPolicyAllOf); err == nil {
		*o = TechsupportmanagementCollectionControlPolicyAllOf(varTechsupportmanagementCollectionControlPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeploymentType")
		delete(additionalProperties, "TechSupportCollection")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTechsupportmanagementCollectionControlPolicyAllOf struct {
	value *TechsupportmanagementCollectionControlPolicyAllOf
	isSet bool
}

func (v NullableTechsupportmanagementCollectionControlPolicyAllOf) Get() *TechsupportmanagementCollectionControlPolicyAllOf {
	return v.value
}

func (v *NullableTechsupportmanagementCollectionControlPolicyAllOf) Set(val *TechsupportmanagementCollectionControlPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTechsupportmanagementCollectionControlPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTechsupportmanagementCollectionControlPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechsupportmanagementCollectionControlPolicyAllOf(val *TechsupportmanagementCollectionControlPolicyAllOf) *NullableTechsupportmanagementCollectionControlPolicyAllOf {
	return &NullableTechsupportmanagementCollectionControlPolicyAllOf{value: val, isSet: true}
}

func (v NullableTechsupportmanagementCollectionControlPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechsupportmanagementCollectionControlPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
