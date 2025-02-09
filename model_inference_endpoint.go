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

// checks if the InferenceEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InferenceEndpoint{}

// InferenceEndpoint Represents an inference endpoint
type InferenceEndpoint struct {
	// The unique identifier
	Id string `json:"id"`
	// The inference endpoint name
	Name string `json:"name"`
	// The inference endpoint display name
	DisplayName string `json:"display_name" validate:"regexp=^[ ,-.0-9A-Za-z]+$"`
	// a brief description of the inference endpoint
	Description string `json:"description"`
	// The URL of the inference endpoint
	EndpointUrl string `json:"endpoint_url"`
	// A markdown file containing a detailed description of the inference endpoint
	Readme string `json:"readme"`
	// A description of the price
	PriceDescription string `json:"price_description"`
	// The URL of the icon image
	IconImage string `json:"icon_image"`
}

type _InferenceEndpoint InferenceEndpoint

// NewInferenceEndpoint instantiates a new InferenceEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInferenceEndpoint(id string, name string, displayName string, description string, endpointUrl string, readme string, priceDescription string, iconImage string) *InferenceEndpoint {
	this := InferenceEndpoint{}
	this.Id = id
	this.Name = name
	this.DisplayName = displayName
	this.Description = description
	this.EndpointUrl = endpointUrl
	this.Readme = readme
	this.PriceDescription = priceDescription
	this.IconImage = iconImage
	return &this
}

// NewInferenceEndpointWithDefaults instantiates a new InferenceEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInferenceEndpointWithDefaults() *InferenceEndpoint {
	this := InferenceEndpoint{}
	return &this
}

// GetId returns the Id field value
func (o *InferenceEndpoint) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InferenceEndpoint) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InferenceEndpoint) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *InferenceEndpoint) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InferenceEndpoint) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InferenceEndpoint) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
func (o *InferenceEndpoint) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *InferenceEndpoint) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *InferenceEndpoint) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDescription returns the Description field value
func (o *InferenceEndpoint) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *InferenceEndpoint) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *InferenceEndpoint) SetDescription(v string) {
	o.Description = v
}

// GetEndpointUrl returns the EndpointUrl field value
func (o *InferenceEndpoint) GetEndpointUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointUrl
}

// GetEndpointUrlOk returns a tuple with the EndpointUrl field value
// and a boolean to check if the value has been set.
func (o *InferenceEndpoint) GetEndpointUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointUrl, true
}

// SetEndpointUrl sets field value
func (o *InferenceEndpoint) SetEndpointUrl(v string) {
	o.EndpointUrl = v
}

// GetReadme returns the Readme field value
func (o *InferenceEndpoint) GetReadme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Readme
}

// GetReadmeOk returns a tuple with the Readme field value
// and a boolean to check if the value has been set.
func (o *InferenceEndpoint) GetReadmeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Readme, true
}

// SetReadme sets field value
func (o *InferenceEndpoint) SetReadme(v string) {
	o.Readme = v
}

// GetPriceDescription returns the PriceDescription field value
func (o *InferenceEndpoint) GetPriceDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PriceDescription
}

// GetPriceDescriptionOk returns a tuple with the PriceDescription field value
// and a boolean to check if the value has been set.
func (o *InferenceEndpoint) GetPriceDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceDescription, true
}

// SetPriceDescription sets field value
func (o *InferenceEndpoint) SetPriceDescription(v string) {
	o.PriceDescription = v
}

// GetIconImage returns the IconImage field value
func (o *InferenceEndpoint) GetIconImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IconImage
}

// GetIconImageOk returns a tuple with the IconImage field value
// and a boolean to check if the value has been set.
func (o *InferenceEndpoint) GetIconImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IconImage, true
}

// SetIconImage sets field value
func (o *InferenceEndpoint) SetIconImage(v string) {
	o.IconImage = v
}

func (o InferenceEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InferenceEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["display_name"] = o.DisplayName
	toSerialize["description"] = o.Description
	toSerialize["endpoint_url"] = o.EndpointUrl
	toSerialize["readme"] = o.Readme
	toSerialize["price_description"] = o.PriceDescription
	toSerialize["icon_image"] = o.IconImage
	return toSerialize, nil
}

func (o *InferenceEndpoint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"display_name",
		"description",
		"endpoint_url",
		"readme",
		"price_description",
		"icon_image",
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

	varInferenceEndpoint := _InferenceEndpoint{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInferenceEndpoint)

	if err != nil {
		return err
	}

	*o = InferenceEndpoint(varInferenceEndpoint)

	return err
}

type NullableInferenceEndpoint struct {
	value *InferenceEndpoint
	isSet bool
}

func (v NullableInferenceEndpoint) Get() *InferenceEndpoint {
	return v.value
}

func (v *NullableInferenceEndpoint) Set(val *InferenceEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableInferenceEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableInferenceEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInferenceEndpoint(val *InferenceEndpoint) *NullableInferenceEndpoint {
	return &NullableInferenceEndpoint{value: val, isSet: true}
}

func (v NullableInferenceEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInferenceEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


