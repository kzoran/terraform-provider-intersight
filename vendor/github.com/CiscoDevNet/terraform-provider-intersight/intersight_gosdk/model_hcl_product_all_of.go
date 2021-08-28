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

// HclProductAllOf Definition of the list of properties defined in 'hcl.Product', excluding properties defined in parent classes.
type HclProductAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType  string   `json:"ObjectType"`
	DriverNames []string `json:"DriverNames,omitempty"`
	// Error code indicating the support status. * `Success` - The input combination is valid. * `Unknown` - Unknown API request to the service. * `UnknownServer` - An invalid server model is given API requests or the server model is not present in the HCL database. * `InvalidUcsVersion` - UCS Version is not in expected format. * `ProcessorNotSupported` - Processor is not supported with the given Server or the Processor does not exist in the HCL database. * `OSNotSupported` - OS version is not supported with the given server, processor combination or OS information is not present in the HCL database. * `OSUnknown` - OS vendor or version is not known as per the HCL database. * `UCSVersionNotSupported` - UCS Version is not supported with the given server, processor and OS combination or the UCS version is not present in the HCL database. * `UcsVersionServerOSCombinationNotSupported` - Combination of UCS version, server (model and processor) and os version is not supported. * `ProductUnknown` - Product is not known as per the HCL database. * `ProductNotSupported` - Product is not supported in the given UCS version, server (model and processor) and operating system version. * `DriverNameNotSupported` - Driver protocol or name is not supported for the given product. * `FirmwareVersionNotSupported` - Firmware version not supported for the component and the server, operating system combination. * `DriverVersionNotSupported` - Driver version not supported for the component and the server, operating system combination. * `FirmwareVersionDriverVersionCombinationNotSupported` - Both Firmware and Driver versions are not supported for the component and the server, operating system combination. * `FirmwareVersionAndDriverVersionNotSupported` - Firmware and Driver version combination not supported for the component and the server, operating system combination. * `FirmwareVersionAndDriverNameNotSupported` - Firmware Version and Driver name or not supported with the component and the server, operating system combination. * `InternalError` - API requests to the service have either failed or timed out. * `MarshallingError` - Error in JSON marshalling. * `Exempted` - An exempted error code means that the product is part of the exempted Catalog and should be ignored for HCL validation purposes.
	ErrorCode *string       `json:"ErrorCode,omitempty"`
	Firmwares []HclFirmware `json:"Firmwares,omitempty"`
	// Identifier of the product.
	Id *string `json:"Id,omitempty"`
	// Model/PID of the product/adapter.
	Model *string `json:"Model,omitempty"`
	// Revision of the adapter model.
	Revision *string `json:"Revision,omitempty"`
	// Type of the product/adapter say OCP, PT, GPU. * `` - Default type of the Product. * `Adapter` - Represents network adapter cards. * `StorageController` - Represents storage controllers. * `GPU` - Represents graphics cards.
	Type *string `json:"Type,omitempty"`
	// Vendor of the product or adapter.
	Vendor               *string `json:"Vendor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HclProductAllOf HclProductAllOf

// NewHclProductAllOf instantiates a new HclProductAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclProductAllOf(classId string, objectType string) *HclProductAllOf {
	this := HclProductAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = ""
	this.Type = &type_
	return &this
}

// NewHclProductAllOfWithDefaults instantiates a new HclProductAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclProductAllOfWithDefaults() *HclProductAllOf {
	this := HclProductAllOf{}
	var classId string = "hcl.Product"
	this.ClassId = classId
	var objectType string = "hcl.Product"
	this.ObjectType = objectType
	var type_ string = ""
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *HclProductAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HclProductAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HclProductAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HclProductAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HclProductAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HclProductAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDriverNames returns the DriverNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HclProductAllOf) GetDriverNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DriverNames
}

// GetDriverNamesOk returns a tuple with the DriverNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HclProductAllOf) GetDriverNamesOk() (*[]string, bool) {
	if o == nil || o.DriverNames == nil {
		return nil, false
	}
	return &o.DriverNames, true
}

// HasDriverNames returns a boolean if a field has been set.
func (o *HclProductAllOf) HasDriverNames() bool {
	if o != nil && o.DriverNames != nil {
		return true
	}

	return false
}

// SetDriverNames gets a reference to the given []string and assigns it to the DriverNames field.
func (o *HclProductAllOf) SetDriverNames(v []string) {
	o.DriverNames = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *HclProductAllOf) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclProductAllOf) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *HclProductAllOf) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *HclProductAllOf) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetFirmwares returns the Firmwares field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HclProductAllOf) GetFirmwares() []HclFirmware {
	if o == nil {
		var ret []HclFirmware
		return ret
	}
	return o.Firmwares
}

// GetFirmwaresOk returns a tuple with the Firmwares field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HclProductAllOf) GetFirmwaresOk() (*[]HclFirmware, bool) {
	if o == nil || o.Firmwares == nil {
		return nil, false
	}
	return &o.Firmwares, true
}

// HasFirmwares returns a boolean if a field has been set.
func (o *HclProductAllOf) HasFirmwares() bool {
	if o != nil && o.Firmwares != nil {
		return true
	}

	return false
}

// SetFirmwares gets a reference to the given []HclFirmware and assigns it to the Firmwares field.
func (o *HclProductAllOf) SetFirmwares(v []HclFirmware) {
	o.Firmwares = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HclProductAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclProductAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HclProductAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HclProductAllOf) SetId(v string) {
	o.Id = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *HclProductAllOf) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclProductAllOf) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *HclProductAllOf) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *HclProductAllOf) SetModel(v string) {
	o.Model = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *HclProductAllOf) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclProductAllOf) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *HclProductAllOf) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *HclProductAllOf) SetRevision(v string) {
	o.Revision = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *HclProductAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclProductAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *HclProductAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *HclProductAllOf) SetType(v string) {
	o.Type = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *HclProductAllOf) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclProductAllOf) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *HclProductAllOf) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *HclProductAllOf) SetVendor(v string) {
	o.Vendor = &v
}

func (o HclProductAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DriverNames != nil {
		toSerialize["DriverNames"] = o.DriverNames
	}
	if o.ErrorCode != nil {
		toSerialize["ErrorCode"] = o.ErrorCode
	}
	if o.Firmwares != nil {
		toSerialize["Firmwares"] = o.Firmwares
	}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.Revision != nil {
		toSerialize["Revision"] = o.Revision
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HclProductAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHclProductAllOf := _HclProductAllOf{}

	if err = json.Unmarshal(bytes, &varHclProductAllOf); err == nil {
		*o = HclProductAllOf(varHclProductAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DriverNames")
		delete(additionalProperties, "ErrorCode")
		delete(additionalProperties, "Firmwares")
		delete(additionalProperties, "Id")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "Revision")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Vendor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHclProductAllOf struct {
	value *HclProductAllOf
	isSet bool
}

func (v NullableHclProductAllOf) Get() *HclProductAllOf {
	return v.value
}

func (v *NullableHclProductAllOf) Set(val *HclProductAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHclProductAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHclProductAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclProductAllOf(val *HclProductAllOf) *NullableHclProductAllOf {
	return &NullableHclProductAllOf{value: val, isSet: true}
}

func (v NullableHclProductAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclProductAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
