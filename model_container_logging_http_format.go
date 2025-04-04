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
	"fmt"
)

// ContainerLoggingHttpFormat The format in which logs will be delivered
type ContainerLoggingHttpFormat string

// List of ContainerLoggingHttpFormat
const (
	CONTAINERLOGGINGHTTPFORMAT_JSON ContainerLoggingHttpFormat = "json"
	CONTAINERLOGGINGHTTPFORMAT_JSON_LINES ContainerLoggingHttpFormat = "json_lines"
)

// All allowed values of ContainerLoggingHttpFormat enum
var AllowedContainerLoggingHttpFormatEnumValues = []ContainerLoggingHttpFormat{
	"json",
	"json_lines",
}

func (v *ContainerLoggingHttpFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContainerLoggingHttpFormat(value)
	for _, existing := range AllowedContainerLoggingHttpFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContainerLoggingHttpFormat", value)
}

// NewContainerLoggingHttpFormatFromValue returns a pointer to a valid ContainerLoggingHttpFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContainerLoggingHttpFormatFromValue(v string) (*ContainerLoggingHttpFormat, error) {
	ev := ContainerLoggingHttpFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContainerLoggingHttpFormat: valid values are %v", v, AllowedContainerLoggingHttpFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContainerLoggingHttpFormat) IsValid() bool {
	for _, existing := range AllowedContainerLoggingHttpFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContainerLoggingHttpFormat value
func (v ContainerLoggingHttpFormat) Ptr() *ContainerLoggingHttpFormat {
	return &v
}

type NullableContainerLoggingHttpFormat struct {
	value *ContainerLoggingHttpFormat
	isSet bool
}

func (v NullableContainerLoggingHttpFormat) Get() *ContainerLoggingHttpFormat {
	return v.value
}

func (v *NullableContainerLoggingHttpFormat) Set(val *ContainerLoggingHttpFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerLoggingHttpFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerLoggingHttpFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerLoggingHttpFormat(val *ContainerLoggingHttpFormat) *NullableContainerLoggingHttpFormat {
	return &NullableContainerLoggingHttpFormat{value: val, isSet: true}
}

func (v NullableContainerLoggingHttpFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerLoggingHttpFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

