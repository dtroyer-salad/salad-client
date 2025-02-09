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

// checks if the HttpHeadersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HttpHeadersInner{}

// HttpHeadersInner struct for HttpHeadersInner
type HttpHeadersInner struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type _HttpHeadersInner HttpHeadersInner

// NewHttpHeadersInner instantiates a new HttpHeadersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpHeadersInner(name string, value string) *HttpHeadersInner {
	this := HttpHeadersInner{}
	this.Name = name
	this.Value = value
	return &this
}

// NewHttpHeadersInnerWithDefaults instantiates a new HttpHeadersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpHeadersInnerWithDefaults() *HttpHeadersInner {
	this := HttpHeadersInner{}
	return &this
}

// GetName returns the Name field value
func (o *HttpHeadersInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HttpHeadersInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HttpHeadersInner) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *HttpHeadersInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *HttpHeadersInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *HttpHeadersInner) SetValue(v string) {
	o.Value = v
}

func (o HttpHeadersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpHeadersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *HttpHeadersInner) UnmarshalJSON(data []byte) (err error) {
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

	varHttpHeadersInner := _HttpHeadersInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHttpHeadersInner)

	if err != nil {
		return err
	}

	*o = HttpHeadersInner(varHttpHeadersInner)

	return err
}

type NullableHttpHeadersInner struct {
	value *HttpHeadersInner
	isSet bool
}

func (v NullableHttpHeadersInner) Get() *HttpHeadersInner {
	return v.value
}

func (v *NullableHttpHeadersInner) Set(val *HttpHeadersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpHeadersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpHeadersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpHeadersInner(val *HttpHeadersInner) *NullableHttpHeadersInner {
	return &NullableHttpHeadersInner{value: val, isSet: true}
}

func (v NullableHttpHeadersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpHeadersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


