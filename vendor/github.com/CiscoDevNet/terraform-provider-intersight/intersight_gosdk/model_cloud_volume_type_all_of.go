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

// CloudVolumeTypeAllOf Definition of the list of properties defined in 'cloud.VolumeType', excluding properties defined in parent classes.
type CloudVolumeTypeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The type of the volume.Types vary from cloud to cloud.
	Name *string `json:"Name,omitempty"`
	// Unique identity of the volume type, assigned by the cloud provider.
	VolumeTypeId         *string `json:"VolumeTypeId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudVolumeTypeAllOf CloudVolumeTypeAllOf

// NewCloudVolumeTypeAllOf instantiates a new CloudVolumeTypeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudVolumeTypeAllOf(classId string, objectType string) *CloudVolumeTypeAllOf {
	this := CloudVolumeTypeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudVolumeTypeAllOfWithDefaults instantiates a new CloudVolumeTypeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudVolumeTypeAllOfWithDefaults() *CloudVolumeTypeAllOf {
	this := CloudVolumeTypeAllOf{}
	var classId string = "cloud.VolumeType"
	this.ClassId = classId
	var objectType string = "cloud.VolumeType"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudVolumeTypeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudVolumeTypeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudVolumeTypeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudVolumeTypeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudVolumeTypeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudVolumeTypeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudVolumeTypeAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeTypeAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudVolumeTypeAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudVolumeTypeAllOf) SetName(v string) {
	o.Name = &v
}

// GetVolumeTypeId returns the VolumeTypeId field value if set, zero value otherwise.
func (o *CloudVolumeTypeAllOf) GetVolumeTypeId() string {
	if o == nil || o.VolumeTypeId == nil {
		var ret string
		return ret
	}
	return *o.VolumeTypeId
}

// GetVolumeTypeIdOk returns a tuple with the VolumeTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeTypeAllOf) GetVolumeTypeIdOk() (*string, bool) {
	if o == nil || o.VolumeTypeId == nil {
		return nil, false
	}
	return o.VolumeTypeId, true
}

// HasVolumeTypeId returns a boolean if a field has been set.
func (o *CloudVolumeTypeAllOf) HasVolumeTypeId() bool {
	if o != nil && o.VolumeTypeId != nil {
		return true
	}

	return false
}

// SetVolumeTypeId gets a reference to the given string and assigns it to the VolumeTypeId field.
func (o *CloudVolumeTypeAllOf) SetVolumeTypeId(v string) {
	o.VolumeTypeId = &v
}

func (o CloudVolumeTypeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.VolumeTypeId != nil {
		toSerialize["VolumeTypeId"] = o.VolumeTypeId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudVolumeTypeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCloudVolumeTypeAllOf := _CloudVolumeTypeAllOf{}

	if err = json.Unmarshal(bytes, &varCloudVolumeTypeAllOf); err == nil {
		*o = CloudVolumeTypeAllOf(varCloudVolumeTypeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "VolumeTypeId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudVolumeTypeAllOf struct {
	value *CloudVolumeTypeAllOf
	isSet bool
}

func (v NullableCloudVolumeTypeAllOf) Get() *CloudVolumeTypeAllOf {
	return v.value
}

func (v *NullableCloudVolumeTypeAllOf) Set(val *CloudVolumeTypeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudVolumeTypeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudVolumeTypeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudVolumeTypeAllOf(val *CloudVolumeTypeAllOf) *NullableCloudVolumeTypeAllOf {
	return &NullableCloudVolumeTypeAllOf{value: val, isSet: true}
}

func (v NullableCloudVolumeTypeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudVolumeTypeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
