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

// MostPopularData struct for MostPopularData
type MostPopularData struct {
	Items []MostPopularDataDetail `json:"items,omitempty"`
	ErrorMessage NullableString `json:"errorMessage,omitempty"`
}

// NewMostPopularData instantiates a new MostPopularData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMostPopularData() *MostPopularData {
	this := MostPopularData{}
	return &this
}

// NewMostPopularDataWithDefaults instantiates a new MostPopularData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMostPopularDataWithDefaults() *MostPopularData {
	this := MostPopularData{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MostPopularData) GetItems() []MostPopularDataDetail {
	if o == nil {
		var ret []MostPopularDataDetail
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MostPopularData) GetItemsOk() ([]MostPopularDataDetail, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *MostPopularData) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []MostPopularDataDetail and assigns it to the Items field.
func (o *MostPopularData) SetItems(v []MostPopularDataDetail) {
	o.Items = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MostPopularData) GetErrorMessage() string {
	if o == nil || o.ErrorMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MostPopularData) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *MostPopularData) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *MostPopularData) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *MostPopularData) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *MostPopularData) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

func (o MostPopularData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMostPopularData struct {
	value *MostPopularData
	isSet bool
}

func (v NullableMostPopularData) Get() *MostPopularData {
	return v.value
}

func (v *NullableMostPopularData) Set(val *MostPopularData) {
	v.value = val
	v.isSet = true
}

func (v NullableMostPopularData) IsSet() bool {
	return v.isSet
}

func (v *NullableMostPopularData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMostPopularData(val *MostPopularData) *NullableMostPopularData {
	return &NullableMostPopularData{value: val, isSet: true}
}

func (v NullableMostPopularData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMostPopularData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

