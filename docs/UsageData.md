# UsageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Maximum** | Pointer to **int32** |  | [optional] 
**Account** | Pointer to **NullableString** |  | [optional] 
**ExpireDate** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUsageData

`func NewUsageData() *UsageData`

NewUsageData instantiates a new UsageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageDataWithDefaults

`func NewUsageDataWithDefaults() *UsageData`

NewUsageDataWithDefaults instantiates a new UsageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UsageData) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UsageData) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UsageData) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UsageData) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetMaximum

`func (o *UsageData) GetMaximum() int32`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *UsageData) GetMaximumOk() (*int32, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *UsageData) SetMaximum(v int32)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *UsageData) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### GetAccount

`func (o *UsageData) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UsageData) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UsageData) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UsageData) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *UsageData) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *UsageData) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetExpireDate

`func (o *UsageData) GetExpireDate() string`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *UsageData) GetExpireDateOk() (*string, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *UsageData) SetExpireDate(v string)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *UsageData) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### SetExpireDateNil

`func (o *UsageData) SetExpireDateNil(b bool)`

 SetExpireDateNil sets the value for ExpireDate to be an explicit nil

### UnsetExpireDate
`func (o *UsageData) UnsetExpireDate()`

UnsetExpireDate ensures that no value is present for ExpireDate, not even an explicit nil
### GetErrorMessage

`func (o *UsageData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *UsageData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *UsageData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *UsageData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *UsageData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *UsageData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


