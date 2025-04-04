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

// checks if the ContainerGroupQueueAutoscaler type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerGroupQueueAutoscaler{}

// ContainerGroupQueueAutoscaler Defines configuration for automatically scaling container instances based on queue length. The autoscaler monitors a queue and adjusts the number of running replicas to maintain the desired queue length.
type ContainerGroupQueueAutoscaler struct {
	// The target number of items in the queue that the autoscaler attempts to maintain by scaling the containers up or down
	DesiredQueueLength int32 `json:"desired_queue_length"`
	// The maximum number of instances the container can scale up to
	MaxReplicas int32 `json:"max_replicas"`
	// The maximum number of instances that can be removed per minute to prevent rapid downscaling
	MaxDownscalePerMinute *int32 `json:"max_downscale_per_minute,omitempty"`
	// The maximum number of instances that can be added per minute to prevent rapid upscaling
	MaxUpscalePerMinute *int32 `json:"max_upscale_per_minute,omitempty"`
	// The minimum number of instances the container can scale down to, ensuring baseline availability
	MinReplicas int32 `json:"min_replicas"`
	// The period (in seconds) in which the autoscaler checks the queue length and applies the scaling formula
	PollingPeriod *int32 `json:"polling_period,omitempty"`
}

type _ContainerGroupQueueAutoscaler ContainerGroupQueueAutoscaler

// NewContainerGroupQueueAutoscaler instantiates a new ContainerGroupQueueAutoscaler object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerGroupQueueAutoscaler(desiredQueueLength int32, maxReplicas int32, minReplicas int32) *ContainerGroupQueueAutoscaler {
	this := ContainerGroupQueueAutoscaler{}
	this.DesiredQueueLength = desiredQueueLength
	this.MaxReplicas = maxReplicas
	this.MinReplicas = minReplicas
	return &this
}

// NewContainerGroupQueueAutoscalerWithDefaults instantiates a new ContainerGroupQueueAutoscaler object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerGroupQueueAutoscalerWithDefaults() *ContainerGroupQueueAutoscaler {
	this := ContainerGroupQueueAutoscaler{}
	return &this
}

// GetDesiredQueueLength returns the DesiredQueueLength field value
func (o *ContainerGroupQueueAutoscaler) GetDesiredQueueLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DesiredQueueLength
}

// GetDesiredQueueLengthOk returns a tuple with the DesiredQueueLength field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupQueueAutoscaler) GetDesiredQueueLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DesiredQueueLength, true
}

// SetDesiredQueueLength sets field value
func (o *ContainerGroupQueueAutoscaler) SetDesiredQueueLength(v int32) {
	o.DesiredQueueLength = v
}

// GetMaxReplicas returns the MaxReplicas field value
func (o *ContainerGroupQueueAutoscaler) GetMaxReplicas() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxReplicas
}

// GetMaxReplicasOk returns a tuple with the MaxReplicas field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupQueueAutoscaler) GetMaxReplicasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxReplicas, true
}

// SetMaxReplicas sets field value
func (o *ContainerGroupQueueAutoscaler) SetMaxReplicas(v int32) {
	o.MaxReplicas = v
}

// GetMaxDownscalePerMinute returns the MaxDownscalePerMinute field value if set, zero value otherwise.
func (o *ContainerGroupQueueAutoscaler) GetMaxDownscalePerMinute() int32 {
	if o == nil || IsNil(o.MaxDownscalePerMinute) {
		var ret int32
		return ret
	}
	return *o.MaxDownscalePerMinute
}

// GetMaxDownscalePerMinuteOk returns a tuple with the MaxDownscalePerMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroupQueueAutoscaler) GetMaxDownscalePerMinuteOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxDownscalePerMinute) {
		return nil, false
	}
	return o.MaxDownscalePerMinute, true
}

// HasMaxDownscalePerMinute returns a boolean if a field has been set.
func (o *ContainerGroupQueueAutoscaler) HasMaxDownscalePerMinute() bool {
	if o != nil && !IsNil(o.MaxDownscalePerMinute) {
		return true
	}

	return false
}

// SetMaxDownscalePerMinute gets a reference to the given int32 and assigns it to the MaxDownscalePerMinute field.
func (o *ContainerGroupQueueAutoscaler) SetMaxDownscalePerMinute(v int32) {
	o.MaxDownscalePerMinute = &v
}

// GetMaxUpscalePerMinute returns the MaxUpscalePerMinute field value if set, zero value otherwise.
func (o *ContainerGroupQueueAutoscaler) GetMaxUpscalePerMinute() int32 {
	if o == nil || IsNil(o.MaxUpscalePerMinute) {
		var ret int32
		return ret
	}
	return *o.MaxUpscalePerMinute
}

// GetMaxUpscalePerMinuteOk returns a tuple with the MaxUpscalePerMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroupQueueAutoscaler) GetMaxUpscalePerMinuteOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxUpscalePerMinute) {
		return nil, false
	}
	return o.MaxUpscalePerMinute, true
}

// HasMaxUpscalePerMinute returns a boolean if a field has been set.
func (o *ContainerGroupQueueAutoscaler) HasMaxUpscalePerMinute() bool {
	if o != nil && !IsNil(o.MaxUpscalePerMinute) {
		return true
	}

	return false
}

// SetMaxUpscalePerMinute gets a reference to the given int32 and assigns it to the MaxUpscalePerMinute field.
func (o *ContainerGroupQueueAutoscaler) SetMaxUpscalePerMinute(v int32) {
	o.MaxUpscalePerMinute = &v
}

// GetMinReplicas returns the MinReplicas field value
func (o *ContainerGroupQueueAutoscaler) GetMinReplicas() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinReplicas
}

// GetMinReplicasOk returns a tuple with the MinReplicas field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupQueueAutoscaler) GetMinReplicasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinReplicas, true
}

// SetMinReplicas sets field value
func (o *ContainerGroupQueueAutoscaler) SetMinReplicas(v int32) {
	o.MinReplicas = v
}

// GetPollingPeriod returns the PollingPeriod field value if set, zero value otherwise.
func (o *ContainerGroupQueueAutoscaler) GetPollingPeriod() int32 {
	if o == nil || IsNil(o.PollingPeriod) {
		var ret int32
		return ret
	}
	return *o.PollingPeriod
}

// GetPollingPeriodOk returns a tuple with the PollingPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroupQueueAutoscaler) GetPollingPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.PollingPeriod) {
		return nil, false
	}
	return o.PollingPeriod, true
}

// HasPollingPeriod returns a boolean if a field has been set.
func (o *ContainerGroupQueueAutoscaler) HasPollingPeriod() bool {
	if o != nil && !IsNil(o.PollingPeriod) {
		return true
	}

	return false
}

// SetPollingPeriod gets a reference to the given int32 and assigns it to the PollingPeriod field.
func (o *ContainerGroupQueueAutoscaler) SetPollingPeriod(v int32) {
	o.PollingPeriod = &v
}

func (o ContainerGroupQueueAutoscaler) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerGroupQueueAutoscaler) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["desired_queue_length"] = o.DesiredQueueLength
	toSerialize["max_replicas"] = o.MaxReplicas
	if !IsNil(o.MaxDownscalePerMinute) {
		toSerialize["max_downscale_per_minute"] = o.MaxDownscalePerMinute
	}
	if !IsNil(o.MaxUpscalePerMinute) {
		toSerialize["max_upscale_per_minute"] = o.MaxUpscalePerMinute
	}
	toSerialize["min_replicas"] = o.MinReplicas
	if !IsNil(o.PollingPeriod) {
		toSerialize["polling_period"] = o.PollingPeriod
	}
	return toSerialize, nil
}

func (o *ContainerGroupQueueAutoscaler) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"desired_queue_length",
		"max_replicas",
		"min_replicas",
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

	varContainerGroupQueueAutoscaler := _ContainerGroupQueueAutoscaler{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerGroupQueueAutoscaler)

	if err != nil {
		return err
	}

	*o = ContainerGroupQueueAutoscaler(varContainerGroupQueueAutoscaler)

	return err
}

type NullableContainerGroupQueueAutoscaler struct {
	value *ContainerGroupQueueAutoscaler
	isSet bool
}

func (v NullableContainerGroupQueueAutoscaler) Get() *ContainerGroupQueueAutoscaler {
	return v.value
}

func (v *NullableContainerGroupQueueAutoscaler) Set(val *ContainerGroupQueueAutoscaler) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerGroupQueueAutoscaler) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerGroupQueueAutoscaler) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerGroupQueueAutoscaler(val *ContainerGroupQueueAutoscaler) *NullableContainerGroupQueueAutoscaler {
	return &NullableContainerGroupQueueAutoscaler{value: val, isSet: true}
}

func (v NullableContainerGroupQueueAutoscaler) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerGroupQueueAutoscaler) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


