# SeasonEpisodeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImDbId** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**FullTitle** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Year** | Pointer to **NullableString** |  | [optional] 
**Episodes** | Pointer to [**[]EpisodeShortDetail**](EpisodeShortDetail.md) |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSeasonEpisodeData

`func NewSeasonEpisodeData() *SeasonEpisodeData`

NewSeasonEpisodeData instantiates a new SeasonEpisodeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeasonEpisodeDataWithDefaults

`func NewSeasonEpisodeDataWithDefaults() *SeasonEpisodeData`

NewSeasonEpisodeDataWithDefaults instantiates a new SeasonEpisodeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImDbId

`func (o *SeasonEpisodeData) GetImDbId() string`

GetImDbId returns the ImDbId field if non-nil, zero value otherwise.

### GetImDbIdOk

`func (o *SeasonEpisodeData) GetImDbIdOk() (*string, bool)`

GetImDbIdOk returns a tuple with the ImDbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImDbId

`func (o *SeasonEpisodeData) SetImDbId(v string)`

SetImDbId sets ImDbId field to given value.

### HasImDbId

`func (o *SeasonEpisodeData) HasImDbId() bool`

HasImDbId returns a boolean if a field has been set.

### SetImDbIdNil

`func (o *SeasonEpisodeData) SetImDbIdNil(b bool)`

 SetImDbIdNil sets the value for ImDbId to be an explicit nil

### UnsetImDbId
`func (o *SeasonEpisodeData) UnsetImDbId()`

UnsetImDbId ensures that no value is present for ImDbId, not even an explicit nil
### GetTitle

`func (o *SeasonEpisodeData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SeasonEpisodeData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SeasonEpisodeData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SeasonEpisodeData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SeasonEpisodeData) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SeasonEpisodeData) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetFullTitle

`func (o *SeasonEpisodeData) GetFullTitle() string`

GetFullTitle returns the FullTitle field if non-nil, zero value otherwise.

### GetFullTitleOk

`func (o *SeasonEpisodeData) GetFullTitleOk() (*string, bool)`

GetFullTitleOk returns a tuple with the FullTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTitle

`func (o *SeasonEpisodeData) SetFullTitle(v string)`

SetFullTitle sets FullTitle field to given value.

### HasFullTitle

`func (o *SeasonEpisodeData) HasFullTitle() bool`

HasFullTitle returns a boolean if a field has been set.

### SetFullTitleNil

`func (o *SeasonEpisodeData) SetFullTitleNil(b bool)`

 SetFullTitleNil sets the value for FullTitle to be an explicit nil

### UnsetFullTitle
`func (o *SeasonEpisodeData) UnsetFullTitle()`

UnsetFullTitle ensures that no value is present for FullTitle, not even an explicit nil
### GetType

`func (o *SeasonEpisodeData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SeasonEpisodeData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SeasonEpisodeData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SeasonEpisodeData) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SeasonEpisodeData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SeasonEpisodeData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetYear

`func (o *SeasonEpisodeData) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *SeasonEpisodeData) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *SeasonEpisodeData) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *SeasonEpisodeData) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *SeasonEpisodeData) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *SeasonEpisodeData) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetEpisodes

`func (o *SeasonEpisodeData) GetEpisodes() []EpisodeShortDetail`

GetEpisodes returns the Episodes field if non-nil, zero value otherwise.

### GetEpisodesOk

`func (o *SeasonEpisodeData) GetEpisodesOk() (*[]EpisodeShortDetail, bool)`

GetEpisodesOk returns a tuple with the Episodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodes

`func (o *SeasonEpisodeData) SetEpisodes(v []EpisodeShortDetail)`

SetEpisodes sets Episodes field to given value.

### HasEpisodes

`func (o *SeasonEpisodeData) HasEpisodes() bool`

HasEpisodes returns a boolean if a field has been set.

### SetEpisodesNil

`func (o *SeasonEpisodeData) SetEpisodesNil(b bool)`

 SetEpisodesNil sets the value for Episodes to be an explicit nil

### UnsetEpisodes
`func (o *SeasonEpisodeData) UnsetEpisodes()`

UnsetEpisodes ensures that no value is present for Episodes, not even an explicit nil
### GetErrorMessage

`func (o *SeasonEpisodeData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SeasonEpisodeData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *SeasonEpisodeData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *SeasonEpisodeData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *SeasonEpisodeData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *SeasonEpisodeData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


