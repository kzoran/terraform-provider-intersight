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

// ConnectorCommandControlMessage A Command Message is sent from a cloud service to the connectors command plugin to execute a given command on the platform and begin tunneling input/output to/from the command.
type ConnectorCommandControlMessage struct {
	ConnectorAuthMessage
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The working directory of the command. If empty command is executed in the same directory the device connector process was called.
	Dir *string `json:"Dir,omitempty"`
	// Message carrying the operation to perform.
	MsgType *string `json:"MsgType,omitempty"`
	// The command to execute. Commands must be whitelisted by platform implementation, if a command does not match any whitelisted command patterns an error will be returned to the requesting service on command start.
	Stream *string `json:"Stream,omitempty"`
	// Indicates that a pseudo terminal should be attached to the command. Used for interactive commands. e.g A cross launch cli.
	Terminal *bool `json:"Terminal,omitempty"`
	// The timeout for the command to complete and exit after starting or receiving input. If timeout is not set a default of 10 minutes will be used. If there is input to the command stream the timeout is extended.
	Timeout              *int64 `json:"Timeout,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorCommandControlMessage ConnectorCommandControlMessage

// NewConnectorCommandControlMessage instantiates a new ConnectorCommandControlMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorCommandControlMessage(classId string, objectType string) *ConnectorCommandControlMessage {
	this := ConnectorCommandControlMessage{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorCommandControlMessageWithDefaults instantiates a new ConnectorCommandControlMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorCommandControlMessageWithDefaults() *ConnectorCommandControlMessage {
	this := ConnectorCommandControlMessage{}
	var classId string = "connector.CommandControlMessage"
	this.ClassId = classId
	var objectType string = "connector.CommandControlMessage"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorCommandControlMessage) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorCommandControlMessage) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorCommandControlMessage) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorCommandControlMessage) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorCommandControlMessage) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorCommandControlMessage) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDir returns the Dir field value if set, zero value otherwise.
func (o *ConnectorCommandControlMessage) GetDir() string {
	if o == nil || o.Dir == nil {
		var ret string
		return ret
	}
	return *o.Dir
}

// GetDirOk returns a tuple with the Dir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorCommandControlMessage) GetDirOk() (*string, bool) {
	if o == nil || o.Dir == nil {
		return nil, false
	}
	return o.Dir, true
}

// HasDir returns a boolean if a field has been set.
func (o *ConnectorCommandControlMessage) HasDir() bool {
	if o != nil && o.Dir != nil {
		return true
	}

	return false
}

// SetDir gets a reference to the given string and assigns it to the Dir field.
func (o *ConnectorCommandControlMessage) SetDir(v string) {
	o.Dir = &v
}

// GetMsgType returns the MsgType field value if set, zero value otherwise.
func (o *ConnectorCommandControlMessage) GetMsgType() string {
	if o == nil || o.MsgType == nil {
		var ret string
		return ret
	}
	return *o.MsgType
}

// GetMsgTypeOk returns a tuple with the MsgType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorCommandControlMessage) GetMsgTypeOk() (*string, bool) {
	if o == nil || o.MsgType == nil {
		return nil, false
	}
	return o.MsgType, true
}

// HasMsgType returns a boolean if a field has been set.
func (o *ConnectorCommandControlMessage) HasMsgType() bool {
	if o != nil && o.MsgType != nil {
		return true
	}

	return false
}

// SetMsgType gets a reference to the given string and assigns it to the MsgType field.
func (o *ConnectorCommandControlMessage) SetMsgType(v string) {
	o.MsgType = &v
}

// GetStream returns the Stream field value if set, zero value otherwise.
func (o *ConnectorCommandControlMessage) GetStream() string {
	if o == nil || o.Stream == nil {
		var ret string
		return ret
	}
	return *o.Stream
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorCommandControlMessage) GetStreamOk() (*string, bool) {
	if o == nil || o.Stream == nil {
		return nil, false
	}
	return o.Stream, true
}

// HasStream returns a boolean if a field has been set.
func (o *ConnectorCommandControlMessage) HasStream() bool {
	if o != nil && o.Stream != nil {
		return true
	}

	return false
}

// SetStream gets a reference to the given string and assigns it to the Stream field.
func (o *ConnectorCommandControlMessage) SetStream(v string) {
	o.Stream = &v
}

// GetTerminal returns the Terminal field value if set, zero value otherwise.
func (o *ConnectorCommandControlMessage) GetTerminal() bool {
	if o == nil || o.Terminal == nil {
		var ret bool
		return ret
	}
	return *o.Terminal
}

// GetTerminalOk returns a tuple with the Terminal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorCommandControlMessage) GetTerminalOk() (*bool, bool) {
	if o == nil || o.Terminal == nil {
		return nil, false
	}
	return o.Terminal, true
}

// HasTerminal returns a boolean if a field has been set.
func (o *ConnectorCommandControlMessage) HasTerminal() bool {
	if o != nil && o.Terminal != nil {
		return true
	}

	return false
}

// SetTerminal gets a reference to the given bool and assigns it to the Terminal field.
func (o *ConnectorCommandControlMessage) SetTerminal(v bool) {
	o.Terminal = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *ConnectorCommandControlMessage) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorCommandControlMessage) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *ConnectorCommandControlMessage) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *ConnectorCommandControlMessage) SetTimeout(v int64) {
	o.Timeout = &v
}

func (o ConnectorCommandControlMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorAuthMessage, errConnectorAuthMessage := json.Marshal(o.ConnectorAuthMessage)
	if errConnectorAuthMessage != nil {
		return []byte{}, errConnectorAuthMessage
	}
	errConnectorAuthMessage = json.Unmarshal([]byte(serializedConnectorAuthMessage), &toSerialize)
	if errConnectorAuthMessage != nil {
		return []byte{}, errConnectorAuthMessage
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Dir != nil {
		toSerialize["Dir"] = o.Dir
	}
	if o.MsgType != nil {
		toSerialize["MsgType"] = o.MsgType
	}
	if o.Stream != nil {
		toSerialize["Stream"] = o.Stream
	}
	if o.Terminal != nil {
		toSerialize["Terminal"] = o.Terminal
	}
	if o.Timeout != nil {
		toSerialize["Timeout"] = o.Timeout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorCommandControlMessage) UnmarshalJSON(bytes []byte) (err error) {
	type ConnectorCommandControlMessageWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The working directory of the command. If empty command is executed in the same directory the device connector process was called.
		Dir *string `json:"Dir,omitempty"`
		// Message carrying the operation to perform.
		MsgType *string `json:"MsgType,omitempty"`
		// The command to execute. Commands must be whitelisted by platform implementation, if a command does not match any whitelisted command patterns an error will be returned to the requesting service on command start.
		Stream *string `json:"Stream,omitempty"`
		// Indicates that a pseudo terminal should be attached to the command. Used for interactive commands. e.g A cross launch cli.
		Terminal *bool `json:"Terminal,omitempty"`
		// The timeout for the command to complete and exit after starting or receiving input. If timeout is not set a default of 10 minutes will be used. If there is input to the command stream the timeout is extended.
		Timeout *int64 `json:"Timeout,omitempty"`
	}

	varConnectorCommandControlMessageWithoutEmbeddedStruct := ConnectorCommandControlMessageWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConnectorCommandControlMessageWithoutEmbeddedStruct)
	if err == nil {
		varConnectorCommandControlMessage := _ConnectorCommandControlMessage{}
		varConnectorCommandControlMessage.ClassId = varConnectorCommandControlMessageWithoutEmbeddedStruct.ClassId
		varConnectorCommandControlMessage.ObjectType = varConnectorCommandControlMessageWithoutEmbeddedStruct.ObjectType
		varConnectorCommandControlMessage.Dir = varConnectorCommandControlMessageWithoutEmbeddedStruct.Dir
		varConnectorCommandControlMessage.MsgType = varConnectorCommandControlMessageWithoutEmbeddedStruct.MsgType
		varConnectorCommandControlMessage.Stream = varConnectorCommandControlMessageWithoutEmbeddedStruct.Stream
		varConnectorCommandControlMessage.Terminal = varConnectorCommandControlMessageWithoutEmbeddedStruct.Terminal
		varConnectorCommandControlMessage.Timeout = varConnectorCommandControlMessageWithoutEmbeddedStruct.Timeout
		*o = ConnectorCommandControlMessage(varConnectorCommandControlMessage)
	} else {
		return err
	}

	varConnectorCommandControlMessage := _ConnectorCommandControlMessage{}

	err = json.Unmarshal(bytes, &varConnectorCommandControlMessage)
	if err == nil {
		o.ConnectorAuthMessage = varConnectorCommandControlMessage.ConnectorAuthMessage
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dir")
		delete(additionalProperties, "MsgType")
		delete(additionalProperties, "Stream")
		delete(additionalProperties, "Terminal")
		delete(additionalProperties, "Timeout")

		// remove fields from embedded structs
		reflectConnectorAuthMessage := reflect.ValueOf(o.ConnectorAuthMessage)
		for i := 0; i < reflectConnectorAuthMessage.Type().NumField(); i++ {
			t := reflectConnectorAuthMessage.Type().Field(i)

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

type NullableConnectorCommandControlMessage struct {
	value *ConnectorCommandControlMessage
	isSet bool
}

func (v NullableConnectorCommandControlMessage) Get() *ConnectorCommandControlMessage {
	return v.value
}

func (v *NullableConnectorCommandControlMessage) Set(val *ConnectorCommandControlMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorCommandControlMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorCommandControlMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorCommandControlMessage(val *ConnectorCommandControlMessage) *NullableConnectorCommandControlMessage {
	return &NullableConnectorCommandControlMessage{value: val, isSet: true}
}

func (v NullableConnectorCommandControlMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorCommandControlMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
