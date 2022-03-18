# SearchData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchType** | **string** |  | 
**Expression** | **string** |  | 
**Results** | Pointer to [**[]SearchResult**](SearchResult.md) |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSearchData

`func NewSearchData(searchType string, expression string, ) *SearchData`

NewSearchData instantiates a new SearchData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchDataWithDefaults

`func NewSearchDataWithDefaults() *SearchData`

NewSearchDataWithDefaults instantiates a new SearchData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchType

`func (o *SearchData) GetSearchType() string`

GetSearchType returns the SearchType field if non-nil, zero value otherwise.

### GetSearchTypeOk

`func (o *SearchData) GetSearchTypeOk() (*string, bool)`

GetSearchTypeOk returns a tuple with the SearchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchType

`func (o *SearchData) SetSearchType(v string)`

SetSearchType sets SearchType field to given value.


### GetExpression

`func (o *SearchData) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *SearchData) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *SearchData) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetResults

`func (o *SearchData) GetResults() []SearchResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchData) GetResultsOk() (*[]SearchResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchData) SetResults(v []SearchResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *SearchData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *SearchData) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *SearchData) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil
### GetErrorMessage

`func (o *SearchData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SearchData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *SearchData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *SearchData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *SearchData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *SearchData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


