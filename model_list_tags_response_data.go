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

// ListTagsResponseData struct for ListTagsResponseData
type ListTagsResponseData struct {
	// Returns the value of a video tag used in your project.
	Value *string `json:"value,omitempty"`
	// Returns the number of times a video tag is used.
	VideoCount *int32 `json:"videoCount,omitempty"`
}

// NewListTagsResponseData instantiates a new ListTagsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTagsResponseData() *ListTagsResponseData {
	this := ListTagsResponseData{}
	return &this
}

// NewListTagsResponseDataWithDefaults instantiates a new ListTagsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTagsResponseDataWithDefaults() *ListTagsResponseData {
	this := ListTagsResponseData{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ListTagsResponseData) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTagsResponseData) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ListTagsResponseData) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ListTagsResponseData) SetValue(v string) {
	o.Value = &v
}

// GetVideoCount returns the VideoCount field value if set, zero value otherwise.
func (o *ListTagsResponseData) GetVideoCount() int32 {
	if o == nil || o.VideoCount == nil {
		var ret int32
		return ret
	}
	return *o.VideoCount
}

// GetVideoCountOk returns a tuple with the VideoCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTagsResponseData) GetVideoCountOk() (*int32, bool) {
	if o == nil || o.VideoCount == nil {
		return nil, false
	}
	return o.VideoCount, true
}

// HasVideoCount returns a boolean if a field has been set.
func (o *ListTagsResponseData) HasVideoCount() bool {
	if o != nil && o.VideoCount != nil {
		return true
	}

	return false
}

// SetVideoCount gets a reference to the given int32 and assigns it to the VideoCount field.
func (o *ListTagsResponseData) SetVideoCount(v int32) {
	o.VideoCount = &v
}

type NullableListTagsResponseData struct {
	value *ListTagsResponseData
	isSet bool
}

func (v NullableListTagsResponseData) Get() *ListTagsResponseData {
	return v.value
}

func (v *NullableListTagsResponseData) Set(val *ListTagsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableListTagsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableListTagsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTagsResponseData(val *ListTagsResponseData) *NullableListTagsResponseData {
	return &NullableListTagsResponseData{value: val, isSet: true}
}