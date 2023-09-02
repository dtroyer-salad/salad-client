/*
SaladCloud Public API

The SaladCloud Public API.

API version: 1.0.0-alpha.56
Contact: support@salad.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package salad-client

import (
	"encoding/json"
)

// checks if the ContainerGroupProbeExec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerGroupProbeExec{}

// ContainerGroupProbeExec struct for ContainerGroupProbeExec
type ContainerGroupProbeExec struct {
	Command []string `json:"command"`
}

// NewContainerGroupProbeExec instantiates a new ContainerGroupProbeExec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerGroupProbeExec(command []string) *ContainerGroupProbeExec {
	this := ContainerGroupProbeExec{}
	this.Command = command
	return &this
}

// NewContainerGroupProbeExecWithDefaults instantiates a new ContainerGroupProbeExec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerGroupProbeExecWithDefaults() *ContainerGroupProbeExec {
	this := ContainerGroupProbeExec{}
	return &this
}

// GetCommand returns the Command field value
func (o *ContainerGroupProbeExec) GetCommand() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbeExec) GetCommandOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Command, true
}

// SetCommand sets field value
func (o *ContainerGroupProbeExec) SetCommand(v []string) {
	o.Command = v
}

func (o ContainerGroupProbeExec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerGroupProbeExec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["command"] = o.Command
	return toSerialize, nil
}

type NullableContainerGroupProbeExec struct {
	value *ContainerGroupProbeExec
	isSet bool
}

func (v NullableContainerGroupProbeExec) Get() *ContainerGroupProbeExec {
	return v.value
}

func (v *NullableContainerGroupProbeExec) Set(val *ContainerGroupProbeExec) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerGroupProbeExec) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerGroupProbeExec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerGroupProbeExec(val *ContainerGroupProbeExec) *NullableContainerGroupProbeExec {
	return &NullableContainerGroupProbeExec{value: val, isSet: true}
}

func (v NullableContainerGroupProbeExec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerGroupProbeExec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


