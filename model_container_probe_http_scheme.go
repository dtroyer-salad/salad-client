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
	"fmt"
)

// ContainerProbeHttpScheme the model 'ContainerProbeHttpScheme'
type ContainerProbeHttpScheme string

// List of ContainerProbeHttpScheme
const (
	HTTP ContainerProbeHttpScheme = "http"
)

// All allowed values of ContainerProbeHttpScheme enum
var AllowedContainerProbeHttpSchemeEnumValues = []ContainerProbeHttpScheme{
	"http",
}

func (v *ContainerProbeHttpScheme) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContainerProbeHttpScheme(value)
	for _, existing := range AllowedContainerProbeHttpSchemeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContainerProbeHttpScheme", value)
}

// NewContainerProbeHttpSchemeFromValue returns a pointer to a valid ContainerProbeHttpScheme
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContainerProbeHttpSchemeFromValue(v string) (*ContainerProbeHttpScheme, error) {
	ev := ContainerProbeHttpScheme(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContainerProbeHttpScheme: valid values are %v", v, AllowedContainerProbeHttpSchemeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContainerProbeHttpScheme) IsValid() bool {
	for _, existing := range AllowedContainerProbeHttpSchemeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContainerProbeHttpScheme value
func (v ContainerProbeHttpScheme) Ptr() *ContainerProbeHttpScheme {
	return &v
}

type NullableContainerProbeHttpScheme struct {
	value *ContainerProbeHttpScheme
	isSet bool
}

func (v NullableContainerProbeHttpScheme) Get() *ContainerProbeHttpScheme {
	return v.value
}

func (v *NullableContainerProbeHttpScheme) Set(val *ContainerProbeHttpScheme) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerProbeHttpScheme) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerProbeHttpScheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerProbeHttpScheme(val *ContainerProbeHttpScheme) *NullableContainerProbeHttpScheme {
	return &NullableContainerProbeHttpScheme{value: val, isSet: true}
}

func (v NullableContainerProbeHttpScheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerProbeHttpScheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

