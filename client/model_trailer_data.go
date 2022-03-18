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

// TrailerData struct for TrailerData
type TrailerData struct {
	ImDbId NullableString `json:"imDbId,omitempty"`
	Title NullableString `json:"title,omitempty"`
	FullTitle NullableString `json:"fullTitle,omitempty"`
	Type NullableString `json:"type,omitempty"`
	Year NullableString `json:"year,omitempty"`
	VideoId NullableString `json:"videoId,omitempty"`
	VideoTitle NullableString `json:"videoTitle,omitempty"`
	VideoDescription NullableString `json:"videoDescription,omitempty"`
	ThumbnailUrl NullableString `json:"thumbnailUrl,omitempty"`
	UploadDate NullableString `json:"uploadDate,omitempty"`
	Link NullableString `json:"link,omitempty"`
	LinkEmbed NullableString `json:"linkEmbed,omitempty"`
	ErrorMessage NullableString `json:"errorMessage,omitempty"`
}

// NewTrailerData instantiates a new TrailerData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrailerData() *TrailerData {
	this := TrailerData{}
	return &this
}

// NewTrailerDataWithDefaults instantiates a new TrailerData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrailerDataWithDefaults() *TrailerData {
	this := TrailerData{}
	return &this
}

// GetImDbId returns the ImDbId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerData) GetImDbId() string {
	if o == nil || o.ImDbId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ImDbId.Get()
}

// GetImDbIdOk returns a tuple with the ImDbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerData) GetImDbIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImDbId.Get(), o.ImDbId.IsSet()
}

// HasImDbId returns a boolean if a field has been set.
func (o *TrailerData) HasImDbId() bool {
	if o != nil && o.ImDbId.IsSet() {
		return true
	}

	return false
}

// SetImDbId gets a reference to the given NullableString and assigns it to the ImDbId field.
func (o *TrailerData) SetImDbId(v string) {
	o.ImDbId.Set(&v)
}
// SetImDbIdNil sets the value for ImDbId to be an explicit nil
func (o *TrailerData) SetImDbIdNil() {
	o.ImDbId.Set(nil)
}

// UnsetImDbId ensures that no value is present for ImDbId, not even an explicit nil
func (o *TrailerData) UnsetImDbId() {
	o.ImDbId.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerData) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerData) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *TrailerData) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *TrailerData) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *TrailerData) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *TrailerData) UnsetTitle() {
	o.Title.Unset()
}

// GetFullTitle returns the FullTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerData) GetFullTitle() string {
	if o == nil || o.FullTitle.Get() == nil {
		var ret string
		return ret
	}
	return *o.FullTitle.Get()
}

// GetFullTitleOk returns a tuple with the FullTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerData) GetFullTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FullTitle.Get(), o.FullTitle.IsSet()
}

// HasFullTitle returns a boolean if a field has been set.
func (o *TrailerData) HasFullTitle() bool {
	if o != nil && o.FullTitle.IsSet() {
		return true
	}

	return false
}

// SetFullTitle gets a reference to the given NullableString and assigns it to the FullTitle field.
func (o *TrailerData) SetFullTitle(v string) {
	o.FullTitle.Set(&v)
}
// SetFullTitleNil sets the value for FullTitle to be an explicit nil
func (o *TrailerData) SetFullTitleNil() {
	o.FullTitle.Set(nil)
}

// UnsetFullTitle ensures that no value is present for FullTitle, not even an explicit nil
func (o *TrailerData) UnsetFullTitle() {
	o.FullTitle.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerData) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *TrailerData) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *TrailerData) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *TrailerData) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *TrailerData) UnsetType() {
	o.Type.Unset()
}

// GetYear returns the Year field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerData) GetYear() string {
	if o == nil || o.Year.Get() == nil {
		var ret string
		return ret
	}
	return *o.Year.Get()
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerData) GetYearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Year.Get(), o.Year.IsSet()
}

// HasYear returns a boolean if a field has been set.
func (o *TrailerData) HasYear() bool {
	if o != nil && o.Year.IsSet() {
		return true
	}

	return false
}

// SetYear gets a reference to the given NullableString and assigns it to the Year field.
func (o *TrailerData) SetYear(v string) {
	o.Year.Set(&v)
}
// SetYearNil sets the value for Year to be an explicit nil
func (o *TrailerData) SetYearNil() {
	o.Year.Set(nil)
}

// UnsetYear ensures that no value is present for Year, not even an explicit nil
func (o *TrailerData) UnsetYear() {
	o.Year.Unset()
}

// GetVideoId returns the VideoId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerData) GetVideoId() string {
	if o == nil || o.VideoId.Get() == nil {
		var ret string
		return ret
	}
	return *o.VideoId.Get()
}

// GetVideoIdOk returns a tuple with the VideoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerData) GetVideoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoId.Get(), o.VideoId.IsSet()
}

// HasVideoId returns a boolean if a field has been set.
func (o *TrailerData) HasVideoId() bool {
	if o != nil && o.VideoId.IsSet() {
		return true
	}

	return false
}

// SetVideoId gets a reference to the given NullableString and assigns it to the VideoId field.
func (o *TrailerData) SetVideoId(v string) {
	o.VideoId.Set(&v)
}
// SetVideoIdNil sets the value for VideoId to be an explicit nil
func (o *TrailerData) SetVideoIdNil() {
	o.VideoId.Set(nil)
}

// UnsetVideoId ensures that no value is present for VideoId, not even an explicit nil
func (o *TrailerData) UnsetVideoId() {
	o.VideoId.Unset()
}

// GetVideoTitle returns the VideoTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerData) GetVideoTitle() string {
	if o == nil || o.VideoTitle.Get() == nil {
		var ret string
		return ret
	}
	return *o.VideoTitle.Get()
}

// GetVideoTitleOk returns a tuple with the VideoTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerData) GetVideoTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoTitle.Get(), o.VideoTitle.IsSet()
}

// HasVideoTitle returns a boolean if a field has been set.
func (o *TrailerData) HasVideoTitle() bool {
	if o != nil && o.VideoTitle.IsSet() {
		return true
	}

	return false
}

// SetVideoTitle gets a reference to the given NullableString and assigns it to the VideoTitle field.
func (o *TrailerData) SetVideoTitle(v string) {
	o.VideoTitle.Set(&v)
}
// SetVideoTitleNil sets the value for VideoTitle to be an explicit nil
func (o *TrailerData) SetVideoTitleNil() {
	o.VideoTitle.Set(nil)
}

// UnsetVideoTitle ensures that no value is present for VideoTitle, not even an explicit nil
func (o *TrailerData) UnsetVideoTitle() {
	o.VideoTitle.Unset()
}

// GetVideoDescription returns the VideoDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerData) GetVideoDescription() string {
	if o == nil || o.VideoDescription.Get() == nil {
		var ret string
		return ret
	}
	return *o.VideoDescription.Get()
}

// GetVideoDescriptionOk returns a tuple with the VideoDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerData) GetVideoDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoDescription.Get(), o.VideoDescription.IsSet()
}

// HasVideoDescription returns a boolean if a field has been set.
func (o *TrailerData) HasVideoDescription() bool {
	if o != nil && o.VideoDescription.IsSet() {
		return true
	}

	return false
}

// SetVideoDescription gets a reference to the given NullableString and assigns it to the VideoDescription field.
func (o *TrailerData) SetVideoDescription(v string) {
	o.VideoDescription.Set(&v)
}
// SetVideoDescriptionNil sets the value for VideoDescription to be an explicit nil
func (o *TrailerData) SetVideoDescriptionNil() {
	o.VideoDescription.Set(nil)
}

// UnsetVideoDescription ensures that no value is present for VideoDescription, not even an explicit nil
func (o *TrailerData) UnsetVideoDescription() {
	o.VideoDescription.Unset()
}

// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerData) GetThumbnailUrl() string {
	if o == nil || o.ThumbnailUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl.Get()
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerData) GetThumbnailUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThumbnailUrl.Get(), o.ThumbnailUrl.IsSet()
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *TrailerData) HasThumbnailUrl() bool {
	if o != nil && o.ThumbnailUrl.IsSet() {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given NullableString and assigns it to the ThumbnailUrl field.
func (o *TrailerData) SetThumbnailUrl(v string) {
	o.ThumbnailUrl.Set(&v)
}
// SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil
func (o *TrailerData) SetThumbnailUrlNil() {
	o.ThumbnailUrl.Set(nil)
}

// UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
func (o *TrailerData) UnsetThumbnailUrl() {
	o.ThumbnailUrl.Unset()
}

// GetUploadDate returns the UploadDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerData) GetUploadDate() string {
	if o == nil || o.UploadDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.UploadDate.Get()
}

// GetUploadDateOk returns a tuple with the UploadDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerData) GetUploadDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UploadDate.Get(), o.UploadDate.IsSet()
}

// HasUploadDate returns a boolean if a field has been set.
func (o *TrailerData) HasUploadDate() bool {
	if o != nil && o.UploadDate.IsSet() {
		return true
	}

	return false
}

// SetUploadDate gets a reference to the given NullableString and assigns it to the UploadDate field.
func (o *TrailerData) SetUploadDate(v string) {
	o.UploadDate.Set(&v)
}
// SetUploadDateNil sets the value for UploadDate to be an explicit nil
func (o *TrailerData) SetUploadDateNil() {
	o.UploadDate.Set(nil)
}

// UnsetUploadDate ensures that no value is present for UploadDate, not even an explicit nil
func (o *TrailerData) UnsetUploadDate() {
	o.UploadDate.Unset()
}

// GetLink returns the Link field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerData) GetLink() string {
	if o == nil || o.Link.Get() == nil {
		var ret string
		return ret
	}
	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerData) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// HasLink returns a boolean if a field has been set.
func (o *TrailerData) HasLink() bool {
	if o != nil && o.Link.IsSet() {
		return true
	}

	return false
}

// SetLink gets a reference to the given NullableString and assigns it to the Link field.
func (o *TrailerData) SetLink(v string) {
	o.Link.Set(&v)
}
// SetLinkNil sets the value for Link to be an explicit nil
func (o *TrailerData) SetLinkNil() {
	o.Link.Set(nil)
}

// UnsetLink ensures that no value is present for Link, not even an explicit nil
func (o *TrailerData) UnsetLink() {
	o.Link.Unset()
}

// GetLinkEmbed returns the LinkEmbed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerData) GetLinkEmbed() string {
	if o == nil || o.LinkEmbed.Get() == nil {
		var ret string
		return ret
	}
	return *o.LinkEmbed.Get()
}

// GetLinkEmbedOk returns a tuple with the LinkEmbed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerData) GetLinkEmbedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkEmbed.Get(), o.LinkEmbed.IsSet()
}

// HasLinkEmbed returns a boolean if a field has been set.
func (o *TrailerData) HasLinkEmbed() bool {
	if o != nil && o.LinkEmbed.IsSet() {
		return true
	}

	return false
}

// SetLinkEmbed gets a reference to the given NullableString and assigns it to the LinkEmbed field.
func (o *TrailerData) SetLinkEmbed(v string) {
	o.LinkEmbed.Set(&v)
}
// SetLinkEmbedNil sets the value for LinkEmbed to be an explicit nil
func (o *TrailerData) SetLinkEmbedNil() {
	o.LinkEmbed.Set(nil)
}

// UnsetLinkEmbed ensures that no value is present for LinkEmbed, not even an explicit nil
func (o *TrailerData) UnsetLinkEmbed() {
	o.LinkEmbed.Unset()
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerData) GetErrorMessage() string {
	if o == nil || o.ErrorMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerData) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *TrailerData) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *TrailerData) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *TrailerData) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *TrailerData) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

func (o TrailerData) MarshalJSON() ([]byte, error) {
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
	if o.VideoId.IsSet() {
		toSerialize["videoId"] = o.VideoId.Get()
	}
	if o.VideoTitle.IsSet() {
		toSerialize["videoTitle"] = o.VideoTitle.Get()
	}
	if o.VideoDescription.IsSet() {
		toSerialize["videoDescription"] = o.VideoDescription.Get()
	}
	if o.ThumbnailUrl.IsSet() {
		toSerialize["thumbnailUrl"] = o.ThumbnailUrl.Get()
	}
	if o.UploadDate.IsSet() {
		toSerialize["uploadDate"] = o.UploadDate.Get()
	}
	if o.Link.IsSet() {
		toSerialize["link"] = o.Link.Get()
	}
	if o.LinkEmbed.IsSet() {
		toSerialize["linkEmbed"] = o.LinkEmbed.Get()
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTrailerData struct {
	value *TrailerData
	isSet bool
}

func (v NullableTrailerData) Get() *TrailerData {
	return v.value
}

func (v *NullableTrailerData) Set(val *TrailerData) {
	v.value = val
	v.isSet = true
}

func (v NullableTrailerData) IsSet() bool {
	return v.isSet
}

func (v *NullableTrailerData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrailerData(val *TrailerData) *NullableTrailerData {
	return &NullableTrailerData{value: val, isSet: true}
}

func (v NullableTrailerData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrailerData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

