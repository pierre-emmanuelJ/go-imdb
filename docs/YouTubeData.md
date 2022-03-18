# YouTubeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VideoId** | **string** |  | 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Duration** | Pointer to **NullableString** |  | [optional] 
**UploadDate** | Pointer to **NullableString** |  | [optional] 
**Image** | Pointer to **NullableString** |  | [optional] 
**Videos** | Pointer to [**[]YouTubeDataItem**](YouTubeDataItem.md) |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewYouTubeData

`func NewYouTubeData(videoId string, ) *YouTubeData`

NewYouTubeData instantiates a new YouTubeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYouTubeDataWithDefaults

`func NewYouTubeDataWithDefaults() *YouTubeData`

NewYouTubeDataWithDefaults instantiates a new YouTubeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVideoId

`func (o *YouTubeData) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *YouTubeData) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *YouTubeData) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.


### GetTitle

`func (o *YouTubeData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *YouTubeData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *YouTubeData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *YouTubeData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *YouTubeData) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *YouTubeData) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *YouTubeData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *YouTubeData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *YouTubeData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *YouTubeData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *YouTubeData) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *YouTubeData) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDuration

`func (o *YouTubeData) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *YouTubeData) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *YouTubeData) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *YouTubeData) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *YouTubeData) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *YouTubeData) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetUploadDate

`func (o *YouTubeData) GetUploadDate() string`

GetUploadDate returns the UploadDate field if non-nil, zero value otherwise.

### GetUploadDateOk

`func (o *YouTubeData) GetUploadDateOk() (*string, bool)`

GetUploadDateOk returns a tuple with the UploadDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadDate

`func (o *YouTubeData) SetUploadDate(v string)`

SetUploadDate sets UploadDate field to given value.

### HasUploadDate

`func (o *YouTubeData) HasUploadDate() bool`

HasUploadDate returns a boolean if a field has been set.

### SetUploadDateNil

`func (o *YouTubeData) SetUploadDateNil(b bool)`

 SetUploadDateNil sets the value for UploadDate to be an explicit nil

### UnsetUploadDate
`func (o *YouTubeData) UnsetUploadDate()`

UnsetUploadDate ensures that no value is present for UploadDate, not even an explicit nil
### GetImage

`func (o *YouTubeData) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *YouTubeData) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *YouTubeData) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *YouTubeData) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *YouTubeData) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *YouTubeData) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetVideos

`func (o *YouTubeData) GetVideos() []YouTubeDataItem`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *YouTubeData) GetVideosOk() (*[]YouTubeDataItem, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *YouTubeData) SetVideos(v []YouTubeDataItem)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *YouTubeData) HasVideos() bool`

HasVideos returns a boolean if a field has been set.

### SetVideosNil

`func (o *YouTubeData) SetVideosNil(b bool)`

 SetVideosNil sets the value for Videos to be an explicit nil

### UnsetVideos
`func (o *YouTubeData) UnsetVideos()`

UnsetVideos ensures that no value is present for Videos, not even an explicit nil
### GetErrorMessage

`func (o *YouTubeData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *YouTubeData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *YouTubeData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *YouTubeData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *YouTubeData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *YouTubeData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


