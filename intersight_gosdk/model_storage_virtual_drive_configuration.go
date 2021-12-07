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
	"reflect"
	"strings"
)

// StorageVirtualDriveConfiguration Virtual Drive type models a single virtual drive that needs to be created through this policy.
type StorageVirtualDriveConfiguration struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This flag enables this virtual drive to be used as a boot drive.
	BootDrive *bool `json:"BootDrive,omitempty"`
	// This flag enables the virtual drive to use all the space available in the disk group. When this flag is enabled, the size property is ignored.
	ExpandToAvailable *bool `json:"ExpandToAvailable,omitempty"`
	// The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed.
	Name *string `json:"Name,omitempty"`
	// Virtual drive size in MebiBytes. Size is mandatory field except when the Expand to Available option is enabled.
	Size                 *int64                            `json:"Size,omitempty"`
	VirtualDrivePolicy   NullableStorageVirtualDrivePolicy `json:"VirtualDrivePolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageVirtualDriveConfiguration StorageVirtualDriveConfiguration

// NewStorageVirtualDriveConfiguration instantiates a new StorageVirtualDriveConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageVirtualDriveConfiguration(classId string, objectType string) *StorageVirtualDriveConfiguration {
	this := StorageVirtualDriveConfiguration{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageVirtualDriveConfigurationWithDefaults instantiates a new StorageVirtualDriveConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageVirtualDriveConfigurationWithDefaults() *StorageVirtualDriveConfiguration {
	this := StorageVirtualDriveConfiguration{}
	var classId string = "storage.VirtualDriveConfiguration"
	this.ClassId = classId
	var objectType string = "storage.VirtualDriveConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageVirtualDriveConfiguration) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfiguration) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageVirtualDriveConfiguration) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageVirtualDriveConfiguration) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfiguration) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageVirtualDriveConfiguration) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBootDrive returns the BootDrive field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfiguration) GetBootDrive() bool {
	if o == nil || o.BootDrive == nil {
		var ret bool
		return ret
	}
	return *o.BootDrive
}

// GetBootDriveOk returns a tuple with the BootDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfiguration) GetBootDriveOk() (*bool, bool) {
	if o == nil || o.BootDrive == nil {
		return nil, false
	}
	return o.BootDrive, true
}

// HasBootDrive returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfiguration) HasBootDrive() bool {
	if o != nil && o.BootDrive != nil {
		return true
	}

	return false
}

// SetBootDrive gets a reference to the given bool and assigns it to the BootDrive field.
func (o *StorageVirtualDriveConfiguration) SetBootDrive(v bool) {
	o.BootDrive = &v
}

// GetExpandToAvailable returns the ExpandToAvailable field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfiguration) GetExpandToAvailable() bool {
	if o == nil || o.ExpandToAvailable == nil {
		var ret bool
		return ret
	}
	return *o.ExpandToAvailable
}

// GetExpandToAvailableOk returns a tuple with the ExpandToAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfiguration) GetExpandToAvailableOk() (*bool, bool) {
	if o == nil || o.ExpandToAvailable == nil {
		return nil, false
	}
	return o.ExpandToAvailable, true
}

// HasExpandToAvailable returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfiguration) HasExpandToAvailable() bool {
	if o != nil && o.ExpandToAvailable != nil {
		return true
	}

	return false
}

// SetExpandToAvailable gets a reference to the given bool and assigns it to the ExpandToAvailable field.
func (o *StorageVirtualDriveConfiguration) SetExpandToAvailable(v bool) {
	o.ExpandToAvailable = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfiguration) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfiguration) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfiguration) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageVirtualDriveConfiguration) SetName(v string) {
	o.Name = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfiguration) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfiguration) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfiguration) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *StorageVirtualDriveConfiguration) SetSize(v int64) {
	o.Size = &v
}

// GetVirtualDrivePolicy returns the VirtualDrivePolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageVirtualDriveConfiguration) GetVirtualDrivePolicy() StorageVirtualDrivePolicy {
	if o == nil || o.VirtualDrivePolicy.Get() == nil {
		var ret StorageVirtualDrivePolicy
		return ret
	}
	return *o.VirtualDrivePolicy.Get()
}

// GetVirtualDrivePolicyOk returns a tuple with the VirtualDrivePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageVirtualDriveConfiguration) GetVirtualDrivePolicyOk() (*StorageVirtualDrivePolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualDrivePolicy.Get(), o.VirtualDrivePolicy.IsSet()
}

// HasVirtualDrivePolicy returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfiguration) HasVirtualDrivePolicy() bool {
	if o != nil && o.VirtualDrivePolicy.IsSet() {
		return true
	}

	return false
}

// SetVirtualDrivePolicy gets a reference to the given NullableStorageVirtualDrivePolicy and assigns it to the VirtualDrivePolicy field.
func (o *StorageVirtualDriveConfiguration) SetVirtualDrivePolicy(v StorageVirtualDrivePolicy) {
	o.VirtualDrivePolicy.Set(&v)
}

// SetVirtualDrivePolicyNil sets the value for VirtualDrivePolicy to be an explicit nil
func (o *StorageVirtualDriveConfiguration) SetVirtualDrivePolicyNil() {
	o.VirtualDrivePolicy.Set(nil)
}

// UnsetVirtualDrivePolicy ensures that no value is present for VirtualDrivePolicy, not even an explicit nil
func (o *StorageVirtualDriveConfiguration) UnsetVirtualDrivePolicy() {
	o.VirtualDrivePolicy.Unset()
}

func (o StorageVirtualDriveConfiguration) MarshalJSON() ([]byte, error) {
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
	if o.BootDrive != nil {
		toSerialize["BootDrive"] = o.BootDrive
	}
	if o.ExpandToAvailable != nil {
		toSerialize["ExpandToAvailable"] = o.ExpandToAvailable
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.VirtualDrivePolicy.IsSet() {
		toSerialize["VirtualDrivePolicy"] = o.VirtualDrivePolicy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageVirtualDriveConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	type StorageVirtualDriveConfigurationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// This flag enables this virtual drive to be used as a boot drive.
		BootDrive *bool `json:"BootDrive,omitempty"`
		// This flag enables the virtual drive to use all the space available in the disk group. When this flag is enabled, the size property is ignored.
		ExpandToAvailable *bool `json:"ExpandToAvailable,omitempty"`
		// The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed.
		Name *string `json:"Name,omitempty"`
		// Virtual drive size in MebiBytes. Size is mandatory field except when the Expand to Available option is enabled.
		Size               *int64                            `json:"Size,omitempty"`
		VirtualDrivePolicy NullableStorageVirtualDrivePolicy `json:"VirtualDrivePolicy,omitempty"`
	}

	varStorageVirtualDriveConfigurationWithoutEmbeddedStruct := StorageVirtualDriveConfigurationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageVirtualDriveConfigurationWithoutEmbeddedStruct)
	if err == nil {
		varStorageVirtualDriveConfiguration := _StorageVirtualDriveConfiguration{}
		varStorageVirtualDriveConfiguration.ClassId = varStorageVirtualDriveConfigurationWithoutEmbeddedStruct.ClassId
		varStorageVirtualDriveConfiguration.ObjectType = varStorageVirtualDriveConfigurationWithoutEmbeddedStruct.ObjectType
		varStorageVirtualDriveConfiguration.BootDrive = varStorageVirtualDriveConfigurationWithoutEmbeddedStruct.BootDrive
		varStorageVirtualDriveConfiguration.ExpandToAvailable = varStorageVirtualDriveConfigurationWithoutEmbeddedStruct.ExpandToAvailable
		varStorageVirtualDriveConfiguration.Name = varStorageVirtualDriveConfigurationWithoutEmbeddedStruct.Name
		varStorageVirtualDriveConfiguration.Size = varStorageVirtualDriveConfigurationWithoutEmbeddedStruct.Size
		varStorageVirtualDriveConfiguration.VirtualDrivePolicy = varStorageVirtualDriveConfigurationWithoutEmbeddedStruct.VirtualDrivePolicy
		*o = StorageVirtualDriveConfiguration(varStorageVirtualDriveConfiguration)
	} else {
		return err
	}

	varStorageVirtualDriveConfiguration := _StorageVirtualDriveConfiguration{}

	err = json.Unmarshal(bytes, &varStorageVirtualDriveConfiguration)
	if err == nil {
		o.MoBaseComplexType = varStorageVirtualDriveConfiguration.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BootDrive")
		delete(additionalProperties, "ExpandToAvailable")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "VirtualDrivePolicy")

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

type NullableStorageVirtualDriveConfiguration struct {
	value *StorageVirtualDriveConfiguration
	isSet bool
}

func (v NullableStorageVirtualDriveConfiguration) Get() *StorageVirtualDriveConfiguration {
	return v.value
}

func (v *NullableStorageVirtualDriveConfiguration) Set(val *StorageVirtualDriveConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageVirtualDriveConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageVirtualDriveConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageVirtualDriveConfiguration(val *StorageVirtualDriveConfiguration) *NullableStorageVirtualDriveConfiguration {
	return &NullableStorageVirtualDriveConfiguration{value: val, isSet: true}
}

func (v NullableStorageVirtualDriveConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageVirtualDriveConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
