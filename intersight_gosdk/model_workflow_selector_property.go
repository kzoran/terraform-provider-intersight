/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4929
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// WorkflowSelectorProperty Selector properties to define HTTP method and 'body' in case of upsert (short for update or insert) operations. When requested resource (endpoint) does not exists, the user interface shows empty list.
type WorkflowSelectorProperty struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Content of the request body to send for POST request.
	Body interface{} `json:"Body,omitempty"`
	// The HTTP method to be used. * `GET` - The HTTP GET method requests a representation of the specified resource. * `POST` - The HTTP POST method sends data to the server.
	Method               *string `json:"Method,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowSelectorProperty WorkflowSelectorProperty

// NewWorkflowSelectorProperty instantiates a new WorkflowSelectorProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSelectorProperty(classId string, objectType string) *WorkflowSelectorProperty {
	this := WorkflowSelectorProperty{}
	this.ClassId = classId
	this.ObjectType = objectType
	var method string = "GET"
	this.Method = &method
	return &this
}

// NewWorkflowSelectorPropertyWithDefaults instantiates a new WorkflowSelectorProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSelectorPropertyWithDefaults() *WorkflowSelectorProperty {
	this := WorkflowSelectorProperty{}
	var classId string = "workflow.SelectorProperty"
	this.ClassId = classId
	var objectType string = "workflow.SelectorProperty"
	this.ObjectType = objectType
	var method string = "GET"
	this.Method = &method
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowSelectorProperty) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowSelectorProperty) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowSelectorProperty) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowSelectorProperty) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowSelectorProperty) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowSelectorProperty) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowSelectorProperty) GetBody() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowSelectorProperty) GetBodyOk() (*interface{}, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return &o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *WorkflowSelectorProperty) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given interface{} and assigns it to the Body field.
func (o *WorkflowSelectorProperty) SetBody(v interface{}) {
	o.Body = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *WorkflowSelectorProperty) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSelectorProperty) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *WorkflowSelectorProperty) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *WorkflowSelectorProperty) SetMethod(v string) {
	o.Method = &v
}

func (o WorkflowSelectorProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Body != nil {
		toSerialize["Body"] = o.Body
	}
	if o.Method != nil {
		toSerialize["Method"] = o.Method
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowSelectorProperty) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowSelectorPropertyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Content of the request body to send for POST request.
		Body interface{} `json:"Body,omitempty"`
		// The HTTP method to be used. * `GET` - The HTTP GET method requests a representation of the specified resource. * `POST` - The HTTP POST method sends data to the server.
		Method *string `json:"Method,omitempty"`
	}

	varWorkflowSelectorPropertyWithoutEmbeddedStruct := WorkflowSelectorPropertyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowSelectorPropertyWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowSelectorProperty := _WorkflowSelectorProperty{}
		varWorkflowSelectorProperty.ClassId = varWorkflowSelectorPropertyWithoutEmbeddedStruct.ClassId
		varWorkflowSelectorProperty.ObjectType = varWorkflowSelectorPropertyWithoutEmbeddedStruct.ObjectType
		varWorkflowSelectorProperty.Body = varWorkflowSelectorPropertyWithoutEmbeddedStruct.Body
		varWorkflowSelectorProperty.Method = varWorkflowSelectorPropertyWithoutEmbeddedStruct.Method
		*o = WorkflowSelectorProperty(varWorkflowSelectorProperty)
	} else {
		return err
	}

	varWorkflowSelectorProperty := _WorkflowSelectorProperty{}

	err = json.Unmarshal(bytes, &varWorkflowSelectorProperty)
	if err == nil {
		o.MoBaseComplexType = varWorkflowSelectorProperty.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Body")
		delete(additionalProperties, "Method")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableWorkflowSelectorProperty struct {
	value *WorkflowSelectorProperty
	isSet bool
}

func (v NullableWorkflowSelectorProperty) Get() *WorkflowSelectorProperty {
	return v.value
}

func (v *NullableWorkflowSelectorProperty) Set(val *WorkflowSelectorProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSelectorProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSelectorProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSelectorProperty(val *WorkflowSelectorProperty) *NullableWorkflowSelectorProperty {
	return &NullableWorkflowSelectorProperty{value: val, isSet: true}
}

func (v NullableWorkflowSelectorProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSelectorProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
