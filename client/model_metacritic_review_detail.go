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

// MetacriticReviewDetail struct for MetacriticReviewDetail
type MetacriticReviewDetail struct {
	Publisher NullableString `json:"publisher,omitempty"`
	Author NullableString `json:"author,omitempty"`
	Link NullableString `json:"link,omitempty"`
	Rate NullableString `json:"rate,omitempty"`
	Content NullableString `json:"content,omitempty"`
}

// NewMetacriticReviewDetail instantiates a new MetacriticReviewDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetacriticReviewDetail() *MetacriticReviewDetail {
	this := MetacriticReviewDetail{}
	return &this
}

// NewMetacriticReviewDetailWithDefaults instantiates a new MetacriticReviewDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetacriticReviewDetailWithDefaults() *MetacriticReviewDetail {
	this := MetacriticReviewDetail{}
	return &this
}

// GetPublisher returns the Publisher field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetacriticReviewDetail) GetPublisher() string {
	if o == nil || o.Publisher.Get() == nil {
		var ret string
		return ret
	}
	return *o.Publisher.Get()
}

// GetPublisherOk returns a tuple with the Publisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetacriticReviewDetail) GetPublisherOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Publisher.Get(), o.Publisher.IsSet()
}

// HasPublisher returns a boolean if a field has been set.
func (o *MetacriticReviewDetail) HasPublisher() bool {
	if o != nil && o.Publisher.IsSet() {
		return true
	}

	return false
}

// SetPublisher gets a reference to the given NullableString and assigns it to the Publisher field.
func (o *MetacriticReviewDetail) SetPublisher(v string) {
	o.Publisher.Set(&v)
}
// SetPublisherNil sets the value for Publisher to be an explicit nil
func (o *MetacriticReviewDetail) SetPublisherNil() {
	o.Publisher.Set(nil)
}

// UnsetPublisher ensures that no value is present for Publisher, not even an explicit nil
func (o *MetacriticReviewDetail) UnsetPublisher() {
	o.Publisher.Unset()
}

// GetAuthor returns the Author field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetacriticReviewDetail) GetAuthor() string {
	if o == nil || o.Author.Get() == nil {
		var ret string
		return ret
	}
	return *o.Author.Get()
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetacriticReviewDetail) GetAuthorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Author.Get(), o.Author.IsSet()
}

// HasAuthor returns a boolean if a field has been set.
func (o *MetacriticReviewDetail) HasAuthor() bool {
	if o != nil && o.Author.IsSet() {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given NullableString and assigns it to the Author field.
func (o *MetacriticReviewDetail) SetAuthor(v string) {
	o.Author.Set(&v)
}
// SetAuthorNil sets the value for Author to be an explicit nil
func (o *MetacriticReviewDetail) SetAuthorNil() {
	o.Author.Set(nil)
}

// UnsetAuthor ensures that no value is present for Author, not even an explicit nil
func (o *MetacriticReviewDetail) UnsetAuthor() {
	o.Author.Unset()
}

// GetLink returns the Link field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetacriticReviewDetail) GetLink() string {
	if o == nil || o.Link.Get() == nil {
		var ret string
		return ret
	}
	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetacriticReviewDetail) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// HasLink returns a boolean if a field has been set.
func (o *MetacriticReviewDetail) HasLink() bool {
	if o != nil && o.Link.IsSet() {
		return true
	}

	return false
}

// SetLink gets a reference to the given NullableString and assigns it to the Link field.
func (o *MetacriticReviewDetail) SetLink(v string) {
	o.Link.Set(&v)
}
// SetLinkNil sets the value for Link to be an explicit nil
func (o *MetacriticReviewDetail) SetLinkNil() {
	o.Link.Set(nil)
}

// UnsetLink ensures that no value is present for Link, not even an explicit nil
func (o *MetacriticReviewDetail) UnsetLink() {
	o.Link.Unset()
}

// GetRate returns the Rate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetacriticReviewDetail) GetRate() string {
	if o == nil || o.Rate.Get() == nil {
		var ret string
		return ret
	}
	return *o.Rate.Get()
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetacriticReviewDetail) GetRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rate.Get(), o.Rate.IsSet()
}

// HasRate returns a boolean if a field has been set.
func (o *MetacriticReviewDetail) HasRate() bool {
	if o != nil && o.Rate.IsSet() {
		return true
	}

	return false
}

// SetRate gets a reference to the given NullableString and assigns it to the Rate field.
func (o *MetacriticReviewDetail) SetRate(v string) {
	o.Rate.Set(&v)
}
// SetRateNil sets the value for Rate to be an explicit nil
func (o *MetacriticReviewDetail) SetRateNil() {
	o.Rate.Set(nil)
}

// UnsetRate ensures that no value is present for Rate, not even an explicit nil
func (o *MetacriticReviewDetail) UnsetRate() {
	o.Rate.Unset()
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetacriticReviewDetail) GetContent() string {
	if o == nil || o.Content.Get() == nil {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetacriticReviewDetail) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *MetacriticReviewDetail) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *MetacriticReviewDetail) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *MetacriticReviewDetail) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *MetacriticReviewDetail) UnsetContent() {
	o.Content.Unset()
}

func (o MetacriticReviewDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Publisher.IsSet() {
		toSerialize["publisher"] = o.Publisher.Get()
	}
	if o.Author.IsSet() {
		toSerialize["author"] = o.Author.Get()
	}
	if o.Link.IsSet() {
		toSerialize["link"] = o.Link.Get()
	}
	if o.Rate.IsSet() {
		toSerialize["rate"] = o.Rate.Get()
	}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMetacriticReviewDetail struct {
	value *MetacriticReviewDetail
	isSet bool
}

func (v NullableMetacriticReviewDetail) Get() *MetacriticReviewDetail {
	return v.value
}

func (v *NullableMetacriticReviewDetail) Set(val *MetacriticReviewDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableMetacriticReviewDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableMetacriticReviewDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetacriticReviewDetail(val *MetacriticReviewDetail) *NullableMetacriticReviewDetail {
	return &NullableMetacriticReviewDetail{value: val, isSet: true}
}

func (v NullableMetacriticReviewDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetacriticReviewDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

