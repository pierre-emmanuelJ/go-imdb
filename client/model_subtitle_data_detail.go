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

// SubtitleDataDetail struct for SubtitleDataDetail
type SubtitleDataDetail struct {
	SeasonNumber NullableInt32 `json:"seasonNumber,omitempty"`
	Id NullableString `json:"id,omitempty"`
	Rate NullableString `json:"rate,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Owner NullableString `json:"owner,omitempty"`
	Comment NullableString `json:"comment,omitempty"`
	Link NullableString `json:"link,omitempty"`
}

// NewSubtitleDataDetail instantiates a new SubtitleDataDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubtitleDataDetail() *SubtitleDataDetail {
	this := SubtitleDataDetail{}
	return &this
}

// NewSubtitleDataDetailWithDefaults instantiates a new SubtitleDataDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubtitleDataDetailWithDefaults() *SubtitleDataDetail {
	this := SubtitleDataDetail{}
	return &this
}

// GetSeasonNumber returns the SeasonNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubtitleDataDetail) GetSeasonNumber() int32 {
	if o == nil || o.SeasonNumber.Get() == nil {
		var ret int32
		return ret
	}
	return *o.SeasonNumber.Get()
}

// GetSeasonNumberOk returns a tuple with the SeasonNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubtitleDataDetail) GetSeasonNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeasonNumber.Get(), o.SeasonNumber.IsSet()
}

// HasSeasonNumber returns a boolean if a field has been set.
func (o *SubtitleDataDetail) HasSeasonNumber() bool {
	if o != nil && o.SeasonNumber.IsSet() {
		return true
	}

	return false
}

// SetSeasonNumber gets a reference to the given NullableInt32 and assigns it to the SeasonNumber field.
func (o *SubtitleDataDetail) SetSeasonNumber(v int32) {
	o.SeasonNumber.Set(&v)
}
// SetSeasonNumberNil sets the value for SeasonNumber to be an explicit nil
func (o *SubtitleDataDetail) SetSeasonNumberNil() {
	o.SeasonNumber.Set(nil)
}

// UnsetSeasonNumber ensures that no value is present for SeasonNumber, not even an explicit nil
func (o *SubtitleDataDetail) UnsetSeasonNumber() {
	o.SeasonNumber.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubtitleDataDetail) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubtitleDataDetail) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *SubtitleDataDetail) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *SubtitleDataDetail) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *SubtitleDataDetail) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *SubtitleDataDetail) UnsetId() {
	o.Id.Unset()
}

// GetRate returns the Rate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubtitleDataDetail) GetRate() string {
	if o == nil || o.Rate.Get() == nil {
		var ret string
		return ret
	}
	return *o.Rate.Get()
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubtitleDataDetail) GetRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rate.Get(), o.Rate.IsSet()
}

// HasRate returns a boolean if a field has been set.
func (o *SubtitleDataDetail) HasRate() bool {
	if o != nil && o.Rate.IsSet() {
		return true
	}

	return false
}

// SetRate gets a reference to the given NullableString and assigns it to the Rate field.
func (o *SubtitleDataDetail) SetRate(v string) {
	o.Rate.Set(&v)
}
// SetRateNil sets the value for Rate to be an explicit nil
func (o *SubtitleDataDetail) SetRateNil() {
	o.Rate.Set(nil)
}

// UnsetRate ensures that no value is present for Rate, not even an explicit nil
func (o *SubtitleDataDetail) UnsetRate() {
	o.Rate.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubtitleDataDetail) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubtitleDataDetail) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *SubtitleDataDetail) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *SubtitleDataDetail) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *SubtitleDataDetail) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *SubtitleDataDetail) UnsetTitle() {
	o.Title.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubtitleDataDetail) GetOwner() string {
	if o == nil || o.Owner.Get() == nil {
		var ret string
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubtitleDataDetail) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *SubtitleDataDetail) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableString and assigns it to the Owner field.
func (o *SubtitleDataDetail) SetOwner(v string) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *SubtitleDataDetail) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *SubtitleDataDetail) UnsetOwner() {
	o.Owner.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubtitleDataDetail) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubtitleDataDetail) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *SubtitleDataDetail) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *SubtitleDataDetail) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *SubtitleDataDetail) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *SubtitleDataDetail) UnsetComment() {
	o.Comment.Unset()
}

// GetLink returns the Link field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubtitleDataDetail) GetLink() string {
	if o == nil || o.Link.Get() == nil {
		var ret string
		return ret
	}
	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubtitleDataDetail) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// HasLink returns a boolean if a field has been set.
func (o *SubtitleDataDetail) HasLink() bool {
	if o != nil && o.Link.IsSet() {
		return true
	}

	return false
}

// SetLink gets a reference to the given NullableString and assigns it to the Link field.
func (o *SubtitleDataDetail) SetLink(v string) {
	o.Link.Set(&v)
}
// SetLinkNil sets the value for Link to be an explicit nil
func (o *SubtitleDataDetail) SetLinkNil() {
	o.Link.Set(nil)
}

// UnsetLink ensures that no value is present for Link, not even an explicit nil
func (o *SubtitleDataDetail) UnsetLink() {
	o.Link.Unset()
}

func (o SubtitleDataDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SeasonNumber.IsSet() {
		toSerialize["seasonNumber"] = o.SeasonNumber.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Rate.IsSet() {
		toSerialize["rate"] = o.Rate.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Link.IsSet() {
		toSerialize["link"] = o.Link.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSubtitleDataDetail struct {
	value *SubtitleDataDetail
	isSet bool
}

func (v NullableSubtitleDataDetail) Get() *SubtitleDataDetail {
	return v.value
}

func (v *NullableSubtitleDataDetail) Set(val *SubtitleDataDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableSubtitleDataDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableSubtitleDataDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubtitleDataDetail(val *SubtitleDataDetail) *NullableSubtitleDataDetail {
	return &NullableSubtitleDataDetail{value: val, isSet: true}
}

func (v NullableSubtitleDataDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubtitleDataDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


