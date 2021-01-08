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
	"reflect"
	"strings"
)

// IaasUcsdManagedInfra Describes about UCSD Managed infrastructure statistics.
type IaasUcsdManagedInfra struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Total advanced catalogs in UCSD.
	AdvancedCatalogCount *int64 `json:"AdvancedCatalogCount,omitempty"`
	// Total bare metal catalogs in UCSD.
	BmCatalogCount *int64 `json:"BmCatalogCount,omitempty"`
	// Total service container catalogs in UCSD.
	ContainerCatalogCount *int64 `json:"ContainerCatalogCount,omitempty"`
	// Total ESXi hosts in UCSD.
	EsxiHostCount *int64 `json:"EsxiHostCount,omitempty"`
	// Total external (Ldap) groups in UCSD.
	ExternalGroupCount *int64 `json:"ExternalGroupCount,omitempty"`
	// Total HyperV hosts in UCSD.
	HypervHostCount *int64 `json:"HypervHostCount,omitempty"`
	// Total local groups in UCSD.
	LocalGroupCount *int64 `json:"LocalGroupCount,omitempty"`
	// Total standard catalogs in UCSD.
	StandardCatalogCount *int64 `json:"StandardCatalogCount,omitempty"`
	// Total user accounts in UCSD.
	UserCount *int64 `json:"UserCount,omitempty"`
	// Total virtual datacenters in UCSD.
	VdcCount *int64 `json:"VdcCount,omitempty"`
	// Total Virtual machines in UCSD.
	VmCount              *int64                    `json:"VmCount,omitempty"`
	Guid                 *IaasUcsdInfoRelationship `json:"Guid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IaasUcsdManagedInfra IaasUcsdManagedInfra

// NewIaasUcsdManagedInfra instantiates a new IaasUcsdManagedInfra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasUcsdManagedInfra(classId string, objectType string) *IaasUcsdManagedInfra {
	this := IaasUcsdManagedInfra{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIaasUcsdManagedInfraWithDefaults instantiates a new IaasUcsdManagedInfra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasUcsdManagedInfraWithDefaults() *IaasUcsdManagedInfra {
	this := IaasUcsdManagedInfra{}
	var classId string = "iaas.UcsdManagedInfra"
	this.ClassId = classId
	var objectType string = "iaas.UcsdManagedInfra"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IaasUcsdManagedInfra) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IaasUcsdManagedInfra) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IaasUcsdManagedInfra) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IaasUcsdManagedInfra) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IaasUcsdManagedInfra) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IaasUcsdManagedInfra) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdvancedCatalogCount returns the AdvancedCatalogCount field value if set, zero value otherwise.
func (o *IaasUcsdManagedInfra) GetAdvancedCatalogCount() int64 {
	if o == nil || o.AdvancedCatalogCount == nil {
		var ret int64
		return ret
	}
	return *o.AdvancedCatalogCount
}

// GetAdvancedCatalogCountOk returns a tuple with the AdvancedCatalogCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdManagedInfra) GetAdvancedCatalogCountOk() (*int64, bool) {
	if o == nil || o.AdvancedCatalogCount == nil {
		return nil, false
	}
	return o.AdvancedCatalogCount, true
}

// HasAdvancedCatalogCount returns a boolean if a field has been set.
func (o *IaasUcsdManagedInfra) HasAdvancedCatalogCount() bool {
	if o != nil && o.AdvancedCatalogCount != nil {
		return true
	}

	return false
}

// SetAdvancedCatalogCount gets a reference to the given int64 and assigns it to the AdvancedCatalogCount field.
func (o *IaasUcsdManagedInfra) SetAdvancedCatalogCount(v int64) {
	o.AdvancedCatalogCount = &v
}

// GetBmCatalogCount returns the BmCatalogCount field value if set, zero value otherwise.
func (o *IaasUcsdManagedInfra) GetBmCatalogCount() int64 {
	if o == nil || o.BmCatalogCount == nil {
		var ret int64
		return ret
	}
	return *o.BmCatalogCount
}

// GetBmCatalogCountOk returns a tuple with the BmCatalogCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdManagedInfra) GetBmCatalogCountOk() (*int64, bool) {
	if o == nil || o.BmCatalogCount == nil {
		return nil, false
	}
	return o.BmCatalogCount, true
}

// HasBmCatalogCount returns a boolean if a field has been set.
func (o *IaasUcsdManagedInfra) HasBmCatalogCount() bool {
	if o != nil && o.BmCatalogCount != nil {
		return true
	}

	return false
}

// SetBmCatalogCount gets a reference to the given int64 and assigns it to the BmCatalogCount field.
func (o *IaasUcsdManagedInfra) SetBmCatalogCount(v int64) {
	o.BmCatalogCount = &v
}

// GetContainerCatalogCount returns the ContainerCatalogCount field value if set, zero value otherwise.
func (o *IaasUcsdManagedInfra) GetContainerCatalogCount() int64 {
	if o == nil || o.ContainerCatalogCount == nil {
		var ret int64
		return ret
	}
	return *o.ContainerCatalogCount
}

// GetContainerCatalogCountOk returns a tuple with the ContainerCatalogCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdManagedInfra) GetContainerCatalogCountOk() (*int64, bool) {
	if o == nil || o.ContainerCatalogCount == nil {
		return nil, false
	}
	return o.ContainerCatalogCount, true
}

// HasContainerCatalogCount returns a boolean if a field has been set.
func (o *IaasUcsdManagedInfra) HasContainerCatalogCount() bool {
	if o != nil && o.ContainerCatalogCount != nil {
		return true
	}

	return false
}

// SetContainerCatalogCount gets a reference to the given int64 and assigns it to the ContainerCatalogCount field.
func (o *IaasUcsdManagedInfra) SetContainerCatalogCount(v int64) {
	o.ContainerCatalogCount = &v
}

// GetEsxiHostCount returns the EsxiHostCount field value if set, zero value otherwise.
func (o *IaasUcsdManagedInfra) GetEsxiHostCount() int64 {
	if o == nil || o.EsxiHostCount == nil {
		var ret int64
		return ret
	}
	return *o.EsxiHostCount
}

// GetEsxiHostCountOk returns a tuple with the EsxiHostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdManagedInfra) GetEsxiHostCountOk() (*int64, bool) {
	if o == nil || o.EsxiHostCount == nil {
		return nil, false
	}
	return o.EsxiHostCount, true
}

// HasEsxiHostCount returns a boolean if a field has been set.
func (o *IaasUcsdManagedInfra) HasEsxiHostCount() bool {
	if o != nil && o.EsxiHostCount != nil {
		return true
	}

	return false
}

// SetEsxiHostCount gets a reference to the given int64 and assigns it to the EsxiHostCount field.
func (o *IaasUcsdManagedInfra) SetEsxiHostCount(v int64) {
	o.EsxiHostCount = &v
}

// GetExternalGroupCount returns the ExternalGroupCount field value if set, zero value otherwise.
func (o *IaasUcsdManagedInfra) GetExternalGroupCount() int64 {
	if o == nil || o.ExternalGroupCount == nil {
		var ret int64
		return ret
	}
	return *o.ExternalGroupCount
}

// GetExternalGroupCountOk returns a tuple with the ExternalGroupCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdManagedInfra) GetExternalGroupCountOk() (*int64, bool) {
	if o == nil || o.ExternalGroupCount == nil {
		return nil, false
	}
	return o.ExternalGroupCount, true
}

// HasExternalGroupCount returns a boolean if a field has been set.
func (o *IaasUcsdManagedInfra) HasExternalGroupCount() bool {
	if o != nil && o.ExternalGroupCount != nil {
		return true
	}

	return false
}

// SetExternalGroupCount gets a reference to the given int64 and assigns it to the ExternalGroupCount field.
func (o *IaasUcsdManagedInfra) SetExternalGroupCount(v int64) {
	o.ExternalGroupCount = &v
}

// GetHypervHostCount returns the HypervHostCount field value if set, zero value otherwise.
func (o *IaasUcsdManagedInfra) GetHypervHostCount() int64 {
	if o == nil || o.HypervHostCount == nil {
		var ret int64
		return ret
	}
	return *o.HypervHostCount
}

// GetHypervHostCountOk returns a tuple with the HypervHostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdManagedInfra) GetHypervHostCountOk() (*int64, bool) {
	if o == nil || o.HypervHostCount == nil {
		return nil, false
	}
	return o.HypervHostCount, true
}

// HasHypervHostCount returns a boolean if a field has been set.
func (o *IaasUcsdManagedInfra) HasHypervHostCount() bool {
	if o != nil && o.HypervHostCount != nil {
		return true
	}

	return false
}

// SetHypervHostCount gets a reference to the given int64 and assigns it to the HypervHostCount field.
func (o *IaasUcsdManagedInfra) SetHypervHostCount(v int64) {
	o.HypervHostCount = &v
}

// GetLocalGroupCount returns the LocalGroupCount field value if set, zero value otherwise.
func (o *IaasUcsdManagedInfra) GetLocalGroupCount() int64 {
	if o == nil || o.LocalGroupCount == nil {
		var ret int64
		return ret
	}
	return *o.LocalGroupCount
}

// GetLocalGroupCountOk returns a tuple with the LocalGroupCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdManagedInfra) GetLocalGroupCountOk() (*int64, bool) {
	if o == nil || o.LocalGroupCount == nil {
		return nil, false
	}
	return o.LocalGroupCount, true
}

// HasLocalGroupCount returns a boolean if a field has been set.
func (o *IaasUcsdManagedInfra) HasLocalGroupCount() bool {
	if o != nil && o.LocalGroupCount != nil {
		return true
	}

	return false
}

// SetLocalGroupCount gets a reference to the given int64 and assigns it to the LocalGroupCount field.
func (o *IaasUcsdManagedInfra) SetLocalGroupCount(v int64) {
	o.LocalGroupCount = &v
}

// GetStandardCatalogCount returns the StandardCatalogCount field value if set, zero value otherwise.
func (o *IaasUcsdManagedInfra) GetStandardCatalogCount() int64 {
	if o == nil || o.StandardCatalogCount == nil {
		var ret int64
		return ret
	}
	return *o.StandardCatalogCount
}

// GetStandardCatalogCountOk returns a tuple with the StandardCatalogCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdManagedInfra) GetStandardCatalogCountOk() (*int64, bool) {
	if o == nil || o.StandardCatalogCount == nil {
		return nil, false
	}
	return o.StandardCatalogCount, true
}

// HasStandardCatalogCount returns a boolean if a field has been set.
func (o *IaasUcsdManagedInfra) HasStandardCatalogCount() bool {
	if o != nil && o.StandardCatalogCount != nil {
		return true
	}

	return false
}

// SetStandardCatalogCount gets a reference to the given int64 and assigns it to the StandardCatalogCount field.
func (o *IaasUcsdManagedInfra) SetStandardCatalogCount(v int64) {
	o.StandardCatalogCount = &v
}

// GetUserCount returns the UserCount field value if set, zero value otherwise.
func (o *IaasUcsdManagedInfra) GetUserCount() int64 {
	if o == nil || o.UserCount == nil {
		var ret int64
		return ret
	}
	return *o.UserCount
}

// GetUserCountOk returns a tuple with the UserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdManagedInfra) GetUserCountOk() (*int64, bool) {
	if o == nil || o.UserCount == nil {
		return nil, false
	}
	return o.UserCount, true
}

// HasUserCount returns a boolean if a field has been set.
func (o *IaasUcsdManagedInfra) HasUserCount() bool {
	if o != nil && o.UserCount != nil {
		return true
	}

	return false
}

// SetUserCount gets a reference to the given int64 and assigns it to the UserCount field.
func (o *IaasUcsdManagedInfra) SetUserCount(v int64) {
	o.UserCount = &v
}

// GetVdcCount returns the VdcCount field value if set, zero value otherwise.
func (o *IaasUcsdManagedInfra) GetVdcCount() int64 {
	if o == nil || o.VdcCount == nil {
		var ret int64
		return ret
	}
	return *o.VdcCount
}

// GetVdcCountOk returns a tuple with the VdcCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdManagedInfra) GetVdcCountOk() (*int64, bool) {
	if o == nil || o.VdcCount == nil {
		return nil, false
	}
	return o.VdcCount, true
}

// HasVdcCount returns a boolean if a field has been set.
func (o *IaasUcsdManagedInfra) HasVdcCount() bool {
	if o != nil && o.VdcCount != nil {
		return true
	}

	return false
}

// SetVdcCount gets a reference to the given int64 and assigns it to the VdcCount field.
func (o *IaasUcsdManagedInfra) SetVdcCount(v int64) {
	o.VdcCount = &v
}

// GetVmCount returns the VmCount field value if set, zero value otherwise.
func (o *IaasUcsdManagedInfra) GetVmCount() int64 {
	if o == nil || o.VmCount == nil {
		var ret int64
		return ret
	}
	return *o.VmCount
}

// GetVmCountOk returns a tuple with the VmCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdManagedInfra) GetVmCountOk() (*int64, bool) {
	if o == nil || o.VmCount == nil {
		return nil, false
	}
	return o.VmCount, true
}

// HasVmCount returns a boolean if a field has been set.
func (o *IaasUcsdManagedInfra) HasVmCount() bool {
	if o != nil && o.VmCount != nil {
		return true
	}

	return false
}

// SetVmCount gets a reference to the given int64 and assigns it to the VmCount field.
func (o *IaasUcsdManagedInfra) SetVmCount(v int64) {
	o.VmCount = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *IaasUcsdManagedInfra) GetGuid() IaasUcsdInfoRelationship {
	if o == nil || o.Guid == nil {
		var ret IaasUcsdInfoRelationship
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdManagedInfra) GetGuidOk() (*IaasUcsdInfoRelationship, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *IaasUcsdManagedInfra) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given IaasUcsdInfoRelationship and assigns it to the Guid field.
func (o *IaasUcsdManagedInfra) SetGuid(v IaasUcsdInfoRelationship) {
	o.Guid = &v
}

func (o IaasUcsdManagedInfra) MarshalJSON() ([]byte, error) {
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
	if o.AdvancedCatalogCount != nil {
		toSerialize["AdvancedCatalogCount"] = o.AdvancedCatalogCount
	}
	if o.BmCatalogCount != nil {
		toSerialize["BmCatalogCount"] = o.BmCatalogCount
	}
	if o.ContainerCatalogCount != nil {
		toSerialize["ContainerCatalogCount"] = o.ContainerCatalogCount
	}
	if o.EsxiHostCount != nil {
		toSerialize["EsxiHostCount"] = o.EsxiHostCount
	}
	if o.ExternalGroupCount != nil {
		toSerialize["ExternalGroupCount"] = o.ExternalGroupCount
	}
	if o.HypervHostCount != nil {
		toSerialize["HypervHostCount"] = o.HypervHostCount
	}
	if o.LocalGroupCount != nil {
		toSerialize["LocalGroupCount"] = o.LocalGroupCount
	}
	if o.StandardCatalogCount != nil {
		toSerialize["StandardCatalogCount"] = o.StandardCatalogCount
	}
	if o.UserCount != nil {
		toSerialize["UserCount"] = o.UserCount
	}
	if o.VdcCount != nil {
		toSerialize["VdcCount"] = o.VdcCount
	}
	if o.VmCount != nil {
		toSerialize["VmCount"] = o.VmCount
	}
	if o.Guid != nil {
		toSerialize["Guid"] = o.Guid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IaasUcsdManagedInfra) UnmarshalJSON(bytes []byte) (err error) {
	type IaasUcsdManagedInfraWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Total advanced catalogs in UCSD.
		AdvancedCatalogCount *int64 `json:"AdvancedCatalogCount,omitempty"`
		// Total bare metal catalogs in UCSD.
		BmCatalogCount *int64 `json:"BmCatalogCount,omitempty"`
		// Total service container catalogs in UCSD.
		ContainerCatalogCount *int64 `json:"ContainerCatalogCount,omitempty"`
		// Total ESXi hosts in UCSD.
		EsxiHostCount *int64 `json:"EsxiHostCount,omitempty"`
		// Total external (Ldap) groups in UCSD.
		ExternalGroupCount *int64 `json:"ExternalGroupCount,omitempty"`
		// Total HyperV hosts in UCSD.
		HypervHostCount *int64 `json:"HypervHostCount,omitempty"`
		// Total local groups in UCSD.
		LocalGroupCount *int64 `json:"LocalGroupCount,omitempty"`
		// Total standard catalogs in UCSD.
		StandardCatalogCount *int64 `json:"StandardCatalogCount,omitempty"`
		// Total user accounts in UCSD.
		UserCount *int64 `json:"UserCount,omitempty"`
		// Total virtual datacenters in UCSD.
		VdcCount *int64 `json:"VdcCount,omitempty"`
		// Total Virtual machines in UCSD.
		VmCount *int64                    `json:"VmCount,omitempty"`
		Guid    *IaasUcsdInfoRelationship `json:"Guid,omitempty"`
	}

	varIaasUcsdManagedInfraWithoutEmbeddedStruct := IaasUcsdManagedInfraWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIaasUcsdManagedInfraWithoutEmbeddedStruct)
	if err == nil {
		varIaasUcsdManagedInfra := _IaasUcsdManagedInfra{}
		varIaasUcsdManagedInfra.ClassId = varIaasUcsdManagedInfraWithoutEmbeddedStruct.ClassId
		varIaasUcsdManagedInfra.ObjectType = varIaasUcsdManagedInfraWithoutEmbeddedStruct.ObjectType
		varIaasUcsdManagedInfra.AdvancedCatalogCount = varIaasUcsdManagedInfraWithoutEmbeddedStruct.AdvancedCatalogCount
		varIaasUcsdManagedInfra.BmCatalogCount = varIaasUcsdManagedInfraWithoutEmbeddedStruct.BmCatalogCount
		varIaasUcsdManagedInfra.ContainerCatalogCount = varIaasUcsdManagedInfraWithoutEmbeddedStruct.ContainerCatalogCount
		varIaasUcsdManagedInfra.EsxiHostCount = varIaasUcsdManagedInfraWithoutEmbeddedStruct.EsxiHostCount
		varIaasUcsdManagedInfra.ExternalGroupCount = varIaasUcsdManagedInfraWithoutEmbeddedStruct.ExternalGroupCount
		varIaasUcsdManagedInfra.HypervHostCount = varIaasUcsdManagedInfraWithoutEmbeddedStruct.HypervHostCount
		varIaasUcsdManagedInfra.LocalGroupCount = varIaasUcsdManagedInfraWithoutEmbeddedStruct.LocalGroupCount
		varIaasUcsdManagedInfra.StandardCatalogCount = varIaasUcsdManagedInfraWithoutEmbeddedStruct.StandardCatalogCount
		varIaasUcsdManagedInfra.UserCount = varIaasUcsdManagedInfraWithoutEmbeddedStruct.UserCount
		varIaasUcsdManagedInfra.VdcCount = varIaasUcsdManagedInfraWithoutEmbeddedStruct.VdcCount
		varIaasUcsdManagedInfra.VmCount = varIaasUcsdManagedInfraWithoutEmbeddedStruct.VmCount
		varIaasUcsdManagedInfra.Guid = varIaasUcsdManagedInfraWithoutEmbeddedStruct.Guid
		*o = IaasUcsdManagedInfra(varIaasUcsdManagedInfra)
	} else {
		return err
	}

	varIaasUcsdManagedInfra := _IaasUcsdManagedInfra{}

	err = json.Unmarshal(bytes, &varIaasUcsdManagedInfra)
	if err == nil {
		o.MoBaseMo = varIaasUcsdManagedInfra.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdvancedCatalogCount")
		delete(additionalProperties, "BmCatalogCount")
		delete(additionalProperties, "ContainerCatalogCount")
		delete(additionalProperties, "EsxiHostCount")
		delete(additionalProperties, "ExternalGroupCount")
		delete(additionalProperties, "HypervHostCount")
		delete(additionalProperties, "LocalGroupCount")
		delete(additionalProperties, "StandardCatalogCount")
		delete(additionalProperties, "UserCount")
		delete(additionalProperties, "VdcCount")
		delete(additionalProperties, "VmCount")
		delete(additionalProperties, "Guid")

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

type NullableIaasUcsdManagedInfra struct {
	value *IaasUcsdManagedInfra
	isSet bool
}

func (v NullableIaasUcsdManagedInfra) Get() *IaasUcsdManagedInfra {
	return v.value
}

func (v *NullableIaasUcsdManagedInfra) Set(val *IaasUcsdManagedInfra) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasUcsdManagedInfra) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasUcsdManagedInfra) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasUcsdManagedInfra(val *IaasUcsdManagedInfra) *NullableIaasUcsdManagedInfra {
	return &NullableIaasUcsdManagedInfra{value: val, isSet: true}
}

func (v NullableIaasUcsdManagedInfra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasUcsdManagedInfra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
