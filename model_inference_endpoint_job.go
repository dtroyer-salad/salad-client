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

// checks if the InferenceEndpointJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InferenceEndpointJob{}

// InferenceEndpointJob Represents a inference endpoint job
type InferenceEndpointJob struct {
	// The inference endpoint job identifier.
	Id string `json:"id"`
	// The inference endpoint name.
	InferenceEndpointName string `json:"inference_endpoint_name" validate:"regexp=^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	// The organization name.
	OrganizationName string `json:"organization_name" validate:"regexp=^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	Input interface{} `json:"input"`
	// The job metadata. May be any valid JSON.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The webhook URL called when the job completes.
	// Deprecated
	Webhook *string `json:"webhook,omitempty"`
	// The webhook URL called when the job completes.
	WebhookUrl *string `json:"webhook_url,omitempty"`
	Status InferenceEndpointJobStatus `json:"status"`
	// The list of events.
	Events []InferenceEndpointJobEvent `json:"events"`
	Output interface{} `json:"output,omitempty"`
	// The time the job was created.
	CreateTime time.Time `json:"create_time"`
	// The time the job was last updated.
	UpdateTime time.Time `json:"update_time"`
}

type _InferenceEndpointJob InferenceEndpointJob

// NewInferenceEndpointJob instantiates a new InferenceEndpointJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInferenceEndpointJob(id string, inferenceEndpointName string, organizationName string, input interface{}, status InferenceEndpointJobStatus, events []InferenceEndpointJobEvent, createTime time.Time, updateTime time.Time) *InferenceEndpointJob {
	this := InferenceEndpointJob{}
	this.Id = id
	this.InferenceEndpointName = inferenceEndpointName
	this.OrganizationName = organizationName
	this.Input = input
	this.Status = status
	this.Events = events
	this.CreateTime = createTime
	this.UpdateTime = updateTime
	return &this
}

// NewInferenceEndpointJobWithDefaults instantiates a new InferenceEndpointJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInferenceEndpointJobWithDefaults() *InferenceEndpointJob {
	this := InferenceEndpointJob{}
	return &this
}

// GetId returns the Id field value
func (o *InferenceEndpointJob) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InferenceEndpointJob) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InferenceEndpointJob) SetId(v string) {
	o.Id = v
}

// GetInferenceEndpointName returns the InferenceEndpointName field value
func (o *InferenceEndpointJob) GetInferenceEndpointName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InferenceEndpointName
}

// GetInferenceEndpointNameOk returns a tuple with the InferenceEndpointName field value
// and a boolean to check if the value has been set.
func (o *InferenceEndpointJob) GetInferenceEndpointNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InferenceEndpointName, true
}

// SetInferenceEndpointName sets field value
func (o *InferenceEndpointJob) SetInferenceEndpointName(v string) {
	o.InferenceEndpointName = v
}

// GetOrganizationName returns the OrganizationName field value
func (o *InferenceEndpointJob) GetOrganizationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value
// and a boolean to check if the value has been set.
func (o *InferenceEndpointJob) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationName, true
}

// SetOrganizationName sets field value
func (o *InferenceEndpointJob) SetOrganizationName(v string) {
	o.OrganizationName = v
}

// GetInput returns the Input field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *InferenceEndpointJob) GetInput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InferenceEndpointJob) GetInputOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Input) {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *InferenceEndpointJob) SetInput(v interface{}) {
	o.Input = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *InferenceEndpointJob) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InferenceEndpointJob) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *InferenceEndpointJob) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *InferenceEndpointJob) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
// Deprecated
func (o *InferenceEndpointJob) GetWebhook() string {
	if o == nil || IsNil(o.Webhook) {
		var ret string
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *InferenceEndpointJob) GetWebhookOk() (*string, bool) {
	if o == nil || IsNil(o.Webhook) {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *InferenceEndpointJob) HasWebhook() bool {
	if o != nil && !IsNil(o.Webhook) {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given string and assigns it to the Webhook field.
// Deprecated
func (o *InferenceEndpointJob) SetWebhook(v string) {
	o.Webhook = &v
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *InferenceEndpointJob) GetWebhookUrl() string {
	if o == nil || IsNil(o.WebhookUrl) {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InferenceEndpointJob) GetWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookUrl) {
		return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *InferenceEndpointJob) HasWebhookUrl() bool {
	if o != nil && !IsNil(o.WebhookUrl) {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *InferenceEndpointJob) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

// GetStatus returns the Status field value
func (o *InferenceEndpointJob) GetStatus() InferenceEndpointJobStatus {
	if o == nil {
		var ret InferenceEndpointJobStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InferenceEndpointJob) GetStatusOk() (*InferenceEndpointJobStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InferenceEndpointJob) SetStatus(v InferenceEndpointJobStatus) {
	o.Status = v
}

// GetEvents returns the Events field value
func (o *InferenceEndpointJob) GetEvents() []InferenceEndpointJobEvent {
	if o == nil {
		var ret []InferenceEndpointJobEvent
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *InferenceEndpointJob) GetEventsOk() ([]InferenceEndpointJobEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *InferenceEndpointJob) SetEvents(v []InferenceEndpointJobEvent) {
	o.Events = v
}

// GetOutput returns the Output field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InferenceEndpointJob) GetOutput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InferenceEndpointJob) GetOutputOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Output) {
		return nil, false
	}
	return &o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *InferenceEndpointJob) HasOutput() bool {
	if o != nil && !IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given interface{} and assigns it to the Output field.
func (o *InferenceEndpointJob) SetOutput(v interface{}) {
	o.Output = v
}

// GetCreateTime returns the CreateTime field value
func (o *InferenceEndpointJob) GetCreateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value
// and a boolean to check if the value has been set.
func (o *InferenceEndpointJob) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateTime, true
}

// SetCreateTime sets field value
func (o *InferenceEndpointJob) SetCreateTime(v time.Time) {
	o.CreateTime = v
}

// GetUpdateTime returns the UpdateTime field value
func (o *InferenceEndpointJob) GetUpdateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value
// and a boolean to check if the value has been set.
func (o *InferenceEndpointJob) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateTime, true
}

// SetUpdateTime sets field value
func (o *InferenceEndpointJob) SetUpdateTime(v time.Time) {
	o.UpdateTime = v
}

func (o InferenceEndpointJob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InferenceEndpointJob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["inference_endpoint_name"] = o.InferenceEndpointName
	toSerialize["organization_name"] = o.OrganizationName
	if o.Input != nil {
		toSerialize["input"] = o.Input
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Webhook) {
		toSerialize["webhook"] = o.Webhook
	}
	if !IsNil(o.WebhookUrl) {
		toSerialize["webhook_url"] = o.WebhookUrl
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

func (o *InferenceEndpointJob) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"inference_endpoint_name",
		"organization_name",
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

	varInferenceEndpointJob := _InferenceEndpointJob{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInferenceEndpointJob)

	if err != nil {
		return err
	}

	*o = InferenceEndpointJob(varInferenceEndpointJob)

	return err
}

type NullableInferenceEndpointJob struct {
	value *InferenceEndpointJob
	isSet bool
}

func (v NullableInferenceEndpointJob) Get() *InferenceEndpointJob {
	return v.value
}

func (v *NullableInferenceEndpointJob) Set(val *InferenceEndpointJob) {
	v.value = val
	v.isSet = true
}

func (v NullableInferenceEndpointJob) IsSet() bool {
	return v.isSet
}

func (v *NullableInferenceEndpointJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInferenceEndpointJob(val *InferenceEndpointJob) *NullableInferenceEndpointJob {
	return &NullableInferenceEndpointJob{value: val, isSet: true}
}

func (v NullableInferenceEndpointJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInferenceEndpointJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


