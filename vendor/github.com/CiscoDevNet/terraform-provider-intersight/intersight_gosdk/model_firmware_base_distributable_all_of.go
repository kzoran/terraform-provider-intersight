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

// FirmwareBaseDistributableAllOf Definition of the list of properties defined in 'firmware.BaseDistributable', excluding properties defined in parent classes.
type FirmwareBaseDistributableAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The bundle type of the image, as published on cisco.com.
	BundleType    *string                 `json:"BundleType,omitempty"`
	ComponentMeta []FirmwareComponentMeta `json:"ComponentMeta,omitempty"`
	// The unique identifier for an image in a Cisco repository.
	Guid *string `json:"Guid,omitempty"`
	// The type of image which the distributable falls into according to the component it can upgrade. For e.g.; Standalone server, Intersight managed server, UCS Managed Fabric Interconnect. The field is used in private appliance mode, where image does not have description populated from CCO.
	ImageType *string `json:"ImageType,omitempty"`
	// The mdfid of the image provided by cisco.com.
	Mdfid *string `json:"Mdfid,omitempty"`
	// The endpoint model for which this firmware image is applicable.
	Model *string `json:"Model,omitempty"`
	// The platform type of the image.
	PlatformType *string `json:"PlatformType,omitempty"`
	// The build which is recommended by Cisco.
	RecommendedBuild *string `json:"RecommendedBuild,omitempty"`
	// The url for the release notes of this image.
	ReleaseNotesUrl *string `json:"ReleaseNotesUrl,omitempty"`
	// The software type id provided by cisco.com.
	SoftwareTypeId  *string  `json:"SoftwareTypeId,omitempty"`
	SupportedModels []string `json:"SupportedModels,omitempty"`
	// The vendor or publisher of this file.
	Vendor *string `json:"Vendor,omitempty"`
	// An array of relationships to firmwareDistributableMeta resources.
	DistributableMetas   []FirmwareDistributableMetaRelationship `json:"DistributableMetas,omitempty"`
	Release              *SoftwarerepositoryReleaseRelationship  `json:"Release,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareBaseDistributableAllOf FirmwareBaseDistributableAllOf

// NewFirmwareBaseDistributableAllOf instantiates a new FirmwareBaseDistributableAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareBaseDistributableAllOf(classId string, objectType string) *FirmwareBaseDistributableAllOf {
	this := FirmwareBaseDistributableAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var vendor string = "Cisco"
	this.Vendor = &vendor
	return &this
}

// NewFirmwareBaseDistributableAllOfWithDefaults instantiates a new FirmwareBaseDistributableAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareBaseDistributableAllOfWithDefaults() *FirmwareBaseDistributableAllOf {
	this := FirmwareBaseDistributableAllOf{}
	var vendor string = "Cisco"
	this.Vendor = &vendor
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareBaseDistributableAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareBaseDistributableAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareBaseDistributableAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareBaseDistributableAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareBaseDistributableAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareBaseDistributableAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBundleType returns the BundleType field value if set, zero value otherwise.
func (o *FirmwareBaseDistributableAllOf) GetBundleType() string {
	if o == nil || o.BundleType == nil {
		var ret string
		return ret
	}
	return *o.BundleType
}

// GetBundleTypeOk returns a tuple with the BundleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBaseDistributableAllOf) GetBundleTypeOk() (*string, bool) {
	if o == nil || o.BundleType == nil {
		return nil, false
	}
	return o.BundleType, true
}

// HasBundleType returns a boolean if a field has been set.
func (o *FirmwareBaseDistributableAllOf) HasBundleType() bool {
	if o != nil && o.BundleType != nil {
		return true
	}

	return false
}

// SetBundleType gets a reference to the given string and assigns it to the BundleType field.
func (o *FirmwareBaseDistributableAllOf) SetBundleType(v string) {
	o.BundleType = &v
}

// GetComponentMeta returns the ComponentMeta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareBaseDistributableAllOf) GetComponentMeta() []FirmwareComponentMeta {
	if o == nil {
		var ret []FirmwareComponentMeta
		return ret
	}
	return o.ComponentMeta
}

// GetComponentMetaOk returns a tuple with the ComponentMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareBaseDistributableAllOf) GetComponentMetaOk() (*[]FirmwareComponentMeta, bool) {
	if o == nil || o.ComponentMeta == nil {
		return nil, false
	}
	return &o.ComponentMeta, true
}

// HasComponentMeta returns a boolean if a field has been set.
func (o *FirmwareBaseDistributableAllOf) HasComponentMeta() bool {
	if o != nil && o.ComponentMeta != nil {
		return true
	}

	return false
}

// SetComponentMeta gets a reference to the given []FirmwareComponentMeta and assigns it to the ComponentMeta field.
func (o *FirmwareBaseDistributableAllOf) SetComponentMeta(v []FirmwareComponentMeta) {
	o.ComponentMeta = v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *FirmwareBaseDistributableAllOf) GetGuid() string {
	if o == nil || o.Guid == nil {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBaseDistributableAllOf) GetGuidOk() (*string, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *FirmwareBaseDistributableAllOf) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *FirmwareBaseDistributableAllOf) SetGuid(v string) {
	o.Guid = &v
}

// GetImageType returns the ImageType field value if set, zero value otherwise.
func (o *FirmwareBaseDistributableAllOf) GetImageType() string {
	if o == nil || o.ImageType == nil {
		var ret string
		return ret
	}
	return *o.ImageType
}

// GetImageTypeOk returns a tuple with the ImageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBaseDistributableAllOf) GetImageTypeOk() (*string, bool) {
	if o == nil || o.ImageType == nil {
		return nil, false
	}
	return o.ImageType, true
}

// HasImageType returns a boolean if a field has been set.
func (o *FirmwareBaseDistributableAllOf) HasImageType() bool {
	if o != nil && o.ImageType != nil {
		return true
	}

	return false
}

// SetImageType gets a reference to the given string and assigns it to the ImageType field.
func (o *FirmwareBaseDistributableAllOf) SetImageType(v string) {
	o.ImageType = &v
}

// GetMdfid returns the Mdfid field value if set, zero value otherwise.
func (o *FirmwareBaseDistributableAllOf) GetMdfid() string {
	if o == nil || o.Mdfid == nil {
		var ret string
		return ret
	}
	return *o.Mdfid
}

// GetMdfidOk returns a tuple with the Mdfid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBaseDistributableAllOf) GetMdfidOk() (*string, bool) {
	if o == nil || o.Mdfid == nil {
		return nil, false
	}
	return o.Mdfid, true
}

// HasMdfid returns a boolean if a field has been set.
func (o *FirmwareBaseDistributableAllOf) HasMdfid() bool {
	if o != nil && o.Mdfid != nil {
		return true
	}

	return false
}

// SetMdfid gets a reference to the given string and assigns it to the Mdfid field.
func (o *FirmwareBaseDistributableAllOf) SetMdfid(v string) {
	o.Mdfid = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *FirmwareBaseDistributableAllOf) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBaseDistributableAllOf) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *FirmwareBaseDistributableAllOf) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *FirmwareBaseDistributableAllOf) SetModel(v string) {
	o.Model = &v
}

// GetPlatformType returns the PlatformType field value if set, zero value otherwise.
func (o *FirmwareBaseDistributableAllOf) GetPlatformType() string {
	if o == nil || o.PlatformType == nil {
		var ret string
		return ret
	}
	return *o.PlatformType
}

// GetPlatformTypeOk returns a tuple with the PlatformType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBaseDistributableAllOf) GetPlatformTypeOk() (*string, bool) {
	if o == nil || o.PlatformType == nil {
		return nil, false
	}
	return o.PlatformType, true
}

// HasPlatformType returns a boolean if a field has been set.
func (o *FirmwareBaseDistributableAllOf) HasPlatformType() bool {
	if o != nil && o.PlatformType != nil {
		return true
	}

	return false
}

// SetPlatformType gets a reference to the given string and assigns it to the PlatformType field.
func (o *FirmwareBaseDistributableAllOf) SetPlatformType(v string) {
	o.PlatformType = &v
}

// GetRecommendedBuild returns the RecommendedBuild field value if set, zero value otherwise.
func (o *FirmwareBaseDistributableAllOf) GetRecommendedBuild() string {
	if o == nil || o.RecommendedBuild == nil {
		var ret string
		return ret
	}
	return *o.RecommendedBuild
}

// GetRecommendedBuildOk returns a tuple with the RecommendedBuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBaseDistributableAllOf) GetRecommendedBuildOk() (*string, bool) {
	if o == nil || o.RecommendedBuild == nil {
		return nil, false
	}
	return o.RecommendedBuild, true
}

// HasRecommendedBuild returns a boolean if a field has been set.
func (o *FirmwareBaseDistributableAllOf) HasRecommendedBuild() bool {
	if o != nil && o.RecommendedBuild != nil {
		return true
	}

	return false
}

// SetRecommendedBuild gets a reference to the given string and assigns it to the RecommendedBuild field.
func (o *FirmwareBaseDistributableAllOf) SetRecommendedBuild(v string) {
	o.RecommendedBuild = &v
}

// GetReleaseNotesUrl returns the ReleaseNotesUrl field value if set, zero value otherwise.
func (o *FirmwareBaseDistributableAllOf) GetReleaseNotesUrl() string {
	if o == nil || o.ReleaseNotesUrl == nil {
		var ret string
		return ret
	}
	return *o.ReleaseNotesUrl
}

// GetReleaseNotesUrlOk returns a tuple with the ReleaseNotesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBaseDistributableAllOf) GetReleaseNotesUrlOk() (*string, bool) {
	if o == nil || o.ReleaseNotesUrl == nil {
		return nil, false
	}
	return o.ReleaseNotesUrl, true
}

// HasReleaseNotesUrl returns a boolean if a field has been set.
func (o *FirmwareBaseDistributableAllOf) HasReleaseNotesUrl() bool {
	if o != nil && o.ReleaseNotesUrl != nil {
		return true
	}

	return false
}

// SetReleaseNotesUrl gets a reference to the given string and assigns it to the ReleaseNotesUrl field.
func (o *FirmwareBaseDistributableAllOf) SetReleaseNotesUrl(v string) {
	o.ReleaseNotesUrl = &v
}

// GetSoftwareTypeId returns the SoftwareTypeId field value if set, zero value otherwise.
func (o *FirmwareBaseDistributableAllOf) GetSoftwareTypeId() string {
	if o == nil || o.SoftwareTypeId == nil {
		var ret string
		return ret
	}
	return *o.SoftwareTypeId
}

// GetSoftwareTypeIdOk returns a tuple with the SoftwareTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBaseDistributableAllOf) GetSoftwareTypeIdOk() (*string, bool) {
	if o == nil || o.SoftwareTypeId == nil {
		return nil, false
	}
	return o.SoftwareTypeId, true
}

// HasSoftwareTypeId returns a boolean if a field has been set.
func (o *FirmwareBaseDistributableAllOf) HasSoftwareTypeId() bool {
	if o != nil && o.SoftwareTypeId != nil {
		return true
	}

	return false
}

// SetSoftwareTypeId gets a reference to the given string and assigns it to the SoftwareTypeId field.
func (o *FirmwareBaseDistributableAllOf) SetSoftwareTypeId(v string) {
	o.SoftwareTypeId = &v
}

// GetSupportedModels returns the SupportedModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareBaseDistributableAllOf) GetSupportedModels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedModels
}

// GetSupportedModelsOk returns a tuple with the SupportedModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareBaseDistributableAllOf) GetSupportedModelsOk() (*[]string, bool) {
	if o == nil || o.SupportedModels == nil {
		return nil, false
	}
	return &o.SupportedModels, true
}

// HasSupportedModels returns a boolean if a field has been set.
func (o *FirmwareBaseDistributableAllOf) HasSupportedModels() bool {
	if o != nil && o.SupportedModels != nil {
		return true
	}

	return false
}

// SetSupportedModels gets a reference to the given []string and assigns it to the SupportedModels field.
func (o *FirmwareBaseDistributableAllOf) SetSupportedModels(v []string) {
	o.SupportedModels = v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *FirmwareBaseDistributableAllOf) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBaseDistributableAllOf) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *FirmwareBaseDistributableAllOf) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *FirmwareBaseDistributableAllOf) SetVendor(v string) {
	o.Vendor = &v
}

// GetDistributableMetas returns the DistributableMetas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareBaseDistributableAllOf) GetDistributableMetas() []FirmwareDistributableMetaRelationship {
	if o == nil {
		var ret []FirmwareDistributableMetaRelationship
		return ret
	}
	return o.DistributableMetas
}

// GetDistributableMetasOk returns a tuple with the DistributableMetas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareBaseDistributableAllOf) GetDistributableMetasOk() (*[]FirmwareDistributableMetaRelationship, bool) {
	if o == nil || o.DistributableMetas == nil {
		return nil, false
	}
	return &o.DistributableMetas, true
}

// HasDistributableMetas returns a boolean if a field has been set.
func (o *FirmwareBaseDistributableAllOf) HasDistributableMetas() bool {
	if o != nil && o.DistributableMetas != nil {
		return true
	}

	return false
}

// SetDistributableMetas gets a reference to the given []FirmwareDistributableMetaRelationship and assigns it to the DistributableMetas field.
func (o *FirmwareBaseDistributableAllOf) SetDistributableMetas(v []FirmwareDistributableMetaRelationship) {
	o.DistributableMetas = v
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *FirmwareBaseDistributableAllOf) GetRelease() SoftwarerepositoryReleaseRelationship {
	if o == nil || o.Release == nil {
		var ret SoftwarerepositoryReleaseRelationship
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBaseDistributableAllOf) GetReleaseOk() (*SoftwarerepositoryReleaseRelationship, bool) {
	if o == nil || o.Release == nil {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *FirmwareBaseDistributableAllOf) HasRelease() bool {
	if o != nil && o.Release != nil {
		return true
	}

	return false
}

// SetRelease gets a reference to the given SoftwarerepositoryReleaseRelationship and assigns it to the Release field.
func (o *FirmwareBaseDistributableAllOf) SetRelease(v SoftwarerepositoryReleaseRelationship) {
	o.Release = &v
}

func (o FirmwareBaseDistributableAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BundleType != nil {
		toSerialize["BundleType"] = o.BundleType
	}
	if o.ComponentMeta != nil {
		toSerialize["ComponentMeta"] = o.ComponentMeta
	}
	if o.Guid != nil {
		toSerialize["Guid"] = o.Guid
	}
	if o.ImageType != nil {
		toSerialize["ImageType"] = o.ImageType
	}
	if o.Mdfid != nil {
		toSerialize["Mdfid"] = o.Mdfid
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.PlatformType != nil {
		toSerialize["PlatformType"] = o.PlatformType
	}
	if o.RecommendedBuild != nil {
		toSerialize["RecommendedBuild"] = o.RecommendedBuild
	}
	if o.ReleaseNotesUrl != nil {
		toSerialize["ReleaseNotesUrl"] = o.ReleaseNotesUrl
	}
	if o.SoftwareTypeId != nil {
		toSerialize["SoftwareTypeId"] = o.SoftwareTypeId
	}
	if o.SupportedModels != nil {
		toSerialize["SupportedModels"] = o.SupportedModels
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}
	if o.DistributableMetas != nil {
		toSerialize["DistributableMetas"] = o.DistributableMetas
	}
	if o.Release != nil {
		toSerialize["Release"] = o.Release
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareBaseDistributableAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwareBaseDistributableAllOf := _FirmwareBaseDistributableAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwareBaseDistributableAllOf); err == nil {
		*o = FirmwareBaseDistributableAllOf(varFirmwareBaseDistributableAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BundleType")
		delete(additionalProperties, "ComponentMeta")
		delete(additionalProperties, "Guid")
		delete(additionalProperties, "ImageType")
		delete(additionalProperties, "Mdfid")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "PlatformType")
		delete(additionalProperties, "RecommendedBuild")
		delete(additionalProperties, "ReleaseNotesUrl")
		delete(additionalProperties, "SoftwareTypeId")
		delete(additionalProperties, "SupportedModels")
		delete(additionalProperties, "Vendor")
		delete(additionalProperties, "DistributableMetas")
		delete(additionalProperties, "Release")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareBaseDistributableAllOf struct {
	value *FirmwareBaseDistributableAllOf
	isSet bool
}

func (v NullableFirmwareBaseDistributableAllOf) Get() *FirmwareBaseDistributableAllOf {
	return v.value
}

func (v *NullableFirmwareBaseDistributableAllOf) Set(val *FirmwareBaseDistributableAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareBaseDistributableAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareBaseDistributableAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareBaseDistributableAllOf(val *FirmwareBaseDistributableAllOf) *NullableFirmwareBaseDistributableAllOf {
	return &NullableFirmwareBaseDistributableAllOf{value: val, isSet: true}
}

func (v NullableFirmwareBaseDistributableAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareBaseDistributableAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
