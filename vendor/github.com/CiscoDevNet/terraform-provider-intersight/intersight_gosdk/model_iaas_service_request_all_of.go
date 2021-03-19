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

// IaasServiceRequestAllOf Definition of the list of properties defined in 'iaas.ServiceRequest', excluding properties defined in parent classes.
type IaasServiceRequestAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Service request duration.
	Duration *string `json:"Duration,omitempty"`
	// Service Request initiating user.
	InitiatingUser *string `json:"InitiatingUser,omitempty"`
	// Service request end time.
	RequestEndTime *string `json:"RequestEndTime,omitempty"`
	// Service request id of an SR.
	RequestId *string `json:"RequestId,omitempty"`
	// Service request start time.
	RequestStartTime *string `json:"RequestStartTime,omitempty"`
	// Service request type of an SR.
	RequestType *string `json:"RequestType,omitempty"`
	// UCSD service request status.
	Status *string `json:"Status,omitempty"`
	// Executed workflow name for an SR.
	WorkflowName         *string                              `json:"WorkflowName,omitempty"`
	WorkflowSteps        []IaasWorkflowSteps                  `json:"WorkflowSteps,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IaasServiceRequestAllOf IaasServiceRequestAllOf

// NewIaasServiceRequestAllOf instantiates a new IaasServiceRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasServiceRequestAllOf(classId string, objectType string) *IaasServiceRequestAllOf {
	this := IaasServiceRequestAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIaasServiceRequestAllOfWithDefaults instantiates a new IaasServiceRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasServiceRequestAllOfWithDefaults() *IaasServiceRequestAllOf {
	this := IaasServiceRequestAllOf{}
	var classId string = "iaas.ServiceRequest"
	this.ClassId = classId
	var objectType string = "iaas.ServiceRequest"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IaasServiceRequestAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IaasServiceRequestAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IaasServiceRequestAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IaasServiceRequestAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IaasServiceRequestAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IaasServiceRequestAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *IaasServiceRequestAllOf) GetDuration() string {
	if o == nil || o.Duration == nil {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequestAllOf) GetDurationOk() (*string, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *IaasServiceRequestAllOf) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *IaasServiceRequestAllOf) SetDuration(v string) {
	o.Duration = &v
}

// GetInitiatingUser returns the InitiatingUser field value if set, zero value otherwise.
func (o *IaasServiceRequestAllOf) GetInitiatingUser() string {
	if o == nil || o.InitiatingUser == nil {
		var ret string
		return ret
	}
	return *o.InitiatingUser
}

// GetInitiatingUserOk returns a tuple with the InitiatingUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequestAllOf) GetInitiatingUserOk() (*string, bool) {
	if o == nil || o.InitiatingUser == nil {
		return nil, false
	}
	return o.InitiatingUser, true
}

// HasInitiatingUser returns a boolean if a field has been set.
func (o *IaasServiceRequestAllOf) HasInitiatingUser() bool {
	if o != nil && o.InitiatingUser != nil {
		return true
	}

	return false
}

// SetInitiatingUser gets a reference to the given string and assigns it to the InitiatingUser field.
func (o *IaasServiceRequestAllOf) SetInitiatingUser(v string) {
	o.InitiatingUser = &v
}

// GetRequestEndTime returns the RequestEndTime field value if set, zero value otherwise.
func (o *IaasServiceRequestAllOf) GetRequestEndTime() string {
	if o == nil || o.RequestEndTime == nil {
		var ret string
		return ret
	}
	return *o.RequestEndTime
}

// GetRequestEndTimeOk returns a tuple with the RequestEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequestAllOf) GetRequestEndTimeOk() (*string, bool) {
	if o == nil || o.RequestEndTime == nil {
		return nil, false
	}
	return o.RequestEndTime, true
}

// HasRequestEndTime returns a boolean if a field has been set.
func (o *IaasServiceRequestAllOf) HasRequestEndTime() bool {
	if o != nil && o.RequestEndTime != nil {
		return true
	}

	return false
}

// SetRequestEndTime gets a reference to the given string and assigns it to the RequestEndTime field.
func (o *IaasServiceRequestAllOf) SetRequestEndTime(v string) {
	o.RequestEndTime = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *IaasServiceRequestAllOf) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequestAllOf) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *IaasServiceRequestAllOf) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *IaasServiceRequestAllOf) SetRequestId(v string) {
	o.RequestId = &v
}

// GetRequestStartTime returns the RequestStartTime field value if set, zero value otherwise.
func (o *IaasServiceRequestAllOf) GetRequestStartTime() string {
	if o == nil || o.RequestStartTime == nil {
		var ret string
		return ret
	}
	return *o.RequestStartTime
}

// GetRequestStartTimeOk returns a tuple with the RequestStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequestAllOf) GetRequestStartTimeOk() (*string, bool) {
	if o == nil || o.RequestStartTime == nil {
		return nil, false
	}
	return o.RequestStartTime, true
}

// HasRequestStartTime returns a boolean if a field has been set.
func (o *IaasServiceRequestAllOf) HasRequestStartTime() bool {
	if o != nil && o.RequestStartTime != nil {
		return true
	}

	return false
}

// SetRequestStartTime gets a reference to the given string and assigns it to the RequestStartTime field.
func (o *IaasServiceRequestAllOf) SetRequestStartTime(v string) {
	o.RequestStartTime = &v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *IaasServiceRequestAllOf) GetRequestType() string {
	if o == nil || o.RequestType == nil {
		var ret string
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequestAllOf) GetRequestTypeOk() (*string, bool) {
	if o == nil || o.RequestType == nil {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *IaasServiceRequestAllOf) HasRequestType() bool {
	if o != nil && o.RequestType != nil {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given string and assigns it to the RequestType field.
func (o *IaasServiceRequestAllOf) SetRequestType(v string) {
	o.RequestType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IaasServiceRequestAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequestAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IaasServiceRequestAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IaasServiceRequestAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetWorkflowName returns the WorkflowName field value if set, zero value otherwise.
func (o *IaasServiceRequestAllOf) GetWorkflowName() string {
	if o == nil || o.WorkflowName == nil {
		var ret string
		return ret
	}
	return *o.WorkflowName
}

// GetWorkflowNameOk returns a tuple with the WorkflowName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequestAllOf) GetWorkflowNameOk() (*string, bool) {
	if o == nil || o.WorkflowName == nil {
		return nil, false
	}
	return o.WorkflowName, true
}

// HasWorkflowName returns a boolean if a field has been set.
func (o *IaasServiceRequestAllOf) HasWorkflowName() bool {
	if o != nil && o.WorkflowName != nil {
		return true
	}

	return false
}

// SetWorkflowName gets a reference to the given string and assigns it to the WorkflowName field.
func (o *IaasServiceRequestAllOf) SetWorkflowName(v string) {
	o.WorkflowName = &v
}

// GetWorkflowSteps returns the WorkflowSteps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasServiceRequestAllOf) GetWorkflowSteps() []IaasWorkflowSteps {
	if o == nil {
		var ret []IaasWorkflowSteps
		return ret
	}
	return o.WorkflowSteps
}

// GetWorkflowStepsOk returns a tuple with the WorkflowSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasServiceRequestAllOf) GetWorkflowStepsOk() (*[]IaasWorkflowSteps, bool) {
	if o == nil || o.WorkflowSteps == nil {
		return nil, false
	}
	return &o.WorkflowSteps, true
}

// HasWorkflowSteps returns a boolean if a field has been set.
func (o *IaasServiceRequestAllOf) HasWorkflowSteps() bool {
	if o != nil && o.WorkflowSteps != nil {
		return true
	}

	return false
}

// SetWorkflowSteps gets a reference to the given []IaasWorkflowSteps and assigns it to the WorkflowSteps field.
func (o *IaasServiceRequestAllOf) SetWorkflowSteps(v []IaasWorkflowSteps) {
	o.WorkflowSteps = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *IaasServiceRequestAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequestAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *IaasServiceRequestAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *IaasServiceRequestAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o IaasServiceRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Duration != nil {
		toSerialize["Duration"] = o.Duration
	}
	if o.InitiatingUser != nil {
		toSerialize["InitiatingUser"] = o.InitiatingUser
	}
	if o.RequestEndTime != nil {
		toSerialize["RequestEndTime"] = o.RequestEndTime
	}
	if o.RequestId != nil {
		toSerialize["RequestId"] = o.RequestId
	}
	if o.RequestStartTime != nil {
		toSerialize["RequestStartTime"] = o.RequestStartTime
	}
	if o.RequestType != nil {
		toSerialize["RequestType"] = o.RequestType
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.WorkflowName != nil {
		toSerialize["WorkflowName"] = o.WorkflowName
	}
	if o.WorkflowSteps != nil {
		toSerialize["WorkflowSteps"] = o.WorkflowSteps
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IaasServiceRequestAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIaasServiceRequestAllOf := _IaasServiceRequestAllOf{}

	if err = json.Unmarshal(bytes, &varIaasServiceRequestAllOf); err == nil {
		*o = IaasServiceRequestAllOf(varIaasServiceRequestAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Duration")
		delete(additionalProperties, "InitiatingUser")
		delete(additionalProperties, "RequestEndTime")
		delete(additionalProperties, "RequestId")
		delete(additionalProperties, "RequestStartTime")
		delete(additionalProperties, "RequestType")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "WorkflowName")
		delete(additionalProperties, "WorkflowSteps")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIaasServiceRequestAllOf struct {
	value *IaasServiceRequestAllOf
	isSet bool
}

func (v NullableIaasServiceRequestAllOf) Get() *IaasServiceRequestAllOf {
	return v.value
}

func (v *NullableIaasServiceRequestAllOf) Set(val *IaasServiceRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasServiceRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasServiceRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasServiceRequestAllOf(val *IaasServiceRequestAllOf) *NullableIaasServiceRequestAllOf {
	return &NullableIaasServiceRequestAllOf{value: val, isSet: true}
}

func (v NullableIaasServiceRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasServiceRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
