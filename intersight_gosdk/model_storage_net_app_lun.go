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
	"reflect"
	"strings"
)

// StorageNetAppLun NetApp LUN (logical unit number) is an identifier for a device called a logical unit addressed by a SAN protocol.
type StorageNetAppLun struct {
	StorageBaseVolume
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The operating system (OS) type for this LUN. * `Linux` - Family of open source Unix-like operating systems based on the Linux kernel. * `AIX` - Advanced Interactive Executive (AIX). * `HP-UX` - HP-UX is implementation of the Unix operating system, based on Unix System V. * `Hyper-V` - Windows Server 2008 or Windows Server 2012 Hyper-V. * `OpenVMS` - OpenVMS is multi-user, multiprocessing virtual memory-based operating system. * `Solaris` - Solaris is a Unix operating system. * `NetWare` - NetWare is a computer network operating system. * `VMware` - An enterprise-class, type-1 hypervisor developed by VMware for deploying and serving virtual computers. * `Windows` - Single-partition Windows disk using the Master Boot Record (MBR) partitioning style. * `Xen` - Xen is a type-1 hypervisor, providing services that allow multiple computer operating systems to execute on the same computer hardware concurrently.
	OsType *string `json:"OsType,omitempty"`
	// Serial number for the provisioned LUN.
	Serial *string `json:"Serial,omitempty"`
	// The administrative state of a LUN. * `offline` - The LUN is administratively offline, or a more detailed offline reason is not available. * `online` - The LUN is online.
	State *string `json:"State,omitempty"`
	// UUID of the LUN.
	Uuid  *string                           `json:"Uuid,omitempty"`
	Array *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	// An array of relationships to storageNetAppInitiatorGroup resources.
	Host                 []StorageNetAppInitiatorGroupRelationship `json:"Host,omitempty"`
	StorageContainer     *StorageNetAppVolumeRelationship          `json:"StorageContainer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppLun StorageNetAppLun

// NewStorageNetAppLun instantiates a new StorageNetAppLun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppLun(classId string, objectType string) *StorageNetAppLun {
	this := StorageNetAppLun{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppLunWithDefaults instantiates a new StorageNetAppLun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppLunWithDefaults() *StorageNetAppLun {
	this := StorageNetAppLun{}
	var classId string = "storage.NetAppLun"
	this.ClassId = classId
	var objectType string = "storage.NetAppLun"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppLun) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppLun) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppLun) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppLun) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppLun) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppLun) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOsType returns the OsType field value if set, zero value otherwise.
func (o *StorageNetAppLun) GetOsType() string {
	if o == nil || o.OsType == nil {
		var ret string
		return ret
	}
	return *o.OsType
}

// GetOsTypeOk returns a tuple with the OsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLun) GetOsTypeOk() (*string, bool) {
	if o == nil || o.OsType == nil {
		return nil, false
	}
	return o.OsType, true
}

// HasOsType returns a boolean if a field has been set.
func (o *StorageNetAppLun) HasOsType() bool {
	if o != nil && o.OsType != nil {
		return true
	}

	return false
}

// SetOsType gets a reference to the given string and assigns it to the OsType field.
func (o *StorageNetAppLun) SetOsType(v string) {
	o.OsType = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *StorageNetAppLun) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLun) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *StorageNetAppLun) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *StorageNetAppLun) SetSerial(v string) {
	o.Serial = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageNetAppLun) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLun) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageNetAppLun) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageNetAppLun) SetState(v string) {
	o.State = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppLun) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLun) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppLun) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppLun) SetUuid(v string) {
	o.Uuid = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppLun) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLun) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppLun) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppLun) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppLun) GetHost() []StorageNetAppInitiatorGroupRelationship {
	if o == nil {
		var ret []StorageNetAppInitiatorGroupRelationship
		return ret
	}
	return o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppLun) GetHostOk() (*[]StorageNetAppInitiatorGroupRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return &o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *StorageNetAppLun) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given []StorageNetAppInitiatorGroupRelationship and assigns it to the Host field.
func (o *StorageNetAppLun) SetHost(v []StorageNetAppInitiatorGroupRelationship) {
	o.Host = v
}

// GetStorageContainer returns the StorageContainer field value if set, zero value otherwise.
func (o *StorageNetAppLun) GetStorageContainer() StorageNetAppVolumeRelationship {
	if o == nil || o.StorageContainer == nil {
		var ret StorageNetAppVolumeRelationship
		return ret
	}
	return *o.StorageContainer
}

// GetStorageContainerOk returns a tuple with the StorageContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLun) GetStorageContainerOk() (*StorageNetAppVolumeRelationship, bool) {
	if o == nil || o.StorageContainer == nil {
		return nil, false
	}
	return o.StorageContainer, true
}

// HasStorageContainer returns a boolean if a field has been set.
func (o *StorageNetAppLun) HasStorageContainer() bool {
	if o != nil && o.StorageContainer != nil {
		return true
	}

	return false
}

// SetStorageContainer gets a reference to the given StorageNetAppVolumeRelationship and assigns it to the StorageContainer field.
func (o *StorageNetAppLun) SetStorageContainer(v StorageNetAppVolumeRelationship) {
	o.StorageContainer = &v
}

func (o StorageNetAppLun) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseVolume, errStorageBaseVolume := json.Marshal(o.StorageBaseVolume)
	if errStorageBaseVolume != nil {
		return []byte{}, errStorageBaseVolume
	}
	errStorageBaseVolume = json.Unmarshal([]byte(serializedStorageBaseVolume), &toSerialize)
	if errStorageBaseVolume != nil {
		return []byte{}, errStorageBaseVolume
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.OsType != nil {
		toSerialize["OsType"] = o.OsType
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
	}
	if o.StorageContainer != nil {
		toSerialize["StorageContainer"] = o.StorageContainer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppLun) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppLunWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The operating system (OS) type for this LUN. * `Linux` - Family of open source Unix-like operating systems based on the Linux kernel. * `AIX` - Advanced Interactive Executive (AIX). * `HP-UX` - HP-UX is implementation of the Unix operating system, based on Unix System V. * `Hyper-V` - Windows Server 2008 or Windows Server 2012 Hyper-V. * `OpenVMS` - OpenVMS is multi-user, multiprocessing virtual memory-based operating system. * `Solaris` - Solaris is a Unix operating system. * `NetWare` - NetWare is a computer network operating system. * `VMware` - An enterprise-class, type-1 hypervisor developed by VMware for deploying and serving virtual computers. * `Windows` - Single-partition Windows disk using the Master Boot Record (MBR) partitioning style. * `Xen` - Xen is a type-1 hypervisor, providing services that allow multiple computer operating systems to execute on the same computer hardware concurrently.
		OsType *string `json:"OsType,omitempty"`
		// Serial number for the provisioned LUN.
		Serial *string `json:"Serial,omitempty"`
		// The administrative state of a LUN. * `offline` - The LUN is administratively offline, or a more detailed offline reason is not available. * `online` - The LUN is online.
		State *string `json:"State,omitempty"`
		// UUID of the LUN.
		Uuid  *string                           `json:"Uuid,omitempty"`
		Array *StorageNetAppClusterRelationship `json:"Array,omitempty"`
		// An array of relationships to storageNetAppInitiatorGroup resources.
		Host             []StorageNetAppInitiatorGroupRelationship `json:"Host,omitempty"`
		StorageContainer *StorageNetAppVolumeRelationship          `json:"StorageContainer,omitempty"`
	}

	varStorageNetAppLunWithoutEmbeddedStruct := StorageNetAppLunWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppLunWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppLun := _StorageNetAppLun{}
		varStorageNetAppLun.ClassId = varStorageNetAppLunWithoutEmbeddedStruct.ClassId
		varStorageNetAppLun.ObjectType = varStorageNetAppLunWithoutEmbeddedStruct.ObjectType
		varStorageNetAppLun.OsType = varStorageNetAppLunWithoutEmbeddedStruct.OsType
		varStorageNetAppLun.Serial = varStorageNetAppLunWithoutEmbeddedStruct.Serial
		varStorageNetAppLun.State = varStorageNetAppLunWithoutEmbeddedStruct.State
		varStorageNetAppLun.Uuid = varStorageNetAppLunWithoutEmbeddedStruct.Uuid
		varStorageNetAppLun.Array = varStorageNetAppLunWithoutEmbeddedStruct.Array
		varStorageNetAppLun.Host = varStorageNetAppLunWithoutEmbeddedStruct.Host
		varStorageNetAppLun.StorageContainer = varStorageNetAppLunWithoutEmbeddedStruct.StorageContainer
		*o = StorageNetAppLun(varStorageNetAppLun)
	} else {
		return err
	}

	varStorageNetAppLun := _StorageNetAppLun{}

	err = json.Unmarshal(bytes, &varStorageNetAppLun)
	if err == nil {
		o.StorageBaseVolume = varStorageNetAppLun.StorageBaseVolume
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "OsType")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "Host")
		delete(additionalProperties, "StorageContainer")

		// remove fields from embedded structs
		reflectStorageBaseVolume := reflect.ValueOf(o.StorageBaseVolume)
		for i := 0; i < reflectStorageBaseVolume.Type().NumField(); i++ {
			t := reflectStorageBaseVolume.Type().Field(i)

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

type NullableStorageNetAppLun struct {
	value *StorageNetAppLun
	isSet bool
}

func (v NullableStorageNetAppLun) Get() *StorageNetAppLun {
	return v.value
}

func (v *NullableStorageNetAppLun) Set(val *StorageNetAppLun) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppLun) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppLun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppLun(val *StorageNetAppLun) *NullableStorageNetAppLun {
	return &NullableStorageNetAppLun{value: val, isSet: true}
}

func (v NullableStorageNetAppLun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppLun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
