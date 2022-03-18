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

// KeywordData struct for KeywordData
type KeywordData struct {
	Keyword NullableString `json:"keyword,omitempty"`
	Items []MovieShort `json:"items,omitempty"`
	ErrorMessage NullableString `json:"errorMessage,omitempty"`
}

// NewKeywordData instantiates a new KeywordData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeywordData() *KeywordData {
	this := KeywordData{}
	return &this
}

// NewKeywordDataWithDefaults instantiates a new KeywordData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeywordDataWithDefaults() *KeywordData {
	this := KeywordData{}
	return &this
}

// GetKeyword returns the Keyword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeywordData) GetKeyword() string {
	if o == nil || o.Keyword.Get() == nil {
		var ret string
		return ret
	}
	return *o.Keyword.Get()
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeywordData) GetKeywordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Keyword.Get(), o.Keyword.IsSet()
}

// HasKeyword returns a boolean if a field has been set.
func (o *KeywordData) HasKeyword() bool {
	if o != nil && o.Keyword.IsSet() {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given NullableString and assigns it to the Keyword field.
func (o *KeywordData) SetKeyword(v string) {
	o.Keyword.Set(&v)
}
// SetKeywordNil sets the value for Keyword to be an explicit nil
func (o *KeywordData) SetKeywordNil() {
	o.Keyword.Set(nil)
}

// UnsetKeyword ensures that no value is present for Keyword, not even an explicit nil
func (o *KeywordData) UnsetKeyword() {
	o.Keyword.Unset()
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeywordData) GetItems() []MovieShort {
	if o == nil {
		var ret []MovieShort
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeywordData) GetItemsOk() ([]MovieShort, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *KeywordData) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []MovieShort and assigns it to the Items field.
func (o *KeywordData) SetItems(v []MovieShort) {
	o.Items = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeywordData) GetErrorMessage() string {
	if o == nil || o.ErrorMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeywordData) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *KeywordData) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *KeywordData) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *KeywordData) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *KeywordData) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

func (o KeywordData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Keyword.IsSet() {
		toSerialize["keyword"] = o.Keyword.Get()
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableKeywordData struct {
	value *KeywordData
	isSet bool
}

func (v NullableKeywordData) Get() *KeywordData {
	return v.value
}

func (v *NullableKeywordData) Set(val *KeywordData) {
	v.value = val
	v.isSet = true
}

func (v NullableKeywordData) IsSet() bool {
	return v.isSet
}

func (v *NullableKeywordData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeywordData(val *KeywordData) *NullableKeywordData {
	return &NullableKeywordData{value: val, isSet: true}
}

func (v NullableKeywordData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeywordData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

