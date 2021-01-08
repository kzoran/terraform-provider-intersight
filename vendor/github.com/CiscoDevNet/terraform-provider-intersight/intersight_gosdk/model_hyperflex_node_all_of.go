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

// HyperflexNodeAllOf Definition of the list of properties defined in 'hyperflex.Node', excluding properties defined in parent classes.
type HyperflexNodeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The build number of the hypervisor running on the host.
	BuildNumber *string `json:"BuildNumber,omitempty"`
	// The user-friendly string representation of the hypervisor version of the host.
	DisplayVersion *string `json:"DisplayVersion,omitempty"`
	// The hostname configured for the hypervisor running on the host.
	HostName *string `json:"HostName,omitempty"`
	// The type of hypervisor running on the host.
	Hypervisor *string                             `json:"Hypervisor,omitempty"`
	Identity   NullableHyperflexHxUuIdDt           `json:"Identity,omitempty"`
	Ip         NullableHyperflexHxNetworkAddressDt `json:"Ip,omitempty"`
	// The admin state of lockdown mode on the host. If 'true', lockdown mode is enabled.
	Lockdown *bool `json:"Lockdown,omitempty"`
	// The model of the host server.
	ModelNumber *string `json:"ModelNumber,omitempty"`
	// The role of the host in the HyperFlex cluster. Specifies whether this host is used for compute or for both compute and storage. * `UNKNOWN` - The role of the HyperFlex cluster node is not known. * `STORAGE` - The HyperFlex cluster node provides both storage and compute resources for the cluster. * `COMPUTE` - The HyperFlex cluster node provides compute resources for the cluster.
	Role *string `json:"Role,omitempty"`
	// The serial of the host server.
	SerialNumber *string `json:"SerialNumber,omitempty"`
	// The status of the host. Indicates whether the hypervisor is online. * `UNKNOWN` - The host status cannot be determined. * `ONLINE` - The host is online and operational. * `OFFLINE` - The host is offline and is currently not participating in the HyperFlex cluster. * `INMAINTENANCE` - The host is not participating in the HyperFlex cluster because of a maintenance operation, such as firmware or data platform upgrade. * `DEGRADED` - The host is degraded and may not be performing in its full operational capacity.
	Status *string `json:"Status,omitempty"`
	// The version of the hypervisor running on the host.
	Version              *string                         `json:"Version,omitempty"`
	Cluster              *HyperflexClusterRelationship   `json:"Cluster,omitempty"`
	ClusterMember        *AssetClusterMemberRelationship `json:"ClusterMember,omitempty"`
	PhysicalServer       *ComputePhysicalRelationship    `json:"PhysicalServer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexNodeAllOf HyperflexNodeAllOf

// NewHyperflexNodeAllOf instantiates a new HyperflexNodeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexNodeAllOf(classId string, objectType string) *HyperflexNodeAllOf {
	this := HyperflexNodeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var role string = "UNKNOWN"
	this.Role = &role
	var status string = "UNKNOWN"
	this.Status = &status
	return &this
}

// NewHyperflexNodeAllOfWithDefaults instantiates a new HyperflexNodeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexNodeAllOfWithDefaults() *HyperflexNodeAllOf {
	this := HyperflexNodeAllOf{}
	var classId string = "hyperflex.Node"
	this.ClassId = classId
	var objectType string = "hyperflex.Node"
	this.ObjectType = objectType
	var role string = "UNKNOWN"
	this.Role = &role
	var status string = "UNKNOWN"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexNodeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexNodeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexNodeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexNodeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexNodeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexNodeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBuildNumber returns the BuildNumber field value if set, zero value otherwise.
func (o *HyperflexNodeAllOf) GetBuildNumber() string {
	if o == nil || o.BuildNumber == nil {
		var ret string
		return ret
	}
	return *o.BuildNumber
}

// GetBuildNumberOk returns a tuple with the BuildNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeAllOf) GetBuildNumberOk() (*string, bool) {
	if o == nil || o.BuildNumber == nil {
		return nil, false
	}
	return o.BuildNumber, true
}

// HasBuildNumber returns a boolean if a field has been set.
func (o *HyperflexNodeAllOf) HasBuildNumber() bool {
	if o != nil && o.BuildNumber != nil {
		return true
	}

	return false
}

// SetBuildNumber gets a reference to the given string and assigns it to the BuildNumber field.
func (o *HyperflexNodeAllOf) SetBuildNumber(v string) {
	o.BuildNumber = &v
}

// GetDisplayVersion returns the DisplayVersion field value if set, zero value otherwise.
func (o *HyperflexNodeAllOf) GetDisplayVersion() string {
	if o == nil || o.DisplayVersion == nil {
		var ret string
		return ret
	}
	return *o.DisplayVersion
}

// GetDisplayVersionOk returns a tuple with the DisplayVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeAllOf) GetDisplayVersionOk() (*string, bool) {
	if o == nil || o.DisplayVersion == nil {
		return nil, false
	}
	return o.DisplayVersion, true
}

// HasDisplayVersion returns a boolean if a field has been set.
func (o *HyperflexNodeAllOf) HasDisplayVersion() bool {
	if o != nil && o.DisplayVersion != nil {
		return true
	}

	return false
}

// SetDisplayVersion gets a reference to the given string and assigns it to the DisplayVersion field.
func (o *HyperflexNodeAllOf) SetDisplayVersion(v string) {
	o.DisplayVersion = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *HyperflexNodeAllOf) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeAllOf) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *HyperflexNodeAllOf) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *HyperflexNodeAllOf) SetHostName(v string) {
	o.HostName = &v
}

// GetHypervisor returns the Hypervisor field value if set, zero value otherwise.
func (o *HyperflexNodeAllOf) GetHypervisor() string {
	if o == nil || o.Hypervisor == nil {
		var ret string
		return ret
	}
	return *o.Hypervisor
}

// GetHypervisorOk returns a tuple with the Hypervisor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeAllOf) GetHypervisorOk() (*string, bool) {
	if o == nil || o.Hypervisor == nil {
		return nil, false
	}
	return o.Hypervisor, true
}

// HasHypervisor returns a boolean if a field has been set.
func (o *HyperflexNodeAllOf) HasHypervisor() bool {
	if o != nil && o.Hypervisor != nil {
		return true
	}

	return false
}

// SetHypervisor gets a reference to the given string and assigns it to the Hypervisor field.
func (o *HyperflexNodeAllOf) SetHypervisor(v string) {
	o.Hypervisor = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexNodeAllOf) GetIdentity() HyperflexHxUuIdDt {
	if o == nil || o.Identity.Get() == nil {
		var ret HyperflexHxUuIdDt
		return ret
	}
	return *o.Identity.Get()
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexNodeAllOf) GetIdentityOk() (*HyperflexHxUuIdDt, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identity.Get(), o.Identity.IsSet()
}

// HasIdentity returns a boolean if a field has been set.
func (o *HyperflexNodeAllOf) HasIdentity() bool {
	if o != nil && o.Identity.IsSet() {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given NullableHyperflexHxUuIdDt and assigns it to the Identity field.
func (o *HyperflexNodeAllOf) SetIdentity(v HyperflexHxUuIdDt) {
	o.Identity.Set(&v)
}

// SetIdentityNil sets the value for Identity to be an explicit nil
func (o *HyperflexNodeAllOf) SetIdentityNil() {
	o.Identity.Set(nil)
}

// UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
func (o *HyperflexNodeAllOf) UnsetIdentity() {
	o.Identity.Unset()
}

// GetIp returns the Ip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexNodeAllOf) GetIp() HyperflexHxNetworkAddressDt {
	if o == nil || o.Ip.Get() == nil {
		var ret HyperflexHxNetworkAddressDt
		return ret
	}
	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexNodeAllOf) GetIpOk() (*HyperflexHxNetworkAddressDt, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// HasIp returns a boolean if a field has been set.
func (o *HyperflexNodeAllOf) HasIp() bool {
	if o != nil && o.Ip.IsSet() {
		return true
	}

	return false
}

// SetIp gets a reference to the given NullableHyperflexHxNetworkAddressDt and assigns it to the Ip field.
func (o *HyperflexNodeAllOf) SetIp(v HyperflexHxNetworkAddressDt) {
	o.Ip.Set(&v)
}

// SetIpNil sets the value for Ip to be an explicit nil
func (o *HyperflexNodeAllOf) SetIpNil() {
	o.Ip.Set(nil)
}

// UnsetIp ensures that no value is present for Ip, not even an explicit nil
func (o *HyperflexNodeAllOf) UnsetIp() {
	o.Ip.Unset()
}

// GetLockdown returns the Lockdown field value if set, zero value otherwise.
func (o *HyperflexNodeAllOf) GetLockdown() bool {
	if o == nil || o.Lockdown == nil {
		var ret bool
		return ret
	}
	return *o.Lockdown
}

// GetLockdownOk returns a tuple with the Lockdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeAllOf) GetLockdownOk() (*bool, bool) {
	if o == nil || o.Lockdown == nil {
		return nil, false
	}
	return o.Lockdown, true
}

// HasLockdown returns a boolean if a field has been set.
func (o *HyperflexNodeAllOf) HasLockdown() bool {
	if o != nil && o.Lockdown != nil {
		return true
	}

	return false
}

// SetLockdown gets a reference to the given bool and assigns it to the Lockdown field.
func (o *HyperflexNodeAllOf) SetLockdown(v bool) {
	o.Lockdown = &v
}

// GetModelNumber returns the ModelNumber field value if set, zero value otherwise.
func (o *HyperflexNodeAllOf) GetModelNumber() string {
	if o == nil || o.ModelNumber == nil {
		var ret string
		return ret
	}
	return *o.ModelNumber
}

// GetModelNumberOk returns a tuple with the ModelNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeAllOf) GetModelNumberOk() (*string, bool) {
	if o == nil || o.ModelNumber == nil {
		return nil, false
	}
	return o.ModelNumber, true
}

// HasModelNumber returns a boolean if a field has been set.
func (o *HyperflexNodeAllOf) HasModelNumber() bool {
	if o != nil && o.ModelNumber != nil {
		return true
	}

	return false
}

// SetModelNumber gets a reference to the given string and assigns it to the ModelNumber field.
func (o *HyperflexNodeAllOf) SetModelNumber(v string) {
	o.ModelNumber = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *HyperflexNodeAllOf) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeAllOf) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *HyperflexNodeAllOf) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *HyperflexNodeAllOf) SetRole(v string) {
	o.Role = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *HyperflexNodeAllOf) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeAllOf) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *HyperflexNodeAllOf) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *HyperflexNodeAllOf) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HyperflexNodeAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HyperflexNodeAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HyperflexNodeAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexNodeAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexNodeAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexNodeAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexNodeAllOf) GetCluster() HyperflexClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexNodeAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexNodeAllOf) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster = &v
}

// GetClusterMember returns the ClusterMember field value if set, zero value otherwise.
func (o *HyperflexNodeAllOf) GetClusterMember() AssetClusterMemberRelationship {
	if o == nil || o.ClusterMember == nil {
		var ret AssetClusterMemberRelationship
		return ret
	}
	return *o.ClusterMember
}

// GetClusterMemberOk returns a tuple with the ClusterMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeAllOf) GetClusterMemberOk() (*AssetClusterMemberRelationship, bool) {
	if o == nil || o.ClusterMember == nil {
		return nil, false
	}
	return o.ClusterMember, true
}

// HasClusterMember returns a boolean if a field has been set.
func (o *HyperflexNodeAllOf) HasClusterMember() bool {
	if o != nil && o.ClusterMember != nil {
		return true
	}

	return false
}

// SetClusterMember gets a reference to the given AssetClusterMemberRelationship and assigns it to the ClusterMember field.
func (o *HyperflexNodeAllOf) SetClusterMember(v AssetClusterMemberRelationship) {
	o.ClusterMember = &v
}

// GetPhysicalServer returns the PhysicalServer field value if set, zero value otherwise.
func (o *HyperflexNodeAllOf) GetPhysicalServer() ComputePhysicalRelationship {
	if o == nil || o.PhysicalServer == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.PhysicalServer
}

// GetPhysicalServerOk returns a tuple with the PhysicalServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNodeAllOf) GetPhysicalServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.PhysicalServer == nil {
		return nil, false
	}
	return o.PhysicalServer, true
}

// HasPhysicalServer returns a boolean if a field has been set.
func (o *HyperflexNodeAllOf) HasPhysicalServer() bool {
	if o != nil && o.PhysicalServer != nil {
		return true
	}

	return false
}

// SetPhysicalServer gets a reference to the given ComputePhysicalRelationship and assigns it to the PhysicalServer field.
func (o *HyperflexNodeAllOf) SetPhysicalServer(v ComputePhysicalRelationship) {
	o.PhysicalServer = &v
}

func (o HyperflexNodeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BuildNumber != nil {
		toSerialize["BuildNumber"] = o.BuildNumber
	}
	if o.DisplayVersion != nil {
		toSerialize["DisplayVersion"] = o.DisplayVersion
	}
	if o.HostName != nil {
		toSerialize["HostName"] = o.HostName
	}
	if o.Hypervisor != nil {
		toSerialize["Hypervisor"] = o.Hypervisor
	}
	if o.Identity.IsSet() {
		toSerialize["Identity"] = o.Identity.Get()
	}
	if o.Ip.IsSet() {
		toSerialize["Ip"] = o.Ip.Get()
	}
	if o.Lockdown != nil {
		toSerialize["Lockdown"] = o.Lockdown
	}
	if o.ModelNumber != nil {
		toSerialize["ModelNumber"] = o.ModelNumber
	}
	if o.Role != nil {
		toSerialize["Role"] = o.Role
	}
	if o.SerialNumber != nil {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.ClusterMember != nil {
		toSerialize["ClusterMember"] = o.ClusterMember
	}
	if o.PhysicalServer != nil {
		toSerialize["PhysicalServer"] = o.PhysicalServer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexNodeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexNodeAllOf := _HyperflexNodeAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexNodeAllOf); err == nil {
		*o = HyperflexNodeAllOf(varHyperflexNodeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BuildNumber")
		delete(additionalProperties, "DisplayVersion")
		delete(additionalProperties, "HostName")
		delete(additionalProperties, "Hypervisor")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "Ip")
		delete(additionalProperties, "Lockdown")
		delete(additionalProperties, "ModelNumber")
		delete(additionalProperties, "Role")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "ClusterMember")
		delete(additionalProperties, "PhysicalServer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexNodeAllOf struct {
	value *HyperflexNodeAllOf
	isSet bool
}

func (v NullableHyperflexNodeAllOf) Get() *HyperflexNodeAllOf {
	return v.value
}

func (v *NullableHyperflexNodeAllOf) Set(val *HyperflexNodeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexNodeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexNodeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexNodeAllOf(val *HyperflexNodeAllOf) *NullableHyperflexNodeAllOf {
	return &NullableHyperflexNodeAllOf{value: val, isSet: true}
}

func (v NullableHyperflexNodeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexNodeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
