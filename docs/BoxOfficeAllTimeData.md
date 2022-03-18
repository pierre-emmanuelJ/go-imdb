# BoxOfficeAllTimeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]BoxOfficeAllTimeDataDetail**](BoxOfficeAllTimeDataDetail.md) |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBoxOfficeAllTimeData

`func NewBoxOfficeAllTimeData() *BoxOfficeAllTimeData`

NewBoxOfficeAllTimeData instantiates a new BoxOfficeAllTimeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoxOfficeAllTimeDataWithDefaults

`func NewBoxOfficeAllTimeDataWithDefaults() *BoxOfficeAllTimeData`

NewBoxOfficeAllTimeDataWithDefaults instantiates a new BoxOfficeAllTimeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *BoxOfficeAllTimeData) GetItems() []BoxOfficeAllTimeDataDetail`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BoxOfficeAllTimeData) GetItemsOk() (*[]BoxOfficeAllTimeDataDetail, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BoxOfficeAllTimeData) SetItems(v []BoxOfficeAllTimeDataDetail)`

SetItems sets Items field to given value.

### HasItems

`func (o *BoxOfficeAllTimeData) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *BoxOfficeAllTimeData) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *BoxOfficeAllTimeData) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetErrorMessage

`func (o *BoxOfficeAllTimeData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BoxOfficeAllTimeData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BoxOfficeAllTimeData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BoxOfficeAllTimeData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *BoxOfficeAllTimeData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *BoxOfficeAllTimeData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


