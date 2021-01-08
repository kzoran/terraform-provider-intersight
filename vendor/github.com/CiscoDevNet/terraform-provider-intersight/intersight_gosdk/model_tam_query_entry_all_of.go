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

// TamQueryEntryAllOf Definition of the list of properties defined in 'tam.QueryEntry', excluding properties defined in parent classes.
type TamQueryEntryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source.
	Name *string `json:"Name,omitempty"`
	// An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection.
	Priority *int64 `json:"Priority,omitempty"`
	// A SparkSQL query to be used on a given data source.
	Query                *string `json:"Query,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamQueryEntryAllOf TamQueryEntryAllOf

// NewTamQueryEntryAllOf instantiates a new TamQueryEntryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamQueryEntryAllOf(classId string, objectType string) *TamQueryEntryAllOf {
	this := TamQueryEntryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTamQueryEntryAllOfWithDefaults instantiates a new TamQueryEntryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamQueryEntryAllOfWithDefaults() *TamQueryEntryAllOf {
	this := TamQueryEntryAllOf{}
	var classId string = "tam.QueryEntry"
	this.ClassId = classId
	var objectType string = "tam.QueryEntry"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamQueryEntryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamQueryEntryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamQueryEntryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TamQueryEntryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamQueryEntryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamQueryEntryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TamQueryEntryAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamQueryEntryAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TamQueryEntryAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TamQueryEntryAllOf) SetName(v string) {
	o.Name = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *TamQueryEntryAllOf) GetPriority() int64 {
	if o == nil || o.Priority == nil {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamQueryEntryAllOf) GetPriorityOk() (*int64, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *TamQueryEntryAllOf) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *TamQueryEntryAllOf) SetPriority(v int64) {
	o.Priority = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *TamQueryEntryAllOf) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamQueryEntryAllOf) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *TamQueryEntryAllOf) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *TamQueryEntryAllOf) SetQuery(v string) {
	o.Query = &v
}

func (o TamQueryEntryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Priority != nil {
		toSerialize["Priority"] = o.Priority
	}
	if o.Query != nil {
		toSerialize["Query"] = o.Query
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamQueryEntryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTamQueryEntryAllOf := _TamQueryEntryAllOf{}

	if err = json.Unmarshal(bytes, &varTamQueryEntryAllOf); err == nil {
		*o = TamQueryEntryAllOf(varTamQueryEntryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Priority")
		delete(additionalProperties, "Query")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTamQueryEntryAllOf struct {
	value *TamQueryEntryAllOf
	isSet bool
}

func (v NullableTamQueryEntryAllOf) Get() *TamQueryEntryAllOf {
	return v.value
}

func (v *NullableTamQueryEntryAllOf) Set(val *TamQueryEntryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTamQueryEntryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTamQueryEntryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamQueryEntryAllOf(val *TamQueryEntryAllOf) *NullableTamQueryEntryAllOf {
	return &NullableTamQueryEntryAllOf{value: val, isSet: true}
}

func (v NullableTamQueryEntryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamQueryEntryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
