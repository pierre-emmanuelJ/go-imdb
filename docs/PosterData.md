# PosterData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImDbId** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**FullTitle** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Year** | Pointer to **NullableString** |  | [optional] 
**Posters** | Pointer to [**[]PosterDataItem**](PosterDataItem.md) |  | [optional] 
**Backdrops** | Pointer to [**[]PosterDataItem**](PosterDataItem.md) |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPosterData

`func NewPosterData() *PosterData`

NewPosterData instantiates a new PosterData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPosterDataWithDefaults

`func NewPosterDataWithDefaults() *PosterData`

NewPosterDataWithDefaults instantiates a new PosterData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImDbId

`func (o *PosterData) GetImDbId() string`

GetImDbId returns the ImDbId field if non-nil, zero value otherwise.

### GetImDbIdOk

`func (o *PosterData) GetImDbIdOk() (*string, bool)`

GetImDbIdOk returns a tuple with the ImDbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImDbId

`func (o *PosterData) SetImDbId(v string)`

SetImDbId sets ImDbId field to given value.

### HasImDbId

`func (o *PosterData) HasImDbId() bool`

HasImDbId returns a boolean if a field has been set.

### SetImDbIdNil

`func (o *PosterData) SetImDbIdNil(b bool)`

 SetImDbIdNil sets the value for ImDbId to be an explicit nil

### UnsetImDbId
`func (o *PosterData) UnsetImDbId()`

UnsetImDbId ensures that no value is present for ImDbId, not even an explicit nil
### GetTitle

`func (o *PosterData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PosterData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PosterData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PosterData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *PosterData) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *PosterData) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetFullTitle

`func (o *PosterData) GetFullTitle() string`

GetFullTitle returns the FullTitle field if non-nil, zero value otherwise.

### GetFullTitleOk

`func (o *PosterData) GetFullTitleOk() (*string, bool)`

GetFullTitleOk returns a tuple with the FullTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTitle

`func (o *PosterData) SetFullTitle(v string)`

SetFullTitle sets FullTitle field to given value.

### HasFullTitle

`func (o *PosterData) HasFullTitle() bool`

HasFullTitle returns a boolean if a field has been set.

### SetFullTitleNil

`func (o *PosterData) SetFullTitleNil(b bool)`

 SetFullTitleNil sets the value for FullTitle to be an explicit nil

### UnsetFullTitle
`func (o *PosterData) UnsetFullTitle()`

UnsetFullTitle ensures that no value is present for FullTitle, not even an explicit nil
### GetType

`func (o *PosterData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PosterData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PosterData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PosterData) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PosterData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PosterData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetYear

`func (o *PosterData) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *PosterData) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *PosterData) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *PosterData) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *PosterData) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *PosterData) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetPosters

`func (o *PosterData) GetPosters() []PosterDataItem`

GetPosters returns the Posters field if non-nil, zero value otherwise.

### GetPostersOk

`func (o *PosterData) GetPostersOk() (*[]PosterDataItem, bool)`

GetPostersOk returns a tuple with the Posters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosters

`func (o *PosterData) SetPosters(v []PosterDataItem)`

SetPosters sets Posters field to given value.

### HasPosters

`func (o *PosterData) HasPosters() bool`

HasPosters returns a boolean if a field has been set.

### SetPostersNil

`func (o *PosterData) SetPostersNil(b bool)`

 SetPostersNil sets the value for Posters to be an explicit nil

### UnsetPosters
`func (o *PosterData) UnsetPosters()`

UnsetPosters ensures that no value is present for Posters, not even an explicit nil
### GetBackdrops

`func (o *PosterData) GetBackdrops() []PosterDataItem`

GetBackdrops returns the Backdrops field if non-nil, zero value otherwise.

### GetBackdropsOk

`func (o *PosterData) GetBackdropsOk() (*[]PosterDataItem, bool)`

GetBackdropsOk returns a tuple with the Backdrops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackdrops

`func (o *PosterData) SetBackdrops(v []PosterDataItem)`

SetBackdrops sets Backdrops field to given value.

### HasBackdrops

`func (o *PosterData) HasBackdrops() bool`

HasBackdrops returns a boolean if a field has been set.

### SetBackdropsNil

`func (o *PosterData) SetBackdropsNil(b bool)`

 SetBackdropsNil sets the value for Backdrops to be an explicit nil

### UnsetBackdrops
`func (o *PosterData) UnsetBackdrops()`

UnsetBackdrops ensures that no value is present for Backdrops, not even an explicit nil
### GetErrorMessage

`func (o *PosterData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PosterData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PosterData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PosterData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *PosterData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *PosterData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


