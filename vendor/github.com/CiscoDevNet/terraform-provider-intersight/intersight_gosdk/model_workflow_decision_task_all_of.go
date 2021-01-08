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
)

// WorkflowDecisionTaskAllOf Definition of the list of properties defined in 'workflow.DecisionTask', excluding properties defined in parent classes.
type WorkflowDecisionTaskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The condition to evaluate for this decision task. The condition can be a workflow or task variable or an JavaScript expression. Example value for condition could be a variable like \"${task1.output.var1} or ${workflow.input.var2}\" which evaluates to a value matching any of the decision case values. Example value for condition if it's an expression - \"if ( ${task1.output.var1} ! = null && ${task1.output.var1} > 0 ) 'true'; else 'false'; \" which evaluates to 'true' or 'false' and will match one of the decision case values. You can also use JavaScript functions like indexOf, toUpperCase in the expression which will be evaluated by the expression evaluator.
	Condition     *string                `json:"Condition,omitempty"`
	DecisionCases []WorkflowDecisionCase `json:"DecisionCases,omitempty"`
	// The default next Task to execute if the decision cannot be evaluated to any of the DecisionCases.
	DefaultTask *string `json:"DefaultTask,omitempty"`
	// This field is deprecated. Decision case conditions can be added using the workflow input or task output variables in the Condition field. Refer to Condition field for more details.
	InputParameters      interface{} `json:"InputParameters,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowDecisionTaskAllOf WorkflowDecisionTaskAllOf

// NewWorkflowDecisionTaskAllOf instantiates a new WorkflowDecisionTaskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowDecisionTaskAllOf(classId string, objectType string) *WorkflowDecisionTaskAllOf {
	this := WorkflowDecisionTaskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowDecisionTaskAllOfWithDefaults instantiates a new WorkflowDecisionTaskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowDecisionTaskAllOfWithDefaults() *WorkflowDecisionTaskAllOf {
	this := WorkflowDecisionTaskAllOf{}
	var classId string = "workflow.DecisionTask"
	this.ClassId = classId
	var objectType string = "workflow.DecisionTask"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowDecisionTaskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowDecisionTaskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowDecisionTaskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowDecisionTaskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowDecisionTaskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowDecisionTaskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *WorkflowDecisionTaskAllOf) GetCondition() string {
	if o == nil || o.Condition == nil {
		var ret string
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDecisionTaskAllOf) GetConditionOk() (*string, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *WorkflowDecisionTaskAllOf) HasCondition() bool {
	if o != nil && o.Condition != nil {
		return true
	}

	return false
}

// SetCondition gets a reference to the given string and assigns it to the Condition field.
func (o *WorkflowDecisionTaskAllOf) SetCondition(v string) {
	o.Condition = &v
}

// GetDecisionCases returns the DecisionCases field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowDecisionTaskAllOf) GetDecisionCases() []WorkflowDecisionCase {
	if o == nil {
		var ret []WorkflowDecisionCase
		return ret
	}
	return o.DecisionCases
}

// GetDecisionCasesOk returns a tuple with the DecisionCases field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowDecisionTaskAllOf) GetDecisionCasesOk() (*[]WorkflowDecisionCase, bool) {
	if o == nil || o.DecisionCases == nil {
		return nil, false
	}
	return &o.DecisionCases, true
}

// HasDecisionCases returns a boolean if a field has been set.
func (o *WorkflowDecisionTaskAllOf) HasDecisionCases() bool {
	if o != nil && o.DecisionCases != nil {
		return true
	}

	return false
}

// SetDecisionCases gets a reference to the given []WorkflowDecisionCase and assigns it to the DecisionCases field.
func (o *WorkflowDecisionTaskAllOf) SetDecisionCases(v []WorkflowDecisionCase) {
	o.DecisionCases = v
}

// GetDefaultTask returns the DefaultTask field value if set, zero value otherwise.
func (o *WorkflowDecisionTaskAllOf) GetDefaultTask() string {
	if o == nil || o.DefaultTask == nil {
		var ret string
		return ret
	}
	return *o.DefaultTask
}

// GetDefaultTaskOk returns a tuple with the DefaultTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDecisionTaskAllOf) GetDefaultTaskOk() (*string, bool) {
	if o == nil || o.DefaultTask == nil {
		return nil, false
	}
	return o.DefaultTask, true
}

// HasDefaultTask returns a boolean if a field has been set.
func (o *WorkflowDecisionTaskAllOf) HasDefaultTask() bool {
	if o != nil && o.DefaultTask != nil {
		return true
	}

	return false
}

// SetDefaultTask gets a reference to the given string and assigns it to the DefaultTask field.
func (o *WorkflowDecisionTaskAllOf) SetDefaultTask(v string) {
	o.DefaultTask = &v
}

// GetInputParameters returns the InputParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowDecisionTaskAllOf) GetInputParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.InputParameters
}

// GetInputParametersOk returns a tuple with the InputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowDecisionTaskAllOf) GetInputParametersOk() (*interface{}, bool) {
	if o == nil || o.InputParameters == nil {
		return nil, false
	}
	return &o.InputParameters, true
}

// HasInputParameters returns a boolean if a field has been set.
func (o *WorkflowDecisionTaskAllOf) HasInputParameters() bool {
	if o != nil && o.InputParameters != nil {
		return true
	}

	return false
}

// SetInputParameters gets a reference to the given interface{} and assigns it to the InputParameters field.
func (o *WorkflowDecisionTaskAllOf) SetInputParameters(v interface{}) {
	o.InputParameters = v
}

func (o WorkflowDecisionTaskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Condition != nil {
		toSerialize["Condition"] = o.Condition
	}
	if o.DecisionCases != nil {
		toSerialize["DecisionCases"] = o.DecisionCases
	}
	if o.DefaultTask != nil {
		toSerialize["DefaultTask"] = o.DefaultTask
	}
	if o.InputParameters != nil {
		toSerialize["InputParameters"] = o.InputParameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowDecisionTaskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowDecisionTaskAllOf := _WorkflowDecisionTaskAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowDecisionTaskAllOf); err == nil {
		*o = WorkflowDecisionTaskAllOf(varWorkflowDecisionTaskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Condition")
		delete(additionalProperties, "DecisionCases")
		delete(additionalProperties, "DefaultTask")
		delete(additionalProperties, "InputParameters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowDecisionTaskAllOf struct {
	value *WorkflowDecisionTaskAllOf
	isSet bool
}

func (v NullableWorkflowDecisionTaskAllOf) Get() *WorkflowDecisionTaskAllOf {
	return v.value
}

func (v *NullableWorkflowDecisionTaskAllOf) Set(val *WorkflowDecisionTaskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowDecisionTaskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowDecisionTaskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowDecisionTaskAllOf(val *WorkflowDecisionTaskAllOf) *NullableWorkflowDecisionTaskAllOf {
	return &NullableWorkflowDecisionTaskAllOf{value: val, isSet: true}
}

func (v NullableWorkflowDecisionTaskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowDecisionTaskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
