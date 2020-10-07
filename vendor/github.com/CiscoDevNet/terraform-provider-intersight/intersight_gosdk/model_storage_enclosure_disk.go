/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-31T04:35:53Z.
 *
 * API version: 1.0.9-2110
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// StorageEnclosureDisk Physical Disk on the enclosure.
type StorageEnclosureDisk struct {
	EquipmentBase
	// The block size of the physical disk in bytes.
	BlockSize *string `json:"BlockSize,omitempty"`
	// This field represents the disk Id in the storage enclosure.
	DiskId *string `json:"DiskId,omitempty"`
	// This field identifies the current disk configuration applied in the physical disk.
	DiskState *string `json:"DiskState,omitempty"`
	// The current health state of the enclosure disk.
	Health *string `json:"Health,omitempty"`
	// The number of blocks present on the physical disk.
	NumBlocks *string `json:"NumBlocks,omitempty"`
	// This field identifies the Product ID for physicalDisk.
	Pid *string `json:"Pid,omitempty"`
	// This field identifies the SAS address assigned to the disk SAS port-1.
	SasAddress1 *string `json:"SasAddress1,omitempty"`
	// This field identifies the SAS address assigned to the disk SAS port-2.
	SasAddress2 *string `json:"SasAddress2,omitempty"`
	// The size of the physical disk in MB.
	Size                 *string                              `json:"Size,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	PhysicalDisk         *StoragePhysicalDiskRelationship     `json:"PhysicalDisk,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	StorageEnclosure     *StorageEnclosureRelationship        `json:"StorageEnclosure,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageEnclosureDisk StorageEnclosureDisk

// NewStorageEnclosureDisk instantiates a new StorageEnclosureDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageEnclosureDisk() *StorageEnclosureDisk {
	this := StorageEnclosureDisk{}
	return &this
}

// NewStorageEnclosureDiskWithDefaults instantiates a new StorageEnclosureDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageEnclosureDiskWithDefaults() *StorageEnclosureDisk {
	this := StorageEnclosureDisk{}
	return &this
}

// GetBlockSize returns the BlockSize field value if set, zero value otherwise.
func (o *StorageEnclosureDisk) GetBlockSize() string {
	if o == nil || o.BlockSize == nil {
		var ret string
		return ret
	}
	return *o.BlockSize
}

// GetBlockSizeOk returns a tuple with the BlockSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDisk) GetBlockSizeOk() (*string, bool) {
	if o == nil || o.BlockSize == nil {
		return nil, false
	}
	return o.BlockSize, true
}

// HasBlockSize returns a boolean if a field has been set.
func (o *StorageEnclosureDisk) HasBlockSize() bool {
	if o != nil && o.BlockSize != nil {
		return true
	}

	return false
}

// SetBlockSize gets a reference to the given string and assigns it to the BlockSize field.
func (o *StorageEnclosureDisk) SetBlockSize(v string) {
	o.BlockSize = &v
}

// GetDiskId returns the DiskId field value if set, zero value otherwise.
func (o *StorageEnclosureDisk) GetDiskId() string {
	if o == nil || o.DiskId == nil {
		var ret string
		return ret
	}
	return *o.DiskId
}

// GetDiskIdOk returns a tuple with the DiskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDisk) GetDiskIdOk() (*string, bool) {
	if o == nil || o.DiskId == nil {
		return nil, false
	}
	return o.DiskId, true
}

// HasDiskId returns a boolean if a field has been set.
func (o *StorageEnclosureDisk) HasDiskId() bool {
	if o != nil && o.DiskId != nil {
		return true
	}

	return false
}

// SetDiskId gets a reference to the given string and assigns it to the DiskId field.
func (o *StorageEnclosureDisk) SetDiskId(v string) {
	o.DiskId = &v
}

// GetDiskState returns the DiskState field value if set, zero value otherwise.
func (o *StorageEnclosureDisk) GetDiskState() string {
	if o == nil || o.DiskState == nil {
		var ret string
		return ret
	}
	return *o.DiskState
}

// GetDiskStateOk returns a tuple with the DiskState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDisk) GetDiskStateOk() (*string, bool) {
	if o == nil || o.DiskState == nil {
		return nil, false
	}
	return o.DiskState, true
}

// HasDiskState returns a boolean if a field has been set.
func (o *StorageEnclosureDisk) HasDiskState() bool {
	if o != nil && o.DiskState != nil {
		return true
	}

	return false
}

// SetDiskState gets a reference to the given string and assigns it to the DiskState field.
func (o *StorageEnclosureDisk) SetDiskState(v string) {
	o.DiskState = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *StorageEnclosureDisk) GetHealth() string {
	if o == nil || o.Health == nil {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDisk) GetHealthOk() (*string, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *StorageEnclosureDisk) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *StorageEnclosureDisk) SetHealth(v string) {
	o.Health = &v
}

// GetNumBlocks returns the NumBlocks field value if set, zero value otherwise.
func (o *StorageEnclosureDisk) GetNumBlocks() string {
	if o == nil || o.NumBlocks == nil {
		var ret string
		return ret
	}
	return *o.NumBlocks
}

// GetNumBlocksOk returns a tuple with the NumBlocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDisk) GetNumBlocksOk() (*string, bool) {
	if o == nil || o.NumBlocks == nil {
		return nil, false
	}
	return o.NumBlocks, true
}

// HasNumBlocks returns a boolean if a field has been set.
func (o *StorageEnclosureDisk) HasNumBlocks() bool {
	if o != nil && o.NumBlocks != nil {
		return true
	}

	return false
}

// SetNumBlocks gets a reference to the given string and assigns it to the NumBlocks field.
func (o *StorageEnclosureDisk) SetNumBlocks(v string) {
	o.NumBlocks = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *StorageEnclosureDisk) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDisk) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *StorageEnclosureDisk) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *StorageEnclosureDisk) SetPid(v string) {
	o.Pid = &v
}

// GetSasAddress1 returns the SasAddress1 field value if set, zero value otherwise.
func (o *StorageEnclosureDisk) GetSasAddress1() string {
	if o == nil || o.SasAddress1 == nil {
		var ret string
		return ret
	}
	return *o.SasAddress1
}

// GetSasAddress1Ok returns a tuple with the SasAddress1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDisk) GetSasAddress1Ok() (*string, bool) {
	if o == nil || o.SasAddress1 == nil {
		return nil, false
	}
	return o.SasAddress1, true
}

// HasSasAddress1 returns a boolean if a field has been set.
func (o *StorageEnclosureDisk) HasSasAddress1() bool {
	if o != nil && o.SasAddress1 != nil {
		return true
	}

	return false
}

// SetSasAddress1 gets a reference to the given string and assigns it to the SasAddress1 field.
func (o *StorageEnclosureDisk) SetSasAddress1(v string) {
	o.SasAddress1 = &v
}

// GetSasAddress2 returns the SasAddress2 field value if set, zero value otherwise.
func (o *StorageEnclosureDisk) GetSasAddress2() string {
	if o == nil || o.SasAddress2 == nil {
		var ret string
		return ret
	}
	return *o.SasAddress2
}

// GetSasAddress2Ok returns a tuple with the SasAddress2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDisk) GetSasAddress2Ok() (*string, bool) {
	if o == nil || o.SasAddress2 == nil {
		return nil, false
	}
	return o.SasAddress2, true
}

// HasSasAddress2 returns a boolean if a field has been set.
func (o *StorageEnclosureDisk) HasSasAddress2() bool {
	if o != nil && o.SasAddress2 != nil {
		return true
	}

	return false
}

// SetSasAddress2 gets a reference to the given string and assigns it to the SasAddress2 field.
func (o *StorageEnclosureDisk) SetSasAddress2(v string) {
	o.SasAddress2 = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageEnclosureDisk) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDisk) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageEnclosureDisk) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *StorageEnclosureDisk) SetSize(v string) {
	o.Size = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageEnclosureDisk) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDisk) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageEnclosureDisk) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageEnclosureDisk) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetPhysicalDisk returns the PhysicalDisk field value if set, zero value otherwise.
func (o *StorageEnclosureDisk) GetPhysicalDisk() StoragePhysicalDiskRelationship {
	if o == nil || o.PhysicalDisk == nil {
		var ret StoragePhysicalDiskRelationship
		return ret
	}
	return *o.PhysicalDisk
}

// GetPhysicalDiskOk returns a tuple with the PhysicalDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDisk) GetPhysicalDiskOk() (*StoragePhysicalDiskRelationship, bool) {
	if o == nil || o.PhysicalDisk == nil {
		return nil, false
	}
	return o.PhysicalDisk, true
}

// HasPhysicalDisk returns a boolean if a field has been set.
func (o *StorageEnclosureDisk) HasPhysicalDisk() bool {
	if o != nil && o.PhysicalDisk != nil {
		return true
	}

	return false
}

// SetPhysicalDisk gets a reference to the given StoragePhysicalDiskRelationship and assigns it to the PhysicalDisk field.
func (o *StorageEnclosureDisk) SetPhysicalDisk(v StoragePhysicalDiskRelationship) {
	o.PhysicalDisk = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageEnclosureDisk) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDisk) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageEnclosureDisk) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageEnclosureDisk) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageEnclosure returns the StorageEnclosure field value if set, zero value otherwise.
func (o *StorageEnclosureDisk) GetStorageEnclosure() StorageEnclosureRelationship {
	if o == nil || o.StorageEnclosure == nil {
		var ret StorageEnclosureRelationship
		return ret
	}
	return *o.StorageEnclosure
}

// GetStorageEnclosureOk returns a tuple with the StorageEnclosure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDisk) GetStorageEnclosureOk() (*StorageEnclosureRelationship, bool) {
	if o == nil || o.StorageEnclosure == nil {
		return nil, false
	}
	return o.StorageEnclosure, true
}

// HasStorageEnclosure returns a boolean if a field has been set.
func (o *StorageEnclosureDisk) HasStorageEnclosure() bool {
	if o != nil && o.StorageEnclosure != nil {
		return true
	}

	return false
}

// SetStorageEnclosure gets a reference to the given StorageEnclosureRelationship and assigns it to the StorageEnclosure field.
func (o *StorageEnclosureDisk) SetStorageEnclosure(v StorageEnclosureRelationship) {
	o.StorageEnclosure = &v
}

func (o StorageEnclosureDisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if o.BlockSize != nil {
		toSerialize["BlockSize"] = o.BlockSize
	}
	if o.DiskId != nil {
		toSerialize["DiskId"] = o.DiskId
	}
	if o.DiskState != nil {
		toSerialize["DiskState"] = o.DiskState
	}
	if o.Health != nil {
		toSerialize["Health"] = o.Health
	}
	if o.NumBlocks != nil {
		toSerialize["NumBlocks"] = o.NumBlocks
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	if o.SasAddress1 != nil {
		toSerialize["SasAddress1"] = o.SasAddress1
	}
	if o.SasAddress2 != nil {
		toSerialize["SasAddress2"] = o.SasAddress2
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.PhysicalDisk != nil {
		toSerialize["PhysicalDisk"] = o.PhysicalDisk
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageEnclosure != nil {
		toSerialize["StorageEnclosure"] = o.StorageEnclosure
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageEnclosureDisk) UnmarshalJSON(bytes []byte) (err error) {
	type StorageEnclosureDiskWithoutEmbeddedStruct struct {
		// The block size of the physical disk in bytes.
		BlockSize *string `json:"BlockSize,omitempty"`
		// This field represents the disk Id in the storage enclosure.
		DiskId *string `json:"DiskId,omitempty"`
		// This field identifies the current disk configuration applied in the physical disk.
		DiskState *string `json:"DiskState,omitempty"`
		// The current health state of the enclosure disk.
		Health *string `json:"Health,omitempty"`
		// The number of blocks present on the physical disk.
		NumBlocks *string `json:"NumBlocks,omitempty"`
		// This field identifies the Product ID for physicalDisk.
		Pid *string `json:"Pid,omitempty"`
		// This field identifies the SAS address assigned to the disk SAS port-1.
		SasAddress1 *string `json:"SasAddress1,omitempty"`
		// This field identifies the SAS address assigned to the disk SAS port-2.
		SasAddress2 *string `json:"SasAddress2,omitempty"`
		// The size of the physical disk in MB.
		Size                *string                              `json:"Size,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		PhysicalDisk        *StoragePhysicalDiskRelationship     `json:"PhysicalDisk,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		StorageEnclosure    *StorageEnclosureRelationship        `json:"StorageEnclosure,omitempty"`
	}

	varStorageEnclosureDiskWithoutEmbeddedStruct := StorageEnclosureDiskWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageEnclosureDiskWithoutEmbeddedStruct)
	if err == nil {
		varStorageEnclosureDisk := _StorageEnclosureDisk{}
		varStorageEnclosureDisk.BlockSize = varStorageEnclosureDiskWithoutEmbeddedStruct.BlockSize
		varStorageEnclosureDisk.DiskId = varStorageEnclosureDiskWithoutEmbeddedStruct.DiskId
		varStorageEnclosureDisk.DiskState = varStorageEnclosureDiskWithoutEmbeddedStruct.DiskState
		varStorageEnclosureDisk.Health = varStorageEnclosureDiskWithoutEmbeddedStruct.Health
		varStorageEnclosureDisk.NumBlocks = varStorageEnclosureDiskWithoutEmbeddedStruct.NumBlocks
		varStorageEnclosureDisk.Pid = varStorageEnclosureDiskWithoutEmbeddedStruct.Pid
		varStorageEnclosureDisk.SasAddress1 = varStorageEnclosureDiskWithoutEmbeddedStruct.SasAddress1
		varStorageEnclosureDisk.SasAddress2 = varStorageEnclosureDiskWithoutEmbeddedStruct.SasAddress2
		varStorageEnclosureDisk.Size = varStorageEnclosureDiskWithoutEmbeddedStruct.Size
		varStorageEnclosureDisk.InventoryDeviceInfo = varStorageEnclosureDiskWithoutEmbeddedStruct.InventoryDeviceInfo
		varStorageEnclosureDisk.PhysicalDisk = varStorageEnclosureDiskWithoutEmbeddedStruct.PhysicalDisk
		varStorageEnclosureDisk.RegisteredDevice = varStorageEnclosureDiskWithoutEmbeddedStruct.RegisteredDevice
		varStorageEnclosureDisk.StorageEnclosure = varStorageEnclosureDiskWithoutEmbeddedStruct.StorageEnclosure
		*o = StorageEnclosureDisk(varStorageEnclosureDisk)
	} else {
		return err
	}

	varStorageEnclosureDisk := _StorageEnclosureDisk{}

	err = json.Unmarshal(bytes, &varStorageEnclosureDisk)
	if err == nil {
		o.EquipmentBase = varStorageEnclosureDisk.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "BlockSize")
		delete(additionalProperties, "DiskId")
		delete(additionalProperties, "DiskState")
		delete(additionalProperties, "Health")
		delete(additionalProperties, "NumBlocks")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "SasAddress1")
		delete(additionalProperties, "SasAddress2")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "PhysicalDisk")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageEnclosure")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

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

type NullableStorageEnclosureDisk struct {
	value *StorageEnclosureDisk
	isSet bool
}

func (v NullableStorageEnclosureDisk) Get() *StorageEnclosureDisk {
	return v.value
}

func (v *NullableStorageEnclosureDisk) Set(val *StorageEnclosureDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageEnclosureDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageEnclosureDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageEnclosureDisk(val *StorageEnclosureDisk) *NullableStorageEnclosureDisk {
	return &NullableStorageEnclosureDisk{value: val, isSet: true}
}

func (v NullableStorageEnclosureDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageEnclosureDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}