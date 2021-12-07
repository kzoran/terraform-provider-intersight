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

// StorageFlexFlashVirtualDriveAllOf Definition of the list of properties defined in 'storage.FlexFlashVirtualDrive', excluding properties defined in parent classes.
type StorageFlexFlashVirtualDriveAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The drive scope of the flex flash virtual drive.
	DriveScope *string `json:"DriveScope,omitempty"`
	// Status of virtual drive on the flex controller.
	DriveStatus *string `json:"DriveStatus,omitempty"`
	// The partition Id of the flex flash virtual Drive.
	PartitionId *string `json:"PartitionId,omitempty"`
	// The resident image on the flex flash virtual Drive.
	ResidentImage *string `json:"ResidentImage,omitempty"`
	// Size of virtual drive on the flex controller.
	Size *string `json:"Size,omitempty"`
	// Virtual drive on the flex flash controller.
	VirtualDrive               *string                                 `json:"VirtualDrive,omitempty"`
	InventoryDeviceInfo        *InventoryDeviceInfoRelationship        `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice           *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	StorageFlexFlashController *StorageFlexFlashControllerRelationship `json:"StorageFlexFlashController,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _StorageFlexFlashVirtualDriveAllOf StorageFlexFlashVirtualDriveAllOf

// NewStorageFlexFlashVirtualDriveAllOf instantiates a new StorageFlexFlashVirtualDriveAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageFlexFlashVirtualDriveAllOf(classId string, objectType string) *StorageFlexFlashVirtualDriveAllOf {
	this := StorageFlexFlashVirtualDriveAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageFlexFlashVirtualDriveAllOfWithDefaults instantiates a new StorageFlexFlashVirtualDriveAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageFlexFlashVirtualDriveAllOfWithDefaults() *StorageFlexFlashVirtualDriveAllOf {
	this := StorageFlexFlashVirtualDriveAllOf{}
	var classId string = "storage.FlexFlashVirtualDrive"
	this.ClassId = classId
	var objectType string = "storage.FlexFlashVirtualDrive"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageFlexFlashVirtualDriveAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDriveAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageFlexFlashVirtualDriveAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageFlexFlashVirtualDriveAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDriveAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageFlexFlashVirtualDriveAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDriveScope returns the DriveScope field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDriveAllOf) GetDriveScope() string {
	if o == nil || o.DriveScope == nil {
		var ret string
		return ret
	}
	return *o.DriveScope
}

// GetDriveScopeOk returns a tuple with the DriveScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDriveAllOf) GetDriveScopeOk() (*string, bool) {
	if o == nil || o.DriveScope == nil {
		return nil, false
	}
	return o.DriveScope, true
}

// HasDriveScope returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDriveAllOf) HasDriveScope() bool {
	if o != nil && o.DriveScope != nil {
		return true
	}

	return false
}

// SetDriveScope gets a reference to the given string and assigns it to the DriveScope field.
func (o *StorageFlexFlashVirtualDriveAllOf) SetDriveScope(v string) {
	o.DriveScope = &v
}

// GetDriveStatus returns the DriveStatus field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDriveAllOf) GetDriveStatus() string {
	if o == nil || o.DriveStatus == nil {
		var ret string
		return ret
	}
	return *o.DriveStatus
}

// GetDriveStatusOk returns a tuple with the DriveStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDriveAllOf) GetDriveStatusOk() (*string, bool) {
	if o == nil || o.DriveStatus == nil {
		return nil, false
	}
	return o.DriveStatus, true
}

// HasDriveStatus returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDriveAllOf) HasDriveStatus() bool {
	if o != nil && o.DriveStatus != nil {
		return true
	}

	return false
}

// SetDriveStatus gets a reference to the given string and assigns it to the DriveStatus field.
func (o *StorageFlexFlashVirtualDriveAllOf) SetDriveStatus(v string) {
	o.DriveStatus = &v
}

// GetPartitionId returns the PartitionId field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDriveAllOf) GetPartitionId() string {
	if o == nil || o.PartitionId == nil {
		var ret string
		return ret
	}
	return *o.PartitionId
}

// GetPartitionIdOk returns a tuple with the PartitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDriveAllOf) GetPartitionIdOk() (*string, bool) {
	if o == nil || o.PartitionId == nil {
		return nil, false
	}
	return o.PartitionId, true
}

// HasPartitionId returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDriveAllOf) HasPartitionId() bool {
	if o != nil && o.PartitionId != nil {
		return true
	}

	return false
}

// SetPartitionId gets a reference to the given string and assigns it to the PartitionId field.
func (o *StorageFlexFlashVirtualDriveAllOf) SetPartitionId(v string) {
	o.PartitionId = &v
}

// GetResidentImage returns the ResidentImage field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDriveAllOf) GetResidentImage() string {
	if o == nil || o.ResidentImage == nil {
		var ret string
		return ret
	}
	return *o.ResidentImage
}

// GetResidentImageOk returns a tuple with the ResidentImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDriveAllOf) GetResidentImageOk() (*string, bool) {
	if o == nil || o.ResidentImage == nil {
		return nil, false
	}
	return o.ResidentImage, true
}

// HasResidentImage returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDriveAllOf) HasResidentImage() bool {
	if o != nil && o.ResidentImage != nil {
		return true
	}

	return false
}

// SetResidentImage gets a reference to the given string and assigns it to the ResidentImage field.
func (o *StorageFlexFlashVirtualDriveAllOf) SetResidentImage(v string) {
	o.ResidentImage = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDriveAllOf) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDriveAllOf) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDriveAllOf) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *StorageFlexFlashVirtualDriveAllOf) SetSize(v string) {
	o.Size = &v
}

// GetVirtualDrive returns the VirtualDrive field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDriveAllOf) GetVirtualDrive() string {
	if o == nil || o.VirtualDrive == nil {
		var ret string
		return ret
	}
	return *o.VirtualDrive
}

// GetVirtualDriveOk returns a tuple with the VirtualDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDriveAllOf) GetVirtualDriveOk() (*string, bool) {
	if o == nil || o.VirtualDrive == nil {
		return nil, false
	}
	return o.VirtualDrive, true
}

// HasVirtualDrive returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDriveAllOf) HasVirtualDrive() bool {
	if o != nil && o.VirtualDrive != nil {
		return true
	}

	return false
}

// SetVirtualDrive gets a reference to the given string and assigns it to the VirtualDrive field.
func (o *StorageFlexFlashVirtualDriveAllOf) SetVirtualDrive(v string) {
	o.VirtualDrive = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDriveAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDriveAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDriveAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageFlexFlashVirtualDriveAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDriveAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDriveAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDriveAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageFlexFlashVirtualDriveAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageFlexFlashController returns the StorageFlexFlashController field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDriveAllOf) GetStorageFlexFlashController() StorageFlexFlashControllerRelationship {
	if o == nil || o.StorageFlexFlashController == nil {
		var ret StorageFlexFlashControllerRelationship
		return ret
	}
	return *o.StorageFlexFlashController
}

// GetStorageFlexFlashControllerOk returns a tuple with the StorageFlexFlashController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDriveAllOf) GetStorageFlexFlashControllerOk() (*StorageFlexFlashControllerRelationship, bool) {
	if o == nil || o.StorageFlexFlashController == nil {
		return nil, false
	}
	return o.StorageFlexFlashController, true
}

// HasStorageFlexFlashController returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDriveAllOf) HasStorageFlexFlashController() bool {
	if o != nil && o.StorageFlexFlashController != nil {
		return true
	}

	return false
}

// SetStorageFlexFlashController gets a reference to the given StorageFlexFlashControllerRelationship and assigns it to the StorageFlexFlashController field.
func (o *StorageFlexFlashVirtualDriveAllOf) SetStorageFlexFlashController(v StorageFlexFlashControllerRelationship) {
	o.StorageFlexFlashController = &v
}

func (o StorageFlexFlashVirtualDriveAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DriveScope != nil {
		toSerialize["DriveScope"] = o.DriveScope
	}
	if o.DriveStatus != nil {
		toSerialize["DriveStatus"] = o.DriveStatus
	}
	if o.PartitionId != nil {
		toSerialize["PartitionId"] = o.PartitionId
	}
	if o.ResidentImage != nil {
		toSerialize["ResidentImage"] = o.ResidentImage
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.VirtualDrive != nil {
		toSerialize["VirtualDrive"] = o.VirtualDrive
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageFlexFlashController != nil {
		toSerialize["StorageFlexFlashController"] = o.StorageFlexFlashController
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageFlexFlashVirtualDriveAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageFlexFlashVirtualDriveAllOf := _StorageFlexFlashVirtualDriveAllOf{}

	if err = json.Unmarshal(bytes, &varStorageFlexFlashVirtualDriveAllOf); err == nil {
		*o = StorageFlexFlashVirtualDriveAllOf(varStorageFlexFlashVirtualDriveAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DriveScope")
		delete(additionalProperties, "DriveStatus")
		delete(additionalProperties, "PartitionId")
		delete(additionalProperties, "ResidentImage")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "VirtualDrive")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageFlexFlashController")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageFlexFlashVirtualDriveAllOf struct {
	value *StorageFlexFlashVirtualDriveAllOf
	isSet bool
}

func (v NullableStorageFlexFlashVirtualDriveAllOf) Get() *StorageFlexFlashVirtualDriveAllOf {
	return v.value
}

func (v *NullableStorageFlexFlashVirtualDriveAllOf) Set(val *StorageFlexFlashVirtualDriveAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageFlexFlashVirtualDriveAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageFlexFlashVirtualDriveAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageFlexFlashVirtualDriveAllOf(val *StorageFlexFlashVirtualDriveAllOf) *NullableStorageFlexFlashVirtualDriveAllOf {
	return &NullableStorageFlexFlashVirtualDriveAllOf{value: val, isSet: true}
}

func (v NullableStorageFlexFlashVirtualDriveAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageFlexFlashVirtualDriveAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
