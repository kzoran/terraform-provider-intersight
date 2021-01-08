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

// TelemetryDruidFilteredAggregatorAllOf struct for TelemetryDruidFilteredAggregatorAllOf
type TelemetryDruidFilteredAggregatorAllOf struct {
	Filter               TelemetryDruidFilter     `json:"filter"`
	Aggregator           TelemetryDruidAggregator `json:"aggregator"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidFilteredAggregatorAllOf TelemetryDruidFilteredAggregatorAllOf

// NewTelemetryDruidFilteredAggregatorAllOf instantiates a new TelemetryDruidFilteredAggregatorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidFilteredAggregatorAllOf(filter TelemetryDruidFilter, aggregator TelemetryDruidAggregator) *TelemetryDruidFilteredAggregatorAllOf {
	this := TelemetryDruidFilteredAggregatorAllOf{}
	this.Filter = filter
	this.Aggregator = aggregator
	return &this
}

// NewTelemetryDruidFilteredAggregatorAllOfWithDefaults instantiates a new TelemetryDruidFilteredAggregatorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidFilteredAggregatorAllOfWithDefaults() *TelemetryDruidFilteredAggregatorAllOf {
	this := TelemetryDruidFilteredAggregatorAllOf{}
	return &this
}

// GetFilter returns the Filter field value
func (o *TelemetryDruidFilteredAggregatorAllOf) GetFilter() TelemetryDruidFilter {
	if o == nil {
		var ret TelemetryDruidFilter
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidFilteredAggregatorAllOf) GetFilterOk() (*TelemetryDruidFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *TelemetryDruidFilteredAggregatorAllOf) SetFilter(v TelemetryDruidFilter) {
	o.Filter = v
}

// GetAggregator returns the Aggregator field value
func (o *TelemetryDruidFilteredAggregatorAllOf) GetAggregator() TelemetryDruidAggregator {
	if o == nil {
		var ret TelemetryDruidAggregator
		return ret
	}

	return o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidFilteredAggregatorAllOf) GetAggregatorOk() (*TelemetryDruidAggregator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregator, true
}

// SetAggregator sets field value
func (o *TelemetryDruidFilteredAggregatorAllOf) SetAggregator(v TelemetryDruidAggregator) {
	o.Aggregator = v
}

func (o TelemetryDruidFilteredAggregatorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["filter"] = o.Filter
	}
	if true {
		toSerialize["aggregator"] = o.Aggregator
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidFilteredAggregatorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidFilteredAggregatorAllOf := _TelemetryDruidFilteredAggregatorAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidFilteredAggregatorAllOf); err == nil {
		*o = TelemetryDruidFilteredAggregatorAllOf(varTelemetryDruidFilteredAggregatorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "filter")
		delete(additionalProperties, "aggregator")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidFilteredAggregatorAllOf struct {
	value *TelemetryDruidFilteredAggregatorAllOf
	isSet bool
}

func (v NullableTelemetryDruidFilteredAggregatorAllOf) Get() *TelemetryDruidFilteredAggregatorAllOf {
	return v.value
}

func (v *NullableTelemetryDruidFilteredAggregatorAllOf) Set(val *TelemetryDruidFilteredAggregatorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidFilteredAggregatorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidFilteredAggregatorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidFilteredAggregatorAllOf(val *TelemetryDruidFilteredAggregatorAllOf) *NullableTelemetryDruidFilteredAggregatorAllOf {
	return &NullableTelemetryDruidFilteredAggregatorAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidFilteredAggregatorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidFilteredAggregatorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
