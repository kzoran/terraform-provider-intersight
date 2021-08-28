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

// WorkflowCommentsAllOf Definition of the list of properties defined in 'workflow.Comments', excluding properties defined in parent classes.
type WorkflowCommentsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description field provides comment about the template function.
	Description          *string  `json:"Description,omitempty"`
	Examples             []string `json:"Examples,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowCommentsAllOf WorkflowCommentsAllOf

// NewWorkflowCommentsAllOf instantiates a new WorkflowCommentsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowCommentsAllOf(classId string, objectType string) *WorkflowCommentsAllOf {
	this := WorkflowCommentsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowCommentsAllOfWithDefaults instantiates a new WorkflowCommentsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowCommentsAllOfWithDefaults() *WorkflowCommentsAllOf {
	this := WorkflowCommentsAllOf{}
	var classId string = "workflow.Comments"
	this.ClassId = classId
	var objectType string = "workflow.Comments"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowCommentsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowCommentsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowCommentsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowCommentsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowCommentsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowCommentsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowCommentsAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCommentsAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowCommentsAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowCommentsAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetExamples returns the Examples field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowCommentsAllOf) GetExamples() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Examples
}

// GetExamplesOk returns a tuple with the Examples field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowCommentsAllOf) GetExamplesOk() (*[]string, bool) {
	if o == nil || o.Examples == nil {
		return nil, false
	}
	return &o.Examples, true
}

// HasExamples returns a boolean if a field has been set.
func (o *WorkflowCommentsAllOf) HasExamples() bool {
	if o != nil && o.Examples != nil {
		return true
	}

	return false
}

// SetExamples gets a reference to the given []string and assigns it to the Examples field.
func (o *WorkflowCommentsAllOf) SetExamples(v []string) {
	o.Examples = v
}

func (o WorkflowCommentsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Examples != nil {
		toSerialize["Examples"] = o.Examples
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowCommentsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowCommentsAllOf := _WorkflowCommentsAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowCommentsAllOf); err == nil {
		*o = WorkflowCommentsAllOf(varWorkflowCommentsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Examples")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowCommentsAllOf struct {
	value *WorkflowCommentsAllOf
	isSet bool
}

func (v NullableWorkflowCommentsAllOf) Get() *WorkflowCommentsAllOf {
	return v.value
}

func (v *NullableWorkflowCommentsAllOf) Set(val *WorkflowCommentsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowCommentsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowCommentsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowCommentsAllOf(val *WorkflowCommentsAllOf) *NullableWorkflowCommentsAllOf {
	return &NullableWorkflowCommentsAllOf{value: val, isSet: true}
}

func (v NullableWorkflowCommentsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowCommentsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
