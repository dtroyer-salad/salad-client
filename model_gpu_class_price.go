/*
SaladCloud API

The SaladCloud REST API. Please refer to the [SaladCloud API Documentation](https://docs.salad.com/api-reference) for more details.

API version: 0.9.0-alpha.6
Contact: cloud@salad.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package saladclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GpuClassPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GpuClassPrice{}

// GpuClassPrice Represents the price of a GPU class for a given container group priority
type GpuClassPrice struct {
	Priority NullableContainerGroupPriority `json:"priority"`
	// The price
	Price string `json:"price"`
}

type _GpuClassPrice GpuClassPrice

// NewGpuClassPrice instantiates a new GpuClassPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGpuClassPrice(priority NullableContainerGroupPriority, price string) *GpuClassPrice {
	this := GpuClassPrice{}
	this.Priority = priority
	this.Price = price
	return &this
}

// NewGpuClassPriceWithDefaults instantiates a new GpuClassPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGpuClassPriceWithDefaults() *GpuClassPrice {
	this := GpuClassPrice{}
	return &this
}

// GetPriority returns the Priority field value
// If the value is explicit nil, the zero value for ContainerGroupPriority will be returned
func (o *GpuClassPrice) GetPriority() ContainerGroupPriority {
	if o == nil || o.Priority.Get() == nil {
		var ret ContainerGroupPriority
		return ret
	}

	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GpuClassPrice) GetPriorityOk() (*ContainerGroupPriority, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// SetPriority sets field value
func (o *GpuClassPrice) SetPriority(v ContainerGroupPriority) {
	o.Priority.Set(&v)
}

// GetPrice returns the Price field value
func (o *GpuClassPrice) GetPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *GpuClassPrice) GetPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *GpuClassPrice) SetPrice(v string) {
	o.Price = v
}

func (o GpuClassPrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GpuClassPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["priority"] = o.Priority.Get()
	toSerialize["price"] = o.Price
	return toSerialize, nil
}

func (o *GpuClassPrice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"priority",
		"price",
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

	varGpuClassPrice := _GpuClassPrice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGpuClassPrice)

	if err != nil {
		return err
	}

	*o = GpuClassPrice(varGpuClassPrice)

	return err
}

type NullableGpuClassPrice struct {
	value *GpuClassPrice
	isSet bool
}

func (v NullableGpuClassPrice) Get() *GpuClassPrice {
	return v.value
}

func (v *NullableGpuClassPrice) Set(val *GpuClassPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableGpuClassPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableGpuClassPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGpuClassPrice(val *GpuClassPrice) *NullableGpuClassPrice {
	return &NullableGpuClassPrice{value: val, isSet: true}
}

func (v NullableGpuClassPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGpuClassPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


