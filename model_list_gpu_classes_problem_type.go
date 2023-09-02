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

// ListGpuClassesProblemType the model 'ListGpuClassesProblemType'
type ListGpuClassesProblemType string

// List of ListGpuClassesProblemType
const (
	NULL ListGpuClassesProblemType = "null"
)

// All allowed values of ListGpuClassesProblemType enum
var AllowedListGpuClassesProblemTypeEnumValues = []ListGpuClassesProblemType{
	"null",
}

func (v *ListGpuClassesProblemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ListGpuClassesProblemType(value)
	for _, existing := range AllowedListGpuClassesProblemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ListGpuClassesProblemType", value)
}

// NewListGpuClassesProblemTypeFromValue returns a pointer to a valid ListGpuClassesProblemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewListGpuClassesProblemTypeFromValue(v string) (*ListGpuClassesProblemType, error) {
	ev := ListGpuClassesProblemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ListGpuClassesProblemType: valid values are %v", v, AllowedListGpuClassesProblemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ListGpuClassesProblemType) IsValid() bool {
	for _, existing := range AllowedListGpuClassesProblemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListGpuClassesProblemType value
func (v ListGpuClassesProblemType) Ptr() *ListGpuClassesProblemType {
	return &v
}

type NullableListGpuClassesProblemType struct {
	value *ListGpuClassesProblemType
	isSet bool
}

func (v NullableListGpuClassesProblemType) Get() *ListGpuClassesProblemType {
	return v.value
}

func (v *NullableListGpuClassesProblemType) Set(val *ListGpuClassesProblemType) {
	v.value = val
	v.isSet = true
}

func (v NullableListGpuClassesProblemType) IsSet() bool {
	return v.isSet
}

func (v *NullableListGpuClassesProblemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListGpuClassesProblemType(val *ListGpuClassesProblemType) *NullableListGpuClassesProblemType {
	return &NullableListGpuClassesProblemType{value: val, isSet: true}
}

func (v NullableListGpuClassesProblemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListGpuClassesProblemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

