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

// StoragePureReplicationSchedule Pure snapshot replication schedule entity.
type StoragePureReplicationSchedule struct {
	StorageBaseReplicationSchedule
	// Total number of snapshots per day to be available on target above and over the specified retention time. PureStorage FlashArray maintains all created snapshot until retention period. Daily limit is applied only on the snapshots once retention time is expired. In case of, daily limit is less than the number of snapshot available on source, system select snapshots evenly spaced out throughout the day.
	DailyLimit                   *int64                            `json:"DailyLimit,omitempty"`
	ReplicationBlackoutIntervals *[]StoragePureReplicationBlackout `json:"ReplicationBlackoutIntervals,omitempty"`
	// Duration to keep the daily limit snapshots on target array. StorageArray deletes the snapshots permanently from the targets beyond this period.
	SnapshotExpiryTime   *string                                 `json:"SnapshotExpiryTime,omitempty"`
	Array                *StoragePureArrayRelationship           `json:"Array,omitempty"`
	ProtectionGroup      *StoragePureProtectionGroupRelationship `json:"ProtectionGroup,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoragePureReplicationSchedule StoragePureReplicationSchedule

// NewStoragePureReplicationSchedule instantiates a new StoragePureReplicationSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureReplicationSchedule() *StoragePureReplicationSchedule {
	this := StoragePureReplicationSchedule{}
	return &this
}

// NewStoragePureReplicationScheduleWithDefaults instantiates a new StoragePureReplicationSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureReplicationScheduleWithDefaults() *StoragePureReplicationSchedule {
	this := StoragePureReplicationSchedule{}
	return &this
}

// GetDailyLimit returns the DailyLimit field value if set, zero value otherwise.
func (o *StoragePureReplicationSchedule) GetDailyLimit() int64 {
	if o == nil || o.DailyLimit == nil {
		var ret int64
		return ret
	}
	return *o.DailyLimit
}

// GetDailyLimitOk returns a tuple with the DailyLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureReplicationSchedule) GetDailyLimitOk() (*int64, bool) {
	if o == nil || o.DailyLimit == nil {
		return nil, false
	}
	return o.DailyLimit, true
}

// HasDailyLimit returns a boolean if a field has been set.
func (o *StoragePureReplicationSchedule) HasDailyLimit() bool {
	if o != nil && o.DailyLimit != nil {
		return true
	}

	return false
}

// SetDailyLimit gets a reference to the given int64 and assigns it to the DailyLimit field.
func (o *StoragePureReplicationSchedule) SetDailyLimit(v int64) {
	o.DailyLimit = &v
}

// GetReplicationBlackoutIntervals returns the ReplicationBlackoutIntervals field value if set, zero value otherwise.
func (o *StoragePureReplicationSchedule) GetReplicationBlackoutIntervals() []StoragePureReplicationBlackout {
	if o == nil || o.ReplicationBlackoutIntervals == nil {
		var ret []StoragePureReplicationBlackout
		return ret
	}
	return *o.ReplicationBlackoutIntervals
}

// GetReplicationBlackoutIntervalsOk returns a tuple with the ReplicationBlackoutIntervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureReplicationSchedule) GetReplicationBlackoutIntervalsOk() (*[]StoragePureReplicationBlackout, bool) {
	if o == nil || o.ReplicationBlackoutIntervals == nil {
		return nil, false
	}
	return o.ReplicationBlackoutIntervals, true
}

// HasReplicationBlackoutIntervals returns a boolean if a field has been set.
func (o *StoragePureReplicationSchedule) HasReplicationBlackoutIntervals() bool {
	if o != nil && o.ReplicationBlackoutIntervals != nil {
		return true
	}

	return false
}

// SetReplicationBlackoutIntervals gets a reference to the given []StoragePureReplicationBlackout and assigns it to the ReplicationBlackoutIntervals field.
func (o *StoragePureReplicationSchedule) SetReplicationBlackoutIntervals(v []StoragePureReplicationBlackout) {
	o.ReplicationBlackoutIntervals = &v
}

// GetSnapshotExpiryTime returns the SnapshotExpiryTime field value if set, zero value otherwise.
func (o *StoragePureReplicationSchedule) GetSnapshotExpiryTime() string {
	if o == nil || o.SnapshotExpiryTime == nil {
		var ret string
		return ret
	}
	return *o.SnapshotExpiryTime
}

// GetSnapshotExpiryTimeOk returns a tuple with the SnapshotExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureReplicationSchedule) GetSnapshotExpiryTimeOk() (*string, bool) {
	if o == nil || o.SnapshotExpiryTime == nil {
		return nil, false
	}
	return o.SnapshotExpiryTime, true
}

// HasSnapshotExpiryTime returns a boolean if a field has been set.
func (o *StoragePureReplicationSchedule) HasSnapshotExpiryTime() bool {
	if o != nil && o.SnapshotExpiryTime != nil {
		return true
	}

	return false
}

// SetSnapshotExpiryTime gets a reference to the given string and assigns it to the SnapshotExpiryTime field.
func (o *StoragePureReplicationSchedule) SetSnapshotExpiryTime(v string) {
	o.SnapshotExpiryTime = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StoragePureReplicationSchedule) GetArray() StoragePureArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureReplicationSchedule) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePureReplicationSchedule) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePureReplicationSchedule) SetArray(v StoragePureArrayRelationship) {
	o.Array = &v
}

// GetProtectionGroup returns the ProtectionGroup field value if set, zero value otherwise.
func (o *StoragePureReplicationSchedule) GetProtectionGroup() StoragePureProtectionGroupRelationship {
	if o == nil || o.ProtectionGroup == nil {
		var ret StoragePureProtectionGroupRelationship
		return ret
	}
	return *o.ProtectionGroup
}

// GetProtectionGroupOk returns a tuple with the ProtectionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureReplicationSchedule) GetProtectionGroupOk() (*StoragePureProtectionGroupRelationship, bool) {
	if o == nil || o.ProtectionGroup == nil {
		return nil, false
	}
	return o.ProtectionGroup, true
}

// HasProtectionGroup returns a boolean if a field has been set.
func (o *StoragePureReplicationSchedule) HasProtectionGroup() bool {
	if o != nil && o.ProtectionGroup != nil {
		return true
	}

	return false
}

// SetProtectionGroup gets a reference to the given StoragePureProtectionGroupRelationship and assigns it to the ProtectionGroup field.
func (o *StoragePureReplicationSchedule) SetProtectionGroup(v StoragePureProtectionGroupRelationship) {
	o.ProtectionGroup = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePureReplicationSchedule) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureReplicationSchedule) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePureReplicationSchedule) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePureReplicationSchedule) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StoragePureReplicationSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseReplicationSchedule, errStorageBaseReplicationSchedule := json.Marshal(o.StorageBaseReplicationSchedule)
	if errStorageBaseReplicationSchedule != nil {
		return []byte{}, errStorageBaseReplicationSchedule
	}
	errStorageBaseReplicationSchedule = json.Unmarshal([]byte(serializedStorageBaseReplicationSchedule), &toSerialize)
	if errStorageBaseReplicationSchedule != nil {
		return []byte{}, errStorageBaseReplicationSchedule
	}
	if o.DailyLimit != nil {
		toSerialize["DailyLimit"] = o.DailyLimit
	}
	if o.ReplicationBlackoutIntervals != nil {
		toSerialize["ReplicationBlackoutIntervals"] = o.ReplicationBlackoutIntervals
	}
	if o.SnapshotExpiryTime != nil {
		toSerialize["SnapshotExpiryTime"] = o.SnapshotExpiryTime
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.ProtectionGroup != nil {
		toSerialize["ProtectionGroup"] = o.ProtectionGroup
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoragePureReplicationSchedule) UnmarshalJSON(bytes []byte) (err error) {
	type StoragePureReplicationScheduleWithoutEmbeddedStruct struct {
		// Total number of snapshots per day to be available on target above and over the specified retention time. PureStorage FlashArray maintains all created snapshot until retention period. Daily limit is applied only on the snapshots once retention time is expired. In case of, daily limit is less than the number of snapshot available on source, system select snapshots evenly spaced out throughout the day.
		DailyLimit                   *int64                            `json:"DailyLimit,omitempty"`
		ReplicationBlackoutIntervals *[]StoragePureReplicationBlackout `json:"ReplicationBlackoutIntervals,omitempty"`
		// Duration to keep the daily limit snapshots on target array. StorageArray deletes the snapshots permanently from the targets beyond this period.
		SnapshotExpiryTime *string                                 `json:"SnapshotExpiryTime,omitempty"`
		Array              *StoragePureArrayRelationship           `json:"Array,omitempty"`
		ProtectionGroup    *StoragePureProtectionGroupRelationship `json:"ProtectionGroup,omitempty"`
		RegisteredDevice   *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	}

	varStoragePureReplicationScheduleWithoutEmbeddedStruct := StoragePureReplicationScheduleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStoragePureReplicationScheduleWithoutEmbeddedStruct)
	if err == nil {
		varStoragePureReplicationSchedule := _StoragePureReplicationSchedule{}
		varStoragePureReplicationSchedule.DailyLimit = varStoragePureReplicationScheduleWithoutEmbeddedStruct.DailyLimit
		varStoragePureReplicationSchedule.ReplicationBlackoutIntervals = varStoragePureReplicationScheduleWithoutEmbeddedStruct.ReplicationBlackoutIntervals
		varStoragePureReplicationSchedule.SnapshotExpiryTime = varStoragePureReplicationScheduleWithoutEmbeddedStruct.SnapshotExpiryTime
		varStoragePureReplicationSchedule.Array = varStoragePureReplicationScheduleWithoutEmbeddedStruct.Array
		varStoragePureReplicationSchedule.ProtectionGroup = varStoragePureReplicationScheduleWithoutEmbeddedStruct.ProtectionGroup
		varStoragePureReplicationSchedule.RegisteredDevice = varStoragePureReplicationScheduleWithoutEmbeddedStruct.RegisteredDevice
		*o = StoragePureReplicationSchedule(varStoragePureReplicationSchedule)
	} else {
		return err
	}

	varStoragePureReplicationSchedule := _StoragePureReplicationSchedule{}

	err = json.Unmarshal(bytes, &varStoragePureReplicationSchedule)
	if err == nil {
		o.StorageBaseReplicationSchedule = varStoragePureReplicationSchedule.StorageBaseReplicationSchedule
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "DailyLimit")
		delete(additionalProperties, "ReplicationBlackoutIntervals")
		delete(additionalProperties, "SnapshotExpiryTime")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "ProtectionGroup")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectStorageBaseReplicationSchedule := reflect.ValueOf(o.StorageBaseReplicationSchedule)
		for i := 0; i < reflectStorageBaseReplicationSchedule.Type().NumField(); i++ {
			t := reflectStorageBaseReplicationSchedule.Type().Field(i)

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

type NullableStoragePureReplicationSchedule struct {
	value *StoragePureReplicationSchedule
	isSet bool
}

func (v NullableStoragePureReplicationSchedule) Get() *StoragePureReplicationSchedule {
	return v.value
}

func (v *NullableStoragePureReplicationSchedule) Set(val *StoragePureReplicationSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureReplicationSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureReplicationSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureReplicationSchedule(val *StoragePureReplicationSchedule) *NullableStoragePureReplicationSchedule {
	return &NullableStoragePureReplicationSchedule{value: val, isSet: true}
}

func (v NullableStoragePureReplicationSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureReplicationSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}