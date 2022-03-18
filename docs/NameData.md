# NameData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Role** | Pointer to **NullableString** |  | [optional] 
**Image** | Pointer to **NullableString** |  | [optional] 
**Summary** | Pointer to **NullableString** |  | [optional] 
**BirthDate** | Pointer to **NullableString** |  | [optional] 
**DeathDate** | Pointer to **NullableString** |  | [optional] 
**Awards** | Pointer to **NullableString** |  | [optional] 
**Height** | Pointer to **NullableString** |  | [optional] 
**KnownFor** | Pointer to [**[]KnownFor**](KnownFor.md) |  | [optional] 
**CastMovies** | Pointer to [**[]CastMovie**](CastMovie.md) |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNameData

`func NewNameData() *NameData`

NewNameData instantiates a new NameData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNameDataWithDefaults

`func NewNameDataWithDefaults() *NameData`

NewNameDataWithDefaults instantiates a new NameData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NameData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NameData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NameData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NameData) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *NameData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NameData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *NameData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NameData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NameData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NameData) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *NameData) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *NameData) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRole

`func (o *NameData) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *NameData) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *NameData) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *NameData) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *NameData) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *NameData) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetImage

`func (o *NameData) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *NameData) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *NameData) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *NameData) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *NameData) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *NameData) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetSummary

`func (o *NameData) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *NameData) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *NameData) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *NameData) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *NameData) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *NameData) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetBirthDate

`func (o *NameData) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *NameData) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *NameData) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *NameData) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### SetBirthDateNil

`func (o *NameData) SetBirthDateNil(b bool)`

 SetBirthDateNil sets the value for BirthDate to be an explicit nil

### UnsetBirthDate
`func (o *NameData) UnsetBirthDate()`

UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
### GetDeathDate

`func (o *NameData) GetDeathDate() string`

GetDeathDate returns the DeathDate field if non-nil, zero value otherwise.

### GetDeathDateOk

`func (o *NameData) GetDeathDateOk() (*string, bool)`

GetDeathDateOk returns a tuple with the DeathDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeathDate

`func (o *NameData) SetDeathDate(v string)`

SetDeathDate sets DeathDate field to given value.

### HasDeathDate

`func (o *NameData) HasDeathDate() bool`

HasDeathDate returns a boolean if a field has been set.

### SetDeathDateNil

`func (o *NameData) SetDeathDateNil(b bool)`

 SetDeathDateNil sets the value for DeathDate to be an explicit nil

### UnsetDeathDate
`func (o *NameData) UnsetDeathDate()`

UnsetDeathDate ensures that no value is present for DeathDate, not even an explicit nil
### GetAwards

`func (o *NameData) GetAwards() string`

GetAwards returns the Awards field if non-nil, zero value otherwise.

### GetAwardsOk

`func (o *NameData) GetAwardsOk() (*string, bool)`

GetAwardsOk returns a tuple with the Awards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwards

`func (o *NameData) SetAwards(v string)`

SetAwards sets Awards field to given value.

### HasAwards

`func (o *NameData) HasAwards() bool`

HasAwards returns a boolean if a field has been set.

### SetAwardsNil

`func (o *NameData) SetAwardsNil(b bool)`

 SetAwardsNil sets the value for Awards to be an explicit nil

### UnsetAwards
`func (o *NameData) UnsetAwards()`

UnsetAwards ensures that no value is present for Awards, not even an explicit nil
### GetHeight

`func (o *NameData) GetHeight() string`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *NameData) GetHeightOk() (*string, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *NameData) SetHeight(v string)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *NameData) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *NameData) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *NameData) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetKnownFor

`func (o *NameData) GetKnownFor() []KnownFor`

GetKnownFor returns the KnownFor field if non-nil, zero value otherwise.

### GetKnownForOk

`func (o *NameData) GetKnownForOk() (*[]KnownFor, bool)`

GetKnownForOk returns a tuple with the KnownFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownFor

`func (o *NameData) SetKnownFor(v []KnownFor)`

SetKnownFor sets KnownFor field to given value.

### HasKnownFor

`func (o *NameData) HasKnownFor() bool`

HasKnownFor returns a boolean if a field has been set.

### SetKnownForNil

`func (o *NameData) SetKnownForNil(b bool)`

 SetKnownForNil sets the value for KnownFor to be an explicit nil

### UnsetKnownFor
`func (o *NameData) UnsetKnownFor()`

UnsetKnownFor ensures that no value is present for KnownFor, not even an explicit nil
### GetCastMovies

`func (o *NameData) GetCastMovies() []CastMovie`

GetCastMovies returns the CastMovies field if non-nil, zero value otherwise.

### GetCastMoviesOk

`func (o *NameData) GetCastMoviesOk() (*[]CastMovie, bool)`

GetCastMoviesOk returns a tuple with the CastMovies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastMovies

`func (o *NameData) SetCastMovies(v []CastMovie)`

SetCastMovies sets CastMovies field to given value.

### HasCastMovies

`func (o *NameData) HasCastMovies() bool`

HasCastMovies returns a boolean if a field has been set.

### SetCastMoviesNil

`func (o *NameData) SetCastMoviesNil(b bool)`

 SetCastMoviesNil sets the value for CastMovies to be an explicit nil

### UnsetCastMovies
`func (o *NameData) UnsetCastMovies()`

UnsetCastMovies ensures that no value is present for CastMovies, not even an explicit nil
### GetErrorMessage

`func (o *NameData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *NameData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *NameData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *NameData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *NameData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *NameData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


