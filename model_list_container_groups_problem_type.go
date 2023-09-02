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

// ListContainerGroupsProblemType the model 'ListContainerGroupsProblemType'
type ListContainerGroupsProblemType string

// List of ListContainerGroupsProblemType
const (
	NULL ListContainerGroupsProblemType = "null"
)

// All allowed values of ListContainerGroupsProblemType enum
var AllowedListContainerGroupsProblemTypeEnumValues = []ListContainerGroupsProblemType{
	"null",
}

func (v *ListContainerGroupsProblemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ListContainerGroupsProblemType(value)
	for _, existing := range AllowedListContainerGroupsProblemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ListContainerGroupsProblemType", value)
}

// NewListContainerGroupsProblemTypeFromValue returns a pointer to a valid ListContainerGroupsProblemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewListContainerGroupsProblemTypeFromValue(v string) (*ListContainerGroupsProblemType, error) {
	ev := ListContainerGroupsProblemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ListContainerGroupsProblemType: valid values are %v", v, AllowedListContainerGroupsProblemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ListContainerGroupsProblemType) IsValid() bool {
	for _, existing := range AllowedListContainerGroupsProblemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListContainerGroupsProblemType value
func (v ListContainerGroupsProblemType) Ptr() *ListContainerGroupsProblemType {
	return &v
}

type NullableListContainerGroupsProblemType struct {
	value *ListContainerGroupsProblemType
	isSet bool
}

func (v NullableListContainerGroupsProblemType) Get() *ListContainerGroupsProblemType {
	return v.value
}

func (v *NullableListContainerGroupsProblemType) Set(val *ListContainerGroupsProblemType) {
	v.value = val
	v.isSet = true
}

func (v NullableListContainerGroupsProblemType) IsSet() bool {
	return v.isSet
}

func (v *NullableListContainerGroupsProblemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListContainerGroupsProblemType(val *ListContainerGroupsProblemType) *NullableListContainerGroupsProblemType {
	return &NullableListContainerGroupsProblemType{value: val, isSet: true}
}

func (v NullableListContainerGroupsProblemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListContainerGroupsProblemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

