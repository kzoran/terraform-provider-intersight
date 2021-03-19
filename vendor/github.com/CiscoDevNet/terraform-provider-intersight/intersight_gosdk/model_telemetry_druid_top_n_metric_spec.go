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
	"fmt"
)

// TelemetryDruidTopNMetricSpec - A Druid TopN metric spec.
type TelemetryDruidTopNMetricSpec struct {
	TelemetryDruidDimensionTopNMetricSpec *TelemetryDruidDimensionTopNMetricSpec
	TelemetryDruidInvertedTopNMetricSpec  *TelemetryDruidInvertedTopNMetricSpec
	TelemetryDruidNumericTopNMetricSpec   *TelemetryDruidNumericTopNMetricSpec
}

// TelemetryDruidDimensionTopNMetricSpecAsTelemetryDruidTopNMetricSpec is a convenience function that returns TelemetryDruidDimensionTopNMetricSpec wrapped in TelemetryDruidTopNMetricSpec
func TelemetryDruidDimensionTopNMetricSpecAsTelemetryDruidTopNMetricSpec(v *TelemetryDruidDimensionTopNMetricSpec) TelemetryDruidTopNMetricSpec {
	return TelemetryDruidTopNMetricSpec{TelemetryDruidDimensionTopNMetricSpec: v}
}

// TelemetryDruidInvertedTopNMetricSpecAsTelemetryDruidTopNMetricSpec is a convenience function that returns TelemetryDruidInvertedTopNMetricSpec wrapped in TelemetryDruidTopNMetricSpec
func TelemetryDruidInvertedTopNMetricSpecAsTelemetryDruidTopNMetricSpec(v *TelemetryDruidInvertedTopNMetricSpec) TelemetryDruidTopNMetricSpec {
	return TelemetryDruidTopNMetricSpec{TelemetryDruidInvertedTopNMetricSpec: v}
}

// TelemetryDruidNumericTopNMetricSpecAsTelemetryDruidTopNMetricSpec is a convenience function that returns TelemetryDruidNumericTopNMetricSpec wrapped in TelemetryDruidTopNMetricSpec
func TelemetryDruidNumericTopNMetricSpecAsTelemetryDruidTopNMetricSpec(v *TelemetryDruidNumericTopNMetricSpec) TelemetryDruidTopNMetricSpec {
	return TelemetryDruidTopNMetricSpec{TelemetryDruidNumericTopNMetricSpec: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TelemetryDruidTopNMetricSpec) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'dimension'
	if jsonDict["type"] == "dimension" {
		// try to unmarshal JSON data into TelemetryDruidDimensionTopNMetricSpec
		err = json.Unmarshal(data, &dst.TelemetryDruidDimensionTopNMetricSpec)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidDimensionTopNMetricSpec, return on the first match
		} else {
			dst.TelemetryDruidDimensionTopNMetricSpec = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidTopNMetricSpec as TelemetryDruidDimensionTopNMetricSpec: %s", err.Error())
		}
	}

	// check if the discriminator value is 'inverted'
	if jsonDict["type"] == "inverted" {
		// try to unmarshal JSON data into TelemetryDruidInvertedTopNMetricSpec
		err = json.Unmarshal(data, &dst.TelemetryDruidInvertedTopNMetricSpec)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidInvertedTopNMetricSpec, return on the first match
		} else {
			dst.TelemetryDruidInvertedTopNMetricSpec = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidTopNMetricSpec as TelemetryDruidInvertedTopNMetricSpec: %s", err.Error())
		}
	}

	// check if the discriminator value is 'numeric'
	if jsonDict["type"] == "numeric" {
		// try to unmarshal JSON data into TelemetryDruidNumericTopNMetricSpec
		err = json.Unmarshal(data, &dst.TelemetryDruidNumericTopNMetricSpec)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidNumericTopNMetricSpec, return on the first match
		} else {
			dst.TelemetryDruidNumericTopNMetricSpec = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidTopNMetricSpec as TelemetryDruidNumericTopNMetricSpec: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidDimensionTopNMetricSpec'
	if jsonDict["type"] == "telemetry.DruidDimensionTopNMetricSpec" {
		// try to unmarshal JSON data into TelemetryDruidDimensionTopNMetricSpec
		err = json.Unmarshal(data, &dst.TelemetryDruidDimensionTopNMetricSpec)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidDimensionTopNMetricSpec, return on the first match
		} else {
			dst.TelemetryDruidDimensionTopNMetricSpec = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidTopNMetricSpec as TelemetryDruidDimensionTopNMetricSpec: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidInvertedTopNMetricSpec'
	if jsonDict["type"] == "telemetry.DruidInvertedTopNMetricSpec" {
		// try to unmarshal JSON data into TelemetryDruidInvertedTopNMetricSpec
		err = json.Unmarshal(data, &dst.TelemetryDruidInvertedTopNMetricSpec)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidInvertedTopNMetricSpec, return on the first match
		} else {
			dst.TelemetryDruidInvertedTopNMetricSpec = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidTopNMetricSpec as TelemetryDruidInvertedTopNMetricSpec: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidNumericTopNMetricSpec'
	if jsonDict["type"] == "telemetry.DruidNumericTopNMetricSpec" {
		// try to unmarshal JSON data into TelemetryDruidNumericTopNMetricSpec
		err = json.Unmarshal(data, &dst.TelemetryDruidNumericTopNMetricSpec)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidNumericTopNMetricSpec, return on the first match
		} else {
			dst.TelemetryDruidNumericTopNMetricSpec = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidTopNMetricSpec as TelemetryDruidNumericTopNMetricSpec: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TelemetryDruidTopNMetricSpec) MarshalJSON() ([]byte, error) {
	if src.TelemetryDruidDimensionTopNMetricSpec != nil {
		return json.Marshal(&src.TelemetryDruidDimensionTopNMetricSpec)
	}

	if src.TelemetryDruidInvertedTopNMetricSpec != nil {
		return json.Marshal(&src.TelemetryDruidInvertedTopNMetricSpec)
	}

	if src.TelemetryDruidNumericTopNMetricSpec != nil {
		return json.Marshal(&src.TelemetryDruidNumericTopNMetricSpec)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TelemetryDruidTopNMetricSpec) GetActualInstance() interface{} {
	if obj.TelemetryDruidDimensionTopNMetricSpec != nil {
		return obj.TelemetryDruidDimensionTopNMetricSpec
	}

	if obj.TelemetryDruidInvertedTopNMetricSpec != nil {
		return obj.TelemetryDruidInvertedTopNMetricSpec
	}

	if obj.TelemetryDruidNumericTopNMetricSpec != nil {
		return obj.TelemetryDruidNumericTopNMetricSpec
	}

	// all schemas are nil
	return nil
}

type NullableTelemetryDruidTopNMetricSpec struct {
	value *TelemetryDruidTopNMetricSpec
	isSet bool
}

func (v NullableTelemetryDruidTopNMetricSpec) Get() *TelemetryDruidTopNMetricSpec {
	return v.value
}

func (v *NullableTelemetryDruidTopNMetricSpec) Set(val *TelemetryDruidTopNMetricSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidTopNMetricSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidTopNMetricSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidTopNMetricSpec(val *TelemetryDruidTopNMetricSpec) *NullableTelemetryDruidTopNMetricSpec {
	return &NullableTelemetryDruidTopNMetricSpec{value: val, isSet: true}
}

func (v NullableTelemetryDruidTopNMetricSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidTopNMetricSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
