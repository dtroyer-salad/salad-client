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

// checks if the ContainerGroupProbeHttp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerGroupProbeHttp{}

// ContainerGroupProbeHttp struct for ContainerGroupProbeHttp
type ContainerGroupProbeHttp struct {
	Path string `json:"path"`
	Port int32 `json:"port"`
	Scheme *ContainerProbeHttpScheme `json:"scheme,omitempty"`
	Headers []HttpHeadersInner `json:"headers,omitempty"`
}

// NewContainerGroupProbeHttp instantiates a new ContainerGroupProbeHttp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerGroupProbeHttp(path string, port int32) *ContainerGroupProbeHttp {
	this := ContainerGroupProbeHttp{}
	this.Path = path
	this.Port = port
	return &this
}

// NewContainerGroupProbeHttpWithDefaults instantiates a new ContainerGroupProbeHttp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerGroupProbeHttpWithDefaults() *ContainerGroupProbeHttp {
	this := ContainerGroupProbeHttp{}
	return &this
}

// GetPath returns the Path field value
func (o *ContainerGroupProbeHttp) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbeHttp) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *ContainerGroupProbeHttp) SetPath(v string) {
	o.Path = v
}

// GetPort returns the Port field value
func (o *ContainerGroupProbeHttp) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbeHttp) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *ContainerGroupProbeHttp) SetPort(v int32) {
	o.Port = v
}

// GetScheme returns the Scheme field value if set, zero value otherwise.
func (o *ContainerGroupProbeHttp) GetScheme() ContainerProbeHttpScheme {
	if o == nil || IsNil(o.Scheme) {
		var ret ContainerProbeHttpScheme
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbeHttp) GetSchemeOk() (*ContainerProbeHttpScheme, bool) {
	if o == nil || IsNil(o.Scheme) {
		return nil, false
	}
	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *ContainerGroupProbeHttp) HasScheme() bool {
	if o != nil && !IsNil(o.Scheme) {
		return true
	}

	return false
}

// SetScheme gets a reference to the given ContainerProbeHttpScheme and assigns it to the Scheme field.
func (o *ContainerGroupProbeHttp) SetScheme(v ContainerProbeHttpScheme) {
	o.Scheme = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *ContainerGroupProbeHttp) GetHeaders() []HttpHeadersInner {
	if o == nil || IsNil(o.Headers) {
		var ret []HttpHeadersInner
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbeHttp) GetHeadersOk() ([]HttpHeadersInner, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *ContainerGroupProbeHttp) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []HttpHeadersInner and assigns it to the Headers field.
func (o *ContainerGroupProbeHttp) SetHeaders(v []HttpHeadersInner) {
	o.Headers = v
}

func (o ContainerGroupProbeHttp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerGroupProbeHttp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path
	toSerialize["port"] = o.Port
	if !IsNil(o.Scheme) {
		toSerialize["scheme"] = o.Scheme
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	return toSerialize, nil
}

type NullableContainerGroupProbeHttp struct {
	value *ContainerGroupProbeHttp
	isSet bool
}

func (v NullableContainerGroupProbeHttp) Get() *ContainerGroupProbeHttp {
	return v.value
}

func (v *NullableContainerGroupProbeHttp) Set(val *ContainerGroupProbeHttp) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerGroupProbeHttp) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerGroupProbeHttp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerGroupProbeHttp(val *ContainerGroupProbeHttp) *NullableContainerGroupProbeHttp {
	return &NullableContainerGroupProbeHttp{value: val, isSet: true}
}

func (v NullableContainerGroupProbeHttp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerGroupProbeHttp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


