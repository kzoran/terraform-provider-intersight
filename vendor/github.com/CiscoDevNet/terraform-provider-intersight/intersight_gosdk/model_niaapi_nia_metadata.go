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
	"reflect"
	"strings"
	"time"
)

// NiaapiNiaMetadata Contains the latest Metadata available for download from server.
type NiaapiNiaMetadata struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string         `json:"ObjectType"`
	Content    []NiaapiDetail `json:"Content,omitempty"`
	// The date when this package is generated.
	Date *time.Time `json:"Date,omitempty"`
	// Chksum used to check the integrity of the Metadata file downloaded.
	MetadataChksum *string `json:"MetadataChksum,omitempty"`
	// The Filename of this Metadata package.
	MetadataFilename *string `json:"MetadataFilename,omitempty"`
	// The version number of the Metadata package.
	Version              *int64 `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiNiaMetadata NiaapiNiaMetadata

// NewNiaapiNiaMetadata instantiates a new NiaapiNiaMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiNiaMetadata(classId string, objectType string) *NiaapiNiaMetadata {
	this := NiaapiNiaMetadata{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiNiaMetadataWithDefaults instantiates a new NiaapiNiaMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiNiaMetadataWithDefaults() *NiaapiNiaMetadata {
	this := NiaapiNiaMetadata{}
	var classId string = "niaapi.NiaMetadata"
	this.ClassId = classId
	var objectType string = "niaapi.NiaMetadata"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiaapiNiaMetadata) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiaapiNiaMetadata) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiaapiNiaMetadata) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiaapiNiaMetadata) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiaapiNiaMetadata) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiaapiNiaMetadata) SetObjectType(v string) {
	o.ObjectType = v
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiaapiNiaMetadata) GetContent() []NiaapiDetail {
	if o == nil {
		var ret []NiaapiDetail
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiaapiNiaMetadata) GetContentOk() (*[]NiaapiDetail, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return &o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *NiaapiNiaMetadata) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []NiaapiDetail and assigns it to the Content field.
func (o *NiaapiNiaMetadata) SetContent(v []NiaapiDetail) {
	o.Content = v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *NiaapiNiaMetadata) GetDate() time.Time {
	if o == nil || o.Date == nil {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNiaMetadata) GetDateOk() (*time.Time, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *NiaapiNiaMetadata) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *NiaapiNiaMetadata) SetDate(v time.Time) {
	o.Date = &v
}

// GetMetadataChksum returns the MetadataChksum field value if set, zero value otherwise.
func (o *NiaapiNiaMetadata) GetMetadataChksum() string {
	if o == nil || o.MetadataChksum == nil {
		var ret string
		return ret
	}
	return *o.MetadataChksum
}

// GetMetadataChksumOk returns a tuple with the MetadataChksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNiaMetadata) GetMetadataChksumOk() (*string, bool) {
	if o == nil || o.MetadataChksum == nil {
		return nil, false
	}
	return o.MetadataChksum, true
}

// HasMetadataChksum returns a boolean if a field has been set.
func (o *NiaapiNiaMetadata) HasMetadataChksum() bool {
	if o != nil && o.MetadataChksum != nil {
		return true
	}

	return false
}

// SetMetadataChksum gets a reference to the given string and assigns it to the MetadataChksum field.
func (o *NiaapiNiaMetadata) SetMetadataChksum(v string) {
	o.MetadataChksum = &v
}

// GetMetadataFilename returns the MetadataFilename field value if set, zero value otherwise.
func (o *NiaapiNiaMetadata) GetMetadataFilename() string {
	if o == nil || o.MetadataFilename == nil {
		var ret string
		return ret
	}
	return *o.MetadataFilename
}

// GetMetadataFilenameOk returns a tuple with the MetadataFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNiaMetadata) GetMetadataFilenameOk() (*string, bool) {
	if o == nil || o.MetadataFilename == nil {
		return nil, false
	}
	return o.MetadataFilename, true
}

// HasMetadataFilename returns a boolean if a field has been set.
func (o *NiaapiNiaMetadata) HasMetadataFilename() bool {
	if o != nil && o.MetadataFilename != nil {
		return true
	}

	return false
}

// SetMetadataFilename gets a reference to the given string and assigns it to the MetadataFilename field.
func (o *NiaapiNiaMetadata) SetMetadataFilename(v string) {
	o.MetadataFilename = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NiaapiNiaMetadata) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNiaMetadata) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NiaapiNiaMetadata) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *NiaapiNiaMetadata) SetVersion(v int64) {
	o.Version = &v
}

func (o NiaapiNiaMetadata) MarshalJSON() ([]byte, error) {
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
	if o.Content != nil {
		toSerialize["Content"] = o.Content
	}
	if o.Date != nil {
		toSerialize["Date"] = o.Date
	}
	if o.MetadataChksum != nil {
		toSerialize["MetadataChksum"] = o.MetadataChksum
	}
	if o.MetadataFilename != nil {
		toSerialize["MetadataFilename"] = o.MetadataFilename
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiNiaMetadata) UnmarshalJSON(bytes []byte) (err error) {
	type NiaapiNiaMetadataWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string         `json:"ObjectType"`
		Content    []NiaapiDetail `json:"Content,omitempty"`
		// The date when this package is generated.
		Date *time.Time `json:"Date,omitempty"`
		// Chksum used to check the integrity of the Metadata file downloaded.
		MetadataChksum *string `json:"MetadataChksum,omitempty"`
		// The Filename of this Metadata package.
		MetadataFilename *string `json:"MetadataFilename,omitempty"`
		// The version number of the Metadata package.
		Version *int64 `json:"Version,omitempty"`
	}

	varNiaapiNiaMetadataWithoutEmbeddedStruct := NiaapiNiaMetadataWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiaapiNiaMetadataWithoutEmbeddedStruct)
	if err == nil {
		varNiaapiNiaMetadata := _NiaapiNiaMetadata{}
		varNiaapiNiaMetadata.ClassId = varNiaapiNiaMetadataWithoutEmbeddedStruct.ClassId
		varNiaapiNiaMetadata.ObjectType = varNiaapiNiaMetadataWithoutEmbeddedStruct.ObjectType
		varNiaapiNiaMetadata.Content = varNiaapiNiaMetadataWithoutEmbeddedStruct.Content
		varNiaapiNiaMetadata.Date = varNiaapiNiaMetadataWithoutEmbeddedStruct.Date
		varNiaapiNiaMetadata.MetadataChksum = varNiaapiNiaMetadataWithoutEmbeddedStruct.MetadataChksum
		varNiaapiNiaMetadata.MetadataFilename = varNiaapiNiaMetadataWithoutEmbeddedStruct.MetadataFilename
		varNiaapiNiaMetadata.Version = varNiaapiNiaMetadataWithoutEmbeddedStruct.Version
		*o = NiaapiNiaMetadata(varNiaapiNiaMetadata)
	} else {
		return err
	}

	varNiaapiNiaMetadata := _NiaapiNiaMetadata{}

	err = json.Unmarshal(bytes, &varNiaapiNiaMetadata)
	if err == nil {
		o.MoBaseMo = varNiaapiNiaMetadata.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Content")
		delete(additionalProperties, "Date")
		delete(additionalProperties, "MetadataChksum")
		delete(additionalProperties, "MetadataFilename")
		delete(additionalProperties, "Version")

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

type NullableNiaapiNiaMetadata struct {
	value *NiaapiNiaMetadata
	isSet bool
}

func (v NullableNiaapiNiaMetadata) Get() *NiaapiNiaMetadata {
	return v.value
}

func (v *NullableNiaapiNiaMetadata) Set(val *NiaapiNiaMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiNiaMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiNiaMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiNiaMetadata(val *NiaapiNiaMetadata) *NullableNiaapiNiaMetadata {
	return &NullableNiaapiNiaMetadata{value: val, isSet: true}
}

func (v NullableNiaapiNiaMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiNiaMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
