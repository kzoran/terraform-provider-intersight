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
	"fmt"
)

// SoftwarerepositoryOperatingSystemFileRelationship - A relationship to the 'softwarerepository.OperatingSystemFile' resource, or the expanded 'softwarerepository.OperatingSystemFile' resource, or the 'null' value.
type SoftwarerepositoryOperatingSystemFileRelationship struct {
	MoMoRef                               *MoMoRef
	SoftwarerepositoryOperatingSystemFile *SoftwarerepositoryOperatingSystemFile
}

// MoMoRefAsSoftwarerepositoryOperatingSystemFileRelationship is a convenience function that returns MoMoRef wrapped in SoftwarerepositoryOperatingSystemFileRelationship
func MoMoRefAsSoftwarerepositoryOperatingSystemFileRelationship(v *MoMoRef) SoftwarerepositoryOperatingSystemFileRelationship {
	return SoftwarerepositoryOperatingSystemFileRelationship{MoMoRef: v}
}

// SoftwarerepositoryOperatingSystemFileAsSoftwarerepositoryOperatingSystemFileRelationship is a convenience function that returns SoftwarerepositoryOperatingSystemFile wrapped in SoftwarerepositoryOperatingSystemFileRelationship
func SoftwarerepositoryOperatingSystemFileAsSoftwarerepositoryOperatingSystemFileRelationship(v *SoftwarerepositoryOperatingSystemFile) SoftwarerepositoryOperatingSystemFileRelationship {
	return SoftwarerepositoryOperatingSystemFileRelationship{SoftwarerepositoryOperatingSystemFile: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SoftwarerepositoryOperatingSystemFileRelationship) UnmarshalJSON(data []byte) error {
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
			return fmt.Errorf("Failed to unmarshal SoftwarerepositoryOperatingSystemFileRelationship as MoMoRef: %s", err.Error())
		}
	}

	// check if the discriminator value is 'softwarerepository.OperatingSystemFile'
	if jsonDict["ClassId"] == "softwarerepository.OperatingSystemFile" {
		// try to unmarshal JSON data into SoftwarerepositoryOperatingSystemFile
		err = json.Unmarshal(data, &dst.SoftwarerepositoryOperatingSystemFile)
		if err == nil {
			return nil // data stored in dst.SoftwarerepositoryOperatingSystemFile, return on the first match
		} else {
			dst.SoftwarerepositoryOperatingSystemFile = nil
			return fmt.Errorf("Failed to unmarshal SoftwarerepositoryOperatingSystemFileRelationship as SoftwarerepositoryOperatingSystemFile: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SoftwarerepositoryOperatingSystemFileRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.SoftwarerepositoryOperatingSystemFile != nil {
		return json.Marshal(&src.SoftwarerepositoryOperatingSystemFile)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SoftwarerepositoryOperatingSystemFileRelationship) GetActualInstance() interface{} {
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.SoftwarerepositoryOperatingSystemFile != nil {
		return obj.SoftwarerepositoryOperatingSystemFile
	}

	// all schemas are nil
	return nil
}

type NullableSoftwarerepositoryOperatingSystemFileRelationship struct {
	value *SoftwarerepositoryOperatingSystemFileRelationship
	isSet bool
}

func (v NullableSoftwarerepositoryOperatingSystemFileRelationship) Get() *SoftwarerepositoryOperatingSystemFileRelationship {
	return v.value
}

func (v *NullableSoftwarerepositoryOperatingSystemFileRelationship) Set(val *SoftwarerepositoryOperatingSystemFileRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryOperatingSystemFileRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryOperatingSystemFileRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryOperatingSystemFileRelationship(val *SoftwarerepositoryOperatingSystemFileRelationship) *NullableSoftwarerepositoryOperatingSystemFileRelationship {
	return &NullableSoftwarerepositoryOperatingSystemFileRelationship{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryOperatingSystemFileRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryOperatingSystemFileRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
