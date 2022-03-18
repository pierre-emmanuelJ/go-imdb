# FAQData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImDbId** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**FullTitle** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Year** | Pointer to **NullableString** |  | [optional] 
**Items** | Pointer to [**[]FAQDetail**](FAQDetail.md) |  | [optional] 
**SpoilerItems** | Pointer to [**[]FAQDetail**](FAQDetail.md) |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFAQData

`func NewFAQData() *FAQData`

NewFAQData instantiates a new FAQData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFAQDataWithDefaults

`func NewFAQDataWithDefaults() *FAQData`

NewFAQDataWithDefaults instantiates a new FAQData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImDbId

`func (o *FAQData) GetImDbId() string`

GetImDbId returns the ImDbId field if non-nil, zero value otherwise.

### GetImDbIdOk

`func (o *FAQData) GetImDbIdOk() (*string, bool)`

GetImDbIdOk returns a tuple with the ImDbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImDbId

`func (o *FAQData) SetImDbId(v string)`

SetImDbId sets ImDbId field to given value.

### HasImDbId

`func (o *FAQData) HasImDbId() bool`

HasImDbId returns a boolean if a field has been set.

### SetImDbIdNil

`func (o *FAQData) SetImDbIdNil(b bool)`

 SetImDbIdNil sets the value for ImDbId to be an explicit nil

### UnsetImDbId
`func (o *FAQData) UnsetImDbId()`

UnsetImDbId ensures that no value is present for ImDbId, not even an explicit nil
### GetTitle

`func (o *FAQData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FAQData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FAQData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FAQData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *FAQData) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *FAQData) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetFullTitle

`func (o *FAQData) GetFullTitle() string`

GetFullTitle returns the FullTitle field if non-nil, zero value otherwise.

### GetFullTitleOk

`func (o *FAQData) GetFullTitleOk() (*string, bool)`

GetFullTitleOk returns a tuple with the FullTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTitle

`func (o *FAQData) SetFullTitle(v string)`

SetFullTitle sets FullTitle field to given value.

### HasFullTitle

`func (o *FAQData) HasFullTitle() bool`

HasFullTitle returns a boolean if a field has been set.

### SetFullTitleNil

`func (o *FAQData) SetFullTitleNil(b bool)`

 SetFullTitleNil sets the value for FullTitle to be an explicit nil

### UnsetFullTitle
`func (o *FAQData) UnsetFullTitle()`

UnsetFullTitle ensures that no value is present for FullTitle, not even an explicit nil
### GetType

`func (o *FAQData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FAQData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FAQData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FAQData) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *FAQData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FAQData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetYear

`func (o *FAQData) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *FAQData) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *FAQData) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *FAQData) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *FAQData) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *FAQData) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetItems

`func (o *FAQData) GetItems() []FAQDetail`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FAQData) GetItemsOk() (*[]FAQDetail, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FAQData) SetItems(v []FAQDetail)`

SetItems sets Items field to given value.

### HasItems

`func (o *FAQData) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *FAQData) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *FAQData) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetSpoilerItems

`func (o *FAQData) GetSpoilerItems() []FAQDetail`

GetSpoilerItems returns the SpoilerItems field if non-nil, zero value otherwise.

### GetSpoilerItemsOk

`func (o *FAQData) GetSpoilerItemsOk() (*[]FAQDetail, bool)`

GetSpoilerItemsOk returns a tuple with the SpoilerItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoilerItems

`func (o *FAQData) SetSpoilerItems(v []FAQDetail)`

SetSpoilerItems sets SpoilerItems field to given value.

### HasSpoilerItems

`func (o *FAQData) HasSpoilerItems() bool`

HasSpoilerItems returns a boolean if a field has been set.

### SetSpoilerItemsNil

`func (o *FAQData) SetSpoilerItemsNil(b bool)`

 SetSpoilerItemsNil sets the value for SpoilerItems to be an explicit nil

### UnsetSpoilerItems
`func (o *FAQData) UnsetSpoilerItems()`

UnsetSpoilerItems ensures that no value is present for SpoilerItems, not even an explicit nil
### GetErrorMessage

`func (o *FAQData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *FAQData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *FAQData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *FAQData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *FAQData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *FAQData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


