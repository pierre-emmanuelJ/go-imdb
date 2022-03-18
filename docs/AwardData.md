# AwardData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImDbId** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**FullTitle** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Year** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Items** | Pointer to [**[]AwardEvent**](AwardEvent.md) |  | [optional] 
**AwardsHtml** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAwardData

`func NewAwardData() *AwardData`

NewAwardData instantiates a new AwardData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwardDataWithDefaults

`func NewAwardDataWithDefaults() *AwardData`

NewAwardDataWithDefaults instantiates a new AwardData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImDbId

`func (o *AwardData) GetImDbId() string`

GetImDbId returns the ImDbId field if non-nil, zero value otherwise.

### GetImDbIdOk

`func (o *AwardData) GetImDbIdOk() (*string, bool)`

GetImDbIdOk returns a tuple with the ImDbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImDbId

`func (o *AwardData) SetImDbId(v string)`

SetImDbId sets ImDbId field to given value.

### HasImDbId

`func (o *AwardData) HasImDbId() bool`

HasImDbId returns a boolean if a field has been set.

### SetImDbIdNil

`func (o *AwardData) SetImDbIdNil(b bool)`

 SetImDbIdNil sets the value for ImDbId to be an explicit nil

### UnsetImDbId
`func (o *AwardData) UnsetImDbId()`

UnsetImDbId ensures that no value is present for ImDbId, not even an explicit nil
### GetTitle

`func (o *AwardData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AwardData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AwardData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AwardData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AwardData) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AwardData) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetFullTitle

`func (o *AwardData) GetFullTitle() string`

GetFullTitle returns the FullTitle field if non-nil, zero value otherwise.

### GetFullTitleOk

`func (o *AwardData) GetFullTitleOk() (*string, bool)`

GetFullTitleOk returns a tuple with the FullTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTitle

`func (o *AwardData) SetFullTitle(v string)`

SetFullTitle sets FullTitle field to given value.

### HasFullTitle

`func (o *AwardData) HasFullTitle() bool`

HasFullTitle returns a boolean if a field has been set.

### SetFullTitleNil

`func (o *AwardData) SetFullTitleNil(b bool)`

 SetFullTitleNil sets the value for FullTitle to be an explicit nil

### UnsetFullTitle
`func (o *AwardData) UnsetFullTitle()`

UnsetFullTitle ensures that no value is present for FullTitle, not even an explicit nil
### GetType

`func (o *AwardData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwardData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwardData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AwardData) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AwardData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AwardData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetYear

`func (o *AwardData) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *AwardData) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *AwardData) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *AwardData) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *AwardData) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *AwardData) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetDescription

`func (o *AwardData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AwardData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AwardData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AwardData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AwardData) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AwardData) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetItems

`func (o *AwardData) GetItems() []AwardEvent`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AwardData) GetItemsOk() (*[]AwardEvent, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AwardData) SetItems(v []AwardEvent)`

SetItems sets Items field to given value.

### HasItems

`func (o *AwardData) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *AwardData) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *AwardData) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetAwardsHtml

`func (o *AwardData) GetAwardsHtml() string`

GetAwardsHtml returns the AwardsHtml field if non-nil, zero value otherwise.

### GetAwardsHtmlOk

`func (o *AwardData) GetAwardsHtmlOk() (*string, bool)`

GetAwardsHtmlOk returns a tuple with the AwardsHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwardsHtml

`func (o *AwardData) SetAwardsHtml(v string)`

SetAwardsHtml sets AwardsHtml field to given value.

### HasAwardsHtml

`func (o *AwardData) HasAwardsHtml() bool`

HasAwardsHtml returns a boolean if a field has been set.

### SetAwardsHtmlNil

`func (o *AwardData) SetAwardsHtmlNil(b bool)`

 SetAwardsHtmlNil sets the value for AwardsHtml to be an explicit nil

### UnsetAwardsHtml
`func (o *AwardData) UnsetAwardsHtml()`

UnsetAwardsHtml ensures that no value is present for AwardsHtml, not even an explicit nil
### GetErrorMessage

`func (o *AwardData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AwardData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AwardData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AwardData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *AwardData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *AwardData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


