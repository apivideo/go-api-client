/*
 * api.video
 *
 * api.video is an API that encodes on the go to facilitate immediate playback, enhancing viewer streaming experiences across multiple devices and platforms. You can stream live or on-demand online videos within minutes.
 *
 * API version: 1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apivideosdk

import (
//"encoding/json"
)

// UnrecognizedRequestUrl struct for UnrecognizedRequestUrl
type UnrecognizedRequestUrl struct {
	// A link to the error documentation.
	Type *string `json:"type,omitempty"`
	// A description of the error that occurred.
	Title *string `json:"title,omitempty"`
	// The HTTP status code.
	Status *int32 `json:"status,omitempty"`
}

// NewUnrecognizedRequestUrl instantiates a new UnrecognizedRequestUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnrecognizedRequestUrl() *UnrecognizedRequestUrl {
	this := UnrecognizedRequestUrl{}
	return &this
}

// NewUnrecognizedRequestUrlWithDefaults instantiates a new UnrecognizedRequestUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnrecognizedRequestUrlWithDefaults() *UnrecognizedRequestUrl {
	this := UnrecognizedRequestUrl{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UnrecognizedRequestUrl) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnrecognizedRequestUrl) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UnrecognizedRequestUrl) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UnrecognizedRequestUrl) SetType(v string) {
	o.Type = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UnrecognizedRequestUrl) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnrecognizedRequestUrl) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UnrecognizedRequestUrl) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UnrecognizedRequestUrl) SetTitle(v string) {
	o.Title = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnrecognizedRequestUrl) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnrecognizedRequestUrl) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnrecognizedRequestUrl) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *UnrecognizedRequestUrl) SetStatus(v int32) {
	o.Status = &v
}

type NullableUnrecognizedRequestUrl struct {
	value *UnrecognizedRequestUrl
	isSet bool
}

func (v NullableUnrecognizedRequestUrl) Get() *UnrecognizedRequestUrl {
	return v.value
}

func (v *NullableUnrecognizedRequestUrl) Set(val *UnrecognizedRequestUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableUnrecognizedRequestUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableUnrecognizedRequestUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnrecognizedRequestUrl(val *UnrecognizedRequestUrl) *NullableUnrecognizedRequestUrl {
	return &NullableUnrecognizedRequestUrl{value: val, isSet: true}
}