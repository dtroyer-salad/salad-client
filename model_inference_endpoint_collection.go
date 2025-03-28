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

// checks if the InferenceEndpointCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InferenceEndpointCollection{}

// InferenceEndpointCollection Represents a page from the collection of inference endpoints.
type InferenceEndpointCollection struct {
	// The list of inference endpoints.
	Items []InferenceEndpoint `json:"items"`
	// The page number.
	Page int32 `json:"page"`
	// The maximum number of items per page.
	PageSize int32 `json:"page_size"`
	// The total number of items in the collection.
	TotalSize int32 `json:"total_size"`
}

type _InferenceEndpointCollection InferenceEndpointCollection

// NewInferenceEndpointCollection instantiates a new InferenceEndpointCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInferenceEndpointCollection(items []InferenceEndpoint, page int32, pageSize int32, totalSize int32) *InferenceEndpointCollection {
	this := InferenceEndpointCollection{}
	this.Items = items
	this.Page = page
	this.PageSize = pageSize
	this.TotalSize = totalSize
	return &this
}

// NewInferenceEndpointCollectionWithDefaults instantiates a new InferenceEndpointCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInferenceEndpointCollectionWithDefaults() *InferenceEndpointCollection {
	this := InferenceEndpointCollection{}
	return &this
}

// GetItems returns the Items field value
func (o *InferenceEndpointCollection) GetItems() []InferenceEndpoint {
	if o == nil {
		var ret []InferenceEndpoint
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *InferenceEndpointCollection) GetItemsOk() ([]InferenceEndpoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *InferenceEndpointCollection) SetItems(v []InferenceEndpoint) {
	o.Items = v
}

// GetPage returns the Page field value
func (o *InferenceEndpointCollection) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *InferenceEndpointCollection) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *InferenceEndpointCollection) SetPage(v int32) {
	o.Page = v
}

// GetPageSize returns the PageSize field value
func (o *InferenceEndpointCollection) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *InferenceEndpointCollection) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *InferenceEndpointCollection) SetPageSize(v int32) {
	o.PageSize = v
}

// GetTotalSize returns the TotalSize field value
func (o *InferenceEndpointCollection) GetTotalSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value
// and a boolean to check if the value has been set.
func (o *InferenceEndpointCollection) GetTotalSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSize, true
}

// SetTotalSize sets field value
func (o *InferenceEndpointCollection) SetTotalSize(v int32) {
	o.TotalSize = v
}

func (o InferenceEndpointCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InferenceEndpointCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["page"] = o.Page
	toSerialize["page_size"] = o.PageSize
	toSerialize["total_size"] = o.TotalSize
	return toSerialize, nil
}

func (o *InferenceEndpointCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
		"page",
		"page_size",
		"total_size",
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

	varInferenceEndpointCollection := _InferenceEndpointCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInferenceEndpointCollection)

	if err != nil {
		return err
	}

	*o = InferenceEndpointCollection(varInferenceEndpointCollection)

	return err
}

type NullableInferenceEndpointCollection struct {
	value *InferenceEndpointCollection
	isSet bool
}

func (v NullableInferenceEndpointCollection) Get() *InferenceEndpointCollection {
	return v.value
}

func (v *NullableInferenceEndpointCollection) Set(val *InferenceEndpointCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableInferenceEndpointCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableInferenceEndpointCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInferenceEndpointCollection(val *InferenceEndpointCollection) *NullableInferenceEndpointCollection {
	return &NullableInferenceEndpointCollection{value: val, isSet: true}
}

func (v NullableInferenceEndpointCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInferenceEndpointCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


