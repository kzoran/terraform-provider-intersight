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

// ConnectorpackConnectorPackUpgrade Used to download or install connector packs on the target device.
type ConnectorpackConnectorPackUpgrade struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The type of operation to be performed on UCS Director. * `Install` - Installs the requisite connector packs on UCS Director. * `Push` - Pushes the requisite connector packs to UCS Director.
	ConnectorPackOpType  *string                           `json:"ConnectorPackOpType,omitempty"`
	UcsdInfo             *IaasUcsdInfoRelationship         `json:"UcsdInfo,omitempty"`
	Workflow             *WorkflowWorkflowInfoRelationship `json:"Workflow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorpackConnectorPackUpgrade ConnectorpackConnectorPackUpgrade

// NewConnectorpackConnectorPackUpgrade instantiates a new ConnectorpackConnectorPackUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorpackConnectorPackUpgrade(classId string, objectType string) *ConnectorpackConnectorPackUpgrade {
	this := ConnectorpackConnectorPackUpgrade{}
	this.ClassId = classId
	this.ObjectType = objectType
	var connectorPackOpType string = "Install"
	this.ConnectorPackOpType = &connectorPackOpType
	return &this
}

// NewConnectorpackConnectorPackUpgradeWithDefaults instantiates a new ConnectorpackConnectorPackUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorpackConnectorPackUpgradeWithDefaults() *ConnectorpackConnectorPackUpgrade {
	this := ConnectorpackConnectorPackUpgrade{}
	var classId string = "connectorpack.ConnectorPackUpgrade"
	this.ClassId = classId
	var objectType string = "connectorpack.ConnectorPackUpgrade"
	this.ObjectType = objectType
	var connectorPackOpType string = "Install"
	this.ConnectorPackOpType = &connectorPackOpType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorpackConnectorPackUpgrade) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorpackConnectorPackUpgrade) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorpackConnectorPackUpgrade) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorpackConnectorPackUpgrade) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorpackConnectorPackUpgrade) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorpackConnectorPackUpgrade) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConnectorPackOpType returns the ConnectorPackOpType field value if set, zero value otherwise.
func (o *ConnectorpackConnectorPackUpgrade) GetConnectorPackOpType() string {
	if o == nil || o.ConnectorPackOpType == nil {
		var ret string
		return ret
	}
	return *o.ConnectorPackOpType
}

// GetConnectorPackOpTypeOk returns a tuple with the ConnectorPackOpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorpackConnectorPackUpgrade) GetConnectorPackOpTypeOk() (*string, bool) {
	if o == nil || o.ConnectorPackOpType == nil {
		return nil, false
	}
	return o.ConnectorPackOpType, true
}

// HasConnectorPackOpType returns a boolean if a field has been set.
func (o *ConnectorpackConnectorPackUpgrade) HasConnectorPackOpType() bool {
	if o != nil && o.ConnectorPackOpType != nil {
		return true
	}

	return false
}

// SetConnectorPackOpType gets a reference to the given string and assigns it to the ConnectorPackOpType field.
func (o *ConnectorpackConnectorPackUpgrade) SetConnectorPackOpType(v string) {
	o.ConnectorPackOpType = &v
}

// GetUcsdInfo returns the UcsdInfo field value if set, zero value otherwise.
func (o *ConnectorpackConnectorPackUpgrade) GetUcsdInfo() IaasUcsdInfoRelationship {
	if o == nil || o.UcsdInfo == nil {
		var ret IaasUcsdInfoRelationship
		return ret
	}
	return *o.UcsdInfo
}

// GetUcsdInfoOk returns a tuple with the UcsdInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorpackConnectorPackUpgrade) GetUcsdInfoOk() (*IaasUcsdInfoRelationship, bool) {
	if o == nil || o.UcsdInfo == nil {
		return nil, false
	}
	return o.UcsdInfo, true
}

// HasUcsdInfo returns a boolean if a field has been set.
func (o *ConnectorpackConnectorPackUpgrade) HasUcsdInfo() bool {
	if o != nil && o.UcsdInfo != nil {
		return true
	}

	return false
}

// SetUcsdInfo gets a reference to the given IaasUcsdInfoRelationship and assigns it to the UcsdInfo field.
func (o *ConnectorpackConnectorPackUpgrade) SetUcsdInfo(v IaasUcsdInfoRelationship) {
	o.UcsdInfo = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *ConnectorpackConnectorPackUpgrade) GetWorkflow() WorkflowWorkflowInfoRelationship {
	if o == nil || o.Workflow == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorpackConnectorPackUpgrade) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.Workflow == nil {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *ConnectorpackConnectorPackUpgrade) HasWorkflow() bool {
	if o != nil && o.Workflow != nil {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the Workflow field.
func (o *ConnectorpackConnectorPackUpgrade) SetWorkflow(v WorkflowWorkflowInfoRelationship) {
	o.Workflow = &v
}

func (o ConnectorpackConnectorPackUpgrade) MarshalJSON() ([]byte, error) {
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
	if o.ConnectorPackOpType != nil {
		toSerialize["ConnectorPackOpType"] = o.ConnectorPackOpType
	}
	if o.UcsdInfo != nil {
		toSerialize["UcsdInfo"] = o.UcsdInfo
	}
	if o.Workflow != nil {
		toSerialize["Workflow"] = o.Workflow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorpackConnectorPackUpgrade) UnmarshalJSON(bytes []byte) (err error) {
	type ConnectorpackConnectorPackUpgradeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The type of operation to be performed on UCS Director. * `Install` - Installs the requisite connector packs on UCS Director. * `Push` - Pushes the requisite connector packs to UCS Director.
		ConnectorPackOpType *string                           `json:"ConnectorPackOpType,omitempty"`
		UcsdInfo            *IaasUcsdInfoRelationship         `json:"UcsdInfo,omitempty"`
		Workflow            *WorkflowWorkflowInfoRelationship `json:"Workflow,omitempty"`
	}

	varConnectorpackConnectorPackUpgradeWithoutEmbeddedStruct := ConnectorpackConnectorPackUpgradeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConnectorpackConnectorPackUpgradeWithoutEmbeddedStruct)
	if err == nil {
		varConnectorpackConnectorPackUpgrade := _ConnectorpackConnectorPackUpgrade{}
		varConnectorpackConnectorPackUpgrade.ClassId = varConnectorpackConnectorPackUpgradeWithoutEmbeddedStruct.ClassId
		varConnectorpackConnectorPackUpgrade.ObjectType = varConnectorpackConnectorPackUpgradeWithoutEmbeddedStruct.ObjectType
		varConnectorpackConnectorPackUpgrade.ConnectorPackOpType = varConnectorpackConnectorPackUpgradeWithoutEmbeddedStruct.ConnectorPackOpType
		varConnectorpackConnectorPackUpgrade.UcsdInfo = varConnectorpackConnectorPackUpgradeWithoutEmbeddedStruct.UcsdInfo
		varConnectorpackConnectorPackUpgrade.Workflow = varConnectorpackConnectorPackUpgradeWithoutEmbeddedStruct.Workflow
		*o = ConnectorpackConnectorPackUpgrade(varConnectorpackConnectorPackUpgrade)
	} else {
		return err
	}

	varConnectorpackConnectorPackUpgrade := _ConnectorpackConnectorPackUpgrade{}

	err = json.Unmarshal(bytes, &varConnectorpackConnectorPackUpgrade)
	if err == nil {
		o.MoBaseMo = varConnectorpackConnectorPackUpgrade.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConnectorPackOpType")
		delete(additionalProperties, "UcsdInfo")
		delete(additionalProperties, "Workflow")

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

type NullableConnectorpackConnectorPackUpgrade struct {
	value *ConnectorpackConnectorPackUpgrade
	isSet bool
}

func (v NullableConnectorpackConnectorPackUpgrade) Get() *ConnectorpackConnectorPackUpgrade {
	return v.value
}

func (v *NullableConnectorpackConnectorPackUpgrade) Set(val *ConnectorpackConnectorPackUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorpackConnectorPackUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorpackConnectorPackUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorpackConnectorPackUpgrade(val *ConnectorpackConnectorPackUpgrade) *NullableConnectorpackConnectorPackUpgrade {
	return &NullableConnectorpackConnectorPackUpgrade{value: val, isSet: true}
}

func (v NullableConnectorpackConnectorPackUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorpackConnectorPackUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
