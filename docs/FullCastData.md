# FullCastData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImDbId** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**FullTitle** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Year** | Pointer to **NullableString** |  | [optional] 
**Directors** | Pointer to [**CastShort**](CastShort.md) |  | [optional] 
**Writers** | Pointer to [**CastShort**](CastShort.md) |  | [optional] 
**Actors** | Pointer to [**[]ActorShort**](ActorShort.md) |  | [optional] 
**Others** | Pointer to [**[]CastShort**](CastShort.md) |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFullCastData

`func NewFullCastData() *FullCastData`

NewFullCastData instantiates a new FullCastData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullCastDataWithDefaults

`func NewFullCastDataWithDefaults() *FullCastData`

NewFullCastDataWithDefaults instantiates a new FullCastData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImDbId

`func (o *FullCastData) GetImDbId() string`

GetImDbId returns the ImDbId field if non-nil, zero value otherwise.

### GetImDbIdOk

`func (o *FullCastData) GetImDbIdOk() (*string, bool)`

GetImDbIdOk returns a tuple with the ImDbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImDbId

`func (o *FullCastData) SetImDbId(v string)`

SetImDbId sets ImDbId field to given value.

### HasImDbId

`func (o *FullCastData) HasImDbId() bool`

HasImDbId returns a boolean if a field has been set.

### SetImDbIdNil

`func (o *FullCastData) SetImDbIdNil(b bool)`

 SetImDbIdNil sets the value for ImDbId to be an explicit nil

### UnsetImDbId
`func (o *FullCastData) UnsetImDbId()`

UnsetImDbId ensures that no value is present for ImDbId, not even an explicit nil
### GetTitle

`func (o *FullCastData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FullCastData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FullCastData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FullCastData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *FullCastData) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *FullCastData) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetFullTitle

`func (o *FullCastData) GetFullTitle() string`

GetFullTitle returns the FullTitle field if non-nil, zero value otherwise.

### GetFullTitleOk

`func (o *FullCastData) GetFullTitleOk() (*string, bool)`

GetFullTitleOk returns a tuple with the FullTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTitle

`func (o *FullCastData) SetFullTitle(v string)`

SetFullTitle sets FullTitle field to given value.

### HasFullTitle

`func (o *FullCastData) HasFullTitle() bool`

HasFullTitle returns a boolean if a field has been set.

### SetFullTitleNil

`func (o *FullCastData) SetFullTitleNil(b bool)`

 SetFullTitleNil sets the value for FullTitle to be an explicit nil

### UnsetFullTitle
`func (o *FullCastData) UnsetFullTitle()`

UnsetFullTitle ensures that no value is present for FullTitle, not even an explicit nil
### GetType

`func (o *FullCastData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FullCastData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FullCastData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FullCastData) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *FullCastData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FullCastData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetYear

`func (o *FullCastData) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *FullCastData) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *FullCastData) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *FullCastData) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *FullCastData) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *FullCastData) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetDirectors

`func (o *FullCastData) GetDirectors() CastShort`

GetDirectors returns the Directors field if non-nil, zero value otherwise.

### GetDirectorsOk

`func (o *FullCastData) GetDirectorsOk() (*CastShort, bool)`

GetDirectorsOk returns a tuple with the Directors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectors

`func (o *FullCastData) SetDirectors(v CastShort)`

SetDirectors sets Directors field to given value.

### HasDirectors

`func (o *FullCastData) HasDirectors() bool`

HasDirectors returns a boolean if a field has been set.

### GetWriters

`func (o *FullCastData) GetWriters() CastShort`

GetWriters returns the Writers field if non-nil, zero value otherwise.

### GetWritersOk

`func (o *FullCastData) GetWritersOk() (*CastShort, bool)`

GetWritersOk returns a tuple with the Writers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriters

`func (o *FullCastData) SetWriters(v CastShort)`

SetWriters sets Writers field to given value.

### HasWriters

`func (o *FullCastData) HasWriters() bool`

HasWriters returns a boolean if a field has been set.

### GetActors

`func (o *FullCastData) GetActors() []ActorShort`

GetActors returns the Actors field if non-nil, zero value otherwise.

### GetActorsOk

`func (o *FullCastData) GetActorsOk() (*[]ActorShort, bool)`

GetActorsOk returns a tuple with the Actors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActors

`func (o *FullCastData) SetActors(v []ActorShort)`

SetActors sets Actors field to given value.

### HasActors

`func (o *FullCastData) HasActors() bool`

HasActors returns a boolean if a field has been set.

### SetActorsNil

`func (o *FullCastData) SetActorsNil(b bool)`

 SetActorsNil sets the value for Actors to be an explicit nil

### UnsetActors
`func (o *FullCastData) UnsetActors()`

UnsetActors ensures that no value is present for Actors, not even an explicit nil
### GetOthers

`func (o *FullCastData) GetOthers() []CastShort`

GetOthers returns the Others field if non-nil, zero value otherwise.

### GetOthersOk

`func (o *FullCastData) GetOthersOk() (*[]CastShort, bool)`

GetOthersOk returns a tuple with the Others field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOthers

`func (o *FullCastData) SetOthers(v []CastShort)`

SetOthers sets Others field to given value.

### HasOthers

`func (o *FullCastData) HasOthers() bool`

HasOthers returns a boolean if a field has been set.

### SetOthersNil

`func (o *FullCastData) SetOthersNil(b bool)`

 SetOthersNil sets the value for Others to be an explicit nil

### UnsetOthers
`func (o *FullCastData) UnsetOthers()`

UnsetOthers ensures that no value is present for Others, not even an explicit nil
### GetErrorMessage

`func (o *FullCastData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *FullCastData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *FullCastData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *FullCastData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *FullCastData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *FullCastData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


