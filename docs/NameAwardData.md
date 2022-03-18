# NameAwardData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImDbId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Items** | Pointer to [**[]NameAwardEvent**](NameAwardEvent.md) |  | [optional] 
**NameAwardsHtml** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNameAwardData

`func NewNameAwardData() *NameAwardData`

NewNameAwardData instantiates a new NameAwardData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNameAwardDataWithDefaults

`func NewNameAwardDataWithDefaults() *NameAwardData`

NewNameAwardDataWithDefaults instantiates a new NameAwardData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImDbId

`func (o *NameAwardData) GetImDbId() string`

GetImDbId returns the ImDbId field if non-nil, zero value otherwise.

### GetImDbIdOk

`func (o *NameAwardData) GetImDbIdOk() (*string, bool)`

GetImDbIdOk returns a tuple with the ImDbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImDbId

`func (o *NameAwardData) SetImDbId(v string)`

SetImDbId sets ImDbId field to given value.

### HasImDbId

`func (o *NameAwardData) HasImDbId() bool`

HasImDbId returns a boolean if a field has been set.

### SetImDbIdNil

`func (o *NameAwardData) SetImDbIdNil(b bool)`

 SetImDbIdNil sets the value for ImDbId to be an explicit nil

### UnsetImDbId
`func (o *NameAwardData) UnsetImDbId()`

UnsetImDbId ensures that no value is present for ImDbId, not even an explicit nil
### GetName

`func (o *NameAwardData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NameAwardData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NameAwardData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NameAwardData) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *NameAwardData) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *NameAwardData) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *NameAwardData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NameAwardData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NameAwardData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NameAwardData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NameAwardData) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NameAwardData) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetItems

`func (o *NameAwardData) GetItems() []NameAwardEvent`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NameAwardData) GetItemsOk() (*[]NameAwardEvent, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NameAwardData) SetItems(v []NameAwardEvent)`

SetItems sets Items field to given value.

### HasItems

`func (o *NameAwardData) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *NameAwardData) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *NameAwardData) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetNameAwardsHtml

`func (o *NameAwardData) GetNameAwardsHtml() string`

GetNameAwardsHtml returns the NameAwardsHtml field if non-nil, zero value otherwise.

### GetNameAwardsHtmlOk

`func (o *NameAwardData) GetNameAwardsHtmlOk() (*string, bool)`

GetNameAwardsHtmlOk returns a tuple with the NameAwardsHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameAwardsHtml

`func (o *NameAwardData) SetNameAwardsHtml(v string)`

SetNameAwardsHtml sets NameAwardsHtml field to given value.

### HasNameAwardsHtml

`func (o *NameAwardData) HasNameAwardsHtml() bool`

HasNameAwardsHtml returns a boolean if a field has been set.

### SetNameAwardsHtmlNil

`func (o *NameAwardData) SetNameAwardsHtmlNil(b bool)`

 SetNameAwardsHtmlNil sets the value for NameAwardsHtml to be an explicit nil

### UnsetNameAwardsHtml
`func (o *NameAwardData) UnsetNameAwardsHtml()`

UnsetNameAwardsHtml ensures that no value is present for NameAwardsHtml, not even an explicit nil
### GetErrorMessage

`func (o *NameAwardData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *NameAwardData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *NameAwardData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *NameAwardData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *NameAwardData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *NameAwardData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


