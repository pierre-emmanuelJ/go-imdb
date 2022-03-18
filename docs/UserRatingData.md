# UserRatingData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImDbId** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**FullTitle** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Year** | Pointer to **NullableString** |  | [optional] 
**TotalRating** | Pointer to **NullableString** |  | [optional] 
**TotalRatingVotes** | Pointer to **NullableString** |  | [optional] 
**Ratings** | Pointer to [**[]UserRatingDataDetail**](UserRatingDataDetail.md) |  | [optional] 
**DemographicAll** | Pointer to [**UserRatingDataDemographic**](UserRatingDataDemographic.md) |  | [optional] 
**DemographicMales** | Pointer to [**UserRatingDataDemographic**](UserRatingDataDemographic.md) |  | [optional] 
**DemographicFemales** | Pointer to [**UserRatingDataDemographic**](UserRatingDataDemographic.md) |  | [optional] 
**Top1000Voters** | Pointer to [**UserRatingDataDemographicDetail**](UserRatingDataDemographicDetail.md) |  | [optional] 
**UsUsers** | Pointer to [**UserRatingDataDemographicDetail**](UserRatingDataDemographicDetail.md) |  | [optional] 
**NonUSUsers** | Pointer to [**UserRatingDataDemographicDetail**](UserRatingDataDemographicDetail.md) |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserRatingData

`func NewUserRatingData() *UserRatingData`

NewUserRatingData instantiates a new UserRatingData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRatingDataWithDefaults

`func NewUserRatingDataWithDefaults() *UserRatingData`

NewUserRatingDataWithDefaults instantiates a new UserRatingData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImDbId

`func (o *UserRatingData) GetImDbId() string`

GetImDbId returns the ImDbId field if non-nil, zero value otherwise.

### GetImDbIdOk

`func (o *UserRatingData) GetImDbIdOk() (*string, bool)`

GetImDbIdOk returns a tuple with the ImDbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImDbId

`func (o *UserRatingData) SetImDbId(v string)`

SetImDbId sets ImDbId field to given value.

### HasImDbId

`func (o *UserRatingData) HasImDbId() bool`

HasImDbId returns a boolean if a field has been set.

### SetImDbIdNil

`func (o *UserRatingData) SetImDbIdNil(b bool)`

 SetImDbIdNil sets the value for ImDbId to be an explicit nil

### UnsetImDbId
`func (o *UserRatingData) UnsetImDbId()`

UnsetImDbId ensures that no value is present for ImDbId, not even an explicit nil
### GetTitle

`func (o *UserRatingData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UserRatingData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UserRatingData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UserRatingData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *UserRatingData) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UserRatingData) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetFullTitle

`func (o *UserRatingData) GetFullTitle() string`

GetFullTitle returns the FullTitle field if non-nil, zero value otherwise.

### GetFullTitleOk

`func (o *UserRatingData) GetFullTitleOk() (*string, bool)`

GetFullTitleOk returns a tuple with the FullTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTitle

`func (o *UserRatingData) SetFullTitle(v string)`

SetFullTitle sets FullTitle field to given value.

### HasFullTitle

`func (o *UserRatingData) HasFullTitle() bool`

HasFullTitle returns a boolean if a field has been set.

### SetFullTitleNil

`func (o *UserRatingData) SetFullTitleNil(b bool)`

 SetFullTitleNil sets the value for FullTitle to be an explicit nil

### UnsetFullTitle
`func (o *UserRatingData) UnsetFullTitle()`

UnsetFullTitle ensures that no value is present for FullTitle, not even an explicit nil
### GetType

`func (o *UserRatingData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserRatingData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserRatingData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserRatingData) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *UserRatingData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *UserRatingData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetYear

`func (o *UserRatingData) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *UserRatingData) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *UserRatingData) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *UserRatingData) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *UserRatingData) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *UserRatingData) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetTotalRating

`func (o *UserRatingData) GetTotalRating() string`

GetTotalRating returns the TotalRating field if non-nil, zero value otherwise.

### GetTotalRatingOk

`func (o *UserRatingData) GetTotalRatingOk() (*string, bool)`

GetTotalRatingOk returns a tuple with the TotalRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRating

`func (o *UserRatingData) SetTotalRating(v string)`

SetTotalRating sets TotalRating field to given value.

### HasTotalRating

`func (o *UserRatingData) HasTotalRating() bool`

HasTotalRating returns a boolean if a field has been set.

### SetTotalRatingNil

`func (o *UserRatingData) SetTotalRatingNil(b bool)`

 SetTotalRatingNil sets the value for TotalRating to be an explicit nil

### UnsetTotalRating
`func (o *UserRatingData) UnsetTotalRating()`

UnsetTotalRating ensures that no value is present for TotalRating, not even an explicit nil
### GetTotalRatingVotes

`func (o *UserRatingData) GetTotalRatingVotes() string`

GetTotalRatingVotes returns the TotalRatingVotes field if non-nil, zero value otherwise.

### GetTotalRatingVotesOk

`func (o *UserRatingData) GetTotalRatingVotesOk() (*string, bool)`

GetTotalRatingVotesOk returns a tuple with the TotalRatingVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRatingVotes

`func (o *UserRatingData) SetTotalRatingVotes(v string)`

SetTotalRatingVotes sets TotalRatingVotes field to given value.

### HasTotalRatingVotes

`func (o *UserRatingData) HasTotalRatingVotes() bool`

HasTotalRatingVotes returns a boolean if a field has been set.

### SetTotalRatingVotesNil

`func (o *UserRatingData) SetTotalRatingVotesNil(b bool)`

 SetTotalRatingVotesNil sets the value for TotalRatingVotes to be an explicit nil

### UnsetTotalRatingVotes
`func (o *UserRatingData) UnsetTotalRatingVotes()`

UnsetTotalRatingVotes ensures that no value is present for TotalRatingVotes, not even an explicit nil
### GetRatings

`func (o *UserRatingData) GetRatings() []UserRatingDataDetail`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *UserRatingData) GetRatingsOk() (*[]UserRatingDataDetail, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *UserRatingData) SetRatings(v []UserRatingDataDetail)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *UserRatingData) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### SetRatingsNil

`func (o *UserRatingData) SetRatingsNil(b bool)`

 SetRatingsNil sets the value for Ratings to be an explicit nil

### UnsetRatings
`func (o *UserRatingData) UnsetRatings()`

UnsetRatings ensures that no value is present for Ratings, not even an explicit nil
### GetDemographicAll

`func (o *UserRatingData) GetDemographicAll() UserRatingDataDemographic`

GetDemographicAll returns the DemographicAll field if non-nil, zero value otherwise.

### GetDemographicAllOk

`func (o *UserRatingData) GetDemographicAllOk() (*UserRatingDataDemographic, bool)`

GetDemographicAllOk returns a tuple with the DemographicAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemographicAll

`func (o *UserRatingData) SetDemographicAll(v UserRatingDataDemographic)`

SetDemographicAll sets DemographicAll field to given value.

### HasDemographicAll

`func (o *UserRatingData) HasDemographicAll() bool`

HasDemographicAll returns a boolean if a field has been set.

### GetDemographicMales

`func (o *UserRatingData) GetDemographicMales() UserRatingDataDemographic`

GetDemographicMales returns the DemographicMales field if non-nil, zero value otherwise.

### GetDemographicMalesOk

`func (o *UserRatingData) GetDemographicMalesOk() (*UserRatingDataDemographic, bool)`

GetDemographicMalesOk returns a tuple with the DemographicMales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemographicMales

`func (o *UserRatingData) SetDemographicMales(v UserRatingDataDemographic)`

SetDemographicMales sets DemographicMales field to given value.

### HasDemographicMales

`func (o *UserRatingData) HasDemographicMales() bool`

HasDemographicMales returns a boolean if a field has been set.

### GetDemographicFemales

`func (o *UserRatingData) GetDemographicFemales() UserRatingDataDemographic`

GetDemographicFemales returns the DemographicFemales field if non-nil, zero value otherwise.

### GetDemographicFemalesOk

`func (o *UserRatingData) GetDemographicFemalesOk() (*UserRatingDataDemographic, bool)`

GetDemographicFemalesOk returns a tuple with the DemographicFemales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemographicFemales

`func (o *UserRatingData) SetDemographicFemales(v UserRatingDataDemographic)`

SetDemographicFemales sets DemographicFemales field to given value.

### HasDemographicFemales

`func (o *UserRatingData) HasDemographicFemales() bool`

HasDemographicFemales returns a boolean if a field has been set.

### GetTop1000Voters

`func (o *UserRatingData) GetTop1000Voters() UserRatingDataDemographicDetail`

GetTop1000Voters returns the Top1000Voters field if non-nil, zero value otherwise.

### GetTop1000VotersOk

`func (o *UserRatingData) GetTop1000VotersOk() (*UserRatingDataDemographicDetail, bool)`

GetTop1000VotersOk returns a tuple with the Top1000Voters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop1000Voters

`func (o *UserRatingData) SetTop1000Voters(v UserRatingDataDemographicDetail)`

SetTop1000Voters sets Top1000Voters field to given value.

### HasTop1000Voters

`func (o *UserRatingData) HasTop1000Voters() bool`

HasTop1000Voters returns a boolean if a field has been set.

### GetUsUsers

`func (o *UserRatingData) GetUsUsers() UserRatingDataDemographicDetail`

GetUsUsers returns the UsUsers field if non-nil, zero value otherwise.

### GetUsUsersOk

`func (o *UserRatingData) GetUsUsersOk() (*UserRatingDataDemographicDetail, bool)`

GetUsUsersOk returns a tuple with the UsUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsUsers

`func (o *UserRatingData) SetUsUsers(v UserRatingDataDemographicDetail)`

SetUsUsers sets UsUsers field to given value.

### HasUsUsers

`func (o *UserRatingData) HasUsUsers() bool`

HasUsUsers returns a boolean if a field has been set.

### GetNonUSUsers

`func (o *UserRatingData) GetNonUSUsers() UserRatingDataDemographicDetail`

GetNonUSUsers returns the NonUSUsers field if non-nil, zero value otherwise.

### GetNonUSUsersOk

`func (o *UserRatingData) GetNonUSUsersOk() (*UserRatingDataDemographicDetail, bool)`

GetNonUSUsersOk returns a tuple with the NonUSUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonUSUsers

`func (o *UserRatingData) SetNonUSUsers(v UserRatingDataDemographicDetail)`

SetNonUSUsers sets NonUSUsers field to given value.

### HasNonUSUsers

`func (o *UserRatingData) HasNonUSUsers() bool`

HasNonUSUsers returns a boolean if a field has been set.

### GetErrorMessage

`func (o *UserRatingData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *UserRatingData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *UserRatingData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *UserRatingData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *UserRatingData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *UserRatingData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


