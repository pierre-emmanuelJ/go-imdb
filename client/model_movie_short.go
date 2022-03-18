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

// MovieShort struct for MovieShort
type MovieShort struct {
	Id NullableString `json:"id,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Year NullableString `json:"year,omitempty"`
	Image NullableString `json:"image,omitempty"`
	ImDbRating NullableString `json:"imDbRating,omitempty"`
}

// NewMovieShort instantiates a new MovieShort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMovieShort() *MovieShort {
	this := MovieShort{}
	return &this
}

// NewMovieShortWithDefaults instantiates a new MovieShort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMovieShortWithDefaults() *MovieShort {
	this := MovieShort{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieShort) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieShort) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MovieShort) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *MovieShort) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MovieShort) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MovieShort) UnsetId() {
	o.Id.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieShort) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieShort) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *MovieShort) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *MovieShort) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *MovieShort) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *MovieShort) UnsetTitle() {
	o.Title.Unset()
}

// GetYear returns the Year field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieShort) GetYear() string {
	if o == nil || o.Year.Get() == nil {
		var ret string
		return ret
	}
	return *o.Year.Get()
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieShort) GetYearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Year.Get(), o.Year.IsSet()
}

// HasYear returns a boolean if a field has been set.
func (o *MovieShort) HasYear() bool {
	if o != nil && o.Year.IsSet() {
		return true
	}

	return false
}

// SetYear gets a reference to the given NullableString and assigns it to the Year field.
func (o *MovieShort) SetYear(v string) {
	o.Year.Set(&v)
}
// SetYearNil sets the value for Year to be an explicit nil
func (o *MovieShort) SetYearNil() {
	o.Year.Set(nil)
}

// UnsetYear ensures that no value is present for Year, not even an explicit nil
func (o *MovieShort) UnsetYear() {
	o.Year.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieShort) GetImage() string {
	if o == nil || o.Image.Get() == nil {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieShort) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *MovieShort) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *MovieShort) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *MovieShort) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *MovieShort) UnsetImage() {
	o.Image.Unset()
}

// GetImDbRating returns the ImDbRating field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieShort) GetImDbRating() string {
	if o == nil || o.ImDbRating.Get() == nil {
		var ret string
		return ret
	}
	return *o.ImDbRating.Get()
}

// GetImDbRatingOk returns a tuple with the ImDbRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieShort) GetImDbRatingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImDbRating.Get(), o.ImDbRating.IsSet()
}

// HasImDbRating returns a boolean if a field has been set.
func (o *MovieShort) HasImDbRating() bool {
	if o != nil && o.ImDbRating.IsSet() {
		return true
	}

	return false
}

// SetImDbRating gets a reference to the given NullableString and assigns it to the ImDbRating field.
func (o *MovieShort) SetImDbRating(v string) {
	o.ImDbRating.Set(&v)
}
// SetImDbRatingNil sets the value for ImDbRating to be an explicit nil
func (o *MovieShort) SetImDbRatingNil() {
	o.ImDbRating.Set(nil)
}

// UnsetImDbRating ensures that no value is present for ImDbRating, not even an explicit nil
func (o *MovieShort) UnsetImDbRating() {
	o.ImDbRating.Unset()
}

func (o MovieShort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Year.IsSet() {
		toSerialize["year"] = o.Year.Get()
	}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	if o.ImDbRating.IsSet() {
		toSerialize["imDbRating"] = o.ImDbRating.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMovieShort struct {
	value *MovieShort
	isSet bool
}

func (v NullableMovieShort) Get() *MovieShort {
	return v.value
}

func (v *NullableMovieShort) Set(val *MovieShort) {
	v.value = val
	v.isSet = true
}

func (v NullableMovieShort) IsSet() bool {
	return v.isSet
}

func (v *NullableMovieShort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovieShort(val *MovieShort) *NullableMovieShort {
	return &NullableMovieShort{value: val, isSet: true}
}

func (v NullableMovieShort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovieShort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


