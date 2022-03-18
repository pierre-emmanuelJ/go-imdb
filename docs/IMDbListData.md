# IMDbListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**By** | Pointer to **NullableString** |  | [optional] 
**Created** | Pointer to **NullableString** |  | [optional] 
**Updated** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Items** | Pointer to [**[]IMDbListDataDetail**](IMDbListDataDetail.md) |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIMDbListData

`func NewIMDbListData() *IMDbListData`

NewIMDbListData instantiates a new IMDbListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIMDbListDataWithDefaults

`func NewIMDbListDataWithDefaults() *IMDbListData`

NewIMDbListDataWithDefaults instantiates a new IMDbListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *IMDbListData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IMDbListData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IMDbListData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IMDbListData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *IMDbListData) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *IMDbListData) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetBy

`func (o *IMDbListData) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *IMDbListData) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *IMDbListData) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *IMDbListData) HasBy() bool`

HasBy returns a boolean if a field has been set.

### SetByNil

`func (o *IMDbListData) SetByNil(b bool)`

 SetByNil sets the value for By to be an explicit nil

### UnsetBy
`func (o *IMDbListData) UnsetBy()`

UnsetBy ensures that no value is present for By, not even an explicit nil
### GetCreated

`func (o *IMDbListData) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IMDbListData) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IMDbListData) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IMDbListData) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *IMDbListData) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *IMDbListData) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetUpdated

`func (o *IMDbListData) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *IMDbListData) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *IMDbListData) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *IMDbListData) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *IMDbListData) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *IMDbListData) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
### GetDescription

`func (o *IMDbListData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IMDbListData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IMDbListData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IMDbListData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IMDbListData) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IMDbListData) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetItems

`func (o *IMDbListData) GetItems() []IMDbListDataDetail`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IMDbListData) GetItemsOk() (*[]IMDbListDataDetail, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IMDbListData) SetItems(v []IMDbListDataDetail)`

SetItems sets Items field to given value.

### HasItems

`func (o *IMDbListData) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *IMDbListData) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *IMDbListData) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetErrorMessage

`func (o *IMDbListData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *IMDbListData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *IMDbListData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *IMDbListData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *IMDbListData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *IMDbListData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


