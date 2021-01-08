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

// TelemetryDruidGroupByRequestAllOf struct for TelemetryDruidGroupByRequestAllOf
type TelemetryDruidGroupByRequestAllOf struct {
	DataSource TelemetryDruidDataSource `json:"dataSource"`
	// A JSON list of dimensions to do the groupBy over; or see DimensionSpec for ways to extract dimensions..
	Dimensions  []TelemetryDruidDimensionSpec   `json:"dimensions"`
	LimitSpec   *TelemetryDruidDefaultLimitSpec `json:"limitSpec,omitempty"`
	Having      *TelemetryDruidHavingFilter     `json:"having,omitempty"`
	Granularity TelemetryDruidGranularity       `json:"granularity"`
	Filter      *TelemetryDruidFilter           `json:"filter,omitempty"`
	// Aggregation functions are used to summarize data in buckets. Summarization functions include counting rows, calculating the min/max/sum of metrics and retrieving the first/last value of metrics for each bucket. Additional summarization functions are available with extensions. If no aggregator is provided, the results will be empty for each bucket.
	Aggregations *[]TelemetryDruidAggregator `json:"aggregations,omitempty"`
	// Post-aggregations are specifications of processing that should happen on aggregated values as they come out of Apache Druid. If you include a post aggregation as part of a query, make sure to include all aggregators the post-aggregator requires.
	PostAggregations *[]TelemetryDruidPostAggregator `json:"postAggregations,omitempty"`
	// A JSON Object representing ISO-8601 Intervals. This defines the time ranges to run the query over.
	Intervals []string `json:"intervals"`
	// A JSON array of arrays to return additional result sets for groupings of subsets of top level dimensions. The subtotals feature allows computation of multiple sub-groupings in a single query. To use this feature, add a \"subtotalsSpec\" to your query, which should be a list of subgroup dimension sets. It should contain the \"outputName\" from dimensions in your \"dimensions\" attribute, in the same order as they appear in the \"dimensions\" attribute.
	SubtotalsSpec        *map[string]interface{}     `json:"subtotalsSpec,omitempty"`
	Context              *TelemetryDruidQueryContext `json:"context,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidGroupByRequestAllOf TelemetryDruidGroupByRequestAllOf

// NewTelemetryDruidGroupByRequestAllOf instantiates a new TelemetryDruidGroupByRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidGroupByRequestAllOf(dataSource TelemetryDruidDataSource, dimensions []TelemetryDruidDimensionSpec, granularity TelemetryDruidGranularity, intervals []string) *TelemetryDruidGroupByRequestAllOf {
	this := TelemetryDruidGroupByRequestAllOf{}
	this.DataSource = dataSource
	this.Dimensions = dimensions
	this.Granularity = granularity
	this.Intervals = intervals
	return &this
}

// NewTelemetryDruidGroupByRequestAllOfWithDefaults instantiates a new TelemetryDruidGroupByRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidGroupByRequestAllOfWithDefaults() *TelemetryDruidGroupByRequestAllOf {
	this := TelemetryDruidGroupByRequestAllOf{}
	return &this
}

// GetDataSource returns the DataSource field value
func (o *TelemetryDruidGroupByRequestAllOf) GetDataSource() TelemetryDruidDataSource {
	if o == nil {
		var ret TelemetryDruidDataSource
		return ret
	}

	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequestAllOf) GetDataSourceOk() (*TelemetryDruidDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value
func (o *TelemetryDruidGroupByRequestAllOf) SetDataSource(v TelemetryDruidDataSource) {
	o.DataSource = v
}

// GetDimensions returns the Dimensions field value
func (o *TelemetryDruidGroupByRequestAllOf) GetDimensions() []TelemetryDruidDimensionSpec {
	if o == nil {
		var ret []TelemetryDruidDimensionSpec
		return ret
	}

	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequestAllOf) GetDimensionsOk() (*[]TelemetryDruidDimensionSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value
func (o *TelemetryDruidGroupByRequestAllOf) SetDimensions(v []TelemetryDruidDimensionSpec) {
	o.Dimensions = v
}

// GetLimitSpec returns the LimitSpec field value if set, zero value otherwise.
func (o *TelemetryDruidGroupByRequestAllOf) GetLimitSpec() TelemetryDruidDefaultLimitSpec {
	if o == nil || o.LimitSpec == nil {
		var ret TelemetryDruidDefaultLimitSpec
		return ret
	}
	return *o.LimitSpec
}

// GetLimitSpecOk returns a tuple with the LimitSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequestAllOf) GetLimitSpecOk() (*TelemetryDruidDefaultLimitSpec, bool) {
	if o == nil || o.LimitSpec == nil {
		return nil, false
	}
	return o.LimitSpec, true
}

// HasLimitSpec returns a boolean if a field has been set.
func (o *TelemetryDruidGroupByRequestAllOf) HasLimitSpec() bool {
	if o != nil && o.LimitSpec != nil {
		return true
	}

	return false
}

// SetLimitSpec gets a reference to the given TelemetryDruidDefaultLimitSpec and assigns it to the LimitSpec field.
func (o *TelemetryDruidGroupByRequestAllOf) SetLimitSpec(v TelemetryDruidDefaultLimitSpec) {
	o.LimitSpec = &v
}

// GetHaving returns the Having field value if set, zero value otherwise.
func (o *TelemetryDruidGroupByRequestAllOf) GetHaving() TelemetryDruidHavingFilter {
	if o == nil || o.Having == nil {
		var ret TelemetryDruidHavingFilter
		return ret
	}
	return *o.Having
}

// GetHavingOk returns a tuple with the Having field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequestAllOf) GetHavingOk() (*TelemetryDruidHavingFilter, bool) {
	if o == nil || o.Having == nil {
		return nil, false
	}
	return o.Having, true
}

// HasHaving returns a boolean if a field has been set.
func (o *TelemetryDruidGroupByRequestAllOf) HasHaving() bool {
	if o != nil && o.Having != nil {
		return true
	}

	return false
}

// SetHaving gets a reference to the given TelemetryDruidHavingFilter and assigns it to the Having field.
func (o *TelemetryDruidGroupByRequestAllOf) SetHaving(v TelemetryDruidHavingFilter) {
	o.Having = &v
}

// GetGranularity returns the Granularity field value
func (o *TelemetryDruidGroupByRequestAllOf) GetGranularity() TelemetryDruidGranularity {
	if o == nil {
		var ret TelemetryDruidGranularity
		return ret
	}

	return o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequestAllOf) GetGranularityOk() (*TelemetryDruidGranularity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Granularity, true
}

// SetGranularity sets field value
func (o *TelemetryDruidGroupByRequestAllOf) SetGranularity(v TelemetryDruidGranularity) {
	o.Granularity = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *TelemetryDruidGroupByRequestAllOf) GetFilter() TelemetryDruidFilter {
	if o == nil || o.Filter == nil {
		var ret TelemetryDruidFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequestAllOf) GetFilterOk() (*TelemetryDruidFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *TelemetryDruidGroupByRequestAllOf) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given TelemetryDruidFilter and assigns it to the Filter field.
func (o *TelemetryDruidGroupByRequestAllOf) SetFilter(v TelemetryDruidFilter) {
	o.Filter = &v
}

// GetAggregations returns the Aggregations field value if set, zero value otherwise.
func (o *TelemetryDruidGroupByRequestAllOf) GetAggregations() []TelemetryDruidAggregator {
	if o == nil || o.Aggregations == nil {
		var ret []TelemetryDruidAggregator
		return ret
	}
	return *o.Aggregations
}

// GetAggregationsOk returns a tuple with the Aggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequestAllOf) GetAggregationsOk() (*[]TelemetryDruidAggregator, bool) {
	if o == nil || o.Aggregations == nil {
		return nil, false
	}
	return o.Aggregations, true
}

// HasAggregations returns a boolean if a field has been set.
func (o *TelemetryDruidGroupByRequestAllOf) HasAggregations() bool {
	if o != nil && o.Aggregations != nil {
		return true
	}

	return false
}

// SetAggregations gets a reference to the given []TelemetryDruidAggregator and assigns it to the Aggregations field.
func (o *TelemetryDruidGroupByRequestAllOf) SetAggregations(v []TelemetryDruidAggregator) {
	o.Aggregations = &v
}

// GetPostAggregations returns the PostAggregations field value if set, zero value otherwise.
func (o *TelemetryDruidGroupByRequestAllOf) GetPostAggregations() []TelemetryDruidPostAggregator {
	if o == nil || o.PostAggregations == nil {
		var ret []TelemetryDruidPostAggregator
		return ret
	}
	return *o.PostAggregations
}

// GetPostAggregationsOk returns a tuple with the PostAggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequestAllOf) GetPostAggregationsOk() (*[]TelemetryDruidPostAggregator, bool) {
	if o == nil || o.PostAggregations == nil {
		return nil, false
	}
	return o.PostAggregations, true
}

// HasPostAggregations returns a boolean if a field has been set.
func (o *TelemetryDruidGroupByRequestAllOf) HasPostAggregations() bool {
	if o != nil && o.PostAggregations != nil {
		return true
	}

	return false
}

// SetPostAggregations gets a reference to the given []TelemetryDruidPostAggregator and assigns it to the PostAggregations field.
func (o *TelemetryDruidGroupByRequestAllOf) SetPostAggregations(v []TelemetryDruidPostAggregator) {
	o.PostAggregations = &v
}

// GetIntervals returns the Intervals field value
func (o *TelemetryDruidGroupByRequestAllOf) GetIntervals() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Intervals
}

// GetIntervalsOk returns a tuple with the Intervals field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequestAllOf) GetIntervalsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Intervals, true
}

// SetIntervals sets field value
func (o *TelemetryDruidGroupByRequestAllOf) SetIntervals(v []string) {
	o.Intervals = v
}

// GetSubtotalsSpec returns the SubtotalsSpec field value if set, zero value otherwise.
func (o *TelemetryDruidGroupByRequestAllOf) GetSubtotalsSpec() map[string]interface{} {
	if o == nil || o.SubtotalsSpec == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.SubtotalsSpec
}

// GetSubtotalsSpecOk returns a tuple with the SubtotalsSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequestAllOf) GetSubtotalsSpecOk() (*map[string]interface{}, bool) {
	if o == nil || o.SubtotalsSpec == nil {
		return nil, false
	}
	return o.SubtotalsSpec, true
}

// HasSubtotalsSpec returns a boolean if a field has been set.
func (o *TelemetryDruidGroupByRequestAllOf) HasSubtotalsSpec() bool {
	if o != nil && o.SubtotalsSpec != nil {
		return true
	}

	return false
}

// SetSubtotalsSpec gets a reference to the given map[string]interface{} and assigns it to the SubtotalsSpec field.
func (o *TelemetryDruidGroupByRequestAllOf) SetSubtotalsSpec(v map[string]interface{}) {
	o.SubtotalsSpec = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TelemetryDruidGroupByRequestAllOf) GetContext() TelemetryDruidQueryContext {
	if o == nil || o.Context == nil {
		var ret TelemetryDruidQueryContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByRequestAllOf) GetContextOk() (*TelemetryDruidQueryContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TelemetryDruidGroupByRequestAllOf) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given TelemetryDruidQueryContext and assigns it to the Context field.
func (o *TelemetryDruidGroupByRequestAllOf) SetContext(v TelemetryDruidQueryContext) {
	o.Context = &v
}

func (o TelemetryDruidGroupByRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dataSource"] = o.DataSource
	}
	if true {
		toSerialize["dimensions"] = o.Dimensions
	}
	if o.LimitSpec != nil {
		toSerialize["limitSpec"] = o.LimitSpec
	}
	if o.Having != nil {
		toSerialize["having"] = o.Having
	}
	if true {
		toSerialize["granularity"] = o.Granularity
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Aggregations != nil {
		toSerialize["aggregations"] = o.Aggregations
	}
	if o.PostAggregations != nil {
		toSerialize["postAggregations"] = o.PostAggregations
	}
	if true {
		toSerialize["intervals"] = o.Intervals
	}
	if o.SubtotalsSpec != nil {
		toSerialize["subtotalsSpec"] = o.SubtotalsSpec
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidGroupByRequestAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidGroupByRequestAllOf := _TelemetryDruidGroupByRequestAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidGroupByRequestAllOf); err == nil {
		*o = TelemetryDruidGroupByRequestAllOf(varTelemetryDruidGroupByRequestAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dataSource")
		delete(additionalProperties, "dimensions")
		delete(additionalProperties, "limitSpec")
		delete(additionalProperties, "having")
		delete(additionalProperties, "granularity")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "aggregations")
		delete(additionalProperties, "postAggregations")
		delete(additionalProperties, "intervals")
		delete(additionalProperties, "subtotalsSpec")
		delete(additionalProperties, "context")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidGroupByRequestAllOf struct {
	value *TelemetryDruidGroupByRequestAllOf
	isSet bool
}

func (v NullableTelemetryDruidGroupByRequestAllOf) Get() *TelemetryDruidGroupByRequestAllOf {
	return v.value
}

func (v *NullableTelemetryDruidGroupByRequestAllOf) Set(val *TelemetryDruidGroupByRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidGroupByRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidGroupByRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidGroupByRequestAllOf(val *TelemetryDruidGroupByRequestAllOf) *NullableTelemetryDruidGroupByRequestAllOf {
	return &NullableTelemetryDruidGroupByRequestAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidGroupByRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidGroupByRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
