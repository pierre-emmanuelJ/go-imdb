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

// YouTubePlaylistDataItem struct for YouTubePlaylistDataItem
type YouTubePlaylistDataItem struct {
	VideoId NullableString `json:"videoId,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Duration NullableString `json:"duration,omitempty"`
	UploadDate NullableString `json:"uploadDate,omitempty"`
	Image NullableString `json:"image,omitempty"`
	Url NullableString `json:"url,omitempty"`
}

// NewYouTubePlaylistDataItem instantiates a new YouTubePlaylistDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewYouTubePlaylistDataItem() *YouTubePlaylistDataItem {
	this := YouTubePlaylistDataItem{}
	return &this
}

// NewYouTubePlaylistDataItemWithDefaults instantiates a new YouTubePlaylistDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewYouTubePlaylistDataItemWithDefaults() *YouTubePlaylistDataItem {
	this := YouTubePlaylistDataItem{}
	return &this
}

// GetVideoId returns the VideoId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *YouTubePlaylistDataItem) GetVideoId() string {
	if o == nil || o.VideoId.Get() == nil {
		var ret string
		return ret
	}
	return *o.VideoId.Get()
}

// GetVideoIdOk returns a tuple with the VideoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *YouTubePlaylistDataItem) GetVideoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoId.Get(), o.VideoId.IsSet()
}

// HasVideoId returns a boolean if a field has been set.
func (o *YouTubePlaylistDataItem) HasVideoId() bool {
	if o != nil && o.VideoId.IsSet() {
		return true
	}

	return false
}

// SetVideoId gets a reference to the given NullableString and assigns it to the VideoId field.
func (o *YouTubePlaylistDataItem) SetVideoId(v string) {
	o.VideoId.Set(&v)
}
// SetVideoIdNil sets the value for VideoId to be an explicit nil
func (o *YouTubePlaylistDataItem) SetVideoIdNil() {
	o.VideoId.Set(nil)
}

// UnsetVideoId ensures that no value is present for VideoId, not even an explicit nil
func (o *YouTubePlaylistDataItem) UnsetVideoId() {
	o.VideoId.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *YouTubePlaylistDataItem) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *YouTubePlaylistDataItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *YouTubePlaylistDataItem) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *YouTubePlaylistDataItem) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *YouTubePlaylistDataItem) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *YouTubePlaylistDataItem) UnsetTitle() {
	o.Title.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *YouTubePlaylistDataItem) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *YouTubePlaylistDataItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *YouTubePlaylistDataItem) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *YouTubePlaylistDataItem) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *YouTubePlaylistDataItem) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *YouTubePlaylistDataItem) UnsetDescription() {
	o.Description.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *YouTubePlaylistDataItem) GetDuration() string {
	if o == nil || o.Duration.Get() == nil {
		var ret string
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *YouTubePlaylistDataItem) GetDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *YouTubePlaylistDataItem) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableString and assigns it to the Duration field.
func (o *YouTubePlaylistDataItem) SetDuration(v string) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *YouTubePlaylistDataItem) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *YouTubePlaylistDataItem) UnsetDuration() {
	o.Duration.Unset()
}

// GetUploadDate returns the UploadDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *YouTubePlaylistDataItem) GetUploadDate() string {
	if o == nil || o.UploadDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.UploadDate.Get()
}

// GetUploadDateOk returns a tuple with the UploadDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *YouTubePlaylistDataItem) GetUploadDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UploadDate.Get(), o.UploadDate.IsSet()
}

// HasUploadDate returns a boolean if a field has been set.
func (o *YouTubePlaylistDataItem) HasUploadDate() bool {
	if o != nil && o.UploadDate.IsSet() {
		return true
	}

	return false
}

// SetUploadDate gets a reference to the given NullableString and assigns it to the UploadDate field.
func (o *YouTubePlaylistDataItem) SetUploadDate(v string) {
	o.UploadDate.Set(&v)
}
// SetUploadDateNil sets the value for UploadDate to be an explicit nil
func (o *YouTubePlaylistDataItem) SetUploadDateNil() {
	o.UploadDate.Set(nil)
}

// UnsetUploadDate ensures that no value is present for UploadDate, not even an explicit nil
func (o *YouTubePlaylistDataItem) UnsetUploadDate() {
	o.UploadDate.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *YouTubePlaylistDataItem) GetImage() string {
	if o == nil || o.Image.Get() == nil {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *YouTubePlaylistDataItem) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *YouTubePlaylistDataItem) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *YouTubePlaylistDataItem) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *YouTubePlaylistDataItem) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *YouTubePlaylistDataItem) UnsetImage() {
	o.Image.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *YouTubePlaylistDataItem) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *YouTubePlaylistDataItem) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *YouTubePlaylistDataItem) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *YouTubePlaylistDataItem) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *YouTubePlaylistDataItem) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *YouTubePlaylistDataItem) UnsetUrl() {
	o.Url.Unset()
}

func (o YouTubePlaylistDataItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VideoId.IsSet() {
		toSerialize["videoId"] = o.VideoId.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if o.UploadDate.IsSet() {
		toSerialize["uploadDate"] = o.UploadDate.Get()
	}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableYouTubePlaylistDataItem struct {
	value *YouTubePlaylistDataItem
	isSet bool
}

func (v NullableYouTubePlaylistDataItem) Get() *YouTubePlaylistDataItem {
	return v.value
}

func (v *NullableYouTubePlaylistDataItem) Set(val *YouTubePlaylistDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableYouTubePlaylistDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableYouTubePlaylistDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableYouTubePlaylistDataItem(val *YouTubePlaylistDataItem) *NullableYouTubePlaylistDataItem {
	return &NullableYouTubePlaylistDataItem{value: val, isSet: true}
}

func (v NullableYouTubePlaylistDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableYouTubePlaylistDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


