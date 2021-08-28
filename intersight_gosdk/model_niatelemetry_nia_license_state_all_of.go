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

// NiatelemetryNiaLicenseStateAllOf Definition of the list of properties defined in 'niatelemetry.NiaLicenseState', excluding properties defined in parent classes.
type NiatelemetryNiaLicenseStateAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Features activated on device being inventoried. This determines which features are currently enabled on the device that the license API can check.
	FeatureActivated *string `json:"FeatureActivated,omitempty"`
	// Licenses activated on device being inventoried. This determines which lienceses are currently enabled on the device.
	LicenseActivated *string `json:"LicenseActivated,omitempty"`
	// PID of device being inventoried. This determines the hardware model type of the device.
	PidType *string `json:"PidType,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Serial number of device being inventoried. The serial number is unique per device.
	Serial *string `json:"Serial,omitempty"`
	// Name of fabric domain of the controller.
	SiteName             *string                               `json:"SiteName,omitempty"`
	Device               *NiatelemetryNiaInventoryRelationship `json:"Device,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNiaLicenseStateAllOf NiatelemetryNiaLicenseStateAllOf

// NewNiatelemetryNiaLicenseStateAllOf instantiates a new NiatelemetryNiaLicenseStateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNiaLicenseStateAllOf(classId string, objectType string) *NiatelemetryNiaLicenseStateAllOf {
	this := NiatelemetryNiaLicenseStateAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNiaLicenseStateAllOfWithDefaults instantiates a new NiatelemetryNiaLicenseStateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNiaLicenseStateAllOfWithDefaults() *NiatelemetryNiaLicenseStateAllOf {
	this := NiatelemetryNiaLicenseStateAllOf{}
	var classId string = "niatelemetry.NiaLicenseState"
	this.ClassId = classId
	var objectType string = "niatelemetry.NiaLicenseState"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNiaLicenseStateAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaLicenseStateAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNiaLicenseStateAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNiaLicenseStateAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaLicenseStateAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNiaLicenseStateAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFeatureActivated returns the FeatureActivated field value if set, zero value otherwise.
func (o *NiatelemetryNiaLicenseStateAllOf) GetFeatureActivated() string {
	if o == nil || o.FeatureActivated == nil {
		var ret string
		return ret
	}
	return *o.FeatureActivated
}

// GetFeatureActivatedOk returns a tuple with the FeatureActivated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaLicenseStateAllOf) GetFeatureActivatedOk() (*string, bool) {
	if o == nil || o.FeatureActivated == nil {
		return nil, false
	}
	return o.FeatureActivated, true
}

// HasFeatureActivated returns a boolean if a field has been set.
func (o *NiatelemetryNiaLicenseStateAllOf) HasFeatureActivated() bool {
	if o != nil && o.FeatureActivated != nil {
		return true
	}

	return false
}

// SetFeatureActivated gets a reference to the given string and assigns it to the FeatureActivated field.
func (o *NiatelemetryNiaLicenseStateAllOf) SetFeatureActivated(v string) {
	o.FeatureActivated = &v
}

// GetLicenseActivated returns the LicenseActivated field value if set, zero value otherwise.
func (o *NiatelemetryNiaLicenseStateAllOf) GetLicenseActivated() string {
	if o == nil || o.LicenseActivated == nil {
		var ret string
		return ret
	}
	return *o.LicenseActivated
}

// GetLicenseActivatedOk returns a tuple with the LicenseActivated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaLicenseStateAllOf) GetLicenseActivatedOk() (*string, bool) {
	if o == nil || o.LicenseActivated == nil {
		return nil, false
	}
	return o.LicenseActivated, true
}

// HasLicenseActivated returns a boolean if a field has been set.
func (o *NiatelemetryNiaLicenseStateAllOf) HasLicenseActivated() bool {
	if o != nil && o.LicenseActivated != nil {
		return true
	}

	return false
}

// SetLicenseActivated gets a reference to the given string and assigns it to the LicenseActivated field.
func (o *NiatelemetryNiaLicenseStateAllOf) SetLicenseActivated(v string) {
	o.LicenseActivated = &v
}

// GetPidType returns the PidType field value if set, zero value otherwise.
func (o *NiatelemetryNiaLicenseStateAllOf) GetPidType() string {
	if o == nil || o.PidType == nil {
		var ret string
		return ret
	}
	return *o.PidType
}

// GetPidTypeOk returns a tuple with the PidType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaLicenseStateAllOf) GetPidTypeOk() (*string, bool) {
	if o == nil || o.PidType == nil {
		return nil, false
	}
	return o.PidType, true
}

// HasPidType returns a boolean if a field has been set.
func (o *NiatelemetryNiaLicenseStateAllOf) HasPidType() bool {
	if o != nil && o.PidType != nil {
		return true
	}

	return false
}

// SetPidType gets a reference to the given string and assigns it to the PidType field.
func (o *NiatelemetryNiaLicenseStateAllOf) SetPidType(v string) {
	o.PidType = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryNiaLicenseStateAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaLicenseStateAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryNiaLicenseStateAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryNiaLicenseStateAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryNiaLicenseStateAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaLicenseStateAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryNiaLicenseStateAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryNiaLicenseStateAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *NiatelemetryNiaLicenseStateAllOf) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaLicenseStateAllOf) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *NiatelemetryNiaLicenseStateAllOf) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *NiatelemetryNiaLicenseStateAllOf) SetSerial(v string) {
	o.Serial = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryNiaLicenseStateAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaLicenseStateAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryNiaLicenseStateAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryNiaLicenseStateAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *NiatelemetryNiaLicenseStateAllOf) GetDevice() NiatelemetryNiaInventoryRelationship {
	if o == nil || o.Device == nil {
		var ret NiatelemetryNiaInventoryRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaLicenseStateAllOf) GetDeviceOk() (*NiatelemetryNiaInventoryRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *NiatelemetryNiaLicenseStateAllOf) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NiatelemetryNiaInventoryRelationship and assigns it to the Device field.
func (o *NiatelemetryNiaLicenseStateAllOf) SetDevice(v NiatelemetryNiaInventoryRelationship) {
	o.Device = &v
}

func (o NiatelemetryNiaLicenseStateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FeatureActivated != nil {
		toSerialize["FeatureActivated"] = o.FeatureActivated
	}
	if o.LicenseActivated != nil {
		toSerialize["LicenseActivated"] = o.LicenseActivated
	}
	if o.PidType != nil {
		toSerialize["PidType"] = o.PidType
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.RecordVersion != nil {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNiaLicenseStateAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryNiaLicenseStateAllOf := _NiatelemetryNiaLicenseStateAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryNiaLicenseStateAllOf); err == nil {
		*o = NiatelemetryNiaLicenseStateAllOf(varNiatelemetryNiaLicenseStateAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FeatureActivated")
		delete(additionalProperties, "LicenseActivated")
		delete(additionalProperties, "PidType")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "Device")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryNiaLicenseStateAllOf struct {
	value *NiatelemetryNiaLicenseStateAllOf
	isSet bool
}

func (v NullableNiatelemetryNiaLicenseStateAllOf) Get() *NiatelemetryNiaLicenseStateAllOf {
	return v.value
}

func (v *NullableNiatelemetryNiaLicenseStateAllOf) Set(val *NiatelemetryNiaLicenseStateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNiaLicenseStateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNiaLicenseStateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNiaLicenseStateAllOf(val *NiatelemetryNiaLicenseStateAllOf) *NullableNiatelemetryNiaLicenseStateAllOf {
	return &NullableNiatelemetryNiaLicenseStateAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryNiaLicenseStateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNiaLicenseStateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
