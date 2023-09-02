/*
SaladCloud Public API

The SaladCloud Public API.

API version: 1.0.0-alpha.56
Contact: support@salad.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package salad-client

import (
	"encoding/json"
)

// checks if the RecipeResources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecipeResources{}

// RecipeResources Represents a recipe resources
type RecipeResources struct {
	Cpu int32 `json:"cpu"`
	Ram int32 `json:"ram"`
	GpuClass string `json:"gpu_class"`
}

// NewRecipeResources instantiates a new RecipeResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipeResources(cpu int32, ram int32, gpuClass string) *RecipeResources {
	this := RecipeResources{}
	this.Cpu = cpu
	this.Ram = ram
	this.GpuClass = gpuClass
	return &this
}

// NewRecipeResourcesWithDefaults instantiates a new RecipeResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipeResourcesWithDefaults() *RecipeResources {
	this := RecipeResources{}
	return &this
}

// GetCpu returns the Cpu field value
func (o *RecipeResources) GetCpu() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *RecipeResources) GetCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value
func (o *RecipeResources) SetCpu(v int32) {
	o.Cpu = v
}

// GetRam returns the Ram field value
func (o *RecipeResources) GetRam() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Ram
}

// GetRamOk returns a tuple with the Ram field value
// and a boolean to check if the value has been set.
func (o *RecipeResources) GetRamOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ram, true
}

// SetRam sets field value
func (o *RecipeResources) SetRam(v int32) {
	o.Ram = v
}

// GetGpuClass returns the GpuClass field value
func (o *RecipeResources) GetGpuClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GpuClass
}

// GetGpuClassOk returns a tuple with the GpuClass field value
// and a boolean to check if the value has been set.
func (o *RecipeResources) GetGpuClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GpuClass, true
}

// SetGpuClass sets field value
func (o *RecipeResources) SetGpuClass(v string) {
	o.GpuClass = v
}

func (o RecipeResources) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecipeResources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cpu"] = o.Cpu
	toSerialize["ram"] = o.Ram
	toSerialize["gpu_class"] = o.GpuClass
	return toSerialize, nil
}

type NullableRecipeResources struct {
	value *RecipeResources
	isSet bool
}

func (v NullableRecipeResources) Get() *RecipeResources {
	return v.value
}

func (v *NullableRecipeResources) Set(val *RecipeResources) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipeResources) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipeResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipeResources(val *RecipeResources) *NullableRecipeResources {
	return &NullableRecipeResources{value: val, isSet: true}
}

func (v NullableRecipeResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipeResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


