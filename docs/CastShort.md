# CastShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | Pointer to **NullableString** |  | [optional] 
**Items** | Pointer to [**[]CastShortItem**](CastShortItem.md) |  | [optional] 

## Methods

### NewCastShort

`func NewCastShort() *CastShort`

NewCastShort instantiates a new CastShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastShortWithDefaults

`func NewCastShortWithDefaults() *CastShort`

NewCastShortWithDefaults instantiates a new CastShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *CastShort) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *CastShort) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *CastShort) SetJob(v string)`

SetJob sets Job field to given value.

### HasJob

`func (o *CastShort) HasJob() bool`

HasJob returns a boolean if a field has been set.

### SetJobNil

`func (o *CastShort) SetJobNil(b bool)`

 SetJobNil sets the value for Job to be an explicit nil

### UnsetJob
`func (o *CastShort) UnsetJob()`

UnsetJob ensures that no value is present for Job, not even an explicit nil
### GetItems

`func (o *CastShort) GetItems() []CastShortItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CastShort) GetItemsOk() (*[]CastShortItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CastShort) SetItems(v []CastShortItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *CastShort) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *CastShort) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *CastShort) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


