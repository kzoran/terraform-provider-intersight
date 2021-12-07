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

// OsValidationInformation Used to show the state of the CSV validaton and its state.
type OsValidationInformation struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Validation error message.
	ErrorMsg *string `json:"ErrorMsg,omitempty"`
	// The status of the validation step. * `NotValidated` - The validation not started. * `Valid` - The step status marked as valid when respective step validation condition is successful. * `Invalid` - The step status marked as invalid when respective step validation condition is failed. * `InProgress` - The validation is in progress.
	Status *string `json:"Status,omitempty"`
	// The validation step name. * `OS Install Schema Validation` - The step to validate the CSV file schema. * `OS Image Validation` - The Operating System Image parameter validation step. * `SCU Image Validation` - The SCU Image parameter validation step. * `Configuration source and file validation` - The Configuration Source and Configuration file validation step. * `Server level data validation` - The server level parameters validation.
	StepName             *string `json:"StepName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsValidationInformation OsValidationInformation

// NewOsValidationInformation instantiates a new OsValidationInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsValidationInformation(classId string, objectType string) *OsValidationInformation {
	this := OsValidationInformation{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsValidationInformationWithDefaults instantiates a new OsValidationInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsValidationInformationWithDefaults() *OsValidationInformation {
	this := OsValidationInformation{}
	var classId string = "os.ValidationInformation"
	this.ClassId = classId
	var objectType string = "os.ValidationInformation"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsValidationInformation) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsValidationInformation) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsValidationInformation) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsValidationInformation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsValidationInformation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsValidationInformation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetErrorMsg returns the ErrorMsg field value if set, zero value otherwise.
func (o *OsValidationInformation) GetErrorMsg() string {
	if o == nil || o.ErrorMsg == nil {
		var ret string
		return ret
	}
	return *o.ErrorMsg
}

// GetErrorMsgOk returns a tuple with the ErrorMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsValidationInformation) GetErrorMsgOk() (*string, bool) {
	if o == nil || o.ErrorMsg == nil {
		return nil, false
	}
	return o.ErrorMsg, true
}

// HasErrorMsg returns a boolean if a field has been set.
func (o *OsValidationInformation) HasErrorMsg() bool {
	if o != nil && o.ErrorMsg != nil {
		return true
	}

	return false
}

// SetErrorMsg gets a reference to the given string and assigns it to the ErrorMsg field.
func (o *OsValidationInformation) SetErrorMsg(v string) {
	o.ErrorMsg = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OsValidationInformation) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsValidationInformation) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OsValidationInformation) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OsValidationInformation) SetStatus(v string) {
	o.Status = &v
}

// GetStepName returns the StepName field value if set, zero value otherwise.
func (o *OsValidationInformation) GetStepName() string {
	if o == nil || o.StepName == nil {
		var ret string
		return ret
	}
	return *o.StepName
}

// GetStepNameOk returns a tuple with the StepName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsValidationInformation) GetStepNameOk() (*string, bool) {
	if o == nil || o.StepName == nil {
		return nil, false
	}
	return o.StepName, true
}

// HasStepName returns a boolean if a field has been set.
func (o *OsValidationInformation) HasStepName() bool {
	if o != nil && o.StepName != nil {
		return true
	}

	return false
}

// SetStepName gets a reference to the given string and assigns it to the StepName field.
func (o *OsValidationInformation) SetStepName(v string) {
	o.StepName = &v
}

func (o OsValidationInformation) MarshalJSON() ([]byte, error) {
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
	if o.ErrorMsg != nil {
		toSerialize["ErrorMsg"] = o.ErrorMsg
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StepName != nil {
		toSerialize["StepName"] = o.StepName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsValidationInformation) UnmarshalJSON(bytes []byte) (err error) {
	type OsValidationInformationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Validation error message.
		ErrorMsg *string `json:"ErrorMsg,omitempty"`
		// The status of the validation step. * `NotValidated` - The validation not started. * `Valid` - The step status marked as valid when respective step validation condition is successful. * `Invalid` - The step status marked as invalid when respective step validation condition is failed. * `InProgress` - The validation is in progress.
		Status *string `json:"Status,omitempty"`
		// The validation step name. * `OS Install Schema Validation` - The step to validate the CSV file schema. * `OS Image Validation` - The Operating System Image parameter validation step. * `SCU Image Validation` - The SCU Image parameter validation step. * `Configuration source and file validation` - The Configuration Source and Configuration file validation step. * `Server level data validation` - The server level parameters validation.
		StepName *string `json:"StepName,omitempty"`
	}

	varOsValidationInformationWithoutEmbeddedStruct := OsValidationInformationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varOsValidationInformationWithoutEmbeddedStruct)
	if err == nil {
		varOsValidationInformation := _OsValidationInformation{}
		varOsValidationInformation.ClassId = varOsValidationInformationWithoutEmbeddedStruct.ClassId
		varOsValidationInformation.ObjectType = varOsValidationInformationWithoutEmbeddedStruct.ObjectType
		varOsValidationInformation.ErrorMsg = varOsValidationInformationWithoutEmbeddedStruct.ErrorMsg
		varOsValidationInformation.Status = varOsValidationInformationWithoutEmbeddedStruct.Status
		varOsValidationInformation.StepName = varOsValidationInformationWithoutEmbeddedStruct.StepName
		*o = OsValidationInformation(varOsValidationInformation)
	} else {
		return err
	}

	varOsValidationInformation := _OsValidationInformation{}

	err = json.Unmarshal(bytes, &varOsValidationInformation)
	if err == nil {
		o.MoBaseComplexType = varOsValidationInformation.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ErrorMsg")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StepName")

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

type NullableOsValidationInformation struct {
	value *OsValidationInformation
	isSet bool
}

func (v NullableOsValidationInformation) Get() *OsValidationInformation {
	return v.value
}

func (v *NullableOsValidationInformation) Set(val *OsValidationInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableOsValidationInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableOsValidationInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsValidationInformation(val *OsValidationInformation) *NullableOsValidationInformation {
	return &NullableOsValidationInformation{value: val, isSet: true}
}

func (v NullableOsValidationInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsValidationInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
