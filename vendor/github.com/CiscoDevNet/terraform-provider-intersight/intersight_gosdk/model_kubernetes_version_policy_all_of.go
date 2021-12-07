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

// KubernetesVersionPolicyAllOf Definition of the list of properties defined in 'kubernetes.VersionPolicy', excluding properties defined in parent classes.
type KubernetesVersionPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType   string                                `json:"ObjectType"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to kubernetesNodeGroupProfile resources.
	Profiles             []KubernetesNodeGroupProfileRelationship `json:"Profiles,omitempty"`
	Version              *KubernetesVersionRelationship           `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesVersionPolicyAllOf KubernetesVersionPolicyAllOf

// NewKubernetesVersionPolicyAllOf instantiates a new KubernetesVersionPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesVersionPolicyAllOf(classId string, objectType string) *KubernetesVersionPolicyAllOf {
	this := KubernetesVersionPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesVersionPolicyAllOfWithDefaults instantiates a new KubernetesVersionPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesVersionPolicyAllOfWithDefaults() *KubernetesVersionPolicyAllOf {
	this := KubernetesVersionPolicyAllOf{}
	var classId string = "kubernetes.VersionPolicy"
	this.ClassId = classId
	var objectType string = "kubernetes.VersionPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesVersionPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesVersionPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesVersionPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesVersionPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesVersionPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesVersionPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesVersionPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVersionPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesVersionPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesVersionPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesVersionPolicyAllOf) GetProfiles() []KubernetesNodeGroupProfileRelationship {
	if o == nil {
		var ret []KubernetesNodeGroupProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesVersionPolicyAllOf) GetProfilesOk() (*[]KubernetesNodeGroupProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *KubernetesVersionPolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []KubernetesNodeGroupProfileRelationship and assigns it to the Profiles field.
func (o *KubernetesVersionPolicyAllOf) SetProfiles(v []KubernetesNodeGroupProfileRelationship) {
	o.Profiles = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *KubernetesVersionPolicyAllOf) GetVersion() KubernetesVersionRelationship {
	if o == nil || o.Version == nil {
		var ret KubernetesVersionRelationship
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVersionPolicyAllOf) GetVersionOk() (*KubernetesVersionRelationship, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *KubernetesVersionPolicyAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given KubernetesVersionRelationship and assigns it to the Version field.
func (o *KubernetesVersionPolicyAllOf) SetVersion(v KubernetesVersionRelationship) {
	o.Version = &v
}

func (o KubernetesVersionPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesVersionPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesVersionPolicyAllOf := _KubernetesVersionPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesVersionPolicyAllOf); err == nil {
		*o = KubernetesVersionPolicyAllOf(varKubernetesVersionPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesVersionPolicyAllOf struct {
	value *KubernetesVersionPolicyAllOf
	isSet bool
}

func (v NullableKubernetesVersionPolicyAllOf) Get() *KubernetesVersionPolicyAllOf {
	return v.value
}

func (v *NullableKubernetesVersionPolicyAllOf) Set(val *KubernetesVersionPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesVersionPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesVersionPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesVersionPolicyAllOf(val *KubernetesVersionPolicyAllOf) *NullableKubernetesVersionPolicyAllOf {
	return &NullableKubernetesVersionPolicyAllOf{value: val, isSet: true}
}

func (v NullableKubernetesVersionPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesVersionPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
