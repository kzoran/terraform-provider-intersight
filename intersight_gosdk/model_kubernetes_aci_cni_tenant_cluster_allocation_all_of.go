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

// KubernetesAciCniTenantClusterAllocationAllOf Definition of the list of properties defined in 'kubernetes.AciCniTenantClusterAllocation', excluding properties defined in parent classes.
type KubernetesAciCniTenantClusterAllocationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// CIDR allocated for ACI node service IPs in this tenant cluster.
	NodeSvcIpSubnet *string `json:"NodeSvcIpSubnet,omitempty"`
	// CIDR allocated for pod IPs in this tenant cluster.
	PodIpSubnet *string `json:"PodIpSubnet,omitempty"`
	// End of VLAN range allocated to this tenant cluster.
	VlanEnd *string `json:"VlanEnd,omitempty"`
	// Start of VLAN range allocated to this tenant cluster.
	VlanStart            *string                               `json:"VlanStart,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesAciCniTenantClusterAllocationAllOf KubernetesAciCniTenantClusterAllocationAllOf

// NewKubernetesAciCniTenantClusterAllocationAllOf instantiates a new KubernetesAciCniTenantClusterAllocationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAciCniTenantClusterAllocationAllOf(classId string, objectType string) *KubernetesAciCniTenantClusterAllocationAllOf {
	this := KubernetesAciCniTenantClusterAllocationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesAciCniTenantClusterAllocationAllOfWithDefaults instantiates a new KubernetesAciCniTenantClusterAllocationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAciCniTenantClusterAllocationAllOfWithDefaults() *KubernetesAciCniTenantClusterAllocationAllOf {
	this := KubernetesAciCniTenantClusterAllocationAllOf{}
	var classId string = "kubernetes.AciCniTenantClusterAllocation"
	this.ClassId = classId
	var objectType string = "kubernetes.AciCniTenantClusterAllocation"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesAciCniTenantClusterAllocationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesAciCniTenantClusterAllocationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetNodeSvcIpSubnet returns the NodeSvcIpSubnet field value if set, zero value otherwise.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetNodeSvcIpSubnet() string {
	if o == nil || o.NodeSvcIpSubnet == nil {
		var ret string
		return ret
	}
	return *o.NodeSvcIpSubnet
}

// GetNodeSvcIpSubnetOk returns a tuple with the NodeSvcIpSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetNodeSvcIpSubnetOk() (*string, bool) {
	if o == nil || o.NodeSvcIpSubnet == nil {
		return nil, false
	}
	return o.NodeSvcIpSubnet, true
}

// HasNodeSvcIpSubnet returns a boolean if a field has been set.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) HasNodeSvcIpSubnet() bool {
	if o != nil && o.NodeSvcIpSubnet != nil {
		return true
	}

	return false
}

// SetNodeSvcIpSubnet gets a reference to the given string and assigns it to the NodeSvcIpSubnet field.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) SetNodeSvcIpSubnet(v string) {
	o.NodeSvcIpSubnet = &v
}

// GetPodIpSubnet returns the PodIpSubnet field value if set, zero value otherwise.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetPodIpSubnet() string {
	if o == nil || o.PodIpSubnet == nil {
		var ret string
		return ret
	}
	return *o.PodIpSubnet
}

// GetPodIpSubnetOk returns a tuple with the PodIpSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetPodIpSubnetOk() (*string, bool) {
	if o == nil || o.PodIpSubnet == nil {
		return nil, false
	}
	return o.PodIpSubnet, true
}

// HasPodIpSubnet returns a boolean if a field has been set.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) HasPodIpSubnet() bool {
	if o != nil && o.PodIpSubnet != nil {
		return true
	}

	return false
}

// SetPodIpSubnet gets a reference to the given string and assigns it to the PodIpSubnet field.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) SetPodIpSubnet(v string) {
	o.PodIpSubnet = &v
}

// GetVlanEnd returns the VlanEnd field value if set, zero value otherwise.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetVlanEnd() string {
	if o == nil || o.VlanEnd == nil {
		var ret string
		return ret
	}
	return *o.VlanEnd
}

// GetVlanEndOk returns a tuple with the VlanEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetVlanEndOk() (*string, bool) {
	if o == nil || o.VlanEnd == nil {
		return nil, false
	}
	return o.VlanEnd, true
}

// HasVlanEnd returns a boolean if a field has been set.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) HasVlanEnd() bool {
	if o != nil && o.VlanEnd != nil {
		return true
	}

	return false
}

// SetVlanEnd gets a reference to the given string and assigns it to the VlanEnd field.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) SetVlanEnd(v string) {
	o.VlanEnd = &v
}

// GetVlanStart returns the VlanStart field value if set, zero value otherwise.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetVlanStart() string {
	if o == nil || o.VlanStart == nil {
		var ret string
		return ret
	}
	return *o.VlanStart
}

// GetVlanStartOk returns a tuple with the VlanStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetVlanStartOk() (*string, bool) {
	if o == nil || o.VlanStart == nil {
		return nil, false
	}
	return o.VlanStart, true
}

// HasVlanStart returns a boolean if a field has been set.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) HasVlanStart() bool {
	if o != nil && o.VlanStart != nil {
		return true
	}

	return false
}

// SetVlanStart gets a reference to the given string and assigns it to the VlanStart field.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) SetVlanStart(v string) {
	o.VlanStart = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesAciCniTenantClusterAllocationAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o KubernetesAciCniTenantClusterAllocationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.NodeSvcIpSubnet != nil {
		toSerialize["NodeSvcIpSubnet"] = o.NodeSvcIpSubnet
	}
	if o.PodIpSubnet != nil {
		toSerialize["PodIpSubnet"] = o.PodIpSubnet
	}
	if o.VlanEnd != nil {
		toSerialize["VlanEnd"] = o.VlanEnd
	}
	if o.VlanStart != nil {
		toSerialize["VlanStart"] = o.VlanStart
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesAciCniTenantClusterAllocationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesAciCniTenantClusterAllocationAllOf := _KubernetesAciCniTenantClusterAllocationAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesAciCniTenantClusterAllocationAllOf); err == nil {
		*o = KubernetesAciCniTenantClusterAllocationAllOf(varKubernetesAciCniTenantClusterAllocationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "NodeSvcIpSubnet")
		delete(additionalProperties, "PodIpSubnet")
		delete(additionalProperties, "VlanEnd")
		delete(additionalProperties, "VlanStart")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesAciCniTenantClusterAllocationAllOf struct {
	value *KubernetesAciCniTenantClusterAllocationAllOf
	isSet bool
}

func (v NullableKubernetesAciCniTenantClusterAllocationAllOf) Get() *KubernetesAciCniTenantClusterAllocationAllOf {
	return v.value
}

func (v *NullableKubernetesAciCniTenantClusterAllocationAllOf) Set(val *KubernetesAciCniTenantClusterAllocationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAciCniTenantClusterAllocationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAciCniTenantClusterAllocationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAciCniTenantClusterAllocationAllOf(val *KubernetesAciCniTenantClusterAllocationAllOf) *NullableKubernetesAciCniTenantClusterAllocationAllOf {
	return &NullableKubernetesAciCniTenantClusterAllocationAllOf{value: val, isSet: true}
}

func (v NullableKubernetesAciCniTenantClusterAllocationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAciCniTenantClusterAllocationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
