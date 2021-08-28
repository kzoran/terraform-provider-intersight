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
	"time"
)

// CondAlarmAllOf Definition of the list of properties defined in 'cond.Alarm', excluding properties defined in parent classes.
type CondAlarmAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Alarm acknowledgment state. Default value is None. * `None` - The Enum value None represents that the alarm is not acknowledged and is included as part of health status and overall alarm count. * `Acknowledge` - The Enum value Acknowledge represents that the alarm is acknowledged by user. The alarm will be ignored from the health status and overall alarm count.
	Acknowledge *string `json:"Acknowledge,omitempty"`
	// User who acknowledged the alarm.
	AcknowledgeBy *string `json:"AcknowledgeBy,omitempty"`
	// Time at which the alarm was acknowledged by the user.
	AcknowledgeTime *time.Time `json:"AcknowledgeTime,omitempty"`
	// Display name of the affected object on which the alarm is raised. The name uniquely identifies an alarm target such that it can be distinguished from similar objects managed by Intersight.
	AffectedMoDisplayName *string `json:"AffectedMoDisplayName,omitempty"`
	// MoId of the affected object from the managed system's point of view.
	// Deprecated
	AffectedMoId *string `json:"AffectedMoId,omitempty"`
	// Managed system affected object type. For example Adaptor, FI, CIMC.
	// Deprecated
	AffectedMoType *string `json:"AffectedMoType,omitempty"`
	// A unique key for an alarm instance, consists of a combination of a unique system name and msAffectedObject.
	// Deprecated
	AffectedObject *string `json:"AffectedObject,omitempty"`
	// Parent MoId of the fault from managed system. For example, Blade moid for adaptor fault.
	AncestorMoId *string `json:"AncestorMoId,omitempty"`
	// Parent MO type of the fault from managed system. For example, Blade for adaptor fault.
	AncestorMoType *string `json:"AncestorMoType,omitempty"`
	// A unique alarm code. For alarms mapped from UCS faults, this will be the same as the UCS fault code.
	Code *string `json:"Code,omitempty"`
	// The time the alarm was created.
	CreationTime *time.Time `json:"CreationTime,omitempty"`
	// A longer description of the alarm than the name. The description contains details of the component reporting the issue.
	Description *string `json:"Description,omitempty"`
	// The time the alarm last had a change in severity.
	LastTransitionTime *time.Time `json:"LastTransitionTime,omitempty"`
	// A unique key for the alarm from the managed system's point of view. For example, in the case of UCS, this is the fault's dn.
	// Deprecated
	MsAffectedObject *string `json:"MsAffectedObject,omitempty"`
	// Uniquely identifies the type of alarm. For alarms originating from Intersight, this will be a descriptive name. For alarms that are mapped from faults, the name will be derived from fault properties. For example, alarms mapped from UCS faults will use a prefix of UCS and appended with the fault code.
	Name *string `json:"Name,omitempty"`
	// The original severity when the alarm was first created. * `None` - The Enum value None represents that there is no severity. * `Info` - The Enum value Info represents the Informational level of severity. * `Critical` - The Enum value Critical represents the Critical level of severity. * `Warning` - The Enum value Warning represents the Warning level of severity. * `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared.
	OrigSeverity *string `json:"OrigSeverity,omitempty"`
	// The severity of the alarm. Valid values are Critical, Warning, Info, and Cleared. * `None` - The Enum value None represents that there is no severity. * `Info` - The Enum value Info represents the Informational level of severity. * `Critical` - The Enum value Critical represents the Critical level of severity. * `Warning` - The Enum value Warning represents the Warning level of severity. * `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared.
	Severity             *string                              `json:"Severity,omitempty"`
	AffectedMo           *MoBaseMoRelationship                `json:"AffectedMo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CondAlarmAllOf CondAlarmAllOf

// NewCondAlarmAllOf instantiates a new CondAlarmAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCondAlarmAllOf(classId string, objectType string) *CondAlarmAllOf {
	this := CondAlarmAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var acknowledge string = "None"
	this.Acknowledge = &acknowledge
	return &this
}

// NewCondAlarmAllOfWithDefaults instantiates a new CondAlarmAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCondAlarmAllOfWithDefaults() *CondAlarmAllOf {
	this := CondAlarmAllOf{}
	var classId string = "cond.Alarm"
	this.ClassId = classId
	var objectType string = "cond.Alarm"
	this.ObjectType = objectType
	var acknowledge string = "None"
	this.Acknowledge = &acknowledge
	return &this
}

// GetClassId returns the ClassId field value
func (o *CondAlarmAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CondAlarmAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CondAlarmAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CondAlarmAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAcknowledge returns the Acknowledge field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetAcknowledge() string {
	if o == nil || o.Acknowledge == nil {
		var ret string
		return ret
	}
	return *o.Acknowledge
}

// GetAcknowledgeOk returns a tuple with the Acknowledge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetAcknowledgeOk() (*string, bool) {
	if o == nil || o.Acknowledge == nil {
		return nil, false
	}
	return o.Acknowledge, true
}

// HasAcknowledge returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasAcknowledge() bool {
	if o != nil && o.Acknowledge != nil {
		return true
	}

	return false
}

// SetAcknowledge gets a reference to the given string and assigns it to the Acknowledge field.
func (o *CondAlarmAllOf) SetAcknowledge(v string) {
	o.Acknowledge = &v
}

// GetAcknowledgeBy returns the AcknowledgeBy field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetAcknowledgeBy() string {
	if o == nil || o.AcknowledgeBy == nil {
		var ret string
		return ret
	}
	return *o.AcknowledgeBy
}

// GetAcknowledgeByOk returns a tuple with the AcknowledgeBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetAcknowledgeByOk() (*string, bool) {
	if o == nil || o.AcknowledgeBy == nil {
		return nil, false
	}
	return o.AcknowledgeBy, true
}

// HasAcknowledgeBy returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasAcknowledgeBy() bool {
	if o != nil && o.AcknowledgeBy != nil {
		return true
	}

	return false
}

// SetAcknowledgeBy gets a reference to the given string and assigns it to the AcknowledgeBy field.
func (o *CondAlarmAllOf) SetAcknowledgeBy(v string) {
	o.AcknowledgeBy = &v
}

// GetAcknowledgeTime returns the AcknowledgeTime field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetAcknowledgeTime() time.Time {
	if o == nil || o.AcknowledgeTime == nil {
		var ret time.Time
		return ret
	}
	return *o.AcknowledgeTime
}

// GetAcknowledgeTimeOk returns a tuple with the AcknowledgeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetAcknowledgeTimeOk() (*time.Time, bool) {
	if o == nil || o.AcknowledgeTime == nil {
		return nil, false
	}
	return o.AcknowledgeTime, true
}

// HasAcknowledgeTime returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasAcknowledgeTime() bool {
	if o != nil && o.AcknowledgeTime != nil {
		return true
	}

	return false
}

// SetAcknowledgeTime gets a reference to the given time.Time and assigns it to the AcknowledgeTime field.
func (o *CondAlarmAllOf) SetAcknowledgeTime(v time.Time) {
	o.AcknowledgeTime = &v
}

// GetAffectedMoDisplayName returns the AffectedMoDisplayName field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetAffectedMoDisplayName() string {
	if o == nil || o.AffectedMoDisplayName == nil {
		var ret string
		return ret
	}
	return *o.AffectedMoDisplayName
}

// GetAffectedMoDisplayNameOk returns a tuple with the AffectedMoDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetAffectedMoDisplayNameOk() (*string, bool) {
	if o == nil || o.AffectedMoDisplayName == nil {
		return nil, false
	}
	return o.AffectedMoDisplayName, true
}

// HasAffectedMoDisplayName returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasAffectedMoDisplayName() bool {
	if o != nil && o.AffectedMoDisplayName != nil {
		return true
	}

	return false
}

// SetAffectedMoDisplayName gets a reference to the given string and assigns it to the AffectedMoDisplayName field.
func (o *CondAlarmAllOf) SetAffectedMoDisplayName(v string) {
	o.AffectedMoDisplayName = &v
}

// GetAffectedMoId returns the AffectedMoId field value if set, zero value otherwise.
// Deprecated
func (o *CondAlarmAllOf) GetAffectedMoId() string {
	if o == nil || o.AffectedMoId == nil {
		var ret string
		return ret
	}
	return *o.AffectedMoId
}

// GetAffectedMoIdOk returns a tuple with the AffectedMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CondAlarmAllOf) GetAffectedMoIdOk() (*string, bool) {
	if o == nil || o.AffectedMoId == nil {
		return nil, false
	}
	return o.AffectedMoId, true
}

// HasAffectedMoId returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasAffectedMoId() bool {
	if o != nil && o.AffectedMoId != nil {
		return true
	}

	return false
}

// SetAffectedMoId gets a reference to the given string and assigns it to the AffectedMoId field.
// Deprecated
func (o *CondAlarmAllOf) SetAffectedMoId(v string) {
	o.AffectedMoId = &v
}

// GetAffectedMoType returns the AffectedMoType field value if set, zero value otherwise.
// Deprecated
func (o *CondAlarmAllOf) GetAffectedMoType() string {
	if o == nil || o.AffectedMoType == nil {
		var ret string
		return ret
	}
	return *o.AffectedMoType
}

// GetAffectedMoTypeOk returns a tuple with the AffectedMoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CondAlarmAllOf) GetAffectedMoTypeOk() (*string, bool) {
	if o == nil || o.AffectedMoType == nil {
		return nil, false
	}
	return o.AffectedMoType, true
}

// HasAffectedMoType returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasAffectedMoType() bool {
	if o != nil && o.AffectedMoType != nil {
		return true
	}

	return false
}

// SetAffectedMoType gets a reference to the given string and assigns it to the AffectedMoType field.
// Deprecated
func (o *CondAlarmAllOf) SetAffectedMoType(v string) {
	o.AffectedMoType = &v
}

// GetAffectedObject returns the AffectedObject field value if set, zero value otherwise.
// Deprecated
func (o *CondAlarmAllOf) GetAffectedObject() string {
	if o == nil || o.AffectedObject == nil {
		var ret string
		return ret
	}
	return *o.AffectedObject
}

// GetAffectedObjectOk returns a tuple with the AffectedObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CondAlarmAllOf) GetAffectedObjectOk() (*string, bool) {
	if o == nil || o.AffectedObject == nil {
		return nil, false
	}
	return o.AffectedObject, true
}

// HasAffectedObject returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasAffectedObject() bool {
	if o != nil && o.AffectedObject != nil {
		return true
	}

	return false
}

// SetAffectedObject gets a reference to the given string and assigns it to the AffectedObject field.
// Deprecated
func (o *CondAlarmAllOf) SetAffectedObject(v string) {
	o.AffectedObject = &v
}

// GetAncestorMoId returns the AncestorMoId field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetAncestorMoId() string {
	if o == nil || o.AncestorMoId == nil {
		var ret string
		return ret
	}
	return *o.AncestorMoId
}

// GetAncestorMoIdOk returns a tuple with the AncestorMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetAncestorMoIdOk() (*string, bool) {
	if o == nil || o.AncestorMoId == nil {
		return nil, false
	}
	return o.AncestorMoId, true
}

// HasAncestorMoId returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasAncestorMoId() bool {
	if o != nil && o.AncestorMoId != nil {
		return true
	}

	return false
}

// SetAncestorMoId gets a reference to the given string and assigns it to the AncestorMoId field.
func (o *CondAlarmAllOf) SetAncestorMoId(v string) {
	o.AncestorMoId = &v
}

// GetAncestorMoType returns the AncestorMoType field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetAncestorMoType() string {
	if o == nil || o.AncestorMoType == nil {
		var ret string
		return ret
	}
	return *o.AncestorMoType
}

// GetAncestorMoTypeOk returns a tuple with the AncestorMoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetAncestorMoTypeOk() (*string, bool) {
	if o == nil || o.AncestorMoType == nil {
		return nil, false
	}
	return o.AncestorMoType, true
}

// HasAncestorMoType returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasAncestorMoType() bool {
	if o != nil && o.AncestorMoType != nil {
		return true
	}

	return false
}

// SetAncestorMoType gets a reference to the given string and assigns it to the AncestorMoType field.
func (o *CondAlarmAllOf) SetAncestorMoType(v string) {
	o.AncestorMoType = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CondAlarmAllOf) SetCode(v string) {
	o.Code = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *CondAlarmAllOf) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CondAlarmAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetLastTransitionTime returns the LastTransitionTime field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetLastTransitionTime() time.Time {
	if o == nil || o.LastTransitionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastTransitionTime
}

// GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetLastTransitionTimeOk() (*time.Time, bool) {
	if o == nil || o.LastTransitionTime == nil {
		return nil, false
	}
	return o.LastTransitionTime, true
}

// HasLastTransitionTime returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasLastTransitionTime() bool {
	if o != nil && o.LastTransitionTime != nil {
		return true
	}

	return false
}

// SetLastTransitionTime gets a reference to the given time.Time and assigns it to the LastTransitionTime field.
func (o *CondAlarmAllOf) SetLastTransitionTime(v time.Time) {
	o.LastTransitionTime = &v
}

// GetMsAffectedObject returns the MsAffectedObject field value if set, zero value otherwise.
// Deprecated
func (o *CondAlarmAllOf) GetMsAffectedObject() string {
	if o == nil || o.MsAffectedObject == nil {
		var ret string
		return ret
	}
	return *o.MsAffectedObject
}

// GetMsAffectedObjectOk returns a tuple with the MsAffectedObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CondAlarmAllOf) GetMsAffectedObjectOk() (*string, bool) {
	if o == nil || o.MsAffectedObject == nil {
		return nil, false
	}
	return o.MsAffectedObject, true
}

// HasMsAffectedObject returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasMsAffectedObject() bool {
	if o != nil && o.MsAffectedObject != nil {
		return true
	}

	return false
}

// SetMsAffectedObject gets a reference to the given string and assigns it to the MsAffectedObject field.
// Deprecated
func (o *CondAlarmAllOf) SetMsAffectedObject(v string) {
	o.MsAffectedObject = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CondAlarmAllOf) SetName(v string) {
	o.Name = &v
}

// GetOrigSeverity returns the OrigSeverity field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetOrigSeverity() string {
	if o == nil || o.OrigSeverity == nil {
		var ret string
		return ret
	}
	return *o.OrigSeverity
}

// GetOrigSeverityOk returns a tuple with the OrigSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetOrigSeverityOk() (*string, bool) {
	if o == nil || o.OrigSeverity == nil {
		return nil, false
	}
	return o.OrigSeverity, true
}

// HasOrigSeverity returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasOrigSeverity() bool {
	if o != nil && o.OrigSeverity != nil {
		return true
	}

	return false
}

// SetOrigSeverity gets a reference to the given string and assigns it to the OrigSeverity field.
func (o *CondAlarmAllOf) SetOrigSeverity(v string) {
	o.OrigSeverity = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *CondAlarmAllOf) SetSeverity(v string) {
	o.Severity = &v
}

// GetAffectedMo returns the AffectedMo field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetAffectedMo() MoBaseMoRelationship {
	if o == nil || o.AffectedMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AffectedMo
}

// GetAffectedMoOk returns a tuple with the AffectedMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetAffectedMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AffectedMo == nil {
		return nil, false
	}
	return o.AffectedMo, true
}

// HasAffectedMo returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasAffectedMo() bool {
	if o != nil && o.AffectedMo != nil {
		return true
	}

	return false
}

// SetAffectedMo gets a reference to the given MoBaseMoRelationship and assigns it to the AffectedMo field.
func (o *CondAlarmAllOf) SetAffectedMo(v MoBaseMoRelationship) {
	o.AffectedMo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *CondAlarmAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *CondAlarmAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *CondAlarmAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o CondAlarmAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Acknowledge != nil {
		toSerialize["Acknowledge"] = o.Acknowledge
	}
	if o.AcknowledgeBy != nil {
		toSerialize["AcknowledgeBy"] = o.AcknowledgeBy
	}
	if o.AcknowledgeTime != nil {
		toSerialize["AcknowledgeTime"] = o.AcknowledgeTime
	}
	if o.AffectedMoDisplayName != nil {
		toSerialize["AffectedMoDisplayName"] = o.AffectedMoDisplayName
	}
	if o.AffectedMoId != nil {
		toSerialize["AffectedMoId"] = o.AffectedMoId
	}
	if o.AffectedMoType != nil {
		toSerialize["AffectedMoType"] = o.AffectedMoType
	}
	if o.AffectedObject != nil {
		toSerialize["AffectedObject"] = o.AffectedObject
	}
	if o.AncestorMoId != nil {
		toSerialize["AncestorMoId"] = o.AncestorMoId
	}
	if o.AncestorMoType != nil {
		toSerialize["AncestorMoType"] = o.AncestorMoType
	}
	if o.Code != nil {
		toSerialize["Code"] = o.Code
	}
	if o.CreationTime != nil {
		toSerialize["CreationTime"] = o.CreationTime
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.LastTransitionTime != nil {
		toSerialize["LastTransitionTime"] = o.LastTransitionTime
	}
	if o.MsAffectedObject != nil {
		toSerialize["MsAffectedObject"] = o.MsAffectedObject
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OrigSeverity != nil {
		toSerialize["OrigSeverity"] = o.OrigSeverity
	}
	if o.Severity != nil {
		toSerialize["Severity"] = o.Severity
	}
	if o.AffectedMo != nil {
		toSerialize["AffectedMo"] = o.AffectedMo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CondAlarmAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCondAlarmAllOf := _CondAlarmAllOf{}

	if err = json.Unmarshal(bytes, &varCondAlarmAllOf); err == nil {
		*o = CondAlarmAllOf(varCondAlarmAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Acknowledge")
		delete(additionalProperties, "AcknowledgeBy")
		delete(additionalProperties, "AcknowledgeTime")
		delete(additionalProperties, "AffectedMoDisplayName")
		delete(additionalProperties, "AffectedMoId")
		delete(additionalProperties, "AffectedMoType")
		delete(additionalProperties, "AffectedObject")
		delete(additionalProperties, "AncestorMoId")
		delete(additionalProperties, "AncestorMoType")
		delete(additionalProperties, "Code")
		delete(additionalProperties, "CreationTime")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "LastTransitionTime")
		delete(additionalProperties, "MsAffectedObject")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OrigSeverity")
		delete(additionalProperties, "Severity")
		delete(additionalProperties, "AffectedMo")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCondAlarmAllOf struct {
	value *CondAlarmAllOf
	isSet bool
}

func (v NullableCondAlarmAllOf) Get() *CondAlarmAllOf {
	return v.value
}

func (v *NullableCondAlarmAllOf) Set(val *CondAlarmAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCondAlarmAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCondAlarmAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCondAlarmAllOf(val *CondAlarmAllOf) *NullableCondAlarmAllOf {
	return &NullableCondAlarmAllOf{value: val, isSet: true}
}

func (v NullableCondAlarmAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCondAlarmAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
