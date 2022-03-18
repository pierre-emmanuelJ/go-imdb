/*
IMDb-API

The IMDb-API Documentation. You need a <a href='/Identity/Account/Manage' target='_blank'><code>API Key</code></a> for testing APIs.<br/><a class='link' href='/API'>Back to API Tester</a>

API version: 1.8.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ImageDataDetail struct for ImageDataDetail
type ImageDataDetail struct {
	Title NullableString `json:"title,omitempty"`
	Image NullableString `json:"image,omitempty"`
}

// NewImageDataDetail instantiates a new ImageDataDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageDataDetail() *ImageDataDetail {
	this := ImageDataDetail{}
	return &this
}

// NewImageDataDetailWithDefaults instantiates a new ImageDataDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageDataDetailWithDefaults() *ImageDataDetail {
	this := ImageDataDetail{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageDataDetail) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageDataDetail) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ImageDataDetail) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ImageDataDetail) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ImageDataDetail) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ImageDataDetail) UnsetTitle() {
	o.Title.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageDataDetail) GetImage() string {
	if o == nil || o.Image.Get() == nil {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageDataDetail) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *ImageDataDetail) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *ImageDataDetail) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *ImageDataDetail) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *ImageDataDetail) UnsetImage() {
	o.Image.Unset()
}

func (o ImageDataDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableImageDataDetail struct {
	value *ImageDataDetail
	isSet bool
}

func (v NullableImageDataDetail) Get() *ImageDataDetail {
	return v.value
}

func (v *NullableImageDataDetail) Set(val *ImageDataDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableImageDataDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableImageDataDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageDataDetail(val *ImageDataDetail) *NullableImageDataDetail {
	return &NullableImageDataDetail{value: val, isSet: true}
}

func (v NullableImageDataDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageDataDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

