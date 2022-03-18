# SubtitleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImDbId** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**FullTitle** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Year** | Pointer to **NullableString** |  | [optional] 
**Subtitles** | Pointer to [**[]SubtitleDataDetail**](SubtitleDataDetail.md) |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSubtitleData

`func NewSubtitleData() *SubtitleData`

NewSubtitleData instantiates a new SubtitleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubtitleDataWithDefaults

`func NewSubtitleDataWithDefaults() *SubtitleData`

NewSubtitleDataWithDefaults instantiates a new SubtitleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImDbId

`func (o *SubtitleData) GetImDbId() string`

GetImDbId returns the ImDbId field if non-nil, zero value otherwise.

### GetImDbIdOk

`func (o *SubtitleData) GetImDbIdOk() (*string, bool)`

GetImDbIdOk returns a tuple with the ImDbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImDbId

`func (o *SubtitleData) SetImDbId(v string)`

SetImDbId sets ImDbId field to given value.

### HasImDbId

`func (o *SubtitleData) HasImDbId() bool`

HasImDbId returns a boolean if a field has been set.

### SetImDbIdNil

`func (o *SubtitleData) SetImDbIdNil(b bool)`

 SetImDbIdNil sets the value for ImDbId to be an explicit nil

### UnsetImDbId
`func (o *SubtitleData) UnsetImDbId()`

UnsetImDbId ensures that no value is present for ImDbId, not even an explicit nil
### GetTitle

`func (o *SubtitleData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SubtitleData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SubtitleData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SubtitleData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SubtitleData) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SubtitleData) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetFullTitle

`func (o *SubtitleData) GetFullTitle() string`

GetFullTitle returns the FullTitle field if non-nil, zero value otherwise.

### GetFullTitleOk

`func (o *SubtitleData) GetFullTitleOk() (*string, bool)`

GetFullTitleOk returns a tuple with the FullTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTitle

`func (o *SubtitleData) SetFullTitle(v string)`

SetFullTitle sets FullTitle field to given value.

### HasFullTitle

`func (o *SubtitleData) HasFullTitle() bool`

HasFullTitle returns a boolean if a field has been set.

### SetFullTitleNil

`func (o *SubtitleData) SetFullTitleNil(b bool)`

 SetFullTitleNil sets the value for FullTitle to be an explicit nil

### UnsetFullTitle
`func (o *SubtitleData) UnsetFullTitle()`

UnsetFullTitle ensures that no value is present for FullTitle, not even an explicit nil
### GetType

`func (o *SubtitleData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubtitleData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubtitleData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SubtitleData) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SubtitleData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SubtitleData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetYear

`func (o *SubtitleData) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *SubtitleData) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *SubtitleData) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *SubtitleData) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *SubtitleData) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *SubtitleData) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetSubtitles

`func (o *SubtitleData) GetSubtitles() []SubtitleDataDetail`

GetSubtitles returns the Subtitles field if non-nil, zero value otherwise.

### GetSubtitlesOk

`func (o *SubtitleData) GetSubtitlesOk() (*[]SubtitleDataDetail, bool)`

GetSubtitlesOk returns a tuple with the Subtitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitles

`func (o *SubtitleData) SetSubtitles(v []SubtitleDataDetail)`

SetSubtitles sets Subtitles field to given value.

### HasSubtitles

`func (o *SubtitleData) HasSubtitles() bool`

HasSubtitles returns a boolean if a field has been set.

### SetSubtitlesNil

`func (o *SubtitleData) SetSubtitlesNil(b bool)`

 SetSubtitlesNil sets the value for Subtitles to be an explicit nil

### UnsetSubtitles
`func (o *SubtitleData) UnsetSubtitles()`

UnsetSubtitles ensures that no value is present for Subtitles, not even an explicit nil
### GetErrorMessage

`func (o *SubtitleData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SubtitleData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *SubtitleData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *SubtitleData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *SubtitleData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *SubtitleData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


