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
	"reflect"
	"strings"
)

// HyperflexHxapVirtualDisk The Virtual disk created on HyperFlex Application Platform compute cluster.
type HyperflexHxapVirtualDisk struct {
	VirtualizationBaseVirtualDisk
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Access mode of the virtual disk. * `ReadWriteOnce` - Read write permisisons to a Virtual disk by a single virtual machine. * `ReadWriteMany` - Read write permisisons to a Virtual disk by multiple virtual machines. * `ReadOnlyMany` - Read only permisisons to a Virtual disk by multiple virtual machines.
	AccessMode *string `json:"AccessMode,omitempty"`
	// File mode of the disk  example - Filesystem, Block. * `Block` - It is a Block virtual disk. * `Filesystem` - It is a File system virtual disk.
	Mode *string `json:"Mode,omitempty"`
	// Source file path associated with virtual machine disk.
	SourceFilePath *string `json:"SourceFilePath,omitempty"`
	// Source disk name from where the clone is done.
	SourceVirtualDisk *string                     `json:"SourceVirtualDisk,omitempty"`
	Status            NullableHyperflexDiskStatus `json:"Status,omitempty"`
	// UUID of the virtual disk.
	Uuid                 *string                                  `json:"Uuid,omitempty"`
	Cluster              *HyperflexHxapClusterRelationship        `json:"Cluster,omitempty"`
	VirtualMachine       *HyperflexHxapVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxapVirtualDisk HyperflexHxapVirtualDisk

// NewHyperflexHxapVirtualDisk instantiates a new HyperflexHxapVirtualDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxapVirtualDisk(classId string, objectType string) *HyperflexHxapVirtualDisk {
	this := HyperflexHxapVirtualDisk{}
	this.ClassId = classId
	this.ObjectType = objectType
	var accessMode string = "ReadWriteOnce"
	this.AccessMode = &accessMode
	var mode string = "Block"
	this.Mode = &mode
	return &this
}

// NewHyperflexHxapVirtualDiskWithDefaults instantiates a new HyperflexHxapVirtualDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxapVirtualDiskWithDefaults() *HyperflexHxapVirtualDisk {
	this := HyperflexHxapVirtualDisk{}
	var classId string = "hyperflex.HxapVirtualDisk"
	this.ClassId = classId
	var objectType string = "hyperflex.HxapVirtualDisk"
	this.ObjectType = objectType
	var accessMode string = "ReadWriteOnce"
	this.AccessMode = &accessMode
	var mode string = "Block"
	this.Mode = &mode
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxapVirtualDisk) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualDisk) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxapVirtualDisk) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxapVirtualDisk) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualDisk) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxapVirtualDisk) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccessMode returns the AccessMode field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualDisk) GetAccessMode() string {
	if o == nil || o.AccessMode == nil {
		var ret string
		return ret
	}
	return *o.AccessMode
}

// GetAccessModeOk returns a tuple with the AccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualDisk) GetAccessModeOk() (*string, bool) {
	if o == nil || o.AccessMode == nil {
		return nil, false
	}
	return o.AccessMode, true
}

// HasAccessMode returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualDisk) HasAccessMode() bool {
	if o != nil && o.AccessMode != nil {
		return true
	}

	return false
}

// SetAccessMode gets a reference to the given string and assigns it to the AccessMode field.
func (o *HyperflexHxapVirtualDisk) SetAccessMode(v string) {
	o.AccessMode = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualDisk) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualDisk) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualDisk) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *HyperflexHxapVirtualDisk) SetMode(v string) {
	o.Mode = &v
}

// GetSourceFilePath returns the SourceFilePath field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualDisk) GetSourceFilePath() string {
	if o == nil || o.SourceFilePath == nil {
		var ret string
		return ret
	}
	return *o.SourceFilePath
}

// GetSourceFilePathOk returns a tuple with the SourceFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualDisk) GetSourceFilePathOk() (*string, bool) {
	if o == nil || o.SourceFilePath == nil {
		return nil, false
	}
	return o.SourceFilePath, true
}

// HasSourceFilePath returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualDisk) HasSourceFilePath() bool {
	if o != nil && o.SourceFilePath != nil {
		return true
	}

	return false
}

// SetSourceFilePath gets a reference to the given string and assigns it to the SourceFilePath field.
func (o *HyperflexHxapVirtualDisk) SetSourceFilePath(v string) {
	o.SourceFilePath = &v
}

// GetSourceVirtualDisk returns the SourceVirtualDisk field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualDisk) GetSourceVirtualDisk() string {
	if o == nil || o.SourceVirtualDisk == nil {
		var ret string
		return ret
	}
	return *o.SourceVirtualDisk
}

// GetSourceVirtualDiskOk returns a tuple with the SourceVirtualDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualDisk) GetSourceVirtualDiskOk() (*string, bool) {
	if o == nil || o.SourceVirtualDisk == nil {
		return nil, false
	}
	return o.SourceVirtualDisk, true
}

// HasSourceVirtualDisk returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualDisk) HasSourceVirtualDisk() bool {
	if o != nil && o.SourceVirtualDisk != nil {
		return true
	}

	return false
}

// SetSourceVirtualDisk gets a reference to the given string and assigns it to the SourceVirtualDisk field.
func (o *HyperflexHxapVirtualDisk) SetSourceVirtualDisk(v string) {
	o.SourceVirtualDisk = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHxapVirtualDisk) GetStatus() HyperflexDiskStatus {
	if o == nil || o.Status.Get() == nil {
		var ret HyperflexDiskStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHxapVirtualDisk) GetStatusOk() (*HyperflexDiskStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualDisk) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableHyperflexDiskStatus and assigns it to the Status field.
func (o *HyperflexHxapVirtualDisk) SetStatus(v HyperflexDiskStatus) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *HyperflexHxapVirtualDisk) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *HyperflexHxapVirtualDisk) UnsetStatus() {
	o.Status.Unset()
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualDisk) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualDisk) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualDisk) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HyperflexHxapVirtualDisk) SetUuid(v string) {
	o.Uuid = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualDisk) GetCluster() HyperflexHxapClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexHxapClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualDisk) GetClusterOk() (*HyperflexHxapClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualDisk) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexHxapClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexHxapVirtualDisk) SetCluster(v HyperflexHxapClusterRelationship) {
	o.Cluster = &v
}

// GetVirtualMachine returns the VirtualMachine field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualDisk) GetVirtualMachine() HyperflexHxapVirtualMachineRelationship {
	if o == nil || o.VirtualMachine == nil {
		var ret HyperflexHxapVirtualMachineRelationship
		return ret
	}
	return *o.VirtualMachine
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualDisk) GetVirtualMachineOk() (*HyperflexHxapVirtualMachineRelationship, bool) {
	if o == nil || o.VirtualMachine == nil {
		return nil, false
	}
	return o.VirtualMachine, true
}

// HasVirtualMachine returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualDisk) HasVirtualMachine() bool {
	if o != nil && o.VirtualMachine != nil {
		return true
	}

	return false
}

// SetVirtualMachine gets a reference to the given HyperflexHxapVirtualMachineRelationship and assigns it to the VirtualMachine field.
func (o *HyperflexHxapVirtualDisk) SetVirtualMachine(v HyperflexHxapVirtualMachineRelationship) {
	o.VirtualMachine = &v
}

func (o HyperflexHxapVirtualDisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseVirtualDisk, errVirtualizationBaseVirtualDisk := json.Marshal(o.VirtualizationBaseVirtualDisk)
	if errVirtualizationBaseVirtualDisk != nil {
		return []byte{}, errVirtualizationBaseVirtualDisk
	}
	errVirtualizationBaseVirtualDisk = json.Unmarshal([]byte(serializedVirtualizationBaseVirtualDisk), &toSerialize)
	if errVirtualizationBaseVirtualDisk != nil {
		return []byte{}, errVirtualizationBaseVirtualDisk
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AccessMode != nil {
		toSerialize["AccessMode"] = o.AccessMode
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	if o.SourceFilePath != nil {
		toSerialize["SourceFilePath"] = o.SourceFilePath
	}
	if o.SourceVirtualDisk != nil {
		toSerialize["SourceVirtualDisk"] = o.SourceVirtualDisk
	}
	if o.Status.IsSet() {
		toSerialize["Status"] = o.Status.Get()
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.VirtualMachine != nil {
		toSerialize["VirtualMachine"] = o.VirtualMachine
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxapVirtualDisk) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexHxapVirtualDiskWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Access mode of the virtual disk. * `ReadWriteOnce` - Read write permisisons to a Virtual disk by a single virtual machine. * `ReadWriteMany` - Read write permisisons to a Virtual disk by multiple virtual machines. * `ReadOnlyMany` - Read only permisisons to a Virtual disk by multiple virtual machines.
		AccessMode *string `json:"AccessMode,omitempty"`
		// File mode of the disk  example - Filesystem, Block. * `Block` - It is a Block virtual disk. * `Filesystem` - It is a File system virtual disk.
		Mode *string `json:"Mode,omitempty"`
		// Source file path associated with virtual machine disk.
		SourceFilePath *string `json:"SourceFilePath,omitempty"`
		// Source disk name from where the clone is done.
		SourceVirtualDisk *string                     `json:"SourceVirtualDisk,omitempty"`
		Status            NullableHyperflexDiskStatus `json:"Status,omitempty"`
		// UUID of the virtual disk.
		Uuid           *string                                  `json:"Uuid,omitempty"`
		Cluster        *HyperflexHxapClusterRelationship        `json:"Cluster,omitempty"`
		VirtualMachine *HyperflexHxapVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	}

	varHyperflexHxapVirtualDiskWithoutEmbeddedStruct := HyperflexHxapVirtualDiskWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexHxapVirtualDiskWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexHxapVirtualDisk := _HyperflexHxapVirtualDisk{}
		varHyperflexHxapVirtualDisk.ClassId = varHyperflexHxapVirtualDiskWithoutEmbeddedStruct.ClassId
		varHyperflexHxapVirtualDisk.ObjectType = varHyperflexHxapVirtualDiskWithoutEmbeddedStruct.ObjectType
		varHyperflexHxapVirtualDisk.AccessMode = varHyperflexHxapVirtualDiskWithoutEmbeddedStruct.AccessMode
		varHyperflexHxapVirtualDisk.Mode = varHyperflexHxapVirtualDiskWithoutEmbeddedStruct.Mode
		varHyperflexHxapVirtualDisk.SourceFilePath = varHyperflexHxapVirtualDiskWithoutEmbeddedStruct.SourceFilePath
		varHyperflexHxapVirtualDisk.SourceVirtualDisk = varHyperflexHxapVirtualDiskWithoutEmbeddedStruct.SourceVirtualDisk
		varHyperflexHxapVirtualDisk.Status = varHyperflexHxapVirtualDiskWithoutEmbeddedStruct.Status
		varHyperflexHxapVirtualDisk.Uuid = varHyperflexHxapVirtualDiskWithoutEmbeddedStruct.Uuid
		varHyperflexHxapVirtualDisk.Cluster = varHyperflexHxapVirtualDiskWithoutEmbeddedStruct.Cluster
		varHyperflexHxapVirtualDisk.VirtualMachine = varHyperflexHxapVirtualDiskWithoutEmbeddedStruct.VirtualMachine
		*o = HyperflexHxapVirtualDisk(varHyperflexHxapVirtualDisk)
	} else {
		return err
	}

	varHyperflexHxapVirtualDisk := _HyperflexHxapVirtualDisk{}

	err = json.Unmarshal(bytes, &varHyperflexHxapVirtualDisk)
	if err == nil {
		o.VirtualizationBaseVirtualDisk = varHyperflexHxapVirtualDisk.VirtualizationBaseVirtualDisk
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccessMode")
		delete(additionalProperties, "Mode")
		delete(additionalProperties, "SourceFilePath")
		delete(additionalProperties, "SourceVirtualDisk")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "VirtualMachine")

		// remove fields from embedded structs
		reflectVirtualizationBaseVirtualDisk := reflect.ValueOf(o.VirtualizationBaseVirtualDisk)
		for i := 0; i < reflectVirtualizationBaseVirtualDisk.Type().NumField(); i++ {
			t := reflectVirtualizationBaseVirtualDisk.Type().Field(i)

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

type NullableHyperflexHxapVirtualDisk struct {
	value *HyperflexHxapVirtualDisk
	isSet bool
}

func (v NullableHyperflexHxapVirtualDisk) Get() *HyperflexHxapVirtualDisk {
	return v.value
}

func (v *NullableHyperflexHxapVirtualDisk) Set(val *HyperflexHxapVirtualDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxapVirtualDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxapVirtualDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxapVirtualDisk(val *HyperflexHxapVirtualDisk) *NullableHyperflexHxapVirtualDisk {
	return &NullableHyperflexHxapVirtualDisk{value: val, isSet: true}
}

func (v NullableHyperflexHxapVirtualDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxapVirtualDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
