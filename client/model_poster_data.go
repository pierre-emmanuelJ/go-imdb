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

// PosterData struct for PosterData
type PosterData struct {
	ImDbId NullableString `json:"imDbId,omitempty"`
	Title NullableString `json:"title,omitempty"`
	FullTitle NullableString `json:"fullTitle,omitempty"`
	Type NullableString `json:"type,omitempty"`
	Year NullableString `json:"year,omitempty"`
	Posters []PosterDataItem `json:"posters,omitempty"`
	Backdrops []PosterDataItem `json:"backdrops,omitempty"`
	ErrorMessage NullableString `json:"errorMessage,omitempty"`
}

// NewPosterData instantiates a new PosterData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPosterData() *PosterData {
	this := PosterData{}
	return &this
}

// NewPosterDataWithDefaults instantiates a new PosterData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPosterDataWithDefaults() *PosterData {
	this := PosterData{}
	return &this
}

// GetImDbId returns the ImDbId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PosterData) GetImDbId() string {
	if o == nil || o.ImDbId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ImDbId.Get()
}

// GetImDbIdOk returns a tuple with the ImDbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PosterData) GetImDbIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImDbId.Get(), o.ImDbId.IsSet()
}

// HasImDbId returns a boolean if a field has been set.
func (o *PosterData) HasImDbId() bool {
	if o != nil && o.ImDbId.IsSet() {
		return true
	}

	return false
}

// SetImDbId gets a reference to the given NullableString and assigns it to the ImDbId field.
func (o *PosterData) SetImDbId(v string) {
	o.ImDbId.Set(&v)
}
// SetImDbIdNil sets the value for ImDbId to be an explicit nil
func (o *PosterData) SetImDbIdNil() {
	o.ImDbId.Set(nil)
}

// UnsetImDbId ensures that no value is present for ImDbId, not even an explicit nil
func (o *PosterData) UnsetImDbId() {
	o.ImDbId.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PosterData) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PosterData) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *PosterData) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *PosterData) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *PosterData) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *PosterData) UnsetTitle() {
	o.Title.Unset()
}

// GetFullTitle returns the FullTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PosterData) GetFullTitle() string {
	if o == nil || o.FullTitle.Get() == nil {
		var ret string
		return ret
	}
	return *o.FullTitle.Get()
}

// GetFullTitleOk returns a tuple with the FullTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PosterData) GetFullTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FullTitle.Get(), o.FullTitle.IsSet()
}

// HasFullTitle returns a boolean if a field has been set.
func (o *PosterData) HasFullTitle() bool {
	if o != nil && o.FullTitle.IsSet() {
		return true
	}

	return false
}

// SetFullTitle gets a reference to the given NullableString and assigns it to the FullTitle field.
func (o *PosterData) SetFullTitle(v string) {
	o.FullTitle.Set(&v)
}
// SetFullTitleNil sets the value for FullTitle to be an explicit nil
func (o *PosterData) SetFullTitleNil() {
	o.FullTitle.Set(nil)
}

// UnsetFullTitle ensures that no value is present for FullTitle, not even an explicit nil
func (o *PosterData) UnsetFullTitle() {
	o.FullTitle.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PosterData) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PosterData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *PosterData) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *PosterData) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *PosterData) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *PosterData) UnsetType() {
	o.Type.Unset()
}

// GetYear returns the Year field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PosterData) GetYear() string {
	if o == nil || o.Year.Get() == nil {
		var ret string
		return ret
	}
	return *o.Year.Get()
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PosterData) GetYearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Year.Get(), o.Year.IsSet()
}

// HasYear returns a boolean if a field has been set.
func (o *PosterData) HasYear() bool {
	if o != nil && o.Year.IsSet() {
		return true
	}

	return false
}

// SetYear gets a reference to the given NullableString and assigns it to the Year field.
func (o *PosterData) SetYear(v string) {
	o.Year.Set(&v)
}
// SetYearNil sets the value for Year to be an explicit nil
func (o *PosterData) SetYearNil() {
	o.Year.Set(nil)
}

// UnsetYear ensures that no value is present for Year, not even an explicit nil
func (o *PosterData) UnsetYear() {
	o.Year.Unset()
}

// GetPosters returns the Posters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PosterData) GetPosters() []PosterDataItem {
	if o == nil {
		var ret []PosterDataItem
		return ret
	}
	return o.Posters
}

// GetPostersOk returns a tuple with the Posters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PosterData) GetPostersOk() ([]PosterDataItem, bool) {
	if o == nil || o.Posters == nil {
		return nil, false
	}
	return o.Posters, true
}

// HasPosters returns a boolean if a field has been set.
func (o *PosterData) HasPosters() bool {
	if o != nil && o.Posters != nil {
		return true
	}

	return false
}

// SetPosters gets a reference to the given []PosterDataItem and assigns it to the Posters field.
func (o *PosterData) SetPosters(v []PosterDataItem) {
	o.Posters = v
}

// GetBackdrops returns the Backdrops field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PosterData) GetBackdrops() []PosterDataItem {
	if o == nil {
		var ret []PosterDataItem
		return ret
	}
	return o.Backdrops
}

// GetBackdropsOk returns a tuple with the Backdrops field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PosterData) GetBackdropsOk() ([]PosterDataItem, bool) {
	if o == nil || o.Backdrops == nil {
		return nil, false
	}
	return o.Backdrops, true
}

// HasBackdrops returns a boolean if a field has been set.
func (o *PosterData) HasBackdrops() bool {
	if o != nil && o.Backdrops != nil {
		return true
	}

	return false
}

// SetBackdrops gets a reference to the given []PosterDataItem and assigns it to the Backdrops field.
func (o *PosterData) SetBackdrops(v []PosterDataItem) {
	o.Backdrops = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PosterData) GetErrorMessage() string {
	if o == nil || o.ErrorMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PosterData) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *PosterData) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *PosterData) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *PosterData) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *PosterData) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

func (o PosterData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImDbId.IsSet() {
		toSerialize["imDbId"] = o.ImDbId.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.FullTitle.IsSet() {
		toSerialize["fullTitle"] = o.FullTitle.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Year.IsSet() {
		toSerialize["year"] = o.Year.Get()
	}
	if o.Posters != nil {
		toSerialize["posters"] = o.Posters
	}
	if o.Backdrops != nil {
		toSerialize["backdrops"] = o.Backdrops
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePosterData struct {
	value *PosterData
	isSet bool
}

func (v NullablePosterData) Get() *PosterData {
	return v.value
}

func (v *NullablePosterData) Set(val *PosterData) {
	v.value = val
	v.isSet = true
}

func (v NullablePosterData) IsSet() bool {
	return v.isSet
}

func (v *NullablePosterData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePosterData(val *PosterData) *NullablePosterData {
	return &NullablePosterData{value: val, isSet: true}
}

func (v NullablePosterData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePosterData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


