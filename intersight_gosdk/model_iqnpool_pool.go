/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4929
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// IqnpoolPool Pool represents a collection of iSCSI Qualified Names (IQNs) for use as initiator identifiers by iSCSI vNICs.
type IqnpoolPool struct {
	PoolAbstractPool
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType      string                  `json:"ObjectType"`
	IqnSuffixBlocks []IqnpoolIqnSuffixBlock `json:"IqnSuffixBlocks,omitempty"`
	// The prefix for any IQN blocks created for this pool. IQN Prefix must have the following format \"iqn.yyyy-mm.naming-authority\", where naming-authority is usually the reverse syntax of the Internet domain name of the naming authority.
	Prefix *string `json:"Prefix,omitempty"`
	// An array of relationships to iqnpoolBlock resources.
	BlockHeads           []IqnpoolBlockRelationship            `json:"BlockHeads,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IqnpoolPool IqnpoolPool

// NewIqnpoolPool instantiates a new IqnpoolPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIqnpoolPool(classId string, objectType string) *IqnpoolPool {
	this := IqnpoolPool{}
	this.ClassId = classId
	this.ObjectType = objectType
	var assignmentOrder string = "sequential"
	this.AssignmentOrder = &assignmentOrder
	return &this
}

// NewIqnpoolPoolWithDefaults instantiates a new IqnpoolPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIqnpoolPoolWithDefaults() *IqnpoolPool {
	this := IqnpoolPool{}
	var classId string = "iqnpool.Pool"
	this.ClassId = classId
	var objectType string = "iqnpool.Pool"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IqnpoolPool) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IqnpoolPool) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IqnpoolPool) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IqnpoolPool) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IqnpoolPool) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IqnpoolPool) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIqnSuffixBlocks returns the IqnSuffixBlocks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IqnpoolPool) GetIqnSuffixBlocks() []IqnpoolIqnSuffixBlock {
	if o == nil {
		var ret []IqnpoolIqnSuffixBlock
		return ret
	}
	return o.IqnSuffixBlocks
}

// GetIqnSuffixBlocksOk returns a tuple with the IqnSuffixBlocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IqnpoolPool) GetIqnSuffixBlocksOk() (*[]IqnpoolIqnSuffixBlock, bool) {
	if o == nil || o.IqnSuffixBlocks == nil {
		return nil, false
	}
	return &o.IqnSuffixBlocks, true
}

// HasIqnSuffixBlocks returns a boolean if a field has been set.
func (o *IqnpoolPool) HasIqnSuffixBlocks() bool {
	if o != nil && o.IqnSuffixBlocks != nil {
		return true
	}

	return false
}

// SetIqnSuffixBlocks gets a reference to the given []IqnpoolIqnSuffixBlock and assigns it to the IqnSuffixBlocks field.
func (o *IqnpoolPool) SetIqnSuffixBlocks(v []IqnpoolIqnSuffixBlock) {
	o.IqnSuffixBlocks = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *IqnpoolPool) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolPool) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *IqnpoolPool) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *IqnpoolPool) SetPrefix(v string) {
	o.Prefix = &v
}

// GetBlockHeads returns the BlockHeads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IqnpoolPool) GetBlockHeads() []IqnpoolBlockRelationship {
	if o == nil {
		var ret []IqnpoolBlockRelationship
		return ret
	}
	return o.BlockHeads
}

// GetBlockHeadsOk returns a tuple with the BlockHeads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IqnpoolPool) GetBlockHeadsOk() (*[]IqnpoolBlockRelationship, bool) {
	if o == nil || o.BlockHeads == nil {
		return nil, false
	}
	return &o.BlockHeads, true
}

// HasBlockHeads returns a boolean if a field has been set.
func (o *IqnpoolPool) HasBlockHeads() bool {
	if o != nil && o.BlockHeads != nil {
		return true
	}

	return false
}

// SetBlockHeads gets a reference to the given []IqnpoolBlockRelationship and assigns it to the BlockHeads field.
func (o *IqnpoolPool) SetBlockHeads(v []IqnpoolBlockRelationship) {
	o.BlockHeads = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *IqnpoolPool) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolPool) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *IqnpoolPool) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *IqnpoolPool) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o IqnpoolPool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractPool, errPoolAbstractPool := json.Marshal(o.PoolAbstractPool)
	if errPoolAbstractPool != nil {
		return []byte{}, errPoolAbstractPool
	}
	errPoolAbstractPool = json.Unmarshal([]byte(serializedPoolAbstractPool), &toSerialize)
	if errPoolAbstractPool != nil {
		return []byte{}, errPoolAbstractPool
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IqnSuffixBlocks != nil {
		toSerialize["IqnSuffixBlocks"] = o.IqnSuffixBlocks
	}
	if o.Prefix != nil {
		toSerialize["Prefix"] = o.Prefix
	}
	if o.BlockHeads != nil {
		toSerialize["BlockHeads"] = o.BlockHeads
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IqnpoolPool) UnmarshalJSON(bytes []byte) (err error) {
	type IqnpoolPoolWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType      string                  `json:"ObjectType"`
		IqnSuffixBlocks []IqnpoolIqnSuffixBlock `json:"IqnSuffixBlocks,omitempty"`
		// The prefix for any IQN blocks created for this pool. IQN Prefix must have the following format \"iqn.yyyy-mm.naming-authority\", where naming-authority is usually the reverse syntax of the Internet domain name of the naming authority.
		Prefix *string `json:"Prefix,omitempty"`
		// An array of relationships to iqnpoolBlock resources.
		BlockHeads   []IqnpoolBlockRelationship            `json:"BlockHeads,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varIqnpoolPoolWithoutEmbeddedStruct := IqnpoolPoolWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIqnpoolPoolWithoutEmbeddedStruct)
	if err == nil {
		varIqnpoolPool := _IqnpoolPool{}
		varIqnpoolPool.ClassId = varIqnpoolPoolWithoutEmbeddedStruct.ClassId
		varIqnpoolPool.ObjectType = varIqnpoolPoolWithoutEmbeddedStruct.ObjectType
		varIqnpoolPool.IqnSuffixBlocks = varIqnpoolPoolWithoutEmbeddedStruct.IqnSuffixBlocks
		varIqnpoolPool.Prefix = varIqnpoolPoolWithoutEmbeddedStruct.Prefix
		varIqnpoolPool.BlockHeads = varIqnpoolPoolWithoutEmbeddedStruct.BlockHeads
		varIqnpoolPool.Organization = varIqnpoolPoolWithoutEmbeddedStruct.Organization
		*o = IqnpoolPool(varIqnpoolPool)
	} else {
		return err
	}

	varIqnpoolPool := _IqnpoolPool{}

	err = json.Unmarshal(bytes, &varIqnpoolPool)
	if err == nil {
		o.PoolAbstractPool = varIqnpoolPool.PoolAbstractPool
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IqnSuffixBlocks")
		delete(additionalProperties, "Prefix")
		delete(additionalProperties, "BlockHeads")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectPoolAbstractPool := reflect.ValueOf(o.PoolAbstractPool)
		for i := 0; i < reflectPoolAbstractPool.Type().NumField(); i++ {
			t := reflectPoolAbstractPool.Type().Field(i)

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

type NullableIqnpoolPool struct {
	value *IqnpoolPool
	isSet bool
}

func (v NullableIqnpoolPool) Get() *IqnpoolPool {
	return v.value
}

func (v *NullableIqnpoolPool) Set(val *IqnpoolPool) {
	v.value = val
	v.isSet = true
}

func (v NullableIqnpoolPool) IsSet() bool {
	return v.isSet
}

func (v *NullableIqnpoolPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIqnpoolPool(val *IqnpoolPool) *NullableIqnpoolPool {
	return &NullableIqnpoolPool{value: val, isSet: true}
}

func (v NullableIqnpoolPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIqnpoolPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
