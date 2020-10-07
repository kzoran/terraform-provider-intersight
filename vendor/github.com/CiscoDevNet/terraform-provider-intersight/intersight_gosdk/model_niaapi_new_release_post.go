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
	"time"
)

// NiaapiNewReleasePost This contains the new release notice.
type NiaapiNewReleasePost struct {
	MoBaseMo
	// The date when this new release notice is posted.
	PostDate   *time.Time              `json:"PostDate,omitempty"`
	PostDetail *NiaapiNewReleaseDetail `json:"PostDetail,omitempty"`
	// The document type of this post.
	PostType *string `json:"PostType,omitempty"`
	// Identificator of this inbox post.
	Postid *string `json:"Postid,omitempty"`
	// Revision number of this notice.
	Revision             *string `json:"Revision,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiNewReleasePost NiaapiNewReleasePost

// NewNiaapiNewReleasePost instantiates a new NiaapiNewReleasePost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiNewReleasePost() *NiaapiNewReleasePost {
	this := NiaapiNewReleasePost{}
	return &this
}

// NewNiaapiNewReleasePostWithDefaults instantiates a new NiaapiNewReleasePost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiNewReleasePostWithDefaults() *NiaapiNewReleasePost {
	this := NiaapiNewReleasePost{}
	return &this
}

// GetPostDate returns the PostDate field value if set, zero value otherwise.
func (o *NiaapiNewReleasePost) GetPostDate() time.Time {
	if o == nil || o.PostDate == nil {
		var ret time.Time
		return ret
	}
	return *o.PostDate
}

// GetPostDateOk returns a tuple with the PostDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleasePost) GetPostDateOk() (*time.Time, bool) {
	if o == nil || o.PostDate == nil {
		return nil, false
	}
	return o.PostDate, true
}

// HasPostDate returns a boolean if a field has been set.
func (o *NiaapiNewReleasePost) HasPostDate() bool {
	if o != nil && o.PostDate != nil {
		return true
	}

	return false
}

// SetPostDate gets a reference to the given time.Time and assigns it to the PostDate field.
func (o *NiaapiNewReleasePost) SetPostDate(v time.Time) {
	o.PostDate = &v
}

// GetPostDetail returns the PostDetail field value if set, zero value otherwise.
func (o *NiaapiNewReleasePost) GetPostDetail() NiaapiNewReleaseDetail {
	if o == nil || o.PostDetail == nil {
		var ret NiaapiNewReleaseDetail
		return ret
	}
	return *o.PostDetail
}

// GetPostDetailOk returns a tuple with the PostDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleasePost) GetPostDetailOk() (*NiaapiNewReleaseDetail, bool) {
	if o == nil || o.PostDetail == nil {
		return nil, false
	}
	return o.PostDetail, true
}

// HasPostDetail returns a boolean if a field has been set.
func (o *NiaapiNewReleasePost) HasPostDetail() bool {
	if o != nil && o.PostDetail != nil {
		return true
	}

	return false
}

// SetPostDetail gets a reference to the given NiaapiNewReleaseDetail and assigns it to the PostDetail field.
func (o *NiaapiNewReleasePost) SetPostDetail(v NiaapiNewReleaseDetail) {
	o.PostDetail = &v
}

// GetPostType returns the PostType field value if set, zero value otherwise.
func (o *NiaapiNewReleasePost) GetPostType() string {
	if o == nil || o.PostType == nil {
		var ret string
		return ret
	}
	return *o.PostType
}

// GetPostTypeOk returns a tuple with the PostType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleasePost) GetPostTypeOk() (*string, bool) {
	if o == nil || o.PostType == nil {
		return nil, false
	}
	return o.PostType, true
}

// HasPostType returns a boolean if a field has been set.
func (o *NiaapiNewReleasePost) HasPostType() bool {
	if o != nil && o.PostType != nil {
		return true
	}

	return false
}

// SetPostType gets a reference to the given string and assigns it to the PostType field.
func (o *NiaapiNewReleasePost) SetPostType(v string) {
	o.PostType = &v
}

// GetPostid returns the Postid field value if set, zero value otherwise.
func (o *NiaapiNewReleasePost) GetPostid() string {
	if o == nil || o.Postid == nil {
		var ret string
		return ret
	}
	return *o.Postid
}

// GetPostidOk returns a tuple with the Postid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleasePost) GetPostidOk() (*string, bool) {
	if o == nil || o.Postid == nil {
		return nil, false
	}
	return o.Postid, true
}

// HasPostid returns a boolean if a field has been set.
func (o *NiaapiNewReleasePost) HasPostid() bool {
	if o != nil && o.Postid != nil {
		return true
	}

	return false
}

// SetPostid gets a reference to the given string and assigns it to the Postid field.
func (o *NiaapiNewReleasePost) SetPostid(v string) {
	o.Postid = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *NiaapiNewReleasePost) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleasePost) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *NiaapiNewReleasePost) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *NiaapiNewReleasePost) SetRevision(v string) {
	o.Revision = &v
}

func (o NiaapiNewReleasePost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.PostDate != nil {
		toSerialize["PostDate"] = o.PostDate
	}
	if o.PostDetail != nil {
		toSerialize["PostDetail"] = o.PostDetail
	}
	if o.PostType != nil {
		toSerialize["PostType"] = o.PostType
	}
	if o.Postid != nil {
		toSerialize["Postid"] = o.Postid
	}
	if o.Revision != nil {
		toSerialize["Revision"] = o.Revision
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiNewReleasePost) UnmarshalJSON(bytes []byte) (err error) {
	type NiaapiNewReleasePostWithoutEmbeddedStruct struct {
		// The date when this new release notice is posted.
		PostDate   *time.Time              `json:"PostDate,omitempty"`
		PostDetail *NiaapiNewReleaseDetail `json:"PostDetail,omitempty"`
		// The document type of this post.
		PostType *string `json:"PostType,omitempty"`
		// Identificator of this inbox post.
		Postid *string `json:"Postid,omitempty"`
		// Revision number of this notice.
		Revision *string `json:"Revision,omitempty"`
	}

	varNiaapiNewReleasePostWithoutEmbeddedStruct := NiaapiNewReleasePostWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiaapiNewReleasePostWithoutEmbeddedStruct)
	if err == nil {
		varNiaapiNewReleasePost := _NiaapiNewReleasePost{}
		varNiaapiNewReleasePost.PostDate = varNiaapiNewReleasePostWithoutEmbeddedStruct.PostDate
		varNiaapiNewReleasePost.PostDetail = varNiaapiNewReleasePostWithoutEmbeddedStruct.PostDetail
		varNiaapiNewReleasePost.PostType = varNiaapiNewReleasePostWithoutEmbeddedStruct.PostType
		varNiaapiNewReleasePost.Postid = varNiaapiNewReleasePostWithoutEmbeddedStruct.Postid
		varNiaapiNewReleasePost.Revision = varNiaapiNewReleasePostWithoutEmbeddedStruct.Revision
		*o = NiaapiNewReleasePost(varNiaapiNewReleasePost)
	} else {
		return err
	}

	varNiaapiNewReleasePost := _NiaapiNewReleasePost{}

	err = json.Unmarshal(bytes, &varNiaapiNewReleasePost)
	if err == nil {
		o.MoBaseMo = varNiaapiNewReleasePost.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "PostDate")
		delete(additionalProperties, "PostDetail")
		delete(additionalProperties, "PostType")
		delete(additionalProperties, "Postid")
		delete(additionalProperties, "Revision")

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

type NullableNiaapiNewReleasePost struct {
	value *NiaapiNewReleasePost
	isSet bool
}

func (v NullableNiaapiNewReleasePost) Get() *NiaapiNewReleasePost {
	return v.value
}

func (v *NullableNiaapiNewReleasePost) Set(val *NiaapiNewReleasePost) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiNewReleasePost) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiNewReleasePost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiNewReleasePost(val *NiaapiNewReleasePost) *NullableNiaapiNewReleasePost {
	return &NullableNiaapiNewReleasePost{value: val, isSet: true}
}

func (v NullableNiaapiNewReleasePost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiNewReleasePost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}