# KeywordData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keyword** | Pointer to **NullableString** |  | [optional] 
**Items** | Pointer to [**[]MovieShort**](MovieShort.md) |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewKeywordData

`func NewKeywordData() *KeywordData`

NewKeywordData instantiates a new KeywordData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeywordDataWithDefaults

`func NewKeywordDataWithDefaults() *KeywordData`

NewKeywordDataWithDefaults instantiates a new KeywordData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyword

`func (o *KeywordData) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *KeywordData) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *KeywordData) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *KeywordData) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### SetKeywordNil

`func (o *KeywordData) SetKeywordNil(b bool)`

 SetKeywordNil sets the value for Keyword to be an explicit nil

### UnsetKeyword
`func (o *KeywordData) UnsetKeyword()`

UnsetKeyword ensures that no value is present for Keyword, not even an explicit nil
### GetItems

`func (o *KeywordData) GetItems() []MovieShort`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *KeywordData) GetItemsOk() (*[]MovieShort, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *KeywordData) SetItems(v []MovieShort)`

SetItems sets Items field to given value.

### HasItems

`func (o *KeywordData) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *KeywordData) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *KeywordData) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetErrorMessage

`func (o *KeywordData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *KeywordData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *KeywordData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *KeywordData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *KeywordData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *KeywordData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


