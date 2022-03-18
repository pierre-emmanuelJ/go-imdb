# FAQDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Question** | Pointer to **NullableString** |  | [optional] 
**Answer** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFAQDetail

`func NewFAQDetail() *FAQDetail`

NewFAQDetail instantiates a new FAQDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFAQDetailWithDefaults

`func NewFAQDetailWithDefaults() *FAQDetail`

NewFAQDetailWithDefaults instantiates a new FAQDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestion

`func (o *FAQDetail) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *FAQDetail) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *FAQDetail) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *FAQDetail) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### SetQuestionNil

`func (o *FAQDetail) SetQuestionNil(b bool)`

 SetQuestionNil sets the value for Question to be an explicit nil

### UnsetQuestion
`func (o *FAQDetail) UnsetQuestion()`

UnsetQuestion ensures that no value is present for Question, not even an explicit nil
### GetAnswer

`func (o *FAQDetail) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *FAQDetail) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *FAQDetail) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *FAQDetail) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### SetAnswerNil

`func (o *FAQDetail) SetAnswerNil(b bool)`

 SetAnswerNil sets the value for Answer to be an explicit nil

### UnsetAnswer
`func (o *FAQDetail) UnsetAnswer()`

UnsetAnswer ensures that no value is present for Answer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


