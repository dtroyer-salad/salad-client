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

// checks if the WorkloadErrorList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkloadErrorList{}

// WorkloadErrorList Represents a list of workload errors
type WorkloadErrorList struct {
	// A list of workload errors
	Items []WorkloadError `json:"items"`
}

type _WorkloadErrorList WorkloadErrorList

// NewWorkloadErrorList instantiates a new WorkloadErrorList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkloadErrorList(items []WorkloadError) *WorkloadErrorList {
	this := WorkloadErrorList{}
	this.Items = items
	return &this
}

// NewWorkloadErrorListWithDefaults instantiates a new WorkloadErrorList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkloadErrorListWithDefaults() *WorkloadErrorList {
	this := WorkloadErrorList{}
	return &this
}

// GetItems returns the Items field value
func (o *WorkloadErrorList) GetItems() []WorkloadError {
	if o == nil {
		var ret []WorkloadError
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *WorkloadErrorList) GetItemsOk() ([]WorkloadError, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *WorkloadErrorList) SetItems(v []WorkloadError) {
	o.Items = v
}

func (o WorkloadErrorList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkloadErrorList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *WorkloadErrorList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
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

	varWorkloadErrorList := _WorkloadErrorList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkloadErrorList)

	if err != nil {
		return err
	}

	*o = WorkloadErrorList(varWorkloadErrorList)

	return err
}

type NullableWorkloadErrorList struct {
	value *WorkloadErrorList
	isSet bool
}

func (v NullableWorkloadErrorList) Get() *WorkloadErrorList {
	return v.value
}

func (v *NullableWorkloadErrorList) Set(val *WorkloadErrorList) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkloadErrorList) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkloadErrorList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkloadErrorList(val *WorkloadErrorList) *NullableWorkloadErrorList {
	return &NullableWorkloadErrorList{value: val, isSet: true}
}

func (v NullableWorkloadErrorList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkloadErrorList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


