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

// checks if the RecipesQuotas type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecipesQuotas{}

// RecipesQuotas struct for RecipesQuotas
type RecipesQuotas struct {
	MaxCreatedRecipeDeployments int32 `json:"max_created_recipe_deployments"`
	RecipeInstanceQuota int32 `json:"recipe_instance_quota"`
}

// NewRecipesQuotas instantiates a new RecipesQuotas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipesQuotas(maxCreatedRecipeDeployments int32, recipeInstanceQuota int32) *RecipesQuotas {
	this := RecipesQuotas{}
	this.MaxCreatedRecipeDeployments = maxCreatedRecipeDeployments
	this.RecipeInstanceQuota = recipeInstanceQuota
	return &this
}

// NewRecipesQuotasWithDefaults instantiates a new RecipesQuotas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipesQuotasWithDefaults() *RecipesQuotas {
	this := RecipesQuotas{}
	return &this
}

// GetMaxCreatedRecipeDeployments returns the MaxCreatedRecipeDeployments field value
func (o *RecipesQuotas) GetMaxCreatedRecipeDeployments() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxCreatedRecipeDeployments
}

// GetMaxCreatedRecipeDeploymentsOk returns a tuple with the MaxCreatedRecipeDeployments field value
// and a boolean to check if the value has been set.
func (o *RecipesQuotas) GetMaxCreatedRecipeDeploymentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxCreatedRecipeDeployments, true
}

// SetMaxCreatedRecipeDeployments sets field value
func (o *RecipesQuotas) SetMaxCreatedRecipeDeployments(v int32) {
	o.MaxCreatedRecipeDeployments = v
}

// GetRecipeInstanceQuota returns the RecipeInstanceQuota field value
func (o *RecipesQuotas) GetRecipeInstanceQuota() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RecipeInstanceQuota
}

// GetRecipeInstanceQuotaOk returns a tuple with the RecipeInstanceQuota field value
// and a boolean to check if the value has been set.
func (o *RecipesQuotas) GetRecipeInstanceQuotaOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipeInstanceQuota, true
}

// SetRecipeInstanceQuota sets field value
func (o *RecipesQuotas) SetRecipeInstanceQuota(v int32) {
	o.RecipeInstanceQuota = v
}

func (o RecipesQuotas) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecipesQuotas) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["max_created_recipe_deployments"] = o.MaxCreatedRecipeDeployments
	toSerialize["recipe_instance_quota"] = o.RecipeInstanceQuota
	return toSerialize, nil
}

type NullableRecipesQuotas struct {
	value *RecipesQuotas
	isSet bool
}

func (v NullableRecipesQuotas) Get() *RecipesQuotas {
	return v.value
}

func (v *NullableRecipesQuotas) Set(val *RecipesQuotas) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipesQuotas) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipesQuotas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipesQuotas(val *RecipesQuotas) *NullableRecipesQuotas {
	return &NullableRecipesQuotas{value: val, isSet: true}
}

func (v NullableRecipesQuotas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipesQuotas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


