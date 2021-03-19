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

// StorageBasePhysicalPort Common attributes for a physical port in a storage array.
type StorageBasePhysicalPort struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// ISCSI qualified name applicable for ethernet port. Example - 'iqn.2008-05.com.storage:fnm00151300643-514f0c50141faf05'.
	Iqn *string `json:"Iqn,omitempty"`
	// Name of the physical port available in storage array.
	Name *string `json:"Name,omitempty"`
	// Operational speed of physical port measured in Gbps.
	Speed *int64 `json:"Speed,omitempty"`
	// Status of storage array port. * `Unknown` - Component status is not available. * `Ok` - Component is healthy and no issues found. * `Degraded` - Functioning, but not at full capability due to a non-fatal failure. * `Critical` - Not functioning or requiring immediate attention. * `Offline` - Component is installed, but powered off. * `Identifying` - Component is in initialization process. * `NotAvailable` - Component is not installed or not available. * `Updating` - Software update is in progress. * `Unrecognized` - Component is not recognized. It may be because the component is not installed properly or it is not supported.
	Status *string `json:"Status,omitempty"`
	// Port type - possible values are FC, FCoE and iSCSI. * `FC` - Port supports fibre channel protocol. * `iSCSI` - Port supports iSCSI protocol. * `FCoE` - Port supports fibre channel over ethernet protocol.
	Type *string `json:"Type,omitempty"`
	// World wide name of FC port. It is a combination of WWNN and WWPN represented in 128 bit hexadecimal formatted with colon. Example: '51:4f:0c:50:14:1f:af:01:51:4f:0c:51:14:1f:af:01'.
	Wwn *string `json:"Wwn,omitempty"`
	// World wide node name of FC port. Represented in 64 bit hex digits, formatted with colon. Example - '51:4f:0c:50:14:1f:af:01'.
	Wwnn *string `json:"Wwnn,omitempty"`
	// World wide port name of FC port. Represented in 64 bit hex digits, formatted with colon. Example - '51:4f:0c:51:14:1f:af:01'.
	Wwpn                 *string `json:"Wwpn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBasePhysicalPort StorageBasePhysicalPort

// NewStorageBasePhysicalPort instantiates a new StorageBasePhysicalPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBasePhysicalPort(classId string, objectType string) *StorageBasePhysicalPort {
	this := StorageBasePhysicalPort{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "Unknown"
	this.Status = &status
	var type_ string = "FC"
	this.Type = &type_
	return &this
}

// NewStorageBasePhysicalPortWithDefaults instantiates a new StorageBasePhysicalPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBasePhysicalPortWithDefaults() *StorageBasePhysicalPort {
	this := StorageBasePhysicalPort{}
	var status string = "Unknown"
	this.Status = &status
	var type_ string = "FC"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBasePhysicalPort) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBasePhysicalPort) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBasePhysicalPort) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBasePhysicalPort) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBasePhysicalPort) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBasePhysicalPort) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIqn returns the Iqn field value if set, zero value otherwise.
func (o *StorageBasePhysicalPort) GetIqn() string {
	if o == nil || o.Iqn == nil {
		var ret string
		return ret
	}
	return *o.Iqn
}

// GetIqnOk returns a tuple with the Iqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBasePhysicalPort) GetIqnOk() (*string, bool) {
	if o == nil || o.Iqn == nil {
		return nil, false
	}
	return o.Iqn, true
}

// HasIqn returns a boolean if a field has been set.
func (o *StorageBasePhysicalPort) HasIqn() bool {
	if o != nil && o.Iqn != nil {
		return true
	}

	return false
}

// SetIqn gets a reference to the given string and assigns it to the Iqn field.
func (o *StorageBasePhysicalPort) SetIqn(v string) {
	o.Iqn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBasePhysicalPort) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBasePhysicalPort) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBasePhysicalPort) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBasePhysicalPort) SetName(v string) {
	o.Name = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *StorageBasePhysicalPort) GetSpeed() int64 {
	if o == nil || o.Speed == nil {
		var ret int64
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBasePhysicalPort) GetSpeedOk() (*int64, bool) {
	if o == nil || o.Speed == nil {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *StorageBasePhysicalPort) HasSpeed() bool {
	if o != nil && o.Speed != nil {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int64 and assigns it to the Speed field.
func (o *StorageBasePhysicalPort) SetSpeed(v int64) {
	o.Speed = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StorageBasePhysicalPort) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBasePhysicalPort) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StorageBasePhysicalPort) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StorageBasePhysicalPort) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StorageBasePhysicalPort) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBasePhysicalPort) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StorageBasePhysicalPort) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StorageBasePhysicalPort) SetType(v string) {
	o.Type = &v
}

// GetWwn returns the Wwn field value if set, zero value otherwise.
func (o *StorageBasePhysicalPort) GetWwn() string {
	if o == nil || o.Wwn == nil {
		var ret string
		return ret
	}
	return *o.Wwn
}

// GetWwnOk returns a tuple with the Wwn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBasePhysicalPort) GetWwnOk() (*string, bool) {
	if o == nil || o.Wwn == nil {
		return nil, false
	}
	return o.Wwn, true
}

// HasWwn returns a boolean if a field has been set.
func (o *StorageBasePhysicalPort) HasWwn() bool {
	if o != nil && o.Wwn != nil {
		return true
	}

	return false
}

// SetWwn gets a reference to the given string and assigns it to the Wwn field.
func (o *StorageBasePhysicalPort) SetWwn(v string) {
	o.Wwn = &v
}

// GetWwnn returns the Wwnn field value if set, zero value otherwise.
func (o *StorageBasePhysicalPort) GetWwnn() string {
	if o == nil || o.Wwnn == nil {
		var ret string
		return ret
	}
	return *o.Wwnn
}

// GetWwnnOk returns a tuple with the Wwnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBasePhysicalPort) GetWwnnOk() (*string, bool) {
	if o == nil || o.Wwnn == nil {
		return nil, false
	}
	return o.Wwnn, true
}

// HasWwnn returns a boolean if a field has been set.
func (o *StorageBasePhysicalPort) HasWwnn() bool {
	if o != nil && o.Wwnn != nil {
		return true
	}

	return false
}

// SetWwnn gets a reference to the given string and assigns it to the Wwnn field.
func (o *StorageBasePhysicalPort) SetWwnn(v string) {
	o.Wwnn = &v
}

// GetWwpn returns the Wwpn field value if set, zero value otherwise.
func (o *StorageBasePhysicalPort) GetWwpn() string {
	if o == nil || o.Wwpn == nil {
		var ret string
		return ret
	}
	return *o.Wwpn
}

// GetWwpnOk returns a tuple with the Wwpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBasePhysicalPort) GetWwpnOk() (*string, bool) {
	if o == nil || o.Wwpn == nil {
		return nil, false
	}
	return o.Wwpn, true
}

// HasWwpn returns a boolean if a field has been set.
func (o *StorageBasePhysicalPort) HasWwpn() bool {
	if o != nil && o.Wwpn != nil {
		return true
	}

	return false
}

// SetWwpn gets a reference to the given string and assigns it to the Wwpn field.
func (o *StorageBasePhysicalPort) SetWwpn(v string) {
	o.Wwpn = &v
}

func (o StorageBasePhysicalPort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Iqn != nil {
		toSerialize["Iqn"] = o.Iqn
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Speed != nil {
		toSerialize["Speed"] = o.Speed
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Wwn != nil {
		toSerialize["Wwn"] = o.Wwn
	}
	if o.Wwnn != nil {
		toSerialize["Wwnn"] = o.Wwnn
	}
	if o.Wwpn != nil {
		toSerialize["Wwpn"] = o.Wwpn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBasePhysicalPort) UnmarshalJSON(bytes []byte) (err error) {
	type StorageBasePhysicalPortWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// ISCSI qualified name applicable for ethernet port. Example - 'iqn.2008-05.com.storage:fnm00151300643-514f0c50141faf05'.
		Iqn *string `json:"Iqn,omitempty"`
		// Name of the physical port available in storage array.
		Name *string `json:"Name,omitempty"`
		// Operational speed of physical port measured in Gbps.
		Speed *int64 `json:"Speed,omitempty"`
		// Status of storage array port. * `Unknown` - Component status is not available. * `Ok` - Component is healthy and no issues found. * `Degraded` - Functioning, but not at full capability due to a non-fatal failure. * `Critical` - Not functioning or requiring immediate attention. * `Offline` - Component is installed, but powered off. * `Identifying` - Component is in initialization process. * `NotAvailable` - Component is not installed or not available. * `Updating` - Software update is in progress. * `Unrecognized` - Component is not recognized. It may be because the component is not installed properly or it is not supported.
		Status *string `json:"Status,omitempty"`
		// Port type - possible values are FC, FCoE and iSCSI. * `FC` - Port supports fibre channel protocol. * `iSCSI` - Port supports iSCSI protocol. * `FCoE` - Port supports fibre channel over ethernet protocol.
		Type *string `json:"Type,omitempty"`
		// World wide name of FC port. It is a combination of WWNN and WWPN represented in 128 bit hexadecimal formatted with colon. Example: '51:4f:0c:50:14:1f:af:01:51:4f:0c:51:14:1f:af:01'.
		Wwn *string `json:"Wwn,omitempty"`
		// World wide node name of FC port. Represented in 64 bit hex digits, formatted with colon. Example - '51:4f:0c:50:14:1f:af:01'.
		Wwnn *string `json:"Wwnn,omitempty"`
		// World wide port name of FC port. Represented in 64 bit hex digits, formatted with colon. Example - '51:4f:0c:51:14:1f:af:01'.
		Wwpn *string `json:"Wwpn,omitempty"`
	}

	varStorageBasePhysicalPortWithoutEmbeddedStruct := StorageBasePhysicalPortWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageBasePhysicalPortWithoutEmbeddedStruct)
	if err == nil {
		varStorageBasePhysicalPort := _StorageBasePhysicalPort{}
		varStorageBasePhysicalPort.ClassId = varStorageBasePhysicalPortWithoutEmbeddedStruct.ClassId
		varStorageBasePhysicalPort.ObjectType = varStorageBasePhysicalPortWithoutEmbeddedStruct.ObjectType
		varStorageBasePhysicalPort.Iqn = varStorageBasePhysicalPortWithoutEmbeddedStruct.Iqn
		varStorageBasePhysicalPort.Name = varStorageBasePhysicalPortWithoutEmbeddedStruct.Name
		varStorageBasePhysicalPort.Speed = varStorageBasePhysicalPortWithoutEmbeddedStruct.Speed
		varStorageBasePhysicalPort.Status = varStorageBasePhysicalPortWithoutEmbeddedStruct.Status
		varStorageBasePhysicalPort.Type = varStorageBasePhysicalPortWithoutEmbeddedStruct.Type
		varStorageBasePhysicalPort.Wwn = varStorageBasePhysicalPortWithoutEmbeddedStruct.Wwn
		varStorageBasePhysicalPort.Wwnn = varStorageBasePhysicalPortWithoutEmbeddedStruct.Wwnn
		varStorageBasePhysicalPort.Wwpn = varStorageBasePhysicalPortWithoutEmbeddedStruct.Wwpn
		*o = StorageBasePhysicalPort(varStorageBasePhysicalPort)
	} else {
		return err
	}

	varStorageBasePhysicalPort := _StorageBasePhysicalPort{}

	err = json.Unmarshal(bytes, &varStorageBasePhysicalPort)
	if err == nil {
		o.MoBaseMo = varStorageBasePhysicalPort.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Iqn")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Speed")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Wwn")
		delete(additionalProperties, "Wwnn")
		delete(additionalProperties, "Wwpn")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableStorageBasePhysicalPort struct {
	value *StorageBasePhysicalPort
	isSet bool
}

func (v NullableStorageBasePhysicalPort) Get() *StorageBasePhysicalPort {
	return v.value
}

func (v *NullableStorageBasePhysicalPort) Set(val *StorageBasePhysicalPort) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBasePhysicalPort) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBasePhysicalPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBasePhysicalPort(val *StorageBasePhysicalPort) *NullableStorageBasePhysicalPort {
	return &NullableStorageBasePhysicalPort{value: val, isSet: true}
}

func (v NullableStorageBasePhysicalPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBasePhysicalPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
