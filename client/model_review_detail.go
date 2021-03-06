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

// ReviewDetail struct for ReviewDetail
type ReviewDetail struct {
	Username NullableString `json:"username,omitempty"`
	UserUrl NullableString `json:"userUrl,omitempty"`
	ReviewLink NullableString `json:"reviewLink,omitempty"`
	WarningSpoilers *bool `json:"warningSpoilers,omitempty"`
	Date NullableString `json:"date,omitempty"`
	Rate NullableString `json:"rate,omitempty"`
	Helpful NullableString `json:"helpful,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Content NullableString `json:"content,omitempty"`
}

// NewReviewDetail instantiates a new ReviewDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewDetail() *ReviewDetail {
	this := ReviewDetail{}
	return &this
}

// NewReviewDetailWithDefaults instantiates a new ReviewDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewDetailWithDefaults() *ReviewDetail {
	this := ReviewDetail{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewDetail) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewDetail) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *ReviewDetail) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *ReviewDetail) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *ReviewDetail) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *ReviewDetail) UnsetUsername() {
	o.Username.Unset()
}

// GetUserUrl returns the UserUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewDetail) GetUserUrl() string {
	if o == nil || o.UserUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserUrl.Get()
}

// GetUserUrlOk returns a tuple with the UserUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewDetail) GetUserUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserUrl.Get(), o.UserUrl.IsSet()
}

// HasUserUrl returns a boolean if a field has been set.
func (o *ReviewDetail) HasUserUrl() bool {
	if o != nil && o.UserUrl.IsSet() {
		return true
	}

	return false
}

// SetUserUrl gets a reference to the given NullableString and assigns it to the UserUrl field.
func (o *ReviewDetail) SetUserUrl(v string) {
	o.UserUrl.Set(&v)
}
// SetUserUrlNil sets the value for UserUrl to be an explicit nil
func (o *ReviewDetail) SetUserUrlNil() {
	o.UserUrl.Set(nil)
}

// UnsetUserUrl ensures that no value is present for UserUrl, not even an explicit nil
func (o *ReviewDetail) UnsetUserUrl() {
	o.UserUrl.Unset()
}

// GetReviewLink returns the ReviewLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewDetail) GetReviewLink() string {
	if o == nil || o.ReviewLink.Get() == nil {
		var ret string
		return ret
	}
	return *o.ReviewLink.Get()
}

// GetReviewLinkOk returns a tuple with the ReviewLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewDetail) GetReviewLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReviewLink.Get(), o.ReviewLink.IsSet()
}

// HasReviewLink returns a boolean if a field has been set.
func (o *ReviewDetail) HasReviewLink() bool {
	if o != nil && o.ReviewLink.IsSet() {
		return true
	}

	return false
}

// SetReviewLink gets a reference to the given NullableString and assigns it to the ReviewLink field.
func (o *ReviewDetail) SetReviewLink(v string) {
	o.ReviewLink.Set(&v)
}
// SetReviewLinkNil sets the value for ReviewLink to be an explicit nil
func (o *ReviewDetail) SetReviewLinkNil() {
	o.ReviewLink.Set(nil)
}

// UnsetReviewLink ensures that no value is present for ReviewLink, not even an explicit nil
func (o *ReviewDetail) UnsetReviewLink() {
	o.ReviewLink.Unset()
}

// GetWarningSpoilers returns the WarningSpoilers field value if set, zero value otherwise.
func (o *ReviewDetail) GetWarningSpoilers() bool {
	if o == nil || o.WarningSpoilers == nil {
		var ret bool
		return ret
	}
	return *o.WarningSpoilers
}

// GetWarningSpoilersOk returns a tuple with the WarningSpoilers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewDetail) GetWarningSpoilersOk() (*bool, bool) {
	if o == nil || o.WarningSpoilers == nil {
		return nil, false
	}
	return o.WarningSpoilers, true
}

// HasWarningSpoilers returns a boolean if a field has been set.
func (o *ReviewDetail) HasWarningSpoilers() bool {
	if o != nil && o.WarningSpoilers != nil {
		return true
	}

	return false
}

// SetWarningSpoilers gets a reference to the given bool and assigns it to the WarningSpoilers field.
func (o *ReviewDetail) SetWarningSpoilers(v bool) {
	o.WarningSpoilers = &v
}

// GetDate returns the Date field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewDetail) GetDate() string {
	if o == nil || o.Date.Get() == nil {
		var ret string
		return ret
	}
	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewDetail) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// HasDate returns a boolean if a field has been set.
func (o *ReviewDetail) HasDate() bool {
	if o != nil && o.Date.IsSet() {
		return true
	}

	return false
}

// SetDate gets a reference to the given NullableString and assigns it to the Date field.
func (o *ReviewDetail) SetDate(v string) {
	o.Date.Set(&v)
}
// SetDateNil sets the value for Date to be an explicit nil
func (o *ReviewDetail) SetDateNil() {
	o.Date.Set(nil)
}

// UnsetDate ensures that no value is present for Date, not even an explicit nil
func (o *ReviewDetail) UnsetDate() {
	o.Date.Unset()
}

// GetRate returns the Rate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewDetail) GetRate() string {
	if o == nil || o.Rate.Get() == nil {
		var ret string
		return ret
	}
	return *o.Rate.Get()
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewDetail) GetRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rate.Get(), o.Rate.IsSet()
}

// HasRate returns a boolean if a field has been set.
func (o *ReviewDetail) HasRate() bool {
	if o != nil && o.Rate.IsSet() {
		return true
	}

	return false
}

// SetRate gets a reference to the given NullableString and assigns it to the Rate field.
func (o *ReviewDetail) SetRate(v string) {
	o.Rate.Set(&v)
}
// SetRateNil sets the value for Rate to be an explicit nil
func (o *ReviewDetail) SetRateNil() {
	o.Rate.Set(nil)
}

// UnsetRate ensures that no value is present for Rate, not even an explicit nil
func (o *ReviewDetail) UnsetRate() {
	o.Rate.Unset()
}

// GetHelpful returns the Helpful field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewDetail) GetHelpful() string {
	if o == nil || o.Helpful.Get() == nil {
		var ret string
		return ret
	}
	return *o.Helpful.Get()
}

// GetHelpfulOk returns a tuple with the Helpful field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewDetail) GetHelpfulOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Helpful.Get(), o.Helpful.IsSet()
}

// HasHelpful returns a boolean if a field has been set.
func (o *ReviewDetail) HasHelpful() bool {
	if o != nil && o.Helpful.IsSet() {
		return true
	}

	return false
}

// SetHelpful gets a reference to the given NullableString and assigns it to the Helpful field.
func (o *ReviewDetail) SetHelpful(v string) {
	o.Helpful.Set(&v)
}
// SetHelpfulNil sets the value for Helpful to be an explicit nil
func (o *ReviewDetail) SetHelpfulNil() {
	o.Helpful.Set(nil)
}

// UnsetHelpful ensures that no value is present for Helpful, not even an explicit nil
func (o *ReviewDetail) UnsetHelpful() {
	o.Helpful.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewDetail) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewDetail) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ReviewDetail) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ReviewDetail) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ReviewDetail) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ReviewDetail) UnsetTitle() {
	o.Title.Unset()
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewDetail) GetContent() string {
	if o == nil || o.Content.Get() == nil {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewDetail) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *ReviewDetail) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *ReviewDetail) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *ReviewDetail) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *ReviewDetail) UnsetContent() {
	o.Content.Unset()
}

func (o ReviewDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.UserUrl.IsSet() {
		toSerialize["userUrl"] = o.UserUrl.Get()
	}
	if o.ReviewLink.IsSet() {
		toSerialize["reviewLink"] = o.ReviewLink.Get()
	}
	if o.WarningSpoilers != nil {
		toSerialize["warningSpoilers"] = o.WarningSpoilers
	}
	if o.Date.IsSet() {
		toSerialize["date"] = o.Date.Get()
	}
	if o.Rate.IsSet() {
		toSerialize["rate"] = o.Rate.Get()
	}
	if o.Helpful.IsSet() {
		toSerialize["helpful"] = o.Helpful.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableReviewDetail struct {
	value *ReviewDetail
	isSet bool
}

func (v NullableReviewDetail) Get() *ReviewDetail {
	return v.value
}

func (v *NullableReviewDetail) Set(val *ReviewDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewDetail(val *ReviewDetail) *NullableReviewDetail {
	return &NullableReviewDetail{value: val, isSet: true}
}

func (v NullableReviewDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


