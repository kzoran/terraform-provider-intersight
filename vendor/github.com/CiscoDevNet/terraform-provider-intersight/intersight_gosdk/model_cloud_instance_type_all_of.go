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

// CloudInstanceTypeAllOf Definition of the list of properties defined in 'cloud.InstanceType', excluding properties defined in parent classes.
type CloudInstanceTypeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Cpu architecture, for example, x86. * `x86_64` - The virtual machine cpu architecture type x86_64. * `x86` - The virtual machine cpu architecture type x86.
	Architecture *string `json:"Architecture,omitempty"`
	// The speed of the processor, in GHz.
	CpuSpeed *int64 `json:"CpuSpeed,omitempty"`
	// The number of CPUs for the instance type.
	Cpus *int64 `json:"Cpus,omitempty"`
	// The ID of the instance type.
	InstanceTypeId *string `json:"InstanceTypeId,omitempty"`
	// The size of the memory, in bytes.
	Memory *int64 `json:"Memory,omitempty"`
	// Name to identity the instance type.
	Name *string `json:"Name,omitempty"`
	// Operation System, for example, Linux.
	Platform             *string `json:"Platform,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudInstanceTypeAllOf CloudInstanceTypeAllOf

// NewCloudInstanceTypeAllOf instantiates a new CloudInstanceTypeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudInstanceTypeAllOf(classId string, objectType string) *CloudInstanceTypeAllOf {
	this := CloudInstanceTypeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudInstanceTypeAllOfWithDefaults instantiates a new CloudInstanceTypeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudInstanceTypeAllOfWithDefaults() *CloudInstanceTypeAllOf {
	this := CloudInstanceTypeAllOf{}
	var classId string = "cloud.InstanceType"
	this.ClassId = classId
	var objectType string = "cloud.InstanceType"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudInstanceTypeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudInstanceTypeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudInstanceTypeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudInstanceTypeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudInstanceTypeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudInstanceTypeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *CloudInstanceTypeAllOf) GetArchitecture() string {
	if o == nil || o.Architecture == nil {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInstanceTypeAllOf) GetArchitectureOk() (*string, bool) {
	if o == nil || o.Architecture == nil {
		return nil, false
	}
	return o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *CloudInstanceTypeAllOf) HasArchitecture() bool {
	if o != nil && o.Architecture != nil {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *CloudInstanceTypeAllOf) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetCpuSpeed returns the CpuSpeed field value if set, zero value otherwise.
func (o *CloudInstanceTypeAllOf) GetCpuSpeed() int64 {
	if o == nil || o.CpuSpeed == nil {
		var ret int64
		return ret
	}
	return *o.CpuSpeed
}

// GetCpuSpeedOk returns a tuple with the CpuSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInstanceTypeAllOf) GetCpuSpeedOk() (*int64, bool) {
	if o == nil || o.CpuSpeed == nil {
		return nil, false
	}
	return o.CpuSpeed, true
}

// HasCpuSpeed returns a boolean if a field has been set.
func (o *CloudInstanceTypeAllOf) HasCpuSpeed() bool {
	if o != nil && o.CpuSpeed != nil {
		return true
	}

	return false
}

// SetCpuSpeed gets a reference to the given int64 and assigns it to the CpuSpeed field.
func (o *CloudInstanceTypeAllOf) SetCpuSpeed(v int64) {
	o.CpuSpeed = &v
}

// GetCpus returns the Cpus field value if set, zero value otherwise.
func (o *CloudInstanceTypeAllOf) GetCpus() int64 {
	if o == nil || o.Cpus == nil {
		var ret int64
		return ret
	}
	return *o.Cpus
}

// GetCpusOk returns a tuple with the Cpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInstanceTypeAllOf) GetCpusOk() (*int64, bool) {
	if o == nil || o.Cpus == nil {
		return nil, false
	}
	return o.Cpus, true
}

// HasCpus returns a boolean if a field has been set.
func (o *CloudInstanceTypeAllOf) HasCpus() bool {
	if o != nil && o.Cpus != nil {
		return true
	}

	return false
}

// SetCpus gets a reference to the given int64 and assigns it to the Cpus field.
func (o *CloudInstanceTypeAllOf) SetCpus(v int64) {
	o.Cpus = &v
}

// GetInstanceTypeId returns the InstanceTypeId field value if set, zero value otherwise.
func (o *CloudInstanceTypeAllOf) GetInstanceTypeId() string {
	if o == nil || o.InstanceTypeId == nil {
		var ret string
		return ret
	}
	return *o.InstanceTypeId
}

// GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInstanceTypeAllOf) GetInstanceTypeIdOk() (*string, bool) {
	if o == nil || o.InstanceTypeId == nil {
		return nil, false
	}
	return o.InstanceTypeId, true
}

// HasInstanceTypeId returns a boolean if a field has been set.
func (o *CloudInstanceTypeAllOf) HasInstanceTypeId() bool {
	if o != nil && o.InstanceTypeId != nil {
		return true
	}

	return false
}

// SetInstanceTypeId gets a reference to the given string and assigns it to the InstanceTypeId field.
func (o *CloudInstanceTypeAllOf) SetInstanceTypeId(v string) {
	o.InstanceTypeId = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *CloudInstanceTypeAllOf) GetMemory() int64 {
	if o == nil || o.Memory == nil {
		var ret int64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInstanceTypeAllOf) GetMemoryOk() (*int64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *CloudInstanceTypeAllOf) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int64 and assigns it to the Memory field.
func (o *CloudInstanceTypeAllOf) SetMemory(v int64) {
	o.Memory = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudInstanceTypeAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInstanceTypeAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudInstanceTypeAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudInstanceTypeAllOf) SetName(v string) {
	o.Name = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *CloudInstanceTypeAllOf) GetPlatform() string {
	if o == nil || o.Platform == nil {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInstanceTypeAllOf) GetPlatformOk() (*string, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *CloudInstanceTypeAllOf) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *CloudInstanceTypeAllOf) SetPlatform(v string) {
	o.Platform = &v
}

func (o CloudInstanceTypeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Architecture != nil {
		toSerialize["Architecture"] = o.Architecture
	}
	if o.CpuSpeed != nil {
		toSerialize["CpuSpeed"] = o.CpuSpeed
	}
	if o.Cpus != nil {
		toSerialize["Cpus"] = o.Cpus
	}
	if o.InstanceTypeId != nil {
		toSerialize["InstanceTypeId"] = o.InstanceTypeId
	}
	if o.Memory != nil {
		toSerialize["Memory"] = o.Memory
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Platform != nil {
		toSerialize["Platform"] = o.Platform
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudInstanceTypeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCloudInstanceTypeAllOf := _CloudInstanceTypeAllOf{}

	if err = json.Unmarshal(bytes, &varCloudInstanceTypeAllOf); err == nil {
		*o = CloudInstanceTypeAllOf(varCloudInstanceTypeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Architecture")
		delete(additionalProperties, "CpuSpeed")
		delete(additionalProperties, "Cpus")
		delete(additionalProperties, "InstanceTypeId")
		delete(additionalProperties, "Memory")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Platform")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudInstanceTypeAllOf struct {
	value *CloudInstanceTypeAllOf
	isSet bool
}

func (v NullableCloudInstanceTypeAllOf) Get() *CloudInstanceTypeAllOf {
	return v.value
}

func (v *NullableCloudInstanceTypeAllOf) Set(val *CloudInstanceTypeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudInstanceTypeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudInstanceTypeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudInstanceTypeAllOf(val *CloudInstanceTypeAllOf) *NullableCloudInstanceTypeAllOf {
	return &NullableCloudInstanceTypeAllOf{value: val, isSet: true}
}

func (v NullableCloudInstanceTypeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudInstanceTypeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
