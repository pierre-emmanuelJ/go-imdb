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

// SimilarShort struct for SimilarShort
type SimilarShort struct {
	Id NullableString `json:"id,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Image NullableString `json:"image,omitempty"`
	ImDbRating NullableString `json:"imDbRating,omitempty"`
}

// NewSimilarShort instantiates a new SimilarShort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimilarShort() *SimilarShort {
	this := SimilarShort{}
	return &this
}

// NewSimilarShortWithDefaults instantiates a new SimilarShort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimilarShortWithDefaults() *SimilarShort {
	this := SimilarShort{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimilarShort) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimilarShort) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *SimilarShort) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *SimilarShort) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *SimilarShort) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *SimilarShort) UnsetId() {
	o.Id.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimilarShort) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimilarShort) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *SimilarShort) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *SimilarShort) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *SimilarShort) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *SimilarShort) UnsetTitle() {
	o.Title.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimilarShort) GetImage() string {
	if o == nil || o.Image.Get() == nil {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimilarShort) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *SimilarShort) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *SimilarShort) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *SimilarShort) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *SimilarShort) UnsetImage() {
	o.Image.Unset()
}

// GetImDbRating returns the ImDbRating field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimilarShort) GetImDbRating() string {
	if o == nil || o.ImDbRating.Get() == nil {
		var ret string
		return ret
	}
	return *o.ImDbRating.Get()
}

// GetImDbRatingOk returns a tuple with the ImDbRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimilarShort) GetImDbRatingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImDbRating.Get(), o.ImDbRating.IsSet()
}

// HasImDbRating returns a boolean if a field has been set.
func (o *SimilarShort) HasImDbRating() bool {
	if o != nil && o.ImDbRating.IsSet() {
		return true
	}

	return false
}

// SetImDbRating gets a reference to the given NullableString and assigns it to the ImDbRating field.
func (o *SimilarShort) SetImDbRating(v string) {
	o.ImDbRating.Set(&v)
}
// SetImDbRatingNil sets the value for ImDbRating to be an explicit nil
func (o *SimilarShort) SetImDbRatingNil() {
	o.ImDbRating.Set(nil)
}

// UnsetImDbRating ensures that no value is present for ImDbRating, not even an explicit nil
func (o *SimilarShort) UnsetImDbRating() {
	o.ImDbRating.Unset()
}

func (o SimilarShort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	if o.ImDbRating.IsSet() {
		toSerialize["imDbRating"] = o.ImDbRating.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSimilarShort struct {
	value *SimilarShort
	isSet bool
}

func (v NullableSimilarShort) Get() *SimilarShort {
	return v.value
}

func (v *NullableSimilarShort) Set(val *SimilarShort) {
	v.value = val
	v.isSet = true
}

func (v NullableSimilarShort) IsSet() bool {
	return v.isSet
}

func (v *NullableSimilarShort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimilarShort(val *SimilarShort) *NullableSimilarShort {
	return &NullableSimilarShort{value: val, isSet: true}
}

func (v NullableSimilarShort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimilarShort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


