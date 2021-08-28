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

// FabricTransceiverRole Configuration object sent by user to apply on a transceiver.
type FabricTransceiverRole struct {
	FabricPortRole
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Admin configured speed for the port. * `Auto` - Admin configurable speed AUTO ( default ). * `1Gbps` - Admin configurable speed 1Gbps. * `10Gbps` - Admin configurable speed 10Gbps. * `25Gbps` - Admin configurable speed 25Gbps. * `40Gbps` - Admin configurable speed 40Gbps. * `100Gbps` - Admin configurable speed 100Gbps.
	AdminSpeed *string `json:"AdminSpeed,omitempty"`
	// Forward error correction configuration for the port. * `Auto` - Forward error correction option 'Auto'. * `Cl91` - Forward error correction option 'cl91'. * `Cl74` - Forward error correction option 'cl74'.
	Fec                  *string `json:"Fec,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricTransceiverRole FabricTransceiverRole

// NewFabricTransceiverRole instantiates a new FabricTransceiverRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricTransceiverRole(classId string, objectType string) *FabricTransceiverRole {
	this := FabricTransceiverRole{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	var fec string = "Auto"
	this.Fec = &fec
	return &this
}

// NewFabricTransceiverRoleWithDefaults instantiates a new FabricTransceiverRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricTransceiverRoleWithDefaults() *FabricTransceiverRole {
	this := FabricTransceiverRole{}
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	var fec string = "Auto"
	this.Fec = &fec
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricTransceiverRole) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricTransceiverRole) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricTransceiverRole) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricTransceiverRole) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricTransceiverRole) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricTransceiverRole) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminSpeed returns the AdminSpeed field value if set, zero value otherwise.
func (o *FabricTransceiverRole) GetAdminSpeed() string {
	if o == nil || o.AdminSpeed == nil {
		var ret string
		return ret
	}
	return *o.AdminSpeed
}

// GetAdminSpeedOk returns a tuple with the AdminSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricTransceiverRole) GetAdminSpeedOk() (*string, bool) {
	if o == nil || o.AdminSpeed == nil {
		return nil, false
	}
	return o.AdminSpeed, true
}

// HasAdminSpeed returns a boolean if a field has been set.
func (o *FabricTransceiverRole) HasAdminSpeed() bool {
	if o != nil && o.AdminSpeed != nil {
		return true
	}

	return false
}

// SetAdminSpeed gets a reference to the given string and assigns it to the AdminSpeed field.
func (o *FabricTransceiverRole) SetAdminSpeed(v string) {
	o.AdminSpeed = &v
}

// GetFec returns the Fec field value if set, zero value otherwise.
func (o *FabricTransceiverRole) GetFec() string {
	if o == nil || o.Fec == nil {
		var ret string
		return ret
	}
	return *o.Fec
}

// GetFecOk returns a tuple with the Fec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricTransceiverRole) GetFecOk() (*string, bool) {
	if o == nil || o.Fec == nil {
		return nil, false
	}
	return o.Fec, true
}

// HasFec returns a boolean if a field has been set.
func (o *FabricTransceiverRole) HasFec() bool {
	if o != nil && o.Fec != nil {
		return true
	}

	return false
}

// SetFec gets a reference to the given string and assigns it to the Fec field.
func (o *FabricTransceiverRole) SetFec(v string) {
	o.Fec = &v
}

func (o FabricTransceiverRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricPortRole, errFabricPortRole := json.Marshal(o.FabricPortRole)
	if errFabricPortRole != nil {
		return []byte{}, errFabricPortRole
	}
	errFabricPortRole = json.Unmarshal([]byte(serializedFabricPortRole), &toSerialize)
	if errFabricPortRole != nil {
		return []byte{}, errFabricPortRole
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminSpeed != nil {
		toSerialize["AdminSpeed"] = o.AdminSpeed
	}
	if o.Fec != nil {
		toSerialize["Fec"] = o.Fec
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricTransceiverRole) UnmarshalJSON(bytes []byte) (err error) {
	type FabricTransceiverRoleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Admin configured speed for the port. * `Auto` - Admin configurable speed AUTO ( default ). * `1Gbps` - Admin configurable speed 1Gbps. * `10Gbps` - Admin configurable speed 10Gbps. * `25Gbps` - Admin configurable speed 25Gbps. * `40Gbps` - Admin configurable speed 40Gbps. * `100Gbps` - Admin configurable speed 100Gbps.
		AdminSpeed *string `json:"AdminSpeed,omitempty"`
		// Forward error correction configuration for the port. * `Auto` - Forward error correction option 'Auto'. * `Cl91` - Forward error correction option 'cl91'. * `Cl74` - Forward error correction option 'cl74'.
		Fec *string `json:"Fec,omitempty"`
	}

	varFabricTransceiverRoleWithoutEmbeddedStruct := FabricTransceiverRoleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricTransceiverRoleWithoutEmbeddedStruct)
	if err == nil {
		varFabricTransceiverRole := _FabricTransceiverRole{}
		varFabricTransceiverRole.ClassId = varFabricTransceiverRoleWithoutEmbeddedStruct.ClassId
		varFabricTransceiverRole.ObjectType = varFabricTransceiverRoleWithoutEmbeddedStruct.ObjectType
		varFabricTransceiverRole.AdminSpeed = varFabricTransceiverRoleWithoutEmbeddedStruct.AdminSpeed
		varFabricTransceiverRole.Fec = varFabricTransceiverRoleWithoutEmbeddedStruct.Fec
		*o = FabricTransceiverRole(varFabricTransceiverRole)
	} else {
		return err
	}

	varFabricTransceiverRole := _FabricTransceiverRole{}

	err = json.Unmarshal(bytes, &varFabricTransceiverRole)
	if err == nil {
		o.FabricPortRole = varFabricTransceiverRole.FabricPortRole
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminSpeed")
		delete(additionalProperties, "Fec")

		// remove fields from embedded structs
		reflectFabricPortRole := reflect.ValueOf(o.FabricPortRole)
		for i := 0; i < reflectFabricPortRole.Type().NumField(); i++ {
			t := reflectFabricPortRole.Type().Field(i)

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

type NullableFabricTransceiverRole struct {
	value *FabricTransceiverRole
	isSet bool
}

func (v NullableFabricTransceiverRole) Get() *FabricTransceiverRole {
	return v.value
}

func (v *NullableFabricTransceiverRole) Set(val *FabricTransceiverRole) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricTransceiverRole) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricTransceiverRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricTransceiverRole(val *FabricTransceiverRole) *NullableFabricTransceiverRole {
	return &NullableFabricTransceiverRole{value: val, isSet: true}
}

func (v NullableFabricTransceiverRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricTransceiverRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
