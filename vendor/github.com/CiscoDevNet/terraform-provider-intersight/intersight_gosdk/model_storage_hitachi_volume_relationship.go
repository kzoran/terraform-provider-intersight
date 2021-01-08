/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-01-04T05:15:49Z.
 *
 * API version: 1.0.9-3144
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// StorageHitachiVolumeRelationship - A relationship to the 'storage.HitachiVolume' resource, or the expanded 'storage.HitachiVolume' resource, or the 'null' value.
type StorageHitachiVolumeRelationship struct {
	MoMoRef              *MoMoRef
	StorageHitachiVolume *StorageHitachiVolume
}

// MoMoRefAsStorageHitachiVolumeRelationship is a convenience function that returns MoMoRef wrapped in StorageHitachiVolumeRelationship
func MoMoRefAsStorageHitachiVolumeRelationship(v *MoMoRef) StorageHitachiVolumeRelationship {
	return StorageHitachiVolumeRelationship{MoMoRef: v}
}

// StorageHitachiVolumeAsStorageHitachiVolumeRelationship is a convenience function that returns StorageHitachiVolume wrapped in StorageHitachiVolumeRelationship
func StorageHitachiVolumeAsStorageHitachiVolumeRelationship(v *StorageHitachiVolume) StorageHitachiVolumeRelationship {
	return StorageHitachiVolumeRelationship{StorageHitachiVolume: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *StorageHitachiVolumeRelationship) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			return nil // data stored in dst.MoMoRef, return on the first match
		} else {
			dst.MoMoRef = nil
			return fmt.Errorf("Failed to unmarshal StorageHitachiVolumeRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'storage.HitachiVolume'
	if jsonDict["ClassId"] == "storage.HitachiVolume" {
		// try to unmarshal JSON data into StorageHitachiVolume
		err = json.Unmarshal(data, &dst.StorageHitachiVolume)
		if err == nil {
			return nil // data stored in dst.StorageHitachiVolume, return on the first match
		} else {
			dst.StorageHitachiVolume = nil
			return fmt.Errorf("Failed to unmarshal StorageHitachiVolumeRelationship as StorageHitachiVolume: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StorageHitachiVolumeRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.StorageHitachiVolume != nil {
		return json.Marshal(&src.StorageHitachiVolume)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StorageHitachiVolumeRelationship) GetActualInstance() interface{} {
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.StorageHitachiVolume != nil {
		return obj.StorageHitachiVolume
	}

	// all schemas are nil
	return nil
}

type NullableStorageHitachiVolumeRelationship struct {
	value *StorageHitachiVolumeRelationship
	isSet bool
}

func (v NullableStorageHitachiVolumeRelationship) Get() *StorageHitachiVolumeRelationship {
	return v.value
}

func (v *NullableStorageHitachiVolumeRelationship) Set(val *StorageHitachiVolumeRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiVolumeRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiVolumeRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiVolumeRelationship(val *StorageHitachiVolumeRelationship) *NullableStorageHitachiVolumeRelationship {
	return &NullableStorageHitachiVolumeRelationship{value: val, isSet: true}
}

func (v NullableStorageHitachiVolumeRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiVolumeRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
