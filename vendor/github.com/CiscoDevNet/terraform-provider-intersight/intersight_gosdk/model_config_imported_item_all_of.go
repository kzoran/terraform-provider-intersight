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
)

// ConfigImportedItemAllOf Definition of the list of properties defined in 'config.ImportedItem', excluding properties defined in parent classes.
type ConfigImportedItemAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Specifies whether this item MO was in shared scope or user scope when exported.
	IsShared *bool `json:"IsShared,omitempty"`
	// Specifies whether this item MO was updated or created while importing in desired service.
	IsUpdated *bool               `json:"IsUpdated,omitempty"`
	Item      NullableConfigMoRef `json:"Item,omitempty"`
	// MO item identity (the moref corresponding to item) expressed as a string.
	Name *string `json:"Name,omitempty"`
	// Moid of the MO created/updated during import for the item.
	NewMoid *string `json:"NewMoid,omitempty"`
	// Version of the service that owned the item MO when the item was exported.
	ServiceVersion *string `json:"ServiceVersion,omitempty"`
	// Status of the item's import operation. * `` - The operation has not started. * `InProgress` - The operation is in progress. * `Success` - The operation has succeeded. * `Failed` - The operation has failed. * `RollBackInitiated` - The rollback has been inititiated for import failure. * `RollBackFailed` - The rollback has failed for import failure. * `RollbackCompleted` - The rollback has completed for import failure. * `RollbackAborted` - The rollback has been aborted for import failure. * `OperationTimedOut` - The operation has timed out.
	Status *string `json:"Status,omitempty"`
	// Progress or error message for the MO's import operation.
	StatusMessage        *string                     `json:"StatusMessage,omitempty"`
	Importer             *ConfigImporterRelationship `json:"Importer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConfigImportedItemAllOf ConfigImportedItemAllOf

// NewConfigImportedItemAllOf instantiates a new ConfigImportedItemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigImportedItemAllOf(classId string, objectType string) *ConfigImportedItemAllOf {
	this := ConfigImportedItemAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = ""
	this.Status = &status
	return &this
}

// NewConfigImportedItemAllOfWithDefaults instantiates a new ConfigImportedItemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigImportedItemAllOfWithDefaults() *ConfigImportedItemAllOf {
	this := ConfigImportedItemAllOf{}
	var classId string = "config.ImportedItem"
	this.ClassId = classId
	var objectType string = "config.ImportedItem"
	this.ObjectType = objectType
	var status string = ""
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConfigImportedItemAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConfigImportedItemAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConfigImportedItemAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConfigImportedItemAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConfigImportedItemAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConfigImportedItemAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsShared returns the IsShared field value if set, zero value otherwise.
func (o *ConfigImportedItemAllOf) GetIsShared() bool {
	if o == nil || o.IsShared == nil {
		var ret bool
		return ret
	}
	return *o.IsShared
}

// GetIsSharedOk returns a tuple with the IsShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImportedItemAllOf) GetIsSharedOk() (*bool, bool) {
	if o == nil || o.IsShared == nil {
		return nil, false
	}
	return o.IsShared, true
}

// HasIsShared returns a boolean if a field has been set.
func (o *ConfigImportedItemAllOf) HasIsShared() bool {
	if o != nil && o.IsShared != nil {
		return true
	}

	return false
}

// SetIsShared gets a reference to the given bool and assigns it to the IsShared field.
func (o *ConfigImportedItemAllOf) SetIsShared(v bool) {
	o.IsShared = &v
}

// GetIsUpdated returns the IsUpdated field value if set, zero value otherwise.
func (o *ConfigImportedItemAllOf) GetIsUpdated() bool {
	if o == nil || o.IsUpdated == nil {
		var ret bool
		return ret
	}
	return *o.IsUpdated
}

// GetIsUpdatedOk returns a tuple with the IsUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImportedItemAllOf) GetIsUpdatedOk() (*bool, bool) {
	if o == nil || o.IsUpdated == nil {
		return nil, false
	}
	return o.IsUpdated, true
}

// HasIsUpdated returns a boolean if a field has been set.
func (o *ConfigImportedItemAllOf) HasIsUpdated() bool {
	if o != nil && o.IsUpdated != nil {
		return true
	}

	return false
}

// SetIsUpdated gets a reference to the given bool and assigns it to the IsUpdated field.
func (o *ConfigImportedItemAllOf) SetIsUpdated(v bool) {
	o.IsUpdated = &v
}

// GetItem returns the Item field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigImportedItemAllOf) GetItem() ConfigMoRef {
	if o == nil || o.Item.Get() == nil {
		var ret ConfigMoRef
		return ret
	}
	return *o.Item.Get()
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigImportedItemAllOf) GetItemOk() (*ConfigMoRef, bool) {
	if o == nil {
		return nil, false
	}
	return o.Item.Get(), o.Item.IsSet()
}

// HasItem returns a boolean if a field has been set.
func (o *ConfigImportedItemAllOf) HasItem() bool {
	if o != nil && o.Item.IsSet() {
		return true
	}

	return false
}

// SetItem gets a reference to the given NullableConfigMoRef and assigns it to the Item field.
func (o *ConfigImportedItemAllOf) SetItem(v ConfigMoRef) {
	o.Item.Set(&v)
}

// SetItemNil sets the value for Item to be an explicit nil
func (o *ConfigImportedItemAllOf) SetItemNil() {
	o.Item.Set(nil)
}

// UnsetItem ensures that no value is present for Item, not even an explicit nil
func (o *ConfigImportedItemAllOf) UnsetItem() {
	o.Item.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConfigImportedItemAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImportedItemAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConfigImportedItemAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConfigImportedItemAllOf) SetName(v string) {
	o.Name = &v
}

// GetNewMoid returns the NewMoid field value if set, zero value otherwise.
func (o *ConfigImportedItemAllOf) GetNewMoid() string {
	if o == nil || o.NewMoid == nil {
		var ret string
		return ret
	}
	return *o.NewMoid
}

// GetNewMoidOk returns a tuple with the NewMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImportedItemAllOf) GetNewMoidOk() (*string, bool) {
	if o == nil || o.NewMoid == nil {
		return nil, false
	}
	return o.NewMoid, true
}

// HasNewMoid returns a boolean if a field has been set.
func (o *ConfigImportedItemAllOf) HasNewMoid() bool {
	if o != nil && o.NewMoid != nil {
		return true
	}

	return false
}

// SetNewMoid gets a reference to the given string and assigns it to the NewMoid field.
func (o *ConfigImportedItemAllOf) SetNewMoid(v string) {
	o.NewMoid = &v
}

// GetServiceVersion returns the ServiceVersion field value if set, zero value otherwise.
func (o *ConfigImportedItemAllOf) GetServiceVersion() string {
	if o == nil || o.ServiceVersion == nil {
		var ret string
		return ret
	}
	return *o.ServiceVersion
}

// GetServiceVersionOk returns a tuple with the ServiceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImportedItemAllOf) GetServiceVersionOk() (*string, bool) {
	if o == nil || o.ServiceVersion == nil {
		return nil, false
	}
	return o.ServiceVersion, true
}

// HasServiceVersion returns a boolean if a field has been set.
func (o *ConfigImportedItemAllOf) HasServiceVersion() bool {
	if o != nil && o.ServiceVersion != nil {
		return true
	}

	return false
}

// SetServiceVersion gets a reference to the given string and assigns it to the ServiceVersion field.
func (o *ConfigImportedItemAllOf) SetServiceVersion(v string) {
	o.ServiceVersion = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConfigImportedItemAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImportedItemAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConfigImportedItemAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ConfigImportedItemAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *ConfigImportedItemAllOf) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImportedItemAllOf) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *ConfigImportedItemAllOf) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *ConfigImportedItemAllOf) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetImporter returns the Importer field value if set, zero value otherwise.
func (o *ConfigImportedItemAllOf) GetImporter() ConfigImporterRelationship {
	if o == nil || o.Importer == nil {
		var ret ConfigImporterRelationship
		return ret
	}
	return *o.Importer
}

// GetImporterOk returns a tuple with the Importer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImportedItemAllOf) GetImporterOk() (*ConfigImporterRelationship, bool) {
	if o == nil || o.Importer == nil {
		return nil, false
	}
	return o.Importer, true
}

// HasImporter returns a boolean if a field has been set.
func (o *ConfigImportedItemAllOf) HasImporter() bool {
	if o != nil && o.Importer != nil {
		return true
	}

	return false
}

// SetImporter gets a reference to the given ConfigImporterRelationship and assigns it to the Importer field.
func (o *ConfigImportedItemAllOf) SetImporter(v ConfigImporterRelationship) {
	o.Importer = &v
}

func (o ConfigImportedItemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IsShared != nil {
		toSerialize["IsShared"] = o.IsShared
	}
	if o.IsUpdated != nil {
		toSerialize["IsUpdated"] = o.IsUpdated
	}
	if o.Item.IsSet() {
		toSerialize["Item"] = o.Item.Get()
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.NewMoid != nil {
		toSerialize["NewMoid"] = o.NewMoid
	}
	if o.ServiceVersion != nil {
		toSerialize["ServiceVersion"] = o.ServiceVersion
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusMessage != nil {
		toSerialize["StatusMessage"] = o.StatusMessage
	}
	if o.Importer != nil {
		toSerialize["Importer"] = o.Importer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConfigImportedItemAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConfigImportedItemAllOf := _ConfigImportedItemAllOf{}

	if err = json.Unmarshal(bytes, &varConfigImportedItemAllOf); err == nil {
		*o = ConfigImportedItemAllOf(varConfigImportedItemAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsShared")
		delete(additionalProperties, "IsUpdated")
		delete(additionalProperties, "Item")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NewMoid")
		delete(additionalProperties, "ServiceVersion")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StatusMessage")
		delete(additionalProperties, "Importer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConfigImportedItemAllOf struct {
	value *ConfigImportedItemAllOf
	isSet bool
}

func (v NullableConfigImportedItemAllOf) Get() *ConfigImportedItemAllOf {
	return v.value
}

func (v *NullableConfigImportedItemAllOf) Set(val *ConfigImportedItemAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigImportedItemAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigImportedItemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigImportedItemAllOf(val *ConfigImportedItemAllOf) *NullableConfigImportedItemAllOf {
	return &NullableConfigImportedItemAllOf{value: val, isSet: true}
}

func (v NullableConfigImportedItemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigImportedItemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
