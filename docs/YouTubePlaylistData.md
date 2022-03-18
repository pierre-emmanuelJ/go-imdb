# YouTubePlaylistData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Auhtor** | Pointer to **NullableString** |  | [optional] 
**Videos** | Pointer to [**[]YouTubePlaylistDataItem**](YouTubePlaylistDataItem.md) |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewYouTubePlaylistData

`func NewYouTubePlaylistData() *YouTubePlaylistData`

NewYouTubePlaylistData instantiates a new YouTubePlaylistData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYouTubePlaylistDataWithDefaults

`func NewYouTubePlaylistDataWithDefaults() *YouTubePlaylistData`

NewYouTubePlaylistDataWithDefaults instantiates a new YouTubePlaylistData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *YouTubePlaylistData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *YouTubePlaylistData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *YouTubePlaylistData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *YouTubePlaylistData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *YouTubePlaylistData) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *YouTubePlaylistData) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAuhtor

`func (o *YouTubePlaylistData) GetAuhtor() string`

GetAuhtor returns the Auhtor field if non-nil, zero value otherwise.

### GetAuhtorOk

`func (o *YouTubePlaylistData) GetAuhtorOk() (*string, bool)`

GetAuhtorOk returns a tuple with the Auhtor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuhtor

`func (o *YouTubePlaylistData) SetAuhtor(v string)`

SetAuhtor sets Auhtor field to given value.

### HasAuhtor

`func (o *YouTubePlaylistData) HasAuhtor() bool`

HasAuhtor returns a boolean if a field has been set.

### SetAuhtorNil

`func (o *YouTubePlaylistData) SetAuhtorNil(b bool)`

 SetAuhtorNil sets the value for Auhtor to be an explicit nil

### UnsetAuhtor
`func (o *YouTubePlaylistData) UnsetAuhtor()`

UnsetAuhtor ensures that no value is present for Auhtor, not even an explicit nil
### GetVideos

`func (o *YouTubePlaylistData) GetVideos() []YouTubePlaylistDataItem`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *YouTubePlaylistData) GetVideosOk() (*[]YouTubePlaylistDataItem, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *YouTubePlaylistData) SetVideos(v []YouTubePlaylistDataItem)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *YouTubePlaylistData) HasVideos() bool`

HasVideos returns a boolean if a field has been set.

### SetVideosNil

`func (o *YouTubePlaylistData) SetVideosNil(b bool)`

 SetVideosNil sets the value for Videos to be an explicit nil

### UnsetVideos
`func (o *YouTubePlaylistData) UnsetVideos()`

UnsetVideos ensures that no value is present for Videos, not even an explicit nil
### GetErrorMessage

`func (o *YouTubePlaylistData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *YouTubePlaylistData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *YouTubePlaylistData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *YouTubePlaylistData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *YouTubePlaylistData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *YouTubePlaylistData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


