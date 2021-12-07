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

// AssetWorkloadOptimizerNewRelicOptions Captures configuration specific to the NewRelic target for the Workload Optimizer service.
type AssetWorkloadOptimizerNewRelicOptions struct {
	AssetServiceOptions
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Your NewRelic account id.
	AccountId *string `json:"AccountId,omitempty"`
	// The region associated with the NewRelic account. * `US` - The United States (US) region. * `EU` - The European Union (EU) region.
	Region               *string `json:"Region,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetWorkloadOptimizerNewRelicOptions AssetWorkloadOptimizerNewRelicOptions

// NewAssetWorkloadOptimizerNewRelicOptions instantiates a new AssetWorkloadOptimizerNewRelicOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWorkloadOptimizerNewRelicOptions(classId string, objectType string) *AssetWorkloadOptimizerNewRelicOptions {
	this := AssetWorkloadOptimizerNewRelicOptions{}
	this.ClassId = classId
	this.ObjectType = objectType
	var region string = "US"
	this.Region = &region
	return &this
}

// NewAssetWorkloadOptimizerNewRelicOptionsWithDefaults instantiates a new AssetWorkloadOptimizerNewRelicOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWorkloadOptimizerNewRelicOptionsWithDefaults() *AssetWorkloadOptimizerNewRelicOptions {
	this := AssetWorkloadOptimizerNewRelicOptions{}
	var classId string = "asset.WorkloadOptimizerNewRelicOptions"
	this.ClassId = classId
	var objectType string = "asset.WorkloadOptimizerNewRelicOptions"
	this.ObjectType = objectType
	var region string = "US"
	this.Region = &region
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetWorkloadOptimizerNewRelicOptions) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerNewRelicOptions) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetWorkloadOptimizerNewRelicOptions) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetWorkloadOptimizerNewRelicOptions) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerNewRelicOptions) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetWorkloadOptimizerNewRelicOptions) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerNewRelicOptions) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerNewRelicOptions) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerNewRelicOptions) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AssetWorkloadOptimizerNewRelicOptions) SetAccountId(v string) {
	o.AccountId = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerNewRelicOptions) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerNewRelicOptions) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerNewRelicOptions) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AssetWorkloadOptimizerNewRelicOptions) SetRegion(v string) {
	o.Region = &v
}

func (o AssetWorkloadOptimizerNewRelicOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetServiceOptions, errAssetServiceOptions := json.Marshal(o.AssetServiceOptions)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	errAssetServiceOptions = json.Unmarshal([]byte(serializedAssetServiceOptions), &toSerialize)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AccountId != nil {
		toSerialize["AccountId"] = o.AccountId
	}
	if o.Region != nil {
		toSerialize["Region"] = o.Region
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetWorkloadOptimizerNewRelicOptions) UnmarshalJSON(bytes []byte) (err error) {
	type AssetWorkloadOptimizerNewRelicOptionsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Your NewRelic account id.
		AccountId *string `json:"AccountId,omitempty"`
		// The region associated with the NewRelic account. * `US` - The United States (US) region. * `EU` - The European Union (EU) region.
		Region *string `json:"Region,omitempty"`
	}

	varAssetWorkloadOptimizerNewRelicOptionsWithoutEmbeddedStruct := AssetWorkloadOptimizerNewRelicOptionsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerNewRelicOptionsWithoutEmbeddedStruct)
	if err == nil {
		varAssetWorkloadOptimizerNewRelicOptions := _AssetWorkloadOptimizerNewRelicOptions{}
		varAssetWorkloadOptimizerNewRelicOptions.ClassId = varAssetWorkloadOptimizerNewRelicOptionsWithoutEmbeddedStruct.ClassId
		varAssetWorkloadOptimizerNewRelicOptions.ObjectType = varAssetWorkloadOptimizerNewRelicOptionsWithoutEmbeddedStruct.ObjectType
		varAssetWorkloadOptimizerNewRelicOptions.AccountId = varAssetWorkloadOptimizerNewRelicOptionsWithoutEmbeddedStruct.AccountId
		varAssetWorkloadOptimizerNewRelicOptions.Region = varAssetWorkloadOptimizerNewRelicOptionsWithoutEmbeddedStruct.Region
		*o = AssetWorkloadOptimizerNewRelicOptions(varAssetWorkloadOptimizerNewRelicOptions)
	} else {
		return err
	}

	varAssetWorkloadOptimizerNewRelicOptions := _AssetWorkloadOptimizerNewRelicOptions{}

	err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerNewRelicOptions)
	if err == nil {
		o.AssetServiceOptions = varAssetWorkloadOptimizerNewRelicOptions.AssetServiceOptions
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccountId")
		delete(additionalProperties, "Region")

		// remove fields from embedded structs
		reflectAssetServiceOptions := reflect.ValueOf(o.AssetServiceOptions)
		for i := 0; i < reflectAssetServiceOptions.Type().NumField(); i++ {
			t := reflectAssetServiceOptions.Type().Field(i)

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

type NullableAssetWorkloadOptimizerNewRelicOptions struct {
	value *AssetWorkloadOptimizerNewRelicOptions
	isSet bool
}

func (v NullableAssetWorkloadOptimizerNewRelicOptions) Get() *AssetWorkloadOptimizerNewRelicOptions {
	return v.value
}

func (v *NullableAssetWorkloadOptimizerNewRelicOptions) Set(val *AssetWorkloadOptimizerNewRelicOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWorkloadOptimizerNewRelicOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWorkloadOptimizerNewRelicOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWorkloadOptimizerNewRelicOptions(val *AssetWorkloadOptimizerNewRelicOptions) *NullableAssetWorkloadOptimizerNewRelicOptions {
	return &NullableAssetWorkloadOptimizerNewRelicOptions{value: val, isSet: true}
}

func (v NullableAssetWorkloadOptimizerNewRelicOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWorkloadOptimizerNewRelicOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
