# MetacriticReviewData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImDbId** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**FullTitle** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Year** | Pointer to **NullableString** |  | [optional] 
**Items** | Pointer to [**[]MetacriticReviewDetail**](MetacriticReviewDetail.md) |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMetacriticReviewData

`func NewMetacriticReviewData() *MetacriticReviewData`

NewMetacriticReviewData instantiates a new MetacriticReviewData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetacriticReviewDataWithDefaults

`func NewMetacriticReviewDataWithDefaults() *MetacriticReviewData`

NewMetacriticReviewDataWithDefaults instantiates a new MetacriticReviewData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImDbId

`func (o *MetacriticReviewData) GetImDbId() string`

GetImDbId returns the ImDbId field if non-nil, zero value otherwise.

### GetImDbIdOk

`func (o *MetacriticReviewData) GetImDbIdOk() (*string, bool)`

GetImDbIdOk returns a tuple with the ImDbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImDbId

`func (o *MetacriticReviewData) SetImDbId(v string)`

SetImDbId sets ImDbId field to given value.

### HasImDbId

`func (o *MetacriticReviewData) HasImDbId() bool`

HasImDbId returns a boolean if a field has been set.

### SetImDbIdNil

`func (o *MetacriticReviewData) SetImDbIdNil(b bool)`

 SetImDbIdNil sets the value for ImDbId to be an explicit nil

### UnsetImDbId
`func (o *MetacriticReviewData) UnsetImDbId()`

UnsetImDbId ensures that no value is present for ImDbId, not even an explicit nil
### GetTitle

`func (o *MetacriticReviewData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MetacriticReviewData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MetacriticReviewData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MetacriticReviewData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MetacriticReviewData) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MetacriticReviewData) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetFullTitle

`func (o *MetacriticReviewData) GetFullTitle() string`

GetFullTitle returns the FullTitle field if non-nil, zero value otherwise.

### GetFullTitleOk

`func (o *MetacriticReviewData) GetFullTitleOk() (*string, bool)`

GetFullTitleOk returns a tuple with the FullTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTitle

`func (o *MetacriticReviewData) SetFullTitle(v string)`

SetFullTitle sets FullTitle field to given value.

### HasFullTitle

`func (o *MetacriticReviewData) HasFullTitle() bool`

HasFullTitle returns a boolean if a field has been set.

### SetFullTitleNil

`func (o *MetacriticReviewData) SetFullTitleNil(b bool)`

 SetFullTitleNil sets the value for FullTitle to be an explicit nil

### UnsetFullTitle
`func (o *MetacriticReviewData) UnsetFullTitle()`

UnsetFullTitle ensures that no value is present for FullTitle, not even an explicit nil
### GetType

`func (o *MetacriticReviewData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetacriticReviewData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetacriticReviewData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MetacriticReviewData) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MetacriticReviewData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MetacriticReviewData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetYear

`func (o *MetacriticReviewData) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *MetacriticReviewData) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *MetacriticReviewData) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *MetacriticReviewData) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *MetacriticReviewData) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *MetacriticReviewData) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetItems

`func (o *MetacriticReviewData) GetItems() []MetacriticReviewDetail`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MetacriticReviewData) GetItemsOk() (*[]MetacriticReviewDetail, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MetacriticReviewData) SetItems(v []MetacriticReviewDetail)`

SetItems sets Items field to given value.

### HasItems

`func (o *MetacriticReviewData) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *MetacriticReviewData) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *MetacriticReviewData) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetErrorMessage

`func (o *MetacriticReviewData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *MetacriticReviewData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *MetacriticReviewData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *MetacriticReviewData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *MetacriticReviewData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *MetacriticReviewData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


