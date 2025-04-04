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
	"time"
	"bytes"
	"fmt"
)

// checks if the InferenceEndpointJobEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InferenceEndpointJobEvent{}

// InferenceEndpointJobEvent Represents an event for inference endpoint job
type InferenceEndpointJobEvent struct {
	Action InferenceEndpointJobEventAction `json:"action"`
	// The time the event occurred.
	Time time.Time `json:"time"`
}

type _InferenceEndpointJobEvent InferenceEndpointJobEvent

// NewInferenceEndpointJobEvent instantiates a new InferenceEndpointJobEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInferenceEndpointJobEvent(action InferenceEndpointJobEventAction, time time.Time) *InferenceEndpointJobEvent {
	this := InferenceEndpointJobEvent{}
	this.Action = action
	this.Time = time
	return &this
}

// NewInferenceEndpointJobEventWithDefaults instantiates a new InferenceEndpointJobEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInferenceEndpointJobEventWithDefaults() *InferenceEndpointJobEvent {
	this := InferenceEndpointJobEvent{}
	return &this
}

// GetAction returns the Action field value
func (o *InferenceEndpointJobEvent) GetAction() InferenceEndpointJobEventAction {
	if o == nil {
		var ret InferenceEndpointJobEventAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *InferenceEndpointJobEvent) GetActionOk() (*InferenceEndpointJobEventAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *InferenceEndpointJobEvent) SetAction(v InferenceEndpointJobEventAction) {
	o.Action = v
}

// GetTime returns the Time field value
func (o *InferenceEndpointJobEvent) GetTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *InferenceEndpointJobEvent) GetTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *InferenceEndpointJobEvent) SetTime(v time.Time) {
	o.Time = v
}

func (o InferenceEndpointJobEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InferenceEndpointJobEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["time"] = o.Time
	return toSerialize, nil
}

func (o *InferenceEndpointJobEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"time",
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

	varInferenceEndpointJobEvent := _InferenceEndpointJobEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInferenceEndpointJobEvent)

	if err != nil {
		return err
	}

	*o = InferenceEndpointJobEvent(varInferenceEndpointJobEvent)

	return err
}

type NullableInferenceEndpointJobEvent struct {
	value *InferenceEndpointJobEvent
	isSet bool
}

func (v NullableInferenceEndpointJobEvent) Get() *InferenceEndpointJobEvent {
	return v.value
}

func (v *NullableInferenceEndpointJobEvent) Set(val *InferenceEndpointJobEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableInferenceEndpointJobEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableInferenceEndpointJobEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInferenceEndpointJobEvent(val *InferenceEndpointJobEvent) *NullableInferenceEndpointJobEvent {
	return &NullableInferenceEndpointJobEvent{value: val, isSet: true}
}

func (v NullableInferenceEndpointJobEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInferenceEndpointJobEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


