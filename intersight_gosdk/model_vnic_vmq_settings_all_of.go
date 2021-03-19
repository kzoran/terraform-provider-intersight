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
)

// VnicVmqSettingsAllOf Definition of the list of properties defined in 'vnic.VmqSettings', excluding properties defined in parent classes.
type VnicVmqSettingsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enables VMQ feature on the virtual interface.
	Enabled *bool `json:"Enabled,omitempty"`
	// Enables Virtual Machine Multi-Queue feature on the virtual interface. VMMQ allows configuration of multiple I/O queues for a single VM and thus distributes traffic across multiple CPU cores in a VM.
	MultiQueueSupport *bool `json:"MultiQueueSupport,omitempty"`
	// The number of interrupt resources to be allocated. Recommended value is the number of CPU threads or logical processors available in the server.
	NumInterrupts *int64 `json:"NumInterrupts,omitempty"`
	// The number of sub vNICs to be created.
	NumSubVnics *int64 `json:"NumSubVnics,omitempty"`
	// The number of hardware Virtual Machine Queues to be allocated. The number of VMQs per adapter must be one more than the maximum number of VM NICs.
	NumVmqs *int64 `json:"NumVmqs,omitempty"`
	// Ethernet Adapter policy to be associated with the Sub vNICs. The Transmit Queue and Receive Queue resource value of VMMQ adapter policy should be greater than or equal to the configured number of sub vNICs.
	VmmqAdapterPolicy    *string `json:"VmmqAdapterPolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicVmqSettingsAllOf VnicVmqSettingsAllOf

// NewVnicVmqSettingsAllOf instantiates a new VnicVmqSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicVmqSettingsAllOf(classId string, objectType string) *VnicVmqSettingsAllOf {
	this := VnicVmqSettingsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = false
	this.Enabled = &enabled
	var multiQueueSupport bool = false
	this.MultiQueueSupport = &multiQueueSupport
	var numInterrupts int64 = 16
	this.NumInterrupts = &numInterrupts
	var numSubVnics int64 = 64
	this.NumSubVnics = &numSubVnics
	var numVmqs int64 = 4
	this.NumVmqs = &numVmqs
	return &this
}

// NewVnicVmqSettingsAllOfWithDefaults instantiates a new VnicVmqSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicVmqSettingsAllOfWithDefaults() *VnicVmqSettingsAllOf {
	this := VnicVmqSettingsAllOf{}
	var classId string = "vnic.VmqSettings"
	this.ClassId = classId
	var objectType string = "vnic.VmqSettings"
	this.ObjectType = objectType
	var enabled bool = false
	this.Enabled = &enabled
	var multiQueueSupport bool = false
	this.MultiQueueSupport = &multiQueueSupport
	var numInterrupts int64 = 16
	this.NumInterrupts = &numInterrupts
	var numSubVnics int64 = 64
	this.NumSubVnics = &numSubVnics
	var numVmqs int64 = 4
	this.NumVmqs = &numVmqs
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicVmqSettingsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicVmqSettingsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicVmqSettingsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicVmqSettingsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicVmqSettingsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicVmqSettingsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *VnicVmqSettingsAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVmqSettingsAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *VnicVmqSettingsAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *VnicVmqSettingsAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMultiQueueSupport returns the MultiQueueSupport field value if set, zero value otherwise.
func (o *VnicVmqSettingsAllOf) GetMultiQueueSupport() bool {
	if o == nil || o.MultiQueueSupport == nil {
		var ret bool
		return ret
	}
	return *o.MultiQueueSupport
}

// GetMultiQueueSupportOk returns a tuple with the MultiQueueSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVmqSettingsAllOf) GetMultiQueueSupportOk() (*bool, bool) {
	if o == nil || o.MultiQueueSupport == nil {
		return nil, false
	}
	return o.MultiQueueSupport, true
}

// HasMultiQueueSupport returns a boolean if a field has been set.
func (o *VnicVmqSettingsAllOf) HasMultiQueueSupport() bool {
	if o != nil && o.MultiQueueSupport != nil {
		return true
	}

	return false
}

// SetMultiQueueSupport gets a reference to the given bool and assigns it to the MultiQueueSupport field.
func (o *VnicVmqSettingsAllOf) SetMultiQueueSupport(v bool) {
	o.MultiQueueSupport = &v
}

// GetNumInterrupts returns the NumInterrupts field value if set, zero value otherwise.
func (o *VnicVmqSettingsAllOf) GetNumInterrupts() int64 {
	if o == nil || o.NumInterrupts == nil {
		var ret int64
		return ret
	}
	return *o.NumInterrupts
}

// GetNumInterruptsOk returns a tuple with the NumInterrupts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVmqSettingsAllOf) GetNumInterruptsOk() (*int64, bool) {
	if o == nil || o.NumInterrupts == nil {
		return nil, false
	}
	return o.NumInterrupts, true
}

// HasNumInterrupts returns a boolean if a field has been set.
func (o *VnicVmqSettingsAllOf) HasNumInterrupts() bool {
	if o != nil && o.NumInterrupts != nil {
		return true
	}

	return false
}

// SetNumInterrupts gets a reference to the given int64 and assigns it to the NumInterrupts field.
func (o *VnicVmqSettingsAllOf) SetNumInterrupts(v int64) {
	o.NumInterrupts = &v
}

// GetNumSubVnics returns the NumSubVnics field value if set, zero value otherwise.
func (o *VnicVmqSettingsAllOf) GetNumSubVnics() int64 {
	if o == nil || o.NumSubVnics == nil {
		var ret int64
		return ret
	}
	return *o.NumSubVnics
}

// GetNumSubVnicsOk returns a tuple with the NumSubVnics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVmqSettingsAllOf) GetNumSubVnicsOk() (*int64, bool) {
	if o == nil || o.NumSubVnics == nil {
		return nil, false
	}
	return o.NumSubVnics, true
}

// HasNumSubVnics returns a boolean if a field has been set.
func (o *VnicVmqSettingsAllOf) HasNumSubVnics() bool {
	if o != nil && o.NumSubVnics != nil {
		return true
	}

	return false
}

// SetNumSubVnics gets a reference to the given int64 and assigns it to the NumSubVnics field.
func (o *VnicVmqSettingsAllOf) SetNumSubVnics(v int64) {
	o.NumSubVnics = &v
}

// GetNumVmqs returns the NumVmqs field value if set, zero value otherwise.
func (o *VnicVmqSettingsAllOf) GetNumVmqs() int64 {
	if o == nil || o.NumVmqs == nil {
		var ret int64
		return ret
	}
	return *o.NumVmqs
}

// GetNumVmqsOk returns a tuple with the NumVmqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVmqSettingsAllOf) GetNumVmqsOk() (*int64, bool) {
	if o == nil || o.NumVmqs == nil {
		return nil, false
	}
	return o.NumVmqs, true
}

// HasNumVmqs returns a boolean if a field has been set.
func (o *VnicVmqSettingsAllOf) HasNumVmqs() bool {
	if o != nil && o.NumVmqs != nil {
		return true
	}

	return false
}

// SetNumVmqs gets a reference to the given int64 and assigns it to the NumVmqs field.
func (o *VnicVmqSettingsAllOf) SetNumVmqs(v int64) {
	o.NumVmqs = &v
}

// GetVmmqAdapterPolicy returns the VmmqAdapterPolicy field value if set, zero value otherwise.
func (o *VnicVmqSettingsAllOf) GetVmmqAdapterPolicy() string {
	if o == nil || o.VmmqAdapterPolicy == nil {
		var ret string
		return ret
	}
	return *o.VmmqAdapterPolicy
}

// GetVmmqAdapterPolicyOk returns a tuple with the VmmqAdapterPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVmqSettingsAllOf) GetVmmqAdapterPolicyOk() (*string, bool) {
	if o == nil || o.VmmqAdapterPolicy == nil {
		return nil, false
	}
	return o.VmmqAdapterPolicy, true
}

// HasVmmqAdapterPolicy returns a boolean if a field has been set.
func (o *VnicVmqSettingsAllOf) HasVmmqAdapterPolicy() bool {
	if o != nil && o.VmmqAdapterPolicy != nil {
		return true
	}

	return false
}

// SetVmmqAdapterPolicy gets a reference to the given string and assigns it to the VmmqAdapterPolicy field.
func (o *VnicVmqSettingsAllOf) SetVmmqAdapterPolicy(v string) {
	o.VmmqAdapterPolicy = &v
}

func (o VnicVmqSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.MultiQueueSupport != nil {
		toSerialize["MultiQueueSupport"] = o.MultiQueueSupport
	}
	if o.NumInterrupts != nil {
		toSerialize["NumInterrupts"] = o.NumInterrupts
	}
	if o.NumSubVnics != nil {
		toSerialize["NumSubVnics"] = o.NumSubVnics
	}
	if o.NumVmqs != nil {
		toSerialize["NumVmqs"] = o.NumVmqs
	}
	if o.VmmqAdapterPolicy != nil {
		toSerialize["VmmqAdapterPolicy"] = o.VmmqAdapterPolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicVmqSettingsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicVmqSettingsAllOf := _VnicVmqSettingsAllOf{}

	if err = json.Unmarshal(bytes, &varVnicVmqSettingsAllOf); err == nil {
		*o = VnicVmqSettingsAllOf(varVnicVmqSettingsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "MultiQueueSupport")
		delete(additionalProperties, "NumInterrupts")
		delete(additionalProperties, "NumSubVnics")
		delete(additionalProperties, "NumVmqs")
		delete(additionalProperties, "VmmqAdapterPolicy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicVmqSettingsAllOf struct {
	value *VnicVmqSettingsAllOf
	isSet bool
}

func (v NullableVnicVmqSettingsAllOf) Get() *VnicVmqSettingsAllOf {
	return v.value
}

func (v *NullableVnicVmqSettingsAllOf) Set(val *VnicVmqSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicVmqSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicVmqSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicVmqSettingsAllOf(val *VnicVmqSettingsAllOf) *NullableVnicVmqSettingsAllOf {
	return &NullableVnicVmqSettingsAllOf{value: val, isSet: true}
}

func (v NullableVnicVmqSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicVmqSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
