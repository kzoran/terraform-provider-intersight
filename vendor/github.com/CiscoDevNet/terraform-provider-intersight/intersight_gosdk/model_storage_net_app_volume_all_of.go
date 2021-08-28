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

// StorageNetAppVolumeAllOf Definition of the list of properties defined in 'storage.NetAppVolume', excluding properties defined in parent classes.
type StorageNetAppVolumeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The autosize mode for NetApp Volume. Modes can be off or grow or grow_shrink. * `off` - The volume will not grow or shrink in size in response to the amount of used space. * `grow` - The volume will automatically grow when used space in the volume is above the grow threshold. * `grow_shrink` - The volume will grow or shrink in size in response to the amount of used space.
	AutosizeMode *string `json:"AutosizeMode,omitempty"`
	// Name of the Export Policy associated with the volume.
	ExportPolicyName *string `json:"ExportPolicyName,omitempty"`
	// Name of the snapshot policy.
	SnapshotPolicyName *string `json:"SnapshotPolicyName,omitempty"`
	// Uuid of the snapshot policy.
	SnapshotPolicyUuid *string `json:"SnapshotPolicyUuid,omitempty"`
	// The total space used by snapshot copies in the volume represented in bytes.
	SnapshotUtilizedCapacity *int64 `json:"SnapshotUtilizedCapacity,omitempty"`
	// The current state of a NetApp volume. * `offline` - Read and write access to the volume is not allowed. * `online` - Read and write access to the volume is allowed. * `error` - Storage volume state of error type. * `mixed` - The constituents of a FlexGroup volume are not all in the same state.
	State *string `json:"State,omitempty"`
	// NetApp volume type. The volume type can be Read-write or Data-protection, Load-sharing, or Data-cache. * `data-protection` - Prevents modification of the data on the Volume. * `read-write` - Data on the Volume can be modified. * `load-sharing` - Load Sharing.
	Type *string `json:"Type,omitempty"`
	// UUID of NetApp Volume.
	Uuid  *string                           `json:"Uuid,omitempty"`
	Array *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	// An array of relationships to storageNetAppAggregate resources.
	DiskPool             []StorageNetAppAggregateRelationship `json:"DiskPool,omitempty"`
	Tenant               *StorageNetAppStorageVmRelationship  `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppVolumeAllOf StorageNetAppVolumeAllOf

// NewStorageNetAppVolumeAllOf instantiates a new StorageNetAppVolumeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppVolumeAllOf(classId string, objectType string) *StorageNetAppVolumeAllOf {
	this := StorageNetAppVolumeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppVolumeAllOfWithDefaults instantiates a new StorageNetAppVolumeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppVolumeAllOfWithDefaults() *StorageNetAppVolumeAllOf {
	this := StorageNetAppVolumeAllOf{}
	var classId string = "storage.NetAppVolume"
	this.ClassId = classId
	var objectType string = "storage.NetAppVolume"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppVolumeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppVolumeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppVolumeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppVolumeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAutosizeMode returns the AutosizeMode field value if set, zero value otherwise.
func (o *StorageNetAppVolumeAllOf) GetAutosizeMode() string {
	if o == nil || o.AutosizeMode == nil {
		var ret string
		return ret
	}
	return *o.AutosizeMode
}

// GetAutosizeModeOk returns a tuple with the AutosizeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeAllOf) GetAutosizeModeOk() (*string, bool) {
	if o == nil || o.AutosizeMode == nil {
		return nil, false
	}
	return o.AutosizeMode, true
}

// HasAutosizeMode returns a boolean if a field has been set.
func (o *StorageNetAppVolumeAllOf) HasAutosizeMode() bool {
	if o != nil && o.AutosizeMode != nil {
		return true
	}

	return false
}

// SetAutosizeMode gets a reference to the given string and assigns it to the AutosizeMode field.
func (o *StorageNetAppVolumeAllOf) SetAutosizeMode(v string) {
	o.AutosizeMode = &v
}

// GetExportPolicyName returns the ExportPolicyName field value if set, zero value otherwise.
func (o *StorageNetAppVolumeAllOf) GetExportPolicyName() string {
	if o == nil || o.ExportPolicyName == nil {
		var ret string
		return ret
	}
	return *o.ExportPolicyName
}

// GetExportPolicyNameOk returns a tuple with the ExportPolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeAllOf) GetExportPolicyNameOk() (*string, bool) {
	if o == nil || o.ExportPolicyName == nil {
		return nil, false
	}
	return o.ExportPolicyName, true
}

// HasExportPolicyName returns a boolean if a field has been set.
func (o *StorageNetAppVolumeAllOf) HasExportPolicyName() bool {
	if o != nil && o.ExportPolicyName != nil {
		return true
	}

	return false
}

// SetExportPolicyName gets a reference to the given string and assigns it to the ExportPolicyName field.
func (o *StorageNetAppVolumeAllOf) SetExportPolicyName(v string) {
	o.ExportPolicyName = &v
}

// GetSnapshotPolicyName returns the SnapshotPolicyName field value if set, zero value otherwise.
func (o *StorageNetAppVolumeAllOf) GetSnapshotPolicyName() string {
	if o == nil || o.SnapshotPolicyName == nil {
		var ret string
		return ret
	}
	return *o.SnapshotPolicyName
}

// GetSnapshotPolicyNameOk returns a tuple with the SnapshotPolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeAllOf) GetSnapshotPolicyNameOk() (*string, bool) {
	if o == nil || o.SnapshotPolicyName == nil {
		return nil, false
	}
	return o.SnapshotPolicyName, true
}

// HasSnapshotPolicyName returns a boolean if a field has been set.
func (o *StorageNetAppVolumeAllOf) HasSnapshotPolicyName() bool {
	if o != nil && o.SnapshotPolicyName != nil {
		return true
	}

	return false
}

// SetSnapshotPolicyName gets a reference to the given string and assigns it to the SnapshotPolicyName field.
func (o *StorageNetAppVolumeAllOf) SetSnapshotPolicyName(v string) {
	o.SnapshotPolicyName = &v
}

// GetSnapshotPolicyUuid returns the SnapshotPolicyUuid field value if set, zero value otherwise.
func (o *StorageNetAppVolumeAllOf) GetSnapshotPolicyUuid() string {
	if o == nil || o.SnapshotPolicyUuid == nil {
		var ret string
		return ret
	}
	return *o.SnapshotPolicyUuid
}

// GetSnapshotPolicyUuidOk returns a tuple with the SnapshotPolicyUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeAllOf) GetSnapshotPolicyUuidOk() (*string, bool) {
	if o == nil || o.SnapshotPolicyUuid == nil {
		return nil, false
	}
	return o.SnapshotPolicyUuid, true
}

// HasSnapshotPolicyUuid returns a boolean if a field has been set.
func (o *StorageNetAppVolumeAllOf) HasSnapshotPolicyUuid() bool {
	if o != nil && o.SnapshotPolicyUuid != nil {
		return true
	}

	return false
}

// SetSnapshotPolicyUuid gets a reference to the given string and assigns it to the SnapshotPolicyUuid field.
func (o *StorageNetAppVolumeAllOf) SetSnapshotPolicyUuid(v string) {
	o.SnapshotPolicyUuid = &v
}

// GetSnapshotUtilizedCapacity returns the SnapshotUtilizedCapacity field value if set, zero value otherwise.
func (o *StorageNetAppVolumeAllOf) GetSnapshotUtilizedCapacity() int64 {
	if o == nil || o.SnapshotUtilizedCapacity == nil {
		var ret int64
		return ret
	}
	return *o.SnapshotUtilizedCapacity
}

// GetSnapshotUtilizedCapacityOk returns a tuple with the SnapshotUtilizedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeAllOf) GetSnapshotUtilizedCapacityOk() (*int64, bool) {
	if o == nil || o.SnapshotUtilizedCapacity == nil {
		return nil, false
	}
	return o.SnapshotUtilizedCapacity, true
}

// HasSnapshotUtilizedCapacity returns a boolean if a field has been set.
func (o *StorageNetAppVolumeAllOf) HasSnapshotUtilizedCapacity() bool {
	if o != nil && o.SnapshotUtilizedCapacity != nil {
		return true
	}

	return false
}

// SetSnapshotUtilizedCapacity gets a reference to the given int64 and assigns it to the SnapshotUtilizedCapacity field.
func (o *StorageNetAppVolumeAllOf) SetSnapshotUtilizedCapacity(v int64) {
	o.SnapshotUtilizedCapacity = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageNetAppVolumeAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageNetAppVolumeAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageNetAppVolumeAllOf) SetState(v string) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StorageNetAppVolumeAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StorageNetAppVolumeAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StorageNetAppVolumeAllOf) SetType(v string) {
	o.Type = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppVolumeAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppVolumeAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppVolumeAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppVolumeAllOf) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppVolumeAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppVolumeAllOf) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

// GetDiskPool returns the DiskPool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppVolumeAllOf) GetDiskPool() []StorageNetAppAggregateRelationship {
	if o == nil {
		var ret []StorageNetAppAggregateRelationship
		return ret
	}
	return o.DiskPool
}

// GetDiskPoolOk returns a tuple with the DiskPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppVolumeAllOf) GetDiskPoolOk() (*[]StorageNetAppAggregateRelationship, bool) {
	if o == nil || o.DiskPool == nil {
		return nil, false
	}
	return &o.DiskPool, true
}

// HasDiskPool returns a boolean if a field has been set.
func (o *StorageNetAppVolumeAllOf) HasDiskPool() bool {
	if o != nil && o.DiskPool != nil {
		return true
	}

	return false
}

// SetDiskPool gets a reference to the given []StorageNetAppAggregateRelationship and assigns it to the DiskPool field.
func (o *StorageNetAppVolumeAllOf) SetDiskPool(v []StorageNetAppAggregateRelationship) {
	o.DiskPool = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *StorageNetAppVolumeAllOf) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || o.Tenant == nil {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeAllOf) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppVolumeAllOf) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given StorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppVolumeAllOf) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant = &v
}

func (o StorageNetAppVolumeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AutosizeMode != nil {
		toSerialize["AutosizeMode"] = o.AutosizeMode
	}
	if o.ExportPolicyName != nil {
		toSerialize["ExportPolicyName"] = o.ExportPolicyName
	}
	if o.SnapshotPolicyName != nil {
		toSerialize["SnapshotPolicyName"] = o.SnapshotPolicyName
	}
	if o.SnapshotPolicyUuid != nil {
		toSerialize["SnapshotPolicyUuid"] = o.SnapshotPolicyUuid
	}
	if o.SnapshotUtilizedCapacity != nil {
		toSerialize["SnapshotUtilizedCapacity"] = o.SnapshotUtilizedCapacity
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.DiskPool != nil {
		toSerialize["DiskPool"] = o.DiskPool
	}
	if o.Tenant != nil {
		toSerialize["Tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppVolumeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppVolumeAllOf := _StorageNetAppVolumeAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppVolumeAllOf); err == nil {
		*o = StorageNetAppVolumeAllOf(varStorageNetAppVolumeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AutosizeMode")
		delete(additionalProperties, "ExportPolicyName")
		delete(additionalProperties, "SnapshotPolicyName")
		delete(additionalProperties, "SnapshotPolicyUuid")
		delete(additionalProperties, "SnapshotUtilizedCapacity")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "DiskPool")
		delete(additionalProperties, "Tenant")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppVolumeAllOf struct {
	value *StorageNetAppVolumeAllOf
	isSet bool
}

func (v NullableStorageNetAppVolumeAllOf) Get() *StorageNetAppVolumeAllOf {
	return v.value
}

func (v *NullableStorageNetAppVolumeAllOf) Set(val *StorageNetAppVolumeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppVolumeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppVolumeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppVolumeAllOf(val *StorageNetAppVolumeAllOf) *NullableStorageNetAppVolumeAllOf {
	return &NullableStorageNetAppVolumeAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppVolumeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppVolumeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
