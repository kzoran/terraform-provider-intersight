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

// TelemetryDruidUnionDataSourceAllOf struct for TelemetryDruidUnionDataSourceAllOf
type TelemetryDruidUnionDataSourceAllOf struct {
	// A list of data sources.
	DataSources          []string `json:"dataSources"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidUnionDataSourceAllOf TelemetryDruidUnionDataSourceAllOf

// NewTelemetryDruidUnionDataSourceAllOf instantiates a new TelemetryDruidUnionDataSourceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidUnionDataSourceAllOf(dataSources []string) *TelemetryDruidUnionDataSourceAllOf {
	this := TelemetryDruidUnionDataSourceAllOf{}
	this.DataSources = dataSources
	return &this
}

// NewTelemetryDruidUnionDataSourceAllOfWithDefaults instantiates a new TelemetryDruidUnionDataSourceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidUnionDataSourceAllOfWithDefaults() *TelemetryDruidUnionDataSourceAllOf {
	this := TelemetryDruidUnionDataSourceAllOf{}
	return &this
}

// GetDataSources returns the DataSources field value
func (o *TelemetryDruidUnionDataSourceAllOf) GetDataSources() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DataSources
}

// GetDataSourcesOk returns a tuple with the DataSources field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidUnionDataSourceAllOf) GetDataSourcesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSources, true
}

// SetDataSources sets field value
func (o *TelemetryDruidUnionDataSourceAllOf) SetDataSources(v []string) {
	o.DataSources = v
}

func (o TelemetryDruidUnionDataSourceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dataSources"] = o.DataSources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidUnionDataSourceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidUnionDataSourceAllOf := _TelemetryDruidUnionDataSourceAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidUnionDataSourceAllOf); err == nil {
		*o = TelemetryDruidUnionDataSourceAllOf(varTelemetryDruidUnionDataSourceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dataSources")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidUnionDataSourceAllOf struct {
	value *TelemetryDruidUnionDataSourceAllOf
	isSet bool
}

func (v NullableTelemetryDruidUnionDataSourceAllOf) Get() *TelemetryDruidUnionDataSourceAllOf {
	return v.value
}

func (v *NullableTelemetryDruidUnionDataSourceAllOf) Set(val *TelemetryDruidUnionDataSourceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidUnionDataSourceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidUnionDataSourceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidUnionDataSourceAllOf(val *TelemetryDruidUnionDataSourceAllOf) *NullableTelemetryDruidUnionDataSourceAllOf {
	return &NullableTelemetryDruidUnionDataSourceAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidUnionDataSourceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidUnionDataSourceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
