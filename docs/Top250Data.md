# Top250Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Top250DataDetail**](Top250DataDetail.md) |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTop250Data

`func NewTop250Data() *Top250Data`

NewTop250Data instantiates a new Top250Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTop250DataWithDefaults

`func NewTop250DataWithDefaults() *Top250Data`

NewTop250DataWithDefaults instantiates a new Top250Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *Top250Data) GetItems() []Top250DataDetail`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Top250Data) GetItemsOk() (*[]Top250DataDetail, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Top250Data) SetItems(v []Top250DataDetail)`

SetItems sets Items field to given value.

### HasItems

`func (o *Top250Data) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *Top250Data) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *Top250Data) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetErrorMessage

`func (o *Top250Data) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Top250Data) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Top250Data) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Top250Data) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *Top250Data) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *Top250Data) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


