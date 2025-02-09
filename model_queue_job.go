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
	"time"
	"bytes"
	"fmt"
)

// checks if the QueueJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueueJob{}

// QueueJob Represents a queue job
type QueueJob struct {
	Id string `json:"id"`
	// The job input. May be any valid JSON.
	Input interface{} `json:"input"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Webhook NullableString `json:"webhook,omitempty"`
	Status string `json:"status"`
	Events []QueueJobEvent `json:"events"`
	// The job output. May be any valid JSON.
	Output interface{} `json:"output,omitempty"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type _QueueJob QueueJob

// NewQueueJob instantiates a new QueueJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueueJob(id string, input interface{}, status string, events []QueueJobEvent, createTime time.Time, updateTime time.Time) *QueueJob {
	this := QueueJob{}
	this.Id = id
	this.Input = input
	this.Status = status
	this.Events = events
	this.CreateTime = createTime
	this.UpdateTime = updateTime
	return &this
}

// NewQueueJobWithDefaults instantiates a new QueueJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueueJobWithDefaults() *QueueJob {
	this := QueueJob{}
	return &this
}

// GetId returns the Id field value
func (o *QueueJob) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QueueJob) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QueueJob) SetId(v string) {
	o.Id = v
}

// GetInput returns the Input field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *QueueJob) GetInput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueueJob) GetInputOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Input) {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *QueueJob) SetInput(v interface{}) {
	o.Input = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueueJob) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueueJob) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *QueueJob) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *QueueJob) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueueJob) GetWebhook() string {
	if o == nil || IsNil(o.Webhook.Get()) {
		var ret string
		return ret
	}
	return *o.Webhook.Get()
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueueJob) GetWebhookOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Webhook.Get(), o.Webhook.IsSet()
}

// HasWebhook returns a boolean if a field has been set.
func (o *QueueJob) HasWebhook() bool {
	if o != nil && o.Webhook.IsSet() {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given NullableString and assigns it to the Webhook field.
func (o *QueueJob) SetWebhook(v string) {
	o.Webhook.Set(&v)
}
// SetWebhookNil sets the value for Webhook to be an explicit nil
func (o *QueueJob) SetWebhookNil() {
	o.Webhook.Set(nil)
}

// UnsetWebhook ensures that no value is present for Webhook, not even an explicit nil
func (o *QueueJob) UnsetWebhook() {
	o.Webhook.Unset()
}

// GetStatus returns the Status field value
func (o *QueueJob) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *QueueJob) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *QueueJob) SetStatus(v string) {
	o.Status = v
}

// GetEvents returns the Events field value
func (o *QueueJob) GetEvents() []QueueJobEvent {
	if o == nil {
		var ret []QueueJobEvent
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *QueueJob) GetEventsOk() ([]QueueJobEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *QueueJob) SetEvents(v []QueueJobEvent) {
	o.Events = v
}

// GetOutput returns the Output field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueueJob) GetOutput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueueJob) GetOutputOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Output) {
		return nil, false
	}
	return &o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *QueueJob) HasOutput() bool {
	if o != nil && !IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given interface{} and assigns it to the Output field.
func (o *QueueJob) SetOutput(v interface{}) {
	o.Output = v
}

// GetCreateTime returns the CreateTime field value
func (o *QueueJob) GetCreateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value
// and a boolean to check if the value has been set.
func (o *QueueJob) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateTime, true
}

// SetCreateTime sets field value
func (o *QueueJob) SetCreateTime(v time.Time) {
	o.CreateTime = v
}

// GetUpdateTime returns the UpdateTime field value
func (o *QueueJob) GetUpdateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value
// and a boolean to check if the value has been set.
func (o *QueueJob) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateTime, true
}

// SetUpdateTime sets field value
func (o *QueueJob) SetUpdateTime(v time.Time) {
	o.UpdateTime = v
}

func (o QueueJob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueueJob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Input != nil {
		toSerialize["input"] = o.Input
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Webhook.IsSet() {
		toSerialize["webhook"] = o.Webhook.Get()
	}
	toSerialize["status"] = o.Status
	toSerialize["events"] = o.Events
	if o.Output != nil {
		toSerialize["output"] = o.Output
	}
	toSerialize["create_time"] = o.CreateTime
	toSerialize["update_time"] = o.UpdateTime
	return toSerialize, nil
}

func (o *QueueJob) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"input",
		"status",
		"events",
		"create_time",
		"update_time",
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

	varQueueJob := _QueueJob{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQueueJob)

	if err != nil {
		return err
	}

	*o = QueueJob(varQueueJob)

	return err
}

type NullableQueueJob struct {
	value *QueueJob
	isSet bool
}

func (v NullableQueueJob) Get() *QueueJob {
	return v.value
}

func (v *NullableQueueJob) Set(val *QueueJob) {
	v.value = val
	v.isSet = true
}

func (v NullableQueueJob) IsSet() bool {
	return v.isSet
}

func (v *NullableQueueJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueueJob(val *QueueJob) *NullableQueueJob {
	return &NullableQueueJob{value: val, isSet: true}
}

func (v NullableQueueJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueueJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


