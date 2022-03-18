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

// UserRatingDataDemographicDetail struct for UserRatingDataDemographicDetail
type UserRatingDataDemographicDetail struct {
	Rating NullableString `json:"rating,omitempty"`
	Votes NullableString `json:"votes,omitempty"`
}

// NewUserRatingDataDemographicDetail instantiates a new UserRatingDataDemographicDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRatingDataDemographicDetail() *UserRatingDataDemographicDetail {
	this := UserRatingDataDemographicDetail{}
	return &this
}

// NewUserRatingDataDemographicDetailWithDefaults instantiates a new UserRatingDataDemographicDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRatingDataDemographicDetailWithDefaults() *UserRatingDataDemographicDetail {
	this := UserRatingDataDemographicDetail{}
	return &this
}

// GetRating returns the Rating field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserRatingDataDemographicDetail) GetRating() string {
	if o == nil || o.Rating.Get() == nil {
		var ret string
		return ret
	}
	return *o.Rating.Get()
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserRatingDataDemographicDetail) GetRatingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rating.Get(), o.Rating.IsSet()
}

// HasRating returns a boolean if a field has been set.
func (o *UserRatingDataDemographicDetail) HasRating() bool {
	if o != nil && o.Rating.IsSet() {
		return true
	}

	return false
}

// SetRating gets a reference to the given NullableString and assigns it to the Rating field.
func (o *UserRatingDataDemographicDetail) SetRating(v string) {
	o.Rating.Set(&v)
}
// SetRatingNil sets the value for Rating to be an explicit nil
func (o *UserRatingDataDemographicDetail) SetRatingNil() {
	o.Rating.Set(nil)
}

// UnsetRating ensures that no value is present for Rating, not even an explicit nil
func (o *UserRatingDataDemographicDetail) UnsetRating() {
	o.Rating.Unset()
}

// GetVotes returns the Votes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserRatingDataDemographicDetail) GetVotes() string {
	if o == nil || o.Votes.Get() == nil {
		var ret string
		return ret
	}
	return *o.Votes.Get()
}

// GetVotesOk returns a tuple with the Votes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserRatingDataDemographicDetail) GetVotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Votes.Get(), o.Votes.IsSet()
}

// HasVotes returns a boolean if a field has been set.
func (o *UserRatingDataDemographicDetail) HasVotes() bool {
	if o != nil && o.Votes.IsSet() {
		return true
	}

	return false
}

// SetVotes gets a reference to the given NullableString and assigns it to the Votes field.
func (o *UserRatingDataDemographicDetail) SetVotes(v string) {
	o.Votes.Set(&v)
}
// SetVotesNil sets the value for Votes to be an explicit nil
func (o *UserRatingDataDemographicDetail) SetVotesNil() {
	o.Votes.Set(nil)
}

// UnsetVotes ensures that no value is present for Votes, not even an explicit nil
func (o *UserRatingDataDemographicDetail) UnsetVotes() {
	o.Votes.Unset()
}

func (o UserRatingDataDemographicDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rating.IsSet() {
		toSerialize["rating"] = o.Rating.Get()
	}
	if o.Votes.IsSet() {
		toSerialize["votes"] = o.Votes.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUserRatingDataDemographicDetail struct {
	value *UserRatingDataDemographicDetail
	isSet bool
}

func (v NullableUserRatingDataDemographicDetail) Get() *UserRatingDataDemographicDetail {
	return v.value
}

func (v *NullableUserRatingDataDemographicDetail) Set(val *UserRatingDataDemographicDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRatingDataDemographicDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRatingDataDemographicDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRatingDataDemographicDetail(val *UserRatingDataDemographicDetail) *NullableUserRatingDataDemographicDetail {
	return &NullableUserRatingDataDemographicDetail{value: val, isSet: true}
}

func (v NullableUserRatingDataDemographicDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRatingDataDemographicDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


