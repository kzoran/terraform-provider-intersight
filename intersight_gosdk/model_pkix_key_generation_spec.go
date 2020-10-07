/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-31T04:35:53Z.
 *
 * API version: 1.0.9-2110
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// PkixKeyGenerationSpec The key generation spec provides the algorithm and the parameters required for this algorithm to generate a key.
type PkixKeyGenerationSpec struct {
	MoBaseComplexType
	// Name of the key generation algorithm. * `RSA` - Key pairs should be generated by the RSA algorithm.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PkixKeyGenerationSpec PkixKeyGenerationSpec

// NewPkixKeyGenerationSpec instantiates a new PkixKeyGenerationSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkixKeyGenerationSpec() *PkixKeyGenerationSpec {
	this := PkixKeyGenerationSpec{}
	var name string = "RSA"
	this.Name = &name
	return &this
}

// NewPkixKeyGenerationSpecWithDefaults instantiates a new PkixKeyGenerationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkixKeyGenerationSpecWithDefaults() *PkixKeyGenerationSpec {
	this := PkixKeyGenerationSpec{}
	var name string = "RSA"
	this.Name = &name
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PkixKeyGenerationSpec) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkixKeyGenerationSpec) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PkixKeyGenerationSpec) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PkixKeyGenerationSpec) SetName(v string) {
	o.Name = &v
}

func (o PkixKeyGenerationSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PkixKeyGenerationSpec) UnmarshalJSON(bytes []byte) (err error) {
	type PkixKeyGenerationSpecWithoutEmbeddedStruct struct {
		// Name of the key generation algorithm. * `RSA` - Key pairs should be generated by the RSA algorithm.
		Name *string `json:"Name,omitempty"`
	}

	varPkixKeyGenerationSpecWithoutEmbeddedStruct := PkixKeyGenerationSpecWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPkixKeyGenerationSpecWithoutEmbeddedStruct)
	if err == nil {
		varPkixKeyGenerationSpec := _PkixKeyGenerationSpec{}
		varPkixKeyGenerationSpec.Name = varPkixKeyGenerationSpecWithoutEmbeddedStruct.Name
		*o = PkixKeyGenerationSpec(varPkixKeyGenerationSpec)
	} else {
		return err
	}

	varPkixKeyGenerationSpec := _PkixKeyGenerationSpec{}

	err = json.Unmarshal(bytes, &varPkixKeyGenerationSpec)
	if err == nil {
		o.MoBaseComplexType = varPkixKeyGenerationSpec.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Name")

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

type NullablePkixKeyGenerationSpec struct {
	value *PkixKeyGenerationSpec
	isSet bool
}

func (v NullablePkixKeyGenerationSpec) Get() *PkixKeyGenerationSpec {
	return v.value
}

func (v *NullablePkixKeyGenerationSpec) Set(val *PkixKeyGenerationSpec) {
	v.value = val
	v.isSet = true
}

func (v NullablePkixKeyGenerationSpec) IsSet() bool {
	return v.isSet
}

func (v *NullablePkixKeyGenerationSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkixKeyGenerationSpec(val *PkixKeyGenerationSpec) *NullablePkixKeyGenerationSpec {
	return &NullablePkixKeyGenerationSpec{value: val, isSet: true}
}

func (v NullablePkixKeyGenerationSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkixKeyGenerationSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}