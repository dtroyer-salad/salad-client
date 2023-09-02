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

// checks if the RecipeDeploymentInstances type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecipeDeploymentInstances{}

// RecipeDeploymentInstances Represents a list of recipe deployment instances
type RecipeDeploymentInstances struct {
	Instances []RecipeDeploymentInstancesInstancesInner `json:"instances"`
}

// NewRecipeDeploymentInstances instantiates a new RecipeDeploymentInstances object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipeDeploymentInstances(instances []RecipeDeploymentInstancesInstancesInner) *RecipeDeploymentInstances {
	this := RecipeDeploymentInstances{}
	this.Instances = instances
	return &this
}

// NewRecipeDeploymentInstancesWithDefaults instantiates a new RecipeDeploymentInstances object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipeDeploymentInstancesWithDefaults() *RecipeDeploymentInstances {
	this := RecipeDeploymentInstances{}
	return &this
}

// GetInstances returns the Instances field value
func (o *RecipeDeploymentInstances) GetInstances() []RecipeDeploymentInstancesInstancesInner {
	if o == nil {
		var ret []RecipeDeploymentInstancesInstancesInner
		return ret
	}

	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value
// and a boolean to check if the value has been set.
func (o *RecipeDeploymentInstances) GetInstancesOk() ([]RecipeDeploymentInstancesInstancesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instances, true
}

// SetInstances sets field value
func (o *RecipeDeploymentInstances) SetInstances(v []RecipeDeploymentInstancesInstancesInner) {
	o.Instances = v
}

func (o RecipeDeploymentInstances) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecipeDeploymentInstances) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instances"] = o.Instances
	return toSerialize, nil
}

type NullableRecipeDeploymentInstances struct {
	value *RecipeDeploymentInstances
	isSet bool
}

func (v NullableRecipeDeploymentInstances) Get() *RecipeDeploymentInstances {
	return v.value
}

func (v *NullableRecipeDeploymentInstances) Set(val *RecipeDeploymentInstances) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipeDeploymentInstances) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipeDeploymentInstances) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipeDeploymentInstances(val *RecipeDeploymentInstances) *NullableRecipeDeploymentInstances {
	return &NullableRecipeDeploymentInstances{value: val, isSet: true}
}

func (v NullableRecipeDeploymentInstances) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipeDeploymentInstances) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


