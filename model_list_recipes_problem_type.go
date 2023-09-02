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

// ListRecipesProblemType the model 'ListRecipesProblemType'
type ListRecipesProblemType string

// List of ListRecipesProblemType
const (
	NULL ListRecipesProblemType = "null"
)

// All allowed values of ListRecipesProblemType enum
var AllowedListRecipesProblemTypeEnumValues = []ListRecipesProblemType{
	"null",
}

func (v *ListRecipesProblemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ListRecipesProblemType(value)
	for _, existing := range AllowedListRecipesProblemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ListRecipesProblemType", value)
}

// NewListRecipesProblemTypeFromValue returns a pointer to a valid ListRecipesProblemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewListRecipesProblemTypeFromValue(v string) (*ListRecipesProblemType, error) {
	ev := ListRecipesProblemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ListRecipesProblemType: valid values are %v", v, AllowedListRecipesProblemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ListRecipesProblemType) IsValid() bool {
	for _, existing := range AllowedListRecipesProblemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListRecipesProblemType value
func (v ListRecipesProblemType) Ptr() *ListRecipesProblemType {
	return &v
}

type NullableListRecipesProblemType struct {
	value *ListRecipesProblemType
	isSet bool
}

func (v NullableListRecipesProblemType) Get() *ListRecipesProblemType {
	return v.value
}

func (v *NullableListRecipesProblemType) Set(val *ListRecipesProblemType) {
	v.value = val
	v.isSet = true
}

func (v NullableListRecipesProblemType) IsSet() bool {
	return v.isSet
}

func (v *NullableListRecipesProblemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRecipesProblemType(val *ListRecipesProblemType) *NullableListRecipesProblemType {
	return &NullableListRecipesProblemType{value: val, isSet: true}
}

func (v NullableListRecipesProblemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRecipesProblemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

