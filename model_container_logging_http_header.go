/*
SaladCloud API

The SaladCloud REST API. Please refer to the [SaladCloud API Documentation](https://docs.salad.com/api-reference) for more details.

API version: 0.9.0-alpha.11
Contact: cloud@salad.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package saladclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ContainerLoggingHttpHeader type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerLoggingHttpHeader{}

// ContainerLoggingHttpHeader Represents an HTTP header used for container logging configuration.
type ContainerLoggingHttpHeader struct {
	// The name of the HTTP header
	Name string `json:"name" validate:"regexp=^.*$"`
	// The value of the HTTP header
	Value string `json:"value" validate:"regexp=^.*$"`
}

type _ContainerLoggingHttpHeader ContainerLoggingHttpHeader

// NewContainerLoggingHttpHeader instantiates a new ContainerLoggingHttpHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerLoggingHttpHeader(name string, value string) *ContainerLoggingHttpHeader {
	this := ContainerLoggingHttpHeader{}
	this.Name = name
	this.Value = value
	return &this
}

// NewContainerLoggingHttpHeaderWithDefaults instantiates a new ContainerLoggingHttpHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerLoggingHttpHeaderWithDefaults() *ContainerLoggingHttpHeader {
	this := ContainerLoggingHttpHeader{}
	return &this
}

// GetName returns the Name field value
func (o *ContainerLoggingHttpHeader) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerLoggingHttpHeader) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerLoggingHttpHeader) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *ContainerLoggingHttpHeader) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ContainerLoggingHttpHeader) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ContainerLoggingHttpHeader) SetValue(v string) {
	o.Value = v
}

func (o ContainerLoggingHttpHeader) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerLoggingHttpHeader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *ContainerLoggingHttpHeader) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varContainerLoggingHttpHeader := _ContainerLoggingHttpHeader{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerLoggingHttpHeader)

	if err != nil {
		return err
	}

	*o = ContainerLoggingHttpHeader(varContainerLoggingHttpHeader)

	return err
}

type NullableContainerLoggingHttpHeader struct {
	value *ContainerLoggingHttpHeader
	isSet bool
}

func (v NullableContainerLoggingHttpHeader) Get() *ContainerLoggingHttpHeader {
	return v.value
}

func (v *NullableContainerLoggingHttpHeader) Set(val *ContainerLoggingHttpHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerLoggingHttpHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerLoggingHttpHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerLoggingHttpHeader(val *ContainerLoggingHttpHeader) *NullableContainerLoggingHttpHeader {
	return &NullableContainerLoggingHttpHeader{value: val, isSet: true}
}

func (v NullableContainerLoggingHttpHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerLoggingHttpHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


