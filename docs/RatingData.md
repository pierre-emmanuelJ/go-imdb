# RatingData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImDbId** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**FullTitle** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Year** | Pointer to **NullableString** |  | [optional] 
**ImDb** | Pointer to **NullableString** |  | [optional] 
**Metacritic** | Pointer to **NullableString** |  | [optional] 
**TheMovieDb** | Pointer to **NullableString** |  | [optional] 
**RottenTomatoes** | Pointer to **NullableString** |  | [optional] 
**FilmAffinity** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRatingData

`func NewRatingData() *RatingData`

NewRatingData instantiates a new RatingData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatingDataWithDefaults

`func NewRatingDataWithDefaults() *RatingData`

NewRatingDataWithDefaults instantiates a new RatingData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImDbId

`func (o *RatingData) GetImDbId() string`

GetImDbId returns the ImDbId field if non-nil, zero value otherwise.

### GetImDbIdOk

`func (o *RatingData) GetImDbIdOk() (*string, bool)`

GetImDbIdOk returns a tuple with the ImDbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImDbId

`func (o *RatingData) SetImDbId(v string)`

SetImDbId sets ImDbId field to given value.

### HasImDbId

`func (o *RatingData) HasImDbId() bool`

HasImDbId returns a boolean if a field has been set.

### SetImDbIdNil

`func (o *RatingData) SetImDbIdNil(b bool)`

 SetImDbIdNil sets the value for ImDbId to be an explicit nil

### UnsetImDbId
`func (o *RatingData) UnsetImDbId()`

UnsetImDbId ensures that no value is present for ImDbId, not even an explicit nil
### GetTitle

`func (o *RatingData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RatingData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RatingData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RatingData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *RatingData) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *RatingData) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetFullTitle

`func (o *RatingData) GetFullTitle() string`

GetFullTitle returns the FullTitle field if non-nil, zero value otherwise.

### GetFullTitleOk

`func (o *RatingData) GetFullTitleOk() (*string, bool)`

GetFullTitleOk returns a tuple with the FullTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTitle

`func (o *RatingData) SetFullTitle(v string)`

SetFullTitle sets FullTitle field to given value.

### HasFullTitle

`func (o *RatingData) HasFullTitle() bool`

HasFullTitle returns a boolean if a field has been set.

### SetFullTitleNil

`func (o *RatingData) SetFullTitleNil(b bool)`

 SetFullTitleNil sets the value for FullTitle to be an explicit nil

### UnsetFullTitle
`func (o *RatingData) UnsetFullTitle()`

UnsetFullTitle ensures that no value is present for FullTitle, not even an explicit nil
### GetType

`func (o *RatingData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RatingData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RatingData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RatingData) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *RatingData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *RatingData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetYear

`func (o *RatingData) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *RatingData) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *RatingData) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *RatingData) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *RatingData) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *RatingData) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetImDb

`func (o *RatingData) GetImDb() string`

GetImDb returns the ImDb field if non-nil, zero value otherwise.

### GetImDbOk

`func (o *RatingData) GetImDbOk() (*string, bool)`

GetImDbOk returns a tuple with the ImDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImDb

`func (o *RatingData) SetImDb(v string)`

SetImDb sets ImDb field to given value.

### HasImDb

`func (o *RatingData) HasImDb() bool`

HasImDb returns a boolean if a field has been set.

### SetImDbNil

`func (o *RatingData) SetImDbNil(b bool)`

 SetImDbNil sets the value for ImDb to be an explicit nil

### UnsetImDb
`func (o *RatingData) UnsetImDb()`

UnsetImDb ensures that no value is present for ImDb, not even an explicit nil
### GetMetacritic

`func (o *RatingData) GetMetacritic() string`

GetMetacritic returns the Metacritic field if non-nil, zero value otherwise.

### GetMetacriticOk

`func (o *RatingData) GetMetacriticOk() (*string, bool)`

GetMetacriticOk returns a tuple with the Metacritic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetacritic

`func (o *RatingData) SetMetacritic(v string)`

SetMetacritic sets Metacritic field to given value.

### HasMetacritic

`func (o *RatingData) HasMetacritic() bool`

HasMetacritic returns a boolean if a field has been set.

### SetMetacriticNil

`func (o *RatingData) SetMetacriticNil(b bool)`

 SetMetacriticNil sets the value for Metacritic to be an explicit nil

### UnsetMetacritic
`func (o *RatingData) UnsetMetacritic()`

UnsetMetacritic ensures that no value is present for Metacritic, not even an explicit nil
### GetTheMovieDb

`func (o *RatingData) GetTheMovieDb() string`

GetTheMovieDb returns the TheMovieDb field if non-nil, zero value otherwise.

### GetTheMovieDbOk

`func (o *RatingData) GetTheMovieDbOk() (*string, bool)`

GetTheMovieDbOk returns a tuple with the TheMovieDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheMovieDb

`func (o *RatingData) SetTheMovieDb(v string)`

SetTheMovieDb sets TheMovieDb field to given value.

### HasTheMovieDb

`func (o *RatingData) HasTheMovieDb() bool`

HasTheMovieDb returns a boolean if a field has been set.

### SetTheMovieDbNil

`func (o *RatingData) SetTheMovieDbNil(b bool)`

 SetTheMovieDbNil sets the value for TheMovieDb to be an explicit nil

### UnsetTheMovieDb
`func (o *RatingData) UnsetTheMovieDb()`

UnsetTheMovieDb ensures that no value is present for TheMovieDb, not even an explicit nil
### GetRottenTomatoes

`func (o *RatingData) GetRottenTomatoes() string`

GetRottenTomatoes returns the RottenTomatoes field if non-nil, zero value otherwise.

### GetRottenTomatoesOk

`func (o *RatingData) GetRottenTomatoesOk() (*string, bool)`

GetRottenTomatoesOk returns a tuple with the RottenTomatoes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRottenTomatoes

`func (o *RatingData) SetRottenTomatoes(v string)`

SetRottenTomatoes sets RottenTomatoes field to given value.

### HasRottenTomatoes

`func (o *RatingData) HasRottenTomatoes() bool`

HasRottenTomatoes returns a boolean if a field has been set.

### SetRottenTomatoesNil

`func (o *RatingData) SetRottenTomatoesNil(b bool)`

 SetRottenTomatoesNil sets the value for RottenTomatoes to be an explicit nil

### UnsetRottenTomatoes
`func (o *RatingData) UnsetRottenTomatoes()`

UnsetRottenTomatoes ensures that no value is present for RottenTomatoes, not even an explicit nil
### GetFilmAffinity

`func (o *RatingData) GetFilmAffinity() string`

GetFilmAffinity returns the FilmAffinity field if non-nil, zero value otherwise.

### GetFilmAffinityOk

`func (o *RatingData) GetFilmAffinityOk() (*string, bool)`

GetFilmAffinityOk returns a tuple with the FilmAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilmAffinity

`func (o *RatingData) SetFilmAffinity(v string)`

SetFilmAffinity sets FilmAffinity field to given value.

### HasFilmAffinity

`func (o *RatingData) HasFilmAffinity() bool`

HasFilmAffinity returns a boolean if a field has been set.

### SetFilmAffinityNil

`func (o *RatingData) SetFilmAffinityNil(b bool)`

 SetFilmAffinityNil sets the value for FilmAffinity to be an explicit nil

### UnsetFilmAffinity
`func (o *RatingData) UnsetFilmAffinity()`

UnsetFilmAffinity ensures that no value is present for FilmAffinity, not even an explicit nil
### GetErrorMessage

`func (o *RatingData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *RatingData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *RatingData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *RatingData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *RatingData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *RatingData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


