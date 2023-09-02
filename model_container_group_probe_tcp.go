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

// checks if the ContainerGroupProbeTcp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerGroupProbeTcp{}

// ContainerGroupProbeTcp struct for ContainerGroupProbeTcp
type ContainerGroupProbeTcp struct {
	Port int32 `json:"port"`
}

// NewContainerGroupProbeTcp instantiates a new ContainerGroupProbeTcp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerGroupProbeTcp(port int32) *ContainerGroupProbeTcp {
	this := ContainerGroupProbeTcp{}
	this.Port = port
	return &this
}

// NewContainerGroupProbeTcpWithDefaults instantiates a new ContainerGroupProbeTcp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerGroupProbeTcpWithDefaults() *ContainerGroupProbeTcp {
	this := ContainerGroupProbeTcp{}
	return &this
}

// GetPort returns the Port field value
func (o *ContainerGroupProbeTcp) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbeTcp) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *ContainerGroupProbeTcp) SetPort(v int32) {
	o.Port = v
}

func (o ContainerGroupProbeTcp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerGroupProbeTcp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["port"] = o.Port
	return toSerialize, nil
}

type NullableContainerGroupProbeTcp struct {
	value *ContainerGroupProbeTcp
	isSet bool
}

func (v NullableContainerGroupProbeTcp) Get() *ContainerGroupProbeTcp {
	return v.value
}

func (v *NullableContainerGroupProbeTcp) Set(val *ContainerGroupProbeTcp) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerGroupProbeTcp) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerGroupProbeTcp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerGroupProbeTcp(val *ContainerGroupProbeTcp) *NullableContainerGroupProbeTcp {
	return &NullableContainerGroupProbeTcp{value: val, isSet: true}
}

func (v NullableContainerGroupProbeTcp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerGroupProbeTcp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


