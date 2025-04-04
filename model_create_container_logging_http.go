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

// checks if the CreateContainerLoggingHttp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateContainerLoggingHttp{}

// CreateContainerLoggingHttp Configuration for sending container logs to an HTTP endpoint. Defines how logs are formatted, compressed, and transmitted.
type CreateContainerLoggingHttp struct {
	// The hostname or IP address of the HTTP logging endpoint
	Host string `json:"host" validate:"regexp=^.*$"`
	// The port number of the HTTP logging endpoint (1-65535)
	Port int32 `json:"port"`
	// Optional username for HTTP authentication
	User NullableString `json:"user,omitempty" validate:"regexp=^.*$"`
	// Optional password for HTTP authentication
	Password NullableString `json:"password,omitempty" validate:"regexp=^.*$"`
	// Optional URL path for the HTTP endpoint
	Path NullableString `json:"path,omitempty" validate:"regexp=^.*$"`
	Format ContainerLoggingHttpFormat `json:"format"`
	// Optional HTTP headers to include in log transmission requests
	Headers []ContainerLoggingHttpHeader `json:"headers,omitempty"`
	Compression ContainerLoggingHttpCompression `json:"compression"`
}

type _CreateContainerLoggingHttp CreateContainerLoggingHttp

// NewCreateContainerLoggingHttp instantiates a new CreateContainerLoggingHttp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateContainerLoggingHttp(host string, port int32, format ContainerLoggingHttpFormat, compression ContainerLoggingHttpCompression) *CreateContainerLoggingHttp {
	this := CreateContainerLoggingHttp{}
	this.Host = host
	this.Port = port
	this.Format = format
	this.Compression = compression
	return &this
}

// NewCreateContainerLoggingHttpWithDefaults instantiates a new CreateContainerLoggingHttp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateContainerLoggingHttpWithDefaults() *CreateContainerLoggingHttp {
	this := CreateContainerLoggingHttp{}
	return &this
}

// GetHost returns the Host field value
func (o *CreateContainerLoggingHttp) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *CreateContainerLoggingHttp) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *CreateContainerLoggingHttp) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *CreateContainerLoggingHttp) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *CreateContainerLoggingHttp) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *CreateContainerLoggingHttp) SetPort(v int32) {
	o.Port = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateContainerLoggingHttp) GetUser() string {
	if o == nil || IsNil(o.User.Get()) {
		var ret string
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateContainerLoggingHttp) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *CreateContainerLoggingHttp) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableString and assigns it to the User field.
func (o *CreateContainerLoggingHttp) SetUser(v string) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *CreateContainerLoggingHttp) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *CreateContainerLoggingHttp) UnsetUser() {
	o.User.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateContainerLoggingHttp) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateContainerLoggingHttp) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *CreateContainerLoggingHttp) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *CreateContainerLoggingHttp) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *CreateContainerLoggingHttp) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *CreateContainerLoggingHttp) UnsetPassword() {
	o.Password.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateContainerLoggingHttp) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateContainerLoggingHttp) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *CreateContainerLoggingHttp) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *CreateContainerLoggingHttp) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *CreateContainerLoggingHttp) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *CreateContainerLoggingHttp) UnsetPath() {
	o.Path.Unset()
}

// GetFormat returns the Format field value
func (o *CreateContainerLoggingHttp) GetFormat() ContainerLoggingHttpFormat {
	if o == nil {
		var ret ContainerLoggingHttpFormat
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *CreateContainerLoggingHttp) GetFormatOk() (*ContainerLoggingHttpFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *CreateContainerLoggingHttp) SetFormat(v ContainerLoggingHttpFormat) {
	o.Format = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *CreateContainerLoggingHttp) GetHeaders() []ContainerLoggingHttpHeader {
	if o == nil || IsNil(o.Headers) {
		var ret []ContainerLoggingHttpHeader
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainerLoggingHttp) GetHeadersOk() ([]ContainerLoggingHttpHeader, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *CreateContainerLoggingHttp) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []ContainerLoggingHttpHeader and assigns it to the Headers field.
func (o *CreateContainerLoggingHttp) SetHeaders(v []ContainerLoggingHttpHeader) {
	o.Headers = v
}

// GetCompression returns the Compression field value
func (o *CreateContainerLoggingHttp) GetCompression() ContainerLoggingHttpCompression {
	if o == nil {
		var ret ContainerLoggingHttpCompression
		return ret
	}

	return o.Compression
}

// GetCompressionOk returns a tuple with the Compression field value
// and a boolean to check if the value has been set.
func (o *CreateContainerLoggingHttp) GetCompressionOk() (*ContainerLoggingHttpCompression, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compression, true
}

// SetCompression sets field value
func (o *CreateContainerLoggingHttp) SetCompression(v ContainerLoggingHttpCompression) {
	o.Compression = v
}

func (o CreateContainerLoggingHttp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateContainerLoggingHttp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	toSerialize["port"] = o.Port
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	toSerialize["format"] = o.Format
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	toSerialize["compression"] = o.Compression
	return toSerialize, nil
}

func (o *CreateContainerLoggingHttp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"host",
		"port",
		"format",
		"compression",
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

	varCreateContainerLoggingHttp := _CreateContainerLoggingHttp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateContainerLoggingHttp)

	if err != nil {
		return err
	}

	*o = CreateContainerLoggingHttp(varCreateContainerLoggingHttp)

	return err
}

type NullableCreateContainerLoggingHttp struct {
	value *CreateContainerLoggingHttp
	isSet bool
}

func (v NullableCreateContainerLoggingHttp) Get() *CreateContainerLoggingHttp {
	return v.value
}

func (v *NullableCreateContainerLoggingHttp) Set(val *CreateContainerLoggingHttp) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateContainerLoggingHttp) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateContainerLoggingHttp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateContainerLoggingHttp(val *CreateContainerLoggingHttp) *NullableCreateContainerLoggingHttp {
	return &NullableCreateContainerLoggingHttp{value: val, isSet: true}
}

func (v NullableCreateContainerLoggingHttp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateContainerLoggingHttp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


