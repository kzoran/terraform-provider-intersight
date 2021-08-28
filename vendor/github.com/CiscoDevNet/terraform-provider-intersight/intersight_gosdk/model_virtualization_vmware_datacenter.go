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
)

// VirtualizationVmwareDatacenter Datacenter object in VMware inventory. It is the logical container for all other objects like Datastore, Host, VirtualMachine, etc.
type VirtualizationVmwareDatacenter struct {
	VirtualizationBaseDatacenter
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Count of all clusters associated with this DC.
	ClusterCount *int64 `json:"ClusterCount,omitempty"`
	// Count of all datastores associated with this DC.
	DatastoreCount *int64 `json:"DatastoreCount,omitempty"`
	// Count of all hosts associated with this DC.
	HostCount *int64 `json:"HostCount,omitempty"`
	// Inventory path of the DC.
	InventoryPath *string `json:"InventoryPath,omitempty"`
	// Count of all networks associated with this datacenter (DC).
	NetworkCount *int64 `json:"NetworkCount,omitempty"`
	// Count of all virtual machines (VMs) associated with this DC.
	VmCount              *int64                                   `json:"VmCount,omitempty"`
	HypervisorManager    *VirtualizationVmwareVcenterRelationship `json:"HypervisorManager,omitempty"`
	ParentFolder         *VirtualizationVmwareFolderRelationship  `json:"ParentFolder,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareDatacenter VirtualizationVmwareDatacenter

// NewVirtualizationVmwareDatacenter instantiates a new VirtualizationVmwareDatacenter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareDatacenter(classId string, objectType string) *VirtualizationVmwareDatacenter {
	this := VirtualizationVmwareDatacenter{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareDatacenterWithDefaults instantiates a new VirtualizationVmwareDatacenter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareDatacenterWithDefaults() *VirtualizationVmwareDatacenter {
	this := VirtualizationVmwareDatacenter{}
	var classId string = "virtualization.VmwareDatacenter"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareDatacenter"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareDatacenter) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareDatacenter) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareDatacenter) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareDatacenter) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterCount returns the ClusterCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetClusterCount() int64 {
	if o == nil || o.ClusterCount == nil {
		var ret int64
		return ret
	}
	return *o.ClusterCount
}

// GetClusterCountOk returns a tuple with the ClusterCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetClusterCountOk() (*int64, bool) {
	if o == nil || o.ClusterCount == nil {
		return nil, false
	}
	return o.ClusterCount, true
}

// HasClusterCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasClusterCount() bool {
	if o != nil && o.ClusterCount != nil {
		return true
	}

	return false
}

// SetClusterCount gets a reference to the given int64 and assigns it to the ClusterCount field.
func (o *VirtualizationVmwareDatacenter) SetClusterCount(v int64) {
	o.ClusterCount = &v
}

// GetDatastoreCount returns the DatastoreCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetDatastoreCount() int64 {
	if o == nil || o.DatastoreCount == nil {
		var ret int64
		return ret
	}
	return *o.DatastoreCount
}

// GetDatastoreCountOk returns a tuple with the DatastoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetDatastoreCountOk() (*int64, bool) {
	if o == nil || o.DatastoreCount == nil {
		return nil, false
	}
	return o.DatastoreCount, true
}

// HasDatastoreCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasDatastoreCount() bool {
	if o != nil && o.DatastoreCount != nil {
		return true
	}

	return false
}

// SetDatastoreCount gets a reference to the given int64 and assigns it to the DatastoreCount field.
func (o *VirtualizationVmwareDatacenter) SetDatastoreCount(v int64) {
	o.DatastoreCount = &v
}

// GetHostCount returns the HostCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetHostCount() int64 {
	if o == nil || o.HostCount == nil {
		var ret int64
		return ret
	}
	return *o.HostCount
}

// GetHostCountOk returns a tuple with the HostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetHostCountOk() (*int64, bool) {
	if o == nil || o.HostCount == nil {
		return nil, false
	}
	return o.HostCount, true
}

// HasHostCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasHostCount() bool {
	if o != nil && o.HostCount != nil {
		return true
	}

	return false
}

// SetHostCount gets a reference to the given int64 and assigns it to the HostCount field.
func (o *VirtualizationVmwareDatacenter) SetHostCount(v int64) {
	o.HostCount = &v
}

// GetInventoryPath returns the InventoryPath field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetInventoryPath() string {
	if o == nil || o.InventoryPath == nil {
		var ret string
		return ret
	}
	return *o.InventoryPath
}

// GetInventoryPathOk returns a tuple with the InventoryPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetInventoryPathOk() (*string, bool) {
	if o == nil || o.InventoryPath == nil {
		return nil, false
	}
	return o.InventoryPath, true
}

// HasInventoryPath returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasInventoryPath() bool {
	if o != nil && o.InventoryPath != nil {
		return true
	}

	return false
}

// SetInventoryPath gets a reference to the given string and assigns it to the InventoryPath field.
func (o *VirtualizationVmwareDatacenter) SetInventoryPath(v string) {
	o.InventoryPath = &v
}

// GetNetworkCount returns the NetworkCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetNetworkCount() int64 {
	if o == nil || o.NetworkCount == nil {
		var ret int64
		return ret
	}
	return *o.NetworkCount
}

// GetNetworkCountOk returns a tuple with the NetworkCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetNetworkCountOk() (*int64, bool) {
	if o == nil || o.NetworkCount == nil {
		return nil, false
	}
	return o.NetworkCount, true
}

// HasNetworkCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasNetworkCount() bool {
	if o != nil && o.NetworkCount != nil {
		return true
	}

	return false
}

// SetNetworkCount gets a reference to the given int64 and assigns it to the NetworkCount field.
func (o *VirtualizationVmwareDatacenter) SetNetworkCount(v int64) {
	o.NetworkCount = &v
}

// GetVmCount returns the VmCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetVmCount() int64 {
	if o == nil || o.VmCount == nil {
		var ret int64
		return ret
	}
	return *o.VmCount
}

// GetVmCountOk returns a tuple with the VmCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetVmCountOk() (*int64, bool) {
	if o == nil || o.VmCount == nil {
		return nil, false
	}
	return o.VmCount, true
}

// HasVmCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasVmCount() bool {
	if o != nil && o.VmCount != nil {
		return true
	}

	return false
}

// SetVmCount gets a reference to the given int64 and assigns it to the VmCount field.
func (o *VirtualizationVmwareDatacenter) SetVmCount(v int64) {
	o.VmCount = &v
}

// GetHypervisorManager returns the HypervisorManager field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetHypervisorManager() VirtualizationVmwareVcenterRelationship {
	if o == nil || o.HypervisorManager == nil {
		var ret VirtualizationVmwareVcenterRelationship
		return ret
	}
	return *o.HypervisorManager
}

// GetHypervisorManagerOk returns a tuple with the HypervisorManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetHypervisorManagerOk() (*VirtualizationVmwareVcenterRelationship, bool) {
	if o == nil || o.HypervisorManager == nil {
		return nil, false
	}
	return o.HypervisorManager, true
}

// HasHypervisorManager returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasHypervisorManager() bool {
	if o != nil && o.HypervisorManager != nil {
		return true
	}

	return false
}

// SetHypervisorManager gets a reference to the given VirtualizationVmwareVcenterRelationship and assigns it to the HypervisorManager field.
func (o *VirtualizationVmwareDatacenter) SetHypervisorManager(v VirtualizationVmwareVcenterRelationship) {
	o.HypervisorManager = &v
}

// GetParentFolder returns the ParentFolder field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatacenter) GetParentFolder() VirtualizationVmwareFolderRelationship {
	if o == nil || o.ParentFolder == nil {
		var ret VirtualizationVmwareFolderRelationship
		return ret
	}
	return *o.ParentFolder
}

// GetParentFolderOk returns a tuple with the ParentFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatacenter) GetParentFolderOk() (*VirtualizationVmwareFolderRelationship, bool) {
	if o == nil || o.ParentFolder == nil {
		return nil, false
	}
	return o.ParentFolder, true
}

// HasParentFolder returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatacenter) HasParentFolder() bool {
	if o != nil && o.ParentFolder != nil {
		return true
	}

	return false
}

// SetParentFolder gets a reference to the given VirtualizationVmwareFolderRelationship and assigns it to the ParentFolder field.
func (o *VirtualizationVmwareDatacenter) SetParentFolder(v VirtualizationVmwareFolderRelationship) {
	o.ParentFolder = &v
}

func (o VirtualizationVmwareDatacenter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseDatacenter, errVirtualizationBaseDatacenter := json.Marshal(o.VirtualizationBaseDatacenter)
	if errVirtualizationBaseDatacenter != nil {
		return []byte{}, errVirtualizationBaseDatacenter
	}
	errVirtualizationBaseDatacenter = json.Unmarshal([]byte(serializedVirtualizationBaseDatacenter), &toSerialize)
	if errVirtualizationBaseDatacenter != nil {
		return []byte{}, errVirtualizationBaseDatacenter
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClusterCount != nil {
		toSerialize["ClusterCount"] = o.ClusterCount
	}
	if o.DatastoreCount != nil {
		toSerialize["DatastoreCount"] = o.DatastoreCount
	}
	if o.HostCount != nil {
		toSerialize["HostCount"] = o.HostCount
	}
	if o.InventoryPath != nil {
		toSerialize["InventoryPath"] = o.InventoryPath
	}
	if o.NetworkCount != nil {
		toSerialize["NetworkCount"] = o.NetworkCount
	}
	if o.VmCount != nil {
		toSerialize["VmCount"] = o.VmCount
	}
	if o.HypervisorManager != nil {
		toSerialize["HypervisorManager"] = o.HypervisorManager
	}
	if o.ParentFolder != nil {
		toSerialize["ParentFolder"] = o.ParentFolder
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareDatacenter) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVmwareDatacenterWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Count of all clusters associated with this DC.
		ClusterCount *int64 `json:"ClusterCount,omitempty"`
		// Count of all datastores associated with this DC.
		DatastoreCount *int64 `json:"DatastoreCount,omitempty"`
		// Count of all hosts associated with this DC.
		HostCount *int64 `json:"HostCount,omitempty"`
		// Inventory path of the DC.
		InventoryPath *string `json:"InventoryPath,omitempty"`
		// Count of all networks associated with this datacenter (DC).
		NetworkCount *int64 `json:"NetworkCount,omitempty"`
		// Count of all virtual machines (VMs) associated with this DC.
		VmCount           *int64                                   `json:"VmCount,omitempty"`
		HypervisorManager *VirtualizationVmwareVcenterRelationship `json:"HypervisorManager,omitempty"`
		ParentFolder      *VirtualizationVmwareFolderRelationship  `json:"ParentFolder,omitempty"`
	}

	varVirtualizationVmwareDatacenterWithoutEmbeddedStruct := VirtualizationVmwareDatacenterWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareDatacenterWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareDatacenter := _VirtualizationVmwareDatacenter{}
		varVirtualizationVmwareDatacenter.ClassId = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareDatacenter.ObjectType = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareDatacenter.ClusterCount = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.ClusterCount
		varVirtualizationVmwareDatacenter.DatastoreCount = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.DatastoreCount
		varVirtualizationVmwareDatacenter.HostCount = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.HostCount
		varVirtualizationVmwareDatacenter.InventoryPath = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.InventoryPath
		varVirtualizationVmwareDatacenter.NetworkCount = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.NetworkCount
		varVirtualizationVmwareDatacenter.VmCount = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.VmCount
		varVirtualizationVmwareDatacenter.HypervisorManager = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.HypervisorManager
		varVirtualizationVmwareDatacenter.ParentFolder = varVirtualizationVmwareDatacenterWithoutEmbeddedStruct.ParentFolder
		*o = VirtualizationVmwareDatacenter(varVirtualizationVmwareDatacenter)
	} else {
		return err
	}

	varVirtualizationVmwareDatacenter := _VirtualizationVmwareDatacenter{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareDatacenter)
	if err == nil {
		o.VirtualizationBaseDatacenter = varVirtualizationVmwareDatacenter.VirtualizationBaseDatacenter
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterCount")
		delete(additionalProperties, "DatastoreCount")
		delete(additionalProperties, "HostCount")
		delete(additionalProperties, "InventoryPath")
		delete(additionalProperties, "NetworkCount")
		delete(additionalProperties, "VmCount")
		delete(additionalProperties, "HypervisorManager")
		delete(additionalProperties, "ParentFolder")

		// remove fields from embedded structs
		reflectVirtualizationBaseDatacenter := reflect.ValueOf(o.VirtualizationBaseDatacenter)
		for i := 0; i < reflectVirtualizationBaseDatacenter.Type().NumField(); i++ {
			t := reflectVirtualizationBaseDatacenter.Type().Field(i)

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

type NullableVirtualizationVmwareDatacenter struct {
	value *VirtualizationVmwareDatacenter
	isSet bool
}

func (v NullableVirtualizationVmwareDatacenter) Get() *VirtualizationVmwareDatacenter {
	return v.value
}

func (v *NullableVirtualizationVmwareDatacenter) Set(val *VirtualizationVmwareDatacenter) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareDatacenter) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareDatacenter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareDatacenter(val *VirtualizationVmwareDatacenter) *NullableVirtualizationVmwareDatacenter {
	return &NullableVirtualizationVmwareDatacenter{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareDatacenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareDatacenter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
