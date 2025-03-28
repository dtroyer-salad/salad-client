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

// checks if the QueuePrototype type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueuePrototype{}

// QueuePrototype Represents a request to create a new queue.
type QueuePrototype struct {
	// The queue name. This must be unique within the project.
	Name string `json:"name" validate:"regexp=^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	// The display name. This may be used as a more human-readable name.
	DisplayName *string `json:"display_name,omitempty" validate:"regexp=^[ ,-.0-9A-Za-z]+$"`
	// The description. This may be used as a space for notes or other information about the queue.
	Description *string `json:"description,omitempty" validate:"regexp=^.*$"`
}

type _QueuePrototype QueuePrototype

// NewQueuePrototype instantiates a new QueuePrototype object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueuePrototype(name string) *QueuePrototype {
	this := QueuePrototype{}
	this.Name = name
	return &this
}

// NewQueuePrototypeWithDefaults instantiates a new QueuePrototype object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueuePrototypeWithDefaults() *QueuePrototype {
	this := QueuePrototype{}
	return &this
}

// GetName returns the Name field value
func (o *QueuePrototype) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *QueuePrototype) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *QueuePrototype) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *QueuePrototype) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueuePrototype) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *QueuePrototype) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *QueuePrototype) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *QueuePrototype) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueuePrototype) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *QueuePrototype) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *QueuePrototype) SetDescription(v string) {
	o.Description = &v
}

func (o QueuePrototype) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueuePrototype) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

func (o *QueuePrototype) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varQueuePrototype := _QueuePrototype{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQueuePrototype)

	if err != nil {
		return err
	}

	*o = QueuePrototype(varQueuePrototype)

	return err
}

type NullableQueuePrototype struct {
	value *QueuePrototype
	isSet bool
}

func (v NullableQueuePrototype) Get() *QueuePrototype {
	return v.value
}

func (v *NullableQueuePrototype) Set(val *QueuePrototype) {
	v.value = val
	v.isSet = true
}

func (v NullableQueuePrototype) IsSet() bool {
	return v.isSet
}

func (v *NullableQueuePrototype) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueuePrototype(val *QueuePrototype) *NullableQueuePrototype {
	return &NullableQueuePrototype{value: val, isSet: true}
}

func (v NullableQueuePrototype) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueuePrototype) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


